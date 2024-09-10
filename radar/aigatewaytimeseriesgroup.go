// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// AIGatewayTimeseriesGroupService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIGatewayTimeseriesGroupService] method instead.
type AIGatewayTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewAIGatewayTimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAIGatewayTimeseriesGroupService(opts ...option.RequestOption) (r *AIGatewayTimeseriesGroupService) {
	r = &AIGatewayTimeseriesGroupService{}
	r.Options = opts
	return
}

// Percentage distribution of unique accounts per model over time.
func (r *AIGatewayTimeseriesGroupService) Model(ctx context.Context, query AIGatewayTimeseriesGroupModelParams, opts ...option.RequestOption) (res *AIGatewayTimeseriesGroupModelResponse, err error) {
	var env AIGatewayTimeseriesGroupModelResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/inference/timeseries_groups/model"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of unique accounts per provider over time.
func (r *AIGatewayTimeseriesGroupService) Provider(ctx context.Context, query AIGatewayTimeseriesGroupProviderParams, opts ...option.RequestOption) (res *AIGatewayTimeseriesGroupProviderResponse, err error) {
	var env AIGatewayTimeseriesGroupProviderResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/gateway/timeseries_groups/provider"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Percentage distribution of unique accounts per task over time.
func (r *AIGatewayTimeseriesGroupService) Task(ctx context.Context, query AIGatewayTimeseriesGroupTaskParams, opts ...option.RequestOption) (res *AIGatewayTimeseriesGroupTaskResponse, err error) {
	var env AIGatewayTimeseriesGroupTaskResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/inference/timeseries_groups/task"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AIGatewayTimeseriesGroupModelResponse struct {
	Meta   interface{}                                 `json:"meta,required"`
	Serie0 AIGatewayTimeseriesGroupModelResponseSerie0 `json:"serie_0,required"`
	JSON   aiGatewayTimeseriesGroupModelResponseJSON   `json:"-"`
}

// aiGatewayTimeseriesGroupModelResponseJSON contains the JSON metadata for the
// struct [AIGatewayTimeseriesGroupModelResponse]
type aiGatewayTimeseriesGroupModelResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayTimeseriesGroupModelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayTimeseriesGroupModelResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayTimeseriesGroupModelResponseSerie0 struct {
	Timestamps  []string                                        `json:"timestamps,required"`
	ExtraFields map[string][]string                             `json:"-,extras"`
	JSON        aiGatewayTimeseriesGroupModelResponseSerie0JSON `json:"-"`
}

// aiGatewayTimeseriesGroupModelResponseSerie0JSON contains the JSON metadata for
// the struct [AIGatewayTimeseriesGroupModelResponseSerie0]
type aiGatewayTimeseriesGroupModelResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayTimeseriesGroupModelResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayTimeseriesGroupModelResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AIGatewayTimeseriesGroupProviderResponse struct {
	Meta   interface{}                                    `json:"meta,required"`
	Serie0 AIGatewayTimeseriesGroupProviderResponseSerie0 `json:"serie_0,required"`
	JSON   aiGatewayTimeseriesGroupProviderResponseJSON   `json:"-"`
}

// aiGatewayTimeseriesGroupProviderResponseJSON contains the JSON metadata for the
// struct [AIGatewayTimeseriesGroupProviderResponse]
type aiGatewayTimeseriesGroupProviderResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayTimeseriesGroupProviderResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayTimeseriesGroupProviderResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayTimeseriesGroupProviderResponseSerie0 struct {
	Timestamps  []string                                           `json:"timestamps,required"`
	ExtraFields map[string][]string                                `json:"-,extras"`
	JSON        aiGatewayTimeseriesGroupProviderResponseSerie0JSON `json:"-"`
}

// aiGatewayTimeseriesGroupProviderResponseSerie0JSON contains the JSON metadata
// for the struct [AIGatewayTimeseriesGroupProviderResponseSerie0]
type aiGatewayTimeseriesGroupProviderResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayTimeseriesGroupProviderResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayTimeseriesGroupProviderResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AIGatewayTimeseriesGroupTaskResponse struct {
	Meta   interface{}                                `json:"meta,required"`
	Serie0 AIGatewayTimeseriesGroupTaskResponseSerie0 `json:"serie_0,required"`
	JSON   aiGatewayTimeseriesGroupTaskResponseJSON   `json:"-"`
}

// aiGatewayTimeseriesGroupTaskResponseJSON contains the JSON metadata for the
// struct [AIGatewayTimeseriesGroupTaskResponse]
type aiGatewayTimeseriesGroupTaskResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayTimeseriesGroupTaskResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayTimeseriesGroupTaskResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayTimeseriesGroupTaskResponseSerie0 struct {
	Timestamps  []string                                       `json:"timestamps,required"`
	ExtraFields map[string][]string                            `json:"-,extras"`
	JSON        aiGatewayTimeseriesGroupTaskResponseSerie0JSON `json:"-"`
}

// aiGatewayTimeseriesGroupTaskResponseSerie0JSON contains the JSON metadata for
// the struct [AIGatewayTimeseriesGroupTaskResponseSerie0]
type aiGatewayTimeseriesGroupTaskResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayTimeseriesGroupTaskResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayTimeseriesGroupTaskResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AIGatewayTimeseriesGroupModelParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AIGatewayTimeseriesGroupModelParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AIGatewayTimeseriesGroupModelParamsFormat] `query:"format"`
	// Limit the number of objects (eg browsers, verticals, etc) to the top items over
	// the time range.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AIGatewayTimeseriesGroupModelParams]'s query parameters as
// `url.Values`.
func (r AIGatewayTimeseriesGroupModelParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIGatewayTimeseriesGroupModelParamsAggInterval string

const (
	AIGatewayTimeseriesGroupModelParamsAggInterval15m AIGatewayTimeseriesGroupModelParamsAggInterval = "15m"
	AIGatewayTimeseriesGroupModelParamsAggInterval1h  AIGatewayTimeseriesGroupModelParamsAggInterval = "1h"
	AIGatewayTimeseriesGroupModelParamsAggInterval1d  AIGatewayTimeseriesGroupModelParamsAggInterval = "1d"
	AIGatewayTimeseriesGroupModelParamsAggInterval1w  AIGatewayTimeseriesGroupModelParamsAggInterval = "1w"
)

func (r AIGatewayTimeseriesGroupModelParamsAggInterval) IsKnown() bool {
	switch r {
	case AIGatewayTimeseriesGroupModelParamsAggInterval15m, AIGatewayTimeseriesGroupModelParamsAggInterval1h, AIGatewayTimeseriesGroupModelParamsAggInterval1d, AIGatewayTimeseriesGroupModelParamsAggInterval1w:
		return true
	}
	return false
}

// Format results are returned in.
type AIGatewayTimeseriesGroupModelParamsFormat string

const (
	AIGatewayTimeseriesGroupModelParamsFormatJson AIGatewayTimeseriesGroupModelParamsFormat = "JSON"
	AIGatewayTimeseriesGroupModelParamsFormatCsv  AIGatewayTimeseriesGroupModelParamsFormat = "CSV"
)

func (r AIGatewayTimeseriesGroupModelParamsFormat) IsKnown() bool {
	switch r {
	case AIGatewayTimeseriesGroupModelParamsFormatJson, AIGatewayTimeseriesGroupModelParamsFormatCsv:
		return true
	}
	return false
}

type AIGatewayTimeseriesGroupModelResponseEnvelope struct {
	Result  AIGatewayTimeseriesGroupModelResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    aiGatewayTimeseriesGroupModelResponseEnvelopeJSON `json:"-"`
}

// aiGatewayTimeseriesGroupModelResponseEnvelopeJSON contains the JSON metadata for
// the struct [AIGatewayTimeseriesGroupModelResponseEnvelope]
type aiGatewayTimeseriesGroupModelResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayTimeseriesGroupModelResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayTimeseriesGroupModelResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIGatewayTimeseriesGroupProviderParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AIGatewayTimeseriesGroupProviderParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AIGatewayTimeseriesGroupProviderParamsFormat] `query:"format"`
	// Limit the number of objects (eg browsers, verticals, etc) to the top items over
	// the time range.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AIGatewayTimeseriesGroupProviderParams]'s query parameters
// as `url.Values`.
func (r AIGatewayTimeseriesGroupProviderParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIGatewayTimeseriesGroupProviderParamsAggInterval string

const (
	AIGatewayTimeseriesGroupProviderParamsAggInterval15m AIGatewayTimeseriesGroupProviderParamsAggInterval = "15m"
	AIGatewayTimeseriesGroupProviderParamsAggInterval1h  AIGatewayTimeseriesGroupProviderParamsAggInterval = "1h"
	AIGatewayTimeseriesGroupProviderParamsAggInterval1d  AIGatewayTimeseriesGroupProviderParamsAggInterval = "1d"
	AIGatewayTimeseriesGroupProviderParamsAggInterval1w  AIGatewayTimeseriesGroupProviderParamsAggInterval = "1w"
)

func (r AIGatewayTimeseriesGroupProviderParamsAggInterval) IsKnown() bool {
	switch r {
	case AIGatewayTimeseriesGroupProviderParamsAggInterval15m, AIGatewayTimeseriesGroupProviderParamsAggInterval1h, AIGatewayTimeseriesGroupProviderParamsAggInterval1d, AIGatewayTimeseriesGroupProviderParamsAggInterval1w:
		return true
	}
	return false
}

// Format results are returned in.
type AIGatewayTimeseriesGroupProviderParamsFormat string

const (
	AIGatewayTimeseriesGroupProviderParamsFormatJson AIGatewayTimeseriesGroupProviderParamsFormat = "JSON"
	AIGatewayTimeseriesGroupProviderParamsFormatCsv  AIGatewayTimeseriesGroupProviderParamsFormat = "CSV"
)

func (r AIGatewayTimeseriesGroupProviderParamsFormat) IsKnown() bool {
	switch r {
	case AIGatewayTimeseriesGroupProviderParamsFormatJson, AIGatewayTimeseriesGroupProviderParamsFormatCsv:
		return true
	}
	return false
}

type AIGatewayTimeseriesGroupProviderResponseEnvelope struct {
	Result  AIGatewayTimeseriesGroupProviderResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    aiGatewayTimeseriesGroupProviderResponseEnvelopeJSON `json:"-"`
}

// aiGatewayTimeseriesGroupProviderResponseEnvelopeJSON contains the JSON metadata
// for the struct [AIGatewayTimeseriesGroupProviderResponseEnvelope]
type aiGatewayTimeseriesGroupProviderResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayTimeseriesGroupProviderResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayTimeseriesGroupProviderResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIGatewayTimeseriesGroupTaskParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AIGatewayTimeseriesGroupTaskParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AIGatewayTimeseriesGroupTaskParamsFormat] `query:"format"`
	// Limit the number of objects (eg browsers, verticals, etc) to the top items over
	// the time range.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AIGatewayTimeseriesGroupTaskParams]'s query parameters as
// `url.Values`.
func (r AIGatewayTimeseriesGroupTaskParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIGatewayTimeseriesGroupTaskParamsAggInterval string

const (
	AIGatewayTimeseriesGroupTaskParamsAggInterval15m AIGatewayTimeseriesGroupTaskParamsAggInterval = "15m"
	AIGatewayTimeseriesGroupTaskParamsAggInterval1h  AIGatewayTimeseriesGroupTaskParamsAggInterval = "1h"
	AIGatewayTimeseriesGroupTaskParamsAggInterval1d  AIGatewayTimeseriesGroupTaskParamsAggInterval = "1d"
	AIGatewayTimeseriesGroupTaskParamsAggInterval1w  AIGatewayTimeseriesGroupTaskParamsAggInterval = "1w"
)

func (r AIGatewayTimeseriesGroupTaskParamsAggInterval) IsKnown() bool {
	switch r {
	case AIGatewayTimeseriesGroupTaskParamsAggInterval15m, AIGatewayTimeseriesGroupTaskParamsAggInterval1h, AIGatewayTimeseriesGroupTaskParamsAggInterval1d, AIGatewayTimeseriesGroupTaskParamsAggInterval1w:
		return true
	}
	return false
}

// Format results are returned in.
type AIGatewayTimeseriesGroupTaskParamsFormat string

const (
	AIGatewayTimeseriesGroupTaskParamsFormatJson AIGatewayTimeseriesGroupTaskParamsFormat = "JSON"
	AIGatewayTimeseriesGroupTaskParamsFormatCsv  AIGatewayTimeseriesGroupTaskParamsFormat = "CSV"
)

func (r AIGatewayTimeseriesGroupTaskParamsFormat) IsKnown() bool {
	switch r {
	case AIGatewayTimeseriesGroupTaskParamsFormatJson, AIGatewayTimeseriesGroupTaskParamsFormatCsv:
		return true
	}
	return false
}

type AIGatewayTimeseriesGroupTaskResponseEnvelope struct {
	Result  AIGatewayTimeseriesGroupTaskResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    aiGatewayTimeseriesGroupTaskResponseEnvelopeJSON `json:"-"`
}

// aiGatewayTimeseriesGroupTaskResponseEnvelopeJSON contains the JSON metadata for
// the struct [AIGatewayTimeseriesGroupTaskResponseEnvelope]
type aiGatewayTimeseriesGroupTaskResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayTimeseriesGroupTaskResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayTimeseriesGroupTaskResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
