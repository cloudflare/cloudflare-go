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

// RadarEmailSecurityTopTldMaliciousService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopTldMaliciousService] method instead.
type RadarEmailSecurityTopTldMaliciousService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopTldMaliciousService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopTldMaliciousService(opts ...option.RequestOption) (r *RadarEmailSecurityTopTldMaliciousService) {
	r = &RadarEmailSecurityTopTldMaliciousService{}
	r.Options = opts
	return
}

// Get the TLDs by emails classified as malicious or not.
func (r *RadarEmailSecurityTopTldMaliciousService) Get(ctx context.Context, malicious RadarEmailSecurityTopTldMaliciousGetParamsMalicious, query RadarEmailSecurityTopTldMaliciousGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopTldMaliciousGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopTldMaliciousGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/tlds/malicious/%v", malicious)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopTldMaliciousGetResponse struct {
	Meta RadarEmailSecurityTopTldMaliciousGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopTldMaliciousGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopTldMaliciousGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopTldMaliciousGetResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopTldMaliciousGetResponse]
type radarEmailSecurityTopTldMaliciousGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldMaliciousGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopTldMaliciousGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopTldMaliciousGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                         `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopTldMaliciousGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopTldMaliciousGetResponseMetaJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopTldMaliciousGetResponseMeta]
type radarEmailSecurityTopTldMaliciousGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldMaliciousGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopTldMaliciousGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopTldMaliciousGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopTldMaliciousGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopTldMaliciousGetResponseMetaDateRange]
type radarEmailSecurityTopTldMaliciousGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldMaliciousGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                      `json:"level"`
	JSON        radarEmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                       `json:"dataSource,required"`
	Description     string                                                                       `json:"description,required"`
	EventType       string                                                                       `json:"eventType,required"`
	IsInstantaneous interface{}                                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                                    `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopTldMaliciousGetResponseTop0 struct {
	Name  string                                               `json:"name,required"`
	Value string                                               `json:"value,required"`
	JSON  radarEmailSecurityTopTldMaliciousGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopTldMaliciousGetResponseTop0JSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopTldMaliciousGetResponseTop0]
type radarEmailSecurityTopTldMaliciousGetResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldMaliciousGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopTldMaliciousGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTopTldMaliciousGetParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopTldMaliciousGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopTldMaliciousGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTopTldMaliciousGetParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopTldMaliciousGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopTldMaliciousGetParamsSPF] `query:"spf"`
	// Filter for TLDs by category.
	TldCategory param.Field[RadarEmailSecurityTopTldMaliciousGetParamsTldCategory] `query:"tldCategory"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecurityTopTldMaliciousGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecurityTopTldMaliciousGetParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTopTldMaliciousGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Malicious.
type RadarEmailSecurityTopTldMaliciousGetParamsMalicious string

const (
	RadarEmailSecurityTopTldMaliciousGetParamsMaliciousMalicious    RadarEmailSecurityTopTldMaliciousGetParamsMalicious = "MALICIOUS"
	RadarEmailSecurityTopTldMaliciousGetParamsMaliciousNotMalicious RadarEmailSecurityTopTldMaliciousGetParamsMalicious = "NOT_MALICIOUS"
)

type RadarEmailSecurityTopTldMaliciousGetParamsARC string

const (
	RadarEmailSecurityTopTldMaliciousGetParamsARCPass RadarEmailSecurityTopTldMaliciousGetParamsARC = "PASS"
	RadarEmailSecurityTopTldMaliciousGetParamsARCNone RadarEmailSecurityTopTldMaliciousGetParamsARC = "NONE"
	RadarEmailSecurityTopTldMaliciousGetParamsARCFail RadarEmailSecurityTopTldMaliciousGetParamsARC = "FAIL"
)

type RadarEmailSecurityTopTldMaliciousGetParamsDateRange string

const (
	RadarEmailSecurityTopTldMaliciousGetParamsDateRange1d         RadarEmailSecurityTopTldMaliciousGetParamsDateRange = "1d"
	RadarEmailSecurityTopTldMaliciousGetParamsDateRange2d         RadarEmailSecurityTopTldMaliciousGetParamsDateRange = "2d"
	RadarEmailSecurityTopTldMaliciousGetParamsDateRange7d         RadarEmailSecurityTopTldMaliciousGetParamsDateRange = "7d"
	RadarEmailSecurityTopTldMaliciousGetParamsDateRange14d        RadarEmailSecurityTopTldMaliciousGetParamsDateRange = "14d"
	RadarEmailSecurityTopTldMaliciousGetParamsDateRange28d        RadarEmailSecurityTopTldMaliciousGetParamsDateRange = "28d"
	RadarEmailSecurityTopTldMaliciousGetParamsDateRange12w        RadarEmailSecurityTopTldMaliciousGetParamsDateRange = "12w"
	RadarEmailSecurityTopTldMaliciousGetParamsDateRange24w        RadarEmailSecurityTopTldMaliciousGetParamsDateRange = "24w"
	RadarEmailSecurityTopTldMaliciousGetParamsDateRange52w        RadarEmailSecurityTopTldMaliciousGetParamsDateRange = "52w"
	RadarEmailSecurityTopTldMaliciousGetParamsDateRange1dControl  RadarEmailSecurityTopTldMaliciousGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopTldMaliciousGetParamsDateRange2dControl  RadarEmailSecurityTopTldMaliciousGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopTldMaliciousGetParamsDateRange7dControl  RadarEmailSecurityTopTldMaliciousGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopTldMaliciousGetParamsDateRange14dControl RadarEmailSecurityTopTldMaliciousGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopTldMaliciousGetParamsDateRange28dControl RadarEmailSecurityTopTldMaliciousGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopTldMaliciousGetParamsDateRange12wControl RadarEmailSecurityTopTldMaliciousGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopTldMaliciousGetParamsDateRange24wControl RadarEmailSecurityTopTldMaliciousGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopTldMaliciousGetParamsDKIM string

const (
	RadarEmailSecurityTopTldMaliciousGetParamsDKIMPass RadarEmailSecurityTopTldMaliciousGetParamsDKIM = "PASS"
	RadarEmailSecurityTopTldMaliciousGetParamsDKIMNone RadarEmailSecurityTopTldMaliciousGetParamsDKIM = "NONE"
	RadarEmailSecurityTopTldMaliciousGetParamsDKIMFail RadarEmailSecurityTopTldMaliciousGetParamsDKIM = "FAIL"
)

type RadarEmailSecurityTopTldMaliciousGetParamsDMARC string

const (
	RadarEmailSecurityTopTldMaliciousGetParamsDMARCPass RadarEmailSecurityTopTldMaliciousGetParamsDMARC = "PASS"
	RadarEmailSecurityTopTldMaliciousGetParamsDMARCNone RadarEmailSecurityTopTldMaliciousGetParamsDMARC = "NONE"
	RadarEmailSecurityTopTldMaliciousGetParamsDMARCFail RadarEmailSecurityTopTldMaliciousGetParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopTldMaliciousGetParamsFormat string

const (
	RadarEmailSecurityTopTldMaliciousGetParamsFormatJson RadarEmailSecurityTopTldMaliciousGetParamsFormat = "JSON"
	RadarEmailSecurityTopTldMaliciousGetParamsFormatCsv  RadarEmailSecurityTopTldMaliciousGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopTldMaliciousGetParamsSPF string

const (
	RadarEmailSecurityTopTldMaliciousGetParamsSPFPass RadarEmailSecurityTopTldMaliciousGetParamsSPF = "PASS"
	RadarEmailSecurityTopTldMaliciousGetParamsSPFNone RadarEmailSecurityTopTldMaliciousGetParamsSPF = "NONE"
	RadarEmailSecurityTopTldMaliciousGetParamsSPFFail RadarEmailSecurityTopTldMaliciousGetParamsSPF = "FAIL"
)

// Filter for TLDs by category.
type RadarEmailSecurityTopTldMaliciousGetParamsTldCategory string

const (
	RadarEmailSecurityTopTldMaliciousGetParamsTldCategoryClassic RadarEmailSecurityTopTldMaliciousGetParamsTldCategory = "CLASSIC"
	RadarEmailSecurityTopTldMaliciousGetParamsTldCategoryCountry RadarEmailSecurityTopTldMaliciousGetParamsTldCategory = "COUNTRY"
)

type RadarEmailSecurityTopTldMaliciousGetParamsTLSVersion string

const (
	RadarEmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_0 RadarEmailSecurityTopTldMaliciousGetParamsTLSVersion = "TLSv1_0"
	RadarEmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_1 RadarEmailSecurityTopTldMaliciousGetParamsTLSVersion = "TLSv1_1"
	RadarEmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_2 RadarEmailSecurityTopTldMaliciousGetParamsTLSVersion = "TLSv1_2"
	RadarEmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_3 RadarEmailSecurityTopTldMaliciousGetParamsTLSVersion = "TLSv1_3"
)

type RadarEmailSecurityTopTldMaliciousGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopTldMaliciousGetResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarEmailSecurityTopTldMaliciousGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopTldMaliciousGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopTldMaliciousGetResponseEnvelope]
type radarEmailSecurityTopTldMaliciousGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldMaliciousGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
