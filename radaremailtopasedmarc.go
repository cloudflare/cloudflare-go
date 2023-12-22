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

// RadarEmailTopAseDmarcService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailTopAseDmarcService]
// method instead.
type RadarEmailTopAseDmarcService struct {
	Options []option.RequestOption
}

// NewRadarEmailTopAseDmarcService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailTopAseDmarcService(opts ...option.RequestOption) (r *RadarEmailTopAseDmarcService) {
	r = &RadarEmailTopAseDmarcService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by emails classified, of DMARC validations.
func (r *RadarEmailTopAseDmarcService) Get(ctx context.Context, dmarc RadarEmailTopAseDmarcGetParamsDmarc, query RadarEmailTopAseDmarcGetParams, opts ...option.RequestOption) (res *RadarEmailTopAseDmarcGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/top/ases/dmarc/%v", dmarc)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTopAseDmarcGetResponse struct {
	Result  RadarEmailTopAseDmarcGetResponseResult `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarEmailTopAseDmarcGetResponseJSON   `json:"-"`
}

// radarEmailTopAseDmarcGetResponseJSON contains the JSON metadata for the struct
// [RadarEmailTopAseDmarcGetResponse]
type radarEmailTopAseDmarcGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseDmarcGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseDmarcGetResponseResult struct {
	Meta RadarEmailTopAseDmarcGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailTopAseDmarcGetResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailTopAseDmarcGetResponseResultJSON   `json:"-"`
}

// radarEmailTopAseDmarcGetResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailTopAseDmarcGetResponseResult]
type radarEmailTopAseDmarcGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseDmarcGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseDmarcGetResponseResultMeta struct {
	ConfidenceInfo RadarEmailTopAseDmarcGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarEmailTopAseDmarcGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarEmailTopAseDmarcGetResponseResultMetaJSON           `json:"-"`
}

// radarEmailTopAseDmarcGetResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarEmailTopAseDmarcGetResponseResultMeta]
type radarEmailTopAseDmarcGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailTopAseDmarcGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseDmarcGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailTopAseDmarcGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                `json:"level,required"`
	JSON        radarEmailTopAseDmarcGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailTopAseDmarcGetResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailTopAseDmarcGetResponseResultMetaConfidenceInfo]
type radarEmailTopAseDmarcGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseDmarcGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseDmarcGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                 `json:"dataSource,required"`
	Description string                                                                 `json:"description,required"`
	EndTime     time.Time                                                              `json:"endTime,required" format:"date-time"`
	EventType   string                                                                 `json:"eventType,required"`
	StartTime   time.Time                                                              `json:"startTime,required" format:"date-time"`
	JSON        radarEmailTopAseDmarcGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailTopAseDmarcGetResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailTopAseDmarcGetResponseResultMetaConfidenceInfoAnnotation]
type radarEmailTopAseDmarcGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseDmarcGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseDmarcGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                               `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarEmailTopAseDmarcGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailTopAseDmarcGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailTopAseDmarcGetResponseResultMetaDateRange]
type radarEmailTopAseDmarcGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseDmarcGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseDmarcGetResponseResultTop0 struct {
	ClientASN    int64                                          `json:"clientASN,required"`
	ClientAsName string                                         `json:"clientASName,required"`
	Value        string                                         `json:"value,required"`
	JSON         radarEmailTopAseDmarcGetResponseResultTop0JSON `json:"-"`
}

// radarEmailTopAseDmarcGetResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarEmailTopAseDmarcGetResponseResultTop0]
type radarEmailTopAseDmarcGetResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailTopAseDmarcGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseDmarcGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailTopAseDmarcGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTopAseDmarcGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailTopAseDmarcGetParamsDkim] `query:"dkim"`
	// Format results are returned in.
	Format param.Field[RadarEmailTopAseDmarcGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailTopAseDmarcGetParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailTopAseDmarcGetParams]'s query parameters as
// `url.Values`.
func (r RadarEmailTopAseDmarcGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailTopAseDmarcGetParamsDmarc string

const (
	RadarEmailTopAseDmarcGetParamsDmarcPass RadarEmailTopAseDmarcGetParamsDmarc = "PASS"
	RadarEmailTopAseDmarcGetParamsDmarcNone RadarEmailTopAseDmarcGetParamsDmarc = "NONE"
	RadarEmailTopAseDmarcGetParamsDmarcFail RadarEmailTopAseDmarcGetParamsDmarc = "FAIL"
)

type RadarEmailTopAseDmarcGetParamsArc string

const (
	RadarEmailTopAseDmarcGetParamsArcPass RadarEmailTopAseDmarcGetParamsArc = "PASS"
	RadarEmailTopAseDmarcGetParamsArcNone RadarEmailTopAseDmarcGetParamsArc = "NONE"
	RadarEmailTopAseDmarcGetParamsArcFail RadarEmailTopAseDmarcGetParamsArc = "FAIL"
)

type RadarEmailTopAseDmarcGetParamsDateRange string

const (
	RadarEmailTopAseDmarcGetParamsDateRange1d         RadarEmailTopAseDmarcGetParamsDateRange = "1d"
	RadarEmailTopAseDmarcGetParamsDateRange7d         RadarEmailTopAseDmarcGetParamsDateRange = "7d"
	RadarEmailTopAseDmarcGetParamsDateRange14d        RadarEmailTopAseDmarcGetParamsDateRange = "14d"
	RadarEmailTopAseDmarcGetParamsDateRange28d        RadarEmailTopAseDmarcGetParamsDateRange = "28d"
	RadarEmailTopAseDmarcGetParamsDateRange12w        RadarEmailTopAseDmarcGetParamsDateRange = "12w"
	RadarEmailTopAseDmarcGetParamsDateRange24w        RadarEmailTopAseDmarcGetParamsDateRange = "24w"
	RadarEmailTopAseDmarcGetParamsDateRange52w        RadarEmailTopAseDmarcGetParamsDateRange = "52w"
	RadarEmailTopAseDmarcGetParamsDateRange1dControl  RadarEmailTopAseDmarcGetParamsDateRange = "1dControl"
	RadarEmailTopAseDmarcGetParamsDateRange7dControl  RadarEmailTopAseDmarcGetParamsDateRange = "7dControl"
	RadarEmailTopAseDmarcGetParamsDateRange14dControl RadarEmailTopAseDmarcGetParamsDateRange = "14dControl"
	RadarEmailTopAseDmarcGetParamsDateRange28dControl RadarEmailTopAseDmarcGetParamsDateRange = "28dControl"
	RadarEmailTopAseDmarcGetParamsDateRange12wControl RadarEmailTopAseDmarcGetParamsDateRange = "12wControl"
	RadarEmailTopAseDmarcGetParamsDateRange24wControl RadarEmailTopAseDmarcGetParamsDateRange = "24wControl"
)

type RadarEmailTopAseDmarcGetParamsDkim string

const (
	RadarEmailTopAseDmarcGetParamsDkimPass RadarEmailTopAseDmarcGetParamsDkim = "PASS"
	RadarEmailTopAseDmarcGetParamsDkimNone RadarEmailTopAseDmarcGetParamsDkim = "NONE"
	RadarEmailTopAseDmarcGetParamsDkimFail RadarEmailTopAseDmarcGetParamsDkim = "FAIL"
)

// Format results are returned in.
type RadarEmailTopAseDmarcGetParamsFormat string

const (
	RadarEmailTopAseDmarcGetParamsFormatJson RadarEmailTopAseDmarcGetParamsFormat = "JSON"
	RadarEmailTopAseDmarcGetParamsFormatCsv  RadarEmailTopAseDmarcGetParamsFormat = "CSV"
)

type RadarEmailTopAseDmarcGetParamsSpf string

const (
	RadarEmailTopAseDmarcGetParamsSpfPass RadarEmailTopAseDmarcGetParamsSpf = "PASS"
	RadarEmailTopAseDmarcGetParamsSpfNone RadarEmailTopAseDmarcGetParamsSpf = "NONE"
	RadarEmailTopAseDmarcGetParamsSpfFail RadarEmailTopAseDmarcGetParamsSpf = "FAIL"
)
