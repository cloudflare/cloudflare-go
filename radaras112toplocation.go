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

// RadarAs112TopLocationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAs112TopLocationService]
// method instead.
type RadarAs112TopLocationService struct {
	Options    []option.RequestOption
	Dnssecs    *RadarAs112TopLocationDnssecService
	Edns       *RadarAs112TopLocationEdnService
	IPVersions *RadarAs112TopLocationIPVersionService
}

// NewRadarAs112TopLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAs112TopLocationService(opts ...option.RequestOption) (r *RadarAs112TopLocationService) {
	r = &RadarAs112TopLocationService{}
	r.Options = opts
	r.Dnssecs = NewRadarAs112TopLocationDnssecService(opts...)
	r.Edns = NewRadarAs112TopLocationEdnService(opts...)
	r.IPVersions = NewRadarAs112TopLocationIPVersionService(opts...)
	return
}

// Get the top locations by HTTP traffic. Values are a percentage out of the total
// traffic.
func (r *RadarAs112TopLocationService) List(ctx context.Context, query RadarAs112TopLocationListParams, opts ...option.RequestOption) (res *RadarAs112TopLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TopLocationListResponse struct {
	Result  RadarAs112TopLocationListResponseResult `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarAs112TopLocationListResponseJSON   `json:"-"`
}

// radarAs112TopLocationListResponseJSON contains the JSON metadata for the struct
// [RadarAs112TopLocationListResponse]
type radarAs112TopLocationListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationListResponseResult struct {
	Meta RadarAs112TopLocationListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAs112TopLocationListResponseResultTop0 `json:"top_0,required"`
	JSON radarAs112TopLocationListResponseResultJSON   `json:"-"`
}

// radarAs112TopLocationListResponseResultJSON contains the JSON metadata for the
// struct [RadarAs112TopLocationListResponseResult]
type radarAs112TopLocationListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationListResponseResultMeta struct {
	ConfidenceInfo RadarAs112TopLocationListResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarAs112TopLocationListResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarAs112TopLocationListResponseResultMetaJSON           `json:"-"`
}

// radarAs112TopLocationListResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarAs112TopLocationListResponseResultMeta]
type radarAs112TopLocationListResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112TopLocationListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112TopLocationListResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                 `json:"level,required"`
	JSON        radarAs112TopLocationListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112TopLocationListResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAs112TopLocationListResponseResultMetaConfidenceInfo]
type radarAs112TopLocationListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                  `json:"dataSource,required"`
	Description string                                                                  `json:"description,required"`
	EndTime     time.Time                                                               `json:"endTime,required" format:"date-time"`
	EventType   string                                                                  `json:"eventType,required"`
	StartTime   time.Time                                                               `json:"startTime,required" format:"date-time"`
	JSON        radarAs112TopLocationListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112TopLocationListResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarAs112TopLocationListResponseResultMetaConfidenceInfoAnnotation]
type radarAs112TopLocationListResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationListResponseResultMetaDateRange struct {
	EndTime   time.Time                                                `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      radarAs112TopLocationListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112TopLocationListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAs112TopLocationListResponseResultMetaDateRange]
type radarAs112TopLocationListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationListResponseResultTop0 struct {
	ClientCountryAlpha2 string                                          `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                          `json:"clientCountryName,required"`
	Value               string                                          `json:"value,required"`
	JSON                radarAs112TopLocationListResponseResultTop0JSON `json:"-"`
}

// radarAs112TopLocationListResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarAs112TopLocationListResponseResultTop0]
type radarAs112TopLocationListResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAs112TopLocationListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TopLocationListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TopLocationListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TopLocationListParams]'s query parameters as
// `url.Values`.
func (r RadarAs112TopLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAs112TopLocationListParamsDateRange string

const (
	RadarAs112TopLocationListParamsDateRange1d         RadarAs112TopLocationListParamsDateRange = "1d"
	RadarAs112TopLocationListParamsDateRange7d         RadarAs112TopLocationListParamsDateRange = "7d"
	RadarAs112TopLocationListParamsDateRange14d        RadarAs112TopLocationListParamsDateRange = "14d"
	RadarAs112TopLocationListParamsDateRange28d        RadarAs112TopLocationListParamsDateRange = "28d"
	RadarAs112TopLocationListParamsDateRange12w        RadarAs112TopLocationListParamsDateRange = "12w"
	RadarAs112TopLocationListParamsDateRange24w        RadarAs112TopLocationListParamsDateRange = "24w"
	RadarAs112TopLocationListParamsDateRange52w        RadarAs112TopLocationListParamsDateRange = "52w"
	RadarAs112TopLocationListParamsDateRange1dControl  RadarAs112TopLocationListParamsDateRange = "1dControl"
	RadarAs112TopLocationListParamsDateRange7dControl  RadarAs112TopLocationListParamsDateRange = "7dControl"
	RadarAs112TopLocationListParamsDateRange14dControl RadarAs112TopLocationListParamsDateRange = "14dControl"
	RadarAs112TopLocationListParamsDateRange28dControl RadarAs112TopLocationListParamsDateRange = "28dControl"
	RadarAs112TopLocationListParamsDateRange12wControl RadarAs112TopLocationListParamsDateRange = "12wControl"
	RadarAs112TopLocationListParamsDateRange24wControl RadarAs112TopLocationListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TopLocationListParamsFormat string

const (
	RadarAs112TopLocationListParamsFormatJson RadarAs112TopLocationListParamsFormat = "JSON"
	RadarAs112TopLocationListParamsFormatCsv  RadarAs112TopLocationListParamsFormat = "CSV"
)
