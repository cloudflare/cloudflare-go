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

// AITimeseriesGroupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAITimeseriesGroupService] method instead.
type AITimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewAITimeseriesGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAITimeseriesGroupService(opts ...option.RequestOption) (r *AITimeseriesGroupService) {
	r = &AITimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of traffic by AI user agent over time.
func (r *AITimeseriesGroupService) UserAgent(ctx context.Context, query AITimeseriesGroupUserAgentParams, opts ...option.RequestOption) (res *AITimeseriesGroupUserAgentResponse, err error) {
	var env AITimeseriesGroupUserAgentResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/bots/timeseries_groups/user_agent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AITimeseriesGroupUserAgentResponse struct {
	Meta   interface{}                            `json:"meta,required"`
	Serie0 interface{}                            `json:"serie_0,required"`
	JSON   aiTimeseriesGroupUserAgentResponseJSON `json:"-"`
}

// aiTimeseriesGroupUserAgentResponseJSON contains the JSON metadata for the struct
// [AITimeseriesGroupUserAgentResponse]
type aiTimeseriesGroupUserAgentResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupUserAgentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupUserAgentResponseJSON) RawJSON() string {
	return r.raw
}

type AITimeseriesGroupUserAgentParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AITimeseriesGroupUserAgentParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AITimeseriesGroupUserAgentParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. If there are more items than the limit, the response will include
	// the count of items, with any remaining items grouped together under an "other"
	// category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AITimeseriesGroupUserAgentParams]'s query parameters as
// `url.Values`.
func (r AITimeseriesGroupUserAgentParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AITimeseriesGroupUserAgentParamsAggInterval string

const (
	AITimeseriesGroupUserAgentParamsAggInterval15m AITimeseriesGroupUserAgentParamsAggInterval = "15m"
	AITimeseriesGroupUserAgentParamsAggInterval1h  AITimeseriesGroupUserAgentParamsAggInterval = "1h"
	AITimeseriesGroupUserAgentParamsAggInterval1d  AITimeseriesGroupUserAgentParamsAggInterval = "1d"
	AITimeseriesGroupUserAgentParamsAggInterval1w  AITimeseriesGroupUserAgentParamsAggInterval = "1w"
)

func (r AITimeseriesGroupUserAgentParamsAggInterval) IsKnown() bool {
	switch r {
	case AITimeseriesGroupUserAgentParamsAggInterval15m, AITimeseriesGroupUserAgentParamsAggInterval1h, AITimeseriesGroupUserAgentParamsAggInterval1d, AITimeseriesGroupUserAgentParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AITimeseriesGroupUserAgentParamsFormat string

const (
	AITimeseriesGroupUserAgentParamsFormatJson AITimeseriesGroupUserAgentParamsFormat = "JSON"
	AITimeseriesGroupUserAgentParamsFormatCsv  AITimeseriesGroupUserAgentParamsFormat = "CSV"
)

func (r AITimeseriesGroupUserAgentParamsFormat) IsKnown() bool {
	switch r {
	case AITimeseriesGroupUserAgentParamsFormatJson, AITimeseriesGroupUserAgentParamsFormatCsv:
		return true
	}
	return false
}

type AITimeseriesGroupUserAgentResponseEnvelope struct {
	Result  AITimeseriesGroupUserAgentResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    aiTimeseriesGroupUserAgentResponseEnvelopeJSON `json:"-"`
}

// aiTimeseriesGroupUserAgentResponseEnvelopeJSON contains the JSON metadata for
// the struct [AITimeseriesGroupUserAgentResponseEnvelope]
type aiTimeseriesGroupUserAgentResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupUserAgentResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupUserAgentResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
