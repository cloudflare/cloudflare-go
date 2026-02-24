// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_data_catalog

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// R2DataCatalogService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewR2DataCatalogService] method instead.
type R2DataCatalogService struct {
	Options            []option.RequestOption
	MaintenanceConfigs *MaintenanceConfigService
	Credentials        *CredentialService
	Namespaces         *NamespaceService
}

// NewR2DataCatalogService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewR2DataCatalogService(opts ...option.RequestOption) (r *R2DataCatalogService) {
	r = &R2DataCatalogService{}
	r.Options = opts
	r.MaintenanceConfigs = NewMaintenanceConfigService(opts...)
	r.Credentials = NewCredentialService(opts...)
	r.Namespaces = NewNamespaceService(opts...)
	return
}

// Returns a list of R2 buckets that have been enabled as Apache Iceberg catalogs
// for the specified account. Each catalog represents an R2 bucket configured to
// store Iceberg metadata and data files.
func (r *R2DataCatalogService) List(ctx context.Context, query R2DataCatalogListParams, opts ...option.RequestOption) (res *R2DataCatalogListResponse, err error) {
	var env R2DataCatalogListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2-catalog", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disable an R2 bucket as a catalog. This operation deactivates the catalog but
// preserves existing metadata and data files. The catalog can be re-enabled later.
func (r *R2DataCatalogService) Disable(ctx context.Context, bucketName string, body R2DataCatalogDisableParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2-catalog/%s/disable", body.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Enable an R2 bucket as an Apache Iceberg catalog. This operation creates the
// necessary catalog infrastructure and activates the bucket for storing Iceberg
// metadata and data files.
func (r *R2DataCatalogService) Enable(ctx context.Context, bucketName string, body R2DataCatalogEnableParams, opts ...option.RequestOption) (res *R2DataCatalogEnableResponse, err error) {
	var env R2DataCatalogEnableResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2-catalog/%s/enable", body.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve detailed information about a specific R2 catalog by bucket name.
// Returns catalog status, maintenance configuration, and credential status.
func (r *R2DataCatalogService) Get(ctx context.Context, bucketName string, query R2DataCatalogGetParams, opts ...option.RequestOption) (res *R2DataCatalogGetResponse, err error) {
	var env R2DataCatalogGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2-catalog/%s", query.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Contains the list of catalogs.
type R2DataCatalogListResponse struct {
	// Lists catalogs in the account.
	Warehouses []R2DataCatalogListResponseWarehouse `json:"warehouses,required"`
	JSON       r2DataCatalogListResponseJSON        `json:"-"`
}

// r2DataCatalogListResponseJSON contains the JSON metadata for the struct
// [R2DataCatalogListResponse]
type r2DataCatalogListResponseJSON struct {
	Warehouses  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2DataCatalogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2DataCatalogListResponseJSON) RawJSON() string {
	return r.raw
}

// Contains R2 Data Catalog information.
type R2DataCatalogListResponseWarehouse struct {
	// Use this to uniquely identify the catalog.
	ID string `json:"id,required" format:"uuid"`
	// Specifies the associated R2 bucket name.
	Bucket string `json:"bucket,required"`
	// Specifies the catalog name (generated from account and bucket name).
	Name string `json:"name,required"`
	// Indicates the status of the catalog.
	Status R2DataCatalogListResponseWarehousesStatus `json:"status,required"`
	// Shows the credential configuration status.
	CredentialStatus R2DataCatalogListResponseWarehousesCredentialStatus `json:"credential_status,nullable"`
	// Configures maintenance for the catalog.
	MaintenanceConfig R2DataCatalogListResponseWarehousesMaintenanceConfig `json:"maintenance_config,nullable"`
	JSON              r2DataCatalogListResponseWarehouseJSON               `json:"-"`
}

// r2DataCatalogListResponseWarehouseJSON contains the JSON metadata for the struct
// [R2DataCatalogListResponseWarehouse]
type r2DataCatalogListResponseWarehouseJSON struct {
	ID                apijson.Field
	Bucket            apijson.Field
	Name              apijson.Field
	Status            apijson.Field
	CredentialStatus  apijson.Field
	MaintenanceConfig apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *R2DataCatalogListResponseWarehouse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2DataCatalogListResponseWarehouseJSON) RawJSON() string {
	return r.raw
}

// Indicates the status of the catalog.
type R2DataCatalogListResponseWarehousesStatus string

const (
	R2DataCatalogListResponseWarehousesStatusActive   R2DataCatalogListResponseWarehousesStatus = "active"
	R2DataCatalogListResponseWarehousesStatusInactive R2DataCatalogListResponseWarehousesStatus = "inactive"
)

func (r R2DataCatalogListResponseWarehousesStatus) IsKnown() bool {
	switch r {
	case R2DataCatalogListResponseWarehousesStatusActive, R2DataCatalogListResponseWarehousesStatusInactive:
		return true
	}
	return false
}

// Shows the credential configuration status.
type R2DataCatalogListResponseWarehousesCredentialStatus string

const (
	R2DataCatalogListResponseWarehousesCredentialStatusPresent R2DataCatalogListResponseWarehousesCredentialStatus = "present"
	R2DataCatalogListResponseWarehousesCredentialStatusAbsent  R2DataCatalogListResponseWarehousesCredentialStatus = "absent"
)

func (r R2DataCatalogListResponseWarehousesCredentialStatus) IsKnown() bool {
	switch r {
	case R2DataCatalogListResponseWarehousesCredentialStatusPresent, R2DataCatalogListResponseWarehousesCredentialStatusAbsent:
		return true
	}
	return false
}

// Configures maintenance for the catalog.
type R2DataCatalogListResponseWarehousesMaintenanceConfig struct {
	// Configures compaction for catalog maintenance.
	Compaction R2DataCatalogListResponseWarehousesMaintenanceConfigCompaction `json:"compaction"`
	// Configures snapshot expiration settings.
	SnapshotExpiration R2DataCatalogListResponseWarehousesMaintenanceConfigSnapshotExpiration `json:"snapshot_expiration"`
	JSON               r2DataCatalogListResponseWarehousesMaintenanceConfigJSON               `json:"-"`
}

// r2DataCatalogListResponseWarehousesMaintenanceConfigJSON contains the JSON
// metadata for the struct [R2DataCatalogListResponseWarehousesMaintenanceConfig]
type r2DataCatalogListResponseWarehousesMaintenanceConfigJSON struct {
	Compaction         apijson.Field
	SnapshotExpiration apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *R2DataCatalogListResponseWarehousesMaintenanceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2DataCatalogListResponseWarehousesMaintenanceConfigJSON) RawJSON() string {
	return r.raw
}

// Configures compaction for catalog maintenance.
type R2DataCatalogListResponseWarehousesMaintenanceConfigCompaction struct {
	// Specifies the state of maintenance operations.
	State R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionState `json:"state,required"`
	// Sets the target file size for compaction in megabytes. Defaults to "128".
	TargetSizeMB R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB `json:"target_size_mb,required"`
	JSON         r2DataCatalogListResponseWarehousesMaintenanceConfigCompactionJSON         `json:"-"`
}

// r2DataCatalogListResponseWarehousesMaintenanceConfigCompactionJSON contains the
// JSON metadata for the struct
// [R2DataCatalogListResponseWarehousesMaintenanceConfigCompaction]
type r2DataCatalogListResponseWarehousesMaintenanceConfigCompactionJSON struct {
	State        apijson.Field
	TargetSizeMB apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *R2DataCatalogListResponseWarehousesMaintenanceConfigCompaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2DataCatalogListResponseWarehousesMaintenanceConfigCompactionJSON) RawJSON() string {
	return r.raw
}

// Specifies the state of maintenance operations.
type R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionState string

const (
	R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionStateEnabled  R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionState = "enabled"
	R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionStateDisabled R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionState = "disabled"
)

func (r R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionState) IsKnown() bool {
	switch r {
	case R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionStateEnabled, R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionStateDisabled:
		return true
	}
	return false
}

// Sets the target file size for compaction in megabytes. Defaults to "128".
type R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB string

const (
	R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB64  R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB = "64"
	R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB128 R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB = "128"
	R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB256 R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB = "256"
	R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB512 R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB = "512"
)

func (r R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB) IsKnown() bool {
	switch r {
	case R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB64, R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB128, R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB256, R2DataCatalogListResponseWarehousesMaintenanceConfigCompactionTargetSizeMB512:
		return true
	}
	return false
}

// Configures snapshot expiration settings.
type R2DataCatalogListResponseWarehousesMaintenanceConfigSnapshotExpiration struct {
	// Specifies the maximum age for snapshots. The system deletes snapshots older than
	// this age. Format: <number><unit> where unit is d (days), h (hours), m (minutes),
	// or s (seconds). Examples: "7d" (7 days), "48h" (48 hours), "2880m" (2,880
	// minutes). Defaults to "7d".
	MaxSnapshotAge string `json:"max_snapshot_age,required"`
	// Specifies the minimum number of snapshots to retain. Defaults to 100.
	MinSnapshotsToKeep int64 `json:"min_snapshots_to_keep,required"`
	// Specifies the state of maintenance operations.
	State R2DataCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationState `json:"state,required"`
	JSON  r2DataCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationJSON  `json:"-"`
}

// r2DataCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationJSON
// contains the JSON metadata for the struct
// [R2DataCatalogListResponseWarehousesMaintenanceConfigSnapshotExpiration]
type r2DataCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationJSON struct {
	MaxSnapshotAge     apijson.Field
	MinSnapshotsToKeep apijson.Field
	State              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *R2DataCatalogListResponseWarehousesMaintenanceConfigSnapshotExpiration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2DataCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationJSON) RawJSON() string {
	return r.raw
}

// Specifies the state of maintenance operations.
type R2DataCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationState string

const (
	R2DataCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationStateEnabled  R2DataCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationState = "enabled"
	R2DataCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationStateDisabled R2DataCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationState = "disabled"
)

func (r R2DataCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationState) IsKnown() bool {
	switch r {
	case R2DataCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationStateEnabled, R2DataCatalogListResponseWarehousesMaintenanceConfigSnapshotExpirationStateDisabled:
		return true
	}
	return false
}

// Contains response from activating an R2 bucket as a catalog.
type R2DataCatalogEnableResponse struct {
	// Use this to uniquely identify the activated catalog.
	ID string `json:"id,required" format:"uuid"`
	// Specifies the name of the activated catalog.
	Name string                          `json:"name,required"`
	JSON r2DataCatalogEnableResponseJSON `json:"-"`
}

// r2DataCatalogEnableResponseJSON contains the JSON metadata for the struct
// [R2DataCatalogEnableResponse]
type r2DataCatalogEnableResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2DataCatalogEnableResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2DataCatalogEnableResponseJSON) RawJSON() string {
	return r.raw
}

// Contains R2 Data Catalog information.
type R2DataCatalogGetResponse struct {
	// Use this to uniquely identify the catalog.
	ID string `json:"id,required" format:"uuid"`
	// Specifies the associated R2 bucket name.
	Bucket string `json:"bucket,required"`
	// Specifies the catalog name (generated from account and bucket name).
	Name string `json:"name,required"`
	// Indicates the status of the catalog.
	Status R2DataCatalogGetResponseStatus `json:"status,required"`
	// Shows the credential configuration status.
	CredentialStatus R2DataCatalogGetResponseCredentialStatus `json:"credential_status,nullable"`
	// Configures maintenance for the catalog.
	MaintenanceConfig R2DataCatalogGetResponseMaintenanceConfig `json:"maintenance_config,nullable"`
	JSON              r2DataCatalogGetResponseJSON              `json:"-"`
}

// r2DataCatalogGetResponseJSON contains the JSON metadata for the struct
// [R2DataCatalogGetResponse]
type r2DataCatalogGetResponseJSON struct {
	ID                apijson.Field
	Bucket            apijson.Field
	Name              apijson.Field
	Status            apijson.Field
	CredentialStatus  apijson.Field
	MaintenanceConfig apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *R2DataCatalogGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2DataCatalogGetResponseJSON) RawJSON() string {
	return r.raw
}

// Indicates the status of the catalog.
type R2DataCatalogGetResponseStatus string

const (
	R2DataCatalogGetResponseStatusActive   R2DataCatalogGetResponseStatus = "active"
	R2DataCatalogGetResponseStatusInactive R2DataCatalogGetResponseStatus = "inactive"
)

func (r R2DataCatalogGetResponseStatus) IsKnown() bool {
	switch r {
	case R2DataCatalogGetResponseStatusActive, R2DataCatalogGetResponseStatusInactive:
		return true
	}
	return false
}

// Shows the credential configuration status.
type R2DataCatalogGetResponseCredentialStatus string

const (
	R2DataCatalogGetResponseCredentialStatusPresent R2DataCatalogGetResponseCredentialStatus = "present"
	R2DataCatalogGetResponseCredentialStatusAbsent  R2DataCatalogGetResponseCredentialStatus = "absent"
)

func (r R2DataCatalogGetResponseCredentialStatus) IsKnown() bool {
	switch r {
	case R2DataCatalogGetResponseCredentialStatusPresent, R2DataCatalogGetResponseCredentialStatusAbsent:
		return true
	}
	return false
}

// Configures maintenance for the catalog.
type R2DataCatalogGetResponseMaintenanceConfig struct {
	// Configures compaction for catalog maintenance.
	Compaction R2DataCatalogGetResponseMaintenanceConfigCompaction `json:"compaction"`
	// Configures snapshot expiration settings.
	SnapshotExpiration R2DataCatalogGetResponseMaintenanceConfigSnapshotExpiration `json:"snapshot_expiration"`
	JSON               r2DataCatalogGetResponseMaintenanceConfigJSON               `json:"-"`
}

// r2DataCatalogGetResponseMaintenanceConfigJSON contains the JSON metadata for the
// struct [R2DataCatalogGetResponseMaintenanceConfig]
type r2DataCatalogGetResponseMaintenanceConfigJSON struct {
	Compaction         apijson.Field
	SnapshotExpiration apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *R2DataCatalogGetResponseMaintenanceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2DataCatalogGetResponseMaintenanceConfigJSON) RawJSON() string {
	return r.raw
}

// Configures compaction for catalog maintenance.
type R2DataCatalogGetResponseMaintenanceConfigCompaction struct {
	// Specifies the state of maintenance operations.
	State R2DataCatalogGetResponseMaintenanceConfigCompactionState `json:"state,required"`
	// Sets the target file size for compaction in megabytes. Defaults to "128".
	TargetSizeMB R2DataCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB `json:"target_size_mb,required"`
	JSON         r2DataCatalogGetResponseMaintenanceConfigCompactionJSON         `json:"-"`
}

// r2DataCatalogGetResponseMaintenanceConfigCompactionJSON contains the JSON
// metadata for the struct [R2DataCatalogGetResponseMaintenanceConfigCompaction]
type r2DataCatalogGetResponseMaintenanceConfigCompactionJSON struct {
	State        apijson.Field
	TargetSizeMB apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *R2DataCatalogGetResponseMaintenanceConfigCompaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2DataCatalogGetResponseMaintenanceConfigCompactionJSON) RawJSON() string {
	return r.raw
}

// Specifies the state of maintenance operations.
type R2DataCatalogGetResponseMaintenanceConfigCompactionState string

const (
	R2DataCatalogGetResponseMaintenanceConfigCompactionStateEnabled  R2DataCatalogGetResponseMaintenanceConfigCompactionState = "enabled"
	R2DataCatalogGetResponseMaintenanceConfigCompactionStateDisabled R2DataCatalogGetResponseMaintenanceConfigCompactionState = "disabled"
)

func (r R2DataCatalogGetResponseMaintenanceConfigCompactionState) IsKnown() bool {
	switch r {
	case R2DataCatalogGetResponseMaintenanceConfigCompactionStateEnabled, R2DataCatalogGetResponseMaintenanceConfigCompactionStateDisabled:
		return true
	}
	return false
}

// Sets the target file size for compaction in megabytes. Defaults to "128".
type R2DataCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB string

const (
	R2DataCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB64  R2DataCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB = "64"
	R2DataCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB128 R2DataCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB = "128"
	R2DataCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB256 R2DataCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB = "256"
	R2DataCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB512 R2DataCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB = "512"
)

func (r R2DataCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB) IsKnown() bool {
	switch r {
	case R2DataCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB64, R2DataCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB128, R2DataCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB256, R2DataCatalogGetResponseMaintenanceConfigCompactionTargetSizeMB512:
		return true
	}
	return false
}

// Configures snapshot expiration settings.
type R2DataCatalogGetResponseMaintenanceConfigSnapshotExpiration struct {
	// Specifies the maximum age for snapshots. The system deletes snapshots older than
	// this age. Format: <number><unit> where unit is d (days), h (hours), m (minutes),
	// or s (seconds). Examples: "7d" (7 days), "48h" (48 hours), "2880m" (2,880
	// minutes). Defaults to "7d".
	MaxSnapshotAge string `json:"max_snapshot_age,required"`
	// Specifies the minimum number of snapshots to retain. Defaults to 100.
	MinSnapshotsToKeep int64 `json:"min_snapshots_to_keep,required"`
	// Specifies the state of maintenance operations.
	State R2DataCatalogGetResponseMaintenanceConfigSnapshotExpirationState `json:"state,required"`
	JSON  r2DataCatalogGetResponseMaintenanceConfigSnapshotExpirationJSON  `json:"-"`
}

// r2DataCatalogGetResponseMaintenanceConfigSnapshotExpirationJSON contains the
// JSON metadata for the struct
// [R2DataCatalogGetResponseMaintenanceConfigSnapshotExpiration]
type r2DataCatalogGetResponseMaintenanceConfigSnapshotExpirationJSON struct {
	MaxSnapshotAge     apijson.Field
	MinSnapshotsToKeep apijson.Field
	State              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *R2DataCatalogGetResponseMaintenanceConfigSnapshotExpiration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2DataCatalogGetResponseMaintenanceConfigSnapshotExpirationJSON) RawJSON() string {
	return r.raw
}

// Specifies the state of maintenance operations.
type R2DataCatalogGetResponseMaintenanceConfigSnapshotExpirationState string

const (
	R2DataCatalogGetResponseMaintenanceConfigSnapshotExpirationStateEnabled  R2DataCatalogGetResponseMaintenanceConfigSnapshotExpirationState = "enabled"
	R2DataCatalogGetResponseMaintenanceConfigSnapshotExpirationStateDisabled R2DataCatalogGetResponseMaintenanceConfigSnapshotExpirationState = "disabled"
)

func (r R2DataCatalogGetResponseMaintenanceConfigSnapshotExpirationState) IsKnown() bool {
	switch r {
	case R2DataCatalogGetResponseMaintenanceConfigSnapshotExpirationStateEnabled, R2DataCatalogGetResponseMaintenanceConfigSnapshotExpirationStateDisabled:
		return true
	}
	return false
}

type R2DataCatalogListParams struct {
	// Use this to identify the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type R2DataCatalogListResponseEnvelope struct {
	// Contains errors if the API call was unsuccessful.
	Errors []R2DataCatalogListResponseEnvelopeErrors `json:"errors,required"`
	// Contains informational messages.
	Messages []R2DataCatalogListResponseEnvelopeMessages `json:"messages,required"`
	// Indicates whether the API call was successful.
	Success bool `json:"success,required"`
	// Contains the list of catalogs.
	Result R2DataCatalogListResponse             `json:"result"`
	JSON   r2DataCatalogListResponseEnvelopeJSON `json:"-"`
}

// r2DataCatalogListResponseEnvelopeJSON contains the JSON metadata for the struct
// [R2DataCatalogListResponseEnvelope]
type r2DataCatalogListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2DataCatalogListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2DataCatalogListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type R2DataCatalogListResponseEnvelopeErrors struct {
	// Specifies the error code.
	Code int64 `json:"code,required"`
	// Describes the error.
	Message string                                      `json:"message,required"`
	JSON    r2DataCatalogListResponseEnvelopeErrorsJSON `json:"-"`
}

// r2DataCatalogListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [R2DataCatalogListResponseEnvelopeErrors]
type r2DataCatalogListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2DataCatalogListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2DataCatalogListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type R2DataCatalogListResponseEnvelopeMessages struct {
	// Specifies the message code.
	Code int64 `json:"code,required"`
	// Contains the message text.
	Message string                                        `json:"message,required"`
	JSON    r2DataCatalogListResponseEnvelopeMessagesJSON `json:"-"`
}

// r2DataCatalogListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [R2DataCatalogListResponseEnvelopeMessages]
type r2DataCatalogListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2DataCatalogListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2DataCatalogListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type R2DataCatalogDisableParams struct {
	// Use this to identify the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type R2DataCatalogEnableParams struct {
	// Use this to identify the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type R2DataCatalogEnableResponseEnvelope struct {
	// Contains errors if the API call was unsuccessful.
	Errors []R2DataCatalogEnableResponseEnvelopeErrors `json:"errors,required"`
	// Contains informational messages.
	Messages []R2DataCatalogEnableResponseEnvelopeMessages `json:"messages,required"`
	// Indicates whether the API call was successful.
	Success bool `json:"success,required"`
	// Contains response from activating an R2 bucket as a catalog.
	Result R2DataCatalogEnableResponse             `json:"result"`
	JSON   r2DataCatalogEnableResponseEnvelopeJSON `json:"-"`
}

// r2DataCatalogEnableResponseEnvelopeJSON contains the JSON metadata for the
// struct [R2DataCatalogEnableResponseEnvelope]
type r2DataCatalogEnableResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2DataCatalogEnableResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2DataCatalogEnableResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type R2DataCatalogEnableResponseEnvelopeErrors struct {
	// Specifies the error code.
	Code int64 `json:"code,required"`
	// Describes the error.
	Message string                                        `json:"message,required"`
	JSON    r2DataCatalogEnableResponseEnvelopeErrorsJSON `json:"-"`
}

// r2DataCatalogEnableResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [R2DataCatalogEnableResponseEnvelopeErrors]
type r2DataCatalogEnableResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2DataCatalogEnableResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2DataCatalogEnableResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type R2DataCatalogEnableResponseEnvelopeMessages struct {
	// Specifies the message code.
	Code int64 `json:"code,required"`
	// Contains the message text.
	Message string                                          `json:"message,required"`
	JSON    r2DataCatalogEnableResponseEnvelopeMessagesJSON `json:"-"`
}

// r2DataCatalogEnableResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [R2DataCatalogEnableResponseEnvelopeMessages]
type r2DataCatalogEnableResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2DataCatalogEnableResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2DataCatalogEnableResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type R2DataCatalogGetParams struct {
	// Use this to identify the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type R2DataCatalogGetResponseEnvelope struct {
	// Contains errors if the API call was unsuccessful.
	Errors []R2DataCatalogGetResponseEnvelopeErrors `json:"errors,required"`
	// Contains informational messages.
	Messages []R2DataCatalogGetResponseEnvelopeMessages `json:"messages,required"`
	// Indicates whether the API call was successful.
	Success bool `json:"success,required"`
	// Contains R2 Data Catalog information.
	Result R2DataCatalogGetResponse             `json:"result"`
	JSON   r2DataCatalogGetResponseEnvelopeJSON `json:"-"`
}

// r2DataCatalogGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [R2DataCatalogGetResponseEnvelope]
type r2DataCatalogGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2DataCatalogGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2DataCatalogGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type R2DataCatalogGetResponseEnvelopeErrors struct {
	// Specifies the error code.
	Code int64 `json:"code,required"`
	// Describes the error.
	Message string                                     `json:"message,required"`
	JSON    r2DataCatalogGetResponseEnvelopeErrorsJSON `json:"-"`
}

// r2DataCatalogGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [R2DataCatalogGetResponseEnvelopeErrors]
type r2DataCatalogGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2DataCatalogGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2DataCatalogGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type R2DataCatalogGetResponseEnvelopeMessages struct {
	// Specifies the message code.
	Code int64 `json:"code,required"`
	// Contains the message text.
	Message string                                       `json:"message,required"`
	JSON    r2DataCatalogGetResponseEnvelopeMessagesJSON `json:"-"`
}

// r2DataCatalogGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [R2DataCatalogGetResponseEnvelopeMessages]
type r2DataCatalogGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2DataCatalogGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2DataCatalogGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
