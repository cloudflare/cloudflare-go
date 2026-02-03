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

// AttackLayer3Service contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackLayer3Service] method instead.
type AttackLayer3Service struct {
	Options          []option.RequestOption
	Summary          *AttackLayer3SummaryService
	TimeseriesGroups *AttackLayer3TimeseriesGroupService
	Top              *AttackLayer3TopService
}

// NewAttackLayer3Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAttackLayer3Service(opts ...option.RequestOption) (r *AttackLayer3Service) {
	r = &AttackLayer3Service{}
	r.Options = opts
	r.Summary = NewAttackLayer3SummaryService(opts...)
	r.TimeseriesGroups = NewAttackLayer3TimeseriesGroupService(opts...)
	r.Top = NewAttackLayer3TopService(opts...)
	return
}

// Retrieves the distribution of layer 3 attacks by the specified dimension.
func (r *AttackLayer3Service) SummaryV2(ctx context.Context, dimension AttackLayer3SummaryV2ParamsDimension, query AttackLayer3SummaryV2Params, opts ...option.RequestOption) (res *AttackLayer3SummaryV2Response, err error) {
	var env AttackLayer3SummaryV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/attacks/layer3/summary/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves layer 3 attacks over time.
func (r *AttackLayer3Service) Timeseries(ctx context.Context, query AttackLayer3TimeseriesParams, opts ...option.RequestOption) (res *AttackLayer3TimeseriesResponse, err error) {
	var env AttackLayer3TimeseriesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/attacks/layer3/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of layer 3 attacks grouped by dimension over time.
func (r *AttackLayer3Service) TimeseriesGroupsV2(ctx context.Context, dimension AttackLayer3TimeseriesGroupsV2ParamsDimension, query AttackLayer3TimeseriesGroupsV2Params, opts ...option.RequestOption) (res *AttackLayer3TimeseriesGroupsV2Response, err error) {
	var env AttackLayer3TimeseriesGroupsV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/attacks/layer3/timeseries_groups/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer3SummaryV2Response struct {
	// Metadata for the results.
	Meta     AttackLayer3SummaryV2ResponseMeta `json:"meta,required"`
	Summary0 map[string]string                 `json:"summary_0,required"`
	JSON     attackLayer3SummaryV2ResponseJSON `json:"-"`
}

// attackLayer3SummaryV2ResponseJSON contains the JSON metadata for the struct
// [AttackLayer3SummaryV2Response]
type attackLayer3SummaryV2ResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3SummaryV2ResponseMeta struct {
	ConfidenceInfo AttackLayer3SummaryV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer3SummaryV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3SummaryV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3SummaryV2ResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3SummaryV2ResponseMetaJSON   `json:"-"`
}

// attackLayer3SummaryV2ResponseMetaJSON contains the JSON metadata for the struct
// [AttackLayer3SummaryV2ResponseMeta]
type attackLayer3SummaryV2ResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3SummaryV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryV2ResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                               `json:"level,required"`
	JSON  attackLayer3SummaryV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3SummaryV2ResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AttackLayer3SummaryV2ResponseMetaConfidenceInfo]
type attackLayer3SummaryV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                               `json:"description,required"`
	EndDate     time.Time                                                            `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                          `json:"isInstantaneous,required"`
	LinkedURL       string                                                        `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                     `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotation]
type attackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll                AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots               AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos                AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet                AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AttackLayer3SummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AttackLayer3SummaryV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                      `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3SummaryV2ResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3SummaryV2ResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer3SummaryV2ResponseMetaDateRange]
type attackLayer3SummaryV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3SummaryV2ResponseMetaNormalization string

const (
	AttackLayer3SummaryV2ResponseMetaNormalizationPercentage           AttackLayer3SummaryV2ResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3SummaryV2ResponseMetaNormalizationMin0Max              AttackLayer3SummaryV2ResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3SummaryV2ResponseMetaNormalizationMinMax               AttackLayer3SummaryV2ResponseMetaNormalization = "MIN_MAX"
	AttackLayer3SummaryV2ResponseMetaNormalizationRawValues            AttackLayer3SummaryV2ResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3SummaryV2ResponseMetaNormalizationPercentageChange     AttackLayer3SummaryV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3SummaryV2ResponseMetaNormalizationRollingAverage       AttackLayer3SummaryV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3SummaryV2ResponseMetaNormalizationOverlappedPercentage AttackLayer3SummaryV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3SummaryV2ResponseMetaNormalizationRatio                AttackLayer3SummaryV2ResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3SummaryV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryV2ResponseMetaNormalizationPercentage, AttackLayer3SummaryV2ResponseMetaNormalizationMin0Max, AttackLayer3SummaryV2ResponseMetaNormalizationMinMax, AttackLayer3SummaryV2ResponseMetaNormalizationRawValues, AttackLayer3SummaryV2ResponseMetaNormalizationPercentageChange, AttackLayer3SummaryV2ResponseMetaNormalizationRollingAverage, AttackLayer3SummaryV2ResponseMetaNormalizationOverlappedPercentage, AttackLayer3SummaryV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3SummaryV2ResponseMetaUnit struct {
	Name  string                                    `json:"name,required"`
	Value string                                    `json:"value,required"`
	JSON  attackLayer3SummaryV2ResponseMetaUnitJSON `json:"-"`
}

// attackLayer3SummaryV2ResponseMetaUnitJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryV2ResponseMetaUnit]
type attackLayer3SummaryV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesResponse struct {
	// Metadata for the results.
	Meta        AttackLayer3TimeseriesResponseMeta        `json:"meta,required"`
	ExtraFields map[string]AttackLayer3TimeseriesResponse `json:"-,extras"`
	JSON        attackLayer3TimeseriesResponseJSON        `json:"-"`
}

// attackLayer3TimeseriesResponseJSON contains the JSON metadata for the struct
// [AttackLayer3TimeseriesResponse]
type attackLayer3TimeseriesResponseJSON struct {
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3TimeseriesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AttackLayer3TimeseriesResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AttackLayer3TimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer3TimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3TimeseriesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3TimeseriesResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3TimeseriesResponseMetaJSON   `json:"-"`
}

// attackLayer3TimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [AttackLayer3TimeseriesResponseMeta]
type attackLayer3TimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesResponseMetaAggInterval string

const (
	AttackLayer3TimeseriesResponseMetaAggIntervalFifteenMinutes AttackLayer3TimeseriesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AttackLayer3TimeseriesResponseMetaAggIntervalOneHour        AttackLayer3TimeseriesResponseMetaAggInterval = "ONE_HOUR"
	AttackLayer3TimeseriesResponseMetaAggIntervalOneDay         AttackLayer3TimeseriesResponseMetaAggInterval = "ONE_DAY"
	AttackLayer3TimeseriesResponseMetaAggIntervalOneWeek        AttackLayer3TimeseriesResponseMetaAggInterval = "ONE_WEEK"
	AttackLayer3TimeseriesResponseMetaAggIntervalOneMonth       AttackLayer3TimeseriesResponseMetaAggInterval = "ONE_MONTH"
)

func (r AttackLayer3TimeseriesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesResponseMetaAggIntervalFifteenMinutes, AttackLayer3TimeseriesResponseMetaAggIntervalOneHour, AttackLayer3TimeseriesResponseMetaAggIntervalOneDay, AttackLayer3TimeseriesResponseMetaAggIntervalOneWeek, AttackLayer3TimeseriesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AttackLayer3TimeseriesResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                `json:"level,required"`
	JSON  attackLayer3TimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3TimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesResponseMetaConfidenceInfo]
type attackLayer3TimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                `json:"description,required"`
	EndDate     time.Time                                                             `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                      `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotation]
type attackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll                AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots               AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos                AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet                AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCT, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AttackLayer3TimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AttackLayer3TimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3TimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3TimeseriesResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer3TimeseriesResponseMetaDateRange]
type attackLayer3TimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesResponseMetaNormalization string

const (
	AttackLayer3TimeseriesResponseMetaNormalizationPercentage           AttackLayer3TimeseriesResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesResponseMetaNormalizationMin0Max              AttackLayer3TimeseriesResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3TimeseriesResponseMetaNormalizationMinMax               AttackLayer3TimeseriesResponseMetaNormalization = "MIN_MAX"
	AttackLayer3TimeseriesResponseMetaNormalizationRawValues            AttackLayer3TimeseriesResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3TimeseriesResponseMetaNormalizationPercentageChange     AttackLayer3TimeseriesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3TimeseriesResponseMetaNormalizationRollingAverage       AttackLayer3TimeseriesResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3TimeseriesResponseMetaNormalizationOverlappedPercentage AttackLayer3TimeseriesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3TimeseriesResponseMetaNormalizationRatio                AttackLayer3TimeseriesResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3TimeseriesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesResponseMetaNormalizationPercentage, AttackLayer3TimeseriesResponseMetaNormalizationMin0Max, AttackLayer3TimeseriesResponseMetaNormalizationMinMax, AttackLayer3TimeseriesResponseMetaNormalizationRawValues, AttackLayer3TimeseriesResponseMetaNormalizationPercentageChange, AttackLayer3TimeseriesResponseMetaNormalizationRollingAverage, AttackLayer3TimeseriesResponseMetaNormalizationOverlappedPercentage, AttackLayer3TimeseriesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3TimeseriesResponseMetaUnit struct {
	Name  string                                     `json:"name,required"`
	Value string                                     `json:"value,required"`
	JSON  attackLayer3TimeseriesResponseMetaUnitJSON `json:"-"`
}

// attackLayer3TimeseriesResponseMetaUnitJSON contains the JSON metadata for the
// struct [AttackLayer3TimeseriesResponseMetaUnit]
type attackLayer3TimeseriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupsV2Response struct {
	// Metadata for the results.
	Meta   AttackLayer3TimeseriesGroupsV2ResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer3TimeseriesGroupsV2ResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer3TimeseriesGroupsV2ResponseJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupsV2ResponseJSON contains the JSON metadata for the
// struct [AttackLayer3TimeseriesGroupsV2Response]
type attackLayer3TimeseriesGroupsV2ResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupsV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupsV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer3TimeseriesGroupsV2ResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AttackLayer3TimeseriesGroupsV2ResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AttackLayer3TimeseriesGroupsV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer3TimeseriesGroupsV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer3TimeseriesGroupsV2ResponseMetaUnit `json:"units,required"`
	JSON  attackLayer3TimeseriesGroupsV2ResponseMetaJSON   `json:"-"`
}

// attackLayer3TimeseriesGroupsV2ResponseMetaJSON contains the JSON metadata for
// the struct [AttackLayer3TimeseriesGroupsV2ResponseMeta]
type attackLayer3TimeseriesGroupsV2ResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupsV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupsV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupsV2ResponseMetaAggInterval string

const (
	AttackLayer3TimeseriesGroupsV2ResponseMetaAggIntervalFifteenMinutes AttackLayer3TimeseriesGroupsV2ResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AttackLayer3TimeseriesGroupsV2ResponseMetaAggIntervalOneHour        AttackLayer3TimeseriesGroupsV2ResponseMetaAggInterval = "ONE_HOUR"
	AttackLayer3TimeseriesGroupsV2ResponseMetaAggIntervalOneDay         AttackLayer3TimeseriesGroupsV2ResponseMetaAggInterval = "ONE_DAY"
	AttackLayer3TimeseriesGroupsV2ResponseMetaAggIntervalOneWeek        AttackLayer3TimeseriesGroupsV2ResponseMetaAggInterval = "ONE_WEEK"
	AttackLayer3TimeseriesGroupsV2ResponseMetaAggIntervalOneMonth       AttackLayer3TimeseriesGroupsV2ResponseMetaAggInterval = "ONE_MONTH"
)

func (r AttackLayer3TimeseriesGroupsV2ResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupsV2ResponseMetaAggIntervalFifteenMinutes, AttackLayer3TimeseriesGroupsV2ResponseMetaAggIntervalOneHour, AttackLayer3TimeseriesGroupsV2ResponseMetaAggIntervalOneDay, AttackLayer3TimeseriesGroupsV2ResponseMetaAggIntervalOneWeek, AttackLayer3TimeseriesGroupsV2ResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                        `json:"level,required"`
	JSON  attackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfo]
type attackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                        `json:"description,required"`
	EndDate     time.Time                                                                     `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                              `json:"startDate,required" format:"date-time"`
	JSON            attackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation]
type attackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll                AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots               AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos                AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet                AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AttackLayer3TimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupsV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3TimeseriesGroupsV2ResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupsV2ResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AttackLayer3TimeseriesGroupsV2ResponseMetaDateRange]
type attackLayer3TimeseriesGroupsV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupsV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupsV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupsV2ResponseMetaNormalization string

const (
	AttackLayer3TimeseriesGroupsV2ResponseMetaNormalizationPercentage           AttackLayer3TimeseriesGroupsV2ResponseMetaNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupsV2ResponseMetaNormalizationMin0Max              AttackLayer3TimeseriesGroupsV2ResponseMetaNormalization = "MIN0_MAX"
	AttackLayer3TimeseriesGroupsV2ResponseMetaNormalizationMinMax               AttackLayer3TimeseriesGroupsV2ResponseMetaNormalization = "MIN_MAX"
	AttackLayer3TimeseriesGroupsV2ResponseMetaNormalizationRawValues            AttackLayer3TimeseriesGroupsV2ResponseMetaNormalization = "RAW_VALUES"
	AttackLayer3TimeseriesGroupsV2ResponseMetaNormalizationPercentageChange     AttackLayer3TimeseriesGroupsV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3TimeseriesGroupsV2ResponseMetaNormalizationRollingAverage       AttackLayer3TimeseriesGroupsV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer3TimeseriesGroupsV2ResponseMetaNormalizationOverlappedPercentage AttackLayer3TimeseriesGroupsV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer3TimeseriesGroupsV2ResponseMetaNormalizationRatio                AttackLayer3TimeseriesGroupsV2ResponseMetaNormalization = "RATIO"
)

func (r AttackLayer3TimeseriesGroupsV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupsV2ResponseMetaNormalizationPercentage, AttackLayer3TimeseriesGroupsV2ResponseMetaNormalizationMin0Max, AttackLayer3TimeseriesGroupsV2ResponseMetaNormalizationMinMax, AttackLayer3TimeseriesGroupsV2ResponseMetaNormalizationRawValues, AttackLayer3TimeseriesGroupsV2ResponseMetaNormalizationPercentageChange, AttackLayer3TimeseriesGroupsV2ResponseMetaNormalizationRollingAverage, AttackLayer3TimeseriesGroupsV2ResponseMetaNormalizationOverlappedPercentage, AttackLayer3TimeseriesGroupsV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupsV2ResponseMetaUnit struct {
	Name  string                                             `json:"name,required"`
	Value string                                             `json:"value,required"`
	JSON  attackLayer3TimeseriesGroupsV2ResponseMetaUnitJSON `json:"-"`
}

// attackLayer3TimeseriesGroupsV2ResponseMetaUnitJSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupsV2ResponseMetaUnit]
type attackLayer3TimeseriesGroupsV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupsV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupsV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupsV2ResponseSerie0 struct {
	Timestamps  []time.Time                                      `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                              `json:"-,extras"`
	JSON        attackLayer3TimeseriesGroupsV2ResponseSerie0JSON `json:"-"`
}

// attackLayer3TimeseriesGroupsV2ResponseSerie0JSON contains the JSON metadata for
// the struct [AttackLayer3TimeseriesGroupsV2ResponseSerie0]
type attackLayer3TimeseriesGroupsV2ResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupsV2ResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupsV2ResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3SummaryV2Params struct {
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[AttackLayer3SummaryV2ParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[AttackLayer3SummaryV2ParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3SummaryV2ParamsIPVersion] `query:"ipVersion"`
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
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3SummaryV2ParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3SummaryV2Params]'s query parameters as
// `url.Values`.
func (r AttackLayer3SummaryV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type AttackLayer3SummaryV2ParamsDimension string

const (
	AttackLayer3SummaryV2ParamsDimensionProtocol  AttackLayer3SummaryV2ParamsDimension = "PROTOCOL"
	AttackLayer3SummaryV2ParamsDimensionIPVersion AttackLayer3SummaryV2ParamsDimension = "IP_VERSION"
	AttackLayer3SummaryV2ParamsDimensionVector    AttackLayer3SummaryV2ParamsDimension = "VECTOR"
	AttackLayer3SummaryV2ParamsDimensionDuration  AttackLayer3SummaryV2ParamsDimension = "DURATION"
	AttackLayer3SummaryV2ParamsDimensionBitrate   AttackLayer3SummaryV2ParamsDimension = "BITRATE"
	AttackLayer3SummaryV2ParamsDimensionVertical  AttackLayer3SummaryV2ParamsDimension = "VERTICAL"
	AttackLayer3SummaryV2ParamsDimensionIndustry  AttackLayer3SummaryV2ParamsDimension = "INDUSTRY"
)

func (r AttackLayer3SummaryV2ParamsDimension) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryV2ParamsDimensionProtocol, AttackLayer3SummaryV2ParamsDimensionIPVersion, AttackLayer3SummaryV2ParamsDimensionVector, AttackLayer3SummaryV2ParamsDimensionDuration, AttackLayer3SummaryV2ParamsDimensionBitrate, AttackLayer3SummaryV2ParamsDimensionVertical, AttackLayer3SummaryV2ParamsDimensionIndustry:
		return true
	}
	return false
}

// Specifies whether the `location` filter applies to the source or target
// location.
type AttackLayer3SummaryV2ParamsDirection string

const (
	AttackLayer3SummaryV2ParamsDirectionOrigin AttackLayer3SummaryV2ParamsDirection = "ORIGIN"
	AttackLayer3SummaryV2ParamsDirectionTarget AttackLayer3SummaryV2ParamsDirection = "TARGET"
)

func (r AttackLayer3SummaryV2ParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryV2ParamsDirectionOrigin, AttackLayer3SummaryV2ParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer3SummaryV2ParamsFormat string

const (
	AttackLayer3SummaryV2ParamsFormatJson AttackLayer3SummaryV2ParamsFormat = "JSON"
	AttackLayer3SummaryV2ParamsFormatCsv  AttackLayer3SummaryV2ParamsFormat = "CSV"
)

func (r AttackLayer3SummaryV2ParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryV2ParamsFormatJson, AttackLayer3SummaryV2ParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3SummaryV2ParamsIPVersion string

const (
	AttackLayer3SummaryV2ParamsIPVersionIPv4 AttackLayer3SummaryV2ParamsIPVersion = "IPv4"
	AttackLayer3SummaryV2ParamsIPVersionIPv6 AttackLayer3SummaryV2ParamsIPVersion = "IPv6"
)

func (r AttackLayer3SummaryV2ParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryV2ParamsIPVersionIPv4, AttackLayer3SummaryV2ParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer3SummaryV2ParamsProtocol string

const (
	AttackLayer3SummaryV2ParamsProtocolUdp  AttackLayer3SummaryV2ParamsProtocol = "UDP"
	AttackLayer3SummaryV2ParamsProtocolTCP  AttackLayer3SummaryV2ParamsProtocol = "TCP"
	AttackLayer3SummaryV2ParamsProtocolIcmp AttackLayer3SummaryV2ParamsProtocol = "ICMP"
	AttackLayer3SummaryV2ParamsProtocolGRE  AttackLayer3SummaryV2ParamsProtocol = "GRE"
)

func (r AttackLayer3SummaryV2ParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3SummaryV2ParamsProtocolUdp, AttackLayer3SummaryV2ParamsProtocolTCP, AttackLayer3SummaryV2ParamsProtocolIcmp, AttackLayer3SummaryV2ParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3SummaryV2ResponseEnvelope struct {
	Result  AttackLayer3SummaryV2Response             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    attackLayer3SummaryV2ResponseEnvelopeJSON `json:"-"`
}

// attackLayer3SummaryV2ResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer3SummaryV2ResponseEnvelope]
type attackLayer3SummaryV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3SummaryV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3SummaryV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer3TimeseriesParamsAggInterval] `query:"aggInterval"`
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[AttackLayer3TimeseriesParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[AttackLayer3TimeseriesParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3TimeseriesParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Measurement units, eg. bytes.
	Metric param.Field[AttackLayer3TimeseriesParamsMetric] `query:"metric"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer3TimeseriesParamsNormalization] `query:"normalization"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3TimeseriesParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TimeseriesParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3TimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesParamsAggInterval string

const (
	AttackLayer3TimeseriesParamsAggInterval15m AttackLayer3TimeseriesParamsAggInterval = "15m"
	AttackLayer3TimeseriesParamsAggInterval1h  AttackLayer3TimeseriesParamsAggInterval = "1h"
	AttackLayer3TimeseriesParamsAggInterval1d  AttackLayer3TimeseriesParamsAggInterval = "1d"
	AttackLayer3TimeseriesParamsAggInterval1w  AttackLayer3TimeseriesParamsAggInterval = "1w"
)

func (r AttackLayer3TimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesParamsAggInterval15m, AttackLayer3TimeseriesParamsAggInterval1h, AttackLayer3TimeseriesParamsAggInterval1d, AttackLayer3TimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Specifies whether the `location` filter applies to the source or target
// location.
type AttackLayer3TimeseriesParamsDirection string

const (
	AttackLayer3TimeseriesParamsDirectionOrigin AttackLayer3TimeseriesParamsDirection = "ORIGIN"
	AttackLayer3TimeseriesParamsDirectionTarget AttackLayer3TimeseriesParamsDirection = "TARGET"
)

func (r AttackLayer3TimeseriesParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesParamsDirectionOrigin, AttackLayer3TimeseriesParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer3TimeseriesParamsFormat string

const (
	AttackLayer3TimeseriesParamsFormatJson AttackLayer3TimeseriesParamsFormat = "JSON"
	AttackLayer3TimeseriesParamsFormatCsv  AttackLayer3TimeseriesParamsFormat = "CSV"
)

func (r AttackLayer3TimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesParamsFormatJson, AttackLayer3TimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3TimeseriesParamsIPVersion string

const (
	AttackLayer3TimeseriesParamsIPVersionIPv4 AttackLayer3TimeseriesParamsIPVersion = "IPv4"
	AttackLayer3TimeseriesParamsIPVersionIPv6 AttackLayer3TimeseriesParamsIPVersion = "IPv6"
)

func (r AttackLayer3TimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesParamsIPVersionIPv4, AttackLayer3TimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

// Measurement units, eg. bytes.
type AttackLayer3TimeseriesParamsMetric string

const (
	AttackLayer3TimeseriesParamsMetricBytes    AttackLayer3TimeseriesParamsMetric = "BYTES"
	AttackLayer3TimeseriesParamsMetricBytesOld AttackLayer3TimeseriesParamsMetric = "BYTES_OLD"
)

func (r AttackLayer3TimeseriesParamsMetric) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesParamsMetricBytes, AttackLayer3TimeseriesParamsMetricBytesOld:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesParamsNormalization string

const (
	AttackLayer3TimeseriesParamsNormalizationPercentageChange AttackLayer3TimeseriesParamsNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3TimeseriesParamsNormalizationMin0Max          AttackLayer3TimeseriesParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer3TimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesParamsNormalizationPercentageChange, AttackLayer3TimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer3TimeseriesParamsProtocol string

const (
	AttackLayer3TimeseriesParamsProtocolUdp  AttackLayer3TimeseriesParamsProtocol = "UDP"
	AttackLayer3TimeseriesParamsProtocolTCP  AttackLayer3TimeseriesParamsProtocol = "TCP"
	AttackLayer3TimeseriesParamsProtocolIcmp AttackLayer3TimeseriesParamsProtocol = "ICMP"
	AttackLayer3TimeseriesParamsProtocolGRE  AttackLayer3TimeseriesParamsProtocol = "GRE"
)

func (r AttackLayer3TimeseriesParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesParamsProtocolUdp, AttackLayer3TimeseriesParamsProtocolTCP, AttackLayer3TimeseriesParamsProtocolIcmp, AttackLayer3TimeseriesParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3TimeseriesResponseEnvelope struct {
	Result  AttackLayer3TimeseriesResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    attackLayer3TimeseriesResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TimeseriesResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer3TimeseriesResponseEnvelope]
type attackLayer3TimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesGroupsV2Params struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer3TimeseriesGroupsV2ParamsAggInterval] `query:"aggInterval"`
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[AttackLayer3TimeseriesGroupsV2ParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[AttackLayer3TimeseriesGroupsV2ParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3TimeseriesGroupsV2ParamsIPVersion] `query:"ipVersion"`
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
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer3TimeseriesGroupsV2ParamsNormalization] `query:"normalization"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3TimeseriesGroupsV2ParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TimeseriesGroupsV2Params]'s query parameters as
// `url.Values`.
func (r AttackLayer3TimeseriesGroupsV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type AttackLayer3TimeseriesGroupsV2ParamsDimension string

const (
	AttackLayer3TimeseriesGroupsV2ParamsDimensionProtocol  AttackLayer3TimeseriesGroupsV2ParamsDimension = "PROTOCOL"
	AttackLayer3TimeseriesGroupsV2ParamsDimensionIPVersion AttackLayer3TimeseriesGroupsV2ParamsDimension = "IP_VERSION"
	AttackLayer3TimeseriesGroupsV2ParamsDimensionVector    AttackLayer3TimeseriesGroupsV2ParamsDimension = "VECTOR"
	AttackLayer3TimeseriesGroupsV2ParamsDimensionDuration  AttackLayer3TimeseriesGroupsV2ParamsDimension = "DURATION"
	AttackLayer3TimeseriesGroupsV2ParamsDimensionBitrate   AttackLayer3TimeseriesGroupsV2ParamsDimension = "BITRATE"
	AttackLayer3TimeseriesGroupsV2ParamsDimensionVertical  AttackLayer3TimeseriesGroupsV2ParamsDimension = "VERTICAL"
	AttackLayer3TimeseriesGroupsV2ParamsDimensionIndustry  AttackLayer3TimeseriesGroupsV2ParamsDimension = "INDUSTRY"
)

func (r AttackLayer3TimeseriesGroupsV2ParamsDimension) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupsV2ParamsDimensionProtocol, AttackLayer3TimeseriesGroupsV2ParamsDimensionIPVersion, AttackLayer3TimeseriesGroupsV2ParamsDimensionVector, AttackLayer3TimeseriesGroupsV2ParamsDimensionDuration, AttackLayer3TimeseriesGroupsV2ParamsDimensionBitrate, AttackLayer3TimeseriesGroupsV2ParamsDimensionVertical, AttackLayer3TimeseriesGroupsV2ParamsDimensionIndustry:
		return true
	}
	return false
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesGroupsV2ParamsAggInterval string

const (
	AttackLayer3TimeseriesGroupsV2ParamsAggInterval15m AttackLayer3TimeseriesGroupsV2ParamsAggInterval = "15m"
	AttackLayer3TimeseriesGroupsV2ParamsAggInterval1h  AttackLayer3TimeseriesGroupsV2ParamsAggInterval = "1h"
	AttackLayer3TimeseriesGroupsV2ParamsAggInterval1d  AttackLayer3TimeseriesGroupsV2ParamsAggInterval = "1d"
	AttackLayer3TimeseriesGroupsV2ParamsAggInterval1w  AttackLayer3TimeseriesGroupsV2ParamsAggInterval = "1w"
)

func (r AttackLayer3TimeseriesGroupsV2ParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupsV2ParamsAggInterval15m, AttackLayer3TimeseriesGroupsV2ParamsAggInterval1h, AttackLayer3TimeseriesGroupsV2ParamsAggInterval1d, AttackLayer3TimeseriesGroupsV2ParamsAggInterval1w:
		return true
	}
	return false
}

// Specifies whether the `location` filter applies to the source or target
// location.
type AttackLayer3TimeseriesGroupsV2ParamsDirection string

const (
	AttackLayer3TimeseriesGroupsV2ParamsDirectionOrigin AttackLayer3TimeseriesGroupsV2ParamsDirection = "ORIGIN"
	AttackLayer3TimeseriesGroupsV2ParamsDirectionTarget AttackLayer3TimeseriesGroupsV2ParamsDirection = "TARGET"
)

func (r AttackLayer3TimeseriesGroupsV2ParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupsV2ParamsDirectionOrigin, AttackLayer3TimeseriesGroupsV2ParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer3TimeseriesGroupsV2ParamsFormat string

const (
	AttackLayer3TimeseriesGroupsV2ParamsFormatJson AttackLayer3TimeseriesGroupsV2ParamsFormat = "JSON"
	AttackLayer3TimeseriesGroupsV2ParamsFormatCsv  AttackLayer3TimeseriesGroupsV2ParamsFormat = "CSV"
)

func (r AttackLayer3TimeseriesGroupsV2ParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupsV2ParamsFormatJson, AttackLayer3TimeseriesGroupsV2ParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupsV2ParamsIPVersion string

const (
	AttackLayer3TimeseriesGroupsV2ParamsIPVersionIPv4 AttackLayer3TimeseriesGroupsV2ParamsIPVersion = "IPv4"
	AttackLayer3TimeseriesGroupsV2ParamsIPVersionIPv6 AttackLayer3TimeseriesGroupsV2ParamsIPVersion = "IPv6"
)

func (r AttackLayer3TimeseriesGroupsV2ParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupsV2ParamsIPVersionIPv4, AttackLayer3TimeseriesGroupsV2ParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesGroupsV2ParamsNormalization string

const (
	AttackLayer3TimeseriesGroupsV2ParamsNormalizationPercentage AttackLayer3TimeseriesGroupsV2ParamsNormalization = "PERCENTAGE"
	AttackLayer3TimeseriesGroupsV2ParamsNormalizationMin0Max    AttackLayer3TimeseriesGroupsV2ParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer3TimeseriesGroupsV2ParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupsV2ParamsNormalizationPercentage, AttackLayer3TimeseriesGroupsV2ParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupsV2ParamsProtocol string

const (
	AttackLayer3TimeseriesGroupsV2ParamsProtocolUdp  AttackLayer3TimeseriesGroupsV2ParamsProtocol = "UDP"
	AttackLayer3TimeseriesGroupsV2ParamsProtocolTCP  AttackLayer3TimeseriesGroupsV2ParamsProtocol = "TCP"
	AttackLayer3TimeseriesGroupsV2ParamsProtocolIcmp AttackLayer3TimeseriesGroupsV2ParamsProtocol = "ICMP"
	AttackLayer3TimeseriesGroupsV2ParamsProtocolGRE  AttackLayer3TimeseriesGroupsV2ParamsProtocol = "GRE"
)

func (r AttackLayer3TimeseriesGroupsV2ParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesGroupsV2ParamsProtocolUdp, AttackLayer3TimeseriesGroupsV2ParamsProtocolTCP, AttackLayer3TimeseriesGroupsV2ParamsProtocolIcmp, AttackLayer3TimeseriesGroupsV2ParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3TimeseriesGroupsV2ResponseEnvelope struct {
	Result  AttackLayer3TimeseriesGroupsV2Response             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    attackLayer3TimeseriesGroupsV2ResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TimeseriesGroupsV2ResponseEnvelopeJSON contains the JSON metadata
// for the struct [AttackLayer3TimeseriesGroupsV2ResponseEnvelope]
type attackLayer3TimeseriesGroupsV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesGroupsV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesGroupsV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
