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

// RadarAttackLayer3SummaryBitrateService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer3SummaryBitrateService] method instead.
type RadarAttackLayer3SummaryBitrateService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3SummaryBitrateService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3SummaryBitrateService(opts ...option.RequestOption) (r *RadarAttackLayer3SummaryBitrateService) {
	r = &RadarAttackLayer3SummaryBitrateService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by bitrate.
func (r *RadarAttackLayer3SummaryBitrateService) Get(ctx context.Context, query RadarAttackLayer3SummaryBitrateGetParams, opts ...option.RequestOption) (res *RadarAttackLayer3SummaryBitrateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/bitrate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3SummaryBitrateGetResponse struct {
	Result  RadarAttackLayer3SummaryBitrateGetResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarAttackLayer3SummaryBitrateGetResponseJSON   `json:"-"`
}

// radarAttackLayer3SummaryBitrateGetResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer3SummaryBitrateGetResponse]
type radarAttackLayer3SummaryBitrateGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryBitrateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryBitrateGetResponseResult struct {
	Meta     RadarAttackLayer3SummaryBitrateGetResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer3SummaryBitrateGetResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer3SummaryBitrateGetResponseResultJSON     `json:"-"`
}

// radarAttackLayer3SummaryBitrateGetResponseResultJSON contains the JSON metadata
// for the struct [RadarAttackLayer3SummaryBitrateGetResponseResult]
type radarAttackLayer3SummaryBitrateGetResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryBitrateGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryBitrateGetResponseResultMeta struct {
	DateRange      []RadarAttackLayer3SummaryBitrateGetResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                             `json:"lastUpdated,required"`
	Normalization  string                                                             `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer3SummaryBitrateGetResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3SummaryBitrateGetResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3SummaryBitrateGetResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryBitrateGetResponseResultMeta]
type radarAttackLayer3SummaryBitrateGetResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryBitrateGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryBitrateGetResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                         `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3SummaryBitrateGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3SummaryBitrateGetResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryBitrateGetResponseResultMetaDateRange]
type radarAttackLayer3SummaryBitrateGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryBitrateGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryBitrateGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryBitrateGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                          `json:"level"`
	JSON        radarAttackLayer3SummaryBitrateGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3SummaryBitrateGetResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3SummaryBitrateGetResponseResultMetaConfidenceInfo]
type radarAttackLayer3SummaryBitrateGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryBitrateGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryBitrateGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                           `json:"dataSource,required"`
	Description     string                                                                           `json:"description,required"`
	EventType       string                                                                           `json:"eventType,required"`
	IsInstantaneous interface{}                                                                      `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                        `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                           `json:"linkedUrl"`
	StartTime       time.Time                                                                        `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3SummaryBitrateGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryBitrateGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryBitrateGetResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryBitrateGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryBitrateGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryBitrateGetResponseResultSummary0 struct {
	Number1GbpsTo10Gbps   string                                                       `json:"_1_GBPS_TO_10_GBPS,required"`
	Number10GbpsTo100Gbps string                                                       `json:"_10_GBPS_TO_100_GBPS,required"`
	Number500MbpsTo1Gbps  string                                                       `json:"_500_MBPS_TO_1_GBPS,required"`
	Over100Gbps           string                                                       `json:"OVER_100_GBPS,required"`
	Under500Mbps          string                                                       `json:"UNDER_500_MBPS,required"`
	JSON                  radarAttackLayer3SummaryBitrateGetResponseResultSummary0JSON `json:"-"`
}

// radarAttackLayer3SummaryBitrateGetResponseResultSummary0JSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3SummaryBitrateGetResponseResultSummary0]
type radarAttackLayer3SummaryBitrateGetResponseResultSummary0JSON struct {
	Number1GbpsTo10Gbps   apijson.Field
	Number10GbpsTo100Gbps apijson.Field
	Number500MbpsTo1Gbps  apijson.Field
	Over100Gbps           apijson.Field
	Under500Mbps          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryBitrateGetResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3SummaryBitrateGetParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3SummaryBitrateGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3SummaryBitrateGetParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3SummaryBitrateGetParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3SummaryBitrateGetParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3SummaryBitrateGetParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3SummaryBitrateGetParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3SummaryBitrateGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3SummaryBitrateGetParamsDateRange string

const (
	RadarAttackLayer3SummaryBitrateGetParamsDateRange1d         RadarAttackLayer3SummaryBitrateGetParamsDateRange = "1d"
	RadarAttackLayer3SummaryBitrateGetParamsDateRange2d         RadarAttackLayer3SummaryBitrateGetParamsDateRange = "2d"
	RadarAttackLayer3SummaryBitrateGetParamsDateRange7d         RadarAttackLayer3SummaryBitrateGetParamsDateRange = "7d"
	RadarAttackLayer3SummaryBitrateGetParamsDateRange14d        RadarAttackLayer3SummaryBitrateGetParamsDateRange = "14d"
	RadarAttackLayer3SummaryBitrateGetParamsDateRange28d        RadarAttackLayer3SummaryBitrateGetParamsDateRange = "28d"
	RadarAttackLayer3SummaryBitrateGetParamsDateRange12w        RadarAttackLayer3SummaryBitrateGetParamsDateRange = "12w"
	RadarAttackLayer3SummaryBitrateGetParamsDateRange24w        RadarAttackLayer3SummaryBitrateGetParamsDateRange = "24w"
	RadarAttackLayer3SummaryBitrateGetParamsDateRange52w        RadarAttackLayer3SummaryBitrateGetParamsDateRange = "52w"
	RadarAttackLayer3SummaryBitrateGetParamsDateRange1dControl  RadarAttackLayer3SummaryBitrateGetParamsDateRange = "1dControl"
	RadarAttackLayer3SummaryBitrateGetParamsDateRange2dControl  RadarAttackLayer3SummaryBitrateGetParamsDateRange = "2dControl"
	RadarAttackLayer3SummaryBitrateGetParamsDateRange7dControl  RadarAttackLayer3SummaryBitrateGetParamsDateRange = "7dControl"
	RadarAttackLayer3SummaryBitrateGetParamsDateRange14dControl RadarAttackLayer3SummaryBitrateGetParamsDateRange = "14dControl"
	RadarAttackLayer3SummaryBitrateGetParamsDateRange28dControl RadarAttackLayer3SummaryBitrateGetParamsDateRange = "28dControl"
	RadarAttackLayer3SummaryBitrateGetParamsDateRange12wControl RadarAttackLayer3SummaryBitrateGetParamsDateRange = "12wControl"
	RadarAttackLayer3SummaryBitrateGetParamsDateRange24wControl RadarAttackLayer3SummaryBitrateGetParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3SummaryBitrateGetParamsDirection string

const (
	RadarAttackLayer3SummaryBitrateGetParamsDirectionOrigin RadarAttackLayer3SummaryBitrateGetParamsDirection = "ORIGIN"
	RadarAttackLayer3SummaryBitrateGetParamsDirectionTarget RadarAttackLayer3SummaryBitrateGetParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3SummaryBitrateGetParamsFormat string

const (
	RadarAttackLayer3SummaryBitrateGetParamsFormatJson RadarAttackLayer3SummaryBitrateGetParamsFormat = "JSON"
	RadarAttackLayer3SummaryBitrateGetParamsFormatCsv  RadarAttackLayer3SummaryBitrateGetParamsFormat = "CSV"
)

type RadarAttackLayer3SummaryBitrateGetParamsIPVersion string

const (
	RadarAttackLayer3SummaryBitrateGetParamsIPVersionIPv4 RadarAttackLayer3SummaryBitrateGetParamsIPVersion = "IPv4"
	RadarAttackLayer3SummaryBitrateGetParamsIPVersionIPv6 RadarAttackLayer3SummaryBitrateGetParamsIPVersion = "IPv6"
)

type RadarAttackLayer3SummaryBitrateGetParamsProtocol string

const (
	RadarAttackLayer3SummaryBitrateGetParamsProtocolUdp  RadarAttackLayer3SummaryBitrateGetParamsProtocol = "UDP"
	RadarAttackLayer3SummaryBitrateGetParamsProtocolTcp  RadarAttackLayer3SummaryBitrateGetParamsProtocol = "TCP"
	RadarAttackLayer3SummaryBitrateGetParamsProtocolIcmp RadarAttackLayer3SummaryBitrateGetParamsProtocol = "ICMP"
	RadarAttackLayer3SummaryBitrateGetParamsProtocolGre  RadarAttackLayer3SummaryBitrateGetParamsProtocol = "GRE"
)
