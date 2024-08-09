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

// DNSService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSService] method instead.
type DNSService struct {
	Options          []option.RequestOption
	Top              *DNSTopService
	Summary          *DNSSummaryService
	TimeseriesGroups *DNSTimeseriesGroupService
}

// NewDNSService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDNSService(opts ...option.RequestOption) (r *DNSService) {
	r = &DNSService{}
	r.Options = opts
	r.Top = NewDNSTopService(opts...)
	r.Summary = NewDNSSummaryService(opts...)
	r.TimeseriesGroups = NewDNSTimeseriesGroupService(opts...)
	return
}

// Get DNS queries change over time.
func (r *DNSService) Timeseries(ctx context.Context, query DNSTimeseriesParams, opts ...option.RequestOption) (res *DNSTimeseriesResponse, err error) {
	var env DNSTimeseriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSTimeseriesResponse struct {
	Meta   DNSTimeseriesResponseMeta   `json:"meta,required"`
	Serie0 DNSTimeseriesResponseSerie0 `json:"serie_0,required"`
	JSON   dnsTimeseriesResponseJSON   `json:"-"`
}

// dnsTimeseriesResponseJSON contains the JSON metadata for the struct
// [DNSTimeseriesResponse]
type dnsTimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesResponseMeta struct {
	AggInterval    string                                  `json:"aggInterval,required"`
	DateRange      []DNSTimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                               `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo DNSTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           dnsTimeseriesResponseMetaJSON           `json:"-"`
}

// dnsTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [DNSTimeseriesResponseMeta]
type dnsTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                              `json:"startTime,required" format:"date-time"`
	JSON      dnsTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// dnsTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [DNSTimeseriesResponseMetaDateRange]
type dnsTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []DNSTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                               `json:"level"`
	JSON        dnsTimeseriesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// dnsTimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [DNSTimeseriesResponseMetaConfidenceInfo]
type dnsTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                `json:"dataSource,required"`
	Description     string                                                `json:"description,required"`
	EventType       string                                                `json:"eventType,required"`
	IsInstantaneous bool                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                `json:"linkedUrl"`
	StartTime       time.Time                                             `json:"startTime" format:"date-time"`
	JSON            dnsTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata
// for the struct [DNSTimeseriesResponseMetaConfidenceInfoAnnotation]
type dnsTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *DNSTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesResponseSerie0 struct {
	Timestamps []time.Time                     `json:"timestamps,required" format:"date-time"`
	Values     []string                        `json:"values,required"`
	JSON       dnsTimeseriesResponseSerie0JSON `json:"-"`
}

// dnsTimeseriesResponseSerie0JSON contains the JSON metadata for the struct
// [DNSTimeseriesResponseSerie0]
type dnsTimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type DNSTimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[DNSTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[DNSTimeseriesParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for ccTLD.
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [DNSTimeseriesParams]'s query parameters as `url.Values`.
func (r DNSTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type DNSTimeseriesParamsAggInterval string

const (
	DNSTimeseriesParamsAggInterval15m DNSTimeseriesParamsAggInterval = "15m"
	DNSTimeseriesParamsAggInterval1h  DNSTimeseriesParamsAggInterval = "1h"
	DNSTimeseriesParamsAggInterval1d  DNSTimeseriesParamsAggInterval = "1d"
	DNSTimeseriesParamsAggInterval1w  DNSTimeseriesParamsAggInterval = "1w"
)

func (r DNSTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsAggInterval15m, DNSTimeseriesParamsAggInterval1h, DNSTimeseriesParamsAggInterval1d, DNSTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format results are returned in.
type DNSTimeseriesParamsFormat string

const (
	DNSTimeseriesParamsFormatJson DNSTimeseriesParamsFormat = "JSON"
	DNSTimeseriesParamsFormatCsv  DNSTimeseriesParamsFormat = "CSV"
)

func (r DNSTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case DNSTimeseriesParamsFormatJson, DNSTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type DNSTimeseriesResponseEnvelope struct {
	Result  DNSTimeseriesResponse             `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    dnsTimeseriesResponseEnvelopeJSON `json:"-"`
}

// dnsTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSTimeseriesResponseEnvelope]
type dnsTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
