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
func (r *RadarQualitySpeedHistogramService) List(ctx context.Context, query RadarQualitySpeedHistogramListParams, opts ...option.RequestOption) (res *RadarQualitySpeedHistogramListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarQualitySpeedHistogramListResponseEnvelope
	path := "radar/quality/speed/histogram"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarQualitySpeedHistogramListResponse struct {
	Histogram0 RadarQualitySpeedHistogramListResponseHistogram0 `json:"histogram_0,required"`
	Meta       interface{}                                      `json:"meta,required"`
	JSON       radarQualitySpeedHistogramListResponseJSON       `json:"-"`
}

// radarQualitySpeedHistogramListResponseJSON contains the JSON metadata for the
// struct [RadarQualitySpeedHistogramListResponse]
type radarQualitySpeedHistogramListResponseJSON struct {
	Histogram0  apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedHistogramListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedHistogramListResponseHistogram0 struct {
	BandwidthDownload []string                                             `json:"bandwidthDownload,required"`
	BandwidthUpload   []string                                             `json:"bandwidthUpload,required"`
	BucketMin         []string                                             `json:"bucketMin,required"`
	JSON              radarQualitySpeedHistogramListResponseHistogram0JSON `json:"-"`
}

// radarQualitySpeedHistogramListResponseHistogram0JSON contains the JSON metadata
// for the struct [RadarQualitySpeedHistogramListResponseHistogram0]
type radarQualitySpeedHistogramListResponseHistogram0JSON struct {
	BandwidthDownload apijson.Field
	BandwidthUpload   apijson.Field
	BucketMin         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RadarQualitySpeedHistogramListResponseHistogram0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedHistogramListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// The width for every bucket in the histogram.
	BucketSize param.Field[int64] `query:"bucketSize"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarQualitySpeedHistogramListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Metrics to be returned.
	MetricGroup param.Field[RadarQualitySpeedHistogramListParamsMetricGroup] `query:"metricGroup"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarQualitySpeedHistogramListParams]'s query parameters as
// `url.Values`.
func (r RadarQualitySpeedHistogramListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarQualitySpeedHistogramListParamsFormat string

const (
	RadarQualitySpeedHistogramListParamsFormatJson RadarQualitySpeedHistogramListParamsFormat = "JSON"
	RadarQualitySpeedHistogramListParamsFormatCsv  RadarQualitySpeedHistogramListParamsFormat = "CSV"
)

// Metrics to be returned.
type RadarQualitySpeedHistogramListParamsMetricGroup string

const (
	RadarQualitySpeedHistogramListParamsMetricGroupBandwidth RadarQualitySpeedHistogramListParamsMetricGroup = "BANDWIDTH"
	RadarQualitySpeedHistogramListParamsMetricGroupLatency   RadarQualitySpeedHistogramListParamsMetricGroup = "LATENCY"
	RadarQualitySpeedHistogramListParamsMetricGroupJitter    RadarQualitySpeedHistogramListParamsMetricGroup = "JITTER"
)

type RadarQualitySpeedHistogramListResponseEnvelope struct {
	Result  RadarQualitySpeedHistogramListResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarQualitySpeedHistogramListResponseEnvelopeJSON `json:"-"`
}

// radarQualitySpeedHistogramListResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarQualitySpeedHistogramListResponseEnvelope]
type radarQualitySpeedHistogramListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedHistogramListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
