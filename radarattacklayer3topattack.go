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

// RadarAttackLayer3TopAttackService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAttackLayer3TopAttackService] method instead.
type RadarAttackLayer3TopAttackService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3TopAttackService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3TopAttackService(opts ...option.RequestOption) (r *RadarAttackLayer3TopAttackService) {
	r = &RadarAttackLayer3TopAttackService{}
	r.Options = opts
	return
}

// Get the top attacks from origin to target location. Values are a percentage out
// of the total layer 3 attacks (with billing country). You can optionally limit
// the number of attacks per origin/target location (useful if all the top attacks
// are from or to the same location).
func (r *RadarAttackLayer3TopAttackService) List(ctx context.Context, query RadarAttackLayer3TopAttackListParams, opts ...option.RequestOption) (res *RadarAttackLayer3TopAttackListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/top/attacks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3TopAttackListResponse struct {
	Result  RadarAttackLayer3TopAttackListResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarAttackLayer3TopAttackListResponseJSON   `json:"-"`
}

// radarAttackLayer3TopAttackListResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3TopAttackListResponse]
type radarAttackLayer3TopAttackListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopAttackListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopAttackListResponseResult struct {
	Meta RadarAttackLayer3TopAttackListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3TopAttackListResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer3TopAttackListResponseResultJSON   `json:"-"`
}

// radarAttackLayer3TopAttackListResponseResultJSON contains the JSON metadata for
// the struct [RadarAttackLayer3TopAttackListResponseResult]
type radarAttackLayer3TopAttackListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopAttackListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopAttackListResponseResultMeta struct {
	DateRange      []RadarAttackLayer3TopAttackListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                         `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3TopAttackListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3TopAttackListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3TopAttackListResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TopAttackListResponseResultMeta]
type radarAttackLayer3TopAttackListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TopAttackListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopAttackListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TopAttackListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TopAttackListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TopAttackListResponseResultMetaDateRange]
type radarAttackLayer3TopAttackListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopAttackListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopAttackListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TopAttackListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                      `json:"level"`
	JSON        radarAttackLayer3TopAttackListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3TopAttackListResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TopAttackListResponseResultMetaConfidenceInfo]
type radarAttackLayer3TopAttackListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopAttackListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopAttackListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                       `json:"dataSource,required"`
	Description     string                                                                       `json:"description,required"`
	EventType       string                                                                       `json:"eventType,required"`
	IsInstantaneous interface{}                                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                                    `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3TopAttackListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TopAttackListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TopAttackListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3TopAttackListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TopAttackListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopAttackListResponseResultTop0 struct {
	OriginCountryAlpha2 string                                               `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                               `json:"originCountryName,required"`
	Value               string                                               `json:"value,required"`
	JSON                radarAttackLayer3TopAttackListResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer3TopAttackListResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarAttackLayer3TopAttackListResponseResultTop0]
type radarAttackLayer3TopAttackListResponseResultTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer3TopAttackListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TopAttackListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TopAttackListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TopAttackListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TopAttackListParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of attack origin/target location attack limits. Together with
	// `limitPerLocation`, limits how many objects will be fetched per origin/target
	// location.
	LimitDirection param.Field[RadarAttackLayer3TopAttackListParamsLimitDirection] `query:"limitDirection"`
	// Limit the number of attacks per origin/target (refer to `limitDirection`
	// parameter) location.
	LimitPerLocation param.Field[int64] `query:"limitPerLocation"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TopAttackListParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TopAttackListParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer3TopAttackListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3TopAttackListParamsDateRange string

const (
	RadarAttackLayer3TopAttackListParamsDateRange1d         RadarAttackLayer3TopAttackListParamsDateRange = "1d"
	RadarAttackLayer3TopAttackListParamsDateRange2d         RadarAttackLayer3TopAttackListParamsDateRange = "2d"
	RadarAttackLayer3TopAttackListParamsDateRange7d         RadarAttackLayer3TopAttackListParamsDateRange = "7d"
	RadarAttackLayer3TopAttackListParamsDateRange14d        RadarAttackLayer3TopAttackListParamsDateRange = "14d"
	RadarAttackLayer3TopAttackListParamsDateRange28d        RadarAttackLayer3TopAttackListParamsDateRange = "28d"
	RadarAttackLayer3TopAttackListParamsDateRange12w        RadarAttackLayer3TopAttackListParamsDateRange = "12w"
	RadarAttackLayer3TopAttackListParamsDateRange24w        RadarAttackLayer3TopAttackListParamsDateRange = "24w"
	RadarAttackLayer3TopAttackListParamsDateRange52w        RadarAttackLayer3TopAttackListParamsDateRange = "52w"
	RadarAttackLayer3TopAttackListParamsDateRange1dControl  RadarAttackLayer3TopAttackListParamsDateRange = "1dControl"
	RadarAttackLayer3TopAttackListParamsDateRange2dControl  RadarAttackLayer3TopAttackListParamsDateRange = "2dControl"
	RadarAttackLayer3TopAttackListParamsDateRange7dControl  RadarAttackLayer3TopAttackListParamsDateRange = "7dControl"
	RadarAttackLayer3TopAttackListParamsDateRange14dControl RadarAttackLayer3TopAttackListParamsDateRange = "14dControl"
	RadarAttackLayer3TopAttackListParamsDateRange28dControl RadarAttackLayer3TopAttackListParamsDateRange = "28dControl"
	RadarAttackLayer3TopAttackListParamsDateRange12wControl RadarAttackLayer3TopAttackListParamsDateRange = "12wControl"
	RadarAttackLayer3TopAttackListParamsDateRange24wControl RadarAttackLayer3TopAttackListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3TopAttackListParamsFormat string

const (
	RadarAttackLayer3TopAttackListParamsFormatJson RadarAttackLayer3TopAttackListParamsFormat = "JSON"
	RadarAttackLayer3TopAttackListParamsFormatCsv  RadarAttackLayer3TopAttackListParamsFormat = "CSV"
)

type RadarAttackLayer3TopAttackListParamsIPVersion string

const (
	RadarAttackLayer3TopAttackListParamsIPVersionIPv4 RadarAttackLayer3TopAttackListParamsIPVersion = "IPv4"
	RadarAttackLayer3TopAttackListParamsIPVersionIPv6 RadarAttackLayer3TopAttackListParamsIPVersion = "IPv6"
)

// Array of attack origin/target location attack limits. Together with
// `limitPerLocation`, limits how many objects will be fetched per origin/target
// location.
type RadarAttackLayer3TopAttackListParamsLimitDirection string

const (
	RadarAttackLayer3TopAttackListParamsLimitDirectionOrigin RadarAttackLayer3TopAttackListParamsLimitDirection = "ORIGIN"
	RadarAttackLayer3TopAttackListParamsLimitDirectionTarget RadarAttackLayer3TopAttackListParamsLimitDirection = "TARGET"
)

type RadarAttackLayer3TopAttackListParamsProtocol string

const (
	RadarAttackLayer3TopAttackListParamsProtocolUdp  RadarAttackLayer3TopAttackListParamsProtocol = "UDP"
	RadarAttackLayer3TopAttackListParamsProtocolTcp  RadarAttackLayer3TopAttackListParamsProtocol = "TCP"
	RadarAttackLayer3TopAttackListParamsProtocolIcmp RadarAttackLayer3TopAttackListParamsProtocol = "ICMP"
	RadarAttackLayer3TopAttackListParamsProtocolGre  RadarAttackLayer3TopAttackListParamsProtocol = "GRE"
)
