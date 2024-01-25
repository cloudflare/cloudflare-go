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

// RadarAttackLayer7TopLocationTargetService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer7TopLocationTargetService] method instead.
type RadarAttackLayer7TopLocationTargetService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7TopLocationTargetService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarAttackLayer7TopLocationTargetService(opts ...option.RequestOption) (r *RadarAttackLayer7TopLocationTargetService) {
	r = &RadarAttackLayer7TopLocationTargetService{}
	r.Options = opts
	return
}

// Get the top target locations of and by layer 7 attacks. Values are a percentage
// out of the total layer 7 attacks. The target location is determined by the
// attacked zone's billing country, when available.
func (r *RadarAttackLayer7TopLocationTargetService) List(ctx context.Context, query RadarAttackLayer7TopLocationTargetListParams, opts ...option.RequestOption) (res *RadarAttackLayer7TopLocationTargetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/locations/target"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7TopLocationTargetListResponse struct {
	Result  RadarAttackLayer7TopLocationTargetListResponseResult `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarAttackLayer7TopLocationTargetListResponseJSON   `json:"-"`
}

// radarAttackLayer7TopLocationTargetListResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TopLocationTargetListResponse]
type radarAttackLayer7TopLocationTargetListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationTargetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationTargetListResponseResult struct {
	Meta RadarAttackLayer7TopLocationTargetListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer7TopLocationTargetListResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer7TopLocationTargetListResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TopLocationTargetListResponseResultJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TopLocationTargetListResponseResult]
type radarAttackLayer7TopLocationTargetListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationTargetListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationTargetListResponseResultMeta struct {
	DateRange      []RadarAttackLayer7TopLocationTargetListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                 `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer7TopLocationTargetListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7TopLocationTargetListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7TopLocationTargetListResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TopLocationTargetListResponseResultMeta]
type radarAttackLayer7TopLocationTargetListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationTargetListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationTargetListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                             `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TopLocationTargetListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TopLocationTargetListResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7TopLocationTargetListResponseResultMetaDateRange]
type radarAttackLayer7TopLocationTargetListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationTargetListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationTargetListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TopLocationTargetListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                              `json:"level"`
	JSON        radarAttackLayer7TopLocationTargetListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7TopLocationTargetListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TopLocationTargetListResponseResultMetaConfidenceInfo]
type radarAttackLayer7TopLocationTargetListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationTargetListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationTargetListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                               `json:"dataSource,required"`
	Description     string                                                                               `json:"description,required"`
	EventType       string                                                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                                            `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7TopLocationTargetListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TopLocationTargetListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TopLocationTargetListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7TopLocationTargetListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TopLocationTargetListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationTargetListResponseResultTop0 struct {
	Rank                float64                                                      `json:"rank,required"`
	TargetCountryAlpha2 string                                                       `json:"targetCountryAlpha2,required"`
	TargetCountryName   string                                                       `json:"targetCountryName,required"`
	Value               string                                                       `json:"value,required"`
	JSON                radarAttackLayer7TopLocationTargetListResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer7TopLocationTargetListResponseResultTop0JSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TopLocationTargetListResponseResultTop0]
type radarAttackLayer7TopLocationTargetListResponseResultTop0JSON struct {
	Rank                apijson.Field
	TargetCountryAlpha2 apijson.Field
	TargetCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationTargetListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TopLocationTargetListParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TopLocationTargetListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TopLocationTargetListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7TopLocationTargetListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TopLocationTargetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7TopLocationTargetListParamsDateRange string

const (
	RadarAttackLayer7TopLocationTargetListParamsDateRange1d         RadarAttackLayer7TopLocationTargetListParamsDateRange = "1d"
	RadarAttackLayer7TopLocationTargetListParamsDateRange2d         RadarAttackLayer7TopLocationTargetListParamsDateRange = "2d"
	RadarAttackLayer7TopLocationTargetListParamsDateRange7d         RadarAttackLayer7TopLocationTargetListParamsDateRange = "7d"
	RadarAttackLayer7TopLocationTargetListParamsDateRange14d        RadarAttackLayer7TopLocationTargetListParamsDateRange = "14d"
	RadarAttackLayer7TopLocationTargetListParamsDateRange28d        RadarAttackLayer7TopLocationTargetListParamsDateRange = "28d"
	RadarAttackLayer7TopLocationTargetListParamsDateRange12w        RadarAttackLayer7TopLocationTargetListParamsDateRange = "12w"
	RadarAttackLayer7TopLocationTargetListParamsDateRange24w        RadarAttackLayer7TopLocationTargetListParamsDateRange = "24w"
	RadarAttackLayer7TopLocationTargetListParamsDateRange52w        RadarAttackLayer7TopLocationTargetListParamsDateRange = "52w"
	RadarAttackLayer7TopLocationTargetListParamsDateRange1dControl  RadarAttackLayer7TopLocationTargetListParamsDateRange = "1dControl"
	RadarAttackLayer7TopLocationTargetListParamsDateRange2dControl  RadarAttackLayer7TopLocationTargetListParamsDateRange = "2dControl"
	RadarAttackLayer7TopLocationTargetListParamsDateRange7dControl  RadarAttackLayer7TopLocationTargetListParamsDateRange = "7dControl"
	RadarAttackLayer7TopLocationTargetListParamsDateRange14dControl RadarAttackLayer7TopLocationTargetListParamsDateRange = "14dControl"
	RadarAttackLayer7TopLocationTargetListParamsDateRange28dControl RadarAttackLayer7TopLocationTargetListParamsDateRange = "28dControl"
	RadarAttackLayer7TopLocationTargetListParamsDateRange12wControl RadarAttackLayer7TopLocationTargetListParamsDateRange = "12wControl"
	RadarAttackLayer7TopLocationTargetListParamsDateRange24wControl RadarAttackLayer7TopLocationTargetListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7TopLocationTargetListParamsFormat string

const (
	RadarAttackLayer7TopLocationTargetListParamsFormatJson RadarAttackLayer7TopLocationTargetListParamsFormat = "JSON"
	RadarAttackLayer7TopLocationTargetListParamsFormatCsv  RadarAttackLayer7TopLocationTargetListParamsFormat = "CSV"
)
