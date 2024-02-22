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

// RadarEmailSecurityTopAseSpamService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopAseSpamService] method instead.
type RadarEmailSecurityTopAseSpamService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopAseSpamService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopAseSpamService(opts ...option.RequestOption) (r *RadarEmailSecurityTopAseSpamService) {
	r = &RadarEmailSecurityTopAseSpamService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by emails classified, of Spam validations.
func (r *RadarEmailSecurityTopAseSpamService) Get(ctx context.Context, spam RadarEmailSecurityTopAseSpamGetParamsSpam, query RadarEmailSecurityTopAseSpamGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopAseSpamGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopAseSpamGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/ases/spam/%v", spam)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopAseSpamGetResponse struct {
	Meta RadarEmailSecurityTopAseSpamGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopAseSpamGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopAseSpamGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopAseSpamGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopAseSpamGetResponse]
type radarEmailSecurityTopAseSpamGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseSpamGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseSpamGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopAseSpamGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                    `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopAseSpamGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopAseSpamGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopAseSpamGetResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseSpamGetResponseMeta]
type radarEmailSecurityTopAseSpamGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseSpamGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseSpamGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopAseSpamGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopAseSpamGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopAseSpamGetResponseMetaDateRange]
type radarEmailSecurityTopAseSpamGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseSpamGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseSpamGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopAseSpamGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                 `json:"level"`
	JSON        radarEmailSecurityTopAseSpamGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopAseSpamGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopAseSpamGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopAseSpamGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseSpamGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseSpamGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                  `json:"dataSource,required"`
	Description     string                                                                  `json:"description,required"`
	EventType       string                                                                  `json:"eventType,required"`
	IsInstantaneous interface{}                                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                                               `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                  `json:"linkedUrl"`
	StartTime       time.Time                                                               `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopAseSpamGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopAseSpamGetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTopAseSpamGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopAseSpamGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopAseSpamGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseSpamGetResponseTop0 struct {
	ClientASN    int64                                           `json:"clientASN,required"`
	ClientAsName string                                          `json:"clientASName,required"`
	Value        string                                          `json:"value,required"`
	JSON         radarEmailSecurityTopAseSpamGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopAseSpamGetResponseTop0JSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseSpamGetResponseTop0]
type radarEmailSecurityTopAseSpamGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseSpamGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseSpamGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTopAseSpamGetParamsARC] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopAseSpamGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopAseSpamGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTopAseSpamGetParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopAseSpamGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopAseSpamGetParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopAseSpamGetParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecurityTopAseSpamGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Spam.
type RadarEmailSecurityTopAseSpamGetParamsSpam string

const (
	RadarEmailSecurityTopAseSpamGetParamsSpamSpam    RadarEmailSecurityTopAseSpamGetParamsSpam = "SPAM"
	RadarEmailSecurityTopAseSpamGetParamsSpamNotSpam RadarEmailSecurityTopAseSpamGetParamsSpam = "NOT_SPAM"
)

type RadarEmailSecurityTopAseSpamGetParamsARC string

const (
	RadarEmailSecurityTopAseSpamGetParamsARCPass RadarEmailSecurityTopAseSpamGetParamsARC = "PASS"
	RadarEmailSecurityTopAseSpamGetParamsARCNone RadarEmailSecurityTopAseSpamGetParamsARC = "NONE"
	RadarEmailSecurityTopAseSpamGetParamsARCFail RadarEmailSecurityTopAseSpamGetParamsARC = "FAIL"
)

type RadarEmailSecurityTopAseSpamGetParamsDateRange string

const (
	RadarEmailSecurityTopAseSpamGetParamsDateRange1d         RadarEmailSecurityTopAseSpamGetParamsDateRange = "1d"
	RadarEmailSecurityTopAseSpamGetParamsDateRange2d         RadarEmailSecurityTopAseSpamGetParamsDateRange = "2d"
	RadarEmailSecurityTopAseSpamGetParamsDateRange7d         RadarEmailSecurityTopAseSpamGetParamsDateRange = "7d"
	RadarEmailSecurityTopAseSpamGetParamsDateRange14d        RadarEmailSecurityTopAseSpamGetParamsDateRange = "14d"
	RadarEmailSecurityTopAseSpamGetParamsDateRange28d        RadarEmailSecurityTopAseSpamGetParamsDateRange = "28d"
	RadarEmailSecurityTopAseSpamGetParamsDateRange12w        RadarEmailSecurityTopAseSpamGetParamsDateRange = "12w"
	RadarEmailSecurityTopAseSpamGetParamsDateRange24w        RadarEmailSecurityTopAseSpamGetParamsDateRange = "24w"
	RadarEmailSecurityTopAseSpamGetParamsDateRange52w        RadarEmailSecurityTopAseSpamGetParamsDateRange = "52w"
	RadarEmailSecurityTopAseSpamGetParamsDateRange1dControl  RadarEmailSecurityTopAseSpamGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopAseSpamGetParamsDateRange2dControl  RadarEmailSecurityTopAseSpamGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopAseSpamGetParamsDateRange7dControl  RadarEmailSecurityTopAseSpamGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopAseSpamGetParamsDateRange14dControl RadarEmailSecurityTopAseSpamGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopAseSpamGetParamsDateRange28dControl RadarEmailSecurityTopAseSpamGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopAseSpamGetParamsDateRange12wControl RadarEmailSecurityTopAseSpamGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopAseSpamGetParamsDateRange24wControl RadarEmailSecurityTopAseSpamGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopAseSpamGetParamsDKIM string

const (
	RadarEmailSecurityTopAseSpamGetParamsDKIMPass RadarEmailSecurityTopAseSpamGetParamsDKIM = "PASS"
	RadarEmailSecurityTopAseSpamGetParamsDKIMNone RadarEmailSecurityTopAseSpamGetParamsDKIM = "NONE"
	RadarEmailSecurityTopAseSpamGetParamsDKIMFail RadarEmailSecurityTopAseSpamGetParamsDKIM = "FAIL"
)

type RadarEmailSecurityTopAseSpamGetParamsDMARC string

const (
	RadarEmailSecurityTopAseSpamGetParamsDMARCPass RadarEmailSecurityTopAseSpamGetParamsDMARC = "PASS"
	RadarEmailSecurityTopAseSpamGetParamsDMARCNone RadarEmailSecurityTopAseSpamGetParamsDMARC = "NONE"
	RadarEmailSecurityTopAseSpamGetParamsDMARCFail RadarEmailSecurityTopAseSpamGetParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopAseSpamGetParamsFormat string

const (
	RadarEmailSecurityTopAseSpamGetParamsFormatJson RadarEmailSecurityTopAseSpamGetParamsFormat = "JSON"
	RadarEmailSecurityTopAseSpamGetParamsFormatCsv  RadarEmailSecurityTopAseSpamGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopAseSpamGetParamsSPF string

const (
	RadarEmailSecurityTopAseSpamGetParamsSPFPass RadarEmailSecurityTopAseSpamGetParamsSPF = "PASS"
	RadarEmailSecurityTopAseSpamGetParamsSPFNone RadarEmailSecurityTopAseSpamGetParamsSPF = "NONE"
	RadarEmailSecurityTopAseSpamGetParamsSPFFail RadarEmailSecurityTopAseSpamGetParamsSPF = "FAIL"
)

type RadarEmailSecurityTopAseSpamGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopAseSpamGetResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarEmailSecurityTopAseSpamGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopAseSpamGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopAseSpamGetResponseEnvelope]
type radarEmailSecurityTopAseSpamGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseSpamGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
