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

// ConnectionTamperingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewConnectionTamperingService]
// method instead.
type ConnectionTamperingService struct {
	Options []option.RequestOption
}

// NewConnectionTamperingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectionTamperingService(opts ...option.RequestOption) (r *ConnectionTamperingService) {
	r = &ConnectionTamperingService{}
	r.Options = opts
	return
}

// Distribution of connection tampering types over a given time period.
func (r *ConnectionTamperingService) Summary(ctx context.Context, query ConnectionTamperingSummaryParams, opts ...option.RequestOption) (res *ConnectionTamperingSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ConnectionTamperingSummaryResponseEnvelope
	path := "radar/connection_tampering/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Distribution of connection tampering types over time.
func (r *ConnectionTamperingService) TimeseriesGroups(ctx context.Context, query ConnectionTamperingTimeseriesGroupsParams, opts ...option.RequestOption) (res *ConnectionTamperingTimeseriesGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ConnectionTamperingTimeseriesGroupsResponseEnvelope
	path := "radar/connection_tampering/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ConnectionTamperingSummaryResponse struct {
	Meta     ConnectionTamperingSummaryResponseMeta     `json:"meta,required"`
	Summary0 ConnectionTamperingSummaryResponseSummary0 `json:"summary_0,required"`
	JSON     connectionTamperingSummaryResponseJSON     `json:"-"`
}

// connectionTamperingSummaryResponseJSON contains the JSON metadata for the struct
// [ConnectionTamperingSummaryResponse]
type connectionTamperingSummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionTamperingSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionTamperingSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectionTamperingSummaryResponseMeta struct {
	DateRange      []UnnamedSchemaRef175                                `json:"dateRange,required"`
	ConfidenceInfo ConnectionTamperingSummaryResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           connectionTamperingSummaryResponseMetaJSON           `json:"-"`
}

// connectionTamperingSummaryResponseMetaJSON contains the JSON metadata for the
// struct [ConnectionTamperingSummaryResponseMeta]
type connectionTamperingSummaryResponseMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConnectionTamperingSummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionTamperingSummaryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type ConnectionTamperingSummaryResponseMetaConfidenceInfo struct {
	Annotations []UnnamedSchemaRef174                                    `json:"annotations"`
	Level       int64                                                    `json:"level"`
	JSON        connectionTamperingSummaryResponseMetaConfidenceInfoJSON `json:"-"`
}

// connectionTamperingSummaryResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [ConnectionTamperingSummaryResponseMetaConfidenceInfo]
type connectionTamperingSummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionTamperingSummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionTamperingSummaryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type ConnectionTamperingSummaryResponseSummary0 struct {
	// Connections matching signatures for tampering later in the connection, after the
	// transfer of multiple data packets.
	LaterInFlow string `json:"later_in_flow,required"`
	// Connections that do not match any tampering signatures.
	NoMatch string `json:"no_match,required"`
	// Connections matching signatures for tampering after the receipt of a SYN packet
	// and ACK packet, meaning the TCP connection was successfully established but the
	// server did not receive any subsequent packets.
	PostAck string `json:"post_ack,required"`
	// Connections matching signatures for tampering after the receipt of a packet with
	// PSH flag set, following connection establishment.
	PostPsh string `json:"post_psh,required"`
	// Connections matching signatures for tampering after the receipt of only a single
	// SYN packet, and before the handshake completes.
	PostSyn string                                         `json:"post_syn,required"`
	JSON    connectionTamperingSummaryResponseSummary0JSON `json:"-"`
}

// connectionTamperingSummaryResponseSummary0JSON contains the JSON metadata for
// the struct [ConnectionTamperingSummaryResponseSummary0]
type connectionTamperingSummaryResponseSummary0JSON struct {
	LaterInFlow apijson.Field
	NoMatch     apijson.Field
	PostAck     apijson.Field
	PostPsh     apijson.Field
	PostSyn     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionTamperingSummaryResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionTamperingSummaryResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type ConnectionTamperingTimeseriesGroupsResponse struct {
	Meta   ConnectionTamperingTimeseriesGroupsResponseMeta   `json:"meta,required"`
	Serie0 ConnectionTamperingTimeseriesGroupsResponseSerie0 `json:"serie_0,required"`
	JSON   connectionTamperingTimeseriesGroupsResponseJSON   `json:"-"`
}

// connectionTamperingTimeseriesGroupsResponseJSON contains the JSON metadata for
// the struct [ConnectionTamperingTimeseriesGroupsResponse]
type connectionTamperingTimeseriesGroupsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionTamperingTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionTamperingTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectionTamperingTimeseriesGroupsResponseMeta struct {
	AggInterval    string                                                        `json:"aggInterval,required"`
	DateRange      []UnnamedSchemaRef175                                         `json:"dateRange,required"`
	LastUpdated    time.Time                                                     `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo ConnectionTamperingTimeseriesGroupsResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           connectionTamperingTimeseriesGroupsResponseMetaJSON           `json:"-"`
}

// connectionTamperingTimeseriesGroupsResponseMetaJSON contains the JSON metadata
// for the struct [ConnectionTamperingTimeseriesGroupsResponseMeta]
type connectionTamperingTimeseriesGroupsResponseMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConnectionTamperingTimeseriesGroupsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionTamperingTimeseriesGroupsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type ConnectionTamperingTimeseriesGroupsResponseMetaConfidenceInfo struct {
	Annotations []UnnamedSchemaRef174                                             `json:"annotations"`
	Level       int64                                                             `json:"level"`
	JSON        connectionTamperingTimeseriesGroupsResponseMetaConfidenceInfoJSON `json:"-"`
}

// connectionTamperingTimeseriesGroupsResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [ConnectionTamperingTimeseriesGroupsResponseMetaConfidenceInfo]
type connectionTamperingTimeseriesGroupsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionTamperingTimeseriesGroupsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionTamperingTimeseriesGroupsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type ConnectionTamperingTimeseriesGroupsResponseSerie0 struct {
	// Connections matching signatures for tampering later in the connection, after the
	// transfer of multiple data packets.
	LaterInFlow []string `json:"later_in_flow,required"`
	// Connections that do not match any tampering signatures.
	NoMatch []string `json:"no_match,required"`
	// Connections matching signatures for tampering after the receipt of a SYN packet
	// and ACK packet, meaning the TCP connection was successfully established but the
	// server did not receive any subsequent packets.
	PostAck []string `json:"post_ack,required"`
	// Connections matching signatures for tampering after the receipt of a packet with
	// PSH flag set, following connection establishment.
	PostPsh []string `json:"post_psh,required"`
	// Connections matching signatures for tampering after the receipt of only a single
	// SYN packet, and before the handshake completes.
	PostSyn    []string                                              `json:"post_syn,required"`
	Timestamps []time.Time                                           `json:"timestamps,required" format:"date-time"`
	JSON       connectionTamperingTimeseriesGroupsResponseSerie0JSON `json:"-"`
}

// connectionTamperingTimeseriesGroupsResponseSerie0JSON contains the JSON metadata
// for the struct [ConnectionTamperingTimeseriesGroupsResponseSerie0]
type connectionTamperingTimeseriesGroupsResponseSerie0JSON struct {
	LaterInFlow apijson.Field
	NoMatch     apijson.Field
	PostAck     apijson.Field
	PostPsh     apijson.Field
	PostSyn     apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionTamperingTimeseriesGroupsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionTamperingTimeseriesGroupsResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type ConnectionTamperingSummaryParams struct {
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
	DateRange param.Field[[]ConnectionTamperingSummaryParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[ConnectionTamperingSummaryParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [ConnectionTamperingSummaryParams]'s query parameters as
// `url.Values`.
func (r ConnectionTamperingSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConnectionTamperingSummaryParamsDateRange string

const (
	ConnectionTamperingSummaryParamsDateRange1d         ConnectionTamperingSummaryParamsDateRange = "1d"
	ConnectionTamperingSummaryParamsDateRange2d         ConnectionTamperingSummaryParamsDateRange = "2d"
	ConnectionTamperingSummaryParamsDateRange7d         ConnectionTamperingSummaryParamsDateRange = "7d"
	ConnectionTamperingSummaryParamsDateRange14d        ConnectionTamperingSummaryParamsDateRange = "14d"
	ConnectionTamperingSummaryParamsDateRange28d        ConnectionTamperingSummaryParamsDateRange = "28d"
	ConnectionTamperingSummaryParamsDateRange12w        ConnectionTamperingSummaryParamsDateRange = "12w"
	ConnectionTamperingSummaryParamsDateRange24w        ConnectionTamperingSummaryParamsDateRange = "24w"
	ConnectionTamperingSummaryParamsDateRange52w        ConnectionTamperingSummaryParamsDateRange = "52w"
	ConnectionTamperingSummaryParamsDateRange1dControl  ConnectionTamperingSummaryParamsDateRange = "1dControl"
	ConnectionTamperingSummaryParamsDateRange2dControl  ConnectionTamperingSummaryParamsDateRange = "2dControl"
	ConnectionTamperingSummaryParamsDateRange7dControl  ConnectionTamperingSummaryParamsDateRange = "7dControl"
	ConnectionTamperingSummaryParamsDateRange14dControl ConnectionTamperingSummaryParamsDateRange = "14dControl"
	ConnectionTamperingSummaryParamsDateRange28dControl ConnectionTamperingSummaryParamsDateRange = "28dControl"
	ConnectionTamperingSummaryParamsDateRange12wControl ConnectionTamperingSummaryParamsDateRange = "12wControl"
	ConnectionTamperingSummaryParamsDateRange24wControl ConnectionTamperingSummaryParamsDateRange = "24wControl"
)

func (r ConnectionTamperingSummaryParamsDateRange) IsKnown() bool {
	switch r {
	case ConnectionTamperingSummaryParamsDateRange1d, ConnectionTamperingSummaryParamsDateRange2d, ConnectionTamperingSummaryParamsDateRange7d, ConnectionTamperingSummaryParamsDateRange14d, ConnectionTamperingSummaryParamsDateRange28d, ConnectionTamperingSummaryParamsDateRange12w, ConnectionTamperingSummaryParamsDateRange24w, ConnectionTamperingSummaryParamsDateRange52w, ConnectionTamperingSummaryParamsDateRange1dControl, ConnectionTamperingSummaryParamsDateRange2dControl, ConnectionTamperingSummaryParamsDateRange7dControl, ConnectionTamperingSummaryParamsDateRange14dControl, ConnectionTamperingSummaryParamsDateRange28dControl, ConnectionTamperingSummaryParamsDateRange12wControl, ConnectionTamperingSummaryParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type ConnectionTamperingSummaryParamsFormat string

const (
	ConnectionTamperingSummaryParamsFormatJson ConnectionTamperingSummaryParamsFormat = "JSON"
	ConnectionTamperingSummaryParamsFormatCsv  ConnectionTamperingSummaryParamsFormat = "CSV"
)

func (r ConnectionTamperingSummaryParamsFormat) IsKnown() bool {
	switch r {
	case ConnectionTamperingSummaryParamsFormatJson, ConnectionTamperingSummaryParamsFormatCsv:
		return true
	}
	return false
}

type ConnectionTamperingSummaryResponseEnvelope struct {
	Result  ConnectionTamperingSummaryResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    connectionTamperingSummaryResponseEnvelopeJSON `json:"-"`
}

// connectionTamperingSummaryResponseEnvelopeJSON contains the JSON metadata for
// the struct [ConnectionTamperingSummaryResponseEnvelope]
type connectionTamperingSummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionTamperingSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionTamperingSummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectionTamperingTimeseriesGroupsParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[ConnectionTamperingTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]ConnectionTamperingTimeseriesGroupsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[ConnectionTamperingTimeseriesGroupsParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [ConnectionTamperingTimeseriesGroupsParams]'s query
// parameters as `url.Values`.
func (r ConnectionTamperingTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type ConnectionTamperingTimeseriesGroupsParamsAggInterval string

const (
	ConnectionTamperingTimeseriesGroupsParamsAggInterval15m ConnectionTamperingTimeseriesGroupsParamsAggInterval = "15m"
	ConnectionTamperingTimeseriesGroupsParamsAggInterval1h  ConnectionTamperingTimeseriesGroupsParamsAggInterval = "1h"
	ConnectionTamperingTimeseriesGroupsParamsAggInterval1d  ConnectionTamperingTimeseriesGroupsParamsAggInterval = "1d"
	ConnectionTamperingTimeseriesGroupsParamsAggInterval1w  ConnectionTamperingTimeseriesGroupsParamsAggInterval = "1w"
)

func (r ConnectionTamperingTimeseriesGroupsParamsAggInterval) IsKnown() bool {
	switch r {
	case ConnectionTamperingTimeseriesGroupsParamsAggInterval15m, ConnectionTamperingTimeseriesGroupsParamsAggInterval1h, ConnectionTamperingTimeseriesGroupsParamsAggInterval1d, ConnectionTamperingTimeseriesGroupsParamsAggInterval1w:
		return true
	}
	return false
}

type ConnectionTamperingTimeseriesGroupsParamsDateRange string

const (
	ConnectionTamperingTimeseriesGroupsParamsDateRange1d         ConnectionTamperingTimeseriesGroupsParamsDateRange = "1d"
	ConnectionTamperingTimeseriesGroupsParamsDateRange2d         ConnectionTamperingTimeseriesGroupsParamsDateRange = "2d"
	ConnectionTamperingTimeseriesGroupsParamsDateRange7d         ConnectionTamperingTimeseriesGroupsParamsDateRange = "7d"
	ConnectionTamperingTimeseriesGroupsParamsDateRange14d        ConnectionTamperingTimeseriesGroupsParamsDateRange = "14d"
	ConnectionTamperingTimeseriesGroupsParamsDateRange28d        ConnectionTamperingTimeseriesGroupsParamsDateRange = "28d"
	ConnectionTamperingTimeseriesGroupsParamsDateRange12w        ConnectionTamperingTimeseriesGroupsParamsDateRange = "12w"
	ConnectionTamperingTimeseriesGroupsParamsDateRange24w        ConnectionTamperingTimeseriesGroupsParamsDateRange = "24w"
	ConnectionTamperingTimeseriesGroupsParamsDateRange52w        ConnectionTamperingTimeseriesGroupsParamsDateRange = "52w"
	ConnectionTamperingTimeseriesGroupsParamsDateRange1dControl  ConnectionTamperingTimeseriesGroupsParamsDateRange = "1dControl"
	ConnectionTamperingTimeseriesGroupsParamsDateRange2dControl  ConnectionTamperingTimeseriesGroupsParamsDateRange = "2dControl"
	ConnectionTamperingTimeseriesGroupsParamsDateRange7dControl  ConnectionTamperingTimeseriesGroupsParamsDateRange = "7dControl"
	ConnectionTamperingTimeseriesGroupsParamsDateRange14dControl ConnectionTamperingTimeseriesGroupsParamsDateRange = "14dControl"
	ConnectionTamperingTimeseriesGroupsParamsDateRange28dControl ConnectionTamperingTimeseriesGroupsParamsDateRange = "28dControl"
	ConnectionTamperingTimeseriesGroupsParamsDateRange12wControl ConnectionTamperingTimeseriesGroupsParamsDateRange = "12wControl"
	ConnectionTamperingTimeseriesGroupsParamsDateRange24wControl ConnectionTamperingTimeseriesGroupsParamsDateRange = "24wControl"
)

func (r ConnectionTamperingTimeseriesGroupsParamsDateRange) IsKnown() bool {
	switch r {
	case ConnectionTamperingTimeseriesGroupsParamsDateRange1d, ConnectionTamperingTimeseriesGroupsParamsDateRange2d, ConnectionTamperingTimeseriesGroupsParamsDateRange7d, ConnectionTamperingTimeseriesGroupsParamsDateRange14d, ConnectionTamperingTimeseriesGroupsParamsDateRange28d, ConnectionTamperingTimeseriesGroupsParamsDateRange12w, ConnectionTamperingTimeseriesGroupsParamsDateRange24w, ConnectionTamperingTimeseriesGroupsParamsDateRange52w, ConnectionTamperingTimeseriesGroupsParamsDateRange1dControl, ConnectionTamperingTimeseriesGroupsParamsDateRange2dControl, ConnectionTamperingTimeseriesGroupsParamsDateRange7dControl, ConnectionTamperingTimeseriesGroupsParamsDateRange14dControl, ConnectionTamperingTimeseriesGroupsParamsDateRange28dControl, ConnectionTamperingTimeseriesGroupsParamsDateRange12wControl, ConnectionTamperingTimeseriesGroupsParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type ConnectionTamperingTimeseriesGroupsParamsFormat string

const (
	ConnectionTamperingTimeseriesGroupsParamsFormatJson ConnectionTamperingTimeseriesGroupsParamsFormat = "JSON"
	ConnectionTamperingTimeseriesGroupsParamsFormatCsv  ConnectionTamperingTimeseriesGroupsParamsFormat = "CSV"
)

func (r ConnectionTamperingTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case ConnectionTamperingTimeseriesGroupsParamsFormatJson, ConnectionTamperingTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}

type ConnectionTamperingTimeseriesGroupsResponseEnvelope struct {
	Result  ConnectionTamperingTimeseriesGroupsResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    connectionTamperingTimeseriesGroupsResponseEnvelopeJSON `json:"-"`
}

// connectionTamperingTimeseriesGroupsResponseEnvelopeJSON contains the JSON
// metadata for the struct [ConnectionTamperingTimeseriesGroupsResponseEnvelope]
type connectionTamperingTimeseriesGroupsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionTamperingTimeseriesGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionTamperingTimeseriesGroupsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
