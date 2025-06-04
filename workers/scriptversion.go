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
	"reflect"

	"github.com/cloudflare/cloudflare-go/v4/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/tidwall/gjson"
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
	Resources     ScriptVersionNewResponseResources `json:"resources,required"`
	ID            string                            `json:"id"`
	Metadata      ScriptVersionNewResponseMetadata  `json:"metadata"`
	Number        float64                           `json:"number"`
	StartupTimeMs int64                             `json:"startup_time_ms"`
	JSON          scriptVersionNewResponseJSON      `json:"-"`
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

type ScriptVersionNewResponseResources struct {
	Bindings      ScriptVersionNewResponseResourcesBindings      `json:"bindings"`
	Script        ScriptVersionNewResponseResourcesScript        `json:"script"`
	ScriptRuntime ScriptVersionNewResponseResourcesScriptRuntime `json:"script_runtime"`
	JSON          scriptVersionNewResponseResourcesJSON          `json:"-"`
}

// scriptVersionNewResponseResourcesJSON contains the JSON metadata for the struct
// [ScriptVersionNewResponseResources]
type scriptVersionNewResponseResourcesJSON struct {
	Bindings      apijson.Field
	Script        apijson.Field
	ScriptRuntime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResources) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseResourcesBindings struct {
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Result []ScriptVersionNewResponseResourcesBindingsResult `json:"result"`
	JSON   scriptVersionNewResponseResourcesBindingsJSON     `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsJSON contains the JSON metadata for the
// struct [ScriptVersionNewResponseResourcesBindings]
type scriptVersionNewResponseResourcesBindingsJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources.
type ScriptVersionNewResponseResourcesBindingsResult struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultType `json:"type,required"`
	// Identifier of the D1 database to bind to.
	ID string `json:"id"`
	// This field can have the runtime type of [interface{}].
	Algorithm interface{} `json:"algorithm"`
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name"`
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The name of the dataset to bind to.
	Dataset string `json:"dataset"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptVersionNewResponseResourcesBindingsResultFormat `json:"format"`
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name"`
	// JSON data to use.
	Json string `json:"json"`
	// This field can have the runtime type of [interface{}].
	KeyJwk interface{} `json:"key_jwk"`
	// Namespace to bind to.
	Namespace string `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// This field can have the runtime type of
	// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutbound].
	Outbound interface{} `json:"outbound"`
	// Name of the Pipeline to bind to.
	Pipeline string `json:"pipeline"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string `json:"script_name"`
	// Name of the secret in the store.
	SecretName string `json:"secret_name"`
	// Name of Worker to bind to.
	Service string `json:"service"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id"`
	// The text value to use.
	Text string `json:"text"`
	// This field can have the runtime type of
	// [[]ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage].
	Usages interface{} `json:"usages"`
	// Name of the Workflow to bind to.
	WorkflowName string                                              `json:"workflow_name"`
	JSON         scriptVersionNewResponseResourcesBindingsResultJSON `json:"-"`
	union        ScriptVersionNewResponseResourcesBindingsResultUnion
}

// scriptVersionNewResponseResourcesBindingsResultJSON contains the JSON metadata
// for the struct [ScriptVersionNewResponseResourcesBindingsResult]
type scriptVersionNewResponseResourcesBindingsResultJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	ID            apijson.Field
	Algorithm     apijson.Field
	BucketName    apijson.Field
	CertificateID apijson.Field
	ClassName     apijson.Field
	Dataset       apijson.Field
	Environment   apijson.Field
	Format        apijson.Field
	IndexName     apijson.Field
	Json          apijson.Field
	KeyJwk        apijson.Field
	Namespace     apijson.Field
	NamespaceID   apijson.Field
	Outbound      apijson.Field
	Pipeline      apijson.Field
	QueueName     apijson.Field
	ScriptName    apijson.Field
	SecretName    apijson.Field
	Service       apijson.Field
	StoreID       apijson.Field
	Text          apijson.Field
	Usages        apijson.Field
	WorkflowName  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r scriptVersionNewResponseResourcesBindingsResultJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptVersionNewResponseResourcesBindingsResult) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptVersionNewResponseResourcesBindingsResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptVersionNewResponseResourcesBindingsResultUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAI],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngine],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAssets],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindBrowser],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindD1],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespace],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespace],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindHyperdrive],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindJson],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindKVNamespace],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindMTLSCertificate],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPlainText],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPipelines],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindQueue],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindR2Bucket],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretText],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindService],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindTailConsumer],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVectorize],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVersionMetadata],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecret],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKey],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindWorkflow].
func (r ScriptVersionNewResponseResourcesBindingsResult) AsUnion() ScriptVersionNewResponseResourcesBindingsResultUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources.
//
// Union satisfied by
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAI],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngine],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAssets],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindBrowser],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindD1],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespace],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespace],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindHyperdrive],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindJson],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindKVNamespace],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindMTLSCertificate],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPlainText],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPipelines],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindQueue],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindR2Bucket],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretText],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindService],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindTailConsumer],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVectorize],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVersionMetadata],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecret],
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKey] or
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindWorkflow].
type ScriptVersionNewResponseResourcesBindingsResultUnion interface {
	implementsScriptVersionNewResponseResourcesBindingsResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptVersionNewResponseResourcesBindingsResultUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAI{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAssets{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindBrowser{}),
			DiscriminatorValue: "browser",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindD1{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPlainText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPipelines{}),
			DiscriminatorValue: "pipelines",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindService{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVectorize{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecret{}),
			DiscriminatorValue: "secrets_store_secret",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindWorkflow{}),
			DiscriminatorValue: "workflow",
		},
	)
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAIType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAIJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAIJSON contains
// the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAI]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAIJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAIJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAI) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAIType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAITypeAI ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAIType = "ai"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngine]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineJSON struct {
	Dataset     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngine) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAssetsType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAssetsJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAssetsJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAssets]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAssetsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAssets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAssetsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAssets) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAssetsType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAssetsTypeAssets ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAssetsType = "assets"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindBrowserType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindBrowserJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindBrowserJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindBrowser]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindBrowserJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindBrowserJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindBrowser) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindBrowserType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindBrowserTypeBrowser ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindBrowserType = "browser"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindD1Type `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindD1JSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindD1JSON contains
// the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindD1]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindD1JSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindD1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindD1JSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindD1) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindD1Type string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindD1TypeD1 ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindD1Type = "d1"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to.
	Namespace string `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceType `json:"type,required"`
	// Outbound worker.
	Outbound ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutbound `json:"outbound"`
	JSON     scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceJSON     `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespace]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespace) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceTypeDispatchNamespace ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params []string `json:"params"`
	// Outbound worker.
	Worker ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundWorker `json:"worker"`
	JSON   scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundJSON   `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutbound]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment string `json:"environment"`
	// Name of the outbound worker.
	Service string                                                                                               `json:"service"`
	JSON    scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundWorkerJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundWorkerJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundWorker]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceType `json:"type,required"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string                                                                                      `json:"script_name"`
	JSON       scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespace]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	ClassName   apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespace) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindHyperdriveType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindHyperdriveJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindHyperdriveJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindHyperdrive]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindHyperdriveJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindHyperdriveJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindHyperdrive) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindHyperdriveType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindHyperdriveTypeHyperdrive ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindJson struct {
	// JSON data to use.
	Json string `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindJsonType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindJsonJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindJsonJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindJson]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindJsonJSON struct {
	Json        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindJsonJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindJson) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindJsonType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindJsonTypeJson ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindJsonType = "json"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindKVNamespaceType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindKVNamespaceJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindKVNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindKVNamespace]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindKVNamespaceJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindKVNamespace) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindKVNamespaceType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindKVNamespaceTypeKVNamespace ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindMTLSCertificate]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateJSON struct {
	CertificateID apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindMTLSCertificate) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateTypeMTLSCertificate ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The text value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPlainTextType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPlainTextJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPlainTextJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPlainText]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPlainTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPlainText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPlainTextJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPlainText) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPlainTextType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPlainTextTypePlainText ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPlainTextType = "plain_text"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline string `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPipelinesType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPipelinesJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPipelinesJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPipelines]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPipelinesJSON struct {
	Name        apijson.Field
	Pipeline    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPipelines) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPipelinesJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPipelines) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPipelinesType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPipelinesTypePipelines ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPipelinesType = "pipelines"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindQueueType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindQueueJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindQueueJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindQueue]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindQueueJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindQueueJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindQueue) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindQueueType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindQueueTypeQueue ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindQueueType = "queue"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindR2BucketType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindR2BucketJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindR2BucketJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindR2Bucket]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindR2BucketJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindR2Bucket) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindR2BucketType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindR2BucketTypeR2Bucket ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretTextType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretTextJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretTextJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretText]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretText) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretTextType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretTextTypeSecretText ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretTextType = "secret_text"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindServiceType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindServiceJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindServiceJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindService]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindServiceJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindServiceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindService) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindServiceType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindServiceTypeService ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindServiceType = "service"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindTailConsumerType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindTailConsumerJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindTailConsumerJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindTailConsumer]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindTailConsumerJSON struct {
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindTailConsumerJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindTailConsumer) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindTailConsumerType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindTailConsumerTypeTailConsumer ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindTailConsumerType = "tail_consumer"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindTailConsumerTypeTailConsumer:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVectorizeType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVectorizeJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVectorizeJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVectorize]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVectorizeJSON struct {
	IndexName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVectorize) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVectorizeType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVectorizeTypeVectorize ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVectorizeType = "vectorize"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVersionMetadataType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVersionMetadataJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVersionMetadataJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVersionMetadata]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVersionMetadataJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVersionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVersionMetadataJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVersionMetadata) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVersionMetadataType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVersionMetadataTypeVersionMetadata ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the secret in the store.
	SecretName string `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecret]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretJSON struct {
	Name        apijson.Field
	SecretName  apijson.Field
	StoreID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecret) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormat `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyType `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage `json:"usages,required"`
	JSON   scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKey]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKey) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormat string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormatRaw   ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormat = "raw"
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormatPkcs8 ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormat = "pkcs8"
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormatSpki  ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormat = "spki"
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormatJwk   ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormatRaw, ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormatPkcs8, ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormatSpki, ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyTypeSecretKey ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyType = "secret_key"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageEncrypt    ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage = "encrypt"
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageDecrypt    ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage = "decrypt"
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageSign       ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage = "sign"
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageVerify     ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage = "verify"
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageDeriveKey  ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage = "deriveKey"
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageDeriveBits ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage = "deriveBits"
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageWrapKey    ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage = "wrapKey"
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageUnwrapKey  ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageEncrypt, ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageDecrypt, ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageSign, ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageVerify, ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageDeriveKey, ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageDeriveBits, ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageWrapKey, ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindWorkflowType `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName string `json:"workflow_name,required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName string `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName string                                                                        `json:"script_name"`
	JSON       scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindWorkflowJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindWorkflowJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindWorkflow]
type scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindWorkflowJSON struct {
	Name         apijson.Field
	Type         apijson.Field
	WorkflowName apijson.Field
	ClassName    apijson.Field
	ScriptName   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsResultWorkersBindingKindWorkflowJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindWorkflow) implementsScriptVersionNewResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindWorkflowType string

const (
	ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindWorkflowTypeWorkflow ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindWorkflowType = "workflow"
)

func (r ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsResultType string

const (
	ScriptVersionNewResponseResourcesBindingsResultTypeAI                     ScriptVersionNewResponseResourcesBindingsResultType = "ai"
	ScriptVersionNewResponseResourcesBindingsResultTypeAnalyticsEngine        ScriptVersionNewResponseResourcesBindingsResultType = "analytics_engine"
	ScriptVersionNewResponseResourcesBindingsResultTypeAssets                 ScriptVersionNewResponseResourcesBindingsResultType = "assets"
	ScriptVersionNewResponseResourcesBindingsResultTypeBrowser                ScriptVersionNewResponseResourcesBindingsResultType = "browser"
	ScriptVersionNewResponseResourcesBindingsResultTypeD1                     ScriptVersionNewResponseResourcesBindingsResultType = "d1"
	ScriptVersionNewResponseResourcesBindingsResultTypeDispatchNamespace      ScriptVersionNewResponseResourcesBindingsResultType = "dispatch_namespace"
	ScriptVersionNewResponseResourcesBindingsResultTypeDurableObjectNamespace ScriptVersionNewResponseResourcesBindingsResultType = "durable_object_namespace"
	ScriptVersionNewResponseResourcesBindingsResultTypeHyperdrive             ScriptVersionNewResponseResourcesBindingsResultType = "hyperdrive"
	ScriptVersionNewResponseResourcesBindingsResultTypeJson                   ScriptVersionNewResponseResourcesBindingsResultType = "json"
	ScriptVersionNewResponseResourcesBindingsResultTypeKVNamespace            ScriptVersionNewResponseResourcesBindingsResultType = "kv_namespace"
	ScriptVersionNewResponseResourcesBindingsResultTypeMTLSCertificate        ScriptVersionNewResponseResourcesBindingsResultType = "mtls_certificate"
	ScriptVersionNewResponseResourcesBindingsResultTypePlainText              ScriptVersionNewResponseResourcesBindingsResultType = "plain_text"
	ScriptVersionNewResponseResourcesBindingsResultTypePipelines              ScriptVersionNewResponseResourcesBindingsResultType = "pipelines"
	ScriptVersionNewResponseResourcesBindingsResultTypeQueue                  ScriptVersionNewResponseResourcesBindingsResultType = "queue"
	ScriptVersionNewResponseResourcesBindingsResultTypeR2Bucket               ScriptVersionNewResponseResourcesBindingsResultType = "r2_bucket"
	ScriptVersionNewResponseResourcesBindingsResultTypeSecretText             ScriptVersionNewResponseResourcesBindingsResultType = "secret_text"
	ScriptVersionNewResponseResourcesBindingsResultTypeService                ScriptVersionNewResponseResourcesBindingsResultType = "service"
	ScriptVersionNewResponseResourcesBindingsResultTypeTailConsumer           ScriptVersionNewResponseResourcesBindingsResultType = "tail_consumer"
	ScriptVersionNewResponseResourcesBindingsResultTypeVectorize              ScriptVersionNewResponseResourcesBindingsResultType = "vectorize"
	ScriptVersionNewResponseResourcesBindingsResultTypeVersionMetadata        ScriptVersionNewResponseResourcesBindingsResultType = "version_metadata"
	ScriptVersionNewResponseResourcesBindingsResultTypeSecretsStoreSecret     ScriptVersionNewResponseResourcesBindingsResultType = "secrets_store_secret"
	ScriptVersionNewResponseResourcesBindingsResultTypeSecretKey              ScriptVersionNewResponseResourcesBindingsResultType = "secret_key"
	ScriptVersionNewResponseResourcesBindingsResultTypeWorkflow               ScriptVersionNewResponseResourcesBindingsResultType = "workflow"
)

func (r ScriptVersionNewResponseResourcesBindingsResultType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultTypeAI, ScriptVersionNewResponseResourcesBindingsResultTypeAnalyticsEngine, ScriptVersionNewResponseResourcesBindingsResultTypeAssets, ScriptVersionNewResponseResourcesBindingsResultTypeBrowser, ScriptVersionNewResponseResourcesBindingsResultTypeD1, ScriptVersionNewResponseResourcesBindingsResultTypeDispatchNamespace, ScriptVersionNewResponseResourcesBindingsResultTypeDurableObjectNamespace, ScriptVersionNewResponseResourcesBindingsResultTypeHyperdrive, ScriptVersionNewResponseResourcesBindingsResultTypeJson, ScriptVersionNewResponseResourcesBindingsResultTypeKVNamespace, ScriptVersionNewResponseResourcesBindingsResultTypeMTLSCertificate, ScriptVersionNewResponseResourcesBindingsResultTypePlainText, ScriptVersionNewResponseResourcesBindingsResultTypePipelines, ScriptVersionNewResponseResourcesBindingsResultTypeQueue, ScriptVersionNewResponseResourcesBindingsResultTypeR2Bucket, ScriptVersionNewResponseResourcesBindingsResultTypeSecretText, ScriptVersionNewResponseResourcesBindingsResultTypeService, ScriptVersionNewResponseResourcesBindingsResultTypeTailConsumer, ScriptVersionNewResponseResourcesBindingsResultTypeVectorize, ScriptVersionNewResponseResourcesBindingsResultTypeVersionMetadata, ScriptVersionNewResponseResourcesBindingsResultTypeSecretsStoreSecret, ScriptVersionNewResponseResourcesBindingsResultTypeSecretKey, ScriptVersionNewResponseResourcesBindingsResultTypeWorkflow:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptVersionNewResponseResourcesBindingsResultFormat string

const (
	ScriptVersionNewResponseResourcesBindingsResultFormatRaw   ScriptVersionNewResponseResourcesBindingsResultFormat = "raw"
	ScriptVersionNewResponseResourcesBindingsResultFormatPkcs8 ScriptVersionNewResponseResourcesBindingsResultFormat = "pkcs8"
	ScriptVersionNewResponseResourcesBindingsResultFormatSpki  ScriptVersionNewResponseResourcesBindingsResultFormat = "spki"
	ScriptVersionNewResponseResourcesBindingsResultFormatJwk   ScriptVersionNewResponseResourcesBindingsResultFormat = "jwk"
)

func (r ScriptVersionNewResponseResourcesBindingsResultFormat) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsResultFormatRaw, ScriptVersionNewResponseResourcesBindingsResultFormatPkcs8, ScriptVersionNewResponseResourcesBindingsResultFormatSpki, ScriptVersionNewResponseResourcesBindingsResultFormatJwk:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesScript struct {
	Etag             string                                                `json:"etag"`
	Handlers         []string                                              `json:"handlers"`
	LastDeployedFrom string                                                `json:"last_deployed_from"`
	NamedHandlers    []ScriptVersionNewResponseResourcesScriptNamedHandler `json:"named_handlers"`
	JSON             scriptVersionNewResponseResourcesScriptJSON           `json:"-"`
}

// scriptVersionNewResponseResourcesScriptJSON contains the JSON metadata for the
// struct [ScriptVersionNewResponseResourcesScript]
type scriptVersionNewResponseResourcesScriptJSON struct {
	Etag             apijson.Field
	Handlers         apijson.Field
	LastDeployedFrom apijson.Field
	NamedHandlers    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesScript) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesScriptJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseResourcesScriptNamedHandler struct {
	Handlers []string                                                `json:"handlers"`
	Name     string                                                  `json:"name"`
	JSON     scriptVersionNewResponseResourcesScriptNamedHandlerJSON `json:"-"`
}

// scriptVersionNewResponseResourcesScriptNamedHandlerJSON contains the JSON
// metadata for the struct [ScriptVersionNewResponseResourcesScriptNamedHandler]
type scriptVersionNewResponseResourcesScriptNamedHandlerJSON struct {
	Handlers    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesScriptNamedHandler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesScriptNamedHandlerJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseResourcesScriptRuntime struct {
	CompatibilityDate  string                                                   `json:"compatibility_date"`
	CompatibilityFlags []string                                                 `json:"compatibility_flags"`
	Limits             ScriptVersionNewResponseResourcesScriptRuntimeLimits     `json:"limits"`
	MigrationTag       string                                                   `json:"migration_tag"`
	UsageModel         ScriptVersionNewResponseResourcesScriptRuntimeUsageModel `json:"usage_model"`
	JSON               scriptVersionNewResponseResourcesScriptRuntimeJSON       `json:"-"`
}

// scriptVersionNewResponseResourcesScriptRuntimeJSON contains the JSON metadata
// for the struct [ScriptVersionNewResponseResourcesScriptRuntime]
type scriptVersionNewResponseResourcesScriptRuntimeJSON struct {
	CompatibilityDate  apijson.Field
	CompatibilityFlags apijson.Field
	Limits             apijson.Field
	MigrationTag       apijson.Field
	UsageModel         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesScriptRuntime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesScriptRuntimeJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseResourcesScriptRuntimeLimits struct {
	CPUMs int64                                                    `json:"cpu_ms"`
	JSON  scriptVersionNewResponseResourcesScriptRuntimeLimitsJSON `json:"-"`
}

// scriptVersionNewResponseResourcesScriptRuntimeLimitsJSON contains the JSON
// metadata for the struct [ScriptVersionNewResponseResourcesScriptRuntimeLimits]
type scriptVersionNewResponseResourcesScriptRuntimeLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesScriptRuntimeLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesScriptRuntimeLimitsJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseResourcesScriptRuntimeUsageModel string

const (
	ScriptVersionNewResponseResourcesScriptRuntimeUsageModelBundled  ScriptVersionNewResponseResourcesScriptRuntimeUsageModel = "bundled"
	ScriptVersionNewResponseResourcesScriptRuntimeUsageModelUnbound  ScriptVersionNewResponseResourcesScriptRuntimeUsageModel = "unbound"
	ScriptVersionNewResponseResourcesScriptRuntimeUsageModelStandard ScriptVersionNewResponseResourcesScriptRuntimeUsageModel = "standard"
)

func (r ScriptVersionNewResponseResourcesScriptRuntimeUsageModel) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesScriptRuntimeUsageModelBundled, ScriptVersionNewResponseResourcesScriptRuntimeUsageModelUnbound, ScriptVersionNewResponseResourcesScriptRuntimeUsageModelStandard:
		return true
	}
	return false
}

type ScriptVersionNewResponseMetadata struct {
	AuthorEmail string                                 `json:"author_email"`
	AuthorID    string                                 `json:"author_id"`
	CreatedOn   string                                 `json:"created_on"`
	HasPreview  bool                                   `json:"hasPreview"`
	ModifiedOn  string                                 `json:"modified_on"`
	Source      ScriptVersionNewResponseMetadataSource `json:"source"`
	JSON        scriptVersionNewResponseMetadataJSON   `json:"-"`
}

// scriptVersionNewResponseMetadataJSON contains the JSON metadata for the struct
// [ScriptVersionNewResponseMetadata]
type scriptVersionNewResponseMetadataJSON struct {
	AuthorEmail apijson.Field
	AuthorID    apijson.Field
	CreatedOn   apijson.Field
	HasPreview  apijson.Field
	ModifiedOn  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseMetadataSource string

const (
	ScriptVersionNewResponseMetadataSourceUnknown      ScriptVersionNewResponseMetadataSource = "unknown"
	ScriptVersionNewResponseMetadataSourceAPI          ScriptVersionNewResponseMetadataSource = "api"
	ScriptVersionNewResponseMetadataSourceWrangler     ScriptVersionNewResponseMetadataSource = "wrangler"
	ScriptVersionNewResponseMetadataSourceTerraform    ScriptVersionNewResponseMetadataSource = "terraform"
	ScriptVersionNewResponseMetadataSourceDash         ScriptVersionNewResponseMetadataSource = "dash"
	ScriptVersionNewResponseMetadataSourceDashTemplate ScriptVersionNewResponseMetadataSource = "dash_template"
	ScriptVersionNewResponseMetadataSourceIntegration  ScriptVersionNewResponseMetadataSource = "integration"
	ScriptVersionNewResponseMetadataSourceQuickEditor  ScriptVersionNewResponseMetadataSource = "quick_editor"
	ScriptVersionNewResponseMetadataSourcePlayground   ScriptVersionNewResponseMetadataSource = "playground"
	ScriptVersionNewResponseMetadataSourceWorkersci    ScriptVersionNewResponseMetadataSource = "workersci"
)

func (r ScriptVersionNewResponseMetadataSource) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseMetadataSourceUnknown, ScriptVersionNewResponseMetadataSourceAPI, ScriptVersionNewResponseMetadataSourceWrangler, ScriptVersionNewResponseMetadataSourceTerraform, ScriptVersionNewResponseMetadataSourceDash, ScriptVersionNewResponseMetadataSourceDashTemplate, ScriptVersionNewResponseMetadataSourceIntegration, ScriptVersionNewResponseMetadataSourceQuickEditor, ScriptVersionNewResponseMetadataSourcePlayground, ScriptVersionNewResponseMetadataSourceWorkersci:
		return true
	}
	return false
}

type ScriptVersionListResponse struct {
	ID       string                            `json:"id"`
	Metadata ScriptVersionListResponseMetadata `json:"metadata"`
	Number   float64                           `json:"number"`
	JSON     scriptVersionListResponseJSON     `json:"-"`
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

type ScriptVersionListResponseMetadata struct {
	AuthorEmail string                                  `json:"author_email"`
	AuthorID    string                                  `json:"author_id"`
	CreatedOn   string                                  `json:"created_on"`
	HasPreview  bool                                    `json:"hasPreview"`
	ModifiedOn  string                                  `json:"modified_on"`
	Source      ScriptVersionListResponseMetadataSource `json:"source"`
	JSON        scriptVersionListResponseMetadataJSON   `json:"-"`
}

// scriptVersionListResponseMetadataJSON contains the JSON metadata for the struct
// [ScriptVersionListResponseMetadata]
type scriptVersionListResponseMetadataJSON struct {
	AuthorEmail apijson.Field
	AuthorID    apijson.Field
	CreatedOn   apijson.Field
	HasPreview  apijson.Field
	ModifiedOn  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionListResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionListResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionListResponseMetadataSource string

const (
	ScriptVersionListResponseMetadataSourceUnknown      ScriptVersionListResponseMetadataSource = "unknown"
	ScriptVersionListResponseMetadataSourceAPI          ScriptVersionListResponseMetadataSource = "api"
	ScriptVersionListResponseMetadataSourceWrangler     ScriptVersionListResponseMetadataSource = "wrangler"
	ScriptVersionListResponseMetadataSourceTerraform    ScriptVersionListResponseMetadataSource = "terraform"
	ScriptVersionListResponseMetadataSourceDash         ScriptVersionListResponseMetadataSource = "dash"
	ScriptVersionListResponseMetadataSourceDashTemplate ScriptVersionListResponseMetadataSource = "dash_template"
	ScriptVersionListResponseMetadataSourceIntegration  ScriptVersionListResponseMetadataSource = "integration"
	ScriptVersionListResponseMetadataSourceQuickEditor  ScriptVersionListResponseMetadataSource = "quick_editor"
	ScriptVersionListResponseMetadataSourcePlayground   ScriptVersionListResponseMetadataSource = "playground"
	ScriptVersionListResponseMetadataSourceWorkersci    ScriptVersionListResponseMetadataSource = "workersci"
)

func (r ScriptVersionListResponseMetadataSource) IsKnown() bool {
	switch r {
	case ScriptVersionListResponseMetadataSourceUnknown, ScriptVersionListResponseMetadataSourceAPI, ScriptVersionListResponseMetadataSourceWrangler, ScriptVersionListResponseMetadataSourceTerraform, ScriptVersionListResponseMetadataSourceDash, ScriptVersionListResponseMetadataSourceDashTemplate, ScriptVersionListResponseMetadataSourceIntegration, ScriptVersionListResponseMetadataSourceQuickEditor, ScriptVersionListResponseMetadataSourcePlayground, ScriptVersionListResponseMetadataSourceWorkersci:
		return true
	}
	return false
}

type ScriptVersionGetResponse struct {
	Resources ScriptVersionGetResponseResources `json:"resources,required"`
	ID        string                            `json:"id"`
	Metadata  ScriptVersionGetResponseMetadata  `json:"metadata"`
	Number    float64                           `json:"number"`
	JSON      scriptVersionGetResponseJSON      `json:"-"`
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

type ScriptVersionGetResponseResources struct {
	Bindings      ScriptVersionGetResponseResourcesBindings      `json:"bindings"`
	Script        ScriptVersionGetResponseResourcesScript        `json:"script"`
	ScriptRuntime ScriptVersionGetResponseResourcesScriptRuntime `json:"script_runtime"`
	JSON          scriptVersionGetResponseResourcesJSON          `json:"-"`
}

// scriptVersionGetResponseResourcesJSON contains the JSON metadata for the struct
// [ScriptVersionGetResponseResources]
type scriptVersionGetResponseResourcesJSON struct {
	Bindings      apijson.Field
	Script        apijson.Field
	ScriptRuntime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResources) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseResourcesBindings struct {
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Result []ScriptVersionGetResponseResourcesBindingsResult `json:"result"`
	JSON   scriptVersionGetResponseResourcesBindingsJSON     `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsJSON contains the JSON metadata for the
// struct [ScriptVersionGetResponseResourcesBindings]
type scriptVersionGetResponseResourcesBindingsJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources.
type ScriptVersionGetResponseResourcesBindingsResult struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultType `json:"type,required"`
	// Identifier of the D1 database to bind to.
	ID string `json:"id"`
	// This field can have the runtime type of [interface{}].
	Algorithm interface{} `json:"algorithm"`
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name"`
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The name of the dataset to bind to.
	Dataset string `json:"dataset"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptVersionGetResponseResourcesBindingsResultFormat `json:"format"`
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name"`
	// JSON data to use.
	Json string `json:"json"`
	// This field can have the runtime type of [interface{}].
	KeyJwk interface{} `json:"key_jwk"`
	// Namespace to bind to.
	Namespace string `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// This field can have the runtime type of
	// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutbound].
	Outbound interface{} `json:"outbound"`
	// Name of the Pipeline to bind to.
	Pipeline string `json:"pipeline"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string `json:"script_name"`
	// Name of the secret in the store.
	SecretName string `json:"secret_name"`
	// Name of Worker to bind to.
	Service string `json:"service"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id"`
	// The text value to use.
	Text string `json:"text"`
	// This field can have the runtime type of
	// [[]ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage].
	Usages interface{} `json:"usages"`
	// Name of the Workflow to bind to.
	WorkflowName string                                              `json:"workflow_name"`
	JSON         scriptVersionGetResponseResourcesBindingsResultJSON `json:"-"`
	union        ScriptVersionGetResponseResourcesBindingsResultUnion
}

// scriptVersionGetResponseResourcesBindingsResultJSON contains the JSON metadata
// for the struct [ScriptVersionGetResponseResourcesBindingsResult]
type scriptVersionGetResponseResourcesBindingsResultJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	ID            apijson.Field
	Algorithm     apijson.Field
	BucketName    apijson.Field
	CertificateID apijson.Field
	ClassName     apijson.Field
	Dataset       apijson.Field
	Environment   apijson.Field
	Format        apijson.Field
	IndexName     apijson.Field
	Json          apijson.Field
	KeyJwk        apijson.Field
	Namespace     apijson.Field
	NamespaceID   apijson.Field
	Outbound      apijson.Field
	Pipeline      apijson.Field
	QueueName     apijson.Field
	ScriptName    apijson.Field
	SecretName    apijson.Field
	Service       apijson.Field
	StoreID       apijson.Field
	Text          apijson.Field
	Usages        apijson.Field
	WorkflowName  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r scriptVersionGetResponseResourcesBindingsResultJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptVersionGetResponseResourcesBindingsResult) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptVersionGetResponseResourcesBindingsResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptVersionGetResponseResourcesBindingsResultUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAI],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngine],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAssets],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindBrowser],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindD1],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespace],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespace],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindHyperdrive],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindJson],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindKVNamespace],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindMTLSCertificate],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPlainText],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPipelines],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindQueue],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindR2Bucket],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretText],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindService],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindTailConsumer],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVectorize],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVersionMetadata],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecret],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKey],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindWorkflow].
func (r ScriptVersionGetResponseResourcesBindingsResult) AsUnion() ScriptVersionGetResponseResourcesBindingsResultUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources.
//
// Union satisfied by
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAI],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngine],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAssets],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindBrowser],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindD1],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespace],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespace],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindHyperdrive],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindJson],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindKVNamespace],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindMTLSCertificate],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPlainText],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPipelines],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindQueue],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindR2Bucket],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretText],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindService],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindTailConsumer],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVectorize],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVersionMetadata],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecret],
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKey] or
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindWorkflow].
type ScriptVersionGetResponseResourcesBindingsResultUnion interface {
	implementsScriptVersionGetResponseResourcesBindingsResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptVersionGetResponseResourcesBindingsResultUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAI{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAssets{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindBrowser{}),
			DiscriminatorValue: "browser",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindD1{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPlainText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPipelines{}),
			DiscriminatorValue: "pipelines",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindService{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVectorize{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecret{}),
			DiscriminatorValue: "secrets_store_secret",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindWorkflow{}),
			DiscriminatorValue: "workflow",
		},
	)
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAIType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAIJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAIJSON contains
// the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAI]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAIJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAIJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAI) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAIType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAITypeAI ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAIType = "ai"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngine]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineJSON struct {
	Dataset     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngine) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAssetsType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAssetsJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAssetsJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAssets]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAssetsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAssets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAssetsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAssets) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAssetsType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAssetsTypeAssets ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAssetsType = "assets"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindBrowserType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindBrowserJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindBrowserJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindBrowser]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindBrowserJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindBrowserJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindBrowser) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindBrowserType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindBrowserTypeBrowser ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindBrowserType = "browser"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindD1Type `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindD1JSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindD1JSON contains
// the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindD1]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindD1JSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindD1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindD1JSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindD1) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindD1Type string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindD1TypeD1 ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindD1Type = "d1"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to.
	Namespace string `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceType `json:"type,required"`
	// Outbound worker.
	Outbound ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutbound `json:"outbound"`
	JSON     scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceJSON     `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespace]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespace) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceTypeDispatchNamespace ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params []string `json:"params"`
	// Outbound worker.
	Worker ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundWorker `json:"worker"`
	JSON   scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundJSON   `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutbound]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment string `json:"environment"`
	// Name of the outbound worker.
	Service string                                                                                               `json:"service"`
	JSON    scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundWorkerJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundWorkerJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundWorker]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDispatchNamespaceOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceType `json:"type,required"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string                                                                                      `json:"script_name"`
	JSON       scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespace]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	ClassName   apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespace) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindHyperdriveType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindHyperdriveJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindHyperdriveJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindHyperdrive]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindHyperdriveJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindHyperdriveJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindHyperdrive) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindHyperdriveType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindHyperdriveTypeHyperdrive ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindJson struct {
	// JSON data to use.
	Json string `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindJsonType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindJsonJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindJsonJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindJson]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindJsonJSON struct {
	Json        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindJsonJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindJson) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindJsonType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindJsonTypeJson ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindJsonType = "json"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindKVNamespaceType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindKVNamespaceJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindKVNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindKVNamespace]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindKVNamespaceJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindKVNamespace) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindKVNamespaceType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindKVNamespaceTypeKVNamespace ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindMTLSCertificate]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateJSON struct {
	CertificateID apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindMTLSCertificate) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateTypeMTLSCertificate ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The text value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPlainTextType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPlainTextJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPlainTextJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPlainText]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPlainTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPlainText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPlainTextJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPlainText) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPlainTextType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPlainTextTypePlainText ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPlainTextType = "plain_text"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline string `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPipelinesType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPipelinesJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPipelinesJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPipelines]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPipelinesJSON struct {
	Name        apijson.Field
	Pipeline    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPipelines) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPipelinesJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPipelines) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPipelinesType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPipelinesTypePipelines ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPipelinesType = "pipelines"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindQueueType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindQueueJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindQueueJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindQueue]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindQueueJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindQueueJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindQueue) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindQueueType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindQueueTypeQueue ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindQueueType = "queue"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindR2BucketType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindR2BucketJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindR2BucketJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindR2Bucket]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindR2BucketJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindR2Bucket) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindR2BucketType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindR2BucketTypeR2Bucket ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretTextType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretTextJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretTextJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretText]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretText) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretTextType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretTextTypeSecretText ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretTextType = "secret_text"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindServiceType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindServiceJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindServiceJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindService]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindServiceJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindServiceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindService) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindServiceType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindServiceTypeService ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindServiceType = "service"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindTailConsumerType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindTailConsumerJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindTailConsumerJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindTailConsumer]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindTailConsumerJSON struct {
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindTailConsumerJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindTailConsumer) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindTailConsumerType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindTailConsumerTypeTailConsumer ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindTailConsumerType = "tail_consumer"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindTailConsumerTypeTailConsumer:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVectorizeType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVectorizeJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVectorizeJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVectorize]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVectorizeJSON struct {
	IndexName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVectorize) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVectorizeType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVectorizeTypeVectorize ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVectorizeType = "vectorize"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVersionMetadataType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVersionMetadataJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVersionMetadataJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVersionMetadata]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVersionMetadataJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVersionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVersionMetadataJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVersionMetadata) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVersionMetadataType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVersionMetadataTypeVersionMetadata ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the secret in the store.
	SecretName string `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecret]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretJSON struct {
	Name        apijson.Field
	SecretName  apijson.Field
	StoreID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecret) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormat `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyType `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage `json:"usages,required"`
	JSON   scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKey]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKey) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormat string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormatRaw   ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormat = "raw"
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormatPkcs8 ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormat = "pkcs8"
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormatSpki  ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormat = "spki"
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormatJwk   ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormatRaw, ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormatPkcs8, ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormatSpki, ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyTypeSecretKey ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyType = "secret_key"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageEncrypt    ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage = "encrypt"
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageDecrypt    ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage = "decrypt"
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageSign       ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage = "sign"
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageVerify     ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage = "verify"
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageDeriveKey  ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage = "deriveKey"
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageDeriveBits ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage = "deriveBits"
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageWrapKey    ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage = "wrapKey"
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageUnwrapKey  ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageEncrypt, ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageDecrypt, ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageSign, ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageVerify, ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageDeriveKey, ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageDeriveBits, ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageWrapKey, ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindWorkflowType `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName string `json:"workflow_name,required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName string `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName string                                                                        `json:"script_name"`
	JSON       scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindWorkflowJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindWorkflowJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindWorkflow]
type scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindWorkflowJSON struct {
	Name         apijson.Field
	Type         apijson.Field
	WorkflowName apijson.Field
	ClassName    apijson.Field
	ScriptName   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsResultWorkersBindingKindWorkflowJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindWorkflow) implementsScriptVersionGetResponseResourcesBindingsResult() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindWorkflowType string

const (
	ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindWorkflowTypeWorkflow ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindWorkflowType = "workflow"
)

func (r ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsResultType string

const (
	ScriptVersionGetResponseResourcesBindingsResultTypeAI                     ScriptVersionGetResponseResourcesBindingsResultType = "ai"
	ScriptVersionGetResponseResourcesBindingsResultTypeAnalyticsEngine        ScriptVersionGetResponseResourcesBindingsResultType = "analytics_engine"
	ScriptVersionGetResponseResourcesBindingsResultTypeAssets                 ScriptVersionGetResponseResourcesBindingsResultType = "assets"
	ScriptVersionGetResponseResourcesBindingsResultTypeBrowser                ScriptVersionGetResponseResourcesBindingsResultType = "browser"
	ScriptVersionGetResponseResourcesBindingsResultTypeD1                     ScriptVersionGetResponseResourcesBindingsResultType = "d1"
	ScriptVersionGetResponseResourcesBindingsResultTypeDispatchNamespace      ScriptVersionGetResponseResourcesBindingsResultType = "dispatch_namespace"
	ScriptVersionGetResponseResourcesBindingsResultTypeDurableObjectNamespace ScriptVersionGetResponseResourcesBindingsResultType = "durable_object_namespace"
	ScriptVersionGetResponseResourcesBindingsResultTypeHyperdrive             ScriptVersionGetResponseResourcesBindingsResultType = "hyperdrive"
	ScriptVersionGetResponseResourcesBindingsResultTypeJson                   ScriptVersionGetResponseResourcesBindingsResultType = "json"
	ScriptVersionGetResponseResourcesBindingsResultTypeKVNamespace            ScriptVersionGetResponseResourcesBindingsResultType = "kv_namespace"
	ScriptVersionGetResponseResourcesBindingsResultTypeMTLSCertificate        ScriptVersionGetResponseResourcesBindingsResultType = "mtls_certificate"
	ScriptVersionGetResponseResourcesBindingsResultTypePlainText              ScriptVersionGetResponseResourcesBindingsResultType = "plain_text"
	ScriptVersionGetResponseResourcesBindingsResultTypePipelines              ScriptVersionGetResponseResourcesBindingsResultType = "pipelines"
	ScriptVersionGetResponseResourcesBindingsResultTypeQueue                  ScriptVersionGetResponseResourcesBindingsResultType = "queue"
	ScriptVersionGetResponseResourcesBindingsResultTypeR2Bucket               ScriptVersionGetResponseResourcesBindingsResultType = "r2_bucket"
	ScriptVersionGetResponseResourcesBindingsResultTypeSecretText             ScriptVersionGetResponseResourcesBindingsResultType = "secret_text"
	ScriptVersionGetResponseResourcesBindingsResultTypeService                ScriptVersionGetResponseResourcesBindingsResultType = "service"
	ScriptVersionGetResponseResourcesBindingsResultTypeTailConsumer           ScriptVersionGetResponseResourcesBindingsResultType = "tail_consumer"
	ScriptVersionGetResponseResourcesBindingsResultTypeVectorize              ScriptVersionGetResponseResourcesBindingsResultType = "vectorize"
	ScriptVersionGetResponseResourcesBindingsResultTypeVersionMetadata        ScriptVersionGetResponseResourcesBindingsResultType = "version_metadata"
	ScriptVersionGetResponseResourcesBindingsResultTypeSecretsStoreSecret     ScriptVersionGetResponseResourcesBindingsResultType = "secrets_store_secret"
	ScriptVersionGetResponseResourcesBindingsResultTypeSecretKey              ScriptVersionGetResponseResourcesBindingsResultType = "secret_key"
	ScriptVersionGetResponseResourcesBindingsResultTypeWorkflow               ScriptVersionGetResponseResourcesBindingsResultType = "workflow"
)

func (r ScriptVersionGetResponseResourcesBindingsResultType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultTypeAI, ScriptVersionGetResponseResourcesBindingsResultTypeAnalyticsEngine, ScriptVersionGetResponseResourcesBindingsResultTypeAssets, ScriptVersionGetResponseResourcesBindingsResultTypeBrowser, ScriptVersionGetResponseResourcesBindingsResultTypeD1, ScriptVersionGetResponseResourcesBindingsResultTypeDispatchNamespace, ScriptVersionGetResponseResourcesBindingsResultTypeDurableObjectNamespace, ScriptVersionGetResponseResourcesBindingsResultTypeHyperdrive, ScriptVersionGetResponseResourcesBindingsResultTypeJson, ScriptVersionGetResponseResourcesBindingsResultTypeKVNamespace, ScriptVersionGetResponseResourcesBindingsResultTypeMTLSCertificate, ScriptVersionGetResponseResourcesBindingsResultTypePlainText, ScriptVersionGetResponseResourcesBindingsResultTypePipelines, ScriptVersionGetResponseResourcesBindingsResultTypeQueue, ScriptVersionGetResponseResourcesBindingsResultTypeR2Bucket, ScriptVersionGetResponseResourcesBindingsResultTypeSecretText, ScriptVersionGetResponseResourcesBindingsResultTypeService, ScriptVersionGetResponseResourcesBindingsResultTypeTailConsumer, ScriptVersionGetResponseResourcesBindingsResultTypeVectorize, ScriptVersionGetResponseResourcesBindingsResultTypeVersionMetadata, ScriptVersionGetResponseResourcesBindingsResultTypeSecretsStoreSecret, ScriptVersionGetResponseResourcesBindingsResultTypeSecretKey, ScriptVersionGetResponseResourcesBindingsResultTypeWorkflow:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptVersionGetResponseResourcesBindingsResultFormat string

const (
	ScriptVersionGetResponseResourcesBindingsResultFormatRaw   ScriptVersionGetResponseResourcesBindingsResultFormat = "raw"
	ScriptVersionGetResponseResourcesBindingsResultFormatPkcs8 ScriptVersionGetResponseResourcesBindingsResultFormat = "pkcs8"
	ScriptVersionGetResponseResourcesBindingsResultFormatSpki  ScriptVersionGetResponseResourcesBindingsResultFormat = "spki"
	ScriptVersionGetResponseResourcesBindingsResultFormatJwk   ScriptVersionGetResponseResourcesBindingsResultFormat = "jwk"
)

func (r ScriptVersionGetResponseResourcesBindingsResultFormat) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsResultFormatRaw, ScriptVersionGetResponseResourcesBindingsResultFormatPkcs8, ScriptVersionGetResponseResourcesBindingsResultFormatSpki, ScriptVersionGetResponseResourcesBindingsResultFormatJwk:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesScript struct {
	Etag             string                                                `json:"etag"`
	Handlers         []string                                              `json:"handlers"`
	LastDeployedFrom string                                                `json:"last_deployed_from"`
	NamedHandlers    []ScriptVersionGetResponseResourcesScriptNamedHandler `json:"named_handlers"`
	JSON             scriptVersionGetResponseResourcesScriptJSON           `json:"-"`
}

// scriptVersionGetResponseResourcesScriptJSON contains the JSON metadata for the
// struct [ScriptVersionGetResponseResourcesScript]
type scriptVersionGetResponseResourcesScriptJSON struct {
	Etag             apijson.Field
	Handlers         apijson.Field
	LastDeployedFrom apijson.Field
	NamedHandlers    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesScript) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesScriptJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseResourcesScriptNamedHandler struct {
	Handlers []string                                                `json:"handlers"`
	Name     string                                                  `json:"name"`
	JSON     scriptVersionGetResponseResourcesScriptNamedHandlerJSON `json:"-"`
}

// scriptVersionGetResponseResourcesScriptNamedHandlerJSON contains the JSON
// metadata for the struct [ScriptVersionGetResponseResourcesScriptNamedHandler]
type scriptVersionGetResponseResourcesScriptNamedHandlerJSON struct {
	Handlers    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesScriptNamedHandler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesScriptNamedHandlerJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseResourcesScriptRuntime struct {
	CompatibilityDate  string                                                   `json:"compatibility_date"`
	CompatibilityFlags []string                                                 `json:"compatibility_flags"`
	Limits             ScriptVersionGetResponseResourcesScriptRuntimeLimits     `json:"limits"`
	MigrationTag       string                                                   `json:"migration_tag"`
	UsageModel         ScriptVersionGetResponseResourcesScriptRuntimeUsageModel `json:"usage_model"`
	JSON               scriptVersionGetResponseResourcesScriptRuntimeJSON       `json:"-"`
}

// scriptVersionGetResponseResourcesScriptRuntimeJSON contains the JSON metadata
// for the struct [ScriptVersionGetResponseResourcesScriptRuntime]
type scriptVersionGetResponseResourcesScriptRuntimeJSON struct {
	CompatibilityDate  apijson.Field
	CompatibilityFlags apijson.Field
	Limits             apijson.Field
	MigrationTag       apijson.Field
	UsageModel         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesScriptRuntime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesScriptRuntimeJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseResourcesScriptRuntimeLimits struct {
	CPUMs int64                                                    `json:"cpu_ms"`
	JSON  scriptVersionGetResponseResourcesScriptRuntimeLimitsJSON `json:"-"`
}

// scriptVersionGetResponseResourcesScriptRuntimeLimitsJSON contains the JSON
// metadata for the struct [ScriptVersionGetResponseResourcesScriptRuntimeLimits]
type scriptVersionGetResponseResourcesScriptRuntimeLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesScriptRuntimeLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesScriptRuntimeLimitsJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseResourcesScriptRuntimeUsageModel string

const (
	ScriptVersionGetResponseResourcesScriptRuntimeUsageModelBundled  ScriptVersionGetResponseResourcesScriptRuntimeUsageModel = "bundled"
	ScriptVersionGetResponseResourcesScriptRuntimeUsageModelUnbound  ScriptVersionGetResponseResourcesScriptRuntimeUsageModel = "unbound"
	ScriptVersionGetResponseResourcesScriptRuntimeUsageModelStandard ScriptVersionGetResponseResourcesScriptRuntimeUsageModel = "standard"
)

func (r ScriptVersionGetResponseResourcesScriptRuntimeUsageModel) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesScriptRuntimeUsageModelBundled, ScriptVersionGetResponseResourcesScriptRuntimeUsageModelUnbound, ScriptVersionGetResponseResourcesScriptRuntimeUsageModelStandard:
		return true
	}
	return false
}

type ScriptVersionGetResponseMetadata struct {
	AuthorEmail string                                 `json:"author_email"`
	AuthorID    string                                 `json:"author_id"`
	CreatedOn   string                                 `json:"created_on"`
	HasPreview  bool                                   `json:"hasPreview"`
	ModifiedOn  string                                 `json:"modified_on"`
	Source      ScriptVersionGetResponseMetadataSource `json:"source"`
	JSON        scriptVersionGetResponseMetadataJSON   `json:"-"`
}

// scriptVersionGetResponseMetadataJSON contains the JSON metadata for the struct
// [ScriptVersionGetResponseMetadata]
type scriptVersionGetResponseMetadataJSON struct {
	AuthorEmail apijson.Field
	AuthorID    apijson.Field
	CreatedOn   apijson.Field
	HasPreview  apijson.Field
	ModifiedOn  apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseMetadataSource string

const (
	ScriptVersionGetResponseMetadataSourceUnknown      ScriptVersionGetResponseMetadataSource = "unknown"
	ScriptVersionGetResponseMetadataSourceAPI          ScriptVersionGetResponseMetadataSource = "api"
	ScriptVersionGetResponseMetadataSourceWrangler     ScriptVersionGetResponseMetadataSource = "wrangler"
	ScriptVersionGetResponseMetadataSourceTerraform    ScriptVersionGetResponseMetadataSource = "terraform"
	ScriptVersionGetResponseMetadataSourceDash         ScriptVersionGetResponseMetadataSource = "dash"
	ScriptVersionGetResponseMetadataSourceDashTemplate ScriptVersionGetResponseMetadataSource = "dash_template"
	ScriptVersionGetResponseMetadataSourceIntegration  ScriptVersionGetResponseMetadataSource = "integration"
	ScriptVersionGetResponseMetadataSourceQuickEditor  ScriptVersionGetResponseMetadataSource = "quick_editor"
	ScriptVersionGetResponseMetadataSourcePlayground   ScriptVersionGetResponseMetadataSource = "playground"
	ScriptVersionGetResponseMetadataSourceWorkersci    ScriptVersionGetResponseMetadataSource = "workersci"
)

func (r ScriptVersionGetResponseMetadataSource) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseMetadataSourceUnknown, ScriptVersionGetResponseMetadataSourceAPI, ScriptVersionGetResponseMetadataSourceWrangler, ScriptVersionGetResponseMetadataSourceTerraform, ScriptVersionGetResponseMetadataSourceDash, ScriptVersionGetResponseMetadataSourceDashTemplate, ScriptVersionGetResponseMetadataSourceIntegration, ScriptVersionGetResponseMetadataSourceQuickEditor, ScriptVersionGetResponseMetadataSourcePlayground, ScriptVersionGetResponseMetadataSourceWorkersci:
		return true
	}
	return false
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

// A binding to allow the Worker to communicate with resources.
type ScriptVersionNewParamsMetadataBinding struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsType] `json:"type,required"`
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
	Format param.Field[ScriptVersionNewParamsMetadataBindingsFormat] `json:"format"`
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

func (r ScriptVersionNewParamsMetadataBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBinding) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// A binding to allow the Worker to communicate with resources.
//
// Satisfied by
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAI],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngine],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssets],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowser],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespace],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdrive],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJson],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespace],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificate],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainText],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelines],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueue],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2Bucket],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretText],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindService],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumer],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorize],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadata],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKey],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflow],
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssets) implementsScriptVersionNewParamsMetadataBindingUnion() {
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

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowser) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserTypeBrowser ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserType = "browser"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserTypeBrowser:
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespace) implementsScriptVersionNewParamsMetadataBindingUnion() {
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

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType] `json:"type,required"`
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
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
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
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
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeJson ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType = "json"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonTypeJson:
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
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
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
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
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
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline param.Field[string] `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelinesType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelines) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelines) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelinesType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelinesTypePipelines ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelinesType = "pipelines"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelinesTypePipelines:
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
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeQueue ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType = "queue"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueTypeQueue:
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
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeR2Bucket ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketTypeR2Bucket:
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
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeSecretText ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType = "secret_text"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextTypeSecretText:
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTailConsumer) implementsScriptVersionNewParamsMetadataBindingUnion() {
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorize) implementsScriptVersionNewParamsMetadataBindingUnion() {
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

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadata) implementsScriptVersionNewParamsMetadataBindingUnion() {
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

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the secret in the store.
	SecretName param.Field[string] `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID param.Field[string] `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm param.Field[interface{}] `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormat] `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyType] `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages param.Field[[]ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage] `json:"usages,required"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string] `json:"key_base64"`
	// Key data in
	// [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key)
	// format. Required if `format` is "jwk".
	KeyJwk param.Field[interface{}] `json:"key_jwk"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKey) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormat string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormatRaw   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "raw"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormatPkcs8 ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "pkcs8"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormatSpki  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "spki"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormatJwk   ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormatRaw, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormatPkcs8, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormatSpki, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyTypeSecretKey ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyType = "secret_key"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageEncrypt    ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "encrypt"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDecrypt    ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "decrypt"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageSign       ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "sign"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageVerify     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "verify"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveKey  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "deriveKey"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveBits ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "deriveBits"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageWrapKey    ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "wrapKey"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageUnwrapKey  ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageEncrypt, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDecrypt, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageSign, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageVerify, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveKey, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveBits, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageWrapKey, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflowType] `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName param.Field[string] `json:"workflow_name,required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName param.Field[string] `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflow) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflow) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflowType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflowTypeWorkflow ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflowType = "workflow"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflowTypeWorkflow:
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
	ScriptVersionNewParamsMetadataBindingsTypeBrowser                ScriptVersionNewParamsMetadataBindingsType = "browser"
	ScriptVersionNewParamsMetadataBindingsTypeD1                     ScriptVersionNewParamsMetadataBindingsType = "d1"
	ScriptVersionNewParamsMetadataBindingsTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsTypeJson                   ScriptVersionNewParamsMetadataBindingsType = "json"
	ScriptVersionNewParamsMetadataBindingsTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsTypePlainText              ScriptVersionNewParamsMetadataBindingsType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsTypePipelines              ScriptVersionNewParamsMetadataBindingsType = "pipelines"
	ScriptVersionNewParamsMetadataBindingsTypeQueue                  ScriptVersionNewParamsMetadataBindingsType = "queue"
	ScriptVersionNewParamsMetadataBindingsTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsTypeSecretText             ScriptVersionNewParamsMetadataBindingsType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsTypeService                ScriptVersionNewParamsMetadataBindingsType = "service"
	ScriptVersionNewParamsMetadataBindingsTypeTailConsumer           ScriptVersionNewParamsMetadataBindingsType = "tail_consumer"
	ScriptVersionNewParamsMetadataBindingsTypeVectorize              ScriptVersionNewParamsMetadataBindingsType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsType = "version_metadata"
	ScriptVersionNewParamsMetadataBindingsTypeSecretsStoreSecret     ScriptVersionNewParamsMetadataBindingsType = "secrets_store_secret"
	ScriptVersionNewParamsMetadataBindingsTypeSecretKey              ScriptVersionNewParamsMetadataBindingsType = "secret_key"
	ScriptVersionNewParamsMetadataBindingsTypeWorkflow               ScriptVersionNewParamsMetadataBindingsType = "workflow"
)

func (r ScriptVersionNewParamsMetadataBindingsType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsTypeAI, ScriptVersionNewParamsMetadataBindingsTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsTypeAssets, ScriptVersionNewParamsMetadataBindingsTypeBrowser, ScriptVersionNewParamsMetadataBindingsTypeD1, ScriptVersionNewParamsMetadataBindingsTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsTypeJson, ScriptVersionNewParamsMetadataBindingsTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsTypePlainText, ScriptVersionNewParamsMetadataBindingsTypePipelines, ScriptVersionNewParamsMetadataBindingsTypeQueue, ScriptVersionNewParamsMetadataBindingsTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsTypeSecretText, ScriptVersionNewParamsMetadataBindingsTypeService, ScriptVersionNewParamsMetadataBindingsTypeTailConsumer, ScriptVersionNewParamsMetadataBindingsTypeVectorize, ScriptVersionNewParamsMetadataBindingsTypeVersionMetadata, ScriptVersionNewParamsMetadataBindingsTypeSecretsStoreSecret, ScriptVersionNewParamsMetadataBindingsTypeSecretKey, ScriptVersionNewParamsMetadataBindingsTypeWorkflow:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptVersionNewParamsMetadataBindingsFormat string

const (
	ScriptVersionNewParamsMetadataBindingsFormatRaw   ScriptVersionNewParamsMetadataBindingsFormat = "raw"
	ScriptVersionNewParamsMetadataBindingsFormatPkcs8 ScriptVersionNewParamsMetadataBindingsFormat = "pkcs8"
	ScriptVersionNewParamsMetadataBindingsFormatSpki  ScriptVersionNewParamsMetadataBindingsFormat = "spki"
	ScriptVersionNewParamsMetadataBindingsFormatJwk   ScriptVersionNewParamsMetadataBindingsFormat = "jwk"
)

func (r ScriptVersionNewParamsMetadataBindingsFormat) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsFormatRaw, ScriptVersionNewParamsMetadataBindingsFormatPkcs8, ScriptVersionNewParamsMetadataBindingsFormatSpki, ScriptVersionNewParamsMetadataBindingsFormatJwk:
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
	Result   ScriptVersionNewResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptVersionNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptVersionNewResponseEnvelopeJSON    `json:"-"`
}

// scriptVersionNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptVersionNewResponseEnvelope]
type scriptVersionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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
	Result   ScriptVersionGetResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ScriptVersionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptVersionGetResponseEnvelopeJSON    `json:"-"`
}

// scriptVersionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptVersionGetResponseEnvelope]
type scriptVersionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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
