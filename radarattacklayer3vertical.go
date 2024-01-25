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

// RadarAttackLayer3VerticalService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAttackLayer3VerticalService] method instead.
type RadarAttackLayer3VerticalService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3VerticalService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3VerticalService(opts ...option.RequestOption) (r *RadarAttackLayer3VerticalService) {
	r = &RadarAttackLayer3VerticalService{}
	r.Options = opts
	return
}

// Get the Verticals of attacks.
func (r *RadarAttackLayer3VerticalService) List(ctx context.Context, query RadarAttackLayer3VerticalListParams, opts ...option.RequestOption) (res *RadarAttackLayer3VerticalListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/top/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3VerticalListResponse struct {
	Result  RadarAttackLayer3VerticalListResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarAttackLayer3VerticalListResponseJSON   `json:"-"`
}

// radarAttackLayer3VerticalListResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3VerticalListResponse]
type radarAttackLayer3VerticalListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3VerticalListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3VerticalListResponseResult struct {
	Meta RadarAttackLayer3VerticalListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3VerticalListResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer3VerticalListResponseResultJSON   `json:"-"`
}

// radarAttackLayer3VerticalListResponseResultJSON contains the JSON metadata for
// the struct [RadarAttackLayer3VerticalListResponseResult]
type radarAttackLayer3VerticalListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3VerticalListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3VerticalListResponseResultMeta struct {
	DateRange      []RadarAttackLayer3VerticalListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                        `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3VerticalListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3VerticalListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3VerticalListResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAttackLayer3VerticalListResponseResultMeta]
type radarAttackLayer3VerticalListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3VerticalListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3VerticalListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3VerticalListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3VerticalListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3VerticalListResponseResultMetaDateRange]
type radarAttackLayer3VerticalListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3VerticalListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3VerticalListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3VerticalListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                     `json:"level"`
	JSON        radarAttackLayer3VerticalListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3VerticalListResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3VerticalListResponseResultMetaConfidenceInfo]
type radarAttackLayer3VerticalListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3VerticalListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3VerticalListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                      `json:"dataSource,required"`
	Description     string                                                                      `json:"description,required"`
	EventType       string                                                                      `json:"eventType,required"`
	IsInstantaneous interface{}                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                                   `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3VerticalListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3VerticalListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3VerticalListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3VerticalListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3VerticalListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3VerticalListResponseResultTop0 struct {
	Name  string                                              `json:"name,required"`
	Value string                                              `json:"value,required"`
	JSON  radarAttackLayer3VerticalListResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer3VerticalListResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarAttackLayer3VerticalListResponseResultTop0]
type radarAttackLayer3VerticalListResponseResultTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3VerticalListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3VerticalListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3VerticalListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3VerticalListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3VerticalListParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3VerticalListParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3VerticalListParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer3VerticalListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3VerticalListParamsDateRange string

const (
	RadarAttackLayer3VerticalListParamsDateRange1d         RadarAttackLayer3VerticalListParamsDateRange = "1d"
	RadarAttackLayer3VerticalListParamsDateRange2d         RadarAttackLayer3VerticalListParamsDateRange = "2d"
	RadarAttackLayer3VerticalListParamsDateRange7d         RadarAttackLayer3VerticalListParamsDateRange = "7d"
	RadarAttackLayer3VerticalListParamsDateRange14d        RadarAttackLayer3VerticalListParamsDateRange = "14d"
	RadarAttackLayer3VerticalListParamsDateRange28d        RadarAttackLayer3VerticalListParamsDateRange = "28d"
	RadarAttackLayer3VerticalListParamsDateRange12w        RadarAttackLayer3VerticalListParamsDateRange = "12w"
	RadarAttackLayer3VerticalListParamsDateRange24w        RadarAttackLayer3VerticalListParamsDateRange = "24w"
	RadarAttackLayer3VerticalListParamsDateRange52w        RadarAttackLayer3VerticalListParamsDateRange = "52w"
	RadarAttackLayer3VerticalListParamsDateRange1dControl  RadarAttackLayer3VerticalListParamsDateRange = "1dControl"
	RadarAttackLayer3VerticalListParamsDateRange2dControl  RadarAttackLayer3VerticalListParamsDateRange = "2dControl"
	RadarAttackLayer3VerticalListParamsDateRange7dControl  RadarAttackLayer3VerticalListParamsDateRange = "7dControl"
	RadarAttackLayer3VerticalListParamsDateRange14dControl RadarAttackLayer3VerticalListParamsDateRange = "14dControl"
	RadarAttackLayer3VerticalListParamsDateRange28dControl RadarAttackLayer3VerticalListParamsDateRange = "28dControl"
	RadarAttackLayer3VerticalListParamsDateRange12wControl RadarAttackLayer3VerticalListParamsDateRange = "12wControl"
	RadarAttackLayer3VerticalListParamsDateRange24wControl RadarAttackLayer3VerticalListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3VerticalListParamsFormat string

const (
	RadarAttackLayer3VerticalListParamsFormatJson RadarAttackLayer3VerticalListParamsFormat = "JSON"
	RadarAttackLayer3VerticalListParamsFormatCsv  RadarAttackLayer3VerticalListParamsFormat = "CSV"
)

type RadarAttackLayer3VerticalListParamsIPVersion string

const (
	RadarAttackLayer3VerticalListParamsIPVersionIPv4 RadarAttackLayer3VerticalListParamsIPVersion = "IPv4"
	RadarAttackLayer3VerticalListParamsIPVersionIPv6 RadarAttackLayer3VerticalListParamsIPVersion = "IPv6"
)

type RadarAttackLayer3VerticalListParamsProtocol string

const (
	RadarAttackLayer3VerticalListParamsProtocolUdp  RadarAttackLayer3VerticalListParamsProtocol = "UDP"
	RadarAttackLayer3VerticalListParamsProtocolTcp  RadarAttackLayer3VerticalListParamsProtocol = "TCP"
	RadarAttackLayer3VerticalListParamsProtocolIcmp RadarAttackLayer3VerticalListParamsProtocol = "ICMP"
	RadarAttackLayer3VerticalListParamsProtocolGre  RadarAttackLayer3VerticalListParamsProtocol = "GRE"
)
