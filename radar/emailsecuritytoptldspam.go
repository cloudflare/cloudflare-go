// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

// EmailSecurityTopTldSpamService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailSecurityTopTldSpamService] method instead.
type EmailSecurityTopTldSpamService struct {
	Options []option.RequestOption
}

// NewEmailSecurityTopTldSpamService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailSecurityTopTldSpamService(opts ...option.RequestOption) (r *EmailSecurityTopTldSpamService) {
	r = &EmailSecurityTopTldSpamService{}
	r.Options = opts
	return
}

// Get the top TLDs by emails classified as Spam or not.
func (r *EmailSecurityTopTldSpamService) Get(ctx context.Context, spam EmailSecurityTopTldSpamGetParamsSpam, query EmailSecurityTopTldSpamGetParams, opts ...option.RequestOption) (res *EmailSecurityTopTldSpamGetResponse, err error) {
	var env EmailSecurityTopTldSpamGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/security/top/tlds/spam/%v", spam)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailSecurityTopTldSpamGetResponse struct {
	Meta EmailSecurityTopTldSpamGetResponseMeta `json:"meta,required"`
	Top0 []Browser                              `json:"top_0,required"`
	JSON emailSecurityTopTldSpamGetResponseJSON `json:"-"`
}

// emailSecurityTopTldSpamGetResponseJSON contains the JSON metadata for the struct
// [EmailSecurityTopTldSpamGetResponse]
type emailSecurityTopTldSpamGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpamGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpamGetResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldSpamGetResponseMeta struct {
	DateRange      []EmailSecurityTopTldSpamGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                               `json:"lastUpdated,required"`
	ConfidenceInfo EmailSecurityTopTldSpamGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           emailSecurityTopTldSpamGetResponseMetaJSON           `json:"-"`
}

// emailSecurityTopTldSpamGetResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecurityTopTldSpamGetResponseMeta]
type emailSecurityTopTldSpamGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpamGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpamGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldSpamGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                           `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTopTldSpamGetResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTopTldSpamGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [EmailSecurityTopTldSpamGetResponseMetaDateRange]
type emailSecurityTopTldSpamGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpamGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpamGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldSpamGetResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                            `json:"level"`
	JSON        emailSecurityTopTldSpamGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// emailSecurityTopTldSpamGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [EmailSecurityTopTldSpamGetResponseMetaConfidenceInfo]
type emailSecurityTopTldSpamGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpamGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpamGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                             `json:"dataSource,required"`
	Description     string                                                             `json:"description,required"`
	EventType       string                                                             `json:"eventType,required"`
	IsInstantaneous bool                                                               `json:"isInstantaneous,required"`
	EndTime         time.Time                                                          `json:"endTime" format:"date-time"`
	LinkedURL       string                                                             `json:"linkedUrl"`
	StartTime       time.Time                                                          `json:"startTime" format:"date-time"`
	JSON            emailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotation]
type emailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpamGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldSpamGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecurityTopTldSpamGetParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecurityTopTldSpamGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecurityTopTldSpamGetParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecurityTopTldSpamGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecurityTopTldSpamGetParamsSPF] `query:"spf"`
	// Filter for TLDs by category.
	TldCategory param.Field[EmailSecurityTopTldSpamGetParamsTldCategory] `query:"tldCategory"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecurityTopTldSpamGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTopTldSpamGetParams]'s query parameters as
// `url.Values`.
func (r EmailSecurityTopTldSpamGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Spam.
type EmailSecurityTopTldSpamGetParamsSpam string

const (
	EmailSecurityTopTldSpamGetParamsSpamSpam    EmailSecurityTopTldSpamGetParamsSpam = "SPAM"
	EmailSecurityTopTldSpamGetParamsSpamNotSpam EmailSecurityTopTldSpamGetParamsSpam = "NOT_SPAM"
)

func (r EmailSecurityTopTldSpamGetParamsSpam) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpamGetParamsSpamSpam, EmailSecurityTopTldSpamGetParamsSpamNotSpam:
		return true
	}
	return false
}

type EmailSecurityTopTldSpamGetParamsARC string

const (
	EmailSecurityTopTldSpamGetParamsARCPass EmailSecurityTopTldSpamGetParamsARC = "PASS"
	EmailSecurityTopTldSpamGetParamsARCNone EmailSecurityTopTldSpamGetParamsARC = "NONE"
	EmailSecurityTopTldSpamGetParamsARCFail EmailSecurityTopTldSpamGetParamsARC = "FAIL"
)

func (r EmailSecurityTopTldSpamGetParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpamGetParamsARCPass, EmailSecurityTopTldSpamGetParamsARCNone, EmailSecurityTopTldSpamGetParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTopTldSpamGetParamsDKIM string

const (
	EmailSecurityTopTldSpamGetParamsDKIMPass EmailSecurityTopTldSpamGetParamsDKIM = "PASS"
	EmailSecurityTopTldSpamGetParamsDKIMNone EmailSecurityTopTldSpamGetParamsDKIM = "NONE"
	EmailSecurityTopTldSpamGetParamsDKIMFail EmailSecurityTopTldSpamGetParamsDKIM = "FAIL"
)

func (r EmailSecurityTopTldSpamGetParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpamGetParamsDKIMPass, EmailSecurityTopTldSpamGetParamsDKIMNone, EmailSecurityTopTldSpamGetParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTopTldSpamGetParamsDMARC string

const (
	EmailSecurityTopTldSpamGetParamsDMARCPass EmailSecurityTopTldSpamGetParamsDMARC = "PASS"
	EmailSecurityTopTldSpamGetParamsDMARCNone EmailSecurityTopTldSpamGetParamsDMARC = "NONE"
	EmailSecurityTopTldSpamGetParamsDMARCFail EmailSecurityTopTldSpamGetParamsDMARC = "FAIL"
)

func (r EmailSecurityTopTldSpamGetParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpamGetParamsDMARCPass, EmailSecurityTopTldSpamGetParamsDMARCNone, EmailSecurityTopTldSpamGetParamsDMARCFail:
		return true
	}
	return false
}

// Format results are returned in.
type EmailSecurityTopTldSpamGetParamsFormat string

const (
	EmailSecurityTopTldSpamGetParamsFormatJson EmailSecurityTopTldSpamGetParamsFormat = "JSON"
	EmailSecurityTopTldSpamGetParamsFormatCsv  EmailSecurityTopTldSpamGetParamsFormat = "CSV"
)

func (r EmailSecurityTopTldSpamGetParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpamGetParamsFormatJson, EmailSecurityTopTldSpamGetParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTopTldSpamGetParamsSPF string

const (
	EmailSecurityTopTldSpamGetParamsSPFPass EmailSecurityTopTldSpamGetParamsSPF = "PASS"
	EmailSecurityTopTldSpamGetParamsSPFNone EmailSecurityTopTldSpamGetParamsSPF = "NONE"
	EmailSecurityTopTldSpamGetParamsSPFFail EmailSecurityTopTldSpamGetParamsSPF = "FAIL"
)

func (r EmailSecurityTopTldSpamGetParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpamGetParamsSPFPass, EmailSecurityTopTldSpamGetParamsSPFNone, EmailSecurityTopTldSpamGetParamsSPFFail:
		return true
	}
	return false
}

// Filter for TLDs by category.
type EmailSecurityTopTldSpamGetParamsTldCategory string

const (
	EmailSecurityTopTldSpamGetParamsTldCategoryClassic EmailSecurityTopTldSpamGetParamsTldCategory = "CLASSIC"
	EmailSecurityTopTldSpamGetParamsTldCategoryCountry EmailSecurityTopTldSpamGetParamsTldCategory = "COUNTRY"
)

func (r EmailSecurityTopTldSpamGetParamsTldCategory) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpamGetParamsTldCategoryClassic, EmailSecurityTopTldSpamGetParamsTldCategoryCountry:
		return true
	}
	return false
}

type EmailSecurityTopTldSpamGetParamsTLSVersion string

const (
	EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_0 EmailSecurityTopTldSpamGetParamsTLSVersion = "TLSv1_0"
	EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_1 EmailSecurityTopTldSpamGetParamsTLSVersion = "TLSv1_1"
	EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_2 EmailSecurityTopTldSpamGetParamsTLSVersion = "TLSv1_2"
	EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_3 EmailSecurityTopTldSpamGetParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTopTldSpamGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_0, EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_1, EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_2, EmailSecurityTopTldSpamGetParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTopTldSpamGetResponseEnvelope struct {
	Result  EmailSecurityTopTldSpamGetResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    emailSecurityTopTldSpamGetResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTopTldSpamGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailSecurityTopTldSpamGetResponseEnvelope]
type emailSecurityTopTldSpamGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpamGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpamGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
