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

// RadarEmailSecurityTopAseDkimService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopAseDkimService] method instead.
type RadarEmailSecurityTopAseDkimService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopAseDkimService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopAseDkimService(opts ...option.RequestOption) (r *RadarEmailSecurityTopAseDkimService) {
	r = &RadarEmailSecurityTopAseDkimService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by email DKIM validation.
func (r *RadarEmailSecurityTopAseDkimService) Get(ctx context.Context, dkim RadarEmailSecurityTopAseDkimGetParamsDkim, query RadarEmailSecurityTopAseDkimGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopAseDkimGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/security/top/ases/dkim/%v", dkim)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityTopAseDkimGetResponse struct {
	Result  RadarEmailSecurityTopAseDkimGetResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarEmailSecurityTopAseDkimGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopAseDkimGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopAseDkimGetResponse]
type radarEmailSecurityTopAseDkimGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDkimGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDkimGetResponseResult struct {
	Meta RadarEmailSecurityTopAseDkimGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopAseDkimGetResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopAseDkimGetResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTopAseDkimGetResponseResultJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseDkimGetResponseResult]
type radarEmailSecurityTopAseDkimGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDkimGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDkimGetResponseResultMeta struct {
	DateRange      []RadarEmailSecurityTopAseDkimGetResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                          `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopAseDkimGetResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopAseDkimGetResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityTopAseDkimGetResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopAseDkimGetResponseResultMeta]
type radarEmailSecurityTopAseDkimGetResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDkimGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDkimGetResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                      `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopAseDkimGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopAseDkimGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopAseDkimGetResponseResultMetaDateRange]
type radarEmailSecurityTopAseDkimGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDkimGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDkimGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopAseDkimGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                       `json:"level"`
	JSON        radarEmailSecurityTopAseDkimGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopAseDkimGetResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopAseDkimGetResponseResultMetaConfidenceInfo]
type radarEmailSecurityTopAseDkimGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDkimGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDkimGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                        `json:"dataSource,required"`
	Description     string                                                                        `json:"description,required"`
	EventType       string                                                                        `json:"eventType,required"`
	IsInstantaneous interface{}                                                                   `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                     `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                        `json:"linkedUrl"`
	StartTime       time.Time                                                                     `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopAseDkimGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopAseDkimGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopAseDkimGetResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopAseDkimGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopAseDkimGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDkimGetResponseResultTop0 struct {
	ClientAsn    int64                                                 `json:"clientASN,required"`
	ClientAsName string                                                `json:"clientASName,required"`
	Value        string                                                `json:"value,required"`
	JSON         radarEmailSecurityTopAseDkimGetResponseResultTop0JSON `json:"-"`
}

// radarEmailSecurityTopAseDkimGetResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopAseDkimGetResponseResultTop0]
type radarEmailSecurityTopAseDkimGetResponseResultTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDkimGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDkimGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTopAseDkimGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopAseDkimGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTopAseDkimGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopAseDkimGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityTopAseDkimGetParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopAseDkimGetParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecurityTopAseDkimGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// DKIM.
type RadarEmailSecurityTopAseDkimGetParamsDkim string

const (
	RadarEmailSecurityTopAseDkimGetParamsDkimPass RadarEmailSecurityTopAseDkimGetParamsDkim = "PASS"
	RadarEmailSecurityTopAseDkimGetParamsDkimNone RadarEmailSecurityTopAseDkimGetParamsDkim = "NONE"
	RadarEmailSecurityTopAseDkimGetParamsDkimFail RadarEmailSecurityTopAseDkimGetParamsDkim = "FAIL"
)

type RadarEmailSecurityTopAseDkimGetParamsArc string

const (
	RadarEmailSecurityTopAseDkimGetParamsArcPass RadarEmailSecurityTopAseDkimGetParamsArc = "PASS"
	RadarEmailSecurityTopAseDkimGetParamsArcNone RadarEmailSecurityTopAseDkimGetParamsArc = "NONE"
	RadarEmailSecurityTopAseDkimGetParamsArcFail RadarEmailSecurityTopAseDkimGetParamsArc = "FAIL"
)

type RadarEmailSecurityTopAseDkimGetParamsDateRange string

const (
	RadarEmailSecurityTopAseDkimGetParamsDateRange1d         RadarEmailSecurityTopAseDkimGetParamsDateRange = "1d"
	RadarEmailSecurityTopAseDkimGetParamsDateRange2d         RadarEmailSecurityTopAseDkimGetParamsDateRange = "2d"
	RadarEmailSecurityTopAseDkimGetParamsDateRange7d         RadarEmailSecurityTopAseDkimGetParamsDateRange = "7d"
	RadarEmailSecurityTopAseDkimGetParamsDateRange14d        RadarEmailSecurityTopAseDkimGetParamsDateRange = "14d"
	RadarEmailSecurityTopAseDkimGetParamsDateRange28d        RadarEmailSecurityTopAseDkimGetParamsDateRange = "28d"
	RadarEmailSecurityTopAseDkimGetParamsDateRange12w        RadarEmailSecurityTopAseDkimGetParamsDateRange = "12w"
	RadarEmailSecurityTopAseDkimGetParamsDateRange24w        RadarEmailSecurityTopAseDkimGetParamsDateRange = "24w"
	RadarEmailSecurityTopAseDkimGetParamsDateRange52w        RadarEmailSecurityTopAseDkimGetParamsDateRange = "52w"
	RadarEmailSecurityTopAseDkimGetParamsDateRange1dControl  RadarEmailSecurityTopAseDkimGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopAseDkimGetParamsDateRange2dControl  RadarEmailSecurityTopAseDkimGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopAseDkimGetParamsDateRange7dControl  RadarEmailSecurityTopAseDkimGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopAseDkimGetParamsDateRange14dControl RadarEmailSecurityTopAseDkimGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopAseDkimGetParamsDateRange28dControl RadarEmailSecurityTopAseDkimGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopAseDkimGetParamsDateRange12wControl RadarEmailSecurityTopAseDkimGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopAseDkimGetParamsDateRange24wControl RadarEmailSecurityTopAseDkimGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopAseDkimGetParamsDmarc string

const (
	RadarEmailSecurityTopAseDkimGetParamsDmarcPass RadarEmailSecurityTopAseDkimGetParamsDmarc = "PASS"
	RadarEmailSecurityTopAseDkimGetParamsDmarcNone RadarEmailSecurityTopAseDkimGetParamsDmarc = "NONE"
	RadarEmailSecurityTopAseDkimGetParamsDmarcFail RadarEmailSecurityTopAseDkimGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopAseDkimGetParamsFormat string

const (
	RadarEmailSecurityTopAseDkimGetParamsFormatJson RadarEmailSecurityTopAseDkimGetParamsFormat = "JSON"
	RadarEmailSecurityTopAseDkimGetParamsFormatCsv  RadarEmailSecurityTopAseDkimGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopAseDkimGetParamsSpf string

const (
	RadarEmailSecurityTopAseDkimGetParamsSpfPass RadarEmailSecurityTopAseDkimGetParamsSpf = "PASS"
	RadarEmailSecurityTopAseDkimGetParamsSpfNone RadarEmailSecurityTopAseDkimGetParamsSpf = "NONE"
	RadarEmailSecurityTopAseDkimGetParamsSpfFail RadarEmailSecurityTopAseDkimGetParamsSpf = "FAIL"
)
