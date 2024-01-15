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

// RadarAttackLayer7TopLocationOriginService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer7TopLocationOriginService] method instead.
type RadarAttackLayer7TopLocationOriginService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7TopLocationOriginService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarAttackLayer7TopLocationOriginService(opts ...option.RequestOption) (r *RadarAttackLayer7TopLocationOriginService) {
	r = &RadarAttackLayer7TopLocationOriginService{}
	r.Options = opts
	return
}

// Get the top origin locations of and by layer 7 attacks. Values are a percentage
// out of the total layer 7 attacks. The origin location is determined by the
// client IP.
func (r *RadarAttackLayer7TopLocationOriginService) List(ctx context.Context, query RadarAttackLayer7TopLocationOriginListParams, opts ...option.RequestOption) (res *RadarAttackLayer7TopLocationOriginListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/locations/origin"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7TopLocationOriginListResponse struct {
	Result  RadarAttackLayer7TopLocationOriginListResponseResult `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarAttackLayer7TopLocationOriginListResponseJSON   `json:"-"`
}

// radarAttackLayer7TopLocationOriginListResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TopLocationOriginListResponse]
type radarAttackLayer7TopLocationOriginListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationOriginListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationOriginListResponseResult struct {
	Meta RadarAttackLayer7TopLocationOriginListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer7TopLocationOriginListResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer7TopLocationOriginListResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TopLocationOriginListResponseResultJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TopLocationOriginListResponseResult]
type radarAttackLayer7TopLocationOriginListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationOriginListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationOriginListResponseResultMeta struct {
	DateRange      []RadarAttackLayer7TopLocationOriginListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                 `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer7TopLocationOriginListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7TopLocationOriginListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7TopLocationOriginListResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TopLocationOriginListResponseResultMeta]
type radarAttackLayer7TopLocationOriginListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationOriginListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationOriginListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                             `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TopLocationOriginListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TopLocationOriginListResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7TopLocationOriginListResponseResultMetaDateRange]
type radarAttackLayer7TopLocationOriginListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationOriginListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationOriginListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TopLocationOriginListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                              `json:"level"`
	JSON        radarAttackLayer7TopLocationOriginListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7TopLocationOriginListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TopLocationOriginListResponseResultMetaConfidenceInfo]
type radarAttackLayer7TopLocationOriginListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationOriginListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationOriginListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                               `json:"dataSource,required"`
	Description     string                                                                               `json:"description,required"`
	EventType       string                                                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                                            `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7TopLocationOriginListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TopLocationOriginListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TopLocationOriginListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7TopLocationOriginListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TopLocationOriginListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationOriginListResponseResultTop0 struct {
	OriginCountryAlpha2 string                                                       `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                                       `json:"originCountryName,required"`
	Rank                float64                                                      `json:"rank,required"`
	Value               string                                                       `json:"value,required"`
	JSON                radarAttackLayer7TopLocationOriginListResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer7TopLocationOriginListResponseResultTop0JSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TopLocationOriginListResponseResultTop0]
type radarAttackLayer7TopLocationOriginListResponseResultTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	Rank                apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationOriginListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationOriginListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TopLocationOriginListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TopLocationOriginListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7TopLocationOriginListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TopLocationOriginListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7TopLocationOriginListParamsDateRange string

const (
	RadarAttackLayer7TopLocationOriginListParamsDateRange1d         RadarAttackLayer7TopLocationOriginListParamsDateRange = "1d"
	RadarAttackLayer7TopLocationOriginListParamsDateRange2d         RadarAttackLayer7TopLocationOriginListParamsDateRange = "2d"
	RadarAttackLayer7TopLocationOriginListParamsDateRange7d         RadarAttackLayer7TopLocationOriginListParamsDateRange = "7d"
	RadarAttackLayer7TopLocationOriginListParamsDateRange14d        RadarAttackLayer7TopLocationOriginListParamsDateRange = "14d"
	RadarAttackLayer7TopLocationOriginListParamsDateRange28d        RadarAttackLayer7TopLocationOriginListParamsDateRange = "28d"
	RadarAttackLayer7TopLocationOriginListParamsDateRange12w        RadarAttackLayer7TopLocationOriginListParamsDateRange = "12w"
	RadarAttackLayer7TopLocationOriginListParamsDateRange24w        RadarAttackLayer7TopLocationOriginListParamsDateRange = "24w"
	RadarAttackLayer7TopLocationOriginListParamsDateRange52w        RadarAttackLayer7TopLocationOriginListParamsDateRange = "52w"
	RadarAttackLayer7TopLocationOriginListParamsDateRange1dControl  RadarAttackLayer7TopLocationOriginListParamsDateRange = "1dControl"
	RadarAttackLayer7TopLocationOriginListParamsDateRange2dControl  RadarAttackLayer7TopLocationOriginListParamsDateRange = "2dControl"
	RadarAttackLayer7TopLocationOriginListParamsDateRange7dControl  RadarAttackLayer7TopLocationOriginListParamsDateRange = "7dControl"
	RadarAttackLayer7TopLocationOriginListParamsDateRange14dControl RadarAttackLayer7TopLocationOriginListParamsDateRange = "14dControl"
	RadarAttackLayer7TopLocationOriginListParamsDateRange28dControl RadarAttackLayer7TopLocationOriginListParamsDateRange = "28dControl"
	RadarAttackLayer7TopLocationOriginListParamsDateRange12wControl RadarAttackLayer7TopLocationOriginListParamsDateRange = "12wControl"
	RadarAttackLayer7TopLocationOriginListParamsDateRange24wControl RadarAttackLayer7TopLocationOriginListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TopLocationOriginListParamsFormat string

const (
	RadarAttackLayer7TopLocationOriginListParamsFormatJson RadarAttackLayer7TopLocationOriginListParamsFormat = "JSON"
	RadarAttackLayer7TopLocationOriginListParamsFormatCsv  RadarAttackLayer7TopLocationOriginListParamsFormat = "CSV"
)
