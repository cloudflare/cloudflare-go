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

// RadarAttackLayer3SummaryVectorService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer3SummaryVectorService] method instead.
type RadarAttackLayer3SummaryVectorService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3SummaryVectorService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3SummaryVectorService(opts ...option.RequestOption) (r *RadarAttackLayer3SummaryVectorService) {
	r = &RadarAttackLayer3SummaryVectorService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by vector.
func (r *RadarAttackLayer3SummaryVectorService) List(ctx context.Context, query RadarAttackLayer3SummaryVectorListParams, opts ...option.RequestOption) (res *RadarAttackLayer3SummaryVectorListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/vector"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3SummaryVectorListResponse struct {
	Result  RadarAttackLayer3SummaryVectorListResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarAttackLayer3SummaryVectorListResponseJSON   `json:"-"`
}

// radarAttackLayer3SummaryVectorListResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer3SummaryVectorListResponse]
type radarAttackLayer3SummaryVectorListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryVectorListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryVectorListResponseResult struct {
	Meta     RadarAttackLayer3SummaryVectorListResponseResultMeta `json:"meta,required"`
	Summary0 interface{}                                          `json:"summary_0,required"`
	JSON     radarAttackLayer3SummaryVectorListResponseResultJSON `json:"-"`
}

// radarAttackLayer3SummaryVectorListResponseResultJSON contains the JSON metadata
// for the struct [RadarAttackLayer3SummaryVectorListResponseResult]
type radarAttackLayer3SummaryVectorListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryVectorListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryVectorListResponseResultMeta struct {
	DateRange      []RadarAttackLayer3SummaryVectorListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                             `json:"lastUpdated,required"`
	Normalization  string                                                             `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer3SummaryVectorListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3SummaryVectorListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3SummaryVectorListResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryVectorListResponseResultMeta]
type radarAttackLayer3SummaryVectorListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryVectorListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryVectorListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                         `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3SummaryVectorListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3SummaryVectorListResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryVectorListResponseResultMetaDateRange]
type radarAttackLayer3SummaryVectorListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryVectorListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryVectorListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryVectorListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                          `json:"level"`
	JSON        radarAttackLayer3SummaryVectorListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3SummaryVectorListResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3SummaryVectorListResponseResultMetaConfidenceInfo]
type radarAttackLayer3SummaryVectorListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryVectorListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryVectorListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                           `json:"dataSource,required"`
	Description     string                                                                           `json:"description,required"`
	EventType       string                                                                           `json:"eventType,required"`
	IsInstantaneous interface{}                                                                      `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                        `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                           `json:"linkedUrl"`
	StartTime       time.Time                                                                        `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3SummaryVectorListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryVectorListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryVectorListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryVectorListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryVectorListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryVectorListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3SummaryVectorListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3SummaryVectorListParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3SummaryVectorListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3SummaryVectorListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3SummaryVectorListParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3SummaryVectorListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3SummaryVectorListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3SummaryVectorListParamsDateRange string

const (
	RadarAttackLayer3SummaryVectorListParamsDateRange1d         RadarAttackLayer3SummaryVectorListParamsDateRange = "1d"
	RadarAttackLayer3SummaryVectorListParamsDateRange2d         RadarAttackLayer3SummaryVectorListParamsDateRange = "2d"
	RadarAttackLayer3SummaryVectorListParamsDateRange7d         RadarAttackLayer3SummaryVectorListParamsDateRange = "7d"
	RadarAttackLayer3SummaryVectorListParamsDateRange14d        RadarAttackLayer3SummaryVectorListParamsDateRange = "14d"
	RadarAttackLayer3SummaryVectorListParamsDateRange28d        RadarAttackLayer3SummaryVectorListParamsDateRange = "28d"
	RadarAttackLayer3SummaryVectorListParamsDateRange12w        RadarAttackLayer3SummaryVectorListParamsDateRange = "12w"
	RadarAttackLayer3SummaryVectorListParamsDateRange24w        RadarAttackLayer3SummaryVectorListParamsDateRange = "24w"
	RadarAttackLayer3SummaryVectorListParamsDateRange52w        RadarAttackLayer3SummaryVectorListParamsDateRange = "52w"
	RadarAttackLayer3SummaryVectorListParamsDateRange1dControl  RadarAttackLayer3SummaryVectorListParamsDateRange = "1dControl"
	RadarAttackLayer3SummaryVectorListParamsDateRange2dControl  RadarAttackLayer3SummaryVectorListParamsDateRange = "2dControl"
	RadarAttackLayer3SummaryVectorListParamsDateRange7dControl  RadarAttackLayer3SummaryVectorListParamsDateRange = "7dControl"
	RadarAttackLayer3SummaryVectorListParamsDateRange14dControl RadarAttackLayer3SummaryVectorListParamsDateRange = "14dControl"
	RadarAttackLayer3SummaryVectorListParamsDateRange28dControl RadarAttackLayer3SummaryVectorListParamsDateRange = "28dControl"
	RadarAttackLayer3SummaryVectorListParamsDateRange12wControl RadarAttackLayer3SummaryVectorListParamsDateRange = "12wControl"
	RadarAttackLayer3SummaryVectorListParamsDateRange24wControl RadarAttackLayer3SummaryVectorListParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3SummaryVectorListParamsDirection string

const (
	RadarAttackLayer3SummaryVectorListParamsDirectionOrigin RadarAttackLayer3SummaryVectorListParamsDirection = "ORIGIN"
	RadarAttackLayer3SummaryVectorListParamsDirectionTarget RadarAttackLayer3SummaryVectorListParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3SummaryVectorListParamsFormat string

const (
	RadarAttackLayer3SummaryVectorListParamsFormatJson RadarAttackLayer3SummaryVectorListParamsFormat = "JSON"
	RadarAttackLayer3SummaryVectorListParamsFormatCsv  RadarAttackLayer3SummaryVectorListParamsFormat = "CSV"
)

type RadarAttackLayer3SummaryVectorListParamsIPVersion string

const (
	RadarAttackLayer3SummaryVectorListParamsIPVersionIPv4 RadarAttackLayer3SummaryVectorListParamsIPVersion = "IPv4"
	RadarAttackLayer3SummaryVectorListParamsIPVersionIPv6 RadarAttackLayer3SummaryVectorListParamsIPVersion = "IPv6"
)

type RadarAttackLayer3SummaryVectorListParamsProtocol string

const (
	RadarAttackLayer3SummaryVectorListParamsProtocolUdp  RadarAttackLayer3SummaryVectorListParamsProtocol = "UDP"
	RadarAttackLayer3SummaryVectorListParamsProtocolTcp  RadarAttackLayer3SummaryVectorListParamsProtocol = "TCP"
	RadarAttackLayer3SummaryVectorListParamsProtocolIcmp RadarAttackLayer3SummaryVectorListParamsProtocol = "ICMP"
	RadarAttackLayer3SummaryVectorListParamsProtocolGre  RadarAttackLayer3SummaryVectorListParamsProtocol = "GRE"
)
