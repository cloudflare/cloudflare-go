// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v3/shared"
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
	PlacementMode ScriptPlacementMode `json:"placement_mode"`
	// Status of
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
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
	ScriptPlacementStatusNoValidHosts            ScriptPlacementStatus = "NO_VALID_HOSTS"
	ScriptPlacementStatusNoValidBindings         ScriptPlacementStatus = "NO_VALID_BINDINGS"
	ScriptPlacementStatusUnsupportedApplication  ScriptPlacementStatus = "UNSUPPORTED_APPLICATION"
	ScriptPlacementStatusInsufficientInvocations ScriptPlacementStatus = "INSUFFICIENT_INVOCATIONS"
	ScriptPlacementStatusInsufficientSubrequests ScriptPlacementStatus = "INSUFFICIENT_SUBREQUESTS"
)

func (r ScriptPlacementStatus) IsKnown() bool {
	switch r {
	case ScriptPlacementStatusSuccess, ScriptPlacementStatusNoValidHosts, ScriptPlacementStatusNoValidBindings, ScriptPlacementStatusUnsupportedApplication, ScriptPlacementStatusInsufficientInvocations, ScriptPlacementStatusInsufficientSubrequests:
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
	PlacementMode ScriptUpdateResponsePlacementMode `json:"placement_mode"`
	// Status of
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
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
	ScriptUpdateResponsePlacementStatusNoValidHosts            ScriptUpdateResponsePlacementStatus = "NO_VALID_HOSTS"
	ScriptUpdateResponsePlacementStatusNoValidBindings         ScriptUpdateResponsePlacementStatus = "NO_VALID_BINDINGS"
	ScriptUpdateResponsePlacementStatusUnsupportedApplication  ScriptUpdateResponsePlacementStatus = "UNSUPPORTED_APPLICATION"
	ScriptUpdateResponsePlacementStatusInsufficientInvocations ScriptUpdateResponsePlacementStatus = "INSUFFICIENT_INVOCATIONS"
	ScriptUpdateResponsePlacementStatusInsufficientSubrequests ScriptUpdateResponsePlacementStatus = "INSUFFICIENT_SUBREQUESTS"
)

func (r ScriptUpdateResponsePlacementStatus) IsKnown() bool {
	switch r {
	case ScriptUpdateResponsePlacementStatusSuccess, ScriptUpdateResponsePlacementStatusNoValidHosts, ScriptUpdateResponsePlacementStatusNoValidBindings, ScriptUpdateResponsePlacementStatusUnsupportedApplication, ScriptUpdateResponsePlacementStatusInsufficientInvocations, ScriptUpdateResponsePlacementStatusInsufficientSubrequests:
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
	// When true and the incoming request matches an asset, that will be served instead
	// of invoking the Worker script. When false, requests will always invoke the
	// Worker script.
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
	Type param.Field[string] `json:"type,required"`
	// Identifier of the D1 database to bind to.
	ID param.Field[string] `json:"id"`
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name"`
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id"`
	// The exported class name of the Durable Object.
	ClassName param.Field[string] `json:"class_name"`
	// The dataset name to bind to.
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

func (r ScriptUpdateParamsMetadataBinding) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindAny],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindAI],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRendering],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindDo],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindJson],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCERT],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecret],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindService],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumer],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize],
// [workers.ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata],
// [ScriptUpdateParamsMetadataBinding].
type ScriptUpdateParamsMetadataBindingUnion interface {
	implementsWorkersScriptUpdateParamsMetadataBindingUnion()
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindAny struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type        param.Field[string]    `json:"type,required"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAny) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAny) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
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

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAI) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAI ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "ai"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine struct {
	// The dataset name to bind to.
	Dataset param.Field[string] `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
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

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAssets ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "assets"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAssets:
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

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRendering) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "browser_rendering"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering:
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

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeD1 ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "d1"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeD1:
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

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
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

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindDo struct {
	// The exported class name of the Durable Object.
	ClassName param.Field[string] `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindDoType] `json:"type,required"`
	// The environment of the script_name to bind to.
	Environment param.Field[string] `json:"environment"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDo) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindDoType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindDoTypeDurableObjectNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindDoType = "durable_object_namespace"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindDoType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindDoTypeDurableObjectNamespace:
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

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
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

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindJson) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeJson ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "json"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeJson:
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

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCERT struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCERTType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCERT) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCERT) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCERTType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCERTTypeMTLSCertificate ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCERTType = "mtls_certificate"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCERTType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCERTTypeMTLSCertificate:
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

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText:
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

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeQueue ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "queue"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2 struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Type] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Type string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2TypeR2Bucket ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Type = "r2_bucket"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Type) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindR2TypeR2Bucket:
		return true
	}
	return false
}

type ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecret struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The secret value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretType] `json:"type,required"`
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecret) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTypeSecretText ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretType = "secret_text"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTypeSecretText:
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

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindService) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeService ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "service"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeService:
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

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumer) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeTailConsumer ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeTailConsumer:
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

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeVectorize ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "vectorize"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeVectorize:
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

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata) implementsWorkersScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType string

const (
	ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
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

func (r ScriptUpdateParamsMetadataMigrations) implementsWorkersScriptUpdateParamsMetadataMigrationsUnion() {
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by [workers.SingleStepMigrationParam],
// [workers.ScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations],
// [ScriptUpdateParamsMetadataMigrations].
type ScriptUpdateParamsMetadataMigrationsUnion interface {
	implementsWorkersScriptUpdateParamsMetadataMigrationsUnion()
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

func (r ScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations) implementsWorkersScriptUpdateParamsMetadataMigrationsUnion() {
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
	ScriptUpdateParamsMetadataPlacementStatusNoValidHosts            ScriptUpdateParamsMetadataPlacementStatus = "NO_VALID_HOSTS"
	ScriptUpdateParamsMetadataPlacementStatusNoValidBindings         ScriptUpdateParamsMetadataPlacementStatus = "NO_VALID_BINDINGS"
	ScriptUpdateParamsMetadataPlacementStatusUnsupportedApplication  ScriptUpdateParamsMetadataPlacementStatus = "UNSUPPORTED_APPLICATION"
	ScriptUpdateParamsMetadataPlacementStatusInsufficientInvocations ScriptUpdateParamsMetadataPlacementStatus = "INSUFFICIENT_INVOCATIONS"
	ScriptUpdateParamsMetadataPlacementStatusInsufficientSubrequests ScriptUpdateParamsMetadataPlacementStatus = "INSUFFICIENT_SUBREQUESTS"
)

func (r ScriptUpdateParamsMetadataPlacementStatus) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsMetadataPlacementStatusSuccess, ScriptUpdateParamsMetadataPlacementStatusNoValidHosts, ScriptUpdateParamsMetadataPlacementStatusNoValidBindings, ScriptUpdateParamsMetadataPlacementStatusUnsupportedApplication, ScriptUpdateParamsMetadataPlacementStatusInsufficientInvocations, ScriptUpdateParamsMetadataPlacementStatusInsufficientSubrequests:
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
