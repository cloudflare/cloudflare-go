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

// PostQuantumOriginService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPostQuantumOriginService] method instead.
type PostQuantumOriginService struct {
	Options []option.RequestOption
}

// NewPostQuantumOriginService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPostQuantumOriginService(opts ...option.RequestOption) (r *PostQuantumOriginService) {
	r = &PostQuantumOriginService{}
	r.Options = opts
	return
}

// Returns a summary of origin post-quantum data grouped by the specified
// dimension.
func (r *PostQuantumOriginService) Summary(ctx context.Context, dimension PostQuantumOriginSummaryParamsDimension, query PostQuantumOriginSummaryParams, opts ...option.RequestOption) (res *PostQuantumOriginSummaryResponse, err error) {
	var env PostQuantumOriginSummaryResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/post_quantum/origin/summary/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns a timeseries of origin post-quantum data grouped by the specified
// dimension.
func (r *PostQuantumOriginService) TimeseriesGroups(ctx context.Context, dimension PostQuantumOriginTimeseriesGroupsParamsDimension, query PostQuantumOriginTimeseriesGroupsParams, opts ...option.RequestOption) (res *PostQuantumOriginTimeseriesGroupsResponse, err error) {
	var env PostQuantumOriginTimeseriesGroupsResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/post_quantum/origin/timeseries_groups/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PostQuantumOriginSummaryResponse struct {
	// Metadata for the results.
	Meta     PostQuantumOriginSummaryResponseMeta `json:"meta,required"`
	Summary0 map[string]string                    `json:"summary_0,required"`
	JSON     postQuantumOriginSummaryResponseJSON `json:"-"`
}

// postQuantumOriginSummaryResponseJSON contains the JSON metadata for the struct
// [PostQuantumOriginSummaryResponse]
type postQuantumOriginSummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostQuantumOriginSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postQuantumOriginSummaryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type PostQuantumOriginSummaryResponseMeta struct {
	ConfidenceInfo PostQuantumOriginSummaryResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []PostQuantumOriginSummaryResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization PostQuantumOriginSummaryResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []PostQuantumOriginSummaryResponseMetaUnit `json:"units,required"`
	JSON  postQuantumOriginSummaryResponseMetaJSON   `json:"-"`
}

// postQuantumOriginSummaryResponseMetaJSON contains the JSON metadata for the
// struct [PostQuantumOriginSummaryResponseMeta]
type postQuantumOriginSummaryResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostQuantumOriginSummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postQuantumOriginSummaryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type PostQuantumOriginSummaryResponseMetaConfidenceInfo struct {
	Annotations []PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                  `json:"level,required"`
	JSON  postQuantumOriginSummaryResponseMetaConfidenceInfoJSON `json:"-"`
}

// postQuantumOriginSummaryResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [PostQuantumOriginSummaryResponseMetaConfidenceInfo]
type postQuantumOriginSummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostQuantumOriginSummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postQuantumOriginSummaryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                  `json:"description,required"`
	EndDate     time.Time                                                               `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                             `json:"isInstantaneous,required"`
	LinkedURL       string                                                           `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                        `json:"startDate,required" format:"date-time"`
	JSON            postQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// postQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotation]
type postQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAll                PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBGP                PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBots               PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceCT                 PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNS                PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDos                PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFw                 PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceIQI                PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceNet                PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAll, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBGP, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBots, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceCT, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNS, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDos, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFw, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceIQI, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceNet, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventType string

const (
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventTypeEvent             PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventTypeOutage            PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventTypePipeline          PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventTypeEvent, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventTypeOutage, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventTypePipeline, PostQuantumOriginSummaryResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type PostQuantumOriginSummaryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      postQuantumOriginSummaryResponseMetaDateRangeJSON `json:"-"`
}

// postQuantumOriginSummaryResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [PostQuantumOriginSummaryResponseMetaDateRange]
type postQuantumOriginSummaryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostQuantumOriginSummaryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postQuantumOriginSummaryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type PostQuantumOriginSummaryResponseMetaNormalization string

const (
	PostQuantumOriginSummaryResponseMetaNormalizationPercentage           PostQuantumOriginSummaryResponseMetaNormalization = "PERCENTAGE"
	PostQuantumOriginSummaryResponseMetaNormalizationMin0Max              PostQuantumOriginSummaryResponseMetaNormalization = "MIN0_MAX"
	PostQuantumOriginSummaryResponseMetaNormalizationMinMax               PostQuantumOriginSummaryResponseMetaNormalization = "MIN_MAX"
	PostQuantumOriginSummaryResponseMetaNormalizationRawValues            PostQuantumOriginSummaryResponseMetaNormalization = "RAW_VALUES"
	PostQuantumOriginSummaryResponseMetaNormalizationPercentageChange     PostQuantumOriginSummaryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	PostQuantumOriginSummaryResponseMetaNormalizationRollingAverage       PostQuantumOriginSummaryResponseMetaNormalization = "ROLLING_AVERAGE"
	PostQuantumOriginSummaryResponseMetaNormalizationOverlappedPercentage PostQuantumOriginSummaryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	PostQuantumOriginSummaryResponseMetaNormalizationRatio                PostQuantumOriginSummaryResponseMetaNormalization = "RATIO"
)

func (r PostQuantumOriginSummaryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case PostQuantumOriginSummaryResponseMetaNormalizationPercentage, PostQuantumOriginSummaryResponseMetaNormalizationMin0Max, PostQuantumOriginSummaryResponseMetaNormalizationMinMax, PostQuantumOriginSummaryResponseMetaNormalizationRawValues, PostQuantumOriginSummaryResponseMetaNormalizationPercentageChange, PostQuantumOriginSummaryResponseMetaNormalizationRollingAverage, PostQuantumOriginSummaryResponseMetaNormalizationOverlappedPercentage, PostQuantumOriginSummaryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type PostQuantumOriginSummaryResponseMetaUnit struct {
	Name  string                                       `json:"name,required"`
	Value string                                       `json:"value,required"`
	JSON  postQuantumOriginSummaryResponseMetaUnitJSON `json:"-"`
}

// postQuantumOriginSummaryResponseMetaUnitJSON contains the JSON metadata for the
// struct [PostQuantumOriginSummaryResponseMetaUnit]
type postQuantumOriginSummaryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostQuantumOriginSummaryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postQuantumOriginSummaryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type PostQuantumOriginTimeseriesGroupsResponse struct {
	// Metadata for the results.
	Meta   PostQuantumOriginTimeseriesGroupsResponseMeta   `json:"meta,required"`
	Serie0 PostQuantumOriginTimeseriesGroupsResponseSerie0 `json:"serie_0,required"`
	JSON   postQuantumOriginTimeseriesGroupsResponseJSON   `json:"-"`
}

// postQuantumOriginTimeseriesGroupsResponseJSON contains the JSON metadata for the
// struct [PostQuantumOriginTimeseriesGroupsResponse]
type postQuantumOriginTimeseriesGroupsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostQuantumOriginTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postQuantumOriginTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type PostQuantumOriginTimeseriesGroupsResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    PostQuantumOriginTimeseriesGroupsResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []PostQuantumOriginTimeseriesGroupsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization PostQuantumOriginTimeseriesGroupsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []PostQuantumOriginTimeseriesGroupsResponseMetaUnit `json:"units,required"`
	JSON  postQuantumOriginTimeseriesGroupsResponseMetaJSON   `json:"-"`
}

// postQuantumOriginTimeseriesGroupsResponseMetaJSON contains the JSON metadata for
// the struct [PostQuantumOriginTimeseriesGroupsResponseMeta]
type postQuantumOriginTimeseriesGroupsResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PostQuantumOriginTimeseriesGroupsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postQuantumOriginTimeseriesGroupsResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type PostQuantumOriginTimeseriesGroupsResponseMetaAggInterval string

const (
	PostQuantumOriginTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes PostQuantumOriginTimeseriesGroupsResponseMetaAggInterval = "FIFTEEN_MINUTES"
	PostQuantumOriginTimeseriesGroupsResponseMetaAggIntervalOneHour        PostQuantumOriginTimeseriesGroupsResponseMetaAggInterval = "ONE_HOUR"
	PostQuantumOriginTimeseriesGroupsResponseMetaAggIntervalOneDay         PostQuantumOriginTimeseriesGroupsResponseMetaAggInterval = "ONE_DAY"
	PostQuantumOriginTimeseriesGroupsResponseMetaAggIntervalOneWeek        PostQuantumOriginTimeseriesGroupsResponseMetaAggInterval = "ONE_WEEK"
	PostQuantumOriginTimeseriesGroupsResponseMetaAggIntervalOneMonth       PostQuantumOriginTimeseriesGroupsResponseMetaAggInterval = "ONE_MONTH"
)

func (r PostQuantumOriginTimeseriesGroupsResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case PostQuantumOriginTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes, PostQuantumOriginTimeseriesGroupsResponseMetaAggIntervalOneHour, PostQuantumOriginTimeseriesGroupsResponseMetaAggIntervalOneDay, PostQuantumOriginTimeseriesGroupsResponseMetaAggIntervalOneWeek, PostQuantumOriginTimeseriesGroupsResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfo struct {
	Annotations []PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                           `json:"level,required"`
	JSON  postQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoJSON `json:"-"`
}

// postQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfo]
type postQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                           `json:"description,required"`
	EndDate     time.Time                                                                        `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                      `json:"isInstantaneous,required"`
	LinkedURL       string                                                                    `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                 `json:"startDate,required" format:"date-time"`
	JSON            postQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// postQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotation]
type postQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAll                PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceBGP                PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceBots               PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceCT                 PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNS                PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDos                PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceFw                 PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceIQI                PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceNet                PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAll, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceBGP, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceBots, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceCT, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNS, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDos, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceFw, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceIQI, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceNet, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType string

const (
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeEvent             PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeOutage            PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypePipeline          PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeEvent, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeOutage, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypePipeline, PostQuantumOriginTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type PostQuantumOriginTimeseriesGroupsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      postQuantumOriginTimeseriesGroupsResponseMetaDateRangeJSON `json:"-"`
}

// postQuantumOriginTimeseriesGroupsResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [PostQuantumOriginTimeseriesGroupsResponseMetaDateRange]
type postQuantumOriginTimeseriesGroupsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostQuantumOriginTimeseriesGroupsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postQuantumOriginTimeseriesGroupsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type PostQuantumOriginTimeseriesGroupsResponseMetaNormalization string

const (
	PostQuantumOriginTimeseriesGroupsResponseMetaNormalizationPercentage           PostQuantumOriginTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE"
	PostQuantumOriginTimeseriesGroupsResponseMetaNormalizationMin0Max              PostQuantumOriginTimeseriesGroupsResponseMetaNormalization = "MIN0_MAX"
	PostQuantumOriginTimeseriesGroupsResponseMetaNormalizationMinMax               PostQuantumOriginTimeseriesGroupsResponseMetaNormalization = "MIN_MAX"
	PostQuantumOriginTimeseriesGroupsResponseMetaNormalizationRawValues            PostQuantumOriginTimeseriesGroupsResponseMetaNormalization = "RAW_VALUES"
	PostQuantumOriginTimeseriesGroupsResponseMetaNormalizationPercentageChange     PostQuantumOriginTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	PostQuantumOriginTimeseriesGroupsResponseMetaNormalizationRollingAverage       PostQuantumOriginTimeseriesGroupsResponseMetaNormalization = "ROLLING_AVERAGE"
	PostQuantumOriginTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage PostQuantumOriginTimeseriesGroupsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	PostQuantumOriginTimeseriesGroupsResponseMetaNormalizationRatio                PostQuantumOriginTimeseriesGroupsResponseMetaNormalization = "RATIO"
)

func (r PostQuantumOriginTimeseriesGroupsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case PostQuantumOriginTimeseriesGroupsResponseMetaNormalizationPercentage, PostQuantumOriginTimeseriesGroupsResponseMetaNormalizationMin0Max, PostQuantumOriginTimeseriesGroupsResponseMetaNormalizationMinMax, PostQuantumOriginTimeseriesGroupsResponseMetaNormalizationRawValues, PostQuantumOriginTimeseriesGroupsResponseMetaNormalizationPercentageChange, PostQuantumOriginTimeseriesGroupsResponseMetaNormalizationRollingAverage, PostQuantumOriginTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage, PostQuantumOriginTimeseriesGroupsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type PostQuantumOriginTimeseriesGroupsResponseMetaUnit struct {
	Name  string                                                `json:"name,required"`
	Value string                                                `json:"value,required"`
	JSON  postQuantumOriginTimeseriesGroupsResponseMetaUnitJSON `json:"-"`
}

// postQuantumOriginTimeseriesGroupsResponseMetaUnitJSON contains the JSON metadata
// for the struct [PostQuantumOriginTimeseriesGroupsResponseMetaUnit]
type postQuantumOriginTimeseriesGroupsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostQuantumOriginTimeseriesGroupsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postQuantumOriginTimeseriesGroupsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type PostQuantumOriginTimeseriesGroupsResponseSerie0 struct {
	Timestamps  []time.Time                                         `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                 `json:"-,extras"`
	JSON        postQuantumOriginTimeseriesGroupsResponseSerie0JSON `json:"-"`
}

// postQuantumOriginTimeseriesGroupsResponseSerie0JSON contains the JSON metadata
// for the struct [PostQuantumOriginTimeseriesGroupsResponseSerie0]
type postQuantumOriginTimeseriesGroupsResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostQuantumOriginTimeseriesGroupsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postQuantumOriginTimeseriesGroupsResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type PostQuantumOriginSummaryParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[PostQuantumOriginSummaryParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [PostQuantumOriginSummaryParams]'s query parameters as
// `url.Values`.
func (r PostQuantumOriginSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the origin post-quantum data dimension by which to group the results.
type PostQuantumOriginSummaryParamsDimension string

const (
	PostQuantumOriginSummaryParamsDimensionKeyAgreement PostQuantumOriginSummaryParamsDimension = "KEY_AGREEMENT"
)

func (r PostQuantumOriginSummaryParamsDimension) IsKnown() bool {
	switch r {
	case PostQuantumOriginSummaryParamsDimensionKeyAgreement:
		return true
	}
	return false
}

// Format in which results will be returned.
type PostQuantumOriginSummaryParamsFormat string

const (
	PostQuantumOriginSummaryParamsFormatJson PostQuantumOriginSummaryParamsFormat = "JSON"
	PostQuantumOriginSummaryParamsFormatCsv  PostQuantumOriginSummaryParamsFormat = "CSV"
)

func (r PostQuantumOriginSummaryParamsFormat) IsKnown() bool {
	switch r {
	case PostQuantumOriginSummaryParamsFormatJson, PostQuantumOriginSummaryParamsFormatCsv:
		return true
	}
	return false
}

type PostQuantumOriginSummaryResponseEnvelope struct {
	Result  PostQuantumOriginSummaryResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    postQuantumOriginSummaryResponseEnvelopeJSON `json:"-"`
}

// postQuantumOriginSummaryResponseEnvelopeJSON contains the JSON metadata for the
// struct [PostQuantumOriginSummaryResponseEnvelope]
type postQuantumOriginSummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostQuantumOriginSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postQuantumOriginSummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PostQuantumOriginTimeseriesGroupsParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[PostQuantumOriginTimeseriesGroupsParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [PostQuantumOriginTimeseriesGroupsParams]'s query parameters
// as `url.Values`.
func (r PostQuantumOriginTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the origin post-quantum data dimension by which to group the results.
type PostQuantumOriginTimeseriesGroupsParamsDimension string

const (
	PostQuantumOriginTimeseriesGroupsParamsDimensionKeyAgreement PostQuantumOriginTimeseriesGroupsParamsDimension = "KEY_AGREEMENT"
)

func (r PostQuantumOriginTimeseriesGroupsParamsDimension) IsKnown() bool {
	switch r {
	case PostQuantumOriginTimeseriesGroupsParamsDimensionKeyAgreement:
		return true
	}
	return false
}

// Format in which results will be returned.
type PostQuantumOriginTimeseriesGroupsParamsFormat string

const (
	PostQuantumOriginTimeseriesGroupsParamsFormatJson PostQuantumOriginTimeseriesGroupsParamsFormat = "JSON"
	PostQuantumOriginTimeseriesGroupsParamsFormatCsv  PostQuantumOriginTimeseriesGroupsParamsFormat = "CSV"
)

func (r PostQuantumOriginTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case PostQuantumOriginTimeseriesGroupsParamsFormatJson, PostQuantumOriginTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}

type PostQuantumOriginTimeseriesGroupsResponseEnvelope struct {
	Result  PostQuantumOriginTimeseriesGroupsResponse             `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    postQuantumOriginTimeseriesGroupsResponseEnvelopeJSON `json:"-"`
}

// postQuantumOriginTimeseriesGroupsResponseEnvelopeJSON contains the JSON metadata
// for the struct [PostQuantumOriginTimeseriesGroupsResponseEnvelope]
type postQuantumOriginTimeseriesGroupsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PostQuantumOriginTimeseriesGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r postQuantumOriginTimeseriesGroupsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
