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

// RadarAttackLayer7TopAseOriginService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer7TopAseOriginService] method instead.
type RadarAttackLayer7TopAseOriginService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7TopAseOriginService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7TopAseOriginService(opts ...option.RequestOption) (r *RadarAttackLayer7TopAseOriginService) {
	r = &RadarAttackLayer7TopAseOriginService{}
	r.Options = opts
	return
}

// Get the top origin Autonomous Systems of and by layer 7 attacks. Values are a
// percentage out of the total layer 7 attacks. The origin Autonomous Systems is
// determined by the client IP.
func (r *RadarAttackLayer7TopAseOriginService) List(ctx context.Context, query RadarAttackLayer7TopAseOriginListParams, opts ...option.RequestOption) (res *RadarAttackLayer7TopAseOriginListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/ases/origin"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7TopAseOriginListResponse struct {
	Result  RadarAttackLayer7TopAseOriginListResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarAttackLayer7TopAseOriginListResponseJSON   `json:"-"`
}

// radarAttackLayer7TopAseOriginListResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7TopAseOriginListResponse]
type radarAttackLayer7TopAseOriginListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseOriginListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopAseOriginListResponseResult struct {
	Meta RadarAttackLayer7TopAseOriginListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer7TopAseOriginListResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer7TopAseOriginListResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TopAseOriginListResponseResultJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TopAseOriginListResponseResult]
type radarAttackLayer7TopAseOriginListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseOriginListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopAseOriginListResponseResultMeta struct {
	DateRange      []RadarAttackLayer7TopAseOriginListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                            `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer7TopAseOriginListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7TopAseOriginListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7TopAseOriginListResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TopAseOriginListResponseResultMeta]
type radarAttackLayer7TopAseOriginListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseOriginListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopAseOriginListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TopAseOriginListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TopAseOriginListResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TopAseOriginListResponseResultMetaDateRange]
type radarAttackLayer7TopAseOriginListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseOriginListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopAseOriginListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TopAseOriginListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                         `json:"level"`
	JSON        radarAttackLayer7TopAseOriginListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7TopAseOriginListResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7TopAseOriginListResponseResultMetaConfidenceInfo]
type radarAttackLayer7TopAseOriginListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseOriginListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopAseOriginListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                          `json:"dataSource,required"`
	Description     string                                                                          `json:"description,required"`
	EventType       string                                                                          `json:"eventType,required"`
	IsInstantaneous interface{}                                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                                       `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7TopAseOriginListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TopAseOriginListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TopAseOriginListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7TopAseOriginListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TopAseOriginListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopAseOriginListResponseResultTop0 struct {
	OriginASN     string                                                  `json:"originAsn,required"`
	OriginASNName string                                                  `json:"originAsnName,required"`
	Rank          float64                                                 `json:"rank,required"`
	Value         string                                                  `json:"value,required"`
	JSON          radarAttackLayer7TopAseOriginListResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer7TopAseOriginListResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarAttackLayer7TopAseOriginListResponseResultTop0]
type radarAttackLayer7TopAseOriginListResponseResultTop0JSON struct {
	OriginASN     apijson.Field
	OriginASNName apijson.Field
	Rank          apijson.Field
	Value         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseOriginListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopAseOriginListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TopAseOriginListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TopAseOriginListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7TopAseOriginListParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer7TopAseOriginListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7TopAseOriginListParamsDateRange string

const (
	RadarAttackLayer7TopAseOriginListParamsDateRange1d         RadarAttackLayer7TopAseOriginListParamsDateRange = "1d"
	RadarAttackLayer7TopAseOriginListParamsDateRange2d         RadarAttackLayer7TopAseOriginListParamsDateRange = "2d"
	RadarAttackLayer7TopAseOriginListParamsDateRange7d         RadarAttackLayer7TopAseOriginListParamsDateRange = "7d"
	RadarAttackLayer7TopAseOriginListParamsDateRange14d        RadarAttackLayer7TopAseOriginListParamsDateRange = "14d"
	RadarAttackLayer7TopAseOriginListParamsDateRange28d        RadarAttackLayer7TopAseOriginListParamsDateRange = "28d"
	RadarAttackLayer7TopAseOriginListParamsDateRange12w        RadarAttackLayer7TopAseOriginListParamsDateRange = "12w"
	RadarAttackLayer7TopAseOriginListParamsDateRange24w        RadarAttackLayer7TopAseOriginListParamsDateRange = "24w"
	RadarAttackLayer7TopAseOriginListParamsDateRange52w        RadarAttackLayer7TopAseOriginListParamsDateRange = "52w"
	RadarAttackLayer7TopAseOriginListParamsDateRange1dControl  RadarAttackLayer7TopAseOriginListParamsDateRange = "1dControl"
	RadarAttackLayer7TopAseOriginListParamsDateRange2dControl  RadarAttackLayer7TopAseOriginListParamsDateRange = "2dControl"
	RadarAttackLayer7TopAseOriginListParamsDateRange7dControl  RadarAttackLayer7TopAseOriginListParamsDateRange = "7dControl"
	RadarAttackLayer7TopAseOriginListParamsDateRange14dControl RadarAttackLayer7TopAseOriginListParamsDateRange = "14dControl"
	RadarAttackLayer7TopAseOriginListParamsDateRange28dControl RadarAttackLayer7TopAseOriginListParamsDateRange = "28dControl"
	RadarAttackLayer7TopAseOriginListParamsDateRange12wControl RadarAttackLayer7TopAseOriginListParamsDateRange = "12wControl"
	RadarAttackLayer7TopAseOriginListParamsDateRange24wControl RadarAttackLayer7TopAseOriginListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TopAseOriginListParamsFormat string

const (
	RadarAttackLayer7TopAseOriginListParamsFormatJson RadarAttackLayer7TopAseOriginListParamsFormat = "JSON"
	RadarAttackLayer7TopAseOriginListParamsFormatCsv  RadarAttackLayer7TopAseOriginListParamsFormat = "CSV"
)
