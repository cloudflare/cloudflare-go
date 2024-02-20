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
func (r *RadarDatasetDownloadService) New(ctx context.Context, params RadarDatasetDownloadNewParams, opts ...option.RequestOption) (res *RadarDatasetDownloadNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarDatasetDownloadNewResponseEnvelope
	path := "radar/datasets/download"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarDatasetDownloadNewResponse struct {
	Dataset RadarDatasetDownloadNewResponseDataset `json:"dataset,required"`
	JSON    radarDatasetDownloadNewResponseJSON    `json:"-"`
}

// radarDatasetDownloadNewResponseJSON contains the JSON metadata for the struct
// [RadarDatasetDownloadNewResponse]
type radarDatasetDownloadNewResponseJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDatasetDownloadNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDatasetDownloadNewResponseDataset struct {
	URL  string                                     `json:"url,required"`
	JSON radarDatasetDownloadNewResponseDatasetJSON `json:"-"`
}

// radarDatasetDownloadNewResponseDatasetJSON contains the JSON metadata for the
// struct [RadarDatasetDownloadNewResponseDataset]
type radarDatasetDownloadNewResponseDatasetJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDatasetDownloadNewResponseDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDatasetDownloadNewParams struct {
	DatasetID param.Field[int64] `json:"datasetId,required"`
	// Format results are returned in.
	Format param.Field[RadarDatasetDownloadNewParamsFormat] `query:"format"`
}

func (r RadarDatasetDownloadNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [RadarDatasetDownloadNewParams]'s query parameters as
// `url.Values`.
func (r RadarDatasetDownloadNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarDatasetDownloadNewParamsFormat string

const (
	RadarDatasetDownloadNewParamsFormatJson RadarDatasetDownloadNewParamsFormat = "JSON"
	RadarDatasetDownloadNewParamsFormatCsv  RadarDatasetDownloadNewParamsFormat = "CSV"
)

type RadarDatasetDownloadNewResponseEnvelope struct {
	Result  RadarDatasetDownloadNewResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarDatasetDownloadNewResponseEnvelopeJSON `json:"-"`
}

// radarDatasetDownloadNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarDatasetDownloadNewResponseEnvelope]
type radarDatasetDownloadNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDatasetDownloadNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
