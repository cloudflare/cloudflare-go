// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarAs112TopLocationEdnService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAs112TopLocationEdnService] method instead.
type RadarAs112TopLocationEdnService struct {
	Options []option.RequestOption
}

// NewRadarAs112TopLocationEdnService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAs112TopLocationEdnService(opts ...option.RequestOption) (r *RadarAs112TopLocationEdnService) {
	r = &RadarAs112TopLocationEdnService{}
	r.Options = opts
	return
}

// Get the locations, by AS112 queries, of EDNS validations.
func (r *RadarAs112TopLocationEdnService) Get(ctx context.Context, edns RadarAs112TopLocationEdnGetParamsEdns, query RadarAs112TopLocationEdnGetParams, opts ...option.RequestOption) (res *RadarAs112TopLocationEdnGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/as112/top/locations/edns/%v", edns)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TopLocationEdnGetResponse struct {
	Result  RadarAs112TopLocationEdnGetResponseResult `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarAs112TopLocationEdnGetResponseJSON   `json:"-"`
}

// radarAs112TopLocationEdnGetResponseJSON contains the JSON metadata for the
// struct [RadarAs112TopLocationEdnGetResponse]
type radarAs112TopLocationEdnGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationEdnGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationEdnGetResponseResult struct {
	Meta RadarAs112TopLocationEdnGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAs112TopLocationEdnGetResponseResultTop0 `json:"top_0,required"`
	JSON radarAs112TopLocationEdnGetResponseResultJSON   `json:"-"`
}

// radarAs112TopLocationEdnGetResponseResultJSON contains the JSON metadata for the
// struct [RadarAs112TopLocationEdnGetResponseResult]
type radarAs112TopLocationEdnGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationEdnGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationEdnGetResponseResultMeta struct {
	ConfidenceInfo RadarAs112TopLocationEdnGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarAs112TopLocationEdnGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarAs112TopLocationEdnGetResponseResultMetaJSON           `json:"-"`
}

// radarAs112TopLocationEdnGetResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarAs112TopLocationEdnGetResponseResultMeta]
type radarAs112TopLocationEdnGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112TopLocationEdnGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationEdnGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112TopLocationEdnGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                   `json:"level,required"`
	JSON        radarAs112TopLocationEdnGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112TopLocationEdnGetResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAs112TopLocationEdnGetResponseResultMetaConfidenceInfo]
type radarAs112TopLocationEdnGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationEdnGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationEdnGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                    `json:"dataSource,required"`
	Description string                                                                    `json:"description,required"`
	EndTime     time.Time                                                                 `json:"endTime,required" format:"date-time"`
	EventType   string                                                                    `json:"eventType,required"`
	StartTime   time.Time                                                                 `json:"startTime,required" format:"date-time"`
	JSON        radarAs112TopLocationEdnGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112TopLocationEdnGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAs112TopLocationEdnGetResponseResultMetaConfidenceInfoAnnotation]
type radarAs112TopLocationEdnGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationEdnGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationEdnGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                                  `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarAs112TopLocationEdnGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112TopLocationEdnGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAs112TopLocationEdnGetResponseResultMetaDateRange]
type radarAs112TopLocationEdnGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationEdnGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationEdnGetResponseResultTop0 struct {
	ClientCountryAlpha2 string                                            `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                            `json:"clientCountryName,required"`
	Value               string                                            `json:"value,required"`
	JSON                radarAs112TopLocationEdnGetResponseResultTop0JSON `json:"-"`
}

// radarAs112TopLocationEdnGetResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarAs112TopLocationEdnGetResponseResultTop0]
type radarAs112TopLocationEdnGetResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAs112TopLocationEdnGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationEdnGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TopLocationEdnGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TopLocationEdnGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TopLocationEdnGetParams]'s query parameters as
// `url.Values`.
func (r RadarAs112TopLocationEdnGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAs112TopLocationEdnGetParamsEdns string

const (
	RadarAs112TopLocationEdnGetParamsEdnsSupported    RadarAs112TopLocationEdnGetParamsEdns = "SUPPORTED"
	RadarAs112TopLocationEdnGetParamsEdnsNotSupported RadarAs112TopLocationEdnGetParamsEdns = "NOT_SUPPORTED"
)

type RadarAs112TopLocationEdnGetParamsDateRange string

const (
	RadarAs112TopLocationEdnGetParamsDateRange1d         RadarAs112TopLocationEdnGetParamsDateRange = "1d"
	RadarAs112TopLocationEdnGetParamsDateRange7d         RadarAs112TopLocationEdnGetParamsDateRange = "7d"
	RadarAs112TopLocationEdnGetParamsDateRange14d        RadarAs112TopLocationEdnGetParamsDateRange = "14d"
	RadarAs112TopLocationEdnGetParamsDateRange28d        RadarAs112TopLocationEdnGetParamsDateRange = "28d"
	RadarAs112TopLocationEdnGetParamsDateRange12w        RadarAs112TopLocationEdnGetParamsDateRange = "12w"
	RadarAs112TopLocationEdnGetParamsDateRange24w        RadarAs112TopLocationEdnGetParamsDateRange = "24w"
	RadarAs112TopLocationEdnGetParamsDateRange52w        RadarAs112TopLocationEdnGetParamsDateRange = "52w"
	RadarAs112TopLocationEdnGetParamsDateRange1dControl  RadarAs112TopLocationEdnGetParamsDateRange = "1dControl"
	RadarAs112TopLocationEdnGetParamsDateRange7dControl  RadarAs112TopLocationEdnGetParamsDateRange = "7dControl"
	RadarAs112TopLocationEdnGetParamsDateRange14dControl RadarAs112TopLocationEdnGetParamsDateRange = "14dControl"
	RadarAs112TopLocationEdnGetParamsDateRange28dControl RadarAs112TopLocationEdnGetParamsDateRange = "28dControl"
	RadarAs112TopLocationEdnGetParamsDateRange12wControl RadarAs112TopLocationEdnGetParamsDateRange = "12wControl"
	RadarAs112TopLocationEdnGetParamsDateRange24wControl RadarAs112TopLocationEdnGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TopLocationEdnGetParamsFormat string

const (
	RadarAs112TopLocationEdnGetParamsFormatJson RadarAs112TopLocationEdnGetParamsFormat = "JSON"
	RadarAs112TopLocationEdnGetParamsFormatCsv  RadarAs112TopLocationEdnGetParamsFormat = "CSV"
)
