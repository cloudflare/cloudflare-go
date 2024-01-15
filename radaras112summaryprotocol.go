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

// RadarAs112SummaryProtocolService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAs112SummaryProtocolService] method instead.
type RadarAs112SummaryProtocolService struct {
	Options []option.RequestOption
}

// NewRadarAs112SummaryProtocolService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAs112SummaryProtocolService(opts ...option.RequestOption) (r *RadarAs112SummaryProtocolService) {
	r = &RadarAs112SummaryProtocolService{}
	r.Options = opts
	return
}

// Percentage distribution of DNS queries to AS112 per protocol.
func (r *RadarAs112SummaryProtocolService) List(ctx context.Context, query RadarAs112SummaryProtocolListParams, opts ...option.RequestOption) (res *RadarAs112SummaryProtocolListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112SummaryProtocolListResponse struct {
	Result  RadarAs112SummaryProtocolListResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarAs112SummaryProtocolListResponseJSON   `json:"-"`
}

// radarAs112SummaryProtocolListResponseJSON contains the JSON metadata for the
// struct [RadarAs112SummaryProtocolListResponse]
type radarAs112SummaryProtocolListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryProtocolListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryProtocolListResponseResult struct {
	Meta     RadarAs112SummaryProtocolListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAs112SummaryProtocolListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAs112SummaryProtocolListResponseResultJSON     `json:"-"`
}

// radarAs112SummaryProtocolListResponseResultJSON contains the JSON metadata for
// the struct [RadarAs112SummaryProtocolListResponseResult]
type radarAs112SummaryProtocolListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryProtocolListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryProtocolListResponseResultMeta struct {
	DateRange      []RadarAs112SummaryProtocolListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                        `json:"lastUpdated,required"`
	Normalization  string                                                        `json:"normalization,required"`
	ConfidenceInfo RadarAs112SummaryProtocolListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112SummaryProtocolListResponseResultMetaJSON           `json:"-"`
}

// radarAs112SummaryProtocolListResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAs112SummaryProtocolListResponseResultMeta]
type radarAs112SummaryProtocolListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112SummaryProtocolListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryProtocolListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarAs112SummaryProtocolListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112SummaryProtocolListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAs112SummaryProtocolListResponseResultMetaDateRange]
type radarAs112SummaryProtocolListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryProtocolListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryProtocolListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112SummaryProtocolListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                     `json:"level"`
	JSON        radarAs112SummaryProtocolListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112SummaryProtocolListResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAs112SummaryProtocolListResponseResultMetaConfidenceInfo]
type radarAs112SummaryProtocolListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryProtocolListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryProtocolListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                      `json:"dataSource,required"`
	Description     string                                                                      `json:"description,required"`
	EventType       string                                                                      `json:"eventType,required"`
	IsInstantaneous interface{}                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                                   `json:"startTime" format:"date-time"`
	JSON            radarAs112SummaryProtocolListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112SummaryProtocolListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAs112SummaryProtocolListResponseResultMetaConfidenceInfoAnnotation]
type radarAs112SummaryProtocolListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112SummaryProtocolListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryProtocolListResponseResultSummary0 struct {
	Tcp  string                                                  `json:"tcp,required"`
	Udp  string                                                  `json:"udp,required"`
	JSON radarAs112SummaryProtocolListResponseResultSummary0JSON `json:"-"`
}

// radarAs112SummaryProtocolListResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarAs112SummaryProtocolListResponseResultSummary0]
type radarAs112SummaryProtocolListResponseResultSummary0JSON struct {
	Tcp         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryProtocolListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112SummaryProtocolListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112SummaryProtocolListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112SummaryProtocolListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112SummaryProtocolListParams]'s query parameters as
// `url.Values`.
func (r RadarAs112SummaryProtocolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAs112SummaryProtocolListParamsDateRange string

const (
	RadarAs112SummaryProtocolListParamsDateRange1d         RadarAs112SummaryProtocolListParamsDateRange = "1d"
	RadarAs112SummaryProtocolListParamsDateRange2d         RadarAs112SummaryProtocolListParamsDateRange = "2d"
	RadarAs112SummaryProtocolListParamsDateRange7d         RadarAs112SummaryProtocolListParamsDateRange = "7d"
	RadarAs112SummaryProtocolListParamsDateRange14d        RadarAs112SummaryProtocolListParamsDateRange = "14d"
	RadarAs112SummaryProtocolListParamsDateRange28d        RadarAs112SummaryProtocolListParamsDateRange = "28d"
	RadarAs112SummaryProtocolListParamsDateRange12w        RadarAs112SummaryProtocolListParamsDateRange = "12w"
	RadarAs112SummaryProtocolListParamsDateRange24w        RadarAs112SummaryProtocolListParamsDateRange = "24w"
	RadarAs112SummaryProtocolListParamsDateRange52w        RadarAs112SummaryProtocolListParamsDateRange = "52w"
	RadarAs112SummaryProtocolListParamsDateRange1dControl  RadarAs112SummaryProtocolListParamsDateRange = "1dControl"
	RadarAs112SummaryProtocolListParamsDateRange2dControl  RadarAs112SummaryProtocolListParamsDateRange = "2dControl"
	RadarAs112SummaryProtocolListParamsDateRange7dControl  RadarAs112SummaryProtocolListParamsDateRange = "7dControl"
	RadarAs112SummaryProtocolListParamsDateRange14dControl RadarAs112SummaryProtocolListParamsDateRange = "14dControl"
	RadarAs112SummaryProtocolListParamsDateRange28dControl RadarAs112SummaryProtocolListParamsDateRange = "28dControl"
	RadarAs112SummaryProtocolListParamsDateRange12wControl RadarAs112SummaryProtocolListParamsDateRange = "12wControl"
	RadarAs112SummaryProtocolListParamsDateRange24wControl RadarAs112SummaryProtocolListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112SummaryProtocolListParamsFormat string

const (
	RadarAs112SummaryProtocolListParamsFormatJson RadarAs112SummaryProtocolListParamsFormat = "JSON"
	RadarAs112SummaryProtocolListParamsFormatCsv  RadarAs112SummaryProtocolListParamsFormat = "CSV"
)
