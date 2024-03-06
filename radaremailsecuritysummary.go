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

// RadarEmailSecuritySummaryService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailSecuritySummaryService] method instead.
type RadarEmailSecuritySummaryService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecuritySummaryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecuritySummaryService(opts ...option.RequestOption) (r *RadarEmailSecuritySummaryService) {
	r = &RadarEmailSecuritySummaryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per ARC validation.
func (r *RadarEmailSecuritySummaryService) ARC(ctx context.Context, query RadarEmailSecuritySummaryARCParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryARCResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySummaryARCResponseEnvelope
	path := "radar/email/security/summary/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per DKIM validation.
func (r *RadarEmailSecuritySummaryService) DKIM(ctx context.Context, query RadarEmailSecuritySummaryDKIMParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryDKIMResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySummaryDKIMResponseEnvelope
	path := "radar/email/security/summary/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per DMARC validation.
func (r *RadarEmailSecuritySummaryService) DMARC(ctx context.Context, query RadarEmailSecuritySummaryDMARCParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryDMARCResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySummaryDMARCResponseEnvelope
	path := "radar/email/security/summary/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified as MALICIOUS.
func (r *RadarEmailSecuritySummaryService) Malicious(ctx context.Context, query RadarEmailSecuritySummaryMaliciousParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryMaliciousResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySummaryMaliciousResponseEnvelope
	path := "radar/email/security/summary/malicious"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Proportion of emails categorized as either spam or legitimate (non-spam).
func (r *RadarEmailSecuritySummaryService) Spam(ctx context.Context, query RadarEmailSecuritySummarySpamParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummarySpamResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySummarySpamResponseEnvelope
	path := "radar/email/security/summary/spam"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per SPF validation.
func (r *RadarEmailSecuritySummaryService) SPF(ctx context.Context, query RadarEmailSecuritySummarySPFParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummarySPFResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySummarySPFResponseEnvelope
	path := "radar/email/security/summary/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Proportion of emails categorized as either spoof or legitimate (non-spoof).
func (r *RadarEmailSecuritySummaryService) Spoof(ctx context.Context, query RadarEmailSecuritySummarySpoofParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummarySpoofResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySummarySpoofResponseEnvelope
	path := "radar/email/security/summary/spoof"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified in Threat Categories.
func (r *RadarEmailSecuritySummaryService) ThreatCategory(ctx context.Context, query RadarEmailSecuritySummaryThreatCategoryParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryThreatCategoryResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySummaryThreatCategoryResponseEnvelope
	path := "radar/email/security/summary/threat_category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per TLS Version.
func (r *RadarEmailSecuritySummaryService) TLSVersion(ctx context.Context, query RadarEmailSecuritySummaryTLSVersionParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryTLSVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySummaryTLSVersionResponseEnvelope
	path := "radar/email/security/summary/tls_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecuritySummaryARCResponse struct {
	Meta     RadarEmailSecuritySummaryARCResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryARCResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryARCResponseJSON     `json:"-"`
}

// radarEmailSecuritySummaryARCResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummaryARCResponse]
type radarEmailSecuritySummaryARCResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryARCResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryARCResponseMeta struct {
	DateRange      []RadarEmailSecuritySummaryARCResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                 `json:"lastUpdated,required"`
	Normalization  string                                                 `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySummaryARCResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySummaryARCResponseMetaJSON           `json:"-"`
}

// radarEmailSecuritySummaryARCResponseMetaJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummaryARCResponseMeta]
type radarEmailSecuritySummaryARCResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryARCResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryARCResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryARCResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryARCResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryARCResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryARCResponseMetaDateRange]
type radarEmailSecuritySummaryARCResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryARCResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryARCResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryARCResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        radarEmailSecuritySummaryARCResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySummaryARCResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryARCResponseMetaConfidenceInfo]
type radarEmailSecuritySummaryARCResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryARCResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryARCResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryARCResponseSummary0 struct {
	Fail string                                           `json:"FAIL,required"`
	None string                                           `json:"NONE,required"`
	Pass string                                           `json:"PASS,required"`
	JSON radarEmailSecuritySummaryARCResponseSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryARCResponseSummary0JSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummaryARCResponseSummary0]
type radarEmailSecuritySummaryARCResponseSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryARCResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryARCResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryDKIMResponse struct {
	Meta     RadarEmailSecuritySummaryDKIMResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryDKIMResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryDKIMResponseJSON     `json:"-"`
}

// radarEmailSecuritySummaryDKIMResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummaryDKIMResponse]
type radarEmailSecuritySummaryDKIMResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDKIMResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryDKIMResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryDKIMResponseMeta struct {
	DateRange      []RadarEmailSecuritySummaryDKIMResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                  `json:"lastUpdated,required"`
	Normalization  string                                                  `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySummaryDKIMResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySummaryDKIMResponseMetaJSON           `json:"-"`
}

// radarEmailSecuritySummaryDKIMResponseMetaJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummaryDKIMResponseMeta]
type radarEmailSecuritySummaryDKIMResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDKIMResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryDKIMResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryDKIMResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryDKIMResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryDKIMResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryDKIMResponseMetaDateRange]
type radarEmailSecuritySummaryDKIMResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDKIMResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryDKIMResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryDKIMResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                               `json:"level"`
	JSON        radarEmailSecuritySummaryDKIMResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySummaryDKIMResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummaryDKIMResponseMetaConfidenceInfo]
type radarEmailSecuritySummaryDKIMResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDKIMResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryDKIMResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                `json:"dataSource,required"`
	Description     string                                                                `json:"description,required"`
	EventType       string                                                                `json:"eventType,required"`
	IsInstantaneous interface{}                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                             `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryDKIMResponseSummary0 struct {
	Fail string                                            `json:"FAIL,required"`
	None string                                            `json:"NONE,required"`
	Pass string                                            `json:"PASS,required"`
	JSON radarEmailSecuritySummaryDKIMResponseSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryDKIMResponseSummary0JSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummaryDKIMResponseSummary0]
type radarEmailSecuritySummaryDKIMResponseSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDKIMResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryDKIMResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryDMARCResponse struct {
	Meta     RadarEmailSecuritySummaryDMARCResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryDMARCResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryDMARCResponseJSON     `json:"-"`
}

// radarEmailSecuritySummaryDMARCResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummaryDMARCResponse]
type radarEmailSecuritySummaryDMARCResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDMARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryDMARCResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryDMARCResponseMeta struct {
	DateRange      []RadarEmailSecuritySummaryDMARCResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	Normalization  string                                                   `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySummaryDMARCResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySummaryDMARCResponseMetaJSON           `json:"-"`
}

// radarEmailSecuritySummaryDMARCResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummaryDMARCResponseMeta]
type radarEmailSecuritySummaryDMARCResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDMARCResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryDMARCResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryDMARCResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryDMARCResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryDMARCResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryDMARCResponseMetaDateRange]
type radarEmailSecuritySummaryDMARCResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDMARCResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryDMARCResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryDMARCResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        radarEmailSecuritySummaryDMARCResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySummaryDMARCResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummaryDMARCResponseMetaConfidenceInfo]
type radarEmailSecuritySummaryDMARCResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDMARCResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryDMARCResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous interface{}                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryDMARCResponseSummary0 struct {
	Fail string                                             `json:"FAIL,required"`
	None string                                             `json:"NONE,required"`
	Pass string                                             `json:"PASS,required"`
	JSON radarEmailSecuritySummaryDMARCResponseSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryDMARCResponseSummary0JSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryDMARCResponseSummary0]
type radarEmailSecuritySummaryDMARCResponseSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDMARCResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryDMARCResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryMaliciousResponse struct {
	Meta     RadarEmailSecuritySummaryMaliciousResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryMaliciousResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryMaliciousResponseJSON     `json:"-"`
}

// radarEmailSecuritySummaryMaliciousResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummaryMaliciousResponse]
type radarEmailSecuritySummaryMaliciousResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryMaliciousResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryMaliciousResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryMaliciousResponseMeta struct {
	DateRange      []RadarEmailSecuritySummaryMaliciousResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                       `json:"lastUpdated,required"`
	Normalization  string                                                       `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySummaryMaliciousResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySummaryMaliciousResponseMetaJSON           `json:"-"`
}

// radarEmailSecuritySummaryMaliciousResponseMetaJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryMaliciousResponseMeta]
type radarEmailSecuritySummaryMaliciousResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryMaliciousResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryMaliciousResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryMaliciousResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryMaliciousResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryMaliciousResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummaryMaliciousResponseMetaDateRange]
type radarEmailSecuritySummaryMaliciousResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryMaliciousResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryMaliciousResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryMaliciousResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                    `json:"level"`
	JSON        radarEmailSecuritySummaryMaliciousResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySummaryMaliciousResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummaryMaliciousResponseMetaConfidenceInfo]
type radarEmailSecuritySummaryMaliciousResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryMaliciousResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryMaliciousResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                     `json:"dataSource,required"`
	Description     string                                                                     `json:"description,required"`
	EventType       string                                                                     `json:"eventType,required"`
	IsInstantaneous interface{}                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                                  `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryMaliciousResponseSummary0 struct {
	Malicious    string                                                 `json:"MALICIOUS,required"`
	NotMalicious string                                                 `json:"NOT_MALICIOUS,required"`
	JSON         radarEmailSecuritySummaryMaliciousResponseSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryMaliciousResponseSummary0JSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryMaliciousResponseSummary0]
type radarEmailSecuritySummaryMaliciousResponseSummary0JSON struct {
	Malicious    apijson.Field
	NotMalicious apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryMaliciousResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryMaliciousResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySpamResponse struct {
	Meta     RadarEmailSecuritySummarySpamResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummarySpamResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummarySpamResponseJSON     `json:"-"`
}

// radarEmailSecuritySummarySpamResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummarySpamResponse]
type radarEmailSecuritySummarySpamResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpamResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySpamResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySpamResponseMeta struct {
	DateRange      []RadarEmailSecuritySummarySpamResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                  `json:"lastUpdated,required"`
	Normalization  string                                                  `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySummarySpamResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySummarySpamResponseMetaJSON           `json:"-"`
}

// radarEmailSecuritySummarySpamResponseMetaJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummarySpamResponseMeta]
type radarEmailSecuritySummarySpamResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpamResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySpamResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySpamResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummarySpamResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummarySpamResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummarySpamResponseMetaDateRange]
type radarEmailSecuritySummarySpamResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpamResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySpamResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySpamResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                               `json:"level"`
	JSON        radarEmailSecuritySummarySpamResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySummarySpamResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummarySpamResponseMetaConfidenceInfo]
type radarEmailSecuritySummarySpamResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpamResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySpamResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                `json:"dataSource,required"`
	Description     string                                                                `json:"description,required"`
	EventType       string                                                                `json:"eventType,required"`
	IsInstantaneous interface{}                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                             `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySpamResponseSummary0 struct {
	NotSpam string                                            `json:"NOT_SPAM,required"`
	Spam    string                                            `json:"SPAM,required"`
	JSON    radarEmailSecuritySummarySpamResponseSummary0JSON `json:"-"`
}

// radarEmailSecuritySummarySpamResponseSummary0JSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummarySpamResponseSummary0]
type radarEmailSecuritySummarySpamResponseSummary0JSON struct {
	NotSpam     apijson.Field
	Spam        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpamResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySpamResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySPFResponse struct {
	Meta     RadarEmailSecuritySummarySPFResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummarySPFResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummarySPFResponseJSON     `json:"-"`
}

// radarEmailSecuritySummarySPFResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummarySPFResponse]
type radarEmailSecuritySummarySPFResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySPFResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySPFResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySPFResponseMeta struct {
	DateRange      []RadarEmailSecuritySummarySPFResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                 `json:"lastUpdated,required"`
	Normalization  string                                                 `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySummarySPFResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySummarySPFResponseMetaJSON           `json:"-"`
}

// radarEmailSecuritySummarySPFResponseMetaJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummarySPFResponseMeta]
type radarEmailSecuritySummarySPFResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySPFResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySPFResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySPFResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummarySPFResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummarySPFResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummarySPFResponseMetaDateRange]
type radarEmailSecuritySummarySPFResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySPFResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySPFResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySPFResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        radarEmailSecuritySummarySPFResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySummarySPFResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummarySPFResponseMetaConfidenceInfo]
type radarEmailSecuritySummarySPFResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySPFResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySPFResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySPFResponseSummary0 struct {
	Fail string                                           `json:"FAIL,required"`
	None string                                           `json:"NONE,required"`
	Pass string                                           `json:"PASS,required"`
	JSON radarEmailSecuritySummarySPFResponseSummary0JSON `json:"-"`
}

// radarEmailSecuritySummarySPFResponseSummary0JSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummarySPFResponseSummary0]
type radarEmailSecuritySummarySPFResponseSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySPFResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySPFResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySpoofResponse struct {
	Meta     RadarEmailSecuritySummarySpoofResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummarySpoofResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummarySpoofResponseJSON     `json:"-"`
}

// radarEmailSecuritySummarySpoofResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummarySpoofResponse]
type radarEmailSecuritySummarySpoofResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpoofResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySpoofResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySpoofResponseMeta struct {
	DateRange      []RadarEmailSecuritySummarySpoofResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	Normalization  string                                                   `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySummarySpoofResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySummarySpoofResponseMetaJSON           `json:"-"`
}

// radarEmailSecuritySummarySpoofResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummarySpoofResponseMeta]
type radarEmailSecuritySummarySpoofResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpoofResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySpoofResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySpoofResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummarySpoofResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummarySpoofResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummarySpoofResponseMetaDateRange]
type radarEmailSecuritySummarySpoofResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpoofResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySpoofResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySpoofResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        radarEmailSecuritySummarySpoofResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySummarySpoofResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummarySpoofResponseMetaConfidenceInfo]
type radarEmailSecuritySummarySpoofResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpoofResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySpoofResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous interface{}                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySpoofResponseSummary0 struct {
	NotSpoof string                                             `json:"NOT_SPOOF,required"`
	Spoof    string                                             `json:"SPOOF,required"`
	JSON     radarEmailSecuritySummarySpoofResponseSummary0JSON `json:"-"`
}

// radarEmailSecuritySummarySpoofResponseSummary0JSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummarySpoofResponseSummary0]
type radarEmailSecuritySummarySpoofResponseSummary0JSON struct {
	NotSpoof    apijson.Field
	Spoof       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpoofResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySpoofResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryThreatCategoryResponse struct {
	Meta     RadarEmailSecuritySummaryThreatCategoryResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryThreatCategoryResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryThreatCategoryResponseJSON     `json:"-"`
}

// radarEmailSecuritySummaryThreatCategoryResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryThreatCategoryResponse]
type radarEmailSecuritySummaryThreatCategoryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryThreatCategoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryThreatCategoryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryThreatCategoryResponseMeta struct {
	DateRange      []RadarEmailSecuritySummaryThreatCategoryResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                            `json:"lastUpdated,required"`
	Normalization  string                                                            `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySummaryThreatCategoryResponseMetaJSON           `json:"-"`
}

// radarEmailSecuritySummaryThreatCategoryResponseMetaJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryThreatCategoryResponseMeta]
type radarEmailSecuritySummaryThreatCategoryResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryThreatCategoryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryThreatCategoryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryThreatCategoryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryThreatCategoryResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryThreatCategoryResponseMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummaryThreatCategoryResponseMetaDateRange]
type radarEmailSecuritySummaryThreatCategoryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryThreatCategoryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryThreatCategoryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                         `json:"level"`
	JSON        radarEmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfo]
type radarEmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                          `json:"dataSource,required"`
	Description     string                                                                          `json:"description,required"`
	EventType       string                                                                          `json:"eventType,required"`
	IsInstantaneous interface{}                                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                                       `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryThreatCategoryResponseSummary0 struct {
	BrandImpersonation  string                                                      `json:"BrandImpersonation,required"`
	CredentialHarvester string                                                      `json:"CredentialHarvester,required"`
	IdentityDeception   string                                                      `json:"IdentityDeception,required"`
	Link                string                                                      `json:"Link,required"`
	JSON                radarEmailSecuritySummaryThreatCategoryResponseSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryThreatCategoryResponseSummary0JSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummaryThreatCategoryResponseSummary0]
type radarEmailSecuritySummaryThreatCategoryResponseSummary0JSON struct {
	BrandImpersonation  apijson.Field
	CredentialHarvester apijson.Field
	IdentityDeception   apijson.Field
	Link                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryThreatCategoryResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryThreatCategoryResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryTLSVersionResponse struct {
	Meta     RadarEmailSecuritySummaryTLSVersionResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryTLSVersionResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryTLSVersionResponseJSON     `json:"-"`
}

// radarEmailSecuritySummaryTLSVersionResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummaryTLSVersionResponse]
type radarEmailSecuritySummaryTLSVersionResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryTLSVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryTLSVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryTLSVersionResponseMeta struct {
	DateRange      []RadarEmailSecuritySummaryTLSVersionResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                        `json:"lastUpdated,required"`
	Normalization  string                                                        `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySummaryTLSVersionResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySummaryTLSVersionResponseMetaJSON           `json:"-"`
}

// radarEmailSecuritySummaryTLSVersionResponseMetaJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryTLSVersionResponseMeta]
type radarEmailSecuritySummaryTLSVersionResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryTLSVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryTLSVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryTLSVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryTLSVersionResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryTLSVersionResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummaryTLSVersionResponseMetaDateRange]
type radarEmailSecuritySummaryTLSVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryTLSVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryTLSVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryTLSVersionResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                     `json:"level"`
	JSON        radarEmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummaryTLSVersionResponseMetaConfidenceInfo]
type radarEmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryTLSVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                      `json:"dataSource,required"`
	Description     string                                                                      `json:"description,required"`
	EventType       string                                                                      `json:"eventType,required"`
	IsInstantaneous interface{}                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                                   `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryTLSVersionResponseSummary0 struct {
	TLS1_0 string                                                  `json:"TLS 1.0,required"`
	TLS1_1 string                                                  `json:"TLS 1.1,required"`
	TLS1_2 string                                                  `json:"TLS 1.2,required"`
	TLS1_3 string                                                  `json:"TLS 1.3,required"`
	JSON   radarEmailSecuritySummaryTLSVersionResponseSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryTLSVersionResponseSummary0JSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryTLSVersionResponseSummary0]
type radarEmailSecuritySummaryTLSVersionResponseSummary0JSON struct {
	TLS1_0      apijson.Field
	TLS1_1      apijson.Field
	TLS1_2      apijson.Field
	TLS1_3      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryTLSVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryTLSVersionResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryARCParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummaryARCParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecuritySummaryARCParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecuritySummaryARCParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummaryARCParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummaryARCParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecuritySummaryARCParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecuritySummaryARCParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecuritySummaryARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryARCParamsDateRange string

const (
	RadarEmailSecuritySummaryARCParamsDateRange1d         RadarEmailSecuritySummaryARCParamsDateRange = "1d"
	RadarEmailSecuritySummaryARCParamsDateRange2d         RadarEmailSecuritySummaryARCParamsDateRange = "2d"
	RadarEmailSecuritySummaryARCParamsDateRange7d         RadarEmailSecuritySummaryARCParamsDateRange = "7d"
	RadarEmailSecuritySummaryARCParamsDateRange14d        RadarEmailSecuritySummaryARCParamsDateRange = "14d"
	RadarEmailSecuritySummaryARCParamsDateRange28d        RadarEmailSecuritySummaryARCParamsDateRange = "28d"
	RadarEmailSecuritySummaryARCParamsDateRange12w        RadarEmailSecuritySummaryARCParamsDateRange = "12w"
	RadarEmailSecuritySummaryARCParamsDateRange24w        RadarEmailSecuritySummaryARCParamsDateRange = "24w"
	RadarEmailSecuritySummaryARCParamsDateRange52w        RadarEmailSecuritySummaryARCParamsDateRange = "52w"
	RadarEmailSecuritySummaryARCParamsDateRange1dControl  RadarEmailSecuritySummaryARCParamsDateRange = "1dControl"
	RadarEmailSecuritySummaryARCParamsDateRange2dControl  RadarEmailSecuritySummaryARCParamsDateRange = "2dControl"
	RadarEmailSecuritySummaryARCParamsDateRange7dControl  RadarEmailSecuritySummaryARCParamsDateRange = "7dControl"
	RadarEmailSecuritySummaryARCParamsDateRange14dControl RadarEmailSecuritySummaryARCParamsDateRange = "14dControl"
	RadarEmailSecuritySummaryARCParamsDateRange28dControl RadarEmailSecuritySummaryARCParamsDateRange = "28dControl"
	RadarEmailSecuritySummaryARCParamsDateRange12wControl RadarEmailSecuritySummaryARCParamsDateRange = "12wControl"
	RadarEmailSecuritySummaryARCParamsDateRange24wControl RadarEmailSecuritySummaryARCParamsDateRange = "24wControl"
)

type RadarEmailSecuritySummaryARCParamsDKIM string

const (
	RadarEmailSecuritySummaryARCParamsDKIMPass RadarEmailSecuritySummaryARCParamsDKIM = "PASS"
	RadarEmailSecuritySummaryARCParamsDKIMNone RadarEmailSecuritySummaryARCParamsDKIM = "NONE"
	RadarEmailSecuritySummaryARCParamsDKIMFail RadarEmailSecuritySummaryARCParamsDKIM = "FAIL"
)

type RadarEmailSecuritySummaryARCParamsDMARC string

const (
	RadarEmailSecuritySummaryARCParamsDMARCPass RadarEmailSecuritySummaryARCParamsDMARC = "PASS"
	RadarEmailSecuritySummaryARCParamsDMARCNone RadarEmailSecuritySummaryARCParamsDMARC = "NONE"
	RadarEmailSecuritySummaryARCParamsDMARCFail RadarEmailSecuritySummaryARCParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummaryARCParamsFormat string

const (
	RadarEmailSecuritySummaryARCParamsFormatJson RadarEmailSecuritySummaryARCParamsFormat = "JSON"
	RadarEmailSecuritySummaryARCParamsFormatCsv  RadarEmailSecuritySummaryARCParamsFormat = "CSV"
)

type RadarEmailSecuritySummaryARCParamsSPF string

const (
	RadarEmailSecuritySummaryARCParamsSPFPass RadarEmailSecuritySummaryARCParamsSPF = "PASS"
	RadarEmailSecuritySummaryARCParamsSPFNone RadarEmailSecuritySummaryARCParamsSPF = "NONE"
	RadarEmailSecuritySummaryARCParamsSPFFail RadarEmailSecuritySummaryARCParamsSPF = "FAIL"
)

type RadarEmailSecuritySummaryARCParamsTLSVersion string

const (
	RadarEmailSecuritySummaryARCParamsTLSVersionTlSv1_0 RadarEmailSecuritySummaryARCParamsTLSVersion = "TLSv1_0"
	RadarEmailSecuritySummaryARCParamsTLSVersionTlSv1_1 RadarEmailSecuritySummaryARCParamsTLSVersion = "TLSv1_1"
	RadarEmailSecuritySummaryARCParamsTLSVersionTlSv1_2 RadarEmailSecuritySummaryARCParamsTLSVersion = "TLSv1_2"
	RadarEmailSecuritySummaryARCParamsTLSVersionTlSv1_3 RadarEmailSecuritySummaryARCParamsTLSVersion = "TLSv1_3"
)

type RadarEmailSecuritySummaryARCResponseEnvelope struct {
	Result  RadarEmailSecuritySummaryARCResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarEmailSecuritySummaryARCResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySummaryARCResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummaryARCResponseEnvelope]
type radarEmailSecuritySummaryARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryDKIMParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecuritySummaryDKIMParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummaryDKIMParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecuritySummaryDKIMParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummaryDKIMParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummaryDKIMParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecuritySummaryDKIMParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecuritySummaryDKIMParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecuritySummaryDKIMParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryDKIMParamsARC string

const (
	RadarEmailSecuritySummaryDKIMParamsARCPass RadarEmailSecuritySummaryDKIMParamsARC = "PASS"
	RadarEmailSecuritySummaryDKIMParamsARCNone RadarEmailSecuritySummaryDKIMParamsARC = "NONE"
	RadarEmailSecuritySummaryDKIMParamsARCFail RadarEmailSecuritySummaryDKIMParamsARC = "FAIL"
)

type RadarEmailSecuritySummaryDKIMParamsDateRange string

const (
	RadarEmailSecuritySummaryDKIMParamsDateRange1d         RadarEmailSecuritySummaryDKIMParamsDateRange = "1d"
	RadarEmailSecuritySummaryDKIMParamsDateRange2d         RadarEmailSecuritySummaryDKIMParamsDateRange = "2d"
	RadarEmailSecuritySummaryDKIMParamsDateRange7d         RadarEmailSecuritySummaryDKIMParamsDateRange = "7d"
	RadarEmailSecuritySummaryDKIMParamsDateRange14d        RadarEmailSecuritySummaryDKIMParamsDateRange = "14d"
	RadarEmailSecuritySummaryDKIMParamsDateRange28d        RadarEmailSecuritySummaryDKIMParamsDateRange = "28d"
	RadarEmailSecuritySummaryDKIMParamsDateRange12w        RadarEmailSecuritySummaryDKIMParamsDateRange = "12w"
	RadarEmailSecuritySummaryDKIMParamsDateRange24w        RadarEmailSecuritySummaryDKIMParamsDateRange = "24w"
	RadarEmailSecuritySummaryDKIMParamsDateRange52w        RadarEmailSecuritySummaryDKIMParamsDateRange = "52w"
	RadarEmailSecuritySummaryDKIMParamsDateRange1dControl  RadarEmailSecuritySummaryDKIMParamsDateRange = "1dControl"
	RadarEmailSecuritySummaryDKIMParamsDateRange2dControl  RadarEmailSecuritySummaryDKIMParamsDateRange = "2dControl"
	RadarEmailSecuritySummaryDKIMParamsDateRange7dControl  RadarEmailSecuritySummaryDKIMParamsDateRange = "7dControl"
	RadarEmailSecuritySummaryDKIMParamsDateRange14dControl RadarEmailSecuritySummaryDKIMParamsDateRange = "14dControl"
	RadarEmailSecuritySummaryDKIMParamsDateRange28dControl RadarEmailSecuritySummaryDKIMParamsDateRange = "28dControl"
	RadarEmailSecuritySummaryDKIMParamsDateRange12wControl RadarEmailSecuritySummaryDKIMParamsDateRange = "12wControl"
	RadarEmailSecuritySummaryDKIMParamsDateRange24wControl RadarEmailSecuritySummaryDKIMParamsDateRange = "24wControl"
)

type RadarEmailSecuritySummaryDKIMParamsDMARC string

const (
	RadarEmailSecuritySummaryDKIMParamsDMARCPass RadarEmailSecuritySummaryDKIMParamsDMARC = "PASS"
	RadarEmailSecuritySummaryDKIMParamsDMARCNone RadarEmailSecuritySummaryDKIMParamsDMARC = "NONE"
	RadarEmailSecuritySummaryDKIMParamsDMARCFail RadarEmailSecuritySummaryDKIMParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummaryDKIMParamsFormat string

const (
	RadarEmailSecuritySummaryDKIMParamsFormatJson RadarEmailSecuritySummaryDKIMParamsFormat = "JSON"
	RadarEmailSecuritySummaryDKIMParamsFormatCsv  RadarEmailSecuritySummaryDKIMParamsFormat = "CSV"
)

type RadarEmailSecuritySummaryDKIMParamsSPF string

const (
	RadarEmailSecuritySummaryDKIMParamsSPFPass RadarEmailSecuritySummaryDKIMParamsSPF = "PASS"
	RadarEmailSecuritySummaryDKIMParamsSPFNone RadarEmailSecuritySummaryDKIMParamsSPF = "NONE"
	RadarEmailSecuritySummaryDKIMParamsSPFFail RadarEmailSecuritySummaryDKIMParamsSPF = "FAIL"
)

type RadarEmailSecuritySummaryDKIMParamsTLSVersion string

const (
	RadarEmailSecuritySummaryDKIMParamsTLSVersionTlSv1_0 RadarEmailSecuritySummaryDKIMParamsTLSVersion = "TLSv1_0"
	RadarEmailSecuritySummaryDKIMParamsTLSVersionTlSv1_1 RadarEmailSecuritySummaryDKIMParamsTLSVersion = "TLSv1_1"
	RadarEmailSecuritySummaryDKIMParamsTLSVersionTlSv1_2 RadarEmailSecuritySummaryDKIMParamsTLSVersion = "TLSv1_2"
	RadarEmailSecuritySummaryDKIMParamsTLSVersionTlSv1_3 RadarEmailSecuritySummaryDKIMParamsTLSVersion = "TLSv1_3"
)

type RadarEmailSecuritySummaryDKIMResponseEnvelope struct {
	Result  RadarEmailSecuritySummaryDKIMResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarEmailSecuritySummaryDKIMResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySummaryDKIMResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummaryDKIMResponseEnvelope]
type radarEmailSecuritySummaryDKIMResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDKIMResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryDKIMResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryDMARCParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecuritySummaryDMARCParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummaryDMARCParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecuritySummaryDMARCParamsDKIM] `query:"dkim"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummaryDMARCParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummaryDMARCParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecuritySummaryDMARCParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecuritySummaryDMARCParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecuritySummaryDMARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryDMARCParamsARC string

const (
	RadarEmailSecuritySummaryDMARCParamsARCPass RadarEmailSecuritySummaryDMARCParamsARC = "PASS"
	RadarEmailSecuritySummaryDMARCParamsARCNone RadarEmailSecuritySummaryDMARCParamsARC = "NONE"
	RadarEmailSecuritySummaryDMARCParamsARCFail RadarEmailSecuritySummaryDMARCParamsARC = "FAIL"
)

type RadarEmailSecuritySummaryDMARCParamsDateRange string

const (
	RadarEmailSecuritySummaryDMARCParamsDateRange1d         RadarEmailSecuritySummaryDMARCParamsDateRange = "1d"
	RadarEmailSecuritySummaryDMARCParamsDateRange2d         RadarEmailSecuritySummaryDMARCParamsDateRange = "2d"
	RadarEmailSecuritySummaryDMARCParamsDateRange7d         RadarEmailSecuritySummaryDMARCParamsDateRange = "7d"
	RadarEmailSecuritySummaryDMARCParamsDateRange14d        RadarEmailSecuritySummaryDMARCParamsDateRange = "14d"
	RadarEmailSecuritySummaryDMARCParamsDateRange28d        RadarEmailSecuritySummaryDMARCParamsDateRange = "28d"
	RadarEmailSecuritySummaryDMARCParamsDateRange12w        RadarEmailSecuritySummaryDMARCParamsDateRange = "12w"
	RadarEmailSecuritySummaryDMARCParamsDateRange24w        RadarEmailSecuritySummaryDMARCParamsDateRange = "24w"
	RadarEmailSecuritySummaryDMARCParamsDateRange52w        RadarEmailSecuritySummaryDMARCParamsDateRange = "52w"
	RadarEmailSecuritySummaryDMARCParamsDateRange1dControl  RadarEmailSecuritySummaryDMARCParamsDateRange = "1dControl"
	RadarEmailSecuritySummaryDMARCParamsDateRange2dControl  RadarEmailSecuritySummaryDMARCParamsDateRange = "2dControl"
	RadarEmailSecuritySummaryDMARCParamsDateRange7dControl  RadarEmailSecuritySummaryDMARCParamsDateRange = "7dControl"
	RadarEmailSecuritySummaryDMARCParamsDateRange14dControl RadarEmailSecuritySummaryDMARCParamsDateRange = "14dControl"
	RadarEmailSecuritySummaryDMARCParamsDateRange28dControl RadarEmailSecuritySummaryDMARCParamsDateRange = "28dControl"
	RadarEmailSecuritySummaryDMARCParamsDateRange12wControl RadarEmailSecuritySummaryDMARCParamsDateRange = "12wControl"
	RadarEmailSecuritySummaryDMARCParamsDateRange24wControl RadarEmailSecuritySummaryDMARCParamsDateRange = "24wControl"
)

type RadarEmailSecuritySummaryDMARCParamsDKIM string

const (
	RadarEmailSecuritySummaryDMARCParamsDKIMPass RadarEmailSecuritySummaryDMARCParamsDKIM = "PASS"
	RadarEmailSecuritySummaryDMARCParamsDKIMNone RadarEmailSecuritySummaryDMARCParamsDKIM = "NONE"
	RadarEmailSecuritySummaryDMARCParamsDKIMFail RadarEmailSecuritySummaryDMARCParamsDKIM = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummaryDMARCParamsFormat string

const (
	RadarEmailSecuritySummaryDMARCParamsFormatJson RadarEmailSecuritySummaryDMARCParamsFormat = "JSON"
	RadarEmailSecuritySummaryDMARCParamsFormatCsv  RadarEmailSecuritySummaryDMARCParamsFormat = "CSV"
)

type RadarEmailSecuritySummaryDMARCParamsSPF string

const (
	RadarEmailSecuritySummaryDMARCParamsSPFPass RadarEmailSecuritySummaryDMARCParamsSPF = "PASS"
	RadarEmailSecuritySummaryDMARCParamsSPFNone RadarEmailSecuritySummaryDMARCParamsSPF = "NONE"
	RadarEmailSecuritySummaryDMARCParamsSPFFail RadarEmailSecuritySummaryDMARCParamsSPF = "FAIL"
)

type RadarEmailSecuritySummaryDMARCParamsTLSVersion string

const (
	RadarEmailSecuritySummaryDMARCParamsTLSVersionTlSv1_0 RadarEmailSecuritySummaryDMARCParamsTLSVersion = "TLSv1_0"
	RadarEmailSecuritySummaryDMARCParamsTLSVersionTlSv1_1 RadarEmailSecuritySummaryDMARCParamsTLSVersion = "TLSv1_1"
	RadarEmailSecuritySummaryDMARCParamsTLSVersionTlSv1_2 RadarEmailSecuritySummaryDMARCParamsTLSVersion = "TLSv1_2"
	RadarEmailSecuritySummaryDMARCParamsTLSVersionTlSv1_3 RadarEmailSecuritySummaryDMARCParamsTLSVersion = "TLSv1_3"
)

type RadarEmailSecuritySummaryDMARCResponseEnvelope struct {
	Result  RadarEmailSecuritySummaryDMARCResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarEmailSecuritySummaryDMARCResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySummaryDMARCResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryDMARCResponseEnvelope]
type radarEmailSecuritySummaryDMARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDMARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryDMARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryMaliciousParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecuritySummaryMaliciousParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummaryMaliciousParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecuritySummaryMaliciousParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecuritySummaryMaliciousParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummaryMaliciousParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummaryMaliciousParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecuritySummaryMaliciousParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecuritySummaryMaliciousParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecuritySummaryMaliciousParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryMaliciousParamsARC string

const (
	RadarEmailSecuritySummaryMaliciousParamsARCPass RadarEmailSecuritySummaryMaliciousParamsARC = "PASS"
	RadarEmailSecuritySummaryMaliciousParamsARCNone RadarEmailSecuritySummaryMaliciousParamsARC = "NONE"
	RadarEmailSecuritySummaryMaliciousParamsARCFail RadarEmailSecuritySummaryMaliciousParamsARC = "FAIL"
)

type RadarEmailSecuritySummaryMaliciousParamsDateRange string

const (
	RadarEmailSecuritySummaryMaliciousParamsDateRange1d         RadarEmailSecuritySummaryMaliciousParamsDateRange = "1d"
	RadarEmailSecuritySummaryMaliciousParamsDateRange2d         RadarEmailSecuritySummaryMaliciousParamsDateRange = "2d"
	RadarEmailSecuritySummaryMaliciousParamsDateRange7d         RadarEmailSecuritySummaryMaliciousParamsDateRange = "7d"
	RadarEmailSecuritySummaryMaliciousParamsDateRange14d        RadarEmailSecuritySummaryMaliciousParamsDateRange = "14d"
	RadarEmailSecuritySummaryMaliciousParamsDateRange28d        RadarEmailSecuritySummaryMaliciousParamsDateRange = "28d"
	RadarEmailSecuritySummaryMaliciousParamsDateRange12w        RadarEmailSecuritySummaryMaliciousParamsDateRange = "12w"
	RadarEmailSecuritySummaryMaliciousParamsDateRange24w        RadarEmailSecuritySummaryMaliciousParamsDateRange = "24w"
	RadarEmailSecuritySummaryMaliciousParamsDateRange52w        RadarEmailSecuritySummaryMaliciousParamsDateRange = "52w"
	RadarEmailSecuritySummaryMaliciousParamsDateRange1dControl  RadarEmailSecuritySummaryMaliciousParamsDateRange = "1dControl"
	RadarEmailSecuritySummaryMaliciousParamsDateRange2dControl  RadarEmailSecuritySummaryMaliciousParamsDateRange = "2dControl"
	RadarEmailSecuritySummaryMaliciousParamsDateRange7dControl  RadarEmailSecuritySummaryMaliciousParamsDateRange = "7dControl"
	RadarEmailSecuritySummaryMaliciousParamsDateRange14dControl RadarEmailSecuritySummaryMaliciousParamsDateRange = "14dControl"
	RadarEmailSecuritySummaryMaliciousParamsDateRange28dControl RadarEmailSecuritySummaryMaliciousParamsDateRange = "28dControl"
	RadarEmailSecuritySummaryMaliciousParamsDateRange12wControl RadarEmailSecuritySummaryMaliciousParamsDateRange = "12wControl"
	RadarEmailSecuritySummaryMaliciousParamsDateRange24wControl RadarEmailSecuritySummaryMaliciousParamsDateRange = "24wControl"
)

type RadarEmailSecuritySummaryMaliciousParamsDKIM string

const (
	RadarEmailSecuritySummaryMaliciousParamsDKIMPass RadarEmailSecuritySummaryMaliciousParamsDKIM = "PASS"
	RadarEmailSecuritySummaryMaliciousParamsDKIMNone RadarEmailSecuritySummaryMaliciousParamsDKIM = "NONE"
	RadarEmailSecuritySummaryMaliciousParamsDKIMFail RadarEmailSecuritySummaryMaliciousParamsDKIM = "FAIL"
)

type RadarEmailSecuritySummaryMaliciousParamsDMARC string

const (
	RadarEmailSecuritySummaryMaliciousParamsDMARCPass RadarEmailSecuritySummaryMaliciousParamsDMARC = "PASS"
	RadarEmailSecuritySummaryMaliciousParamsDMARCNone RadarEmailSecuritySummaryMaliciousParamsDMARC = "NONE"
	RadarEmailSecuritySummaryMaliciousParamsDMARCFail RadarEmailSecuritySummaryMaliciousParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummaryMaliciousParamsFormat string

const (
	RadarEmailSecuritySummaryMaliciousParamsFormatJson RadarEmailSecuritySummaryMaliciousParamsFormat = "JSON"
	RadarEmailSecuritySummaryMaliciousParamsFormatCsv  RadarEmailSecuritySummaryMaliciousParamsFormat = "CSV"
)

type RadarEmailSecuritySummaryMaliciousParamsSPF string

const (
	RadarEmailSecuritySummaryMaliciousParamsSPFPass RadarEmailSecuritySummaryMaliciousParamsSPF = "PASS"
	RadarEmailSecuritySummaryMaliciousParamsSPFNone RadarEmailSecuritySummaryMaliciousParamsSPF = "NONE"
	RadarEmailSecuritySummaryMaliciousParamsSPFFail RadarEmailSecuritySummaryMaliciousParamsSPF = "FAIL"
)

type RadarEmailSecuritySummaryMaliciousParamsTLSVersion string

const (
	RadarEmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_0 RadarEmailSecuritySummaryMaliciousParamsTLSVersion = "TLSv1_0"
	RadarEmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_1 RadarEmailSecuritySummaryMaliciousParamsTLSVersion = "TLSv1_1"
	RadarEmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_2 RadarEmailSecuritySummaryMaliciousParamsTLSVersion = "TLSv1_2"
	RadarEmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_3 RadarEmailSecuritySummaryMaliciousParamsTLSVersion = "TLSv1_3"
)

type RadarEmailSecuritySummaryMaliciousResponseEnvelope struct {
	Result  RadarEmailSecuritySummaryMaliciousResponse             `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarEmailSecuritySummaryMaliciousResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySummaryMaliciousResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryMaliciousResponseEnvelope]
type radarEmailSecuritySummaryMaliciousResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryMaliciousResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryMaliciousResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySpamParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecuritySummarySpamParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummarySpamParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecuritySummarySpamParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecuritySummarySpamParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummarySpamParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummarySpamParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecuritySummarySpamParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecuritySummarySpamParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecuritySummarySpamParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummarySpamParamsARC string

const (
	RadarEmailSecuritySummarySpamParamsARCPass RadarEmailSecuritySummarySpamParamsARC = "PASS"
	RadarEmailSecuritySummarySpamParamsARCNone RadarEmailSecuritySummarySpamParamsARC = "NONE"
	RadarEmailSecuritySummarySpamParamsARCFail RadarEmailSecuritySummarySpamParamsARC = "FAIL"
)

type RadarEmailSecuritySummarySpamParamsDateRange string

const (
	RadarEmailSecuritySummarySpamParamsDateRange1d         RadarEmailSecuritySummarySpamParamsDateRange = "1d"
	RadarEmailSecuritySummarySpamParamsDateRange2d         RadarEmailSecuritySummarySpamParamsDateRange = "2d"
	RadarEmailSecuritySummarySpamParamsDateRange7d         RadarEmailSecuritySummarySpamParamsDateRange = "7d"
	RadarEmailSecuritySummarySpamParamsDateRange14d        RadarEmailSecuritySummarySpamParamsDateRange = "14d"
	RadarEmailSecuritySummarySpamParamsDateRange28d        RadarEmailSecuritySummarySpamParamsDateRange = "28d"
	RadarEmailSecuritySummarySpamParamsDateRange12w        RadarEmailSecuritySummarySpamParamsDateRange = "12w"
	RadarEmailSecuritySummarySpamParamsDateRange24w        RadarEmailSecuritySummarySpamParamsDateRange = "24w"
	RadarEmailSecuritySummarySpamParamsDateRange52w        RadarEmailSecuritySummarySpamParamsDateRange = "52w"
	RadarEmailSecuritySummarySpamParamsDateRange1dControl  RadarEmailSecuritySummarySpamParamsDateRange = "1dControl"
	RadarEmailSecuritySummarySpamParamsDateRange2dControl  RadarEmailSecuritySummarySpamParamsDateRange = "2dControl"
	RadarEmailSecuritySummarySpamParamsDateRange7dControl  RadarEmailSecuritySummarySpamParamsDateRange = "7dControl"
	RadarEmailSecuritySummarySpamParamsDateRange14dControl RadarEmailSecuritySummarySpamParamsDateRange = "14dControl"
	RadarEmailSecuritySummarySpamParamsDateRange28dControl RadarEmailSecuritySummarySpamParamsDateRange = "28dControl"
	RadarEmailSecuritySummarySpamParamsDateRange12wControl RadarEmailSecuritySummarySpamParamsDateRange = "12wControl"
	RadarEmailSecuritySummarySpamParamsDateRange24wControl RadarEmailSecuritySummarySpamParamsDateRange = "24wControl"
)

type RadarEmailSecuritySummarySpamParamsDKIM string

const (
	RadarEmailSecuritySummarySpamParamsDKIMPass RadarEmailSecuritySummarySpamParamsDKIM = "PASS"
	RadarEmailSecuritySummarySpamParamsDKIMNone RadarEmailSecuritySummarySpamParamsDKIM = "NONE"
	RadarEmailSecuritySummarySpamParamsDKIMFail RadarEmailSecuritySummarySpamParamsDKIM = "FAIL"
)

type RadarEmailSecuritySummarySpamParamsDMARC string

const (
	RadarEmailSecuritySummarySpamParamsDMARCPass RadarEmailSecuritySummarySpamParamsDMARC = "PASS"
	RadarEmailSecuritySummarySpamParamsDMARCNone RadarEmailSecuritySummarySpamParamsDMARC = "NONE"
	RadarEmailSecuritySummarySpamParamsDMARCFail RadarEmailSecuritySummarySpamParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummarySpamParamsFormat string

const (
	RadarEmailSecuritySummarySpamParamsFormatJson RadarEmailSecuritySummarySpamParamsFormat = "JSON"
	RadarEmailSecuritySummarySpamParamsFormatCsv  RadarEmailSecuritySummarySpamParamsFormat = "CSV"
)

type RadarEmailSecuritySummarySpamParamsSPF string

const (
	RadarEmailSecuritySummarySpamParamsSPFPass RadarEmailSecuritySummarySpamParamsSPF = "PASS"
	RadarEmailSecuritySummarySpamParamsSPFNone RadarEmailSecuritySummarySpamParamsSPF = "NONE"
	RadarEmailSecuritySummarySpamParamsSPFFail RadarEmailSecuritySummarySpamParamsSPF = "FAIL"
)

type RadarEmailSecuritySummarySpamParamsTLSVersion string

const (
	RadarEmailSecuritySummarySpamParamsTLSVersionTlSv1_0 RadarEmailSecuritySummarySpamParamsTLSVersion = "TLSv1_0"
	RadarEmailSecuritySummarySpamParamsTLSVersionTlSv1_1 RadarEmailSecuritySummarySpamParamsTLSVersion = "TLSv1_1"
	RadarEmailSecuritySummarySpamParamsTLSVersionTlSv1_2 RadarEmailSecuritySummarySpamParamsTLSVersion = "TLSv1_2"
	RadarEmailSecuritySummarySpamParamsTLSVersionTlSv1_3 RadarEmailSecuritySummarySpamParamsTLSVersion = "TLSv1_3"
)

type RadarEmailSecuritySummarySpamResponseEnvelope struct {
	Result  RadarEmailSecuritySummarySpamResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarEmailSecuritySummarySpamResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySummarySpamResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummarySpamResponseEnvelope]
type radarEmailSecuritySummarySpamResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpamResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySpamResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySPFParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecuritySummarySPFParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummarySPFParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecuritySummarySPFParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecuritySummarySPFParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummarySPFParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecuritySummarySPFParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecuritySummarySPFParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecuritySummarySPFParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummarySPFParamsARC string

const (
	RadarEmailSecuritySummarySPFParamsARCPass RadarEmailSecuritySummarySPFParamsARC = "PASS"
	RadarEmailSecuritySummarySPFParamsARCNone RadarEmailSecuritySummarySPFParamsARC = "NONE"
	RadarEmailSecuritySummarySPFParamsARCFail RadarEmailSecuritySummarySPFParamsARC = "FAIL"
)

type RadarEmailSecuritySummarySPFParamsDateRange string

const (
	RadarEmailSecuritySummarySPFParamsDateRange1d         RadarEmailSecuritySummarySPFParamsDateRange = "1d"
	RadarEmailSecuritySummarySPFParamsDateRange2d         RadarEmailSecuritySummarySPFParamsDateRange = "2d"
	RadarEmailSecuritySummarySPFParamsDateRange7d         RadarEmailSecuritySummarySPFParamsDateRange = "7d"
	RadarEmailSecuritySummarySPFParamsDateRange14d        RadarEmailSecuritySummarySPFParamsDateRange = "14d"
	RadarEmailSecuritySummarySPFParamsDateRange28d        RadarEmailSecuritySummarySPFParamsDateRange = "28d"
	RadarEmailSecuritySummarySPFParamsDateRange12w        RadarEmailSecuritySummarySPFParamsDateRange = "12w"
	RadarEmailSecuritySummarySPFParamsDateRange24w        RadarEmailSecuritySummarySPFParamsDateRange = "24w"
	RadarEmailSecuritySummarySPFParamsDateRange52w        RadarEmailSecuritySummarySPFParamsDateRange = "52w"
	RadarEmailSecuritySummarySPFParamsDateRange1dControl  RadarEmailSecuritySummarySPFParamsDateRange = "1dControl"
	RadarEmailSecuritySummarySPFParamsDateRange2dControl  RadarEmailSecuritySummarySPFParamsDateRange = "2dControl"
	RadarEmailSecuritySummarySPFParamsDateRange7dControl  RadarEmailSecuritySummarySPFParamsDateRange = "7dControl"
	RadarEmailSecuritySummarySPFParamsDateRange14dControl RadarEmailSecuritySummarySPFParamsDateRange = "14dControl"
	RadarEmailSecuritySummarySPFParamsDateRange28dControl RadarEmailSecuritySummarySPFParamsDateRange = "28dControl"
	RadarEmailSecuritySummarySPFParamsDateRange12wControl RadarEmailSecuritySummarySPFParamsDateRange = "12wControl"
	RadarEmailSecuritySummarySPFParamsDateRange24wControl RadarEmailSecuritySummarySPFParamsDateRange = "24wControl"
)

type RadarEmailSecuritySummarySPFParamsDKIM string

const (
	RadarEmailSecuritySummarySPFParamsDKIMPass RadarEmailSecuritySummarySPFParamsDKIM = "PASS"
	RadarEmailSecuritySummarySPFParamsDKIMNone RadarEmailSecuritySummarySPFParamsDKIM = "NONE"
	RadarEmailSecuritySummarySPFParamsDKIMFail RadarEmailSecuritySummarySPFParamsDKIM = "FAIL"
)

type RadarEmailSecuritySummarySPFParamsDMARC string

const (
	RadarEmailSecuritySummarySPFParamsDMARCPass RadarEmailSecuritySummarySPFParamsDMARC = "PASS"
	RadarEmailSecuritySummarySPFParamsDMARCNone RadarEmailSecuritySummarySPFParamsDMARC = "NONE"
	RadarEmailSecuritySummarySPFParamsDMARCFail RadarEmailSecuritySummarySPFParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummarySPFParamsFormat string

const (
	RadarEmailSecuritySummarySPFParamsFormatJson RadarEmailSecuritySummarySPFParamsFormat = "JSON"
	RadarEmailSecuritySummarySPFParamsFormatCsv  RadarEmailSecuritySummarySPFParamsFormat = "CSV"
)

type RadarEmailSecuritySummarySPFParamsTLSVersion string

const (
	RadarEmailSecuritySummarySPFParamsTLSVersionTlSv1_0 RadarEmailSecuritySummarySPFParamsTLSVersion = "TLSv1_0"
	RadarEmailSecuritySummarySPFParamsTLSVersionTlSv1_1 RadarEmailSecuritySummarySPFParamsTLSVersion = "TLSv1_1"
	RadarEmailSecuritySummarySPFParamsTLSVersionTlSv1_2 RadarEmailSecuritySummarySPFParamsTLSVersion = "TLSv1_2"
	RadarEmailSecuritySummarySPFParamsTLSVersionTlSv1_3 RadarEmailSecuritySummarySPFParamsTLSVersion = "TLSv1_3"
)

type RadarEmailSecuritySummarySPFResponseEnvelope struct {
	Result  RadarEmailSecuritySummarySPFResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarEmailSecuritySummarySPFResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySummarySPFResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummarySPFResponseEnvelope]
type radarEmailSecuritySummarySPFResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySPFResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySPFResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummarySpoofParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecuritySummarySpoofParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummarySpoofParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecuritySummarySpoofParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecuritySummarySpoofParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummarySpoofParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummarySpoofParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecuritySummarySpoofParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecuritySummarySpoofParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecuritySummarySpoofParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummarySpoofParamsARC string

const (
	RadarEmailSecuritySummarySpoofParamsARCPass RadarEmailSecuritySummarySpoofParamsARC = "PASS"
	RadarEmailSecuritySummarySpoofParamsARCNone RadarEmailSecuritySummarySpoofParamsARC = "NONE"
	RadarEmailSecuritySummarySpoofParamsARCFail RadarEmailSecuritySummarySpoofParamsARC = "FAIL"
)

type RadarEmailSecuritySummarySpoofParamsDateRange string

const (
	RadarEmailSecuritySummarySpoofParamsDateRange1d         RadarEmailSecuritySummarySpoofParamsDateRange = "1d"
	RadarEmailSecuritySummarySpoofParamsDateRange2d         RadarEmailSecuritySummarySpoofParamsDateRange = "2d"
	RadarEmailSecuritySummarySpoofParamsDateRange7d         RadarEmailSecuritySummarySpoofParamsDateRange = "7d"
	RadarEmailSecuritySummarySpoofParamsDateRange14d        RadarEmailSecuritySummarySpoofParamsDateRange = "14d"
	RadarEmailSecuritySummarySpoofParamsDateRange28d        RadarEmailSecuritySummarySpoofParamsDateRange = "28d"
	RadarEmailSecuritySummarySpoofParamsDateRange12w        RadarEmailSecuritySummarySpoofParamsDateRange = "12w"
	RadarEmailSecuritySummarySpoofParamsDateRange24w        RadarEmailSecuritySummarySpoofParamsDateRange = "24w"
	RadarEmailSecuritySummarySpoofParamsDateRange52w        RadarEmailSecuritySummarySpoofParamsDateRange = "52w"
	RadarEmailSecuritySummarySpoofParamsDateRange1dControl  RadarEmailSecuritySummarySpoofParamsDateRange = "1dControl"
	RadarEmailSecuritySummarySpoofParamsDateRange2dControl  RadarEmailSecuritySummarySpoofParamsDateRange = "2dControl"
	RadarEmailSecuritySummarySpoofParamsDateRange7dControl  RadarEmailSecuritySummarySpoofParamsDateRange = "7dControl"
	RadarEmailSecuritySummarySpoofParamsDateRange14dControl RadarEmailSecuritySummarySpoofParamsDateRange = "14dControl"
	RadarEmailSecuritySummarySpoofParamsDateRange28dControl RadarEmailSecuritySummarySpoofParamsDateRange = "28dControl"
	RadarEmailSecuritySummarySpoofParamsDateRange12wControl RadarEmailSecuritySummarySpoofParamsDateRange = "12wControl"
	RadarEmailSecuritySummarySpoofParamsDateRange24wControl RadarEmailSecuritySummarySpoofParamsDateRange = "24wControl"
)

type RadarEmailSecuritySummarySpoofParamsDKIM string

const (
	RadarEmailSecuritySummarySpoofParamsDKIMPass RadarEmailSecuritySummarySpoofParamsDKIM = "PASS"
	RadarEmailSecuritySummarySpoofParamsDKIMNone RadarEmailSecuritySummarySpoofParamsDKIM = "NONE"
	RadarEmailSecuritySummarySpoofParamsDKIMFail RadarEmailSecuritySummarySpoofParamsDKIM = "FAIL"
)

type RadarEmailSecuritySummarySpoofParamsDMARC string

const (
	RadarEmailSecuritySummarySpoofParamsDMARCPass RadarEmailSecuritySummarySpoofParamsDMARC = "PASS"
	RadarEmailSecuritySummarySpoofParamsDMARCNone RadarEmailSecuritySummarySpoofParamsDMARC = "NONE"
	RadarEmailSecuritySummarySpoofParamsDMARCFail RadarEmailSecuritySummarySpoofParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummarySpoofParamsFormat string

const (
	RadarEmailSecuritySummarySpoofParamsFormatJson RadarEmailSecuritySummarySpoofParamsFormat = "JSON"
	RadarEmailSecuritySummarySpoofParamsFormatCsv  RadarEmailSecuritySummarySpoofParamsFormat = "CSV"
)

type RadarEmailSecuritySummarySpoofParamsSPF string

const (
	RadarEmailSecuritySummarySpoofParamsSPFPass RadarEmailSecuritySummarySpoofParamsSPF = "PASS"
	RadarEmailSecuritySummarySpoofParamsSPFNone RadarEmailSecuritySummarySpoofParamsSPF = "NONE"
	RadarEmailSecuritySummarySpoofParamsSPFFail RadarEmailSecuritySummarySpoofParamsSPF = "FAIL"
)

type RadarEmailSecuritySummarySpoofParamsTLSVersion string

const (
	RadarEmailSecuritySummarySpoofParamsTLSVersionTlSv1_0 RadarEmailSecuritySummarySpoofParamsTLSVersion = "TLSv1_0"
	RadarEmailSecuritySummarySpoofParamsTLSVersionTlSv1_1 RadarEmailSecuritySummarySpoofParamsTLSVersion = "TLSv1_1"
	RadarEmailSecuritySummarySpoofParamsTLSVersionTlSv1_2 RadarEmailSecuritySummarySpoofParamsTLSVersion = "TLSv1_2"
	RadarEmailSecuritySummarySpoofParamsTLSVersionTlSv1_3 RadarEmailSecuritySummarySpoofParamsTLSVersion = "TLSv1_3"
)

type RadarEmailSecuritySummarySpoofResponseEnvelope struct {
	Result  RadarEmailSecuritySummarySpoofResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarEmailSecuritySummarySpoofResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySummarySpoofResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummarySpoofResponseEnvelope]
type radarEmailSecuritySummarySpoofResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpoofResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummarySpoofResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryThreatCategoryParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecuritySummaryThreatCategoryParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummaryThreatCategoryParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecuritySummaryThreatCategoryParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecuritySummaryThreatCategoryParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummaryThreatCategoryParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummaryThreatCategoryParamsSPF] `query:"spf"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarEmailSecuritySummaryThreatCategoryParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecuritySummaryThreatCategoryParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecuritySummaryThreatCategoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryThreatCategoryParamsARC string

const (
	RadarEmailSecuritySummaryThreatCategoryParamsARCPass RadarEmailSecuritySummaryThreatCategoryParamsARC = "PASS"
	RadarEmailSecuritySummaryThreatCategoryParamsARCNone RadarEmailSecuritySummaryThreatCategoryParamsARC = "NONE"
	RadarEmailSecuritySummaryThreatCategoryParamsARCFail RadarEmailSecuritySummaryThreatCategoryParamsARC = "FAIL"
)

type RadarEmailSecuritySummaryThreatCategoryParamsDateRange string

const (
	RadarEmailSecuritySummaryThreatCategoryParamsDateRange1d         RadarEmailSecuritySummaryThreatCategoryParamsDateRange = "1d"
	RadarEmailSecuritySummaryThreatCategoryParamsDateRange2d         RadarEmailSecuritySummaryThreatCategoryParamsDateRange = "2d"
	RadarEmailSecuritySummaryThreatCategoryParamsDateRange7d         RadarEmailSecuritySummaryThreatCategoryParamsDateRange = "7d"
	RadarEmailSecuritySummaryThreatCategoryParamsDateRange14d        RadarEmailSecuritySummaryThreatCategoryParamsDateRange = "14d"
	RadarEmailSecuritySummaryThreatCategoryParamsDateRange28d        RadarEmailSecuritySummaryThreatCategoryParamsDateRange = "28d"
	RadarEmailSecuritySummaryThreatCategoryParamsDateRange12w        RadarEmailSecuritySummaryThreatCategoryParamsDateRange = "12w"
	RadarEmailSecuritySummaryThreatCategoryParamsDateRange24w        RadarEmailSecuritySummaryThreatCategoryParamsDateRange = "24w"
	RadarEmailSecuritySummaryThreatCategoryParamsDateRange52w        RadarEmailSecuritySummaryThreatCategoryParamsDateRange = "52w"
	RadarEmailSecuritySummaryThreatCategoryParamsDateRange1dControl  RadarEmailSecuritySummaryThreatCategoryParamsDateRange = "1dControl"
	RadarEmailSecuritySummaryThreatCategoryParamsDateRange2dControl  RadarEmailSecuritySummaryThreatCategoryParamsDateRange = "2dControl"
	RadarEmailSecuritySummaryThreatCategoryParamsDateRange7dControl  RadarEmailSecuritySummaryThreatCategoryParamsDateRange = "7dControl"
	RadarEmailSecuritySummaryThreatCategoryParamsDateRange14dControl RadarEmailSecuritySummaryThreatCategoryParamsDateRange = "14dControl"
	RadarEmailSecuritySummaryThreatCategoryParamsDateRange28dControl RadarEmailSecuritySummaryThreatCategoryParamsDateRange = "28dControl"
	RadarEmailSecuritySummaryThreatCategoryParamsDateRange12wControl RadarEmailSecuritySummaryThreatCategoryParamsDateRange = "12wControl"
	RadarEmailSecuritySummaryThreatCategoryParamsDateRange24wControl RadarEmailSecuritySummaryThreatCategoryParamsDateRange = "24wControl"
)

type RadarEmailSecuritySummaryThreatCategoryParamsDKIM string

const (
	RadarEmailSecuritySummaryThreatCategoryParamsDKIMPass RadarEmailSecuritySummaryThreatCategoryParamsDKIM = "PASS"
	RadarEmailSecuritySummaryThreatCategoryParamsDKIMNone RadarEmailSecuritySummaryThreatCategoryParamsDKIM = "NONE"
	RadarEmailSecuritySummaryThreatCategoryParamsDKIMFail RadarEmailSecuritySummaryThreatCategoryParamsDKIM = "FAIL"
)

type RadarEmailSecuritySummaryThreatCategoryParamsDMARC string

const (
	RadarEmailSecuritySummaryThreatCategoryParamsDMARCPass RadarEmailSecuritySummaryThreatCategoryParamsDMARC = "PASS"
	RadarEmailSecuritySummaryThreatCategoryParamsDMARCNone RadarEmailSecuritySummaryThreatCategoryParamsDMARC = "NONE"
	RadarEmailSecuritySummaryThreatCategoryParamsDMARCFail RadarEmailSecuritySummaryThreatCategoryParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummaryThreatCategoryParamsFormat string

const (
	RadarEmailSecuritySummaryThreatCategoryParamsFormatJson RadarEmailSecuritySummaryThreatCategoryParamsFormat = "JSON"
	RadarEmailSecuritySummaryThreatCategoryParamsFormatCsv  RadarEmailSecuritySummaryThreatCategoryParamsFormat = "CSV"
)

type RadarEmailSecuritySummaryThreatCategoryParamsSPF string

const (
	RadarEmailSecuritySummaryThreatCategoryParamsSPFPass RadarEmailSecuritySummaryThreatCategoryParamsSPF = "PASS"
	RadarEmailSecuritySummaryThreatCategoryParamsSPFNone RadarEmailSecuritySummaryThreatCategoryParamsSPF = "NONE"
	RadarEmailSecuritySummaryThreatCategoryParamsSPFFail RadarEmailSecuritySummaryThreatCategoryParamsSPF = "FAIL"
)

type RadarEmailSecuritySummaryThreatCategoryParamsTLSVersion string

const (
	RadarEmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_0 RadarEmailSecuritySummaryThreatCategoryParamsTLSVersion = "TLSv1_0"
	RadarEmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_1 RadarEmailSecuritySummaryThreatCategoryParamsTLSVersion = "TLSv1_1"
	RadarEmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_2 RadarEmailSecuritySummaryThreatCategoryParamsTLSVersion = "TLSv1_2"
	RadarEmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_3 RadarEmailSecuritySummaryThreatCategoryParamsTLSVersion = "TLSv1_3"
)

type RadarEmailSecuritySummaryThreatCategoryResponseEnvelope struct {
	Result  RadarEmailSecuritySummaryThreatCategoryResponse             `json:"result,required"`
	Success bool                                                        `json:"success,required"`
	JSON    radarEmailSecuritySummaryThreatCategoryResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySummaryThreatCategoryResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummaryThreatCategoryResponseEnvelope]
type radarEmailSecuritySummaryThreatCategoryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryThreatCategoryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryThreatCategoryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryTLSVersionParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecuritySummaryTLSVersionParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummaryTLSVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecuritySummaryTLSVersionParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecuritySummaryTLSVersionParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummaryTLSVersionParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummaryTLSVersionParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecuritySummaryTLSVersionParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecuritySummaryTLSVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryTLSVersionParamsARC string

const (
	RadarEmailSecuritySummaryTLSVersionParamsARCPass RadarEmailSecuritySummaryTLSVersionParamsARC = "PASS"
	RadarEmailSecuritySummaryTLSVersionParamsARCNone RadarEmailSecuritySummaryTLSVersionParamsARC = "NONE"
	RadarEmailSecuritySummaryTLSVersionParamsARCFail RadarEmailSecuritySummaryTLSVersionParamsARC = "FAIL"
)

type RadarEmailSecuritySummaryTLSVersionParamsDateRange string

const (
	RadarEmailSecuritySummaryTLSVersionParamsDateRange1d         RadarEmailSecuritySummaryTLSVersionParamsDateRange = "1d"
	RadarEmailSecuritySummaryTLSVersionParamsDateRange2d         RadarEmailSecuritySummaryTLSVersionParamsDateRange = "2d"
	RadarEmailSecuritySummaryTLSVersionParamsDateRange7d         RadarEmailSecuritySummaryTLSVersionParamsDateRange = "7d"
	RadarEmailSecuritySummaryTLSVersionParamsDateRange14d        RadarEmailSecuritySummaryTLSVersionParamsDateRange = "14d"
	RadarEmailSecuritySummaryTLSVersionParamsDateRange28d        RadarEmailSecuritySummaryTLSVersionParamsDateRange = "28d"
	RadarEmailSecuritySummaryTLSVersionParamsDateRange12w        RadarEmailSecuritySummaryTLSVersionParamsDateRange = "12w"
	RadarEmailSecuritySummaryTLSVersionParamsDateRange24w        RadarEmailSecuritySummaryTLSVersionParamsDateRange = "24w"
	RadarEmailSecuritySummaryTLSVersionParamsDateRange52w        RadarEmailSecuritySummaryTLSVersionParamsDateRange = "52w"
	RadarEmailSecuritySummaryTLSVersionParamsDateRange1dControl  RadarEmailSecuritySummaryTLSVersionParamsDateRange = "1dControl"
	RadarEmailSecuritySummaryTLSVersionParamsDateRange2dControl  RadarEmailSecuritySummaryTLSVersionParamsDateRange = "2dControl"
	RadarEmailSecuritySummaryTLSVersionParamsDateRange7dControl  RadarEmailSecuritySummaryTLSVersionParamsDateRange = "7dControl"
	RadarEmailSecuritySummaryTLSVersionParamsDateRange14dControl RadarEmailSecuritySummaryTLSVersionParamsDateRange = "14dControl"
	RadarEmailSecuritySummaryTLSVersionParamsDateRange28dControl RadarEmailSecuritySummaryTLSVersionParamsDateRange = "28dControl"
	RadarEmailSecuritySummaryTLSVersionParamsDateRange12wControl RadarEmailSecuritySummaryTLSVersionParamsDateRange = "12wControl"
	RadarEmailSecuritySummaryTLSVersionParamsDateRange24wControl RadarEmailSecuritySummaryTLSVersionParamsDateRange = "24wControl"
)

type RadarEmailSecuritySummaryTLSVersionParamsDKIM string

const (
	RadarEmailSecuritySummaryTLSVersionParamsDKIMPass RadarEmailSecuritySummaryTLSVersionParamsDKIM = "PASS"
	RadarEmailSecuritySummaryTLSVersionParamsDKIMNone RadarEmailSecuritySummaryTLSVersionParamsDKIM = "NONE"
	RadarEmailSecuritySummaryTLSVersionParamsDKIMFail RadarEmailSecuritySummaryTLSVersionParamsDKIM = "FAIL"
)

type RadarEmailSecuritySummaryTLSVersionParamsDMARC string

const (
	RadarEmailSecuritySummaryTLSVersionParamsDMARCPass RadarEmailSecuritySummaryTLSVersionParamsDMARC = "PASS"
	RadarEmailSecuritySummaryTLSVersionParamsDMARCNone RadarEmailSecuritySummaryTLSVersionParamsDMARC = "NONE"
	RadarEmailSecuritySummaryTLSVersionParamsDMARCFail RadarEmailSecuritySummaryTLSVersionParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummaryTLSVersionParamsFormat string

const (
	RadarEmailSecuritySummaryTLSVersionParamsFormatJson RadarEmailSecuritySummaryTLSVersionParamsFormat = "JSON"
	RadarEmailSecuritySummaryTLSVersionParamsFormatCsv  RadarEmailSecuritySummaryTLSVersionParamsFormat = "CSV"
)

type RadarEmailSecuritySummaryTLSVersionParamsSPF string

const (
	RadarEmailSecuritySummaryTLSVersionParamsSPFPass RadarEmailSecuritySummaryTLSVersionParamsSPF = "PASS"
	RadarEmailSecuritySummaryTLSVersionParamsSPFNone RadarEmailSecuritySummaryTLSVersionParamsSPF = "NONE"
	RadarEmailSecuritySummaryTLSVersionParamsSPFFail RadarEmailSecuritySummaryTLSVersionParamsSPF = "FAIL"
)

type RadarEmailSecuritySummaryTLSVersionResponseEnvelope struct {
	Result  RadarEmailSecuritySummaryTLSVersionResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarEmailSecuritySummaryTLSVersionResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySummaryTLSVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryTLSVersionResponseEnvelope]
type radarEmailSecuritySummaryTLSVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryTLSVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryTLSVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
