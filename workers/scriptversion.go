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

	"github.com/cloudflare/cloudflare-go/v3/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v3/shared"
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
	// Identifier
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
	JSON param.Field[string] `json:"json"`
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

func (r ScriptVersionNewParamsMetadataBinding) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAny],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAI],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngine],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssets],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRendering],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespace],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDo],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdrive],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJSON],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespace],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCERT],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainText],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueue],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecret],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindService],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumer],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorize],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadata],
// [ScriptVersionNewParamsMetadataBinding].
type ScriptVersionNewParamsMetadataBindingUnion interface {
	implementsWorkersScriptVersionNewParamsMetadataBindingUnion()
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAny struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type        param.Field[string]    `json:"type,required"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAny) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAny) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAI) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeAI ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType = "ai"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngine struct {
	// The dataset name to bind to.
	Dataset param.Field[string] `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngine) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngine) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssets) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeAssets ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType = "assets"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsTypeAssets:
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRendering) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType = "browser_rendering"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering:
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeD1 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type = "d1"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1TypeD1:
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespace) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
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

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDo struct {
	// The exported class name of the Durable Object.
	ClassName param.Field[string] `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDoType] `json:"type,required"`
	// The environment of the script_name to bind to.
	Environment param.Field[string] `json:"environment"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDo) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDoType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDoTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDoType = "durable_object_namespace"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDoType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDoTypeDurableObjectNamespace:
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdrive) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJSON struct {
	// JSON data to use.
	JSON param.Field[string] `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJSONType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJSON) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJSON) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJSONType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJSONTypeJSON ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJSONType = "json"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJSONType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJSONTypeJSON:
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespace) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCERT struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCERTType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCERT) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCERT) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCERTType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCERTTypeMTLSCertificate ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCERTType = "mtls_certificate"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCERTType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCERTTypeMTLSCertificate:
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainText) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText:
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueue) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeQueue ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "queue"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2 struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2Type] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2Type string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2TypeR2Bucket ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2Type = "r2_bucket"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2Type) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2TypeR2Bucket:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecret struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The secret value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecret) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTypeSecretText ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretType = "secret_text"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTypeSecretText:
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindService) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeService ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType = "service"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceTypeService:
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumer) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeTailConsumer ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumerTypeTailConsumer:
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorize) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeVectorize ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType = "vectorize"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeTypeVectorize:
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadata) implementsWorkersScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
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

// Whether the API call was successful
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
	// Identifier
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
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptVersionGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
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

// Whether the API call was successful
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
