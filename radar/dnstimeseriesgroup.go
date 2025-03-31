// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// DNSTimeseriesGroupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSTimeseriesGroupService] method instead.
type DNSTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewDNSTimeseriesGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDNSTimeseriesGroupService(opts ...option.RequestOption) (r *DNSTimeseriesGroupService) {
	r = &DNSTimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of DNS queries by cache status over time.
func (r *DNSTimeseriesGroupService) CacheHit(ctx context.Context, query DNSTimeseriesGroupCacheHitParams, opts ...option.RequestOption) (res *DNSTimeseriesGroupCacheHitResponse, err error) {
	var env DNSTimeseriesGroupCacheHitResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/cache_hit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS responses by DNSSEC (DNS Security Extensions)
// support over time.
func (r *DNSTimeseriesGroupService) DNSSEC(ctx context.Context, query DNSTimeseriesGroupDNSSECParams, opts ...option.RequestOption) (res *DNSTimeseriesGroupDNSSECResponse, err error) {
	var env DNSTimeseriesGroupDNSSECResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/dnssec"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS queries by DNSSEC (DNS Security Extensions)
// client awareness over time.
func (r *DNSTimeseriesGroupService) DNSSECAware(ctx context.Context, query DNSTimeseriesGroupDNSSECAwareParams, opts ...option.RequestOption) (res *DNSTimeseriesGroupDNSSECAwareResponse, err error) {
	var env DNSTimeseriesGroupDNSSECAwareResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/dnssec_aware"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNSSEC-validated answers by end-to-end security
// status over time.
func (r *DNSTimeseriesGroupService) DNSSECE2E(ctx context.Context, query DNSTimeseriesGroupDNSSECE2EParams, opts ...option.RequestOption) (res *DNSTimeseriesGroupDnssece2EResponse, err error) {
	var env DNSTimeseriesGroupDnssece2EResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/dnssec_e2e"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS queries by IP version over time.
func (r *DNSTimeseriesGroupService) IPVersion(ctx context.Context, query DNSTimeseriesGroupIPVersionParams, opts ...option.RequestOption) (res *DNSTimeseriesGroupIPVersionResponse, err error) {
	var env DNSTimeseriesGroupIPVersionResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS queries by matching answers over time.
func (r *DNSTimeseriesGroupService) MatchingAnswer(ctx context.Context, query DNSTimeseriesGroupMatchingAnswerParams, opts ...option.RequestOption) (res *DNSTimeseriesGroupMatchingAnswerResponse, err error) {
	var env DNSTimeseriesGroupMatchingAnswerResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/matching_answer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS queries by DNS transport protocol over time.
func (r *DNSTimeseriesGroupService) Protocol(ctx context.Context, query DNSTimeseriesGroupProtocolParams, opts ...option.RequestOption) (res *DNSTimeseriesGroupProtocolResponse, err error) {
	var env DNSTimeseriesGroupProtocolResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS queries by type over time.
func (r *DNSTimeseriesGroupService) QueryType(ctx context.Context, query DNSTimeseriesGroupQueryTypeParams, opts ...option.RequestOption) (res *DNSTimeseriesGroupQueryTypeResponse, err error) {
	var env DNSTimeseriesGroupQueryTypeResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/query_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS queries by response code over time.
func (r *DNSTimeseriesGroupService) ResponseCode(ctx context.Context, query DNSTimeseriesGroupResponseCodeParams, opts ...option.RequestOption) (res *DNSTimeseriesGroupResponseCodeResponse, err error) {
	var env DNSTimeseriesGroupResponseCodeResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/response_code"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of DNS queries by minimum answer TTL over time.
func (r *DNSTimeseriesGroupService) ResponseTTL(ctx context.Context, query DNSTimeseriesGroupResponseTTLParams, opts ...option.RequestOption) (res *DNSTimeseriesGroupResponseTTLResponse, err error) {
	var env DNSTimeseriesGroupResponseTTLResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/response_ttl"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSTimeseriesGroupCacheHitResponse struct {
	Meta   interface{}                              `json:"meta,required"`
	Serie0 DNSTimeseriesGroupCacheHitResponseSerie0 `json:"serie_0,required"`
	JSON   dnsTimeseriesGroupCacheHitResponseJSON   `json:"-"`
}

// dnsTimeseriesGroupCacheHitResponseJSON contains the JSON metadata for the struct
// [DNSTimeseriesGroupCacheHitResponse]
type dnsTimeseriesGroupCacheHitResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupCacheHitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupCacheHitResponseJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupCacheHitResponseSerie0 struct {
	Negative []string                                     `json:"NEGATIVE,required"`
	Positive []string                                     `json:"POSITIVE,required"`
	JSON     dnsTimeseriesGroupCacheHitResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesGroupCacheHitResponseSerie0JSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupCacheHitResponseSerie0]
type dnsTimeseriesGroupCacheHitResponseSerie0JSON struct {
	Negative    apijson.Field
	Positive    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupCacheHitResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupCacheHitResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupDNSSECResponse struct {
	Meta   interface{}                            `json:"meta,required"`
	Serie0 DNSTimeseriesGroupDNSSECResponseSerie0 `json:"serie_0,required"`
	JSON   dnsTimeseriesGroupDNSSECResponseJSON   `json:"-"`
}

// dnsTimeseriesGroupDNSSECResponseJSON contains the JSON metadata for the struct
// [DNSTimeseriesGroupDNSSECResponse]
type dnsTimeseriesGroupDNSSECResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupDNSSECResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupDNSSECResponseJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupDNSSECResponseSerie0 struct {
	Insecure []string                                   `json:"INSECURE,required"`
	Invalid  []string                                   `json:"INVALID,required"`
	Other    []string                                   `json:"OTHER,required"`
	Secure   []string                                   `json:"SECURE,required"`
	JSON     dnsTimeseriesGroupDNSSECResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesGroupDNSSECResponseSerie0JSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupDNSSECResponseSerie0]
type dnsTimeseriesGroupDNSSECResponseSerie0JSON struct {
	Insecure    apijson.Field
	Invalid     apijson.Field
	Other       apijson.Field
	Secure      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupDNSSECResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupDNSSECResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupDNSSECAwareResponse struct {
	Meta   interface{}                                 `json:"meta,required"`
	Serie0 DNSTimeseriesGroupDNSSECAwareResponseSerie0 `json:"serie_0,required"`
	JSON   dnsTimeseriesGroupDNSSECAwareResponseJSON   `json:"-"`
}

// dnsTimeseriesGroupDNSSECAwareResponseJSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupDNSSECAwareResponse]
type dnsTimeseriesGroupDNSSECAwareResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupDNSSECAwareResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupDNSSECAwareResponseJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupDNSSECAwareResponseSerie0 struct {
	NotSupported []string                                        `json:"NOT_SUPPORTED,required"`
	Supported    []string                                        `json:"SUPPORTED,required"`
	JSON         dnsTimeseriesGroupDNSSECAwareResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesGroupDNSSECAwareResponseSerie0JSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupDNSSECAwareResponseSerie0]
type dnsTimeseriesGroupDNSSECAwareResponseSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSTimeseriesGroupDNSSECAwareResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupDNSSECAwareResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupDnssece2EResponse struct {
	Meta   interface{}                               `json:"meta,required"`
	Serie0 DNSTimeseriesGroupDnssece2EResponseSerie0 `json:"serie_0,required"`
	JSON   dnsTimeseriesGroupDnssece2EResponseJSON   `json:"-"`
}

// dnsTimeseriesGroupDnssece2EResponseJSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupDnssece2EResponse]
type dnsTimeseriesGroupDnssece2EResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupDnssece2EResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupDnssece2EResponseJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupDnssece2EResponseSerie0 struct {
	Negative []string                                      `json:"NEGATIVE,required"`
	Positive []string                                      `json:"POSITIVE,required"`
	JSON     dnsTimeseriesGroupDnssece2EResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesGroupDnssece2EResponseSerie0JSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupDnssece2EResponseSerie0]
type dnsTimeseriesGroupDnssece2EResponseSerie0JSON struct {
	Negative    apijson.Field
	Positive    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupDnssece2EResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupDnssece2EResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupIPVersionResponse struct {
	Meta   interface{}                               `json:"meta,required"`
	Serie0 DNSTimeseriesGroupIPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   dnsTimeseriesGroupIPVersionResponseJSON   `json:"-"`
}

// dnsTimeseriesGroupIPVersionResponseJSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupIPVersionResponse]
type dnsTimeseriesGroupIPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupIPVersionResponseSerie0 struct {
	IPv4 []string                                      `json:"IPv4,required"`
	IPv6 []string                                      `json:"IPv6,required"`
	JSON dnsTimeseriesGroupIPVersionResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesGroupIPVersionResponseSerie0JSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupIPVersionResponseSerie0]
type dnsTimeseriesGroupIPVersionResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupIPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupIPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupMatchingAnswerResponse struct {
	Meta   interface{}                                    `json:"meta,required"`
	Serie0 DNSTimeseriesGroupMatchingAnswerResponseSerie0 `json:"serie_0,required"`
	JSON   dnsTimeseriesGroupMatchingAnswerResponseJSON   `json:"-"`
}

// dnsTimeseriesGroupMatchingAnswerResponseJSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupMatchingAnswerResponse]
type dnsTimeseriesGroupMatchingAnswerResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupMatchingAnswerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupMatchingAnswerResponseJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupMatchingAnswerResponseSerie0 struct {
	Negative []string                                           `json:"NEGATIVE,required"`
	Positive []string                                           `json:"POSITIVE,required"`
	JSON     dnsTimeseriesGroupMatchingAnswerResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesGroupMatchingAnswerResponseSerie0JSON contains the JSON metadata
// for the struct [DNSTimeseriesGroupMatchingAnswerResponseSerie0]
type dnsTimeseriesGroupMatchingAnswerResponseSerie0JSON struct {
	Negative    apijson.Field
	Positive    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupMatchingAnswerResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupMatchingAnswerResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupProtocolResponse struct {
	Meta   interface{}                              `json:"meta,required"`
	Serie0 DNSTimeseriesGroupProtocolResponseSerie0 `json:"serie_0,required"`
	JSON   dnsTimeseriesGroupProtocolResponseJSON   `json:"-"`
}

// dnsTimeseriesGroupProtocolResponseJSON contains the JSON metadata for the struct
// [DNSTimeseriesGroupProtocolResponse]
type dnsTimeseriesGroupProtocolResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupProtocolResponseJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupProtocolResponseSerie0 struct {
	HTTPS []string                                     `json:"HTTPS,required"`
	TCP   []string                                     `json:"TCP,required"`
	TLS   []string                                     `json:"TLS,required"`
	Udp   []string                                     `json:"UDP,required"`
	JSON  dnsTimeseriesGroupProtocolResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesGroupProtocolResponseSerie0JSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupProtocolResponseSerie0]
type dnsTimeseriesGroupProtocolResponseSerie0JSON struct {
	HTTPS       apijson.Field
	TCP         apijson.Field
	TLS         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupProtocolResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupProtocolResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupQueryTypeResponse struct {
	Meta   interface{}                             `json:"meta,required"`
	Serie0 interface{}                             `json:"serie_0,required"`
	JSON   dnsTimeseriesGroupQueryTypeResponseJSON `json:"-"`
}

// dnsTimeseriesGroupQueryTypeResponseJSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupQueryTypeResponse]
type dnsTimeseriesGroupQueryTypeResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupQueryTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupQueryTypeResponseJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupResponseCodeResponse struct {
	Meta   interface{}                                `json:"meta,required"`
	Serie0 interface{}                                `json:"serie_0,required"`
	JSON   dnsTimeseriesGroupResponseCodeResponseJSON `json:"-"`
}

// dnsTimeseriesGroupResponseCodeResponseJSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupResponseCodeResponse]
type dnsTimeseriesGroupResponseCodeResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupResponseCodeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupResponseCodeResponseJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupResponseTTLResponse struct {
	Meta   interface{}                                 `json:"meta,required"`
	Serie0 DNSTimeseriesGroupResponseTTLResponseSerie0 `json:"serie_0,required"`
	JSON   dnsTimeseriesGroupResponseTTLResponseJSON   `json:"-"`
}

// dnsTimeseriesGroupResponseTTLResponseJSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupResponseTTLResponse]
type dnsTimeseriesGroupResponseTTLResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupResponseTTLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupResponseTTLResponseJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupResponseTTLResponseSerie0 struct {
	Gt15mLte1h []string                                        `json:"gt_15m_lte_1h,required"`
	Gt1dLte1w  []string                                        `json:"gt_1d_lte_1w,required"`
	Gt1hLte1d  []string                                        `json:"gt_1h_lte_1d,required"`
	Gt1mLte5m  []string                                        `json:"gt_1m_lte_5m,required"`
	Gt1w       []string                                        `json:"gt_1w,required"`
	Gt5mLte15m []string                                        `json:"gt_5m_lte_15m,required"`
	Lte1m      []string                                        `json:"lte_1m,required"`
	JSON       dnsTimeseriesGroupResponseTTLResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesGroupResponseTTLResponseSerie0JSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupResponseTTLResponseSerie0]
type dnsTimeseriesGroupResponseTTLResponseSerie0JSON struct {
	Gt15mLte1h  apijson.Field
	Gt1dLte1w   apijson.Field
	Gt1hLte1d   apijson.Field
	Gt1mLte5m   apijson.Field
	Gt1w        apijson.Field
	Gt5mLte15m  apijson.Field
	Lte1m       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupResponseTTLResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupResponseTTLResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupCacheHitParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[DNSTimeseriesGroupCacheHitParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[DNSTimeseriesGroupCacheHitParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSTimeseriesGroupCacheHitParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSTimeseriesGroupCacheHitParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSTimeseriesGroupCacheHitParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupCacheHitParams]'s query parameters as
// `url.Values`.
func (r DNSTimeseriesGroupCacheHitParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupCacheHitParamsAggInterval string

const (
	DNSTimeseriesGroupCacheHitParamsAggInterval15m DNSTimeseriesGroupCacheHitParamsAggInterval = "15m"
	DNSTimeseriesGroupCacheHitParamsAggInterval1h  DNSTimeseriesGroupCacheHitParamsAggInterval = "1h"
	DNSTimeseriesGroupCacheHitParamsAggInterval1d  DNSTimeseriesGroupCacheHitParamsAggInterval = "1d"
	DNSTimeseriesGroupCacheHitParamsAggInterval1w  DNSTimeseriesGroupCacheHitParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupCacheHitParamsAggInterval) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupCacheHitParamsAggInterval15m, DNSTimeseriesGroupCacheHitParamsAggInterval1h, DNSTimeseriesGroupCacheHitParamsAggInterval1d, DNSTimeseriesGroupCacheHitParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type DNSTimeseriesGroupCacheHitParamsFormat string

const (
	DNSTimeseriesGroupCacheHitParamsFormatJson DNSTimeseriesGroupCacheHitParamsFormat = "JSON"
	DNSTimeseriesGroupCacheHitParamsFormatCsv  DNSTimeseriesGroupCacheHitParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupCacheHitParamsFormat) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupCacheHitParamsFormatJson, DNSTimeseriesGroupCacheHitParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSTimeseriesGroupCacheHitParamsProtocol string

const (
	DNSTimeseriesGroupCacheHitParamsProtocolUdp   DNSTimeseriesGroupCacheHitParamsProtocol = "UDP"
	DNSTimeseriesGroupCacheHitParamsProtocolTCP   DNSTimeseriesGroupCacheHitParamsProtocol = "TCP"
	DNSTimeseriesGroupCacheHitParamsProtocolHTTPS DNSTimeseriesGroupCacheHitParamsProtocol = "HTTPS"
	DNSTimeseriesGroupCacheHitParamsProtocolTLS   DNSTimeseriesGroupCacheHitParamsProtocol = "TLS"
)

func (r DNSTimeseriesGroupCacheHitParamsProtocol) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupCacheHitParamsProtocolUdp, DNSTimeseriesGroupCacheHitParamsProtocolTCP, DNSTimeseriesGroupCacheHitParamsProtocolHTTPS, DNSTimeseriesGroupCacheHitParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSTimeseriesGroupCacheHitParamsQueryType string

const (
	DNSTimeseriesGroupCacheHitParamsQueryTypeA          DNSTimeseriesGroupCacheHitParamsQueryType = "A"
	DNSTimeseriesGroupCacheHitParamsQueryTypeAAAA       DNSTimeseriesGroupCacheHitParamsQueryType = "AAAA"
	DNSTimeseriesGroupCacheHitParamsQueryTypeA6         DNSTimeseriesGroupCacheHitParamsQueryType = "A6"
	DNSTimeseriesGroupCacheHitParamsQueryTypeAfsdb      DNSTimeseriesGroupCacheHitParamsQueryType = "AFSDB"
	DNSTimeseriesGroupCacheHitParamsQueryTypeAny        DNSTimeseriesGroupCacheHitParamsQueryType = "ANY"
	DNSTimeseriesGroupCacheHitParamsQueryTypeApl        DNSTimeseriesGroupCacheHitParamsQueryType = "APL"
	DNSTimeseriesGroupCacheHitParamsQueryTypeAtma       DNSTimeseriesGroupCacheHitParamsQueryType = "ATMA"
	DNSTimeseriesGroupCacheHitParamsQueryTypeAXFR       DNSTimeseriesGroupCacheHitParamsQueryType = "AXFR"
	DNSTimeseriesGroupCacheHitParamsQueryTypeCAA        DNSTimeseriesGroupCacheHitParamsQueryType = "CAA"
	DNSTimeseriesGroupCacheHitParamsQueryTypeCdnskey    DNSTimeseriesGroupCacheHitParamsQueryType = "CDNSKEY"
	DNSTimeseriesGroupCacheHitParamsQueryTypeCds        DNSTimeseriesGroupCacheHitParamsQueryType = "CDS"
	DNSTimeseriesGroupCacheHitParamsQueryTypeCERT       DNSTimeseriesGroupCacheHitParamsQueryType = "CERT"
	DNSTimeseriesGroupCacheHitParamsQueryTypeCNAME      DNSTimeseriesGroupCacheHitParamsQueryType = "CNAME"
	DNSTimeseriesGroupCacheHitParamsQueryTypeCsync      DNSTimeseriesGroupCacheHitParamsQueryType = "CSYNC"
	DNSTimeseriesGroupCacheHitParamsQueryTypeDhcid      DNSTimeseriesGroupCacheHitParamsQueryType = "DHCID"
	DNSTimeseriesGroupCacheHitParamsQueryTypeDlv        DNSTimeseriesGroupCacheHitParamsQueryType = "DLV"
	DNSTimeseriesGroupCacheHitParamsQueryTypeDname      DNSTimeseriesGroupCacheHitParamsQueryType = "DNAME"
	DNSTimeseriesGroupCacheHitParamsQueryTypeDNSKEY     DNSTimeseriesGroupCacheHitParamsQueryType = "DNSKEY"
	DNSTimeseriesGroupCacheHitParamsQueryTypeDoa        DNSTimeseriesGroupCacheHitParamsQueryType = "DOA"
	DNSTimeseriesGroupCacheHitParamsQueryTypeDS         DNSTimeseriesGroupCacheHitParamsQueryType = "DS"
	DNSTimeseriesGroupCacheHitParamsQueryTypeEid        DNSTimeseriesGroupCacheHitParamsQueryType = "EID"
	DNSTimeseriesGroupCacheHitParamsQueryTypeEui48      DNSTimeseriesGroupCacheHitParamsQueryType = "EUI48"
	DNSTimeseriesGroupCacheHitParamsQueryTypeEui64      DNSTimeseriesGroupCacheHitParamsQueryType = "EUI64"
	DNSTimeseriesGroupCacheHitParamsQueryTypeGpos       DNSTimeseriesGroupCacheHitParamsQueryType = "GPOS"
	DNSTimeseriesGroupCacheHitParamsQueryTypeGid        DNSTimeseriesGroupCacheHitParamsQueryType = "GID"
	DNSTimeseriesGroupCacheHitParamsQueryTypeHinfo      DNSTimeseriesGroupCacheHitParamsQueryType = "HINFO"
	DNSTimeseriesGroupCacheHitParamsQueryTypeHip        DNSTimeseriesGroupCacheHitParamsQueryType = "HIP"
	DNSTimeseriesGroupCacheHitParamsQueryTypeHTTPS      DNSTimeseriesGroupCacheHitParamsQueryType = "HTTPS"
	DNSTimeseriesGroupCacheHitParamsQueryTypeIpseckey   DNSTimeseriesGroupCacheHitParamsQueryType = "IPSECKEY"
	DNSTimeseriesGroupCacheHitParamsQueryTypeIsdn       DNSTimeseriesGroupCacheHitParamsQueryType = "ISDN"
	DNSTimeseriesGroupCacheHitParamsQueryTypeIxfr       DNSTimeseriesGroupCacheHitParamsQueryType = "IXFR"
	DNSTimeseriesGroupCacheHitParamsQueryTypeKey        DNSTimeseriesGroupCacheHitParamsQueryType = "KEY"
	DNSTimeseriesGroupCacheHitParamsQueryTypeKx         DNSTimeseriesGroupCacheHitParamsQueryType = "KX"
	DNSTimeseriesGroupCacheHitParamsQueryTypeL32        DNSTimeseriesGroupCacheHitParamsQueryType = "L32"
	DNSTimeseriesGroupCacheHitParamsQueryTypeL64        DNSTimeseriesGroupCacheHitParamsQueryType = "L64"
	DNSTimeseriesGroupCacheHitParamsQueryTypeLOC        DNSTimeseriesGroupCacheHitParamsQueryType = "LOC"
	DNSTimeseriesGroupCacheHitParamsQueryTypeLp         DNSTimeseriesGroupCacheHitParamsQueryType = "LP"
	DNSTimeseriesGroupCacheHitParamsQueryTypeMaila      DNSTimeseriesGroupCacheHitParamsQueryType = "MAILA"
	DNSTimeseriesGroupCacheHitParamsQueryTypeMailb      DNSTimeseriesGroupCacheHitParamsQueryType = "MAILB"
	DNSTimeseriesGroupCacheHitParamsQueryTypeMB         DNSTimeseriesGroupCacheHitParamsQueryType = "MB"
	DNSTimeseriesGroupCacheHitParamsQueryTypeMd         DNSTimeseriesGroupCacheHitParamsQueryType = "MD"
	DNSTimeseriesGroupCacheHitParamsQueryTypeMf         DNSTimeseriesGroupCacheHitParamsQueryType = "MF"
	DNSTimeseriesGroupCacheHitParamsQueryTypeMg         DNSTimeseriesGroupCacheHitParamsQueryType = "MG"
	DNSTimeseriesGroupCacheHitParamsQueryTypeMinfo      DNSTimeseriesGroupCacheHitParamsQueryType = "MINFO"
	DNSTimeseriesGroupCacheHitParamsQueryTypeMr         DNSTimeseriesGroupCacheHitParamsQueryType = "MR"
	DNSTimeseriesGroupCacheHitParamsQueryTypeMX         DNSTimeseriesGroupCacheHitParamsQueryType = "MX"
	DNSTimeseriesGroupCacheHitParamsQueryTypeNAPTR      DNSTimeseriesGroupCacheHitParamsQueryType = "NAPTR"
	DNSTimeseriesGroupCacheHitParamsQueryTypeNb         DNSTimeseriesGroupCacheHitParamsQueryType = "NB"
	DNSTimeseriesGroupCacheHitParamsQueryTypeNbstat     DNSTimeseriesGroupCacheHitParamsQueryType = "NBSTAT"
	DNSTimeseriesGroupCacheHitParamsQueryTypeNid        DNSTimeseriesGroupCacheHitParamsQueryType = "NID"
	DNSTimeseriesGroupCacheHitParamsQueryTypeNimloc     DNSTimeseriesGroupCacheHitParamsQueryType = "NIMLOC"
	DNSTimeseriesGroupCacheHitParamsQueryTypeNinfo      DNSTimeseriesGroupCacheHitParamsQueryType = "NINFO"
	DNSTimeseriesGroupCacheHitParamsQueryTypeNS         DNSTimeseriesGroupCacheHitParamsQueryType = "NS"
	DNSTimeseriesGroupCacheHitParamsQueryTypeNsap       DNSTimeseriesGroupCacheHitParamsQueryType = "NSAP"
	DNSTimeseriesGroupCacheHitParamsQueryTypeNsec       DNSTimeseriesGroupCacheHitParamsQueryType = "NSEC"
	DNSTimeseriesGroupCacheHitParamsQueryTypeNsec3      DNSTimeseriesGroupCacheHitParamsQueryType = "NSEC3"
	DNSTimeseriesGroupCacheHitParamsQueryTypeNsec3Param DNSTimeseriesGroupCacheHitParamsQueryType = "NSEC3PARAM"
	DNSTimeseriesGroupCacheHitParamsQueryTypeNull       DNSTimeseriesGroupCacheHitParamsQueryType = "NULL"
	DNSTimeseriesGroupCacheHitParamsQueryTypeNxt        DNSTimeseriesGroupCacheHitParamsQueryType = "NXT"
	DNSTimeseriesGroupCacheHitParamsQueryTypeOpenpgpkey DNSTimeseriesGroupCacheHitParamsQueryType = "OPENPGPKEY"
	DNSTimeseriesGroupCacheHitParamsQueryTypeOpt        DNSTimeseriesGroupCacheHitParamsQueryType = "OPT"
	DNSTimeseriesGroupCacheHitParamsQueryTypePTR        DNSTimeseriesGroupCacheHitParamsQueryType = "PTR"
	DNSTimeseriesGroupCacheHitParamsQueryTypePx         DNSTimeseriesGroupCacheHitParamsQueryType = "PX"
	DNSTimeseriesGroupCacheHitParamsQueryTypeRkey       DNSTimeseriesGroupCacheHitParamsQueryType = "RKEY"
	DNSTimeseriesGroupCacheHitParamsQueryTypeRp         DNSTimeseriesGroupCacheHitParamsQueryType = "RP"
	DNSTimeseriesGroupCacheHitParamsQueryTypeRrsig      DNSTimeseriesGroupCacheHitParamsQueryType = "RRSIG"
	DNSTimeseriesGroupCacheHitParamsQueryTypeRt         DNSTimeseriesGroupCacheHitParamsQueryType = "RT"
	DNSTimeseriesGroupCacheHitParamsQueryTypeSig        DNSTimeseriesGroupCacheHitParamsQueryType = "SIG"
	DNSTimeseriesGroupCacheHitParamsQueryTypeSink       DNSTimeseriesGroupCacheHitParamsQueryType = "SINK"
	DNSTimeseriesGroupCacheHitParamsQueryTypeSMIMEA     DNSTimeseriesGroupCacheHitParamsQueryType = "SMIMEA"
	DNSTimeseriesGroupCacheHitParamsQueryTypeSOA        DNSTimeseriesGroupCacheHitParamsQueryType = "SOA"
	DNSTimeseriesGroupCacheHitParamsQueryTypeSPF        DNSTimeseriesGroupCacheHitParamsQueryType = "SPF"
	DNSTimeseriesGroupCacheHitParamsQueryTypeSRV        DNSTimeseriesGroupCacheHitParamsQueryType = "SRV"
	DNSTimeseriesGroupCacheHitParamsQueryTypeSSHFP      DNSTimeseriesGroupCacheHitParamsQueryType = "SSHFP"
	DNSTimeseriesGroupCacheHitParamsQueryTypeSVCB       DNSTimeseriesGroupCacheHitParamsQueryType = "SVCB"
	DNSTimeseriesGroupCacheHitParamsQueryTypeTa         DNSTimeseriesGroupCacheHitParamsQueryType = "TA"
	DNSTimeseriesGroupCacheHitParamsQueryTypeTalink     DNSTimeseriesGroupCacheHitParamsQueryType = "TALINK"
	DNSTimeseriesGroupCacheHitParamsQueryTypeTkey       DNSTimeseriesGroupCacheHitParamsQueryType = "TKEY"
	DNSTimeseriesGroupCacheHitParamsQueryTypeTLSA       DNSTimeseriesGroupCacheHitParamsQueryType = "TLSA"
	DNSTimeseriesGroupCacheHitParamsQueryTypeTSIG       DNSTimeseriesGroupCacheHitParamsQueryType = "TSIG"
	DNSTimeseriesGroupCacheHitParamsQueryTypeTXT        DNSTimeseriesGroupCacheHitParamsQueryType = "TXT"
	DNSTimeseriesGroupCacheHitParamsQueryTypeUinfo      DNSTimeseriesGroupCacheHitParamsQueryType = "UINFO"
	DNSTimeseriesGroupCacheHitParamsQueryTypeUID        DNSTimeseriesGroupCacheHitParamsQueryType = "UID"
	DNSTimeseriesGroupCacheHitParamsQueryTypeUnspec     DNSTimeseriesGroupCacheHitParamsQueryType = "UNSPEC"
	DNSTimeseriesGroupCacheHitParamsQueryTypeURI        DNSTimeseriesGroupCacheHitParamsQueryType = "URI"
	DNSTimeseriesGroupCacheHitParamsQueryTypeWks        DNSTimeseriesGroupCacheHitParamsQueryType = "WKS"
	DNSTimeseriesGroupCacheHitParamsQueryTypeX25        DNSTimeseriesGroupCacheHitParamsQueryType = "X25"
	DNSTimeseriesGroupCacheHitParamsQueryTypeZonemd     DNSTimeseriesGroupCacheHitParamsQueryType = "ZONEMD"
)

func (r DNSTimeseriesGroupCacheHitParamsQueryType) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupCacheHitParamsQueryTypeA, DNSTimeseriesGroupCacheHitParamsQueryTypeAAAA, DNSTimeseriesGroupCacheHitParamsQueryTypeA6, DNSTimeseriesGroupCacheHitParamsQueryTypeAfsdb, DNSTimeseriesGroupCacheHitParamsQueryTypeAny, DNSTimeseriesGroupCacheHitParamsQueryTypeApl, DNSTimeseriesGroupCacheHitParamsQueryTypeAtma, DNSTimeseriesGroupCacheHitParamsQueryTypeAXFR, DNSTimeseriesGroupCacheHitParamsQueryTypeCAA, DNSTimeseriesGroupCacheHitParamsQueryTypeCdnskey, DNSTimeseriesGroupCacheHitParamsQueryTypeCds, DNSTimeseriesGroupCacheHitParamsQueryTypeCERT, DNSTimeseriesGroupCacheHitParamsQueryTypeCNAME, DNSTimeseriesGroupCacheHitParamsQueryTypeCsync, DNSTimeseriesGroupCacheHitParamsQueryTypeDhcid, DNSTimeseriesGroupCacheHitParamsQueryTypeDlv, DNSTimeseriesGroupCacheHitParamsQueryTypeDname, DNSTimeseriesGroupCacheHitParamsQueryTypeDNSKEY, DNSTimeseriesGroupCacheHitParamsQueryTypeDoa, DNSTimeseriesGroupCacheHitParamsQueryTypeDS, DNSTimeseriesGroupCacheHitParamsQueryTypeEid, DNSTimeseriesGroupCacheHitParamsQueryTypeEui48, DNSTimeseriesGroupCacheHitParamsQueryTypeEui64, DNSTimeseriesGroupCacheHitParamsQueryTypeGpos, DNSTimeseriesGroupCacheHitParamsQueryTypeGid, DNSTimeseriesGroupCacheHitParamsQueryTypeHinfo, DNSTimeseriesGroupCacheHitParamsQueryTypeHip, DNSTimeseriesGroupCacheHitParamsQueryTypeHTTPS, DNSTimeseriesGroupCacheHitParamsQueryTypeIpseckey, DNSTimeseriesGroupCacheHitParamsQueryTypeIsdn, DNSTimeseriesGroupCacheHitParamsQueryTypeIxfr, DNSTimeseriesGroupCacheHitParamsQueryTypeKey, DNSTimeseriesGroupCacheHitParamsQueryTypeKx, DNSTimeseriesGroupCacheHitParamsQueryTypeL32, DNSTimeseriesGroupCacheHitParamsQueryTypeL64, DNSTimeseriesGroupCacheHitParamsQueryTypeLOC, DNSTimeseriesGroupCacheHitParamsQueryTypeLp, DNSTimeseriesGroupCacheHitParamsQueryTypeMaila, DNSTimeseriesGroupCacheHitParamsQueryTypeMailb, DNSTimeseriesGroupCacheHitParamsQueryTypeMB, DNSTimeseriesGroupCacheHitParamsQueryTypeMd, DNSTimeseriesGroupCacheHitParamsQueryTypeMf, DNSTimeseriesGroupCacheHitParamsQueryTypeMg, DNSTimeseriesGroupCacheHitParamsQueryTypeMinfo, DNSTimeseriesGroupCacheHitParamsQueryTypeMr, DNSTimeseriesGroupCacheHitParamsQueryTypeMX, DNSTimeseriesGroupCacheHitParamsQueryTypeNAPTR, DNSTimeseriesGroupCacheHitParamsQueryTypeNb, DNSTimeseriesGroupCacheHitParamsQueryTypeNbstat, DNSTimeseriesGroupCacheHitParamsQueryTypeNid, DNSTimeseriesGroupCacheHitParamsQueryTypeNimloc, DNSTimeseriesGroupCacheHitParamsQueryTypeNinfo, DNSTimeseriesGroupCacheHitParamsQueryTypeNS, DNSTimeseriesGroupCacheHitParamsQueryTypeNsap, DNSTimeseriesGroupCacheHitParamsQueryTypeNsec, DNSTimeseriesGroupCacheHitParamsQueryTypeNsec3, DNSTimeseriesGroupCacheHitParamsQueryTypeNsec3Param, DNSTimeseriesGroupCacheHitParamsQueryTypeNull, DNSTimeseriesGroupCacheHitParamsQueryTypeNxt, DNSTimeseriesGroupCacheHitParamsQueryTypeOpenpgpkey, DNSTimeseriesGroupCacheHitParamsQueryTypeOpt, DNSTimeseriesGroupCacheHitParamsQueryTypePTR, DNSTimeseriesGroupCacheHitParamsQueryTypePx, DNSTimeseriesGroupCacheHitParamsQueryTypeRkey, DNSTimeseriesGroupCacheHitParamsQueryTypeRp, DNSTimeseriesGroupCacheHitParamsQueryTypeRrsig, DNSTimeseriesGroupCacheHitParamsQueryTypeRt, DNSTimeseriesGroupCacheHitParamsQueryTypeSig, DNSTimeseriesGroupCacheHitParamsQueryTypeSink, DNSTimeseriesGroupCacheHitParamsQueryTypeSMIMEA, DNSTimeseriesGroupCacheHitParamsQueryTypeSOA, DNSTimeseriesGroupCacheHitParamsQueryTypeSPF, DNSTimeseriesGroupCacheHitParamsQueryTypeSRV, DNSTimeseriesGroupCacheHitParamsQueryTypeSSHFP, DNSTimeseriesGroupCacheHitParamsQueryTypeSVCB, DNSTimeseriesGroupCacheHitParamsQueryTypeTa, DNSTimeseriesGroupCacheHitParamsQueryTypeTalink, DNSTimeseriesGroupCacheHitParamsQueryTypeTkey, DNSTimeseriesGroupCacheHitParamsQueryTypeTLSA, DNSTimeseriesGroupCacheHitParamsQueryTypeTSIG, DNSTimeseriesGroupCacheHitParamsQueryTypeTXT, DNSTimeseriesGroupCacheHitParamsQueryTypeUinfo, DNSTimeseriesGroupCacheHitParamsQueryTypeUID, DNSTimeseriesGroupCacheHitParamsQueryTypeUnspec, DNSTimeseriesGroupCacheHitParamsQueryTypeURI, DNSTimeseriesGroupCacheHitParamsQueryTypeWks, DNSTimeseriesGroupCacheHitParamsQueryTypeX25, DNSTimeseriesGroupCacheHitParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSTimeseriesGroupCacheHitParamsResponseCode string

const (
	DNSTimeseriesGroupCacheHitParamsResponseCodeNoerror   DNSTimeseriesGroupCacheHitParamsResponseCode = "NOERROR"
	DNSTimeseriesGroupCacheHitParamsResponseCodeFormerr   DNSTimeseriesGroupCacheHitParamsResponseCode = "FORMERR"
	DNSTimeseriesGroupCacheHitParamsResponseCodeServfail  DNSTimeseriesGroupCacheHitParamsResponseCode = "SERVFAIL"
	DNSTimeseriesGroupCacheHitParamsResponseCodeNxdomain  DNSTimeseriesGroupCacheHitParamsResponseCode = "NXDOMAIN"
	DNSTimeseriesGroupCacheHitParamsResponseCodeNotimp    DNSTimeseriesGroupCacheHitParamsResponseCode = "NOTIMP"
	DNSTimeseriesGroupCacheHitParamsResponseCodeRefused   DNSTimeseriesGroupCacheHitParamsResponseCode = "REFUSED"
	DNSTimeseriesGroupCacheHitParamsResponseCodeYxdomain  DNSTimeseriesGroupCacheHitParamsResponseCode = "YXDOMAIN"
	DNSTimeseriesGroupCacheHitParamsResponseCodeYxrrset   DNSTimeseriesGroupCacheHitParamsResponseCode = "YXRRSET"
	DNSTimeseriesGroupCacheHitParamsResponseCodeNxrrset   DNSTimeseriesGroupCacheHitParamsResponseCode = "NXRRSET"
	DNSTimeseriesGroupCacheHitParamsResponseCodeNotauth   DNSTimeseriesGroupCacheHitParamsResponseCode = "NOTAUTH"
	DNSTimeseriesGroupCacheHitParamsResponseCodeNotzone   DNSTimeseriesGroupCacheHitParamsResponseCode = "NOTZONE"
	DNSTimeseriesGroupCacheHitParamsResponseCodeBadsig    DNSTimeseriesGroupCacheHitParamsResponseCode = "BADSIG"
	DNSTimeseriesGroupCacheHitParamsResponseCodeBadkey    DNSTimeseriesGroupCacheHitParamsResponseCode = "BADKEY"
	DNSTimeseriesGroupCacheHitParamsResponseCodeBadtime   DNSTimeseriesGroupCacheHitParamsResponseCode = "BADTIME"
	DNSTimeseriesGroupCacheHitParamsResponseCodeBadmode   DNSTimeseriesGroupCacheHitParamsResponseCode = "BADMODE"
	DNSTimeseriesGroupCacheHitParamsResponseCodeBadname   DNSTimeseriesGroupCacheHitParamsResponseCode = "BADNAME"
	DNSTimeseriesGroupCacheHitParamsResponseCodeBadalg    DNSTimeseriesGroupCacheHitParamsResponseCode = "BADALG"
	DNSTimeseriesGroupCacheHitParamsResponseCodeBadtrunc  DNSTimeseriesGroupCacheHitParamsResponseCode = "BADTRUNC"
	DNSTimeseriesGroupCacheHitParamsResponseCodeBadcookie DNSTimeseriesGroupCacheHitParamsResponseCode = "BADCOOKIE"
)

func (r DNSTimeseriesGroupCacheHitParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupCacheHitParamsResponseCodeNoerror, DNSTimeseriesGroupCacheHitParamsResponseCodeFormerr, DNSTimeseriesGroupCacheHitParamsResponseCodeServfail, DNSTimeseriesGroupCacheHitParamsResponseCodeNxdomain, DNSTimeseriesGroupCacheHitParamsResponseCodeNotimp, DNSTimeseriesGroupCacheHitParamsResponseCodeRefused, DNSTimeseriesGroupCacheHitParamsResponseCodeYxdomain, DNSTimeseriesGroupCacheHitParamsResponseCodeYxrrset, DNSTimeseriesGroupCacheHitParamsResponseCodeNxrrset, DNSTimeseriesGroupCacheHitParamsResponseCodeNotauth, DNSTimeseriesGroupCacheHitParamsResponseCodeNotzone, DNSTimeseriesGroupCacheHitParamsResponseCodeBadsig, DNSTimeseriesGroupCacheHitParamsResponseCodeBadkey, DNSTimeseriesGroupCacheHitParamsResponseCodeBadtime, DNSTimeseriesGroupCacheHitParamsResponseCodeBadmode, DNSTimeseriesGroupCacheHitParamsResponseCodeBadname, DNSTimeseriesGroupCacheHitParamsResponseCodeBadalg, DNSTimeseriesGroupCacheHitParamsResponseCodeBadtrunc, DNSTimeseriesGroupCacheHitParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSTimeseriesGroupCacheHitResponseEnvelope struct {
	Result  DNSTimeseriesGroupCacheHitResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    dnsTimeseriesGroupCacheHitResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupCacheHitResponseEnvelopeJSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupCacheHitResponseEnvelope]
type dnsTimeseriesGroupCacheHitResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupCacheHitResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupCacheHitResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupDNSSECParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[DNSTimeseriesGroupDNSSECParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[DNSTimeseriesGroupDNSSECParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSTimeseriesGroupDNSSECParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSTimeseriesGroupDNSSECParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSTimeseriesGroupDNSSECParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupDNSSECParams]'s query parameters as
// `url.Values`.
func (r DNSTimeseriesGroupDNSSECParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupDNSSECParamsAggInterval string

const (
	DNSTimeseriesGroupDNSSECParamsAggInterval15m DNSTimeseriesGroupDNSSECParamsAggInterval = "15m"
	DNSTimeseriesGroupDNSSECParamsAggInterval1h  DNSTimeseriesGroupDNSSECParamsAggInterval = "1h"
	DNSTimeseriesGroupDNSSECParamsAggInterval1d  DNSTimeseriesGroupDNSSECParamsAggInterval = "1d"
	DNSTimeseriesGroupDNSSECParamsAggInterval1w  DNSTimeseriesGroupDNSSECParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupDNSSECParamsAggInterval) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupDNSSECParamsAggInterval15m, DNSTimeseriesGroupDNSSECParamsAggInterval1h, DNSTimeseriesGroupDNSSECParamsAggInterval1d, DNSTimeseriesGroupDNSSECParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type DNSTimeseriesGroupDNSSECParamsFormat string

const (
	DNSTimeseriesGroupDNSSECParamsFormatJson DNSTimeseriesGroupDNSSECParamsFormat = "JSON"
	DNSTimeseriesGroupDNSSECParamsFormatCsv  DNSTimeseriesGroupDNSSECParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupDNSSECParamsFormat) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupDNSSECParamsFormatJson, DNSTimeseriesGroupDNSSECParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSTimeseriesGroupDNSSECParamsProtocol string

const (
	DNSTimeseriesGroupDNSSECParamsProtocolUdp   DNSTimeseriesGroupDNSSECParamsProtocol = "UDP"
	DNSTimeseriesGroupDNSSECParamsProtocolTCP   DNSTimeseriesGroupDNSSECParamsProtocol = "TCP"
	DNSTimeseriesGroupDNSSECParamsProtocolHTTPS DNSTimeseriesGroupDNSSECParamsProtocol = "HTTPS"
	DNSTimeseriesGroupDNSSECParamsProtocolTLS   DNSTimeseriesGroupDNSSECParamsProtocol = "TLS"
)

func (r DNSTimeseriesGroupDNSSECParamsProtocol) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupDNSSECParamsProtocolUdp, DNSTimeseriesGroupDNSSECParamsProtocolTCP, DNSTimeseriesGroupDNSSECParamsProtocolHTTPS, DNSTimeseriesGroupDNSSECParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSTimeseriesGroupDNSSECParamsQueryType string

const (
	DNSTimeseriesGroupDNSSECParamsQueryTypeA          DNSTimeseriesGroupDNSSECParamsQueryType = "A"
	DNSTimeseriesGroupDNSSECParamsQueryTypeAAAA       DNSTimeseriesGroupDNSSECParamsQueryType = "AAAA"
	DNSTimeseriesGroupDNSSECParamsQueryTypeA6         DNSTimeseriesGroupDNSSECParamsQueryType = "A6"
	DNSTimeseriesGroupDNSSECParamsQueryTypeAfsdb      DNSTimeseriesGroupDNSSECParamsQueryType = "AFSDB"
	DNSTimeseriesGroupDNSSECParamsQueryTypeAny        DNSTimeseriesGroupDNSSECParamsQueryType = "ANY"
	DNSTimeseriesGroupDNSSECParamsQueryTypeApl        DNSTimeseriesGroupDNSSECParamsQueryType = "APL"
	DNSTimeseriesGroupDNSSECParamsQueryTypeAtma       DNSTimeseriesGroupDNSSECParamsQueryType = "ATMA"
	DNSTimeseriesGroupDNSSECParamsQueryTypeAXFR       DNSTimeseriesGroupDNSSECParamsQueryType = "AXFR"
	DNSTimeseriesGroupDNSSECParamsQueryTypeCAA        DNSTimeseriesGroupDNSSECParamsQueryType = "CAA"
	DNSTimeseriesGroupDNSSECParamsQueryTypeCdnskey    DNSTimeseriesGroupDNSSECParamsQueryType = "CDNSKEY"
	DNSTimeseriesGroupDNSSECParamsQueryTypeCds        DNSTimeseriesGroupDNSSECParamsQueryType = "CDS"
	DNSTimeseriesGroupDNSSECParamsQueryTypeCERT       DNSTimeseriesGroupDNSSECParamsQueryType = "CERT"
	DNSTimeseriesGroupDNSSECParamsQueryTypeCNAME      DNSTimeseriesGroupDNSSECParamsQueryType = "CNAME"
	DNSTimeseriesGroupDNSSECParamsQueryTypeCsync      DNSTimeseriesGroupDNSSECParamsQueryType = "CSYNC"
	DNSTimeseriesGroupDNSSECParamsQueryTypeDhcid      DNSTimeseriesGroupDNSSECParamsQueryType = "DHCID"
	DNSTimeseriesGroupDNSSECParamsQueryTypeDlv        DNSTimeseriesGroupDNSSECParamsQueryType = "DLV"
	DNSTimeseriesGroupDNSSECParamsQueryTypeDname      DNSTimeseriesGroupDNSSECParamsQueryType = "DNAME"
	DNSTimeseriesGroupDNSSECParamsQueryTypeDNSKEY     DNSTimeseriesGroupDNSSECParamsQueryType = "DNSKEY"
	DNSTimeseriesGroupDNSSECParamsQueryTypeDoa        DNSTimeseriesGroupDNSSECParamsQueryType = "DOA"
	DNSTimeseriesGroupDNSSECParamsQueryTypeDS         DNSTimeseriesGroupDNSSECParamsQueryType = "DS"
	DNSTimeseriesGroupDNSSECParamsQueryTypeEid        DNSTimeseriesGroupDNSSECParamsQueryType = "EID"
	DNSTimeseriesGroupDNSSECParamsQueryTypeEui48      DNSTimeseriesGroupDNSSECParamsQueryType = "EUI48"
	DNSTimeseriesGroupDNSSECParamsQueryTypeEui64      DNSTimeseriesGroupDNSSECParamsQueryType = "EUI64"
	DNSTimeseriesGroupDNSSECParamsQueryTypeGpos       DNSTimeseriesGroupDNSSECParamsQueryType = "GPOS"
	DNSTimeseriesGroupDNSSECParamsQueryTypeGid        DNSTimeseriesGroupDNSSECParamsQueryType = "GID"
	DNSTimeseriesGroupDNSSECParamsQueryTypeHinfo      DNSTimeseriesGroupDNSSECParamsQueryType = "HINFO"
	DNSTimeseriesGroupDNSSECParamsQueryTypeHip        DNSTimeseriesGroupDNSSECParamsQueryType = "HIP"
	DNSTimeseriesGroupDNSSECParamsQueryTypeHTTPS      DNSTimeseriesGroupDNSSECParamsQueryType = "HTTPS"
	DNSTimeseriesGroupDNSSECParamsQueryTypeIpseckey   DNSTimeseriesGroupDNSSECParamsQueryType = "IPSECKEY"
	DNSTimeseriesGroupDNSSECParamsQueryTypeIsdn       DNSTimeseriesGroupDNSSECParamsQueryType = "ISDN"
	DNSTimeseriesGroupDNSSECParamsQueryTypeIxfr       DNSTimeseriesGroupDNSSECParamsQueryType = "IXFR"
	DNSTimeseriesGroupDNSSECParamsQueryTypeKey        DNSTimeseriesGroupDNSSECParamsQueryType = "KEY"
	DNSTimeseriesGroupDNSSECParamsQueryTypeKx         DNSTimeseriesGroupDNSSECParamsQueryType = "KX"
	DNSTimeseriesGroupDNSSECParamsQueryTypeL32        DNSTimeseriesGroupDNSSECParamsQueryType = "L32"
	DNSTimeseriesGroupDNSSECParamsQueryTypeL64        DNSTimeseriesGroupDNSSECParamsQueryType = "L64"
	DNSTimeseriesGroupDNSSECParamsQueryTypeLOC        DNSTimeseriesGroupDNSSECParamsQueryType = "LOC"
	DNSTimeseriesGroupDNSSECParamsQueryTypeLp         DNSTimeseriesGroupDNSSECParamsQueryType = "LP"
	DNSTimeseriesGroupDNSSECParamsQueryTypeMaila      DNSTimeseriesGroupDNSSECParamsQueryType = "MAILA"
	DNSTimeseriesGroupDNSSECParamsQueryTypeMailb      DNSTimeseriesGroupDNSSECParamsQueryType = "MAILB"
	DNSTimeseriesGroupDNSSECParamsQueryTypeMB         DNSTimeseriesGroupDNSSECParamsQueryType = "MB"
	DNSTimeseriesGroupDNSSECParamsQueryTypeMd         DNSTimeseriesGroupDNSSECParamsQueryType = "MD"
	DNSTimeseriesGroupDNSSECParamsQueryTypeMf         DNSTimeseriesGroupDNSSECParamsQueryType = "MF"
	DNSTimeseriesGroupDNSSECParamsQueryTypeMg         DNSTimeseriesGroupDNSSECParamsQueryType = "MG"
	DNSTimeseriesGroupDNSSECParamsQueryTypeMinfo      DNSTimeseriesGroupDNSSECParamsQueryType = "MINFO"
	DNSTimeseriesGroupDNSSECParamsQueryTypeMr         DNSTimeseriesGroupDNSSECParamsQueryType = "MR"
	DNSTimeseriesGroupDNSSECParamsQueryTypeMX         DNSTimeseriesGroupDNSSECParamsQueryType = "MX"
	DNSTimeseriesGroupDNSSECParamsQueryTypeNAPTR      DNSTimeseriesGroupDNSSECParamsQueryType = "NAPTR"
	DNSTimeseriesGroupDNSSECParamsQueryTypeNb         DNSTimeseriesGroupDNSSECParamsQueryType = "NB"
	DNSTimeseriesGroupDNSSECParamsQueryTypeNbstat     DNSTimeseriesGroupDNSSECParamsQueryType = "NBSTAT"
	DNSTimeseriesGroupDNSSECParamsQueryTypeNid        DNSTimeseriesGroupDNSSECParamsQueryType = "NID"
	DNSTimeseriesGroupDNSSECParamsQueryTypeNimloc     DNSTimeseriesGroupDNSSECParamsQueryType = "NIMLOC"
	DNSTimeseriesGroupDNSSECParamsQueryTypeNinfo      DNSTimeseriesGroupDNSSECParamsQueryType = "NINFO"
	DNSTimeseriesGroupDNSSECParamsQueryTypeNS         DNSTimeseriesGroupDNSSECParamsQueryType = "NS"
	DNSTimeseriesGroupDNSSECParamsQueryTypeNsap       DNSTimeseriesGroupDNSSECParamsQueryType = "NSAP"
	DNSTimeseriesGroupDNSSECParamsQueryTypeNsec       DNSTimeseriesGroupDNSSECParamsQueryType = "NSEC"
	DNSTimeseriesGroupDNSSECParamsQueryTypeNsec3      DNSTimeseriesGroupDNSSECParamsQueryType = "NSEC3"
	DNSTimeseriesGroupDNSSECParamsQueryTypeNsec3Param DNSTimeseriesGroupDNSSECParamsQueryType = "NSEC3PARAM"
	DNSTimeseriesGroupDNSSECParamsQueryTypeNull       DNSTimeseriesGroupDNSSECParamsQueryType = "NULL"
	DNSTimeseriesGroupDNSSECParamsQueryTypeNxt        DNSTimeseriesGroupDNSSECParamsQueryType = "NXT"
	DNSTimeseriesGroupDNSSECParamsQueryTypeOpenpgpkey DNSTimeseriesGroupDNSSECParamsQueryType = "OPENPGPKEY"
	DNSTimeseriesGroupDNSSECParamsQueryTypeOpt        DNSTimeseriesGroupDNSSECParamsQueryType = "OPT"
	DNSTimeseriesGroupDNSSECParamsQueryTypePTR        DNSTimeseriesGroupDNSSECParamsQueryType = "PTR"
	DNSTimeseriesGroupDNSSECParamsQueryTypePx         DNSTimeseriesGroupDNSSECParamsQueryType = "PX"
	DNSTimeseriesGroupDNSSECParamsQueryTypeRkey       DNSTimeseriesGroupDNSSECParamsQueryType = "RKEY"
	DNSTimeseriesGroupDNSSECParamsQueryTypeRp         DNSTimeseriesGroupDNSSECParamsQueryType = "RP"
	DNSTimeseriesGroupDNSSECParamsQueryTypeRrsig      DNSTimeseriesGroupDNSSECParamsQueryType = "RRSIG"
	DNSTimeseriesGroupDNSSECParamsQueryTypeRt         DNSTimeseriesGroupDNSSECParamsQueryType = "RT"
	DNSTimeseriesGroupDNSSECParamsQueryTypeSig        DNSTimeseriesGroupDNSSECParamsQueryType = "SIG"
	DNSTimeseriesGroupDNSSECParamsQueryTypeSink       DNSTimeseriesGroupDNSSECParamsQueryType = "SINK"
	DNSTimeseriesGroupDNSSECParamsQueryTypeSMIMEA     DNSTimeseriesGroupDNSSECParamsQueryType = "SMIMEA"
	DNSTimeseriesGroupDNSSECParamsQueryTypeSOA        DNSTimeseriesGroupDNSSECParamsQueryType = "SOA"
	DNSTimeseriesGroupDNSSECParamsQueryTypeSPF        DNSTimeseriesGroupDNSSECParamsQueryType = "SPF"
	DNSTimeseriesGroupDNSSECParamsQueryTypeSRV        DNSTimeseriesGroupDNSSECParamsQueryType = "SRV"
	DNSTimeseriesGroupDNSSECParamsQueryTypeSSHFP      DNSTimeseriesGroupDNSSECParamsQueryType = "SSHFP"
	DNSTimeseriesGroupDNSSECParamsQueryTypeSVCB       DNSTimeseriesGroupDNSSECParamsQueryType = "SVCB"
	DNSTimeseriesGroupDNSSECParamsQueryTypeTa         DNSTimeseriesGroupDNSSECParamsQueryType = "TA"
	DNSTimeseriesGroupDNSSECParamsQueryTypeTalink     DNSTimeseriesGroupDNSSECParamsQueryType = "TALINK"
	DNSTimeseriesGroupDNSSECParamsQueryTypeTkey       DNSTimeseriesGroupDNSSECParamsQueryType = "TKEY"
	DNSTimeseriesGroupDNSSECParamsQueryTypeTLSA       DNSTimeseriesGroupDNSSECParamsQueryType = "TLSA"
	DNSTimeseriesGroupDNSSECParamsQueryTypeTSIG       DNSTimeseriesGroupDNSSECParamsQueryType = "TSIG"
	DNSTimeseriesGroupDNSSECParamsQueryTypeTXT        DNSTimeseriesGroupDNSSECParamsQueryType = "TXT"
	DNSTimeseriesGroupDNSSECParamsQueryTypeUinfo      DNSTimeseriesGroupDNSSECParamsQueryType = "UINFO"
	DNSTimeseriesGroupDNSSECParamsQueryTypeUID        DNSTimeseriesGroupDNSSECParamsQueryType = "UID"
	DNSTimeseriesGroupDNSSECParamsQueryTypeUnspec     DNSTimeseriesGroupDNSSECParamsQueryType = "UNSPEC"
	DNSTimeseriesGroupDNSSECParamsQueryTypeURI        DNSTimeseriesGroupDNSSECParamsQueryType = "URI"
	DNSTimeseriesGroupDNSSECParamsQueryTypeWks        DNSTimeseriesGroupDNSSECParamsQueryType = "WKS"
	DNSTimeseriesGroupDNSSECParamsQueryTypeX25        DNSTimeseriesGroupDNSSECParamsQueryType = "X25"
	DNSTimeseriesGroupDNSSECParamsQueryTypeZonemd     DNSTimeseriesGroupDNSSECParamsQueryType = "ZONEMD"
)

func (r DNSTimeseriesGroupDNSSECParamsQueryType) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupDNSSECParamsQueryTypeA, DNSTimeseriesGroupDNSSECParamsQueryTypeAAAA, DNSTimeseriesGroupDNSSECParamsQueryTypeA6, DNSTimeseriesGroupDNSSECParamsQueryTypeAfsdb, DNSTimeseriesGroupDNSSECParamsQueryTypeAny, DNSTimeseriesGroupDNSSECParamsQueryTypeApl, DNSTimeseriesGroupDNSSECParamsQueryTypeAtma, DNSTimeseriesGroupDNSSECParamsQueryTypeAXFR, DNSTimeseriesGroupDNSSECParamsQueryTypeCAA, DNSTimeseriesGroupDNSSECParamsQueryTypeCdnskey, DNSTimeseriesGroupDNSSECParamsQueryTypeCds, DNSTimeseriesGroupDNSSECParamsQueryTypeCERT, DNSTimeseriesGroupDNSSECParamsQueryTypeCNAME, DNSTimeseriesGroupDNSSECParamsQueryTypeCsync, DNSTimeseriesGroupDNSSECParamsQueryTypeDhcid, DNSTimeseriesGroupDNSSECParamsQueryTypeDlv, DNSTimeseriesGroupDNSSECParamsQueryTypeDname, DNSTimeseriesGroupDNSSECParamsQueryTypeDNSKEY, DNSTimeseriesGroupDNSSECParamsQueryTypeDoa, DNSTimeseriesGroupDNSSECParamsQueryTypeDS, DNSTimeseriesGroupDNSSECParamsQueryTypeEid, DNSTimeseriesGroupDNSSECParamsQueryTypeEui48, DNSTimeseriesGroupDNSSECParamsQueryTypeEui64, DNSTimeseriesGroupDNSSECParamsQueryTypeGpos, DNSTimeseriesGroupDNSSECParamsQueryTypeGid, DNSTimeseriesGroupDNSSECParamsQueryTypeHinfo, DNSTimeseriesGroupDNSSECParamsQueryTypeHip, DNSTimeseriesGroupDNSSECParamsQueryTypeHTTPS, DNSTimeseriesGroupDNSSECParamsQueryTypeIpseckey, DNSTimeseriesGroupDNSSECParamsQueryTypeIsdn, DNSTimeseriesGroupDNSSECParamsQueryTypeIxfr, DNSTimeseriesGroupDNSSECParamsQueryTypeKey, DNSTimeseriesGroupDNSSECParamsQueryTypeKx, DNSTimeseriesGroupDNSSECParamsQueryTypeL32, DNSTimeseriesGroupDNSSECParamsQueryTypeL64, DNSTimeseriesGroupDNSSECParamsQueryTypeLOC, DNSTimeseriesGroupDNSSECParamsQueryTypeLp, DNSTimeseriesGroupDNSSECParamsQueryTypeMaila, DNSTimeseriesGroupDNSSECParamsQueryTypeMailb, DNSTimeseriesGroupDNSSECParamsQueryTypeMB, DNSTimeseriesGroupDNSSECParamsQueryTypeMd, DNSTimeseriesGroupDNSSECParamsQueryTypeMf, DNSTimeseriesGroupDNSSECParamsQueryTypeMg, DNSTimeseriesGroupDNSSECParamsQueryTypeMinfo, DNSTimeseriesGroupDNSSECParamsQueryTypeMr, DNSTimeseriesGroupDNSSECParamsQueryTypeMX, DNSTimeseriesGroupDNSSECParamsQueryTypeNAPTR, DNSTimeseriesGroupDNSSECParamsQueryTypeNb, DNSTimeseriesGroupDNSSECParamsQueryTypeNbstat, DNSTimeseriesGroupDNSSECParamsQueryTypeNid, DNSTimeseriesGroupDNSSECParamsQueryTypeNimloc, DNSTimeseriesGroupDNSSECParamsQueryTypeNinfo, DNSTimeseriesGroupDNSSECParamsQueryTypeNS, DNSTimeseriesGroupDNSSECParamsQueryTypeNsap, DNSTimeseriesGroupDNSSECParamsQueryTypeNsec, DNSTimeseriesGroupDNSSECParamsQueryTypeNsec3, DNSTimeseriesGroupDNSSECParamsQueryTypeNsec3Param, DNSTimeseriesGroupDNSSECParamsQueryTypeNull, DNSTimeseriesGroupDNSSECParamsQueryTypeNxt, DNSTimeseriesGroupDNSSECParamsQueryTypeOpenpgpkey, DNSTimeseriesGroupDNSSECParamsQueryTypeOpt, DNSTimeseriesGroupDNSSECParamsQueryTypePTR, DNSTimeseriesGroupDNSSECParamsQueryTypePx, DNSTimeseriesGroupDNSSECParamsQueryTypeRkey, DNSTimeseriesGroupDNSSECParamsQueryTypeRp, DNSTimeseriesGroupDNSSECParamsQueryTypeRrsig, DNSTimeseriesGroupDNSSECParamsQueryTypeRt, DNSTimeseriesGroupDNSSECParamsQueryTypeSig, DNSTimeseriesGroupDNSSECParamsQueryTypeSink, DNSTimeseriesGroupDNSSECParamsQueryTypeSMIMEA, DNSTimeseriesGroupDNSSECParamsQueryTypeSOA, DNSTimeseriesGroupDNSSECParamsQueryTypeSPF, DNSTimeseriesGroupDNSSECParamsQueryTypeSRV, DNSTimeseriesGroupDNSSECParamsQueryTypeSSHFP, DNSTimeseriesGroupDNSSECParamsQueryTypeSVCB, DNSTimeseriesGroupDNSSECParamsQueryTypeTa, DNSTimeseriesGroupDNSSECParamsQueryTypeTalink, DNSTimeseriesGroupDNSSECParamsQueryTypeTkey, DNSTimeseriesGroupDNSSECParamsQueryTypeTLSA, DNSTimeseriesGroupDNSSECParamsQueryTypeTSIG, DNSTimeseriesGroupDNSSECParamsQueryTypeTXT, DNSTimeseriesGroupDNSSECParamsQueryTypeUinfo, DNSTimeseriesGroupDNSSECParamsQueryTypeUID, DNSTimeseriesGroupDNSSECParamsQueryTypeUnspec, DNSTimeseriesGroupDNSSECParamsQueryTypeURI, DNSTimeseriesGroupDNSSECParamsQueryTypeWks, DNSTimeseriesGroupDNSSECParamsQueryTypeX25, DNSTimeseriesGroupDNSSECParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSTimeseriesGroupDNSSECParamsResponseCode string

const (
	DNSTimeseriesGroupDNSSECParamsResponseCodeNoerror   DNSTimeseriesGroupDNSSECParamsResponseCode = "NOERROR"
	DNSTimeseriesGroupDNSSECParamsResponseCodeFormerr   DNSTimeseriesGroupDNSSECParamsResponseCode = "FORMERR"
	DNSTimeseriesGroupDNSSECParamsResponseCodeServfail  DNSTimeseriesGroupDNSSECParamsResponseCode = "SERVFAIL"
	DNSTimeseriesGroupDNSSECParamsResponseCodeNxdomain  DNSTimeseriesGroupDNSSECParamsResponseCode = "NXDOMAIN"
	DNSTimeseriesGroupDNSSECParamsResponseCodeNotimp    DNSTimeseriesGroupDNSSECParamsResponseCode = "NOTIMP"
	DNSTimeseriesGroupDNSSECParamsResponseCodeRefused   DNSTimeseriesGroupDNSSECParamsResponseCode = "REFUSED"
	DNSTimeseriesGroupDNSSECParamsResponseCodeYxdomain  DNSTimeseriesGroupDNSSECParamsResponseCode = "YXDOMAIN"
	DNSTimeseriesGroupDNSSECParamsResponseCodeYxrrset   DNSTimeseriesGroupDNSSECParamsResponseCode = "YXRRSET"
	DNSTimeseriesGroupDNSSECParamsResponseCodeNxrrset   DNSTimeseriesGroupDNSSECParamsResponseCode = "NXRRSET"
	DNSTimeseriesGroupDNSSECParamsResponseCodeNotauth   DNSTimeseriesGroupDNSSECParamsResponseCode = "NOTAUTH"
	DNSTimeseriesGroupDNSSECParamsResponseCodeNotzone   DNSTimeseriesGroupDNSSECParamsResponseCode = "NOTZONE"
	DNSTimeseriesGroupDNSSECParamsResponseCodeBadsig    DNSTimeseriesGroupDNSSECParamsResponseCode = "BADSIG"
	DNSTimeseriesGroupDNSSECParamsResponseCodeBadkey    DNSTimeseriesGroupDNSSECParamsResponseCode = "BADKEY"
	DNSTimeseriesGroupDNSSECParamsResponseCodeBadtime   DNSTimeseriesGroupDNSSECParamsResponseCode = "BADTIME"
	DNSTimeseriesGroupDNSSECParamsResponseCodeBadmode   DNSTimeseriesGroupDNSSECParamsResponseCode = "BADMODE"
	DNSTimeseriesGroupDNSSECParamsResponseCodeBadname   DNSTimeseriesGroupDNSSECParamsResponseCode = "BADNAME"
	DNSTimeseriesGroupDNSSECParamsResponseCodeBadalg    DNSTimeseriesGroupDNSSECParamsResponseCode = "BADALG"
	DNSTimeseriesGroupDNSSECParamsResponseCodeBadtrunc  DNSTimeseriesGroupDNSSECParamsResponseCode = "BADTRUNC"
	DNSTimeseriesGroupDNSSECParamsResponseCodeBadcookie DNSTimeseriesGroupDNSSECParamsResponseCode = "BADCOOKIE"
)

func (r DNSTimeseriesGroupDNSSECParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupDNSSECParamsResponseCodeNoerror, DNSTimeseriesGroupDNSSECParamsResponseCodeFormerr, DNSTimeseriesGroupDNSSECParamsResponseCodeServfail, DNSTimeseriesGroupDNSSECParamsResponseCodeNxdomain, DNSTimeseriesGroupDNSSECParamsResponseCodeNotimp, DNSTimeseriesGroupDNSSECParamsResponseCodeRefused, DNSTimeseriesGroupDNSSECParamsResponseCodeYxdomain, DNSTimeseriesGroupDNSSECParamsResponseCodeYxrrset, DNSTimeseriesGroupDNSSECParamsResponseCodeNxrrset, DNSTimeseriesGroupDNSSECParamsResponseCodeNotauth, DNSTimeseriesGroupDNSSECParamsResponseCodeNotzone, DNSTimeseriesGroupDNSSECParamsResponseCodeBadsig, DNSTimeseriesGroupDNSSECParamsResponseCodeBadkey, DNSTimeseriesGroupDNSSECParamsResponseCodeBadtime, DNSTimeseriesGroupDNSSECParamsResponseCodeBadmode, DNSTimeseriesGroupDNSSECParamsResponseCodeBadname, DNSTimeseriesGroupDNSSECParamsResponseCodeBadalg, DNSTimeseriesGroupDNSSECParamsResponseCodeBadtrunc, DNSTimeseriesGroupDNSSECParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSTimeseriesGroupDNSSECResponseEnvelope struct {
	Result  DNSTimeseriesGroupDNSSECResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    dnsTimeseriesGroupDNSSECResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupDNSSECResponseEnvelopeJSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupDNSSECResponseEnvelope]
type dnsTimeseriesGroupDNSSECResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupDNSSECResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupDNSSECResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupDNSSECAwareParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[DNSTimeseriesGroupDNSSECAwareParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[DNSTimeseriesGroupDNSSECAwareParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSTimeseriesGroupDNSSECAwareParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSTimeseriesGroupDNSSECAwareParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSTimeseriesGroupDNSSECAwareParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupDNSSECAwareParams]'s query parameters as
// `url.Values`.
func (r DNSTimeseriesGroupDNSSECAwareParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupDNSSECAwareParamsAggInterval string

const (
	DNSTimeseriesGroupDNSSECAwareParamsAggInterval15m DNSTimeseriesGroupDNSSECAwareParamsAggInterval = "15m"
	DNSTimeseriesGroupDNSSECAwareParamsAggInterval1h  DNSTimeseriesGroupDNSSECAwareParamsAggInterval = "1h"
	DNSTimeseriesGroupDNSSECAwareParamsAggInterval1d  DNSTimeseriesGroupDNSSECAwareParamsAggInterval = "1d"
	DNSTimeseriesGroupDNSSECAwareParamsAggInterval1w  DNSTimeseriesGroupDNSSECAwareParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupDNSSECAwareParamsAggInterval) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupDNSSECAwareParamsAggInterval15m, DNSTimeseriesGroupDNSSECAwareParamsAggInterval1h, DNSTimeseriesGroupDNSSECAwareParamsAggInterval1d, DNSTimeseriesGroupDNSSECAwareParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type DNSTimeseriesGroupDNSSECAwareParamsFormat string

const (
	DNSTimeseriesGroupDNSSECAwareParamsFormatJson DNSTimeseriesGroupDNSSECAwareParamsFormat = "JSON"
	DNSTimeseriesGroupDNSSECAwareParamsFormatCsv  DNSTimeseriesGroupDNSSECAwareParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupDNSSECAwareParamsFormat) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupDNSSECAwareParamsFormatJson, DNSTimeseriesGroupDNSSECAwareParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSTimeseriesGroupDNSSECAwareParamsProtocol string

const (
	DNSTimeseriesGroupDNSSECAwareParamsProtocolUdp   DNSTimeseriesGroupDNSSECAwareParamsProtocol = "UDP"
	DNSTimeseriesGroupDNSSECAwareParamsProtocolTCP   DNSTimeseriesGroupDNSSECAwareParamsProtocol = "TCP"
	DNSTimeseriesGroupDNSSECAwareParamsProtocolHTTPS DNSTimeseriesGroupDNSSECAwareParamsProtocol = "HTTPS"
	DNSTimeseriesGroupDNSSECAwareParamsProtocolTLS   DNSTimeseriesGroupDNSSECAwareParamsProtocol = "TLS"
)

func (r DNSTimeseriesGroupDNSSECAwareParamsProtocol) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupDNSSECAwareParamsProtocolUdp, DNSTimeseriesGroupDNSSECAwareParamsProtocolTCP, DNSTimeseriesGroupDNSSECAwareParamsProtocolHTTPS, DNSTimeseriesGroupDNSSECAwareParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSTimeseriesGroupDNSSECAwareParamsQueryType string

const (
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeA          DNSTimeseriesGroupDNSSECAwareParamsQueryType = "A"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeAAAA       DNSTimeseriesGroupDNSSECAwareParamsQueryType = "AAAA"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeA6         DNSTimeseriesGroupDNSSECAwareParamsQueryType = "A6"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeAfsdb      DNSTimeseriesGroupDNSSECAwareParamsQueryType = "AFSDB"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeAny        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "ANY"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeApl        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "APL"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeAtma       DNSTimeseriesGroupDNSSECAwareParamsQueryType = "ATMA"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeAXFR       DNSTimeseriesGroupDNSSECAwareParamsQueryType = "AXFR"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeCAA        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "CAA"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeCdnskey    DNSTimeseriesGroupDNSSECAwareParamsQueryType = "CDNSKEY"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeCds        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "CDS"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeCERT       DNSTimeseriesGroupDNSSECAwareParamsQueryType = "CERT"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeCNAME      DNSTimeseriesGroupDNSSECAwareParamsQueryType = "CNAME"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeCsync      DNSTimeseriesGroupDNSSECAwareParamsQueryType = "CSYNC"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeDhcid      DNSTimeseriesGroupDNSSECAwareParamsQueryType = "DHCID"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeDlv        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "DLV"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeDname      DNSTimeseriesGroupDNSSECAwareParamsQueryType = "DNAME"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeDNSKEY     DNSTimeseriesGroupDNSSECAwareParamsQueryType = "DNSKEY"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeDoa        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "DOA"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeDS         DNSTimeseriesGroupDNSSECAwareParamsQueryType = "DS"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeEid        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "EID"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeEui48      DNSTimeseriesGroupDNSSECAwareParamsQueryType = "EUI48"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeEui64      DNSTimeseriesGroupDNSSECAwareParamsQueryType = "EUI64"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeGpos       DNSTimeseriesGroupDNSSECAwareParamsQueryType = "GPOS"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeGid        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "GID"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeHinfo      DNSTimeseriesGroupDNSSECAwareParamsQueryType = "HINFO"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeHip        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "HIP"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeHTTPS      DNSTimeseriesGroupDNSSECAwareParamsQueryType = "HTTPS"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeIpseckey   DNSTimeseriesGroupDNSSECAwareParamsQueryType = "IPSECKEY"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeIsdn       DNSTimeseriesGroupDNSSECAwareParamsQueryType = "ISDN"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeIxfr       DNSTimeseriesGroupDNSSECAwareParamsQueryType = "IXFR"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeKey        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "KEY"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeKx         DNSTimeseriesGroupDNSSECAwareParamsQueryType = "KX"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeL32        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "L32"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeL64        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "L64"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeLOC        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "LOC"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeLp         DNSTimeseriesGroupDNSSECAwareParamsQueryType = "LP"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeMaila      DNSTimeseriesGroupDNSSECAwareParamsQueryType = "MAILA"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeMailb      DNSTimeseriesGroupDNSSECAwareParamsQueryType = "MAILB"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeMB         DNSTimeseriesGroupDNSSECAwareParamsQueryType = "MB"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeMd         DNSTimeseriesGroupDNSSECAwareParamsQueryType = "MD"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeMf         DNSTimeseriesGroupDNSSECAwareParamsQueryType = "MF"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeMg         DNSTimeseriesGroupDNSSECAwareParamsQueryType = "MG"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeMinfo      DNSTimeseriesGroupDNSSECAwareParamsQueryType = "MINFO"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeMr         DNSTimeseriesGroupDNSSECAwareParamsQueryType = "MR"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeMX         DNSTimeseriesGroupDNSSECAwareParamsQueryType = "MX"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNAPTR      DNSTimeseriesGroupDNSSECAwareParamsQueryType = "NAPTR"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNb         DNSTimeseriesGroupDNSSECAwareParamsQueryType = "NB"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNbstat     DNSTimeseriesGroupDNSSECAwareParamsQueryType = "NBSTAT"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNid        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "NID"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNimloc     DNSTimeseriesGroupDNSSECAwareParamsQueryType = "NIMLOC"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNinfo      DNSTimeseriesGroupDNSSECAwareParamsQueryType = "NINFO"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNS         DNSTimeseriesGroupDNSSECAwareParamsQueryType = "NS"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNsap       DNSTimeseriesGroupDNSSECAwareParamsQueryType = "NSAP"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNsec       DNSTimeseriesGroupDNSSECAwareParamsQueryType = "NSEC"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNsec3      DNSTimeseriesGroupDNSSECAwareParamsQueryType = "NSEC3"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNsec3Param DNSTimeseriesGroupDNSSECAwareParamsQueryType = "NSEC3PARAM"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNull       DNSTimeseriesGroupDNSSECAwareParamsQueryType = "NULL"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNxt        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "NXT"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeOpenpgpkey DNSTimeseriesGroupDNSSECAwareParamsQueryType = "OPENPGPKEY"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeOpt        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "OPT"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypePTR        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "PTR"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypePx         DNSTimeseriesGroupDNSSECAwareParamsQueryType = "PX"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeRkey       DNSTimeseriesGroupDNSSECAwareParamsQueryType = "RKEY"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeRp         DNSTimeseriesGroupDNSSECAwareParamsQueryType = "RP"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeRrsig      DNSTimeseriesGroupDNSSECAwareParamsQueryType = "RRSIG"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeRt         DNSTimeseriesGroupDNSSECAwareParamsQueryType = "RT"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeSig        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "SIG"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeSink       DNSTimeseriesGroupDNSSECAwareParamsQueryType = "SINK"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeSMIMEA     DNSTimeseriesGroupDNSSECAwareParamsQueryType = "SMIMEA"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeSOA        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "SOA"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeSPF        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "SPF"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeSRV        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "SRV"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeSSHFP      DNSTimeseriesGroupDNSSECAwareParamsQueryType = "SSHFP"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeSVCB       DNSTimeseriesGroupDNSSECAwareParamsQueryType = "SVCB"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeTa         DNSTimeseriesGroupDNSSECAwareParamsQueryType = "TA"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeTalink     DNSTimeseriesGroupDNSSECAwareParamsQueryType = "TALINK"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeTkey       DNSTimeseriesGroupDNSSECAwareParamsQueryType = "TKEY"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeTLSA       DNSTimeseriesGroupDNSSECAwareParamsQueryType = "TLSA"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeTSIG       DNSTimeseriesGroupDNSSECAwareParamsQueryType = "TSIG"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeTXT        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "TXT"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeUinfo      DNSTimeseriesGroupDNSSECAwareParamsQueryType = "UINFO"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeUID        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "UID"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeUnspec     DNSTimeseriesGroupDNSSECAwareParamsQueryType = "UNSPEC"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeURI        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "URI"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeWks        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "WKS"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeX25        DNSTimeseriesGroupDNSSECAwareParamsQueryType = "X25"
	DNSTimeseriesGroupDNSSECAwareParamsQueryTypeZonemd     DNSTimeseriesGroupDNSSECAwareParamsQueryType = "ZONEMD"
)

func (r DNSTimeseriesGroupDNSSECAwareParamsQueryType) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupDNSSECAwareParamsQueryTypeA, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeAAAA, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeA6, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeAfsdb, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeAny, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeApl, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeAtma, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeAXFR, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeCAA, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeCdnskey, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeCds, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeCERT, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeCNAME, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeCsync, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeDhcid, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeDlv, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeDname, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeDNSKEY, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeDoa, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeDS, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeEid, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeEui48, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeEui64, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeGpos, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeGid, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeHinfo, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeHip, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeHTTPS, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeIpseckey, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeIsdn, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeIxfr, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeKey, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeKx, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeL32, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeL64, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeLOC, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeLp, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeMaila, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeMailb, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeMB, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeMd, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeMf, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeMg, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeMinfo, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeMr, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeMX, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNAPTR, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNb, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNbstat, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNid, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNimloc, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNinfo, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNS, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNsap, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNsec, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNsec3, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNsec3Param, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNull, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeNxt, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeOpenpgpkey, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeOpt, DNSTimeseriesGroupDNSSECAwareParamsQueryTypePTR, DNSTimeseriesGroupDNSSECAwareParamsQueryTypePx, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeRkey, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeRp, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeRrsig, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeRt, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeSig, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeSink, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeSMIMEA, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeSOA, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeSPF, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeSRV, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeSSHFP, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeSVCB, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeTa, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeTalink, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeTkey, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeTLSA, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeTSIG, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeTXT, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeUinfo, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeUID, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeUnspec, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeURI, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeWks, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeX25, DNSTimeseriesGroupDNSSECAwareParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSTimeseriesGroupDNSSECAwareParamsResponseCode string

const (
	DNSTimeseriesGroupDNSSECAwareParamsResponseCodeNoerror   DNSTimeseriesGroupDNSSECAwareParamsResponseCode = "NOERROR"
	DNSTimeseriesGroupDNSSECAwareParamsResponseCodeFormerr   DNSTimeseriesGroupDNSSECAwareParamsResponseCode = "FORMERR"
	DNSTimeseriesGroupDNSSECAwareParamsResponseCodeServfail  DNSTimeseriesGroupDNSSECAwareParamsResponseCode = "SERVFAIL"
	DNSTimeseriesGroupDNSSECAwareParamsResponseCodeNxdomain  DNSTimeseriesGroupDNSSECAwareParamsResponseCode = "NXDOMAIN"
	DNSTimeseriesGroupDNSSECAwareParamsResponseCodeNotimp    DNSTimeseriesGroupDNSSECAwareParamsResponseCode = "NOTIMP"
	DNSTimeseriesGroupDNSSECAwareParamsResponseCodeRefused   DNSTimeseriesGroupDNSSECAwareParamsResponseCode = "REFUSED"
	DNSTimeseriesGroupDNSSECAwareParamsResponseCodeYxdomain  DNSTimeseriesGroupDNSSECAwareParamsResponseCode = "YXDOMAIN"
	DNSTimeseriesGroupDNSSECAwareParamsResponseCodeYxrrset   DNSTimeseriesGroupDNSSECAwareParamsResponseCode = "YXRRSET"
	DNSTimeseriesGroupDNSSECAwareParamsResponseCodeNxrrset   DNSTimeseriesGroupDNSSECAwareParamsResponseCode = "NXRRSET"
	DNSTimeseriesGroupDNSSECAwareParamsResponseCodeNotauth   DNSTimeseriesGroupDNSSECAwareParamsResponseCode = "NOTAUTH"
	DNSTimeseriesGroupDNSSECAwareParamsResponseCodeNotzone   DNSTimeseriesGroupDNSSECAwareParamsResponseCode = "NOTZONE"
	DNSTimeseriesGroupDNSSECAwareParamsResponseCodeBadsig    DNSTimeseriesGroupDNSSECAwareParamsResponseCode = "BADSIG"
	DNSTimeseriesGroupDNSSECAwareParamsResponseCodeBadkey    DNSTimeseriesGroupDNSSECAwareParamsResponseCode = "BADKEY"
	DNSTimeseriesGroupDNSSECAwareParamsResponseCodeBadtime   DNSTimeseriesGroupDNSSECAwareParamsResponseCode = "BADTIME"
	DNSTimeseriesGroupDNSSECAwareParamsResponseCodeBadmode   DNSTimeseriesGroupDNSSECAwareParamsResponseCode = "BADMODE"
	DNSTimeseriesGroupDNSSECAwareParamsResponseCodeBadname   DNSTimeseriesGroupDNSSECAwareParamsResponseCode = "BADNAME"
	DNSTimeseriesGroupDNSSECAwareParamsResponseCodeBadalg    DNSTimeseriesGroupDNSSECAwareParamsResponseCode = "BADALG"
	DNSTimeseriesGroupDNSSECAwareParamsResponseCodeBadtrunc  DNSTimeseriesGroupDNSSECAwareParamsResponseCode = "BADTRUNC"
	DNSTimeseriesGroupDNSSECAwareParamsResponseCodeBadcookie DNSTimeseriesGroupDNSSECAwareParamsResponseCode = "BADCOOKIE"
)

func (r DNSTimeseriesGroupDNSSECAwareParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupDNSSECAwareParamsResponseCodeNoerror, DNSTimeseriesGroupDNSSECAwareParamsResponseCodeFormerr, DNSTimeseriesGroupDNSSECAwareParamsResponseCodeServfail, DNSTimeseriesGroupDNSSECAwareParamsResponseCodeNxdomain, DNSTimeseriesGroupDNSSECAwareParamsResponseCodeNotimp, DNSTimeseriesGroupDNSSECAwareParamsResponseCodeRefused, DNSTimeseriesGroupDNSSECAwareParamsResponseCodeYxdomain, DNSTimeseriesGroupDNSSECAwareParamsResponseCodeYxrrset, DNSTimeseriesGroupDNSSECAwareParamsResponseCodeNxrrset, DNSTimeseriesGroupDNSSECAwareParamsResponseCodeNotauth, DNSTimeseriesGroupDNSSECAwareParamsResponseCodeNotzone, DNSTimeseriesGroupDNSSECAwareParamsResponseCodeBadsig, DNSTimeseriesGroupDNSSECAwareParamsResponseCodeBadkey, DNSTimeseriesGroupDNSSECAwareParamsResponseCodeBadtime, DNSTimeseriesGroupDNSSECAwareParamsResponseCodeBadmode, DNSTimeseriesGroupDNSSECAwareParamsResponseCodeBadname, DNSTimeseriesGroupDNSSECAwareParamsResponseCodeBadalg, DNSTimeseriesGroupDNSSECAwareParamsResponseCodeBadtrunc, DNSTimeseriesGroupDNSSECAwareParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSTimeseriesGroupDNSSECAwareResponseEnvelope struct {
	Result  DNSTimeseriesGroupDNSSECAwareResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    dnsTimeseriesGroupDNSSECAwareResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupDNSSECAwareResponseEnvelopeJSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupDNSSECAwareResponseEnvelope]
type dnsTimeseriesGroupDNSSECAwareResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupDNSSECAwareResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupDNSSECAwareResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupDNSSECE2EParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[DNSTimeseriesGroupDnssece2EParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[DNSTimeseriesGroupDnssece2EParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSTimeseriesGroupDnssece2EParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSTimeseriesGroupDnssece2EParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSTimeseriesGroupDnssece2EParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupDNSSECE2EParams]'s query parameters as
// `url.Values`.
func (r DNSTimeseriesGroupDNSSECE2EParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupDnssece2EParamsAggInterval string

const (
	DNSTimeseriesGroupDnssece2EParamsAggInterval15m DNSTimeseriesGroupDnssece2EParamsAggInterval = "15m"
	DNSTimeseriesGroupDnssece2EParamsAggInterval1h  DNSTimeseriesGroupDnssece2EParamsAggInterval = "1h"
	DNSTimeseriesGroupDnssece2EParamsAggInterval1d  DNSTimeseriesGroupDnssece2EParamsAggInterval = "1d"
	DNSTimeseriesGroupDnssece2EParamsAggInterval1w  DNSTimeseriesGroupDnssece2EParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupDnssece2EParamsAggInterval) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupDnssece2EParamsAggInterval15m, DNSTimeseriesGroupDnssece2EParamsAggInterval1h, DNSTimeseriesGroupDnssece2EParamsAggInterval1d, DNSTimeseriesGroupDnssece2EParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type DNSTimeseriesGroupDnssece2EParamsFormat string

const (
	DNSTimeseriesGroupDnssece2EParamsFormatJson DNSTimeseriesGroupDnssece2EParamsFormat = "JSON"
	DNSTimeseriesGroupDnssece2EParamsFormatCsv  DNSTimeseriesGroupDnssece2EParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupDnssece2EParamsFormat) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupDnssece2EParamsFormatJson, DNSTimeseriesGroupDnssece2EParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSTimeseriesGroupDnssece2EParamsProtocol string

const (
	DNSTimeseriesGroupDnssece2EParamsProtocolUdp   DNSTimeseriesGroupDnssece2EParamsProtocol = "UDP"
	DNSTimeseriesGroupDnssece2EParamsProtocolTCP   DNSTimeseriesGroupDnssece2EParamsProtocol = "TCP"
	DNSTimeseriesGroupDnssece2EParamsProtocolHTTPS DNSTimeseriesGroupDnssece2EParamsProtocol = "HTTPS"
	DNSTimeseriesGroupDnssece2EParamsProtocolTLS   DNSTimeseriesGroupDnssece2EParamsProtocol = "TLS"
)

func (r DNSTimeseriesGroupDnssece2EParamsProtocol) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupDnssece2EParamsProtocolUdp, DNSTimeseriesGroupDnssece2EParamsProtocolTCP, DNSTimeseriesGroupDnssece2EParamsProtocolHTTPS, DNSTimeseriesGroupDnssece2EParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSTimeseriesGroupDnssece2EParamsQueryType string

const (
	DNSTimeseriesGroupDnssece2EParamsQueryTypeA          DNSTimeseriesGroupDnssece2EParamsQueryType = "A"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeAAAA       DNSTimeseriesGroupDnssece2EParamsQueryType = "AAAA"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeA6         DNSTimeseriesGroupDnssece2EParamsQueryType = "A6"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeAfsdb      DNSTimeseriesGroupDnssece2EParamsQueryType = "AFSDB"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeAny        DNSTimeseriesGroupDnssece2EParamsQueryType = "ANY"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeApl        DNSTimeseriesGroupDnssece2EParamsQueryType = "APL"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeAtma       DNSTimeseriesGroupDnssece2EParamsQueryType = "ATMA"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeAXFR       DNSTimeseriesGroupDnssece2EParamsQueryType = "AXFR"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeCAA        DNSTimeseriesGroupDnssece2EParamsQueryType = "CAA"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeCdnskey    DNSTimeseriesGroupDnssece2EParamsQueryType = "CDNSKEY"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeCds        DNSTimeseriesGroupDnssece2EParamsQueryType = "CDS"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeCERT       DNSTimeseriesGroupDnssece2EParamsQueryType = "CERT"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeCNAME      DNSTimeseriesGroupDnssece2EParamsQueryType = "CNAME"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeCsync      DNSTimeseriesGroupDnssece2EParamsQueryType = "CSYNC"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeDhcid      DNSTimeseriesGroupDnssece2EParamsQueryType = "DHCID"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeDlv        DNSTimeseriesGroupDnssece2EParamsQueryType = "DLV"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeDname      DNSTimeseriesGroupDnssece2EParamsQueryType = "DNAME"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeDNSKEY     DNSTimeseriesGroupDnssece2EParamsQueryType = "DNSKEY"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeDoa        DNSTimeseriesGroupDnssece2EParamsQueryType = "DOA"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeDS         DNSTimeseriesGroupDnssece2EParamsQueryType = "DS"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeEid        DNSTimeseriesGroupDnssece2EParamsQueryType = "EID"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeEui48      DNSTimeseriesGroupDnssece2EParamsQueryType = "EUI48"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeEui64      DNSTimeseriesGroupDnssece2EParamsQueryType = "EUI64"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeGpos       DNSTimeseriesGroupDnssece2EParamsQueryType = "GPOS"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeGid        DNSTimeseriesGroupDnssece2EParamsQueryType = "GID"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeHinfo      DNSTimeseriesGroupDnssece2EParamsQueryType = "HINFO"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeHip        DNSTimeseriesGroupDnssece2EParamsQueryType = "HIP"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeHTTPS      DNSTimeseriesGroupDnssece2EParamsQueryType = "HTTPS"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeIpseckey   DNSTimeseriesGroupDnssece2EParamsQueryType = "IPSECKEY"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeIsdn       DNSTimeseriesGroupDnssece2EParamsQueryType = "ISDN"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeIxfr       DNSTimeseriesGroupDnssece2EParamsQueryType = "IXFR"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeKey        DNSTimeseriesGroupDnssece2EParamsQueryType = "KEY"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeKx         DNSTimeseriesGroupDnssece2EParamsQueryType = "KX"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeL32        DNSTimeseriesGroupDnssece2EParamsQueryType = "L32"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeL64        DNSTimeseriesGroupDnssece2EParamsQueryType = "L64"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeLOC        DNSTimeseriesGroupDnssece2EParamsQueryType = "LOC"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeLp         DNSTimeseriesGroupDnssece2EParamsQueryType = "LP"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeMaila      DNSTimeseriesGroupDnssece2EParamsQueryType = "MAILA"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeMailb      DNSTimeseriesGroupDnssece2EParamsQueryType = "MAILB"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeMB         DNSTimeseriesGroupDnssece2EParamsQueryType = "MB"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeMd         DNSTimeseriesGroupDnssece2EParamsQueryType = "MD"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeMf         DNSTimeseriesGroupDnssece2EParamsQueryType = "MF"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeMg         DNSTimeseriesGroupDnssece2EParamsQueryType = "MG"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeMinfo      DNSTimeseriesGroupDnssece2EParamsQueryType = "MINFO"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeMr         DNSTimeseriesGroupDnssece2EParamsQueryType = "MR"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeMX         DNSTimeseriesGroupDnssece2EParamsQueryType = "MX"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeNAPTR      DNSTimeseriesGroupDnssece2EParamsQueryType = "NAPTR"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeNb         DNSTimeseriesGroupDnssece2EParamsQueryType = "NB"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeNbstat     DNSTimeseriesGroupDnssece2EParamsQueryType = "NBSTAT"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeNid        DNSTimeseriesGroupDnssece2EParamsQueryType = "NID"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeNimloc     DNSTimeseriesGroupDnssece2EParamsQueryType = "NIMLOC"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeNinfo      DNSTimeseriesGroupDnssece2EParamsQueryType = "NINFO"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeNS         DNSTimeseriesGroupDnssece2EParamsQueryType = "NS"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeNsap       DNSTimeseriesGroupDnssece2EParamsQueryType = "NSAP"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeNsec       DNSTimeseriesGroupDnssece2EParamsQueryType = "NSEC"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeNsec3      DNSTimeseriesGroupDnssece2EParamsQueryType = "NSEC3"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeNsec3Param DNSTimeseriesGroupDnssece2EParamsQueryType = "NSEC3PARAM"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeNull       DNSTimeseriesGroupDnssece2EParamsQueryType = "NULL"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeNxt        DNSTimeseriesGroupDnssece2EParamsQueryType = "NXT"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeOpenpgpkey DNSTimeseriesGroupDnssece2EParamsQueryType = "OPENPGPKEY"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeOpt        DNSTimeseriesGroupDnssece2EParamsQueryType = "OPT"
	DNSTimeseriesGroupDnssece2EParamsQueryTypePTR        DNSTimeseriesGroupDnssece2EParamsQueryType = "PTR"
	DNSTimeseriesGroupDnssece2EParamsQueryTypePx         DNSTimeseriesGroupDnssece2EParamsQueryType = "PX"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeRkey       DNSTimeseriesGroupDnssece2EParamsQueryType = "RKEY"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeRp         DNSTimeseriesGroupDnssece2EParamsQueryType = "RP"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeRrsig      DNSTimeseriesGroupDnssece2EParamsQueryType = "RRSIG"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeRt         DNSTimeseriesGroupDnssece2EParamsQueryType = "RT"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeSig        DNSTimeseriesGroupDnssece2EParamsQueryType = "SIG"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeSink       DNSTimeseriesGroupDnssece2EParamsQueryType = "SINK"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeSMIMEA     DNSTimeseriesGroupDnssece2EParamsQueryType = "SMIMEA"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeSOA        DNSTimeseriesGroupDnssece2EParamsQueryType = "SOA"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeSPF        DNSTimeseriesGroupDnssece2EParamsQueryType = "SPF"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeSRV        DNSTimeseriesGroupDnssece2EParamsQueryType = "SRV"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeSSHFP      DNSTimeseriesGroupDnssece2EParamsQueryType = "SSHFP"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeSVCB       DNSTimeseriesGroupDnssece2EParamsQueryType = "SVCB"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeTa         DNSTimeseriesGroupDnssece2EParamsQueryType = "TA"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeTalink     DNSTimeseriesGroupDnssece2EParamsQueryType = "TALINK"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeTkey       DNSTimeseriesGroupDnssece2EParamsQueryType = "TKEY"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeTLSA       DNSTimeseriesGroupDnssece2EParamsQueryType = "TLSA"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeTSIG       DNSTimeseriesGroupDnssece2EParamsQueryType = "TSIG"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeTXT        DNSTimeseriesGroupDnssece2EParamsQueryType = "TXT"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeUinfo      DNSTimeseriesGroupDnssece2EParamsQueryType = "UINFO"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeUID        DNSTimeseriesGroupDnssece2EParamsQueryType = "UID"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeUnspec     DNSTimeseriesGroupDnssece2EParamsQueryType = "UNSPEC"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeURI        DNSTimeseriesGroupDnssece2EParamsQueryType = "URI"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeWks        DNSTimeseriesGroupDnssece2EParamsQueryType = "WKS"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeX25        DNSTimeseriesGroupDnssece2EParamsQueryType = "X25"
	DNSTimeseriesGroupDnssece2EParamsQueryTypeZonemd     DNSTimeseriesGroupDnssece2EParamsQueryType = "ZONEMD"
)

func (r DNSTimeseriesGroupDnssece2EParamsQueryType) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupDnssece2EParamsQueryTypeA, DNSTimeseriesGroupDnssece2EParamsQueryTypeAAAA, DNSTimeseriesGroupDnssece2EParamsQueryTypeA6, DNSTimeseriesGroupDnssece2EParamsQueryTypeAfsdb, DNSTimeseriesGroupDnssece2EParamsQueryTypeAny, DNSTimeseriesGroupDnssece2EParamsQueryTypeApl, DNSTimeseriesGroupDnssece2EParamsQueryTypeAtma, DNSTimeseriesGroupDnssece2EParamsQueryTypeAXFR, DNSTimeseriesGroupDnssece2EParamsQueryTypeCAA, DNSTimeseriesGroupDnssece2EParamsQueryTypeCdnskey, DNSTimeseriesGroupDnssece2EParamsQueryTypeCds, DNSTimeseriesGroupDnssece2EParamsQueryTypeCERT, DNSTimeseriesGroupDnssece2EParamsQueryTypeCNAME, DNSTimeseriesGroupDnssece2EParamsQueryTypeCsync, DNSTimeseriesGroupDnssece2EParamsQueryTypeDhcid, DNSTimeseriesGroupDnssece2EParamsQueryTypeDlv, DNSTimeseriesGroupDnssece2EParamsQueryTypeDname, DNSTimeseriesGroupDnssece2EParamsQueryTypeDNSKEY, DNSTimeseriesGroupDnssece2EParamsQueryTypeDoa, DNSTimeseriesGroupDnssece2EParamsQueryTypeDS, DNSTimeseriesGroupDnssece2EParamsQueryTypeEid, DNSTimeseriesGroupDnssece2EParamsQueryTypeEui48, DNSTimeseriesGroupDnssece2EParamsQueryTypeEui64, DNSTimeseriesGroupDnssece2EParamsQueryTypeGpos, DNSTimeseriesGroupDnssece2EParamsQueryTypeGid, DNSTimeseriesGroupDnssece2EParamsQueryTypeHinfo, DNSTimeseriesGroupDnssece2EParamsQueryTypeHip, DNSTimeseriesGroupDnssece2EParamsQueryTypeHTTPS, DNSTimeseriesGroupDnssece2EParamsQueryTypeIpseckey, DNSTimeseriesGroupDnssece2EParamsQueryTypeIsdn, DNSTimeseriesGroupDnssece2EParamsQueryTypeIxfr, DNSTimeseriesGroupDnssece2EParamsQueryTypeKey, DNSTimeseriesGroupDnssece2EParamsQueryTypeKx, DNSTimeseriesGroupDnssece2EParamsQueryTypeL32, DNSTimeseriesGroupDnssece2EParamsQueryTypeL64, DNSTimeseriesGroupDnssece2EParamsQueryTypeLOC, DNSTimeseriesGroupDnssece2EParamsQueryTypeLp, DNSTimeseriesGroupDnssece2EParamsQueryTypeMaila, DNSTimeseriesGroupDnssece2EParamsQueryTypeMailb, DNSTimeseriesGroupDnssece2EParamsQueryTypeMB, DNSTimeseriesGroupDnssece2EParamsQueryTypeMd, DNSTimeseriesGroupDnssece2EParamsQueryTypeMf, DNSTimeseriesGroupDnssece2EParamsQueryTypeMg, DNSTimeseriesGroupDnssece2EParamsQueryTypeMinfo, DNSTimeseriesGroupDnssece2EParamsQueryTypeMr, DNSTimeseriesGroupDnssece2EParamsQueryTypeMX, DNSTimeseriesGroupDnssece2EParamsQueryTypeNAPTR, DNSTimeseriesGroupDnssece2EParamsQueryTypeNb, DNSTimeseriesGroupDnssece2EParamsQueryTypeNbstat, DNSTimeseriesGroupDnssece2EParamsQueryTypeNid, DNSTimeseriesGroupDnssece2EParamsQueryTypeNimloc, DNSTimeseriesGroupDnssece2EParamsQueryTypeNinfo, DNSTimeseriesGroupDnssece2EParamsQueryTypeNS, DNSTimeseriesGroupDnssece2EParamsQueryTypeNsap, DNSTimeseriesGroupDnssece2EParamsQueryTypeNsec, DNSTimeseriesGroupDnssece2EParamsQueryTypeNsec3, DNSTimeseriesGroupDnssece2EParamsQueryTypeNsec3Param, DNSTimeseriesGroupDnssece2EParamsQueryTypeNull, DNSTimeseriesGroupDnssece2EParamsQueryTypeNxt, DNSTimeseriesGroupDnssece2EParamsQueryTypeOpenpgpkey, DNSTimeseriesGroupDnssece2EParamsQueryTypeOpt, DNSTimeseriesGroupDnssece2EParamsQueryTypePTR, DNSTimeseriesGroupDnssece2EParamsQueryTypePx, DNSTimeseriesGroupDnssece2EParamsQueryTypeRkey, DNSTimeseriesGroupDnssece2EParamsQueryTypeRp, DNSTimeseriesGroupDnssece2EParamsQueryTypeRrsig, DNSTimeseriesGroupDnssece2EParamsQueryTypeRt, DNSTimeseriesGroupDnssece2EParamsQueryTypeSig, DNSTimeseriesGroupDnssece2EParamsQueryTypeSink, DNSTimeseriesGroupDnssece2EParamsQueryTypeSMIMEA, DNSTimeseriesGroupDnssece2EParamsQueryTypeSOA, DNSTimeseriesGroupDnssece2EParamsQueryTypeSPF, DNSTimeseriesGroupDnssece2EParamsQueryTypeSRV, DNSTimeseriesGroupDnssece2EParamsQueryTypeSSHFP, DNSTimeseriesGroupDnssece2EParamsQueryTypeSVCB, DNSTimeseriesGroupDnssece2EParamsQueryTypeTa, DNSTimeseriesGroupDnssece2EParamsQueryTypeTalink, DNSTimeseriesGroupDnssece2EParamsQueryTypeTkey, DNSTimeseriesGroupDnssece2EParamsQueryTypeTLSA, DNSTimeseriesGroupDnssece2EParamsQueryTypeTSIG, DNSTimeseriesGroupDnssece2EParamsQueryTypeTXT, DNSTimeseriesGroupDnssece2EParamsQueryTypeUinfo, DNSTimeseriesGroupDnssece2EParamsQueryTypeUID, DNSTimeseriesGroupDnssece2EParamsQueryTypeUnspec, DNSTimeseriesGroupDnssece2EParamsQueryTypeURI, DNSTimeseriesGroupDnssece2EParamsQueryTypeWks, DNSTimeseriesGroupDnssece2EParamsQueryTypeX25, DNSTimeseriesGroupDnssece2EParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSTimeseriesGroupDnssece2EParamsResponseCode string

const (
	DNSTimeseriesGroupDnssece2EParamsResponseCodeNoerror   DNSTimeseriesGroupDnssece2EParamsResponseCode = "NOERROR"
	DNSTimeseriesGroupDnssece2EParamsResponseCodeFormerr   DNSTimeseriesGroupDnssece2EParamsResponseCode = "FORMERR"
	DNSTimeseriesGroupDnssece2EParamsResponseCodeServfail  DNSTimeseriesGroupDnssece2EParamsResponseCode = "SERVFAIL"
	DNSTimeseriesGroupDnssece2EParamsResponseCodeNxdomain  DNSTimeseriesGroupDnssece2EParamsResponseCode = "NXDOMAIN"
	DNSTimeseriesGroupDnssece2EParamsResponseCodeNotimp    DNSTimeseriesGroupDnssece2EParamsResponseCode = "NOTIMP"
	DNSTimeseriesGroupDnssece2EParamsResponseCodeRefused   DNSTimeseriesGroupDnssece2EParamsResponseCode = "REFUSED"
	DNSTimeseriesGroupDnssece2EParamsResponseCodeYxdomain  DNSTimeseriesGroupDnssece2EParamsResponseCode = "YXDOMAIN"
	DNSTimeseriesGroupDnssece2EParamsResponseCodeYxrrset   DNSTimeseriesGroupDnssece2EParamsResponseCode = "YXRRSET"
	DNSTimeseriesGroupDnssece2EParamsResponseCodeNxrrset   DNSTimeseriesGroupDnssece2EParamsResponseCode = "NXRRSET"
	DNSTimeseriesGroupDnssece2EParamsResponseCodeNotauth   DNSTimeseriesGroupDnssece2EParamsResponseCode = "NOTAUTH"
	DNSTimeseriesGroupDnssece2EParamsResponseCodeNotzone   DNSTimeseriesGroupDnssece2EParamsResponseCode = "NOTZONE"
	DNSTimeseriesGroupDnssece2EParamsResponseCodeBadsig    DNSTimeseriesGroupDnssece2EParamsResponseCode = "BADSIG"
	DNSTimeseriesGroupDnssece2EParamsResponseCodeBadkey    DNSTimeseriesGroupDnssece2EParamsResponseCode = "BADKEY"
	DNSTimeseriesGroupDnssece2EParamsResponseCodeBadtime   DNSTimeseriesGroupDnssece2EParamsResponseCode = "BADTIME"
	DNSTimeseriesGroupDnssece2EParamsResponseCodeBadmode   DNSTimeseriesGroupDnssece2EParamsResponseCode = "BADMODE"
	DNSTimeseriesGroupDnssece2EParamsResponseCodeBadname   DNSTimeseriesGroupDnssece2EParamsResponseCode = "BADNAME"
	DNSTimeseriesGroupDnssece2EParamsResponseCodeBadalg    DNSTimeseriesGroupDnssece2EParamsResponseCode = "BADALG"
	DNSTimeseriesGroupDnssece2EParamsResponseCodeBadtrunc  DNSTimeseriesGroupDnssece2EParamsResponseCode = "BADTRUNC"
	DNSTimeseriesGroupDnssece2EParamsResponseCodeBadcookie DNSTimeseriesGroupDnssece2EParamsResponseCode = "BADCOOKIE"
)

func (r DNSTimeseriesGroupDnssece2EParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupDnssece2EParamsResponseCodeNoerror, DNSTimeseriesGroupDnssece2EParamsResponseCodeFormerr, DNSTimeseriesGroupDnssece2EParamsResponseCodeServfail, DNSTimeseriesGroupDnssece2EParamsResponseCodeNxdomain, DNSTimeseriesGroupDnssece2EParamsResponseCodeNotimp, DNSTimeseriesGroupDnssece2EParamsResponseCodeRefused, DNSTimeseriesGroupDnssece2EParamsResponseCodeYxdomain, DNSTimeseriesGroupDnssece2EParamsResponseCodeYxrrset, DNSTimeseriesGroupDnssece2EParamsResponseCodeNxrrset, DNSTimeseriesGroupDnssece2EParamsResponseCodeNotauth, DNSTimeseriesGroupDnssece2EParamsResponseCodeNotzone, DNSTimeseriesGroupDnssece2EParamsResponseCodeBadsig, DNSTimeseriesGroupDnssece2EParamsResponseCodeBadkey, DNSTimeseriesGroupDnssece2EParamsResponseCodeBadtime, DNSTimeseriesGroupDnssece2EParamsResponseCodeBadmode, DNSTimeseriesGroupDnssece2EParamsResponseCodeBadname, DNSTimeseriesGroupDnssece2EParamsResponseCodeBadalg, DNSTimeseriesGroupDnssece2EParamsResponseCodeBadtrunc, DNSTimeseriesGroupDnssece2EParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSTimeseriesGroupDnssece2EResponseEnvelope struct {
	Result  DNSTimeseriesGroupDnssece2EResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    dnsTimeseriesGroupDnssece2EResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupDnssece2EResponseEnvelopeJSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupDnssece2EResponseEnvelope]
type dnsTimeseriesGroupDnssece2EResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupDnssece2EResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupDnssece2EResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupIPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[DNSTimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[DNSTimeseriesGroupIPVersionParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSTimeseriesGroupIPVersionParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSTimeseriesGroupIPVersionParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSTimeseriesGroupIPVersionParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupIPVersionParams]'s query parameters as
// `url.Values`.
func (r DNSTimeseriesGroupIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupIPVersionParamsAggInterval string

const (
	DNSTimeseriesGroupIPVersionParamsAggInterval15m DNSTimeseriesGroupIPVersionParamsAggInterval = "15m"
	DNSTimeseriesGroupIPVersionParamsAggInterval1h  DNSTimeseriesGroupIPVersionParamsAggInterval = "1h"
	DNSTimeseriesGroupIPVersionParamsAggInterval1d  DNSTimeseriesGroupIPVersionParamsAggInterval = "1d"
	DNSTimeseriesGroupIPVersionParamsAggInterval1w  DNSTimeseriesGroupIPVersionParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupIPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupIPVersionParamsAggInterval15m, DNSTimeseriesGroupIPVersionParamsAggInterval1h, DNSTimeseriesGroupIPVersionParamsAggInterval1d, DNSTimeseriesGroupIPVersionParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type DNSTimeseriesGroupIPVersionParamsFormat string

const (
	DNSTimeseriesGroupIPVersionParamsFormatJson DNSTimeseriesGroupIPVersionParamsFormat = "JSON"
	DNSTimeseriesGroupIPVersionParamsFormatCsv  DNSTimeseriesGroupIPVersionParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupIPVersionParamsFormatJson, DNSTimeseriesGroupIPVersionParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSTimeseriesGroupIPVersionParamsProtocol string

const (
	DNSTimeseriesGroupIPVersionParamsProtocolUdp   DNSTimeseriesGroupIPVersionParamsProtocol = "UDP"
	DNSTimeseriesGroupIPVersionParamsProtocolTCP   DNSTimeseriesGroupIPVersionParamsProtocol = "TCP"
	DNSTimeseriesGroupIPVersionParamsProtocolHTTPS DNSTimeseriesGroupIPVersionParamsProtocol = "HTTPS"
	DNSTimeseriesGroupIPVersionParamsProtocolTLS   DNSTimeseriesGroupIPVersionParamsProtocol = "TLS"
)

func (r DNSTimeseriesGroupIPVersionParamsProtocol) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupIPVersionParamsProtocolUdp, DNSTimeseriesGroupIPVersionParamsProtocolTCP, DNSTimeseriesGroupIPVersionParamsProtocolHTTPS, DNSTimeseriesGroupIPVersionParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSTimeseriesGroupIPVersionParamsQueryType string

const (
	DNSTimeseriesGroupIPVersionParamsQueryTypeA          DNSTimeseriesGroupIPVersionParamsQueryType = "A"
	DNSTimeseriesGroupIPVersionParamsQueryTypeAAAA       DNSTimeseriesGroupIPVersionParamsQueryType = "AAAA"
	DNSTimeseriesGroupIPVersionParamsQueryTypeA6         DNSTimeseriesGroupIPVersionParamsQueryType = "A6"
	DNSTimeseriesGroupIPVersionParamsQueryTypeAfsdb      DNSTimeseriesGroupIPVersionParamsQueryType = "AFSDB"
	DNSTimeseriesGroupIPVersionParamsQueryTypeAny        DNSTimeseriesGroupIPVersionParamsQueryType = "ANY"
	DNSTimeseriesGroupIPVersionParamsQueryTypeApl        DNSTimeseriesGroupIPVersionParamsQueryType = "APL"
	DNSTimeseriesGroupIPVersionParamsQueryTypeAtma       DNSTimeseriesGroupIPVersionParamsQueryType = "ATMA"
	DNSTimeseriesGroupIPVersionParamsQueryTypeAXFR       DNSTimeseriesGroupIPVersionParamsQueryType = "AXFR"
	DNSTimeseriesGroupIPVersionParamsQueryTypeCAA        DNSTimeseriesGroupIPVersionParamsQueryType = "CAA"
	DNSTimeseriesGroupIPVersionParamsQueryTypeCdnskey    DNSTimeseriesGroupIPVersionParamsQueryType = "CDNSKEY"
	DNSTimeseriesGroupIPVersionParamsQueryTypeCds        DNSTimeseriesGroupIPVersionParamsQueryType = "CDS"
	DNSTimeseriesGroupIPVersionParamsQueryTypeCERT       DNSTimeseriesGroupIPVersionParamsQueryType = "CERT"
	DNSTimeseriesGroupIPVersionParamsQueryTypeCNAME      DNSTimeseriesGroupIPVersionParamsQueryType = "CNAME"
	DNSTimeseriesGroupIPVersionParamsQueryTypeCsync      DNSTimeseriesGroupIPVersionParamsQueryType = "CSYNC"
	DNSTimeseriesGroupIPVersionParamsQueryTypeDhcid      DNSTimeseriesGroupIPVersionParamsQueryType = "DHCID"
	DNSTimeseriesGroupIPVersionParamsQueryTypeDlv        DNSTimeseriesGroupIPVersionParamsQueryType = "DLV"
	DNSTimeseriesGroupIPVersionParamsQueryTypeDname      DNSTimeseriesGroupIPVersionParamsQueryType = "DNAME"
	DNSTimeseriesGroupIPVersionParamsQueryTypeDNSKEY     DNSTimeseriesGroupIPVersionParamsQueryType = "DNSKEY"
	DNSTimeseriesGroupIPVersionParamsQueryTypeDoa        DNSTimeseriesGroupIPVersionParamsQueryType = "DOA"
	DNSTimeseriesGroupIPVersionParamsQueryTypeDS         DNSTimeseriesGroupIPVersionParamsQueryType = "DS"
	DNSTimeseriesGroupIPVersionParamsQueryTypeEid        DNSTimeseriesGroupIPVersionParamsQueryType = "EID"
	DNSTimeseriesGroupIPVersionParamsQueryTypeEui48      DNSTimeseriesGroupIPVersionParamsQueryType = "EUI48"
	DNSTimeseriesGroupIPVersionParamsQueryTypeEui64      DNSTimeseriesGroupIPVersionParamsQueryType = "EUI64"
	DNSTimeseriesGroupIPVersionParamsQueryTypeGpos       DNSTimeseriesGroupIPVersionParamsQueryType = "GPOS"
	DNSTimeseriesGroupIPVersionParamsQueryTypeGid        DNSTimeseriesGroupIPVersionParamsQueryType = "GID"
	DNSTimeseriesGroupIPVersionParamsQueryTypeHinfo      DNSTimeseriesGroupIPVersionParamsQueryType = "HINFO"
	DNSTimeseriesGroupIPVersionParamsQueryTypeHip        DNSTimeseriesGroupIPVersionParamsQueryType = "HIP"
	DNSTimeseriesGroupIPVersionParamsQueryTypeHTTPS      DNSTimeseriesGroupIPVersionParamsQueryType = "HTTPS"
	DNSTimeseriesGroupIPVersionParamsQueryTypeIpseckey   DNSTimeseriesGroupIPVersionParamsQueryType = "IPSECKEY"
	DNSTimeseriesGroupIPVersionParamsQueryTypeIsdn       DNSTimeseriesGroupIPVersionParamsQueryType = "ISDN"
	DNSTimeseriesGroupIPVersionParamsQueryTypeIxfr       DNSTimeseriesGroupIPVersionParamsQueryType = "IXFR"
	DNSTimeseriesGroupIPVersionParamsQueryTypeKey        DNSTimeseriesGroupIPVersionParamsQueryType = "KEY"
	DNSTimeseriesGroupIPVersionParamsQueryTypeKx         DNSTimeseriesGroupIPVersionParamsQueryType = "KX"
	DNSTimeseriesGroupIPVersionParamsQueryTypeL32        DNSTimeseriesGroupIPVersionParamsQueryType = "L32"
	DNSTimeseriesGroupIPVersionParamsQueryTypeL64        DNSTimeseriesGroupIPVersionParamsQueryType = "L64"
	DNSTimeseriesGroupIPVersionParamsQueryTypeLOC        DNSTimeseriesGroupIPVersionParamsQueryType = "LOC"
	DNSTimeseriesGroupIPVersionParamsQueryTypeLp         DNSTimeseriesGroupIPVersionParamsQueryType = "LP"
	DNSTimeseriesGroupIPVersionParamsQueryTypeMaila      DNSTimeseriesGroupIPVersionParamsQueryType = "MAILA"
	DNSTimeseriesGroupIPVersionParamsQueryTypeMailb      DNSTimeseriesGroupIPVersionParamsQueryType = "MAILB"
	DNSTimeseriesGroupIPVersionParamsQueryTypeMB         DNSTimeseriesGroupIPVersionParamsQueryType = "MB"
	DNSTimeseriesGroupIPVersionParamsQueryTypeMd         DNSTimeseriesGroupIPVersionParamsQueryType = "MD"
	DNSTimeseriesGroupIPVersionParamsQueryTypeMf         DNSTimeseriesGroupIPVersionParamsQueryType = "MF"
	DNSTimeseriesGroupIPVersionParamsQueryTypeMg         DNSTimeseriesGroupIPVersionParamsQueryType = "MG"
	DNSTimeseriesGroupIPVersionParamsQueryTypeMinfo      DNSTimeseriesGroupIPVersionParamsQueryType = "MINFO"
	DNSTimeseriesGroupIPVersionParamsQueryTypeMr         DNSTimeseriesGroupIPVersionParamsQueryType = "MR"
	DNSTimeseriesGroupIPVersionParamsQueryTypeMX         DNSTimeseriesGroupIPVersionParamsQueryType = "MX"
	DNSTimeseriesGroupIPVersionParamsQueryTypeNAPTR      DNSTimeseriesGroupIPVersionParamsQueryType = "NAPTR"
	DNSTimeseriesGroupIPVersionParamsQueryTypeNb         DNSTimeseriesGroupIPVersionParamsQueryType = "NB"
	DNSTimeseriesGroupIPVersionParamsQueryTypeNbstat     DNSTimeseriesGroupIPVersionParamsQueryType = "NBSTAT"
	DNSTimeseriesGroupIPVersionParamsQueryTypeNid        DNSTimeseriesGroupIPVersionParamsQueryType = "NID"
	DNSTimeseriesGroupIPVersionParamsQueryTypeNimloc     DNSTimeseriesGroupIPVersionParamsQueryType = "NIMLOC"
	DNSTimeseriesGroupIPVersionParamsQueryTypeNinfo      DNSTimeseriesGroupIPVersionParamsQueryType = "NINFO"
	DNSTimeseriesGroupIPVersionParamsQueryTypeNS         DNSTimeseriesGroupIPVersionParamsQueryType = "NS"
	DNSTimeseriesGroupIPVersionParamsQueryTypeNsap       DNSTimeseriesGroupIPVersionParamsQueryType = "NSAP"
	DNSTimeseriesGroupIPVersionParamsQueryTypeNsec       DNSTimeseriesGroupIPVersionParamsQueryType = "NSEC"
	DNSTimeseriesGroupIPVersionParamsQueryTypeNsec3      DNSTimeseriesGroupIPVersionParamsQueryType = "NSEC3"
	DNSTimeseriesGroupIPVersionParamsQueryTypeNsec3Param DNSTimeseriesGroupIPVersionParamsQueryType = "NSEC3PARAM"
	DNSTimeseriesGroupIPVersionParamsQueryTypeNull       DNSTimeseriesGroupIPVersionParamsQueryType = "NULL"
	DNSTimeseriesGroupIPVersionParamsQueryTypeNxt        DNSTimeseriesGroupIPVersionParamsQueryType = "NXT"
	DNSTimeseriesGroupIPVersionParamsQueryTypeOpenpgpkey DNSTimeseriesGroupIPVersionParamsQueryType = "OPENPGPKEY"
	DNSTimeseriesGroupIPVersionParamsQueryTypeOpt        DNSTimeseriesGroupIPVersionParamsQueryType = "OPT"
	DNSTimeseriesGroupIPVersionParamsQueryTypePTR        DNSTimeseriesGroupIPVersionParamsQueryType = "PTR"
	DNSTimeseriesGroupIPVersionParamsQueryTypePx         DNSTimeseriesGroupIPVersionParamsQueryType = "PX"
	DNSTimeseriesGroupIPVersionParamsQueryTypeRkey       DNSTimeseriesGroupIPVersionParamsQueryType = "RKEY"
	DNSTimeseriesGroupIPVersionParamsQueryTypeRp         DNSTimeseriesGroupIPVersionParamsQueryType = "RP"
	DNSTimeseriesGroupIPVersionParamsQueryTypeRrsig      DNSTimeseriesGroupIPVersionParamsQueryType = "RRSIG"
	DNSTimeseriesGroupIPVersionParamsQueryTypeRt         DNSTimeseriesGroupIPVersionParamsQueryType = "RT"
	DNSTimeseriesGroupIPVersionParamsQueryTypeSig        DNSTimeseriesGroupIPVersionParamsQueryType = "SIG"
	DNSTimeseriesGroupIPVersionParamsQueryTypeSink       DNSTimeseriesGroupIPVersionParamsQueryType = "SINK"
	DNSTimeseriesGroupIPVersionParamsQueryTypeSMIMEA     DNSTimeseriesGroupIPVersionParamsQueryType = "SMIMEA"
	DNSTimeseriesGroupIPVersionParamsQueryTypeSOA        DNSTimeseriesGroupIPVersionParamsQueryType = "SOA"
	DNSTimeseriesGroupIPVersionParamsQueryTypeSPF        DNSTimeseriesGroupIPVersionParamsQueryType = "SPF"
	DNSTimeseriesGroupIPVersionParamsQueryTypeSRV        DNSTimeseriesGroupIPVersionParamsQueryType = "SRV"
	DNSTimeseriesGroupIPVersionParamsQueryTypeSSHFP      DNSTimeseriesGroupIPVersionParamsQueryType = "SSHFP"
	DNSTimeseriesGroupIPVersionParamsQueryTypeSVCB       DNSTimeseriesGroupIPVersionParamsQueryType = "SVCB"
	DNSTimeseriesGroupIPVersionParamsQueryTypeTa         DNSTimeseriesGroupIPVersionParamsQueryType = "TA"
	DNSTimeseriesGroupIPVersionParamsQueryTypeTalink     DNSTimeseriesGroupIPVersionParamsQueryType = "TALINK"
	DNSTimeseriesGroupIPVersionParamsQueryTypeTkey       DNSTimeseriesGroupIPVersionParamsQueryType = "TKEY"
	DNSTimeseriesGroupIPVersionParamsQueryTypeTLSA       DNSTimeseriesGroupIPVersionParamsQueryType = "TLSA"
	DNSTimeseriesGroupIPVersionParamsQueryTypeTSIG       DNSTimeseriesGroupIPVersionParamsQueryType = "TSIG"
	DNSTimeseriesGroupIPVersionParamsQueryTypeTXT        DNSTimeseriesGroupIPVersionParamsQueryType = "TXT"
	DNSTimeseriesGroupIPVersionParamsQueryTypeUinfo      DNSTimeseriesGroupIPVersionParamsQueryType = "UINFO"
	DNSTimeseriesGroupIPVersionParamsQueryTypeUID        DNSTimeseriesGroupIPVersionParamsQueryType = "UID"
	DNSTimeseriesGroupIPVersionParamsQueryTypeUnspec     DNSTimeseriesGroupIPVersionParamsQueryType = "UNSPEC"
	DNSTimeseriesGroupIPVersionParamsQueryTypeURI        DNSTimeseriesGroupIPVersionParamsQueryType = "URI"
	DNSTimeseriesGroupIPVersionParamsQueryTypeWks        DNSTimeseriesGroupIPVersionParamsQueryType = "WKS"
	DNSTimeseriesGroupIPVersionParamsQueryTypeX25        DNSTimeseriesGroupIPVersionParamsQueryType = "X25"
	DNSTimeseriesGroupIPVersionParamsQueryTypeZonemd     DNSTimeseriesGroupIPVersionParamsQueryType = "ZONEMD"
)

func (r DNSTimeseriesGroupIPVersionParamsQueryType) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupIPVersionParamsQueryTypeA, DNSTimeseriesGroupIPVersionParamsQueryTypeAAAA, DNSTimeseriesGroupIPVersionParamsQueryTypeA6, DNSTimeseriesGroupIPVersionParamsQueryTypeAfsdb, DNSTimeseriesGroupIPVersionParamsQueryTypeAny, DNSTimeseriesGroupIPVersionParamsQueryTypeApl, DNSTimeseriesGroupIPVersionParamsQueryTypeAtma, DNSTimeseriesGroupIPVersionParamsQueryTypeAXFR, DNSTimeseriesGroupIPVersionParamsQueryTypeCAA, DNSTimeseriesGroupIPVersionParamsQueryTypeCdnskey, DNSTimeseriesGroupIPVersionParamsQueryTypeCds, DNSTimeseriesGroupIPVersionParamsQueryTypeCERT, DNSTimeseriesGroupIPVersionParamsQueryTypeCNAME, DNSTimeseriesGroupIPVersionParamsQueryTypeCsync, DNSTimeseriesGroupIPVersionParamsQueryTypeDhcid, DNSTimeseriesGroupIPVersionParamsQueryTypeDlv, DNSTimeseriesGroupIPVersionParamsQueryTypeDname, DNSTimeseriesGroupIPVersionParamsQueryTypeDNSKEY, DNSTimeseriesGroupIPVersionParamsQueryTypeDoa, DNSTimeseriesGroupIPVersionParamsQueryTypeDS, DNSTimeseriesGroupIPVersionParamsQueryTypeEid, DNSTimeseriesGroupIPVersionParamsQueryTypeEui48, DNSTimeseriesGroupIPVersionParamsQueryTypeEui64, DNSTimeseriesGroupIPVersionParamsQueryTypeGpos, DNSTimeseriesGroupIPVersionParamsQueryTypeGid, DNSTimeseriesGroupIPVersionParamsQueryTypeHinfo, DNSTimeseriesGroupIPVersionParamsQueryTypeHip, DNSTimeseriesGroupIPVersionParamsQueryTypeHTTPS, DNSTimeseriesGroupIPVersionParamsQueryTypeIpseckey, DNSTimeseriesGroupIPVersionParamsQueryTypeIsdn, DNSTimeseriesGroupIPVersionParamsQueryTypeIxfr, DNSTimeseriesGroupIPVersionParamsQueryTypeKey, DNSTimeseriesGroupIPVersionParamsQueryTypeKx, DNSTimeseriesGroupIPVersionParamsQueryTypeL32, DNSTimeseriesGroupIPVersionParamsQueryTypeL64, DNSTimeseriesGroupIPVersionParamsQueryTypeLOC, DNSTimeseriesGroupIPVersionParamsQueryTypeLp, DNSTimeseriesGroupIPVersionParamsQueryTypeMaila, DNSTimeseriesGroupIPVersionParamsQueryTypeMailb, DNSTimeseriesGroupIPVersionParamsQueryTypeMB, DNSTimeseriesGroupIPVersionParamsQueryTypeMd, DNSTimeseriesGroupIPVersionParamsQueryTypeMf, DNSTimeseriesGroupIPVersionParamsQueryTypeMg, DNSTimeseriesGroupIPVersionParamsQueryTypeMinfo, DNSTimeseriesGroupIPVersionParamsQueryTypeMr, DNSTimeseriesGroupIPVersionParamsQueryTypeMX, DNSTimeseriesGroupIPVersionParamsQueryTypeNAPTR, DNSTimeseriesGroupIPVersionParamsQueryTypeNb, DNSTimeseriesGroupIPVersionParamsQueryTypeNbstat, DNSTimeseriesGroupIPVersionParamsQueryTypeNid, DNSTimeseriesGroupIPVersionParamsQueryTypeNimloc, DNSTimeseriesGroupIPVersionParamsQueryTypeNinfo, DNSTimeseriesGroupIPVersionParamsQueryTypeNS, DNSTimeseriesGroupIPVersionParamsQueryTypeNsap, DNSTimeseriesGroupIPVersionParamsQueryTypeNsec, DNSTimeseriesGroupIPVersionParamsQueryTypeNsec3, DNSTimeseriesGroupIPVersionParamsQueryTypeNsec3Param, DNSTimeseriesGroupIPVersionParamsQueryTypeNull, DNSTimeseriesGroupIPVersionParamsQueryTypeNxt, DNSTimeseriesGroupIPVersionParamsQueryTypeOpenpgpkey, DNSTimeseriesGroupIPVersionParamsQueryTypeOpt, DNSTimeseriesGroupIPVersionParamsQueryTypePTR, DNSTimeseriesGroupIPVersionParamsQueryTypePx, DNSTimeseriesGroupIPVersionParamsQueryTypeRkey, DNSTimeseriesGroupIPVersionParamsQueryTypeRp, DNSTimeseriesGroupIPVersionParamsQueryTypeRrsig, DNSTimeseriesGroupIPVersionParamsQueryTypeRt, DNSTimeseriesGroupIPVersionParamsQueryTypeSig, DNSTimeseriesGroupIPVersionParamsQueryTypeSink, DNSTimeseriesGroupIPVersionParamsQueryTypeSMIMEA, DNSTimeseriesGroupIPVersionParamsQueryTypeSOA, DNSTimeseriesGroupIPVersionParamsQueryTypeSPF, DNSTimeseriesGroupIPVersionParamsQueryTypeSRV, DNSTimeseriesGroupIPVersionParamsQueryTypeSSHFP, DNSTimeseriesGroupIPVersionParamsQueryTypeSVCB, DNSTimeseriesGroupIPVersionParamsQueryTypeTa, DNSTimeseriesGroupIPVersionParamsQueryTypeTalink, DNSTimeseriesGroupIPVersionParamsQueryTypeTkey, DNSTimeseriesGroupIPVersionParamsQueryTypeTLSA, DNSTimeseriesGroupIPVersionParamsQueryTypeTSIG, DNSTimeseriesGroupIPVersionParamsQueryTypeTXT, DNSTimeseriesGroupIPVersionParamsQueryTypeUinfo, DNSTimeseriesGroupIPVersionParamsQueryTypeUID, DNSTimeseriesGroupIPVersionParamsQueryTypeUnspec, DNSTimeseriesGroupIPVersionParamsQueryTypeURI, DNSTimeseriesGroupIPVersionParamsQueryTypeWks, DNSTimeseriesGroupIPVersionParamsQueryTypeX25, DNSTimeseriesGroupIPVersionParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSTimeseriesGroupIPVersionParamsResponseCode string

const (
	DNSTimeseriesGroupIPVersionParamsResponseCodeNoerror   DNSTimeseriesGroupIPVersionParamsResponseCode = "NOERROR"
	DNSTimeseriesGroupIPVersionParamsResponseCodeFormerr   DNSTimeseriesGroupIPVersionParamsResponseCode = "FORMERR"
	DNSTimeseriesGroupIPVersionParamsResponseCodeServfail  DNSTimeseriesGroupIPVersionParamsResponseCode = "SERVFAIL"
	DNSTimeseriesGroupIPVersionParamsResponseCodeNxdomain  DNSTimeseriesGroupIPVersionParamsResponseCode = "NXDOMAIN"
	DNSTimeseriesGroupIPVersionParamsResponseCodeNotimp    DNSTimeseriesGroupIPVersionParamsResponseCode = "NOTIMP"
	DNSTimeseriesGroupIPVersionParamsResponseCodeRefused   DNSTimeseriesGroupIPVersionParamsResponseCode = "REFUSED"
	DNSTimeseriesGroupIPVersionParamsResponseCodeYxdomain  DNSTimeseriesGroupIPVersionParamsResponseCode = "YXDOMAIN"
	DNSTimeseriesGroupIPVersionParamsResponseCodeYxrrset   DNSTimeseriesGroupIPVersionParamsResponseCode = "YXRRSET"
	DNSTimeseriesGroupIPVersionParamsResponseCodeNxrrset   DNSTimeseriesGroupIPVersionParamsResponseCode = "NXRRSET"
	DNSTimeseriesGroupIPVersionParamsResponseCodeNotauth   DNSTimeseriesGroupIPVersionParamsResponseCode = "NOTAUTH"
	DNSTimeseriesGroupIPVersionParamsResponseCodeNotzone   DNSTimeseriesGroupIPVersionParamsResponseCode = "NOTZONE"
	DNSTimeseriesGroupIPVersionParamsResponseCodeBadsig    DNSTimeseriesGroupIPVersionParamsResponseCode = "BADSIG"
	DNSTimeseriesGroupIPVersionParamsResponseCodeBadkey    DNSTimeseriesGroupIPVersionParamsResponseCode = "BADKEY"
	DNSTimeseriesGroupIPVersionParamsResponseCodeBadtime   DNSTimeseriesGroupIPVersionParamsResponseCode = "BADTIME"
	DNSTimeseriesGroupIPVersionParamsResponseCodeBadmode   DNSTimeseriesGroupIPVersionParamsResponseCode = "BADMODE"
	DNSTimeseriesGroupIPVersionParamsResponseCodeBadname   DNSTimeseriesGroupIPVersionParamsResponseCode = "BADNAME"
	DNSTimeseriesGroupIPVersionParamsResponseCodeBadalg    DNSTimeseriesGroupIPVersionParamsResponseCode = "BADALG"
	DNSTimeseriesGroupIPVersionParamsResponseCodeBadtrunc  DNSTimeseriesGroupIPVersionParamsResponseCode = "BADTRUNC"
	DNSTimeseriesGroupIPVersionParamsResponseCodeBadcookie DNSTimeseriesGroupIPVersionParamsResponseCode = "BADCOOKIE"
)

func (r DNSTimeseriesGroupIPVersionParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupIPVersionParamsResponseCodeNoerror, DNSTimeseriesGroupIPVersionParamsResponseCodeFormerr, DNSTimeseriesGroupIPVersionParamsResponseCodeServfail, DNSTimeseriesGroupIPVersionParamsResponseCodeNxdomain, DNSTimeseriesGroupIPVersionParamsResponseCodeNotimp, DNSTimeseriesGroupIPVersionParamsResponseCodeRefused, DNSTimeseriesGroupIPVersionParamsResponseCodeYxdomain, DNSTimeseriesGroupIPVersionParamsResponseCodeYxrrset, DNSTimeseriesGroupIPVersionParamsResponseCodeNxrrset, DNSTimeseriesGroupIPVersionParamsResponseCodeNotauth, DNSTimeseriesGroupIPVersionParamsResponseCodeNotzone, DNSTimeseriesGroupIPVersionParamsResponseCodeBadsig, DNSTimeseriesGroupIPVersionParamsResponseCodeBadkey, DNSTimeseriesGroupIPVersionParamsResponseCodeBadtime, DNSTimeseriesGroupIPVersionParamsResponseCodeBadmode, DNSTimeseriesGroupIPVersionParamsResponseCodeBadname, DNSTimeseriesGroupIPVersionParamsResponseCodeBadalg, DNSTimeseriesGroupIPVersionParamsResponseCodeBadtrunc, DNSTimeseriesGroupIPVersionParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSTimeseriesGroupIPVersionResponseEnvelope struct {
	Result  DNSTimeseriesGroupIPVersionResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    dnsTimeseriesGroupIPVersionResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupIPVersionResponseEnvelopeJSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupIPVersionResponseEnvelope]
type dnsTimeseriesGroupIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupMatchingAnswerParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[DNSTimeseriesGroupMatchingAnswerParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[DNSTimeseriesGroupMatchingAnswerParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSTimeseriesGroupMatchingAnswerParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSTimeseriesGroupMatchingAnswerParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSTimeseriesGroupMatchingAnswerParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupMatchingAnswerParams]'s query parameters
// as `url.Values`.
func (r DNSTimeseriesGroupMatchingAnswerParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupMatchingAnswerParamsAggInterval string

const (
	DNSTimeseriesGroupMatchingAnswerParamsAggInterval15m DNSTimeseriesGroupMatchingAnswerParamsAggInterval = "15m"
	DNSTimeseriesGroupMatchingAnswerParamsAggInterval1h  DNSTimeseriesGroupMatchingAnswerParamsAggInterval = "1h"
	DNSTimeseriesGroupMatchingAnswerParamsAggInterval1d  DNSTimeseriesGroupMatchingAnswerParamsAggInterval = "1d"
	DNSTimeseriesGroupMatchingAnswerParamsAggInterval1w  DNSTimeseriesGroupMatchingAnswerParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupMatchingAnswerParamsAggInterval) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupMatchingAnswerParamsAggInterval15m, DNSTimeseriesGroupMatchingAnswerParamsAggInterval1h, DNSTimeseriesGroupMatchingAnswerParamsAggInterval1d, DNSTimeseriesGroupMatchingAnswerParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type DNSTimeseriesGroupMatchingAnswerParamsFormat string

const (
	DNSTimeseriesGroupMatchingAnswerParamsFormatJson DNSTimeseriesGroupMatchingAnswerParamsFormat = "JSON"
	DNSTimeseriesGroupMatchingAnswerParamsFormatCsv  DNSTimeseriesGroupMatchingAnswerParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupMatchingAnswerParamsFormat) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupMatchingAnswerParamsFormatJson, DNSTimeseriesGroupMatchingAnswerParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSTimeseriesGroupMatchingAnswerParamsProtocol string

const (
	DNSTimeseriesGroupMatchingAnswerParamsProtocolUdp   DNSTimeseriesGroupMatchingAnswerParamsProtocol = "UDP"
	DNSTimeseriesGroupMatchingAnswerParamsProtocolTCP   DNSTimeseriesGroupMatchingAnswerParamsProtocol = "TCP"
	DNSTimeseriesGroupMatchingAnswerParamsProtocolHTTPS DNSTimeseriesGroupMatchingAnswerParamsProtocol = "HTTPS"
	DNSTimeseriesGroupMatchingAnswerParamsProtocolTLS   DNSTimeseriesGroupMatchingAnswerParamsProtocol = "TLS"
)

func (r DNSTimeseriesGroupMatchingAnswerParamsProtocol) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupMatchingAnswerParamsProtocolUdp, DNSTimeseriesGroupMatchingAnswerParamsProtocolTCP, DNSTimeseriesGroupMatchingAnswerParamsProtocolHTTPS, DNSTimeseriesGroupMatchingAnswerParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSTimeseriesGroupMatchingAnswerParamsQueryType string

const (
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeA          DNSTimeseriesGroupMatchingAnswerParamsQueryType = "A"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeAAAA       DNSTimeseriesGroupMatchingAnswerParamsQueryType = "AAAA"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeA6         DNSTimeseriesGroupMatchingAnswerParamsQueryType = "A6"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeAfsdb      DNSTimeseriesGroupMatchingAnswerParamsQueryType = "AFSDB"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeAny        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "ANY"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeApl        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "APL"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeAtma       DNSTimeseriesGroupMatchingAnswerParamsQueryType = "ATMA"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeAXFR       DNSTimeseriesGroupMatchingAnswerParamsQueryType = "AXFR"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeCAA        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "CAA"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeCdnskey    DNSTimeseriesGroupMatchingAnswerParamsQueryType = "CDNSKEY"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeCds        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "CDS"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeCERT       DNSTimeseriesGroupMatchingAnswerParamsQueryType = "CERT"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeCNAME      DNSTimeseriesGroupMatchingAnswerParamsQueryType = "CNAME"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeCsync      DNSTimeseriesGroupMatchingAnswerParamsQueryType = "CSYNC"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeDhcid      DNSTimeseriesGroupMatchingAnswerParamsQueryType = "DHCID"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeDlv        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "DLV"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeDname      DNSTimeseriesGroupMatchingAnswerParamsQueryType = "DNAME"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeDNSKEY     DNSTimeseriesGroupMatchingAnswerParamsQueryType = "DNSKEY"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeDoa        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "DOA"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeDS         DNSTimeseriesGroupMatchingAnswerParamsQueryType = "DS"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeEid        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "EID"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeEui48      DNSTimeseriesGroupMatchingAnswerParamsQueryType = "EUI48"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeEui64      DNSTimeseriesGroupMatchingAnswerParamsQueryType = "EUI64"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeGpos       DNSTimeseriesGroupMatchingAnswerParamsQueryType = "GPOS"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeGid        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "GID"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeHinfo      DNSTimeseriesGroupMatchingAnswerParamsQueryType = "HINFO"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeHip        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "HIP"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeHTTPS      DNSTimeseriesGroupMatchingAnswerParamsQueryType = "HTTPS"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeIpseckey   DNSTimeseriesGroupMatchingAnswerParamsQueryType = "IPSECKEY"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeIsdn       DNSTimeseriesGroupMatchingAnswerParamsQueryType = "ISDN"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeIxfr       DNSTimeseriesGroupMatchingAnswerParamsQueryType = "IXFR"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeKey        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "KEY"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeKx         DNSTimeseriesGroupMatchingAnswerParamsQueryType = "KX"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeL32        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "L32"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeL64        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "L64"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeLOC        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "LOC"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeLp         DNSTimeseriesGroupMatchingAnswerParamsQueryType = "LP"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeMaila      DNSTimeseriesGroupMatchingAnswerParamsQueryType = "MAILA"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeMailb      DNSTimeseriesGroupMatchingAnswerParamsQueryType = "MAILB"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeMB         DNSTimeseriesGroupMatchingAnswerParamsQueryType = "MB"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeMd         DNSTimeseriesGroupMatchingAnswerParamsQueryType = "MD"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeMf         DNSTimeseriesGroupMatchingAnswerParamsQueryType = "MF"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeMg         DNSTimeseriesGroupMatchingAnswerParamsQueryType = "MG"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeMinfo      DNSTimeseriesGroupMatchingAnswerParamsQueryType = "MINFO"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeMr         DNSTimeseriesGroupMatchingAnswerParamsQueryType = "MR"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeMX         DNSTimeseriesGroupMatchingAnswerParamsQueryType = "MX"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNAPTR      DNSTimeseriesGroupMatchingAnswerParamsQueryType = "NAPTR"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNb         DNSTimeseriesGroupMatchingAnswerParamsQueryType = "NB"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNbstat     DNSTimeseriesGroupMatchingAnswerParamsQueryType = "NBSTAT"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNid        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "NID"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNimloc     DNSTimeseriesGroupMatchingAnswerParamsQueryType = "NIMLOC"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNinfo      DNSTimeseriesGroupMatchingAnswerParamsQueryType = "NINFO"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNS         DNSTimeseriesGroupMatchingAnswerParamsQueryType = "NS"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNsap       DNSTimeseriesGroupMatchingAnswerParamsQueryType = "NSAP"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNsec       DNSTimeseriesGroupMatchingAnswerParamsQueryType = "NSEC"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNsec3      DNSTimeseriesGroupMatchingAnswerParamsQueryType = "NSEC3"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNsec3Param DNSTimeseriesGroupMatchingAnswerParamsQueryType = "NSEC3PARAM"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNull       DNSTimeseriesGroupMatchingAnswerParamsQueryType = "NULL"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNxt        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "NXT"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeOpenpgpkey DNSTimeseriesGroupMatchingAnswerParamsQueryType = "OPENPGPKEY"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeOpt        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "OPT"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypePTR        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "PTR"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypePx         DNSTimeseriesGroupMatchingAnswerParamsQueryType = "PX"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeRkey       DNSTimeseriesGroupMatchingAnswerParamsQueryType = "RKEY"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeRp         DNSTimeseriesGroupMatchingAnswerParamsQueryType = "RP"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeRrsig      DNSTimeseriesGroupMatchingAnswerParamsQueryType = "RRSIG"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeRt         DNSTimeseriesGroupMatchingAnswerParamsQueryType = "RT"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeSig        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "SIG"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeSink       DNSTimeseriesGroupMatchingAnswerParamsQueryType = "SINK"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeSMIMEA     DNSTimeseriesGroupMatchingAnswerParamsQueryType = "SMIMEA"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeSOA        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "SOA"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeSPF        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "SPF"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeSRV        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "SRV"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeSSHFP      DNSTimeseriesGroupMatchingAnswerParamsQueryType = "SSHFP"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeSVCB       DNSTimeseriesGroupMatchingAnswerParamsQueryType = "SVCB"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeTa         DNSTimeseriesGroupMatchingAnswerParamsQueryType = "TA"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeTalink     DNSTimeseriesGroupMatchingAnswerParamsQueryType = "TALINK"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeTkey       DNSTimeseriesGroupMatchingAnswerParamsQueryType = "TKEY"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeTLSA       DNSTimeseriesGroupMatchingAnswerParamsQueryType = "TLSA"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeTSIG       DNSTimeseriesGroupMatchingAnswerParamsQueryType = "TSIG"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeTXT        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "TXT"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeUinfo      DNSTimeseriesGroupMatchingAnswerParamsQueryType = "UINFO"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeUID        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "UID"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeUnspec     DNSTimeseriesGroupMatchingAnswerParamsQueryType = "UNSPEC"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeURI        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "URI"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeWks        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "WKS"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeX25        DNSTimeseriesGroupMatchingAnswerParamsQueryType = "X25"
	DNSTimeseriesGroupMatchingAnswerParamsQueryTypeZonemd     DNSTimeseriesGroupMatchingAnswerParamsQueryType = "ZONEMD"
)

func (r DNSTimeseriesGroupMatchingAnswerParamsQueryType) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupMatchingAnswerParamsQueryTypeA, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeAAAA, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeA6, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeAfsdb, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeAny, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeApl, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeAtma, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeAXFR, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeCAA, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeCdnskey, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeCds, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeCERT, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeCNAME, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeCsync, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeDhcid, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeDlv, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeDname, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeDNSKEY, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeDoa, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeDS, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeEid, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeEui48, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeEui64, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeGpos, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeGid, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeHinfo, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeHip, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeHTTPS, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeIpseckey, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeIsdn, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeIxfr, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeKey, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeKx, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeL32, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeL64, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeLOC, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeLp, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeMaila, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeMailb, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeMB, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeMd, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeMf, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeMg, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeMinfo, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeMr, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeMX, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNAPTR, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNb, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNbstat, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNid, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNimloc, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNinfo, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNS, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNsap, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNsec, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNsec3, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNsec3Param, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNull, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeNxt, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeOpenpgpkey, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeOpt, DNSTimeseriesGroupMatchingAnswerParamsQueryTypePTR, DNSTimeseriesGroupMatchingAnswerParamsQueryTypePx, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeRkey, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeRp, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeRrsig, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeRt, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeSig, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeSink, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeSMIMEA, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeSOA, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeSPF, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeSRV, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeSSHFP, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeSVCB, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeTa, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeTalink, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeTkey, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeTLSA, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeTSIG, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeTXT, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeUinfo, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeUID, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeUnspec, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeURI, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeWks, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeX25, DNSTimeseriesGroupMatchingAnswerParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSTimeseriesGroupMatchingAnswerParamsResponseCode string

const (
	DNSTimeseriesGroupMatchingAnswerParamsResponseCodeNoerror   DNSTimeseriesGroupMatchingAnswerParamsResponseCode = "NOERROR"
	DNSTimeseriesGroupMatchingAnswerParamsResponseCodeFormerr   DNSTimeseriesGroupMatchingAnswerParamsResponseCode = "FORMERR"
	DNSTimeseriesGroupMatchingAnswerParamsResponseCodeServfail  DNSTimeseriesGroupMatchingAnswerParamsResponseCode = "SERVFAIL"
	DNSTimeseriesGroupMatchingAnswerParamsResponseCodeNxdomain  DNSTimeseriesGroupMatchingAnswerParamsResponseCode = "NXDOMAIN"
	DNSTimeseriesGroupMatchingAnswerParamsResponseCodeNotimp    DNSTimeseriesGroupMatchingAnswerParamsResponseCode = "NOTIMP"
	DNSTimeseriesGroupMatchingAnswerParamsResponseCodeRefused   DNSTimeseriesGroupMatchingAnswerParamsResponseCode = "REFUSED"
	DNSTimeseriesGroupMatchingAnswerParamsResponseCodeYxdomain  DNSTimeseriesGroupMatchingAnswerParamsResponseCode = "YXDOMAIN"
	DNSTimeseriesGroupMatchingAnswerParamsResponseCodeYxrrset   DNSTimeseriesGroupMatchingAnswerParamsResponseCode = "YXRRSET"
	DNSTimeseriesGroupMatchingAnswerParamsResponseCodeNxrrset   DNSTimeseriesGroupMatchingAnswerParamsResponseCode = "NXRRSET"
	DNSTimeseriesGroupMatchingAnswerParamsResponseCodeNotauth   DNSTimeseriesGroupMatchingAnswerParamsResponseCode = "NOTAUTH"
	DNSTimeseriesGroupMatchingAnswerParamsResponseCodeNotzone   DNSTimeseriesGroupMatchingAnswerParamsResponseCode = "NOTZONE"
	DNSTimeseriesGroupMatchingAnswerParamsResponseCodeBadsig    DNSTimeseriesGroupMatchingAnswerParamsResponseCode = "BADSIG"
	DNSTimeseriesGroupMatchingAnswerParamsResponseCodeBadkey    DNSTimeseriesGroupMatchingAnswerParamsResponseCode = "BADKEY"
	DNSTimeseriesGroupMatchingAnswerParamsResponseCodeBadtime   DNSTimeseriesGroupMatchingAnswerParamsResponseCode = "BADTIME"
	DNSTimeseriesGroupMatchingAnswerParamsResponseCodeBadmode   DNSTimeseriesGroupMatchingAnswerParamsResponseCode = "BADMODE"
	DNSTimeseriesGroupMatchingAnswerParamsResponseCodeBadname   DNSTimeseriesGroupMatchingAnswerParamsResponseCode = "BADNAME"
	DNSTimeseriesGroupMatchingAnswerParamsResponseCodeBadalg    DNSTimeseriesGroupMatchingAnswerParamsResponseCode = "BADALG"
	DNSTimeseriesGroupMatchingAnswerParamsResponseCodeBadtrunc  DNSTimeseriesGroupMatchingAnswerParamsResponseCode = "BADTRUNC"
	DNSTimeseriesGroupMatchingAnswerParamsResponseCodeBadcookie DNSTimeseriesGroupMatchingAnswerParamsResponseCode = "BADCOOKIE"
)

func (r DNSTimeseriesGroupMatchingAnswerParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupMatchingAnswerParamsResponseCodeNoerror, DNSTimeseriesGroupMatchingAnswerParamsResponseCodeFormerr, DNSTimeseriesGroupMatchingAnswerParamsResponseCodeServfail, DNSTimeseriesGroupMatchingAnswerParamsResponseCodeNxdomain, DNSTimeseriesGroupMatchingAnswerParamsResponseCodeNotimp, DNSTimeseriesGroupMatchingAnswerParamsResponseCodeRefused, DNSTimeseriesGroupMatchingAnswerParamsResponseCodeYxdomain, DNSTimeseriesGroupMatchingAnswerParamsResponseCodeYxrrset, DNSTimeseriesGroupMatchingAnswerParamsResponseCodeNxrrset, DNSTimeseriesGroupMatchingAnswerParamsResponseCodeNotauth, DNSTimeseriesGroupMatchingAnswerParamsResponseCodeNotzone, DNSTimeseriesGroupMatchingAnswerParamsResponseCodeBadsig, DNSTimeseriesGroupMatchingAnswerParamsResponseCodeBadkey, DNSTimeseriesGroupMatchingAnswerParamsResponseCodeBadtime, DNSTimeseriesGroupMatchingAnswerParamsResponseCodeBadmode, DNSTimeseriesGroupMatchingAnswerParamsResponseCodeBadname, DNSTimeseriesGroupMatchingAnswerParamsResponseCodeBadalg, DNSTimeseriesGroupMatchingAnswerParamsResponseCodeBadtrunc, DNSTimeseriesGroupMatchingAnswerParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSTimeseriesGroupMatchingAnswerResponseEnvelope struct {
	Result  DNSTimeseriesGroupMatchingAnswerResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    dnsTimeseriesGroupMatchingAnswerResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupMatchingAnswerResponseEnvelopeJSON contains the JSON metadata
// for the struct [DNSTimeseriesGroupMatchingAnswerResponseEnvelope]
type dnsTimeseriesGroupMatchingAnswerResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupMatchingAnswerResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupMatchingAnswerResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupProtocolParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[DNSTimeseriesGroupProtocolParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[DNSTimeseriesGroupProtocolParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSTimeseriesGroupProtocolParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSTimeseriesGroupProtocolParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupProtocolParams]'s query parameters as
// `url.Values`.
func (r DNSTimeseriesGroupProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupProtocolParamsAggInterval string

const (
	DNSTimeseriesGroupProtocolParamsAggInterval15m DNSTimeseriesGroupProtocolParamsAggInterval = "15m"
	DNSTimeseriesGroupProtocolParamsAggInterval1h  DNSTimeseriesGroupProtocolParamsAggInterval = "1h"
	DNSTimeseriesGroupProtocolParamsAggInterval1d  DNSTimeseriesGroupProtocolParamsAggInterval = "1d"
	DNSTimeseriesGroupProtocolParamsAggInterval1w  DNSTimeseriesGroupProtocolParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupProtocolParamsAggInterval) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupProtocolParamsAggInterval15m, DNSTimeseriesGroupProtocolParamsAggInterval1h, DNSTimeseriesGroupProtocolParamsAggInterval1d, DNSTimeseriesGroupProtocolParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type DNSTimeseriesGroupProtocolParamsFormat string

const (
	DNSTimeseriesGroupProtocolParamsFormatJson DNSTimeseriesGroupProtocolParamsFormat = "JSON"
	DNSTimeseriesGroupProtocolParamsFormatCsv  DNSTimeseriesGroupProtocolParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupProtocolParamsFormat) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupProtocolParamsFormatJson, DNSTimeseriesGroupProtocolParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSTimeseriesGroupProtocolParamsQueryType string

const (
	DNSTimeseriesGroupProtocolParamsQueryTypeA          DNSTimeseriesGroupProtocolParamsQueryType = "A"
	DNSTimeseriesGroupProtocolParamsQueryTypeAAAA       DNSTimeseriesGroupProtocolParamsQueryType = "AAAA"
	DNSTimeseriesGroupProtocolParamsQueryTypeA6         DNSTimeseriesGroupProtocolParamsQueryType = "A6"
	DNSTimeseriesGroupProtocolParamsQueryTypeAfsdb      DNSTimeseriesGroupProtocolParamsQueryType = "AFSDB"
	DNSTimeseriesGroupProtocolParamsQueryTypeAny        DNSTimeseriesGroupProtocolParamsQueryType = "ANY"
	DNSTimeseriesGroupProtocolParamsQueryTypeApl        DNSTimeseriesGroupProtocolParamsQueryType = "APL"
	DNSTimeseriesGroupProtocolParamsQueryTypeAtma       DNSTimeseriesGroupProtocolParamsQueryType = "ATMA"
	DNSTimeseriesGroupProtocolParamsQueryTypeAXFR       DNSTimeseriesGroupProtocolParamsQueryType = "AXFR"
	DNSTimeseriesGroupProtocolParamsQueryTypeCAA        DNSTimeseriesGroupProtocolParamsQueryType = "CAA"
	DNSTimeseriesGroupProtocolParamsQueryTypeCdnskey    DNSTimeseriesGroupProtocolParamsQueryType = "CDNSKEY"
	DNSTimeseriesGroupProtocolParamsQueryTypeCds        DNSTimeseriesGroupProtocolParamsQueryType = "CDS"
	DNSTimeseriesGroupProtocolParamsQueryTypeCERT       DNSTimeseriesGroupProtocolParamsQueryType = "CERT"
	DNSTimeseriesGroupProtocolParamsQueryTypeCNAME      DNSTimeseriesGroupProtocolParamsQueryType = "CNAME"
	DNSTimeseriesGroupProtocolParamsQueryTypeCsync      DNSTimeseriesGroupProtocolParamsQueryType = "CSYNC"
	DNSTimeseriesGroupProtocolParamsQueryTypeDhcid      DNSTimeseriesGroupProtocolParamsQueryType = "DHCID"
	DNSTimeseriesGroupProtocolParamsQueryTypeDlv        DNSTimeseriesGroupProtocolParamsQueryType = "DLV"
	DNSTimeseriesGroupProtocolParamsQueryTypeDname      DNSTimeseriesGroupProtocolParamsQueryType = "DNAME"
	DNSTimeseriesGroupProtocolParamsQueryTypeDNSKEY     DNSTimeseriesGroupProtocolParamsQueryType = "DNSKEY"
	DNSTimeseriesGroupProtocolParamsQueryTypeDoa        DNSTimeseriesGroupProtocolParamsQueryType = "DOA"
	DNSTimeseriesGroupProtocolParamsQueryTypeDS         DNSTimeseriesGroupProtocolParamsQueryType = "DS"
	DNSTimeseriesGroupProtocolParamsQueryTypeEid        DNSTimeseriesGroupProtocolParamsQueryType = "EID"
	DNSTimeseriesGroupProtocolParamsQueryTypeEui48      DNSTimeseriesGroupProtocolParamsQueryType = "EUI48"
	DNSTimeseriesGroupProtocolParamsQueryTypeEui64      DNSTimeseriesGroupProtocolParamsQueryType = "EUI64"
	DNSTimeseriesGroupProtocolParamsQueryTypeGpos       DNSTimeseriesGroupProtocolParamsQueryType = "GPOS"
	DNSTimeseriesGroupProtocolParamsQueryTypeGid        DNSTimeseriesGroupProtocolParamsQueryType = "GID"
	DNSTimeseriesGroupProtocolParamsQueryTypeHinfo      DNSTimeseriesGroupProtocolParamsQueryType = "HINFO"
	DNSTimeseriesGroupProtocolParamsQueryTypeHip        DNSTimeseriesGroupProtocolParamsQueryType = "HIP"
	DNSTimeseriesGroupProtocolParamsQueryTypeHTTPS      DNSTimeseriesGroupProtocolParamsQueryType = "HTTPS"
	DNSTimeseriesGroupProtocolParamsQueryTypeIpseckey   DNSTimeseriesGroupProtocolParamsQueryType = "IPSECKEY"
	DNSTimeseriesGroupProtocolParamsQueryTypeIsdn       DNSTimeseriesGroupProtocolParamsQueryType = "ISDN"
	DNSTimeseriesGroupProtocolParamsQueryTypeIxfr       DNSTimeseriesGroupProtocolParamsQueryType = "IXFR"
	DNSTimeseriesGroupProtocolParamsQueryTypeKey        DNSTimeseriesGroupProtocolParamsQueryType = "KEY"
	DNSTimeseriesGroupProtocolParamsQueryTypeKx         DNSTimeseriesGroupProtocolParamsQueryType = "KX"
	DNSTimeseriesGroupProtocolParamsQueryTypeL32        DNSTimeseriesGroupProtocolParamsQueryType = "L32"
	DNSTimeseriesGroupProtocolParamsQueryTypeL64        DNSTimeseriesGroupProtocolParamsQueryType = "L64"
	DNSTimeseriesGroupProtocolParamsQueryTypeLOC        DNSTimeseriesGroupProtocolParamsQueryType = "LOC"
	DNSTimeseriesGroupProtocolParamsQueryTypeLp         DNSTimeseriesGroupProtocolParamsQueryType = "LP"
	DNSTimeseriesGroupProtocolParamsQueryTypeMaila      DNSTimeseriesGroupProtocolParamsQueryType = "MAILA"
	DNSTimeseriesGroupProtocolParamsQueryTypeMailb      DNSTimeseriesGroupProtocolParamsQueryType = "MAILB"
	DNSTimeseriesGroupProtocolParamsQueryTypeMB         DNSTimeseriesGroupProtocolParamsQueryType = "MB"
	DNSTimeseriesGroupProtocolParamsQueryTypeMd         DNSTimeseriesGroupProtocolParamsQueryType = "MD"
	DNSTimeseriesGroupProtocolParamsQueryTypeMf         DNSTimeseriesGroupProtocolParamsQueryType = "MF"
	DNSTimeseriesGroupProtocolParamsQueryTypeMg         DNSTimeseriesGroupProtocolParamsQueryType = "MG"
	DNSTimeseriesGroupProtocolParamsQueryTypeMinfo      DNSTimeseriesGroupProtocolParamsQueryType = "MINFO"
	DNSTimeseriesGroupProtocolParamsQueryTypeMr         DNSTimeseriesGroupProtocolParamsQueryType = "MR"
	DNSTimeseriesGroupProtocolParamsQueryTypeMX         DNSTimeseriesGroupProtocolParamsQueryType = "MX"
	DNSTimeseriesGroupProtocolParamsQueryTypeNAPTR      DNSTimeseriesGroupProtocolParamsQueryType = "NAPTR"
	DNSTimeseriesGroupProtocolParamsQueryTypeNb         DNSTimeseriesGroupProtocolParamsQueryType = "NB"
	DNSTimeseriesGroupProtocolParamsQueryTypeNbstat     DNSTimeseriesGroupProtocolParamsQueryType = "NBSTAT"
	DNSTimeseriesGroupProtocolParamsQueryTypeNid        DNSTimeseriesGroupProtocolParamsQueryType = "NID"
	DNSTimeseriesGroupProtocolParamsQueryTypeNimloc     DNSTimeseriesGroupProtocolParamsQueryType = "NIMLOC"
	DNSTimeseriesGroupProtocolParamsQueryTypeNinfo      DNSTimeseriesGroupProtocolParamsQueryType = "NINFO"
	DNSTimeseriesGroupProtocolParamsQueryTypeNS         DNSTimeseriesGroupProtocolParamsQueryType = "NS"
	DNSTimeseriesGroupProtocolParamsQueryTypeNsap       DNSTimeseriesGroupProtocolParamsQueryType = "NSAP"
	DNSTimeseriesGroupProtocolParamsQueryTypeNsec       DNSTimeseriesGroupProtocolParamsQueryType = "NSEC"
	DNSTimeseriesGroupProtocolParamsQueryTypeNsec3      DNSTimeseriesGroupProtocolParamsQueryType = "NSEC3"
	DNSTimeseriesGroupProtocolParamsQueryTypeNsec3Param DNSTimeseriesGroupProtocolParamsQueryType = "NSEC3PARAM"
	DNSTimeseriesGroupProtocolParamsQueryTypeNull       DNSTimeseriesGroupProtocolParamsQueryType = "NULL"
	DNSTimeseriesGroupProtocolParamsQueryTypeNxt        DNSTimeseriesGroupProtocolParamsQueryType = "NXT"
	DNSTimeseriesGroupProtocolParamsQueryTypeOpenpgpkey DNSTimeseriesGroupProtocolParamsQueryType = "OPENPGPKEY"
	DNSTimeseriesGroupProtocolParamsQueryTypeOpt        DNSTimeseriesGroupProtocolParamsQueryType = "OPT"
	DNSTimeseriesGroupProtocolParamsQueryTypePTR        DNSTimeseriesGroupProtocolParamsQueryType = "PTR"
	DNSTimeseriesGroupProtocolParamsQueryTypePx         DNSTimeseriesGroupProtocolParamsQueryType = "PX"
	DNSTimeseriesGroupProtocolParamsQueryTypeRkey       DNSTimeseriesGroupProtocolParamsQueryType = "RKEY"
	DNSTimeseriesGroupProtocolParamsQueryTypeRp         DNSTimeseriesGroupProtocolParamsQueryType = "RP"
	DNSTimeseriesGroupProtocolParamsQueryTypeRrsig      DNSTimeseriesGroupProtocolParamsQueryType = "RRSIG"
	DNSTimeseriesGroupProtocolParamsQueryTypeRt         DNSTimeseriesGroupProtocolParamsQueryType = "RT"
	DNSTimeseriesGroupProtocolParamsQueryTypeSig        DNSTimeseriesGroupProtocolParamsQueryType = "SIG"
	DNSTimeseriesGroupProtocolParamsQueryTypeSink       DNSTimeseriesGroupProtocolParamsQueryType = "SINK"
	DNSTimeseriesGroupProtocolParamsQueryTypeSMIMEA     DNSTimeseriesGroupProtocolParamsQueryType = "SMIMEA"
	DNSTimeseriesGroupProtocolParamsQueryTypeSOA        DNSTimeseriesGroupProtocolParamsQueryType = "SOA"
	DNSTimeseriesGroupProtocolParamsQueryTypeSPF        DNSTimeseriesGroupProtocolParamsQueryType = "SPF"
	DNSTimeseriesGroupProtocolParamsQueryTypeSRV        DNSTimeseriesGroupProtocolParamsQueryType = "SRV"
	DNSTimeseriesGroupProtocolParamsQueryTypeSSHFP      DNSTimeseriesGroupProtocolParamsQueryType = "SSHFP"
	DNSTimeseriesGroupProtocolParamsQueryTypeSVCB       DNSTimeseriesGroupProtocolParamsQueryType = "SVCB"
	DNSTimeseriesGroupProtocolParamsQueryTypeTa         DNSTimeseriesGroupProtocolParamsQueryType = "TA"
	DNSTimeseriesGroupProtocolParamsQueryTypeTalink     DNSTimeseriesGroupProtocolParamsQueryType = "TALINK"
	DNSTimeseriesGroupProtocolParamsQueryTypeTkey       DNSTimeseriesGroupProtocolParamsQueryType = "TKEY"
	DNSTimeseriesGroupProtocolParamsQueryTypeTLSA       DNSTimeseriesGroupProtocolParamsQueryType = "TLSA"
	DNSTimeseriesGroupProtocolParamsQueryTypeTSIG       DNSTimeseriesGroupProtocolParamsQueryType = "TSIG"
	DNSTimeseriesGroupProtocolParamsQueryTypeTXT        DNSTimeseriesGroupProtocolParamsQueryType = "TXT"
	DNSTimeseriesGroupProtocolParamsQueryTypeUinfo      DNSTimeseriesGroupProtocolParamsQueryType = "UINFO"
	DNSTimeseriesGroupProtocolParamsQueryTypeUID        DNSTimeseriesGroupProtocolParamsQueryType = "UID"
	DNSTimeseriesGroupProtocolParamsQueryTypeUnspec     DNSTimeseriesGroupProtocolParamsQueryType = "UNSPEC"
	DNSTimeseriesGroupProtocolParamsQueryTypeURI        DNSTimeseriesGroupProtocolParamsQueryType = "URI"
	DNSTimeseriesGroupProtocolParamsQueryTypeWks        DNSTimeseriesGroupProtocolParamsQueryType = "WKS"
	DNSTimeseriesGroupProtocolParamsQueryTypeX25        DNSTimeseriesGroupProtocolParamsQueryType = "X25"
	DNSTimeseriesGroupProtocolParamsQueryTypeZonemd     DNSTimeseriesGroupProtocolParamsQueryType = "ZONEMD"
)

func (r DNSTimeseriesGroupProtocolParamsQueryType) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupProtocolParamsQueryTypeA, DNSTimeseriesGroupProtocolParamsQueryTypeAAAA, DNSTimeseriesGroupProtocolParamsQueryTypeA6, DNSTimeseriesGroupProtocolParamsQueryTypeAfsdb, DNSTimeseriesGroupProtocolParamsQueryTypeAny, DNSTimeseriesGroupProtocolParamsQueryTypeApl, DNSTimeseriesGroupProtocolParamsQueryTypeAtma, DNSTimeseriesGroupProtocolParamsQueryTypeAXFR, DNSTimeseriesGroupProtocolParamsQueryTypeCAA, DNSTimeseriesGroupProtocolParamsQueryTypeCdnskey, DNSTimeseriesGroupProtocolParamsQueryTypeCds, DNSTimeseriesGroupProtocolParamsQueryTypeCERT, DNSTimeseriesGroupProtocolParamsQueryTypeCNAME, DNSTimeseriesGroupProtocolParamsQueryTypeCsync, DNSTimeseriesGroupProtocolParamsQueryTypeDhcid, DNSTimeseriesGroupProtocolParamsQueryTypeDlv, DNSTimeseriesGroupProtocolParamsQueryTypeDname, DNSTimeseriesGroupProtocolParamsQueryTypeDNSKEY, DNSTimeseriesGroupProtocolParamsQueryTypeDoa, DNSTimeseriesGroupProtocolParamsQueryTypeDS, DNSTimeseriesGroupProtocolParamsQueryTypeEid, DNSTimeseriesGroupProtocolParamsQueryTypeEui48, DNSTimeseriesGroupProtocolParamsQueryTypeEui64, DNSTimeseriesGroupProtocolParamsQueryTypeGpos, DNSTimeseriesGroupProtocolParamsQueryTypeGid, DNSTimeseriesGroupProtocolParamsQueryTypeHinfo, DNSTimeseriesGroupProtocolParamsQueryTypeHip, DNSTimeseriesGroupProtocolParamsQueryTypeHTTPS, DNSTimeseriesGroupProtocolParamsQueryTypeIpseckey, DNSTimeseriesGroupProtocolParamsQueryTypeIsdn, DNSTimeseriesGroupProtocolParamsQueryTypeIxfr, DNSTimeseriesGroupProtocolParamsQueryTypeKey, DNSTimeseriesGroupProtocolParamsQueryTypeKx, DNSTimeseriesGroupProtocolParamsQueryTypeL32, DNSTimeseriesGroupProtocolParamsQueryTypeL64, DNSTimeseriesGroupProtocolParamsQueryTypeLOC, DNSTimeseriesGroupProtocolParamsQueryTypeLp, DNSTimeseriesGroupProtocolParamsQueryTypeMaila, DNSTimeseriesGroupProtocolParamsQueryTypeMailb, DNSTimeseriesGroupProtocolParamsQueryTypeMB, DNSTimeseriesGroupProtocolParamsQueryTypeMd, DNSTimeseriesGroupProtocolParamsQueryTypeMf, DNSTimeseriesGroupProtocolParamsQueryTypeMg, DNSTimeseriesGroupProtocolParamsQueryTypeMinfo, DNSTimeseriesGroupProtocolParamsQueryTypeMr, DNSTimeseriesGroupProtocolParamsQueryTypeMX, DNSTimeseriesGroupProtocolParamsQueryTypeNAPTR, DNSTimeseriesGroupProtocolParamsQueryTypeNb, DNSTimeseriesGroupProtocolParamsQueryTypeNbstat, DNSTimeseriesGroupProtocolParamsQueryTypeNid, DNSTimeseriesGroupProtocolParamsQueryTypeNimloc, DNSTimeseriesGroupProtocolParamsQueryTypeNinfo, DNSTimeseriesGroupProtocolParamsQueryTypeNS, DNSTimeseriesGroupProtocolParamsQueryTypeNsap, DNSTimeseriesGroupProtocolParamsQueryTypeNsec, DNSTimeseriesGroupProtocolParamsQueryTypeNsec3, DNSTimeseriesGroupProtocolParamsQueryTypeNsec3Param, DNSTimeseriesGroupProtocolParamsQueryTypeNull, DNSTimeseriesGroupProtocolParamsQueryTypeNxt, DNSTimeseriesGroupProtocolParamsQueryTypeOpenpgpkey, DNSTimeseriesGroupProtocolParamsQueryTypeOpt, DNSTimeseriesGroupProtocolParamsQueryTypePTR, DNSTimeseriesGroupProtocolParamsQueryTypePx, DNSTimeseriesGroupProtocolParamsQueryTypeRkey, DNSTimeseriesGroupProtocolParamsQueryTypeRp, DNSTimeseriesGroupProtocolParamsQueryTypeRrsig, DNSTimeseriesGroupProtocolParamsQueryTypeRt, DNSTimeseriesGroupProtocolParamsQueryTypeSig, DNSTimeseriesGroupProtocolParamsQueryTypeSink, DNSTimeseriesGroupProtocolParamsQueryTypeSMIMEA, DNSTimeseriesGroupProtocolParamsQueryTypeSOA, DNSTimeseriesGroupProtocolParamsQueryTypeSPF, DNSTimeseriesGroupProtocolParamsQueryTypeSRV, DNSTimeseriesGroupProtocolParamsQueryTypeSSHFP, DNSTimeseriesGroupProtocolParamsQueryTypeSVCB, DNSTimeseriesGroupProtocolParamsQueryTypeTa, DNSTimeseriesGroupProtocolParamsQueryTypeTalink, DNSTimeseriesGroupProtocolParamsQueryTypeTkey, DNSTimeseriesGroupProtocolParamsQueryTypeTLSA, DNSTimeseriesGroupProtocolParamsQueryTypeTSIG, DNSTimeseriesGroupProtocolParamsQueryTypeTXT, DNSTimeseriesGroupProtocolParamsQueryTypeUinfo, DNSTimeseriesGroupProtocolParamsQueryTypeUID, DNSTimeseriesGroupProtocolParamsQueryTypeUnspec, DNSTimeseriesGroupProtocolParamsQueryTypeURI, DNSTimeseriesGroupProtocolParamsQueryTypeWks, DNSTimeseriesGroupProtocolParamsQueryTypeX25, DNSTimeseriesGroupProtocolParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSTimeseriesGroupProtocolParamsResponseCode string

const (
	DNSTimeseriesGroupProtocolParamsResponseCodeNoerror   DNSTimeseriesGroupProtocolParamsResponseCode = "NOERROR"
	DNSTimeseriesGroupProtocolParamsResponseCodeFormerr   DNSTimeseriesGroupProtocolParamsResponseCode = "FORMERR"
	DNSTimeseriesGroupProtocolParamsResponseCodeServfail  DNSTimeseriesGroupProtocolParamsResponseCode = "SERVFAIL"
	DNSTimeseriesGroupProtocolParamsResponseCodeNxdomain  DNSTimeseriesGroupProtocolParamsResponseCode = "NXDOMAIN"
	DNSTimeseriesGroupProtocolParamsResponseCodeNotimp    DNSTimeseriesGroupProtocolParamsResponseCode = "NOTIMP"
	DNSTimeseriesGroupProtocolParamsResponseCodeRefused   DNSTimeseriesGroupProtocolParamsResponseCode = "REFUSED"
	DNSTimeseriesGroupProtocolParamsResponseCodeYxdomain  DNSTimeseriesGroupProtocolParamsResponseCode = "YXDOMAIN"
	DNSTimeseriesGroupProtocolParamsResponseCodeYxrrset   DNSTimeseriesGroupProtocolParamsResponseCode = "YXRRSET"
	DNSTimeseriesGroupProtocolParamsResponseCodeNxrrset   DNSTimeseriesGroupProtocolParamsResponseCode = "NXRRSET"
	DNSTimeseriesGroupProtocolParamsResponseCodeNotauth   DNSTimeseriesGroupProtocolParamsResponseCode = "NOTAUTH"
	DNSTimeseriesGroupProtocolParamsResponseCodeNotzone   DNSTimeseriesGroupProtocolParamsResponseCode = "NOTZONE"
	DNSTimeseriesGroupProtocolParamsResponseCodeBadsig    DNSTimeseriesGroupProtocolParamsResponseCode = "BADSIG"
	DNSTimeseriesGroupProtocolParamsResponseCodeBadkey    DNSTimeseriesGroupProtocolParamsResponseCode = "BADKEY"
	DNSTimeseriesGroupProtocolParamsResponseCodeBadtime   DNSTimeseriesGroupProtocolParamsResponseCode = "BADTIME"
	DNSTimeseriesGroupProtocolParamsResponseCodeBadmode   DNSTimeseriesGroupProtocolParamsResponseCode = "BADMODE"
	DNSTimeseriesGroupProtocolParamsResponseCodeBadname   DNSTimeseriesGroupProtocolParamsResponseCode = "BADNAME"
	DNSTimeseriesGroupProtocolParamsResponseCodeBadalg    DNSTimeseriesGroupProtocolParamsResponseCode = "BADALG"
	DNSTimeseriesGroupProtocolParamsResponseCodeBadtrunc  DNSTimeseriesGroupProtocolParamsResponseCode = "BADTRUNC"
	DNSTimeseriesGroupProtocolParamsResponseCodeBadcookie DNSTimeseriesGroupProtocolParamsResponseCode = "BADCOOKIE"
)

func (r DNSTimeseriesGroupProtocolParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupProtocolParamsResponseCodeNoerror, DNSTimeseriesGroupProtocolParamsResponseCodeFormerr, DNSTimeseriesGroupProtocolParamsResponseCodeServfail, DNSTimeseriesGroupProtocolParamsResponseCodeNxdomain, DNSTimeseriesGroupProtocolParamsResponseCodeNotimp, DNSTimeseriesGroupProtocolParamsResponseCodeRefused, DNSTimeseriesGroupProtocolParamsResponseCodeYxdomain, DNSTimeseriesGroupProtocolParamsResponseCodeYxrrset, DNSTimeseriesGroupProtocolParamsResponseCodeNxrrset, DNSTimeseriesGroupProtocolParamsResponseCodeNotauth, DNSTimeseriesGroupProtocolParamsResponseCodeNotzone, DNSTimeseriesGroupProtocolParamsResponseCodeBadsig, DNSTimeseriesGroupProtocolParamsResponseCodeBadkey, DNSTimeseriesGroupProtocolParamsResponseCodeBadtime, DNSTimeseriesGroupProtocolParamsResponseCodeBadmode, DNSTimeseriesGroupProtocolParamsResponseCodeBadname, DNSTimeseriesGroupProtocolParamsResponseCodeBadalg, DNSTimeseriesGroupProtocolParamsResponseCodeBadtrunc, DNSTimeseriesGroupProtocolParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSTimeseriesGroupProtocolResponseEnvelope struct {
	Result  DNSTimeseriesGroupProtocolResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    dnsTimeseriesGroupProtocolResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupProtocolResponseEnvelopeJSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupProtocolResponseEnvelope]
type dnsTimeseriesGroupProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupProtocolResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupQueryTypeParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[DNSTimeseriesGroupQueryTypeParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[DNSTimeseriesGroupQueryTypeParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. If there are more items than the limit, the response will include
	// the count of items, with any remaining items grouped together under an "other"
	// category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSTimeseriesGroupQueryTypeParamsProtocol] `query:"protocol"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSTimeseriesGroupQueryTypeParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupQueryTypeParams]'s query parameters as
// `url.Values`.
func (r DNSTimeseriesGroupQueryTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupQueryTypeParamsAggInterval string

const (
	DNSTimeseriesGroupQueryTypeParamsAggInterval15m DNSTimeseriesGroupQueryTypeParamsAggInterval = "15m"
	DNSTimeseriesGroupQueryTypeParamsAggInterval1h  DNSTimeseriesGroupQueryTypeParamsAggInterval = "1h"
	DNSTimeseriesGroupQueryTypeParamsAggInterval1d  DNSTimeseriesGroupQueryTypeParamsAggInterval = "1d"
	DNSTimeseriesGroupQueryTypeParamsAggInterval1w  DNSTimeseriesGroupQueryTypeParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupQueryTypeParamsAggInterval) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupQueryTypeParamsAggInterval15m, DNSTimeseriesGroupQueryTypeParamsAggInterval1h, DNSTimeseriesGroupQueryTypeParamsAggInterval1d, DNSTimeseriesGroupQueryTypeParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type DNSTimeseriesGroupQueryTypeParamsFormat string

const (
	DNSTimeseriesGroupQueryTypeParamsFormatJson DNSTimeseriesGroupQueryTypeParamsFormat = "JSON"
	DNSTimeseriesGroupQueryTypeParamsFormatCsv  DNSTimeseriesGroupQueryTypeParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupQueryTypeParamsFormat) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupQueryTypeParamsFormatJson, DNSTimeseriesGroupQueryTypeParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSTimeseriesGroupQueryTypeParamsProtocol string

const (
	DNSTimeseriesGroupQueryTypeParamsProtocolUdp   DNSTimeseriesGroupQueryTypeParamsProtocol = "UDP"
	DNSTimeseriesGroupQueryTypeParamsProtocolTCP   DNSTimeseriesGroupQueryTypeParamsProtocol = "TCP"
	DNSTimeseriesGroupQueryTypeParamsProtocolHTTPS DNSTimeseriesGroupQueryTypeParamsProtocol = "HTTPS"
	DNSTimeseriesGroupQueryTypeParamsProtocolTLS   DNSTimeseriesGroupQueryTypeParamsProtocol = "TLS"
)

func (r DNSTimeseriesGroupQueryTypeParamsProtocol) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupQueryTypeParamsProtocolUdp, DNSTimeseriesGroupQueryTypeParamsProtocolTCP, DNSTimeseriesGroupQueryTypeParamsProtocolHTTPS, DNSTimeseriesGroupQueryTypeParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSTimeseriesGroupQueryTypeParamsResponseCode string

const (
	DNSTimeseriesGroupQueryTypeParamsResponseCodeNoerror   DNSTimeseriesGroupQueryTypeParamsResponseCode = "NOERROR"
	DNSTimeseriesGroupQueryTypeParamsResponseCodeFormerr   DNSTimeseriesGroupQueryTypeParamsResponseCode = "FORMERR"
	DNSTimeseriesGroupQueryTypeParamsResponseCodeServfail  DNSTimeseriesGroupQueryTypeParamsResponseCode = "SERVFAIL"
	DNSTimeseriesGroupQueryTypeParamsResponseCodeNxdomain  DNSTimeseriesGroupQueryTypeParamsResponseCode = "NXDOMAIN"
	DNSTimeseriesGroupQueryTypeParamsResponseCodeNotimp    DNSTimeseriesGroupQueryTypeParamsResponseCode = "NOTIMP"
	DNSTimeseriesGroupQueryTypeParamsResponseCodeRefused   DNSTimeseriesGroupQueryTypeParamsResponseCode = "REFUSED"
	DNSTimeseriesGroupQueryTypeParamsResponseCodeYxdomain  DNSTimeseriesGroupQueryTypeParamsResponseCode = "YXDOMAIN"
	DNSTimeseriesGroupQueryTypeParamsResponseCodeYxrrset   DNSTimeseriesGroupQueryTypeParamsResponseCode = "YXRRSET"
	DNSTimeseriesGroupQueryTypeParamsResponseCodeNxrrset   DNSTimeseriesGroupQueryTypeParamsResponseCode = "NXRRSET"
	DNSTimeseriesGroupQueryTypeParamsResponseCodeNotauth   DNSTimeseriesGroupQueryTypeParamsResponseCode = "NOTAUTH"
	DNSTimeseriesGroupQueryTypeParamsResponseCodeNotzone   DNSTimeseriesGroupQueryTypeParamsResponseCode = "NOTZONE"
	DNSTimeseriesGroupQueryTypeParamsResponseCodeBadsig    DNSTimeseriesGroupQueryTypeParamsResponseCode = "BADSIG"
	DNSTimeseriesGroupQueryTypeParamsResponseCodeBadkey    DNSTimeseriesGroupQueryTypeParamsResponseCode = "BADKEY"
	DNSTimeseriesGroupQueryTypeParamsResponseCodeBadtime   DNSTimeseriesGroupQueryTypeParamsResponseCode = "BADTIME"
	DNSTimeseriesGroupQueryTypeParamsResponseCodeBadmode   DNSTimeseriesGroupQueryTypeParamsResponseCode = "BADMODE"
	DNSTimeseriesGroupQueryTypeParamsResponseCodeBadname   DNSTimeseriesGroupQueryTypeParamsResponseCode = "BADNAME"
	DNSTimeseriesGroupQueryTypeParamsResponseCodeBadalg    DNSTimeseriesGroupQueryTypeParamsResponseCode = "BADALG"
	DNSTimeseriesGroupQueryTypeParamsResponseCodeBadtrunc  DNSTimeseriesGroupQueryTypeParamsResponseCode = "BADTRUNC"
	DNSTimeseriesGroupQueryTypeParamsResponseCodeBadcookie DNSTimeseriesGroupQueryTypeParamsResponseCode = "BADCOOKIE"
)

func (r DNSTimeseriesGroupQueryTypeParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupQueryTypeParamsResponseCodeNoerror, DNSTimeseriesGroupQueryTypeParamsResponseCodeFormerr, DNSTimeseriesGroupQueryTypeParamsResponseCodeServfail, DNSTimeseriesGroupQueryTypeParamsResponseCodeNxdomain, DNSTimeseriesGroupQueryTypeParamsResponseCodeNotimp, DNSTimeseriesGroupQueryTypeParamsResponseCodeRefused, DNSTimeseriesGroupQueryTypeParamsResponseCodeYxdomain, DNSTimeseriesGroupQueryTypeParamsResponseCodeYxrrset, DNSTimeseriesGroupQueryTypeParamsResponseCodeNxrrset, DNSTimeseriesGroupQueryTypeParamsResponseCodeNotauth, DNSTimeseriesGroupQueryTypeParamsResponseCodeNotzone, DNSTimeseriesGroupQueryTypeParamsResponseCodeBadsig, DNSTimeseriesGroupQueryTypeParamsResponseCodeBadkey, DNSTimeseriesGroupQueryTypeParamsResponseCodeBadtime, DNSTimeseriesGroupQueryTypeParamsResponseCodeBadmode, DNSTimeseriesGroupQueryTypeParamsResponseCodeBadname, DNSTimeseriesGroupQueryTypeParamsResponseCodeBadalg, DNSTimeseriesGroupQueryTypeParamsResponseCodeBadtrunc, DNSTimeseriesGroupQueryTypeParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSTimeseriesGroupQueryTypeResponseEnvelope struct {
	Result  DNSTimeseriesGroupQueryTypeResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    dnsTimeseriesGroupQueryTypeResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupQueryTypeResponseEnvelopeJSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupQueryTypeResponseEnvelope]
type dnsTimeseriesGroupQueryTypeResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupQueryTypeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupQueryTypeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupResponseCodeParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[DNSTimeseriesGroupResponseCodeParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[DNSTimeseriesGroupResponseCodeParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. If there are more items than the limit, the response will include
	// the count of items, with any remaining items grouped together under an "other"
	// category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSTimeseriesGroupResponseCodeParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSTimeseriesGroupResponseCodeParamsQueryType] `query:"queryType"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupResponseCodeParams]'s query parameters as
// `url.Values`.
func (r DNSTimeseriesGroupResponseCodeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupResponseCodeParamsAggInterval string

const (
	DNSTimeseriesGroupResponseCodeParamsAggInterval15m DNSTimeseriesGroupResponseCodeParamsAggInterval = "15m"
	DNSTimeseriesGroupResponseCodeParamsAggInterval1h  DNSTimeseriesGroupResponseCodeParamsAggInterval = "1h"
	DNSTimeseriesGroupResponseCodeParamsAggInterval1d  DNSTimeseriesGroupResponseCodeParamsAggInterval = "1d"
	DNSTimeseriesGroupResponseCodeParamsAggInterval1w  DNSTimeseriesGroupResponseCodeParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupResponseCodeParamsAggInterval) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupResponseCodeParamsAggInterval15m, DNSTimeseriesGroupResponseCodeParamsAggInterval1h, DNSTimeseriesGroupResponseCodeParamsAggInterval1d, DNSTimeseriesGroupResponseCodeParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type DNSTimeseriesGroupResponseCodeParamsFormat string

const (
	DNSTimeseriesGroupResponseCodeParamsFormatJson DNSTimeseriesGroupResponseCodeParamsFormat = "JSON"
	DNSTimeseriesGroupResponseCodeParamsFormatCsv  DNSTimeseriesGroupResponseCodeParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupResponseCodeParamsFormat) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupResponseCodeParamsFormatJson, DNSTimeseriesGroupResponseCodeParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSTimeseriesGroupResponseCodeParamsProtocol string

const (
	DNSTimeseriesGroupResponseCodeParamsProtocolUdp   DNSTimeseriesGroupResponseCodeParamsProtocol = "UDP"
	DNSTimeseriesGroupResponseCodeParamsProtocolTCP   DNSTimeseriesGroupResponseCodeParamsProtocol = "TCP"
	DNSTimeseriesGroupResponseCodeParamsProtocolHTTPS DNSTimeseriesGroupResponseCodeParamsProtocol = "HTTPS"
	DNSTimeseriesGroupResponseCodeParamsProtocolTLS   DNSTimeseriesGroupResponseCodeParamsProtocol = "TLS"
)

func (r DNSTimeseriesGroupResponseCodeParamsProtocol) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupResponseCodeParamsProtocolUdp, DNSTimeseriesGroupResponseCodeParamsProtocolTCP, DNSTimeseriesGroupResponseCodeParamsProtocolHTTPS, DNSTimeseriesGroupResponseCodeParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSTimeseriesGroupResponseCodeParamsQueryType string

const (
	DNSTimeseriesGroupResponseCodeParamsQueryTypeA          DNSTimeseriesGroupResponseCodeParamsQueryType = "A"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeAAAA       DNSTimeseriesGroupResponseCodeParamsQueryType = "AAAA"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeA6         DNSTimeseriesGroupResponseCodeParamsQueryType = "A6"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeAfsdb      DNSTimeseriesGroupResponseCodeParamsQueryType = "AFSDB"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeAny        DNSTimeseriesGroupResponseCodeParamsQueryType = "ANY"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeApl        DNSTimeseriesGroupResponseCodeParamsQueryType = "APL"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeAtma       DNSTimeseriesGroupResponseCodeParamsQueryType = "ATMA"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeAXFR       DNSTimeseriesGroupResponseCodeParamsQueryType = "AXFR"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeCAA        DNSTimeseriesGroupResponseCodeParamsQueryType = "CAA"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeCdnskey    DNSTimeseriesGroupResponseCodeParamsQueryType = "CDNSKEY"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeCds        DNSTimeseriesGroupResponseCodeParamsQueryType = "CDS"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeCERT       DNSTimeseriesGroupResponseCodeParamsQueryType = "CERT"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeCNAME      DNSTimeseriesGroupResponseCodeParamsQueryType = "CNAME"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeCsync      DNSTimeseriesGroupResponseCodeParamsQueryType = "CSYNC"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeDhcid      DNSTimeseriesGroupResponseCodeParamsQueryType = "DHCID"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeDlv        DNSTimeseriesGroupResponseCodeParamsQueryType = "DLV"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeDname      DNSTimeseriesGroupResponseCodeParamsQueryType = "DNAME"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeDNSKEY     DNSTimeseriesGroupResponseCodeParamsQueryType = "DNSKEY"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeDoa        DNSTimeseriesGroupResponseCodeParamsQueryType = "DOA"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeDS         DNSTimeseriesGroupResponseCodeParamsQueryType = "DS"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeEid        DNSTimeseriesGroupResponseCodeParamsQueryType = "EID"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeEui48      DNSTimeseriesGroupResponseCodeParamsQueryType = "EUI48"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeEui64      DNSTimeseriesGroupResponseCodeParamsQueryType = "EUI64"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeGpos       DNSTimeseriesGroupResponseCodeParamsQueryType = "GPOS"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeGid        DNSTimeseriesGroupResponseCodeParamsQueryType = "GID"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeHinfo      DNSTimeseriesGroupResponseCodeParamsQueryType = "HINFO"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeHip        DNSTimeseriesGroupResponseCodeParamsQueryType = "HIP"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeHTTPS      DNSTimeseriesGroupResponseCodeParamsQueryType = "HTTPS"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeIpseckey   DNSTimeseriesGroupResponseCodeParamsQueryType = "IPSECKEY"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeIsdn       DNSTimeseriesGroupResponseCodeParamsQueryType = "ISDN"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeIxfr       DNSTimeseriesGroupResponseCodeParamsQueryType = "IXFR"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeKey        DNSTimeseriesGroupResponseCodeParamsQueryType = "KEY"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeKx         DNSTimeseriesGroupResponseCodeParamsQueryType = "KX"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeL32        DNSTimeseriesGroupResponseCodeParamsQueryType = "L32"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeL64        DNSTimeseriesGroupResponseCodeParamsQueryType = "L64"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeLOC        DNSTimeseriesGroupResponseCodeParamsQueryType = "LOC"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeLp         DNSTimeseriesGroupResponseCodeParamsQueryType = "LP"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeMaila      DNSTimeseriesGroupResponseCodeParamsQueryType = "MAILA"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeMailb      DNSTimeseriesGroupResponseCodeParamsQueryType = "MAILB"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeMB         DNSTimeseriesGroupResponseCodeParamsQueryType = "MB"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeMd         DNSTimeseriesGroupResponseCodeParamsQueryType = "MD"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeMf         DNSTimeseriesGroupResponseCodeParamsQueryType = "MF"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeMg         DNSTimeseriesGroupResponseCodeParamsQueryType = "MG"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeMinfo      DNSTimeseriesGroupResponseCodeParamsQueryType = "MINFO"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeMr         DNSTimeseriesGroupResponseCodeParamsQueryType = "MR"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeMX         DNSTimeseriesGroupResponseCodeParamsQueryType = "MX"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeNAPTR      DNSTimeseriesGroupResponseCodeParamsQueryType = "NAPTR"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeNb         DNSTimeseriesGroupResponseCodeParamsQueryType = "NB"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeNbstat     DNSTimeseriesGroupResponseCodeParamsQueryType = "NBSTAT"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeNid        DNSTimeseriesGroupResponseCodeParamsQueryType = "NID"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeNimloc     DNSTimeseriesGroupResponseCodeParamsQueryType = "NIMLOC"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeNinfo      DNSTimeseriesGroupResponseCodeParamsQueryType = "NINFO"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeNS         DNSTimeseriesGroupResponseCodeParamsQueryType = "NS"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeNsap       DNSTimeseriesGroupResponseCodeParamsQueryType = "NSAP"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeNsec       DNSTimeseriesGroupResponseCodeParamsQueryType = "NSEC"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeNsec3      DNSTimeseriesGroupResponseCodeParamsQueryType = "NSEC3"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeNsec3Param DNSTimeseriesGroupResponseCodeParamsQueryType = "NSEC3PARAM"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeNull       DNSTimeseriesGroupResponseCodeParamsQueryType = "NULL"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeNxt        DNSTimeseriesGroupResponseCodeParamsQueryType = "NXT"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeOpenpgpkey DNSTimeseriesGroupResponseCodeParamsQueryType = "OPENPGPKEY"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeOpt        DNSTimeseriesGroupResponseCodeParamsQueryType = "OPT"
	DNSTimeseriesGroupResponseCodeParamsQueryTypePTR        DNSTimeseriesGroupResponseCodeParamsQueryType = "PTR"
	DNSTimeseriesGroupResponseCodeParamsQueryTypePx         DNSTimeseriesGroupResponseCodeParamsQueryType = "PX"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeRkey       DNSTimeseriesGroupResponseCodeParamsQueryType = "RKEY"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeRp         DNSTimeseriesGroupResponseCodeParamsQueryType = "RP"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeRrsig      DNSTimeseriesGroupResponseCodeParamsQueryType = "RRSIG"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeRt         DNSTimeseriesGroupResponseCodeParamsQueryType = "RT"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeSig        DNSTimeseriesGroupResponseCodeParamsQueryType = "SIG"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeSink       DNSTimeseriesGroupResponseCodeParamsQueryType = "SINK"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeSMIMEA     DNSTimeseriesGroupResponseCodeParamsQueryType = "SMIMEA"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeSOA        DNSTimeseriesGroupResponseCodeParamsQueryType = "SOA"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeSPF        DNSTimeseriesGroupResponseCodeParamsQueryType = "SPF"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeSRV        DNSTimeseriesGroupResponseCodeParamsQueryType = "SRV"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeSSHFP      DNSTimeseriesGroupResponseCodeParamsQueryType = "SSHFP"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeSVCB       DNSTimeseriesGroupResponseCodeParamsQueryType = "SVCB"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeTa         DNSTimeseriesGroupResponseCodeParamsQueryType = "TA"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeTalink     DNSTimeseriesGroupResponseCodeParamsQueryType = "TALINK"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeTkey       DNSTimeseriesGroupResponseCodeParamsQueryType = "TKEY"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeTLSA       DNSTimeseriesGroupResponseCodeParamsQueryType = "TLSA"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeTSIG       DNSTimeseriesGroupResponseCodeParamsQueryType = "TSIG"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeTXT        DNSTimeseriesGroupResponseCodeParamsQueryType = "TXT"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeUinfo      DNSTimeseriesGroupResponseCodeParamsQueryType = "UINFO"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeUID        DNSTimeseriesGroupResponseCodeParamsQueryType = "UID"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeUnspec     DNSTimeseriesGroupResponseCodeParamsQueryType = "UNSPEC"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeURI        DNSTimeseriesGroupResponseCodeParamsQueryType = "URI"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeWks        DNSTimeseriesGroupResponseCodeParamsQueryType = "WKS"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeX25        DNSTimeseriesGroupResponseCodeParamsQueryType = "X25"
	DNSTimeseriesGroupResponseCodeParamsQueryTypeZonemd     DNSTimeseriesGroupResponseCodeParamsQueryType = "ZONEMD"
)

func (r DNSTimeseriesGroupResponseCodeParamsQueryType) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupResponseCodeParamsQueryTypeA, DNSTimeseriesGroupResponseCodeParamsQueryTypeAAAA, DNSTimeseriesGroupResponseCodeParamsQueryTypeA6, DNSTimeseriesGroupResponseCodeParamsQueryTypeAfsdb, DNSTimeseriesGroupResponseCodeParamsQueryTypeAny, DNSTimeseriesGroupResponseCodeParamsQueryTypeApl, DNSTimeseriesGroupResponseCodeParamsQueryTypeAtma, DNSTimeseriesGroupResponseCodeParamsQueryTypeAXFR, DNSTimeseriesGroupResponseCodeParamsQueryTypeCAA, DNSTimeseriesGroupResponseCodeParamsQueryTypeCdnskey, DNSTimeseriesGroupResponseCodeParamsQueryTypeCds, DNSTimeseriesGroupResponseCodeParamsQueryTypeCERT, DNSTimeseriesGroupResponseCodeParamsQueryTypeCNAME, DNSTimeseriesGroupResponseCodeParamsQueryTypeCsync, DNSTimeseriesGroupResponseCodeParamsQueryTypeDhcid, DNSTimeseriesGroupResponseCodeParamsQueryTypeDlv, DNSTimeseriesGroupResponseCodeParamsQueryTypeDname, DNSTimeseriesGroupResponseCodeParamsQueryTypeDNSKEY, DNSTimeseriesGroupResponseCodeParamsQueryTypeDoa, DNSTimeseriesGroupResponseCodeParamsQueryTypeDS, DNSTimeseriesGroupResponseCodeParamsQueryTypeEid, DNSTimeseriesGroupResponseCodeParamsQueryTypeEui48, DNSTimeseriesGroupResponseCodeParamsQueryTypeEui64, DNSTimeseriesGroupResponseCodeParamsQueryTypeGpos, DNSTimeseriesGroupResponseCodeParamsQueryTypeGid, DNSTimeseriesGroupResponseCodeParamsQueryTypeHinfo, DNSTimeseriesGroupResponseCodeParamsQueryTypeHip, DNSTimeseriesGroupResponseCodeParamsQueryTypeHTTPS, DNSTimeseriesGroupResponseCodeParamsQueryTypeIpseckey, DNSTimeseriesGroupResponseCodeParamsQueryTypeIsdn, DNSTimeseriesGroupResponseCodeParamsQueryTypeIxfr, DNSTimeseriesGroupResponseCodeParamsQueryTypeKey, DNSTimeseriesGroupResponseCodeParamsQueryTypeKx, DNSTimeseriesGroupResponseCodeParamsQueryTypeL32, DNSTimeseriesGroupResponseCodeParamsQueryTypeL64, DNSTimeseriesGroupResponseCodeParamsQueryTypeLOC, DNSTimeseriesGroupResponseCodeParamsQueryTypeLp, DNSTimeseriesGroupResponseCodeParamsQueryTypeMaila, DNSTimeseriesGroupResponseCodeParamsQueryTypeMailb, DNSTimeseriesGroupResponseCodeParamsQueryTypeMB, DNSTimeseriesGroupResponseCodeParamsQueryTypeMd, DNSTimeseriesGroupResponseCodeParamsQueryTypeMf, DNSTimeseriesGroupResponseCodeParamsQueryTypeMg, DNSTimeseriesGroupResponseCodeParamsQueryTypeMinfo, DNSTimeseriesGroupResponseCodeParamsQueryTypeMr, DNSTimeseriesGroupResponseCodeParamsQueryTypeMX, DNSTimeseriesGroupResponseCodeParamsQueryTypeNAPTR, DNSTimeseriesGroupResponseCodeParamsQueryTypeNb, DNSTimeseriesGroupResponseCodeParamsQueryTypeNbstat, DNSTimeseriesGroupResponseCodeParamsQueryTypeNid, DNSTimeseriesGroupResponseCodeParamsQueryTypeNimloc, DNSTimeseriesGroupResponseCodeParamsQueryTypeNinfo, DNSTimeseriesGroupResponseCodeParamsQueryTypeNS, DNSTimeseriesGroupResponseCodeParamsQueryTypeNsap, DNSTimeseriesGroupResponseCodeParamsQueryTypeNsec, DNSTimeseriesGroupResponseCodeParamsQueryTypeNsec3, DNSTimeseriesGroupResponseCodeParamsQueryTypeNsec3Param, DNSTimeseriesGroupResponseCodeParamsQueryTypeNull, DNSTimeseriesGroupResponseCodeParamsQueryTypeNxt, DNSTimeseriesGroupResponseCodeParamsQueryTypeOpenpgpkey, DNSTimeseriesGroupResponseCodeParamsQueryTypeOpt, DNSTimeseriesGroupResponseCodeParamsQueryTypePTR, DNSTimeseriesGroupResponseCodeParamsQueryTypePx, DNSTimeseriesGroupResponseCodeParamsQueryTypeRkey, DNSTimeseriesGroupResponseCodeParamsQueryTypeRp, DNSTimeseriesGroupResponseCodeParamsQueryTypeRrsig, DNSTimeseriesGroupResponseCodeParamsQueryTypeRt, DNSTimeseriesGroupResponseCodeParamsQueryTypeSig, DNSTimeseriesGroupResponseCodeParamsQueryTypeSink, DNSTimeseriesGroupResponseCodeParamsQueryTypeSMIMEA, DNSTimeseriesGroupResponseCodeParamsQueryTypeSOA, DNSTimeseriesGroupResponseCodeParamsQueryTypeSPF, DNSTimeseriesGroupResponseCodeParamsQueryTypeSRV, DNSTimeseriesGroupResponseCodeParamsQueryTypeSSHFP, DNSTimeseriesGroupResponseCodeParamsQueryTypeSVCB, DNSTimeseriesGroupResponseCodeParamsQueryTypeTa, DNSTimeseriesGroupResponseCodeParamsQueryTypeTalink, DNSTimeseriesGroupResponseCodeParamsQueryTypeTkey, DNSTimeseriesGroupResponseCodeParamsQueryTypeTLSA, DNSTimeseriesGroupResponseCodeParamsQueryTypeTSIG, DNSTimeseriesGroupResponseCodeParamsQueryTypeTXT, DNSTimeseriesGroupResponseCodeParamsQueryTypeUinfo, DNSTimeseriesGroupResponseCodeParamsQueryTypeUID, DNSTimeseriesGroupResponseCodeParamsQueryTypeUnspec, DNSTimeseriesGroupResponseCodeParamsQueryTypeURI, DNSTimeseriesGroupResponseCodeParamsQueryTypeWks, DNSTimeseriesGroupResponseCodeParamsQueryTypeX25, DNSTimeseriesGroupResponseCodeParamsQueryTypeZonemd:
		return true
	}
	return false
}

type DNSTimeseriesGroupResponseCodeResponseEnvelope struct {
	Result  DNSTimeseriesGroupResponseCodeResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    dnsTimeseriesGroupResponseCodeResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupResponseCodeResponseEnvelopeJSON contains the JSON metadata
// for the struct [DNSTimeseriesGroupResponseCodeResponseEnvelope]
type dnsTimeseriesGroupResponseCodeResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupResponseCodeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupResponseCodeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesGroupResponseTTLParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[DNSTimeseriesGroupResponseTTLParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[DNSTimeseriesGroupResponseTTLParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[DNSTimeseriesGroupResponseTTLParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[DNSTimeseriesGroupResponseTTLParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[DNSTimeseriesGroupResponseTTLParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupResponseTTLParams]'s query parameters as
// `url.Values`.
func (r DNSTimeseriesGroupResponseTTLParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupResponseTTLParamsAggInterval string

const (
	DNSTimeseriesGroupResponseTTLParamsAggInterval15m DNSTimeseriesGroupResponseTTLParamsAggInterval = "15m"
	DNSTimeseriesGroupResponseTTLParamsAggInterval1h  DNSTimeseriesGroupResponseTTLParamsAggInterval = "1h"
	DNSTimeseriesGroupResponseTTLParamsAggInterval1d  DNSTimeseriesGroupResponseTTLParamsAggInterval = "1d"
	DNSTimeseriesGroupResponseTTLParamsAggInterval1w  DNSTimeseriesGroupResponseTTLParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupResponseTTLParamsAggInterval) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupResponseTTLParamsAggInterval15m, DNSTimeseriesGroupResponseTTLParamsAggInterval1h, DNSTimeseriesGroupResponseTTLParamsAggInterval1d, DNSTimeseriesGroupResponseTTLParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type DNSTimeseriesGroupResponseTTLParamsFormat string

const (
	DNSTimeseriesGroupResponseTTLParamsFormatJson DNSTimeseriesGroupResponseTTLParamsFormat = "JSON"
	DNSTimeseriesGroupResponseTTLParamsFormatCsv  DNSTimeseriesGroupResponseTTLParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupResponseTTLParamsFormat) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupResponseTTLParamsFormatJson, DNSTimeseriesGroupResponseTTLParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type DNSTimeseriesGroupResponseTTLParamsProtocol string

const (
	DNSTimeseriesGroupResponseTTLParamsProtocolUdp   DNSTimeseriesGroupResponseTTLParamsProtocol = "UDP"
	DNSTimeseriesGroupResponseTTLParamsProtocolTCP   DNSTimeseriesGroupResponseTTLParamsProtocol = "TCP"
	DNSTimeseriesGroupResponseTTLParamsProtocolHTTPS DNSTimeseriesGroupResponseTTLParamsProtocol = "HTTPS"
	DNSTimeseriesGroupResponseTTLParamsProtocolTLS   DNSTimeseriesGroupResponseTTLParamsProtocol = "TLS"
)

func (r DNSTimeseriesGroupResponseTTLParamsProtocol) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupResponseTTLParamsProtocolUdp, DNSTimeseriesGroupResponseTTLParamsProtocolTCP, DNSTimeseriesGroupResponseTTLParamsProtocolHTTPS, DNSTimeseriesGroupResponseTTLParamsProtocolTLS:
		return true
	}
	return false
}

// Filters results by DNS query type.
type DNSTimeseriesGroupResponseTTLParamsQueryType string

const (
	DNSTimeseriesGroupResponseTTLParamsQueryTypeA          DNSTimeseriesGroupResponseTTLParamsQueryType = "A"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeAAAA       DNSTimeseriesGroupResponseTTLParamsQueryType = "AAAA"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeA6         DNSTimeseriesGroupResponseTTLParamsQueryType = "A6"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeAfsdb      DNSTimeseriesGroupResponseTTLParamsQueryType = "AFSDB"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeAny        DNSTimeseriesGroupResponseTTLParamsQueryType = "ANY"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeApl        DNSTimeseriesGroupResponseTTLParamsQueryType = "APL"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeAtma       DNSTimeseriesGroupResponseTTLParamsQueryType = "ATMA"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeAXFR       DNSTimeseriesGroupResponseTTLParamsQueryType = "AXFR"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeCAA        DNSTimeseriesGroupResponseTTLParamsQueryType = "CAA"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeCdnskey    DNSTimeseriesGroupResponseTTLParamsQueryType = "CDNSKEY"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeCds        DNSTimeseriesGroupResponseTTLParamsQueryType = "CDS"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeCERT       DNSTimeseriesGroupResponseTTLParamsQueryType = "CERT"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeCNAME      DNSTimeseriesGroupResponseTTLParamsQueryType = "CNAME"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeCsync      DNSTimeseriesGroupResponseTTLParamsQueryType = "CSYNC"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeDhcid      DNSTimeseriesGroupResponseTTLParamsQueryType = "DHCID"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeDlv        DNSTimeseriesGroupResponseTTLParamsQueryType = "DLV"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeDname      DNSTimeseriesGroupResponseTTLParamsQueryType = "DNAME"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeDNSKEY     DNSTimeseriesGroupResponseTTLParamsQueryType = "DNSKEY"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeDoa        DNSTimeseriesGroupResponseTTLParamsQueryType = "DOA"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeDS         DNSTimeseriesGroupResponseTTLParamsQueryType = "DS"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeEid        DNSTimeseriesGroupResponseTTLParamsQueryType = "EID"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeEui48      DNSTimeseriesGroupResponseTTLParamsQueryType = "EUI48"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeEui64      DNSTimeseriesGroupResponseTTLParamsQueryType = "EUI64"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeGpos       DNSTimeseriesGroupResponseTTLParamsQueryType = "GPOS"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeGid        DNSTimeseriesGroupResponseTTLParamsQueryType = "GID"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeHinfo      DNSTimeseriesGroupResponseTTLParamsQueryType = "HINFO"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeHip        DNSTimeseriesGroupResponseTTLParamsQueryType = "HIP"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeHTTPS      DNSTimeseriesGroupResponseTTLParamsQueryType = "HTTPS"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeIpseckey   DNSTimeseriesGroupResponseTTLParamsQueryType = "IPSECKEY"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeIsdn       DNSTimeseriesGroupResponseTTLParamsQueryType = "ISDN"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeIxfr       DNSTimeseriesGroupResponseTTLParamsQueryType = "IXFR"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeKey        DNSTimeseriesGroupResponseTTLParamsQueryType = "KEY"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeKx         DNSTimeseriesGroupResponseTTLParamsQueryType = "KX"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeL32        DNSTimeseriesGroupResponseTTLParamsQueryType = "L32"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeL64        DNSTimeseriesGroupResponseTTLParamsQueryType = "L64"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeLOC        DNSTimeseriesGroupResponseTTLParamsQueryType = "LOC"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeLp         DNSTimeseriesGroupResponseTTLParamsQueryType = "LP"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeMaila      DNSTimeseriesGroupResponseTTLParamsQueryType = "MAILA"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeMailb      DNSTimeseriesGroupResponseTTLParamsQueryType = "MAILB"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeMB         DNSTimeseriesGroupResponseTTLParamsQueryType = "MB"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeMd         DNSTimeseriesGroupResponseTTLParamsQueryType = "MD"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeMf         DNSTimeseriesGroupResponseTTLParamsQueryType = "MF"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeMg         DNSTimeseriesGroupResponseTTLParamsQueryType = "MG"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeMinfo      DNSTimeseriesGroupResponseTTLParamsQueryType = "MINFO"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeMr         DNSTimeseriesGroupResponseTTLParamsQueryType = "MR"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeMX         DNSTimeseriesGroupResponseTTLParamsQueryType = "MX"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeNAPTR      DNSTimeseriesGroupResponseTTLParamsQueryType = "NAPTR"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeNb         DNSTimeseriesGroupResponseTTLParamsQueryType = "NB"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeNbstat     DNSTimeseriesGroupResponseTTLParamsQueryType = "NBSTAT"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeNid        DNSTimeseriesGroupResponseTTLParamsQueryType = "NID"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeNimloc     DNSTimeseriesGroupResponseTTLParamsQueryType = "NIMLOC"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeNinfo      DNSTimeseriesGroupResponseTTLParamsQueryType = "NINFO"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeNS         DNSTimeseriesGroupResponseTTLParamsQueryType = "NS"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeNsap       DNSTimeseriesGroupResponseTTLParamsQueryType = "NSAP"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeNsec       DNSTimeseriesGroupResponseTTLParamsQueryType = "NSEC"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeNsec3      DNSTimeseriesGroupResponseTTLParamsQueryType = "NSEC3"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeNsec3Param DNSTimeseriesGroupResponseTTLParamsQueryType = "NSEC3PARAM"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeNull       DNSTimeseriesGroupResponseTTLParamsQueryType = "NULL"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeNxt        DNSTimeseriesGroupResponseTTLParamsQueryType = "NXT"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeOpenpgpkey DNSTimeseriesGroupResponseTTLParamsQueryType = "OPENPGPKEY"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeOpt        DNSTimeseriesGroupResponseTTLParamsQueryType = "OPT"
	DNSTimeseriesGroupResponseTTLParamsQueryTypePTR        DNSTimeseriesGroupResponseTTLParamsQueryType = "PTR"
	DNSTimeseriesGroupResponseTTLParamsQueryTypePx         DNSTimeseriesGroupResponseTTLParamsQueryType = "PX"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeRkey       DNSTimeseriesGroupResponseTTLParamsQueryType = "RKEY"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeRp         DNSTimeseriesGroupResponseTTLParamsQueryType = "RP"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeRrsig      DNSTimeseriesGroupResponseTTLParamsQueryType = "RRSIG"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeRt         DNSTimeseriesGroupResponseTTLParamsQueryType = "RT"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeSig        DNSTimeseriesGroupResponseTTLParamsQueryType = "SIG"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeSink       DNSTimeseriesGroupResponseTTLParamsQueryType = "SINK"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeSMIMEA     DNSTimeseriesGroupResponseTTLParamsQueryType = "SMIMEA"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeSOA        DNSTimeseriesGroupResponseTTLParamsQueryType = "SOA"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeSPF        DNSTimeseriesGroupResponseTTLParamsQueryType = "SPF"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeSRV        DNSTimeseriesGroupResponseTTLParamsQueryType = "SRV"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeSSHFP      DNSTimeseriesGroupResponseTTLParamsQueryType = "SSHFP"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeSVCB       DNSTimeseriesGroupResponseTTLParamsQueryType = "SVCB"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeTa         DNSTimeseriesGroupResponseTTLParamsQueryType = "TA"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeTalink     DNSTimeseriesGroupResponseTTLParamsQueryType = "TALINK"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeTkey       DNSTimeseriesGroupResponseTTLParamsQueryType = "TKEY"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeTLSA       DNSTimeseriesGroupResponseTTLParamsQueryType = "TLSA"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeTSIG       DNSTimeseriesGroupResponseTTLParamsQueryType = "TSIG"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeTXT        DNSTimeseriesGroupResponseTTLParamsQueryType = "TXT"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeUinfo      DNSTimeseriesGroupResponseTTLParamsQueryType = "UINFO"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeUID        DNSTimeseriesGroupResponseTTLParamsQueryType = "UID"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeUnspec     DNSTimeseriesGroupResponseTTLParamsQueryType = "UNSPEC"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeURI        DNSTimeseriesGroupResponseTTLParamsQueryType = "URI"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeWks        DNSTimeseriesGroupResponseTTLParamsQueryType = "WKS"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeX25        DNSTimeseriesGroupResponseTTLParamsQueryType = "X25"
	DNSTimeseriesGroupResponseTTLParamsQueryTypeZonemd     DNSTimeseriesGroupResponseTTLParamsQueryType = "ZONEMD"
)

func (r DNSTimeseriesGroupResponseTTLParamsQueryType) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupResponseTTLParamsQueryTypeA, DNSTimeseriesGroupResponseTTLParamsQueryTypeAAAA, DNSTimeseriesGroupResponseTTLParamsQueryTypeA6, DNSTimeseriesGroupResponseTTLParamsQueryTypeAfsdb, DNSTimeseriesGroupResponseTTLParamsQueryTypeAny, DNSTimeseriesGroupResponseTTLParamsQueryTypeApl, DNSTimeseriesGroupResponseTTLParamsQueryTypeAtma, DNSTimeseriesGroupResponseTTLParamsQueryTypeAXFR, DNSTimeseriesGroupResponseTTLParamsQueryTypeCAA, DNSTimeseriesGroupResponseTTLParamsQueryTypeCdnskey, DNSTimeseriesGroupResponseTTLParamsQueryTypeCds, DNSTimeseriesGroupResponseTTLParamsQueryTypeCERT, DNSTimeseriesGroupResponseTTLParamsQueryTypeCNAME, DNSTimeseriesGroupResponseTTLParamsQueryTypeCsync, DNSTimeseriesGroupResponseTTLParamsQueryTypeDhcid, DNSTimeseriesGroupResponseTTLParamsQueryTypeDlv, DNSTimeseriesGroupResponseTTLParamsQueryTypeDname, DNSTimeseriesGroupResponseTTLParamsQueryTypeDNSKEY, DNSTimeseriesGroupResponseTTLParamsQueryTypeDoa, DNSTimeseriesGroupResponseTTLParamsQueryTypeDS, DNSTimeseriesGroupResponseTTLParamsQueryTypeEid, DNSTimeseriesGroupResponseTTLParamsQueryTypeEui48, DNSTimeseriesGroupResponseTTLParamsQueryTypeEui64, DNSTimeseriesGroupResponseTTLParamsQueryTypeGpos, DNSTimeseriesGroupResponseTTLParamsQueryTypeGid, DNSTimeseriesGroupResponseTTLParamsQueryTypeHinfo, DNSTimeseriesGroupResponseTTLParamsQueryTypeHip, DNSTimeseriesGroupResponseTTLParamsQueryTypeHTTPS, DNSTimeseriesGroupResponseTTLParamsQueryTypeIpseckey, DNSTimeseriesGroupResponseTTLParamsQueryTypeIsdn, DNSTimeseriesGroupResponseTTLParamsQueryTypeIxfr, DNSTimeseriesGroupResponseTTLParamsQueryTypeKey, DNSTimeseriesGroupResponseTTLParamsQueryTypeKx, DNSTimeseriesGroupResponseTTLParamsQueryTypeL32, DNSTimeseriesGroupResponseTTLParamsQueryTypeL64, DNSTimeseriesGroupResponseTTLParamsQueryTypeLOC, DNSTimeseriesGroupResponseTTLParamsQueryTypeLp, DNSTimeseriesGroupResponseTTLParamsQueryTypeMaila, DNSTimeseriesGroupResponseTTLParamsQueryTypeMailb, DNSTimeseriesGroupResponseTTLParamsQueryTypeMB, DNSTimeseriesGroupResponseTTLParamsQueryTypeMd, DNSTimeseriesGroupResponseTTLParamsQueryTypeMf, DNSTimeseriesGroupResponseTTLParamsQueryTypeMg, DNSTimeseriesGroupResponseTTLParamsQueryTypeMinfo, DNSTimeseriesGroupResponseTTLParamsQueryTypeMr, DNSTimeseriesGroupResponseTTLParamsQueryTypeMX, DNSTimeseriesGroupResponseTTLParamsQueryTypeNAPTR, DNSTimeseriesGroupResponseTTLParamsQueryTypeNb, DNSTimeseriesGroupResponseTTLParamsQueryTypeNbstat, DNSTimeseriesGroupResponseTTLParamsQueryTypeNid, DNSTimeseriesGroupResponseTTLParamsQueryTypeNimloc, DNSTimeseriesGroupResponseTTLParamsQueryTypeNinfo, DNSTimeseriesGroupResponseTTLParamsQueryTypeNS, DNSTimeseriesGroupResponseTTLParamsQueryTypeNsap, DNSTimeseriesGroupResponseTTLParamsQueryTypeNsec, DNSTimeseriesGroupResponseTTLParamsQueryTypeNsec3, DNSTimeseriesGroupResponseTTLParamsQueryTypeNsec3Param, DNSTimeseriesGroupResponseTTLParamsQueryTypeNull, DNSTimeseriesGroupResponseTTLParamsQueryTypeNxt, DNSTimeseriesGroupResponseTTLParamsQueryTypeOpenpgpkey, DNSTimeseriesGroupResponseTTLParamsQueryTypeOpt, DNSTimeseriesGroupResponseTTLParamsQueryTypePTR, DNSTimeseriesGroupResponseTTLParamsQueryTypePx, DNSTimeseriesGroupResponseTTLParamsQueryTypeRkey, DNSTimeseriesGroupResponseTTLParamsQueryTypeRp, DNSTimeseriesGroupResponseTTLParamsQueryTypeRrsig, DNSTimeseriesGroupResponseTTLParamsQueryTypeRt, DNSTimeseriesGroupResponseTTLParamsQueryTypeSig, DNSTimeseriesGroupResponseTTLParamsQueryTypeSink, DNSTimeseriesGroupResponseTTLParamsQueryTypeSMIMEA, DNSTimeseriesGroupResponseTTLParamsQueryTypeSOA, DNSTimeseriesGroupResponseTTLParamsQueryTypeSPF, DNSTimeseriesGroupResponseTTLParamsQueryTypeSRV, DNSTimeseriesGroupResponseTTLParamsQueryTypeSSHFP, DNSTimeseriesGroupResponseTTLParamsQueryTypeSVCB, DNSTimeseriesGroupResponseTTLParamsQueryTypeTa, DNSTimeseriesGroupResponseTTLParamsQueryTypeTalink, DNSTimeseriesGroupResponseTTLParamsQueryTypeTkey, DNSTimeseriesGroupResponseTTLParamsQueryTypeTLSA, DNSTimeseriesGroupResponseTTLParamsQueryTypeTSIG, DNSTimeseriesGroupResponseTTLParamsQueryTypeTXT, DNSTimeseriesGroupResponseTTLParamsQueryTypeUinfo, DNSTimeseriesGroupResponseTTLParamsQueryTypeUID, DNSTimeseriesGroupResponseTTLParamsQueryTypeUnspec, DNSTimeseriesGroupResponseTTLParamsQueryTypeURI, DNSTimeseriesGroupResponseTTLParamsQueryTypeWks, DNSTimeseriesGroupResponseTTLParamsQueryTypeX25, DNSTimeseriesGroupResponseTTLParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type DNSTimeseriesGroupResponseTTLParamsResponseCode string

const (
	DNSTimeseriesGroupResponseTTLParamsResponseCodeNoerror   DNSTimeseriesGroupResponseTTLParamsResponseCode = "NOERROR"
	DNSTimeseriesGroupResponseTTLParamsResponseCodeFormerr   DNSTimeseriesGroupResponseTTLParamsResponseCode = "FORMERR"
	DNSTimeseriesGroupResponseTTLParamsResponseCodeServfail  DNSTimeseriesGroupResponseTTLParamsResponseCode = "SERVFAIL"
	DNSTimeseriesGroupResponseTTLParamsResponseCodeNxdomain  DNSTimeseriesGroupResponseTTLParamsResponseCode = "NXDOMAIN"
	DNSTimeseriesGroupResponseTTLParamsResponseCodeNotimp    DNSTimeseriesGroupResponseTTLParamsResponseCode = "NOTIMP"
	DNSTimeseriesGroupResponseTTLParamsResponseCodeRefused   DNSTimeseriesGroupResponseTTLParamsResponseCode = "REFUSED"
	DNSTimeseriesGroupResponseTTLParamsResponseCodeYxdomain  DNSTimeseriesGroupResponseTTLParamsResponseCode = "YXDOMAIN"
	DNSTimeseriesGroupResponseTTLParamsResponseCodeYxrrset   DNSTimeseriesGroupResponseTTLParamsResponseCode = "YXRRSET"
	DNSTimeseriesGroupResponseTTLParamsResponseCodeNxrrset   DNSTimeseriesGroupResponseTTLParamsResponseCode = "NXRRSET"
	DNSTimeseriesGroupResponseTTLParamsResponseCodeNotauth   DNSTimeseriesGroupResponseTTLParamsResponseCode = "NOTAUTH"
	DNSTimeseriesGroupResponseTTLParamsResponseCodeNotzone   DNSTimeseriesGroupResponseTTLParamsResponseCode = "NOTZONE"
	DNSTimeseriesGroupResponseTTLParamsResponseCodeBadsig    DNSTimeseriesGroupResponseTTLParamsResponseCode = "BADSIG"
	DNSTimeseriesGroupResponseTTLParamsResponseCodeBadkey    DNSTimeseriesGroupResponseTTLParamsResponseCode = "BADKEY"
	DNSTimeseriesGroupResponseTTLParamsResponseCodeBadtime   DNSTimeseriesGroupResponseTTLParamsResponseCode = "BADTIME"
	DNSTimeseriesGroupResponseTTLParamsResponseCodeBadmode   DNSTimeseriesGroupResponseTTLParamsResponseCode = "BADMODE"
	DNSTimeseriesGroupResponseTTLParamsResponseCodeBadname   DNSTimeseriesGroupResponseTTLParamsResponseCode = "BADNAME"
	DNSTimeseriesGroupResponseTTLParamsResponseCodeBadalg    DNSTimeseriesGroupResponseTTLParamsResponseCode = "BADALG"
	DNSTimeseriesGroupResponseTTLParamsResponseCodeBadtrunc  DNSTimeseriesGroupResponseTTLParamsResponseCode = "BADTRUNC"
	DNSTimeseriesGroupResponseTTLParamsResponseCodeBadcookie DNSTimeseriesGroupResponseTTLParamsResponseCode = "BADCOOKIE"
)

func (r DNSTimeseriesGroupResponseTTLParamsResponseCode) IsKnown() bool {
	switch r {
	case DNSTimeseriesGroupResponseTTLParamsResponseCodeNoerror, DNSTimeseriesGroupResponseTTLParamsResponseCodeFormerr, DNSTimeseriesGroupResponseTTLParamsResponseCodeServfail, DNSTimeseriesGroupResponseTTLParamsResponseCodeNxdomain, DNSTimeseriesGroupResponseTTLParamsResponseCodeNotimp, DNSTimeseriesGroupResponseTTLParamsResponseCodeRefused, DNSTimeseriesGroupResponseTTLParamsResponseCodeYxdomain, DNSTimeseriesGroupResponseTTLParamsResponseCodeYxrrset, DNSTimeseriesGroupResponseTTLParamsResponseCodeNxrrset, DNSTimeseriesGroupResponseTTLParamsResponseCodeNotauth, DNSTimeseriesGroupResponseTTLParamsResponseCodeNotzone, DNSTimeseriesGroupResponseTTLParamsResponseCodeBadsig, DNSTimeseriesGroupResponseTTLParamsResponseCodeBadkey, DNSTimeseriesGroupResponseTTLParamsResponseCodeBadtime, DNSTimeseriesGroupResponseTTLParamsResponseCodeBadmode, DNSTimeseriesGroupResponseTTLParamsResponseCodeBadname, DNSTimeseriesGroupResponseTTLParamsResponseCodeBadalg, DNSTimeseriesGroupResponseTTLParamsResponseCodeBadtrunc, DNSTimeseriesGroupResponseTTLParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type DNSTimeseriesGroupResponseTTLResponseEnvelope struct {
	Result  DNSTimeseriesGroupResponseTTLResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    dnsTimeseriesGroupResponseTTLResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupResponseTTLResponseEnvelopeJSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupResponseTTLResponseEnvelope]
type dnsTimeseriesGroupResponseTTLResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupResponseTTLResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupResponseTTLResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
