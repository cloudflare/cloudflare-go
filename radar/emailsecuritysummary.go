// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// EmailSecuritySummaryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailSecuritySummaryService] method instead.
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

// Retrieves the distribution of emails by ARC (Authenticated Received Chain)
// validation.
//
// Deprecated: Use
// [Radar Email Security Summary By Dimension](https://developers.cloudflare.com/api/resources/radar/subresources/email/subresources/security/methods/summary_v2/)
// instead.
func (r *EmailSecuritySummaryService) ARC(ctx context.Context, query EmailSecuritySummaryARCParams, opts ...option.RequestOption) (res *EmailSecuritySummaryARCResponse, err error) {
	var env EmailSecuritySummaryARCResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/email/security/summary/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by DKIM (DomainKeys Identified Mail)
// validation.
//
// Deprecated: Use
// [Radar Email Security Summary By Dimension](https://developers.cloudflare.com/api/resources/radar/subresources/email/subresources/security/methods/summary_v2/)
// instead.
func (r *EmailSecuritySummaryService) DKIM(ctx context.Context, query EmailSecuritySummaryDKIMParams, opts ...option.RequestOption) (res *EmailSecuritySummaryDKIMResponse, err error) {
	var env EmailSecuritySummaryDKIMResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/email/security/summary/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by DMARC (Domain-based Message
// Authentication, Reporting and Conformance) validation.
//
// Deprecated: Use
// [Radar Email Security Summary By Dimension](https://developers.cloudflare.com/api/resources/radar/subresources/email/subresources/security/methods/summary_v2/)
// instead.
func (r *EmailSecuritySummaryService) DMARC(ctx context.Context, query EmailSecuritySummaryDMARCParams, opts ...option.RequestOption) (res *EmailSecuritySummaryDMARCResponse, err error) {
	var env EmailSecuritySummaryDMARCResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/email/security/summary/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by malicious classification.
//
// Deprecated: Use
// [Radar Email Security Summary By Dimension](https://developers.cloudflare.com/api/resources/radar/subresources/email/subresources/security/methods/summary_v2/)
// instead.
func (r *EmailSecuritySummaryService) Malicious(ctx context.Context, query EmailSecuritySummaryMaliciousParams, opts ...option.RequestOption) (res *EmailSecuritySummaryMaliciousResponse, err error) {
	var env EmailSecuritySummaryMaliciousResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/email/security/summary/malicious"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the proportion of emails by spam classification (spam vs. non-spam).
//
// Deprecated: Use
// [Radar Email Security Summary By Dimension](https://developers.cloudflare.com/api/resources/radar/subresources/email/subresources/security/methods/summary_v2/)
// instead.
func (r *EmailSecuritySummaryService) Spam(ctx context.Context, query EmailSecuritySummarySpamParams, opts ...option.RequestOption) (res *EmailSecuritySummarySpamResponse, err error) {
	var env EmailSecuritySummarySpamResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/email/security/summary/spam"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by SPF (Sender Policy Framework)
// validation.
//
// Deprecated: Use
// [Radar Email Security Summary By Dimension](https://developers.cloudflare.com/api/resources/radar/subresources/email/subresources/security/methods/summary_v2/)
// instead.
func (r *EmailSecuritySummaryService) SPF(ctx context.Context, query EmailSecuritySummarySPFParams, opts ...option.RequestOption) (res *EmailSecuritySummarySPFResponse, err error) {
	var env EmailSecuritySummarySPFResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/email/security/summary/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the proportion of emails by spoof classification (spoof vs.
// non-spoof).
//
// Deprecated: Use
// [Radar Email Security Summary By Dimension](https://developers.cloudflare.com/api/resources/radar/subresources/email/subresources/security/methods/summary_v2/)
// instead.
func (r *EmailSecuritySummaryService) Spoof(ctx context.Context, query EmailSecuritySummarySpoofParams, opts ...option.RequestOption) (res *EmailSecuritySummarySpoofResponse, err error) {
	var env EmailSecuritySummarySpoofResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/email/security/summary/spoof"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by threat categories.
//
// Deprecated: Use
// [Radar Email Security Summary By Dimension](https://developers.cloudflare.com/api/resources/radar/subresources/email/subresources/security/methods/summary_v2/)
// instead.
func (r *EmailSecuritySummaryService) ThreatCategory(ctx context.Context, query EmailSecuritySummaryThreatCategoryParams, opts ...option.RequestOption) (res *EmailSecuritySummaryThreatCategoryResponse, err error) {
	var env EmailSecuritySummaryThreatCategoryResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/email/security/summary/threat_category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by TLS version.
//
// Deprecated: Use
// [Radar Email Security Summary By Dimension](https://developers.cloudflare.com/api/resources/radar/subresources/email/subresources/security/methods/summary_v2/)
// instead.
func (r *EmailSecuritySummaryService) TLSVersion(ctx context.Context, query EmailSecuritySummaryTLSVersionParams, opts ...option.RequestOption) (res *EmailSecuritySummaryTLSVersionResponse, err error) {
	var env EmailSecuritySummaryTLSVersionResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/email/security/summary/tls_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailSecuritySummaryARCResponse struct {
	// Metadata for the results.
	Meta     EmailSecuritySummaryARCResponseMeta `json:"meta,required"`
	Summary0 RadarEmailSummary                   `json:"summary_0,required"`
	JSON     emailSecuritySummaryARCResponseJSON `json:"-"`
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

// Metadata for the results.
type EmailSecuritySummaryARCResponseMeta struct {
	ConfidenceInfo EmailSecuritySummaryARCResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecuritySummaryARCResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecuritySummaryARCResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecuritySummaryARCResponseMetaUnit `json:"units,required"`
	JSON  emailSecuritySummaryARCResponseMetaJSON   `json:"-"`
}

// emailSecuritySummaryARCResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecuritySummaryARCResponseMeta]
type emailSecuritySummaryARCResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecuritySummaryARCResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryARCResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryARCResponseMetaConfidenceInfo struct {
	Annotations []EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                 `json:"level,required"`
	JSON  emailSecuritySummaryARCResponseMetaConfidenceInfoJSON `json:"-"`
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

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                 `json:"description,required"`
	EndDate     time.Time                                                              `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                            `json:"isInstantaneous,required"`
	LinkedURL       string                                                          `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                       `json:"startDate,required" format:"date-time"`
	JSON            emailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotation]
type emailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailSecuritySummaryARCResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
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

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecuritySummaryARCResponseMetaNormalization string

const (
	EmailSecuritySummaryARCResponseMetaNormalizationPercentage           EmailSecuritySummaryARCResponseMetaNormalization = "PERCENTAGE"
	EmailSecuritySummaryARCResponseMetaNormalizationMin0Max              EmailSecuritySummaryARCResponseMetaNormalization = "MIN0_MAX"
	EmailSecuritySummaryARCResponseMetaNormalizationMinMax               EmailSecuritySummaryARCResponseMetaNormalization = "MIN_MAX"
	EmailSecuritySummaryARCResponseMetaNormalizationRawValues            EmailSecuritySummaryARCResponseMetaNormalization = "RAW_VALUES"
	EmailSecuritySummaryARCResponseMetaNormalizationPercentageChange     EmailSecuritySummaryARCResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecuritySummaryARCResponseMetaNormalizationRollingAverage       EmailSecuritySummaryARCResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecuritySummaryARCResponseMetaNormalizationOverlappedPercentage EmailSecuritySummaryARCResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecuritySummaryARCResponseMetaNormalizationRatio                EmailSecuritySummaryARCResponseMetaNormalization = "RATIO"
)

func (r EmailSecuritySummaryARCResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryARCResponseMetaNormalizationPercentage, EmailSecuritySummaryARCResponseMetaNormalizationMin0Max, EmailSecuritySummaryARCResponseMetaNormalizationMinMax, EmailSecuritySummaryARCResponseMetaNormalizationRawValues, EmailSecuritySummaryARCResponseMetaNormalizationPercentageChange, EmailSecuritySummaryARCResponseMetaNormalizationRollingAverage, EmailSecuritySummaryARCResponseMetaNormalizationOverlappedPercentage, EmailSecuritySummaryARCResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecuritySummaryARCResponseMetaUnit struct {
	Name  string                                      `json:"name,required"`
	Value string                                      `json:"value,required"`
	JSON  emailSecuritySummaryARCResponseMetaUnitJSON `json:"-"`
}

// emailSecuritySummaryARCResponseMetaUnitJSON contains the JSON metadata for the
// struct [EmailSecuritySummaryARCResponseMetaUnit]
type emailSecuritySummaryARCResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryARCResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryARCResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryDKIMResponse struct {
	// Metadata for the results.
	Meta     EmailSecuritySummaryDKIMResponseMeta `json:"meta,required"`
	Summary0 RadarEmailSummary                    `json:"summary_0,required"`
	JSON     emailSecuritySummaryDKIMResponseJSON `json:"-"`
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

// Metadata for the results.
type EmailSecuritySummaryDKIMResponseMeta struct {
	ConfidenceInfo EmailSecuritySummaryDKIMResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecuritySummaryDKIMResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecuritySummaryDKIMResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecuritySummaryDKIMResponseMetaUnit `json:"units,required"`
	JSON  emailSecuritySummaryDKIMResponseMetaJSON   `json:"-"`
}

// emailSecuritySummaryDKIMResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecuritySummaryDKIMResponseMeta]
type emailSecuritySummaryDKIMResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecuritySummaryDKIMResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDKIMResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryDKIMResponseMetaConfidenceInfo struct {
	Annotations []EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                  `json:"level,required"`
	JSON  emailSecuritySummaryDKIMResponseMetaConfidenceInfoJSON `json:"-"`
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

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                  `json:"description,required"`
	EndDate     time.Time                                                               `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                             `json:"isInstantaneous,required"`
	LinkedURL       string                                                           `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                        `json:"startDate,required" format:"date-time"`
	JSON            emailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotation]
type emailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailSecuritySummaryDKIMResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
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

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecuritySummaryDKIMResponseMetaNormalization string

const (
	EmailSecuritySummaryDKIMResponseMetaNormalizationPercentage           EmailSecuritySummaryDKIMResponseMetaNormalization = "PERCENTAGE"
	EmailSecuritySummaryDKIMResponseMetaNormalizationMin0Max              EmailSecuritySummaryDKIMResponseMetaNormalization = "MIN0_MAX"
	EmailSecuritySummaryDKIMResponseMetaNormalizationMinMax               EmailSecuritySummaryDKIMResponseMetaNormalization = "MIN_MAX"
	EmailSecuritySummaryDKIMResponseMetaNormalizationRawValues            EmailSecuritySummaryDKIMResponseMetaNormalization = "RAW_VALUES"
	EmailSecuritySummaryDKIMResponseMetaNormalizationPercentageChange     EmailSecuritySummaryDKIMResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecuritySummaryDKIMResponseMetaNormalizationRollingAverage       EmailSecuritySummaryDKIMResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecuritySummaryDKIMResponseMetaNormalizationOverlappedPercentage EmailSecuritySummaryDKIMResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecuritySummaryDKIMResponseMetaNormalizationRatio                EmailSecuritySummaryDKIMResponseMetaNormalization = "RATIO"
)

func (r EmailSecuritySummaryDKIMResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryDKIMResponseMetaNormalizationPercentage, EmailSecuritySummaryDKIMResponseMetaNormalizationMin0Max, EmailSecuritySummaryDKIMResponseMetaNormalizationMinMax, EmailSecuritySummaryDKIMResponseMetaNormalizationRawValues, EmailSecuritySummaryDKIMResponseMetaNormalizationPercentageChange, EmailSecuritySummaryDKIMResponseMetaNormalizationRollingAverage, EmailSecuritySummaryDKIMResponseMetaNormalizationOverlappedPercentage, EmailSecuritySummaryDKIMResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecuritySummaryDKIMResponseMetaUnit struct {
	Name  string                                       `json:"name,required"`
	Value string                                       `json:"value,required"`
	JSON  emailSecuritySummaryDKIMResponseMetaUnitJSON `json:"-"`
}

// emailSecuritySummaryDKIMResponseMetaUnitJSON contains the JSON metadata for the
// struct [EmailSecuritySummaryDKIMResponseMetaUnit]
type emailSecuritySummaryDKIMResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryDKIMResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDKIMResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryDMARCResponse struct {
	// Metadata for the results.
	Meta     EmailSecuritySummaryDMARCResponseMeta `json:"meta,required"`
	Summary0 RadarEmailSummary                     `json:"summary_0,required"`
	JSON     emailSecuritySummaryDMARCResponseJSON `json:"-"`
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

// Metadata for the results.
type EmailSecuritySummaryDMARCResponseMeta struct {
	ConfidenceInfo EmailSecuritySummaryDMARCResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecuritySummaryDMARCResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecuritySummaryDMARCResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecuritySummaryDMARCResponseMetaUnit `json:"units,required"`
	JSON  emailSecuritySummaryDMARCResponseMetaJSON   `json:"-"`
}

// emailSecuritySummaryDMARCResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecuritySummaryDMARCResponseMeta]
type emailSecuritySummaryDMARCResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecuritySummaryDMARCResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDMARCResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryDMARCResponseMetaConfidenceInfo struct {
	Annotations []EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                   `json:"level,required"`
	JSON  emailSecuritySummaryDMARCResponseMetaConfidenceInfoJSON `json:"-"`
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

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                   `json:"description,required"`
	EndDate     time.Time                                                                `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                              `json:"isInstantaneous,required"`
	LinkedURL       string                                                            `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                         `json:"startDate,required" format:"date-time"`
	JSON            emailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotation]
type emailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailSecuritySummaryDMARCResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
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

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecuritySummaryDMARCResponseMetaNormalization string

const (
	EmailSecuritySummaryDMARCResponseMetaNormalizationPercentage           EmailSecuritySummaryDMARCResponseMetaNormalization = "PERCENTAGE"
	EmailSecuritySummaryDMARCResponseMetaNormalizationMin0Max              EmailSecuritySummaryDMARCResponseMetaNormalization = "MIN0_MAX"
	EmailSecuritySummaryDMARCResponseMetaNormalizationMinMax               EmailSecuritySummaryDMARCResponseMetaNormalization = "MIN_MAX"
	EmailSecuritySummaryDMARCResponseMetaNormalizationRawValues            EmailSecuritySummaryDMARCResponseMetaNormalization = "RAW_VALUES"
	EmailSecuritySummaryDMARCResponseMetaNormalizationPercentageChange     EmailSecuritySummaryDMARCResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecuritySummaryDMARCResponseMetaNormalizationRollingAverage       EmailSecuritySummaryDMARCResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecuritySummaryDMARCResponseMetaNormalizationOverlappedPercentage EmailSecuritySummaryDMARCResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecuritySummaryDMARCResponseMetaNormalizationRatio                EmailSecuritySummaryDMARCResponseMetaNormalization = "RATIO"
)

func (r EmailSecuritySummaryDMARCResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryDMARCResponseMetaNormalizationPercentage, EmailSecuritySummaryDMARCResponseMetaNormalizationMin0Max, EmailSecuritySummaryDMARCResponseMetaNormalizationMinMax, EmailSecuritySummaryDMARCResponseMetaNormalizationRawValues, EmailSecuritySummaryDMARCResponseMetaNormalizationPercentageChange, EmailSecuritySummaryDMARCResponseMetaNormalizationRollingAverage, EmailSecuritySummaryDMARCResponseMetaNormalizationOverlappedPercentage, EmailSecuritySummaryDMARCResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecuritySummaryDMARCResponseMetaUnit struct {
	Name  string                                        `json:"name,required"`
	Value string                                        `json:"value,required"`
	JSON  emailSecuritySummaryDMARCResponseMetaUnitJSON `json:"-"`
}

// emailSecuritySummaryDMARCResponseMetaUnitJSON contains the JSON metadata for the
// struct [EmailSecuritySummaryDMARCResponseMetaUnit]
type emailSecuritySummaryDMARCResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryDMARCResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryDMARCResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryMaliciousResponse struct {
	// Metadata for the results.
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

// Metadata for the results.
type EmailSecuritySummaryMaliciousResponseMeta struct {
	ConfidenceInfo EmailSecuritySummaryMaliciousResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecuritySummaryMaliciousResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecuritySummaryMaliciousResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecuritySummaryMaliciousResponseMetaUnit `json:"units,required"`
	JSON  emailSecuritySummaryMaliciousResponseMetaJSON   `json:"-"`
}

// emailSecuritySummaryMaliciousResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecuritySummaryMaliciousResponseMeta]
type emailSecuritySummaryMaliciousResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecuritySummaryMaliciousResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryMaliciousResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryMaliciousResponseMetaConfidenceInfo struct {
	Annotations []EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                       `json:"level,required"`
	JSON  emailSecuritySummaryMaliciousResponseMetaConfidenceInfoJSON `json:"-"`
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

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                       `json:"description,required"`
	EndDate     time.Time                                                                    `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                             `json:"startDate,required" format:"date-time"`
	JSON            emailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotation]
type emailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailSecuritySummaryMaliciousResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
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

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecuritySummaryMaliciousResponseMetaNormalization string

const (
	EmailSecuritySummaryMaliciousResponseMetaNormalizationPercentage           EmailSecuritySummaryMaliciousResponseMetaNormalization = "PERCENTAGE"
	EmailSecuritySummaryMaliciousResponseMetaNormalizationMin0Max              EmailSecuritySummaryMaliciousResponseMetaNormalization = "MIN0_MAX"
	EmailSecuritySummaryMaliciousResponseMetaNormalizationMinMax               EmailSecuritySummaryMaliciousResponseMetaNormalization = "MIN_MAX"
	EmailSecuritySummaryMaliciousResponseMetaNormalizationRawValues            EmailSecuritySummaryMaliciousResponseMetaNormalization = "RAW_VALUES"
	EmailSecuritySummaryMaliciousResponseMetaNormalizationPercentageChange     EmailSecuritySummaryMaliciousResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecuritySummaryMaliciousResponseMetaNormalizationRollingAverage       EmailSecuritySummaryMaliciousResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecuritySummaryMaliciousResponseMetaNormalizationOverlappedPercentage EmailSecuritySummaryMaliciousResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecuritySummaryMaliciousResponseMetaNormalizationRatio                EmailSecuritySummaryMaliciousResponseMetaNormalization = "RATIO"
)

func (r EmailSecuritySummaryMaliciousResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryMaliciousResponseMetaNormalizationPercentage, EmailSecuritySummaryMaliciousResponseMetaNormalizationMin0Max, EmailSecuritySummaryMaliciousResponseMetaNormalizationMinMax, EmailSecuritySummaryMaliciousResponseMetaNormalizationRawValues, EmailSecuritySummaryMaliciousResponseMetaNormalizationPercentageChange, EmailSecuritySummaryMaliciousResponseMetaNormalizationRollingAverage, EmailSecuritySummaryMaliciousResponseMetaNormalizationOverlappedPercentage, EmailSecuritySummaryMaliciousResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecuritySummaryMaliciousResponseMetaUnit struct {
	Name  string                                            `json:"name,required"`
	Value string                                            `json:"value,required"`
	JSON  emailSecuritySummaryMaliciousResponseMetaUnitJSON `json:"-"`
}

// emailSecuritySummaryMaliciousResponseMetaUnitJSON contains the JSON metadata for
// the struct [EmailSecuritySummaryMaliciousResponseMetaUnit]
type emailSecuritySummaryMaliciousResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryMaliciousResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryMaliciousResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryMaliciousResponseSummary0 struct {
	// A numeric string.
	Malicious string `json:"MALICIOUS,required"`
	// A numeric string.
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
	// Metadata for the results.
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

// Metadata for the results.
type EmailSecuritySummarySpamResponseMeta struct {
	ConfidenceInfo EmailSecuritySummarySpamResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecuritySummarySpamResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecuritySummarySpamResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecuritySummarySpamResponseMetaUnit `json:"units,required"`
	JSON  emailSecuritySummarySpamResponseMetaJSON   `json:"-"`
}

// emailSecuritySummarySpamResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecuritySummarySpamResponseMeta]
type emailSecuritySummarySpamResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecuritySummarySpamResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpamResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySpamResponseMetaConfidenceInfo struct {
	Annotations []EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                  `json:"level,required"`
	JSON  emailSecuritySummarySpamResponseMetaConfidenceInfoJSON `json:"-"`
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

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                  `json:"description,required"`
	EndDate     time.Time                                                               `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                             `json:"isInstantaneous,required"`
	LinkedURL       string                                                           `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                        `json:"startDate,required" format:"date-time"`
	JSON            emailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotation]
type emailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailSecuritySummarySpamResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
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

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecuritySummarySpamResponseMetaNormalization string

const (
	EmailSecuritySummarySpamResponseMetaNormalizationPercentage           EmailSecuritySummarySpamResponseMetaNormalization = "PERCENTAGE"
	EmailSecuritySummarySpamResponseMetaNormalizationMin0Max              EmailSecuritySummarySpamResponseMetaNormalization = "MIN0_MAX"
	EmailSecuritySummarySpamResponseMetaNormalizationMinMax               EmailSecuritySummarySpamResponseMetaNormalization = "MIN_MAX"
	EmailSecuritySummarySpamResponseMetaNormalizationRawValues            EmailSecuritySummarySpamResponseMetaNormalization = "RAW_VALUES"
	EmailSecuritySummarySpamResponseMetaNormalizationPercentageChange     EmailSecuritySummarySpamResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecuritySummarySpamResponseMetaNormalizationRollingAverage       EmailSecuritySummarySpamResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecuritySummarySpamResponseMetaNormalizationOverlappedPercentage EmailSecuritySummarySpamResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecuritySummarySpamResponseMetaNormalizationRatio                EmailSecuritySummarySpamResponseMetaNormalization = "RATIO"
)

func (r EmailSecuritySummarySpamResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySpamResponseMetaNormalizationPercentage, EmailSecuritySummarySpamResponseMetaNormalizationMin0Max, EmailSecuritySummarySpamResponseMetaNormalizationMinMax, EmailSecuritySummarySpamResponseMetaNormalizationRawValues, EmailSecuritySummarySpamResponseMetaNormalizationPercentageChange, EmailSecuritySummarySpamResponseMetaNormalizationRollingAverage, EmailSecuritySummarySpamResponseMetaNormalizationOverlappedPercentage, EmailSecuritySummarySpamResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecuritySummarySpamResponseMetaUnit struct {
	Name  string                                       `json:"name,required"`
	Value string                                       `json:"value,required"`
	JSON  emailSecuritySummarySpamResponseMetaUnitJSON `json:"-"`
}

// emailSecuritySummarySpamResponseMetaUnitJSON contains the JSON metadata for the
// struct [EmailSecuritySummarySpamResponseMetaUnit]
type emailSecuritySummarySpamResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummarySpamResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpamResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySpamResponseSummary0 struct {
	// A numeric string.
	NotSpam string `json:"NOT_SPAM,required"`
	// A numeric string.
	Spam string                                       `json:"SPAM,required"`
	JSON emailSecuritySummarySpamResponseSummary0JSON `json:"-"`
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
	// Metadata for the results.
	Meta     EmailSecuritySummarySPFResponseMeta `json:"meta,required"`
	Summary0 RadarEmailSummary                   `json:"summary_0,required"`
	JSON     emailSecuritySummarySPFResponseJSON `json:"-"`
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

// Metadata for the results.
type EmailSecuritySummarySPFResponseMeta struct {
	ConfidenceInfo EmailSecuritySummarySPFResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecuritySummarySPFResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecuritySummarySPFResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecuritySummarySPFResponseMetaUnit `json:"units,required"`
	JSON  emailSecuritySummarySPFResponseMetaJSON   `json:"-"`
}

// emailSecuritySummarySPFResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecuritySummarySPFResponseMeta]
type emailSecuritySummarySPFResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecuritySummarySPFResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySPFResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySPFResponseMetaConfidenceInfo struct {
	Annotations []EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                 `json:"level,required"`
	JSON  emailSecuritySummarySPFResponseMetaConfidenceInfoJSON `json:"-"`
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

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                 `json:"description,required"`
	EndDate     time.Time                                                              `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                            `json:"isInstantaneous,required"`
	LinkedURL       string                                                          `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                       `json:"startDate,required" format:"date-time"`
	JSON            emailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotation]
type emailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailSecuritySummarySPFResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
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

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecuritySummarySPFResponseMetaNormalization string

const (
	EmailSecuritySummarySPFResponseMetaNormalizationPercentage           EmailSecuritySummarySPFResponseMetaNormalization = "PERCENTAGE"
	EmailSecuritySummarySPFResponseMetaNormalizationMin0Max              EmailSecuritySummarySPFResponseMetaNormalization = "MIN0_MAX"
	EmailSecuritySummarySPFResponseMetaNormalizationMinMax               EmailSecuritySummarySPFResponseMetaNormalization = "MIN_MAX"
	EmailSecuritySummarySPFResponseMetaNormalizationRawValues            EmailSecuritySummarySPFResponseMetaNormalization = "RAW_VALUES"
	EmailSecuritySummarySPFResponseMetaNormalizationPercentageChange     EmailSecuritySummarySPFResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecuritySummarySPFResponseMetaNormalizationRollingAverage       EmailSecuritySummarySPFResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecuritySummarySPFResponseMetaNormalizationOverlappedPercentage EmailSecuritySummarySPFResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecuritySummarySPFResponseMetaNormalizationRatio                EmailSecuritySummarySPFResponseMetaNormalization = "RATIO"
)

func (r EmailSecuritySummarySPFResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySPFResponseMetaNormalizationPercentage, EmailSecuritySummarySPFResponseMetaNormalizationMin0Max, EmailSecuritySummarySPFResponseMetaNormalizationMinMax, EmailSecuritySummarySPFResponseMetaNormalizationRawValues, EmailSecuritySummarySPFResponseMetaNormalizationPercentageChange, EmailSecuritySummarySPFResponseMetaNormalizationRollingAverage, EmailSecuritySummarySPFResponseMetaNormalizationOverlappedPercentage, EmailSecuritySummarySPFResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecuritySummarySPFResponseMetaUnit struct {
	Name  string                                      `json:"name,required"`
	Value string                                      `json:"value,required"`
	JSON  emailSecuritySummarySPFResponseMetaUnitJSON `json:"-"`
}

// emailSecuritySummarySPFResponseMetaUnitJSON contains the JSON metadata for the
// struct [EmailSecuritySummarySPFResponseMetaUnit]
type emailSecuritySummarySPFResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummarySPFResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySPFResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySpoofResponse struct {
	// Metadata for the results.
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

// Metadata for the results.
type EmailSecuritySummarySpoofResponseMeta struct {
	ConfidenceInfo EmailSecuritySummarySpoofResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecuritySummarySpoofResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecuritySummarySpoofResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecuritySummarySpoofResponseMetaUnit `json:"units,required"`
	JSON  emailSecuritySummarySpoofResponseMetaJSON   `json:"-"`
}

// emailSecuritySummarySpoofResponseMetaJSON contains the JSON metadata for the
// struct [EmailSecuritySummarySpoofResponseMeta]
type emailSecuritySummarySpoofResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecuritySummarySpoofResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpoofResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySpoofResponseMetaConfidenceInfo struct {
	Annotations []EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                   `json:"level,required"`
	JSON  emailSecuritySummarySpoofResponseMetaConfidenceInfoJSON `json:"-"`
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

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                   `json:"description,required"`
	EndDate     time.Time                                                                `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                              `json:"isInstantaneous,required"`
	LinkedURL       string                                                            `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                         `json:"startDate,required" format:"date-time"`
	JSON            emailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotation]
type emailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailSecuritySummarySpoofResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
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

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecuritySummarySpoofResponseMetaNormalization string

const (
	EmailSecuritySummarySpoofResponseMetaNormalizationPercentage           EmailSecuritySummarySpoofResponseMetaNormalization = "PERCENTAGE"
	EmailSecuritySummarySpoofResponseMetaNormalizationMin0Max              EmailSecuritySummarySpoofResponseMetaNormalization = "MIN0_MAX"
	EmailSecuritySummarySpoofResponseMetaNormalizationMinMax               EmailSecuritySummarySpoofResponseMetaNormalization = "MIN_MAX"
	EmailSecuritySummarySpoofResponseMetaNormalizationRawValues            EmailSecuritySummarySpoofResponseMetaNormalization = "RAW_VALUES"
	EmailSecuritySummarySpoofResponseMetaNormalizationPercentageChange     EmailSecuritySummarySpoofResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecuritySummarySpoofResponseMetaNormalizationRollingAverage       EmailSecuritySummarySpoofResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecuritySummarySpoofResponseMetaNormalizationOverlappedPercentage EmailSecuritySummarySpoofResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecuritySummarySpoofResponseMetaNormalizationRatio                EmailSecuritySummarySpoofResponseMetaNormalization = "RATIO"
)

func (r EmailSecuritySummarySpoofResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySpoofResponseMetaNormalizationPercentage, EmailSecuritySummarySpoofResponseMetaNormalizationMin0Max, EmailSecuritySummarySpoofResponseMetaNormalizationMinMax, EmailSecuritySummarySpoofResponseMetaNormalizationRawValues, EmailSecuritySummarySpoofResponseMetaNormalizationPercentageChange, EmailSecuritySummarySpoofResponseMetaNormalizationRollingAverage, EmailSecuritySummarySpoofResponseMetaNormalizationOverlappedPercentage, EmailSecuritySummarySpoofResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecuritySummarySpoofResponseMetaUnit struct {
	Name  string                                        `json:"name,required"`
	Value string                                        `json:"value,required"`
	JSON  emailSecuritySummarySpoofResponseMetaUnitJSON `json:"-"`
}

// emailSecuritySummarySpoofResponseMetaUnitJSON contains the JSON metadata for the
// struct [EmailSecuritySummarySpoofResponseMetaUnit]
type emailSecuritySummarySpoofResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummarySpoofResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummarySpoofResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummarySpoofResponseSummary0 struct {
	// A numeric string.
	NotSpoof string `json:"NOT_SPOOF,required"`
	// A numeric string.
	Spoof string                                        `json:"SPOOF,required"`
	JSON  emailSecuritySummarySpoofResponseSummary0JSON `json:"-"`
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
	// Metadata for the results.
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

// Metadata for the results.
type EmailSecuritySummaryThreatCategoryResponseMeta struct {
	ConfidenceInfo EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecuritySummaryThreatCategoryResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecuritySummaryThreatCategoryResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecuritySummaryThreatCategoryResponseMetaUnit `json:"units,required"`
	JSON  emailSecuritySummaryThreatCategoryResponseMetaJSON   `json:"-"`
}

// emailSecuritySummaryThreatCategoryResponseMetaJSON contains the JSON metadata
// for the struct [EmailSecuritySummaryThreatCategoryResponseMeta]
type emailSecuritySummaryThreatCategoryResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecuritySummaryThreatCategoryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryThreatCategoryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfo struct {
	Annotations []EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                            `json:"level,required"`
	JSON  emailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoJSON `json:"-"`
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

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                            `json:"description,required"`
	EndDate     time.Time                                                                         `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                       `json:"isInstantaneous,required"`
	LinkedURL       string                                                                     `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                  `json:"startDate,required" format:"date-time"`
	JSON            emailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotation]
type emailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailSecuritySummaryThreatCategoryResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
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

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecuritySummaryThreatCategoryResponseMetaNormalization string

const (
	EmailSecuritySummaryThreatCategoryResponseMetaNormalizationPercentage           EmailSecuritySummaryThreatCategoryResponseMetaNormalization = "PERCENTAGE"
	EmailSecuritySummaryThreatCategoryResponseMetaNormalizationMin0Max              EmailSecuritySummaryThreatCategoryResponseMetaNormalization = "MIN0_MAX"
	EmailSecuritySummaryThreatCategoryResponseMetaNormalizationMinMax               EmailSecuritySummaryThreatCategoryResponseMetaNormalization = "MIN_MAX"
	EmailSecuritySummaryThreatCategoryResponseMetaNormalizationRawValues            EmailSecuritySummaryThreatCategoryResponseMetaNormalization = "RAW_VALUES"
	EmailSecuritySummaryThreatCategoryResponseMetaNormalizationPercentageChange     EmailSecuritySummaryThreatCategoryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecuritySummaryThreatCategoryResponseMetaNormalizationRollingAverage       EmailSecuritySummaryThreatCategoryResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecuritySummaryThreatCategoryResponseMetaNormalizationOverlappedPercentage EmailSecuritySummaryThreatCategoryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecuritySummaryThreatCategoryResponseMetaNormalizationRatio                EmailSecuritySummaryThreatCategoryResponseMetaNormalization = "RATIO"
)

func (r EmailSecuritySummaryThreatCategoryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryThreatCategoryResponseMetaNormalizationPercentage, EmailSecuritySummaryThreatCategoryResponseMetaNormalizationMin0Max, EmailSecuritySummaryThreatCategoryResponseMetaNormalizationMinMax, EmailSecuritySummaryThreatCategoryResponseMetaNormalizationRawValues, EmailSecuritySummaryThreatCategoryResponseMetaNormalizationPercentageChange, EmailSecuritySummaryThreatCategoryResponseMetaNormalizationRollingAverage, EmailSecuritySummaryThreatCategoryResponseMetaNormalizationOverlappedPercentage, EmailSecuritySummaryThreatCategoryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecuritySummaryThreatCategoryResponseMetaUnit struct {
	Name  string                                                 `json:"name,required"`
	Value string                                                 `json:"value,required"`
	JSON  emailSecuritySummaryThreatCategoryResponseMetaUnitJSON `json:"-"`
}

// emailSecuritySummaryThreatCategoryResponseMetaUnitJSON contains the JSON
// metadata for the struct [EmailSecuritySummaryThreatCategoryResponseMetaUnit]
type emailSecuritySummaryThreatCategoryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryThreatCategoryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryThreatCategoryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryThreatCategoryResponseSummary0 struct {
	// A numeric string.
	BrandImpersonation string `json:"BrandImpersonation,required"`
	// A numeric string.
	CredentialHarvester string `json:"CredentialHarvester,required"`
	// A numeric string.
	IdentityDeception string `json:"IdentityDeception,required"`
	// A numeric string.
	Link string                                                 `json:"Link,required"`
	JSON emailSecuritySummaryThreatCategoryResponseSummary0JSON `json:"-"`
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
	// Metadata for the results.
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

// Metadata for the results.
type EmailSecuritySummaryTLSVersionResponseMeta struct {
	ConfidenceInfo EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailSecuritySummaryTLSVersionResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailSecuritySummaryTLSVersionResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailSecuritySummaryTLSVersionResponseMetaUnit `json:"units,required"`
	JSON  emailSecuritySummaryTLSVersionResponseMetaJSON   `json:"-"`
}

// emailSecuritySummaryTLSVersionResponseMetaJSON contains the JSON metadata for
// the struct [EmailSecuritySummaryTLSVersionResponseMeta]
type emailSecuritySummaryTLSVersionResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailSecuritySummaryTLSVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryTLSVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfo struct {
	Annotations []EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                        `json:"level,required"`
	JSON  emailSecuritySummaryTLSVersionResponseMetaConfidenceInfoJSON `json:"-"`
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

// Annotation associated with the result (e.g. outage or other type of event).
type EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                        `json:"description,required"`
	EndDate     time.Time                                                                     `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                              `json:"startDate,required" format:"date-time"`
	JSON            emailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotation]
type emailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailSecuritySummaryTLSVersionResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
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

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailSecuritySummaryTLSVersionResponseMetaNormalization string

const (
	EmailSecuritySummaryTLSVersionResponseMetaNormalizationPercentage           EmailSecuritySummaryTLSVersionResponseMetaNormalization = "PERCENTAGE"
	EmailSecuritySummaryTLSVersionResponseMetaNormalizationMin0Max              EmailSecuritySummaryTLSVersionResponseMetaNormalization = "MIN0_MAX"
	EmailSecuritySummaryTLSVersionResponseMetaNormalizationMinMax               EmailSecuritySummaryTLSVersionResponseMetaNormalization = "MIN_MAX"
	EmailSecuritySummaryTLSVersionResponseMetaNormalizationRawValues            EmailSecuritySummaryTLSVersionResponseMetaNormalization = "RAW_VALUES"
	EmailSecuritySummaryTLSVersionResponseMetaNormalizationPercentageChange     EmailSecuritySummaryTLSVersionResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailSecuritySummaryTLSVersionResponseMetaNormalizationRollingAverage       EmailSecuritySummaryTLSVersionResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailSecuritySummaryTLSVersionResponseMetaNormalizationOverlappedPercentage EmailSecuritySummaryTLSVersionResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailSecuritySummaryTLSVersionResponseMetaNormalizationRatio                EmailSecuritySummaryTLSVersionResponseMetaNormalization = "RATIO"
)

func (r EmailSecuritySummaryTLSVersionResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryTLSVersionResponseMetaNormalizationPercentage, EmailSecuritySummaryTLSVersionResponseMetaNormalizationMin0Max, EmailSecuritySummaryTLSVersionResponseMetaNormalizationMinMax, EmailSecuritySummaryTLSVersionResponseMetaNormalizationRawValues, EmailSecuritySummaryTLSVersionResponseMetaNormalizationPercentageChange, EmailSecuritySummaryTLSVersionResponseMetaNormalizationRollingAverage, EmailSecuritySummaryTLSVersionResponseMetaNormalizationOverlappedPercentage, EmailSecuritySummaryTLSVersionResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailSecuritySummaryTLSVersionResponseMetaUnit struct {
	Name  string                                             `json:"name,required"`
	Value string                                             `json:"value,required"`
	JSON  emailSecuritySummaryTLSVersionResponseMetaUnitJSON `json:"-"`
}

// emailSecuritySummaryTLSVersionResponseMetaUnitJSON contains the JSON metadata
// for the struct [EmailSecuritySummaryTLSVersionResponseMetaUnit]
type emailSecuritySummaryTLSVersionResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecuritySummaryTLSVersionResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecuritySummaryTLSVersionResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailSecuritySummaryTLSVersionResponseSummary0 struct {
	// A numeric string.
	TLS1_0 string `json:"TLS 1.0,required"`
	// A numeric string.
	TLS1_1 string `json:"TLS 1.1,required"`
	// A numeric string.
	TLS1_2 string `json:"TLS 1.2,required"`
	// A numeric string.
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
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecuritySummaryARCParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecuritySummaryARCParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecuritySummaryARCParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecuritySummaryARCParamsSPF] `query:"spf"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecuritySummaryARCParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecuritySummaryARCParams]'s query parameters as
// `url.Values`.
func (r EmailSecuritySummaryARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailSecuritySummaryARCParamsDKIM string

const (
	EmailSecuritySummaryARCParamsDKIMPass EmailSecuritySummaryARCParamsDKIM = "PASS"
	EmailSecuritySummaryARCParamsDKIMNone EmailSecuritySummaryARCParamsDKIM = "NONE"
	EmailSecuritySummaryARCParamsDKIMFail EmailSecuritySummaryARCParamsDKIM = "FAIL"
)

func (r EmailSecuritySummaryARCParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryARCParamsDKIMPass, EmailSecuritySummaryARCParamsDKIMNone, EmailSecuritySummaryARCParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecuritySummaryARCParamsDMARC string

const (
	EmailSecuritySummaryARCParamsDMARCPass EmailSecuritySummaryARCParamsDMARC = "PASS"
	EmailSecuritySummaryARCParamsDMARCNone EmailSecuritySummaryARCParamsDMARC = "NONE"
	EmailSecuritySummaryARCParamsDMARCFail EmailSecuritySummaryARCParamsDMARC = "FAIL"
)

func (r EmailSecuritySummaryARCParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryARCParamsDMARCPass, EmailSecuritySummaryARCParamsDMARCNone, EmailSecuritySummaryARCParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecuritySummaryARCParamsFormat string

const (
	EmailSecuritySummaryARCParamsFormatJson EmailSecuritySummaryARCParamsFormat = "JSON"
	EmailSecuritySummaryARCParamsFormatCsv  EmailSecuritySummaryARCParamsFormat = "CSV"
)

func (r EmailSecuritySummaryARCParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryARCParamsFormatJson, EmailSecuritySummaryARCParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecuritySummaryARCParamsSPF string

const (
	EmailSecuritySummaryARCParamsSPFPass EmailSecuritySummaryARCParamsSPF = "PASS"
	EmailSecuritySummaryARCParamsSPFNone EmailSecuritySummaryARCParamsSPF = "NONE"
	EmailSecuritySummaryARCParamsSPFFail EmailSecuritySummaryARCParamsSPF = "FAIL"
)

func (r EmailSecuritySummaryARCParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryARCParamsSPFPass, EmailSecuritySummaryARCParamsSPFNone, EmailSecuritySummaryARCParamsSPFFail:
		return true
	}
	return false
}

type EmailSecuritySummaryARCParamsTLSVersion string

const (
	EmailSecuritySummaryARCParamsTLSVersionTlSv1_0 EmailSecuritySummaryARCParamsTLSVersion = "TLSv1_0"
	EmailSecuritySummaryARCParamsTLSVersionTlSv1_1 EmailSecuritySummaryARCParamsTLSVersion = "TLSv1_1"
	EmailSecuritySummaryARCParamsTLSVersionTlSv1_2 EmailSecuritySummaryARCParamsTLSVersion = "TLSv1_2"
	EmailSecuritySummaryARCParamsTLSVersionTlSv1_3 EmailSecuritySummaryARCParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecuritySummaryARCParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryARCParamsTLSVersionTlSv1_0, EmailSecuritySummaryARCParamsTLSVersionTlSv1_1, EmailSecuritySummaryARCParamsTLSVersionTlSv1_2, EmailSecuritySummaryARCParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

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
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecuritySummaryDKIMParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecuritySummaryDKIMParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecuritySummaryDKIMParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecuritySummaryDKIMParamsSPF] `query:"spf"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecuritySummaryDKIMParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecuritySummaryDKIMParams]'s query parameters as
// `url.Values`.
func (r EmailSecuritySummaryDKIMParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailSecuritySummaryDKIMParamsARC string

const (
	EmailSecuritySummaryDKIMParamsARCPass EmailSecuritySummaryDKIMParamsARC = "PASS"
	EmailSecuritySummaryDKIMParamsARCNone EmailSecuritySummaryDKIMParamsARC = "NONE"
	EmailSecuritySummaryDKIMParamsARCFail EmailSecuritySummaryDKIMParamsARC = "FAIL"
)

func (r EmailSecuritySummaryDKIMParamsARC) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryDKIMParamsARCPass, EmailSecuritySummaryDKIMParamsARCNone, EmailSecuritySummaryDKIMParamsARCFail:
		return true
	}
	return false
}

type EmailSecuritySummaryDKIMParamsDMARC string

const (
	EmailSecuritySummaryDKIMParamsDMARCPass EmailSecuritySummaryDKIMParamsDMARC = "PASS"
	EmailSecuritySummaryDKIMParamsDMARCNone EmailSecuritySummaryDKIMParamsDMARC = "NONE"
	EmailSecuritySummaryDKIMParamsDMARCFail EmailSecuritySummaryDKIMParamsDMARC = "FAIL"
)

func (r EmailSecuritySummaryDKIMParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryDKIMParamsDMARCPass, EmailSecuritySummaryDKIMParamsDMARCNone, EmailSecuritySummaryDKIMParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecuritySummaryDKIMParamsFormat string

const (
	EmailSecuritySummaryDKIMParamsFormatJson EmailSecuritySummaryDKIMParamsFormat = "JSON"
	EmailSecuritySummaryDKIMParamsFormatCsv  EmailSecuritySummaryDKIMParamsFormat = "CSV"
)

func (r EmailSecuritySummaryDKIMParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryDKIMParamsFormatJson, EmailSecuritySummaryDKIMParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecuritySummaryDKIMParamsSPF string

const (
	EmailSecuritySummaryDKIMParamsSPFPass EmailSecuritySummaryDKIMParamsSPF = "PASS"
	EmailSecuritySummaryDKIMParamsSPFNone EmailSecuritySummaryDKIMParamsSPF = "NONE"
	EmailSecuritySummaryDKIMParamsSPFFail EmailSecuritySummaryDKIMParamsSPF = "FAIL"
)

func (r EmailSecuritySummaryDKIMParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryDKIMParamsSPFPass, EmailSecuritySummaryDKIMParamsSPFNone, EmailSecuritySummaryDKIMParamsSPFFail:
		return true
	}
	return false
}

type EmailSecuritySummaryDKIMParamsTLSVersion string

const (
	EmailSecuritySummaryDKIMParamsTLSVersionTlSv1_0 EmailSecuritySummaryDKIMParamsTLSVersion = "TLSv1_0"
	EmailSecuritySummaryDKIMParamsTLSVersionTlSv1_1 EmailSecuritySummaryDKIMParamsTLSVersion = "TLSv1_1"
	EmailSecuritySummaryDKIMParamsTLSVersionTlSv1_2 EmailSecuritySummaryDKIMParamsTLSVersion = "TLSv1_2"
	EmailSecuritySummaryDKIMParamsTLSVersionTlSv1_3 EmailSecuritySummaryDKIMParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecuritySummaryDKIMParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryDKIMParamsTLSVersionTlSv1_0, EmailSecuritySummaryDKIMParamsTLSVersionTlSv1_1, EmailSecuritySummaryDKIMParamsTLSVersionTlSv1_2, EmailSecuritySummaryDKIMParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

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
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecuritySummaryDMARCParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecuritySummaryDMARCParamsDKIM] `query:"dkim"`
	// Format in which results will be returned.
	Format param.Field[EmailSecuritySummaryDMARCParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecuritySummaryDMARCParamsSPF] `query:"spf"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecuritySummaryDMARCParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecuritySummaryDMARCParams]'s query parameters as
// `url.Values`.
func (r EmailSecuritySummaryDMARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailSecuritySummaryDMARCParamsARC string

const (
	EmailSecuritySummaryDMARCParamsARCPass EmailSecuritySummaryDMARCParamsARC = "PASS"
	EmailSecuritySummaryDMARCParamsARCNone EmailSecuritySummaryDMARCParamsARC = "NONE"
	EmailSecuritySummaryDMARCParamsARCFail EmailSecuritySummaryDMARCParamsARC = "FAIL"
)

func (r EmailSecuritySummaryDMARCParamsARC) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryDMARCParamsARCPass, EmailSecuritySummaryDMARCParamsARCNone, EmailSecuritySummaryDMARCParamsARCFail:
		return true
	}
	return false
}

type EmailSecuritySummaryDMARCParamsDKIM string

const (
	EmailSecuritySummaryDMARCParamsDKIMPass EmailSecuritySummaryDMARCParamsDKIM = "PASS"
	EmailSecuritySummaryDMARCParamsDKIMNone EmailSecuritySummaryDMARCParamsDKIM = "NONE"
	EmailSecuritySummaryDMARCParamsDKIMFail EmailSecuritySummaryDMARCParamsDKIM = "FAIL"
)

func (r EmailSecuritySummaryDMARCParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryDMARCParamsDKIMPass, EmailSecuritySummaryDMARCParamsDKIMNone, EmailSecuritySummaryDMARCParamsDKIMFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecuritySummaryDMARCParamsFormat string

const (
	EmailSecuritySummaryDMARCParamsFormatJson EmailSecuritySummaryDMARCParamsFormat = "JSON"
	EmailSecuritySummaryDMARCParamsFormatCsv  EmailSecuritySummaryDMARCParamsFormat = "CSV"
)

func (r EmailSecuritySummaryDMARCParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryDMARCParamsFormatJson, EmailSecuritySummaryDMARCParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecuritySummaryDMARCParamsSPF string

const (
	EmailSecuritySummaryDMARCParamsSPFPass EmailSecuritySummaryDMARCParamsSPF = "PASS"
	EmailSecuritySummaryDMARCParamsSPFNone EmailSecuritySummaryDMARCParamsSPF = "NONE"
	EmailSecuritySummaryDMARCParamsSPFFail EmailSecuritySummaryDMARCParamsSPF = "FAIL"
)

func (r EmailSecuritySummaryDMARCParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryDMARCParamsSPFPass, EmailSecuritySummaryDMARCParamsSPFNone, EmailSecuritySummaryDMARCParamsSPFFail:
		return true
	}
	return false
}

type EmailSecuritySummaryDMARCParamsTLSVersion string

const (
	EmailSecuritySummaryDMARCParamsTLSVersionTlSv1_0 EmailSecuritySummaryDMARCParamsTLSVersion = "TLSv1_0"
	EmailSecuritySummaryDMARCParamsTLSVersionTlSv1_1 EmailSecuritySummaryDMARCParamsTLSVersion = "TLSv1_1"
	EmailSecuritySummaryDMARCParamsTLSVersionTlSv1_2 EmailSecuritySummaryDMARCParamsTLSVersion = "TLSv1_2"
	EmailSecuritySummaryDMARCParamsTLSVersionTlSv1_3 EmailSecuritySummaryDMARCParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecuritySummaryDMARCParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryDMARCParamsTLSVersionTlSv1_0, EmailSecuritySummaryDMARCParamsTLSVersionTlSv1_1, EmailSecuritySummaryDMARCParamsTLSVersionTlSv1_2, EmailSecuritySummaryDMARCParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

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
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecuritySummaryMaliciousParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecuritySummaryMaliciousParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecuritySummaryMaliciousParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecuritySummaryMaliciousParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecuritySummaryMaliciousParamsSPF] `query:"spf"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecuritySummaryMaliciousParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecuritySummaryMaliciousParams]'s query parameters as
// `url.Values`.
func (r EmailSecuritySummaryMaliciousParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailSecuritySummaryMaliciousParamsARC string

const (
	EmailSecuritySummaryMaliciousParamsARCPass EmailSecuritySummaryMaliciousParamsARC = "PASS"
	EmailSecuritySummaryMaliciousParamsARCNone EmailSecuritySummaryMaliciousParamsARC = "NONE"
	EmailSecuritySummaryMaliciousParamsARCFail EmailSecuritySummaryMaliciousParamsARC = "FAIL"
)

func (r EmailSecuritySummaryMaliciousParamsARC) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryMaliciousParamsARCPass, EmailSecuritySummaryMaliciousParamsARCNone, EmailSecuritySummaryMaliciousParamsARCFail:
		return true
	}
	return false
}

type EmailSecuritySummaryMaliciousParamsDKIM string

const (
	EmailSecuritySummaryMaliciousParamsDKIMPass EmailSecuritySummaryMaliciousParamsDKIM = "PASS"
	EmailSecuritySummaryMaliciousParamsDKIMNone EmailSecuritySummaryMaliciousParamsDKIM = "NONE"
	EmailSecuritySummaryMaliciousParamsDKIMFail EmailSecuritySummaryMaliciousParamsDKIM = "FAIL"
)

func (r EmailSecuritySummaryMaliciousParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryMaliciousParamsDKIMPass, EmailSecuritySummaryMaliciousParamsDKIMNone, EmailSecuritySummaryMaliciousParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecuritySummaryMaliciousParamsDMARC string

const (
	EmailSecuritySummaryMaliciousParamsDMARCPass EmailSecuritySummaryMaliciousParamsDMARC = "PASS"
	EmailSecuritySummaryMaliciousParamsDMARCNone EmailSecuritySummaryMaliciousParamsDMARC = "NONE"
	EmailSecuritySummaryMaliciousParamsDMARCFail EmailSecuritySummaryMaliciousParamsDMARC = "FAIL"
)

func (r EmailSecuritySummaryMaliciousParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryMaliciousParamsDMARCPass, EmailSecuritySummaryMaliciousParamsDMARCNone, EmailSecuritySummaryMaliciousParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecuritySummaryMaliciousParamsFormat string

const (
	EmailSecuritySummaryMaliciousParamsFormatJson EmailSecuritySummaryMaliciousParamsFormat = "JSON"
	EmailSecuritySummaryMaliciousParamsFormatCsv  EmailSecuritySummaryMaliciousParamsFormat = "CSV"
)

func (r EmailSecuritySummaryMaliciousParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryMaliciousParamsFormatJson, EmailSecuritySummaryMaliciousParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecuritySummaryMaliciousParamsSPF string

const (
	EmailSecuritySummaryMaliciousParamsSPFPass EmailSecuritySummaryMaliciousParamsSPF = "PASS"
	EmailSecuritySummaryMaliciousParamsSPFNone EmailSecuritySummaryMaliciousParamsSPF = "NONE"
	EmailSecuritySummaryMaliciousParamsSPFFail EmailSecuritySummaryMaliciousParamsSPF = "FAIL"
)

func (r EmailSecuritySummaryMaliciousParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryMaliciousParamsSPFPass, EmailSecuritySummaryMaliciousParamsSPFNone, EmailSecuritySummaryMaliciousParamsSPFFail:
		return true
	}
	return false
}

type EmailSecuritySummaryMaliciousParamsTLSVersion string

const (
	EmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_0 EmailSecuritySummaryMaliciousParamsTLSVersion = "TLSv1_0"
	EmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_1 EmailSecuritySummaryMaliciousParamsTLSVersion = "TLSv1_1"
	EmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_2 EmailSecuritySummaryMaliciousParamsTLSVersion = "TLSv1_2"
	EmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_3 EmailSecuritySummaryMaliciousParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecuritySummaryMaliciousParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_0, EmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_1, EmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_2, EmailSecuritySummaryMaliciousParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

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
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecuritySummarySpamParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecuritySummarySpamParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecuritySummarySpamParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecuritySummarySpamParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecuritySummarySpamParamsSPF] `query:"spf"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecuritySummarySpamParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecuritySummarySpamParams]'s query parameters as
// `url.Values`.
func (r EmailSecuritySummarySpamParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailSecuritySummarySpamParamsARC string

const (
	EmailSecuritySummarySpamParamsARCPass EmailSecuritySummarySpamParamsARC = "PASS"
	EmailSecuritySummarySpamParamsARCNone EmailSecuritySummarySpamParamsARC = "NONE"
	EmailSecuritySummarySpamParamsARCFail EmailSecuritySummarySpamParamsARC = "FAIL"
)

func (r EmailSecuritySummarySpamParamsARC) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySpamParamsARCPass, EmailSecuritySummarySpamParamsARCNone, EmailSecuritySummarySpamParamsARCFail:
		return true
	}
	return false
}

type EmailSecuritySummarySpamParamsDKIM string

const (
	EmailSecuritySummarySpamParamsDKIMPass EmailSecuritySummarySpamParamsDKIM = "PASS"
	EmailSecuritySummarySpamParamsDKIMNone EmailSecuritySummarySpamParamsDKIM = "NONE"
	EmailSecuritySummarySpamParamsDKIMFail EmailSecuritySummarySpamParamsDKIM = "FAIL"
)

func (r EmailSecuritySummarySpamParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySpamParamsDKIMPass, EmailSecuritySummarySpamParamsDKIMNone, EmailSecuritySummarySpamParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecuritySummarySpamParamsDMARC string

const (
	EmailSecuritySummarySpamParamsDMARCPass EmailSecuritySummarySpamParamsDMARC = "PASS"
	EmailSecuritySummarySpamParamsDMARCNone EmailSecuritySummarySpamParamsDMARC = "NONE"
	EmailSecuritySummarySpamParamsDMARCFail EmailSecuritySummarySpamParamsDMARC = "FAIL"
)

func (r EmailSecuritySummarySpamParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySpamParamsDMARCPass, EmailSecuritySummarySpamParamsDMARCNone, EmailSecuritySummarySpamParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecuritySummarySpamParamsFormat string

const (
	EmailSecuritySummarySpamParamsFormatJson EmailSecuritySummarySpamParamsFormat = "JSON"
	EmailSecuritySummarySpamParamsFormatCsv  EmailSecuritySummarySpamParamsFormat = "CSV"
)

func (r EmailSecuritySummarySpamParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySpamParamsFormatJson, EmailSecuritySummarySpamParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecuritySummarySpamParamsSPF string

const (
	EmailSecuritySummarySpamParamsSPFPass EmailSecuritySummarySpamParamsSPF = "PASS"
	EmailSecuritySummarySpamParamsSPFNone EmailSecuritySummarySpamParamsSPF = "NONE"
	EmailSecuritySummarySpamParamsSPFFail EmailSecuritySummarySpamParamsSPF = "FAIL"
)

func (r EmailSecuritySummarySpamParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySpamParamsSPFPass, EmailSecuritySummarySpamParamsSPFNone, EmailSecuritySummarySpamParamsSPFFail:
		return true
	}
	return false
}

type EmailSecuritySummarySpamParamsTLSVersion string

const (
	EmailSecuritySummarySpamParamsTLSVersionTlSv1_0 EmailSecuritySummarySpamParamsTLSVersion = "TLSv1_0"
	EmailSecuritySummarySpamParamsTLSVersionTlSv1_1 EmailSecuritySummarySpamParamsTLSVersion = "TLSv1_1"
	EmailSecuritySummarySpamParamsTLSVersionTlSv1_2 EmailSecuritySummarySpamParamsTLSVersion = "TLSv1_2"
	EmailSecuritySummarySpamParamsTLSVersionTlSv1_3 EmailSecuritySummarySpamParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecuritySummarySpamParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySpamParamsTLSVersionTlSv1_0, EmailSecuritySummarySpamParamsTLSVersionTlSv1_1, EmailSecuritySummarySpamParamsTLSVersionTlSv1_2, EmailSecuritySummarySpamParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

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
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecuritySummarySPFParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecuritySummarySPFParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecuritySummarySPFParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecuritySummarySPFParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecuritySummarySPFParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecuritySummarySPFParams]'s query parameters as
// `url.Values`.
func (r EmailSecuritySummarySPFParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailSecuritySummarySPFParamsARC string

const (
	EmailSecuritySummarySPFParamsARCPass EmailSecuritySummarySPFParamsARC = "PASS"
	EmailSecuritySummarySPFParamsARCNone EmailSecuritySummarySPFParamsARC = "NONE"
	EmailSecuritySummarySPFParamsARCFail EmailSecuritySummarySPFParamsARC = "FAIL"
)

func (r EmailSecuritySummarySPFParamsARC) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySPFParamsARCPass, EmailSecuritySummarySPFParamsARCNone, EmailSecuritySummarySPFParamsARCFail:
		return true
	}
	return false
}

type EmailSecuritySummarySPFParamsDKIM string

const (
	EmailSecuritySummarySPFParamsDKIMPass EmailSecuritySummarySPFParamsDKIM = "PASS"
	EmailSecuritySummarySPFParamsDKIMNone EmailSecuritySummarySPFParamsDKIM = "NONE"
	EmailSecuritySummarySPFParamsDKIMFail EmailSecuritySummarySPFParamsDKIM = "FAIL"
)

func (r EmailSecuritySummarySPFParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySPFParamsDKIMPass, EmailSecuritySummarySPFParamsDKIMNone, EmailSecuritySummarySPFParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecuritySummarySPFParamsDMARC string

const (
	EmailSecuritySummarySPFParamsDMARCPass EmailSecuritySummarySPFParamsDMARC = "PASS"
	EmailSecuritySummarySPFParamsDMARCNone EmailSecuritySummarySPFParamsDMARC = "NONE"
	EmailSecuritySummarySPFParamsDMARCFail EmailSecuritySummarySPFParamsDMARC = "FAIL"
)

func (r EmailSecuritySummarySPFParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySPFParamsDMARCPass, EmailSecuritySummarySPFParamsDMARCNone, EmailSecuritySummarySPFParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecuritySummarySPFParamsFormat string

const (
	EmailSecuritySummarySPFParamsFormatJson EmailSecuritySummarySPFParamsFormat = "JSON"
	EmailSecuritySummarySPFParamsFormatCsv  EmailSecuritySummarySPFParamsFormat = "CSV"
)

func (r EmailSecuritySummarySPFParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySPFParamsFormatJson, EmailSecuritySummarySPFParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecuritySummarySPFParamsTLSVersion string

const (
	EmailSecuritySummarySPFParamsTLSVersionTlSv1_0 EmailSecuritySummarySPFParamsTLSVersion = "TLSv1_0"
	EmailSecuritySummarySPFParamsTLSVersionTlSv1_1 EmailSecuritySummarySPFParamsTLSVersion = "TLSv1_1"
	EmailSecuritySummarySPFParamsTLSVersionTlSv1_2 EmailSecuritySummarySPFParamsTLSVersion = "TLSv1_2"
	EmailSecuritySummarySPFParamsTLSVersionTlSv1_3 EmailSecuritySummarySPFParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecuritySummarySPFParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySPFParamsTLSVersionTlSv1_0, EmailSecuritySummarySPFParamsTLSVersionTlSv1_1, EmailSecuritySummarySPFParamsTLSVersionTlSv1_2, EmailSecuritySummarySPFParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

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
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecuritySummarySpoofParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecuritySummarySpoofParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecuritySummarySpoofParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecuritySummarySpoofParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecuritySummarySpoofParamsSPF] `query:"spf"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecuritySummarySpoofParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecuritySummarySpoofParams]'s query parameters as
// `url.Values`.
func (r EmailSecuritySummarySpoofParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailSecuritySummarySpoofParamsARC string

const (
	EmailSecuritySummarySpoofParamsARCPass EmailSecuritySummarySpoofParamsARC = "PASS"
	EmailSecuritySummarySpoofParamsARCNone EmailSecuritySummarySpoofParamsARC = "NONE"
	EmailSecuritySummarySpoofParamsARCFail EmailSecuritySummarySpoofParamsARC = "FAIL"
)

func (r EmailSecuritySummarySpoofParamsARC) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySpoofParamsARCPass, EmailSecuritySummarySpoofParamsARCNone, EmailSecuritySummarySpoofParamsARCFail:
		return true
	}
	return false
}

type EmailSecuritySummarySpoofParamsDKIM string

const (
	EmailSecuritySummarySpoofParamsDKIMPass EmailSecuritySummarySpoofParamsDKIM = "PASS"
	EmailSecuritySummarySpoofParamsDKIMNone EmailSecuritySummarySpoofParamsDKIM = "NONE"
	EmailSecuritySummarySpoofParamsDKIMFail EmailSecuritySummarySpoofParamsDKIM = "FAIL"
)

func (r EmailSecuritySummarySpoofParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySpoofParamsDKIMPass, EmailSecuritySummarySpoofParamsDKIMNone, EmailSecuritySummarySpoofParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecuritySummarySpoofParamsDMARC string

const (
	EmailSecuritySummarySpoofParamsDMARCPass EmailSecuritySummarySpoofParamsDMARC = "PASS"
	EmailSecuritySummarySpoofParamsDMARCNone EmailSecuritySummarySpoofParamsDMARC = "NONE"
	EmailSecuritySummarySpoofParamsDMARCFail EmailSecuritySummarySpoofParamsDMARC = "FAIL"
)

func (r EmailSecuritySummarySpoofParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySpoofParamsDMARCPass, EmailSecuritySummarySpoofParamsDMARCNone, EmailSecuritySummarySpoofParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecuritySummarySpoofParamsFormat string

const (
	EmailSecuritySummarySpoofParamsFormatJson EmailSecuritySummarySpoofParamsFormat = "JSON"
	EmailSecuritySummarySpoofParamsFormatCsv  EmailSecuritySummarySpoofParamsFormat = "CSV"
)

func (r EmailSecuritySummarySpoofParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySpoofParamsFormatJson, EmailSecuritySummarySpoofParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecuritySummarySpoofParamsSPF string

const (
	EmailSecuritySummarySpoofParamsSPFPass EmailSecuritySummarySpoofParamsSPF = "PASS"
	EmailSecuritySummarySpoofParamsSPFNone EmailSecuritySummarySpoofParamsSPF = "NONE"
	EmailSecuritySummarySpoofParamsSPFFail EmailSecuritySummarySpoofParamsSPF = "FAIL"
)

func (r EmailSecuritySummarySpoofParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySpoofParamsSPFPass, EmailSecuritySummarySpoofParamsSPFNone, EmailSecuritySummarySpoofParamsSPFFail:
		return true
	}
	return false
}

type EmailSecuritySummarySpoofParamsTLSVersion string

const (
	EmailSecuritySummarySpoofParamsTLSVersionTlSv1_0 EmailSecuritySummarySpoofParamsTLSVersion = "TLSv1_0"
	EmailSecuritySummarySpoofParamsTLSVersionTlSv1_1 EmailSecuritySummarySpoofParamsTLSVersion = "TLSv1_1"
	EmailSecuritySummarySpoofParamsTLSVersionTlSv1_2 EmailSecuritySummarySpoofParamsTLSVersion = "TLSv1_2"
	EmailSecuritySummarySpoofParamsTLSVersionTlSv1_3 EmailSecuritySummarySpoofParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecuritySummarySpoofParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecuritySummarySpoofParamsTLSVersionTlSv1_0, EmailSecuritySummarySpoofParamsTLSVersionTlSv1_1, EmailSecuritySummarySpoofParamsTLSVersionTlSv1_2, EmailSecuritySummarySpoofParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

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
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecuritySummaryThreatCategoryParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecuritySummaryThreatCategoryParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecuritySummaryThreatCategoryParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecuritySummaryThreatCategoryParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecuritySummaryThreatCategoryParamsSPF] `query:"spf"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]EmailSecuritySummaryThreatCategoryParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [EmailSecuritySummaryThreatCategoryParams]'s query
// parameters as `url.Values`.
func (r EmailSecuritySummaryThreatCategoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailSecuritySummaryThreatCategoryParamsARC string

const (
	EmailSecuritySummaryThreatCategoryParamsARCPass EmailSecuritySummaryThreatCategoryParamsARC = "PASS"
	EmailSecuritySummaryThreatCategoryParamsARCNone EmailSecuritySummaryThreatCategoryParamsARC = "NONE"
	EmailSecuritySummaryThreatCategoryParamsARCFail EmailSecuritySummaryThreatCategoryParamsARC = "FAIL"
)

func (r EmailSecuritySummaryThreatCategoryParamsARC) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryThreatCategoryParamsARCPass, EmailSecuritySummaryThreatCategoryParamsARCNone, EmailSecuritySummaryThreatCategoryParamsARCFail:
		return true
	}
	return false
}

type EmailSecuritySummaryThreatCategoryParamsDKIM string

const (
	EmailSecuritySummaryThreatCategoryParamsDKIMPass EmailSecuritySummaryThreatCategoryParamsDKIM = "PASS"
	EmailSecuritySummaryThreatCategoryParamsDKIMNone EmailSecuritySummaryThreatCategoryParamsDKIM = "NONE"
	EmailSecuritySummaryThreatCategoryParamsDKIMFail EmailSecuritySummaryThreatCategoryParamsDKIM = "FAIL"
)

func (r EmailSecuritySummaryThreatCategoryParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryThreatCategoryParamsDKIMPass, EmailSecuritySummaryThreatCategoryParamsDKIMNone, EmailSecuritySummaryThreatCategoryParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecuritySummaryThreatCategoryParamsDMARC string

const (
	EmailSecuritySummaryThreatCategoryParamsDMARCPass EmailSecuritySummaryThreatCategoryParamsDMARC = "PASS"
	EmailSecuritySummaryThreatCategoryParamsDMARCNone EmailSecuritySummaryThreatCategoryParamsDMARC = "NONE"
	EmailSecuritySummaryThreatCategoryParamsDMARCFail EmailSecuritySummaryThreatCategoryParamsDMARC = "FAIL"
)

func (r EmailSecuritySummaryThreatCategoryParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryThreatCategoryParamsDMARCPass, EmailSecuritySummaryThreatCategoryParamsDMARCNone, EmailSecuritySummaryThreatCategoryParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecuritySummaryThreatCategoryParamsFormat string

const (
	EmailSecuritySummaryThreatCategoryParamsFormatJson EmailSecuritySummaryThreatCategoryParamsFormat = "JSON"
	EmailSecuritySummaryThreatCategoryParamsFormatCsv  EmailSecuritySummaryThreatCategoryParamsFormat = "CSV"
)

func (r EmailSecuritySummaryThreatCategoryParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryThreatCategoryParamsFormatJson, EmailSecuritySummaryThreatCategoryParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecuritySummaryThreatCategoryParamsSPF string

const (
	EmailSecuritySummaryThreatCategoryParamsSPFPass EmailSecuritySummaryThreatCategoryParamsSPF = "PASS"
	EmailSecuritySummaryThreatCategoryParamsSPFNone EmailSecuritySummaryThreatCategoryParamsSPF = "NONE"
	EmailSecuritySummaryThreatCategoryParamsSPFFail EmailSecuritySummaryThreatCategoryParamsSPF = "FAIL"
)

func (r EmailSecuritySummaryThreatCategoryParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryThreatCategoryParamsSPFPass, EmailSecuritySummaryThreatCategoryParamsSPFNone, EmailSecuritySummaryThreatCategoryParamsSPFFail:
		return true
	}
	return false
}

type EmailSecuritySummaryThreatCategoryParamsTLSVersion string

const (
	EmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_0 EmailSecuritySummaryThreatCategoryParamsTLSVersion = "TLSv1_0"
	EmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_1 EmailSecuritySummaryThreatCategoryParamsTLSVersion = "TLSv1_1"
	EmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_2 EmailSecuritySummaryThreatCategoryParamsTLSVersion = "TLSv1_2"
	EmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_3 EmailSecuritySummaryThreatCategoryParamsTLSVersion = "TLSv1_3"
)

func (r EmailSecuritySummaryThreatCategoryParamsTLSVersion) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_0, EmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_1, EmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_2, EmailSecuritySummaryThreatCategoryParamsTLSVersionTlSv1_3:
		return true
	}
	return false
}

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
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailSecuritySummaryTLSVersionParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailSecuritySummaryTLSVersionParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailSecuritySummaryTLSVersionParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailSecuritySummaryTLSVersionParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailSecuritySummaryTLSVersionParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailSecuritySummaryTLSVersionParams]'s query parameters as
// `url.Values`.
func (r EmailSecuritySummaryTLSVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EmailSecuritySummaryTLSVersionParamsARC string

const (
	EmailSecuritySummaryTLSVersionParamsARCPass EmailSecuritySummaryTLSVersionParamsARC = "PASS"
	EmailSecuritySummaryTLSVersionParamsARCNone EmailSecuritySummaryTLSVersionParamsARC = "NONE"
	EmailSecuritySummaryTLSVersionParamsARCFail EmailSecuritySummaryTLSVersionParamsARC = "FAIL"
)

func (r EmailSecuritySummaryTLSVersionParamsARC) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryTLSVersionParamsARCPass, EmailSecuritySummaryTLSVersionParamsARCNone, EmailSecuritySummaryTLSVersionParamsARCFail:
		return true
	}
	return false
}

type EmailSecuritySummaryTLSVersionParamsDKIM string

const (
	EmailSecuritySummaryTLSVersionParamsDKIMPass EmailSecuritySummaryTLSVersionParamsDKIM = "PASS"
	EmailSecuritySummaryTLSVersionParamsDKIMNone EmailSecuritySummaryTLSVersionParamsDKIM = "NONE"
	EmailSecuritySummaryTLSVersionParamsDKIMFail EmailSecuritySummaryTLSVersionParamsDKIM = "FAIL"
)

func (r EmailSecuritySummaryTLSVersionParamsDKIM) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryTLSVersionParamsDKIMPass, EmailSecuritySummaryTLSVersionParamsDKIMNone, EmailSecuritySummaryTLSVersionParamsDKIMFail:
		return true
	}
	return false
}

type EmailSecuritySummaryTLSVersionParamsDMARC string

const (
	EmailSecuritySummaryTLSVersionParamsDMARCPass EmailSecuritySummaryTLSVersionParamsDMARC = "PASS"
	EmailSecuritySummaryTLSVersionParamsDMARCNone EmailSecuritySummaryTLSVersionParamsDMARC = "NONE"
	EmailSecuritySummaryTLSVersionParamsDMARCFail EmailSecuritySummaryTLSVersionParamsDMARC = "FAIL"
)

func (r EmailSecuritySummaryTLSVersionParamsDMARC) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryTLSVersionParamsDMARCPass, EmailSecuritySummaryTLSVersionParamsDMARCNone, EmailSecuritySummaryTLSVersionParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailSecuritySummaryTLSVersionParamsFormat string

const (
	EmailSecuritySummaryTLSVersionParamsFormatJson EmailSecuritySummaryTLSVersionParamsFormat = "JSON"
	EmailSecuritySummaryTLSVersionParamsFormatCsv  EmailSecuritySummaryTLSVersionParamsFormat = "CSV"
)

func (r EmailSecuritySummaryTLSVersionParamsFormat) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryTLSVersionParamsFormatJson, EmailSecuritySummaryTLSVersionParamsFormatCsv:
		return true
	}
	return false
}

type EmailSecuritySummaryTLSVersionParamsSPF string

const (
	EmailSecuritySummaryTLSVersionParamsSPFPass EmailSecuritySummaryTLSVersionParamsSPF = "PASS"
	EmailSecuritySummaryTLSVersionParamsSPFNone EmailSecuritySummaryTLSVersionParamsSPF = "NONE"
	EmailSecuritySummaryTLSVersionParamsSPFFail EmailSecuritySummaryTLSVersionParamsSPF = "FAIL"
)

func (r EmailSecuritySummaryTLSVersionParamsSPF) IsKnown() bool {
	switch r {
	case EmailSecuritySummaryTLSVersionParamsSPFPass, EmailSecuritySummaryTLSVersionParamsSPFNone, EmailSecuritySummaryTLSVersionParamsSPFFail:
		return true
	}
	return false
}

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
