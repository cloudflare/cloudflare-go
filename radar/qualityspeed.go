// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// QualitySpeedService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQualitySpeedService] method instead.
type QualitySpeedService struct {
	Options []option.RequestOption
	Top     *QualitySpeedTopService
}

// NewQualitySpeedService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQualitySpeedService(opts ...option.RequestOption) (r *QualitySpeedService) {
	r = &QualitySpeedService{}
	r.Options = opts
	r.Top = NewQualitySpeedTopService(opts...)
	return
}

// Get an histogram from the previous 90 days of Cloudflare Speed Test data, split
// into fixed bandwidth (Mbps), latency (ms) or jitter (ms) buckets.
func (r *QualitySpeedService) Histogram(ctx context.Context, query QualitySpeedHistogramParams, opts ...option.RequestOption) (res *QualitySpeedHistogramResponse, err error) {
	var env QualitySpeedHistogramResponseEnvelope
	opts = append(r.Options[:], opts...)
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
func (r *QualitySpeedService) Summary(ctx context.Context, query QualitySpeedSummaryParams, opts ...option.RequestOption) (res *QualitySpeedSummaryResponse, err error) {
	var env QualitySpeedSummaryResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/quality/speed/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type QualitySpeedHistogramResponse struct {
	Histogram0 QualitySpeedHistogramResponseHistogram0 `json:"histogram_0,required"`
	Meta       interface{}                             `json:"meta,required"`
	JSON       qualitySpeedHistogramResponseJSON       `json:"-"`
}

// qualitySpeedHistogramResponseJSON contains the JSON metadata for the struct
// [QualitySpeedHistogramResponse]
type qualitySpeedHistogramResponseJSON struct {
	Histogram0  apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedHistogramResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedHistogramResponseJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedHistogramResponseHistogram0 struct {
	BandwidthDownload []string                                    `json:"bandwidthDownload,required"`
	BandwidthUpload   []string                                    `json:"bandwidthUpload,required"`
	BucketMin         []string                                    `json:"bucketMin,required"`
	JSON              qualitySpeedHistogramResponseHistogram0JSON `json:"-"`
}

// qualitySpeedHistogramResponseHistogram0JSON contains the JSON metadata for the
// struct [QualitySpeedHistogramResponseHistogram0]
type qualitySpeedHistogramResponseHistogram0JSON struct {
	BandwidthDownload apijson.Field
	BandwidthUpload   apijson.Field
	BucketMin         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *QualitySpeedHistogramResponseHistogram0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedHistogramResponseHistogram0JSON) RawJSON() string {
	return r.raw
}

type QualitySpeedSummaryResponse struct {
	Meta     QualitySpeedSummaryResponseMeta     `json:"meta,required"`
	Summary0 QualitySpeedSummaryResponseSummary0 `json:"summary_0,required"`
	JSON     qualitySpeedSummaryResponseJSON     `json:"-"`
}

// qualitySpeedSummaryResponseJSON contains the JSON metadata for the struct
// [QualitySpeedSummaryResponse]
type qualitySpeedSummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedSummaryResponseMeta struct {
	DateRange      []QualitySpeedSummaryResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                        `json:"lastUpdated,required"`
	Normalization  string                                        `json:"normalization,required"`
	ConfidenceInfo QualitySpeedSummaryResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           qualitySpeedSummaryResponseMetaJSON           `json:"-"`
}

// qualitySpeedSummaryResponseMetaJSON contains the JSON metadata for the struct
// [QualitySpeedSummaryResponseMeta]
type qualitySpeedSummaryResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *QualitySpeedSummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedSummaryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedSummaryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                    `json:"startTime,required" format:"date-time"`
	JSON      qualitySpeedSummaryResponseMetaDateRangeJSON `json:"-"`
}

// qualitySpeedSummaryResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [QualitySpeedSummaryResponseMetaDateRange]
type qualitySpeedSummaryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedSummaryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedSummaryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedSummaryResponseMetaConfidenceInfo struct {
	Annotations []QualitySpeedSummaryResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                     `json:"level"`
	JSON        qualitySpeedSummaryResponseMetaConfidenceInfoJSON         `json:"-"`
}

// qualitySpeedSummaryResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [QualitySpeedSummaryResponseMetaConfidenceInfo]
type qualitySpeedSummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedSummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedSummaryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedSummaryResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                      `json:"dataSource,required"`
	Description     string                                                      `json:"description,required"`
	EventType       string                                                      `json:"eventType,required"`
	IsInstantaneous bool                                                        `json:"isInstantaneous,required"`
	EndTime         time.Time                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                   `json:"startTime" format:"date-time"`
	JSON            qualitySpeedSummaryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// qualitySpeedSummaryResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [QualitySpeedSummaryResponseMetaConfidenceInfoAnnotation]
type qualitySpeedSummaryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *QualitySpeedSummaryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedSummaryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedSummaryResponseSummary0 struct {
	BandwidthDownload string                                  `json:"bandwidthDownload,required"`
	BandwidthUpload   string                                  `json:"bandwidthUpload,required"`
	JitterIdle        string                                  `json:"jitterIdle,required"`
	JitterLoaded      string                                  `json:"jitterLoaded,required"`
	LatencyIdle       string                                  `json:"latencyIdle,required"`
	LatencyLoaded     string                                  `json:"latencyLoaded,required"`
	PacketLoss        string                                  `json:"packetLoss,required"`
	JSON              qualitySpeedSummaryResponseSummary0JSON `json:"-"`
}

// qualitySpeedSummaryResponseSummary0JSON contains the JSON metadata for the
// struct [QualitySpeedSummaryResponseSummary0]
type qualitySpeedSummaryResponseSummary0JSON struct {
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

func (r *QualitySpeedSummaryResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedSummaryResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type QualitySpeedHistogramParams struct {
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
	Format param.Field[QualitySpeedHistogramParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Metrics to be returned.
	MetricGroup param.Field[QualitySpeedHistogramParamsMetricGroup] `query:"metricGroup"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [QualitySpeedHistogramParams]'s query parameters as
// `url.Values`.
func (r QualitySpeedHistogramParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type QualitySpeedHistogramParamsFormat string

const (
	QualitySpeedHistogramParamsFormatJson QualitySpeedHistogramParamsFormat = "JSON"
	QualitySpeedHistogramParamsFormatCsv  QualitySpeedHistogramParamsFormat = "CSV"
)

func (r QualitySpeedHistogramParamsFormat) IsKnown() bool {
	switch r {
	case QualitySpeedHistogramParamsFormatJson, QualitySpeedHistogramParamsFormatCsv:
		return true
	}
	return false
}

// Metrics to be returned.
type QualitySpeedHistogramParamsMetricGroup string

const (
	QualitySpeedHistogramParamsMetricGroupBandwidth QualitySpeedHistogramParamsMetricGroup = "BANDWIDTH"
	QualitySpeedHistogramParamsMetricGroupLatency   QualitySpeedHistogramParamsMetricGroup = "LATENCY"
	QualitySpeedHistogramParamsMetricGroupJitter    QualitySpeedHistogramParamsMetricGroup = "JITTER"
)

func (r QualitySpeedHistogramParamsMetricGroup) IsKnown() bool {
	switch r {
	case QualitySpeedHistogramParamsMetricGroupBandwidth, QualitySpeedHistogramParamsMetricGroupLatency, QualitySpeedHistogramParamsMetricGroupJitter:
		return true
	}
	return false
}

type QualitySpeedHistogramResponseEnvelope struct {
	Result  QualitySpeedHistogramResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    qualitySpeedHistogramResponseEnvelopeJSON `json:"-"`
}

// qualitySpeedHistogramResponseEnvelopeJSON contains the JSON metadata for the
// struct [QualitySpeedHistogramResponseEnvelope]
type qualitySpeedHistogramResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedHistogramResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedHistogramResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedSummaryParams struct {
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
	Format param.Field[QualitySpeedSummaryParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [QualitySpeedSummaryParams]'s query parameters as
// `url.Values`.
func (r QualitySpeedSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type QualitySpeedSummaryParamsFormat string

const (
	QualitySpeedSummaryParamsFormatJson QualitySpeedSummaryParamsFormat = "JSON"
	QualitySpeedSummaryParamsFormatCsv  QualitySpeedSummaryParamsFormat = "CSV"
)

func (r QualitySpeedSummaryParamsFormat) IsKnown() bool {
	switch r {
	case QualitySpeedSummaryParamsFormatJson, QualitySpeedSummaryParamsFormatCsv:
		return true
	}
	return false
}

type QualitySpeedSummaryResponseEnvelope struct {
	Result  QualitySpeedSummaryResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    qualitySpeedSummaryResponseEnvelopeJSON `json:"-"`
}

// qualitySpeedSummaryResponseEnvelopeJSON contains the JSON metadata for the
// struct [QualitySpeedSummaryResponseEnvelope]
type qualitySpeedSummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedSummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
