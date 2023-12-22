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

// RadarEmailTopAseSpamService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailTopAseSpamService]
// method instead.
type RadarEmailTopAseSpamService struct {
	Options []option.RequestOption
}

// NewRadarEmailTopAseSpamService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailTopAseSpamService(opts ...option.RequestOption) (r *RadarEmailTopAseSpamService) {
	r = &RadarEmailTopAseSpamService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by emails classified, of Spam validations.
func (r *RadarEmailTopAseSpamService) Get(ctx context.Context, spam RadarEmailTopAseSpamGetParamsSpam, query RadarEmailTopAseSpamGetParams, opts ...option.RequestOption) (res *RadarEmailTopAseSpamGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/top/ases/spam/%v", spam)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTopAseSpamGetResponse struct {
	Result  RadarEmailTopAseSpamGetResponseResult `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarEmailTopAseSpamGetResponseJSON   `json:"-"`
}

// radarEmailTopAseSpamGetResponseJSON contains the JSON metadata for the struct
// [RadarEmailTopAseSpamGetResponse]
type radarEmailTopAseSpamGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseSpamGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseSpamGetResponseResult struct {
	Meta RadarEmailTopAseSpamGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailTopAseSpamGetResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailTopAseSpamGetResponseResultJSON   `json:"-"`
}

// radarEmailTopAseSpamGetResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailTopAseSpamGetResponseResult]
type radarEmailTopAseSpamGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseSpamGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseSpamGetResponseResultMeta struct {
	ConfidenceInfo RadarEmailTopAseSpamGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarEmailTopAseSpamGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarEmailTopAseSpamGetResponseResultMetaJSON           `json:"-"`
}

// radarEmailTopAseSpamGetResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarEmailTopAseSpamGetResponseResultMeta]
type radarEmailTopAseSpamGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailTopAseSpamGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseSpamGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailTopAseSpamGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                               `json:"level,required"`
	JSON        radarEmailTopAseSpamGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailTopAseSpamGetResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailTopAseSpamGetResponseResultMetaConfidenceInfo]
type radarEmailTopAseSpamGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseSpamGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseSpamGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                `json:"dataSource,required"`
	Description string                                                                `json:"description,required"`
	EndTime     time.Time                                                             `json:"endTime,required" format:"date-time"`
	EventType   string                                                                `json:"eventType,required"`
	StartTime   time.Time                                                             `json:"startTime,required" format:"date-time"`
	JSON        radarEmailTopAseSpamGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailTopAseSpamGetResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailTopAseSpamGetResponseResultMetaConfidenceInfoAnnotation]
type radarEmailTopAseSpamGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseSpamGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseSpamGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                              `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      radarEmailTopAseSpamGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailTopAseSpamGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailTopAseSpamGetResponseResultMetaDateRange]
type radarEmailTopAseSpamGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseSpamGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseSpamGetResponseResultTop0 struct {
	ClientASN    int64                                         `json:"clientASN,required"`
	ClientAsName string                                        `json:"clientASName,required"`
	Value        string                                        `json:"value,required"`
	JSON         radarEmailTopAseSpamGetResponseResultTop0JSON `json:"-"`
}

// radarEmailTopAseSpamGetResponseResultTop0JSON contains the JSON metadata for the
// struct [RadarEmailTopAseSpamGetResponseResultTop0]
type radarEmailTopAseSpamGetResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailTopAseSpamGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseSpamGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailTopAseSpamGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTopAseSpamGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailTopAseSpamGetParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailTopAseSpamGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailTopAseSpamGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailTopAseSpamGetParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailTopAseSpamGetParams]'s query parameters as
// `url.Values`.
func (r RadarEmailTopAseSpamGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailTopAseSpamGetParamsSpam string

const (
	RadarEmailTopAseSpamGetParamsSpamSpam    RadarEmailTopAseSpamGetParamsSpam = "SPAM"
	RadarEmailTopAseSpamGetParamsSpamNotSpam RadarEmailTopAseSpamGetParamsSpam = "NOT_SPAM"
)

type RadarEmailTopAseSpamGetParamsArc string

const (
	RadarEmailTopAseSpamGetParamsArcPass RadarEmailTopAseSpamGetParamsArc = "PASS"
	RadarEmailTopAseSpamGetParamsArcNone RadarEmailTopAseSpamGetParamsArc = "NONE"
	RadarEmailTopAseSpamGetParamsArcFail RadarEmailTopAseSpamGetParamsArc = "FAIL"
)

type RadarEmailTopAseSpamGetParamsDateRange string

const (
	RadarEmailTopAseSpamGetParamsDateRange1d         RadarEmailTopAseSpamGetParamsDateRange = "1d"
	RadarEmailTopAseSpamGetParamsDateRange7d         RadarEmailTopAseSpamGetParamsDateRange = "7d"
	RadarEmailTopAseSpamGetParamsDateRange14d        RadarEmailTopAseSpamGetParamsDateRange = "14d"
	RadarEmailTopAseSpamGetParamsDateRange28d        RadarEmailTopAseSpamGetParamsDateRange = "28d"
	RadarEmailTopAseSpamGetParamsDateRange12w        RadarEmailTopAseSpamGetParamsDateRange = "12w"
	RadarEmailTopAseSpamGetParamsDateRange24w        RadarEmailTopAseSpamGetParamsDateRange = "24w"
	RadarEmailTopAseSpamGetParamsDateRange52w        RadarEmailTopAseSpamGetParamsDateRange = "52w"
	RadarEmailTopAseSpamGetParamsDateRange1dControl  RadarEmailTopAseSpamGetParamsDateRange = "1dControl"
	RadarEmailTopAseSpamGetParamsDateRange7dControl  RadarEmailTopAseSpamGetParamsDateRange = "7dControl"
	RadarEmailTopAseSpamGetParamsDateRange14dControl RadarEmailTopAseSpamGetParamsDateRange = "14dControl"
	RadarEmailTopAseSpamGetParamsDateRange28dControl RadarEmailTopAseSpamGetParamsDateRange = "28dControl"
	RadarEmailTopAseSpamGetParamsDateRange12wControl RadarEmailTopAseSpamGetParamsDateRange = "12wControl"
	RadarEmailTopAseSpamGetParamsDateRange24wControl RadarEmailTopAseSpamGetParamsDateRange = "24wControl"
)

type RadarEmailTopAseSpamGetParamsDkim string

const (
	RadarEmailTopAseSpamGetParamsDkimPass RadarEmailTopAseSpamGetParamsDkim = "PASS"
	RadarEmailTopAseSpamGetParamsDkimNone RadarEmailTopAseSpamGetParamsDkim = "NONE"
	RadarEmailTopAseSpamGetParamsDkimFail RadarEmailTopAseSpamGetParamsDkim = "FAIL"
)

type RadarEmailTopAseSpamGetParamsDmarc string

const (
	RadarEmailTopAseSpamGetParamsDmarcPass RadarEmailTopAseSpamGetParamsDmarc = "PASS"
	RadarEmailTopAseSpamGetParamsDmarcNone RadarEmailTopAseSpamGetParamsDmarc = "NONE"
	RadarEmailTopAseSpamGetParamsDmarcFail RadarEmailTopAseSpamGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailTopAseSpamGetParamsFormat string

const (
	RadarEmailTopAseSpamGetParamsFormatJson RadarEmailTopAseSpamGetParamsFormat = "JSON"
	RadarEmailTopAseSpamGetParamsFormatCsv  RadarEmailTopAseSpamGetParamsFormat = "CSV"
)

type RadarEmailTopAseSpamGetParamsSpf string

const (
	RadarEmailTopAseSpamGetParamsSpfPass RadarEmailTopAseSpamGetParamsSpf = "PASS"
	RadarEmailTopAseSpamGetParamsSpfNone RadarEmailTopAseSpamGetParamsSpf = "NONE"
	RadarEmailTopAseSpamGetParamsSpfFail RadarEmailTopAseSpamGetParamsSpf = "FAIL"
)
