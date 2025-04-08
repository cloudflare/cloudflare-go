// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v4/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
)

// ScriptVersionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptVersionService] method instead.
type ScriptVersionService struct {
	Options []option.RequestOption
}

// NewScriptVersionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptVersionService(opts ...option.RequestOption) (r *ScriptVersionService) {
	r = &ScriptVersionService{}
	r.Options = opts
	return
}

// Upload a Worker Version without deploying to Cloudflare's network. You can find
// more about the multipart metadata on our docs:
// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.
func (r *ScriptVersionService) New(ctx context.Context, scriptName string, params ScriptVersionNewParams, opts ...option.RequestOption) (res *ScriptVersionNewResponse, err error) {
	var env ScriptVersionNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/versions", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List of Worker Versions. The first version in the list is the latest version.
func (r *ScriptVersionService) List(ctx context.Context, scriptName string, params ScriptVersionListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[ScriptVersionListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/versions", params.AccountID, scriptName)
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

// List of Worker Versions. The first version in the list is the latest version.
func (r *ScriptVersionService) ListAutoPaging(ctx context.Context, scriptName string, params ScriptVersionListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[ScriptVersionListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, scriptName, params, opts...))
}

// Get Version Detail
func (r *ScriptVersionService) Get(ctx context.Context, scriptName string, versionID string, query ScriptVersionGetParams, opts ...option.RequestOption) (res *ScriptVersionGetResponse, err error) {
	var env ScriptVersionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/versions/%s", query.AccountID, scriptName, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScriptVersionNewResponse struct {
	Resources     interface{}                  `json:"resources,required"`
	ID            string                       `json:"id"`
	Metadata      interface{}                  `json:"metadata"`
	Number        float64                      `json:"number"`
	StartupTimeMs int64                        `json:"startup_time_ms"`
	JSON          scriptVersionNewResponseJSON `json:"-"`
}

// scriptVersionNewResponseJSON contains the JSON metadata for the struct
// [ScriptVersionNewResponse]
type scriptVersionNewResponseJSON struct {
	Resources     apijson.Field
	ID            apijson.Field
	Metadata      apijson.Field
	Number        apijson.Field
	StartupTimeMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptVersionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionListResponse struct {
	ID       string                        `json:"id"`
	Metadata interface{}                   `json:"metadata"`
	Number   float64                       `json:"number"`
	JSON     scriptVersionListResponseJSON `json:"-"`
}

// scriptVersionListResponseJSON contains the JSON metadata for the struct
// [ScriptVersionListResponse]
type scriptVersionListResponseJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Number      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionListResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponse struct {
	Resources interface{}                  `json:"resources,required"`
	ID        string                       `json:"id"`
	Metadata  interface{}                  `json:"metadata"`
	Number    float64                      `json:"number"`
	JSON      scriptVersionGetResponseJSON `json:"-"`
}

// scriptVersionGetResponseJSON contains the JSON metadata for the struct
// [ScriptVersionGetResponse]
type scriptVersionGetResponseJSON struct {
	Resources   apijson.Field
	ID          apijson.Field
	Metadata    apijson.Field
	Number      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[ScriptVersionNewParamsMetadata] `json:"metadata,required"`
}

func (r ScriptVersionNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// JSON encoded metadata about the uploaded parts and Worker configuration.
type ScriptVersionNewParamsMetadata struct {
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker, which
	// is required for Version Upload.
	MainModule  param.Field[string]                                    `json:"main_module,required"`
	Annotations param.Field[ScriptVersionNewParamsMetadataAnnotations] `json:"annotations"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings param.Field[[]ScriptVersionNewParamsMetadataBindingUnion] `json:"bindings"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// List of binding types to keep from previous_upload.
	KeepBindings param.Field[[]string] `json:"keep_bindings"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[ScriptVersionNewParamsMetadataUsageModel] `json:"usage_model"`
}

func (r ScriptVersionNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptVersionNewParamsMetadataAnnotations struct {
	// Human-readable message about the version. Truncated to 100 bytes.
	WorkersMessage param.Field[string] `json:"workers/message"`
	// User-provided identifier for the version.
	WorkersTag param.Field[string] `json:"workers/tag"`
}

func (r ScriptVersionNewParamsMetadataAnnotations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
type ScriptVersionNewParamsMetadataBinding struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsType] `json:"type,required"`
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

func (r ScriptVersionNewParamsMetadataBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBinding) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAI],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngine],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssets],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRendering],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespace],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdrive],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJson],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespace],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificate],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainText],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueue],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2Bucket],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretText],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindService],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumer],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorize],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadata],
// [ScriptVersionNewParamsMetadataBinding].
type ScriptVersionNewParamsMetadataBindingUnion interface {
	implementsScriptVersionNewParamsMetadataBindingUnion()
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAI) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeAI                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "ai"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeAssets                 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "assets"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeD1                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "d1"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeHyperdrive             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeJson                   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "json"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeKVNamespace            ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypePlainText              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeQueue                  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "queue"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeR2Bucket               ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeSecretText             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeService                ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "service"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeTailConsumer           ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeVectorize              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeAI, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeAssets, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeD1, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeHyperdrive, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeJson, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeKVNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypePlainText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeQueue, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeR2Bucket, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeSecretText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeService, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeTailConsumer, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeVectorize, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngine) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngine) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAI                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "ai"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAssets                 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "assets"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeD1                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "d1"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeJson                   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "json"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypePlainText              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeQueue                  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "queue"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeSecretText             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeService                ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "service"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeVectorize              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAI, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAssets, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeD1, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeJson, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypePlainText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeQueue, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeSecretText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeService, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeVectorize, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssets) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeAI                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "ai"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeAssets                 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "assets"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeD1                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "d1"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeJson                   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "json"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypePlainText              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeQueue                  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "queue"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeSecretText             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeService                ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "service"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeVectorize              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeAI, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeAssets, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeD1, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeJson, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypePlainText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeQueue, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeSecretText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeService, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeVectorize, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRendering struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRendering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRendering) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeAI                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "ai"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeAssets                 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "assets"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeD1                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "d1"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeJson                   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "json"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypePlainText              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeQueue                  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "queue"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeSecretText             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeService                ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "service"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeVectorize              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeAI, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeAssets, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeD1, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeJson, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypePlainText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeQueue, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeSecretText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeService, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeVectorize, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeAI                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "ai"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeAssets                 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "assets"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeD1                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "d1"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeHyperdrive             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeJson                   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "json"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeKVNamespace            ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypePlainText              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "plain_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeQueue                  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "queue"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeR2Bucket               ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeSecretText             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "secret_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeService                ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "service"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeTailConsumer           ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeVectorize              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "vectorize"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeAI, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeAssets, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeD1, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeHyperdrive, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeJson, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeKVNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypePlainText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeQueue, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeR2Bucket, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeSecretText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeService, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeTailConsumer, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeVectorize, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace to bind to.
	Namespace param.Field[string] `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType] `json:"type,required"`
	// Outbound worker.
	Outbound param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound] `json:"outbound"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespace) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeAI                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "ai"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeAssets                 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "assets"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeD1                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "d1"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeJson                   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "json"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypePlainText              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeQueue                  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "queue"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeSecretText             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeService                ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "service"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeVectorize              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeAI, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeAssets, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeD1, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeJson, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypePlainText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeQueue, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeSecretText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeService, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeVectorize, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

// Outbound worker.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params param.Field[[]string] `json:"params"`
	// Outbound worker.
	Worker param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker] `json:"worker"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker.
	Service param.Field[string] `json:"service"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace struct {
	// The exported class name of the Durable Object.
	ClassName param.Field[string] `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType] `json:"type,required"`
	// The environment of the script_name to bind to.
	Environment param.Field[string] `json:"environment"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeAI                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "ai"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeAssets                 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "assets"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeD1                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "d1"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeJson                   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "json"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypePlainText              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeQueue                  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "queue"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeSecretText             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeService                ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "service"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeVectorize              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeAI, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeAssets, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeD1, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeJson, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypePlainText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeQueue, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeSecretText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeService, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeVectorize, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdrive) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdrive) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeAI                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "ai"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeAssets                 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "assets"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeD1                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "d1"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeJson                   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "json"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypePlainText              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeQueue                  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "queue"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeSecretText             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeService                ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "service"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeVectorize              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeAI, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeAssets, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeD1, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeJson, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypePlainText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeQueue, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeSecretText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeService, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeVectorize, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json param.Field[string] `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJson) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeAI                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "ai"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeAssets                 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "assets"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeD1                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "d1"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeJson                   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "json"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypePlainText              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeQueue                  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "queue"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeSecretText             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeService                ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "service"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeVectorize              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeAI, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeAssets, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeD1, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeJson, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypePlainText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeQueue, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeSecretText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeService, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeVectorize, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespace) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeAI                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "ai"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeAssets                 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "assets"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeD1                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "d1"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeJson                   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "json"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypePlainText              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeQueue                  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "queue"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeSecretText             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeService                ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "service"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeVectorize              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeAI, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeAssets, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeD1, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeJson, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypePlainText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeQueue, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeSecretText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeService, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeVectorize, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificate) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeAI                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "ai"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeAssets                 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "assets"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeD1                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "d1"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeJson                   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "json"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypePlainText              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeQueue                  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "queue"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeSecretText             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeService                ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "service"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeVectorize              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeAI, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeAssets, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeD1, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeJson, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypePlainText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeQueue, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeSecretText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeService, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeVectorize, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The text value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainText) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeAI                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "ai"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeAssets                 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "assets"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeD1                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "d1"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeJson                   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "json"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeQueue                  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "queue"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeSecretText             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeService                ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "service"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeVectorize              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeAI, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeAssets, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeD1, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeJson, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeQueue, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeSecretText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeService, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeVectorize, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueue) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeAI                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "ai"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeAssets                 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "assets"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeD1                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "d1"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeJson                   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "json"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypePlainText              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeQueue                  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "queue"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeSecretText             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeService                ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "service"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeVectorize              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeAI, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeAssets, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeD1, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeJson, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypePlainText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeQueue, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeSecretText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeService, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeVectorize, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2Bucket) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2Bucket) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeAI                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "ai"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeAssets                 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "assets"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeD1                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "d1"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeJson                   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "json"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypePlainText              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeQueue                  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "queue"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeSecretText             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeService                ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "service"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeVectorize              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeAI, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeAssets, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeD1, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeJson, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypePlainText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeQueue, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeSecretText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeService, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeVectorize, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The secret value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretText) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeAI                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "ai"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeAssets                 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "assets"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeD1                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "d1"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeJson                   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "json"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypePlainText              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeQueue                  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "queue"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeSecretText             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeService                ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "service"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeVectorize              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeAI, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeAssets, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeD1, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeJson, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypePlainText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeQueue, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeSecretText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeService, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeVectorize, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindService) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeAI                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "ai"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeAssets                 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "assets"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeD1                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "d1"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeJson                   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "json"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypePlainText              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeQueue                  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "queue"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeSecretText             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeService                ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "service"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeVectorize              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeAI, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeAssets, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeD1, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeJson, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypePlainText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeQueue, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeSecretText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeService, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeVectorize, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumer) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeAI                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "ai"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeAssets                 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "assets"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeD1                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "d1"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeJson                   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "json"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypePlainText              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeQueue                  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "queue"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeSecretText             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeService                ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "service"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeVectorize              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeAI, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeAssets, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeD1, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeJson, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypePlainText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeQueue, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeSecretText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeService, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeVectorize, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorize) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeAI                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "ai"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeAssets                 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "assets"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeD1                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "d1"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeJson                   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "json"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypePlainText              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeQueue                  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "queue"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeSecretText             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeService                ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "service"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeVectorize              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeAI, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeAssets, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeD1, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeJson, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypePlainText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeQueue, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeSecretText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeService, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeVectorize, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadata) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeAI                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "ai"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeAssets                 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "assets"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeD1                     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "d1"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeJson                   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "json"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypePlainText              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeQueue                  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "queue"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeSecretText             ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeService                ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "service"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVectorize              ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeAI, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeAssets, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeD1, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeJson, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypePlainText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeQueue, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeSecretText, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeService, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVectorize, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsType string

const (
	ScriptVersionNewParamsMetadataBindingsTypeAI                     ScriptVersionNewParamsMetadataBindingsType = "ai"
	ScriptVersionNewParamsMetadataBindingsTypeAnalyticsEngine        ScriptVersionNewParamsMetadataBindingsType = "analytics_engine"
	ScriptVersionNewParamsMetadataBindingsTypeAssets                 ScriptVersionNewParamsMetadataBindingsType = "assets"
	ScriptVersionNewParamsMetadataBindingsTypeBrowserRendering       ScriptVersionNewParamsMetadataBindingsType = "browser_rendering"
	ScriptVersionNewParamsMetadataBindingsTypeD1                     ScriptVersionNewParamsMetadataBindingsType = "d1"
	ScriptVersionNewParamsMetadataBindingsTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsTypeJson                   ScriptVersionNewParamsMetadataBindingsType = "json"
	ScriptVersionNewParamsMetadataBindingsTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsTypePlainText              ScriptVersionNewParamsMetadataBindingsType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsTypeQueue                  ScriptVersionNewParamsMetadataBindingsType = "queue"
	ScriptVersionNewParamsMetadataBindingsTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsTypeSecretText             ScriptVersionNewParamsMetadataBindingsType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsTypeService                ScriptVersionNewParamsMetadataBindingsType = "service"
	ScriptVersionNewParamsMetadataBindingsTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsTypeVectorize              ScriptVersionNewParamsMetadataBindingsType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsTypeAI, ScriptVersionNewParamsMetadataBindingsTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsTypeAssets, ScriptVersionNewParamsMetadataBindingsTypeBrowserRendering, ScriptVersionNewParamsMetadataBindingsTypeD1, ScriptVersionNewParamsMetadataBindingsTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsTypeJson, ScriptVersionNewParamsMetadataBindingsTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsTypePlainText, ScriptVersionNewParamsMetadataBindingsTypeQueue, ScriptVersionNewParamsMetadataBindingsTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsTypeSecretText, ScriptVersionNewParamsMetadataBindingsTypeService, ScriptVersionNewParamsMetadataBindingsTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsTypeVectorize, ScriptVersionNewParamsMetadataBindingsTypeVersionMetadata:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type ScriptVersionNewParamsMetadataUsageModel string

const (
	ScriptVersionNewParamsMetadataUsageModelStandard ScriptVersionNewParamsMetadataUsageModel = "standard"
)

func (r ScriptVersionNewParamsMetadataUsageModel) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataUsageModelStandard:
		return true
	}
	return false
}

type ScriptVersionNewResponseEnvelope struct {
	Errors   []ScriptVersionNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptVersionNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ScriptVersionNewResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptVersionNewResponse                `json:"result"`
	JSON    scriptVersionNewResponseEnvelopeJSON    `json:"-"`
}

// scriptVersionNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptVersionNewResponseEnvelope]
type scriptVersionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ScriptVersionNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptVersionNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptVersionNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptVersionNewResponseEnvelopeErrors]
type scriptVersionNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptVersionNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    scriptVersionNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptVersionNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ScriptVersionNewResponseEnvelopeErrorsSource]
type scriptVersionNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ScriptVersionNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptVersionNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptVersionNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScriptVersionNewResponseEnvelopeMessages]
type scriptVersionNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptVersionNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    scriptVersionNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptVersionNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ScriptVersionNewResponseEnvelopeMessagesSource]
type scriptVersionNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptVersionNewResponseEnvelopeSuccess bool

const (
	ScriptVersionNewResponseEnvelopeSuccessTrue ScriptVersionNewResponseEnvelopeSuccess = true
)

func (r ScriptVersionNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptVersionListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Only return versions that can be used in a deployment. Ignores pagination.
	Deployable param.Field[bool] `query:"deployable"`
	// Current page.
	Page param.Field[int64] `query:"page"`
	// Items per-page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [ScriptVersionListParams]'s query parameters as
// `url.Values`.
func (r ScriptVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ScriptVersionGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptVersionGetResponseEnvelope struct {
	Errors   []ScriptVersionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptVersionGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ScriptVersionGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptVersionGetResponse                `json:"result"`
	JSON    scriptVersionGetResponseEnvelopeJSON    `json:"-"`
}

// scriptVersionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptVersionGetResponseEnvelope]
type scriptVersionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ScriptVersionGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptVersionGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptVersionGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptVersionGetResponseEnvelopeErrors]
type scriptVersionGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptVersionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    scriptVersionGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptVersionGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ScriptVersionGetResponseEnvelopeErrorsSource]
type scriptVersionGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ScriptVersionGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptVersionGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptVersionGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScriptVersionGetResponseEnvelopeMessages]
type scriptVersionGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptVersionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    scriptVersionGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptVersionGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ScriptVersionGetResponseEnvelopeMessagesSource]
type scriptVersionGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptVersionGetResponseEnvelopeSuccess bool

const (
	ScriptVersionGetResponseEnvelopeSuccessTrue ScriptVersionGetResponseEnvelopeSuccess = true
)

func (r ScriptVersionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
