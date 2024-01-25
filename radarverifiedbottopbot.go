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

// RadarVerifiedBotTopBotService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarVerifiedBotTopBotService]
// method instead.
type RadarVerifiedBotTopBotService struct {
	Options []option.RequestOption
}

// NewRadarVerifiedBotTopBotService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarVerifiedBotTopBotService(opts ...option.RequestOption) (r *RadarVerifiedBotTopBotService) {
	r = &RadarVerifiedBotTopBotService{}
	r.Options = opts
	return
}

// Get top verified bots by HTTP requests, with owner and category.
func (r *RadarVerifiedBotTopBotService) List(ctx context.Context, query RadarVerifiedBotTopBotListParams, opts ...option.RequestOption) (res *RadarVerifiedBotTopBotListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/verified_bots/top/bots"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarVerifiedBotTopBotListResponse struct {
	Result  RadarVerifiedBotTopBotListResponseResult `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarVerifiedBotTopBotListResponseJSON   `json:"-"`
}

// radarVerifiedBotTopBotListResponseJSON contains the JSON metadata for the struct
// [RadarVerifiedBotTopBotListResponse]
type radarVerifiedBotTopBotListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopBotListResponseResult struct {
	Meta RadarVerifiedBotTopBotListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarVerifiedBotTopBotListResponseResultTop0 `json:"top_0,required"`
	JSON radarVerifiedBotTopBotListResponseResultJSON   `json:"-"`
}

// radarVerifiedBotTopBotListResponseResultJSON contains the JSON metadata for the
// struct [RadarVerifiedBotTopBotListResponseResult]
type radarVerifiedBotTopBotListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopBotListResponseResultMeta struct {
	DateRange      []RadarVerifiedBotTopBotListResponseResultMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarVerifiedBotTopBotListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarVerifiedBotTopBotListResponseResultMetaJSON           `json:"-"`
}

// radarVerifiedBotTopBotListResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarVerifiedBotTopBotListResponseResultMeta]
type radarVerifiedBotTopBotListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopBotListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarVerifiedBotTopBotListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarVerifiedBotTopBotListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarVerifiedBotTopBotListResponseResultMetaDateRange]
type radarVerifiedBotTopBotListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopBotListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarVerifiedBotTopBotListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarVerifiedBotTopBotListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarVerifiedBotTopBotListResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarVerifiedBotTopBotListResponseResultMetaConfidenceInfo]
type radarVerifiedBotTopBotListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopBotListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous interface{}                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarVerifiedBotTopBotListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarVerifiedBotTopBotListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarVerifiedBotTopBotListResponseResultMetaConfidenceInfoAnnotation]
type radarVerifiedBotTopBotListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarVerifiedBotTopBotListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopBotListResponseResultTop0 struct {
	BotCategory string                                           `json:"botCategory,required"`
	BotName     string                                           `json:"botName,required"`
	BotOwner    string                                           `json:"botOwner,required"`
	Value       string                                           `json:"value,required"`
	JSON        radarVerifiedBotTopBotListResponseResultTop0JSON `json:"-"`
}

// radarVerifiedBotTopBotListResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarVerifiedBotTopBotListResponseResultTop0]
type radarVerifiedBotTopBotListResponseResultTop0JSON struct {
	BotCategory apijson.Field
	BotName     apijson.Field
	BotOwner    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopBotListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarVerifiedBotTopBotListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarVerifiedBotTopBotListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarVerifiedBotTopBotListParams]'s query parameters as
// `url.Values`.
func (r RadarVerifiedBotTopBotListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarVerifiedBotTopBotListParamsDateRange string

const (
	RadarVerifiedBotTopBotListParamsDateRange1d         RadarVerifiedBotTopBotListParamsDateRange = "1d"
	RadarVerifiedBotTopBotListParamsDateRange2d         RadarVerifiedBotTopBotListParamsDateRange = "2d"
	RadarVerifiedBotTopBotListParamsDateRange7d         RadarVerifiedBotTopBotListParamsDateRange = "7d"
	RadarVerifiedBotTopBotListParamsDateRange14d        RadarVerifiedBotTopBotListParamsDateRange = "14d"
	RadarVerifiedBotTopBotListParamsDateRange28d        RadarVerifiedBotTopBotListParamsDateRange = "28d"
	RadarVerifiedBotTopBotListParamsDateRange12w        RadarVerifiedBotTopBotListParamsDateRange = "12w"
	RadarVerifiedBotTopBotListParamsDateRange24w        RadarVerifiedBotTopBotListParamsDateRange = "24w"
	RadarVerifiedBotTopBotListParamsDateRange52w        RadarVerifiedBotTopBotListParamsDateRange = "52w"
	RadarVerifiedBotTopBotListParamsDateRange1dControl  RadarVerifiedBotTopBotListParamsDateRange = "1dControl"
	RadarVerifiedBotTopBotListParamsDateRange2dControl  RadarVerifiedBotTopBotListParamsDateRange = "2dControl"
	RadarVerifiedBotTopBotListParamsDateRange7dControl  RadarVerifiedBotTopBotListParamsDateRange = "7dControl"
	RadarVerifiedBotTopBotListParamsDateRange14dControl RadarVerifiedBotTopBotListParamsDateRange = "14dControl"
	RadarVerifiedBotTopBotListParamsDateRange28dControl RadarVerifiedBotTopBotListParamsDateRange = "28dControl"
	RadarVerifiedBotTopBotListParamsDateRange12wControl RadarVerifiedBotTopBotListParamsDateRange = "12wControl"
	RadarVerifiedBotTopBotListParamsDateRange24wControl RadarVerifiedBotTopBotListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarVerifiedBotTopBotListParamsFormat string

const (
	RadarVerifiedBotTopBotListParamsFormatJson RadarVerifiedBotTopBotListParamsFormat = "JSON"
	RadarVerifiedBotTopBotListParamsFormatCsv  RadarVerifiedBotTopBotListParamsFormat = "CSV"
)
