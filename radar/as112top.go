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

// AS112TopService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAS112TopService] method instead.
type AS112TopService struct {
	Options []option.RequestOption
}

// NewAS112TopService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAS112TopService(opts ...option.RequestOption) (r *AS112TopService) {
	r = &AS112TopService{}
	r.Options = opts
	return
}

// Retrieves the top locations of DNS queries to AS112 with DNSSEC (DNS Security
// Extensions) support.
func (r *AS112TopService) DNSSEC(ctx context.Context, dnssec AS112TopDNSSECParamsDNSSEC, query AS112TopDNSSECParams, opts ...option.RequestOption) (res *AS112TopDNSSECResponse, err error) {
	var env AS112TopDNSSECResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/as112/top/locations/dnssec/%v", dnssec)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the top locations of DNS queries to AS112 with EDNS (Extension
// Mechanisms for DNS) support.
func (r *AS112TopService) Edns(ctx context.Context, edns AS112TopEdnsParamsEdns, query AS112TopEdnsParams, opts ...option.RequestOption) (res *AS112TopEdnsResponse, err error) {
	var env AS112TopEdnsResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/as112/top/locations/edns/%v", edns)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the top locations of DNS queries to AS112 for an IP version.
func (r *AS112TopService) IPVersion(ctx context.Context, ipVersion AS112TopIPVersionParamsIPVersion, query AS112TopIPVersionParams, opts ...option.RequestOption) (res *AS112TopIPVersionResponse, err error) {
	var env AS112TopIPVersionResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/as112/top/locations/ip_version/%v", ipVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the top locations by AS112 DNS queries.
func (r *AS112TopService) Locations(ctx context.Context, query AS112TopLocationsParams, opts ...option.RequestOption) (res *AS112TopLocationsResponse, err error) {
	var env AS112TopLocationsResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/as112/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AS112TopDNSSECResponse struct {
	// Metadata for the results.
	Meta AS112TopDNSSECResponseMeta   `json:"meta,required"`
	Top0 []AS112TopDNSSECResponseTop0 `json:"top_0,required"`
	JSON as112TopDNSSECResponseJSON   `json:"-"`
}

// as112TopDNSSECResponseJSON contains the JSON metadata for the struct
// [AS112TopDNSSECResponse]
type as112TopDNSSECResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopDNSSECResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopDNSSECResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112TopDNSSECResponseMeta struct {
	ConfidenceInfo AS112TopDNSSECResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []AS112TopDNSSECResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112TopDNSSECResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112TopDNSSECResponseMetaUnit `json:"units,required"`
	JSON  as112TopDNSSECResponseMetaJSON   `json:"-"`
}

// as112TopDNSSECResponseMetaJSON contains the JSON metadata for the struct
// [AS112TopDNSSECResponseMeta]
type as112TopDNSSECResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112TopDNSSECResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopDNSSECResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112TopDNSSECResponseMetaConfidenceInfo struct {
	Annotations []AS112TopDNSSECResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                        `json:"level,required"`
	JSON  as112TopDNSSECResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112TopDNSSECResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [AS112TopDNSSECResponseMetaConfidenceInfo]
type as112TopDNSSECResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopDNSSECResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopDNSSECResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112TopDNSSECResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                        `json:"description,required"`
	EndDate     time.Time                                                     `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                              `json:"startDate,required" format:"date-time"`
	JSON            as112TopDNSSECResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112TopDNSSECResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [AS112TopDNSSECResponseMetaConfidenceInfoAnnotation]
type as112TopDNSSECResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112TopDNSSECResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopDNSSECResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceAll                AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceBots               AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceDos                AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceNet                AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceAll, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceBots, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceCT, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceDos, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceFw, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceNet, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AS112TopDNSSECResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AS112TopDNSSECResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                               `json:"startTime,required" format:"date-time"`
	JSON      as112TopDNSSECResponseMetaDateRangeJSON `json:"-"`
}

// as112TopDNSSECResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [AS112TopDNSSECResponseMetaDateRange]
type as112TopDNSSECResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopDNSSECResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopDNSSECResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112TopDNSSECResponseMetaNormalization string

const (
	AS112TopDNSSECResponseMetaNormalizationPercentage           AS112TopDNSSECResponseMetaNormalization = "PERCENTAGE"
	AS112TopDNSSECResponseMetaNormalizationMin0Max              AS112TopDNSSECResponseMetaNormalization = "MIN0_MAX"
	AS112TopDNSSECResponseMetaNormalizationMinMax               AS112TopDNSSECResponseMetaNormalization = "MIN_MAX"
	AS112TopDNSSECResponseMetaNormalizationRawValues            AS112TopDNSSECResponseMetaNormalization = "RAW_VALUES"
	AS112TopDNSSECResponseMetaNormalizationPercentageChange     AS112TopDNSSECResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112TopDNSSECResponseMetaNormalizationRollingAverage       AS112TopDNSSECResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112TopDNSSECResponseMetaNormalizationOverlappedPercentage AS112TopDNSSECResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112TopDNSSECResponseMetaNormalizationRatio                AS112TopDNSSECResponseMetaNormalization = "RATIO"
)

func (r AS112TopDNSSECResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112TopDNSSECResponseMetaNormalizationPercentage, AS112TopDNSSECResponseMetaNormalizationMin0Max, AS112TopDNSSECResponseMetaNormalizationMinMax, AS112TopDNSSECResponseMetaNormalizationRawValues, AS112TopDNSSECResponseMetaNormalizationPercentageChange, AS112TopDNSSECResponseMetaNormalizationRollingAverage, AS112TopDNSSECResponseMetaNormalizationOverlappedPercentage, AS112TopDNSSECResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112TopDNSSECResponseMetaUnit struct {
	Name  string                             `json:"name,required"`
	Value string                             `json:"value,required"`
	JSON  as112TopDNSSECResponseMetaUnitJSON `json:"-"`
}

// as112TopDNSSECResponseMetaUnitJSON contains the JSON metadata for the struct
// [AS112TopDNSSECResponseMetaUnit]
type as112TopDNSSECResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopDNSSECResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopDNSSECResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112TopDNSSECResponseTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                         `json:"value,required"`
	JSON  as112TopDNSSECResponseTop0JSON `json:"-"`
}

// as112TopDNSSECResponseTop0JSON contains the JSON metadata for the struct
// [AS112TopDNSSECResponseTop0]
type as112TopDNSSECResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AS112TopDNSSECResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopDNSSECResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AS112TopEdnsResponse struct {
	// Metadata for the results.
	Meta AS112TopEdnsResponseMeta   `json:"meta,required"`
	Top0 []AS112TopEdnsResponseTop0 `json:"top_0,required"`
	JSON as112TopEdnsResponseJSON   `json:"-"`
}

// as112TopEdnsResponseJSON contains the JSON metadata for the struct
// [AS112TopEdnsResponse]
type as112TopEdnsResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopEdnsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopEdnsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112TopEdnsResponseMeta struct {
	ConfidenceInfo AS112TopEdnsResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []AS112TopEdnsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112TopEdnsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112TopEdnsResponseMetaUnit `json:"units,required"`
	JSON  as112TopEdnsResponseMetaJSON   `json:"-"`
}

// as112TopEdnsResponseMetaJSON contains the JSON metadata for the struct
// [AS112TopEdnsResponseMeta]
type as112TopEdnsResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112TopEdnsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopEdnsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112TopEdnsResponseMetaConfidenceInfo struct {
	Annotations []AS112TopEdnsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                      `json:"level,required"`
	JSON  as112TopEdnsResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112TopEdnsResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [AS112TopEdnsResponseMetaConfidenceInfo]
type as112TopEdnsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopEdnsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopEdnsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112TopEdnsResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                      `json:"description,required"`
	EndDate     time.Time                                                   `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                            `json:"startDate,required" format:"date-time"`
	JSON            as112TopEdnsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112TopEdnsResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata
// for the struct [AS112TopEdnsResponseMetaConfidenceInfoAnnotation]
type as112TopEdnsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112TopEdnsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopEdnsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceAll                AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceBots               AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceDos                AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceNet                AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceAll, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceBots, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceCT, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceDos, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceFw, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceNet, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AS112TopEdnsResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AS112TopEdnsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                             `json:"startTime,required" format:"date-time"`
	JSON      as112TopEdnsResponseMetaDateRangeJSON `json:"-"`
}

// as112TopEdnsResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [AS112TopEdnsResponseMetaDateRange]
type as112TopEdnsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopEdnsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopEdnsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112TopEdnsResponseMetaNormalization string

const (
	AS112TopEdnsResponseMetaNormalizationPercentage           AS112TopEdnsResponseMetaNormalization = "PERCENTAGE"
	AS112TopEdnsResponseMetaNormalizationMin0Max              AS112TopEdnsResponseMetaNormalization = "MIN0_MAX"
	AS112TopEdnsResponseMetaNormalizationMinMax               AS112TopEdnsResponseMetaNormalization = "MIN_MAX"
	AS112TopEdnsResponseMetaNormalizationRawValues            AS112TopEdnsResponseMetaNormalization = "RAW_VALUES"
	AS112TopEdnsResponseMetaNormalizationPercentageChange     AS112TopEdnsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112TopEdnsResponseMetaNormalizationRollingAverage       AS112TopEdnsResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112TopEdnsResponseMetaNormalizationOverlappedPercentage AS112TopEdnsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112TopEdnsResponseMetaNormalizationRatio                AS112TopEdnsResponseMetaNormalization = "RATIO"
)

func (r AS112TopEdnsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112TopEdnsResponseMetaNormalizationPercentage, AS112TopEdnsResponseMetaNormalizationMin0Max, AS112TopEdnsResponseMetaNormalizationMinMax, AS112TopEdnsResponseMetaNormalizationRawValues, AS112TopEdnsResponseMetaNormalizationPercentageChange, AS112TopEdnsResponseMetaNormalizationRollingAverage, AS112TopEdnsResponseMetaNormalizationOverlappedPercentage, AS112TopEdnsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112TopEdnsResponseMetaUnit struct {
	Name  string                           `json:"name,required"`
	Value string                           `json:"value,required"`
	JSON  as112TopEdnsResponseMetaUnitJSON `json:"-"`
}

// as112TopEdnsResponseMetaUnitJSON contains the JSON metadata for the struct
// [AS112TopEdnsResponseMetaUnit]
type as112TopEdnsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopEdnsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopEdnsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112TopEdnsResponseTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                       `json:"value,required"`
	JSON  as112TopEdnsResponseTop0JSON `json:"-"`
}

// as112TopEdnsResponseTop0JSON contains the JSON metadata for the struct
// [AS112TopEdnsResponseTop0]
type as112TopEdnsResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AS112TopEdnsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopEdnsResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AS112TopIPVersionResponse struct {
	// Metadata for the results.
	Meta AS112TopIPVersionResponseMeta   `json:"meta,required"`
	Top0 []AS112TopIPVersionResponseTop0 `json:"top_0,required"`
	JSON as112TopIPVersionResponseJSON   `json:"-"`
}

// as112TopIPVersionResponseJSON contains the JSON metadata for the struct
// [AS112TopIPVersionResponse]
type as112TopIPVersionResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112TopIPVersionResponseMeta struct {
	ConfidenceInfo AS112TopIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []AS112TopIPVersionResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112TopIPVersionResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112TopIPVersionResponseMetaUnit `json:"units,required"`
	JSON  as112TopIPVersionResponseMetaJSON   `json:"-"`
}

// as112TopIPVersionResponseMetaJSON contains the JSON metadata for the struct
// [AS112TopIPVersionResponseMeta]
type as112TopIPVersionResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112TopIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopIPVersionResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112TopIPVersionResponseMetaConfidenceInfo struct {
	Annotations []AS112TopIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                           `json:"level,required"`
	JSON  as112TopIPVersionResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112TopIPVersionResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [AS112TopIPVersionResponseMetaConfidenceInfo]
type as112TopIPVersionResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopIPVersionResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112TopIPVersionResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                           `json:"description,required"`
	EndDate     time.Time                                                        `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                      `json:"isInstantaneous,required"`
	LinkedURL       string                                                    `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                 `json:"startDate,required" format:"date-time"`
	JSON            as112TopIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112TopIPVersionResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [AS112TopIPVersionResponseMetaConfidenceInfoAnnotation]
type as112TopIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112TopIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceAll                AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceBots               AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceDos                AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceNet                AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceAll, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceBots, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceCT, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceDos, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceFw, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceNet, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AS112TopIPVersionResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AS112TopIPVersionResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                  `json:"startTime,required" format:"date-time"`
	JSON      as112TopIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// as112TopIPVersionResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [AS112TopIPVersionResponseMetaDateRange]
type as112TopIPVersionResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopIPVersionResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112TopIPVersionResponseMetaNormalization string

const (
	AS112TopIPVersionResponseMetaNormalizationPercentage           AS112TopIPVersionResponseMetaNormalization = "PERCENTAGE"
	AS112TopIPVersionResponseMetaNormalizationMin0Max              AS112TopIPVersionResponseMetaNormalization = "MIN0_MAX"
	AS112TopIPVersionResponseMetaNormalizationMinMax               AS112TopIPVersionResponseMetaNormalization = "MIN_MAX"
	AS112TopIPVersionResponseMetaNormalizationRawValues            AS112TopIPVersionResponseMetaNormalization = "RAW_VALUES"
	AS112TopIPVersionResponseMetaNormalizationPercentageChange     AS112TopIPVersionResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112TopIPVersionResponseMetaNormalizationRollingAverage       AS112TopIPVersionResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112TopIPVersionResponseMetaNormalizationOverlappedPercentage AS112TopIPVersionResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112TopIPVersionResponseMetaNormalizationRatio                AS112TopIPVersionResponseMetaNormalization = "RATIO"
)

func (r AS112TopIPVersionResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112TopIPVersionResponseMetaNormalizationPercentage, AS112TopIPVersionResponseMetaNormalizationMin0Max, AS112TopIPVersionResponseMetaNormalizationMinMax, AS112TopIPVersionResponseMetaNormalizationRawValues, AS112TopIPVersionResponseMetaNormalizationPercentageChange, AS112TopIPVersionResponseMetaNormalizationRollingAverage, AS112TopIPVersionResponseMetaNormalizationOverlappedPercentage, AS112TopIPVersionResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112TopIPVersionResponseMetaUnit struct {
	Name  string                                `json:"name,required"`
	Value string                                `json:"value,required"`
	JSON  as112TopIPVersionResponseMetaUnitJSON `json:"-"`
}

// as112TopIPVersionResponseMetaUnitJSON contains the JSON metadata for the struct
// [AS112TopIPVersionResponseMetaUnit]
type as112TopIPVersionResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopIPVersionResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopIPVersionResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112TopIPVersionResponseTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                            `json:"value,required"`
	JSON  as112TopIPVersionResponseTop0JSON `json:"-"`
}

// as112TopIPVersionResponseTop0JSON contains the JSON metadata for the struct
// [AS112TopIPVersionResponseTop0]
type as112TopIPVersionResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AS112TopIPVersionResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopIPVersionResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AS112TopLocationsResponse struct {
	// Metadata for the results.
	Meta AS112TopLocationsResponseMeta   `json:"meta,required"`
	Top0 []AS112TopLocationsResponseTop0 `json:"top_0,required"`
	JSON as112TopLocationsResponseJSON   `json:"-"`
}

// as112TopLocationsResponseJSON contains the JSON metadata for the struct
// [AS112TopLocationsResponse]
type as112TopLocationsResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopLocationsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AS112TopLocationsResponseMeta struct {
	ConfidenceInfo AS112TopLocationsResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []AS112TopLocationsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AS112TopLocationsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AS112TopLocationsResponseMetaUnit `json:"units,required"`
	JSON  as112TopLocationsResponseMetaJSON   `json:"-"`
}

// as112TopLocationsResponseMetaJSON contains the JSON metadata for the struct
// [AS112TopLocationsResponseMeta]
type as112TopLocationsResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AS112TopLocationsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopLocationsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AS112TopLocationsResponseMetaConfidenceInfo struct {
	Annotations []AS112TopLocationsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                           `json:"level,required"`
	JSON  as112TopLocationsResponseMetaConfidenceInfoJSON `json:"-"`
}

// as112TopLocationsResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [AS112TopLocationsResponseMetaConfidenceInfo]
type as112TopLocationsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopLocationsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopLocationsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AS112TopLocationsResponseMetaConfidenceInfoAnnotation struct {
	// Data source for annotations.
	DataSource  AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource `json:"dataSource,required"`
	Description string                                                           `json:"description,required"`
	EndDate     time.Time                                                        `json:"endDate,required" format:"date-time"`
	// Event type for annotations.
	EventType AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventType `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                      `json:"isInstantaneous,required"`
	LinkedURL       string                                                    `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                 `json:"startDate,required" format:"date-time"`
	JSON            as112TopLocationsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// as112TopLocationsResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [AS112TopLocationsResponseMetaConfidenceInfoAnnotation]
type as112TopLocationsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AS112TopLocationsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopLocationsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

// Data source for annotations.
type AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource string

const (
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceAll                AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "ALL"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceAIBots             AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "AI_BOTS"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway          AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "AI_GATEWAY"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceBGP                AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "BGP"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceBots               AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "BOTS"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly  AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "CONNECTION_ANOMALY"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceCT                 AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "CT"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDNS                AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude       AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_MAGNITUDE"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112           AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "DNS_AS112"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDos                AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "DOS"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting       AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_ROUTING"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity      AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "EMAIL_SECURITY"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceFw                 AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "FW"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceFwPg               AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "FW_PG"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTP               AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl        AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CONTROL"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_CRAWLER_REFERER"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins        AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "HTTP_ORIGINS"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceIQI                AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "IQI"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials  AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "LEAKED_CREDENTIALS"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceNet                AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "NET"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT          AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "ROBOTS_TXT"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceSpeed              AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "SPEED"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI          AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource = "WORKERS_AI"
)

func (r AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSource) IsKnown() bool {
	switch r {
	case AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceAll, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceAIBots, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceAIGateway, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceBGP, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceBots, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceConnectionAnomaly, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceCT, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDNS, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDNSMagnitude, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDNSAS112, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceDos, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceEmailRouting, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceEmailSecurity, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceFw, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceFwPg, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTP, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPControl, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPCrawlerReferer, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceHTTPOrigins, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceIQI, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceLeakedCredentials, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceNet, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceRobotsTXT, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceSpeed, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsDataSourceWorkersAI:
		return true
	}
	return false
}

// Event type for annotations.
type AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventType string

const (
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeEvent             AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventType = "EVENT"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeGeneral           AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventType = "GENERAL"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeOutage            AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventType = "OUTAGE"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventType = "PARTIAL_PROJECTION"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventTypePipeline          AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventType = "PIPELINE"
	AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly    AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventType = "TRAFFIC_ANOMALY"
)

func (r AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventType) IsKnown() bool {
	switch r {
	case AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeEvent, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeGeneral, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeOutage, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventTypePartialProjection, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventTypePipeline, AS112TopLocationsResponseMetaConfidenceInfoAnnotationsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

type AS112TopLocationsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                  `json:"startTime,required" format:"date-time"`
	JSON      as112TopLocationsResponseMetaDateRangeJSON `json:"-"`
}

// as112TopLocationsResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [AS112TopLocationsResponseMetaDateRange]
type as112TopLocationsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopLocationsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopLocationsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AS112TopLocationsResponseMetaNormalization string

const (
	AS112TopLocationsResponseMetaNormalizationPercentage           AS112TopLocationsResponseMetaNormalization = "PERCENTAGE"
	AS112TopLocationsResponseMetaNormalizationMin0Max              AS112TopLocationsResponseMetaNormalization = "MIN0_MAX"
	AS112TopLocationsResponseMetaNormalizationMinMax               AS112TopLocationsResponseMetaNormalization = "MIN_MAX"
	AS112TopLocationsResponseMetaNormalizationRawValues            AS112TopLocationsResponseMetaNormalization = "RAW_VALUES"
	AS112TopLocationsResponseMetaNormalizationPercentageChange     AS112TopLocationsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AS112TopLocationsResponseMetaNormalizationRollingAverage       AS112TopLocationsResponseMetaNormalization = "ROLLING_AVERAGE"
	AS112TopLocationsResponseMetaNormalizationOverlappedPercentage AS112TopLocationsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AS112TopLocationsResponseMetaNormalizationRatio                AS112TopLocationsResponseMetaNormalization = "RATIO"
)

func (r AS112TopLocationsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AS112TopLocationsResponseMetaNormalizationPercentage, AS112TopLocationsResponseMetaNormalizationMin0Max, AS112TopLocationsResponseMetaNormalizationMinMax, AS112TopLocationsResponseMetaNormalizationRawValues, AS112TopLocationsResponseMetaNormalizationPercentageChange, AS112TopLocationsResponseMetaNormalizationRollingAverage, AS112TopLocationsResponseMetaNormalizationOverlappedPercentage, AS112TopLocationsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AS112TopLocationsResponseMetaUnit struct {
	Name  string                                `json:"name,required"`
	Value string                                `json:"value,required"`
	JSON  as112TopLocationsResponseMetaUnitJSON `json:"-"`
}

// as112TopLocationsResponseMetaUnitJSON contains the JSON metadata for the struct
// [AS112TopLocationsResponseMetaUnit]
type as112TopLocationsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopLocationsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopLocationsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AS112TopLocationsResponseTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                            `json:"value,required"`
	JSON  as112TopLocationsResponseTop0JSON `json:"-"`
}

// as112TopLocationsResponseTop0JSON contains the JSON metadata for the struct
// [AS112TopLocationsResponseTop0]
type as112TopLocationsResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AS112TopLocationsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopLocationsResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AS112TopDNSSECParams struct {
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
	Format param.Field[AS112TopDNSSECParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112TopDNSSECParams]'s query parameters as `url.Values`.
func (r AS112TopDNSSECParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// DNSSEC (DNS Security Extensions) status.
type AS112TopDNSSECParamsDNSSEC string

const (
	AS112TopDNSSECParamsDNSSECSupported    AS112TopDNSSECParamsDNSSEC = "SUPPORTED"
	AS112TopDNSSECParamsDNSSECNotSupported AS112TopDNSSECParamsDNSSEC = "NOT_SUPPORTED"
)

func (r AS112TopDNSSECParamsDNSSEC) IsKnown() bool {
	switch r {
	case AS112TopDNSSECParamsDNSSECSupported, AS112TopDNSSECParamsDNSSECNotSupported:
		return true
	}
	return false
}

// Format in which results will be returned.
type AS112TopDNSSECParamsFormat string

const (
	AS112TopDNSSECParamsFormatJson AS112TopDNSSECParamsFormat = "JSON"
	AS112TopDNSSECParamsFormatCsv  AS112TopDNSSECParamsFormat = "CSV"
)

func (r AS112TopDNSSECParamsFormat) IsKnown() bool {
	switch r {
	case AS112TopDNSSECParamsFormatJson, AS112TopDNSSECParamsFormatCsv:
		return true
	}
	return false
}

type AS112TopDNSSECResponseEnvelope struct {
	Result  AS112TopDNSSECResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    as112TopDNSSECResponseEnvelopeJSON `json:"-"`
}

// as112TopDNSSECResponseEnvelopeJSON contains the JSON metadata for the struct
// [AS112TopDNSSECResponseEnvelope]
type as112TopDNSSECResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopDNSSECResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopDNSSECResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112TopEdnsParams struct {
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
	Format param.Field[AS112TopEdnsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112TopEdnsParams]'s query parameters as `url.Values`.
func (r AS112TopEdnsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// EDNS (Extension Mechanisms for DNS) status.
type AS112TopEdnsParamsEdns string

const (
	AS112TopEdnsParamsEdnsSupported    AS112TopEdnsParamsEdns = "SUPPORTED"
	AS112TopEdnsParamsEdnsNotSupported AS112TopEdnsParamsEdns = "NOT_SUPPORTED"
)

func (r AS112TopEdnsParamsEdns) IsKnown() bool {
	switch r {
	case AS112TopEdnsParamsEdnsSupported, AS112TopEdnsParamsEdnsNotSupported:
		return true
	}
	return false
}

// Format in which results will be returned.
type AS112TopEdnsParamsFormat string

const (
	AS112TopEdnsParamsFormatJson AS112TopEdnsParamsFormat = "JSON"
	AS112TopEdnsParamsFormatCsv  AS112TopEdnsParamsFormat = "CSV"
)

func (r AS112TopEdnsParamsFormat) IsKnown() bool {
	switch r {
	case AS112TopEdnsParamsFormatJson, AS112TopEdnsParamsFormatCsv:
		return true
	}
	return false
}

type AS112TopEdnsResponseEnvelope struct {
	Result  AS112TopEdnsResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    as112TopEdnsResponseEnvelopeJSON `json:"-"`
}

// as112TopEdnsResponseEnvelopeJSON contains the JSON metadata for the struct
// [AS112TopEdnsResponseEnvelope]
type as112TopEdnsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopEdnsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopEdnsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112TopIPVersionParams struct {
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
	Format param.Field[AS112TopIPVersionParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112TopIPVersionParams]'s query parameters as
// `url.Values`.
func (r AS112TopIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// IP version.
type AS112TopIPVersionParamsIPVersion string

const (
	AS112TopIPVersionParamsIPVersionIPv4 AS112TopIPVersionParamsIPVersion = "IPv4"
	AS112TopIPVersionParamsIPVersionIPv6 AS112TopIPVersionParamsIPVersion = "IPv6"
)

func (r AS112TopIPVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case AS112TopIPVersionParamsIPVersionIPv4, AS112TopIPVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

// Format in which results will be returned.
type AS112TopIPVersionParamsFormat string

const (
	AS112TopIPVersionParamsFormatJson AS112TopIPVersionParamsFormat = "JSON"
	AS112TopIPVersionParamsFormatCsv  AS112TopIPVersionParamsFormat = "CSV"
)

func (r AS112TopIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case AS112TopIPVersionParamsFormatJson, AS112TopIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type AS112TopIPVersionResponseEnvelope struct {
	Result  AS112TopIPVersionResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    as112TopIPVersionResponseEnvelopeJSON `json:"-"`
}

// as112TopIPVersionResponseEnvelopeJSON contains the JSON metadata for the struct
// [AS112TopIPVersionResponseEnvelope]
type as112TopIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112TopLocationsParams struct {
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
	Format param.Field[AS112TopLocationsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112TopLocationsParams]'s query parameters as
// `url.Values`.
func (r AS112TopLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AS112TopLocationsParamsFormat string

const (
	AS112TopLocationsParamsFormatJson AS112TopLocationsParamsFormat = "JSON"
	AS112TopLocationsParamsFormatCsv  AS112TopLocationsParamsFormat = "CSV"
)

func (r AS112TopLocationsParamsFormat) IsKnown() bool {
	switch r {
	case AS112TopLocationsParamsFormatJson, AS112TopLocationsParamsFormatCsv:
		return true
	}
	return false
}

type AS112TopLocationsResponseEnvelope struct {
	Result  AS112TopLocationsResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    as112TopLocationsResponseEnvelopeJSON `json:"-"`
}

// as112TopLocationsResponseEnvelopeJSON contains the JSON metadata for the struct
// [AS112TopLocationsResponseEnvelope]
type as112TopLocationsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TopLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TopLocationsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
