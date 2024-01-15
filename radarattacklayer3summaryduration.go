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

// RadarAttackLayer3SummaryDurationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer3SummaryDurationService] method instead.
type RadarAttackLayer3SummaryDurationService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3SummaryDurationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3SummaryDurationService(opts ...option.RequestOption) (r *RadarAttackLayer3SummaryDurationService) {
	r = &RadarAttackLayer3SummaryDurationService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by duration.
func (r *RadarAttackLayer3SummaryDurationService) List(ctx context.Context, query RadarAttackLayer3SummaryDurationListParams, opts ...option.RequestOption) (res *RadarAttackLayer3SummaryDurationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/duration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3SummaryDurationListResponse struct {
	Result  RadarAttackLayer3SummaryDurationListResponseResult `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarAttackLayer3SummaryDurationListResponseJSON   `json:"-"`
}

// radarAttackLayer3SummaryDurationListResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer3SummaryDurationListResponse]
type radarAttackLayer3SummaryDurationListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryDurationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryDurationListResponseResult struct {
	Meta     RadarAttackLayer3SummaryDurationListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer3SummaryDurationListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer3SummaryDurationListResponseResultJSON     `json:"-"`
}

// radarAttackLayer3SummaryDurationListResponseResultJSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryDurationListResponseResult]
type radarAttackLayer3SummaryDurationListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryDurationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryDurationListResponseResultMeta struct {
	DateRange      []RadarAttackLayer3SummaryDurationListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                               `json:"lastUpdated,required"`
	Normalization  string                                                               `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer3SummaryDurationListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3SummaryDurationListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3SummaryDurationListResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryDurationListResponseResultMeta]
type radarAttackLayer3SummaryDurationListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryDurationListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryDurationListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                           `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3SummaryDurationListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3SummaryDurationListResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryDurationListResponseResultMetaDateRange]
type radarAttackLayer3SummaryDurationListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryDurationListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryDurationListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryDurationListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                            `json:"level"`
	JSON        radarAttackLayer3SummaryDurationListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3SummaryDurationListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryDurationListResponseResultMetaConfidenceInfo]
type radarAttackLayer3SummaryDurationListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryDurationListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryDurationListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                             `json:"dataSource,required"`
	Description     string                                                                             `json:"description,required"`
	EventType       string                                                                             `json:"eventType,required"`
	IsInstantaneous interface{}                                                                        `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                          `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                             `json:"linkedUrl"`
	StartTime       time.Time                                                                          `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3SummaryDurationListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryDurationListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryDurationListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryDurationListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryDurationListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryDurationListResponseResultSummary0 struct {
	_1HourTo3Hours  string                                                         `json:"_1_HOUR_TO_3_HOURS,required"`
	_10MinsTo20Mins string                                                         `json:"_10_MINS_TO_20_MINS,required"`
	_20MinsTo40Mins string                                                         `json:"_20_MINS_TO_40_MINS,required"`
	_40MinsTo1Hour  string                                                         `json:"_40_MINS_TO_1_HOUR,required"`
	Over3Hours      string                                                         `json:"OVER_3_HOURS,required"`
	Under10Mins     string                                                         `json:"UNDER_10_MINS,required"`
	JSON            radarAttackLayer3SummaryDurationListResponseResultSummary0JSON `json:"-"`
}

// radarAttackLayer3SummaryDurationListResponseResultSummary0JSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3SummaryDurationListResponseResultSummary0]
type radarAttackLayer3SummaryDurationListResponseResultSummary0JSON struct {
	_1HourTo3Hours  apijson.Field
	_10MinsTo20Mins apijson.Field
	_20MinsTo40Mins apijson.Field
	_40MinsTo1Hour  apijson.Field
	Over3Hours      apijson.Field
	Under10Mins     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryDurationListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryDurationListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3SummaryDurationListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3SummaryDurationListParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3SummaryDurationListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3SummaryDurationListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3SummaryDurationListParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3SummaryDurationListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3SummaryDurationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3SummaryDurationListParamsDateRange string

const (
	RadarAttackLayer3SummaryDurationListParamsDateRange1d         RadarAttackLayer3SummaryDurationListParamsDateRange = "1d"
	RadarAttackLayer3SummaryDurationListParamsDateRange2d         RadarAttackLayer3SummaryDurationListParamsDateRange = "2d"
	RadarAttackLayer3SummaryDurationListParamsDateRange7d         RadarAttackLayer3SummaryDurationListParamsDateRange = "7d"
	RadarAttackLayer3SummaryDurationListParamsDateRange14d        RadarAttackLayer3SummaryDurationListParamsDateRange = "14d"
	RadarAttackLayer3SummaryDurationListParamsDateRange28d        RadarAttackLayer3SummaryDurationListParamsDateRange = "28d"
	RadarAttackLayer3SummaryDurationListParamsDateRange12w        RadarAttackLayer3SummaryDurationListParamsDateRange = "12w"
	RadarAttackLayer3SummaryDurationListParamsDateRange24w        RadarAttackLayer3SummaryDurationListParamsDateRange = "24w"
	RadarAttackLayer3SummaryDurationListParamsDateRange52w        RadarAttackLayer3SummaryDurationListParamsDateRange = "52w"
	RadarAttackLayer3SummaryDurationListParamsDateRange1dControl  RadarAttackLayer3SummaryDurationListParamsDateRange = "1dControl"
	RadarAttackLayer3SummaryDurationListParamsDateRange2dControl  RadarAttackLayer3SummaryDurationListParamsDateRange = "2dControl"
	RadarAttackLayer3SummaryDurationListParamsDateRange7dControl  RadarAttackLayer3SummaryDurationListParamsDateRange = "7dControl"
	RadarAttackLayer3SummaryDurationListParamsDateRange14dControl RadarAttackLayer3SummaryDurationListParamsDateRange = "14dControl"
	RadarAttackLayer3SummaryDurationListParamsDateRange28dControl RadarAttackLayer3SummaryDurationListParamsDateRange = "28dControl"
	RadarAttackLayer3SummaryDurationListParamsDateRange12wControl RadarAttackLayer3SummaryDurationListParamsDateRange = "12wControl"
	RadarAttackLayer3SummaryDurationListParamsDateRange24wControl RadarAttackLayer3SummaryDurationListParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3SummaryDurationListParamsDirection string

const (
	RadarAttackLayer3SummaryDurationListParamsDirectionOrigin RadarAttackLayer3SummaryDurationListParamsDirection = "ORIGIN"
	RadarAttackLayer3SummaryDurationListParamsDirectionTarget RadarAttackLayer3SummaryDurationListParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3SummaryDurationListParamsFormat string

const (
	RadarAttackLayer3SummaryDurationListParamsFormatJson RadarAttackLayer3SummaryDurationListParamsFormat = "JSON"
	RadarAttackLayer3SummaryDurationListParamsFormatCsv  RadarAttackLayer3SummaryDurationListParamsFormat = "CSV"
)

type RadarAttackLayer3SummaryDurationListParamsIPVersion string

const (
	RadarAttackLayer3SummaryDurationListParamsIPVersionIPv4 RadarAttackLayer3SummaryDurationListParamsIPVersion = "IPv4"
	RadarAttackLayer3SummaryDurationListParamsIPVersionIPv6 RadarAttackLayer3SummaryDurationListParamsIPVersion = "IPv6"
)

type RadarAttackLayer3SummaryDurationListParamsProtocol string

const (
	RadarAttackLayer3SummaryDurationListParamsProtocolUdp  RadarAttackLayer3SummaryDurationListParamsProtocol = "UDP"
	RadarAttackLayer3SummaryDurationListParamsProtocolTcp  RadarAttackLayer3SummaryDurationListParamsProtocol = "TCP"
	RadarAttackLayer3SummaryDurationListParamsProtocolIcmp RadarAttackLayer3SummaryDurationListParamsProtocol = "ICMP"
	RadarAttackLayer3SummaryDurationListParamsProtocolGre  RadarAttackLayer3SummaryDurationListParamsProtocol = "GRE"
)
