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

// RadarEmailSecurityTopTldSpoofService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopTldSpoofService] method instead.
type RadarEmailSecurityTopTldSpoofService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopTldSpoofService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopTldSpoofService(opts ...option.RequestOption) (r *RadarEmailSecurityTopTldSpoofService) {
	r = &RadarEmailSecurityTopTldSpoofService{}
	r.Options = opts
	return
}

// Get the TLDs by emails classified as spoof or not.
func (r *RadarEmailSecurityTopTldSpoofService) Get(ctx context.Context, spoof RadarEmailSecurityTopTldSpoofGetParamsSpoof, query RadarEmailSecurityTopTldSpoofGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopTldSpoofGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopTldSpoofGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/tlds/spoof/%v", spoof)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopTldSpoofGetResponse struct {
	Meta RadarEmailSecurityTopTldSpoofGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopTldSpoofGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopTldSpoofGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopTldSpoofGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopTldSpoofGetResponse]
type radarEmailSecurityTopTldSpoofGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldSpoofGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldSpoofGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldSpoofGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopTldSpoofGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                     `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopTldSpoofGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopTldSpoofGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopTldSpoofGetResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopTldSpoofGetResponseMeta]
type radarEmailSecurityTopTldSpoofGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldSpoofGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldSpoofGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldSpoofGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopTldSpoofGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopTldSpoofGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopTldSpoofGetResponseMetaDateRange]
type radarEmailSecurityTopTldSpoofGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldSpoofGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldSpoofGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldSpoofGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarEmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopTldSpoofGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldSpoofGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous interface{}                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldSpoofGetResponseTop0 struct {
	Name  string                                           `json:"name,required"`
	Value string                                           `json:"value,required"`
	JSON  radarEmailSecurityTopTldSpoofGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopTldSpoofGetResponseTop0JSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopTldSpoofGetResponseTop0]
type radarEmailSecurityTopTldSpoofGetResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldSpoofGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldSpoofGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldSpoofGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTopTldSpoofGetParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopTldSpoofGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopTldSpoofGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTopTldSpoofGetParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopTldSpoofGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopTldSpoofGetParamsSPF] `query:"spf"`
	// Filter for TLDs by category.
	TldCategory param.Field[RadarEmailSecurityTopTldSpoofGetParamsTldCategory] `query:"tldCategory"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecurityTopTldSpoofGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecurityTopTldSpoofGetParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecurityTopTldSpoofGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Spoof.
type RadarEmailSecurityTopTldSpoofGetParamsSpoof string

const (
	RadarEmailSecurityTopTldSpoofGetParamsSpoofSpoof    RadarEmailSecurityTopTldSpoofGetParamsSpoof = "SPOOF"
	RadarEmailSecurityTopTldSpoofGetParamsSpoofNotSpoof RadarEmailSecurityTopTldSpoofGetParamsSpoof = "NOT_SPOOF"
)

type RadarEmailSecurityTopTldSpoofGetParamsARC string

const (
	RadarEmailSecurityTopTldSpoofGetParamsARCPass RadarEmailSecurityTopTldSpoofGetParamsARC = "PASS"
	RadarEmailSecurityTopTldSpoofGetParamsARCNone RadarEmailSecurityTopTldSpoofGetParamsARC = "NONE"
	RadarEmailSecurityTopTldSpoofGetParamsARCFail RadarEmailSecurityTopTldSpoofGetParamsARC = "FAIL"
)

type RadarEmailSecurityTopTldSpoofGetParamsDateRange string

const (
	RadarEmailSecurityTopTldSpoofGetParamsDateRange1d         RadarEmailSecurityTopTldSpoofGetParamsDateRange = "1d"
	RadarEmailSecurityTopTldSpoofGetParamsDateRange2d         RadarEmailSecurityTopTldSpoofGetParamsDateRange = "2d"
	RadarEmailSecurityTopTldSpoofGetParamsDateRange7d         RadarEmailSecurityTopTldSpoofGetParamsDateRange = "7d"
	RadarEmailSecurityTopTldSpoofGetParamsDateRange14d        RadarEmailSecurityTopTldSpoofGetParamsDateRange = "14d"
	RadarEmailSecurityTopTldSpoofGetParamsDateRange28d        RadarEmailSecurityTopTldSpoofGetParamsDateRange = "28d"
	RadarEmailSecurityTopTldSpoofGetParamsDateRange12w        RadarEmailSecurityTopTldSpoofGetParamsDateRange = "12w"
	RadarEmailSecurityTopTldSpoofGetParamsDateRange24w        RadarEmailSecurityTopTldSpoofGetParamsDateRange = "24w"
	RadarEmailSecurityTopTldSpoofGetParamsDateRange52w        RadarEmailSecurityTopTldSpoofGetParamsDateRange = "52w"
	RadarEmailSecurityTopTldSpoofGetParamsDateRange1dControl  RadarEmailSecurityTopTldSpoofGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopTldSpoofGetParamsDateRange2dControl  RadarEmailSecurityTopTldSpoofGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopTldSpoofGetParamsDateRange7dControl  RadarEmailSecurityTopTldSpoofGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopTldSpoofGetParamsDateRange14dControl RadarEmailSecurityTopTldSpoofGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopTldSpoofGetParamsDateRange28dControl RadarEmailSecurityTopTldSpoofGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopTldSpoofGetParamsDateRange12wControl RadarEmailSecurityTopTldSpoofGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopTldSpoofGetParamsDateRange24wControl RadarEmailSecurityTopTldSpoofGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopTldSpoofGetParamsDKIM string

const (
	RadarEmailSecurityTopTldSpoofGetParamsDKIMPass RadarEmailSecurityTopTldSpoofGetParamsDKIM = "PASS"
	RadarEmailSecurityTopTldSpoofGetParamsDKIMNone RadarEmailSecurityTopTldSpoofGetParamsDKIM = "NONE"
	RadarEmailSecurityTopTldSpoofGetParamsDKIMFail RadarEmailSecurityTopTldSpoofGetParamsDKIM = "FAIL"
)

type RadarEmailSecurityTopTldSpoofGetParamsDMARC string

const (
	RadarEmailSecurityTopTldSpoofGetParamsDMARCPass RadarEmailSecurityTopTldSpoofGetParamsDMARC = "PASS"
	RadarEmailSecurityTopTldSpoofGetParamsDMARCNone RadarEmailSecurityTopTldSpoofGetParamsDMARC = "NONE"
	RadarEmailSecurityTopTldSpoofGetParamsDMARCFail RadarEmailSecurityTopTldSpoofGetParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopTldSpoofGetParamsFormat string

const (
	RadarEmailSecurityTopTldSpoofGetParamsFormatJson RadarEmailSecurityTopTldSpoofGetParamsFormat = "JSON"
	RadarEmailSecurityTopTldSpoofGetParamsFormatCsv  RadarEmailSecurityTopTldSpoofGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopTldSpoofGetParamsSPF string

const (
	RadarEmailSecurityTopTldSpoofGetParamsSPFPass RadarEmailSecurityTopTldSpoofGetParamsSPF = "PASS"
	RadarEmailSecurityTopTldSpoofGetParamsSPFNone RadarEmailSecurityTopTldSpoofGetParamsSPF = "NONE"
	RadarEmailSecurityTopTldSpoofGetParamsSPFFail RadarEmailSecurityTopTldSpoofGetParamsSPF = "FAIL"
)

// Filter for TLDs by category.
type RadarEmailSecurityTopTldSpoofGetParamsTldCategory string

const (
	RadarEmailSecurityTopTldSpoofGetParamsTldCategoryClassic RadarEmailSecurityTopTldSpoofGetParamsTldCategory = "CLASSIC"
	RadarEmailSecurityTopTldSpoofGetParamsTldCategoryCountry RadarEmailSecurityTopTldSpoofGetParamsTldCategory = "COUNTRY"
)

type RadarEmailSecurityTopTldSpoofGetParamsTLSVersion string

const (
	RadarEmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_0 RadarEmailSecurityTopTldSpoofGetParamsTLSVersion = "TLSv1_0"
	RadarEmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_1 RadarEmailSecurityTopTldSpoofGetParamsTLSVersion = "TLSv1_1"
	RadarEmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_2 RadarEmailSecurityTopTldSpoofGetParamsTLSVersion = "TLSv1_2"
	RadarEmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_3 RadarEmailSecurityTopTldSpoofGetParamsTLSVersion = "TLSv1_3"
)

type RadarEmailSecurityTopTldSpoofGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopTldSpoofGetResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarEmailSecurityTopTldSpoofGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopTldSpoofGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopTldSpoofGetResponseEnvelope]
type radarEmailSecurityTopTldSpoofGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldSpoofGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldSpoofGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
