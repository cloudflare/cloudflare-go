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

// AS112TimeseriesGroupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAS112TimeseriesGroupService] method instead.
type AS112TimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewAS112TimeseriesGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAS112TimeseriesGroupService(opts ...option.RequestOption) (r *AS112TimeseriesGroupService) {
	r = &AS112TimeseriesGroupService{}
	r.Options = opts
	return
}

// Percentage distribution of DNS AS112 queries by DNSSEC support over time.
func (r *AS112TimeseriesGroupService) DNSSEC(ctx context.Context, query AS112TimeseriesGroupDNSSECParams, opts ...option.RequestOption) (res *AS112TimeseriesGroupDNSSECResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AS112TimeseriesGroupDNSSECResponseEnvelope
	path := "radar/as112/timeseries_groups/dnssec"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of AS112 DNS queries by EDNS support over time.
func (r *AS112TimeseriesGroupService) Edns(ctx context.Context, query AS112TimeseriesGroupEdnsParams, opts ...option.RequestOption) (res *AS112TimeseriesGroupEdnsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AS112TimeseriesGroupEdnsResponseEnvelope
	path := "radar/as112/timeseries_groups/edns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of AS112 DNS queries by IP Version over time.
func (r *AS112TimeseriesGroupService) IPVersion(ctx context.Context, query AS112TimeseriesGroupIPVersionParams, opts ...option.RequestOption) (res *AS112TimeseriesGroupIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AS112TimeseriesGroupIPVersionResponseEnvelope
	path := "radar/as112/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of AS112 dns requests classified per Protocol over time.
func (r *AS112TimeseriesGroupService) Protocol(ctx context.Context, query AS112TimeseriesGroupProtocolParams, opts ...option.RequestOption) (res *AS112TimeseriesGroupProtocolResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AS112TimeseriesGroupProtocolResponseEnvelope
	path := "radar/as112/timeseries_groups/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of AS112 DNS queries by Query Type over time.
func (r *AS112TimeseriesGroupService) QueryType(ctx context.Context, query AS112TimeseriesGroupQueryTypeParams, opts ...option.RequestOption) (res *AS112TimeseriesGroupQueryTypeResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AS112TimeseriesGroupQueryTypeResponseEnvelope
	path := "radar/as112/timeseries_groups/query_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of AS112 dns requests classified per Response Codes over
// time.
func (r *AS112TimeseriesGroupService) ResponseCodes(ctx context.Context, query AS112TimeseriesGroupResponseCodesParams, opts ...option.RequestOption) (res *AS112TimeseriesGroupResponseCodesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AS112TimeseriesGroupResponseCodesResponseEnvelope
	path := "radar/as112/timeseries_groups/response_codes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AS112TimeseriesGroupDNSSECResponse struct {
	Meta   interface{}                              `json:"meta,required"`
	Serie0 AS112TimeseriesGroupDNSSECResponseSerie0 `json:"serie_0,required"`
	JSON   as112TimeseriesGroupDNSSECResponseJSON   `json:"-"`
}

// as112TimeseriesGroupDNSSECResponseJSON contains the JSON metadata for the struct
// [AS112TimeseriesGroupDNSSECResponse]
type as112TimeseriesGroupDNSSECResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupDNSSECResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupDNSSECResponseJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupDNSSECResponseSerie0 struct {
	NotSupported []string                                     `json:"NOT_SUPPORTED,required"`
	Supported    []string                                     `json:"SUPPORTED,required"`
	JSON         as112TimeseriesGroupDNSSECResponseSerie0JSON `json:"-"`
}

// as112TimeseriesGroupDNSSECResponseSerie0JSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupDNSSECResponseSerie0]
type as112TimeseriesGroupDNSSECResponseSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AS112TimeseriesGroupDNSSECResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupDNSSECResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupEdnsResponse struct {
	Meta   interface{}                            `json:"meta,required"`
	Serie0 AS112TimeseriesGroupEdnsResponseSerie0 `json:"serie_0,required"`
	JSON   as112TimeseriesGroupEdnsResponseJSON   `json:"-"`
}

// as112TimeseriesGroupEdnsResponseJSON contains the JSON metadata for the struct
// [AS112TimeseriesGroupEdnsResponse]
type as112TimeseriesGroupEdnsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupEdnsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupEdnsResponseJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupEdnsResponseSerie0 struct {
	NotSupported []string                                   `json:"NOT_SUPPORTED,required"`
	Supported    []string                                   `json:"SUPPORTED,required"`
	JSON         as112TimeseriesGroupEdnsResponseSerie0JSON `json:"-"`
}

// as112TimeseriesGroupEdnsResponseSerie0JSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupEdnsResponseSerie0]
type as112TimeseriesGroupEdnsResponseSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AS112TimeseriesGroupEdnsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupEdnsResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupIPVersionResponse struct {
	Meta   interface{}                                 `json:"meta,required"`
	Serie0 AS112TimeseriesGroupIPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   as112TimeseriesGroupIPVersionResponseJSON   `json:"-"`
}

// as112TimeseriesGroupIPVersionResponseJSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupIPVersionResponse]
type as112TimeseriesGroupIPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupIPVersionResponseSerie0 struct {
	IPv4 []string                                        `json:"IPv4,required"`
	IPv6 []string                                        `json:"IPv6,required"`
	JSON as112TimeseriesGroupIPVersionResponseSerie0JSON `json:"-"`
}

// as112TimeseriesGroupIPVersionResponseSerie0JSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupIPVersionResponseSerie0]
type as112TimeseriesGroupIPVersionResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupIPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupIPVersionResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupProtocolResponse struct {
	Meta   interface{}                                `json:"meta,required"`
	Serie0 AS112TimeseriesGroupProtocolResponseSerie0 `json:"serie_0,required"`
	JSON   as112TimeseriesGroupProtocolResponseJSON   `json:"-"`
}

// as112TimeseriesGroupProtocolResponseJSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupProtocolResponse]
type as112TimeseriesGroupProtocolResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupProtocolResponseJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupProtocolResponseSerie0 struct {
	TCP  []string                                       `json:"tcp,required"`
	Udp  []string                                       `json:"udp,required"`
	JSON as112TimeseriesGroupProtocolResponseSerie0JSON `json:"-"`
}

// as112TimeseriesGroupProtocolResponseSerie0JSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupProtocolResponseSerie0]
type as112TimeseriesGroupProtocolResponseSerie0JSON struct {
	TCP         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupProtocolResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupProtocolResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupQueryTypeResponse struct {
	Meta   interface{}                                 `json:"meta,required"`
	Serie0 AS112TimeseriesGroupQueryTypeResponseSerie0 `json:"serie_0,required"`
	JSON   as112TimeseriesGroupQueryTypeResponseJSON   `json:"-"`
}

// as112TimeseriesGroupQueryTypeResponseJSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupQueryTypeResponse]
type as112TimeseriesGroupQueryTypeResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupQueryTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupQueryTypeResponseJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupQueryTypeResponseSerie0 struct {
	A    []string                                        `json:"A,required"`
	AAAA []string                                        `json:"AAAA,required"`
	PTR  []string                                        `json:"PTR,required"`
	SOA  []string                                        `json:"SOA,required"`
	SRV  []string                                        `json:"SRV,required"`
	JSON as112TimeseriesGroupQueryTypeResponseSerie0JSON `json:"-"`
}

// as112TimeseriesGroupQueryTypeResponseSerie0JSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupQueryTypeResponseSerie0]
type as112TimeseriesGroupQueryTypeResponseSerie0JSON struct {
	A           apijson.Field
	AAAA        apijson.Field
	PTR         apijson.Field
	SOA         apijson.Field
	SRV         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupQueryTypeResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupQueryTypeResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupResponseCodesResponse struct {
	Meta   interface{}                                     `json:"meta,required"`
	Serie0 AS112TimeseriesGroupResponseCodesResponseSerie0 `json:"serie_0,required"`
	JSON   as112TimeseriesGroupResponseCodesResponseJSON   `json:"-"`
}

// as112TimeseriesGroupResponseCodesResponseJSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupResponseCodesResponse]
type as112TimeseriesGroupResponseCodesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupResponseCodesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupResponseCodesResponseJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupResponseCodesResponseSerie0 struct {
	Noerror  []string                                            `json:"NOERROR,required"`
	Nxdomain []string                                            `json:"NXDOMAIN,required"`
	JSON     as112TimeseriesGroupResponseCodesResponseSerie0JSON `json:"-"`
}

// as112TimeseriesGroupResponseCodesResponseSerie0JSON contains the JSON metadata
// for the struct [AS112TimeseriesGroupResponseCodesResponseSerie0]
type as112TimeseriesGroupResponseCodesResponseSerie0JSON struct {
	Noerror     apijson.Field
	Nxdomain    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupResponseCodesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupResponseCodesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupDNSSECParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AS112TimeseriesGroupDNSSECParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]AS112TimeseriesGroupDNSSECParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AS112TimeseriesGroupDNSSECParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112TimeseriesGroupDNSSECParams]'s query parameters as
// `url.Values`.
func (r AS112TimeseriesGroupDNSSECParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupDNSSECParamsAggInterval string

const (
	AS112TimeseriesGroupDNSSECParamsAggInterval15m AS112TimeseriesGroupDNSSECParamsAggInterval = "15m"
	AS112TimeseriesGroupDNSSECParamsAggInterval1h  AS112TimeseriesGroupDNSSECParamsAggInterval = "1h"
	AS112TimeseriesGroupDNSSECParamsAggInterval1d  AS112TimeseriesGroupDNSSECParamsAggInterval = "1d"
	AS112TimeseriesGroupDNSSECParamsAggInterval1w  AS112TimeseriesGroupDNSSECParamsAggInterval = "1w"
)

func (r AS112TimeseriesGroupDNSSECParamsAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupDNSSECParamsAggInterval15m, AS112TimeseriesGroupDNSSECParamsAggInterval1h, AS112TimeseriesGroupDNSSECParamsAggInterval1d, AS112TimeseriesGroupDNSSECParamsAggInterval1w:
		return true
	}
	return false
}

type AS112TimeseriesGroupDNSSECParamsDateRange string

const (
	AS112TimeseriesGroupDNSSECParamsDateRange1d         AS112TimeseriesGroupDNSSECParamsDateRange = "1d"
	AS112TimeseriesGroupDNSSECParamsDateRange2d         AS112TimeseriesGroupDNSSECParamsDateRange = "2d"
	AS112TimeseriesGroupDNSSECParamsDateRange7d         AS112TimeseriesGroupDNSSECParamsDateRange = "7d"
	AS112TimeseriesGroupDNSSECParamsDateRange14d        AS112TimeseriesGroupDNSSECParamsDateRange = "14d"
	AS112TimeseriesGroupDNSSECParamsDateRange28d        AS112TimeseriesGroupDNSSECParamsDateRange = "28d"
	AS112TimeseriesGroupDNSSECParamsDateRange12w        AS112TimeseriesGroupDNSSECParamsDateRange = "12w"
	AS112TimeseriesGroupDNSSECParamsDateRange24w        AS112TimeseriesGroupDNSSECParamsDateRange = "24w"
	AS112TimeseriesGroupDNSSECParamsDateRange52w        AS112TimeseriesGroupDNSSECParamsDateRange = "52w"
	AS112TimeseriesGroupDNSSECParamsDateRange1dControl  AS112TimeseriesGroupDNSSECParamsDateRange = "1dControl"
	AS112TimeseriesGroupDNSSECParamsDateRange2dControl  AS112TimeseriesGroupDNSSECParamsDateRange = "2dControl"
	AS112TimeseriesGroupDNSSECParamsDateRange7dControl  AS112TimeseriesGroupDNSSECParamsDateRange = "7dControl"
	AS112TimeseriesGroupDNSSECParamsDateRange14dControl AS112TimeseriesGroupDNSSECParamsDateRange = "14dControl"
	AS112TimeseriesGroupDNSSECParamsDateRange28dControl AS112TimeseriesGroupDNSSECParamsDateRange = "28dControl"
	AS112TimeseriesGroupDNSSECParamsDateRange12wControl AS112TimeseriesGroupDNSSECParamsDateRange = "12wControl"
	AS112TimeseriesGroupDNSSECParamsDateRange24wControl AS112TimeseriesGroupDNSSECParamsDateRange = "24wControl"
)

func (r AS112TimeseriesGroupDNSSECParamsDateRange) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupDNSSECParamsDateRange1d, AS112TimeseriesGroupDNSSECParamsDateRange2d, AS112TimeseriesGroupDNSSECParamsDateRange7d, AS112TimeseriesGroupDNSSECParamsDateRange14d, AS112TimeseriesGroupDNSSECParamsDateRange28d, AS112TimeseriesGroupDNSSECParamsDateRange12w, AS112TimeseriesGroupDNSSECParamsDateRange24w, AS112TimeseriesGroupDNSSECParamsDateRange52w, AS112TimeseriesGroupDNSSECParamsDateRange1dControl, AS112TimeseriesGroupDNSSECParamsDateRange2dControl, AS112TimeseriesGroupDNSSECParamsDateRange7dControl, AS112TimeseriesGroupDNSSECParamsDateRange14dControl, AS112TimeseriesGroupDNSSECParamsDateRange28dControl, AS112TimeseriesGroupDNSSECParamsDateRange12wControl, AS112TimeseriesGroupDNSSECParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AS112TimeseriesGroupDNSSECParamsFormat string

const (
	AS112TimeseriesGroupDNSSECParamsFormatJson AS112TimeseriesGroupDNSSECParamsFormat = "JSON"
	AS112TimeseriesGroupDNSSECParamsFormatCsv  AS112TimeseriesGroupDNSSECParamsFormat = "CSV"
)

func (r AS112TimeseriesGroupDNSSECParamsFormat) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupDNSSECParamsFormatJson, AS112TimeseriesGroupDNSSECParamsFormatCsv:
		return true
	}
	return false
}

type AS112TimeseriesGroupDNSSECResponseEnvelope struct {
	Result  AS112TimeseriesGroupDNSSECResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    as112TimeseriesGroupDNSSECResponseEnvelopeJSON `json:"-"`
}

// as112TimeseriesGroupDNSSECResponseEnvelopeJSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupDNSSECResponseEnvelope]
type as112TimeseriesGroupDNSSECResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupDNSSECResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupDNSSECResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupEdnsParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AS112TimeseriesGroupEdnsParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]AS112TimeseriesGroupEdnsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AS112TimeseriesGroupEdnsParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112TimeseriesGroupEdnsParams]'s query parameters as
// `url.Values`.
func (r AS112TimeseriesGroupEdnsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupEdnsParamsAggInterval string

const (
	AS112TimeseriesGroupEdnsParamsAggInterval15m AS112TimeseriesGroupEdnsParamsAggInterval = "15m"
	AS112TimeseriesGroupEdnsParamsAggInterval1h  AS112TimeseriesGroupEdnsParamsAggInterval = "1h"
	AS112TimeseriesGroupEdnsParamsAggInterval1d  AS112TimeseriesGroupEdnsParamsAggInterval = "1d"
	AS112TimeseriesGroupEdnsParamsAggInterval1w  AS112TimeseriesGroupEdnsParamsAggInterval = "1w"
)

func (r AS112TimeseriesGroupEdnsParamsAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupEdnsParamsAggInterval15m, AS112TimeseriesGroupEdnsParamsAggInterval1h, AS112TimeseriesGroupEdnsParamsAggInterval1d, AS112TimeseriesGroupEdnsParamsAggInterval1w:
		return true
	}
	return false
}

type AS112TimeseriesGroupEdnsParamsDateRange string

const (
	AS112TimeseriesGroupEdnsParamsDateRange1d         AS112TimeseriesGroupEdnsParamsDateRange = "1d"
	AS112TimeseriesGroupEdnsParamsDateRange2d         AS112TimeseriesGroupEdnsParamsDateRange = "2d"
	AS112TimeseriesGroupEdnsParamsDateRange7d         AS112TimeseriesGroupEdnsParamsDateRange = "7d"
	AS112TimeseriesGroupEdnsParamsDateRange14d        AS112TimeseriesGroupEdnsParamsDateRange = "14d"
	AS112TimeseriesGroupEdnsParamsDateRange28d        AS112TimeseriesGroupEdnsParamsDateRange = "28d"
	AS112TimeseriesGroupEdnsParamsDateRange12w        AS112TimeseriesGroupEdnsParamsDateRange = "12w"
	AS112TimeseriesGroupEdnsParamsDateRange24w        AS112TimeseriesGroupEdnsParamsDateRange = "24w"
	AS112TimeseriesGroupEdnsParamsDateRange52w        AS112TimeseriesGroupEdnsParamsDateRange = "52w"
	AS112TimeseriesGroupEdnsParamsDateRange1dControl  AS112TimeseriesGroupEdnsParamsDateRange = "1dControl"
	AS112TimeseriesGroupEdnsParamsDateRange2dControl  AS112TimeseriesGroupEdnsParamsDateRange = "2dControl"
	AS112TimeseriesGroupEdnsParamsDateRange7dControl  AS112TimeseriesGroupEdnsParamsDateRange = "7dControl"
	AS112TimeseriesGroupEdnsParamsDateRange14dControl AS112TimeseriesGroupEdnsParamsDateRange = "14dControl"
	AS112TimeseriesGroupEdnsParamsDateRange28dControl AS112TimeseriesGroupEdnsParamsDateRange = "28dControl"
	AS112TimeseriesGroupEdnsParamsDateRange12wControl AS112TimeseriesGroupEdnsParamsDateRange = "12wControl"
	AS112TimeseriesGroupEdnsParamsDateRange24wControl AS112TimeseriesGroupEdnsParamsDateRange = "24wControl"
)

func (r AS112TimeseriesGroupEdnsParamsDateRange) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupEdnsParamsDateRange1d, AS112TimeseriesGroupEdnsParamsDateRange2d, AS112TimeseriesGroupEdnsParamsDateRange7d, AS112TimeseriesGroupEdnsParamsDateRange14d, AS112TimeseriesGroupEdnsParamsDateRange28d, AS112TimeseriesGroupEdnsParamsDateRange12w, AS112TimeseriesGroupEdnsParamsDateRange24w, AS112TimeseriesGroupEdnsParamsDateRange52w, AS112TimeseriesGroupEdnsParamsDateRange1dControl, AS112TimeseriesGroupEdnsParamsDateRange2dControl, AS112TimeseriesGroupEdnsParamsDateRange7dControl, AS112TimeseriesGroupEdnsParamsDateRange14dControl, AS112TimeseriesGroupEdnsParamsDateRange28dControl, AS112TimeseriesGroupEdnsParamsDateRange12wControl, AS112TimeseriesGroupEdnsParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AS112TimeseriesGroupEdnsParamsFormat string

const (
	AS112TimeseriesGroupEdnsParamsFormatJson AS112TimeseriesGroupEdnsParamsFormat = "JSON"
	AS112TimeseriesGroupEdnsParamsFormatCsv  AS112TimeseriesGroupEdnsParamsFormat = "CSV"
)

func (r AS112TimeseriesGroupEdnsParamsFormat) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupEdnsParamsFormatJson, AS112TimeseriesGroupEdnsParamsFormatCsv:
		return true
	}
	return false
}

type AS112TimeseriesGroupEdnsResponseEnvelope struct {
	Result  AS112TimeseriesGroupEdnsResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    as112TimeseriesGroupEdnsResponseEnvelopeJSON `json:"-"`
}

// as112TimeseriesGroupEdnsResponseEnvelopeJSON contains the JSON metadata for the
// struct [AS112TimeseriesGroupEdnsResponseEnvelope]
type as112TimeseriesGroupEdnsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupEdnsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupEdnsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupIPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AS112TimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]AS112TimeseriesGroupIPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AS112TimeseriesGroupIPVersionParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112TimeseriesGroupIPVersionParams]'s query parameters as
// `url.Values`.
func (r AS112TimeseriesGroupIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupIPVersionParamsAggInterval string

const (
	AS112TimeseriesGroupIPVersionParamsAggInterval15m AS112TimeseriesGroupIPVersionParamsAggInterval = "15m"
	AS112TimeseriesGroupIPVersionParamsAggInterval1h  AS112TimeseriesGroupIPVersionParamsAggInterval = "1h"
	AS112TimeseriesGroupIPVersionParamsAggInterval1d  AS112TimeseriesGroupIPVersionParamsAggInterval = "1d"
	AS112TimeseriesGroupIPVersionParamsAggInterval1w  AS112TimeseriesGroupIPVersionParamsAggInterval = "1w"
)

func (r AS112TimeseriesGroupIPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupIPVersionParamsAggInterval15m, AS112TimeseriesGroupIPVersionParamsAggInterval1h, AS112TimeseriesGroupIPVersionParamsAggInterval1d, AS112TimeseriesGroupIPVersionParamsAggInterval1w:
		return true
	}
	return false
}

type AS112TimeseriesGroupIPVersionParamsDateRange string

const (
	AS112TimeseriesGroupIPVersionParamsDateRange1d         AS112TimeseriesGroupIPVersionParamsDateRange = "1d"
	AS112TimeseriesGroupIPVersionParamsDateRange2d         AS112TimeseriesGroupIPVersionParamsDateRange = "2d"
	AS112TimeseriesGroupIPVersionParamsDateRange7d         AS112TimeseriesGroupIPVersionParamsDateRange = "7d"
	AS112TimeseriesGroupIPVersionParamsDateRange14d        AS112TimeseriesGroupIPVersionParamsDateRange = "14d"
	AS112TimeseriesGroupIPVersionParamsDateRange28d        AS112TimeseriesGroupIPVersionParamsDateRange = "28d"
	AS112TimeseriesGroupIPVersionParamsDateRange12w        AS112TimeseriesGroupIPVersionParamsDateRange = "12w"
	AS112TimeseriesGroupIPVersionParamsDateRange24w        AS112TimeseriesGroupIPVersionParamsDateRange = "24w"
	AS112TimeseriesGroupIPVersionParamsDateRange52w        AS112TimeseriesGroupIPVersionParamsDateRange = "52w"
	AS112TimeseriesGroupIPVersionParamsDateRange1dControl  AS112TimeseriesGroupIPVersionParamsDateRange = "1dControl"
	AS112TimeseriesGroupIPVersionParamsDateRange2dControl  AS112TimeseriesGroupIPVersionParamsDateRange = "2dControl"
	AS112TimeseriesGroupIPVersionParamsDateRange7dControl  AS112TimeseriesGroupIPVersionParamsDateRange = "7dControl"
	AS112TimeseriesGroupIPVersionParamsDateRange14dControl AS112TimeseriesGroupIPVersionParamsDateRange = "14dControl"
	AS112TimeseriesGroupIPVersionParamsDateRange28dControl AS112TimeseriesGroupIPVersionParamsDateRange = "28dControl"
	AS112TimeseriesGroupIPVersionParamsDateRange12wControl AS112TimeseriesGroupIPVersionParamsDateRange = "12wControl"
	AS112TimeseriesGroupIPVersionParamsDateRange24wControl AS112TimeseriesGroupIPVersionParamsDateRange = "24wControl"
)

func (r AS112TimeseriesGroupIPVersionParamsDateRange) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupIPVersionParamsDateRange1d, AS112TimeseriesGroupIPVersionParamsDateRange2d, AS112TimeseriesGroupIPVersionParamsDateRange7d, AS112TimeseriesGroupIPVersionParamsDateRange14d, AS112TimeseriesGroupIPVersionParamsDateRange28d, AS112TimeseriesGroupIPVersionParamsDateRange12w, AS112TimeseriesGroupIPVersionParamsDateRange24w, AS112TimeseriesGroupIPVersionParamsDateRange52w, AS112TimeseriesGroupIPVersionParamsDateRange1dControl, AS112TimeseriesGroupIPVersionParamsDateRange2dControl, AS112TimeseriesGroupIPVersionParamsDateRange7dControl, AS112TimeseriesGroupIPVersionParamsDateRange14dControl, AS112TimeseriesGroupIPVersionParamsDateRange28dControl, AS112TimeseriesGroupIPVersionParamsDateRange12wControl, AS112TimeseriesGroupIPVersionParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AS112TimeseriesGroupIPVersionParamsFormat string

const (
	AS112TimeseriesGroupIPVersionParamsFormatJson AS112TimeseriesGroupIPVersionParamsFormat = "JSON"
	AS112TimeseriesGroupIPVersionParamsFormatCsv  AS112TimeseriesGroupIPVersionParamsFormat = "CSV"
)

func (r AS112TimeseriesGroupIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupIPVersionParamsFormatJson, AS112TimeseriesGroupIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type AS112TimeseriesGroupIPVersionResponseEnvelope struct {
	Result  AS112TimeseriesGroupIPVersionResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    as112TimeseriesGroupIPVersionResponseEnvelopeJSON `json:"-"`
}

// as112TimeseriesGroupIPVersionResponseEnvelopeJSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupIPVersionResponseEnvelope]
type as112TimeseriesGroupIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupIPVersionResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupProtocolParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AS112TimeseriesGroupProtocolParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]AS112TimeseriesGroupProtocolParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AS112TimeseriesGroupProtocolParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112TimeseriesGroupProtocolParams]'s query parameters as
// `url.Values`.
func (r AS112TimeseriesGroupProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupProtocolParamsAggInterval string

const (
	AS112TimeseriesGroupProtocolParamsAggInterval15m AS112TimeseriesGroupProtocolParamsAggInterval = "15m"
	AS112TimeseriesGroupProtocolParamsAggInterval1h  AS112TimeseriesGroupProtocolParamsAggInterval = "1h"
	AS112TimeseriesGroupProtocolParamsAggInterval1d  AS112TimeseriesGroupProtocolParamsAggInterval = "1d"
	AS112TimeseriesGroupProtocolParamsAggInterval1w  AS112TimeseriesGroupProtocolParamsAggInterval = "1w"
)

func (r AS112TimeseriesGroupProtocolParamsAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupProtocolParamsAggInterval15m, AS112TimeseriesGroupProtocolParamsAggInterval1h, AS112TimeseriesGroupProtocolParamsAggInterval1d, AS112TimeseriesGroupProtocolParamsAggInterval1w:
		return true
	}
	return false
}

type AS112TimeseriesGroupProtocolParamsDateRange string

const (
	AS112TimeseriesGroupProtocolParamsDateRange1d         AS112TimeseriesGroupProtocolParamsDateRange = "1d"
	AS112TimeseriesGroupProtocolParamsDateRange2d         AS112TimeseriesGroupProtocolParamsDateRange = "2d"
	AS112TimeseriesGroupProtocolParamsDateRange7d         AS112TimeseriesGroupProtocolParamsDateRange = "7d"
	AS112TimeseriesGroupProtocolParamsDateRange14d        AS112TimeseriesGroupProtocolParamsDateRange = "14d"
	AS112TimeseriesGroupProtocolParamsDateRange28d        AS112TimeseriesGroupProtocolParamsDateRange = "28d"
	AS112TimeseriesGroupProtocolParamsDateRange12w        AS112TimeseriesGroupProtocolParamsDateRange = "12w"
	AS112TimeseriesGroupProtocolParamsDateRange24w        AS112TimeseriesGroupProtocolParamsDateRange = "24w"
	AS112TimeseriesGroupProtocolParamsDateRange52w        AS112TimeseriesGroupProtocolParamsDateRange = "52w"
	AS112TimeseriesGroupProtocolParamsDateRange1dControl  AS112TimeseriesGroupProtocolParamsDateRange = "1dControl"
	AS112TimeseriesGroupProtocolParamsDateRange2dControl  AS112TimeseriesGroupProtocolParamsDateRange = "2dControl"
	AS112TimeseriesGroupProtocolParamsDateRange7dControl  AS112TimeseriesGroupProtocolParamsDateRange = "7dControl"
	AS112TimeseriesGroupProtocolParamsDateRange14dControl AS112TimeseriesGroupProtocolParamsDateRange = "14dControl"
	AS112TimeseriesGroupProtocolParamsDateRange28dControl AS112TimeseriesGroupProtocolParamsDateRange = "28dControl"
	AS112TimeseriesGroupProtocolParamsDateRange12wControl AS112TimeseriesGroupProtocolParamsDateRange = "12wControl"
	AS112TimeseriesGroupProtocolParamsDateRange24wControl AS112TimeseriesGroupProtocolParamsDateRange = "24wControl"
)

func (r AS112TimeseriesGroupProtocolParamsDateRange) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupProtocolParamsDateRange1d, AS112TimeseriesGroupProtocolParamsDateRange2d, AS112TimeseriesGroupProtocolParamsDateRange7d, AS112TimeseriesGroupProtocolParamsDateRange14d, AS112TimeseriesGroupProtocolParamsDateRange28d, AS112TimeseriesGroupProtocolParamsDateRange12w, AS112TimeseriesGroupProtocolParamsDateRange24w, AS112TimeseriesGroupProtocolParamsDateRange52w, AS112TimeseriesGroupProtocolParamsDateRange1dControl, AS112TimeseriesGroupProtocolParamsDateRange2dControl, AS112TimeseriesGroupProtocolParamsDateRange7dControl, AS112TimeseriesGroupProtocolParamsDateRange14dControl, AS112TimeseriesGroupProtocolParamsDateRange28dControl, AS112TimeseriesGroupProtocolParamsDateRange12wControl, AS112TimeseriesGroupProtocolParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AS112TimeseriesGroupProtocolParamsFormat string

const (
	AS112TimeseriesGroupProtocolParamsFormatJson AS112TimeseriesGroupProtocolParamsFormat = "JSON"
	AS112TimeseriesGroupProtocolParamsFormatCsv  AS112TimeseriesGroupProtocolParamsFormat = "CSV"
)

func (r AS112TimeseriesGroupProtocolParamsFormat) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupProtocolParamsFormatJson, AS112TimeseriesGroupProtocolParamsFormatCsv:
		return true
	}
	return false
}

type AS112TimeseriesGroupProtocolResponseEnvelope struct {
	Result  AS112TimeseriesGroupProtocolResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    as112TimeseriesGroupProtocolResponseEnvelopeJSON `json:"-"`
}

// as112TimeseriesGroupProtocolResponseEnvelopeJSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupProtocolResponseEnvelope]
type as112TimeseriesGroupProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupProtocolResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupQueryTypeParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AS112TimeseriesGroupQueryTypeParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]AS112TimeseriesGroupQueryTypeParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AS112TimeseriesGroupQueryTypeParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112TimeseriesGroupQueryTypeParams]'s query parameters as
// `url.Values`.
func (r AS112TimeseriesGroupQueryTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupQueryTypeParamsAggInterval string

const (
	AS112TimeseriesGroupQueryTypeParamsAggInterval15m AS112TimeseriesGroupQueryTypeParamsAggInterval = "15m"
	AS112TimeseriesGroupQueryTypeParamsAggInterval1h  AS112TimeseriesGroupQueryTypeParamsAggInterval = "1h"
	AS112TimeseriesGroupQueryTypeParamsAggInterval1d  AS112TimeseriesGroupQueryTypeParamsAggInterval = "1d"
	AS112TimeseriesGroupQueryTypeParamsAggInterval1w  AS112TimeseriesGroupQueryTypeParamsAggInterval = "1w"
)

func (r AS112TimeseriesGroupQueryTypeParamsAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupQueryTypeParamsAggInterval15m, AS112TimeseriesGroupQueryTypeParamsAggInterval1h, AS112TimeseriesGroupQueryTypeParamsAggInterval1d, AS112TimeseriesGroupQueryTypeParamsAggInterval1w:
		return true
	}
	return false
}

type AS112TimeseriesGroupQueryTypeParamsDateRange string

const (
	AS112TimeseriesGroupQueryTypeParamsDateRange1d         AS112TimeseriesGroupQueryTypeParamsDateRange = "1d"
	AS112TimeseriesGroupQueryTypeParamsDateRange2d         AS112TimeseriesGroupQueryTypeParamsDateRange = "2d"
	AS112TimeseriesGroupQueryTypeParamsDateRange7d         AS112TimeseriesGroupQueryTypeParamsDateRange = "7d"
	AS112TimeseriesGroupQueryTypeParamsDateRange14d        AS112TimeseriesGroupQueryTypeParamsDateRange = "14d"
	AS112TimeseriesGroupQueryTypeParamsDateRange28d        AS112TimeseriesGroupQueryTypeParamsDateRange = "28d"
	AS112TimeseriesGroupQueryTypeParamsDateRange12w        AS112TimeseriesGroupQueryTypeParamsDateRange = "12w"
	AS112TimeseriesGroupQueryTypeParamsDateRange24w        AS112TimeseriesGroupQueryTypeParamsDateRange = "24w"
	AS112TimeseriesGroupQueryTypeParamsDateRange52w        AS112TimeseriesGroupQueryTypeParamsDateRange = "52w"
	AS112TimeseriesGroupQueryTypeParamsDateRange1dControl  AS112TimeseriesGroupQueryTypeParamsDateRange = "1dControl"
	AS112TimeseriesGroupQueryTypeParamsDateRange2dControl  AS112TimeseriesGroupQueryTypeParamsDateRange = "2dControl"
	AS112TimeseriesGroupQueryTypeParamsDateRange7dControl  AS112TimeseriesGroupQueryTypeParamsDateRange = "7dControl"
	AS112TimeseriesGroupQueryTypeParamsDateRange14dControl AS112TimeseriesGroupQueryTypeParamsDateRange = "14dControl"
	AS112TimeseriesGroupQueryTypeParamsDateRange28dControl AS112TimeseriesGroupQueryTypeParamsDateRange = "28dControl"
	AS112TimeseriesGroupQueryTypeParamsDateRange12wControl AS112TimeseriesGroupQueryTypeParamsDateRange = "12wControl"
	AS112TimeseriesGroupQueryTypeParamsDateRange24wControl AS112TimeseriesGroupQueryTypeParamsDateRange = "24wControl"
)

func (r AS112TimeseriesGroupQueryTypeParamsDateRange) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupQueryTypeParamsDateRange1d, AS112TimeseriesGroupQueryTypeParamsDateRange2d, AS112TimeseriesGroupQueryTypeParamsDateRange7d, AS112TimeseriesGroupQueryTypeParamsDateRange14d, AS112TimeseriesGroupQueryTypeParamsDateRange28d, AS112TimeseriesGroupQueryTypeParamsDateRange12w, AS112TimeseriesGroupQueryTypeParamsDateRange24w, AS112TimeseriesGroupQueryTypeParamsDateRange52w, AS112TimeseriesGroupQueryTypeParamsDateRange1dControl, AS112TimeseriesGroupQueryTypeParamsDateRange2dControl, AS112TimeseriesGroupQueryTypeParamsDateRange7dControl, AS112TimeseriesGroupQueryTypeParamsDateRange14dControl, AS112TimeseriesGroupQueryTypeParamsDateRange28dControl, AS112TimeseriesGroupQueryTypeParamsDateRange12wControl, AS112TimeseriesGroupQueryTypeParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AS112TimeseriesGroupQueryTypeParamsFormat string

const (
	AS112TimeseriesGroupQueryTypeParamsFormatJson AS112TimeseriesGroupQueryTypeParamsFormat = "JSON"
	AS112TimeseriesGroupQueryTypeParamsFormatCsv  AS112TimeseriesGroupQueryTypeParamsFormat = "CSV"
)

func (r AS112TimeseriesGroupQueryTypeParamsFormat) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupQueryTypeParamsFormatJson, AS112TimeseriesGroupQueryTypeParamsFormatCsv:
		return true
	}
	return false
}

type AS112TimeseriesGroupQueryTypeResponseEnvelope struct {
	Result  AS112TimeseriesGroupQueryTypeResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    as112TimeseriesGroupQueryTypeResponseEnvelopeJSON `json:"-"`
}

// as112TimeseriesGroupQueryTypeResponseEnvelopeJSON contains the JSON metadata for
// the struct [AS112TimeseriesGroupQueryTypeResponseEnvelope]
type as112TimeseriesGroupQueryTypeResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupQueryTypeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupQueryTypeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AS112TimeseriesGroupResponseCodesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AS112TimeseriesGroupResponseCodesParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]AS112TimeseriesGroupResponseCodesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AS112TimeseriesGroupResponseCodesParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AS112TimeseriesGroupResponseCodesParams]'s query parameters
// as `url.Values`.
func (r AS112TimeseriesGroupResponseCodesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AS112TimeseriesGroupResponseCodesParamsAggInterval string

const (
	AS112TimeseriesGroupResponseCodesParamsAggInterval15m AS112TimeseriesGroupResponseCodesParamsAggInterval = "15m"
	AS112TimeseriesGroupResponseCodesParamsAggInterval1h  AS112TimeseriesGroupResponseCodesParamsAggInterval = "1h"
	AS112TimeseriesGroupResponseCodesParamsAggInterval1d  AS112TimeseriesGroupResponseCodesParamsAggInterval = "1d"
	AS112TimeseriesGroupResponseCodesParamsAggInterval1w  AS112TimeseriesGroupResponseCodesParamsAggInterval = "1w"
)

func (r AS112TimeseriesGroupResponseCodesParamsAggInterval) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupResponseCodesParamsAggInterval15m, AS112TimeseriesGroupResponseCodesParamsAggInterval1h, AS112TimeseriesGroupResponseCodesParamsAggInterval1d, AS112TimeseriesGroupResponseCodesParamsAggInterval1w:
		return true
	}
	return false
}

type AS112TimeseriesGroupResponseCodesParamsDateRange string

const (
	AS112TimeseriesGroupResponseCodesParamsDateRange1d         AS112TimeseriesGroupResponseCodesParamsDateRange = "1d"
	AS112TimeseriesGroupResponseCodesParamsDateRange2d         AS112TimeseriesGroupResponseCodesParamsDateRange = "2d"
	AS112TimeseriesGroupResponseCodesParamsDateRange7d         AS112TimeseriesGroupResponseCodesParamsDateRange = "7d"
	AS112TimeseriesGroupResponseCodesParamsDateRange14d        AS112TimeseriesGroupResponseCodesParamsDateRange = "14d"
	AS112TimeseriesGroupResponseCodesParamsDateRange28d        AS112TimeseriesGroupResponseCodesParamsDateRange = "28d"
	AS112TimeseriesGroupResponseCodesParamsDateRange12w        AS112TimeseriesGroupResponseCodesParamsDateRange = "12w"
	AS112TimeseriesGroupResponseCodesParamsDateRange24w        AS112TimeseriesGroupResponseCodesParamsDateRange = "24w"
	AS112TimeseriesGroupResponseCodesParamsDateRange52w        AS112TimeseriesGroupResponseCodesParamsDateRange = "52w"
	AS112TimeseriesGroupResponseCodesParamsDateRange1dControl  AS112TimeseriesGroupResponseCodesParamsDateRange = "1dControl"
	AS112TimeseriesGroupResponseCodesParamsDateRange2dControl  AS112TimeseriesGroupResponseCodesParamsDateRange = "2dControl"
	AS112TimeseriesGroupResponseCodesParamsDateRange7dControl  AS112TimeseriesGroupResponseCodesParamsDateRange = "7dControl"
	AS112TimeseriesGroupResponseCodesParamsDateRange14dControl AS112TimeseriesGroupResponseCodesParamsDateRange = "14dControl"
	AS112TimeseriesGroupResponseCodesParamsDateRange28dControl AS112TimeseriesGroupResponseCodesParamsDateRange = "28dControl"
	AS112TimeseriesGroupResponseCodesParamsDateRange12wControl AS112TimeseriesGroupResponseCodesParamsDateRange = "12wControl"
	AS112TimeseriesGroupResponseCodesParamsDateRange24wControl AS112TimeseriesGroupResponseCodesParamsDateRange = "24wControl"
)

func (r AS112TimeseriesGroupResponseCodesParamsDateRange) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupResponseCodesParamsDateRange1d, AS112TimeseriesGroupResponseCodesParamsDateRange2d, AS112TimeseriesGroupResponseCodesParamsDateRange7d, AS112TimeseriesGroupResponseCodesParamsDateRange14d, AS112TimeseriesGroupResponseCodesParamsDateRange28d, AS112TimeseriesGroupResponseCodesParamsDateRange12w, AS112TimeseriesGroupResponseCodesParamsDateRange24w, AS112TimeseriesGroupResponseCodesParamsDateRange52w, AS112TimeseriesGroupResponseCodesParamsDateRange1dControl, AS112TimeseriesGroupResponseCodesParamsDateRange2dControl, AS112TimeseriesGroupResponseCodesParamsDateRange7dControl, AS112TimeseriesGroupResponseCodesParamsDateRange14dControl, AS112TimeseriesGroupResponseCodesParamsDateRange28dControl, AS112TimeseriesGroupResponseCodesParamsDateRange12wControl, AS112TimeseriesGroupResponseCodesParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AS112TimeseriesGroupResponseCodesParamsFormat string

const (
	AS112TimeseriesGroupResponseCodesParamsFormatJson AS112TimeseriesGroupResponseCodesParamsFormat = "JSON"
	AS112TimeseriesGroupResponseCodesParamsFormatCsv  AS112TimeseriesGroupResponseCodesParamsFormat = "CSV"
)

func (r AS112TimeseriesGroupResponseCodesParamsFormat) IsKnown() bool {
	switch r {
	case AS112TimeseriesGroupResponseCodesParamsFormatJson, AS112TimeseriesGroupResponseCodesParamsFormatCsv:
		return true
	}
	return false
}

type AS112TimeseriesGroupResponseCodesResponseEnvelope struct {
	Result  AS112TimeseriesGroupResponseCodesResponse             `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    as112TimeseriesGroupResponseCodesResponseEnvelopeJSON `json:"-"`
}

// as112TimeseriesGroupResponseCodesResponseEnvelopeJSON contains the JSON metadata
// for the struct [AS112TimeseriesGroupResponseCodesResponseEnvelope]
type as112TimeseriesGroupResponseCodesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AS112TimeseriesGroupResponseCodesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r as112TimeseriesGroupResponseCodesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
