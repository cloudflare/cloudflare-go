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

// RadarEmailRoutingSummaryService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailRoutingSummaryService] method instead.
type RadarEmailRoutingSummaryService struct {
	Options []option.RequestOption
}

// NewRadarEmailRoutingSummaryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailRoutingSummaryService(opts ...option.RequestOption) (r *RadarEmailRoutingSummaryService) {
	r = &RadarEmailRoutingSummaryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per ARC validation.
func (r *RadarEmailRoutingSummaryService) ARC(ctx context.Context, query RadarEmailRoutingSummaryARCParams, opts ...option.RequestOption) (res *RadarEmailRoutingSummaryARCResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailRoutingSummaryARCResponseEnvelope
	path := "radar/email/routing/summary/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per DKIM validation.
func (r *RadarEmailRoutingSummaryService) DKIM(ctx context.Context, query RadarEmailRoutingSummaryDKIMParams, opts ...option.RequestOption) (res *RadarEmailRoutingSummaryDKIMResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailRoutingSummaryDKIMResponseEnvelope
	path := "radar/email/routing/summary/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per DMARC validation.
func (r *RadarEmailRoutingSummaryService) DMARC(ctx context.Context, query RadarEmailRoutingSummaryDMARCParams, opts ...option.RequestOption) (res *RadarEmailRoutingSummaryDMARCResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailRoutingSummaryDMARCResponseEnvelope
	path := "radar/email/routing/summary/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails by Encrypted
func (r *RadarEmailRoutingSummaryService) Encrypted(ctx context.Context, query RadarEmailRoutingSummaryEncryptedParams, opts ...option.RequestOption) (res *RadarEmailRoutingSummaryEncryptedResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailRoutingSummaryEncryptedResponseEnvelope
	path := "radar/email/routing/summary/encrypted"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails by Ip Version.
func (r *RadarEmailRoutingSummaryService) IPVersion(ctx context.Context, query RadarEmailRoutingSummaryIPVersionParams, opts ...option.RequestOption) (res *RadarEmailRoutingSummaryIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailRoutingSummaryIPVersionResponseEnvelope
	path := "radar/email/routing/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified per SPF validation.
func (r *RadarEmailRoutingSummaryService) SPF(ctx context.Context, query RadarEmailRoutingSummarySPFParams, opts ...option.RequestOption) (res *RadarEmailRoutingSummarySPFResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailRoutingSummarySPFResponseEnvelope
	path := "radar/email/routing/summary/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailRoutingSummaryARCResponse struct {
	Meta     RadarEmailRoutingSummaryARCResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailRoutingSummaryARCResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailRoutingSummaryARCResponseJSON     `json:"-"`
}

// radarEmailRoutingSummaryARCResponseJSON contains the JSON metadata for the
// struct [RadarEmailRoutingSummaryARCResponse]
type radarEmailRoutingSummaryARCResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryARCResponseMeta struct {
	DateRange      []RadarEmailRoutingSummaryARCResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                `json:"lastUpdated,required"`
	Normalization  string                                                `json:"normalization,required"`
	ConfidenceInfo RadarEmailRoutingSummaryARCResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailRoutingSummaryARCResponseMetaJSON           `json:"-"`
}

// radarEmailRoutingSummaryARCResponseMetaJSON contains the JSON metadata for the
// struct [RadarEmailRoutingSummaryARCResponseMeta]
type radarEmailRoutingSummaryARCResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryARCResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryARCResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      radarEmailRoutingSummaryARCResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailRoutingSummaryARCResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarEmailRoutingSummaryARCResponseMetaDateRange]
type radarEmailRoutingSummaryARCResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryARCResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryARCResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailRoutingSummaryARCResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                             `json:"level"`
	JSON        radarEmailRoutingSummaryARCResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailRoutingSummaryARCResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarEmailRoutingSummaryARCResponseMetaConfidenceInfo]
type radarEmailRoutingSummaryARCResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryARCResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryARCResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                              `json:"dataSource,required"`
	Description     string                                                              `json:"description,required"`
	EventType       string                                                              `json:"eventType,required"`
	IsInstantaneous interface{}                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                           `json:"startTime" format:"date-time"`
	JSON            radarEmailRoutingSummaryARCResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailRoutingSummaryARCResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingSummaryARCResponseMetaConfidenceInfoAnnotation]
type radarEmailRoutingSummaryARCResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailRoutingSummaryARCResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryARCResponseSummary0 struct {
	Fail string                                          `json:"FAIL,required"`
	None string                                          `json:"NONE,required"`
	Pass string                                          `json:"PASS,required"`
	JSON radarEmailRoutingSummaryARCResponseSummary0JSON `json:"-"`
}

// radarEmailRoutingSummaryARCResponseSummary0JSON contains the JSON metadata for
// the struct [RadarEmailRoutingSummaryARCResponseSummary0]
type radarEmailRoutingSummaryARCResponseSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryARCResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryDKIMResponse struct {
	Meta     RadarEmailRoutingSummaryDKIMResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailRoutingSummaryDKIMResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailRoutingSummaryDKIMResponseJSON     `json:"-"`
}

// radarEmailRoutingSummaryDKIMResponseJSON contains the JSON metadata for the
// struct [RadarEmailRoutingSummaryDKIMResponse]
type radarEmailRoutingSummaryDKIMResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryDKIMResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryDKIMResponseMeta struct {
	DateRange      []RadarEmailRoutingSummaryDKIMResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                 `json:"lastUpdated,required"`
	Normalization  string                                                 `json:"normalization,required"`
	ConfidenceInfo RadarEmailRoutingSummaryDKIMResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailRoutingSummaryDKIMResponseMetaJSON           `json:"-"`
}

// radarEmailRoutingSummaryDKIMResponseMetaJSON contains the JSON metadata for the
// struct [RadarEmailRoutingSummaryDKIMResponseMeta]
type radarEmailRoutingSummaryDKIMResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryDKIMResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryDKIMResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      radarEmailRoutingSummaryDKIMResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailRoutingSummaryDKIMResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarEmailRoutingSummaryDKIMResponseMetaDateRange]
type radarEmailRoutingSummaryDKIMResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryDKIMResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryDKIMResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        radarEmailRoutingSummaryDKIMResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailRoutingSummaryDKIMResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarEmailRoutingSummaryDKIMResponseMetaConfidenceInfo]
type radarEmailRoutingSummaryDKIMResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryDKIMResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            radarEmailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotation]
type radarEmailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryDKIMResponseSummary0 struct {
	Fail string                                           `json:"FAIL,required"`
	None string                                           `json:"NONE,required"`
	Pass string                                           `json:"PASS,required"`
	JSON radarEmailRoutingSummaryDKIMResponseSummary0JSON `json:"-"`
}

// radarEmailRoutingSummaryDKIMResponseSummary0JSON contains the JSON metadata for
// the struct [RadarEmailRoutingSummaryDKIMResponseSummary0]
type radarEmailRoutingSummaryDKIMResponseSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryDKIMResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryDMARCResponse struct {
	Meta     RadarEmailRoutingSummaryDMARCResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailRoutingSummaryDMARCResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailRoutingSummaryDMARCResponseJSON     `json:"-"`
}

// radarEmailRoutingSummaryDMARCResponseJSON contains the JSON metadata for the
// struct [RadarEmailRoutingSummaryDMARCResponse]
type radarEmailRoutingSummaryDMARCResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryDMARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryDMARCResponseMeta struct {
	DateRange      []RadarEmailRoutingSummaryDMARCResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                  `json:"lastUpdated,required"`
	Normalization  string                                                  `json:"normalization,required"`
	ConfidenceInfo RadarEmailRoutingSummaryDMARCResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailRoutingSummaryDMARCResponseMetaJSON           `json:"-"`
}

// radarEmailRoutingSummaryDMARCResponseMetaJSON contains the JSON metadata for the
// struct [RadarEmailRoutingSummaryDMARCResponseMeta]
type radarEmailRoutingSummaryDMARCResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryDMARCResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryDMARCResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      radarEmailRoutingSummaryDMARCResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailRoutingSummaryDMARCResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailRoutingSummaryDMARCResponseMetaDateRange]
type radarEmailRoutingSummaryDMARCResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryDMARCResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryDMARCResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                               `json:"level"`
	JSON        radarEmailRoutingSummaryDMARCResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailRoutingSummaryDMARCResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingSummaryDMARCResponseMetaConfidenceInfo]
type radarEmailRoutingSummaryDMARCResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryDMARCResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                `json:"dataSource,required"`
	Description     string                                                                `json:"description,required"`
	EventType       string                                                                `json:"eventType,required"`
	IsInstantaneous interface{}                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                             `json:"startTime" format:"date-time"`
	JSON            radarEmailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotation]
type radarEmailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryDMARCResponseSummary0 struct {
	Fail string                                            `json:"FAIL,required"`
	None string                                            `json:"NONE,required"`
	Pass string                                            `json:"PASS,required"`
	JSON radarEmailRoutingSummaryDMARCResponseSummary0JSON `json:"-"`
}

// radarEmailRoutingSummaryDMARCResponseSummary0JSON contains the JSON metadata for
// the struct [RadarEmailRoutingSummaryDMARCResponseSummary0]
type radarEmailRoutingSummaryDMARCResponseSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryDMARCResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryEncryptedResponse struct {
	Meta     RadarEmailRoutingSummaryEncryptedResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailRoutingSummaryEncryptedResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailRoutingSummaryEncryptedResponseJSON     `json:"-"`
}

// radarEmailRoutingSummaryEncryptedResponseJSON contains the JSON metadata for the
// struct [RadarEmailRoutingSummaryEncryptedResponse]
type radarEmailRoutingSummaryEncryptedResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryEncryptedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryEncryptedResponseMeta struct {
	DateRange      []RadarEmailRoutingSummaryEncryptedResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                      `json:"lastUpdated,required"`
	Normalization  string                                                      `json:"normalization,required"`
	ConfidenceInfo RadarEmailRoutingSummaryEncryptedResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailRoutingSummaryEncryptedResponseMetaJSON           `json:"-"`
}

// radarEmailRoutingSummaryEncryptedResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailRoutingSummaryEncryptedResponseMeta]
type radarEmailRoutingSummaryEncryptedResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryEncryptedResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryEncryptedResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarEmailRoutingSummaryEncryptedResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailRoutingSummaryEncryptedResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailRoutingSummaryEncryptedResponseMetaDateRange]
type radarEmailRoutingSummaryEncryptedResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryEncryptedResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryEncryptedResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                   `json:"level"`
	JSON        radarEmailRoutingSummaryEncryptedResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailRoutingSummaryEncryptedResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingSummaryEncryptedResponseMetaConfidenceInfo]
type radarEmailRoutingSummaryEncryptedResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryEncryptedResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                    `json:"dataSource,required"`
	Description     string                                                                    `json:"description,required"`
	EventType       string                                                                    `json:"eventType,required"`
	IsInstantaneous interface{}                                                               `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                                 `json:"startTime" format:"date-time"`
	JSON            radarEmailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotation]
type radarEmailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryEncryptedResponseSummary0 struct {
	Encrypted    string                                                `json:"ENCRYPTED,required"`
	NotEncrypted string                                                `json:"NOT_ENCRYPTED,required"`
	JSON         radarEmailRoutingSummaryEncryptedResponseSummary0JSON `json:"-"`
}

// radarEmailRoutingSummaryEncryptedResponseSummary0JSON contains the JSON metadata
// for the struct [RadarEmailRoutingSummaryEncryptedResponseSummary0]
type radarEmailRoutingSummaryEncryptedResponseSummary0JSON struct {
	Encrypted    apijson.Field
	NotEncrypted apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryEncryptedResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryIPVersionResponse struct {
	Meta     RadarEmailRoutingSummaryIPVersionResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailRoutingSummaryIPVersionResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailRoutingSummaryIPVersionResponseJSON     `json:"-"`
}

// radarEmailRoutingSummaryIPVersionResponseJSON contains the JSON metadata for the
// struct [RadarEmailRoutingSummaryIPVersionResponse]
type radarEmailRoutingSummaryIPVersionResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryIPVersionResponseMeta struct {
	DateRange      []RadarEmailRoutingSummaryIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                      `json:"lastUpdated,required"`
	Normalization  string                                                      `json:"normalization,required"`
	ConfidenceInfo RadarEmailRoutingSummaryIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailRoutingSummaryIPVersionResponseMetaJSON           `json:"-"`
}

// radarEmailRoutingSummaryIPVersionResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailRoutingSummaryIPVersionResponseMeta]
type radarEmailRoutingSummaryIPVersionResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarEmailRoutingSummaryIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailRoutingSummaryIPVersionResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailRoutingSummaryIPVersionResponseMetaDateRange]
type radarEmailRoutingSummaryIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryIPVersionResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                   `json:"level"`
	JSON        radarEmailRoutingSummaryIPVersionResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailRoutingSummaryIPVersionResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingSummaryIPVersionResponseMetaConfidenceInfo]
type radarEmailRoutingSummaryIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                    `json:"dataSource,required"`
	Description     string                                                                    `json:"description,required"`
	EventType       string                                                                    `json:"eventType,required"`
	IsInstantaneous interface{}                                                               `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                                 `json:"startTime" format:"date-time"`
	JSON            radarEmailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotation]
type radarEmailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryIPVersionResponseSummary0 struct {
	IPv4 string                                                `json:"IPv4,required"`
	IPv6 string                                                `json:"IPv6,required"`
	JSON radarEmailRoutingSummaryIPVersionResponseSummary0JSON `json:"-"`
}

// radarEmailRoutingSummaryIPVersionResponseSummary0JSON contains the JSON metadata
// for the struct [RadarEmailRoutingSummaryIPVersionResponseSummary0]
type radarEmailRoutingSummaryIPVersionResponseSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryIPVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummarySPFResponse struct {
	Meta     RadarEmailRoutingSummarySPFResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailRoutingSummarySPFResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailRoutingSummarySPFResponseJSON     `json:"-"`
}

// radarEmailRoutingSummarySPFResponseJSON contains the JSON metadata for the
// struct [RadarEmailRoutingSummarySPFResponse]
type radarEmailRoutingSummarySPFResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummarySPFResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummarySPFResponseMeta struct {
	DateRange      []RadarEmailRoutingSummarySPFResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                `json:"lastUpdated,required"`
	Normalization  string                                                `json:"normalization,required"`
	ConfidenceInfo RadarEmailRoutingSummarySPFResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailRoutingSummarySPFResponseMetaJSON           `json:"-"`
}

// radarEmailRoutingSummarySPFResponseMetaJSON contains the JSON metadata for the
// struct [RadarEmailRoutingSummarySPFResponseMeta]
type radarEmailRoutingSummarySPFResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailRoutingSummarySPFResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummarySPFResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      radarEmailRoutingSummarySPFResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailRoutingSummarySPFResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarEmailRoutingSummarySPFResponseMetaDateRange]
type radarEmailRoutingSummarySPFResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummarySPFResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummarySPFResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailRoutingSummarySPFResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                             `json:"level"`
	JSON        radarEmailRoutingSummarySPFResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailRoutingSummarySPFResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarEmailRoutingSummarySPFResponseMetaConfidenceInfo]
type radarEmailRoutingSummarySPFResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummarySPFResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummarySPFResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                              `json:"dataSource,required"`
	Description     string                                                              `json:"description,required"`
	EventType       string                                                              `json:"eventType,required"`
	IsInstantaneous interface{}                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                           `json:"startTime" format:"date-time"`
	JSON            radarEmailRoutingSummarySPFResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailRoutingSummarySPFResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingSummarySPFResponseMetaConfidenceInfoAnnotation]
type radarEmailRoutingSummarySPFResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailRoutingSummarySPFResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummarySPFResponseSummary0 struct {
	Fail string                                          `json:"FAIL,required"`
	None string                                          `json:"NONE,required"`
	Pass string                                          `json:"PASS,required"`
	JSON radarEmailRoutingSummarySPFResponseSummary0JSON `json:"-"`
}

// radarEmailRoutingSummarySPFResponseSummary0JSON contains the JSON metadata for
// the struct [RadarEmailRoutingSummarySPFResponseSummary0]
type radarEmailRoutingSummarySPFResponseSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummarySPFResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryARCParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailRoutingSummaryARCParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailRoutingSummaryARCParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailRoutingSummaryARCParamsDMARC] `query:"dmarc"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]RadarEmailRoutingSummaryARCParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[RadarEmailRoutingSummaryARCParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarEmailRoutingSummaryARCParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailRoutingSummaryARCParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingSummaryARCParams]'s query parameters as
// `url.Values`.
func (r RadarEmailRoutingSummaryARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailRoutingSummaryARCParamsDateRange string

const (
	RadarEmailRoutingSummaryARCParamsDateRange1d         RadarEmailRoutingSummaryARCParamsDateRange = "1d"
	RadarEmailRoutingSummaryARCParamsDateRange2d         RadarEmailRoutingSummaryARCParamsDateRange = "2d"
	RadarEmailRoutingSummaryARCParamsDateRange7d         RadarEmailRoutingSummaryARCParamsDateRange = "7d"
	RadarEmailRoutingSummaryARCParamsDateRange14d        RadarEmailRoutingSummaryARCParamsDateRange = "14d"
	RadarEmailRoutingSummaryARCParamsDateRange28d        RadarEmailRoutingSummaryARCParamsDateRange = "28d"
	RadarEmailRoutingSummaryARCParamsDateRange12w        RadarEmailRoutingSummaryARCParamsDateRange = "12w"
	RadarEmailRoutingSummaryARCParamsDateRange24w        RadarEmailRoutingSummaryARCParamsDateRange = "24w"
	RadarEmailRoutingSummaryARCParamsDateRange52w        RadarEmailRoutingSummaryARCParamsDateRange = "52w"
	RadarEmailRoutingSummaryARCParamsDateRange1dControl  RadarEmailRoutingSummaryARCParamsDateRange = "1dControl"
	RadarEmailRoutingSummaryARCParamsDateRange2dControl  RadarEmailRoutingSummaryARCParamsDateRange = "2dControl"
	RadarEmailRoutingSummaryARCParamsDateRange7dControl  RadarEmailRoutingSummaryARCParamsDateRange = "7dControl"
	RadarEmailRoutingSummaryARCParamsDateRange14dControl RadarEmailRoutingSummaryARCParamsDateRange = "14dControl"
	RadarEmailRoutingSummaryARCParamsDateRange28dControl RadarEmailRoutingSummaryARCParamsDateRange = "28dControl"
	RadarEmailRoutingSummaryARCParamsDateRange12wControl RadarEmailRoutingSummaryARCParamsDateRange = "12wControl"
	RadarEmailRoutingSummaryARCParamsDateRange24wControl RadarEmailRoutingSummaryARCParamsDateRange = "24wControl"
)

type RadarEmailRoutingSummaryARCParamsDKIM string

const (
	RadarEmailRoutingSummaryARCParamsDKIMPass RadarEmailRoutingSummaryARCParamsDKIM = "PASS"
	RadarEmailRoutingSummaryARCParamsDKIMNone RadarEmailRoutingSummaryARCParamsDKIM = "NONE"
	RadarEmailRoutingSummaryARCParamsDKIMFail RadarEmailRoutingSummaryARCParamsDKIM = "FAIL"
)

type RadarEmailRoutingSummaryARCParamsDMARC string

const (
	RadarEmailRoutingSummaryARCParamsDMARCPass RadarEmailRoutingSummaryARCParamsDMARC = "PASS"
	RadarEmailRoutingSummaryARCParamsDMARCNone RadarEmailRoutingSummaryARCParamsDMARC = "NONE"
	RadarEmailRoutingSummaryARCParamsDMARCFail RadarEmailRoutingSummaryARCParamsDMARC = "FAIL"
)

type RadarEmailRoutingSummaryARCParamsEncrypted string

const (
	RadarEmailRoutingSummaryARCParamsEncryptedEncrypted    RadarEmailRoutingSummaryARCParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingSummaryARCParamsEncryptedNotEncrypted RadarEmailRoutingSummaryARCParamsEncrypted = "NOT_ENCRYPTED"
)

// Format results are returned in.
type RadarEmailRoutingSummaryARCParamsFormat string

const (
	RadarEmailRoutingSummaryARCParamsFormatJson RadarEmailRoutingSummaryARCParamsFormat = "JSON"
	RadarEmailRoutingSummaryARCParamsFormatCsv  RadarEmailRoutingSummaryARCParamsFormat = "CSV"
)

type RadarEmailRoutingSummaryARCParamsIPVersion string

const (
	RadarEmailRoutingSummaryARCParamsIPVersionIPv4 RadarEmailRoutingSummaryARCParamsIPVersion = "IPv4"
	RadarEmailRoutingSummaryARCParamsIPVersionIPv6 RadarEmailRoutingSummaryARCParamsIPVersion = "IPv6"
)

type RadarEmailRoutingSummaryARCParamsSPF string

const (
	RadarEmailRoutingSummaryARCParamsSPFPass RadarEmailRoutingSummaryARCParamsSPF = "PASS"
	RadarEmailRoutingSummaryARCParamsSPFNone RadarEmailRoutingSummaryARCParamsSPF = "NONE"
	RadarEmailRoutingSummaryARCParamsSPFFail RadarEmailRoutingSummaryARCParamsSPF = "FAIL"
)

type RadarEmailRoutingSummaryARCResponseEnvelope struct {
	Result  RadarEmailRoutingSummaryARCResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarEmailRoutingSummaryARCResponseEnvelopeJSON `json:"-"`
}

// radarEmailRoutingSummaryARCResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarEmailRoutingSummaryARCResponseEnvelope]
type radarEmailRoutingSummaryARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryDKIMParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailRoutingSummaryDKIMParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailRoutingSummaryDKIMParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailRoutingSummaryDKIMParamsDMARC] `query:"dmarc"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]RadarEmailRoutingSummaryDKIMParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[RadarEmailRoutingSummaryDKIMParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarEmailRoutingSummaryDKIMParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailRoutingSummaryDKIMParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingSummaryDKIMParams]'s query parameters as
// `url.Values`.
func (r RadarEmailRoutingSummaryDKIMParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailRoutingSummaryDKIMParamsARC string

const (
	RadarEmailRoutingSummaryDKIMParamsARCPass RadarEmailRoutingSummaryDKIMParamsARC = "PASS"
	RadarEmailRoutingSummaryDKIMParamsARCNone RadarEmailRoutingSummaryDKIMParamsARC = "NONE"
	RadarEmailRoutingSummaryDKIMParamsARCFail RadarEmailRoutingSummaryDKIMParamsARC = "FAIL"
)

type RadarEmailRoutingSummaryDKIMParamsDateRange string

const (
	RadarEmailRoutingSummaryDKIMParamsDateRange1d         RadarEmailRoutingSummaryDKIMParamsDateRange = "1d"
	RadarEmailRoutingSummaryDKIMParamsDateRange2d         RadarEmailRoutingSummaryDKIMParamsDateRange = "2d"
	RadarEmailRoutingSummaryDKIMParamsDateRange7d         RadarEmailRoutingSummaryDKIMParamsDateRange = "7d"
	RadarEmailRoutingSummaryDKIMParamsDateRange14d        RadarEmailRoutingSummaryDKIMParamsDateRange = "14d"
	RadarEmailRoutingSummaryDKIMParamsDateRange28d        RadarEmailRoutingSummaryDKIMParamsDateRange = "28d"
	RadarEmailRoutingSummaryDKIMParamsDateRange12w        RadarEmailRoutingSummaryDKIMParamsDateRange = "12w"
	RadarEmailRoutingSummaryDKIMParamsDateRange24w        RadarEmailRoutingSummaryDKIMParamsDateRange = "24w"
	RadarEmailRoutingSummaryDKIMParamsDateRange52w        RadarEmailRoutingSummaryDKIMParamsDateRange = "52w"
	RadarEmailRoutingSummaryDKIMParamsDateRange1dControl  RadarEmailRoutingSummaryDKIMParamsDateRange = "1dControl"
	RadarEmailRoutingSummaryDKIMParamsDateRange2dControl  RadarEmailRoutingSummaryDKIMParamsDateRange = "2dControl"
	RadarEmailRoutingSummaryDKIMParamsDateRange7dControl  RadarEmailRoutingSummaryDKIMParamsDateRange = "7dControl"
	RadarEmailRoutingSummaryDKIMParamsDateRange14dControl RadarEmailRoutingSummaryDKIMParamsDateRange = "14dControl"
	RadarEmailRoutingSummaryDKIMParamsDateRange28dControl RadarEmailRoutingSummaryDKIMParamsDateRange = "28dControl"
	RadarEmailRoutingSummaryDKIMParamsDateRange12wControl RadarEmailRoutingSummaryDKIMParamsDateRange = "12wControl"
	RadarEmailRoutingSummaryDKIMParamsDateRange24wControl RadarEmailRoutingSummaryDKIMParamsDateRange = "24wControl"
)

type RadarEmailRoutingSummaryDKIMParamsDMARC string

const (
	RadarEmailRoutingSummaryDKIMParamsDMARCPass RadarEmailRoutingSummaryDKIMParamsDMARC = "PASS"
	RadarEmailRoutingSummaryDKIMParamsDMARCNone RadarEmailRoutingSummaryDKIMParamsDMARC = "NONE"
	RadarEmailRoutingSummaryDKIMParamsDMARCFail RadarEmailRoutingSummaryDKIMParamsDMARC = "FAIL"
)

type RadarEmailRoutingSummaryDKIMParamsEncrypted string

const (
	RadarEmailRoutingSummaryDKIMParamsEncryptedEncrypted    RadarEmailRoutingSummaryDKIMParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingSummaryDKIMParamsEncryptedNotEncrypted RadarEmailRoutingSummaryDKIMParamsEncrypted = "NOT_ENCRYPTED"
)

// Format results are returned in.
type RadarEmailRoutingSummaryDKIMParamsFormat string

const (
	RadarEmailRoutingSummaryDKIMParamsFormatJson RadarEmailRoutingSummaryDKIMParamsFormat = "JSON"
	RadarEmailRoutingSummaryDKIMParamsFormatCsv  RadarEmailRoutingSummaryDKIMParamsFormat = "CSV"
)

type RadarEmailRoutingSummaryDKIMParamsIPVersion string

const (
	RadarEmailRoutingSummaryDKIMParamsIPVersionIPv4 RadarEmailRoutingSummaryDKIMParamsIPVersion = "IPv4"
	RadarEmailRoutingSummaryDKIMParamsIPVersionIPv6 RadarEmailRoutingSummaryDKIMParamsIPVersion = "IPv6"
)

type RadarEmailRoutingSummaryDKIMParamsSPF string

const (
	RadarEmailRoutingSummaryDKIMParamsSPFPass RadarEmailRoutingSummaryDKIMParamsSPF = "PASS"
	RadarEmailRoutingSummaryDKIMParamsSPFNone RadarEmailRoutingSummaryDKIMParamsSPF = "NONE"
	RadarEmailRoutingSummaryDKIMParamsSPFFail RadarEmailRoutingSummaryDKIMParamsSPF = "FAIL"
)

type RadarEmailRoutingSummaryDKIMResponseEnvelope struct {
	Result  RadarEmailRoutingSummaryDKIMResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarEmailRoutingSummaryDKIMResponseEnvelopeJSON `json:"-"`
}

// radarEmailRoutingSummaryDKIMResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarEmailRoutingSummaryDKIMResponseEnvelope]
type radarEmailRoutingSummaryDKIMResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryDKIMResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryDMARCParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailRoutingSummaryDMARCParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailRoutingSummaryDMARCParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailRoutingSummaryDMARCParamsDKIM] `query:"dkim"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]RadarEmailRoutingSummaryDMARCParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[RadarEmailRoutingSummaryDMARCParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarEmailRoutingSummaryDMARCParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailRoutingSummaryDMARCParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingSummaryDMARCParams]'s query parameters as
// `url.Values`.
func (r RadarEmailRoutingSummaryDMARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailRoutingSummaryDMARCParamsARC string

const (
	RadarEmailRoutingSummaryDMARCParamsARCPass RadarEmailRoutingSummaryDMARCParamsARC = "PASS"
	RadarEmailRoutingSummaryDMARCParamsARCNone RadarEmailRoutingSummaryDMARCParamsARC = "NONE"
	RadarEmailRoutingSummaryDMARCParamsARCFail RadarEmailRoutingSummaryDMARCParamsARC = "FAIL"
)

type RadarEmailRoutingSummaryDMARCParamsDateRange string

const (
	RadarEmailRoutingSummaryDMARCParamsDateRange1d         RadarEmailRoutingSummaryDMARCParamsDateRange = "1d"
	RadarEmailRoutingSummaryDMARCParamsDateRange2d         RadarEmailRoutingSummaryDMARCParamsDateRange = "2d"
	RadarEmailRoutingSummaryDMARCParamsDateRange7d         RadarEmailRoutingSummaryDMARCParamsDateRange = "7d"
	RadarEmailRoutingSummaryDMARCParamsDateRange14d        RadarEmailRoutingSummaryDMARCParamsDateRange = "14d"
	RadarEmailRoutingSummaryDMARCParamsDateRange28d        RadarEmailRoutingSummaryDMARCParamsDateRange = "28d"
	RadarEmailRoutingSummaryDMARCParamsDateRange12w        RadarEmailRoutingSummaryDMARCParamsDateRange = "12w"
	RadarEmailRoutingSummaryDMARCParamsDateRange24w        RadarEmailRoutingSummaryDMARCParamsDateRange = "24w"
	RadarEmailRoutingSummaryDMARCParamsDateRange52w        RadarEmailRoutingSummaryDMARCParamsDateRange = "52w"
	RadarEmailRoutingSummaryDMARCParamsDateRange1dControl  RadarEmailRoutingSummaryDMARCParamsDateRange = "1dControl"
	RadarEmailRoutingSummaryDMARCParamsDateRange2dControl  RadarEmailRoutingSummaryDMARCParamsDateRange = "2dControl"
	RadarEmailRoutingSummaryDMARCParamsDateRange7dControl  RadarEmailRoutingSummaryDMARCParamsDateRange = "7dControl"
	RadarEmailRoutingSummaryDMARCParamsDateRange14dControl RadarEmailRoutingSummaryDMARCParamsDateRange = "14dControl"
	RadarEmailRoutingSummaryDMARCParamsDateRange28dControl RadarEmailRoutingSummaryDMARCParamsDateRange = "28dControl"
	RadarEmailRoutingSummaryDMARCParamsDateRange12wControl RadarEmailRoutingSummaryDMARCParamsDateRange = "12wControl"
	RadarEmailRoutingSummaryDMARCParamsDateRange24wControl RadarEmailRoutingSummaryDMARCParamsDateRange = "24wControl"
)

type RadarEmailRoutingSummaryDMARCParamsDKIM string

const (
	RadarEmailRoutingSummaryDMARCParamsDKIMPass RadarEmailRoutingSummaryDMARCParamsDKIM = "PASS"
	RadarEmailRoutingSummaryDMARCParamsDKIMNone RadarEmailRoutingSummaryDMARCParamsDKIM = "NONE"
	RadarEmailRoutingSummaryDMARCParamsDKIMFail RadarEmailRoutingSummaryDMARCParamsDKIM = "FAIL"
)

type RadarEmailRoutingSummaryDMARCParamsEncrypted string

const (
	RadarEmailRoutingSummaryDMARCParamsEncryptedEncrypted    RadarEmailRoutingSummaryDMARCParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingSummaryDMARCParamsEncryptedNotEncrypted RadarEmailRoutingSummaryDMARCParamsEncrypted = "NOT_ENCRYPTED"
)

// Format results are returned in.
type RadarEmailRoutingSummaryDMARCParamsFormat string

const (
	RadarEmailRoutingSummaryDMARCParamsFormatJson RadarEmailRoutingSummaryDMARCParamsFormat = "JSON"
	RadarEmailRoutingSummaryDMARCParamsFormatCsv  RadarEmailRoutingSummaryDMARCParamsFormat = "CSV"
)

type RadarEmailRoutingSummaryDMARCParamsIPVersion string

const (
	RadarEmailRoutingSummaryDMARCParamsIPVersionIPv4 RadarEmailRoutingSummaryDMARCParamsIPVersion = "IPv4"
	RadarEmailRoutingSummaryDMARCParamsIPVersionIPv6 RadarEmailRoutingSummaryDMARCParamsIPVersion = "IPv6"
)

type RadarEmailRoutingSummaryDMARCParamsSPF string

const (
	RadarEmailRoutingSummaryDMARCParamsSPFPass RadarEmailRoutingSummaryDMARCParamsSPF = "PASS"
	RadarEmailRoutingSummaryDMARCParamsSPFNone RadarEmailRoutingSummaryDMARCParamsSPF = "NONE"
	RadarEmailRoutingSummaryDMARCParamsSPFFail RadarEmailRoutingSummaryDMARCParamsSPF = "FAIL"
)

type RadarEmailRoutingSummaryDMARCResponseEnvelope struct {
	Result  RadarEmailRoutingSummaryDMARCResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarEmailRoutingSummaryDMARCResponseEnvelopeJSON `json:"-"`
}

// radarEmailRoutingSummaryDMARCResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarEmailRoutingSummaryDMARCResponseEnvelope]
type radarEmailRoutingSummaryDMARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryDMARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryEncryptedParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailRoutingSummaryEncryptedParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailRoutingSummaryEncryptedParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailRoutingSummaryEncryptedParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailRoutingSummaryEncryptedParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailRoutingSummaryEncryptedParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarEmailRoutingSummaryEncryptedParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailRoutingSummaryEncryptedParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingSummaryEncryptedParams]'s query parameters
// as `url.Values`.
func (r RadarEmailRoutingSummaryEncryptedParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailRoutingSummaryEncryptedParamsARC string

const (
	RadarEmailRoutingSummaryEncryptedParamsARCPass RadarEmailRoutingSummaryEncryptedParamsARC = "PASS"
	RadarEmailRoutingSummaryEncryptedParamsARCNone RadarEmailRoutingSummaryEncryptedParamsARC = "NONE"
	RadarEmailRoutingSummaryEncryptedParamsARCFail RadarEmailRoutingSummaryEncryptedParamsARC = "FAIL"
)

type RadarEmailRoutingSummaryEncryptedParamsDateRange string

const (
	RadarEmailRoutingSummaryEncryptedParamsDateRange1d         RadarEmailRoutingSummaryEncryptedParamsDateRange = "1d"
	RadarEmailRoutingSummaryEncryptedParamsDateRange2d         RadarEmailRoutingSummaryEncryptedParamsDateRange = "2d"
	RadarEmailRoutingSummaryEncryptedParamsDateRange7d         RadarEmailRoutingSummaryEncryptedParamsDateRange = "7d"
	RadarEmailRoutingSummaryEncryptedParamsDateRange14d        RadarEmailRoutingSummaryEncryptedParamsDateRange = "14d"
	RadarEmailRoutingSummaryEncryptedParamsDateRange28d        RadarEmailRoutingSummaryEncryptedParamsDateRange = "28d"
	RadarEmailRoutingSummaryEncryptedParamsDateRange12w        RadarEmailRoutingSummaryEncryptedParamsDateRange = "12w"
	RadarEmailRoutingSummaryEncryptedParamsDateRange24w        RadarEmailRoutingSummaryEncryptedParamsDateRange = "24w"
	RadarEmailRoutingSummaryEncryptedParamsDateRange52w        RadarEmailRoutingSummaryEncryptedParamsDateRange = "52w"
	RadarEmailRoutingSummaryEncryptedParamsDateRange1dControl  RadarEmailRoutingSummaryEncryptedParamsDateRange = "1dControl"
	RadarEmailRoutingSummaryEncryptedParamsDateRange2dControl  RadarEmailRoutingSummaryEncryptedParamsDateRange = "2dControl"
	RadarEmailRoutingSummaryEncryptedParamsDateRange7dControl  RadarEmailRoutingSummaryEncryptedParamsDateRange = "7dControl"
	RadarEmailRoutingSummaryEncryptedParamsDateRange14dControl RadarEmailRoutingSummaryEncryptedParamsDateRange = "14dControl"
	RadarEmailRoutingSummaryEncryptedParamsDateRange28dControl RadarEmailRoutingSummaryEncryptedParamsDateRange = "28dControl"
	RadarEmailRoutingSummaryEncryptedParamsDateRange12wControl RadarEmailRoutingSummaryEncryptedParamsDateRange = "12wControl"
	RadarEmailRoutingSummaryEncryptedParamsDateRange24wControl RadarEmailRoutingSummaryEncryptedParamsDateRange = "24wControl"
)

type RadarEmailRoutingSummaryEncryptedParamsDKIM string

const (
	RadarEmailRoutingSummaryEncryptedParamsDKIMPass RadarEmailRoutingSummaryEncryptedParamsDKIM = "PASS"
	RadarEmailRoutingSummaryEncryptedParamsDKIMNone RadarEmailRoutingSummaryEncryptedParamsDKIM = "NONE"
	RadarEmailRoutingSummaryEncryptedParamsDKIMFail RadarEmailRoutingSummaryEncryptedParamsDKIM = "FAIL"
)

type RadarEmailRoutingSummaryEncryptedParamsDMARC string

const (
	RadarEmailRoutingSummaryEncryptedParamsDMARCPass RadarEmailRoutingSummaryEncryptedParamsDMARC = "PASS"
	RadarEmailRoutingSummaryEncryptedParamsDMARCNone RadarEmailRoutingSummaryEncryptedParamsDMARC = "NONE"
	RadarEmailRoutingSummaryEncryptedParamsDMARCFail RadarEmailRoutingSummaryEncryptedParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailRoutingSummaryEncryptedParamsFormat string

const (
	RadarEmailRoutingSummaryEncryptedParamsFormatJson RadarEmailRoutingSummaryEncryptedParamsFormat = "JSON"
	RadarEmailRoutingSummaryEncryptedParamsFormatCsv  RadarEmailRoutingSummaryEncryptedParamsFormat = "CSV"
)

type RadarEmailRoutingSummaryEncryptedParamsIPVersion string

const (
	RadarEmailRoutingSummaryEncryptedParamsIPVersionIPv4 RadarEmailRoutingSummaryEncryptedParamsIPVersion = "IPv4"
	RadarEmailRoutingSummaryEncryptedParamsIPVersionIPv6 RadarEmailRoutingSummaryEncryptedParamsIPVersion = "IPv6"
)

type RadarEmailRoutingSummaryEncryptedParamsSPF string

const (
	RadarEmailRoutingSummaryEncryptedParamsSPFPass RadarEmailRoutingSummaryEncryptedParamsSPF = "PASS"
	RadarEmailRoutingSummaryEncryptedParamsSPFNone RadarEmailRoutingSummaryEncryptedParamsSPF = "NONE"
	RadarEmailRoutingSummaryEncryptedParamsSPFFail RadarEmailRoutingSummaryEncryptedParamsSPF = "FAIL"
)

type RadarEmailRoutingSummaryEncryptedResponseEnvelope struct {
	Result  RadarEmailRoutingSummaryEncryptedResponse             `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarEmailRoutingSummaryEncryptedResponseEnvelopeJSON `json:"-"`
}

// radarEmailRoutingSummaryEncryptedResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailRoutingSummaryEncryptedResponseEnvelope]
type radarEmailRoutingSummaryEncryptedResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryEncryptedResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummaryIPVersionParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailRoutingSummaryIPVersionParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailRoutingSummaryIPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailRoutingSummaryIPVersionParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailRoutingSummaryIPVersionParamsDMARC] `query:"dmarc"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]RadarEmailRoutingSummaryIPVersionParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[RadarEmailRoutingSummaryIPVersionParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailRoutingSummaryIPVersionParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingSummaryIPVersionParams]'s query parameters
// as `url.Values`.
func (r RadarEmailRoutingSummaryIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailRoutingSummaryIPVersionParamsARC string

const (
	RadarEmailRoutingSummaryIPVersionParamsARCPass RadarEmailRoutingSummaryIPVersionParamsARC = "PASS"
	RadarEmailRoutingSummaryIPVersionParamsARCNone RadarEmailRoutingSummaryIPVersionParamsARC = "NONE"
	RadarEmailRoutingSummaryIPVersionParamsARCFail RadarEmailRoutingSummaryIPVersionParamsARC = "FAIL"
)

type RadarEmailRoutingSummaryIPVersionParamsDateRange string

const (
	RadarEmailRoutingSummaryIPVersionParamsDateRange1d         RadarEmailRoutingSummaryIPVersionParamsDateRange = "1d"
	RadarEmailRoutingSummaryIPVersionParamsDateRange2d         RadarEmailRoutingSummaryIPVersionParamsDateRange = "2d"
	RadarEmailRoutingSummaryIPVersionParamsDateRange7d         RadarEmailRoutingSummaryIPVersionParamsDateRange = "7d"
	RadarEmailRoutingSummaryIPVersionParamsDateRange14d        RadarEmailRoutingSummaryIPVersionParamsDateRange = "14d"
	RadarEmailRoutingSummaryIPVersionParamsDateRange28d        RadarEmailRoutingSummaryIPVersionParamsDateRange = "28d"
	RadarEmailRoutingSummaryIPVersionParamsDateRange12w        RadarEmailRoutingSummaryIPVersionParamsDateRange = "12w"
	RadarEmailRoutingSummaryIPVersionParamsDateRange24w        RadarEmailRoutingSummaryIPVersionParamsDateRange = "24w"
	RadarEmailRoutingSummaryIPVersionParamsDateRange52w        RadarEmailRoutingSummaryIPVersionParamsDateRange = "52w"
	RadarEmailRoutingSummaryIPVersionParamsDateRange1dControl  RadarEmailRoutingSummaryIPVersionParamsDateRange = "1dControl"
	RadarEmailRoutingSummaryIPVersionParamsDateRange2dControl  RadarEmailRoutingSummaryIPVersionParamsDateRange = "2dControl"
	RadarEmailRoutingSummaryIPVersionParamsDateRange7dControl  RadarEmailRoutingSummaryIPVersionParamsDateRange = "7dControl"
	RadarEmailRoutingSummaryIPVersionParamsDateRange14dControl RadarEmailRoutingSummaryIPVersionParamsDateRange = "14dControl"
	RadarEmailRoutingSummaryIPVersionParamsDateRange28dControl RadarEmailRoutingSummaryIPVersionParamsDateRange = "28dControl"
	RadarEmailRoutingSummaryIPVersionParamsDateRange12wControl RadarEmailRoutingSummaryIPVersionParamsDateRange = "12wControl"
	RadarEmailRoutingSummaryIPVersionParamsDateRange24wControl RadarEmailRoutingSummaryIPVersionParamsDateRange = "24wControl"
)

type RadarEmailRoutingSummaryIPVersionParamsDKIM string

const (
	RadarEmailRoutingSummaryIPVersionParamsDKIMPass RadarEmailRoutingSummaryIPVersionParamsDKIM = "PASS"
	RadarEmailRoutingSummaryIPVersionParamsDKIMNone RadarEmailRoutingSummaryIPVersionParamsDKIM = "NONE"
	RadarEmailRoutingSummaryIPVersionParamsDKIMFail RadarEmailRoutingSummaryIPVersionParamsDKIM = "FAIL"
)

type RadarEmailRoutingSummaryIPVersionParamsDMARC string

const (
	RadarEmailRoutingSummaryIPVersionParamsDMARCPass RadarEmailRoutingSummaryIPVersionParamsDMARC = "PASS"
	RadarEmailRoutingSummaryIPVersionParamsDMARCNone RadarEmailRoutingSummaryIPVersionParamsDMARC = "NONE"
	RadarEmailRoutingSummaryIPVersionParamsDMARCFail RadarEmailRoutingSummaryIPVersionParamsDMARC = "FAIL"
)

type RadarEmailRoutingSummaryIPVersionParamsEncrypted string

const (
	RadarEmailRoutingSummaryIPVersionParamsEncryptedEncrypted    RadarEmailRoutingSummaryIPVersionParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingSummaryIPVersionParamsEncryptedNotEncrypted RadarEmailRoutingSummaryIPVersionParamsEncrypted = "NOT_ENCRYPTED"
)

// Format results are returned in.
type RadarEmailRoutingSummaryIPVersionParamsFormat string

const (
	RadarEmailRoutingSummaryIPVersionParamsFormatJson RadarEmailRoutingSummaryIPVersionParamsFormat = "JSON"
	RadarEmailRoutingSummaryIPVersionParamsFormatCsv  RadarEmailRoutingSummaryIPVersionParamsFormat = "CSV"
)

type RadarEmailRoutingSummaryIPVersionParamsSPF string

const (
	RadarEmailRoutingSummaryIPVersionParamsSPFPass RadarEmailRoutingSummaryIPVersionParamsSPF = "PASS"
	RadarEmailRoutingSummaryIPVersionParamsSPFNone RadarEmailRoutingSummaryIPVersionParamsSPF = "NONE"
	RadarEmailRoutingSummaryIPVersionParamsSPFFail RadarEmailRoutingSummaryIPVersionParamsSPF = "FAIL"
)

type RadarEmailRoutingSummaryIPVersionResponseEnvelope struct {
	Result  RadarEmailRoutingSummaryIPVersionResponse             `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarEmailRoutingSummaryIPVersionResponseEnvelopeJSON `json:"-"`
}

// radarEmailRoutingSummaryIPVersionResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailRoutingSummaryIPVersionResponseEnvelope]
type radarEmailRoutingSummaryIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailRoutingSummarySPFParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailRoutingSummarySPFParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailRoutingSummarySPFParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailRoutingSummarySPFParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailRoutingSummarySPFParamsDMARC] `query:"dmarc"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]RadarEmailRoutingSummarySPFParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[RadarEmailRoutingSummarySPFParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarEmailRoutingSummarySPFParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailRoutingSummarySPFParams]'s query parameters as
// `url.Values`.
func (r RadarEmailRoutingSummarySPFParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailRoutingSummarySPFParamsARC string

const (
	RadarEmailRoutingSummarySPFParamsARCPass RadarEmailRoutingSummarySPFParamsARC = "PASS"
	RadarEmailRoutingSummarySPFParamsARCNone RadarEmailRoutingSummarySPFParamsARC = "NONE"
	RadarEmailRoutingSummarySPFParamsARCFail RadarEmailRoutingSummarySPFParamsARC = "FAIL"
)

type RadarEmailRoutingSummarySPFParamsDateRange string

const (
	RadarEmailRoutingSummarySPFParamsDateRange1d         RadarEmailRoutingSummarySPFParamsDateRange = "1d"
	RadarEmailRoutingSummarySPFParamsDateRange2d         RadarEmailRoutingSummarySPFParamsDateRange = "2d"
	RadarEmailRoutingSummarySPFParamsDateRange7d         RadarEmailRoutingSummarySPFParamsDateRange = "7d"
	RadarEmailRoutingSummarySPFParamsDateRange14d        RadarEmailRoutingSummarySPFParamsDateRange = "14d"
	RadarEmailRoutingSummarySPFParamsDateRange28d        RadarEmailRoutingSummarySPFParamsDateRange = "28d"
	RadarEmailRoutingSummarySPFParamsDateRange12w        RadarEmailRoutingSummarySPFParamsDateRange = "12w"
	RadarEmailRoutingSummarySPFParamsDateRange24w        RadarEmailRoutingSummarySPFParamsDateRange = "24w"
	RadarEmailRoutingSummarySPFParamsDateRange52w        RadarEmailRoutingSummarySPFParamsDateRange = "52w"
	RadarEmailRoutingSummarySPFParamsDateRange1dControl  RadarEmailRoutingSummarySPFParamsDateRange = "1dControl"
	RadarEmailRoutingSummarySPFParamsDateRange2dControl  RadarEmailRoutingSummarySPFParamsDateRange = "2dControl"
	RadarEmailRoutingSummarySPFParamsDateRange7dControl  RadarEmailRoutingSummarySPFParamsDateRange = "7dControl"
	RadarEmailRoutingSummarySPFParamsDateRange14dControl RadarEmailRoutingSummarySPFParamsDateRange = "14dControl"
	RadarEmailRoutingSummarySPFParamsDateRange28dControl RadarEmailRoutingSummarySPFParamsDateRange = "28dControl"
	RadarEmailRoutingSummarySPFParamsDateRange12wControl RadarEmailRoutingSummarySPFParamsDateRange = "12wControl"
	RadarEmailRoutingSummarySPFParamsDateRange24wControl RadarEmailRoutingSummarySPFParamsDateRange = "24wControl"
)

type RadarEmailRoutingSummarySPFParamsDKIM string

const (
	RadarEmailRoutingSummarySPFParamsDKIMPass RadarEmailRoutingSummarySPFParamsDKIM = "PASS"
	RadarEmailRoutingSummarySPFParamsDKIMNone RadarEmailRoutingSummarySPFParamsDKIM = "NONE"
	RadarEmailRoutingSummarySPFParamsDKIMFail RadarEmailRoutingSummarySPFParamsDKIM = "FAIL"
)

type RadarEmailRoutingSummarySPFParamsDMARC string

const (
	RadarEmailRoutingSummarySPFParamsDMARCPass RadarEmailRoutingSummarySPFParamsDMARC = "PASS"
	RadarEmailRoutingSummarySPFParamsDMARCNone RadarEmailRoutingSummarySPFParamsDMARC = "NONE"
	RadarEmailRoutingSummarySPFParamsDMARCFail RadarEmailRoutingSummarySPFParamsDMARC = "FAIL"
)

type RadarEmailRoutingSummarySPFParamsEncrypted string

const (
	RadarEmailRoutingSummarySPFParamsEncryptedEncrypted    RadarEmailRoutingSummarySPFParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingSummarySPFParamsEncryptedNotEncrypted RadarEmailRoutingSummarySPFParamsEncrypted = "NOT_ENCRYPTED"
)

// Format results are returned in.
type RadarEmailRoutingSummarySPFParamsFormat string

const (
	RadarEmailRoutingSummarySPFParamsFormatJson RadarEmailRoutingSummarySPFParamsFormat = "JSON"
	RadarEmailRoutingSummarySPFParamsFormatCsv  RadarEmailRoutingSummarySPFParamsFormat = "CSV"
)

type RadarEmailRoutingSummarySPFParamsIPVersion string

const (
	RadarEmailRoutingSummarySPFParamsIPVersionIPv4 RadarEmailRoutingSummarySPFParamsIPVersion = "IPv4"
	RadarEmailRoutingSummarySPFParamsIPVersionIPv6 RadarEmailRoutingSummarySPFParamsIPVersion = "IPv6"
)

type RadarEmailRoutingSummarySPFResponseEnvelope struct {
	Result  RadarEmailRoutingSummarySPFResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarEmailRoutingSummarySPFResponseEnvelopeJSON `json:"-"`
}

// radarEmailRoutingSummarySPFResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarEmailRoutingSummarySPFResponseEnvelope]
type radarEmailRoutingSummarySPFResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummarySPFResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
