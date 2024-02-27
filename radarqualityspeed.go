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

// RadarQualitySpeedService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarQualitySpeedService] method
// instead.
type RadarQualitySpeedService struct {
	Options []option.RequestOption
	Top     *RadarQualitySpeedTopService
}

// NewRadarQualitySpeedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarQualitySpeedService(opts ...option.RequestOption) (r *RadarQualitySpeedService) {
	r = &RadarQualitySpeedService{}
	r.Options = opts
	r.Top = NewRadarQualitySpeedTopService(opts...)
	return
}

// Get an histogram from the previous 90 days of Cloudflare Speed Test data, split
// into fixed bandwidth (Mbps), latency (ms) or jitter (ms) buckets.
func (r *RadarQualitySpeedService) Histogram(ctx context.Context, query RadarQualitySpeedHistogramParams, opts ...option.RequestOption) (res *RadarQualitySpeedHistogramResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarQualitySpeedHistogramResponseEnvelope
	path := "radar/quality/speed/histogram"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a summary of bandwidth, latency, jitter and packet loss, from the previous
// 90 days of Cloudflare Speed Test data.
func (r *RadarQualitySpeedService) Summary(ctx context.Context, query RadarQualitySpeedSummaryParams, opts ...option.RequestOption) (res *RadarQualitySpeedSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarQualitySpeedSummaryResponseEnvelope
	path := "radar/quality/speed/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarQualitySpeedHistogramResponse struct {
	Histogram0 RadarQualitySpeedHistogramResponseHistogram0 `json:"histogram_0,required"`
	Meta       interface{}                                  `json:"meta,required"`
	JSON       radarQualitySpeedHistogramResponseJSON       `json:"-"`
}

// radarQualitySpeedHistogramResponseJSON contains the JSON metadata for the struct
// [RadarQualitySpeedHistogramResponse]
type radarQualitySpeedHistogramResponseJSON struct {
	Histogram0  apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedHistogramResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedHistogramResponseHistogram0 struct {
	BandwidthDownload []string                                         `json:"bandwidthDownload,required"`
	BandwidthUpload   []string                                         `json:"bandwidthUpload,required"`
	BucketMin         []string                                         `json:"bucketMin,required"`
	JSON              radarQualitySpeedHistogramResponseHistogram0JSON `json:"-"`
}

// radarQualitySpeedHistogramResponseHistogram0JSON contains the JSON metadata for
// the struct [RadarQualitySpeedHistogramResponseHistogram0]
type radarQualitySpeedHistogramResponseHistogram0JSON struct {
	BandwidthDownload apijson.Field
	BandwidthUpload   apijson.Field
	BucketMin         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RadarQualitySpeedHistogramResponseHistogram0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryResponse struct {
	Meta     RadarQualitySpeedSummaryResponseMeta     `json:"meta,required"`
	Summary0 RadarQualitySpeedSummaryResponseSummary0 `json:"summary_0,required"`
	JSON     radarQualitySpeedSummaryResponseJSON     `json:"-"`
}

// radarQualitySpeedSummaryResponseJSON contains the JSON metadata for the struct
// [RadarQualitySpeedSummaryResponse]
type radarQualitySpeedSummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryResponseMeta struct {
	DateRange      []RadarQualitySpeedSummaryResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                             `json:"lastUpdated,required"`
	Normalization  string                                             `json:"normalization,required"`
	ConfidenceInfo RadarQualitySpeedSummaryResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarQualitySpeedSummaryResponseMetaJSON           `json:"-"`
}

// radarQualitySpeedSummaryResponseMetaJSON contains the JSON metadata for the
// struct [RadarQualitySpeedSummaryResponseMeta]
type radarQualitySpeedSummaryResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarQualitySpeedSummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      radarQualitySpeedSummaryResponseMetaDateRangeJSON `json:"-"`
}

// radarQualitySpeedSummaryResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarQualitySpeedSummaryResponseMetaDateRange]
type radarQualitySpeedSummaryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedSummaryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryResponseMetaConfidenceInfo struct {
	Annotations []RadarQualitySpeedSummaryResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                          `json:"level"`
	JSON        radarQualitySpeedSummaryResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarQualitySpeedSummaryResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarQualitySpeedSummaryResponseMetaConfidenceInfo]
type radarQualitySpeedSummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedSummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                           `json:"dataSource,required"`
	Description     string                                                           `json:"description,required"`
	EventType       string                                                           `json:"eventType,required"`
	IsInstantaneous interface{}                                                      `json:"isInstantaneous,required"`
	EndTime         time.Time                                                        `json:"endTime" format:"date-time"`
	LinkedURL       string                                                           `json:"linkedUrl"`
	StartTime       time.Time                                                        `json:"startTime" format:"date-time"`
	JSON            radarQualitySpeedSummaryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarQualitySpeedSummaryResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarQualitySpeedSummaryResponseMetaConfidenceInfoAnnotation]
type radarQualitySpeedSummaryResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	EndTime         apijson.Field
	LinkedURL       apijson.Field
	StartTime       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarQualitySpeedSummaryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryResponseSummary0 struct {
	BandwidthDownload string                                       `json:"bandwidthDownload,required"`
	BandwidthUpload   string                                       `json:"bandwidthUpload,required"`
	JitterIdle        string                                       `json:"jitterIdle,required"`
	JitterLoaded      string                                       `json:"jitterLoaded,required"`
	LatencyIdle       string                                       `json:"latencyIdle,required"`
	LatencyLoaded     string                                       `json:"latencyLoaded,required"`
	PacketLoss        string                                       `json:"packetLoss,required"`
	JSON              radarQualitySpeedSummaryResponseSummary0JSON `json:"-"`
}

// radarQualitySpeedSummaryResponseSummary0JSON contains the JSON metadata for the
// struct [RadarQualitySpeedSummaryResponseSummary0]
type radarQualitySpeedSummaryResponseSummary0JSON struct {
	BandwidthDownload apijson.Field
	BandwidthUpload   apijson.Field
	JitterIdle        apijson.Field
	JitterLoaded      apijson.Field
	LatencyIdle       apijson.Field
	LatencyLoaded     apijson.Field
	PacketLoss        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RadarQualitySpeedSummaryResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedHistogramParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// The width for every bucket in the histogram.
	BucketSize param.Field[int64] `query:"bucketSize"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarQualitySpeedHistogramParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Metrics to be returned.
	MetricGroup param.Field[RadarQualitySpeedHistogramParamsMetricGroup] `query:"metricGroup"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarQualitySpeedHistogramParams]'s query parameters as
// `url.Values`.
func (r RadarQualitySpeedHistogramParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarQualitySpeedHistogramParamsFormat string

const (
	RadarQualitySpeedHistogramParamsFormatJson RadarQualitySpeedHistogramParamsFormat = "JSON"
	RadarQualitySpeedHistogramParamsFormatCsv  RadarQualitySpeedHistogramParamsFormat = "CSV"
)

// Metrics to be returned.
type RadarQualitySpeedHistogramParamsMetricGroup string

const (
	RadarQualitySpeedHistogramParamsMetricGroupBandwidth RadarQualitySpeedHistogramParamsMetricGroup = "BANDWIDTH"
	RadarQualitySpeedHistogramParamsMetricGroupLatency   RadarQualitySpeedHistogramParamsMetricGroup = "LATENCY"
	RadarQualitySpeedHistogramParamsMetricGroupJitter    RadarQualitySpeedHistogramParamsMetricGroup = "JITTER"
)

type RadarQualitySpeedHistogramResponseEnvelope struct {
	Result  RadarQualitySpeedHistogramResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarQualitySpeedHistogramResponseEnvelopeJSON `json:"-"`
}

// radarQualitySpeedHistogramResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarQualitySpeedHistogramResponseEnvelope]
type radarQualitySpeedHistogramResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedHistogramResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarQualitySpeedSummaryParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarQualitySpeedSummaryParams]'s query parameters as
// `url.Values`.
func (r RadarQualitySpeedSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarQualitySpeedSummaryParamsFormat string

const (
	RadarQualitySpeedSummaryParamsFormatJson RadarQualitySpeedSummaryParamsFormat = "JSON"
	RadarQualitySpeedSummaryParamsFormatCsv  RadarQualitySpeedSummaryParamsFormat = "CSV"
)

type RadarQualitySpeedSummaryResponseEnvelope struct {
	Result  RadarQualitySpeedSummaryResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarQualitySpeedSummaryResponseEnvelopeJSON `json:"-"`
}

// radarQualitySpeedSummaryResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarQualitySpeedSummaryResponseEnvelope]
type radarQualitySpeedSummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
