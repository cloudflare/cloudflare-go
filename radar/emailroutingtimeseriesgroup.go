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

// EmailRoutingTimeseriesGroupService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailRoutingTimeseriesGroupService] method instead.
type EmailRoutingTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewEmailRoutingTimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmailRoutingTimeseriesGroupService(opts ...option.RequestOption) (r *EmailRoutingTimeseriesGroupService) {
	r = &EmailRoutingTimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of emails by ARC (Authenticated Received Chain)
// validation over time.
//
// Deprecated: Use
// [Radar Email Routing Timeseries Groups By Dimension](https://developers.cloudflare.com/api/resources/radar/subresources/email/subresources/routing/methods/timeseries_groups_v2/)
// instead.
func (r *EmailRoutingTimeseriesGroupService) ARC(ctx context.Context, query EmailRoutingTimeseriesGroupARCParams, opts ...option.RequestOption) (res *EmailRoutingTimeseriesGroupARCResponse, err error) {
	var env EmailRoutingTimeseriesGroupARCResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/email/routing/timeseries_groups/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by DKIM (DomainKeys Identified Mail)
// validation over time.
//
// Deprecated: Use
// [Radar Email Routing Timeseries Groups By Dimension](https://developers.cloudflare.com/api/resources/radar/subresources/email/subresources/routing/methods/timeseries_groups_v2/)
// instead.
func (r *EmailRoutingTimeseriesGroupService) DKIM(ctx context.Context, query EmailRoutingTimeseriesGroupDKIMParams, opts ...option.RequestOption) (res *EmailRoutingTimeseriesGroupDKIMResponse, err error) {
	var env EmailRoutingTimeseriesGroupDKIMResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/email/routing/timeseries_groups/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by DMARC (Domain-based Message
// Authentication, Reporting and Conformance) validation over time.
//
// Deprecated: Use
// [Radar Email Routing Timeseries Groups By Dimension](https://developers.cloudflare.com/api/resources/radar/subresources/email/subresources/routing/methods/timeseries_groups_v2/)
// instead.
func (r *EmailRoutingTimeseriesGroupService) DMARC(ctx context.Context, query EmailRoutingTimeseriesGroupDMARCParams, opts ...option.RequestOption) (res *EmailRoutingTimeseriesGroupDMARCResponse, err error) {
	var env EmailRoutingTimeseriesGroupDMARCResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/email/routing/timeseries_groups/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by encryption status (encrypted vs.
// not-encrypted) over time.
//
// Deprecated: Use
// [Radar Email Routing Timeseries Groups By Dimension](https://developers.cloudflare.com/api/resources/radar/subresources/email/subresources/routing/methods/timeseries_groups_v2/)
// instead.
func (r *EmailRoutingTimeseriesGroupService) Encrypted(ctx context.Context, query EmailRoutingTimeseriesGroupEncryptedParams, opts ...option.RequestOption) (res *EmailRoutingTimeseriesGroupEncryptedResponse, err error) {
	var env EmailRoutingTimeseriesGroupEncryptedResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/email/routing/timeseries_groups/encrypted"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by IP version over time.
//
// Deprecated: Use
// [Radar Email Routing Timeseries Groups By Dimension](https://developers.cloudflare.com/api/resources/radar/subresources/email/subresources/routing/methods/timeseries_groups_v2/)
// instead.
func (r *EmailRoutingTimeseriesGroupService) IPVersion(ctx context.Context, query EmailRoutingTimeseriesGroupIPVersionParams, opts ...option.RequestOption) (res *EmailRoutingTimeseriesGroupIPVersionResponse, err error) {
	var env EmailRoutingTimeseriesGroupIPVersionResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/email/routing/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of emails by SPF (Sender Policy Framework) validation
// over time.
//
// Deprecated: Use
// [Radar Email Routing Timeseries Groups By Dimension](https://developers.cloudflare.com/api/resources/radar/subresources/email/subresources/routing/methods/timeseries_groups_v2/)
// instead.
func (r *EmailRoutingTimeseriesGroupService) SPF(ctx context.Context, query EmailRoutingTimeseriesGroupSPFParams, opts ...option.RequestOption) (res *EmailRoutingTimeseriesGroupSPFResponse, err error) {
	var env EmailRoutingTimeseriesGroupSPFResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/email/routing/timeseries_groups/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingTimeseriesGroupARCResponse struct {
	// Metadata for the results.
	Meta   EmailRoutingTimeseriesGroupARCResponseMeta `json:"meta,required"`
	Serie0 RadarEmailSeries                           `json:"serie_0,required"`
	JSON   emailRoutingTimeseriesGroupARCResponseJSON `json:"-"`
}

// emailRoutingTimeseriesGroupARCResponseJSON contains the JSON metadata for the
// struct [EmailRoutingTimeseriesGroupARCResponse]
type emailRoutingTimeseriesGroupARCResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupARCResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingTimeseriesGroupARCResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailRoutingTimeseriesGroupARCResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingTimeseriesGroupARCResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingTimeseriesGroupARCResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingTimeseriesGroupARCResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingTimeseriesGroupARCResponseMetaJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupARCResponseMetaJSON contains the JSON metadata for
// the struct [EmailRoutingTimeseriesGroupARCResponseMeta]
type emailRoutingTimeseriesGroupARCResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupARCResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupARCResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupARCResponseMetaAggInterval string

const (
	EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalFifteenMinutes EmailRoutingTimeseriesGroupARCResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalOneHour        EmailRoutingTimeseriesGroupARCResponseMetaAggInterval = "ONE_HOUR"
	EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalOneDay         EmailRoutingTimeseriesGroupARCResponseMetaAggInterval = "ONE_DAY"
	EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalOneWeek        EmailRoutingTimeseriesGroupARCResponseMetaAggInterval = "ONE_WEEK"
	EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalOneMonth       EmailRoutingTimeseriesGroupARCResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailRoutingTimeseriesGroupARCResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalFifteenMinutes, EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalOneHour, EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalOneDay, EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalOneWeek, EmailRoutingTimeseriesGroupARCResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                        `json:"level,required"`
	JSON  emailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfo]
type emailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                        `json:"description,required"`
	EndDate     time.Time                                                                     `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                              `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotation]
type emailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailRoutingTimeseriesGroupARCResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupARCResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingTimeseriesGroupARCResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupARCResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupARCResponseMetaDateRange]
type emailRoutingTimeseriesGroupARCResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupARCResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupARCResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingTimeseriesGroupARCResponseMetaNormalization string

const (
	EmailRoutingTimeseriesGroupARCResponseMetaNormalizationPercentage           EmailRoutingTimeseriesGroupARCResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingTimeseriesGroupARCResponseMetaNormalizationMin0Max              EmailRoutingTimeseriesGroupARCResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingTimeseriesGroupARCResponseMetaNormalizationMinMax               EmailRoutingTimeseriesGroupARCResponseMetaNormalization = "MIN_MAX"
	EmailRoutingTimeseriesGroupARCResponseMetaNormalizationRawValues            EmailRoutingTimeseriesGroupARCResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingTimeseriesGroupARCResponseMetaNormalizationPercentageChange     EmailRoutingTimeseriesGroupARCResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingTimeseriesGroupARCResponseMetaNormalizationRollingAverage       EmailRoutingTimeseriesGroupARCResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingTimeseriesGroupARCResponseMetaNormalizationOverlappedPercentage EmailRoutingTimeseriesGroupARCResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingTimeseriesGroupARCResponseMetaNormalizationRatio                EmailRoutingTimeseriesGroupARCResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingTimeseriesGroupARCResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCResponseMetaNormalizationPercentage, EmailRoutingTimeseriesGroupARCResponseMetaNormalizationMin0Max, EmailRoutingTimeseriesGroupARCResponseMetaNormalizationMinMax, EmailRoutingTimeseriesGroupARCResponseMetaNormalizationRawValues, EmailRoutingTimeseriesGroupARCResponseMetaNormalizationPercentageChange, EmailRoutingTimeseriesGroupARCResponseMetaNormalizationRollingAverage, EmailRoutingTimeseriesGroupARCResponseMetaNormalizationOverlappedPercentage, EmailRoutingTimeseriesGroupARCResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupARCResponseMetaUnit struct {
	Name  string                                             `json:"name,required"`
	Value string                                             `json:"value,required"`
	JSON  emailRoutingTimeseriesGroupARCResponseMetaUnitJSON `json:"-"`
}

// emailRoutingTimeseriesGroupARCResponseMetaUnitJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupARCResponseMetaUnit]
type emailRoutingTimeseriesGroupARCResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupARCResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupARCResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupDKIMResponse struct {
	// Metadata for the results.
	Meta   EmailRoutingTimeseriesGroupDKIMResponseMeta `json:"meta,required"`
	Serie0 RadarEmailSeries                            `json:"serie_0,required"`
	JSON   emailRoutingTimeseriesGroupDKIMResponseJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDKIMResponseJSON contains the JSON metadata for the
// struct [EmailRoutingTimeseriesGroupDKIMResponse]
type emailRoutingTimeseriesGroupDKIMResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDKIMResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDKIMResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingTimeseriesGroupDKIMResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailRoutingTimeseriesGroupDKIMResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingTimeseriesGroupDKIMResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingTimeseriesGroupDKIMResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingTimeseriesGroupDKIMResponseMetaJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupDKIMResponseMetaJSON contains the JSON metadata for
// the struct [EmailRoutingTimeseriesGroupDKIMResponseMeta]
type emailRoutingTimeseriesGroupDKIMResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDKIMResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDKIMResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupDKIMResponseMetaAggInterval string

const (
	EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalFifteenMinutes EmailRoutingTimeseriesGroupDKIMResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalOneHour        EmailRoutingTimeseriesGroupDKIMResponseMetaAggInterval = "ONE_HOUR"
	EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalOneDay         EmailRoutingTimeseriesGroupDKIMResponseMetaAggInterval = "ONE_DAY"
	EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalOneWeek        EmailRoutingTimeseriesGroupDKIMResponseMetaAggInterval = "ONE_WEEK"
	EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalOneMonth       EmailRoutingTimeseriesGroupDKIMResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailRoutingTimeseriesGroupDKIMResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalFifteenMinutes, EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalOneHour, EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalOneDay, EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalOneWeek, EmailRoutingTimeseriesGroupDKIMResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                         `json:"level,required"`
	JSON  emailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfo]
type emailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                         `json:"description,required"`
	EndDate     time.Time                                                                      `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                               `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotation]
type emailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailRoutingTimeseriesGroupDKIMResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDKIMResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingTimeseriesGroupDKIMResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDKIMResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupDKIMResponseMetaDateRange]
type emailRoutingTimeseriesGroupDKIMResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDKIMResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDKIMResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization string

const (
	EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationPercentage           EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationMin0Max              EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationMinMax               EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization = "MIN_MAX"
	EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationRawValues            EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationPercentageChange     EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationRollingAverage       EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationOverlappedPercentage EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationRatio                EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingTimeseriesGroupDKIMResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationPercentage, EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationMin0Max, EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationMinMax, EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationRawValues, EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationPercentageChange, EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationRollingAverage, EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationOverlappedPercentage, EmailRoutingTimeseriesGroupDKIMResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDKIMResponseMetaUnit struct {
	Name  string                                              `json:"name,required"`
	Value string                                              `json:"value,required"`
	JSON  emailRoutingTimeseriesGroupDKIMResponseMetaUnitJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDKIMResponseMetaUnitJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupDKIMResponseMetaUnit]
type emailRoutingTimeseriesGroupDKIMResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDKIMResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDKIMResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupDMARCResponse struct {
	// Metadata for the results.
	Meta   EmailRoutingTimeseriesGroupDMARCResponseMeta `json:"meta,required"`
	Serie0 RadarEmailSeries                             `json:"serie_0,required"`
	JSON   emailRoutingTimeseriesGroupDMARCResponseJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDMARCResponseJSON contains the JSON metadata for the
// struct [EmailRoutingTimeseriesGroupDMARCResponse]
type emailRoutingTimeseriesGroupDMARCResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDMARCResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDMARCResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingTimeseriesGroupDMARCResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailRoutingTimeseriesGroupDMARCResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingTimeseriesGroupDMARCResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingTimeseriesGroupDMARCResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingTimeseriesGroupDMARCResponseMetaJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupDMARCResponseMetaJSON contains the JSON metadata for
// the struct [EmailRoutingTimeseriesGroupDMARCResponseMeta]
type emailRoutingTimeseriesGroupDMARCResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDMARCResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDMARCResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupDMARCResponseMetaAggInterval string

const (
	EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalFifteenMinutes EmailRoutingTimeseriesGroupDMARCResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalOneHour        EmailRoutingTimeseriesGroupDMARCResponseMetaAggInterval = "ONE_HOUR"
	EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalOneDay         EmailRoutingTimeseriesGroupDMARCResponseMetaAggInterval = "ONE_DAY"
	EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalOneWeek        EmailRoutingTimeseriesGroupDMARCResponseMetaAggInterval = "ONE_WEEK"
	EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalOneMonth       EmailRoutingTimeseriesGroupDMARCResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailRoutingTimeseriesGroupDMARCResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalFifteenMinutes, EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalOneHour, EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalOneDay, EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalOneWeek, EmailRoutingTimeseriesGroupDMARCResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                          `json:"level,required"`
	JSON  emailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfo]
type emailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                          `json:"description,required"`
	EndDate     time.Time                                                                       `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                     `json:"isInstantaneous,required"`
	LinkedURL       string                                                                   `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotation]
type emailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailRoutingTimeseriesGroupDMARCResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDMARCResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingTimeseriesGroupDMARCResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDMARCResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupDMARCResponseMetaDateRange]
type emailRoutingTimeseriesGroupDMARCResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDMARCResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDMARCResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization string

const (
	EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationPercentage           EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationMin0Max              EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationMinMax               EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization = "MIN_MAX"
	EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationRawValues            EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationPercentageChange     EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationRollingAverage       EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationOverlappedPercentage EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationRatio                EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingTimeseriesGroupDMARCResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationPercentage, EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationMin0Max, EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationMinMax, EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationRawValues, EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationPercentageChange, EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationRollingAverage, EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationOverlappedPercentage, EmailRoutingTimeseriesGroupDMARCResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDMARCResponseMetaUnit struct {
	Name  string                                               `json:"name,required"`
	Value string                                               `json:"value,required"`
	JSON  emailRoutingTimeseriesGroupDMARCResponseMetaUnitJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDMARCResponseMetaUnitJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupDMARCResponseMetaUnit]
type emailRoutingTimeseriesGroupDMARCResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDMARCResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDMARCResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupEncryptedResponse struct {
	// Metadata for the results.
	Meta   EmailRoutingTimeseriesGroupEncryptedResponseMeta   `json:"meta,required"`
	Serie0 EmailRoutingTimeseriesGroupEncryptedResponseSerie0 `json:"serie_0,required"`
	JSON   emailRoutingTimeseriesGroupEncryptedResponseJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupEncryptedResponseJSON contains the JSON metadata for
// the struct [EmailRoutingTimeseriesGroupEncryptedResponse]
type emailRoutingTimeseriesGroupEncryptedResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupEncryptedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupEncryptedResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingTimeseriesGroupEncryptedResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailRoutingTimeseriesGroupEncryptedResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingTimeseriesGroupEncryptedResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingTimeseriesGroupEncryptedResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingTimeseriesGroupEncryptedResponseMetaJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupEncryptedResponseMetaJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupEncryptedResponseMeta]
type emailRoutingTimeseriesGroupEncryptedResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupEncryptedResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupEncryptedResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupEncryptedResponseMetaAggInterval string

const (
	EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalFifteenMinutes EmailRoutingTimeseriesGroupEncryptedResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalOneHour        EmailRoutingTimeseriesGroupEncryptedResponseMetaAggInterval = "ONE_HOUR"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalOneDay         EmailRoutingTimeseriesGroupEncryptedResponseMetaAggInterval = "ONE_DAY"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalOneWeek        EmailRoutingTimeseriesGroupEncryptedResponseMetaAggInterval = "ONE_WEEK"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalOneMonth       EmailRoutingTimeseriesGroupEncryptedResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailRoutingTimeseriesGroupEncryptedResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalFifteenMinutes, EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalOneHour, EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalOneDay, EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalOneWeek, EmailRoutingTimeseriesGroupEncryptedResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                              `json:"level,required"`
	JSON  emailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfo]
type emailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                              `json:"description,required"`
	EndDate     time.Time                                                                           `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                    `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotation]
type emailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailRoutingTimeseriesGroupEncryptedResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupEncryptedResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingTimeseriesGroupEncryptedResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupEncryptedResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [EmailRoutingTimeseriesGroupEncryptedResponseMetaDateRange]
type emailRoutingTimeseriesGroupEncryptedResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupEncryptedResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupEncryptedResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization string

const (
	EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationPercentage           EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationMin0Max              EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationMinMax               EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization = "MIN_MAX"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationRawValues            EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationPercentageChange     EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationRollingAverage       EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationOverlappedPercentage EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationRatio                EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationPercentage, EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationMin0Max, EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationMinMax, EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationRawValues, EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationPercentageChange, EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationRollingAverage, EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationOverlappedPercentage, EmailRoutingTimeseriesGroupEncryptedResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupEncryptedResponseMetaUnit struct {
	Name  string                                                   `json:"name,required"`
	Value string                                                   `json:"value,required"`
	JSON  emailRoutingTimeseriesGroupEncryptedResponseMetaUnitJSON `json:"-"`
}

// emailRoutingTimeseriesGroupEncryptedResponseMetaUnitJSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupEncryptedResponseMetaUnit]
type emailRoutingTimeseriesGroupEncryptedResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupEncryptedResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupEncryptedResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupEncryptedResponseSerie0 struct {
	Encrypted    []string                                               `json:"ENCRYPTED,required"`
	NotEncrypted []string                                               `json:"NOT_ENCRYPTED,required"`
	JSON         emailRoutingTimeseriesGroupEncryptedResponseSerie0JSON `json:"-"`
}

// emailRoutingTimeseriesGroupEncryptedResponseSerie0JSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupEncryptedResponseSerie0]
type emailRoutingTimeseriesGroupEncryptedResponseSerie0JSON struct {
	Encrypted    apijson.Field
	NotEncrypted apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupEncryptedResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupEncryptedResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupIPVersionResponse struct {
	// Metadata for the results.
	Meta   EmailRoutingTimeseriesGroupIPVersionResponseMeta   `json:"meta,required"`
	Serie0 EmailRoutingTimeseriesGroupIPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   emailRoutingTimeseriesGroupIPVersionResponseJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupIPVersionResponseJSON contains the JSON metadata for
// the struct [EmailRoutingTimeseriesGroupIPVersionResponse]
type emailRoutingTimeseriesGroupIPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingTimeseriesGroupIPVersionResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailRoutingTimeseriesGroupIPVersionResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingTimeseriesGroupIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingTimeseriesGroupIPVersionResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingTimeseriesGroupIPVersionResponseMetaJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupIPVersionResponseMetaJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupIPVersionResponseMeta]
type emailRoutingTimeseriesGroupIPVersionResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupIPVersionResponseMetaAggInterval string

const (
	EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalFifteenMinutes EmailRoutingTimeseriesGroupIPVersionResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalOneHour        EmailRoutingTimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_HOUR"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalOneDay         EmailRoutingTimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_DAY"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalOneWeek        EmailRoutingTimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_WEEK"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalOneMonth       EmailRoutingTimeseriesGroupIPVersionResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailRoutingTimeseriesGroupIPVersionResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalFifteenMinutes, EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalOneHour, EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalOneDay, EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalOneWeek, EmailRoutingTimeseriesGroupIPVersionResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                              `json:"level,required"`
	JSON  emailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfo]
type emailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                              `json:"description,required"`
	EndDate     time.Time                                                                           `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                    `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation]
type emailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailRoutingTimeseriesGroupIPVersionResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingTimeseriesGroupIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupIPVersionResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [EmailRoutingTimeseriesGroupIPVersionResponseMetaDateRange]
type emailRoutingTimeseriesGroupIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupIPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization string

const (
	EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationPercentage           EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationMin0Max              EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationMinMax               EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization = "MIN_MAX"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationRawValues            EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationPercentageChange     EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationRollingAverage       EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationOverlappedPercentage EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationRatio                EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationPercentage, EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationMin0Max, EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationMinMax, EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationRawValues, EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationPercentageChange, EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationRollingAverage, EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationOverlappedPercentage, EmailRoutingTimeseriesGroupIPVersionResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupIPVersionResponseMetaUnit struct {
	Name  string                                                   `json:"name,required"`
	Value string                                                   `json:"value,required"`
	JSON  emailRoutingTimeseriesGroupIPVersionResponseMetaUnitJSON `json:"-"`
}

// emailRoutingTimeseriesGroupIPVersionResponseMetaUnitJSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupIPVersionResponseMetaUnit]
type emailRoutingTimeseriesGroupIPVersionResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupIPVersionResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupIPVersionResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupIPVersionResponseSerie0 struct {
	IPv4 []string                                               `json:"IPv4,required"`
	IPv6 []string                                               `json:"IPv6,required"`
	JSON emailRoutingTimeseriesGroupIPVersionResponseSerie0JSON `json:"-"`
}

// emailRoutingTimeseriesGroupIPVersionResponseSerie0JSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupIPVersionResponseSerie0]
type emailRoutingTimeseriesGroupIPVersionResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupIPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupIPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupSPFResponse struct {
	// Metadata for the results.
	Meta   EmailRoutingTimeseriesGroupSPFResponseMeta `json:"meta,required"`
	Serie0 RadarEmailSeries                           `json:"serie_0,required"`
	JSON   emailRoutingTimeseriesGroupSPFResponseJSON `json:"-"`
}

// emailRoutingTimeseriesGroupSPFResponseJSON contains the JSON metadata for the
// struct [EmailRoutingTimeseriesGroupSPFResponse]
type emailRoutingTimeseriesGroupSPFResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupSPFResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupSPFResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type EmailRoutingTimeseriesGroupSPFResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    EmailRoutingTimeseriesGroupSPFResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []EmailRoutingTimeseriesGroupSPFResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization EmailRoutingTimeseriesGroupSPFResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []EmailRoutingTimeseriesGroupSPFResponseMetaUnit `json:"units,required"`
	JSON  emailRoutingTimeseriesGroupSPFResponseMetaJSON   `json:"-"`
}

// emailRoutingTimeseriesGroupSPFResponseMetaJSON contains the JSON metadata for
// the struct [EmailRoutingTimeseriesGroupSPFResponseMeta]
type emailRoutingTimeseriesGroupSPFResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupSPFResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupSPFResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupSPFResponseMetaAggInterval string

const (
	EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalFifteenMinutes EmailRoutingTimeseriesGroupSPFResponseMetaAggInterval = "FIFTEEN_MINUTES"
	EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalOneHour        EmailRoutingTimeseriesGroupSPFResponseMetaAggInterval = "ONE_HOUR"
	EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalOneDay         EmailRoutingTimeseriesGroupSPFResponseMetaAggInterval = "ONE_DAY"
	EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalOneWeek        EmailRoutingTimeseriesGroupSPFResponseMetaAggInterval = "ONE_WEEK"
	EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalOneMonth       EmailRoutingTimeseriesGroupSPFResponseMetaAggInterval = "ONE_MONTH"
)

func (r EmailRoutingTimeseriesGroupSPFResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalFifteenMinutes, EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalOneHour, EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalOneDay, EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalOneWeek, EmailRoutingTimeseriesGroupSPFResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfo struct {
	Annotations []EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                        `json:"level,required"`
	JSON  emailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoJSON `json:"-"`
}

// emailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfo]
type emailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                        `json:"description,required"`
	EndDate     time.Time                                                                     `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                              `json:"startDate,required" format:"date-time"`
	JSON            emailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// emailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotation]
type emailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceAll                EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceBGP                EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceBots               EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceCT                 EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceDNS                EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceDos                EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceFw                 EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceIQI                EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceNet                EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceAll, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceBGP, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceBots, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceCT, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceDNS, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceDos, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceFw, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceIQI, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceNet, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventType string

const (
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventTypeEvent             EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventTypeOutage            EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventTypePipeline          EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventTypeEvent, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventTypeOutage, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventTypePipeline, EmailRoutingTimeseriesGroupSPFResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupSPFResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      emailRoutingTimeseriesGroupSPFResponseMetaDateRangeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupSPFResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupSPFResponseMetaDateRange]
type emailRoutingTimeseriesGroupSPFResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupSPFResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupSPFResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type EmailRoutingTimeseriesGroupSPFResponseMetaNormalization string

const (
	EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationPercentage           EmailRoutingTimeseriesGroupSPFResponseMetaNormalization = "PERCENTAGE"
	EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationMin0Max              EmailRoutingTimeseriesGroupSPFResponseMetaNormalization = "MIN0_MAX"
	EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationMinMax               EmailRoutingTimeseriesGroupSPFResponseMetaNormalization = "MIN_MAX"
	EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationRawValues            EmailRoutingTimeseriesGroupSPFResponseMetaNormalization = "RAW_VALUES"
	EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationPercentageChange     EmailRoutingTimeseriesGroupSPFResponseMetaNormalization = "PERCENTAGE_CHANGE"
	EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationRollingAverage       EmailRoutingTimeseriesGroupSPFResponseMetaNormalization = "ROLLING_AVERAGE"
	EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationOverlappedPercentage EmailRoutingTimeseriesGroupSPFResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationRatio                EmailRoutingTimeseriesGroupSPFResponseMetaNormalization = "RATIO"
)

func (r EmailRoutingTimeseriesGroupSPFResponseMetaNormalization) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationPercentage, EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationMin0Max, EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationMinMax, EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationRawValues, EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationPercentageChange, EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationRollingAverage, EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationOverlappedPercentage, EmailRoutingTimeseriesGroupSPFResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupSPFResponseMetaUnit struct {
	Name  string                                             `json:"name,required"`
	Value string                                             `json:"value,required"`
	JSON  emailRoutingTimeseriesGroupSPFResponseMetaUnitJSON `json:"-"`
}

// emailRoutingTimeseriesGroupSPFResponseMetaUnitJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupSPFResponseMetaUnit]
type emailRoutingTimeseriesGroupSPFResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupSPFResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupSPFResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupARCParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailRoutingTimeseriesGroupARCParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailRoutingTimeseriesGroupARCParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailRoutingTimeseriesGroupARCParamsDMARC] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]EmailRoutingTimeseriesGroupARCParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingTimeseriesGroupARCParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]EmailRoutingTimeseriesGroupARCParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailRoutingTimeseriesGroupARCParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingTimeseriesGroupARCParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingTimeseriesGroupARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupARCParamsAggInterval string

const (
	EmailRoutingTimeseriesGroupARCParamsAggInterval15m EmailRoutingTimeseriesGroupARCParamsAggInterval = "15m"
	EmailRoutingTimeseriesGroupARCParamsAggInterval1h  EmailRoutingTimeseriesGroupARCParamsAggInterval = "1h"
	EmailRoutingTimeseriesGroupARCParamsAggInterval1d  EmailRoutingTimeseriesGroupARCParamsAggInterval = "1d"
	EmailRoutingTimeseriesGroupARCParamsAggInterval1w  EmailRoutingTimeseriesGroupARCParamsAggInterval = "1w"
)

func (r EmailRoutingTimeseriesGroupARCParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCParamsAggInterval15m, EmailRoutingTimeseriesGroupARCParamsAggInterval1h, EmailRoutingTimeseriesGroupARCParamsAggInterval1d, EmailRoutingTimeseriesGroupARCParamsAggInterval1w:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupARCParamsDKIM string

const (
	EmailRoutingTimeseriesGroupARCParamsDKIMPass EmailRoutingTimeseriesGroupARCParamsDKIM = "PASS"
	EmailRoutingTimeseriesGroupARCParamsDKIMNone EmailRoutingTimeseriesGroupARCParamsDKIM = "NONE"
	EmailRoutingTimeseriesGroupARCParamsDKIMFail EmailRoutingTimeseriesGroupARCParamsDKIM = "FAIL"
)

func (r EmailRoutingTimeseriesGroupARCParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCParamsDKIMPass, EmailRoutingTimeseriesGroupARCParamsDKIMNone, EmailRoutingTimeseriesGroupARCParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupARCParamsDMARC string

const (
	EmailRoutingTimeseriesGroupARCParamsDMARCPass EmailRoutingTimeseriesGroupARCParamsDMARC = "PASS"
	EmailRoutingTimeseriesGroupARCParamsDMARCNone EmailRoutingTimeseriesGroupARCParamsDMARC = "NONE"
	EmailRoutingTimeseriesGroupARCParamsDMARCFail EmailRoutingTimeseriesGroupARCParamsDMARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupARCParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCParamsDMARCPass, EmailRoutingTimeseriesGroupARCParamsDMARCNone, EmailRoutingTimeseriesGroupARCParamsDMARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupARCParamsEncrypted string

const (
	EmailRoutingTimeseriesGroupARCParamsEncryptedEncrypted    EmailRoutingTimeseriesGroupARCParamsEncrypted = "ENCRYPTED"
	EmailRoutingTimeseriesGroupARCParamsEncryptedNotEncrypted EmailRoutingTimeseriesGroupARCParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingTimeseriesGroupARCParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCParamsEncryptedEncrypted, EmailRoutingTimeseriesGroupARCParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingTimeseriesGroupARCParamsFormat string

const (
	EmailRoutingTimeseriesGroupARCParamsFormatJson EmailRoutingTimeseriesGroupARCParamsFormat = "JSON"
	EmailRoutingTimeseriesGroupARCParamsFormatCsv  EmailRoutingTimeseriesGroupARCParamsFormat = "CSV"
)

func (r EmailRoutingTimeseriesGroupARCParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCParamsFormatJson, EmailRoutingTimeseriesGroupARCParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupARCParamsIPVersion string

const (
	EmailRoutingTimeseriesGroupARCParamsIPVersionIPv4 EmailRoutingTimeseriesGroupARCParamsIPVersion = "IPv4"
	EmailRoutingTimeseriesGroupARCParamsIPVersionIPv6 EmailRoutingTimeseriesGroupARCParamsIPVersion = "IPv6"
)

func (r EmailRoutingTimeseriesGroupARCParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCParamsIPVersionIPv4, EmailRoutingTimeseriesGroupARCParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupARCParamsSPF string

const (
	EmailRoutingTimeseriesGroupARCParamsSPFPass EmailRoutingTimeseriesGroupARCParamsSPF = "PASS"
	EmailRoutingTimeseriesGroupARCParamsSPFNone EmailRoutingTimeseriesGroupARCParamsSPF = "NONE"
	EmailRoutingTimeseriesGroupARCParamsSPFFail EmailRoutingTimeseriesGroupARCParamsSPF = "FAIL"
)

func (r EmailRoutingTimeseriesGroupARCParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupARCParamsSPFPass, EmailRoutingTimeseriesGroupARCParamsSPFNone, EmailRoutingTimeseriesGroupARCParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupARCResponseEnvelope struct {
	Result  EmailRoutingTimeseriesGroupARCResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    emailRoutingTimeseriesGroupARCResponseEnvelopeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupARCResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupARCResponseEnvelope]
type emailRoutingTimeseriesGroupARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupDKIMParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailRoutingTimeseriesGroupDKIMParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailRoutingTimeseriesGroupDKIMParamsARC] `query:"arc"`
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
	DMARC param.Field[[]EmailRoutingTimeseriesGroupDKIMParamsDMARC] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]EmailRoutingTimeseriesGroupDKIMParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingTimeseriesGroupDKIMParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]EmailRoutingTimeseriesGroupDKIMParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailRoutingTimeseriesGroupDKIMParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingTimeseriesGroupDKIMParams]'s query parameters
// as `url.Values`.
func (r EmailRoutingTimeseriesGroupDKIMParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupDKIMParamsAggInterval string

const (
	EmailRoutingTimeseriesGroupDKIMParamsAggInterval15m EmailRoutingTimeseriesGroupDKIMParamsAggInterval = "15m"
	EmailRoutingTimeseriesGroupDKIMParamsAggInterval1h  EmailRoutingTimeseriesGroupDKIMParamsAggInterval = "1h"
	EmailRoutingTimeseriesGroupDKIMParamsAggInterval1d  EmailRoutingTimeseriesGroupDKIMParamsAggInterval = "1d"
	EmailRoutingTimeseriesGroupDKIMParamsAggInterval1w  EmailRoutingTimeseriesGroupDKIMParamsAggInterval = "1w"
)

func (r EmailRoutingTimeseriesGroupDKIMParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMParamsAggInterval15m, EmailRoutingTimeseriesGroupDKIMParamsAggInterval1h, EmailRoutingTimeseriesGroupDKIMParamsAggInterval1d, EmailRoutingTimeseriesGroupDKIMParamsAggInterval1w:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDKIMParamsARC string

const (
	EmailRoutingTimeseriesGroupDKIMParamsARCPass EmailRoutingTimeseriesGroupDKIMParamsARC = "PASS"
	EmailRoutingTimeseriesGroupDKIMParamsARCNone EmailRoutingTimeseriesGroupDKIMParamsARC = "NONE"
	EmailRoutingTimeseriesGroupDKIMParamsARCFail EmailRoutingTimeseriesGroupDKIMParamsARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupDKIMParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMParamsARCPass, EmailRoutingTimeseriesGroupDKIMParamsARCNone, EmailRoutingTimeseriesGroupDKIMParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDKIMParamsDMARC string

const (
	EmailRoutingTimeseriesGroupDKIMParamsDMARCPass EmailRoutingTimeseriesGroupDKIMParamsDMARC = "PASS"
	EmailRoutingTimeseriesGroupDKIMParamsDMARCNone EmailRoutingTimeseriesGroupDKIMParamsDMARC = "NONE"
	EmailRoutingTimeseriesGroupDKIMParamsDMARCFail EmailRoutingTimeseriesGroupDKIMParamsDMARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupDKIMParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMParamsDMARCPass, EmailRoutingTimeseriesGroupDKIMParamsDMARCNone, EmailRoutingTimeseriesGroupDKIMParamsDMARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDKIMParamsEncrypted string

const (
	EmailRoutingTimeseriesGroupDKIMParamsEncryptedEncrypted    EmailRoutingTimeseriesGroupDKIMParamsEncrypted = "ENCRYPTED"
	EmailRoutingTimeseriesGroupDKIMParamsEncryptedNotEncrypted EmailRoutingTimeseriesGroupDKIMParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingTimeseriesGroupDKIMParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMParamsEncryptedEncrypted, EmailRoutingTimeseriesGroupDKIMParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingTimeseriesGroupDKIMParamsFormat string

const (
	EmailRoutingTimeseriesGroupDKIMParamsFormatJson EmailRoutingTimeseriesGroupDKIMParamsFormat = "JSON"
	EmailRoutingTimeseriesGroupDKIMParamsFormatCsv  EmailRoutingTimeseriesGroupDKIMParamsFormat = "CSV"
)

func (r EmailRoutingTimeseriesGroupDKIMParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMParamsFormatJson, EmailRoutingTimeseriesGroupDKIMParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDKIMParamsIPVersion string

const (
	EmailRoutingTimeseriesGroupDKIMParamsIPVersionIPv4 EmailRoutingTimeseriesGroupDKIMParamsIPVersion = "IPv4"
	EmailRoutingTimeseriesGroupDKIMParamsIPVersionIPv6 EmailRoutingTimeseriesGroupDKIMParamsIPVersion = "IPv6"
)

func (r EmailRoutingTimeseriesGroupDKIMParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMParamsIPVersionIPv4, EmailRoutingTimeseriesGroupDKIMParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDKIMParamsSPF string

const (
	EmailRoutingTimeseriesGroupDKIMParamsSPFPass EmailRoutingTimeseriesGroupDKIMParamsSPF = "PASS"
	EmailRoutingTimeseriesGroupDKIMParamsSPFNone EmailRoutingTimeseriesGroupDKIMParamsSPF = "NONE"
	EmailRoutingTimeseriesGroupDKIMParamsSPFFail EmailRoutingTimeseriesGroupDKIMParamsSPF = "FAIL"
)

func (r EmailRoutingTimeseriesGroupDKIMParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDKIMParamsSPFPass, EmailRoutingTimeseriesGroupDKIMParamsSPFNone, EmailRoutingTimeseriesGroupDKIMParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDKIMResponseEnvelope struct {
	Result  EmailRoutingTimeseriesGroupDKIMResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    emailRoutingTimeseriesGroupDKIMResponseEnvelopeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDKIMResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupDKIMResponseEnvelope]
type emailRoutingTimeseriesGroupDKIMResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDKIMResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDKIMResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupDMARCParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailRoutingTimeseriesGroupDMARCParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailRoutingTimeseriesGroupDMARCParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailRoutingTimeseriesGroupDMARCParamsDKIM] `query:"dkim"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]EmailRoutingTimeseriesGroupDMARCParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingTimeseriesGroupDMARCParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]EmailRoutingTimeseriesGroupDMARCParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailRoutingTimeseriesGroupDMARCParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingTimeseriesGroupDMARCParams]'s query parameters
// as `url.Values`.
func (r EmailRoutingTimeseriesGroupDMARCParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupDMARCParamsAggInterval string

const (
	EmailRoutingTimeseriesGroupDMARCParamsAggInterval15m EmailRoutingTimeseriesGroupDMARCParamsAggInterval = "15m"
	EmailRoutingTimeseriesGroupDMARCParamsAggInterval1h  EmailRoutingTimeseriesGroupDMARCParamsAggInterval = "1h"
	EmailRoutingTimeseriesGroupDMARCParamsAggInterval1d  EmailRoutingTimeseriesGroupDMARCParamsAggInterval = "1d"
	EmailRoutingTimeseriesGroupDMARCParamsAggInterval1w  EmailRoutingTimeseriesGroupDMARCParamsAggInterval = "1w"
)

func (r EmailRoutingTimeseriesGroupDMARCParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCParamsAggInterval15m, EmailRoutingTimeseriesGroupDMARCParamsAggInterval1h, EmailRoutingTimeseriesGroupDMARCParamsAggInterval1d, EmailRoutingTimeseriesGroupDMARCParamsAggInterval1w:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDMARCParamsARC string

const (
	EmailRoutingTimeseriesGroupDMARCParamsARCPass EmailRoutingTimeseriesGroupDMARCParamsARC = "PASS"
	EmailRoutingTimeseriesGroupDMARCParamsARCNone EmailRoutingTimeseriesGroupDMARCParamsARC = "NONE"
	EmailRoutingTimeseriesGroupDMARCParamsARCFail EmailRoutingTimeseriesGroupDMARCParamsARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupDMARCParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCParamsARCPass, EmailRoutingTimeseriesGroupDMARCParamsARCNone, EmailRoutingTimeseriesGroupDMARCParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDMARCParamsDKIM string

const (
	EmailRoutingTimeseriesGroupDMARCParamsDKIMPass EmailRoutingTimeseriesGroupDMARCParamsDKIM = "PASS"
	EmailRoutingTimeseriesGroupDMARCParamsDKIMNone EmailRoutingTimeseriesGroupDMARCParamsDKIM = "NONE"
	EmailRoutingTimeseriesGroupDMARCParamsDKIMFail EmailRoutingTimeseriesGroupDMARCParamsDKIM = "FAIL"
)

func (r EmailRoutingTimeseriesGroupDMARCParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCParamsDKIMPass, EmailRoutingTimeseriesGroupDMARCParamsDKIMNone, EmailRoutingTimeseriesGroupDMARCParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDMARCParamsEncrypted string

const (
	EmailRoutingTimeseriesGroupDMARCParamsEncryptedEncrypted    EmailRoutingTimeseriesGroupDMARCParamsEncrypted = "ENCRYPTED"
	EmailRoutingTimeseriesGroupDMARCParamsEncryptedNotEncrypted EmailRoutingTimeseriesGroupDMARCParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingTimeseriesGroupDMARCParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCParamsEncryptedEncrypted, EmailRoutingTimeseriesGroupDMARCParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingTimeseriesGroupDMARCParamsFormat string

const (
	EmailRoutingTimeseriesGroupDMARCParamsFormatJson EmailRoutingTimeseriesGroupDMARCParamsFormat = "JSON"
	EmailRoutingTimeseriesGroupDMARCParamsFormatCsv  EmailRoutingTimeseriesGroupDMARCParamsFormat = "CSV"
)

func (r EmailRoutingTimeseriesGroupDMARCParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCParamsFormatJson, EmailRoutingTimeseriesGroupDMARCParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDMARCParamsIPVersion string

const (
	EmailRoutingTimeseriesGroupDMARCParamsIPVersionIPv4 EmailRoutingTimeseriesGroupDMARCParamsIPVersion = "IPv4"
	EmailRoutingTimeseriesGroupDMARCParamsIPVersionIPv6 EmailRoutingTimeseriesGroupDMARCParamsIPVersion = "IPv6"
)

func (r EmailRoutingTimeseriesGroupDMARCParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCParamsIPVersionIPv4, EmailRoutingTimeseriesGroupDMARCParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDMARCParamsSPF string

const (
	EmailRoutingTimeseriesGroupDMARCParamsSPFPass EmailRoutingTimeseriesGroupDMARCParamsSPF = "PASS"
	EmailRoutingTimeseriesGroupDMARCParamsSPFNone EmailRoutingTimeseriesGroupDMARCParamsSPF = "NONE"
	EmailRoutingTimeseriesGroupDMARCParamsSPFFail EmailRoutingTimeseriesGroupDMARCParamsSPF = "FAIL"
)

func (r EmailRoutingTimeseriesGroupDMARCParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupDMARCParamsSPFPass, EmailRoutingTimeseriesGroupDMARCParamsSPFNone, EmailRoutingTimeseriesGroupDMARCParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupDMARCResponseEnvelope struct {
	Result  EmailRoutingTimeseriesGroupDMARCResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    emailRoutingTimeseriesGroupDMARCResponseEnvelopeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupDMARCResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupDMARCResponseEnvelope]
type emailRoutingTimeseriesGroupDMARCResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupDMARCResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupDMARCResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupEncryptedParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailRoutingTimeseriesGroupEncryptedParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailRoutingTimeseriesGroupEncryptedParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailRoutingTimeseriesGroupEncryptedParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailRoutingTimeseriesGroupEncryptedParamsDMARC] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingTimeseriesGroupEncryptedParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]EmailRoutingTimeseriesGroupEncryptedParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailRoutingTimeseriesGroupEncryptedParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingTimeseriesGroupEncryptedParams]'s query
// parameters as `url.Values`.
func (r EmailRoutingTimeseriesGroupEncryptedParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupEncryptedParamsAggInterval string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsAggInterval15m EmailRoutingTimeseriesGroupEncryptedParamsAggInterval = "15m"
	EmailRoutingTimeseriesGroupEncryptedParamsAggInterval1h  EmailRoutingTimeseriesGroupEncryptedParamsAggInterval = "1h"
	EmailRoutingTimeseriesGroupEncryptedParamsAggInterval1d  EmailRoutingTimeseriesGroupEncryptedParamsAggInterval = "1d"
	EmailRoutingTimeseriesGroupEncryptedParamsAggInterval1w  EmailRoutingTimeseriesGroupEncryptedParamsAggInterval = "1w"
)

func (r EmailRoutingTimeseriesGroupEncryptedParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedParamsAggInterval15m, EmailRoutingTimeseriesGroupEncryptedParamsAggInterval1h, EmailRoutingTimeseriesGroupEncryptedParamsAggInterval1d, EmailRoutingTimeseriesGroupEncryptedParamsAggInterval1w:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupEncryptedParamsARC string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsARCPass EmailRoutingTimeseriesGroupEncryptedParamsARC = "PASS"
	EmailRoutingTimeseriesGroupEncryptedParamsARCNone EmailRoutingTimeseriesGroupEncryptedParamsARC = "NONE"
	EmailRoutingTimeseriesGroupEncryptedParamsARCFail EmailRoutingTimeseriesGroupEncryptedParamsARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupEncryptedParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedParamsARCPass, EmailRoutingTimeseriesGroupEncryptedParamsARCNone, EmailRoutingTimeseriesGroupEncryptedParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupEncryptedParamsDKIM string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsDKIMPass EmailRoutingTimeseriesGroupEncryptedParamsDKIM = "PASS"
	EmailRoutingTimeseriesGroupEncryptedParamsDKIMNone EmailRoutingTimeseriesGroupEncryptedParamsDKIM = "NONE"
	EmailRoutingTimeseriesGroupEncryptedParamsDKIMFail EmailRoutingTimeseriesGroupEncryptedParamsDKIM = "FAIL"
)

func (r EmailRoutingTimeseriesGroupEncryptedParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedParamsDKIMPass, EmailRoutingTimeseriesGroupEncryptedParamsDKIMNone, EmailRoutingTimeseriesGroupEncryptedParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupEncryptedParamsDMARC string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsDMARCPass EmailRoutingTimeseriesGroupEncryptedParamsDMARC = "PASS"
	EmailRoutingTimeseriesGroupEncryptedParamsDMARCNone EmailRoutingTimeseriesGroupEncryptedParamsDMARC = "NONE"
	EmailRoutingTimeseriesGroupEncryptedParamsDMARCFail EmailRoutingTimeseriesGroupEncryptedParamsDMARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupEncryptedParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedParamsDMARCPass, EmailRoutingTimeseriesGroupEncryptedParamsDMARCNone, EmailRoutingTimeseriesGroupEncryptedParamsDMARCFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingTimeseriesGroupEncryptedParamsFormat string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsFormatJson EmailRoutingTimeseriesGroupEncryptedParamsFormat = "JSON"
	EmailRoutingTimeseriesGroupEncryptedParamsFormatCsv  EmailRoutingTimeseriesGroupEncryptedParamsFormat = "CSV"
)

func (r EmailRoutingTimeseriesGroupEncryptedParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedParamsFormatJson, EmailRoutingTimeseriesGroupEncryptedParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupEncryptedParamsIPVersion string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsIPVersionIPv4 EmailRoutingTimeseriesGroupEncryptedParamsIPVersion = "IPv4"
	EmailRoutingTimeseriesGroupEncryptedParamsIPVersionIPv6 EmailRoutingTimeseriesGroupEncryptedParamsIPVersion = "IPv6"
)

func (r EmailRoutingTimeseriesGroupEncryptedParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedParamsIPVersionIPv4, EmailRoutingTimeseriesGroupEncryptedParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupEncryptedParamsSPF string

const (
	EmailRoutingTimeseriesGroupEncryptedParamsSPFPass EmailRoutingTimeseriesGroupEncryptedParamsSPF = "PASS"
	EmailRoutingTimeseriesGroupEncryptedParamsSPFNone EmailRoutingTimeseriesGroupEncryptedParamsSPF = "NONE"
	EmailRoutingTimeseriesGroupEncryptedParamsSPFFail EmailRoutingTimeseriesGroupEncryptedParamsSPF = "FAIL"
)

func (r EmailRoutingTimeseriesGroupEncryptedParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupEncryptedParamsSPFPass, EmailRoutingTimeseriesGroupEncryptedParamsSPFNone, EmailRoutingTimeseriesGroupEncryptedParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupEncryptedResponseEnvelope struct {
	Result  EmailRoutingTimeseriesGroupEncryptedResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    emailRoutingTimeseriesGroupEncryptedResponseEnvelopeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupEncryptedResponseEnvelopeJSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupEncryptedResponseEnvelope]
type emailRoutingTimeseriesGroupEncryptedResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupEncryptedResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupEncryptedResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupIPVersionParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailRoutingTimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailRoutingTimeseriesGroupIPVersionParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailRoutingTimeseriesGroupIPVersionParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailRoutingTimeseriesGroupIPVersionParamsDMARC] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]EmailRoutingTimeseriesGroupIPVersionParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingTimeseriesGroupIPVersionParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	SPF param.Field[[]EmailRoutingTimeseriesGroupIPVersionParamsSPF] `query:"spf"`
}

// URLQuery serializes [EmailRoutingTimeseriesGroupIPVersionParams]'s query
// parameters as `url.Values`.
func (r EmailRoutingTimeseriesGroupIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupIPVersionParamsAggInterval string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsAggInterval15m EmailRoutingTimeseriesGroupIPVersionParamsAggInterval = "15m"
	EmailRoutingTimeseriesGroupIPVersionParamsAggInterval1h  EmailRoutingTimeseriesGroupIPVersionParamsAggInterval = "1h"
	EmailRoutingTimeseriesGroupIPVersionParamsAggInterval1d  EmailRoutingTimeseriesGroupIPVersionParamsAggInterval = "1d"
	EmailRoutingTimeseriesGroupIPVersionParamsAggInterval1w  EmailRoutingTimeseriesGroupIPVersionParamsAggInterval = "1w"
)

func (r EmailRoutingTimeseriesGroupIPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionParamsAggInterval15m, EmailRoutingTimeseriesGroupIPVersionParamsAggInterval1h, EmailRoutingTimeseriesGroupIPVersionParamsAggInterval1d, EmailRoutingTimeseriesGroupIPVersionParamsAggInterval1w:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupIPVersionParamsARC string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsARCPass EmailRoutingTimeseriesGroupIPVersionParamsARC = "PASS"
	EmailRoutingTimeseriesGroupIPVersionParamsARCNone EmailRoutingTimeseriesGroupIPVersionParamsARC = "NONE"
	EmailRoutingTimeseriesGroupIPVersionParamsARCFail EmailRoutingTimeseriesGroupIPVersionParamsARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupIPVersionParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionParamsARCPass, EmailRoutingTimeseriesGroupIPVersionParamsARCNone, EmailRoutingTimeseriesGroupIPVersionParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupIPVersionParamsDKIM string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsDKIMPass EmailRoutingTimeseriesGroupIPVersionParamsDKIM = "PASS"
	EmailRoutingTimeseriesGroupIPVersionParamsDKIMNone EmailRoutingTimeseriesGroupIPVersionParamsDKIM = "NONE"
	EmailRoutingTimeseriesGroupIPVersionParamsDKIMFail EmailRoutingTimeseriesGroupIPVersionParamsDKIM = "FAIL"
)

func (r EmailRoutingTimeseriesGroupIPVersionParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionParamsDKIMPass, EmailRoutingTimeseriesGroupIPVersionParamsDKIMNone, EmailRoutingTimeseriesGroupIPVersionParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupIPVersionParamsDMARC string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsDMARCPass EmailRoutingTimeseriesGroupIPVersionParamsDMARC = "PASS"
	EmailRoutingTimeseriesGroupIPVersionParamsDMARCNone EmailRoutingTimeseriesGroupIPVersionParamsDMARC = "NONE"
	EmailRoutingTimeseriesGroupIPVersionParamsDMARCFail EmailRoutingTimeseriesGroupIPVersionParamsDMARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupIPVersionParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionParamsDMARCPass, EmailRoutingTimeseriesGroupIPVersionParamsDMARCNone, EmailRoutingTimeseriesGroupIPVersionParamsDMARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupIPVersionParamsEncrypted string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsEncryptedEncrypted    EmailRoutingTimeseriesGroupIPVersionParamsEncrypted = "ENCRYPTED"
	EmailRoutingTimeseriesGroupIPVersionParamsEncryptedNotEncrypted EmailRoutingTimeseriesGroupIPVersionParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingTimeseriesGroupIPVersionParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionParamsEncryptedEncrypted, EmailRoutingTimeseriesGroupIPVersionParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingTimeseriesGroupIPVersionParamsFormat string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsFormatJson EmailRoutingTimeseriesGroupIPVersionParamsFormat = "JSON"
	EmailRoutingTimeseriesGroupIPVersionParamsFormatCsv  EmailRoutingTimeseriesGroupIPVersionParamsFormat = "CSV"
)

func (r EmailRoutingTimeseriesGroupIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionParamsFormatJson, EmailRoutingTimeseriesGroupIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupIPVersionParamsSPF string

const (
	EmailRoutingTimeseriesGroupIPVersionParamsSPFPass EmailRoutingTimeseriesGroupIPVersionParamsSPF = "PASS"
	EmailRoutingTimeseriesGroupIPVersionParamsSPFNone EmailRoutingTimeseriesGroupIPVersionParamsSPF = "NONE"
	EmailRoutingTimeseriesGroupIPVersionParamsSPFFail EmailRoutingTimeseriesGroupIPVersionParamsSPF = "FAIL"
)

func (r EmailRoutingTimeseriesGroupIPVersionParamsSPF) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupIPVersionParamsSPFPass, EmailRoutingTimeseriesGroupIPVersionParamsSPFNone, EmailRoutingTimeseriesGroupIPVersionParamsSPFFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupIPVersionResponseEnvelope struct {
	Result  EmailRoutingTimeseriesGroupIPVersionResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    emailRoutingTimeseriesGroupIPVersionResponseEnvelopeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupIPVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct [EmailRoutingTimeseriesGroupIPVersionResponseEnvelope]
type emailRoutingTimeseriesGroupIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingTimeseriesGroupSPFParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[EmailRoutingTimeseriesGroupSPFParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	ARC param.Field[[]EmailRoutingTimeseriesGroupSPFParamsARC] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	DKIM param.Field[[]EmailRoutingTimeseriesGroupSPFParamsDKIM] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	DMARC param.Field[[]EmailRoutingTimeseriesGroupSPFParamsDMARC] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]EmailRoutingTimeseriesGroupSPFParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[EmailRoutingTimeseriesGroupSPFParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]EmailRoutingTimeseriesGroupSPFParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [EmailRoutingTimeseriesGroupSPFParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingTimeseriesGroupSPFParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type EmailRoutingTimeseriesGroupSPFParamsAggInterval string

const (
	EmailRoutingTimeseriesGroupSPFParamsAggInterval15m EmailRoutingTimeseriesGroupSPFParamsAggInterval = "15m"
	EmailRoutingTimeseriesGroupSPFParamsAggInterval1h  EmailRoutingTimeseriesGroupSPFParamsAggInterval = "1h"
	EmailRoutingTimeseriesGroupSPFParamsAggInterval1d  EmailRoutingTimeseriesGroupSPFParamsAggInterval = "1d"
	EmailRoutingTimeseriesGroupSPFParamsAggInterval1w  EmailRoutingTimeseriesGroupSPFParamsAggInterval = "1w"
)

func (r EmailRoutingTimeseriesGroupSPFParamsAggInterval) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFParamsAggInterval15m, EmailRoutingTimeseriesGroupSPFParamsAggInterval1h, EmailRoutingTimeseriesGroupSPFParamsAggInterval1d, EmailRoutingTimeseriesGroupSPFParamsAggInterval1w:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupSPFParamsARC string

const (
	EmailRoutingTimeseriesGroupSPFParamsARCPass EmailRoutingTimeseriesGroupSPFParamsARC = "PASS"
	EmailRoutingTimeseriesGroupSPFParamsARCNone EmailRoutingTimeseriesGroupSPFParamsARC = "NONE"
	EmailRoutingTimeseriesGroupSPFParamsARCFail EmailRoutingTimeseriesGroupSPFParamsARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupSPFParamsARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFParamsARCPass, EmailRoutingTimeseriesGroupSPFParamsARCNone, EmailRoutingTimeseriesGroupSPFParamsARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupSPFParamsDKIM string

const (
	EmailRoutingTimeseriesGroupSPFParamsDKIMPass EmailRoutingTimeseriesGroupSPFParamsDKIM = "PASS"
	EmailRoutingTimeseriesGroupSPFParamsDKIMNone EmailRoutingTimeseriesGroupSPFParamsDKIM = "NONE"
	EmailRoutingTimeseriesGroupSPFParamsDKIMFail EmailRoutingTimeseriesGroupSPFParamsDKIM = "FAIL"
)

func (r EmailRoutingTimeseriesGroupSPFParamsDKIM) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFParamsDKIMPass, EmailRoutingTimeseriesGroupSPFParamsDKIMNone, EmailRoutingTimeseriesGroupSPFParamsDKIMFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupSPFParamsDMARC string

const (
	EmailRoutingTimeseriesGroupSPFParamsDMARCPass EmailRoutingTimeseriesGroupSPFParamsDMARC = "PASS"
	EmailRoutingTimeseriesGroupSPFParamsDMARCNone EmailRoutingTimeseriesGroupSPFParamsDMARC = "NONE"
	EmailRoutingTimeseriesGroupSPFParamsDMARCFail EmailRoutingTimeseriesGroupSPFParamsDMARC = "FAIL"
)

func (r EmailRoutingTimeseriesGroupSPFParamsDMARC) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFParamsDMARCPass, EmailRoutingTimeseriesGroupSPFParamsDMARCNone, EmailRoutingTimeseriesGroupSPFParamsDMARCFail:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupSPFParamsEncrypted string

const (
	EmailRoutingTimeseriesGroupSPFParamsEncryptedEncrypted    EmailRoutingTimeseriesGroupSPFParamsEncrypted = "ENCRYPTED"
	EmailRoutingTimeseriesGroupSPFParamsEncryptedNotEncrypted EmailRoutingTimeseriesGroupSPFParamsEncrypted = "NOT_ENCRYPTED"
)

func (r EmailRoutingTimeseriesGroupSPFParamsEncrypted) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFParamsEncryptedEncrypted, EmailRoutingTimeseriesGroupSPFParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type EmailRoutingTimeseriesGroupSPFParamsFormat string

const (
	EmailRoutingTimeseriesGroupSPFParamsFormatJson EmailRoutingTimeseriesGroupSPFParamsFormat = "JSON"
	EmailRoutingTimeseriesGroupSPFParamsFormatCsv  EmailRoutingTimeseriesGroupSPFParamsFormat = "CSV"
)

func (r EmailRoutingTimeseriesGroupSPFParamsFormat) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFParamsFormatJson, EmailRoutingTimeseriesGroupSPFParamsFormatCsv:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupSPFParamsIPVersion string

const (
	EmailRoutingTimeseriesGroupSPFParamsIPVersionIPv4 EmailRoutingTimeseriesGroupSPFParamsIPVersion = "IPv4"
	EmailRoutingTimeseriesGroupSPFParamsIPVersionIPv6 EmailRoutingTimeseriesGroupSPFParamsIPVersion = "IPv6"
)

func (r EmailRoutingTimeseriesGroupSPFParamsIPVersion) IsKnown() bool {
	switch r {
	case EmailRoutingTimeseriesGroupSPFParamsIPVersionIPv4, EmailRoutingTimeseriesGroupSPFParamsIPVersionIPv6:
		return true
	}
	return false
}

type EmailRoutingTimeseriesGroupSPFResponseEnvelope struct {
	Result  EmailRoutingTimeseriesGroupSPFResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    emailRoutingTimeseriesGroupSPFResponseEnvelopeJSON `json:"-"`
}

// emailRoutingTimeseriesGroupSPFResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailRoutingTimeseriesGroupSPFResponseEnvelope]
type emailRoutingTimeseriesGroupSPFResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingTimeseriesGroupSPFResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingTimeseriesGroupSPFResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
