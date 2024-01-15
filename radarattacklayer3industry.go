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

// RadarAttackLayer3IndustryService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAttackLayer3IndustryService] method instead.
type RadarAttackLayer3IndustryService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3IndustryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3IndustryService(opts ...option.RequestOption) (r *RadarAttackLayer3IndustryService) {
	r = &RadarAttackLayer3IndustryService{}
	r.Options = opts
	return
}

// Get the Industry of attacks.
func (r *RadarAttackLayer3IndustryService) List(ctx context.Context, query RadarAttackLayer3IndustryListParams, opts ...option.RequestOption) (res *RadarAttackLayer3IndustryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/top/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3IndustryListResponse struct {
	Result  RadarAttackLayer3IndustryListResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarAttackLayer3IndustryListResponseJSON   `json:"-"`
}

// radarAttackLayer3IndustryListResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3IndustryListResponse]
type radarAttackLayer3IndustryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3IndustryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3IndustryListResponseResult struct {
	Meta RadarAttackLayer3IndustryListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3IndustryListResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer3IndustryListResponseResultJSON   `json:"-"`
}

// radarAttackLayer3IndustryListResponseResultJSON contains the JSON metadata for
// the struct [RadarAttackLayer3IndustryListResponseResult]
type radarAttackLayer3IndustryListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3IndustryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3IndustryListResponseResultMeta struct {
	DateRange      []RadarAttackLayer3IndustryListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                        `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3IndustryListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3IndustryListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3IndustryListResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAttackLayer3IndustryListResponseResultMeta]
type radarAttackLayer3IndustryListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3IndustryListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3IndustryListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3IndustryListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3IndustryListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3IndustryListResponseResultMetaDateRange]
type radarAttackLayer3IndustryListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3IndustryListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3IndustryListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3IndustryListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                     `json:"level"`
	JSON        radarAttackLayer3IndustryListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3IndustryListResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3IndustryListResponseResultMetaConfidenceInfo]
type radarAttackLayer3IndustryListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3IndustryListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3IndustryListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                      `json:"dataSource,required"`
	Description     string                                                                      `json:"description,required"`
	EventType       string                                                                      `json:"eventType,required"`
	IsInstantaneous interface{}                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                                   `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3IndustryListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3IndustryListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3IndustryListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3IndustryListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3IndustryListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3IndustryListResponseResultTop0 struct {
	Name  string                                              `json:"name,required"`
	Value string                                              `json:"value,required"`
	JSON  radarAttackLayer3IndustryListResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer3IndustryListResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarAttackLayer3IndustryListResponseResultTop0]
type radarAttackLayer3IndustryListResponseResultTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3IndustryListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3IndustryListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3IndustryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3IndustryListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3IndustryListParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3IndustryListParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3IndustryListParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer3IndustryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3IndustryListParamsDateRange string

const (
	RadarAttackLayer3IndustryListParamsDateRange1d         RadarAttackLayer3IndustryListParamsDateRange = "1d"
	RadarAttackLayer3IndustryListParamsDateRange2d         RadarAttackLayer3IndustryListParamsDateRange = "2d"
	RadarAttackLayer3IndustryListParamsDateRange7d         RadarAttackLayer3IndustryListParamsDateRange = "7d"
	RadarAttackLayer3IndustryListParamsDateRange14d        RadarAttackLayer3IndustryListParamsDateRange = "14d"
	RadarAttackLayer3IndustryListParamsDateRange28d        RadarAttackLayer3IndustryListParamsDateRange = "28d"
	RadarAttackLayer3IndustryListParamsDateRange12w        RadarAttackLayer3IndustryListParamsDateRange = "12w"
	RadarAttackLayer3IndustryListParamsDateRange24w        RadarAttackLayer3IndustryListParamsDateRange = "24w"
	RadarAttackLayer3IndustryListParamsDateRange52w        RadarAttackLayer3IndustryListParamsDateRange = "52w"
	RadarAttackLayer3IndustryListParamsDateRange1dControl  RadarAttackLayer3IndustryListParamsDateRange = "1dControl"
	RadarAttackLayer3IndustryListParamsDateRange2dControl  RadarAttackLayer3IndustryListParamsDateRange = "2dControl"
	RadarAttackLayer3IndustryListParamsDateRange7dControl  RadarAttackLayer3IndustryListParamsDateRange = "7dControl"
	RadarAttackLayer3IndustryListParamsDateRange14dControl RadarAttackLayer3IndustryListParamsDateRange = "14dControl"
	RadarAttackLayer3IndustryListParamsDateRange28dControl RadarAttackLayer3IndustryListParamsDateRange = "28dControl"
	RadarAttackLayer3IndustryListParamsDateRange12wControl RadarAttackLayer3IndustryListParamsDateRange = "12wControl"
	RadarAttackLayer3IndustryListParamsDateRange24wControl RadarAttackLayer3IndustryListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3IndustryListParamsFormat string

const (
	RadarAttackLayer3IndustryListParamsFormatJson RadarAttackLayer3IndustryListParamsFormat = "JSON"
	RadarAttackLayer3IndustryListParamsFormatCsv  RadarAttackLayer3IndustryListParamsFormat = "CSV"
)

type RadarAttackLayer3IndustryListParamsIPVersion string

const (
	RadarAttackLayer3IndustryListParamsIPVersionIPv4 RadarAttackLayer3IndustryListParamsIPVersion = "IPv4"
	RadarAttackLayer3IndustryListParamsIPVersionIPv6 RadarAttackLayer3IndustryListParamsIPVersion = "IPv6"
)

type RadarAttackLayer3IndustryListParamsProtocol string

const (
	RadarAttackLayer3IndustryListParamsProtocolUdp  RadarAttackLayer3IndustryListParamsProtocol = "UDP"
	RadarAttackLayer3IndustryListParamsProtocolTcp  RadarAttackLayer3IndustryListParamsProtocol = "TCP"
	RadarAttackLayer3IndustryListParamsProtocolIcmp RadarAttackLayer3IndustryListParamsProtocol = "ICMP"
	RadarAttackLayer3IndustryListParamsProtocolGre  RadarAttackLayer3IndustryListParamsProtocol = "GRE"
)
