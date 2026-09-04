// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// AnnotationService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnnotationService] method instead.
type AnnotationService struct {
	Options []option.RequestOption
	Outages *AnnotationOutageService
}

// NewAnnotationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAnnotationService(opts ...option.RequestOption) (r *AnnotationService) {
	r = &AnnotationService{}
	r.Options = opts
	r.Outages = NewAnnotationOutageService(opts...)
	return
}

// Retrieves the latest annotations.
func (r *AnnotationService) List(ctx context.Context, query AnnotationListParams, opts ...option.RequestOption) (res *AnnotationListResponse, err error) {
	var env AnnotationListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/annotations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AnnotationListResponse struct {
	Annotations []AnnotationListResponseAnnotation `json:"annotations" api:"required"`
	JSON        annotationListResponseJSON         `json:"-"`
}

// annotationListResponseJSON contains the JSON metadata for the struct
// [AnnotationListResponse]
type annotationListResponseJSON struct {
	Annotations apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationListResponseJSON) RawJSON() string {
	return r.raw
}

type AnnotationListResponseAnnotation struct {
	ID               string                                             `json:"id" api:"required"`
	ASNs             []int64                                            `json:"asns" api:"required"`
	ASNsDetails      []AnnotationListResponseAnnotationsASNsDetail      `json:"asnsDetails" api:"required"`
	DataSource       string                                             `json:"dataSource" api:"required"`
	Description      string                                             `json:"description" api:"required,nullable"`
	EndDate          string                                             `json:"endDate" api:"required,nullable"`
	Entities         []AnnotationListResponseAnnotationsEntity          `json:"entities" api:"required"`
	EventType        string                                             `json:"eventType" api:"required"`
	GeoIDs           []string                                           `json:"geoIds" api:"required"`
	LinkedURL        string                                             `json:"linkedUrl" api:"required,nullable"`
	Locations        []string                                           `json:"locations" api:"required"`
	LocationsDetails []AnnotationListResponseAnnotationsLocationsDetail `json:"locationsDetails" api:"required"`
	Origins          []string                                           `json:"origins" api:"required"`
	OriginsDetails   []AnnotationListResponseAnnotationsOriginsDetail   `json:"originsDetails" api:"required"`
	Outage           AnnotationListResponseAnnotationsOutage            `json:"outage" api:"required,nullable"`
	Scope            string                                             `json:"scope" api:"required,nullable"`
	StartDate        string                                             `json:"startDate" api:"required"`
	Tags             []string                                           `json:"tags" api:"required"`
	JSON             annotationListResponseAnnotationJSON               `json:"-"`
}

// annotationListResponseAnnotationJSON contains the JSON metadata for the struct
// [AnnotationListResponseAnnotation]
type annotationListResponseAnnotationJSON struct {
	ID               apijson.Field
	ASNs             apijson.Field
	ASNsDetails      apijson.Field
	DataSource       apijson.Field
	Description      apijson.Field
	EndDate          apijson.Field
	Entities         apijson.Field
	EventType        apijson.Field
	GeoIDs           apijson.Field
	LinkedURL        apijson.Field
	Locations        apijson.Field
	LocationsDetails apijson.Field
	Origins          apijson.Field
	OriginsDetails   apijson.Field
	Outage           apijson.Field
	Scope            apijson.Field
	StartDate        apijson.Field
	Tags             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AnnotationListResponseAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationListResponseAnnotationJSON) RawJSON() string {
	return r.raw
}

type AnnotationListResponseAnnotationsASNsDetail struct {
	ASN      string                                               `json:"asn" api:"required"`
	Location AnnotationListResponseAnnotationsASNsDetailsLocation `json:"location" api:"required,nullable"`
	Name     string                                               `json:"name" api:"required,nullable"`
	JSON     annotationListResponseAnnotationsASNsDetailJSON      `json:"-"`
}

// annotationListResponseAnnotationsASNsDetailJSON contains the JSON metadata for
// the struct [AnnotationListResponseAnnotationsASNsDetail]
type annotationListResponseAnnotationsASNsDetailJSON struct {
	ASN         apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationListResponseAnnotationsASNsDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationListResponseAnnotationsASNsDetailJSON) RawJSON() string {
	return r.raw
}

type AnnotationListResponseAnnotationsASNsDetailsLocation struct {
	Code string                                                   `json:"code" api:"required"`
	Name string                                                   `json:"name" api:"required"`
	JSON annotationListResponseAnnotationsASNsDetailsLocationJSON `json:"-"`
}

// annotationListResponseAnnotationsASNsDetailsLocationJSON contains the JSON
// metadata for the struct [AnnotationListResponseAnnotationsASNsDetailsLocation]
type annotationListResponseAnnotationsASNsDetailsLocationJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationListResponseAnnotationsASNsDetailsLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationListResponseAnnotationsASNsDetailsLocationJSON) RawJSON() string {
	return r.raw
}

type AnnotationListResponseAnnotationsEntity struct {
	EntityName  string                                      `json:"entityName" api:"required,nullable"`
	EntityType  string                                      `json:"entityType" api:"required"`
	EntityValue string                                      `json:"entityValue" api:"required"`
	JSON        annotationListResponseAnnotationsEntityJSON `json:"-"`
}

// annotationListResponseAnnotationsEntityJSON contains the JSON metadata for the
// struct [AnnotationListResponseAnnotationsEntity]
type annotationListResponseAnnotationsEntityJSON struct {
	EntityName  apijson.Field
	EntityType  apijson.Field
	EntityValue apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationListResponseAnnotationsEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationListResponseAnnotationsEntityJSON) RawJSON() string {
	return r.raw
}

type AnnotationListResponseAnnotationsLocationsDetail struct {
	Code string                                               `json:"code" api:"required"`
	Name string                                               `json:"name" api:"required"`
	JSON annotationListResponseAnnotationsLocationsDetailJSON `json:"-"`
}

// annotationListResponseAnnotationsLocationsDetailJSON contains the JSON metadata
// for the struct [AnnotationListResponseAnnotationsLocationsDetail]
type annotationListResponseAnnotationsLocationsDetailJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationListResponseAnnotationsLocationsDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationListResponseAnnotationsLocationsDetailJSON) RawJSON() string {
	return r.raw
}

type AnnotationListResponseAnnotationsOriginsDetail struct {
	Name   string                                             `json:"name" api:"required,nullable"`
	Origin string                                             `json:"origin" api:"required"`
	JSON   annotationListResponseAnnotationsOriginsDetailJSON `json:"-"`
}

// annotationListResponseAnnotationsOriginsDetailJSON contains the JSON metadata
// for the struct [AnnotationListResponseAnnotationsOriginsDetail]
type annotationListResponseAnnotationsOriginsDetailJSON struct {
	Name        apijson.Field
	Origin      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationListResponseAnnotationsOriginsDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationListResponseAnnotationsOriginsDetailJSON) RawJSON() string {
	return r.raw
}

type AnnotationListResponseAnnotationsOutage struct {
	OutageCause string                                      `json:"outageCause" api:"required"`
	OutageType  string                                      `json:"outageType" api:"required"`
	JSON        annotationListResponseAnnotationsOutageJSON `json:"-"`
}

// annotationListResponseAnnotationsOutageJSON contains the JSON metadata for the
// struct [AnnotationListResponseAnnotationsOutage]
type annotationListResponseAnnotationsOutageJSON struct {
	OutageCause apijson.Field
	OutageType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationListResponseAnnotationsOutage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationListResponseAnnotationsOutageJSON) RawJSON() string {
	return r.raw
}

type AnnotationListParams struct {
	// Filters results by Autonomous System. Specify a single Autonomous System Number
	// (ASN) as integer.
	ASN param.Field[int64] `query:"asn"`
	// Filters results by bot.
	Bot param.Field[string] `query:"bot"`
	// Filters results by certificate authority.
	CA param.Field[string] `query:"ca"`
	// Filters results by data source.
	DataSource param.Field[AnnotationListParamsDataSource] `query:"dataSource"`
	// End of the date range (inclusive). Alternative to `dateRange`; provide together
	// with `dateStart`.
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by a relative date range ending at the current time. Use `<n>d`
	// for days (up to `364d`) or `<n>w` for weeks (up to `52w`), e.g. `7d`. Append
	// `control` to request the equivalent previous period for comparison: the
	// comparison window is shifted back by the current window's length rounded up to a
	// whole number of weeks, so it keeps the same weekday alignment and does not
	// overlap the current window (e.g. `3dcontrol` covers days -10 to -7, `7dcontrol`
	// covers days -14 to -7, `28dcontrol` covers days -56 to -28, and `10dcontrol`
	// covers days -24 to -14). Mutually exclusive with `dateStart`/`dateEnd`.
	DateRange param.Field[string] `query:"dateRange"`
	// Start of the date range (inclusive). Alternative to `dateRange`; provide
	// together with `dateEnd`.
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by event type. EVENT is a legacy alias for GENERAL.
	EventType param.Field[AnnotationListParamsEventType] `query:"eventType"`
	// Format in which results will be returned.
	Format param.Field[AnnotationListParamsFormat] `query:"format"`
	// Filters results by geolocation. Refer to
	// [GeoNames](https://download.geonames.org/export/dump/readme.txt).
	GeoID param.Field[string] `query:"geoId"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify an alpha-2 location code.
	Location param.Field[string] `query:"location"`
	// Filters results by certificate log.
	Log param.Field[string] `query:"log"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64] `query:"offset"`
	// Filters results by origin.
	Origin param.Field[string] `query:"origin"`
	// Filters results by outage cause.
	OutageCause param.Field[AnnotationListParamsOutageCause] `query:"outageCause"`
	// Filters results by outage type.
	OutageType param.Field[AnnotationListParamsOutageType] `query:"outageType"`
	// Filters results by a free-text match on the annotation description, id, or
	// linked entities (location, ASN, origin).
	Query param.Field[string] `query:"query"`
	// Filters results by annotation tag. Matches annotations carrying at least one of
	// the given tags.
	Tags param.Field[[]AnnotationListParamsTag] `query:"tags"`
	// Filters results by top-level domain.
	TLD param.Field[string] `query:"tld"`
}

// URLQuery serializes [AnnotationListParams]'s query parameters as `url.Values`.
func (r AnnotationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filters results by data source.
type AnnotationListParamsDataSource string

const (
	AnnotationListParamsDataSourceAll                AnnotationListParamsDataSource = "ALL"
	AnnotationListParamsDataSourceAIBots             AnnotationListParamsDataSource = "AI_BOTS"
	AnnotationListParamsDataSourceAIGateway          AnnotationListParamsDataSource = "AI_GATEWAY"
	AnnotationListParamsDataSourceBGP                AnnotationListParamsDataSource = "BGP"
	AnnotationListParamsDataSourceBots               AnnotationListParamsDataSource = "BOTS"
	AnnotationListParamsDataSourceConnectionAnomaly  AnnotationListParamsDataSource = "CONNECTION_ANOMALY"
	AnnotationListParamsDataSourceCT                 AnnotationListParamsDataSource = "CT"
	AnnotationListParamsDataSourceDNS                AnnotationListParamsDataSource = "DNS"
	AnnotationListParamsDataSourceDNSMagnitude       AnnotationListParamsDataSource = "DNS_MAGNITUDE"
	AnnotationListParamsDataSourceDNSAS112           AnnotationListParamsDataSource = "DNS_AS112"
	AnnotationListParamsDataSourceDos                AnnotationListParamsDataSource = "DOS"
	AnnotationListParamsDataSourceEmailRouting       AnnotationListParamsDataSource = "EMAIL_ROUTING"
	AnnotationListParamsDataSourceEmailSecurity      AnnotationListParamsDataSource = "EMAIL_SECURITY"
	AnnotationListParamsDataSourceFw                 AnnotationListParamsDataSource = "FW"
	AnnotationListParamsDataSourceFwPg               AnnotationListParamsDataSource = "FW_PG"
	AnnotationListParamsDataSourceHTTP               AnnotationListParamsDataSource = "HTTP"
	AnnotationListParamsDataSourceHTTPControl        AnnotationListParamsDataSource = "HTTP_CONTROL"
	AnnotationListParamsDataSourceHTTPCrawlerReferer AnnotationListParamsDataSource = "HTTP_CRAWLER_REFERER"
	AnnotationListParamsDataSourceHTTPOrigins        AnnotationListParamsDataSource = "HTTP_ORIGINS"
	AnnotationListParamsDataSourceIQI                AnnotationListParamsDataSource = "IQI"
	AnnotationListParamsDataSourceLeakedCredentials  AnnotationListParamsDataSource = "LEAKED_CREDENTIALS"
	AnnotationListParamsDataSourceNet                AnnotationListParamsDataSource = "NET"
	AnnotationListParamsDataSourceRobotsTXT          AnnotationListParamsDataSource = "ROBOTS_TXT"
	AnnotationListParamsDataSourceSpeed              AnnotationListParamsDataSource = "SPEED"
	AnnotationListParamsDataSourceWorkersAI          AnnotationListParamsDataSource = "WORKERS_AI"
)

func (r AnnotationListParamsDataSource) IsKnown() bool {
	switch r {
	case AnnotationListParamsDataSourceAll, AnnotationListParamsDataSourceAIBots, AnnotationListParamsDataSourceAIGateway, AnnotationListParamsDataSourceBGP, AnnotationListParamsDataSourceBots, AnnotationListParamsDataSourceConnectionAnomaly, AnnotationListParamsDataSourceCT, AnnotationListParamsDataSourceDNS, AnnotationListParamsDataSourceDNSMagnitude, AnnotationListParamsDataSourceDNSAS112, AnnotationListParamsDataSourceDos, AnnotationListParamsDataSourceEmailRouting, AnnotationListParamsDataSourceEmailSecurity, AnnotationListParamsDataSourceFw, AnnotationListParamsDataSourceFwPg, AnnotationListParamsDataSourceHTTP, AnnotationListParamsDataSourceHTTPControl, AnnotationListParamsDataSourceHTTPCrawlerReferer, AnnotationListParamsDataSourceHTTPOrigins, AnnotationListParamsDataSourceIQI, AnnotationListParamsDataSourceLeakedCredentials, AnnotationListParamsDataSourceNet, AnnotationListParamsDataSourceRobotsTXT, AnnotationListParamsDataSourceSpeed, AnnotationListParamsDataSourceWorkersAI:
		return true
	}
	return false
}

// Filters results by event type. EVENT is a legacy alias for GENERAL.
type AnnotationListParamsEventType string

const (
	AnnotationListParamsEventTypeEvent             AnnotationListParamsEventType = "EVENT"
	AnnotationListParamsEventTypeGeneral           AnnotationListParamsEventType = "GENERAL"
	AnnotationListParamsEventTypeOutage            AnnotationListParamsEventType = "OUTAGE"
	AnnotationListParamsEventTypePartialProjection AnnotationListParamsEventType = "PARTIAL_PROJECTION"
	AnnotationListParamsEventTypePipeline          AnnotationListParamsEventType = "PIPELINE"
	AnnotationListParamsEventTypeTrafficAnomaly    AnnotationListParamsEventType = "TRAFFIC_ANOMALY"
)

func (r AnnotationListParamsEventType) IsKnown() bool {
	switch r {
	case AnnotationListParamsEventTypeEvent, AnnotationListParamsEventTypeGeneral, AnnotationListParamsEventTypeOutage, AnnotationListParamsEventTypePartialProjection, AnnotationListParamsEventTypePipeline, AnnotationListParamsEventTypeTrafficAnomaly:
		return true
	}
	return false
}

// Format in which results will be returned.
type AnnotationListParamsFormat string

const (
	AnnotationListParamsFormatJson AnnotationListParamsFormat = "JSON"
	AnnotationListParamsFormatCsv  AnnotationListParamsFormat = "CSV"
)

func (r AnnotationListParamsFormat) IsKnown() bool {
	switch r {
	case AnnotationListParamsFormatJson, AnnotationListParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by outage cause.
type AnnotationListParamsOutageCause string

const (
	AnnotationListParamsOutageCauseBlocking           AnnotationListParamsOutageCause = "BLOCKING"
	AnnotationListParamsOutageCauseCableCut           AnnotationListParamsOutageCause = "CABLE_CUT"
	AnnotationListParamsOutageCauseCyberattack        AnnotationListParamsOutageCause = "CYBERATTACK"
	AnnotationListParamsOutageCauseDNS                AnnotationListParamsOutageCause = "DNS"
	AnnotationListParamsOutageCauseFire               AnnotationListParamsOutageCause = "FIRE"
	AnnotationListParamsOutageCauseGovernmentDirected AnnotationListParamsOutageCause = "GOVERNMENT_DIRECTED"
	AnnotationListParamsOutageCauseMaintenance        AnnotationListParamsOutageCause = "MAINTENANCE"
	AnnotationListParamsOutageCauseMechanical         AnnotationListParamsOutageCause = "MECHANICAL"
	AnnotationListParamsOutageCauseMilitaryAction     AnnotationListParamsOutageCause = "MILITARY_ACTION"
	AnnotationListParamsOutageCauseMisconfiguration   AnnotationListParamsOutageCause = "MISCONFIGURATION"
	AnnotationListParamsOutageCauseNaturalDisaster    AnnotationListParamsOutageCause = "NATURAL_DISASTER"
	AnnotationListParamsOutageCauseNetworkProblem     AnnotationListParamsOutageCause = "NETWORK_PROBLEM"
	AnnotationListParamsOutageCausePowerOutage        AnnotationListParamsOutageCause = "POWER_OUTAGE"
	AnnotationListParamsOutageCauseSoftware           AnnotationListParamsOutageCause = "SOFTWARE"
	AnnotationListParamsOutageCauseTechnicalProblem   AnnotationListParamsOutageCause = "TECHNICAL_PROBLEM"
	AnnotationListParamsOutageCauseUnknown            AnnotationListParamsOutageCause = "UNKNOWN"
	AnnotationListParamsOutageCauseWeather            AnnotationListParamsOutageCause = "WEATHER"
)

func (r AnnotationListParamsOutageCause) IsKnown() bool {
	switch r {
	case AnnotationListParamsOutageCauseBlocking, AnnotationListParamsOutageCauseCableCut, AnnotationListParamsOutageCauseCyberattack, AnnotationListParamsOutageCauseDNS, AnnotationListParamsOutageCauseFire, AnnotationListParamsOutageCauseGovernmentDirected, AnnotationListParamsOutageCauseMaintenance, AnnotationListParamsOutageCauseMechanical, AnnotationListParamsOutageCauseMilitaryAction, AnnotationListParamsOutageCauseMisconfiguration, AnnotationListParamsOutageCauseNaturalDisaster, AnnotationListParamsOutageCauseNetworkProblem, AnnotationListParamsOutageCausePowerOutage, AnnotationListParamsOutageCauseSoftware, AnnotationListParamsOutageCauseTechnicalProblem, AnnotationListParamsOutageCauseUnknown, AnnotationListParamsOutageCauseWeather:
		return true
	}
	return false
}

// Filters results by outage type.
type AnnotationListParamsOutageType string

const (
	AnnotationListParamsOutageTypeNationwide AnnotationListParamsOutageType = "NATIONWIDE"
	AnnotationListParamsOutageTypeRegional   AnnotationListParamsOutageType = "REGIONAL"
	AnnotationListParamsOutageTypeNetwork    AnnotationListParamsOutageType = "NETWORK"
	AnnotationListParamsOutageTypePlatform   AnnotationListParamsOutageType = "PLATFORM"
)

func (r AnnotationListParamsOutageType) IsKnown() bool {
	switch r {
	case AnnotationListParamsOutageTypeNationwide, AnnotationListParamsOutageTypeRegional, AnnotationListParamsOutageTypeNetwork, AnnotationListParamsOutageTypePlatform:
		return true
	}
	return false
}

// Dimension tag an annotation applies to.
type AnnotationListParamsTag string

const (
	AnnotationListParamsTagAdm1                     AnnotationListParamsTag = "ADM1"
	AnnotationListParamsTagAdm2                     AnnotationListParamsTag = "ADM2"
	AnnotationListParamsTagAPITraffic               AnnotationListParamsTag = "API_TRAFFIC"
	AnnotationListParamsTagARC                      AnnotationListParamsTag = "ARC"
	AnnotationListParamsTagAs                       AnnotationListParamsTag = "AS"
	AnnotationListParamsTagASN                      AnnotationListParamsTag = "ASN"
	AnnotationListParamsTagAttacks                  AnnotationListParamsTag = "ATTACKS"
	AnnotationListParamsTagAuthor                   AnnotationListParamsTag = "AUTHOR"
	AnnotationListParamsTagBandwidth                AnnotationListParamsTag = "BANDWIDTH"
	AnnotationListParamsTagBitrate                  AnnotationListParamsTag = "BITRATE"
	AnnotationListParamsTagBot                      AnnotationListParamsTag = "BOT"
	AnnotationListParamsTagBotCategory              AnnotationListParamsTag = "BOT_CATEGORY"
	AnnotationListParamsTagBotClass                 AnnotationListParamsTag = "BOT_CLASS"
	AnnotationListParamsTagBotKind                  AnnotationListParamsTag = "BOT_KIND"
	AnnotationListParamsTagBotOperator              AnnotationListParamsTag = "BOT_OPERATOR"
	AnnotationListParamsTagBrowser                  AnnotationListParamsTag = "BROWSER"
	AnnotationListParamsTagBrowserFamily            AnnotationListParamsTag = "BROWSER_FAMILY"
	AnnotationListParamsTagBytes                    AnnotationListParamsTag = "BYTES"
	AnnotationListParamsTagCA                       AnnotationListParamsTag = "CA"
	AnnotationListParamsTagCacheHit                 AnnotationListParamsTag = "CACHE_HIT"
	AnnotationListParamsTagCAOwner                  AnnotationListParamsTag = "CA_OWNER"
	AnnotationListParamsTagCheckResult              AnnotationListParamsTag = "CHECK_RESULT"
	AnnotationListParamsTagClientType               AnnotationListParamsTag = "CLIENT_TYPE"
	AnnotationListParamsTagCompromised              AnnotationListParamsTag = "COMPROMISED"
	AnnotationListParamsTagContentType              AnnotationListParamsTag = "CONTENT_TYPE"
	AnnotationListParamsTagCrawlPurpose             AnnotationListParamsTag = "CRAWL_PURPOSE"
	AnnotationListParamsTagCrawlReferRatio          AnnotationListParamsTag = "CRAWL_REFER_RATIO"
	AnnotationListParamsTagDeviceType               AnnotationListParamsTag = "DEVICE_TYPE"
	AnnotationListParamsTagDKIM                     AnnotationListParamsTag = "DKIM"
	AnnotationListParamsTagDMARC                    AnnotationListParamsTag = "DMARC"
	AnnotationListParamsTagDNS                      AnnotationListParamsTag = "DNS"
	AnnotationListParamsTagDNSSEC                   AnnotationListParamsTag = "DNSSEC"
	AnnotationListParamsTagDNSSECAware              AnnotationListParamsTag = "DNSSEC_AWARE"
	AnnotationListParamsTagDNSSECE2E                AnnotationListParamsTag = "DNSSEC_E2E"
	AnnotationListParamsTagDomainCategory           AnnotationListParamsTag = "DOMAIN_CATEGORY"
	AnnotationListParamsTagDuration                 AnnotationListParamsTag = "DURATION"
	AnnotationListParamsTagEdns                     AnnotationListParamsTag = "EDNS"
	AnnotationListParamsTagEncrypted                AnnotationListParamsTag = "ENCRYPTED"
	AnnotationListParamsTagEntryType                AnnotationListParamsTag = "ENTRY_TYPE"
	AnnotationListParamsTagExpirationStatus         AnnotationListParamsTag = "EXPIRATION_STATUS"
	AnnotationListParamsTagHasIPs                   AnnotationListParamsTag = "HAS_IPS"
	AnnotationListParamsTagHasMatchingAnswer        AnnotationListParamsTag = "HAS_MATCHING_ANSWER"
	AnnotationListParamsTagHasWildcards             AnnotationListParamsTag = "HAS_WILDCARDS"
	AnnotationListParamsTagHTTPMethod               AnnotationListParamsTag = "HTTP_METHOD"
	AnnotationListParamsTagHTTPProtocol             AnnotationListParamsTag = "HTTP_PROTOCOL"
	AnnotationListParamsTagHTTPVersion              AnnotationListParamsTag = "HTTP_VERSION"
	AnnotationListParamsTagIndustry                 AnnotationListParamsTag = "INDUSTRY"
	AnnotationListParamsTagIPVersion                AnnotationListParamsTag = "IP_VERSION"
	AnnotationListParamsTagJitter                   AnnotationListParamsTag = "JITTER"
	AnnotationListParamsTagKeyAgreement             AnnotationListParamsTag = "KEY_AGREEMENT"
	AnnotationListParamsTagLatency                  AnnotationListParamsTag = "LATENCY"
	AnnotationListParamsTagLocation                 AnnotationListParamsTag = "LOCATION"
	AnnotationListParamsTagLocationLatency          AnnotationListParamsTag = "LOCATION_LATENCY"
	AnnotationListParamsTagLog                      AnnotationListParamsTag = "LOG"
	AnnotationListParamsTagLogAPI                   AnnotationListParamsTag = "LOG_API"
	AnnotationListParamsTagLogOperator              AnnotationListParamsTag = "LOG_OPERATOR"
	AnnotationListParamsTagMalicious                AnnotationListParamsTag = "MALICIOUS"
	AnnotationListParamsTagManagedRules             AnnotationListParamsTag = "MANAGED_RULES"
	AnnotationListParamsTagMitigationProduct        AnnotationListParamsTag = "MITIGATION_PRODUCT"
	AnnotationListParamsTagModel                    AnnotationListParamsTag = "MODEL"
	AnnotationListParamsTagNameserverLatency        AnnotationListParamsTag = "NAMESERVER_LATENCY"
	AnnotationListParamsTagOrigin                   AnnotationListParamsTag = "ORIGIN"
	AnnotationListParamsTagOriginAs                 AnnotationListParamsTag = "ORIGIN_AS"
	AnnotationListParamsTagOriginLocation           AnnotationListParamsTag = "ORIGIN_LOCATION"
	AnnotationListParamsTagOriginTargetLocationPair AnnotationListParamsTag = "ORIGIN_TARGET_LOCATION_PAIR"
	AnnotationListParamsTagOS                       AnnotationListParamsTag = "OS"
	AnnotationListParamsTagPercentile               AnnotationListParamsTag = "PERCENTILE"
	AnnotationListParamsTagPostQuantum              AnnotationListParamsTag = "POST_QUANTUM"
	AnnotationListParamsTagPrefix                   AnnotationListParamsTag = "PREFIX"
	AnnotationListParamsTagProduct                  AnnotationListParamsTag = "PRODUCT"
	AnnotationListParamsTagProtocol                 AnnotationListParamsTag = "PROTOCOL"
	AnnotationListParamsTagProvider                 AnnotationListParamsTag = "PROVIDER"
	AnnotationListParamsTagPublicKeyAlgorithm       AnnotationListParamsTag = "PUBLIC_KEY_ALGORITHM"
	AnnotationListParamsTagQueryType                AnnotationListParamsTag = "QUERY_TYPE"
	AnnotationListParamsTagReferer                  AnnotationListParamsTag = "REFERER"
	AnnotationListParamsTagRegion                   AnnotationListParamsTag = "REGION"
	AnnotationListParamsTagResponseCode             AnnotationListParamsTag = "RESPONSE_CODE"
	AnnotationListParamsTagResponseStatus           AnnotationListParamsTag = "RESPONSE_STATUS"
	AnnotationListParamsTagResponseStatusCategory   AnnotationListParamsTag = "RESPONSE_STATUS_CATEGORY"
	AnnotationListParamsTagResponseTTL              AnnotationListParamsTag = "RESPONSE_TTL"
	AnnotationListParamsTagSignatureAlgorithm       AnnotationListParamsTag = "SIGNATURE_ALGORITHM"
	AnnotationListParamsTagSpam                     AnnotationListParamsTag = "SPAM"
	AnnotationListParamsTagSPF                      AnnotationListParamsTag = "SPF"
	AnnotationListParamsTagSpoof                    AnnotationListParamsTag = "SPOOF"
	AnnotationListParamsTagSuccessRate              AnnotationListParamsTag = "SUCCESS_RATE"
	AnnotationListParamsTagTargetLocation           AnnotationListParamsTag = "TARGET_LOCATION"
	AnnotationListParamsTagTask                     AnnotationListParamsTag = "TASK"
	AnnotationListParamsTagThreatCategory           AnnotationListParamsTag = "THREAT_CATEGORY"
	AnnotationListParamsTagTLD                      AnnotationListParamsTag = "TLD"
	AnnotationListParamsTagTLDDNSMagnitude          AnnotationListParamsTag = "TLD_DNS_MAGNITUDE"
	AnnotationListParamsTagTLSVersion               AnnotationListParamsTag = "TLS_VERSION"
	AnnotationListParamsTagUpdateType               AnnotationListParamsTag = "UPDATE_TYPE"
	AnnotationListParamsTagUserAgent                AnnotationListParamsTag = "USER_AGENT"
	AnnotationListParamsTagValidationLevel          AnnotationListParamsTag = "VALIDATION_LEVEL"
	AnnotationListParamsTagVector                   AnnotationListParamsTag = "VECTOR"
	AnnotationListParamsTagVertical                 AnnotationListParamsTag = "VERTICAL"
)

func (r AnnotationListParamsTag) IsKnown() bool {
	switch r {
	case AnnotationListParamsTagAdm1, AnnotationListParamsTagAdm2, AnnotationListParamsTagAPITraffic, AnnotationListParamsTagARC, AnnotationListParamsTagAs, AnnotationListParamsTagASN, AnnotationListParamsTagAttacks, AnnotationListParamsTagAuthor, AnnotationListParamsTagBandwidth, AnnotationListParamsTagBitrate, AnnotationListParamsTagBot, AnnotationListParamsTagBotCategory, AnnotationListParamsTagBotClass, AnnotationListParamsTagBotKind, AnnotationListParamsTagBotOperator, AnnotationListParamsTagBrowser, AnnotationListParamsTagBrowserFamily, AnnotationListParamsTagBytes, AnnotationListParamsTagCA, AnnotationListParamsTagCacheHit, AnnotationListParamsTagCAOwner, AnnotationListParamsTagCheckResult, AnnotationListParamsTagClientType, AnnotationListParamsTagCompromised, AnnotationListParamsTagContentType, AnnotationListParamsTagCrawlPurpose, AnnotationListParamsTagCrawlReferRatio, AnnotationListParamsTagDeviceType, AnnotationListParamsTagDKIM, AnnotationListParamsTagDMARC, AnnotationListParamsTagDNS, AnnotationListParamsTagDNSSEC, AnnotationListParamsTagDNSSECAware, AnnotationListParamsTagDNSSECE2E, AnnotationListParamsTagDomainCategory, AnnotationListParamsTagDuration, AnnotationListParamsTagEdns, AnnotationListParamsTagEncrypted, AnnotationListParamsTagEntryType, AnnotationListParamsTagExpirationStatus, AnnotationListParamsTagHasIPs, AnnotationListParamsTagHasMatchingAnswer, AnnotationListParamsTagHasWildcards, AnnotationListParamsTagHTTPMethod, AnnotationListParamsTagHTTPProtocol, AnnotationListParamsTagHTTPVersion, AnnotationListParamsTagIndustry, AnnotationListParamsTagIPVersion, AnnotationListParamsTagJitter, AnnotationListParamsTagKeyAgreement, AnnotationListParamsTagLatency, AnnotationListParamsTagLocation, AnnotationListParamsTagLocationLatency, AnnotationListParamsTagLog, AnnotationListParamsTagLogAPI, AnnotationListParamsTagLogOperator, AnnotationListParamsTagMalicious, AnnotationListParamsTagManagedRules, AnnotationListParamsTagMitigationProduct, AnnotationListParamsTagModel, AnnotationListParamsTagNameserverLatency, AnnotationListParamsTagOrigin, AnnotationListParamsTagOriginAs, AnnotationListParamsTagOriginLocation, AnnotationListParamsTagOriginTargetLocationPair, AnnotationListParamsTagOS, AnnotationListParamsTagPercentile, AnnotationListParamsTagPostQuantum, AnnotationListParamsTagPrefix, AnnotationListParamsTagProduct, AnnotationListParamsTagProtocol, AnnotationListParamsTagProvider, AnnotationListParamsTagPublicKeyAlgorithm, AnnotationListParamsTagQueryType, AnnotationListParamsTagReferer, AnnotationListParamsTagRegion, AnnotationListParamsTagResponseCode, AnnotationListParamsTagResponseStatus, AnnotationListParamsTagResponseStatusCategory, AnnotationListParamsTagResponseTTL, AnnotationListParamsTagSignatureAlgorithm, AnnotationListParamsTagSpam, AnnotationListParamsTagSPF, AnnotationListParamsTagSpoof, AnnotationListParamsTagSuccessRate, AnnotationListParamsTagTargetLocation, AnnotationListParamsTagTask, AnnotationListParamsTagThreatCategory, AnnotationListParamsTagTLD, AnnotationListParamsTagTLDDNSMagnitude, AnnotationListParamsTagTLSVersion, AnnotationListParamsTagUpdateType, AnnotationListParamsTagUserAgent, AnnotationListParamsTagValidationLevel, AnnotationListParamsTagVector, AnnotationListParamsTagVertical:
		return true
	}
	return false
}

type AnnotationListResponseEnvelope struct {
	Result  AnnotationListResponse             `json:"result" api:"required"`
	Success bool                               `json:"success" api:"required"`
	JSON    annotationListResponseEnvelopeJSON `json:"-"`
}

// annotationListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AnnotationListResponseEnvelope]
type annotationListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
