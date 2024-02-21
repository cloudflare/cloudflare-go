// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarDatasetService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarDatasetService] method
// instead.
type RadarDatasetService struct {
	Options []option.RequestOption
}

// NewRadarDatasetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarDatasetService(opts ...option.RequestOption) (r *RadarDatasetService) {
	r = &RadarDatasetService{}
	r.Options = opts
	return
}

// Get a list of datasets.
func (r *RadarDatasetService) List(ctx context.Context, query RadarDatasetListParams, opts ...option.RequestOption) (res *RadarDatasetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarDatasetListResponseEnvelope
	path := "radar/datasets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a url to download a single dataset.
func (r *RadarDatasetService) Download(ctx context.Context, params RadarDatasetDownloadParams, opts ...option.RequestOption) (res *RadarDatasetDownloadResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarDatasetDownloadResponseEnvelope
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
func (r *RadarDatasetService) Get(ctx context.Context, alias string, query RadarDatasetGetParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/csv")}, opts...)
	path := fmt.Sprintf("radar/datasets/%s", alias)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarDatasetListResponse struct {
	Datasets []RadarDatasetListResponseDataset `json:"datasets,required"`
	JSON     radarDatasetListResponseJSON      `json:"-"`
}

// radarDatasetListResponseJSON contains the JSON metadata for the struct
// [RadarDatasetListResponse]
type radarDatasetListResponseJSON struct {
	Datasets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDatasetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDatasetListResponseDataset struct {
	ID          int64                               `json:"id,required"`
	Description string                              `json:"description,required"`
	Meta        interface{}                         `json:"meta,required"`
	Tags        []string                            `json:"tags,required"`
	Title       string                              `json:"title,required"`
	Type        string                              `json:"type,required"`
	JSON        radarDatasetListResponseDatasetJSON `json:"-"`
}

// radarDatasetListResponseDatasetJSON contains the JSON metadata for the struct
// [RadarDatasetListResponseDataset]
type radarDatasetListResponseDatasetJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Meta        apijson.Field
	Tags        apijson.Field
	Title       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDatasetListResponseDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDatasetDownloadResponse struct {
	Dataset RadarDatasetDownloadResponseDataset `json:"dataset,required"`
	JSON    radarDatasetDownloadResponseJSON    `json:"-"`
}

// radarDatasetDownloadResponseJSON contains the JSON metadata for the struct
// [RadarDatasetDownloadResponse]
type radarDatasetDownloadResponseJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDatasetDownloadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDatasetDownloadResponseDataset struct {
	URL  string                                  `json:"url,required"`
	JSON radarDatasetDownloadResponseDatasetJSON `json:"-"`
}

// radarDatasetDownloadResponseDatasetJSON contains the JSON metadata for the
// struct [RadarDatasetDownloadResponseDataset]
type radarDatasetDownloadResponseDatasetJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDatasetDownloadResponseDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDatasetListParams struct {
	// Dataset type.
	DatasetType param.Field[RadarDatasetListParamsDatasetType] `query:"datasetType"`
	// Format results are returned in.
	Format param.Field[RadarDatasetListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Number of objects to skip before grabbing results.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [RadarDatasetListParams]'s query parameters as `url.Values`.
func (r RadarDatasetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Dataset type.
type RadarDatasetListParamsDatasetType string

const (
	RadarDatasetListParamsDatasetTypeRankingBucket RadarDatasetListParamsDatasetType = "RANKING_BUCKET"
	RadarDatasetListParamsDatasetTypeReport        RadarDatasetListParamsDatasetType = "REPORT"
)

// Format results are returned in.
type RadarDatasetListParamsFormat string

const (
	RadarDatasetListParamsFormatJson RadarDatasetListParamsFormat = "JSON"
	RadarDatasetListParamsFormatCsv  RadarDatasetListParamsFormat = "CSV"
)

type RadarDatasetListResponseEnvelope struct {
	Result  RadarDatasetListResponse             `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    radarDatasetListResponseEnvelopeJSON `json:"-"`
}

// radarDatasetListResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarDatasetListResponseEnvelope]
type radarDatasetListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDatasetListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDatasetDownloadParams struct {
	DatasetID param.Field[int64] `json:"datasetId,required"`
	// Format results are returned in.
	Format param.Field[RadarDatasetDownloadParamsFormat] `query:"format"`
}

func (r RadarDatasetDownloadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [RadarDatasetDownloadParams]'s query parameters as
// `url.Values`.
func (r RadarDatasetDownloadParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarDatasetDownloadParamsFormat string

const (
	RadarDatasetDownloadParamsFormatJson RadarDatasetDownloadParamsFormat = "JSON"
	RadarDatasetDownloadParamsFormatCsv  RadarDatasetDownloadParamsFormat = "CSV"
)

type RadarDatasetDownloadResponseEnvelope struct {
	Result  RadarDatasetDownloadResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarDatasetDownloadResponseEnvelopeJSON `json:"-"`
}

// radarDatasetDownloadResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarDatasetDownloadResponseEnvelope]
type radarDatasetDownloadResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDatasetDownloadResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDatasetGetParams struct {
	// Filter dataset alias by date
	Date param.Field[string] `query:"date"`
}

// URLQuery serializes [RadarDatasetGetParams]'s query parameters as `url.Values`.
func (r RadarDatasetGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
