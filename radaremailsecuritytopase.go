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

// RadarEmailSecurityTopAseService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopAseService] method instead.
type RadarEmailSecurityTopAseService struct {
	Options []option.RequestOption
	Arc     *RadarEmailSecurityTopAseArcService
	Dkim    *RadarEmailSecurityTopAseDkimService
	Dmarc   *RadarEmailSecurityTopAseDmarcService
}

// NewRadarEmailSecurityTopAseService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopAseService(opts ...option.RequestOption) (r *RadarEmailSecurityTopAseService) {
	r = &RadarEmailSecurityTopAseService{}
	r.Options = opts
	r.Arc = NewRadarEmailSecurityTopAseArcService(opts...)
	r.Dkim = NewRadarEmailSecurityTopAseDkimService(opts...)
	r.Dmarc = NewRadarEmailSecurityTopAseDmarcService(opts...)
	return
}

// Get the top autonomous systems (AS) by email messages. Values are a percentage
// out of the total emails.
func (r *RadarEmailSecurityTopAseService) List(ctx context.Context, query RadarEmailSecurityTopAseListParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopAseListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopAseListResponseEnvelope
	path := "radar/email/security/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopAseListResponse struct {
	Meta RadarEmailSecurityTopAseListResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopAseListResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopAseListResponseJSON   `json:"-"`
}

// radarEmailSecurityTopAseListResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopAseListResponse]
type radarEmailSecurityTopAseListResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseListResponseMeta struct {
	DateRange      []RadarEmailSecurityTopAseListResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                 `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopAseListResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopAseListResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopAseListResponseMetaJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopAseListResponseMeta]
type radarEmailSecurityTopAseListResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopAseListResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopAseListResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopAseListResponseMetaDateRange]
type radarEmailSecurityTopAseListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseListResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopAseListResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        radarEmailSecurityTopAseListResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopAseListResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopAseListResponseMetaConfidenceInfo]
type radarEmailSecurityTopAseListResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseListResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseListResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopAseListResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopAseListResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTopAseListResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopAseListResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopAseListResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseListResponseTop0 struct {
	ClientAsn    int64                                        `json:"clientASN,required"`
	ClientAsName string                                       `json:"clientASName,required"`
	Value        string                                       `json:"value,required"`
	JSON         radarEmailSecurityTopAseListResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopAseListResponseTop0JSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopAseListResponseTop0]
type radarEmailSecurityTopAseListResponseTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseListResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTopAseListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopAseListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityTopAseListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTopAseListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopAseListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityTopAseListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopAseListParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecurityTopAseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecurityTopAseListParamsArc string

const (
	RadarEmailSecurityTopAseListParamsArcPass RadarEmailSecurityTopAseListParamsArc = "PASS"
	RadarEmailSecurityTopAseListParamsArcNone RadarEmailSecurityTopAseListParamsArc = "NONE"
	RadarEmailSecurityTopAseListParamsArcFail RadarEmailSecurityTopAseListParamsArc = "FAIL"
)

type RadarEmailSecurityTopAseListParamsDateRange string

const (
	RadarEmailSecurityTopAseListParamsDateRange1d         RadarEmailSecurityTopAseListParamsDateRange = "1d"
	RadarEmailSecurityTopAseListParamsDateRange2d         RadarEmailSecurityTopAseListParamsDateRange = "2d"
	RadarEmailSecurityTopAseListParamsDateRange7d         RadarEmailSecurityTopAseListParamsDateRange = "7d"
	RadarEmailSecurityTopAseListParamsDateRange14d        RadarEmailSecurityTopAseListParamsDateRange = "14d"
	RadarEmailSecurityTopAseListParamsDateRange28d        RadarEmailSecurityTopAseListParamsDateRange = "28d"
	RadarEmailSecurityTopAseListParamsDateRange12w        RadarEmailSecurityTopAseListParamsDateRange = "12w"
	RadarEmailSecurityTopAseListParamsDateRange24w        RadarEmailSecurityTopAseListParamsDateRange = "24w"
	RadarEmailSecurityTopAseListParamsDateRange52w        RadarEmailSecurityTopAseListParamsDateRange = "52w"
	RadarEmailSecurityTopAseListParamsDateRange1dControl  RadarEmailSecurityTopAseListParamsDateRange = "1dControl"
	RadarEmailSecurityTopAseListParamsDateRange2dControl  RadarEmailSecurityTopAseListParamsDateRange = "2dControl"
	RadarEmailSecurityTopAseListParamsDateRange7dControl  RadarEmailSecurityTopAseListParamsDateRange = "7dControl"
	RadarEmailSecurityTopAseListParamsDateRange14dControl RadarEmailSecurityTopAseListParamsDateRange = "14dControl"
	RadarEmailSecurityTopAseListParamsDateRange28dControl RadarEmailSecurityTopAseListParamsDateRange = "28dControl"
	RadarEmailSecurityTopAseListParamsDateRange12wControl RadarEmailSecurityTopAseListParamsDateRange = "12wControl"
	RadarEmailSecurityTopAseListParamsDateRange24wControl RadarEmailSecurityTopAseListParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopAseListParamsDkim string

const (
	RadarEmailSecurityTopAseListParamsDkimPass RadarEmailSecurityTopAseListParamsDkim = "PASS"
	RadarEmailSecurityTopAseListParamsDkimNone RadarEmailSecurityTopAseListParamsDkim = "NONE"
	RadarEmailSecurityTopAseListParamsDkimFail RadarEmailSecurityTopAseListParamsDkim = "FAIL"
)

type RadarEmailSecurityTopAseListParamsDmarc string

const (
	RadarEmailSecurityTopAseListParamsDmarcPass RadarEmailSecurityTopAseListParamsDmarc = "PASS"
	RadarEmailSecurityTopAseListParamsDmarcNone RadarEmailSecurityTopAseListParamsDmarc = "NONE"
	RadarEmailSecurityTopAseListParamsDmarcFail RadarEmailSecurityTopAseListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopAseListParamsFormat string

const (
	RadarEmailSecurityTopAseListParamsFormatJson RadarEmailSecurityTopAseListParamsFormat = "JSON"
	RadarEmailSecurityTopAseListParamsFormatCsv  RadarEmailSecurityTopAseListParamsFormat = "CSV"
)

type RadarEmailSecurityTopAseListParamsSpf string

const (
	RadarEmailSecurityTopAseListParamsSpfPass RadarEmailSecurityTopAseListParamsSpf = "PASS"
	RadarEmailSecurityTopAseListParamsSpfNone RadarEmailSecurityTopAseListParamsSpf = "NONE"
	RadarEmailSecurityTopAseListParamsSpfFail RadarEmailSecurityTopAseListParamsSpf = "FAIL"
)

type RadarEmailSecurityTopAseListResponseEnvelope struct {
	Result  RadarEmailSecurityTopAseListResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarEmailSecurityTopAseListResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopAseListResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseListResponseEnvelope]
type radarEmailSecurityTopAseListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
