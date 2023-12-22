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

// RadarEmailTopAseMaliciousService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailTopAseMaliciousService] method instead.
type RadarEmailTopAseMaliciousService struct {
	Options []option.RequestOption
}

// NewRadarEmailTopAseMaliciousService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailTopAseMaliciousService(opts ...option.RequestOption) (r *RadarEmailTopAseMaliciousService) {
	r = &RadarEmailTopAseMaliciousService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by emails classified, of Malicious
// validations.
func (r *RadarEmailTopAseMaliciousService) Get(ctx context.Context, malicious RadarEmailTopAseMaliciousGetParamsMalicious, query RadarEmailTopAseMaliciousGetParams, opts ...option.RequestOption) (res *RadarEmailTopAseMaliciousGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/top/ases/malicious/%v", malicious)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTopAseMaliciousGetResponse struct {
	Result  RadarEmailTopAseMaliciousGetResponseResult `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarEmailTopAseMaliciousGetResponseJSON   `json:"-"`
}

// radarEmailTopAseMaliciousGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailTopAseMaliciousGetResponse]
type radarEmailTopAseMaliciousGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseMaliciousGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseMaliciousGetResponseResult struct {
	Meta RadarEmailTopAseMaliciousGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailTopAseMaliciousGetResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailTopAseMaliciousGetResponseResultJSON   `json:"-"`
}

// radarEmailTopAseMaliciousGetResponseResultJSON contains the JSON metadata for
// the struct [RadarEmailTopAseMaliciousGetResponseResult]
type radarEmailTopAseMaliciousGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseMaliciousGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseMaliciousGetResponseResultMeta struct {
	ConfidenceInfo RadarEmailTopAseMaliciousGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarEmailTopAseMaliciousGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarEmailTopAseMaliciousGetResponseResultMetaJSON           `json:"-"`
}

// radarEmailTopAseMaliciousGetResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarEmailTopAseMaliciousGetResponseResultMeta]
type radarEmailTopAseMaliciousGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailTopAseMaliciousGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseMaliciousGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailTopAseMaliciousGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                    `json:"level,required"`
	JSON        radarEmailTopAseMaliciousGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailTopAseMaliciousGetResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailTopAseMaliciousGetResponseResultMetaConfidenceInfo]
type radarEmailTopAseMaliciousGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseMaliciousGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseMaliciousGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                     `json:"dataSource,required"`
	Description string                                                                     `json:"description,required"`
	EndTime     time.Time                                                                  `json:"endTime,required" format:"date-time"`
	EventType   string                                                                     `json:"eventType,required"`
	StartTime   time.Time                                                                  `json:"startTime,required" format:"date-time"`
	JSON        radarEmailTopAseMaliciousGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailTopAseMaliciousGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailTopAseMaliciousGetResponseResultMetaConfidenceInfoAnnotation]
type radarEmailTopAseMaliciousGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseMaliciousGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseMaliciousGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                                   `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarEmailTopAseMaliciousGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailTopAseMaliciousGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailTopAseMaliciousGetResponseResultMetaDateRange]
type radarEmailTopAseMaliciousGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseMaliciousGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseMaliciousGetResponseResultTop0 struct {
	ClientASN    int64                                              `json:"clientASN,required"`
	ClientAsName string                                             `json:"clientASName,required"`
	Value        string                                             `json:"value,required"`
	JSON         radarEmailTopAseMaliciousGetResponseResultTop0JSON `json:"-"`
}

// radarEmailTopAseMaliciousGetResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarEmailTopAseMaliciousGetResponseResultTop0]
type radarEmailTopAseMaliciousGetResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailTopAseMaliciousGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseMaliciousGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailTopAseMaliciousGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTopAseMaliciousGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailTopAseMaliciousGetParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailTopAseMaliciousGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailTopAseMaliciousGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailTopAseMaliciousGetParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailTopAseMaliciousGetParams]'s query parameters as
// `url.Values`.
func (r RadarEmailTopAseMaliciousGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailTopAseMaliciousGetParamsMalicious string

const (
	RadarEmailTopAseMaliciousGetParamsMaliciousMalicious    RadarEmailTopAseMaliciousGetParamsMalicious = "MALICIOUS"
	RadarEmailTopAseMaliciousGetParamsMaliciousNotMalicious RadarEmailTopAseMaliciousGetParamsMalicious = "NOT_MALICIOUS"
)

type RadarEmailTopAseMaliciousGetParamsArc string

const (
	RadarEmailTopAseMaliciousGetParamsArcPass RadarEmailTopAseMaliciousGetParamsArc = "PASS"
	RadarEmailTopAseMaliciousGetParamsArcNone RadarEmailTopAseMaliciousGetParamsArc = "NONE"
	RadarEmailTopAseMaliciousGetParamsArcFail RadarEmailTopAseMaliciousGetParamsArc = "FAIL"
)

type RadarEmailTopAseMaliciousGetParamsDateRange string

const (
	RadarEmailTopAseMaliciousGetParamsDateRange1d         RadarEmailTopAseMaliciousGetParamsDateRange = "1d"
	RadarEmailTopAseMaliciousGetParamsDateRange7d         RadarEmailTopAseMaliciousGetParamsDateRange = "7d"
	RadarEmailTopAseMaliciousGetParamsDateRange14d        RadarEmailTopAseMaliciousGetParamsDateRange = "14d"
	RadarEmailTopAseMaliciousGetParamsDateRange28d        RadarEmailTopAseMaliciousGetParamsDateRange = "28d"
	RadarEmailTopAseMaliciousGetParamsDateRange12w        RadarEmailTopAseMaliciousGetParamsDateRange = "12w"
	RadarEmailTopAseMaliciousGetParamsDateRange24w        RadarEmailTopAseMaliciousGetParamsDateRange = "24w"
	RadarEmailTopAseMaliciousGetParamsDateRange52w        RadarEmailTopAseMaliciousGetParamsDateRange = "52w"
	RadarEmailTopAseMaliciousGetParamsDateRange1dControl  RadarEmailTopAseMaliciousGetParamsDateRange = "1dControl"
	RadarEmailTopAseMaliciousGetParamsDateRange7dControl  RadarEmailTopAseMaliciousGetParamsDateRange = "7dControl"
	RadarEmailTopAseMaliciousGetParamsDateRange14dControl RadarEmailTopAseMaliciousGetParamsDateRange = "14dControl"
	RadarEmailTopAseMaliciousGetParamsDateRange28dControl RadarEmailTopAseMaliciousGetParamsDateRange = "28dControl"
	RadarEmailTopAseMaliciousGetParamsDateRange12wControl RadarEmailTopAseMaliciousGetParamsDateRange = "12wControl"
	RadarEmailTopAseMaliciousGetParamsDateRange24wControl RadarEmailTopAseMaliciousGetParamsDateRange = "24wControl"
)

type RadarEmailTopAseMaliciousGetParamsDkim string

const (
	RadarEmailTopAseMaliciousGetParamsDkimPass RadarEmailTopAseMaliciousGetParamsDkim = "PASS"
	RadarEmailTopAseMaliciousGetParamsDkimNone RadarEmailTopAseMaliciousGetParamsDkim = "NONE"
	RadarEmailTopAseMaliciousGetParamsDkimFail RadarEmailTopAseMaliciousGetParamsDkim = "FAIL"
)

type RadarEmailTopAseMaliciousGetParamsDmarc string

const (
	RadarEmailTopAseMaliciousGetParamsDmarcPass RadarEmailTopAseMaliciousGetParamsDmarc = "PASS"
	RadarEmailTopAseMaliciousGetParamsDmarcNone RadarEmailTopAseMaliciousGetParamsDmarc = "NONE"
	RadarEmailTopAseMaliciousGetParamsDmarcFail RadarEmailTopAseMaliciousGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailTopAseMaliciousGetParamsFormat string

const (
	RadarEmailTopAseMaliciousGetParamsFormatJson RadarEmailTopAseMaliciousGetParamsFormat = "JSON"
	RadarEmailTopAseMaliciousGetParamsFormatCsv  RadarEmailTopAseMaliciousGetParamsFormat = "CSV"
)

type RadarEmailTopAseMaliciousGetParamsSpf string

const (
	RadarEmailTopAseMaliciousGetParamsSpfPass RadarEmailTopAseMaliciousGetParamsSpf = "PASS"
	RadarEmailTopAseMaliciousGetParamsSpfNone RadarEmailTopAseMaliciousGetParamsSpf = "NONE"
	RadarEmailTopAseMaliciousGetParamsSpfFail RadarEmailTopAseMaliciousGetParamsSpf = "FAIL"
)
