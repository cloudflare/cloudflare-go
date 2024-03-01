// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarAS112TimeseriesGroupService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAS112TimeseriesGroupService] method instead.
type RadarAS112TimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarAS112TimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAS112TimeseriesGroupService(opts ...option.RequestOption) (r *RadarAS112TimeseriesGroupService) {
	r = &RadarAS112TimeseriesGroupService{}
	r.Options = opts
	return
}

// Percentage distribution of DNS AS112 queries by DNSSEC support over time.
func (r *RadarAS112TimeseriesGroupService) DNSSEC(ctx context.Context, query RadarAS112TimeseriesGroupDNSSECParams, opts ...option.RequestOption) (res *RadarAS112TimeseriesGroupDNSSECResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAS112TimeseriesGroupDNSSECResponseEnvelope
	path := "radar/as112/timeseries_groups/dnssec"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of AS112 DNS queries by EDNS support over time.
func (r *RadarAS112TimeseriesGroupService) Edns(ctx context.Context, query RadarAS112TimeseriesGroupEdnsParams, opts ...option.RequestOption) (res *RadarAS112TimeseriesGroupEdnsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAS112TimeseriesGroupEdnsResponseEnvelope
	path := "radar/as112/timeseries_groups/edns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of AS112 DNS queries by IP Version over time.
func (r *RadarAS112TimeseriesGroupService) IPVersion(ctx context.Context, query RadarAS112TimeseriesGroupIPVersionParams, opts ...option.RequestOption) (res *RadarAS112TimeseriesGroupIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAS112TimeseriesGroupIPVersionResponseEnvelope
	path := "radar/as112/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of AS112 dns requests classified per Protocol over time.
func (r *RadarAS112TimeseriesGroupService) Protocol(ctx context.Context, query RadarAS112TimeseriesGroupProtocolParams, opts ...option.RequestOption) (res *RadarAS112TimeseriesGroupProtocolResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAS112TimeseriesGroupProtocolResponseEnvelope
	path := "radar/as112/timeseries_groups/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of AS112 DNS queries by Query Type over time.
func (r *RadarAS112TimeseriesGroupService) QueryType(ctx context.Context, query RadarAS112TimeseriesGroupQueryTypeParams, opts ...option.RequestOption) (res *RadarAS112TimeseriesGroupQueryTypeResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAS112TimeseriesGroupQueryTypeResponseEnvelope
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
func (r *RadarAS112TimeseriesGroupService) ResponseCodes(ctx context.Context, query RadarAS112TimeseriesGroupResponseCodesParams, opts ...option.RequestOption) (res *RadarAS112TimeseriesGroupResponseCodesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAS112TimeseriesGroupResponseCodesResponseEnvelope
	path := "radar/as112/timeseries_groups/response_codes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAS112TimeseriesGroupDNSSECResponse struct {
	Meta   interface{}                                   `json:"meta,required"`
	Serie0 RadarAS112TimeseriesGroupDNSSECResponseSerie0 `json:"serie_0,required"`
	JSON   radarAS112TimeseriesGroupDNSSECResponseJSON   `json:"-"`
}

// radarAS112TimeseriesGroupDNSSECResponseJSON contains the JSON metadata for the
// struct [RadarAS112TimeseriesGroupDNSSECResponse]
type radarAS112TimeseriesGroupDNSSECResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesGroupDNSSECResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAS112TimeseriesGroupDNSSECResponseSerie0 struct {
	NotSupported []string                                          `json:"NOT_SUPPORTED,required"`
	Supported    []string                                          `json:"SUPPORTED,required"`
	JSON         radarAS112TimeseriesGroupDNSSECResponseSerie0JSON `json:"-"`
}

// radarAS112TimeseriesGroupDNSSECResponseSerie0JSON contains the JSON metadata for
// the struct [RadarAS112TimeseriesGroupDNSSECResponseSerie0]
type radarAS112TimeseriesGroupDNSSECResponseSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAS112TimeseriesGroupDNSSECResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAS112TimeseriesGroupEdnsResponse struct {
	Meta   interface{}                                 `json:"meta,required"`
	Serie0 RadarAS112TimeseriesGroupEdnsResponseSerie0 `json:"serie_0,required"`
	JSON   radarAS112TimeseriesGroupEdnsResponseJSON   `json:"-"`
}

// radarAS112TimeseriesGroupEdnsResponseJSON contains the JSON metadata for the
// struct [RadarAS112TimeseriesGroupEdnsResponse]
type radarAS112TimeseriesGroupEdnsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesGroupEdnsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAS112TimeseriesGroupEdnsResponseSerie0 struct {
	NotSupported []string                                        `json:"NOT_SUPPORTED,required"`
	Supported    []string                                        `json:"SUPPORTED,required"`
	JSON         radarAS112TimeseriesGroupEdnsResponseSerie0JSON `json:"-"`
}

// radarAS112TimeseriesGroupEdnsResponseSerie0JSON contains the JSON metadata for
// the struct [RadarAS112TimeseriesGroupEdnsResponseSerie0]
type radarAS112TimeseriesGroupEdnsResponseSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAS112TimeseriesGroupEdnsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAS112TimeseriesGroupIPVersionResponse struct {
	Meta   interface{}                                      `json:"meta,required"`
	Serie0 RadarAS112TimeseriesGroupIPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   radarAS112TimeseriesGroupIPVersionResponseJSON   `json:"-"`
}

// radarAS112TimeseriesGroupIPVersionResponseJSON contains the JSON metadata for
// the struct [RadarAS112TimeseriesGroupIPVersionResponse]
type radarAS112TimeseriesGroupIPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesGroupIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAS112TimeseriesGroupIPVersionResponseSerie0 struct {
	IPv4 []string                                             `json:"IPv4,required"`
	IPv6 []string                                             `json:"IPv6,required"`
	JSON radarAS112TimeseriesGroupIPVersionResponseSerie0JSON `json:"-"`
}

// radarAS112TimeseriesGroupIPVersionResponseSerie0JSON contains the JSON metadata
// for the struct [RadarAS112TimeseriesGroupIPVersionResponseSerie0]
type radarAS112TimeseriesGroupIPVersionResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesGroupIPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAS112TimeseriesGroupProtocolResponse struct {
	Meta   interface{}                                     `json:"meta,required"`
	Serie0 RadarAS112TimeseriesGroupProtocolResponseSerie0 `json:"serie_0,required"`
	JSON   radarAS112TimeseriesGroupProtocolResponseJSON   `json:"-"`
}

// radarAS112TimeseriesGroupProtocolResponseJSON contains the JSON metadata for the
// struct [RadarAS112TimeseriesGroupProtocolResponse]
type radarAS112TimeseriesGroupProtocolResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesGroupProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAS112TimeseriesGroupProtocolResponseSerie0 struct {
	Tcp  []string                                            `json:"tcp,required"`
	Udp  []string                                            `json:"udp,required"`
	JSON radarAS112TimeseriesGroupProtocolResponseSerie0JSON `json:"-"`
}

// radarAS112TimeseriesGroupProtocolResponseSerie0JSON contains the JSON metadata
// for the struct [RadarAS112TimeseriesGroupProtocolResponseSerie0]
type radarAS112TimeseriesGroupProtocolResponseSerie0JSON struct {
	Tcp         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesGroupProtocolResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAS112TimeseriesGroupQueryTypeResponse struct {
	Meta   interface{}                                      `json:"meta,required"`
	Serie0 RadarAS112TimeseriesGroupQueryTypeResponseSerie0 `json:"serie_0,required"`
	JSON   radarAS112TimeseriesGroupQueryTypeResponseJSON   `json:"-"`
}

// radarAS112TimeseriesGroupQueryTypeResponseJSON contains the JSON metadata for
// the struct [RadarAS112TimeseriesGroupQueryTypeResponse]
type radarAS112TimeseriesGroupQueryTypeResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesGroupQueryTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAS112TimeseriesGroupQueryTypeResponseSerie0 struct {
	A    []string                                             `json:"A,required"`
	Aaaa []string                                             `json:"AAAA,required"`
	Ptr  []string                                             `json:"PTR,required"`
	Soa  []string                                             `json:"SOA,required"`
	Srv  []string                                             `json:"SRV,required"`
	JSON radarAS112TimeseriesGroupQueryTypeResponseSerie0JSON `json:"-"`
}

// radarAS112TimeseriesGroupQueryTypeResponseSerie0JSON contains the JSON metadata
// for the struct [RadarAS112TimeseriesGroupQueryTypeResponseSerie0]
type radarAS112TimeseriesGroupQueryTypeResponseSerie0JSON struct {
	A           apijson.Field
	Aaaa        apijson.Field
	Ptr         apijson.Field
	Soa         apijson.Field
	Srv         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesGroupQueryTypeResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAS112TimeseriesGroupResponseCodesResponse struct {
	Meta   interface{}                                          `json:"meta,required"`
	Serie0 RadarAS112TimeseriesGroupResponseCodesResponseSerie0 `json:"serie_0,required"`
	JSON   radarAS112TimeseriesGroupResponseCodesResponseJSON   `json:"-"`
}

// radarAS112TimeseriesGroupResponseCodesResponseJSON contains the JSON metadata
// for the struct [RadarAS112TimeseriesGroupResponseCodesResponse]
type radarAS112TimeseriesGroupResponseCodesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesGroupResponseCodesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAS112TimeseriesGroupResponseCodesResponseSerie0 struct {
	Noerror  []string                                                 `json:"NOERROR,required"`
	Nxdomain []string                                                 `json:"NXDOMAIN,required"`
	JSON     radarAS112TimeseriesGroupResponseCodesResponseSerie0JSON `json:"-"`
}

// radarAS112TimeseriesGroupResponseCodesResponseSerie0JSON contains the JSON
// metadata for the struct [RadarAS112TimeseriesGroupResponseCodesResponseSerie0]
type radarAS112TimeseriesGroupResponseCodesResponseSerie0JSON struct {
	Noerror     apijson.Field
	Nxdomain    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesGroupResponseCodesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAS112TimeseriesGroupDNSSECParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAS112TimeseriesGroupDNSSECParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]RadarAS112TimeseriesGroupDNSSECParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAS112TimeseriesGroupDNSSECParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAS112TimeseriesGroupDNSSECParams]'s query parameters
// as `url.Values`.
func (r RadarAS112TimeseriesGroupDNSSECParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAS112TimeseriesGroupDNSSECParamsAggInterval string

const (
	RadarAS112TimeseriesGroupDNSSECParamsAggInterval15m RadarAS112TimeseriesGroupDNSSECParamsAggInterval = "15m"
	RadarAS112TimeseriesGroupDNSSECParamsAggInterval1h  RadarAS112TimeseriesGroupDNSSECParamsAggInterval = "1h"
	RadarAS112TimeseriesGroupDNSSECParamsAggInterval1d  RadarAS112TimeseriesGroupDNSSECParamsAggInterval = "1d"
	RadarAS112TimeseriesGroupDNSSECParamsAggInterval1w  RadarAS112TimeseriesGroupDNSSECParamsAggInterval = "1w"
)

type RadarAS112TimeseriesGroupDNSSECParamsDateRange string

const (
	RadarAS112TimeseriesGroupDNSSECParamsDateRange1d         RadarAS112TimeseriesGroupDNSSECParamsDateRange = "1d"
	RadarAS112TimeseriesGroupDNSSECParamsDateRange2d         RadarAS112TimeseriesGroupDNSSECParamsDateRange = "2d"
	RadarAS112TimeseriesGroupDNSSECParamsDateRange7d         RadarAS112TimeseriesGroupDNSSECParamsDateRange = "7d"
	RadarAS112TimeseriesGroupDNSSECParamsDateRange14d        RadarAS112TimeseriesGroupDNSSECParamsDateRange = "14d"
	RadarAS112TimeseriesGroupDNSSECParamsDateRange28d        RadarAS112TimeseriesGroupDNSSECParamsDateRange = "28d"
	RadarAS112TimeseriesGroupDNSSECParamsDateRange12w        RadarAS112TimeseriesGroupDNSSECParamsDateRange = "12w"
	RadarAS112TimeseriesGroupDNSSECParamsDateRange24w        RadarAS112TimeseriesGroupDNSSECParamsDateRange = "24w"
	RadarAS112TimeseriesGroupDNSSECParamsDateRange52w        RadarAS112TimeseriesGroupDNSSECParamsDateRange = "52w"
	RadarAS112TimeseriesGroupDNSSECParamsDateRange1dControl  RadarAS112TimeseriesGroupDNSSECParamsDateRange = "1dControl"
	RadarAS112TimeseriesGroupDNSSECParamsDateRange2dControl  RadarAS112TimeseriesGroupDNSSECParamsDateRange = "2dControl"
	RadarAS112TimeseriesGroupDNSSECParamsDateRange7dControl  RadarAS112TimeseriesGroupDNSSECParamsDateRange = "7dControl"
	RadarAS112TimeseriesGroupDNSSECParamsDateRange14dControl RadarAS112TimeseriesGroupDNSSECParamsDateRange = "14dControl"
	RadarAS112TimeseriesGroupDNSSECParamsDateRange28dControl RadarAS112TimeseriesGroupDNSSECParamsDateRange = "28dControl"
	RadarAS112TimeseriesGroupDNSSECParamsDateRange12wControl RadarAS112TimeseriesGroupDNSSECParamsDateRange = "12wControl"
	RadarAS112TimeseriesGroupDNSSECParamsDateRange24wControl RadarAS112TimeseriesGroupDNSSECParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAS112TimeseriesGroupDNSSECParamsFormat string

const (
	RadarAS112TimeseriesGroupDNSSECParamsFormatJson RadarAS112TimeseriesGroupDNSSECParamsFormat = "JSON"
	RadarAS112TimeseriesGroupDNSSECParamsFormatCsv  RadarAS112TimeseriesGroupDNSSECParamsFormat = "CSV"
)

type RadarAS112TimeseriesGroupDNSSECResponseEnvelope struct {
	Result  RadarAS112TimeseriesGroupDNSSECResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarAS112TimeseriesGroupDNSSECResponseEnvelopeJSON `json:"-"`
}

// radarAS112TimeseriesGroupDNSSECResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAS112TimeseriesGroupDNSSECResponseEnvelope]
type radarAS112TimeseriesGroupDNSSECResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesGroupDNSSECResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAS112TimeseriesGroupEdnsParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAS112TimeseriesGroupEdnsParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]RadarAS112TimeseriesGroupEdnsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAS112TimeseriesGroupEdnsParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAS112TimeseriesGroupEdnsParams]'s query parameters as
// `url.Values`.
func (r RadarAS112TimeseriesGroupEdnsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAS112TimeseriesGroupEdnsParamsAggInterval string

const (
	RadarAS112TimeseriesGroupEdnsParamsAggInterval15m RadarAS112TimeseriesGroupEdnsParamsAggInterval = "15m"
	RadarAS112TimeseriesGroupEdnsParamsAggInterval1h  RadarAS112TimeseriesGroupEdnsParamsAggInterval = "1h"
	RadarAS112TimeseriesGroupEdnsParamsAggInterval1d  RadarAS112TimeseriesGroupEdnsParamsAggInterval = "1d"
	RadarAS112TimeseriesGroupEdnsParamsAggInterval1w  RadarAS112TimeseriesGroupEdnsParamsAggInterval = "1w"
)

type RadarAS112TimeseriesGroupEdnsParamsDateRange string

const (
	RadarAS112TimeseriesGroupEdnsParamsDateRange1d         RadarAS112TimeseriesGroupEdnsParamsDateRange = "1d"
	RadarAS112TimeseriesGroupEdnsParamsDateRange2d         RadarAS112TimeseriesGroupEdnsParamsDateRange = "2d"
	RadarAS112TimeseriesGroupEdnsParamsDateRange7d         RadarAS112TimeseriesGroupEdnsParamsDateRange = "7d"
	RadarAS112TimeseriesGroupEdnsParamsDateRange14d        RadarAS112TimeseriesGroupEdnsParamsDateRange = "14d"
	RadarAS112TimeseriesGroupEdnsParamsDateRange28d        RadarAS112TimeseriesGroupEdnsParamsDateRange = "28d"
	RadarAS112TimeseriesGroupEdnsParamsDateRange12w        RadarAS112TimeseriesGroupEdnsParamsDateRange = "12w"
	RadarAS112TimeseriesGroupEdnsParamsDateRange24w        RadarAS112TimeseriesGroupEdnsParamsDateRange = "24w"
	RadarAS112TimeseriesGroupEdnsParamsDateRange52w        RadarAS112TimeseriesGroupEdnsParamsDateRange = "52w"
	RadarAS112TimeseriesGroupEdnsParamsDateRange1dControl  RadarAS112TimeseriesGroupEdnsParamsDateRange = "1dControl"
	RadarAS112TimeseriesGroupEdnsParamsDateRange2dControl  RadarAS112TimeseriesGroupEdnsParamsDateRange = "2dControl"
	RadarAS112TimeseriesGroupEdnsParamsDateRange7dControl  RadarAS112TimeseriesGroupEdnsParamsDateRange = "7dControl"
	RadarAS112TimeseriesGroupEdnsParamsDateRange14dControl RadarAS112TimeseriesGroupEdnsParamsDateRange = "14dControl"
	RadarAS112TimeseriesGroupEdnsParamsDateRange28dControl RadarAS112TimeseriesGroupEdnsParamsDateRange = "28dControl"
	RadarAS112TimeseriesGroupEdnsParamsDateRange12wControl RadarAS112TimeseriesGroupEdnsParamsDateRange = "12wControl"
	RadarAS112TimeseriesGroupEdnsParamsDateRange24wControl RadarAS112TimeseriesGroupEdnsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAS112TimeseriesGroupEdnsParamsFormat string

const (
	RadarAS112TimeseriesGroupEdnsParamsFormatJson RadarAS112TimeseriesGroupEdnsParamsFormat = "JSON"
	RadarAS112TimeseriesGroupEdnsParamsFormatCsv  RadarAS112TimeseriesGroupEdnsParamsFormat = "CSV"
)

type RadarAS112TimeseriesGroupEdnsResponseEnvelope struct {
	Result  RadarAS112TimeseriesGroupEdnsResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarAS112TimeseriesGroupEdnsResponseEnvelopeJSON `json:"-"`
}

// radarAS112TimeseriesGroupEdnsResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarAS112TimeseriesGroupEdnsResponseEnvelope]
type radarAS112TimeseriesGroupEdnsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesGroupEdnsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAS112TimeseriesGroupIPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAS112TimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]RadarAS112TimeseriesGroupIPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAS112TimeseriesGroupIPVersionParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAS112TimeseriesGroupIPVersionParams]'s query
// parameters as `url.Values`.
func (r RadarAS112TimeseriesGroupIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAS112TimeseriesGroupIPVersionParamsAggInterval string

const (
	RadarAS112TimeseriesGroupIPVersionParamsAggInterval15m RadarAS112TimeseriesGroupIPVersionParamsAggInterval = "15m"
	RadarAS112TimeseriesGroupIPVersionParamsAggInterval1h  RadarAS112TimeseriesGroupIPVersionParamsAggInterval = "1h"
	RadarAS112TimeseriesGroupIPVersionParamsAggInterval1d  RadarAS112TimeseriesGroupIPVersionParamsAggInterval = "1d"
	RadarAS112TimeseriesGroupIPVersionParamsAggInterval1w  RadarAS112TimeseriesGroupIPVersionParamsAggInterval = "1w"
)

type RadarAS112TimeseriesGroupIPVersionParamsDateRange string

const (
	RadarAS112TimeseriesGroupIPVersionParamsDateRange1d         RadarAS112TimeseriesGroupIPVersionParamsDateRange = "1d"
	RadarAS112TimeseriesGroupIPVersionParamsDateRange2d         RadarAS112TimeseriesGroupIPVersionParamsDateRange = "2d"
	RadarAS112TimeseriesGroupIPVersionParamsDateRange7d         RadarAS112TimeseriesGroupIPVersionParamsDateRange = "7d"
	RadarAS112TimeseriesGroupIPVersionParamsDateRange14d        RadarAS112TimeseriesGroupIPVersionParamsDateRange = "14d"
	RadarAS112TimeseriesGroupIPVersionParamsDateRange28d        RadarAS112TimeseriesGroupIPVersionParamsDateRange = "28d"
	RadarAS112TimeseriesGroupIPVersionParamsDateRange12w        RadarAS112TimeseriesGroupIPVersionParamsDateRange = "12w"
	RadarAS112TimeseriesGroupIPVersionParamsDateRange24w        RadarAS112TimeseriesGroupIPVersionParamsDateRange = "24w"
	RadarAS112TimeseriesGroupIPVersionParamsDateRange52w        RadarAS112TimeseriesGroupIPVersionParamsDateRange = "52w"
	RadarAS112TimeseriesGroupIPVersionParamsDateRange1dControl  RadarAS112TimeseriesGroupIPVersionParamsDateRange = "1dControl"
	RadarAS112TimeseriesGroupIPVersionParamsDateRange2dControl  RadarAS112TimeseriesGroupIPVersionParamsDateRange = "2dControl"
	RadarAS112TimeseriesGroupIPVersionParamsDateRange7dControl  RadarAS112TimeseriesGroupIPVersionParamsDateRange = "7dControl"
	RadarAS112TimeseriesGroupIPVersionParamsDateRange14dControl RadarAS112TimeseriesGroupIPVersionParamsDateRange = "14dControl"
	RadarAS112TimeseriesGroupIPVersionParamsDateRange28dControl RadarAS112TimeseriesGroupIPVersionParamsDateRange = "28dControl"
	RadarAS112TimeseriesGroupIPVersionParamsDateRange12wControl RadarAS112TimeseriesGroupIPVersionParamsDateRange = "12wControl"
	RadarAS112TimeseriesGroupIPVersionParamsDateRange24wControl RadarAS112TimeseriesGroupIPVersionParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAS112TimeseriesGroupIPVersionParamsFormat string

const (
	RadarAS112TimeseriesGroupIPVersionParamsFormatJson RadarAS112TimeseriesGroupIPVersionParamsFormat = "JSON"
	RadarAS112TimeseriesGroupIPVersionParamsFormatCsv  RadarAS112TimeseriesGroupIPVersionParamsFormat = "CSV"
)

type RadarAS112TimeseriesGroupIPVersionResponseEnvelope struct {
	Result  RadarAS112TimeseriesGroupIPVersionResponse             `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarAS112TimeseriesGroupIPVersionResponseEnvelopeJSON `json:"-"`
}

// radarAS112TimeseriesGroupIPVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAS112TimeseriesGroupIPVersionResponseEnvelope]
type radarAS112TimeseriesGroupIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesGroupIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAS112TimeseriesGroupProtocolParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAS112TimeseriesGroupProtocolParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]RadarAS112TimeseriesGroupProtocolParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAS112TimeseriesGroupProtocolParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAS112TimeseriesGroupProtocolParams]'s query parameters
// as `url.Values`.
func (r RadarAS112TimeseriesGroupProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAS112TimeseriesGroupProtocolParamsAggInterval string

const (
	RadarAS112TimeseriesGroupProtocolParamsAggInterval15m RadarAS112TimeseriesGroupProtocolParamsAggInterval = "15m"
	RadarAS112TimeseriesGroupProtocolParamsAggInterval1h  RadarAS112TimeseriesGroupProtocolParamsAggInterval = "1h"
	RadarAS112TimeseriesGroupProtocolParamsAggInterval1d  RadarAS112TimeseriesGroupProtocolParamsAggInterval = "1d"
	RadarAS112TimeseriesGroupProtocolParamsAggInterval1w  RadarAS112TimeseriesGroupProtocolParamsAggInterval = "1w"
)

type RadarAS112TimeseriesGroupProtocolParamsDateRange string

const (
	RadarAS112TimeseriesGroupProtocolParamsDateRange1d         RadarAS112TimeseriesGroupProtocolParamsDateRange = "1d"
	RadarAS112TimeseriesGroupProtocolParamsDateRange2d         RadarAS112TimeseriesGroupProtocolParamsDateRange = "2d"
	RadarAS112TimeseriesGroupProtocolParamsDateRange7d         RadarAS112TimeseriesGroupProtocolParamsDateRange = "7d"
	RadarAS112TimeseriesGroupProtocolParamsDateRange14d        RadarAS112TimeseriesGroupProtocolParamsDateRange = "14d"
	RadarAS112TimeseriesGroupProtocolParamsDateRange28d        RadarAS112TimeseriesGroupProtocolParamsDateRange = "28d"
	RadarAS112TimeseriesGroupProtocolParamsDateRange12w        RadarAS112TimeseriesGroupProtocolParamsDateRange = "12w"
	RadarAS112TimeseriesGroupProtocolParamsDateRange24w        RadarAS112TimeseriesGroupProtocolParamsDateRange = "24w"
	RadarAS112TimeseriesGroupProtocolParamsDateRange52w        RadarAS112TimeseriesGroupProtocolParamsDateRange = "52w"
	RadarAS112TimeseriesGroupProtocolParamsDateRange1dControl  RadarAS112TimeseriesGroupProtocolParamsDateRange = "1dControl"
	RadarAS112TimeseriesGroupProtocolParamsDateRange2dControl  RadarAS112TimeseriesGroupProtocolParamsDateRange = "2dControl"
	RadarAS112TimeseriesGroupProtocolParamsDateRange7dControl  RadarAS112TimeseriesGroupProtocolParamsDateRange = "7dControl"
	RadarAS112TimeseriesGroupProtocolParamsDateRange14dControl RadarAS112TimeseriesGroupProtocolParamsDateRange = "14dControl"
	RadarAS112TimeseriesGroupProtocolParamsDateRange28dControl RadarAS112TimeseriesGroupProtocolParamsDateRange = "28dControl"
	RadarAS112TimeseriesGroupProtocolParamsDateRange12wControl RadarAS112TimeseriesGroupProtocolParamsDateRange = "12wControl"
	RadarAS112TimeseriesGroupProtocolParamsDateRange24wControl RadarAS112TimeseriesGroupProtocolParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAS112TimeseriesGroupProtocolParamsFormat string

const (
	RadarAS112TimeseriesGroupProtocolParamsFormatJson RadarAS112TimeseriesGroupProtocolParamsFormat = "JSON"
	RadarAS112TimeseriesGroupProtocolParamsFormatCsv  RadarAS112TimeseriesGroupProtocolParamsFormat = "CSV"
)

type RadarAS112TimeseriesGroupProtocolResponseEnvelope struct {
	Result  RadarAS112TimeseriesGroupProtocolResponse             `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarAS112TimeseriesGroupProtocolResponseEnvelopeJSON `json:"-"`
}

// radarAS112TimeseriesGroupProtocolResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAS112TimeseriesGroupProtocolResponseEnvelope]
type radarAS112TimeseriesGroupProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesGroupProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAS112TimeseriesGroupQueryTypeParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAS112TimeseriesGroupQueryTypeParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]RadarAS112TimeseriesGroupQueryTypeParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAS112TimeseriesGroupQueryTypeParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAS112TimeseriesGroupQueryTypeParams]'s query
// parameters as `url.Values`.
func (r RadarAS112TimeseriesGroupQueryTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAS112TimeseriesGroupQueryTypeParamsAggInterval string

const (
	RadarAS112TimeseriesGroupQueryTypeParamsAggInterval15m RadarAS112TimeseriesGroupQueryTypeParamsAggInterval = "15m"
	RadarAS112TimeseriesGroupQueryTypeParamsAggInterval1h  RadarAS112TimeseriesGroupQueryTypeParamsAggInterval = "1h"
	RadarAS112TimeseriesGroupQueryTypeParamsAggInterval1d  RadarAS112TimeseriesGroupQueryTypeParamsAggInterval = "1d"
	RadarAS112TimeseriesGroupQueryTypeParamsAggInterval1w  RadarAS112TimeseriesGroupQueryTypeParamsAggInterval = "1w"
)

type RadarAS112TimeseriesGroupQueryTypeParamsDateRange string

const (
	RadarAS112TimeseriesGroupQueryTypeParamsDateRange1d         RadarAS112TimeseriesGroupQueryTypeParamsDateRange = "1d"
	RadarAS112TimeseriesGroupQueryTypeParamsDateRange2d         RadarAS112TimeseriesGroupQueryTypeParamsDateRange = "2d"
	RadarAS112TimeseriesGroupQueryTypeParamsDateRange7d         RadarAS112TimeseriesGroupQueryTypeParamsDateRange = "7d"
	RadarAS112TimeseriesGroupQueryTypeParamsDateRange14d        RadarAS112TimeseriesGroupQueryTypeParamsDateRange = "14d"
	RadarAS112TimeseriesGroupQueryTypeParamsDateRange28d        RadarAS112TimeseriesGroupQueryTypeParamsDateRange = "28d"
	RadarAS112TimeseriesGroupQueryTypeParamsDateRange12w        RadarAS112TimeseriesGroupQueryTypeParamsDateRange = "12w"
	RadarAS112TimeseriesGroupQueryTypeParamsDateRange24w        RadarAS112TimeseriesGroupQueryTypeParamsDateRange = "24w"
	RadarAS112TimeseriesGroupQueryTypeParamsDateRange52w        RadarAS112TimeseriesGroupQueryTypeParamsDateRange = "52w"
	RadarAS112TimeseriesGroupQueryTypeParamsDateRange1dControl  RadarAS112TimeseriesGroupQueryTypeParamsDateRange = "1dControl"
	RadarAS112TimeseriesGroupQueryTypeParamsDateRange2dControl  RadarAS112TimeseriesGroupQueryTypeParamsDateRange = "2dControl"
	RadarAS112TimeseriesGroupQueryTypeParamsDateRange7dControl  RadarAS112TimeseriesGroupQueryTypeParamsDateRange = "7dControl"
	RadarAS112TimeseriesGroupQueryTypeParamsDateRange14dControl RadarAS112TimeseriesGroupQueryTypeParamsDateRange = "14dControl"
	RadarAS112TimeseriesGroupQueryTypeParamsDateRange28dControl RadarAS112TimeseriesGroupQueryTypeParamsDateRange = "28dControl"
	RadarAS112TimeseriesGroupQueryTypeParamsDateRange12wControl RadarAS112TimeseriesGroupQueryTypeParamsDateRange = "12wControl"
	RadarAS112TimeseriesGroupQueryTypeParamsDateRange24wControl RadarAS112TimeseriesGroupQueryTypeParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAS112TimeseriesGroupQueryTypeParamsFormat string

const (
	RadarAS112TimeseriesGroupQueryTypeParamsFormatJson RadarAS112TimeseriesGroupQueryTypeParamsFormat = "JSON"
	RadarAS112TimeseriesGroupQueryTypeParamsFormatCsv  RadarAS112TimeseriesGroupQueryTypeParamsFormat = "CSV"
)

type RadarAS112TimeseriesGroupQueryTypeResponseEnvelope struct {
	Result  RadarAS112TimeseriesGroupQueryTypeResponse             `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarAS112TimeseriesGroupQueryTypeResponseEnvelopeJSON `json:"-"`
}

// radarAS112TimeseriesGroupQueryTypeResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAS112TimeseriesGroupQueryTypeResponseEnvelope]
type radarAS112TimeseriesGroupQueryTypeResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesGroupQueryTypeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAS112TimeseriesGroupResponseCodesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAS112TimeseriesGroupResponseCodesParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]RadarAS112TimeseriesGroupResponseCodesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAS112TimeseriesGroupResponseCodesParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAS112TimeseriesGroupResponseCodesParams]'s query
// parameters as `url.Values`.
func (r RadarAS112TimeseriesGroupResponseCodesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAS112TimeseriesGroupResponseCodesParamsAggInterval string

const (
	RadarAS112TimeseriesGroupResponseCodesParamsAggInterval15m RadarAS112TimeseriesGroupResponseCodesParamsAggInterval = "15m"
	RadarAS112TimeseriesGroupResponseCodesParamsAggInterval1h  RadarAS112TimeseriesGroupResponseCodesParamsAggInterval = "1h"
	RadarAS112TimeseriesGroupResponseCodesParamsAggInterval1d  RadarAS112TimeseriesGroupResponseCodesParamsAggInterval = "1d"
	RadarAS112TimeseriesGroupResponseCodesParamsAggInterval1w  RadarAS112TimeseriesGroupResponseCodesParamsAggInterval = "1w"
)

type RadarAS112TimeseriesGroupResponseCodesParamsDateRange string

const (
	RadarAS112TimeseriesGroupResponseCodesParamsDateRange1d         RadarAS112TimeseriesGroupResponseCodesParamsDateRange = "1d"
	RadarAS112TimeseriesGroupResponseCodesParamsDateRange2d         RadarAS112TimeseriesGroupResponseCodesParamsDateRange = "2d"
	RadarAS112TimeseriesGroupResponseCodesParamsDateRange7d         RadarAS112TimeseriesGroupResponseCodesParamsDateRange = "7d"
	RadarAS112TimeseriesGroupResponseCodesParamsDateRange14d        RadarAS112TimeseriesGroupResponseCodesParamsDateRange = "14d"
	RadarAS112TimeseriesGroupResponseCodesParamsDateRange28d        RadarAS112TimeseriesGroupResponseCodesParamsDateRange = "28d"
	RadarAS112TimeseriesGroupResponseCodesParamsDateRange12w        RadarAS112TimeseriesGroupResponseCodesParamsDateRange = "12w"
	RadarAS112TimeseriesGroupResponseCodesParamsDateRange24w        RadarAS112TimeseriesGroupResponseCodesParamsDateRange = "24w"
	RadarAS112TimeseriesGroupResponseCodesParamsDateRange52w        RadarAS112TimeseriesGroupResponseCodesParamsDateRange = "52w"
	RadarAS112TimeseriesGroupResponseCodesParamsDateRange1dControl  RadarAS112TimeseriesGroupResponseCodesParamsDateRange = "1dControl"
	RadarAS112TimeseriesGroupResponseCodesParamsDateRange2dControl  RadarAS112TimeseriesGroupResponseCodesParamsDateRange = "2dControl"
	RadarAS112TimeseriesGroupResponseCodesParamsDateRange7dControl  RadarAS112TimeseriesGroupResponseCodesParamsDateRange = "7dControl"
	RadarAS112TimeseriesGroupResponseCodesParamsDateRange14dControl RadarAS112TimeseriesGroupResponseCodesParamsDateRange = "14dControl"
	RadarAS112TimeseriesGroupResponseCodesParamsDateRange28dControl RadarAS112TimeseriesGroupResponseCodesParamsDateRange = "28dControl"
	RadarAS112TimeseriesGroupResponseCodesParamsDateRange12wControl RadarAS112TimeseriesGroupResponseCodesParamsDateRange = "12wControl"
	RadarAS112TimeseriesGroupResponseCodesParamsDateRange24wControl RadarAS112TimeseriesGroupResponseCodesParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAS112TimeseriesGroupResponseCodesParamsFormat string

const (
	RadarAS112TimeseriesGroupResponseCodesParamsFormatJson RadarAS112TimeseriesGroupResponseCodesParamsFormat = "JSON"
	RadarAS112TimeseriesGroupResponseCodesParamsFormatCsv  RadarAS112TimeseriesGroupResponseCodesParamsFormat = "CSV"
)

type RadarAS112TimeseriesGroupResponseCodesResponseEnvelope struct {
	Result  RadarAS112TimeseriesGroupResponseCodesResponse             `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    radarAS112TimeseriesGroupResponseCodesResponseEnvelopeJSON `json:"-"`
}

// radarAS112TimeseriesGroupResponseCodesResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAS112TimeseriesGroupResponseCodesResponseEnvelope]
type radarAS112TimeseriesGroupResponseCodesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesGroupResponseCodesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
