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

// RadarAttackLayer3SummaryService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAttackLayer3SummaryService] method instead.
type RadarAttackLayer3SummaryService struct {
	Options    []option.RequestOption
	Bitrate    *RadarAttackLayer3SummaryBitrateService
	Durations  *RadarAttackLayer3SummaryDurationService
	IPVersions *RadarAttackLayer3SummaryIPVersionService
	Protocols  *RadarAttackLayer3SummaryProtocolService
	Vectors    *RadarAttackLayer3SummaryVectorService
}

// NewRadarAttackLayer3SummaryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3SummaryService(opts ...option.RequestOption) (r *RadarAttackLayer3SummaryService) {
	r = &RadarAttackLayer3SummaryService{}
	r.Options = opts
	r.Bitrate = NewRadarAttackLayer3SummaryBitrateService(opts...)
	r.Durations = NewRadarAttackLayer3SummaryDurationService(opts...)
	r.IPVersions = NewRadarAttackLayer3SummaryIPVersionService(opts...)
	r.Protocols = NewRadarAttackLayer3SummaryProtocolService(opts...)
	r.Vectors = NewRadarAttackLayer3SummaryVectorService(opts...)
	return
}

// Percentage distribution of network protocols in layer 3/4 attacks over a given
// time period.
func (r *RadarAttackLayer3SummaryService) List(ctx context.Context, query RadarAttackLayer3SummaryListParams, opts ...option.RequestOption) (res *RadarAttackLayer3SummaryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3SummaryListResponse struct {
	Result  RadarAttackLayer3SummaryListResponseResult `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarAttackLayer3SummaryListResponseJSON   `json:"-"`
}

// radarAttackLayer3SummaryListResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3SummaryListResponse]
type radarAttackLayer3SummaryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryListResponseResult struct {
	Meta     RadarAttackLayer3SummaryListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer3SummaryListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer3SummaryListResponseResultJSON     `json:"-"`
}

// radarAttackLayer3SummaryListResponseResultJSON contains the JSON metadata for
// the struct [RadarAttackLayer3SummaryListResponseResult]
type radarAttackLayer3SummaryListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryListResponseResultMeta struct {
	DateRange      []RadarAttackLayer3SummaryListResponseResultMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarAttackLayer3SummaryListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3SummaryListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3SummaryListResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAttackLayer3SummaryListResponseResultMeta]
type radarAttackLayer3SummaryListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3SummaryListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3SummaryListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3SummaryListResponseResultMetaDateRange]
type radarAttackLayer3SummaryListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                    `json:"level"`
	JSON        radarAttackLayer3SummaryListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3SummaryListResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryListResponseResultMetaConfidenceInfo]
type radarAttackLayer3SummaryListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                     `json:"dataSource,required"`
	Description     string                                                                     `json:"description,required"`
	EventType       string                                                                     `json:"eventType,required"`
	IsInstantaneous interface{}                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                                  `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3SummaryListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryListResponseResultSummary0 struct {
	Gre  string                                                 `json:"gre,required"`
	Icmp string                                                 `json:"icmp,required"`
	Tcp  string                                                 `json:"tcp,required"`
	Udp  string                                                 `json:"udp,required"`
	JSON radarAttackLayer3SummaryListResponseResultSummary0JSON `json:"-"`
}

// radarAttackLayer3SummaryListResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryListResponseResultSummary0]
type radarAttackLayer3SummaryListResponseResultSummary0JSON struct {
	Gre         apijson.Field
	Icmp        apijson.Field
	Tcp         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3SummaryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3SummaryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer3SummaryListParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer3SummaryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3SummaryListParamsDateRange string

const (
	RadarAttackLayer3SummaryListParamsDateRange1d         RadarAttackLayer3SummaryListParamsDateRange = "1d"
	RadarAttackLayer3SummaryListParamsDateRange2d         RadarAttackLayer3SummaryListParamsDateRange = "2d"
	RadarAttackLayer3SummaryListParamsDateRange7d         RadarAttackLayer3SummaryListParamsDateRange = "7d"
	RadarAttackLayer3SummaryListParamsDateRange14d        RadarAttackLayer3SummaryListParamsDateRange = "14d"
	RadarAttackLayer3SummaryListParamsDateRange28d        RadarAttackLayer3SummaryListParamsDateRange = "28d"
	RadarAttackLayer3SummaryListParamsDateRange12w        RadarAttackLayer3SummaryListParamsDateRange = "12w"
	RadarAttackLayer3SummaryListParamsDateRange24w        RadarAttackLayer3SummaryListParamsDateRange = "24w"
	RadarAttackLayer3SummaryListParamsDateRange52w        RadarAttackLayer3SummaryListParamsDateRange = "52w"
	RadarAttackLayer3SummaryListParamsDateRange1dControl  RadarAttackLayer3SummaryListParamsDateRange = "1dControl"
	RadarAttackLayer3SummaryListParamsDateRange2dControl  RadarAttackLayer3SummaryListParamsDateRange = "2dControl"
	RadarAttackLayer3SummaryListParamsDateRange7dControl  RadarAttackLayer3SummaryListParamsDateRange = "7dControl"
	RadarAttackLayer3SummaryListParamsDateRange14dControl RadarAttackLayer3SummaryListParamsDateRange = "14dControl"
	RadarAttackLayer3SummaryListParamsDateRange28dControl RadarAttackLayer3SummaryListParamsDateRange = "28dControl"
	RadarAttackLayer3SummaryListParamsDateRange12wControl RadarAttackLayer3SummaryListParamsDateRange = "12wControl"
	RadarAttackLayer3SummaryListParamsDateRange24wControl RadarAttackLayer3SummaryListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3SummaryListParamsFormat string

const (
	RadarAttackLayer3SummaryListParamsFormatJson RadarAttackLayer3SummaryListParamsFormat = "JSON"
	RadarAttackLayer3SummaryListParamsFormatCsv  RadarAttackLayer3SummaryListParamsFormat = "CSV"
)
