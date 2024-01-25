// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarDatasetDownloadService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarDatasetDownloadService]
// method instead.
type RadarDatasetDownloadService struct {
	Options []option.RequestOption
}

// NewRadarDatasetDownloadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarDatasetDownloadService(opts ...option.RequestOption) (r *RadarDatasetDownloadService) {
	r = &RadarDatasetDownloadService{}
	r.Options = opts
	return
}

// Get a url to download a single dataset.
func (r *RadarDatasetDownloadService) RadarPostDatasetDownload(ctx context.Context, params RadarDatasetDownloadRadarPostDatasetDownloadParams, opts ...option.RequestOption) (res *RadarDatasetDownloadRadarPostDatasetDownloadResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/datasets/download"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type RadarDatasetDownloadRadarPostDatasetDownloadResponse struct {
	Result  RadarDatasetDownloadRadarPostDatasetDownloadResponseResult `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    radarDatasetDownloadRadarPostDatasetDownloadResponseJSON   `json:"-"`
}

// radarDatasetDownloadRadarPostDatasetDownloadResponseJSON contains the JSON
// metadata for the struct [RadarDatasetDownloadRadarPostDatasetDownloadResponse]
type radarDatasetDownloadRadarPostDatasetDownloadResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDatasetDownloadRadarPostDatasetDownloadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDatasetDownloadRadarPostDatasetDownloadResponseResult struct {
	Dataset RadarDatasetDownloadRadarPostDatasetDownloadResponseResultDataset `json:"dataset,required"`
	JSON    radarDatasetDownloadRadarPostDatasetDownloadResponseResultJSON    `json:"-"`
}

// radarDatasetDownloadRadarPostDatasetDownloadResponseResultJSON contains the JSON
// metadata for the struct
// [RadarDatasetDownloadRadarPostDatasetDownloadResponseResult]
type radarDatasetDownloadRadarPostDatasetDownloadResponseResultJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDatasetDownloadRadarPostDatasetDownloadResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDatasetDownloadRadarPostDatasetDownloadResponseResultDataset struct {
	URL  string                                                                `json:"url,required"`
	JSON radarDatasetDownloadRadarPostDatasetDownloadResponseResultDatasetJSON `json:"-"`
}

// radarDatasetDownloadRadarPostDatasetDownloadResponseResultDatasetJSON contains
// the JSON metadata for the struct
// [RadarDatasetDownloadRadarPostDatasetDownloadResponseResultDataset]
type radarDatasetDownloadRadarPostDatasetDownloadResponseResultDatasetJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDatasetDownloadRadarPostDatasetDownloadResponseResultDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDatasetDownloadRadarPostDatasetDownloadParams struct {
	DatasetID param.Field[int64] `json:"datasetId,required"`
	// Format results are returned in.
	Format param.Field[RadarDatasetDownloadRadarPostDatasetDownloadParamsFormat] `query:"format"`
}

func (r RadarDatasetDownloadRadarPostDatasetDownloadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [RadarDatasetDownloadRadarPostDatasetDownloadParams]'s query
// parameters as `url.Values`.
func (r RadarDatasetDownloadRadarPostDatasetDownloadParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarDatasetDownloadRadarPostDatasetDownloadParamsFormat string

const (
	RadarDatasetDownloadRadarPostDatasetDownloadParamsFormatJson RadarDatasetDownloadRadarPostDatasetDownloadParamsFormat = "JSON"
	RadarDatasetDownloadRadarPostDatasetDownloadParamsFormatCsv  RadarDatasetDownloadRadarPostDatasetDownloadParamsFormat = "CSV"
)
