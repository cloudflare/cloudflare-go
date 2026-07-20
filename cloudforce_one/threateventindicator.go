// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// ThreatEventIndicatorService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventIndicatorService] method instead.
type ThreatEventIndicatorService struct {
	Options   []option.RequestOption
	Aggregate *ThreatEventIndicatorAggregateService
	Types     *ThreatEventIndicatorTypeService
	ByDataset *ThreatEventIndicatorByDatasetService
}

// NewThreatEventIndicatorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatEventIndicatorService(opts ...option.RequestOption) (r *ThreatEventIndicatorService) {
	r = &ThreatEventIndicatorService{}
	r.Options = opts
	r.Aggregate = NewThreatEventIndicatorAggregateService(opts...)
	r.Types = NewThreatEventIndicatorTypeService(opts...)
	r.ByDataset = NewThreatEventIndicatorByDatasetService(opts...)
	return
}

// Retrieves a paginated list of indicators across specified datasets. Use
// datasetIds=all or datasetIds=\* to query all datasets for the account. If no
// datasetIds provided, uses the default dataset.
func (r *ThreatEventIndicatorService) List(ctx context.Context, params ThreatEventIndicatorListParams, opts ...option.RequestOption) (res *ThreatEventIndicatorListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/indicators", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type ThreatEventIndicatorListResponse struct {
	Properties ThreatEventIndicatorListResponseProperties `json:"properties" api:"required"`
	Type       string                                     `json:"type" api:"required"`
	JSON       threatEventIndicatorListResponseJSON       `json:"-"`
}

// threatEventIndicatorListResponseJSON contains the JSON metadata for the struct
// [ThreatEventIndicatorListResponse]
type threatEventIndicatorListResponseJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventIndicatorListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorListResponseProperties struct {
	Indicators ThreatEventIndicatorListResponsePropertiesIndicators `json:"indicators" api:"required"`
	Pagination ThreatEventIndicatorListResponsePropertiesPagination `json:"pagination" api:"required"`
	JSON       threatEventIndicatorListResponsePropertiesJSON       `json:"-"`
}

// threatEventIndicatorListResponsePropertiesJSON contains the JSON metadata for
// the struct [ThreatEventIndicatorListResponseProperties]
type threatEventIndicatorListResponsePropertiesJSON struct {
	Indicators  apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventIndicatorListResponseProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorListResponsePropertiesJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorListResponsePropertiesIndicators struct {
	Items ThreatEventIndicatorListResponsePropertiesIndicatorsItems `json:"items" api:"required"`
	Type  string                                                    `json:"type" api:"required"`
	JSON  threatEventIndicatorListResponsePropertiesIndicatorsJSON  `json:"-"`
}

// threatEventIndicatorListResponsePropertiesIndicatorsJSON contains the JSON
// metadata for the struct [ThreatEventIndicatorListResponsePropertiesIndicators]
type threatEventIndicatorListResponsePropertiesIndicatorsJSON struct {
	Items       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventIndicatorListResponsePropertiesIndicators) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorListResponsePropertiesIndicatorsJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorListResponsePropertiesIndicatorsItems struct {
	CreatedAt     time.Time `json:"createdAt" api:"required" format:"date-time"`
	IndicatorType string    `json:"indicatorType" api:"required"`
	UpdatedAt     time.Time `json:"updatedAt" api:"required" format:"date-time"`
	UUID          string    `json:"uuid" api:"required"`
	Value         string    `json:"value" api:"required"`
	// The dataset ID this indicator belongs to. Included in list responses.
	DatasetID     string                                                                  `json:"datasetId"`
	RelatedEvents []ThreatEventIndicatorListResponsePropertiesIndicatorsItemsRelatedEvent `json:"relatedEvents"`
	Tags          []ThreatEventIndicatorListResponsePropertiesIndicatorsItemsTag          `json:"tags"`
	JSON          threatEventIndicatorListResponsePropertiesIndicatorsItemsJSON           `json:"-"`
}

// threatEventIndicatorListResponsePropertiesIndicatorsItemsJSON contains the JSON
// metadata for the struct
// [ThreatEventIndicatorListResponsePropertiesIndicatorsItems]
type threatEventIndicatorListResponsePropertiesIndicatorsItemsJSON struct {
	CreatedAt     apijson.Field
	IndicatorType apijson.Field
	UpdatedAt     apijson.Field
	UUID          apijson.Field
	Value         apijson.Field
	DatasetID     apijson.Field
	RelatedEvents apijson.Field
	Tags          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ThreatEventIndicatorListResponsePropertiesIndicatorsItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorListResponsePropertiesIndicatorsItemsJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorListResponsePropertiesIndicatorsItemsRelatedEvent struct {
	DatasetID string                                                                    `json:"datasetId" api:"required"`
	EventID   string                                                                    `json:"eventId" api:"required"`
	JSON      threatEventIndicatorListResponsePropertiesIndicatorsItemsRelatedEventJSON `json:"-"`
}

// threatEventIndicatorListResponsePropertiesIndicatorsItemsRelatedEventJSON
// contains the JSON metadata for the struct
// [ThreatEventIndicatorListResponsePropertiesIndicatorsItemsRelatedEvent]
type threatEventIndicatorListResponsePropertiesIndicatorsItemsRelatedEventJSON struct {
	DatasetID   apijson.Field
	EventID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventIndicatorListResponsePropertiesIndicatorsItemsRelatedEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorListResponsePropertiesIndicatorsItemsRelatedEventJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorListResponsePropertiesIndicatorsItemsTag struct {
	CategoryName string                                                           `json:"categoryName"`
	UUID         string                                                           `json:"uuid"`
	Value        string                                                           `json:"value"`
	JSON         threatEventIndicatorListResponsePropertiesIndicatorsItemsTagJSON `json:"-"`
}

// threatEventIndicatorListResponsePropertiesIndicatorsItemsTagJSON contains the
// JSON metadata for the struct
// [ThreatEventIndicatorListResponsePropertiesIndicatorsItemsTag]
type threatEventIndicatorListResponsePropertiesIndicatorsItemsTagJSON struct {
	CategoryName apijson.Field
	UUID         apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ThreatEventIndicatorListResponsePropertiesIndicatorsItemsTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorListResponsePropertiesIndicatorsItemsTagJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorListResponsePropertiesPagination struct {
	Properties ThreatEventIndicatorListResponsePropertiesPaginationProperties `json:"properties" api:"required"`
	Type       string                                                         `json:"type" api:"required"`
	JSON       threatEventIndicatorListResponsePropertiesPaginationJSON       `json:"-"`
}

// threatEventIndicatorListResponsePropertiesPaginationJSON contains the JSON
// metadata for the struct [ThreatEventIndicatorListResponsePropertiesPagination]
type threatEventIndicatorListResponsePropertiesPaginationJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventIndicatorListResponsePropertiesPagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorListResponsePropertiesPaginationJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorListResponsePropertiesPaginationProperties struct {
	Count      ThreatEventIndicatorListResponsePropertiesPaginationPropertiesCount      `json:"count" api:"required"`
	Page       ThreatEventIndicatorListResponsePropertiesPaginationPropertiesPage       `json:"page" api:"required"`
	PerPage    ThreatEventIndicatorListResponsePropertiesPaginationPropertiesPerPage    `json:"per_page" api:"required"`
	TotalCount ThreatEventIndicatorListResponsePropertiesPaginationPropertiesTotalCount `json:"total_count" api:"required"`
	JSON       threatEventIndicatorListResponsePropertiesPaginationPropertiesJSON       `json:"-"`
}

// threatEventIndicatorListResponsePropertiesPaginationPropertiesJSON contains the
// JSON metadata for the struct
// [ThreatEventIndicatorListResponsePropertiesPaginationProperties]
type threatEventIndicatorListResponsePropertiesPaginationPropertiesJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventIndicatorListResponsePropertiesPaginationProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorListResponsePropertiesPaginationPropertiesJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorListResponsePropertiesPaginationPropertiesCount struct {
	Type string                                                                  `json:"type" api:"required"`
	JSON threatEventIndicatorListResponsePropertiesPaginationPropertiesCountJSON `json:"-"`
}

// threatEventIndicatorListResponsePropertiesPaginationPropertiesCountJSON contains
// the JSON metadata for the struct
// [ThreatEventIndicatorListResponsePropertiesPaginationPropertiesCount]
type threatEventIndicatorListResponsePropertiesPaginationPropertiesCountJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventIndicatorListResponsePropertiesPaginationPropertiesCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorListResponsePropertiesPaginationPropertiesCountJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorListResponsePropertiesPaginationPropertiesPage struct {
	Type string                                                                 `json:"type" api:"required"`
	JSON threatEventIndicatorListResponsePropertiesPaginationPropertiesPageJSON `json:"-"`
}

// threatEventIndicatorListResponsePropertiesPaginationPropertiesPageJSON contains
// the JSON metadata for the struct
// [ThreatEventIndicatorListResponsePropertiesPaginationPropertiesPage]
type threatEventIndicatorListResponsePropertiesPaginationPropertiesPageJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventIndicatorListResponsePropertiesPaginationPropertiesPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorListResponsePropertiesPaginationPropertiesPageJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorListResponsePropertiesPaginationPropertiesPerPage struct {
	Type string                                                                    `json:"type" api:"required"`
	JSON threatEventIndicatorListResponsePropertiesPaginationPropertiesPerPageJSON `json:"-"`
}

// threatEventIndicatorListResponsePropertiesPaginationPropertiesPerPageJSON
// contains the JSON metadata for the struct
// [ThreatEventIndicatorListResponsePropertiesPaginationPropertiesPerPage]
type threatEventIndicatorListResponsePropertiesPaginationPropertiesPerPageJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventIndicatorListResponsePropertiesPaginationPropertiesPerPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorListResponsePropertiesPaginationPropertiesPerPageJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorListResponsePropertiesPaginationPropertiesTotalCount struct {
	Type string                                                                       `json:"type" api:"required"`
	JSON threatEventIndicatorListResponsePropertiesPaginationPropertiesTotalCountJSON `json:"-"`
}

// threatEventIndicatorListResponsePropertiesPaginationPropertiesTotalCountJSON
// contains the JSON metadata for the struct
// [ThreatEventIndicatorListResponsePropertiesPaginationPropertiesTotalCount]
type threatEventIndicatorListResponsePropertiesPaginationPropertiesTotalCountJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventIndicatorListResponsePropertiesPaginationPropertiesTotalCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorListResponsePropertiesPaginationPropertiesTotalCountJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter indicators created on or after this date. Must use ISO 8601 format (e.g.,
	// '2024-01-15T00:00:00Z').
	CreatedAfter param.Field[time.Time] `query:"createdAfter" format:"date-time"`
	// Filter indicators created on or before this date. Must use ISO 8601 format
	// (e.g., '2024-12-31T23:59:59Z').
	CreatedBefore param.Field[time.Time] `query:"createdBefore" format:"date-time"`
	// Dataset IDs to query indicators from (array of UUIDs), or special value 'all' or
	// '\*' to query all datasets. If not provided, uses the default dataset.
	DatasetIDs param.Field[[]string] `query:"datasetIds"`
	// Output format for indicator data. 'json' returns the default format, 'stix2'
	// returns STIX 2.1 Indicator SDOs, 'taxii' returns a TAXII 2.1 Envelope with
	// Content-Type application/taxii+json;version=2.1.
	Format param.Field[ThreatEventIndicatorListParamsFormat] `query:"format"`
	// Whether to include full tag details for each indicator. Defaults to true.
	IncludeTags param.Field[bool] `query:"includeTags"`
	// Whether to compute accurate total count via COUNT(\*). Defaults to false for
	// performance. When false, total_count is an approximation.
	IncludeTotalCount param.Field[bool]   `query:"includeTotalCount"`
	IndicatorType     param.Field[string] `query:"indicatorType"`
	// Filter indicators by value using substring match (LIKE). Legacy alternative to
	// structured search.
	Name     param.Field[string]  `query:"name"`
	Page     param.Field[float64] `query:"page"`
	PageSize param.Field[float64] `query:"pageSize"`
	// Filter by related event IDs
	RelatedEvents param.Field[[]string] `query:"relatedEvents"`
	// Limit the number of related events returned per indicator. Default: 2. Set to 0
	// for none, -1 for all events.
	RelatedEventsLimit param.Field[float64] `query:"relatedEventsLimit"`
	// Structured search as a JSON array of {field, op, value} objects. Searchable
	// fields: value, indicatorType. Supports operators: equals, not, contains,
	// startsWith, endsWith, gt, lt, gte, lte, like, in, find. Use the 'in' operator
	// with an array value to bulk-check up to 100 indicators in a single request, e.g.
	// search=[{"field":"value","op":"in","value":["evil.com","bad.org"]}]. Multiple
	// conditions are AND'd together. Max 10 conditions per request.
	Search param.Field[[]ThreatEventIndicatorListParamsSearch] `query:"search"`
	// Read backend. 'do' (default) reads Durable Object storage. 'r2catalog' reads R2
	// Data Catalog (admin-only, experimental; supports a subset of search fields).
	Source param.Field[ThreatEventIndicatorListParamsSource] `query:"source"`
	// Filter by tag values or UUIDs. Indicators must have at least one of the
	// specified tags (OR logic). Supports both tag UUID and tag value.
	Tags param.Field[[]string] `query:"tags"`
	// Structured tag-metadata filter as a JSON array of {field, op, value} objects.
	// Operates against the per-dataset IndicatorTag mirror so you can find indicators
	// by tag attributes (origin country, motive, sophistication, priority, etc.)
	// without a separate Tags lookup. Common dashboard usage: drill from a country
	// into indicators, e.g.
	// tagSearch=[{"field":"originCountryISO","op":"in","value":["IR","CN"]}]. Country
	// values may be passed as alpha-2, alpha-3, name, or alias (e.g. "iran").
	// Operators: equals, not, gt/gte/lt/lte (numeric only),
	// contains/like/find/startsWith/endsWith (string only), in. AND-joined across
	// entries; combined with `tags`, a matching tag must satisfy both. Max 10 entries
	// per request, max 100 values per 'in'. Performance notes: `originCountryISO` uses
	// its B-tree index for equals/not/in. `priority` uses its B-tree index for numeric
	// comparisons. Other string columns (`actorCategory`, `motive`, etc.) are
	// case-insensitive and unindexed; current catalog size makes this a non-issue.
	// `endsWith` and `aliasGroupNames` contains/like are leading-wildcard scans and
	// slow on large result sets. `aliasGroupNames` matches on the JSON-encoded text,
	// so substrings can cross alias boundaries ("apt28" also matches "apt280" when
	// both appear in the same tag's alias list).
	TagSearch param.Field[[]ThreatEventIndicatorListParamsTagSearch] `query:"tagSearch"`
}

// URLQuery serializes [ThreatEventIndicatorListParams]'s query parameters as
// `url.Values`.
func (r ThreatEventIndicatorListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Output format for indicator data. 'json' returns the default format, 'stix2'
// returns STIX 2.1 Indicator SDOs, 'taxii' returns a TAXII 2.1 Envelope with
// Content-Type application/taxii+json;version=2.1.
type ThreatEventIndicatorListParamsFormat string

const (
	ThreatEventIndicatorListParamsFormatJson  ThreatEventIndicatorListParamsFormat = "json"
	ThreatEventIndicatorListParamsFormatStix2 ThreatEventIndicatorListParamsFormat = "stix2"
	ThreatEventIndicatorListParamsFormatTaxii ThreatEventIndicatorListParamsFormat = "taxii"
)

func (r ThreatEventIndicatorListParamsFormat) IsKnown() bool {
	switch r {
	case ThreatEventIndicatorListParamsFormatJson, ThreatEventIndicatorListParamsFormatStix2, ThreatEventIndicatorListParamsFormatTaxii:
		return true
	}
	return false
}

type ThreatEventIndicatorListParamsSearch struct {
	// The indicator field to search on. Allowed: value, indicatorType.
	Field param.Field[ThreatEventIndicatorListParamsSearchField] `query:"field" api:"required"`
	// Search operator. Use 'in' for bulk lookup of up to 100 values at once, e.g.
	// {field:'value', op:'in', value:['evil.com','bad.org']}.
	Op param.Field[ThreatEventIndicatorListParamsSearchOp] `query:"op" api:"required"`
	// Search value. String for most operators. Array of strings for 'in' operator (max
	// 100 items).
	Value param.Field[ThreatEventIndicatorListParamsSearchValueUnion] `query:"value" api:"required"`
}

// URLQuery serializes [ThreatEventIndicatorListParamsSearch]'s query parameters as
// `url.Values`.
func (r ThreatEventIndicatorListParamsSearch) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The indicator field to search on. Allowed: value, indicatorType.
type ThreatEventIndicatorListParamsSearchField string

const (
	ThreatEventIndicatorListParamsSearchFieldValue         ThreatEventIndicatorListParamsSearchField = "value"
	ThreatEventIndicatorListParamsSearchFieldIndicatorType ThreatEventIndicatorListParamsSearchField = "indicatorType"
)

func (r ThreatEventIndicatorListParamsSearchField) IsKnown() bool {
	switch r {
	case ThreatEventIndicatorListParamsSearchFieldValue, ThreatEventIndicatorListParamsSearchFieldIndicatorType:
		return true
	}
	return false
}

// Search operator. Use 'in' for bulk lookup of up to 100 values at once, e.g.
// {field:'value', op:'in', value:['evil.com','bad.org']}.
type ThreatEventIndicatorListParamsSearchOp string

const (
	ThreatEventIndicatorListParamsSearchOpEquals     ThreatEventIndicatorListParamsSearchOp = "equals"
	ThreatEventIndicatorListParamsSearchOpNot        ThreatEventIndicatorListParamsSearchOp = "not"
	ThreatEventIndicatorListParamsSearchOpGt         ThreatEventIndicatorListParamsSearchOp = "gt"
	ThreatEventIndicatorListParamsSearchOpGte        ThreatEventIndicatorListParamsSearchOp = "gte"
	ThreatEventIndicatorListParamsSearchOpLt         ThreatEventIndicatorListParamsSearchOp = "lt"
	ThreatEventIndicatorListParamsSearchOpLte        ThreatEventIndicatorListParamsSearchOp = "lte"
	ThreatEventIndicatorListParamsSearchOpLike       ThreatEventIndicatorListParamsSearchOp = "like"
	ThreatEventIndicatorListParamsSearchOpContains   ThreatEventIndicatorListParamsSearchOp = "contains"
	ThreatEventIndicatorListParamsSearchOpStartsWith ThreatEventIndicatorListParamsSearchOp = "startsWith"
	ThreatEventIndicatorListParamsSearchOpEndsWith   ThreatEventIndicatorListParamsSearchOp = "endsWith"
	ThreatEventIndicatorListParamsSearchOpIn         ThreatEventIndicatorListParamsSearchOp = "in"
	ThreatEventIndicatorListParamsSearchOpFind       ThreatEventIndicatorListParamsSearchOp = "find"
)

func (r ThreatEventIndicatorListParamsSearchOp) IsKnown() bool {
	switch r {
	case ThreatEventIndicatorListParamsSearchOpEquals, ThreatEventIndicatorListParamsSearchOpNot, ThreatEventIndicatorListParamsSearchOpGt, ThreatEventIndicatorListParamsSearchOpGte, ThreatEventIndicatorListParamsSearchOpLt, ThreatEventIndicatorListParamsSearchOpLte, ThreatEventIndicatorListParamsSearchOpLike, ThreatEventIndicatorListParamsSearchOpContains, ThreatEventIndicatorListParamsSearchOpStartsWith, ThreatEventIndicatorListParamsSearchOpEndsWith, ThreatEventIndicatorListParamsSearchOpIn, ThreatEventIndicatorListParamsSearchOpFind:
		return true
	}
	return false
}

// Search value. String for most operators. Array of strings for 'in' operator (max
// 100 items).
//
// Satisfied by [shared.UnionString],
// [cloudforce_one.ThreatEventIndicatorListParamsSearchValueArray].
type ThreatEventIndicatorListParamsSearchValueUnion interface {
	ImplementsThreatEventIndicatorListParamsSearchValueUnion()
}

type ThreatEventIndicatorListParamsSearchValueArray []string

func (r ThreatEventIndicatorListParamsSearchValueArray) ImplementsThreatEventIndicatorListParamsSearchValueUnion() {
}

// Read backend. 'do' (default) reads Durable Object storage. 'r2catalog' reads R2
// Data Catalog (admin-only, experimental; supports a subset of search fields).
type ThreatEventIndicatorListParamsSource string

const (
	ThreatEventIndicatorListParamsSourceDo        ThreatEventIndicatorListParamsSource = "do"
	ThreatEventIndicatorListParamsSourceR2catalog ThreatEventIndicatorListParamsSource = "r2catalog"
)

func (r ThreatEventIndicatorListParamsSource) IsKnown() bool {
	switch r {
	case ThreatEventIndicatorListParamsSourceDo, ThreatEventIndicatorListParamsSourceR2catalog:
		return true
	}
	return false
}

type ThreatEventIndicatorListParamsTagSearch struct {
	// Tag mirror field to filter on. Allowed: value, categoryId, actorCategory,
	// aliasGroupNames, attributionConfidence, attributionOrganization, motive,
	// opsecLevel, originCountryISO, sophisticationLevel, priority, analyticPriority.
	// Filters operate against the per-dataset IndicatorTag mirror (which is kept in
	// sync with the Tags SoT by the tag-propagation workflow).
	Field param.Field[ThreatEventIndicatorListParamsTagSearchField] `query:"field" api:"required"`
	// Search operator. Use 'in' for bulk OR within a single field, e.g.
	// {field:"originCountryISO", op:"in", value:["IR","CN"]}.
	Op param.Field[ThreatEventIndicatorListParamsTagSearchOp] `query:"op" api:"required"`
	// Search value. String or number for most operators. Array for 'in' (max 100
	// items). Country values may be passed as alpha-2, alpha-3, name, or common alias
	// (e.g. "iran", "IR", "IRN") and are normalized to alpha-2 server-side.
	Value param.Field[ThreatEventIndicatorListParamsTagSearchValueUnion] `query:"value"`
}

// URLQuery serializes [ThreatEventIndicatorListParamsTagSearch]'s query parameters
// as `url.Values`.
func (r ThreatEventIndicatorListParamsTagSearch) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Tag mirror field to filter on. Allowed: value, categoryId, actorCategory,
// aliasGroupNames, attributionConfidence, attributionOrganization, motive,
// opsecLevel, originCountryISO, sophisticationLevel, priority, analyticPriority.
// Filters operate against the per-dataset IndicatorTag mirror (which is kept in
// sync with the Tags SoT by the tag-propagation workflow).
type ThreatEventIndicatorListParamsTagSearchField string

const (
	ThreatEventIndicatorListParamsTagSearchFieldValue                   ThreatEventIndicatorListParamsTagSearchField = "value"
	ThreatEventIndicatorListParamsTagSearchFieldCategoryID              ThreatEventIndicatorListParamsTagSearchField = "categoryId"
	ThreatEventIndicatorListParamsTagSearchFieldActorCategory           ThreatEventIndicatorListParamsTagSearchField = "actorCategory"
	ThreatEventIndicatorListParamsTagSearchFieldAliasGroupNames         ThreatEventIndicatorListParamsTagSearchField = "aliasGroupNames"
	ThreatEventIndicatorListParamsTagSearchFieldAttributionConfidence   ThreatEventIndicatorListParamsTagSearchField = "attributionConfidence"
	ThreatEventIndicatorListParamsTagSearchFieldAttributionOrganization ThreatEventIndicatorListParamsTagSearchField = "attributionOrganization"
	ThreatEventIndicatorListParamsTagSearchFieldMotive                  ThreatEventIndicatorListParamsTagSearchField = "motive"
	ThreatEventIndicatorListParamsTagSearchFieldOpsecLevel              ThreatEventIndicatorListParamsTagSearchField = "opsecLevel"
	ThreatEventIndicatorListParamsTagSearchFieldOriginCountryISO        ThreatEventIndicatorListParamsTagSearchField = "originCountryISO"
	ThreatEventIndicatorListParamsTagSearchFieldSophisticationLevel     ThreatEventIndicatorListParamsTagSearchField = "sophisticationLevel"
	ThreatEventIndicatorListParamsTagSearchFieldPriority                ThreatEventIndicatorListParamsTagSearchField = "priority"
	ThreatEventIndicatorListParamsTagSearchFieldAnalyticPriority        ThreatEventIndicatorListParamsTagSearchField = "analyticPriority"
)

func (r ThreatEventIndicatorListParamsTagSearchField) IsKnown() bool {
	switch r {
	case ThreatEventIndicatorListParamsTagSearchFieldValue, ThreatEventIndicatorListParamsTagSearchFieldCategoryID, ThreatEventIndicatorListParamsTagSearchFieldActorCategory, ThreatEventIndicatorListParamsTagSearchFieldAliasGroupNames, ThreatEventIndicatorListParamsTagSearchFieldAttributionConfidence, ThreatEventIndicatorListParamsTagSearchFieldAttributionOrganization, ThreatEventIndicatorListParamsTagSearchFieldMotive, ThreatEventIndicatorListParamsTagSearchFieldOpsecLevel, ThreatEventIndicatorListParamsTagSearchFieldOriginCountryISO, ThreatEventIndicatorListParamsTagSearchFieldSophisticationLevel, ThreatEventIndicatorListParamsTagSearchFieldPriority, ThreatEventIndicatorListParamsTagSearchFieldAnalyticPriority:
		return true
	}
	return false
}

// Search operator. Use 'in' for bulk OR within a single field, e.g.
// {field:"originCountryISO", op:"in", value:["IR","CN"]}.
type ThreatEventIndicatorListParamsTagSearchOp string

const (
	ThreatEventIndicatorListParamsTagSearchOpEquals     ThreatEventIndicatorListParamsTagSearchOp = "equals"
	ThreatEventIndicatorListParamsTagSearchOpNot        ThreatEventIndicatorListParamsTagSearchOp = "not"
	ThreatEventIndicatorListParamsTagSearchOpGt         ThreatEventIndicatorListParamsTagSearchOp = "gt"
	ThreatEventIndicatorListParamsTagSearchOpGte        ThreatEventIndicatorListParamsTagSearchOp = "gte"
	ThreatEventIndicatorListParamsTagSearchOpLt         ThreatEventIndicatorListParamsTagSearchOp = "lt"
	ThreatEventIndicatorListParamsTagSearchOpLte        ThreatEventIndicatorListParamsTagSearchOp = "lte"
	ThreatEventIndicatorListParamsTagSearchOpLike       ThreatEventIndicatorListParamsTagSearchOp = "like"
	ThreatEventIndicatorListParamsTagSearchOpContains   ThreatEventIndicatorListParamsTagSearchOp = "contains"
	ThreatEventIndicatorListParamsTagSearchOpStartsWith ThreatEventIndicatorListParamsTagSearchOp = "startsWith"
	ThreatEventIndicatorListParamsTagSearchOpEndsWith   ThreatEventIndicatorListParamsTagSearchOp = "endsWith"
	ThreatEventIndicatorListParamsTagSearchOpIn         ThreatEventIndicatorListParamsTagSearchOp = "in"
	ThreatEventIndicatorListParamsTagSearchOpFind       ThreatEventIndicatorListParamsTagSearchOp = "find"
)

func (r ThreatEventIndicatorListParamsTagSearchOp) IsKnown() bool {
	switch r {
	case ThreatEventIndicatorListParamsTagSearchOpEquals, ThreatEventIndicatorListParamsTagSearchOpNot, ThreatEventIndicatorListParamsTagSearchOpGt, ThreatEventIndicatorListParamsTagSearchOpGte, ThreatEventIndicatorListParamsTagSearchOpLt, ThreatEventIndicatorListParamsTagSearchOpLte, ThreatEventIndicatorListParamsTagSearchOpLike, ThreatEventIndicatorListParamsTagSearchOpContains, ThreatEventIndicatorListParamsTagSearchOpStartsWith, ThreatEventIndicatorListParamsTagSearchOpEndsWith, ThreatEventIndicatorListParamsTagSearchOpIn, ThreatEventIndicatorListParamsTagSearchOpFind:
		return true
	}
	return false
}

// Search value. String or number for most operators. Array for 'in' (max 100
// items). Country values may be passed as alpha-2, alpha-3, name, or common alias
// (e.g. "iran", "IR", "IRN") and are normalized to alpha-2 server-side.
//
// Satisfied by [shared.UnionString], [shared.UnionFloat],
// [cloudforce_one.ThreatEventIndicatorListParamsTagSearchValueArray].
type ThreatEventIndicatorListParamsTagSearchValueUnion interface {
	ImplementsThreatEventIndicatorListParamsTagSearchValueUnion()
}

type ThreatEventIndicatorListParamsTagSearchValueArray []ThreatEventIndicatorListParamsTagSearchValueArrayItemUnion

func (r ThreatEventIndicatorListParamsTagSearchValueArray) ImplementsThreatEventIndicatorListParamsTagSearchValueUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type ThreatEventIndicatorListParamsTagSearchValueArrayItemUnion interface {
	ImplementsThreatEventIndicatorListParamsTagSearchValueArrayItemUnion()
}
