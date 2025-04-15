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

// AIInferenceTimeseriesGroupSummaryService contains methods and other services
// that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIInferenceTimeseriesGroupSummaryService] method instead.
type AIInferenceTimeseriesGroupSummaryService struct {
	Options []option.RequestOption
}

// NewAIInferenceTimeseriesGroupSummaryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAIInferenceTimeseriesGroupSummaryService(opts ...option.RequestOption) (r *AIInferenceTimeseriesGroupSummaryService) {
	r = &AIInferenceTimeseriesGroupSummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of unique accounts by model over time.
func (r *AIInferenceTimeseriesGroupSummaryService) Model(ctx context.Context, query AIInferenceTimeseriesGroupSummaryModelParams, opts ...option.RequestOption) (res *AIInferenceTimeseriesGroupSummaryModelResponse, err error) {
	var env AIInferenceTimeseriesGroupSummaryModelResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/inference/timeseries_groups/model"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of unique accounts by task over time.
func (r *AIInferenceTimeseriesGroupSummaryService) Task(ctx context.Context, query AIInferenceTimeseriesGroupSummaryTaskParams, opts ...option.RequestOption) (res *AIInferenceTimeseriesGroupSummaryTaskResponse, err error) {
	var env AIInferenceTimeseriesGroupSummaryTaskResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/inference/timeseries_groups/task"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AIInferenceTimeseriesGroupSummaryModelResponse struct {
	Meta   interface{}                                          `json:"meta,required"`
	Serie0 AIInferenceTimeseriesGroupSummaryModelResponseSerie0 `json:"serie_0,required"`
	JSON   aiInferenceTimeseriesGroupSummaryModelResponseJSON   `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryModelResponseJSON contains the JSON metadata
// for the struct [AIInferenceTimeseriesGroupSummaryModelResponse]
type aiInferenceTimeseriesGroupSummaryModelResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryModelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryModelResponseJSON) RawJSON() string {
	return r.raw
}

type AIInferenceTimeseriesGroupSummaryModelResponseSerie0 struct {
	Timestamps  []string                                                 `json:"timestamps,required"`
	ExtraFields map[string][]string                                      `json:"-,extras"`
	JSON        aiInferenceTimeseriesGroupSummaryModelResponseSerie0JSON `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryModelResponseSerie0JSON contains the JSON
// metadata for the struct [AIInferenceTimeseriesGroupSummaryModelResponseSerie0]
type aiInferenceTimeseriesGroupSummaryModelResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryModelResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryModelResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AIInferenceTimeseriesGroupSummaryTaskResponse struct {
	Meta   interface{}                                         `json:"meta,required"`
	Serie0 AIInferenceTimeseriesGroupSummaryTaskResponseSerie0 `json:"serie_0,required"`
	JSON   aiInferenceTimeseriesGroupSummaryTaskResponseJSON   `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryTaskResponseJSON contains the JSON metadata for
// the struct [AIInferenceTimeseriesGroupSummaryTaskResponse]
type aiInferenceTimeseriesGroupSummaryTaskResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryTaskResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryTaskResponseJSON) RawJSON() string {
	return r.raw
}

type AIInferenceTimeseriesGroupSummaryTaskResponseSerie0 struct {
	Timestamps  []string                                                `json:"timestamps,required"`
	ExtraFields map[string][]string                                     `json:"-,extras"`
	JSON        aiInferenceTimeseriesGroupSummaryTaskResponseSerie0JSON `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryTaskResponseSerie0JSON contains the JSON
// metadata for the struct [AIInferenceTimeseriesGroupSummaryTaskResponseSerie0]
type aiInferenceTimeseriesGroupSummaryTaskResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryTaskResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryTaskResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AIInferenceTimeseriesGroupSummaryModelParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AIInferenceTimeseriesGroupSummaryModelParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AIInferenceTimeseriesGroupSummaryModelParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AIInferenceTimeseriesGroupSummaryModelParams]'s query
// parameters as `url.Values`.
func (r AIInferenceTimeseriesGroupSummaryModelParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIInferenceTimeseriesGroupSummaryModelParamsAggInterval string

const (
	AIInferenceTimeseriesGroupSummaryModelParamsAggInterval15m AIInferenceTimeseriesGroupSummaryModelParamsAggInterval = "15m"
	AIInferenceTimeseriesGroupSummaryModelParamsAggInterval1h  AIInferenceTimeseriesGroupSummaryModelParamsAggInterval = "1h"
	AIInferenceTimeseriesGroupSummaryModelParamsAggInterval1d  AIInferenceTimeseriesGroupSummaryModelParamsAggInterval = "1d"
	AIInferenceTimeseriesGroupSummaryModelParamsAggInterval1w  AIInferenceTimeseriesGroupSummaryModelParamsAggInterval = "1w"
)

func (r AIInferenceTimeseriesGroupSummaryModelParamsAggInterval) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupSummaryModelParamsAggInterval15m, AIInferenceTimeseriesGroupSummaryModelParamsAggInterval1h, AIInferenceTimeseriesGroupSummaryModelParamsAggInterval1d, AIInferenceTimeseriesGroupSummaryModelParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AIInferenceTimeseriesGroupSummaryModelParamsFormat string

const (
	AIInferenceTimeseriesGroupSummaryModelParamsFormatJson AIInferenceTimeseriesGroupSummaryModelParamsFormat = "JSON"
	AIInferenceTimeseriesGroupSummaryModelParamsFormatCsv  AIInferenceTimeseriesGroupSummaryModelParamsFormat = "CSV"
)

func (r AIInferenceTimeseriesGroupSummaryModelParamsFormat) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupSummaryModelParamsFormatJson, AIInferenceTimeseriesGroupSummaryModelParamsFormatCsv:
		return true
	}
	return false
}

type AIInferenceTimeseriesGroupSummaryModelResponseEnvelope struct {
	Result  AIInferenceTimeseriesGroupSummaryModelResponse             `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    aiInferenceTimeseriesGroupSummaryModelResponseEnvelopeJSON `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryModelResponseEnvelopeJSON contains the JSON
// metadata for the struct [AIInferenceTimeseriesGroupSummaryModelResponseEnvelope]
type aiInferenceTimeseriesGroupSummaryModelResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryModelResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryModelResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIInferenceTimeseriesGroupSummaryTaskParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AIInferenceTimeseriesGroupSummaryTaskParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AIInferenceTimeseriesGroupSummaryTaskParams]'s query
// parameters as `url.Values`.
func (r AIInferenceTimeseriesGroupSummaryTaskParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval string

const (
	AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval15m AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval = "15m"
	AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval1h  AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval = "1h"
	AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval1d  AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval = "1d"
	AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval1w  AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval = "1w"
)

func (r AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval15m, AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval1h, AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval1d, AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AIInferenceTimeseriesGroupSummaryTaskParamsFormat string

const (
	AIInferenceTimeseriesGroupSummaryTaskParamsFormatJson AIInferenceTimeseriesGroupSummaryTaskParamsFormat = "JSON"
	AIInferenceTimeseriesGroupSummaryTaskParamsFormatCsv  AIInferenceTimeseriesGroupSummaryTaskParamsFormat = "CSV"
)

func (r AIInferenceTimeseriesGroupSummaryTaskParamsFormat) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupSummaryTaskParamsFormatJson, AIInferenceTimeseriesGroupSummaryTaskParamsFormatCsv:
		return true
	}
	return false
}

type AIInferenceTimeseriesGroupSummaryTaskResponseEnvelope struct {
	Result  AIInferenceTimeseriesGroupSummaryTaskResponse             `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    aiInferenceTimeseriesGroupSummaryTaskResponseEnvelopeJSON `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryTaskResponseEnvelopeJSON contains the JSON
// metadata for the struct [AIInferenceTimeseriesGroupSummaryTaskResponseEnvelope]
type aiInferenceTimeseriesGroupSummaryTaskResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryTaskResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryTaskResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
