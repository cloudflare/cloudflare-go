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

// RadarAttackLayer3SummaryProtocolService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer3SummaryProtocolService] method instead.
type RadarAttackLayer3SummaryProtocolService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3SummaryProtocolService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3SummaryProtocolService(opts ...option.RequestOption) (r *RadarAttackLayer3SummaryProtocolService) {
	r = &RadarAttackLayer3SummaryProtocolService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by protocol used.
func (r *RadarAttackLayer3SummaryProtocolService) List(ctx context.Context, query RadarAttackLayer3SummaryProtocolListParams, opts ...option.RequestOption) (res *RadarAttackLayer3SummaryProtocolListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3SummaryProtocolListResponse struct {
	Result  RadarAttackLayer3SummaryProtocolListResponseResult `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarAttackLayer3SummaryProtocolListResponseJSON   `json:"-"`
}

// radarAttackLayer3SummaryProtocolListResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer3SummaryProtocolListResponse]
type radarAttackLayer3SummaryProtocolListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryProtocolListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryProtocolListResponseResult struct {
	Meta     RadarAttackLayer3SummaryProtocolListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer3SummaryProtocolListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer3SummaryProtocolListResponseResultJSON     `json:"-"`
}

// radarAttackLayer3SummaryProtocolListResponseResultJSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryProtocolListResponseResult]
type radarAttackLayer3SummaryProtocolListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryProtocolListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryProtocolListResponseResultMeta struct {
	DateRange      []RadarAttackLayer3SummaryProtocolListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                               `json:"lastUpdated,required"`
	Normalization  string                                                               `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer3SummaryProtocolListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3SummaryProtocolListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3SummaryProtocolListResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryProtocolListResponseResultMeta]
type radarAttackLayer3SummaryProtocolListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryProtocolListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryProtocolListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                           `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3SummaryProtocolListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3SummaryProtocolListResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryProtocolListResponseResultMetaDateRange]
type radarAttackLayer3SummaryProtocolListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryProtocolListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryProtocolListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryProtocolListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                            `json:"level"`
	JSON        radarAttackLayer3SummaryProtocolListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3SummaryProtocolListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryProtocolListResponseResultMetaConfidenceInfo]
type radarAttackLayer3SummaryProtocolListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryProtocolListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryProtocolListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                             `json:"dataSource,required"`
	Description     string                                                                             `json:"description,required"`
	EventType       string                                                                             `json:"eventType,required"`
	IsInstantaneous interface{}                                                                        `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                          `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                             `json:"linkedUrl"`
	StartTime       time.Time                                                                          `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3SummaryProtocolListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryProtocolListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryProtocolListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryProtocolListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryProtocolListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryProtocolListResponseResultSummary0 struct {
	Gre  string                                                         `json:"GRE,required"`
	Icmp string                                                         `json:"ICMP,required"`
	Tcp  string                                                         `json:"TCP,required"`
	Udp  string                                                         `json:"UDP,required"`
	JSON radarAttackLayer3SummaryProtocolListResponseResultSummary0JSON `json:"-"`
}

// radarAttackLayer3SummaryProtocolListResponseResultSummary0JSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3SummaryProtocolListResponseResultSummary0]
type radarAttackLayer3SummaryProtocolListResponseResultSummary0JSON struct {
	Gre         apijson.Field
	Icmp        apijson.Field
	Tcp         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryProtocolListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryProtocolListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3SummaryProtocolListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3SummaryProtocolListParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3SummaryProtocolListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3SummaryProtocolListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer3SummaryProtocolListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3SummaryProtocolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3SummaryProtocolListParamsDateRange string

const (
	RadarAttackLayer3SummaryProtocolListParamsDateRange1d         RadarAttackLayer3SummaryProtocolListParamsDateRange = "1d"
	RadarAttackLayer3SummaryProtocolListParamsDateRange2d         RadarAttackLayer3SummaryProtocolListParamsDateRange = "2d"
	RadarAttackLayer3SummaryProtocolListParamsDateRange7d         RadarAttackLayer3SummaryProtocolListParamsDateRange = "7d"
	RadarAttackLayer3SummaryProtocolListParamsDateRange14d        RadarAttackLayer3SummaryProtocolListParamsDateRange = "14d"
	RadarAttackLayer3SummaryProtocolListParamsDateRange28d        RadarAttackLayer3SummaryProtocolListParamsDateRange = "28d"
	RadarAttackLayer3SummaryProtocolListParamsDateRange12w        RadarAttackLayer3SummaryProtocolListParamsDateRange = "12w"
	RadarAttackLayer3SummaryProtocolListParamsDateRange24w        RadarAttackLayer3SummaryProtocolListParamsDateRange = "24w"
	RadarAttackLayer3SummaryProtocolListParamsDateRange52w        RadarAttackLayer3SummaryProtocolListParamsDateRange = "52w"
	RadarAttackLayer3SummaryProtocolListParamsDateRange1dControl  RadarAttackLayer3SummaryProtocolListParamsDateRange = "1dControl"
	RadarAttackLayer3SummaryProtocolListParamsDateRange2dControl  RadarAttackLayer3SummaryProtocolListParamsDateRange = "2dControl"
	RadarAttackLayer3SummaryProtocolListParamsDateRange7dControl  RadarAttackLayer3SummaryProtocolListParamsDateRange = "7dControl"
	RadarAttackLayer3SummaryProtocolListParamsDateRange14dControl RadarAttackLayer3SummaryProtocolListParamsDateRange = "14dControl"
	RadarAttackLayer3SummaryProtocolListParamsDateRange28dControl RadarAttackLayer3SummaryProtocolListParamsDateRange = "28dControl"
	RadarAttackLayer3SummaryProtocolListParamsDateRange12wControl RadarAttackLayer3SummaryProtocolListParamsDateRange = "12wControl"
	RadarAttackLayer3SummaryProtocolListParamsDateRange24wControl RadarAttackLayer3SummaryProtocolListParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3SummaryProtocolListParamsDirection string

const (
	RadarAttackLayer3SummaryProtocolListParamsDirectionOrigin RadarAttackLayer3SummaryProtocolListParamsDirection = "ORIGIN"
	RadarAttackLayer3SummaryProtocolListParamsDirectionTarget RadarAttackLayer3SummaryProtocolListParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3SummaryProtocolListParamsFormat string

const (
	RadarAttackLayer3SummaryProtocolListParamsFormatJson RadarAttackLayer3SummaryProtocolListParamsFormat = "JSON"
	RadarAttackLayer3SummaryProtocolListParamsFormatCsv  RadarAttackLayer3SummaryProtocolListParamsFormat = "CSV"
)

type RadarAttackLayer3SummaryProtocolListParamsIPVersion string

const (
	RadarAttackLayer3SummaryProtocolListParamsIPVersionIPv4 RadarAttackLayer3SummaryProtocolListParamsIPVersion = "IPv4"
	RadarAttackLayer3SummaryProtocolListParamsIPVersionIPv6 RadarAttackLayer3SummaryProtocolListParamsIPVersion = "IPv6"
)
