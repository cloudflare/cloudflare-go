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

// DNSSummaryService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSSummaryService] method instead.
type DNSSummaryService struct {
Options []option.RequestOption
}

// NewDNSSummaryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDNSSummaryService(opts ...option.RequestOption) (r *DNSSummaryService) {
  r = &DNSSummaryService{}
  r.Options = opts
  return
}

// Percentage breakdown of DNS queries/responses per Cache Hit status.
func (r *DNSSummaryService) CacheHit(ctx context.Context, query DNSSummaryCacheHitParams, opts ...option.RequestOption) (res *DNSSummaryCacheHitResponse, err error) {
  var env DNSSummaryCacheHitResponseEnvelope
  opts = append(r.Options[:], opts...)
  path := "radar/dns/summary/cache_hit"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

// Percentage breakdown of DNS responses by DNSSEC support.
func (r *DNSSummaryService) DNSSEC(ctx context.Context, query DNSSummaryDNSSECParams, opts ...option.RequestOption) (res *DNSSummaryDNSSECResponse, err error) {
  var env DNSSummaryDNSSECResponseEnvelope
  opts = append(r.Options[:], opts...)
  path := "radar/dns/summary/dnssec"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

// Percentage distribution of DNS queries by DNSSEC awareness.
func (r *DNSSummaryService) DNSSECAware(ctx context.Context, query DNSSummaryDNSSECAwareParams, opts ...option.RequestOption) (res *DNSSummaryDNSSECAwareResponse, err error) {
  var env DNSSummaryDNSSECAwareResponseEnvelope
  opts = append(r.Options[:], opts...)
  path := "radar/dns/summary/dnssec_aware"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

// Percentage breakdown of DNS queries/responses per end-to-end security status.
func (r *DNSSummaryService) DNSSECE2E(ctx context.Context, query DNSSummaryDNSSECE2EParams, opts ...option.RequestOption) (res *DNSSummaryDnssece2EResponse, err error) {
  var env DNSSummaryDnssece2EResponseEnvelope
  opts = append(r.Options[:], opts...)
  path := "radar/dns/summary/dnssec_e2e"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

// Percentage breakdown of DNS queries per IP version used.
func (r *DNSSummaryService) IPVersion(ctx context.Context, query DNSSummaryIPVersionParams, opts ...option.RequestOption) (res *DNSSummaryIPVersionResponse, err error) {
  var env DNSSummaryIPVersionResponseEnvelope
  opts = append(r.Options[:], opts...)
  path := "radar/dns/summary/ip_version"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

// Percentage breakdown of DNS queries with/without matching answers.
func (r *DNSSummaryService) MatchingAnswer(ctx context.Context, query DNSSummaryMatchingAnswerParams, opts ...option.RequestOption) (res *DNSSummaryMatchingAnswerResponse, err error) {
  var env DNSSummaryMatchingAnswerResponseEnvelope
  opts = append(r.Options[:], opts...)
  path := "radar/dns/summary/matching_answer"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

// Percentage breakdown of DNS queries per DNS transport protocol used.
func (r *DNSSummaryService) Protocol(ctx context.Context, query DNSSummaryProtocolParams, opts ...option.RequestOption) (res *DNSSummaryProtocolResponse, err error) {
  var env DNSSummaryProtocolResponseEnvelope
  opts = append(r.Options[:], opts...)
  path := "radar/dns/summary/protocol"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

// Percentage distribution of DNS queries per query type.
func (r *DNSSummaryService) QueryType(ctx context.Context, query DNSSummaryQueryTypeParams, opts ...option.RequestOption) (res *DNSSummaryQueryTypeResponse, err error) {
  var env DNSSummaryQueryTypeResponseEnvelope
  opts = append(r.Options[:], opts...)
  path := "radar/dns/summary/query_type"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

// Percentage breakdown of DNS responses per response code.
func (r *DNSSummaryService) ResponseCodes(ctx context.Context, query DNSSummaryResponseCodesParams, opts ...option.RequestOption) (res *DNSSummaryResponseCodesResponse, err error) {
  var env DNSSummaryResponseCodesResponseEnvelope
  opts = append(r.Options[:], opts...)
  path := "radar/dns/summary/response_codes"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

// Percentage breakdown of DNS queries per minimum answer TTL.
func (r *DNSSummaryService) ResponseTTL(ctx context.Context, query DNSSummaryResponseTTLParams, opts ...option.RequestOption) (res *DNSSummaryResponseTTLResponse, err error) {
  var env DNSSummaryResponseTTLResponseEnvelope
  opts = append(r.Options[:], opts...)
  path := "radar/dns/summary/response_ttl"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

type DNSSummaryCacheHitResponse struct {
Meta DNSSummaryCacheHitResponseMeta `json:"meta,required"`
Summary0 DNSSummaryCacheHitResponseSummary0 `json:"summary_0,required"`
JSON dnsSummaryCacheHitResponseJSON `json:"-"`
}

// dnsSummaryCacheHitResponseJSON contains the JSON metadata for the struct
// [DNSSummaryCacheHitResponse]
type dnsSummaryCacheHitResponseJSON struct {
Meta apijson.Field
Summary0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryCacheHitResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryCacheHitResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryCacheHitResponseMeta struct {
DateRange []DNSSummaryCacheHitResponseMetaDateRange `json:"dateRange,required"`
LastUpdated string `json:"lastUpdated,required"`
Normalization string `json:"normalization,required"`
ConfidenceInfo DNSSummaryCacheHitResponseMetaConfidenceInfo `json:"confidenceInfo"`
JSON dnsSummaryCacheHitResponseMetaJSON `json:"-"`
}

// dnsSummaryCacheHitResponseMetaJSON contains the JSON metadata for the struct
// [DNSSummaryCacheHitResponseMeta]
type dnsSummaryCacheHitResponseMetaJSON struct {
DateRange apijson.Field
LastUpdated apijson.Field
Normalization apijson.Field
ConfidenceInfo apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryCacheHitResponseMeta) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryCacheHitResponseMetaJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryCacheHitResponseMetaDateRange struct {
// Adjusted end of date range.
EndTime time.Time `json:"endTime,required" format:"date-time"`
// Adjusted start of date range.
StartTime time.Time `json:"startTime,required" format:"date-time"`
JSON dnsSummaryCacheHitResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryCacheHitResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [DNSSummaryCacheHitResponseMetaDateRange]
type dnsSummaryCacheHitResponseMetaDateRangeJSON struct {
EndTime apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryCacheHitResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryCacheHitResponseMetaDateRangeJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryCacheHitResponseMetaConfidenceInfo struct {
Annotations []DNSSummaryCacheHitResponseMetaConfidenceInfoAnnotation `json:"annotations"`
Level int64 `json:"level"`
JSON dnsSummaryCacheHitResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryCacheHitResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [DNSSummaryCacheHitResponseMetaConfidenceInfo]
type dnsSummaryCacheHitResponseMetaConfidenceInfoJSON struct {
Annotations apijson.Field
Level apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryCacheHitResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryCacheHitResponseMetaConfidenceInfoJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryCacheHitResponseMetaConfidenceInfoAnnotation struct {
DataSource string `json:"dataSource,required"`
Description string `json:"description,required"`
EventType string `json:"eventType,required"`
IsInstantaneous bool `json:"isInstantaneous,required"`
EndTime time.Time `json:"endTime" format:"date-time"`
LinkedURL string `json:"linkedUrl"`
StartTime time.Time `json:"startTime" format:"date-time"`
JSON dnsSummaryCacheHitResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryCacheHitResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [DNSSummaryCacheHitResponseMetaConfidenceInfoAnnotation]
type dnsSummaryCacheHitResponseMetaConfidenceInfoAnnotationJSON struct {
DataSource apijson.Field
Description apijson.Field
EventType apijson.Field
IsInstantaneous apijson.Field
EndTime apijson.Field
LinkedURL apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryCacheHitResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryCacheHitResponseMetaConfidenceInfoAnnotationJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryCacheHitResponseSummary0 struct {
Negative string `json:"NEGATIVE,required"`
Positive string `json:"POSITIVE,required"`
JSON dnsSummaryCacheHitResponseSummary0JSON `json:"-"`
}

// dnsSummaryCacheHitResponseSummary0JSON contains the JSON metadata for the struct
// [DNSSummaryCacheHitResponseSummary0]
type dnsSummaryCacheHitResponseSummary0JSON struct {
Negative apijson.Field
Positive apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryCacheHitResponseSummary0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryCacheHitResponseSummary0JSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDNSSECResponse struct {
Meta DNSSummaryDNSSECResponseMeta `json:"meta,required"`
Summary0 DNSSummaryDNSSECResponseSummary0 `json:"summary_0,required"`
JSON dnsSummaryDNSSECResponseJSON `json:"-"`
}

// dnsSummaryDNSSECResponseJSON contains the JSON metadata for the struct
// [DNSSummaryDNSSECResponse]
type dnsSummaryDNSSECResponseJSON struct {
Meta apijson.Field
Summary0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDNSSECResponseMeta struct {
DateRange []DNSSummaryDNSSECResponseMetaDateRange `json:"dateRange,required"`
LastUpdated string `json:"lastUpdated,required"`
Normalization string `json:"normalization,required"`
ConfidenceInfo DNSSummaryDNSSECResponseMetaConfidenceInfo `json:"confidenceInfo"`
JSON dnsSummaryDNSSECResponseMetaJSON `json:"-"`
}

// dnsSummaryDNSSECResponseMetaJSON contains the JSON metadata for the struct
// [DNSSummaryDNSSECResponseMeta]
type dnsSummaryDNSSECResponseMetaJSON struct {
DateRange apijson.Field
LastUpdated apijson.Field
Normalization apijson.Field
ConfidenceInfo apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECResponseMeta) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECResponseMetaJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDNSSECResponseMetaDateRange struct {
// Adjusted end of date range.
EndTime time.Time `json:"endTime,required" format:"date-time"`
// Adjusted start of date range.
StartTime time.Time `json:"startTime,required" format:"date-time"`
JSON dnsSummaryDNSSECResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryDNSSECResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [DNSSummaryDNSSECResponseMetaDateRange]
type dnsSummaryDNSSECResponseMetaDateRangeJSON struct {
EndTime apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECResponseMetaDateRangeJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDNSSECResponseMetaConfidenceInfo struct {
Annotations []DNSSummaryDNSSECResponseMetaConfidenceInfoAnnotation `json:"annotations"`
Level int64 `json:"level"`
JSON dnsSummaryDNSSECResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryDNSSECResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [DNSSummaryDNSSECResponseMetaConfidenceInfo]
type dnsSummaryDNSSECResponseMetaConfidenceInfoJSON struct {
Annotations apijson.Field
Level apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECResponseMetaConfidenceInfoJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDNSSECResponseMetaConfidenceInfoAnnotation struct {
DataSource string `json:"dataSource,required"`
Description string `json:"description,required"`
EventType string `json:"eventType,required"`
IsInstantaneous bool `json:"isInstantaneous,required"`
EndTime time.Time `json:"endTime" format:"date-time"`
LinkedURL string `json:"linkedUrl"`
StartTime time.Time `json:"startTime" format:"date-time"`
JSON dnsSummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [DNSSummaryDNSSECResponseMetaConfidenceInfoAnnotation]
type dnsSummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON struct {
DataSource apijson.Field
Description apijson.Field
EventType apijson.Field
IsInstantaneous apijson.Field
EndTime apijson.Field
LinkedURL apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECResponseMetaConfidenceInfoAnnotationJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDNSSECResponseSummary0 struct {
Insecure string `json:"INSECURE,required"`
Invalid string `json:"INVALID,required"`
Other string `json:"OTHER,required"`
Secure string `json:"SECURE,required"`
JSON dnsSummaryDNSSECResponseSummary0JSON `json:"-"`
}

// dnsSummaryDNSSECResponseSummary0JSON contains the JSON metadata for the struct
// [DNSSummaryDNSSECResponseSummary0]
type dnsSummaryDNSSECResponseSummary0JSON struct {
Insecure apijson.Field
Invalid apijson.Field
Other apijson.Field
Secure apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECResponseSummary0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECResponseSummary0JSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDNSSECAwareResponse struct {
Meta DNSSummaryDNSSECAwareResponseMeta `json:"meta,required"`
Summary0 DNSSummaryDNSSECAwareResponseSummary0 `json:"summary_0,required"`
JSON dnsSummaryDNSSECAwareResponseJSON `json:"-"`
}

// dnsSummaryDNSSECAwareResponseJSON contains the JSON metadata for the struct
// [DNSSummaryDNSSECAwareResponse]
type dnsSummaryDNSSECAwareResponseJSON struct {
Meta apijson.Field
Summary0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECAwareResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECAwareResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDNSSECAwareResponseMeta struct {
DateRange []DNSSummaryDNSSECAwareResponseMetaDateRange `json:"dateRange,required"`
LastUpdated string `json:"lastUpdated,required"`
Normalization string `json:"normalization,required"`
ConfidenceInfo DNSSummaryDNSSECAwareResponseMetaConfidenceInfo `json:"confidenceInfo"`
JSON dnsSummaryDNSSECAwareResponseMetaJSON `json:"-"`
}

// dnsSummaryDNSSECAwareResponseMetaJSON contains the JSON metadata for the struct
// [DNSSummaryDNSSECAwareResponseMeta]
type dnsSummaryDNSSECAwareResponseMetaJSON struct {
DateRange apijson.Field
LastUpdated apijson.Field
Normalization apijson.Field
ConfidenceInfo apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECAwareResponseMeta) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECAwareResponseMetaJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDNSSECAwareResponseMetaDateRange struct {
// Adjusted end of date range.
EndTime time.Time `json:"endTime,required" format:"date-time"`
// Adjusted start of date range.
StartTime time.Time `json:"startTime,required" format:"date-time"`
JSON dnsSummaryDNSSECAwareResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryDNSSECAwareResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [DNSSummaryDNSSECAwareResponseMetaDateRange]
type dnsSummaryDNSSECAwareResponseMetaDateRangeJSON struct {
EndTime apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECAwareResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECAwareResponseMetaDateRangeJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDNSSECAwareResponseMetaConfidenceInfo struct {
Annotations []DNSSummaryDNSSECAwareResponseMetaConfidenceInfoAnnotation `json:"annotations"`
Level int64 `json:"level"`
JSON dnsSummaryDNSSECAwareResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryDNSSECAwareResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [DNSSummaryDNSSECAwareResponseMetaConfidenceInfo]
type dnsSummaryDNSSECAwareResponseMetaConfidenceInfoJSON struct {
Annotations apijson.Field
Level apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECAwareResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECAwareResponseMetaConfidenceInfoJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDNSSECAwareResponseMetaConfidenceInfoAnnotation struct {
DataSource string `json:"dataSource,required"`
Description string `json:"description,required"`
EventType string `json:"eventType,required"`
IsInstantaneous bool `json:"isInstantaneous,required"`
EndTime time.Time `json:"endTime" format:"date-time"`
LinkedURL string `json:"linkedUrl"`
StartTime time.Time `json:"startTime" format:"date-time"`
JSON dnsSummaryDNSSECAwareResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryDNSSECAwareResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [DNSSummaryDNSSECAwareResponseMetaConfidenceInfoAnnotation]
type dnsSummaryDNSSECAwareResponseMetaConfidenceInfoAnnotationJSON struct {
DataSource apijson.Field
Description apijson.Field
EventType apijson.Field
IsInstantaneous apijson.Field
EndTime apijson.Field
LinkedURL apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECAwareResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECAwareResponseMetaConfidenceInfoAnnotationJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDNSSECAwareResponseSummary0 struct {
NotSupported string `json:"NOT_SUPPORTED,required"`
Supported string `json:"SUPPORTED,required"`
JSON dnsSummaryDNSSECAwareResponseSummary0JSON `json:"-"`
}

// dnsSummaryDNSSECAwareResponseSummary0JSON contains the JSON metadata for the
// struct [DNSSummaryDNSSECAwareResponseSummary0]
type dnsSummaryDNSSECAwareResponseSummary0JSON struct {
NotSupported apijson.Field
Supported apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECAwareResponseSummary0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECAwareResponseSummary0JSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDnssece2EResponse struct {
Meta DNSSummaryDnssece2EResponseMeta `json:"meta,required"`
Summary0 DNSSummaryDnssece2EResponseSummary0 `json:"summary_0,required"`
JSON dnsSummaryDnssece2EResponseJSON `json:"-"`
}

// dnsSummaryDnssece2EResponseJSON contains the JSON metadata for the struct
// [DNSSummaryDnssece2EResponse]
type dnsSummaryDnssece2EResponseJSON struct {
Meta apijson.Field
Summary0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDnssece2EResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDnssece2EResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDnssece2EResponseMeta struct {
DateRange []DNSSummaryDnssece2EResponseMetaDateRange `json:"dateRange,required"`
LastUpdated string `json:"lastUpdated,required"`
Normalization string `json:"normalization,required"`
ConfidenceInfo DNSSummaryDnssece2EResponseMetaConfidenceInfo `json:"confidenceInfo"`
JSON dnsSummaryDnssece2EResponseMetaJSON `json:"-"`
}

// dnsSummaryDnssece2EResponseMetaJSON contains the JSON metadata for the struct
// [DNSSummaryDnssece2EResponseMeta]
type dnsSummaryDnssece2EResponseMetaJSON struct {
DateRange apijson.Field
LastUpdated apijson.Field
Normalization apijson.Field
ConfidenceInfo apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDnssece2EResponseMeta) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDnssece2EResponseMetaJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDnssece2EResponseMetaDateRange struct {
// Adjusted end of date range.
EndTime time.Time `json:"endTime,required" format:"date-time"`
// Adjusted start of date range.
StartTime time.Time `json:"startTime,required" format:"date-time"`
JSON dnsSummaryDnssece2EResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryDnssece2EResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [DNSSummaryDnssece2EResponseMetaDateRange]
type dnsSummaryDnssece2EResponseMetaDateRangeJSON struct {
EndTime apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDnssece2EResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDnssece2EResponseMetaDateRangeJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDnssece2EResponseMetaConfidenceInfo struct {
Annotations []DNSSummaryDnssece2EResponseMetaConfidenceInfoAnnotation `json:"annotations"`
Level int64 `json:"level"`
JSON dnsSummaryDnssece2EResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryDnssece2EResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [DNSSummaryDnssece2EResponseMetaConfidenceInfo]
type dnsSummaryDnssece2EResponseMetaConfidenceInfoJSON struct {
Annotations apijson.Field
Level apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDnssece2EResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDnssece2EResponseMetaConfidenceInfoJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDnssece2EResponseMetaConfidenceInfoAnnotation struct {
DataSource string `json:"dataSource,required"`
Description string `json:"description,required"`
EventType string `json:"eventType,required"`
IsInstantaneous bool `json:"isInstantaneous,required"`
EndTime time.Time `json:"endTime" format:"date-time"`
LinkedURL string `json:"linkedUrl"`
StartTime time.Time `json:"startTime" format:"date-time"`
JSON dnsSummaryDnssece2EResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryDnssece2EResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [DNSSummaryDnssece2EResponseMetaConfidenceInfoAnnotation]
type dnsSummaryDnssece2EResponseMetaConfidenceInfoAnnotationJSON struct {
DataSource apijson.Field
Description apijson.Field
EventType apijson.Field
IsInstantaneous apijson.Field
EndTime apijson.Field
LinkedURL apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDnssece2EResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDnssece2EResponseMetaConfidenceInfoAnnotationJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDnssece2EResponseSummary0 struct {
Negative string `json:"NEGATIVE,required"`
Positive string `json:"POSITIVE,required"`
JSON dnsSummaryDnssece2EResponseSummary0JSON `json:"-"`
}

// dnsSummaryDnssece2EResponseSummary0JSON contains the JSON metadata for the
// struct [DNSSummaryDnssece2EResponseSummary0]
type dnsSummaryDnssece2EResponseSummary0JSON struct {
Negative apijson.Field
Positive apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDnssece2EResponseSummary0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDnssece2EResponseSummary0JSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryIPVersionResponse struct {
Meta DNSSummaryIPVersionResponseMeta `json:"meta,required"`
Summary0 DNSSummaryIPVersionResponseSummary0 `json:"summary_0,required"`
JSON dnsSummaryIPVersionResponseJSON `json:"-"`
}

// dnsSummaryIPVersionResponseJSON contains the JSON metadata for the struct
// [DNSSummaryIPVersionResponse]
type dnsSummaryIPVersionResponseJSON struct {
Meta apijson.Field
Summary0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryIPVersionResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryIPVersionResponseMeta struct {
DateRange []DNSSummaryIPVersionResponseMetaDateRange `json:"dateRange,required"`
LastUpdated string `json:"lastUpdated,required"`
Normalization string `json:"normalization,required"`
ConfidenceInfo DNSSummaryIPVersionResponseMetaConfidenceInfo `json:"confidenceInfo"`
JSON dnsSummaryIPVersionResponseMetaJSON `json:"-"`
}

// dnsSummaryIPVersionResponseMetaJSON contains the JSON metadata for the struct
// [DNSSummaryIPVersionResponseMeta]
type dnsSummaryIPVersionResponseMetaJSON struct {
DateRange apijson.Field
LastUpdated apijson.Field
Normalization apijson.Field
ConfidenceInfo apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryIPVersionResponseMeta) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryIPVersionResponseMetaJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryIPVersionResponseMetaDateRange struct {
// Adjusted end of date range.
EndTime time.Time `json:"endTime,required" format:"date-time"`
// Adjusted start of date range.
StartTime time.Time `json:"startTime,required" format:"date-time"`
JSON dnsSummaryIPVersionResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryIPVersionResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [DNSSummaryIPVersionResponseMetaDateRange]
type dnsSummaryIPVersionResponseMetaDateRangeJSON struct {
EndTime apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryIPVersionResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryIPVersionResponseMetaDateRangeJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryIPVersionResponseMetaConfidenceInfo struct {
Annotations []DNSSummaryIPVersionResponseMetaConfidenceInfoAnnotation `json:"annotations"`
Level int64 `json:"level"`
JSON dnsSummaryIPVersionResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryIPVersionResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [DNSSummaryIPVersionResponseMetaConfidenceInfo]
type dnsSummaryIPVersionResponseMetaConfidenceInfoJSON struct {
Annotations apijson.Field
Level apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryIPVersionResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryIPVersionResponseMetaConfidenceInfoJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryIPVersionResponseMetaConfidenceInfoAnnotation struct {
DataSource string `json:"dataSource,required"`
Description string `json:"description,required"`
EventType string `json:"eventType,required"`
IsInstantaneous bool `json:"isInstantaneous,required"`
EndTime time.Time `json:"endTime" format:"date-time"`
LinkedURL string `json:"linkedUrl"`
StartTime time.Time `json:"startTime" format:"date-time"`
JSON dnsSummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [DNSSummaryIPVersionResponseMetaConfidenceInfoAnnotation]
type dnsSummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON struct {
DataSource apijson.Field
Description apijson.Field
EventType apijson.Field
IsInstantaneous apijson.Field
EndTime apijson.Field
LinkedURL apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryIPVersionResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryIPVersionResponseMetaConfidenceInfoAnnotationJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryIPVersionResponseSummary0 struct {
IPv4 string `json:"IPv4,required"`
IPv6 string `json:"IPv6,required"`
JSON dnsSummaryIPVersionResponseSummary0JSON `json:"-"`
}

// dnsSummaryIPVersionResponseSummary0JSON contains the JSON metadata for the
// struct [DNSSummaryIPVersionResponseSummary0]
type dnsSummaryIPVersionResponseSummary0JSON struct {
IPv4 apijson.Field
IPv6 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryIPVersionResponseSummary0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryIPVersionResponseSummary0JSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryMatchingAnswerResponse struct {
Meta DNSSummaryMatchingAnswerResponseMeta `json:"meta,required"`
Summary0 DNSSummaryMatchingAnswerResponseSummary0 `json:"summary_0,required"`
JSON dnsSummaryMatchingAnswerResponseJSON `json:"-"`
}

// dnsSummaryMatchingAnswerResponseJSON contains the JSON metadata for the struct
// [DNSSummaryMatchingAnswerResponse]
type dnsSummaryMatchingAnswerResponseJSON struct {
Meta apijson.Field
Summary0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryMatchingAnswerResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryMatchingAnswerResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryMatchingAnswerResponseMeta struct {
DateRange []DNSSummaryMatchingAnswerResponseMetaDateRange `json:"dateRange,required"`
LastUpdated string `json:"lastUpdated,required"`
Normalization string `json:"normalization,required"`
ConfidenceInfo DNSSummaryMatchingAnswerResponseMetaConfidenceInfo `json:"confidenceInfo"`
JSON dnsSummaryMatchingAnswerResponseMetaJSON `json:"-"`
}

// dnsSummaryMatchingAnswerResponseMetaJSON contains the JSON metadata for the
// struct [DNSSummaryMatchingAnswerResponseMeta]
type dnsSummaryMatchingAnswerResponseMetaJSON struct {
DateRange apijson.Field
LastUpdated apijson.Field
Normalization apijson.Field
ConfidenceInfo apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryMatchingAnswerResponseMeta) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryMatchingAnswerResponseMetaJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryMatchingAnswerResponseMetaDateRange struct {
// Adjusted end of date range.
EndTime time.Time `json:"endTime,required" format:"date-time"`
// Adjusted start of date range.
StartTime time.Time `json:"startTime,required" format:"date-time"`
JSON dnsSummaryMatchingAnswerResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryMatchingAnswerResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [DNSSummaryMatchingAnswerResponseMetaDateRange]
type dnsSummaryMatchingAnswerResponseMetaDateRangeJSON struct {
EndTime apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryMatchingAnswerResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryMatchingAnswerResponseMetaDateRangeJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryMatchingAnswerResponseMetaConfidenceInfo struct {
Annotations []DNSSummaryMatchingAnswerResponseMetaConfidenceInfoAnnotation `json:"annotations"`
Level int64 `json:"level"`
JSON dnsSummaryMatchingAnswerResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryMatchingAnswerResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [DNSSummaryMatchingAnswerResponseMetaConfidenceInfo]
type dnsSummaryMatchingAnswerResponseMetaConfidenceInfoJSON struct {
Annotations apijson.Field
Level apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryMatchingAnswerResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryMatchingAnswerResponseMetaConfidenceInfoJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryMatchingAnswerResponseMetaConfidenceInfoAnnotation struct {
DataSource string `json:"dataSource,required"`
Description string `json:"description,required"`
EventType string `json:"eventType,required"`
IsInstantaneous bool `json:"isInstantaneous,required"`
EndTime time.Time `json:"endTime" format:"date-time"`
LinkedURL string `json:"linkedUrl"`
StartTime time.Time `json:"startTime" format:"date-time"`
JSON dnsSummaryMatchingAnswerResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryMatchingAnswerResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [DNSSummaryMatchingAnswerResponseMetaConfidenceInfoAnnotation]
type dnsSummaryMatchingAnswerResponseMetaConfidenceInfoAnnotationJSON struct {
DataSource apijson.Field
Description apijson.Field
EventType apijson.Field
IsInstantaneous apijson.Field
EndTime apijson.Field
LinkedURL apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryMatchingAnswerResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryMatchingAnswerResponseMetaConfidenceInfoAnnotationJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryMatchingAnswerResponseSummary0 struct {
Negative string `json:"NEGATIVE,required"`
Positive string `json:"POSITIVE,required"`
JSON dnsSummaryMatchingAnswerResponseSummary0JSON `json:"-"`
}

// dnsSummaryMatchingAnswerResponseSummary0JSON contains the JSON metadata for the
// struct [DNSSummaryMatchingAnswerResponseSummary0]
type dnsSummaryMatchingAnswerResponseSummary0JSON struct {
Negative apijson.Field
Positive apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryMatchingAnswerResponseSummary0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryMatchingAnswerResponseSummary0JSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryProtocolResponse struct {
Meta DNSSummaryProtocolResponseMeta `json:"meta,required"`
Summary0 DNSSummaryProtocolResponseSummary0 `json:"summary_0,required"`
JSON dnsSummaryProtocolResponseJSON `json:"-"`
}

// dnsSummaryProtocolResponseJSON contains the JSON metadata for the struct
// [DNSSummaryProtocolResponse]
type dnsSummaryProtocolResponseJSON struct {
Meta apijson.Field
Summary0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryProtocolResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryProtocolResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryProtocolResponseMeta struct {
DateRange []DNSSummaryProtocolResponseMetaDateRange `json:"dateRange,required"`
LastUpdated string `json:"lastUpdated,required"`
Normalization string `json:"normalization,required"`
ConfidenceInfo DNSSummaryProtocolResponseMetaConfidenceInfo `json:"confidenceInfo"`
JSON dnsSummaryProtocolResponseMetaJSON `json:"-"`
}

// dnsSummaryProtocolResponseMetaJSON contains the JSON metadata for the struct
// [DNSSummaryProtocolResponseMeta]
type dnsSummaryProtocolResponseMetaJSON struct {
DateRange apijson.Field
LastUpdated apijson.Field
Normalization apijson.Field
ConfidenceInfo apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryProtocolResponseMeta) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryProtocolResponseMetaJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryProtocolResponseMetaDateRange struct {
// Adjusted end of date range.
EndTime time.Time `json:"endTime,required" format:"date-time"`
// Adjusted start of date range.
StartTime time.Time `json:"startTime,required" format:"date-time"`
JSON dnsSummaryProtocolResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryProtocolResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [DNSSummaryProtocolResponseMetaDateRange]
type dnsSummaryProtocolResponseMetaDateRangeJSON struct {
EndTime apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryProtocolResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryProtocolResponseMetaDateRangeJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryProtocolResponseMetaConfidenceInfo struct {
Annotations []DNSSummaryProtocolResponseMetaConfidenceInfoAnnotation `json:"annotations"`
Level int64 `json:"level"`
JSON dnsSummaryProtocolResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryProtocolResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [DNSSummaryProtocolResponseMetaConfidenceInfo]
type dnsSummaryProtocolResponseMetaConfidenceInfoJSON struct {
Annotations apijson.Field
Level apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryProtocolResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryProtocolResponseMetaConfidenceInfoJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryProtocolResponseMetaConfidenceInfoAnnotation struct {
DataSource string `json:"dataSource,required"`
Description string `json:"description,required"`
EventType string `json:"eventType,required"`
IsInstantaneous bool `json:"isInstantaneous,required"`
EndTime time.Time `json:"endTime" format:"date-time"`
LinkedURL string `json:"linkedUrl"`
StartTime time.Time `json:"startTime" format:"date-time"`
JSON dnsSummaryProtocolResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryProtocolResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [DNSSummaryProtocolResponseMetaConfidenceInfoAnnotation]
type dnsSummaryProtocolResponseMetaConfidenceInfoAnnotationJSON struct {
DataSource apijson.Field
Description apijson.Field
EventType apijson.Field
IsInstantaneous apijson.Field
EndTime apijson.Field
LinkedURL apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryProtocolResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryProtocolResponseMetaConfidenceInfoAnnotationJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryProtocolResponseSummary0 struct {
HTTPS string `json:"HTTPS,required"`
TCP string `json:"TCP,required"`
TLS string `json:"TLS,required"`
Udp string `json:"UDP,required"`
JSON dnsSummaryProtocolResponseSummary0JSON `json:"-"`
}

// dnsSummaryProtocolResponseSummary0JSON contains the JSON metadata for the struct
// [DNSSummaryProtocolResponseSummary0]
type dnsSummaryProtocolResponseSummary0JSON struct {
HTTPS apijson.Field
TCP apijson.Field
TLS apijson.Field
Udp apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryProtocolResponseSummary0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryProtocolResponseSummary0JSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryQueryTypeResponse struct {
Meta DNSSummaryQueryTypeResponseMeta `json:"meta,required"`
Summary0 DNSSummaryQueryTypeResponseSummary0 `json:"summary_0,required"`
JSON dnsSummaryQueryTypeResponseJSON `json:"-"`
}

// dnsSummaryQueryTypeResponseJSON contains the JSON metadata for the struct
// [DNSSummaryQueryTypeResponse]
type dnsSummaryQueryTypeResponseJSON struct {
Meta apijson.Field
Summary0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryQueryTypeResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryQueryTypeResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryQueryTypeResponseMeta struct {
DateRange []DNSSummaryQueryTypeResponseMetaDateRange `json:"dateRange,required"`
LastUpdated string `json:"lastUpdated,required"`
Normalization string `json:"normalization,required"`
ConfidenceInfo DNSSummaryQueryTypeResponseMetaConfidenceInfo `json:"confidenceInfo"`
JSON dnsSummaryQueryTypeResponseMetaJSON `json:"-"`
}

// dnsSummaryQueryTypeResponseMetaJSON contains the JSON metadata for the struct
// [DNSSummaryQueryTypeResponseMeta]
type dnsSummaryQueryTypeResponseMetaJSON struct {
DateRange apijson.Field
LastUpdated apijson.Field
Normalization apijson.Field
ConfidenceInfo apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryQueryTypeResponseMeta) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryQueryTypeResponseMetaJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryQueryTypeResponseMetaDateRange struct {
// Adjusted end of date range.
EndTime time.Time `json:"endTime,required" format:"date-time"`
// Adjusted start of date range.
StartTime time.Time `json:"startTime,required" format:"date-time"`
JSON dnsSummaryQueryTypeResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryQueryTypeResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [DNSSummaryQueryTypeResponseMetaDateRange]
type dnsSummaryQueryTypeResponseMetaDateRangeJSON struct {
EndTime apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryQueryTypeResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryQueryTypeResponseMetaDateRangeJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryQueryTypeResponseMetaConfidenceInfo struct {
Annotations []DNSSummaryQueryTypeResponseMetaConfidenceInfoAnnotation `json:"annotations"`
Level int64 `json:"level"`
JSON dnsSummaryQueryTypeResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryQueryTypeResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [DNSSummaryQueryTypeResponseMetaConfidenceInfo]
type dnsSummaryQueryTypeResponseMetaConfidenceInfoJSON struct {
Annotations apijson.Field
Level apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryQueryTypeResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryQueryTypeResponseMetaConfidenceInfoJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryQueryTypeResponseMetaConfidenceInfoAnnotation struct {
DataSource string `json:"dataSource,required"`
Description string `json:"description,required"`
EventType string `json:"eventType,required"`
IsInstantaneous bool `json:"isInstantaneous,required"`
EndTime time.Time `json:"endTime" format:"date-time"`
LinkedURL string `json:"linkedUrl"`
StartTime time.Time `json:"startTime" format:"date-time"`
JSON dnsSummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [DNSSummaryQueryTypeResponseMetaConfidenceInfoAnnotation]
type dnsSummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON struct {
DataSource apijson.Field
Description apijson.Field
EventType apijson.Field
IsInstantaneous apijson.Field
EndTime apijson.Field
LinkedURL apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryQueryTypeResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryQueryTypeResponseMetaConfidenceInfoAnnotationJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryQueryTypeResponseSummary0 struct {
A string `json:"A,required"`
AAAA string `json:"AAAA,required"`
HTTPS string `json:"HTTPS,required"`
NS string `json:"NS,required"`
PTR string `json:"PTR,required"`
JSON dnsSummaryQueryTypeResponseSummary0JSON `json:"-"`
}

// dnsSummaryQueryTypeResponseSummary0JSON contains the JSON metadata for the
// struct [DNSSummaryQueryTypeResponseSummary0]
type dnsSummaryQueryTypeResponseSummary0JSON struct {
A apijson.Field
AAAA apijson.Field
HTTPS apijson.Field
NS apijson.Field
PTR apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryQueryTypeResponseSummary0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryQueryTypeResponseSummary0JSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryResponseCodesResponse struct {
Meta DNSSummaryResponseCodesResponseMeta `json:"meta,required"`
Summary0 DNSSummaryResponseCodesResponseSummary0 `json:"summary_0,required"`
JSON dnsSummaryResponseCodesResponseJSON `json:"-"`
}

// dnsSummaryResponseCodesResponseJSON contains the JSON metadata for the struct
// [DNSSummaryResponseCodesResponse]
type dnsSummaryResponseCodesResponseJSON struct {
Meta apijson.Field
Summary0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseCodesResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseCodesResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryResponseCodesResponseMeta struct {
DateRange []DNSSummaryResponseCodesResponseMetaDateRange `json:"dateRange,required"`
LastUpdated string `json:"lastUpdated,required"`
Normalization string `json:"normalization,required"`
ConfidenceInfo DNSSummaryResponseCodesResponseMetaConfidenceInfo `json:"confidenceInfo"`
JSON dnsSummaryResponseCodesResponseMetaJSON `json:"-"`
}

// dnsSummaryResponseCodesResponseMetaJSON contains the JSON metadata for the
// struct [DNSSummaryResponseCodesResponseMeta]
type dnsSummaryResponseCodesResponseMetaJSON struct {
DateRange apijson.Field
LastUpdated apijson.Field
Normalization apijson.Field
ConfidenceInfo apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseCodesResponseMeta) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseCodesResponseMetaJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryResponseCodesResponseMetaDateRange struct {
// Adjusted end of date range.
EndTime time.Time `json:"endTime,required" format:"date-time"`
// Adjusted start of date range.
StartTime time.Time `json:"startTime,required" format:"date-time"`
JSON dnsSummaryResponseCodesResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryResponseCodesResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [DNSSummaryResponseCodesResponseMetaDateRange]
type dnsSummaryResponseCodesResponseMetaDateRangeJSON struct {
EndTime apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseCodesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseCodesResponseMetaDateRangeJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryResponseCodesResponseMetaConfidenceInfo struct {
Annotations []DNSSummaryResponseCodesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
Level int64 `json:"level"`
JSON dnsSummaryResponseCodesResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryResponseCodesResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [DNSSummaryResponseCodesResponseMetaConfidenceInfo]
type dnsSummaryResponseCodesResponseMetaConfidenceInfoJSON struct {
Annotations apijson.Field
Level apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseCodesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseCodesResponseMetaConfidenceInfoJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryResponseCodesResponseMetaConfidenceInfoAnnotation struct {
DataSource string `json:"dataSource,required"`
Description string `json:"description,required"`
EventType string `json:"eventType,required"`
IsInstantaneous bool `json:"isInstantaneous,required"`
EndTime time.Time `json:"endTime" format:"date-time"`
LinkedURL string `json:"linkedUrl"`
StartTime time.Time `json:"startTime" format:"date-time"`
JSON dnsSummaryResponseCodesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryResponseCodesResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [DNSSummaryResponseCodesResponseMetaConfidenceInfoAnnotation]
type dnsSummaryResponseCodesResponseMetaConfidenceInfoAnnotationJSON struct {
DataSource apijson.Field
Description apijson.Field
EventType apijson.Field
IsInstantaneous apijson.Field
EndTime apijson.Field
LinkedURL apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseCodesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseCodesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryResponseCodesResponseSummary0 struct {
Formerr string `json:"FORMERR,required"`
Noerror string `json:"NOERROR,required"`
Notimp string `json:"NOTIMP,required"`
Nxdomain string `json:"NXDOMAIN,required"`
Refused string `json:"REFUSED,required"`
Servfail string `json:"SERVFAIL,required"`
JSON dnsSummaryResponseCodesResponseSummary0JSON `json:"-"`
}

// dnsSummaryResponseCodesResponseSummary0JSON contains the JSON metadata for the
// struct [DNSSummaryResponseCodesResponseSummary0]
type dnsSummaryResponseCodesResponseSummary0JSON struct {
Formerr apijson.Field
Noerror apijson.Field
Notimp apijson.Field
Nxdomain apijson.Field
Refused apijson.Field
Servfail apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseCodesResponseSummary0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseCodesResponseSummary0JSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryResponseTTLResponse struct {
Meta DNSSummaryResponseTTLResponseMeta `json:"meta,required"`
Summary0 DNSSummaryResponseTTLResponseSummary0 `json:"summary_0,required"`
JSON dnsSummaryResponseTTLResponseJSON `json:"-"`
}

// dnsSummaryResponseTTLResponseJSON contains the JSON metadata for the struct
// [DNSSummaryResponseTTLResponse]
type dnsSummaryResponseTTLResponseJSON struct {
Meta apijson.Field
Summary0 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseTTLResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseTTLResponseJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryResponseTTLResponseMeta struct {
DateRange []DNSSummaryResponseTTLResponseMetaDateRange `json:"dateRange,required"`
LastUpdated string `json:"lastUpdated,required"`
Normalization string `json:"normalization,required"`
ConfidenceInfo DNSSummaryResponseTTLResponseMetaConfidenceInfo `json:"confidenceInfo"`
JSON dnsSummaryResponseTTLResponseMetaJSON `json:"-"`
}

// dnsSummaryResponseTTLResponseMetaJSON contains the JSON metadata for the struct
// [DNSSummaryResponseTTLResponseMeta]
type dnsSummaryResponseTTLResponseMetaJSON struct {
DateRange apijson.Field
LastUpdated apijson.Field
Normalization apijson.Field
ConfidenceInfo apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseTTLResponseMeta) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseTTLResponseMetaJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryResponseTTLResponseMetaDateRange struct {
// Adjusted end of date range.
EndTime time.Time `json:"endTime,required" format:"date-time"`
// Adjusted start of date range.
StartTime time.Time `json:"startTime,required" format:"date-time"`
JSON dnsSummaryResponseTTLResponseMetaDateRangeJSON `json:"-"`
}

// dnsSummaryResponseTTLResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [DNSSummaryResponseTTLResponseMetaDateRange]
type dnsSummaryResponseTTLResponseMetaDateRangeJSON struct {
EndTime apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseTTLResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseTTLResponseMetaDateRangeJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryResponseTTLResponseMetaConfidenceInfo struct {
Annotations []DNSSummaryResponseTTLResponseMetaConfidenceInfoAnnotation `json:"annotations"`
Level int64 `json:"level"`
JSON dnsSummaryResponseTTLResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsSummaryResponseTTLResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [DNSSummaryResponseTTLResponseMetaConfidenceInfo]
type dnsSummaryResponseTTLResponseMetaConfidenceInfoJSON struct {
Annotations apijson.Field
Level apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseTTLResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseTTLResponseMetaConfidenceInfoJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryResponseTTLResponseMetaConfidenceInfoAnnotation struct {
DataSource string `json:"dataSource,required"`
Description string `json:"description,required"`
EventType string `json:"eventType,required"`
IsInstantaneous bool `json:"isInstantaneous,required"`
EndTime time.Time `json:"endTime" format:"date-time"`
LinkedURL string `json:"linkedUrl"`
StartTime time.Time `json:"startTime" format:"date-time"`
JSON dnsSummaryResponseTTLResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsSummaryResponseTTLResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [DNSSummaryResponseTTLResponseMetaConfidenceInfoAnnotation]
type dnsSummaryResponseTTLResponseMetaConfidenceInfoAnnotationJSON struct {
DataSource apijson.Field
Description apijson.Field
EventType apijson.Field
IsInstantaneous apijson.Field
EndTime apijson.Field
LinkedURL apijson.Field
StartTime apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseTTLResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseTTLResponseMetaConfidenceInfoAnnotationJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryResponseTTLResponseSummary0 struct {
0m string `json:"<=0m,required"`
15m string `json:"<=15m,required"`
1d string `json:"<=1d,required"`
1h string `json:"<=1h,required"`
1m string `json:"<=1m,required"`
1w string `json:"<=1w,required"`
1y string `json:"<=1y,required"`
5m string `json:"<=5m,required"`
1y string `json:">1y,required"`
JSON dnsSummaryResponseTTLResponseSummary0JSON `json:"-"`
}

// dnsSummaryResponseTTLResponseSummary0JSON contains the JSON metadata for the
// struct [DNSSummaryResponseTTLResponseSummary0]
type dnsSummaryResponseTTLResponseSummary0JSON struct {
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

func (r *DNSSummaryResponseTTLResponseSummary0) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseTTLResponseSummary0JSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryCacheHitParams struct {
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
Format param.Field[DNSSummaryCacheHitParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryCacheHitParams]'s query parameters as
// `url.Values`.
func (r DNSSummaryCacheHitParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Format results are returned in.
type DNSSummaryCacheHitParamsFormat string

const (
  DNSSummaryCacheHitParamsFormatJson DNSSummaryCacheHitParamsFormat = "JSON"
  DNSSummaryCacheHitParamsFormatCsv DNSSummaryCacheHitParamsFormat = "CSV"
)

func (r DNSSummaryCacheHitParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSSummaryCacheHitParamsFormatJson, DNSSummaryCacheHitParamsFormatCsv:
      return true
  }
  return false
}

type DNSSummaryCacheHitResponseEnvelope struct {
Result DNSSummaryCacheHitResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsSummaryCacheHitResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryCacheHitResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSSummaryCacheHitResponseEnvelope]
type dnsSummaryCacheHitResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryCacheHitResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryCacheHitResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDNSSECParams struct {
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
Format param.Field[DNSSummaryDNSSECParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryDNSSECParams]'s query parameters as `url.Values`.
func (r DNSSummaryDNSSECParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Format results are returned in.
type DNSSummaryDNSSECParamsFormat string

const (
  DNSSummaryDNSSECParamsFormatJson DNSSummaryDNSSECParamsFormat = "JSON"
  DNSSummaryDNSSECParamsFormatCsv DNSSummaryDNSSECParamsFormat = "CSV"
)

func (r DNSSummaryDNSSECParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSSummaryDNSSECParamsFormatJson, DNSSummaryDNSSECParamsFormatCsv:
      return true
  }
  return false
}

type DNSSummaryDNSSECResponseEnvelope struct {
Result DNSSummaryDNSSECResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsSummaryDNSSECResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryDNSSECResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSSummaryDNSSECResponseEnvelope]
type dnsSummaryDNSSECResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDNSSECAwareParams struct {
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
Format param.Field[DNSSummaryDNSSECAwareParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryDNSSECAwareParams]'s query parameters as
// `url.Values`.
func (r DNSSummaryDNSSECAwareParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Format results are returned in.
type DNSSummaryDNSSECAwareParamsFormat string

const (
  DNSSummaryDNSSECAwareParamsFormatJson DNSSummaryDNSSECAwareParamsFormat = "JSON"
  DNSSummaryDNSSECAwareParamsFormatCsv DNSSummaryDNSSECAwareParamsFormat = "CSV"
)

func (r DNSSummaryDNSSECAwareParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSSummaryDNSSECAwareParamsFormatJson, DNSSummaryDNSSECAwareParamsFormatCsv:
      return true
  }
  return false
}

type DNSSummaryDNSSECAwareResponseEnvelope struct {
Result DNSSummaryDNSSECAwareResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsSummaryDNSSECAwareResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryDNSSECAwareResponseEnvelopeJSON contains the JSON metadata for the
// struct [DNSSummaryDNSSECAwareResponseEnvelope]
type dnsSummaryDNSSECAwareResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDNSSECAwareResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDNSSECAwareResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryDNSSECE2EParams struct {
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
Format param.Field[DNSSummaryDnssece2EParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryDNSSECE2EParams]'s query parameters as
// `url.Values`.
func (r DNSSummaryDNSSECE2EParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Format results are returned in.
type DNSSummaryDnssece2EParamsFormat string

const (
  DNSSummaryDnssece2EParamsFormatJson DNSSummaryDnssece2EParamsFormat = "JSON"
  DNSSummaryDnssece2EParamsFormatCsv DNSSummaryDnssece2EParamsFormat = "CSV"
)

func (r DNSSummaryDnssece2EParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSSummaryDnssece2EParamsFormatJson, DNSSummaryDnssece2EParamsFormatCsv:
      return true
  }
  return false
}

type DNSSummaryDnssece2EResponseEnvelope struct {
Result DNSSummaryDnssece2EResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsSummaryDnssece2EResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryDnssece2EResponseEnvelopeJSON contains the JSON metadata for the
// struct [DNSSummaryDnssece2EResponseEnvelope]
type dnsSummaryDnssece2EResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryDnssece2EResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryDnssece2EResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryIPVersionParams struct {
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
Format param.Field[DNSSummaryIPVersionParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryIPVersionParams]'s query parameters as
// `url.Values`.
func (r DNSSummaryIPVersionParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Format results are returned in.
type DNSSummaryIPVersionParamsFormat string

const (
  DNSSummaryIPVersionParamsFormatJson DNSSummaryIPVersionParamsFormat = "JSON"
  DNSSummaryIPVersionParamsFormatCsv DNSSummaryIPVersionParamsFormat = "CSV"
)

func (r DNSSummaryIPVersionParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSSummaryIPVersionParamsFormatJson, DNSSummaryIPVersionParamsFormatCsv:
      return true
  }
  return false
}

type DNSSummaryIPVersionResponseEnvelope struct {
Result DNSSummaryIPVersionResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsSummaryIPVersionResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryIPVersionResponseEnvelopeJSON contains the JSON metadata for the
// struct [DNSSummaryIPVersionResponseEnvelope]
type dnsSummaryIPVersionResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryIPVersionResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryMatchingAnswerParams struct {
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
Format param.Field[DNSSummaryMatchingAnswerParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryMatchingAnswerParams]'s query parameters as
// `url.Values`.
func (r DNSSummaryMatchingAnswerParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Format results are returned in.
type DNSSummaryMatchingAnswerParamsFormat string

const (
  DNSSummaryMatchingAnswerParamsFormatJson DNSSummaryMatchingAnswerParamsFormat = "JSON"
  DNSSummaryMatchingAnswerParamsFormatCsv DNSSummaryMatchingAnswerParamsFormat = "CSV"
)

func (r DNSSummaryMatchingAnswerParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSSummaryMatchingAnswerParamsFormatJson, DNSSummaryMatchingAnswerParamsFormatCsv:
      return true
  }
  return false
}

type DNSSummaryMatchingAnswerResponseEnvelope struct {
Result DNSSummaryMatchingAnswerResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsSummaryMatchingAnswerResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryMatchingAnswerResponseEnvelopeJSON contains the JSON metadata for the
// struct [DNSSummaryMatchingAnswerResponseEnvelope]
type dnsSummaryMatchingAnswerResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryMatchingAnswerResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryMatchingAnswerResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryProtocolParams struct {
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
Format param.Field[DNSSummaryProtocolParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryProtocolParams]'s query parameters as
// `url.Values`.
func (r DNSSummaryProtocolParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Format results are returned in.
type DNSSummaryProtocolParamsFormat string

const (
  DNSSummaryProtocolParamsFormatJson DNSSummaryProtocolParamsFormat = "JSON"
  DNSSummaryProtocolParamsFormatCsv DNSSummaryProtocolParamsFormat = "CSV"
)

func (r DNSSummaryProtocolParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSSummaryProtocolParamsFormatJson, DNSSummaryProtocolParamsFormatCsv:
      return true
  }
  return false
}

type DNSSummaryProtocolResponseEnvelope struct {
Result DNSSummaryProtocolResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsSummaryProtocolResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryProtocolResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSSummaryProtocolResponseEnvelope]
type dnsSummaryProtocolResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryProtocolResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryQueryTypeParams struct {
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
Format param.Field[DNSSummaryQueryTypeParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryQueryTypeParams]'s query parameters as
// `url.Values`.
func (r DNSSummaryQueryTypeParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Format results are returned in.
type DNSSummaryQueryTypeParamsFormat string

const (
  DNSSummaryQueryTypeParamsFormatJson DNSSummaryQueryTypeParamsFormat = "JSON"
  DNSSummaryQueryTypeParamsFormatCsv DNSSummaryQueryTypeParamsFormat = "CSV"
)

func (r DNSSummaryQueryTypeParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSSummaryQueryTypeParamsFormatJson, DNSSummaryQueryTypeParamsFormatCsv:
      return true
  }
  return false
}

type DNSSummaryQueryTypeResponseEnvelope struct {
Result DNSSummaryQueryTypeResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsSummaryQueryTypeResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryQueryTypeResponseEnvelopeJSON contains the JSON metadata for the
// struct [DNSSummaryQueryTypeResponseEnvelope]
type dnsSummaryQueryTypeResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryQueryTypeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryQueryTypeResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryResponseCodesParams struct {
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
Format param.Field[DNSSummaryResponseCodesParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryResponseCodesParams]'s query parameters as
// `url.Values`.
func (r DNSSummaryResponseCodesParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Format results are returned in.
type DNSSummaryResponseCodesParamsFormat string

const (
  DNSSummaryResponseCodesParamsFormatJson DNSSummaryResponseCodesParamsFormat = "JSON"
  DNSSummaryResponseCodesParamsFormatCsv DNSSummaryResponseCodesParamsFormat = "CSV"
)

func (r DNSSummaryResponseCodesParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSSummaryResponseCodesParamsFormatJson, DNSSummaryResponseCodesParamsFormatCsv:
      return true
  }
  return false
}

type DNSSummaryResponseCodesResponseEnvelope struct {
Result DNSSummaryResponseCodesResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsSummaryResponseCodesResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryResponseCodesResponseEnvelopeJSON contains the JSON metadata for the
// struct [DNSSummaryResponseCodesResponseEnvelope]
type dnsSummaryResponseCodesResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseCodesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseCodesResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

type DNSSummaryResponseTTLParams struct {
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
Format param.Field[DNSSummaryResponseTTLParamsFormat] `query:"format"`
// Array of comma separated list of locations (alpha-2 country codes). Start with
// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
// but includes results from PT.
Location param.Field[[]string] `query:"location"`
// Array of names that will be used to name the series in responses.
Name param.Field[[]string] `query:"name"`
// Filter for ccTLD.
Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSSummaryResponseTTLParams]'s query parameters as
// `url.Values`.
func (r DNSSummaryResponseTTLParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Format results are returned in.
type DNSSummaryResponseTTLParamsFormat string

const (
  DNSSummaryResponseTTLParamsFormatJson DNSSummaryResponseTTLParamsFormat = "JSON"
  DNSSummaryResponseTTLParamsFormatCsv DNSSummaryResponseTTLParamsFormat = "CSV"
)

func (r DNSSummaryResponseTTLParamsFormat) IsKnown() (bool) {
  switch r {
  case DNSSummaryResponseTTLParamsFormatJson, DNSSummaryResponseTTLParamsFormatCsv:
      return true
  }
  return false
}

type DNSSummaryResponseTTLResponseEnvelope struct {
Result DNSSummaryResponseTTLResponse `json:"result,required"`
Success bool `json:"success,required"`
JSON dnsSummaryResponseTTLResponseEnvelopeJSON `json:"-"`
}

// dnsSummaryResponseTTLResponseEnvelopeJSON contains the JSON metadata for the
// struct [DNSSummaryResponseTTLResponseEnvelope]
type dnsSummaryResponseTTLResponseEnvelopeJSON struct {
Result apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DNSSummaryResponseTTLResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r dnsSummaryResponseTTLResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}
