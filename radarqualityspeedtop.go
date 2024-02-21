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

// RadarQualitySpeedTopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarQualitySpeedTopService]
// method instead.
type RadarQualitySpeedTopService struct {
	Options []option.RequestOption
}

// NewRadarQualitySpeedTopService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarQualitySpeedTopService(opts ...option.RequestOption) (r *RadarQualitySpeedTopService) {
	r = &RadarQualitySpeedTopService{}
	r.Options = opts
	return
}

// Get the top autonomous systems by bandwidth, latency, jitter or packet loss,
// from the previous 90 days of Cloudflare Speed Test data.
func (r *RadarQualitySpeedTopService) Ases(ctx context.Context, query RadarQualitySpeedTopAsesParams, opts ...option.RequestOption) (res *RadarQualitySpeedTopAsesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarQualitySpeedTopAsesResponseEnvelope
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
func (r *RadarQualitySpeedTopService) Locations(ctx context.Context, query RadarQualitySpeedTopLocationsParams, opts ...option.RequestOption) (res *RadarQualitySpeedTopLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarQualitySpeedTopLocationsResponseEnvelope
	path := "radar/quality/speed/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarQualitySpeedTopAsesResponse struct {
	Meta RadarQualitySpeedTopAsesResponseMeta   `json:"meta,required"`
	Top0 []RadarQualitySpeedTopAsesResponseTop0 `json:"top_0,required"`
	JSON radarQualitySpeedTopAsesResponseJSON   `json:"-"`
}

// radarQualitySpeedTopAsesResponseJSON contains the JSON metadata for the struct
// [RadarQualitySpeedTopAsesResponse]
type radarQualitySpeedTopAsesResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopAsesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopAsesResponseMeta struct {
	DateRange      []RadarQualitySpeedTopAsesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                             `json:"lastUpdated,required"`
	ConfidenceInfo RadarQualitySpeedTopAsesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarQualitySpeedTopAsesResponseMetaJSON           `json:"-"`
}

// radarQualitySpeedTopAsesResponseMetaJSON contains the JSON metadata for the
// struct [RadarQualitySpeedTopAsesResponseMeta]
type radarQualitySpeedTopAsesResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarQualitySpeedTopAsesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopAsesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      radarQualitySpeedTopAsesResponseMetaDateRangeJSON `json:"-"`
}

// radarQualitySpeedTopAsesResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarQualitySpeedTopAsesResponseMetaDateRange]
type radarQualitySpeedTopAsesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopAsesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopAsesResponseMetaConfidenceInfo struct {
	Annotations []RadarQualitySpeedTopAsesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                          `json:"level"`
	JSON        radarQualitySpeedTopAsesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarQualitySpeedTopAsesResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarQualitySpeedTopAsesResponseMetaConfidenceInfo]
type radarQualitySpeedTopAsesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopAsesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopAsesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                           `json:"dataSource,required"`
	Description     string                                                           `json:"description,required"`
	EventType       string                                                           `json:"eventType,required"`
	IsInstantaneous interface{}                                                      `json:"isInstantaneous,required"`
	EndTime         time.Time                                                        `json:"endTime" format:"date-time"`
	LinkedURL       string                                                           `json:"linkedUrl"`
	StartTime       time.Time                                                        `json:"startTime" format:"date-time"`
	JSON            radarQualitySpeedTopAsesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarQualitySpeedTopAsesResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarQualitySpeedTopAsesResponseMetaConfidenceInfoAnnotation]
type radarQualitySpeedTopAsesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarQualitySpeedTopAsesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopAsesResponseTop0 struct {
	BandwidthDownload string                                   `json:"bandwidthDownload,required"`
	BandwidthUpload   string                                   `json:"bandwidthUpload,required"`
	ClientAsn         float64                                  `json:"clientASN,required"`
	ClientAsName      string                                   `json:"clientASName,required"`
	JitterIdle        string                                   `json:"jitterIdle,required"`
	JitterLoaded      string                                   `json:"jitterLoaded,required"`
	LatencyIdle       string                                   `json:"latencyIdle,required"`
	LatencyLoaded     string                                   `json:"latencyLoaded,required"`
	NumTests          float64                                  `json:"numTests,required"`
	RankPower         float64                                  `json:"rankPower,required"`
	JSON              radarQualitySpeedTopAsesResponseTop0JSON `json:"-"`
}

// radarQualitySpeedTopAsesResponseTop0JSON contains the JSON metadata for the
// struct [RadarQualitySpeedTopAsesResponseTop0]
type radarQualitySpeedTopAsesResponseTop0JSON struct {
	BandwidthDownload apijson.Field
	BandwidthUpload   apijson.Field
	ClientAsn         apijson.Field
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

func (r *RadarQualitySpeedTopAsesResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopLocationsResponse struct {
	Meta RadarQualitySpeedTopLocationsResponseMeta   `json:"meta,required"`
	Top0 []RadarQualitySpeedTopLocationsResponseTop0 `json:"top_0,required"`
	JSON radarQualitySpeedTopLocationsResponseJSON   `json:"-"`
}

// radarQualitySpeedTopLocationsResponseJSON contains the JSON metadata for the
// struct [RadarQualitySpeedTopLocationsResponse]
type radarQualitySpeedTopLocationsResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopLocationsResponseMeta struct {
	DateRange      []RadarQualitySpeedTopLocationsResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                  `json:"lastUpdated,required"`
	ConfidenceInfo RadarQualitySpeedTopLocationsResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarQualitySpeedTopLocationsResponseMetaJSON           `json:"-"`
}

// radarQualitySpeedTopLocationsResponseMetaJSON contains the JSON metadata for the
// struct [RadarQualitySpeedTopLocationsResponseMeta]
type radarQualitySpeedTopLocationsResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarQualitySpeedTopLocationsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopLocationsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      radarQualitySpeedTopLocationsResponseMetaDateRangeJSON `json:"-"`
}

// radarQualitySpeedTopLocationsResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarQualitySpeedTopLocationsResponseMetaDateRange]
type radarQualitySpeedTopLocationsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopLocationsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopLocationsResponseMetaConfidenceInfo struct {
	Annotations []RadarQualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                               `json:"level"`
	JSON        radarQualitySpeedTopLocationsResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarQualitySpeedTopLocationsResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarQualitySpeedTopLocationsResponseMetaConfidenceInfo]
type radarQualitySpeedTopLocationsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopLocationsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                `json:"dataSource,required"`
	Description     string                                                                `json:"description,required"`
	EventType       string                                                                `json:"eventType,required"`
	IsInstantaneous interface{}                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                             `json:"startTime" format:"date-time"`
	JSON            radarQualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarQualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarQualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotation]
type radarQualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarQualitySpeedTopLocationsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopLocationsResponseTop0 struct {
	BandwidthDownload   string                                        `json:"bandwidthDownload,required"`
	BandwidthUpload     string                                        `json:"bandwidthUpload,required"`
	ClientCountryAlpha2 string                                        `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                        `json:"clientCountryName,required"`
	JitterIdle          string                                        `json:"jitterIdle,required"`
	JitterLoaded        string                                        `json:"jitterLoaded,required"`
	LatencyIdle         string                                        `json:"latencyIdle,required"`
	LatencyLoaded       string                                        `json:"latencyLoaded,required"`
	NumTests            float64                                       `json:"numTests,required"`
	RankPower           float64                                       `json:"rankPower,required"`
	JSON                radarQualitySpeedTopLocationsResponseTop0JSON `json:"-"`
}

// radarQualitySpeedTopLocationsResponseTop0JSON contains the JSON metadata for the
// struct [RadarQualitySpeedTopLocationsResponseTop0]
type radarQualitySpeedTopLocationsResponseTop0JSON struct {
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

func (r *RadarQualitySpeedTopLocationsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopAsesParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarQualitySpeedTopAsesParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Metric to order the results by.
	OrderBy param.Field[RadarQualitySpeedTopAsesParamsOrderBy] `query:"orderBy"`
	// Reverse the order of results.
	Reverse param.Field[bool] `query:"reverse"`
}

// URLQuery serializes [RadarQualitySpeedTopAsesParams]'s query parameters as
// `url.Values`.
func (r RadarQualitySpeedTopAsesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarQualitySpeedTopAsesParamsFormat string

const (
	RadarQualitySpeedTopAsesParamsFormatJson RadarQualitySpeedTopAsesParamsFormat = "JSON"
	RadarQualitySpeedTopAsesParamsFormatCsv  RadarQualitySpeedTopAsesParamsFormat = "CSV"
)

// Metric to order the results by.
type RadarQualitySpeedTopAsesParamsOrderBy string

const (
	RadarQualitySpeedTopAsesParamsOrderByBandwidthDownload RadarQualitySpeedTopAsesParamsOrderBy = "BANDWIDTH_DOWNLOAD"
	RadarQualitySpeedTopAsesParamsOrderByBandwidthUpload   RadarQualitySpeedTopAsesParamsOrderBy = "BANDWIDTH_UPLOAD"
	RadarQualitySpeedTopAsesParamsOrderByLatencyIdle       RadarQualitySpeedTopAsesParamsOrderBy = "LATENCY_IDLE"
	RadarQualitySpeedTopAsesParamsOrderByLatencyLoaded     RadarQualitySpeedTopAsesParamsOrderBy = "LATENCY_LOADED"
	RadarQualitySpeedTopAsesParamsOrderByJitterIdle        RadarQualitySpeedTopAsesParamsOrderBy = "JITTER_IDLE"
	RadarQualitySpeedTopAsesParamsOrderByJitterLoaded      RadarQualitySpeedTopAsesParamsOrderBy = "JITTER_LOADED"
)

type RadarQualitySpeedTopAsesResponseEnvelope struct {
	Result  RadarQualitySpeedTopAsesResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarQualitySpeedTopAsesResponseEnvelopeJSON `json:"-"`
}

// radarQualitySpeedTopAsesResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarQualitySpeedTopAsesResponseEnvelope]
type radarQualitySpeedTopAsesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopAsesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopLocationsParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarQualitySpeedTopLocationsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Metric to order the results by.
	OrderBy param.Field[RadarQualitySpeedTopLocationsParamsOrderBy] `query:"orderBy"`
	// Reverse the order of results.
	Reverse param.Field[bool] `query:"reverse"`
}

// URLQuery serializes [RadarQualitySpeedTopLocationsParams]'s query parameters as
// `url.Values`.
func (r RadarQualitySpeedTopLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarQualitySpeedTopLocationsParamsFormat string

const (
	RadarQualitySpeedTopLocationsParamsFormatJson RadarQualitySpeedTopLocationsParamsFormat = "JSON"
	RadarQualitySpeedTopLocationsParamsFormatCsv  RadarQualitySpeedTopLocationsParamsFormat = "CSV"
)

// Metric to order the results by.
type RadarQualitySpeedTopLocationsParamsOrderBy string

const (
	RadarQualitySpeedTopLocationsParamsOrderByBandwidthDownload RadarQualitySpeedTopLocationsParamsOrderBy = "BANDWIDTH_DOWNLOAD"
	RadarQualitySpeedTopLocationsParamsOrderByBandwidthUpload   RadarQualitySpeedTopLocationsParamsOrderBy = "BANDWIDTH_UPLOAD"
	RadarQualitySpeedTopLocationsParamsOrderByLatencyIdle       RadarQualitySpeedTopLocationsParamsOrderBy = "LATENCY_IDLE"
	RadarQualitySpeedTopLocationsParamsOrderByLatencyLoaded     RadarQualitySpeedTopLocationsParamsOrderBy = "LATENCY_LOADED"
	RadarQualitySpeedTopLocationsParamsOrderByJitterIdle        RadarQualitySpeedTopLocationsParamsOrderBy = "JITTER_IDLE"
	RadarQualitySpeedTopLocationsParamsOrderByJitterLoaded      RadarQualitySpeedTopLocationsParamsOrderBy = "JITTER_LOADED"
)

type RadarQualitySpeedTopLocationsResponseEnvelope struct {
	Result  RadarQualitySpeedTopLocationsResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarQualitySpeedTopLocationsResponseEnvelopeJSON `json:"-"`
}

// radarQualitySpeedTopLocationsResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarQualitySpeedTopLocationsResponseEnvelope]
type radarQualitySpeedTopLocationsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
