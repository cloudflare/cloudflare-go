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

// AnnotationOutageService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnnotationOutageService] method instead.
type AnnotationOutageService struct {
	Options []option.RequestOption
}

// NewAnnotationOutageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAnnotationOutageService(opts ...option.RequestOption) (r *AnnotationOutageService) {
	r = &AnnotationOutageService{}
	r.Options = opts
	return
}

// Retrieves the latest Internet outages and anomalies.
func (r *AnnotationOutageService) Get(ctx context.Context, query AnnotationOutageGetParams, opts ...option.RequestOption) (res *AnnotationOutageGetResponse, err error) {
	var env AnnotationOutageGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/annotations/outages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves the number of outages by location.
func (r *AnnotationOutageService) Locations(ctx context.Context, query AnnotationOutageLocationsParams, opts ...option.RequestOption) (res *AnnotationOutageLocationsResponse, err error) {
	var env AnnotationOutageLocationsResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/annotations/outages/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AnnotationOutageGetResponse struct {
	Annotations []AnnotationOutageGetResponseAnnotation `json:"annotations" api:"required"`
	JSON        annotationOutageGetResponseJSON         `json:"-"`
}

// annotationOutageGetResponseJSON contains the JSON metadata for the struct
// [AnnotationOutageGetResponse]
type annotationOutageGetResponseJSON struct {
	Annotations apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageGetResponseAnnotation struct {
	ID               string                                                  `json:"id" api:"required"`
	ASNs             []int64                                                 `json:"asns" api:"required"`
	ASNsDetails      []AnnotationOutageGetResponseAnnotationsASNsDetail      `json:"asnsDetails" api:"required"`
	DataSource       string                                                  `json:"dataSource" api:"required"`
	Description      string                                                  `json:"description" api:"required,nullable"`
	EndDate          time.Time                                               `json:"endDate" api:"required,nullable" format:"date-time"`
	Entities         []AnnotationOutageGetResponseAnnotationsEntity          `json:"entities" api:"required"`
	EventType        string                                                  `json:"eventType" api:"required"`
	GeoIDs           []string                                                `json:"geoIds" api:"required"`
	LinkedURL        string                                                  `json:"linkedUrl" api:"required,nullable"`
	Locations        []string                                                `json:"locations" api:"required"`
	LocationsDetails []AnnotationOutageGetResponseAnnotationsLocationsDetail `json:"locationsDetails" api:"required"`
	Origins          []string                                                `json:"origins" api:"required"`
	OriginsDetails   []AnnotationOutageGetResponseAnnotationsOriginsDetail   `json:"originsDetails" api:"required"`
	Outage           AnnotationOutageGetResponseAnnotationsOutage            `json:"outage" api:"required,nullable"`
	Scope            string                                                  `json:"scope" api:"required,nullable"`
	StartDate        time.Time                                               `json:"startDate" api:"required" format:"date-time"`
	Tags             []string                                                `json:"tags" api:"required"`
	JSON             annotationOutageGetResponseAnnotationJSON               `json:"-"`
}

// annotationOutageGetResponseAnnotationJSON contains the JSON metadata for the
// struct [AnnotationOutageGetResponseAnnotation]
type annotationOutageGetResponseAnnotationJSON struct {
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

func (r *AnnotationOutageGetResponseAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseAnnotationJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageGetResponseAnnotationsASNsDetail struct {
	ASN      string                                                    `json:"asn" api:"required"`
	Location AnnotationOutageGetResponseAnnotationsASNsDetailsLocation `json:"location" api:"required,nullable"`
	Name     string                                                    `json:"name" api:"required,nullable"`
	JSON     annotationOutageGetResponseAnnotationsASNsDetailJSON      `json:"-"`
}

// annotationOutageGetResponseAnnotationsASNsDetailJSON contains the JSON metadata
// for the struct [AnnotationOutageGetResponseAnnotationsASNsDetail]
type annotationOutageGetResponseAnnotationsASNsDetailJSON struct {
	ASN         apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageGetResponseAnnotationsASNsDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseAnnotationsASNsDetailJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageGetResponseAnnotationsASNsDetailsLocation struct {
	Code string                                                        `json:"code" api:"required"`
	Name string                                                        `json:"name" api:"required"`
	JSON annotationOutageGetResponseAnnotationsASNsDetailsLocationJSON `json:"-"`
}

// annotationOutageGetResponseAnnotationsASNsDetailsLocationJSON contains the JSON
// metadata for the struct
// [AnnotationOutageGetResponseAnnotationsASNsDetailsLocation]
type annotationOutageGetResponseAnnotationsASNsDetailsLocationJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageGetResponseAnnotationsASNsDetailsLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseAnnotationsASNsDetailsLocationJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageGetResponseAnnotationsEntity struct {
	EntityName  string                                           `json:"entityName" api:"required,nullable"`
	EntityType  string                                           `json:"entityType" api:"required"`
	EntityValue string                                           `json:"entityValue" api:"required"`
	JSON        annotationOutageGetResponseAnnotationsEntityJSON `json:"-"`
}

// annotationOutageGetResponseAnnotationsEntityJSON contains the JSON metadata for
// the struct [AnnotationOutageGetResponseAnnotationsEntity]
type annotationOutageGetResponseAnnotationsEntityJSON struct {
	EntityName  apijson.Field
	EntityType  apijson.Field
	EntityValue apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageGetResponseAnnotationsEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseAnnotationsEntityJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageGetResponseAnnotationsLocationsDetail struct {
	Code string                                                    `json:"code" api:"required"`
	Name string                                                    `json:"name" api:"required"`
	JSON annotationOutageGetResponseAnnotationsLocationsDetailJSON `json:"-"`
}

// annotationOutageGetResponseAnnotationsLocationsDetailJSON contains the JSON
// metadata for the struct [AnnotationOutageGetResponseAnnotationsLocationsDetail]
type annotationOutageGetResponseAnnotationsLocationsDetailJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageGetResponseAnnotationsLocationsDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseAnnotationsLocationsDetailJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageGetResponseAnnotationsOriginsDetail struct {
	Name   string                                                  `json:"name" api:"required,nullable"`
	Origin string                                                  `json:"origin" api:"required"`
	JSON   annotationOutageGetResponseAnnotationsOriginsDetailJSON `json:"-"`
}

// annotationOutageGetResponseAnnotationsOriginsDetailJSON contains the JSON
// metadata for the struct [AnnotationOutageGetResponseAnnotationsOriginsDetail]
type annotationOutageGetResponseAnnotationsOriginsDetailJSON struct {
	Name        apijson.Field
	Origin      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageGetResponseAnnotationsOriginsDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseAnnotationsOriginsDetailJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageGetResponseAnnotationsOutage struct {
	OutageCause string                                           `json:"outageCause" api:"required"`
	OutageType  string                                           `json:"outageType" api:"required"`
	JSON        annotationOutageGetResponseAnnotationsOutageJSON `json:"-"`
}

// annotationOutageGetResponseAnnotationsOutageJSON contains the JSON metadata for
// the struct [AnnotationOutageGetResponseAnnotationsOutage]
type annotationOutageGetResponseAnnotationsOutageJSON struct {
	OutageCause apijson.Field
	OutageType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageGetResponseAnnotationsOutage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseAnnotationsOutageJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageLocationsResponse struct {
	Annotations []AnnotationOutageLocationsResponseAnnotation `json:"annotations" api:"required"`
	JSON        annotationOutageLocationsResponseJSON         `json:"-"`
}

// annotationOutageLocationsResponseJSON contains the JSON metadata for the struct
// [AnnotationOutageLocationsResponse]
type annotationOutageLocationsResponseJSON struct {
	Annotations apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageLocationsResponseJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageLocationsResponseAnnotation struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2" api:"required"`
	ClientCountryName   string `json:"clientCountryName" api:"required"`
	// A numeric string.
	Value string                                          `json:"value" api:"required"`
	JSON  annotationOutageLocationsResponseAnnotationJSON `json:"-"`
}

// annotationOutageLocationsResponseAnnotationJSON contains the JSON metadata for
// the struct [AnnotationOutageLocationsResponseAnnotation]
type annotationOutageLocationsResponseAnnotationJSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AnnotationOutageLocationsResponseAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageLocationsResponseAnnotationJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageGetParams struct {
	// Filters results by Autonomous System. Specify a single Autonomous System Number
	// (ASN) as integer.
	ASN param.Field[int64] `query:"asn"`
	// Filters results by bot.
	Bot param.Field[string] `query:"bot"`
	// Filters results by certificate authority.
	CA param.Field[string] `query:"ca"`
	// Filters results by data source.
	DataSource param.Field[AnnotationOutageGetParamsDataSource] `query:"dataSource"`
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
	// Format in which results will be returned.
	Format param.Field[AnnotationOutageGetParamsFormat] `query:"format"`
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
	OutageCause param.Field[AnnotationOutageGetParamsOutageCause] `query:"outageCause"`
	// Filters results by outage type.
	OutageType param.Field[AnnotationOutageGetParamsOutageType] `query:"outageType"`
	// Filters results by a free-text match on the annotation description, id, or
	// linked entities (location, ASN, origin).
	Query param.Field[string] `query:"query"`
	// Filters results by annotation tag. Matches annotations carrying at least one of
	// the given tags.
	Tags param.Field[[]AnnotationOutageGetParamsTag] `query:"tags"`
	// Filters results by top-level domain.
	TLD param.Field[string] `query:"tld"`
}

// URLQuery serializes [AnnotationOutageGetParams]'s query parameters as
// `url.Values`.
func (r AnnotationOutageGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filters results by data source.
type AnnotationOutageGetParamsDataSource string

const (
	AnnotationOutageGetParamsDataSourceAll                AnnotationOutageGetParamsDataSource = "ALL"
	AnnotationOutageGetParamsDataSourceAIBots             AnnotationOutageGetParamsDataSource = "AI_BOTS"
	AnnotationOutageGetParamsDataSourceAIGateway          AnnotationOutageGetParamsDataSource = "AI_GATEWAY"
	AnnotationOutageGetParamsDataSourceBGP                AnnotationOutageGetParamsDataSource = "BGP"
	AnnotationOutageGetParamsDataSourceBots               AnnotationOutageGetParamsDataSource = "BOTS"
	AnnotationOutageGetParamsDataSourceConnectionAnomaly  AnnotationOutageGetParamsDataSource = "CONNECTION_ANOMALY"
	AnnotationOutageGetParamsDataSourceCT                 AnnotationOutageGetParamsDataSource = "CT"
	AnnotationOutageGetParamsDataSourceDNS                AnnotationOutageGetParamsDataSource = "DNS"
	AnnotationOutageGetParamsDataSourceDNSMagnitude       AnnotationOutageGetParamsDataSource = "DNS_MAGNITUDE"
	AnnotationOutageGetParamsDataSourceDNSAS112           AnnotationOutageGetParamsDataSource = "DNS_AS112"
	AnnotationOutageGetParamsDataSourceDos                AnnotationOutageGetParamsDataSource = "DOS"
	AnnotationOutageGetParamsDataSourceEmailRouting       AnnotationOutageGetParamsDataSource = "EMAIL_ROUTING"
	AnnotationOutageGetParamsDataSourceEmailSecurity      AnnotationOutageGetParamsDataSource = "EMAIL_SECURITY"
	AnnotationOutageGetParamsDataSourceFw                 AnnotationOutageGetParamsDataSource = "FW"
	AnnotationOutageGetParamsDataSourceFwPg               AnnotationOutageGetParamsDataSource = "FW_PG"
	AnnotationOutageGetParamsDataSourceHTTP               AnnotationOutageGetParamsDataSource = "HTTP"
	AnnotationOutageGetParamsDataSourceHTTPControl        AnnotationOutageGetParamsDataSource = "HTTP_CONTROL"
	AnnotationOutageGetParamsDataSourceHTTPCrawlerReferer AnnotationOutageGetParamsDataSource = "HTTP_CRAWLER_REFERER"
	AnnotationOutageGetParamsDataSourceHTTPOrigins        AnnotationOutageGetParamsDataSource = "HTTP_ORIGINS"
	AnnotationOutageGetParamsDataSourceIQI                AnnotationOutageGetParamsDataSource = "IQI"
	AnnotationOutageGetParamsDataSourceLeakedCredentials  AnnotationOutageGetParamsDataSource = "LEAKED_CREDENTIALS"
	AnnotationOutageGetParamsDataSourceNet                AnnotationOutageGetParamsDataSource = "NET"
	AnnotationOutageGetParamsDataSourceRobotsTXT          AnnotationOutageGetParamsDataSource = "ROBOTS_TXT"
	AnnotationOutageGetParamsDataSourceSpeed              AnnotationOutageGetParamsDataSource = "SPEED"
	AnnotationOutageGetParamsDataSourceWorkersAI          AnnotationOutageGetParamsDataSource = "WORKERS_AI"
)

func (r AnnotationOutageGetParamsDataSource) IsKnown() bool {
	switch r {
	case AnnotationOutageGetParamsDataSourceAll, AnnotationOutageGetParamsDataSourceAIBots, AnnotationOutageGetParamsDataSourceAIGateway, AnnotationOutageGetParamsDataSourceBGP, AnnotationOutageGetParamsDataSourceBots, AnnotationOutageGetParamsDataSourceConnectionAnomaly, AnnotationOutageGetParamsDataSourceCT, AnnotationOutageGetParamsDataSourceDNS, AnnotationOutageGetParamsDataSourceDNSMagnitude, AnnotationOutageGetParamsDataSourceDNSAS112, AnnotationOutageGetParamsDataSourceDos, AnnotationOutageGetParamsDataSourceEmailRouting, AnnotationOutageGetParamsDataSourceEmailSecurity, AnnotationOutageGetParamsDataSourceFw, AnnotationOutageGetParamsDataSourceFwPg, AnnotationOutageGetParamsDataSourceHTTP, AnnotationOutageGetParamsDataSourceHTTPControl, AnnotationOutageGetParamsDataSourceHTTPCrawlerReferer, AnnotationOutageGetParamsDataSourceHTTPOrigins, AnnotationOutageGetParamsDataSourceIQI, AnnotationOutageGetParamsDataSourceLeakedCredentials, AnnotationOutageGetParamsDataSourceNet, AnnotationOutageGetParamsDataSourceRobotsTXT, AnnotationOutageGetParamsDataSourceSpeed, AnnotationOutageGetParamsDataSourceWorkersAI:
		return true
	}
	return false
}

// Format in which results will be returned.
type AnnotationOutageGetParamsFormat string

const (
	AnnotationOutageGetParamsFormatJson AnnotationOutageGetParamsFormat = "JSON"
	AnnotationOutageGetParamsFormatCsv  AnnotationOutageGetParamsFormat = "CSV"
)

func (r AnnotationOutageGetParamsFormat) IsKnown() bool {
	switch r {
	case AnnotationOutageGetParamsFormatJson, AnnotationOutageGetParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by outage cause.
type AnnotationOutageGetParamsOutageCause string

const (
	AnnotationOutageGetParamsOutageCauseBlocking           AnnotationOutageGetParamsOutageCause = "BLOCKING"
	AnnotationOutageGetParamsOutageCauseCableCut           AnnotationOutageGetParamsOutageCause = "CABLE_CUT"
	AnnotationOutageGetParamsOutageCauseCyberattack        AnnotationOutageGetParamsOutageCause = "CYBERATTACK"
	AnnotationOutageGetParamsOutageCauseDNS                AnnotationOutageGetParamsOutageCause = "DNS"
	AnnotationOutageGetParamsOutageCauseFire               AnnotationOutageGetParamsOutageCause = "FIRE"
	AnnotationOutageGetParamsOutageCauseGovernmentDirected AnnotationOutageGetParamsOutageCause = "GOVERNMENT_DIRECTED"
	AnnotationOutageGetParamsOutageCauseMaintenance        AnnotationOutageGetParamsOutageCause = "MAINTENANCE"
	AnnotationOutageGetParamsOutageCauseMechanical         AnnotationOutageGetParamsOutageCause = "MECHANICAL"
	AnnotationOutageGetParamsOutageCauseMilitaryAction     AnnotationOutageGetParamsOutageCause = "MILITARY_ACTION"
	AnnotationOutageGetParamsOutageCauseMisconfiguration   AnnotationOutageGetParamsOutageCause = "MISCONFIGURATION"
	AnnotationOutageGetParamsOutageCauseNaturalDisaster    AnnotationOutageGetParamsOutageCause = "NATURAL_DISASTER"
	AnnotationOutageGetParamsOutageCauseNetworkProblem     AnnotationOutageGetParamsOutageCause = "NETWORK_PROBLEM"
	AnnotationOutageGetParamsOutageCausePowerOutage        AnnotationOutageGetParamsOutageCause = "POWER_OUTAGE"
	AnnotationOutageGetParamsOutageCauseSoftware           AnnotationOutageGetParamsOutageCause = "SOFTWARE"
	AnnotationOutageGetParamsOutageCauseTechnicalProblem   AnnotationOutageGetParamsOutageCause = "TECHNICAL_PROBLEM"
	AnnotationOutageGetParamsOutageCauseUnknown            AnnotationOutageGetParamsOutageCause = "UNKNOWN"
	AnnotationOutageGetParamsOutageCauseWeather            AnnotationOutageGetParamsOutageCause = "WEATHER"
)

func (r AnnotationOutageGetParamsOutageCause) IsKnown() bool {
	switch r {
	case AnnotationOutageGetParamsOutageCauseBlocking, AnnotationOutageGetParamsOutageCauseCableCut, AnnotationOutageGetParamsOutageCauseCyberattack, AnnotationOutageGetParamsOutageCauseDNS, AnnotationOutageGetParamsOutageCauseFire, AnnotationOutageGetParamsOutageCauseGovernmentDirected, AnnotationOutageGetParamsOutageCauseMaintenance, AnnotationOutageGetParamsOutageCauseMechanical, AnnotationOutageGetParamsOutageCauseMilitaryAction, AnnotationOutageGetParamsOutageCauseMisconfiguration, AnnotationOutageGetParamsOutageCauseNaturalDisaster, AnnotationOutageGetParamsOutageCauseNetworkProblem, AnnotationOutageGetParamsOutageCausePowerOutage, AnnotationOutageGetParamsOutageCauseSoftware, AnnotationOutageGetParamsOutageCauseTechnicalProblem, AnnotationOutageGetParamsOutageCauseUnknown, AnnotationOutageGetParamsOutageCauseWeather:
		return true
	}
	return false
}

// Filters results by outage type.
type AnnotationOutageGetParamsOutageType string

const (
	AnnotationOutageGetParamsOutageTypeNationwide AnnotationOutageGetParamsOutageType = "NATIONWIDE"
	AnnotationOutageGetParamsOutageTypeRegional   AnnotationOutageGetParamsOutageType = "REGIONAL"
	AnnotationOutageGetParamsOutageTypeNetwork    AnnotationOutageGetParamsOutageType = "NETWORK"
	AnnotationOutageGetParamsOutageTypePlatform   AnnotationOutageGetParamsOutageType = "PLATFORM"
)

func (r AnnotationOutageGetParamsOutageType) IsKnown() bool {
	switch r {
	case AnnotationOutageGetParamsOutageTypeNationwide, AnnotationOutageGetParamsOutageTypeRegional, AnnotationOutageGetParamsOutageTypeNetwork, AnnotationOutageGetParamsOutageTypePlatform:
		return true
	}
	return false
}

// Dimension tag an annotation applies to.
type AnnotationOutageGetParamsTag string

const (
	AnnotationOutageGetParamsTagAdm1                     AnnotationOutageGetParamsTag = "ADM1"
	AnnotationOutageGetParamsTagAdm2                     AnnotationOutageGetParamsTag = "ADM2"
	AnnotationOutageGetParamsTagAPITraffic               AnnotationOutageGetParamsTag = "API_TRAFFIC"
	AnnotationOutageGetParamsTagARC                      AnnotationOutageGetParamsTag = "ARC"
	AnnotationOutageGetParamsTagAs                       AnnotationOutageGetParamsTag = "AS"
	AnnotationOutageGetParamsTagASN                      AnnotationOutageGetParamsTag = "ASN"
	AnnotationOutageGetParamsTagAttacks                  AnnotationOutageGetParamsTag = "ATTACKS"
	AnnotationOutageGetParamsTagAuthor                   AnnotationOutageGetParamsTag = "AUTHOR"
	AnnotationOutageGetParamsTagBandwidth                AnnotationOutageGetParamsTag = "BANDWIDTH"
	AnnotationOutageGetParamsTagBitrate                  AnnotationOutageGetParamsTag = "BITRATE"
	AnnotationOutageGetParamsTagBot                      AnnotationOutageGetParamsTag = "BOT"
	AnnotationOutageGetParamsTagBotCategory              AnnotationOutageGetParamsTag = "BOT_CATEGORY"
	AnnotationOutageGetParamsTagBotClass                 AnnotationOutageGetParamsTag = "BOT_CLASS"
	AnnotationOutageGetParamsTagBotKind                  AnnotationOutageGetParamsTag = "BOT_KIND"
	AnnotationOutageGetParamsTagBotOperator              AnnotationOutageGetParamsTag = "BOT_OPERATOR"
	AnnotationOutageGetParamsTagBrowser                  AnnotationOutageGetParamsTag = "BROWSER"
	AnnotationOutageGetParamsTagBrowserFamily            AnnotationOutageGetParamsTag = "BROWSER_FAMILY"
	AnnotationOutageGetParamsTagBytes                    AnnotationOutageGetParamsTag = "BYTES"
	AnnotationOutageGetParamsTagCA                       AnnotationOutageGetParamsTag = "CA"
	AnnotationOutageGetParamsTagCacheHit                 AnnotationOutageGetParamsTag = "CACHE_HIT"
	AnnotationOutageGetParamsTagCAOwner                  AnnotationOutageGetParamsTag = "CA_OWNER"
	AnnotationOutageGetParamsTagCheckResult              AnnotationOutageGetParamsTag = "CHECK_RESULT"
	AnnotationOutageGetParamsTagClientType               AnnotationOutageGetParamsTag = "CLIENT_TYPE"
	AnnotationOutageGetParamsTagCompromised              AnnotationOutageGetParamsTag = "COMPROMISED"
	AnnotationOutageGetParamsTagContentType              AnnotationOutageGetParamsTag = "CONTENT_TYPE"
	AnnotationOutageGetParamsTagCrawlPurpose             AnnotationOutageGetParamsTag = "CRAWL_PURPOSE"
	AnnotationOutageGetParamsTagCrawlReferRatio          AnnotationOutageGetParamsTag = "CRAWL_REFER_RATIO"
	AnnotationOutageGetParamsTagDeviceType               AnnotationOutageGetParamsTag = "DEVICE_TYPE"
	AnnotationOutageGetParamsTagDKIM                     AnnotationOutageGetParamsTag = "DKIM"
	AnnotationOutageGetParamsTagDMARC                    AnnotationOutageGetParamsTag = "DMARC"
	AnnotationOutageGetParamsTagDNS                      AnnotationOutageGetParamsTag = "DNS"
	AnnotationOutageGetParamsTagDNSSEC                   AnnotationOutageGetParamsTag = "DNSSEC"
	AnnotationOutageGetParamsTagDNSSECAware              AnnotationOutageGetParamsTag = "DNSSEC_AWARE"
	AnnotationOutageGetParamsTagDNSSECE2E                AnnotationOutageGetParamsTag = "DNSSEC_E2E"
	AnnotationOutageGetParamsTagDomainCategory           AnnotationOutageGetParamsTag = "DOMAIN_CATEGORY"
	AnnotationOutageGetParamsTagDuration                 AnnotationOutageGetParamsTag = "DURATION"
	AnnotationOutageGetParamsTagEdns                     AnnotationOutageGetParamsTag = "EDNS"
	AnnotationOutageGetParamsTagEncrypted                AnnotationOutageGetParamsTag = "ENCRYPTED"
	AnnotationOutageGetParamsTagEntryType                AnnotationOutageGetParamsTag = "ENTRY_TYPE"
	AnnotationOutageGetParamsTagExpirationStatus         AnnotationOutageGetParamsTag = "EXPIRATION_STATUS"
	AnnotationOutageGetParamsTagHasIPs                   AnnotationOutageGetParamsTag = "HAS_IPS"
	AnnotationOutageGetParamsTagHasMatchingAnswer        AnnotationOutageGetParamsTag = "HAS_MATCHING_ANSWER"
	AnnotationOutageGetParamsTagHasWildcards             AnnotationOutageGetParamsTag = "HAS_WILDCARDS"
	AnnotationOutageGetParamsTagHTTPMethod               AnnotationOutageGetParamsTag = "HTTP_METHOD"
	AnnotationOutageGetParamsTagHTTPProtocol             AnnotationOutageGetParamsTag = "HTTP_PROTOCOL"
	AnnotationOutageGetParamsTagHTTPVersion              AnnotationOutageGetParamsTag = "HTTP_VERSION"
	AnnotationOutageGetParamsTagIndustry                 AnnotationOutageGetParamsTag = "INDUSTRY"
	AnnotationOutageGetParamsTagIPVersion                AnnotationOutageGetParamsTag = "IP_VERSION"
	AnnotationOutageGetParamsTagJitter                   AnnotationOutageGetParamsTag = "JITTER"
	AnnotationOutageGetParamsTagKeyAgreement             AnnotationOutageGetParamsTag = "KEY_AGREEMENT"
	AnnotationOutageGetParamsTagLatency                  AnnotationOutageGetParamsTag = "LATENCY"
	AnnotationOutageGetParamsTagLocation                 AnnotationOutageGetParamsTag = "LOCATION"
	AnnotationOutageGetParamsTagLocationLatency          AnnotationOutageGetParamsTag = "LOCATION_LATENCY"
	AnnotationOutageGetParamsTagLog                      AnnotationOutageGetParamsTag = "LOG"
	AnnotationOutageGetParamsTagLogAPI                   AnnotationOutageGetParamsTag = "LOG_API"
	AnnotationOutageGetParamsTagLogOperator              AnnotationOutageGetParamsTag = "LOG_OPERATOR"
	AnnotationOutageGetParamsTagMalicious                AnnotationOutageGetParamsTag = "MALICIOUS"
	AnnotationOutageGetParamsTagManagedRules             AnnotationOutageGetParamsTag = "MANAGED_RULES"
	AnnotationOutageGetParamsTagMitigationProduct        AnnotationOutageGetParamsTag = "MITIGATION_PRODUCT"
	AnnotationOutageGetParamsTagModel                    AnnotationOutageGetParamsTag = "MODEL"
	AnnotationOutageGetParamsTagNameserverLatency        AnnotationOutageGetParamsTag = "NAMESERVER_LATENCY"
	AnnotationOutageGetParamsTagOrigin                   AnnotationOutageGetParamsTag = "ORIGIN"
	AnnotationOutageGetParamsTagOriginAs                 AnnotationOutageGetParamsTag = "ORIGIN_AS"
	AnnotationOutageGetParamsTagOriginLocation           AnnotationOutageGetParamsTag = "ORIGIN_LOCATION"
	AnnotationOutageGetParamsTagOriginTargetLocationPair AnnotationOutageGetParamsTag = "ORIGIN_TARGET_LOCATION_PAIR"
	AnnotationOutageGetParamsTagOS                       AnnotationOutageGetParamsTag = "OS"
	AnnotationOutageGetParamsTagPercentile               AnnotationOutageGetParamsTag = "PERCENTILE"
	AnnotationOutageGetParamsTagPostQuantum              AnnotationOutageGetParamsTag = "POST_QUANTUM"
	AnnotationOutageGetParamsTagPrefix                   AnnotationOutageGetParamsTag = "PREFIX"
	AnnotationOutageGetParamsTagProduct                  AnnotationOutageGetParamsTag = "PRODUCT"
	AnnotationOutageGetParamsTagProtocol                 AnnotationOutageGetParamsTag = "PROTOCOL"
	AnnotationOutageGetParamsTagProvider                 AnnotationOutageGetParamsTag = "PROVIDER"
	AnnotationOutageGetParamsTagPublicKeyAlgorithm       AnnotationOutageGetParamsTag = "PUBLIC_KEY_ALGORITHM"
	AnnotationOutageGetParamsTagQueryType                AnnotationOutageGetParamsTag = "QUERY_TYPE"
	AnnotationOutageGetParamsTagReferer                  AnnotationOutageGetParamsTag = "REFERER"
	AnnotationOutageGetParamsTagRegion                   AnnotationOutageGetParamsTag = "REGION"
	AnnotationOutageGetParamsTagResponseCode             AnnotationOutageGetParamsTag = "RESPONSE_CODE"
	AnnotationOutageGetParamsTagResponseStatus           AnnotationOutageGetParamsTag = "RESPONSE_STATUS"
	AnnotationOutageGetParamsTagResponseStatusCategory   AnnotationOutageGetParamsTag = "RESPONSE_STATUS_CATEGORY"
	AnnotationOutageGetParamsTagResponseTTL              AnnotationOutageGetParamsTag = "RESPONSE_TTL"
	AnnotationOutageGetParamsTagSignatureAlgorithm       AnnotationOutageGetParamsTag = "SIGNATURE_ALGORITHM"
	AnnotationOutageGetParamsTagSpam                     AnnotationOutageGetParamsTag = "SPAM"
	AnnotationOutageGetParamsTagSPF                      AnnotationOutageGetParamsTag = "SPF"
	AnnotationOutageGetParamsTagSpoof                    AnnotationOutageGetParamsTag = "SPOOF"
	AnnotationOutageGetParamsTagSuccessRate              AnnotationOutageGetParamsTag = "SUCCESS_RATE"
	AnnotationOutageGetParamsTagTargetLocation           AnnotationOutageGetParamsTag = "TARGET_LOCATION"
	AnnotationOutageGetParamsTagTask                     AnnotationOutageGetParamsTag = "TASK"
	AnnotationOutageGetParamsTagThreatCategory           AnnotationOutageGetParamsTag = "THREAT_CATEGORY"
	AnnotationOutageGetParamsTagTLD                      AnnotationOutageGetParamsTag = "TLD"
	AnnotationOutageGetParamsTagTLDDNSMagnitude          AnnotationOutageGetParamsTag = "TLD_DNS_MAGNITUDE"
	AnnotationOutageGetParamsTagTLSVersion               AnnotationOutageGetParamsTag = "TLS_VERSION"
	AnnotationOutageGetParamsTagUpdateType               AnnotationOutageGetParamsTag = "UPDATE_TYPE"
	AnnotationOutageGetParamsTagUserAgent                AnnotationOutageGetParamsTag = "USER_AGENT"
	AnnotationOutageGetParamsTagValidationLevel          AnnotationOutageGetParamsTag = "VALIDATION_LEVEL"
	AnnotationOutageGetParamsTagVector                   AnnotationOutageGetParamsTag = "VECTOR"
	AnnotationOutageGetParamsTagVertical                 AnnotationOutageGetParamsTag = "VERTICAL"
)

func (r AnnotationOutageGetParamsTag) IsKnown() bool {
	switch r {
	case AnnotationOutageGetParamsTagAdm1, AnnotationOutageGetParamsTagAdm2, AnnotationOutageGetParamsTagAPITraffic, AnnotationOutageGetParamsTagARC, AnnotationOutageGetParamsTagAs, AnnotationOutageGetParamsTagASN, AnnotationOutageGetParamsTagAttacks, AnnotationOutageGetParamsTagAuthor, AnnotationOutageGetParamsTagBandwidth, AnnotationOutageGetParamsTagBitrate, AnnotationOutageGetParamsTagBot, AnnotationOutageGetParamsTagBotCategory, AnnotationOutageGetParamsTagBotClass, AnnotationOutageGetParamsTagBotKind, AnnotationOutageGetParamsTagBotOperator, AnnotationOutageGetParamsTagBrowser, AnnotationOutageGetParamsTagBrowserFamily, AnnotationOutageGetParamsTagBytes, AnnotationOutageGetParamsTagCA, AnnotationOutageGetParamsTagCacheHit, AnnotationOutageGetParamsTagCAOwner, AnnotationOutageGetParamsTagCheckResult, AnnotationOutageGetParamsTagClientType, AnnotationOutageGetParamsTagCompromised, AnnotationOutageGetParamsTagContentType, AnnotationOutageGetParamsTagCrawlPurpose, AnnotationOutageGetParamsTagCrawlReferRatio, AnnotationOutageGetParamsTagDeviceType, AnnotationOutageGetParamsTagDKIM, AnnotationOutageGetParamsTagDMARC, AnnotationOutageGetParamsTagDNS, AnnotationOutageGetParamsTagDNSSEC, AnnotationOutageGetParamsTagDNSSECAware, AnnotationOutageGetParamsTagDNSSECE2E, AnnotationOutageGetParamsTagDomainCategory, AnnotationOutageGetParamsTagDuration, AnnotationOutageGetParamsTagEdns, AnnotationOutageGetParamsTagEncrypted, AnnotationOutageGetParamsTagEntryType, AnnotationOutageGetParamsTagExpirationStatus, AnnotationOutageGetParamsTagHasIPs, AnnotationOutageGetParamsTagHasMatchingAnswer, AnnotationOutageGetParamsTagHasWildcards, AnnotationOutageGetParamsTagHTTPMethod, AnnotationOutageGetParamsTagHTTPProtocol, AnnotationOutageGetParamsTagHTTPVersion, AnnotationOutageGetParamsTagIndustry, AnnotationOutageGetParamsTagIPVersion, AnnotationOutageGetParamsTagJitter, AnnotationOutageGetParamsTagKeyAgreement, AnnotationOutageGetParamsTagLatency, AnnotationOutageGetParamsTagLocation, AnnotationOutageGetParamsTagLocationLatency, AnnotationOutageGetParamsTagLog, AnnotationOutageGetParamsTagLogAPI, AnnotationOutageGetParamsTagLogOperator, AnnotationOutageGetParamsTagMalicious, AnnotationOutageGetParamsTagManagedRules, AnnotationOutageGetParamsTagMitigationProduct, AnnotationOutageGetParamsTagModel, AnnotationOutageGetParamsTagNameserverLatency, AnnotationOutageGetParamsTagOrigin, AnnotationOutageGetParamsTagOriginAs, AnnotationOutageGetParamsTagOriginLocation, AnnotationOutageGetParamsTagOriginTargetLocationPair, AnnotationOutageGetParamsTagOS, AnnotationOutageGetParamsTagPercentile, AnnotationOutageGetParamsTagPostQuantum, AnnotationOutageGetParamsTagPrefix, AnnotationOutageGetParamsTagProduct, AnnotationOutageGetParamsTagProtocol, AnnotationOutageGetParamsTagProvider, AnnotationOutageGetParamsTagPublicKeyAlgorithm, AnnotationOutageGetParamsTagQueryType, AnnotationOutageGetParamsTagReferer, AnnotationOutageGetParamsTagRegion, AnnotationOutageGetParamsTagResponseCode, AnnotationOutageGetParamsTagResponseStatus, AnnotationOutageGetParamsTagResponseStatusCategory, AnnotationOutageGetParamsTagResponseTTL, AnnotationOutageGetParamsTagSignatureAlgorithm, AnnotationOutageGetParamsTagSpam, AnnotationOutageGetParamsTagSPF, AnnotationOutageGetParamsTagSpoof, AnnotationOutageGetParamsTagSuccessRate, AnnotationOutageGetParamsTagTargetLocation, AnnotationOutageGetParamsTagTask, AnnotationOutageGetParamsTagThreatCategory, AnnotationOutageGetParamsTagTLD, AnnotationOutageGetParamsTagTLDDNSMagnitude, AnnotationOutageGetParamsTagTLSVersion, AnnotationOutageGetParamsTagUpdateType, AnnotationOutageGetParamsTagUserAgent, AnnotationOutageGetParamsTagValidationLevel, AnnotationOutageGetParamsTagVector, AnnotationOutageGetParamsTagVertical:
		return true
	}
	return false
}

type AnnotationOutageGetResponseEnvelope struct {
	Result  AnnotationOutageGetResponse             `json:"result" api:"required"`
	Success bool                                    `json:"success" api:"required"`
	JSON    annotationOutageGetResponseEnvelopeJSON `json:"-"`
}

// annotationOutageGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AnnotationOutageGetResponseEnvelope]
type annotationOutageGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageLocationsParams struct {
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
	// Format in which results will be returned.
	Format param.Field[AnnotationOutageLocationsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [AnnotationOutageLocationsParams]'s query parameters as
// `url.Values`.
func (r AnnotationOutageLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AnnotationOutageLocationsParamsFormat string

const (
	AnnotationOutageLocationsParamsFormatJson AnnotationOutageLocationsParamsFormat = "JSON"
	AnnotationOutageLocationsParamsFormatCsv  AnnotationOutageLocationsParamsFormat = "CSV"
)

func (r AnnotationOutageLocationsParamsFormat) IsKnown() bool {
	switch r {
	case AnnotationOutageLocationsParamsFormatJson, AnnotationOutageLocationsParamsFormatCsv:
		return true
	}
	return false
}

type AnnotationOutageLocationsResponseEnvelope struct {
	Result  AnnotationOutageLocationsResponse             `json:"result" api:"required"`
	Success bool                                          `json:"success" api:"required"`
	JSON    annotationOutageLocationsResponseEnvelopeJSON `json:"-"`
}

// annotationOutageLocationsResponseEnvelopeJSON contains the JSON metadata for the
// struct [AnnotationOutageLocationsResponseEnvelope]
type annotationOutageLocationsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageLocationsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
