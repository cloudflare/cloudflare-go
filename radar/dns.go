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

// DNSService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSService] method instead.
type DNSService struct {
	Options          []option.RequestOption
	Top              *DNSTopService
	Summary          *DNSSummaryService
	TimeseriesGroups *DNSTimeseriesGroupService
}

// NewDNSService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDNSService(opts ...option.RequestOption) (r *DNSService) {
	r = &DNSService{}
	r.Options = opts
	r.Top = NewDNSTopService(opts...)
	r.Summary = NewDNSSummaryService(opts...)
	r.TimeseriesGroups = NewDNSTimeseriesGroupService(opts...)
	return
}

// Retrieves normalized query volume to the 1.1.1.1 DNS resolver over time.
func (r *DNSService) Timeseries(ctx context.Context, query DNSTimeseriesParams, opts ...option.RequestOption) (res *DNSTimeseriesResponse, err error) {
	var env DNSTimeseriesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/dns/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type DNSTimeseriesResponse struct {
	// Metadata for the results.
	Meta        DNSTimeseriesResponseMeta        `json:"meta" api:"required"`
	ExtraFields map[string]DNSTimeseriesResponse `json:"-" api:"extrafields"`
	JSON        dnsTimeseriesResponseJSON        `json:"-"`
}

// dnsTimeseriesResponseJSON contains the JSON metadata for the struct
// [DNSTimeseriesResponse]
type dnsTimeseriesResponseJSON struct {
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type DNSTimeseriesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    DNSTimeseriesResponseMetaAggInterval    `json:"aggInterval" api:"required"`
	ConfidenceInfo DNSTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo" api:"required"`
	DateRange      []DNSTimeseriesResponseMetaDateRange    `json:"dateRange" api:"required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated" api:"required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization DNSTimeseriesResponseMetaNormalization `json:"normalization" api:"required"`
	// Measurement units for the results.
	Units []DNSTimeseriesResponseMetaUnit `json:"units" api:"required"`
	JSON  dnsTimeseriesResponseMetaJSON   `json:"-"`
}

// dnsTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [DNSTimeseriesResponseMeta]
type dnsTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesResponseMetaAggInterval string

const (
	DNSTimeseriesResponseMetaAggIntervalFifteenMinutes DNSTimeseriesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	DNSTimeseriesResponseMetaAggIntervalOneHour        DNSTimeseriesResponseMetaAggInterval = "ONE_HOUR"
	DNSTimeseriesResponseMetaAggIntervalOneDay         DNSTimeseriesResponseMetaAggInterval = "ONE_DAY"
	DNSTimeseriesResponseMetaAggIntervalOneWeek        DNSTimeseriesResponseMetaAggInterval = "ONE_WEEK"
	DNSTimeseriesResponseMetaAggIntervalOneMonth       DNSTimeseriesResponseMetaAggInterval = "ONE_MONTH"
)

func (r DNSTimeseriesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case DNSTimeseriesResponseMetaAggIntervalFifteenMinutes, DNSTimeseriesResponseMetaAggIntervalOneHour, DNSTimeseriesResponseMetaAggIntervalOneDay, DNSTimeseriesResponseMetaAggIntervalOneWeek, DNSTimeseriesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type DNSTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []DNSTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations" api:"required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                       `json:"level" api:"required"`
	JSON  dnsTimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsTimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [DNSTimeseriesResponseMetaConfidenceInfo]
type dnsTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type DNSTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource" api:"required"`
	Description string                                                       `json:"description" api:"required"`
	EndDate     time.Time                                                    `json:"endDate" api:"required" format:"date-time"`
	// Event type for annotations.
	EventType DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType" api:"required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                  `json:"isInstantaneous" api:"required"`
	LinkedURL       string                                                `json:"linkedUrl" api:"required" format:"uri"`
	StartDate       time.Time                                             `json:"startDate" api:"required" format:"date-time"`
	JSON            dnsTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata
// for the struct [DNSTimeseriesResponseMetaConfidenceInfoAnnotation]
type dnsTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *DNSTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll                DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP                DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots               DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCt                 DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS                DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos                DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw                 DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI                DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet                DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAll, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBGP, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceBots, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceCt, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNS, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceDos, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFw, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceIQI, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceNet, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventType string

const (
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent             DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage            DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline          DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeEvent, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeOutage, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypePipeline, DNSTimeseriesResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type DNSTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime" api:"required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                              `json:"startTime" api:"required" format:"date-time"`
	JSON      dnsTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// dnsTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [DNSTimeseriesResponseMetaDateRange]
type dnsTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type DNSTimeseriesResponseMetaNormalization string

const (
	DNSTimeseriesResponseMetaNormalizationPercentage           DNSTimeseriesResponseMetaNormalization = "PERCENTAGE"
	DNSTimeseriesResponseMetaNormalizationMin0Max              DNSTimeseriesResponseMetaNormalization = "MIN0_MAX"
	DNSTimeseriesResponseMetaNormalizationMinMax               DNSTimeseriesResponseMetaNormalization = "MIN_MAX"
	DNSTimeseriesResponseMetaNormalizationRawValues            DNSTimeseriesResponseMetaNormalization = "RAW_VALUES"
	DNSTimeseriesResponseMetaNormalizationPercentageChange     DNSTimeseriesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	DNSTimeseriesResponseMetaNormalizationRollingAverage       DNSTimeseriesResponseMetaNormalization = "ROLLING_AVERAGE"
	DNSTimeseriesResponseMetaNormalizationOverlappedPercentage DNSTimeseriesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	DNSTimeseriesResponseMetaNormalizationRatio                DNSTimeseriesResponseMetaNormalization = "RATIO"
)

func (r DNSTimeseriesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case DNSTimeseriesResponseMetaNormalizationPercentage, DNSTimeseriesResponseMetaNormalizationMin0Max, DNSTimeseriesResponseMetaNormalizationMinMax, DNSTimeseriesResponseMetaNormalizationRawValues, DNSTimeseriesResponseMetaNormalizationPercentageChange, DNSTimeseriesResponseMetaNormalizationRollingAverage, DNSTimeseriesResponseMetaNormalizationOverlappedPercentage, DNSTimeseriesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type DNSTimeseriesResponseMetaUnit struct {
	Name  string                            `json:"name" api:"required"`
	Value string                            `json:"value" api:"required"`
	JSON  dnsTimeseriesResponseMetaUnitJSON `json:"-"`
}

// dnsTimeseriesResponseMetaUnitJSON contains the JSON metadata for the struct
// [DNSTimeseriesResponseMetaUnit]
type dnsTimeseriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[DNSTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	DNSSEC param.Field[[]DNSTimeseriesParamsDNSSEC] `query:"dnssec"`
	// Filters results based on DNSSEC (DNS Security Extensions) client awareness.
	DNSSECAware param.Field[[]DNSTimeseriesParamsDNSSECAware] `query:"dnssecAware"`
	// Filters results based on DNSSEC-validated answers by end-to-end security status.
	DNSSECE2E param.Field[[]bool] `query:"dnssecE2e"`
	// Format in which results will be returned.
	Format param.Field[DNSTimeseriesParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]DNSTimeseriesParamsIPVersion] `query:"ipVersion"`
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
	Protocol param.Field[[]DNSTimeseriesParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[[]DNSTimeseriesParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[[]DNSTimeseriesParamsResponseCode] `query:"responseCode"`
	// Filters results by DNS response TTL.
	ResponseTTL param.Field[[]DNSTimeseriesParamsResponseTTL] `query:"responseTtl"`
	// Filters results by top-level domain.
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesParams]'s query parameters as `url.Values`.
func (r DNSTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesParamsAggInterval string

const (
	DNSTimeseriesParamsAggInterval15m DNSTimeseriesParamsAggInterval = "15m"
	DNSTimeseriesParamsAggInterval1h  DNSTimeseriesParamsAggInterval = "1h"
	DNSTimeseriesParamsAggInterval1d  DNSTimeseriesParamsAggInterval = "1d"
	DNSTimeseriesParamsAggInterval1w  DNSTimeseriesParamsAggInterval = "1w"
)

func (r DNSTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsAggInterval15m, DNSTimeseriesParamsAggInterval1h, DNSTimeseriesParamsAggInterval1d, DNSTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

type DNSTimeseriesParamsDNSSEC string

const (
	DNSTimeseriesParamsDNSSECInvalid  DNSTimeseriesParamsDNSSEC = "INVALID"
	DNSTimeseriesParamsDNSSECInsecure DNSTimeseriesParamsDNSSEC = "INSECURE"
	DNSTimeseriesParamsDNSSECSecure   DNSTimeseriesParamsDNSSEC = "SECURE"
	DNSTimeseriesParamsDNSSECOther    DNSTimeseriesParamsDNSSEC = "OTHER"
)

func (r DNSTimeseriesParamsDNSSEC) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsDNSSECInvalid, DNSTimeseriesParamsDNSSECInsecure, DNSTimeseriesParamsDNSSECSecure, DNSTimeseriesParamsDNSSECOther:
		return true
	}
	return false
}

type DNSTimeseriesParamsDNSSECAware string

const (
	DNSTimeseriesParamsDNSSECAwareSupported    DNSTimeseriesParamsDNSSECAware = "SUPPORTED"
	DNSTimeseriesParamsDNSSECAwareNotSupported DNSTimeseriesParamsDNSSECAware = "NOT_SUPPORTED"
)

func (r DNSTimeseriesParamsDNSSECAware) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsDNSSECAwareSupported, DNSTimeseriesParamsDNSSECAwareNotSupported:
		return true
	}
	return false
}

// Format in which results will be returned.
type DNSTimeseriesParamsFormat string

const (
	DNSTimeseriesParamsFormatJson DNSTimeseriesParamsFormat = "JSON"
	DNSTimeseriesParamsFormatCsv  DNSTimeseriesParamsFormat = "CSV"
)

func (r DNSTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsFormatJson, DNSTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type DNSTimeseriesParamsIPVersion string

const (
	DNSTimeseriesParamsIPVersionIPv4 DNSTimeseriesParamsIPVersion = "IPv4"
	DNSTimeseriesParamsIPVersionIPv6 DNSTimeseriesParamsIPVersion = "IPv6"
)

func (r DNSTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsIPVersionIPv4, DNSTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

type DNSTimeseriesParamsProtocol string

const (
	DNSTimeseriesParamsProtocolUdp   DNSTimeseriesParamsProtocol = "UDP"
	DNSTimeseriesParamsProtocolTCP   DNSTimeseriesParamsProtocol = "TCP"
	DNSTimeseriesParamsProtocolHTTPS DNSTimeseriesParamsProtocol = "HTTPS"
	DNSTimeseriesParamsProtocolTLS   DNSTimeseriesParamsProtocol = "TLS"
)

func (r DNSTimeseriesParamsProtocol) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsProtocolUdp, DNSTimeseriesParamsProtocolTCP, DNSTimeseriesParamsProtocolHTTPS, DNSTimeseriesParamsProtocolTLS:
		return true
	}
	return false
}

type DNSTimeseriesParamsQueryType string

const (
	DNSTimeseriesParamsQueryTypeA          DNSTimeseriesParamsQueryType = "A"
	DNSTimeseriesParamsQueryTypeAAAA       DNSTimeseriesParamsQueryType = "AAAA"
	DNSTimeseriesParamsQueryTypeA6         DNSTimeseriesParamsQueryType = "A6"
	DNSTimeseriesParamsQueryTypeAfsdb      DNSTimeseriesParamsQueryType = "AFSDB"
	DNSTimeseriesParamsQueryTypeAny        DNSTimeseriesParamsQueryType = "ANY"
	DNSTimeseriesParamsQueryTypeApl        DNSTimeseriesParamsQueryType = "APL"
	DNSTimeseriesParamsQueryTypeAtma       DNSTimeseriesParamsQueryType = "ATMA"
	DNSTimeseriesParamsQueryTypeAXFR       DNSTimeseriesParamsQueryType = "AXFR"
	DNSTimeseriesParamsQueryTypeCAA        DNSTimeseriesParamsQueryType = "CAA"
	DNSTimeseriesParamsQueryTypeCdnskey    DNSTimeseriesParamsQueryType = "CDNSKEY"
	DNSTimeseriesParamsQueryTypeCds        DNSTimeseriesParamsQueryType = "CDS"
	DNSTimeseriesParamsQueryTypeCERT       DNSTimeseriesParamsQueryType = "CERT"
	DNSTimeseriesParamsQueryTypeCNAME      DNSTimeseriesParamsQueryType = "CNAME"
	DNSTimeseriesParamsQueryTypeCsync      DNSTimeseriesParamsQueryType = "CSYNC"
	DNSTimeseriesParamsQueryTypeDhcid      DNSTimeseriesParamsQueryType = "DHCID"
	DNSTimeseriesParamsQueryTypeDlv        DNSTimeseriesParamsQueryType = "DLV"
	DNSTimeseriesParamsQueryTypeDname      DNSTimeseriesParamsQueryType = "DNAME"
	DNSTimeseriesParamsQueryTypeDNSKEY     DNSTimeseriesParamsQueryType = "DNSKEY"
	DNSTimeseriesParamsQueryTypeDoa        DNSTimeseriesParamsQueryType = "DOA"
	DNSTimeseriesParamsQueryTypeDS         DNSTimeseriesParamsQueryType = "DS"
	DNSTimeseriesParamsQueryTypeEid        DNSTimeseriesParamsQueryType = "EID"
	DNSTimeseriesParamsQueryTypeEui48      DNSTimeseriesParamsQueryType = "EUI48"
	DNSTimeseriesParamsQueryTypeEui64      DNSTimeseriesParamsQueryType = "EUI64"
	DNSTimeseriesParamsQueryTypeGpos       DNSTimeseriesParamsQueryType = "GPOS"
	DNSTimeseriesParamsQueryTypeGid        DNSTimeseriesParamsQueryType = "GID"
	DNSTimeseriesParamsQueryTypeHinfo      DNSTimeseriesParamsQueryType = "HINFO"
	DNSTimeseriesParamsQueryTypeHip        DNSTimeseriesParamsQueryType = "HIP"
	DNSTimeseriesParamsQueryTypeHTTPS      DNSTimeseriesParamsQueryType = "HTTPS"
	DNSTimeseriesParamsQueryTypeIpseckey   DNSTimeseriesParamsQueryType = "IPSECKEY"
	DNSTimeseriesParamsQueryTypeIsdn       DNSTimeseriesParamsQueryType = "ISDN"
	DNSTimeseriesParamsQueryTypeIxfr       DNSTimeseriesParamsQueryType = "IXFR"
	DNSTimeseriesParamsQueryTypeKey        DNSTimeseriesParamsQueryType = "KEY"
	DNSTimeseriesParamsQueryTypeKx         DNSTimeseriesParamsQueryType = "KX"
	DNSTimeseriesParamsQueryTypeL32        DNSTimeseriesParamsQueryType = "L32"
	DNSTimeseriesParamsQueryTypeL64        DNSTimeseriesParamsQueryType = "L64"
	DNSTimeseriesParamsQueryTypeLOC        DNSTimeseriesParamsQueryType = "LOC"
	DNSTimeseriesParamsQueryTypeLp         DNSTimeseriesParamsQueryType = "LP"
	DNSTimeseriesParamsQueryTypeMaila      DNSTimeseriesParamsQueryType = "MAILA"
	DNSTimeseriesParamsQueryTypeMailb      DNSTimeseriesParamsQueryType = "MAILB"
	DNSTimeseriesParamsQueryTypeMB         DNSTimeseriesParamsQueryType = "MB"
	DNSTimeseriesParamsQueryTypeMd         DNSTimeseriesParamsQueryType = "MD"
	DNSTimeseriesParamsQueryTypeMf         DNSTimeseriesParamsQueryType = "MF"
	DNSTimeseriesParamsQueryTypeMg         DNSTimeseriesParamsQueryType = "MG"
	DNSTimeseriesParamsQueryTypeMinfo      DNSTimeseriesParamsQueryType = "MINFO"
	DNSTimeseriesParamsQueryTypeMr         DNSTimeseriesParamsQueryType = "MR"
	DNSTimeseriesParamsQueryTypeMX         DNSTimeseriesParamsQueryType = "MX"
	DNSTimeseriesParamsQueryTypeNAPTR      DNSTimeseriesParamsQueryType = "NAPTR"
	DNSTimeseriesParamsQueryTypeNb         DNSTimeseriesParamsQueryType = "NB"
	DNSTimeseriesParamsQueryTypeNbstat     DNSTimeseriesParamsQueryType = "NBSTAT"
	DNSTimeseriesParamsQueryTypeNid        DNSTimeseriesParamsQueryType = "NID"
	DNSTimeseriesParamsQueryTypeNimloc     DNSTimeseriesParamsQueryType = "NIMLOC"
	DNSTimeseriesParamsQueryTypeNinfo      DNSTimeseriesParamsQueryType = "NINFO"
	DNSTimeseriesParamsQueryTypeNS         DNSTimeseriesParamsQueryType = "NS"
	DNSTimeseriesParamsQueryTypeNsap       DNSTimeseriesParamsQueryType = "NSAP"
	DNSTimeseriesParamsQueryTypeNsec       DNSTimeseriesParamsQueryType = "NSEC"
	DNSTimeseriesParamsQueryTypeNsec3      DNSTimeseriesParamsQueryType = "NSEC3"
	DNSTimeseriesParamsQueryTypeNsec3Param DNSTimeseriesParamsQueryType = "NSEC3PARAM"
	DNSTimeseriesParamsQueryTypeNull       DNSTimeseriesParamsQueryType = "NULL"
	DNSTimeseriesParamsQueryTypeNxt        DNSTimeseriesParamsQueryType = "NXT"
	DNSTimeseriesParamsQueryTypeOpenpgpkey DNSTimeseriesParamsQueryType = "OPENPGPKEY"
	DNSTimeseriesParamsQueryTypeOpt        DNSTimeseriesParamsQueryType = "OPT"
	DNSTimeseriesParamsQueryTypePTR        DNSTimeseriesParamsQueryType = "PTR"
	DNSTimeseriesParamsQueryTypePx         DNSTimeseriesParamsQueryType = "PX"
	DNSTimeseriesParamsQueryTypeRkey       DNSTimeseriesParamsQueryType = "RKEY"
	DNSTimeseriesParamsQueryTypeRp         DNSTimeseriesParamsQueryType = "RP"
	DNSTimeseriesParamsQueryTypeRrsig      DNSTimeseriesParamsQueryType = "RRSIG"
	DNSTimeseriesParamsQueryTypeRt         DNSTimeseriesParamsQueryType = "RT"
	DNSTimeseriesParamsQueryTypeSig        DNSTimeseriesParamsQueryType = "SIG"
	DNSTimeseriesParamsQueryTypeSink       DNSTimeseriesParamsQueryType = "SINK"
	DNSTimeseriesParamsQueryTypeSMIMEA     DNSTimeseriesParamsQueryType = "SMIMEA"
	DNSTimeseriesParamsQueryTypeSOA        DNSTimeseriesParamsQueryType = "SOA"
	DNSTimeseriesParamsQueryTypeSPF        DNSTimeseriesParamsQueryType = "SPF"
	DNSTimeseriesParamsQueryTypeSRV        DNSTimeseriesParamsQueryType = "SRV"
	DNSTimeseriesParamsQueryTypeSSHFP      DNSTimeseriesParamsQueryType = "SSHFP"
	DNSTimeseriesParamsQueryTypeSVCB       DNSTimeseriesParamsQueryType = "SVCB"
	DNSTimeseriesParamsQueryTypeTa         DNSTimeseriesParamsQueryType = "TA"
	DNSTimeseriesParamsQueryTypeTalink     DNSTimeseriesParamsQueryType = "TALINK"
	DNSTimeseriesParamsQueryTypeTkey       DNSTimeseriesParamsQueryType = "TKEY"
	DNSTimeseriesParamsQueryTypeTLSA       DNSTimeseriesParamsQueryType = "TLSA"
	DNSTimeseriesParamsQueryTypeTSIG       DNSTimeseriesParamsQueryType = "TSIG"
	DNSTimeseriesParamsQueryTypeTXT        DNSTimeseriesParamsQueryType = "TXT"
	DNSTimeseriesParamsQueryTypeUinfo      DNSTimeseriesParamsQueryType = "UINFO"
	DNSTimeseriesParamsQueryTypeUID        DNSTimeseriesParamsQueryType = "UID"
	DNSTimeseriesParamsQueryTypeUnspec     DNSTimeseriesParamsQueryType = "UNSPEC"
	DNSTimeseriesParamsQueryTypeURI        DNSTimeseriesParamsQueryType = "URI"
	DNSTimeseriesParamsQueryTypeWks        DNSTimeseriesParamsQueryType = "WKS"
	DNSTimeseriesParamsQueryTypeX25        DNSTimeseriesParamsQueryType = "X25"
	DNSTimeseriesParamsQueryTypeZonemd     DNSTimeseriesParamsQueryType = "ZONEMD"
)

func (r DNSTimeseriesParamsQueryType) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsQueryTypeA, DNSTimeseriesParamsQueryTypeAAAA, DNSTimeseriesParamsQueryTypeA6, DNSTimeseriesParamsQueryTypeAfsdb, DNSTimeseriesParamsQueryTypeAny, DNSTimeseriesParamsQueryTypeApl, DNSTimeseriesParamsQueryTypeAtma, DNSTimeseriesParamsQueryTypeAXFR, DNSTimeseriesParamsQueryTypeCAA, DNSTimeseriesParamsQueryTypeCdnskey, DNSTimeseriesParamsQueryTypeCds, DNSTimeseriesParamsQueryTypeCERT, DNSTimeseriesParamsQueryTypeCNAME, DNSTimeseriesParamsQueryTypeCsync, DNSTimeseriesParamsQueryTypeDhcid, DNSTimeseriesParamsQueryTypeDlv, DNSTimeseriesParamsQueryTypeDname, DNSTimeseriesParamsQueryTypeDNSKEY, DNSTimeseriesParamsQueryTypeDoa, DNSTimeseriesParamsQueryTypeDS, DNSTimeseriesParamsQueryTypeEid, DNSTimeseriesParamsQueryTypeEui48, DNSTimeseriesParamsQueryTypeEui64, DNSTimeseriesParamsQueryTypeGpos, DNSTimeseriesParamsQueryTypeGid, DNSTimeseriesParamsQueryTypeHinfo, DNSTimeseriesParamsQueryTypeHip, DNSTimeseriesParamsQueryTypeHTTPS, DNSTimeseriesParamsQueryTypeIpseckey, DNSTimeseriesParamsQueryTypeIsdn, DNSTimeseriesParamsQueryTypeIxfr, DNSTimeseriesParamsQueryTypeKey, DNSTimeseriesParamsQueryTypeKx, DNSTimeseriesParamsQueryTypeL32, DNSTimeseriesParamsQueryTypeL64, DNSTimeseriesParamsQueryTypeLOC, DNSTimeseriesParamsQueryTypeLp, DNSTimeseriesParamsQueryTypeMaila, DNSTimeseriesParamsQueryTypeMailb, DNSTimeseriesParamsQueryTypeMB, DNSTimeseriesParamsQueryTypeMd, DNSTimeseriesParamsQueryTypeMf, DNSTimeseriesParamsQueryTypeMg, DNSTimeseriesParamsQueryTypeMinfo, DNSTimeseriesParamsQueryTypeMr, DNSTimeseriesParamsQueryTypeMX, DNSTimeseriesParamsQueryTypeNAPTR, DNSTimeseriesParamsQueryTypeNb, DNSTimeseriesParamsQueryTypeNbstat, DNSTimeseriesParamsQueryTypeNid, DNSTimeseriesParamsQueryTypeNimloc, DNSTimeseriesParamsQueryTypeNinfo, DNSTimeseriesParamsQueryTypeNS, DNSTimeseriesParamsQueryTypeNsap, DNSTimeseriesParamsQueryTypeNsec, DNSTimeseriesParamsQueryTypeNsec3, DNSTimeseriesParamsQueryTypeNsec3Param, DNSTimeseriesParamsQueryTypeNull, DNSTimeseriesParamsQueryTypeNxt, DNSTimeseriesParamsQueryTypeOpenpgpkey, DNSTimeseriesParamsQueryTypeOpt, DNSTimeseriesParamsQueryTypePTR, DNSTimeseriesParamsQueryTypePx, DNSTimeseriesParamsQueryTypeRkey, DNSTimeseriesParamsQueryTypeRp, DNSTimeseriesParamsQueryTypeRrsig, DNSTimeseriesParamsQueryTypeRt, DNSTimeseriesParamsQueryTypeSig, DNSTimeseriesParamsQueryTypeSink, DNSTimeseriesParamsQueryTypeSMIMEA, DNSTimeseriesParamsQueryTypeSOA, DNSTimeseriesParamsQueryTypeSPF, DNSTimeseriesParamsQueryTypeSRV, DNSTimeseriesParamsQueryTypeSSHFP, DNSTimeseriesParamsQueryTypeSVCB, DNSTimeseriesParamsQueryTypeTa, DNSTimeseriesParamsQueryTypeTalink, DNSTimeseriesParamsQueryTypeTkey, DNSTimeseriesParamsQueryTypeTLSA, DNSTimeseriesParamsQueryTypeTSIG, DNSTimeseriesParamsQueryTypeTXT, DNSTimeseriesParamsQueryTypeUinfo, DNSTimeseriesParamsQueryTypeUID, DNSTimeseriesParamsQueryTypeUnspec, DNSTimeseriesParamsQueryTypeURI, DNSTimeseriesParamsQueryTypeWks, DNSTimeseriesParamsQueryTypeX25, DNSTimeseriesParamsQueryTypeZonemd:
		return true
	}
	return false
}

type DNSTimeseriesParamsResponseCode string

const (
	DNSTimeseriesParamsResponseCodeNoerror   DNSTimeseriesParamsResponseCode = "NOERROR"
	DNSTimeseriesParamsResponseCodeFormerr   DNSTimeseriesParamsResponseCode = "FORMERR"
	DNSTimeseriesParamsResponseCodeServfail  DNSTimeseriesParamsResponseCode = "SERVFAIL"
	DNSTimeseriesParamsResponseCodeNxdomain  DNSTimeseriesParamsResponseCode = "NXDOMAIN"
	DNSTimeseriesParamsResponseCodeNotimp    DNSTimeseriesParamsResponseCode = "NOTIMP"
	DNSTimeseriesParamsResponseCodeRefused   DNSTimeseriesParamsResponseCode = "REFUSED"
	DNSTimeseriesParamsResponseCodeYxdomain  DNSTimeseriesParamsResponseCode = "YXDOMAIN"
	DNSTimeseriesParamsResponseCodeYxrrset   DNSTimeseriesParamsResponseCode = "YXRRSET"
	DNSTimeseriesParamsResponseCodeNxrrset   DNSTimeseriesParamsResponseCode = "NXRRSET"
	DNSTimeseriesParamsResponseCodeNotauth   DNSTimeseriesParamsResponseCode = "NOTAUTH"
	DNSTimeseriesParamsResponseCodeNotzone   DNSTimeseriesParamsResponseCode = "NOTZONE"
	DNSTimeseriesParamsResponseCodeBadsig    DNSTimeseriesParamsResponseCode = "BADSIG"
	DNSTimeseriesParamsResponseCodeBadkey    DNSTimeseriesParamsResponseCode = "BADKEY"
	DNSTimeseriesParamsResponseCodeBadtime   DNSTimeseriesParamsResponseCode = "BADTIME"
	DNSTimeseriesParamsResponseCodeBadmode   DNSTimeseriesParamsResponseCode = "BADMODE"
	DNSTimeseriesParamsResponseCodeBadname   DNSTimeseriesParamsResponseCode = "BADNAME"
	DNSTimeseriesParamsResponseCodeBadalg    DNSTimeseriesParamsResponseCode = "BADALG"
	DNSTimeseriesParamsResponseCodeBadtrunc  DNSTimeseriesParamsResponseCode = "BADTRUNC"
	DNSTimeseriesParamsResponseCodeBadcookie DNSTimeseriesParamsResponseCode = "BADCOOKIE"
)

func (r DNSTimeseriesParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsResponseCodeNoerror, DNSTimeseriesParamsResponseCodeFormerr, DNSTimeseriesParamsResponseCodeServfail, DNSTimeseriesParamsResponseCodeNxdomain, DNSTimeseriesParamsResponseCodeNotimp, DNSTimeseriesParamsResponseCodeRefused, DNSTimeseriesParamsResponseCodeYxdomain, DNSTimeseriesParamsResponseCodeYxrrset, DNSTimeseriesParamsResponseCodeNxrrset, DNSTimeseriesParamsResponseCodeNotauth, DNSTimeseriesParamsResponseCodeNotzone, DNSTimeseriesParamsResponseCodeBadsig, DNSTimeseriesParamsResponseCodeBadkey, DNSTimeseriesParamsResponseCodeBadtime, DNSTimeseriesParamsResponseCodeBadmode, DNSTimeseriesParamsResponseCodeBadname, DNSTimeseriesParamsResponseCodeBadalg, DNSTimeseriesParamsResponseCodeBadtrunc, DNSTimeseriesParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSTimeseriesParamsResponseTTL string

const (
	DNSTimeseriesParamsResponseTTLLte1M      DNSTimeseriesParamsResponseTTL = "LTE_1M"
	DNSTimeseriesParamsResponseTTLGt1MLte5M  DNSTimeseriesParamsResponseTTL = "GT_1M_LTE_5M"
	DNSTimeseriesParamsResponseTTLGt5MLte15M DNSTimeseriesParamsResponseTTL = "GT_5M_LTE_15M"
	DNSTimeseriesParamsResponseTTLGt15MLte1H DNSTimeseriesParamsResponseTTL = "GT_15M_LTE_1H"
	DNSTimeseriesParamsResponseTTLGt1HLte1D  DNSTimeseriesParamsResponseTTL = "GT_1H_LTE_1D"
	DNSTimeseriesParamsResponseTTLGt1DLte1W  DNSTimeseriesParamsResponseTTL = "GT_1D_LTE_1W"
	DNSTimeseriesParamsResponseTTLGt1W       DNSTimeseriesParamsResponseTTL = "GT_1W"
)

func (r DNSTimeseriesParamsResponseTTL) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsResponseTTLLte1M, DNSTimeseriesParamsResponseTTLGt1MLte5M, DNSTimeseriesParamsResponseTTLGt5MLte15M, DNSTimeseriesParamsResponseTTLGt15MLte1H, DNSTimeseriesParamsResponseTTLGt1HLte1D, DNSTimeseriesParamsResponseTTLGt1DLte1W, DNSTimeseriesParamsResponseTTLGt1W:
		return true
	}
	return false
}

type DNSTimeseriesResponseEnvelope struct {
	Result  DNSTimeseriesResponse             `json:"result" api:"required"`
	Success bool                              `json:"success" api:"required"`
	JSON    dnsTimeseriesResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSTimeseriesResponseEnvelope]
type dnsTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
