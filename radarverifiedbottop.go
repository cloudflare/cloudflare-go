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

// RadarVerifiedBotTopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarVerifiedBotTopService]
// method instead.
type RadarVerifiedBotTopService struct {
	Options []option.RequestOption
}

// NewRadarVerifiedBotTopService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarVerifiedBotTopService(opts ...option.RequestOption) (r *RadarVerifiedBotTopService) {
	r = &RadarVerifiedBotTopService{}
	r.Options = opts
	return
}

// Get top verified bots by HTTP requests, with owner and category.
func (r *RadarVerifiedBotTopService) Bots(ctx context.Context, query RadarVerifiedBotTopBotsParams, opts ...option.RequestOption) (res *RadarVerifiedBotTopBotsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarVerifiedBotTopBotsResponseEnvelope
	path := "radar/verified_bots/top/bots"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get top verified bot categories by HTTP requests, along with their corresponding
// percentage, over the total verified bot HTTP requests.
func (r *RadarVerifiedBotTopService) Categories(ctx context.Context, query RadarVerifiedBotTopCategoriesParams, opts ...option.RequestOption) (res *RadarVerifiedBotTopCategoriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarVerifiedBotTopCategoriesResponseEnvelope
	path := "radar/verified_bots/top/categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarVerifiedBotTopBotsResponse struct {
	Meta RadarVerifiedBotTopBotsResponseMeta   `json:"meta,required"`
	Top0 []RadarVerifiedBotTopBotsResponseTop0 `json:"top_0,required"`
	JSON radarVerifiedBotTopBotsResponseJSON   `json:"-"`
}

// radarVerifiedBotTopBotsResponseJSON contains the JSON metadata for the struct
// [RadarVerifiedBotTopBotsResponse]
type radarVerifiedBotTopBotsResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopBotsResponseMeta struct {
	DateRange      []RadarVerifiedBotTopBotsResponseMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarVerifiedBotTopBotsResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarVerifiedBotTopBotsResponseMetaJSON           `json:"-"`
}

// radarVerifiedBotTopBotsResponseMetaJSON contains the JSON metadata for the
// struct [RadarVerifiedBotTopBotsResponseMeta]
type radarVerifiedBotTopBotsResponseMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopBotsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      radarVerifiedBotTopBotsResponseMetaDateRangeJSON `json:"-"`
}

// radarVerifiedBotTopBotsResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarVerifiedBotTopBotsResponseMetaDateRange]
type radarVerifiedBotTopBotsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopBotsResponseMetaConfidenceInfo struct {
	Annotations []RadarVerifiedBotTopBotsResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                         `json:"level"`
	JSON        radarVerifiedBotTopBotsResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarVerifiedBotTopBotsResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarVerifiedBotTopBotsResponseMetaConfidenceInfo]
type radarVerifiedBotTopBotsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopBotsResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                          `json:"dataSource,required"`
	Description     string                                                          `json:"description,required"`
	EventType       string                                                          `json:"eventType,required"`
	IsInstantaneous interface{}                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                       `json:"startTime" format:"date-time"`
	JSON            radarVerifiedBotTopBotsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarVerifiedBotTopBotsResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarVerifiedBotTopBotsResponseMetaConfidenceInfoAnnotation]
type radarVerifiedBotTopBotsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarVerifiedBotTopBotsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopBotsResponseTop0 struct {
	BotCategory string                                  `json:"botCategory,required"`
	BotName     string                                  `json:"botName,required"`
	BotOwner    string                                  `json:"botOwner,required"`
	Value       string                                  `json:"value,required"`
	JSON        radarVerifiedBotTopBotsResponseTop0JSON `json:"-"`
}

// radarVerifiedBotTopBotsResponseTop0JSON contains the JSON metadata for the
// struct [RadarVerifiedBotTopBotsResponseTop0]
type radarVerifiedBotTopBotsResponseTop0JSON struct {
	BotCategory apijson.Field
	BotName     apijson.Field
	BotOwner    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoriesResponse struct {
	Meta RadarVerifiedBotTopCategoriesResponseMeta   `json:"meta,required"`
	Top0 []RadarVerifiedBotTopCategoriesResponseTop0 `json:"top_0,required"`
	JSON radarVerifiedBotTopCategoriesResponseJSON   `json:"-"`
}

// radarVerifiedBotTopCategoriesResponseJSON contains the JSON metadata for the
// struct [RadarVerifiedBotTopCategoriesResponse]
type radarVerifiedBotTopCategoriesResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoriesResponseMeta struct {
	DateRange      []RadarVerifiedBotTopCategoriesResponseMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarVerifiedBotTopCategoriesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarVerifiedBotTopCategoriesResponseMetaJSON           `json:"-"`
}

// radarVerifiedBotTopCategoriesResponseMetaJSON contains the JSON metadata for the
// struct [RadarVerifiedBotTopCategoriesResponseMeta]
type radarVerifiedBotTopCategoriesResponseMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      radarVerifiedBotTopCategoriesResponseMetaDateRangeJSON `json:"-"`
}

// radarVerifiedBotTopCategoriesResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarVerifiedBotTopCategoriesResponseMetaDateRange]
type radarVerifiedBotTopCategoriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoriesResponseMetaConfidenceInfo struct {
	Annotations []RadarVerifiedBotTopCategoriesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                               `json:"level"`
	JSON        radarVerifiedBotTopCategoriesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarVerifiedBotTopCategoriesResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarVerifiedBotTopCategoriesResponseMetaConfidenceInfo]
type radarVerifiedBotTopCategoriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                `json:"dataSource,required"`
	Description     string                                                                `json:"description,required"`
	EventType       string                                                                `json:"eventType,required"`
	IsInstantaneous interface{}                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                             `json:"startTime" format:"date-time"`
	JSON            radarVerifiedBotTopCategoriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarVerifiedBotTopCategoriesResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarVerifiedBotTopCategoriesResponseMetaConfidenceInfoAnnotation]
type radarVerifiedBotTopCategoriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarVerifiedBotTopCategoriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoriesResponseTop0 struct {
	BotCategory string                                        `json:"botCategory,required"`
	Value       string                                        `json:"value,required"`
	JSON        radarVerifiedBotTopCategoriesResponseTop0JSON `json:"-"`
}

// radarVerifiedBotTopCategoriesResponseTop0JSON contains the JSON metadata for the
// struct [RadarVerifiedBotTopCategoriesResponseTop0]
type radarVerifiedBotTopCategoriesResponseTop0JSON struct {
	BotCategory apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoriesResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopBotsParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarVerifiedBotTopBotsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarVerifiedBotTopBotsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarVerifiedBotTopBotsParams]'s query parameters as
// `url.Values`.
func (r RadarVerifiedBotTopBotsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarVerifiedBotTopBotsParamsDateRange string

const (
	RadarVerifiedBotTopBotsParamsDateRange1d         RadarVerifiedBotTopBotsParamsDateRange = "1d"
	RadarVerifiedBotTopBotsParamsDateRange2d         RadarVerifiedBotTopBotsParamsDateRange = "2d"
	RadarVerifiedBotTopBotsParamsDateRange7d         RadarVerifiedBotTopBotsParamsDateRange = "7d"
	RadarVerifiedBotTopBotsParamsDateRange14d        RadarVerifiedBotTopBotsParamsDateRange = "14d"
	RadarVerifiedBotTopBotsParamsDateRange28d        RadarVerifiedBotTopBotsParamsDateRange = "28d"
	RadarVerifiedBotTopBotsParamsDateRange12w        RadarVerifiedBotTopBotsParamsDateRange = "12w"
	RadarVerifiedBotTopBotsParamsDateRange24w        RadarVerifiedBotTopBotsParamsDateRange = "24w"
	RadarVerifiedBotTopBotsParamsDateRange52w        RadarVerifiedBotTopBotsParamsDateRange = "52w"
	RadarVerifiedBotTopBotsParamsDateRange1dControl  RadarVerifiedBotTopBotsParamsDateRange = "1dControl"
	RadarVerifiedBotTopBotsParamsDateRange2dControl  RadarVerifiedBotTopBotsParamsDateRange = "2dControl"
	RadarVerifiedBotTopBotsParamsDateRange7dControl  RadarVerifiedBotTopBotsParamsDateRange = "7dControl"
	RadarVerifiedBotTopBotsParamsDateRange14dControl RadarVerifiedBotTopBotsParamsDateRange = "14dControl"
	RadarVerifiedBotTopBotsParamsDateRange28dControl RadarVerifiedBotTopBotsParamsDateRange = "28dControl"
	RadarVerifiedBotTopBotsParamsDateRange12wControl RadarVerifiedBotTopBotsParamsDateRange = "12wControl"
	RadarVerifiedBotTopBotsParamsDateRange24wControl RadarVerifiedBotTopBotsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarVerifiedBotTopBotsParamsFormat string

const (
	RadarVerifiedBotTopBotsParamsFormatJson RadarVerifiedBotTopBotsParamsFormat = "JSON"
	RadarVerifiedBotTopBotsParamsFormatCsv  RadarVerifiedBotTopBotsParamsFormat = "CSV"
)

type RadarVerifiedBotTopBotsResponseEnvelope struct {
	Result  RadarVerifiedBotTopBotsResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarVerifiedBotTopBotsResponseEnvelopeJSON `json:"-"`
}

// radarVerifiedBotTopBotsResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarVerifiedBotTopBotsResponseEnvelope]
type radarVerifiedBotTopBotsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarVerifiedBotTopCategoriesParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarVerifiedBotTopCategoriesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarVerifiedBotTopCategoriesParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarVerifiedBotTopCategoriesParams]'s query parameters as
// `url.Values`.
func (r RadarVerifiedBotTopCategoriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarVerifiedBotTopCategoriesParamsDateRange string

const (
	RadarVerifiedBotTopCategoriesParamsDateRange1d         RadarVerifiedBotTopCategoriesParamsDateRange = "1d"
	RadarVerifiedBotTopCategoriesParamsDateRange2d         RadarVerifiedBotTopCategoriesParamsDateRange = "2d"
	RadarVerifiedBotTopCategoriesParamsDateRange7d         RadarVerifiedBotTopCategoriesParamsDateRange = "7d"
	RadarVerifiedBotTopCategoriesParamsDateRange14d        RadarVerifiedBotTopCategoriesParamsDateRange = "14d"
	RadarVerifiedBotTopCategoriesParamsDateRange28d        RadarVerifiedBotTopCategoriesParamsDateRange = "28d"
	RadarVerifiedBotTopCategoriesParamsDateRange12w        RadarVerifiedBotTopCategoriesParamsDateRange = "12w"
	RadarVerifiedBotTopCategoriesParamsDateRange24w        RadarVerifiedBotTopCategoriesParamsDateRange = "24w"
	RadarVerifiedBotTopCategoriesParamsDateRange52w        RadarVerifiedBotTopCategoriesParamsDateRange = "52w"
	RadarVerifiedBotTopCategoriesParamsDateRange1dControl  RadarVerifiedBotTopCategoriesParamsDateRange = "1dControl"
	RadarVerifiedBotTopCategoriesParamsDateRange2dControl  RadarVerifiedBotTopCategoriesParamsDateRange = "2dControl"
	RadarVerifiedBotTopCategoriesParamsDateRange7dControl  RadarVerifiedBotTopCategoriesParamsDateRange = "7dControl"
	RadarVerifiedBotTopCategoriesParamsDateRange14dControl RadarVerifiedBotTopCategoriesParamsDateRange = "14dControl"
	RadarVerifiedBotTopCategoriesParamsDateRange28dControl RadarVerifiedBotTopCategoriesParamsDateRange = "28dControl"
	RadarVerifiedBotTopCategoriesParamsDateRange12wControl RadarVerifiedBotTopCategoriesParamsDateRange = "12wControl"
	RadarVerifiedBotTopCategoriesParamsDateRange24wControl RadarVerifiedBotTopCategoriesParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarVerifiedBotTopCategoriesParamsFormat string

const (
	RadarVerifiedBotTopCategoriesParamsFormatJson RadarVerifiedBotTopCategoriesParamsFormat = "JSON"
	RadarVerifiedBotTopCategoriesParamsFormatCsv  RadarVerifiedBotTopCategoriesParamsFormat = "CSV"
)

type RadarVerifiedBotTopCategoriesResponseEnvelope struct {
	Result  RadarVerifiedBotTopCategoriesResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarVerifiedBotTopCategoriesResponseEnvelopeJSON `json:"-"`
}

// radarVerifiedBotTopCategoriesResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarVerifiedBotTopCategoriesResponseEnvelope]
type radarVerifiedBotTopCategoriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
