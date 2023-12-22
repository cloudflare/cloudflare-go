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

// RadarDNSTopLocationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarDNSTopLocationService]
// method instead.
type RadarDNSTopLocationService struct {
	Options []option.RequestOption
}

// NewRadarDNSTopLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarDNSTopLocationService(opts ...option.RequestOption) (r *RadarDNSTopLocationService) {
	r = &RadarDNSTopLocationService{}
	r.Options = opts
	return
}

// Get top locations by DNS queries made to Cloudflare's public DNS resolver.
func (r *RadarDNSTopLocationService) List(ctx context.Context, query RadarDNSTopLocationListParams, opts ...option.RequestOption) (res *RadarDNSTopLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarDNSTopLocationListResponse struct {
	Result  RadarDNSTopLocationListResponseResult `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarDNSTopLocationListResponseJSON   `json:"-"`
}

// radarDNSTopLocationListResponseJSON contains the JSON metadata for the struct
// [RadarDNSTopLocationListResponse]
type radarDNSTopLocationListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopLocationListResponseResult struct {
	Meta RadarDNSTopLocationListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarDNSTopLocationListResponseResultTop0 `json:"top_0,required"`
	JSON radarDNSTopLocationListResponseResultJSON   `json:"-"`
}

// radarDNSTopLocationListResponseResultJSON contains the JSON metadata for the
// struct [RadarDNSTopLocationListResponseResult]
type radarDNSTopLocationListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopLocationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopLocationListResponseResultMeta struct {
	ConfidenceInfo RadarDNSTopLocationListResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarDNSTopLocationListResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarDNSTopLocationListResponseResultMetaJSON           `json:"-"`
}

// radarDNSTopLocationListResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarDNSTopLocationListResponseResultMeta]
type radarDNSTopLocationListResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarDNSTopLocationListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopLocationListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarDNSTopLocationListResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                               `json:"level,required"`
	JSON        radarDNSTopLocationListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarDNSTopLocationListResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarDNSTopLocationListResponseResultMetaConfidenceInfo]
type radarDNSTopLocationListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopLocationListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopLocationListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                `json:"dataSource,required"`
	Description string                                                                `json:"description,required"`
	EndTime     time.Time                                                             `json:"endTime,required" format:"date-time"`
	EventType   string                                                                `json:"eventType,required"`
	StartTime   time.Time                                                             `json:"startTime,required" format:"date-time"`
	JSON        radarDNSTopLocationListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarDNSTopLocationListResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarDNSTopLocationListResponseResultMetaConfidenceInfoAnnotation]
type radarDNSTopLocationListResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopLocationListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopLocationListResponseResultMetaDateRange struct {
	EndTime   time.Time                                              `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      radarDNSTopLocationListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarDNSTopLocationListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarDNSTopLocationListResponseResultMetaDateRange]
type radarDNSTopLocationListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopLocationListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopLocationListResponseResultTop0 struct {
	ClientCountryAlpha2 string                                        `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                        `json:"clientCountryName,required"`
	Value               string                                        `json:"value,required"`
	JSON                radarDNSTopLocationListResponseResultTop0JSON `json:"-"`
}

// radarDNSTopLocationListResponseResultTop0JSON contains the JSON metadata for the
// struct [RadarDNSTopLocationListResponseResultTop0]
type radarDNSTopLocationListResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarDNSTopLocationListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopLocationListParams struct {
	// Array of domain names.
	Domain param.Field[[]string] `query:"domain,required"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarDNSTopLocationListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarDNSTopLocationListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarDNSTopLocationListParams]'s query parameters as
// `url.Values`.
func (r RadarDNSTopLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarDNSTopLocationListParamsDateRange string

const (
	RadarDNSTopLocationListParamsDateRange1d         RadarDNSTopLocationListParamsDateRange = "1d"
	RadarDNSTopLocationListParamsDateRange7d         RadarDNSTopLocationListParamsDateRange = "7d"
	RadarDNSTopLocationListParamsDateRange14d        RadarDNSTopLocationListParamsDateRange = "14d"
	RadarDNSTopLocationListParamsDateRange28d        RadarDNSTopLocationListParamsDateRange = "28d"
	RadarDNSTopLocationListParamsDateRange12w        RadarDNSTopLocationListParamsDateRange = "12w"
	RadarDNSTopLocationListParamsDateRange24w        RadarDNSTopLocationListParamsDateRange = "24w"
	RadarDNSTopLocationListParamsDateRange52w        RadarDNSTopLocationListParamsDateRange = "52w"
	RadarDNSTopLocationListParamsDateRange1dControl  RadarDNSTopLocationListParamsDateRange = "1dControl"
	RadarDNSTopLocationListParamsDateRange7dControl  RadarDNSTopLocationListParamsDateRange = "7dControl"
	RadarDNSTopLocationListParamsDateRange14dControl RadarDNSTopLocationListParamsDateRange = "14dControl"
	RadarDNSTopLocationListParamsDateRange28dControl RadarDNSTopLocationListParamsDateRange = "28dControl"
	RadarDNSTopLocationListParamsDateRange12wControl RadarDNSTopLocationListParamsDateRange = "12wControl"
	RadarDNSTopLocationListParamsDateRange24wControl RadarDNSTopLocationListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarDNSTopLocationListParamsFormat string

const (
	RadarDNSTopLocationListParamsFormatJson RadarDNSTopLocationListParamsFormat = "JSON"
	RadarDNSTopLocationListParamsFormatCsv  RadarDNSTopLocationListParamsFormat = "CSV"
)
