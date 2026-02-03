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

// AttackLayer7Service contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackLayer7Service] method instead.
type AttackLayer7Service struct {
	Options          []option.RequestOption
	Summary          *AttackLayer7SummaryService
	TimeseriesGroups *AttackLayer7TimeseriesGroupService
	Top              *AttackLayer7TopService
}

// NewAttackLayer7Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAttackLayer7Service(opts ...option.RequestOption) (r *AttackLayer7Service) {
	r = &AttackLayer7Service{}
	r.Options = opts
	r.Summary = NewAttackLayer7SummaryService(opts...)
	r.TimeseriesGroups = NewAttackLayer7TimeseriesGroupService(opts...)
	r.Top = NewAttackLayer7TopService(opts...)
	return
}

// Retrieves the distribution of layer 7 attacks by the specified dimension.
func (r *AttackLayer7Service) SummaryV2(ctx context.Context, dimension AttackLayer7SummaryV2ParamsDimension, query AttackLayer7SummaryV2Params, opts ...option.RequestOption) (res *AttackLayer7SummaryV2Response, err error) {
	var env AttackLayer7SummaryV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/attacks/layer7/summary/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves layer 7 attacks over time.
func (r *AttackLayer7Service) Timeseries(ctx context.Context, query AttackLayer7TimeseriesParams, opts ...option.RequestOption) (res *AttackLayer7TimeseriesResponse, err error) {
	var env AttackLayer7TimeseriesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/attacks/layer7/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 7 attacks grouped by dimension over time.
func (r *AttackLayer7Service) TimeseriesGroupsV2(ctx context.Context, dimension AttackLayer7TimeseriesGroupsV2ParamsDimension, query AttackLayer7TimeseriesGroupsV2Params, opts ...option.RequestOption) (res *AttackLayer7TimeseriesGroupsV2Response, err error) {
	var env AttackLayer7TimeseriesGroupsV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/attacks/layer7/timeseries_groups/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer7SummaryV2Response struct {
	// Metadata for the results.
	Meta     AttackLayer7SummaryV2ResponseMeta `json:"meta,required"`
	Summary0 map[string]string                 `json:"summary_0,required"`
	JSON     attackLayer7SummaryV2ResponseJSON `json:"-"`
}

// attackLayer7SummaryV2ResponseJSON contains the JSON metadata for the struct
// [AttackLayer7SummaryV2Response]
type attackLayer7SummaryV2ResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7SummaryV2ResponseMeta struct {
	ConfidenceInfo AttackLayer7SummaryV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer7SummaryV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7SummaryV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7SummaryV2ResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7SummaryV2ResponseMetaJSON   `json:"-"`
}

// attackLayer7SummaryV2ResponseMetaJSON contains the JSON metadata for the struct
// [AttackLayer7SummaryV2ResponseMeta]
type attackLayer7SummaryV2ResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7SummaryV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryV2ResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                               `json:"level,required"`
	JSON  attackLayer7SummaryV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7SummaryV2ResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AttackLayer7SummaryV2ResponseMetaConfidenceInfo]
type attackLayer7SummaryV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                               `json:"description,required"`
	EndDate     time.Time                                                            `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                          `json:"isInstantaneous,required"`
	LinkedURL       string                                                        `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                     `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotation]
type attackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll                AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots               AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos                AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet                AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AttackLayer7SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AttackLayer7SummaryV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                      `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7SummaryV2ResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7SummaryV2ResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer7SummaryV2ResponseMetaDateRange]
type attackLayer7SummaryV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7SummaryV2ResponseMetaNormalization string

const (
	AttackLayer7SummaryV2ResponseMetaNormalizationPercentage           AttackLayer7SummaryV2ResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7SummaryV2ResponseMetaNormalizationMin0Max              AttackLayer7SummaryV2ResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7SummaryV2ResponseMetaNormalizationMinMax               AttackLayer7SummaryV2ResponseMetaNormalization = "MIN_MAX"
	AttackLayer7SummaryV2ResponseMetaNormalizationRawValues            AttackLayer7SummaryV2ResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7SummaryV2ResponseMetaNormalizationPercentageChange     AttackLayer7SummaryV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7SummaryV2ResponseMetaNormalizationRollingAverage       AttackLayer7SummaryV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7SummaryV2ResponseMetaNormalizationOverlappedPercentage AttackLayer7SummaryV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7SummaryV2ResponseMetaNormalizationRatio                AttackLayer7SummaryV2ResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7SummaryV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryV2ResponseMetaNormalizationPercentage, AttackLayer7SummaryV2ResponseMetaNormalizationMin0Max, AttackLayer7SummaryV2ResponseMetaNormalizationMinMax, AttackLayer7SummaryV2ResponseMetaNormalizationRawValues, AttackLayer7SummaryV2ResponseMetaNormalizationPercentageChange, AttackLayer7SummaryV2ResponseMetaNormalizationRollingAverage, AttackLayer7SummaryV2ResponseMetaNormalizationOverlappedPercentage, AttackLayer7SummaryV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7SummaryV2ResponseMetaUnit struct {
	Name  string                                    `json:"name,required"`
	Value string                                    `json:"value,required"`
	JSON  attackLayer7SummaryV2ResponseMetaUnitJSON `json:"-"`
}

// attackLayer7SummaryV2ResponseMetaUnitJSON contains the JSON metadata for the
// struct [AttackLayer7SummaryV2ResponseMetaUnit]
type attackLayer7SummaryV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesResponse struct {
	// Metadata for the results.
	Meta   AttackLayer7TimeseriesResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer7TimeseriesResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer7TimeseriesResponseJSON   `json:"-"`
}

// attackLayer7TimeseriesResponseJSON contains the JSON metadata for the struct
// [AttackLayer7TimeseriesResponse]
type attackLayer7TimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7TimeseriesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AttackLayer7TimeseriesResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AttackLayer7TimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer7TimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7TimeseriesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7TimeseriesResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7TimeseriesResponseMetaJSON   `json:"-"`
}

// attackLayer7TimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [AttackLayer7TimeseriesResponseMeta]
type attackLayer7TimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesResponseMetaAggInterval string

const (
	AttackLayer7TimeseriesResponseMetaAggIntervalFifteenMinutes AttackLayer7TimeseriesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AttackLayer7TimeseriesResponseMetaAggIntervalOneHour        AttackLayer7TimeseriesResponseMetaAggInterval = "ONE_HOUR"
	AttackLayer7TimeseriesResponseMetaAggIntervalOneDay         AttackLayer7TimeseriesResponseMetaAggInterval = "ONE_DAY"
	AttackLayer7TimeseriesResponseMetaAggIntervalOneWeek        AttackLayer7TimeseriesResponseMetaAggInterval = "ONE_WEEK"
	AttackLayer7TimeseriesResponseMetaAggIntervalOneMonth       AttackLayer7TimeseriesResponseMetaAggInterval = "ONE_MONTH"
)

func (r AttackLayer7TimeseriesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesResponseMetaAggIntervalFifteenMinutes, AttackLayer7TimeseriesResponseMetaAggIntervalOneHour, AttackLayer7TimeseriesResponseMetaAggIntervalOneDay, AttackLayer7TimeseriesResponseMetaAggIntervalOneWeek, AttackLayer7TimeseriesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AttackLayer7TimeseriesResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                `json:"level,required"`
	JSON  attackLayer7TimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7TimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AttackLayer7TimeseriesResponseMetaConfidenceInfo]
type attackLayer7TimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                `json:"description,required"`
	EndDate     time.Time                                                             `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                      `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotation]
type attackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll                AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots               AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos                AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet                AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCT, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AttackLayer7TimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TimeseriesResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer7TimeseriesResponseMetaDateRange]
type attackLayer7TimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesResponseMetaNormalization string

const (
	AttackLayer7TimeseriesResponseMetaNormalizationPercentage           AttackLayer7TimeseriesResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesResponseMetaNormalizationMin0Max              AttackLayer7TimeseriesResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7TimeseriesResponseMetaNormalizationMinMax               AttackLayer7TimeseriesResponseMetaNormalization = "MIN_MAX"
	AttackLayer7TimeseriesResponseMetaNormalizationRawValues            AttackLayer7TimeseriesResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7TimeseriesResponseMetaNormalizationPercentageChange     AttackLayer7TimeseriesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7TimeseriesResponseMetaNormalizationRollingAverage       AttackLayer7TimeseriesResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7TimeseriesResponseMetaNormalizationOverlappedPercentage AttackLayer7TimeseriesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7TimeseriesResponseMetaNormalizationRatio                AttackLayer7TimeseriesResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7TimeseriesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesResponseMetaNormalizationPercentage, AttackLayer7TimeseriesResponseMetaNormalizationMin0Max, AttackLayer7TimeseriesResponseMetaNormalizationMinMax, AttackLayer7TimeseriesResponseMetaNormalizationRawValues, AttackLayer7TimeseriesResponseMetaNormalizationPercentageChange, AttackLayer7TimeseriesResponseMetaNormalizationRollingAverage, AttackLayer7TimeseriesResponseMetaNormalizationOverlappedPercentage, AttackLayer7TimeseriesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7TimeseriesResponseMetaUnit struct {
	Name  string                                     `json:"name,required"`
	Value string                                     `json:"value,required"`
	JSON  attackLayer7TimeseriesResponseMetaUnitJSON `json:"-"`
}

// attackLayer7TimeseriesResponseMetaUnitJSON contains the JSON metadata for the
// struct [AttackLayer7TimeseriesResponseMetaUnit]
type attackLayer7TimeseriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesResponseSerie0 struct {
	Timestamps []time.Time                              `json:"timestamps,required" format:"date-time"`
	Values     []string                                 `json:"values,required"`
	JSON       attackLayer7TimeseriesResponseSerie0JSON `json:"-"`
}

// attackLayer7TimeseriesResponseSerie0JSON contains the JSON metadata for the
// struct [AttackLayer7TimeseriesResponseSerie0]
type attackLayer7TimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupsV2Response struct {
	// Metadata for the results.
	Meta   AttackLayer7TimeseriesGroupsV2ResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer7TimeseriesGroupsV2ResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer7TimeseriesGroupsV2ResponseJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupsV2ResponseJSON contains the JSON metadata for the
// struct [AttackLayer7TimeseriesGroupsV2Response]
type attackLayer7TimeseriesGroupsV2ResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupsV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupsV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7TimeseriesGroupsV2ResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AttackLayer7TimeseriesGroupsV2ResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer7TimeseriesGroupsV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7TimeseriesGroupsV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7TimeseriesGroupsV2ResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7TimeseriesGroupsV2ResponseMetaJSON   `json:"-"`
}

// attackLayer7TimeseriesGroupsV2ResponseMetaJSON contains the JSON metadata for
// the struct [AttackLayer7TimeseriesGroupsV2ResponseMeta]
type attackLayer7TimeseriesGroupsV2ResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupsV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupsV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupsV2ResponseMetaAggInterval string

const (
	AttackLayer7TimeseriesGroupsV2ResponseMetaAggIntervalFifteenMinutes AttackLayer7TimeseriesGroupsV2ResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AttackLayer7TimeseriesGroupsV2ResponseMetaAggIntervalOneHour        AttackLayer7TimeseriesGroupsV2ResponseMetaAggInterval = "ONE_HOUR"
	AttackLayer7TimeseriesGroupsV2ResponseMetaAggIntervalOneDay         AttackLayer7TimeseriesGroupsV2ResponseMetaAggInterval = "ONE_DAY"
	AttackLayer7TimeseriesGroupsV2ResponseMetaAggIntervalOneWeek        AttackLayer7TimeseriesGroupsV2ResponseMetaAggInterval = "ONE_WEEK"
	AttackLayer7TimeseriesGroupsV2ResponseMetaAggIntervalOneMonth       AttackLayer7TimeseriesGroupsV2ResponseMetaAggInterval = "ONE_MONTH"
)

func (r AttackLayer7TimeseriesGroupsV2ResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupsV2ResponseMetaAggIntervalFifteenMinutes, AttackLayer7TimeseriesGroupsV2ResponseMetaAggIntervalOneHour, AttackLayer7TimeseriesGroupsV2ResponseMetaAggIntervalOneDay, AttackLayer7TimeseriesGroupsV2ResponseMetaAggIntervalOneWeek, AttackLayer7TimeseriesGroupsV2ResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                        `json:"level,required"`
	JSON  attackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfo]
type attackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                        `json:"description,required"`
	EndDate     time.Time                                                                     `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                              `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation]
type attackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll                AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots               AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos                AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet                AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AttackLayer7TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupsV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TimeseriesGroupsV2ResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupsV2ResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AttackLayer7TimeseriesGroupsV2ResponseMetaDateRange]
type attackLayer7TimeseriesGroupsV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupsV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupsV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupsV2ResponseMetaNormalization string

const (
	AttackLayer7TimeseriesGroupsV2ResponseMetaNormalizationPercentage           AttackLayer7TimeseriesGroupsV2ResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupsV2ResponseMetaNormalizationMin0Max              AttackLayer7TimeseriesGroupsV2ResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7TimeseriesGroupsV2ResponseMetaNormalizationMinMax               AttackLayer7TimeseriesGroupsV2ResponseMetaNormalization = "MIN_MAX"
	AttackLayer7TimeseriesGroupsV2ResponseMetaNormalizationRawValues            AttackLayer7TimeseriesGroupsV2ResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7TimeseriesGroupsV2ResponseMetaNormalizationPercentageChange     AttackLayer7TimeseriesGroupsV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7TimeseriesGroupsV2ResponseMetaNormalizationRollingAverage       AttackLayer7TimeseriesGroupsV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7TimeseriesGroupsV2ResponseMetaNormalizationOverlappedPercentage AttackLayer7TimeseriesGroupsV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7TimeseriesGroupsV2ResponseMetaNormalizationRatio                AttackLayer7TimeseriesGroupsV2ResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7TimeseriesGroupsV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupsV2ResponseMetaNormalizationPercentage, AttackLayer7TimeseriesGroupsV2ResponseMetaNormalizationMin0Max, AttackLayer7TimeseriesGroupsV2ResponseMetaNormalizationMinMax, AttackLayer7TimeseriesGroupsV2ResponseMetaNormalizationRawValues, AttackLayer7TimeseriesGroupsV2ResponseMetaNormalizationPercentageChange, AttackLayer7TimeseriesGroupsV2ResponseMetaNormalizationRollingAverage, AttackLayer7TimeseriesGroupsV2ResponseMetaNormalizationOverlappedPercentage, AttackLayer7TimeseriesGroupsV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupsV2ResponseMetaUnit struct {
	Name  string                                             `json:"name,required"`
	Value string                                             `json:"value,required"`
	JSON  attackLayer7TimeseriesGroupsV2ResponseMetaUnitJSON `json:"-"`
}

// attackLayer7TimeseriesGroupsV2ResponseMetaUnitJSON contains the JSON metadata
// for the struct [AttackLayer7TimeseriesGroupsV2ResponseMetaUnit]
type attackLayer7TimeseriesGroupsV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupsV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupsV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupsV2ResponseSerie0 struct {
	Timestamps  []time.Time                                      `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                              `json:"-,extras"`
	JSON        attackLayer7TimeseriesGroupsV2ResponseSerie0JSON `json:"-"`
}

// attackLayer7TimeseriesGroupsV2ResponseSerie0JSON contains the JSON metadata for
// the struct [AttackLayer7TimeseriesGroupsV2ResponseSerie0]
type attackLayer7TimeseriesGroupsV2ResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupsV2ResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupsV2ResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7SummaryV2Params struct {
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
	Format param.Field[AttackLayer7SummaryV2ParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]AttackLayer7SummaryV2ParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]AttackLayer7SummaryV2ParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer7SummaryV2ParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7SummaryV2ParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7SummaryV2Params]'s query parameters as
// `url.Values`.
func (r AttackLayer7SummaryV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type AttackLayer7SummaryV2ParamsDimension string

const (
	AttackLayer7SummaryV2ParamsDimensionHTTPMethod        AttackLayer7SummaryV2ParamsDimension = "HTTP_METHOD"
	AttackLayer7SummaryV2ParamsDimensionHTTPVersion       AttackLayer7SummaryV2ParamsDimension = "HTTP_VERSION"
	AttackLayer7SummaryV2ParamsDimensionIPVersion         AttackLayer7SummaryV2ParamsDimension = "IP_VERSION"
	AttackLayer7SummaryV2ParamsDimensionManagedRules      AttackLayer7SummaryV2ParamsDimension = "MANAGED_RULES"
	AttackLayer7SummaryV2ParamsDimensionMitigationProduct AttackLayer7SummaryV2ParamsDimension = "MITIGATION_PRODUCT"
	AttackLayer7SummaryV2ParamsDimensionVertical          AttackLayer7SummaryV2ParamsDimension = "VERTICAL"
	AttackLayer7SummaryV2ParamsDimensionIndustry          AttackLayer7SummaryV2ParamsDimension = "INDUSTRY"
)

func (r AttackLayer7SummaryV2ParamsDimension) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryV2ParamsDimensionHTTPMethod, AttackLayer7SummaryV2ParamsDimensionHTTPVersion, AttackLayer7SummaryV2ParamsDimensionIPVersion, AttackLayer7SummaryV2ParamsDimensionManagedRules, AttackLayer7SummaryV2ParamsDimensionMitigationProduct, AttackLayer7SummaryV2ParamsDimensionVertical, AttackLayer7SummaryV2ParamsDimensionIndustry:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer7SummaryV2ParamsFormat string

const (
	AttackLayer7SummaryV2ParamsFormatJson AttackLayer7SummaryV2ParamsFormat = "JSON"
	AttackLayer7SummaryV2ParamsFormatCsv  AttackLayer7SummaryV2ParamsFormat = "CSV"
)

func (r AttackLayer7SummaryV2ParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryV2ParamsFormatJson, AttackLayer7SummaryV2ParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7SummaryV2ParamsHTTPMethod string

const (
	AttackLayer7SummaryV2ParamsHTTPMethodGet             AttackLayer7SummaryV2ParamsHTTPMethod = "GET"
	AttackLayer7SummaryV2ParamsHTTPMethodPost            AttackLayer7SummaryV2ParamsHTTPMethod = "POST"
	AttackLayer7SummaryV2ParamsHTTPMethodDelete          AttackLayer7SummaryV2ParamsHTTPMethod = "DELETE"
	AttackLayer7SummaryV2ParamsHTTPMethodPut             AttackLayer7SummaryV2ParamsHTTPMethod = "PUT"
	AttackLayer7SummaryV2ParamsHTTPMethodHead            AttackLayer7SummaryV2ParamsHTTPMethod = "HEAD"
	AttackLayer7SummaryV2ParamsHTTPMethodPurge           AttackLayer7SummaryV2ParamsHTTPMethod = "PURGE"
	AttackLayer7SummaryV2ParamsHTTPMethodOptions         AttackLayer7SummaryV2ParamsHTTPMethod = "OPTIONS"
	AttackLayer7SummaryV2ParamsHTTPMethodPropfind        AttackLayer7SummaryV2ParamsHTTPMethod = "PROPFIND"
	AttackLayer7SummaryV2ParamsHTTPMethodMkcol           AttackLayer7SummaryV2ParamsHTTPMethod = "MKCOL"
	AttackLayer7SummaryV2ParamsHTTPMethodPatch           AttackLayer7SummaryV2ParamsHTTPMethod = "PATCH"
	AttackLayer7SummaryV2ParamsHTTPMethodACL             AttackLayer7SummaryV2ParamsHTTPMethod = "ACL"
	AttackLayer7SummaryV2ParamsHTTPMethodBcopy           AttackLayer7SummaryV2ParamsHTTPMethod = "BCOPY"
	AttackLayer7SummaryV2ParamsHTTPMethodBdelete         AttackLayer7SummaryV2ParamsHTTPMethod = "BDELETE"
	AttackLayer7SummaryV2ParamsHTTPMethodBmove           AttackLayer7SummaryV2ParamsHTTPMethod = "BMOVE"
	AttackLayer7SummaryV2ParamsHTTPMethodBpropfind       AttackLayer7SummaryV2ParamsHTTPMethod = "BPROPFIND"
	AttackLayer7SummaryV2ParamsHTTPMethodBproppatch      AttackLayer7SummaryV2ParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7SummaryV2ParamsHTTPMethodCheckin         AttackLayer7SummaryV2ParamsHTTPMethod = "CHECKIN"
	AttackLayer7SummaryV2ParamsHTTPMethodCheckout        AttackLayer7SummaryV2ParamsHTTPMethod = "CHECKOUT"
	AttackLayer7SummaryV2ParamsHTTPMethodConnect         AttackLayer7SummaryV2ParamsHTTPMethod = "CONNECT"
	AttackLayer7SummaryV2ParamsHTTPMethodCopy            AttackLayer7SummaryV2ParamsHTTPMethod = "COPY"
	AttackLayer7SummaryV2ParamsHTTPMethodLabel           AttackLayer7SummaryV2ParamsHTTPMethod = "LABEL"
	AttackLayer7SummaryV2ParamsHTTPMethodLock            AttackLayer7SummaryV2ParamsHTTPMethod = "LOCK"
	AttackLayer7SummaryV2ParamsHTTPMethodMerge           AttackLayer7SummaryV2ParamsHTTPMethod = "MERGE"
	AttackLayer7SummaryV2ParamsHTTPMethodMkactivity      AttackLayer7SummaryV2ParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7SummaryV2ParamsHTTPMethodMkworkspace     AttackLayer7SummaryV2ParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7SummaryV2ParamsHTTPMethodMove            AttackLayer7SummaryV2ParamsHTTPMethod = "MOVE"
	AttackLayer7SummaryV2ParamsHTTPMethodNotify          AttackLayer7SummaryV2ParamsHTTPMethod = "NOTIFY"
	AttackLayer7SummaryV2ParamsHTTPMethodOrderpatch      AttackLayer7SummaryV2ParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7SummaryV2ParamsHTTPMethodPoll            AttackLayer7SummaryV2ParamsHTTPMethod = "POLL"
	AttackLayer7SummaryV2ParamsHTTPMethodProppatch       AttackLayer7SummaryV2ParamsHTTPMethod = "PROPPATCH"
	AttackLayer7SummaryV2ParamsHTTPMethodReport          AttackLayer7SummaryV2ParamsHTTPMethod = "REPORT"
	AttackLayer7SummaryV2ParamsHTTPMethodSearch          AttackLayer7SummaryV2ParamsHTTPMethod = "SEARCH"
	AttackLayer7SummaryV2ParamsHTTPMethodSubscribe       AttackLayer7SummaryV2ParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7SummaryV2ParamsHTTPMethodTrace           AttackLayer7SummaryV2ParamsHTTPMethod = "TRACE"
	AttackLayer7SummaryV2ParamsHTTPMethodUncheckout      AttackLayer7SummaryV2ParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7SummaryV2ParamsHTTPMethodUnlock          AttackLayer7SummaryV2ParamsHTTPMethod = "UNLOCK"
	AttackLayer7SummaryV2ParamsHTTPMethodUnsubscribe     AttackLayer7SummaryV2ParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7SummaryV2ParamsHTTPMethodUpdate          AttackLayer7SummaryV2ParamsHTTPMethod = "UPDATE"
	AttackLayer7SummaryV2ParamsHTTPMethodVersioncontrol  AttackLayer7SummaryV2ParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7SummaryV2ParamsHTTPMethodBaselinecontrol AttackLayer7SummaryV2ParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7SummaryV2ParamsHTTPMethodXmsenumatts     AttackLayer7SummaryV2ParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7SummaryV2ParamsHTTPMethodRpcOutData      AttackLayer7SummaryV2ParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7SummaryV2ParamsHTTPMethodRpcInData       AttackLayer7SummaryV2ParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7SummaryV2ParamsHTTPMethodJson            AttackLayer7SummaryV2ParamsHTTPMethod = "JSON"
	AttackLayer7SummaryV2ParamsHTTPMethodCook            AttackLayer7SummaryV2ParamsHTTPMethod = "COOK"
	AttackLayer7SummaryV2ParamsHTTPMethodTrack           AttackLayer7SummaryV2ParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7SummaryV2ParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryV2ParamsHTTPMethodGet, AttackLayer7SummaryV2ParamsHTTPMethodPost, AttackLayer7SummaryV2ParamsHTTPMethodDelete, AttackLayer7SummaryV2ParamsHTTPMethodPut, AttackLayer7SummaryV2ParamsHTTPMethodHead, AttackLayer7SummaryV2ParamsHTTPMethodPurge, AttackLayer7SummaryV2ParamsHTTPMethodOptions, AttackLayer7SummaryV2ParamsHTTPMethodPropfind, AttackLayer7SummaryV2ParamsHTTPMethodMkcol, AttackLayer7SummaryV2ParamsHTTPMethodPatch, AttackLayer7SummaryV2ParamsHTTPMethodACL, AttackLayer7SummaryV2ParamsHTTPMethodBcopy, AttackLayer7SummaryV2ParamsHTTPMethodBdelete, AttackLayer7SummaryV2ParamsHTTPMethodBmove, AttackLayer7SummaryV2ParamsHTTPMethodBpropfind, AttackLayer7SummaryV2ParamsHTTPMethodBproppatch, AttackLayer7SummaryV2ParamsHTTPMethodCheckin, AttackLayer7SummaryV2ParamsHTTPMethodCheckout, AttackLayer7SummaryV2ParamsHTTPMethodConnect, AttackLayer7SummaryV2ParamsHTTPMethodCopy, AttackLayer7SummaryV2ParamsHTTPMethodLabel, AttackLayer7SummaryV2ParamsHTTPMethodLock, AttackLayer7SummaryV2ParamsHTTPMethodMerge, AttackLayer7SummaryV2ParamsHTTPMethodMkactivity, AttackLayer7SummaryV2ParamsHTTPMethodMkworkspace, AttackLayer7SummaryV2ParamsHTTPMethodMove, AttackLayer7SummaryV2ParamsHTTPMethodNotify, AttackLayer7SummaryV2ParamsHTTPMethodOrderpatch, AttackLayer7SummaryV2ParamsHTTPMethodPoll, AttackLayer7SummaryV2ParamsHTTPMethodProppatch, AttackLayer7SummaryV2ParamsHTTPMethodReport, AttackLayer7SummaryV2ParamsHTTPMethodSearch, AttackLayer7SummaryV2ParamsHTTPMethodSubscribe, AttackLayer7SummaryV2ParamsHTTPMethodTrace, AttackLayer7SummaryV2ParamsHTTPMethodUncheckout, AttackLayer7SummaryV2ParamsHTTPMethodUnlock, AttackLayer7SummaryV2ParamsHTTPMethodUnsubscribe, AttackLayer7SummaryV2ParamsHTTPMethodUpdate, AttackLayer7SummaryV2ParamsHTTPMethodVersioncontrol, AttackLayer7SummaryV2ParamsHTTPMethodBaselinecontrol, AttackLayer7SummaryV2ParamsHTTPMethodXmsenumatts, AttackLayer7SummaryV2ParamsHTTPMethodRpcOutData, AttackLayer7SummaryV2ParamsHTTPMethodRpcInData, AttackLayer7SummaryV2ParamsHTTPMethodJson, AttackLayer7SummaryV2ParamsHTTPMethodCook, AttackLayer7SummaryV2ParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7SummaryV2ParamsHTTPVersion string

const (
	AttackLayer7SummaryV2ParamsHTTPVersionHttPv1 AttackLayer7SummaryV2ParamsHTTPVersion = "HTTPv1"
	AttackLayer7SummaryV2ParamsHTTPVersionHttPv2 AttackLayer7SummaryV2ParamsHTTPVersion = "HTTPv2"
	AttackLayer7SummaryV2ParamsHTTPVersionHttPv3 AttackLayer7SummaryV2ParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7SummaryV2ParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryV2ParamsHTTPVersionHttPv1, AttackLayer7SummaryV2ParamsHTTPVersionHttPv2, AttackLayer7SummaryV2ParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7SummaryV2ParamsIPVersion string

const (
	AttackLayer7SummaryV2ParamsIPVersionIPv4 AttackLayer7SummaryV2ParamsIPVersion = "IPv4"
	AttackLayer7SummaryV2ParamsIPVersionIPv6 AttackLayer7SummaryV2ParamsIPVersion = "IPv6"
)

func (r AttackLayer7SummaryV2ParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryV2ParamsIPVersionIPv4, AttackLayer7SummaryV2ParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7SummaryV2ParamsMitigationProduct string

const (
	AttackLayer7SummaryV2ParamsMitigationProductDDoS               AttackLayer7SummaryV2ParamsMitigationProduct = "DDOS"
	AttackLayer7SummaryV2ParamsMitigationProductWAF                AttackLayer7SummaryV2ParamsMitigationProduct = "WAF"
	AttackLayer7SummaryV2ParamsMitigationProductBotManagement      AttackLayer7SummaryV2ParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7SummaryV2ParamsMitigationProductAccessRules        AttackLayer7SummaryV2ParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7SummaryV2ParamsMitigationProductIPReputation       AttackLayer7SummaryV2ParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7SummaryV2ParamsMitigationProductAPIShield          AttackLayer7SummaryV2ParamsMitigationProduct = "API_SHIELD"
	AttackLayer7SummaryV2ParamsMitigationProductDataLossPrevention AttackLayer7SummaryV2ParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7SummaryV2ParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7SummaryV2ParamsMitigationProductDDoS, AttackLayer7SummaryV2ParamsMitigationProductWAF, AttackLayer7SummaryV2ParamsMitigationProductBotManagement, AttackLayer7SummaryV2ParamsMitigationProductAccessRules, AttackLayer7SummaryV2ParamsMitigationProductIPReputation, AttackLayer7SummaryV2ParamsMitigationProductAPIShield, AttackLayer7SummaryV2ParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type AttackLayer7SummaryV2ResponseEnvelope struct {
	Result  AttackLayer7SummaryV2Response             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    attackLayer7SummaryV2ResponseEnvelopeJSON `json:"-"`
}

// attackLayer7SummaryV2ResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer7SummaryV2ResponseEnvelope]
type attackLayer7SummaryV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7SummaryV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7SummaryV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer7TimeseriesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AttackLayer7TimeseriesParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]AttackLayer7TimeseriesParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]AttackLayer7TimeseriesParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer7TimeseriesParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7TimeseriesParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer7TimeseriesParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer7TimeseriesParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7TimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesParamsAggInterval string

const (
	AttackLayer7TimeseriesParamsAggInterval15m AttackLayer7TimeseriesParamsAggInterval = "15m"
	AttackLayer7TimeseriesParamsAggInterval1h  AttackLayer7TimeseriesParamsAggInterval = "1h"
	AttackLayer7TimeseriesParamsAggInterval1d  AttackLayer7TimeseriesParamsAggInterval = "1d"
	AttackLayer7TimeseriesParamsAggInterval1w  AttackLayer7TimeseriesParamsAggInterval = "1w"
)

func (r AttackLayer7TimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesParamsAggInterval15m, AttackLayer7TimeseriesParamsAggInterval1h, AttackLayer7TimeseriesParamsAggInterval1d, AttackLayer7TimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer7TimeseriesParamsFormat string

const (
	AttackLayer7TimeseriesParamsFormatJson AttackLayer7TimeseriesParamsFormat = "JSON"
	AttackLayer7TimeseriesParamsFormatCsv  AttackLayer7TimeseriesParamsFormat = "CSV"
)

func (r AttackLayer7TimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesParamsFormatJson, AttackLayer7TimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TimeseriesParamsHTTPMethod string

const (
	AttackLayer7TimeseriesParamsHTTPMethodGet             AttackLayer7TimeseriesParamsHTTPMethod = "GET"
	AttackLayer7TimeseriesParamsHTTPMethodPost            AttackLayer7TimeseriesParamsHTTPMethod = "POST"
	AttackLayer7TimeseriesParamsHTTPMethodDelete          AttackLayer7TimeseriesParamsHTTPMethod = "DELETE"
	AttackLayer7TimeseriesParamsHTTPMethodPut             AttackLayer7TimeseriesParamsHTTPMethod = "PUT"
	AttackLayer7TimeseriesParamsHTTPMethodHead            AttackLayer7TimeseriesParamsHTTPMethod = "HEAD"
	AttackLayer7TimeseriesParamsHTTPMethodPurge           AttackLayer7TimeseriesParamsHTTPMethod = "PURGE"
	AttackLayer7TimeseriesParamsHTTPMethodOptions         AttackLayer7TimeseriesParamsHTTPMethod = "OPTIONS"
	AttackLayer7TimeseriesParamsHTTPMethodPropfind        AttackLayer7TimeseriesParamsHTTPMethod = "PROPFIND"
	AttackLayer7TimeseriesParamsHTTPMethodMkcol           AttackLayer7TimeseriesParamsHTTPMethod = "MKCOL"
	AttackLayer7TimeseriesParamsHTTPMethodPatch           AttackLayer7TimeseriesParamsHTTPMethod = "PATCH"
	AttackLayer7TimeseriesParamsHTTPMethodACL             AttackLayer7TimeseriesParamsHTTPMethod = "ACL"
	AttackLayer7TimeseriesParamsHTTPMethodBcopy           AttackLayer7TimeseriesParamsHTTPMethod = "BCOPY"
	AttackLayer7TimeseriesParamsHTTPMethodBdelete         AttackLayer7TimeseriesParamsHTTPMethod = "BDELETE"
	AttackLayer7TimeseriesParamsHTTPMethodBmove           AttackLayer7TimeseriesParamsHTTPMethod = "BMOVE"
	AttackLayer7TimeseriesParamsHTTPMethodBpropfind       AttackLayer7TimeseriesParamsHTTPMethod = "BPROPFIND"
	AttackLayer7TimeseriesParamsHTTPMethodBproppatch      AttackLayer7TimeseriesParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7TimeseriesParamsHTTPMethodCheckin         AttackLayer7TimeseriesParamsHTTPMethod = "CHECKIN"
	AttackLayer7TimeseriesParamsHTTPMethodCheckout        AttackLayer7TimeseriesParamsHTTPMethod = "CHECKOUT"
	AttackLayer7TimeseriesParamsHTTPMethodConnect         AttackLayer7TimeseriesParamsHTTPMethod = "CONNECT"
	AttackLayer7TimeseriesParamsHTTPMethodCopy            AttackLayer7TimeseriesParamsHTTPMethod = "COPY"
	AttackLayer7TimeseriesParamsHTTPMethodLabel           AttackLayer7TimeseriesParamsHTTPMethod = "LABEL"
	AttackLayer7TimeseriesParamsHTTPMethodLock            AttackLayer7TimeseriesParamsHTTPMethod = "LOCK"
	AttackLayer7TimeseriesParamsHTTPMethodMerge           AttackLayer7TimeseriesParamsHTTPMethod = "MERGE"
	AttackLayer7TimeseriesParamsHTTPMethodMkactivity      AttackLayer7TimeseriesParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7TimeseriesParamsHTTPMethodMkworkspace     AttackLayer7TimeseriesParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7TimeseriesParamsHTTPMethodMove            AttackLayer7TimeseriesParamsHTTPMethod = "MOVE"
	AttackLayer7TimeseriesParamsHTTPMethodNotify          AttackLayer7TimeseriesParamsHTTPMethod = "NOTIFY"
	AttackLayer7TimeseriesParamsHTTPMethodOrderpatch      AttackLayer7TimeseriesParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7TimeseriesParamsHTTPMethodPoll            AttackLayer7TimeseriesParamsHTTPMethod = "POLL"
	AttackLayer7TimeseriesParamsHTTPMethodProppatch       AttackLayer7TimeseriesParamsHTTPMethod = "PROPPATCH"
	AttackLayer7TimeseriesParamsHTTPMethodReport          AttackLayer7TimeseriesParamsHTTPMethod = "REPORT"
	AttackLayer7TimeseriesParamsHTTPMethodSearch          AttackLayer7TimeseriesParamsHTTPMethod = "SEARCH"
	AttackLayer7TimeseriesParamsHTTPMethodSubscribe       AttackLayer7TimeseriesParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7TimeseriesParamsHTTPMethodTrace           AttackLayer7TimeseriesParamsHTTPMethod = "TRACE"
	AttackLayer7TimeseriesParamsHTTPMethodUncheckout      AttackLayer7TimeseriesParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7TimeseriesParamsHTTPMethodUnlock          AttackLayer7TimeseriesParamsHTTPMethod = "UNLOCK"
	AttackLayer7TimeseriesParamsHTTPMethodUnsubscribe     AttackLayer7TimeseriesParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7TimeseriesParamsHTTPMethodUpdate          AttackLayer7TimeseriesParamsHTTPMethod = "UPDATE"
	AttackLayer7TimeseriesParamsHTTPMethodVersioncontrol  AttackLayer7TimeseriesParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7TimeseriesParamsHTTPMethodBaselinecontrol AttackLayer7TimeseriesParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7TimeseriesParamsHTTPMethodXmsenumatts     AttackLayer7TimeseriesParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7TimeseriesParamsHTTPMethodRpcOutData      AttackLayer7TimeseriesParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7TimeseriesParamsHTTPMethodRpcInData       AttackLayer7TimeseriesParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7TimeseriesParamsHTTPMethodJson            AttackLayer7TimeseriesParamsHTTPMethod = "JSON"
	AttackLayer7TimeseriesParamsHTTPMethodCook            AttackLayer7TimeseriesParamsHTTPMethod = "COOK"
	AttackLayer7TimeseriesParamsHTTPMethodTrack           AttackLayer7TimeseriesParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7TimeseriesParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesParamsHTTPMethodGet, AttackLayer7TimeseriesParamsHTTPMethodPost, AttackLayer7TimeseriesParamsHTTPMethodDelete, AttackLayer7TimeseriesParamsHTTPMethodPut, AttackLayer7TimeseriesParamsHTTPMethodHead, AttackLayer7TimeseriesParamsHTTPMethodPurge, AttackLayer7TimeseriesParamsHTTPMethodOptions, AttackLayer7TimeseriesParamsHTTPMethodPropfind, AttackLayer7TimeseriesParamsHTTPMethodMkcol, AttackLayer7TimeseriesParamsHTTPMethodPatch, AttackLayer7TimeseriesParamsHTTPMethodACL, AttackLayer7TimeseriesParamsHTTPMethodBcopy, AttackLayer7TimeseriesParamsHTTPMethodBdelete, AttackLayer7TimeseriesParamsHTTPMethodBmove, AttackLayer7TimeseriesParamsHTTPMethodBpropfind, AttackLayer7TimeseriesParamsHTTPMethodBproppatch, AttackLayer7TimeseriesParamsHTTPMethodCheckin, AttackLayer7TimeseriesParamsHTTPMethodCheckout, AttackLayer7TimeseriesParamsHTTPMethodConnect, AttackLayer7TimeseriesParamsHTTPMethodCopy, AttackLayer7TimeseriesParamsHTTPMethodLabel, AttackLayer7TimeseriesParamsHTTPMethodLock, AttackLayer7TimeseriesParamsHTTPMethodMerge, AttackLayer7TimeseriesParamsHTTPMethodMkactivity, AttackLayer7TimeseriesParamsHTTPMethodMkworkspace, AttackLayer7TimeseriesParamsHTTPMethodMove, AttackLayer7TimeseriesParamsHTTPMethodNotify, AttackLayer7TimeseriesParamsHTTPMethodOrderpatch, AttackLayer7TimeseriesParamsHTTPMethodPoll, AttackLayer7TimeseriesParamsHTTPMethodProppatch, AttackLayer7TimeseriesParamsHTTPMethodReport, AttackLayer7TimeseriesParamsHTTPMethodSearch, AttackLayer7TimeseriesParamsHTTPMethodSubscribe, AttackLayer7TimeseriesParamsHTTPMethodTrace, AttackLayer7TimeseriesParamsHTTPMethodUncheckout, AttackLayer7TimeseriesParamsHTTPMethodUnlock, AttackLayer7TimeseriesParamsHTTPMethodUnsubscribe, AttackLayer7TimeseriesParamsHTTPMethodUpdate, AttackLayer7TimeseriesParamsHTTPMethodVersioncontrol, AttackLayer7TimeseriesParamsHTTPMethodBaselinecontrol, AttackLayer7TimeseriesParamsHTTPMethodXmsenumatts, AttackLayer7TimeseriesParamsHTTPMethodRpcOutData, AttackLayer7TimeseriesParamsHTTPMethodRpcInData, AttackLayer7TimeseriesParamsHTTPMethodJson, AttackLayer7TimeseriesParamsHTTPMethodCook, AttackLayer7TimeseriesParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7TimeseriesParamsHTTPVersion string

const (
	AttackLayer7TimeseriesParamsHTTPVersionHttPv1 AttackLayer7TimeseriesParamsHTTPVersion = "HTTPv1"
	AttackLayer7TimeseriesParamsHTTPVersionHttPv2 AttackLayer7TimeseriesParamsHTTPVersion = "HTTPv2"
	AttackLayer7TimeseriesParamsHTTPVersionHttPv3 AttackLayer7TimeseriesParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7TimeseriesParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesParamsHTTPVersionHttPv1, AttackLayer7TimeseriesParamsHTTPVersionHttPv2, AttackLayer7TimeseriesParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7TimeseriesParamsIPVersion string

const (
	AttackLayer7TimeseriesParamsIPVersionIPv4 AttackLayer7TimeseriesParamsIPVersion = "IPv4"
	AttackLayer7TimeseriesParamsIPVersionIPv6 AttackLayer7TimeseriesParamsIPVersion = "IPv6"
)

func (r AttackLayer7TimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesParamsIPVersionIPv4, AttackLayer7TimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7TimeseriesParamsMitigationProduct string

const (
	AttackLayer7TimeseriesParamsMitigationProductDDoS               AttackLayer7TimeseriesParamsMitigationProduct = "DDOS"
	AttackLayer7TimeseriesParamsMitigationProductWAF                AttackLayer7TimeseriesParamsMitigationProduct = "WAF"
	AttackLayer7TimeseriesParamsMitigationProductBotManagement      AttackLayer7TimeseriesParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TimeseriesParamsMitigationProductAccessRules        AttackLayer7TimeseriesParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TimeseriesParamsMitigationProductIPReputation       AttackLayer7TimeseriesParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TimeseriesParamsMitigationProductAPIShield          AttackLayer7TimeseriesParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TimeseriesParamsMitigationProductDataLossPrevention AttackLayer7TimeseriesParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TimeseriesParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesParamsMitigationProductDDoS, AttackLayer7TimeseriesParamsMitigationProductWAF, AttackLayer7TimeseriesParamsMitigationProductBotManagement, AttackLayer7TimeseriesParamsMitigationProductAccessRules, AttackLayer7TimeseriesParamsMitigationProductIPReputation, AttackLayer7TimeseriesParamsMitigationProductAPIShield, AttackLayer7TimeseriesParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesParamsNormalization string

const (
	AttackLayer7TimeseriesParamsNormalizationPercentageChange AttackLayer7TimeseriesParamsNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7TimeseriesParamsNormalizationMin0Max          AttackLayer7TimeseriesParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer7TimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesParamsNormalizationPercentageChange, AttackLayer7TimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer7TimeseriesResponseEnvelope struct {
	Result  AttackLayer7TimeseriesResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    attackLayer7TimeseriesResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TimeseriesResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer7TimeseriesResponseEnvelope]
type attackLayer7TimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesGroupsV2Params struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer7TimeseriesGroupsV2ParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AttackLayer7TimeseriesGroupsV2ParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]AttackLayer7TimeseriesGroupsV2ParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer7TimeseriesGroupsV2ParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7TimeseriesGroupsV2ParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer7TimeseriesGroupsV2ParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer7TimeseriesGroupsV2Params]'s query parameters as
// `url.Values`.
func (r AttackLayer7TimeseriesGroupsV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type AttackLayer7TimeseriesGroupsV2ParamsDimension string

const (
	AttackLayer7TimeseriesGroupsV2ParamsDimensionHTTPMethod        AttackLayer7TimeseriesGroupsV2ParamsDimension = "HTTP_METHOD"
	AttackLayer7TimeseriesGroupsV2ParamsDimensionHTTPVersion       AttackLayer7TimeseriesGroupsV2ParamsDimension = "HTTP_VERSION"
	AttackLayer7TimeseriesGroupsV2ParamsDimensionIPVersion         AttackLayer7TimeseriesGroupsV2ParamsDimension = "IP_VERSION"
	AttackLayer7TimeseriesGroupsV2ParamsDimensionManagedRules      AttackLayer7TimeseriesGroupsV2ParamsDimension = "MANAGED_RULES"
	AttackLayer7TimeseriesGroupsV2ParamsDimensionMitigationProduct AttackLayer7TimeseriesGroupsV2ParamsDimension = "MITIGATION_PRODUCT"
	AttackLayer7TimeseriesGroupsV2ParamsDimensionVertical          AttackLayer7TimeseriesGroupsV2ParamsDimension = "VERTICAL"
	AttackLayer7TimeseriesGroupsV2ParamsDimensionIndustry          AttackLayer7TimeseriesGroupsV2ParamsDimension = "INDUSTRY"
)

func (r AttackLayer7TimeseriesGroupsV2ParamsDimension) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupsV2ParamsDimensionHTTPMethod, AttackLayer7TimeseriesGroupsV2ParamsDimensionHTTPVersion, AttackLayer7TimeseriesGroupsV2ParamsDimensionIPVersion, AttackLayer7TimeseriesGroupsV2ParamsDimensionManagedRules, AttackLayer7TimeseriesGroupsV2ParamsDimensionMitigationProduct, AttackLayer7TimeseriesGroupsV2ParamsDimensionVertical, AttackLayer7TimeseriesGroupsV2ParamsDimensionIndustry:
		return true
	}
	return false
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesGroupsV2ParamsAggInterval string

const (
	AttackLayer7TimeseriesGroupsV2ParamsAggInterval15m AttackLayer7TimeseriesGroupsV2ParamsAggInterval = "15m"
	AttackLayer7TimeseriesGroupsV2ParamsAggInterval1h  AttackLayer7TimeseriesGroupsV2ParamsAggInterval = "1h"
	AttackLayer7TimeseriesGroupsV2ParamsAggInterval1d  AttackLayer7TimeseriesGroupsV2ParamsAggInterval = "1d"
	AttackLayer7TimeseriesGroupsV2ParamsAggInterval1w  AttackLayer7TimeseriesGroupsV2ParamsAggInterval = "1w"
)

func (r AttackLayer7TimeseriesGroupsV2ParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupsV2ParamsAggInterval15m, AttackLayer7TimeseriesGroupsV2ParamsAggInterval1h, AttackLayer7TimeseriesGroupsV2ParamsAggInterval1d, AttackLayer7TimeseriesGroupsV2ParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer7TimeseriesGroupsV2ParamsFormat string

const (
	AttackLayer7TimeseriesGroupsV2ParamsFormatJson AttackLayer7TimeseriesGroupsV2ParamsFormat = "JSON"
	AttackLayer7TimeseriesGroupsV2ParamsFormatCsv  AttackLayer7TimeseriesGroupsV2ParamsFormat = "CSV"
)

func (r AttackLayer7TimeseriesGroupsV2ParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupsV2ParamsFormatJson, AttackLayer7TimeseriesGroupsV2ParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod string

const (
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodGet             AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "GET"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodPost            AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "POST"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodDelete          AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "DELETE"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodPut             AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "PUT"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodHead            AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "HEAD"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodPurge           AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "PURGE"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodOptions         AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "OPTIONS"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodPropfind        AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "PROPFIND"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodMkcol           AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "MKCOL"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodPatch           AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "PATCH"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodACL             AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "ACL"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodBcopy           AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "BCOPY"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodBdelete         AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "BDELETE"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodBmove           AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "BMOVE"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodBpropfind       AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "BPROPFIND"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodBproppatch      AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodCheckin         AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "CHECKIN"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodCheckout        AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "CHECKOUT"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodConnect         AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "CONNECT"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodCopy            AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "COPY"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodLabel           AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "LABEL"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodLock            AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "LOCK"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodMerge           AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "MERGE"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodMkactivity      AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodMkworkspace     AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodMove            AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "MOVE"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodNotify          AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "NOTIFY"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodOrderpatch      AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodPoll            AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "POLL"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodProppatch       AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "PROPPATCH"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodReport          AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "REPORT"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodSearch          AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "SEARCH"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodSubscribe       AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodTrace           AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "TRACE"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodUncheckout      AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodUnlock          AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "UNLOCK"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodUnsubscribe     AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodUpdate          AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "UPDATE"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodVersioncontrol  AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodBaselinecontrol AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodXmsenumatts     AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodRpcOutData      AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodRpcInData       AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodJson            AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "JSON"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodCook            AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "COOK"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodTrack           AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7TimeseriesGroupsV2ParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodGet, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodPost, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodDelete, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodPut, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodHead, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodPurge, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodOptions, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodPropfind, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodMkcol, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodPatch, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodACL, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodBcopy, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodBdelete, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodBmove, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodBpropfind, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodBproppatch, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodCheckin, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodCheckout, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodConnect, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodCopy, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodLabel, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodLock, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodMerge, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodMkactivity, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodMkworkspace, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodMove, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodNotify, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodOrderpatch, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodPoll, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodProppatch, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodReport, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodSearch, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodSubscribe, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodTrace, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodUncheckout, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodUnlock, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodUnsubscribe, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodUpdate, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodVersioncontrol, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodBaselinecontrol, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodXmsenumatts, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodRpcOutData, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodRpcInData, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodJson, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodCook, AttackLayer7TimeseriesGroupsV2ParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupsV2ParamsHTTPVersion string

const (
	AttackLayer7TimeseriesGroupsV2ParamsHTTPVersionHttPv1 AttackLayer7TimeseriesGroupsV2ParamsHTTPVersion = "HTTPv1"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPVersionHttPv2 AttackLayer7TimeseriesGroupsV2ParamsHTTPVersion = "HTTPv2"
	AttackLayer7TimeseriesGroupsV2ParamsHTTPVersionHttPv3 AttackLayer7TimeseriesGroupsV2ParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7TimeseriesGroupsV2ParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupsV2ParamsHTTPVersionHttPv1, AttackLayer7TimeseriesGroupsV2ParamsHTTPVersionHttPv2, AttackLayer7TimeseriesGroupsV2ParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupsV2ParamsIPVersion string

const (
	AttackLayer7TimeseriesGroupsV2ParamsIPVersionIPv4 AttackLayer7TimeseriesGroupsV2ParamsIPVersion = "IPv4"
	AttackLayer7TimeseriesGroupsV2ParamsIPVersionIPv6 AttackLayer7TimeseriesGroupsV2ParamsIPVersion = "IPv6"
)

func (r AttackLayer7TimeseriesGroupsV2ParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupsV2ParamsIPVersionIPv4, AttackLayer7TimeseriesGroupsV2ParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupsV2ParamsMitigationProduct string

const (
	AttackLayer7TimeseriesGroupsV2ParamsMitigationProductDDoS               AttackLayer7TimeseriesGroupsV2ParamsMitigationProduct = "DDOS"
	AttackLayer7TimeseriesGroupsV2ParamsMitigationProductWAF                AttackLayer7TimeseriesGroupsV2ParamsMitigationProduct = "WAF"
	AttackLayer7TimeseriesGroupsV2ParamsMitigationProductBotManagement      AttackLayer7TimeseriesGroupsV2ParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TimeseriesGroupsV2ParamsMitigationProductAccessRules        AttackLayer7TimeseriesGroupsV2ParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TimeseriesGroupsV2ParamsMitigationProductIPReputation       AttackLayer7TimeseriesGroupsV2ParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TimeseriesGroupsV2ParamsMitigationProductAPIShield          AttackLayer7TimeseriesGroupsV2ParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TimeseriesGroupsV2ParamsMitigationProductDataLossPrevention AttackLayer7TimeseriesGroupsV2ParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TimeseriesGroupsV2ParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupsV2ParamsMitigationProductDDoS, AttackLayer7TimeseriesGroupsV2ParamsMitigationProductWAF, AttackLayer7TimeseriesGroupsV2ParamsMitigationProductBotManagement, AttackLayer7TimeseriesGroupsV2ParamsMitigationProductAccessRules, AttackLayer7TimeseriesGroupsV2ParamsMitigationProductIPReputation, AttackLayer7TimeseriesGroupsV2ParamsMitigationProductAPIShield, AttackLayer7TimeseriesGroupsV2ParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesGroupsV2ParamsNormalization string

const (
	AttackLayer7TimeseriesGroupsV2ParamsNormalizationPercentage AttackLayer7TimeseriesGroupsV2ParamsNormalization = "PERCENTAGE"
	AttackLayer7TimeseriesGroupsV2ParamsNormalizationMin0Max    AttackLayer7TimeseriesGroupsV2ParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer7TimeseriesGroupsV2ParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesGroupsV2ParamsNormalizationPercentage, AttackLayer7TimeseriesGroupsV2ParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer7TimeseriesGroupsV2ResponseEnvelope struct {
	Result  AttackLayer7TimeseriesGroupsV2Response             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    attackLayer7TimeseriesGroupsV2ResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TimeseriesGroupsV2ResponseEnvelopeJSON contains the JSON metadata
// for the struct [AttackLayer7TimeseriesGroupsV2ResponseEnvelope]
type attackLayer7TimeseriesGroupsV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesGroupsV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesGroupsV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
