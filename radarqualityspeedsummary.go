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

// RadarQualitySpeedSummaryService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarQualitySpeedSummaryService] method instead.
type RadarQualitySpeedSummaryService struct {
	Options []option.RequestOption
}

// NewRadarQualitySpeedSummaryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarQualitySpeedSummaryService(opts ...option.RequestOption) (r *RadarQualitySpeedSummaryService) {
	r = &RadarQualitySpeedSummaryService{}
	r.Options = opts
	return
}

// Get a summary of bandwidth, latency, jitter and packet loss, from the previous
// 90 days of Cloudflare Speed Test data.
func (r *RadarQualitySpeedSummaryService) Get(ctx context.Context, query RadarQualitySpeedSummaryGetParams, opts ...option.RequestOption) (res *RadarQualitySpeedSummaryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/quality/speed/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarQualitySpeedSummaryGetResponse struct {
	Result  RadarQualitySpeedSummaryGetResponseResult `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarQualitySpeedSummaryGetResponseJSON   `json:"-"`
}

// radarQualitySpeedSummaryGetResponseJSON contains the JSON metadata for the
// struct [RadarQualitySpeedSummaryGetResponse]
type radarQualitySpeedSummaryGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedSummaryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryGetResponseResult struct {
	Meta     RadarQualitySpeedSummaryGetResponseResultMeta     `json:"meta,required"`
	Summary0 RadarQualitySpeedSummaryGetResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarQualitySpeedSummaryGetResponseResultJSON     `json:"-"`
}

// radarQualitySpeedSummaryGetResponseResultJSON contains the JSON metadata for the
// struct [RadarQualitySpeedSummaryGetResponseResult]
type radarQualitySpeedSummaryGetResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedSummaryGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryGetResponseResultMeta struct {
	DateRange      []RadarQualitySpeedSummaryGetResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                      `json:"lastUpdated,required"`
	Normalization  string                                                      `json:"normalization,required"`
	ConfidenceInfo RadarQualitySpeedSummaryGetResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarQualitySpeedSummaryGetResponseResultMetaJSON           `json:"-"`
}

// radarQualitySpeedSummaryGetResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarQualitySpeedSummaryGetResponseResultMeta]
type radarQualitySpeedSummaryGetResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarQualitySpeedSummaryGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryGetResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarQualitySpeedSummaryGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarQualitySpeedSummaryGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarQualitySpeedSummaryGetResponseResultMetaDateRange]
type radarQualitySpeedSummaryGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedSummaryGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarQualitySpeedSummaryGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                   `json:"level"`
	JSON        radarQualitySpeedSummaryGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarQualitySpeedSummaryGetResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarQualitySpeedSummaryGetResponseResultMetaConfidenceInfo]
type radarQualitySpeedSummaryGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedSummaryGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                    `json:"dataSource,required"`
	Description     string                                                                    `json:"description,required"`
	EventType       string                                                                    `json:"eventType,required"`
	IsInstantaneous interface{}                                                               `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                                 `json:"startTime" format:"date-time"`
	JSON            radarQualitySpeedSummaryGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarQualitySpeedSummaryGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarQualitySpeedSummaryGetResponseResultMetaConfidenceInfoAnnotation]
type radarQualitySpeedSummaryGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarQualitySpeedSummaryGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryGetResponseResultSummary0 struct {
	BandwidthDownload string                                                `json:"bandwidthDownload,required"`
	BandwidthUpload   string                                                `json:"bandwidthUpload,required"`
	JitterIdle        string                                                `json:"jitterIdle,required"`
	JitterLoaded      string                                                `json:"jitterLoaded,required"`
	LatencyIdle       string                                                `json:"latencyIdle,required"`
	LatencyLoaded     string                                                `json:"latencyLoaded,required"`
	PacketLoss        string                                                `json:"packetLoss,required"`
	JSON              radarQualitySpeedSummaryGetResponseResultSummary0JSON `json:"-"`
}

// radarQualitySpeedSummaryGetResponseResultSummary0JSON contains the JSON metadata
// for the struct [RadarQualitySpeedSummaryGetResponseResultSummary0]
type radarQualitySpeedSummaryGetResponseResultSummary0JSON struct {
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

func (r *RadarQualitySpeedSummaryGetResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarQualitySpeedSummaryGetParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarQualitySpeedSummaryGetParams]'s query parameters as
// `url.Values`.
func (r RadarQualitySpeedSummaryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarQualitySpeedSummaryGetParamsFormat string

const (
	RadarQualitySpeedSummaryGetParamsFormatJson RadarQualitySpeedSummaryGetParamsFormat = "JSON"
	RadarQualitySpeedSummaryGetParamsFormatCsv  RadarQualitySpeedSummaryGetParamsFormat = "CSV"
)
