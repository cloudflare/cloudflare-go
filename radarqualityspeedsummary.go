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
func (r *RadarQualitySpeedSummaryService) List(ctx context.Context, query RadarQualitySpeedSummaryListParams, opts ...option.RequestOption) (res *RadarQualitySpeedSummaryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarQualitySpeedSummaryListResponseEnvelope
	path := "radar/quality/speed/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarQualitySpeedSummaryListResponse struct {
	Meta     RadarQualitySpeedSummaryListResponseMeta     `json:"meta,required"`
	Summary0 RadarQualitySpeedSummaryListResponseSummary0 `json:"summary_0,required"`
	JSON     radarQualitySpeedSummaryListResponseJSON     `json:"-"`
}

// radarQualitySpeedSummaryListResponseJSON contains the JSON metadata for the
// struct [RadarQualitySpeedSummaryListResponse]
type radarQualitySpeedSummaryListResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedSummaryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryListResponseMeta struct {
	DateRange      []RadarQualitySpeedSummaryListResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                 `json:"lastUpdated,required"`
	Normalization  string                                                 `json:"normalization,required"`
	ConfidenceInfo RadarQualitySpeedSummaryListResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarQualitySpeedSummaryListResponseMetaJSON           `json:"-"`
}

// radarQualitySpeedSummaryListResponseMetaJSON contains the JSON metadata for the
// struct [RadarQualitySpeedSummaryListResponseMeta]
type radarQualitySpeedSummaryListResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarQualitySpeedSummaryListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      radarQualitySpeedSummaryListResponseMetaDateRangeJSON `json:"-"`
}

// radarQualitySpeedSummaryListResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarQualitySpeedSummaryListResponseMetaDateRange]
type radarQualitySpeedSummaryListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedSummaryListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryListResponseMetaConfidenceInfo struct {
	Annotations []RadarQualitySpeedSummaryListResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        radarQualitySpeedSummaryListResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarQualitySpeedSummaryListResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarQualitySpeedSummaryListResponseMetaConfidenceInfo]
type radarQualitySpeedSummaryListResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedSummaryListResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryListResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            radarQualitySpeedSummaryListResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarQualitySpeedSummaryListResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarQualitySpeedSummaryListResponseMetaConfidenceInfoAnnotation]
type radarQualitySpeedSummaryListResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarQualitySpeedSummaryListResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryListResponseSummary0 struct {
	BandwidthDownload string                                           `json:"bandwidthDownload,required"`
	BandwidthUpload   string                                           `json:"bandwidthUpload,required"`
	JitterIdle        string                                           `json:"jitterIdle,required"`
	JitterLoaded      string                                           `json:"jitterLoaded,required"`
	LatencyIdle       string                                           `json:"latencyIdle,required"`
	LatencyLoaded     string                                           `json:"latencyLoaded,required"`
	PacketLoss        string                                           `json:"packetLoss,required"`
	JSON              radarQualitySpeedSummaryListResponseSummary0JSON `json:"-"`
}

// radarQualitySpeedSummaryListResponseSummary0JSON contains the JSON metadata for
// the struct [RadarQualitySpeedSummaryListResponseSummary0]
type radarQualitySpeedSummaryListResponseSummary0JSON struct {
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

func (r *RadarQualitySpeedSummaryListResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualitySpeedSummaryListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarQualitySpeedSummaryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarQualitySpeedSummaryListParams]'s query parameters as
// `url.Values`.
func (r RadarQualitySpeedSummaryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarQualitySpeedSummaryListParamsFormat string

const (
	RadarQualitySpeedSummaryListParamsFormatJson RadarQualitySpeedSummaryListParamsFormat = "JSON"
	RadarQualitySpeedSummaryListParamsFormatCsv  RadarQualitySpeedSummaryListParamsFormat = "CSV"
)

type RadarQualitySpeedSummaryListResponseEnvelope struct {
	Result  RadarQualitySpeedSummaryListResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarQualitySpeedSummaryListResponseEnvelopeJSON `json:"-"`
}

// radarQualitySpeedSummaryListResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarQualitySpeedSummaryListResponseEnvelope]
type radarQualitySpeedSummaryListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedSummaryListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
