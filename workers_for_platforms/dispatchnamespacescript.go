// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

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
	"github.com/cloudflare/cloudflare-go/v5/workers"
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
func (r *DispatchNamespaceScriptService) Delete(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptDeleteParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptDeleteResponse, err error) {
	var env DispatchNamespaceScriptDeleteResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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
	Placement DispatchNamespaceScriptUpdateResponsePlacement `json:"placement"`
	// Deprecated: deprecated
	PlacementMode DispatchNamespaceScriptUpdateResponsePlacementMode `json:"placement_mode"`
	// Deprecated: deprecated
	PlacementStatus DispatchNamespaceScriptUpdateResponsePlacementStatus `json:"placement_status"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []workers.ConsumerScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel DispatchNamespaceScriptUpdateResponseUsageModel `json:"usage_model"`
	JSON       dispatchNamespaceScriptUpdateResponseJSON       `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseJSON contains the JSON metadata for the
// struct [DispatchNamespaceScriptUpdateResponse]
type dispatchNamespaceScriptUpdateResponseJSON struct {
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

func (r *DispatchNamespaceScriptUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateResponsePlacement struct {
	// The last time the script was analyzed for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	LastAnalyzedAt time.Time `json:"last_analyzed_at" format:"date-time"`
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
	LastAnalyzedAt apijson.Field
	Mode           apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
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

type DispatchNamespaceScriptDeleteResponse = interface{}

type DispatchNamespaceScriptUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// JSON-encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[DispatchNamespaceScriptUpdateParamsMetadataUnion] `json:"metadata,required"`
	// An array of modules (often JavaScript files) comprising a Worker script. At
	// least one module must be present and referenced in the metadata as `main_module`
	// or `body_part` by filename.<br/>Possible Content-Type(s) are:
	// `application/javascript+module`, `text/javascript+module`,
	// `application/javascript`, `text/javascript`, `text/x-python`,
	// `text/x-python-requirement`, `application/wasm`, `text/plain`,
	// `application/octet-stream`, `application/source-map`.
	Files param.Field[[]io.Reader] `json:"files" format:"binary"`
}

func (r DispatchNamespaceScriptUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type DispatchNamespaceScriptUpdateParamsMetadata struct {
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
	UsageModel param.Field[DispatchNamespaceScriptUpdateParamsMetadataUsageModel] `json:"usage_model"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadata) implementsDispatchNamespaceScriptUpdateParamsMetadataUnion() {
}

// JSON-encoded metadata about the uploaded parts and Worker configuration.
//
// Satisfied by
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObject],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObject],
// [DispatchNamespaceScriptUpdateParamsMetadata].
type DispatchNamespaceScriptUpdateParamsMetadataUnion interface {
	implementsDispatchNamespaceScriptUpdateParamsMetadataUnion()
}

type DispatchNamespaceScriptUpdateParamsMetadataObject struct {
	// Name of the uploaded file that contains the main module (e.g. the file exporting
	// a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module,required"`
	// Configuration for assets within a Worker.
	Assets param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectAssets] `json:"assets"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings param.Field[[]DispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion] `json:"bindings"`
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
	Migrations param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectMigrationsUnion] `json:"migrations"`
	// Observability settings for the Worker.
	Observability param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectObservability] `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectPlacement] `json:"placement"`
	// List of strings to use as tags for this Worker.
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]workers.ConsumerScriptParam] `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectUsageModel] `json:"usage_model"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObject) implementsDispatchNamespaceScriptUpdateParamsMetadataUnion() {
}

// Configuration for assets within a Worker.
type DispatchNamespaceScriptUpdateParamsMetadataObjectAssets struct {
	// Configuration for assets within a Worker.
	Config param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfig] `json:"config"`
	// Token provided upon successful upload of all files from a registered manifest.
	JWT param.Field[string] `json:"jwt"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker.
type DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfig struct {
	// The contents of a \_headers file (used to attach custom headers on asset
	// responses).
	Headers param.Field[string] `json:"_headers"`
	// The contents of a \_redirects file (used to apply redirects or proxy paths ahead
	// of asset serving).
	Redirects param.Field[string] `json:"_redirects"`
	// Determines the redirects and rewrites of requests for HTML content.
	HTMLHandling param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandling] `json:"html_handling"`
	// Determines the response when a request does not match a static asset, and there
	// is no Worker script.
	NotFoundHandling param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandling] `json:"not_found_handling"`
	// Contains a list path rules to control routing to either the Worker or assets.
	// Glob (\*) and negative (!) rules are supported. Rules must start with either '/'
	// or '!/'. At least one non-negative rule must be provided, and negative rules
	// have higher precedence than non-negative rules.
	RunWorkerFirst param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigRunWorkerFirstUnion] `json:"run_worker_first"`
	// When true and the incoming request matches an asset, that will be served instead
	// of invoking the Worker script. When false, requests will always invoke the
	// Worker script.
	//
	// Deprecated: deprecated
	ServeDirectly param.Field[bool] `json:"serve_directly"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the redirects and rewrites of requests for HTML content.
type DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandling string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandlingAutoTrailingSlash  DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandling = "auto-trailing-slash"
	DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandlingForceTrailingSlash DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandling = "force-trailing-slash"
	DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandlingDropTrailingSlash  DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandling = "drop-trailing-slash"
	DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandlingNone               DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandling = "none"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandling) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandlingAutoTrailingSlash, DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandlingForceTrailingSlash, DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandlingDropTrailingSlash, DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigHTMLHandlingNone:
		return true
	}
	return false
}

// Determines the response when a request does not match a static asset, and there
// is no Worker script.
type DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandling string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandlingNone                  DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandling = "none"
	DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandling404Page               DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandling = "404-page"
	DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandlingSinglePageApplication DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandling = "single-page-application"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandling) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandlingNone, DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandling404Page, DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigNotFoundHandlingSinglePageApplication:
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
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigRunWorkerFirstArray],
// [shared.UnionBool].
type DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigRunWorkerFirstUnion interface {
	ImplementsDispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigRunWorkerFirstUnion()
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigRunWorkerFirstArray []string

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigRunWorkerFirstArray) ImplementsDispatchNamespaceScriptUpdateParamsMetadataObjectAssetsConfigRunWorkerFirstUnion() {
}

// A binding to allow the Worker to communicate with resources.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBinding struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType] `json:"type,required"`
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
	Format param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsFormat] `json:"format"`
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

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBinding) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// A binding to allow the Worker to communicate with resources.
//
// Satisfied by
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAI],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngine],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssets],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowser],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespace],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespace],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdrive],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJson],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespace],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificate],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainText],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelines],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueue],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2Bucket],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretText],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindService],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumer],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorize],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadata],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecret],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKey],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflow],
// [DispatchNamespaceScriptUpdateParamsMetadataObjectBinding].
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion interface {
	implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion()
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAIType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAI) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAIType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAITypeAI DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAIType = "ai"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngineType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngine) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngine) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngineType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssetsType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssets) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssetsType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssetsTypeAssets DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssetsType = "assets"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowserType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowser) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowserType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowserTypeBrowser DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowserType = "browser"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1Type] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1Type string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1TypeD1 DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1Type = "d1"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace to bind to.
	Namespace param.Field[string] `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceType] `json:"type,required"`
	// Outbound worker.
	Outbound param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceOutbound] `json:"outbound"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespace) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params param.Field[[]string] `json:"params"`
	// Outbound worker.
	Worker param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceOutboundWorker] `json:"worker"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker.
	Service param.Field[string] `json:"service"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespaceType] `json:"type,required"`
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

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespace) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdriveType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdrive) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdrive) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdriveType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdriveTypeHyperdrive DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json param.Field[string] `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJsonType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJson) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJsonType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJsonTypeJson DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJsonType = "json"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespaceType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespace) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespaceType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespaceTypeKVNamespace DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificateType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificate) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificateType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The text value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainTextType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainText) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainTextType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainTextTypePlainText DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline param.Field[string] `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelinesType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelines) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelines) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelinesType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelinesTypePipelines DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelinesType = "pipelines"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueueType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueue) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueueType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueueTypeQueue DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueueType = "queue"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2BucketType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2Bucket) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2Bucket) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2BucketType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2BucketTypeR2Bucket DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The secret value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretTextType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretText) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretTextType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretTextTypeSecretText DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretTextType = "secret_text"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindServiceType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindService) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindServiceType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindServiceTypeService DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindServiceType = "service"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumerType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumer) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumerType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumerTypeTailConsumer DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindTailConsumerTypeTailConsumer:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorizeType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorize) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorizeType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorizeTypeVectorize DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorizeType = "vectorize"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadataType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadata) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadataType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the secret in the store.
	SecretName param.Field[string] `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID param.Field[string] `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecretType] `json:"type,required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecret) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecretType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm param.Field[interface{}] `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormat] `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyType] `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages param.Field[[]DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage] `json:"usages,required"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string] `json:"key_base64"`
	// Key data in
	// [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key)
	// format. Required if `format` is "jwk".
	KeyJwk param.Field[interface{}] `json:"key_jwk"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKey) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormat string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormatRaw   DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormat = "raw"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormatPkcs8 DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormat = "pkcs8"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormatSpki  DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormat = "spki"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormatJwk   DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormatRaw, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormatPkcs8, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormatSpki, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyTypeSecretKey DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyType = "secret_key"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageEncrypt    DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage = "encrypt"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageDecrypt    DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage = "decrypt"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageSign       DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage = "sign"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageVerify     DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage = "verify"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageDeriveKey  DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage = "deriveKey"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageDeriveBits DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage = "deriveBits"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageWrapKey    DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage = "wrapKey"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageUnwrapKey  DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageEncrypt, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageDecrypt, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageSign, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageVerify, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageDeriveKey, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageDeriveBits, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageWrapKey, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflowType] `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName param.Field[string] `json:"workflow_name,required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName param.Field[string] `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflow) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflow) implementsDispatchNamespaceScriptUpdateParamsMetadataObjectBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflowType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflowTypeWorkflow DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflowType = "workflow"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeBrowser                DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "browser"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypePipelines              DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "pipelines"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeService                DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeTailConsumer           DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "tail_consumer"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "version_metadata"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeSecretsStoreSecret     DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "secrets_store_secret"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeSecretKey              DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "secret_key"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeWorkflow               DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType = "workflow"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeAI, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeBrowser, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeD1, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeJson, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypePipelines, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeService, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeTailConsumer, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeVersionMetadata, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeSecretsStoreSecret, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeSecretKey, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsTypeWorkflow:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsFormat string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsFormatRaw   DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsFormat = "raw"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsFormatPkcs8 DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsFormat = "pkcs8"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsFormatSpki  DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsFormat = "spki"
	DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsFormatJwk   DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsFormat = "jwk"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsFormatRaw, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsFormatPkcs8, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsFormatSpki, DispatchNamespaceScriptUpdateParamsMetadataObjectBindingsFormatJwk:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
type DispatchNamespaceScriptUpdateParamsMetadataObjectMigrations struct {
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

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectMigrations) ImplementsDispatchNamespaceScriptUpdateParamsMetadataObjectMigrationsUnion() {
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by [workers.SingleStepMigrationParam],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataObjectMigrationsWorkersMultipleStepMigrations],
// [DispatchNamespaceScriptUpdateParamsMetadataObjectMigrations].
type DispatchNamespaceScriptUpdateParamsMetadataObjectMigrationsUnion interface {
	ImplementsDispatchNamespaceScriptUpdateParamsMetadataObjectMigrationsUnion()
}

type DispatchNamespaceScriptUpdateParamsMetadataObjectMigrationsWorkersMultipleStepMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]workers.MigrationStepParam] `json:"steps"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectMigrationsWorkersMultipleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectMigrationsWorkersMultipleStepMigrations) ImplementsDispatchNamespaceScriptUpdateParamsMetadataObjectMigrationsUnion() {
}

// Observability settings for the Worker.
type DispatchNamespaceScriptUpdateParamsMetadataObjectObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
	// Log settings for the Worker.
	Logs param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectObservabilityLogs] `json:"logs"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectObservability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Log settings for the Worker.
type DispatchNamespaceScriptUpdateParamsMetadataObjectObservabilityLogs struct {
	// Whether logs are enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs param.Field[bool] `json:"invocation_logs,required"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectObservabilityLogs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateParamsMetadataObjectPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode param.Field[DispatchNamespaceScriptUpdateParamsMetadataObjectPlacementMode] `json:"mode"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateParamsMetadataObjectPlacementMode string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectPlacementModeSmart DispatchNamespaceScriptUpdateParamsMetadataObjectPlacementMode = "smart"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectPlacementMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectPlacementModeSmart:
		return true
	}
	return false
}

// Status of
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateParamsMetadataObjectPlacementStatus string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectPlacementStatusSuccess                 DispatchNamespaceScriptUpdateParamsMetadataObjectPlacementStatus = "SUCCESS"
	DispatchNamespaceScriptUpdateParamsMetadataObjectPlacementStatusUnsupportedApplication  DispatchNamespaceScriptUpdateParamsMetadataObjectPlacementStatus = "UNSUPPORTED_APPLICATION"
	DispatchNamespaceScriptUpdateParamsMetadataObjectPlacementStatusInsufficientInvocations DispatchNamespaceScriptUpdateParamsMetadataObjectPlacementStatus = "INSUFFICIENT_INVOCATIONS"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectPlacementStatus) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectPlacementStatusSuccess, DispatchNamespaceScriptUpdateParamsMetadataObjectPlacementStatusUnsupportedApplication, DispatchNamespaceScriptUpdateParamsMetadataObjectPlacementStatusInsufficientInvocations:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type DispatchNamespaceScriptUpdateParamsMetadataObjectUsageModel string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObjectUsageModelStandard DispatchNamespaceScriptUpdateParamsMetadataObjectUsageModel = "standard"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObjectUsageModel) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObjectUsageModelStandard:
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
	Errors   []DispatchNamespaceScriptUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceScriptUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   DispatchNamespaceScriptUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success DispatchNamespaceScriptUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatchNamespaceScriptUpdateResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptUpdateResponseEnvelope]
type dispatchNamespaceScriptUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptUpdateResponseEnvelopeErrors struct {
	Code             int64                                                     `json:"code,required"`
	Message          string                                                    `json:"message,required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           DispatchNamespaceScriptUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dispatchNamespaceScriptUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptUpdateResponseEnvelopeErrors]
type dispatchNamespaceScriptUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    dispatchNamespaceScriptUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptUpdateResponseEnvelopeErrorsSource]
type dispatchNamespaceScriptUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptUpdateResponseEnvelopeMessages struct {
	Code             int64                                                       `json:"code,required"`
	Message          string                                                      `json:"message,required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           DispatchNamespaceScriptUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dispatchNamespaceScriptUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptUpdateResponseEnvelopeMessages]
type dispatchNamespaceScriptUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    dispatchNamespaceScriptUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptUpdateResponseEnvelopeMessagesSource]
type dispatchNamespaceScriptUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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
	// Identifier.
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

type DispatchNamespaceScriptDeleteResponseEnvelope struct {
	Errors   []DispatchNamespaceScriptDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceScriptDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DispatchNamespaceScriptDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceScriptDeleteResponse                `json:"result,nullable"`
	JSON    dispatchNamespaceScriptDeleteResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptDeleteResponseEnvelope]
type dispatchNamespaceScriptDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptDeleteResponseEnvelopeErrors struct {
	Code             int64                                                     `json:"code,required"`
	Message          string                                                    `json:"message,required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           DispatchNamespaceScriptDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dispatchNamespaceScriptDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dispatchNamespaceScriptDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptDeleteResponseEnvelopeErrors]
type dispatchNamespaceScriptDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    dispatchNamespaceScriptDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceScriptDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptDeleteResponseEnvelopeErrorsSource]
type dispatchNamespaceScriptDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptDeleteResponseEnvelopeMessages struct {
	Code             int64                                                       `json:"code,required"`
	Message          string                                                      `json:"message,required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           DispatchNamespaceScriptDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dispatchNamespaceScriptDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dispatchNamespaceScriptDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptDeleteResponseEnvelopeMessages]
type dispatchNamespaceScriptDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    dispatchNamespaceScriptDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceScriptDeleteResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptDeleteResponseEnvelopeMessagesSource]
type dispatchNamespaceScriptDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DispatchNamespaceScriptDeleteResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptDeleteResponseEnvelopeSuccessTrue DispatchNamespaceScriptDeleteResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceScriptGetResponseEnvelope struct {
	Errors   []DispatchNamespaceScriptGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceScriptGetResponseEnvelopeMessages `json:"messages,required"`
	// Details about a worker uploaded to a Workers for Platforms namespace.
	Result Script `json:"result,required"`
	// Whether the API call was successful.
	Success DispatchNamespaceScriptGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatchNamespaceScriptGetResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptGetResponseEnvelope]
type dispatchNamespaceScriptGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptGetResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           DispatchNamespaceScriptGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dispatchNamespaceScriptGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dispatchNamespaceScriptGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptGetResponseEnvelopeErrors]
type dispatchNamespaceScriptGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    dispatchNamespaceScriptGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceScriptGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptGetResponseEnvelopeErrorsSource]
type dispatchNamespaceScriptGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptGetResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           DispatchNamespaceScriptGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dispatchNamespaceScriptGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dispatchNamespaceScriptGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptGetResponseEnvelopeMessages]
type dispatchNamespaceScriptGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    dispatchNamespaceScriptGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceScriptGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptGetResponseEnvelopeMessagesSource]
type dispatchNamespaceScriptGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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
