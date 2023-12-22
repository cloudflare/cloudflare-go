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

// RadarAs112TopLocationIPVersionService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAs112TopLocationIPVersionService] method instead.
type RadarAs112TopLocationIPVersionService struct {
	Options []option.RequestOption
}

// NewRadarAs112TopLocationIPVersionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAs112TopLocationIPVersionService(opts ...option.RequestOption) (r *RadarAs112TopLocationIPVersionService) {
	r = &RadarAs112TopLocationIPVersionService{}
	r.Options = opts
	return
}

// Get the locations, by AS112 queries, of IP version.
func (r *RadarAs112TopLocationIPVersionService) Get(ctx context.Context, ipVersion RadarAs112TopLocationIPVersionGetParamsIPVersion, query RadarAs112TopLocationIPVersionGetParams, opts ...option.RequestOption) (res *RadarAs112TopLocationIPVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/as112/top/locations/ip_version/%v", ipVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TopLocationIPVersionGetResponse struct {
	Result  RadarAs112TopLocationIPVersionGetResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarAs112TopLocationIPVersionGetResponseJSON   `json:"-"`
}

// radarAs112TopLocationIPVersionGetResponseJSON contains the JSON metadata for the
// struct [RadarAs112TopLocationIPVersionGetResponse]
type radarAs112TopLocationIPVersionGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationIPVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationIPVersionGetResponseResult struct {
	Meta RadarAs112TopLocationIPVersionGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAs112TopLocationIPVersionGetResponseResultTop0 `json:"top_0,required"`
	JSON radarAs112TopLocationIPVersionGetResponseResultJSON   `json:"-"`
}

// radarAs112TopLocationIPVersionGetResponseResultJSON contains the JSON metadata
// for the struct [RadarAs112TopLocationIPVersionGetResponseResult]
type radarAs112TopLocationIPVersionGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationIPVersionGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationIPVersionGetResponseResultMeta struct {
	ConfidenceInfo RadarAs112TopLocationIPVersionGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarAs112TopLocationIPVersionGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarAs112TopLocationIPVersionGetResponseResultMetaJSON           `json:"-"`
}

// radarAs112TopLocationIPVersionGetResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAs112TopLocationIPVersionGetResponseResultMeta]
type radarAs112TopLocationIPVersionGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112TopLocationIPVersionGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationIPVersionGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112TopLocationIPVersionGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                         `json:"level,required"`
	JSON        radarAs112TopLocationIPVersionGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112TopLocationIPVersionGetResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAs112TopLocationIPVersionGetResponseResultMetaConfidenceInfo]
type radarAs112TopLocationIPVersionGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationIPVersionGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationIPVersionGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                          `json:"dataSource,required"`
	Description string                                                                          `json:"description,required"`
	EndTime     time.Time                                                                       `json:"endTime,required" format:"date-time"`
	EventType   string                                                                          `json:"eventType,required"`
	StartTime   time.Time                                                                       `json:"startTime,required" format:"date-time"`
	JSON        radarAs112TopLocationIPVersionGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112TopLocationIPVersionGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAs112TopLocationIPVersionGetResponseResultMetaConfidenceInfoAnnotation]
type radarAs112TopLocationIPVersionGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationIPVersionGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationIPVersionGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                                        `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarAs112TopLocationIPVersionGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112TopLocationIPVersionGetResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAs112TopLocationIPVersionGetResponseResultMetaDateRange]
type radarAs112TopLocationIPVersionGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationIPVersionGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationIPVersionGetResponseResultTop0 struct {
	ClientCountryAlpha2 string                                                  `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                  `json:"clientCountryName,required"`
	Value               string                                                  `json:"value,required"`
	JSON                radarAs112TopLocationIPVersionGetResponseResultTop0JSON `json:"-"`
}

// radarAs112TopLocationIPVersionGetResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarAs112TopLocationIPVersionGetResponseResultTop0]
type radarAs112TopLocationIPVersionGetResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAs112TopLocationIPVersionGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TopLocationIPVersionGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TopLocationIPVersionGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TopLocationIPVersionGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TopLocationIPVersionGetParams]'s query parameters
// as `url.Values`.
func (r RadarAs112TopLocationIPVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAs112TopLocationIPVersionGetParamsIPVersion string

const (
	RadarAs112TopLocationIPVersionGetParamsIPVersionIPv4 RadarAs112TopLocationIPVersionGetParamsIPVersion = "IPv4"
	RadarAs112TopLocationIPVersionGetParamsIPVersionIPv6 RadarAs112TopLocationIPVersionGetParamsIPVersion = "IPv6"
)

type RadarAs112TopLocationIPVersionGetParamsDateRange string

const (
	RadarAs112TopLocationIPVersionGetParamsDateRange1d         RadarAs112TopLocationIPVersionGetParamsDateRange = "1d"
	RadarAs112TopLocationIPVersionGetParamsDateRange7d         RadarAs112TopLocationIPVersionGetParamsDateRange = "7d"
	RadarAs112TopLocationIPVersionGetParamsDateRange14d        RadarAs112TopLocationIPVersionGetParamsDateRange = "14d"
	RadarAs112TopLocationIPVersionGetParamsDateRange28d        RadarAs112TopLocationIPVersionGetParamsDateRange = "28d"
	RadarAs112TopLocationIPVersionGetParamsDateRange12w        RadarAs112TopLocationIPVersionGetParamsDateRange = "12w"
	RadarAs112TopLocationIPVersionGetParamsDateRange24w        RadarAs112TopLocationIPVersionGetParamsDateRange = "24w"
	RadarAs112TopLocationIPVersionGetParamsDateRange52w        RadarAs112TopLocationIPVersionGetParamsDateRange = "52w"
	RadarAs112TopLocationIPVersionGetParamsDateRange1dControl  RadarAs112TopLocationIPVersionGetParamsDateRange = "1dControl"
	RadarAs112TopLocationIPVersionGetParamsDateRange7dControl  RadarAs112TopLocationIPVersionGetParamsDateRange = "7dControl"
	RadarAs112TopLocationIPVersionGetParamsDateRange14dControl RadarAs112TopLocationIPVersionGetParamsDateRange = "14dControl"
	RadarAs112TopLocationIPVersionGetParamsDateRange28dControl RadarAs112TopLocationIPVersionGetParamsDateRange = "28dControl"
	RadarAs112TopLocationIPVersionGetParamsDateRange12wControl RadarAs112TopLocationIPVersionGetParamsDateRange = "12wControl"
	RadarAs112TopLocationIPVersionGetParamsDateRange24wControl RadarAs112TopLocationIPVersionGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TopLocationIPVersionGetParamsFormat string

const (
	RadarAs112TopLocationIPVersionGetParamsFormatJson RadarAs112TopLocationIPVersionGetParamsFormat = "JSON"
	RadarAs112TopLocationIPVersionGetParamsFormatCsv  RadarAs112TopLocationIPVersionGetParamsFormat = "CSV"
)
