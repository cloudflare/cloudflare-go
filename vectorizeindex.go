// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// VectorizeIndexService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewVectorizeIndexService] method
// instead.
type VectorizeIndexService struct {
	Options []option.RequestOption
}

// NewVectorizeIndexService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVectorizeIndexService(opts ...option.RequestOption) (r *VectorizeIndexService) {
	r = &VectorizeIndexService{}
	r.Options = opts
	return
}

// Creates and returns a new Vectorize Index.
func (r *VectorizeIndexService) New(ctx context.Context, accountIdentifier string, body VectorizeIndexNewParams, opts ...option.RequestOption) (res *VectorizeCreateIndex, err error) {
	opts = append(r.Options[:], opts...)
	var env VectorizeIndexNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/vectorize/indexes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates and returns the specified Vectorize Index.
func (r *VectorizeIndexService) Update(ctx context.Context, accountIdentifier string, indexName string, body VectorizeIndexUpdateParams, opts ...option.RequestOption) (res *VectorizeCreateIndex, err error) {
	opts = append(r.Options[:], opts...)
	var env VectorizeIndexUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns a list of Vectorize Indexes
func (r *VectorizeIndexService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]VectorizeCreateIndex, err error) {
	opts = append(r.Options[:], opts...)
	var env VectorizeIndexListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/vectorize/indexes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes the specified Vectorize Index.
func (r *VectorizeIndexService) Delete(ctx context.Context, accountIdentifier string, indexName string, opts ...option.RequestOption) (res *VectorizeIndexDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env VectorizeIndexDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a set of vectors from an index by their vector identifiers.
func (r *VectorizeIndexService) DeleteByIDs(ctx context.Context, accountIdentifier string, indexName string, body VectorizeIndexDeleteByIDsParams, opts ...option.RequestOption) (res *VectorizeIndexDeleteVectorsByID, err error) {
	opts = append(r.Options[:], opts...)
	var env VectorizeIndexDeleteByIDsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/delete-by-ids", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the specified Vectorize Index.
func (r *VectorizeIndexService) Get(ctx context.Context, accountIdentifier string, indexName string, opts ...option.RequestOption) (res *VectorizeCreateIndex, err error) {
	opts = append(r.Options[:], opts...)
	var env VectorizeIndexGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a set of vectors from an index by their vector identifiers.
func (r *VectorizeIndexService) GetByIDs(ctx context.Context, accountIdentifier string, indexName string, body VectorizeIndexGetByIDsParams, opts ...option.RequestOption) (res *VectorizeIndexGetByIDsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env VectorizeIndexGetByIDsResponseEnvelope
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
func (r *VectorizeIndexService) Insert(ctx context.Context, accountIdentifier string, indexName string, opts ...option.RequestOption) (res *VectorizeIndexInsert, err error) {
	opts = append(r.Options[:], opts...)
	var env VectorizeIndexInsertResponseEnvelope
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/insert", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Finds vectors closest to a given vector in an index.
func (r *VectorizeIndexService) Query(ctx context.Context, accountIdentifier string, indexName string, body VectorizeIndexQueryParams, opts ...option.RequestOption) (res *VectorizeIndexQuery, err error) {
	opts = append(r.Options[:], opts...)
	var env VectorizeIndexQueryResponseEnvelope
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
func (r *VectorizeIndexService) Upsert(ctx context.Context, accountIdentifier string, indexName string, opts ...option.RequestOption) (res *VectorizeIndexUpsert, err error) {
	opts = append(r.Options[:], opts...)
	var env VectorizeIndexUpsertResponseEnvelope
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/upsert", accountIdentifier, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type VectorizeCreateIndex struct {
	Config VectorizeCreateIndexConfig `json:"config"`
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedOn interface{} `json:"created_on"`
	// Specifies the description of the index.
	Description string `json:"description"`
	// Specifies the timestamp the resource was modified as an ISO8601 string.
	ModifiedOn interface{}              `json:"modified_on"`
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

// Specifies the type of metric to use calculating distance.
type VectorizeCreateIndexConfigMetric string

const (
	VectorizeCreateIndexConfigMetricCosine     VectorizeCreateIndexConfigMetric = "cosine"
	VectorizeCreateIndexConfigMetricEuclidean  VectorizeCreateIndexConfigMetric = "euclidean"
	VectorizeCreateIndexConfigMetricDotProduct VectorizeCreateIndexConfigMetric = "dot-product"
)

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

// Union satisfied by [VectorizeIndexDeleteResponseUnknown] or
// [shared.UnionString].
type VectorizeIndexDeleteResponse interface {
	ImplementsVectorizeIndexDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VectorizeIndexDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type VectorizeIndexGetByIDsResponse = interface{}

type VectorizeIndexNewParams struct {
	// Specifies the type of configuration to use for the index.
	Config param.Field[VectorizeIndexNewParamsConfig] `json:"config,required"`
	Name   param.Field[string]                        `json:"name,required"`
	// Specifies the description of the index.
	Description param.Field[string] `json:"description"`
}

func (r VectorizeIndexNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of configuration to use for the index.
//
// Satisfied by [VectorizeIndexNewParamsConfigVectorizeIndexPresetConfiguration],
// [VectorizeIndexNewParamsConfigVectorizeIndexDimensionConfiguration].
type VectorizeIndexNewParamsConfig interface {
	implementsVectorizeIndexNewParamsConfig()
}

type VectorizeIndexNewParamsConfigVectorizeIndexPresetConfiguration struct {
	// Specifies the preset to use for the index.
	Preset param.Field[VectorizeIndexNewParamsConfigVectorizeIndexPresetConfigurationPreset] `json:"preset,required"`
}

func (r VectorizeIndexNewParamsConfigVectorizeIndexPresetConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VectorizeIndexNewParamsConfigVectorizeIndexPresetConfiguration) implementsVectorizeIndexNewParamsConfig() {
}

// Specifies the preset to use for the index.
type VectorizeIndexNewParamsConfigVectorizeIndexPresetConfigurationPreset string

const (
	VectorizeIndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeSmallEnV1_5        VectorizeIndexNewParamsConfigVectorizeIndexPresetConfigurationPreset = "@cf/baai/bge-small-en-v1.5"
	VectorizeIndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeBaseEnV1_5         VectorizeIndexNewParamsConfigVectorizeIndexPresetConfigurationPreset = "@cf/baai/bge-base-en-v1.5"
	VectorizeIndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeLargeEnV1_5        VectorizeIndexNewParamsConfigVectorizeIndexPresetConfigurationPreset = "@cf/baai/bge-large-en-v1.5"
	VectorizeIndexNewParamsConfigVectorizeIndexPresetConfigurationPresetOpenAITextEmbeddingAda002   VectorizeIndexNewParamsConfigVectorizeIndexPresetConfigurationPreset = "openai/text-embedding-ada-002"
	VectorizeIndexNewParamsConfigVectorizeIndexPresetConfigurationPresetCohereEmbedMultilingualV2_0 VectorizeIndexNewParamsConfigVectorizeIndexPresetConfigurationPreset = "cohere/embed-multilingual-v2.0"
)

type VectorizeIndexNewParamsConfigVectorizeIndexDimensionConfiguration struct {
	// Specifies the number of dimensions for the index
	Dimensions param.Field[int64] `json:"dimensions,required"`
	// Specifies the type of metric to use calculating distance.
	Metric param.Field[VectorizeIndexNewParamsConfigVectorizeIndexDimensionConfigurationMetric] `json:"metric,required"`
}

func (r VectorizeIndexNewParamsConfigVectorizeIndexDimensionConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VectorizeIndexNewParamsConfigVectorizeIndexDimensionConfiguration) implementsVectorizeIndexNewParamsConfig() {
}

// Specifies the type of metric to use calculating distance.
type VectorizeIndexNewParamsConfigVectorizeIndexDimensionConfigurationMetric string

const (
	VectorizeIndexNewParamsConfigVectorizeIndexDimensionConfigurationMetricCosine     VectorizeIndexNewParamsConfigVectorizeIndexDimensionConfigurationMetric = "cosine"
	VectorizeIndexNewParamsConfigVectorizeIndexDimensionConfigurationMetricEuclidean  VectorizeIndexNewParamsConfigVectorizeIndexDimensionConfigurationMetric = "euclidean"
	VectorizeIndexNewParamsConfigVectorizeIndexDimensionConfigurationMetricDotProduct VectorizeIndexNewParamsConfigVectorizeIndexDimensionConfigurationMetric = "dot-product"
)

type VectorizeIndexNewResponseEnvelope struct {
	Errors   []VectorizeIndexNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []VectorizeIndexNewResponseEnvelopeMessages `json:"messages,required"`
	Result   VectorizeCreateIndex                        `json:"result,required,nullable"`
	// Whether the API call was successful
	Success VectorizeIndexNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    vectorizeIndexNewResponseEnvelopeJSON    `json:"-"`
}

// vectorizeIndexNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [VectorizeIndexNewResponseEnvelope]
type vectorizeIndexNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexNewResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    vectorizeIndexNewResponseEnvelopeErrorsJSON `json:"-"`
}

// vectorizeIndexNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [VectorizeIndexNewResponseEnvelopeErrors]
type vectorizeIndexNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexNewResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    vectorizeIndexNewResponseEnvelopeMessagesJSON `json:"-"`
}

// vectorizeIndexNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [VectorizeIndexNewResponseEnvelopeMessages]
type vectorizeIndexNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type VectorizeIndexNewResponseEnvelopeSuccess bool

const (
	VectorizeIndexNewResponseEnvelopeSuccessTrue VectorizeIndexNewResponseEnvelopeSuccess = true
)

type VectorizeIndexUpdateParams struct {
	// Specifies the description of the index.
	Description param.Field[string] `json:"description,required"`
}

func (r VectorizeIndexUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VectorizeIndexUpdateResponseEnvelope struct {
	Errors   []VectorizeIndexUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []VectorizeIndexUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   VectorizeCreateIndex                           `json:"result,required,nullable"`
	// Whether the API call was successful
	Success VectorizeIndexUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    vectorizeIndexUpdateResponseEnvelopeJSON    `json:"-"`
}

// vectorizeIndexUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [VectorizeIndexUpdateResponseEnvelope]
type vectorizeIndexUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexUpdateResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    vectorizeIndexUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// vectorizeIndexUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [VectorizeIndexUpdateResponseEnvelopeErrors]
type vectorizeIndexUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexUpdateResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    vectorizeIndexUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// vectorizeIndexUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [VectorizeIndexUpdateResponseEnvelopeMessages]
type vectorizeIndexUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type VectorizeIndexUpdateResponseEnvelopeSuccess bool

const (
	VectorizeIndexUpdateResponseEnvelopeSuccessTrue VectorizeIndexUpdateResponseEnvelopeSuccess = true
)

type VectorizeIndexListResponseEnvelope struct {
	Errors   []VectorizeIndexListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []VectorizeIndexListResponseEnvelopeMessages `json:"messages,required"`
	Result   []VectorizeCreateIndex                       `json:"result,required"`
	// Whether the API call was successful
	Success VectorizeIndexListResponseEnvelopeSuccess `json:"success,required"`
	JSON    vectorizeIndexListResponseEnvelopeJSON    `json:"-"`
}

// vectorizeIndexListResponseEnvelopeJSON contains the JSON metadata for the struct
// [VectorizeIndexListResponseEnvelope]
type vectorizeIndexListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexListResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    vectorizeIndexListResponseEnvelopeErrorsJSON `json:"-"`
}

// vectorizeIndexListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [VectorizeIndexListResponseEnvelopeErrors]
type vectorizeIndexListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexListResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    vectorizeIndexListResponseEnvelopeMessagesJSON `json:"-"`
}

// vectorizeIndexListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [VectorizeIndexListResponseEnvelopeMessages]
type vectorizeIndexListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type VectorizeIndexListResponseEnvelopeSuccess bool

const (
	VectorizeIndexListResponseEnvelopeSuccessTrue VectorizeIndexListResponseEnvelopeSuccess = true
)

type VectorizeIndexDeleteResponseEnvelope struct {
	Errors   []VectorizeIndexDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []VectorizeIndexDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   VectorizeIndexDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success VectorizeIndexDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    vectorizeIndexDeleteResponseEnvelopeJSON    `json:"-"`
}

// vectorizeIndexDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [VectorizeIndexDeleteResponseEnvelope]
type vectorizeIndexDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexDeleteResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    vectorizeIndexDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// vectorizeIndexDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [VectorizeIndexDeleteResponseEnvelopeErrors]
type vectorizeIndexDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexDeleteResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    vectorizeIndexDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// vectorizeIndexDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [VectorizeIndexDeleteResponseEnvelopeMessages]
type vectorizeIndexDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type VectorizeIndexDeleteResponseEnvelopeSuccess bool

const (
	VectorizeIndexDeleteResponseEnvelopeSuccessTrue VectorizeIndexDeleteResponseEnvelopeSuccess = true
)

type VectorizeIndexDeleteByIDsParams struct {
	// A list of vector identifiers to delete from the index indicated by the path.
	IDs param.Field[[]string] `json:"ids"`
}

func (r VectorizeIndexDeleteByIDsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VectorizeIndexDeleteByIDsResponseEnvelope struct {
	Errors   []VectorizeIndexDeleteByIDsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []VectorizeIndexDeleteByIDsResponseEnvelopeMessages `json:"messages,required"`
	Result   VectorizeIndexDeleteVectorsByID                     `json:"result,required,nullable"`
	// Whether the API call was successful
	Success VectorizeIndexDeleteByIDsResponseEnvelopeSuccess `json:"success,required"`
	JSON    vectorizeIndexDeleteByIDsResponseEnvelopeJSON    `json:"-"`
}

// vectorizeIndexDeleteByIDsResponseEnvelopeJSON contains the JSON metadata for the
// struct [VectorizeIndexDeleteByIDsResponseEnvelope]
type vectorizeIndexDeleteByIDsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexDeleteByIDsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexDeleteByIDsResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    vectorizeIndexDeleteByIDsResponseEnvelopeErrorsJSON `json:"-"`
}

// vectorizeIndexDeleteByIDsResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [VectorizeIndexDeleteByIDsResponseEnvelopeErrors]
type vectorizeIndexDeleteByIDsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexDeleteByIDsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexDeleteByIDsResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    vectorizeIndexDeleteByIDsResponseEnvelopeMessagesJSON `json:"-"`
}

// vectorizeIndexDeleteByIDsResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [VectorizeIndexDeleteByIDsResponseEnvelopeMessages]
type vectorizeIndexDeleteByIDsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexDeleteByIDsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type VectorizeIndexDeleteByIDsResponseEnvelopeSuccess bool

const (
	VectorizeIndexDeleteByIDsResponseEnvelopeSuccessTrue VectorizeIndexDeleteByIDsResponseEnvelopeSuccess = true
)

type VectorizeIndexGetResponseEnvelope struct {
	Errors   []VectorizeIndexGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []VectorizeIndexGetResponseEnvelopeMessages `json:"messages,required"`
	Result   VectorizeCreateIndex                        `json:"result,required,nullable"`
	// Whether the API call was successful
	Success VectorizeIndexGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    vectorizeIndexGetResponseEnvelopeJSON    `json:"-"`
}

// vectorizeIndexGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [VectorizeIndexGetResponseEnvelope]
type vectorizeIndexGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    vectorizeIndexGetResponseEnvelopeErrorsJSON `json:"-"`
}

// vectorizeIndexGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [VectorizeIndexGetResponseEnvelopeErrors]
type vectorizeIndexGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    vectorizeIndexGetResponseEnvelopeMessagesJSON `json:"-"`
}

// vectorizeIndexGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [VectorizeIndexGetResponseEnvelopeMessages]
type vectorizeIndexGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type VectorizeIndexGetResponseEnvelopeSuccess bool

const (
	VectorizeIndexGetResponseEnvelopeSuccessTrue VectorizeIndexGetResponseEnvelopeSuccess = true
)

type VectorizeIndexGetByIDsParams struct {
	// A list of vector identifiers to retrieve from the index indicated by the path.
	IDs param.Field[[]string] `json:"ids"`
}

func (r VectorizeIndexGetByIDsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VectorizeIndexGetByIDsResponseEnvelope struct {
	Errors   []VectorizeIndexGetByIDsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []VectorizeIndexGetByIDsResponseEnvelopeMessages `json:"messages,required"`
	// Array of vectors with matching ids.
	Result VectorizeIndexGetByIDsResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success VectorizeIndexGetByIDsResponseEnvelopeSuccess `json:"success,required"`
	JSON    vectorizeIndexGetByIDsResponseEnvelopeJSON    `json:"-"`
}

// vectorizeIndexGetByIDsResponseEnvelopeJSON contains the JSON metadata for the
// struct [VectorizeIndexGetByIDsResponseEnvelope]
type vectorizeIndexGetByIDsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexGetByIDsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexGetByIDsResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    vectorizeIndexGetByIDsResponseEnvelopeErrorsJSON `json:"-"`
}

// vectorizeIndexGetByIDsResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [VectorizeIndexGetByIDsResponseEnvelopeErrors]
type vectorizeIndexGetByIDsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexGetByIDsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexGetByIDsResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    vectorizeIndexGetByIDsResponseEnvelopeMessagesJSON `json:"-"`
}

// vectorizeIndexGetByIDsResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [VectorizeIndexGetByIDsResponseEnvelopeMessages]
type vectorizeIndexGetByIDsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexGetByIDsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type VectorizeIndexGetByIDsResponseEnvelopeSuccess bool

const (
	VectorizeIndexGetByIDsResponseEnvelopeSuccessTrue VectorizeIndexGetByIDsResponseEnvelopeSuccess = true
)

type VectorizeIndexInsertResponseEnvelope struct {
	Errors   []VectorizeIndexInsertResponseEnvelopeErrors   `json:"errors,required"`
	Messages []VectorizeIndexInsertResponseEnvelopeMessages `json:"messages,required"`
	Result   VectorizeIndexInsert                           `json:"result,required,nullable"`
	// Whether the API call was successful
	Success VectorizeIndexInsertResponseEnvelopeSuccess `json:"success,required"`
	JSON    vectorizeIndexInsertResponseEnvelopeJSON    `json:"-"`
}

// vectorizeIndexInsertResponseEnvelopeJSON contains the JSON metadata for the
// struct [VectorizeIndexInsertResponseEnvelope]
type vectorizeIndexInsertResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexInsertResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexInsertResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    vectorizeIndexInsertResponseEnvelopeErrorsJSON `json:"-"`
}

// vectorizeIndexInsertResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [VectorizeIndexInsertResponseEnvelopeErrors]
type vectorizeIndexInsertResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexInsertResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexInsertResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    vectorizeIndexInsertResponseEnvelopeMessagesJSON `json:"-"`
}

// vectorizeIndexInsertResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [VectorizeIndexInsertResponseEnvelopeMessages]
type vectorizeIndexInsertResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexInsertResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type VectorizeIndexInsertResponseEnvelopeSuccess bool

const (
	VectorizeIndexInsertResponseEnvelopeSuccessTrue VectorizeIndexInsertResponseEnvelopeSuccess = true
)

type VectorizeIndexQueryParams struct {
	// Whether to return the metadata associated with the closest vectors.
	ReturnMetadata param.Field[bool] `json:"returnMetadata"`
	// Whether to return the values associated with the closest vectors.
	ReturnValues param.Field[bool] `json:"returnValues"`
	// The number of nearest neighbors to find.
	TopK param.Field[float64] `json:"topK"`
	// The search vector that will be used to find the nearest neighbors.
	Vector param.Field[[]float64] `json:"vector"`
}

func (r VectorizeIndexQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VectorizeIndexQueryResponseEnvelope struct {
	Errors   []VectorizeIndexQueryResponseEnvelopeErrors   `json:"errors,required"`
	Messages []VectorizeIndexQueryResponseEnvelopeMessages `json:"messages,required"`
	Result   VectorizeIndexQuery                           `json:"result,required,nullable"`
	// Whether the API call was successful
	Success VectorizeIndexQueryResponseEnvelopeSuccess `json:"success,required"`
	JSON    vectorizeIndexQueryResponseEnvelopeJSON    `json:"-"`
}

// vectorizeIndexQueryResponseEnvelopeJSON contains the JSON metadata for the
// struct [VectorizeIndexQueryResponseEnvelope]
type vectorizeIndexQueryResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexQueryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexQueryResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    vectorizeIndexQueryResponseEnvelopeErrorsJSON `json:"-"`
}

// vectorizeIndexQueryResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [VectorizeIndexQueryResponseEnvelopeErrors]
type vectorizeIndexQueryResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexQueryResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexQueryResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    vectorizeIndexQueryResponseEnvelopeMessagesJSON `json:"-"`
}

// vectorizeIndexQueryResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [VectorizeIndexQueryResponseEnvelopeMessages]
type vectorizeIndexQueryResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexQueryResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type VectorizeIndexQueryResponseEnvelopeSuccess bool

const (
	VectorizeIndexQueryResponseEnvelopeSuccessTrue VectorizeIndexQueryResponseEnvelopeSuccess = true
)

type VectorizeIndexUpsertResponseEnvelope struct {
	Errors   []VectorizeIndexUpsertResponseEnvelopeErrors   `json:"errors,required"`
	Messages []VectorizeIndexUpsertResponseEnvelopeMessages `json:"messages,required"`
	Result   VectorizeIndexUpsert                           `json:"result,required,nullable"`
	// Whether the API call was successful
	Success VectorizeIndexUpsertResponseEnvelopeSuccess `json:"success,required"`
	JSON    vectorizeIndexUpsertResponseEnvelopeJSON    `json:"-"`
}

// vectorizeIndexUpsertResponseEnvelopeJSON contains the JSON metadata for the
// struct [VectorizeIndexUpsertResponseEnvelope]
type vectorizeIndexUpsertResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexUpsertResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexUpsertResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    vectorizeIndexUpsertResponseEnvelopeErrorsJSON `json:"-"`
}

// vectorizeIndexUpsertResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [VectorizeIndexUpsertResponseEnvelopeErrors]
type vectorizeIndexUpsertResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexUpsertResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexUpsertResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    vectorizeIndexUpsertResponseEnvelopeMessagesJSON `json:"-"`
}

// vectorizeIndexUpsertResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [VectorizeIndexUpsertResponseEnvelopeMessages]
type vectorizeIndexUpsertResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexUpsertResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type VectorizeIndexUpsertResponseEnvelopeSuccess bool

const (
	VectorizeIndexUpsertResponseEnvelopeSuccessTrue VectorizeIndexUpsertResponseEnvelopeSuccess = true
)
