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

// RadarQualitySpeedTopAseService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarQualitySpeedTopAseService] method instead.
type RadarQualitySpeedTopAseService struct {
	Options []option.RequestOption
}

// NewRadarQualitySpeedTopAseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarQualitySpeedTopAseService(opts ...option.RequestOption) (r *RadarQualitySpeedTopAseService) {
	r = &RadarQualitySpeedTopAseService{}
	r.Options = opts
	return
}

// Get the top autonomous systems by bandwidth, latency, jitter or packet loss,
// from the previous 90 days of Cloudflare Speed Test data.
func (r *RadarQualitySpeedTopAseService) List(ctx context.Context, query RadarQualitySpeedTopAseListParams, opts ...option.RequestOption) (res *RadarQualitySpeedTopAseListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarQualitySpeedTopAseListResponseEnvelope
	path := "radar/quality/speed/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarQualitySpeedTopAseListResponse struct {
	Meta RadarQualitySpeedTopAseListResponseMeta   `json:"meta,required"`
	Top0 []RadarQualitySpeedTopAseListResponseTop0 `json:"top_0,required"`
	JSON radarQualitySpeedTopAseListResponseJSON   `json:"-"`
}

// radarQualitySpeedTopAseListResponseJSON contains the JSON metadata for the
// struct [RadarQualitySpeedTopAseListResponse]
type radarQualitySpeedTopAseListResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopAseListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopAseListResponseMeta struct {
	DateRange      []RadarQualitySpeedTopAseListResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                `json:"lastUpdated,required"`
	ConfidenceInfo RadarQualitySpeedTopAseListResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarQualitySpeedTopAseListResponseMetaJSON           `json:"-"`
}

// radarQualitySpeedTopAseListResponseMetaJSON contains the JSON metadata for the
// struct [RadarQualitySpeedTopAseListResponseMeta]
type radarQualitySpeedTopAseListResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarQualitySpeedTopAseListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopAseListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      radarQualitySpeedTopAseListResponseMetaDateRangeJSON `json:"-"`
}

// radarQualitySpeedTopAseListResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarQualitySpeedTopAseListResponseMetaDateRange]
type radarQualitySpeedTopAseListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopAseListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopAseListResponseMetaConfidenceInfo struct {
	Annotations []RadarQualitySpeedTopAseListResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                             `json:"level"`
	JSON        radarQualitySpeedTopAseListResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarQualitySpeedTopAseListResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarQualitySpeedTopAseListResponseMetaConfidenceInfo]
type radarQualitySpeedTopAseListResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopAseListResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopAseListResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                              `json:"dataSource,required"`
	Description     string                                                              `json:"description,required"`
	EventType       string                                                              `json:"eventType,required"`
	IsInstantaneous interface{}                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                           `json:"startTime" format:"date-time"`
	JSON            radarQualitySpeedTopAseListResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarQualitySpeedTopAseListResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarQualitySpeedTopAseListResponseMetaConfidenceInfoAnnotation]
type radarQualitySpeedTopAseListResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarQualitySpeedTopAseListResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopAseListResponseTop0 struct {
	BandwidthDownload string                                      `json:"bandwidthDownload,required"`
	BandwidthUpload   string                                      `json:"bandwidthUpload,required"`
	ClientAsn         float64                                     `json:"clientASN,required"`
	ClientAsName      string                                      `json:"clientASName,required"`
	JitterIdle        string                                      `json:"jitterIdle,required"`
	JitterLoaded      string                                      `json:"jitterLoaded,required"`
	LatencyIdle       string                                      `json:"latencyIdle,required"`
	LatencyLoaded     string                                      `json:"latencyLoaded,required"`
	NumTests          float64                                     `json:"numTests,required"`
	RankPower         float64                                     `json:"rankPower,required"`
	JSON              radarQualitySpeedTopAseListResponseTop0JSON `json:"-"`
}

// radarQualitySpeedTopAseListResponseTop0JSON contains the JSON metadata for the
// struct [RadarQualitySpeedTopAseListResponseTop0]
type radarQualitySpeedTopAseListResponseTop0JSON struct {
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

func (r *RadarQualitySpeedTopAseListResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopAseListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarQualitySpeedTopAseListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Metric to order the results by.
	OrderBy param.Field[RadarQualitySpeedTopAseListParamsOrderBy] `query:"orderBy"`
	// Reverse the order of results.
	Reverse param.Field[bool] `query:"reverse"`
}

// URLQuery serializes [RadarQualitySpeedTopAseListParams]'s query parameters as
// `url.Values`.
func (r RadarQualitySpeedTopAseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarQualitySpeedTopAseListParamsFormat string

const (
	RadarQualitySpeedTopAseListParamsFormatJson RadarQualitySpeedTopAseListParamsFormat = "JSON"
	RadarQualitySpeedTopAseListParamsFormatCsv  RadarQualitySpeedTopAseListParamsFormat = "CSV"
)

// Metric to order the results by.
type RadarQualitySpeedTopAseListParamsOrderBy string

const (
	RadarQualitySpeedTopAseListParamsOrderByBandwidthDownload RadarQualitySpeedTopAseListParamsOrderBy = "BANDWIDTH_DOWNLOAD"
	RadarQualitySpeedTopAseListParamsOrderByBandwidthUpload   RadarQualitySpeedTopAseListParamsOrderBy = "BANDWIDTH_UPLOAD"
	RadarQualitySpeedTopAseListParamsOrderByLatencyIdle       RadarQualitySpeedTopAseListParamsOrderBy = "LATENCY_IDLE"
	RadarQualitySpeedTopAseListParamsOrderByLatencyLoaded     RadarQualitySpeedTopAseListParamsOrderBy = "LATENCY_LOADED"
	RadarQualitySpeedTopAseListParamsOrderByJitterIdle        RadarQualitySpeedTopAseListParamsOrderBy = "JITTER_IDLE"
	RadarQualitySpeedTopAseListParamsOrderByJitterLoaded      RadarQualitySpeedTopAseListParamsOrderBy = "JITTER_LOADED"
)

type RadarQualitySpeedTopAseListResponseEnvelope struct {
	Result  RadarQualitySpeedTopAseListResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarQualitySpeedTopAseListResponseEnvelopeJSON `json:"-"`
}

// radarQualitySpeedTopAseListResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarQualitySpeedTopAseListResponseEnvelope]
type radarQualitySpeedTopAseListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopAseListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
