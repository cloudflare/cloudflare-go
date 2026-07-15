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
)

// AnalyticsQueryDataSecurityFindingService contains methods and other services
// that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsQueryDataSecurityFindingService] method instead.
type AnalyticsQueryDataSecurityFindingService struct {
	Options []option.RequestOption
}

// NewAnalyticsQueryDataSecurityFindingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAnalyticsQueryDataSecurityFindingService(opts ...option.RequestOption) (r *AnalyticsQueryDataSecurityFindingService) {
	r = &AnalyticsQueryDataSecurityFindingService{}
	r.Options = opts
	return
}

// Returns aggregate current-period and previous-period totals for CASB findings.
func (r *AnalyticsQueryDataSecurityFindingService) Summary(ctx context.Context, params AnalyticsQueryDataSecurityFindingSummaryParams, opts ...option.RequestOption) (res *AnalyticsQueryDataSecurityFindingSummaryResponse, err error) {
	var env AnalyticsQueryDataSecurityFindingSummaryResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/analytics/query/data-security/findings/summary", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns merged time-bucketed CASB findings.
func (r *AnalyticsQueryDataSecurityFindingService) Timeseries(ctx context.Context, params AnalyticsQueryDataSecurityFindingTimeseriesParams, opts ...option.RequestOption) (res *AnalyticsQueryDataSecurityFindingTimeseriesResponse, err error) {
	var env AnalyticsQueryDataSecurityFindingTimeseriesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/analytics/query/data-security/findings/timeseries", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AnalyticsQueryDataSecurityFindingSummaryResponse struct {
	// Aggregated stats for the requested time range.
	CurrentTotal []map[string]interface{} `json:"currentTotal" api:"required"`
	// Aggregated stats for the equivalent preceding time range, for trend comparison.
	PreviousTotal []map[string]interface{}                             `json:"previousTotal" api:"required"`
	JSON          analyticsQueryDataSecurityFindingSummaryResponseJSON `json:"-"`
}

// analyticsQueryDataSecurityFindingSummaryResponseJSON contains the JSON metadata
// for the struct [AnalyticsQueryDataSecurityFindingSummaryResponse]
type analyticsQueryDataSecurityFindingSummaryResponseJSON struct {
	CurrentTotal  apijson.Field
	PreviousTotal apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AnalyticsQueryDataSecurityFindingSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsQueryDataSecurityFindingSummaryResponseJSON) RawJSON() string {
	return r.raw
}

// Merged CASB and CDE findings timeseries result.
type AnalyticsQueryDataSecurityFindingTimeseriesResponse struct {
	// Contains time-bucketed result rows. Each slot includes a `timestamp` plus
	// `content` and `posture` maps with `cloud` and `saas` keys.
	Slots []map[string]interface{} `json:"slots" api:"required"`
	// Always null for this endpoint.
	Resolution string                                                  `json:"resolution" api:"nullable"`
	JSON       analyticsQueryDataSecurityFindingTimeseriesResponseJSON `json:"-"`
}

// analyticsQueryDataSecurityFindingTimeseriesResponseJSON contains the JSON
// metadata for the struct [AnalyticsQueryDataSecurityFindingTimeseriesResponse]
type analyticsQueryDataSecurityFindingTimeseriesResponseJSON struct {
	Slots       apijson.Field
	Resolution  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsQueryDataSecurityFindingTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsQueryDataSecurityFindingTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type AnalyticsQueryDataSecurityFindingSummaryParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filters to apply.
	Filters param.Field[[]AnalyticsQueryDataSecurityFindingSummaryParamsFilter] `json:"filters" api:"required"`
	// Start of the query time range (inclusive). RFC3339.
	From param.Field[time.Time] `json:"from" api:"required" format:"date-time"`
	// End of the query time range (exclusive). RFC3339.
	To param.Field[time.Time] `json:"to" api:"required" format:"date-time"`
}

func (r AnalyticsQueryDataSecurityFindingSummaryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A filter to apply to the query.
type AnalyticsQueryDataSecurityFindingSummaryParamsFilter struct {
	// Specifies the column name to filter on. Requires a valid column for the target
	// dataset (e.g. `country`, `allowed`, `appId`).
	Name param.Field[string] `json:"name" api:"required"`
	// Filter operator. Common values: `eq`, `neq`, `in`, `not_in`, `gt`, `lt`, `gte`,
	// `lte`.
	Op param.Field[string] `json:"op" api:"required"`
	// Values to match against. Type depends on the column.
	Values param.Field[[]AnalyticsQueryDataSecurityFindingSummaryParamsFiltersValueUnion] `json:"values" api:"required"`
}

func (r AnalyticsQueryDataSecurityFindingSummaryParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionBool], [shared.UnionFloat].
type AnalyticsQueryDataSecurityFindingSummaryParamsFiltersValueUnion interface {
	ImplementsAnalyticsQueryDataSecurityFindingSummaryParamsFiltersValueUnion()
}

type AnalyticsQueryDataSecurityFindingSummaryResponseEnvelope struct {
	Errors   []AnalyticsQueryDataSecurityFindingSummaryResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AnalyticsQueryDataSecurityFindingSummaryResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   AnalyticsQueryDataSecurityFindingSummaryResponse                   `json:"result" api:"required"`
	Success  bool                                                               `json:"success" api:"required"`
	JSON     analyticsQueryDataSecurityFindingSummaryResponseEnvelopeJSON       `json:"-"`
}

// analyticsQueryDataSecurityFindingSummaryResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [AnalyticsQueryDataSecurityFindingSummaryResponseEnvelope]
type analyticsQueryDataSecurityFindingSummaryResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsQueryDataSecurityFindingSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsQueryDataSecurityFindingSummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AnalyticsQueryDataSecurityFindingSummaryResponseEnvelopeErrors struct {
	Code    int64                                                              `json:"code"`
	Message string                                                             `json:"message"`
	JSON    analyticsQueryDataSecurityFindingSummaryResponseEnvelopeErrorsJSON `json:"-"`
}

// analyticsQueryDataSecurityFindingSummaryResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [AnalyticsQueryDataSecurityFindingSummaryResponseEnvelopeErrors]
type analyticsQueryDataSecurityFindingSummaryResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsQueryDataSecurityFindingSummaryResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsQueryDataSecurityFindingSummaryResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AnalyticsQueryDataSecurityFindingSummaryResponseEnvelopeMessages struct {
	Code    int64                                                                `json:"code"`
	Message string                                                               `json:"message"`
	JSON    analyticsQueryDataSecurityFindingSummaryResponseEnvelopeMessagesJSON `json:"-"`
}

// analyticsQueryDataSecurityFindingSummaryResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [AnalyticsQueryDataSecurityFindingSummaryResponseEnvelopeMessages]
type analyticsQueryDataSecurityFindingSummaryResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsQueryDataSecurityFindingSummaryResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsQueryDataSecurityFindingSummaryResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AnalyticsQueryDataSecurityFindingTimeseriesParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filters to apply.
	Filters param.Field[[]AnalyticsQueryDataSecurityFindingTimeseriesParamsFilter] `json:"filters" api:"required"`
	// Start of the query time range (inclusive). RFC3339.
	From param.Field[time.Time] `json:"from" api:"required" format:"date-time"`
	// End of the query time range (exclusive). RFC3339.
	To param.Field[time.Time] `json:"to" api:"required" format:"date-time"`
}

func (r AnalyticsQueryDataSecurityFindingTimeseriesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A filter to apply to the query.
type AnalyticsQueryDataSecurityFindingTimeseriesParamsFilter struct {
	// Specifies the column name to filter on. Requires a valid column for the target
	// dataset (e.g. `country`, `allowed`, `appId`).
	Name param.Field[string] `json:"name" api:"required"`
	// Filter operator. Common values: `eq`, `neq`, `in`, `not_in`, `gt`, `lt`, `gte`,
	// `lte`.
	Op param.Field[string] `json:"op" api:"required"`
	// Values to match against. Type depends on the column.
	Values param.Field[[]AnalyticsQueryDataSecurityFindingTimeseriesParamsFiltersValueUnion] `json:"values" api:"required"`
}

func (r AnalyticsQueryDataSecurityFindingTimeseriesParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionBool], [shared.UnionFloat].
type AnalyticsQueryDataSecurityFindingTimeseriesParamsFiltersValueUnion interface {
	ImplementsAnalyticsQueryDataSecurityFindingTimeseriesParamsFiltersValueUnion()
}

type AnalyticsQueryDataSecurityFindingTimeseriesResponseEnvelope struct {
	Errors   []AnalyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AnalyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeMessages `json:"messages" api:"required"`
	// Merged CASB and CDE findings timeseries result.
	Result  AnalyticsQueryDataSecurityFindingTimeseriesResponse             `json:"result" api:"required"`
	Success bool                                                            `json:"success" api:"required"`
	JSON    analyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeJSON `json:"-"`
}

// analyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [AnalyticsQueryDataSecurityFindingTimeseriesResponseEnvelope]
type analyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsQueryDataSecurityFindingTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AnalyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeErrors struct {
	Code    int64                                                                 `json:"code"`
	Message string                                                                `json:"message"`
	JSON    analyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeErrorsJSON `json:"-"`
}

// analyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [AnalyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeErrors]
type analyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AnalyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeMessages struct {
	Code    int64                                                                   `json:"code"`
	Message string                                                                  `json:"message"`
	JSON    analyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeMessagesJSON `json:"-"`
}

// analyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [AnalyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeMessages]
type analyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsQueryDataSecurityFindingTimeseriesResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
