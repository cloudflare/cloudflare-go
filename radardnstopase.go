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
	var env RadarDNSTopAseListResponseEnvelope
	path := "radar/dns/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarDNSTopAseListResponse struct {
	Meta RadarDNSTopAseListResponseMeta   `json:"meta,required"`
	Top0 []RadarDNSTopAseListResponseTop0 `json:"top_0,required"`
	JSON radarDNSTopAseListResponseJSON   `json:"-"`
}

// radarDNSTopAseListResponseJSON contains the JSON metadata for the struct
// [RadarDNSTopAseListResponse]
type radarDNSTopAseListResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopAseListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopAseListResponseMeta struct {
	DateRange      []RadarDNSTopAseListResponseMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarDNSTopAseListResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarDNSTopAseListResponseMetaJSON           `json:"-"`
}

// radarDNSTopAseListResponseMetaJSON contains the JSON metadata for the struct
// [RadarDNSTopAseListResponseMeta]
type radarDNSTopAseListResponseMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarDNSTopAseListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopAseListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                   `json:"startTime,required" format:"date-time"`
	JSON      radarDNSTopAseListResponseMetaDateRangeJSON `json:"-"`
}

// radarDNSTopAseListResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarDNSTopAseListResponseMetaDateRange]
type radarDNSTopAseListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopAseListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopAseListResponseMetaConfidenceInfo struct {
	Annotations []RadarDNSTopAseListResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                    `json:"level"`
	JSON        radarDNSTopAseListResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarDNSTopAseListResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [RadarDNSTopAseListResponseMetaConfidenceInfo]
type radarDNSTopAseListResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopAseListResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopAseListResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                     `json:"dataSource,required"`
	Description     string                                                     `json:"description,required"`
	EventType       string                                                     `json:"eventType,required"`
	IsInstantaneous interface{}                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                  `json:"startTime" format:"date-time"`
	JSON            radarDNSTopAseListResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarDNSTopAseListResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [RadarDNSTopAseListResponseMetaConfidenceInfoAnnotation]
type radarDNSTopAseListResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarDNSTopAseListResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopAseListResponseTop0 struct {
	ClientAsn    int64                              `json:"clientASN,required"`
	ClientAsName string                             `json:"clientASName,required"`
	Value        string                             `json:"value,required"`
	JSON         radarDNSTopAseListResponseTop0JSON `json:"-"`
}

// radarDNSTopAseListResponseTop0JSON contains the JSON metadata for the struct
// [RadarDNSTopAseListResponseTop0]
type radarDNSTopAseListResponseTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarDNSTopAseListResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarDNSTopAseListParams struct {
	// Array of domain names.
	Domain param.Field[[]string] `query:"domain,required"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
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
	RadarDNSTopAseListParamsDateRange2d         RadarDNSTopAseListParamsDateRange = "2d"
	RadarDNSTopAseListParamsDateRange7d         RadarDNSTopAseListParamsDateRange = "7d"
	RadarDNSTopAseListParamsDateRange14d        RadarDNSTopAseListParamsDateRange = "14d"
	RadarDNSTopAseListParamsDateRange28d        RadarDNSTopAseListParamsDateRange = "28d"
	RadarDNSTopAseListParamsDateRange12w        RadarDNSTopAseListParamsDateRange = "12w"
	RadarDNSTopAseListParamsDateRange24w        RadarDNSTopAseListParamsDateRange = "24w"
	RadarDNSTopAseListParamsDateRange52w        RadarDNSTopAseListParamsDateRange = "52w"
	RadarDNSTopAseListParamsDateRange1dControl  RadarDNSTopAseListParamsDateRange = "1dControl"
	RadarDNSTopAseListParamsDateRange2dControl  RadarDNSTopAseListParamsDateRange = "2dControl"
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

type RadarDNSTopAseListResponseEnvelope struct {
	Result  RadarDNSTopAseListResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarDNSTopAseListResponseEnvelopeJSON `json:"-"`
}

// radarDNSTopAseListResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarDNSTopAseListResponseEnvelope]
type radarDNSTopAseListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopAseListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
