// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// EmailSecurityTopTldService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailSecurityTopTldService] method instead.
type EmailSecurityTopTldService struct {
	Options   []option.RequestOption
	Malicious *EmailSecurityTopTldMaliciousService
	Spam      *EmailSecurityTopTldSpamService
	Spoof     *EmailSecurityTopTldSpoofService
}

// NewEmailSecurityTopTldService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailSecurityTopTldService(opts ...option.RequestOption) (r *EmailSecurityTopTldService) {
	r = &EmailSecurityTopTldService{}
	r.Options = opts
	r.Malicious = NewEmailSecurityTopTldMaliciousService(opts...)
	r.Spam = NewEmailSecurityTopTldSpamService(opts...)
	r.Spoof = NewEmailSecurityTopTldSpoofService(opts...)
	return
}

// Get the top TLDs by email messages. Values are a percentage out of the total
// emails.
func (r *EmailSecurityTopTldService) Get(ctx context.Context, query EmailSecurityTopTldGetParams, opts ...option.RequestOption) (res *EmailSecurityTopTldGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecurityTopTldGetResponseEnvelope
	path := "radar/email/security/top/tlds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailSecurityTopTldGetResponse struct {
	Meta EmailSecurityTopTldGetResponseMeta `json:"meta,required"`
	Top0 []Browser                          `json:"top_0,required"`
	JSON emailSecurityTopTldGetResponseJSON `json:"-"`
}

// emailSecurityTopTldGetResponseJSON contains the JSON metadata for the struct
// [EmailSecurityTopTldGetResponse]
type emailSecurityTopTldGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldGetResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldGetResponseMeta struct {
	DateRange      []EmailSecurityTopTldGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                           `json:"lastUpdated,required"`
	ConfidenceInfo EmailSecurityTopTldGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           emailSecurityTopTldGetResponseMetaJSON           `json:"-"`
}

// emailSecurityTopTldGetResponseMetaJSON contains the JSON metadata for the struct
// [EmailSecurityTopTldGetResponseMeta]
type emailSecurityTopTldGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTopTldGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTopTldGetResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTopTldGetResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [EmailSecurityTopTldGetResponseMetaDateRange]
type emailSecurityTopTldGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldGetResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTopTldGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                        `json:"level"`
	JSON        emailSecurityTopTldGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// emailSecurityTopTldGetResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [EmailSecurityTopTldGetResponseMetaConfidenceInfo]
type emailSecurityTopTldGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                         `json:"dataSource,required"`
	Description     string                                                         `json:"description,required"`
	EventType       string                                                         `json:"eventType,required"`
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                      `json:"startTime" format:"date-time"`
	JSON            emailSecurityTopTldGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTopTldGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [EmailSecurityTopTldGetResponseMetaConfidenceInfoAnnotation]
type emailSecurityTopTldGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTopTldGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecurityTopTldGetParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecurityTopTldGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecurityTopTldGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecurityTopTldGetParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecurityTopTldGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecurityTopTldGetParamsSPF] `query:"spf"`
	// Filter for TLDs by category.
	TldCategory param.Field[EmailSecurityTopTldGetParamsTldCategory] `query:"tldCategory"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecurityTopTldGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTopTldGetParams]'s query parameters as
// `url.Values`.
func (r EmailSecurityTopTldGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailSecurityTopTldGetParamsARC string

const (
	EmailSecurityTopTldGetParamsARCPass EmailSecurityTopTldGetParamsARC = "PASS"
	EmailSecurityTopTldGetParamsARCNone EmailSecurityTopTldGetParamsARC = "NONE"
	EmailSecurityTopTldGetParamsARCFail EmailSecurityTopTldGetParamsARC = "FAIL"
)

func (r EmailSecurityTopTldGetParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldGetParamsARCPass, EmailSecurityTopTldGetParamsARCNone, EmailSecurityTopTldGetParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTopTldGetParamsDateRange string

const (
	EmailSecurityTopTldGetParamsDateRange1d         EmailSecurityTopTldGetParamsDateRange = "1d"
	EmailSecurityTopTldGetParamsDateRange2d         EmailSecurityTopTldGetParamsDateRange = "2d"
	EmailSecurityTopTldGetParamsDateRange7d         EmailSecurityTopTldGetParamsDateRange = "7d"
	EmailSecurityTopTldGetParamsDateRange14d        EmailSecurityTopTldGetParamsDateRange = "14d"
	EmailSecurityTopTldGetParamsDateRange28d        EmailSecurityTopTldGetParamsDateRange = "28d"
	EmailSecurityTopTldGetParamsDateRange12w        EmailSecurityTopTldGetParamsDateRange = "12w"
	EmailSecurityTopTldGetParamsDateRange24w        EmailSecurityTopTldGetParamsDateRange = "24w"
	EmailSecurityTopTldGetParamsDateRange52w        EmailSecurityTopTldGetParamsDateRange = "52w"
	EmailSecurityTopTldGetParamsDateRange1dControl  EmailSecurityTopTldGetParamsDateRange = "1dControl"
	EmailSecurityTopTldGetParamsDateRange2dControl  EmailSecurityTopTldGetParamsDateRange = "2dControl"
	EmailSecurityTopTldGetParamsDateRange7dControl  EmailSecurityTopTldGetParamsDateRange = "7dControl"
	EmailSecurityTopTldGetParamsDateRange14dControl EmailSecurityTopTldGetParamsDateRange = "14dControl"
	EmailSecurityTopTldGetParamsDateRange28dControl EmailSecurityTopTldGetParamsDateRange = "28dControl"
	EmailSecurityTopTldGetParamsDateRange12wControl EmailSecurityTopTldGetParamsDateRange = "12wControl"
	EmailSecurityTopTldGetParamsDateRange24wControl EmailSecurityTopTldGetParamsDateRange = "24wControl"
)

func (r EmailSecurityTopTldGetParamsDateRange) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldGetParamsDateRange1d, EmailSecurityTopTldGetParamsDateRange2d, EmailSecurityTopTldGetParamsDateRange7d, EmailSecurityTopTldGetParamsDateRange14d, EmailSecurityTopTldGetParamsDateRange28d, EmailSecurityTopTldGetParamsDateRange12w, EmailSecurityTopTldGetParamsDateRange24w, EmailSecurityTopTldGetParamsDateRange52w, EmailSecurityTopTldGetParamsDateRange1dControl, EmailSecurityTopTldGetParamsDateRange2dControl, EmailSecurityTopTldGetParamsDateRange7dControl, EmailSecurityTopTldGetParamsDateRange14dControl, EmailSecurityTopTldGetParamsDateRange28dControl, EmailSecurityTopTldGetParamsDateRange12wControl, EmailSecurityTopTldGetParamsDateRange24wControl:
		return true
	}
	return false
}

type EmailSecurityTopTldGetParamsDKIM string

const (
	EmailSecurityTopTldGetParamsDKIMPass EmailSecurityTopTldGetParamsDKIM = "PASS"
	EmailSecurityTopTldGetParamsDKIMNone EmailSecurityTopTldGetParamsDKIM = "NONE"
	EmailSecurityTopTldGetParamsDKIMFail EmailSecurityTopTldGetParamsDKIM = "FAIL"
)

func (r EmailSecurityTopTldGetParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldGetParamsDKIMPass, EmailSecurityTopTldGetParamsDKIMNone, EmailSecurityTopTldGetParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTopTldGetParamsDMARC string

const (
	EmailSecurityTopTldGetParamsDMARCPass EmailSecurityTopTldGetParamsDMARC = "PASS"
	EmailSecurityTopTldGetParamsDMARCNone EmailSecurityTopTldGetParamsDMARC = "NONE"
	EmailSecurityTopTldGetParamsDMARCFail EmailSecurityTopTldGetParamsDMARC = "FAIL"
)

func (r EmailSecurityTopTldGetParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldGetParamsDMARCPass, EmailSecurityTopTldGetParamsDMARCNone, EmailSecurityTopTldGetParamsDMARCFail:
		return true
	}
	return false
}

// Format results are returned in.
type EmailSecurityTopTldGetParamsFormat string

const (
	EmailSecurityTopTldGetParamsFormatJson EmailSecurityTopTldGetParamsFormat = "JSON"
	EmailSecurityTopTldGetParamsFormatCsv  EmailSecurityTopTldGetParamsFormat = "CSV"
)

func (r EmailSecurityTopTldGetParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldGetParamsFormatJson, EmailSecurityTopTldGetParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTopTldGetParamsSPF string

const (
	EmailSecurityTopTldGetParamsSPFPass EmailSecurityTopTldGetParamsSPF = "PASS"
	EmailSecurityTopTldGetParamsSPFNone EmailSecurityTopTldGetParamsSPF = "NONE"
	EmailSecurityTopTldGetParamsSPFFail EmailSecurityTopTldGetParamsSPF = "FAIL"
)

func (r EmailSecurityTopTldGetParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldGetParamsSPFPass, EmailSecurityTopTldGetParamsSPFNone, EmailSecurityTopTldGetParamsSPFFail:
		return true
	}
	return false
}

// Filter for TLDs by category.
type EmailSecurityTopTldGetParamsTldCategory string

const (
	EmailSecurityTopTldGetParamsTldCategoryClassic EmailSecurityTopTldGetParamsTldCategory = "CLASSIC"
	EmailSecurityTopTldGetParamsTldCategoryCountry EmailSecurityTopTldGetParamsTldCategory = "COUNTRY"
)

func (r EmailSecurityTopTldGetParamsTldCategory) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldGetParamsTldCategoryClassic, EmailSecurityTopTldGetParamsTldCategoryCountry:
		return true
	}
	return false
}

type EmailSecurityTopTldGetParamsTLSVersion string

const (
	EmailSecurityTopTldGetParamsTLSVersionTlSv1_0 EmailSecurityTopTldGetParamsTLSVersion = "TLSv1_0"
	EmailSecurityTopTldGetParamsTLSVersionTlSv1_1 EmailSecurityTopTldGetParamsTLSVersion = "TLSv1_1"
	EmailSecurityTopTldGetParamsTLSVersionTlSv1_2 EmailSecurityTopTldGetParamsTLSVersion = "TLSv1_2"
	EmailSecurityTopTldGetParamsTLSVersionTlSv1_3 EmailSecurityTopTldGetParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTopTldGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldGetParamsTLSVersionTlSv1_0, EmailSecurityTopTldGetParamsTLSVersionTlSv1_1, EmailSecurityTopTldGetParamsTLSVersionTlSv1_2, EmailSecurityTopTldGetParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTopTldGetResponseEnvelope struct {
	Result  EmailSecurityTopTldGetResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    emailSecurityTopTldGetResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTopTldGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailSecurityTopTldGetResponseEnvelope]
type emailSecurityTopTldGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
