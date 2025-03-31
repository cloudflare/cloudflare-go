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

// NetflowService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetflowService] method instead.
type NetflowService struct {
	Options []option.RequestOption
	Top     *NetflowTopService
}

// NewNetflowService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNetflowService(opts ...option.RequestOption) (r *NetflowService) {
	r = &NetflowService{}
	r.Options = opts
	r.Top = NewNetflowTopService(opts...)
	return
}

// Retrieves the distribution of network traffic (NetFlows) by HTTP vs other
// protocols.
func (r *NetflowService) Summary(ctx context.Context, query NetflowSummaryParams, opts ...option.RequestOption) (res *NetflowSummaryResponse, err error) {
	var env NetflowSummaryResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/netflows/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves network traffic (NetFlows) over time.
func (r *NetflowService) Timeseries(ctx context.Context, query NetflowTimeseriesParams, opts ...option.RequestOption) (res *NetflowTimeseriesResponse, err error) {
	var env NetflowTimeseriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/netflows/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NetflowSummaryResponse struct {
	Meta     NetflowSummaryResponseMeta     `json:"meta,required"`
	Summary0 NetflowSummaryResponseSummary0 `json:"summary_0,required"`
	JSON     netflowSummaryResponseJSON     `json:"-"`
}

// netflowSummaryResponseJSON contains the JSON metadata for the struct
// [NetflowSummaryResponse]
type netflowSummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type NetflowSummaryResponseMeta struct {
	DateRange      []interface{}                            `json:"dateRange,required"`
	ConfidenceInfo NetflowSummaryResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           netflowSummaryResponseMetaJSON           `json:"-"`
}

// netflowSummaryResponseMetaJSON contains the JSON metadata for the struct
// [NetflowSummaryResponseMeta]
type netflowSummaryResponseMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NetflowSummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowSummaryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type NetflowSummaryResponseMetaConfidenceInfo struct {
	Annotations []interface{}                                `json:"annotations"`
	Level       int64                                        `json:"level"`
	JSON        netflowSummaryResponseMetaConfidenceInfoJSON `json:"-"`
}

// netflowSummaryResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [NetflowSummaryResponseMetaConfidenceInfo]
type netflowSummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowSummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowSummaryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type NetflowSummaryResponseSummary0 struct {
	HTTP  string                             `json:"HTTP,required"`
	Other string                             `json:"OTHER,required"`
	JSON  netflowSummaryResponseSummary0JSON `json:"-"`
}

// netflowSummaryResponseSummary0JSON contains the JSON metadata for the struct
// [NetflowSummaryResponseSummary0]
type netflowSummaryResponseSummary0JSON struct {
	HTTP        apijson.Field
	Other       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowSummaryResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowSummaryResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type NetflowTimeseriesResponse struct {
	Meta   NetflowTimeseriesResponseMeta `json:"meta,required"`
	Serie0 interface{}                   `json:"serie_0,required"`
	JSON   netflowTimeseriesResponseJSON `json:"-"`
}

// netflowTimeseriesResponseJSON contains the JSON metadata for the struct
// [NetflowTimeseriesResponse]
type netflowTimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type NetflowTimeseriesResponseMeta struct {
	AggInterval    string                                      `json:"aggInterval,required"`
	DateRange      []interface{}                               `json:"dateRange,required"`
	LastUpdated    time.Time                                   `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo NetflowTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           netflowTimeseriesResponseMetaJSON           `json:"-"`
}

// netflowTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [NetflowTimeseriesResponseMeta]
type netflowTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NetflowTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type NetflowTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []interface{}                                   `json:"annotations"`
	Level       int64                                           `json:"level"`
	JSON        netflowTimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// netflowTimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [NetflowTimeseriesResponseMetaConfidenceInfo]
type netflowTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type NetflowSummaryParams struct {
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
	Format param.Field[NetflowSummaryParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [NetflowSummaryParams]'s query parameters as `url.Values`.
func (r NetflowSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type NetflowSummaryParamsFormat string

const (
	NetflowSummaryParamsFormatJson NetflowSummaryParamsFormat = "JSON"
	NetflowSummaryParamsFormatCsv  NetflowSummaryParamsFormat = "CSV"
)

func (r NetflowSummaryParamsFormat) IsKnown() bool {
	switch r {
	case NetflowSummaryParamsFormatJson, NetflowSummaryParamsFormatCsv:
		return true
	}
	return false
}

type NetflowSummaryResponseEnvelope struct {
	Result  NetflowSummaryResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    netflowSummaryResponseEnvelopeJSON `json:"-"`
}

// netflowSummaryResponseEnvelopeJSON contains the JSON metadata for the struct
// [NetflowSummaryResponseEnvelope]
type netflowSummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowSummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NetflowTimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[NetflowTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[NetflowTimeseriesParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[NetflowTimeseriesParamsNormalization] `query:"normalization"`
	// Array of network traffic product types to filters results by.
	Product param.Field[[]NetflowTimeseriesParamsProduct] `query:"product"`
}

// URLQuery serializes [NetflowTimeseriesParams]'s query parameters as
// `url.Values`.
func (r NetflowTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type NetflowTimeseriesParamsAggInterval string

const (
	NetflowTimeseriesParamsAggInterval15m NetflowTimeseriesParamsAggInterval = "15m"
	NetflowTimeseriesParamsAggInterval1h  NetflowTimeseriesParamsAggInterval = "1h"
	NetflowTimeseriesParamsAggInterval1d  NetflowTimeseriesParamsAggInterval = "1d"
	NetflowTimeseriesParamsAggInterval1w  NetflowTimeseriesParamsAggInterval = "1w"
)

func (r NetflowTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case NetflowTimeseriesParamsAggInterval15m, NetflowTimeseriesParamsAggInterval1h, NetflowTimeseriesParamsAggInterval1d, NetflowTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type NetflowTimeseriesParamsFormat string

const (
	NetflowTimeseriesParamsFormatJson NetflowTimeseriesParamsFormat = "JSON"
	NetflowTimeseriesParamsFormatCsv  NetflowTimeseriesParamsFormat = "CSV"
)

func (r NetflowTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case NetflowTimeseriesParamsFormatJson, NetflowTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type NetflowTimeseriesParamsNormalization string

const (
	NetflowTimeseriesParamsNormalizationPercentageChange NetflowTimeseriesParamsNormalization = "PERCENTAGE_CHANGE"
	NetflowTimeseriesParamsNormalizationMin0Max          NetflowTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r NetflowTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case NetflowTimeseriesParamsNormalizationPercentageChange, NetflowTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type NetflowTimeseriesParamsProduct string

const (
	NetflowTimeseriesParamsProductHTTP NetflowTimeseriesParamsProduct = "HTTP"
	NetflowTimeseriesParamsProductAll  NetflowTimeseriesParamsProduct = "ALL"
)

func (r NetflowTimeseriesParamsProduct) IsKnown() bool {
	switch r {
	case NetflowTimeseriesParamsProductHTTP, NetflowTimeseriesParamsProductAll:
		return true
	}
	return false
}

type NetflowTimeseriesResponseEnvelope struct {
	Result  NetflowTimeseriesResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    netflowTimeseriesResponseEnvelopeJSON `json:"-"`
}

// netflowTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [NetflowTimeseriesResponseEnvelope]
type netflowTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
