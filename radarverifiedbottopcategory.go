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

// RadarVerifiedBotTopCategoryService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarVerifiedBotTopCategoryService] method instead.
type RadarVerifiedBotTopCategoryService struct {
	Options []option.RequestOption
}

// NewRadarVerifiedBotTopCategoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarVerifiedBotTopCategoryService(opts ...option.RequestOption) (r *RadarVerifiedBotTopCategoryService) {
	r = &RadarVerifiedBotTopCategoryService{}
	r.Options = opts
	return
}

// Get top verified bot categories by HTTP requests, along with their corresponding
// percentage, over the total verified bot HTTP requests.
func (r *RadarVerifiedBotTopCategoryService) List(ctx context.Context, query RadarVerifiedBotTopCategoryListParams, opts ...option.RequestOption) (res *RadarVerifiedBotTopCategoryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/verified_bots/top/categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarVerifiedBotTopCategoryListResponse struct {
	Result  RadarVerifiedBotTopCategoryListResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarVerifiedBotTopCategoryListResponseJSON   `json:"-"`
}

// radarVerifiedBotTopCategoryListResponseJSON contains the JSON metadata for the
// struct [RadarVerifiedBotTopCategoryListResponse]
type radarVerifiedBotTopCategoryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoryListResponseResult struct {
	Meta RadarVerifiedBotTopCategoryListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarVerifiedBotTopCategoryListResponseResultTop0 `json:"top_0,required"`
	JSON radarVerifiedBotTopCategoryListResponseResultJSON   `json:"-"`
}

// radarVerifiedBotTopCategoryListResponseResultJSON contains the JSON metadata for
// the struct [RadarVerifiedBotTopCategoryListResponseResult]
type radarVerifiedBotTopCategoryListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoryListResponseResultMeta struct {
	DateRange      []RadarVerifiedBotTopCategoryListResponseResultMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarVerifiedBotTopCategoryListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarVerifiedBotTopCategoryListResponseResultMetaJSON           `json:"-"`
}

// radarVerifiedBotTopCategoryListResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarVerifiedBotTopCategoryListResponseResultMeta]
type radarVerifiedBotTopCategoryListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoryListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoryListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                      `json:"startTime,required" format:"date-time"`
	JSON      radarVerifiedBotTopCategoryListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarVerifiedBotTopCategoryListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarVerifiedBotTopCategoryListResponseResultMetaDateRange]
type radarVerifiedBotTopCategoryListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoryListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoryListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarVerifiedBotTopCategoryListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                       `json:"level"`
	JSON        radarVerifiedBotTopCategoryListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarVerifiedBotTopCategoryListResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarVerifiedBotTopCategoryListResponseResultMetaConfidenceInfo]
type radarVerifiedBotTopCategoryListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoryListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoryListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                        `json:"dataSource,required"`
	Description     string                                                                        `json:"description,required"`
	EventType       string                                                                        `json:"eventType,required"`
	IsInstantaneous interface{}                                                                   `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                     `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                        `json:"linkedUrl"`
	StartTime       time.Time                                                                     `json:"startTime" format:"date-time"`
	JSON            radarVerifiedBotTopCategoryListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarVerifiedBotTopCategoryListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarVerifiedBotTopCategoryListResponseResultMetaConfidenceInfoAnnotation]
type radarVerifiedBotTopCategoryListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarVerifiedBotTopCategoryListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoryListResponseResultTop0 struct {
	BotCategory string                                                `json:"botCategory,required"`
	Value       string                                                `json:"value,required"`
	JSON        radarVerifiedBotTopCategoryListResponseResultTop0JSON `json:"-"`
}

// radarVerifiedBotTopCategoryListResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarVerifiedBotTopCategoryListResponseResultTop0]
type radarVerifiedBotTopCategoryListResponseResultTop0JSON struct {
	BotCategory apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoryListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoryListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarVerifiedBotTopCategoryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarVerifiedBotTopCategoryListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarVerifiedBotTopCategoryListParams]'s query parameters
// as `url.Values`.
func (r RadarVerifiedBotTopCategoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarVerifiedBotTopCategoryListParamsDateRange string

const (
	RadarVerifiedBotTopCategoryListParamsDateRange1d         RadarVerifiedBotTopCategoryListParamsDateRange = "1d"
	RadarVerifiedBotTopCategoryListParamsDateRange2d         RadarVerifiedBotTopCategoryListParamsDateRange = "2d"
	RadarVerifiedBotTopCategoryListParamsDateRange7d         RadarVerifiedBotTopCategoryListParamsDateRange = "7d"
	RadarVerifiedBotTopCategoryListParamsDateRange14d        RadarVerifiedBotTopCategoryListParamsDateRange = "14d"
	RadarVerifiedBotTopCategoryListParamsDateRange28d        RadarVerifiedBotTopCategoryListParamsDateRange = "28d"
	RadarVerifiedBotTopCategoryListParamsDateRange12w        RadarVerifiedBotTopCategoryListParamsDateRange = "12w"
	RadarVerifiedBotTopCategoryListParamsDateRange24w        RadarVerifiedBotTopCategoryListParamsDateRange = "24w"
	RadarVerifiedBotTopCategoryListParamsDateRange52w        RadarVerifiedBotTopCategoryListParamsDateRange = "52w"
	RadarVerifiedBotTopCategoryListParamsDateRange1dControl  RadarVerifiedBotTopCategoryListParamsDateRange = "1dControl"
	RadarVerifiedBotTopCategoryListParamsDateRange2dControl  RadarVerifiedBotTopCategoryListParamsDateRange = "2dControl"
	RadarVerifiedBotTopCategoryListParamsDateRange7dControl  RadarVerifiedBotTopCategoryListParamsDateRange = "7dControl"
	RadarVerifiedBotTopCategoryListParamsDateRange14dControl RadarVerifiedBotTopCategoryListParamsDateRange = "14dControl"
	RadarVerifiedBotTopCategoryListParamsDateRange28dControl RadarVerifiedBotTopCategoryListParamsDateRange = "28dControl"
	RadarVerifiedBotTopCategoryListParamsDateRange12wControl RadarVerifiedBotTopCategoryListParamsDateRange = "12wControl"
	RadarVerifiedBotTopCategoryListParamsDateRange24wControl RadarVerifiedBotTopCategoryListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarVerifiedBotTopCategoryListParamsFormat string

const (
	RadarVerifiedBotTopCategoryListParamsFormatJson RadarVerifiedBotTopCategoryListParamsFormat = "JSON"
	RadarVerifiedBotTopCategoryListParamsFormatCsv  RadarVerifiedBotTopCategoryListParamsFormat = "CSV"
)
