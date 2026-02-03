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

// HTTPService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPService] method instead.
type HTTPService struct {
	Options          []option.RequestOption
	Locations        *HTTPLocationService
	Ases             *HTTPAseService
	Summary          *HTTPSummaryService
	TimeseriesGroups *HTTPTimeseriesGroupService
	Top              *HTTPTopService
}

// NewHTTPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHTTPService(opts ...option.RequestOption) (r *HTTPService) {
	r = &HTTPService{}
	r.Options = opts
	r.Locations = NewHTTPLocationService(opts...)
	r.Ases = NewHTTPAseService(opts...)
	r.Summary = NewHTTPSummaryService(opts...)
	r.TimeseriesGroups = NewHTTPTimeseriesGroupService(opts...)
	r.Top = NewHTTPTopService(opts...)
	return
}

// Retrieves the distribution of HTTP requests by the specified dimension.
func (r *HTTPService) SummaryV2(ctx context.Context, dimension HTTPSummaryV2ParamsDimension, query HTTPSummaryV2Params, opts ...option.RequestOption) (res *HTTPSummaryV2Response, err error) {
	var env HTTPSummaryV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/http/summary/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the HTTP requests over time.
func (r *HTTPService) Timeseries(ctx context.Context, query HTTPTimeseriesParams, opts ...option.RequestOption) (res *HTTPTimeseriesResponse, err error) {
	var env HTTPTimeseriesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/http/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of HTTP requests grouped by dimension.
func (r *HTTPService) TimeseriesGroupsV2(ctx context.Context, dimension HTTPTimeseriesGroupsV2ParamsDimension, query HTTPTimeseriesGroupsV2Params, opts ...option.RequestOption) (res *HTTPTimeseriesGroupsV2Response, err error) {
	var env HTTPTimeseriesGroupsV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/http/timeseries_groups/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPSummaryV2Response struct {
	// Metadata for the results.
	Meta     HTTPSummaryV2ResponseMeta `json:"meta,required"`
	Summary0 map[string]string         `json:"summary_0,required"`
	JSON     httpSummaryV2ResponseJSON `json:"-"`
}

// httpSummaryV2ResponseJSON contains the JSON metadata for the struct
// [HTTPSummaryV2Response]
type httpSummaryV2ResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPSummaryV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpSummaryV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPSummaryV2ResponseMeta struct {
	ConfidenceInfo HTTPSummaryV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []HTTPSummaryV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPSummaryV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPSummaryV2ResponseMetaUnit `json:"units,required"`
	JSON  httpSummaryV2ResponseMetaJSON   `json:"-"`
}

// httpSummaryV2ResponseMetaJSON contains the JSON metadata for the struct
// [HTTPSummaryV2ResponseMeta]
type httpSummaryV2ResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPSummaryV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpSummaryV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPSummaryV2ResponseMetaConfidenceInfo struct {
	Annotations []HTTPSummaryV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                       `json:"level,required"`
	JSON  httpSummaryV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpSummaryV2ResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [HTTPSummaryV2ResponseMetaConfidenceInfo]
type httpSummaryV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPSummaryV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpSummaryV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPSummaryV2ResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                       `json:"description,required"`
	EndDate     time.Time                                                    `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                             `json:"startDate,required" format:"date-time"`
	JSON            httpSummaryV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpSummaryV2ResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata
// for the struct [HTTPSummaryV2ResponseMetaConfidenceInfoAnnotation]
type httpSummaryV2ResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPSummaryV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpSummaryV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll                HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP                HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots               HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT                 HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS                HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos                HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw                 HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI                HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet                HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType string

const (
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent             HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage            HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline          HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline, HTTPSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type HTTPSummaryV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                              `json:"startTime,required" format:"date-time"`
	JSON      httpSummaryV2ResponseMetaDateRangeJSON `json:"-"`
}

// httpSummaryV2ResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [HTTPSummaryV2ResponseMetaDateRange]
type httpSummaryV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPSummaryV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpSummaryV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPSummaryV2ResponseMetaNormalization string

const (
	HTTPSummaryV2ResponseMetaNormalizationPercentage           HTTPSummaryV2ResponseMetaNormalization = "PERCENTAGE"
	HTTPSummaryV2ResponseMetaNormalizationMin0Max              HTTPSummaryV2ResponseMetaNormalization = "MIN0_MAX"
	HTTPSummaryV2ResponseMetaNormalizationMinMax               HTTPSummaryV2ResponseMetaNormalization = "MIN_MAX"
	HTTPSummaryV2ResponseMetaNormalizationRawValues            HTTPSummaryV2ResponseMetaNormalization = "RAW_VALUES"
	HTTPSummaryV2ResponseMetaNormalizationPercentageChange     HTTPSummaryV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPSummaryV2ResponseMetaNormalizationRollingAverage       HTTPSummaryV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPSummaryV2ResponseMetaNormalizationOverlappedPercentage HTTPSummaryV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPSummaryV2ResponseMetaNormalizationRatio                HTTPSummaryV2ResponseMetaNormalization = "RATIO"
)

func (r HTTPSummaryV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPSummaryV2ResponseMetaNormalizationPercentage, HTTPSummaryV2ResponseMetaNormalizationMin0Max, HTTPSummaryV2ResponseMetaNormalizationMinMax, HTTPSummaryV2ResponseMetaNormalizationRawValues, HTTPSummaryV2ResponseMetaNormalizationPercentageChange, HTTPSummaryV2ResponseMetaNormalizationRollingAverage, HTTPSummaryV2ResponseMetaNormalizationOverlappedPercentage, HTTPSummaryV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPSummaryV2ResponseMetaUnit struct {
	Name  string                            `json:"name,required"`
	Value string                            `json:"value,required"`
	JSON  httpSummaryV2ResponseMetaUnitJSON `json:"-"`
}

// httpSummaryV2ResponseMetaUnitJSON contains the JSON metadata for the struct
// [HTTPSummaryV2ResponseMetaUnit]
type httpSummaryV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPSummaryV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpSummaryV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesResponse struct {
	// Metadata for the results.
	Meta        HTTPTimeseriesResponseMeta        `json:"meta,required"`
	ExtraFields map[string]HTTPTimeseriesResponse `json:"-,extras"`
	JSON        httpTimeseriesResponseJSON        `json:"-"`
}

// httpTimeseriesResponseJSON contains the JSON metadata for the struct
// [HTTPTimeseriesResponse]
type httpTimeseriesResponseJSON struct {
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPTimeseriesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    HTTPTimeseriesResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo HTTPTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []HTTPTimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPTimeseriesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPTimeseriesResponseMetaUnit `json:"units,required"`
	JSON  httpTimeseriesResponseMetaJSON   `json:"-"`
}

// httpTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [HTTPTimeseriesResponseMeta]
type httpTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesResponseMetaAggInterval string

const (
	HTTPTimeseriesResponseMetaAggIntervalFifteenMinutes HTTPTimeseriesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	HTTPTimeseriesResponseMetaAggIntervalOneHour        HTTPTimeseriesResponseMetaAggInterval = "ONE_HOUR"
	HTTPTimeseriesResponseMetaAggIntervalOneDay         HTTPTimeseriesResponseMetaAggInterval = "ONE_DAY"
	HTTPTimeseriesResponseMetaAggIntervalOneWeek        HTTPTimeseriesResponseMetaAggInterval = "ONE_WEEK"
	HTTPTimeseriesResponseMetaAggIntervalOneMonth       HTTPTimeseriesResponseMetaAggInterval = "ONE_MONTH"
)

func (r HTTPTimeseriesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesResponseMetaAggIntervalFifteenMinutes, HTTPTimeseriesResponseMetaAggIntervalOneHour, HTTPTimeseriesResponseMetaAggIntervalOneDay, HTTPTimeseriesResponseMetaAggIntervalOneWeek, HTTPTimeseriesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type HTTPTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []HTTPTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                        `json:"level,required"`
	JSON  httpTimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpTimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [HTTPTimeseriesResponseMetaConfidenceInfo]
type httpTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                        `json:"description,required"`
	EndDate     time.Time                                                     `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                              `json:"startDate,required" format:"date-time"`
	JSON            httpTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [HTTPTimeseriesResponseMetaConfidenceInfoAnnotation]
type httpTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll                HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP                HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots               HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCT                 HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS                HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos                HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw                 HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI                HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet                HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCT, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventType string

const (
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent             HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage            HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline          HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline, HTTPTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type HTTPTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                               `json:"startTime,required" format:"date-time"`
	JSON      httpTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// httpTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [HTTPTimeseriesResponseMetaDateRange]
type httpTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPTimeseriesResponseMetaNormalization string

const (
	HTTPTimeseriesResponseMetaNormalizationPercentage           HTTPTimeseriesResponseMetaNormalization = "PERCENTAGE"
	HTTPTimeseriesResponseMetaNormalizationMin0Max              HTTPTimeseriesResponseMetaNormalization = "MIN0_MAX"
	HTTPTimeseriesResponseMetaNormalizationMinMax               HTTPTimeseriesResponseMetaNormalization = "MIN_MAX"
	HTTPTimeseriesResponseMetaNormalizationRawValues            HTTPTimeseriesResponseMetaNormalization = "RAW_VALUES"
	HTTPTimeseriesResponseMetaNormalizationPercentageChange     HTTPTimeseriesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPTimeseriesResponseMetaNormalizationRollingAverage       HTTPTimeseriesResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPTimeseriesResponseMetaNormalizationOverlappedPercentage HTTPTimeseriesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPTimeseriesResponseMetaNormalizationRatio                HTTPTimeseriesResponseMetaNormalization = "RATIO"
)

func (r HTTPTimeseriesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPTimeseriesResponseMetaNormalizationPercentage, HTTPTimeseriesResponseMetaNormalizationMin0Max, HTTPTimeseriesResponseMetaNormalizationMinMax, HTTPTimeseriesResponseMetaNormalizationRawValues, HTTPTimeseriesResponseMetaNormalizationPercentageChange, HTTPTimeseriesResponseMetaNormalizationRollingAverage, HTTPTimeseriesResponseMetaNormalizationOverlappedPercentage, HTTPTimeseriesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPTimeseriesResponseMetaUnit struct {
	Name  string                             `json:"name,required"`
	Value string                             `json:"value,required"`
	JSON  httpTimeseriesResponseMetaUnitJSON `json:"-"`
}

// httpTimeseriesResponseMetaUnitJSON contains the JSON metadata for the struct
// [HTTPTimeseriesResponseMetaUnit]
type httpTimeseriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupsV2Response struct {
	// Metadata for the results.
	Meta   HTTPTimeseriesGroupsV2ResponseMeta   `json:"meta,required"`
	Serie0 HTTPTimeseriesGroupsV2ResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesGroupsV2ResponseJSON   `json:"-"`
}

// httpTimeseriesGroupsV2ResponseJSON contains the JSON metadata for the struct
// [HTTPTimeseriesGroupsV2Response]
type httpTimeseriesGroupsV2ResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupsV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupsV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type HTTPTimeseriesGroupsV2ResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    HTTPTimeseriesGroupsV2ResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []HTTPTimeseriesGroupsV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization HTTPTimeseriesGroupsV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []HTTPTimeseriesGroupsV2ResponseMetaUnit `json:"units,required"`
	JSON  httpTimeseriesGroupsV2ResponseMetaJSON   `json:"-"`
}

// httpTimeseriesGroupsV2ResponseMetaJSON contains the JSON metadata for the struct
// [HTTPTimeseriesGroupsV2ResponseMeta]
type httpTimeseriesGroupsV2ResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupsV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupsV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupsV2ResponseMetaAggInterval string

const (
	HTTPTimeseriesGroupsV2ResponseMetaAggIntervalFifteenMinutes HTTPTimeseriesGroupsV2ResponseMetaAggInterval = "FIFTEEN_MINUTES"
	HTTPTimeseriesGroupsV2ResponseMetaAggIntervalOneHour        HTTPTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_HOUR"
	HTTPTimeseriesGroupsV2ResponseMetaAggIntervalOneDay         HTTPTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_DAY"
	HTTPTimeseriesGroupsV2ResponseMetaAggIntervalOneWeek        HTTPTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_WEEK"
	HTTPTimeseriesGroupsV2ResponseMetaAggIntervalOneMonth       HTTPTimeseriesGroupsV2ResponseMetaAggInterval = "ONE_MONTH"
)

func (r HTTPTimeseriesGroupsV2ResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupsV2ResponseMetaAggIntervalFifteenMinutes, HTTPTimeseriesGroupsV2ResponseMetaAggIntervalOneHour, HTTPTimeseriesGroupsV2ResponseMetaAggIntervalOneDay, HTTPTimeseriesGroupsV2ResponseMetaAggIntervalOneWeek, HTTPTimeseriesGroupsV2ResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfo struct {
	Annotations []HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                `json:"level,required"`
	JSON  httpTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfo]
type httpTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupsV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                `json:"description,required"`
	EndDate     time.Time                                                             `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                      `json:"startDate,required" format:"date-time"`
	JSON            httpTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation]
type httpTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll                HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP                HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots               HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT                 HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS                HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos                HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw                 HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI                HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet                HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType string

const (
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent             HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage            HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline          HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline, HTTPTimeseriesGroupsV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type HTTPTimeseriesGroupsV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      httpTimeseriesGroupsV2ResponseMetaDateRangeJSON `json:"-"`
}

// httpTimeseriesGroupsV2ResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [HTTPTimeseriesGroupsV2ResponseMetaDateRange]
type httpTimeseriesGroupsV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupsV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupsV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPTimeseriesGroupsV2ResponseMetaNormalization string

const (
	HTTPTimeseriesGroupsV2ResponseMetaNormalizationPercentage           HTTPTimeseriesGroupsV2ResponseMetaNormalization = "PERCENTAGE"
	HTTPTimeseriesGroupsV2ResponseMetaNormalizationMin0Max              HTTPTimeseriesGroupsV2ResponseMetaNormalization = "MIN0_MAX"
	HTTPTimeseriesGroupsV2ResponseMetaNormalizationMinMax               HTTPTimeseriesGroupsV2ResponseMetaNormalization = "MIN_MAX"
	HTTPTimeseriesGroupsV2ResponseMetaNormalizationRawValues            HTTPTimeseriesGroupsV2ResponseMetaNormalization = "RAW_VALUES"
	HTTPTimeseriesGroupsV2ResponseMetaNormalizationPercentageChange     HTTPTimeseriesGroupsV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	HTTPTimeseriesGroupsV2ResponseMetaNormalizationRollingAverage       HTTPTimeseriesGroupsV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	HTTPTimeseriesGroupsV2ResponseMetaNormalizationOverlappedPercentage HTTPTimeseriesGroupsV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	HTTPTimeseriesGroupsV2ResponseMetaNormalizationRatio                HTTPTimeseriesGroupsV2ResponseMetaNormalization = "RATIO"
)

func (r HTTPTimeseriesGroupsV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupsV2ResponseMetaNormalizationPercentage, HTTPTimeseriesGroupsV2ResponseMetaNormalizationMin0Max, HTTPTimeseriesGroupsV2ResponseMetaNormalizationMinMax, HTTPTimeseriesGroupsV2ResponseMetaNormalizationRawValues, HTTPTimeseriesGroupsV2ResponseMetaNormalizationPercentageChange, HTTPTimeseriesGroupsV2ResponseMetaNormalizationRollingAverage, HTTPTimeseriesGroupsV2ResponseMetaNormalizationOverlappedPercentage, HTTPTimeseriesGroupsV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type HTTPTimeseriesGroupsV2ResponseMetaUnit struct {
	Name  string                                     `json:"name,required"`
	Value string                                     `json:"value,required"`
	JSON  httpTimeseriesGroupsV2ResponseMetaUnitJSON `json:"-"`
}

// httpTimeseriesGroupsV2ResponseMetaUnitJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupsV2ResponseMetaUnit]
type httpTimeseriesGroupsV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupsV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupsV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupsV2ResponseSerie0 struct {
	Timestamps  []time.Time                              `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                      `json:"-,extras"`
	JSON        httpTimeseriesGroupsV2ResponseSerie0JSON `json:"-"`
}

// httpTimeseriesGroupsV2ResponseSerie0JSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupsV2ResponseSerie0]
type httpTimeseriesGroupsV2ResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupsV2ResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupsV2ResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPSummaryV2Params struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPSummaryV2ParamsBotClass] `query:"botClass"`
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
	// Filters results by device type.
	DeviceType param.Field[[]HTTPSummaryV2ParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPSummaryV2ParamsFormat] `query:"format"`
	// Filters results by Geolocation. Specify a comma-separated list of GeoNames IDs.
	// Prefix with `-` to exclude geoIds from results. For example, `-2267056,360689`
	// excludes results from the 2267056 (Lisbon), but includes results from 5128638
	// (New York).
	GeoID param.Field[[]string] `query:"geoId"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPSummaryV2ParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPSummaryV2ParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPSummaryV2ParamsIPVersion] `query:"ipVersion"`
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
	// Filters results by operating system.
	OS param.Field[[]HTTPSummaryV2ParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPSummaryV2ParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPSummaryV2Params]'s query parameters as `url.Values`.
func (r HTTPSummaryV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the HTTP attribute by which to group the results.
type HTTPSummaryV2ParamsDimension string

const (
	HTTPSummaryV2ParamsDimensionAdm1          HTTPSummaryV2ParamsDimension = "ADM1"
	HTTPSummaryV2ParamsDimensionBotClass      HTTPSummaryV2ParamsDimension = "BOT_CLASS"
	HTTPSummaryV2ParamsDimensionBrowser       HTTPSummaryV2ParamsDimension = "BROWSER"
	HTTPSummaryV2ParamsDimensionBrowserFamily HTTPSummaryV2ParamsDimension = "BROWSER_FAMILY"
	HTTPSummaryV2ParamsDimensionDeviceType    HTTPSummaryV2ParamsDimension = "DEVICE_TYPE"
	HTTPSummaryV2ParamsDimensionHTTPProtocol  HTTPSummaryV2ParamsDimension = "HTTP_PROTOCOL"
	HTTPSummaryV2ParamsDimensionHTTPVersion   HTTPSummaryV2ParamsDimension = "HTTP_VERSION"
	HTTPSummaryV2ParamsDimensionIPVersion     HTTPSummaryV2ParamsDimension = "IP_VERSION"
	HTTPSummaryV2ParamsDimensionOS            HTTPSummaryV2ParamsDimension = "OS"
	HTTPSummaryV2ParamsDimensionPostQuantum   HTTPSummaryV2ParamsDimension = "POST_QUANTUM"
	HTTPSummaryV2ParamsDimensionTLSVersion    HTTPSummaryV2ParamsDimension = "TLS_VERSION"
)

func (r HTTPSummaryV2ParamsDimension) IsKnown() bool {
	switch r {
	case HTTPSummaryV2ParamsDimensionAdm1, HTTPSummaryV2ParamsDimensionBotClass, HTTPSummaryV2ParamsDimensionBrowser, HTTPSummaryV2ParamsDimensionBrowserFamily, HTTPSummaryV2ParamsDimensionDeviceType, HTTPSummaryV2ParamsDimensionHTTPProtocol, HTTPSummaryV2ParamsDimensionHTTPVersion, HTTPSummaryV2ParamsDimensionIPVersion, HTTPSummaryV2ParamsDimensionOS, HTTPSummaryV2ParamsDimensionPostQuantum, HTTPSummaryV2ParamsDimensionTLSVersion:
		return true
	}
	return false
}

type HTTPSummaryV2ParamsBotClass string

const (
	HTTPSummaryV2ParamsBotClassLikelyAutomated HTTPSummaryV2ParamsBotClass = "LIKELY_AUTOMATED"
	HTTPSummaryV2ParamsBotClassLikelyHuman     HTTPSummaryV2ParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPSummaryV2ParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPSummaryV2ParamsBotClassLikelyAutomated, HTTPSummaryV2ParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPSummaryV2ParamsDeviceType string

const (
	HTTPSummaryV2ParamsDeviceTypeDesktop HTTPSummaryV2ParamsDeviceType = "DESKTOP"
	HTTPSummaryV2ParamsDeviceTypeMobile  HTTPSummaryV2ParamsDeviceType = "MOBILE"
	HTTPSummaryV2ParamsDeviceTypeOther   HTTPSummaryV2ParamsDeviceType = "OTHER"
)

func (r HTTPSummaryV2ParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPSummaryV2ParamsDeviceTypeDesktop, HTTPSummaryV2ParamsDeviceTypeMobile, HTTPSummaryV2ParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPSummaryV2ParamsFormat string

const (
	HTTPSummaryV2ParamsFormatJson HTTPSummaryV2ParamsFormat = "JSON"
	HTTPSummaryV2ParamsFormatCsv  HTTPSummaryV2ParamsFormat = "CSV"
)

func (r HTTPSummaryV2ParamsFormat) IsKnown() bool {
	switch r {
	case HTTPSummaryV2ParamsFormatJson, HTTPSummaryV2ParamsFormatCsv:
		return true
	}
	return false
}

type HTTPSummaryV2ParamsHTTPProtocol string

const (
	HTTPSummaryV2ParamsHTTPProtocolHTTP  HTTPSummaryV2ParamsHTTPProtocol = "HTTP"
	HTTPSummaryV2ParamsHTTPProtocolHTTPS HTTPSummaryV2ParamsHTTPProtocol = "HTTPS"
)

func (r HTTPSummaryV2ParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPSummaryV2ParamsHTTPProtocolHTTP, HTTPSummaryV2ParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPSummaryV2ParamsHTTPVersion string

const (
	HTTPSummaryV2ParamsHTTPVersionHttPv1 HTTPSummaryV2ParamsHTTPVersion = "HTTPv1"
	HTTPSummaryV2ParamsHTTPVersionHttPv2 HTTPSummaryV2ParamsHTTPVersion = "HTTPv2"
	HTTPSummaryV2ParamsHTTPVersionHttPv3 HTTPSummaryV2ParamsHTTPVersion = "HTTPv3"
)

func (r HTTPSummaryV2ParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPSummaryV2ParamsHTTPVersionHttPv1, HTTPSummaryV2ParamsHTTPVersionHttPv2, HTTPSummaryV2ParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPSummaryV2ParamsIPVersion string

const (
	HTTPSummaryV2ParamsIPVersionIPv4 HTTPSummaryV2ParamsIPVersion = "IPv4"
	HTTPSummaryV2ParamsIPVersionIPv6 HTTPSummaryV2ParamsIPVersion = "IPv6"
)

func (r HTTPSummaryV2ParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPSummaryV2ParamsIPVersionIPv4, HTTPSummaryV2ParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPSummaryV2ParamsOS string

const (
	HTTPSummaryV2ParamsOSWindows  HTTPSummaryV2ParamsOS = "WINDOWS"
	HTTPSummaryV2ParamsOSMacosx   HTTPSummaryV2ParamsOS = "MACOSX"
	HTTPSummaryV2ParamsOSIos      HTTPSummaryV2ParamsOS = "IOS"
	HTTPSummaryV2ParamsOSAndroid  HTTPSummaryV2ParamsOS = "ANDROID"
	HTTPSummaryV2ParamsOSChromeos HTTPSummaryV2ParamsOS = "CHROMEOS"
	HTTPSummaryV2ParamsOSLinux    HTTPSummaryV2ParamsOS = "LINUX"
	HTTPSummaryV2ParamsOSSmartTv  HTTPSummaryV2ParamsOS = "SMART_TV"
)

func (r HTTPSummaryV2ParamsOS) IsKnown() bool {
	switch r {
	case HTTPSummaryV2ParamsOSWindows, HTTPSummaryV2ParamsOSMacosx, HTTPSummaryV2ParamsOSIos, HTTPSummaryV2ParamsOSAndroid, HTTPSummaryV2ParamsOSChromeos, HTTPSummaryV2ParamsOSLinux, HTTPSummaryV2ParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPSummaryV2ParamsTLSVersion string

const (
	HTTPSummaryV2ParamsTLSVersionTlSv1_0  HTTPSummaryV2ParamsTLSVersion = "TLSv1_0"
	HTTPSummaryV2ParamsTLSVersionTlSv1_1  HTTPSummaryV2ParamsTLSVersion = "TLSv1_1"
	HTTPSummaryV2ParamsTLSVersionTlSv1_2  HTTPSummaryV2ParamsTLSVersion = "TLSv1_2"
	HTTPSummaryV2ParamsTLSVersionTlSv1_3  HTTPSummaryV2ParamsTLSVersion = "TLSv1_3"
	HTTPSummaryV2ParamsTLSVersionTlSvQuic HTTPSummaryV2ParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPSummaryV2ParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPSummaryV2ParamsTLSVersionTlSv1_0, HTTPSummaryV2ParamsTLSVersionTlSv1_1, HTTPSummaryV2ParamsTLSVersionTlSv1_2, HTTPSummaryV2ParamsTLSVersionTlSv1_3, HTTPSummaryV2ParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPSummaryV2ResponseEnvelope struct {
	Result  HTTPSummaryV2Response             `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    httpSummaryV2ResponseEnvelopeJSON `json:"-"`
}

// httpSummaryV2ResponseEnvelopeJSON contains the JSON metadata for the struct
// [HTTPSummaryV2ResponseEnvelope]
type httpSummaryV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPSummaryV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpSummaryV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]HTTPTimeseriesParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]HTTPTimeseriesParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPTimeseriesParamsFormat] `query:"format"`
	// Filters results by Geolocation. Specify a comma-separated list of GeoNames IDs.
	// Prefix with `-` to exclude geoIds from results. For example, `-2267056,360689`
	// excludes results from the 2267056 (Lisbon), but includes results from 5128638
	// (New York).
	GeoID param.Field[[]string] `query:"geoId"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPTimeseriesParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPTimeseriesParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[HTTPTimeseriesParamsNormalization] `query:"normalization"`
	// Filters results by operating system.
	OS param.Field[[]HTTPTimeseriesParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPTimeseriesParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesParams]'s query parameters as `url.Values`.
func (r HTTPTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesParamsAggInterval string

const (
	HTTPTimeseriesParamsAggInterval15m HTTPTimeseriesParamsAggInterval = "15m"
	HTTPTimeseriesParamsAggInterval1h  HTTPTimeseriesParamsAggInterval = "1h"
	HTTPTimeseriesParamsAggInterval1d  HTTPTimeseriesParamsAggInterval = "1d"
	HTTPTimeseriesParamsAggInterval1w  HTTPTimeseriesParamsAggInterval = "1w"
)

func (r HTTPTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsAggInterval15m, HTTPTimeseriesParamsAggInterval1h, HTTPTimeseriesParamsAggInterval1d, HTTPTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesParamsBotClass string

const (
	HTTPTimeseriesParamsBotClassLikelyAutomated HTTPTimeseriesParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesParamsBotClassLikelyHuman     HTTPTimeseriesParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsBotClassLikelyAutomated, HTTPTimeseriesParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesParamsBrowserFamily string

const (
	HTTPTimeseriesParamsBrowserFamilyChrome  HTTPTimeseriesParamsBrowserFamily = "CHROME"
	HTTPTimeseriesParamsBrowserFamilyEdge    HTTPTimeseriesParamsBrowserFamily = "EDGE"
	HTTPTimeseriesParamsBrowserFamilyFirefox HTTPTimeseriesParamsBrowserFamily = "FIREFOX"
	HTTPTimeseriesParamsBrowserFamilySafari  HTTPTimeseriesParamsBrowserFamily = "SAFARI"
)

func (r HTTPTimeseriesParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsBrowserFamilyChrome, HTTPTimeseriesParamsBrowserFamilyEdge, HTTPTimeseriesParamsBrowserFamilyFirefox, HTTPTimeseriesParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPTimeseriesParamsDeviceType string

const (
	HTTPTimeseriesParamsDeviceTypeDesktop HTTPTimeseriesParamsDeviceType = "DESKTOP"
	HTTPTimeseriesParamsDeviceTypeMobile  HTTPTimeseriesParamsDeviceType = "MOBILE"
	HTTPTimeseriesParamsDeviceTypeOther   HTTPTimeseriesParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsDeviceTypeDesktop, HTTPTimeseriesParamsDeviceTypeMobile, HTTPTimeseriesParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPTimeseriesParamsFormat string

const (
	HTTPTimeseriesParamsFormatJson HTTPTimeseriesParamsFormat = "JSON"
	HTTPTimeseriesParamsFormatCsv  HTTPTimeseriesParamsFormat = "CSV"
)

func (r HTTPTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsFormatJson, HTTPTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesParamsHTTPProtocol string

const (
	HTTPTimeseriesParamsHTTPProtocolHTTP  HTTPTimeseriesParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesParamsHTTPProtocolHTTPS HTTPTimeseriesParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsHTTPProtocolHTTP, HTTPTimeseriesParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesParamsHTTPVersion string

const (
	HTTPTimeseriesParamsHTTPVersionHttPv1 HTTPTimeseriesParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesParamsHTTPVersionHttPv2 HTTPTimeseriesParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesParamsHTTPVersionHttPv3 HTTPTimeseriesParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsHTTPVersionHttPv1, HTTPTimeseriesParamsHTTPVersionHttPv2, HTTPTimeseriesParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesParamsIPVersion string

const (
	HTTPTimeseriesParamsIPVersionIPv4 HTTPTimeseriesParamsIPVersion = "IPv4"
	HTTPTimeseriesParamsIPVersionIPv6 HTTPTimeseriesParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsIPVersionIPv4, HTTPTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPTimeseriesParamsNormalization string

const (
	HTTPTimeseriesParamsNormalizationPercentageChange HTTPTimeseriesParamsNormalization = "PERCENTAGE_CHANGE"
	HTTPTimeseriesParamsNormalizationMin0Max          HTTPTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r HTTPTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsNormalizationPercentageChange, HTTPTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type HTTPTimeseriesParamsOS string

const (
	HTTPTimeseriesParamsOSWindows  HTTPTimeseriesParamsOS = "WINDOWS"
	HTTPTimeseriesParamsOSMacosx   HTTPTimeseriesParamsOS = "MACOSX"
	HTTPTimeseriesParamsOSIos      HTTPTimeseriesParamsOS = "IOS"
	HTTPTimeseriesParamsOSAndroid  HTTPTimeseriesParamsOS = "ANDROID"
	HTTPTimeseriesParamsOSChromeos HTTPTimeseriesParamsOS = "CHROMEOS"
	HTTPTimeseriesParamsOSLinux    HTTPTimeseriesParamsOS = "LINUX"
	HTTPTimeseriesParamsOSSmartTv  HTTPTimeseriesParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsOSWindows, HTTPTimeseriesParamsOSMacosx, HTTPTimeseriesParamsOSIos, HTTPTimeseriesParamsOSAndroid, HTTPTimeseriesParamsOSChromeos, HTTPTimeseriesParamsOSLinux, HTTPTimeseriesParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesParamsTLSVersion string

const (
	HTTPTimeseriesParamsTLSVersionTlSv1_0  HTTPTimeseriesParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesParamsTLSVersionTlSv1_1  HTTPTimeseriesParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesParamsTLSVersionTlSv1_2  HTTPTimeseriesParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesParamsTLSVersionTlSv1_3  HTTPTimeseriesParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesParamsTLSVersionTlSvQuic HTTPTimeseriesParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsTLSVersionTlSv1_0, HTTPTimeseriesParamsTLSVersionTlSv1_1, HTTPTimeseriesParamsTLSVersionTlSv1_2, HTTPTimeseriesParamsTLSVersionTlSv1_3, HTTPTimeseriesParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesResponseEnvelope struct {
	Result  HTTPTimeseriesResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    httpTimeseriesResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [HTTPTimeseriesResponseEnvelope]
type httpTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesGroupsV2Params struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesGroupsV2ParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesGroupsV2ParamsBotClass] `query:"botClass"`
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
	// Filters results by device type.
	DeviceType param.Field[[]HTTPTimeseriesGroupsV2ParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPTimeseriesGroupsV2ParamsFormat] `query:"format"`
	// Filters results by Geolocation. Specify a comma-separated list of GeoNames IDs.
	// Prefix with `-` to exclude geoIds from results. For example, `-2267056,360689`
	// excludes results from the 2267056 (Lisbon), but includes results from 5128638
	// (New York).
	GeoID param.Field[[]string] `query:"geoId"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPTimeseriesGroupsV2ParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPTimeseriesGroupsV2ParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPTimeseriesGroupsV2ParamsIPVersion] `query:"ipVersion"`
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
	Normalization param.Field[HTTPTimeseriesGroupsV2ParamsNormalization] `query:"normalization"`
	// Filters results by operating system.
	OS param.Field[[]HTTPTimeseriesGroupsV2ParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPTimeseriesGroupsV2ParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesGroupsV2Params]'s query parameters as
// `url.Values`.
func (r HTTPTimeseriesGroupsV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the HTTP attribute by which to group the results.
type HTTPTimeseriesGroupsV2ParamsDimension string

const (
	HTTPTimeseriesGroupsV2ParamsDimensionAdm1          HTTPTimeseriesGroupsV2ParamsDimension = "ADM1"
	HTTPTimeseriesGroupsV2ParamsDimensionBotClass      HTTPTimeseriesGroupsV2ParamsDimension = "BOT_CLASS"
	HTTPTimeseriesGroupsV2ParamsDimensionBrowser       HTTPTimeseriesGroupsV2ParamsDimension = "BROWSER"
	HTTPTimeseriesGroupsV2ParamsDimensionBrowserFamily HTTPTimeseriesGroupsV2ParamsDimension = "BROWSER_FAMILY"
	HTTPTimeseriesGroupsV2ParamsDimensionDeviceType    HTTPTimeseriesGroupsV2ParamsDimension = "DEVICE_TYPE"
	HTTPTimeseriesGroupsV2ParamsDimensionHTTPProtocol  HTTPTimeseriesGroupsV2ParamsDimension = "HTTP_PROTOCOL"
	HTTPTimeseriesGroupsV2ParamsDimensionHTTPVersion   HTTPTimeseriesGroupsV2ParamsDimension = "HTTP_VERSION"
	HTTPTimeseriesGroupsV2ParamsDimensionIPVersion     HTTPTimeseriesGroupsV2ParamsDimension = "IP_VERSION"
	HTTPTimeseriesGroupsV2ParamsDimensionOS            HTTPTimeseriesGroupsV2ParamsDimension = "OS"
	HTTPTimeseriesGroupsV2ParamsDimensionPostQuantum   HTTPTimeseriesGroupsV2ParamsDimension = "POST_QUANTUM"
	HTTPTimeseriesGroupsV2ParamsDimensionTLSVersion    HTTPTimeseriesGroupsV2ParamsDimension = "TLS_VERSION"
)

func (r HTTPTimeseriesGroupsV2ParamsDimension) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupsV2ParamsDimensionAdm1, HTTPTimeseriesGroupsV2ParamsDimensionBotClass, HTTPTimeseriesGroupsV2ParamsDimensionBrowser, HTTPTimeseriesGroupsV2ParamsDimensionBrowserFamily, HTTPTimeseriesGroupsV2ParamsDimensionDeviceType, HTTPTimeseriesGroupsV2ParamsDimensionHTTPProtocol, HTTPTimeseriesGroupsV2ParamsDimensionHTTPVersion, HTTPTimeseriesGroupsV2ParamsDimensionIPVersion, HTTPTimeseriesGroupsV2ParamsDimensionOS, HTTPTimeseriesGroupsV2ParamsDimensionPostQuantum, HTTPTimeseriesGroupsV2ParamsDimensionTLSVersion:
		return true
	}
	return false
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesGroupsV2ParamsAggInterval string

const (
	HTTPTimeseriesGroupsV2ParamsAggInterval15m HTTPTimeseriesGroupsV2ParamsAggInterval = "15m"
	HTTPTimeseriesGroupsV2ParamsAggInterval1h  HTTPTimeseriesGroupsV2ParamsAggInterval = "1h"
	HTTPTimeseriesGroupsV2ParamsAggInterval1d  HTTPTimeseriesGroupsV2ParamsAggInterval = "1d"
	HTTPTimeseriesGroupsV2ParamsAggInterval1w  HTTPTimeseriesGroupsV2ParamsAggInterval = "1w"
)

func (r HTTPTimeseriesGroupsV2ParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupsV2ParamsAggInterval15m, HTTPTimeseriesGroupsV2ParamsAggInterval1h, HTTPTimeseriesGroupsV2ParamsAggInterval1d, HTTPTimeseriesGroupsV2ParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesGroupsV2ParamsBotClass string

const (
	HTTPTimeseriesGroupsV2ParamsBotClassLikelyAutomated HTTPTimeseriesGroupsV2ParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesGroupsV2ParamsBotClassLikelyHuman     HTTPTimeseriesGroupsV2ParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesGroupsV2ParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupsV2ParamsBotClassLikelyAutomated, HTTPTimeseriesGroupsV2ParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesGroupsV2ParamsDeviceType string

const (
	HTTPTimeseriesGroupsV2ParamsDeviceTypeDesktop HTTPTimeseriesGroupsV2ParamsDeviceType = "DESKTOP"
	HTTPTimeseriesGroupsV2ParamsDeviceTypeMobile  HTTPTimeseriesGroupsV2ParamsDeviceType = "MOBILE"
	HTTPTimeseriesGroupsV2ParamsDeviceTypeOther   HTTPTimeseriesGroupsV2ParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesGroupsV2ParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupsV2ParamsDeviceTypeDesktop, HTTPTimeseriesGroupsV2ParamsDeviceTypeMobile, HTTPTimeseriesGroupsV2ParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPTimeseriesGroupsV2ParamsFormat string

const (
	HTTPTimeseriesGroupsV2ParamsFormatJson HTTPTimeseriesGroupsV2ParamsFormat = "JSON"
	HTTPTimeseriesGroupsV2ParamsFormatCsv  HTTPTimeseriesGroupsV2ParamsFormat = "CSV"
)

func (r HTTPTimeseriesGroupsV2ParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupsV2ParamsFormatJson, HTTPTimeseriesGroupsV2ParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupsV2ParamsHTTPProtocol string

const (
	HTTPTimeseriesGroupsV2ParamsHTTPProtocolHTTP  HTTPTimeseriesGroupsV2ParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesGroupsV2ParamsHTTPProtocolHTTPS HTTPTimeseriesGroupsV2ParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesGroupsV2ParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupsV2ParamsHTTPProtocolHTTP, HTTPTimeseriesGroupsV2ParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesGroupsV2ParamsHTTPVersion string

const (
	HTTPTimeseriesGroupsV2ParamsHTTPVersionHttPv1 HTTPTimeseriesGroupsV2ParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesGroupsV2ParamsHTTPVersionHttPv2 HTTPTimeseriesGroupsV2ParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesGroupsV2ParamsHTTPVersionHttPv3 HTTPTimeseriesGroupsV2ParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesGroupsV2ParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupsV2ParamsHTTPVersionHttPv1, HTTPTimeseriesGroupsV2ParamsHTTPVersionHttPv2, HTTPTimeseriesGroupsV2ParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesGroupsV2ParamsIPVersion string

const (
	HTTPTimeseriesGroupsV2ParamsIPVersionIPv4 HTTPTimeseriesGroupsV2ParamsIPVersion = "IPv4"
	HTTPTimeseriesGroupsV2ParamsIPVersionIPv6 HTTPTimeseriesGroupsV2ParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesGroupsV2ParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupsV2ParamsIPVersionIPv4, HTTPTimeseriesGroupsV2ParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPTimeseriesGroupsV2ParamsNormalization string

const (
	HTTPTimeseriesGroupsV2ParamsNormalizationPercentage HTTPTimeseriesGroupsV2ParamsNormalization = "PERCENTAGE"
	HTTPTimeseriesGroupsV2ParamsNormalizationMin0Max    HTTPTimeseriesGroupsV2ParamsNormalization = "MIN0_MAX"
)

func (r HTTPTimeseriesGroupsV2ParamsNormalization) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupsV2ParamsNormalizationPercentage, HTTPTimeseriesGroupsV2ParamsNormalizationMin0Max:
		return true
	}
	return false
}

type HTTPTimeseriesGroupsV2ParamsOS string

const (
	HTTPTimeseriesGroupsV2ParamsOSWindows  HTTPTimeseriesGroupsV2ParamsOS = "WINDOWS"
	HTTPTimeseriesGroupsV2ParamsOSMacosx   HTTPTimeseriesGroupsV2ParamsOS = "MACOSX"
	HTTPTimeseriesGroupsV2ParamsOSIos      HTTPTimeseriesGroupsV2ParamsOS = "IOS"
	HTTPTimeseriesGroupsV2ParamsOSAndroid  HTTPTimeseriesGroupsV2ParamsOS = "ANDROID"
	HTTPTimeseriesGroupsV2ParamsOSChromeos HTTPTimeseriesGroupsV2ParamsOS = "CHROMEOS"
	HTTPTimeseriesGroupsV2ParamsOSLinux    HTTPTimeseriesGroupsV2ParamsOS = "LINUX"
	HTTPTimeseriesGroupsV2ParamsOSSmartTv  HTTPTimeseriesGroupsV2ParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesGroupsV2ParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupsV2ParamsOSWindows, HTTPTimeseriesGroupsV2ParamsOSMacosx, HTTPTimeseriesGroupsV2ParamsOSIos, HTTPTimeseriesGroupsV2ParamsOSAndroid, HTTPTimeseriesGroupsV2ParamsOSChromeos, HTTPTimeseriesGroupsV2ParamsOSLinux, HTTPTimeseriesGroupsV2ParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesGroupsV2ParamsTLSVersion string

const (
	HTTPTimeseriesGroupsV2ParamsTLSVersionTlSv1_0  HTTPTimeseriesGroupsV2ParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesGroupsV2ParamsTLSVersionTlSv1_1  HTTPTimeseriesGroupsV2ParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesGroupsV2ParamsTLSVersionTlSv1_2  HTTPTimeseriesGroupsV2ParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesGroupsV2ParamsTLSVersionTlSv1_3  HTTPTimeseriesGroupsV2ParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesGroupsV2ParamsTLSVersionTlSvQuic HTTPTimeseriesGroupsV2ParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesGroupsV2ParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesGroupsV2ParamsTLSVersionTlSv1_0, HTTPTimeseriesGroupsV2ParamsTLSVersionTlSv1_1, HTTPTimeseriesGroupsV2ParamsTLSVersionTlSv1_2, HTTPTimeseriesGroupsV2ParamsTLSVersionTlSv1_3, HTTPTimeseriesGroupsV2ParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesGroupsV2ResponseEnvelope struct {
	Result  HTTPTimeseriesGroupsV2Response             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    httpTimeseriesGroupsV2ResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesGroupsV2ResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPTimeseriesGroupsV2ResponseEnvelope]
type httpTimeseriesGroupsV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesGroupsV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesGroupsV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
