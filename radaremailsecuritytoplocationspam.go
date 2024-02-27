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

// RadarEmailSecurityTopLocationSpamService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopLocationSpamService] method instead.
type RadarEmailSecurityTopLocationSpamService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopLocationSpamService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopLocationSpamService(opts ...option.RequestOption) (r *RadarEmailSecurityTopLocationSpamService) {
	r = &RadarEmailSecurityTopLocationSpamService{}
	r.Options = opts
	return
}

// Get the top locations by emails classified as Spam or not.
func (r *RadarEmailSecurityTopLocationSpamService) Get(ctx context.Context, spam RadarEmailSecurityTopLocationSpamGetParamsSpam, query RadarEmailSecurityTopLocationSpamGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopLocationSpamGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopLocationSpamGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/locations/spam/%v", spam)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopLocationSpamGetResponse struct {
	Meta RadarEmailSecurityTopLocationSpamGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopLocationSpamGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopLocationSpamGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopLocationSpamGetResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopLocationSpamGetResponse]
type radarEmailSecurityTopLocationSpamGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationSpamGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationSpamGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopLocationSpamGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                         `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopLocationSpamGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopLocationSpamGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopLocationSpamGetResponseMetaJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationSpamGetResponseMeta]
type radarEmailSecurityTopLocationSpamGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationSpamGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationSpamGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopLocationSpamGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopLocationSpamGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopLocationSpamGetResponseMetaDateRange]
type radarEmailSecurityTopLocationSpamGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationSpamGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationSpamGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopLocationSpamGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                      `json:"level"`
	JSON        radarEmailSecurityTopLocationSpamGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopLocationSpamGetResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopLocationSpamGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopLocationSpamGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationSpamGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationSpamGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                       `json:"dataSource,required"`
	Description     string                                                                       `json:"description,required"`
	EventType       string                                                                       `json:"eventType,required"`
	IsInstantaneous interface{}                                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                                    `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopLocationSpamGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopLocationSpamGetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationSpamGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopLocationSpamGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopLocationSpamGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationSpamGetResponseTop0 struct {
	ClientCountryAlpha2 string                                               `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                               `json:"clientCountryName,required"`
	Value               string                                               `json:"value,required"`
	JSON                radarEmailSecurityTopLocationSpamGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopLocationSpamGetResponseTop0JSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationSpamGetResponseTop0]
type radarEmailSecurityTopLocationSpamGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationSpamGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationSpamGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTopLocationSpamGetParamsARC] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopLocationSpamGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopLocationSpamGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTopLocationSpamGetParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopLocationSpamGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopLocationSpamGetParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopLocationSpamGetParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTopLocationSpamGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Spam.
type RadarEmailSecurityTopLocationSpamGetParamsSpam string

const (
	RadarEmailSecurityTopLocationSpamGetParamsSpamSpam    RadarEmailSecurityTopLocationSpamGetParamsSpam = "SPAM"
	RadarEmailSecurityTopLocationSpamGetParamsSpamNotSpam RadarEmailSecurityTopLocationSpamGetParamsSpam = "NOT_SPAM"
)

type RadarEmailSecurityTopLocationSpamGetParamsARC string

const (
	RadarEmailSecurityTopLocationSpamGetParamsARCPass RadarEmailSecurityTopLocationSpamGetParamsARC = "PASS"
	RadarEmailSecurityTopLocationSpamGetParamsARCNone RadarEmailSecurityTopLocationSpamGetParamsARC = "NONE"
	RadarEmailSecurityTopLocationSpamGetParamsARCFail RadarEmailSecurityTopLocationSpamGetParamsARC = "FAIL"
)

type RadarEmailSecurityTopLocationSpamGetParamsDateRange string

const (
	RadarEmailSecurityTopLocationSpamGetParamsDateRange1d         RadarEmailSecurityTopLocationSpamGetParamsDateRange = "1d"
	RadarEmailSecurityTopLocationSpamGetParamsDateRange2d         RadarEmailSecurityTopLocationSpamGetParamsDateRange = "2d"
	RadarEmailSecurityTopLocationSpamGetParamsDateRange7d         RadarEmailSecurityTopLocationSpamGetParamsDateRange = "7d"
	RadarEmailSecurityTopLocationSpamGetParamsDateRange14d        RadarEmailSecurityTopLocationSpamGetParamsDateRange = "14d"
	RadarEmailSecurityTopLocationSpamGetParamsDateRange28d        RadarEmailSecurityTopLocationSpamGetParamsDateRange = "28d"
	RadarEmailSecurityTopLocationSpamGetParamsDateRange12w        RadarEmailSecurityTopLocationSpamGetParamsDateRange = "12w"
	RadarEmailSecurityTopLocationSpamGetParamsDateRange24w        RadarEmailSecurityTopLocationSpamGetParamsDateRange = "24w"
	RadarEmailSecurityTopLocationSpamGetParamsDateRange52w        RadarEmailSecurityTopLocationSpamGetParamsDateRange = "52w"
	RadarEmailSecurityTopLocationSpamGetParamsDateRange1dControl  RadarEmailSecurityTopLocationSpamGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopLocationSpamGetParamsDateRange2dControl  RadarEmailSecurityTopLocationSpamGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopLocationSpamGetParamsDateRange7dControl  RadarEmailSecurityTopLocationSpamGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopLocationSpamGetParamsDateRange14dControl RadarEmailSecurityTopLocationSpamGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopLocationSpamGetParamsDateRange28dControl RadarEmailSecurityTopLocationSpamGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopLocationSpamGetParamsDateRange12wControl RadarEmailSecurityTopLocationSpamGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopLocationSpamGetParamsDateRange24wControl RadarEmailSecurityTopLocationSpamGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopLocationSpamGetParamsDKIM string

const (
	RadarEmailSecurityTopLocationSpamGetParamsDKIMPass RadarEmailSecurityTopLocationSpamGetParamsDKIM = "PASS"
	RadarEmailSecurityTopLocationSpamGetParamsDKIMNone RadarEmailSecurityTopLocationSpamGetParamsDKIM = "NONE"
	RadarEmailSecurityTopLocationSpamGetParamsDKIMFail RadarEmailSecurityTopLocationSpamGetParamsDKIM = "FAIL"
)

type RadarEmailSecurityTopLocationSpamGetParamsDMARC string

const (
	RadarEmailSecurityTopLocationSpamGetParamsDMARCPass RadarEmailSecurityTopLocationSpamGetParamsDMARC = "PASS"
	RadarEmailSecurityTopLocationSpamGetParamsDMARCNone RadarEmailSecurityTopLocationSpamGetParamsDMARC = "NONE"
	RadarEmailSecurityTopLocationSpamGetParamsDMARCFail RadarEmailSecurityTopLocationSpamGetParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopLocationSpamGetParamsFormat string

const (
	RadarEmailSecurityTopLocationSpamGetParamsFormatJson RadarEmailSecurityTopLocationSpamGetParamsFormat = "JSON"
	RadarEmailSecurityTopLocationSpamGetParamsFormatCsv  RadarEmailSecurityTopLocationSpamGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopLocationSpamGetParamsSPF string

const (
	RadarEmailSecurityTopLocationSpamGetParamsSPFPass RadarEmailSecurityTopLocationSpamGetParamsSPF = "PASS"
	RadarEmailSecurityTopLocationSpamGetParamsSPFNone RadarEmailSecurityTopLocationSpamGetParamsSPF = "NONE"
	RadarEmailSecurityTopLocationSpamGetParamsSPFFail RadarEmailSecurityTopLocationSpamGetParamsSPF = "FAIL"
)

type RadarEmailSecurityTopLocationSpamGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopLocationSpamGetResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarEmailSecurityTopLocationSpamGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopLocationSpamGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopLocationSpamGetResponseEnvelope]
type radarEmailSecurityTopLocationSpamGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationSpamGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
