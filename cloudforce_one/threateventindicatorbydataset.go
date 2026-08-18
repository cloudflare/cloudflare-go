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

// ThreatEventIndicatorByDatasetService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventIndicatorByDatasetService] method instead.
type ThreatEventIndicatorByDatasetService struct {
	Options []option.RequestOption
	Tags    *ThreatEventIndicatorByDatasetTagService
}

// NewThreatEventIndicatorByDatasetService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewThreatEventIndicatorByDatasetService(opts ...option.RequestOption) (r *ThreatEventIndicatorByDatasetService) {
	r = &ThreatEventIndicatorByDatasetService{}
	r.Options = opts
	r.Tags = NewThreatEventIndicatorByDatasetTagService(opts...)
	return
}

// This method is deprecated. Please use /events/indicators to retrieve a paginated
// list of indicators.
//
// Deprecated: Use indicators.list instead (GET
// /accounts/{account_id}/cloudforce-one/events/indicators).
func (r *ThreatEventIndicatorByDatasetService) List(ctx context.Context, datasetID string, params ThreatEventIndicatorByDatasetListParams, opts ...option.RequestOption) (res *ThreatEventIndicatorByDatasetListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/dataset/%s/indicators", params.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Retrieves a specific indicator by its UUID.
func (r *ThreatEventIndicatorByDatasetService) Get(ctx context.Context, datasetID string, indicatorID string, query ThreatEventIndicatorByDatasetGetParams, opts ...option.RequestOption) (res *ThreatEventIndicatorByDatasetGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return nil, err
	}
	if indicatorID == "" {
		err = errors.New("missing required indicator_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/dataset/%s/indicators/%s", query.AccountID, datasetID, indicatorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ThreatEventIndicatorByDatasetListResponse struct {
	Indicators []ThreatEventIndicatorByDatasetListResponseIndicator `json:"indicators" api:"required"`
	Pagination ThreatEventIndicatorByDatasetListResponsePagination  `json:"pagination" api:"required"`
	JSON       threatEventIndicatorByDatasetListResponseJSON        `json:"-"`
}

// threatEventIndicatorByDatasetListResponseJSON contains the JSON metadata for the
// struct [ThreatEventIndicatorByDatasetListResponse]
type threatEventIndicatorByDatasetListResponseJSON struct {
	Indicators  apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventIndicatorByDatasetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorByDatasetListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorByDatasetListResponseIndicator struct {
	CreatedAt     time.Time `json:"createdAt" api:"required" format:"date-time"`
	IndicatorType string    `json:"indicatorType" api:"required"`
	UpdatedAt     time.Time `json:"updatedAt" api:"required" format:"date-time"`
	UUID          string    `json:"uuid" api:"required"`
	Value         string    `json:"value" api:"required"`
	// The dataset ID this indicator belongs to. Included in list responses.
	DatasetID     string                                                            `json:"datasetId"`
	RelatedEvents []ThreatEventIndicatorByDatasetListResponseIndicatorsRelatedEvent `json:"relatedEvents"`
	Tags          []ThreatEventIndicatorByDatasetListResponseIndicatorsTag          `json:"tags"`
	JSON          threatEventIndicatorByDatasetListResponseIndicatorJSON            `json:"-"`
}

// threatEventIndicatorByDatasetListResponseIndicatorJSON contains the JSON
// metadata for the struct [ThreatEventIndicatorByDatasetListResponseIndicator]
type threatEventIndicatorByDatasetListResponseIndicatorJSON struct {
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

func (r *ThreatEventIndicatorByDatasetListResponseIndicator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorByDatasetListResponseIndicatorJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorByDatasetListResponseIndicatorsRelatedEvent struct {
	DatasetID string `json:"datasetId" api:"required"`
	EventID   string `json:"eventId" api:"required"`
	// ISO 8601 date of the related event. Null for legacy relationships created before
	// event-date tracking was added.
	EventDate string                                                              `json:"eventDate" api:"nullable"`
	JSON      threatEventIndicatorByDatasetListResponseIndicatorsRelatedEventJSON `json:"-"`
}

// threatEventIndicatorByDatasetListResponseIndicatorsRelatedEventJSON contains the
// JSON metadata for the struct
// [ThreatEventIndicatorByDatasetListResponseIndicatorsRelatedEvent]
type threatEventIndicatorByDatasetListResponseIndicatorsRelatedEventJSON struct {
	DatasetID   apijson.Field
	EventID     apijson.Field
	EventDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventIndicatorByDatasetListResponseIndicatorsRelatedEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorByDatasetListResponseIndicatorsRelatedEventJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorByDatasetListResponseIndicatorsTag struct {
	CategoryName string                                                     `json:"categoryName"`
	UUID         string                                                     `json:"uuid"`
	Value        string                                                     `json:"value"`
	JSON         threatEventIndicatorByDatasetListResponseIndicatorsTagJSON `json:"-"`
}

// threatEventIndicatorByDatasetListResponseIndicatorsTagJSON contains the JSON
// metadata for the struct [ThreatEventIndicatorByDatasetListResponseIndicatorsTag]
type threatEventIndicatorByDatasetListResponseIndicatorsTagJSON struct {
	CategoryName apijson.Field
	UUID         apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ThreatEventIndicatorByDatasetListResponseIndicatorsTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorByDatasetListResponseIndicatorsTagJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorByDatasetListResponsePagination struct {
	Page       float64                                                 `json:"page" api:"required"`
	PageSize   float64                                                 `json:"pageSize" api:"required"`
	TotalCount float64                                                 `json:"totalCount" api:"required"`
	TotalPages float64                                                 `json:"totalPages" api:"required"`
	JSON       threatEventIndicatorByDatasetListResponsePaginationJSON `json:"-"`
}

// threatEventIndicatorByDatasetListResponsePaginationJSON contains the JSON
// metadata for the struct [ThreatEventIndicatorByDatasetListResponsePagination]
type threatEventIndicatorByDatasetListResponsePaginationJSON struct {
	Page        apijson.Field
	PageSize    apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventIndicatorByDatasetListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorByDatasetListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorByDatasetGetResponse struct {
	CreatedAt     time.Time `json:"createdAt" api:"required" format:"date-time"`
	IndicatorType string    `json:"indicatorType" api:"required"`
	UpdatedAt     time.Time `json:"updatedAt" api:"required" format:"date-time"`
	UUID          string    `json:"uuid" api:"required"`
	Value         string    `json:"value" api:"required"`
	// The dataset ID this indicator belongs to. Included in list responses.
	DatasetID     string                                                 `json:"datasetId"`
	RelatedEvents []ThreatEventIndicatorByDatasetGetResponseRelatedEvent `json:"relatedEvents"`
	Tags          []ThreatEventIndicatorByDatasetGetResponseTag          `json:"tags"`
	JSON          threatEventIndicatorByDatasetGetResponseJSON           `json:"-"`
}

// threatEventIndicatorByDatasetGetResponseJSON contains the JSON metadata for the
// struct [ThreatEventIndicatorByDatasetGetResponse]
type threatEventIndicatorByDatasetGetResponseJSON struct {
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

func (r *ThreatEventIndicatorByDatasetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorByDatasetGetResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorByDatasetGetResponseRelatedEvent struct {
	DatasetID string `json:"datasetId" api:"required"`
	EventID   string `json:"eventId" api:"required"`
	// ISO 8601 date of the related event. Null for legacy relationships created before
	// event-date tracking was added.
	EventDate string                                                   `json:"eventDate" api:"nullable"`
	JSON      threatEventIndicatorByDatasetGetResponseRelatedEventJSON `json:"-"`
}

// threatEventIndicatorByDatasetGetResponseRelatedEventJSON contains the JSON
// metadata for the struct [ThreatEventIndicatorByDatasetGetResponseRelatedEvent]
type threatEventIndicatorByDatasetGetResponseRelatedEventJSON struct {
	DatasetID   apijson.Field
	EventID     apijson.Field
	EventDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventIndicatorByDatasetGetResponseRelatedEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorByDatasetGetResponseRelatedEventJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorByDatasetGetResponseTag struct {
	CategoryName string                                          `json:"categoryName"`
	UUID         string                                          `json:"uuid"`
	Value        string                                          `json:"value"`
	JSON         threatEventIndicatorByDatasetGetResponseTagJSON `json:"-"`
}

// threatEventIndicatorByDatasetGetResponseTagJSON contains the JSON metadata for
// the struct [ThreatEventIndicatorByDatasetGetResponseTag]
type threatEventIndicatorByDatasetGetResponseTagJSON struct {
	CategoryName apijson.Field
	UUID         apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ThreatEventIndicatorByDatasetGetResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventIndicatorByDatasetGetResponseTagJSON) RawJSON() string {
	return r.raw
}

type ThreatEventIndicatorByDatasetListParams struct {
	// Account ID.
	AccountID     param.Field[string] `path:"account_id" api:"required"`
	IndicatorType param.Field[string] `query:"indicatorType"`
	// Filter by indicator value (substring match)
	Name     param.Field[string]  `query:"name"`
	Page     param.Field[float64] `query:"page"`
	PageSize param.Field[float64] `query:"pageSize"`
	// Filter indicators by related event UUID(s). Multiple UUIDs can be provided by
	// repeating the parameter.
	RelatedEvent param.Field[[]string] `query:"relatedEvent"`
}

// URLQuery serializes [ThreatEventIndicatorByDatasetListParams]'s query parameters
// as `url.Values`.
func (r ThreatEventIndicatorByDatasetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ThreatEventIndicatorByDatasetGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
