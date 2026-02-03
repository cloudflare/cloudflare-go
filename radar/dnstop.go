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

// DNSTopService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSTopService] method instead.
type DNSTopService struct {
	Options []option.RequestOption
}

// NewDNSTopService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDNSTopService(opts ...option.RequestOption) (r *DNSTopService) {
	r = &DNSTopService{}
	r.Options = opts
	return
}

// Retrieves the top autonomous systems by DNS queries made to 1.1.1.1 DNS
// resolver.
func (r *DNSTopService) Ases(ctx context.Context, query DNSTopAsesParams, opts ...option.RequestOption) (res *DNSTopAsesResponse, err error) {
	var env DNSTopAsesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/dns/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the top locations by DNS queries made to 1.1.1.1 DNS resolver.
func (r *DNSTopService) Locations(ctx context.Context, query DNSTopLocationsParams, opts ...option.RequestOption) (res *DNSTopLocationsResponse, err error) {
	var env DNSTopLocationsResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/dns/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSTopAsesResponse struct {
	// Metadata for the results.
	Meta DNSTopAsesResponseMeta   `json:"meta,required"`
	Top0 []DNSTopAsesResponseTop0 `json:"top_0,required"`
	JSON dnsTopAsesResponseJSON   `json:"-"`
}

// dnsTopAsesResponseJSON contains the JSON metadata for the struct
// [DNSTopAsesResponse]
type dnsTopAsesResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopAsesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopAsesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type DNSTopAsesResponseMeta struct {
	ConfidenceInfo DNSTopAsesResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []DNSTopAsesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization DNSTopAsesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []DNSTopAsesResponseMetaUnit `json:"units,required"`
	JSON  dnsTopAsesResponseMetaJSON   `json:"-"`
}

// dnsTopAsesResponseMetaJSON contains the JSON metadata for the struct
// [DNSTopAsesResponseMeta]
type dnsTopAsesResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSTopAsesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopAsesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type DNSTopAsesResponseMetaConfidenceInfo struct {
	Annotations []DNSTopAsesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                    `json:"level,required"`
	JSON  dnsTopAsesResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsTopAsesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [DNSTopAsesResponseMetaConfidenceInfo]
type dnsTopAsesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopAsesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopAsesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type DNSTopAsesResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                    `json:"description,required"`
	EndDate     time.Time                                                 `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                               `json:"isInstantaneous,required"`
	LinkedURL       string                                             `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                          `json:"startDate,required" format:"date-time"`
	JSON            dnsTopAsesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsTopAsesResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata
// for the struct [DNSTopAsesResponseMetaConfidenceInfoAnnotation]
type dnsTopAsesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *DNSTopAsesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopAsesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceAll                DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceBGP                DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceBots               DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceCT                 DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceDNS                DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceDos                DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceFw                 DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceIQI                DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceNet                DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceAll, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceBGP, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceBots, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceCT, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceDNS, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceDos, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceFw, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceIQI, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceNet, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, DNSTopAsesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventType string

const (
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventTypeEvent             DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventTypeOutage            DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventTypePipeline          DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventTypeEvent, DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventTypeOutage, DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventTypePipeline, DNSTopAsesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type DNSTopAsesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                           `json:"startTime,required" format:"date-time"`
	JSON      dnsTopAsesResponseMetaDateRangeJSON `json:"-"`
}

// dnsTopAsesResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [DNSTopAsesResponseMetaDateRange]
type dnsTopAsesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopAsesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopAsesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type DNSTopAsesResponseMetaNormalization string

const (
	DNSTopAsesResponseMetaNormalizationPercentage           DNSTopAsesResponseMetaNormalization = "PERCENTAGE"
	DNSTopAsesResponseMetaNormalizationMin0Max              DNSTopAsesResponseMetaNormalization = "MIN0_MAX"
	DNSTopAsesResponseMetaNormalizationMinMax               DNSTopAsesResponseMetaNormalization = "MIN_MAX"
	DNSTopAsesResponseMetaNormalizationRawValues            DNSTopAsesResponseMetaNormalization = "RAW_VALUES"
	DNSTopAsesResponseMetaNormalizationPercentageChange     DNSTopAsesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	DNSTopAsesResponseMetaNormalizationRollingAverage       DNSTopAsesResponseMetaNormalization = "ROLLING_AVERAGE"
	DNSTopAsesResponseMetaNormalizationOverlappedPercentage DNSTopAsesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	DNSTopAsesResponseMetaNormalizationRatio                DNSTopAsesResponseMetaNormalization = "RATIO"
)

func (r DNSTopAsesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case DNSTopAsesResponseMetaNormalizationPercentage, DNSTopAsesResponseMetaNormalizationMin0Max, DNSTopAsesResponseMetaNormalizationMinMax, DNSTopAsesResponseMetaNormalizationRawValues, DNSTopAsesResponseMetaNormalizationPercentageChange, DNSTopAsesResponseMetaNormalizationRollingAverage, DNSTopAsesResponseMetaNormalizationOverlappedPercentage, DNSTopAsesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type DNSTopAsesResponseMetaUnit struct {
	Name  string                         `json:"name,required"`
	Value string                         `json:"value,required"`
	JSON  dnsTopAsesResponseMetaUnitJSON `json:"-"`
}

// dnsTopAsesResponseMetaUnitJSON contains the JSON metadata for the struct
// [DNSTopAsesResponseMetaUnit]
type dnsTopAsesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopAsesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopAsesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type DNSTopAsesResponseTop0 struct {
	ClientASN    int64  `json:"clientASN,required"`
	ClientAsName string `json:"clientASName,required"`
	// A numeric string.
	Value string                     `json:"value,required"`
	JSON  dnsTopAsesResponseTop0JSON `json:"-"`
}

// dnsTopAsesResponseTop0JSON contains the JSON metadata for the struct
// [DNSTopAsesResponseTop0]
type dnsTopAsesResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSTopAsesResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopAsesResponseTop0JSON) RawJSON() string {
	return r.raw
}

type DNSTopLocationsResponse struct {
	// Metadata for the results.
	Meta DNSTopLocationsResponseMeta   `json:"meta,required"`
	Top0 []DNSTopLocationsResponseTop0 `json:"top_0,required"`
	JSON dnsTopLocationsResponseJSON   `json:"-"`
}

// dnsTopLocationsResponseJSON contains the JSON metadata for the struct
// [DNSTopLocationsResponse]
type dnsTopLocationsResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopLocationsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type DNSTopLocationsResponseMeta struct {
	ConfidenceInfo DNSTopLocationsResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []DNSTopLocationsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization DNSTopLocationsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []DNSTopLocationsResponseMetaUnit `json:"units,required"`
	JSON  dnsTopLocationsResponseMetaJSON   `json:"-"`
}

// dnsTopLocationsResponseMetaJSON contains the JSON metadata for the struct
// [DNSTopLocationsResponseMeta]
type dnsTopLocationsResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSTopLocationsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopLocationsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type DNSTopLocationsResponseMetaConfidenceInfo struct {
	Annotations []DNSTopLocationsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                         `json:"level,required"`
	JSON  dnsTopLocationsResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsTopLocationsResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [DNSTopLocationsResponseMetaConfidenceInfo]
type dnsTopLocationsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopLocationsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopLocationsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type DNSTopLocationsResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                         `json:"description,required"`
	EndDate     time.Time                                                      `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                               `json:"startDate,required" format:"date-time"`
	JSON            dnsTopLocationsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsTopLocationsResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [DNSTopLocationsResponseMetaConfidenceInfoAnnotation]
type dnsTopLocationsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *DNSTopLocationsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopLocationsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceAll                DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceBGP                DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceBots               DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceCT                 DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDNS                DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDos                DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceFw                 DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceIQI                DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceNet                DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceAll, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceBGP, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceBots, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceCT, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDNS, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDos, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceFw, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceIQI, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceNet, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventType string

const (
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeEvent             DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeOutage            DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypePipeline          DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeEvent, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeOutage, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypePipeline, DNSTopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type DNSTopLocationsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                `json:"startTime,required" format:"date-time"`
	JSON      dnsTopLocationsResponseMetaDateRangeJSON `json:"-"`
}

// dnsTopLocationsResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [DNSTopLocationsResponseMetaDateRange]
type dnsTopLocationsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopLocationsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopLocationsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type DNSTopLocationsResponseMetaNormalization string

const (
	DNSTopLocationsResponseMetaNormalizationPercentage           DNSTopLocationsResponseMetaNormalization = "PERCENTAGE"
	DNSTopLocationsResponseMetaNormalizationMin0Max              DNSTopLocationsResponseMetaNormalization = "MIN0_MAX"
	DNSTopLocationsResponseMetaNormalizationMinMax               DNSTopLocationsResponseMetaNormalization = "MIN_MAX"
	DNSTopLocationsResponseMetaNormalizationRawValues            DNSTopLocationsResponseMetaNormalization = "RAW_VALUES"
	DNSTopLocationsResponseMetaNormalizationPercentageChange     DNSTopLocationsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	DNSTopLocationsResponseMetaNormalizationRollingAverage       DNSTopLocationsResponseMetaNormalization = "ROLLING_AVERAGE"
	DNSTopLocationsResponseMetaNormalizationOverlappedPercentage DNSTopLocationsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	DNSTopLocationsResponseMetaNormalizationRatio                DNSTopLocationsResponseMetaNormalization = "RATIO"
)

func (r DNSTopLocationsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case DNSTopLocationsResponseMetaNormalizationPercentage, DNSTopLocationsResponseMetaNormalizationMin0Max, DNSTopLocationsResponseMetaNormalizationMinMax, DNSTopLocationsResponseMetaNormalizationRawValues, DNSTopLocationsResponseMetaNormalizationPercentageChange, DNSTopLocationsResponseMetaNormalizationRollingAverage, DNSTopLocationsResponseMetaNormalizationOverlappedPercentage, DNSTopLocationsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type DNSTopLocationsResponseMetaUnit struct {
	Name  string                              `json:"name,required"`
	Value string                              `json:"value,required"`
	JSON  dnsTopLocationsResponseMetaUnitJSON `json:"-"`
}

// dnsTopLocationsResponseMetaUnitJSON contains the JSON metadata for the struct
// [DNSTopLocationsResponseMetaUnit]
type dnsTopLocationsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopLocationsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopLocationsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type DNSTopLocationsResponseTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                          `json:"value,required"`
	JSON  dnsTopLocationsResponseTop0JSON `json:"-"`
}

// dnsTopLocationsResponseTop0JSON contains the JSON metadata for the struct
// [DNSTopLocationsResponseTop0]
type dnsTopLocationsResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DNSTopLocationsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopLocationsResponseTop0JSON) RawJSON() string {
	return r.raw
}

type DNSTopAsesParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results based on cache status.
	CacheHit param.Field[[]bool] `query:"cacheHit"`
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
	// Filters results based on DNSSEC (DNS Security Extensions) support.
	DNSSEC param.Field[[]DNSTopAsesParamsDNSSEC] `query:"dnssec"`
	// Filters results based on DNSSEC (DNS Security Extensions) client awareness.
	DNSSECAware param.Field[[]DNSTopAsesParamsDNSSECAware] `query:"dnssecAware"`
	// Filters results based on DNSSEC-validated answers by end-to-end security status.
	DNSSECE2E param.Field[[]bool] `query:"dnssecE2e"`
	// Filters results by domain name.
	Domain param.Field[[]string] `query:"domain"`
	// Format in which results will be returned.
	Format param.Field[DNSTopAsesParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]DNSTopAsesParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters results based on whether the queries have a matching answer.
	MatchingAnswer param.Field[[]bool] `query:"matchingAnswer"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Specifies whether the response includes empty DNS responses (NODATA).
	Nodata param.Field[[]bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[[]DNSTopAsesParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[[]DNSTopAsesParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[[]DNSTopAsesParamsResponseCode] `query:"responseCode"`
	// Filters results by DNS response TTL.
	ResponseTTL param.Field[[]DNSTopAsesParamsResponseTTL] `query:"responseTtl"`
}

// URLQuery serializes [DNSTopAsesParams]'s query parameters as `url.Values`.
func (r DNSTopAsesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DNSTopAsesParamsDNSSEC string

const (
	DNSTopAsesParamsDNSSECInvalid  DNSTopAsesParamsDNSSEC = "INVALID"
	DNSTopAsesParamsDNSSECInsecure DNSTopAsesParamsDNSSEC = "INSECURE"
	DNSTopAsesParamsDNSSECSecure   DNSTopAsesParamsDNSSEC = "SECURE"
	DNSTopAsesParamsDNSSECOther    DNSTopAsesParamsDNSSEC = "OTHER"
)

func (r DNSTopAsesParamsDNSSEC) IsKnown() bool {
	switch r {
	case DNSTopAsesParamsDNSSECInvalid, DNSTopAsesParamsDNSSECInsecure, DNSTopAsesParamsDNSSECSecure, DNSTopAsesParamsDNSSECOther:
		return true
	}
	return false
}

type DNSTopAsesParamsDNSSECAware string

const (
	DNSTopAsesParamsDNSSECAwareSupported    DNSTopAsesParamsDNSSECAware = "SUPPORTED"
	DNSTopAsesParamsDNSSECAwareNotSupported DNSTopAsesParamsDNSSECAware = "NOT_SUPPORTED"
)

func (r DNSTopAsesParamsDNSSECAware) IsKnown() bool {
	switch r {
	case DNSTopAsesParamsDNSSECAwareSupported, DNSTopAsesParamsDNSSECAwareNotSupported:
		return true
	}
	return false
}

// Format in which results will be returned.
type DNSTopAsesParamsFormat string

const (
	DNSTopAsesParamsFormatJson DNSTopAsesParamsFormat = "JSON"
	DNSTopAsesParamsFormatCsv  DNSTopAsesParamsFormat = "CSV"
)

func (r DNSTopAsesParamsFormat) IsKnown() bool {
	switch r {
	case DNSTopAsesParamsFormatJson, DNSTopAsesParamsFormatCsv:
		return true
	}
	return false
}

type DNSTopAsesParamsIPVersion string

const (
	DNSTopAsesParamsIPVersionIPv4 DNSTopAsesParamsIPVersion = "IPv4"
	DNSTopAsesParamsIPVersionIPv6 DNSTopAsesParamsIPVersion = "IPv6"
)

func (r DNSTopAsesParamsIPVersion) IsKnown() bool {
	switch r {
	case DNSTopAsesParamsIPVersionIPv4, DNSTopAsesParamsIPVersionIPv6:
		return true
	}
	return false
}

type DNSTopAsesParamsProtocol string

const (
	DNSTopAsesParamsProtocolUdp   DNSTopAsesParamsProtocol = "UDP"
	DNSTopAsesParamsProtocolTCP   DNSTopAsesParamsProtocol = "TCP"
	DNSTopAsesParamsProtocolHTTPS DNSTopAsesParamsProtocol = "HTTPS"
	DNSTopAsesParamsProtocolTLS   DNSTopAsesParamsProtocol = "TLS"
)

func (r DNSTopAsesParamsProtocol) IsKnown() bool {
	switch r {
	case DNSTopAsesParamsProtocolUdp, DNSTopAsesParamsProtocolTCP, DNSTopAsesParamsProtocolHTTPS, DNSTopAsesParamsProtocolTLS:
		return true
	}
	return false
}

type DNSTopAsesParamsQueryType string

const (
	DNSTopAsesParamsQueryTypeA          DNSTopAsesParamsQueryType = "A"
	DNSTopAsesParamsQueryTypeAAAA       DNSTopAsesParamsQueryType = "AAAA"
	DNSTopAsesParamsQueryTypeA6         DNSTopAsesParamsQueryType = "A6"
	DNSTopAsesParamsQueryTypeAfsdb      DNSTopAsesParamsQueryType = "AFSDB"
	DNSTopAsesParamsQueryTypeAny        DNSTopAsesParamsQueryType = "ANY"
	DNSTopAsesParamsQueryTypeApl        DNSTopAsesParamsQueryType = "APL"
	DNSTopAsesParamsQueryTypeAtma       DNSTopAsesParamsQueryType = "ATMA"
	DNSTopAsesParamsQueryTypeAXFR       DNSTopAsesParamsQueryType = "AXFR"
	DNSTopAsesParamsQueryTypeCAA        DNSTopAsesParamsQueryType = "CAA"
	DNSTopAsesParamsQueryTypeCdnskey    DNSTopAsesParamsQueryType = "CDNSKEY"
	DNSTopAsesParamsQueryTypeCds        DNSTopAsesParamsQueryType = "CDS"
	DNSTopAsesParamsQueryTypeCERT       DNSTopAsesParamsQueryType = "CERT"
	DNSTopAsesParamsQueryTypeCNAME      DNSTopAsesParamsQueryType = "CNAME"
	DNSTopAsesParamsQueryTypeCsync      DNSTopAsesParamsQueryType = "CSYNC"
	DNSTopAsesParamsQueryTypeDhcid      DNSTopAsesParamsQueryType = "DHCID"
	DNSTopAsesParamsQueryTypeDlv        DNSTopAsesParamsQueryType = "DLV"
	DNSTopAsesParamsQueryTypeDname      DNSTopAsesParamsQueryType = "DNAME"
	DNSTopAsesParamsQueryTypeDNSKEY     DNSTopAsesParamsQueryType = "DNSKEY"
	DNSTopAsesParamsQueryTypeDoa        DNSTopAsesParamsQueryType = "DOA"
	DNSTopAsesParamsQueryTypeDS         DNSTopAsesParamsQueryType = "DS"
	DNSTopAsesParamsQueryTypeEid        DNSTopAsesParamsQueryType = "EID"
	DNSTopAsesParamsQueryTypeEui48      DNSTopAsesParamsQueryType = "EUI48"
	DNSTopAsesParamsQueryTypeEui64      DNSTopAsesParamsQueryType = "EUI64"
	DNSTopAsesParamsQueryTypeGpos       DNSTopAsesParamsQueryType = "GPOS"
	DNSTopAsesParamsQueryTypeGid        DNSTopAsesParamsQueryType = "GID"
	DNSTopAsesParamsQueryTypeHinfo      DNSTopAsesParamsQueryType = "HINFO"
	DNSTopAsesParamsQueryTypeHip        DNSTopAsesParamsQueryType = "HIP"
	DNSTopAsesParamsQueryTypeHTTPS      DNSTopAsesParamsQueryType = "HTTPS"
	DNSTopAsesParamsQueryTypeIpseckey   DNSTopAsesParamsQueryType = "IPSECKEY"
	DNSTopAsesParamsQueryTypeIsdn       DNSTopAsesParamsQueryType = "ISDN"
	DNSTopAsesParamsQueryTypeIxfr       DNSTopAsesParamsQueryType = "IXFR"
	DNSTopAsesParamsQueryTypeKey        DNSTopAsesParamsQueryType = "KEY"
	DNSTopAsesParamsQueryTypeKx         DNSTopAsesParamsQueryType = "KX"
	DNSTopAsesParamsQueryTypeL32        DNSTopAsesParamsQueryType = "L32"
	DNSTopAsesParamsQueryTypeL64        DNSTopAsesParamsQueryType = "L64"
	DNSTopAsesParamsQueryTypeLOC        DNSTopAsesParamsQueryType = "LOC"
	DNSTopAsesParamsQueryTypeLp         DNSTopAsesParamsQueryType = "LP"
	DNSTopAsesParamsQueryTypeMaila      DNSTopAsesParamsQueryType = "MAILA"
	DNSTopAsesParamsQueryTypeMailb      DNSTopAsesParamsQueryType = "MAILB"
	DNSTopAsesParamsQueryTypeMB         DNSTopAsesParamsQueryType = "MB"
	DNSTopAsesParamsQueryTypeMd         DNSTopAsesParamsQueryType = "MD"
	DNSTopAsesParamsQueryTypeMf         DNSTopAsesParamsQueryType = "MF"
	DNSTopAsesParamsQueryTypeMg         DNSTopAsesParamsQueryType = "MG"
	DNSTopAsesParamsQueryTypeMinfo      DNSTopAsesParamsQueryType = "MINFO"
	DNSTopAsesParamsQueryTypeMr         DNSTopAsesParamsQueryType = "MR"
	DNSTopAsesParamsQueryTypeMX         DNSTopAsesParamsQueryType = "MX"
	DNSTopAsesParamsQueryTypeNAPTR      DNSTopAsesParamsQueryType = "NAPTR"
	DNSTopAsesParamsQueryTypeNb         DNSTopAsesParamsQueryType = "NB"
	DNSTopAsesParamsQueryTypeNbstat     DNSTopAsesParamsQueryType = "NBSTAT"
	DNSTopAsesParamsQueryTypeNid        DNSTopAsesParamsQueryType = "NID"
	DNSTopAsesParamsQueryTypeNimloc     DNSTopAsesParamsQueryType = "NIMLOC"
	DNSTopAsesParamsQueryTypeNinfo      DNSTopAsesParamsQueryType = "NINFO"
	DNSTopAsesParamsQueryTypeNS         DNSTopAsesParamsQueryType = "NS"
	DNSTopAsesParamsQueryTypeNsap       DNSTopAsesParamsQueryType = "NSAP"
	DNSTopAsesParamsQueryTypeNsec       DNSTopAsesParamsQueryType = "NSEC"
	DNSTopAsesParamsQueryTypeNsec3      DNSTopAsesParamsQueryType = "NSEC3"
	DNSTopAsesParamsQueryTypeNsec3Param DNSTopAsesParamsQueryType = "NSEC3PARAM"
	DNSTopAsesParamsQueryTypeNull       DNSTopAsesParamsQueryType = "NULL"
	DNSTopAsesParamsQueryTypeNxt        DNSTopAsesParamsQueryType = "NXT"
	DNSTopAsesParamsQueryTypeOpenpgpkey DNSTopAsesParamsQueryType = "OPENPGPKEY"
	DNSTopAsesParamsQueryTypeOpt        DNSTopAsesParamsQueryType = "OPT"
	DNSTopAsesParamsQueryTypePTR        DNSTopAsesParamsQueryType = "PTR"
	DNSTopAsesParamsQueryTypePx         DNSTopAsesParamsQueryType = "PX"
	DNSTopAsesParamsQueryTypeRkey       DNSTopAsesParamsQueryType = "RKEY"
	DNSTopAsesParamsQueryTypeRp         DNSTopAsesParamsQueryType = "RP"
	DNSTopAsesParamsQueryTypeRrsig      DNSTopAsesParamsQueryType = "RRSIG"
	DNSTopAsesParamsQueryTypeRt         DNSTopAsesParamsQueryType = "RT"
	DNSTopAsesParamsQueryTypeSig        DNSTopAsesParamsQueryType = "SIG"
	DNSTopAsesParamsQueryTypeSink       DNSTopAsesParamsQueryType = "SINK"
	DNSTopAsesParamsQueryTypeSMIMEA     DNSTopAsesParamsQueryType = "SMIMEA"
	DNSTopAsesParamsQueryTypeSOA        DNSTopAsesParamsQueryType = "SOA"
	DNSTopAsesParamsQueryTypeSPF        DNSTopAsesParamsQueryType = "SPF"
	DNSTopAsesParamsQueryTypeSRV        DNSTopAsesParamsQueryType = "SRV"
	DNSTopAsesParamsQueryTypeSSHFP      DNSTopAsesParamsQueryType = "SSHFP"
	DNSTopAsesParamsQueryTypeSVCB       DNSTopAsesParamsQueryType = "SVCB"
	DNSTopAsesParamsQueryTypeTa         DNSTopAsesParamsQueryType = "TA"
	DNSTopAsesParamsQueryTypeTalink     DNSTopAsesParamsQueryType = "TALINK"
	DNSTopAsesParamsQueryTypeTkey       DNSTopAsesParamsQueryType = "TKEY"
	DNSTopAsesParamsQueryTypeTLSA       DNSTopAsesParamsQueryType = "TLSA"
	DNSTopAsesParamsQueryTypeTSIG       DNSTopAsesParamsQueryType = "TSIG"
	DNSTopAsesParamsQueryTypeTXT        DNSTopAsesParamsQueryType = "TXT"
	DNSTopAsesParamsQueryTypeUinfo      DNSTopAsesParamsQueryType = "UINFO"
	DNSTopAsesParamsQueryTypeUID        DNSTopAsesParamsQueryType = "UID"
	DNSTopAsesParamsQueryTypeUnspec     DNSTopAsesParamsQueryType = "UNSPEC"
	DNSTopAsesParamsQueryTypeURI        DNSTopAsesParamsQueryType = "URI"
	DNSTopAsesParamsQueryTypeWks        DNSTopAsesParamsQueryType = "WKS"
	DNSTopAsesParamsQueryTypeX25        DNSTopAsesParamsQueryType = "X25"
	DNSTopAsesParamsQueryTypeZonemd     DNSTopAsesParamsQueryType = "ZONEMD"
)

func (r DNSTopAsesParamsQueryType) IsKnown() bool {
	switch r {
	case DNSTopAsesParamsQueryTypeA, DNSTopAsesParamsQueryTypeAAAA, DNSTopAsesParamsQueryTypeA6, DNSTopAsesParamsQueryTypeAfsdb, DNSTopAsesParamsQueryTypeAny, DNSTopAsesParamsQueryTypeApl, DNSTopAsesParamsQueryTypeAtma, DNSTopAsesParamsQueryTypeAXFR, DNSTopAsesParamsQueryTypeCAA, DNSTopAsesParamsQueryTypeCdnskey, DNSTopAsesParamsQueryTypeCds, DNSTopAsesParamsQueryTypeCERT, DNSTopAsesParamsQueryTypeCNAME, DNSTopAsesParamsQueryTypeCsync, DNSTopAsesParamsQueryTypeDhcid, DNSTopAsesParamsQueryTypeDlv, DNSTopAsesParamsQueryTypeDname, DNSTopAsesParamsQueryTypeDNSKEY, DNSTopAsesParamsQueryTypeDoa, DNSTopAsesParamsQueryTypeDS, DNSTopAsesParamsQueryTypeEid, DNSTopAsesParamsQueryTypeEui48, DNSTopAsesParamsQueryTypeEui64, DNSTopAsesParamsQueryTypeGpos, DNSTopAsesParamsQueryTypeGid, DNSTopAsesParamsQueryTypeHinfo, DNSTopAsesParamsQueryTypeHip, DNSTopAsesParamsQueryTypeHTTPS, DNSTopAsesParamsQueryTypeIpseckey, DNSTopAsesParamsQueryTypeIsdn, DNSTopAsesParamsQueryTypeIxfr, DNSTopAsesParamsQueryTypeKey, DNSTopAsesParamsQueryTypeKx, DNSTopAsesParamsQueryTypeL32, DNSTopAsesParamsQueryTypeL64, DNSTopAsesParamsQueryTypeLOC, DNSTopAsesParamsQueryTypeLp, DNSTopAsesParamsQueryTypeMaila, DNSTopAsesParamsQueryTypeMailb, DNSTopAsesParamsQueryTypeMB, DNSTopAsesParamsQueryTypeMd, DNSTopAsesParamsQueryTypeMf, DNSTopAsesParamsQueryTypeMg, DNSTopAsesParamsQueryTypeMinfo, DNSTopAsesParamsQueryTypeMr, DNSTopAsesParamsQueryTypeMX, DNSTopAsesParamsQueryTypeNAPTR, DNSTopAsesParamsQueryTypeNb, DNSTopAsesParamsQueryTypeNbstat, DNSTopAsesParamsQueryTypeNid, DNSTopAsesParamsQueryTypeNimloc, DNSTopAsesParamsQueryTypeNinfo, DNSTopAsesParamsQueryTypeNS, DNSTopAsesParamsQueryTypeNsap, DNSTopAsesParamsQueryTypeNsec, DNSTopAsesParamsQueryTypeNsec3, DNSTopAsesParamsQueryTypeNsec3Param, DNSTopAsesParamsQueryTypeNull, DNSTopAsesParamsQueryTypeNxt, DNSTopAsesParamsQueryTypeOpenpgpkey, DNSTopAsesParamsQueryTypeOpt, DNSTopAsesParamsQueryTypePTR, DNSTopAsesParamsQueryTypePx, DNSTopAsesParamsQueryTypeRkey, DNSTopAsesParamsQueryTypeRp, DNSTopAsesParamsQueryTypeRrsig, DNSTopAsesParamsQueryTypeRt, DNSTopAsesParamsQueryTypeSig, DNSTopAsesParamsQueryTypeSink, DNSTopAsesParamsQueryTypeSMIMEA, DNSTopAsesParamsQueryTypeSOA, DNSTopAsesParamsQueryTypeSPF, DNSTopAsesParamsQueryTypeSRV, DNSTopAsesParamsQueryTypeSSHFP, DNSTopAsesParamsQueryTypeSVCB, DNSTopAsesParamsQueryTypeTa, DNSTopAsesParamsQueryTypeTalink, DNSTopAsesParamsQueryTypeTkey, DNSTopAsesParamsQueryTypeTLSA, DNSTopAsesParamsQueryTypeTSIG, DNSTopAsesParamsQueryTypeTXT, DNSTopAsesParamsQueryTypeUinfo, DNSTopAsesParamsQueryTypeUID, DNSTopAsesParamsQueryTypeUnspec, DNSTopAsesParamsQueryTypeURI, DNSTopAsesParamsQueryTypeWks, DNSTopAsesParamsQueryTypeX25, DNSTopAsesParamsQueryTypeZonemd:
		return true
	}
	return false
}

type DNSTopAsesParamsResponseCode string

const (
	DNSTopAsesParamsResponseCodeNoerror   DNSTopAsesParamsResponseCode = "NOERROR"
	DNSTopAsesParamsResponseCodeFormerr   DNSTopAsesParamsResponseCode = "FORMERR"
	DNSTopAsesParamsResponseCodeServfail  DNSTopAsesParamsResponseCode = "SERVFAIL"
	DNSTopAsesParamsResponseCodeNxdomain  DNSTopAsesParamsResponseCode = "NXDOMAIN"
	DNSTopAsesParamsResponseCodeNotimp    DNSTopAsesParamsResponseCode = "NOTIMP"
	DNSTopAsesParamsResponseCodeRefused   DNSTopAsesParamsResponseCode = "REFUSED"
	DNSTopAsesParamsResponseCodeYxdomain  DNSTopAsesParamsResponseCode = "YXDOMAIN"
	DNSTopAsesParamsResponseCodeYxrrset   DNSTopAsesParamsResponseCode = "YXRRSET"
	DNSTopAsesParamsResponseCodeNxrrset   DNSTopAsesParamsResponseCode = "NXRRSET"
	DNSTopAsesParamsResponseCodeNotauth   DNSTopAsesParamsResponseCode = "NOTAUTH"
	DNSTopAsesParamsResponseCodeNotzone   DNSTopAsesParamsResponseCode = "NOTZONE"
	DNSTopAsesParamsResponseCodeBadsig    DNSTopAsesParamsResponseCode = "BADSIG"
	DNSTopAsesParamsResponseCodeBadkey    DNSTopAsesParamsResponseCode = "BADKEY"
	DNSTopAsesParamsResponseCodeBadtime   DNSTopAsesParamsResponseCode = "BADTIME"
	DNSTopAsesParamsResponseCodeBadmode   DNSTopAsesParamsResponseCode = "BADMODE"
	DNSTopAsesParamsResponseCodeBadname   DNSTopAsesParamsResponseCode = "BADNAME"
	DNSTopAsesParamsResponseCodeBadalg    DNSTopAsesParamsResponseCode = "BADALG"
	DNSTopAsesParamsResponseCodeBadtrunc  DNSTopAsesParamsResponseCode = "BADTRUNC"
	DNSTopAsesParamsResponseCodeBadcookie DNSTopAsesParamsResponseCode = "BADCOOKIE"
)

func (r DNSTopAsesParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSTopAsesParamsResponseCodeNoerror, DNSTopAsesParamsResponseCodeFormerr, DNSTopAsesParamsResponseCodeServfail, DNSTopAsesParamsResponseCodeNxdomain, DNSTopAsesParamsResponseCodeNotimp, DNSTopAsesParamsResponseCodeRefused, DNSTopAsesParamsResponseCodeYxdomain, DNSTopAsesParamsResponseCodeYxrrset, DNSTopAsesParamsResponseCodeNxrrset, DNSTopAsesParamsResponseCodeNotauth, DNSTopAsesParamsResponseCodeNotzone, DNSTopAsesParamsResponseCodeBadsig, DNSTopAsesParamsResponseCodeBadkey, DNSTopAsesParamsResponseCodeBadtime, DNSTopAsesParamsResponseCodeBadmode, DNSTopAsesParamsResponseCodeBadname, DNSTopAsesParamsResponseCodeBadalg, DNSTopAsesParamsResponseCodeBadtrunc, DNSTopAsesParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSTopAsesParamsResponseTTL string

const (
	DNSTopAsesParamsResponseTTLLte1M      DNSTopAsesParamsResponseTTL = "LTE_1M"
	DNSTopAsesParamsResponseTTLGt1MLte5M  DNSTopAsesParamsResponseTTL = "GT_1M_LTE_5M"
	DNSTopAsesParamsResponseTTLGt5MLte15M DNSTopAsesParamsResponseTTL = "GT_5M_LTE_15M"
	DNSTopAsesParamsResponseTTLGt15MLte1H DNSTopAsesParamsResponseTTL = "GT_15M_LTE_1H"
	DNSTopAsesParamsResponseTTLGt1HLte1D  DNSTopAsesParamsResponseTTL = "GT_1H_LTE_1D"
	DNSTopAsesParamsResponseTTLGt1DLte1W  DNSTopAsesParamsResponseTTL = "GT_1D_LTE_1W"
	DNSTopAsesParamsResponseTTLGt1W       DNSTopAsesParamsResponseTTL = "GT_1W"
)

func (r DNSTopAsesParamsResponseTTL) IsKnown() bool {
	switch r {
	case DNSTopAsesParamsResponseTTLLte1M, DNSTopAsesParamsResponseTTLGt1MLte5M, DNSTopAsesParamsResponseTTLGt5MLte15M, DNSTopAsesParamsResponseTTLGt15MLte1H, DNSTopAsesParamsResponseTTLGt1HLte1D, DNSTopAsesParamsResponseTTLGt1DLte1W, DNSTopAsesParamsResponseTTLGt1W:
		return true
	}
	return false
}

type DNSTopAsesResponseEnvelope struct {
	Result  DNSTopAsesResponse             `json:"result,required"`
	Success bool                           `json:"success,required"`
	JSON    dnsTopAsesResponseEnvelopeJSON `json:"-"`
}

// dnsTopAsesResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSTopAsesResponseEnvelope]
type dnsTopAsesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopAsesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopAsesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSTopLocationsParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results based on cache status.
	CacheHit param.Field[[]bool] `query:"cacheHit"`
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
	// Filters results based on DNSSEC (DNS Security Extensions) support.
	DNSSEC param.Field[[]DNSTopLocationsParamsDNSSEC] `query:"dnssec"`
	// Filters results based on DNSSEC (DNS Security Extensions) client awareness.
	DNSSECAware param.Field[[]DNSTopLocationsParamsDNSSECAware] `query:"dnssecAware"`
	// Filters results based on DNSSEC-validated answers by end-to-end security status.
	DNSSECE2E param.Field[[]bool] `query:"dnssecE2e"`
	// Filters results by domain name.
	Domain param.Field[[]string] `query:"domain"`
	// Format in which results will be returned.
	Format param.Field[DNSTopLocationsParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]DNSTopLocationsParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters results based on whether the queries have a matching answer.
	MatchingAnswer param.Field[[]bool] `query:"matchingAnswer"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Specifies whether the response includes empty DNS responses (NODATA).
	Nodata param.Field[[]bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[[]DNSTopLocationsParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[[]DNSTopLocationsParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[[]DNSTopLocationsParamsResponseCode] `query:"responseCode"`
	// Filters results by DNS response TTL.
	ResponseTTL param.Field[[]DNSTopLocationsParamsResponseTTL] `query:"responseTtl"`
	// Filters results by top-level domain.
	TLD param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTopLocationsParams]'s query parameters as `url.Values`.
func (r DNSTopLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DNSTopLocationsParamsDNSSEC string

const (
	DNSTopLocationsParamsDNSSECInvalid  DNSTopLocationsParamsDNSSEC = "INVALID"
	DNSTopLocationsParamsDNSSECInsecure DNSTopLocationsParamsDNSSEC = "INSECURE"
	DNSTopLocationsParamsDNSSECSecure   DNSTopLocationsParamsDNSSEC = "SECURE"
	DNSTopLocationsParamsDNSSECOther    DNSTopLocationsParamsDNSSEC = "OTHER"
)

func (r DNSTopLocationsParamsDNSSEC) IsKnown() bool {
	switch r {
	case DNSTopLocationsParamsDNSSECInvalid, DNSTopLocationsParamsDNSSECInsecure, DNSTopLocationsParamsDNSSECSecure, DNSTopLocationsParamsDNSSECOther:
		return true
	}
	return false
}

type DNSTopLocationsParamsDNSSECAware string

const (
	DNSTopLocationsParamsDNSSECAwareSupported    DNSTopLocationsParamsDNSSECAware = "SUPPORTED"
	DNSTopLocationsParamsDNSSECAwareNotSupported DNSTopLocationsParamsDNSSECAware = "NOT_SUPPORTED"
)

func (r DNSTopLocationsParamsDNSSECAware) IsKnown() bool {
	switch r {
	case DNSTopLocationsParamsDNSSECAwareSupported, DNSTopLocationsParamsDNSSECAwareNotSupported:
		return true
	}
	return false
}

// Format in which results will be returned.
type DNSTopLocationsParamsFormat string

const (
	DNSTopLocationsParamsFormatJson DNSTopLocationsParamsFormat = "JSON"
	DNSTopLocationsParamsFormatCsv  DNSTopLocationsParamsFormat = "CSV"
)

func (r DNSTopLocationsParamsFormat) IsKnown() bool {
	switch r {
	case DNSTopLocationsParamsFormatJson, DNSTopLocationsParamsFormatCsv:
		return true
	}
	return false
}

type DNSTopLocationsParamsIPVersion string

const (
	DNSTopLocationsParamsIPVersionIPv4 DNSTopLocationsParamsIPVersion = "IPv4"
	DNSTopLocationsParamsIPVersionIPv6 DNSTopLocationsParamsIPVersion = "IPv6"
)

func (r DNSTopLocationsParamsIPVersion) IsKnown() bool {
	switch r {
	case DNSTopLocationsParamsIPVersionIPv4, DNSTopLocationsParamsIPVersionIPv6:
		return true
	}
	return false
}

type DNSTopLocationsParamsProtocol string

const (
	DNSTopLocationsParamsProtocolUdp   DNSTopLocationsParamsProtocol = "UDP"
	DNSTopLocationsParamsProtocolTCP   DNSTopLocationsParamsProtocol = "TCP"
	DNSTopLocationsParamsProtocolHTTPS DNSTopLocationsParamsProtocol = "HTTPS"
	DNSTopLocationsParamsProtocolTLS   DNSTopLocationsParamsProtocol = "TLS"
)

func (r DNSTopLocationsParamsProtocol) IsKnown() bool {
	switch r {
	case DNSTopLocationsParamsProtocolUdp, DNSTopLocationsParamsProtocolTCP, DNSTopLocationsParamsProtocolHTTPS, DNSTopLocationsParamsProtocolTLS:
		return true
	}
	return false
}

type DNSTopLocationsParamsQueryType string

const (
	DNSTopLocationsParamsQueryTypeA          DNSTopLocationsParamsQueryType = "A"
	DNSTopLocationsParamsQueryTypeAAAA       DNSTopLocationsParamsQueryType = "AAAA"
	DNSTopLocationsParamsQueryTypeA6         DNSTopLocationsParamsQueryType = "A6"
	DNSTopLocationsParamsQueryTypeAfsdb      DNSTopLocationsParamsQueryType = "AFSDB"
	DNSTopLocationsParamsQueryTypeAny        DNSTopLocationsParamsQueryType = "ANY"
	DNSTopLocationsParamsQueryTypeApl        DNSTopLocationsParamsQueryType = "APL"
	DNSTopLocationsParamsQueryTypeAtma       DNSTopLocationsParamsQueryType = "ATMA"
	DNSTopLocationsParamsQueryTypeAXFR       DNSTopLocationsParamsQueryType = "AXFR"
	DNSTopLocationsParamsQueryTypeCAA        DNSTopLocationsParamsQueryType = "CAA"
	DNSTopLocationsParamsQueryTypeCdnskey    DNSTopLocationsParamsQueryType = "CDNSKEY"
	DNSTopLocationsParamsQueryTypeCds        DNSTopLocationsParamsQueryType = "CDS"
	DNSTopLocationsParamsQueryTypeCERT       DNSTopLocationsParamsQueryType = "CERT"
	DNSTopLocationsParamsQueryTypeCNAME      DNSTopLocationsParamsQueryType = "CNAME"
	DNSTopLocationsParamsQueryTypeCsync      DNSTopLocationsParamsQueryType = "CSYNC"
	DNSTopLocationsParamsQueryTypeDhcid      DNSTopLocationsParamsQueryType = "DHCID"
	DNSTopLocationsParamsQueryTypeDlv        DNSTopLocationsParamsQueryType = "DLV"
	DNSTopLocationsParamsQueryTypeDname      DNSTopLocationsParamsQueryType = "DNAME"
	DNSTopLocationsParamsQueryTypeDNSKEY     DNSTopLocationsParamsQueryType = "DNSKEY"
	DNSTopLocationsParamsQueryTypeDoa        DNSTopLocationsParamsQueryType = "DOA"
	DNSTopLocationsParamsQueryTypeDS         DNSTopLocationsParamsQueryType = "DS"
	DNSTopLocationsParamsQueryTypeEid        DNSTopLocationsParamsQueryType = "EID"
	DNSTopLocationsParamsQueryTypeEui48      DNSTopLocationsParamsQueryType = "EUI48"
	DNSTopLocationsParamsQueryTypeEui64      DNSTopLocationsParamsQueryType = "EUI64"
	DNSTopLocationsParamsQueryTypeGpos       DNSTopLocationsParamsQueryType = "GPOS"
	DNSTopLocationsParamsQueryTypeGid        DNSTopLocationsParamsQueryType = "GID"
	DNSTopLocationsParamsQueryTypeHinfo      DNSTopLocationsParamsQueryType = "HINFO"
	DNSTopLocationsParamsQueryTypeHip        DNSTopLocationsParamsQueryType = "HIP"
	DNSTopLocationsParamsQueryTypeHTTPS      DNSTopLocationsParamsQueryType = "HTTPS"
	DNSTopLocationsParamsQueryTypeIpseckey   DNSTopLocationsParamsQueryType = "IPSECKEY"
	DNSTopLocationsParamsQueryTypeIsdn       DNSTopLocationsParamsQueryType = "ISDN"
	DNSTopLocationsParamsQueryTypeIxfr       DNSTopLocationsParamsQueryType = "IXFR"
	DNSTopLocationsParamsQueryTypeKey        DNSTopLocationsParamsQueryType = "KEY"
	DNSTopLocationsParamsQueryTypeKx         DNSTopLocationsParamsQueryType = "KX"
	DNSTopLocationsParamsQueryTypeL32        DNSTopLocationsParamsQueryType = "L32"
	DNSTopLocationsParamsQueryTypeL64        DNSTopLocationsParamsQueryType = "L64"
	DNSTopLocationsParamsQueryTypeLOC        DNSTopLocationsParamsQueryType = "LOC"
	DNSTopLocationsParamsQueryTypeLp         DNSTopLocationsParamsQueryType = "LP"
	DNSTopLocationsParamsQueryTypeMaila      DNSTopLocationsParamsQueryType = "MAILA"
	DNSTopLocationsParamsQueryTypeMailb      DNSTopLocationsParamsQueryType = "MAILB"
	DNSTopLocationsParamsQueryTypeMB         DNSTopLocationsParamsQueryType = "MB"
	DNSTopLocationsParamsQueryTypeMd         DNSTopLocationsParamsQueryType = "MD"
	DNSTopLocationsParamsQueryTypeMf         DNSTopLocationsParamsQueryType = "MF"
	DNSTopLocationsParamsQueryTypeMg         DNSTopLocationsParamsQueryType = "MG"
	DNSTopLocationsParamsQueryTypeMinfo      DNSTopLocationsParamsQueryType = "MINFO"
	DNSTopLocationsParamsQueryTypeMr         DNSTopLocationsParamsQueryType = "MR"
	DNSTopLocationsParamsQueryTypeMX         DNSTopLocationsParamsQueryType = "MX"
	DNSTopLocationsParamsQueryTypeNAPTR      DNSTopLocationsParamsQueryType = "NAPTR"
	DNSTopLocationsParamsQueryTypeNb         DNSTopLocationsParamsQueryType = "NB"
	DNSTopLocationsParamsQueryTypeNbstat     DNSTopLocationsParamsQueryType = "NBSTAT"
	DNSTopLocationsParamsQueryTypeNid        DNSTopLocationsParamsQueryType = "NID"
	DNSTopLocationsParamsQueryTypeNimloc     DNSTopLocationsParamsQueryType = "NIMLOC"
	DNSTopLocationsParamsQueryTypeNinfo      DNSTopLocationsParamsQueryType = "NINFO"
	DNSTopLocationsParamsQueryTypeNS         DNSTopLocationsParamsQueryType = "NS"
	DNSTopLocationsParamsQueryTypeNsap       DNSTopLocationsParamsQueryType = "NSAP"
	DNSTopLocationsParamsQueryTypeNsec       DNSTopLocationsParamsQueryType = "NSEC"
	DNSTopLocationsParamsQueryTypeNsec3      DNSTopLocationsParamsQueryType = "NSEC3"
	DNSTopLocationsParamsQueryTypeNsec3Param DNSTopLocationsParamsQueryType = "NSEC3PARAM"
	DNSTopLocationsParamsQueryTypeNull       DNSTopLocationsParamsQueryType = "NULL"
	DNSTopLocationsParamsQueryTypeNxt        DNSTopLocationsParamsQueryType = "NXT"
	DNSTopLocationsParamsQueryTypeOpenpgpkey DNSTopLocationsParamsQueryType = "OPENPGPKEY"
	DNSTopLocationsParamsQueryTypeOpt        DNSTopLocationsParamsQueryType = "OPT"
	DNSTopLocationsParamsQueryTypePTR        DNSTopLocationsParamsQueryType = "PTR"
	DNSTopLocationsParamsQueryTypePx         DNSTopLocationsParamsQueryType = "PX"
	DNSTopLocationsParamsQueryTypeRkey       DNSTopLocationsParamsQueryType = "RKEY"
	DNSTopLocationsParamsQueryTypeRp         DNSTopLocationsParamsQueryType = "RP"
	DNSTopLocationsParamsQueryTypeRrsig      DNSTopLocationsParamsQueryType = "RRSIG"
	DNSTopLocationsParamsQueryTypeRt         DNSTopLocationsParamsQueryType = "RT"
	DNSTopLocationsParamsQueryTypeSig        DNSTopLocationsParamsQueryType = "SIG"
	DNSTopLocationsParamsQueryTypeSink       DNSTopLocationsParamsQueryType = "SINK"
	DNSTopLocationsParamsQueryTypeSMIMEA     DNSTopLocationsParamsQueryType = "SMIMEA"
	DNSTopLocationsParamsQueryTypeSOA        DNSTopLocationsParamsQueryType = "SOA"
	DNSTopLocationsParamsQueryTypeSPF        DNSTopLocationsParamsQueryType = "SPF"
	DNSTopLocationsParamsQueryTypeSRV        DNSTopLocationsParamsQueryType = "SRV"
	DNSTopLocationsParamsQueryTypeSSHFP      DNSTopLocationsParamsQueryType = "SSHFP"
	DNSTopLocationsParamsQueryTypeSVCB       DNSTopLocationsParamsQueryType = "SVCB"
	DNSTopLocationsParamsQueryTypeTa         DNSTopLocationsParamsQueryType = "TA"
	DNSTopLocationsParamsQueryTypeTalink     DNSTopLocationsParamsQueryType = "TALINK"
	DNSTopLocationsParamsQueryTypeTkey       DNSTopLocationsParamsQueryType = "TKEY"
	DNSTopLocationsParamsQueryTypeTLSA       DNSTopLocationsParamsQueryType = "TLSA"
	DNSTopLocationsParamsQueryTypeTSIG       DNSTopLocationsParamsQueryType = "TSIG"
	DNSTopLocationsParamsQueryTypeTXT        DNSTopLocationsParamsQueryType = "TXT"
	DNSTopLocationsParamsQueryTypeUinfo      DNSTopLocationsParamsQueryType = "UINFO"
	DNSTopLocationsParamsQueryTypeUID        DNSTopLocationsParamsQueryType = "UID"
	DNSTopLocationsParamsQueryTypeUnspec     DNSTopLocationsParamsQueryType = "UNSPEC"
	DNSTopLocationsParamsQueryTypeURI        DNSTopLocationsParamsQueryType = "URI"
	DNSTopLocationsParamsQueryTypeWks        DNSTopLocationsParamsQueryType = "WKS"
	DNSTopLocationsParamsQueryTypeX25        DNSTopLocationsParamsQueryType = "X25"
	DNSTopLocationsParamsQueryTypeZonemd     DNSTopLocationsParamsQueryType = "ZONEMD"
)

func (r DNSTopLocationsParamsQueryType) IsKnown() bool {
	switch r {
	case DNSTopLocationsParamsQueryTypeA, DNSTopLocationsParamsQueryTypeAAAA, DNSTopLocationsParamsQueryTypeA6, DNSTopLocationsParamsQueryTypeAfsdb, DNSTopLocationsParamsQueryTypeAny, DNSTopLocationsParamsQueryTypeApl, DNSTopLocationsParamsQueryTypeAtma, DNSTopLocationsParamsQueryTypeAXFR, DNSTopLocationsParamsQueryTypeCAA, DNSTopLocationsParamsQueryTypeCdnskey, DNSTopLocationsParamsQueryTypeCds, DNSTopLocationsParamsQueryTypeCERT, DNSTopLocationsParamsQueryTypeCNAME, DNSTopLocationsParamsQueryTypeCsync, DNSTopLocationsParamsQueryTypeDhcid, DNSTopLocationsParamsQueryTypeDlv, DNSTopLocationsParamsQueryTypeDname, DNSTopLocationsParamsQueryTypeDNSKEY, DNSTopLocationsParamsQueryTypeDoa, DNSTopLocationsParamsQueryTypeDS, DNSTopLocationsParamsQueryTypeEid, DNSTopLocationsParamsQueryTypeEui48, DNSTopLocationsParamsQueryTypeEui64, DNSTopLocationsParamsQueryTypeGpos, DNSTopLocationsParamsQueryTypeGid, DNSTopLocationsParamsQueryTypeHinfo, DNSTopLocationsParamsQueryTypeHip, DNSTopLocationsParamsQueryTypeHTTPS, DNSTopLocationsParamsQueryTypeIpseckey, DNSTopLocationsParamsQueryTypeIsdn, DNSTopLocationsParamsQueryTypeIxfr, DNSTopLocationsParamsQueryTypeKey, DNSTopLocationsParamsQueryTypeKx, DNSTopLocationsParamsQueryTypeL32, DNSTopLocationsParamsQueryTypeL64, DNSTopLocationsParamsQueryTypeLOC, DNSTopLocationsParamsQueryTypeLp, DNSTopLocationsParamsQueryTypeMaila, DNSTopLocationsParamsQueryTypeMailb, DNSTopLocationsParamsQueryTypeMB, DNSTopLocationsParamsQueryTypeMd, DNSTopLocationsParamsQueryTypeMf, DNSTopLocationsParamsQueryTypeMg, DNSTopLocationsParamsQueryTypeMinfo, DNSTopLocationsParamsQueryTypeMr, DNSTopLocationsParamsQueryTypeMX, DNSTopLocationsParamsQueryTypeNAPTR, DNSTopLocationsParamsQueryTypeNb, DNSTopLocationsParamsQueryTypeNbstat, DNSTopLocationsParamsQueryTypeNid, DNSTopLocationsParamsQueryTypeNimloc, DNSTopLocationsParamsQueryTypeNinfo, DNSTopLocationsParamsQueryTypeNS, DNSTopLocationsParamsQueryTypeNsap, DNSTopLocationsParamsQueryTypeNsec, DNSTopLocationsParamsQueryTypeNsec3, DNSTopLocationsParamsQueryTypeNsec3Param, DNSTopLocationsParamsQueryTypeNull, DNSTopLocationsParamsQueryTypeNxt, DNSTopLocationsParamsQueryTypeOpenpgpkey, DNSTopLocationsParamsQueryTypeOpt, DNSTopLocationsParamsQueryTypePTR, DNSTopLocationsParamsQueryTypePx, DNSTopLocationsParamsQueryTypeRkey, DNSTopLocationsParamsQueryTypeRp, DNSTopLocationsParamsQueryTypeRrsig, DNSTopLocationsParamsQueryTypeRt, DNSTopLocationsParamsQueryTypeSig, DNSTopLocationsParamsQueryTypeSink, DNSTopLocationsParamsQueryTypeSMIMEA, DNSTopLocationsParamsQueryTypeSOA, DNSTopLocationsParamsQueryTypeSPF, DNSTopLocationsParamsQueryTypeSRV, DNSTopLocationsParamsQueryTypeSSHFP, DNSTopLocationsParamsQueryTypeSVCB, DNSTopLocationsParamsQueryTypeTa, DNSTopLocationsParamsQueryTypeTalink, DNSTopLocationsParamsQueryTypeTkey, DNSTopLocationsParamsQueryTypeTLSA, DNSTopLocationsParamsQueryTypeTSIG, DNSTopLocationsParamsQueryTypeTXT, DNSTopLocationsParamsQueryTypeUinfo, DNSTopLocationsParamsQueryTypeUID, DNSTopLocationsParamsQueryTypeUnspec, DNSTopLocationsParamsQueryTypeURI, DNSTopLocationsParamsQueryTypeWks, DNSTopLocationsParamsQueryTypeX25, DNSTopLocationsParamsQueryTypeZonemd:
		return true
	}
	return false
}

type DNSTopLocationsParamsResponseCode string

const (
	DNSTopLocationsParamsResponseCodeNoerror   DNSTopLocationsParamsResponseCode = "NOERROR"
	DNSTopLocationsParamsResponseCodeFormerr   DNSTopLocationsParamsResponseCode = "FORMERR"
	DNSTopLocationsParamsResponseCodeServfail  DNSTopLocationsParamsResponseCode = "SERVFAIL"
	DNSTopLocationsParamsResponseCodeNxdomain  DNSTopLocationsParamsResponseCode = "NXDOMAIN"
	DNSTopLocationsParamsResponseCodeNotimp    DNSTopLocationsParamsResponseCode = "NOTIMP"
	DNSTopLocationsParamsResponseCodeRefused   DNSTopLocationsParamsResponseCode = "REFUSED"
	DNSTopLocationsParamsResponseCodeYxdomain  DNSTopLocationsParamsResponseCode = "YXDOMAIN"
	DNSTopLocationsParamsResponseCodeYxrrset   DNSTopLocationsParamsResponseCode = "YXRRSET"
	DNSTopLocationsParamsResponseCodeNxrrset   DNSTopLocationsParamsResponseCode = "NXRRSET"
	DNSTopLocationsParamsResponseCodeNotauth   DNSTopLocationsParamsResponseCode = "NOTAUTH"
	DNSTopLocationsParamsResponseCodeNotzone   DNSTopLocationsParamsResponseCode = "NOTZONE"
	DNSTopLocationsParamsResponseCodeBadsig    DNSTopLocationsParamsResponseCode = "BADSIG"
	DNSTopLocationsParamsResponseCodeBadkey    DNSTopLocationsParamsResponseCode = "BADKEY"
	DNSTopLocationsParamsResponseCodeBadtime   DNSTopLocationsParamsResponseCode = "BADTIME"
	DNSTopLocationsParamsResponseCodeBadmode   DNSTopLocationsParamsResponseCode = "BADMODE"
	DNSTopLocationsParamsResponseCodeBadname   DNSTopLocationsParamsResponseCode = "BADNAME"
	DNSTopLocationsParamsResponseCodeBadalg    DNSTopLocationsParamsResponseCode = "BADALG"
	DNSTopLocationsParamsResponseCodeBadtrunc  DNSTopLocationsParamsResponseCode = "BADTRUNC"
	DNSTopLocationsParamsResponseCodeBadcookie DNSTopLocationsParamsResponseCode = "BADCOOKIE"
)

func (r DNSTopLocationsParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSTopLocationsParamsResponseCodeNoerror, DNSTopLocationsParamsResponseCodeFormerr, DNSTopLocationsParamsResponseCodeServfail, DNSTopLocationsParamsResponseCodeNxdomain, DNSTopLocationsParamsResponseCodeNotimp, DNSTopLocationsParamsResponseCodeRefused, DNSTopLocationsParamsResponseCodeYxdomain, DNSTopLocationsParamsResponseCodeYxrrset, DNSTopLocationsParamsResponseCodeNxrrset, DNSTopLocationsParamsResponseCodeNotauth, DNSTopLocationsParamsResponseCodeNotzone, DNSTopLocationsParamsResponseCodeBadsig, DNSTopLocationsParamsResponseCodeBadkey, DNSTopLocationsParamsResponseCodeBadtime, DNSTopLocationsParamsResponseCodeBadmode, DNSTopLocationsParamsResponseCodeBadname, DNSTopLocationsParamsResponseCodeBadalg, DNSTopLocationsParamsResponseCodeBadtrunc, DNSTopLocationsParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSTopLocationsParamsResponseTTL string

const (
	DNSTopLocationsParamsResponseTTLLte1M      DNSTopLocationsParamsResponseTTL = "LTE_1M"
	DNSTopLocationsParamsResponseTTLGt1MLte5M  DNSTopLocationsParamsResponseTTL = "GT_1M_LTE_5M"
	DNSTopLocationsParamsResponseTTLGt5MLte15M DNSTopLocationsParamsResponseTTL = "GT_5M_LTE_15M"
	DNSTopLocationsParamsResponseTTLGt15MLte1H DNSTopLocationsParamsResponseTTL = "GT_15M_LTE_1H"
	DNSTopLocationsParamsResponseTTLGt1HLte1D  DNSTopLocationsParamsResponseTTL = "GT_1H_LTE_1D"
	DNSTopLocationsParamsResponseTTLGt1DLte1W  DNSTopLocationsParamsResponseTTL = "GT_1D_LTE_1W"
	DNSTopLocationsParamsResponseTTLGt1W       DNSTopLocationsParamsResponseTTL = "GT_1W"
)

func (r DNSTopLocationsParamsResponseTTL) IsKnown() bool {
	switch r {
	case DNSTopLocationsParamsResponseTTLLte1M, DNSTopLocationsParamsResponseTTLGt1MLte5M, DNSTopLocationsParamsResponseTTLGt5MLte15M, DNSTopLocationsParamsResponseTTLGt15MLte1H, DNSTopLocationsParamsResponseTTLGt1HLte1D, DNSTopLocationsParamsResponseTTLGt1DLte1W, DNSTopLocationsParamsResponseTTLGt1W:
		return true
	}
	return false
}

type DNSTopLocationsResponseEnvelope struct {
	Result  DNSTopLocationsResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    dnsTopLocationsResponseEnvelopeJSON `json:"-"`
}

// dnsTopLocationsResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSTopLocationsResponseEnvelope]
type dnsTopLocationsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopLocationsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
