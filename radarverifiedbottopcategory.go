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
	var env RadarVerifiedBotTopCategoryListResponseEnvelope
	path := "radar/verified_bots/top/categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarVerifiedBotTopCategoryListResponse struct {
	Meta RadarVerifiedBotTopCategoryListResponseMeta   `json:"meta,required"`
	Top0 []RadarVerifiedBotTopCategoryListResponseTop0 `json:"top_0,required"`
	JSON radarVerifiedBotTopCategoryListResponseJSON   `json:"-"`
}

// radarVerifiedBotTopCategoryListResponseJSON contains the JSON metadata for the
// struct [RadarVerifiedBotTopCategoryListResponse]
type radarVerifiedBotTopCategoryListResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoryListResponseMeta struct {
	DateRange      []RadarVerifiedBotTopCategoryListResponseMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarVerifiedBotTopCategoryListResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarVerifiedBotTopCategoryListResponseMetaJSON           `json:"-"`
}

// radarVerifiedBotTopCategoryListResponseMetaJSON contains the JSON metadata for
// the struct [RadarVerifiedBotTopCategoryListResponseMeta]
type radarVerifiedBotTopCategoryListResponseMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoryListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoryListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      radarVerifiedBotTopCategoryListResponseMetaDateRangeJSON `json:"-"`
}

// radarVerifiedBotTopCategoryListResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarVerifiedBotTopCategoryListResponseMetaDateRange]
type radarVerifiedBotTopCategoryListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoryListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoryListResponseMetaConfidenceInfo struct {
	Annotations []RadarVerifiedBotTopCategoryListResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                 `json:"level"`
	JSON        radarVerifiedBotTopCategoryListResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarVerifiedBotTopCategoryListResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarVerifiedBotTopCategoryListResponseMetaConfidenceInfo]
type radarVerifiedBotTopCategoryListResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoryListResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoryListResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                  `json:"dataSource,required"`
	Description     string                                                                  `json:"description,required"`
	EventType       string                                                                  `json:"eventType,required"`
	IsInstantaneous interface{}                                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                                               `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                  `json:"linkedUrl"`
	StartTime       time.Time                                                               `json:"startTime" format:"date-time"`
	JSON            radarVerifiedBotTopCategoryListResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarVerifiedBotTopCategoryListResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarVerifiedBotTopCategoryListResponseMetaConfidenceInfoAnnotation]
type radarVerifiedBotTopCategoryListResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarVerifiedBotTopCategoryListResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoryListResponseTop0 struct {
	BotCategory string                                          `json:"botCategory,required"`
	Value       string                                          `json:"value,required"`
	JSON        radarVerifiedBotTopCategoryListResponseTop0JSON `json:"-"`
}

// radarVerifiedBotTopCategoryListResponseTop0JSON contains the JSON metadata for
// the struct [RadarVerifiedBotTopCategoryListResponseTop0]
type radarVerifiedBotTopCategoryListResponseTop0JSON struct {
	BotCategory apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoryListResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoryListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
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

type RadarVerifiedBotTopCategoryListResponseEnvelope struct {
	Result  RadarVerifiedBotTopCategoryListResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarVerifiedBotTopCategoryListResponseEnvelopeJSON `json:"-"`
}

// radarVerifiedBotTopCategoryListResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarVerifiedBotTopCategoryListResponseEnvelope]
type radarVerifiedBotTopCategoryListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoryListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
