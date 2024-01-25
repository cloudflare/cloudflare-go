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

// RadarAttackLayer3AttackService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAttackLayer3AttackService] method instead.
type RadarAttackLayer3AttackService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3AttackService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3AttackService(opts ...option.RequestOption) (r *RadarAttackLayer3AttackService) {
	r = &RadarAttackLayer3AttackService{}
	r.Options = opts
	return
}

// Get the top attacks from origin to target location. Values are a percentage out
// of the total layer 3 attacks (with billing country). You can optionally limit
// the number of attacks per origin/target location (useful if all the top attacks
// are from or to the same location).
func (r *RadarAttackLayer3AttackService) List(ctx context.Context, query RadarAttackLayer3AttackListParams, opts ...option.RequestOption) (res *RadarAttackLayer3AttackListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/top/attacks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3AttackListResponse struct {
	Result  RadarAttackLayer3AttackListResponseResult `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarAttackLayer3AttackListResponseJSON   `json:"-"`
}

// radarAttackLayer3AttackListResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3AttackListResponse]
type radarAttackLayer3AttackListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3AttackListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3AttackListResponseResult struct {
	Meta RadarAttackLayer3AttackListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3AttackListResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer3AttackListResponseResultJSON   `json:"-"`
}

// radarAttackLayer3AttackListResponseResultJSON contains the JSON metadata for the
// struct [RadarAttackLayer3AttackListResponseResult]
type radarAttackLayer3AttackListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3AttackListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3AttackListResponseResultMeta struct {
	DateRange      []RadarAttackLayer3AttackListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                      `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3AttackListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3AttackListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3AttackListResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarAttackLayer3AttackListResponseResultMeta]
type radarAttackLayer3AttackListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3AttackListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3AttackListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3AttackListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3AttackListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAttackLayer3AttackListResponseResultMetaDateRange]
type radarAttackLayer3AttackListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3AttackListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3AttackListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3AttackListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                   `json:"level"`
	JSON        radarAttackLayer3AttackListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3AttackListResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3AttackListResponseResultMetaConfidenceInfo]
type radarAttackLayer3AttackListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3AttackListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3AttackListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                    `json:"dataSource,required"`
	Description     string                                                                    `json:"description,required"`
	EventType       string                                                                    `json:"eventType,required"`
	IsInstantaneous interface{}                                                               `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                                 `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3AttackListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3AttackListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3AttackListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3AttackListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3AttackListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3AttackListResponseResultTop0 struct {
	OriginCountryAlpha2 string                                            `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                            `json:"originCountryName,required"`
	Value               string                                            `json:"value,required"`
	JSON                radarAttackLayer3AttackListResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer3AttackListResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarAttackLayer3AttackListResponseResultTop0]
type radarAttackLayer3AttackListResponseResultTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer3AttackListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3AttackListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3AttackListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3AttackListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3AttackListParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of attack origin/target location attack limits. Together with
	// `limitPerLocation`, limits how many objects will be fetched per origin/target
	// location.
	LimitDirection param.Field[RadarAttackLayer3AttackListParamsLimitDirection] `query:"limitDirection"`
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
	Protocol param.Field[[]RadarAttackLayer3AttackListParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3AttackListParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer3AttackListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer3AttackListParamsDateRange string

const (
	RadarAttackLayer3AttackListParamsDateRange1d         RadarAttackLayer3AttackListParamsDateRange = "1d"
	RadarAttackLayer3AttackListParamsDateRange2d         RadarAttackLayer3AttackListParamsDateRange = "2d"
	RadarAttackLayer3AttackListParamsDateRange7d         RadarAttackLayer3AttackListParamsDateRange = "7d"
	RadarAttackLayer3AttackListParamsDateRange14d        RadarAttackLayer3AttackListParamsDateRange = "14d"
	RadarAttackLayer3AttackListParamsDateRange28d        RadarAttackLayer3AttackListParamsDateRange = "28d"
	RadarAttackLayer3AttackListParamsDateRange12w        RadarAttackLayer3AttackListParamsDateRange = "12w"
	RadarAttackLayer3AttackListParamsDateRange24w        RadarAttackLayer3AttackListParamsDateRange = "24w"
	RadarAttackLayer3AttackListParamsDateRange52w        RadarAttackLayer3AttackListParamsDateRange = "52w"
	RadarAttackLayer3AttackListParamsDateRange1dControl  RadarAttackLayer3AttackListParamsDateRange = "1dControl"
	RadarAttackLayer3AttackListParamsDateRange2dControl  RadarAttackLayer3AttackListParamsDateRange = "2dControl"
	RadarAttackLayer3AttackListParamsDateRange7dControl  RadarAttackLayer3AttackListParamsDateRange = "7dControl"
	RadarAttackLayer3AttackListParamsDateRange14dControl RadarAttackLayer3AttackListParamsDateRange = "14dControl"
	RadarAttackLayer3AttackListParamsDateRange28dControl RadarAttackLayer3AttackListParamsDateRange = "28dControl"
	RadarAttackLayer3AttackListParamsDateRange12wControl RadarAttackLayer3AttackListParamsDateRange = "12wControl"
	RadarAttackLayer3AttackListParamsDateRange24wControl RadarAttackLayer3AttackListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3AttackListParamsFormat string

const (
	RadarAttackLayer3AttackListParamsFormatJson RadarAttackLayer3AttackListParamsFormat = "JSON"
	RadarAttackLayer3AttackListParamsFormatCsv  RadarAttackLayer3AttackListParamsFormat = "CSV"
)

type RadarAttackLayer3AttackListParamsIPVersion string

const (
	RadarAttackLayer3AttackListParamsIPVersionIPv4 RadarAttackLayer3AttackListParamsIPVersion = "IPv4"
	RadarAttackLayer3AttackListParamsIPVersionIPv6 RadarAttackLayer3AttackListParamsIPVersion = "IPv6"
)

// Array of attack origin/target location attack limits. Together with
// `limitPerLocation`, limits how many objects will be fetched per origin/target
// location.
type RadarAttackLayer3AttackListParamsLimitDirection string

const (
	RadarAttackLayer3AttackListParamsLimitDirectionOrigin RadarAttackLayer3AttackListParamsLimitDirection = "ORIGIN"
	RadarAttackLayer3AttackListParamsLimitDirectionTarget RadarAttackLayer3AttackListParamsLimitDirection = "TARGET"
)

type RadarAttackLayer3AttackListParamsProtocol string

const (
	RadarAttackLayer3AttackListParamsProtocolUdp  RadarAttackLayer3AttackListParamsProtocol = "UDP"
	RadarAttackLayer3AttackListParamsProtocolTcp  RadarAttackLayer3AttackListParamsProtocol = "TCP"
	RadarAttackLayer3AttackListParamsProtocolIcmp RadarAttackLayer3AttackListParamsProtocol = "ICMP"
	RadarAttackLayer3AttackListParamsProtocolGre  RadarAttackLayer3AttackListParamsProtocol = "GRE"
)
