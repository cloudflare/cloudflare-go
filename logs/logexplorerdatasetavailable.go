// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logs

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// LogExplorerDatasetAvailableService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogExplorerDatasetAvailableService] method instead.
type LogExplorerDatasetAvailableService struct {
	Options []option.RequestOption
}

// NewLogExplorerDatasetAvailableService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLogExplorerDatasetAvailableService(opts ...option.RequestOption) (r *LogExplorerDatasetAvailableService) {
	r = &LogExplorerDatasetAvailableService{}
	r.Options = opts
	return
}

// Returns all dataset types that this account or zone can create. Each entry
// includes the dataset schema and timestamp field.
//
// The schema shows all possible fields for a dataset. However, not all fields may
// be available for your account or zone. When creating or updating a dataset, only
// fields available to your account or zone can be enabled. If you request a field
// that is not available, you will receive an error.
func (r *LogExplorerDatasetAvailableService) List(ctx context.Context, query LogExplorerDatasetAvailableListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AvailableDataset], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	path := fmt.Sprintf("%s/%s/logs/explorer/datasets/available", accountOrZone, accountOrZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Returns all dataset types that this account or zone can create. Each entry
// includes the dataset schema and timestamp field.
//
// The schema shows all possible fields for a dataset. However, not all fields may
// be available for your account or zone. When creating or updating a dataset, only
// fields available to your account or zone can be enabled. If you request a field
// that is not available, you will receive an error.
func (r *LogExplorerDatasetAvailableService) ListAutoPaging(ctx context.Context, query LogExplorerDatasetAvailableListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AvailableDataset] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// A dataset type that the account or zone can create.
type AvailableDataset struct {
	// Dataset type name (e.g. `http_requests`).
	Dataset string `json:"dataset" api:"required"`
	// Whether this dataset type is account-scoped or zone-scoped.
	ObjectType AvailableDatasetObjectType `json:"object_type" api:"required"`
	// JSON Schema that describes the fields this dataset exposes.
	Schema AvailableDatasetSchema `json:"schema" api:"required"`
	// The primary timestamp field name for this dataset.
	TimestampField string               `json:"timestamp_field" api:"required"`
	JSON           availableDatasetJSON `json:"-"`
}

// availableDatasetJSON contains the JSON metadata for the struct
// [AvailableDataset]
type availableDatasetJSON struct {
	Dataset        apijson.Field
	ObjectType     apijson.Field
	Schema         apijson.Field
	TimestampField apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AvailableDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availableDatasetJSON) RawJSON() string {
	return r.raw
}

// Whether this dataset type is account-scoped or zone-scoped.
type AvailableDatasetObjectType string

const (
	AvailableDatasetObjectTypeAccount AvailableDatasetObjectType = "account"
	AvailableDatasetObjectTypeZone    AvailableDatasetObjectType = "zone"
)

func (r AvailableDatasetObjectType) IsKnown() bool {
	switch r {
	case AvailableDatasetObjectTypeAccount, AvailableDatasetObjectTypeZone:
		return true
	}
	return false
}

// JSON Schema that describes the fields this dataset exposes.
type AvailableDatasetSchema struct {
	Properties map[string]interface{}     `json:"properties"`
	Required   []string                   `json:"required"`
	Type       AvailableDatasetSchemaType `json:"type"`
	JSON       availableDatasetSchemaJSON `json:"-"`
}

// availableDatasetSchemaJSON contains the JSON metadata for the struct
// [AvailableDatasetSchema]
type availableDatasetSchemaJSON struct {
	Properties  apijson.Field
	Required    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableDatasetSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availableDatasetSchemaJSON) RawJSON() string {
	return r.raw
}

type AvailableDatasetSchemaType string

const (
	AvailableDatasetSchemaTypeObject AvailableDatasetSchemaType = "object"
)

func (r AvailableDatasetSchemaType) IsKnown() bool {
	switch r {
	case AvailableDatasetSchemaTypeObject:
		return true
	}
	return false
}

type AvailableList struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []string              `json:"messages" api:"required"`
	Success  bool                  `json:"success" api:"required"`
	Result   []AvailableDataset    `json:"result" api:"nullable"`
	JSON     availableListJSON     `json:"-"`
}

// availableListJSON contains the JSON metadata for the struct [AvailableList]
type availableListJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availableListJSON) RawJSON() string {
	return r.raw
}

type LogExplorerDatasetAvailableListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}
