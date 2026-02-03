// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"fmt"
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

// AS112Service contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAS112Service] method instead.
type AS112Service struct {
	Options          []option.RequestOption
	Summary          *AS112SummaryService
	TimeseriesGroups *AS112TimeseriesGroupService
	Top              *AS112TopService
}

// NewAS112Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAS112Service(opts ...option.RequestOption) (r *AS112Service) {
	r = &AS112Service{}
	r.Options = opts
	r.Summary = NewAS112SummaryService(opts...)
	r.TimeseriesGroups = NewAS112TimeseriesGroupService(opts...)
	r.Top = NewAS112TopService(opts...)
	return
}

// Retrieves the distribution of AS112 queries by the specified dimension.
func (r *AS112Service) SummaryV2(ctx context.Context, dimension AS112SummaryV2ParamsDimension, query AS112SummaryV2Params, opts ...option.RequestOption) (res *AS112SummaryV2Response, err error) {
	var env AS112SummaryV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/as112/summary/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the AS112 DNS queries over time.
func (r *AS112Service) Timeseries(ctx context.Context, query AS112TimeseriesParams, opts ...option.RequestOption) (res *AS112TimeseriesResponse, err error) {
	var env AS112TimeseriesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/as112/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of AS112 queries grouped by dimension over time.
func (r *AS112Service) TimeseriesGroupsV2(ctx context.Context, dimension AS112TimeseriesGroupsV2ParamsDimension, query AS112TimeseriesGroupsV2Params, opts ...option.RequestOption) (res *AS112TimeseriesGroupsV2Response, err error) {
	var env AS112TimeseriesGroupsV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/as112/timeseries_groups/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AS112SummaryV2Response struct {
	// Metadata for the results.
	Meta     AS112SummaryV2ResponseMeta `json:"meta,required"`
	Summary0 map[string]string          `json:"summary_0,required"`
	JSON     as112SummaryV2ResponseJSON `json:"-"`
}

// as112SummaryV2ResponseJSON contains the JSON metadata for the struct
// [AS112SummaryV2Response]
type as112SummaryV2ResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112SummaryV2ResponseMeta struct {
	ConfidenceInfo AS112SummaryV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AS112SummaryV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112SummaryV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112SummaryV2ResponseMetaUnit `json:"units,required"`
	JSON  as112SummaryV2ResponseMetaJSON   `json:"-"`
}

// as112SummaryV2ResponseMetaJSON contains the JSON metadata for the struct
// [AS112SummaryV2ResponseMeta]
type as112SummaryV2ResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112SummaryV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112SummaryV2ResponseMetaConfidenceInfo struct {
	Annotations []AS112SummaryV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                        `json:"level,required"`
	JSON  as112SummaryV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112SummaryV2ResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [AS112SummaryV2ResponseMetaConfidenceInfo]
type as112SummaryV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112SummaryV2ResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                        `json:"description,required"`
	EndDate     time.Time                                                     `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                              `json:"startDate,required" format:"date-time"`
	JSON            as112SummaryV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112SummaryV2ResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [AS112SummaryV2ResponseMetaConfidenceInfoAnnotation]
type as112SummaryV2ResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112SummaryV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll                AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots               AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos                AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet                AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AS112SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AS112SummaryV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                               `json:"startTime,required" format:"date-time"`
	JSON      as112SummaryV2ResponseMetaDateRangeJSON `json:"-"`
}

// as112SummaryV2ResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [AS112SummaryV2ResponseMetaDateRange]
type as112SummaryV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112SummaryV2ResponseMetaNormalization string

const (
	AS112SummaryV2ResponseMetaNormalizationPercentage           AS112SummaryV2ResponseMetaNormalization = "PERCENTAGE"
	AS112SummaryV2ResponseMetaNormalizationMin0Max              AS112SummaryV2ResponseMetaNormalization = "MIN0_MAX"
	AS112SummaryV2ResponseMetaNormalizationMinMax               AS112SummaryV2ResponseMetaNormalization = "MIN_MAX"
	AS112SummaryV2ResponseMetaNormalizationRawValues            AS112SummaryV2ResponseMetaNormalization = "RAW_VALUES"
	AS112SummaryV2ResponseMetaNormalizationPercentageChange     AS112SummaryV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112SummaryV2ResponseMetaNormalizationRollingAverage       AS112SummaryV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112SummaryV2ResponseMetaNormalizationOverlappedPercentage AS112SummaryV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112SummaryV2ResponseMetaNormalizationRatio                AS112SummaryV2ResponseMetaNormalization = "RATIO"
)

func (r AS112SummaryV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112SummaryV2ResponseMetaNormalizationPercentage, AS112SummaryV2ResponseMetaNormalizationMin0Max, AS112SummaryV2ResponseMetaNormalizationMinMax, AS112SummaryV2ResponseMetaNormalizationRawValues, AS112SummaryV2ResponseMetaNormalizationPercentageChange, AS112SummaryV2ResponseMetaNormalizationRollingAverage, AS112SummaryV2ResponseMetaNormalizationOverlappedPercentage, AS112SummaryV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112SummaryV2ResponseMetaUnit struct {
	Name  string                             `json:"name,required"`
	Value string                             `json:"value,required"`
	JSON  as112SummaryV2ResponseMetaUnitJSON `json:"-"`
}

// as112SummaryV2ResponseMetaUnitJSON contains the JSON metadata for the struct
// [AS112SummaryV2ResponseMetaUnit]
type as112SummaryV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesResponse struct {
	// Metadata for the results.
	Meta        AS112TimeseriesResponseMeta        `json:"meta,required"`
	ExtraFields map[string]AS112TimeseriesResponse `json:"-,extras"`
	JSON        as112TimeseriesResponseJSON        `json:"-"`
}

// as112TimeseriesResponseJSON contains the JSON metadata for the struct
// [AS112TimeseriesResponse]
type as112TimeseriesResponseJSON struct {
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112TimeseriesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AS112TimeseriesResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AS112TimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AS112TimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112TimeseriesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112TimeseriesResponseMetaUnit `json:"units,required"`
	JSON  as112TimeseriesResponseMetaJSON   `json:"-"`
}

// as112TimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [AS112TimeseriesResponseMeta]
type as112TimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112TimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesResponseMetaAggInterval string

const (
	AS112TimeseriesResponseMetaAggIntervalFifteenMinutes AS112TimeseriesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AS112TimeseriesResponseMetaAggIntervalOneHour        AS112TimeseriesResponseMetaAggInterval = "ONE_HOUR"
	AS112TimeseriesResponseMetaAggIntervalOneDay         AS112TimeseriesResponseMetaAggInterval = "ONE_DAY"
	AS112TimeseriesResponseMetaAggIntervalOneWeek        AS112TimeseriesResponseMetaAggInterval = "ONE_WEEK"
	AS112TimeseriesResponseMetaAggIntervalOneMonth       AS112TimeseriesResponseMetaAggInterval = "ONE_MONTH"
)

func (r AS112TimeseriesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesResponseMetaAggIntervalFifteenMinutes, AS112TimeseriesResponseMetaAggIntervalOneHour, AS112TimeseriesResponseMetaAggIntervalOneDay, AS112TimeseriesResponseMetaAggIntervalOneWeek, AS112TimeseriesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AS112TimeseriesResponseMetaConfidenceInfo struct {
	Annotations []AS112TimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                         `json:"level,required"`
	JSON  as112TimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112TimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [AS112TimeseriesResponseMetaConfidenceInfo]
type as112TimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112TimeseriesResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                         `json:"description,required"`
	EndDate     time.Time                                                      `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                               `json:"startDate,required" format:"date-time"`
	JSON            as112TimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112TimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [AS112TimeseriesResponseMetaConfidenceInfoAnnotation]
type as112TimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112TimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll                AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots               AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos                AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet                AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCT, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AS112TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AS112TimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                `json:"startTime,required" format:"date-time"`
	JSON      as112TimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// as112TimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [AS112TimeseriesResponseMetaDateRange]
type as112TimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112TimeseriesResponseMetaNormalization string

const (
	AS112TimeseriesResponseMetaNormalizationPercentage           AS112TimeseriesResponseMetaNormalization = "PERCENTAGE"
	AS112TimeseriesResponseMetaNormalizationMin0Max              AS112TimeseriesResponseMetaNormalization = "MIN0_MAX"
	AS112TimeseriesResponseMetaNormalizationMinMax               AS112TimeseriesResponseMetaNormalization = "MIN_MAX"
	AS112TimeseriesResponseMetaNormalizationRawValues            AS112TimeseriesResponseMetaNormalization = "RAW_VALUES"
	AS112TimeseriesResponseMetaNormalizationPercentageChange     AS112TimeseriesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112TimeseriesResponseMetaNormalizationRollingAverage       AS112TimeseriesResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112TimeseriesResponseMetaNormalizationOverlappedPercentage AS112TimeseriesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112TimeseriesResponseMetaNormalizationRatio                AS112TimeseriesResponseMetaNormalization = "RATIO"
)

func (r AS112TimeseriesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112TimeseriesResponseMetaNormalizationPercentage, AS112TimeseriesResponseMetaNormalizationMin0Max, AS112TimeseriesResponseMetaNormalizationMinMax, AS112TimeseriesResponseMetaNormalizationRawValues, AS112TimeseriesResponseMetaNormalizationPercentageChange, AS112TimeseriesResponseMetaNormalizationRollingAverage, AS112TimeseriesResponseMetaNormalizationOverlappedPercentage, AS112TimeseriesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112TimeseriesResponseMetaUnit struct {
	Name  string                              `json:"name,required"`
	Value string                              `json:"value,required"`
	JSON  as112TimeseriesResponseMetaUnitJSON `json:"-"`
}

// as112TimeseriesResponseMetaUnitJSON contains the JSON metadata for the struct
// [AS112TimeseriesResponseMetaUnit]
type as112TimeseriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupsV2Response struct {
	// Metadata for the results.
	Meta   AS112TimeseriesGroupsV2ResponseMeta   `json:"meta,required"`
	Serie0 AS112TimeseriesGroupsV2ResponseSerie0 `json:"serie_0,required"`
	JSON   as112TimeseriesGroupsV2ResponseJSON   `json:"-"`
}

// as112TimeseriesGroupsV2ResponseJSON contains the JSON metadata for the struct
// [AS112TimeseriesGroupsV2Response]
type as112TimeseriesGroupsV2ResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupsV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupsV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112TimeseriesGroupsV2ResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AS112TimeseriesGroupsV2ResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AS112TimeseriesGroupsV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AS112TimeseriesGroupsV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112TimeseriesGroupsV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112TimeseriesGroupsV2ResponseMetaUnit `json:"units,required"`
	JSON  as112TimeseriesGroupsV2ResponseMetaJSON   `json:"-"`
}

// as112TimeseriesGroupsV2ResponseMetaJSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupsV2ResponseMeta]
type as112TimeseriesGroupsV2ResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112TimeseriesGroupsV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupsV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupsV2ResponseMetaAggInterval string

const (
	AS112TimeseriesGroupsV2ResponseMetaAggIntervalFifteenMinutes AS112TimeseriesGroupsV2ResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AS112TimeseriesGroupsV2ResponseMetaAggIntervalOneHour        AS112TimeseriesGroupsV2ResponseMetaAggInterval = "ONE_HOUR"
	AS112TimeseriesGroupsV2ResponseMetaAggIntervalOneDay         AS112TimeseriesGroupsV2ResponseMetaAggInterval = "ONE_DAY"
	AS112TimeseriesGroupsV2ResponseMetaAggIntervalOneWeek        AS112TimeseriesGroupsV2ResponseMetaAggInterval = "ONE_WEEK"
	AS112TimeseriesGroupsV2ResponseMetaAggIntervalOneMonth       AS112TimeseriesGroupsV2ResponseMetaAggInterval = "ONE_MONTH"
)

func (r AS112TimeseriesGroupsV2ResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupsV2ResponseMetaAggIntervalFifteenMinutes, AS112TimeseriesGroupsV2ResponseMetaAggIntervalOneHour, AS112TimeseriesGroupsV2ResponseMetaAggIntervalOneDay, AS112TimeseriesGroupsV2ResponseMetaAggIntervalOneWeek, AS112TimeseriesGroupsV2ResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AS112TimeseriesGroupsV2ResponseMetaConfidenceInfo struct {
	Annotations []AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                 `json:"level,required"`
	JSON  as112TimeseriesGroupsV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112TimeseriesGroupsV2ResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AS112TimeseriesGroupsV2ResponseMetaConfidenceInfo]
type as112TimeseriesGroupsV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupsV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupsV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                 `json:"description,required"`
	EndDate     time.Time                                                              `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                            `json:"isInstantaneous,required"`
	LinkedURL       string                                                          `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                       `json:"startDate,required" format:"date-time"`
	JSON            as112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation]
type as112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll                AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots               AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos                AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet                AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AS112TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AS112TimeseriesGroupsV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      as112TimeseriesGroupsV2ResponseMetaDateRangeJSON `json:"-"`
}

// as112TimeseriesGroupsV2ResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupsV2ResponseMetaDateRange]
type as112TimeseriesGroupsV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupsV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupsV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112TimeseriesGroupsV2ResponseMetaNormalization string

const (
	AS112TimeseriesGroupsV2ResponseMetaNormalizationPercentage           AS112TimeseriesGroupsV2ResponseMetaNormalization = "PERCENTAGE"
	AS112TimeseriesGroupsV2ResponseMetaNormalizationMin0Max              AS112TimeseriesGroupsV2ResponseMetaNormalization = "MIN0_MAX"
	AS112TimeseriesGroupsV2ResponseMetaNormalizationMinMax               AS112TimeseriesGroupsV2ResponseMetaNormalization = "MIN_MAX"
	AS112TimeseriesGroupsV2ResponseMetaNormalizationRawValues            AS112TimeseriesGroupsV2ResponseMetaNormalization = "RAW_VALUES"
	AS112TimeseriesGroupsV2ResponseMetaNormalizationPercentageChange     AS112TimeseriesGroupsV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112TimeseriesGroupsV2ResponseMetaNormalizationRollingAverage       AS112TimeseriesGroupsV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112TimeseriesGroupsV2ResponseMetaNormalizationOverlappedPercentage AS112TimeseriesGroupsV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112TimeseriesGroupsV2ResponseMetaNormalizationRatio                AS112TimeseriesGroupsV2ResponseMetaNormalization = "RATIO"
)

func (r AS112TimeseriesGroupsV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupsV2ResponseMetaNormalizationPercentage, AS112TimeseriesGroupsV2ResponseMetaNormalizationMin0Max, AS112TimeseriesGroupsV2ResponseMetaNormalizationMinMax, AS112TimeseriesGroupsV2ResponseMetaNormalizationRawValues, AS112TimeseriesGroupsV2ResponseMetaNormalizationPercentageChange, AS112TimeseriesGroupsV2ResponseMetaNormalizationRollingAverage, AS112TimeseriesGroupsV2ResponseMetaNormalizationOverlappedPercentage, AS112TimeseriesGroupsV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112TimeseriesGroupsV2ResponseMetaUnit struct {
	Name  string                                      `json:"name,required"`
	Value string                                      `json:"value,required"`
	JSON  as112TimeseriesGroupsV2ResponseMetaUnitJSON `json:"-"`
}

// as112TimeseriesGroupsV2ResponseMetaUnitJSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupsV2ResponseMetaUnit]
type as112TimeseriesGroupsV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupsV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupsV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupsV2ResponseSerie0 struct {
	Timestamps  []time.Time                               `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                       `json:"-,extras"`
	JSON        as112TimeseriesGroupsV2ResponseSerie0JSON `json:"-"`
}

// as112TimeseriesGroupsV2ResponseSerie0JSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupsV2ResponseSerie0]
type as112TimeseriesGroupsV2ResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupsV2ResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupsV2ResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AS112SummaryV2Params struct {
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
	Format param.Field[AS112SummaryV2ParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[[]AS112SummaryV2ParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[[]AS112SummaryV2ParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[[]AS112SummaryV2ParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [AS112SummaryV2Params]'s query parameters as `url.Values`.
func (r AS112SummaryV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type AS112SummaryV2ParamsDimension string

const (
	AS112SummaryV2ParamsDimensionDNSSEC       AS112SummaryV2ParamsDimension = "DNSSEC"
	AS112SummaryV2ParamsDimensionEdns         AS112SummaryV2ParamsDimension = "EDNS"
	AS112SummaryV2ParamsDimensionIPVersion    AS112SummaryV2ParamsDimension = "IP_VERSION"
	AS112SummaryV2ParamsDimensionProtocol     AS112SummaryV2ParamsDimension = "PROTOCOL"
	AS112SummaryV2ParamsDimensionQueryType    AS112SummaryV2ParamsDimension = "QUERY_TYPE"
	AS112SummaryV2ParamsDimensionResponseCode AS112SummaryV2ParamsDimension = "RESPONSE_CODE"
)

func (r AS112SummaryV2ParamsDimension) IsKnown() bool {
	switch r {
	case AS112SummaryV2ParamsDimensionDNSSEC, AS112SummaryV2ParamsDimensionEdns, AS112SummaryV2ParamsDimensionIPVersion, AS112SummaryV2ParamsDimensionProtocol, AS112SummaryV2ParamsDimensionQueryType, AS112SummaryV2ParamsDimensionResponseCode:
		return true
	}
	return false
}

// Format in which results will be returned.
type AS112SummaryV2ParamsFormat string

const (
	AS112SummaryV2ParamsFormatJson AS112SummaryV2ParamsFormat = "JSON"
	AS112SummaryV2ParamsFormatCsv  AS112SummaryV2ParamsFormat = "CSV"
)

func (r AS112SummaryV2ParamsFormat) IsKnown() bool {
	switch r {
	case AS112SummaryV2ParamsFormatJson, AS112SummaryV2ParamsFormatCsv:
		return true
	}
	return false
}

type AS112SummaryV2ParamsProtocol string

const (
	AS112SummaryV2ParamsProtocolUdp   AS112SummaryV2ParamsProtocol = "UDP"
	AS112SummaryV2ParamsProtocolTCP   AS112SummaryV2ParamsProtocol = "TCP"
	AS112SummaryV2ParamsProtocolHTTPS AS112SummaryV2ParamsProtocol = "HTTPS"
	AS112SummaryV2ParamsProtocolTLS   AS112SummaryV2ParamsProtocol = "TLS"
)

func (r AS112SummaryV2ParamsProtocol) IsKnown() bool {
	switch r {
	case AS112SummaryV2ParamsProtocolUdp, AS112SummaryV2ParamsProtocolTCP, AS112SummaryV2ParamsProtocolHTTPS, AS112SummaryV2ParamsProtocolTLS:
		return true
	}
	return false
}

type AS112SummaryV2ParamsQueryType string

const (
	AS112SummaryV2ParamsQueryTypeA          AS112SummaryV2ParamsQueryType = "A"
	AS112SummaryV2ParamsQueryTypeAAAA       AS112SummaryV2ParamsQueryType = "AAAA"
	AS112SummaryV2ParamsQueryTypeA6         AS112SummaryV2ParamsQueryType = "A6"
	AS112SummaryV2ParamsQueryTypeAfsdb      AS112SummaryV2ParamsQueryType = "AFSDB"
	AS112SummaryV2ParamsQueryTypeAny        AS112SummaryV2ParamsQueryType = "ANY"
	AS112SummaryV2ParamsQueryTypeApl        AS112SummaryV2ParamsQueryType = "APL"
	AS112SummaryV2ParamsQueryTypeAtma       AS112SummaryV2ParamsQueryType = "ATMA"
	AS112SummaryV2ParamsQueryTypeAXFR       AS112SummaryV2ParamsQueryType = "AXFR"
	AS112SummaryV2ParamsQueryTypeCAA        AS112SummaryV2ParamsQueryType = "CAA"
	AS112SummaryV2ParamsQueryTypeCdnskey    AS112SummaryV2ParamsQueryType = "CDNSKEY"
	AS112SummaryV2ParamsQueryTypeCds        AS112SummaryV2ParamsQueryType = "CDS"
	AS112SummaryV2ParamsQueryTypeCERT       AS112SummaryV2ParamsQueryType = "CERT"
	AS112SummaryV2ParamsQueryTypeCNAME      AS112SummaryV2ParamsQueryType = "CNAME"
	AS112SummaryV2ParamsQueryTypeCsync      AS112SummaryV2ParamsQueryType = "CSYNC"
	AS112SummaryV2ParamsQueryTypeDhcid      AS112SummaryV2ParamsQueryType = "DHCID"
	AS112SummaryV2ParamsQueryTypeDlv        AS112SummaryV2ParamsQueryType = "DLV"
	AS112SummaryV2ParamsQueryTypeDname      AS112SummaryV2ParamsQueryType = "DNAME"
	AS112SummaryV2ParamsQueryTypeDNSKEY     AS112SummaryV2ParamsQueryType = "DNSKEY"
	AS112SummaryV2ParamsQueryTypeDoa        AS112SummaryV2ParamsQueryType = "DOA"
	AS112SummaryV2ParamsQueryTypeDS         AS112SummaryV2ParamsQueryType = "DS"
	AS112SummaryV2ParamsQueryTypeEid        AS112SummaryV2ParamsQueryType = "EID"
	AS112SummaryV2ParamsQueryTypeEui48      AS112SummaryV2ParamsQueryType = "EUI48"
	AS112SummaryV2ParamsQueryTypeEui64      AS112SummaryV2ParamsQueryType = "EUI64"
	AS112SummaryV2ParamsQueryTypeGpos       AS112SummaryV2ParamsQueryType = "GPOS"
	AS112SummaryV2ParamsQueryTypeGid        AS112SummaryV2ParamsQueryType = "GID"
	AS112SummaryV2ParamsQueryTypeHinfo      AS112SummaryV2ParamsQueryType = "HINFO"
	AS112SummaryV2ParamsQueryTypeHip        AS112SummaryV2ParamsQueryType = "HIP"
	AS112SummaryV2ParamsQueryTypeHTTPS      AS112SummaryV2ParamsQueryType = "HTTPS"
	AS112SummaryV2ParamsQueryTypeIpseckey   AS112SummaryV2ParamsQueryType = "IPSECKEY"
	AS112SummaryV2ParamsQueryTypeIsdn       AS112SummaryV2ParamsQueryType = "ISDN"
	AS112SummaryV2ParamsQueryTypeIxfr       AS112SummaryV2ParamsQueryType = "IXFR"
	AS112SummaryV2ParamsQueryTypeKey        AS112SummaryV2ParamsQueryType = "KEY"
	AS112SummaryV2ParamsQueryTypeKx         AS112SummaryV2ParamsQueryType = "KX"
	AS112SummaryV2ParamsQueryTypeL32        AS112SummaryV2ParamsQueryType = "L32"
	AS112SummaryV2ParamsQueryTypeL64        AS112SummaryV2ParamsQueryType = "L64"
	AS112SummaryV2ParamsQueryTypeLOC        AS112SummaryV2ParamsQueryType = "LOC"
	AS112SummaryV2ParamsQueryTypeLp         AS112SummaryV2ParamsQueryType = "LP"
	AS112SummaryV2ParamsQueryTypeMaila      AS112SummaryV2ParamsQueryType = "MAILA"
	AS112SummaryV2ParamsQueryTypeMailb      AS112SummaryV2ParamsQueryType = "MAILB"
	AS112SummaryV2ParamsQueryTypeMB         AS112SummaryV2ParamsQueryType = "MB"
	AS112SummaryV2ParamsQueryTypeMd         AS112SummaryV2ParamsQueryType = "MD"
	AS112SummaryV2ParamsQueryTypeMf         AS112SummaryV2ParamsQueryType = "MF"
	AS112SummaryV2ParamsQueryTypeMg         AS112SummaryV2ParamsQueryType = "MG"
	AS112SummaryV2ParamsQueryTypeMinfo      AS112SummaryV2ParamsQueryType = "MINFO"
	AS112SummaryV2ParamsQueryTypeMr         AS112SummaryV2ParamsQueryType = "MR"
	AS112SummaryV2ParamsQueryTypeMX         AS112SummaryV2ParamsQueryType = "MX"
	AS112SummaryV2ParamsQueryTypeNAPTR      AS112SummaryV2ParamsQueryType = "NAPTR"
	AS112SummaryV2ParamsQueryTypeNb         AS112SummaryV2ParamsQueryType = "NB"
	AS112SummaryV2ParamsQueryTypeNbstat     AS112SummaryV2ParamsQueryType = "NBSTAT"
	AS112SummaryV2ParamsQueryTypeNid        AS112SummaryV2ParamsQueryType = "NID"
	AS112SummaryV2ParamsQueryTypeNimloc     AS112SummaryV2ParamsQueryType = "NIMLOC"
	AS112SummaryV2ParamsQueryTypeNinfo      AS112SummaryV2ParamsQueryType = "NINFO"
	AS112SummaryV2ParamsQueryTypeNS         AS112SummaryV2ParamsQueryType = "NS"
	AS112SummaryV2ParamsQueryTypeNsap       AS112SummaryV2ParamsQueryType = "NSAP"
	AS112SummaryV2ParamsQueryTypeNsec       AS112SummaryV2ParamsQueryType = "NSEC"
	AS112SummaryV2ParamsQueryTypeNsec3      AS112SummaryV2ParamsQueryType = "NSEC3"
	AS112SummaryV2ParamsQueryTypeNsec3Param AS112SummaryV2ParamsQueryType = "NSEC3PARAM"
	AS112SummaryV2ParamsQueryTypeNull       AS112SummaryV2ParamsQueryType = "NULL"
	AS112SummaryV2ParamsQueryTypeNxt        AS112SummaryV2ParamsQueryType = "NXT"
	AS112SummaryV2ParamsQueryTypeOpenpgpkey AS112SummaryV2ParamsQueryType = "OPENPGPKEY"
	AS112SummaryV2ParamsQueryTypeOpt        AS112SummaryV2ParamsQueryType = "OPT"
	AS112SummaryV2ParamsQueryTypePTR        AS112SummaryV2ParamsQueryType = "PTR"
	AS112SummaryV2ParamsQueryTypePx         AS112SummaryV2ParamsQueryType = "PX"
	AS112SummaryV2ParamsQueryTypeRkey       AS112SummaryV2ParamsQueryType = "RKEY"
	AS112SummaryV2ParamsQueryTypeRp         AS112SummaryV2ParamsQueryType = "RP"
	AS112SummaryV2ParamsQueryTypeRrsig      AS112SummaryV2ParamsQueryType = "RRSIG"
	AS112SummaryV2ParamsQueryTypeRt         AS112SummaryV2ParamsQueryType = "RT"
	AS112SummaryV2ParamsQueryTypeSig        AS112SummaryV2ParamsQueryType = "SIG"
	AS112SummaryV2ParamsQueryTypeSink       AS112SummaryV2ParamsQueryType = "SINK"
	AS112SummaryV2ParamsQueryTypeSMIMEA     AS112SummaryV2ParamsQueryType = "SMIMEA"
	AS112SummaryV2ParamsQueryTypeSOA        AS112SummaryV2ParamsQueryType = "SOA"
	AS112SummaryV2ParamsQueryTypeSPF        AS112SummaryV2ParamsQueryType = "SPF"
	AS112SummaryV2ParamsQueryTypeSRV        AS112SummaryV2ParamsQueryType = "SRV"
	AS112SummaryV2ParamsQueryTypeSSHFP      AS112SummaryV2ParamsQueryType = "SSHFP"
	AS112SummaryV2ParamsQueryTypeSVCB       AS112SummaryV2ParamsQueryType = "SVCB"
	AS112SummaryV2ParamsQueryTypeTa         AS112SummaryV2ParamsQueryType = "TA"
	AS112SummaryV2ParamsQueryTypeTalink     AS112SummaryV2ParamsQueryType = "TALINK"
	AS112SummaryV2ParamsQueryTypeTkey       AS112SummaryV2ParamsQueryType = "TKEY"
	AS112SummaryV2ParamsQueryTypeTLSA       AS112SummaryV2ParamsQueryType = "TLSA"
	AS112SummaryV2ParamsQueryTypeTSIG       AS112SummaryV2ParamsQueryType = "TSIG"
	AS112SummaryV2ParamsQueryTypeTXT        AS112SummaryV2ParamsQueryType = "TXT"
	AS112SummaryV2ParamsQueryTypeUinfo      AS112SummaryV2ParamsQueryType = "UINFO"
	AS112SummaryV2ParamsQueryTypeUID        AS112SummaryV2ParamsQueryType = "UID"
	AS112SummaryV2ParamsQueryTypeUnspec     AS112SummaryV2ParamsQueryType = "UNSPEC"
	AS112SummaryV2ParamsQueryTypeURI        AS112SummaryV2ParamsQueryType = "URI"
	AS112SummaryV2ParamsQueryTypeWks        AS112SummaryV2ParamsQueryType = "WKS"
	AS112SummaryV2ParamsQueryTypeX25        AS112SummaryV2ParamsQueryType = "X25"
	AS112SummaryV2ParamsQueryTypeZonemd     AS112SummaryV2ParamsQueryType = "ZONEMD"
)

func (r AS112SummaryV2ParamsQueryType) IsKnown() bool {
	switch r {
	case AS112SummaryV2ParamsQueryTypeA, AS112SummaryV2ParamsQueryTypeAAAA, AS112SummaryV2ParamsQueryTypeA6, AS112SummaryV2ParamsQueryTypeAfsdb, AS112SummaryV2ParamsQueryTypeAny, AS112SummaryV2ParamsQueryTypeApl, AS112SummaryV2ParamsQueryTypeAtma, AS112SummaryV2ParamsQueryTypeAXFR, AS112SummaryV2ParamsQueryTypeCAA, AS112SummaryV2ParamsQueryTypeCdnskey, AS112SummaryV2ParamsQueryTypeCds, AS112SummaryV2ParamsQueryTypeCERT, AS112SummaryV2ParamsQueryTypeCNAME, AS112SummaryV2ParamsQueryTypeCsync, AS112SummaryV2ParamsQueryTypeDhcid, AS112SummaryV2ParamsQueryTypeDlv, AS112SummaryV2ParamsQueryTypeDname, AS112SummaryV2ParamsQueryTypeDNSKEY, AS112SummaryV2ParamsQueryTypeDoa, AS112SummaryV2ParamsQueryTypeDS, AS112SummaryV2ParamsQueryTypeEid, AS112SummaryV2ParamsQueryTypeEui48, AS112SummaryV2ParamsQueryTypeEui64, AS112SummaryV2ParamsQueryTypeGpos, AS112SummaryV2ParamsQueryTypeGid, AS112SummaryV2ParamsQueryTypeHinfo, AS112SummaryV2ParamsQueryTypeHip, AS112SummaryV2ParamsQueryTypeHTTPS, AS112SummaryV2ParamsQueryTypeIpseckey, AS112SummaryV2ParamsQueryTypeIsdn, AS112SummaryV2ParamsQueryTypeIxfr, AS112SummaryV2ParamsQueryTypeKey, AS112SummaryV2ParamsQueryTypeKx, AS112SummaryV2ParamsQueryTypeL32, AS112SummaryV2ParamsQueryTypeL64, AS112SummaryV2ParamsQueryTypeLOC, AS112SummaryV2ParamsQueryTypeLp, AS112SummaryV2ParamsQueryTypeMaila, AS112SummaryV2ParamsQueryTypeMailb, AS112SummaryV2ParamsQueryTypeMB, AS112SummaryV2ParamsQueryTypeMd, AS112SummaryV2ParamsQueryTypeMf, AS112SummaryV2ParamsQueryTypeMg, AS112SummaryV2ParamsQueryTypeMinfo, AS112SummaryV2ParamsQueryTypeMr, AS112SummaryV2ParamsQueryTypeMX, AS112SummaryV2ParamsQueryTypeNAPTR, AS112SummaryV2ParamsQueryTypeNb, AS112SummaryV2ParamsQueryTypeNbstat, AS112SummaryV2ParamsQueryTypeNid, AS112SummaryV2ParamsQueryTypeNimloc, AS112SummaryV2ParamsQueryTypeNinfo, AS112SummaryV2ParamsQueryTypeNS, AS112SummaryV2ParamsQueryTypeNsap, AS112SummaryV2ParamsQueryTypeNsec, AS112SummaryV2ParamsQueryTypeNsec3, AS112SummaryV2ParamsQueryTypeNsec3Param, AS112SummaryV2ParamsQueryTypeNull, AS112SummaryV2ParamsQueryTypeNxt, AS112SummaryV2ParamsQueryTypeOpenpgpkey, AS112SummaryV2ParamsQueryTypeOpt, AS112SummaryV2ParamsQueryTypePTR, AS112SummaryV2ParamsQueryTypePx, AS112SummaryV2ParamsQueryTypeRkey, AS112SummaryV2ParamsQueryTypeRp, AS112SummaryV2ParamsQueryTypeRrsig, AS112SummaryV2ParamsQueryTypeRt, AS112SummaryV2ParamsQueryTypeSig, AS112SummaryV2ParamsQueryTypeSink, AS112SummaryV2ParamsQueryTypeSMIMEA, AS112SummaryV2ParamsQueryTypeSOA, AS112SummaryV2ParamsQueryTypeSPF, AS112SummaryV2ParamsQueryTypeSRV, AS112SummaryV2ParamsQueryTypeSSHFP, AS112SummaryV2ParamsQueryTypeSVCB, AS112SummaryV2ParamsQueryTypeTa, AS112SummaryV2ParamsQueryTypeTalink, AS112SummaryV2ParamsQueryTypeTkey, AS112SummaryV2ParamsQueryTypeTLSA, AS112SummaryV2ParamsQueryTypeTSIG, AS112SummaryV2ParamsQueryTypeTXT, AS112SummaryV2ParamsQueryTypeUinfo, AS112SummaryV2ParamsQueryTypeUID, AS112SummaryV2ParamsQueryTypeUnspec, AS112SummaryV2ParamsQueryTypeURI, AS112SummaryV2ParamsQueryTypeWks, AS112SummaryV2ParamsQueryTypeX25, AS112SummaryV2ParamsQueryTypeZonemd:
		return true
	}
	return false
}

type AS112SummaryV2ParamsResponseCode string

const (
	AS112SummaryV2ParamsResponseCodeNoerror   AS112SummaryV2ParamsResponseCode = "NOERROR"
	AS112SummaryV2ParamsResponseCodeFormerr   AS112SummaryV2ParamsResponseCode = "FORMERR"
	AS112SummaryV2ParamsResponseCodeServfail  AS112SummaryV2ParamsResponseCode = "SERVFAIL"
	AS112SummaryV2ParamsResponseCodeNxdomain  AS112SummaryV2ParamsResponseCode = "NXDOMAIN"
	AS112SummaryV2ParamsResponseCodeNotimp    AS112SummaryV2ParamsResponseCode = "NOTIMP"
	AS112SummaryV2ParamsResponseCodeRefused   AS112SummaryV2ParamsResponseCode = "REFUSED"
	AS112SummaryV2ParamsResponseCodeYxdomain  AS112SummaryV2ParamsResponseCode = "YXDOMAIN"
	AS112SummaryV2ParamsResponseCodeYxrrset   AS112SummaryV2ParamsResponseCode = "YXRRSET"
	AS112SummaryV2ParamsResponseCodeNxrrset   AS112SummaryV2ParamsResponseCode = "NXRRSET"
	AS112SummaryV2ParamsResponseCodeNotauth   AS112SummaryV2ParamsResponseCode = "NOTAUTH"
	AS112SummaryV2ParamsResponseCodeNotzone   AS112SummaryV2ParamsResponseCode = "NOTZONE"
	AS112SummaryV2ParamsResponseCodeBadsig    AS112SummaryV2ParamsResponseCode = "BADSIG"
	AS112SummaryV2ParamsResponseCodeBadkey    AS112SummaryV2ParamsResponseCode = "BADKEY"
	AS112SummaryV2ParamsResponseCodeBadtime   AS112SummaryV2ParamsResponseCode = "BADTIME"
	AS112SummaryV2ParamsResponseCodeBadmode   AS112SummaryV2ParamsResponseCode = "BADMODE"
	AS112SummaryV2ParamsResponseCodeBadname   AS112SummaryV2ParamsResponseCode = "BADNAME"
	AS112SummaryV2ParamsResponseCodeBadalg    AS112SummaryV2ParamsResponseCode = "BADALG"
	AS112SummaryV2ParamsResponseCodeBadtrunc  AS112SummaryV2ParamsResponseCode = "BADTRUNC"
	AS112SummaryV2ParamsResponseCodeBadcookie AS112SummaryV2ParamsResponseCode = "BADCOOKIE"
)

func (r AS112SummaryV2ParamsResponseCode) IsKnown() bool {
	switch r {
	case AS112SummaryV2ParamsResponseCodeNoerror, AS112SummaryV2ParamsResponseCodeFormerr, AS112SummaryV2ParamsResponseCodeServfail, AS112SummaryV2ParamsResponseCodeNxdomain, AS112SummaryV2ParamsResponseCodeNotimp, AS112SummaryV2ParamsResponseCodeRefused, AS112SummaryV2ParamsResponseCodeYxdomain, AS112SummaryV2ParamsResponseCodeYxrrset, AS112SummaryV2ParamsResponseCodeNxrrset, AS112SummaryV2ParamsResponseCodeNotauth, AS112SummaryV2ParamsResponseCodeNotzone, AS112SummaryV2ParamsResponseCodeBadsig, AS112SummaryV2ParamsResponseCodeBadkey, AS112SummaryV2ParamsResponseCodeBadtime, AS112SummaryV2ParamsResponseCodeBadmode, AS112SummaryV2ParamsResponseCodeBadname, AS112SummaryV2ParamsResponseCodeBadalg, AS112SummaryV2ParamsResponseCodeBadtrunc, AS112SummaryV2ParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type AS112SummaryV2ResponseEnvelope struct {
	Result  AS112SummaryV2Response             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    as112SummaryV2ResponseEnvelopeJSON `json:"-"`
}

// as112SummaryV2ResponseEnvelopeJSON contains the JSON metadata for the struct
// [AS112SummaryV2ResponseEnvelope]
type as112SummaryV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112SummaryV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112SummaryV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AS112TimeseriesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AS112TimeseriesParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[[]AS112TimeseriesParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[[]AS112TimeseriesParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[[]AS112TimeseriesParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [AS112TimeseriesParams]'s query parameters as `url.Values`.
func (r AS112TimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesParamsAggInterval string

const (
	AS112TimeseriesParamsAggInterval15m AS112TimeseriesParamsAggInterval = "15m"
	AS112TimeseriesParamsAggInterval1h  AS112TimeseriesParamsAggInterval = "1h"
	AS112TimeseriesParamsAggInterval1d  AS112TimeseriesParamsAggInterval = "1d"
	AS112TimeseriesParamsAggInterval1w  AS112TimeseriesParamsAggInterval = "1w"
)

func (r AS112TimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesParamsAggInterval15m, AS112TimeseriesParamsAggInterval1h, AS112TimeseriesParamsAggInterval1d, AS112TimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AS112TimeseriesParamsFormat string

const (
	AS112TimeseriesParamsFormatJson AS112TimeseriesParamsFormat = "JSON"
	AS112TimeseriesParamsFormatCsv  AS112TimeseriesParamsFormat = "CSV"
)

func (r AS112TimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case AS112TimeseriesParamsFormatJson, AS112TimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type AS112TimeseriesParamsProtocol string

const (
	AS112TimeseriesParamsProtocolUdp   AS112TimeseriesParamsProtocol = "UDP"
	AS112TimeseriesParamsProtocolTCP   AS112TimeseriesParamsProtocol = "TCP"
	AS112TimeseriesParamsProtocolHTTPS AS112TimeseriesParamsProtocol = "HTTPS"
	AS112TimeseriesParamsProtocolTLS   AS112TimeseriesParamsProtocol = "TLS"
)

func (r AS112TimeseriesParamsProtocol) IsKnown() bool {
	switch r {
	case AS112TimeseriesParamsProtocolUdp, AS112TimeseriesParamsProtocolTCP, AS112TimeseriesParamsProtocolHTTPS, AS112TimeseriesParamsProtocolTLS:
		return true
	}
	return false
}

type AS112TimeseriesParamsQueryType string

const (
	AS112TimeseriesParamsQueryTypeA          AS112TimeseriesParamsQueryType = "A"
	AS112TimeseriesParamsQueryTypeAAAA       AS112TimeseriesParamsQueryType = "AAAA"
	AS112TimeseriesParamsQueryTypeA6         AS112TimeseriesParamsQueryType = "A6"
	AS112TimeseriesParamsQueryTypeAfsdb      AS112TimeseriesParamsQueryType = "AFSDB"
	AS112TimeseriesParamsQueryTypeAny        AS112TimeseriesParamsQueryType = "ANY"
	AS112TimeseriesParamsQueryTypeApl        AS112TimeseriesParamsQueryType = "APL"
	AS112TimeseriesParamsQueryTypeAtma       AS112TimeseriesParamsQueryType = "ATMA"
	AS112TimeseriesParamsQueryTypeAXFR       AS112TimeseriesParamsQueryType = "AXFR"
	AS112TimeseriesParamsQueryTypeCAA        AS112TimeseriesParamsQueryType = "CAA"
	AS112TimeseriesParamsQueryTypeCdnskey    AS112TimeseriesParamsQueryType = "CDNSKEY"
	AS112TimeseriesParamsQueryTypeCds        AS112TimeseriesParamsQueryType = "CDS"
	AS112TimeseriesParamsQueryTypeCERT       AS112TimeseriesParamsQueryType = "CERT"
	AS112TimeseriesParamsQueryTypeCNAME      AS112TimeseriesParamsQueryType = "CNAME"
	AS112TimeseriesParamsQueryTypeCsync      AS112TimeseriesParamsQueryType = "CSYNC"
	AS112TimeseriesParamsQueryTypeDhcid      AS112TimeseriesParamsQueryType = "DHCID"
	AS112TimeseriesParamsQueryTypeDlv        AS112TimeseriesParamsQueryType = "DLV"
	AS112TimeseriesParamsQueryTypeDname      AS112TimeseriesParamsQueryType = "DNAME"
	AS112TimeseriesParamsQueryTypeDNSKEY     AS112TimeseriesParamsQueryType = "DNSKEY"
	AS112TimeseriesParamsQueryTypeDoa        AS112TimeseriesParamsQueryType = "DOA"
	AS112TimeseriesParamsQueryTypeDS         AS112TimeseriesParamsQueryType = "DS"
	AS112TimeseriesParamsQueryTypeEid        AS112TimeseriesParamsQueryType = "EID"
	AS112TimeseriesParamsQueryTypeEui48      AS112TimeseriesParamsQueryType = "EUI48"
	AS112TimeseriesParamsQueryTypeEui64      AS112TimeseriesParamsQueryType = "EUI64"
	AS112TimeseriesParamsQueryTypeGpos       AS112TimeseriesParamsQueryType = "GPOS"
	AS112TimeseriesParamsQueryTypeGid        AS112TimeseriesParamsQueryType = "GID"
	AS112TimeseriesParamsQueryTypeHinfo      AS112TimeseriesParamsQueryType = "HINFO"
	AS112TimeseriesParamsQueryTypeHip        AS112TimeseriesParamsQueryType = "HIP"
	AS112TimeseriesParamsQueryTypeHTTPS      AS112TimeseriesParamsQueryType = "HTTPS"
	AS112TimeseriesParamsQueryTypeIpseckey   AS112TimeseriesParamsQueryType = "IPSECKEY"
	AS112TimeseriesParamsQueryTypeIsdn       AS112TimeseriesParamsQueryType = "ISDN"
	AS112TimeseriesParamsQueryTypeIxfr       AS112TimeseriesParamsQueryType = "IXFR"
	AS112TimeseriesParamsQueryTypeKey        AS112TimeseriesParamsQueryType = "KEY"
	AS112TimeseriesParamsQueryTypeKx         AS112TimeseriesParamsQueryType = "KX"
	AS112TimeseriesParamsQueryTypeL32        AS112TimeseriesParamsQueryType = "L32"
	AS112TimeseriesParamsQueryTypeL64        AS112TimeseriesParamsQueryType = "L64"
	AS112TimeseriesParamsQueryTypeLOC        AS112TimeseriesParamsQueryType = "LOC"
	AS112TimeseriesParamsQueryTypeLp         AS112TimeseriesParamsQueryType = "LP"
	AS112TimeseriesParamsQueryTypeMaila      AS112TimeseriesParamsQueryType = "MAILA"
	AS112TimeseriesParamsQueryTypeMailb      AS112TimeseriesParamsQueryType = "MAILB"
	AS112TimeseriesParamsQueryTypeMB         AS112TimeseriesParamsQueryType = "MB"
	AS112TimeseriesParamsQueryTypeMd         AS112TimeseriesParamsQueryType = "MD"
	AS112TimeseriesParamsQueryTypeMf         AS112TimeseriesParamsQueryType = "MF"
	AS112TimeseriesParamsQueryTypeMg         AS112TimeseriesParamsQueryType = "MG"
	AS112TimeseriesParamsQueryTypeMinfo      AS112TimeseriesParamsQueryType = "MINFO"
	AS112TimeseriesParamsQueryTypeMr         AS112TimeseriesParamsQueryType = "MR"
	AS112TimeseriesParamsQueryTypeMX         AS112TimeseriesParamsQueryType = "MX"
	AS112TimeseriesParamsQueryTypeNAPTR      AS112TimeseriesParamsQueryType = "NAPTR"
	AS112TimeseriesParamsQueryTypeNb         AS112TimeseriesParamsQueryType = "NB"
	AS112TimeseriesParamsQueryTypeNbstat     AS112TimeseriesParamsQueryType = "NBSTAT"
	AS112TimeseriesParamsQueryTypeNid        AS112TimeseriesParamsQueryType = "NID"
	AS112TimeseriesParamsQueryTypeNimloc     AS112TimeseriesParamsQueryType = "NIMLOC"
	AS112TimeseriesParamsQueryTypeNinfo      AS112TimeseriesParamsQueryType = "NINFO"
	AS112TimeseriesParamsQueryTypeNS         AS112TimeseriesParamsQueryType = "NS"
	AS112TimeseriesParamsQueryTypeNsap       AS112TimeseriesParamsQueryType = "NSAP"
	AS112TimeseriesParamsQueryTypeNsec       AS112TimeseriesParamsQueryType = "NSEC"
	AS112TimeseriesParamsQueryTypeNsec3      AS112TimeseriesParamsQueryType = "NSEC3"
	AS112TimeseriesParamsQueryTypeNsec3Param AS112TimeseriesParamsQueryType = "NSEC3PARAM"
	AS112TimeseriesParamsQueryTypeNull       AS112TimeseriesParamsQueryType = "NULL"
	AS112TimeseriesParamsQueryTypeNxt        AS112TimeseriesParamsQueryType = "NXT"
	AS112TimeseriesParamsQueryTypeOpenpgpkey AS112TimeseriesParamsQueryType = "OPENPGPKEY"
	AS112TimeseriesParamsQueryTypeOpt        AS112TimeseriesParamsQueryType = "OPT"
	AS112TimeseriesParamsQueryTypePTR        AS112TimeseriesParamsQueryType = "PTR"
	AS112TimeseriesParamsQueryTypePx         AS112TimeseriesParamsQueryType = "PX"
	AS112TimeseriesParamsQueryTypeRkey       AS112TimeseriesParamsQueryType = "RKEY"
	AS112TimeseriesParamsQueryTypeRp         AS112TimeseriesParamsQueryType = "RP"
	AS112TimeseriesParamsQueryTypeRrsig      AS112TimeseriesParamsQueryType = "RRSIG"
	AS112TimeseriesParamsQueryTypeRt         AS112TimeseriesParamsQueryType = "RT"
	AS112TimeseriesParamsQueryTypeSig        AS112TimeseriesParamsQueryType = "SIG"
	AS112TimeseriesParamsQueryTypeSink       AS112TimeseriesParamsQueryType = "SINK"
	AS112TimeseriesParamsQueryTypeSMIMEA     AS112TimeseriesParamsQueryType = "SMIMEA"
	AS112TimeseriesParamsQueryTypeSOA        AS112TimeseriesParamsQueryType = "SOA"
	AS112TimeseriesParamsQueryTypeSPF        AS112TimeseriesParamsQueryType = "SPF"
	AS112TimeseriesParamsQueryTypeSRV        AS112TimeseriesParamsQueryType = "SRV"
	AS112TimeseriesParamsQueryTypeSSHFP      AS112TimeseriesParamsQueryType = "SSHFP"
	AS112TimeseriesParamsQueryTypeSVCB       AS112TimeseriesParamsQueryType = "SVCB"
	AS112TimeseriesParamsQueryTypeTa         AS112TimeseriesParamsQueryType = "TA"
	AS112TimeseriesParamsQueryTypeTalink     AS112TimeseriesParamsQueryType = "TALINK"
	AS112TimeseriesParamsQueryTypeTkey       AS112TimeseriesParamsQueryType = "TKEY"
	AS112TimeseriesParamsQueryTypeTLSA       AS112TimeseriesParamsQueryType = "TLSA"
	AS112TimeseriesParamsQueryTypeTSIG       AS112TimeseriesParamsQueryType = "TSIG"
	AS112TimeseriesParamsQueryTypeTXT        AS112TimeseriesParamsQueryType = "TXT"
	AS112TimeseriesParamsQueryTypeUinfo      AS112TimeseriesParamsQueryType = "UINFO"
	AS112TimeseriesParamsQueryTypeUID        AS112TimeseriesParamsQueryType = "UID"
	AS112TimeseriesParamsQueryTypeUnspec     AS112TimeseriesParamsQueryType = "UNSPEC"
	AS112TimeseriesParamsQueryTypeURI        AS112TimeseriesParamsQueryType = "URI"
	AS112TimeseriesParamsQueryTypeWks        AS112TimeseriesParamsQueryType = "WKS"
	AS112TimeseriesParamsQueryTypeX25        AS112TimeseriesParamsQueryType = "X25"
	AS112TimeseriesParamsQueryTypeZonemd     AS112TimeseriesParamsQueryType = "ZONEMD"
)

func (r AS112TimeseriesParamsQueryType) IsKnown() bool {
	switch r {
	case AS112TimeseriesParamsQueryTypeA, AS112TimeseriesParamsQueryTypeAAAA, AS112TimeseriesParamsQueryTypeA6, AS112TimeseriesParamsQueryTypeAfsdb, AS112TimeseriesParamsQueryTypeAny, AS112TimeseriesParamsQueryTypeApl, AS112TimeseriesParamsQueryTypeAtma, AS112TimeseriesParamsQueryTypeAXFR, AS112TimeseriesParamsQueryTypeCAA, AS112TimeseriesParamsQueryTypeCdnskey, AS112TimeseriesParamsQueryTypeCds, AS112TimeseriesParamsQueryTypeCERT, AS112TimeseriesParamsQueryTypeCNAME, AS112TimeseriesParamsQueryTypeCsync, AS112TimeseriesParamsQueryTypeDhcid, AS112TimeseriesParamsQueryTypeDlv, AS112TimeseriesParamsQueryTypeDname, AS112TimeseriesParamsQueryTypeDNSKEY, AS112TimeseriesParamsQueryTypeDoa, AS112TimeseriesParamsQueryTypeDS, AS112TimeseriesParamsQueryTypeEid, AS112TimeseriesParamsQueryTypeEui48, AS112TimeseriesParamsQueryTypeEui64, AS112TimeseriesParamsQueryTypeGpos, AS112TimeseriesParamsQueryTypeGid, AS112TimeseriesParamsQueryTypeHinfo, AS112TimeseriesParamsQueryTypeHip, AS112TimeseriesParamsQueryTypeHTTPS, AS112TimeseriesParamsQueryTypeIpseckey, AS112TimeseriesParamsQueryTypeIsdn, AS112TimeseriesParamsQueryTypeIxfr, AS112TimeseriesParamsQueryTypeKey, AS112TimeseriesParamsQueryTypeKx, AS112TimeseriesParamsQueryTypeL32, AS112TimeseriesParamsQueryTypeL64, AS112TimeseriesParamsQueryTypeLOC, AS112TimeseriesParamsQueryTypeLp, AS112TimeseriesParamsQueryTypeMaila, AS112TimeseriesParamsQueryTypeMailb, AS112TimeseriesParamsQueryTypeMB, AS112TimeseriesParamsQueryTypeMd, AS112TimeseriesParamsQueryTypeMf, AS112TimeseriesParamsQueryTypeMg, AS112TimeseriesParamsQueryTypeMinfo, AS112TimeseriesParamsQueryTypeMr, AS112TimeseriesParamsQueryTypeMX, AS112TimeseriesParamsQueryTypeNAPTR, AS112TimeseriesParamsQueryTypeNb, AS112TimeseriesParamsQueryTypeNbstat, AS112TimeseriesParamsQueryTypeNid, AS112TimeseriesParamsQueryTypeNimloc, AS112TimeseriesParamsQueryTypeNinfo, AS112TimeseriesParamsQueryTypeNS, AS112TimeseriesParamsQueryTypeNsap, AS112TimeseriesParamsQueryTypeNsec, AS112TimeseriesParamsQueryTypeNsec3, AS112TimeseriesParamsQueryTypeNsec3Param, AS112TimeseriesParamsQueryTypeNull, AS112TimeseriesParamsQueryTypeNxt, AS112TimeseriesParamsQueryTypeOpenpgpkey, AS112TimeseriesParamsQueryTypeOpt, AS112TimeseriesParamsQueryTypePTR, AS112TimeseriesParamsQueryTypePx, AS112TimeseriesParamsQueryTypeRkey, AS112TimeseriesParamsQueryTypeRp, AS112TimeseriesParamsQueryTypeRrsig, AS112TimeseriesParamsQueryTypeRt, AS112TimeseriesParamsQueryTypeSig, AS112TimeseriesParamsQueryTypeSink, AS112TimeseriesParamsQueryTypeSMIMEA, AS112TimeseriesParamsQueryTypeSOA, AS112TimeseriesParamsQueryTypeSPF, AS112TimeseriesParamsQueryTypeSRV, AS112TimeseriesParamsQueryTypeSSHFP, AS112TimeseriesParamsQueryTypeSVCB, AS112TimeseriesParamsQueryTypeTa, AS112TimeseriesParamsQueryTypeTalink, AS112TimeseriesParamsQueryTypeTkey, AS112TimeseriesParamsQueryTypeTLSA, AS112TimeseriesParamsQueryTypeTSIG, AS112TimeseriesParamsQueryTypeTXT, AS112TimeseriesParamsQueryTypeUinfo, AS112TimeseriesParamsQueryTypeUID, AS112TimeseriesParamsQueryTypeUnspec, AS112TimeseriesParamsQueryTypeURI, AS112TimeseriesParamsQueryTypeWks, AS112TimeseriesParamsQueryTypeX25, AS112TimeseriesParamsQueryTypeZonemd:
		return true
	}
	return false
}

type AS112TimeseriesParamsResponseCode string

const (
	AS112TimeseriesParamsResponseCodeNoerror   AS112TimeseriesParamsResponseCode = "NOERROR"
	AS112TimeseriesParamsResponseCodeFormerr   AS112TimeseriesParamsResponseCode = "FORMERR"
	AS112TimeseriesParamsResponseCodeServfail  AS112TimeseriesParamsResponseCode = "SERVFAIL"
	AS112TimeseriesParamsResponseCodeNxdomain  AS112TimeseriesParamsResponseCode = "NXDOMAIN"
	AS112TimeseriesParamsResponseCodeNotimp    AS112TimeseriesParamsResponseCode = "NOTIMP"
	AS112TimeseriesParamsResponseCodeRefused   AS112TimeseriesParamsResponseCode = "REFUSED"
	AS112TimeseriesParamsResponseCodeYxdomain  AS112TimeseriesParamsResponseCode = "YXDOMAIN"
	AS112TimeseriesParamsResponseCodeYxrrset   AS112TimeseriesParamsResponseCode = "YXRRSET"
	AS112TimeseriesParamsResponseCodeNxrrset   AS112TimeseriesParamsResponseCode = "NXRRSET"
	AS112TimeseriesParamsResponseCodeNotauth   AS112TimeseriesParamsResponseCode = "NOTAUTH"
	AS112TimeseriesParamsResponseCodeNotzone   AS112TimeseriesParamsResponseCode = "NOTZONE"
	AS112TimeseriesParamsResponseCodeBadsig    AS112TimeseriesParamsResponseCode = "BADSIG"
	AS112TimeseriesParamsResponseCodeBadkey    AS112TimeseriesParamsResponseCode = "BADKEY"
	AS112TimeseriesParamsResponseCodeBadtime   AS112TimeseriesParamsResponseCode = "BADTIME"
	AS112TimeseriesParamsResponseCodeBadmode   AS112TimeseriesParamsResponseCode = "BADMODE"
	AS112TimeseriesParamsResponseCodeBadname   AS112TimeseriesParamsResponseCode = "BADNAME"
	AS112TimeseriesParamsResponseCodeBadalg    AS112TimeseriesParamsResponseCode = "BADALG"
	AS112TimeseriesParamsResponseCodeBadtrunc  AS112TimeseriesParamsResponseCode = "BADTRUNC"
	AS112TimeseriesParamsResponseCodeBadcookie AS112TimeseriesParamsResponseCode = "BADCOOKIE"
)

func (r AS112TimeseriesParamsResponseCode) IsKnown() bool {
	switch r {
	case AS112TimeseriesParamsResponseCodeNoerror, AS112TimeseriesParamsResponseCodeFormerr, AS112TimeseriesParamsResponseCodeServfail, AS112TimeseriesParamsResponseCodeNxdomain, AS112TimeseriesParamsResponseCodeNotimp, AS112TimeseriesParamsResponseCodeRefused, AS112TimeseriesParamsResponseCodeYxdomain, AS112TimeseriesParamsResponseCodeYxrrset, AS112TimeseriesParamsResponseCodeNxrrset, AS112TimeseriesParamsResponseCodeNotauth, AS112TimeseriesParamsResponseCodeNotzone, AS112TimeseriesParamsResponseCodeBadsig, AS112TimeseriesParamsResponseCodeBadkey, AS112TimeseriesParamsResponseCodeBadtime, AS112TimeseriesParamsResponseCodeBadmode, AS112TimeseriesParamsResponseCodeBadname, AS112TimeseriesParamsResponseCodeBadalg, AS112TimeseriesParamsResponseCodeBadtrunc, AS112TimeseriesParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type AS112TimeseriesResponseEnvelope struct {
	Result  AS112TimeseriesResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    as112TimeseriesResponseEnvelopeJSON `json:"-"`
}

// as112TimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [AS112TimeseriesResponseEnvelope]
type as112TimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupsV2Params struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AS112TimeseriesGroupsV2ParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AS112TimeseriesGroupsV2ParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[[]AS112TimeseriesGroupsV2ParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[[]AS112TimeseriesGroupsV2ParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[[]AS112TimeseriesGroupsV2ParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [AS112TimeseriesGroupsV2Params]'s query parameters as
// `url.Values`.
func (r AS112TimeseriesGroupsV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type AS112TimeseriesGroupsV2ParamsDimension string

const (
	AS112TimeseriesGroupsV2ParamsDimensionDNSSEC       AS112TimeseriesGroupsV2ParamsDimension = "DNSSEC"
	AS112TimeseriesGroupsV2ParamsDimensionEdns         AS112TimeseriesGroupsV2ParamsDimension = "EDNS"
	AS112TimeseriesGroupsV2ParamsDimensionIPVersion    AS112TimeseriesGroupsV2ParamsDimension = "IP_VERSION"
	AS112TimeseriesGroupsV2ParamsDimensionProtocol     AS112TimeseriesGroupsV2ParamsDimension = "PROTOCOL"
	AS112TimeseriesGroupsV2ParamsDimensionQueryType    AS112TimeseriesGroupsV2ParamsDimension = "QUERY_TYPE"
	AS112TimeseriesGroupsV2ParamsDimensionResponseCode AS112TimeseriesGroupsV2ParamsDimension = "RESPONSE_CODE"
)

func (r AS112TimeseriesGroupsV2ParamsDimension) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupsV2ParamsDimensionDNSSEC, AS112TimeseriesGroupsV2ParamsDimensionEdns, AS112TimeseriesGroupsV2ParamsDimensionIPVersion, AS112TimeseriesGroupsV2ParamsDimensionProtocol, AS112TimeseriesGroupsV2ParamsDimensionQueryType, AS112TimeseriesGroupsV2ParamsDimensionResponseCode:
		return true
	}
	return false
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupsV2ParamsAggInterval string

const (
	AS112TimeseriesGroupsV2ParamsAggInterval15m AS112TimeseriesGroupsV2ParamsAggInterval = "15m"
	AS112TimeseriesGroupsV2ParamsAggInterval1h  AS112TimeseriesGroupsV2ParamsAggInterval = "1h"
	AS112TimeseriesGroupsV2ParamsAggInterval1d  AS112TimeseriesGroupsV2ParamsAggInterval = "1d"
	AS112TimeseriesGroupsV2ParamsAggInterval1w  AS112TimeseriesGroupsV2ParamsAggInterval = "1w"
)

func (r AS112TimeseriesGroupsV2ParamsAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupsV2ParamsAggInterval15m, AS112TimeseriesGroupsV2ParamsAggInterval1h, AS112TimeseriesGroupsV2ParamsAggInterval1d, AS112TimeseriesGroupsV2ParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AS112TimeseriesGroupsV2ParamsFormat string

const (
	AS112TimeseriesGroupsV2ParamsFormatJson AS112TimeseriesGroupsV2ParamsFormat = "JSON"
	AS112TimeseriesGroupsV2ParamsFormatCsv  AS112TimeseriesGroupsV2ParamsFormat = "CSV"
)

func (r AS112TimeseriesGroupsV2ParamsFormat) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupsV2ParamsFormatJson, AS112TimeseriesGroupsV2ParamsFormatCsv:
		return true
	}
	return false
}

type AS112TimeseriesGroupsV2ParamsProtocol string

const (
	AS112TimeseriesGroupsV2ParamsProtocolUdp   AS112TimeseriesGroupsV2ParamsProtocol = "UDP"
	AS112TimeseriesGroupsV2ParamsProtocolTCP   AS112TimeseriesGroupsV2ParamsProtocol = "TCP"
	AS112TimeseriesGroupsV2ParamsProtocolHTTPS AS112TimeseriesGroupsV2ParamsProtocol = "HTTPS"
	AS112TimeseriesGroupsV2ParamsProtocolTLS   AS112TimeseriesGroupsV2ParamsProtocol = "TLS"
)

func (r AS112TimeseriesGroupsV2ParamsProtocol) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupsV2ParamsProtocolUdp, AS112TimeseriesGroupsV2ParamsProtocolTCP, AS112TimeseriesGroupsV2ParamsProtocolHTTPS, AS112TimeseriesGroupsV2ParamsProtocolTLS:
		return true
	}
	return false
}

type AS112TimeseriesGroupsV2ParamsQueryType string

const (
	AS112TimeseriesGroupsV2ParamsQueryTypeA          AS112TimeseriesGroupsV2ParamsQueryType = "A"
	AS112TimeseriesGroupsV2ParamsQueryTypeAAAA       AS112TimeseriesGroupsV2ParamsQueryType = "AAAA"
	AS112TimeseriesGroupsV2ParamsQueryTypeA6         AS112TimeseriesGroupsV2ParamsQueryType = "A6"
	AS112TimeseriesGroupsV2ParamsQueryTypeAfsdb      AS112TimeseriesGroupsV2ParamsQueryType = "AFSDB"
	AS112TimeseriesGroupsV2ParamsQueryTypeAny        AS112TimeseriesGroupsV2ParamsQueryType = "ANY"
	AS112TimeseriesGroupsV2ParamsQueryTypeApl        AS112TimeseriesGroupsV2ParamsQueryType = "APL"
	AS112TimeseriesGroupsV2ParamsQueryTypeAtma       AS112TimeseriesGroupsV2ParamsQueryType = "ATMA"
	AS112TimeseriesGroupsV2ParamsQueryTypeAXFR       AS112TimeseriesGroupsV2ParamsQueryType = "AXFR"
	AS112TimeseriesGroupsV2ParamsQueryTypeCAA        AS112TimeseriesGroupsV2ParamsQueryType = "CAA"
	AS112TimeseriesGroupsV2ParamsQueryTypeCdnskey    AS112TimeseriesGroupsV2ParamsQueryType = "CDNSKEY"
	AS112TimeseriesGroupsV2ParamsQueryTypeCds        AS112TimeseriesGroupsV2ParamsQueryType = "CDS"
	AS112TimeseriesGroupsV2ParamsQueryTypeCERT       AS112TimeseriesGroupsV2ParamsQueryType = "CERT"
	AS112TimeseriesGroupsV2ParamsQueryTypeCNAME      AS112TimeseriesGroupsV2ParamsQueryType = "CNAME"
	AS112TimeseriesGroupsV2ParamsQueryTypeCsync      AS112TimeseriesGroupsV2ParamsQueryType = "CSYNC"
	AS112TimeseriesGroupsV2ParamsQueryTypeDhcid      AS112TimeseriesGroupsV2ParamsQueryType = "DHCID"
	AS112TimeseriesGroupsV2ParamsQueryTypeDlv        AS112TimeseriesGroupsV2ParamsQueryType = "DLV"
	AS112TimeseriesGroupsV2ParamsQueryTypeDname      AS112TimeseriesGroupsV2ParamsQueryType = "DNAME"
	AS112TimeseriesGroupsV2ParamsQueryTypeDNSKEY     AS112TimeseriesGroupsV2ParamsQueryType = "DNSKEY"
	AS112TimeseriesGroupsV2ParamsQueryTypeDoa        AS112TimeseriesGroupsV2ParamsQueryType = "DOA"
	AS112TimeseriesGroupsV2ParamsQueryTypeDS         AS112TimeseriesGroupsV2ParamsQueryType = "DS"
	AS112TimeseriesGroupsV2ParamsQueryTypeEid        AS112TimeseriesGroupsV2ParamsQueryType = "EID"
	AS112TimeseriesGroupsV2ParamsQueryTypeEui48      AS112TimeseriesGroupsV2ParamsQueryType = "EUI48"
	AS112TimeseriesGroupsV2ParamsQueryTypeEui64      AS112TimeseriesGroupsV2ParamsQueryType = "EUI64"
	AS112TimeseriesGroupsV2ParamsQueryTypeGpos       AS112TimeseriesGroupsV2ParamsQueryType = "GPOS"
	AS112TimeseriesGroupsV2ParamsQueryTypeGid        AS112TimeseriesGroupsV2ParamsQueryType = "GID"
	AS112TimeseriesGroupsV2ParamsQueryTypeHinfo      AS112TimeseriesGroupsV2ParamsQueryType = "HINFO"
	AS112TimeseriesGroupsV2ParamsQueryTypeHip        AS112TimeseriesGroupsV2ParamsQueryType = "HIP"
	AS112TimeseriesGroupsV2ParamsQueryTypeHTTPS      AS112TimeseriesGroupsV2ParamsQueryType = "HTTPS"
	AS112TimeseriesGroupsV2ParamsQueryTypeIpseckey   AS112TimeseriesGroupsV2ParamsQueryType = "IPSECKEY"
	AS112TimeseriesGroupsV2ParamsQueryTypeIsdn       AS112TimeseriesGroupsV2ParamsQueryType = "ISDN"
	AS112TimeseriesGroupsV2ParamsQueryTypeIxfr       AS112TimeseriesGroupsV2ParamsQueryType = "IXFR"
	AS112TimeseriesGroupsV2ParamsQueryTypeKey        AS112TimeseriesGroupsV2ParamsQueryType = "KEY"
	AS112TimeseriesGroupsV2ParamsQueryTypeKx         AS112TimeseriesGroupsV2ParamsQueryType = "KX"
	AS112TimeseriesGroupsV2ParamsQueryTypeL32        AS112TimeseriesGroupsV2ParamsQueryType = "L32"
	AS112TimeseriesGroupsV2ParamsQueryTypeL64        AS112TimeseriesGroupsV2ParamsQueryType = "L64"
	AS112TimeseriesGroupsV2ParamsQueryTypeLOC        AS112TimeseriesGroupsV2ParamsQueryType = "LOC"
	AS112TimeseriesGroupsV2ParamsQueryTypeLp         AS112TimeseriesGroupsV2ParamsQueryType = "LP"
	AS112TimeseriesGroupsV2ParamsQueryTypeMaila      AS112TimeseriesGroupsV2ParamsQueryType = "MAILA"
	AS112TimeseriesGroupsV2ParamsQueryTypeMailb      AS112TimeseriesGroupsV2ParamsQueryType = "MAILB"
	AS112TimeseriesGroupsV2ParamsQueryTypeMB         AS112TimeseriesGroupsV2ParamsQueryType = "MB"
	AS112TimeseriesGroupsV2ParamsQueryTypeMd         AS112TimeseriesGroupsV2ParamsQueryType = "MD"
	AS112TimeseriesGroupsV2ParamsQueryTypeMf         AS112TimeseriesGroupsV2ParamsQueryType = "MF"
	AS112TimeseriesGroupsV2ParamsQueryTypeMg         AS112TimeseriesGroupsV2ParamsQueryType = "MG"
	AS112TimeseriesGroupsV2ParamsQueryTypeMinfo      AS112TimeseriesGroupsV2ParamsQueryType = "MINFO"
	AS112TimeseriesGroupsV2ParamsQueryTypeMr         AS112TimeseriesGroupsV2ParamsQueryType = "MR"
	AS112TimeseriesGroupsV2ParamsQueryTypeMX         AS112TimeseriesGroupsV2ParamsQueryType = "MX"
	AS112TimeseriesGroupsV2ParamsQueryTypeNAPTR      AS112TimeseriesGroupsV2ParamsQueryType = "NAPTR"
	AS112TimeseriesGroupsV2ParamsQueryTypeNb         AS112TimeseriesGroupsV2ParamsQueryType = "NB"
	AS112TimeseriesGroupsV2ParamsQueryTypeNbstat     AS112TimeseriesGroupsV2ParamsQueryType = "NBSTAT"
	AS112TimeseriesGroupsV2ParamsQueryTypeNid        AS112TimeseriesGroupsV2ParamsQueryType = "NID"
	AS112TimeseriesGroupsV2ParamsQueryTypeNimloc     AS112TimeseriesGroupsV2ParamsQueryType = "NIMLOC"
	AS112TimeseriesGroupsV2ParamsQueryTypeNinfo      AS112TimeseriesGroupsV2ParamsQueryType = "NINFO"
	AS112TimeseriesGroupsV2ParamsQueryTypeNS         AS112TimeseriesGroupsV2ParamsQueryType = "NS"
	AS112TimeseriesGroupsV2ParamsQueryTypeNsap       AS112TimeseriesGroupsV2ParamsQueryType = "NSAP"
	AS112TimeseriesGroupsV2ParamsQueryTypeNsec       AS112TimeseriesGroupsV2ParamsQueryType = "NSEC"
	AS112TimeseriesGroupsV2ParamsQueryTypeNsec3      AS112TimeseriesGroupsV2ParamsQueryType = "NSEC3"
	AS112TimeseriesGroupsV2ParamsQueryTypeNsec3Param AS112TimeseriesGroupsV2ParamsQueryType = "NSEC3PARAM"
	AS112TimeseriesGroupsV2ParamsQueryTypeNull       AS112TimeseriesGroupsV2ParamsQueryType = "NULL"
	AS112TimeseriesGroupsV2ParamsQueryTypeNxt        AS112TimeseriesGroupsV2ParamsQueryType = "NXT"
	AS112TimeseriesGroupsV2ParamsQueryTypeOpenpgpkey AS112TimeseriesGroupsV2ParamsQueryType = "OPENPGPKEY"
	AS112TimeseriesGroupsV2ParamsQueryTypeOpt        AS112TimeseriesGroupsV2ParamsQueryType = "OPT"
	AS112TimeseriesGroupsV2ParamsQueryTypePTR        AS112TimeseriesGroupsV2ParamsQueryType = "PTR"
	AS112TimeseriesGroupsV2ParamsQueryTypePx         AS112TimeseriesGroupsV2ParamsQueryType = "PX"
	AS112TimeseriesGroupsV2ParamsQueryTypeRkey       AS112TimeseriesGroupsV2ParamsQueryType = "RKEY"
	AS112TimeseriesGroupsV2ParamsQueryTypeRp         AS112TimeseriesGroupsV2ParamsQueryType = "RP"
	AS112TimeseriesGroupsV2ParamsQueryTypeRrsig      AS112TimeseriesGroupsV2ParamsQueryType = "RRSIG"
	AS112TimeseriesGroupsV2ParamsQueryTypeRt         AS112TimeseriesGroupsV2ParamsQueryType = "RT"
	AS112TimeseriesGroupsV2ParamsQueryTypeSig        AS112TimeseriesGroupsV2ParamsQueryType = "SIG"
	AS112TimeseriesGroupsV2ParamsQueryTypeSink       AS112TimeseriesGroupsV2ParamsQueryType = "SINK"
	AS112TimeseriesGroupsV2ParamsQueryTypeSMIMEA     AS112TimeseriesGroupsV2ParamsQueryType = "SMIMEA"
	AS112TimeseriesGroupsV2ParamsQueryTypeSOA        AS112TimeseriesGroupsV2ParamsQueryType = "SOA"
	AS112TimeseriesGroupsV2ParamsQueryTypeSPF        AS112TimeseriesGroupsV2ParamsQueryType = "SPF"
	AS112TimeseriesGroupsV2ParamsQueryTypeSRV        AS112TimeseriesGroupsV2ParamsQueryType = "SRV"
	AS112TimeseriesGroupsV2ParamsQueryTypeSSHFP      AS112TimeseriesGroupsV2ParamsQueryType = "SSHFP"
	AS112TimeseriesGroupsV2ParamsQueryTypeSVCB       AS112TimeseriesGroupsV2ParamsQueryType = "SVCB"
	AS112TimeseriesGroupsV2ParamsQueryTypeTa         AS112TimeseriesGroupsV2ParamsQueryType = "TA"
	AS112TimeseriesGroupsV2ParamsQueryTypeTalink     AS112TimeseriesGroupsV2ParamsQueryType = "TALINK"
	AS112TimeseriesGroupsV2ParamsQueryTypeTkey       AS112TimeseriesGroupsV2ParamsQueryType = "TKEY"
	AS112TimeseriesGroupsV2ParamsQueryTypeTLSA       AS112TimeseriesGroupsV2ParamsQueryType = "TLSA"
	AS112TimeseriesGroupsV2ParamsQueryTypeTSIG       AS112TimeseriesGroupsV2ParamsQueryType = "TSIG"
	AS112TimeseriesGroupsV2ParamsQueryTypeTXT        AS112TimeseriesGroupsV2ParamsQueryType = "TXT"
	AS112TimeseriesGroupsV2ParamsQueryTypeUinfo      AS112TimeseriesGroupsV2ParamsQueryType = "UINFO"
	AS112TimeseriesGroupsV2ParamsQueryTypeUID        AS112TimeseriesGroupsV2ParamsQueryType = "UID"
	AS112TimeseriesGroupsV2ParamsQueryTypeUnspec     AS112TimeseriesGroupsV2ParamsQueryType = "UNSPEC"
	AS112TimeseriesGroupsV2ParamsQueryTypeURI        AS112TimeseriesGroupsV2ParamsQueryType = "URI"
	AS112TimeseriesGroupsV2ParamsQueryTypeWks        AS112TimeseriesGroupsV2ParamsQueryType = "WKS"
	AS112TimeseriesGroupsV2ParamsQueryTypeX25        AS112TimeseriesGroupsV2ParamsQueryType = "X25"
	AS112TimeseriesGroupsV2ParamsQueryTypeZonemd     AS112TimeseriesGroupsV2ParamsQueryType = "ZONEMD"
)

func (r AS112TimeseriesGroupsV2ParamsQueryType) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupsV2ParamsQueryTypeA, AS112TimeseriesGroupsV2ParamsQueryTypeAAAA, AS112TimeseriesGroupsV2ParamsQueryTypeA6, AS112TimeseriesGroupsV2ParamsQueryTypeAfsdb, AS112TimeseriesGroupsV2ParamsQueryTypeAny, AS112TimeseriesGroupsV2ParamsQueryTypeApl, AS112TimeseriesGroupsV2ParamsQueryTypeAtma, AS112TimeseriesGroupsV2ParamsQueryTypeAXFR, AS112TimeseriesGroupsV2ParamsQueryTypeCAA, AS112TimeseriesGroupsV2ParamsQueryTypeCdnskey, AS112TimeseriesGroupsV2ParamsQueryTypeCds, AS112TimeseriesGroupsV2ParamsQueryTypeCERT, AS112TimeseriesGroupsV2ParamsQueryTypeCNAME, AS112TimeseriesGroupsV2ParamsQueryTypeCsync, AS112TimeseriesGroupsV2ParamsQueryTypeDhcid, AS112TimeseriesGroupsV2ParamsQueryTypeDlv, AS112TimeseriesGroupsV2ParamsQueryTypeDname, AS112TimeseriesGroupsV2ParamsQueryTypeDNSKEY, AS112TimeseriesGroupsV2ParamsQueryTypeDoa, AS112TimeseriesGroupsV2ParamsQueryTypeDS, AS112TimeseriesGroupsV2ParamsQueryTypeEid, AS112TimeseriesGroupsV2ParamsQueryTypeEui48, AS112TimeseriesGroupsV2ParamsQueryTypeEui64, AS112TimeseriesGroupsV2ParamsQueryTypeGpos, AS112TimeseriesGroupsV2ParamsQueryTypeGid, AS112TimeseriesGroupsV2ParamsQueryTypeHinfo, AS112TimeseriesGroupsV2ParamsQueryTypeHip, AS112TimeseriesGroupsV2ParamsQueryTypeHTTPS, AS112TimeseriesGroupsV2ParamsQueryTypeIpseckey, AS112TimeseriesGroupsV2ParamsQueryTypeIsdn, AS112TimeseriesGroupsV2ParamsQueryTypeIxfr, AS112TimeseriesGroupsV2ParamsQueryTypeKey, AS112TimeseriesGroupsV2ParamsQueryTypeKx, AS112TimeseriesGroupsV2ParamsQueryTypeL32, AS112TimeseriesGroupsV2ParamsQueryTypeL64, AS112TimeseriesGroupsV2ParamsQueryTypeLOC, AS112TimeseriesGroupsV2ParamsQueryTypeLp, AS112TimeseriesGroupsV2ParamsQueryTypeMaila, AS112TimeseriesGroupsV2ParamsQueryTypeMailb, AS112TimeseriesGroupsV2ParamsQueryTypeMB, AS112TimeseriesGroupsV2ParamsQueryTypeMd, AS112TimeseriesGroupsV2ParamsQueryTypeMf, AS112TimeseriesGroupsV2ParamsQueryTypeMg, AS112TimeseriesGroupsV2ParamsQueryTypeMinfo, AS112TimeseriesGroupsV2ParamsQueryTypeMr, AS112TimeseriesGroupsV2ParamsQueryTypeMX, AS112TimeseriesGroupsV2ParamsQueryTypeNAPTR, AS112TimeseriesGroupsV2ParamsQueryTypeNb, AS112TimeseriesGroupsV2ParamsQueryTypeNbstat, AS112TimeseriesGroupsV2ParamsQueryTypeNid, AS112TimeseriesGroupsV2ParamsQueryTypeNimloc, AS112TimeseriesGroupsV2ParamsQueryTypeNinfo, AS112TimeseriesGroupsV2ParamsQueryTypeNS, AS112TimeseriesGroupsV2ParamsQueryTypeNsap, AS112TimeseriesGroupsV2ParamsQueryTypeNsec, AS112TimeseriesGroupsV2ParamsQueryTypeNsec3, AS112TimeseriesGroupsV2ParamsQueryTypeNsec3Param, AS112TimeseriesGroupsV2ParamsQueryTypeNull, AS112TimeseriesGroupsV2ParamsQueryTypeNxt, AS112TimeseriesGroupsV2ParamsQueryTypeOpenpgpkey, AS112TimeseriesGroupsV2ParamsQueryTypeOpt, AS112TimeseriesGroupsV2ParamsQueryTypePTR, AS112TimeseriesGroupsV2ParamsQueryTypePx, AS112TimeseriesGroupsV2ParamsQueryTypeRkey, AS112TimeseriesGroupsV2ParamsQueryTypeRp, AS112TimeseriesGroupsV2ParamsQueryTypeRrsig, AS112TimeseriesGroupsV2ParamsQueryTypeRt, AS112TimeseriesGroupsV2ParamsQueryTypeSig, AS112TimeseriesGroupsV2ParamsQueryTypeSink, AS112TimeseriesGroupsV2ParamsQueryTypeSMIMEA, AS112TimeseriesGroupsV2ParamsQueryTypeSOA, AS112TimeseriesGroupsV2ParamsQueryTypeSPF, AS112TimeseriesGroupsV2ParamsQueryTypeSRV, AS112TimeseriesGroupsV2ParamsQueryTypeSSHFP, AS112TimeseriesGroupsV2ParamsQueryTypeSVCB, AS112TimeseriesGroupsV2ParamsQueryTypeTa, AS112TimeseriesGroupsV2ParamsQueryTypeTalink, AS112TimeseriesGroupsV2ParamsQueryTypeTkey, AS112TimeseriesGroupsV2ParamsQueryTypeTLSA, AS112TimeseriesGroupsV2ParamsQueryTypeTSIG, AS112TimeseriesGroupsV2ParamsQueryTypeTXT, AS112TimeseriesGroupsV2ParamsQueryTypeUinfo, AS112TimeseriesGroupsV2ParamsQueryTypeUID, AS112TimeseriesGroupsV2ParamsQueryTypeUnspec, AS112TimeseriesGroupsV2ParamsQueryTypeURI, AS112TimeseriesGroupsV2ParamsQueryTypeWks, AS112TimeseriesGroupsV2ParamsQueryTypeX25, AS112TimeseriesGroupsV2ParamsQueryTypeZonemd:
		return true
	}
	return false
}

type AS112TimeseriesGroupsV2ParamsResponseCode string

const (
	AS112TimeseriesGroupsV2ParamsResponseCodeNoerror   AS112TimeseriesGroupsV2ParamsResponseCode = "NOERROR"
	AS112TimeseriesGroupsV2ParamsResponseCodeFormerr   AS112TimeseriesGroupsV2ParamsResponseCode = "FORMERR"
	AS112TimeseriesGroupsV2ParamsResponseCodeServfail  AS112TimeseriesGroupsV2ParamsResponseCode = "SERVFAIL"
	AS112TimeseriesGroupsV2ParamsResponseCodeNxdomain  AS112TimeseriesGroupsV2ParamsResponseCode = "NXDOMAIN"
	AS112TimeseriesGroupsV2ParamsResponseCodeNotimp    AS112TimeseriesGroupsV2ParamsResponseCode = "NOTIMP"
	AS112TimeseriesGroupsV2ParamsResponseCodeRefused   AS112TimeseriesGroupsV2ParamsResponseCode = "REFUSED"
	AS112TimeseriesGroupsV2ParamsResponseCodeYxdomain  AS112TimeseriesGroupsV2ParamsResponseCode = "YXDOMAIN"
	AS112TimeseriesGroupsV2ParamsResponseCodeYxrrset   AS112TimeseriesGroupsV2ParamsResponseCode = "YXRRSET"
	AS112TimeseriesGroupsV2ParamsResponseCodeNxrrset   AS112TimeseriesGroupsV2ParamsResponseCode = "NXRRSET"
	AS112TimeseriesGroupsV2ParamsResponseCodeNotauth   AS112TimeseriesGroupsV2ParamsResponseCode = "NOTAUTH"
	AS112TimeseriesGroupsV2ParamsResponseCodeNotzone   AS112TimeseriesGroupsV2ParamsResponseCode = "NOTZONE"
	AS112TimeseriesGroupsV2ParamsResponseCodeBadsig    AS112TimeseriesGroupsV2ParamsResponseCode = "BADSIG"
	AS112TimeseriesGroupsV2ParamsResponseCodeBadkey    AS112TimeseriesGroupsV2ParamsResponseCode = "BADKEY"
	AS112TimeseriesGroupsV2ParamsResponseCodeBadtime   AS112TimeseriesGroupsV2ParamsResponseCode = "BADTIME"
	AS112TimeseriesGroupsV2ParamsResponseCodeBadmode   AS112TimeseriesGroupsV2ParamsResponseCode = "BADMODE"
	AS112TimeseriesGroupsV2ParamsResponseCodeBadname   AS112TimeseriesGroupsV2ParamsResponseCode = "BADNAME"
	AS112TimeseriesGroupsV2ParamsResponseCodeBadalg    AS112TimeseriesGroupsV2ParamsResponseCode = "BADALG"
	AS112TimeseriesGroupsV2ParamsResponseCodeBadtrunc  AS112TimeseriesGroupsV2ParamsResponseCode = "BADTRUNC"
	AS112TimeseriesGroupsV2ParamsResponseCodeBadcookie AS112TimeseriesGroupsV2ParamsResponseCode = "BADCOOKIE"
)

func (r AS112TimeseriesGroupsV2ParamsResponseCode) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupsV2ParamsResponseCodeNoerror, AS112TimeseriesGroupsV2ParamsResponseCodeFormerr, AS112TimeseriesGroupsV2ParamsResponseCodeServfail, AS112TimeseriesGroupsV2ParamsResponseCodeNxdomain, AS112TimeseriesGroupsV2ParamsResponseCodeNotimp, AS112TimeseriesGroupsV2ParamsResponseCodeRefused, AS112TimeseriesGroupsV2ParamsResponseCodeYxdomain, AS112TimeseriesGroupsV2ParamsResponseCodeYxrrset, AS112TimeseriesGroupsV2ParamsResponseCodeNxrrset, AS112TimeseriesGroupsV2ParamsResponseCodeNotauth, AS112TimeseriesGroupsV2ParamsResponseCodeNotzone, AS112TimeseriesGroupsV2ParamsResponseCodeBadsig, AS112TimeseriesGroupsV2ParamsResponseCodeBadkey, AS112TimeseriesGroupsV2ParamsResponseCodeBadtime, AS112TimeseriesGroupsV2ParamsResponseCodeBadmode, AS112TimeseriesGroupsV2ParamsResponseCodeBadname, AS112TimeseriesGroupsV2ParamsResponseCodeBadalg, AS112TimeseriesGroupsV2ParamsResponseCodeBadtrunc, AS112TimeseriesGroupsV2ParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type AS112TimeseriesGroupsV2ResponseEnvelope struct {
	Result  AS112TimeseriesGroupsV2Response             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    as112TimeseriesGroupsV2ResponseEnvelopeJSON `json:"-"`
}

// as112TimeseriesGroupsV2ResponseEnvelopeJSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupsV2ResponseEnvelope]
type as112TimeseriesGroupsV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupsV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupsV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
