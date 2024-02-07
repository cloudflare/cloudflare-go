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

// RadarQualitySpeedTopLocationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarQualitySpeedTopLocationService] method instead.
type RadarQualitySpeedTopLocationService struct {
	Options []option.RequestOption
}

// NewRadarQualitySpeedTopLocationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarQualitySpeedTopLocationService(opts ...option.RequestOption) (r *RadarQualitySpeedTopLocationService) {
	r = &RadarQualitySpeedTopLocationService{}
	r.Options = opts
	return
}

// Get the top locations by bandwidth, latency, jitter or packet loss, from the
// previous 90 days of Cloudflare Speed Test data.
func (r *RadarQualitySpeedTopLocationService) List(ctx context.Context, query RadarQualitySpeedTopLocationListParams, opts ...option.RequestOption) (res *RadarQualitySpeedTopLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/quality/speed/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarQualitySpeedTopLocationListResponse struct {
	Result  RadarQualitySpeedTopLocationListResponseResult `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarQualitySpeedTopLocationListResponseJSON   `json:"-"`
}

// radarQualitySpeedTopLocationListResponseJSON contains the JSON metadata for the
// struct [RadarQualitySpeedTopLocationListResponse]
type radarQualitySpeedTopLocationListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopLocationListResponseResult struct {
	Meta RadarQualitySpeedTopLocationListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarQualitySpeedTopLocationListResponseResultTop0 `json:"top_0,required"`
	JSON radarQualitySpeedTopLocationListResponseResultJSON   `json:"-"`
}

// radarQualitySpeedTopLocationListResponseResultJSON contains the JSON metadata
// for the struct [RadarQualitySpeedTopLocationListResponseResult]
type radarQualitySpeedTopLocationListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopLocationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopLocationListResponseResultMeta struct {
	DateRange      []RadarQualitySpeedTopLocationListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                           `json:"lastUpdated,required"`
	ConfidenceInfo RadarQualitySpeedTopLocationListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarQualitySpeedTopLocationListResponseResultMetaJSON           `json:"-"`
}

// radarQualitySpeedTopLocationListResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarQualitySpeedTopLocationListResponseResultMeta]
type radarQualitySpeedTopLocationListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarQualitySpeedTopLocationListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopLocationListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      radarQualitySpeedTopLocationListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarQualitySpeedTopLocationListResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarQualitySpeedTopLocationListResponseResultMetaDateRange]
type radarQualitySpeedTopLocationListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopLocationListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopLocationListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarQualitySpeedTopLocationListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                        `json:"level"`
	JSON        radarQualitySpeedTopLocationListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarQualitySpeedTopLocationListResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarQualitySpeedTopLocationListResponseResultMetaConfidenceInfo]
type radarQualitySpeedTopLocationListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopLocationListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopLocationListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                         `json:"dataSource,required"`
	Description     string                                                                         `json:"description,required"`
	EventType       string                                                                         `json:"eventType,required"`
	IsInstantaneous interface{}                                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                                      `json:"startTime" format:"date-time"`
	JSON            radarQualitySpeedTopLocationListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarQualitySpeedTopLocationListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarQualitySpeedTopLocationListResponseResultMetaConfidenceInfoAnnotation]
type radarQualitySpeedTopLocationListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarQualitySpeedTopLocationListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopLocationListResponseResultTop0 struct {
	BandwidthDownload   string                                                 `json:"bandwidthDownload,required"`
	BandwidthUpload     string                                                 `json:"bandwidthUpload,required"`
	ClientCountryAlpha2 string                                                 `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                 `json:"clientCountryName,required"`
	JitterIdle          string                                                 `json:"jitterIdle,required"`
	JitterLoaded        string                                                 `json:"jitterLoaded,required"`
	LatencyIdle         string                                                 `json:"latencyIdle,required"`
	LatencyLoaded       string                                                 `json:"latencyLoaded,required"`
	NumTests            float64                                                `json:"numTests,required"`
	RankPower           float64                                                `json:"rankPower,required"`
	JSON                radarQualitySpeedTopLocationListResponseResultTop0JSON `json:"-"`
}

// radarQualitySpeedTopLocationListResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarQualitySpeedTopLocationListResponseResultTop0]
type radarQualitySpeedTopLocationListResponseResultTop0JSON struct {
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

func (r *RadarQualitySpeedTopLocationListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedTopLocationListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarQualitySpeedTopLocationListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Metric to order the results by.
	OrderBy param.Field[RadarQualitySpeedTopLocationListParamsOrderBy] `query:"orderBy"`
	// Reverse the order of results.
	Reverse param.Field[bool] `query:"reverse"`
}

// URLQuery serializes [RadarQualitySpeedTopLocationListParams]'s query parameters
// as `url.Values`.
func (r RadarQualitySpeedTopLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarQualitySpeedTopLocationListParamsFormat string

const (
	RadarQualitySpeedTopLocationListParamsFormatJson RadarQualitySpeedTopLocationListParamsFormat = "JSON"
	RadarQualitySpeedTopLocationListParamsFormatCsv  RadarQualitySpeedTopLocationListParamsFormat = "CSV"
)

// Metric to order the results by.
type RadarQualitySpeedTopLocationListParamsOrderBy string

const (
	RadarQualitySpeedTopLocationListParamsOrderByBandwidthDownload RadarQualitySpeedTopLocationListParamsOrderBy = "BANDWIDTH_DOWNLOAD"
	RadarQualitySpeedTopLocationListParamsOrderByBandwidthUpload   RadarQualitySpeedTopLocationListParamsOrderBy = "BANDWIDTH_UPLOAD"
	RadarQualitySpeedTopLocationListParamsOrderByLatencyIdle       RadarQualitySpeedTopLocationListParamsOrderBy = "LATENCY_IDLE"
	RadarQualitySpeedTopLocationListParamsOrderByLatencyLoaded     RadarQualitySpeedTopLocationListParamsOrderBy = "LATENCY_LOADED"
	RadarQualitySpeedTopLocationListParamsOrderByJitterIdle        RadarQualitySpeedTopLocationListParamsOrderBy = "JITTER_IDLE"
	RadarQualitySpeedTopLocationListParamsOrderByJitterLoaded      RadarQualitySpeedTopLocationListParamsOrderBy = "JITTER_LOADED"
)
