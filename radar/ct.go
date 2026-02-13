// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/tidwall/gjson"
)

// CTService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCTService] method instead.
type CTService struct {
	Options     []option.RequestOption
	Authorities *CTAuthorityService
	Logs        *CTLogService
}

// NewCTService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCTService(opts ...option.RequestOption) (r *CTService) {
	r = &CTService{}
	r.Options = opts
	r.Authorities = NewCTAuthorityService(opts...)
	r.Logs = NewCTLogService(opts...)
	return
}

// Retrieves an aggregated summary of certificates grouped by the specified
// dimension.
func (r *CTService) Summary(ctx context.Context, dimension CTSummaryParamsDimension, query CTSummaryParams, opts ...option.RequestOption) (res *CTSummaryResponse, err error) {
	var env CTSummaryResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/ct/summary/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves certificate volume over time.
func (r *CTService) Timeseries(ctx context.Context, query CTTimeseriesParams, opts ...option.RequestOption) (res *CTTimeseriesResponse, err error) {
	var env CTTimeseriesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/ct/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of certificates grouped by the specified dimension
// over time.
func (r *CTService) TimeseriesGroups(ctx context.Context, dimension CTTimeseriesGroupsParamsDimension, query CTTimeseriesGroupsParams, opts ...option.RequestOption) (res *CTTimeseriesGroupsResponse, err error) {
	var env CTTimeseriesGroupsResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/ct/timeseries_groups/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CTSummaryResponse struct {
	// Metadata for the results.
	Meta     CTSummaryResponseMeta          `json:"meta,required"`
	Summary0 CTSummaryResponseSummary0Union `json:"summary_0,required"`
	JSON     ctSummaryResponseJSON          `json:"-"`
}

// ctSummaryResponseJSON contains the JSON metadata for the struct
// [CTSummaryResponse]
type ctSummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctSummaryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type CTSummaryResponseMeta struct {
	ConfidenceInfo CTSummaryResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []CTSummaryResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization CTSummaryResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []CTSummaryResponseMetaUnit `json:"units,required"`
	JSON  ctSummaryResponseMetaJSON   `json:"-"`
}

// ctSummaryResponseMetaJSON contains the JSON metadata for the struct
// [CTSummaryResponseMeta]
type ctSummaryResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CTSummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctSummaryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type CTSummaryResponseMetaConfidenceInfo struct {
	Annotations []CTSummaryResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                   `json:"level,required"`
	JSON  ctSummaryResponseMetaConfidenceInfoJSON `json:"-"`
}

// ctSummaryResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [CTSummaryResponseMetaConfidenceInfo]
type ctSummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTSummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctSummaryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type CTSummaryResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                   `json:"description,required"`
	EndDate     time.Time                                                `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType CTSummaryResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                              `json:"isInstantaneous,required"`
	LinkedURL       string                                            `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                         `json:"startDate,required" format:"date-time"`
	JSON            ctSummaryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// ctSummaryResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata for
// the struct [CTSummaryResponseMetaConfidenceInfoAnnotation]
type ctSummaryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *CTSummaryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctSummaryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAll                CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBGP                CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBots               CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceCT                 CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNS                CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDos                CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFw                 CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceIQI                CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceNet                CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r CTSummaryResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAll, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBGP, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBots, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceCT, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNS, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDos, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFw, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceIQI, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceNet, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, CTSummaryResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type CTSummaryResponseMetaConfidenceInfoAnnotationsEventType string

const (
	CTSummaryResponseMetaConfidenceInfoAnnotationsEventTypeEvent             CTSummaryResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	CTSummaryResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           CTSummaryResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	CTSummaryResponseMetaConfidenceInfoAnnotationsEventTypeOutage            CTSummaryResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	CTSummaryResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection CTSummaryResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	CTSummaryResponseMetaConfidenceInfoAnnotationsEventTypePipeline          CTSummaryResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	CTSummaryResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    CTSummaryResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r CTSummaryResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case CTSummaryResponseMetaConfidenceInfoAnnotationsEventTypeEvent, CTSummaryResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, CTSummaryResponseMetaConfidenceInfoAnnotationsEventTypeOutage, CTSummaryResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, CTSummaryResponseMetaConfidenceInfoAnnotationsEventTypePipeline, CTSummaryResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type CTSummaryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                          `json:"startTime,required" format:"date-time"`
	JSON      ctSummaryResponseMetaDateRangeJSON `json:"-"`
}

// ctSummaryResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [CTSummaryResponseMetaDateRange]
type ctSummaryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTSummaryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctSummaryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type CTSummaryResponseMetaNormalization string

const (
	CTSummaryResponseMetaNormalizationPercentage           CTSummaryResponseMetaNormalization = "PERCENTAGE"
	CTSummaryResponseMetaNormalizationMin0Max              CTSummaryResponseMetaNormalization = "MIN0_MAX"
	CTSummaryResponseMetaNormalizationMinMax               CTSummaryResponseMetaNormalization = "MIN_MAX"
	CTSummaryResponseMetaNormalizationRawValues            CTSummaryResponseMetaNormalization = "RAW_VALUES"
	CTSummaryResponseMetaNormalizationPercentageChange     CTSummaryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	CTSummaryResponseMetaNormalizationRollingAverage       CTSummaryResponseMetaNormalization = "ROLLING_AVERAGE"
	CTSummaryResponseMetaNormalizationOverlappedPercentage CTSummaryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	CTSummaryResponseMetaNormalizationRatio                CTSummaryResponseMetaNormalization = "RATIO"
)

func (r CTSummaryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case CTSummaryResponseMetaNormalizationPercentage, CTSummaryResponseMetaNormalizationMin0Max, CTSummaryResponseMetaNormalizationMinMax, CTSummaryResponseMetaNormalizationRawValues, CTSummaryResponseMetaNormalizationPercentageChange, CTSummaryResponseMetaNormalizationRollingAverage, CTSummaryResponseMetaNormalizationOverlappedPercentage, CTSummaryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type CTSummaryResponseMetaUnit struct {
	Name  string                        `json:"name,required"`
	Value string                        `json:"value,required"`
	JSON  ctSummaryResponseMetaUnitJSON `json:"-"`
}

// ctSummaryResponseMetaUnitJSON contains the JSON metadata for the struct
// [CTSummaryResponseMetaUnit]
type ctSummaryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTSummaryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctSummaryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [CTSummaryResponseSummary0Map],
// [CTSummaryResponseSummary0Object], [CTSummaryResponseSummary0Object],
// [CTSummaryResponseSummary0Object], [CTSummaryResponseSummary0Object],
// [CTSummaryResponseSummary0Object], [CTSummaryResponseSummary0Object] or
// [CTSummaryResponseSummary0Object].
type CTSummaryResponseSummary0Union interface {
	implementsCTSummaryResponseSummary0Union()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CTSummaryResponseSummary0Union)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CTSummaryResponseSummary0Map{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CTSummaryResponseSummary0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CTSummaryResponseSummary0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CTSummaryResponseSummary0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CTSummaryResponseSummary0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CTSummaryResponseSummary0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CTSummaryResponseSummary0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CTSummaryResponseSummary0Object{}),
		},
	)
}

type CTSummaryResponseSummary0Map map[string]string

func (r CTSummaryResponseSummary0Map) implementsCTSummaryResponseSummary0Union() {}

type CTSummaryResponseSummary0Object struct {
	Rfc6962 string                              `json:"rfc6962,required"`
	Static  string                              `json:"static,required"`
	JSON    ctSummaryResponseSummary0ObjectJSON `json:"-"`
}

// ctSummaryResponseSummary0ObjectJSON contains the JSON metadata for the struct
// [CTSummaryResponseSummary0Object]
type ctSummaryResponseSummary0ObjectJSON struct {
	Rfc6962     apijson.Field
	Static      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTSummaryResponseSummary0Object) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctSummaryResponseSummary0ObjectJSON) RawJSON() string {
	return r.raw
}

func (r CTSummaryResponseSummary0Object) implementsCTSummaryResponseSummary0Union() {}

type CTTimeseriesResponse struct {
	// Metadata for the results.
	Meta        CTTimeseriesResponseMeta        `json:"meta,required"`
	ExtraFields map[string]CTTimeseriesResponse `json:"-,extras"`
	JSON        ctTimeseriesResponseJSON        `json:"-"`
}

// ctTimeseriesResponseJSON contains the JSON metadata for the struct
// [CTTimeseriesResponse]
type ctTimeseriesResponseJSON struct {
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type CTTimeseriesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    CTTimeseriesResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo CTTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []CTTimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization CTTimeseriesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []CTTimeseriesResponseMetaUnit `json:"units,required"`
	JSON  ctTimeseriesResponseMetaJSON   `json:"-"`
}

// ctTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [CTTimeseriesResponseMeta]
type ctTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CTTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type CTTimeseriesResponseMetaAggInterval string

const (
	CTTimeseriesResponseMetaAggIntervalFifteenMinutes CTTimeseriesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	CTTimeseriesResponseMetaAggIntervalOneHour        CTTimeseriesResponseMetaAggInterval = "ONE_HOUR"
	CTTimeseriesResponseMetaAggIntervalOneDay         CTTimeseriesResponseMetaAggInterval = "ONE_DAY"
	CTTimeseriesResponseMetaAggIntervalOneWeek        CTTimeseriesResponseMetaAggInterval = "ONE_WEEK"
	CTTimeseriesResponseMetaAggIntervalOneMonth       CTTimeseriesResponseMetaAggInterval = "ONE_MONTH"
)

func (r CTTimeseriesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case CTTimeseriesResponseMetaAggIntervalFifteenMinutes, CTTimeseriesResponseMetaAggIntervalOneHour, CTTimeseriesResponseMetaAggIntervalOneDay, CTTimeseriesResponseMetaAggIntervalOneWeek, CTTimeseriesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type CTTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []CTTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                      `json:"level,required"`
	JSON  ctTimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// ctTimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [CTTimeseriesResponseMetaConfidenceInfo]
type ctTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type CTTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                      `json:"description,required"`
	EndDate     time.Time                                                   `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                            `json:"startDate,required" format:"date-time"`
	JSON            ctTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// ctTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata
// for the struct [CTTimeseriesResponseMetaConfidenceInfoAnnotation]
type ctTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *CTTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll                CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP                CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots               CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCT                 CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS                CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos                CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw                 CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI                CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet                CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCT, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, CTTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventType string

const (
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent             CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage            CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline          CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent, CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage, CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline, CTTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type CTTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                             `json:"startTime,required" format:"date-time"`
	JSON      ctTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// ctTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [CTTimeseriesResponseMetaDateRange]
type ctTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type CTTimeseriesResponseMetaNormalization string

const (
	CTTimeseriesResponseMetaNormalizationPercentage           CTTimeseriesResponseMetaNormalization = "PERCENTAGE"
	CTTimeseriesResponseMetaNormalizationMin0Max              CTTimeseriesResponseMetaNormalization = "MIN0_MAX"
	CTTimeseriesResponseMetaNormalizationMinMax               CTTimeseriesResponseMetaNormalization = "MIN_MAX"
	CTTimeseriesResponseMetaNormalizationRawValues            CTTimeseriesResponseMetaNormalization = "RAW_VALUES"
	CTTimeseriesResponseMetaNormalizationPercentageChange     CTTimeseriesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	CTTimeseriesResponseMetaNormalizationRollingAverage       CTTimeseriesResponseMetaNormalization = "ROLLING_AVERAGE"
	CTTimeseriesResponseMetaNormalizationOverlappedPercentage CTTimeseriesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	CTTimeseriesResponseMetaNormalizationRatio                CTTimeseriesResponseMetaNormalization = "RATIO"
)

func (r CTTimeseriesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case CTTimeseriesResponseMetaNormalizationPercentage, CTTimeseriesResponseMetaNormalizationMin0Max, CTTimeseriesResponseMetaNormalizationMinMax, CTTimeseriesResponseMetaNormalizationRawValues, CTTimeseriesResponseMetaNormalizationPercentageChange, CTTimeseriesResponseMetaNormalizationRollingAverage, CTTimeseriesResponseMetaNormalizationOverlappedPercentage, CTTimeseriesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type CTTimeseriesResponseMetaUnit struct {
	Name  string                           `json:"name,required"`
	Value string                           `json:"value,required"`
	JSON  ctTimeseriesResponseMetaUnitJSON `json:"-"`
}

// ctTimeseriesResponseMetaUnitJSON contains the JSON metadata for the struct
// [CTTimeseriesResponseMetaUnit]
type ctTimeseriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTTimeseriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type CTTimeseriesGroupsResponse struct {
	// Metadata for the results.
	Meta   CTTimeseriesGroupsResponseMeta   `json:"meta,required"`
	Serie0 CTTimeseriesGroupsResponseSerie0 `json:"serie_0,required"`
	JSON   ctTimeseriesGroupsResponseJSON   `json:"-"`
}

// ctTimeseriesGroupsResponseJSON contains the JSON metadata for the struct
// [CTTimeseriesGroupsResponse]
type ctTimeseriesGroupsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type CTTimeseriesGroupsResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    CTTimeseriesGroupsResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo CTTimeseriesGroupsResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []CTTimeseriesGroupsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization CTTimeseriesGroupsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []CTTimeseriesGroupsResponseMetaUnit `json:"units,required"`
	JSON  ctTimeseriesGroupsResponseMetaJSON   `json:"-"`
}

// ctTimeseriesGroupsResponseMetaJSON contains the JSON metadata for the struct
// [CTTimeseriesGroupsResponseMeta]
type ctTimeseriesGroupsResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CTTimeseriesGroupsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesGroupsResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type CTTimeseriesGroupsResponseMetaAggInterval string

const (
	CTTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes CTTimeseriesGroupsResponseMetaAggInterval = "FIFTEEN_MINUTES"
	CTTimeseriesGroupsResponseMetaAggIntervalOneHour        CTTimeseriesGroupsResponseMetaAggInterval = "ONE_HOUR"
	CTTimeseriesGroupsResponseMetaAggIntervalOneDay         CTTimeseriesGroupsResponseMetaAggInterval = "ONE_DAY"
	CTTimeseriesGroupsResponseMetaAggIntervalOneWeek        CTTimeseriesGroupsResponseMetaAggInterval = "ONE_WEEK"
	CTTimeseriesGroupsResponseMetaAggIntervalOneMonth       CTTimeseriesGroupsResponseMetaAggInterval = "ONE_MONTH"
)

func (r CTTimeseriesGroupsResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case CTTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes, CTTimeseriesGroupsResponseMetaAggIntervalOneHour, CTTimeseriesGroupsResponseMetaAggIntervalOneDay, CTTimeseriesGroupsResponseMetaAggIntervalOneWeek, CTTimeseriesGroupsResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type CTTimeseriesGroupsResponseMetaConfidenceInfo struct {
	Annotations []CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                            `json:"level,required"`
	JSON  ctTimeseriesGroupsResponseMetaConfidenceInfoJSON `json:"-"`
}

// ctTimeseriesGroupsResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [CTTimeseriesGroupsResponseMetaConfidenceInfo]
type ctTimeseriesGroupsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTTimeseriesGroupsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesGroupsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                            `json:"description,required"`
	EndDate     time.Time                                                         `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                       `json:"isInstantaneous,required"`
	LinkedURL       string                                                     `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                  `json:"startDate,required" format:"date-time"`
	JSON            ctTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// ctTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotation]
type ctTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAll                CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceBGP                CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceBots               CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceCT                 CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNS                CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDos                CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceFw                 CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceIQI                CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceNet                CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAll, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceBGP, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceBots, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceCT, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNS, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDos, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceFw, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceIQI, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceNet, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType string

const (
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeEvent             CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeOutage            CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypePipeline          CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeEvent, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeOutage, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypePipeline, CTTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type CTTimeseriesGroupsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                   `json:"startTime,required" format:"date-time"`
	JSON      ctTimeseriesGroupsResponseMetaDateRangeJSON `json:"-"`
}

// ctTimeseriesGroupsResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [CTTimeseriesGroupsResponseMetaDateRange]
type ctTimeseriesGroupsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTTimeseriesGroupsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesGroupsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type CTTimeseriesGroupsResponseMetaNormalization string

const (
	CTTimeseriesGroupsResponseMetaNormalizationPercentage           CTTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE"
	CTTimeseriesGroupsResponseMetaNormalizationMin0Max              CTTimeseriesGroupsResponseMetaNormalization = "MIN0_MAX"
	CTTimeseriesGroupsResponseMetaNormalizationMinMax               CTTimeseriesGroupsResponseMetaNormalization = "MIN_MAX"
	CTTimeseriesGroupsResponseMetaNormalizationRawValues            CTTimeseriesGroupsResponseMetaNormalization = "RAW_VALUES"
	CTTimeseriesGroupsResponseMetaNormalizationPercentageChange     CTTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	CTTimeseriesGroupsResponseMetaNormalizationRollingAverage       CTTimeseriesGroupsResponseMetaNormalization = "ROLLING_AVERAGE"
	CTTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage CTTimeseriesGroupsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	CTTimeseriesGroupsResponseMetaNormalizationRatio                CTTimeseriesGroupsResponseMetaNormalization = "RATIO"
)

func (r CTTimeseriesGroupsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case CTTimeseriesGroupsResponseMetaNormalizationPercentage, CTTimeseriesGroupsResponseMetaNormalizationMin0Max, CTTimeseriesGroupsResponseMetaNormalizationMinMax, CTTimeseriesGroupsResponseMetaNormalizationRawValues, CTTimeseriesGroupsResponseMetaNormalizationPercentageChange, CTTimeseriesGroupsResponseMetaNormalizationRollingAverage, CTTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage, CTTimeseriesGroupsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type CTTimeseriesGroupsResponseMetaUnit struct {
	Name  string                                 `json:"name,required"`
	Value string                                 `json:"value,required"`
	JSON  ctTimeseriesGroupsResponseMetaUnitJSON `json:"-"`
}

// ctTimeseriesGroupsResponseMetaUnitJSON contains the JSON metadata for the struct
// [CTTimeseriesGroupsResponseMetaUnit]
type ctTimeseriesGroupsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTTimeseriesGroupsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesGroupsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type CTTimeseriesGroupsResponseSerie0 struct {
	// This field can have the runtime type of [[]string].
	Certificate interface{} `json:"CERTIFICATE"`
	// This field can have the runtime type of [[]string].
	Domain interface{} `json:"domain"`
	// This field can have the runtime type of [[]string].
	Dsa interface{} `json:"DSA"`
	// This field can have the runtime type of [[]string].
	Ecdsa interface{} `json:"ECDSA"`
	// This field can have the runtime type of [[]string].
	Expired interface{} `json:"EXPIRED"`
	// This field can have the runtime type of [[]string].
	Extended interface{} `json:"extended"`
	// This field can have the runtime type of [[]string].
	Gt121d interface{} `json:"gt_121d"`
	// This field can have the runtime type of [[]string].
	Gt16dLte31d interface{} `json:"gt_16d_lte_31d"`
	// This field can have the runtime type of [[]string].
	Gt31dLte91d interface{} `json:"gt_31d_lte_91d"`
	// This field can have the runtime type of [[]string].
	Gt3dLte16d interface{} `json:"gt_3d_lte_16d"`
	// This field can have the runtime type of [[]string].
	Gt91dLte121d interface{} `json:"gt_91d_lte_121d"`
	// This field can have the runtime type of [[]string].
	Lte3d interface{} `json:"lte_3d"`
	// This field can have the runtime type of [[]string].
	Negative interface{} `json:"NEGATIVE"`
	// This field can have the runtime type of [[]string].
	Organization interface{} `json:"organization"`
	// This field can have the runtime type of [[]string].
	Positive interface{} `json:"POSITIVE"`
	// This field can have the runtime type of [[]string].
	Precertificate interface{} `json:"PRECERTIFICATE"`
	// This field can have the runtime type of [[]string].
	Rfc6962 interface{} `json:"rfc6962"`
	// This field can have the runtime type of [[]string].
	RSA interface{} `json:"RSA"`
	// This field can have the runtime type of [[]string].
	Static interface{} `json:"static"`
	// This field can have the runtime type of [[]time.Time].
	Timestamps interface{} `json:"timestamps"`
	// This field can have the runtime type of [[]string].
	Unknown interface{} `json:"unknown"`
	// This field can have the runtime type of [[]string].
	Valid interface{}                          `json:"VALID"`
	JSON  ctTimeseriesGroupsResponseSerie0JSON `json:"-"`
	union CTTimeseriesGroupsResponseSerie0Union
}

// ctTimeseriesGroupsResponseSerie0JSON contains the JSON metadata for the struct
// [CTTimeseriesGroupsResponseSerie0]
type ctTimeseriesGroupsResponseSerie0JSON struct {
	Certificate    apijson.Field
	Domain         apijson.Field
	Dsa            apijson.Field
	Ecdsa          apijson.Field
	Expired        apijson.Field
	Extended       apijson.Field
	Gt121d         apijson.Field
	Gt16dLte31d    apijson.Field
	Gt31dLte91d    apijson.Field
	Gt3dLte16d     apijson.Field
	Gt91dLte121d   apijson.Field
	Lte3d          apijson.Field
	Negative       apijson.Field
	Organization   apijson.Field
	Positive       apijson.Field
	Precertificate apijson.Field
	Rfc6962        apijson.Field
	RSA            apijson.Field
	Static         apijson.Field
	Timestamps     apijson.Field
	Unknown        apijson.Field
	Valid          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r ctTimeseriesGroupsResponseSerie0JSON) RawJSON() string {
	return r.raw
}

func (r *CTTimeseriesGroupsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	*r = CTTimeseriesGroupsResponseSerie0{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CTTimeseriesGroupsResponseSerie0Union] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CTTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55],
// [CTTimeseriesGroupsResponseSerie0Object],
// [CTTimeseriesGroupsResponseSerie0Object],
// [CTTimeseriesGroupsResponseSerie0Object],
// [CTTimeseriesGroupsResponseSerie0Object],
// [CTTimeseriesGroupsResponseSerie0Object],
// [CTTimeseriesGroupsResponseSerie0Object],
// [CTTimeseriesGroupsResponseSerie0Object].
func (r CTTimeseriesGroupsResponseSerie0) AsUnion() CTTimeseriesGroupsResponseSerie0Union {
	return r.union
}

// Union satisfied by
// [CTTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55],
// [CTTimeseriesGroupsResponseSerie0Object],
// [CTTimeseriesGroupsResponseSerie0Object],
// [CTTimeseriesGroupsResponseSerie0Object],
// [CTTimeseriesGroupsResponseSerie0Object],
// [CTTimeseriesGroupsResponseSerie0Object],
// [CTTimeseriesGroupsResponseSerie0Object] or
// [CTTimeseriesGroupsResponseSerie0Object].
type CTTimeseriesGroupsResponseSerie0Union interface {
	implementsCTTimeseriesGroupsResponseSerie0()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CTTimeseriesGroupsResponseSerie0Union)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CTTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CTTimeseriesGroupsResponseSerie0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CTTimeseriesGroupsResponseSerie0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CTTimeseriesGroupsResponseSerie0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CTTimeseriesGroupsResponseSerie0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CTTimeseriesGroupsResponseSerie0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CTTimeseriesGroupsResponseSerie0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CTTimeseriesGroupsResponseSerie0Object{}),
		},
	)
}

type CTTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55 struct {
	Timestamps  []time.Time                                                                          `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                                                  `json:"-,extras"`
	JSON        ctTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55JSON `json:"-"`
}

// ctTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55JSON
// contains the JSON metadata for the struct
// [CTTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55]
type ctTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55JSON) RawJSON() string {
	return r.raw
}

func (r CTTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55) implementsCTTimeseriesGroupsResponseSerie0() {
}

type CTTimeseriesGroupsResponseSerie0Object struct {
	Rfc6962 []string                                   `json:"rfc6962,required"`
	Static  []string                                   `json:"static,required"`
	JSON    ctTimeseriesGroupsResponseSerie0ObjectJSON `json:"-"`
}

// ctTimeseriesGroupsResponseSerie0ObjectJSON contains the JSON metadata for the
// struct [CTTimeseriesGroupsResponseSerie0Object]
type ctTimeseriesGroupsResponseSerie0ObjectJSON struct {
	Rfc6962     apijson.Field
	Static      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTTimeseriesGroupsResponseSerie0Object) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesGroupsResponseSerie0ObjectJSON) RawJSON() string {
	return r.raw
}

func (r CTTimeseriesGroupsResponseSerie0Object) implementsCTTimeseriesGroupsResponseSerie0() {}

type CTSummaryParams struct {
	// Filters results by certificate authority.
	CA param.Field[[]string] `query:"ca"`
	// Filters results by certificate authority owner.
	CAOwner param.Field[[]string] `query:"caOwner"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by certificate duration.
	Duration param.Field[[]CTSummaryParamsDuration] `query:"duration"`
	// Filters results by entry type (certificate vs. pre-certificate).
	EntryType param.Field[[]CTSummaryParamsEntryType] `query:"entryType"`
	// Filters results by expiration status (expired vs. valid).
	ExpirationStatus param.Field[[]CTSummaryParamsExpirationStatus] `query:"expirationStatus"`
	// Format in which results will be returned.
	Format param.Field[CTSummaryParamsFormat] `query:"format"`
	// Filters results based on whether the certificates are bound to specific IP
	// addresses.
	HasIPs param.Field[[]bool] `query:"hasIps"`
	// Filters results based on whether the certificates contain wildcard domains.
	HasWildcards param.Field[[]bool] `query:"hasWildcards"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by certificate log.
	Log param.Field[[]string] `query:"log"`
	// Filters results by certificate log API (RFC6962 vs. static).
	LogAPI param.Field[[]CTSummaryParamsLogAPI] `query:"logApi"`
	// Filters results by certificate log operator.
	LogOperator param.Field[[]string] `query:"logOperator"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[CTSummaryParamsNormalization] `query:"normalization"`
	// Filters results by public key algorithm.
	PublicKeyAlgorithm param.Field[[]CTSummaryParamsPublicKeyAlgorithm] `query:"publicKeyAlgorithm"`
	// Filters results by signature algorithm.
	SignatureAlgorithm param.Field[[]CTSummaryParamsSignatureAlgorithm] `query:"signatureAlgorithm"`
	// Filters results by top-level domain.
	TLD param.Field[[]string] `query:"tld"`
	// Specifies whether to filter out duplicate certificates and pre-certificates. Set
	// to true for unique entries only.
	UniqueEntries param.Field[[]CTSummaryParamsUniqueEntry] `query:"uniqueEntries"`
	// Filters results by validation level.
	ValidationLevel param.Field[[]CTSummaryParamsValidationLevel] `query:"validationLevel"`
}

// URLQuery serializes [CTSummaryParams]'s query parameters as `url.Values`.
func (r CTSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the certificate attribute by which to group the results.
type CTSummaryParamsDimension string

const (
	CTSummaryParamsDimensionCA                 CTSummaryParamsDimension = "CA"
	CTSummaryParamsDimensionCAOwner            CTSummaryParamsDimension = "CA_OWNER"
	CTSummaryParamsDimensionDuration           CTSummaryParamsDimension = "DURATION"
	CTSummaryParamsDimensionEntryType          CTSummaryParamsDimension = "ENTRY_TYPE"
	CTSummaryParamsDimensionExpirationStatus   CTSummaryParamsDimension = "EXPIRATION_STATUS"
	CTSummaryParamsDimensionHasIPs             CTSummaryParamsDimension = "HAS_IPS"
	CTSummaryParamsDimensionHasWildcards       CTSummaryParamsDimension = "HAS_WILDCARDS"
	CTSummaryParamsDimensionLog                CTSummaryParamsDimension = "LOG"
	CTSummaryParamsDimensionLogAPI             CTSummaryParamsDimension = "LOG_API"
	CTSummaryParamsDimensionLogOperator        CTSummaryParamsDimension = "LOG_OPERATOR"
	CTSummaryParamsDimensionPublicKeyAlgorithm CTSummaryParamsDimension = "PUBLIC_KEY_ALGORITHM"
	CTSummaryParamsDimensionSignatureAlgorithm CTSummaryParamsDimension = "SIGNATURE_ALGORITHM"
	CTSummaryParamsDimensionTLD                CTSummaryParamsDimension = "TLD"
	CTSummaryParamsDimensionValidationLevel    CTSummaryParamsDimension = "VALIDATION_LEVEL"
)

func (r CTSummaryParamsDimension) IsKnown() bool {
	switch r {
	case CTSummaryParamsDimensionCA, CTSummaryParamsDimensionCAOwner, CTSummaryParamsDimensionDuration, CTSummaryParamsDimensionEntryType, CTSummaryParamsDimensionExpirationStatus, CTSummaryParamsDimensionHasIPs, CTSummaryParamsDimensionHasWildcards, CTSummaryParamsDimensionLog, CTSummaryParamsDimensionLogAPI, CTSummaryParamsDimensionLogOperator, CTSummaryParamsDimensionPublicKeyAlgorithm, CTSummaryParamsDimensionSignatureAlgorithm, CTSummaryParamsDimensionTLD, CTSummaryParamsDimensionValidationLevel:
		return true
	}
	return false
}

type CTSummaryParamsDuration string

const (
	CTSummaryParamsDurationLte3D         CTSummaryParamsDuration = "LTE_3D"
	CTSummaryParamsDurationGt3DLte7D     CTSummaryParamsDuration = "GT_3D_LTE_7D"
	CTSummaryParamsDurationGt7DLte10D    CTSummaryParamsDuration = "GT_7D_LTE_10D"
	CTSummaryParamsDurationGt10DLte47D   CTSummaryParamsDuration = "GT_10D_LTE_47D"
	CTSummaryParamsDurationGt47DLte100D  CTSummaryParamsDuration = "GT_47D_LTE_100D"
	CTSummaryParamsDurationGt100DLte200D CTSummaryParamsDuration = "GT_100D_LTE_200D"
	CTSummaryParamsDurationGt200D        CTSummaryParamsDuration = "GT_200D"
)

func (r CTSummaryParamsDuration) IsKnown() bool {
	switch r {
	case CTSummaryParamsDurationLte3D, CTSummaryParamsDurationGt3DLte7D, CTSummaryParamsDurationGt7DLte10D, CTSummaryParamsDurationGt10DLte47D, CTSummaryParamsDurationGt47DLte100D, CTSummaryParamsDurationGt100DLte200D, CTSummaryParamsDurationGt200D:
		return true
	}
	return false
}

type CTSummaryParamsEntryType string

const (
	CTSummaryParamsEntryTypePrecertificate CTSummaryParamsEntryType = "PRECERTIFICATE"
	CTSummaryParamsEntryTypeCertificate    CTSummaryParamsEntryType = "CERTIFICATE"
)

func (r CTSummaryParamsEntryType) IsKnown() bool {
	switch r {
	case CTSummaryParamsEntryTypePrecertificate, CTSummaryParamsEntryTypeCertificate:
		return true
	}
	return false
}

type CTSummaryParamsExpirationStatus string

const (
	CTSummaryParamsExpirationStatusExpired CTSummaryParamsExpirationStatus = "EXPIRED"
	CTSummaryParamsExpirationStatusValid   CTSummaryParamsExpirationStatus = "VALID"
)

func (r CTSummaryParamsExpirationStatus) IsKnown() bool {
	switch r {
	case CTSummaryParamsExpirationStatusExpired, CTSummaryParamsExpirationStatusValid:
		return true
	}
	return false
}

// Format in which results will be returned.
type CTSummaryParamsFormat string

const (
	CTSummaryParamsFormatJson CTSummaryParamsFormat = "JSON"
	CTSummaryParamsFormatCsv  CTSummaryParamsFormat = "CSV"
)

func (r CTSummaryParamsFormat) IsKnown() bool {
	switch r {
	case CTSummaryParamsFormatJson, CTSummaryParamsFormatCsv:
		return true
	}
	return false
}

type CTSummaryParamsLogAPI string

const (
	CTSummaryParamsLogAPIRfc6962 CTSummaryParamsLogAPI = "RFC6962"
	CTSummaryParamsLogAPIStatic  CTSummaryParamsLogAPI = "STATIC"
)

func (r CTSummaryParamsLogAPI) IsKnown() bool {
	switch r {
	case CTSummaryParamsLogAPIRfc6962, CTSummaryParamsLogAPIStatic:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type CTSummaryParamsNormalization string

const (
	CTSummaryParamsNormalizationRawValues  CTSummaryParamsNormalization = "RAW_VALUES"
	CTSummaryParamsNormalizationPercentage CTSummaryParamsNormalization = "PERCENTAGE"
)

func (r CTSummaryParamsNormalization) IsKnown() bool {
	switch r {
	case CTSummaryParamsNormalizationRawValues, CTSummaryParamsNormalizationPercentage:
		return true
	}
	return false
}

type CTSummaryParamsPublicKeyAlgorithm string

const (
	CTSummaryParamsPublicKeyAlgorithmDsa   CTSummaryParamsPublicKeyAlgorithm = "DSA"
	CTSummaryParamsPublicKeyAlgorithmEcdsa CTSummaryParamsPublicKeyAlgorithm = "ECDSA"
	CTSummaryParamsPublicKeyAlgorithmRSA   CTSummaryParamsPublicKeyAlgorithm = "RSA"
)

func (r CTSummaryParamsPublicKeyAlgorithm) IsKnown() bool {
	switch r {
	case CTSummaryParamsPublicKeyAlgorithmDsa, CTSummaryParamsPublicKeyAlgorithmEcdsa, CTSummaryParamsPublicKeyAlgorithmRSA:
		return true
	}
	return false
}

type CTSummaryParamsSignatureAlgorithm string

const (
	CTSummaryParamsSignatureAlgorithmDsaSha1     CTSummaryParamsSignatureAlgorithm = "DSA_SHA_1"
	CTSummaryParamsSignatureAlgorithmDsaSha256   CTSummaryParamsSignatureAlgorithm = "DSA_SHA_256"
	CTSummaryParamsSignatureAlgorithmEcdsaSha1   CTSummaryParamsSignatureAlgorithm = "ECDSA_SHA_1"
	CTSummaryParamsSignatureAlgorithmEcdsaSha256 CTSummaryParamsSignatureAlgorithm = "ECDSA_SHA_256"
	CTSummaryParamsSignatureAlgorithmEcdsaSha384 CTSummaryParamsSignatureAlgorithm = "ECDSA_SHA_384"
	CTSummaryParamsSignatureAlgorithmEcdsaSha512 CTSummaryParamsSignatureAlgorithm = "ECDSA_SHA_512"
	CTSummaryParamsSignatureAlgorithmPssSha256   CTSummaryParamsSignatureAlgorithm = "PSS_SHA_256"
	CTSummaryParamsSignatureAlgorithmPssSha384   CTSummaryParamsSignatureAlgorithm = "PSS_SHA_384"
	CTSummaryParamsSignatureAlgorithmPssSha512   CTSummaryParamsSignatureAlgorithm = "PSS_SHA_512"
	CTSummaryParamsSignatureAlgorithmRSAMd2      CTSummaryParamsSignatureAlgorithm = "RSA_MD2"
	CTSummaryParamsSignatureAlgorithmRSAMd5      CTSummaryParamsSignatureAlgorithm = "RSA_MD5"
	CTSummaryParamsSignatureAlgorithmRSASha1     CTSummaryParamsSignatureAlgorithm = "RSA_SHA_1"
	CTSummaryParamsSignatureAlgorithmRSASha256   CTSummaryParamsSignatureAlgorithm = "RSA_SHA_256"
	CTSummaryParamsSignatureAlgorithmRSASha384   CTSummaryParamsSignatureAlgorithm = "RSA_SHA_384"
	CTSummaryParamsSignatureAlgorithmRSASha512   CTSummaryParamsSignatureAlgorithm = "RSA_SHA_512"
)

func (r CTSummaryParamsSignatureAlgorithm) IsKnown() bool {
	switch r {
	case CTSummaryParamsSignatureAlgorithmDsaSha1, CTSummaryParamsSignatureAlgorithmDsaSha256, CTSummaryParamsSignatureAlgorithmEcdsaSha1, CTSummaryParamsSignatureAlgorithmEcdsaSha256, CTSummaryParamsSignatureAlgorithmEcdsaSha384, CTSummaryParamsSignatureAlgorithmEcdsaSha512, CTSummaryParamsSignatureAlgorithmPssSha256, CTSummaryParamsSignatureAlgorithmPssSha384, CTSummaryParamsSignatureAlgorithmPssSha512, CTSummaryParamsSignatureAlgorithmRSAMd2, CTSummaryParamsSignatureAlgorithmRSAMd5, CTSummaryParamsSignatureAlgorithmRSASha1, CTSummaryParamsSignatureAlgorithmRSASha256, CTSummaryParamsSignatureAlgorithmRSASha384, CTSummaryParamsSignatureAlgorithmRSASha512:
		return true
	}
	return false
}

type CTSummaryParamsUniqueEntry string

const (
	CTSummaryParamsUniqueEntryTrue  CTSummaryParamsUniqueEntry = "true"
	CTSummaryParamsUniqueEntryFalse CTSummaryParamsUniqueEntry = "false"
)

func (r CTSummaryParamsUniqueEntry) IsKnown() bool {
	switch r {
	case CTSummaryParamsUniqueEntryTrue, CTSummaryParamsUniqueEntryFalse:
		return true
	}
	return false
}

type CTSummaryParamsValidationLevel string

const (
	CTSummaryParamsValidationLevelDomain       CTSummaryParamsValidationLevel = "DOMAIN"
	CTSummaryParamsValidationLevelOrganization CTSummaryParamsValidationLevel = "ORGANIZATION"
	CTSummaryParamsValidationLevelExtended     CTSummaryParamsValidationLevel = "EXTENDED"
)

func (r CTSummaryParamsValidationLevel) IsKnown() bool {
	switch r {
	case CTSummaryParamsValidationLevelDomain, CTSummaryParamsValidationLevelOrganization, CTSummaryParamsValidationLevelExtended:
		return true
	}
	return false
}

type CTSummaryResponseEnvelope struct {
	Result  CTSummaryResponse             `json:"result,required"`
	Success bool                          `json:"success,required"`
	JSON    ctSummaryResponseEnvelopeJSON `json:"-"`
}

// ctSummaryResponseEnvelopeJSON contains the JSON metadata for the struct
// [CTSummaryResponseEnvelope]
type ctSummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctSummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CTTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[CTTimeseriesParamsAggInterval] `query:"aggInterval"`
	// Filters results by certificate authority.
	CA param.Field[[]string] `query:"ca"`
	// Filters results by certificate authority owner.
	CAOwner param.Field[[]string] `query:"caOwner"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by certificate duration.
	Duration param.Field[[]CTTimeseriesParamsDuration] `query:"duration"`
	// Filters results by entry type (certificate vs. pre-certificate).
	EntryType param.Field[[]CTTimeseriesParamsEntryType] `query:"entryType"`
	// Filters results by expiration status (expired vs. valid).
	ExpirationStatus param.Field[[]CTTimeseriesParamsExpirationStatus] `query:"expirationStatus"`
	// Format in which results will be returned.
	Format param.Field[CTTimeseriesParamsFormat] `query:"format"`
	// Filters results based on whether the certificates are bound to specific IP
	// addresses.
	HasIPs param.Field[[]bool] `query:"hasIps"`
	// Filters results based on whether the certificates contain wildcard domains.
	HasWildcards param.Field[[]bool] `query:"hasWildcards"`
	// Filters results by certificate log.
	Log param.Field[[]string] `query:"log"`
	// Filters results by certificate log API (RFC6962 vs. static).
	LogAPI param.Field[[]CTTimeseriesParamsLogAPI] `query:"logApi"`
	// Filters results by certificate log operator.
	LogOperator param.Field[[]string] `query:"logOperator"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by public key algorithm.
	PublicKeyAlgorithm param.Field[[]CTTimeseriesParamsPublicKeyAlgorithm] `query:"publicKeyAlgorithm"`
	// Filters results by signature algorithm.
	SignatureAlgorithm param.Field[[]CTTimeseriesParamsSignatureAlgorithm] `query:"signatureAlgorithm"`
	// Filters results by top-level domain.
	TLD param.Field[[]string] `query:"tld"`
	// Specifies whether to filter out duplicate certificates and pre-certificates. Set
	// to true for unique entries only.
	UniqueEntries param.Field[[]CTTimeseriesParamsUniqueEntry] `query:"uniqueEntries"`
	// Filters results by validation level.
	ValidationLevel param.Field[[]CTTimeseriesParamsValidationLevel] `query:"validationLevel"`
}

// URLQuery serializes [CTTimeseriesParams]'s query parameters as `url.Values`.
func (r CTTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type CTTimeseriesParamsAggInterval string

const (
	CTTimeseriesParamsAggInterval15m CTTimeseriesParamsAggInterval = "15m"
	CTTimeseriesParamsAggInterval1h  CTTimeseriesParamsAggInterval = "1h"
	CTTimeseriesParamsAggInterval1d  CTTimeseriesParamsAggInterval = "1d"
	CTTimeseriesParamsAggInterval1w  CTTimeseriesParamsAggInterval = "1w"
)

func (r CTTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case CTTimeseriesParamsAggInterval15m, CTTimeseriesParamsAggInterval1h, CTTimeseriesParamsAggInterval1d, CTTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

type CTTimeseriesParamsDuration string

const (
	CTTimeseriesParamsDurationLte3D         CTTimeseriesParamsDuration = "LTE_3D"
	CTTimeseriesParamsDurationGt3DLte7D     CTTimeseriesParamsDuration = "GT_3D_LTE_7D"
	CTTimeseriesParamsDurationGt7DLte10D    CTTimeseriesParamsDuration = "GT_7D_LTE_10D"
	CTTimeseriesParamsDurationGt10DLte47D   CTTimeseriesParamsDuration = "GT_10D_LTE_47D"
	CTTimeseriesParamsDurationGt47DLte100D  CTTimeseriesParamsDuration = "GT_47D_LTE_100D"
	CTTimeseriesParamsDurationGt100DLte200D CTTimeseriesParamsDuration = "GT_100D_LTE_200D"
	CTTimeseriesParamsDurationGt200D        CTTimeseriesParamsDuration = "GT_200D"
)

func (r CTTimeseriesParamsDuration) IsKnown() bool {
	switch r {
	case CTTimeseriesParamsDurationLte3D, CTTimeseriesParamsDurationGt3DLte7D, CTTimeseriesParamsDurationGt7DLte10D, CTTimeseriesParamsDurationGt10DLte47D, CTTimeseriesParamsDurationGt47DLte100D, CTTimeseriesParamsDurationGt100DLte200D, CTTimeseriesParamsDurationGt200D:
		return true
	}
	return false
}

type CTTimeseriesParamsEntryType string

const (
	CTTimeseriesParamsEntryTypePrecertificate CTTimeseriesParamsEntryType = "PRECERTIFICATE"
	CTTimeseriesParamsEntryTypeCertificate    CTTimeseriesParamsEntryType = "CERTIFICATE"
)

func (r CTTimeseriesParamsEntryType) IsKnown() bool {
	switch r {
	case CTTimeseriesParamsEntryTypePrecertificate, CTTimeseriesParamsEntryTypeCertificate:
		return true
	}
	return false
}

type CTTimeseriesParamsExpirationStatus string

const (
	CTTimeseriesParamsExpirationStatusExpired CTTimeseriesParamsExpirationStatus = "EXPIRED"
	CTTimeseriesParamsExpirationStatusValid   CTTimeseriesParamsExpirationStatus = "VALID"
)

func (r CTTimeseriesParamsExpirationStatus) IsKnown() bool {
	switch r {
	case CTTimeseriesParamsExpirationStatusExpired, CTTimeseriesParamsExpirationStatusValid:
		return true
	}
	return false
}

// Format in which results will be returned.
type CTTimeseriesParamsFormat string

const (
	CTTimeseriesParamsFormatJson CTTimeseriesParamsFormat = "JSON"
	CTTimeseriesParamsFormatCsv  CTTimeseriesParamsFormat = "CSV"
)

func (r CTTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case CTTimeseriesParamsFormatJson, CTTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type CTTimeseriesParamsLogAPI string

const (
	CTTimeseriesParamsLogAPIRfc6962 CTTimeseriesParamsLogAPI = "RFC6962"
	CTTimeseriesParamsLogAPIStatic  CTTimeseriesParamsLogAPI = "STATIC"
)

func (r CTTimeseriesParamsLogAPI) IsKnown() bool {
	switch r {
	case CTTimeseriesParamsLogAPIRfc6962, CTTimeseriesParamsLogAPIStatic:
		return true
	}
	return false
}

type CTTimeseriesParamsPublicKeyAlgorithm string

const (
	CTTimeseriesParamsPublicKeyAlgorithmDsa   CTTimeseriesParamsPublicKeyAlgorithm = "DSA"
	CTTimeseriesParamsPublicKeyAlgorithmEcdsa CTTimeseriesParamsPublicKeyAlgorithm = "ECDSA"
	CTTimeseriesParamsPublicKeyAlgorithmRSA   CTTimeseriesParamsPublicKeyAlgorithm = "RSA"
)

func (r CTTimeseriesParamsPublicKeyAlgorithm) IsKnown() bool {
	switch r {
	case CTTimeseriesParamsPublicKeyAlgorithmDsa, CTTimeseriesParamsPublicKeyAlgorithmEcdsa, CTTimeseriesParamsPublicKeyAlgorithmRSA:
		return true
	}
	return false
}

type CTTimeseriesParamsSignatureAlgorithm string

const (
	CTTimeseriesParamsSignatureAlgorithmDsaSha1     CTTimeseriesParamsSignatureAlgorithm = "DSA_SHA_1"
	CTTimeseriesParamsSignatureAlgorithmDsaSha256   CTTimeseriesParamsSignatureAlgorithm = "DSA_SHA_256"
	CTTimeseriesParamsSignatureAlgorithmEcdsaSha1   CTTimeseriesParamsSignatureAlgorithm = "ECDSA_SHA_1"
	CTTimeseriesParamsSignatureAlgorithmEcdsaSha256 CTTimeseriesParamsSignatureAlgorithm = "ECDSA_SHA_256"
	CTTimeseriesParamsSignatureAlgorithmEcdsaSha384 CTTimeseriesParamsSignatureAlgorithm = "ECDSA_SHA_384"
	CTTimeseriesParamsSignatureAlgorithmEcdsaSha512 CTTimeseriesParamsSignatureAlgorithm = "ECDSA_SHA_512"
	CTTimeseriesParamsSignatureAlgorithmPssSha256   CTTimeseriesParamsSignatureAlgorithm = "PSS_SHA_256"
	CTTimeseriesParamsSignatureAlgorithmPssSha384   CTTimeseriesParamsSignatureAlgorithm = "PSS_SHA_384"
	CTTimeseriesParamsSignatureAlgorithmPssSha512   CTTimeseriesParamsSignatureAlgorithm = "PSS_SHA_512"
	CTTimeseriesParamsSignatureAlgorithmRSAMd2      CTTimeseriesParamsSignatureAlgorithm = "RSA_MD2"
	CTTimeseriesParamsSignatureAlgorithmRSAMd5      CTTimeseriesParamsSignatureAlgorithm = "RSA_MD5"
	CTTimeseriesParamsSignatureAlgorithmRSASha1     CTTimeseriesParamsSignatureAlgorithm = "RSA_SHA_1"
	CTTimeseriesParamsSignatureAlgorithmRSASha256   CTTimeseriesParamsSignatureAlgorithm = "RSA_SHA_256"
	CTTimeseriesParamsSignatureAlgorithmRSASha384   CTTimeseriesParamsSignatureAlgorithm = "RSA_SHA_384"
	CTTimeseriesParamsSignatureAlgorithmRSASha512   CTTimeseriesParamsSignatureAlgorithm = "RSA_SHA_512"
)

func (r CTTimeseriesParamsSignatureAlgorithm) IsKnown() bool {
	switch r {
	case CTTimeseriesParamsSignatureAlgorithmDsaSha1, CTTimeseriesParamsSignatureAlgorithmDsaSha256, CTTimeseriesParamsSignatureAlgorithmEcdsaSha1, CTTimeseriesParamsSignatureAlgorithmEcdsaSha256, CTTimeseriesParamsSignatureAlgorithmEcdsaSha384, CTTimeseriesParamsSignatureAlgorithmEcdsaSha512, CTTimeseriesParamsSignatureAlgorithmPssSha256, CTTimeseriesParamsSignatureAlgorithmPssSha384, CTTimeseriesParamsSignatureAlgorithmPssSha512, CTTimeseriesParamsSignatureAlgorithmRSAMd2, CTTimeseriesParamsSignatureAlgorithmRSAMd5, CTTimeseriesParamsSignatureAlgorithmRSASha1, CTTimeseriesParamsSignatureAlgorithmRSASha256, CTTimeseriesParamsSignatureAlgorithmRSASha384, CTTimeseriesParamsSignatureAlgorithmRSASha512:
		return true
	}
	return false
}

type CTTimeseriesParamsUniqueEntry string

const (
	CTTimeseriesParamsUniqueEntryTrue  CTTimeseriesParamsUniqueEntry = "true"
	CTTimeseriesParamsUniqueEntryFalse CTTimeseriesParamsUniqueEntry = "false"
)

func (r CTTimeseriesParamsUniqueEntry) IsKnown() bool {
	switch r {
	case CTTimeseriesParamsUniqueEntryTrue, CTTimeseriesParamsUniqueEntryFalse:
		return true
	}
	return false
}

type CTTimeseriesParamsValidationLevel string

const (
	CTTimeseriesParamsValidationLevelDomain       CTTimeseriesParamsValidationLevel = "DOMAIN"
	CTTimeseriesParamsValidationLevelOrganization CTTimeseriesParamsValidationLevel = "ORGANIZATION"
	CTTimeseriesParamsValidationLevelExtended     CTTimeseriesParamsValidationLevel = "EXTENDED"
)

func (r CTTimeseriesParamsValidationLevel) IsKnown() bool {
	switch r {
	case CTTimeseriesParamsValidationLevelDomain, CTTimeseriesParamsValidationLevelOrganization, CTTimeseriesParamsValidationLevelExtended:
		return true
	}
	return false
}

type CTTimeseriesResponseEnvelope struct {
	Result  CTTimeseriesResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    ctTimeseriesResponseEnvelopeJSON `json:"-"`
}

// ctTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [CTTimeseriesResponseEnvelope]
type ctTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CTTimeseriesGroupsParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[CTTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
	// Filters results by certificate authority.
	CA param.Field[[]string] `query:"ca"`
	// Filters results by certificate authority owner.
	CAOwner param.Field[[]string] `query:"caOwner"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by certificate duration.
	Duration param.Field[[]CTTimeseriesGroupsParamsDuration] `query:"duration"`
	// Filters results by entry type (certificate vs. pre-certificate).
	EntryType param.Field[[]CTTimeseriesGroupsParamsEntryType] `query:"entryType"`
	// Filters results by expiration status (expired vs. valid).
	ExpirationStatus param.Field[[]CTTimeseriesGroupsParamsExpirationStatus] `query:"expirationStatus"`
	// Format in which results will be returned.
	Format param.Field[CTTimeseriesGroupsParamsFormat] `query:"format"`
	// Filters results based on whether the certificates are bound to specific IP
	// addresses.
	HasIPs param.Field[[]bool] `query:"hasIps"`
	// Filters results based on whether the certificates contain wildcard domains.
	HasWildcards param.Field[[]bool] `query:"hasWildcards"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by certificate log.
	Log param.Field[[]string] `query:"log"`
	// Filters results by certificate log API (RFC6962 vs. static).
	LogAPI param.Field[[]CTTimeseriesGroupsParamsLogAPI] `query:"logApi"`
	// Filters results by certificate log operator.
	LogOperator param.Field[[]string] `query:"logOperator"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[CTTimeseriesGroupsParamsNormalization] `query:"normalization"`
	// Filters results by public key algorithm.
	PublicKeyAlgorithm param.Field[[]CTTimeseriesGroupsParamsPublicKeyAlgorithm] `query:"publicKeyAlgorithm"`
	// Filters results by signature algorithm.
	SignatureAlgorithm param.Field[[]CTTimeseriesGroupsParamsSignatureAlgorithm] `query:"signatureAlgorithm"`
	// Filters results by top-level domain.
	TLD param.Field[[]string] `query:"tld"`
	// Specifies whether to filter out duplicate certificates and pre-certificates. Set
	// to true for unique entries only.
	UniqueEntries param.Field[[]CTTimeseriesGroupsParamsUniqueEntry] `query:"uniqueEntries"`
	// Filters results by validation level.
	ValidationLevel param.Field[[]CTTimeseriesGroupsParamsValidationLevel] `query:"validationLevel"`
}

// URLQuery serializes [CTTimeseriesGroupsParams]'s query parameters as
// `url.Values`.
func (r CTTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the certificate attribute by which to group the results.
type CTTimeseriesGroupsParamsDimension string

const (
	CTTimeseriesGroupsParamsDimensionCA                 CTTimeseriesGroupsParamsDimension = "CA"
	CTTimeseriesGroupsParamsDimensionCAOwner            CTTimeseriesGroupsParamsDimension = "CA_OWNER"
	CTTimeseriesGroupsParamsDimensionDuration           CTTimeseriesGroupsParamsDimension = "DURATION"
	CTTimeseriesGroupsParamsDimensionEntryType          CTTimeseriesGroupsParamsDimension = "ENTRY_TYPE"
	CTTimeseriesGroupsParamsDimensionExpirationStatus   CTTimeseriesGroupsParamsDimension = "EXPIRATION_STATUS"
	CTTimeseriesGroupsParamsDimensionHasIPs             CTTimeseriesGroupsParamsDimension = "HAS_IPS"
	CTTimeseriesGroupsParamsDimensionHasWildcards       CTTimeseriesGroupsParamsDimension = "HAS_WILDCARDS"
	CTTimeseriesGroupsParamsDimensionLog                CTTimeseriesGroupsParamsDimension = "LOG"
	CTTimeseriesGroupsParamsDimensionLogAPI             CTTimeseriesGroupsParamsDimension = "LOG_API"
	CTTimeseriesGroupsParamsDimensionLogOperator        CTTimeseriesGroupsParamsDimension = "LOG_OPERATOR"
	CTTimeseriesGroupsParamsDimensionPublicKeyAlgorithm CTTimeseriesGroupsParamsDimension = "PUBLIC_KEY_ALGORITHM"
	CTTimeseriesGroupsParamsDimensionSignatureAlgorithm CTTimeseriesGroupsParamsDimension = "SIGNATURE_ALGORITHM"
	CTTimeseriesGroupsParamsDimensionTLD                CTTimeseriesGroupsParamsDimension = "TLD"
	CTTimeseriesGroupsParamsDimensionValidationLevel    CTTimeseriesGroupsParamsDimension = "VALIDATION_LEVEL"
)

func (r CTTimeseriesGroupsParamsDimension) IsKnown() bool {
	switch r {
	case CTTimeseriesGroupsParamsDimensionCA, CTTimeseriesGroupsParamsDimensionCAOwner, CTTimeseriesGroupsParamsDimensionDuration, CTTimeseriesGroupsParamsDimensionEntryType, CTTimeseriesGroupsParamsDimensionExpirationStatus, CTTimeseriesGroupsParamsDimensionHasIPs, CTTimeseriesGroupsParamsDimensionHasWildcards, CTTimeseriesGroupsParamsDimensionLog, CTTimeseriesGroupsParamsDimensionLogAPI, CTTimeseriesGroupsParamsDimensionLogOperator, CTTimeseriesGroupsParamsDimensionPublicKeyAlgorithm, CTTimeseriesGroupsParamsDimensionSignatureAlgorithm, CTTimeseriesGroupsParamsDimensionTLD, CTTimeseriesGroupsParamsDimensionValidationLevel:
		return true
	}
	return false
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type CTTimeseriesGroupsParamsAggInterval string

const (
	CTTimeseriesGroupsParamsAggInterval15m CTTimeseriesGroupsParamsAggInterval = "15m"
	CTTimeseriesGroupsParamsAggInterval1h  CTTimeseriesGroupsParamsAggInterval = "1h"
	CTTimeseriesGroupsParamsAggInterval1d  CTTimeseriesGroupsParamsAggInterval = "1d"
	CTTimeseriesGroupsParamsAggInterval1w  CTTimeseriesGroupsParamsAggInterval = "1w"
)

func (r CTTimeseriesGroupsParamsAggInterval) IsKnown() bool {
	switch r {
	case CTTimeseriesGroupsParamsAggInterval15m, CTTimeseriesGroupsParamsAggInterval1h, CTTimeseriesGroupsParamsAggInterval1d, CTTimeseriesGroupsParamsAggInterval1w:
		return true
	}
	return false
}

type CTTimeseriesGroupsParamsDuration string

const (
	CTTimeseriesGroupsParamsDurationLte3D         CTTimeseriesGroupsParamsDuration = "LTE_3D"
	CTTimeseriesGroupsParamsDurationGt3DLte7D     CTTimeseriesGroupsParamsDuration = "GT_3D_LTE_7D"
	CTTimeseriesGroupsParamsDurationGt7DLte10D    CTTimeseriesGroupsParamsDuration = "GT_7D_LTE_10D"
	CTTimeseriesGroupsParamsDurationGt10DLte47D   CTTimeseriesGroupsParamsDuration = "GT_10D_LTE_47D"
	CTTimeseriesGroupsParamsDurationGt47DLte100D  CTTimeseriesGroupsParamsDuration = "GT_47D_LTE_100D"
	CTTimeseriesGroupsParamsDurationGt100DLte200D CTTimeseriesGroupsParamsDuration = "GT_100D_LTE_200D"
	CTTimeseriesGroupsParamsDurationGt200D        CTTimeseriesGroupsParamsDuration = "GT_200D"
)

func (r CTTimeseriesGroupsParamsDuration) IsKnown() bool {
	switch r {
	case CTTimeseriesGroupsParamsDurationLte3D, CTTimeseriesGroupsParamsDurationGt3DLte7D, CTTimeseriesGroupsParamsDurationGt7DLte10D, CTTimeseriesGroupsParamsDurationGt10DLte47D, CTTimeseriesGroupsParamsDurationGt47DLte100D, CTTimeseriesGroupsParamsDurationGt100DLte200D, CTTimeseriesGroupsParamsDurationGt200D:
		return true
	}
	return false
}

type CTTimeseriesGroupsParamsEntryType string

const (
	CTTimeseriesGroupsParamsEntryTypePrecertificate CTTimeseriesGroupsParamsEntryType = "PRECERTIFICATE"
	CTTimeseriesGroupsParamsEntryTypeCertificate    CTTimeseriesGroupsParamsEntryType = "CERTIFICATE"
)

func (r CTTimeseriesGroupsParamsEntryType) IsKnown() bool {
	switch r {
	case CTTimeseriesGroupsParamsEntryTypePrecertificate, CTTimeseriesGroupsParamsEntryTypeCertificate:
		return true
	}
	return false
}

type CTTimeseriesGroupsParamsExpirationStatus string

const (
	CTTimeseriesGroupsParamsExpirationStatusExpired CTTimeseriesGroupsParamsExpirationStatus = "EXPIRED"
	CTTimeseriesGroupsParamsExpirationStatusValid   CTTimeseriesGroupsParamsExpirationStatus = "VALID"
)

func (r CTTimeseriesGroupsParamsExpirationStatus) IsKnown() bool {
	switch r {
	case CTTimeseriesGroupsParamsExpirationStatusExpired, CTTimeseriesGroupsParamsExpirationStatusValid:
		return true
	}
	return false
}

// Format in which results will be returned.
type CTTimeseriesGroupsParamsFormat string

const (
	CTTimeseriesGroupsParamsFormatJson CTTimeseriesGroupsParamsFormat = "JSON"
	CTTimeseriesGroupsParamsFormatCsv  CTTimeseriesGroupsParamsFormat = "CSV"
)

func (r CTTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case CTTimeseriesGroupsParamsFormatJson, CTTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}

type CTTimeseriesGroupsParamsLogAPI string

const (
	CTTimeseriesGroupsParamsLogAPIRfc6962 CTTimeseriesGroupsParamsLogAPI = "RFC6962"
	CTTimeseriesGroupsParamsLogAPIStatic  CTTimeseriesGroupsParamsLogAPI = "STATIC"
)

func (r CTTimeseriesGroupsParamsLogAPI) IsKnown() bool {
	switch r {
	case CTTimeseriesGroupsParamsLogAPIRfc6962, CTTimeseriesGroupsParamsLogAPIStatic:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type CTTimeseriesGroupsParamsNormalization string

const (
	CTTimeseriesGroupsParamsNormalizationRawValues  CTTimeseriesGroupsParamsNormalization = "RAW_VALUES"
	CTTimeseriesGroupsParamsNormalizationPercentage CTTimeseriesGroupsParamsNormalization = "PERCENTAGE"
)

func (r CTTimeseriesGroupsParamsNormalization) IsKnown() bool {
	switch r {
	case CTTimeseriesGroupsParamsNormalizationRawValues, CTTimeseriesGroupsParamsNormalizationPercentage:
		return true
	}
	return false
}

type CTTimeseriesGroupsParamsPublicKeyAlgorithm string

const (
	CTTimeseriesGroupsParamsPublicKeyAlgorithmDsa   CTTimeseriesGroupsParamsPublicKeyAlgorithm = "DSA"
	CTTimeseriesGroupsParamsPublicKeyAlgorithmEcdsa CTTimeseriesGroupsParamsPublicKeyAlgorithm = "ECDSA"
	CTTimeseriesGroupsParamsPublicKeyAlgorithmRSA   CTTimeseriesGroupsParamsPublicKeyAlgorithm = "RSA"
)

func (r CTTimeseriesGroupsParamsPublicKeyAlgorithm) IsKnown() bool {
	switch r {
	case CTTimeseriesGroupsParamsPublicKeyAlgorithmDsa, CTTimeseriesGroupsParamsPublicKeyAlgorithmEcdsa, CTTimeseriesGroupsParamsPublicKeyAlgorithmRSA:
		return true
	}
	return false
}

type CTTimeseriesGroupsParamsSignatureAlgorithm string

const (
	CTTimeseriesGroupsParamsSignatureAlgorithmDsaSha1     CTTimeseriesGroupsParamsSignatureAlgorithm = "DSA_SHA_1"
	CTTimeseriesGroupsParamsSignatureAlgorithmDsaSha256   CTTimeseriesGroupsParamsSignatureAlgorithm = "DSA_SHA_256"
	CTTimeseriesGroupsParamsSignatureAlgorithmEcdsaSha1   CTTimeseriesGroupsParamsSignatureAlgorithm = "ECDSA_SHA_1"
	CTTimeseriesGroupsParamsSignatureAlgorithmEcdsaSha256 CTTimeseriesGroupsParamsSignatureAlgorithm = "ECDSA_SHA_256"
	CTTimeseriesGroupsParamsSignatureAlgorithmEcdsaSha384 CTTimeseriesGroupsParamsSignatureAlgorithm = "ECDSA_SHA_384"
	CTTimeseriesGroupsParamsSignatureAlgorithmEcdsaSha512 CTTimeseriesGroupsParamsSignatureAlgorithm = "ECDSA_SHA_512"
	CTTimeseriesGroupsParamsSignatureAlgorithmPssSha256   CTTimeseriesGroupsParamsSignatureAlgorithm = "PSS_SHA_256"
	CTTimeseriesGroupsParamsSignatureAlgorithmPssSha384   CTTimeseriesGroupsParamsSignatureAlgorithm = "PSS_SHA_384"
	CTTimeseriesGroupsParamsSignatureAlgorithmPssSha512   CTTimeseriesGroupsParamsSignatureAlgorithm = "PSS_SHA_512"
	CTTimeseriesGroupsParamsSignatureAlgorithmRSAMd2      CTTimeseriesGroupsParamsSignatureAlgorithm = "RSA_MD2"
	CTTimeseriesGroupsParamsSignatureAlgorithmRSAMd5      CTTimeseriesGroupsParamsSignatureAlgorithm = "RSA_MD5"
	CTTimeseriesGroupsParamsSignatureAlgorithmRSASha1     CTTimeseriesGroupsParamsSignatureAlgorithm = "RSA_SHA_1"
	CTTimeseriesGroupsParamsSignatureAlgorithmRSASha256   CTTimeseriesGroupsParamsSignatureAlgorithm = "RSA_SHA_256"
	CTTimeseriesGroupsParamsSignatureAlgorithmRSASha384   CTTimeseriesGroupsParamsSignatureAlgorithm = "RSA_SHA_384"
	CTTimeseriesGroupsParamsSignatureAlgorithmRSASha512   CTTimeseriesGroupsParamsSignatureAlgorithm = "RSA_SHA_512"
)

func (r CTTimeseriesGroupsParamsSignatureAlgorithm) IsKnown() bool {
	switch r {
	case CTTimeseriesGroupsParamsSignatureAlgorithmDsaSha1, CTTimeseriesGroupsParamsSignatureAlgorithmDsaSha256, CTTimeseriesGroupsParamsSignatureAlgorithmEcdsaSha1, CTTimeseriesGroupsParamsSignatureAlgorithmEcdsaSha256, CTTimeseriesGroupsParamsSignatureAlgorithmEcdsaSha384, CTTimeseriesGroupsParamsSignatureAlgorithmEcdsaSha512, CTTimeseriesGroupsParamsSignatureAlgorithmPssSha256, CTTimeseriesGroupsParamsSignatureAlgorithmPssSha384, CTTimeseriesGroupsParamsSignatureAlgorithmPssSha512, CTTimeseriesGroupsParamsSignatureAlgorithmRSAMd2, CTTimeseriesGroupsParamsSignatureAlgorithmRSAMd5, CTTimeseriesGroupsParamsSignatureAlgorithmRSASha1, CTTimeseriesGroupsParamsSignatureAlgorithmRSASha256, CTTimeseriesGroupsParamsSignatureAlgorithmRSASha384, CTTimeseriesGroupsParamsSignatureAlgorithmRSASha512:
		return true
	}
	return false
}

type CTTimeseriesGroupsParamsUniqueEntry string

const (
	CTTimeseriesGroupsParamsUniqueEntryTrue  CTTimeseriesGroupsParamsUniqueEntry = "true"
	CTTimeseriesGroupsParamsUniqueEntryFalse CTTimeseriesGroupsParamsUniqueEntry = "false"
)

func (r CTTimeseriesGroupsParamsUniqueEntry) IsKnown() bool {
	switch r {
	case CTTimeseriesGroupsParamsUniqueEntryTrue, CTTimeseriesGroupsParamsUniqueEntryFalse:
		return true
	}
	return false
}

type CTTimeseriesGroupsParamsValidationLevel string

const (
	CTTimeseriesGroupsParamsValidationLevelDomain       CTTimeseriesGroupsParamsValidationLevel = "DOMAIN"
	CTTimeseriesGroupsParamsValidationLevelOrganization CTTimeseriesGroupsParamsValidationLevel = "ORGANIZATION"
	CTTimeseriesGroupsParamsValidationLevelExtended     CTTimeseriesGroupsParamsValidationLevel = "EXTENDED"
)

func (r CTTimeseriesGroupsParamsValidationLevel) IsKnown() bool {
	switch r {
	case CTTimeseriesGroupsParamsValidationLevelDomain, CTTimeseriesGroupsParamsValidationLevelOrganization, CTTimeseriesGroupsParamsValidationLevelExtended:
		return true
	}
	return false
}

type CTTimeseriesGroupsResponseEnvelope struct {
	Result  CTTimeseriesGroupsResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    ctTimeseriesGroupsResponseEnvelopeJSON `json:"-"`
}

// ctTimeseriesGroupsResponseEnvelopeJSON contains the JSON metadata for the struct
// [CTTimeseriesGroupsResponseEnvelope]
type ctTimeseriesGroupsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTTimeseriesGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesGroupsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
