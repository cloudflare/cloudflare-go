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

// Get network traffic change over time. Visit
// https://en.wikipedia.org/wiki/NetFlow for more information on NetFlows.
func (r *NetflowService) Timeseries(ctx context.Context, query NetflowTimeseriesParams, opts ...option.RequestOption) (res *NetflowTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env NetflowTimeseriesResponseEnvelope
	path := "radar/netflows/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NetflowTimeseriesResponse struct {
	Meta   NetflowTimeseriesResponseMeta   `json:"meta,required"`
	Serie0 NetflowTimeseriesResponseSerie0 `json:"serie_0,required"`
	JSON   netflowTimeseriesResponseJSON   `json:"-"`
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
	DateRange      []NetflowTimeseriesResponseMetaDateRange    `json:"dateRange,required"`
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

type NetflowTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                  `json:"startTime,required" format:"date-time"`
	JSON      netflowTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// netflowTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [NetflowTimeseriesResponseMetaDateRange]
type netflowTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type NetflowTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []NetflowTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                   `json:"level"`
	JSON        netflowTimeseriesResponseMetaConfidenceInfoJSON         `json:"-"`
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

type NetflowTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                    `json:"dataSource,required"`
	Description     string                                                    `json:"description,required"`
	EventType       string                                                    `json:"eventType,required"`
	IsInstantaneous bool                                                      `json:"isInstantaneous,required"`
	EndTime         time.Time                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                 `json:"startTime" format:"date-time"`
	JSON            netflowTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// netflowTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [NetflowTimeseriesResponseMetaConfidenceInfoAnnotation]
type netflowTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	EndTime         apijson.Field
	LinkedURL       apijson.Field
	StartTime       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *NetflowTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type NetflowTimeseriesResponseSerie0 struct {
	Timestamps []time.Time                         `json:"timestamps,required" format:"date-time"`
	Values     []string                            `json:"values,required"`
	JSON       netflowTimeseriesResponseSerie0JSON `json:"-"`
}

// netflowTimeseriesResponseSerie0JSON contains the JSON metadata for the struct
// [NetflowTimeseriesResponseSerie0]
type netflowTimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type NetflowTimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[NetflowTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]NetflowTimeseriesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[NetflowTimeseriesParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[NetflowTimeseriesParamsNormalization] `query:"normalization"`
	// Array of network traffic product types.
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

type NetflowTimeseriesParamsDateRange string

const (
	NetflowTimeseriesParamsDateRange1d         NetflowTimeseriesParamsDateRange = "1d"
	NetflowTimeseriesParamsDateRange2d         NetflowTimeseriesParamsDateRange = "2d"
	NetflowTimeseriesParamsDateRange7d         NetflowTimeseriesParamsDateRange = "7d"
	NetflowTimeseriesParamsDateRange14d        NetflowTimeseriesParamsDateRange = "14d"
	NetflowTimeseriesParamsDateRange28d        NetflowTimeseriesParamsDateRange = "28d"
	NetflowTimeseriesParamsDateRange12w        NetflowTimeseriesParamsDateRange = "12w"
	NetflowTimeseriesParamsDateRange24w        NetflowTimeseriesParamsDateRange = "24w"
	NetflowTimeseriesParamsDateRange52w        NetflowTimeseriesParamsDateRange = "52w"
	NetflowTimeseriesParamsDateRange1dControl  NetflowTimeseriesParamsDateRange = "1dControl"
	NetflowTimeseriesParamsDateRange2dControl  NetflowTimeseriesParamsDateRange = "2dControl"
	NetflowTimeseriesParamsDateRange7dControl  NetflowTimeseriesParamsDateRange = "7dControl"
	NetflowTimeseriesParamsDateRange14dControl NetflowTimeseriesParamsDateRange = "14dControl"
	NetflowTimeseriesParamsDateRange28dControl NetflowTimeseriesParamsDateRange = "28dControl"
	NetflowTimeseriesParamsDateRange12wControl NetflowTimeseriesParamsDateRange = "12wControl"
	NetflowTimeseriesParamsDateRange24wControl NetflowTimeseriesParamsDateRange = "24wControl"
)

func (r NetflowTimeseriesParamsDateRange) IsKnown() bool {
	switch r {
	case NetflowTimeseriesParamsDateRange1d, NetflowTimeseriesParamsDateRange2d, NetflowTimeseriesParamsDateRange7d, NetflowTimeseriesParamsDateRange14d, NetflowTimeseriesParamsDateRange28d, NetflowTimeseriesParamsDateRange12w, NetflowTimeseriesParamsDateRange24w, NetflowTimeseriesParamsDateRange52w, NetflowTimeseriesParamsDateRange1dControl, NetflowTimeseriesParamsDateRange2dControl, NetflowTimeseriesParamsDateRange7dControl, NetflowTimeseriesParamsDateRange14dControl, NetflowTimeseriesParamsDateRange28dControl, NetflowTimeseriesParamsDateRange12wControl, NetflowTimeseriesParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
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
