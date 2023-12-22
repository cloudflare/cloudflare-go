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

// RadarDNSTopAseService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarDNSTopAseService] method
// instead.
type RadarDNSTopAseService struct {
	Options []option.RequestOption
}

// NewRadarDNSTopAseService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarDNSTopAseService(opts ...option.RequestOption) (r *RadarDNSTopAseService) {
	r = &RadarDNSTopAseService{}
	r.Options = opts
	return
}

// Get top autonomous systems by DNS queries made to Cloudflare's public DNS
// resolver.
func (r *RadarDNSTopAseService) List(ctx context.Context, query RadarDNSTopAseListParams, opts ...option.RequestOption) (res *RadarDNSTopAseListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarDNSTopAseListResponse struct {
	Result  RadarDNSTopAseListResponseResult `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    radarDNSTopAseListResponseJSON   `json:"-"`
}

// radarDNSTopAseListResponseJSON contains the JSON metadata for the struct
// [RadarDNSTopAseListResponse]
type radarDNSTopAseListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopAseListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopAseListResponseResult struct {
	Meta RadarDNSTopAseListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarDNSTopAseListResponseResultTop0 `json:"top_0,required"`
	JSON radarDNSTopAseListResponseResultJSON   `json:"-"`
}

// radarDNSTopAseListResponseResultJSON contains the JSON metadata for the struct
// [RadarDNSTopAseListResponseResult]
type radarDNSTopAseListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopAseListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopAseListResponseResultMeta struct {
	ConfidenceInfo RadarDNSTopAseListResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarDNSTopAseListResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarDNSTopAseListResponseResultMetaJSON           `json:"-"`
}

// radarDNSTopAseListResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarDNSTopAseListResponseResultMeta]
type radarDNSTopAseListResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarDNSTopAseListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopAseListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarDNSTopAseListResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                          `json:"level,required"`
	JSON        radarDNSTopAseListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarDNSTopAseListResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarDNSTopAseListResponseResultMetaConfidenceInfo]
type radarDNSTopAseListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopAseListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopAseListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                           `json:"dataSource,required"`
	Description string                                                           `json:"description,required"`
	EndTime     time.Time                                                        `json:"endTime,required" format:"date-time"`
	EventType   string                                                           `json:"eventType,required"`
	StartTime   time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON        radarDNSTopAseListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarDNSTopAseListResponseResultMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarDNSTopAseListResponseResultMetaConfidenceInfoAnnotation]
type radarDNSTopAseListResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopAseListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopAseListResponseResultMetaDateRange struct {
	EndTime   time.Time                                         `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      radarDNSTopAseListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarDNSTopAseListResponseResultMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarDNSTopAseListResponseResultMetaDateRange]
type radarDNSTopAseListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopAseListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopAseListResponseResultTop0 struct {
	ClientASN    int64                                    `json:"clientASN,required"`
	ClientAsName string                                   `json:"clientASName,required"`
	Value        string                                   `json:"value,required"`
	JSON         radarDNSTopAseListResponseResultTop0JSON `json:"-"`
}

// radarDNSTopAseListResponseResultTop0JSON contains the JSON metadata for the
// struct [RadarDNSTopAseListResponseResultTop0]
type radarDNSTopAseListResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarDNSTopAseListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopAseListParams struct {
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
	DateRange param.Field[[]RadarDNSTopAseListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarDNSTopAseListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarDNSTopAseListParams]'s query parameters as
// `url.Values`.
func (r RadarDNSTopAseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarDNSTopAseListParamsDateRange string

const (
	RadarDNSTopAseListParamsDateRange1d         RadarDNSTopAseListParamsDateRange = "1d"
	RadarDNSTopAseListParamsDateRange7d         RadarDNSTopAseListParamsDateRange = "7d"
	RadarDNSTopAseListParamsDateRange14d        RadarDNSTopAseListParamsDateRange = "14d"
	RadarDNSTopAseListParamsDateRange28d        RadarDNSTopAseListParamsDateRange = "28d"
	RadarDNSTopAseListParamsDateRange12w        RadarDNSTopAseListParamsDateRange = "12w"
	RadarDNSTopAseListParamsDateRange24w        RadarDNSTopAseListParamsDateRange = "24w"
	RadarDNSTopAseListParamsDateRange52w        RadarDNSTopAseListParamsDateRange = "52w"
	RadarDNSTopAseListParamsDateRange1dControl  RadarDNSTopAseListParamsDateRange = "1dControl"
	RadarDNSTopAseListParamsDateRange7dControl  RadarDNSTopAseListParamsDateRange = "7dControl"
	RadarDNSTopAseListParamsDateRange14dControl RadarDNSTopAseListParamsDateRange = "14dControl"
	RadarDNSTopAseListParamsDateRange28dControl RadarDNSTopAseListParamsDateRange = "28dControl"
	RadarDNSTopAseListParamsDateRange12wControl RadarDNSTopAseListParamsDateRange = "12wControl"
	RadarDNSTopAseListParamsDateRange24wControl RadarDNSTopAseListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarDNSTopAseListParamsFormat string

const (
	RadarDNSTopAseListParamsFormatJson RadarDNSTopAseListParamsFormat = "JSON"
	RadarDNSTopAseListParamsFormatCsv  RadarDNSTopAseListParamsFormat = "CSV"
)
