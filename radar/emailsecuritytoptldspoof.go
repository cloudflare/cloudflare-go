// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// EmailSecurityTopTldSpoofService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailSecurityTopTldSpoofService] method instead.
type EmailSecurityTopTldSpoofService struct {
	Options []option.RequestOption
}

// NewEmailSecurityTopTldSpoofService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmailSecurityTopTldSpoofService(opts ...option.RequestOption) (r *EmailSecurityTopTldSpoofService) {
	r = &EmailSecurityTopTldSpoofService{}
	r.Options = opts
	return
}

// Get the TLDs by emails classified as spoof or not.
func (r *EmailSecurityTopTldSpoofService) Get(ctx context.Context, spoof EmailSecurityTopTldSpoofGetParamsSpoof, query EmailSecurityTopTldSpoofGetParams, opts ...option.RequestOption) (res *EmailSecurityTopTldSpoofGetResponse, err error) {
	var env EmailSecurityTopTldSpoofGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/security/top/tlds/spoof/%v", spoof)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailSecurityTopTldSpoofGetResponse struct {
	Meta EmailSecurityTopTldSpoofGetResponseMeta   `json:"meta,required"`
	Top0 []EmailSecurityTopTldSpoofGetResponseTop0 `json:"top_0,required"`
	JSON emailSecurityTopTldSpoofGetResponseJSON   `json:"-"`
}

// emailSecurityTopTldSpoofGetResponseJSON contains the JSON metadata for the
// struct [EmailSecurityTopTldSpoofGetResponse]
type emailSecurityTopTldSpoofGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpoofGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpoofGetResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldSpoofGetResponseMeta struct {
	DateRange      []EmailSecurityTopTldSpoofGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                `json:"lastUpdated,required"`
	ConfidenceInfo EmailSecurityTopTldSpoofGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           emailSecurityTopTldSpoofGetResponseMetaJSON           `json:"-"`
}

// emailSecurityTopTldSpoofGetResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecurityTopTldSpoofGetResponseMeta]
type emailSecurityTopTldSpoofGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpoofGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpoofGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldSpoofGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      emailSecurityTopTldSpoofGetResponseMetaDateRangeJSON `json:"-"`
}

// emailSecurityTopTldSpoofGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [EmailSecurityTopTldSpoofGetResponseMetaDateRange]
type emailSecurityTopTldSpoofGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpoofGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpoofGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldSpoofGetResponseMetaConfidenceInfo struct {
	Annotations []EmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                             `json:"level"`
	JSON        emailSecurityTopTldSpoofGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// emailSecurityTopTldSpoofGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [EmailSecurityTopTldSpoofGetResponseMetaConfidenceInfo]
type emailSecurityTopTldSpoofGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpoofGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpoofGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                              `json:"dataSource,required"`
	Description     string                                                              `json:"description,required"`
	EventType       string                                                              `json:"eventType,required"`
	IsInstantaneous bool                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                           `json:"startTime" format:"date-time"`
	JSON            emailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotation]
type emailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpoofGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldSpoofGetResponseTop0 struct {
	Name  string                                      `json:"name,required"`
	Value string                                      `json:"value,required"`
	JSON  emailSecurityTopTldSpoofGetResponseTop0JSON `json:"-"`
}

// emailSecurityTopTldSpoofGetResponseTop0JSON contains the JSON metadata for the
// struct [EmailSecurityTopTldSpoofGetResponseTop0]
type emailSecurityTopTldSpoofGetResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpoofGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpoofGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type EmailSecurityTopTldSpoofGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecurityTopTldSpoofGetParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecurityTopTldSpoofGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecurityTopTldSpoofGetParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecurityTopTldSpoofGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecurityTopTldSpoofGetParamsSPF] `query:"spf"`
	// Filter for TLDs by category.
	TldCategory param.Field[EmailSecurityTopTldSpoofGetParamsTldCategory] `query:"tldCategory"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecurityTopTldSpoofGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecurityTopTldSpoofGetParams]'s query parameters as
// `url.Values`.
func (r EmailSecurityTopTldSpoofGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Spoof.
type EmailSecurityTopTldSpoofGetParamsSpoof string

const (
	EmailSecurityTopTldSpoofGetParamsSpoofSpoof    EmailSecurityTopTldSpoofGetParamsSpoof = "SPOOF"
	EmailSecurityTopTldSpoofGetParamsSpoofNotSpoof EmailSecurityTopTldSpoofGetParamsSpoof = "NOT_SPOOF"
)

func (r EmailSecurityTopTldSpoofGetParamsSpoof) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpoofGetParamsSpoofSpoof, EmailSecurityTopTldSpoofGetParamsSpoofNotSpoof:
		return true
	}
	return false
}

type EmailSecurityTopTldSpoofGetParamsARC string

const (
	EmailSecurityTopTldSpoofGetParamsARCPass EmailSecurityTopTldSpoofGetParamsARC = "PASS"
	EmailSecurityTopTldSpoofGetParamsARCNone EmailSecurityTopTldSpoofGetParamsARC = "NONE"
	EmailSecurityTopTldSpoofGetParamsARCFail EmailSecurityTopTldSpoofGetParamsARC = "FAIL"
)

func (r EmailSecurityTopTldSpoofGetParamsARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpoofGetParamsARCPass, EmailSecurityTopTldSpoofGetParamsARCNone, EmailSecurityTopTldSpoofGetParamsARCFail:
		return true
	}
	return false
}

type EmailSecurityTopTldSpoofGetParamsDKIM string

const (
	EmailSecurityTopTldSpoofGetParamsDKIMPass EmailSecurityTopTldSpoofGetParamsDKIM = "PASS"
	EmailSecurityTopTldSpoofGetParamsDKIMNone EmailSecurityTopTldSpoofGetParamsDKIM = "NONE"
	EmailSecurityTopTldSpoofGetParamsDKIMFail EmailSecurityTopTldSpoofGetParamsDKIM = "FAIL"
)

func (r EmailSecurityTopTldSpoofGetParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpoofGetParamsDKIMPass, EmailSecurityTopTldSpoofGetParamsDKIMNone, EmailSecurityTopTldSpoofGetParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecurityTopTldSpoofGetParamsDMARC string

const (
	EmailSecurityTopTldSpoofGetParamsDMARCPass EmailSecurityTopTldSpoofGetParamsDMARC = "PASS"
	EmailSecurityTopTldSpoofGetParamsDMARCNone EmailSecurityTopTldSpoofGetParamsDMARC = "NONE"
	EmailSecurityTopTldSpoofGetParamsDMARCFail EmailSecurityTopTldSpoofGetParamsDMARC = "FAIL"
)

func (r EmailSecurityTopTldSpoofGetParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpoofGetParamsDMARCPass, EmailSecurityTopTldSpoofGetParamsDMARCNone, EmailSecurityTopTldSpoofGetParamsDMARCFail:
		return true
	}
	return false
}

// Format results are returned in.
type EmailSecurityTopTldSpoofGetParamsFormat string

const (
	EmailSecurityTopTldSpoofGetParamsFormatJson EmailSecurityTopTldSpoofGetParamsFormat = "JSON"
	EmailSecurityTopTldSpoofGetParamsFormatCsv  EmailSecurityTopTldSpoofGetParamsFormat = "CSV"
)

func (r EmailSecurityTopTldSpoofGetParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpoofGetParamsFormatJson, EmailSecurityTopTldSpoofGetParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecurityTopTldSpoofGetParamsSPF string

const (
	EmailSecurityTopTldSpoofGetParamsSPFPass EmailSecurityTopTldSpoofGetParamsSPF = "PASS"
	EmailSecurityTopTldSpoofGetParamsSPFNone EmailSecurityTopTldSpoofGetParamsSPF = "NONE"
	EmailSecurityTopTldSpoofGetParamsSPFFail EmailSecurityTopTldSpoofGetParamsSPF = "FAIL"
)

func (r EmailSecurityTopTldSpoofGetParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpoofGetParamsSPFPass, EmailSecurityTopTldSpoofGetParamsSPFNone, EmailSecurityTopTldSpoofGetParamsSPFFail:
		return true
	}
	return false
}

// Filter for TLDs by category.
type EmailSecurityTopTldSpoofGetParamsTldCategory string

const (
	EmailSecurityTopTldSpoofGetParamsTldCategoryClassic EmailSecurityTopTldSpoofGetParamsTldCategory = "CLASSIC"
	EmailSecurityTopTldSpoofGetParamsTldCategoryCountry EmailSecurityTopTldSpoofGetParamsTldCategory = "COUNTRY"
)

func (r EmailSecurityTopTldSpoofGetParamsTldCategory) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpoofGetParamsTldCategoryClassic, EmailSecurityTopTldSpoofGetParamsTldCategoryCountry:
		return true
	}
	return false
}

type EmailSecurityTopTldSpoofGetParamsTLSVersion string

const (
	EmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_0 EmailSecurityTopTldSpoofGetParamsTLSVersion = "TLSv1_0"
	EmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_1 EmailSecurityTopTldSpoofGetParamsTLSVersion = "TLSv1_1"
	EmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_2 EmailSecurityTopTldSpoofGetParamsTLSVersion = "TLSv1_2"
	EmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_3 EmailSecurityTopTldSpoofGetParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecurityTopTldSpoofGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_0, EmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_1, EmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_2, EmailSecurityTopTldSpoofGetParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

type EmailSecurityTopTldSpoofGetResponseEnvelope struct {
	Result  EmailSecurityTopTldSpoofGetResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    emailSecurityTopTldSpoofGetResponseEnvelopeJSON `json:"-"`
}

// emailSecurityTopTldSpoofGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailSecurityTopTldSpoofGetResponseEnvelope]
type emailSecurityTopTldSpoofGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityTopTldSpoofGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityTopTldSpoofGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
