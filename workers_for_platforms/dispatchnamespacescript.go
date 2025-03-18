// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

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
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/cloudflare/cloudflare-go/v4/workers"
)

// DispatchNamespaceScriptService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDispatchNamespaceScriptService] method instead.
type DispatchNamespaceScriptService struct {
	Options     []option.RequestOption
	AssetUpload *DispatchNamespaceScriptAssetUploadService
	Content     *DispatchNamespaceScriptContentService
	Settings    *DispatchNamespaceScriptSettingService
	Bindings    *DispatchNamespaceScriptBindingService
	Secrets     *DispatchNamespaceScriptSecretService
	Tags        *DispatchNamespaceScriptTagService
}

// NewDispatchNamespaceScriptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptService(opts ...option.RequestOption) (r *DispatchNamespaceScriptService) {
	r = &DispatchNamespaceScriptService{}
	r.Options = opts
	r.AssetUpload = NewDispatchNamespaceScriptAssetUploadService(opts...)
	r.Content = NewDispatchNamespaceScriptContentService(opts...)
	r.Settings = NewDispatchNamespaceScriptSettingService(opts...)
	r.Bindings = NewDispatchNamespaceScriptBindingService(opts...)
	r.Secrets = NewDispatchNamespaceScriptSecretService(opts...)
	r.Tags = NewDispatchNamespaceScriptTagService(opts...)
	return
}

// Upload a worker module to a Workers for Platforms namespace. You can find more
// about the multipart metadata on our docs:
// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.
func (r *DispatchNamespaceScriptService) Update(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptUpdateParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptUpdateResponse, err error) {
	var env DispatchNamespaceScriptUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a worker from a Workers for Platforms namespace. This call has no
// response body on a successful delete.
func (r *DispatchNamespaceScriptService) Delete(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

// Fetch information about a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptGetParams, opts ...option.RequestOption) (res *Script, err error) {
	var env DispatchNamespaceScriptGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Details about a worker uploaded to a Workers for Platforms namespace.
type Script struct {
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Name of the Workers for Platforms dispatch namespace.
	DispatchNamespace string `json:"dispatch_namespace"`
	// When the script was last modified.
	ModifiedOn time.Time      `json:"modified_on" format:"date-time"`
	Script     workers.Script `json:"script"`
	JSON       scriptJSON     `json:"-"`
}

// scriptJSON contains the JSON metadata for the struct [Script]
type scriptJSON struct {
	CreatedOn         apijson.Field
	DispatchNamespace apijson.Field
	ModifiedOn        apijson.Field
	Script            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Script) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptUpdateResponse struct {
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
	Placement DispatchNamespaceScriptUpdateResponsePlacement `json:"placement"`
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	//
	// Deprecated: deprecated
	PlacementMode DispatchNamespaceScriptUpdateResponsePlacementMode `json:"placement_mode"`
	// Status of
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	//
	// Deprecated: deprecated
	PlacementStatus DispatchNamespaceScriptUpdateResponsePlacementStatus `json:"placement_status"`
	StartupTimeMs   int64                                                `json:"startup_time_ms"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []workers.ConsumerScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel DispatchNamespaceScriptUpdateResponseUsageModel `json:"usage_model"`
	JSON       dispatchNamespaceScriptUpdateResponseJSON       `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseJSON contains the JSON metadata for the
// struct [DispatchNamespaceScriptUpdateResponse]
type dispatchNamespaceScriptUpdateResponseJSON struct {
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

func (r *DispatchNamespaceScriptUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode DispatchNamespaceScriptUpdateResponsePlacementMode `json:"mode"`
	// Status of
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Status DispatchNamespaceScriptUpdateResponsePlacementStatus `json:"status"`
	JSON   dispatchNamespaceScriptUpdateResponsePlacementJSON   `json:"-"`
}

// dispatchNamespaceScriptUpdateResponsePlacementJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptUpdateResponsePlacement]
type dispatchNamespaceScriptUpdateResponsePlacementJSON struct {
	Mode        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateResponsePlacementMode string

const (
	DispatchNamespaceScriptUpdateResponsePlacementModeSmart DispatchNamespaceScriptUpdateResponsePlacementMode = "smart"
)

func (r DispatchNamespaceScriptUpdateResponsePlacementMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponsePlacementModeSmart:
		return true
	}
	return false
}

// Status of
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateResponsePlacementStatus string

const (
	DispatchNamespaceScriptUpdateResponsePlacementStatusSuccess                 DispatchNamespaceScriptUpdateResponsePlacementStatus = "SUCCESS"
	DispatchNamespaceScriptUpdateResponsePlacementStatusUnsupportedApplication  DispatchNamespaceScriptUpdateResponsePlacementStatus = "UNSUPPORTED_APPLICATION"
	DispatchNamespaceScriptUpdateResponsePlacementStatusInsufficientInvocations DispatchNamespaceScriptUpdateResponsePlacementStatus = "INSUFFICIENT_INVOCATIONS"
)

func (r DispatchNamespaceScriptUpdateResponsePlacementStatus) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponsePlacementStatusSuccess, DispatchNamespaceScriptUpdateResponsePlacementStatusUnsupportedApplication, DispatchNamespaceScriptUpdateResponsePlacementStatusInsufficientInvocations:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type DispatchNamespaceScriptUpdateResponseUsageModel string

const (
	DispatchNamespaceScriptUpdateResponseUsageModelStandard DispatchNamespaceScriptUpdateResponseUsageModel = "standard"
)

func (r DispatchNamespaceScriptUpdateResponseUsageModel) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseUsageModelStandard:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[DispatchNamespaceScriptUpdateParamsMetadata] `json:"metadata,required"`
}

func (r DispatchNamespaceScriptUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// JSON encoded metadata about the uploaded parts and Worker configuration.
type DispatchNamespaceScriptUpdateParamsMetadata struct {
	// Configuration for assets within a Worker
	Assets param.Field[DispatchNamespaceScriptUpdateParamsMetadataAssets] `json:"assets"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings param.Field[[]DispatchNamespaceScriptUpdateParamsMetadataBindingUnion] `json:"bindings"`
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
	Migrations param.Field[DispatchNamespaceScriptUpdateParamsMetadataMigrationsUnion] `json:"migrations"`
	// Observability settings for the Worker.
	Observability param.Field[DispatchNamespaceScriptUpdateParamsMetadataObservability] `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement param.Field[DispatchNamespaceScriptUpdateParamsMetadataPlacement] `json:"placement"`
	// List of strings to use as tags for this Worker.
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]workers.ConsumerScriptParam] `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[DispatchNamespaceScriptUpdateParamsMetadataUsageModel] `json:"usage_model"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker
type DispatchNamespaceScriptUpdateParamsMetadataAssets struct {
	// Configuration for assets within a Worker.
	Config param.Field[DispatchNamespaceScriptUpdateParamsMetadataAssetsConfig] `json:"config"`
	// Token provided upon successful upload of all files from a registered manifest.
	JWT param.Field[string] `json:"jwt"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker.
type DispatchNamespaceScriptUpdateParamsMetadataAssetsConfig struct {
	// Determines the redirects and rewrites of requests for HTML content.
	HTMLHandling param.Field[DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling] `json:"html_handling"`
	// Determines the response when a request does not match a static asset, and there
	// is no Worker script.
	NotFoundHandling param.Field[DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling] `json:"not_found_handling"`
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

func (r DispatchNamespaceScriptUpdateParamsMetadataAssetsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the redirects and rewrites of requests for HTML content.
type DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling string

const (
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingAutoTrailingSlash  DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "auto-trailing-slash"
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingForceTrailingSlash DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "force-trailing-slash"
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingDropTrailingSlash  DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "drop-trailing-slash"
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingNone               DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "none"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingAutoTrailingSlash, DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingForceTrailingSlash, DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingDropTrailingSlash, DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingNone:
		return true
	}
	return false
}

// Determines the response when a request does not match a static asset, and there
// is no Worker script.
type DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling string

const (
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingNone                  DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling = "none"
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling404Page               DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling = "404-page"
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingSinglePageApplication DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling = "single-page-application"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingNone, DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling404Page, DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingSinglePageApplication:
		return true
	}
	return false
}

// A binding to allow the Worker to communicate with resources
type DispatchNamespaceScriptUpdateParamsMetadataBinding struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsType] `json:"type,required"`
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

func (r DispatchNamespaceScriptUpdateParamsMetadataBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBinding) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAI],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRendering],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJson],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindService],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumer],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata],
// [DispatchNamespaceScriptUpdateParamsMetadataBinding].
type DispatchNamespaceScriptUpdateParamsMetadataBindingUnion interface {
	implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion()
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAI) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRendering struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRendering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRendering) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace to bind to.
	Namespace param.Field[string] `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType] `json:"type,required"`
	// Outbound worker.
	Outbound param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound] `json:"outbound"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

// Outbound worker.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params param.Field[[]string] `json:"params"`
	// Outbound worker.
	Worker param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker] `json:"worker"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker.
	Service param.Field[string] `json:"service"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace struct {
	// The exported class name of the Durable Object.
	ClassName param.Field[string] `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType] `json:"type,required"`
	// The environment of the script_name to bind to.
	Environment param.Field[string] `json:"environment"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json param.Field[string] `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJson) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The text value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The secret value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindService) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumer) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTailConsumerTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeBrowserRendering       DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "browser_rendering"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeBrowserRendering, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeVersionMetadata:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
type DispatchNamespaceScriptUpdateParamsMetadataMigrations struct {
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

func (r DispatchNamespaceScriptUpdateParamsMetadataMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataMigrations) ImplementsDispatchNamespaceScriptUpdateParamsMetadataMigrationsUnion() {
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by [workers.SingleStepMigrationParam],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations],
// [DispatchNamespaceScriptUpdateParamsMetadataMigrations].
type DispatchNamespaceScriptUpdateParamsMetadataMigrationsUnion interface {
	ImplementsDispatchNamespaceScriptUpdateParamsMetadataMigrationsUnion()
}

type DispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]workers.MigrationStepParam] `json:"steps"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations) ImplementsDispatchNamespaceScriptUpdateParamsMetadataMigrationsUnion() {
}

// Observability settings for the Worker.
type DispatchNamespaceScriptUpdateParamsMetadataObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObservability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateParamsMetadataPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode param.Field[DispatchNamespaceScriptUpdateParamsMetadataPlacementMode] `json:"mode"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateParamsMetadataPlacementMode string

const (
	DispatchNamespaceScriptUpdateParamsMetadataPlacementModeSmart DispatchNamespaceScriptUpdateParamsMetadataPlacementMode = "smart"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataPlacementMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataPlacementModeSmart:
		return true
	}
	return false
}

// Status of
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateParamsMetadataPlacementStatus string

const (
	DispatchNamespaceScriptUpdateParamsMetadataPlacementStatusSuccess                 DispatchNamespaceScriptUpdateParamsMetadataPlacementStatus = "SUCCESS"
	DispatchNamespaceScriptUpdateParamsMetadataPlacementStatusUnsupportedApplication  DispatchNamespaceScriptUpdateParamsMetadataPlacementStatus = "UNSUPPORTED_APPLICATION"
	DispatchNamespaceScriptUpdateParamsMetadataPlacementStatusInsufficientInvocations DispatchNamespaceScriptUpdateParamsMetadataPlacementStatus = "INSUFFICIENT_INVOCATIONS"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataPlacementStatus) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataPlacementStatusSuccess, DispatchNamespaceScriptUpdateParamsMetadataPlacementStatusUnsupportedApplication, DispatchNamespaceScriptUpdateParamsMetadataPlacementStatusInsufficientInvocations:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type DispatchNamespaceScriptUpdateParamsMetadataUsageModel string

const (
	DispatchNamespaceScriptUpdateParamsMetadataUsageModelStandard DispatchNamespaceScriptUpdateParamsMetadataUsageModel = "standard"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataUsageModel) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataUsageModelStandard:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceScriptUpdateResponse                `json:"result"`
	JSON    dispatchNamespaceScriptUpdateResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptUpdateResponseEnvelope]
type dispatchNamespaceScriptUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptUpdateResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptUpdateResponseEnvelopeSuccessTrue DispatchNamespaceScriptUpdateResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// If set to true, delete will not be stopped by associated service binding,
	// durable object, or other binding. Any of these associated bindings/durable
	// objects will be deleted along with the script.
	Force param.Field[bool] `query:"force"`
}

// URLQuery serializes [DispatchNamespaceScriptDeleteParams]'s query parameters as
// `url.Values`.
func (r DispatchNamespaceScriptDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DispatchNamespaceScriptGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceScriptGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptGetResponseEnvelopeSuccess `json:"success,required"`
	// Details about a worker uploaded to a Workers for Platforms namespace.
	Result Script                                         `json:"result"`
	JSON   dispatchNamespaceScriptGetResponseEnvelopeJSON `json:"-"`
}

// dispatchNamespaceScriptGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptGetResponseEnvelope]
type dispatchNamespaceScriptGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptGetResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptGetResponseEnvelopeSuccessTrue DispatchNamespaceScriptGetResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
