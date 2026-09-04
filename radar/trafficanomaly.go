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

// TrafficAnomalyService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTrafficAnomalyService] method instead.
type TrafficAnomalyService struct {
	Options   []option.RequestOption
	Locations *TrafficAnomalyLocationService
}

// NewTrafficAnomalyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTrafficAnomalyService(opts ...option.RequestOption) (r *TrafficAnomalyService) {
	r = &TrafficAnomalyService{}
	r.Options = opts
	r.Locations = NewTrafficAnomalyLocationService(opts...)
	return
}

// Retrieves the latest Internet traffic anomalies, which are signals that might
// indicate an outage. These alerts are automatically detected by Radar and
// manually verified by our team.
func (r *TrafficAnomalyService) Get(ctx context.Context, query TrafficAnomalyGetParams, opts ...option.RequestOption) (res *TrafficAnomalyGetResponse, err error) {
	var env TrafficAnomalyGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/traffic_anomalies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type TrafficAnomalyGetResponse struct {
	TrafficAnomalies []TrafficAnomalyGetResponseTrafficAnomaly `json:"trafficAnomalies" api:"required"`
	JSON             trafficAnomalyGetResponseJSON             `json:"-"`
}

// trafficAnomalyGetResponseJSON contains the JSON metadata for the struct
// [TrafficAnomalyGetResponse]
type trafficAnomalyGetResponseJSON struct {
	TrafficAnomalies apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TrafficAnomalyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trafficAnomalyGetResponseJSON) RawJSON() string {
	return r.raw
}

type TrafficAnomalyGetResponseTrafficAnomaly struct {
	ASNDetails           TrafficAnomalyGetResponseTrafficAnomaliesASNDetails      `json:"asnDetails" api:"required,nullable"`
	EndDate              time.Time                                                `json:"endDate" api:"required,nullable" format:"date-time"`
	LocationDetails      TrafficAnomalyGetResponseTrafficAnomaliesLocationDetails `json:"locationDetails" api:"required,nullable"`
	OriginDetails        TrafficAnomalyGetResponseTrafficAnomaliesOriginDetails   `json:"originDetails" api:"required,nullable"`
	StartDate            string                                                   `json:"startDate" api:"required"`
	Status               string                                                   `json:"status" api:"required"`
	Type                 string                                                   `json:"type" api:"required"`
	UUID                 string                                                   `json:"uuid" api:"required"`
	VisibleInDataSources []string                                                 `json:"visibleInDataSources" api:"required,nullable"`
	JSON                 trafficAnomalyGetResponseTrafficAnomalyJSON              `json:"-"`
}

// trafficAnomalyGetResponseTrafficAnomalyJSON contains the JSON metadata for the
// struct [TrafficAnomalyGetResponseTrafficAnomaly]
type trafficAnomalyGetResponseTrafficAnomalyJSON struct {
	ASNDetails           apijson.Field
	EndDate              apijson.Field
	LocationDetails      apijson.Field
	OriginDetails        apijson.Field
	StartDate            apijson.Field
	Status               apijson.Field
	Type                 apijson.Field
	UUID                 apijson.Field
	VisibleInDataSources apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TrafficAnomalyGetResponseTrafficAnomaly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trafficAnomalyGetResponseTrafficAnomalyJSON) RawJSON() string {
	return r.raw
}

type TrafficAnomalyGetResponseTrafficAnomaliesASNDetails struct {
	ASN      string                                                      `json:"asn" api:"required"`
	Location TrafficAnomalyGetResponseTrafficAnomaliesASNDetailsLocation `json:"location" api:"required,nullable"`
	Name     string                                                      `json:"name" api:"required,nullable"`
	JSON     trafficAnomalyGetResponseTrafficAnomaliesASNDetailsJSON     `json:"-"`
}

// trafficAnomalyGetResponseTrafficAnomaliesASNDetailsJSON contains the JSON
// metadata for the struct [TrafficAnomalyGetResponseTrafficAnomaliesASNDetails]
type trafficAnomalyGetResponseTrafficAnomaliesASNDetailsJSON struct {
	ASN         apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TrafficAnomalyGetResponseTrafficAnomaliesASNDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trafficAnomalyGetResponseTrafficAnomaliesASNDetailsJSON) RawJSON() string {
	return r.raw
}

type TrafficAnomalyGetResponseTrafficAnomaliesASNDetailsLocation struct {
	Code string                                                          `json:"code" api:"required"`
	Name string                                                          `json:"name" api:"required"`
	JSON trafficAnomalyGetResponseTrafficAnomaliesASNDetailsLocationJSON `json:"-"`
}

// trafficAnomalyGetResponseTrafficAnomaliesASNDetailsLocationJSON contains the
// JSON metadata for the struct
// [TrafficAnomalyGetResponseTrafficAnomaliesASNDetailsLocation]
type trafficAnomalyGetResponseTrafficAnomaliesASNDetailsLocationJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TrafficAnomalyGetResponseTrafficAnomaliesASNDetailsLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trafficAnomalyGetResponseTrafficAnomaliesASNDetailsLocationJSON) RawJSON() string {
	return r.raw
}

type TrafficAnomalyGetResponseTrafficAnomaliesLocationDetails struct {
	Code string                                                       `json:"code" api:"required"`
	Name string                                                       `json:"name" api:"required"`
	JSON trafficAnomalyGetResponseTrafficAnomaliesLocationDetailsJSON `json:"-"`
}

// trafficAnomalyGetResponseTrafficAnomaliesLocationDetailsJSON contains the JSON
// metadata for the struct
// [TrafficAnomalyGetResponseTrafficAnomaliesLocationDetails]
type trafficAnomalyGetResponseTrafficAnomaliesLocationDetailsJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TrafficAnomalyGetResponseTrafficAnomaliesLocationDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trafficAnomalyGetResponseTrafficAnomaliesLocationDetailsJSON) RawJSON() string {
	return r.raw
}

type TrafficAnomalyGetResponseTrafficAnomaliesOriginDetails struct {
	Name   string                                                     `json:"name" api:"required,nullable"`
	Origin string                                                     `json:"origin" api:"required"`
	JSON   trafficAnomalyGetResponseTrafficAnomaliesOriginDetailsJSON `json:"-"`
}

// trafficAnomalyGetResponseTrafficAnomaliesOriginDetailsJSON contains the JSON
// metadata for the struct [TrafficAnomalyGetResponseTrafficAnomaliesOriginDetails]
type trafficAnomalyGetResponseTrafficAnomaliesOriginDetailsJSON struct {
	Name        apijson.Field
	Origin      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TrafficAnomalyGetResponseTrafficAnomaliesOriginDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trafficAnomalyGetResponseTrafficAnomaliesOriginDetailsJSON) RawJSON() string {
	return r.raw
}

type TrafficAnomalyGetParams struct {
	// Filters results by Autonomous System. Specify a single Autonomous System Number
	// (ASN) as integer.
	ASN param.Field[int64] `query:"asn"`
	// Filters results by data source.
	DataSource param.Field[TrafficAnomalyGetParamsDataSource] `query:"dataSource"`
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
	Format param.Field[TrafficAnomalyGetParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify an alpha-2 location code.
	Location param.Field[string] `query:"location"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64] `query:"offset"`
	// Filters results by origin.
	Origin param.Field[string]                        `query:"origin"`
	Status param.Field[TrafficAnomalyGetParamsStatus] `query:"status"`
	// Filters results by entity type (LOCATION, AS, or ORIGIN).
	Type param.Field[[]TrafficAnomalyGetParamsType] `query:"type"`
}

// URLQuery serializes [TrafficAnomalyGetParams]'s query parameters as
// `url.Values`.
func (r TrafficAnomalyGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filters results by data source.
type TrafficAnomalyGetParamsDataSource string

const (
	TrafficAnomalyGetParamsDataSourceAll                TrafficAnomalyGetParamsDataSource = "ALL"
	TrafficAnomalyGetParamsDataSourceAIBots             TrafficAnomalyGetParamsDataSource = "AI_BOTS"
	TrafficAnomalyGetParamsDataSourceAIGateway          TrafficAnomalyGetParamsDataSource = "AI_GATEWAY"
	TrafficAnomalyGetParamsDataSourceBGP                TrafficAnomalyGetParamsDataSource = "BGP"
	TrafficAnomalyGetParamsDataSourceBots               TrafficAnomalyGetParamsDataSource = "BOTS"
	TrafficAnomalyGetParamsDataSourceConnectionAnomaly  TrafficAnomalyGetParamsDataSource = "CONNECTION_ANOMALY"
	TrafficAnomalyGetParamsDataSourceCT                 TrafficAnomalyGetParamsDataSource = "CT"
	TrafficAnomalyGetParamsDataSourceDNS                TrafficAnomalyGetParamsDataSource = "DNS"
	TrafficAnomalyGetParamsDataSourceDNSMagnitude       TrafficAnomalyGetParamsDataSource = "DNS_MAGNITUDE"
	TrafficAnomalyGetParamsDataSourceDNSAS112           TrafficAnomalyGetParamsDataSource = "DNS_AS112"
	TrafficAnomalyGetParamsDataSourceDos                TrafficAnomalyGetParamsDataSource = "DOS"
	TrafficAnomalyGetParamsDataSourceEmailRouting       TrafficAnomalyGetParamsDataSource = "EMAIL_ROUTING"
	TrafficAnomalyGetParamsDataSourceEmailSecurity      TrafficAnomalyGetParamsDataSource = "EMAIL_SECURITY"
	TrafficAnomalyGetParamsDataSourceFw                 TrafficAnomalyGetParamsDataSource = "FW"
	TrafficAnomalyGetParamsDataSourceFwPg               TrafficAnomalyGetParamsDataSource = "FW_PG"
	TrafficAnomalyGetParamsDataSourceHTTP               TrafficAnomalyGetParamsDataSource = "HTTP"
	TrafficAnomalyGetParamsDataSourceHTTPControl        TrafficAnomalyGetParamsDataSource = "HTTP_CONTROL"
	TrafficAnomalyGetParamsDataSourceHTTPCrawlerReferer TrafficAnomalyGetParamsDataSource = "HTTP_CRAWLER_REFERER"
	TrafficAnomalyGetParamsDataSourceHTTPOrigins        TrafficAnomalyGetParamsDataSource = "HTTP_ORIGINS"
	TrafficAnomalyGetParamsDataSourceIQI                TrafficAnomalyGetParamsDataSource = "IQI"
	TrafficAnomalyGetParamsDataSourceLeakedCredentials  TrafficAnomalyGetParamsDataSource = "LEAKED_CREDENTIALS"
	TrafficAnomalyGetParamsDataSourceNet                TrafficAnomalyGetParamsDataSource = "NET"
	TrafficAnomalyGetParamsDataSourceRobotsTXT          TrafficAnomalyGetParamsDataSource = "ROBOTS_TXT"
	TrafficAnomalyGetParamsDataSourceSpeed              TrafficAnomalyGetParamsDataSource = "SPEED"
	TrafficAnomalyGetParamsDataSourceWorkersAI          TrafficAnomalyGetParamsDataSource = "WORKERS_AI"
)

func (r TrafficAnomalyGetParamsDataSource) IsKnown() bool {
	switch r {
	case TrafficAnomalyGetParamsDataSourceAll, TrafficAnomalyGetParamsDataSourceAIBots, TrafficAnomalyGetParamsDataSourceAIGateway, TrafficAnomalyGetParamsDataSourceBGP, TrafficAnomalyGetParamsDataSourceBots, TrafficAnomalyGetParamsDataSourceConnectionAnomaly, TrafficAnomalyGetParamsDataSourceCT, TrafficAnomalyGetParamsDataSourceDNS, TrafficAnomalyGetParamsDataSourceDNSMagnitude, TrafficAnomalyGetParamsDataSourceDNSAS112, TrafficAnomalyGetParamsDataSourceDos, TrafficAnomalyGetParamsDataSourceEmailRouting, TrafficAnomalyGetParamsDataSourceEmailSecurity, TrafficAnomalyGetParamsDataSourceFw, TrafficAnomalyGetParamsDataSourceFwPg, TrafficAnomalyGetParamsDataSourceHTTP, TrafficAnomalyGetParamsDataSourceHTTPControl, TrafficAnomalyGetParamsDataSourceHTTPCrawlerReferer, TrafficAnomalyGetParamsDataSourceHTTPOrigins, TrafficAnomalyGetParamsDataSourceIQI, TrafficAnomalyGetParamsDataSourceLeakedCredentials, TrafficAnomalyGetParamsDataSourceNet, TrafficAnomalyGetParamsDataSourceRobotsTXT, TrafficAnomalyGetParamsDataSourceSpeed, TrafficAnomalyGetParamsDataSourceWorkersAI:
		return true
	}
	return false
}

// Format in which results will be returned.
type TrafficAnomalyGetParamsFormat string

const (
	TrafficAnomalyGetParamsFormatJson TrafficAnomalyGetParamsFormat = "JSON"
	TrafficAnomalyGetParamsFormatCsv  TrafficAnomalyGetParamsFormat = "CSV"
)

func (r TrafficAnomalyGetParamsFormat) IsKnown() bool {
	switch r {
	case TrafficAnomalyGetParamsFormatJson, TrafficAnomalyGetParamsFormatCsv:
		return true
	}
	return false
}

type TrafficAnomalyGetParamsStatus string

const (
	TrafficAnomalyGetParamsStatusVerified   TrafficAnomalyGetParamsStatus = "VERIFIED"
	TrafficAnomalyGetParamsStatusUnverified TrafficAnomalyGetParamsStatus = "UNVERIFIED"
)

func (r TrafficAnomalyGetParamsStatus) IsKnown() bool {
	switch r {
	case TrafficAnomalyGetParamsStatusVerified, TrafficAnomalyGetParamsStatusUnverified:
		return true
	}
	return false
}

type TrafficAnomalyGetParamsType string

const (
	TrafficAnomalyGetParamsTypeLocation TrafficAnomalyGetParamsType = "LOCATION"
	TrafficAnomalyGetParamsTypeAs       TrafficAnomalyGetParamsType = "AS"
	TrafficAnomalyGetParamsTypeOrigin   TrafficAnomalyGetParamsType = "ORIGIN"
)

func (r TrafficAnomalyGetParamsType) IsKnown() bool {
	switch r {
	case TrafficAnomalyGetParamsTypeLocation, TrafficAnomalyGetParamsTypeAs, TrafficAnomalyGetParamsTypeOrigin:
		return true
	}
	return false
}

type TrafficAnomalyGetResponseEnvelope struct {
	Result  TrafficAnomalyGetResponse             `json:"result" api:"required"`
	Success bool                                  `json:"success" api:"required"`
	JSON    trafficAnomalyGetResponseEnvelopeJSON `json:"-"`
}

// trafficAnomalyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TrafficAnomalyGetResponseEnvelope]
type trafficAnomalyGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TrafficAnomalyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trafficAnomalyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
