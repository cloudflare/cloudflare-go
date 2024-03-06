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

// RadarEmailSecurityTopTldService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopTldService] method instead.
type RadarEmailSecurityTopTldService struct {
	Options   []option.RequestOption
	Malicious *RadarEmailSecurityTopTldMaliciousService
	Spam      *RadarEmailSecurityTopTldSpamService
	Spoof     *RadarEmailSecurityTopTldSpoofService
}

// NewRadarEmailSecurityTopTldService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopTldService(opts ...option.RequestOption) (r *RadarEmailSecurityTopTldService) {
	r = &RadarEmailSecurityTopTldService{}
	r.Options = opts
	r.Malicious = NewRadarEmailSecurityTopTldMaliciousService(opts...)
	r.Spam = NewRadarEmailSecurityTopTldSpamService(opts...)
	r.Spoof = NewRadarEmailSecurityTopTldSpoofService(opts...)
	return
}

// Get the top TLDs by email messages. Values are a percentage out of the total
// emails.
func (r *RadarEmailSecurityTopTldService) Get(ctx context.Context, query RadarEmailSecurityTopTldGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopTldGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopTldGetResponseEnvelope
	path := "radar/email/security/top/tlds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopTldGetResponse struct {
	Meta RadarEmailSecurityTopTldGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopTldGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopTldGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopTldGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopTldGetResponse]
type radarEmailSecurityTopTldGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopTldGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopTldGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopTldGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopTldGetResponseMetaJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopTldGetResponseMeta]
type radarEmailSecurityTopTldGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopTldGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopTldGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopTldGetResponseMetaDateRange]
type radarEmailSecurityTopTldGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopTldGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                             `json:"level"`
	JSON        radarEmailSecurityTopTldGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopTldGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopTldGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopTldGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                              `json:"dataSource,required"`
	Description     string                                                              `json:"description,required"`
	EventType       string                                                              `json:"eventType,required"`
	IsInstantaneous interface{}                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                           `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopTldGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopTldGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopTldGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopTldGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopTldGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetResponseTop0 struct {
	Name  string                                      `json:"name,required"`
	Value string                                      `json:"value,required"`
	JSON  radarEmailSecurityTopTldGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopTldGetResponseTop0JSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopTldGetResponseTop0]
type radarEmailSecurityTopTldGetResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTopTldGetParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopTldGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopTldGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTopTldGetParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopTldGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopTldGetParamsSPF] `query:"spf"`
	// Filter for TLDs by category.
	TldCategory param.Field[RadarEmailSecurityTopTldGetParamsTldCategory] `query:"tldCategory"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecurityTopTldGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecurityTopTldGetParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecurityTopTldGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecurityTopTldGetParamsARC string

const (
	RadarEmailSecurityTopTldGetParamsARCPass RadarEmailSecurityTopTldGetParamsARC = "PASS"
	RadarEmailSecurityTopTldGetParamsARCNone RadarEmailSecurityTopTldGetParamsARC = "NONE"
	RadarEmailSecurityTopTldGetParamsARCFail RadarEmailSecurityTopTldGetParamsARC = "FAIL"
)

type RadarEmailSecurityTopTldGetParamsDateRange string

const (
	RadarEmailSecurityTopTldGetParamsDateRange1d         RadarEmailSecurityTopTldGetParamsDateRange = "1d"
	RadarEmailSecurityTopTldGetParamsDateRange2d         RadarEmailSecurityTopTldGetParamsDateRange = "2d"
	RadarEmailSecurityTopTldGetParamsDateRange7d         RadarEmailSecurityTopTldGetParamsDateRange = "7d"
	RadarEmailSecurityTopTldGetParamsDateRange14d        RadarEmailSecurityTopTldGetParamsDateRange = "14d"
	RadarEmailSecurityTopTldGetParamsDateRange28d        RadarEmailSecurityTopTldGetParamsDateRange = "28d"
	RadarEmailSecurityTopTldGetParamsDateRange12w        RadarEmailSecurityTopTldGetParamsDateRange = "12w"
	RadarEmailSecurityTopTldGetParamsDateRange24w        RadarEmailSecurityTopTldGetParamsDateRange = "24w"
	RadarEmailSecurityTopTldGetParamsDateRange52w        RadarEmailSecurityTopTldGetParamsDateRange = "52w"
	RadarEmailSecurityTopTldGetParamsDateRange1dControl  RadarEmailSecurityTopTldGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopTldGetParamsDateRange2dControl  RadarEmailSecurityTopTldGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopTldGetParamsDateRange7dControl  RadarEmailSecurityTopTldGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopTldGetParamsDateRange14dControl RadarEmailSecurityTopTldGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopTldGetParamsDateRange28dControl RadarEmailSecurityTopTldGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopTldGetParamsDateRange12wControl RadarEmailSecurityTopTldGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopTldGetParamsDateRange24wControl RadarEmailSecurityTopTldGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopTldGetParamsDKIM string

const (
	RadarEmailSecurityTopTldGetParamsDKIMPass RadarEmailSecurityTopTldGetParamsDKIM = "PASS"
	RadarEmailSecurityTopTldGetParamsDKIMNone RadarEmailSecurityTopTldGetParamsDKIM = "NONE"
	RadarEmailSecurityTopTldGetParamsDKIMFail RadarEmailSecurityTopTldGetParamsDKIM = "FAIL"
)

type RadarEmailSecurityTopTldGetParamsDMARC string

const (
	RadarEmailSecurityTopTldGetParamsDMARCPass RadarEmailSecurityTopTldGetParamsDMARC = "PASS"
	RadarEmailSecurityTopTldGetParamsDMARCNone RadarEmailSecurityTopTldGetParamsDMARC = "NONE"
	RadarEmailSecurityTopTldGetParamsDMARCFail RadarEmailSecurityTopTldGetParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopTldGetParamsFormat string

const (
	RadarEmailSecurityTopTldGetParamsFormatJson RadarEmailSecurityTopTldGetParamsFormat = "JSON"
	RadarEmailSecurityTopTldGetParamsFormatCsv  RadarEmailSecurityTopTldGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopTldGetParamsSPF string

const (
	RadarEmailSecurityTopTldGetParamsSPFPass RadarEmailSecurityTopTldGetParamsSPF = "PASS"
	RadarEmailSecurityTopTldGetParamsSPFNone RadarEmailSecurityTopTldGetParamsSPF = "NONE"
	RadarEmailSecurityTopTldGetParamsSPFFail RadarEmailSecurityTopTldGetParamsSPF = "FAIL"
)

// Filter for TLDs by category.
type RadarEmailSecurityTopTldGetParamsTldCategory string

const (
	RadarEmailSecurityTopTldGetParamsTldCategoryClassic RadarEmailSecurityTopTldGetParamsTldCategory = "CLASSIC"
	RadarEmailSecurityTopTldGetParamsTldCategoryCountry RadarEmailSecurityTopTldGetParamsTldCategory = "COUNTRY"
)

type RadarEmailSecurityTopTldGetParamsTLSVersion string

const (
	RadarEmailSecurityTopTldGetParamsTLSVersionTlSv1_0 RadarEmailSecurityTopTldGetParamsTLSVersion = "TLSv1_0"
	RadarEmailSecurityTopTldGetParamsTLSVersionTlSv1_1 RadarEmailSecurityTopTldGetParamsTLSVersion = "TLSv1_1"
	RadarEmailSecurityTopTldGetParamsTLSVersionTlSv1_2 RadarEmailSecurityTopTldGetParamsTLSVersion = "TLSv1_2"
	RadarEmailSecurityTopTldGetParamsTLSVersionTlSv1_3 RadarEmailSecurityTopTldGetParamsTLSVersion = "TLSv1_3"
)

type RadarEmailSecurityTopTldGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopTldGetResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarEmailSecurityTopTldGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopTldGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopTldGetResponseEnvelope]
type radarEmailSecurityTopTldGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
