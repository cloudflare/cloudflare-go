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

// AttackLayer7Service contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackLayer7Service] method instead.
type AttackLayer7Service struct {
	Options          []option.RequestOption
	Summary          *AttackLayer7SummaryService
	TimeseriesGroups *AttackLayer7TimeseriesGroupService
	Top              *AttackLayer7TopService
}

// NewAttackLayer7Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAttackLayer7Service(opts ...option.RequestOption) (r *AttackLayer7Service) {
	r = &AttackLayer7Service{}
	r.Options = opts
	r.Summary = NewAttackLayer7SummaryService(opts...)
	r.TimeseriesGroups = NewAttackLayer7TimeseriesGroupService(opts...)
	r.Top = NewAttackLayer7TopService(opts...)
	return
}

// Get a timeseries of Layer 7 attacks. Values represent HTTP requests and are
// normalized using min-max by default.
func (r *AttackLayer7Service) Timeseries(ctx context.Context, query AttackLayer7TimeseriesParams, opts ...option.RequestOption) (res *AttackLayer7TimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer7TimeseriesResponseEnvelope
	path := "radar/attacks/layer7/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer7TimeseriesResponse struct {
	Meta   AttackLayer7TimeseriesResponseMeta   `json:"meta,required"`
	Serie0 AttackLayer7TimeseriesResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer7TimeseriesResponseJSON   `json:"-"`
}

// attackLayer7TimeseriesResponseJSON contains the JSON metadata for the struct
// [AttackLayer7TimeseriesResponse]
type attackLayer7TimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesResponseMeta struct {
	AggInterval    string                                           `json:"aggInterval,required"`
	DateRange      []AttackLayer7TimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                        `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo AttackLayer7TimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer7TimeseriesResponseMetaJSON           `json:"-"`
}

// attackLayer7TimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [AttackLayer7TimeseriesResponseMeta]
type attackLayer7TimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TimeseriesResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer7TimeseriesResponseMetaDateRange]
type attackLayer7TimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                        `json:"level"`
	JSON        attackLayer7TimeseriesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer7TimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AttackLayer7TimeseriesResponseMetaConfidenceInfo]
type attackLayer7TimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                         `json:"dataSource,required"`
	Description     string                                                         `json:"description,required"`
	EventType       string                                                         `json:"eventType,required"`
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                      `json:"startTime" format:"date-time"`
	JSON            attackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotation]
type attackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesResponseSerie0 struct {
	Timestamps []time.Time                              `json:"timestamps,required" format:"date-time"`
	Values     []string                                 `json:"values,required"`
	JSON       attackLayer7TimeseriesResponseSerie0JSON `json:"-"`
}

// attackLayer7TimeseriesResponseSerie0JSON contains the JSON metadata for the
// struct [AttackLayer7TimeseriesResponseSerie0]
type attackLayer7TimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer7TimeseriesParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of L7 attack types.
	Attack param.Field[[]AttackLayer7TimeseriesParamsAttack] `query:"attack"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]AttackLayer7TimeseriesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AttackLayer7TimeseriesParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer7TimeseriesParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AttackLayer7TimeseriesParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7TimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer7TimeseriesParamsAggInterval string

const (
	AttackLayer7TimeseriesParamsAggInterval15m AttackLayer7TimeseriesParamsAggInterval = "15m"
	AttackLayer7TimeseriesParamsAggInterval1h  AttackLayer7TimeseriesParamsAggInterval = "1h"
	AttackLayer7TimeseriesParamsAggInterval1d  AttackLayer7TimeseriesParamsAggInterval = "1d"
	AttackLayer7TimeseriesParamsAggInterval1w  AttackLayer7TimeseriesParamsAggInterval = "1w"
)

func (r AttackLayer7TimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesParamsAggInterval15m, AttackLayer7TimeseriesParamsAggInterval1h, AttackLayer7TimeseriesParamsAggInterval1d, AttackLayer7TimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

type AttackLayer7TimeseriesParamsAttack string

const (
	AttackLayer7TimeseriesParamsAttackDDoS               AttackLayer7TimeseriesParamsAttack = "DDOS"
	AttackLayer7TimeseriesParamsAttackWAF                AttackLayer7TimeseriesParamsAttack = "WAF"
	AttackLayer7TimeseriesParamsAttackBotManagement      AttackLayer7TimeseriesParamsAttack = "BOT_MANAGEMENT"
	AttackLayer7TimeseriesParamsAttackAccessRules        AttackLayer7TimeseriesParamsAttack = "ACCESS_RULES"
	AttackLayer7TimeseriesParamsAttackIPReputation       AttackLayer7TimeseriesParamsAttack = "IP_REPUTATION"
	AttackLayer7TimeseriesParamsAttackAPIShield          AttackLayer7TimeseriesParamsAttack = "API_SHIELD"
	AttackLayer7TimeseriesParamsAttackDataLossPrevention AttackLayer7TimeseriesParamsAttack = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TimeseriesParamsAttack) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesParamsAttackDDoS, AttackLayer7TimeseriesParamsAttackWAF, AttackLayer7TimeseriesParamsAttackBotManagement, AttackLayer7TimeseriesParamsAttackAccessRules, AttackLayer7TimeseriesParamsAttackIPReputation, AttackLayer7TimeseriesParamsAttackAPIShield, AttackLayer7TimeseriesParamsAttackDataLossPrevention:
		return true
	}
	return false
}

type AttackLayer7TimeseriesParamsDateRange string

const (
	AttackLayer7TimeseriesParamsDateRange1d         AttackLayer7TimeseriesParamsDateRange = "1d"
	AttackLayer7TimeseriesParamsDateRange2d         AttackLayer7TimeseriesParamsDateRange = "2d"
	AttackLayer7TimeseriesParamsDateRange7d         AttackLayer7TimeseriesParamsDateRange = "7d"
	AttackLayer7TimeseriesParamsDateRange14d        AttackLayer7TimeseriesParamsDateRange = "14d"
	AttackLayer7TimeseriesParamsDateRange28d        AttackLayer7TimeseriesParamsDateRange = "28d"
	AttackLayer7TimeseriesParamsDateRange12w        AttackLayer7TimeseriesParamsDateRange = "12w"
	AttackLayer7TimeseriesParamsDateRange24w        AttackLayer7TimeseriesParamsDateRange = "24w"
	AttackLayer7TimeseriesParamsDateRange52w        AttackLayer7TimeseriesParamsDateRange = "52w"
	AttackLayer7TimeseriesParamsDateRange1dControl  AttackLayer7TimeseriesParamsDateRange = "1dControl"
	AttackLayer7TimeseriesParamsDateRange2dControl  AttackLayer7TimeseriesParamsDateRange = "2dControl"
	AttackLayer7TimeseriesParamsDateRange7dControl  AttackLayer7TimeseriesParamsDateRange = "7dControl"
	AttackLayer7TimeseriesParamsDateRange14dControl AttackLayer7TimeseriesParamsDateRange = "14dControl"
	AttackLayer7TimeseriesParamsDateRange28dControl AttackLayer7TimeseriesParamsDateRange = "28dControl"
	AttackLayer7TimeseriesParamsDateRange12wControl AttackLayer7TimeseriesParamsDateRange = "12wControl"
	AttackLayer7TimeseriesParamsDateRange24wControl AttackLayer7TimeseriesParamsDateRange = "24wControl"
)

func (r AttackLayer7TimeseriesParamsDateRange) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesParamsDateRange1d, AttackLayer7TimeseriesParamsDateRange2d, AttackLayer7TimeseriesParamsDateRange7d, AttackLayer7TimeseriesParamsDateRange14d, AttackLayer7TimeseriesParamsDateRange28d, AttackLayer7TimeseriesParamsDateRange12w, AttackLayer7TimeseriesParamsDateRange24w, AttackLayer7TimeseriesParamsDateRange52w, AttackLayer7TimeseriesParamsDateRange1dControl, AttackLayer7TimeseriesParamsDateRange2dControl, AttackLayer7TimeseriesParamsDateRange7dControl, AttackLayer7TimeseriesParamsDateRange14dControl, AttackLayer7TimeseriesParamsDateRange28dControl, AttackLayer7TimeseriesParamsDateRange12wControl, AttackLayer7TimeseriesParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer7TimeseriesParamsFormat string

const (
	AttackLayer7TimeseriesParamsFormatJson AttackLayer7TimeseriesParamsFormat = "JSON"
	AttackLayer7TimeseriesParamsFormatCsv  AttackLayer7TimeseriesParamsFormat = "CSV"
)

func (r AttackLayer7TimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesParamsFormatJson, AttackLayer7TimeseriesParamsFormatCsv:
		return true
	}
	return false
}

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TimeseriesParamsNormalization string

const (
	AttackLayer7TimeseriesParamsNormalizationPercentageChange AttackLayer7TimeseriesParamsNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7TimeseriesParamsNormalizationMin0Max          AttackLayer7TimeseriesParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer7TimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TimeseriesParamsNormalizationPercentageChange, AttackLayer7TimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer7TimeseriesResponseEnvelope struct {
	Result  AttackLayer7TimeseriesResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    attackLayer7TimeseriesResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TimeseriesResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer7TimeseriesResponseEnvelope]
type attackLayer7TimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
