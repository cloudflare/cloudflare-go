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

// QualitySpeedTopService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQualitySpeedTopService] method instead.
type QualitySpeedTopService struct {
	Options []option.RequestOption
}

// NewQualitySpeedTopService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQualitySpeedTopService(opts ...option.RequestOption) (r *QualitySpeedTopService) {
	r = &QualitySpeedTopService{}
	r.Options = opts
	return
}

// Get the top autonomous systems by bandwidth, latency, jitter or packet loss,
// from the previous 90 days of Cloudflare Speed Test data.
func (r *QualitySpeedTopService) Ases(ctx context.Context, query QualitySpeedTopAsesParams, opts ...option.RequestOption) (res *QualitySpeedTopAsesResponse, err error) {
	var env QualitySpeedTopAsesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/quality/speed/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the top locations by bandwidth, latency, jitter or packet loss, from the
// previous 90 days of Cloudflare Speed Test data.
func (r *QualitySpeedTopService) Locations(ctx context.Context, query QualitySpeedTopLocationsParams, opts ...option.RequestOption) (res *QualitySpeedTopLocationsResponse, err error) {
	var env QualitySpeedTopLocationsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/quality/speed/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type QualitySpeedTopAsesResponse struct {
	Meta QualitySpeedTopAsesResponseMeta   `json:"meta,required"`
	Top0 []QualitySpeedTopAsesResponseTop0 `json:"top_0,required"`
	JSON qualitySpeedTopAsesResponseJSON   `json:"-"`
}

// qualitySpeedTopAsesResponseJSON contains the JSON metadata for the struct
// [QualitySpeedTopAsesResponse]
type qualitySpeedTopAsesResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedTopAsesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopAsesResponseJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopAsesResponseMeta struct {
	DateRange      []QualitySpeedTopAsesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                        `json:"lastUpdated,required"`
	ConfidenceInfo QualitySpeedTopAsesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           qualitySpeedTopAsesResponseMetaJSON           `json:"-"`
}

// qualitySpeedTopAsesResponseMetaJSON contains the JSON metadata for the struct
// [QualitySpeedTopAsesResponseMeta]
type qualitySpeedTopAsesResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *QualitySpeedTopAsesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopAsesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopAsesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                    `json:"startTime,required" format:"date-time"`
	JSON      qualitySpeedTopAsesResponseMetaDateRangeJSON `json:"-"`
}

// qualitySpeedTopAsesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [QualitySpeedTopAsesResponseMetaDateRange]
type qualitySpeedTopAsesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedTopAsesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopAsesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopAsesResponseMetaConfidenceInfo struct {
	Annotations []QualitySpeedTopAsesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                     `json:"level"`
	JSON        qualitySpeedTopAsesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// qualitySpeedTopAsesResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [QualitySpeedTopAsesResponseMetaConfidenceInfo]
type qualitySpeedTopAsesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedTopAsesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopAsesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopAsesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                      `json:"dataSource,required"`
	Description     string                                                      `json:"description,required"`
	EventType       string                                                      `json:"eventType,required"`
	IsInstantaneous bool                                                        `json:"isInstantaneous,required"`
	EndTime         time.Time                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                   `json:"startTime" format:"date-time"`
	JSON            qualitySpeedTopAsesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// qualitySpeedTopAsesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [QualitySpeedTopAsesResponseMetaConfidenceInfoAnnotation]
type qualitySpeedTopAsesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *QualitySpeedTopAsesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopAsesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopAsesResponseTop0 struct {
	BandwidthDownload string                              `json:"bandwidthDownload,required"`
	BandwidthUpload   string                              `json:"bandwidthUpload,required"`
	ClientASN         float64                             `json:"clientASN,required"`
	ClientAsName      string                              `json:"clientASName,required"`
	JitterIdle        string                              `json:"jitterIdle,required"`
	JitterLoaded      string                              `json:"jitterLoaded,required"`
	LatencyIdle       string                              `json:"latencyIdle,required"`
	LatencyLoaded     string                              `json:"latencyLoaded,required"`
	NumTests          float64                             `json:"numTests,required"`
	RankPower         float64                             `json:"rankPower,required"`
	JSON              qualitySpeedTopAsesResponseTop0JSON `json:"-"`
}

// qualitySpeedTopAsesResponseTop0JSON contains the JSON metadata for the struct
// [QualitySpeedTopAsesResponseTop0]
type qualitySpeedTopAsesResponseTop0JSON struct {
	BandwidthDownload apijson.Field
	BandwidthUpload   apijson.Field
	ClientASN         apijson.Field
	ClientAsName      apijson.Field
	JitterIdle        apijson.Field
	JitterLoaded      apijson.Field
	LatencyIdle       apijson.Field
	LatencyLoaded     apijson.Field
	NumTests          apijson.Field
	RankPower         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *QualitySpeedTopAsesResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopAsesResponseTop0JSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopLocationsResponse struct {
	Meta QualitySpeedTopLocationsResponseMeta   `json:"meta,required"`
	Top0 []QualitySpeedTopLocationsResponseTop0 `json:"top_0,required"`
	JSON qualitySpeedTopLocationsResponseJSON   `json:"-"`
}

// qualitySpeedTopLocationsResponseJSON contains the JSON metadata for the struct
// [QualitySpeedTopLocationsResponse]
type qualitySpeedTopLocationsResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedTopLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopLocationsResponseJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopLocationsResponseMeta struct {
	DateRange      []QualitySpeedTopLocationsResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                             `json:"lastUpdated,required"`
	ConfidenceInfo QualitySpeedTopLocationsResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           qualitySpeedTopLocationsResponseMetaJSON           `json:"-"`
}

// qualitySpeedTopLocationsResponseMetaJSON contains the JSON metadata for the
// struct [QualitySpeedTopLocationsResponseMeta]
type qualitySpeedTopLocationsResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *QualitySpeedTopLocationsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopLocationsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopLocationsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      qualitySpeedTopLocationsResponseMetaDateRangeJSON `json:"-"`
}

// qualitySpeedTopLocationsResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [QualitySpeedTopLocationsResponseMetaDateRange]
type qualitySpeedTopLocationsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedTopLocationsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopLocationsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopLocationsResponseMetaConfidenceInfo struct {
	Annotations []QualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                          `json:"level"`
	JSON        qualitySpeedTopLocationsResponseMetaConfidenceInfoJSON         `json:"-"`
}

// qualitySpeedTopLocationsResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [QualitySpeedTopLocationsResponseMetaConfidenceInfo]
type qualitySpeedTopLocationsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedTopLocationsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopLocationsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                           `json:"dataSource,required"`
	Description     string                                                           `json:"description,required"`
	EventType       string                                                           `json:"eventType,required"`
	IsInstantaneous bool                                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                                        `json:"endTime" format:"date-time"`
	LinkedURL       string                                                           `json:"linkedUrl"`
	StartTime       time.Time                                                        `json:"startTime" format:"date-time"`
	JSON            qualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// qualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [QualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotation]
type qualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *QualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopLocationsResponseTop0 struct {
	BandwidthDownload   string                                   `json:"bandwidthDownload,required"`
	BandwidthUpload     string                                   `json:"bandwidthUpload,required"`
	ClientCountryAlpha2 string                                   `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                   `json:"clientCountryName,required"`
	JitterIdle          string                                   `json:"jitterIdle,required"`
	JitterLoaded        string                                   `json:"jitterLoaded,required"`
	LatencyIdle         string                                   `json:"latencyIdle,required"`
	LatencyLoaded       string                                   `json:"latencyLoaded,required"`
	NumTests            float64                                  `json:"numTests,required"`
	RankPower           float64                                  `json:"rankPower,required"`
	JSON                qualitySpeedTopLocationsResponseTop0JSON `json:"-"`
}

// qualitySpeedTopLocationsResponseTop0JSON contains the JSON metadata for the
// struct [QualitySpeedTopLocationsResponseTop0]
type qualitySpeedTopLocationsResponseTop0JSON struct {
	BandwidthDownload   apijson.Field
	BandwidthUpload     apijson.Field
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	JitterIdle          apijson.Field
	JitterLoaded        apijson.Field
	LatencyIdle         apijson.Field
	LatencyLoaded       apijson.Field
	NumTests            apijson.Field
	RankPower           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *QualitySpeedTopLocationsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopLocationsResponseTop0JSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopAsesParams struct {
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
	Format param.Field[QualitySpeedTopAsesParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Metric to order the results by.
	OrderBy param.Field[QualitySpeedTopAsesParamsOrderBy] `query:"orderBy"`
	// Reverse the order of results.
	Reverse param.Field[bool] `query:"reverse"`
}

// URLQuery serializes [QualitySpeedTopAsesParams]'s query parameters as
// `url.Values`.
func (r QualitySpeedTopAsesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type QualitySpeedTopAsesParamsFormat string

const (
	QualitySpeedTopAsesParamsFormatJson QualitySpeedTopAsesParamsFormat = "JSON"
	QualitySpeedTopAsesParamsFormatCsv  QualitySpeedTopAsesParamsFormat = "CSV"
)

func (r QualitySpeedTopAsesParamsFormat) IsKnown() bool {
	switch r {
	case QualitySpeedTopAsesParamsFormatJson, QualitySpeedTopAsesParamsFormatCsv:
		return true
	}
	return false
}

// Metric to order the results by.
type QualitySpeedTopAsesParamsOrderBy string

const (
	QualitySpeedTopAsesParamsOrderByBandwidthDownload QualitySpeedTopAsesParamsOrderBy = "BANDWIDTH_DOWNLOAD"
	QualitySpeedTopAsesParamsOrderByBandwidthUpload   QualitySpeedTopAsesParamsOrderBy = "BANDWIDTH_UPLOAD"
	QualitySpeedTopAsesParamsOrderByLatencyIdle       QualitySpeedTopAsesParamsOrderBy = "LATENCY_IDLE"
	QualitySpeedTopAsesParamsOrderByLatencyLoaded     QualitySpeedTopAsesParamsOrderBy = "LATENCY_LOADED"
	QualitySpeedTopAsesParamsOrderByJitterIdle        QualitySpeedTopAsesParamsOrderBy = "JITTER_IDLE"
	QualitySpeedTopAsesParamsOrderByJitterLoaded      QualitySpeedTopAsesParamsOrderBy = "JITTER_LOADED"
)

func (r QualitySpeedTopAsesParamsOrderBy) IsKnown() bool {
	switch r {
	case QualitySpeedTopAsesParamsOrderByBandwidthDownload, QualitySpeedTopAsesParamsOrderByBandwidthUpload, QualitySpeedTopAsesParamsOrderByLatencyIdle, QualitySpeedTopAsesParamsOrderByLatencyLoaded, QualitySpeedTopAsesParamsOrderByJitterIdle, QualitySpeedTopAsesParamsOrderByJitterLoaded:
		return true
	}
	return false
}

type QualitySpeedTopAsesResponseEnvelope struct {
	Result  QualitySpeedTopAsesResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    qualitySpeedTopAsesResponseEnvelopeJSON `json:"-"`
}

// qualitySpeedTopAsesResponseEnvelopeJSON contains the JSON metadata for the
// struct [QualitySpeedTopAsesResponseEnvelope]
type qualitySpeedTopAsesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedTopAsesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopAsesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type QualitySpeedTopLocationsParams struct {
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
	Format param.Field[QualitySpeedTopLocationsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Metric to order the results by.
	OrderBy param.Field[QualitySpeedTopLocationsParamsOrderBy] `query:"orderBy"`
	// Reverse the order of results.
	Reverse param.Field[bool] `query:"reverse"`
}

// URLQuery serializes [QualitySpeedTopLocationsParams]'s query parameters as
// `url.Values`.
func (r QualitySpeedTopLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type QualitySpeedTopLocationsParamsFormat string

const (
	QualitySpeedTopLocationsParamsFormatJson QualitySpeedTopLocationsParamsFormat = "JSON"
	QualitySpeedTopLocationsParamsFormatCsv  QualitySpeedTopLocationsParamsFormat = "CSV"
)

func (r QualitySpeedTopLocationsParamsFormat) IsKnown() bool {
	switch r {
	case QualitySpeedTopLocationsParamsFormatJson, QualitySpeedTopLocationsParamsFormatCsv:
		return true
	}
	return false
}

// Metric to order the results by.
type QualitySpeedTopLocationsParamsOrderBy string

const (
	QualitySpeedTopLocationsParamsOrderByBandwidthDownload QualitySpeedTopLocationsParamsOrderBy = "BANDWIDTH_DOWNLOAD"
	QualitySpeedTopLocationsParamsOrderByBandwidthUpload   QualitySpeedTopLocationsParamsOrderBy = "BANDWIDTH_UPLOAD"
	QualitySpeedTopLocationsParamsOrderByLatencyIdle       QualitySpeedTopLocationsParamsOrderBy = "LATENCY_IDLE"
	QualitySpeedTopLocationsParamsOrderByLatencyLoaded     QualitySpeedTopLocationsParamsOrderBy = "LATENCY_LOADED"
	QualitySpeedTopLocationsParamsOrderByJitterIdle        QualitySpeedTopLocationsParamsOrderBy = "JITTER_IDLE"
	QualitySpeedTopLocationsParamsOrderByJitterLoaded      QualitySpeedTopLocationsParamsOrderBy = "JITTER_LOADED"
)

func (r QualitySpeedTopLocationsParamsOrderBy) IsKnown() bool {
	switch r {
	case QualitySpeedTopLocationsParamsOrderByBandwidthDownload, QualitySpeedTopLocationsParamsOrderByBandwidthUpload, QualitySpeedTopLocationsParamsOrderByLatencyIdle, QualitySpeedTopLocationsParamsOrderByLatencyLoaded, QualitySpeedTopLocationsParamsOrderByJitterIdle, QualitySpeedTopLocationsParamsOrderByJitterLoaded:
		return true
	}
	return false
}

type QualitySpeedTopLocationsResponseEnvelope struct {
	Result  QualitySpeedTopLocationsResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    qualitySpeedTopLocationsResponseEnvelopeJSON `json:"-"`
}

// qualitySpeedTopLocationsResponseEnvelopeJSON contains the JSON metadata for the
// struct [QualitySpeedTopLocationsResponseEnvelope]
type qualitySpeedTopLocationsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualitySpeedTopLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualitySpeedTopLocationsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
