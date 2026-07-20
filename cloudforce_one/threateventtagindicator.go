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

// ThreatEventTagIndicatorService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventTagIndicatorService] method instead.
type ThreatEventTagIndicatorService struct {
	Options   []option.RequestOption
	ByDataset *ThreatEventTagIndicatorByDatasetService
}

// NewThreatEventTagIndicatorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatEventTagIndicatorService(opts ...option.RequestOption) (r *ThreatEventTagIndicatorService) {
	r = &ThreatEventTagIndicatorService{}
	r.Options = opts
	r.ByDataset = NewThreatEventTagIndicatorByDatasetService(opts...)
	return
}

// Returns indicators associated with the provided tag UUID, with pagination. By
// default fans out across every indicator dataset the account can read; pass
// datasetIds to scope to specific datasets.
func (r *ThreatEventTagIndicatorService) List(ctx context.Context, tagUUID string, params ThreatEventTagIndicatorListParams, opts ...option.RequestOption) (res *ThreatEventTagIndicatorListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if tagUUID == "" {
		err = errors.New("missing required tag_uuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/tags/%s/indicators", params.AccountID, tagUUID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type ThreatEventTagIndicatorListResponse struct {
	Indicators []ThreatEventTagIndicatorListResponseIndicator `json:"indicators" api:"required"`
	Pagination ThreatEventTagIndicatorListResponsePagination  `json:"pagination" api:"required"`
	JSON       threatEventTagIndicatorListResponseJSON        `json:"-"`
}

// threatEventTagIndicatorListResponseJSON contains the JSON metadata for the
// struct [ThreatEventTagIndicatorListResponse]
type threatEventTagIndicatorListResponseJSON struct {
	Indicators  apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagIndicatorListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagIndicatorListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagIndicatorListResponseIndicator struct {
	CreatedAt     time.Time `json:"createdAt" api:"required" format:"date-time"`
	IndicatorType string    `json:"indicatorType" api:"required"`
	UpdatedAt     time.Time `json:"updatedAt" api:"required" format:"date-time"`
	UUID          string    `json:"uuid" api:"required"`
	Value         string    `json:"value" api:"required"`
	// The dataset ID this indicator belongs to. Included in list responses.
	DatasetID     string                                                      `json:"datasetId"`
	RelatedEvents []ThreatEventTagIndicatorListResponseIndicatorsRelatedEvent `json:"relatedEvents"`
	Tags          []ThreatEventTagIndicatorListResponseIndicatorsTag          `json:"tags"`
	JSON          threatEventTagIndicatorListResponseIndicatorJSON            `json:"-"`
}

// threatEventTagIndicatorListResponseIndicatorJSON contains the JSON metadata for
// the struct [ThreatEventTagIndicatorListResponseIndicator]
type threatEventTagIndicatorListResponseIndicatorJSON struct {
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

func (r *ThreatEventTagIndicatorListResponseIndicator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagIndicatorListResponseIndicatorJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagIndicatorListResponseIndicatorsRelatedEvent struct {
	DatasetID string                                                        `json:"datasetId" api:"required"`
	EventID   string                                                        `json:"eventId" api:"required"`
	JSON      threatEventTagIndicatorListResponseIndicatorsRelatedEventJSON `json:"-"`
}

// threatEventTagIndicatorListResponseIndicatorsRelatedEventJSON contains the JSON
// metadata for the struct
// [ThreatEventTagIndicatorListResponseIndicatorsRelatedEvent]
type threatEventTagIndicatorListResponseIndicatorsRelatedEventJSON struct {
	DatasetID   apijson.Field
	EventID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagIndicatorListResponseIndicatorsRelatedEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagIndicatorListResponseIndicatorsRelatedEventJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagIndicatorListResponseIndicatorsTag struct {
	CategoryName string                                               `json:"categoryName"`
	UUID         string                                               `json:"uuid"`
	Value        string                                               `json:"value"`
	JSON         threatEventTagIndicatorListResponseIndicatorsTagJSON `json:"-"`
}

// threatEventTagIndicatorListResponseIndicatorsTagJSON contains the JSON metadata
// for the struct [ThreatEventTagIndicatorListResponseIndicatorsTag]
type threatEventTagIndicatorListResponseIndicatorsTagJSON struct {
	CategoryName apijson.Field
	UUID         apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ThreatEventTagIndicatorListResponseIndicatorsTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagIndicatorListResponseIndicatorsTagJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagIndicatorListResponsePagination struct {
	Page       float64                                           `json:"page" api:"required"`
	PageSize   float64                                           `json:"pageSize" api:"required"`
	TotalCount float64                                           `json:"totalCount" api:"required"`
	TotalPages float64                                           `json:"totalPages" api:"required"`
	JSON       threatEventTagIndicatorListResponsePaginationJSON `json:"-"`
}

// threatEventTagIndicatorListResponsePaginationJSON contains the JSON metadata for
// the struct [ThreatEventTagIndicatorListResponsePagination]
type threatEventTagIndicatorListResponsePaginationJSON struct {
	Page        apijson.Field
	PageSize    apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagIndicatorListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagIndicatorListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagIndicatorListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Dataset UUIDs to scope to (repeat the param for multiple), or 'all' / '\*' for
	// every readable indicator dataset. Omit to search all readable datasets.
	DatasetIDs    param.Field[[]string] `query:"datasetIds"`
	IndicatorType param.Field[string]   `query:"indicatorType"`
	Page          param.Field[float64]  `query:"page"`
	PageSize      param.Field[float64]  `query:"pageSize"`
	// Filter indicators by related event UUID(s). Multiple UUIDs can be provided by
	// repeating the parameter.
	RelatedEvent param.Field[[]string] `query:"relatedEvent"`
	// Structured search as a JSON array of {field, op, value} objects. Searchable
	// fields: value, indicatorType. Multiple conditions are AND'd together. Max 10
	// conditions per request.
	Search param.Field[[]ThreatEventTagIndicatorListParamsSearch] `query:"search"`
}

// URLQuery serializes [ThreatEventTagIndicatorListParams]'s query parameters as
// `url.Values`.
func (r ThreatEventTagIndicatorListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ThreatEventTagIndicatorListParamsSearch struct {
	// The indicator field to search on. Allowed: value, indicatorType.
	Field param.Field[ThreatEventTagIndicatorListParamsSearchField] `query:"field" api:"required"`
	// Search operator. Use 'in' for bulk lookup of up to 100 values at once, e.g.
	// {field:'value', op:'in', value:['evil.com','bad.org']}.
	Op param.Field[ThreatEventTagIndicatorListParamsSearchOp] `query:"op" api:"required"`
	// Search value. String for most operators. Array of strings for 'in' operator (max
	// 100 items).
	Value param.Field[ThreatEventTagIndicatorListParamsSearchValueUnion] `query:"value" api:"required"`
}

// URLQuery serializes [ThreatEventTagIndicatorListParamsSearch]'s query parameters
// as `url.Values`.
func (r ThreatEventTagIndicatorListParamsSearch) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The indicator field to search on. Allowed: value, indicatorType.
type ThreatEventTagIndicatorListParamsSearchField string

const (
	ThreatEventTagIndicatorListParamsSearchFieldValue         ThreatEventTagIndicatorListParamsSearchField = "value"
	ThreatEventTagIndicatorListParamsSearchFieldIndicatorType ThreatEventTagIndicatorListParamsSearchField = "indicatorType"
)

func (r ThreatEventTagIndicatorListParamsSearchField) IsKnown() bool {
	switch r {
	case ThreatEventTagIndicatorListParamsSearchFieldValue, ThreatEventTagIndicatorListParamsSearchFieldIndicatorType:
		return true
	}
	return false
}

// Search operator. Use 'in' for bulk lookup of up to 100 values at once, e.g.
// {field:'value', op:'in', value:['evil.com','bad.org']}.
type ThreatEventTagIndicatorListParamsSearchOp string

const (
	ThreatEventTagIndicatorListParamsSearchOpEquals     ThreatEventTagIndicatorListParamsSearchOp = "equals"
	ThreatEventTagIndicatorListParamsSearchOpNot        ThreatEventTagIndicatorListParamsSearchOp = "not"
	ThreatEventTagIndicatorListParamsSearchOpGt         ThreatEventTagIndicatorListParamsSearchOp = "gt"
	ThreatEventTagIndicatorListParamsSearchOpGte        ThreatEventTagIndicatorListParamsSearchOp = "gte"
	ThreatEventTagIndicatorListParamsSearchOpLt         ThreatEventTagIndicatorListParamsSearchOp = "lt"
	ThreatEventTagIndicatorListParamsSearchOpLte        ThreatEventTagIndicatorListParamsSearchOp = "lte"
	ThreatEventTagIndicatorListParamsSearchOpLike       ThreatEventTagIndicatorListParamsSearchOp = "like"
	ThreatEventTagIndicatorListParamsSearchOpContains   ThreatEventTagIndicatorListParamsSearchOp = "contains"
	ThreatEventTagIndicatorListParamsSearchOpStartsWith ThreatEventTagIndicatorListParamsSearchOp = "startsWith"
	ThreatEventTagIndicatorListParamsSearchOpEndsWith   ThreatEventTagIndicatorListParamsSearchOp = "endsWith"
	ThreatEventTagIndicatorListParamsSearchOpIn         ThreatEventTagIndicatorListParamsSearchOp = "in"
	ThreatEventTagIndicatorListParamsSearchOpFind       ThreatEventTagIndicatorListParamsSearchOp = "find"
)

func (r ThreatEventTagIndicatorListParamsSearchOp) IsKnown() bool {
	switch r {
	case ThreatEventTagIndicatorListParamsSearchOpEquals, ThreatEventTagIndicatorListParamsSearchOpNot, ThreatEventTagIndicatorListParamsSearchOpGt, ThreatEventTagIndicatorListParamsSearchOpGte, ThreatEventTagIndicatorListParamsSearchOpLt, ThreatEventTagIndicatorListParamsSearchOpLte, ThreatEventTagIndicatorListParamsSearchOpLike, ThreatEventTagIndicatorListParamsSearchOpContains, ThreatEventTagIndicatorListParamsSearchOpStartsWith, ThreatEventTagIndicatorListParamsSearchOpEndsWith, ThreatEventTagIndicatorListParamsSearchOpIn, ThreatEventTagIndicatorListParamsSearchOpFind:
		return true
	}
	return false
}

// Search value. String for most operators. Array of strings for 'in' operator (max
// 100 items).
//
// Satisfied by [shared.UnionString],
// [cloudforce_one.ThreatEventTagIndicatorListParamsSearchValueArray].
type ThreatEventTagIndicatorListParamsSearchValueUnion interface {
	ImplementsThreatEventTagIndicatorListParamsSearchValueUnion()
}

type ThreatEventTagIndicatorListParamsSearchValueArray []string

func (r ThreatEventTagIndicatorListParamsSearchValueArray) ImplementsThreatEventTagIndicatorListParamsSearchValueUnion() {
}
