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

// ThreatEventIndicatorAggregateService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventIndicatorAggregateService] method instead.
type ThreatEventIndicatorAggregateService struct {
	Options []option.RequestOption
}

// NewThreatEventIndicatorAggregateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewThreatEventIndicatorAggregateService(opts ...option.RequestOption) (r *ThreatEventIndicatorAggregateService) {
	r = &ThreatEventIndicatorAggregateService{}
	r.Options = opts
	return
}

// Aggregate threat indicators by one or more columns (e.g., indicatorType, value)
// across datasets. Returns top-N groups ordered by count.
func (r *ThreatEventIndicatorAggregateService) List(ctx context.Context, params ThreatEventIndicatorAggregateListParams, opts ...option.RequestOption) (res *ThreatEventIndicatorAggregateListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/indicators/aggregate", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type ThreatEventIndicatorAggregateListResponse struct {
	// Column(s) that were aggregated by
	AggregateBy string `json:"aggregateBy" api:"required"`
	// Array of aggregation results with dynamic fields based on aggregateBy columns
	Aggregations []ThreatEventIndicatorAggregateListResponseAggregation `json:"aggregations" api:"required"`
	// Number of datasets whose aggregation failed and were excluded from the result
	FailedDatasets float64 `json:"failedDatasets" api:"required"`
	// Total count in the aggregation: indicator rows when measure=indicators, or
	// linked-event rows when measure=relationships
	Total float64                                       `json:"total" api:"required"`
	JSON  threatEventIndicatorAggregateListResponseJSON `json:"-"`
}

// threatEventIndicatorAggregateListResponseJSON contains the JSON metadata for the
// struct [ThreatEventIndicatorAggregateListResponse]
type threatEventIndicatorAggregateListResponseJSON struct {
	AggregateBy    apijson.Field
	Aggregations   apijson.Field
	FailedDatasets apijson.Field
	Total          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ThreatEventIndicatorAggregateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorAggregateListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorAggregateListResponseAggregation struct {
	// Number of indicators for this aggregation
	Count       float64                                                  `json:"count" api:"required"`
	ExtraFields map[string]string                                        `json:"-" api:"extrafields"`
	JSON        threatEventIndicatorAggregateListResponseAggregationJSON `json:"-"`
}

// threatEventIndicatorAggregateListResponseAggregationJSON contains the JSON
// metadata for the struct [ThreatEventIndicatorAggregateListResponseAggregation]
type threatEventIndicatorAggregateListResponseAggregationJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventIndicatorAggregateListResponseAggregation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorAggregateListResponseAggregationJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorAggregateListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Column(s) to aggregate by - single column or comma-separated list (e.g.,
	// 'indicatorType', 'value', 'indicatorType,value')
	AggregateBy param.Field[string] `query:"aggregateBy" api:"required"`
	// Filter indicators created after this date (ISO 8601 format, e.g., '2024-01-01')
	CreatedAfter param.Field[string] `query:"createdAfter"`
	// Filter indicators created before this date (ISO 8601 format, e.g., '2024-12-31')
	CreatedBefore param.Field[string] `query:"createdBefore"`
	// Dataset ID(s) to filter by. Can be a single dataset ID or comma-separated list.
	// If not provided, aggregates across all accessible datasets
	DatasetIDs param.Field[[]string] `query:"datasetIds"`
	// For measure=relationships: only count event links whose eventDate is on/after
	// this date (ISO 8601). Use to bound 'top indicator' to recent activity.
	EventDateAfter param.Field[string] `query:"eventDateAfter"`
	// For measure=relationships: only count event links whose eventDate is on/before
	// this date (ISO 8601).
	EventDateBefore param.Field[string] `query:"eventDateBefore"`
	// Maximum number of aggregation results to return (1-100)
	Limit param.Field[float64] `query:"limit"`
	// What to count per group: 'indicators' (catalog rows, default) or 'relationships'
	// (linked events per indicator). Use 'relationships' for 'top indicator by event
	// activity'.
	Measure param.Field[ThreatEventIndicatorAggregateListParamsMeasure] `query:"measure"`
	// Scope to indicators associated with this tag/actor UUID. Combine with
	// measure=relationships for 'top indicator for an actor'.
	TagUUID param.Field[string] `query:"tagUuid"`
}

// URLQuery serializes [ThreatEventIndicatorAggregateListParams]'s query parameters
// as `url.Values`.
func (r ThreatEventIndicatorAggregateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// What to count per group: 'indicators' (catalog rows, default) or 'relationships'
// (linked events per indicator). Use 'relationships' for 'top indicator by event
// activity'.
type ThreatEventIndicatorAggregateListParamsMeasure string

const (
	ThreatEventIndicatorAggregateListParamsMeasureIndicators    ThreatEventIndicatorAggregateListParamsMeasure = "indicators"
	ThreatEventIndicatorAggregateListParamsMeasureRelationships ThreatEventIndicatorAggregateListParamsMeasure = "relationships"
)

func (r ThreatEventIndicatorAggregateListParamsMeasure) IsKnown() bool {
	switch r {
	case ThreatEventIndicatorAggregateListParamsMeasureIndicators, ThreatEventIndicatorAggregateListParamsMeasureRelationships:
		return true
	}
	return false
}
