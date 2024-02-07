// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarQualitySpeedHistogramService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarQualitySpeedHistogramService] method instead.
type RadarQualitySpeedHistogramService struct {
	Options []option.RequestOption
}

// NewRadarQualitySpeedHistogramService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarQualitySpeedHistogramService(opts ...option.RequestOption) (r *RadarQualitySpeedHistogramService) {
	r = &RadarQualitySpeedHistogramService{}
	r.Options = opts
	return
}

// Get an histogram from the previous 90 days of Cloudflare Speed Test data, split
// into fixed bandwidth (Mbps), latency (ms) or jitter (ms) buckets.
func (r *RadarQualitySpeedHistogramService) Get(ctx context.Context, query RadarQualitySpeedHistogramGetParams, opts ...option.RequestOption) (res *RadarQualitySpeedHistogramGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/quality/speed/histogram"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarQualitySpeedHistogramGetResponse struct {
	Result  RadarQualitySpeedHistogramGetResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarQualitySpeedHistogramGetResponseJSON   `json:"-"`
}

// radarQualitySpeedHistogramGetResponseJSON contains the JSON metadata for the
// struct [RadarQualitySpeedHistogramGetResponse]
type radarQualitySpeedHistogramGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedHistogramGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedHistogramGetResponseResult struct {
	Histogram0 RadarQualitySpeedHistogramGetResponseResultHistogram0 `json:"histogram_0,required"`
	Meta       interface{}                                           `json:"meta,required"`
	JSON       radarQualitySpeedHistogramGetResponseResultJSON       `json:"-"`
}

// radarQualitySpeedHistogramGetResponseResultJSON contains the JSON metadata for
// the struct [RadarQualitySpeedHistogramGetResponseResult]
type radarQualitySpeedHistogramGetResponseResultJSON struct {
	Histogram0  apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedHistogramGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedHistogramGetResponseResultHistogram0 struct {
	BandwidthDownload []string                                                  `json:"bandwidthDownload,required"`
	BandwidthUpload   []string                                                  `json:"bandwidthUpload,required"`
	BucketMin         []string                                                  `json:"bucketMin,required"`
	JSON              radarQualitySpeedHistogramGetResponseResultHistogram0JSON `json:"-"`
}

// radarQualitySpeedHistogramGetResponseResultHistogram0JSON contains the JSON
// metadata for the struct [RadarQualitySpeedHistogramGetResponseResultHistogram0]
type radarQualitySpeedHistogramGetResponseResultHistogram0JSON struct {
	BandwidthDownload apijson.Field
	BandwidthUpload   apijson.Field
	BucketMin         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RadarQualitySpeedHistogramGetResponseResultHistogram0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedHistogramGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// The width for every bucket in the histogram.
	BucketSize param.Field[int64] `query:"bucketSize"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarQualitySpeedHistogramGetParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Metrics to be returned.
	MetricGroup param.Field[RadarQualitySpeedHistogramGetParamsMetricGroup] `query:"metricGroup"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarQualitySpeedHistogramGetParams]'s query parameters as
// `url.Values`.
func (r RadarQualitySpeedHistogramGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarQualitySpeedHistogramGetParamsFormat string

const (
	RadarQualitySpeedHistogramGetParamsFormatJson RadarQualitySpeedHistogramGetParamsFormat = "JSON"
	RadarQualitySpeedHistogramGetParamsFormatCsv  RadarQualitySpeedHistogramGetParamsFormat = "CSV"
)

// Metrics to be returned.
type RadarQualitySpeedHistogramGetParamsMetricGroup string

const (
	RadarQualitySpeedHistogramGetParamsMetricGroupBandwidth RadarQualitySpeedHistogramGetParamsMetricGroup = "BANDWIDTH"
	RadarQualitySpeedHistogramGetParamsMetricGroupLatency   RadarQualitySpeedHistogramGetParamsMetricGroup = "LATENCY"
	RadarQualitySpeedHistogramGetParamsMetricGroupJitter    RadarQualitySpeedHistogramGetParamsMetricGroup = "JITTER"
)
