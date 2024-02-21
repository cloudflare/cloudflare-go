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

// RadarEmailSecurityTopAseDKIMService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopAseDKIMService] method instead.
type RadarEmailSecurityTopAseDKIMService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopAseDKIMService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopAseDKIMService(opts ...option.RequestOption) (r *RadarEmailSecurityTopAseDKIMService) {
	r = &RadarEmailSecurityTopAseDKIMService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by email DKIM validation.
func (r *RadarEmailSecurityTopAseDKIMService) Get(ctx context.Context, dkim RadarEmailSecurityTopAseDKIMGetParamsDKIM, query RadarEmailSecurityTopAseDKIMGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopAseDKIMGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopAseDKIMGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/ases/dkim/%v", dkim)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopAseDKIMGetResponse struct {
	Meta RadarEmailSecurityTopAseDKIMGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopAseDKIMGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopAseDKIMGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopAseDKIMGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopAseDKIMGetResponse]
type radarEmailSecurityTopAseDKIMGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDKIMGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDKIMGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopAseDKIMGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                    `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopAseDKIMGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopAseDKIMGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopAseDKIMGetResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseDKIMGetResponseMeta]
type radarEmailSecurityTopAseDKIMGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDKIMGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDKIMGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopAseDKIMGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopAseDKIMGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopAseDKIMGetResponseMetaDateRange]
type radarEmailSecurityTopAseDKIMGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDKIMGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDKIMGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopAseDKIMGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                 `json:"level"`
	JSON        radarEmailSecurityTopAseDKIMGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopAseDKIMGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopAseDKIMGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopAseDKIMGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDKIMGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDKIMGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                  `json:"dataSource,required"`
	Description     string                                                                  `json:"description,required"`
	EventType       string                                                                  `json:"eventType,required"`
	IsInstantaneous interface{}                                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                                               `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                  `json:"linkedUrl"`
	StartTime       time.Time                                                               `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopAseDKIMGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopAseDKIMGetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTopAseDKIMGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopAseDKIMGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopAseDKIMGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDKIMGetResponseTop0 struct {
	ClientAsn    int64                                           `json:"clientASN,required"`
	ClientAsName string                                          `json:"clientASName,required"`
	Value        string                                          `json:"value,required"`
	JSON         radarEmailSecurityTopAseDKIMGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopAseDKIMGetResponseTop0JSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseDKIMGetResponseTop0]
type radarEmailSecurityTopAseDKIMGetResponseTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDKIMGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDKIMGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTopAseDKIMGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopAseDKIMGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTopAseDKIMGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopAseDKIMGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopAseDKIMGetParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopAseDKIMGetParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecurityTopAseDKIMGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// DKIM.
type RadarEmailSecurityTopAseDKIMGetParamsDKIM string

const (
	RadarEmailSecurityTopAseDKIMGetParamsDKIMPass RadarEmailSecurityTopAseDKIMGetParamsDKIM = "PASS"
	RadarEmailSecurityTopAseDKIMGetParamsDKIMNone RadarEmailSecurityTopAseDKIMGetParamsDKIM = "NONE"
	RadarEmailSecurityTopAseDKIMGetParamsDKIMFail RadarEmailSecurityTopAseDKIMGetParamsDKIM = "FAIL"
)

type RadarEmailSecurityTopAseDKIMGetParamsArc string

const (
	RadarEmailSecurityTopAseDKIMGetParamsArcPass RadarEmailSecurityTopAseDKIMGetParamsArc = "PASS"
	RadarEmailSecurityTopAseDKIMGetParamsArcNone RadarEmailSecurityTopAseDKIMGetParamsArc = "NONE"
	RadarEmailSecurityTopAseDKIMGetParamsArcFail RadarEmailSecurityTopAseDKIMGetParamsArc = "FAIL"
)

type RadarEmailSecurityTopAseDKIMGetParamsDateRange string

const (
	RadarEmailSecurityTopAseDKIMGetParamsDateRange1d         RadarEmailSecurityTopAseDKIMGetParamsDateRange = "1d"
	RadarEmailSecurityTopAseDKIMGetParamsDateRange2d         RadarEmailSecurityTopAseDKIMGetParamsDateRange = "2d"
	RadarEmailSecurityTopAseDKIMGetParamsDateRange7d         RadarEmailSecurityTopAseDKIMGetParamsDateRange = "7d"
	RadarEmailSecurityTopAseDKIMGetParamsDateRange14d        RadarEmailSecurityTopAseDKIMGetParamsDateRange = "14d"
	RadarEmailSecurityTopAseDKIMGetParamsDateRange28d        RadarEmailSecurityTopAseDKIMGetParamsDateRange = "28d"
	RadarEmailSecurityTopAseDKIMGetParamsDateRange12w        RadarEmailSecurityTopAseDKIMGetParamsDateRange = "12w"
	RadarEmailSecurityTopAseDKIMGetParamsDateRange24w        RadarEmailSecurityTopAseDKIMGetParamsDateRange = "24w"
	RadarEmailSecurityTopAseDKIMGetParamsDateRange52w        RadarEmailSecurityTopAseDKIMGetParamsDateRange = "52w"
	RadarEmailSecurityTopAseDKIMGetParamsDateRange1dControl  RadarEmailSecurityTopAseDKIMGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopAseDKIMGetParamsDateRange2dControl  RadarEmailSecurityTopAseDKIMGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopAseDKIMGetParamsDateRange7dControl  RadarEmailSecurityTopAseDKIMGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopAseDKIMGetParamsDateRange14dControl RadarEmailSecurityTopAseDKIMGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopAseDKIMGetParamsDateRange28dControl RadarEmailSecurityTopAseDKIMGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopAseDKIMGetParamsDateRange12wControl RadarEmailSecurityTopAseDKIMGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopAseDKIMGetParamsDateRange24wControl RadarEmailSecurityTopAseDKIMGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopAseDKIMGetParamsDmarc string

const (
	RadarEmailSecurityTopAseDKIMGetParamsDmarcPass RadarEmailSecurityTopAseDKIMGetParamsDmarc = "PASS"
	RadarEmailSecurityTopAseDKIMGetParamsDmarcNone RadarEmailSecurityTopAseDKIMGetParamsDmarc = "NONE"
	RadarEmailSecurityTopAseDKIMGetParamsDmarcFail RadarEmailSecurityTopAseDKIMGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopAseDKIMGetParamsFormat string

const (
	RadarEmailSecurityTopAseDKIMGetParamsFormatJson RadarEmailSecurityTopAseDKIMGetParamsFormat = "JSON"
	RadarEmailSecurityTopAseDKIMGetParamsFormatCsv  RadarEmailSecurityTopAseDKIMGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopAseDKIMGetParamsSPF string

const (
	RadarEmailSecurityTopAseDKIMGetParamsSPFPass RadarEmailSecurityTopAseDKIMGetParamsSPF = "PASS"
	RadarEmailSecurityTopAseDKIMGetParamsSPFNone RadarEmailSecurityTopAseDKIMGetParamsSPF = "NONE"
	RadarEmailSecurityTopAseDKIMGetParamsSPFFail RadarEmailSecurityTopAseDKIMGetParamsSPF = "FAIL"
)

type RadarEmailSecurityTopAseDKIMGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopAseDKIMGetResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarEmailSecurityTopAseDKIMGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopAseDKIMGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopAseDKIMGetResponseEnvelope]
type radarEmailSecurityTopAseDKIMGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDKIMGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
