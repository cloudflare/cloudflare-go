// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DatasetService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDatasetService] method instead.
type DatasetService struct {
	Options []option.RequestOption
}

// NewDatasetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDatasetService(opts ...option.RequestOption) (r *DatasetService) {
	r = &DatasetService{}
	r.Options = opts
	return
}

// Get a list of datasets.
func (r *DatasetService) List(ctx context.Context, query DatasetListParams, opts ...option.RequestOption) (res *DatasetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DatasetListResponseEnvelope
	path := "radar/datasets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a url to download a single dataset.
func (r *DatasetService) Download(ctx context.Context, params DatasetDownloadParams, opts ...option.RequestOption) (res *DatasetDownloadResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DatasetDownloadResponseEnvelope
	path := "radar/datasets/download"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the csv content of a given dataset by alias or id. When getting the content
// by alias the latest dataset is returned, optionally filtered by the latest
// available at a given date.
func (r *DatasetService) Get(ctx context.Context, alias string, query DatasetGetParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/csv")}, opts...)
	path := fmt.Sprintf("radar/datasets/%s", alias)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DatasetListResponse struct {
	Datasets []DatasetListResponseDataset `json:"datasets,required"`
	JSON     datasetListResponseJSON      `json:"-"`
}

// datasetListResponseJSON contains the JSON metadata for the struct
// [DatasetListResponse]
type datasetListResponseJSON struct {
	Datasets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetListResponseJSON) RawJSON() string {
	return r.raw
}

type DatasetListResponseDataset struct {
	ID          int64                          `json:"id,required"`
	Description string                         `json:"description,required"`
	Meta        interface{}                    `json:"meta,required"`
	Tags        []string                       `json:"tags,required"`
	Title       string                         `json:"title,required"`
	Type        string                         `json:"type,required"`
	JSON        datasetListResponseDatasetJSON `json:"-"`
}

// datasetListResponseDatasetJSON contains the JSON metadata for the struct
// [DatasetListResponseDataset]
type datasetListResponseDatasetJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Meta        apijson.Field
	Tags        apijson.Field
	Title       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetListResponseDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetListResponseDatasetJSON) RawJSON() string {
	return r.raw
}

type DatasetDownloadResponse struct {
	Dataset DatasetDownloadResponseDataset `json:"dataset,required"`
	JSON    datasetDownloadResponseJSON    `json:"-"`
}

// datasetDownloadResponseJSON contains the JSON metadata for the struct
// [DatasetDownloadResponse]
type datasetDownloadResponseJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetDownloadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetDownloadResponseJSON) RawJSON() string {
	return r.raw
}

type DatasetDownloadResponseDataset struct {
	URL  string                             `json:"url,required"`
	JSON datasetDownloadResponseDatasetJSON `json:"-"`
}

// datasetDownloadResponseDatasetJSON contains the JSON metadata for the struct
// [DatasetDownloadResponseDataset]
type datasetDownloadResponseDatasetJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetDownloadResponseDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetDownloadResponseDatasetJSON) RawJSON() string {
	return r.raw
}

type DatasetListParams struct {
	// Dataset type.
	DatasetType param.Field[DatasetListParamsDatasetType] `query:"datasetType"`
	// Format results are returned in.
	Format param.Field[DatasetListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Number of objects to skip before grabbing results.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [DatasetListParams]'s query parameters as `url.Values`.
func (r DatasetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Dataset type.
type DatasetListParamsDatasetType string

const (
	DatasetListParamsDatasetTypeRankingBucket DatasetListParamsDatasetType = "RANKING_BUCKET"
	DatasetListParamsDatasetTypeReport        DatasetListParamsDatasetType = "REPORT"
)

func (r DatasetListParamsDatasetType) IsKnown() bool {
	switch r {
	case DatasetListParamsDatasetTypeRankingBucket, DatasetListParamsDatasetTypeReport:
		return true
	}
	return false
}

// Format results are returned in.
type DatasetListParamsFormat string

const (
	DatasetListParamsFormatJson DatasetListParamsFormat = "JSON"
	DatasetListParamsFormatCsv  DatasetListParamsFormat = "CSV"
)

func (r DatasetListParamsFormat) IsKnown() bool {
	switch r {
	case DatasetListParamsFormatJson, DatasetListParamsFormatCsv:
		return true
	}
	return false
}

type DatasetListResponseEnvelope struct {
	Result  DatasetListResponse             `json:"result,required"`
	Success bool                            `json:"success,required"`
	JSON    datasetListResponseEnvelopeJSON `json:"-"`
}

// datasetListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatasetListResponseEnvelope]
type datasetListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DatasetDownloadParams struct {
	DatasetID param.Field[int64] `json:"datasetId,required"`
	// Format results are returned in.
	Format param.Field[DatasetDownloadParamsFormat] `query:"format"`
}

func (r DatasetDownloadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [DatasetDownloadParams]'s query parameters as `url.Values`.
func (r DatasetDownloadParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type DatasetDownloadParamsFormat string

const (
	DatasetDownloadParamsFormatJson DatasetDownloadParamsFormat = "JSON"
	DatasetDownloadParamsFormatCsv  DatasetDownloadParamsFormat = "CSV"
)

func (r DatasetDownloadParamsFormat) IsKnown() bool {
	switch r {
	case DatasetDownloadParamsFormatJson, DatasetDownloadParamsFormatCsv:
		return true
	}
	return false
}

type DatasetDownloadResponseEnvelope struct {
	Result  DatasetDownloadResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    datasetDownloadResponseEnvelopeJSON `json:"-"`
}

// datasetDownloadResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatasetDownloadResponseEnvelope]
type datasetDownloadResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetDownloadResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetDownloadResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DatasetGetParams struct {
	// Filter dataset alias by date
	Date param.Field[string] `query:"date"`
}

// URLQuery serializes [DatasetGetParams]'s query parameters as `url.Values`.
func (r DatasetGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
