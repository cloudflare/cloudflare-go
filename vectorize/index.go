// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vectorize

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// IndexService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIndexService] method instead.
type IndexService struct {
	Options []option.RequestOption
}

// NewIndexService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIndexService(opts ...option.RequestOption) (r *IndexService) {
	r = &IndexService{}
	r.Options = opts
	return
}

// Creates and returns a new Vectorize Index.
func (r *IndexService) New(ctx context.Context, params IndexNewParams, opts ...option.RequestOption) (res *CreateIndex, err error) {
	var env IndexNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates and returns the specified Vectorize Index.
func (r *IndexService) Update(ctx context.Context, indexName string, params IndexUpdateParams, opts ...option.RequestOption) (res *CreateIndex, err error) {
	var env IndexUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s", params.AccountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns a list of Vectorize Indexes
func (r *IndexService) List(ctx context.Context, query IndexListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CreateIndex], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes", query.AccountID)
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

// Returns a list of Vectorize Indexes
func (r *IndexService) ListAutoPaging(ctx context.Context, query IndexListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CreateIndex] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes the specified Vectorize Index.
func (r *IndexService) Delete(ctx context.Context, indexName string, body IndexDeleteParams, opts ...option.RequestOption) (res *IndexDeleteResponseUnion, err error) {
	var env IndexDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s", body.AccountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a set of vectors from an index by their vector identifiers.
func (r *IndexService) DeleteByIDs(ctx context.Context, indexName string, params IndexDeleteByIDsParams, opts ...option.RequestOption) (res *IndexDeleteVectorsByID, err error) {
	var env IndexDeleteByIDsResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/delete-by-ids", params.AccountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the specified Vectorize Index.
func (r *IndexService) Get(ctx context.Context, indexName string, query IndexGetParams, opts ...option.RequestOption) (res *CreateIndex, err error) {
	var env IndexGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s", query.AccountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a set of vectors from an index by their vector identifiers.
func (r *IndexService) GetByIDs(ctx context.Context, indexName string, params IndexGetByIDsParams, opts ...option.RequestOption) (res *IndexGetByIDsResponse, err error) {
	var env IndexGetByIDsResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/get-by-ids", params.AccountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Inserts vectors into the specified index and returns the count of the vectors
// successfully inserted.
func (r *IndexService) Insert(ctx context.Context, indexName string, params IndexInsertParams, opts ...option.RequestOption) (res *IndexInsert, err error) {
	var env IndexInsertResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/insert", params.AccountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Finds vectors closest to a given vector in an index.
func (r *IndexService) Query(ctx context.Context, indexName string, params IndexQueryParams, opts ...option.RequestOption) (res *IndexQuery, err error) {
	var env IndexQueryResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/query", params.AccountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Upserts vectors into the specified index, creating them if they do not exist and
// returns the count of values and ids successfully inserted.
func (r *IndexService) Upsert(ctx context.Context, indexName string, params IndexUpsertParams, opts ...option.RequestOption) (res *IndexUpsert, err error) {
	var env IndexUpsertResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/upsert", params.AccountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CreateIndex struct {
	Config IndexDimensionConfiguration `json:"config"`
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedOn string `json:"created_on"`
	// Specifies the description of the index.
	Description string `json:"description"`
	// Specifies the timestamp the resource was modified as an ISO8601 string.
	ModifiedOn string          `json:"modified_on"`
	Name       string          `json:"name"`
	JSON       createIndexJSON `json:"-"`
}

// createIndexJSON contains the JSON metadata for the struct [CreateIndex]
type createIndexJSON struct {
	Config      apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateIndex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r createIndexJSON) RawJSON() string {
	return r.raw
}

type IndexDeleteVectorsByID struct {
	// The count of the vectors successfully deleted.
	Count int64 `json:"count"`
	// Array of vector identifiers of the vectors that were successfully processed for
	// deletion.
	IDs  []string                   `json:"ids"`
	JSON indexDeleteVectorsByIDJSON `json:"-"`
}

// indexDeleteVectorsByIDJSON contains the JSON metadata for the struct
// [IndexDeleteVectorsByID]
type indexDeleteVectorsByIDJSON struct {
	Count       apijson.Field
	IDs         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexDeleteVectorsByID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexDeleteVectorsByIDJSON) RawJSON() string {
	return r.raw
}

type IndexDimensionConfiguration struct {
	// Specifies the number of dimensions for the index
	Dimensions int64 `json:"dimensions,required"`
	// Specifies the type of metric to use calculating distance.
	Metric IndexDimensionConfigurationMetric `json:"metric,required"`
	JSON   indexDimensionConfigurationJSON   `json:"-"`
}

// indexDimensionConfigurationJSON contains the JSON metadata for the struct
// [IndexDimensionConfiguration]
type indexDimensionConfigurationJSON struct {
	Dimensions  apijson.Field
	Metric      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexDimensionConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexDimensionConfigurationJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of metric to use calculating distance.
type IndexDimensionConfigurationMetric string

const (
	IndexDimensionConfigurationMetricCosine     IndexDimensionConfigurationMetric = "cosine"
	IndexDimensionConfigurationMetricEuclidean  IndexDimensionConfigurationMetric = "euclidean"
	IndexDimensionConfigurationMetricDOTProduct IndexDimensionConfigurationMetric = "dot-product"
)

func (r IndexDimensionConfigurationMetric) IsKnown() bool {
	switch r {
	case IndexDimensionConfigurationMetricCosine, IndexDimensionConfigurationMetricEuclidean, IndexDimensionConfigurationMetricDOTProduct:
		return true
	}
	return false
}

type IndexDimensionConfigurationParam struct {
	// Specifies the number of dimensions for the index
	Dimensions param.Field[int64] `json:"dimensions,required"`
	// Specifies the type of metric to use calculating distance.
	Metric param.Field[IndexDimensionConfigurationMetric] `json:"metric,required"`
}

func (r IndexDimensionConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IndexDimensionConfigurationParam) implementsVectorizeIndexNewParamsConfigUnion() {}

type IndexInsert struct {
	// Specifies the count of the vectors successfully inserted.
	Count int64 `json:"count"`
	// Array of vector identifiers of the vectors successfully inserted.
	IDs  []string        `json:"ids"`
	JSON indexInsertJSON `json:"-"`
}

// indexInsertJSON contains the JSON metadata for the struct [IndexInsert]
type indexInsertJSON struct {
	Count       apijson.Field
	IDs         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexInsert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexInsertJSON) RawJSON() string {
	return r.raw
}

type IndexQuery struct {
	// Specifies the count of vectors returned by the search
	Count int64 `json:"count"`
	// Array of vectors matched by the search
	Matches []IndexQueryMatch `json:"matches"`
	JSON    indexQueryJSON    `json:"-"`
}

// indexQueryJSON contains the JSON metadata for the struct [IndexQuery]
type indexQueryJSON struct {
	Count       apijson.Field
	Matches     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexQueryJSON) RawJSON() string {
	return r.raw
}

type IndexQueryMatch struct {
	// Identifier
	ID       string      `json:"id"`
	Metadata interface{} `json:"metadata,nullable"`
	// The score of the vector according to the index's distance metric
	Score  float64             `json:"score"`
	Values []float64           `json:"values,nullable"`
	JSON   indexQueryMatchJSON `json:"-"`
}

// indexQueryMatchJSON contains the JSON metadata for the struct [IndexQueryMatch]
type indexQueryMatchJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Score       apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexQueryMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexQueryMatchJSON) RawJSON() string {
	return r.raw
}

type IndexUpsert struct {
	// Specifies the count of the vectors successfully inserted.
	Count int64 `json:"count"`
	// Array of vector identifiers of the vectors successfully inserted.
	IDs  []string        `json:"ids"`
	JSON indexUpsertJSON `json:"-"`
}

// indexUpsertJSON contains the JSON metadata for the struct [IndexUpsert]
type indexUpsertJSON struct {
	Count       apijson.Field
	IDs         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexUpsert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexUpsertJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [vectorize.IndexDeleteResponseUnknown] or
// [shared.UnionString].
type IndexDeleteResponseUnion interface {
	ImplementsVectorizeIndexDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IndexDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type IndexGetByIDsResponse = interface{}

type IndexNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Specifies the type of configuration to use for the index.
	Config param.Field[IndexNewParamsConfigUnion] `json:"config,required"`
	Name   param.Field[string]                    `json:"name,required"`
	// Specifies the description of the index.
	Description param.Field[string] `json:"description"`
}

func (r IndexNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of configuration to use for the index.
type IndexNewParamsConfig struct {
	// Specifies the number of dimensions for the index
	Dimensions param.Field[int64] `json:"dimensions"`
	// Specifies the type of metric to use calculating distance.
	Metric param.Field[IndexNewParamsConfigMetric] `json:"metric"`
	// Specifies the preset to use for the index.
	Preset param.Field[IndexNewParamsConfigPreset] `json:"preset"`
}

func (r IndexNewParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IndexNewParamsConfig) implementsVectorizeIndexNewParamsConfigUnion() {}

// Specifies the type of configuration to use for the index.
//
// Satisfied by [vectorize.IndexDimensionConfigurationParam],
// [vectorize.IndexNewParamsConfigVectorizeIndexPresetConfiguration],
// [IndexNewParamsConfig].
type IndexNewParamsConfigUnion interface {
	implementsVectorizeIndexNewParamsConfigUnion()
}

type IndexNewParamsConfigVectorizeIndexPresetConfiguration struct {
	// Specifies the preset to use for the index.
	Preset param.Field[IndexNewParamsConfigVectorizeIndexPresetConfigurationPreset] `json:"preset,required"`
}

func (r IndexNewParamsConfigVectorizeIndexPresetConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IndexNewParamsConfigVectorizeIndexPresetConfiguration) implementsVectorizeIndexNewParamsConfigUnion() {
}

// Specifies the preset to use for the index.
type IndexNewParamsConfigVectorizeIndexPresetConfigurationPreset string

const (
	IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeSmallEnV1_5        IndexNewParamsConfigVectorizeIndexPresetConfigurationPreset = "@cf/baai/bge-small-en-v1.5"
	IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeBaseEnV1_5         IndexNewParamsConfigVectorizeIndexPresetConfigurationPreset = "@cf/baai/bge-base-en-v1.5"
	IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeLargeEnV1_5        IndexNewParamsConfigVectorizeIndexPresetConfigurationPreset = "@cf/baai/bge-large-en-v1.5"
	IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetOpenAITextEmbeddingAda002   IndexNewParamsConfigVectorizeIndexPresetConfigurationPreset = "openai/text-embedding-ada-002"
	IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCohereEmbedMultilingualV2_0 IndexNewParamsConfigVectorizeIndexPresetConfigurationPreset = "cohere/embed-multilingual-v2.0"
)

func (r IndexNewParamsConfigVectorizeIndexPresetConfigurationPreset) IsKnown() bool {
	switch r {
	case IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeSmallEnV1_5, IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeBaseEnV1_5, IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeLargeEnV1_5, IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetOpenAITextEmbeddingAda002, IndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCohereEmbedMultilingualV2_0:
		return true
	}
	return false
}

// Specifies the type of metric to use calculating distance.
type IndexNewParamsConfigMetric string

const (
	IndexNewParamsConfigMetricCosine     IndexNewParamsConfigMetric = "cosine"
	IndexNewParamsConfigMetricEuclidean  IndexNewParamsConfigMetric = "euclidean"
	IndexNewParamsConfigMetricDOTProduct IndexNewParamsConfigMetric = "dot-product"
)

func (r IndexNewParamsConfigMetric) IsKnown() bool {
	switch r {
	case IndexNewParamsConfigMetricCosine, IndexNewParamsConfigMetricEuclidean, IndexNewParamsConfigMetricDOTProduct:
		return true
	}
	return false
}

// Specifies the preset to use for the index.
type IndexNewParamsConfigPreset string

const (
	IndexNewParamsConfigPresetCfBaaiBgeSmallEnV1_5        IndexNewParamsConfigPreset = "@cf/baai/bge-small-en-v1.5"
	IndexNewParamsConfigPresetCfBaaiBgeBaseEnV1_5         IndexNewParamsConfigPreset = "@cf/baai/bge-base-en-v1.5"
	IndexNewParamsConfigPresetCfBaaiBgeLargeEnV1_5        IndexNewParamsConfigPreset = "@cf/baai/bge-large-en-v1.5"
	IndexNewParamsConfigPresetOpenAITextEmbeddingAda002   IndexNewParamsConfigPreset = "openai/text-embedding-ada-002"
	IndexNewParamsConfigPresetCohereEmbedMultilingualV2_0 IndexNewParamsConfigPreset = "cohere/embed-multilingual-v2.0"
)

func (r IndexNewParamsConfigPreset) IsKnown() bool {
	switch r {
	case IndexNewParamsConfigPresetCfBaaiBgeSmallEnV1_5, IndexNewParamsConfigPresetCfBaaiBgeBaseEnV1_5, IndexNewParamsConfigPresetCfBaaiBgeLargeEnV1_5, IndexNewParamsConfigPresetOpenAITextEmbeddingAda002, IndexNewParamsConfigPresetCohereEmbedMultilingualV2_0:
		return true
	}
	return false
}

type IndexNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   CreateIndex           `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexNewResponseEnvelopeJSON    `json:"-"`
}

// indexNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndexNewResponseEnvelope]
type indexNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexNewResponseEnvelopeSuccess bool

const (
	IndexNewResponseEnvelopeSuccessTrue IndexNewResponseEnvelopeSuccess = true
)

func (r IndexNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndexUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Specifies the description of the index.
	Description param.Field[string] `json:"description,required"`
}

func (r IndexUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndexUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   CreateIndex           `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexUpdateResponseEnvelopeJSON    `json:"-"`
}

// indexUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndexUpdateResponseEnvelope]
type indexUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexUpdateResponseEnvelopeSuccess bool

const (
	IndexUpdateResponseEnvelopeSuccessTrue IndexUpdateResponseEnvelopeSuccess = true
)

func (r IndexUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndexListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndexDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndexDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   IndexDeleteResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success IndexDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexDeleteResponseEnvelopeJSON    `json:"-"`
}

// indexDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndexDeleteResponseEnvelope]
type indexDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexDeleteResponseEnvelopeSuccess bool

const (
	IndexDeleteResponseEnvelopeSuccessTrue IndexDeleteResponseEnvelopeSuccess = true
)

func (r IndexDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndexDeleteByIDsParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A list of vector identifiers to delete from the index indicated by the path.
	IDs param.Field[[]string] `json:"ids"`
}

func (r IndexDeleteByIDsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndexDeleteByIDsResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   IndexDeleteVectorsByID `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexDeleteByIDsResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexDeleteByIDsResponseEnvelopeJSON    `json:"-"`
}

// indexDeleteByIDsResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndexDeleteByIDsResponseEnvelope]
type indexDeleteByIDsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexDeleteByIDsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexDeleteByIDsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexDeleteByIDsResponseEnvelopeSuccess bool

const (
	IndexDeleteByIDsResponseEnvelopeSuccessTrue IndexDeleteByIDsResponseEnvelopeSuccess = true
)

func (r IndexDeleteByIDsResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexDeleteByIDsResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndexGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndexGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   CreateIndex           `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexGetResponseEnvelopeJSON    `json:"-"`
}

// indexGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndexGetResponseEnvelope]
type indexGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexGetResponseEnvelopeSuccess bool

const (
	IndexGetResponseEnvelopeSuccessTrue IndexGetResponseEnvelopeSuccess = true
)

func (r IndexGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndexGetByIDsParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A list of vector identifiers to retrieve from the index indicated by the path.
	IDs param.Field[[]string] `json:"ids"`
}

func (r IndexGetByIDsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndexGetByIDsResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Array of vectors with matching ids.
	Result IndexGetByIDsResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexGetByIDsResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexGetByIDsResponseEnvelopeJSON    `json:"-"`
}

// indexGetByIDsResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndexGetByIDsResponseEnvelope]
type indexGetByIDsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexGetByIDsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexGetByIDsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexGetByIDsResponseEnvelopeSuccess bool

const (
	IndexGetByIDsResponseEnvelopeSuccessTrue IndexGetByIDsResponseEnvelopeSuccess = true
)

func (r IndexGetByIDsResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexGetByIDsResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndexInsertParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// ndjson file containing vectors to insert.
	Body io.Reader `json:"body,required" format:"binary"`
}

func (r IndexInsertParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type IndexInsertResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   IndexInsert           `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexInsertResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexInsertResponseEnvelopeJSON    `json:"-"`
}

// indexInsertResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndexInsertResponseEnvelope]
type indexInsertResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexInsertResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexInsertResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexInsertResponseEnvelopeSuccess bool

const (
	IndexInsertResponseEnvelopeSuccessTrue IndexInsertResponseEnvelopeSuccess = true
)

func (r IndexInsertResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexInsertResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndexQueryParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The search vector that will be used to find the nearest neighbors.
	Vector param.Field[[]float64] `json:"vector,required"`
	// A metadata filter expression used to limit nearest neighbor results.
	Filter param.Field[interface{}] `json:"filter"`
	// Whether to return the metadata associated with the closest vectors.
	ReturnMetadata param.Field[bool] `json:"returnMetadata"`
	// Whether to return the values associated with the closest vectors.
	ReturnValues param.Field[bool] `json:"returnValues"`
	// The number of nearest neighbors to find.
	TopK param.Field[float64] `json:"topK"`
}

func (r IndexQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndexQueryResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   IndexQuery            `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexQueryResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexQueryResponseEnvelopeJSON    `json:"-"`
}

// indexQueryResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndexQueryResponseEnvelope]
type indexQueryResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexQueryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexQueryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexQueryResponseEnvelopeSuccess bool

const (
	IndexQueryResponseEnvelopeSuccessTrue IndexQueryResponseEnvelopeSuccess = true
)

func (r IndexQueryResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexQueryResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IndexUpsertParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// ndjson file containing vectors to upsert.
	Body io.Reader `json:"body,required" format:"binary"`
}

func (r IndexUpsertParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type IndexUpsertResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   IndexUpsert           `json:"result,required,nullable"`
	// Whether the API call was successful
	Success IndexUpsertResponseEnvelopeSuccess `json:"success,required"`
	JSON    indexUpsertResponseEnvelopeJSON    `json:"-"`
}

// indexUpsertResponseEnvelopeJSON contains the JSON metadata for the struct
// [IndexUpsertResponseEnvelope]
type indexUpsertResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexUpsertResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexUpsertResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndexUpsertResponseEnvelopeSuccess bool

const (
	IndexUpsertResponseEnvelopeSuccessTrue IndexUpsertResponseEnvelopeSuccess = true
)

func (r IndexUpsertResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndexUpsertResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
