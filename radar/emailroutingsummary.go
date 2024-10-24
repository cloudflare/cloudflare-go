// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// EmailRoutingSummaryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailRoutingSummaryService] method instead.
type EmailRoutingSummaryService struct {
	Options []option.RequestOption
}

// NewEmailRoutingSummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailRoutingSummaryService(opts ...option.RequestOption) (r *EmailRoutingSummaryService) {
	r = &EmailRoutingSummaryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified by ARC validation.
func (r *EmailRoutingSummaryService) ARC(ctx context.Context, query EmailRoutingSummaryARCParams, opts ...option.RequestOption) (res *EmailRoutingSummaryARCResponse, err error) {
	var env EmailRoutingSummaryARCResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/summary/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified by DKIM validation.
func (r *EmailRoutingSummaryService) DKIM(ctx context.Context, query EmailRoutingSummaryDKIMParams, opts ...option.RequestOption) (res *EmailRoutingSummaryDKIMResponse, err error) {
	var env EmailRoutingSummaryDKIMResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/summary/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified by DMARC validation.
func (r *EmailRoutingSummaryService) DMARC(ctx context.Context, query EmailRoutingSummaryDMARCParams, opts ...option.RequestOption) (res *EmailRoutingSummaryDMARCResponse, err error) {
	var env EmailRoutingSummaryDMARCResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/summary/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails by encryption status.
func (r *EmailRoutingSummaryService) Encrypted(ctx context.Context, query EmailRoutingSummaryEncryptedParams, opts ...option.RequestOption) (res *EmailRoutingSummaryEncryptedResponse, err error) {
	var env EmailRoutingSummaryEncryptedResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/summary/encrypted"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails by IP version.
func (r *EmailRoutingSummaryService) IPVersion(ctx context.Context, query EmailRoutingSummaryIPVersionParams, opts ...option.RequestOption) (res *EmailRoutingSummaryIPVersionResponse, err error) {
	var env EmailRoutingSummaryIPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of emails classified by SPF validation.
func (r *EmailRoutingSummaryService) SPF(ctx context.Context, query EmailRoutingSummarySPFParams, opts ...option.RequestOption) (res *EmailRoutingSummarySPFResponse, err error) {
	var env EmailRoutingSummarySPFResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/summary/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingSummaryARCResponse struct {
	Meta     EmailRoutingSummaryARCResponseMeta `json:"meta,required"`
	Summary0 RadarEmailSummary                  `json:"summary_0,required"`
	JSON     emailRoutingSummaryARCResponseJSON `json:"-"`
}

// emailRoutingSummaryARCResponseJSON contains the JSON metadata for the struct
// [EmailRoutingSummaryARCResponse]
type emailRoutingSummaryARCResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryARCResponseJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryARCResponseMeta struct {
	DateRange      []EmailRoutingSummaryARCResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                           `json:"lastUpdated,required"`
	Normalization  string                                           `json:"normalization,required"`
	ConfidenceInfo EmailRoutingSummaryARCResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           emailRoutingSummaryARCResponseMetaJSON           `json:"-"`
}

// emailRoutingSummaryARCResponseMetaJSON contains the JSON metadata for the struct
// [EmailRoutingSummaryARCResponseMeta]
type emailRoutingSummaryARCResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingSummaryARCResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryARCResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryARCResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingSummaryARCResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingSummaryARCResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [EmailRoutingSummaryARCResponseMetaDateRange]
type emailRoutingSummaryARCResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryARCResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryARCResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryARCResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingSummaryARCResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                        `json:"level"`
	JSON        emailRoutingSummaryARCResponseMetaConfidenceInfoJSON         `json:"-"`
}

// emailRoutingSummaryARCResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [EmailRoutingSummaryARCResponseMetaConfidenceInfo]
type emailRoutingSummaryARCResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryARCResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryARCResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryARCResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                         `json:"dataSource,required"`
	Description     string                                                         `json:"description,required"`
	EventType       string                                                         `json:"eventType,required"`
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                      `json:"startTime" format:"date-time"`
	JSON            emailRoutingSummaryARCResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingSummaryARCResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [EmailRoutingSummaryARCResponseMetaConfidenceInfoAnnotation]
type emailRoutingSummaryARCResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingSummaryARCResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryARCResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDKIMResponse struct {
	Meta     EmailRoutingSummaryDKIMResponseMeta `json:"meta,required"`
	Summary0 RadarEmailSummary                   `json:"summary_0,required"`
	JSON     emailRoutingSummaryDKIMResponseJSON `json:"-"`
}

// emailRoutingSummaryDKIMResponseJSON contains the JSON metadata for the struct
// [EmailRoutingSummaryDKIMResponse]
type emailRoutingSummaryDKIMResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryDKIMResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDKIMResponseJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDKIMResponseMeta struct {
	DateRange      []EmailRoutingSummaryDKIMResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                            `json:"lastUpdated,required"`
	Normalization  string                                            `json:"normalization,required"`
	ConfidenceInfo EmailRoutingSummaryDKIMResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           emailRoutingSummaryDKIMResponseMetaJSON           `json:"-"`
}

// emailRoutingSummaryDKIMResponseMetaJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryDKIMResponseMeta]
type emailRoutingSummaryDKIMResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingSummaryDKIMResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDKIMResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDKIMResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingSummaryDKIMResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingSummaryDKIMResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [EmailRoutingSummaryDKIMResponseMetaDateRange]
type emailRoutingSummaryDKIMResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryDKIMResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDKIMResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDKIMResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                         `json:"level"`
	JSON        emailRoutingSummaryDKIMResponseMetaConfidenceInfoJSON         `json:"-"`
}

// emailRoutingSummaryDKIMResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [EmailRoutingSummaryDKIMResponseMetaConfidenceInfo]
type emailRoutingSummaryDKIMResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryDKIMResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDKIMResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                          `json:"dataSource,required"`
	Description     string                                                          `json:"description,required"`
	EventType       string                                                          `json:"eventType,required"`
	IsInstantaneous bool                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                       `json:"startTime" format:"date-time"`
	JSON            emailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotation]
type emailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDKIMResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDMARCResponse struct {
	Meta     EmailRoutingSummaryDMARCResponseMeta `json:"meta,required"`
	Summary0 RadarEmailSummary                    `json:"summary_0,required"`
	JSON     emailRoutingSummaryDMARCResponseJSON `json:"-"`
}

// emailRoutingSummaryDMARCResponseJSON contains the JSON metadata for the struct
// [EmailRoutingSummaryDMARCResponse]
type emailRoutingSummaryDMARCResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryDMARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDMARCResponseJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDMARCResponseMeta struct {
	DateRange      []EmailRoutingSummaryDMARCResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                             `json:"lastUpdated,required"`
	Normalization  string                                             `json:"normalization,required"`
	ConfidenceInfo EmailRoutingSummaryDMARCResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           emailRoutingSummaryDMARCResponseMetaJSON           `json:"-"`
}

// emailRoutingSummaryDMARCResponseMetaJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryDMARCResponseMeta]
type emailRoutingSummaryDMARCResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingSummaryDMARCResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDMARCResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDMARCResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingSummaryDMARCResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingSummaryDMARCResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [EmailRoutingSummaryDMARCResponseMetaDateRange]
type emailRoutingSummaryDMARCResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryDMARCResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDMARCResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDMARCResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                          `json:"level"`
	JSON        emailRoutingSummaryDMARCResponseMetaConfidenceInfoJSON         `json:"-"`
}

// emailRoutingSummaryDMARCResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [EmailRoutingSummaryDMARCResponseMetaConfidenceInfo]
type emailRoutingSummaryDMARCResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryDMARCResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDMARCResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                           `json:"dataSource,required"`
	Description     string                                                           `json:"description,required"`
	EventType       string                                                           `json:"eventType,required"`
	IsInstantaneous bool                                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                                        `json:"endTime" format:"date-time"`
	LinkedURL       string                                                           `json:"linkedUrl"`
	StartTime       time.Time                                                        `json:"startTime" format:"date-time"`
	JSON            emailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotation]
type emailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDMARCResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryEncryptedResponse struct {
	Meta     EmailRoutingSummaryEncryptedResponseMeta     `json:"meta,required"`
	Summary0 EmailRoutingSummaryEncryptedResponseSummary0 `json:"summary_0,required"`
	JSON     emailRoutingSummaryEncryptedResponseJSON     `json:"-"`
}

// emailRoutingSummaryEncryptedResponseJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryEncryptedResponse]
type emailRoutingSummaryEncryptedResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryEncryptedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryEncryptedResponseJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryEncryptedResponseMeta struct {
	DateRange      []EmailRoutingSummaryEncryptedResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                 `json:"lastUpdated,required"`
	Normalization  string                                                 `json:"normalization,required"`
	ConfidenceInfo EmailRoutingSummaryEncryptedResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           emailRoutingSummaryEncryptedResponseMetaJSON           `json:"-"`
}

// emailRoutingSummaryEncryptedResponseMetaJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryEncryptedResponseMeta]
type emailRoutingSummaryEncryptedResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingSummaryEncryptedResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryEncryptedResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryEncryptedResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingSummaryEncryptedResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingSummaryEncryptedResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [EmailRoutingSummaryEncryptedResponseMetaDateRange]
type emailRoutingSummaryEncryptedResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryEncryptedResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryEncryptedResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryEncryptedResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        emailRoutingSummaryEncryptedResponseMetaConfidenceInfoJSON         `json:"-"`
}

// emailRoutingSummaryEncryptedResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [EmailRoutingSummaryEncryptedResponseMetaConfidenceInfo]
type emailRoutingSummaryEncryptedResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryEncryptedResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryEncryptedResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous bool                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            emailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotation]
type emailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryEncryptedResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryEncryptedResponseSummary0 struct {
	Encrypted    string                                           `json:"ENCRYPTED,required"`
	NotEncrypted string                                           `json:"NOT_ENCRYPTED,required"`
	JSON         emailRoutingSummaryEncryptedResponseSummary0JSON `json:"-"`
}

// emailRoutingSummaryEncryptedResponseSummary0JSON contains the JSON metadata for
// the struct [EmailRoutingSummaryEncryptedResponseSummary0]
type emailRoutingSummaryEncryptedResponseSummary0JSON struct {
	Encrypted    apijson.Field
	NotEncrypted apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EmailRoutingSummaryEncryptedResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryEncryptedResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryIPVersionResponse struct {
	Meta     EmailRoutingSummaryIPVersionResponseMeta     `json:"meta,required"`
	Summary0 EmailRoutingSummaryIPVersionResponseSummary0 `json:"summary_0,required"`
	JSON     emailRoutingSummaryIPVersionResponseJSON     `json:"-"`
}

// emailRoutingSummaryIPVersionResponseJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryIPVersionResponse]
type emailRoutingSummaryIPVersionResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryIPVersionResponseMeta struct {
	DateRange      []EmailRoutingSummaryIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                 `json:"lastUpdated,required"`
	Normalization  string                                                 `json:"normalization,required"`
	ConfidenceInfo EmailRoutingSummaryIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           emailRoutingSummaryIPVersionResponseMetaJSON           `json:"-"`
}

// emailRoutingSummaryIPVersionResponseMetaJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryIPVersionResponseMeta]
type emailRoutingSummaryIPVersionResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingSummaryIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingSummaryIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingSummaryIPVersionResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [EmailRoutingSummaryIPVersionResponseMetaDateRange]
type emailRoutingSummaryIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryIPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryIPVersionResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        emailRoutingSummaryIPVersionResponseMetaConfidenceInfoJSON         `json:"-"`
}

// emailRoutingSummaryIPVersionResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [EmailRoutingSummaryIPVersionResponseMetaConfidenceInfo]
type emailRoutingSummaryIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous bool                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            emailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotation]
type emailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryIPVersionResponseSummary0 struct {
	IPv4 string                                           `json:"IPv4,required"`
	IPv6 string                                           `json:"IPv6,required"`
	JSON emailRoutingSummaryIPVersionResponseSummary0JSON `json:"-"`
}

// emailRoutingSummaryIPVersionResponseSummary0JSON contains the JSON metadata for
// the struct [EmailRoutingSummaryIPVersionResponseSummary0]
type emailRoutingSummaryIPVersionResponseSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryIPVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryIPVersionResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummarySPFResponse struct {
	Meta     EmailRoutingSummarySPFResponseMeta `json:"meta,required"`
	Summary0 RadarEmailSummary                  `json:"summary_0,required"`
	JSON     emailRoutingSummarySPFResponseJSON `json:"-"`
}

// emailRoutingSummarySPFResponseJSON contains the JSON metadata for the struct
// [EmailRoutingSummarySPFResponse]
type emailRoutingSummarySPFResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummarySPFResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummarySPFResponseJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummarySPFResponseMeta struct {
	DateRange      []EmailRoutingSummarySPFResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                           `json:"lastUpdated,required"`
	Normalization  string                                           `json:"normalization,required"`
	ConfidenceInfo EmailRoutingSummarySPFResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           emailRoutingSummarySPFResponseMetaJSON           `json:"-"`
}

// emailRoutingSummarySPFResponseMetaJSON contains the JSON metadata for the struct
// [EmailRoutingSummarySPFResponseMeta]
type emailRoutingSummarySPFResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingSummarySPFResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummarySPFResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummarySPFResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingSummarySPFResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingSummarySPFResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [EmailRoutingSummarySPFResponseMetaDateRange]
type emailRoutingSummarySPFResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummarySPFResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummarySPFResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummarySPFResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingSummarySPFResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                        `json:"level"`
	JSON        emailRoutingSummarySPFResponseMetaConfidenceInfoJSON         `json:"-"`
}

// emailRoutingSummarySPFResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [EmailRoutingSummarySPFResponseMetaConfidenceInfo]
type emailRoutingSummarySPFResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummarySPFResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummarySPFResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummarySPFResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                         `json:"dataSource,required"`
	Description     string                                                         `json:"description,required"`
	EventType       string                                                         `json:"eventType,required"`
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                      `json:"startTime" format:"date-time"`
	JSON            emailRoutingSummarySPFResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingSummarySPFResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [EmailRoutingSummarySPFResponseMetaConfidenceInfoAnnotation]
type emailRoutingSummarySPFResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingSummarySPFResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummarySPFResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryARCParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailRoutingSummaryARCParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailRoutingSummaryARCParamsDMARC] `query:"dmarc"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]EmailRoutingSummaryARCParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[EmailRoutingSummaryARCParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]EmailRoutingSummaryARCParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailRoutingSummaryARCParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingSummaryARCParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingSummaryARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailRoutingSummaryARCParamsDKIM string

const (
	EmailRoutingSummaryARCParamsDKIMPass EmailRoutingSummaryARCParamsDKIM = "PASS"
	EmailRoutingSummaryARCParamsDKIMNone EmailRoutingSummaryARCParamsDKIM = "NONE"
	EmailRoutingSummaryARCParamsDKIMFail EmailRoutingSummaryARCParamsDKIM = "FAIL"
)

func (r EmailRoutingSummaryARCParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryARCParamsDKIMPass, EmailRoutingSummaryARCParamsDKIMNone, EmailRoutingSummaryARCParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingSummaryARCParamsDMARC string

const (
	EmailRoutingSummaryARCParamsDMARCPass EmailRoutingSummaryARCParamsDMARC = "PASS"
	EmailRoutingSummaryARCParamsDMARCNone EmailRoutingSummaryARCParamsDMARC = "NONE"
	EmailRoutingSummaryARCParamsDMARCFail EmailRoutingSummaryARCParamsDMARC = "FAIL"
)

func (r EmailRoutingSummaryARCParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryARCParamsDMARCPass, EmailRoutingSummaryARCParamsDMARCNone, EmailRoutingSummaryARCParamsDMARCFail:
		return true
	}
	return false
}

type EmailRoutingSummaryARCParamsEncrypted string

const (
	EmailRoutingSummaryARCParamsEncryptedEncrypted    EmailRoutingSummaryARCParamsEncrypted = "ENCRYPTED"
	EmailRoutingSummaryARCParamsEncryptedNotEncrypted EmailRoutingSummaryARCParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingSummaryARCParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryARCParamsEncryptedEncrypted, EmailRoutingSummaryARCParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format results are returned in.
type EmailRoutingSummaryARCParamsFormat string

const (
	EmailRoutingSummaryARCParamsFormatJson EmailRoutingSummaryARCParamsFormat = "JSON"
	EmailRoutingSummaryARCParamsFormatCsv  EmailRoutingSummaryARCParamsFormat = "CSV"
)

func (r EmailRoutingSummaryARCParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryARCParamsFormatJson, EmailRoutingSummaryARCParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingSummaryARCParamsIPVersion string

const (
	EmailRoutingSummaryARCParamsIPVersionIPv4 EmailRoutingSummaryARCParamsIPVersion = "IPv4"
	EmailRoutingSummaryARCParamsIPVersionIPv6 EmailRoutingSummaryARCParamsIPVersion = "IPv6"
)

func (r EmailRoutingSummaryARCParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryARCParamsIPVersionIPv4, EmailRoutingSummaryARCParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingSummaryARCParamsSPF string

const (
	EmailRoutingSummaryARCParamsSPFPass EmailRoutingSummaryARCParamsSPF = "PASS"
	EmailRoutingSummaryARCParamsSPFNone EmailRoutingSummaryARCParamsSPF = "NONE"
	EmailRoutingSummaryARCParamsSPFFail EmailRoutingSummaryARCParamsSPF = "FAIL"
)

func (r EmailRoutingSummaryARCParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryARCParamsSPFPass, EmailRoutingSummaryARCParamsSPFNone, EmailRoutingSummaryARCParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingSummaryARCResponseEnvelope struct {
	Result  EmailRoutingSummaryARCResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    emailRoutingSummaryARCResponseEnvelopeJSON `json:"-"`
}

// emailRoutingSummaryARCResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryARCResponseEnvelope]
type emailRoutingSummaryARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDKIMParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailRoutingSummaryDKIMParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailRoutingSummaryDKIMParamsDMARC] `query:"dmarc"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]EmailRoutingSummaryDKIMParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[EmailRoutingSummaryDKIMParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]EmailRoutingSummaryDKIMParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailRoutingSummaryDKIMParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingSummaryDKIMParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingSummaryDKIMParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailRoutingSummaryDKIMParamsARC string

const (
	EmailRoutingSummaryDKIMParamsARCPass EmailRoutingSummaryDKIMParamsARC = "PASS"
	EmailRoutingSummaryDKIMParamsARCNone EmailRoutingSummaryDKIMParamsARC = "NONE"
	EmailRoutingSummaryDKIMParamsARCFail EmailRoutingSummaryDKIMParamsARC = "FAIL"
)

func (r EmailRoutingSummaryDKIMParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDKIMParamsARCPass, EmailRoutingSummaryDKIMParamsARCNone, EmailRoutingSummaryDKIMParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingSummaryDKIMParamsDMARC string

const (
	EmailRoutingSummaryDKIMParamsDMARCPass EmailRoutingSummaryDKIMParamsDMARC = "PASS"
	EmailRoutingSummaryDKIMParamsDMARCNone EmailRoutingSummaryDKIMParamsDMARC = "NONE"
	EmailRoutingSummaryDKIMParamsDMARCFail EmailRoutingSummaryDKIMParamsDMARC = "FAIL"
)

func (r EmailRoutingSummaryDKIMParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDKIMParamsDMARCPass, EmailRoutingSummaryDKIMParamsDMARCNone, EmailRoutingSummaryDKIMParamsDMARCFail:
		return true
	}
	return false
}

type EmailRoutingSummaryDKIMParamsEncrypted string

const (
	EmailRoutingSummaryDKIMParamsEncryptedEncrypted    EmailRoutingSummaryDKIMParamsEncrypted = "ENCRYPTED"
	EmailRoutingSummaryDKIMParamsEncryptedNotEncrypted EmailRoutingSummaryDKIMParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingSummaryDKIMParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDKIMParamsEncryptedEncrypted, EmailRoutingSummaryDKIMParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format results are returned in.
type EmailRoutingSummaryDKIMParamsFormat string

const (
	EmailRoutingSummaryDKIMParamsFormatJson EmailRoutingSummaryDKIMParamsFormat = "JSON"
	EmailRoutingSummaryDKIMParamsFormatCsv  EmailRoutingSummaryDKIMParamsFormat = "CSV"
)

func (r EmailRoutingSummaryDKIMParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDKIMParamsFormatJson, EmailRoutingSummaryDKIMParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingSummaryDKIMParamsIPVersion string

const (
	EmailRoutingSummaryDKIMParamsIPVersionIPv4 EmailRoutingSummaryDKIMParamsIPVersion = "IPv4"
	EmailRoutingSummaryDKIMParamsIPVersionIPv6 EmailRoutingSummaryDKIMParamsIPVersion = "IPv6"
)

func (r EmailRoutingSummaryDKIMParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDKIMParamsIPVersionIPv4, EmailRoutingSummaryDKIMParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingSummaryDKIMParamsSPF string

const (
	EmailRoutingSummaryDKIMParamsSPFPass EmailRoutingSummaryDKIMParamsSPF = "PASS"
	EmailRoutingSummaryDKIMParamsSPFNone EmailRoutingSummaryDKIMParamsSPF = "NONE"
	EmailRoutingSummaryDKIMParamsSPFFail EmailRoutingSummaryDKIMParamsSPF = "FAIL"
)

func (r EmailRoutingSummaryDKIMParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDKIMParamsSPFPass, EmailRoutingSummaryDKIMParamsSPFNone, EmailRoutingSummaryDKIMParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingSummaryDKIMResponseEnvelope struct {
	Result  EmailRoutingSummaryDKIMResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    emailRoutingSummaryDKIMResponseEnvelopeJSON `json:"-"`
}

// emailRoutingSummaryDKIMResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryDKIMResponseEnvelope]
type emailRoutingSummaryDKIMResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryDKIMResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDKIMResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryDMARCParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailRoutingSummaryDMARCParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailRoutingSummaryDMARCParamsDKIM] `query:"dkim"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]EmailRoutingSummaryDMARCParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[EmailRoutingSummaryDMARCParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]EmailRoutingSummaryDMARCParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailRoutingSummaryDMARCParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingSummaryDMARCParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingSummaryDMARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailRoutingSummaryDMARCParamsARC string

const (
	EmailRoutingSummaryDMARCParamsARCPass EmailRoutingSummaryDMARCParamsARC = "PASS"
	EmailRoutingSummaryDMARCParamsARCNone EmailRoutingSummaryDMARCParamsARC = "NONE"
	EmailRoutingSummaryDMARCParamsARCFail EmailRoutingSummaryDMARCParamsARC = "FAIL"
)

func (r EmailRoutingSummaryDMARCParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDMARCParamsARCPass, EmailRoutingSummaryDMARCParamsARCNone, EmailRoutingSummaryDMARCParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingSummaryDMARCParamsDKIM string

const (
	EmailRoutingSummaryDMARCParamsDKIMPass EmailRoutingSummaryDMARCParamsDKIM = "PASS"
	EmailRoutingSummaryDMARCParamsDKIMNone EmailRoutingSummaryDMARCParamsDKIM = "NONE"
	EmailRoutingSummaryDMARCParamsDKIMFail EmailRoutingSummaryDMARCParamsDKIM = "FAIL"
)

func (r EmailRoutingSummaryDMARCParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDMARCParamsDKIMPass, EmailRoutingSummaryDMARCParamsDKIMNone, EmailRoutingSummaryDMARCParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingSummaryDMARCParamsEncrypted string

const (
	EmailRoutingSummaryDMARCParamsEncryptedEncrypted    EmailRoutingSummaryDMARCParamsEncrypted = "ENCRYPTED"
	EmailRoutingSummaryDMARCParamsEncryptedNotEncrypted EmailRoutingSummaryDMARCParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingSummaryDMARCParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDMARCParamsEncryptedEncrypted, EmailRoutingSummaryDMARCParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format results are returned in.
type EmailRoutingSummaryDMARCParamsFormat string

const (
	EmailRoutingSummaryDMARCParamsFormatJson EmailRoutingSummaryDMARCParamsFormat = "JSON"
	EmailRoutingSummaryDMARCParamsFormatCsv  EmailRoutingSummaryDMARCParamsFormat = "CSV"
)

func (r EmailRoutingSummaryDMARCParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDMARCParamsFormatJson, EmailRoutingSummaryDMARCParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingSummaryDMARCParamsIPVersion string

const (
	EmailRoutingSummaryDMARCParamsIPVersionIPv4 EmailRoutingSummaryDMARCParamsIPVersion = "IPv4"
	EmailRoutingSummaryDMARCParamsIPVersionIPv6 EmailRoutingSummaryDMARCParamsIPVersion = "IPv6"
)

func (r EmailRoutingSummaryDMARCParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDMARCParamsIPVersionIPv4, EmailRoutingSummaryDMARCParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingSummaryDMARCParamsSPF string

const (
	EmailRoutingSummaryDMARCParamsSPFPass EmailRoutingSummaryDMARCParamsSPF = "PASS"
	EmailRoutingSummaryDMARCParamsSPFNone EmailRoutingSummaryDMARCParamsSPF = "NONE"
	EmailRoutingSummaryDMARCParamsSPFFail EmailRoutingSummaryDMARCParamsSPF = "FAIL"
)

func (r EmailRoutingSummaryDMARCParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryDMARCParamsSPFPass, EmailRoutingSummaryDMARCParamsSPFNone, EmailRoutingSummaryDMARCParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingSummaryDMARCResponseEnvelope struct {
	Result  EmailRoutingSummaryDMARCResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    emailRoutingSummaryDMARCResponseEnvelopeJSON `json:"-"`
}

// emailRoutingSummaryDMARCResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingSummaryDMARCResponseEnvelope]
type emailRoutingSummaryDMARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryDMARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryDMARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryEncryptedParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailRoutingSummaryEncryptedParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailRoutingSummaryEncryptedParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailRoutingSummaryEncryptedParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[EmailRoutingSummaryEncryptedParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]EmailRoutingSummaryEncryptedParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailRoutingSummaryEncryptedParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingSummaryEncryptedParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingSummaryEncryptedParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailRoutingSummaryEncryptedParamsARC string

const (
	EmailRoutingSummaryEncryptedParamsARCPass EmailRoutingSummaryEncryptedParamsARC = "PASS"
	EmailRoutingSummaryEncryptedParamsARCNone EmailRoutingSummaryEncryptedParamsARC = "NONE"
	EmailRoutingSummaryEncryptedParamsARCFail EmailRoutingSummaryEncryptedParamsARC = "FAIL"
)

func (r EmailRoutingSummaryEncryptedParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryEncryptedParamsARCPass, EmailRoutingSummaryEncryptedParamsARCNone, EmailRoutingSummaryEncryptedParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingSummaryEncryptedParamsDKIM string

const (
	EmailRoutingSummaryEncryptedParamsDKIMPass EmailRoutingSummaryEncryptedParamsDKIM = "PASS"
	EmailRoutingSummaryEncryptedParamsDKIMNone EmailRoutingSummaryEncryptedParamsDKIM = "NONE"
	EmailRoutingSummaryEncryptedParamsDKIMFail EmailRoutingSummaryEncryptedParamsDKIM = "FAIL"
)

func (r EmailRoutingSummaryEncryptedParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryEncryptedParamsDKIMPass, EmailRoutingSummaryEncryptedParamsDKIMNone, EmailRoutingSummaryEncryptedParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingSummaryEncryptedParamsDMARC string

const (
	EmailRoutingSummaryEncryptedParamsDMARCPass EmailRoutingSummaryEncryptedParamsDMARC = "PASS"
	EmailRoutingSummaryEncryptedParamsDMARCNone EmailRoutingSummaryEncryptedParamsDMARC = "NONE"
	EmailRoutingSummaryEncryptedParamsDMARCFail EmailRoutingSummaryEncryptedParamsDMARC = "FAIL"
)

func (r EmailRoutingSummaryEncryptedParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryEncryptedParamsDMARCPass, EmailRoutingSummaryEncryptedParamsDMARCNone, EmailRoutingSummaryEncryptedParamsDMARCFail:
		return true
	}
	return false
}

// Format results are returned in.
type EmailRoutingSummaryEncryptedParamsFormat string

const (
	EmailRoutingSummaryEncryptedParamsFormatJson EmailRoutingSummaryEncryptedParamsFormat = "JSON"
	EmailRoutingSummaryEncryptedParamsFormatCsv  EmailRoutingSummaryEncryptedParamsFormat = "CSV"
)

func (r EmailRoutingSummaryEncryptedParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryEncryptedParamsFormatJson, EmailRoutingSummaryEncryptedParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingSummaryEncryptedParamsIPVersion string

const (
	EmailRoutingSummaryEncryptedParamsIPVersionIPv4 EmailRoutingSummaryEncryptedParamsIPVersion = "IPv4"
	EmailRoutingSummaryEncryptedParamsIPVersionIPv6 EmailRoutingSummaryEncryptedParamsIPVersion = "IPv6"
)

func (r EmailRoutingSummaryEncryptedParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryEncryptedParamsIPVersionIPv4, EmailRoutingSummaryEncryptedParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingSummaryEncryptedParamsSPF string

const (
	EmailRoutingSummaryEncryptedParamsSPFPass EmailRoutingSummaryEncryptedParamsSPF = "PASS"
	EmailRoutingSummaryEncryptedParamsSPFNone EmailRoutingSummaryEncryptedParamsSPF = "NONE"
	EmailRoutingSummaryEncryptedParamsSPFFail EmailRoutingSummaryEncryptedParamsSPF = "FAIL"
)

func (r EmailRoutingSummaryEncryptedParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryEncryptedParamsSPFPass, EmailRoutingSummaryEncryptedParamsSPFNone, EmailRoutingSummaryEncryptedParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingSummaryEncryptedResponseEnvelope struct {
	Result  EmailRoutingSummaryEncryptedResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    emailRoutingSummaryEncryptedResponseEnvelopeJSON `json:"-"`
}

// emailRoutingSummaryEncryptedResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailRoutingSummaryEncryptedResponseEnvelope]
type emailRoutingSummaryEncryptedResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryEncryptedResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryEncryptedResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummaryIPVersionParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailRoutingSummaryIPVersionParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailRoutingSummaryIPVersionParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailRoutingSummaryIPVersionParamsDMARC] `query:"dmarc"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]EmailRoutingSummaryIPVersionParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[EmailRoutingSummaryIPVersionParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]EmailRoutingSummaryIPVersionParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingSummaryIPVersionParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingSummaryIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailRoutingSummaryIPVersionParamsARC string

const (
	EmailRoutingSummaryIPVersionParamsARCPass EmailRoutingSummaryIPVersionParamsARC = "PASS"
	EmailRoutingSummaryIPVersionParamsARCNone EmailRoutingSummaryIPVersionParamsARC = "NONE"
	EmailRoutingSummaryIPVersionParamsARCFail EmailRoutingSummaryIPVersionParamsARC = "FAIL"
)

func (r EmailRoutingSummaryIPVersionParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryIPVersionParamsARCPass, EmailRoutingSummaryIPVersionParamsARCNone, EmailRoutingSummaryIPVersionParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingSummaryIPVersionParamsDKIM string

const (
	EmailRoutingSummaryIPVersionParamsDKIMPass EmailRoutingSummaryIPVersionParamsDKIM = "PASS"
	EmailRoutingSummaryIPVersionParamsDKIMNone EmailRoutingSummaryIPVersionParamsDKIM = "NONE"
	EmailRoutingSummaryIPVersionParamsDKIMFail EmailRoutingSummaryIPVersionParamsDKIM = "FAIL"
)

func (r EmailRoutingSummaryIPVersionParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryIPVersionParamsDKIMPass, EmailRoutingSummaryIPVersionParamsDKIMNone, EmailRoutingSummaryIPVersionParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingSummaryIPVersionParamsDMARC string

const (
	EmailRoutingSummaryIPVersionParamsDMARCPass EmailRoutingSummaryIPVersionParamsDMARC = "PASS"
	EmailRoutingSummaryIPVersionParamsDMARCNone EmailRoutingSummaryIPVersionParamsDMARC = "NONE"
	EmailRoutingSummaryIPVersionParamsDMARCFail EmailRoutingSummaryIPVersionParamsDMARC = "FAIL"
)

func (r EmailRoutingSummaryIPVersionParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryIPVersionParamsDMARCPass, EmailRoutingSummaryIPVersionParamsDMARCNone, EmailRoutingSummaryIPVersionParamsDMARCFail:
		return true
	}
	return false
}

type EmailRoutingSummaryIPVersionParamsEncrypted string

const (
	EmailRoutingSummaryIPVersionParamsEncryptedEncrypted    EmailRoutingSummaryIPVersionParamsEncrypted = "ENCRYPTED"
	EmailRoutingSummaryIPVersionParamsEncryptedNotEncrypted EmailRoutingSummaryIPVersionParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingSummaryIPVersionParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryIPVersionParamsEncryptedEncrypted, EmailRoutingSummaryIPVersionParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format results are returned in.
type EmailRoutingSummaryIPVersionParamsFormat string

const (
	EmailRoutingSummaryIPVersionParamsFormatJson EmailRoutingSummaryIPVersionParamsFormat = "JSON"
	EmailRoutingSummaryIPVersionParamsFormatCsv  EmailRoutingSummaryIPVersionParamsFormat = "CSV"
)

func (r EmailRoutingSummaryIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryIPVersionParamsFormatJson, EmailRoutingSummaryIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingSummaryIPVersionParamsSPF string

const (
	EmailRoutingSummaryIPVersionParamsSPFPass EmailRoutingSummaryIPVersionParamsSPF = "PASS"
	EmailRoutingSummaryIPVersionParamsSPFNone EmailRoutingSummaryIPVersionParamsSPF = "NONE"
	EmailRoutingSummaryIPVersionParamsSPFFail EmailRoutingSummaryIPVersionParamsSPF = "FAIL"
)

func (r EmailRoutingSummaryIPVersionParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingSummaryIPVersionParamsSPFPass, EmailRoutingSummaryIPVersionParamsSPFNone, EmailRoutingSummaryIPVersionParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingSummaryIPVersionResponseEnvelope struct {
	Result  EmailRoutingSummaryIPVersionResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    emailRoutingSummaryIPVersionResponseEnvelopeJSON `json:"-"`
}

// emailRoutingSummaryIPVersionResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailRoutingSummaryIPVersionResponseEnvelope]
type emailRoutingSummaryIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummaryIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummaryIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingSummarySPFParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]EmailRoutingSummarySPFParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]EmailRoutingSummarySPFParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]EmailRoutingSummarySPFParamsDMARC] `query:"dmarc"`
	// Filter for encrypted emails.
	Encrypted param.Field[[]EmailRoutingSummarySPFParamsEncrypted] `query:"encrypted"`
	// Format results are returned in.
	Format param.Field[EmailRoutingSummarySPFParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]EmailRoutingSummarySPFParamsIPVersion] `query:"ipVersion"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [EmailRoutingSummarySPFParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingSummarySPFParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailRoutingSummarySPFParamsARC string

const (
	EmailRoutingSummarySPFParamsARCPass EmailRoutingSummarySPFParamsARC = "PASS"
	EmailRoutingSummarySPFParamsARCNone EmailRoutingSummarySPFParamsARC = "NONE"
	EmailRoutingSummarySPFParamsARCFail EmailRoutingSummarySPFParamsARC = "FAIL"
)

func (r EmailRoutingSummarySPFParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummarySPFParamsARCPass, EmailRoutingSummarySPFParamsARCNone, EmailRoutingSummarySPFParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingSummarySPFParamsDKIM string

const (
	EmailRoutingSummarySPFParamsDKIMPass EmailRoutingSummarySPFParamsDKIM = "PASS"
	EmailRoutingSummarySPFParamsDKIMNone EmailRoutingSummarySPFParamsDKIM = "NONE"
	EmailRoutingSummarySPFParamsDKIMFail EmailRoutingSummarySPFParamsDKIM = "FAIL"
)

func (r EmailRoutingSummarySPFParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingSummarySPFParamsDKIMPass, EmailRoutingSummarySPFParamsDKIMNone, EmailRoutingSummarySPFParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingSummarySPFParamsDMARC string

const (
	EmailRoutingSummarySPFParamsDMARCPass EmailRoutingSummarySPFParamsDMARC = "PASS"
	EmailRoutingSummarySPFParamsDMARCNone EmailRoutingSummarySPFParamsDMARC = "NONE"
	EmailRoutingSummarySPFParamsDMARCFail EmailRoutingSummarySPFParamsDMARC = "FAIL"
)

func (r EmailRoutingSummarySPFParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingSummarySPFParamsDMARCPass, EmailRoutingSummarySPFParamsDMARCNone, EmailRoutingSummarySPFParamsDMARCFail:
		return true
	}
	return false
}

type EmailRoutingSummarySPFParamsEncrypted string

const (
	EmailRoutingSummarySPFParamsEncryptedEncrypted    EmailRoutingSummarySPFParamsEncrypted = "ENCRYPTED"
	EmailRoutingSummarySPFParamsEncryptedNotEncrypted EmailRoutingSummarySPFParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingSummarySPFParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingSummarySPFParamsEncryptedEncrypted, EmailRoutingSummarySPFParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format results are returned in.
type EmailRoutingSummarySPFParamsFormat string

const (
	EmailRoutingSummarySPFParamsFormatJson EmailRoutingSummarySPFParamsFormat = "JSON"
	EmailRoutingSummarySPFParamsFormatCsv  EmailRoutingSummarySPFParamsFormat = "CSV"
)

func (r EmailRoutingSummarySPFParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingSummarySPFParamsFormatJson, EmailRoutingSummarySPFParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingSummarySPFParamsIPVersion string

const (
	EmailRoutingSummarySPFParamsIPVersionIPv4 EmailRoutingSummarySPFParamsIPVersion = "IPv4"
	EmailRoutingSummarySPFParamsIPVersionIPv6 EmailRoutingSummarySPFParamsIPVersion = "IPv6"
)

func (r EmailRoutingSummarySPFParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingSummarySPFParamsIPVersionIPv4, EmailRoutingSummarySPFParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingSummarySPFResponseEnvelope struct {
	Result  EmailRoutingSummarySPFResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    emailRoutingSummarySPFResponseEnvelopeJSON `json:"-"`
}

// emailRoutingSummarySPFResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingSummarySPFResponseEnvelope]
type emailRoutingSummarySPFResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingSummarySPFResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingSummarySPFResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
