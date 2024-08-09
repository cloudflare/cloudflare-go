// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
  "context"
  "net/http"
  "net/url"
  "time"

  "github.com/cloudflare/cloudflare-go/v2/internal/apijson"
  "github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
  "github.com/cloudflare/cloudflare-go/v2/internal/param"
  "github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
  "github.com/cloudflare/cloudflare-go/v2/option"
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

// Percentage breakdown of DNS queries/responses per Cache Hit status over time.
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

// Percentage breakdown of DNS responses by DNSSEC support over time.
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

// Percentage distribution of DNS queries by DNSSEC awareness over time.
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

// Percentage breakdown of DNS queries/responses per end-to-end security status
// over time.
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

// Percentage breakdown of DNS queries per IP version used over time.
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

// Percentage breakdown of DNS queries with/without matching answers over time.
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

// Percentage breakdown of DNS queries per DNS transport protocol used over time.
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

// Percentage distribution of DNS queries per query type over time.
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

// Percentage breakdown of DNS responses per response code over time.
func (r *DNSTimeseriesGroupService) ResponseCodes(ctx context.Context, query DNSTimeseriesGroupResponseCodesParams, opts ...option.RequestOption) (res *DNSTimeseriesGroupResponseCodesResponse, err error) {
  var env DNSTimeseriesGroupResponseCodesResponseEnvelope
  opts = append(r.Options[:], opts...)
  path := "radar/dns/timeseries_groups/response_codes"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

// Percentage breakdown of DNS queries per minimum answer TTL over time.
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
Meta interface{} `json:"meta,required"`
Serie0 DNSTimeseriesGroupCacheHitResponseSerie0 `json:"serie_0,required"`
JSON dnsTimeseriesGroupCacheHitResponseJSON `json:"-"`
}

// dnsTimeseriesGroupCacheHitResponseJSON contains the JSON metadata for the struct
// [DNSTimeseriesGroupCacheHitResponse]
type dnsTimeseriesGroupCacheHitResponseJSON struct {
Meta apijson.Field
Serie0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupCacheHitResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupCacheHitResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupCacheHitResponseSerie0 struct {
Negative []string `json:"NEGATIVE,required"`
Positive []string `json:"POSITIVE,required"`
JSON dnsTimeseriesGroupCacheHitResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesGroupCacheHitResponseSerie0JSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupCacheHitResponseSerie0]
type dnsTimeseriesGroupCacheHitResponseSerie0JSON struct {
Negative apijson.Field
Positive apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupCacheHitResponseSerie0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupCacheHitResponseSerie0JSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupDNSSECResponse struct {
Meta interface{} `json:"meta,required"`
Serie0 DNSTimeseriesGroupDNSSECResponseSerie0 `json:"serie_0,required"`
JSON dnsTimeseriesGroupDNSSECResponseJSON `json:"-"`
}

// dnsTimeseriesGroupDNSSECResponseJSON contains the JSON metadata for the struct
// [DNSTimeseriesGroupDNSSECResponse]
type dnsTimeseriesGroupDNSSECResponseJSON struct {
Meta apijson.Field
Serie0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupDNSSECResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupDNSSECResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupDNSSECResponseSerie0 struct {
Insecure []string `json:"INSECURE,required"`
Invalid []string `json:"INVALID,required"`
Other []string `json:"OTHER,required"`
Secure []string `json:"SECURE,required"`
JSON dnsTimeseriesGroupDNSSECResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesGroupDNSSECResponseSerie0JSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupDNSSECResponseSerie0]
type dnsTimeseriesGroupDNSSECResponseSerie0JSON struct {
Insecure apijson.Field
Invalid apijson.Field
Other apijson.Field
Secure apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupDNSSECResponseSerie0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupDNSSECResponseSerie0JSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupDNSSECAwareResponse struct {
Meta interface{} `json:"meta,required"`
Serie0 DNSTimeseriesGroupDNSSECAwareResponseSerie0 `json:"serie_0,required"`
JSON dnsTimeseriesGroupDNSSECAwareResponseJSON `json:"-"`
}

// dnsTimeseriesGroupDNSSECAwareResponseJSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupDNSSECAwareResponse]
type dnsTimeseriesGroupDNSSECAwareResponseJSON struct {
Meta apijson.Field
Serie0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupDNSSECAwareResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupDNSSECAwareResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupDNSSECAwareResponseSerie0 struct {
NotSupported []string `json:"NOT_SUPPORTED,required"`
Supported []string `json:"SUPPORTED,required"`
JSON dnsTimeseriesGroupDNSSECAwareResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesGroupDNSSECAwareResponseSerie0JSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupDNSSECAwareResponseSerie0]
type dnsTimeseriesGroupDNSSECAwareResponseSerie0JSON struct {
NotSupported apijson.Field
Supported apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupDNSSECAwareResponseSerie0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupDNSSECAwareResponseSerie0JSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupDnssece2EResponse struct {
Meta interface{} `json:"meta,required"`
Serie0 DNSTimeseriesGroupDnssece2EResponseSerie0 `json:"serie_0,required"`
JSON dnsTimeseriesGroupDnssece2EResponseJSON `json:"-"`
}

// dnsTimeseriesGroupDnssece2EResponseJSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupDnssece2EResponse]
type dnsTimeseriesGroupDnssece2EResponseJSON struct {
Meta apijson.Field
Serie0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupDnssece2EResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupDnssece2EResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupDnssece2EResponseSerie0 struct {
Negative []string `json:"NEGATIVE,required"`
Positive []string `json:"POSITIVE,required"`
JSON dnsTimeseriesGroupDnssece2EResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesGroupDnssece2EResponseSerie0JSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupDnssece2EResponseSerie0]
type dnsTimeseriesGroupDnssece2EResponseSerie0JSON struct {
Negative apijson.Field
Positive apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupDnssece2EResponseSerie0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupDnssece2EResponseSerie0JSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupIPVersionResponse struct {
Meta interface{} `json:"meta,required"`
Serie0 DNSTimeseriesGroupIPVersionResponseSerie0 `json:"serie_0,required"`
JSON dnsTimeseriesGroupIPVersionResponseJSON `json:"-"`
}

// dnsTimeseriesGroupIPVersionResponseJSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupIPVersionResponse]
type dnsTimeseriesGroupIPVersionResponseJSON struct {
Meta apijson.Field
Serie0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupIPVersionResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupIPVersionResponseSerie0 struct {
IPv4 []string `json:"IPv4,required"`
IPv6 []string `json:"IPv6,required"`
JSON dnsTimeseriesGroupIPVersionResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesGroupIPVersionResponseSerie0JSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupIPVersionResponseSerie0]
type dnsTimeseriesGroupIPVersionResponseSerie0JSON struct {
IPv4 apijson.Field
IPv6 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupIPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupIPVersionResponseSerie0JSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupMatchingAnswerResponse struct {
Meta interface{} `json:"meta,required"`
Serie0 DNSTimeseriesGroupMatchingAnswerResponseSerie0 `json:"serie_0,required"`
JSON dnsTimeseriesGroupMatchingAnswerResponseJSON `json:"-"`
}

// dnsTimeseriesGroupMatchingAnswerResponseJSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupMatchingAnswerResponse]
type dnsTimeseriesGroupMatchingAnswerResponseJSON struct {
Meta apijson.Field
Serie0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupMatchingAnswerResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupMatchingAnswerResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupMatchingAnswerResponseSerie0 struct {
Negative []string `json:"NEGATIVE,required"`
Positive []string `json:"POSITIVE,required"`
JSON dnsTimeseriesGroupMatchingAnswerResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesGroupMatchingAnswerResponseSerie0JSON contains the JSON metadata
// for the struct [DNSTimeseriesGroupMatchingAnswerResponseSerie0]
type dnsTimeseriesGroupMatchingAnswerResponseSerie0JSON struct {
Negative apijson.Field
Positive apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupMatchingAnswerResponseSerie0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupMatchingAnswerResponseSerie0JSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupProtocolResponse struct {
Meta interface{} `json:"meta,required"`
Serie0 DNSTimeseriesGroupProtocolResponseSerie0 `json:"serie_0,required"`
JSON dnsTimeseriesGroupProtocolResponseJSON `json:"-"`
}

// dnsTimeseriesGroupProtocolResponseJSON contains the JSON metadata for the struct
// [DNSTimeseriesGroupProtocolResponse]
type dnsTimeseriesGroupProtocolResponseJSON struct {
Meta apijson.Field
Serie0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupProtocolResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupProtocolResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupProtocolResponseSerie0 struct {
HTTPS []string `json:"HTTPS,required"`
TCP []string `json:"TCP,required"`
TLS []string `json:"TLS,required"`
Udp []string `json:"UDP,required"`
JSON dnsTimeseriesGroupProtocolResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesGroupProtocolResponseSerie0JSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupProtocolResponseSerie0]
type dnsTimeseriesGroupProtocolResponseSerie0JSON struct {
HTTPS apijson.Field
TCP apijson.Field
TLS apijson.Field
Udp apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupProtocolResponseSerie0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupProtocolResponseSerie0JSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupQueryTypeResponse struct {
Meta interface{} `json:"meta,required"`
Serie0 DNSTimeseriesGroupQueryTypeResponseSerie0 `json:"serie_0,required"`
JSON dnsTimeseriesGroupQueryTypeResponseJSON `json:"-"`
}

// dnsTimeseriesGroupQueryTypeResponseJSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupQueryTypeResponse]
type dnsTimeseriesGroupQueryTypeResponseJSON struct {
Meta apijson.Field
Serie0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupQueryTypeResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupQueryTypeResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupQueryTypeResponseSerie0 struct {
A []string `json:"A,required"`
AAAA []string `json:"AAAA,required"`
HTTPS []string `json:"HTTPS,required"`
NS []string `json:"NS,required"`
PTR []string `json:"PTR,required"`
JSON dnsTimeseriesGroupQueryTypeResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesGroupQueryTypeResponseSerie0JSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupQueryTypeResponseSerie0]
type dnsTimeseriesGroupQueryTypeResponseSerie0JSON struct {
A apijson.Field
AAAA apijson.Field
HTTPS apijson.Field
NS apijson.Field
PTR apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupQueryTypeResponseSerie0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupQueryTypeResponseSerie0JSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupResponseCodesResponse struct {
Meta interface{} `json:"meta,required"`
Serie0 DNSTimeseriesGroupResponseCodesResponseSerie0 `json:"serie_0,required"`
JSON dnsTimeseriesGroupResponseCodesResponseJSON `json:"-"`
}

// dnsTimeseriesGroupResponseCodesResponseJSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupResponseCodesResponse]
type dnsTimeseriesGroupResponseCodesResponseJSON struct {
Meta apijson.Field
Serie0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupResponseCodesResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupResponseCodesResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupResponseCodesResponseSerie0 struct {
Formerr []string `json:"FORMERR,required"`
Noerror []string `json:"NOERROR,required"`
Notimp []string `json:"NOTIMP,required"`
Nxdomain []string `json:"NXDOMAIN,required"`
Refused []string `json:"REFUSED,required"`
Servfail []string `json:"SERVFAIL,required"`
JSON dnsTimeseriesGroupResponseCodesResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesGroupResponseCodesResponseSerie0JSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupResponseCodesResponseSerie0]
type dnsTimeseriesGroupResponseCodesResponseSerie0JSON struct {
Formerr apijson.Field
Noerror apijson.Field
Notimp apijson.Field
Nxdomain apijson.Field
Refused apijson.Field
Servfail apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupResponseCodesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupResponseCodesResponseSerie0JSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupResponseTTLResponse struct {
Meta interface{} `json:"meta,required"`
Serie0 DNSTimeseriesGroupResponseTTLResponseSerie0 `json:"serie_0,required"`
JSON dnsTimeseriesGroupResponseTTLResponseJSON `json:"-"`
}

// dnsTimeseriesGroupResponseTTLResponseJSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupResponseTTLResponse]
type dnsTimeseriesGroupResponseTTLResponseJSON struct {
Meta apijson.Field
Serie0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupResponseTTLResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupResponseTTLResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupResponseTTLResponseSerie0 struct {
0m []string `json:"<=0m,required"`
15m []string `json:"<=15m,required"`
1d []string `json:"<=1d,required"`
1h []string `json:"<=1h,required"`
1m []string `json:"<=1m,required"`
1w []string `json:"<=1w,required"`
1y []string `json:"<=1y,required"`
5m []string `json:"<=5m,required"`
1y []string `json:">1y,required"`
JSON dnsTimeseriesGroupResponseTTLResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesGroupResponseTTLResponseSerie0JSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupResponseTTLResponseSerie0]
type dnsTimeseriesGroupResponseTTLResponseSerie0JSON struct {
0m apijson.Field
15m apijson.Field
1d apijson.Field
1h apijson.Field
1m apijson.Field
1w apijson.Field
1y apijson.Field
5m apijson.Field
1y apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupResponseTTLResponseSerie0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupResponseTTLResponseSerie0JSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupCacheHitParams struct {
// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
AggInterval param.Field[DNSTimeseriesGroupCacheHitParamsAggInterval] `query:"aggInterval"`
// Array of comma separated list of ASNs, start with `-` to exclude from results.
// For example, `-174, 3356` excludes results from AS174, but includes results from
// AS3356.
ASN param.Field[[]string] `query:"asn"`
// Array of comma separated list of continents (alpha-2 continent codes). Start
// with `-` to exclude from results. For example, `-EU,NA` excludes results from
// Europe, but includes results from North America.
Continent param.Field[[]string] `query:"continent"`
// End of the date range (inclusive).
DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
// For example, use `7d` and `7dControl` to compare this week with the previous
// week. Use this parameter or set specific start and end dates (`dateStart` and
// `dateEnd` parameters).
DateRange param.Field[[]string] `query:"dateRange"`
// Array of datetimes to filter the start of a series.
DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
// Format results are returned in.
Format param.Field[DNSTimeseriesGroupCacheHitParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupCacheHitParams]'s query parameters as
// `url.Values`.
func (r DNSTimeseriesGroupCacheHitParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupCacheHitParamsAggInterval string

const (
  DNSTimeseriesGroupCacheHitParamsAggInterval15m DNSTimeseriesGroupCacheHitParamsAggInterval = "15m"
  DNSTimeseriesGroupCacheHitParamsAggInterval1h DNSTimeseriesGroupCacheHitParamsAggInterval = "1h"
  DNSTimeseriesGroupCacheHitParamsAggInterval1d DNSTimeseriesGroupCacheHitParamsAggInterval = "1d"
  DNSTimeseriesGroupCacheHitParamsAggInterval1w DNSTimeseriesGroupCacheHitParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupCacheHitParamsAggInterval) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupCacheHitParamsAggInterval15m, DNSTimeseriesGroupCacheHitParamsAggInterval1h, DNSTimeseriesGroupCacheHitParamsAggInterval1d, DNSTimeseriesGroupCacheHitParamsAggInterval1w:
      return true
  }
  return false
}

// Format results are returned in.
type DNSTimeseriesGroupCacheHitParamsFormat string

const (
  DNSTimeseriesGroupCacheHitParamsFormatJson DNSTimeseriesGroupCacheHitParamsFormat = "JSON"
  DNSTimeseriesGroupCacheHitParamsFormatCsv DNSTimeseriesGroupCacheHitParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupCacheHitParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupCacheHitParamsFormatJson, DNSTimeseriesGroupCacheHitParamsFormatCsv:
      return true
  }
  return false
}

type DNSTimeseriesGroupCacheHitResponseEnvelope struct {
Result DNSTimeseriesGroupCacheHitResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsTimeseriesGroupCacheHitResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupCacheHitResponseEnvelopeJSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupCacheHitResponseEnvelope]
type dnsTimeseriesGroupCacheHitResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupCacheHitResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupCacheHitResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupDNSSECParams struct {
// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
AggInterval param.Field[DNSTimeseriesGroupDNSSECParamsAggInterval] `query:"aggInterval"`
// Array of comma separated list of ASNs, start with `-` to exclude from results.
// For example, `-174, 3356` excludes results from AS174, but includes results from
// AS3356.
ASN param.Field[[]string] `query:"asn"`
// Array of comma separated list of continents (alpha-2 continent codes). Start
// with `-` to exclude from results. For example, `-EU,NA` excludes results from
// Europe, but includes results from North America.
Continent param.Field[[]string] `query:"continent"`
// End of the date range (inclusive).
DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
// For example, use `7d` and `7dControl` to compare this week with the previous
// week. Use this parameter or set specific start and end dates (`dateStart` and
// `dateEnd` parameters).
DateRange param.Field[[]string] `query:"dateRange"`
// Array of datetimes to filter the start of a series.
DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
// Format results are returned in.
Format param.Field[DNSTimeseriesGroupDNSSECParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupDNSSECParams]'s query parameters as
// `url.Values`.
func (r DNSTimeseriesGroupDNSSECParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupDNSSECParamsAggInterval string

const (
  DNSTimeseriesGroupDNSSECParamsAggInterval15m DNSTimeseriesGroupDNSSECParamsAggInterval = "15m"
  DNSTimeseriesGroupDNSSECParamsAggInterval1h DNSTimeseriesGroupDNSSECParamsAggInterval = "1h"
  DNSTimeseriesGroupDNSSECParamsAggInterval1d DNSTimeseriesGroupDNSSECParamsAggInterval = "1d"
  DNSTimeseriesGroupDNSSECParamsAggInterval1w DNSTimeseriesGroupDNSSECParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupDNSSECParamsAggInterval) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupDNSSECParamsAggInterval15m, DNSTimeseriesGroupDNSSECParamsAggInterval1h, DNSTimeseriesGroupDNSSECParamsAggInterval1d, DNSTimeseriesGroupDNSSECParamsAggInterval1w:
      return true
  }
  return false
}

// Format results are returned in.
type DNSTimeseriesGroupDNSSECParamsFormat string

const (
  DNSTimeseriesGroupDNSSECParamsFormatJson DNSTimeseriesGroupDNSSECParamsFormat = "JSON"
  DNSTimeseriesGroupDNSSECParamsFormatCsv DNSTimeseriesGroupDNSSECParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupDNSSECParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupDNSSECParamsFormatJson, DNSTimeseriesGroupDNSSECParamsFormatCsv:
      return true
  }
  return false
}

type DNSTimeseriesGroupDNSSECResponseEnvelope struct {
Result DNSTimeseriesGroupDNSSECResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsTimeseriesGroupDNSSECResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupDNSSECResponseEnvelopeJSON contains the JSON metadata for the
// struct [DNSTimeseriesGroupDNSSECResponseEnvelope]
type dnsTimeseriesGroupDNSSECResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupDNSSECResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupDNSSECResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupDNSSECAwareParams struct {
// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
AggInterval param.Field[DNSTimeseriesGroupDNSSECAwareParamsAggInterval] `query:"aggInterval"`
// Array of comma separated list of ASNs, start with `-` to exclude from results.
// For example, `-174, 3356` excludes results from AS174, but includes results from
// AS3356.
ASN param.Field[[]string] `query:"asn"`
// Array of comma separated list of continents (alpha-2 continent codes). Start
// with `-` to exclude from results. For example, `-EU,NA` excludes results from
// Europe, but includes results from North America.
Continent param.Field[[]string] `query:"continent"`
// End of the date range (inclusive).
DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
// For example, use `7d` and `7dControl` to compare this week with the previous
// week. Use this parameter or set specific start and end dates (`dateStart` and
// `dateEnd` parameters).
DateRange param.Field[[]string] `query:"dateRange"`
// Array of datetimes to filter the start of a series.
DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
// Format results are returned in.
Format param.Field[DNSTimeseriesGroupDNSSECAwareParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupDNSSECAwareParams]'s query parameters as
// `url.Values`.
func (r DNSTimeseriesGroupDNSSECAwareParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupDNSSECAwareParamsAggInterval string

const (
  DNSTimeseriesGroupDNSSECAwareParamsAggInterval15m DNSTimeseriesGroupDNSSECAwareParamsAggInterval = "15m"
  DNSTimeseriesGroupDNSSECAwareParamsAggInterval1h DNSTimeseriesGroupDNSSECAwareParamsAggInterval = "1h"
  DNSTimeseriesGroupDNSSECAwareParamsAggInterval1d DNSTimeseriesGroupDNSSECAwareParamsAggInterval = "1d"
  DNSTimeseriesGroupDNSSECAwareParamsAggInterval1w DNSTimeseriesGroupDNSSECAwareParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupDNSSECAwareParamsAggInterval) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupDNSSECAwareParamsAggInterval15m, DNSTimeseriesGroupDNSSECAwareParamsAggInterval1h, DNSTimeseriesGroupDNSSECAwareParamsAggInterval1d, DNSTimeseriesGroupDNSSECAwareParamsAggInterval1w:
      return true
  }
  return false
}

// Format results are returned in.
type DNSTimeseriesGroupDNSSECAwareParamsFormat string

const (
  DNSTimeseriesGroupDNSSECAwareParamsFormatJson DNSTimeseriesGroupDNSSECAwareParamsFormat = "JSON"
  DNSTimeseriesGroupDNSSECAwareParamsFormatCsv DNSTimeseriesGroupDNSSECAwareParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupDNSSECAwareParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupDNSSECAwareParamsFormatJson, DNSTimeseriesGroupDNSSECAwareParamsFormatCsv:
      return true
  }
  return false
}

type DNSTimeseriesGroupDNSSECAwareResponseEnvelope struct {
Result DNSTimeseriesGroupDNSSECAwareResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsTimeseriesGroupDNSSECAwareResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupDNSSECAwareResponseEnvelopeJSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupDNSSECAwareResponseEnvelope]
type dnsTimeseriesGroupDNSSECAwareResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupDNSSECAwareResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupDNSSECAwareResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupDNSSECE2EParams struct {
// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
AggInterval param.Field[DNSTimeseriesGroupDnssece2EParamsAggInterval] `query:"aggInterval"`
// Array of comma separated list of ASNs, start with `-` to exclude from results.
// For example, `-174, 3356` excludes results from AS174, but includes results from
// AS3356.
ASN param.Field[[]string] `query:"asn"`
// Array of comma separated list of continents (alpha-2 continent codes). Start
// with `-` to exclude from results. For example, `-EU,NA` excludes results from
// Europe, but includes results from North America.
Continent param.Field[[]string] `query:"continent"`
// End of the date range (inclusive).
DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
// For example, use `7d` and `7dControl` to compare this week with the previous
// week. Use this parameter or set specific start and end dates (`dateStart` and
// `dateEnd` parameters).
DateRange param.Field[[]string] `query:"dateRange"`
// Array of datetimes to filter the start of a series.
DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
// Format results are returned in.
Format param.Field[DNSTimeseriesGroupDnssece2EParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupDNSSECE2EParams]'s query parameters as
// `url.Values`.
func (r DNSTimeseriesGroupDNSSECE2EParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupDnssece2EParamsAggInterval string

const (
  DNSTimeseriesGroupDnssece2EParamsAggInterval15m DNSTimeseriesGroupDnssece2EParamsAggInterval = "15m"
  DNSTimeseriesGroupDnssece2EParamsAggInterval1h DNSTimeseriesGroupDnssece2EParamsAggInterval = "1h"
  DNSTimeseriesGroupDnssece2EParamsAggInterval1d DNSTimeseriesGroupDnssece2EParamsAggInterval = "1d"
  DNSTimeseriesGroupDnssece2EParamsAggInterval1w DNSTimeseriesGroupDnssece2EParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupDnssece2EParamsAggInterval) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupDnssece2EParamsAggInterval15m, DNSTimeseriesGroupDnssece2EParamsAggInterval1h, DNSTimeseriesGroupDnssece2EParamsAggInterval1d, DNSTimeseriesGroupDnssece2EParamsAggInterval1w:
      return true
  }
  return false
}

// Format results are returned in.
type DNSTimeseriesGroupDnssece2EParamsFormat string

const (
  DNSTimeseriesGroupDnssece2EParamsFormatJson DNSTimeseriesGroupDnssece2EParamsFormat = "JSON"
  DNSTimeseriesGroupDnssece2EParamsFormatCsv DNSTimeseriesGroupDnssece2EParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupDnssece2EParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupDnssece2EParamsFormatJson, DNSTimeseriesGroupDnssece2EParamsFormatCsv:
      return true
  }
  return false
}

type DNSTimeseriesGroupDnssece2EResponseEnvelope struct {
Result DNSTimeseriesGroupDnssece2EResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsTimeseriesGroupDnssece2EResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupDnssece2EResponseEnvelopeJSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupDnssece2EResponseEnvelope]
type dnsTimeseriesGroupDnssece2EResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupDnssece2EResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupDnssece2EResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupIPVersionParams struct {
// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
AggInterval param.Field[DNSTimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
// Array of comma separated list of ASNs, start with `-` to exclude from results.
// For example, `-174, 3356` excludes results from AS174, but includes results from
// AS3356.
ASN param.Field[[]string] `query:"asn"`
// Array of comma separated list of continents (alpha-2 continent codes). Start
// with `-` to exclude from results. For example, `-EU,NA` excludes results from
// Europe, but includes results from North America.
Continent param.Field[[]string] `query:"continent"`
// End of the date range (inclusive).
DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
// For example, use `7d` and `7dControl` to compare this week with the previous
// week. Use this parameter or set specific start and end dates (`dateStart` and
// `dateEnd` parameters).
DateRange param.Field[[]string] `query:"dateRange"`
// Array of datetimes to filter the start of a series.
DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
// Format results are returned in.
Format param.Field[DNSTimeseriesGroupIPVersionParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupIPVersionParams]'s query parameters as
// `url.Values`.
func (r DNSTimeseriesGroupIPVersionParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupIPVersionParamsAggInterval string

const (
  DNSTimeseriesGroupIPVersionParamsAggInterval15m DNSTimeseriesGroupIPVersionParamsAggInterval = "15m"
  DNSTimeseriesGroupIPVersionParamsAggInterval1h DNSTimeseriesGroupIPVersionParamsAggInterval = "1h"
  DNSTimeseriesGroupIPVersionParamsAggInterval1d DNSTimeseriesGroupIPVersionParamsAggInterval = "1d"
  DNSTimeseriesGroupIPVersionParamsAggInterval1w DNSTimeseriesGroupIPVersionParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupIPVersionParamsAggInterval) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupIPVersionParamsAggInterval15m, DNSTimeseriesGroupIPVersionParamsAggInterval1h, DNSTimeseriesGroupIPVersionParamsAggInterval1d, DNSTimeseriesGroupIPVersionParamsAggInterval1w:
      return true
  }
  return false
}

// Format results are returned in.
type DNSTimeseriesGroupIPVersionParamsFormat string

const (
  DNSTimeseriesGroupIPVersionParamsFormatJson DNSTimeseriesGroupIPVersionParamsFormat = "JSON"
  DNSTimeseriesGroupIPVersionParamsFormatCsv DNSTimeseriesGroupIPVersionParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupIPVersionParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupIPVersionParamsFormatJson, DNSTimeseriesGroupIPVersionParamsFormatCsv:
      return true
  }
  return false
}

type DNSTimeseriesGroupIPVersionResponseEnvelope struct {
Result DNSTimeseriesGroupIPVersionResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsTimeseriesGroupIPVersionResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupIPVersionResponseEnvelopeJSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupIPVersionResponseEnvelope]
type dnsTimeseriesGroupIPVersionResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupIPVersionResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupMatchingAnswerParams struct {
// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
AggInterval param.Field[DNSTimeseriesGroupMatchingAnswerParamsAggInterval] `query:"aggInterval"`
// Array of comma separated list of ASNs, start with `-` to exclude from results.
// For example, `-174, 3356` excludes results from AS174, but includes results from
// AS3356.
ASN param.Field[[]string] `query:"asn"`
// Array of comma separated list of continents (alpha-2 continent codes). Start
// with `-` to exclude from results. For example, `-EU,NA` excludes results from
// Europe, but includes results from North America.
Continent param.Field[[]string] `query:"continent"`
// End of the date range (inclusive).
DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
// For example, use `7d` and `7dControl` to compare this week with the previous
// week. Use this parameter or set specific start and end dates (`dateStart` and
// `dateEnd` parameters).
DateRange param.Field[[]string] `query:"dateRange"`
// Array of datetimes to filter the start of a series.
DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
// Format results are returned in.
Format param.Field[DNSTimeseriesGroupMatchingAnswerParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupMatchingAnswerParams]'s query parameters
// as `url.Values`.
func (r DNSTimeseriesGroupMatchingAnswerParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupMatchingAnswerParamsAggInterval string

const (
  DNSTimeseriesGroupMatchingAnswerParamsAggInterval15m DNSTimeseriesGroupMatchingAnswerParamsAggInterval = "15m"
  DNSTimeseriesGroupMatchingAnswerParamsAggInterval1h DNSTimeseriesGroupMatchingAnswerParamsAggInterval = "1h"
  DNSTimeseriesGroupMatchingAnswerParamsAggInterval1d DNSTimeseriesGroupMatchingAnswerParamsAggInterval = "1d"
  DNSTimeseriesGroupMatchingAnswerParamsAggInterval1w DNSTimeseriesGroupMatchingAnswerParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupMatchingAnswerParamsAggInterval) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupMatchingAnswerParamsAggInterval15m, DNSTimeseriesGroupMatchingAnswerParamsAggInterval1h, DNSTimeseriesGroupMatchingAnswerParamsAggInterval1d, DNSTimeseriesGroupMatchingAnswerParamsAggInterval1w:
      return true
  }
  return false
}

// Format results are returned in.
type DNSTimeseriesGroupMatchingAnswerParamsFormat string

const (
  DNSTimeseriesGroupMatchingAnswerParamsFormatJson DNSTimeseriesGroupMatchingAnswerParamsFormat = "JSON"
  DNSTimeseriesGroupMatchingAnswerParamsFormatCsv DNSTimeseriesGroupMatchingAnswerParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupMatchingAnswerParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupMatchingAnswerParamsFormatJson, DNSTimeseriesGroupMatchingAnswerParamsFormatCsv:
      return true
  }
  return false
}

type DNSTimeseriesGroupMatchingAnswerResponseEnvelope struct {
Result DNSTimeseriesGroupMatchingAnswerResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsTimeseriesGroupMatchingAnswerResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupMatchingAnswerResponseEnvelopeJSON contains the JSON metadata
// for the struct [DNSTimeseriesGroupMatchingAnswerResponseEnvelope]
type dnsTimeseriesGroupMatchingAnswerResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupMatchingAnswerResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupMatchingAnswerResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupProtocolParams struct {
// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
AggInterval param.Field[DNSTimeseriesGroupProtocolParamsAggInterval] `query:"aggInterval"`
// Array of comma separated list of ASNs, start with `-` to exclude from results.
// For example, `-174, 3356` excludes results from AS174, but includes results from
// AS3356.
ASN param.Field[[]string] `query:"asn"`
// Array of comma separated list of continents (alpha-2 continent codes). Start
// with `-` to exclude from results. For example, `-EU,NA` excludes results from
// Europe, but includes results from North America.
Continent param.Field[[]string] `query:"continent"`
// End of the date range (inclusive).
DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
// For example, use `7d` and `7dControl` to compare this week with the previous
// week. Use this parameter or set specific start and end dates (`dateStart` and
// `dateEnd` parameters).
DateRange param.Field[[]string] `query:"dateRange"`
// Array of datetimes to filter the start of a series.
DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
// Format results are returned in.
Format param.Field[DNSTimeseriesGroupProtocolParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupProtocolParams]'s query parameters as
// `url.Values`.
func (r DNSTimeseriesGroupProtocolParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupProtocolParamsAggInterval string

const (
  DNSTimeseriesGroupProtocolParamsAggInterval15m DNSTimeseriesGroupProtocolParamsAggInterval = "15m"
  DNSTimeseriesGroupProtocolParamsAggInterval1h DNSTimeseriesGroupProtocolParamsAggInterval = "1h"
  DNSTimeseriesGroupProtocolParamsAggInterval1d DNSTimeseriesGroupProtocolParamsAggInterval = "1d"
  DNSTimeseriesGroupProtocolParamsAggInterval1w DNSTimeseriesGroupProtocolParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupProtocolParamsAggInterval) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupProtocolParamsAggInterval15m, DNSTimeseriesGroupProtocolParamsAggInterval1h, DNSTimeseriesGroupProtocolParamsAggInterval1d, DNSTimeseriesGroupProtocolParamsAggInterval1w:
      return true
  }
  return false
}

// Format results are returned in.
type DNSTimeseriesGroupProtocolParamsFormat string

const (
  DNSTimeseriesGroupProtocolParamsFormatJson DNSTimeseriesGroupProtocolParamsFormat = "JSON"
  DNSTimeseriesGroupProtocolParamsFormatCsv DNSTimeseriesGroupProtocolParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupProtocolParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupProtocolParamsFormatJson, DNSTimeseriesGroupProtocolParamsFormatCsv:
      return true
  }
  return false
}

type DNSTimeseriesGroupProtocolResponseEnvelope struct {
Result DNSTimeseriesGroupProtocolResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsTimeseriesGroupProtocolResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupProtocolResponseEnvelopeJSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupProtocolResponseEnvelope]
type dnsTimeseriesGroupProtocolResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupProtocolResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupQueryTypeParams struct {
// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
AggInterval param.Field[DNSTimeseriesGroupQueryTypeParamsAggInterval] `query:"aggInterval"`
// Array of comma separated list of ASNs, start with `-` to exclude from results.
// For example, `-174, 3356` excludes results from AS174, but includes results from
// AS3356.
ASN param.Field[[]string] `query:"asn"`
// Array of comma separated list of continents (alpha-2 continent codes). Start
// with `-` to exclude from results. For example, `-EU,NA` excludes results from
// Europe, but includes results from North America.
Continent param.Field[[]string] `query:"continent"`
// End of the date range (inclusive).
DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
// For example, use `7d` and `7dControl` to compare this week with the previous
// week. Use this parameter or set specific start and end dates (`dateStart` and
// `dateEnd` parameters).
DateRange param.Field[[]string] `query:"dateRange"`
// Array of datetimes to filter the start of a series.
DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
// Format results are returned in.
Format param.Field[DNSTimeseriesGroupQueryTypeParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupQueryTypeParams]'s query parameters as
// `url.Values`.
func (r DNSTimeseriesGroupQueryTypeParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupQueryTypeParamsAggInterval string

const (
  DNSTimeseriesGroupQueryTypeParamsAggInterval15m DNSTimeseriesGroupQueryTypeParamsAggInterval = "15m"
  DNSTimeseriesGroupQueryTypeParamsAggInterval1h DNSTimeseriesGroupQueryTypeParamsAggInterval = "1h"
  DNSTimeseriesGroupQueryTypeParamsAggInterval1d DNSTimeseriesGroupQueryTypeParamsAggInterval = "1d"
  DNSTimeseriesGroupQueryTypeParamsAggInterval1w DNSTimeseriesGroupQueryTypeParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupQueryTypeParamsAggInterval) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupQueryTypeParamsAggInterval15m, DNSTimeseriesGroupQueryTypeParamsAggInterval1h, DNSTimeseriesGroupQueryTypeParamsAggInterval1d, DNSTimeseriesGroupQueryTypeParamsAggInterval1w:
      return true
  }
  return false
}

// Format results are returned in.
type DNSTimeseriesGroupQueryTypeParamsFormat string

const (
  DNSTimeseriesGroupQueryTypeParamsFormatJson DNSTimeseriesGroupQueryTypeParamsFormat = "JSON"
  DNSTimeseriesGroupQueryTypeParamsFormatCsv DNSTimeseriesGroupQueryTypeParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupQueryTypeParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupQueryTypeParamsFormatJson, DNSTimeseriesGroupQueryTypeParamsFormatCsv:
      return true
  }
  return false
}

type DNSTimeseriesGroupQueryTypeResponseEnvelope struct {
Result DNSTimeseriesGroupQueryTypeResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsTimeseriesGroupQueryTypeResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupQueryTypeResponseEnvelopeJSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupQueryTypeResponseEnvelope]
type dnsTimeseriesGroupQueryTypeResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupQueryTypeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupQueryTypeResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupResponseCodesParams struct {
// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
AggInterval param.Field[DNSTimeseriesGroupResponseCodesParamsAggInterval] `query:"aggInterval"`
// Array of comma separated list of ASNs, start with `-` to exclude from results.
// For example, `-174, 3356` excludes results from AS174, but includes results from
// AS3356.
ASN param.Field[[]string] `query:"asn"`
// Array of comma separated list of continents (alpha-2 continent codes). Start
// with `-` to exclude from results. For example, `-EU,NA` excludes results from
// Europe, but includes results from North America.
Continent param.Field[[]string] `query:"continent"`
// End of the date range (inclusive).
DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
// For example, use `7d` and `7dControl` to compare this week with the previous
// week. Use this parameter or set specific start and end dates (`dateStart` and
// `dateEnd` parameters).
DateRange param.Field[[]string] `query:"dateRange"`
// Array of datetimes to filter the start of a series.
DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
// Format results are returned in.
Format param.Field[DNSTimeseriesGroupResponseCodesParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupResponseCodesParams]'s query parameters
// as `url.Values`.
func (r DNSTimeseriesGroupResponseCodesParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupResponseCodesParamsAggInterval string

const (
  DNSTimeseriesGroupResponseCodesParamsAggInterval15m DNSTimeseriesGroupResponseCodesParamsAggInterval = "15m"
  DNSTimeseriesGroupResponseCodesParamsAggInterval1h DNSTimeseriesGroupResponseCodesParamsAggInterval = "1h"
  DNSTimeseriesGroupResponseCodesParamsAggInterval1d DNSTimeseriesGroupResponseCodesParamsAggInterval = "1d"
  DNSTimeseriesGroupResponseCodesParamsAggInterval1w DNSTimeseriesGroupResponseCodesParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupResponseCodesParamsAggInterval) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupResponseCodesParamsAggInterval15m, DNSTimeseriesGroupResponseCodesParamsAggInterval1h, DNSTimeseriesGroupResponseCodesParamsAggInterval1d, DNSTimeseriesGroupResponseCodesParamsAggInterval1w:
      return true
  }
  return false
}

// Format results are returned in.
type DNSTimeseriesGroupResponseCodesParamsFormat string

const (
  DNSTimeseriesGroupResponseCodesParamsFormatJson DNSTimeseriesGroupResponseCodesParamsFormat = "JSON"
  DNSTimeseriesGroupResponseCodesParamsFormatCsv DNSTimeseriesGroupResponseCodesParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupResponseCodesParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupResponseCodesParamsFormatJson, DNSTimeseriesGroupResponseCodesParamsFormatCsv:
      return true
  }
  return false
}

type DNSTimeseriesGroupResponseCodesResponseEnvelope struct {
Result DNSTimeseriesGroupResponseCodesResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsTimeseriesGroupResponseCodesResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupResponseCodesResponseEnvelopeJSON contains the JSON metadata
// for the struct [DNSTimeseriesGroupResponseCodesResponseEnvelope]
type dnsTimeseriesGroupResponseCodesResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupResponseCodesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupResponseCodesResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

type DNSTimeseriesGroupResponseTTLParams struct {
// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
AggInterval param.Field[DNSTimeseriesGroupResponseTTLParamsAggInterval] `query:"aggInterval"`
// Array of comma separated list of ASNs, start with `-` to exclude from results.
// For example, `-174, 3356` excludes results from AS174, but includes results from
// AS3356.
ASN param.Field[[]string] `query:"asn"`
// Array of comma separated list of continents (alpha-2 continent codes). Start
// with `-` to exclude from results. For example, `-EU,NA` excludes results from
// Europe, but includes results from North America.
Continent param.Field[[]string] `query:"continent"`
// End of the date range (inclusive).
DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
// For example, use `7d` and `7dControl` to compare this week with the previous
// week. Use this parameter or set specific start and end dates (`dateStart` and
// `dateEnd` parameters).
DateRange param.Field[[]string] `query:"dateRange"`
// Array of datetimes to filter the start of a series.
DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
// Format results are returned in.
Format param.Field[DNSTimeseriesGroupResponseTTLParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesGroupResponseTTLParams]'s query parameters as
// `url.Values`.
func (r DNSTimeseriesGroupResponseTTLParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesGroupResponseTTLParamsAggInterval string

const (
  DNSTimeseriesGroupResponseTTLParamsAggInterval15m DNSTimeseriesGroupResponseTTLParamsAggInterval = "15m"
  DNSTimeseriesGroupResponseTTLParamsAggInterval1h DNSTimeseriesGroupResponseTTLParamsAggInterval = "1h"
  DNSTimeseriesGroupResponseTTLParamsAggInterval1d DNSTimeseriesGroupResponseTTLParamsAggInterval = "1d"
  DNSTimeseriesGroupResponseTTLParamsAggInterval1w DNSTimeseriesGroupResponseTTLParamsAggInterval = "1w"
)

func (r DNSTimeseriesGroupResponseTTLParamsAggInterval) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupResponseTTLParamsAggInterval15m, DNSTimeseriesGroupResponseTTLParamsAggInterval1h, DNSTimeseriesGroupResponseTTLParamsAggInterval1d, DNSTimeseriesGroupResponseTTLParamsAggInterval1w:
      return true
  }
  return false
}

// Format results are returned in.
type DNSTimeseriesGroupResponseTTLParamsFormat string

const (
  DNSTimeseriesGroupResponseTTLParamsFormatJson DNSTimeseriesGroupResponseTTLParamsFormat = "JSON"
  DNSTimeseriesGroupResponseTTLParamsFormatCsv DNSTimeseriesGroupResponseTTLParamsFormat = "CSV"
)

func (r DNSTimeseriesGroupResponseTTLParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSTimeseriesGroupResponseTTLParamsFormatJson, DNSTimeseriesGroupResponseTTLParamsFormatCsv:
      return true
  }
  return false
}

type DNSTimeseriesGroupResponseTTLResponseEnvelope struct {
Result DNSTimeseriesGroupResponseTTLResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsTimeseriesGroupResponseTTLResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesGroupResponseTTLResponseEnvelopeJSON contains the JSON metadata for
// the struct [DNSTimeseriesGroupResponseTTLResponseEnvelope]
type dnsTimeseriesGroupResponseTTLResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesGroupResponseTTLResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesGroupResponseTTLResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}
