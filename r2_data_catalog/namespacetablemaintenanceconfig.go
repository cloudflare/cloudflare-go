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

// NamespaceTableMaintenanceConfigService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNamespaceTableMaintenanceConfigService] method instead.
type NamespaceTableMaintenanceConfigService struct {
	Options []option.RequestOption
}

// NewNamespaceTableMaintenanceConfigService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewNamespaceTableMaintenanceConfigService(opts ...option.RequestOption) (r *NamespaceTableMaintenanceConfigService) {
	r = &NamespaceTableMaintenanceConfigService{}
	r.Options = opts
	return
}

// Update the maintenance configuration for a specific table. This allows you to
// enable or disable compaction and adjust target file sizes for optimization.
func (r *NamespaceTableMaintenanceConfigService) Update(ctx context.Context, bucketName string, namespace string, tableName string, params NamespaceTableMaintenanceConfigUpdateParams, opts ...option.RequestOption) (res *NamespaceTableMaintenanceConfigUpdateResponse, err error) {
	var env NamespaceTableMaintenanceConfigUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	if namespace == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	if tableName == "" {
		err = errors.New("missing required table_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2-catalog/%s/namespaces/%s/tables/%s/maintenance-configs", params.AccountID, bucketName, namespace, tableName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve the maintenance configuration for a specific table, including
// compaction settings.
func (r *NamespaceTableMaintenanceConfigService) Get(ctx context.Context, bucketName string, namespace string, tableName string, query NamespaceTableMaintenanceConfigGetParams, opts ...option.RequestOption) (res *NamespaceTableMaintenanceConfigGetResponse, err error) {
	var env NamespaceTableMaintenanceConfigGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	if namespace == "" {
		err = errors.New("missing required namespace parameter")
		return
	}
	if tableName == "" {
		err = errors.New("missing required table_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2-catalog/%s/namespaces/%s/tables/%s/maintenance-configs", query.AccountID, bucketName, namespace, tableName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Configures maintenance for the table.
type NamespaceTableMaintenanceConfigUpdateResponse struct {
	// Configures compaction settings for table optimization.
	Compaction NamespaceTableMaintenanceConfigUpdateResponseCompaction `json:"compaction"`
	JSON       namespaceTableMaintenanceConfigUpdateResponseJSON       `json:"-"`
}

// namespaceTableMaintenanceConfigUpdateResponseJSON contains the JSON metadata for
// the struct [NamespaceTableMaintenanceConfigUpdateResponse]
type namespaceTableMaintenanceConfigUpdateResponseJSON struct {
	Compaction  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceTableMaintenanceConfigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceTableMaintenanceConfigUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Configures compaction settings for table optimization.
type NamespaceTableMaintenanceConfigUpdateResponseCompaction struct {
	// Specifies the state of maintenance operations.
	State NamespaceTableMaintenanceConfigUpdateResponseCompactionState `json:"state,required"`
	// Sets the target file size for compaction in megabytes.
	TargetSizeMB NamespaceTableMaintenanceConfigUpdateResponseCompactionTargetSizeMB `json:"target_size_mb,required"`
	JSON         namespaceTableMaintenanceConfigUpdateResponseCompactionJSON         `json:"-"`
}

// namespaceTableMaintenanceConfigUpdateResponseCompactionJSON contains the JSON
// metadata for the struct
// [NamespaceTableMaintenanceConfigUpdateResponseCompaction]
type namespaceTableMaintenanceConfigUpdateResponseCompactionJSON struct {
	State        apijson.Field
	TargetSizeMB apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NamespaceTableMaintenanceConfigUpdateResponseCompaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceTableMaintenanceConfigUpdateResponseCompactionJSON) RawJSON() string {
	return r.raw
}

// Specifies the state of maintenance operations.
type NamespaceTableMaintenanceConfigUpdateResponseCompactionState string

const (
	NamespaceTableMaintenanceConfigUpdateResponseCompactionStateEnabled  NamespaceTableMaintenanceConfigUpdateResponseCompactionState = "enabled"
	NamespaceTableMaintenanceConfigUpdateResponseCompactionStateDisabled NamespaceTableMaintenanceConfigUpdateResponseCompactionState = "disabled"
)

func (r NamespaceTableMaintenanceConfigUpdateResponseCompactionState) IsKnown() bool {
	switch r {
	case NamespaceTableMaintenanceConfigUpdateResponseCompactionStateEnabled, NamespaceTableMaintenanceConfigUpdateResponseCompactionStateDisabled:
		return true
	}
	return false
}

// Sets the target file size for compaction in megabytes.
type NamespaceTableMaintenanceConfigUpdateResponseCompactionTargetSizeMB string

const (
	NamespaceTableMaintenanceConfigUpdateResponseCompactionTargetSizeMB64  NamespaceTableMaintenanceConfigUpdateResponseCompactionTargetSizeMB = "64"
	NamespaceTableMaintenanceConfigUpdateResponseCompactionTargetSizeMB128 NamespaceTableMaintenanceConfigUpdateResponseCompactionTargetSizeMB = "128"
	NamespaceTableMaintenanceConfigUpdateResponseCompactionTargetSizeMB256 NamespaceTableMaintenanceConfigUpdateResponseCompactionTargetSizeMB = "256"
	NamespaceTableMaintenanceConfigUpdateResponseCompactionTargetSizeMB512 NamespaceTableMaintenanceConfigUpdateResponseCompactionTargetSizeMB = "512"
)

func (r NamespaceTableMaintenanceConfigUpdateResponseCompactionTargetSizeMB) IsKnown() bool {
	switch r {
	case NamespaceTableMaintenanceConfigUpdateResponseCompactionTargetSizeMB64, NamespaceTableMaintenanceConfigUpdateResponseCompactionTargetSizeMB128, NamespaceTableMaintenanceConfigUpdateResponseCompactionTargetSizeMB256, NamespaceTableMaintenanceConfigUpdateResponseCompactionTargetSizeMB512:
		return true
	}
	return false
}

// Contains table maintenance configuration.
type NamespaceTableMaintenanceConfigGetResponse struct {
	// Configures maintenance for the table.
	MaintenanceConfig NamespaceTableMaintenanceConfigGetResponseMaintenanceConfig `json:"maintenance_config,required"`
	JSON              namespaceTableMaintenanceConfigGetResponseJSON              `json:"-"`
}

// namespaceTableMaintenanceConfigGetResponseJSON contains the JSON metadata for
// the struct [NamespaceTableMaintenanceConfigGetResponse]
type namespaceTableMaintenanceConfigGetResponseJSON struct {
	MaintenanceConfig apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *NamespaceTableMaintenanceConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceTableMaintenanceConfigGetResponseJSON) RawJSON() string {
	return r.raw
}

// Configures maintenance for the table.
type NamespaceTableMaintenanceConfigGetResponseMaintenanceConfig struct {
	// Configures compaction settings for table optimization.
	Compaction NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompaction `json:"compaction"`
	JSON       namespaceTableMaintenanceConfigGetResponseMaintenanceConfigJSON       `json:"-"`
}

// namespaceTableMaintenanceConfigGetResponseMaintenanceConfigJSON contains the
// JSON metadata for the struct
// [NamespaceTableMaintenanceConfigGetResponseMaintenanceConfig]
type namespaceTableMaintenanceConfigGetResponseMaintenanceConfigJSON struct {
	Compaction  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceTableMaintenanceConfigGetResponseMaintenanceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceTableMaintenanceConfigGetResponseMaintenanceConfigJSON) RawJSON() string {
	return r.raw
}

// Configures compaction settings for table optimization.
type NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompaction struct {
	// Specifies the state of maintenance operations.
	State NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionState `json:"state,required"`
	// Sets the target file size for compaction in megabytes.
	TargetSizeMB NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB `json:"target_size_mb,required"`
	JSON         namespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionJSON         `json:"-"`
}

// namespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionJSON
// contains the JSON metadata for the struct
// [NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompaction]
type namespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionJSON struct {
	State        apijson.Field
	TargetSizeMB apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionJSON) RawJSON() string {
	return r.raw
}

// Specifies the state of maintenance operations.
type NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionState string

const (
	NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionStateEnabled  NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionState = "enabled"
	NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionStateDisabled NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionState = "disabled"
)

func (r NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionState) IsKnown() bool {
	switch r {
	case NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionStateEnabled, NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionStateDisabled:
		return true
	}
	return false
}

// Sets the target file size for compaction in megabytes.
type NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB string

const (
	NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB64  NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB = "64"
	NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB128 NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB = "128"
	NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB256 NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB = "256"
	NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB512 NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB = "512"
)

func (r NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB) IsKnown() bool {
	switch r {
	case NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB64, NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB128, NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB256, NamespaceTableMaintenanceConfigGetResponseMaintenanceConfigCompactionTargetSizeMB512:
		return true
	}
	return false
}

type NamespaceTableMaintenanceConfigUpdateParams struct {
	// Use this to identify the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// Updates compaction configuration (all fields optional).
	Compaction param.Field[NamespaceTableMaintenanceConfigUpdateParamsCompaction] `json:"compaction"`
}

func (r NamespaceTableMaintenanceConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Updates compaction configuration (all fields optional).
type NamespaceTableMaintenanceConfigUpdateParamsCompaction struct {
	// Updates the state optionally.
	State param.Field[NamespaceTableMaintenanceConfigUpdateParamsCompactionState] `json:"state"`
	// Updates the target file size optionally.
	TargetSizeMB param.Field[NamespaceTableMaintenanceConfigUpdateParamsCompactionTargetSizeMB] `json:"target_size_mb"`
}

func (r NamespaceTableMaintenanceConfigUpdateParamsCompaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Updates the state optionally.
type NamespaceTableMaintenanceConfigUpdateParamsCompactionState string

const (
	NamespaceTableMaintenanceConfigUpdateParamsCompactionStateEnabled  NamespaceTableMaintenanceConfigUpdateParamsCompactionState = "enabled"
	NamespaceTableMaintenanceConfigUpdateParamsCompactionStateDisabled NamespaceTableMaintenanceConfigUpdateParamsCompactionState = "disabled"
)

func (r NamespaceTableMaintenanceConfigUpdateParamsCompactionState) IsKnown() bool {
	switch r {
	case NamespaceTableMaintenanceConfigUpdateParamsCompactionStateEnabled, NamespaceTableMaintenanceConfigUpdateParamsCompactionStateDisabled:
		return true
	}
	return false
}

// Updates the target file size optionally.
type NamespaceTableMaintenanceConfigUpdateParamsCompactionTargetSizeMB string

const (
	NamespaceTableMaintenanceConfigUpdateParamsCompactionTargetSizeMB64  NamespaceTableMaintenanceConfigUpdateParamsCompactionTargetSizeMB = "64"
	NamespaceTableMaintenanceConfigUpdateParamsCompactionTargetSizeMB128 NamespaceTableMaintenanceConfigUpdateParamsCompactionTargetSizeMB = "128"
	NamespaceTableMaintenanceConfigUpdateParamsCompactionTargetSizeMB256 NamespaceTableMaintenanceConfigUpdateParamsCompactionTargetSizeMB = "256"
	NamespaceTableMaintenanceConfigUpdateParamsCompactionTargetSizeMB512 NamespaceTableMaintenanceConfigUpdateParamsCompactionTargetSizeMB = "512"
)

func (r NamespaceTableMaintenanceConfigUpdateParamsCompactionTargetSizeMB) IsKnown() bool {
	switch r {
	case NamespaceTableMaintenanceConfigUpdateParamsCompactionTargetSizeMB64, NamespaceTableMaintenanceConfigUpdateParamsCompactionTargetSizeMB128, NamespaceTableMaintenanceConfigUpdateParamsCompactionTargetSizeMB256, NamespaceTableMaintenanceConfigUpdateParamsCompactionTargetSizeMB512:
		return true
	}
	return false
}

type NamespaceTableMaintenanceConfigUpdateResponseEnvelope struct {
	// Contains errors if the API call was unsuccessful.
	Errors []NamespaceTableMaintenanceConfigUpdateResponseEnvelopeErrors `json:"errors,required"`
	// Contains informational messages.
	Messages []NamespaceTableMaintenanceConfigUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Indicates whether the API call was successful.
	Success bool `json:"success,required"`
	// Configures maintenance for the table.
	Result NamespaceTableMaintenanceConfigUpdateResponse             `json:"result"`
	JSON   namespaceTableMaintenanceConfigUpdateResponseEnvelopeJSON `json:"-"`
}

// namespaceTableMaintenanceConfigUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [NamespaceTableMaintenanceConfigUpdateResponseEnvelope]
type namespaceTableMaintenanceConfigUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceTableMaintenanceConfigUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceTableMaintenanceConfigUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceTableMaintenanceConfigUpdateResponseEnvelopeErrors struct {
	// Specifies the error code.
	Code int64 `json:"code,required"`
	// Describes the error.
	Message string                                                          `json:"message,required"`
	JSON    namespaceTableMaintenanceConfigUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// namespaceTableMaintenanceConfigUpdateResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [NamespaceTableMaintenanceConfigUpdateResponseEnvelopeErrors]
type namespaceTableMaintenanceConfigUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceTableMaintenanceConfigUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceTableMaintenanceConfigUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NamespaceTableMaintenanceConfigUpdateResponseEnvelopeMessages struct {
	// Specifies the message code.
	Code int64 `json:"code,required"`
	// Contains the message text.
	Message string                                                            `json:"message,required"`
	JSON    namespaceTableMaintenanceConfigUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// namespaceTableMaintenanceConfigUpdateResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [NamespaceTableMaintenanceConfigUpdateResponseEnvelopeMessages]
type namespaceTableMaintenanceConfigUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceTableMaintenanceConfigUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceTableMaintenanceConfigUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type NamespaceTableMaintenanceConfigGetParams struct {
	// Use this to identify the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type NamespaceTableMaintenanceConfigGetResponseEnvelope struct {
	// Contains errors if the API call was unsuccessful.
	Errors []NamespaceTableMaintenanceConfigGetResponseEnvelopeErrors `json:"errors,required"`
	// Contains informational messages.
	Messages []NamespaceTableMaintenanceConfigGetResponseEnvelopeMessages `json:"messages,required"`
	// Indicates whether the API call was successful.
	Success bool `json:"success,required"`
	// Contains table maintenance configuration.
	Result NamespaceTableMaintenanceConfigGetResponse             `json:"result"`
	JSON   namespaceTableMaintenanceConfigGetResponseEnvelopeJSON `json:"-"`
}

// namespaceTableMaintenanceConfigGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [NamespaceTableMaintenanceConfigGetResponseEnvelope]
type namespaceTableMaintenanceConfigGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceTableMaintenanceConfigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceTableMaintenanceConfigGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceTableMaintenanceConfigGetResponseEnvelopeErrors struct {
	// Specifies the error code.
	Code int64 `json:"code,required"`
	// Describes the error.
	Message string                                                       `json:"message,required"`
	JSON    namespaceTableMaintenanceConfigGetResponseEnvelopeErrorsJSON `json:"-"`
}

// namespaceTableMaintenanceConfigGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [NamespaceTableMaintenanceConfigGetResponseEnvelopeErrors]
type namespaceTableMaintenanceConfigGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceTableMaintenanceConfigGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceTableMaintenanceConfigGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NamespaceTableMaintenanceConfigGetResponseEnvelopeMessages struct {
	// Specifies the message code.
	Code int64 `json:"code,required"`
	// Contains the message text.
	Message string                                                         `json:"message,required"`
	JSON    namespaceTableMaintenanceConfigGetResponseEnvelopeMessagesJSON `json:"-"`
}

// namespaceTableMaintenanceConfigGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [NamespaceTableMaintenanceConfigGetResponseEnvelopeMessages]
type namespaceTableMaintenanceConfigGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceTableMaintenanceConfigGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceTableMaintenanceConfigGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
