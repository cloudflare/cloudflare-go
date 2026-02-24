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

// MaintenanceConfigService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMaintenanceConfigService] method instead.
type MaintenanceConfigService struct {
	Options []option.RequestOption
}

// NewMaintenanceConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMaintenanceConfigService(opts ...option.RequestOption) (r *MaintenanceConfigService) {
	r = &MaintenanceConfigService{}
	r.Options = opts
	return
}

// Update the maintenance configuration for a catalog. This allows you to enable or
// disable compaction and adjust target file sizes for optimization.
func (r *MaintenanceConfigService) Update(ctx context.Context, bucketName string, params MaintenanceConfigUpdateParams, opts ...option.RequestOption) (res *MaintenanceConfigUpdateResponse, err error) {
	var env MaintenanceConfigUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2-catalog/%s/maintenance-configs", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve the maintenance configuration for a specific catalog, including
// compaction settings and credential status.
func (r *MaintenanceConfigService) Get(ctx context.Context, bucketName string, query MaintenanceConfigGetParams, opts ...option.RequestOption) (res *MaintenanceConfigGetResponse, err error) {
	var env MaintenanceConfigGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2-catalog/%s/maintenance-configs", query.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Configures maintenance for the catalog.
type MaintenanceConfigUpdateResponse struct {
	// Configures compaction for catalog maintenance.
	Compaction MaintenanceConfigUpdateResponseCompaction `json:"compaction"`
	// Configures snapshot expiration settings.
	SnapshotExpiration MaintenanceConfigUpdateResponseSnapshotExpiration `json:"snapshot_expiration"`
	JSON               maintenanceConfigUpdateResponseJSON               `json:"-"`
}

// maintenanceConfigUpdateResponseJSON contains the JSON metadata for the struct
// [MaintenanceConfigUpdateResponse]
type maintenanceConfigUpdateResponseJSON struct {
	Compaction         apijson.Field
	SnapshotExpiration apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MaintenanceConfigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maintenanceConfigUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Configures compaction for catalog maintenance.
type MaintenanceConfigUpdateResponseCompaction struct {
	// Specifies the state of maintenance operations.
	State MaintenanceConfigUpdateResponseCompactionState `json:"state,required"`
	// Sets the target file size for compaction in megabytes. Defaults to "128".
	TargetSizeMB MaintenanceConfigUpdateResponseCompactionTargetSizeMB `json:"target_size_mb,required"`
	JSON         maintenanceConfigUpdateResponseCompactionJSON         `json:"-"`
}

// maintenanceConfigUpdateResponseCompactionJSON contains the JSON metadata for the
// struct [MaintenanceConfigUpdateResponseCompaction]
type maintenanceConfigUpdateResponseCompactionJSON struct {
	State        apijson.Field
	TargetSizeMB apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MaintenanceConfigUpdateResponseCompaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maintenanceConfigUpdateResponseCompactionJSON) RawJSON() string {
	return r.raw
}

// Specifies the state of maintenance operations.
type MaintenanceConfigUpdateResponseCompactionState string

const (
	MaintenanceConfigUpdateResponseCompactionStateEnabled  MaintenanceConfigUpdateResponseCompactionState = "enabled"
	MaintenanceConfigUpdateResponseCompactionStateDisabled MaintenanceConfigUpdateResponseCompactionState = "disabled"
)

func (r MaintenanceConfigUpdateResponseCompactionState) IsKnown() bool {
	switch r {
	case MaintenanceConfigUpdateResponseCompactionStateEnabled, MaintenanceConfigUpdateResponseCompactionStateDisabled:
		return true
	}
	return false
}

// Sets the target file size for compaction in megabytes. Defaults to "128".
type MaintenanceConfigUpdateResponseCompactionTargetSizeMB string

const (
	MaintenanceConfigUpdateResponseCompactionTargetSizeMB64  MaintenanceConfigUpdateResponseCompactionTargetSizeMB = "64"
	MaintenanceConfigUpdateResponseCompactionTargetSizeMB128 MaintenanceConfigUpdateResponseCompactionTargetSizeMB = "128"
	MaintenanceConfigUpdateResponseCompactionTargetSizeMB256 MaintenanceConfigUpdateResponseCompactionTargetSizeMB = "256"
	MaintenanceConfigUpdateResponseCompactionTargetSizeMB512 MaintenanceConfigUpdateResponseCompactionTargetSizeMB = "512"
)

func (r MaintenanceConfigUpdateResponseCompactionTargetSizeMB) IsKnown() bool {
	switch r {
	case MaintenanceConfigUpdateResponseCompactionTargetSizeMB64, MaintenanceConfigUpdateResponseCompactionTargetSizeMB128, MaintenanceConfigUpdateResponseCompactionTargetSizeMB256, MaintenanceConfigUpdateResponseCompactionTargetSizeMB512:
		return true
	}
	return false
}

// Configures snapshot expiration settings.
type MaintenanceConfigUpdateResponseSnapshotExpiration struct {
	// Specifies the maximum age for snapshots. The system deletes snapshots older than
	// this age. Format: <number><unit> where unit is d (days), h (hours), m (minutes),
	// or s (seconds). Examples: "7d" (7 days), "48h" (48 hours), "2880m" (2,880
	// minutes). Defaults to "7d".
	MaxSnapshotAge string `json:"max_snapshot_age,required"`
	// Specifies the minimum number of snapshots to retain. Defaults to 100.
	MinSnapshotsToKeep int64 `json:"min_snapshots_to_keep,required"`
	// Specifies the state of maintenance operations.
	State MaintenanceConfigUpdateResponseSnapshotExpirationState `json:"state,required"`
	JSON  maintenanceConfigUpdateResponseSnapshotExpirationJSON  `json:"-"`
}

// maintenanceConfigUpdateResponseSnapshotExpirationJSON contains the JSON metadata
// for the struct [MaintenanceConfigUpdateResponseSnapshotExpiration]
type maintenanceConfigUpdateResponseSnapshotExpirationJSON struct {
	MaxSnapshotAge     apijson.Field
	MinSnapshotsToKeep apijson.Field
	State              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MaintenanceConfigUpdateResponseSnapshotExpiration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maintenanceConfigUpdateResponseSnapshotExpirationJSON) RawJSON() string {
	return r.raw
}

// Specifies the state of maintenance operations.
type MaintenanceConfigUpdateResponseSnapshotExpirationState string

const (
	MaintenanceConfigUpdateResponseSnapshotExpirationStateEnabled  MaintenanceConfigUpdateResponseSnapshotExpirationState = "enabled"
	MaintenanceConfigUpdateResponseSnapshotExpirationStateDisabled MaintenanceConfigUpdateResponseSnapshotExpirationState = "disabled"
)

func (r MaintenanceConfigUpdateResponseSnapshotExpirationState) IsKnown() bool {
	switch r {
	case MaintenanceConfigUpdateResponseSnapshotExpirationStateEnabled, MaintenanceConfigUpdateResponseSnapshotExpirationStateDisabled:
		return true
	}
	return false
}

// Contains maintenance configuration and credential status.
type MaintenanceConfigGetResponse struct {
	// Shows the credential configuration status.
	CredentialStatus MaintenanceConfigGetResponseCredentialStatus `json:"credential_status,required"`
	// Configures maintenance for the catalog.
	MaintenanceConfig MaintenanceConfigGetResponseMaintenanceConfig `json:"maintenance_config,required"`
	JSON              maintenanceConfigGetResponseJSON              `json:"-"`
}

// maintenanceConfigGetResponseJSON contains the JSON metadata for the struct
// [MaintenanceConfigGetResponse]
type maintenanceConfigGetResponseJSON struct {
	CredentialStatus  apijson.Field
	MaintenanceConfig apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *MaintenanceConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maintenanceConfigGetResponseJSON) RawJSON() string {
	return r.raw
}

// Shows the credential configuration status.
type MaintenanceConfigGetResponseCredentialStatus string

const (
	MaintenanceConfigGetResponseCredentialStatusPresent MaintenanceConfigGetResponseCredentialStatus = "present"
	MaintenanceConfigGetResponseCredentialStatusAbsent  MaintenanceConfigGetResponseCredentialStatus = "absent"
)

func (r MaintenanceConfigGetResponseCredentialStatus) IsKnown() bool {
	switch r {
	case MaintenanceConfigGetResponseCredentialStatusPresent, MaintenanceConfigGetResponseCredentialStatusAbsent:
		return true
	}
	return false
}

// Configures maintenance for the catalog.
type MaintenanceConfigGetResponseMaintenanceConfig struct {
	// Configures compaction for catalog maintenance.
	Compaction MaintenanceConfigGetResponseMaintenanceConfigCompaction `json:"compaction"`
	// Configures snapshot expiration settings.
	SnapshotExpiration MaintenanceConfigGetResponseMaintenanceConfigSnapshotExpiration `json:"snapshot_expiration"`
	JSON               maintenanceConfigGetResponseMaintenanceConfigJSON               `json:"-"`
}

// maintenanceConfigGetResponseMaintenanceConfigJSON contains the JSON metadata for
// the struct [MaintenanceConfigGetResponseMaintenanceConfig]
type maintenanceConfigGetResponseMaintenanceConfigJSON struct {
	Compaction         apijson.Field
	SnapshotExpiration apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MaintenanceConfigGetResponseMaintenanceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maintenanceConfigGetResponseMaintenanceConfigJSON) RawJSON() string {
	return r.raw
}

// Configures compaction for catalog maintenance.
type MaintenanceConfigGetResponseMaintenanceConfigCompaction struct {
	// Specifies the state of maintenance operations.
	State MaintenanceConfigGetResponseMaintenanceConfigCompactionState `json:"state,required"`
	// Sets the target file size for compaction in megabytes. Defaults to "128".
	TargetSizeMB MaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB `json:"target_size_mb,required"`
	JSON         maintenanceConfigGetResponseMaintenanceConfigCompactionJSON         `json:"-"`
}

// maintenanceConfigGetResponseMaintenanceConfigCompactionJSON contains the JSON
// metadata for the struct
// [MaintenanceConfigGetResponseMaintenanceConfigCompaction]
type maintenanceConfigGetResponseMaintenanceConfigCompactionJSON struct {
	State        apijson.Field
	TargetSizeMB apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MaintenanceConfigGetResponseMaintenanceConfigCompaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maintenanceConfigGetResponseMaintenanceConfigCompactionJSON) RawJSON() string {
	return r.raw
}

// Specifies the state of maintenance operations.
type MaintenanceConfigGetResponseMaintenanceConfigCompactionState string

const (
	MaintenanceConfigGetResponseMaintenanceConfigCompactionStateEnabled  MaintenanceConfigGetResponseMaintenanceConfigCompactionState = "enabled"
	MaintenanceConfigGetResponseMaintenanceConfigCompactionStateDisabled MaintenanceConfigGetResponseMaintenanceConfigCompactionState = "disabled"
)

func (r MaintenanceConfigGetResponseMaintenanceConfigCompactionState) IsKnown() bool {
	switch r {
	case MaintenanceConfigGetResponseMaintenanceConfigCompactionStateEnabled, MaintenanceConfigGetResponseMaintenanceConfigCompactionStateDisabled:
		return true
	}
	return false
}

// Sets the target file size for compaction in megabytes. Defaults to "128".
type MaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB string

const (
	MaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB64  MaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB = "64"
	MaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB128 MaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB = "128"
	MaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB256 MaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB = "256"
	MaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB512 MaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB = "512"
)

func (r MaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB) IsKnown() bool {
	switch r {
	case MaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB64, MaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB128, MaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB256, MaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB512:
		return true
	}
	return false
}

// Configures snapshot expiration settings.
type MaintenanceConfigGetResponseMaintenanceConfigSnapshotExpiration struct {
	// Specifies the maximum age for snapshots. The system deletes snapshots older than
	// this age. Format: <number><unit> where unit is d (days), h (hours), m (minutes),
	// or s (seconds). Examples: "7d" (7 days), "48h" (48 hours), "2880m" (2,880
	// minutes). Defaults to "7d".
	MaxSnapshotAge string `json:"max_snapshot_age,required"`
	// Specifies the minimum number of snapshots to retain. Defaults to 100.
	MinSnapshotsToKeep int64 `json:"min_snapshots_to_keep,required"`
	// Specifies the state of maintenance operations.
	State MaintenanceConfigGetResponseMaintenanceConfigSnapshotExpirationState `json:"state,required"`
	JSON  maintenanceConfigGetResponseMaintenanceConfigSnapshotExpirationJSON  `json:"-"`
}

// maintenanceConfigGetResponseMaintenanceConfigSnapshotExpirationJSON contains the
// JSON metadata for the struct
// [MaintenanceConfigGetResponseMaintenanceConfigSnapshotExpiration]
type maintenanceConfigGetResponseMaintenanceConfigSnapshotExpirationJSON struct {
	MaxSnapshotAge     apijson.Field
	MinSnapshotsToKeep apijson.Field
	State              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MaintenanceConfigGetResponseMaintenanceConfigSnapshotExpiration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maintenanceConfigGetResponseMaintenanceConfigSnapshotExpirationJSON) RawJSON() string {
	return r.raw
}

// Specifies the state of maintenance operations.
type MaintenanceConfigGetResponseMaintenanceConfigSnapshotExpirationState string

const (
	MaintenanceConfigGetResponseMaintenanceConfigSnapshotExpirationStateEnabled  MaintenanceConfigGetResponseMaintenanceConfigSnapshotExpirationState = "enabled"
	MaintenanceConfigGetResponseMaintenanceConfigSnapshotExpirationStateDisabled MaintenanceConfigGetResponseMaintenanceConfigSnapshotExpirationState = "disabled"
)

func (r MaintenanceConfigGetResponseMaintenanceConfigSnapshotExpirationState) IsKnown() bool {
	switch r {
	case MaintenanceConfigGetResponseMaintenanceConfigSnapshotExpirationStateEnabled, MaintenanceConfigGetResponseMaintenanceConfigSnapshotExpirationStateDisabled:
		return true
	}
	return false
}

type MaintenanceConfigUpdateParams struct {
	// Use this to identify the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// Updates compaction configuration (all fields optional).
	Compaction param.Field[MaintenanceConfigUpdateParamsCompaction] `json:"compaction"`
	// Updates snapshot expiration configuration (all fields optional).
	SnapshotExpiration param.Field[MaintenanceConfigUpdateParamsSnapshotExpiration] `json:"snapshot_expiration"`
}

func (r MaintenanceConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Updates compaction configuration (all fields optional).
type MaintenanceConfigUpdateParamsCompaction struct {
	// Updates the state optionally.
	State param.Field[MaintenanceConfigUpdateParamsCompactionState] `json:"state"`
	// Updates the target file size optionally.
	TargetSizeMB param.Field[MaintenanceConfigUpdateParamsCompactionTargetSizeMB] `json:"target_size_mb"`
}

func (r MaintenanceConfigUpdateParamsCompaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Updates the state optionally.
type MaintenanceConfigUpdateParamsCompactionState string

const (
	MaintenanceConfigUpdateParamsCompactionStateEnabled  MaintenanceConfigUpdateParamsCompactionState = "enabled"
	MaintenanceConfigUpdateParamsCompactionStateDisabled MaintenanceConfigUpdateParamsCompactionState = "disabled"
)

func (r MaintenanceConfigUpdateParamsCompactionState) IsKnown() bool {
	switch r {
	case MaintenanceConfigUpdateParamsCompactionStateEnabled, MaintenanceConfigUpdateParamsCompactionStateDisabled:
		return true
	}
	return false
}

// Updates the target file size optionally.
type MaintenanceConfigUpdateParamsCompactionTargetSizeMB string

const (
	MaintenanceConfigUpdateParamsCompactionTargetSizeMB64  MaintenanceConfigUpdateParamsCompactionTargetSizeMB = "64"
	MaintenanceConfigUpdateParamsCompactionTargetSizeMB128 MaintenanceConfigUpdateParamsCompactionTargetSizeMB = "128"
	MaintenanceConfigUpdateParamsCompactionTargetSizeMB256 MaintenanceConfigUpdateParamsCompactionTargetSizeMB = "256"
	MaintenanceConfigUpdateParamsCompactionTargetSizeMB512 MaintenanceConfigUpdateParamsCompactionTargetSizeMB = "512"
)

func (r MaintenanceConfigUpdateParamsCompactionTargetSizeMB) IsKnown() bool {
	switch r {
	case MaintenanceConfigUpdateParamsCompactionTargetSizeMB64, MaintenanceConfigUpdateParamsCompactionTargetSizeMB128, MaintenanceConfigUpdateParamsCompactionTargetSizeMB256, MaintenanceConfigUpdateParamsCompactionTargetSizeMB512:
		return true
	}
	return false
}

// Updates snapshot expiration configuration (all fields optional).
type MaintenanceConfigUpdateParamsSnapshotExpiration struct {
	// Updates the maximum age for snapshots optionally.
	MaxSnapshotAge param.Field[string] `json:"max_snapshot_age"`
	// Updates the minimum number of snapshots to retain optionally.
	MinSnapshotsToKeep param.Field[int64] `json:"min_snapshots_to_keep"`
	// Updates the state optionally.
	State param.Field[MaintenanceConfigUpdateParamsSnapshotExpirationState] `json:"state"`
}

func (r MaintenanceConfigUpdateParamsSnapshotExpiration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Updates the state optionally.
type MaintenanceConfigUpdateParamsSnapshotExpirationState string

const (
	MaintenanceConfigUpdateParamsSnapshotExpirationStateEnabled  MaintenanceConfigUpdateParamsSnapshotExpirationState = "enabled"
	MaintenanceConfigUpdateParamsSnapshotExpirationStateDisabled MaintenanceConfigUpdateParamsSnapshotExpirationState = "disabled"
)

func (r MaintenanceConfigUpdateParamsSnapshotExpirationState) IsKnown() bool {
	switch r {
	case MaintenanceConfigUpdateParamsSnapshotExpirationStateEnabled, MaintenanceConfigUpdateParamsSnapshotExpirationStateDisabled:
		return true
	}
	return false
}

type MaintenanceConfigUpdateResponseEnvelope struct {
	// Contains errors if the API call was unsuccessful.
	Errors []MaintenanceConfigUpdateResponseEnvelopeErrors `json:"errors,required"`
	// Contains informational messages.
	Messages []MaintenanceConfigUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Indicates whether the API call was successful.
	Success bool `json:"success,required"`
	// Configures maintenance for the catalog.
	Result MaintenanceConfigUpdateResponse             `json:"result"`
	JSON   maintenanceConfigUpdateResponseEnvelopeJSON `json:"-"`
}

// maintenanceConfigUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [MaintenanceConfigUpdateResponseEnvelope]
type maintenanceConfigUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MaintenanceConfigUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maintenanceConfigUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MaintenanceConfigUpdateResponseEnvelopeErrors struct {
	// Specifies the error code.
	Code int64 `json:"code,required"`
	// Describes the error.
	Message string                                            `json:"message,required"`
	JSON    maintenanceConfigUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// maintenanceConfigUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MaintenanceConfigUpdateResponseEnvelopeErrors]
type maintenanceConfigUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MaintenanceConfigUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maintenanceConfigUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MaintenanceConfigUpdateResponseEnvelopeMessages struct {
	// Specifies the message code.
	Code int64 `json:"code,required"`
	// Contains the message text.
	Message string                                              `json:"message,required"`
	JSON    maintenanceConfigUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// maintenanceConfigUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [MaintenanceConfigUpdateResponseEnvelopeMessages]
type maintenanceConfigUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MaintenanceConfigUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maintenanceConfigUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type MaintenanceConfigGetParams struct {
	// Use this to identify the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type MaintenanceConfigGetResponseEnvelope struct {
	// Contains errors if the API call was unsuccessful.
	Errors []MaintenanceConfigGetResponseEnvelopeErrors `json:"errors,required"`
	// Contains informational messages.
	Messages []MaintenanceConfigGetResponseEnvelopeMessages `json:"messages,required"`
	// Indicates whether the API call was successful.
	Success bool `json:"success,required"`
	// Contains maintenance configuration and credential status.
	Result MaintenanceConfigGetResponse             `json:"result"`
	JSON   maintenanceConfigGetResponseEnvelopeJSON `json:"-"`
}

// maintenanceConfigGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [MaintenanceConfigGetResponseEnvelope]
type maintenanceConfigGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MaintenanceConfigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maintenanceConfigGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MaintenanceConfigGetResponseEnvelopeErrors struct {
	// Specifies the error code.
	Code int64 `json:"code,required"`
	// Describes the error.
	Message string                                         `json:"message,required"`
	JSON    maintenanceConfigGetResponseEnvelopeErrorsJSON `json:"-"`
}

// maintenanceConfigGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MaintenanceConfigGetResponseEnvelopeErrors]
type maintenanceConfigGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MaintenanceConfigGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maintenanceConfigGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MaintenanceConfigGetResponseEnvelopeMessages struct {
	// Specifies the message code.
	Code int64 `json:"code,required"`
	// Contains the message text.
	Message string                                           `json:"message,required"`
	JSON    maintenanceConfigGetResponseEnvelopeMessagesJSON `json:"-"`
}

// maintenanceConfigGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MaintenanceConfigGetResponseEnvelopeMessages]
type maintenanceConfigGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MaintenanceConfigGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r maintenanceConfigGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
