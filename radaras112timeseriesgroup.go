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

// RadarAs112TimeseriesGroupService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAs112TimeseriesGroupService] method instead.
type RadarAs112TimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarAs112TimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAs112TimeseriesGroupService(opts ...option.RequestOption) (r *RadarAs112TimeseriesGroupService) {
	r = &RadarAs112TimeseriesGroupService{}
	r.Options = opts
	return
}

// Percentage distribution of DNS AS112 queries by DNSSEC support over time.
func (r *RadarAs112TimeseriesGroupService) DNSSEC(ctx context.Context, query RadarAs112TimeseriesGroupDNSSECParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupDNSSECResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112TimeseriesGroupDNSSECResponseEnvelope
	path := "radar/as112/timeseries_groups/dnssec"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of AS112 DNS queries by EDNS support over time.
func (r *RadarAs112TimeseriesGroupService) Edns(ctx context.Context, query RadarAs112TimeseriesGroupEdnsParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupEdnsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112TimeseriesGroupEdnsResponseEnvelope
	path := "radar/as112/timeseries_groups/edns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of AS112 DNS queries by IP Version over time.
func (r *RadarAs112TimeseriesGroupService) IPVersion(ctx context.Context, query RadarAs112TimeseriesGroupIPVersionParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112TimeseriesGroupIPVersionResponseEnvelope
	path := "radar/as112/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of AS112 dns requests classified per Protocol over time.
func (r *RadarAs112TimeseriesGroupService) Protocol(ctx context.Context, query RadarAs112TimeseriesGroupProtocolParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupProtocolResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112TimeseriesGroupProtocolResponseEnvelope
	path := "radar/as112/timeseries_groups/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of AS112 DNS queries by Query Type over time.
func (r *RadarAs112TimeseriesGroupService) QueryType(ctx context.Context, query RadarAs112TimeseriesGroupQueryTypeParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupQueryTypeResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112TimeseriesGroupQueryTypeResponseEnvelope
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
func (r *RadarAs112TimeseriesGroupService) ResponseCodes(ctx context.Context, query RadarAs112TimeseriesGroupResponseCodesParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupResponseCodesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112TimeseriesGroupResponseCodesResponseEnvelope
	path := "radar/as112/timeseries_groups/response_codes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAs112TimeseriesGroupDNSSECResponse struct {
	Meta   interface{}                                   `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupDNSSECResponseSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupDNSSECResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupDNSSECResponseJSON contains the JSON metadata for the
// struct [RadarAs112TimeseriesGroupDNSSECResponse]
type radarAs112TimeseriesGroupDNSSECResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupDNSSECResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupDNSSECResponseSerie0 struct {
	NotSupported []string                                          `json:"NOT_SUPPORTED,required"`
	Supported    []string                                          `json:"SUPPORTED,required"`
	JSON         radarAs112TimeseriesGroupDNSSECResponseSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupDNSSECResponseSerie0JSON contains the JSON metadata for
// the struct [RadarAs112TimeseriesGroupDNSSECResponseSerie0]
type radarAs112TimeseriesGroupDNSSECResponseSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupDNSSECResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupEdnsResponse struct {
	Meta   interface{}                                 `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupEdnsResponseSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupEdnsResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupEdnsResponseJSON contains the JSON metadata for the
// struct [RadarAs112TimeseriesGroupEdnsResponse]
type radarAs112TimeseriesGroupEdnsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupEdnsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupEdnsResponseSerie0 struct {
	NotSupported []string                                        `json:"NOT_SUPPORTED,required"`
	Supported    []string                                        `json:"SUPPORTED,required"`
	JSON         radarAs112TimeseriesGroupEdnsResponseSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupEdnsResponseSerie0JSON contains the JSON metadata for
// the struct [RadarAs112TimeseriesGroupEdnsResponseSerie0]
type radarAs112TimeseriesGroupEdnsResponseSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupEdnsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupIPVersionResponse struct {
	Meta   interface{}                                      `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupIPVersionResponseSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupIPVersionResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupIPVersionResponseJSON contains the JSON metadata for
// the struct [RadarAs112TimeseriesGroupIPVersionResponse]
type radarAs112TimeseriesGroupIPVersionResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupIPVersionResponseSerie0 struct {
	IPv4 []string                                             `json:"IPv4,required"`
	IPv6 []string                                             `json:"IPv6,required"`
	JSON radarAs112TimeseriesGroupIPVersionResponseSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupIPVersionResponseSerie0JSON contains the JSON metadata
// for the struct [RadarAs112TimeseriesGroupIPVersionResponseSerie0]
type radarAs112TimeseriesGroupIPVersionResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupIPVersionResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupProtocolResponse struct {
	Meta   interface{}                                     `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupProtocolResponseSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupProtocolResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupProtocolResponseJSON contains the JSON metadata for the
// struct [RadarAs112TimeseriesGroupProtocolResponse]
type radarAs112TimeseriesGroupProtocolResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupProtocolResponseSerie0 struct {
	Tcp  []string                                            `json:"tcp,required"`
	Udp  []string                                            `json:"udp,required"`
	JSON radarAs112TimeseriesGroupProtocolResponseSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupProtocolResponseSerie0JSON contains the JSON metadata
// for the struct [RadarAs112TimeseriesGroupProtocolResponseSerie0]
type radarAs112TimeseriesGroupProtocolResponseSerie0JSON struct {
	Tcp         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupProtocolResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupQueryTypeResponse struct {
	Meta   interface{}                                      `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupQueryTypeResponseSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupQueryTypeResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupQueryTypeResponseJSON contains the JSON metadata for
// the struct [RadarAs112TimeseriesGroupQueryTypeResponse]
type radarAs112TimeseriesGroupQueryTypeResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupQueryTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupQueryTypeResponseSerie0 struct {
	A    []string                                             `json:"A,required"`
	Aaaa []string                                             `json:"AAAA,required"`
	Ptr  []string                                             `json:"PTR,required"`
	Soa  []string                                             `json:"SOA,required"`
	Srv  []string                                             `json:"SRV,required"`
	JSON radarAs112TimeseriesGroupQueryTypeResponseSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupQueryTypeResponseSerie0JSON contains the JSON metadata
// for the struct [RadarAs112TimeseriesGroupQueryTypeResponseSerie0]
type radarAs112TimeseriesGroupQueryTypeResponseSerie0JSON struct {
	A           apijson.Field
	Aaaa        apijson.Field
	Ptr         apijson.Field
	Soa         apijson.Field
	Srv         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupQueryTypeResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupResponseCodesResponse struct {
	Meta   interface{}                                          `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupResponseCodesResponseSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupResponseCodesResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupResponseCodesResponseJSON contains the JSON metadata
// for the struct [RadarAs112TimeseriesGroupResponseCodesResponse]
type radarAs112TimeseriesGroupResponseCodesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupResponseCodesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupResponseCodesResponseSerie0 struct {
	Noerror  []string                                                 `json:"NOERROR,required"`
	Nxdomain []string                                                 `json:"NXDOMAIN,required"`
	JSON     radarAs112TimeseriesGroupResponseCodesResponseSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupResponseCodesResponseSerie0JSON contains the JSON
// metadata for the struct [RadarAs112TimeseriesGroupResponseCodesResponseSerie0]
type radarAs112TimeseriesGroupResponseCodesResponseSerie0JSON struct {
	Noerror     apijson.Field
	Nxdomain    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupResponseCodesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupDNSSECParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupDNSSECParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseriesGroupDNSSECParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseriesGroupDNSSECParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupDNSSECParams]'s query parameters
// as `url.Values`.
func (r RadarAs112TimeseriesGroupDNSSECParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupDNSSECParamsAggInterval string

const (
	RadarAs112TimeseriesGroupDNSSECParamsAggInterval15m RadarAs112TimeseriesGroupDNSSECParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupDNSSECParamsAggInterval1h  RadarAs112TimeseriesGroupDNSSECParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupDNSSECParamsAggInterval1d  RadarAs112TimeseriesGroupDNSSECParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupDNSSECParamsAggInterval1w  RadarAs112TimeseriesGroupDNSSECParamsAggInterval = "1w"
)

type RadarAs112TimeseriesGroupDNSSECParamsDateRange string

const (
	RadarAs112TimeseriesGroupDNSSECParamsDateRange1d         RadarAs112TimeseriesGroupDNSSECParamsDateRange = "1d"
	RadarAs112TimeseriesGroupDNSSECParamsDateRange2d         RadarAs112TimeseriesGroupDNSSECParamsDateRange = "2d"
	RadarAs112TimeseriesGroupDNSSECParamsDateRange7d         RadarAs112TimeseriesGroupDNSSECParamsDateRange = "7d"
	RadarAs112TimeseriesGroupDNSSECParamsDateRange14d        RadarAs112TimeseriesGroupDNSSECParamsDateRange = "14d"
	RadarAs112TimeseriesGroupDNSSECParamsDateRange28d        RadarAs112TimeseriesGroupDNSSECParamsDateRange = "28d"
	RadarAs112TimeseriesGroupDNSSECParamsDateRange12w        RadarAs112TimeseriesGroupDNSSECParamsDateRange = "12w"
	RadarAs112TimeseriesGroupDNSSECParamsDateRange24w        RadarAs112TimeseriesGroupDNSSECParamsDateRange = "24w"
	RadarAs112TimeseriesGroupDNSSECParamsDateRange52w        RadarAs112TimeseriesGroupDNSSECParamsDateRange = "52w"
	RadarAs112TimeseriesGroupDNSSECParamsDateRange1dControl  RadarAs112TimeseriesGroupDNSSECParamsDateRange = "1dControl"
	RadarAs112TimeseriesGroupDNSSECParamsDateRange2dControl  RadarAs112TimeseriesGroupDNSSECParamsDateRange = "2dControl"
	RadarAs112TimeseriesGroupDNSSECParamsDateRange7dControl  RadarAs112TimeseriesGroupDNSSECParamsDateRange = "7dControl"
	RadarAs112TimeseriesGroupDNSSECParamsDateRange14dControl RadarAs112TimeseriesGroupDNSSECParamsDateRange = "14dControl"
	RadarAs112TimeseriesGroupDNSSECParamsDateRange28dControl RadarAs112TimeseriesGroupDNSSECParamsDateRange = "28dControl"
	RadarAs112TimeseriesGroupDNSSECParamsDateRange12wControl RadarAs112TimeseriesGroupDNSSECParamsDateRange = "12wControl"
	RadarAs112TimeseriesGroupDNSSECParamsDateRange24wControl RadarAs112TimeseriesGroupDNSSECParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseriesGroupDNSSECParamsFormat string

const (
	RadarAs112TimeseriesGroupDNSSECParamsFormatJson RadarAs112TimeseriesGroupDNSSECParamsFormat = "JSON"
	RadarAs112TimeseriesGroupDNSSECParamsFormatCsv  RadarAs112TimeseriesGroupDNSSECParamsFormat = "CSV"
)

type RadarAs112TimeseriesGroupDNSSECResponseEnvelope struct {
	Result  RadarAs112TimeseriesGroupDNSSECResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarAs112TimeseriesGroupDNSSECResponseEnvelopeJSON `json:"-"`
}

// radarAs112TimeseriesGroupDNSSECResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAs112TimeseriesGroupDNSSECResponseEnvelope]
type radarAs112TimeseriesGroupDNSSECResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupDNSSECResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupEdnsParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupEdnsParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseriesGroupEdnsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseriesGroupEdnsParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupEdnsParams]'s query parameters as
// `url.Values`.
func (r RadarAs112TimeseriesGroupEdnsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupEdnsParamsAggInterval string

const (
	RadarAs112TimeseriesGroupEdnsParamsAggInterval15m RadarAs112TimeseriesGroupEdnsParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupEdnsParamsAggInterval1h  RadarAs112TimeseriesGroupEdnsParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupEdnsParamsAggInterval1d  RadarAs112TimeseriesGroupEdnsParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupEdnsParamsAggInterval1w  RadarAs112TimeseriesGroupEdnsParamsAggInterval = "1w"
)

type RadarAs112TimeseriesGroupEdnsParamsDateRange string

const (
	RadarAs112TimeseriesGroupEdnsParamsDateRange1d         RadarAs112TimeseriesGroupEdnsParamsDateRange = "1d"
	RadarAs112TimeseriesGroupEdnsParamsDateRange2d         RadarAs112TimeseriesGroupEdnsParamsDateRange = "2d"
	RadarAs112TimeseriesGroupEdnsParamsDateRange7d         RadarAs112TimeseriesGroupEdnsParamsDateRange = "7d"
	RadarAs112TimeseriesGroupEdnsParamsDateRange14d        RadarAs112TimeseriesGroupEdnsParamsDateRange = "14d"
	RadarAs112TimeseriesGroupEdnsParamsDateRange28d        RadarAs112TimeseriesGroupEdnsParamsDateRange = "28d"
	RadarAs112TimeseriesGroupEdnsParamsDateRange12w        RadarAs112TimeseriesGroupEdnsParamsDateRange = "12w"
	RadarAs112TimeseriesGroupEdnsParamsDateRange24w        RadarAs112TimeseriesGroupEdnsParamsDateRange = "24w"
	RadarAs112TimeseriesGroupEdnsParamsDateRange52w        RadarAs112TimeseriesGroupEdnsParamsDateRange = "52w"
	RadarAs112TimeseriesGroupEdnsParamsDateRange1dControl  RadarAs112TimeseriesGroupEdnsParamsDateRange = "1dControl"
	RadarAs112TimeseriesGroupEdnsParamsDateRange2dControl  RadarAs112TimeseriesGroupEdnsParamsDateRange = "2dControl"
	RadarAs112TimeseriesGroupEdnsParamsDateRange7dControl  RadarAs112TimeseriesGroupEdnsParamsDateRange = "7dControl"
	RadarAs112TimeseriesGroupEdnsParamsDateRange14dControl RadarAs112TimeseriesGroupEdnsParamsDateRange = "14dControl"
	RadarAs112TimeseriesGroupEdnsParamsDateRange28dControl RadarAs112TimeseriesGroupEdnsParamsDateRange = "28dControl"
	RadarAs112TimeseriesGroupEdnsParamsDateRange12wControl RadarAs112TimeseriesGroupEdnsParamsDateRange = "12wControl"
	RadarAs112TimeseriesGroupEdnsParamsDateRange24wControl RadarAs112TimeseriesGroupEdnsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseriesGroupEdnsParamsFormat string

const (
	RadarAs112TimeseriesGroupEdnsParamsFormatJson RadarAs112TimeseriesGroupEdnsParamsFormat = "JSON"
	RadarAs112TimeseriesGroupEdnsParamsFormatCsv  RadarAs112TimeseriesGroupEdnsParamsFormat = "CSV"
)

type RadarAs112TimeseriesGroupEdnsResponseEnvelope struct {
	Result  RadarAs112TimeseriesGroupEdnsResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarAs112TimeseriesGroupEdnsResponseEnvelopeJSON `json:"-"`
}

// radarAs112TimeseriesGroupEdnsResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarAs112TimeseriesGroupEdnsResponseEnvelope]
type radarAs112TimeseriesGroupEdnsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupEdnsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupIPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupIPVersionParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseriesGroupIPVersionParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseriesGroupIPVersionParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupIPVersionParams]'s query
// parameters as `url.Values`.
func (r RadarAs112TimeseriesGroupIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupIPVersionParamsAggInterval string

const (
	RadarAs112TimeseriesGroupIPVersionParamsAggInterval15m RadarAs112TimeseriesGroupIPVersionParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupIPVersionParamsAggInterval1h  RadarAs112TimeseriesGroupIPVersionParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupIPVersionParamsAggInterval1d  RadarAs112TimeseriesGroupIPVersionParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupIPVersionParamsAggInterval1w  RadarAs112TimeseriesGroupIPVersionParamsAggInterval = "1w"
)

type RadarAs112TimeseriesGroupIPVersionParamsDateRange string

const (
	RadarAs112TimeseriesGroupIPVersionParamsDateRange1d         RadarAs112TimeseriesGroupIPVersionParamsDateRange = "1d"
	RadarAs112TimeseriesGroupIPVersionParamsDateRange2d         RadarAs112TimeseriesGroupIPVersionParamsDateRange = "2d"
	RadarAs112TimeseriesGroupIPVersionParamsDateRange7d         RadarAs112TimeseriesGroupIPVersionParamsDateRange = "7d"
	RadarAs112TimeseriesGroupIPVersionParamsDateRange14d        RadarAs112TimeseriesGroupIPVersionParamsDateRange = "14d"
	RadarAs112TimeseriesGroupIPVersionParamsDateRange28d        RadarAs112TimeseriesGroupIPVersionParamsDateRange = "28d"
	RadarAs112TimeseriesGroupIPVersionParamsDateRange12w        RadarAs112TimeseriesGroupIPVersionParamsDateRange = "12w"
	RadarAs112TimeseriesGroupIPVersionParamsDateRange24w        RadarAs112TimeseriesGroupIPVersionParamsDateRange = "24w"
	RadarAs112TimeseriesGroupIPVersionParamsDateRange52w        RadarAs112TimeseriesGroupIPVersionParamsDateRange = "52w"
	RadarAs112TimeseriesGroupIPVersionParamsDateRange1dControl  RadarAs112TimeseriesGroupIPVersionParamsDateRange = "1dControl"
	RadarAs112TimeseriesGroupIPVersionParamsDateRange2dControl  RadarAs112TimeseriesGroupIPVersionParamsDateRange = "2dControl"
	RadarAs112TimeseriesGroupIPVersionParamsDateRange7dControl  RadarAs112TimeseriesGroupIPVersionParamsDateRange = "7dControl"
	RadarAs112TimeseriesGroupIPVersionParamsDateRange14dControl RadarAs112TimeseriesGroupIPVersionParamsDateRange = "14dControl"
	RadarAs112TimeseriesGroupIPVersionParamsDateRange28dControl RadarAs112TimeseriesGroupIPVersionParamsDateRange = "28dControl"
	RadarAs112TimeseriesGroupIPVersionParamsDateRange12wControl RadarAs112TimeseriesGroupIPVersionParamsDateRange = "12wControl"
	RadarAs112TimeseriesGroupIPVersionParamsDateRange24wControl RadarAs112TimeseriesGroupIPVersionParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseriesGroupIPVersionParamsFormat string

const (
	RadarAs112TimeseriesGroupIPVersionParamsFormatJson RadarAs112TimeseriesGroupIPVersionParamsFormat = "JSON"
	RadarAs112TimeseriesGroupIPVersionParamsFormatCsv  RadarAs112TimeseriesGroupIPVersionParamsFormat = "CSV"
)

type RadarAs112TimeseriesGroupIPVersionResponseEnvelope struct {
	Result  RadarAs112TimeseriesGroupIPVersionResponse             `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarAs112TimeseriesGroupIPVersionResponseEnvelopeJSON `json:"-"`
}

// radarAs112TimeseriesGroupIPVersionResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAs112TimeseriesGroupIPVersionResponseEnvelope]
type radarAs112TimeseriesGroupIPVersionResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupIPVersionResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupProtocolParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupProtocolParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseriesGroupProtocolParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseriesGroupProtocolParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupProtocolParams]'s query parameters
// as `url.Values`.
func (r RadarAs112TimeseriesGroupProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupProtocolParamsAggInterval string

const (
	RadarAs112TimeseriesGroupProtocolParamsAggInterval15m RadarAs112TimeseriesGroupProtocolParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupProtocolParamsAggInterval1h  RadarAs112TimeseriesGroupProtocolParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupProtocolParamsAggInterval1d  RadarAs112TimeseriesGroupProtocolParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupProtocolParamsAggInterval1w  RadarAs112TimeseriesGroupProtocolParamsAggInterval = "1w"
)

type RadarAs112TimeseriesGroupProtocolParamsDateRange string

const (
	RadarAs112TimeseriesGroupProtocolParamsDateRange1d         RadarAs112TimeseriesGroupProtocolParamsDateRange = "1d"
	RadarAs112TimeseriesGroupProtocolParamsDateRange2d         RadarAs112TimeseriesGroupProtocolParamsDateRange = "2d"
	RadarAs112TimeseriesGroupProtocolParamsDateRange7d         RadarAs112TimeseriesGroupProtocolParamsDateRange = "7d"
	RadarAs112TimeseriesGroupProtocolParamsDateRange14d        RadarAs112TimeseriesGroupProtocolParamsDateRange = "14d"
	RadarAs112TimeseriesGroupProtocolParamsDateRange28d        RadarAs112TimeseriesGroupProtocolParamsDateRange = "28d"
	RadarAs112TimeseriesGroupProtocolParamsDateRange12w        RadarAs112TimeseriesGroupProtocolParamsDateRange = "12w"
	RadarAs112TimeseriesGroupProtocolParamsDateRange24w        RadarAs112TimeseriesGroupProtocolParamsDateRange = "24w"
	RadarAs112TimeseriesGroupProtocolParamsDateRange52w        RadarAs112TimeseriesGroupProtocolParamsDateRange = "52w"
	RadarAs112TimeseriesGroupProtocolParamsDateRange1dControl  RadarAs112TimeseriesGroupProtocolParamsDateRange = "1dControl"
	RadarAs112TimeseriesGroupProtocolParamsDateRange2dControl  RadarAs112TimeseriesGroupProtocolParamsDateRange = "2dControl"
	RadarAs112TimeseriesGroupProtocolParamsDateRange7dControl  RadarAs112TimeseriesGroupProtocolParamsDateRange = "7dControl"
	RadarAs112TimeseriesGroupProtocolParamsDateRange14dControl RadarAs112TimeseriesGroupProtocolParamsDateRange = "14dControl"
	RadarAs112TimeseriesGroupProtocolParamsDateRange28dControl RadarAs112TimeseriesGroupProtocolParamsDateRange = "28dControl"
	RadarAs112TimeseriesGroupProtocolParamsDateRange12wControl RadarAs112TimeseriesGroupProtocolParamsDateRange = "12wControl"
	RadarAs112TimeseriesGroupProtocolParamsDateRange24wControl RadarAs112TimeseriesGroupProtocolParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseriesGroupProtocolParamsFormat string

const (
	RadarAs112TimeseriesGroupProtocolParamsFormatJson RadarAs112TimeseriesGroupProtocolParamsFormat = "JSON"
	RadarAs112TimeseriesGroupProtocolParamsFormatCsv  RadarAs112TimeseriesGroupProtocolParamsFormat = "CSV"
)

type RadarAs112TimeseriesGroupProtocolResponseEnvelope struct {
	Result  RadarAs112TimeseriesGroupProtocolResponse             `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarAs112TimeseriesGroupProtocolResponseEnvelopeJSON `json:"-"`
}

// radarAs112TimeseriesGroupProtocolResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAs112TimeseriesGroupProtocolResponseEnvelope]
type radarAs112TimeseriesGroupProtocolResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupProtocolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupQueryTypeParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupQueryTypeParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseriesGroupQueryTypeParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseriesGroupQueryTypeParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupQueryTypeParams]'s query
// parameters as `url.Values`.
func (r RadarAs112TimeseriesGroupQueryTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupQueryTypeParamsAggInterval string

const (
	RadarAs112TimeseriesGroupQueryTypeParamsAggInterval15m RadarAs112TimeseriesGroupQueryTypeParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupQueryTypeParamsAggInterval1h  RadarAs112TimeseriesGroupQueryTypeParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupQueryTypeParamsAggInterval1d  RadarAs112TimeseriesGroupQueryTypeParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupQueryTypeParamsAggInterval1w  RadarAs112TimeseriesGroupQueryTypeParamsAggInterval = "1w"
)

type RadarAs112TimeseriesGroupQueryTypeParamsDateRange string

const (
	RadarAs112TimeseriesGroupQueryTypeParamsDateRange1d         RadarAs112TimeseriesGroupQueryTypeParamsDateRange = "1d"
	RadarAs112TimeseriesGroupQueryTypeParamsDateRange2d         RadarAs112TimeseriesGroupQueryTypeParamsDateRange = "2d"
	RadarAs112TimeseriesGroupQueryTypeParamsDateRange7d         RadarAs112TimeseriesGroupQueryTypeParamsDateRange = "7d"
	RadarAs112TimeseriesGroupQueryTypeParamsDateRange14d        RadarAs112TimeseriesGroupQueryTypeParamsDateRange = "14d"
	RadarAs112TimeseriesGroupQueryTypeParamsDateRange28d        RadarAs112TimeseriesGroupQueryTypeParamsDateRange = "28d"
	RadarAs112TimeseriesGroupQueryTypeParamsDateRange12w        RadarAs112TimeseriesGroupQueryTypeParamsDateRange = "12w"
	RadarAs112TimeseriesGroupQueryTypeParamsDateRange24w        RadarAs112TimeseriesGroupQueryTypeParamsDateRange = "24w"
	RadarAs112TimeseriesGroupQueryTypeParamsDateRange52w        RadarAs112TimeseriesGroupQueryTypeParamsDateRange = "52w"
	RadarAs112TimeseriesGroupQueryTypeParamsDateRange1dControl  RadarAs112TimeseriesGroupQueryTypeParamsDateRange = "1dControl"
	RadarAs112TimeseriesGroupQueryTypeParamsDateRange2dControl  RadarAs112TimeseriesGroupQueryTypeParamsDateRange = "2dControl"
	RadarAs112TimeseriesGroupQueryTypeParamsDateRange7dControl  RadarAs112TimeseriesGroupQueryTypeParamsDateRange = "7dControl"
	RadarAs112TimeseriesGroupQueryTypeParamsDateRange14dControl RadarAs112TimeseriesGroupQueryTypeParamsDateRange = "14dControl"
	RadarAs112TimeseriesGroupQueryTypeParamsDateRange28dControl RadarAs112TimeseriesGroupQueryTypeParamsDateRange = "28dControl"
	RadarAs112TimeseriesGroupQueryTypeParamsDateRange12wControl RadarAs112TimeseriesGroupQueryTypeParamsDateRange = "12wControl"
	RadarAs112TimeseriesGroupQueryTypeParamsDateRange24wControl RadarAs112TimeseriesGroupQueryTypeParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseriesGroupQueryTypeParamsFormat string

const (
	RadarAs112TimeseriesGroupQueryTypeParamsFormatJson RadarAs112TimeseriesGroupQueryTypeParamsFormat = "JSON"
	RadarAs112TimeseriesGroupQueryTypeParamsFormatCsv  RadarAs112TimeseriesGroupQueryTypeParamsFormat = "CSV"
)

type RadarAs112TimeseriesGroupQueryTypeResponseEnvelope struct {
	Result  RadarAs112TimeseriesGroupQueryTypeResponse             `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarAs112TimeseriesGroupQueryTypeResponseEnvelopeJSON `json:"-"`
}

// radarAs112TimeseriesGroupQueryTypeResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAs112TimeseriesGroupQueryTypeResponseEnvelope]
type radarAs112TimeseriesGroupQueryTypeResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupQueryTypeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupResponseCodesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupResponseCodesParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseriesGroupResponseCodesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseriesGroupResponseCodesParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupResponseCodesParams]'s query
// parameters as `url.Values`.
func (r RadarAs112TimeseriesGroupResponseCodesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupResponseCodesParamsAggInterval string

const (
	RadarAs112TimeseriesGroupResponseCodesParamsAggInterval15m RadarAs112TimeseriesGroupResponseCodesParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupResponseCodesParamsAggInterval1h  RadarAs112TimeseriesGroupResponseCodesParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupResponseCodesParamsAggInterval1d  RadarAs112TimeseriesGroupResponseCodesParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupResponseCodesParamsAggInterval1w  RadarAs112TimeseriesGroupResponseCodesParamsAggInterval = "1w"
)

type RadarAs112TimeseriesGroupResponseCodesParamsDateRange string

const (
	RadarAs112TimeseriesGroupResponseCodesParamsDateRange1d         RadarAs112TimeseriesGroupResponseCodesParamsDateRange = "1d"
	RadarAs112TimeseriesGroupResponseCodesParamsDateRange2d         RadarAs112TimeseriesGroupResponseCodesParamsDateRange = "2d"
	RadarAs112TimeseriesGroupResponseCodesParamsDateRange7d         RadarAs112TimeseriesGroupResponseCodesParamsDateRange = "7d"
	RadarAs112TimeseriesGroupResponseCodesParamsDateRange14d        RadarAs112TimeseriesGroupResponseCodesParamsDateRange = "14d"
	RadarAs112TimeseriesGroupResponseCodesParamsDateRange28d        RadarAs112TimeseriesGroupResponseCodesParamsDateRange = "28d"
	RadarAs112TimeseriesGroupResponseCodesParamsDateRange12w        RadarAs112TimeseriesGroupResponseCodesParamsDateRange = "12w"
	RadarAs112TimeseriesGroupResponseCodesParamsDateRange24w        RadarAs112TimeseriesGroupResponseCodesParamsDateRange = "24w"
	RadarAs112TimeseriesGroupResponseCodesParamsDateRange52w        RadarAs112TimeseriesGroupResponseCodesParamsDateRange = "52w"
	RadarAs112TimeseriesGroupResponseCodesParamsDateRange1dControl  RadarAs112TimeseriesGroupResponseCodesParamsDateRange = "1dControl"
	RadarAs112TimeseriesGroupResponseCodesParamsDateRange2dControl  RadarAs112TimeseriesGroupResponseCodesParamsDateRange = "2dControl"
	RadarAs112TimeseriesGroupResponseCodesParamsDateRange7dControl  RadarAs112TimeseriesGroupResponseCodesParamsDateRange = "7dControl"
	RadarAs112TimeseriesGroupResponseCodesParamsDateRange14dControl RadarAs112TimeseriesGroupResponseCodesParamsDateRange = "14dControl"
	RadarAs112TimeseriesGroupResponseCodesParamsDateRange28dControl RadarAs112TimeseriesGroupResponseCodesParamsDateRange = "28dControl"
	RadarAs112TimeseriesGroupResponseCodesParamsDateRange12wControl RadarAs112TimeseriesGroupResponseCodesParamsDateRange = "12wControl"
	RadarAs112TimeseriesGroupResponseCodesParamsDateRange24wControl RadarAs112TimeseriesGroupResponseCodesParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseriesGroupResponseCodesParamsFormat string

const (
	RadarAs112TimeseriesGroupResponseCodesParamsFormatJson RadarAs112TimeseriesGroupResponseCodesParamsFormat = "JSON"
	RadarAs112TimeseriesGroupResponseCodesParamsFormatCsv  RadarAs112TimeseriesGroupResponseCodesParamsFormat = "CSV"
)

type RadarAs112TimeseriesGroupResponseCodesResponseEnvelope struct {
	Result  RadarAs112TimeseriesGroupResponseCodesResponse             `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    radarAs112TimeseriesGroupResponseCodesResponseEnvelopeJSON `json:"-"`
}

// radarAs112TimeseriesGroupResponseCodesResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAs112TimeseriesGroupResponseCodesResponseEnvelope]
type radarAs112TimeseriesGroupResponseCodesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupResponseCodesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
