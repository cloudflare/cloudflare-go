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
func (r *RadarEmailSecuritySummaryService) Arc(ctx context.Context, query RadarEmailSecuritySummaryArcParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryArcResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySummaryArcResponseEnvelope
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
func (r *RadarEmailSecuritySummaryService) Dmarc(ctx context.Context, query RadarEmailSecuritySummaryDmarcParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryDmarcResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySummaryDmarcResponseEnvelope
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

type RadarEmailSecuritySummaryArcResponse struct {
	Meta     RadarEmailSecuritySummaryArcResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryArcResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryArcResponseJSON     `json:"-"`
}

// radarEmailSecuritySummaryArcResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummaryArcResponse]
type radarEmailSecuritySummaryArcResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryArcResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryArcResponseMeta struct {
	DateRange      []RadarEmailSecuritySummaryArcResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                 `json:"lastUpdated,required"`
	Normalization  string                                                 `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySummaryArcResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySummaryArcResponseMetaJSON           `json:"-"`
}

// radarEmailSecuritySummaryArcResponseMetaJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummaryArcResponseMeta]
type radarEmailSecuritySummaryArcResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryArcResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryArcResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryArcResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryArcResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryArcResponseMetaDateRange]
type radarEmailSecuritySummaryArcResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryArcResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryArcResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryArcResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        radarEmailSecuritySummaryArcResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySummaryArcResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryArcResponseMetaConfidenceInfo]
type radarEmailSecuritySummaryArcResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryArcResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryArcResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySummaryArcResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryArcResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySummaryArcResponseMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryArcResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySummaryArcResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryArcResponseSummary0 struct {
	Fail string                                           `json:"FAIL,required"`
	None string                                           `json:"NONE,required"`
	Pass string                                           `json:"PASS,required"`
	JSON radarEmailSecuritySummaryArcResponseSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryArcResponseSummary0JSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummaryArcResponseSummary0]
type radarEmailSecuritySummaryArcResponseSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryArcResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type RadarEmailSecuritySummaryDmarcResponse struct {
	Meta     RadarEmailSecuritySummaryDmarcResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryDmarcResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryDmarcResponseJSON     `json:"-"`
}

// radarEmailSecuritySummaryDmarcResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummaryDmarcResponse]
type radarEmailSecuritySummaryDmarcResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDmarcResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryDmarcResponseMeta struct {
	DateRange      []RadarEmailSecuritySummaryDmarcResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	Normalization  string                                                   `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySummaryDmarcResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySummaryDmarcResponseMetaJSON           `json:"-"`
}

// radarEmailSecuritySummaryDmarcResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummaryDmarcResponseMeta]
type radarEmailSecuritySummaryDmarcResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDmarcResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryDmarcResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryDmarcResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryDmarcResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryDmarcResponseMetaDateRange]
type radarEmailSecuritySummaryDmarcResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDmarcResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryDmarcResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryDmarcResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        radarEmailSecuritySummaryDmarcResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySummaryDmarcResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummaryDmarcResponseMetaConfidenceInfo]
type radarEmailSecuritySummaryDmarcResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDmarcResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryDmarcResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous interface{}                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySummaryDmarcResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryDmarcResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySummaryDmarcResponseMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryDmarcResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySummaryDmarcResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryDmarcResponseSummary0 struct {
	Fail string                                             `json:"FAIL,required"`
	None string                                             `json:"NONE,required"`
	Pass string                                             `json:"PASS,required"`
	JSON radarEmailSecuritySummaryDmarcResponseSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryDmarcResponseSummary0JSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryDmarcResponseSummary0]
type radarEmailSecuritySummaryDmarcResponseSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDmarcResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type RadarEmailSecuritySummaryArcParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummaryArcParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecuritySummaryArcParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecuritySummaryArcParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummaryArcParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummaryArcParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecuritySummaryArcParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecuritySummaryArcParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryArcParamsDateRange string

const (
	RadarEmailSecuritySummaryArcParamsDateRange1d         RadarEmailSecuritySummaryArcParamsDateRange = "1d"
	RadarEmailSecuritySummaryArcParamsDateRange2d         RadarEmailSecuritySummaryArcParamsDateRange = "2d"
	RadarEmailSecuritySummaryArcParamsDateRange7d         RadarEmailSecuritySummaryArcParamsDateRange = "7d"
	RadarEmailSecuritySummaryArcParamsDateRange14d        RadarEmailSecuritySummaryArcParamsDateRange = "14d"
	RadarEmailSecuritySummaryArcParamsDateRange28d        RadarEmailSecuritySummaryArcParamsDateRange = "28d"
	RadarEmailSecuritySummaryArcParamsDateRange12w        RadarEmailSecuritySummaryArcParamsDateRange = "12w"
	RadarEmailSecuritySummaryArcParamsDateRange24w        RadarEmailSecuritySummaryArcParamsDateRange = "24w"
	RadarEmailSecuritySummaryArcParamsDateRange52w        RadarEmailSecuritySummaryArcParamsDateRange = "52w"
	RadarEmailSecuritySummaryArcParamsDateRange1dControl  RadarEmailSecuritySummaryArcParamsDateRange = "1dControl"
	RadarEmailSecuritySummaryArcParamsDateRange2dControl  RadarEmailSecuritySummaryArcParamsDateRange = "2dControl"
	RadarEmailSecuritySummaryArcParamsDateRange7dControl  RadarEmailSecuritySummaryArcParamsDateRange = "7dControl"
	RadarEmailSecuritySummaryArcParamsDateRange14dControl RadarEmailSecuritySummaryArcParamsDateRange = "14dControl"
	RadarEmailSecuritySummaryArcParamsDateRange28dControl RadarEmailSecuritySummaryArcParamsDateRange = "28dControl"
	RadarEmailSecuritySummaryArcParamsDateRange12wControl RadarEmailSecuritySummaryArcParamsDateRange = "12wControl"
	RadarEmailSecuritySummaryArcParamsDateRange24wControl RadarEmailSecuritySummaryArcParamsDateRange = "24wControl"
)

type RadarEmailSecuritySummaryArcParamsDKIM string

const (
	RadarEmailSecuritySummaryArcParamsDKIMPass RadarEmailSecuritySummaryArcParamsDKIM = "PASS"
	RadarEmailSecuritySummaryArcParamsDKIMNone RadarEmailSecuritySummaryArcParamsDKIM = "NONE"
	RadarEmailSecuritySummaryArcParamsDKIMFail RadarEmailSecuritySummaryArcParamsDKIM = "FAIL"
)

type RadarEmailSecuritySummaryArcParamsDmarc string

const (
	RadarEmailSecuritySummaryArcParamsDmarcPass RadarEmailSecuritySummaryArcParamsDmarc = "PASS"
	RadarEmailSecuritySummaryArcParamsDmarcNone RadarEmailSecuritySummaryArcParamsDmarc = "NONE"
	RadarEmailSecuritySummaryArcParamsDmarcFail RadarEmailSecuritySummaryArcParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummaryArcParamsFormat string

const (
	RadarEmailSecuritySummaryArcParamsFormatJson RadarEmailSecuritySummaryArcParamsFormat = "JSON"
	RadarEmailSecuritySummaryArcParamsFormatCsv  RadarEmailSecuritySummaryArcParamsFormat = "CSV"
)

type RadarEmailSecuritySummaryArcParamsSPF string

const (
	RadarEmailSecuritySummaryArcParamsSPFPass RadarEmailSecuritySummaryArcParamsSPF = "PASS"
	RadarEmailSecuritySummaryArcParamsSPFNone RadarEmailSecuritySummaryArcParamsSPF = "NONE"
	RadarEmailSecuritySummaryArcParamsSPFFail RadarEmailSecuritySummaryArcParamsSPF = "FAIL"
)

type RadarEmailSecuritySummaryArcResponseEnvelope struct {
	Result  RadarEmailSecuritySummaryArcResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarEmailSecuritySummaryArcResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySummaryArcResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummaryArcResponseEnvelope]
type radarEmailSecuritySummaryArcResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryArcResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryDKIMParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySummaryDKIMParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummaryDKIMParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecuritySummaryDKIMParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummaryDKIMParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummaryDKIMParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecuritySummaryDKIMParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecuritySummaryDKIMParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryDKIMParamsArc string

const (
	RadarEmailSecuritySummaryDKIMParamsArcPass RadarEmailSecuritySummaryDKIMParamsArc = "PASS"
	RadarEmailSecuritySummaryDKIMParamsArcNone RadarEmailSecuritySummaryDKIMParamsArc = "NONE"
	RadarEmailSecuritySummaryDKIMParamsArcFail RadarEmailSecuritySummaryDKIMParamsArc = "FAIL"
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

type RadarEmailSecuritySummaryDKIMParamsDmarc string

const (
	RadarEmailSecuritySummaryDKIMParamsDmarcPass RadarEmailSecuritySummaryDKIMParamsDmarc = "PASS"
	RadarEmailSecuritySummaryDKIMParamsDmarcNone RadarEmailSecuritySummaryDKIMParamsDmarc = "NONE"
	RadarEmailSecuritySummaryDKIMParamsDmarcFail RadarEmailSecuritySummaryDKIMParamsDmarc = "FAIL"
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

type RadarEmailSecuritySummaryDmarcParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySummaryDmarcParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummaryDmarcParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecuritySummaryDmarcParamsDKIM] `query:"dkim"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummaryDmarcParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummaryDmarcParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecuritySummaryDmarcParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecuritySummaryDmarcParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryDmarcParamsArc string

const (
	RadarEmailSecuritySummaryDmarcParamsArcPass RadarEmailSecuritySummaryDmarcParamsArc = "PASS"
	RadarEmailSecuritySummaryDmarcParamsArcNone RadarEmailSecuritySummaryDmarcParamsArc = "NONE"
	RadarEmailSecuritySummaryDmarcParamsArcFail RadarEmailSecuritySummaryDmarcParamsArc = "FAIL"
)

type RadarEmailSecuritySummaryDmarcParamsDateRange string

const (
	RadarEmailSecuritySummaryDmarcParamsDateRange1d         RadarEmailSecuritySummaryDmarcParamsDateRange = "1d"
	RadarEmailSecuritySummaryDmarcParamsDateRange2d         RadarEmailSecuritySummaryDmarcParamsDateRange = "2d"
	RadarEmailSecuritySummaryDmarcParamsDateRange7d         RadarEmailSecuritySummaryDmarcParamsDateRange = "7d"
	RadarEmailSecuritySummaryDmarcParamsDateRange14d        RadarEmailSecuritySummaryDmarcParamsDateRange = "14d"
	RadarEmailSecuritySummaryDmarcParamsDateRange28d        RadarEmailSecuritySummaryDmarcParamsDateRange = "28d"
	RadarEmailSecuritySummaryDmarcParamsDateRange12w        RadarEmailSecuritySummaryDmarcParamsDateRange = "12w"
	RadarEmailSecuritySummaryDmarcParamsDateRange24w        RadarEmailSecuritySummaryDmarcParamsDateRange = "24w"
	RadarEmailSecuritySummaryDmarcParamsDateRange52w        RadarEmailSecuritySummaryDmarcParamsDateRange = "52w"
	RadarEmailSecuritySummaryDmarcParamsDateRange1dControl  RadarEmailSecuritySummaryDmarcParamsDateRange = "1dControl"
	RadarEmailSecuritySummaryDmarcParamsDateRange2dControl  RadarEmailSecuritySummaryDmarcParamsDateRange = "2dControl"
	RadarEmailSecuritySummaryDmarcParamsDateRange7dControl  RadarEmailSecuritySummaryDmarcParamsDateRange = "7dControl"
	RadarEmailSecuritySummaryDmarcParamsDateRange14dControl RadarEmailSecuritySummaryDmarcParamsDateRange = "14dControl"
	RadarEmailSecuritySummaryDmarcParamsDateRange28dControl RadarEmailSecuritySummaryDmarcParamsDateRange = "28dControl"
	RadarEmailSecuritySummaryDmarcParamsDateRange12wControl RadarEmailSecuritySummaryDmarcParamsDateRange = "12wControl"
	RadarEmailSecuritySummaryDmarcParamsDateRange24wControl RadarEmailSecuritySummaryDmarcParamsDateRange = "24wControl"
)

type RadarEmailSecuritySummaryDmarcParamsDKIM string

const (
	RadarEmailSecuritySummaryDmarcParamsDKIMPass RadarEmailSecuritySummaryDmarcParamsDKIM = "PASS"
	RadarEmailSecuritySummaryDmarcParamsDKIMNone RadarEmailSecuritySummaryDmarcParamsDKIM = "NONE"
	RadarEmailSecuritySummaryDmarcParamsDKIMFail RadarEmailSecuritySummaryDmarcParamsDKIM = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummaryDmarcParamsFormat string

const (
	RadarEmailSecuritySummaryDmarcParamsFormatJson RadarEmailSecuritySummaryDmarcParamsFormat = "JSON"
	RadarEmailSecuritySummaryDmarcParamsFormatCsv  RadarEmailSecuritySummaryDmarcParamsFormat = "CSV"
)

type RadarEmailSecuritySummaryDmarcParamsSPF string

const (
	RadarEmailSecuritySummaryDmarcParamsSPFPass RadarEmailSecuritySummaryDmarcParamsSPF = "PASS"
	RadarEmailSecuritySummaryDmarcParamsSPFNone RadarEmailSecuritySummaryDmarcParamsSPF = "NONE"
	RadarEmailSecuritySummaryDmarcParamsSPFFail RadarEmailSecuritySummaryDmarcParamsSPF = "FAIL"
)

type RadarEmailSecuritySummaryDmarcResponseEnvelope struct {
	Result  RadarEmailSecuritySummaryDmarcResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarEmailSecuritySummaryDmarcResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySummaryDmarcResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryDmarcResponseEnvelope]
type radarEmailSecuritySummaryDmarcResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDmarcResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryMaliciousParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySummaryMaliciousParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
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
	Dmarc param.Field[[]RadarEmailSecuritySummaryMaliciousParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummaryMaliciousParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummaryMaliciousParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecuritySummaryMaliciousParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecuritySummaryMaliciousParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryMaliciousParamsArc string

const (
	RadarEmailSecuritySummaryMaliciousParamsArcPass RadarEmailSecuritySummaryMaliciousParamsArc = "PASS"
	RadarEmailSecuritySummaryMaliciousParamsArcNone RadarEmailSecuritySummaryMaliciousParamsArc = "NONE"
	RadarEmailSecuritySummaryMaliciousParamsArcFail RadarEmailSecuritySummaryMaliciousParamsArc = "FAIL"
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

type RadarEmailSecuritySummaryMaliciousParamsDmarc string

const (
	RadarEmailSecuritySummaryMaliciousParamsDmarcPass RadarEmailSecuritySummaryMaliciousParamsDmarc = "PASS"
	RadarEmailSecuritySummaryMaliciousParamsDmarcNone RadarEmailSecuritySummaryMaliciousParamsDmarc = "NONE"
	RadarEmailSecuritySummaryMaliciousParamsDmarcFail RadarEmailSecuritySummaryMaliciousParamsDmarc = "FAIL"
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

type RadarEmailSecuritySummarySpamParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySummarySpamParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
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
	Dmarc param.Field[[]RadarEmailSecuritySummarySpamParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummarySpamParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummarySpamParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecuritySummarySpamParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecuritySummarySpamParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummarySpamParamsArc string

const (
	RadarEmailSecuritySummarySpamParamsArcPass RadarEmailSecuritySummarySpamParamsArc = "PASS"
	RadarEmailSecuritySummarySpamParamsArcNone RadarEmailSecuritySummarySpamParamsArc = "NONE"
	RadarEmailSecuritySummarySpamParamsArcFail RadarEmailSecuritySummarySpamParamsArc = "FAIL"
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

type RadarEmailSecuritySummarySpamParamsDmarc string

const (
	RadarEmailSecuritySummarySpamParamsDmarcPass RadarEmailSecuritySummarySpamParamsDmarc = "PASS"
	RadarEmailSecuritySummarySpamParamsDmarcNone RadarEmailSecuritySummarySpamParamsDmarc = "NONE"
	RadarEmailSecuritySummarySpamParamsDmarcFail RadarEmailSecuritySummarySpamParamsDmarc = "FAIL"
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

type RadarEmailSecuritySummarySPFParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySummarySPFParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
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
	Dmarc param.Field[[]RadarEmailSecuritySummarySPFParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummarySPFParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailSecuritySummarySPFParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecuritySummarySPFParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummarySPFParamsArc string

const (
	RadarEmailSecuritySummarySPFParamsArcPass RadarEmailSecuritySummarySPFParamsArc = "PASS"
	RadarEmailSecuritySummarySPFParamsArcNone RadarEmailSecuritySummarySPFParamsArc = "NONE"
	RadarEmailSecuritySummarySPFParamsArcFail RadarEmailSecuritySummarySPFParamsArc = "FAIL"
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

type RadarEmailSecuritySummarySPFParamsDmarc string

const (
	RadarEmailSecuritySummarySPFParamsDmarcPass RadarEmailSecuritySummarySPFParamsDmarc = "PASS"
	RadarEmailSecuritySummarySPFParamsDmarcNone RadarEmailSecuritySummarySPFParamsDmarc = "NONE"
	RadarEmailSecuritySummarySPFParamsDmarcFail RadarEmailSecuritySummarySPFParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummarySPFParamsFormat string

const (
	RadarEmailSecuritySummarySPFParamsFormatJson RadarEmailSecuritySummarySPFParamsFormat = "JSON"
	RadarEmailSecuritySummarySPFParamsFormatCsv  RadarEmailSecuritySummarySPFParamsFormat = "CSV"
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

type RadarEmailSecuritySummaryThreatCategoryParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySummaryThreatCategoryParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
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
	Dmarc param.Field[[]RadarEmailSecuritySummaryThreatCategoryParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummaryThreatCategoryParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummaryThreatCategoryParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecuritySummaryThreatCategoryParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecuritySummaryThreatCategoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryThreatCategoryParamsArc string

const (
	RadarEmailSecuritySummaryThreatCategoryParamsArcPass RadarEmailSecuritySummaryThreatCategoryParamsArc = "PASS"
	RadarEmailSecuritySummaryThreatCategoryParamsArcNone RadarEmailSecuritySummaryThreatCategoryParamsArc = "NONE"
	RadarEmailSecuritySummaryThreatCategoryParamsArcFail RadarEmailSecuritySummaryThreatCategoryParamsArc = "FAIL"
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

type RadarEmailSecuritySummaryThreatCategoryParamsDmarc string

const (
	RadarEmailSecuritySummaryThreatCategoryParamsDmarcPass RadarEmailSecuritySummaryThreatCategoryParamsDmarc = "PASS"
	RadarEmailSecuritySummaryThreatCategoryParamsDmarcNone RadarEmailSecuritySummaryThreatCategoryParamsDmarc = "NONE"
	RadarEmailSecuritySummaryThreatCategoryParamsDmarcFail RadarEmailSecuritySummaryThreatCategoryParamsDmarc = "FAIL"
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
