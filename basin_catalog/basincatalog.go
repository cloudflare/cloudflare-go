// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package basin_catalog

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

// BasinCatalogService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBasinCatalogService] method instead.
type BasinCatalogService struct {
	Options            []option.RequestOption
	MaintenanceConfigs *MaintenanceConfigService
	Credentials        *CredentialService
	Namespaces         *NamespaceService
}

// NewBasinCatalogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBasinCatalogService(opts ...option.RequestOption) (r *BasinCatalogService) {
	r = &BasinCatalogService{}
	r.Options = opts
	r.MaintenanceConfigs = NewMaintenanceConfigService(opts...)
	r.Credentials = NewCredentialService(opts...)
	r.Namespaces = NewNamespaceService(opts...)
	return
}

// Returns a list of R2 buckets that have been enabled as Apache Iceberg catalogs
// for the specified account. Each catalog represents an R2 bucket configured to
// store Iceberg metadata and data files.
func (r *BasinCatalogService) List(ctx context.Context, query BasinCatalogListParams, opts ...option.RequestOption) (res *BasinCatalogListResponse, err error) {
	var env BasinCatalogListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/basin-catalog", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Removes the catalog from the control plane without deleting R2 bucket objects.
// Set force=true to remove catalog namespaces, tables, views, and maintenance
// metadata. Force deletion is limited to a configured catalog object count.
func (r *BasinCatalogService) Delete(ctx context.Context, bucketName string, params BasinCatalogDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return err
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return err
	}
	path := fmt.Sprintf("accounts/%s/basin-catalog/%s/delete", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Disable an R2 bucket as a catalog. This operation deactivates the catalog but
// preserves existing metadata and data files. The catalog can be re-enabled later.
func (r *BasinCatalogService) Disable(ctx context.Context, bucketName string, body BasinCatalogDisableParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return err
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return err
	}
	path := fmt.Sprintf("accounts/%s/basin-catalog/%s/disable", body.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Enable an R2 bucket as an Apache Iceberg catalog. This operation creates the
// necessary catalog infrastructure and activates the bucket for storing Iceberg
// metadata and data files.
func (r *BasinCatalogService) Enable(ctx context.Context, bucketName string, body BasinCatalogEnableParams, opts ...option.RequestOption) (res *BasinCatalogEnableResponse, err error) {
	var env BasinCatalogEnableResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/basin-catalog/%s/enable", body.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve detailed information about a specific Basin Catalog by bucket name.
// Returns catalog status, maintenance configuration, and credential status.
func (r *BasinCatalogService) Get(ctx context.Context, bucketName string, query BasinCatalogGetParams, opts ...option.RequestOption) (res *BasinCatalogGetResponse, err error) {
	var env BasinCatalogGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/basin-catalog/%s", query.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Contains the list of catalogs.
type BasinCatalogListResponse struct {
	// Lists catalogs in the account.
	Warehouses []BasinCatalogListResponseWarehouse `json:"warehouses" api:"required"`
	JSON       basinCatalogListResponseJSON        `json:"-"`
}

// basinCatalogListResponseJSON contains the JSON metadata for the struct
// [BasinCatalogListResponse]
type basinCatalogListResponseJSON struct {
	Warehouses  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BasinCatalogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basinCatalogListResponseJSON) RawJSON() string {
	return r.raw
}

// Contains catalog information.
type BasinCatalogListResponseWarehouse struct {
	// Use this to uniquely identify the catalog.
	ID string `json:"id" api:"required" format:"uuid"`
	// Specifies the associated R2 bucket name.
	Bucket string `json:"bucket" api:"required"`
	// Specifies the catalog name (generated from account and bucket name).
	Name string `json:"name" api:"required"`
	// Indicates the status of the catalog.
	Status BasinCatalogListResponseWarehousesStatus `json:"status" api:"required"`
	// Shows the credential configuration status.
	CredentialStatus BasinCatalogListResponseWarehousesCredentialStatus `json:"credential_status" api:"nullable"`
	// Configures maintenance for the catalog.
	MaintenanceConfig BasinCatalogListResponseWarehousesMaintenanceConfig `json:"maintenance_config" api:"nullable"`
	JSON              basinCatalogListResponseWarehouseJSON               `json:"-"`
}

// basinCatalogListResponseWarehouseJSON contains the JSON metadata for the struct
// [BasinCatalogListResponseWarehouse]
type basinCatalogListResponseWarehouseJSON struct {
	ID                apijson.Field
	Bucket            apijson.Field
	Name              apijson.Field
	Status            apijson.Field
	CredentialStatus  apijson.Field
	MaintenanceConfig apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *BasinCatalogListResponseWarehouse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basinCatalogListResponseWarehouseJSON) RawJSON() string {
	return r.raw
}

// Indicates the status of the catalog.
type BasinCatalogListResponseWarehousesStatus string

const (
	BasinCatalogListResponseWarehousesStatusActive   BasinCatalogListResponseWarehousesStatus = "active"
	BasinCatalogListResponseWarehousesStatusInactive BasinCatalogListResponseWarehousesStatus = "inactive"
)

func (r BasinCatalogListResponseWarehousesStatus) IsKnown() bool {
	switch r {
	case BasinCatalogListResponseWarehousesStatusActive, BasinCatalogListResponseWarehousesStatusInactive:
		return true
	}
	return false
}

// Shows the credential configuration status.
type BasinCatalogListResponseWarehousesCredentialStatus string

const (
	BasinCatalogListResponseWarehousesCredentialStatusPresent BasinCatalogListResponseWarehousesCredentialStatus = "present"
	BasinCatalogListResponseWarehousesCredentialStatusAbsent  BasinCatalogListResponseWarehousesCredentialStatus = "absent"
)

func (r BasinCatalogListResponseWarehousesCredentialStatus) IsKnown() bool {
	switch r {
	case BasinCatalogListResponseWarehousesCredentialStatusPresent, BasinCatalogListResponseWarehousesCredentialStatusAbsent:
		return true
	}
	return false
}

// Configures maintenance for the catalog.
type BasinCatalogListResponseWarehousesMaintenanceConfig struct {
	// Configures compaction for catalog maintenance.
	Compaction BasinCatalogListResponseWarehousesMaintenanceConfigCompaction `json:"compaction"`
	// Scheduling interval between normal table maintenance runs.
	Interval string `json:"interval"`
	// Configures snapshot expiration settings.
	SnapshotExpiration BasinCatalogListResponseWarehousesMaintenanceConfigSnapshotExpiration `json:"snapshot_expiration"`
	JSON               basinCatalogListResponseWarehousesMaintenanceConfigJSON               `json:"-"`
}

// basinCatalogListResponseWarehousesMaintenanceConfigJSON contains the JSON
// metadata for the struct [BasinCatalogListResponseWarehousesMaintenanceConfig]
type basinCatalogListResponseWarehousesMaintenanceConfigJSON struct {
	Compaction         apijson.Field
	Interval           apijson.Field
	SnapshotExpiration apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BasinCatalogListResponseWarehousesMaintenanceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basinCatalogListResponseWarehousesMaintenanceConfigJSON) RawJSON() string {
	return r.raw
}

// Configures compaction for catalog maintenance.
type BasinCatalogListResponseWarehousesMaintenanceConfigCompaction struct {
	// Specifies the state of maintenance operations.
	State BasinCatalogListResponseWarehousesMaintenanceConfigCompactionState `json:"state" api:"required"`
	// Sets the target file size for compaction in megabytes. Defaults to "128".
	TargetSizeMB BasinCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB `json:"target_size_mb" api:"required"`
	JSON         basinCatalogListResponseWarehousesMaintenanceConfigCompactionJSON         `json:"-"`
}

// basinCatalogListResponseWarehousesMaintenanceConfigCompactionJSON contains the
// JSON metadata for the struct
// [BasinCatalogListResponseWarehousesMaintenanceConfigCompaction]
type basinCatalogListResponseWarehousesMaintenanceConfigCompactionJSON struct {
	State        apijson.Field
	TargetSizeMB apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BasinCatalogListResponseWarehousesMaintenanceConfigCompaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basinCatalogListResponseWarehousesMaintenanceConfigCompactionJSON) RawJSON() string {
	return r.raw
}

// Specifies the state of maintenance operations.
type BasinCatalogListResponseWarehousesMaintenanceConfigCompactionState string

const (
	BasinCatalogListResponseWarehousesMaintenanceConfigCompactionStateEnabled  BasinCatalogListResponseWarehousesMaintenanceConfigCompactionState = "enabled"
	BasinCatalogListResponseWarehousesMaintenanceConfigCompactionStateDisabled BasinCatalogListResponseWarehousesMaintenanceConfigCompactionState = "disabled"
)

func (r BasinCatalogListResponseWarehousesMaintenanceConfigCompactionState) IsKnown() bool {
	switch r {
	case BasinCatalogListResponseWarehousesMaintenanceConfigCompactionStateEnabled, BasinCatalogListResponseWarehousesMaintenanceConfigCompactionStateDisabled:
		return true
	}
	return false
}

// Sets the target file size for compaction in megabytes. Defaults to "128".
type BasinCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB string

const (
	BasinCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB64  BasinCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB = "64"
	BasinCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB128 BasinCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB = "128"
	BasinCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB256 BasinCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB = "256"
	BasinCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB512 BasinCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB = "512"
)

func (r BasinCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB) IsKnown() bool {
	switch r {
	case BasinCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB64, BasinCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB128, BasinCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB256, BasinCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB512:
		return true
	}
	return false
}

// Configures snapshot expiration settings.
type BasinCatalogListResponseWarehousesMaintenanceConfigSnapshotExpiration struct {
	// Specifies the maximum age for snapshots. The system deletes snapshots older than
	// this age. Format: <number><unit> where unit is d (days), h (hours), m (minutes),
	// or s (seconds). Examples: "7d" (7 days), "48h" (48 hours), "2880m" (2,880
	// minutes). Defaults to "7d".
	MaxSnapshotAge string `json:"max_snapshot_age" api:"required"`
	// Specifies the minimum number of snapshots to retain. Defaults to 100.
	MinSnapshotsToKeep int64 `json:"min_snapshots_to_keep" api:"required"`
	// Specifies the state of maintenance operations.
	State BasinCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationState `json:"state" api:"required"`
	JSON  basinCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationJSON  `json:"-"`
}

// basinCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationJSON
// contains the JSON metadata for the struct
// [BasinCatalogListResponseWarehousesMaintenanceConfigSnapshotExpiration]
type basinCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationJSON struct {
	MaxSnapshotAge     apijson.Field
	MinSnapshotsToKeep apijson.Field
	State              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BasinCatalogListResponseWarehousesMaintenanceConfigSnapshotExpiration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basinCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationJSON) RawJSON() string {
	return r.raw
}

// Specifies the state of maintenance operations.
type BasinCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationState string

const (
	BasinCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationStateEnabled  BasinCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationState = "enabled"
	BasinCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationStateDisabled BasinCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationState = "disabled"
)

func (r BasinCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationState) IsKnown() bool {
	switch r {
	case BasinCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationStateEnabled, BasinCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationStateDisabled:
		return true
	}
	return false
}

// Contains response from activating an R2 bucket as a catalog.
type BasinCatalogEnableResponse struct {
	// Use this to uniquely identify the activated catalog.
	ID string `json:"id" api:"required" format:"uuid"`
	// Specifies the name of the activated catalog.
	Name string                         `json:"name" api:"required"`
	JSON basinCatalogEnableResponseJSON `json:"-"`
}

// basinCatalogEnableResponseJSON contains the JSON metadata for the struct
// [BasinCatalogEnableResponse]
type basinCatalogEnableResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BasinCatalogEnableResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basinCatalogEnableResponseJSON) RawJSON() string {
	return r.raw
}

// Contains catalog information.
type BasinCatalogGetResponse struct {
	// Use this to uniquely identify the catalog.
	ID string `json:"id" api:"required" format:"uuid"`
	// Specifies the associated R2 bucket name.
	Bucket string `json:"bucket" api:"required"`
	// Specifies the catalog name (generated from account and bucket name).
	Name string `json:"name" api:"required"`
	// Indicates the status of the catalog.
	Status BasinCatalogGetResponseStatus `json:"status" api:"required"`
	// Shows the credential configuration status.
	CredentialStatus BasinCatalogGetResponseCredentialStatus `json:"credential_status" api:"nullable"`
	// Configures maintenance for the catalog.
	MaintenanceConfig BasinCatalogGetResponseMaintenanceConfig `json:"maintenance_config" api:"nullable"`
	JSON              basinCatalogGetResponseJSON              `json:"-"`
}

// basinCatalogGetResponseJSON contains the JSON metadata for the struct
// [BasinCatalogGetResponse]
type basinCatalogGetResponseJSON struct {
	ID                apijson.Field
	Bucket            apijson.Field
	Name              apijson.Field
	Status            apijson.Field
	CredentialStatus  apijson.Field
	MaintenanceConfig apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *BasinCatalogGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basinCatalogGetResponseJSON) RawJSON() string {
	return r.raw
}

// Indicates the status of the catalog.
type BasinCatalogGetResponseStatus string

const (
	BasinCatalogGetResponseStatusActive   BasinCatalogGetResponseStatus = "active"
	BasinCatalogGetResponseStatusInactive BasinCatalogGetResponseStatus = "inactive"
)

func (r BasinCatalogGetResponseStatus) IsKnown() bool {
	switch r {
	case BasinCatalogGetResponseStatusActive, BasinCatalogGetResponseStatusInactive:
		return true
	}
	return false
}

// Shows the credential configuration status.
type BasinCatalogGetResponseCredentialStatus string

const (
	BasinCatalogGetResponseCredentialStatusPresent BasinCatalogGetResponseCredentialStatus = "present"
	BasinCatalogGetResponseCredentialStatusAbsent  BasinCatalogGetResponseCredentialStatus = "absent"
)

func (r BasinCatalogGetResponseCredentialStatus) IsKnown() bool {
	switch r {
	case BasinCatalogGetResponseCredentialStatusPresent, BasinCatalogGetResponseCredentialStatusAbsent:
		return true
	}
	return false
}

// Configures maintenance for the catalog.
type BasinCatalogGetResponseMaintenanceConfig struct {
	// Configures compaction for catalog maintenance.
	Compaction BasinCatalogGetResponseMaintenanceConfigCompaction `json:"compaction"`
	// Scheduling interval between normal table maintenance runs.
	Interval string `json:"interval"`
	// Configures snapshot expiration settings.
	SnapshotExpiration BasinCatalogGetResponseMaintenanceConfigSnapshotExpiration `json:"snapshot_expiration"`
	JSON               basinCatalogGetResponseMaintenanceConfigJSON               `json:"-"`
}

// basinCatalogGetResponseMaintenanceConfigJSON contains the JSON metadata for the
// struct [BasinCatalogGetResponseMaintenanceConfig]
type basinCatalogGetResponseMaintenanceConfigJSON struct {
	Compaction         apijson.Field
	Interval           apijson.Field
	SnapshotExpiration apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BasinCatalogGetResponseMaintenanceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basinCatalogGetResponseMaintenanceConfigJSON) RawJSON() string {
	return r.raw
}

// Configures compaction for catalog maintenance.
type BasinCatalogGetResponseMaintenanceConfigCompaction struct {
	// Specifies the state of maintenance operations.
	State BasinCatalogGetResponseMaintenanceConfigCompactionState `json:"state" api:"required"`
	// Sets the target file size for compaction in megabytes. Defaults to "128".
	TargetSizeMB BasinCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB `json:"target_size_mb" api:"required"`
	JSON         basinCatalogGetResponseMaintenanceConfigCompactionJSON         `json:"-"`
}

// basinCatalogGetResponseMaintenanceConfigCompactionJSON contains the JSON
// metadata for the struct [BasinCatalogGetResponseMaintenanceConfigCompaction]
type basinCatalogGetResponseMaintenanceConfigCompactionJSON struct {
	State        apijson.Field
	TargetSizeMB apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BasinCatalogGetResponseMaintenanceConfigCompaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basinCatalogGetResponseMaintenanceConfigCompactionJSON) RawJSON() string {
	return r.raw
}

// Specifies the state of maintenance operations.
type BasinCatalogGetResponseMaintenanceConfigCompactionState string

const (
	BasinCatalogGetResponseMaintenanceConfigCompactionStateEnabled  BasinCatalogGetResponseMaintenanceConfigCompactionState = "enabled"
	BasinCatalogGetResponseMaintenanceConfigCompactionStateDisabled BasinCatalogGetResponseMaintenanceConfigCompactionState = "disabled"
)

func (r BasinCatalogGetResponseMaintenanceConfigCompactionState) IsKnown() bool {
	switch r {
	case BasinCatalogGetResponseMaintenanceConfigCompactionStateEnabled, BasinCatalogGetResponseMaintenanceConfigCompactionStateDisabled:
		return true
	}
	return false
}

// Sets the target file size for compaction in megabytes. Defaults to "128".
type BasinCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB string

const (
	BasinCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB64  BasinCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB = "64"
	BasinCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB128 BasinCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB = "128"
	BasinCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB256 BasinCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB = "256"
	BasinCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB512 BasinCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB = "512"
)

func (r BasinCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB) IsKnown() bool {
	switch r {
	case BasinCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB64, BasinCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB128, BasinCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB256, BasinCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB512:
		return true
	}
	return false
}

// Configures snapshot expiration settings.
type BasinCatalogGetResponseMaintenanceConfigSnapshotExpiration struct {
	// Specifies the maximum age for snapshots. The system deletes snapshots older than
	// this age. Format: <number><unit> where unit is d (days), h (hours), m (minutes),
	// or s (seconds). Examples: "7d" (7 days), "48h" (48 hours), "2880m" (2,880
	// minutes). Defaults to "7d".
	MaxSnapshotAge string `json:"max_snapshot_age" api:"required"`
	// Specifies the minimum number of snapshots to retain. Defaults to 100.
	MinSnapshotsToKeep int64 `json:"min_snapshots_to_keep" api:"required"`
	// Specifies the state of maintenance operations.
	State BasinCatalogGetResponseMaintenanceConfigSnapshotExpirationState `json:"state" api:"required"`
	JSON  basinCatalogGetResponseMaintenanceConfigSnapshotExpirationJSON  `json:"-"`
}

// basinCatalogGetResponseMaintenanceConfigSnapshotExpirationJSON contains the JSON
// metadata for the struct
// [BasinCatalogGetResponseMaintenanceConfigSnapshotExpiration]
type basinCatalogGetResponseMaintenanceConfigSnapshotExpirationJSON struct {
	MaxSnapshotAge     apijson.Field
	MinSnapshotsToKeep apijson.Field
	State              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BasinCatalogGetResponseMaintenanceConfigSnapshotExpiration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basinCatalogGetResponseMaintenanceConfigSnapshotExpirationJSON) RawJSON() string {
	return r.raw
}

// Specifies the state of maintenance operations.
type BasinCatalogGetResponseMaintenanceConfigSnapshotExpirationState string

const (
	BasinCatalogGetResponseMaintenanceConfigSnapshotExpirationStateEnabled  BasinCatalogGetResponseMaintenanceConfigSnapshotExpirationState = "enabled"
	BasinCatalogGetResponseMaintenanceConfigSnapshotExpirationStateDisabled BasinCatalogGetResponseMaintenanceConfigSnapshotExpirationState = "disabled"
)

func (r BasinCatalogGetResponseMaintenanceConfigSnapshotExpirationState) IsKnown() bool {
	switch r {
	case BasinCatalogGetResponseMaintenanceConfigSnapshotExpirationStateEnabled, BasinCatalogGetResponseMaintenanceConfigSnapshotExpirationStateDisabled:
		return true
	}
	return false
}

type BasinCatalogListParams struct {
	// Use this to identify the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type BasinCatalogListResponseEnvelope struct {
	// Contains errors if the API call was unsuccessful.
	Errors []BasinCatalogListResponseEnvelopeErrors `json:"errors" api:"required"`
	// Contains informational messages.
	Messages []BasinCatalogListResponseEnvelopeMessages `json:"messages" api:"required"`
	// Indicates whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Contains the list of catalogs.
	Result BasinCatalogListResponse             `json:"result"`
	JSON   basinCatalogListResponseEnvelopeJSON `json:"-"`
}

// basinCatalogListResponseEnvelopeJSON contains the JSON metadata for the struct
// [BasinCatalogListResponseEnvelope]
type basinCatalogListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BasinCatalogListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basinCatalogListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BasinCatalogListResponseEnvelopeErrors struct {
	// Specifies the error code.
	Code int64 `json:"code" api:"required"`
	// Describes the error.
	Message string                                     `json:"message" api:"required"`
	JSON    basinCatalogListResponseEnvelopeErrorsJSON `json:"-"`
}

// basinCatalogListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [BasinCatalogListResponseEnvelopeErrors]
type basinCatalogListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BasinCatalogListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basinCatalogListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BasinCatalogListResponseEnvelopeMessages struct {
	// Specifies the message code.
	Code int64 `json:"code" api:"required"`
	// Contains the message text.
	Message string                                       `json:"message" api:"required"`
	JSON    basinCatalogListResponseEnvelopeMessagesJSON `json:"-"`
}

// basinCatalogListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [BasinCatalogListResponseEnvelopeMessages]
type basinCatalogListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BasinCatalogListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basinCatalogListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BasinCatalogDeleteParams struct {
	// Use this to identify the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Remove child metadata before deleting the catalog.
	Force param.Field[bool] `query:"force"`
}

// URLQuery serializes [BasinCatalogDeleteParams]'s query parameters as
// `url.Values`.
func (r BasinCatalogDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type BasinCatalogDisableParams struct {
	// Use this to identify the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type BasinCatalogEnableParams struct {
	// Use this to identify the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type BasinCatalogEnableResponseEnvelope struct {
	// Contains errors if the API call was unsuccessful.
	Errors []BasinCatalogEnableResponseEnvelopeErrors `json:"errors" api:"required"`
	// Contains informational messages.
	Messages []BasinCatalogEnableResponseEnvelopeMessages `json:"messages" api:"required"`
	// Indicates whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Contains response from activating an R2 bucket as a catalog.
	Result BasinCatalogEnableResponse             `json:"result"`
	JSON   basinCatalogEnableResponseEnvelopeJSON `json:"-"`
}

// basinCatalogEnableResponseEnvelopeJSON contains the JSON metadata for the struct
// [BasinCatalogEnableResponseEnvelope]
type basinCatalogEnableResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BasinCatalogEnableResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basinCatalogEnableResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BasinCatalogEnableResponseEnvelopeErrors struct {
	// Specifies the error code.
	Code int64 `json:"code" api:"required"`
	// Describes the error.
	Message string                                       `json:"message" api:"required"`
	JSON    basinCatalogEnableResponseEnvelopeErrorsJSON `json:"-"`
}

// basinCatalogEnableResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [BasinCatalogEnableResponseEnvelopeErrors]
type basinCatalogEnableResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BasinCatalogEnableResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basinCatalogEnableResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BasinCatalogEnableResponseEnvelopeMessages struct {
	// Specifies the message code.
	Code int64 `json:"code" api:"required"`
	// Contains the message text.
	Message string                                         `json:"message" api:"required"`
	JSON    basinCatalogEnableResponseEnvelopeMessagesJSON `json:"-"`
}

// basinCatalogEnableResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [BasinCatalogEnableResponseEnvelopeMessages]
type basinCatalogEnableResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BasinCatalogEnableResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basinCatalogEnableResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BasinCatalogGetParams struct {
	// Use this to identify the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type BasinCatalogGetResponseEnvelope struct {
	// Contains errors if the API call was unsuccessful.
	Errors []BasinCatalogGetResponseEnvelopeErrors `json:"errors" api:"required"`
	// Contains informational messages.
	Messages []BasinCatalogGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Indicates whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Contains catalog information.
	Result BasinCatalogGetResponse             `json:"result"`
	JSON   basinCatalogGetResponseEnvelopeJSON `json:"-"`
}

// basinCatalogGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BasinCatalogGetResponseEnvelope]
type basinCatalogGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BasinCatalogGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basinCatalogGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BasinCatalogGetResponseEnvelopeErrors struct {
	// Specifies the error code.
	Code int64 `json:"code" api:"required"`
	// Describes the error.
	Message string                                    `json:"message" api:"required"`
	JSON    basinCatalogGetResponseEnvelopeErrorsJSON `json:"-"`
}

// basinCatalogGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [BasinCatalogGetResponseEnvelopeErrors]
type basinCatalogGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BasinCatalogGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basinCatalogGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BasinCatalogGetResponseEnvelopeMessages struct {
	// Specifies the message code.
	Code int64 `json:"code" api:"required"`
	// Contains the message text.
	Message string                                      `json:"message" api:"required"`
	JSON    basinCatalogGetResponseEnvelopeMessagesJSON `json:"-"`
}

// basinCatalogGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [BasinCatalogGetResponseEnvelopeMessages]
type basinCatalogGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BasinCatalogGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basinCatalogGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
