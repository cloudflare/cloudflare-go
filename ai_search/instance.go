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

// InstanceService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceService] method instead.
type InstanceService struct {
	Options []option.RequestOption
	Jobs    *InstanceJobService
}

// NewInstanceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceService(opts ...option.RequestOption) (r *InstanceService) {
	r = &InstanceService{}
	r.Options = opts
	r.Jobs = NewInstanceJobService(opts...)
	return
}

// Create a new AI Search instance with the given configuration.
func (r *InstanceService) New(ctx context.Context, params InstanceNewParams, opts ...option.RequestOption) (res *InstanceNewResponse, err error) {
	var env InstanceNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Update the configuration of an AI Search instance.
func (r *InstanceService) Update(ctx context.Context, id string, params InstanceUpdateParams, opts ...option.RequestOption) (res *InstanceUpdateResponse, err error) {
	var env InstanceUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances/%s", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List all AI Search instances in the account.
func (r *InstanceService) List(ctx context.Context, params InstanceListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[InstanceListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
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

// List all AI Search instances in the account.
func (r *InstanceService) ListAutoPaging(ctx context.Context, params InstanceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[InstanceListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Permanently delete an AI Search instance and all its indexed data.
func (r *InstanceService) Delete(ctx context.Context, id string, body InstanceDeleteParams, opts ...option.RequestOption) (res *InstanceDeleteResponse, err error) {
	var env InstanceDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances/%s", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Performs a chat completion request against an AI Search instance, using indexed
// content as context for generating responses.
func (r *InstanceService) ChatCompletions(ctx context.Context, id string, params InstanceChatCompletionsParams, opts ...option.RequestOption) (res *InstanceChatCompletionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances/%s/chat/completions", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve the configuration and status of an AI Search instance.
func (r *InstanceService) Read(ctx context.Context, id string, query InstanceReadParams, opts ...option.RequestOption) (res *InstanceReadResponse, err error) {
	var env InstanceReadResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances/%s", query.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Executes a semantic search query against an AI Search instance to find relevant
// indexed content.
func (r *InstanceService) Search(ctx context.Context, id string, params InstanceSearchParams, opts ...option.RequestOption) (res *InstanceSearchResponse, err error) {
	var env InstanceSearchResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances/%s/search", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve usage and indexing statistics for an AI Search instance.
func (r *InstanceService) Stats(ctx context.Context, id string, query InstanceStatsParams, opts ...option.RequestOption) (res *InstanceStatsResponse, err error) {
	var env InstanceStatsResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances/%s/stats", query.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type InstanceNewResponse struct {
	// AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.
	ID             string                            `json:"id" api:"required"`
	CreatedAt      time.Time                         `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt     time.Time                         `json:"modified_at" api:"required" format:"date-time"`
	AIGatewayID    string                            `json:"ai_gateway_id" api:"nullable"`
	AISearchModel  string                            `json:"ai_search_model" api:"nullable"`
	Cache          bool                              `json:"cache"`
	CacheThreshold InstanceNewResponseCacheThreshold `json:"cache_threshold"`
	// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
	// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
	// (72h), 518400 (6d).
	CacheTTL       InstanceNewResponseCacheTTL         `json:"cache_ttl"`
	ChunkOverlap   int64                               `json:"chunk_overlap"`
	ChunkSize      int64                               `json:"chunk_size"`
	CreatedBy      string                              `json:"created_by" api:"nullable"`
	CustomMetadata []InstanceNewResponseCustomMetadata `json:"custom_metadata"`
	EmbeddingModel string                              `json:"embedding_model" api:"nullable"`
	Enable         bool                                `json:"enable"`
	EngineVersion  float64                             `json:"engine_version"`
	FusionMethod   InstanceNewResponseFusionMethod     `json:"fusion_method"`
	// Deprecated — use index_method instead.
	//
	// Deprecated: deprecated
	HybridSearchEnabled bool `json:"hybrid_search_enabled"`
	// Controls which storage backends are used during indexing. Defaults to
	// vector-only.
	IndexMethod          InstanceNewResponseIndexMethod          `json:"index_method"`
	IndexingOptions      InstanceNewResponseIndexingOptions      `json:"indexing_options" api:"nullable"`
	LastActivity         time.Time                               `json:"last_activity" api:"nullable" format:"date-time"`
	MaxNumResults        int64                                   `json:"max_num_results"`
	Metadata             InstanceNewResponseMetadata             `json:"metadata"`
	ModifiedBy           string                                  `json:"modified_by" api:"nullable"`
	Namespace            string                                  `json:"namespace" api:"nullable"`
	Paused               bool                                    `json:"paused"`
	PublicEndpointID     string                                  `json:"public_endpoint_id" api:"nullable"`
	PublicEndpointParams InstanceNewResponsePublicEndpointParams `json:"public_endpoint_params"`
	Reranking            bool                                    `json:"reranking"`
	RerankingModel       string                                  `json:"reranking_model" api:"nullable"`
	RetrievalOptions     InstanceNewResponseRetrievalOptions     `json:"retrieval_options" api:"nullable"`
	RewriteModel         string                                  `json:"rewrite_model" api:"nullable"`
	RewriteQuery         bool                                    `json:"rewrite_query"`
	ScoreThreshold       float64                                 `json:"score_threshold"`
	Source               string                                  `json:"source" api:"nullable"`
	SourceParams         InstanceNewResponseSourceParams         `json:"source_params" api:"nullable"`
	Status               string                                  `json:"status"`
	// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
	// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
	SyncInterval InstanceNewResponseSyncInterval `json:"sync_interval"`
	TokenID      string                          `json:"token_id" format:"uuid"`
	Type         InstanceNewResponseType         `json:"type" api:"nullable"`
	JSON         instanceNewResponseJSON         `json:"-"`
}

// instanceNewResponseJSON contains the JSON metadata for the struct
// [InstanceNewResponse]
type instanceNewResponseJSON struct {
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

func (r *InstanceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseJSON) RawJSON() string {
	return r.raw
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

// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
// (72h), 518400 (6d).
type InstanceNewResponseCacheTTL float64

const (
	InstanceNewResponseCacheTTL600    InstanceNewResponseCacheTTL = 600
	InstanceNewResponseCacheTTL1800   InstanceNewResponseCacheTTL = 1800
	InstanceNewResponseCacheTTL3600   InstanceNewResponseCacheTTL = 3600
	InstanceNewResponseCacheTTL7200   InstanceNewResponseCacheTTL = 7200
	InstanceNewResponseCacheTTL21600  InstanceNewResponseCacheTTL = 21600
	InstanceNewResponseCacheTTL43200  InstanceNewResponseCacheTTL = 43200
	InstanceNewResponseCacheTTL86400  InstanceNewResponseCacheTTL = 86400
	InstanceNewResponseCacheTTL172800 InstanceNewResponseCacheTTL = 172800
	InstanceNewResponseCacheTTL259200 InstanceNewResponseCacheTTL = 259200
	InstanceNewResponseCacheTTL518400 InstanceNewResponseCacheTTL = 518400
)

func (r InstanceNewResponseCacheTTL) IsKnown() bool {
	switch r {
	case InstanceNewResponseCacheTTL600, InstanceNewResponseCacheTTL1800, InstanceNewResponseCacheTTL3600, InstanceNewResponseCacheTTL7200, InstanceNewResponseCacheTTL21600, InstanceNewResponseCacheTTL43200, InstanceNewResponseCacheTTL86400, InstanceNewResponseCacheTTL172800, InstanceNewResponseCacheTTL259200, InstanceNewResponseCacheTTL518400:
		return true
	}
	return false
}

type InstanceNewResponseCustomMetadata struct {
	DataType  InstanceNewResponseCustomMetadataDataType `json:"data_type" api:"required"`
	FieldName string                                    `json:"field_name" api:"required"`
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
	InstanceNewResponseCustomMetadataDataTypeText     InstanceNewResponseCustomMetadataDataType = "text"
	InstanceNewResponseCustomMetadataDataTypeNumber   InstanceNewResponseCustomMetadataDataType = "number"
	InstanceNewResponseCustomMetadataDataTypeBoolean  InstanceNewResponseCustomMetadataDataType = "boolean"
	InstanceNewResponseCustomMetadataDataTypeDatetime InstanceNewResponseCustomMetadataDataType = "datetime"
)

func (r InstanceNewResponseCustomMetadataDataType) IsKnown() bool {
	switch r {
	case InstanceNewResponseCustomMetadataDataTypeText, InstanceNewResponseCustomMetadataDataTypeNumber, InstanceNewResponseCustomMetadataDataTypeBoolean, InstanceNewResponseCustomMetadataDataTypeDatetime:
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

// Controls which storage backends are used during indexing. Defaults to
// vector-only.
type InstanceNewResponseIndexMethod struct {
	// Enable keyword (BM25) storage backend.
	Keyword bool `json:"keyword" api:"required"`
	// Enable vector (embedding) storage backend.
	Vector bool                               `json:"vector" api:"required"`
	JSON   instanceNewResponseIndexMethodJSON `json:"-"`
}

// instanceNewResponseIndexMethodJSON contains the JSON metadata for the struct
// [InstanceNewResponseIndexMethod]
type instanceNewResponseIndexMethodJSON struct {
	Keyword     apijson.Field
	Vector      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceNewResponseIndexMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseIndexMethodJSON) RawJSON() string {
	return r.raw
}

type InstanceNewResponseIndexingOptions struct {
	// Tokenizer used for keyword search indexing. porter provides word-level
	// tokenization with Porter stemming (good for natural language queries). trigram
	// enables character-level substring matching (good for partial matches, code,
	// identifiers). Changing this triggers a full re-index. Defaults to porter.
	KeywordTokenizer InstanceNewResponseIndexingOptionsKeywordTokenizer `json:"keyword_tokenizer"`
	JSON             instanceNewResponseIndexingOptionsJSON             `json:"-"`
}

// instanceNewResponseIndexingOptionsJSON contains the JSON metadata for the struct
// [InstanceNewResponseIndexingOptions]
type instanceNewResponseIndexingOptionsJSON struct {
	KeywordTokenizer apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InstanceNewResponseIndexingOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseIndexingOptionsJSON) RawJSON() string {
	return r.raw
}

// Tokenizer used for keyword search indexing. porter provides word-level
// tokenization with Porter stemming (good for natural language queries). trigram
// enables character-level substring matching (good for partial matches, code,
// identifiers). Changing this triggers a full re-index. Defaults to porter.
type InstanceNewResponseIndexingOptionsKeywordTokenizer string

const (
	InstanceNewResponseIndexingOptionsKeywordTokenizerPorter  InstanceNewResponseIndexingOptionsKeywordTokenizer = "porter"
	InstanceNewResponseIndexingOptionsKeywordTokenizerTrigram InstanceNewResponseIndexingOptionsKeywordTokenizer = "trigram"
)

func (r InstanceNewResponseIndexingOptionsKeywordTokenizer) IsKnown() bool {
	switch r {
	case InstanceNewResponseIndexingOptionsKeywordTokenizerPorter, InstanceNewResponseIndexingOptionsKeywordTokenizerTrigram:
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
	DefaultDomainEnabled bool                                                  `json:"default_domain_enabled"`
	Enabled              bool                                                  `json:"enabled"`
	Mcp                  InstanceNewResponsePublicEndpointParamsMcp            `json:"mcp"`
	RateLimit            InstanceNewResponsePublicEndpointParamsRateLimit      `json:"rate_limit"`
	SearchEndpoint       InstanceNewResponsePublicEndpointParamsSearchEndpoint `json:"search_endpoint"`
	JSON                 instanceNewResponsePublicEndpointParamsJSON           `json:"-"`
}

// instanceNewResponsePublicEndpointParamsJSON contains the JSON metadata for the
// struct [InstanceNewResponsePublicEndpointParams]
type instanceNewResponsePublicEndpointParamsJSON struct {
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

type InstanceNewResponseRetrievalOptions struct {
	// Metadata fields to boost search results by. Each entry specifies a metadata
	// field and an optional direction. Direction defaults to 'asc' for
	// numeric/datetime fields and 'exists' for text/boolean fields. Fields must match
	// 'timestamp' or a defined custom_metadata field.
	BoostBy []InstanceNewResponseRetrievalOptionsBoostBy `json:"boost_by"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. When omitted on an
	// update, the existing stored value is preserved; when never set, search falls
	// back to 'and'.
	KeywordMatchMode InstanceNewResponseRetrievalOptionsKeywordMatchMode `json:"keyword_match_mode"`
	JSON             instanceNewResponseRetrievalOptionsJSON             `json:"-"`
}

// instanceNewResponseRetrievalOptionsJSON contains the JSON metadata for the
// struct [InstanceNewResponseRetrievalOptions]
type instanceNewResponseRetrievalOptionsJSON struct {
	BoostBy          apijson.Field
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

type InstanceNewResponseRetrievalOptionsBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support all four directions
	// (asc, desc, exists, not_exists); text/boolean fields only support
	// exists/not_exists.
	Field string `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction InstanceNewResponseRetrievalOptionsBoostByDirection `json:"direction"`
	JSON      instanceNewResponseRetrievalOptionsBoostByJSON      `json:"-"`
}

// instanceNewResponseRetrievalOptionsBoostByJSON contains the JSON metadata for
// the struct [InstanceNewResponseRetrievalOptionsBoostBy]
type instanceNewResponseRetrievalOptionsBoostByJSON struct {
	Field       apijson.Field
	Direction   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceNewResponseRetrievalOptionsBoostBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseRetrievalOptionsBoostByJSON) RawJSON() string {
	return r.raw
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
// for numeric/datetime fields, 'exists' for text/boolean fields.
type InstanceNewResponseRetrievalOptionsBoostByDirection string

const (
	InstanceNewResponseRetrievalOptionsBoostByDirectionAsc       InstanceNewResponseRetrievalOptionsBoostByDirection = "asc"
	InstanceNewResponseRetrievalOptionsBoostByDirectionDesc      InstanceNewResponseRetrievalOptionsBoostByDirection = "desc"
	InstanceNewResponseRetrievalOptionsBoostByDirectionExists    InstanceNewResponseRetrievalOptionsBoostByDirection = "exists"
	InstanceNewResponseRetrievalOptionsBoostByDirectionNotExists InstanceNewResponseRetrievalOptionsBoostByDirection = "not_exists"
)

func (r InstanceNewResponseRetrievalOptionsBoostByDirection) IsKnown() bool {
	switch r {
	case InstanceNewResponseRetrievalOptionsBoostByDirectionAsc, InstanceNewResponseRetrievalOptionsBoostByDirectionDesc, InstanceNewResponseRetrievalOptionsBoostByDirectionExists, InstanceNewResponseRetrievalOptionsBoostByDirectionNotExists:
		return true
	}
	return false
}

// Controls which documents are candidates for BM25 scoring. 'and' restricts
// candidates to documents containing all query terms; 'or' includes any document
// containing at least one term, ranked by BM25 relevance. When omitted on an
// update, the existing stored value is preserved; when never set, search falls
// back to 'and'.
type InstanceNewResponseRetrievalOptionsKeywordMatchMode string

const (
	InstanceNewResponseRetrievalOptionsKeywordMatchModeAnd InstanceNewResponseRetrievalOptionsKeywordMatchMode = "and"
	InstanceNewResponseRetrievalOptionsKeywordMatchModeOr  InstanceNewResponseRetrievalOptionsKeywordMatchMode = "or"
)

func (r InstanceNewResponseRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case InstanceNewResponseRetrievalOptionsKeywordMatchModeAnd, InstanceNewResponseRetrievalOptionsKeywordMatchModeOr:
		return true
	}
	return false
}

type InstanceNewResponseSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced). Most accounts are limited to 10
	// rules; contact support to raise it.
	ExcludeItems []string `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post). Most accounts are limited to 10 rules; contact
	// support to raise it.
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
	// Options for parse_type 'discover', where Browser Run discovers URLs by link
	// following and sitemaps. Ignored for 'sitemap'.
	DiscoverOptions InstanceNewResponseSourceParamsWebCrawlerDiscoverOptions `json:"discover_options"`
	ParseOptions    InstanceNewResponseSourceParamsWebCrawlerParseOptions    `json:"parse_options"`
	// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
	// recursively and requires the source to be a Verified zone on this account.
	ParseType InstanceNewResponseSourceParamsWebCrawlerParseType `json:"parse_type"`
	JSON      instanceNewResponseSourceParamsWebCrawlerJSON      `json:"-"`
}

// instanceNewResponseSourceParamsWebCrawlerJSON contains the JSON metadata for the
// struct [InstanceNewResponseSourceParamsWebCrawler]
type instanceNewResponseSourceParamsWebCrawlerJSON struct {
	DiscoverOptions apijson.Field
	ParseOptions    apijson.Field
	ParseType       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InstanceNewResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

// Options for parse_type 'discover', where Browser Run discovers URLs by link
// following and sitemaps. Ignored for 'sitemap'.
type InstanceNewResponseSourceParamsWebCrawlerDiscoverOptions struct {
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
	Source InstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSource `json:"source"`
	JSON   instanceNewResponseSourceParamsWebCrawlerDiscoverOptionsJSON   `json:"-"`
}

// instanceNewResponseSourceParamsWebCrawlerDiscoverOptionsJSON contains the JSON
// metadata for the struct
// [InstanceNewResponseSourceParamsWebCrawlerDiscoverOptions]
type instanceNewResponseSourceParamsWebCrawlerDiscoverOptionsJSON struct {
	Depth                apijson.Field
	IncludeExternalLinks apijson.Field
	IncludeSubdomains    apijson.Field
	Limit                apijson.Field
	MaxAge               apijson.Field
	Source               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InstanceNewResponseSourceParamsWebCrawlerDiscoverOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseSourceParamsWebCrawlerDiscoverOptionsJSON) RawJSON() string {
	return r.raw
}

// Where the crawler looks for URLs: 'sitemaps' reads sitemap XML only, 'links'
// follows page links only, 'all' does both.
type InstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSource string

const (
	InstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll      InstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSource = "all"
	InstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps InstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSource = "sitemaps"
	InstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks    InstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSource = "links"
)

func (r InstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSource) IsKnown() bool {
	switch r {
	case InstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll, InstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps, InstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks:
		return true
	}
	return false
}

type InstanceNewResponseSourceParamsWebCrawlerParseOptions struct {
	// List of path-to-selector mappings for extracting specific content from crawled
	// pages. Each entry pairs a URL glob pattern with a CSS selector. The first
	// matching path wins. Only the matched HTML fragment is stored and indexed. Omit
	// the field to disable content selection — empty arrays are rejected.
	ContentSelector []InstanceNewResponseSourceParamsWebCrawlerParseOptionsContentSelector `json:"content_selector"`
	// Up to 5 custom HTTP headers sent with each crawl request. Names must be RFC-7230
	// token characters (no spaces, colons, or control characters); values must be
	// HTAB + printable ASCII (no CR/LF).
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
	ContentSelector     apijson.Field
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

type InstanceNewResponseSourceParamsWebCrawlerParseOptionsContentSelector struct {
	// Glob pattern to match against the page URL path. Uses standard glob syntax: \*
	// matches within a segment, \*\* crosses directories.
	Path string `json:"path" api:"required"`
	// CSS selector to extract content from pages matching the path pattern. Must not
	// contain disallowed characters (;, `, $, {, }, \). Must target a single element;
	// if multiple elements match, the selector is ignored and the full page is used.
	Selector string                                                                   `json:"selector" api:"required"`
	JSON     instanceNewResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON `json:"-"`
}

// instanceNewResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON
// contains the JSON metadata for the struct
// [InstanceNewResponseSourceParamsWebCrawlerParseOptionsContentSelector]
type instanceNewResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON struct {
	Path        apijson.Field
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceNewResponseSourceParamsWebCrawlerParseOptionsContentSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON) RawJSON() string {
	return r.raw
}

// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
// recursively and requires the source to be a Verified zone on this account.
type InstanceNewResponseSourceParamsWebCrawlerParseType string

const (
	InstanceNewResponseSourceParamsWebCrawlerParseTypeSitemap  InstanceNewResponseSourceParamsWebCrawlerParseType = "sitemap"
	InstanceNewResponseSourceParamsWebCrawlerParseTypeDiscover InstanceNewResponseSourceParamsWebCrawlerParseType = "discover"
)

func (r InstanceNewResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case InstanceNewResponseSourceParamsWebCrawlerParseTypeSitemap, InstanceNewResponseSourceParamsWebCrawlerParseTypeDiscover:
		return true
	}
	return false
}

// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
type InstanceNewResponseSyncInterval float64

const (
	InstanceNewResponseSyncInterval900   InstanceNewResponseSyncInterval = 900
	InstanceNewResponseSyncInterval1800  InstanceNewResponseSyncInterval = 1800
	InstanceNewResponseSyncInterval3600  InstanceNewResponseSyncInterval = 3600
	InstanceNewResponseSyncInterval7200  InstanceNewResponseSyncInterval = 7200
	InstanceNewResponseSyncInterval14400 InstanceNewResponseSyncInterval = 14400
	InstanceNewResponseSyncInterval21600 InstanceNewResponseSyncInterval = 21600
	InstanceNewResponseSyncInterval43200 InstanceNewResponseSyncInterval = 43200
	InstanceNewResponseSyncInterval86400 InstanceNewResponseSyncInterval = 86400
)

func (r InstanceNewResponseSyncInterval) IsKnown() bool {
	switch r {
	case InstanceNewResponseSyncInterval900, InstanceNewResponseSyncInterval1800, InstanceNewResponseSyncInterval3600, InstanceNewResponseSyncInterval7200, InstanceNewResponseSyncInterval14400, InstanceNewResponseSyncInterval21600, InstanceNewResponseSyncInterval43200, InstanceNewResponseSyncInterval86400:
		return true
	}
	return false
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

type InstanceUpdateResponse struct {
	// AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.
	ID             string                               `json:"id" api:"required"`
	CreatedAt      time.Time                            `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt     time.Time                            `json:"modified_at" api:"required" format:"date-time"`
	AIGatewayID    string                               `json:"ai_gateway_id" api:"nullable"`
	AISearchModel  string                               `json:"ai_search_model" api:"nullable"`
	Cache          bool                                 `json:"cache"`
	CacheThreshold InstanceUpdateResponseCacheThreshold `json:"cache_threshold"`
	// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
	// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
	// (72h), 518400 (6d).
	CacheTTL       InstanceUpdateResponseCacheTTL         `json:"cache_ttl"`
	ChunkOverlap   int64                                  `json:"chunk_overlap"`
	ChunkSize      int64                                  `json:"chunk_size"`
	CreatedBy      string                                 `json:"created_by" api:"nullable"`
	CustomMetadata []InstanceUpdateResponseCustomMetadata `json:"custom_metadata"`
	EmbeddingModel string                                 `json:"embedding_model" api:"nullable"`
	Enable         bool                                   `json:"enable"`
	EngineVersion  float64                                `json:"engine_version"`
	FusionMethod   InstanceUpdateResponseFusionMethod     `json:"fusion_method"`
	// Deprecated — use index_method instead.
	//
	// Deprecated: deprecated
	HybridSearchEnabled bool `json:"hybrid_search_enabled"`
	// Controls which storage backends are used during indexing. Defaults to
	// vector-only.
	IndexMethod          InstanceUpdateResponseIndexMethod          `json:"index_method"`
	IndexingOptions      InstanceUpdateResponseIndexingOptions      `json:"indexing_options" api:"nullable"`
	LastActivity         time.Time                                  `json:"last_activity" api:"nullable" format:"date-time"`
	MaxNumResults        int64                                      `json:"max_num_results"`
	Metadata             InstanceUpdateResponseMetadata             `json:"metadata"`
	ModifiedBy           string                                     `json:"modified_by" api:"nullable"`
	Namespace            string                                     `json:"namespace" api:"nullable"`
	Paused               bool                                       `json:"paused"`
	PublicEndpointID     string                                     `json:"public_endpoint_id" api:"nullable"`
	PublicEndpointParams InstanceUpdateResponsePublicEndpointParams `json:"public_endpoint_params"`
	Reranking            bool                                       `json:"reranking"`
	RerankingModel       string                                     `json:"reranking_model" api:"nullable"`
	RetrievalOptions     InstanceUpdateResponseRetrievalOptions     `json:"retrieval_options" api:"nullable"`
	RewriteModel         string                                     `json:"rewrite_model" api:"nullable"`
	RewriteQuery         bool                                       `json:"rewrite_query"`
	ScoreThreshold       float64                                    `json:"score_threshold"`
	Source               string                                     `json:"source" api:"nullable"`
	SourceParams         InstanceUpdateResponseSourceParams         `json:"source_params" api:"nullable"`
	Status               string                                     `json:"status"`
	// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
	// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
	SyncInterval InstanceUpdateResponseSyncInterval `json:"sync_interval"`
	TokenID      string                             `json:"token_id" format:"uuid"`
	Type         InstanceUpdateResponseType         `json:"type" api:"nullable"`
	JSON         instanceUpdateResponseJSON         `json:"-"`
}

// instanceUpdateResponseJSON contains the JSON metadata for the struct
// [InstanceUpdateResponse]
type instanceUpdateResponseJSON struct {
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

func (r *InstanceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponseJSON) RawJSON() string {
	return r.raw
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

// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
// (72h), 518400 (6d).
type InstanceUpdateResponseCacheTTL float64

const (
	InstanceUpdateResponseCacheTTL600    InstanceUpdateResponseCacheTTL = 600
	InstanceUpdateResponseCacheTTL1800   InstanceUpdateResponseCacheTTL = 1800
	InstanceUpdateResponseCacheTTL3600   InstanceUpdateResponseCacheTTL = 3600
	InstanceUpdateResponseCacheTTL7200   InstanceUpdateResponseCacheTTL = 7200
	InstanceUpdateResponseCacheTTL21600  InstanceUpdateResponseCacheTTL = 21600
	InstanceUpdateResponseCacheTTL43200  InstanceUpdateResponseCacheTTL = 43200
	InstanceUpdateResponseCacheTTL86400  InstanceUpdateResponseCacheTTL = 86400
	InstanceUpdateResponseCacheTTL172800 InstanceUpdateResponseCacheTTL = 172800
	InstanceUpdateResponseCacheTTL259200 InstanceUpdateResponseCacheTTL = 259200
	InstanceUpdateResponseCacheTTL518400 InstanceUpdateResponseCacheTTL = 518400
)

func (r InstanceUpdateResponseCacheTTL) IsKnown() bool {
	switch r {
	case InstanceUpdateResponseCacheTTL600, InstanceUpdateResponseCacheTTL1800, InstanceUpdateResponseCacheTTL3600, InstanceUpdateResponseCacheTTL7200, InstanceUpdateResponseCacheTTL21600, InstanceUpdateResponseCacheTTL43200, InstanceUpdateResponseCacheTTL86400, InstanceUpdateResponseCacheTTL172800, InstanceUpdateResponseCacheTTL259200, InstanceUpdateResponseCacheTTL518400:
		return true
	}
	return false
}

type InstanceUpdateResponseCustomMetadata struct {
	DataType  InstanceUpdateResponseCustomMetadataDataType `json:"data_type" api:"required"`
	FieldName string                                       `json:"field_name" api:"required"`
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
	InstanceUpdateResponseCustomMetadataDataTypeText     InstanceUpdateResponseCustomMetadataDataType = "text"
	InstanceUpdateResponseCustomMetadataDataTypeNumber   InstanceUpdateResponseCustomMetadataDataType = "number"
	InstanceUpdateResponseCustomMetadataDataTypeBoolean  InstanceUpdateResponseCustomMetadataDataType = "boolean"
	InstanceUpdateResponseCustomMetadataDataTypeDatetime InstanceUpdateResponseCustomMetadataDataType = "datetime"
)

func (r InstanceUpdateResponseCustomMetadataDataType) IsKnown() bool {
	switch r {
	case InstanceUpdateResponseCustomMetadataDataTypeText, InstanceUpdateResponseCustomMetadataDataTypeNumber, InstanceUpdateResponseCustomMetadataDataTypeBoolean, InstanceUpdateResponseCustomMetadataDataTypeDatetime:
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

// Controls which storage backends are used during indexing. Defaults to
// vector-only.
type InstanceUpdateResponseIndexMethod struct {
	// Enable keyword (BM25) storage backend.
	Keyword bool `json:"keyword" api:"required"`
	// Enable vector (embedding) storage backend.
	Vector bool                                  `json:"vector" api:"required"`
	JSON   instanceUpdateResponseIndexMethodJSON `json:"-"`
}

// instanceUpdateResponseIndexMethodJSON contains the JSON metadata for the struct
// [InstanceUpdateResponseIndexMethod]
type instanceUpdateResponseIndexMethodJSON struct {
	Keyword     apijson.Field
	Vector      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceUpdateResponseIndexMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponseIndexMethodJSON) RawJSON() string {
	return r.raw
}

type InstanceUpdateResponseIndexingOptions struct {
	// Tokenizer used for keyword search indexing. porter provides word-level
	// tokenization with Porter stemming (good for natural language queries). trigram
	// enables character-level substring matching (good for partial matches, code,
	// identifiers). Changing this triggers a full re-index. Defaults to porter.
	KeywordTokenizer InstanceUpdateResponseIndexingOptionsKeywordTokenizer `json:"keyword_tokenizer"`
	JSON             instanceUpdateResponseIndexingOptionsJSON             `json:"-"`
}

// instanceUpdateResponseIndexingOptionsJSON contains the JSON metadata for the
// struct [InstanceUpdateResponseIndexingOptions]
type instanceUpdateResponseIndexingOptionsJSON struct {
	KeywordTokenizer apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InstanceUpdateResponseIndexingOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponseIndexingOptionsJSON) RawJSON() string {
	return r.raw
}

// Tokenizer used for keyword search indexing. porter provides word-level
// tokenization with Porter stemming (good for natural language queries). trigram
// enables character-level substring matching (good for partial matches, code,
// identifiers). Changing this triggers a full re-index. Defaults to porter.
type InstanceUpdateResponseIndexingOptionsKeywordTokenizer string

const (
	InstanceUpdateResponseIndexingOptionsKeywordTokenizerPorter  InstanceUpdateResponseIndexingOptionsKeywordTokenizer = "porter"
	InstanceUpdateResponseIndexingOptionsKeywordTokenizerTrigram InstanceUpdateResponseIndexingOptionsKeywordTokenizer = "trigram"
)

func (r InstanceUpdateResponseIndexingOptionsKeywordTokenizer) IsKnown() bool {
	switch r {
	case InstanceUpdateResponseIndexingOptionsKeywordTokenizerPorter, InstanceUpdateResponseIndexingOptionsKeywordTokenizerTrigram:
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
	DefaultDomainEnabled bool                                                     `json:"default_domain_enabled"`
	Enabled              bool                                                     `json:"enabled"`
	Mcp                  InstanceUpdateResponsePublicEndpointParamsMcp            `json:"mcp"`
	RateLimit            InstanceUpdateResponsePublicEndpointParamsRateLimit      `json:"rate_limit"`
	SearchEndpoint       InstanceUpdateResponsePublicEndpointParamsSearchEndpoint `json:"search_endpoint"`
	JSON                 instanceUpdateResponsePublicEndpointParamsJSON           `json:"-"`
}

// instanceUpdateResponsePublicEndpointParamsJSON contains the JSON metadata for
// the struct [InstanceUpdateResponsePublicEndpointParams]
type instanceUpdateResponsePublicEndpointParamsJSON struct {
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

type InstanceUpdateResponseRetrievalOptions struct {
	// Metadata fields to boost search results by. Each entry specifies a metadata
	// field and an optional direction. Direction defaults to 'asc' for
	// numeric/datetime fields and 'exists' for text/boolean fields. Fields must match
	// 'timestamp' or a defined custom_metadata field.
	BoostBy []InstanceUpdateResponseRetrievalOptionsBoostBy `json:"boost_by"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. When omitted on an
	// update, the existing stored value is preserved; when never set, search falls
	// back to 'and'.
	KeywordMatchMode InstanceUpdateResponseRetrievalOptionsKeywordMatchMode `json:"keyword_match_mode"`
	JSON             instanceUpdateResponseRetrievalOptionsJSON             `json:"-"`
}

// instanceUpdateResponseRetrievalOptionsJSON contains the JSON metadata for the
// struct [InstanceUpdateResponseRetrievalOptions]
type instanceUpdateResponseRetrievalOptionsJSON struct {
	BoostBy          apijson.Field
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

type InstanceUpdateResponseRetrievalOptionsBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support all four directions
	// (asc, desc, exists, not_exists); text/boolean fields only support
	// exists/not_exists.
	Field string `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction InstanceUpdateResponseRetrievalOptionsBoostByDirection `json:"direction"`
	JSON      instanceUpdateResponseRetrievalOptionsBoostByJSON      `json:"-"`
}

// instanceUpdateResponseRetrievalOptionsBoostByJSON contains the JSON metadata for
// the struct [InstanceUpdateResponseRetrievalOptionsBoostBy]
type instanceUpdateResponseRetrievalOptionsBoostByJSON struct {
	Field       apijson.Field
	Direction   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceUpdateResponseRetrievalOptionsBoostBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponseRetrievalOptionsBoostByJSON) RawJSON() string {
	return r.raw
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
// for numeric/datetime fields, 'exists' for text/boolean fields.
type InstanceUpdateResponseRetrievalOptionsBoostByDirection string

const (
	InstanceUpdateResponseRetrievalOptionsBoostByDirectionAsc       InstanceUpdateResponseRetrievalOptionsBoostByDirection = "asc"
	InstanceUpdateResponseRetrievalOptionsBoostByDirectionDesc      InstanceUpdateResponseRetrievalOptionsBoostByDirection = "desc"
	InstanceUpdateResponseRetrievalOptionsBoostByDirectionExists    InstanceUpdateResponseRetrievalOptionsBoostByDirection = "exists"
	InstanceUpdateResponseRetrievalOptionsBoostByDirectionNotExists InstanceUpdateResponseRetrievalOptionsBoostByDirection = "not_exists"
)

func (r InstanceUpdateResponseRetrievalOptionsBoostByDirection) IsKnown() bool {
	switch r {
	case InstanceUpdateResponseRetrievalOptionsBoostByDirectionAsc, InstanceUpdateResponseRetrievalOptionsBoostByDirectionDesc, InstanceUpdateResponseRetrievalOptionsBoostByDirectionExists, InstanceUpdateResponseRetrievalOptionsBoostByDirectionNotExists:
		return true
	}
	return false
}

// Controls which documents are candidates for BM25 scoring. 'and' restricts
// candidates to documents containing all query terms; 'or' includes any document
// containing at least one term, ranked by BM25 relevance. When omitted on an
// update, the existing stored value is preserved; when never set, search falls
// back to 'and'.
type InstanceUpdateResponseRetrievalOptionsKeywordMatchMode string

const (
	InstanceUpdateResponseRetrievalOptionsKeywordMatchModeAnd InstanceUpdateResponseRetrievalOptionsKeywordMatchMode = "and"
	InstanceUpdateResponseRetrievalOptionsKeywordMatchModeOr  InstanceUpdateResponseRetrievalOptionsKeywordMatchMode = "or"
)

func (r InstanceUpdateResponseRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case InstanceUpdateResponseRetrievalOptionsKeywordMatchModeAnd, InstanceUpdateResponseRetrievalOptionsKeywordMatchModeOr:
		return true
	}
	return false
}

type InstanceUpdateResponseSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced). Most accounts are limited to 10
	// rules; contact support to raise it.
	ExcludeItems []string `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post). Most accounts are limited to 10 rules; contact
	// support to raise it.
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
	// Options for parse_type 'discover', where Browser Run discovers URLs by link
	// following and sitemaps. Ignored for 'sitemap'.
	DiscoverOptions InstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptions `json:"discover_options"`
	ParseOptions    InstanceUpdateResponseSourceParamsWebCrawlerParseOptions    `json:"parse_options"`
	// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
	// recursively and requires the source to be a Verified zone on this account.
	ParseType InstanceUpdateResponseSourceParamsWebCrawlerParseType `json:"parse_type"`
	JSON      instanceUpdateResponseSourceParamsWebCrawlerJSON      `json:"-"`
}

// instanceUpdateResponseSourceParamsWebCrawlerJSON contains the JSON metadata for
// the struct [InstanceUpdateResponseSourceParamsWebCrawler]
type instanceUpdateResponseSourceParamsWebCrawlerJSON struct {
	DiscoverOptions apijson.Field
	ParseOptions    apijson.Field
	ParseType       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InstanceUpdateResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

// Options for parse_type 'discover', where Browser Run discovers URLs by link
// following and sitemaps. Ignored for 'sitemap'.
type InstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptions struct {
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
	Source InstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSource `json:"source"`
	JSON   instanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsJSON   `json:"-"`
}

// instanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsJSON contains the
// JSON metadata for the struct
// [InstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptions]
type instanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsJSON struct {
	Depth                apijson.Field
	IncludeExternalLinks apijson.Field
	IncludeSubdomains    apijson.Field
	Limit                apijson.Field
	MaxAge               apijson.Field
	Source               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsJSON) RawJSON() string {
	return r.raw
}

// Where the crawler looks for URLs: 'sitemaps' reads sitemap XML only, 'links'
// follows page links only, 'all' does both.
type InstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSource string

const (
	InstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll      InstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSource = "all"
	InstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps InstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSource = "sitemaps"
	InstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks    InstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSource = "links"
)

func (r InstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSource) IsKnown() bool {
	switch r {
	case InstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll, InstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps, InstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks:
		return true
	}
	return false
}

type InstanceUpdateResponseSourceParamsWebCrawlerParseOptions struct {
	// List of path-to-selector mappings for extracting specific content from crawled
	// pages. Each entry pairs a URL glob pattern with a CSS selector. The first
	// matching path wins. Only the matched HTML fragment is stored and indexed. Omit
	// the field to disable content selection — empty arrays are rejected.
	ContentSelector []InstanceUpdateResponseSourceParamsWebCrawlerParseOptionsContentSelector `json:"content_selector"`
	// Up to 5 custom HTTP headers sent with each crawl request. Names must be RFC-7230
	// token characters (no spaces, colons, or control characters); values must be
	// HTAB + printable ASCII (no CR/LF).
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
	ContentSelector     apijson.Field
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

type InstanceUpdateResponseSourceParamsWebCrawlerParseOptionsContentSelector struct {
	// Glob pattern to match against the page URL path. Uses standard glob syntax: \*
	// matches within a segment, \*\* crosses directories.
	Path string `json:"path" api:"required"`
	// CSS selector to extract content from pages matching the path pattern. Must not
	// contain disallowed characters (;, `, $, {, }, \). Must target a single element;
	// if multiple elements match, the selector is ignored and the full page is used.
	Selector string                                                                      `json:"selector" api:"required"`
	JSON     instanceUpdateResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON `json:"-"`
}

// instanceUpdateResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON
// contains the JSON metadata for the struct
// [InstanceUpdateResponseSourceParamsWebCrawlerParseOptionsContentSelector]
type instanceUpdateResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON struct {
	Path        apijson.Field
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceUpdateResponseSourceParamsWebCrawlerParseOptionsContentSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON) RawJSON() string {
	return r.raw
}

// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
// recursively and requires the source to be a Verified zone on this account.
type InstanceUpdateResponseSourceParamsWebCrawlerParseType string

const (
	InstanceUpdateResponseSourceParamsWebCrawlerParseTypeSitemap  InstanceUpdateResponseSourceParamsWebCrawlerParseType = "sitemap"
	InstanceUpdateResponseSourceParamsWebCrawlerParseTypeDiscover InstanceUpdateResponseSourceParamsWebCrawlerParseType = "discover"
)

func (r InstanceUpdateResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case InstanceUpdateResponseSourceParamsWebCrawlerParseTypeSitemap, InstanceUpdateResponseSourceParamsWebCrawlerParseTypeDiscover:
		return true
	}
	return false
}

// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
type InstanceUpdateResponseSyncInterval float64

const (
	InstanceUpdateResponseSyncInterval900   InstanceUpdateResponseSyncInterval = 900
	InstanceUpdateResponseSyncInterval1800  InstanceUpdateResponseSyncInterval = 1800
	InstanceUpdateResponseSyncInterval3600  InstanceUpdateResponseSyncInterval = 3600
	InstanceUpdateResponseSyncInterval7200  InstanceUpdateResponseSyncInterval = 7200
	InstanceUpdateResponseSyncInterval14400 InstanceUpdateResponseSyncInterval = 14400
	InstanceUpdateResponseSyncInterval21600 InstanceUpdateResponseSyncInterval = 21600
	InstanceUpdateResponseSyncInterval43200 InstanceUpdateResponseSyncInterval = 43200
	InstanceUpdateResponseSyncInterval86400 InstanceUpdateResponseSyncInterval = 86400
)

func (r InstanceUpdateResponseSyncInterval) IsKnown() bool {
	switch r {
	case InstanceUpdateResponseSyncInterval900, InstanceUpdateResponseSyncInterval1800, InstanceUpdateResponseSyncInterval3600, InstanceUpdateResponseSyncInterval7200, InstanceUpdateResponseSyncInterval14400, InstanceUpdateResponseSyncInterval21600, InstanceUpdateResponseSyncInterval43200, InstanceUpdateResponseSyncInterval86400:
		return true
	}
	return false
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

type InstanceListResponse struct {
	ID                             string                                   `json:"id" api:"required"`
	AIGatewayID                    string                                   `json:"ai_gateway_id" api:"required,nullable"`
	AISearchModel                  string                                   `json:"ai_search_model" api:"required,nullable"`
	Cache                          bool                                     `json:"cache" api:"required"`
	CacheThreshold                 InstanceListResponseCacheThreshold       `json:"cache_threshold" api:"required,nullable"`
	CacheTTL                       InstanceListResponseCacheTTL             `json:"cache_ttl" api:"required"`
	Chunk                          bool                                     `json:"chunk" api:"required"`
	ChunkOverlap                   float64                                  `json:"chunk_overlap" api:"required,nullable"`
	ChunkSize                      float64                                  `json:"chunk_size" api:"required,nullable"`
	CreatedAt                      time.Time                                `json:"created_at" api:"required" format:"date-time"`
	CreatedBy                      string                                   `json:"created_by" api:"required,nullable"`
	CustomMetadata                 []InstanceListResponseCustomMetadata     `json:"custom_metadata" api:"required,nullable"`
	EmbeddingModel                 string                                   `json:"embedding_model" api:"required,nullable"`
	Enable                         bool                                     `json:"enable" api:"required"`
	EngineVersion                  float64                                  `json:"engine_version" api:"required"`
	FusionMethod                   InstanceListResponseFusionMethod         `json:"fusion_method" api:"required"`
	HybridSearchEnabled            bool                                     `json:"hybrid_search_enabled" api:"required"`
	IndexMethod                    InstanceListResponseIndexMethod          `json:"index_method" api:"required"`
	IndexingOptions                InstanceListResponseIndexingOptions      `json:"indexing_options" api:"required,nullable"`
	LastActivity                   time.Time                                `json:"last_activity" api:"required,nullable" format:"date-time"`
	MaxNumResults                  float64                                  `json:"max_num_results" api:"required,nullable"`
	Metadata                       InstanceListResponseMetadata             `json:"metadata" api:"required,nullable"`
	ModifiedAt                     time.Time                                `json:"modified_at" api:"required" format:"date-time"`
	ModifiedBy                     string                                   `json:"modified_by" api:"required,nullable"`
	Namespace                      string                                   `json:"namespace" api:"required"`
	Paused                         bool                                     `json:"paused" api:"required"`
	PublicEndpointID               string                                   `json:"public_endpoint_id" api:"required,nullable"`
	PublicEndpointParams           InstanceListResponsePublicEndpointParams `json:"public_endpoint_params" api:"required,nullable"`
	Reranking                      bool                                     `json:"reranking" api:"required"`
	RerankingModel                 string                                   `json:"reranking_model" api:"required,nullable"`
	RetrievalOptions               InstanceListResponseRetrievalOptions     `json:"retrieval_options" api:"required,nullable"`
	RewriteModel                   string                                   `json:"rewrite_model" api:"required,nullable"`
	RewriteQuery                   bool                                     `json:"rewrite_query" api:"required"`
	ScoreThreshold                 float64                                  `json:"score_threshold" api:"required,nullable"`
	Source                         string                                   `json:"source" api:"required,nullable"`
	SourceParams                   InstanceListResponseSourceParams         `json:"source_params" api:"required,nullable"`
	Status                         string                                   `json:"status" api:"required"`
	Summarization                  bool                                     `json:"summarization" api:"required"`
	SummarizationModel             string                                   `json:"summarization_model" api:"required,nullable"`
	SyncInterval                   InstanceListResponseSyncInterval         `json:"sync_interval" api:"required"`
	SystemPromptAISearch           string                                   `json:"system_prompt_ai_search" api:"required,nullable"`
	SystemPromptIndexSummarization string                                   `json:"system_prompt_index_summarization" api:"required,nullable"`
	SystemPromptRewriteQuery       string                                   `json:"system_prompt_rewrite_query" api:"required,nullable"`
	TokenID                        string                                   `json:"token_id" api:"required,nullable"`
	Type                           InstanceListResponseType                 `json:"type" api:"required,nullable"`
	JSON                           instanceListResponseJSON                 `json:"-"`
}

// instanceListResponseJSON contains the JSON metadata for the struct
// [InstanceListResponse]
type instanceListResponseJSON struct {
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

func (r *InstanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponseJSON) RawJSON() string {
	return r.raw
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

type InstanceListResponseCacheTTL float64

const (
	InstanceListResponseCacheTTL600    InstanceListResponseCacheTTL = 600
	InstanceListResponseCacheTTL1800   InstanceListResponseCacheTTL = 1800
	InstanceListResponseCacheTTL3600   InstanceListResponseCacheTTL = 3600
	InstanceListResponseCacheTTL7200   InstanceListResponseCacheTTL = 7200
	InstanceListResponseCacheTTL21600  InstanceListResponseCacheTTL = 21600
	InstanceListResponseCacheTTL43200  InstanceListResponseCacheTTL = 43200
	InstanceListResponseCacheTTL86400  InstanceListResponseCacheTTL = 86400
	InstanceListResponseCacheTTL172800 InstanceListResponseCacheTTL = 172800
	InstanceListResponseCacheTTL259200 InstanceListResponseCacheTTL = 259200
	InstanceListResponseCacheTTL518400 InstanceListResponseCacheTTL = 518400
)

func (r InstanceListResponseCacheTTL) IsKnown() bool {
	switch r {
	case InstanceListResponseCacheTTL600, InstanceListResponseCacheTTL1800, InstanceListResponseCacheTTL3600, InstanceListResponseCacheTTL7200, InstanceListResponseCacheTTL21600, InstanceListResponseCacheTTL43200, InstanceListResponseCacheTTL86400, InstanceListResponseCacheTTL172800, InstanceListResponseCacheTTL259200, InstanceListResponseCacheTTL518400:
		return true
	}
	return false
}

type InstanceListResponseCustomMetadata struct {
	DataType  InstanceListResponseCustomMetadataDataType `json:"data_type" api:"required"`
	FieldName string                                     `json:"field_name" api:"required"`
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
	InstanceListResponseCustomMetadataDataTypeText     InstanceListResponseCustomMetadataDataType = "text"
	InstanceListResponseCustomMetadataDataTypeNumber   InstanceListResponseCustomMetadataDataType = "number"
	InstanceListResponseCustomMetadataDataTypeBoolean  InstanceListResponseCustomMetadataDataType = "boolean"
	InstanceListResponseCustomMetadataDataTypeDatetime InstanceListResponseCustomMetadataDataType = "datetime"
)

func (r InstanceListResponseCustomMetadataDataType) IsKnown() bool {
	switch r {
	case InstanceListResponseCustomMetadataDataTypeText, InstanceListResponseCustomMetadataDataTypeNumber, InstanceListResponseCustomMetadataDataTypeBoolean, InstanceListResponseCustomMetadataDataTypeDatetime:
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

type InstanceListResponseIndexMethod struct {
	Keyword     bool                                `json:"keyword" api:"required"`
	Vector      bool                                `json:"vector" api:"required"`
	ExtraFields map[string]interface{}              `json:"-" api:"extrafields"`
	JSON        instanceListResponseIndexMethodJSON `json:"-"`
}

// instanceListResponseIndexMethodJSON contains the JSON metadata for the struct
// [InstanceListResponseIndexMethod]
type instanceListResponseIndexMethodJSON struct {
	Keyword     apijson.Field
	Vector      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceListResponseIndexMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponseIndexMethodJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponseIndexingOptions struct {
	KeywordTokenizer InstanceListResponseIndexingOptionsKeywordTokenizer `json:"keyword_tokenizer"`
	ExtraFields      map[string]interface{}                              `json:"-" api:"extrafields"`
	JSON             instanceListResponseIndexingOptionsJSON             `json:"-"`
}

// instanceListResponseIndexingOptionsJSON contains the JSON metadata for the
// struct [InstanceListResponseIndexingOptions]
type instanceListResponseIndexingOptionsJSON struct {
	KeywordTokenizer apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InstanceListResponseIndexingOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponseIndexingOptionsJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponseIndexingOptionsKeywordTokenizer string

const (
	InstanceListResponseIndexingOptionsKeywordTokenizerPorter  InstanceListResponseIndexingOptionsKeywordTokenizer = "porter"
	InstanceListResponseIndexingOptionsKeywordTokenizerTrigram InstanceListResponseIndexingOptionsKeywordTokenizer = "trigram"
)

func (r InstanceListResponseIndexingOptionsKeywordTokenizer) IsKnown() bool {
	switch r {
	case InstanceListResponseIndexingOptionsKeywordTokenizerPorter, InstanceListResponseIndexingOptionsKeywordTokenizerTrigram:
		return true
	}
	return false
}

type InstanceListResponseMetadata struct {
	CreatedFromAISearchWizard bool                             `json:"created_from_aisearch_wizard"`
	WorkerDomain              string                           `json:"worker_domain"`
	ExtraFields               map[string]interface{}           `json:"-" api:"extrafields"`
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
	CustomDomains           []string                                                        `json:"custom_domains" api:"nullable"`
	DefaultDomainEnabled    bool                                                            `json:"default_domain_enabled"`
	Enabled                 bool                                                            `json:"enabled"`
	Mcp                     InstanceListResponsePublicEndpointParamsMcp                     `json:"mcp"`
	RateLimit               InstanceListResponsePublicEndpointParamsRateLimit               `json:"rate_limit"`
	SearchEndpoint          InstanceListResponsePublicEndpointParamsSearchEndpoint          `json:"search_endpoint"`
	ExtraFields             map[string]interface{}                                          `json:"-" api:"extrafields"`
	JSON                    instanceListResponsePublicEndpointParamsJSON                    `json:"-"`
}

// instanceListResponsePublicEndpointParamsJSON contains the JSON metadata for the
// struct [InstanceListResponsePublicEndpointParams]
type instanceListResponsePublicEndpointParamsJSON struct {
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

func (r *InstanceListResponsePublicEndpointParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponsePublicEndpointParamsJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponsePublicEndpointParamsChatCompletionsEndpoint struct {
	Disabled    bool                                                                `json:"disabled"`
	ExtraFields map[string]interface{}                                              `json:"-" api:"extrafields"`
	JSON        instanceListResponsePublicEndpointParamsChatCompletionsEndpointJSON `json:"-"`
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
	Description string                                          `json:"description"`
	Disabled    bool                                            `json:"disabled"`
	ExtraFields map[string]interface{}                          `json:"-" api:"extrafields"`
	JSON        instanceListResponsePublicEndpointParamsMcpJSON `json:"-"`
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
	PeriodMs    int64                                                      `json:"period_ms"`
	Requests    int64                                                      `json:"requests"`
	Technique   InstanceListResponsePublicEndpointParamsRateLimitTechnique `json:"technique"`
	ExtraFields map[string]interface{}                                     `json:"-" api:"extrafields"`
	JSON        instanceListResponsePublicEndpointParamsRateLimitJSON      `json:"-"`
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
	Disabled    bool                                                       `json:"disabled"`
	ExtraFields map[string]interface{}                                     `json:"-" api:"extrafields"`
	JSON        instanceListResponsePublicEndpointParamsSearchEndpointJSON `json:"-"`
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

type InstanceListResponseRetrievalOptions struct {
	BoostBy          []InstanceListResponseRetrievalOptionsBoostBy        `json:"boost_by"`
	KeywordMatchMode InstanceListResponseRetrievalOptionsKeywordMatchMode `json:"keyword_match_mode"`
	ExtraFields      map[string]interface{}                               `json:"-" api:"extrafields"`
	JSON             instanceListResponseRetrievalOptionsJSON             `json:"-"`
}

// instanceListResponseRetrievalOptionsJSON contains the JSON metadata for the
// struct [InstanceListResponseRetrievalOptions]
type instanceListResponseRetrievalOptionsJSON struct {
	BoostBy          apijson.Field
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

type InstanceListResponseRetrievalOptionsBoostBy struct {
	Field       string                                               `json:"field" api:"required"`
	DataType    InstanceListResponseRetrievalOptionsBoostByDataType  `json:"dataType"`
	Direction   InstanceListResponseRetrievalOptionsBoostByDirection `json:"direction"`
	ExtraFields map[string]interface{}                               `json:"-" api:"extrafields"`
	JSON        instanceListResponseRetrievalOptionsBoostByJSON      `json:"-"`
}

// instanceListResponseRetrievalOptionsBoostByJSON contains the JSON metadata for
// the struct [InstanceListResponseRetrievalOptionsBoostBy]
type instanceListResponseRetrievalOptionsBoostByJSON struct {
	Field       apijson.Field
	DataType    apijson.Field
	Direction   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceListResponseRetrievalOptionsBoostBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponseRetrievalOptionsBoostByJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponseRetrievalOptionsBoostByDataType string

const (
	InstanceListResponseRetrievalOptionsBoostByDataTypeNumber   InstanceListResponseRetrievalOptionsBoostByDataType = "number"
	InstanceListResponseRetrievalOptionsBoostByDataTypeDatetime InstanceListResponseRetrievalOptionsBoostByDataType = "datetime"
	InstanceListResponseRetrievalOptionsBoostByDataTypeText     InstanceListResponseRetrievalOptionsBoostByDataType = "text"
	InstanceListResponseRetrievalOptionsBoostByDataTypeBoolean  InstanceListResponseRetrievalOptionsBoostByDataType = "boolean"
)

func (r InstanceListResponseRetrievalOptionsBoostByDataType) IsKnown() bool {
	switch r {
	case InstanceListResponseRetrievalOptionsBoostByDataTypeNumber, InstanceListResponseRetrievalOptionsBoostByDataTypeDatetime, InstanceListResponseRetrievalOptionsBoostByDataTypeText, InstanceListResponseRetrievalOptionsBoostByDataTypeBoolean:
		return true
	}
	return false
}

type InstanceListResponseRetrievalOptionsBoostByDirection string

const (
	InstanceListResponseRetrievalOptionsBoostByDirectionAsc       InstanceListResponseRetrievalOptionsBoostByDirection = "asc"
	InstanceListResponseRetrievalOptionsBoostByDirectionDesc      InstanceListResponseRetrievalOptionsBoostByDirection = "desc"
	InstanceListResponseRetrievalOptionsBoostByDirectionExists    InstanceListResponseRetrievalOptionsBoostByDirection = "exists"
	InstanceListResponseRetrievalOptionsBoostByDirectionNotExists InstanceListResponseRetrievalOptionsBoostByDirection = "not_exists"
)

func (r InstanceListResponseRetrievalOptionsBoostByDirection) IsKnown() bool {
	switch r {
	case InstanceListResponseRetrievalOptionsBoostByDirectionAsc, InstanceListResponseRetrievalOptionsBoostByDirectionDesc, InstanceListResponseRetrievalOptionsBoostByDirectionExists, InstanceListResponseRetrievalOptionsBoostByDirectionNotExists:
		return true
	}
	return false
}

type InstanceListResponseRetrievalOptionsKeywordMatchMode string

const (
	InstanceListResponseRetrievalOptionsKeywordMatchModeAnd InstanceListResponseRetrievalOptionsKeywordMatchMode = "and"
	InstanceListResponseRetrievalOptionsKeywordMatchModeOr  InstanceListResponseRetrievalOptionsKeywordMatchMode = "or"
)

func (r InstanceListResponseRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case InstanceListResponseRetrievalOptionsKeywordMatchModeAnd, InstanceListResponseRetrievalOptionsKeywordMatchModeOr:
		return true
	}
	return false
}

type InstanceListResponseSourceParams struct {
	ExcludeItems   []string                                   `json:"exclude_items"`
	IncludeItems   []string                                   `json:"include_items"`
	Prefix         string                                     `json:"prefix"`
	R2Jurisdiction string                                     `json:"r2_jurisdiction"`
	WebCrawler     InstanceListResponseSourceParamsWebCrawler `json:"web_crawler"`
	ExtraFields    map[string]interface{}                     `json:"-" api:"extrafields"`
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
	DiscoverOptions InstanceListResponseSourceParamsWebCrawlerDiscoverOptions `json:"discover_options"`
	ParseOptions    InstanceListResponseSourceParamsWebCrawlerParseOptions    `json:"parse_options"`
	ParseType       InstanceListResponseSourceParamsWebCrawlerParseType       `json:"parse_type"`
	ExtraFields     map[string]interface{}                                    `json:"-" api:"extrafields"`
	JSON            instanceListResponseSourceParamsWebCrawlerJSON            `json:"-"`
}

// instanceListResponseSourceParamsWebCrawlerJSON contains the JSON metadata for
// the struct [InstanceListResponseSourceParamsWebCrawler]
type instanceListResponseSourceParamsWebCrawlerJSON struct {
	DiscoverOptions apijson.Field
	ParseOptions    apijson.Field
	ParseType       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InstanceListResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponseSourceParamsWebCrawlerDiscoverOptions struct {
	Depth                float64 `json:"depth"`
	IncludeExternalLinks bool    `json:"include_external_links"`
	IncludeSubdomains    bool    `json:"include_subdomains"`
	// Maximum number of pages to crawl. New values are capped at 100000; instances
	// configured before that cap may report a higher stored value, which the crawler
	// clamps at run time.
	Limit       float64                                                         `json:"limit"`
	MaxAge      float64                                                         `json:"max_age"`
	Source      InstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSource `json:"source"`
	ExtraFields map[string]interface{}                                          `json:"-" api:"extrafields"`
	JSON        instanceListResponseSourceParamsWebCrawlerDiscoverOptionsJSON   `json:"-"`
}

// instanceListResponseSourceParamsWebCrawlerDiscoverOptionsJSON contains the JSON
// metadata for the struct
// [InstanceListResponseSourceParamsWebCrawlerDiscoverOptions]
type instanceListResponseSourceParamsWebCrawlerDiscoverOptionsJSON struct {
	Depth                apijson.Field
	IncludeExternalLinks apijson.Field
	IncludeSubdomains    apijson.Field
	Limit                apijson.Field
	MaxAge               apijson.Field
	Source               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InstanceListResponseSourceParamsWebCrawlerDiscoverOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponseSourceParamsWebCrawlerDiscoverOptionsJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSource string

const (
	InstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll      InstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSource = "all"
	InstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps InstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSource = "sitemaps"
	InstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks    InstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSource = "links"
)

func (r InstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSource) IsKnown() bool {
	switch r {
	case InstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll, InstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps, InstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks:
		return true
	}
	return false
}

type InstanceListResponseSourceParamsWebCrawlerParseOptions struct {
	ContentSelector     []InstanceListResponseSourceParamsWebCrawlerParseOptionsContentSelector `json:"content_selector"`
	IncludeHeaders      map[string]string                                                       `json:"include_headers"`
	IncludeImages       bool                                                                    `json:"include_images"`
	SpecificSitemaps    []string                                                                `json:"specific_sitemaps" format:"uri"`
	UseBrowserRendering bool                                                                    `json:"use_browser_rendering"`
	ExtraFields         map[string]interface{}                                                  `json:"-" api:"extrafields"`
	JSON                instanceListResponseSourceParamsWebCrawlerParseOptionsJSON              `json:"-"`
}

// instanceListResponseSourceParamsWebCrawlerParseOptionsJSON contains the JSON
// metadata for the struct [InstanceListResponseSourceParamsWebCrawlerParseOptions]
type instanceListResponseSourceParamsWebCrawlerParseOptionsJSON struct {
	ContentSelector     apijson.Field
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

type InstanceListResponseSourceParamsWebCrawlerParseOptionsContentSelector struct {
	Path        string                                                                    `json:"path" api:"required"`
	Selector    string                                                                    `json:"selector" api:"required"`
	ExtraFields map[string]interface{}                                                    `json:"-" api:"extrafields"`
	JSON        instanceListResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON `json:"-"`
}

// instanceListResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON
// contains the JSON metadata for the struct
// [InstanceListResponseSourceParamsWebCrawlerParseOptionsContentSelector]
type instanceListResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON struct {
	Path        apijson.Field
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceListResponseSourceParamsWebCrawlerParseOptionsContentSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponseSourceParamsWebCrawlerParseType string

const (
	InstanceListResponseSourceParamsWebCrawlerParseTypeSitemap  InstanceListResponseSourceParamsWebCrawlerParseType = "sitemap"
	InstanceListResponseSourceParamsWebCrawlerParseTypeDiscover InstanceListResponseSourceParamsWebCrawlerParseType = "discover"
)

func (r InstanceListResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case InstanceListResponseSourceParamsWebCrawlerParseTypeSitemap, InstanceListResponseSourceParamsWebCrawlerParseTypeDiscover:
		return true
	}
	return false
}

type InstanceListResponseSyncInterval float64

const (
	InstanceListResponseSyncInterval900   InstanceListResponseSyncInterval = 900
	InstanceListResponseSyncInterval1800  InstanceListResponseSyncInterval = 1800
	InstanceListResponseSyncInterval3600  InstanceListResponseSyncInterval = 3600
	InstanceListResponseSyncInterval7200  InstanceListResponseSyncInterval = 7200
	InstanceListResponseSyncInterval14400 InstanceListResponseSyncInterval = 14400
	InstanceListResponseSyncInterval21600 InstanceListResponseSyncInterval = 21600
	InstanceListResponseSyncInterval43200 InstanceListResponseSyncInterval = 43200
	InstanceListResponseSyncInterval86400 InstanceListResponseSyncInterval = 86400
)

func (r InstanceListResponseSyncInterval) IsKnown() bool {
	switch r {
	case InstanceListResponseSyncInterval900, InstanceListResponseSyncInterval1800, InstanceListResponseSyncInterval3600, InstanceListResponseSyncInterval7200, InstanceListResponseSyncInterval14400, InstanceListResponseSyncInterval21600, InstanceListResponseSyncInterval43200, InstanceListResponseSyncInterval86400:
		return true
	}
	return false
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

type InstanceDeleteResponse struct {
	// AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.
	ID             string                               `json:"id" api:"required"`
	CreatedAt      time.Time                            `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt     time.Time                            `json:"modified_at" api:"required" format:"date-time"`
	AIGatewayID    string                               `json:"ai_gateway_id" api:"nullable"`
	AISearchModel  string                               `json:"ai_search_model" api:"nullable"`
	Cache          bool                                 `json:"cache"`
	CacheThreshold InstanceDeleteResponseCacheThreshold `json:"cache_threshold"`
	// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
	// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
	// (72h), 518400 (6d).
	CacheTTL       InstanceDeleteResponseCacheTTL         `json:"cache_ttl"`
	ChunkOverlap   int64                                  `json:"chunk_overlap"`
	ChunkSize      int64                                  `json:"chunk_size"`
	CreatedBy      string                                 `json:"created_by" api:"nullable"`
	CustomMetadata []InstanceDeleteResponseCustomMetadata `json:"custom_metadata"`
	EmbeddingModel string                                 `json:"embedding_model" api:"nullable"`
	Enable         bool                                   `json:"enable"`
	EngineVersion  float64                                `json:"engine_version"`
	FusionMethod   InstanceDeleteResponseFusionMethod     `json:"fusion_method"`
	// Deprecated — use index_method instead.
	//
	// Deprecated: deprecated
	HybridSearchEnabled bool `json:"hybrid_search_enabled"`
	// Controls which storage backends are used during indexing. Defaults to
	// vector-only.
	IndexMethod          InstanceDeleteResponseIndexMethod          `json:"index_method"`
	IndexingOptions      InstanceDeleteResponseIndexingOptions      `json:"indexing_options" api:"nullable"`
	LastActivity         time.Time                                  `json:"last_activity" api:"nullable" format:"date-time"`
	MaxNumResults        int64                                      `json:"max_num_results"`
	Metadata             InstanceDeleteResponseMetadata             `json:"metadata"`
	ModifiedBy           string                                     `json:"modified_by" api:"nullable"`
	Namespace            string                                     `json:"namespace" api:"nullable"`
	Paused               bool                                       `json:"paused"`
	PublicEndpointID     string                                     `json:"public_endpoint_id" api:"nullable"`
	PublicEndpointParams InstanceDeleteResponsePublicEndpointParams `json:"public_endpoint_params"`
	Reranking            bool                                       `json:"reranking"`
	RerankingModel       string                                     `json:"reranking_model" api:"nullable"`
	RetrievalOptions     InstanceDeleteResponseRetrievalOptions     `json:"retrieval_options" api:"nullable"`
	RewriteModel         string                                     `json:"rewrite_model" api:"nullable"`
	RewriteQuery         bool                                       `json:"rewrite_query"`
	ScoreThreshold       float64                                    `json:"score_threshold"`
	Source               string                                     `json:"source" api:"nullable"`
	SourceParams         InstanceDeleteResponseSourceParams         `json:"source_params" api:"nullable"`
	Status               string                                     `json:"status"`
	// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
	// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
	SyncInterval InstanceDeleteResponseSyncInterval `json:"sync_interval"`
	TokenID      string                             `json:"token_id" format:"uuid"`
	Type         InstanceDeleteResponseType         `json:"type" api:"nullable"`
	JSON         instanceDeleteResponseJSON         `json:"-"`
}

// instanceDeleteResponseJSON contains the JSON metadata for the struct
// [InstanceDeleteResponse]
type instanceDeleteResponseJSON struct {
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

func (r *InstanceDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponseJSON) RawJSON() string {
	return r.raw
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

// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
// (72h), 518400 (6d).
type InstanceDeleteResponseCacheTTL float64

const (
	InstanceDeleteResponseCacheTTL600    InstanceDeleteResponseCacheTTL = 600
	InstanceDeleteResponseCacheTTL1800   InstanceDeleteResponseCacheTTL = 1800
	InstanceDeleteResponseCacheTTL3600   InstanceDeleteResponseCacheTTL = 3600
	InstanceDeleteResponseCacheTTL7200   InstanceDeleteResponseCacheTTL = 7200
	InstanceDeleteResponseCacheTTL21600  InstanceDeleteResponseCacheTTL = 21600
	InstanceDeleteResponseCacheTTL43200  InstanceDeleteResponseCacheTTL = 43200
	InstanceDeleteResponseCacheTTL86400  InstanceDeleteResponseCacheTTL = 86400
	InstanceDeleteResponseCacheTTL172800 InstanceDeleteResponseCacheTTL = 172800
	InstanceDeleteResponseCacheTTL259200 InstanceDeleteResponseCacheTTL = 259200
	InstanceDeleteResponseCacheTTL518400 InstanceDeleteResponseCacheTTL = 518400
)

func (r InstanceDeleteResponseCacheTTL) IsKnown() bool {
	switch r {
	case InstanceDeleteResponseCacheTTL600, InstanceDeleteResponseCacheTTL1800, InstanceDeleteResponseCacheTTL3600, InstanceDeleteResponseCacheTTL7200, InstanceDeleteResponseCacheTTL21600, InstanceDeleteResponseCacheTTL43200, InstanceDeleteResponseCacheTTL86400, InstanceDeleteResponseCacheTTL172800, InstanceDeleteResponseCacheTTL259200, InstanceDeleteResponseCacheTTL518400:
		return true
	}
	return false
}

type InstanceDeleteResponseCustomMetadata struct {
	DataType  InstanceDeleteResponseCustomMetadataDataType `json:"data_type" api:"required"`
	FieldName string                                       `json:"field_name" api:"required"`
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
	InstanceDeleteResponseCustomMetadataDataTypeText     InstanceDeleteResponseCustomMetadataDataType = "text"
	InstanceDeleteResponseCustomMetadataDataTypeNumber   InstanceDeleteResponseCustomMetadataDataType = "number"
	InstanceDeleteResponseCustomMetadataDataTypeBoolean  InstanceDeleteResponseCustomMetadataDataType = "boolean"
	InstanceDeleteResponseCustomMetadataDataTypeDatetime InstanceDeleteResponseCustomMetadataDataType = "datetime"
)

func (r InstanceDeleteResponseCustomMetadataDataType) IsKnown() bool {
	switch r {
	case InstanceDeleteResponseCustomMetadataDataTypeText, InstanceDeleteResponseCustomMetadataDataTypeNumber, InstanceDeleteResponseCustomMetadataDataTypeBoolean, InstanceDeleteResponseCustomMetadataDataTypeDatetime:
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

// Controls which storage backends are used during indexing. Defaults to
// vector-only.
type InstanceDeleteResponseIndexMethod struct {
	// Enable keyword (BM25) storage backend.
	Keyword bool `json:"keyword" api:"required"`
	// Enable vector (embedding) storage backend.
	Vector bool                                  `json:"vector" api:"required"`
	JSON   instanceDeleteResponseIndexMethodJSON `json:"-"`
}

// instanceDeleteResponseIndexMethodJSON contains the JSON metadata for the struct
// [InstanceDeleteResponseIndexMethod]
type instanceDeleteResponseIndexMethodJSON struct {
	Keyword     apijson.Field
	Vector      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceDeleteResponseIndexMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponseIndexMethodJSON) RawJSON() string {
	return r.raw
}

type InstanceDeleteResponseIndexingOptions struct {
	// Tokenizer used for keyword search indexing. porter provides word-level
	// tokenization with Porter stemming (good for natural language queries). trigram
	// enables character-level substring matching (good for partial matches, code,
	// identifiers). Changing this triggers a full re-index. Defaults to porter.
	KeywordTokenizer InstanceDeleteResponseIndexingOptionsKeywordTokenizer `json:"keyword_tokenizer"`
	JSON             instanceDeleteResponseIndexingOptionsJSON             `json:"-"`
}

// instanceDeleteResponseIndexingOptionsJSON contains the JSON metadata for the
// struct [InstanceDeleteResponseIndexingOptions]
type instanceDeleteResponseIndexingOptionsJSON struct {
	KeywordTokenizer apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InstanceDeleteResponseIndexingOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponseIndexingOptionsJSON) RawJSON() string {
	return r.raw
}

// Tokenizer used for keyword search indexing. porter provides word-level
// tokenization with Porter stemming (good for natural language queries). trigram
// enables character-level substring matching (good for partial matches, code,
// identifiers). Changing this triggers a full re-index. Defaults to porter.
type InstanceDeleteResponseIndexingOptionsKeywordTokenizer string

const (
	InstanceDeleteResponseIndexingOptionsKeywordTokenizerPorter  InstanceDeleteResponseIndexingOptionsKeywordTokenizer = "porter"
	InstanceDeleteResponseIndexingOptionsKeywordTokenizerTrigram InstanceDeleteResponseIndexingOptionsKeywordTokenizer = "trigram"
)

func (r InstanceDeleteResponseIndexingOptionsKeywordTokenizer) IsKnown() bool {
	switch r {
	case InstanceDeleteResponseIndexingOptionsKeywordTokenizerPorter, InstanceDeleteResponseIndexingOptionsKeywordTokenizerTrigram:
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
	DefaultDomainEnabled bool                                                     `json:"default_domain_enabled"`
	Enabled              bool                                                     `json:"enabled"`
	Mcp                  InstanceDeleteResponsePublicEndpointParamsMcp            `json:"mcp"`
	RateLimit            InstanceDeleteResponsePublicEndpointParamsRateLimit      `json:"rate_limit"`
	SearchEndpoint       InstanceDeleteResponsePublicEndpointParamsSearchEndpoint `json:"search_endpoint"`
	JSON                 instanceDeleteResponsePublicEndpointParamsJSON           `json:"-"`
}

// instanceDeleteResponsePublicEndpointParamsJSON contains the JSON metadata for
// the struct [InstanceDeleteResponsePublicEndpointParams]
type instanceDeleteResponsePublicEndpointParamsJSON struct {
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

type InstanceDeleteResponseRetrievalOptions struct {
	// Metadata fields to boost search results by. Each entry specifies a metadata
	// field and an optional direction. Direction defaults to 'asc' for
	// numeric/datetime fields and 'exists' for text/boolean fields. Fields must match
	// 'timestamp' or a defined custom_metadata field.
	BoostBy []InstanceDeleteResponseRetrievalOptionsBoostBy `json:"boost_by"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. When omitted on an
	// update, the existing stored value is preserved; when never set, search falls
	// back to 'and'.
	KeywordMatchMode InstanceDeleteResponseRetrievalOptionsKeywordMatchMode `json:"keyword_match_mode"`
	JSON             instanceDeleteResponseRetrievalOptionsJSON             `json:"-"`
}

// instanceDeleteResponseRetrievalOptionsJSON contains the JSON metadata for the
// struct [InstanceDeleteResponseRetrievalOptions]
type instanceDeleteResponseRetrievalOptionsJSON struct {
	BoostBy          apijson.Field
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

type InstanceDeleteResponseRetrievalOptionsBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support all four directions
	// (asc, desc, exists, not_exists); text/boolean fields only support
	// exists/not_exists.
	Field string `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction InstanceDeleteResponseRetrievalOptionsBoostByDirection `json:"direction"`
	JSON      instanceDeleteResponseRetrievalOptionsBoostByJSON      `json:"-"`
}

// instanceDeleteResponseRetrievalOptionsBoostByJSON contains the JSON metadata for
// the struct [InstanceDeleteResponseRetrievalOptionsBoostBy]
type instanceDeleteResponseRetrievalOptionsBoostByJSON struct {
	Field       apijson.Field
	Direction   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceDeleteResponseRetrievalOptionsBoostBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponseRetrievalOptionsBoostByJSON) RawJSON() string {
	return r.raw
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
// for numeric/datetime fields, 'exists' for text/boolean fields.
type InstanceDeleteResponseRetrievalOptionsBoostByDirection string

const (
	InstanceDeleteResponseRetrievalOptionsBoostByDirectionAsc       InstanceDeleteResponseRetrievalOptionsBoostByDirection = "asc"
	InstanceDeleteResponseRetrievalOptionsBoostByDirectionDesc      InstanceDeleteResponseRetrievalOptionsBoostByDirection = "desc"
	InstanceDeleteResponseRetrievalOptionsBoostByDirectionExists    InstanceDeleteResponseRetrievalOptionsBoostByDirection = "exists"
	InstanceDeleteResponseRetrievalOptionsBoostByDirectionNotExists InstanceDeleteResponseRetrievalOptionsBoostByDirection = "not_exists"
)

func (r InstanceDeleteResponseRetrievalOptionsBoostByDirection) IsKnown() bool {
	switch r {
	case InstanceDeleteResponseRetrievalOptionsBoostByDirectionAsc, InstanceDeleteResponseRetrievalOptionsBoostByDirectionDesc, InstanceDeleteResponseRetrievalOptionsBoostByDirectionExists, InstanceDeleteResponseRetrievalOptionsBoostByDirectionNotExists:
		return true
	}
	return false
}

// Controls which documents are candidates for BM25 scoring. 'and' restricts
// candidates to documents containing all query terms; 'or' includes any document
// containing at least one term, ranked by BM25 relevance. When omitted on an
// update, the existing stored value is preserved; when never set, search falls
// back to 'and'.
type InstanceDeleteResponseRetrievalOptionsKeywordMatchMode string

const (
	InstanceDeleteResponseRetrievalOptionsKeywordMatchModeAnd InstanceDeleteResponseRetrievalOptionsKeywordMatchMode = "and"
	InstanceDeleteResponseRetrievalOptionsKeywordMatchModeOr  InstanceDeleteResponseRetrievalOptionsKeywordMatchMode = "or"
)

func (r InstanceDeleteResponseRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case InstanceDeleteResponseRetrievalOptionsKeywordMatchModeAnd, InstanceDeleteResponseRetrievalOptionsKeywordMatchModeOr:
		return true
	}
	return false
}

type InstanceDeleteResponseSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced). Most accounts are limited to 10
	// rules; contact support to raise it.
	ExcludeItems []string `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post). Most accounts are limited to 10 rules; contact
	// support to raise it.
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
	// Options for parse_type 'discover', where Browser Run discovers URLs by link
	// following and sitemaps. Ignored for 'sitemap'.
	DiscoverOptions InstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptions `json:"discover_options"`
	ParseOptions    InstanceDeleteResponseSourceParamsWebCrawlerParseOptions    `json:"parse_options"`
	// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
	// recursively and requires the source to be a Verified zone on this account.
	ParseType InstanceDeleteResponseSourceParamsWebCrawlerParseType `json:"parse_type"`
	JSON      instanceDeleteResponseSourceParamsWebCrawlerJSON      `json:"-"`
}

// instanceDeleteResponseSourceParamsWebCrawlerJSON contains the JSON metadata for
// the struct [InstanceDeleteResponseSourceParamsWebCrawler]
type instanceDeleteResponseSourceParamsWebCrawlerJSON struct {
	DiscoverOptions apijson.Field
	ParseOptions    apijson.Field
	ParseType       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InstanceDeleteResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

// Options for parse_type 'discover', where Browser Run discovers URLs by link
// following and sitemaps. Ignored for 'sitemap'.
type InstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptions struct {
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
	Source InstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSource `json:"source"`
	JSON   instanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsJSON   `json:"-"`
}

// instanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsJSON contains the
// JSON metadata for the struct
// [InstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptions]
type instanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsJSON struct {
	Depth                apijson.Field
	IncludeExternalLinks apijson.Field
	IncludeSubdomains    apijson.Field
	Limit                apijson.Field
	MaxAge               apijson.Field
	Source               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsJSON) RawJSON() string {
	return r.raw
}

// Where the crawler looks for URLs: 'sitemaps' reads sitemap XML only, 'links'
// follows page links only, 'all' does both.
type InstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSource string

const (
	InstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll      InstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSource = "all"
	InstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps InstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSource = "sitemaps"
	InstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks    InstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSource = "links"
)

func (r InstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSource) IsKnown() bool {
	switch r {
	case InstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll, InstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps, InstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks:
		return true
	}
	return false
}

type InstanceDeleteResponseSourceParamsWebCrawlerParseOptions struct {
	// List of path-to-selector mappings for extracting specific content from crawled
	// pages. Each entry pairs a URL glob pattern with a CSS selector. The first
	// matching path wins. Only the matched HTML fragment is stored and indexed. Omit
	// the field to disable content selection — empty arrays are rejected.
	ContentSelector []InstanceDeleteResponseSourceParamsWebCrawlerParseOptionsContentSelector `json:"content_selector"`
	// Up to 5 custom HTTP headers sent with each crawl request. Names must be RFC-7230
	// token characters (no spaces, colons, or control characters); values must be
	// HTAB + printable ASCII (no CR/LF).
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
	ContentSelector     apijson.Field
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

type InstanceDeleteResponseSourceParamsWebCrawlerParseOptionsContentSelector struct {
	// Glob pattern to match against the page URL path. Uses standard glob syntax: \*
	// matches within a segment, \*\* crosses directories.
	Path string `json:"path" api:"required"`
	// CSS selector to extract content from pages matching the path pattern. Must not
	// contain disallowed characters (;, `, $, {, }, \). Must target a single element;
	// if multiple elements match, the selector is ignored and the full page is used.
	Selector string                                                                      `json:"selector" api:"required"`
	JSON     instanceDeleteResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON `json:"-"`
}

// instanceDeleteResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON
// contains the JSON metadata for the struct
// [InstanceDeleteResponseSourceParamsWebCrawlerParseOptionsContentSelector]
type instanceDeleteResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON struct {
	Path        apijson.Field
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceDeleteResponseSourceParamsWebCrawlerParseOptionsContentSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON) RawJSON() string {
	return r.raw
}

// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
// recursively and requires the source to be a Verified zone on this account.
type InstanceDeleteResponseSourceParamsWebCrawlerParseType string

const (
	InstanceDeleteResponseSourceParamsWebCrawlerParseTypeSitemap  InstanceDeleteResponseSourceParamsWebCrawlerParseType = "sitemap"
	InstanceDeleteResponseSourceParamsWebCrawlerParseTypeDiscover InstanceDeleteResponseSourceParamsWebCrawlerParseType = "discover"
)

func (r InstanceDeleteResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case InstanceDeleteResponseSourceParamsWebCrawlerParseTypeSitemap, InstanceDeleteResponseSourceParamsWebCrawlerParseTypeDiscover:
		return true
	}
	return false
}

// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
type InstanceDeleteResponseSyncInterval float64

const (
	InstanceDeleteResponseSyncInterval900   InstanceDeleteResponseSyncInterval = 900
	InstanceDeleteResponseSyncInterval1800  InstanceDeleteResponseSyncInterval = 1800
	InstanceDeleteResponseSyncInterval3600  InstanceDeleteResponseSyncInterval = 3600
	InstanceDeleteResponseSyncInterval7200  InstanceDeleteResponseSyncInterval = 7200
	InstanceDeleteResponseSyncInterval14400 InstanceDeleteResponseSyncInterval = 14400
	InstanceDeleteResponseSyncInterval21600 InstanceDeleteResponseSyncInterval = 21600
	InstanceDeleteResponseSyncInterval43200 InstanceDeleteResponseSyncInterval = 43200
	InstanceDeleteResponseSyncInterval86400 InstanceDeleteResponseSyncInterval = 86400
)

func (r InstanceDeleteResponseSyncInterval) IsKnown() bool {
	switch r {
	case InstanceDeleteResponseSyncInterval900, InstanceDeleteResponseSyncInterval1800, InstanceDeleteResponseSyncInterval3600, InstanceDeleteResponseSyncInterval7200, InstanceDeleteResponseSyncInterval14400, InstanceDeleteResponseSyncInterval21600, InstanceDeleteResponseSyncInterval43200, InstanceDeleteResponseSyncInterval86400:
		return true
	}
	return false
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

type InstanceChatCompletionsResponse struct {
	Choices     []InstanceChatCompletionsResponseChoice `json:"choices" api:"required"`
	Chunks      []InstanceChatCompletionsResponseChunk  `json:"chunks" api:"required"`
	ID          string                                  `json:"id"`
	Model       string                                  `json:"model"`
	Object      string                                  `json:"object"`
	ExtraFields map[string]interface{}                  `json:"-" api:"extrafields"`
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
	Message InstanceChatCompletionsResponseChoicesMessage `json:"message" api:"required"`
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
	Content     InstanceChatCompletionsResponseChoicesMessageContentUnion `json:"content" api:"required,nullable"`
	Role        InstanceChatCompletionsResponseChoicesMessageRole         `json:"role" api:"required"`
	ExtraFields map[string]interface{}                                    `json:"-" api:"extrafields"`
	JSON        instanceChatCompletionsResponseChoicesMessageJSON         `json:"-"`
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

// Union satisfied by [shared.UnionString] or
// [InstanceChatCompletionsResponseChoicesMessageContentArray].
type InstanceChatCompletionsResponseChoicesMessageContentUnion interface {
	ImplementsInstanceChatCompletionsResponseChoicesMessageContentUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InstanceChatCompletionsResponseChoicesMessageContentUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InstanceChatCompletionsResponseChoicesMessageContentArray{}),
		},
	)
}

type InstanceChatCompletionsResponseChoicesMessageContentArray []InstanceChatCompletionsResponseChoicesMessageContentArrayItem

func (r InstanceChatCompletionsResponseChoicesMessageContentArray) ImplementsInstanceChatCompletionsResponseChoicesMessageContentUnion() {
}

type InstanceChatCompletionsResponseChoicesMessageContentArrayItem struct {
	Type InstanceChatCompletionsResponseChoicesMessageContentArrayType `json:"type" api:"required"`
	// This field can have the runtime type of
	// [InstanceChatCompletionsResponseChoicesMessageContentArrayObjectFile].
	File interface{} `json:"file"`
	// This field can have the runtime type of
	// [InstanceChatCompletionsResponseChoicesMessageContentArrayObjectImageURL].
	ImageURL interface{}                                                       `json:"image_url"`
	Text     string                                                            `json:"text"`
	JSON     instanceChatCompletionsResponseChoicesMessageContentArrayItemJSON `json:"-"`
	union    InstanceChatCompletionsResponseChoicesMessageContentArrayUnionItem
}

// instanceChatCompletionsResponseChoicesMessageContentArrayItemJSON contains the
// JSON metadata for the struct
// [InstanceChatCompletionsResponseChoicesMessageContentArrayItem]
type instanceChatCompletionsResponseChoicesMessageContentArrayItemJSON struct {
	Type        apijson.Field
	File        apijson.Field
	ImageURL    apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r instanceChatCompletionsResponseChoicesMessageContentArrayItemJSON) RawJSON() string {
	return r.raw
}

func (r *InstanceChatCompletionsResponseChoicesMessageContentArrayItem) UnmarshalJSON(data []byte) (err error) {
	*r = InstanceChatCompletionsResponseChoicesMessageContentArrayItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [InstanceChatCompletionsResponseChoicesMessageContentArrayUnionItem] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [InstanceChatCompletionsResponseChoicesMessageContentArrayObject],
// [InstanceChatCompletionsResponseChoicesMessageContentArrayObject],
// [InstanceChatCompletionsResponseChoicesMessageContentArrayObject].
func (r InstanceChatCompletionsResponseChoicesMessageContentArrayItem) AsUnion() InstanceChatCompletionsResponseChoicesMessageContentArrayUnionItem {
	return r.union
}

// Union satisfied by
// [InstanceChatCompletionsResponseChoicesMessageContentArrayObject],
// [InstanceChatCompletionsResponseChoicesMessageContentArrayObject] or
// [InstanceChatCompletionsResponseChoicesMessageContentArrayObject].
type InstanceChatCompletionsResponseChoicesMessageContentArrayUnionItem interface {
	implementsInstanceChatCompletionsResponseChoicesMessageContentArrayItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InstanceChatCompletionsResponseChoicesMessageContentArrayUnionItem)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InstanceChatCompletionsResponseChoicesMessageContentArrayObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InstanceChatCompletionsResponseChoicesMessageContentArrayObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InstanceChatCompletionsResponseChoicesMessageContentArrayObject{}),
		},
	)
}

type InstanceChatCompletionsResponseChoicesMessageContentArrayObject struct {
	Text string                                                              `json:"text" api:"required"`
	Type InstanceChatCompletionsResponseChoicesMessageContentArrayObjectType `json:"type" api:"required"`
	JSON instanceChatCompletionsResponseChoicesMessageContentArrayObjectJSON `json:"-"`
}

// instanceChatCompletionsResponseChoicesMessageContentArrayObjectJSON contains the
// JSON metadata for the struct
// [InstanceChatCompletionsResponseChoicesMessageContentArrayObject]
type instanceChatCompletionsResponseChoicesMessageContentArrayObjectJSON struct {
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceChatCompletionsResponseChoicesMessageContentArrayObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceChatCompletionsResponseChoicesMessageContentArrayObjectJSON) RawJSON() string {
	return r.raw
}

func (r InstanceChatCompletionsResponseChoicesMessageContentArrayObject) implementsInstanceChatCompletionsResponseChoicesMessageContentArrayItem() {
}

type InstanceChatCompletionsResponseChoicesMessageContentArrayObjectType string

const (
	InstanceChatCompletionsResponseChoicesMessageContentArrayObjectTypeText InstanceChatCompletionsResponseChoicesMessageContentArrayObjectType = "text"
)

func (r InstanceChatCompletionsResponseChoicesMessageContentArrayObjectType) IsKnown() bool {
	switch r {
	case InstanceChatCompletionsResponseChoicesMessageContentArrayObjectTypeText:
		return true
	}
	return false
}

type InstanceChatCompletionsResponseChoicesMessageContentArrayType string

const (
	InstanceChatCompletionsResponseChoicesMessageContentArrayTypeText     InstanceChatCompletionsResponseChoicesMessageContentArrayType = "text"
	InstanceChatCompletionsResponseChoicesMessageContentArrayTypeImageURL InstanceChatCompletionsResponseChoicesMessageContentArrayType = "image_url"
	InstanceChatCompletionsResponseChoicesMessageContentArrayTypeFile     InstanceChatCompletionsResponseChoicesMessageContentArrayType = "file"
)

func (r InstanceChatCompletionsResponseChoicesMessageContentArrayType) IsKnown() bool {
	switch r {
	case InstanceChatCompletionsResponseChoicesMessageContentArrayTypeText, InstanceChatCompletionsResponseChoicesMessageContentArrayTypeImageURL, InstanceChatCompletionsResponseChoicesMessageContentArrayTypeFile:
		return true
	}
	return false
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
	ID             string                                              `json:"id" api:"required"`
	Score          float64                                             `json:"score" api:"required"`
	Text           string                                              `json:"text" api:"required"`
	Type           string                                              `json:"type" api:"required"`
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
	Key       string                                        `json:"key" api:"required"`
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
	FusionMethod   InstanceChatCompletionsResponseChunksScoringDetailsFusionMethod `json:"fusion_method"`
	KeywordRank    float64                                                         `json:"keyword_rank"`
	KeywordScore   float64                                                         `json:"keyword_score"`
	RerankingScore float64                                                         `json:"reranking_score"`
	VectorRank     float64                                                         `json:"vector_rank"`
	VectorScore    float64                                                         `json:"vector_score"`
	JSON           instanceChatCompletionsResponseChunksScoringDetailsJSON         `json:"-"`
}

// instanceChatCompletionsResponseChunksScoringDetailsJSON contains the JSON
// metadata for the struct [InstanceChatCompletionsResponseChunksScoringDetails]
type instanceChatCompletionsResponseChunksScoringDetailsJSON struct {
	FusionMethod   apijson.Field
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

type InstanceChatCompletionsResponseChunksScoringDetailsFusionMethod string

const (
	InstanceChatCompletionsResponseChunksScoringDetailsFusionMethodRrf InstanceChatCompletionsResponseChunksScoringDetailsFusionMethod = "rrf"
	InstanceChatCompletionsResponseChunksScoringDetailsFusionMethodMax InstanceChatCompletionsResponseChunksScoringDetailsFusionMethod = "max"
)

func (r InstanceChatCompletionsResponseChunksScoringDetailsFusionMethod) IsKnown() bool {
	switch r {
	case InstanceChatCompletionsResponseChunksScoringDetailsFusionMethodRrf, InstanceChatCompletionsResponseChunksScoringDetailsFusionMethodMax:
		return true
	}
	return false
}

type InstanceReadResponse struct {
	// AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.
	ID             string                             `json:"id" api:"required"`
	CreatedAt      time.Time                          `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt     time.Time                          `json:"modified_at" api:"required" format:"date-time"`
	AIGatewayID    string                             `json:"ai_gateway_id" api:"nullable"`
	AISearchModel  string                             `json:"ai_search_model" api:"nullable"`
	Cache          bool                               `json:"cache"`
	CacheThreshold InstanceReadResponseCacheThreshold `json:"cache_threshold"`
	// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
	// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
	// (72h), 518400 (6d).
	CacheTTL       InstanceReadResponseCacheTTL         `json:"cache_ttl"`
	ChunkOverlap   int64                                `json:"chunk_overlap"`
	ChunkSize      int64                                `json:"chunk_size"`
	CreatedBy      string                               `json:"created_by" api:"nullable"`
	CustomMetadata []InstanceReadResponseCustomMetadata `json:"custom_metadata"`
	EmbeddingModel string                               `json:"embedding_model" api:"nullable"`
	Enable         bool                                 `json:"enable"`
	EngineVersion  float64                              `json:"engine_version"`
	FusionMethod   InstanceReadResponseFusionMethod     `json:"fusion_method"`
	// Deprecated — use index_method instead.
	//
	// Deprecated: deprecated
	HybridSearchEnabled bool `json:"hybrid_search_enabled"`
	// Controls which storage backends are used during indexing. Defaults to
	// vector-only.
	IndexMethod          InstanceReadResponseIndexMethod          `json:"index_method"`
	IndexingOptions      InstanceReadResponseIndexingOptions      `json:"indexing_options" api:"nullable"`
	LastActivity         time.Time                                `json:"last_activity" api:"nullable" format:"date-time"`
	MaxNumResults        int64                                    `json:"max_num_results"`
	Metadata             InstanceReadResponseMetadata             `json:"metadata"`
	ModifiedBy           string                                   `json:"modified_by" api:"nullable"`
	Namespace            string                                   `json:"namespace" api:"nullable"`
	Paused               bool                                     `json:"paused"`
	PublicEndpointID     string                                   `json:"public_endpoint_id" api:"nullable"`
	PublicEndpointParams InstanceReadResponsePublicEndpointParams `json:"public_endpoint_params"`
	Reranking            bool                                     `json:"reranking"`
	RerankingModel       string                                   `json:"reranking_model" api:"nullable"`
	RetrievalOptions     InstanceReadResponseRetrievalOptions     `json:"retrieval_options" api:"nullable"`
	RewriteModel         string                                   `json:"rewrite_model" api:"nullable"`
	RewriteQuery         bool                                     `json:"rewrite_query"`
	ScoreThreshold       float64                                  `json:"score_threshold"`
	Source               string                                   `json:"source" api:"nullable"`
	SourceParams         InstanceReadResponseSourceParams         `json:"source_params" api:"nullable"`
	Status               string                                   `json:"status"`
	// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
	// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
	SyncInterval InstanceReadResponseSyncInterval `json:"sync_interval"`
	TokenID      string                           `json:"token_id" format:"uuid"`
	Type         InstanceReadResponseType         `json:"type" api:"nullable"`
	JSON         instanceReadResponseJSON         `json:"-"`
}

// instanceReadResponseJSON contains the JSON metadata for the struct
// [InstanceReadResponse]
type instanceReadResponseJSON struct {
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

func (r *InstanceReadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponseJSON) RawJSON() string {
	return r.raw
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

// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
// (72h), 518400 (6d).
type InstanceReadResponseCacheTTL float64

const (
	InstanceReadResponseCacheTTL600    InstanceReadResponseCacheTTL = 600
	InstanceReadResponseCacheTTL1800   InstanceReadResponseCacheTTL = 1800
	InstanceReadResponseCacheTTL3600   InstanceReadResponseCacheTTL = 3600
	InstanceReadResponseCacheTTL7200   InstanceReadResponseCacheTTL = 7200
	InstanceReadResponseCacheTTL21600  InstanceReadResponseCacheTTL = 21600
	InstanceReadResponseCacheTTL43200  InstanceReadResponseCacheTTL = 43200
	InstanceReadResponseCacheTTL86400  InstanceReadResponseCacheTTL = 86400
	InstanceReadResponseCacheTTL172800 InstanceReadResponseCacheTTL = 172800
	InstanceReadResponseCacheTTL259200 InstanceReadResponseCacheTTL = 259200
	InstanceReadResponseCacheTTL518400 InstanceReadResponseCacheTTL = 518400
)

func (r InstanceReadResponseCacheTTL) IsKnown() bool {
	switch r {
	case InstanceReadResponseCacheTTL600, InstanceReadResponseCacheTTL1800, InstanceReadResponseCacheTTL3600, InstanceReadResponseCacheTTL7200, InstanceReadResponseCacheTTL21600, InstanceReadResponseCacheTTL43200, InstanceReadResponseCacheTTL86400, InstanceReadResponseCacheTTL172800, InstanceReadResponseCacheTTL259200, InstanceReadResponseCacheTTL518400:
		return true
	}
	return false
}

type InstanceReadResponseCustomMetadata struct {
	DataType  InstanceReadResponseCustomMetadataDataType `json:"data_type" api:"required"`
	FieldName string                                     `json:"field_name" api:"required"`
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
	InstanceReadResponseCustomMetadataDataTypeText     InstanceReadResponseCustomMetadataDataType = "text"
	InstanceReadResponseCustomMetadataDataTypeNumber   InstanceReadResponseCustomMetadataDataType = "number"
	InstanceReadResponseCustomMetadataDataTypeBoolean  InstanceReadResponseCustomMetadataDataType = "boolean"
	InstanceReadResponseCustomMetadataDataTypeDatetime InstanceReadResponseCustomMetadataDataType = "datetime"
)

func (r InstanceReadResponseCustomMetadataDataType) IsKnown() bool {
	switch r {
	case InstanceReadResponseCustomMetadataDataTypeText, InstanceReadResponseCustomMetadataDataTypeNumber, InstanceReadResponseCustomMetadataDataTypeBoolean, InstanceReadResponseCustomMetadataDataTypeDatetime:
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

// Controls which storage backends are used during indexing. Defaults to
// vector-only.
type InstanceReadResponseIndexMethod struct {
	// Enable keyword (BM25) storage backend.
	Keyword bool `json:"keyword" api:"required"`
	// Enable vector (embedding) storage backend.
	Vector bool                                `json:"vector" api:"required"`
	JSON   instanceReadResponseIndexMethodJSON `json:"-"`
}

// instanceReadResponseIndexMethodJSON contains the JSON metadata for the struct
// [InstanceReadResponseIndexMethod]
type instanceReadResponseIndexMethodJSON struct {
	Keyword     apijson.Field
	Vector      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceReadResponseIndexMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponseIndexMethodJSON) RawJSON() string {
	return r.raw
}

type InstanceReadResponseIndexingOptions struct {
	// Tokenizer used for keyword search indexing. porter provides word-level
	// tokenization with Porter stemming (good for natural language queries). trigram
	// enables character-level substring matching (good for partial matches, code,
	// identifiers). Changing this triggers a full re-index. Defaults to porter.
	KeywordTokenizer InstanceReadResponseIndexingOptionsKeywordTokenizer `json:"keyword_tokenizer"`
	JSON             instanceReadResponseIndexingOptionsJSON             `json:"-"`
}

// instanceReadResponseIndexingOptionsJSON contains the JSON metadata for the
// struct [InstanceReadResponseIndexingOptions]
type instanceReadResponseIndexingOptionsJSON struct {
	KeywordTokenizer apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InstanceReadResponseIndexingOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponseIndexingOptionsJSON) RawJSON() string {
	return r.raw
}

// Tokenizer used for keyword search indexing. porter provides word-level
// tokenization with Porter stemming (good for natural language queries). trigram
// enables character-level substring matching (good for partial matches, code,
// identifiers). Changing this triggers a full re-index. Defaults to porter.
type InstanceReadResponseIndexingOptionsKeywordTokenizer string

const (
	InstanceReadResponseIndexingOptionsKeywordTokenizerPorter  InstanceReadResponseIndexingOptionsKeywordTokenizer = "porter"
	InstanceReadResponseIndexingOptionsKeywordTokenizerTrigram InstanceReadResponseIndexingOptionsKeywordTokenizer = "trigram"
)

func (r InstanceReadResponseIndexingOptionsKeywordTokenizer) IsKnown() bool {
	switch r {
	case InstanceReadResponseIndexingOptionsKeywordTokenizerPorter, InstanceReadResponseIndexingOptionsKeywordTokenizerTrigram:
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
	DefaultDomainEnabled bool                                                   `json:"default_domain_enabled"`
	Enabled              bool                                                   `json:"enabled"`
	Mcp                  InstanceReadResponsePublicEndpointParamsMcp            `json:"mcp"`
	RateLimit            InstanceReadResponsePublicEndpointParamsRateLimit      `json:"rate_limit"`
	SearchEndpoint       InstanceReadResponsePublicEndpointParamsSearchEndpoint `json:"search_endpoint"`
	JSON                 instanceReadResponsePublicEndpointParamsJSON           `json:"-"`
}

// instanceReadResponsePublicEndpointParamsJSON contains the JSON metadata for the
// struct [InstanceReadResponsePublicEndpointParams]
type instanceReadResponsePublicEndpointParamsJSON struct {
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

type InstanceReadResponseRetrievalOptions struct {
	// Metadata fields to boost search results by. Each entry specifies a metadata
	// field and an optional direction. Direction defaults to 'asc' for
	// numeric/datetime fields and 'exists' for text/boolean fields. Fields must match
	// 'timestamp' or a defined custom_metadata field.
	BoostBy []InstanceReadResponseRetrievalOptionsBoostBy `json:"boost_by"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. When omitted on an
	// update, the existing stored value is preserved; when never set, search falls
	// back to 'and'.
	KeywordMatchMode InstanceReadResponseRetrievalOptionsKeywordMatchMode `json:"keyword_match_mode"`
	JSON             instanceReadResponseRetrievalOptionsJSON             `json:"-"`
}

// instanceReadResponseRetrievalOptionsJSON contains the JSON metadata for the
// struct [InstanceReadResponseRetrievalOptions]
type instanceReadResponseRetrievalOptionsJSON struct {
	BoostBy          apijson.Field
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

type InstanceReadResponseRetrievalOptionsBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support all four directions
	// (asc, desc, exists, not_exists); text/boolean fields only support
	// exists/not_exists.
	Field string `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction InstanceReadResponseRetrievalOptionsBoostByDirection `json:"direction"`
	JSON      instanceReadResponseRetrievalOptionsBoostByJSON      `json:"-"`
}

// instanceReadResponseRetrievalOptionsBoostByJSON contains the JSON metadata for
// the struct [InstanceReadResponseRetrievalOptionsBoostBy]
type instanceReadResponseRetrievalOptionsBoostByJSON struct {
	Field       apijson.Field
	Direction   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceReadResponseRetrievalOptionsBoostBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponseRetrievalOptionsBoostByJSON) RawJSON() string {
	return r.raw
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
// for numeric/datetime fields, 'exists' for text/boolean fields.
type InstanceReadResponseRetrievalOptionsBoostByDirection string

const (
	InstanceReadResponseRetrievalOptionsBoostByDirectionAsc       InstanceReadResponseRetrievalOptionsBoostByDirection = "asc"
	InstanceReadResponseRetrievalOptionsBoostByDirectionDesc      InstanceReadResponseRetrievalOptionsBoostByDirection = "desc"
	InstanceReadResponseRetrievalOptionsBoostByDirectionExists    InstanceReadResponseRetrievalOptionsBoostByDirection = "exists"
	InstanceReadResponseRetrievalOptionsBoostByDirectionNotExists InstanceReadResponseRetrievalOptionsBoostByDirection = "not_exists"
)

func (r InstanceReadResponseRetrievalOptionsBoostByDirection) IsKnown() bool {
	switch r {
	case InstanceReadResponseRetrievalOptionsBoostByDirectionAsc, InstanceReadResponseRetrievalOptionsBoostByDirectionDesc, InstanceReadResponseRetrievalOptionsBoostByDirectionExists, InstanceReadResponseRetrievalOptionsBoostByDirectionNotExists:
		return true
	}
	return false
}

// Controls which documents are candidates for BM25 scoring. 'and' restricts
// candidates to documents containing all query terms; 'or' includes any document
// containing at least one term, ranked by BM25 relevance. When omitted on an
// update, the existing stored value is preserved; when never set, search falls
// back to 'and'.
type InstanceReadResponseRetrievalOptionsKeywordMatchMode string

const (
	InstanceReadResponseRetrievalOptionsKeywordMatchModeAnd InstanceReadResponseRetrievalOptionsKeywordMatchMode = "and"
	InstanceReadResponseRetrievalOptionsKeywordMatchModeOr  InstanceReadResponseRetrievalOptionsKeywordMatchMode = "or"
)

func (r InstanceReadResponseRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case InstanceReadResponseRetrievalOptionsKeywordMatchModeAnd, InstanceReadResponseRetrievalOptionsKeywordMatchModeOr:
		return true
	}
	return false
}

type InstanceReadResponseSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced). Most accounts are limited to 10
	// rules; contact support to raise it.
	ExcludeItems []string `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post). Most accounts are limited to 10 rules; contact
	// support to raise it.
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
	// Options for parse_type 'discover', where Browser Run discovers URLs by link
	// following and sitemaps. Ignored for 'sitemap'.
	DiscoverOptions InstanceReadResponseSourceParamsWebCrawlerDiscoverOptions `json:"discover_options"`
	ParseOptions    InstanceReadResponseSourceParamsWebCrawlerParseOptions    `json:"parse_options"`
	// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
	// recursively and requires the source to be a Verified zone on this account.
	ParseType InstanceReadResponseSourceParamsWebCrawlerParseType `json:"parse_type"`
	JSON      instanceReadResponseSourceParamsWebCrawlerJSON      `json:"-"`
}

// instanceReadResponseSourceParamsWebCrawlerJSON contains the JSON metadata for
// the struct [InstanceReadResponseSourceParamsWebCrawler]
type instanceReadResponseSourceParamsWebCrawlerJSON struct {
	DiscoverOptions apijson.Field
	ParseOptions    apijson.Field
	ParseType       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InstanceReadResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

// Options for parse_type 'discover', where Browser Run discovers URLs by link
// following and sitemaps. Ignored for 'sitemap'.
type InstanceReadResponseSourceParamsWebCrawlerDiscoverOptions struct {
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
	Source InstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSource `json:"source"`
	JSON   instanceReadResponseSourceParamsWebCrawlerDiscoverOptionsJSON   `json:"-"`
}

// instanceReadResponseSourceParamsWebCrawlerDiscoverOptionsJSON contains the JSON
// metadata for the struct
// [InstanceReadResponseSourceParamsWebCrawlerDiscoverOptions]
type instanceReadResponseSourceParamsWebCrawlerDiscoverOptionsJSON struct {
	Depth                apijson.Field
	IncludeExternalLinks apijson.Field
	IncludeSubdomains    apijson.Field
	Limit                apijson.Field
	MaxAge               apijson.Field
	Source               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InstanceReadResponseSourceParamsWebCrawlerDiscoverOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponseSourceParamsWebCrawlerDiscoverOptionsJSON) RawJSON() string {
	return r.raw
}

// Where the crawler looks for URLs: 'sitemaps' reads sitemap XML only, 'links'
// follows page links only, 'all' does both.
type InstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSource string

const (
	InstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll      InstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSource = "all"
	InstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps InstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSource = "sitemaps"
	InstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks    InstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSource = "links"
)

func (r InstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSource) IsKnown() bool {
	switch r {
	case InstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll, InstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps, InstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks:
		return true
	}
	return false
}

type InstanceReadResponseSourceParamsWebCrawlerParseOptions struct {
	// List of path-to-selector mappings for extracting specific content from crawled
	// pages. Each entry pairs a URL glob pattern with a CSS selector. The first
	// matching path wins. Only the matched HTML fragment is stored and indexed. Omit
	// the field to disable content selection — empty arrays are rejected.
	ContentSelector []InstanceReadResponseSourceParamsWebCrawlerParseOptionsContentSelector `json:"content_selector"`
	// Up to 5 custom HTTP headers sent with each crawl request. Names must be RFC-7230
	// token characters (no spaces, colons, or control characters); values must be
	// HTAB + printable ASCII (no CR/LF).
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
	ContentSelector     apijson.Field
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

type InstanceReadResponseSourceParamsWebCrawlerParseOptionsContentSelector struct {
	// Glob pattern to match against the page URL path. Uses standard glob syntax: \*
	// matches within a segment, \*\* crosses directories.
	Path string `json:"path" api:"required"`
	// CSS selector to extract content from pages matching the path pattern. Must not
	// contain disallowed characters (;, `, $, {, }, \). Must target a single element;
	// if multiple elements match, the selector is ignored and the full page is used.
	Selector string                                                                    `json:"selector" api:"required"`
	JSON     instanceReadResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON `json:"-"`
}

// instanceReadResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON
// contains the JSON metadata for the struct
// [InstanceReadResponseSourceParamsWebCrawlerParseOptionsContentSelector]
type instanceReadResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON struct {
	Path        apijson.Field
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceReadResponseSourceParamsWebCrawlerParseOptionsContentSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON) RawJSON() string {
	return r.raw
}

// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
// recursively and requires the source to be a Verified zone on this account.
type InstanceReadResponseSourceParamsWebCrawlerParseType string

const (
	InstanceReadResponseSourceParamsWebCrawlerParseTypeSitemap  InstanceReadResponseSourceParamsWebCrawlerParseType = "sitemap"
	InstanceReadResponseSourceParamsWebCrawlerParseTypeDiscover InstanceReadResponseSourceParamsWebCrawlerParseType = "discover"
)

func (r InstanceReadResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case InstanceReadResponseSourceParamsWebCrawlerParseTypeSitemap, InstanceReadResponseSourceParamsWebCrawlerParseTypeDiscover:
		return true
	}
	return false
}

// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
type InstanceReadResponseSyncInterval float64

const (
	InstanceReadResponseSyncInterval900   InstanceReadResponseSyncInterval = 900
	InstanceReadResponseSyncInterval1800  InstanceReadResponseSyncInterval = 1800
	InstanceReadResponseSyncInterval3600  InstanceReadResponseSyncInterval = 3600
	InstanceReadResponseSyncInterval7200  InstanceReadResponseSyncInterval = 7200
	InstanceReadResponseSyncInterval14400 InstanceReadResponseSyncInterval = 14400
	InstanceReadResponseSyncInterval21600 InstanceReadResponseSyncInterval = 21600
	InstanceReadResponseSyncInterval43200 InstanceReadResponseSyncInterval = 43200
	InstanceReadResponseSyncInterval86400 InstanceReadResponseSyncInterval = 86400
)

func (r InstanceReadResponseSyncInterval) IsKnown() bool {
	switch r {
	case InstanceReadResponseSyncInterval900, InstanceReadResponseSyncInterval1800, InstanceReadResponseSyncInterval3600, InstanceReadResponseSyncInterval7200, InstanceReadResponseSyncInterval14400, InstanceReadResponseSyncInterval21600, InstanceReadResponseSyncInterval43200, InstanceReadResponseSyncInterval86400:
		return true
	}
	return false
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

type InstanceSearchResponse struct {
	Chunks      []InstanceSearchResponseChunk   `json:"chunks" api:"required"`
	QueryKind   InstanceSearchResponseQueryKind `json:"query_kind" api:"required"`
	SearchQuery string                          `json:"search_query"`
	JSON        instanceSearchResponseJSON      `json:"-"`
}

// instanceSearchResponseJSON contains the JSON metadata for the struct
// [InstanceSearchResponse]
type instanceSearchResponseJSON struct {
	Chunks      apijson.Field
	QueryKind   apijson.Field
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
	ID             string                                     `json:"id" api:"required"`
	Score          float64                                    `json:"score" api:"required"`
	Text           string                                     `json:"text" api:"required"`
	Type           string                                     `json:"type" api:"required"`
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
	Key       string                               `json:"key" api:"required"`
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
	FusionMethod   InstanceSearchResponseChunksScoringDetailsFusionMethod `json:"fusion_method"`
	KeywordRank    float64                                                `json:"keyword_rank"`
	KeywordScore   float64                                                `json:"keyword_score"`
	RerankingScore float64                                                `json:"reranking_score"`
	VectorRank     float64                                                `json:"vector_rank"`
	VectorScore    float64                                                `json:"vector_score"`
	JSON           instanceSearchResponseChunksScoringDetailsJSON         `json:"-"`
}

// instanceSearchResponseChunksScoringDetailsJSON contains the JSON metadata for
// the struct [InstanceSearchResponseChunksScoringDetails]
type instanceSearchResponseChunksScoringDetailsJSON struct {
	FusionMethod   apijson.Field
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

type InstanceSearchResponseChunksScoringDetailsFusionMethod string

const (
	InstanceSearchResponseChunksScoringDetailsFusionMethodRrf InstanceSearchResponseChunksScoringDetailsFusionMethod = "rrf"
	InstanceSearchResponseChunksScoringDetailsFusionMethodMax InstanceSearchResponseChunksScoringDetailsFusionMethod = "max"
)

func (r InstanceSearchResponseChunksScoringDetailsFusionMethod) IsKnown() bool {
	switch r {
	case InstanceSearchResponseChunksScoringDetailsFusionMethodRrf, InstanceSearchResponseChunksScoringDetailsFusionMethodMax:
		return true
	}
	return false
}

type InstanceSearchResponseQueryKind string

const (
	InstanceSearchResponseQueryKindText       InstanceSearchResponseQueryKind = "text"
	InstanceSearchResponseQueryKindImage      InstanceSearchResponseQueryKind = "image"
	InstanceSearchResponseQueryKindMultimodal InstanceSearchResponseQueryKind = "multimodal"
)

func (r InstanceSearchResponseQueryKind) IsKnown() bool {
	switch r {
	case InstanceSearchResponseQueryKindText, InstanceSearchResponseQueryKindImage, InstanceSearchResponseQueryKindMultimodal:
		return true
	}
	return false
}

type InstanceStatsResponse struct {
	Completed int64 `json:"completed"`
	// True when status counts are unavailable (e.g. legacy stats query exceeded D1
	// statement-size limit). Counts are omitted in this case.
	Degraded bool `json:"degraded"`
	// Engine-specific metadata. Present only for managed (v3) instances.
	Engine            InstanceStatsResponseEngine `json:"engine"`
	Error             int64                       `json:"error"`
	FileEmbedErrors   map[string]interface{}      `json:"file_embed_errors"`
	IndexSourceErrors map[string]interface{}      `json:"index_source_errors"`
	LastActivity      time.Time                   `json:"last_activity" format:"date-time"`
	Outdated          int64                       `json:"outdated"`
	Queued            int64                       `json:"queued"`
	Running           int64                       `json:"running"`
	Skipped           int64                       `json:"skipped"`
	JSON              instanceStatsResponseJSON   `json:"-"`
}

// instanceStatsResponseJSON contains the JSON metadata for the struct
// [InstanceStatsResponse]
type instanceStatsResponseJSON struct {
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

func (r *InstanceStatsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceStatsResponseJSON) RawJSON() string {
	return r.raw
}

// Engine-specific metadata. Present only for managed (v3) instances.
type InstanceStatsResponseEngine struct {
	// R2 bucket storage usage in bytes.
	R2 InstanceStatsResponseEngineR2 `json:"r2"`
	// Vectorize index metadata (dimensions, vector count).
	Vectorize InstanceStatsResponseEngineVectorize `json:"vectorize"`
	JSON      instanceStatsResponseEngineJSON      `json:"-"`
}

// instanceStatsResponseEngineJSON contains the JSON metadata for the struct
// [InstanceStatsResponseEngine]
type instanceStatsResponseEngineJSON struct {
	R2          apijson.Field
	Vectorize   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceStatsResponseEngine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceStatsResponseEngineJSON) RawJSON() string {
	return r.raw
}

// R2 bucket storage usage in bytes.
type InstanceStatsResponseEngineR2 struct {
	MetadataSizeBytes int64                             `json:"metadataSizeBytes" api:"required"`
	ObjectCount       int64                             `json:"objectCount" api:"required"`
	PayloadSizeBytes  int64                             `json:"payloadSizeBytes" api:"required"`
	JSON              instanceStatsResponseEngineR2JSON `json:"-"`
}

// instanceStatsResponseEngineR2JSON contains the JSON metadata for the struct
// [InstanceStatsResponseEngineR2]
type instanceStatsResponseEngineR2JSON struct {
	MetadataSizeBytes apijson.Field
	ObjectCount       apijson.Field
	PayloadSizeBytes  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InstanceStatsResponseEngineR2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceStatsResponseEngineR2JSON) RawJSON() string {
	return r.raw
}

// Vectorize index metadata (dimensions, vector count).
type InstanceStatsResponseEngineVectorize struct {
	Dimensions   int64                                    `json:"dimensions" api:"required"`
	VectorsCount int64                                    `json:"vectorsCount" api:"required"`
	JSON         instanceStatsResponseEngineVectorizeJSON `json:"-"`
}

// instanceStatsResponseEngineVectorizeJSON contains the JSON metadata for the
// struct [InstanceStatsResponseEngineVectorize]
type instanceStatsResponseEngineVectorizeJSON struct {
	Dimensions   apijson.Field
	VectorsCount apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InstanceStatsResponseEngineVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceStatsResponseEngineVectorizeJSON) RawJSON() string {
	return r.raw
}

type InstanceNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.
	ID             param.Field[string]                          `json:"id" api:"required"`
	AIGatewayID    param.Field[string]                          `json:"ai_gateway_id"`
	AISearchModel  param.Field[string]                          `json:"ai_search_model"`
	Cache          param.Field[bool]                            `json:"cache"`
	CacheThreshold param.Field[InstanceNewParamsCacheThreshold] `json:"cache_threshold"`
	// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
	// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
	// (72h), 518400 (6d).
	CacheTTL       param.Field[InstanceNewParamsCacheTTL]         `json:"cache_ttl"`
	Chunk          param.Field[bool]                              `json:"chunk"`
	ChunkOverlap   param.Field[int64]                             `json:"chunk_overlap"`
	ChunkSize      param.Field[int64]                             `json:"chunk_size"`
	CustomMetadata param.Field[[]InstanceNewParamsCustomMetadata] `json:"custom_metadata"`
	EmbeddingModel param.Field[string]                            `json:"embedding_model"`
	FusionMethod   param.Field[InstanceNewParamsFusionMethod]     `json:"fusion_method"`
	// Deprecated — use index_method instead.
	HybridSearchEnabled param.Field[bool] `json:"hybrid_search_enabled"`
	// Controls which storage backends are used during indexing. Defaults to
	// vector-only.
	IndexMethod          param.Field[InstanceNewParamsIndexMethod]          `json:"index_method"`
	IndexingOptions      param.Field[InstanceNewParamsIndexingOptions]      `json:"indexing_options"`
	MaxNumResults        param.Field[int64]                                 `json:"max_num_results"`
	Metadata             param.Field[InstanceNewParamsMetadata]             `json:"metadata"`
	PublicEndpointParams param.Field[InstanceNewParamsPublicEndpointParams] `json:"public_endpoint_params"`
	Reranking            param.Field[bool]                                  `json:"reranking"`
	RerankingModel       param.Field[string]                                `json:"reranking_model"`
	RetrievalOptions     param.Field[InstanceNewParamsRetrievalOptions]     `json:"retrieval_options"`
	RewriteModel         param.Field[string]                                `json:"rewrite_model"`
	RewriteQuery         param.Field[bool]                                  `json:"rewrite_query"`
	ScoreThreshold       param.Field[float64]                               `json:"score_threshold"`
	Source               param.Field[string]                                `json:"source"`
	SourceParams         param.Field[InstanceNewParamsSourceParams]         `json:"source_params"`
	// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
	// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
	SyncInterval param.Field[InstanceNewParamsSyncInterval] `json:"sync_interval"`
	TokenID      param.Field[string]                        `json:"token_id" format:"uuid"`
	Type         param.Field[InstanceNewParamsType]         `json:"type"`
}

func (r InstanceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceNewParamsCacheThreshold string

const (
	InstanceNewParamsCacheThresholdSuperStrictMatch InstanceNewParamsCacheThreshold = "super_strict_match"
	InstanceNewParamsCacheThresholdCloseEnough      InstanceNewParamsCacheThreshold = "close_enough"
	InstanceNewParamsCacheThresholdFlexibleFriend   InstanceNewParamsCacheThreshold = "flexible_friend"
	InstanceNewParamsCacheThresholdAnythingGoes     InstanceNewParamsCacheThreshold = "anything_goes"
)

func (r InstanceNewParamsCacheThreshold) IsKnown() bool {
	switch r {
	case InstanceNewParamsCacheThresholdSuperStrictMatch, InstanceNewParamsCacheThresholdCloseEnough, InstanceNewParamsCacheThresholdFlexibleFriend, InstanceNewParamsCacheThresholdAnythingGoes:
		return true
	}
	return false
}

// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
// (72h), 518400 (6d).
type InstanceNewParamsCacheTTL float64

const (
	InstanceNewParamsCacheTTL600    InstanceNewParamsCacheTTL = 600
	InstanceNewParamsCacheTTL1800   InstanceNewParamsCacheTTL = 1800
	InstanceNewParamsCacheTTL3600   InstanceNewParamsCacheTTL = 3600
	InstanceNewParamsCacheTTL7200   InstanceNewParamsCacheTTL = 7200
	InstanceNewParamsCacheTTL21600  InstanceNewParamsCacheTTL = 21600
	InstanceNewParamsCacheTTL43200  InstanceNewParamsCacheTTL = 43200
	InstanceNewParamsCacheTTL86400  InstanceNewParamsCacheTTL = 86400
	InstanceNewParamsCacheTTL172800 InstanceNewParamsCacheTTL = 172800
	InstanceNewParamsCacheTTL259200 InstanceNewParamsCacheTTL = 259200
	InstanceNewParamsCacheTTL518400 InstanceNewParamsCacheTTL = 518400
)

func (r InstanceNewParamsCacheTTL) IsKnown() bool {
	switch r {
	case InstanceNewParamsCacheTTL600, InstanceNewParamsCacheTTL1800, InstanceNewParamsCacheTTL3600, InstanceNewParamsCacheTTL7200, InstanceNewParamsCacheTTL21600, InstanceNewParamsCacheTTL43200, InstanceNewParamsCacheTTL86400, InstanceNewParamsCacheTTL172800, InstanceNewParamsCacheTTL259200, InstanceNewParamsCacheTTL518400:
		return true
	}
	return false
}

type InstanceNewParamsCustomMetadata struct {
	DataType  param.Field[InstanceNewParamsCustomMetadataDataType] `json:"data_type" api:"required"`
	FieldName param.Field[string]                                  `json:"field_name" api:"required"`
}

func (r InstanceNewParamsCustomMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceNewParamsCustomMetadataDataType string

const (
	InstanceNewParamsCustomMetadataDataTypeText     InstanceNewParamsCustomMetadataDataType = "text"
	InstanceNewParamsCustomMetadataDataTypeNumber   InstanceNewParamsCustomMetadataDataType = "number"
	InstanceNewParamsCustomMetadataDataTypeBoolean  InstanceNewParamsCustomMetadataDataType = "boolean"
	InstanceNewParamsCustomMetadataDataTypeDatetime InstanceNewParamsCustomMetadataDataType = "datetime"
)

func (r InstanceNewParamsCustomMetadataDataType) IsKnown() bool {
	switch r {
	case InstanceNewParamsCustomMetadataDataTypeText, InstanceNewParamsCustomMetadataDataTypeNumber, InstanceNewParamsCustomMetadataDataTypeBoolean, InstanceNewParamsCustomMetadataDataTypeDatetime:
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

// Controls which storage backends are used during indexing. Defaults to
// vector-only.
type InstanceNewParamsIndexMethod struct {
	// Enable keyword (BM25) storage backend.
	Keyword param.Field[bool] `json:"keyword" api:"required"`
	// Enable vector (embedding) storage backend.
	Vector param.Field[bool] `json:"vector" api:"required"`
}

func (r InstanceNewParamsIndexMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceNewParamsIndexingOptions struct {
	// Tokenizer used for keyword search indexing. porter provides word-level
	// tokenization with Porter stemming (good for natural language queries). trigram
	// enables character-level substring matching (good for partial matches, code,
	// identifiers). Changing this triggers a full re-index. Defaults to porter.
	KeywordTokenizer param.Field[InstanceNewParamsIndexingOptionsKeywordTokenizer] `json:"keyword_tokenizer"`
}

func (r InstanceNewParamsIndexingOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Tokenizer used for keyword search indexing. porter provides word-level
// tokenization with Porter stemming (good for natural language queries). trigram
// enables character-level substring matching (good for partial matches, code,
// identifiers). Changing this triggers a full re-index. Defaults to porter.
type InstanceNewParamsIndexingOptionsKeywordTokenizer string

const (
	InstanceNewParamsIndexingOptionsKeywordTokenizerPorter  InstanceNewParamsIndexingOptionsKeywordTokenizer = "porter"
	InstanceNewParamsIndexingOptionsKeywordTokenizerTrigram InstanceNewParamsIndexingOptionsKeywordTokenizer = "trigram"
)

func (r InstanceNewParamsIndexingOptionsKeywordTokenizer) IsKnown() bool {
	switch r {
	case InstanceNewParamsIndexingOptionsKeywordTokenizerPorter, InstanceNewParamsIndexingOptionsKeywordTokenizerTrigram:
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
	DefaultDomainEnabled param.Field[bool]                                                `json:"default_domain_enabled"`
	Enabled              param.Field[bool]                                                `json:"enabled"`
	Mcp                  param.Field[InstanceNewParamsPublicEndpointParamsMcp]            `json:"mcp"`
	RateLimit            param.Field[InstanceNewParamsPublicEndpointParamsRateLimit]      `json:"rate_limit"`
	SearchEndpoint       param.Field[InstanceNewParamsPublicEndpointParamsSearchEndpoint] `json:"search_endpoint"`
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

type InstanceNewParamsRetrievalOptions struct {
	// Metadata fields to boost search results by. Each entry specifies a metadata
	// field and an optional direction. Direction defaults to 'asc' for
	// numeric/datetime fields and 'exists' for text/boolean fields. Fields must match
	// 'timestamp' or a defined custom_metadata field.
	BoostBy param.Field[[]InstanceNewParamsRetrievalOptionsBoostBy] `json:"boost_by"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. When omitted on an
	// update, the existing stored value is preserved; when never set, search falls
	// back to 'and'.
	KeywordMatchMode param.Field[InstanceNewParamsRetrievalOptionsKeywordMatchMode] `json:"keyword_match_mode"`
}

func (r InstanceNewParamsRetrievalOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceNewParamsRetrievalOptionsBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support all four directions
	// (asc, desc, exists, not_exists); text/boolean fields only support
	// exists/not_exists.
	Field param.Field[string] `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction param.Field[InstanceNewParamsRetrievalOptionsBoostByDirection] `json:"direction"`
}

func (r InstanceNewParamsRetrievalOptionsBoostBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
// for numeric/datetime fields, 'exists' for text/boolean fields.
type InstanceNewParamsRetrievalOptionsBoostByDirection string

const (
	InstanceNewParamsRetrievalOptionsBoostByDirectionAsc       InstanceNewParamsRetrievalOptionsBoostByDirection = "asc"
	InstanceNewParamsRetrievalOptionsBoostByDirectionDesc      InstanceNewParamsRetrievalOptionsBoostByDirection = "desc"
	InstanceNewParamsRetrievalOptionsBoostByDirectionExists    InstanceNewParamsRetrievalOptionsBoostByDirection = "exists"
	InstanceNewParamsRetrievalOptionsBoostByDirectionNotExists InstanceNewParamsRetrievalOptionsBoostByDirection = "not_exists"
)

func (r InstanceNewParamsRetrievalOptionsBoostByDirection) IsKnown() bool {
	switch r {
	case InstanceNewParamsRetrievalOptionsBoostByDirectionAsc, InstanceNewParamsRetrievalOptionsBoostByDirectionDesc, InstanceNewParamsRetrievalOptionsBoostByDirectionExists, InstanceNewParamsRetrievalOptionsBoostByDirectionNotExists:
		return true
	}
	return false
}

// Controls which documents are candidates for BM25 scoring. 'and' restricts
// candidates to documents containing all query terms; 'or' includes any document
// containing at least one term, ranked by BM25 relevance. When omitted on an
// update, the existing stored value is preserved; when never set, search falls
// back to 'and'.
type InstanceNewParamsRetrievalOptionsKeywordMatchMode string

const (
	InstanceNewParamsRetrievalOptionsKeywordMatchModeAnd InstanceNewParamsRetrievalOptionsKeywordMatchMode = "and"
	InstanceNewParamsRetrievalOptionsKeywordMatchModeOr  InstanceNewParamsRetrievalOptionsKeywordMatchMode = "or"
)

func (r InstanceNewParamsRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case InstanceNewParamsRetrievalOptionsKeywordMatchModeAnd, InstanceNewParamsRetrievalOptionsKeywordMatchModeOr:
		return true
	}
	return false
}

type InstanceNewParamsSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced). Most accounts are limited to 10
	// rules; contact support to raise it.
	ExcludeItems param.Field[[]string] `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post). Most accounts are limited to 10 rules; contact
	// support to raise it.
	IncludeItems   param.Field[[]string]                                `json:"include_items"`
	Prefix         param.Field[string]                                  `json:"prefix"`
	R2Jurisdiction param.Field[string]                                  `json:"r2_jurisdiction"`
	WebCrawler     param.Field[InstanceNewParamsSourceParamsWebCrawler] `json:"web_crawler"`
}

func (r InstanceNewParamsSourceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceNewParamsSourceParamsWebCrawler struct {
	// Options for parse_type 'discover', where Browser Run discovers URLs by link
	// following and sitemaps. Ignored for 'sitemap'.
	DiscoverOptions param.Field[InstanceNewParamsSourceParamsWebCrawlerDiscoverOptions] `json:"discover_options"`
	ParseOptions    param.Field[InstanceNewParamsSourceParamsWebCrawlerParseOptions]    `json:"parse_options"`
	// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
	// recursively and requires the source to be a Verified zone on this account.
	ParseType param.Field[InstanceNewParamsSourceParamsWebCrawlerParseType] `json:"parse_type"`
}

func (r InstanceNewParamsSourceParamsWebCrawler) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Options for parse_type 'discover', where Browser Run discovers URLs by link
// following and sitemaps. Ignored for 'sitemap'.
type InstanceNewParamsSourceParamsWebCrawlerDiscoverOptions struct {
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
	Source param.Field[InstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSource] `json:"source"`
}

func (r InstanceNewParamsSourceParamsWebCrawlerDiscoverOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Where the crawler looks for URLs: 'sitemaps' reads sitemap XML only, 'links'
// follows page links only, 'all' does both.
type InstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSource string

const (
	InstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSourceAll      InstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSource = "all"
	InstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps InstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSource = "sitemaps"
	InstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSourceLinks    InstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSource = "links"
)

func (r InstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSource) IsKnown() bool {
	switch r {
	case InstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSourceAll, InstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps, InstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSourceLinks:
		return true
	}
	return false
}

type InstanceNewParamsSourceParamsWebCrawlerParseOptions struct {
	// List of path-to-selector mappings for extracting specific content from crawled
	// pages. Each entry pairs a URL glob pattern with a CSS selector. The first
	// matching path wins. Only the matched HTML fragment is stored and indexed. Omit
	// the field to disable content selection — empty arrays are rejected.
	ContentSelector param.Field[[]InstanceNewParamsSourceParamsWebCrawlerParseOptionsContentSelector] `json:"content_selector"`
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

func (r InstanceNewParamsSourceParamsWebCrawlerParseOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceNewParamsSourceParamsWebCrawlerParseOptionsContentSelector struct {
	// Glob pattern to match against the page URL path. Uses standard glob syntax: \*
	// matches within a segment, \*\* crosses directories.
	Path param.Field[string] `json:"path" api:"required"`
	// CSS selector to extract content from pages matching the path pattern. Must not
	// contain disallowed characters (;, `, $, {, }, \). Must target a single element;
	// if multiple elements match, the selector is ignored and the full page is used.
	Selector param.Field[string] `json:"selector" api:"required"`
}

func (r InstanceNewParamsSourceParamsWebCrawlerParseOptionsContentSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
// recursively and requires the source to be a Verified zone on this account.
type InstanceNewParamsSourceParamsWebCrawlerParseType string

const (
	InstanceNewParamsSourceParamsWebCrawlerParseTypeSitemap  InstanceNewParamsSourceParamsWebCrawlerParseType = "sitemap"
	InstanceNewParamsSourceParamsWebCrawlerParseTypeDiscover InstanceNewParamsSourceParamsWebCrawlerParseType = "discover"
)

func (r InstanceNewParamsSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case InstanceNewParamsSourceParamsWebCrawlerParseTypeSitemap, InstanceNewParamsSourceParamsWebCrawlerParseTypeDiscover:
		return true
	}
	return false
}

// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
type InstanceNewParamsSyncInterval float64

const (
	InstanceNewParamsSyncInterval900   InstanceNewParamsSyncInterval = 900
	InstanceNewParamsSyncInterval1800  InstanceNewParamsSyncInterval = 1800
	InstanceNewParamsSyncInterval3600  InstanceNewParamsSyncInterval = 3600
	InstanceNewParamsSyncInterval7200  InstanceNewParamsSyncInterval = 7200
	InstanceNewParamsSyncInterval14400 InstanceNewParamsSyncInterval = 14400
	InstanceNewParamsSyncInterval21600 InstanceNewParamsSyncInterval = 21600
	InstanceNewParamsSyncInterval43200 InstanceNewParamsSyncInterval = 43200
	InstanceNewParamsSyncInterval86400 InstanceNewParamsSyncInterval = 86400
)

func (r InstanceNewParamsSyncInterval) IsKnown() bool {
	switch r {
	case InstanceNewParamsSyncInterval900, InstanceNewParamsSyncInterval1800, InstanceNewParamsSyncInterval3600, InstanceNewParamsSyncInterval7200, InstanceNewParamsSyncInterval14400, InstanceNewParamsSyncInterval21600, InstanceNewParamsSyncInterval43200, InstanceNewParamsSyncInterval86400:
		return true
	}
	return false
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

type InstanceNewResponseEnvelope struct {
	Result  InstanceNewResponse             `json:"result" api:"required"`
	Success bool                            `json:"success" api:"required"`
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
	AccountID      param.Field[string]                             `path:"account_id" api:"required"`
	AIGatewayID    param.Field[string]                             `json:"ai_gateway_id"`
	AISearchModel  param.Field[string]                             `json:"ai_search_model"`
	Cache          param.Field[bool]                               `json:"cache"`
	CacheThreshold param.Field[InstanceUpdateParamsCacheThreshold] `json:"cache_threshold"`
	// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
	// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
	// (72h), 518400 (6d).
	CacheTTL       param.Field[InstanceUpdateParamsCacheTTL]         `json:"cache_ttl"`
	Chunk          param.Field[bool]                                 `json:"chunk"`
	ChunkOverlap   param.Field[int64]                                `json:"chunk_overlap"`
	ChunkSize      param.Field[int64]                                `json:"chunk_size"`
	CustomMetadata param.Field[[]InstanceUpdateParamsCustomMetadata] `json:"custom_metadata"`
	EmbeddingModel param.Field[string]                               `json:"embedding_model"`
	FusionMethod   param.Field[InstanceUpdateParamsFusionMethod]     `json:"fusion_method"`
	// Controls which storage backends are used during indexing. Defaults to
	// vector-only.
	IndexMethod          param.Field[InstanceUpdateParamsIndexMethod]          `json:"index_method"`
	IndexingOptions      param.Field[InstanceUpdateParamsIndexingOptions]      `json:"indexing_options"`
	MaxNumResults        param.Field[int64]                                    `json:"max_num_results"`
	Metadata             param.Field[InstanceUpdateParamsMetadata]             `json:"metadata"`
	Paused               param.Field[bool]                                     `json:"paused"`
	PublicEndpointParams param.Field[InstanceUpdateParamsPublicEndpointParams] `json:"public_endpoint_params"`
	Reranking            param.Field[bool]                                     `json:"reranking"`
	RerankingModel       param.Field[string]                                   `json:"reranking_model"`
	RetrievalOptions     param.Field[InstanceUpdateParamsRetrievalOptions]     `json:"retrieval_options"`
	RewriteModel         param.Field[string]                                   `json:"rewrite_model"`
	RewriteQuery         param.Field[bool]                                     `json:"rewrite_query"`
	ScoreThreshold       param.Field[float64]                                  `json:"score_threshold"`
	Source               param.Field[string]                                   `json:"source"`
	SourceParams         param.Field[InstanceUpdateParamsSourceParams]         `json:"source_params"`
	Summarization        param.Field[bool]                                     `json:"summarization"`
	SummarizationModel   param.Field[string]                                   `json:"summarization_model"`
	// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
	// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
	SyncInterval                   param.Field[InstanceUpdateParamsSyncInterval] `json:"sync_interval"`
	SystemPromptAISearch           param.Field[string]                           `json:"system_prompt_ai_search"`
	SystemPromptIndexSummarization param.Field[string]                           `json:"system_prompt_index_summarization"`
	SystemPromptRewriteQuery       param.Field[string]                           `json:"system_prompt_rewrite_query"`
	TokenID                        param.Field[string]                           `json:"token_id" format:"uuid"`
}

func (r InstanceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
// (72h), 518400 (6d).
type InstanceUpdateParamsCacheTTL float64

const (
	InstanceUpdateParamsCacheTTL600    InstanceUpdateParamsCacheTTL = 600
	InstanceUpdateParamsCacheTTL1800   InstanceUpdateParamsCacheTTL = 1800
	InstanceUpdateParamsCacheTTL3600   InstanceUpdateParamsCacheTTL = 3600
	InstanceUpdateParamsCacheTTL7200   InstanceUpdateParamsCacheTTL = 7200
	InstanceUpdateParamsCacheTTL21600  InstanceUpdateParamsCacheTTL = 21600
	InstanceUpdateParamsCacheTTL43200  InstanceUpdateParamsCacheTTL = 43200
	InstanceUpdateParamsCacheTTL86400  InstanceUpdateParamsCacheTTL = 86400
	InstanceUpdateParamsCacheTTL172800 InstanceUpdateParamsCacheTTL = 172800
	InstanceUpdateParamsCacheTTL259200 InstanceUpdateParamsCacheTTL = 259200
	InstanceUpdateParamsCacheTTL518400 InstanceUpdateParamsCacheTTL = 518400
)

func (r InstanceUpdateParamsCacheTTL) IsKnown() bool {
	switch r {
	case InstanceUpdateParamsCacheTTL600, InstanceUpdateParamsCacheTTL1800, InstanceUpdateParamsCacheTTL3600, InstanceUpdateParamsCacheTTL7200, InstanceUpdateParamsCacheTTL21600, InstanceUpdateParamsCacheTTL43200, InstanceUpdateParamsCacheTTL86400, InstanceUpdateParamsCacheTTL172800, InstanceUpdateParamsCacheTTL259200, InstanceUpdateParamsCacheTTL518400:
		return true
	}
	return false
}

type InstanceUpdateParamsCustomMetadata struct {
	DataType  param.Field[InstanceUpdateParamsCustomMetadataDataType] `json:"data_type" api:"required"`
	FieldName param.Field[string]                                     `json:"field_name" api:"required"`
}

func (r InstanceUpdateParamsCustomMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceUpdateParamsCustomMetadataDataType string

const (
	InstanceUpdateParamsCustomMetadataDataTypeText     InstanceUpdateParamsCustomMetadataDataType = "text"
	InstanceUpdateParamsCustomMetadataDataTypeNumber   InstanceUpdateParamsCustomMetadataDataType = "number"
	InstanceUpdateParamsCustomMetadataDataTypeBoolean  InstanceUpdateParamsCustomMetadataDataType = "boolean"
	InstanceUpdateParamsCustomMetadataDataTypeDatetime InstanceUpdateParamsCustomMetadataDataType = "datetime"
)

func (r InstanceUpdateParamsCustomMetadataDataType) IsKnown() bool {
	switch r {
	case InstanceUpdateParamsCustomMetadataDataTypeText, InstanceUpdateParamsCustomMetadataDataTypeNumber, InstanceUpdateParamsCustomMetadataDataTypeBoolean, InstanceUpdateParamsCustomMetadataDataTypeDatetime:
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

// Controls which storage backends are used during indexing. Defaults to
// vector-only.
type InstanceUpdateParamsIndexMethod struct {
	// Enable keyword (BM25) storage backend.
	Keyword param.Field[bool] `json:"keyword" api:"required"`
	// Enable vector (embedding) storage backend.
	Vector param.Field[bool] `json:"vector" api:"required"`
}

func (r InstanceUpdateParamsIndexMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceUpdateParamsIndexingOptions struct {
	// Tokenizer used for keyword search indexing. porter provides word-level
	// tokenization with Porter stemming (good for natural language queries). trigram
	// enables character-level substring matching (good for partial matches, code,
	// identifiers). Changing this triggers a full re-index. Defaults to porter.
	KeywordTokenizer param.Field[InstanceUpdateParamsIndexingOptionsKeywordTokenizer] `json:"keyword_tokenizer"`
}

func (r InstanceUpdateParamsIndexingOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Tokenizer used for keyword search indexing. porter provides word-level
// tokenization with Porter stemming (good for natural language queries). trigram
// enables character-level substring matching (good for partial matches, code,
// identifiers). Changing this triggers a full re-index. Defaults to porter.
type InstanceUpdateParamsIndexingOptionsKeywordTokenizer string

const (
	InstanceUpdateParamsIndexingOptionsKeywordTokenizerPorter  InstanceUpdateParamsIndexingOptionsKeywordTokenizer = "porter"
	InstanceUpdateParamsIndexingOptionsKeywordTokenizerTrigram InstanceUpdateParamsIndexingOptionsKeywordTokenizer = "trigram"
)

func (r InstanceUpdateParamsIndexingOptionsKeywordTokenizer) IsKnown() bool {
	switch r {
	case InstanceUpdateParamsIndexingOptionsKeywordTokenizerPorter, InstanceUpdateParamsIndexingOptionsKeywordTokenizerTrigram:
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
	DefaultDomainEnabled param.Field[bool]                                                   `json:"default_domain_enabled"`
	Enabled              param.Field[bool]                                                   `json:"enabled"`
	Mcp                  param.Field[InstanceUpdateParamsPublicEndpointParamsMcp]            `json:"mcp"`
	RateLimit            param.Field[InstanceUpdateParamsPublicEndpointParamsRateLimit]      `json:"rate_limit"`
	SearchEndpoint       param.Field[InstanceUpdateParamsPublicEndpointParamsSearchEndpoint] `json:"search_endpoint"`
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

type InstanceUpdateParamsRetrievalOptions struct {
	// Metadata fields to boost search results by. Each entry specifies a metadata
	// field and an optional direction. Direction defaults to 'asc' for
	// numeric/datetime fields and 'exists' for text/boolean fields. Fields must match
	// 'timestamp' or a defined custom_metadata field.
	BoostBy param.Field[[]InstanceUpdateParamsRetrievalOptionsBoostBy] `json:"boost_by"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. When omitted on an
	// update, the existing stored value is preserved; when never set, search falls
	// back to 'and'.
	KeywordMatchMode param.Field[InstanceUpdateParamsRetrievalOptionsKeywordMatchMode] `json:"keyword_match_mode"`
}

func (r InstanceUpdateParamsRetrievalOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceUpdateParamsRetrievalOptionsBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support all four directions
	// (asc, desc, exists, not_exists); text/boolean fields only support
	// exists/not_exists.
	Field param.Field[string] `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction param.Field[InstanceUpdateParamsRetrievalOptionsBoostByDirection] `json:"direction"`
}

func (r InstanceUpdateParamsRetrievalOptionsBoostBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
// for numeric/datetime fields, 'exists' for text/boolean fields.
type InstanceUpdateParamsRetrievalOptionsBoostByDirection string

const (
	InstanceUpdateParamsRetrievalOptionsBoostByDirectionAsc       InstanceUpdateParamsRetrievalOptionsBoostByDirection = "asc"
	InstanceUpdateParamsRetrievalOptionsBoostByDirectionDesc      InstanceUpdateParamsRetrievalOptionsBoostByDirection = "desc"
	InstanceUpdateParamsRetrievalOptionsBoostByDirectionExists    InstanceUpdateParamsRetrievalOptionsBoostByDirection = "exists"
	InstanceUpdateParamsRetrievalOptionsBoostByDirectionNotExists InstanceUpdateParamsRetrievalOptionsBoostByDirection = "not_exists"
)

func (r InstanceUpdateParamsRetrievalOptionsBoostByDirection) IsKnown() bool {
	switch r {
	case InstanceUpdateParamsRetrievalOptionsBoostByDirectionAsc, InstanceUpdateParamsRetrievalOptionsBoostByDirectionDesc, InstanceUpdateParamsRetrievalOptionsBoostByDirectionExists, InstanceUpdateParamsRetrievalOptionsBoostByDirectionNotExists:
		return true
	}
	return false
}

// Controls which documents are candidates for BM25 scoring. 'and' restricts
// candidates to documents containing all query terms; 'or' includes any document
// containing at least one term, ranked by BM25 relevance. When omitted on an
// update, the existing stored value is preserved; when never set, search falls
// back to 'and'.
type InstanceUpdateParamsRetrievalOptionsKeywordMatchMode string

const (
	InstanceUpdateParamsRetrievalOptionsKeywordMatchModeAnd InstanceUpdateParamsRetrievalOptionsKeywordMatchMode = "and"
	InstanceUpdateParamsRetrievalOptionsKeywordMatchModeOr  InstanceUpdateParamsRetrievalOptionsKeywordMatchMode = "or"
)

func (r InstanceUpdateParamsRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case InstanceUpdateParamsRetrievalOptionsKeywordMatchModeAnd, InstanceUpdateParamsRetrievalOptionsKeywordMatchModeOr:
		return true
	}
	return false
}

type InstanceUpdateParamsSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced). Most accounts are limited to 10
	// rules; contact support to raise it.
	ExcludeItems param.Field[[]string] `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post). Most accounts are limited to 10 rules; contact
	// support to raise it.
	IncludeItems   param.Field[[]string]                                   `json:"include_items"`
	Prefix         param.Field[string]                                     `json:"prefix"`
	R2Jurisdiction param.Field[string]                                     `json:"r2_jurisdiction"`
	WebCrawler     param.Field[InstanceUpdateParamsSourceParamsWebCrawler] `json:"web_crawler"`
}

func (r InstanceUpdateParamsSourceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceUpdateParamsSourceParamsWebCrawler struct {
	// Options for parse_type 'discover', where Browser Run discovers URLs by link
	// following and sitemaps. Ignored for 'sitemap'.
	DiscoverOptions param.Field[InstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptions] `json:"discover_options"`
	ParseOptions    param.Field[InstanceUpdateParamsSourceParamsWebCrawlerParseOptions]    `json:"parse_options"`
	// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
	// recursively and requires the source to be a Verified zone on this account.
	ParseType param.Field[InstanceUpdateParamsSourceParamsWebCrawlerParseType] `json:"parse_type"`
}

func (r InstanceUpdateParamsSourceParamsWebCrawler) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Options for parse_type 'discover', where Browser Run discovers URLs by link
// following and sitemaps. Ignored for 'sitemap'.
type InstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptions struct {
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
	Source param.Field[InstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSource] `json:"source"`
}

func (r InstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Where the crawler looks for URLs: 'sitemaps' reads sitemap XML only, 'links'
// follows page links only, 'all' does both.
type InstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSource string

const (
	InstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSourceAll      InstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSource = "all"
	InstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps InstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSource = "sitemaps"
	InstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSourceLinks    InstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSource = "links"
)

func (r InstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSource) IsKnown() bool {
	switch r {
	case InstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSourceAll, InstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps, InstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSourceLinks:
		return true
	}
	return false
}

type InstanceUpdateParamsSourceParamsWebCrawlerParseOptions struct {
	// List of path-to-selector mappings for extracting specific content from crawled
	// pages. Each entry pairs a URL glob pattern with a CSS selector. The first
	// matching path wins. Only the matched HTML fragment is stored and indexed. Omit
	// the field to disable content selection — empty arrays are rejected.
	ContentSelector param.Field[[]InstanceUpdateParamsSourceParamsWebCrawlerParseOptionsContentSelector] `json:"content_selector"`
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

func (r InstanceUpdateParamsSourceParamsWebCrawlerParseOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceUpdateParamsSourceParamsWebCrawlerParseOptionsContentSelector struct {
	// Glob pattern to match against the page URL path. Uses standard glob syntax: \*
	// matches within a segment, \*\* crosses directories.
	Path param.Field[string] `json:"path" api:"required"`
	// CSS selector to extract content from pages matching the path pattern. Must not
	// contain disallowed characters (;, `, $, {, }, \). Must target a single element;
	// if multiple elements match, the selector is ignored and the full page is used.
	Selector param.Field[string] `json:"selector" api:"required"`
}

func (r InstanceUpdateParamsSourceParamsWebCrawlerParseOptionsContentSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
// recursively and requires the source to be a Verified zone on this account.
type InstanceUpdateParamsSourceParamsWebCrawlerParseType string

const (
	InstanceUpdateParamsSourceParamsWebCrawlerParseTypeSitemap  InstanceUpdateParamsSourceParamsWebCrawlerParseType = "sitemap"
	InstanceUpdateParamsSourceParamsWebCrawlerParseTypeDiscover InstanceUpdateParamsSourceParamsWebCrawlerParseType = "discover"
)

func (r InstanceUpdateParamsSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case InstanceUpdateParamsSourceParamsWebCrawlerParseTypeSitemap, InstanceUpdateParamsSourceParamsWebCrawlerParseTypeDiscover:
		return true
	}
	return false
}

// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
type InstanceUpdateParamsSyncInterval float64

const (
	InstanceUpdateParamsSyncInterval900   InstanceUpdateParamsSyncInterval = 900
	InstanceUpdateParamsSyncInterval1800  InstanceUpdateParamsSyncInterval = 1800
	InstanceUpdateParamsSyncInterval3600  InstanceUpdateParamsSyncInterval = 3600
	InstanceUpdateParamsSyncInterval7200  InstanceUpdateParamsSyncInterval = 7200
	InstanceUpdateParamsSyncInterval14400 InstanceUpdateParamsSyncInterval = 14400
	InstanceUpdateParamsSyncInterval21600 InstanceUpdateParamsSyncInterval = 21600
	InstanceUpdateParamsSyncInterval43200 InstanceUpdateParamsSyncInterval = 43200
	InstanceUpdateParamsSyncInterval86400 InstanceUpdateParamsSyncInterval = 86400
)

func (r InstanceUpdateParamsSyncInterval) IsKnown() bool {
	switch r {
	case InstanceUpdateParamsSyncInterval900, InstanceUpdateParamsSyncInterval1800, InstanceUpdateParamsSyncInterval3600, InstanceUpdateParamsSyncInterval7200, InstanceUpdateParamsSyncInterval14400, InstanceUpdateParamsSyncInterval21600, InstanceUpdateParamsSyncInterval43200, InstanceUpdateParamsSyncInterval86400:
		return true
	}
	return false
}

type InstanceUpdateResponseEnvelope struct {
	Result  InstanceUpdateResponse             `json:"result" api:"required"`
	Success bool                               `json:"success" api:"required"`
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
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter by namespace.
	Namespace param.Field[string] `query:"namespace"`
	// Field to order results by.
	OrderBy param.Field[InstanceListParamsOrderBy] `query:"order_by"`
	// Order direction.
	OrderByDirection param.Field[InstanceListParamsOrderByDirection] `query:"order_by_direction"`
	// Page number (1-indexed).
	Page param.Field[int64] `query:"page"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter instances whose id contains this string (case-insensitive).
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [InstanceListParams]'s query parameters as `url.Values`.
func (r InstanceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Field to order results by.
type InstanceListParamsOrderBy string

const (
	InstanceListParamsOrderByCreatedAt InstanceListParamsOrderBy = "created_at"
)

func (r InstanceListParamsOrderBy) IsKnown() bool {
	switch r {
	case InstanceListParamsOrderByCreatedAt:
		return true
	}
	return false
}

// Order direction.
type InstanceListParamsOrderByDirection string

const (
	InstanceListParamsOrderByDirectionAsc  InstanceListParamsOrderByDirection = "asc"
	InstanceListParamsOrderByDirectionDesc InstanceListParamsOrderByDirection = "desc"
)

func (r InstanceListParamsOrderByDirection) IsKnown() bool {
	switch r {
	case InstanceListParamsOrderByDirectionAsc, InstanceListParamsOrderByDirectionDesc:
		return true
	}
	return false
}

type InstanceDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type InstanceDeleteResponseEnvelope struct {
	Result  InstanceDeleteResponse             `json:"result" api:"required"`
	Success bool                               `json:"success" api:"required"`
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
	AccountID       param.Field[string]                                       `path:"account_id" api:"required"`
	Messages        param.Field[[]InstanceChatCompletionsParamsMessage]       `json:"messages" api:"required"`
	AISearchOptions param.Field[InstanceChatCompletionsParamsAISearchOptions] `json:"ai_search_options"`
	Model           param.Field[string]                                       `json:"model"`
	Stream          param.Field[bool]                                         `json:"stream"`
}

func (r InstanceChatCompletionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceChatCompletionsParamsMessage struct {
	Content     param.Field[InstanceChatCompletionsParamsMessagesContentUnion] `json:"content" api:"required"`
	Role        param.Field[InstanceChatCompletionsParamsMessagesRole]         `json:"role" api:"required"`
	ExtraFields map[string]interface{}                                         `json:"-,extras"`
}

func (r InstanceChatCompletionsParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [ai_search.InstanceChatCompletionsParamsMessagesContentArray].
type InstanceChatCompletionsParamsMessagesContentUnion interface {
	ImplementsInstanceChatCompletionsParamsMessagesContentUnion()
}

type InstanceChatCompletionsParamsMessagesContentArray []InstanceChatCompletionsParamsMessagesContentArrayItemUnion

func (r InstanceChatCompletionsParamsMessagesContentArray) ImplementsInstanceChatCompletionsParamsMessagesContentUnion() {
}

type InstanceChatCompletionsParamsMessagesContentArrayItem struct {
	Type     param.Field[InstanceChatCompletionsParamsMessagesContentArrayType] `json:"type" api:"required"`
	File     param.Field[interface{}]                                           `json:"file"`
	ImageURL param.Field[interface{}]                                           `json:"image_url"`
	Text     param.Field[string]                                                `json:"text"`
}

func (r InstanceChatCompletionsParamsMessagesContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InstanceChatCompletionsParamsMessagesContentArrayItem) implementsInstanceChatCompletionsParamsMessagesContentArrayItemUnion() {
}

// Satisfied by
// [ai_search.InstanceChatCompletionsParamsMessagesContentArrayObject],
// [ai_search.InstanceChatCompletionsParamsMessagesContentArrayObject],
// [ai_search.InstanceChatCompletionsParamsMessagesContentArrayObject],
// [InstanceChatCompletionsParamsMessagesContentArrayItem].
type InstanceChatCompletionsParamsMessagesContentArrayItemUnion interface {
	implementsInstanceChatCompletionsParamsMessagesContentArrayItemUnion()
}

type InstanceChatCompletionsParamsMessagesContentArrayObject struct {
	Text param.Field[string]                                                      `json:"text" api:"required"`
	Type param.Field[InstanceChatCompletionsParamsMessagesContentArrayObjectType] `json:"type" api:"required"`
}

func (r InstanceChatCompletionsParamsMessagesContentArrayObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InstanceChatCompletionsParamsMessagesContentArrayObject) implementsInstanceChatCompletionsParamsMessagesContentArrayItemUnion() {
}

type InstanceChatCompletionsParamsMessagesContentArrayObjectType string

const (
	InstanceChatCompletionsParamsMessagesContentArrayObjectTypeText InstanceChatCompletionsParamsMessagesContentArrayObjectType = "text"
)

func (r InstanceChatCompletionsParamsMessagesContentArrayObjectType) IsKnown() bool {
	switch r {
	case InstanceChatCompletionsParamsMessagesContentArrayObjectTypeText:
		return true
	}
	return false
}

type InstanceChatCompletionsParamsMessagesContentArrayType string

const (
	InstanceChatCompletionsParamsMessagesContentArrayTypeText     InstanceChatCompletionsParamsMessagesContentArrayType = "text"
	InstanceChatCompletionsParamsMessagesContentArrayTypeImageURL InstanceChatCompletionsParamsMessagesContentArrayType = "image_url"
	InstanceChatCompletionsParamsMessagesContentArrayTypeFile     InstanceChatCompletionsParamsMessagesContentArrayType = "file"
)

func (r InstanceChatCompletionsParamsMessagesContentArrayType) IsKnown() bool {
	switch r {
	case InstanceChatCompletionsParamsMessagesContentArrayTypeText, InstanceChatCompletionsParamsMessagesContentArrayTypeImageURL, InstanceChatCompletionsParamsMessagesContentArrayTypeFile:
		return true
	}
	return false
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
	Cache        param.Field[InstanceChatCompletionsParamsAISearchOptionsCache]        `json:"cache"`
	QueryRewrite param.Field[InstanceChatCompletionsParamsAISearchOptionsQueryRewrite] `json:"query_rewrite"`
	Reranking    param.Field[InstanceChatCompletionsParamsAISearchOptionsReranking]    `json:"reranking"`
	Retrieval    param.Field[InstanceChatCompletionsParamsAISearchOptionsRetrieval]    `json:"retrieval"`
}

func (r InstanceChatCompletionsParamsAISearchOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceChatCompletionsParamsAISearchOptionsCache struct {
	CacheThreshold param.Field[InstanceChatCompletionsParamsAISearchOptionsCacheCacheThreshold] `json:"cache_threshold"`
	Enabled        param.Field[bool]                                                            `json:"enabled"`
}

func (r InstanceChatCompletionsParamsAISearchOptionsCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceChatCompletionsParamsAISearchOptionsCacheCacheThreshold string

const (
	InstanceChatCompletionsParamsAISearchOptionsCacheCacheThresholdSuperStrictMatch InstanceChatCompletionsParamsAISearchOptionsCacheCacheThreshold = "super_strict_match"
	InstanceChatCompletionsParamsAISearchOptionsCacheCacheThresholdCloseEnough      InstanceChatCompletionsParamsAISearchOptionsCacheCacheThreshold = "close_enough"
	InstanceChatCompletionsParamsAISearchOptionsCacheCacheThresholdFlexibleFriend   InstanceChatCompletionsParamsAISearchOptionsCacheCacheThreshold = "flexible_friend"
	InstanceChatCompletionsParamsAISearchOptionsCacheCacheThresholdAnythingGoes     InstanceChatCompletionsParamsAISearchOptionsCacheCacheThreshold = "anything_goes"
)

func (r InstanceChatCompletionsParamsAISearchOptionsCacheCacheThreshold) IsKnown() bool {
	switch r {
	case InstanceChatCompletionsParamsAISearchOptionsCacheCacheThresholdSuperStrictMatch, InstanceChatCompletionsParamsAISearchOptionsCacheCacheThresholdCloseEnough, InstanceChatCompletionsParamsAISearchOptionsCacheCacheThresholdFlexibleFriend, InstanceChatCompletionsParamsAISearchOptionsCacheCacheThresholdAnythingGoes:
		return true
	}
	return false
}

type InstanceChatCompletionsParamsAISearchOptionsQueryRewrite struct {
	Enabled       param.Field[bool]   `json:"enabled"`
	Model         param.Field[string] `json:"model"`
	RewritePrompt param.Field[string] `json:"rewrite_prompt"`
}

func (r InstanceChatCompletionsParamsAISearchOptionsQueryRewrite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceChatCompletionsParamsAISearchOptionsReranking struct {
	Enabled        param.Field[bool]    `json:"enabled"`
	MatchThreshold param.Field[float64] `json:"match_threshold"`
	Model          param.Field[string]  `json:"model"`
}

func (r InstanceChatCompletionsParamsAISearchOptionsReranking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceChatCompletionsParamsAISearchOptionsRetrieval struct {
	// Metadata fields to boost search results by. Overrides the instance-level
	// boost_by config. Direction defaults to 'asc' for numeric/datetime fields,
	// 'exists' for text/boolean fields. Fields must match 'timestamp' or a defined
	// custom_metadata field.
	BoostBy          param.Field[[]InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostBy]    `json:"boost_by"`
	ContextExpansion param.Field[int64]                                                             `json:"context_expansion"`
	Filters          param.Field[map[string]interface{}]                                            `json:"filters"`
	FusionMethod     param.Field[InstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethod] `json:"fusion_method"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. When omitted, falls back
	// to the instance-level retrieval_options.keyword_match_mode, then to 'and'.
	KeywordMatchMode param.Field[InstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode] `json:"keyword_match_mode"`
	MatchThreshold   param.Field[float64]                                                               `json:"match_threshold"`
	MaxNumResults    param.Field[int64]                                                                 `json:"max_num_results"`
	RetrievalType    param.Field[InstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalType]    `json:"retrieval_type"`
	ReturnOnFailure  param.Field[bool]                                                                  `json:"return_on_failure"`
}

func (r InstanceChatCompletionsParamsAISearchOptionsRetrieval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support all four directions
	// (asc, desc, exists, not_exists); text/boolean fields only support
	// exists/not_exists.
	Field param.Field[string] `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction param.Field[InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection] `json:"direction"`
}

func (r InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
// for numeric/datetime fields, 'exists' for text/boolean fields.
type InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection string

const (
	InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionAsc       InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection = "asc"
	InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionDesc      InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection = "desc"
	InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionExists    InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection = "exists"
	InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionNotExists InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection = "not_exists"
)

func (r InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection) IsKnown() bool {
	switch r {
	case InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionAsc, InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionDesc, InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionExists, InstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionNotExists:
		return true
	}
	return false
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

// Controls which documents are candidates for BM25 scoring. 'and' restricts
// candidates to documents containing all query terms; 'or' includes any document
// containing at least one term, ranked by BM25 relevance. When omitted, falls back
// to the instance-level retrieval_options.keyword_match_mode, then to 'and'.
type InstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode string

const (
	InstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchModeAnd InstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode = "and"
	InstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchModeOr  InstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode = "or"
)

func (r InstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode) IsKnown() bool {
	switch r {
	case InstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchModeAnd, InstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchModeOr:
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

type InstanceReadParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type InstanceReadResponseEnvelope struct {
	Result  InstanceReadResponse             `json:"result" api:"required"`
	Success bool                             `json:"success" api:"required"`
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
	AccountID       param.Field[string]                              `path:"account_id" api:"required"`
	AISearchOptions param.Field[InstanceSearchParamsAISearchOptions] `json:"ai_search_options"`
	// OpenAI-compatible message array. For multimodal queries, set the last user
	// message's `content` to an array of typed parts:
	// `[{type:'text', text:'…'}, {type:'image_url', image_url:{url:'…'}}]`. Image
	// inputs require the RAG's embedding_model to declare 'image' in
	// supported_modalities.
	Messages param.Field[[]InstanceSearchParamsMessage] `json:"messages"`
	// A simple text query string. Alternative to 'messages' — provide either this or
	// 'messages', not both.
	Query param.Field[string] `json:"query"`
}

func (r InstanceSearchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceSearchParamsAISearchOptions struct {
	Cache        param.Field[InstanceSearchParamsAISearchOptionsCache]        `json:"cache"`
	QueryRewrite param.Field[InstanceSearchParamsAISearchOptionsQueryRewrite] `json:"query_rewrite"`
	Reranking    param.Field[InstanceSearchParamsAISearchOptionsReranking]    `json:"reranking"`
	Retrieval    param.Field[InstanceSearchParamsAISearchOptionsRetrieval]    `json:"retrieval"`
}

func (r InstanceSearchParamsAISearchOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceSearchParamsAISearchOptionsCache struct {
	CacheThreshold param.Field[InstanceSearchParamsAISearchOptionsCacheCacheThreshold] `json:"cache_threshold"`
	Enabled        param.Field[bool]                                                   `json:"enabled"`
}

func (r InstanceSearchParamsAISearchOptionsCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceSearchParamsAISearchOptionsCacheCacheThreshold string

const (
	InstanceSearchParamsAISearchOptionsCacheCacheThresholdSuperStrictMatch InstanceSearchParamsAISearchOptionsCacheCacheThreshold = "super_strict_match"
	InstanceSearchParamsAISearchOptionsCacheCacheThresholdCloseEnough      InstanceSearchParamsAISearchOptionsCacheCacheThreshold = "close_enough"
	InstanceSearchParamsAISearchOptionsCacheCacheThresholdFlexibleFriend   InstanceSearchParamsAISearchOptionsCacheCacheThreshold = "flexible_friend"
	InstanceSearchParamsAISearchOptionsCacheCacheThresholdAnythingGoes     InstanceSearchParamsAISearchOptionsCacheCacheThreshold = "anything_goes"
)

func (r InstanceSearchParamsAISearchOptionsCacheCacheThreshold) IsKnown() bool {
	switch r {
	case InstanceSearchParamsAISearchOptionsCacheCacheThresholdSuperStrictMatch, InstanceSearchParamsAISearchOptionsCacheCacheThresholdCloseEnough, InstanceSearchParamsAISearchOptionsCacheCacheThresholdFlexibleFriend, InstanceSearchParamsAISearchOptionsCacheCacheThresholdAnythingGoes:
		return true
	}
	return false
}

type InstanceSearchParamsAISearchOptionsQueryRewrite struct {
	Enabled       param.Field[bool]   `json:"enabled"`
	Model         param.Field[string] `json:"model"`
	RewritePrompt param.Field[string] `json:"rewrite_prompt"`
}

func (r InstanceSearchParamsAISearchOptionsQueryRewrite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceSearchParamsAISearchOptionsReranking struct {
	Enabled        param.Field[bool]    `json:"enabled"`
	MatchThreshold param.Field[float64] `json:"match_threshold"`
	Model          param.Field[string]  `json:"model"`
}

func (r InstanceSearchParamsAISearchOptionsReranking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceSearchParamsAISearchOptionsRetrieval struct {
	// Metadata fields to boost search results by. Overrides the instance-level
	// boost_by config. Direction defaults to 'asc' for numeric/datetime fields,
	// 'exists' for text/boolean fields. Fields must match 'timestamp' or a defined
	// custom_metadata field.
	BoostBy          param.Field[[]InstanceSearchParamsAISearchOptionsRetrievalBoostBy]    `json:"boost_by"`
	ContextExpansion param.Field[int64]                                                    `json:"context_expansion"`
	Filters          param.Field[map[string]interface{}]                                   `json:"filters"`
	FusionMethod     param.Field[InstanceSearchParamsAISearchOptionsRetrievalFusionMethod] `json:"fusion_method"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. When omitted, falls back
	// to the instance-level retrieval_options.keyword_match_mode, then to 'and'.
	KeywordMatchMode param.Field[InstanceSearchParamsAISearchOptionsRetrievalKeywordMatchMode] `json:"keyword_match_mode"`
	MatchThreshold   param.Field[float64]                                                      `json:"match_threshold"`
	MaxNumResults    param.Field[int64]                                                        `json:"max_num_results"`
	RetrievalType    param.Field[InstanceSearchParamsAISearchOptionsRetrievalRetrievalType]    `json:"retrieval_type"`
	ReturnOnFailure  param.Field[bool]                                                         `json:"return_on_failure"`
}

func (r InstanceSearchParamsAISearchOptionsRetrieval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceSearchParamsAISearchOptionsRetrievalBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support all four directions
	// (asc, desc, exists, not_exists); text/boolean fields only support
	// exists/not_exists.
	Field param.Field[string] `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction param.Field[InstanceSearchParamsAISearchOptionsRetrievalBoostByDirection] `json:"direction"`
}

func (r InstanceSearchParamsAISearchOptionsRetrievalBoostBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
// for numeric/datetime fields, 'exists' for text/boolean fields.
type InstanceSearchParamsAISearchOptionsRetrievalBoostByDirection string

const (
	InstanceSearchParamsAISearchOptionsRetrievalBoostByDirectionAsc       InstanceSearchParamsAISearchOptionsRetrievalBoostByDirection = "asc"
	InstanceSearchParamsAISearchOptionsRetrievalBoostByDirectionDesc      InstanceSearchParamsAISearchOptionsRetrievalBoostByDirection = "desc"
	InstanceSearchParamsAISearchOptionsRetrievalBoostByDirectionExists    InstanceSearchParamsAISearchOptionsRetrievalBoostByDirection = "exists"
	InstanceSearchParamsAISearchOptionsRetrievalBoostByDirectionNotExists InstanceSearchParamsAISearchOptionsRetrievalBoostByDirection = "not_exists"
)

func (r InstanceSearchParamsAISearchOptionsRetrievalBoostByDirection) IsKnown() bool {
	switch r {
	case InstanceSearchParamsAISearchOptionsRetrievalBoostByDirectionAsc, InstanceSearchParamsAISearchOptionsRetrievalBoostByDirectionDesc, InstanceSearchParamsAISearchOptionsRetrievalBoostByDirectionExists, InstanceSearchParamsAISearchOptionsRetrievalBoostByDirectionNotExists:
		return true
	}
	return false
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

// Controls which documents are candidates for BM25 scoring. 'and' restricts
// candidates to documents containing all query terms; 'or' includes any document
// containing at least one term, ranked by BM25 relevance. When omitted, falls back
// to the instance-level retrieval_options.keyword_match_mode, then to 'and'.
type InstanceSearchParamsAISearchOptionsRetrievalKeywordMatchMode string

const (
	InstanceSearchParamsAISearchOptionsRetrievalKeywordMatchModeAnd InstanceSearchParamsAISearchOptionsRetrievalKeywordMatchMode = "and"
	InstanceSearchParamsAISearchOptionsRetrievalKeywordMatchModeOr  InstanceSearchParamsAISearchOptionsRetrievalKeywordMatchMode = "or"
)

func (r InstanceSearchParamsAISearchOptionsRetrievalKeywordMatchMode) IsKnown() bool {
	switch r {
	case InstanceSearchParamsAISearchOptionsRetrievalKeywordMatchModeAnd, InstanceSearchParamsAISearchOptionsRetrievalKeywordMatchModeOr:
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

type InstanceSearchParamsMessage struct {
	Content     param.Field[InstanceSearchParamsMessagesContentUnion] `json:"content" api:"required"`
	Role        param.Field[InstanceSearchParamsMessagesRole]         `json:"role" api:"required"`
	ExtraFields map[string]interface{}                                `json:"-,extras"`
}

func (r InstanceSearchParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [ai_search.InstanceSearchParamsMessagesContentArray].
type InstanceSearchParamsMessagesContentUnion interface {
	ImplementsInstanceSearchParamsMessagesContentUnion()
}

type InstanceSearchParamsMessagesContentArray []InstanceSearchParamsMessagesContentArrayItemUnion

func (r InstanceSearchParamsMessagesContentArray) ImplementsInstanceSearchParamsMessagesContentUnion() {
}

type InstanceSearchParamsMessagesContentArrayItem struct {
	Type     param.Field[InstanceSearchParamsMessagesContentArrayType] `json:"type" api:"required"`
	File     param.Field[interface{}]                                  `json:"file"`
	ImageURL param.Field[interface{}]                                  `json:"image_url"`
	Text     param.Field[string]                                       `json:"text"`
}

func (r InstanceSearchParamsMessagesContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InstanceSearchParamsMessagesContentArrayItem) implementsInstanceSearchParamsMessagesContentArrayItemUnion() {
}

// Satisfied by [ai_search.InstanceSearchParamsMessagesContentArrayObject],
// [ai_search.InstanceSearchParamsMessagesContentArrayObject],
// [ai_search.InstanceSearchParamsMessagesContentArrayObject],
// [InstanceSearchParamsMessagesContentArrayItem].
type InstanceSearchParamsMessagesContentArrayItemUnion interface {
	implementsInstanceSearchParamsMessagesContentArrayItemUnion()
}

type InstanceSearchParamsMessagesContentArrayObject struct {
	Text param.Field[string]                                             `json:"text" api:"required"`
	Type param.Field[InstanceSearchParamsMessagesContentArrayObjectType] `json:"type" api:"required"`
}

func (r InstanceSearchParamsMessagesContentArrayObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InstanceSearchParamsMessagesContentArrayObject) implementsInstanceSearchParamsMessagesContentArrayItemUnion() {
}

type InstanceSearchParamsMessagesContentArrayObjectType string

const (
	InstanceSearchParamsMessagesContentArrayObjectTypeText InstanceSearchParamsMessagesContentArrayObjectType = "text"
)

func (r InstanceSearchParamsMessagesContentArrayObjectType) IsKnown() bool {
	switch r {
	case InstanceSearchParamsMessagesContentArrayObjectTypeText:
		return true
	}
	return false
}

type InstanceSearchParamsMessagesContentArrayType string

const (
	InstanceSearchParamsMessagesContentArrayTypeText     InstanceSearchParamsMessagesContentArrayType = "text"
	InstanceSearchParamsMessagesContentArrayTypeImageURL InstanceSearchParamsMessagesContentArrayType = "image_url"
	InstanceSearchParamsMessagesContentArrayTypeFile     InstanceSearchParamsMessagesContentArrayType = "file"
)

func (r InstanceSearchParamsMessagesContentArrayType) IsKnown() bool {
	switch r {
	case InstanceSearchParamsMessagesContentArrayTypeText, InstanceSearchParamsMessagesContentArrayTypeImageURL, InstanceSearchParamsMessagesContentArrayTypeFile:
		return true
	}
	return false
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

type InstanceSearchResponseEnvelope struct {
	Result  InstanceSearchResponse             `json:"result" api:"required"`
	Success bool                               `json:"success" api:"required"`
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
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type InstanceStatsResponseEnvelope struct {
	Result  InstanceStatsResponse             `json:"result" api:"required"`
	Success bool                              `json:"success" api:"required"`
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
