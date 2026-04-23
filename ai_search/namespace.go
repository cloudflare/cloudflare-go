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

// Create a new namespace.
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

// Update namespace.
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

// List namespaces.
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

// List namespaces.
func (r *NamespaceService) ListAutoPaging(ctx context.Context, params NamespaceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[NamespaceListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete namespace.
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

// Read namespace.
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

// Multi-Instance Search
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
	Description string                   `json:"description" api:"nullable"`
	JSON        namespaceNewResponseJSON `json:"-"`
}

// namespaceNewResponseJSON contains the JSON metadata for the struct
// [NamespaceNewResponse]
type namespaceNewResponseJSON struct {
	CreatedAt   apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceNewResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceUpdateResponse struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Name      string    `json:"name" api:"required"`
	// Optional description for the namespace. Max 256 characters.
	Description string                      `json:"description" api:"nullable"`
	JSON        namespaceUpdateResponseJSON `json:"-"`
}

// namespaceUpdateResponseJSON contains the JSON metadata for the struct
// [NamespaceUpdateResponse]
type namespaceUpdateResponseJSON struct {
	CreatedAt   apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceListResponse struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Name      string    `json:"name" api:"required"`
	// Optional description for the namespace. Max 256 characters.
	Description string                    `json:"description" api:"nullable"`
	JSON        namespaceListResponseJSON `json:"-"`
}

// namespaceListResponseJSON contains the JSON metadata for the struct
// [NamespaceListResponse]
type namespaceListResponseJSON struct {
	CreatedAt   apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceListResponseJSON) RawJSON() string {
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
	Content     string                                             `json:"content" api:"required,nullable"`
	Role        NamespaceChatCompletionsResponseChoicesMessageRole `json:"role" api:"required"`
	ExtraFields map[string]interface{}                             `json:"-" api:"extrafields"`
	JSON        namespaceChatCompletionsResponseChoicesMessageJSON `json:"-"`
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
	Description string                    `json:"description" api:"nullable"`
	JSON        namespaceReadResponseJSON `json:"-"`
}

// namespaceReadResponseJSON contains the JSON metadata for the struct
// [NamespaceReadResponse]
type namespaceReadResponseJSON struct {
	CreatedAt   apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceReadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceReadResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceSearchResponse struct {
	Chunks      []NamespaceSearchResponseChunk `json:"chunks" api:"required"`
	SearchQuery string                         `json:"search_query" api:"required"`
	Errors      []NamespaceSearchResponseError `json:"errors"`
	JSON        namespaceSearchResponseJSON    `json:"-"`
}

// namespaceSearchResponseJSON contains the JSON metadata for the struct
// [NamespaceSearchResponse]
type namespaceSearchResponseJSON struct {
	Chunks      apijson.Field
	SearchQuery apijson.Field
	Errors      apijson.Field
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
	Description param.Field[string] `json:"description"`
}

func (r NamespaceNewParams) MarshalJSON() (data []byte, err error) {
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
	Description param.Field[string] `json:"description"`
}

func (r NamespaceUpdateParams) MarshalJSON() (data []byte, err error) {
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
	Model           param.Field[NamespaceChatCompletionsParamsModel]           `json:"model"`
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
	Enabled       param.Field[bool]                                                           `json:"enabled"`
	Model         param.Field[NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel] `json:"model"`
	RewritePrompt param.Field[string]                                                         `json:"rewrite_prompt"`
}

func (r NamespaceChatCompletionsParamsAISearchOptionsQueryRewrite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel string

const (
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfZaiOrgGlm4_7Flash                   NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFast         NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFp8          NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama4Scout17b16eInstruct       NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfQwenQwen3_30bA3bFp8                 NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2Instruct            NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfGoogleGemma3_12bIt                  NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/google/gemma-3-12b-it"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfGoogleGemma4_26bA4bIt               NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2_5                  NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_7Sonnet              NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-3-7-sonnet"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaudeSonnet4                NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-sonnet-4"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaudeOpus4                  NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-opus-4"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_5Haiku               NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-3-5-haiku"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bInstruct        NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bThinking        NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasLlama3_3_70b                  NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "cerebras/llama-3.3-70b"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Maverick17b128eInstruct NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Scout17b16eInstruct     NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasGptOSs120b                    NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "cerebras/gpt-oss-120b"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Flash          NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Pro            NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelGrokGrok4                             NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "grok/grok-4"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelGroqLlama3_3_70bVersatile             NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "groq/llama-3.3-70b-versatile"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelGroqLlama3_1_8bInstant                NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "groq/llama-3.1-8b-instant"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelOpenAIGpt5                            NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "openai/gpt-5"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Mini                        NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "openai/gpt-5-mini"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Nano                        NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "openai/gpt-5-nano"
	NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelEmpty                                 NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel = ""
)

func (r NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModel) IsKnown() bool {
	switch r {
	case NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfZaiOrgGlm4_7Flash, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFast, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFp8, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama4Scout17b16eInstruct, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfQwenQwen3_30bA3bFp8, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2Instruct, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfGoogleGemma3_12bIt, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfGoogleGemma4_26bA4bIt, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2_5, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_7Sonnet, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaudeSonnet4, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaudeOpus4, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_5Haiku, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bInstruct, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bThinking, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasLlama3_3_70b, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Maverick17b128eInstruct, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Scout17b16eInstruct, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasGptOSs120b, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Flash, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Pro, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelGrokGrok4, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelGroqLlama3_3_70bVersatile, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelGroqLlama3_1_8bInstant, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelOpenAIGpt5, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Mini, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Nano, NamespaceChatCompletionsParamsAISearchOptionsQueryRewriteModelEmpty:
		return true
	}
	return false
}

type NamespaceChatCompletionsParamsAISearchOptionsReranking struct {
	Enabled        param.Field[bool]                                                        `json:"enabled"`
	MatchThreshold param.Field[float64]                                                     `json:"match_threshold"`
	Model          param.Field[NamespaceChatCompletionsParamsAISearchOptionsRerankingModel] `json:"model"`
}

func (r NamespaceChatCompletionsParamsAISearchOptionsReranking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceChatCompletionsParamsAISearchOptionsRerankingModel string

const (
	NamespaceChatCompletionsParamsAISearchOptionsRerankingModelCfBaaiBgeRerankerBase NamespaceChatCompletionsParamsAISearchOptionsRerankingModel = "@cf/baai/bge-reranker-base"
	NamespaceChatCompletionsParamsAISearchOptionsRerankingModelEmpty                 NamespaceChatCompletionsParamsAISearchOptionsRerankingModel = ""
)

func (r NamespaceChatCompletionsParamsAISearchOptionsRerankingModel) IsKnown() bool {
	switch r {
	case NamespaceChatCompletionsParamsAISearchOptionsRerankingModelCfBaaiBgeRerankerBase, NamespaceChatCompletionsParamsAISearchOptionsRerankingModelEmpty:
		return true
	}
	return false
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
	// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
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
	// custom_metadata field. Numeric and datetime fields support asc/desc directions;
	// text/boolean fields support exists/not_exists.
	Field param.Field[string] `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction param.Field[NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection] `json:"direction"`
}

func (r NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
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
// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
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
	Content     param.Field[string]                                     `json:"content" api:"required"`
	Role        param.Field[NamespaceChatCompletionsParamsMessagesRole] `json:"role" api:"required"`
	ExtraFields map[string]interface{}                                  `json:"-,extras"`
}

func (r NamespaceChatCompletionsParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

type NamespaceChatCompletionsParamsModel string

const (
	NamespaceChatCompletionsParamsModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceChatCompletionsParamsModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceChatCompletionsParamsModelCfZaiOrgGlm4_7Flash                   NamespaceChatCompletionsParamsModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceChatCompletionsParamsModelCfMetaLlama3_1_8bInstructFast         NamespaceChatCompletionsParamsModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceChatCompletionsParamsModelCfMetaLlama3_1_8bInstructFp8          NamespaceChatCompletionsParamsModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceChatCompletionsParamsModelCfMetaLlama4Scout17b16eInstruct       NamespaceChatCompletionsParamsModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceChatCompletionsParamsModelCfQwenQwen3_30bA3bFp8                 NamespaceChatCompletionsParamsModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceChatCompletionsParamsModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceChatCompletionsParamsModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceChatCompletionsParamsModelCfMoonshotaiKimiK2Instruct            NamespaceChatCompletionsParamsModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceChatCompletionsParamsModelCfGoogleGemma3_12bIt                  NamespaceChatCompletionsParamsModel = "@cf/google/gemma-3-12b-it"
	NamespaceChatCompletionsParamsModelCfGoogleGemma4_26bA4bIt               NamespaceChatCompletionsParamsModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceChatCompletionsParamsModelCfMoonshotaiKimiK2_5                  NamespaceChatCompletionsParamsModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceChatCompletionsParamsModelAnthropicClaude3_7Sonnet              NamespaceChatCompletionsParamsModel = "anthropic/claude-3-7-sonnet"
	NamespaceChatCompletionsParamsModelAnthropicClaudeSonnet4                NamespaceChatCompletionsParamsModel = "anthropic/claude-sonnet-4"
	NamespaceChatCompletionsParamsModelAnthropicClaudeOpus4                  NamespaceChatCompletionsParamsModel = "anthropic/claude-opus-4"
	NamespaceChatCompletionsParamsModelAnthropicClaude3_5Haiku               NamespaceChatCompletionsParamsModel = "anthropic/claude-3-5-haiku"
	NamespaceChatCompletionsParamsModelCerebrasQwen3_235bA22bInstruct        NamespaceChatCompletionsParamsModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceChatCompletionsParamsModelCerebrasQwen3_235bA22bThinking        NamespaceChatCompletionsParamsModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceChatCompletionsParamsModelCerebrasLlama3_3_70b                  NamespaceChatCompletionsParamsModel = "cerebras/llama-3.3-70b"
	NamespaceChatCompletionsParamsModelCerebrasLlama4Maverick17b128eInstruct NamespaceChatCompletionsParamsModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceChatCompletionsParamsModelCerebrasLlama4Scout17b16eInstruct     NamespaceChatCompletionsParamsModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceChatCompletionsParamsModelCerebrasGptOSs120b                    NamespaceChatCompletionsParamsModel = "cerebras/gpt-oss-120b"
	NamespaceChatCompletionsParamsModelGoogleAIStudioGemini2_5Flash          NamespaceChatCompletionsParamsModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceChatCompletionsParamsModelGoogleAIStudioGemini2_5Pro            NamespaceChatCompletionsParamsModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceChatCompletionsParamsModelGrokGrok4                             NamespaceChatCompletionsParamsModel = "grok/grok-4"
	NamespaceChatCompletionsParamsModelGroqLlama3_3_70bVersatile             NamespaceChatCompletionsParamsModel = "groq/llama-3.3-70b-versatile"
	NamespaceChatCompletionsParamsModelGroqLlama3_1_8bInstant                NamespaceChatCompletionsParamsModel = "groq/llama-3.1-8b-instant"
	NamespaceChatCompletionsParamsModelOpenAIGpt5                            NamespaceChatCompletionsParamsModel = "openai/gpt-5"
	NamespaceChatCompletionsParamsModelOpenAIGpt5Mini                        NamespaceChatCompletionsParamsModel = "openai/gpt-5-mini"
	NamespaceChatCompletionsParamsModelOpenAIGpt5Nano                        NamespaceChatCompletionsParamsModel = "openai/gpt-5-nano"
	NamespaceChatCompletionsParamsModelEmpty                                 NamespaceChatCompletionsParamsModel = ""
)

func (r NamespaceChatCompletionsParamsModel) IsKnown() bool {
	switch r {
	case NamespaceChatCompletionsParamsModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceChatCompletionsParamsModelCfZaiOrgGlm4_7Flash, NamespaceChatCompletionsParamsModelCfMetaLlama3_1_8bInstructFast, NamespaceChatCompletionsParamsModelCfMetaLlama3_1_8bInstructFp8, NamespaceChatCompletionsParamsModelCfMetaLlama4Scout17b16eInstruct, NamespaceChatCompletionsParamsModelCfQwenQwen3_30bA3bFp8, NamespaceChatCompletionsParamsModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceChatCompletionsParamsModelCfMoonshotaiKimiK2Instruct, NamespaceChatCompletionsParamsModelCfGoogleGemma3_12bIt, NamespaceChatCompletionsParamsModelCfGoogleGemma4_26bA4bIt, NamespaceChatCompletionsParamsModelCfMoonshotaiKimiK2_5, NamespaceChatCompletionsParamsModelAnthropicClaude3_7Sonnet, NamespaceChatCompletionsParamsModelAnthropicClaudeSonnet4, NamespaceChatCompletionsParamsModelAnthropicClaudeOpus4, NamespaceChatCompletionsParamsModelAnthropicClaude3_5Haiku, NamespaceChatCompletionsParamsModelCerebrasQwen3_235bA22bInstruct, NamespaceChatCompletionsParamsModelCerebrasQwen3_235bA22bThinking, NamespaceChatCompletionsParamsModelCerebrasLlama3_3_70b, NamespaceChatCompletionsParamsModelCerebrasLlama4Maverick17b128eInstruct, NamespaceChatCompletionsParamsModelCerebrasLlama4Scout17b16eInstruct, NamespaceChatCompletionsParamsModelCerebrasGptOSs120b, NamespaceChatCompletionsParamsModelGoogleAIStudioGemini2_5Flash, NamespaceChatCompletionsParamsModelGoogleAIStudioGemini2_5Pro, NamespaceChatCompletionsParamsModelGrokGrok4, NamespaceChatCompletionsParamsModelGroqLlama3_3_70bVersatile, NamespaceChatCompletionsParamsModelGroqLlama3_1_8bInstant, NamespaceChatCompletionsParamsModelOpenAIGpt5, NamespaceChatCompletionsParamsModelOpenAIGpt5Mini, NamespaceChatCompletionsParamsModelOpenAIGpt5Nano, NamespaceChatCompletionsParamsModelEmpty:
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
	Messages        param.Field[[]NamespaceSearchParamsMessage]       `json:"messages"`
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
	Enabled       param.Field[bool]                                                  `json:"enabled"`
	Model         param.Field[NamespaceSearchParamsAISearchOptionsQueryRewriteModel] `json:"model"`
	RewritePrompt param.Field[string]                                                `json:"rewrite_prompt"`
}

func (r NamespaceSearchParamsAISearchOptionsQueryRewrite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceSearchParamsAISearchOptionsQueryRewriteModel string

const (
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfZaiOrgGlm4_7Flash                   NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFast         NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFp8          NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama4Scout17b16eInstruct       NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfQwenQwen3_30bA3bFp8                 NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2Instruct            NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfGoogleGemma3_12bIt                  NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/google/gemma-3-12b-it"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfGoogleGemma4_26bA4bIt               NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2_5                  NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_7Sonnet              NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-3-7-sonnet"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaudeSonnet4                NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-sonnet-4"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaudeOpus4                  NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-opus-4"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_5Haiku               NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-3-5-haiku"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bInstruct        NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bThinking        NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelCerebrasLlama3_3_70b                  NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "cerebras/llama-3.3-70b"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Maverick17b128eInstruct NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Scout17b16eInstruct     NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelCerebrasGptOSs120b                    NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "cerebras/gpt-oss-120b"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Flash          NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Pro            NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelGrokGrok4                             NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "grok/grok-4"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelGroqLlama3_3_70bVersatile             NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "groq/llama-3.3-70b-versatile"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelGroqLlama3_1_8bInstant                NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "groq/llama-3.1-8b-instant"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelOpenAIGpt5                            NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "openai/gpt-5"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Mini                        NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "openai/gpt-5-mini"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Nano                        NamespaceSearchParamsAISearchOptionsQueryRewriteModel = "openai/gpt-5-nano"
	NamespaceSearchParamsAISearchOptionsQueryRewriteModelEmpty                                 NamespaceSearchParamsAISearchOptionsQueryRewriteModel = ""
)

func (r NamespaceSearchParamsAISearchOptionsQueryRewriteModel) IsKnown() bool {
	switch r {
	case NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfZaiOrgGlm4_7Flash, NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFast, NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFp8, NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama4Scout17b16eInstruct, NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfQwenQwen3_30bA3bFp8, NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2Instruct, NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfGoogleGemma3_12bIt, NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfGoogleGemma4_26bA4bIt, NamespaceSearchParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2_5, NamespaceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_7Sonnet, NamespaceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaudeSonnet4, NamespaceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaudeOpus4, NamespaceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_5Haiku, NamespaceSearchParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bInstruct, NamespaceSearchParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bThinking, NamespaceSearchParamsAISearchOptionsQueryRewriteModelCerebrasLlama3_3_70b, NamespaceSearchParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Maverick17b128eInstruct, NamespaceSearchParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Scout17b16eInstruct, NamespaceSearchParamsAISearchOptionsQueryRewriteModelCerebrasGptOSs120b, NamespaceSearchParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Flash, NamespaceSearchParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Pro, NamespaceSearchParamsAISearchOptionsQueryRewriteModelGrokGrok4, NamespaceSearchParamsAISearchOptionsQueryRewriteModelGroqLlama3_3_70bVersatile, NamespaceSearchParamsAISearchOptionsQueryRewriteModelGroqLlama3_1_8bInstant, NamespaceSearchParamsAISearchOptionsQueryRewriteModelOpenAIGpt5, NamespaceSearchParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Mini, NamespaceSearchParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Nano, NamespaceSearchParamsAISearchOptionsQueryRewriteModelEmpty:
		return true
	}
	return false
}

type NamespaceSearchParamsAISearchOptionsReranking struct {
	Enabled        param.Field[bool]                                               `json:"enabled"`
	MatchThreshold param.Field[float64]                                            `json:"match_threshold"`
	Model          param.Field[NamespaceSearchParamsAISearchOptionsRerankingModel] `json:"model"`
}

func (r NamespaceSearchParamsAISearchOptionsReranking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceSearchParamsAISearchOptionsRerankingModel string

const (
	NamespaceSearchParamsAISearchOptionsRerankingModelCfBaaiBgeRerankerBase NamespaceSearchParamsAISearchOptionsRerankingModel = "@cf/baai/bge-reranker-base"
	NamespaceSearchParamsAISearchOptionsRerankingModelEmpty                 NamespaceSearchParamsAISearchOptionsRerankingModel = ""
)

func (r NamespaceSearchParamsAISearchOptionsRerankingModel) IsKnown() bool {
	switch r {
	case NamespaceSearchParamsAISearchOptionsRerankingModelCfBaaiBgeRerankerBase, NamespaceSearchParamsAISearchOptionsRerankingModelEmpty:
		return true
	}
	return false
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
	// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
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
	// custom_metadata field. Numeric and datetime fields support asc/desc directions;
	// text/boolean fields support exists/not_exists.
	Field param.Field[string] `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction param.Field[NamespaceSearchParamsAISearchOptionsRetrievalBoostByDirection] `json:"direction"`
}

func (r NamespaceSearchParamsAISearchOptionsRetrievalBoostBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
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
// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
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
	Content     param.Field[string]                            `json:"content" api:"required"`
	Role        param.Field[NamespaceSearchParamsMessagesRole] `json:"role" api:"required"`
	ExtraFields map[string]interface{}                         `json:"-,extras"`
}

func (r NamespaceSearchParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
