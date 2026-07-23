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

// ThreatEventTagIndicatorByDatasetService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventTagIndicatorByDatasetService] method instead.
type ThreatEventTagIndicatorByDatasetService struct {
	Options []option.RequestOption
}

// NewThreatEventTagIndicatorByDatasetService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewThreatEventTagIndicatorByDatasetService(opts ...option.RequestOption) (r *ThreatEventTagIndicatorByDatasetService) {
	r = &ThreatEventTagIndicatorByDatasetService{}
	r.Options = opts
	return
}

// This endpoint is deprecated. Use GET
// /:account_id/events/tags/:tag_uuid/indicators with the optional datasetIds query
// parameter instead. Returns indicators associated with the provided tag UUID
// within a single dataset's indicator shards, with pagination.
//
// Deprecated: Use list instead (GET
// /accounts/{account_id}/cloudforce-one/events/tags/{tag_uuid}/indicators).
func (r *ThreatEventTagIndicatorByDatasetService) List(ctx context.Context, datasetID string, tagUUID string, params ThreatEventTagIndicatorByDatasetListParams, opts ...option.RequestOption) (res *ThreatEventTagIndicatorByDatasetListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return nil, err
	}
	if tagUUID == "" {
		err = errors.New("missing required tag_uuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/dataset/%s/tags/%s/indicators", params.AccountID, datasetID, tagUUID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type ThreatEventTagIndicatorByDatasetListResponse struct {
	Indicators []ThreatEventTagIndicatorByDatasetListResponseIndicator `json:"indicators" api:"required"`
	Pagination ThreatEventTagIndicatorByDatasetListResponsePagination  `json:"pagination" api:"required"`
	JSON       threatEventTagIndicatorByDatasetListResponseJSON        `json:"-"`
}

// threatEventTagIndicatorByDatasetListResponseJSON contains the JSON metadata for
// the struct [ThreatEventTagIndicatorByDatasetListResponse]
type threatEventTagIndicatorByDatasetListResponseJSON struct {
	Indicators  apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagIndicatorByDatasetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagIndicatorByDatasetListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagIndicatorByDatasetListResponseIndicator struct {
	CreatedAt     time.Time `json:"createdAt" api:"required" format:"date-time"`
	IndicatorType string    `json:"indicatorType" api:"required"`
	UpdatedAt     time.Time `json:"updatedAt" api:"required" format:"date-time"`
	UUID          string    `json:"uuid" api:"required"`
	Value         string    `json:"value" api:"required"`
	// The dataset ID this indicator belongs to. Included in list responses.
	DatasetID     string                                                               `json:"datasetId"`
	RelatedEvents []ThreatEventTagIndicatorByDatasetListResponseIndicatorsRelatedEvent `json:"relatedEvents"`
	Tags          []ThreatEventTagIndicatorByDatasetListResponseIndicatorsTag          `json:"tags"`
	JSON          threatEventTagIndicatorByDatasetListResponseIndicatorJSON            `json:"-"`
}

// threatEventTagIndicatorByDatasetListResponseIndicatorJSON contains the JSON
// metadata for the struct [ThreatEventTagIndicatorByDatasetListResponseIndicator]
type threatEventTagIndicatorByDatasetListResponseIndicatorJSON struct {
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

func (r *ThreatEventTagIndicatorByDatasetListResponseIndicator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagIndicatorByDatasetListResponseIndicatorJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagIndicatorByDatasetListResponseIndicatorsRelatedEvent struct {
	DatasetID string                                                                 `json:"datasetId" api:"required"`
	EventID   string                                                                 `json:"eventId" api:"required"`
	JSON      threatEventTagIndicatorByDatasetListResponseIndicatorsRelatedEventJSON `json:"-"`
}

// threatEventTagIndicatorByDatasetListResponseIndicatorsRelatedEventJSON contains
// the JSON metadata for the struct
// [ThreatEventTagIndicatorByDatasetListResponseIndicatorsRelatedEvent]
type threatEventTagIndicatorByDatasetListResponseIndicatorsRelatedEventJSON struct {
	DatasetID   apijson.Field
	EventID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagIndicatorByDatasetListResponseIndicatorsRelatedEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagIndicatorByDatasetListResponseIndicatorsRelatedEventJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagIndicatorByDatasetListResponseIndicatorsTag struct {
	CategoryName string                                                        `json:"categoryName"`
	UUID         string                                                        `json:"uuid"`
	Value        string                                                        `json:"value"`
	JSON         threatEventTagIndicatorByDatasetListResponseIndicatorsTagJSON `json:"-"`
}

// threatEventTagIndicatorByDatasetListResponseIndicatorsTagJSON contains the JSON
// metadata for the struct
// [ThreatEventTagIndicatorByDatasetListResponseIndicatorsTag]
type threatEventTagIndicatorByDatasetListResponseIndicatorsTagJSON struct {
	CategoryName apijson.Field
	UUID         apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ThreatEventTagIndicatorByDatasetListResponseIndicatorsTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagIndicatorByDatasetListResponseIndicatorsTagJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagIndicatorByDatasetListResponsePagination struct {
	Page       float64                                                    `json:"page" api:"required"`
	PageSize   float64                                                    `json:"pageSize" api:"required"`
	TotalCount float64                                                    `json:"totalCount" api:"required"`
	TotalPages float64                                                    `json:"totalPages" api:"required"`
	JSON       threatEventTagIndicatorByDatasetListResponsePaginationJSON `json:"-"`
}

// threatEventTagIndicatorByDatasetListResponsePaginationJSON contains the JSON
// metadata for the struct [ThreatEventTagIndicatorByDatasetListResponsePagination]
type threatEventTagIndicatorByDatasetListResponsePaginationJSON struct {
	Page        apijson.Field
	PageSize    apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagIndicatorByDatasetListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagIndicatorByDatasetListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagIndicatorByDatasetListParams struct {
	// Account ID.
	AccountID     param.Field[string]  `path:"account_id" api:"required"`
	IndicatorType param.Field[string]  `query:"indicatorType"`
	Page          param.Field[float64] `query:"page"`
	PageSize      param.Field[float64] `query:"pageSize"`
	// Filter indicators by related event UUID(s). Multiple UUIDs can be provided by
	// repeating the parameter.
	RelatedEvent param.Field[[]string] `query:"relatedEvent"`
	// Structured search as a JSON array of {field, op, value} objects. Searchable
	// fields: value, indicatorType. Multiple conditions are AND'd together. Max 10
	// conditions per request.
	Search param.Field[[]ThreatEventTagIndicatorByDatasetListParamsSearch] `query:"search"`
}

// URLQuery serializes [ThreatEventTagIndicatorByDatasetListParams]'s query
// parameters as `url.Values`.
func (r ThreatEventTagIndicatorByDatasetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ThreatEventTagIndicatorByDatasetListParamsSearch struct {
	// The indicator field to search on. Allowed: value, indicatorType, uuid.
	Field param.Field[ThreatEventTagIndicatorByDatasetListParamsSearchField] `query:"field" api:"required"`
	// Search operator. Use 'in' for bulk lookup of up to 100 values at once, e.g.
	// {field:'value', op:'in', value:['evil.com','bad.org']}.
	Op param.Field[ThreatEventTagIndicatorByDatasetListParamsSearchOp] `query:"op" api:"required"`
	// Search value. String for most operators. Array of strings for 'in' operator (max
	// 100 items).
	Value param.Field[ThreatEventTagIndicatorByDatasetListParamsSearchValueUnion] `query:"value" api:"required"`
}

// URLQuery serializes [ThreatEventTagIndicatorByDatasetListParamsSearch]'s query
// parameters as `url.Values`.
func (r ThreatEventTagIndicatorByDatasetListParamsSearch) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The indicator field to search on. Allowed: value, indicatorType, uuid.
type ThreatEventTagIndicatorByDatasetListParamsSearchField string

const (
	ThreatEventTagIndicatorByDatasetListParamsSearchFieldValue         ThreatEventTagIndicatorByDatasetListParamsSearchField = "value"
	ThreatEventTagIndicatorByDatasetListParamsSearchFieldIndicatorType ThreatEventTagIndicatorByDatasetListParamsSearchField = "indicatorType"
	ThreatEventTagIndicatorByDatasetListParamsSearchFieldUUID          ThreatEventTagIndicatorByDatasetListParamsSearchField = "uuid"
)

func (r ThreatEventTagIndicatorByDatasetListParamsSearchField) IsKnown() bool {
	switch r {
	case ThreatEventTagIndicatorByDatasetListParamsSearchFieldValue, ThreatEventTagIndicatorByDatasetListParamsSearchFieldIndicatorType, ThreatEventTagIndicatorByDatasetListParamsSearchFieldUUID:
		return true
	}
	return false
}

// Search operator. Use 'in' for bulk lookup of up to 100 values at once, e.g.
// {field:'value', op:'in', value:['evil.com','bad.org']}.
type ThreatEventTagIndicatorByDatasetListParamsSearchOp string

const (
	ThreatEventTagIndicatorByDatasetListParamsSearchOpEquals     ThreatEventTagIndicatorByDatasetListParamsSearchOp = "equals"
	ThreatEventTagIndicatorByDatasetListParamsSearchOpNot        ThreatEventTagIndicatorByDatasetListParamsSearchOp = "not"
	ThreatEventTagIndicatorByDatasetListParamsSearchOpGt         ThreatEventTagIndicatorByDatasetListParamsSearchOp = "gt"
	ThreatEventTagIndicatorByDatasetListParamsSearchOpGte        ThreatEventTagIndicatorByDatasetListParamsSearchOp = "gte"
	ThreatEventTagIndicatorByDatasetListParamsSearchOpLt         ThreatEventTagIndicatorByDatasetListParamsSearchOp = "lt"
	ThreatEventTagIndicatorByDatasetListParamsSearchOpLte        ThreatEventTagIndicatorByDatasetListParamsSearchOp = "lte"
	ThreatEventTagIndicatorByDatasetListParamsSearchOpLike       ThreatEventTagIndicatorByDatasetListParamsSearchOp = "like"
	ThreatEventTagIndicatorByDatasetListParamsSearchOpContains   ThreatEventTagIndicatorByDatasetListParamsSearchOp = "contains"
	ThreatEventTagIndicatorByDatasetListParamsSearchOpStartsWith ThreatEventTagIndicatorByDatasetListParamsSearchOp = "startsWith"
	ThreatEventTagIndicatorByDatasetListParamsSearchOpEndsWith   ThreatEventTagIndicatorByDatasetListParamsSearchOp = "endsWith"
	ThreatEventTagIndicatorByDatasetListParamsSearchOpIn         ThreatEventTagIndicatorByDatasetListParamsSearchOp = "in"
	ThreatEventTagIndicatorByDatasetListParamsSearchOpFind       ThreatEventTagIndicatorByDatasetListParamsSearchOp = "find"
)

func (r ThreatEventTagIndicatorByDatasetListParamsSearchOp) IsKnown() bool {
	switch r {
	case ThreatEventTagIndicatorByDatasetListParamsSearchOpEquals, ThreatEventTagIndicatorByDatasetListParamsSearchOpNot, ThreatEventTagIndicatorByDatasetListParamsSearchOpGt, ThreatEventTagIndicatorByDatasetListParamsSearchOpGte, ThreatEventTagIndicatorByDatasetListParamsSearchOpLt, ThreatEventTagIndicatorByDatasetListParamsSearchOpLte, ThreatEventTagIndicatorByDatasetListParamsSearchOpLike, ThreatEventTagIndicatorByDatasetListParamsSearchOpContains, ThreatEventTagIndicatorByDatasetListParamsSearchOpStartsWith, ThreatEventTagIndicatorByDatasetListParamsSearchOpEndsWith, ThreatEventTagIndicatorByDatasetListParamsSearchOpIn, ThreatEventTagIndicatorByDatasetListParamsSearchOpFind:
		return true
	}
	return false
}

// Search value. String for most operators. Array of strings for 'in' operator (max
// 100 items).
//
// Satisfied by [shared.UnionString],
// [cloudforce_one.ThreatEventTagIndicatorByDatasetListParamsSearchValueArray].
type ThreatEventTagIndicatorByDatasetListParamsSearchValueUnion interface {
	ImplementsThreatEventTagIndicatorByDatasetListParamsSearchValueUnion()
}

type ThreatEventTagIndicatorByDatasetListParamsSearchValueArray []string

func (r ThreatEventTagIndicatorByDatasetListParamsSearchValueArray) ImplementsThreatEventTagIndicatorByDatasetListParamsSearchValueUnion() {
}
