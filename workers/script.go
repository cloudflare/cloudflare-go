// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// ScriptService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptService] method instead.
type ScriptService struct {
	Options     []option.RequestOption
	Assets      *ScriptAssetService
	Subdomain   *ScriptSubdomainService
	Schedules   *ScriptScheduleService
	Tail        *ScriptTailService
	Content     *ScriptContentService
	Settings    *ScriptSettingService
	Deployments *ScriptDeploymentService
	Versions    *ScriptVersionService
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
func (r *ScriptService) List(ctx context.Context, query ScriptListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Script], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts", query.AccountID)
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

// Fetch a list of uploaded workers.
func (r *ScriptService) ListAutoPaging(ctx context.Context, query ScriptListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Script] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete your worker. This call has no response body on a successful delete.
func (r *ScriptService) Delete(ctx context.Context, scriptName string, params ScriptDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
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
	Mode        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	Observability ScriptSettingObservability `json:"observability"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ConsumerScript  `json:"tail_consumers"`
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
	HeadSamplingRate float64                        `json:"head_sampling_rate,nullable"`
	JSON             scriptSettingObservabilityJSON `json:"-"`
}

// scriptSettingObservabilityJSON contains the JSON metadata for the struct
// [ScriptSettingObservability]
type scriptSettingObservabilityJSON struct {
	Enabled          apijson.Field
	HeadSamplingRate apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSettingObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingObservabilityJSON) RawJSON() string {
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
}

func (r ScriptSettingObservabilityParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptUpdateResponse struct {
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
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	//
	// Deprecated: deprecated
	PlacementMode ScriptUpdateResponsePlacementMode `json:"placement_mode"`
	// Status of
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	//
	// Deprecated: deprecated
	PlacementStatus ScriptUpdateResponsePlacementStatus `json:"placement_status"`
	StartupTimeMs   int64                               `json:"startup_time_ms"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ConsumerScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel ScriptUpdateResponseUsageModel `json:"usage_model"`
	JSON       scriptUpdateResponseJSON       `json:"-"`
}

// scriptUpdateResponseJSON contains the JSON metadata for the struct
// [ScriptUpdateResponse]
type scriptUpdateResponseJSON struct {
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
	StartupTimeMs   apijson.Field
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
	Mode        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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

type ScriptUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[ScriptUpdateParamsMetadata] `json:"metadata,required"`
}

func (r ScriptUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// JSON encoded metadata about the uploaded parts and Worker configuration.
type ScriptUpdateParamsMetadata struct {
	// Configuration for assets within a Worker
	Assets param.Field[ScriptUpdateParamsMetadataAssets] `json:"assets"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings param.Field[[]ScriptUpdateParamsMetadataBindingUnion] `json:"bindings"`
	// Name of the part in the multipart request that contains the script (e.g. the
	// file adding a listener to the `fetch` event). Indicates a
	// `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
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
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[ScriptUpdateParamsMetadataMigrationsUnion] `json:"migrations"`
	// Observability settings for the Worker.
	Observability param.Field[ScriptUpdateParamsMetadataObservability] `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement param.Field[ScriptUpdateParamsMetadataPlacement] `json:"placement"`
	// List of strings to use as tags for this Worker.
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]ConsumerScriptParam] `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[ScriptUpdateParamsMetadataUsageModel] `json:"usage_model"`
}

func (r ScriptUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker
type ScriptUpdateParamsMetadataAssets struct {
	// Configuration for assets within a Worker.
	Config param.Field[ScriptUpdateParamsMetadataAssetsConfig] `json:"config"`
	// Token provided upon successful upload of all files from a registered manifest.
	JWT param.Field[string] `json:"jwt"`
}

func (r ScriptUpdateParamsMetadataAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker.
type ScriptUpdateParamsMetadataAssetsConfig struct {
	// Determines the redirects and rewrites of requests for HTML content.
	HTMLHandling param.Field[ScriptUpdateParamsMetadataAssetsConfigHTMLHandling] `json:"html_handling"`
	// Determines the response when a request does not match a static asset, and there
	// is no Worker script.
	NotFoundHandling param.Field[ScriptUpdateParamsMetadataAssetsConfigNotFoundHandling] `json:"not_found_handling"`
	// When true, requests will always invoke the Worker script. Otherwise, attempt to
	// serve an asset matching the request, falling back to the Worker script.
	RunWorkerFirst param.Field[bool] `json:"run_worker_first"`
	// When true and the incoming request matches an asset, that will be served instead
	// of invoking the Worker script. When false, requests will always invoke the
	// Worker script.
	//
	// Deprecated: deprecated
	ServeDirectly param.Field[bool] `json:"serve_directly"`
}

func (r ScriptUpdateParamsMetadataAssetsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the redirects and rewrites of requests for HTML content.
type ScriptUpdateParamsMetadataAssetsConfigHTMLHandling string

const (
	ScriptUpdateParamsMetadataAssetsConfigHTMLHandlingAutoTrailingSlash  ScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "auto-trailing-slash"
	ScriptUpdateParamsMetadataAssetsConfigHTMLHandlingForceTrailingSlash ScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "force-trailing-slash"
	ScriptUpdateParamsMetadataAssetsConfigHTMLHandlingDropTrailingSlash  ScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "drop-trailing-slash"
	ScriptUpdateParamsMetadataAssetsConfigHTMLHandlingNone               ScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "none"
)

func (r ScriptUpdateParamsMetadataAssetsConfigHTMLHandling) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataAssetsConfigHTMLHandlingAutoTrailingSlash, ScriptUpdateParamsMetadataAssetsConfigHTMLHandlingForceTrailingSlash, ScriptUpdateParamsMetadataAssetsConfigHTMLHandlingDropTrailingSlash, ScriptUpdateParamsMetadataAssetsConfigHTMLHandlingNone:
		return true
	}
	return false
}

// Determines the response when a request does not match a static asset, and there
// is no Worker script.
type ScriptUpdateParamsMetadataAssetsConfigNotFoundHandling string

const (
	ScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingNone                  ScriptUpdateParamsMetadataAssetsConfigNotFoundHandling = "none"
	ScriptUpdateParamsMetadataAssetsConfigNotFoundHandling404Page               ScriptUpdateParamsMetadataAssetsConfigNotFoundHandling = "404-page"
	ScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingSinglePageApplication ScriptUpdateParamsMetadataAssetsConfigNotFoundHandling = "single-page-application"
)

func (r ScriptUpdateParamsMetadataAssetsConfigNotFoundHandling) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingNone, ScriptUpdateParamsMetadataAssetsConfigNotFoundHandling404Page, ScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingSinglePageApplication:
		return true
	}
	return false
}

// A binding to allow the Worker to communicate with resources
type ScriptUpdateParamsMetadataBinding struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsType] `json:"type,required"`
	// Identifier of the D1 database to bind to.
	ID param.Field[string] `json:"id"`
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
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name"`
	// JSON data to use.
	Json param.Field[string] `json:"json"`
	// Namespace to bind to.
	Namespace param.Field[string] `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID param.Field[string]      `json:"namespace_id"`
	Outbound    param.Field[interface{}] `json:"outbound"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName param.Field[string] `json:"script_name"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service"`
	// The text value to use.
	Text param.Field[string] `json:"text"`
}

func (r ScriptUpdateParamsMetadataBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBinding) implementsScriptUpdateParamsMetadataBindingUnion() {}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindAI],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRendering],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindJson],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindService],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumer],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata],
// [ScriptUpdateParamsMetadataBinding].
type ScriptUpdateParamsMetadataBindingUnion interface {
	implementsScriptUpdateParamsMetadataBindingUnion()
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAI) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAI                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "ai"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAssets                 ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "assets"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeBrowserRendering       ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeD1                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "d1"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeHyperdrive             ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeJson                   ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "json"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeKVNamespace            ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypePlainText              ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "plain_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeQueue                  ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "queue"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeR2Bucket               ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeSecretText             ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "secret_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeService                ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "service"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeTailConsumer           ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeVectorize              ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "vectorize"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeVersionMetadata        ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAI, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAssets, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeBrowserRendering, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeD1, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeHyperdrive, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeJson, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeKVNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypePlainText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeQueue, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeR2Bucket, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeSecretText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeService, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeTailConsumer, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeVectorize, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeVersionMetadata:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAI                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "ai"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAssets                 ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "assets"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeBrowserRendering       ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeD1                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "d1"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeHyperdrive             ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeJson                   ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "json"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeKVNamespace            ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypePlainText              ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "plain_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeQueue                  ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "queue"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeR2Bucket               ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeSecretText             ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "secret_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeService                ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "service"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeTailConsumer           ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeVectorize              ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "vectorize"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeVersionMetadata        ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAI, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAssets, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeBrowserRendering, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeD1, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeHyperdrive, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeJson, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeKVNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypePlainText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeQueue, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeR2Bucket, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeSecretText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeService, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeTailConsumer, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeVectorize, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAI                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "ai"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAssets                 ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "assets"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeBrowserRendering       ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeD1                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "d1"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeHyperdrive             ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeJson                   ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "json"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeKVNamespace            ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypePlainText              ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "plain_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeQueue                  ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "queue"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeR2Bucket               ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeSecretText             ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "secret_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeService                ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "service"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeTailConsumer           ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeVectorize              ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "vectorize"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeVersionMetadata        ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAI, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAssets, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeBrowserRendering, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeD1, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeHyperdrive, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeJson, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeKVNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypePlainText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeQueue, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeR2Bucket, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeSecretText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeService, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeTailConsumer, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeVectorize, ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRendering struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRendering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRendering) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeAI                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "ai"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeAssets                 ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "assets"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering       ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeD1                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "d1"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeHyperdrive             ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeJson                   ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "json"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeKVNamespace            ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypePlainText              ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "plain_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeQueue                  ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "queue"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeR2Bucket               ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeSecretText             ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "secret_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeService                ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "service"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeTailConsumer           ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeVectorize              ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "vectorize"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeVersionMetadata        ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeAI, ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeAssets, ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering, ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeD1, ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeHyperdrive, ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeJson, ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeKVNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypePlainText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeQueue, ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeR2Bucket, ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeSecretText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeService, ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeTailConsumer, ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeVectorize, ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeAI                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "ai"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeAssets                 ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "assets"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeBrowserRendering       ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeD1                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "d1"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeHyperdrive             ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeJson                   ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "json"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeKVNamespace            ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypePlainText              ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "plain_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeQueue                  ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "queue"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeR2Bucket               ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeSecretText             ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "secret_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeService                ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "service"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeTailConsumer           ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeVectorize              ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "vectorize"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeVersionMetadata        ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeAI, ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeAssets, ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeBrowserRendering, ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeD1, ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeHyperdrive, ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeJson, ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeKVNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypePlainText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeQueue, ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeR2Bucket, ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeSecretText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeService, ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeTailConsumer, ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeVectorize, ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeVersionMetadata:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace to bind to.
	Namespace param.Field[string] `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType] `json:"type,required"`
	// Outbound worker.
	Outbound param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound] `json:"outbound"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeAI                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "ai"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeAssets                 ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "assets"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeBrowserRendering       ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeD1                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "d1"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeHyperdrive             ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeJson                   ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "json"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeKVNamespace            ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypePlainText              ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "plain_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeQueue                  ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "queue"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeR2Bucket               ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeSecretText             ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "secret_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeService                ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "service"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeTailConsumer           ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeVectorize              ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "vectorize"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeVersionMetadata        ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeAI, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeAssets, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeBrowserRendering, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeD1, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeHyperdrive, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeJson, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeKVNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypePlainText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeQueue, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeR2Bucket, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeSecretText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeService, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeTailConsumer, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeVectorize, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

// Outbound worker.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params param.Field[[]string] `json:"params"`
	// Outbound worker.
	Worker param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker] `json:"worker"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker.
	Service param.Field[string] `json:"service"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace struct {
	// The exported class name of the Durable Object.
	ClassName param.Field[string] `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType] `json:"type,required"`
	// The environment of the script_name to bind to.
	Environment param.Field[string] `json:"environment"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeAI                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "ai"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeAssets                 ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "assets"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering       ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeD1                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "d1"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeHyperdrive             ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeJson                   ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "json"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeKVNamespace            ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypePlainText              ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "plain_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeQueue                  ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "queue"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeR2Bucket               ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeSecretText             ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "secret_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeService                ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "service"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeTailConsumer           ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeVectorize              ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "vectorize"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata        ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeAI, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeAssets, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeD1, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeHyperdrive, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeJson, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeKVNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypePlainText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeQueue, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeR2Bucket, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeSecretText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeService, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeTailConsumer, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeVectorize, ScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeAI                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "ai"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeAssets                 ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "assets"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeBrowserRendering       ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeD1                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "d1"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive             ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeJson                   ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "json"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeKVNamespace            ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypePlainText              ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "plain_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeQueue                  ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "queue"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeR2Bucket               ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeSecretText             ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "secret_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeService                ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "service"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeTailConsumer           ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeVectorize              ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "vectorize"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeVersionMetadata        ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeAI, ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeAssets, ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeBrowserRendering, ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeD1, ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive, ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeJson, ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeKVNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypePlainText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeQueue, ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeR2Bucket, ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeSecretText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeService, ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeTailConsumer, ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeVectorize, ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json param.Field[string] `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindJson) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeAI                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "ai"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeAssets                 ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "assets"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeBrowserRendering       ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeD1                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "d1"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeHyperdrive             ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeJson                   ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "json"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeKVNamespace            ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypePlainText              ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "plain_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeQueue                  ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "queue"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeR2Bucket               ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeSecretText             ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "secret_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeService                ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "service"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeTailConsumer           ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeVectorize              ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "vectorize"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeVersionMetadata        ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeAI, ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeAssets, ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeBrowserRendering, ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeD1, ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeHyperdrive, ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeJson, ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeKVNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypePlainText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeQueue, ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeR2Bucket, ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeSecretText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeService, ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeTailConsumer, ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeVectorize, ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeAI                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "ai"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeAssets                 ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "assets"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeBrowserRendering       ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeD1                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "d1"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeHyperdrive             ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeJson                   ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "json"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace            ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypePlainText              ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "plain_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeQueue                  ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "queue"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeR2Bucket               ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeSecretText             ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "secret_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeService                ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "service"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeTailConsumer           ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeVectorize              ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "vectorize"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeVersionMetadata        ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeAI, ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeAssets, ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeBrowserRendering, ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeD1, ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeHyperdrive, ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeJson, ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypePlainText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeQueue, ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeR2Bucket, ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeSecretText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeService, ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeTailConsumer, ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeVectorize, ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeAI                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "ai"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeAssets                 ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "assets"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeBrowserRendering       ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeD1                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "d1"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeHyperdrive             ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeJson                   ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "json"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeKVNamespace            ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypePlainText              ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "plain_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeQueue                  ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "queue"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeR2Bucket               ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeSecretText             ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "secret_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeService                ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "service"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeTailConsumer           ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeVectorize              ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "vectorize"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeVersionMetadata        ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeAI, ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeAssets, ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeBrowserRendering, ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeD1, ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeHyperdrive, ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeJson, ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeKVNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypePlainText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeQueue, ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeR2Bucket, ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeSecretText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeService, ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeTailConsumer, ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeVectorize, ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The text value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeAI                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "ai"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeAssets                 ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "assets"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeBrowserRendering       ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeD1                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "d1"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeHyperdrive             ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeJson                   ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "json"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeKVNamespace            ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText              ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "plain_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeQueue                  ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "queue"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeR2Bucket               ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeSecretText             ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "secret_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeService                ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "service"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeTailConsumer           ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeVectorize              ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "vectorize"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeVersionMetadata        ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeAI, ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeAssets, ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeBrowserRendering, ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeD1, ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeHyperdrive, ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeJson, ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeKVNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeQueue, ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeR2Bucket, ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeSecretText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeService, ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeTailConsumer, ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeVectorize, ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeAI                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "ai"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeAssets                 ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "assets"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeBrowserRendering       ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeD1                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "d1"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeHyperdrive             ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeJson                   ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "json"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeKVNamespace            ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypePlainText              ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "plain_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeQueue                  ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "queue"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeR2Bucket               ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeSecretText             ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "secret_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeService                ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "service"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeTailConsumer           ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeVectorize              ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "vectorize"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeVersionMetadata        ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeAI, ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeAssets, ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeBrowserRendering, ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeD1, ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeHyperdrive, ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeJson, ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeKVNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypePlainText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeQueue, ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeR2Bucket, ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeSecretText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeService, ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeTailConsumer, ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeVectorize, ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeAI                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "ai"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeAssets                 ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "assets"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeBrowserRendering       ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeD1                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "d1"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeHyperdrive             ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeJson                   ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "json"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeKVNamespace            ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypePlainText              ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "plain_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeQueue                  ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "queue"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeR2Bucket               ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeSecretText             ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "secret_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeService                ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "service"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeTailConsumer           ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeVectorize              ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "vectorize"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeVersionMetadata        ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeAI, ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeAssets, ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeBrowserRendering, ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeD1, ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeHyperdrive, ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeJson, ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeKVNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypePlainText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeQueue, ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeR2Bucket, ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeSecretText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeService, ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeTailConsumer, ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeVectorize, ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The secret value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeAI                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "ai"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeAssets                 ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "assets"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeBrowserRendering       ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeD1                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "d1"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeHyperdrive             ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeJson                   ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "json"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeKVNamespace            ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypePlainText              ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "plain_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeQueue                  ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "queue"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeR2Bucket               ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeSecretText             ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "secret_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeService                ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "service"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeTailConsumer           ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeVectorize              ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "vectorize"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeVersionMetadata        ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeAI, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeAssets, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeBrowserRendering, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeD1, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeHyperdrive, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeJson, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeKVNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypePlainText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeQueue, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeR2Bucket, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeSecretText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeService, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeTailConsumer, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeVectorize, ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindService) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeAI                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "ai"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeAssets                 ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "assets"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeBrowserRendering       ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeD1                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "d1"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeHyperdrive             ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeJson                   ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "json"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeKVNamespace            ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypePlainText              ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "plain_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeQueue                  ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "queue"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeR2Bucket               ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeSecretText             ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "secret_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeService                ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "service"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeTailConsumer           ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeVectorize              ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "vectorize"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeVersionMetadata        ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeAI, ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeAssets, ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeBrowserRendering, ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeD1, ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeHyperdrive, ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeJson, ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeKVNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypePlainText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeQueue, ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeR2Bucket, ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeSecretText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeService, ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeTailConsumer, ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeVectorize, ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumer) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeAI                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "ai"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeAssets                 ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "assets"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeBrowserRendering       ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeD1                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "d1"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeHyperdrive             ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeJson                   ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "json"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeKVNamespace            ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypePlainText              ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "plain_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeQueue                  ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "queue"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeR2Bucket               ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeSecretText             ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "secret_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeService                ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "service"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeTailConsumer           ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeVectorize              ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "vectorize"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeVersionMetadata        ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeAI, ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeAssets, ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeBrowserRendering, ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeD1, ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeHyperdrive, ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeJson, ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeKVNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypePlainText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeQueue, ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeR2Bucket, ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeSecretText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeService, ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeTailConsumer, ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeVectorize, ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeAI                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "ai"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeAssets                 ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "assets"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeBrowserRendering       ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeD1                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "d1"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeHyperdrive             ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeJson                   ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "json"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeKVNamespace            ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypePlainText              ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "plain_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeQueue                  ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "queue"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeR2Bucket               ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeSecretText             ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "secret_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeService                ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "service"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeTailConsumer           ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeVectorize              ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "vectorize"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeVersionMetadata        ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeAI, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeAssets, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeBrowserRendering, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeD1, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeHyperdrive, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeJson, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeKVNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypePlainText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeQueue, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeR2Bucket, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeSecretText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeService, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeTailConsumer, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeVectorize, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata) implementsScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeAI                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "ai"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeAssets                 ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "assets"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeBrowserRendering       ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeD1                     ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "d1"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeHyperdrive             ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeJson                   ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "json"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeKVNamespace            ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypePlainText              ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "plain_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeQueue                  ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "queue"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeR2Bucket               ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeSecretText             ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "secret_text"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeService                ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "service"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeTailConsumer           ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVectorize              ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "vectorize"
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata        ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeAI, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeAssets, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeBrowserRendering, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeD1, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeHyperdrive, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeJson, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeKVNamespace, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypePlainText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeQueue, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeR2Bucket, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeSecretText, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeService, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeTailConsumer, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVectorize, ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsType string

const (
	ScriptUpdateParamsMetadataBindingsTypeAI                     ScriptUpdateParamsMetadataBindingsType = "ai"
	ScriptUpdateParamsMetadataBindingsTypeAnalyticsEngine        ScriptUpdateParamsMetadataBindingsType = "analytics_engine"
	ScriptUpdateParamsMetadataBindingsTypeAssets                 ScriptUpdateParamsMetadataBindingsType = "assets"
	ScriptUpdateParamsMetadataBindingsTypeBrowserRendering       ScriptUpdateParamsMetadataBindingsType = "browser_rendering"
	ScriptUpdateParamsMetadataBindingsTypeD1                     ScriptUpdateParamsMetadataBindingsType = "d1"
	ScriptUpdateParamsMetadataBindingsTypeDispatchNamespace      ScriptUpdateParamsMetadataBindingsType = "dispatch_namespace"
	ScriptUpdateParamsMetadataBindingsTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsType = "durable_object_namespace"
	ScriptUpdateParamsMetadataBindingsTypeHyperdrive             ScriptUpdateParamsMetadataBindingsType = "hyperdrive"
	ScriptUpdateParamsMetadataBindingsTypeJson                   ScriptUpdateParamsMetadataBindingsType = "json"
	ScriptUpdateParamsMetadataBindingsTypeKVNamespace            ScriptUpdateParamsMetadataBindingsType = "kv_namespace"
	ScriptUpdateParamsMetadataBindingsTypeMTLSCertificate        ScriptUpdateParamsMetadataBindingsType = "mtls_certificate"
	ScriptUpdateParamsMetadataBindingsTypePlainText              ScriptUpdateParamsMetadataBindingsType = "plain_text"
	ScriptUpdateParamsMetadataBindingsTypeQueue                  ScriptUpdateParamsMetadataBindingsType = "queue"
	ScriptUpdateParamsMetadataBindingsTypeR2Bucket               ScriptUpdateParamsMetadataBindingsType = "r2_bucket"
	ScriptUpdateParamsMetadataBindingsTypeSecretText             ScriptUpdateParamsMetadataBindingsType = "secret_text"
	ScriptUpdateParamsMetadataBindingsTypeService                ScriptUpdateParamsMetadataBindingsType = "service"
	ScriptUpdateParamsMetadataBindingsTypeTailConsumer           ScriptUpdateParamsMetadataBindingsType = "tail_consumer"
	ScriptUpdateParamsMetadataBindingsTypeVectorize              ScriptUpdateParamsMetadataBindingsType = "vectorize"
	ScriptUpdateParamsMetadataBindingsTypeVersionMetadata        ScriptUpdateParamsMetadataBindingsType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsTypeAI, ScriptUpdateParamsMetadataBindingsTypeAnalyticsEngine, ScriptUpdateParamsMetadataBindingsTypeAssets, ScriptUpdateParamsMetadataBindingsTypeBrowserRendering, ScriptUpdateParamsMetadataBindingsTypeD1, ScriptUpdateParamsMetadataBindingsTypeDispatchNamespace, ScriptUpdateParamsMetadataBindingsTypeDurableObjectNamespace, ScriptUpdateParamsMetadataBindingsTypeHyperdrive, ScriptUpdateParamsMetadataBindingsTypeJson, ScriptUpdateParamsMetadataBindingsTypeKVNamespace, ScriptUpdateParamsMetadataBindingsTypeMTLSCertificate, ScriptUpdateParamsMetadataBindingsTypePlainText, ScriptUpdateParamsMetadataBindingsTypeQueue, ScriptUpdateParamsMetadataBindingsTypeR2Bucket, ScriptUpdateParamsMetadataBindingsTypeSecretText, ScriptUpdateParamsMetadataBindingsTypeService, ScriptUpdateParamsMetadataBindingsTypeTailConsumer, ScriptUpdateParamsMetadataBindingsTypeVectorize, ScriptUpdateParamsMetadataBindingsTypeVersionMetadata:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
type ScriptUpdateParamsMetadataMigrations struct {
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

func (r ScriptUpdateParamsMetadataMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataMigrations) implementsScriptUpdateParamsMetadataMigrationsUnion() {}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by [workers.SingleStepMigrationParam],
// [workers.ScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations],
// [ScriptUpdateParamsMetadataMigrations].
type ScriptUpdateParamsMetadataMigrationsUnion interface {
	implementsScriptUpdateParamsMetadataMigrationsUnion()
}

type ScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]MigrationStepParam] `json:"steps"`
}

func (r ScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations) implementsScriptUpdateParamsMetadataMigrationsUnion() {
}

// Observability settings for the Worker.
type ScriptUpdateParamsMetadataObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
}

func (r ScriptUpdateParamsMetadataObservability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptUpdateParamsMetadataPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode param.Field[ScriptUpdateParamsMetadataPlacementMode] `json:"mode"`
}

func (r ScriptUpdateParamsMetadataPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptUpdateParamsMetadataPlacementMode string

const (
	ScriptUpdateParamsMetadataPlacementModeSmart ScriptUpdateParamsMetadataPlacementMode = "smart"
)

func (r ScriptUpdateParamsMetadataPlacementMode) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataPlacementModeSmart:
		return true
	}
	return false
}

// Status of
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptUpdateParamsMetadataPlacementStatus string

const (
	ScriptUpdateParamsMetadataPlacementStatusSuccess                 ScriptUpdateParamsMetadataPlacementStatus = "SUCCESS"
	ScriptUpdateParamsMetadataPlacementStatusUnsupportedApplication  ScriptUpdateParamsMetadataPlacementStatus = "UNSUPPORTED_APPLICATION"
	ScriptUpdateParamsMetadataPlacementStatusInsufficientInvocations ScriptUpdateParamsMetadataPlacementStatus = "INSUFFICIENT_INVOCATIONS"
)

func (r ScriptUpdateParamsMetadataPlacementStatus) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataPlacementStatusSuccess, ScriptUpdateParamsMetadataPlacementStatusUnsupportedApplication, ScriptUpdateParamsMetadataPlacementStatusInsufficientInvocations:
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ScriptUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptUpdateResponse                `json:"result"`
	JSON    scriptUpdateResponseEnvelopeJSON    `json:"-"`
}

// scriptUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptUpdateResponseEnvelope]
type scriptUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
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
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptDeleteParams struct {
	// Identifier
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

type ScriptGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
