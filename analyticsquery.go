// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// AnalyticsQueryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsQueryService] method instead.
type AnalyticsQueryService struct {
	Options      []option.RequestOption
	DataSecurity *AnalyticsQueryDataSecurityService
}

// NewAnalyticsQueryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAnalyticsQueryService(opts ...option.RequestOption) (r *AnalyticsQueryService) {
	r = &AnalyticsQueryService{}
	r.Options = opts
	r.DataSecurity = NewAnalyticsQueryDataSecurityService(opts...)
	return
}

// Returns aggregate summary stats for a dataset. Includes current-period and
// previous-period totals for trend comparison.
func (r *AnalyticsQueryService) Summary(ctx context.Context, dataset string, params AnalyticsQuerySummaryParams, opts ...option.RequestOption) (res *AnalyticsQuerySummaryResponse, err error) {
	var env AnalyticsQuerySummaryResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if dataset == "" {
		err = errors.New("missing required dataset parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/analytics/query/%s/summary", params.AccountID, dataset)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns time-bucketed analytics data for a dataset. Includes time slots, each
// containing the requested stats, group-by dimensions, and resolution-controlled
// bucket size (e.g. `hour`, `day`).
func (r *AnalyticsQueryService) Timeseries(ctx context.Context, dataset string, params AnalyticsQueryTimeseriesParams, opts ...option.RequestOption) (res *AnalyticsQueryTimeseriesResponse, err error) {
	var env AnalyticsQueryTimeseriesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if dataset == "" {
		err = errors.New("missing required dataset parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/analytics/query/%s/timeseries", params.AccountID, dataset)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the top N results for a dataset by a specified stat. Includes an array
// of result rows, each containing the requested stats and group-by dimensions.
func (r *AnalyticsQueryService) TopN(ctx context.Context, dataset string, params AnalyticsQueryTopNParams, opts ...option.RequestOption) (res *pagination.SinglePage[AnalyticsQueryTopNResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if dataset == "" {
		err = errors.New("missing required dataset parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/analytics/query/%s/top-n", params.AccountID, dataset)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns the top N results for a dataset by a specified stat. Includes an array
// of result rows, each containing the requested stats and group-by dimensions.
func (r *AnalyticsQueryService) TopNAutoPaging(ctx context.Context, dataset string, params AnalyticsQueryTopNParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AnalyticsQueryTopNResponse] {
	return pagination.NewSinglePageAutoPager(r.TopN(ctx, dataset, params, opts...))
}

type AnalyticsQuerySummaryResponse struct {
	// Aggregated stats for the requested time range.
	CurrentTotal []map[string]interface{} `json:"currentTotal" api:"required"`
	// Aggregated stats for the equivalent preceding time range, for trend comparison.
	PreviousTotal []map[string]interface{}          `json:"previousTotal" api:"required"`
	JSON          analyticsQuerySummaryResponseJSON `json:"-"`
}

// analyticsQuerySummaryResponseJSON contains the JSON metadata for the struct
// [AnalyticsQuerySummaryResponse]
type analyticsQuerySummaryResponseJSON struct {
	CurrentTotal  apijson.Field
	PreviousTotal apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AnalyticsQuerySummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsQuerySummaryResponseJSON) RawJSON() string {
	return r.raw
}

type AnalyticsQueryTimeseriesResponse struct {
	// The resolution used for time bucketing.
	Resolution string `json:"resolution" api:"required"`
	// Time-bucketed result rows. Each slot contains a `time_bucket` field plus the
	// requested stats and group-by dimensions.
	Slots []map[string]interface{}             `json:"slots" api:"required"`
	JSON  analyticsQueryTimeseriesResponseJSON `json:"-"`
}

// analyticsQueryTimeseriesResponseJSON contains the JSON metadata for the struct
// [AnalyticsQueryTimeseriesResponse]
type analyticsQueryTimeseriesResponseJSON struct {
	Resolution  apijson.Field
	Slots       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsQueryTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsQueryTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type AnalyticsQueryTopNResponse map[string]interface{}

type AnalyticsQuerySummaryParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filters to apply before aggregating results.
	Filters param.Field[[]AnalyticsQuerySummaryParamsFilter] `json:"filters" api:"required"`
	// The start of the query time range (inclusive). RFC3339 format with timezone is
	// required (e.g. `2024-11-05T00:00:00Z`).
	From param.Field[time.Time] `json:"from" api:"required" format:"date-time"`
	// Specifies the column names to group results by. Requires valid columns for the
	// target dataset.
	GroupBy param.Field[[]string] `json:"groupBy" api:"required"`
	// Specifies the stat names to include in results. Requires valid stats for the
	// target dataset (e.g. `attemptsTotal`, `bytesTotal`).
	Stats param.Field[[]string] `json:"stats" api:"required"`
	// Specifies the end of the query time range (exclusive). Requires RFC3339 format
	// with timezone.
	To param.Field[time.Time] `json:"to" api:"required" format:"date-time"`
}

func (r AnalyticsQuerySummaryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A filter to apply to the query.
type AnalyticsQuerySummaryParamsFilter struct {
	// Specifies the column name to filter on. Requires a valid column for the target
	// dataset (e.g. `country`, `allowed`, `appId`).
	Name param.Field[string] `json:"name" api:"required"`
	// Filter operator. Common values: `eq`, `neq`, `in`, `not_in`, `gt`, `lt`, `gte`,
	// `lte`.
	Op param.Field[string] `json:"op" api:"required"`
	// Values to match against. Type depends on the column.
	Values param.Field[[]AnalyticsQuerySummaryParamsFiltersValueUnion] `json:"values" api:"required"`
}

func (r AnalyticsQuerySummaryParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionBool], [shared.UnionFloat].
type AnalyticsQuerySummaryParamsFiltersValueUnion interface {
	ImplementsAnalyticsQuerySummaryParamsFiltersValueUnion()
}

type AnalyticsQuerySummaryResponseEnvelope struct {
	Errors   []AnalyticsQuerySummaryResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AnalyticsQuerySummaryResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   AnalyticsQuerySummaryResponse                   `json:"result" api:"required"`
	Success  bool                                            `json:"success" api:"required"`
	JSON     analyticsQuerySummaryResponseEnvelopeJSON       `json:"-"`
}

// analyticsQuerySummaryResponseEnvelopeJSON contains the JSON metadata for the
// struct [AnalyticsQuerySummaryResponseEnvelope]
type analyticsQuerySummaryResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsQuerySummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsQuerySummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AnalyticsQuerySummaryResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code"`
	Message string                                          `json:"message"`
	JSON    analyticsQuerySummaryResponseEnvelopeErrorsJSON `json:"-"`
}

// analyticsQuerySummaryResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AnalyticsQuerySummaryResponseEnvelopeErrors]
type analyticsQuerySummaryResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsQuerySummaryResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsQuerySummaryResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AnalyticsQuerySummaryResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code"`
	Message string                                            `json:"message"`
	JSON    analyticsQuerySummaryResponseEnvelopeMessagesJSON `json:"-"`
}

// analyticsQuerySummaryResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AnalyticsQuerySummaryResponseEnvelopeMessages]
type analyticsQuerySummaryResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsQuerySummaryResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsQuerySummaryResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AnalyticsQueryTimeseriesParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filters to apply before aggregating results.
	Filters param.Field[[]AnalyticsQueryTimeseriesParamsFilter] `json:"filters" api:"required"`
	// The start of the query time range (inclusive). RFC3339 format with timezone is
	// required (e.g. `2024-11-05T00:00:00Z`).
	From param.Field[time.Time] `json:"from" api:"required" format:"date-time"`
	// Specifies the column names to group results by. Requires valid columns for the
	// target dataset.
	GroupBy param.Field[[]string] `json:"groupBy" api:"required"`
	// Time bucket size for grouping results. Controls the granularity of the returned
	// time slots.
	Resolution param.Field[string] `json:"resolution" api:"required"`
	// Specifies the stat names to include in results. Requires valid stats for the
	// target dataset (e.g. `attemptsTotal`, `bytesTotal`).
	Stats param.Field[[]string] `json:"stats" api:"required"`
	// Specifies the end of the query time range (exclusive). Requires RFC3339 format
	// with timezone.
	To param.Field[time.Time] `json:"to" api:"required" format:"date-time"`
}

func (r AnalyticsQueryTimeseriesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A filter to apply to the query.
type AnalyticsQueryTimeseriesParamsFilter struct {
	// Specifies the column name to filter on. Requires a valid column for the target
	// dataset (e.g. `country`, `allowed`, `appId`).
	Name param.Field[string] `json:"name" api:"required"`
	// Filter operator. Common values: `eq`, `neq`, `in`, `not_in`, `gt`, `lt`, `gte`,
	// `lte`.
	Op param.Field[string] `json:"op" api:"required"`
	// Values to match against. Type depends on the column.
	Values param.Field[[]AnalyticsQueryTimeseriesParamsFiltersValueUnion] `json:"values" api:"required"`
}

func (r AnalyticsQueryTimeseriesParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionBool], [shared.UnionFloat].
type AnalyticsQueryTimeseriesParamsFiltersValueUnion interface {
	ImplementsAnalyticsQueryTimeseriesParamsFiltersValueUnion()
}

type AnalyticsQueryTimeseriesResponseEnvelope struct {
	Errors   []AnalyticsQueryTimeseriesResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AnalyticsQueryTimeseriesResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   AnalyticsQueryTimeseriesResponse                   `json:"result" api:"required"`
	Success  bool                                               `json:"success" api:"required"`
	JSON     analyticsQueryTimeseriesResponseEnvelopeJSON       `json:"-"`
}

// analyticsQueryTimeseriesResponseEnvelopeJSON contains the JSON metadata for the
// struct [AnalyticsQueryTimeseriesResponseEnvelope]
type analyticsQueryTimeseriesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsQueryTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsQueryTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AnalyticsQueryTimeseriesResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code"`
	Message string                                             `json:"message"`
	JSON    analyticsQueryTimeseriesResponseEnvelopeErrorsJSON `json:"-"`
}

// analyticsQueryTimeseriesResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AnalyticsQueryTimeseriesResponseEnvelopeErrors]
type analyticsQueryTimeseriesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsQueryTimeseriesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsQueryTimeseriesResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AnalyticsQueryTimeseriesResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code"`
	Message string                                               `json:"message"`
	JSON    analyticsQueryTimeseriesResponseEnvelopeMessagesJSON `json:"-"`
}

// analyticsQueryTimeseriesResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AnalyticsQueryTimeseriesResponseEnvelopeMessages]
type analyticsQueryTimeseriesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsQueryTimeseriesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsQueryTimeseriesResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AnalyticsQueryTopNParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filters to apply before aggregating results.
	Filters param.Field[[]AnalyticsQueryTopNParamsFilter] `json:"filters" api:"required"`
	// The start of the query time range (inclusive). RFC3339 format with timezone is
	// required (e.g. `2024-11-05T00:00:00Z`).
	From param.Field[time.Time] `json:"from" api:"required" format:"date-time"`
	// Specifies the column names to group results by. Requires valid columns for the
	// target dataset.
	GroupBy param.Field[[]string] `json:"groupBy" api:"required"`
	// Maximum number of results to return.
	N param.Field[int64] `json:"n" api:"required"`
	// Specifies the stat name for sorting results in descending order. Requires a
	// valid stat for the target dataset.
	OrderBy param.Field[string] `json:"orderBy" api:"required"`
	// Specifies the stat names to include in results. Requires valid stats for the
	// target dataset (e.g. `attemptsTotal`, `bytesTotal`).
	Stats param.Field[[]string] `json:"stats" api:"required"`
	// Specifies the end of the query time range (exclusive). Requires RFC3339 format
	// with timezone.
	To param.Field[time.Time] `json:"to" api:"required" format:"date-time"`
}

func (r AnalyticsQueryTopNParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A filter to apply to the query.
type AnalyticsQueryTopNParamsFilter struct {
	// Specifies the column name to filter on. Requires a valid column for the target
	// dataset (e.g. `country`, `allowed`, `appId`).
	Name param.Field[string] `json:"name" api:"required"`
	// Filter operator. Common values: `eq`, `neq`, `in`, `not_in`, `gt`, `lt`, `gte`,
	// `lte`.
	Op param.Field[string] `json:"op" api:"required"`
	// Values to match against. Type depends on the column.
	Values param.Field[[]AnalyticsQueryTopNParamsFiltersValueUnion] `json:"values" api:"required"`
}

func (r AnalyticsQueryTopNParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionBool], [shared.UnionFloat].
type AnalyticsQueryTopNParamsFiltersValueUnion interface {
	ImplementsAnalyticsQueryTopNParamsFiltersValueUnion()
}
