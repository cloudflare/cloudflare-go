// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// ThreatEventAggregateService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventAggregateService] method instead.
type ThreatEventAggregateService struct {
	Options []option.RequestOption
}

// NewThreatEventAggregateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatEventAggregateService(opts ...option.RequestOption) (r *ThreatEventAggregateService) {
	r = &ThreatEventAggregateService{}
	r.Options = opts
	return
}

// Aggregate threat events by one or more columns (e.g., attacker, targetIndustry)
// with optional date filtering and daily grouping. Supports multi-dimensional
// aggregation for cross-analysis.
func (r *ThreatEventAggregateService) List(ctx context.Context, params ThreatEventAggregateListParams, opts ...option.RequestOption) (res *ThreatEventAggregateListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/aggregate", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type ThreatEventAggregateListResponse struct {
	// Column(s) that were aggregated by
	AggregateBy string `json:"aggregateBy" api:"required"`
	// Array of aggregation results with dynamic fields based on aggregateBy columns
	Aggregations []ThreatEventAggregateListResponseAggregation `json:"aggregations" api:"required"`
	// Total number of events in the aggregation
	Total float64 `json:"total" api:"required"`
	// Date range used for filtering
	DateRange ThreatEventAggregateListResponseDateRange `json:"dateRange"`
	JSON      threatEventAggregateListResponseJSON      `json:"-"`
}

// threatEventAggregateListResponseJSON contains the JSON metadata for the struct
// [ThreatEventAggregateListResponse]
type threatEventAggregateListResponseJSON struct {
	AggregateBy  apijson.Field
	Aggregations apijson.Field
	Total        apijson.Field
	DateRange    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ThreatEventAggregateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventAggregateListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventAggregateListResponseAggregation struct {
	// Number of events for this aggregation
	Count float64 `json:"count" api:"required"`
	// Date (if groupByDate is true)
	Date        string                                          `json:"date"`
	ExtraFields map[string]string                               `json:"-" api:"extrafields"`
	JSON        threatEventAggregateListResponseAggregationJSON `json:"-"`
}

// threatEventAggregateListResponseAggregationJSON contains the JSON metadata for
// the struct [ThreatEventAggregateListResponseAggregation]
type threatEventAggregateListResponseAggregationJSON struct {
	Count       apijson.Field
	Date        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventAggregateListResponseAggregation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventAggregateListResponseAggregationJSON) RawJSON() string {
	return r.raw
}

// Date range used for filtering
type ThreatEventAggregateListResponseDateRange struct {
	EndDate   string                                        `json:"endDate"`
	StartDate string                                        `json:"startDate"`
	JSON      threatEventAggregateListResponseDateRangeJSON `json:"-"`
}

// threatEventAggregateListResponseDateRangeJSON contains the JSON metadata for the
// struct [ThreatEventAggregateListResponseDateRange]
type threatEventAggregateListResponseDateRangeJSON struct {
	EndDate     apijson.Field
	StartDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventAggregateListResponseDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventAggregateListResponseDateRangeJSON) RawJSON() string {
	return r.raw
}

type ThreatEventAggregateListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Column(s) to aggregate by - single column or comma-separated list (e.g.,
	// 'attacker', 'targetIndustry', 'attacker,targetIndustry')
	AggregateBy param.Field[string] `query:"aggregateBy" api:"required"`
	// Dataset ID(s) to filter by. Can be a single dataset ID, comma-separated list, or
	// array. If not provided, uses default dataset
	DatasetID param.Field[[]string] `query:"datasetId"`
	// End date for filtering (ISO 8601 format, e.g., '2024-12-31')
	EndDate param.Field[string] `query:"endDate"`
	// Whether to group results by date (daily aggregation)
	GroupByDate param.Field[bool] `query:"groupByDate"`
	// Maximum number of results to return
	Limit param.Field[float64] `query:"limit"`
	// Start date for filtering (ISO 8601 format, e.g., '2024-01-01')
	StartDate param.Field[string] `query:"startDate"`
}

// URLQuery serializes [ThreatEventAggregateListParams]'s query parameters as
// `url.Values`.
func (r ThreatEventAggregateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
