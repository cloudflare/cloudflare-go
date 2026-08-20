// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logs

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
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// LogExplorerDatasetService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogExplorerDatasetService] method instead.
type LogExplorerDatasetService struct {
	Options   []option.RequestOption
	Available *LogExplorerDatasetAvailableService
}

// NewLogExplorerDatasetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLogExplorerDatasetService(opts ...option.RequestOption) (r *LogExplorerDatasetService) {
	r = &LogExplorerDatasetService{}
	r.Options = opts
	r.Available = NewLogExplorerDatasetAvailableService(opts...)
	return
}

// Create a new Log Explorer dataset for the account or zone.
//
// Use the
// `/account or zones/{account or zone_id}/logs/explorer/datasets/available`
// endpoint to list dataset types you can create along with their available fields.
//
// The `fields` property is optional. If not specified, all available fields will
// be enabled.
//
// For zone-level datasets use the zone-scoped endpoint: POST
// /zones/{zone_id}/logs/explorer/datasets
//
// For dataset field definitions, see:
// https://developers.cloudflare.com/logs/logpush/logpush-job/datasets/
func (r *LogExplorerDatasetService) New(ctx context.Context, params LogExplorerDatasetNewParams, opts ...option.RequestOption) (res *Dataset, err error) {
	var env LogExplorerDatasetNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logs/explorer/datasets", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates the enabled state and/or field configuration of an account or zone
// dataset.
func (r *LogExplorerDatasetService) Update(ctx context.Context, datasetID string, params LogExplorerDatasetUpdateParams, opts ...option.RequestOption) (res *Dataset, err error) {
	var env LogExplorerDatasetUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/logs/explorer/datasets/%s", accountOrZone, accountOrZoneID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns all Log Explorer datasets configured for the account or zone.
//
// Pass `include_zones=true` to also include zone-level datasets that belong to
// this account or zone. List responses omit the `fields` property; use the
// single-dataset endpoint to retrieve field configuration.
func (r *LogExplorerDatasetService) List(ctx context.Context, params LogExplorerDatasetListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DatasetSummary], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logs/explorer/datasets", accountOrZone, accountOrZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Returns all Log Explorer datasets configured for the account or zone.
//
// Pass `include_zones=true` to also include zone-level datasets that belong to
// this account or zone. List responses omit the `fields` property; use the
// single-dataset endpoint to retrieve field configuration.
func (r *LogExplorerDatasetService) ListAutoPaging(ctx context.Context, params LogExplorerDatasetListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DatasetSummary] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Retrieve a single Log Explorer dataset by ID for the account or zone.
func (r *LogExplorerDatasetService) Get(ctx context.Context, datasetID string, query LogExplorerDatasetGetParams, opts ...option.RequestOption) (res *Dataset, err error) {
	var env LogExplorerDatasetGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/logs/explorer/datasets/%s", accountOrZone, accountOrZoneID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type CreateRequestParam struct {
	// Dataset type name to create (e.g. `http_requests`).
	Dataset param.Field[string] `json:"dataset" api:"required"`
	// Controls which fields the API ingests. Defaults to all available fields when
	// absent.
	Fields param.Field[[]CreateRequestFieldParam] `json:"fields"`
}

func (r CreateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CreateRequestFieldParam struct {
	// Whether the API includes this field in log ingest.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Field name in lowercase.
	Name param.Field[string] `json:"name" api:"required"`
}

func (r CreateRequestFieldParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A Log Explorer dataset summary. List endpoints return this type and omit field
// configuration; use the single-dataset endpoint to retrieve it.
type Dataset struct {
	// RFC3339 timestamp recording when the API created this dataset.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Dataset type name (e.g. `http_requests`).
	Dataset string `json:"dataset" api:"required"`
	// Unique dataset ID.
	DatasetID string `json:"dataset_id" api:"required"`
	// Whether log ingest is currently active for this dataset.
	Enabled bool `json:"enabled" api:"required"`
	// Public ID of the account or zone that owns this dataset.
	ObjectID string `json:"object_id" api:"required"`
	// Whether this dataset belongs to an account or a zone.
	ObjectType DatasetObjectType `json:"object_type" api:"required"`
	// RFC3339 timestamp recording when the API last updated this dataset.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The field configuration for this dataset.
	Fields []DatasetField `json:"fields"`
	JSON   datasetJSON    `json:"-"`
}

// datasetJSON contains the JSON metadata for the struct [Dataset]
type datasetJSON struct {
	CreatedAt   apijson.Field
	Dataset     apijson.Field
	DatasetID   apijson.Field
	Enabled     apijson.Field
	ObjectID    apijson.Field
	ObjectType  apijson.Field
	UpdatedAt   apijson.Field
	Fields      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Dataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetJSON) RawJSON() string {
	return r.raw
}

// Whether this dataset belongs to an account or a zone.
type DatasetObjectType string

const (
	DatasetObjectTypeAccount DatasetObjectType = "account"
	DatasetObjectTypeZone    DatasetObjectType = "zone"
)

func (r DatasetObjectType) IsKnown() bool {
	switch r {
	case DatasetObjectTypeAccount, DatasetObjectTypeZone:
		return true
	}
	return false
}

type DatasetField struct {
	// Whether the API includes this field in log ingest.
	Enabled bool `json:"enabled" api:"required"`
	// Field name in lowercase.
	Name string           `json:"name" api:"required"`
	JSON datasetFieldJSON `json:"-"`
}

// datasetFieldJSON contains the JSON metadata for the struct [DatasetField]
type datasetFieldJSON struct {
	Enabled     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetFieldJSON) RawJSON() string {
	return r.raw
}

// A Log Explorer dataset summary. List endpoints return this type and omit field
// configuration; use the single-dataset endpoint to retrieve it.
type DatasetSummary struct {
	// RFC3339 timestamp recording when the API created this dataset.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Dataset type name (e.g. `http_requests`).
	Dataset string `json:"dataset" api:"required"`
	// Unique dataset ID.
	DatasetID string `json:"dataset_id" api:"required"`
	// Whether log ingest is currently active for this dataset.
	Enabled bool `json:"enabled" api:"required"`
	// Public ID of the account or zone that owns this dataset.
	ObjectID string `json:"object_id" api:"required"`
	// Whether this dataset belongs to an account or a zone.
	ObjectType DatasetSummaryObjectType `json:"object_type" api:"required"`
	// RFC3339 timestamp recording when the API last updated this dataset.
	UpdatedAt time.Time          `json:"updated_at" api:"required" format:"date-time"`
	JSON      datasetSummaryJSON `json:"-"`
}

// datasetSummaryJSON contains the JSON metadata for the struct [DatasetSummary]
type datasetSummaryJSON struct {
	CreatedAt   apijson.Field
	Dataset     apijson.Field
	DatasetID   apijson.Field
	Enabled     apijson.Field
	ObjectID    apijson.Field
	ObjectType  apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetSummaryJSON) RawJSON() string {
	return r.raw
}

// Whether this dataset belongs to an account or a zone.
type DatasetSummaryObjectType string

const (
	DatasetSummaryObjectTypeAccount DatasetSummaryObjectType = "account"
	DatasetSummaryObjectTypeZone    DatasetSummaryObjectType = "zone"
)

func (r DatasetSummaryObjectType) IsKnown() bool {
	switch r {
	case DatasetSummaryObjectTypeAccount, DatasetSummaryObjectTypeZone:
		return true
	}
	return false
}

type UpdateRequestParam struct {
	// Whether to enable or disable log ingest for this dataset.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Controls which fields the API ingests after the update. Defaults to all
	// available fields when absent.
	Fields param.Field[[]UpdateRequestFieldParam] `json:"fields"`
}

func (r UpdateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdateRequestFieldParam struct {
	// Whether the API includes this field in log ingest.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Field name in lowercase.
	Name param.Field[string] `json:"name" api:"required"`
}

func (r UpdateRequestFieldParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LogExplorerDatasetNewParams struct {
	CreateRequest CreateRequestParam `json:"create_request" api:"required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r LogExplorerDatasetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateRequest)
}

type LogExplorerDatasetNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []string              `json:"messages" api:"required"`
	Success  bool                  `json:"success" api:"required"`
	// A Log Explorer dataset summary. List endpoints return this type and omit field
	// configuration; use the single-dataset endpoint to retrieve it.
	Result Dataset                                   `json:"result"`
	JSON   logExplorerDatasetNewResponseEnvelopeJSON `json:"-"`
}

// logExplorerDatasetNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [LogExplorerDatasetNewResponseEnvelope]
type logExplorerDatasetNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogExplorerDatasetNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logExplorerDatasetNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LogExplorerDatasetUpdateParams struct {
	UpdateRequest UpdateRequestParam `json:"update_request" api:"required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r LogExplorerDatasetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UpdateRequest)
}

type LogExplorerDatasetUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []string              `json:"messages" api:"required"`
	Success  bool                  `json:"success" api:"required"`
	// A Log Explorer dataset summary. List endpoints return this type and omit field
	// configuration; use the single-dataset endpoint to retrieve it.
	Result Dataset                                      `json:"result"`
	JSON   logExplorerDatasetUpdateResponseEnvelopeJSON `json:"-"`
}

// logExplorerDatasetUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [LogExplorerDatasetUpdateResponseEnvelope]
type logExplorerDatasetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogExplorerDatasetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logExplorerDatasetUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LogExplorerDatasetListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Set to true to include zone-scoped datasets belonging to this account.
	IncludeZones param.Field[bool] `query:"include_zones"`
}

// URLQuery serializes [LogExplorerDatasetListParams]'s query parameters as
// `url.Values`.
func (r LogExplorerDatasetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogExplorerDatasetGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type LogExplorerDatasetGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []string              `json:"messages" api:"required"`
	Success  bool                  `json:"success" api:"required"`
	// A Log Explorer dataset summary. List endpoints return this type and omit field
	// configuration; use the single-dataset endpoint to retrieve it.
	Result Dataset                                   `json:"result"`
	JSON   logExplorerDatasetGetResponseEnvelopeJSON `json:"-"`
}

// logExplorerDatasetGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [LogExplorerDatasetGetResponseEnvelope]
type logExplorerDatasetGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogExplorerDatasetGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logExplorerDatasetGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
