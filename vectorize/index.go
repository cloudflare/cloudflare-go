// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vectorize

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// IndexService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewIndexService] method instead.
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
func (r *IndexService) New(ctx context.Context, accountIdentifier string, body IndexNewParams, opts ...option.RequestOption) (res *VectorizeCreateIndex, err error) {
	opts = append(r.Options[:], opts...)
	var env IndexNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/vectorize/indexes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates and returns the specified Vectorize Index.
func (r *IndexService) Update(ctx context.Context, accountIdentifier string, indexName string, body IndexUpdateParams, opts ...option.RequestOption) (res *VectorizeCreateIndex, err error) {
	opts = append(r.Options[:], opts...)
	var env IndexUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns a list of Vectorize Indexes
func (r *IndexService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *pagination.SinglePage[VectorizeCreateIndex], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/vectorize/indexes", accountIdentifier)
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
func (r *IndexService) ListAutoPaging(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) *pagination.SinglePageAutoPager[VectorizeCreateIndex] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, accountIdentifier, opts...))
}

// Deletes the specified Vectorize Index.
func (r *IndexService) Delete(ctx context.Context, accountIdentifier string, indexName string, opts ...option.RequestOption) (res *IndexDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IndexDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a set of vectors from an index by their vector identifiers.
func (r *IndexService) DeleteByIDs(ctx context.Context, accountIdentifier string, indexName string, body IndexDeleteByIDsParams, opts ...option.RequestOption) (res *VectorizeIndexDeleteVectorsByID, err error) {
	opts = append(r.Options[:], opts...)
	var env IndexDeleteByIDsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/delete-by-ids", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the specified Vectorize Index.
func (r *IndexService) Get(ctx context.Context, accountIdentifier string, indexName string, opts ...option.RequestOption) (res *VectorizeCreateIndex, err error) {
	opts = append(r.Options[:], opts...)
	var env IndexGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a set of vectors from an index by their vector identifiers.
func (r *IndexService) GetByIDs(ctx context.Context, accountIdentifier string, indexName string, body IndexGetByIDsParams, opts ...option.RequestOption) (res *IndexGetByIDsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IndexGetByIDsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/get-by-ids", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Inserts vectors into the specified index and returns the count of the vectors
// successfully inserted.
func (r *IndexService) Insert(ctx context.Context, accountIdentifier string, indexName string, body IndexInsertParams, opts ...option.RequestOption) (res *VectorizeIndexInsert, err error) {
	opts = append(r.Options[:], opts...)
	var env IndexInsertResponseEnvelope
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/insert", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Finds vectors closest to a given vector in an index.
func (r *IndexService) Query(ctx context.Context, accountIdentifier string, indexName string, body IndexQueryParams, opts ...option.RequestOption) (res *VectorizeIndexQuery, err error) {
	opts = append(r.Options[:], opts...)
	var env IndexQueryResponseEnvelope
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/query", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Upserts vectors into the specified index, creating them if they do not exist and
// returns the count of values and ids successfully inserted.
func (r *IndexService) Upsert(ctx context.Context, accountIdentifier string, indexName string, body IndexUpsertParams, opts ...option.RequestOption) (res *VectorizeIndexUpsert, err error) {
	opts = append(r.Options[:], opts...)
	var env IndexUpsertResponseEnvelope
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/upsert", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type VectorizeCreateIndex struct {
	Config VectorizeCreateIndexConfig `json:"config"`
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedOn string `json:"created_on"`
	// Specifies the description of the index.
	Description string `json:"description"`
	// Specifies the timestamp the resource was modified as an ISO8601 string.
	ModifiedOn string                   `json:"modified_on"`
	Name       string                   `json:"name"`
	JSON       vectorizeCreateIndexJSON `json:"-"`
}

// vectorizeCreateIndexJSON contains the JSON metadata for the struct
// [VectorizeCreateIndex]
type vectorizeCreateIndexJSON struct {
	Config      apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeCreateIndex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vectorizeCreateIndexJSON) RawJSON() string {
	return r.raw
}

type VectorizeCreateIndexConfig struct {
	// Specifies the number of dimensions for the index
	Dimensions int64 `json:"dimensions,required"`
	// Specifies the type of metric to use calculating distance.
	Metric VectorizeCreateIndexConfigMetric `json:"metric,required"`
	JSON   vectorizeCreateIndexConfigJSON   `json:"-"`
}

// vectorizeCreateIndexConfigJSON contains the JSON metadata for the struct
// [VectorizeCreateIndexConfig]
type vectorizeCreateIndexConfigJSON struct {
	Dimensions  apijson.Field
	Metric      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeCreateIndexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vectorizeCreateIndexConfigJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of metric to use calculating distance.
type VectorizeCreateIndexConfigMetric string

const (
	VectorizeCreateIndexConfigMetricCosine     VectorizeCreateIndexConfigMetric = "cosine"
	VectorizeCreateIndexConfigMetricEuclidean  VectorizeCreateIndexConfigMetric = "euclidean"
	VectorizeCreateIndexConfigMetricDotProduct VectorizeCreateIndexConfigMetric = "dot-product"
)

func (r VectorizeCreateIndexConfigMetric) IsKnown() bool {
	switch r {
	case VectorizeCreateIndexConfigMetricCosine, VectorizeCreateIndexConfigMetricEuclidean, VectorizeCreateIndexConfigMetricDotProduct:
		return true
	}
	return false
}

type VectorizeIndexDeleteVectorsByID struct {
	// The count of the vectors successfully deleted.
	Count int64 `json:"count"`
	// Array of vector identifiers of the vectors that were successfully processed for
	// deletion.
	IDs  []string                            `json:"ids"`
	JSON vectorizeIndexDeleteVectorsByIDJSON `json:"-"`
}

// vectorizeIndexDeleteVectorsByIDJSON contains the JSON metadata for the struct
// [VectorizeIndexDeleteVectorsByID]
type vectorizeIndexDeleteVectorsByIDJSON struct {
	Count       apijson.Field
	IDs         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexDeleteVectorsByID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vectorizeIndexDeleteVectorsByIDJSON) RawJSON() string {
	return r.raw
}

type VectorizeIndexInsert struct {
	// Specifies the count of the vectors successfully inserted.
	Count int64 `json:"count"`
	// Array of vector identifiers of the vectors successfully inserted.
	IDs  []string                 `json:"ids"`
	JSON vectorizeIndexInsertJSON `json:"-"`
}

// vectorizeIndexInsertJSON contains the JSON metadata for the struct
// [VectorizeIndexInsert]
type vectorizeIndexInsertJSON struct {
	Count       apijson.Field
	IDs         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexInsert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vectorizeIndexInsertJSON) RawJSON() string {
	return r.raw
}

type VectorizeIndexQuery struct {
	// Specifies the count of vectors returned by the search
	Count int64 `json:"count"`
	// Array of vectors matched by the search
	Matches []VectorizeIndexQueryMatch `json:"matches"`
	JSON    vectorizeIndexQueryJSON    `json:"-"`
}

// vectorizeIndexQueryJSON contains the JSON metadata for the struct
// [VectorizeIndexQuery]
type vectorizeIndexQueryJSON struct {
	Count       apijson.Field
	Matches     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vectorizeIndexQueryJSON) RawJSON() string {
	return r.raw
}

type VectorizeIndexQueryMatch struct {
	// Identifier
	ID       string      `json:"id"`
	Metadata interface{} `json:"metadata"`
	// The score of the vector according to the index's distance metric
	Score  float64                      `json:"score"`
	Values []float64                    `json:"values"`
	JSON   vectorizeIndexQueryMatchJSON `json:"-"`
}

// vectorizeIndexQueryMatchJSON contains the JSON metadata for the struct
// [VectorizeIndexQueryMatch]
type vectorizeIndexQueryMatchJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Score       apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexQueryMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vectorizeIndexQueryMatchJSON) RawJSON() string {
	return r.raw
}

type VectorizeIndexUpsert struct {
	// Specifies the count of the vectors successfully inserted.
	Count int64 `json:"count"`
	// Array of vector identifiers of the vectors successfully inserted.
	IDs  []string                 `json:"ids"`
	JSON vectorizeIndexUpsertJSON `json:"-"`
}

// vectorizeIndexUpsertJSON contains the JSON metadata for the struct
// [VectorizeIndexUpsert]
type vectorizeIndexUpsertJSON struct {
	Count       apijson.Field
	IDs         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexUpsert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vectorizeIndexUpsertJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [vectorize.IndexDeleteResponseUnknown] or
// [shared.UnionString].
type IndexDeleteResponse interface {
	ImplementsVectorizeIndexDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IndexDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type IndexGetByIDsResponse = interface{}

type IndexNewParams struct {
	// Specifies the type of configuration to use for the index.
	Config param.Field[IndexNewParamsConfig] `json:"config,required"`
	Name   param.Field[string]               `json:"name,required"`
	// Specifies the description of the index.
	Description param.Field[string] `json:"description"`
}

func (r IndexNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of configuration to use for the index.
//
// Satisfied by [vectorize.IndexNewParamsConfigVectorizeIndexPresetConfiguration],
// [vectorize.IndexNewParamsConfigVectorizeIndexDimensionConfiguration].
type IndexNewParamsConfig interface {
	implementsVectorizeIndexNewParamsConfig()
}

type IndexNewParamsConfigVectorizeIndexPresetConfiguration struct {
	// Specifies the preset to use for the index.
	Preset param.Field[IndexNewParamsConfigVectorizeIndexPresetConfigurationPreset] `json:"preset,required"`
}

func (r IndexNewParamsConfigVectorizeIndexPresetConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IndexNewParamsConfigVectorizeIndexPresetConfiguration) implementsVectorizeIndexNewParamsConfig() {
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

type IndexNewParamsConfigVectorizeIndexDimensionConfiguration struct {
	// Specifies the number of dimensions for the index
	Dimensions param.Field[int64] `json:"dimensions,required"`
	// Specifies the type of metric to use calculating distance.
	Metric param.Field[IndexNewParamsConfigVectorizeIndexDimensionConfigurationMetric] `json:"metric,required"`
}

func (r IndexNewParamsConfigVectorizeIndexDimensionConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IndexNewParamsConfigVectorizeIndexDimensionConfiguration) implementsVectorizeIndexNewParamsConfig() {
}

// Specifies the type of metric to use calculating distance.
type IndexNewParamsConfigVectorizeIndexDimensionConfigurationMetric string

const (
	IndexNewParamsConfigVectorizeIndexDimensionConfigurationMetricCosine     IndexNewParamsConfigVectorizeIndexDimensionConfigurationMetric = "cosine"
	IndexNewParamsConfigVectorizeIndexDimensionConfigurationMetricEuclidean  IndexNewParamsConfigVectorizeIndexDimensionConfigurationMetric = "euclidean"
	IndexNewParamsConfigVectorizeIndexDimensionConfigurationMetricDotProduct IndexNewParamsConfigVectorizeIndexDimensionConfigurationMetric = "dot-product"
)

func (r IndexNewParamsConfigVectorizeIndexDimensionConfigurationMetric) IsKnown() bool {
	switch r {
	case IndexNewParamsConfigVectorizeIndexDimensionConfigurationMetricCosine, IndexNewParamsConfigVectorizeIndexDimensionConfigurationMetricEuclidean, IndexNewParamsConfigVectorizeIndexDimensionConfigurationMetricDotProduct:
		return true
	}
	return false
}

type IndexNewResponseEnvelope struct {
	Errors   []IndexNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndexNewResponseEnvelopeMessages `json:"messages,required"`
	Result   VectorizeCreateIndex               `json:"result,required,nullable"`
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

type IndexNewResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    indexNewResponseEnvelopeErrorsJSON `json:"-"`
}

// indexNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [IndexNewResponseEnvelopeErrors]
type indexNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndexNewResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    indexNewResponseEnvelopeMessagesJSON `json:"-"`
}

// indexNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [IndexNewResponseEnvelopeMessages]
type indexNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexNewResponseEnvelopeMessagesJSON) RawJSON() string {
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
	// Specifies the description of the index.
	Description param.Field[string] `json:"description,required"`
}

func (r IndexUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndexUpdateResponseEnvelope struct {
	Errors   []IndexUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndexUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   VectorizeCreateIndex                  `json:"result,required,nullable"`
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

type IndexUpdateResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    indexUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// indexUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [IndexUpdateResponseEnvelopeErrors]
type indexUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndexUpdateResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    indexUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// indexUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IndexUpdateResponseEnvelopeMessages]
type indexUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
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

type IndexDeleteResponseEnvelope struct {
	Errors   []IndexDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndexDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   IndexDeleteResponse                   `json:"result,required,nullable"`
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

type IndexDeleteResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    indexDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// indexDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [IndexDeleteResponseEnvelopeErrors]
type indexDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndexDeleteResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    indexDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// indexDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IndexDeleteResponseEnvelopeMessages]
type indexDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
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
	// A list of vector identifiers to delete from the index indicated by the path.
	IDs param.Field[[]string] `json:"ids"`
}

func (r IndexDeleteByIDsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndexDeleteByIDsResponseEnvelope struct {
	Errors   []IndexDeleteByIDsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndexDeleteByIDsResponseEnvelopeMessages `json:"messages,required"`
	Result   VectorizeIndexDeleteVectorsByID            `json:"result,required,nullable"`
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

type IndexDeleteByIDsResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    indexDeleteByIDsResponseEnvelopeErrorsJSON `json:"-"`
}

// indexDeleteByIDsResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IndexDeleteByIDsResponseEnvelopeErrors]
type indexDeleteByIDsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexDeleteByIDsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexDeleteByIDsResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndexDeleteByIDsResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    indexDeleteByIDsResponseEnvelopeMessagesJSON `json:"-"`
}

// indexDeleteByIDsResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IndexDeleteByIDsResponseEnvelopeMessages]
type indexDeleteByIDsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexDeleteByIDsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexDeleteByIDsResponseEnvelopeMessagesJSON) RawJSON() string {
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

type IndexGetResponseEnvelope struct {
	Errors   []IndexGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndexGetResponseEnvelopeMessages `json:"messages,required"`
	Result   VectorizeCreateIndex               `json:"result,required,nullable"`
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

type IndexGetResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    indexGetResponseEnvelopeErrorsJSON `json:"-"`
}

// indexGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [IndexGetResponseEnvelopeErrors]
type indexGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndexGetResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    indexGetResponseEnvelopeMessagesJSON `json:"-"`
}

// indexGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [IndexGetResponseEnvelopeMessages]
type indexGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexGetResponseEnvelopeMessagesJSON) RawJSON() string {
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
	// A list of vector identifiers to retrieve from the index indicated by the path.
	IDs param.Field[[]string] `json:"ids"`
}

func (r IndexGetByIDsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndexGetByIDsResponseEnvelope struct {
	Errors   []IndexGetByIDsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndexGetByIDsResponseEnvelopeMessages `json:"messages,required"`
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

type IndexGetByIDsResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    indexGetByIDsResponseEnvelopeErrorsJSON `json:"-"`
}

// indexGetByIDsResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IndexGetByIDsResponseEnvelopeErrors]
type indexGetByIDsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexGetByIDsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexGetByIDsResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndexGetByIDsResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    indexGetByIDsResponseEnvelopeMessagesJSON `json:"-"`
}

// indexGetByIDsResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IndexGetByIDsResponseEnvelopeMessages]
type indexGetByIDsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexGetByIDsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexGetByIDsResponseEnvelopeMessagesJSON) RawJSON() string {
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
	Body param.Field[interface{}] `json:"body,required"`
}

func (r IndexInsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type IndexInsertResponseEnvelope struct {
	Errors   []IndexInsertResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndexInsertResponseEnvelopeMessages `json:"messages,required"`
	Result   VectorizeIndexInsert                  `json:"result,required,nullable"`
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

type IndexInsertResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    indexInsertResponseEnvelopeErrorsJSON `json:"-"`
}

// indexInsertResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [IndexInsertResponseEnvelopeErrors]
type indexInsertResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexInsertResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexInsertResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndexInsertResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    indexInsertResponseEnvelopeMessagesJSON `json:"-"`
}

// indexInsertResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IndexInsertResponseEnvelopeMessages]
type indexInsertResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexInsertResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexInsertResponseEnvelopeMessagesJSON) RawJSON() string {
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
	// Whether to return the metadata associated with the closest vectors.
	ReturnMetadata param.Field[bool] `json:"returnMetadata"`
	// Whether to return the values associated with the closest vectors.
	ReturnValues param.Field[bool] `json:"returnValues"`
	// The number of nearest neighbors to find.
	TopK param.Field[float64] `json:"topK"`
	// The search vector that will be used to find the nearest neighbors.
	Vector param.Field[[]float64] `json:"vector"`
}

func (r IndexQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IndexQueryResponseEnvelope struct {
	Errors   []IndexQueryResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndexQueryResponseEnvelopeMessages `json:"messages,required"`
	Result   VectorizeIndexQuery                  `json:"result,required,nullable"`
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

type IndexQueryResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    indexQueryResponseEnvelopeErrorsJSON `json:"-"`
}

// indexQueryResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [IndexQueryResponseEnvelopeErrors]
type indexQueryResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexQueryResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexQueryResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndexQueryResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    indexQueryResponseEnvelopeMessagesJSON `json:"-"`
}

// indexQueryResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [IndexQueryResponseEnvelopeMessages]
type indexQueryResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexQueryResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexQueryResponseEnvelopeMessagesJSON) RawJSON() string {
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
	Body param.Field[interface{}] `json:"body,required"`
}

func (r IndexUpsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type IndexUpsertResponseEnvelope struct {
	Errors   []IndexUpsertResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IndexUpsertResponseEnvelopeMessages `json:"messages,required"`
	Result   VectorizeIndexUpsert                  `json:"result,required,nullable"`
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

type IndexUpsertResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    indexUpsertResponseEnvelopeErrorsJSON `json:"-"`
}

// indexUpsertResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [IndexUpsertResponseEnvelopeErrors]
type indexUpsertResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexUpsertResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexUpsertResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IndexUpsertResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    indexUpsertResponseEnvelopeMessagesJSON `json:"-"`
}

// indexUpsertResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IndexUpsertResponseEnvelopeMessages]
type indexUpsertResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndexUpsertResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indexUpsertResponseEnvelopeMessagesJSON) RawJSON() string {
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
