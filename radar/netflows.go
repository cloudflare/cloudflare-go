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

// NetFlowsService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetFlowsService] method instead.
type NetFlowsService struct {
	Options []option.RequestOption
	Top     *NetFlowsTopService
}

// NewNetFlowsService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetFlowsService(opts ...option.RequestOption) (r *NetFlowsService) {
	r = &NetFlowsService{}
	r.Options = opts
	r.Top = NewNetFlowsTopService(opts...)
	return
}

// Retrieves the distribution of network traffic (NetFlows) by HTTP vs other
// protocols.
//
// Deprecated: Use
// [Get Network Traffic Distribution By Dimension](https://developers.cloudflare.com/api/resources/radar/subresources/netflows/methods/summary_v2/)
// instead.
func (r *NetFlowsService) Summary(ctx context.Context, query NetFlowsSummaryParams, opts ...option.RequestOption) (res *NetFlowsSummaryResponse, err error) {
	var env NetFlowsSummaryResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/netflows/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of network traffic (NetFlows) by the specified
// dimension.
func (r *NetFlowsService) SummaryV2(ctx context.Context, dimension NetFlowsSummaryV2ParamsDimension, query NetFlowsSummaryV2Params, opts ...option.RequestOption) (res *NetFlowsSummaryV2Response, err error) {
	var env NetFlowsSummaryV2ResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/netflows/summary/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves network traffic (NetFlows) over time.
func (r *NetFlowsService) Timeseries(ctx context.Context, query NetFlowsTimeseriesParams, opts ...option.RequestOption) (res *NetFlowsTimeseriesResponse, err error) {
	var env NetFlowsTimeseriesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/netflows/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of NetFlows traffic, grouped by chosen the specified
// dimension over time.
func (r *NetFlowsService) TimeseriesGroups(ctx context.Context, dimension NetFlowsTimeseriesGroupsParamsDimension, query NetFlowsTimeseriesGroupsParams, opts ...option.RequestOption) (res *NetFlowsTimeseriesGroupsResponse, err error) {
	var env NetFlowsTimeseriesGroupsResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/netflows/timeseries_groups/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NetFlowsSummaryResponse struct {
	// Metadata for the results.
	Meta     NetFlowsSummaryResponseMeta     `json:"meta,required"`
	Summary0 NetFlowsSummaryResponseSummary0 `json:"summary_0,required"`
	JSON     netFlowsSummaryResponseJSON     `json:"-"`
}

// netFlowsSummaryResponseJSON contains the JSON metadata for the struct
// [NetFlowsSummaryResponse]
type netFlowsSummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsSummaryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type NetFlowsSummaryResponseMeta struct {
	ConfidenceInfo NetFlowsSummaryResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []NetFlowsSummaryResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization NetFlowsSummaryResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []NetFlowsSummaryResponseMetaUnit `json:"units,required"`
	JSON  netFlowsSummaryResponseMetaJSON   `json:"-"`
}

// netFlowsSummaryResponseMetaJSON contains the JSON metadata for the struct
// [NetFlowsSummaryResponseMeta]
type netFlowsSummaryResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NetFlowsSummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsSummaryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type NetFlowsSummaryResponseMetaConfidenceInfo struct {
	Annotations []NetFlowsSummaryResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                         `json:"level,required"`
	JSON  netFlowsSummaryResponseMetaConfidenceInfoJSON `json:"-"`
}

// netFlowsSummaryResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [NetFlowsSummaryResponseMetaConfidenceInfo]
type netFlowsSummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsSummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsSummaryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type NetFlowsSummaryResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                         `json:"description,required"`
	EndDate     time.Time                                                      `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                               `json:"startDate,required" format:"date-time"`
	JSON            netFlowsSummaryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// netFlowsSummaryResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [NetFlowsSummaryResponseMetaConfidenceInfoAnnotation]
type netFlowsSummaryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *NetFlowsSummaryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsSummaryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAll                NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBGP                NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBots               NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceCT                 NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNS                NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDos                NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFw                 NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceIQI                NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceNet                NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAll, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBGP, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceBots, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceCT, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNS, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceDos, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFw, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceIQI, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceNet, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventType string

const (
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventTypeEvent             NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventTypeOutage            NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventTypePipeline          NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventTypeEvent, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventTypeOutage, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventTypePipeline, NetFlowsSummaryResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type NetFlowsSummaryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                `json:"startTime,required" format:"date-time"`
	JSON      netFlowsSummaryResponseMetaDateRangeJSON `json:"-"`
}

// netFlowsSummaryResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [NetFlowsSummaryResponseMetaDateRange]
type netFlowsSummaryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsSummaryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsSummaryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type NetFlowsSummaryResponseMetaNormalization string

const (
	NetFlowsSummaryResponseMetaNormalizationPercentage           NetFlowsSummaryResponseMetaNormalization = "PERCENTAGE"
	NetFlowsSummaryResponseMetaNormalizationMin0Max              NetFlowsSummaryResponseMetaNormalization = "MIN0_MAX"
	NetFlowsSummaryResponseMetaNormalizationMinMax               NetFlowsSummaryResponseMetaNormalization = "MIN_MAX"
	NetFlowsSummaryResponseMetaNormalizationRawValues            NetFlowsSummaryResponseMetaNormalization = "RAW_VALUES"
	NetFlowsSummaryResponseMetaNormalizationPercentageChange     NetFlowsSummaryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	NetFlowsSummaryResponseMetaNormalizationRollingAverage       NetFlowsSummaryResponseMetaNormalization = "ROLLING_AVERAGE"
	NetFlowsSummaryResponseMetaNormalizationOverlappedPercentage NetFlowsSummaryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	NetFlowsSummaryResponseMetaNormalizationRatio                NetFlowsSummaryResponseMetaNormalization = "RATIO"
)

func (r NetFlowsSummaryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case NetFlowsSummaryResponseMetaNormalizationPercentage, NetFlowsSummaryResponseMetaNormalizationMin0Max, NetFlowsSummaryResponseMetaNormalizationMinMax, NetFlowsSummaryResponseMetaNormalizationRawValues, NetFlowsSummaryResponseMetaNormalizationPercentageChange, NetFlowsSummaryResponseMetaNormalizationRollingAverage, NetFlowsSummaryResponseMetaNormalizationOverlappedPercentage, NetFlowsSummaryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type NetFlowsSummaryResponseMetaUnit struct {
	Name  string                              `json:"name,required"`
	Value string                              `json:"value,required"`
	JSON  netFlowsSummaryResponseMetaUnitJSON `json:"-"`
}

// netFlowsSummaryResponseMetaUnitJSON contains the JSON metadata for the struct
// [NetFlowsSummaryResponseMetaUnit]
type netFlowsSummaryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsSummaryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsSummaryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type NetFlowsSummaryResponseSummary0 struct {
	// A numeric string.
	HTTP string `json:"HTTP,required"`
	// A numeric string.
	Other string                              `json:"OTHER,required"`
	JSON  netFlowsSummaryResponseSummary0JSON `json:"-"`
}

// netFlowsSummaryResponseSummary0JSON contains the JSON metadata for the struct
// [NetFlowsSummaryResponseSummary0]
type netFlowsSummaryResponseSummary0JSON struct {
	HTTP        apijson.Field
	Other       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsSummaryResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsSummaryResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type NetFlowsSummaryV2Response struct {
	// Metadata for the results.
	Meta     NetFlowsSummaryV2ResponseMeta `json:"meta,required"`
	Summary0 map[string]string             `json:"summary_0,required"`
	JSON     netFlowsSummaryV2ResponseJSON `json:"-"`
}

// netFlowsSummaryV2ResponseJSON contains the JSON metadata for the struct
// [NetFlowsSummaryV2Response]
type netFlowsSummaryV2ResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsSummaryV2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsSummaryV2ResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type NetFlowsSummaryV2ResponseMeta struct {
	ConfidenceInfo NetFlowsSummaryV2ResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []NetFlowsSummaryV2ResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization NetFlowsSummaryV2ResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []NetFlowsSummaryV2ResponseMetaUnit `json:"units,required"`
	JSON  netFlowsSummaryV2ResponseMetaJSON   `json:"-"`
}

// netFlowsSummaryV2ResponseMetaJSON contains the JSON metadata for the struct
// [NetFlowsSummaryV2ResponseMeta]
type netFlowsSummaryV2ResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NetFlowsSummaryV2ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsSummaryV2ResponseMetaJSON) RawJSON() string {
	return r.raw
}

type NetFlowsSummaryV2ResponseMetaConfidenceInfo struct {
	Annotations []NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                           `json:"level,required"`
	JSON  netFlowsSummaryV2ResponseMetaConfidenceInfoJSON `json:"-"`
}

// netFlowsSummaryV2ResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [NetFlowsSummaryV2ResponseMetaConfidenceInfo]
type netFlowsSummaryV2ResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsSummaryV2ResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsSummaryV2ResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                           `json:"description,required"`
	EndDate     time.Time                                                        `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                      `json:"isInstantaneous,required"`
	LinkedURL       string                                                    `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                 `json:"startDate,required" format:"date-time"`
	JSON            netFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// netFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotation]
type netFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll                NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP                NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots               NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT                 NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS                NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos                NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw                 NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI                NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet                NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAll, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBGP, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceBots, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceCT, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNS, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceDos, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFw, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceIQI, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceNet, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType string

const (
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent             NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage            NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline          NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeEvent, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeOutage, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypePipeline, NetFlowsSummaryV2ResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type NetFlowsSummaryV2ResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                  `json:"startTime,required" format:"date-time"`
	JSON      netFlowsSummaryV2ResponseMetaDateRangeJSON `json:"-"`
}

// netFlowsSummaryV2ResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [NetFlowsSummaryV2ResponseMetaDateRange]
type netFlowsSummaryV2ResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsSummaryV2ResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsSummaryV2ResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type NetFlowsSummaryV2ResponseMetaNormalization string

const (
	NetFlowsSummaryV2ResponseMetaNormalizationPercentage           NetFlowsSummaryV2ResponseMetaNormalization = "PERCENTAGE"
	NetFlowsSummaryV2ResponseMetaNormalizationMin0Max              NetFlowsSummaryV2ResponseMetaNormalization = "MIN0_MAX"
	NetFlowsSummaryV2ResponseMetaNormalizationMinMax               NetFlowsSummaryV2ResponseMetaNormalization = "MIN_MAX"
	NetFlowsSummaryV2ResponseMetaNormalizationRawValues            NetFlowsSummaryV2ResponseMetaNormalization = "RAW_VALUES"
	NetFlowsSummaryV2ResponseMetaNormalizationPercentageChange     NetFlowsSummaryV2ResponseMetaNormalization = "PERCENTAGE_CHANGE"
	NetFlowsSummaryV2ResponseMetaNormalizationRollingAverage       NetFlowsSummaryV2ResponseMetaNormalization = "ROLLING_AVERAGE"
	NetFlowsSummaryV2ResponseMetaNormalizationOverlappedPercentage NetFlowsSummaryV2ResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	NetFlowsSummaryV2ResponseMetaNormalizationRatio                NetFlowsSummaryV2ResponseMetaNormalization = "RATIO"
)

func (r NetFlowsSummaryV2ResponseMetaNormalization) IsKnown() bool {
	switch r {
	case NetFlowsSummaryV2ResponseMetaNormalizationPercentage, NetFlowsSummaryV2ResponseMetaNormalizationMin0Max, NetFlowsSummaryV2ResponseMetaNormalizationMinMax, NetFlowsSummaryV2ResponseMetaNormalizationRawValues, NetFlowsSummaryV2ResponseMetaNormalizationPercentageChange, NetFlowsSummaryV2ResponseMetaNormalizationRollingAverage, NetFlowsSummaryV2ResponseMetaNormalizationOverlappedPercentage, NetFlowsSummaryV2ResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type NetFlowsSummaryV2ResponseMetaUnit struct {
	Name  string                                `json:"name,required"`
	Value string                                `json:"value,required"`
	JSON  netFlowsSummaryV2ResponseMetaUnitJSON `json:"-"`
}

// netFlowsSummaryV2ResponseMetaUnitJSON contains the JSON metadata for the struct
// [NetFlowsSummaryV2ResponseMetaUnit]
type netFlowsSummaryV2ResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsSummaryV2ResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsSummaryV2ResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type NetFlowsTimeseriesResponse struct {
	// Metadata for the results.
	Meta   NetFlowsTimeseriesResponseMeta   `json:"meta,required"`
	Serie0 NetFlowsTimeseriesResponseSerie0 `json:"serie_0,required"`
	JSON   netFlowsTimeseriesResponseJSON   `json:"-"`
}

// netFlowsTimeseriesResponseJSON contains the JSON metadata for the struct
// [NetFlowsTimeseriesResponse]
type netFlowsTimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type NetFlowsTimeseriesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    NetFlowsTimeseriesResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo NetFlowsTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []NetFlowsTimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization NetFlowsTimeseriesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []NetFlowsTimeseriesResponseMetaUnit `json:"units,required"`
	JSON  netFlowsTimeseriesResponseMetaJSON   `json:"-"`
}

// netFlowsTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [NetFlowsTimeseriesResponseMeta]
type netFlowsTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NetFlowsTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type NetFlowsTimeseriesResponseMetaAggInterval string

const (
	NetFlowsTimeseriesResponseMetaAggIntervalFifteenMinutes NetFlowsTimeseriesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	NetFlowsTimeseriesResponseMetaAggIntervalOneHour        NetFlowsTimeseriesResponseMetaAggInterval = "ONE_HOUR"
	NetFlowsTimeseriesResponseMetaAggIntervalOneDay         NetFlowsTimeseriesResponseMetaAggInterval = "ONE_DAY"
	NetFlowsTimeseriesResponseMetaAggIntervalOneWeek        NetFlowsTimeseriesResponseMetaAggInterval = "ONE_WEEK"
	NetFlowsTimeseriesResponseMetaAggIntervalOneMonth       NetFlowsTimeseriesResponseMetaAggInterval = "ONE_MONTH"
)

func (r NetFlowsTimeseriesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case NetFlowsTimeseriesResponseMetaAggIntervalFifteenMinutes, NetFlowsTimeseriesResponseMetaAggIntervalOneHour, NetFlowsTimeseriesResponseMetaAggIntervalOneDay, NetFlowsTimeseriesResponseMetaAggIntervalOneWeek, NetFlowsTimeseriesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type NetFlowsTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                            `json:"level,required"`
	JSON  netFlowsTimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// netFlowsTimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [NetFlowsTimeseriesResponseMetaConfidenceInfo]
type netFlowsTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                            `json:"description,required"`
	EndDate     time.Time                                                         `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                       `json:"isInstantaneous,required"`
	LinkedURL       string                                                     `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                  `json:"startDate,required" format:"date-time"`
	JSON            netFlowsTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// netFlowsTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotation]
type netFlowsTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll                NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP                NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots               NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCT                 NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS                NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos                NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw                 NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI                NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet                NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCT, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventType string

const (
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent             NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage            NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline          NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline, NetFlowsTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type NetFlowsTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                   `json:"startTime,required" format:"date-time"`
	JSON      netFlowsTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// netFlowsTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [NetFlowsTimeseriesResponseMetaDateRange]
type netFlowsTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type NetFlowsTimeseriesResponseMetaNormalization string

const (
	NetFlowsTimeseriesResponseMetaNormalizationPercentage           NetFlowsTimeseriesResponseMetaNormalization = "PERCENTAGE"
	NetFlowsTimeseriesResponseMetaNormalizationMin0Max              NetFlowsTimeseriesResponseMetaNormalization = "MIN0_MAX"
	NetFlowsTimeseriesResponseMetaNormalizationMinMax               NetFlowsTimeseriesResponseMetaNormalization = "MIN_MAX"
	NetFlowsTimeseriesResponseMetaNormalizationRawValues            NetFlowsTimeseriesResponseMetaNormalization = "RAW_VALUES"
	NetFlowsTimeseriesResponseMetaNormalizationPercentageChange     NetFlowsTimeseriesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	NetFlowsTimeseriesResponseMetaNormalizationRollingAverage       NetFlowsTimeseriesResponseMetaNormalization = "ROLLING_AVERAGE"
	NetFlowsTimeseriesResponseMetaNormalizationOverlappedPercentage NetFlowsTimeseriesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	NetFlowsTimeseriesResponseMetaNormalizationRatio                NetFlowsTimeseriesResponseMetaNormalization = "RATIO"
)

func (r NetFlowsTimeseriesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case NetFlowsTimeseriesResponseMetaNormalizationPercentage, NetFlowsTimeseriesResponseMetaNormalizationMin0Max, NetFlowsTimeseriesResponseMetaNormalizationMinMax, NetFlowsTimeseriesResponseMetaNormalizationRawValues, NetFlowsTimeseriesResponseMetaNormalizationPercentageChange, NetFlowsTimeseriesResponseMetaNormalizationRollingAverage, NetFlowsTimeseriesResponseMetaNormalizationOverlappedPercentage, NetFlowsTimeseriesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type NetFlowsTimeseriesResponseMetaUnit struct {
	Name  string                                 `json:"name,required"`
	Value string                                 `json:"value,required"`
	JSON  netFlowsTimeseriesResponseMetaUnitJSON `json:"-"`
}

// netFlowsTimeseriesResponseMetaUnitJSON contains the JSON metadata for the struct
// [NetFlowsTimeseriesResponseMetaUnit]
type netFlowsTimeseriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTimeseriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTimeseriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type NetFlowsTimeseriesResponseSerie0 struct {
	Timestamps []time.Time                          `json:"timestamps,required" format:"date-time"`
	Values     []string                             `json:"values,required"`
	JSON       netFlowsTimeseriesResponseSerie0JSON `json:"-"`
}

// netFlowsTimeseriesResponseSerie0JSON contains the JSON metadata for the struct
// [NetFlowsTimeseriesResponseSerie0]
type netFlowsTimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type NetFlowsTimeseriesGroupsResponse struct {
	// Metadata for the results.
	Meta   NetFlowsTimeseriesGroupsResponseMeta   `json:"meta,required"`
	Serie0 NetFlowsTimeseriesGroupsResponseSerie0 `json:"serie_0,required"`
	JSON   netFlowsTimeseriesGroupsResponseJSON   `json:"-"`
}

// netFlowsTimeseriesGroupsResponseJSON contains the JSON metadata for the struct
// [NetFlowsTimeseriesGroupsResponse]
type netFlowsTimeseriesGroupsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type NetFlowsTimeseriesGroupsResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    NetFlowsTimeseriesGroupsResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo NetFlowsTimeseriesGroupsResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []NetFlowsTimeseriesGroupsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization NetFlowsTimeseriesGroupsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []NetFlowsTimeseriesGroupsResponseMetaUnit `json:"units,required"`
	JSON  netFlowsTimeseriesGroupsResponseMetaJSON   `json:"-"`
}

// netFlowsTimeseriesGroupsResponseMetaJSON contains the JSON metadata for the
// struct [NetFlowsTimeseriesGroupsResponseMeta]
type netFlowsTimeseriesGroupsResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NetFlowsTimeseriesGroupsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTimeseriesGroupsResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type NetFlowsTimeseriesGroupsResponseMetaAggInterval string

const (
	NetFlowsTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes NetFlowsTimeseriesGroupsResponseMetaAggInterval = "FIFTEEN_MINUTES"
	NetFlowsTimeseriesGroupsResponseMetaAggIntervalOneHour        NetFlowsTimeseriesGroupsResponseMetaAggInterval = "ONE_HOUR"
	NetFlowsTimeseriesGroupsResponseMetaAggIntervalOneDay         NetFlowsTimeseriesGroupsResponseMetaAggInterval = "ONE_DAY"
	NetFlowsTimeseriesGroupsResponseMetaAggIntervalOneWeek        NetFlowsTimeseriesGroupsResponseMetaAggInterval = "ONE_WEEK"
	NetFlowsTimeseriesGroupsResponseMetaAggIntervalOneMonth       NetFlowsTimeseriesGroupsResponseMetaAggInterval = "ONE_MONTH"
)

func (r NetFlowsTimeseriesGroupsResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case NetFlowsTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes, NetFlowsTimeseriesGroupsResponseMetaAggIntervalOneHour, NetFlowsTimeseriesGroupsResponseMetaAggIntervalOneDay, NetFlowsTimeseriesGroupsResponseMetaAggIntervalOneWeek, NetFlowsTimeseriesGroupsResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type NetFlowsTimeseriesGroupsResponseMetaConfidenceInfo struct {
	Annotations []NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                  `json:"level,required"`
	JSON  netFlowsTimeseriesGroupsResponseMetaConfidenceInfoJSON `json:"-"`
}

// netFlowsTimeseriesGroupsResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [NetFlowsTimeseriesGroupsResponseMetaConfidenceInfo]
type netFlowsTimeseriesGroupsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTimeseriesGroupsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTimeseriesGroupsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                                  `json:"description,required"`
	EndDate     time.Time                                                               `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                             `json:"isInstantaneous,required"`
	LinkedURL       string                                                           `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                        `json:"startDate,required" format:"date-time"`
	JSON            netFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// netFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotation]
type netFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAll                NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceBGP                NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceBots               NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceCT                 NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNS                NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDos                NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceFw                 NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceIQI                NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceNet                NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAll, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceBGP, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceBots, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceCT, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNS, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceDos, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceFw, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceIQI, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceNet, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType string

const (
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeEvent             NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeOutage            NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypePipeline          NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeEvent, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeOutage, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypePipeline, NetFlowsTimeseriesGroupsResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type NetFlowsTimeseriesGroupsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      netFlowsTimeseriesGroupsResponseMetaDateRangeJSON `json:"-"`
}

// netFlowsTimeseriesGroupsResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [NetFlowsTimeseriesGroupsResponseMetaDateRange]
type netFlowsTimeseriesGroupsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTimeseriesGroupsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTimeseriesGroupsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type NetFlowsTimeseriesGroupsResponseMetaNormalization string

const (
	NetFlowsTimeseriesGroupsResponseMetaNormalizationPercentage           NetFlowsTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE"
	NetFlowsTimeseriesGroupsResponseMetaNormalizationMin0Max              NetFlowsTimeseriesGroupsResponseMetaNormalization = "MIN0_MAX"
	NetFlowsTimeseriesGroupsResponseMetaNormalizationMinMax               NetFlowsTimeseriesGroupsResponseMetaNormalization = "MIN_MAX"
	NetFlowsTimeseriesGroupsResponseMetaNormalizationRawValues            NetFlowsTimeseriesGroupsResponseMetaNormalization = "RAW_VALUES"
	NetFlowsTimeseriesGroupsResponseMetaNormalizationPercentageChange     NetFlowsTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	NetFlowsTimeseriesGroupsResponseMetaNormalizationRollingAverage       NetFlowsTimeseriesGroupsResponseMetaNormalization = "ROLLING_AVERAGE"
	NetFlowsTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage NetFlowsTimeseriesGroupsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	NetFlowsTimeseriesGroupsResponseMetaNormalizationRatio                NetFlowsTimeseriesGroupsResponseMetaNormalization = "RATIO"
)

func (r NetFlowsTimeseriesGroupsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case NetFlowsTimeseriesGroupsResponseMetaNormalizationPercentage, NetFlowsTimeseriesGroupsResponseMetaNormalizationMin0Max, NetFlowsTimeseriesGroupsResponseMetaNormalizationMinMax, NetFlowsTimeseriesGroupsResponseMetaNormalizationRawValues, NetFlowsTimeseriesGroupsResponseMetaNormalizationPercentageChange, NetFlowsTimeseriesGroupsResponseMetaNormalizationRollingAverage, NetFlowsTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage, NetFlowsTimeseriesGroupsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type NetFlowsTimeseriesGroupsResponseMetaUnit struct {
	Name  string                                       `json:"name,required"`
	Value string                                       `json:"value,required"`
	JSON  netFlowsTimeseriesGroupsResponseMetaUnitJSON `json:"-"`
}

// netFlowsTimeseriesGroupsResponseMetaUnitJSON contains the JSON metadata for the
// struct [NetFlowsTimeseriesGroupsResponseMetaUnit]
type netFlowsTimeseriesGroupsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTimeseriesGroupsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTimeseriesGroupsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type NetFlowsTimeseriesGroupsResponseSerie0 struct {
	Timestamps  []time.Time                                `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                        `json:"-,extras"`
	JSON        netFlowsTimeseriesGroupsResponseSerie0JSON `json:"-"`
}

// netFlowsTimeseriesGroupsResponseSerie0JSON contains the JSON metadata for the
// struct [NetFlowsTimeseriesGroupsResponseSerie0]
type netFlowsTimeseriesGroupsResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTimeseriesGroupsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTimeseriesGroupsResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type NetFlowsSummaryParams struct {
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
	Format param.Field[NetFlowsSummaryParamsFormat] `query:"format"`
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

// URLQuery serializes [NetFlowsSummaryParams]'s query parameters as `url.Values`.
func (r NetFlowsSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type NetFlowsSummaryParamsFormat string

const (
	NetFlowsSummaryParamsFormatJson NetFlowsSummaryParamsFormat = "JSON"
	NetFlowsSummaryParamsFormatCsv  NetFlowsSummaryParamsFormat = "CSV"
)

func (r NetFlowsSummaryParamsFormat) IsKnown() bool {
	switch r {
	case NetFlowsSummaryParamsFormatJson, NetFlowsSummaryParamsFormatCsv:
		return true
	}
	return false
}

type NetFlowsSummaryResponseEnvelope struct {
	Result  NetFlowsSummaryResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    netFlowsSummaryResponseEnvelopeJSON `json:"-"`
}

// netFlowsSummaryResponseEnvelopeJSON contains the JSON metadata for the struct
// [NetFlowsSummaryResponseEnvelope]
type netFlowsSummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsSummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NetFlowsSummaryV2Params struct {
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
	Format param.Field[NetFlowsSummaryV2ParamsFormat] `query:"format"`
	// Filters results by Geolocation. Specify a comma-separated list of GeoNames IDs.
	// Prefix with `-` to exclude geoIds from results. For example, `-2267056,360689`
	// excludes results from the 2267056 (Lisbon), but includes results from 5128638
	// (New York).
	GeoID param.Field[[]string] `query:"geoId"`
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
	// Filters the results by network traffic product types.
	Product param.Field[[]NetFlowsSummaryV2ParamsProduct] `query:"product"`
}

// URLQuery serializes [NetFlowsSummaryV2Params]'s query parameters as
// `url.Values`.
func (r NetFlowsSummaryV2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the NetFlows attribute by which to group the results.
type NetFlowsSummaryV2ParamsDimension string

const (
	NetFlowsSummaryV2ParamsDimensionAdm1    NetFlowsSummaryV2ParamsDimension = "ADM1"
	NetFlowsSummaryV2ParamsDimensionProduct NetFlowsSummaryV2ParamsDimension = "PRODUCT"
)

func (r NetFlowsSummaryV2ParamsDimension) IsKnown() bool {
	switch r {
	case NetFlowsSummaryV2ParamsDimensionAdm1, NetFlowsSummaryV2ParamsDimensionProduct:
		return true
	}
	return false
}

// Format in which results will be returned.
type NetFlowsSummaryV2ParamsFormat string

const (
	NetFlowsSummaryV2ParamsFormatJson NetFlowsSummaryV2ParamsFormat = "JSON"
	NetFlowsSummaryV2ParamsFormatCsv  NetFlowsSummaryV2ParamsFormat = "CSV"
)

func (r NetFlowsSummaryV2ParamsFormat) IsKnown() bool {
	switch r {
	case NetFlowsSummaryV2ParamsFormatJson, NetFlowsSummaryV2ParamsFormatCsv:
		return true
	}
	return false
}

type NetFlowsSummaryV2ParamsProduct string

const (
	NetFlowsSummaryV2ParamsProductHTTP NetFlowsSummaryV2ParamsProduct = "HTTP"
	NetFlowsSummaryV2ParamsProductAll  NetFlowsSummaryV2ParamsProduct = "ALL"
)

func (r NetFlowsSummaryV2ParamsProduct) IsKnown() bool {
	switch r {
	case NetFlowsSummaryV2ParamsProductHTTP, NetFlowsSummaryV2ParamsProductAll:
		return true
	}
	return false
}

type NetFlowsSummaryV2ResponseEnvelope struct {
	Result  NetFlowsSummaryV2Response             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    netFlowsSummaryV2ResponseEnvelopeJSON `json:"-"`
}

// netFlowsSummaryV2ResponseEnvelopeJSON contains the JSON metadata for the struct
// [NetFlowsSummaryV2ResponseEnvelope]
type netFlowsSummaryV2ResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsSummaryV2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsSummaryV2ResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NetFlowsTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[NetFlowsTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[NetFlowsTimeseriesParamsFormat] `query:"format"`
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
	Normalization param.Field[NetFlowsTimeseriesParamsNormalization] `query:"normalization"`
	// Filters the results by network traffic product types.
	Product param.Field[[]NetFlowsTimeseriesParamsProduct] `query:"product"`
}

// URLQuery serializes [NetFlowsTimeseriesParams]'s query parameters as
// `url.Values`.
func (r NetFlowsTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type NetFlowsTimeseriesParamsAggInterval string

const (
	NetFlowsTimeseriesParamsAggInterval15m NetFlowsTimeseriesParamsAggInterval = "15m"
	NetFlowsTimeseriesParamsAggInterval1h  NetFlowsTimeseriesParamsAggInterval = "1h"
	NetFlowsTimeseriesParamsAggInterval1d  NetFlowsTimeseriesParamsAggInterval = "1d"
	NetFlowsTimeseriesParamsAggInterval1w  NetFlowsTimeseriesParamsAggInterval = "1w"
)

func (r NetFlowsTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case NetFlowsTimeseriesParamsAggInterval15m, NetFlowsTimeseriesParamsAggInterval1h, NetFlowsTimeseriesParamsAggInterval1d, NetFlowsTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type NetFlowsTimeseriesParamsFormat string

const (
	NetFlowsTimeseriesParamsFormatJson NetFlowsTimeseriesParamsFormat = "JSON"
	NetFlowsTimeseriesParamsFormatCsv  NetFlowsTimeseriesParamsFormat = "CSV"
)

func (r NetFlowsTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case NetFlowsTimeseriesParamsFormatJson, NetFlowsTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type NetFlowsTimeseriesParamsNormalization string

const (
	NetFlowsTimeseriesParamsNormalizationPercentageChange NetFlowsTimeseriesParamsNormalization = "PERCENTAGE_CHANGE"
	NetFlowsTimeseriesParamsNormalizationMin0Max          NetFlowsTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r NetFlowsTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case NetFlowsTimeseriesParamsNormalizationPercentageChange, NetFlowsTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type NetFlowsTimeseriesParamsProduct string

const (
	NetFlowsTimeseriesParamsProductHTTP NetFlowsTimeseriesParamsProduct = "HTTP"
	NetFlowsTimeseriesParamsProductAll  NetFlowsTimeseriesParamsProduct = "ALL"
)

func (r NetFlowsTimeseriesParamsProduct) IsKnown() bool {
	switch r {
	case NetFlowsTimeseriesParamsProductHTTP, NetFlowsTimeseriesParamsProductAll:
		return true
	}
	return false
}

type NetFlowsTimeseriesResponseEnvelope struct {
	Result  NetFlowsTimeseriesResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    netFlowsTimeseriesResponseEnvelopeJSON `json:"-"`
}

// netFlowsTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [NetFlowsTimeseriesResponseEnvelope]
type netFlowsTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NetFlowsTimeseriesGroupsParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[NetFlowsTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[NetFlowsTimeseriesGroupsParamsFormat] `query:"format"`
	// Filters results by Geolocation. Specify a comma-separated list of GeoNames IDs.
	// Prefix with `-` to exclude geoIds from results. For example, `-2267056,360689`
	// excludes results from the 2267056 (Lisbon), but includes results from 5128638
	// (New York).
	GeoID param.Field[[]string] `query:"geoId"`
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
	Normalization param.Field[NetFlowsTimeseriesGroupsParamsNormalization] `query:"normalization"`
	// Filters the results by network traffic product types.
	Product param.Field[[]NetFlowsTimeseriesGroupsParamsProduct] `query:"product"`
}

// URLQuery serializes [NetFlowsTimeseriesGroupsParams]'s query parameters as
// `url.Values`.
func (r NetFlowsTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the NetFlows attribute by which to group the results.
type NetFlowsTimeseriesGroupsParamsDimension string

const (
	NetFlowsTimeseriesGroupsParamsDimensionAdm1    NetFlowsTimeseriesGroupsParamsDimension = "ADM1"
	NetFlowsTimeseriesGroupsParamsDimensionProduct NetFlowsTimeseriesGroupsParamsDimension = "PRODUCT"
)

func (r NetFlowsTimeseriesGroupsParamsDimension) IsKnown() bool {
	switch r {
	case NetFlowsTimeseriesGroupsParamsDimensionAdm1, NetFlowsTimeseriesGroupsParamsDimensionProduct:
		return true
	}
	return false
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type NetFlowsTimeseriesGroupsParamsAggInterval string

const (
	NetFlowsTimeseriesGroupsParamsAggInterval15m NetFlowsTimeseriesGroupsParamsAggInterval = "15m"
	NetFlowsTimeseriesGroupsParamsAggInterval1h  NetFlowsTimeseriesGroupsParamsAggInterval = "1h"
	NetFlowsTimeseriesGroupsParamsAggInterval1d  NetFlowsTimeseriesGroupsParamsAggInterval = "1d"
	NetFlowsTimeseriesGroupsParamsAggInterval1w  NetFlowsTimeseriesGroupsParamsAggInterval = "1w"
)

func (r NetFlowsTimeseriesGroupsParamsAggInterval) IsKnown() bool {
	switch r {
	case NetFlowsTimeseriesGroupsParamsAggInterval15m, NetFlowsTimeseriesGroupsParamsAggInterval1h, NetFlowsTimeseriesGroupsParamsAggInterval1d, NetFlowsTimeseriesGroupsParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type NetFlowsTimeseriesGroupsParamsFormat string

const (
	NetFlowsTimeseriesGroupsParamsFormatJson NetFlowsTimeseriesGroupsParamsFormat = "JSON"
	NetFlowsTimeseriesGroupsParamsFormatCsv  NetFlowsTimeseriesGroupsParamsFormat = "CSV"
)

func (r NetFlowsTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case NetFlowsTimeseriesGroupsParamsFormatJson, NetFlowsTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type NetFlowsTimeseriesGroupsParamsNormalization string

const (
	NetFlowsTimeseriesGroupsParamsNormalizationPercentage NetFlowsTimeseriesGroupsParamsNormalization = "PERCENTAGE"
	NetFlowsTimeseriesGroupsParamsNormalizationMin0Max    NetFlowsTimeseriesGroupsParamsNormalization = "MIN0_MAX"
)

func (r NetFlowsTimeseriesGroupsParamsNormalization) IsKnown() bool {
	switch r {
	case NetFlowsTimeseriesGroupsParamsNormalizationPercentage, NetFlowsTimeseriesGroupsParamsNormalizationMin0Max:
		return true
	}
	return false
}

type NetFlowsTimeseriesGroupsParamsProduct string

const (
	NetFlowsTimeseriesGroupsParamsProductHTTP NetFlowsTimeseriesGroupsParamsProduct = "HTTP"
	NetFlowsTimeseriesGroupsParamsProductAll  NetFlowsTimeseriesGroupsParamsProduct = "ALL"
)

func (r NetFlowsTimeseriesGroupsParamsProduct) IsKnown() bool {
	switch r {
	case NetFlowsTimeseriesGroupsParamsProductHTTP, NetFlowsTimeseriesGroupsParamsProductAll:
		return true
	}
	return false
}

type NetFlowsTimeseriesGroupsResponseEnvelope struct {
	Result  NetFlowsTimeseriesGroupsResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    netFlowsTimeseriesGroupsResponseEnvelopeJSON `json:"-"`
}

// netFlowsTimeseriesGroupsResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetFlowsTimeseriesGroupsResponseEnvelope]
type netFlowsTimeseriesGroupsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTimeseriesGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTimeseriesGroupsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
