// File generated from our OpenAPI spec by Stainless.

package radar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// EmailSecurityTopTldMaliciousService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewEmailSecurityTopTldMaliciousService] method instead.
type EmailSecurityTopTldMaliciousService struct {
	Options []option.RequestOption
}

// NewEmailSecurityTopTldMaliciousService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmailSecurityTopTldMaliciousService(opts ...option.RequestOption) (r *EmailSecurityTopTldMaliciousService) {
	r = &EmailSecurityTopTldMaliciousService{}
	r.Options = opts
	return
}

// Get the TLDs by emails classified as malicious or not.
func (r *EmailSecurityTopTldMaliciousService) Get(ctx context.Context, malicious EmailSecurityTopTldMaliciousGetParamsMalicious, query EmailSecurityTopTldMaliciousGetParams, opts ...option.RequestOption) (res *EmailSecurityTopTldMaliciousGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecurityTopTldMaliciousGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/tlds/malicious/%v", malicious)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailSecurityTopTldMaliciousGetResponse struct {
	Meta EmailSecurityTopTldMaliciousGetResponseMeta   `json:"meta,required"`
	Top0 []EmailSecurityTopTldMaliciousGetResponseTop0 `json:"top_0,required"`
	JSON emailSecurityTopTldMaliciousGetResponseJSON   `json:"-"`
}

// emailSecurityTopTldMaliciousGetResponseJSON contains the JSON metadata for the
// struct [EmailSecurityTopTldMaliciousGetResponse]
type emailSecurityTopTldMaliciousGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldMaliciousGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldMaliciousGetResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldMaliciousGetResponseMeta struct {
	DateRange      []EmailSecurityTopTldMaliciousGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                    `json:"lastUpdated,required"`
	ConfidenceInfo EmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           emailSecurityTopTldMaliciousGetResponseMetaJSON           `json:"-"`
}

// emailSecurityTopTldMaliciousGetResponseMetaJSON contains the JSON metadata for
// the struct [EmailSecurityTopTldMaliciousGetResponseMeta]
type emailSecurityTopTldMaliciousGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTopTldMaliciousGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldMaliciousGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldMaliciousGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTopTldMaliciousGetResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTopTldMaliciousGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailSecurityTopTldMaliciousGetResponseMetaDateRange]
type emailSecurityTopTldMaliciousGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldMaliciousGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldMaliciousGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                 `json:"level"`
	JSON        emailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// emailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [EmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfo]
type emailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                  `json:"dataSource,required"`
	Description     string                                                                  `json:"description,required"`
	EventType       string                                                                  `json:"eventType,required"`
	IsInstantaneous interface{}                                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                                               `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                  `json:"linkedUrl"`
	StartTime       time.Time                                                               `json:"startTime" format:"date-time"`
	JSON            emailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotation]
type emailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldMaliciousGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldMaliciousGetResponseTop0 struct {
	Name  string                                          `json:"name,required"`
	Value string                                          `json:"value,required"`
	JSON  emailSecurityTopTldMaliciousGetResponseTop0JSON `json:"-"`
}

// emailSecurityTopTldMaliciousGetResponseTop0JSON contains the JSON metadata for
// the struct [EmailSecurityTopTldMaliciousGetResponseTop0]
type emailSecurityTopTldMaliciousGetResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldMaliciousGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldMaliciousGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldMaliciousGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecurityTopTldMaliciousGetParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecurityTopTldMaliciousGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecurityTopTldMaliciousGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecurityTopTldMaliciousGetParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecurityTopTldMaliciousGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecurityTopTldMaliciousGetParamsSPF] `query:"spf"`
	// Filter for TLDs by category.
	TldCategory param.Field[EmailSecurityTopTldMaliciousGetParamsTldCategory] `query:"tldCategory"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecurityTopTldMaliciousGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTopTldMaliciousGetParams]'s query parameters
// as `url.Values`.
func (r EmailSecurityTopTldMaliciousGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Malicious.
type EmailSecurityTopTldMaliciousGetParamsMalicious string

const (
	EmailSecurityTopTldMaliciousGetParamsMaliciousMalicious    EmailSecurityTopTldMaliciousGetParamsMalicious = "MALICIOUS"
	EmailSecurityTopTldMaliciousGetParamsMaliciousNotMalicious EmailSecurityTopTldMaliciousGetParamsMalicious = "NOT_MALICIOUS"
)

type EmailSecurityTopTldMaliciousGetParamsARC string

const (
	EmailSecurityTopTldMaliciousGetParamsARCPass EmailSecurityTopTldMaliciousGetParamsARC = "PASS"
	EmailSecurityTopTldMaliciousGetParamsARCNone EmailSecurityTopTldMaliciousGetParamsARC = "NONE"
	EmailSecurityTopTldMaliciousGetParamsARCFail EmailSecurityTopTldMaliciousGetParamsARC = "FAIL"
)

type EmailSecurityTopTldMaliciousGetParamsDateRange string

const (
	EmailSecurityTopTldMaliciousGetParamsDateRange1d         EmailSecurityTopTldMaliciousGetParamsDateRange = "1d"
	EmailSecurityTopTldMaliciousGetParamsDateRange2d         EmailSecurityTopTldMaliciousGetParamsDateRange = "2d"
	EmailSecurityTopTldMaliciousGetParamsDateRange7d         EmailSecurityTopTldMaliciousGetParamsDateRange = "7d"
	EmailSecurityTopTldMaliciousGetParamsDateRange14d        EmailSecurityTopTldMaliciousGetParamsDateRange = "14d"
	EmailSecurityTopTldMaliciousGetParamsDateRange28d        EmailSecurityTopTldMaliciousGetParamsDateRange = "28d"
	EmailSecurityTopTldMaliciousGetParamsDateRange12w        EmailSecurityTopTldMaliciousGetParamsDateRange = "12w"
	EmailSecurityTopTldMaliciousGetParamsDateRange24w        EmailSecurityTopTldMaliciousGetParamsDateRange = "24w"
	EmailSecurityTopTldMaliciousGetParamsDateRange52w        EmailSecurityTopTldMaliciousGetParamsDateRange = "52w"
	EmailSecurityTopTldMaliciousGetParamsDateRange1dControl  EmailSecurityTopTldMaliciousGetParamsDateRange = "1dControl"
	EmailSecurityTopTldMaliciousGetParamsDateRange2dControl  EmailSecurityTopTldMaliciousGetParamsDateRange = "2dControl"
	EmailSecurityTopTldMaliciousGetParamsDateRange7dControl  EmailSecurityTopTldMaliciousGetParamsDateRange = "7dControl"
	EmailSecurityTopTldMaliciousGetParamsDateRange14dControl EmailSecurityTopTldMaliciousGetParamsDateRange = "14dControl"
	EmailSecurityTopTldMaliciousGetParamsDateRange28dControl EmailSecurityTopTldMaliciousGetParamsDateRange = "28dControl"
	EmailSecurityTopTldMaliciousGetParamsDateRange12wControl EmailSecurityTopTldMaliciousGetParamsDateRange = "12wControl"
	EmailSecurityTopTldMaliciousGetParamsDateRange24wControl EmailSecurityTopTldMaliciousGetParamsDateRange = "24wControl"
)

type EmailSecurityTopTldMaliciousGetParamsDKIM string

const (
	EmailSecurityTopTldMaliciousGetParamsDKIMPass EmailSecurityTopTldMaliciousGetParamsDKIM = "PASS"
	EmailSecurityTopTldMaliciousGetParamsDKIMNone EmailSecurityTopTldMaliciousGetParamsDKIM = "NONE"
	EmailSecurityTopTldMaliciousGetParamsDKIMFail EmailSecurityTopTldMaliciousGetParamsDKIM = "FAIL"
)

type EmailSecurityTopTldMaliciousGetParamsDMARC string

const (
	EmailSecurityTopTldMaliciousGetParamsDMARCPass EmailSecurityTopTldMaliciousGetParamsDMARC = "PASS"
	EmailSecurityTopTldMaliciousGetParamsDMARCNone EmailSecurityTopTldMaliciousGetParamsDMARC = "NONE"
	EmailSecurityTopTldMaliciousGetParamsDMARCFail EmailSecurityTopTldMaliciousGetParamsDMARC = "FAIL"
)

// Format results are returned in.
type EmailSecurityTopTldMaliciousGetParamsFormat string

const (
	EmailSecurityTopTldMaliciousGetParamsFormatJson EmailSecurityTopTldMaliciousGetParamsFormat = "JSON"
	EmailSecurityTopTldMaliciousGetParamsFormatCsv  EmailSecurityTopTldMaliciousGetParamsFormat = "CSV"
)

type EmailSecurityTopTldMaliciousGetParamsSPF string

const (
	EmailSecurityTopTldMaliciousGetParamsSPFPass EmailSecurityTopTldMaliciousGetParamsSPF = "PASS"
	EmailSecurityTopTldMaliciousGetParamsSPFNone EmailSecurityTopTldMaliciousGetParamsSPF = "NONE"
	EmailSecurityTopTldMaliciousGetParamsSPFFail EmailSecurityTopTldMaliciousGetParamsSPF = "FAIL"
)

// Filter for TLDs by category.
type EmailSecurityTopTldMaliciousGetParamsTldCategory string

const (
	EmailSecurityTopTldMaliciousGetParamsTldCategoryClassic EmailSecurityTopTldMaliciousGetParamsTldCategory = "CLASSIC"
	EmailSecurityTopTldMaliciousGetParamsTldCategoryCountry EmailSecurityTopTldMaliciousGetParamsTldCategory = "COUNTRY"
)

type EmailSecurityTopTldMaliciousGetParamsTLSVersion string

const (
	EmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_0 EmailSecurityTopTldMaliciousGetParamsTLSVersion = "TLSv1_0"
	EmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_1 EmailSecurityTopTldMaliciousGetParamsTLSVersion = "TLSv1_1"
	EmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_2 EmailSecurityTopTldMaliciousGetParamsTLSVersion = "TLSv1_2"
	EmailSecurityTopTldMaliciousGetParamsTLSVersionTlSv1_3 EmailSecurityTopTldMaliciousGetParamsTLSVersion = "TLSv1_3"
)

type EmailSecurityTopTldMaliciousGetResponseEnvelope struct {
	Result  EmailSecurityTopTldMaliciousGetResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    emailSecurityTopTldMaliciousGetResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTopTldMaliciousGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailSecurityTopTldMaliciousGetResponseEnvelope]
type emailSecurityTopTldMaliciousGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldMaliciousGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldMaliciousGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
