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

// RadarAttackLayer3SummaryIPVersionService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer3SummaryIPVersionService] method instead.
type RadarAttackLayer3SummaryIPVersionService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3SummaryIPVersionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3SummaryIPVersionService(opts ...option.RequestOption) (r *RadarAttackLayer3SummaryIPVersionService) {
	r = &RadarAttackLayer3SummaryIPVersionService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by ip version used.
func (r *RadarAttackLayer3SummaryIPVersionService) List(ctx context.Context, query RadarAttackLayer3SummaryIPVersionListParams, opts ...option.RequestOption) (res *RadarAttackLayer3SummaryIPVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3SummaryIPVersionListResponse struct {
	Result  RadarAttackLayer3SummaryIPVersionListResponseResult `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarAttackLayer3SummaryIPVersionListResponseJSON   `json:"-"`
}

// radarAttackLayer3SummaryIPVersionListResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer3SummaryIPVersionListResponse]
type radarAttackLayer3SummaryIPVersionListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryIPVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryIPVersionListResponseResult struct {
	Meta     RadarAttackLayer3SummaryIPVersionListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer3SummaryIPVersionListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer3SummaryIPVersionListResponseResultJSON     `json:"-"`
}

// radarAttackLayer3SummaryIPVersionListResponseResultJSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryIPVersionListResponseResult]
type radarAttackLayer3SummaryIPVersionListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryIPVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryIPVersionListResponseResultMeta struct {
	DateRange      []RadarAttackLayer3SummaryIPVersionListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                `json:"lastUpdated,required"`
	Normalization  string                                                                `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer3SummaryIPVersionListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3SummaryIPVersionListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3SummaryIPVersionListResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3SummaryIPVersionListResponseResultMeta]
type radarAttackLayer3SummaryIPVersionListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryIPVersionListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryIPVersionListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                            `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3SummaryIPVersionListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3SummaryIPVersionListResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3SummaryIPVersionListResponseResultMetaDateRange]
type radarAttackLayer3SummaryIPVersionListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryIPVersionListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryIPVersionListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryIPVersionListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                             `json:"level"`
	JSON        radarAttackLayer3SummaryIPVersionListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3SummaryIPVersionListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryIPVersionListResponseResultMetaConfidenceInfo]
type radarAttackLayer3SummaryIPVersionListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryIPVersionListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryIPVersionListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                              `json:"dataSource,required"`
	Description     string                                                                              `json:"description,required"`
	EventType       string                                                                              `json:"eventType,required"`
	IsInstantaneous interface{}                                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                                           `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3SummaryIPVersionListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryIPVersionListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryIPVersionListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryIPVersionListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryIPVersionListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryIPVersionListResponseResultSummary0 struct {
	IPv4 string                                                          `json:"IPv4,required"`
	IPv6 string                                                          `json:"IPv6,required"`
	JSON radarAttackLayer3SummaryIPVersionListResponseResultSummary0JSON `json:"-"`
}

// radarAttackLayer3SummaryIPVersionListResponseResultSummary0JSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryIPVersionListResponseResultSummary0]
type radarAttackLayer3SummaryIPVersionListResponseResultSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryIPVersionListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryIPVersionListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3SummaryIPVersionListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3SummaryIPVersionListParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3SummaryIPVersionListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3SummaryIPVersionListParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3SummaryIPVersionListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3SummaryIPVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3SummaryIPVersionListParamsDateRange string

const (
	RadarAttackLayer3SummaryIPVersionListParamsDateRange1d         RadarAttackLayer3SummaryIPVersionListParamsDateRange = "1d"
	RadarAttackLayer3SummaryIPVersionListParamsDateRange2d         RadarAttackLayer3SummaryIPVersionListParamsDateRange = "2d"
	RadarAttackLayer3SummaryIPVersionListParamsDateRange7d         RadarAttackLayer3SummaryIPVersionListParamsDateRange = "7d"
	RadarAttackLayer3SummaryIPVersionListParamsDateRange14d        RadarAttackLayer3SummaryIPVersionListParamsDateRange = "14d"
	RadarAttackLayer3SummaryIPVersionListParamsDateRange28d        RadarAttackLayer3SummaryIPVersionListParamsDateRange = "28d"
	RadarAttackLayer3SummaryIPVersionListParamsDateRange12w        RadarAttackLayer3SummaryIPVersionListParamsDateRange = "12w"
	RadarAttackLayer3SummaryIPVersionListParamsDateRange24w        RadarAttackLayer3SummaryIPVersionListParamsDateRange = "24w"
	RadarAttackLayer3SummaryIPVersionListParamsDateRange52w        RadarAttackLayer3SummaryIPVersionListParamsDateRange = "52w"
	RadarAttackLayer3SummaryIPVersionListParamsDateRange1dControl  RadarAttackLayer3SummaryIPVersionListParamsDateRange = "1dControl"
	RadarAttackLayer3SummaryIPVersionListParamsDateRange2dControl  RadarAttackLayer3SummaryIPVersionListParamsDateRange = "2dControl"
	RadarAttackLayer3SummaryIPVersionListParamsDateRange7dControl  RadarAttackLayer3SummaryIPVersionListParamsDateRange = "7dControl"
	RadarAttackLayer3SummaryIPVersionListParamsDateRange14dControl RadarAttackLayer3SummaryIPVersionListParamsDateRange = "14dControl"
	RadarAttackLayer3SummaryIPVersionListParamsDateRange28dControl RadarAttackLayer3SummaryIPVersionListParamsDateRange = "28dControl"
	RadarAttackLayer3SummaryIPVersionListParamsDateRange12wControl RadarAttackLayer3SummaryIPVersionListParamsDateRange = "12wControl"
	RadarAttackLayer3SummaryIPVersionListParamsDateRange24wControl RadarAttackLayer3SummaryIPVersionListParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3SummaryIPVersionListParamsDirection string

const (
	RadarAttackLayer3SummaryIPVersionListParamsDirectionOrigin RadarAttackLayer3SummaryIPVersionListParamsDirection = "ORIGIN"
	RadarAttackLayer3SummaryIPVersionListParamsDirectionTarget RadarAttackLayer3SummaryIPVersionListParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3SummaryIPVersionListParamsFormat string

const (
	RadarAttackLayer3SummaryIPVersionListParamsFormatJson RadarAttackLayer3SummaryIPVersionListParamsFormat = "JSON"
	RadarAttackLayer3SummaryIPVersionListParamsFormatCsv  RadarAttackLayer3SummaryIPVersionListParamsFormat = "CSV"
)

type RadarAttackLayer3SummaryIPVersionListParamsProtocol string

const (
	RadarAttackLayer3SummaryIPVersionListParamsProtocolUdp  RadarAttackLayer3SummaryIPVersionListParamsProtocol = "UDP"
	RadarAttackLayer3SummaryIPVersionListParamsProtocolTcp  RadarAttackLayer3SummaryIPVersionListParamsProtocol = "TCP"
	RadarAttackLayer3SummaryIPVersionListParamsProtocolIcmp RadarAttackLayer3SummaryIPVersionListParamsProtocol = "ICMP"
	RadarAttackLayer3SummaryIPVersionListParamsProtocolGre  RadarAttackLayer3SummaryIPVersionListParamsProtocol = "GRE"
)
