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

// NetflowService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetflowService] method instead.
type NetflowService struct {
	Options []option.RequestOption
	Top     *NetflowTopService
}

// NewNetflowService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNetflowService(opts ...option.RequestOption) (r *NetflowService) {
	r = &NetflowService{}
	r.Options = opts
	r.Top = NewNetflowTopService(opts...)
	return
}

// Retrieves the distribution of network traffic (NetFlows) by HTTP vs other
// protocols.
//
// Deprecated: Use
// [Get Network Traffic Distribution By Dimension](https://developers.cloudflare.com/api/resources/radar/subresources/netflows/methods/summary_v2/)
// instead.
func (r *NetflowService) Summary(ctx context.Context, query NetflowSummaryParams, opts ...option.RequestOption) (res *NetflowSummaryResponse, err error) {
	var env NetflowSummaryResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/netflows/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves network traffic (NetFlows) over time.
func (r *NetflowService) Timeseries(ctx context.Context, query NetflowTimeseriesParams, opts ...option.RequestOption) (res *NetflowTimeseriesResponse, err error) {
	var env NetflowTimeseriesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/netflows/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type NetflowSummaryResponse struct {
	// Metadata for the results.
	Meta     NetflowSummaryResponseMeta     `json:"meta" api:"required"`
	Summary0 NetflowSummaryResponseSummary0 `json:"summary_0" api:"required"`
	JSON     netflowSummaryResponseJSON     `json:"-"`
}

// netflowSummaryResponseJSON contains the JSON metadata for the struct
// [NetflowSummaryResponse]
type netflowSummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowSummaryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type NetflowSummaryResponseMeta struct {
	ConfidenceInfo NetflowSummaryResponseMetaConfidenceInfo `json:"confidenceInfo" api:"required"`
	DateRange      []NetflowSummaryResponseMetaDateRange    `json:"dateRange" api:"required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated" api:"required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization NetflowSummaryResponseMetaNormalization `json:"normalization" api:"required"`
	// Measurement units for the results.
	Units []NetflowSummaryResponseMetaUnit `json:"units" api:"required"`
	JSON  netflowSummaryResponseMetaJSON   `json:"-"`
}

// netflowSummaryResponseMetaJSON contains the JSON metadata for the struct
// [NetflowSummaryResponseMeta]
type netflowSummaryResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NetflowSummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowSummaryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type NetflowSummaryResponseMetaConfidenceInfo struct {
	Annotations []NetflowSummaryResponseMetaConfidenceInfoAnnotation `json:"annotations" api:"required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                        `json:"level" api:"required"`
	JSON  netflowSummaryResponseMetaConfidenceInfoJSON `json:"-"`
}

// netflowSummaryResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [NetflowSummaryResponseMetaConfidenceInfo]
type netflowSummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowSummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowSummaryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type NetflowSummaryResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource" api:"required"`
	Description string                                                        `json:"description" api:"required"`
	EndDate     time.Time                                                     `json:"endDate" api:"required" format:"date-time"`
	// Event type for annotations.
	EventType NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType" api:"required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                   `json:"isInstantaneous" api:"required"`
	LinkedURL       string                                                 `json:"linkedUrl" api:"required" format:"uri"`
	StartDate       time.Time                                              `json:"startDate" api:"required" format:"date-time"`
	JSON            netflowSummaryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// netflowSummaryResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [NetflowSummaryResponseMetaConfidenceInfoAnnotation]
type netflowSummaryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *NetflowSummaryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowSummaryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAll                NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBGP                NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBots               NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceCt                 NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNS                NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDos                NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFw                 NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceIQI                NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceNet                NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAll, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBGP, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBots, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceCt, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNS, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDos, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFw, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceIQI, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceNet, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, NetflowSummaryResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventType string

const (
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventTypeEvent             NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventTypeOutage            NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventTypePipeline          NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventTypeEvent, NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventTypeOutage, NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventTypePipeline, NetflowSummaryResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type NetflowSummaryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime" api:"required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                               `json:"startTime" api:"required" format:"date-time"`
	JSON      netflowSummaryResponseMetaDateRangeJSON `json:"-"`
}

// netflowSummaryResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [NetflowSummaryResponseMetaDateRange]
type netflowSummaryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowSummaryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowSummaryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type NetflowSummaryResponseMetaNormalization string

const (
	NetflowSummaryResponseMetaNormalizationPercentage           NetflowSummaryResponseMetaNormalization = "PERCENTAGE"
	NetflowSummaryResponseMetaNormalizationMin0Max              NetflowSummaryResponseMetaNormalization = "MIN0_MAX"
	NetflowSummaryResponseMetaNormalizationMinMax               NetflowSummaryResponseMetaNormalization = "MIN_MAX"
	NetflowSummaryResponseMetaNormalizationRawValues            NetflowSummaryResponseMetaNormalization = "RAW_VALUES"
	NetflowSummaryResponseMetaNormalizationPercentageChange     NetflowSummaryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	NetflowSummaryResponseMetaNormalizationRollingAverage       NetflowSummaryResponseMetaNormalization = "ROLLING_AVERAGE"
	NetflowSummaryResponseMetaNormalizationOverlappedPercentage NetflowSummaryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	NetflowSummaryResponseMetaNormalizationRatio                NetflowSummaryResponseMetaNormalization = "RATIO"
)

func (r NetflowSummaryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case NetflowSummaryResponseMetaNormalizationPercentage, NetflowSummaryResponseMetaNormalizationMin0Max, NetflowSummaryResponseMetaNormalizationMinMax, NetflowSummaryResponseMetaNormalizationRawValues, NetflowSummaryResponseMetaNormalizationPercentageChange, NetflowSummaryResponseMetaNormalizationRollingAverage, NetflowSummaryResponseMetaNormalizationOverlappedPercentage, NetflowSummaryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type NetflowSummaryResponseMetaUnit struct {
	Name  string                             `json:"name" api:"required"`
	Value string                             `json:"value" api:"required"`
	JSON  netflowSummaryResponseMetaUnitJSON `json:"-"`
}

// netflowSummaryResponseMetaUnitJSON contains the JSON metadata for the struct
// [NetflowSummaryResponseMetaUnit]
type netflowSummaryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowSummaryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowSummaryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type NetflowSummaryResponseSummary0 struct {
	// A numeric string.
	HTTP string `json:"HTTP" api:"required"`
	// A numeric string.
	Other string                             `json:"OTHER" api:"required"`
	JSON  netflowSummaryResponseSummary0JSON `json:"-"`
}

// netflowSummaryResponseSummary0JSON contains the JSON metadata for the struct
// [NetflowSummaryResponseSummary0]
type netflowSummaryResponseSummary0JSON struct {
	HTTP        apijson.Field
	Other       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowSummaryResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowSummaryResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type NetflowTimeseriesResponse struct {
	// Metadata for the results.
	Meta   NetflowTimeseriesResponseMeta   `json:"meta" api:"required"`
	Serie0 NetflowTimeseriesResponseSerie0 `json:"serie_0" api:"required"`
	JSON   netflowTimeseriesResponseJSON   `json:"-"`
}

// netflowTimeseriesResponseJSON contains the JSON metadata for the struct
// [NetflowTimeseriesResponse]
type netflowTimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type NetflowTimeseriesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    NetflowTimeseriesResponseMetaAggInterval    `json:"aggInterval" api:"required"`
	ConfidenceInfo NetflowTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo" api:"required"`
	DateRange      []NetflowTimeseriesResponseMetaDateRange    `json:"dateRange" api:"required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated" api:"required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization NetflowTimeseriesResponseMetaNormalization `json:"normalization" api:"required"`
	// Measurement units for the results.
	Units []NetflowTimeseriesResponseMetaUnit `json:"units" api:"required"`
	JSON  netflowTimeseriesResponseMetaJSON   `json:"-"`
}

// netflowTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [NetflowTimeseriesResponseMeta]
type netflowTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NetflowTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type NetflowTimeseriesResponseMetaAggInterval string

const (
	NetflowTimeseriesResponseMetaAggIntervalFifteenMinutes NetflowTimeseriesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	NetflowTimeseriesResponseMetaAggIntervalOneHour        NetflowTimeseriesResponseMetaAggInterval = "ONE_HOUR"
	NetflowTimeseriesResponseMetaAggIntervalOneDay         NetflowTimeseriesResponseMetaAggInterval = "ONE_DAY"
	NetflowTimeseriesResponseMetaAggIntervalOneWeek        NetflowTimeseriesResponseMetaAggInterval = "ONE_WEEK"
	NetflowTimeseriesResponseMetaAggIntervalOneMonth       NetflowTimeseriesResponseMetaAggInterval = "ONE_MONTH"
)

func (r NetflowTimeseriesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case NetflowTimeseriesResponseMetaAggIntervalFifteenMinutes, NetflowTimeseriesResponseMetaAggIntervalOneHour, NetflowTimeseriesResponseMetaAggIntervalOneDay, NetflowTimeseriesResponseMetaAggIntervalOneWeek, NetflowTimeseriesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type NetflowTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []NetflowTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations" api:"required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                           `json:"level" api:"required"`
	JSON  netflowTimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// netflowTimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [NetflowTimeseriesResponseMetaConfidenceInfo]
type netflowTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type NetflowTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource" api:"required"`
	Description string                                                           `json:"description" api:"required"`
	EndDate     time.Time                                                        `json:"endDate" api:"required" format:"date-time"`
	// Event type for annotations.
	EventType NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType" api:"required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                      `json:"isInstantaneous" api:"required"`
	LinkedURL       string                                                    `json:"linkedUrl" api:"required" format:"uri"`
	StartDate       time.Time                                                 `json:"startDate" api:"required" format:"date-time"`
	JSON            netflowTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// netflowTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [NetflowTimeseriesResponseMetaConfidenceInfoAnnotation]
type netflowTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *NetflowTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll                NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP                NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots               NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCt                 NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS                NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos                NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw                 NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI                NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet                NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCt, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventType string

const (
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent             NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage            NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline          NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline, NetflowTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type NetflowTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime" api:"required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                  `json:"startTime" api:"required" format:"date-time"`
	JSON      netflowTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// netflowTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [NetflowTimeseriesResponseMetaDateRange]
type netflowTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type NetflowTimeseriesResponseMetaNormalization string

const (
	NetflowTimeseriesResponseMetaNormalizationPercentage           NetflowTimeseriesResponseMetaNormalization = "PERCENTAGE"
	NetflowTimeseriesResponseMetaNormalizationMin0Max              NetflowTimeseriesResponseMetaNormalization = "MIN0_MAX"
	NetflowTimeseriesResponseMetaNormalizationMinMax               NetflowTimeseriesResponseMetaNormalization = "MIN_MAX"
	NetflowTimeseriesResponseMetaNormalizationRawValues            NetflowTimeseriesResponseMetaNormalization = "RAW_VALUES"
	NetflowTimeseriesResponseMetaNormalizationPercentageChange     NetflowTimeseriesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	NetflowTimeseriesResponseMetaNormalizationRollingAverage       NetflowTimeseriesResponseMetaNormalization = "ROLLING_AVERAGE"
	NetflowTimeseriesResponseMetaNormalizationOverlappedPercentage NetflowTimeseriesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	NetflowTimeseriesResponseMetaNormalizationRatio                NetflowTimeseriesResponseMetaNormalization = "RATIO"
)

func (r NetflowTimeseriesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case NetflowTimeseriesResponseMetaNormalizationPercentage, NetflowTimeseriesResponseMetaNormalizationMin0Max, NetflowTimeseriesResponseMetaNormalizationMinMax, NetflowTimeseriesResponseMetaNormalizationRawValues, NetflowTimeseriesResponseMetaNormalizationPercentageChange, NetflowTimeseriesResponseMetaNormalizationRollingAverage, NetflowTimeseriesResponseMetaNormalizationOverlappedPercentage, NetflowTimeseriesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type NetflowTimeseriesResponseMetaUnit struct {
	Name  string                                `json:"name" api:"required"`
	Value string                                `json:"value" api:"required"`
	JSON  netflowTimeseriesResponseMetaUnitJSON `json:"-"`
}

// netflowTimeseriesResponseMetaUnitJSON contains the JSON metadata for the struct
// [NetflowTimeseriesResponseMetaUnit]
type netflowTimeseriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTimeseriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTimeseriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type NetflowTimeseriesResponseSerie0 struct {
	Timestamps []time.Time                         `json:"timestamps" api:"required" format:"date-time"`
	Values     []string                            `json:"values" api:"required"`
	JSON       netflowTimeseriesResponseSerie0JSON `json:"-"`
}

// netflowTimeseriesResponseSerie0JSON contains the JSON metadata for the struct
// [NetflowTimeseriesResponseSerie0]
type netflowTimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type NetflowSummaryParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[NetflowSummaryParamsFormat] `query:"format"`
	// Filters results by Geolocation. Specify a comma-separated list of GeoNames IDs.
	// Prefix with `-` to exclude geoIds from results. For example, `-2267056,360689`
	// excludes results from the 2267056 (Lisbon), but includes results from 5128638
	// (New York).
	GeoID param.Field[[]string] `query:"geoId"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [NetflowSummaryParams]'s query parameters as `url.Values`.
func (r NetflowSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type NetflowSummaryParamsFormat string

const (
	NetflowSummaryParamsFormatJson NetflowSummaryParamsFormat = "JSON"
	NetflowSummaryParamsFormatCsv  NetflowSummaryParamsFormat = "CSV"
)

func (r NetflowSummaryParamsFormat) IsKnown() bool {
	switch r {
	case NetflowSummaryParamsFormatJson, NetflowSummaryParamsFormatCsv:
		return true
	}
	return false
}

type NetflowSummaryResponseEnvelope struct {
	Result  NetflowSummaryResponse             `json:"result" api:"required"`
	Success bool                               `json:"success" api:"required"`
	JSON    netflowSummaryResponseEnvelopeJSON `json:"-"`
}

// netflowSummaryResponseEnvelopeJSON contains the JSON metadata for the struct
// [NetflowSummaryResponseEnvelope]
type netflowSummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowSummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NetflowTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[NetflowTimeseriesParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[NetflowTimeseriesParamsFormat] `query:"format"`
	// Filters results by Geolocation. Specify a comma-separated list of GeoNames IDs.
	// Prefix with `-` to exclude geoIds from results. For example, `-2267056,360689`
	// excludes results from the 2267056 (Lisbon), but includes results from 5128638
	// (New York).
	GeoID param.Field[[]string] `query:"geoId"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[NetflowTimeseriesParamsNormalization] `query:"normalization"`
	// Filters the results by network traffic product types.
	Product param.Field[[]NetflowTimeseriesParamsProduct] `query:"product"`
}

// URLQuery serializes [NetflowTimeseriesParams]'s query parameters as
// `url.Values`.
func (r NetflowTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type NetflowTimeseriesParamsAggInterval string

const (
	NetflowTimeseriesParamsAggInterval15m NetflowTimeseriesParamsAggInterval = "15m"
	NetflowTimeseriesParamsAggInterval1h  NetflowTimeseriesParamsAggInterval = "1h"
	NetflowTimeseriesParamsAggInterval1d  NetflowTimeseriesParamsAggInterval = "1d"
	NetflowTimeseriesParamsAggInterval1w  NetflowTimeseriesParamsAggInterval = "1w"
)

func (r NetflowTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case NetflowTimeseriesParamsAggInterval15m, NetflowTimeseriesParamsAggInterval1h, NetflowTimeseriesParamsAggInterval1d, NetflowTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type NetflowTimeseriesParamsFormat string

const (
	NetflowTimeseriesParamsFormatJson NetflowTimeseriesParamsFormat = "JSON"
	NetflowTimeseriesParamsFormatCsv  NetflowTimeseriesParamsFormat = "CSV"
)

func (r NetflowTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case NetflowTimeseriesParamsFormatJson, NetflowTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type NetflowTimeseriesParamsNormalization string

const (
	NetflowTimeseriesParamsNormalizationPercentageChange NetflowTimeseriesParamsNormalization = "PERCENTAGE_CHANGE"
	NetflowTimeseriesParamsNormalizationMin0Max          NetflowTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r NetflowTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case NetflowTimeseriesParamsNormalizationPercentageChange, NetflowTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type NetflowTimeseriesParamsProduct string

const (
	NetflowTimeseriesParamsProductHTTP NetflowTimeseriesParamsProduct = "HTTP"
	NetflowTimeseriesParamsProductAll  NetflowTimeseriesParamsProduct = "ALL"
)

func (r NetflowTimeseriesParamsProduct) IsKnown() bool {
	switch r {
	case NetflowTimeseriesParamsProductHTTP, NetflowTimeseriesParamsProductAll:
		return true
	}
	return false
}

type NetflowTimeseriesResponseEnvelope struct {
	Result  NetflowTimeseriesResponse             `json:"result" api:"required"`
	Success bool                                  `json:"success" api:"required"`
	JSON    netflowTimeseriesResponseEnvelopeJSON `json:"-"`
}

// netflowTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [NetflowTimeseriesResponseEnvelope]
type netflowTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
