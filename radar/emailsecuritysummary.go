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

// EmailSecuritySummaryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewEmailSecuritySummaryService]
// method instead.
type EmailSecuritySummaryService struct {
	Options []option.RequestOption
}

// NewEmailSecuritySummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailSecuritySummaryService(opts ...option.RequestOption) (r *EmailSecuritySummaryService) {
	r = &EmailSecuritySummaryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per ARC validation.
func (r *EmailSecuritySummaryService) ARC(ctx context.Context, query EmailSecuritySummaryARCParams, opts ...option.RequestOption) (res *EmailSecuritySummaryARCResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecuritySummaryARCResponseEnvelope
	path := "radar/email/security/summary/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per DKIM validation.
func (r *EmailSecuritySummaryService) DKIM(ctx context.Context, query EmailSecuritySummaryDKIMParams, opts ...option.RequestOption) (res *EmailSecuritySummaryDKIMResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecuritySummaryDKIMResponseEnvelope
	path := "radar/email/security/summary/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per DMARC validation.
func (r *EmailSecuritySummaryService) DMARC(ctx context.Context, query EmailSecuritySummaryDMARCParams, opts ...option.RequestOption) (res *EmailSecuritySummaryDMARCResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecuritySummaryDMARCResponseEnvelope
	path := "radar/email/security/summary/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified as MALICIOUS.
func (r *EmailSecuritySummaryService) Malicious(ctx context.Context, query EmailSecuritySummaryMaliciousParams, opts ...option.RequestOption) (res *EmailSecuritySummaryMaliciousResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecuritySummaryMaliciousResponseEnvelope
	path := "radar/email/security/summary/malicious"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Proportion of emails categorized as either spam or legitimate (non-spam).
func (r *EmailSecuritySummaryService) Spam(ctx context.Context, query EmailSecuritySummarySpamParams, opts ...option.RequestOption) (res *EmailSecuritySummarySpamResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecuritySummarySpamResponseEnvelope
	path := "radar/email/security/summary/spam"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per SPF validation.
func (r *EmailSecuritySummaryService) SPF(ctx context.Context, query EmailSecuritySummarySPFParams, opts ...option.RequestOption) (res *EmailSecuritySummarySPFResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecuritySummarySPFResponseEnvelope
	path := "radar/email/security/summary/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Proportion of emails categorized as either spoof or legitimate (non-spoof).
func (r *EmailSecuritySummaryService) Spoof(ctx context.Context, query EmailSecuritySummarySpoofParams, opts ...option.RequestOption) (res *EmailSecuritySummarySpoofResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecuritySummarySpoofResponseEnvelope
	path := "radar/email/security/summary/spoof"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified in Threat Categories.
func (r *EmailSecuritySummaryService) ThreatCategory(ctx context.Context, query EmailSecuritySummaryThreatCategoryParams, opts ...option.RequestOption) (res *EmailSecuritySummaryThreatCategoryResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecuritySummaryThreatCategoryResponseEnvelope
	path := "radar/email/security/summary/threat_category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per TLS Version.
func (r *EmailSecuritySummaryService) TLSVersion(ctx context.Context, query EmailSecuritySummaryTLSVersionParams, opts ...option.RequestOption) (res *EmailSecuritySummaryTLSVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailSecuritySummaryTLSVersionResponseEnvelope
	path := "radar/email/security/summary/tls_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailSecuritySummaryARCResponse struct {
	Meta     EmailSecuritySummaryARCResponseMeta     `json:"meta,required"`
	Summary0 EmailSecuritySummaryARCResponseSummary0 `json:"summary_0,required"`
	JSON     emailSecuritySummaryARCResponseJSON     `json:"-"`
}

// emailSecuritySummaryARCResponseJSON contains the JSON metadata for the struct
// [EmailSecuritySummaryARCResponse]
type emailSecuritySummaryARCResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryARCResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryARCResponseMeta struct {
	DateRange      []EmailSecuritySummaryARCResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                            `json:"lastUpdated,required"`
	Normalization  string                                            `json:"normalization,required"`
	ConfidenceInfo EmailSecuritySummaryARCResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           emailSecuritySummaryARCResponseMetaJSON           `json:"-"`
}

// emailSecuritySummaryARCResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecuritySummaryARCResponseMeta]
type emailSecuritySummaryARCResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecuritySummaryARCResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryARCResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryARCResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      emailSecuritySummaryARCResponseMetaDateRangeJSON `json:"-"`
}

// emailSecuritySummaryARCResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [EmailSecuritySummaryARCResponseMetaDateRange]
type emailSecuritySummaryARCResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryARCResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryARCResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryARCResponseMetaConfidenceInfo struct {
	Annotations []EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                         `json:"level"`
	JSON        emailSecuritySummaryARCResponseMetaConfidenceInfoJSON         `json:"-"`
}

// emailSecuritySummaryARCResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [EmailSecuritySummaryARCResponseMetaConfidenceInfo]
type emailSecuritySummaryARCResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryARCResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryARCResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                          `json:"dataSource,required"`
	Description     string                                                          `json:"description,required"`
	EventType       string                                                          `json:"eventType,required"`
	IsInstantaneous interface{}                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                       `json:"startTime" format:"date-time"`
	JSON            emailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotation]
type emailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryARCResponseSummary0 struct {
	Fail string                                      `json:"FAIL,required"`
	None string                                      `json:"NONE,required"`
	Pass string                                      `json:"PASS,required"`
	JSON emailSecuritySummaryARCResponseSummary0JSON `json:"-"`
}

// emailSecuritySummaryARCResponseSummary0JSON contains the JSON metadata for the
// struct [EmailSecuritySummaryARCResponseSummary0]
type emailSecuritySummaryARCResponseSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryARCResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryARCResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryDKIMResponse struct {
	Meta     EmailSecuritySummaryDKIMResponseMeta     `json:"meta,required"`
	Summary0 EmailSecuritySummaryDKIMResponseSummary0 `json:"summary_0,required"`
	JSON     emailSecuritySummaryDKIMResponseJSON     `json:"-"`
}

// emailSecuritySummaryDKIMResponseJSON contains the JSON metadata for the struct
// [EmailSecuritySummaryDKIMResponse]
type emailSecuritySummaryDKIMResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryDKIMResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDKIMResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryDKIMResponseMeta struct {
	DateRange      []EmailSecuritySummaryDKIMResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                             `json:"lastUpdated,required"`
	Normalization  string                                             `json:"normalization,required"`
	ConfidenceInfo EmailSecuritySummaryDKIMResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           emailSecuritySummaryDKIMResponseMetaJSON           `json:"-"`
}

// emailSecuritySummaryDKIMResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecuritySummaryDKIMResponseMeta]
type emailSecuritySummaryDKIMResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecuritySummaryDKIMResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDKIMResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryDKIMResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      emailSecuritySummaryDKIMResponseMetaDateRangeJSON `json:"-"`
}

// emailSecuritySummaryDKIMResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [EmailSecuritySummaryDKIMResponseMetaDateRange]
type emailSecuritySummaryDKIMResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryDKIMResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDKIMResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryDKIMResponseMetaConfidenceInfo struct {
	Annotations []EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                          `json:"level"`
	JSON        emailSecuritySummaryDKIMResponseMetaConfidenceInfoJSON         `json:"-"`
}

// emailSecuritySummaryDKIMResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [EmailSecuritySummaryDKIMResponseMetaConfidenceInfo]
type emailSecuritySummaryDKIMResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryDKIMResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDKIMResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                           `json:"dataSource,required"`
	Description     string                                                           `json:"description,required"`
	EventType       string                                                           `json:"eventType,required"`
	IsInstantaneous interface{}                                                      `json:"isInstantaneous,required"`
	EndTime         time.Time                                                        `json:"endTime" format:"date-time"`
	LinkedURL       string                                                           `json:"linkedUrl"`
	StartTime       time.Time                                                        `json:"startTime" format:"date-time"`
	JSON            emailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotation]
type emailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryDKIMResponseSummary0 struct {
	Fail string                                       `json:"FAIL,required"`
	None string                                       `json:"NONE,required"`
	Pass string                                       `json:"PASS,required"`
	JSON emailSecuritySummaryDKIMResponseSummary0JSON `json:"-"`
}

// emailSecuritySummaryDKIMResponseSummary0JSON contains the JSON metadata for the
// struct [EmailSecuritySummaryDKIMResponseSummary0]
type emailSecuritySummaryDKIMResponseSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryDKIMResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDKIMResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryDMARCResponse struct {
	Meta     EmailSecuritySummaryDMARCResponseMeta     `json:"meta,required"`
	Summary0 EmailSecuritySummaryDMARCResponseSummary0 `json:"summary_0,required"`
	JSON     emailSecuritySummaryDMARCResponseJSON     `json:"-"`
}

// emailSecuritySummaryDMARCResponseJSON contains the JSON metadata for the struct
// [EmailSecuritySummaryDMARCResponse]
type emailSecuritySummaryDMARCResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryDMARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDMARCResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryDMARCResponseMeta struct {
	DateRange      []EmailSecuritySummaryDMARCResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                              `json:"lastUpdated,required"`
	Normalization  string                                              `json:"normalization,required"`
	ConfidenceInfo EmailSecuritySummaryDMARCResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           emailSecuritySummaryDMARCResponseMetaJSON           `json:"-"`
}

// emailSecuritySummaryDMARCResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecuritySummaryDMARCResponseMeta]
type emailSecuritySummaryDMARCResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecuritySummaryDMARCResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDMARCResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryDMARCResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      emailSecuritySummaryDMARCResponseMetaDateRangeJSON `json:"-"`
}

// emailSecuritySummaryDMARCResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [EmailSecuritySummaryDMARCResponseMetaDateRange]
type emailSecuritySummaryDMARCResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryDMARCResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDMARCResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryDMARCResponseMetaConfidenceInfo struct {
	Annotations []EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                           `json:"level"`
	JSON        emailSecuritySummaryDMARCResponseMetaConfidenceInfoJSON         `json:"-"`
}

// emailSecuritySummaryDMARCResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [EmailSecuritySummaryDMARCResponseMetaConfidenceInfo]
type emailSecuritySummaryDMARCResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryDMARCResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDMARCResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                            `json:"dataSource,required"`
	Description     string                                                            `json:"description,required"`
	EventType       string                                                            `json:"eventType,required"`
	IsInstantaneous interface{}                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                         `json:"startTime" format:"date-time"`
	JSON            emailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotation]
type emailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryDMARCResponseSummary0 struct {
	Fail string                                        `json:"FAIL,required"`
	None string                                        `json:"NONE,required"`
	Pass string                                        `json:"PASS,required"`
	JSON emailSecuritySummaryDMARCResponseSummary0JSON `json:"-"`
}

// emailSecuritySummaryDMARCResponseSummary0JSON contains the JSON metadata for the
// struct [EmailSecuritySummaryDMARCResponseSummary0]
type emailSecuritySummaryDMARCResponseSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryDMARCResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDMARCResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryMaliciousResponse struct {
	Meta     EmailSecuritySummaryMaliciousResponseMeta     `json:"meta,required"`
	Summary0 EmailSecuritySummaryMaliciousResponseSummary0 `json:"summary_0,required"`
	JSON     emailSecuritySummaryMaliciousResponseJSON     `json:"-"`
}

// emailSecuritySummaryMaliciousResponseJSON contains the JSON metadata for the
// struct [EmailSecuritySummaryMaliciousResponse]
type emailSecuritySummaryMaliciousResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryMaliciousResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryMaliciousResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryMaliciousResponseMeta struct {
	DateRange      []EmailSecuritySummaryMaliciousResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                  `json:"lastUpdated,required"`
	Normalization  string                                                  `json:"normalization,required"`
	ConfidenceInfo EmailSecuritySummaryMaliciousResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           emailSecuritySummaryMaliciousResponseMetaJSON           `json:"-"`
}

// emailSecuritySummaryMaliciousResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecuritySummaryMaliciousResponseMeta]
type emailSecuritySummaryMaliciousResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecuritySummaryMaliciousResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryMaliciousResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryMaliciousResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      emailSecuritySummaryMaliciousResponseMetaDateRangeJSON `json:"-"`
}

// emailSecuritySummaryMaliciousResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailSecuritySummaryMaliciousResponseMetaDateRange]
type emailSecuritySummaryMaliciousResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryMaliciousResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryMaliciousResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryMaliciousResponseMetaConfidenceInfo struct {
	Annotations []EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                               `json:"level"`
	JSON        emailSecuritySummaryMaliciousResponseMetaConfidenceInfoJSON         `json:"-"`
}

// emailSecuritySummaryMaliciousResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [EmailSecuritySummaryMaliciousResponseMetaConfidenceInfo]
type emailSecuritySummaryMaliciousResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryMaliciousResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryMaliciousResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                `json:"dataSource,required"`
	Description     string                                                                `json:"description,required"`
	EventType       string                                                                `json:"eventType,required"`
	IsInstantaneous interface{}                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                             `json:"startTime" format:"date-time"`
	JSON            emailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotation]
type emailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryMaliciousResponseSummary0 struct {
	Malicious    string                                            `json:"MALICIOUS,required"`
	NotMalicious string                                            `json:"NOT_MALICIOUS,required"`
	JSON         emailSecuritySummaryMaliciousResponseSummary0JSON `json:"-"`
}

// emailSecuritySummaryMaliciousResponseSummary0JSON contains the JSON metadata for
// the struct [EmailSecuritySummaryMaliciousResponseSummary0]
type emailSecuritySummaryMaliciousResponseSummary0JSON struct {
	Malicious    apijson.Field
	NotMalicious apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EmailSecuritySummaryMaliciousResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryMaliciousResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySpamResponse struct {
	Meta     EmailSecuritySummarySpamResponseMeta     `json:"meta,required"`
	Summary0 EmailSecuritySummarySpamResponseSummary0 `json:"summary_0,required"`
	JSON     emailSecuritySummarySpamResponseJSON     `json:"-"`
}

// emailSecuritySummarySpamResponseJSON contains the JSON metadata for the struct
// [EmailSecuritySummarySpamResponse]
type emailSecuritySummarySpamResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummarySpamResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpamResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySpamResponseMeta struct {
	DateRange      []EmailSecuritySummarySpamResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                             `json:"lastUpdated,required"`
	Normalization  string                                             `json:"normalization,required"`
	ConfidenceInfo EmailSecuritySummarySpamResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           emailSecuritySummarySpamResponseMetaJSON           `json:"-"`
}

// emailSecuritySummarySpamResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecuritySummarySpamResponseMeta]
type emailSecuritySummarySpamResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecuritySummarySpamResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpamResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySpamResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      emailSecuritySummarySpamResponseMetaDateRangeJSON `json:"-"`
}

// emailSecuritySummarySpamResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [EmailSecuritySummarySpamResponseMetaDateRange]
type emailSecuritySummarySpamResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummarySpamResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpamResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySpamResponseMetaConfidenceInfo struct {
	Annotations []EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                          `json:"level"`
	JSON        emailSecuritySummarySpamResponseMetaConfidenceInfoJSON         `json:"-"`
}

// emailSecuritySummarySpamResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [EmailSecuritySummarySpamResponseMetaConfidenceInfo]
type emailSecuritySummarySpamResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummarySpamResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpamResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                           `json:"dataSource,required"`
	Description     string                                                           `json:"description,required"`
	EventType       string                                                           `json:"eventType,required"`
	IsInstantaneous interface{}                                                      `json:"isInstantaneous,required"`
	EndTime         time.Time                                                        `json:"endTime" format:"date-time"`
	LinkedURL       string                                                           `json:"linkedUrl"`
	StartTime       time.Time                                                        `json:"startTime" format:"date-time"`
	JSON            emailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotation]
type emailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySpamResponseSummary0 struct {
	NotSpam string                                       `json:"NOT_SPAM,required"`
	Spam    string                                       `json:"SPAM,required"`
	JSON    emailSecuritySummarySpamResponseSummary0JSON `json:"-"`
}

// emailSecuritySummarySpamResponseSummary0JSON contains the JSON metadata for the
// struct [EmailSecuritySummarySpamResponseSummary0]
type emailSecuritySummarySpamResponseSummary0JSON struct {
	NotSpam     apijson.Field
	Spam        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummarySpamResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpamResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySPFResponse struct {
	Meta     EmailSecuritySummarySPFResponseMeta     `json:"meta,required"`
	Summary0 EmailSecuritySummarySPFResponseSummary0 `json:"summary_0,required"`
	JSON     emailSecuritySummarySPFResponseJSON     `json:"-"`
}

// emailSecuritySummarySPFResponseJSON contains the JSON metadata for the struct
// [EmailSecuritySummarySPFResponse]
type emailSecuritySummarySPFResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummarySPFResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySPFResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySPFResponseMeta struct {
	DateRange      []EmailSecuritySummarySPFResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                            `json:"lastUpdated,required"`
	Normalization  string                                            `json:"normalization,required"`
	ConfidenceInfo EmailSecuritySummarySPFResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           emailSecuritySummarySPFResponseMetaJSON           `json:"-"`
}

// emailSecuritySummarySPFResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecuritySummarySPFResponseMeta]
type emailSecuritySummarySPFResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecuritySummarySPFResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySPFResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySPFResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      emailSecuritySummarySPFResponseMetaDateRangeJSON `json:"-"`
}

// emailSecuritySummarySPFResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [EmailSecuritySummarySPFResponseMetaDateRange]
type emailSecuritySummarySPFResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummarySPFResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySPFResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySPFResponseMetaConfidenceInfo struct {
	Annotations []EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                         `json:"level"`
	JSON        emailSecuritySummarySPFResponseMetaConfidenceInfoJSON         `json:"-"`
}

// emailSecuritySummarySPFResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [EmailSecuritySummarySPFResponseMetaConfidenceInfo]
type emailSecuritySummarySPFResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummarySPFResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySPFResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                          `json:"dataSource,required"`
	Description     string                                                          `json:"description,required"`
	EventType       string                                                          `json:"eventType,required"`
	IsInstantaneous interface{}                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                       `json:"startTime" format:"date-time"`
	JSON            emailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotation]
type emailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySPFResponseSummary0 struct {
	Fail string                                      `json:"FAIL,required"`
	None string                                      `json:"NONE,required"`
	Pass string                                      `json:"PASS,required"`
	JSON emailSecuritySummarySPFResponseSummary0JSON `json:"-"`
}

// emailSecuritySummarySPFResponseSummary0JSON contains the JSON metadata for the
// struct [EmailSecuritySummarySPFResponseSummary0]
type emailSecuritySummarySPFResponseSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummarySPFResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySPFResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySpoofResponse struct {
	Meta     EmailSecuritySummarySpoofResponseMeta     `json:"meta,required"`
	Summary0 EmailSecuritySummarySpoofResponseSummary0 `json:"summary_0,required"`
	JSON     emailSecuritySummarySpoofResponseJSON     `json:"-"`
}

// emailSecuritySummarySpoofResponseJSON contains the JSON metadata for the struct
// [EmailSecuritySummarySpoofResponse]
type emailSecuritySummarySpoofResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummarySpoofResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpoofResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySpoofResponseMeta struct {
	DateRange      []EmailSecuritySummarySpoofResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                              `json:"lastUpdated,required"`
	Normalization  string                                              `json:"normalization,required"`
	ConfidenceInfo EmailSecuritySummarySpoofResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           emailSecuritySummarySpoofResponseMetaJSON           `json:"-"`
}

// emailSecuritySummarySpoofResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecuritySummarySpoofResponseMeta]
type emailSecuritySummarySpoofResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecuritySummarySpoofResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpoofResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySpoofResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      emailSecuritySummarySpoofResponseMetaDateRangeJSON `json:"-"`
}

// emailSecuritySummarySpoofResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [EmailSecuritySummarySpoofResponseMetaDateRange]
type emailSecuritySummarySpoofResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummarySpoofResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpoofResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySpoofResponseMetaConfidenceInfo struct {
	Annotations []EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                           `json:"level"`
	JSON        emailSecuritySummarySpoofResponseMetaConfidenceInfoJSON         `json:"-"`
}

// emailSecuritySummarySpoofResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [EmailSecuritySummarySpoofResponseMetaConfidenceInfo]
type emailSecuritySummarySpoofResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummarySpoofResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpoofResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                            `json:"dataSource,required"`
	Description     string                                                            `json:"description,required"`
	EventType       string                                                            `json:"eventType,required"`
	IsInstantaneous interface{}                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                         `json:"startTime" format:"date-time"`
	JSON            emailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotation]
type emailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySpoofResponseSummary0 struct {
	NotSpoof string                                        `json:"NOT_SPOOF,required"`
	Spoof    string                                        `json:"SPOOF,required"`
	JSON     emailSecuritySummarySpoofResponseSummary0JSON `json:"-"`
}

// emailSecuritySummarySpoofResponseSummary0JSON contains the JSON metadata for the
// struct [EmailSecuritySummarySpoofResponseSummary0]
type emailSecuritySummarySpoofResponseSummary0JSON struct {
	NotSpoof    apijson.Field
	Spoof       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummarySpoofResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpoofResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryThreatCategoryResponse struct {
	Meta     EmailSecuritySummaryThreatCategoryResponseMeta     `json:"meta,required"`
	Summary0 EmailSecuritySummaryThreatCategoryResponseSummary0 `json:"summary_0,required"`
	JSON     emailSecuritySummaryThreatCategoryResponseJSON     `json:"-"`
}

// emailSecuritySummaryThreatCategoryResponseJSON contains the JSON metadata for
// the struct [EmailSecuritySummaryThreatCategoryResponse]
type emailSecuritySummaryThreatCategoryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryThreatCategoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryThreatCategoryResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryThreatCategoryResponseMeta struct {
	DateRange      []EmailSecuritySummaryThreatCategoryResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                       `json:"lastUpdated,required"`
	Normalization  string                                                       `json:"normalization,required"`
	ConfidenceInfo EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           emailSecuritySummaryThreatCategoryResponseMetaJSON           `json:"-"`
}

// emailSecuritySummaryThreatCategoryResponseMetaJSON contains the JSON metadata
// for the struct [EmailSecuritySummaryThreatCategoryResponseMeta]
type emailSecuritySummaryThreatCategoryResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecuritySummaryThreatCategoryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryThreatCategoryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryThreatCategoryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      emailSecuritySummaryThreatCategoryResponseMetaDateRangeJSON `json:"-"`
}

// emailSecuritySummaryThreatCategoryResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [EmailSecuritySummaryThreatCategoryResponseMetaDateRange]
type emailSecuritySummaryThreatCategoryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryThreatCategoryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryThreatCategoryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfo struct {
	Annotations []EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                    `json:"level"`
	JSON        emailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoJSON         `json:"-"`
}

// emailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfo]
type emailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                     `json:"dataSource,required"`
	Description     string                                                                     `json:"description,required"`
	EventType       string                                                                     `json:"eventType,required"`
	IsInstantaneous interface{}                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                                  `json:"startTime" format:"date-time"`
	JSON            emailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotation]
type emailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryThreatCategoryResponseSummary0 struct {
	BrandImpersonation  string                                                 `json:"BrandImpersonation,required"`
	CredentialHarvester string                                                 `json:"CredentialHarvester,required"`
	IdentityDeception   string                                                 `json:"IdentityDeception,required"`
	Link                string                                                 `json:"Link,required"`
	JSON                emailSecuritySummaryThreatCategoryResponseSummary0JSON `json:"-"`
}

// emailSecuritySummaryThreatCategoryResponseSummary0JSON contains the JSON
// metadata for the struct [EmailSecuritySummaryThreatCategoryResponseSummary0]
type emailSecuritySummaryThreatCategoryResponseSummary0JSON struct {
	BrandImpersonation  apijson.Field
	CredentialHarvester apijson.Field
	IdentityDeception   apijson.Field
	Link                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EmailSecuritySummaryThreatCategoryResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryThreatCategoryResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryTLSVersionResponse struct {
	Meta     EmailSecuritySummaryTLSVersionResponseMeta     `json:"meta,required"`
	Summary0 EmailSecuritySummaryTLSVersionResponseSummary0 `json:"summary_0,required"`
	JSON     emailSecuritySummaryTLSVersionResponseJSON     `json:"-"`
}

// emailSecuritySummaryTLSVersionResponseJSON contains the JSON metadata for the
// struct [EmailSecuritySummaryTLSVersionResponse]
type emailSecuritySummaryTLSVersionResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryTLSVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryTLSVersionResponseJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryTLSVersionResponseMeta struct {
	DateRange      []EmailSecuritySummaryTLSVersionResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	Normalization  string                                                   `json:"normalization,required"`
	ConfidenceInfo EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           emailSecuritySummaryTLSVersionResponseMetaJSON           `json:"-"`
}

// emailSecuritySummaryTLSVersionResponseMetaJSON contains the JSON metadata for
// the struct [EmailSecuritySummaryTLSVersionResponseMeta]
type emailSecuritySummaryTLSVersionResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecuritySummaryTLSVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryTLSVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryTLSVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      emailSecuritySummaryTLSVersionResponseMetaDateRangeJSON `json:"-"`
}

// emailSecuritySummaryTLSVersionResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailSecuritySummaryTLSVersionResponseMetaDateRange]
type emailSecuritySummaryTLSVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryTLSVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryTLSVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfo struct {
	Annotations []EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        emailSecuritySummaryTLSVersionResponseMetaConfidenceInfoJSON         `json:"-"`
}

// emailSecuritySummaryTLSVersionResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfo]
type emailSecuritySummaryTLSVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryTLSVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous interface{}                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            emailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotation]
type emailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryTLSVersionResponseSummary0 struct {
	TLS1_0 string                                             `json:"TLS 1.0,required"`
	TLS1_1 string                                             `json:"TLS 1.1,required"`
	TLS1_2 string                                             `json:"TLS 1.2,required"`
	TLS1_3 string                                             `json:"TLS 1.3,required"`
	JSON   emailSecuritySummaryTLSVersionResponseSummary0JSON `json:"-"`
}

// emailSecuritySummaryTLSVersionResponseSummary0JSON contains the JSON metadata
// for the struct [EmailSecuritySummaryTLSVersionResponseSummary0]
type emailSecuritySummaryTLSVersionResponseSummary0JSON struct {
	TLS1_0      apijson.Field
	TLS1_1      apijson.Field
	TLS1_2      apijson.Field
	TLS1_3      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryTLSVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryTLSVersionResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryARCParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecuritySummaryARCParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecuritySummaryARCParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecuritySummaryARCParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecuritySummaryARCParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecuritySummaryARCParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecuritySummaryARCParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecuritySummaryARCParams]'s query parameters as
// `url.Values`.
func (r EmailSecuritySummaryARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailSecuritySummaryARCParamsDateRange string

const (
	EmailSecuritySummaryARCParamsDateRange1d         EmailSecuritySummaryARCParamsDateRange = "1d"
	EmailSecuritySummaryARCParamsDateRange2d         EmailSecuritySummaryARCParamsDateRange = "2d"
	EmailSecuritySummaryARCParamsDateRange7d         EmailSecuritySummaryARCParamsDateRange = "7d"
	EmailSecuritySummaryARCParamsDateRange14d        EmailSecuritySummaryARCParamsDateRange = "14d"
	EmailSecuritySummaryARCParamsDateRange28d        EmailSecuritySummaryARCParamsDateRange = "28d"
	EmailSecuritySummaryARCParamsDateRange12w        EmailSecuritySummaryARCParamsDateRange = "12w"
	EmailSecuritySummaryARCParamsDateRange24w        EmailSecuritySummaryARCParamsDateRange = "24w"
	EmailSecuritySummaryARCParamsDateRange52w        EmailSecuritySummaryARCParamsDateRange = "52w"
	EmailSecuritySummaryARCParamsDateRange1dControl  EmailSecuritySummaryARCParamsDateRange = "1dControl"
	EmailSecuritySummaryARCParamsDateRange2dControl  EmailSecuritySummaryARCParamsDateRange = "2dControl"
	EmailSecuritySummaryARCParamsDateRange7dControl  EmailSecuritySummaryARCParamsDateRange = "7dControl"
	EmailSecuritySummaryARCParamsDateRange14dControl EmailSecuritySummaryARCParamsDateRange = "14dControl"
	EmailSecuritySummaryARCParamsDateRange28dControl EmailSecuritySummaryARCParamsDateRange = "28dControl"
	EmailSecuritySummaryARCParamsDateRange12wControl EmailSecuritySummaryARCParamsDateRange = "12wControl"
	EmailSecuritySummaryARCParamsDateRange24wControl EmailSecuritySummaryARCParamsDateRange = "24wControl"
)

type EmailSecuritySummaryARCParamsDKIM string

const (
	EmailSecuritySummaryARCParamsDKIMPass EmailSecuritySummaryARCParamsDKIM = "PASS"
	EmailSecuritySummaryARCParamsDKIMNone EmailSecuritySummaryARCParamsDKIM = "NONE"
	EmailSecuritySummaryARCParamsDKIMFail EmailSecuritySummaryARCParamsDKIM = "FAIL"
)

type EmailSecuritySummaryARCParamsDMARC string

const (
	EmailSecuritySummaryARCParamsDMARCPass EmailSecuritySummaryARCParamsDMARC = "PASS"
	EmailSecuritySummaryARCParamsDMARCNone EmailSecuritySummaryARCParamsDMARC = "NONE"
	EmailSecuritySummaryARCParamsDMARCFail EmailSecuritySummaryARCParamsDMARC = "FAIL"
)

// Format results are returned in.
type EmailSecuritySummaryARCParamsFormat string

const (
	EmailSecuritySummaryARCParamsFormatJson EmailSecuritySummaryARCParamsFormat = "JSON"
	EmailSecuritySummaryARCParamsFormatCsv  EmailSecuritySummaryARCParamsFormat = "CSV"
)

type EmailSecuritySummaryARCParamsSPF string

const (
	EmailSecuritySummaryARCParamsSPFPass EmailSecuritySummaryARCParamsSPF = "PASS"
	EmailSecuritySummaryARCParamsSPFNone EmailSecuritySummaryARCParamsSPF = "NONE"
	EmailSecuritySummaryARCParamsSPFFail EmailSecuritySummaryARCParamsSPF = "FAIL"
)

type EmailSecuritySummaryARCParamsTLSVersion string

const (
	EmailSecuritySummaryARCParamsTLSVersionTlSv1_0 EmailSecuritySummaryARCParamsTLSVersion = "TLSv1_0"
	EmailSecuritySummaryARCParamsTLSVersionTlSv1_1 EmailSecuritySummaryARCParamsTLSVersion = "TLSv1_1"
	EmailSecuritySummaryARCParamsTLSVersionTlSv1_2 EmailSecuritySummaryARCParamsTLSVersion = "TLSv1_2"
	EmailSecuritySummaryARCParamsTLSVersionTlSv1_3 EmailSecuritySummaryARCParamsTLSVersion = "TLSv1_3"
)

type EmailSecuritySummaryARCResponseEnvelope struct {
	Result  EmailSecuritySummaryARCResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    emailSecuritySummaryARCResponseEnvelopeJSON `json:"-"`
}

// emailSecuritySummaryARCResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailSecuritySummaryARCResponseEnvelope]
type emailSecuritySummaryARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryDKIMParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecuritySummaryDKIMParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecuritySummaryDKIMParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecuritySummaryDKIMParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecuritySummaryDKIMParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecuritySummaryDKIMParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecuritySummaryDKIMParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecuritySummaryDKIMParams]'s query parameters as
// `url.Values`.
func (r EmailSecuritySummaryDKIMParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailSecuritySummaryDKIMParamsARC string

const (
	EmailSecuritySummaryDKIMParamsARCPass EmailSecuritySummaryDKIMParamsARC = "PASS"
	EmailSecuritySummaryDKIMParamsARCNone EmailSecuritySummaryDKIMParamsARC = "NONE"
	EmailSecuritySummaryDKIMParamsARCFail EmailSecuritySummaryDKIMParamsARC = "FAIL"
)

type EmailSecuritySummaryDKIMParamsDateRange string

const (
	EmailSecuritySummaryDKIMParamsDateRange1d         EmailSecuritySummaryDKIMParamsDateRange = "1d"
	EmailSecuritySummaryDKIMParamsDateRange2d         EmailSecuritySummaryDKIMParamsDateRange = "2d"
	EmailSecuritySummaryDKIMParamsDateRange7d         EmailSecuritySummaryDKIMParamsDateRange = "7d"
	EmailSecuritySummaryDKIMParamsDateRange14d        EmailSecuritySummaryDKIMParamsDateRange = "14d"
	EmailSecuritySummaryDKIMParamsDateRange28d        EmailSecuritySummaryDKIMParamsDateRange = "28d"
	EmailSecuritySummaryDKIMParamsDateRange12w        EmailSecuritySummaryDKIMParamsDateRange = "12w"
	EmailSecuritySummaryDKIMParamsDateRange24w        EmailSecuritySummaryDKIMParamsDateRange = "24w"
	EmailSecuritySummaryDKIMParamsDateRange52w        EmailSecuritySummaryDKIMParamsDateRange = "52w"
	EmailSecuritySummaryDKIMParamsDateRange1dControl  EmailSecuritySummaryDKIMParamsDateRange = "1dControl"
	EmailSecuritySummaryDKIMParamsDateRange2dControl  EmailSecuritySummaryDKIMParamsDateRange = "2dControl"
	EmailSecuritySummaryDKIMParamsDateRange7dControl  EmailSecuritySummaryDKIMParamsDateRange = "7dControl"
	EmailSecuritySummaryDKIMParamsDateRange14dControl EmailSecuritySummaryDKIMParamsDateRange = "14dControl"
	EmailSecuritySummaryDKIMParamsDateRange28dControl EmailSecuritySummaryDKIMParamsDateRange = "28dControl"
	EmailSecuritySummaryDKIMParamsDateRange12wControl EmailSecuritySummaryDKIMParamsDateRange = "12wControl"
	EmailSecuritySummaryDKIMParamsDateRange24wControl EmailSecuritySummaryDKIMParamsDateRange = "24wControl"
)

type EmailSecuritySummaryDKIMParamsDMARC string

const (
	EmailSecuritySummaryDKIMParamsDMARCPass EmailSecuritySummaryDKIMParamsDMARC = "PASS"
	EmailSecuritySummaryDKIMParamsDMARCNone EmailSecuritySummaryDKIMParamsDMARC = "NONE"
	EmailSecuritySummaryDKIMParamsDMARCFail EmailSecuritySummaryDKIMParamsDMARC = "FAIL"
)

// Format results are returned in.
type EmailSecuritySummaryDKIMParamsFormat string

const (
	EmailSecuritySummaryDKIMParamsFormatJson EmailSecuritySummaryDKIMParamsFormat = "JSON"
	EmailSecuritySummaryDKIMParamsFormatCsv  EmailSecuritySummaryDKIMParamsFormat = "CSV"
)

type EmailSecuritySummaryDKIMParamsSPF string

const (
	EmailSecuritySummaryDKIMParamsSPFPass EmailSecuritySummaryDKIMParamsSPF = "PASS"
	EmailSecuritySummaryDKIMParamsSPFNone EmailSecuritySummaryDKIMParamsSPF = "NONE"
	EmailSecuritySummaryDKIMParamsSPFFail EmailSecuritySummaryDKIMParamsSPF = "FAIL"
)

type EmailSecuritySummaryDKIMParamsTLSVersion string

const (
	EmailSecuritySummaryDKIMParamsTLSVersionTlSv1_0 EmailSecuritySummaryDKIMParamsTLSVersion = "TLSv1_0"
	EmailSecuritySummaryDKIMParamsTLSVersionTlSv1_1 EmailSecuritySummaryDKIMParamsTLSVersion = "TLSv1_1"
	EmailSecuritySummaryDKIMParamsTLSVersionTlSv1_2 EmailSecuritySummaryDKIMParamsTLSVersion = "TLSv1_2"
	EmailSecuritySummaryDKIMParamsTLSVersionTlSv1_3 EmailSecuritySummaryDKIMParamsTLSVersion = "TLSv1_3"
)

type EmailSecuritySummaryDKIMResponseEnvelope struct {
	Result  EmailSecuritySummaryDKIMResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    emailSecuritySummaryDKIMResponseEnvelopeJSON `json:"-"`
}

// emailSecuritySummaryDKIMResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailSecuritySummaryDKIMResponseEnvelope]
type emailSecuritySummaryDKIMResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryDKIMResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDKIMResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryDMARCParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecuritySummaryDMARCParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecuritySummaryDMARCParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecuritySummaryDMARCParamsDKIM] `query:"dkim"`
	// Format results are returned in.
	Format param.Field[EmailSecuritySummaryDMARCParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecuritySummaryDMARCParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecuritySummaryDMARCParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecuritySummaryDMARCParams]'s query parameters as
// `url.Values`.
func (r EmailSecuritySummaryDMARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailSecuritySummaryDMARCParamsARC string

const (
	EmailSecuritySummaryDMARCParamsARCPass EmailSecuritySummaryDMARCParamsARC = "PASS"
	EmailSecuritySummaryDMARCParamsARCNone EmailSecuritySummaryDMARCParamsARC = "NONE"
	EmailSecuritySummaryDMARCParamsARCFail EmailSecuritySummaryDMARCParamsARC = "FAIL"
)

type EmailSecuritySummaryDMARCParamsDateRange string

const (
	EmailSecuritySummaryDMARCParamsDateRange1d         EmailSecuritySummaryDMARCParamsDateRange = "1d"
	EmailSecuritySummaryDMARCParamsDateRange2d         EmailSecuritySummaryDMARCParamsDateRange = "2d"
	EmailSecuritySummaryDMARCParamsDateRange7d         EmailSecuritySummaryDMARCParamsDateRange = "7d"
	EmailSecuritySummaryDMARCParamsDateRange14d        EmailSecuritySummaryDMARCParamsDateRange = "14d"
	EmailSecuritySummaryDMARCParamsDateRange28d        EmailSecuritySummaryDMARCParamsDateRange = "28d"
	EmailSecuritySummaryDMARCParamsDateRange12w        EmailSecuritySummaryDMARCParamsDateRange = "12w"
	EmailSecuritySummaryDMARCParamsDateRange24w        EmailSecuritySummaryDMARCParamsDateRange = "24w"
	EmailSecuritySummaryDMARCParamsDateRange52w        EmailSecuritySummaryDMARCParamsDateRange = "52w"
	EmailSecuritySummaryDMARCParamsDateRange1dControl  EmailSecuritySummaryDMARCParamsDateRange = "1dControl"
	EmailSecuritySummaryDMARCParamsDateRange2dControl  EmailSecuritySummaryDMARCParamsDateRange = "2dControl"
	EmailSecuritySummaryDMARCParamsDateRange7dControl  EmailSecuritySummaryDMARCParamsDateRange = "7dControl"
	EmailSecuritySummaryDMARCParamsDateRange14dControl EmailSecuritySummaryDMARCParamsDateRange = "14dControl"
	EmailSecuritySummaryDMARCParamsDateRange28dControl EmailSecuritySummaryDMARCParamsDateRange = "28dControl"
	EmailSecuritySummaryDMARCParamsDateRange12wControl EmailSecuritySummaryDMARCParamsDateRange = "12wControl"
	EmailSecuritySummaryDMARCParamsDateRange24wControl EmailSecuritySummaryDMARCParamsDateRange = "24wControl"
)

type EmailSecuritySummaryDMARCParamsDKIM string

const (
	EmailSecuritySummaryDMARCParamsDKIMPass EmailSecuritySummaryDMARCParamsDKIM = "PASS"
	EmailSecuritySummaryDMARCParamsDKIMNone EmailSecuritySummaryDMARCParamsDKIM = "NONE"
	EmailSecuritySummaryDMARCParamsDKIMFail EmailSecuritySummaryDMARCParamsDKIM = "FAIL"
)

// Format results are returned in.
type EmailSecuritySummaryDMARCParamsFormat string

const (
	EmailSecuritySummaryDMARCParamsFormatJson EmailSecuritySummaryDMARCParamsFormat = "JSON"
	EmailSecuritySummaryDMARCParamsFormatCsv  EmailSecuritySummaryDMARCParamsFormat = "CSV"
)

type EmailSecuritySummaryDMARCParamsSPF string

const (
	EmailSecuritySummaryDMARCParamsSPFPass EmailSecuritySummaryDMARCParamsSPF = "PASS"
	EmailSecuritySummaryDMARCParamsSPFNone EmailSecuritySummaryDMARCParamsSPF = "NONE"
	EmailSecuritySummaryDMARCParamsSPFFail EmailSecuritySummaryDMARCParamsSPF = "FAIL"
)

type EmailSecuritySummaryDMARCParamsTLSVersion string

const (
	EmailSecuritySummaryDMARCParamsTLSVersionTlSv1_0 EmailSecuritySummaryDMARCParamsTLSVersion = "TLSv1_0"
	EmailSecuritySummaryDMARCParamsTLSVersionTlSv1_1 EmailSecuritySummaryDMARCParamsTLSVersion = "TLSv1_1"
	EmailSecuritySummaryDMARCParamsTLSVersionTlSv1_2 EmailSecuritySummaryDMARCParamsTLSVersion = "TLSv1_2"
	EmailSecuritySummaryDMARCParamsTLSVersionTlSv1_3 EmailSecuritySummaryDMARCParamsTLSVersion = "TLSv1_3"
)

type EmailSecuritySummaryDMARCResponseEnvelope struct {
	Result  EmailSecuritySummaryDMARCResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    emailSecuritySummaryDMARCResponseEnvelopeJSON `json:"-"`
}

// emailSecuritySummaryDMARCResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailSecuritySummaryDMARCResponseEnvelope]
type emailSecuritySummaryDMARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryDMARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDMARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryMaliciousParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecuritySummaryMaliciousParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecuritySummaryMaliciousParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecuritySummaryMaliciousParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecuritySummaryMaliciousParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecuritySummaryMaliciousParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecuritySummaryMaliciousParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecuritySummaryMaliciousParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecuritySummaryMaliciousParams]'s query parameters as
// `url.Values`.
func (r EmailSecuritySummaryMaliciousParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailSecuritySummaryMaliciousParamsARC string

const (
	EmailSecuritySummaryMaliciousParamsARCPass EmailSecuritySummaryMaliciousParamsARC = "PASS"
	EmailSecuritySummaryMaliciousParamsARCNone EmailSecuritySummaryMaliciousParamsARC = "NONE"
	EmailSecuritySummaryMaliciousParamsARCFail EmailSecuritySummaryMaliciousParamsARC = "FAIL"
)

type EmailSecuritySummaryMaliciousParamsDateRange string

const (
	EmailSecuritySummaryMaliciousParamsDateRange1d         EmailSecuritySummaryMaliciousParamsDateRange = "1d"
	EmailSecuritySummaryMaliciousParamsDateRange2d         EmailSecuritySummaryMaliciousParamsDateRange = "2d"
	EmailSecuritySummaryMaliciousParamsDateRange7d         EmailSecuritySummaryMaliciousParamsDateRange = "7d"
	EmailSecuritySummaryMaliciousParamsDateRange14d        EmailSecuritySummaryMaliciousParamsDateRange = "14d"
	EmailSecuritySummaryMaliciousParamsDateRange28d        EmailSecuritySummaryMaliciousParamsDateRange = "28d"
	EmailSecuritySummaryMaliciousParamsDateRange12w        EmailSecuritySummaryMaliciousParamsDateRange = "12w"
	EmailSecuritySummaryMaliciousParamsDateRange24w        EmailSecuritySummaryMaliciousParamsDateRange = "24w"
	EmailSecuritySummaryMaliciousParamsDateRange52w        EmailSecuritySummaryMaliciousParamsDateRange = "52w"
	EmailSecuritySummaryMaliciousParamsDateRange1dControl  EmailSecuritySummaryMaliciousParamsDateRange = "1dControl"
	EmailSecuritySummaryMaliciousParamsDateRange2dControl  EmailSecuritySummaryMaliciousParamsDateRange = "2dControl"
	EmailSecuritySummaryMaliciousParamsDateRange7dControl  EmailSecuritySummaryMaliciousParamsDateRange = "7dControl"
	EmailSecuritySummaryMaliciousParamsDateRange14dControl EmailSecuritySummaryMaliciousParamsDateRange = "14dControl"
	EmailSecuritySummaryMaliciousParamsDateRange28dControl EmailSecuritySummaryMaliciousParamsDateRange = "28dControl"
	EmailSecuritySummaryMaliciousParamsDateRange12wControl EmailSecuritySummaryMaliciousParamsDateRange = "12wControl"
	EmailSecuritySummaryMaliciousParamsDateRange24wControl EmailSecuritySummaryMaliciousParamsDateRange = "24wControl"
)

type EmailSecuritySummaryMaliciousParamsDKIM string

const (
	EmailSecuritySummaryMaliciousParamsDKIMPass EmailSecuritySummaryMaliciousParamsDKIM = "PASS"
	EmailSecuritySummaryMaliciousParamsDKIMNone EmailSecuritySummaryMaliciousParamsDKIM = "NONE"
	EmailSecuritySummaryMaliciousParamsDKIMFail EmailSecuritySummaryMaliciousParamsDKIM = "FAIL"
)

type EmailSecuritySummaryMaliciousParamsDMARC string

const (
	EmailSecuritySummaryMaliciousParamsDMARCPass EmailSecuritySummaryMaliciousParamsDMARC = "PASS"
	EmailSecuritySummaryMaliciousParamsDMARCNone EmailSecuritySummaryMaliciousParamsDMARC = "NONE"
	EmailSecuritySummaryMaliciousParamsDMARCFail EmailSecuritySummaryMaliciousParamsDMARC = "FAIL"
)

// Format results are returned in.
type EmailSecuritySummaryMaliciousParamsFormat string

const (
	EmailSecuritySummaryMaliciousParamsFormatJson EmailSecuritySummaryMaliciousParamsFormat = "JSON"
	EmailSecuritySummaryMaliciousParamsFormatCsv  EmailSecuritySummaryMaliciousParamsFormat = "CSV"
)

type EmailSecuritySummaryMaliciousParamsSPF string

const (
	EmailSecuritySummaryMaliciousParamsSPFPass EmailSecuritySummaryMaliciousParamsSPF = "PASS"
	EmailSecuritySummaryMaliciousParamsSPFNone EmailSecuritySummaryMaliciousParamsSPF = "NONE"
	EmailSecuritySummaryMaliciousParamsSPFFail EmailSecuritySummaryMaliciousParamsSPF = "FAIL"
)

type EmailSecuritySummaryMaliciousParamsTLSVersion string

const (
	EmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_0 EmailSecuritySummaryMaliciousParamsTLSVersion = "TLSv1_0"
	EmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_1 EmailSecuritySummaryMaliciousParamsTLSVersion = "TLSv1_1"
	EmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_2 EmailSecuritySummaryMaliciousParamsTLSVersion = "TLSv1_2"
	EmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_3 EmailSecuritySummaryMaliciousParamsTLSVersion = "TLSv1_3"
)

type EmailSecuritySummaryMaliciousResponseEnvelope struct {
	Result  EmailSecuritySummaryMaliciousResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    emailSecuritySummaryMaliciousResponseEnvelopeJSON `json:"-"`
}

// emailSecuritySummaryMaliciousResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailSecuritySummaryMaliciousResponseEnvelope]
type emailSecuritySummaryMaliciousResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryMaliciousResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryMaliciousResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySpamParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecuritySummarySpamParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecuritySummarySpamParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecuritySummarySpamParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecuritySummarySpamParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecuritySummarySpamParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecuritySummarySpamParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecuritySummarySpamParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecuritySummarySpamParams]'s query parameters as
// `url.Values`.
func (r EmailSecuritySummarySpamParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailSecuritySummarySpamParamsARC string

const (
	EmailSecuritySummarySpamParamsARCPass EmailSecuritySummarySpamParamsARC = "PASS"
	EmailSecuritySummarySpamParamsARCNone EmailSecuritySummarySpamParamsARC = "NONE"
	EmailSecuritySummarySpamParamsARCFail EmailSecuritySummarySpamParamsARC = "FAIL"
)

type EmailSecuritySummarySpamParamsDateRange string

const (
	EmailSecuritySummarySpamParamsDateRange1d         EmailSecuritySummarySpamParamsDateRange = "1d"
	EmailSecuritySummarySpamParamsDateRange2d         EmailSecuritySummarySpamParamsDateRange = "2d"
	EmailSecuritySummarySpamParamsDateRange7d         EmailSecuritySummarySpamParamsDateRange = "7d"
	EmailSecuritySummarySpamParamsDateRange14d        EmailSecuritySummarySpamParamsDateRange = "14d"
	EmailSecuritySummarySpamParamsDateRange28d        EmailSecuritySummarySpamParamsDateRange = "28d"
	EmailSecuritySummarySpamParamsDateRange12w        EmailSecuritySummarySpamParamsDateRange = "12w"
	EmailSecuritySummarySpamParamsDateRange24w        EmailSecuritySummarySpamParamsDateRange = "24w"
	EmailSecuritySummarySpamParamsDateRange52w        EmailSecuritySummarySpamParamsDateRange = "52w"
	EmailSecuritySummarySpamParamsDateRange1dControl  EmailSecuritySummarySpamParamsDateRange = "1dControl"
	EmailSecuritySummarySpamParamsDateRange2dControl  EmailSecuritySummarySpamParamsDateRange = "2dControl"
	EmailSecuritySummarySpamParamsDateRange7dControl  EmailSecuritySummarySpamParamsDateRange = "7dControl"
	EmailSecuritySummarySpamParamsDateRange14dControl EmailSecuritySummarySpamParamsDateRange = "14dControl"
	EmailSecuritySummarySpamParamsDateRange28dControl EmailSecuritySummarySpamParamsDateRange = "28dControl"
	EmailSecuritySummarySpamParamsDateRange12wControl EmailSecuritySummarySpamParamsDateRange = "12wControl"
	EmailSecuritySummarySpamParamsDateRange24wControl EmailSecuritySummarySpamParamsDateRange = "24wControl"
)

type EmailSecuritySummarySpamParamsDKIM string

const (
	EmailSecuritySummarySpamParamsDKIMPass EmailSecuritySummarySpamParamsDKIM = "PASS"
	EmailSecuritySummarySpamParamsDKIMNone EmailSecuritySummarySpamParamsDKIM = "NONE"
	EmailSecuritySummarySpamParamsDKIMFail EmailSecuritySummarySpamParamsDKIM = "FAIL"
)

type EmailSecuritySummarySpamParamsDMARC string

const (
	EmailSecuritySummarySpamParamsDMARCPass EmailSecuritySummarySpamParamsDMARC = "PASS"
	EmailSecuritySummarySpamParamsDMARCNone EmailSecuritySummarySpamParamsDMARC = "NONE"
	EmailSecuritySummarySpamParamsDMARCFail EmailSecuritySummarySpamParamsDMARC = "FAIL"
)

// Format results are returned in.
type EmailSecuritySummarySpamParamsFormat string

const (
	EmailSecuritySummarySpamParamsFormatJson EmailSecuritySummarySpamParamsFormat = "JSON"
	EmailSecuritySummarySpamParamsFormatCsv  EmailSecuritySummarySpamParamsFormat = "CSV"
)

type EmailSecuritySummarySpamParamsSPF string

const (
	EmailSecuritySummarySpamParamsSPFPass EmailSecuritySummarySpamParamsSPF = "PASS"
	EmailSecuritySummarySpamParamsSPFNone EmailSecuritySummarySpamParamsSPF = "NONE"
	EmailSecuritySummarySpamParamsSPFFail EmailSecuritySummarySpamParamsSPF = "FAIL"
)

type EmailSecuritySummarySpamParamsTLSVersion string

const (
	EmailSecuritySummarySpamParamsTLSVersionTlSv1_0 EmailSecuritySummarySpamParamsTLSVersion = "TLSv1_0"
	EmailSecuritySummarySpamParamsTLSVersionTlSv1_1 EmailSecuritySummarySpamParamsTLSVersion = "TLSv1_1"
	EmailSecuritySummarySpamParamsTLSVersionTlSv1_2 EmailSecuritySummarySpamParamsTLSVersion = "TLSv1_2"
	EmailSecuritySummarySpamParamsTLSVersionTlSv1_3 EmailSecuritySummarySpamParamsTLSVersion = "TLSv1_3"
)

type EmailSecuritySummarySpamResponseEnvelope struct {
	Result  EmailSecuritySummarySpamResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    emailSecuritySummarySpamResponseEnvelopeJSON `json:"-"`
}

// emailSecuritySummarySpamResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailSecuritySummarySpamResponseEnvelope]
type emailSecuritySummarySpamResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummarySpamResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpamResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySPFParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecuritySummarySPFParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecuritySummarySPFParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecuritySummarySPFParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecuritySummarySPFParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecuritySummarySPFParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecuritySummarySPFParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecuritySummarySPFParams]'s query parameters as
// `url.Values`.
func (r EmailSecuritySummarySPFParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailSecuritySummarySPFParamsARC string

const (
	EmailSecuritySummarySPFParamsARCPass EmailSecuritySummarySPFParamsARC = "PASS"
	EmailSecuritySummarySPFParamsARCNone EmailSecuritySummarySPFParamsARC = "NONE"
	EmailSecuritySummarySPFParamsARCFail EmailSecuritySummarySPFParamsARC = "FAIL"
)

type EmailSecuritySummarySPFParamsDateRange string

const (
	EmailSecuritySummarySPFParamsDateRange1d         EmailSecuritySummarySPFParamsDateRange = "1d"
	EmailSecuritySummarySPFParamsDateRange2d         EmailSecuritySummarySPFParamsDateRange = "2d"
	EmailSecuritySummarySPFParamsDateRange7d         EmailSecuritySummarySPFParamsDateRange = "7d"
	EmailSecuritySummarySPFParamsDateRange14d        EmailSecuritySummarySPFParamsDateRange = "14d"
	EmailSecuritySummarySPFParamsDateRange28d        EmailSecuritySummarySPFParamsDateRange = "28d"
	EmailSecuritySummarySPFParamsDateRange12w        EmailSecuritySummarySPFParamsDateRange = "12w"
	EmailSecuritySummarySPFParamsDateRange24w        EmailSecuritySummarySPFParamsDateRange = "24w"
	EmailSecuritySummarySPFParamsDateRange52w        EmailSecuritySummarySPFParamsDateRange = "52w"
	EmailSecuritySummarySPFParamsDateRange1dControl  EmailSecuritySummarySPFParamsDateRange = "1dControl"
	EmailSecuritySummarySPFParamsDateRange2dControl  EmailSecuritySummarySPFParamsDateRange = "2dControl"
	EmailSecuritySummarySPFParamsDateRange7dControl  EmailSecuritySummarySPFParamsDateRange = "7dControl"
	EmailSecuritySummarySPFParamsDateRange14dControl EmailSecuritySummarySPFParamsDateRange = "14dControl"
	EmailSecuritySummarySPFParamsDateRange28dControl EmailSecuritySummarySPFParamsDateRange = "28dControl"
	EmailSecuritySummarySPFParamsDateRange12wControl EmailSecuritySummarySPFParamsDateRange = "12wControl"
	EmailSecuritySummarySPFParamsDateRange24wControl EmailSecuritySummarySPFParamsDateRange = "24wControl"
)

type EmailSecuritySummarySPFParamsDKIM string

const (
	EmailSecuritySummarySPFParamsDKIMPass EmailSecuritySummarySPFParamsDKIM = "PASS"
	EmailSecuritySummarySPFParamsDKIMNone EmailSecuritySummarySPFParamsDKIM = "NONE"
	EmailSecuritySummarySPFParamsDKIMFail EmailSecuritySummarySPFParamsDKIM = "FAIL"
)

type EmailSecuritySummarySPFParamsDMARC string

const (
	EmailSecuritySummarySPFParamsDMARCPass EmailSecuritySummarySPFParamsDMARC = "PASS"
	EmailSecuritySummarySPFParamsDMARCNone EmailSecuritySummarySPFParamsDMARC = "NONE"
	EmailSecuritySummarySPFParamsDMARCFail EmailSecuritySummarySPFParamsDMARC = "FAIL"
)

// Format results are returned in.
type EmailSecuritySummarySPFParamsFormat string

const (
	EmailSecuritySummarySPFParamsFormatJson EmailSecuritySummarySPFParamsFormat = "JSON"
	EmailSecuritySummarySPFParamsFormatCsv  EmailSecuritySummarySPFParamsFormat = "CSV"
)

type EmailSecuritySummarySPFParamsTLSVersion string

const (
	EmailSecuritySummarySPFParamsTLSVersionTlSv1_0 EmailSecuritySummarySPFParamsTLSVersion = "TLSv1_0"
	EmailSecuritySummarySPFParamsTLSVersionTlSv1_1 EmailSecuritySummarySPFParamsTLSVersion = "TLSv1_1"
	EmailSecuritySummarySPFParamsTLSVersionTlSv1_2 EmailSecuritySummarySPFParamsTLSVersion = "TLSv1_2"
	EmailSecuritySummarySPFParamsTLSVersionTlSv1_3 EmailSecuritySummarySPFParamsTLSVersion = "TLSv1_3"
)

type EmailSecuritySummarySPFResponseEnvelope struct {
	Result  EmailSecuritySummarySPFResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    emailSecuritySummarySPFResponseEnvelopeJSON `json:"-"`
}

// emailSecuritySummarySPFResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailSecuritySummarySPFResponseEnvelope]
type emailSecuritySummarySPFResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummarySPFResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySPFResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySpoofParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecuritySummarySpoofParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecuritySummarySpoofParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecuritySummarySpoofParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecuritySummarySpoofParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecuritySummarySpoofParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecuritySummarySpoofParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecuritySummarySpoofParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecuritySummarySpoofParams]'s query parameters as
// `url.Values`.
func (r EmailSecuritySummarySpoofParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailSecuritySummarySpoofParamsARC string

const (
	EmailSecuritySummarySpoofParamsARCPass EmailSecuritySummarySpoofParamsARC = "PASS"
	EmailSecuritySummarySpoofParamsARCNone EmailSecuritySummarySpoofParamsARC = "NONE"
	EmailSecuritySummarySpoofParamsARCFail EmailSecuritySummarySpoofParamsARC = "FAIL"
)

type EmailSecuritySummarySpoofParamsDateRange string

const (
	EmailSecuritySummarySpoofParamsDateRange1d         EmailSecuritySummarySpoofParamsDateRange = "1d"
	EmailSecuritySummarySpoofParamsDateRange2d         EmailSecuritySummarySpoofParamsDateRange = "2d"
	EmailSecuritySummarySpoofParamsDateRange7d         EmailSecuritySummarySpoofParamsDateRange = "7d"
	EmailSecuritySummarySpoofParamsDateRange14d        EmailSecuritySummarySpoofParamsDateRange = "14d"
	EmailSecuritySummarySpoofParamsDateRange28d        EmailSecuritySummarySpoofParamsDateRange = "28d"
	EmailSecuritySummarySpoofParamsDateRange12w        EmailSecuritySummarySpoofParamsDateRange = "12w"
	EmailSecuritySummarySpoofParamsDateRange24w        EmailSecuritySummarySpoofParamsDateRange = "24w"
	EmailSecuritySummarySpoofParamsDateRange52w        EmailSecuritySummarySpoofParamsDateRange = "52w"
	EmailSecuritySummarySpoofParamsDateRange1dControl  EmailSecuritySummarySpoofParamsDateRange = "1dControl"
	EmailSecuritySummarySpoofParamsDateRange2dControl  EmailSecuritySummarySpoofParamsDateRange = "2dControl"
	EmailSecuritySummarySpoofParamsDateRange7dControl  EmailSecuritySummarySpoofParamsDateRange = "7dControl"
	EmailSecuritySummarySpoofParamsDateRange14dControl EmailSecuritySummarySpoofParamsDateRange = "14dControl"
	EmailSecuritySummarySpoofParamsDateRange28dControl EmailSecuritySummarySpoofParamsDateRange = "28dControl"
	EmailSecuritySummarySpoofParamsDateRange12wControl EmailSecuritySummarySpoofParamsDateRange = "12wControl"
	EmailSecuritySummarySpoofParamsDateRange24wControl EmailSecuritySummarySpoofParamsDateRange = "24wControl"
)

type EmailSecuritySummarySpoofParamsDKIM string

const (
	EmailSecuritySummarySpoofParamsDKIMPass EmailSecuritySummarySpoofParamsDKIM = "PASS"
	EmailSecuritySummarySpoofParamsDKIMNone EmailSecuritySummarySpoofParamsDKIM = "NONE"
	EmailSecuritySummarySpoofParamsDKIMFail EmailSecuritySummarySpoofParamsDKIM = "FAIL"
)

type EmailSecuritySummarySpoofParamsDMARC string

const (
	EmailSecuritySummarySpoofParamsDMARCPass EmailSecuritySummarySpoofParamsDMARC = "PASS"
	EmailSecuritySummarySpoofParamsDMARCNone EmailSecuritySummarySpoofParamsDMARC = "NONE"
	EmailSecuritySummarySpoofParamsDMARCFail EmailSecuritySummarySpoofParamsDMARC = "FAIL"
)

// Format results are returned in.
type EmailSecuritySummarySpoofParamsFormat string

const (
	EmailSecuritySummarySpoofParamsFormatJson EmailSecuritySummarySpoofParamsFormat = "JSON"
	EmailSecuritySummarySpoofParamsFormatCsv  EmailSecuritySummarySpoofParamsFormat = "CSV"
)

type EmailSecuritySummarySpoofParamsSPF string

const (
	EmailSecuritySummarySpoofParamsSPFPass EmailSecuritySummarySpoofParamsSPF = "PASS"
	EmailSecuritySummarySpoofParamsSPFNone EmailSecuritySummarySpoofParamsSPF = "NONE"
	EmailSecuritySummarySpoofParamsSPFFail EmailSecuritySummarySpoofParamsSPF = "FAIL"
)

type EmailSecuritySummarySpoofParamsTLSVersion string

const (
	EmailSecuritySummarySpoofParamsTLSVersionTlSv1_0 EmailSecuritySummarySpoofParamsTLSVersion = "TLSv1_0"
	EmailSecuritySummarySpoofParamsTLSVersionTlSv1_1 EmailSecuritySummarySpoofParamsTLSVersion = "TLSv1_1"
	EmailSecuritySummarySpoofParamsTLSVersionTlSv1_2 EmailSecuritySummarySpoofParamsTLSVersion = "TLSv1_2"
	EmailSecuritySummarySpoofParamsTLSVersionTlSv1_3 EmailSecuritySummarySpoofParamsTLSVersion = "TLSv1_3"
)

type EmailSecuritySummarySpoofResponseEnvelope struct {
	Result  EmailSecuritySummarySpoofResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    emailSecuritySummarySpoofResponseEnvelopeJSON `json:"-"`
}

// emailSecuritySummarySpoofResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailSecuritySummarySpoofResponseEnvelope]
type emailSecuritySummarySpoofResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummarySpoofResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpoofResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryThreatCategoryParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecuritySummaryThreatCategoryParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecuritySummaryThreatCategoryParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecuritySummaryThreatCategoryParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecuritySummaryThreatCategoryParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecuritySummaryThreatCategoryParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecuritySummaryThreatCategoryParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]EmailSecuritySummaryThreatCategoryParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecuritySummaryThreatCategoryParams]'s query
// parameters as `url.Values`.
func (r EmailSecuritySummaryThreatCategoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailSecuritySummaryThreatCategoryParamsARC string

const (
	EmailSecuritySummaryThreatCategoryParamsARCPass EmailSecuritySummaryThreatCategoryParamsARC = "PASS"
	EmailSecuritySummaryThreatCategoryParamsARCNone EmailSecuritySummaryThreatCategoryParamsARC = "NONE"
	EmailSecuritySummaryThreatCategoryParamsARCFail EmailSecuritySummaryThreatCategoryParamsARC = "FAIL"
)

type EmailSecuritySummaryThreatCategoryParamsDateRange string

const (
	EmailSecuritySummaryThreatCategoryParamsDateRange1d         EmailSecuritySummaryThreatCategoryParamsDateRange = "1d"
	EmailSecuritySummaryThreatCategoryParamsDateRange2d         EmailSecuritySummaryThreatCategoryParamsDateRange = "2d"
	EmailSecuritySummaryThreatCategoryParamsDateRange7d         EmailSecuritySummaryThreatCategoryParamsDateRange = "7d"
	EmailSecuritySummaryThreatCategoryParamsDateRange14d        EmailSecuritySummaryThreatCategoryParamsDateRange = "14d"
	EmailSecuritySummaryThreatCategoryParamsDateRange28d        EmailSecuritySummaryThreatCategoryParamsDateRange = "28d"
	EmailSecuritySummaryThreatCategoryParamsDateRange12w        EmailSecuritySummaryThreatCategoryParamsDateRange = "12w"
	EmailSecuritySummaryThreatCategoryParamsDateRange24w        EmailSecuritySummaryThreatCategoryParamsDateRange = "24w"
	EmailSecuritySummaryThreatCategoryParamsDateRange52w        EmailSecuritySummaryThreatCategoryParamsDateRange = "52w"
	EmailSecuritySummaryThreatCategoryParamsDateRange1dControl  EmailSecuritySummaryThreatCategoryParamsDateRange = "1dControl"
	EmailSecuritySummaryThreatCategoryParamsDateRange2dControl  EmailSecuritySummaryThreatCategoryParamsDateRange = "2dControl"
	EmailSecuritySummaryThreatCategoryParamsDateRange7dControl  EmailSecuritySummaryThreatCategoryParamsDateRange = "7dControl"
	EmailSecuritySummaryThreatCategoryParamsDateRange14dControl EmailSecuritySummaryThreatCategoryParamsDateRange = "14dControl"
	EmailSecuritySummaryThreatCategoryParamsDateRange28dControl EmailSecuritySummaryThreatCategoryParamsDateRange = "28dControl"
	EmailSecuritySummaryThreatCategoryParamsDateRange12wControl EmailSecuritySummaryThreatCategoryParamsDateRange = "12wControl"
	EmailSecuritySummaryThreatCategoryParamsDateRange24wControl EmailSecuritySummaryThreatCategoryParamsDateRange = "24wControl"
)

type EmailSecuritySummaryThreatCategoryParamsDKIM string

const (
	EmailSecuritySummaryThreatCategoryParamsDKIMPass EmailSecuritySummaryThreatCategoryParamsDKIM = "PASS"
	EmailSecuritySummaryThreatCategoryParamsDKIMNone EmailSecuritySummaryThreatCategoryParamsDKIM = "NONE"
	EmailSecuritySummaryThreatCategoryParamsDKIMFail EmailSecuritySummaryThreatCategoryParamsDKIM = "FAIL"
)

type EmailSecuritySummaryThreatCategoryParamsDMARC string

const (
	EmailSecuritySummaryThreatCategoryParamsDMARCPass EmailSecuritySummaryThreatCategoryParamsDMARC = "PASS"
	EmailSecuritySummaryThreatCategoryParamsDMARCNone EmailSecuritySummaryThreatCategoryParamsDMARC = "NONE"
	EmailSecuritySummaryThreatCategoryParamsDMARCFail EmailSecuritySummaryThreatCategoryParamsDMARC = "FAIL"
)

// Format results are returned in.
type EmailSecuritySummaryThreatCategoryParamsFormat string

const (
	EmailSecuritySummaryThreatCategoryParamsFormatJson EmailSecuritySummaryThreatCategoryParamsFormat = "JSON"
	EmailSecuritySummaryThreatCategoryParamsFormatCsv  EmailSecuritySummaryThreatCategoryParamsFormat = "CSV"
)

type EmailSecuritySummaryThreatCategoryParamsSPF string

const (
	EmailSecuritySummaryThreatCategoryParamsSPFPass EmailSecuritySummaryThreatCategoryParamsSPF = "PASS"
	EmailSecuritySummaryThreatCategoryParamsSPFNone EmailSecuritySummaryThreatCategoryParamsSPF = "NONE"
	EmailSecuritySummaryThreatCategoryParamsSPFFail EmailSecuritySummaryThreatCategoryParamsSPF = "FAIL"
)

type EmailSecuritySummaryThreatCategoryParamsTLSVersion string

const (
	EmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_0 EmailSecuritySummaryThreatCategoryParamsTLSVersion = "TLSv1_0"
	EmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_1 EmailSecuritySummaryThreatCategoryParamsTLSVersion = "TLSv1_1"
	EmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_2 EmailSecuritySummaryThreatCategoryParamsTLSVersion = "TLSv1_2"
	EmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_3 EmailSecuritySummaryThreatCategoryParamsTLSVersion = "TLSv1_3"
)

type EmailSecuritySummaryThreatCategoryResponseEnvelope struct {
	Result  EmailSecuritySummaryThreatCategoryResponse             `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    emailSecuritySummaryThreatCategoryResponseEnvelopeJSON `json:"-"`
}

// emailSecuritySummaryThreatCategoryResponseEnvelopeJSON contains the JSON
// metadata for the struct [EmailSecuritySummaryThreatCategoryResponseEnvelope]
type emailSecuritySummaryThreatCategoryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryThreatCategoryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryThreatCategoryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryTLSVersionParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailSecuritySummaryTLSVersionParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]EmailSecuritySummaryTLSVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailSecuritySummaryTLSVersionParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailSecuritySummaryTLSVersionParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailSecuritySummaryTLSVersionParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailSecuritySummaryTLSVersionParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailSecuritySummaryTLSVersionParams]'s query parameters as
// `url.Values`.
func (r EmailSecuritySummaryTLSVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailSecuritySummaryTLSVersionParamsARC string

const (
	EmailSecuritySummaryTLSVersionParamsARCPass EmailSecuritySummaryTLSVersionParamsARC = "PASS"
	EmailSecuritySummaryTLSVersionParamsARCNone EmailSecuritySummaryTLSVersionParamsARC = "NONE"
	EmailSecuritySummaryTLSVersionParamsARCFail EmailSecuritySummaryTLSVersionParamsARC = "FAIL"
)

type EmailSecuritySummaryTLSVersionParamsDateRange string

const (
	EmailSecuritySummaryTLSVersionParamsDateRange1d         EmailSecuritySummaryTLSVersionParamsDateRange = "1d"
	EmailSecuritySummaryTLSVersionParamsDateRange2d         EmailSecuritySummaryTLSVersionParamsDateRange = "2d"
	EmailSecuritySummaryTLSVersionParamsDateRange7d         EmailSecuritySummaryTLSVersionParamsDateRange = "7d"
	EmailSecuritySummaryTLSVersionParamsDateRange14d        EmailSecuritySummaryTLSVersionParamsDateRange = "14d"
	EmailSecuritySummaryTLSVersionParamsDateRange28d        EmailSecuritySummaryTLSVersionParamsDateRange = "28d"
	EmailSecuritySummaryTLSVersionParamsDateRange12w        EmailSecuritySummaryTLSVersionParamsDateRange = "12w"
	EmailSecuritySummaryTLSVersionParamsDateRange24w        EmailSecuritySummaryTLSVersionParamsDateRange = "24w"
	EmailSecuritySummaryTLSVersionParamsDateRange52w        EmailSecuritySummaryTLSVersionParamsDateRange = "52w"
	EmailSecuritySummaryTLSVersionParamsDateRange1dControl  EmailSecuritySummaryTLSVersionParamsDateRange = "1dControl"
	EmailSecuritySummaryTLSVersionParamsDateRange2dControl  EmailSecuritySummaryTLSVersionParamsDateRange = "2dControl"
	EmailSecuritySummaryTLSVersionParamsDateRange7dControl  EmailSecuritySummaryTLSVersionParamsDateRange = "7dControl"
	EmailSecuritySummaryTLSVersionParamsDateRange14dControl EmailSecuritySummaryTLSVersionParamsDateRange = "14dControl"
	EmailSecuritySummaryTLSVersionParamsDateRange28dControl EmailSecuritySummaryTLSVersionParamsDateRange = "28dControl"
	EmailSecuritySummaryTLSVersionParamsDateRange12wControl EmailSecuritySummaryTLSVersionParamsDateRange = "12wControl"
	EmailSecuritySummaryTLSVersionParamsDateRange24wControl EmailSecuritySummaryTLSVersionParamsDateRange = "24wControl"
)

type EmailSecuritySummaryTLSVersionParamsDKIM string

const (
	EmailSecuritySummaryTLSVersionParamsDKIMPass EmailSecuritySummaryTLSVersionParamsDKIM = "PASS"
	EmailSecuritySummaryTLSVersionParamsDKIMNone EmailSecuritySummaryTLSVersionParamsDKIM = "NONE"
	EmailSecuritySummaryTLSVersionParamsDKIMFail EmailSecuritySummaryTLSVersionParamsDKIM = "FAIL"
)

type EmailSecuritySummaryTLSVersionParamsDMARC string

const (
	EmailSecuritySummaryTLSVersionParamsDMARCPass EmailSecuritySummaryTLSVersionParamsDMARC = "PASS"
	EmailSecuritySummaryTLSVersionParamsDMARCNone EmailSecuritySummaryTLSVersionParamsDMARC = "NONE"
	EmailSecuritySummaryTLSVersionParamsDMARCFail EmailSecuritySummaryTLSVersionParamsDMARC = "FAIL"
)

// Format results are returned in.
type EmailSecuritySummaryTLSVersionParamsFormat string

const (
	EmailSecuritySummaryTLSVersionParamsFormatJson EmailSecuritySummaryTLSVersionParamsFormat = "JSON"
	EmailSecuritySummaryTLSVersionParamsFormatCsv  EmailSecuritySummaryTLSVersionParamsFormat = "CSV"
)

type EmailSecuritySummaryTLSVersionParamsSPF string

const (
	EmailSecuritySummaryTLSVersionParamsSPFPass EmailSecuritySummaryTLSVersionParamsSPF = "PASS"
	EmailSecuritySummaryTLSVersionParamsSPFNone EmailSecuritySummaryTLSVersionParamsSPF = "NONE"
	EmailSecuritySummaryTLSVersionParamsSPFFail EmailSecuritySummaryTLSVersionParamsSPF = "FAIL"
)

type EmailSecuritySummaryTLSVersionResponseEnvelope struct {
	Result  EmailSecuritySummaryTLSVersionResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    emailSecuritySummaryTLSVersionResponseEnvelopeJSON `json:"-"`
}

// emailSecuritySummaryTLSVersionResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailSecuritySummaryTLSVersionResponseEnvelope]
type emailSecuritySummaryTLSVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryTLSVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryTLSVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
