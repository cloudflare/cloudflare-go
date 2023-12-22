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

// RadarAs112TopLocationDnssecService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAs112TopLocationDnssecService] method instead.
type RadarAs112TopLocationDnssecService struct {
	Options []option.RequestOption
}

// NewRadarAs112TopLocationDnssecService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAs112TopLocationDnssecService(opts ...option.RequestOption) (r *RadarAs112TopLocationDnssecService) {
	r = &RadarAs112TopLocationDnssecService{}
	r.Options = opts
	return
}

// Get the locations, by AS112 queries, of DNSSEC validations.
func (r *RadarAs112TopLocationDnssecService) Get(ctx context.Context, dnssec RadarAs112TopLocationDnssecGetParamsDnssec, query RadarAs112TopLocationDnssecGetParams, opts ...option.RequestOption) (res *RadarAs112TopLocationDnssecGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/as112/top/locations/dnssec/%v", dnssec)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TopLocationDnssecGetResponse struct {
	Result  RadarAs112TopLocationDnssecGetResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarAs112TopLocationDnssecGetResponseJSON   `json:"-"`
}

// radarAs112TopLocationDnssecGetResponseJSON contains the JSON metadata for the
// struct [RadarAs112TopLocationDnssecGetResponse]
type radarAs112TopLocationDnssecGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationDnssecGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationDnssecGetResponseResult struct {
	Meta RadarAs112TopLocationDnssecGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAs112TopLocationDnssecGetResponseResultTop0 `json:"top_0,required"`
	JSON radarAs112TopLocationDnssecGetResponseResultJSON   `json:"-"`
}

// radarAs112TopLocationDnssecGetResponseResultJSON contains the JSON metadata for
// the struct [RadarAs112TopLocationDnssecGetResponseResult]
type radarAs112TopLocationDnssecGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationDnssecGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationDnssecGetResponseResultMeta struct {
	ConfidenceInfo RadarAs112TopLocationDnssecGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarAs112TopLocationDnssecGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarAs112TopLocationDnssecGetResponseResultMetaJSON           `json:"-"`
}

// radarAs112TopLocationDnssecGetResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAs112TopLocationDnssecGetResponseResultMeta]
type radarAs112TopLocationDnssecGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112TopLocationDnssecGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationDnssecGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112TopLocationDnssecGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                      `json:"level,required"`
	JSON        radarAs112TopLocationDnssecGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112TopLocationDnssecGetResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAs112TopLocationDnssecGetResponseResultMetaConfidenceInfo]
type radarAs112TopLocationDnssecGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationDnssecGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationDnssecGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                       `json:"dataSource,required"`
	Description string                                                                       `json:"description,required"`
	EndTime     time.Time                                                                    `json:"endTime,required" format:"date-time"`
	EventType   string                                                                       `json:"eventType,required"`
	StartTime   time.Time                                                                    `json:"startTime,required" format:"date-time"`
	JSON        radarAs112TopLocationDnssecGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112TopLocationDnssecGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAs112TopLocationDnssecGetResponseResultMetaConfidenceInfoAnnotation]
type radarAs112TopLocationDnssecGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationDnssecGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationDnssecGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                                     `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarAs112TopLocationDnssecGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112TopLocationDnssecGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAs112TopLocationDnssecGetResponseResultMetaDateRange]
type radarAs112TopLocationDnssecGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationDnssecGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationDnssecGetResponseResultTop0 struct {
	ClientCountryAlpha2 string                                               `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                               `json:"clientCountryName,required"`
	Value               string                                               `json:"value,required"`
	JSON                radarAs112TopLocationDnssecGetResponseResultTop0JSON `json:"-"`
}

// radarAs112TopLocationDnssecGetResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarAs112TopLocationDnssecGetResponseResultTop0]
type radarAs112TopLocationDnssecGetResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAs112TopLocationDnssecGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationDnssecGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TopLocationDnssecGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TopLocationDnssecGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TopLocationDnssecGetParams]'s query parameters as
// `url.Values`.
func (r RadarAs112TopLocationDnssecGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAs112TopLocationDnssecGetParamsDnssec string

const (
	RadarAs112TopLocationDnssecGetParamsDnssecSupported    RadarAs112TopLocationDnssecGetParamsDnssec = "SUPPORTED"
	RadarAs112TopLocationDnssecGetParamsDnssecNotSupported RadarAs112TopLocationDnssecGetParamsDnssec = "NOT_SUPPORTED"
)

type RadarAs112TopLocationDnssecGetParamsDateRange string

const (
	RadarAs112TopLocationDnssecGetParamsDateRange1d         RadarAs112TopLocationDnssecGetParamsDateRange = "1d"
	RadarAs112TopLocationDnssecGetParamsDateRange7d         RadarAs112TopLocationDnssecGetParamsDateRange = "7d"
	RadarAs112TopLocationDnssecGetParamsDateRange14d        RadarAs112TopLocationDnssecGetParamsDateRange = "14d"
	RadarAs112TopLocationDnssecGetParamsDateRange28d        RadarAs112TopLocationDnssecGetParamsDateRange = "28d"
	RadarAs112TopLocationDnssecGetParamsDateRange12w        RadarAs112TopLocationDnssecGetParamsDateRange = "12w"
	RadarAs112TopLocationDnssecGetParamsDateRange24w        RadarAs112TopLocationDnssecGetParamsDateRange = "24w"
	RadarAs112TopLocationDnssecGetParamsDateRange52w        RadarAs112TopLocationDnssecGetParamsDateRange = "52w"
	RadarAs112TopLocationDnssecGetParamsDateRange1dControl  RadarAs112TopLocationDnssecGetParamsDateRange = "1dControl"
	RadarAs112TopLocationDnssecGetParamsDateRange7dControl  RadarAs112TopLocationDnssecGetParamsDateRange = "7dControl"
	RadarAs112TopLocationDnssecGetParamsDateRange14dControl RadarAs112TopLocationDnssecGetParamsDateRange = "14dControl"
	RadarAs112TopLocationDnssecGetParamsDateRange28dControl RadarAs112TopLocationDnssecGetParamsDateRange = "28dControl"
	RadarAs112TopLocationDnssecGetParamsDateRange12wControl RadarAs112TopLocationDnssecGetParamsDateRange = "12wControl"
	RadarAs112TopLocationDnssecGetParamsDateRange24wControl RadarAs112TopLocationDnssecGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TopLocationDnssecGetParamsFormat string

const (
	RadarAs112TopLocationDnssecGetParamsFormatJson RadarAs112TopLocationDnssecGetParamsFormat = "JSON"
	RadarAs112TopLocationDnssecGetParamsFormatCsv  RadarAs112TopLocationDnssecGetParamsFormat = "CSV"
)
