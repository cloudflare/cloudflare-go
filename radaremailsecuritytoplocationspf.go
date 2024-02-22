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

// RadarEmailSecurityTopLocationSPFService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopLocationSPFService] method instead.
type RadarEmailSecurityTopLocationSPFService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopLocationSPFService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopLocationSPFService(opts ...option.RequestOption) (r *RadarEmailSecurityTopLocationSPFService) {
	r = &RadarEmailSecurityTopLocationSPFService{}
	r.Options = opts
	return
}

// Get the top locations by email SPF validation.
func (r *RadarEmailSecurityTopLocationSPFService) Get(ctx context.Context, spf RadarEmailSecurityTopLocationSPFGetParamsSPF, query RadarEmailSecurityTopLocationSPFGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopLocationSPFGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopLocationSPFGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/locations/spf/%v", spf)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopLocationSPFGetResponse struct {
	Meta RadarEmailSecurityTopLocationSPFGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopLocationSPFGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopLocationSPFGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopLocationSPFGetResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopLocationSPFGetResponse]
type radarEmailSecurityTopLocationSPFGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationSPFGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationSPFGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopLocationSPFGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                        `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopLocationSPFGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopLocationSPFGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopLocationSPFGetResponseMetaJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationSPFGetResponseMeta]
type radarEmailSecurityTopLocationSPFGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationSPFGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationSPFGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopLocationSPFGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopLocationSPFGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopLocationSPFGetResponseMetaDateRange]
type radarEmailSecurityTopLocationSPFGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationSPFGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationSPFGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopLocationSPFGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                     `json:"level"`
	JSON        radarEmailSecurityTopLocationSPFGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopLocationSPFGetResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopLocationSPFGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopLocationSPFGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationSPFGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationSPFGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                      `json:"dataSource,required"`
	Description     string                                                                      `json:"description,required"`
	EventType       string                                                                      `json:"eventType,required"`
	IsInstantaneous interface{}                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                                   `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopLocationSPFGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopLocationSPFGetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationSPFGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopLocationSPFGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopLocationSPFGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationSPFGetResponseTop0 struct {
	ClientCountryAlpha2 string                                              `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                              `json:"clientCountryName,required"`
	Value               string                                              `json:"value,required"`
	JSON                radarEmailSecurityTopLocationSPFGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopLocationSPFGetResponseTop0JSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationSPFGetResponseTop0]
type radarEmailSecurityTopLocationSPFGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationSPFGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationSPFGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTopLocationSPFGetParamsARC] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopLocationSPFGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopLocationSPFGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTopLocationSPFGetParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopLocationSPFGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailSecurityTopLocationSPFGetParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTopLocationSPFGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// SPF.
type RadarEmailSecurityTopLocationSPFGetParamsSPF string

const (
	RadarEmailSecurityTopLocationSPFGetParamsSPFPass RadarEmailSecurityTopLocationSPFGetParamsSPF = "PASS"
	RadarEmailSecurityTopLocationSPFGetParamsSPFNone RadarEmailSecurityTopLocationSPFGetParamsSPF = "NONE"
	RadarEmailSecurityTopLocationSPFGetParamsSPFFail RadarEmailSecurityTopLocationSPFGetParamsSPF = "FAIL"
)

type RadarEmailSecurityTopLocationSPFGetParamsARC string

const (
	RadarEmailSecurityTopLocationSPFGetParamsARCPass RadarEmailSecurityTopLocationSPFGetParamsARC = "PASS"
	RadarEmailSecurityTopLocationSPFGetParamsARCNone RadarEmailSecurityTopLocationSPFGetParamsARC = "NONE"
	RadarEmailSecurityTopLocationSPFGetParamsARCFail RadarEmailSecurityTopLocationSPFGetParamsARC = "FAIL"
)

type RadarEmailSecurityTopLocationSPFGetParamsDateRange string

const (
	RadarEmailSecurityTopLocationSPFGetParamsDateRange1d         RadarEmailSecurityTopLocationSPFGetParamsDateRange = "1d"
	RadarEmailSecurityTopLocationSPFGetParamsDateRange2d         RadarEmailSecurityTopLocationSPFGetParamsDateRange = "2d"
	RadarEmailSecurityTopLocationSPFGetParamsDateRange7d         RadarEmailSecurityTopLocationSPFGetParamsDateRange = "7d"
	RadarEmailSecurityTopLocationSPFGetParamsDateRange14d        RadarEmailSecurityTopLocationSPFGetParamsDateRange = "14d"
	RadarEmailSecurityTopLocationSPFGetParamsDateRange28d        RadarEmailSecurityTopLocationSPFGetParamsDateRange = "28d"
	RadarEmailSecurityTopLocationSPFGetParamsDateRange12w        RadarEmailSecurityTopLocationSPFGetParamsDateRange = "12w"
	RadarEmailSecurityTopLocationSPFGetParamsDateRange24w        RadarEmailSecurityTopLocationSPFGetParamsDateRange = "24w"
	RadarEmailSecurityTopLocationSPFGetParamsDateRange52w        RadarEmailSecurityTopLocationSPFGetParamsDateRange = "52w"
	RadarEmailSecurityTopLocationSPFGetParamsDateRange1dControl  RadarEmailSecurityTopLocationSPFGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopLocationSPFGetParamsDateRange2dControl  RadarEmailSecurityTopLocationSPFGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopLocationSPFGetParamsDateRange7dControl  RadarEmailSecurityTopLocationSPFGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopLocationSPFGetParamsDateRange14dControl RadarEmailSecurityTopLocationSPFGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopLocationSPFGetParamsDateRange28dControl RadarEmailSecurityTopLocationSPFGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopLocationSPFGetParamsDateRange12wControl RadarEmailSecurityTopLocationSPFGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopLocationSPFGetParamsDateRange24wControl RadarEmailSecurityTopLocationSPFGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopLocationSPFGetParamsDKIM string

const (
	RadarEmailSecurityTopLocationSPFGetParamsDKIMPass RadarEmailSecurityTopLocationSPFGetParamsDKIM = "PASS"
	RadarEmailSecurityTopLocationSPFGetParamsDKIMNone RadarEmailSecurityTopLocationSPFGetParamsDKIM = "NONE"
	RadarEmailSecurityTopLocationSPFGetParamsDKIMFail RadarEmailSecurityTopLocationSPFGetParamsDKIM = "FAIL"
)

type RadarEmailSecurityTopLocationSPFGetParamsDMARC string

const (
	RadarEmailSecurityTopLocationSPFGetParamsDMARCPass RadarEmailSecurityTopLocationSPFGetParamsDMARC = "PASS"
	RadarEmailSecurityTopLocationSPFGetParamsDMARCNone RadarEmailSecurityTopLocationSPFGetParamsDMARC = "NONE"
	RadarEmailSecurityTopLocationSPFGetParamsDMARCFail RadarEmailSecurityTopLocationSPFGetParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopLocationSPFGetParamsFormat string

const (
	RadarEmailSecurityTopLocationSPFGetParamsFormatJson RadarEmailSecurityTopLocationSPFGetParamsFormat = "JSON"
	RadarEmailSecurityTopLocationSPFGetParamsFormatCsv  RadarEmailSecurityTopLocationSPFGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopLocationSPFGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopLocationSPFGetResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarEmailSecurityTopLocationSPFGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopLocationSPFGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopLocationSPFGetResponseEnvelope]
type radarEmailSecurityTopLocationSPFGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationSPFGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
