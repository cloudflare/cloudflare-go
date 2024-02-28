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

// RadarEmailSecurityTopTldSpamService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopTldSpamService] method instead.
type RadarEmailSecurityTopTldSpamService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopTldSpamService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopTldSpamService(opts ...option.RequestOption) (r *RadarEmailSecurityTopTldSpamService) {
	r = &RadarEmailSecurityTopTldSpamService{}
	r.Options = opts
	return
}

// Get the top TLDs by emails classified as Spam or not.
func (r *RadarEmailSecurityTopTldSpamService) Get(ctx context.Context, spam RadarEmailSecurityTopTldSpamGetParamsSpam, query RadarEmailSecurityTopTldSpamGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopTldSpamGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopTldSpamGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/tlds/spam/%v", spam)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopTldSpamGetResponse struct {
	Meta RadarEmailSecurityTopTldSpamGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopTldSpamGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopTldSpamGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopTldSpamGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopTldSpamGetResponse]
type radarEmailSecurityTopTldSpamGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldSpamGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopTldSpamGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopTldSpamGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                    `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopTldSpamGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopTldSpamGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopTldSpamGetResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopTldSpamGetResponseMeta]
type radarEmailSecurityTopTldSpamGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldSpamGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopTldSpamGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopTldSpamGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopTldSpamGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopTldSpamGetResponseMetaDateRange]
type radarEmailSecurityTopTldSpamGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldSpamGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopTldSpamGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                 `json:"level"`
	JSON        radarEmailSecurityTopTldSpamGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopTldSpamGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopTldSpamGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopTldSpamGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldSpamGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                  `json:"dataSource,required"`
	Description     string                                                                  `json:"description,required"`
	EventType       string                                                                  `json:"eventType,required"`
	IsInstantaneous interface{}                                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                                               `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                  `json:"linkedUrl"`
	StartTime       time.Time                                                               `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopTldSpamGetResponseTop0 struct {
	Name  string                                          `json:"name,required"`
	Value string                                          `json:"value,required"`
	JSON  radarEmailSecurityTopTldSpamGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopTldSpamGetResponseTop0JSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopTldSpamGetResponseTop0]
type radarEmailSecurityTopTldSpamGetResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldSpamGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopTldSpamGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTopTldSpamGetParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopTldSpamGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopTldSpamGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTopTldSpamGetParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopTldSpamGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopTldSpamGetParamsSPF] `query:"spf"`
	// Filter for TLDs by category.
	TldCategory param.Field[RadarEmailSecurityTopTldSpamGetParamsTldCategory] `query:"tldCategory"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecurityTopTldSpamGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecurityTopTldSpamGetParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecurityTopTldSpamGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Spam.
type RadarEmailSecurityTopTldSpamGetParamsSpam string

const (
	RadarEmailSecurityTopTldSpamGetParamsSpamSpam    RadarEmailSecurityTopTldSpamGetParamsSpam = "SPAM"
	RadarEmailSecurityTopTldSpamGetParamsSpamNotSpam RadarEmailSecurityTopTldSpamGetParamsSpam = "NOT_SPAM"
)

type RadarEmailSecurityTopTldSpamGetParamsARC string

const (
	RadarEmailSecurityTopTldSpamGetParamsARCPass RadarEmailSecurityTopTldSpamGetParamsARC = "PASS"
	RadarEmailSecurityTopTldSpamGetParamsARCNone RadarEmailSecurityTopTldSpamGetParamsARC = "NONE"
	RadarEmailSecurityTopTldSpamGetParamsARCFail RadarEmailSecurityTopTldSpamGetParamsARC = "FAIL"
)

type RadarEmailSecurityTopTldSpamGetParamsDateRange string

const (
	RadarEmailSecurityTopTldSpamGetParamsDateRange1d         RadarEmailSecurityTopTldSpamGetParamsDateRange = "1d"
	RadarEmailSecurityTopTldSpamGetParamsDateRange2d         RadarEmailSecurityTopTldSpamGetParamsDateRange = "2d"
	RadarEmailSecurityTopTldSpamGetParamsDateRange7d         RadarEmailSecurityTopTldSpamGetParamsDateRange = "7d"
	RadarEmailSecurityTopTldSpamGetParamsDateRange14d        RadarEmailSecurityTopTldSpamGetParamsDateRange = "14d"
	RadarEmailSecurityTopTldSpamGetParamsDateRange28d        RadarEmailSecurityTopTldSpamGetParamsDateRange = "28d"
	RadarEmailSecurityTopTldSpamGetParamsDateRange12w        RadarEmailSecurityTopTldSpamGetParamsDateRange = "12w"
	RadarEmailSecurityTopTldSpamGetParamsDateRange24w        RadarEmailSecurityTopTldSpamGetParamsDateRange = "24w"
	RadarEmailSecurityTopTldSpamGetParamsDateRange52w        RadarEmailSecurityTopTldSpamGetParamsDateRange = "52w"
	RadarEmailSecurityTopTldSpamGetParamsDateRange1dControl  RadarEmailSecurityTopTldSpamGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopTldSpamGetParamsDateRange2dControl  RadarEmailSecurityTopTldSpamGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopTldSpamGetParamsDateRange7dControl  RadarEmailSecurityTopTldSpamGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopTldSpamGetParamsDateRange14dControl RadarEmailSecurityTopTldSpamGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopTldSpamGetParamsDateRange28dControl RadarEmailSecurityTopTldSpamGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopTldSpamGetParamsDateRange12wControl RadarEmailSecurityTopTldSpamGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopTldSpamGetParamsDateRange24wControl RadarEmailSecurityTopTldSpamGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopTldSpamGetParamsDKIM string

const (
	RadarEmailSecurityTopTldSpamGetParamsDKIMPass RadarEmailSecurityTopTldSpamGetParamsDKIM = "PASS"
	RadarEmailSecurityTopTldSpamGetParamsDKIMNone RadarEmailSecurityTopTldSpamGetParamsDKIM = "NONE"
	RadarEmailSecurityTopTldSpamGetParamsDKIMFail RadarEmailSecurityTopTldSpamGetParamsDKIM = "FAIL"
)

type RadarEmailSecurityTopTldSpamGetParamsDMARC string

const (
	RadarEmailSecurityTopTldSpamGetParamsDMARCPass RadarEmailSecurityTopTldSpamGetParamsDMARC = "PASS"
	RadarEmailSecurityTopTldSpamGetParamsDMARCNone RadarEmailSecurityTopTldSpamGetParamsDMARC = "NONE"
	RadarEmailSecurityTopTldSpamGetParamsDMARCFail RadarEmailSecurityTopTldSpamGetParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopTldSpamGetParamsFormat string

const (
	RadarEmailSecurityTopTldSpamGetParamsFormatJson RadarEmailSecurityTopTldSpamGetParamsFormat = "JSON"
	RadarEmailSecurityTopTldSpamGetParamsFormatCsv  RadarEmailSecurityTopTldSpamGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopTldSpamGetParamsSPF string

const (
	RadarEmailSecurityTopTldSpamGetParamsSPFPass RadarEmailSecurityTopTldSpamGetParamsSPF = "PASS"
	RadarEmailSecurityTopTldSpamGetParamsSPFNone RadarEmailSecurityTopTldSpamGetParamsSPF = "NONE"
	RadarEmailSecurityTopTldSpamGetParamsSPFFail RadarEmailSecurityTopTldSpamGetParamsSPF = "FAIL"
)

// Filter for TLDs by category.
type RadarEmailSecurityTopTldSpamGetParamsTldCategory string

const (
	RadarEmailSecurityTopTldSpamGetParamsTldCategoryClassic RadarEmailSecurityTopTldSpamGetParamsTldCategory = "CLASSIC"
	RadarEmailSecurityTopTldSpamGetParamsTldCategoryCountry RadarEmailSecurityTopTldSpamGetParamsTldCategory = "COUNTRY"
)

type RadarEmailSecurityTopTldSpamGetParamsTLSVersion string

const (
	RadarEmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_0 RadarEmailSecurityTopTldSpamGetParamsTLSVersion = "TLSv1_0"
	RadarEmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_1 RadarEmailSecurityTopTldSpamGetParamsTLSVersion = "TLSv1_1"
	RadarEmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_2 RadarEmailSecurityTopTldSpamGetParamsTLSVersion = "TLSv1_2"
	RadarEmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_3 RadarEmailSecurityTopTldSpamGetParamsTLSVersion = "TLSv1_3"
)

type RadarEmailSecurityTopTldSpamGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopTldSpamGetResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarEmailSecurityTopTldSpamGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopTldSpamGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopTldSpamGetResponseEnvelope]
type radarEmailSecurityTopTldSpamGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldSpamGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
