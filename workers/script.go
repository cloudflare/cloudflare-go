// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// ScriptService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptService] method instead.
type ScriptService struct {
	Options                  []option.RequestOption
	Assets                   *ScriptAssetService
	Subdomain                *ScriptSubdomainService
	Schedules                *ScriptScheduleService
	Tail                     *ScriptTailService
	Content                  *ScriptContentService
	Settings                 *ScriptSettingService
	Deployments              *ScriptDeploymentService
	Versions                 *ScriptVersionService
	Secrets                  *ScriptSecretService
	ScriptAndVersionSettings *ScriptScriptAndVersionSettingService
}

// NewScriptService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScriptService(opts ...option.RequestOption) (r *ScriptService) {
	r = &ScriptService{}
	r.Options = opts
	r.Assets = NewScriptAssetService(opts...)
	r.Subdomain = NewScriptSubdomainService(opts...)
	r.Schedules = NewScriptScheduleService(opts...)
	r.Tail = NewScriptTailService(opts...)
	r.Content = NewScriptContentService(opts...)
	r.Settings = NewScriptSettingService(opts...)
	r.Deployments = NewScriptDeploymentService(opts...)
	r.Versions = NewScriptVersionService(opts...)
	r.Secrets = NewScriptSecretService(opts...)
	r.ScriptAndVersionSettings = NewScriptScriptAndVersionSettingService(opts...)
	return
}

// Upload a worker module. You can find more about the multipart metadata on our
// docs:
// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.
func (r *ScriptService) Update(ctx context.Context, scriptName string, params ScriptUpdateParams, opts ...option.RequestOption) (res *ScriptUpdateResponse, err error) {
	var env ScriptUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of uploaded workers.
func (r *ScriptService) List(ctx context.Context, params ScriptListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Script], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts", params.AccountID)
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

// Fetch a list of uploaded workers.
func (r *ScriptService) ListAutoPaging(ctx context.Context, params ScriptListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Script] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Delete your worker. This call has no response body on a successful delete.
func (r *ScriptService) Delete(ctx context.Context, scriptName string, params ScriptDeleteParams, opts ...option.RequestOption) (res *ScriptDeleteResponse, err error) {
	var env ScriptDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch raw script content for your worker. Note this is the original script
// content, not JSON encoded.
func (r *ScriptService) Get(ctx context.Context, scriptName string, query ScriptGetParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/javascript")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Script struct {
	// The id of the script in the Workers system. Usually the script name.
	ID string `json:"id"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Hashed script content, can be used in a If-None-Match header when updating.
	Etag string `json:"etag"`
	// Whether a Worker contains assets.
	HasAssets bool `json:"has_assets"`
	// Whether a Worker contains modules.
	HasModules bool `json:"has_modules"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement ScriptPlacement `json:"placement"`
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	//
	// Deprecated: deprecated
	PlacementMode ScriptPlacementMode `json:"placement_mode"`
	// Status of
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	//
	// Deprecated: deprecated
	PlacementStatus ScriptPlacementStatus `json:"placement_status"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ConsumerScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel ScriptUsageModel `json:"usage_model"`
	JSON       scriptJSON       `json:"-"`
}

// scriptJSON contains the JSON metadata for the struct [Script]
type scriptJSON struct {
	ID              apijson.Field
	CreatedOn       apijson.Field
	Etag            apijson.Field
	HasAssets       apijson.Field
	HasModules      apijson.Field
	Logpush         apijson.Field
	ModifiedOn      apijson.Field
	Placement       apijson.Field
	PlacementMode   apijson.Field
	PlacementStatus apijson.Field
	TailConsumers   apijson.Field
	UsageModel      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Script) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptJSON) RawJSON() string {
	return r.raw
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptPlacement struct {
	// The last time the script was analyzed for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	LastAnalyzedAt time.Time `json:"last_analyzed_at" format:"date-time"`
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode ScriptPlacementMode `json:"mode"`
	// Status of
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Status ScriptPlacementStatus `json:"status"`
	JSON   scriptPlacementJSON   `json:"-"`
}

// scriptPlacementJSON contains the JSON metadata for the struct [ScriptPlacement]
type scriptPlacementJSON struct {
	LastAnalyzedAt apijson.Field
	Mode           apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ScriptPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptPlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptPlacementMode string

const (
	ScriptPlacementModeSmart ScriptPlacementMode = "smart"
)

func (r ScriptPlacementMode) IsKnown() bool {
	switch r {
	case ScriptPlacementModeSmart:
		return true
	}
	return false
}

// Status of
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptPlacementStatus string

const (
	ScriptPlacementStatusSuccess                 ScriptPlacementStatus = "SUCCESS"
	ScriptPlacementStatusUnsupportedApplication  ScriptPlacementStatus = "UNSUPPORTED_APPLICATION"
	ScriptPlacementStatusInsufficientInvocations ScriptPlacementStatus = "INSUFFICIENT_INVOCATIONS"
)

func (r ScriptPlacementStatus) IsKnown() bool {
	switch r {
	case ScriptPlacementStatusSuccess, ScriptPlacementStatusUnsupportedApplication, ScriptPlacementStatusInsufficientInvocations:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type ScriptUsageModel string

const (
	ScriptUsageModelStandard ScriptUsageModel = "standard"
)

func (r ScriptUsageModel) IsKnown() bool {
	switch r {
	case ScriptUsageModelStandard:
		return true
	}
	return false
}

type ScriptSetting struct {
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Observability settings for the Worker.
	Observability ScriptSettingObservability `json:"observability,nullable"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ConsumerScript  `json:"tail_consumers,nullable"`
	JSON          scriptSettingJSON `json:"-"`
}

// scriptSettingJSON contains the JSON metadata for the struct [ScriptSetting]
type scriptSettingJSON struct {
	Logpush       apijson.Field
	Observability apijson.Field
	TailConsumers apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptSetting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingJSON) RawJSON() string {
	return r.raw
}

// Observability settings for the Worker.
type ScriptSettingObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled bool `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate float64 `json:"head_sampling_rate,nullable"`
	// Log settings for the Worker.
	Logs ScriptSettingObservabilityLogs `json:"logs,nullable"`
	JSON scriptSettingObservabilityJSON `json:"-"`
}

// scriptSettingObservabilityJSON contains the JSON metadata for the struct
// [ScriptSettingObservability]
type scriptSettingObservabilityJSON struct {
	Enabled          apijson.Field
	HeadSamplingRate apijson.Field
	Logs             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSettingObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingObservabilityJSON) RawJSON() string {
	return r.raw
}

// Log settings for the Worker.
type ScriptSettingObservabilityLogs struct {
	// Whether logs are enabled for the Worker.
	Enabled bool `json:"enabled,required"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs bool `json:"invocation_logs,required"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate float64                            `json:"head_sampling_rate,nullable"`
	JSON             scriptSettingObservabilityLogsJSON `json:"-"`
}

// scriptSettingObservabilityLogsJSON contains the JSON metadata for the struct
// [ScriptSettingObservabilityLogs]
type scriptSettingObservabilityLogsJSON struct {
	Enabled          apijson.Field
	InvocationLogs   apijson.Field
	HeadSamplingRate apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSettingObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingParam struct {
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Observability settings for the Worker.
	Observability param.Field[ScriptSettingObservabilityParam] `json:"observability"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]ConsumerScriptParam] `json:"tail_consumers"`
}

func (r ScriptSettingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Observability settings for the Worker.
type ScriptSettingObservabilityParam struct {
	// Whether observability is enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
	// Log settings for the Worker.
	Logs param.Field[ScriptSettingObservabilityLogsParam] `json:"logs"`
}

func (r ScriptSettingObservabilityParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Log settings for the Worker.
type ScriptSettingObservabilityLogsParam struct {
	// Whether logs are enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs param.Field[bool] `json:"invocation_logs,required"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
}

func (r ScriptSettingObservabilityLogsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptUpdateResponse struct {
	StartupTimeMs int64 `json:"startup_time_ms,required"`
	// The id of the script in the Workers system. Usually the script name.
	ID string `json:"id"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Hashed script content, can be used in a If-None-Match header when updating.
	Etag string `json:"etag"`
	// Whether a Worker contains assets.
	HasAssets bool `json:"has_assets"`
	// Whether a Worker contains modules.
	HasModules bool `json:"has_modules"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement ScriptUpdateResponsePlacement `json:"placement"`
	// Deprecated: deprecated
	PlacementMode ScriptUpdateResponsePlacementMode `json:"placement_mode"`
	// Deprecated: deprecated
	PlacementStatus ScriptUpdateResponsePlacementStatus `json:"placement_status"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ConsumerScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel ScriptUpdateResponseUsageModel `json:"usage_model"`
	JSON       scriptUpdateResponseJSON       `json:"-"`
}

// scriptUpdateResponseJSON contains the JSON metadata for the struct
// [ScriptUpdateResponse]
type scriptUpdateResponseJSON struct {
	StartupTimeMs   apijson.Field
	ID              apijson.Field
	CreatedOn       apijson.Field
	Etag            apijson.Field
	HasAssets       apijson.Field
	HasModules      apijson.Field
	Logpush         apijson.Field
	ModifiedOn      apijson.Field
	Placement       apijson.Field
	PlacementMode   apijson.Field
	PlacementStatus apijson.Field
	TailConsumers   apijson.Field
	UsageModel      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScriptUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptUpdateResponsePlacement struct {
	// The last time the script was analyzed for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	LastAnalyzedAt time.Time `json:"last_analyzed_at" format:"date-time"`
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode ScriptUpdateResponsePlacementMode `json:"mode"`
	// Status of
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Status ScriptUpdateResponsePlacementStatus `json:"status"`
	JSON   scriptUpdateResponsePlacementJSON   `json:"-"`
}

// scriptUpdateResponsePlacementJSON contains the JSON metadata for the struct
// [ScriptUpdateResponsePlacement]
type scriptUpdateResponsePlacementJSON struct {
	LastAnalyzedAt apijson.Field
	Mode           apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ScriptUpdateResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUpdateResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptUpdateResponsePlacementMode string

const (
	ScriptUpdateResponsePlacementModeSmart ScriptUpdateResponsePlacementMode = "smart"
)

func (r ScriptUpdateResponsePlacementMode) IsKnown() bool {
	switch r {
	case ScriptUpdateResponsePlacementModeSmart:
		return true
	}
	return false
}

// Status of
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptUpdateResponsePlacementStatus string

const (
	ScriptUpdateResponsePlacementStatusSuccess                 ScriptUpdateResponsePlacementStatus = "SUCCESS"
	ScriptUpdateResponsePlacementStatusUnsupportedApplication  ScriptUpdateResponsePlacementStatus = "UNSUPPORTED_APPLICATION"
	ScriptUpdateResponsePlacementStatusInsufficientInvocations ScriptUpdateResponsePlacementStatus = "INSUFFICIENT_INVOCATIONS"
)

func (r ScriptUpdateResponsePlacementStatus) IsKnown() bool {
	switch r {
	case ScriptUpdateResponsePlacementStatusSuccess, ScriptUpdateResponsePlacementStatusUnsupportedApplication, ScriptUpdateResponsePlacementStatusInsufficientInvocations:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type ScriptUpdateResponseUsageModel string

const (
	ScriptUpdateResponseUsageModelStandard ScriptUpdateResponseUsageModel = "standard"
)

func (r ScriptUpdateResponseUsageModel) IsKnown() bool {
	switch r {
	case ScriptUpdateResponseUsageModelStandard:
		return true
	}
	return false
}

type ScriptDeleteResponse = interface{}

type ScriptUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// JSON-encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[ScriptUpdateParamsMetadataUnion] `json:"metadata,required"`
	// An array of modules (often JavaScript files) comprising a Worker script. At
	// least one module must be present and referenced in the metadata as `main_module`
	// or `body_part` by filename.<br/>Possible Content-Type(s) are:
	// `application/javascript+module`, `text/javascript+module`,
	// `application/javascript`, `text/javascript`, `text/x-python`,
	// `text/x-python-requirement`, `application/wasm`, `text/plain`,
	// `application/octet-stream`, `application/source-map`.
	Files param.Field[[]io.Reader] `json:"files" format:"binary"`
}

func (r ScriptUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// JSON-encoded metadata about the uploaded parts and Worker configuration.
type ScriptUpdateParamsMetadata struct {
	Assets   param.Field[interface{}] `json:"assets"`
	Bindings param.Field[interface{}] `json:"bindings"`
	// Name of the uploaded file that contains the Worker script (e.g. the file adding
	// a listener to the `fetch` event). Indicates a `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate  param.Field[string]      `json:"compatibility_date"`
	CompatibilityFlags param.Field[interface{}] `json:"compatibility_flags"`
	// Retain assets which exist for a previously uploaded Worker version; used in lieu
	// of providing a completion token.
	KeepAssets   param.Field[bool]        `json:"keep_assets"`
	KeepBindings param.Field[interface{}] `json:"keep_bindings"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Name of the uploaded file that contains the main module (e.g. the file exporting
	// a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule    param.Field[string]      `json:"main_module"`
	Migrations    param.Field[interface{}] `json:"migrations"`
	Observability param.Field[interface{}] `json:"observability"`
	Placement     param.Field[interface{}] `json:"placement"`
	Tags          param.Field[interface{}] `json:"tags"`
	TailConsumers param.Field[interface{}] `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[ScriptUpdateParamsMetadataUsageModel] `json:"usage_model"`
}

func (r ScriptUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadata) implementsScriptUpdateParamsMetadataUnion() {}

// JSON-encoded metadata about the uploaded parts and Worker configuration.
//
// Satisfied by [workers.ScriptUpdateParamsMetadataObject],
// [workers.ScriptUpdateParamsMetadataObject], [ScriptUpdateParamsMetadata].
type ScriptUpdateParamsMetadataUnion interface {
	implementsScriptUpdateParamsMetadataUnion()
}

type ScriptUpdateParamsMetadataObject struct {
	// Name of the uploaded file that contains the main module (e.g. the file exporting
	// a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module,required"`
	// Configuration for assets within a Worker.
	Assets param.Field[ScriptUpdateParamsMetadataObjectAssets] `json:"assets"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings param.Field[[]ScriptUpdateParamsMetadataObjectBindingUnion] `json:"bindings"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Retain assets which exist for a previously uploaded Worker version; used in lieu
	// of providing a completion token.
	KeepAssets param.Field[bool] `json:"keep_assets"`
	// List of binding types to keep from previous_upload.
	KeepBindings param.Field[[]string] `json:"keep_bindings"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[ScriptUpdateParamsMetadataObjectMigrationsUnion] `json:"migrations"`
	// Observability settings for the Worker.
	Observability param.Field[ScriptUpdateParamsMetadataObjectObservability] `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement param.Field[ScriptUpdateParamsMetadataObjectPlacement] `json:"placement"`
	// List of strings to use as tags for this Worker.
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]ConsumerScriptParam] `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[ScriptUpdateParamsMetadataObjectUsageModel] `json:"usage_model"`
}

func (r ScriptUpdateParamsMetadataObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObject) implementsScriptUpdateParamsMetadataUnion() {}

// Configuration for assets within a Worker.
type ScriptUpdateParamsMetadataObjectAssets struct {
	// Configuration for assets within a Worker.
	Config param.Field[ScriptUpdateParamsMetadataObjectAssetsConfig] `json:"config"`
	// Token provided upon successful upload of all files from a registered manifest.
	JWT param.Field[string] `json:"jwt"`
}

func (r ScriptUpdateParamsMetadataObjectAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker.
type ScriptUpdateParamsMetadataObjectAssetsConfig struct {
	// The contents of a \_headers file (used to attach custom headers on asset
	// responses).
	Headers param.Field[string] `json:"_headers"`
	// The contents of a \_redirects file (used to apply redirects or proxy paths ahead
	// of asset serving).
	Redirects param.Field[string] `json:"_redirects"`
	// Determines the redirects and rewrites of requests for HTML content.
	HTMLHandling param.Field[ScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandling] `json:"html_handling"`
	// Determines the response when a request does not match a static asset, and there
	// is no Worker script.
	NotFoundHandling param.Field[ScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandling] `json:"not_found_handling"`
	// Contains a list path rules to control routing to either the Worker or assets.
	// Glob (\*) and negative (!) rules are supported. Rules must start with either '/'
	// or '!/'. At least one non-negative rule must be provided, and negative rules
	// have higher precedence than non-negative rules.
	RunWorkerFirst param.Field[ScriptUpdateParamsMetadataObjectAssetsConfigRunWorkerFirstUnion] `json:"run_worker_first"`
	// When true and the incoming request matches an asset, that will be served instead
	// of invoking the Worker script. When false, requests will always invoke the
	// Worker script.
	//
	// Deprecated: deprecated
	ServeDirectly param.Field[bool] `json:"serve_directly"`
}

func (r ScriptUpdateParamsMetadataObjectAssetsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the redirects and rewrites of requests for HTML content.
type ScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandling string

const (
	ScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandlingAutoTrailingSlash  ScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandling = "auto-trailing-slash"
	ScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandlingForceTrailingSlash ScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandling = "force-trailing-slash"
	ScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandlingDropTrailingSlash  ScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandling = "drop-trailing-slash"
	ScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandlingNone               ScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandling = "none"
)

func (r ScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandling) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandlingAutoTrailingSlash, ScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandlingForceTrailingSlash, ScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandlingDropTrailingSlash, ScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandlingNone:
		return true
	}
	return false
}

// Determines the response when a request does not match a static asset, and there
// is no Worker script.
type ScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandling string

const (
	ScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandlingNone                  ScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandling = "none"
	ScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandling404Page               ScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandling = "404-page"
	ScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandlingSinglePageApplication ScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandling = "single-page-application"
)

func (r ScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandling) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandlingNone, ScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandling404Page, ScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandlingSinglePageApplication:
		return true
	}
	return false
}

// Contains a list path rules to control routing to either the Worker or assets.
// Glob (\*) and negative (!) rules are supported. Rules must start with either '/'
// or '!/'. At least one non-negative rule must be provided, and negative rules
// have higher precedence than non-negative rules.
//
// Satisfied by
// [workers.ScriptUpdateParamsMetadataObjectAssetsConfigRunWorkerFirstArray],
// [shared.UnionBool].
type ScriptUpdateParamsMetadataObjectAssetsConfigRunWorkerFirstUnion interface {
	ImplementsScriptUpdateParamsMetadataObjectAssetsConfigRunWorkerFirstUnion()
}

type ScriptUpdateParamsMetadataObjectAssetsConfigRunWorkerFirstArray []string

func (r ScriptUpdateParamsMetadataObjectAssetsConfigRunWorkerFirstArray) ImplementsScriptUpdateParamsMetadataObjectAssetsConfigRunWorkerFirstUnion() {
}

// A binding to allow the Worker to communicate with resources.
type ScriptUpdateParamsMetadataObjectBinding struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsType] `json:"type,required"`
	// Identifier of the D1 database to bind to.
	ID        param.Field[string]      `json:"id"`
	Algorithm param.Field[interface{}] `json:"algorithm"`
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name"`
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id"`
	// The exported class name of the Durable Object.
	ClassName param.Field[string] `json:"class_name"`
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset"`
	// The environment of the script_name to bind to.
	Environment param.Field[string] `json:"environment"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[ScriptUpdateParamsMetadataObjectBindingsFormat] `json:"format"`
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name"`
	// JSON data to use.
	Json param.Field[string] `json:"json"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string]      `json:"key_base64"`
	KeyJwk    param.Field[interface{}] `json:"key_jwk"`
	// Namespace to bind to.
	Namespace param.Field[string] `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID param.Field[string]      `json:"namespace_id"`
	Outbound    param.Field[interface{}] `json:"outbound"`
	// Name of the Pipeline to bind to.
	Pipeline param.Field[string] `json:"pipeline"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName param.Field[string] `json:"script_name"`
	// Name of the secret in the store.
	SecretName param.Field[string] `json:"secret_name"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service"`
	// ID of the store containing the secret.
	StoreID param.Field[string] `json:"store_id"`
	// The text value to use.
	Text   param.Field[string]      `json:"text"`
	Usages param.Field[interface{}] `json:"usages"`
	// Name of the Workflow to bind to.
	WorkflowName param.Field[string] `json:"workflow_name"`
}

func (r ScriptUpdateParamsMetadataObjectBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBinding) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// A binding to allow the Worker to communicate with resources.
//
// Satisfied by
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAI],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngine],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssets],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowser],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespace],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespace],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdrive],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJson],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespace],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificate],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainText],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelines],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueue],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2Bucket],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretText],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindService],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumer],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorize],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadata],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecret],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKey],
// [workers.ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflow],
// [ScriptUpdateParamsMetadataObjectBinding].
type ScriptUpdateParamsMetadataObjectBindingUnion interface {
	implementsScriptUpdateParamsMetadataObjectBindingUnion()
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAIType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAI) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAIType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAITypeAI ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAIType = "ai"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngineType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngine) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngine) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngineType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssetsType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssets) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssetsType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssetsTypeAssets ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssetsType = "assets"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowserType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowser) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowserType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowserTypeBrowser ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowserType = "browser"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1Type] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1Type string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1TypeD1 ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1Type = "d1"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace to bind to.
	Namespace param.Field[string] `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceType] `json:"type,required"`
	// Outbound worker.
	Outbound param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceOutbound] `json:"outbound"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespace) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params param.Field[[]string] `json:"params"`
	// Outbound worker.
	Worker param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceOutboundWorker] `json:"worker"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker.
	Service param.Field[string] `json:"service"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespaceType] `json:"type,required"`
	// The exported class name of the Durable Object.
	ClassName param.Field[string] `json:"class_name"`
	// The environment of the script_name to bind to.
	Environment param.Field[string] `json:"environment"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespace) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdriveType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdrive) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdrive) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdriveType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdriveTypeHyperdrive ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json param.Field[string] `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJsonType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJson) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJsonType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJsonTypeJson ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJsonType = "json"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespaceType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespace) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespaceType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespaceTypeKVNamespace ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificateType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificate) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificateType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The text value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainTextType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainText) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainTextType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainTextTypePlainText ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline param.Field[string] `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelinesType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelines) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelines) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelinesType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelinesTypePipelines ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelinesType = "pipelines"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueueType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueue) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueueType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueueTypeQueue ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueueType = "queue"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2BucketType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2Bucket) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2Bucket) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2BucketType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2BucketTypeR2Bucket ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The secret value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretTextType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretText) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretTextType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretTextTypeSecretText ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretTextType = "secret_text"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindServiceType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindService) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindServiceType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindServiceTypeService ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindServiceType = "service"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumerType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumer) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumerType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumerTypeTailConsumer ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumerTypeTailConsumer:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorizeType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorize) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorizeType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorizeTypeVectorize ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorizeType = "vectorize"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadataType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadata) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadataType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the secret in the store.
	SecretName param.Field[string] `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID param.Field[string] `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecretType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecret) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecretType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm param.Field[interface{}] `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormat] `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyType] `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages param.Field[[]ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage] `json:"usages,required"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string] `json:"key_base64"`
	// Key data in
	// [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key)
	// format. Required if `format` is "jwk".
	KeyJwk param.Field[interface{}] `json:"key_jwk"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKey) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormat string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormatRaw   ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormat = "raw"
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormatPkcs8 ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormat = "pkcs8"
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormatSpki  ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormat = "spki"
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormatJwk   ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormatRaw, ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormatPkcs8, ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormatSpki, ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyTypeSecretKey ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyType = "secret_key"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageEncrypt    ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage = "encrypt"
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageDecrypt    ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage = "decrypt"
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageSign       ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage = "sign"
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageVerify     ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage = "verify"
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageDeriveKey  ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage = "deriveKey"
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageDeriveBits ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage = "deriveBits"
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageWrapKey    ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage = "wrapKey"
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageUnwrapKey  ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageEncrypt, ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageDecrypt, ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageSign, ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageVerify, ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageDeriveKey, ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageDeriveBits, ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageWrapKey, ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflowType] `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName param.Field[string] `json:"workflow_name,required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName param.Field[string] `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflow) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflow) implementsScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflowType string

const (
	ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflowTypeWorkflow ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflowType = "workflow"
)

func (r ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataObjectBindingsType string

const (
	ScriptUpdateParamsMetadataObjectBindingsTypeAI                     ScriptUpdateParamsMetadataObjectBindingsType = "ai"
	ScriptUpdateParamsMetadataObjectBindingsTypeAnalyticsEngine        ScriptUpdateParamsMetadataObjectBindingsType = "analytics_engine"
	ScriptUpdateParamsMetadataObjectBindingsTypeAssets                 ScriptUpdateParamsMetadataObjectBindingsType = "assets"
	ScriptUpdateParamsMetadataObjectBindingsTypeBrowser                ScriptUpdateParamsMetadataObjectBindingsType = "browser"
	ScriptUpdateParamsMetadataObjectBindingsTypeD1                     ScriptUpdateParamsMetadataObjectBindingsType = "d1"
	ScriptUpdateParamsMetadataObjectBindingsTypeDispatchNamespace      ScriptUpdateParamsMetadataObjectBindingsType = "dispatch_namespace"
	ScriptUpdateParamsMetadataObjectBindingsTypeDurableObjectNamespace ScriptUpdateParamsMetadataObjectBindingsType = "durable_object_namespace"
	ScriptUpdateParamsMetadataObjectBindingsTypeHyperdrive             ScriptUpdateParamsMetadataObjectBindingsType = "hyperdrive"
	ScriptUpdateParamsMetadataObjectBindingsTypeJson                   ScriptUpdateParamsMetadataObjectBindingsType = "json"
	ScriptUpdateParamsMetadataObjectBindingsTypeKVNamespace            ScriptUpdateParamsMetadataObjectBindingsType = "kv_namespace"
	ScriptUpdateParamsMetadataObjectBindingsTypeMTLSCertificate        ScriptUpdateParamsMetadataObjectBindingsType = "mtls_certificate"
	ScriptUpdateParamsMetadataObjectBindingsTypePlainText              ScriptUpdateParamsMetadataObjectBindingsType = "plain_text"
	ScriptUpdateParamsMetadataObjectBindingsTypePipelines              ScriptUpdateParamsMetadataObjectBindingsType = "pipelines"
	ScriptUpdateParamsMetadataObjectBindingsTypeQueue                  ScriptUpdateParamsMetadataObjectBindingsType = "queue"
	ScriptUpdateParamsMetadataObjectBindingsTypeR2Bucket               ScriptUpdateParamsMetadataObjectBindingsType = "r2_bucket"
	ScriptUpdateParamsMetadataObjectBindingsTypeSecretText             ScriptUpdateParamsMetadataObjectBindingsType = "secret_text"
	ScriptUpdateParamsMetadataObjectBindingsTypeService                ScriptUpdateParamsMetadataObjectBindingsType = "service"
	ScriptUpdateParamsMetadataObjectBindingsTypeTailConsumer           ScriptUpdateParamsMetadataObjectBindingsType = "tail_consumer"
	ScriptUpdateParamsMetadataObjectBindingsTypeVectorize              ScriptUpdateParamsMetadataObjectBindingsType = "vectorize"
	ScriptUpdateParamsMetadataObjectBindingsTypeVersionMetadata        ScriptUpdateParamsMetadataObjectBindingsType = "version_metadata"
	ScriptUpdateParamsMetadataObjectBindingsTypeSecretsStoreSecret     ScriptUpdateParamsMetadataObjectBindingsType = "secrets_store_secret"
	ScriptUpdateParamsMetadataObjectBindingsTypeSecretKey              ScriptUpdateParamsMetadataObjectBindingsType = "secret_key"
	ScriptUpdateParamsMetadataObjectBindingsTypeWorkflow               ScriptUpdateParamsMetadataObjectBindingsType = "workflow"
)

func (r ScriptUpdateParamsMetadataObjectBindingsType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsTypeAI, ScriptUpdateParamsMetadataObjectBindingsTypeAnalyticsEngine, ScriptUpdateParamsMetadataObjectBindingsTypeAssets, ScriptUpdateParamsMetadataObjectBindingsTypeBrowser, ScriptUpdateParamsMetadataObjectBindingsTypeD1, ScriptUpdateParamsMetadataObjectBindingsTypeDispatchNamespace, ScriptUpdateParamsMetadataObjectBindingsTypeDurableObjectNamespace, ScriptUpdateParamsMetadataObjectBindingsTypeHyperdrive, ScriptUpdateParamsMetadataObjectBindingsTypeJson, ScriptUpdateParamsMetadataObjectBindingsTypeKVNamespace, ScriptUpdateParamsMetadataObjectBindingsTypeMTLSCertificate, ScriptUpdateParamsMetadataObjectBindingsTypePlainText, ScriptUpdateParamsMetadataObjectBindingsTypePipelines, ScriptUpdateParamsMetadataObjectBindingsTypeQueue, ScriptUpdateParamsMetadataObjectBindingsTypeR2Bucket, ScriptUpdateParamsMetadataObjectBindingsTypeSecretText, ScriptUpdateParamsMetadataObjectBindingsTypeService, ScriptUpdateParamsMetadataObjectBindingsTypeTailConsumer, ScriptUpdateParamsMetadataObjectBindingsTypeVectorize, ScriptUpdateParamsMetadataObjectBindingsTypeVersionMetadata, ScriptUpdateParamsMetadataObjectBindingsTypeSecretsStoreSecret, ScriptUpdateParamsMetadataObjectBindingsTypeSecretKey, ScriptUpdateParamsMetadataObjectBindingsTypeWorkflow:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptUpdateParamsMetadataObjectBindingsFormat string

const (
	ScriptUpdateParamsMetadataObjectBindingsFormatRaw   ScriptUpdateParamsMetadataObjectBindingsFormat = "raw"
	ScriptUpdateParamsMetadataObjectBindingsFormatPkcs8 ScriptUpdateParamsMetadataObjectBindingsFormat = "pkcs8"
	ScriptUpdateParamsMetadataObjectBindingsFormatSpki  ScriptUpdateParamsMetadataObjectBindingsFormat = "spki"
	ScriptUpdateParamsMetadataObjectBindingsFormatJwk   ScriptUpdateParamsMetadataObjectBindingsFormat = "jwk"
)

func (r ScriptUpdateParamsMetadataObjectBindingsFormat) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectBindingsFormatRaw, ScriptUpdateParamsMetadataObjectBindingsFormatPkcs8, ScriptUpdateParamsMetadataObjectBindingsFormatSpki, ScriptUpdateParamsMetadataObjectBindingsFormatJwk:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
type ScriptUpdateParamsMetadataObjectMigrations struct {
	DeletedClasses   param.Field[interface{}] `json:"deleted_classes"`
	NewClasses       param.Field[interface{}] `json:"new_classes"`
	NewSqliteClasses param.Field[interface{}] `json:"new_sqlite_classes"`
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag             param.Field[string]      `json:"old_tag"`
	RenamedClasses     param.Field[interface{}] `json:"renamed_classes"`
	Steps              param.Field[interface{}] `json:"steps"`
	TransferredClasses param.Field[interface{}] `json:"transferred_classes"`
}

func (r ScriptUpdateParamsMetadataObjectMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectMigrations) implementsScriptUpdateParamsMetadataObjectMigrationsUnion() {
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by [workers.SingleStepMigrationParam],
// [workers.ScriptUpdateParamsMetadataObjectMigrationsWorkersMultipleStepMigrations],
// [ScriptUpdateParamsMetadataObjectMigrations].
type ScriptUpdateParamsMetadataObjectMigrationsUnion interface {
	implementsScriptUpdateParamsMetadataObjectMigrationsUnion()
}

type ScriptUpdateParamsMetadataObjectMigrationsWorkersMultipleStepMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]MigrationStepParam] `json:"steps"`
}

func (r ScriptUpdateParamsMetadataObjectMigrationsWorkersMultipleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataObjectMigrationsWorkersMultipleStepMigrations) implementsScriptUpdateParamsMetadataObjectMigrationsUnion() {
}

// Observability settings for the Worker.
type ScriptUpdateParamsMetadataObjectObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
	// Log settings for the Worker.
	Logs param.Field[ScriptUpdateParamsMetadataObjectObservabilityLogs] `json:"logs"`
}

func (r ScriptUpdateParamsMetadataObjectObservability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Log settings for the Worker.
type ScriptUpdateParamsMetadataObjectObservabilityLogs struct {
	// Whether logs are enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs param.Field[bool] `json:"invocation_logs,required"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
}

func (r ScriptUpdateParamsMetadataObjectObservabilityLogs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptUpdateParamsMetadataObjectPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode param.Field[ScriptUpdateParamsMetadataObjectPlacementMode] `json:"mode"`
}

func (r ScriptUpdateParamsMetadataObjectPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptUpdateParamsMetadataObjectPlacementMode string

const (
	ScriptUpdateParamsMetadataObjectPlacementModeSmart ScriptUpdateParamsMetadataObjectPlacementMode = "smart"
)

func (r ScriptUpdateParamsMetadataObjectPlacementMode) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectPlacementModeSmart:
		return true
	}
	return false
}

// Status of
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptUpdateParamsMetadataObjectPlacementStatus string

const (
	ScriptUpdateParamsMetadataObjectPlacementStatusSuccess                 ScriptUpdateParamsMetadataObjectPlacementStatus = "SUCCESS"
	ScriptUpdateParamsMetadataObjectPlacementStatusUnsupportedApplication  ScriptUpdateParamsMetadataObjectPlacementStatus = "UNSUPPORTED_APPLICATION"
	ScriptUpdateParamsMetadataObjectPlacementStatusInsufficientInvocations ScriptUpdateParamsMetadataObjectPlacementStatus = "INSUFFICIENT_INVOCATIONS"
)

func (r ScriptUpdateParamsMetadataObjectPlacementStatus) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectPlacementStatusSuccess, ScriptUpdateParamsMetadataObjectPlacementStatusUnsupportedApplication, ScriptUpdateParamsMetadataObjectPlacementStatusInsufficientInvocations:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type ScriptUpdateParamsMetadataObjectUsageModel string

const (
	ScriptUpdateParamsMetadataObjectUsageModelStandard ScriptUpdateParamsMetadataObjectUsageModel = "standard"
)

func (r ScriptUpdateParamsMetadataObjectUsageModel) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataObjectUsageModelStandard:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type ScriptUpdateParamsMetadataUsageModel string

const (
	ScriptUpdateParamsMetadataUsageModelStandard ScriptUpdateParamsMetadataUsageModel = "standard"
)

func (r ScriptUpdateParamsMetadataUsageModel) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataUsageModelStandard:
		return true
	}
	return false
}

type ScriptUpdateResponseEnvelope struct {
	Errors   []ScriptUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ScriptUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptUpdateResponseEnvelopeJSON    `json:"-"`
}

// scriptUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptUpdateResponseEnvelope]
type scriptUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptUpdateResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           ScriptUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ScriptUpdateResponseEnvelopeErrors]
type scriptUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    scriptUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ScriptUpdateResponseEnvelopeErrorsSource]
type scriptUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptUpdateResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           ScriptUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScriptUpdateResponseEnvelopeMessages]
type scriptUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    scriptUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ScriptUpdateResponseEnvelopeMessagesSource]
type scriptUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptUpdateResponseEnvelopeSuccess bool

const (
	ScriptUpdateResponseEnvelopeSuccessTrue ScriptUpdateResponseEnvelopeSuccess = true
)

func (r ScriptUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Filter scripts by tags. Format: comma-separated list of tag:allowed pairs where
	// allowed is 'yes' or 'no'.
	Tags param.Field[string] `query:"tags"`
}

// URLQuery serializes [ScriptListParams]'s query parameters as `url.Values`.
func (r ScriptListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ScriptDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// If set to true, delete will not be stopped by associated service binding,
	// durable object, or other binding. Any of these associated bindings/durable
	// objects will be deleted along with the script.
	Force param.Field[bool] `query:"force"`
}

// URLQuery serializes [ScriptDeleteParams]'s query parameters as `url.Values`.
func (r ScriptDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ScriptDeleteResponseEnvelope struct {
	Errors   []ScriptDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ScriptDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptDeleteResponse                `json:"result,nullable"`
	JSON    scriptDeleteResponseEnvelopeJSON    `json:"-"`
}

// scriptDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptDeleteResponseEnvelope]
type scriptDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptDeleteResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           ScriptDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ScriptDeleteResponseEnvelopeErrors]
type scriptDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    scriptDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ScriptDeleteResponseEnvelopeErrorsSource]
type scriptDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptDeleteResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           ScriptDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScriptDeleteResponseEnvelopeMessages]
type scriptDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    scriptDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ScriptDeleteResponseEnvelopeMessagesSource]
type scriptDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptDeleteResponseEnvelopeSuccess bool

const (
	ScriptDeleteResponseEnvelopeSuccessTrue ScriptDeleteResponseEnvelopeSuccess = true
)

func (r ScriptDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}
