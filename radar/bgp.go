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

// BGPService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBGPService] method instead.
type BGPService struct {
	Options []option.RequestOption
	Leaks   *BGPLeakService
	Top     *BGPTopService
	Hijacks *BGPHijackService
	Routes  *BGPRouteService
	IPs     *BGPIPService
}

// NewBGPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBGPService(opts ...option.RequestOption) (r *BGPService) {
	r = &BGPService{}
	r.Options = opts
	r.Leaks = NewBGPLeakService(opts...)
	r.Top = NewBGPTopService(opts...)
	r.Hijacks = NewBGPHijackService(opts...)
	r.Routes = NewBGPRouteService(opts...)
	r.IPs = NewBGPIPService(opts...)
	return
}

// Get BGP updates change over time. Raw values are returned. When requesting
// updates for an autonomous system (AS), only BGP updates of type announcement are
// returned.
func (r *BGPService) Timeseries(ctx context.Context, query BGPTimeseriesParams, opts ...option.RequestOption) (res *BGPTimeseriesResponse, err error) {
	var env BGPTimeseriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BGPTimeseriesResponse struct {
	Meta   BGPTimeseriesResponseMeta   `json:"meta,required"`
	Serie0 BGPTimeseriesResponseSerie0 `json:"serie_0,required"`
	JSON   bgpTimeseriesResponseJSON   `json:"-"`
}

// bgpTimeseriesResponseJSON contains the JSON metadata for the struct
// [BGPTimeseriesResponse]
type bgpTimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type BGPTimeseriesResponseMeta struct {
	AggInterval    string                                  `json:"aggInterval,required"`
	DateRange      []BGPTimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                               `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo BGPTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           bgpTimeseriesResponseMetaJSON           `json:"-"`
}

// bgpTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [BGPTimeseriesResponseMeta]
type bgpTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BGPTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type BGPTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                              `json:"startTime,required" format:"date-time"`
	JSON      bgpTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// bgpTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [BGPTimeseriesResponseMetaDateRange]
type bgpTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type BGPTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []BGPTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                               `json:"level"`
	JSON        bgpTimeseriesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// bgpTimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [BGPTimeseriesResponseMetaConfidenceInfo]
type bgpTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type BGPTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                `json:"dataSource,required"`
	Description     string                                                `json:"description,required"`
	EventType       string                                                `json:"eventType,required"`
	IsInstantaneous bool                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                `json:"linkedUrl"`
	StartTime       time.Time                                             `json:"startTime" format:"date-time"`
	JSON            bgpTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// bgpTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata
// for the struct [BGPTimeseriesResponseMetaConfidenceInfoAnnotation]
type bgpTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *BGPTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type BGPTimeseriesResponseSerie0 struct {
	Timestamps []time.Time                     `json:"timestamps,required" format:"date-time"`
	Values     []string                        `json:"values,required"`
	JSON       bgpTimeseriesResponseSerie0JSON `json:"-"`
}

// bgpTimeseriesResponseSerie0JSON contains the JSON metadata for the struct
// [BGPTimeseriesResponseSerie0]
type bgpTimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type BGPTimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[BGPTimeseriesParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[BGPTimeseriesParamsFormat] `query:"format"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of BGP network prefixes.
	Prefix param.Field[[]string] `query:"prefix"`
	// Array of BGP update types.
	UpdateType param.Field[[]BGPTimeseriesParamsUpdateType] `query:"updateType"`
}

// URLQuery serializes [BGPTimeseriesParams]'s query parameters as `url.Values`.
func (r BGPTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type BGPTimeseriesParamsAggInterval string

const (
	BGPTimeseriesParamsAggInterval15m BGPTimeseriesParamsAggInterval = "15m"
	BGPTimeseriesParamsAggInterval1h  BGPTimeseriesParamsAggInterval = "1h"
	BGPTimeseriesParamsAggInterval1d  BGPTimeseriesParamsAggInterval = "1d"
	BGPTimeseriesParamsAggInterval1w  BGPTimeseriesParamsAggInterval = "1w"
)

func (r BGPTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case BGPTimeseriesParamsAggInterval15m, BGPTimeseriesParamsAggInterval1h, BGPTimeseriesParamsAggInterval1d, BGPTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format results are returned in.
type BGPTimeseriesParamsFormat string

const (
	BGPTimeseriesParamsFormatJson BGPTimeseriesParamsFormat = "JSON"
	BGPTimeseriesParamsFormatCsv  BGPTimeseriesParamsFormat = "CSV"
)

func (r BGPTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case BGPTimeseriesParamsFormatJson, BGPTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type BGPTimeseriesParamsUpdateType string

const (
	BGPTimeseriesParamsUpdateTypeAnnouncement BGPTimeseriesParamsUpdateType = "ANNOUNCEMENT"
	BGPTimeseriesParamsUpdateTypeWithdrawal   BGPTimeseriesParamsUpdateType = "WITHDRAWAL"
)

func (r BGPTimeseriesParamsUpdateType) IsKnown() bool {
	switch r {
	case BGPTimeseriesParamsUpdateTypeAnnouncement, BGPTimeseriesParamsUpdateTypeWithdrawal:
		return true
	}
	return false
}

type BGPTimeseriesResponseEnvelope struct {
	Result  BGPTimeseriesResponse             `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    bgpTimeseriesResponseEnvelopeJSON `json:"-"`
}

// bgpTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [BGPTimeseriesResponseEnvelope]
type bgpTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
