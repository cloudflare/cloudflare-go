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

// RadarAttackLayer3LocationTargetService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer3LocationTargetService] method instead.
type RadarAttackLayer3LocationTargetService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3LocationTargetService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3LocationTargetService(opts ...option.RequestOption) (r *RadarAttackLayer3LocationTargetService) {
	r = &RadarAttackLayer3LocationTargetService{}
	r.Options = opts
	return
}

// Get the target locations of attacks.
func (r *RadarAttackLayer3LocationTargetService) List(ctx context.Context, query RadarAttackLayer3LocationTargetListParams, opts ...option.RequestOption) (res *RadarAttackLayer3LocationTargetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/top/locations/target"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3LocationTargetListResponse struct {
	Result  RadarAttackLayer3LocationTargetListResponseResult `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarAttackLayer3LocationTargetListResponseJSON   `json:"-"`
}

// radarAttackLayer3LocationTargetListResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer3LocationTargetListResponse]
type radarAttackLayer3LocationTargetListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3LocationTargetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3LocationTargetListResponseResult struct {
	Meta RadarAttackLayer3LocationTargetListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3LocationTargetListResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer3LocationTargetListResponseResultJSON   `json:"-"`
}

// radarAttackLayer3LocationTargetListResponseResultJSON contains the JSON metadata
// for the struct [RadarAttackLayer3LocationTargetListResponseResult]
type radarAttackLayer3LocationTargetListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3LocationTargetListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3LocationTargetListResponseResultMeta struct {
	DateRange      []RadarAttackLayer3LocationTargetListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                              `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3LocationTargetListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3LocationTargetListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3LocationTargetListResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer3LocationTargetListResponseResultMeta]
type radarAttackLayer3LocationTargetListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3LocationTargetListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3LocationTargetListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                          `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3LocationTargetListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3LocationTargetListResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3LocationTargetListResponseResultMetaDateRange]
type radarAttackLayer3LocationTargetListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3LocationTargetListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3LocationTargetListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3LocationTargetListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                           `json:"level"`
	JSON        radarAttackLayer3LocationTargetListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3LocationTargetListResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3LocationTargetListResponseResultMetaConfidenceInfo]
type radarAttackLayer3LocationTargetListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3LocationTargetListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3LocationTargetListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                            `json:"dataSource,required"`
	Description     string                                                                            `json:"description,required"`
	EventType       string                                                                            `json:"eventType,required"`
	IsInstantaneous interface{}                                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                                         `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3LocationTargetListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3LocationTargetListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3LocationTargetListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3LocationTargetListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3LocationTargetListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3LocationTargetListResponseResultTop0 struct {
	Rank                float64                                                   `json:"rank,required"`
	TargetCountryAlpha2 string                                                    `json:"targetCountryAlpha2,required"`
	TargetCountryName   string                                                    `json:"targetCountryName,required"`
	Value               string                                                    `json:"value,required"`
	JSON                radarAttackLayer3LocationTargetListResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer3LocationTargetListResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarAttackLayer3LocationTargetListResponseResultTop0]
type radarAttackLayer3LocationTargetListResponseResultTop0JSON struct {
	Rank                apijson.Field
	TargetCountryAlpha2 apijson.Field
	TargetCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer3LocationTargetListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3LocationTargetListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3LocationTargetListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3LocationTargetListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3LocationTargetListParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3LocationTargetListParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3LocationTargetListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3LocationTargetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3LocationTargetListParamsDateRange string

const (
	RadarAttackLayer3LocationTargetListParamsDateRange1d         RadarAttackLayer3LocationTargetListParamsDateRange = "1d"
	RadarAttackLayer3LocationTargetListParamsDateRange2d         RadarAttackLayer3LocationTargetListParamsDateRange = "2d"
	RadarAttackLayer3LocationTargetListParamsDateRange7d         RadarAttackLayer3LocationTargetListParamsDateRange = "7d"
	RadarAttackLayer3LocationTargetListParamsDateRange14d        RadarAttackLayer3LocationTargetListParamsDateRange = "14d"
	RadarAttackLayer3LocationTargetListParamsDateRange28d        RadarAttackLayer3LocationTargetListParamsDateRange = "28d"
	RadarAttackLayer3LocationTargetListParamsDateRange12w        RadarAttackLayer3LocationTargetListParamsDateRange = "12w"
	RadarAttackLayer3LocationTargetListParamsDateRange24w        RadarAttackLayer3LocationTargetListParamsDateRange = "24w"
	RadarAttackLayer3LocationTargetListParamsDateRange52w        RadarAttackLayer3LocationTargetListParamsDateRange = "52w"
	RadarAttackLayer3LocationTargetListParamsDateRange1dControl  RadarAttackLayer3LocationTargetListParamsDateRange = "1dControl"
	RadarAttackLayer3LocationTargetListParamsDateRange2dControl  RadarAttackLayer3LocationTargetListParamsDateRange = "2dControl"
	RadarAttackLayer3LocationTargetListParamsDateRange7dControl  RadarAttackLayer3LocationTargetListParamsDateRange = "7dControl"
	RadarAttackLayer3LocationTargetListParamsDateRange14dControl RadarAttackLayer3LocationTargetListParamsDateRange = "14dControl"
	RadarAttackLayer3LocationTargetListParamsDateRange28dControl RadarAttackLayer3LocationTargetListParamsDateRange = "28dControl"
	RadarAttackLayer3LocationTargetListParamsDateRange12wControl RadarAttackLayer3LocationTargetListParamsDateRange = "12wControl"
	RadarAttackLayer3LocationTargetListParamsDateRange24wControl RadarAttackLayer3LocationTargetListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3LocationTargetListParamsFormat string

const (
	RadarAttackLayer3LocationTargetListParamsFormatJson RadarAttackLayer3LocationTargetListParamsFormat = "JSON"
	RadarAttackLayer3LocationTargetListParamsFormatCsv  RadarAttackLayer3LocationTargetListParamsFormat = "CSV"
)

type RadarAttackLayer3LocationTargetListParamsIPVersion string

const (
	RadarAttackLayer3LocationTargetListParamsIPVersionIPv4 RadarAttackLayer3LocationTargetListParamsIPVersion = "IPv4"
	RadarAttackLayer3LocationTargetListParamsIPVersionIPv6 RadarAttackLayer3LocationTargetListParamsIPVersion = "IPv6"
)

type RadarAttackLayer3LocationTargetListParamsProtocol string

const (
	RadarAttackLayer3LocationTargetListParamsProtocolUdp  RadarAttackLayer3LocationTargetListParamsProtocol = "UDP"
	RadarAttackLayer3LocationTargetListParamsProtocolTcp  RadarAttackLayer3LocationTargetListParamsProtocol = "TCP"
	RadarAttackLayer3LocationTargetListParamsProtocolIcmp RadarAttackLayer3LocationTargetListParamsProtocol = "ICMP"
	RadarAttackLayer3LocationTargetListParamsProtocolGre  RadarAttackLayer3LocationTargetListParamsProtocol = "GRE"
)
