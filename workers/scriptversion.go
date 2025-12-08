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
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	Resources ScriptVersionNewResponseResources `json:"resources,required"`
	// Unique identifier for the version.
	ID       string                           `json:"id"`
	Metadata ScriptVersionNewResponseMetadata `json:"metadata"`
	// Sequential version number.
	Number float64 `json:"number"`
	// Time in milliseconds spent on
	// [Worker startup](https://developers.cloudflare.com/workers/platform/limits/#worker-startup-time).
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

type ScriptVersionNewResponseResources struct {
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings []ScriptVersionNewResponseResourcesBinding `json:"bindings"`
	Script   ScriptVersionNewResponseResourcesScript    `json:"script"`
	// Runtime configuration for the Worker.
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

// A binding to allow the Worker to communicate with resources.
type ScriptVersionNewResponseResourcesBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsType `json:"type,required"`
	// Identifier of the D1 database to bind to.
	ID string `json:"id"`
	// This field can have the runtime type of [interface{}].
	Algorithm interface{} `json:"algorithm"`
	// This field can have the runtime type of [[]string].
	AllowedDestinationAddresses interface{} `json:"allowed_destination_addresses"`
	// This field can have the runtime type of [[]string].
	AllowedSenderAddresses interface{} `json:"allowed_sender_addresses"`
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name"`
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The name of the dataset to bind to.
	Dataset string `json:"dataset"`
	// Destination address for the email.
	DestinationAddress string `json:"destination_address" format:"email"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptVersionNewResponseResourcesBindingsFormat `json:"format"`
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name"`
	// JSON data to use.
	Json string `json:"json"`
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction ScriptVersionNewResponseResourcesBindingsJurisdiction `json:"jurisdiction"`
	// This field can have the runtime type of [interface{}].
	KeyJwk interface{} `json:"key_jwk"`
	// The name of the dispatch namespace.
	Namespace string `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The old name of the inherited binding. If set, the binding will be renamed from
	// `old_name` to `name` in the new version. If not set, the binding will keep the
	// same name between versions.
	OldName string `json:"old_name"`
	// This field can have the runtime type of
	// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound].
	Outbound interface{} `json:"outbound"`
	// The name of the file containing the data content. Only accepted for
	// `service worker syntax` Workers.
	Part string `json:"part"`
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
	// [[]ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage].
	Usages interface{} `json:"usages"`
	// Identifier for the version to inherit the binding from, which can be the version
	// ID or the literal "latest" to inherit from the latest version. Defaults to
	// inheriting the binding from the latest version.
	VersionID string `json:"version_id"`
	// Name of the Workflow to bind to.
	WorkflowName string                                       `json:"workflow_name"`
	JSON         scriptVersionNewResponseResourcesBindingJSON `json:"-"`
	union        ScriptVersionNewResponseResourcesBindingsUnion
}

// scriptVersionNewResponseResourcesBindingJSON contains the JSON metadata for the
// struct [ScriptVersionNewResponseResourcesBinding]
type scriptVersionNewResponseResourcesBindingJSON struct {
	Name                        apijson.Field
	Type                        apijson.Field
	ID                          apijson.Field
	Algorithm                   apijson.Field
	AllowedDestinationAddresses apijson.Field
	AllowedSenderAddresses      apijson.Field
	BucketName                  apijson.Field
	CertificateID               apijson.Field
	ClassName                   apijson.Field
	Dataset                     apijson.Field
	DestinationAddress          apijson.Field
	Environment                 apijson.Field
	Format                      apijson.Field
	IndexName                   apijson.Field
	Json                        apijson.Field
	Jurisdiction                apijson.Field
	KeyJwk                      apijson.Field
	Namespace                   apijson.Field
	NamespaceID                 apijson.Field
	OldName                     apijson.Field
	Outbound                    apijson.Field
	Part                        apijson.Field
	Pipeline                    apijson.Field
	QueueName                   apijson.Field
	ScriptName                  apijson.Field
	SecretName                  apijson.Field
	Service                     apijson.Field
	StoreID                     apijson.Field
	Text                        apijson.Field
	Usages                      apijson.Field
	VersionID                   apijson.Field
	WorkflowName                apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r scriptVersionNewResponseResourcesBindingJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptVersionNewResponseResourcesBinding) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptVersionNewResponseResourcesBinding{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptVersionNewResponseResourcesBindingsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAI],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngine],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssets],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowser],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDataBlob],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespace],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdrive],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindInherit],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindImages],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJson],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespace],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificate],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainText],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelines],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueue],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2Bucket],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretText],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSendEmail],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindService],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTextBlob],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorize],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadata],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKey],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflow],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWasmModule].
func (r ScriptVersionNewResponseResourcesBinding) AsUnion() ScriptVersionNewResponseResourcesBindingsUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources.
//
// Union satisfied by
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAI],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngine],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssets],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowser],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDataBlob],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespace],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdrive],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindInherit],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindImages],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJson],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespace],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificate],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainText],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelines],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueue],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2Bucket],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretText],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSendEmail],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindService],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTextBlob],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorize],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadata],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKey],
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflow] or
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWasmModule].
type ScriptVersionNewResponseResourcesBindingsUnion interface {
	implementsScriptVersionNewResponseResourcesBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptVersionNewResponseResourcesBindingsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowser{}),
			DiscriminatorValue: "browser",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDataBlob{}),
			DiscriminatorValue: "data_blob",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindInherit{}),
			DiscriminatorValue: "inherit",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindImages{}),
			DiscriminatorValue: "images",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelines{}),
			DiscriminatorValue: "pipelines",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSendEmail{}),
			DiscriminatorValue: "send_email",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTextBlob{}),
			DiscriminatorValue: "text_blob",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret{}),
			DiscriminatorValue: "secrets_store_secret",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflow{}),
			DiscriminatorValue: "workflow",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWasmModule{}),
			DiscriminatorValue: "wasm_module",
		},
	)
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAIType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindAIJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindAIJSON contains the
// JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAI]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindAIJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindAIJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAI) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAIType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAITypeAI ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAIType = "ai"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngine]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineJSON struct {
	Dataset     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngine) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsJSON contains
// the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssets]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssets) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsTypeAssets ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsType = "assets"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserJSON contains
// the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowser]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowser) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserTypeBrowser ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserType = "browser"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1Type `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindD1JSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindD1JSON contains the
// JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindD1JSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindD1JSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1Type string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1TypeD1 ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1Type = "d1"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDataBlob struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The name of the file containing the data content. Only accepted for
	// `service worker syntax` Workers.
	Part string `json:"part,required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDataBlobType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindDataBlobJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindDataBlobJSON contains
// the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDataBlob]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindDataBlobJSON struct {
	Name        apijson.Field
	Part        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDataBlob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindDataBlobJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDataBlob) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDataBlobType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDataBlobTypeDataBlob ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDataBlobType = "data_blob"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDataBlobType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDataBlobTypeDataBlob:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The name of the dispatch namespace.
	Namespace string `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceType `json:"type,required"`
	// Outbound worker.
	Outbound ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound `json:"outbound"`
	JSON     scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceJSON     `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespace]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespace) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params []string `json:"params"`
	// Outbound worker.
	Worker ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorker `json:"worker"`
	JSON   scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundJSON   `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment string `json:"environment"`
	// Name of the outbound worker.
	Service string                                                                                         `json:"service"`
	JSON    scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorker]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceType `json:"type,required"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string                                                                                `json:"script_name"`
	JSON       scriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	ClassName   apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdrive]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdrive) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveTypeHyperdrive ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindInherit struct {
	// The name of the inherited binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindInheritType `json:"type,required"`
	// The old name of the inherited binding. If set, the binding will be renamed from
	// `old_name` to `name` in the new version. If not set, the binding will keep the
	// same name between versions.
	OldName string `json:"old_name"`
	// Identifier for the version to inherit the binding from, which can be the version
	// ID or the literal "latest" to inherit from the latest version. Defaults to
	// inheriting the binding from the latest version.
	VersionID string                                                                 `json:"version_id"`
	JSON      scriptVersionNewResponseResourcesBindingsWorkersBindingKindInheritJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindInheritJSON contains
// the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindInherit]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindInheritJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	OldName     apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindInherit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindInheritJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindInherit) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindInheritType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindInheritTypeInherit ScriptVersionNewResponseResourcesBindingsWorkersBindingKindInheritType = "inherit"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindInheritType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindInheritTypeInherit:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindImages struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindImagesType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindImagesJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindImagesJSON contains
// the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindImages]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindImagesJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindImages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindImagesJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindImages) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindImagesType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindImagesTypeImages ScriptVersionNewResponseResourcesBindingsWorkersBindingKindImagesType = "images"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindImagesType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindImagesTypeImages:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json string `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonJSON contains the
// JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJson]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonJSON struct {
	Json        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJson) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonTypeJson ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonType = "json"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespace]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespace) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceTypeKVNamespace ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificate]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateJSON struct {
	CertificateID apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificate) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The text value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainText]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainText) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextTypePlainText ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline string `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelines]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesJSON struct {
	Name        apijson.Field
	Pipeline    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelines) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelines) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesTypePipelines ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesType = "pipelines"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueJSON contains
// the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueue]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueue) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueTypeQueue ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueType = "queue"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketType `json:"type,required"`
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketJurisdiction `json:"jurisdiction"`
	JSON         scriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketJSON         `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketJSON contains
// the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2Bucket]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketJSON struct {
	BucketName   apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	Jurisdiction apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2Bucket) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketTypeR2Bucket ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

// The
// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
// of the R2 bucket.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketJurisdiction string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketJurisdictionEu      ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketJurisdiction = "eu"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketJurisdictionFedramp ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketJurisdiction = "fedramp"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketJurisdiction) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketJurisdictionEu, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindR2BucketJurisdictionFedramp:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretText]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretText) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextTypeSecretText ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextType = "secret_text"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSendEmail struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSendEmailType `json:"type,required"`
	// List of allowed destination addresses.
	AllowedDestinationAddresses []string `json:"allowed_destination_addresses" format:"email"`
	// List of allowed sender addresses.
	AllowedSenderAddresses []string `json:"allowed_sender_addresses" format:"email"`
	// Destination address for the email.
	DestinationAddress string                                                                   `json:"destination_address" format:"email"`
	JSON               scriptVersionNewResponseResourcesBindingsWorkersBindingKindSendEmailJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindSendEmailJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSendEmail]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindSendEmailJSON struct {
	Name                        apijson.Field
	Type                        apijson.Field
	AllowedDestinationAddresses apijson.Field
	AllowedSenderAddresses      apijson.Field
	DestinationAddress          apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSendEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindSendEmailJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSendEmail) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSendEmailType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSendEmailTypeSendEmail ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSendEmailType = "send_email"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSendEmailType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSendEmailTypeSendEmail:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindService struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceType `json:"type,required"`
	// Optional environment if the Worker utilizes one.
	Environment string                                                                 `json:"environment"`
	JSON        scriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceJSON contains
// the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindService]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceJSON struct {
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindService) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceTypeService ScriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceType = "service"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTextBlob struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The name of the file containing the text content. Only accepted for
	// `service worker syntax` Workers.
	Part string `json:"part,required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTextBlobType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindTextBlobJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindTextBlobJSON contains
// the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTextBlob]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindTextBlobJSON struct {
	Name        apijson.Field
	Part        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTextBlob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindTextBlobJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTextBlob) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTextBlobType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTextBlobTypeTextBlob ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTextBlobType = "text_blob"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTextBlobType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindTextBlobTypeTextBlob:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorize]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeJSON struct {
	IndexName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorize) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeTypeVectorize ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeType = "vectorize"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadata]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadata) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the secret in the store.
	SecretName string `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretJSON struct {
	Name        apijson.Field
	SecretName  apijson.Field
	StoreID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormat `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyType `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage `json:"usages,required"`
	JSON   scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKey]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKey) implementsScriptVersionNewResponseResourcesBinding() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormat string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormatRaw   ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormat = "raw"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormatPkcs8 ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormat = "pkcs8"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormatSpki  ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormat = "spki"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormatJwk   ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormatRaw, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormatPkcs8, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormatSpki, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyTypeSecretKey ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyType = "secret_key"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageEncrypt    ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "encrypt"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDecrypt    ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "decrypt"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageSign       ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "sign"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageVerify     ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "verify"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDeriveKey  ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "deriveKey"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDeriveBits ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "deriveBits"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageWrapKey    ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "wrapKey"
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageUnwrapKey  ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageEncrypt, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDecrypt, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageSign, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageVerify, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDeriveKey, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDeriveBits, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageWrapKey, ScriptVersionNewResponseResourcesBindingsWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowType `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName string `json:"workflow_name,required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName string `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName string                                                                  `json:"script_name"`
	JSON       scriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowJSON contains
// the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflow]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowJSON struct {
	Name         apijson.Field
	Type         apijson.Field
	WorkflowName apijson.Field
	ClassName    apijson.Field
	ScriptName   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflow) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowTypeWorkflow ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowType = "workflow"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWasmModule struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The name of the file containing the WebAssembly module content. Only accepted
	// for `service worker syntax` Workers.
	Part string `json:"part,required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWasmModuleType `json:"type,required"`
	JSON scriptVersionNewResponseResourcesBindingsWorkersBindingKindWasmModuleJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsWorkersBindingKindWasmModuleJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWasmModule]
type scriptVersionNewResponseResourcesBindingsWorkersBindingKindWasmModuleJSON struct {
	Name        apijson.Field
	Part        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWasmModule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsWorkersBindingKindWasmModuleJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWasmModule) implementsScriptVersionNewResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWasmModuleType string

const (
	ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWasmModuleTypeWasmModule ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWasmModuleType = "wasm_module"
)

func (r ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWasmModuleType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsWorkersBindingKindWasmModuleTypeWasmModule:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptVersionNewResponseResourcesBindingsType string

const (
	ScriptVersionNewResponseResourcesBindingsTypeAI                     ScriptVersionNewResponseResourcesBindingsType = "ai"
	ScriptVersionNewResponseResourcesBindingsTypeAnalyticsEngine        ScriptVersionNewResponseResourcesBindingsType = "analytics_engine"
	ScriptVersionNewResponseResourcesBindingsTypeAssets                 ScriptVersionNewResponseResourcesBindingsType = "assets"
	ScriptVersionNewResponseResourcesBindingsTypeBrowser                ScriptVersionNewResponseResourcesBindingsType = "browser"
	ScriptVersionNewResponseResourcesBindingsTypeD1                     ScriptVersionNewResponseResourcesBindingsType = "d1"
	ScriptVersionNewResponseResourcesBindingsTypeDataBlob               ScriptVersionNewResponseResourcesBindingsType = "data_blob"
	ScriptVersionNewResponseResourcesBindingsTypeDispatchNamespace      ScriptVersionNewResponseResourcesBindingsType = "dispatch_namespace"
	ScriptVersionNewResponseResourcesBindingsTypeDurableObjectNamespace ScriptVersionNewResponseResourcesBindingsType = "durable_object_namespace"
	ScriptVersionNewResponseResourcesBindingsTypeHyperdrive             ScriptVersionNewResponseResourcesBindingsType = "hyperdrive"
	ScriptVersionNewResponseResourcesBindingsTypeInherit                ScriptVersionNewResponseResourcesBindingsType = "inherit"
	ScriptVersionNewResponseResourcesBindingsTypeImages                 ScriptVersionNewResponseResourcesBindingsType = "images"
	ScriptVersionNewResponseResourcesBindingsTypeJson                   ScriptVersionNewResponseResourcesBindingsType = "json"
	ScriptVersionNewResponseResourcesBindingsTypeKVNamespace            ScriptVersionNewResponseResourcesBindingsType = "kv_namespace"
	ScriptVersionNewResponseResourcesBindingsTypeMTLSCertificate        ScriptVersionNewResponseResourcesBindingsType = "mtls_certificate"
	ScriptVersionNewResponseResourcesBindingsTypePlainText              ScriptVersionNewResponseResourcesBindingsType = "plain_text"
	ScriptVersionNewResponseResourcesBindingsTypePipelines              ScriptVersionNewResponseResourcesBindingsType = "pipelines"
	ScriptVersionNewResponseResourcesBindingsTypeQueue                  ScriptVersionNewResponseResourcesBindingsType = "queue"
	ScriptVersionNewResponseResourcesBindingsTypeR2Bucket               ScriptVersionNewResponseResourcesBindingsType = "r2_bucket"
	ScriptVersionNewResponseResourcesBindingsTypeSecretText             ScriptVersionNewResponseResourcesBindingsType = "secret_text"
	ScriptVersionNewResponseResourcesBindingsTypeSendEmail              ScriptVersionNewResponseResourcesBindingsType = "send_email"
	ScriptVersionNewResponseResourcesBindingsTypeService                ScriptVersionNewResponseResourcesBindingsType = "service"
	ScriptVersionNewResponseResourcesBindingsTypeTextBlob               ScriptVersionNewResponseResourcesBindingsType = "text_blob"
	ScriptVersionNewResponseResourcesBindingsTypeVectorize              ScriptVersionNewResponseResourcesBindingsType = "vectorize"
	ScriptVersionNewResponseResourcesBindingsTypeVersionMetadata        ScriptVersionNewResponseResourcesBindingsType = "version_metadata"
	ScriptVersionNewResponseResourcesBindingsTypeSecretsStoreSecret     ScriptVersionNewResponseResourcesBindingsType = "secrets_store_secret"
	ScriptVersionNewResponseResourcesBindingsTypeSecretKey              ScriptVersionNewResponseResourcesBindingsType = "secret_key"
	ScriptVersionNewResponseResourcesBindingsTypeWorkflow               ScriptVersionNewResponseResourcesBindingsType = "workflow"
	ScriptVersionNewResponseResourcesBindingsTypeWasmModule             ScriptVersionNewResponseResourcesBindingsType = "wasm_module"
)

func (r ScriptVersionNewResponseResourcesBindingsType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsTypeAI, ScriptVersionNewResponseResourcesBindingsTypeAnalyticsEngine, ScriptVersionNewResponseResourcesBindingsTypeAssets, ScriptVersionNewResponseResourcesBindingsTypeBrowser, ScriptVersionNewResponseResourcesBindingsTypeD1, ScriptVersionNewResponseResourcesBindingsTypeDataBlob, ScriptVersionNewResponseResourcesBindingsTypeDispatchNamespace, ScriptVersionNewResponseResourcesBindingsTypeDurableObjectNamespace, ScriptVersionNewResponseResourcesBindingsTypeHyperdrive, ScriptVersionNewResponseResourcesBindingsTypeInherit, ScriptVersionNewResponseResourcesBindingsTypeImages, ScriptVersionNewResponseResourcesBindingsTypeJson, ScriptVersionNewResponseResourcesBindingsTypeKVNamespace, ScriptVersionNewResponseResourcesBindingsTypeMTLSCertificate, ScriptVersionNewResponseResourcesBindingsTypePlainText, ScriptVersionNewResponseResourcesBindingsTypePipelines, ScriptVersionNewResponseResourcesBindingsTypeQueue, ScriptVersionNewResponseResourcesBindingsTypeR2Bucket, ScriptVersionNewResponseResourcesBindingsTypeSecretText, ScriptVersionNewResponseResourcesBindingsTypeSendEmail, ScriptVersionNewResponseResourcesBindingsTypeService, ScriptVersionNewResponseResourcesBindingsTypeTextBlob, ScriptVersionNewResponseResourcesBindingsTypeVectorize, ScriptVersionNewResponseResourcesBindingsTypeVersionMetadata, ScriptVersionNewResponseResourcesBindingsTypeSecretsStoreSecret, ScriptVersionNewResponseResourcesBindingsTypeSecretKey, ScriptVersionNewResponseResourcesBindingsTypeWorkflow, ScriptVersionNewResponseResourcesBindingsTypeWasmModule:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptVersionNewResponseResourcesBindingsFormat string

const (
	ScriptVersionNewResponseResourcesBindingsFormatRaw   ScriptVersionNewResponseResourcesBindingsFormat = "raw"
	ScriptVersionNewResponseResourcesBindingsFormatPkcs8 ScriptVersionNewResponseResourcesBindingsFormat = "pkcs8"
	ScriptVersionNewResponseResourcesBindingsFormatSpki  ScriptVersionNewResponseResourcesBindingsFormat = "spki"
	ScriptVersionNewResponseResourcesBindingsFormatJwk   ScriptVersionNewResponseResourcesBindingsFormat = "jwk"
)

func (r ScriptVersionNewResponseResourcesBindingsFormat) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsFormatRaw, ScriptVersionNewResponseResourcesBindingsFormatPkcs8, ScriptVersionNewResponseResourcesBindingsFormatSpki, ScriptVersionNewResponseResourcesBindingsFormatJwk:
		return true
	}
	return false
}

// The
// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
// of the R2 bucket.
type ScriptVersionNewResponseResourcesBindingsJurisdiction string

const (
	ScriptVersionNewResponseResourcesBindingsJurisdictionEu      ScriptVersionNewResponseResourcesBindingsJurisdiction = "eu"
	ScriptVersionNewResponseResourcesBindingsJurisdictionFedramp ScriptVersionNewResponseResourcesBindingsJurisdiction = "fedramp"
)

func (r ScriptVersionNewResponseResourcesBindingsJurisdiction) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesBindingsJurisdictionEu, ScriptVersionNewResponseResourcesBindingsJurisdictionFedramp:
		return true
	}
	return false
}

type ScriptVersionNewResponseResourcesScript struct {
	// Hashed script content
	Etag string `json:"etag"`
	// The names of handlers exported as part of the default export.
	Handlers []string `json:"handlers"`
	// The client most recently used to deploy this Worker.
	LastDeployedFrom string `json:"last_deployed_from"`
	// Named exports, such as Durable Object class implementations and named
	// entrypoints.
	NamedHandlers []ScriptVersionNewResponseResourcesScriptNamedHandler `json:"named_handlers"`
	JSON          scriptVersionNewResponseResourcesScriptJSON           `json:"-"`
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
	// The names of handlers exported as part of the named export.
	Handlers []string `json:"handlers"`
	// The name of the exported class or entrypoint.
	Name string                                                  `json:"name"`
	JSON scriptVersionNewResponseResourcesScriptNamedHandlerJSON `json:"-"`
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

// Runtime configuration for the Worker.
type ScriptVersionNewResponseResourcesScriptRuntime struct {
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate string `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Resource limits for the Worker.
	Limits ScriptVersionNewResponseResourcesScriptRuntimeLimits `json:"limits"`
	// The tag of the Durable Object migration that was most recently applied for this
	// Worker.
	MigrationTag string `json:"migration_tag"`
	// Usage model for the Worker invocations.
	UsageModel ScriptVersionNewResponseResourcesScriptRuntimeUsageModel `json:"usage_model"`
	JSON       scriptVersionNewResponseResourcesScriptRuntimeJSON       `json:"-"`
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

// Resource limits for the Worker.
type ScriptVersionNewResponseResourcesScriptRuntimeLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
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

// Usage model for the Worker invocations.
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
	// Email of the user who created the version.
	AuthorEmail string `json:"author_email"`
	// Identifier of the user who created the version.
	AuthorID string `json:"author_id"`
	// When the version was created.
	CreatedOn string `json:"created_on"`
	// Whether the version can be previewed.
	HasPreview bool `json:"hasPreview"`
	// When the version was last modified.
	ModifiedOn string `json:"modified_on"`
	// The source of the version upload.
	Source ScriptVersionNewResponseMetadataSource `json:"source"`
	JSON   scriptVersionNewResponseMetadataJSON   `json:"-"`
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

// The source of the version upload.
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
	// Unique identifier for the version.
	ID       string                            `json:"id"`
	Metadata ScriptVersionListResponseMetadata `json:"metadata"`
	// Sequential version number.
	Number float64                       `json:"number"`
	JSON   scriptVersionListResponseJSON `json:"-"`
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
	// Email of the user who created the version.
	AuthorEmail string `json:"author_email"`
	// Identifier of the user who created the version.
	AuthorID string `json:"author_id"`
	// When the version was created.
	CreatedOn string `json:"created_on"`
	// Whether the version can be previewed.
	HasPreview bool `json:"hasPreview"`
	// When the version was last modified.
	ModifiedOn string `json:"modified_on"`
	// The source of the version upload.
	Source ScriptVersionListResponseMetadataSource `json:"source"`
	JSON   scriptVersionListResponseMetadataJSON   `json:"-"`
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

// The source of the version upload.
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
	// Unique identifier for the version.
	ID       string                           `json:"id"`
	Metadata ScriptVersionGetResponseMetadata `json:"metadata"`
	// Sequential version number.
	Number float64                      `json:"number"`
	JSON   scriptVersionGetResponseJSON `json:"-"`
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
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings []ScriptVersionGetResponseResourcesBinding `json:"bindings"`
	Script   ScriptVersionGetResponseResourcesScript    `json:"script"`
	// Runtime configuration for the Worker.
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

// A binding to allow the Worker to communicate with resources.
type ScriptVersionGetResponseResourcesBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsType `json:"type,required"`
	// Identifier of the D1 database to bind to.
	ID string `json:"id"`
	// This field can have the runtime type of [interface{}].
	Algorithm interface{} `json:"algorithm"`
	// This field can have the runtime type of [[]string].
	AllowedDestinationAddresses interface{} `json:"allowed_destination_addresses"`
	// This field can have the runtime type of [[]string].
	AllowedSenderAddresses interface{} `json:"allowed_sender_addresses"`
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name"`
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The name of the dataset to bind to.
	Dataset string `json:"dataset"`
	// Destination address for the email.
	DestinationAddress string `json:"destination_address" format:"email"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptVersionGetResponseResourcesBindingsFormat `json:"format"`
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name"`
	// JSON data to use.
	Json string `json:"json"`
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction ScriptVersionGetResponseResourcesBindingsJurisdiction `json:"jurisdiction"`
	// This field can have the runtime type of [interface{}].
	KeyJwk interface{} `json:"key_jwk"`
	// The name of the dispatch namespace.
	Namespace string `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The old name of the inherited binding. If set, the binding will be renamed from
	// `old_name` to `name` in the new version. If not set, the binding will keep the
	// same name between versions.
	OldName string `json:"old_name"`
	// This field can have the runtime type of
	// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound].
	Outbound interface{} `json:"outbound"`
	// The name of the file containing the data content. Only accepted for
	// `service worker syntax` Workers.
	Part string `json:"part"`
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
	// [[]ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage].
	Usages interface{} `json:"usages"`
	// Identifier for the version to inherit the binding from, which can be the version
	// ID or the literal "latest" to inherit from the latest version. Defaults to
	// inheriting the binding from the latest version.
	VersionID string `json:"version_id"`
	// Name of the Workflow to bind to.
	WorkflowName string                                       `json:"workflow_name"`
	JSON         scriptVersionGetResponseResourcesBindingJSON `json:"-"`
	union        ScriptVersionGetResponseResourcesBindingsUnion
}

// scriptVersionGetResponseResourcesBindingJSON contains the JSON metadata for the
// struct [ScriptVersionGetResponseResourcesBinding]
type scriptVersionGetResponseResourcesBindingJSON struct {
	Name                        apijson.Field
	Type                        apijson.Field
	ID                          apijson.Field
	Algorithm                   apijson.Field
	AllowedDestinationAddresses apijson.Field
	AllowedSenderAddresses      apijson.Field
	BucketName                  apijson.Field
	CertificateID               apijson.Field
	ClassName                   apijson.Field
	Dataset                     apijson.Field
	DestinationAddress          apijson.Field
	Environment                 apijson.Field
	Format                      apijson.Field
	IndexName                   apijson.Field
	Json                        apijson.Field
	Jurisdiction                apijson.Field
	KeyJwk                      apijson.Field
	Namespace                   apijson.Field
	NamespaceID                 apijson.Field
	OldName                     apijson.Field
	Outbound                    apijson.Field
	Part                        apijson.Field
	Pipeline                    apijson.Field
	QueueName                   apijson.Field
	ScriptName                  apijson.Field
	SecretName                  apijson.Field
	Service                     apijson.Field
	StoreID                     apijson.Field
	Text                        apijson.Field
	Usages                      apijson.Field
	VersionID                   apijson.Field
	WorkflowName                apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r scriptVersionGetResponseResourcesBindingJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptVersionGetResponseResourcesBinding) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptVersionGetResponseResourcesBinding{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptVersionGetResponseResourcesBindingsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAI],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngine],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssets],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowser],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDataBlob],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespace],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdrive],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindInherit],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindImages],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJson],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespace],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificate],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainText],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelines],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueue],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2Bucket],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretText],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSendEmail],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindService],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTextBlob],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorize],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadata],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKey],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflow],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWasmModule].
func (r ScriptVersionGetResponseResourcesBinding) AsUnion() ScriptVersionGetResponseResourcesBindingsUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources.
//
// Union satisfied by
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAI],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngine],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssets],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowser],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDataBlob],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespace],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdrive],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindInherit],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindImages],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJson],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespace],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificate],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainText],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelines],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueue],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2Bucket],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretText],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSendEmail],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindService],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTextBlob],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorize],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadata],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKey],
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflow] or
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWasmModule].
type ScriptVersionGetResponseResourcesBindingsUnion interface {
	implementsScriptVersionGetResponseResourcesBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptVersionGetResponseResourcesBindingsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowser{}),
			DiscriminatorValue: "browser",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDataBlob{}),
			DiscriminatorValue: "data_blob",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindInherit{}),
			DiscriminatorValue: "inherit",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindImages{}),
			DiscriminatorValue: "images",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelines{}),
			DiscriminatorValue: "pipelines",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSendEmail{}),
			DiscriminatorValue: "send_email",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTextBlob{}),
			DiscriminatorValue: "text_blob",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret{}),
			DiscriminatorValue: "secrets_store_secret",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflow{}),
			DiscriminatorValue: "workflow",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWasmModule{}),
			DiscriminatorValue: "wasm_module",
		},
	)
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAIType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindAIJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindAIJSON contains the
// JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAI]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindAIJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindAIJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAI) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAIType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAITypeAI ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAIType = "ai"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngine]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineJSON struct {
	Dataset     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngine) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsJSON contains
// the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssets]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssets) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsTypeAssets ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsType = "assets"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserJSON contains
// the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowser]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowser) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserTypeBrowser ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserType = "browser"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1Type `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindD1JSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindD1JSON contains the
// JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindD1JSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindD1JSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1Type string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1TypeD1 ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1Type = "d1"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDataBlob struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The name of the file containing the data content. Only accepted for
	// `service worker syntax` Workers.
	Part string `json:"part,required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDataBlobType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindDataBlobJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindDataBlobJSON contains
// the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDataBlob]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindDataBlobJSON struct {
	Name        apijson.Field
	Part        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDataBlob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindDataBlobJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDataBlob) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDataBlobType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDataBlobTypeDataBlob ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDataBlobType = "data_blob"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDataBlobType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDataBlobTypeDataBlob:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The name of the dispatch namespace.
	Namespace string `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceType `json:"type,required"`
	// Outbound worker.
	Outbound ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound `json:"outbound"`
	JSON     scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceJSON     `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespace]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespace) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params []string `json:"params"`
	// Outbound worker.
	Worker ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorker `json:"worker"`
	JSON   scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundJSON   `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment string `json:"environment"`
	// Name of the outbound worker.
	Service string                                                                                         `json:"service"`
	JSON    scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorker]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceType `json:"type,required"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string                                                                                `json:"script_name"`
	JSON       scriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	ClassName   apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespace) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdrive]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdrive) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveTypeHyperdrive ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindInherit struct {
	// The name of the inherited binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindInheritType `json:"type,required"`
	// The old name of the inherited binding. If set, the binding will be renamed from
	// `old_name` to `name` in the new version. If not set, the binding will keep the
	// same name between versions.
	OldName string `json:"old_name"`
	// Identifier for the version to inherit the binding from, which can be the version
	// ID or the literal "latest" to inherit from the latest version. Defaults to
	// inheriting the binding from the latest version.
	VersionID string                                                                 `json:"version_id"`
	JSON      scriptVersionGetResponseResourcesBindingsWorkersBindingKindInheritJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindInheritJSON contains
// the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindInherit]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindInheritJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	OldName     apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindInherit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindInheritJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindInherit) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindInheritType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindInheritTypeInherit ScriptVersionGetResponseResourcesBindingsWorkersBindingKindInheritType = "inherit"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindInheritType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindInheritTypeInherit:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindImages struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindImagesType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindImagesJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindImagesJSON contains
// the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindImages]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindImagesJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindImages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindImagesJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindImages) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindImagesType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindImagesTypeImages ScriptVersionGetResponseResourcesBindingsWorkersBindingKindImagesType = "images"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindImagesType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindImagesTypeImages:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json string `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonJSON contains the
// JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJson]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonJSON struct {
	Json        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJson) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonTypeJson ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonType = "json"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespace]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespace) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceTypeKVNamespace ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificate]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateJSON struct {
	CertificateID apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificate) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The text value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainText]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainText) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextTypePlainText ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline string `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelines]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesJSON struct {
	Name        apijson.Field
	Pipeline    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelines) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelines) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesTypePipelines ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesType = "pipelines"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueJSON contains
// the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueue]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueue) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueTypeQueue ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueType = "queue"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketType `json:"type,required"`
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketJurisdiction `json:"jurisdiction"`
	JSON         scriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketJSON         `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketJSON contains
// the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2Bucket]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketJSON struct {
	BucketName   apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	Jurisdiction apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2Bucket) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketTypeR2Bucket ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

// The
// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
// of the R2 bucket.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketJurisdiction string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketJurisdictionEu      ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketJurisdiction = "eu"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketJurisdictionFedramp ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketJurisdiction = "fedramp"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketJurisdiction) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketJurisdictionEu, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindR2BucketJurisdictionFedramp:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretText]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretText) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextTypeSecretText ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextType = "secret_text"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSendEmail struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSendEmailType `json:"type,required"`
	// List of allowed destination addresses.
	AllowedDestinationAddresses []string `json:"allowed_destination_addresses" format:"email"`
	// List of allowed sender addresses.
	AllowedSenderAddresses []string `json:"allowed_sender_addresses" format:"email"`
	// Destination address for the email.
	DestinationAddress string                                                                   `json:"destination_address" format:"email"`
	JSON               scriptVersionGetResponseResourcesBindingsWorkersBindingKindSendEmailJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindSendEmailJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSendEmail]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindSendEmailJSON struct {
	Name                        apijson.Field
	Type                        apijson.Field
	AllowedDestinationAddresses apijson.Field
	AllowedSenderAddresses      apijson.Field
	DestinationAddress          apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSendEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindSendEmailJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSendEmail) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSendEmailType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSendEmailTypeSendEmail ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSendEmailType = "send_email"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSendEmailType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSendEmailTypeSendEmail:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindService struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceType `json:"type,required"`
	// Optional environment if the Worker utilizes one.
	Environment string                                                                 `json:"environment"`
	JSON        scriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceJSON contains
// the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindService]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceJSON struct {
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindService) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceTypeService ScriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceType = "service"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTextBlob struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The name of the file containing the text content. Only accepted for
	// `service worker syntax` Workers.
	Part string `json:"part,required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTextBlobType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindTextBlobJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindTextBlobJSON contains
// the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTextBlob]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindTextBlobJSON struct {
	Name        apijson.Field
	Part        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTextBlob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindTextBlobJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTextBlob) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTextBlobType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTextBlobTypeTextBlob ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTextBlobType = "text_blob"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTextBlobType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindTextBlobTypeTextBlob:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorize]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeJSON struct {
	IndexName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorize) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeTypeVectorize ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeType = "vectorize"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadata]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadata) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the secret in the store.
	SecretName string `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretJSON struct {
	Name        apijson.Field
	SecretName  apijson.Field
	StoreID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecret) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormat `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyType `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage `json:"usages,required"`
	JSON   scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKey]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKey) implementsScriptVersionGetResponseResourcesBinding() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormat string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormatRaw   ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormat = "raw"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormatPkcs8 ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormat = "pkcs8"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormatSpki  ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormat = "spki"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormatJwk   ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormatRaw, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormatPkcs8, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormatSpki, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyTypeSecretKey ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyType = "secret_key"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageEncrypt    ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "encrypt"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDecrypt    ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "decrypt"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageSign       ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "sign"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageVerify     ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "verify"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDeriveKey  ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "deriveKey"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDeriveBits ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "deriveBits"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageWrapKey    ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "wrapKey"
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageUnwrapKey  ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageEncrypt, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDecrypt, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageSign, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageVerify, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDeriveKey, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageDeriveBits, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageWrapKey, ScriptVersionGetResponseResourcesBindingsWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowType `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName string `json:"workflow_name,required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName string `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName string                                                                  `json:"script_name"`
	JSON       scriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowJSON contains
// the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflow]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowJSON struct {
	Name         apijson.Field
	Type         apijson.Field
	WorkflowName apijson.Field
	ClassName    apijson.Field
	ScriptName   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflow) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowTypeWorkflow ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowType = "workflow"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWasmModule struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The name of the file containing the WebAssembly module content. Only accepted
	// for `service worker syntax` Workers.
	Part string `json:"part,required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWasmModuleType `json:"type,required"`
	JSON scriptVersionGetResponseResourcesBindingsWorkersBindingKindWasmModuleJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsWorkersBindingKindWasmModuleJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWasmModule]
type scriptVersionGetResponseResourcesBindingsWorkersBindingKindWasmModuleJSON struct {
	Name        apijson.Field
	Part        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWasmModule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsWorkersBindingKindWasmModuleJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWasmModule) implementsScriptVersionGetResponseResourcesBinding() {
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWasmModuleType string

const (
	ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWasmModuleTypeWasmModule ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWasmModuleType = "wasm_module"
)

func (r ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWasmModuleType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsWorkersBindingKindWasmModuleTypeWasmModule:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptVersionGetResponseResourcesBindingsType string

const (
	ScriptVersionGetResponseResourcesBindingsTypeAI                     ScriptVersionGetResponseResourcesBindingsType = "ai"
	ScriptVersionGetResponseResourcesBindingsTypeAnalyticsEngine        ScriptVersionGetResponseResourcesBindingsType = "analytics_engine"
	ScriptVersionGetResponseResourcesBindingsTypeAssets                 ScriptVersionGetResponseResourcesBindingsType = "assets"
	ScriptVersionGetResponseResourcesBindingsTypeBrowser                ScriptVersionGetResponseResourcesBindingsType = "browser"
	ScriptVersionGetResponseResourcesBindingsTypeD1                     ScriptVersionGetResponseResourcesBindingsType = "d1"
	ScriptVersionGetResponseResourcesBindingsTypeDataBlob               ScriptVersionGetResponseResourcesBindingsType = "data_blob"
	ScriptVersionGetResponseResourcesBindingsTypeDispatchNamespace      ScriptVersionGetResponseResourcesBindingsType = "dispatch_namespace"
	ScriptVersionGetResponseResourcesBindingsTypeDurableObjectNamespace ScriptVersionGetResponseResourcesBindingsType = "durable_object_namespace"
	ScriptVersionGetResponseResourcesBindingsTypeHyperdrive             ScriptVersionGetResponseResourcesBindingsType = "hyperdrive"
	ScriptVersionGetResponseResourcesBindingsTypeInherit                ScriptVersionGetResponseResourcesBindingsType = "inherit"
	ScriptVersionGetResponseResourcesBindingsTypeImages                 ScriptVersionGetResponseResourcesBindingsType = "images"
	ScriptVersionGetResponseResourcesBindingsTypeJson                   ScriptVersionGetResponseResourcesBindingsType = "json"
	ScriptVersionGetResponseResourcesBindingsTypeKVNamespace            ScriptVersionGetResponseResourcesBindingsType = "kv_namespace"
	ScriptVersionGetResponseResourcesBindingsTypeMTLSCertificate        ScriptVersionGetResponseResourcesBindingsType = "mtls_certificate"
	ScriptVersionGetResponseResourcesBindingsTypePlainText              ScriptVersionGetResponseResourcesBindingsType = "plain_text"
	ScriptVersionGetResponseResourcesBindingsTypePipelines              ScriptVersionGetResponseResourcesBindingsType = "pipelines"
	ScriptVersionGetResponseResourcesBindingsTypeQueue                  ScriptVersionGetResponseResourcesBindingsType = "queue"
	ScriptVersionGetResponseResourcesBindingsTypeR2Bucket               ScriptVersionGetResponseResourcesBindingsType = "r2_bucket"
	ScriptVersionGetResponseResourcesBindingsTypeSecretText             ScriptVersionGetResponseResourcesBindingsType = "secret_text"
	ScriptVersionGetResponseResourcesBindingsTypeSendEmail              ScriptVersionGetResponseResourcesBindingsType = "send_email"
	ScriptVersionGetResponseResourcesBindingsTypeService                ScriptVersionGetResponseResourcesBindingsType = "service"
	ScriptVersionGetResponseResourcesBindingsTypeTextBlob               ScriptVersionGetResponseResourcesBindingsType = "text_blob"
	ScriptVersionGetResponseResourcesBindingsTypeVectorize              ScriptVersionGetResponseResourcesBindingsType = "vectorize"
	ScriptVersionGetResponseResourcesBindingsTypeVersionMetadata        ScriptVersionGetResponseResourcesBindingsType = "version_metadata"
	ScriptVersionGetResponseResourcesBindingsTypeSecretsStoreSecret     ScriptVersionGetResponseResourcesBindingsType = "secrets_store_secret"
	ScriptVersionGetResponseResourcesBindingsTypeSecretKey              ScriptVersionGetResponseResourcesBindingsType = "secret_key"
	ScriptVersionGetResponseResourcesBindingsTypeWorkflow               ScriptVersionGetResponseResourcesBindingsType = "workflow"
	ScriptVersionGetResponseResourcesBindingsTypeWasmModule             ScriptVersionGetResponseResourcesBindingsType = "wasm_module"
)

func (r ScriptVersionGetResponseResourcesBindingsType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsTypeAI, ScriptVersionGetResponseResourcesBindingsTypeAnalyticsEngine, ScriptVersionGetResponseResourcesBindingsTypeAssets, ScriptVersionGetResponseResourcesBindingsTypeBrowser, ScriptVersionGetResponseResourcesBindingsTypeD1, ScriptVersionGetResponseResourcesBindingsTypeDataBlob, ScriptVersionGetResponseResourcesBindingsTypeDispatchNamespace, ScriptVersionGetResponseResourcesBindingsTypeDurableObjectNamespace, ScriptVersionGetResponseResourcesBindingsTypeHyperdrive, ScriptVersionGetResponseResourcesBindingsTypeInherit, ScriptVersionGetResponseResourcesBindingsTypeImages, ScriptVersionGetResponseResourcesBindingsTypeJson, ScriptVersionGetResponseResourcesBindingsTypeKVNamespace, ScriptVersionGetResponseResourcesBindingsTypeMTLSCertificate, ScriptVersionGetResponseResourcesBindingsTypePlainText, ScriptVersionGetResponseResourcesBindingsTypePipelines, ScriptVersionGetResponseResourcesBindingsTypeQueue, ScriptVersionGetResponseResourcesBindingsTypeR2Bucket, ScriptVersionGetResponseResourcesBindingsTypeSecretText, ScriptVersionGetResponseResourcesBindingsTypeSendEmail, ScriptVersionGetResponseResourcesBindingsTypeService, ScriptVersionGetResponseResourcesBindingsTypeTextBlob, ScriptVersionGetResponseResourcesBindingsTypeVectorize, ScriptVersionGetResponseResourcesBindingsTypeVersionMetadata, ScriptVersionGetResponseResourcesBindingsTypeSecretsStoreSecret, ScriptVersionGetResponseResourcesBindingsTypeSecretKey, ScriptVersionGetResponseResourcesBindingsTypeWorkflow, ScriptVersionGetResponseResourcesBindingsTypeWasmModule:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptVersionGetResponseResourcesBindingsFormat string

const (
	ScriptVersionGetResponseResourcesBindingsFormatRaw   ScriptVersionGetResponseResourcesBindingsFormat = "raw"
	ScriptVersionGetResponseResourcesBindingsFormatPkcs8 ScriptVersionGetResponseResourcesBindingsFormat = "pkcs8"
	ScriptVersionGetResponseResourcesBindingsFormatSpki  ScriptVersionGetResponseResourcesBindingsFormat = "spki"
	ScriptVersionGetResponseResourcesBindingsFormatJwk   ScriptVersionGetResponseResourcesBindingsFormat = "jwk"
)

func (r ScriptVersionGetResponseResourcesBindingsFormat) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsFormatRaw, ScriptVersionGetResponseResourcesBindingsFormatPkcs8, ScriptVersionGetResponseResourcesBindingsFormatSpki, ScriptVersionGetResponseResourcesBindingsFormatJwk:
		return true
	}
	return false
}

// The
// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
// of the R2 bucket.
type ScriptVersionGetResponseResourcesBindingsJurisdiction string

const (
	ScriptVersionGetResponseResourcesBindingsJurisdictionEu      ScriptVersionGetResponseResourcesBindingsJurisdiction = "eu"
	ScriptVersionGetResponseResourcesBindingsJurisdictionFedramp ScriptVersionGetResponseResourcesBindingsJurisdiction = "fedramp"
)

func (r ScriptVersionGetResponseResourcesBindingsJurisdiction) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesBindingsJurisdictionEu, ScriptVersionGetResponseResourcesBindingsJurisdictionFedramp:
		return true
	}
	return false
}

type ScriptVersionGetResponseResourcesScript struct {
	// Hashed script content
	Etag string `json:"etag"`
	// The names of handlers exported as part of the default export.
	Handlers []string `json:"handlers"`
	// The client most recently used to deploy this Worker.
	LastDeployedFrom string `json:"last_deployed_from"`
	// Named exports, such as Durable Object class implementations and named
	// entrypoints.
	NamedHandlers []ScriptVersionGetResponseResourcesScriptNamedHandler `json:"named_handlers"`
	JSON          scriptVersionGetResponseResourcesScriptJSON           `json:"-"`
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
	// The names of handlers exported as part of the named export.
	Handlers []string `json:"handlers"`
	// The name of the exported class or entrypoint.
	Name string                                                  `json:"name"`
	JSON scriptVersionGetResponseResourcesScriptNamedHandlerJSON `json:"-"`
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

// Runtime configuration for the Worker.
type ScriptVersionGetResponseResourcesScriptRuntime struct {
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate string `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Resource limits for the Worker.
	Limits ScriptVersionGetResponseResourcesScriptRuntimeLimits `json:"limits"`
	// The tag of the Durable Object migration that was most recently applied for this
	// Worker.
	MigrationTag string `json:"migration_tag"`
	// Usage model for the Worker invocations.
	UsageModel ScriptVersionGetResponseResourcesScriptRuntimeUsageModel `json:"usage_model"`
	JSON       scriptVersionGetResponseResourcesScriptRuntimeJSON       `json:"-"`
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

// Resource limits for the Worker.
type ScriptVersionGetResponseResourcesScriptRuntimeLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
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

// Usage model for the Worker invocations.
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
	// Email of the user who created the version.
	AuthorEmail string `json:"author_email"`
	// Identifier of the user who created the version.
	AuthorID string `json:"author_id"`
	// When the version was created.
	CreatedOn string `json:"created_on"`
	// Whether the version can be previewed.
	HasPreview bool `json:"hasPreview"`
	// When the version was last modified.
	ModifiedOn string `json:"modified_on"`
	// The source of the version upload.
	Source ScriptVersionGetResponseMetadataSource `json:"source"`
	JSON   scriptVersionGetResponseMetadataJSON   `json:"-"`
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

// The source of the version upload.
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
	// JSON-encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[ScriptVersionNewParamsMetadata] `json:"metadata,required"`
	// An array of modules (often JavaScript files) comprising a Worker script. At
	// least one module must be present and referenced in the metadata as `main_module`
	// or `body_part` by filename.<br/>Possible Content-Type(s) are:
	// `application/javascript+module`, `text/javascript+module`,
	// `application/javascript`, `text/javascript`, `text/x-python`,
	// `text/x-python-requirement`, `application/wasm`, `text/plain`,
	// `application/octet-stream`, `application/source-map`.
	Files param.Field[[]io.Reader] `json:"files" format:"binary"`
}

func (r ScriptVersionNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRootWithJSON(r, writer)
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
type ScriptVersionNewParamsMetadata struct {
	// Name of the uploaded file that contains the main module (e.g. the file exporting
	// a `fetch` handler). Indicates a `module syntax` Worker, which is required for
	// Version Upload.
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
	// Associated alias for a version.
	WorkersAlias param.Field[string] `json:"workers/alias"`
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
	ID                          param.Field[string]      `json:"id"`
	Algorithm                   param.Field[interface{}] `json:"algorithm"`
	AllowedDestinationAddresses param.Field[interface{}] `json:"allowed_destination_addresses"`
	AllowedSenderAddresses      param.Field[interface{}] `json:"allowed_sender_addresses"`
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name"`
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id"`
	// The exported class name of the Durable Object.
	ClassName param.Field[string] `json:"class_name"`
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset"`
	// Destination address for the email.
	DestinationAddress param.Field[string] `json:"destination_address" format:"email"`
	// The environment of the script_name to bind to.
	Environment param.Field[string] `json:"environment"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[ScriptVersionNewParamsMetadataBindingsFormat] `json:"format"`
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name"`
	// JSON data to use.
	Json param.Field[string] `json:"json"`
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction param.Field[ScriptVersionNewParamsMetadataBindingsJurisdiction] `json:"jurisdiction"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string]      `json:"key_base64"`
	KeyJwk    param.Field[interface{}] `json:"key_jwk"`
	// The name of the dispatch namespace.
	Namespace param.Field[string] `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id"`
	// The old name of the inherited binding. If set, the binding will be renamed from
	// `old_name` to `name` in the new version. If not set, the binding will keep the
	// same name between versions.
	OldName  param.Field[string]      `json:"old_name"`
	Outbound param.Field[interface{}] `json:"outbound"`
	// The name of the file containing the data content. Only accepted for
	// `service worker syntax` Workers.
	Part param.Field[string] `json:"part"`
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
	// Identifier for the version to inherit the binding from, which can be the version
	// ID or the literal "latest" to inherit from the latest version. Defaults to
	// inheriting the binding from the latest version.
	VersionID param.Field[string] `json:"version_id"`
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
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDataBlob],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespace],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdrive],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindInherit],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindImages],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJson],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespace],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificate],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainText],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelines],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueue],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2Bucket],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretText],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSendEmail],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindService],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTextBlob],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorize],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadata],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKey],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflow],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWasmModule],
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

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDataBlob struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The name of the file containing the data content. Only accepted for
	// `service worker syntax` Workers.
	Part param.Field[string] `json:"part,required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDataBlobType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDataBlob) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDataBlob) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDataBlobType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDataBlobTypeDataBlob ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDataBlobType = "data_blob"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDataBlobType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDataBlobTypeDataBlob:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The name of the dispatch namespace.
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

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindInherit struct {
	// The name of the inherited binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindInheritType] `json:"type,required"`
	// The old name of the inherited binding. If set, the binding will be renamed from
	// `old_name` to `name` in the new version. If not set, the binding will keep the
	// same name between versions.
	OldName param.Field[string] `json:"old_name"`
	// Identifier for the version to inherit the binding from, which can be the version
	// ID or the literal "latest" to inherit from the latest version. Defaults to
	// inheriting the binding from the latest version.
	VersionID param.Field[string] `json:"version_id"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindInherit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindInherit) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindInheritType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindInheritTypeInherit ScriptVersionNewParamsMetadataBindingsWorkersBindingKindInheritType = "inherit"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindInheritType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindInheritTypeInherit:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindImages struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindImagesType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindImages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindImages) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindImagesType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindImagesTypeImages ScriptVersionNewParamsMetadataBindingsWorkersBindingKindImagesType = "images"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindImagesType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindImagesTypeImages:
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
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdiction] `json:"jurisdiction"`
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

// The
// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
// of the R2 bucket.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdiction string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdictionEu      ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdiction = "eu"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdictionFedramp ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdiction = "fedramp"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdiction) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdictionEu, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdictionFedramp:
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

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSendEmail struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSendEmailType] `json:"type,required"`
	// List of allowed destination addresses.
	AllowedDestinationAddresses param.Field[[]string] `json:"allowed_destination_addresses" format:"email"`
	// List of allowed sender addresses.
	AllowedSenderAddresses param.Field[[]string] `json:"allowed_sender_addresses" format:"email"`
	// Destination address for the email.
	DestinationAddress param.Field[string] `json:"destination_address" format:"email"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSendEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSendEmail) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSendEmailType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSendEmailTypeSendEmail ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSendEmailType = "send_email"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSendEmailType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSendEmailTypeSendEmail:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindService struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType] `json:"type,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
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

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTextBlob struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The name of the file containing the text content. Only accepted for
	// `service worker syntax` Workers.
	Part param.Field[string] `json:"part,required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTextBlobType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTextBlob) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTextBlob) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTextBlobType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTextBlobTypeTextBlob ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTextBlobType = "text_blob"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTextBlobType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTextBlobTypeTextBlob:
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

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWasmModule struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The name of the file containing the WebAssembly module content. Only accepted
	// for `service worker syntax` Workers.
	Part param.Field[string] `json:"part,required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWasmModuleType] `json:"type,required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWasmModule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWasmModule) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWasmModuleType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWasmModuleTypeWasmModule ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWasmModuleType = "wasm_module"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWasmModuleType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWasmModuleTypeWasmModule:
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
	ScriptVersionNewParamsMetadataBindingsTypeDataBlob               ScriptVersionNewParamsMetadataBindingsType = "data_blob"
	ScriptVersionNewParamsMetadataBindingsTypeDispatchNamespace      ScriptVersionNewParamsMetadataBindingsType = "dispatch_namespace"
	ScriptVersionNewParamsMetadataBindingsTypeDurableObjectNamespace ScriptVersionNewParamsMetadataBindingsType = "durable_object_namespace"
	ScriptVersionNewParamsMetadataBindingsTypeHyperdrive             ScriptVersionNewParamsMetadataBindingsType = "hyperdrive"
	ScriptVersionNewParamsMetadataBindingsTypeInherit                ScriptVersionNewParamsMetadataBindingsType = "inherit"
	ScriptVersionNewParamsMetadataBindingsTypeImages                 ScriptVersionNewParamsMetadataBindingsType = "images"
	ScriptVersionNewParamsMetadataBindingsTypeJson                   ScriptVersionNewParamsMetadataBindingsType = "json"
	ScriptVersionNewParamsMetadataBindingsTypeKVNamespace            ScriptVersionNewParamsMetadataBindingsType = "kv_namespace"
	ScriptVersionNewParamsMetadataBindingsTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsTypePlainText              ScriptVersionNewParamsMetadataBindingsType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsTypePipelines              ScriptVersionNewParamsMetadataBindingsType = "pipelines"
	ScriptVersionNewParamsMetadataBindingsTypeQueue                  ScriptVersionNewParamsMetadataBindingsType = "queue"
	ScriptVersionNewParamsMetadataBindingsTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsTypeSecretText             ScriptVersionNewParamsMetadataBindingsType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsTypeSendEmail              ScriptVersionNewParamsMetadataBindingsType = "send_email"
	ScriptVersionNewParamsMetadataBindingsTypeService                ScriptVersionNewParamsMetadataBindingsType = "service"
	ScriptVersionNewParamsMetadataBindingsTypeTextBlob               ScriptVersionNewParamsMetadataBindingsType = "text_blob"
	ScriptVersionNewParamsMetadataBindingsTypeVectorize              ScriptVersionNewParamsMetadataBindingsType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsType = "version_metadata"
	ScriptVersionNewParamsMetadataBindingsTypeSecretsStoreSecret     ScriptVersionNewParamsMetadataBindingsType = "secrets_store_secret"
	ScriptVersionNewParamsMetadataBindingsTypeSecretKey              ScriptVersionNewParamsMetadataBindingsType = "secret_key"
	ScriptVersionNewParamsMetadataBindingsTypeWorkflow               ScriptVersionNewParamsMetadataBindingsType = "workflow"
	ScriptVersionNewParamsMetadataBindingsTypeWasmModule             ScriptVersionNewParamsMetadataBindingsType = "wasm_module"
)

func (r ScriptVersionNewParamsMetadataBindingsType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsTypeAI, ScriptVersionNewParamsMetadataBindingsTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsTypeAssets, ScriptVersionNewParamsMetadataBindingsTypeBrowser, ScriptVersionNewParamsMetadataBindingsTypeD1, ScriptVersionNewParamsMetadataBindingsTypeDataBlob, ScriptVersionNewParamsMetadataBindingsTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsTypeInherit, ScriptVersionNewParamsMetadataBindingsTypeImages, ScriptVersionNewParamsMetadataBindingsTypeJson, ScriptVersionNewParamsMetadataBindingsTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsTypePlainText, ScriptVersionNewParamsMetadataBindingsTypePipelines, ScriptVersionNewParamsMetadataBindingsTypeQueue, ScriptVersionNewParamsMetadataBindingsTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsTypeSecretText, ScriptVersionNewParamsMetadataBindingsTypeSendEmail, ScriptVersionNewParamsMetadataBindingsTypeService, ScriptVersionNewParamsMetadataBindingsTypeTextBlob, ScriptVersionNewParamsMetadataBindingsTypeVectorize, ScriptVersionNewParamsMetadataBindingsTypeVersionMetadata, ScriptVersionNewParamsMetadataBindingsTypeSecretsStoreSecret, ScriptVersionNewParamsMetadataBindingsTypeSecretKey, ScriptVersionNewParamsMetadataBindingsTypeWorkflow, ScriptVersionNewParamsMetadataBindingsTypeWasmModule:
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

// The
// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
// of the R2 bucket.
type ScriptVersionNewParamsMetadataBindingsJurisdiction string

const (
	ScriptVersionNewParamsMetadataBindingsJurisdictionEu      ScriptVersionNewParamsMetadataBindingsJurisdiction = "eu"
	ScriptVersionNewParamsMetadataBindingsJurisdictionFedramp ScriptVersionNewParamsMetadataBindingsJurisdiction = "fedramp"
)

func (r ScriptVersionNewParamsMetadataBindingsJurisdiction) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsJurisdictionEu, ScriptVersionNewParamsMetadataBindingsJurisdictionFedramp:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type ScriptVersionNewParamsMetadataUsageModel string

const (
	ScriptVersionNewParamsMetadataUsageModelStandard ScriptVersionNewParamsMetadataUsageModel = "standard"
	ScriptVersionNewParamsMetadataUsageModelBundled  ScriptVersionNewParamsMetadataUsageModel = "bundled"
	ScriptVersionNewParamsMetadataUsageModelUnbound  ScriptVersionNewParamsMetadataUsageModel = "unbound"
)

func (r ScriptVersionNewParamsMetadataUsageModel) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataUsageModelStandard, ScriptVersionNewParamsMetadataUsageModelBundled, ScriptVersionNewParamsMetadataUsageModelUnbound:
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
