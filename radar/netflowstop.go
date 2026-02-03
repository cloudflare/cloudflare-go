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

// NetFlowsTopService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetFlowsTopService] method instead.
type NetFlowsTopService struct {
	Options []option.RequestOption
}

// NewNetFlowsTopService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetFlowsTopService(opts ...option.RequestOption) (r *NetFlowsTopService) {
	r = &NetFlowsTopService{}
	r.Options = opts
	return
}

// Retrieves the top autonomous systems by network traffic (NetFlows).
func (r *NetFlowsTopService) Ases(ctx context.Context, query NetFlowsTopAsesParams, opts ...option.RequestOption) (res *NetFlowsTopAsesResponse, err error) {
	var env NetFlowsTopAsesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/netflows/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the top locations by network traffic (NetFlows).
func (r *NetFlowsTopService) Locations(ctx context.Context, query NetFlowsTopLocationsParams, opts ...option.RequestOption) (res *NetFlowsTopLocationsResponse, err error) {
	var env NetFlowsTopLocationsResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/netflows/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NetFlowsTopAsesResponse struct {
	// Metadata for the results.
	Meta NetFlowsTopAsesResponseMeta   `json:"meta,required"`
	Top0 []NetFlowsTopAsesResponseTop0 `json:"top_0,required"`
	JSON netFlowsTopAsesResponseJSON   `json:"-"`
}

// netFlowsTopAsesResponseJSON contains the JSON metadata for the struct
// [NetFlowsTopAsesResponse]
type netFlowsTopAsesResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTopAsesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTopAsesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type NetFlowsTopAsesResponseMeta struct {
	ConfidenceInfo NetFlowsTopAsesResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []NetFlowsTopAsesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization NetFlowsTopAsesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []NetFlowsTopAsesResponseMetaUnit `json:"units,required"`
	JSON  netFlowsTopAsesResponseMetaJSON   `json:"-"`
}

// netFlowsTopAsesResponseMetaJSON contains the JSON metadata for the struct
// [NetFlowsTopAsesResponseMeta]
type netFlowsTopAsesResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NetFlowsTopAsesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTopAsesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type NetFlowsTopAsesResponseMetaConfidenceInfo struct {
	Annotations []NetFlowsTopAsesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                         `json:"level,required"`
	JSON  netFlowsTopAsesResponseMetaConfidenceInfoJSON `json:"-"`
}

// netFlowsTopAsesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [NetFlowsTopAsesResponseMetaConfidenceInfo]
type netFlowsTopAsesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTopAsesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTopAsesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type NetFlowsTopAsesResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                         `json:"description,required"`
	EndDate     time.Time                                                      `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                               `json:"startDate,required" format:"date-time"`
	JSON            netFlowsTopAsesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// netFlowsTopAsesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [NetFlowsTopAsesResponseMetaConfidenceInfoAnnotation]
type netFlowsTopAsesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *NetFlowsTopAsesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTopAsesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceAll                NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceBGP                NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceBots               NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceCT                 NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceDNS                NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceDos                NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceFw                 NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceIQI                NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceNet                NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceAll, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceBGP, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceBots, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceCT, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceDNS, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceDos, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceFw, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceIQI, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceNet, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventType string

const (
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventTypeEvent             NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventTypeOutage            NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventTypePipeline          NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventTypeEvent, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventTypeOutage, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventTypePipeline, NetFlowsTopAsesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type NetFlowsTopAsesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                `json:"startTime,required" format:"date-time"`
	JSON      netFlowsTopAsesResponseMetaDateRangeJSON `json:"-"`
}

// netFlowsTopAsesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [NetFlowsTopAsesResponseMetaDateRange]
type netFlowsTopAsesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTopAsesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTopAsesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type NetFlowsTopAsesResponseMetaNormalization string

const (
	NetFlowsTopAsesResponseMetaNormalizationPercentage           NetFlowsTopAsesResponseMetaNormalization = "PERCENTAGE"
	NetFlowsTopAsesResponseMetaNormalizationMin0Max              NetFlowsTopAsesResponseMetaNormalization = "MIN0_MAX"
	NetFlowsTopAsesResponseMetaNormalizationMinMax               NetFlowsTopAsesResponseMetaNormalization = "MIN_MAX"
	NetFlowsTopAsesResponseMetaNormalizationRawValues            NetFlowsTopAsesResponseMetaNormalization = "RAW_VALUES"
	NetFlowsTopAsesResponseMetaNormalizationPercentageChange     NetFlowsTopAsesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	NetFlowsTopAsesResponseMetaNormalizationRollingAverage       NetFlowsTopAsesResponseMetaNormalization = "ROLLING_AVERAGE"
	NetFlowsTopAsesResponseMetaNormalizationOverlappedPercentage NetFlowsTopAsesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	NetFlowsTopAsesResponseMetaNormalizationRatio                NetFlowsTopAsesResponseMetaNormalization = "RATIO"
)

func (r NetFlowsTopAsesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case NetFlowsTopAsesResponseMetaNormalizationPercentage, NetFlowsTopAsesResponseMetaNormalizationMin0Max, NetFlowsTopAsesResponseMetaNormalizationMinMax, NetFlowsTopAsesResponseMetaNormalizationRawValues, NetFlowsTopAsesResponseMetaNormalizationPercentageChange, NetFlowsTopAsesResponseMetaNormalizationRollingAverage, NetFlowsTopAsesResponseMetaNormalizationOverlappedPercentage, NetFlowsTopAsesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type NetFlowsTopAsesResponseMetaUnit struct {
	Name  string                              `json:"name,required"`
	Value string                              `json:"value,required"`
	JSON  netFlowsTopAsesResponseMetaUnitJSON `json:"-"`
}

// netFlowsTopAsesResponseMetaUnitJSON contains the JSON metadata for the struct
// [NetFlowsTopAsesResponseMetaUnit]
type netFlowsTopAsesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTopAsesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTopAsesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type NetFlowsTopAsesResponseTop0 struct {
	ClientASN    float64 `json:"clientASN,required"`
	ClientAsName string  `json:"clientASName,required"`
	// A numeric string.
	Value string                          `json:"value,required"`
	JSON  netFlowsTopAsesResponseTop0JSON `json:"-"`
}

// netFlowsTopAsesResponseTop0JSON contains the JSON metadata for the struct
// [NetFlowsTopAsesResponseTop0]
type netFlowsTopAsesResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NetFlowsTopAsesResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTopAsesResponseTop0JSON) RawJSON() string {
	return r.raw
}

type NetFlowsTopLocationsResponse struct {
	// Metadata for the results.
	Meta NetFlowsTopLocationsResponseMeta   `json:"meta,required"`
	Top0 []NetFlowsTopLocationsResponseTop0 `json:"top_0,required"`
	JSON netFlowsTopLocationsResponseJSON   `json:"-"`
}

// netFlowsTopLocationsResponseJSON contains the JSON metadata for the struct
// [NetFlowsTopLocationsResponse]
type netFlowsTopLocationsResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTopLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTopLocationsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type NetFlowsTopLocationsResponseMeta struct {
	ConfidenceInfo NetFlowsTopLocationsResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []NetFlowsTopLocationsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization NetFlowsTopLocationsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []NetFlowsTopLocationsResponseMetaUnit `json:"units,required"`
	JSON  netFlowsTopLocationsResponseMetaJSON   `json:"-"`
}

// netFlowsTopLocationsResponseMetaJSON contains the JSON metadata for the struct
// [NetFlowsTopLocationsResponseMeta]
type netFlowsTopLocationsResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NetFlowsTopLocationsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTopLocationsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type NetFlowsTopLocationsResponseMetaConfidenceInfo struct {
	Annotations []NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                              `json:"level,required"`
	JSON  netFlowsTopLocationsResponseMetaConfidenceInfoJSON `json:"-"`
}

// netFlowsTopLocationsResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [NetFlowsTopLocationsResponseMetaConfidenceInfo]
type netFlowsTopLocationsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTopLocationsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTopLocationsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                              `json:"description,required"`
	EndDate     time.Time                                                           `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                    `json:"startDate,required" format:"date-time"`
	JSON            netFlowsTopLocationsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// netFlowsTopLocationsResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotation]
type netFlowsTopLocationsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTopLocationsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceAll                NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceBGP                NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceBots               NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceCT                 NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDNS                NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDos                NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceFw                 NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceIQI                NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceNet                NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceAll, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceBGP, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceBots, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceCT, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDNS, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDos, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceFw, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceIQI, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceNet, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventType string

const (
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeEvent             NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeOutage            NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypePipeline          NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeEvent, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeOutage, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypePipeline, NetFlowsTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type NetFlowsTopLocationsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                     `json:"startTime,required" format:"date-time"`
	JSON      netFlowsTopLocationsResponseMetaDateRangeJSON `json:"-"`
}

// netFlowsTopLocationsResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [NetFlowsTopLocationsResponseMetaDateRange]
type netFlowsTopLocationsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTopLocationsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTopLocationsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type NetFlowsTopLocationsResponseMetaNormalization string

const (
	NetFlowsTopLocationsResponseMetaNormalizationPercentage           NetFlowsTopLocationsResponseMetaNormalization = "PERCENTAGE"
	NetFlowsTopLocationsResponseMetaNormalizationMin0Max              NetFlowsTopLocationsResponseMetaNormalization = "MIN0_MAX"
	NetFlowsTopLocationsResponseMetaNormalizationMinMax               NetFlowsTopLocationsResponseMetaNormalization = "MIN_MAX"
	NetFlowsTopLocationsResponseMetaNormalizationRawValues            NetFlowsTopLocationsResponseMetaNormalization = "RAW_VALUES"
	NetFlowsTopLocationsResponseMetaNormalizationPercentageChange     NetFlowsTopLocationsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	NetFlowsTopLocationsResponseMetaNormalizationRollingAverage       NetFlowsTopLocationsResponseMetaNormalization = "ROLLING_AVERAGE"
	NetFlowsTopLocationsResponseMetaNormalizationOverlappedPercentage NetFlowsTopLocationsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	NetFlowsTopLocationsResponseMetaNormalizationRatio                NetFlowsTopLocationsResponseMetaNormalization = "RATIO"
)

func (r NetFlowsTopLocationsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case NetFlowsTopLocationsResponseMetaNormalizationPercentage, NetFlowsTopLocationsResponseMetaNormalizationMin0Max, NetFlowsTopLocationsResponseMetaNormalizationMinMax, NetFlowsTopLocationsResponseMetaNormalizationRawValues, NetFlowsTopLocationsResponseMetaNormalizationPercentageChange, NetFlowsTopLocationsResponseMetaNormalizationRollingAverage, NetFlowsTopLocationsResponseMetaNormalizationOverlappedPercentage, NetFlowsTopLocationsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type NetFlowsTopLocationsResponseMetaUnit struct {
	Name  string                                   `json:"name,required"`
	Value string                                   `json:"value,required"`
	JSON  netFlowsTopLocationsResponseMetaUnitJSON `json:"-"`
}

// netFlowsTopLocationsResponseMetaUnitJSON contains the JSON metadata for the
// struct [NetFlowsTopLocationsResponseMetaUnit]
type netFlowsTopLocationsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTopLocationsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTopLocationsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type NetFlowsTopLocationsResponseTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                               `json:"value,required"`
	JSON  netFlowsTopLocationsResponseTop0JSON `json:"-"`
}

// netFlowsTopLocationsResponseTop0JSON contains the JSON metadata for the struct
// [NetFlowsTopLocationsResponseTop0]
type netFlowsTopLocationsResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *NetFlowsTopLocationsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTopLocationsResponseTop0JSON) RawJSON() string {
	return r.raw
}

type NetFlowsTopAsesParams struct {
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
	Format param.Field[NetFlowsTopAsesParamsFormat] `query:"format"`
	// Filters results by Geolocation. Specify a comma-separated list of GeoNames IDs.
	// Prefix with `-` to exclude geoIds from results. For example, `-2267056,360689`
	// excludes results from the 2267056 (Lisbon), but includes results from 5128638
	// (New York).
	GeoID param.Field[[]string] `query:"geoId"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [NetFlowsTopAsesParams]'s query parameters as `url.Values`.
func (r NetFlowsTopAsesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type NetFlowsTopAsesParamsFormat string

const (
	NetFlowsTopAsesParamsFormatJson NetFlowsTopAsesParamsFormat = "JSON"
	NetFlowsTopAsesParamsFormatCsv  NetFlowsTopAsesParamsFormat = "CSV"
)

func (r NetFlowsTopAsesParamsFormat) IsKnown() bool {
	switch r {
	case NetFlowsTopAsesParamsFormatJson, NetFlowsTopAsesParamsFormatCsv:
		return true
	}
	return false
}

type NetFlowsTopAsesResponseEnvelope struct {
	Result  NetFlowsTopAsesResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    netFlowsTopAsesResponseEnvelopeJSON `json:"-"`
}

// netFlowsTopAsesResponseEnvelopeJSON contains the JSON metadata for the struct
// [NetFlowsTopAsesResponseEnvelope]
type netFlowsTopAsesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTopAsesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTopAsesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NetFlowsTopLocationsParams struct {
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
	Format param.Field[NetFlowsTopLocationsParamsFormat] `query:"format"`
	// Filters results by Geolocation. Specify a comma-separated list of GeoNames IDs.
	// Prefix with `-` to exclude geoIds from results. For example, `-2267056,360689`
	// excludes results from the 2267056 (Lisbon), but includes results from 5128638
	// (New York).
	GeoID param.Field[[]string] `query:"geoId"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [NetFlowsTopLocationsParams]'s query parameters as
// `url.Values`.
func (r NetFlowsTopLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type NetFlowsTopLocationsParamsFormat string

const (
	NetFlowsTopLocationsParamsFormatJson NetFlowsTopLocationsParamsFormat = "JSON"
	NetFlowsTopLocationsParamsFormatCsv  NetFlowsTopLocationsParamsFormat = "CSV"
)

func (r NetFlowsTopLocationsParamsFormat) IsKnown() bool {
	switch r {
	case NetFlowsTopLocationsParamsFormatJson, NetFlowsTopLocationsParamsFormatCsv:
		return true
	}
	return false
}

type NetFlowsTopLocationsResponseEnvelope struct {
	Result  NetFlowsTopLocationsResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    netFlowsTopLocationsResponseEnvelopeJSON `json:"-"`
}

// netFlowsTopLocationsResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetFlowsTopLocationsResponseEnvelope]
type netFlowsTopLocationsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetFlowsTopLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netFlowsTopLocationsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
