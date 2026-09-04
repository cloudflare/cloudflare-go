// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/tidwall/gjson"
)

// ScriptScriptAndVersionSettingService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptScriptAndVersionSettingService] method instead.
type ScriptScriptAndVersionSettingService struct {
	Options []option.RequestOption
}

// NewScriptScriptAndVersionSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewScriptScriptAndVersionSettingService(opts ...option.RequestOption) (r *ScriptScriptAndVersionSettingService) {
	r = &ScriptScriptAndVersionSettingService{}
	r.Options = opts
	return
}

// Patch metadata or config, such as bindings or usage model.
func (r *ScriptScriptAndVersionSettingService) Edit(ctx context.Context, scriptName string, params ScriptScriptAndVersionSettingEditParams, opts ...option.RequestOption) (res *ScriptScriptAndVersionSettingEditResponse, err error) {
	var env ScriptScriptAndVersionSettingEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get metadata and config, such as bindings or usage model.
func (r *ScriptScriptAndVersionSettingService) Get(ctx context.Context, scriptName string, query ScriptScriptAndVersionSettingGetParams, opts ...option.RequestOption) (res *ScriptScriptAndVersionSettingGetResponse, err error) {
	var env ScriptScriptAndVersionSettingGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ScriptScriptAndVersionSettingEditResponse struct {
	// Annotations for the Worker version. Annotations are not inherited across
	// settings updates; omitting this field means the new version will have no
	// annotations.
	Annotations ScriptScriptAndVersionSettingEditResponseAnnotations `json:"annotations"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings []ScriptScriptAndVersionSettingEditResponseBinding `json:"bindings"`
	// Global CacheW configuration for the Worker. When caching is on, the platform
	// provisions a `cloudflare.app` zone for the Worker. A `type: worker` entry in the
	// `exports` map can override this value for a single entrypoint.
	CacheOptions ScriptScriptAndVersionSettingEditResponseCacheOptions `json:"cache_options"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate string `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Summary of the declarative exports reconciliation that ran on this upload.
	// Populated only when the uploaded metadata included an `exports` block. Durable
	// Object entries drive reconciliation; `type: worker` entries do not contribute to
	// this summary.
	ExportsReconciliation ScriptScriptAndVersionSettingEditResponseExportsReconciliation `json:"exports_reconciliation"`
	// Limits to apply for this Worker.
	Limits ScriptScriptAndVersionSettingEditResponseLimits `json:"limits"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Observability settings for the Worker.
	Observability ScriptScriptAndVersionSettingEditResponseObservability `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
	Placement ScriptScriptAndVersionSettingEditResponsePlacement `json:"placement"`
	// Tags associated with the Worker.
	Tags []string `json:"tags" api:"nullable"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ConsumerScript `json:"tail_consumers" api:"nullable"`
	// Usage model for the Worker invocations.
	UsageModel ScriptScriptAndVersionSettingEditResponseUsageModel `json:"usage_model"`
	JSON       scriptScriptAndVersionSettingEditResponseJSON       `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseJSON contains the JSON metadata for the
// struct [ScriptScriptAndVersionSettingEditResponse]
type scriptScriptAndVersionSettingEditResponseJSON struct {
	Annotations           apijson.Field
	Bindings              apijson.Field
	CacheOptions          apijson.Field
	CompatibilityDate     apijson.Field
	CompatibilityFlags    apijson.Field
	ExportsReconciliation apijson.Field
	Limits                apijson.Field
	Logpush               apijson.Field
	Observability         apijson.Field
	Placement             apijson.Field
	Tags                  apijson.Field
	TailConsumers         apijson.Field
	UsageModel            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseJSON) RawJSON() string {
	return r.raw
}

// Annotations for the Worker version. Annotations are not inherited across
// settings updates; omitting this field means the new version will have no
// annotations.
type ScriptScriptAndVersionSettingEditResponseAnnotations struct {
	// Human-readable message about the version. Truncated to 1000 bytes if longer.
	WorkersMessage string `json:"workers/message"`
	// User-provided identifier for the version. Maximum 100 bytes.
	WorkersTag string `json:"workers/tag"`
	// Operation that triggered the creation of the version. This is read-only and set
	// by the server.
	WorkersTriggeredBy string                                                   `json:"workers/triggered_by"`
	JSON               scriptScriptAndVersionSettingEditResponseAnnotationsJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseAnnotationsJSON contains the JSON
// metadata for the struct [ScriptScriptAndVersionSettingEditResponseAnnotations]
type scriptScriptAndVersionSettingEditResponseAnnotationsJSON struct {
	WorkersMessage     apijson.Field
	WorkersTag         apijson.Field
	WorkersTriggeredBy apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseAnnotations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseAnnotationsJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources.
type ScriptScriptAndVersionSettingEditResponseBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsType `json:"type" api:"required"`
	// Identifier of the D1 database to bind to.
	//
	// Deprecated: This property has been renamed to `database_id`.
	ID string `json:"id"`
	// This field can have the runtime type of [interface{}].
	Algorithm interface{} `json:"algorithm"`
	// This field can have the runtime type of [[]string].
	AllowedDestinationAddresses interface{} `json:"allowed_destination_addresses"`
	// This field can have the runtime type of [[]string].
	AllowedSenderAddresses interface{} `json:"allowed_sender_addresses"`
	// ID of the Flagship app to bind to for feature flag evaluation.
	AppID string `json:"app_id"`
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name"`
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// Identifier of the D1 database to bind to.
	DatabaseID string `json:"database_id"`
	// The name of the dataset to bind to.
	Dataset string `json:"dataset"`
	// Destination address for the email.
	DestinationAddress string `json:"destination_address" format:"email"`
	// The dispatch namespace the Durable Object script belongs to.
	DispatchNamespace string `json:"dispatch_namespace"`
	// Entrypoint to invoke on the target Worker.
	Entrypoint string `json:"entrypoint"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptScriptAndVersionSettingEditResponseBindingsFormat `json:"format"`
	// Enables Gateway identity for the binding. Requires network_id to be
	// "cf1:network" and cannot be combined with tunnel_id.
	Identity ScriptScriptAndVersionSettingEditResponseBindingsIdentity `json:"identity"`
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name"`
	// The user-chosen instance name. Must exist at deploy time. The worker can search,
	// chat, update, and manage items/jobs on this instance.
	InstanceName string `json:"instance_name"`
	// This field can have the runtime type of [interface{}].
	Json interface{} `json:"json"`
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction ScriptScriptAndVersionSettingEditResponseBindingsJurisdiction `json:"jurisdiction"`
	// This field can have the runtime type of [interface{}].
	KeyJwk interface{} `json:"key_jwk"`
	// The namespace the instance belongs to. Defaults to "default" if omitted.
	// Customers who don't use namespaces can simply omit this field.
	Namespace string `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// Identifier of the network to bind to. Only "cf1:network" is currently supported.
	// Mutually exclusive with tunnel_id.
	NetworkID string `json:"network_id"`
	// The old name of the inherited binding. If set, the binding will be renamed from
	// `old_name` to `name` in the new version. If not set, the binding will keep the
	// same name between versions.
	OldName string `json:"old_name"`
	// This field can have the runtime type of
	// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutbound].
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
	// Identifier of the VPC service to bind to.
	ServiceID string `json:"service_id"`
	// This field can have the runtime type of
	// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimitSimple].
	Simple interface{} `json:"simple"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id"`
	// The text value to use.
	Text string `json:"text"`
	// UUID of the Cloudflare Tunnel to bind to. Mutually exclusive with network_id.
	TunnelID string `json:"tunnel_id"`
	// This field can have the runtime type of
	// [[]ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage].
	Usages interface{} `json:"usages"`
	// Identifier for the version to inherit the binding from, which can be the version
	// ID or the literal "latest" to inherit from the latest version. Defaults to
	// inheriting the binding from the latest version.
	VersionID string `json:"version_id"`
	// Name of the Workflow to bind to.
	WorkflowName string                                               `json:"workflow_name"`
	JSON         scriptScriptAndVersionSettingEditResponseBindingJSON `json:"-"`
	union        ScriptScriptAndVersionSettingEditResponseBindingsUnion
}

// scriptScriptAndVersionSettingEditResponseBindingJSON contains the JSON metadata
// for the struct [ScriptScriptAndVersionSettingEditResponseBinding]
type scriptScriptAndVersionSettingEditResponseBindingJSON struct {
	Name                        apijson.Field
	Type                        apijson.Field
	ID                          apijson.Field
	Algorithm                   apijson.Field
	AllowedDestinationAddresses apijson.Field
	AllowedSenderAddresses      apijson.Field
	AppID                       apijson.Field
	BucketName                  apijson.Field
	CertificateID               apijson.Field
	ClassName                   apijson.Field
	DatabaseID                  apijson.Field
	Dataset                     apijson.Field
	DestinationAddress          apijson.Field
	DispatchNamespace           apijson.Field
	Entrypoint                  apijson.Field
	Environment                 apijson.Field
	Format                      apijson.Field
	Identity                    apijson.Field
	IndexName                   apijson.Field
	InstanceName                apijson.Field
	Json                        apijson.Field
	Jurisdiction                apijson.Field
	KeyJwk                      apijson.Field
	Namespace                   apijson.Field
	NamespaceID                 apijson.Field
	NetworkID                   apijson.Field
	OldName                     apijson.Field
	Outbound                    apijson.Field
	Part                        apijson.Field
	Pipeline                    apijson.Field
	QueueName                   apijson.Field
	ScriptName                  apijson.Field
	SecretName                  apijson.Field
	Service                     apijson.Field
	ServiceID                   apijson.Field
	Simple                      apijson.Field
	StoreID                     apijson.Field
	Text                        apijson.Field
	TunnelID                    apijson.Field
	Usages                      apijson.Field
	VersionID                   apijson.Field
	WorkflowName                apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r scriptScriptAndVersionSettingEditResponseBindingJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptScriptAndVersionSettingEditResponseBinding) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptScriptAndVersionSettingEditResponseBinding{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptScriptAndVersionSettingEditResponseBindingsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearch],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchNamespace],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMessaging],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowser],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDataBlob],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindInherit],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindImages],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMedia],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPipelines],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimit],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSendEmail],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTextBlob],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecret],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindFlagship],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKey],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWorkflow],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWasmModule],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCService],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetwork].
func (r ScriptScriptAndVersionSettingEditResponseBinding) AsUnion() ScriptScriptAndVersionSettingEditResponseBindingsUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources.
//
// Union satisfied by
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearch],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchNamespace],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMessaging],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowser],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDataBlob],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindInherit],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindImages],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMedia],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPipelines],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimit],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSendEmail],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTextBlob],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecret],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindFlagship],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKey],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWorkflow],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWasmModule],
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCService]
// or
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetwork].
type ScriptScriptAndVersionSettingEditResponseBindingsUnion interface {
	implementsScriptScriptAndVersionSettingEditResponseBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptScriptAndVersionSettingEditResponseBindingsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearch{}),
			DiscriminatorValue: "ai_search",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchNamespace{}),
			DiscriminatorValue: "ai_search_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMessaging{}),
			DiscriminatorValue: "messaging",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowser{}),
			DiscriminatorValue: "browser",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDataBlob{}),
			DiscriminatorValue: "data_blob",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindInherit{}),
			DiscriminatorValue: "inherit",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindImages{}),
			DiscriminatorValue: "images",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMedia{}),
			DiscriminatorValue: "media",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPipelines{}),
			DiscriminatorValue: "pipelines",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimit{}),
			DiscriminatorValue: "ratelimit",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSendEmail{}),
			DiscriminatorValue: "send_email",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTextBlob{}),
			DiscriminatorValue: "text_blob",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecret{}),
			DiscriminatorValue: "secrets_store_secret",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindFlagship{}),
			DiscriminatorValue: "flagship",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWorkflow{}),
			DiscriminatorValue: "workflow",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWasmModule{}),
			DiscriminatorValue: "wasm_module",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCService{}),
			DiscriminatorValue: "vpc_service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetwork{}),
			DiscriminatorValue: "vpc_network",
		},
	)
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeAI ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "ai"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearch struct {
	// The user-chosen instance name. Must exist at deploy time. The worker can search,
	// chat, update, and manage items/jobs on this instance.
	InstanceName string `json:"instance_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchType `json:"type" api:"required"`
	// The namespace the instance belongs to. Defaults to "default" if omitted.
	// Customers who don't use namespaces can simply omit this field.
	Namespace string                                                                          `json:"namespace"`
	JSON      scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearch]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchJSON struct {
	InstanceName apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	Namespace    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearch) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchTypeAISearch ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchType = "ai_search"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchTypeAISearch:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The user-chosen namespace name. Must exist before deploy -- Wrangler handles
	// auto-creation on deploy failure (R2 bucket pattern). The "default" namespace is
	// auto-created by config-api for new accounts. Grants full access (CRUD + search +
	// chat) to all instances within the namespace.
	Namespace string `json:"namespace" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchNamespaceType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchNamespaceJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchNamespace]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchNamespace) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchNamespaceType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchNamespaceTypeAISearchNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchNamespaceType = "ai_search_namespace"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAISearchNamespaceTypeAISearchNamespace:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMessaging struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The Messaging namespace to bind to.
	Namespace string `json:"namespace" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMessagingType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMessagingJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMessagingJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMessaging]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMessagingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMessaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMessagingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMessaging) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMessagingType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMessagingTypeMessaging ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMessagingType = "messaging"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMessagingType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMessagingTypeMessaging:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineJSON struct {
	Dataset     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeAssets ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "assets"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowser]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowser) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserTypeBrowser ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserType = "browser"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	DatabaseID string `json:"database_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type `json:"type" api:"required"`
	// Identifier of the D1 database to bind to.
	//
	// Deprecated: This property has been renamed to `database_id`.
	ID   string                                                                    `json:"id"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1JSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1JSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1JSON struct {
	DatabaseID  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1JSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeD1 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "d1"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDataBlob struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The name of the file containing the data content. Only accepted for
	// `service worker syntax` Workers.
	Part string `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDataBlobType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDataBlobJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDataBlobJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDataBlob]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDataBlobJSON struct {
	Name        apijson.Field
	Part        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDataBlob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDataBlobJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDataBlob) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDataBlobType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDataBlobTypeDataBlob ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDataBlobType = "data_blob"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDataBlobType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDataBlobTypeDataBlob:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The name of the dispatch namespace.
	Namespace string `json:"namespace" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType `json:"type" api:"required"`
	// Outbound worker.
	Outbound ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutbound `json:"outbound"`
	JSON     scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceJSON     `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params []ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundParam `json:"params"`
	// Outbound worker.
	Worker ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker `json:"worker"`
	JSON   scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON   `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutbound]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON) RawJSON() string {
	return r.raw
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundParam struct {
	// Name of the parameter.
	Name string                                                                                                `json:"name" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundParamJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundParamJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundParam]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundParamJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundParam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundParamJSON) RawJSON() string {
	return r.raw
}

// Outbound worker.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Entrypoint to invoke on the outbound worker.
	Entrypoint string `json:"entrypoint"`
	// Environment of the outbound worker.
	Environment string `json:"environment"`
	// Name of the outbound worker.
	Service string                                                                                                 `json:"service"`
	JSON    scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON struct {
	Entrypoint  apijson.Field
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType `json:"type" api:"required"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The dispatch namespace the Durable Object script belongs to.
	DispatchNamespace string `json:"dispatch_namespace"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string                                                                                        `json:"script_name"`
	JSON       scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON struct {
	Name              apijson.Field
	Type              apijson.Field
	ClassName         apijson.Field
	DispatchNamespace apijson.Field
	Environment       apijson.Field
	NamespaceID       apijson.Field
	ScriptName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeHyperdrive ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindInherit struct {
	// The name of the inherited binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindInheritType `json:"type" api:"required"`
	// The old name of the inherited binding. If set, the binding will be renamed from
	// `old_name` to `name` in the new version. If not set, the binding will keep the
	// same name between versions.
	OldName string `json:"old_name"`
	// Identifier for the version to inherit the binding from, which can be the version
	// ID or the literal "latest" to inherit from the latest version. Defaults to
	// inheriting the binding from the latest version.
	VersionID string                                                                         `json:"version_id"`
	JSON      scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindInheritJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindInheritJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindInherit]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindInheritJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	OldName     apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindInherit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindInheritJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindInherit) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindInheritType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindInheritTypeInherit ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindInheritType = "inherit"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindInheritType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindInheritTypeInherit:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindImages struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindImagesType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindImagesJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindImagesJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindImages]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindImagesJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindImages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindImagesJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindImages) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindImagesType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindImagesTypeImages ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindImagesType = "images"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindImagesType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindImagesTypeImages:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json interface{} `json:"json" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonJSON struct {
	Json        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeJson ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "json"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeKVNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMedia struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMediaType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMediaJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMediaJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMedia]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMediaJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMedia) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMediaJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMedia) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMediaType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMediaTypeMedia ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMediaType = "media"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMediaType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMediaTypeMedia:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateJSON struct {
	CertificateID apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The text value to use.
	Text string `json:"text" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypePlainText ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Name of the Pipeline to bind to.
	Pipeline string `json:"pipeline" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPipelinesType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPipelinesJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPipelinesJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPipelines]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPipelinesJSON struct {
	Name        apijson.Field
	Pipeline    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPipelines) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPipelinesJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPipelines) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPipelinesType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPipelinesTypePipelines ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPipelinesType = "pipelines"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeQueue ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "queue"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimit struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Identifier of the rate limit namespace to bind to.
	NamespaceID string `json:"namespace_id" api:"required"`
	// The rate limit configuration.
	Simple ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimitSimple `json:"simple" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimitType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimitJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimitJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimit]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimitJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Simple      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimitJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimit) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The rate limit configuration.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimitSimple struct {
	// The limit (requests per period).
	Limit float64 `json:"limit" api:"required"`
	// The period in seconds.
	Period int64 `json:"period" api:"required"`
	// Duration in seconds to apply the mitigation action after the rate limit is
	// exceeded. Valid values are 0 (disabled), 10, or multiples of 60 up to 86400.
	// Must be greater than or equal to the period when non-zero.
	MitigationTimeout int64                                                                                  `json:"mitigation_timeout"`
	JSON              scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimitSimpleJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimitSimpleJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimitSimple]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimitSimpleJSON struct {
	Limit             apijson.Field
	Period            apijson.Field
	MitigationTimeout apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimitSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimitSimpleJSON) RawJSON() string {
	return r.raw
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimitType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimitTypeRatelimit ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimitType = "ratelimit"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimitType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindRatelimitTypeRatelimit:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType `json:"type" api:"required"`
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJurisdiction `json:"jurisdiction"`
	JSON         scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJSON         `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJSON struct {
	BucketName   apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	Jurisdiction apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeR2Bucket ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

// The
// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
// of the R2 bucket.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJurisdiction string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJurisdictionEu          ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJurisdiction = "eu"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJurisdictionFedramp     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJurisdiction = "fedramp"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJurisdictionFedrampHigh ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJurisdiction = "fedramp-high"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJurisdictionUs          ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJurisdiction = "us"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJurisdiction) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJurisdictionEu, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJurisdictionFedramp, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJurisdictionFedrampHigh, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJurisdictionUs:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeSecretText ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "secret_text"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSendEmail struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSendEmailType `json:"type" api:"required"`
	// List of allowed destination addresses.
	AllowedDestinationAddresses []string `json:"allowed_destination_addresses" format:"email"`
	// List of allowed sender addresses.
	AllowedSenderAddresses []string `json:"allowed_sender_addresses" format:"email"`
	// Destination address for the email.
	DestinationAddress string                                                                           `json:"destination_address" format:"email"`
	JSON               scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSendEmailJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSendEmailJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSendEmail]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSendEmailJSON struct {
	Name                        apijson.Field
	Type                        apijson.Field
	AllowedDestinationAddresses apijson.Field
	AllowedSenderAddresses      apijson.Field
	DestinationAddress          apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSendEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSendEmailJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSendEmail) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSendEmailType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSendEmailTypeSendEmail ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSendEmailType = "send_email"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSendEmailType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSendEmailTypeSendEmail:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Name of Worker to bind to.
	Service string `json:"service" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType `json:"type" api:"required"`
	// Entrypoint to invoke on the target Worker.
	Entrypoint string `json:"entrypoint"`
	// Optional environment if the Worker utilizes one.
	Environment string                                                                         `json:"environment"`
	JSON        scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceJSON struct {
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	Entrypoint  apijson.Field
	Environment apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeService ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "service"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTextBlob struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The name of the file containing the text content. Only accepted for
	// `service worker syntax` Workers.
	Part string `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTextBlobType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTextBlobJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTextBlobJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTextBlob]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTextBlobJSON struct {
	Name        apijson.Field
	Part        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTextBlob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTextBlobJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTextBlob) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTextBlobType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTextBlobTypeTextBlob ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTextBlobType = "text_blob"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTextBlobType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTextBlobTypeTextBlob:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeJSON struct {
	IndexName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeVectorize ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "vectorize"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Name of the secret in the store.
	SecretName string `json:"secret_name" api:"required"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecret]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretJSON struct {
	Name        apijson.Field
	SecretName  apijson.Field
	StoreID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecret) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindFlagship struct {
	// ID of the Flagship app to bind to for feature flag evaluation.
	AppID string `json:"app_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindFlagshipType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindFlagshipJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindFlagshipJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindFlagship]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindFlagshipJSON struct {
	AppID       apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindFlagship) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindFlagshipJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindFlagship) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindFlagshipType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindFlagshipTypeFlagship ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindFlagshipType = "flagship"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindFlagshipType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindFlagshipTypeFlagship:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm" api:"required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyFormat `json:"format" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyType `json:"type" api:"required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage `json:"usages" api:"required"`
	JSON   scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKey]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKey) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyFormat string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyFormatRaw   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyFormat = "raw"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyFormatPkcs8 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyFormat = "pkcs8"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyFormatSpki  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyFormat = "spki"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyFormatJwk   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyFormatRaw, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyFormatPkcs8, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyFormatSpki, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyTypeSecretKey ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyType = "secret_key"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageEncrypt    ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage = "encrypt"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageDecrypt    ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage = "decrypt"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageSign       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage = "sign"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageVerify     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage = "verify"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageDeriveKey  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage = "deriveKey"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageDeriveBits ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage = "deriveBits"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageWrapKey    ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage = "wrapKey"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageUnwrapKey  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageEncrypt, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageDecrypt, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageSign, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageVerify, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageDeriveKey, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageDeriveBits, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageWrapKey, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWorkflowType `json:"type" api:"required"`
	// Name of the Workflow to bind to.
	WorkflowName string `json:"workflow_name" api:"required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName string `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName string                                                                          `json:"script_name"`
	JSON       scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWorkflowJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWorkflowJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWorkflow]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWorkflowJSON struct {
	Name         apijson.Field
	Type         apijson.Field
	WorkflowName apijson.Field
	ClassName    apijson.Field
	ScriptName   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWorkflowJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWorkflow) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWorkflowType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWorkflowTypeWorkflow ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWorkflowType = "workflow"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWasmModule struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The name of the file containing the WebAssembly module content. Only accepted
	// for `service worker syntax` Workers.
	Part string `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWasmModuleType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWasmModuleJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWasmModuleJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWasmModule]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWasmModuleJSON struct {
	Name        apijson.Field
	Part        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWasmModule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWasmModuleJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWasmModule) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWasmModuleType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWasmModuleTypeWasmModule ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWasmModuleType = "wasm_module"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWasmModuleType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindWasmModuleTypeWasmModule:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCService struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Identifier of the VPC service to bind to.
	ServiceID string `json:"service_id" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCServiceType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCServiceJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCServiceJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCService]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCServiceJSON struct {
	Name        apijson.Field
	ServiceID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCServiceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCService) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCServiceType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCServiceTypeVPCService ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCServiceType = "vpc_service"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCServiceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCServiceTypeVPCService:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetwork struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetworkType `json:"type" api:"required"`
	// Enables Gateway identity for the binding. Requires network_id to be
	// "cf1:network" and cannot be combined with tunnel_id.
	Identity ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetworkIdentity `json:"identity"`
	// Identifier of the network to bind to. Only "cf1:network" is currently supported.
	// Mutually exclusive with tunnel_id.
	NetworkID string `json:"network_id"`
	// UUID of the Cloudflare Tunnel to bind to. Mutually exclusive with network_id.
	TunnelID string                                                                            `json:"tunnel_id"`
	JSON     scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetworkJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetworkJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetwork]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetworkJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Identity    apijson.Field
	NetworkID   apijson.Field
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetworkJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetwork) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetworkType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetworkTypeVPCNetwork ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetworkType = "vpc_network"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetworkType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetworkTypeVPCNetwork:
		return true
	}
	return false
}

// Enables Gateway identity for the binding. Requires network_id to be
// "cf1:network" and cannot be combined with tunnel_id.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetworkIdentity string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetworkIdentityRuntimeEmailAlpha ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetworkIdentity = "runtime-email-alpha"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetworkIdentity) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVPCNetworkIdentityRuntimeEmailAlpha:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsTypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeAISearch               ScriptScriptAndVersionSettingEditResponseBindingsType = "ai_search"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeAISearchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsType = "ai_search_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeMessaging              ScriptScriptAndVersionSettingEditResponseBindingsType = "messaging"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeBrowser                ScriptScriptAndVersionSettingEditResponseBindingsType = "browser"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeDataBlob               ScriptScriptAndVersionSettingEditResponseBindingsType = "data_blob"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeInherit                ScriptScriptAndVersionSettingEditResponseBindingsType = "inherit"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeImages                 ScriptScriptAndVersionSettingEditResponseBindingsType = "images"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeMedia                  ScriptScriptAndVersionSettingEditResponseBindingsType = "media"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsTypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsTypePipelines              ScriptScriptAndVersionSettingEditResponseBindingsType = "pipelines"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeRatelimit              ScriptScriptAndVersionSettingEditResponseBindingsType = "ratelimit"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeSendEmail              ScriptScriptAndVersionSettingEditResponseBindingsType = "send_email"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeService                ScriptScriptAndVersionSettingEditResponseBindingsType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeTextBlob               ScriptScriptAndVersionSettingEditResponseBindingsType = "text_blob"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsType = "version_metadata"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeSecretsStoreSecret     ScriptScriptAndVersionSettingEditResponseBindingsType = "secrets_store_secret"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeFlagship               ScriptScriptAndVersionSettingEditResponseBindingsType = "flagship"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeSecretKey              ScriptScriptAndVersionSettingEditResponseBindingsType = "secret_key"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeWorkflow               ScriptScriptAndVersionSettingEditResponseBindingsType = "workflow"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeWasmModule             ScriptScriptAndVersionSettingEditResponseBindingsType = "wasm_module"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeVPCService             ScriptScriptAndVersionSettingEditResponseBindingsType = "vpc_service"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeVPCNetwork             ScriptScriptAndVersionSettingEditResponseBindingsType = "vpc_network"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsTypeAI, ScriptScriptAndVersionSettingEditResponseBindingsTypeAISearch, ScriptScriptAndVersionSettingEditResponseBindingsTypeAISearchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsTypeMessaging, ScriptScriptAndVersionSettingEditResponseBindingsTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsTypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsTypeBrowser, ScriptScriptAndVersionSettingEditResponseBindingsTypeD1, ScriptScriptAndVersionSettingEditResponseBindingsTypeDataBlob, ScriptScriptAndVersionSettingEditResponseBindingsTypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsTypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsTypeInherit, ScriptScriptAndVersionSettingEditResponseBindingsTypeImages, ScriptScriptAndVersionSettingEditResponseBindingsTypeJson, ScriptScriptAndVersionSettingEditResponseBindingsTypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsTypeMedia, ScriptScriptAndVersionSettingEditResponseBindingsTypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsTypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsTypePipelines, ScriptScriptAndVersionSettingEditResponseBindingsTypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsTypeRatelimit, ScriptScriptAndVersionSettingEditResponseBindingsTypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsTypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsTypeSendEmail, ScriptScriptAndVersionSettingEditResponseBindingsTypeService, ScriptScriptAndVersionSettingEditResponseBindingsTypeTextBlob, ScriptScriptAndVersionSettingEditResponseBindingsTypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsTypeVersionMetadata, ScriptScriptAndVersionSettingEditResponseBindingsTypeSecretsStoreSecret, ScriptScriptAndVersionSettingEditResponseBindingsTypeFlagship, ScriptScriptAndVersionSettingEditResponseBindingsTypeSecretKey, ScriptScriptAndVersionSettingEditResponseBindingsTypeWorkflow, ScriptScriptAndVersionSettingEditResponseBindingsTypeWasmModule, ScriptScriptAndVersionSettingEditResponseBindingsTypeVPCService, ScriptScriptAndVersionSettingEditResponseBindingsTypeVPCNetwork:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptScriptAndVersionSettingEditResponseBindingsFormat string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsFormatRaw   ScriptScriptAndVersionSettingEditResponseBindingsFormat = "raw"
	ScriptScriptAndVersionSettingEditResponseBindingsFormatPkcs8 ScriptScriptAndVersionSettingEditResponseBindingsFormat = "pkcs8"
	ScriptScriptAndVersionSettingEditResponseBindingsFormatSpki  ScriptScriptAndVersionSettingEditResponseBindingsFormat = "spki"
	ScriptScriptAndVersionSettingEditResponseBindingsFormatJwk   ScriptScriptAndVersionSettingEditResponseBindingsFormat = "jwk"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsFormat) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsFormatRaw, ScriptScriptAndVersionSettingEditResponseBindingsFormatPkcs8, ScriptScriptAndVersionSettingEditResponseBindingsFormatSpki, ScriptScriptAndVersionSettingEditResponseBindingsFormatJwk:
		return true
	}
	return false
}

// Enables Gateway identity for the binding. Requires network_id to be
// "cf1:network" and cannot be combined with tunnel_id.
type ScriptScriptAndVersionSettingEditResponseBindingsIdentity string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsIdentityRuntimeEmailAlpha ScriptScriptAndVersionSettingEditResponseBindingsIdentity = "runtime-email-alpha"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsIdentity) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsIdentityRuntimeEmailAlpha:
		return true
	}
	return false
}

// The
// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
// of the R2 bucket.
type ScriptScriptAndVersionSettingEditResponseBindingsJurisdiction string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsJurisdictionEu          ScriptScriptAndVersionSettingEditResponseBindingsJurisdiction = "eu"
	ScriptScriptAndVersionSettingEditResponseBindingsJurisdictionFedramp     ScriptScriptAndVersionSettingEditResponseBindingsJurisdiction = "fedramp"
	ScriptScriptAndVersionSettingEditResponseBindingsJurisdictionFedrampHigh ScriptScriptAndVersionSettingEditResponseBindingsJurisdiction = "fedramp-high"
	ScriptScriptAndVersionSettingEditResponseBindingsJurisdictionUs          ScriptScriptAndVersionSettingEditResponseBindingsJurisdiction = "us"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsJurisdiction) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsJurisdictionEu, ScriptScriptAndVersionSettingEditResponseBindingsJurisdictionFedramp, ScriptScriptAndVersionSettingEditResponseBindingsJurisdictionFedrampHigh, ScriptScriptAndVersionSettingEditResponseBindingsJurisdictionUs:
		return true
	}
	return false
}

// Global CacheW configuration for the Worker. When caching is on, the platform
// provisions a `cloudflare.app` zone for the Worker. A `type: worker` entry in the
// `exports` map can override this value for a single entrypoint.
type ScriptScriptAndVersionSettingEditResponseCacheOptions struct {
	// Whether caching is enabled for this Worker.
	Enabled bool `json:"enabled" api:"required"`
	// Whether cached responses are shared across Worker version uploads. This is
	// independent of `enabled`. It can stay true while caching is off, so the
	// preference survives turning caching off and back on.
	CrossVersionCache bool                                                      `json:"cross_version_cache"`
	JSON              scriptScriptAndVersionSettingEditResponseCacheOptionsJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseCacheOptionsJSON contains the JSON
// metadata for the struct [ScriptScriptAndVersionSettingEditResponseCacheOptions]
type scriptScriptAndVersionSettingEditResponseCacheOptionsJSON struct {
	Enabled           apijson.Field
	CrossVersionCache apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseCacheOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseCacheOptionsJSON) RawJSON() string {
	return r.raw
}

// A single entry in the `exports` map, keyed by export name (a `WorkerEntrypoint`
// class name, a Durable Object class name, or `default` for the Worker's default
// export). The `type` discriminator selects the top-level shape: `worker`
// entrypoint entries may carry `cache` configuration, while `durable-object`
// entries are further refined by the optional `state` field (default `created`).
// Tombstone states (`deleted`, `renamed`, `transferred`) express destructive
// lifecycle operations declaratively; `expecting-transfer` is the live target side
// of a transfer. The server validates the exact per-(type, state) field
// combinations; fields not listed for a variant are rejected.
type ScriptScriptAndVersionSettingEditResponseExport struct {
	// Marks this entry as a Worker entrypoint export.
	Type ScriptScriptAndVersionSettingEditResponseExportsType `json:"type" api:"required"`
	// This field can have the runtime type of
	// [ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportCache].
	Cache interface{} `json:"cache"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container string `json:"container"`
	// Live export. May be omitted; defaults to `created`.
	State ScriptScriptAndVersionSettingEditResponseExportsState `json:"state"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage ScriptScriptAndVersionSettingEditResponseExportsStorage `json:"storage"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom string                                              `json:"transfer_from"`
	JSON         scriptScriptAndVersionSettingEditResponseExportJSON `json:"-"`
	union        ScriptScriptAndVersionSettingEditResponseExportsUnion
}

// scriptScriptAndVersionSettingEditResponseExportJSON contains the JSON metadata
// for the struct [ScriptScriptAndVersionSettingEditResponseExport]
type scriptScriptAndVersionSettingEditResponseExportJSON struct {
	Type         apijson.Field
	Cache        apijson.Field
	Container    apijson.Field
	State        apijson.Field
	Storage      apijson.Field
	TransferFrom apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r scriptScriptAndVersionSettingEditResponseExportJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptScriptAndVersionSettingEditResponseExport) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptScriptAndVersionSettingEditResponseExport{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptScriptAndVersionSettingEditResponseExportsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExport],
// [ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExport],
// [ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExport],
// [ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExport],
// [ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExport],
// [ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExport].
func (r ScriptScriptAndVersionSettingEditResponseExport) AsUnion() ScriptScriptAndVersionSettingEditResponseExportsUnion {
	return r.union
}

// A single entry in the `exports` map, keyed by export name (a `WorkerEntrypoint`
// class name, a Durable Object class name, or `default` for the Worker's default
// export). The `type` discriminator selects the top-level shape: `worker`
// entrypoint entries may carry `cache` configuration, while `durable-object`
// entries are further refined by the optional `state` field (default `created`).
// Tombstone states (`deleted`, `renamed`, `transferred`) express destructive
// lifecycle operations declaratively; `expecting-transfer` is the live target side
// of a transfer. The server validates the exact per-(type, state) field
// combinations; fields not listed for a variant are rejected.
//
// Union satisfied by
// [ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExport],
// [ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExport],
// [ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExport],
// [ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExport],
// [ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExport]
// or
// [ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExport].
type ScriptScriptAndVersionSettingEditResponseExportsUnion interface {
	implementsScriptScriptAndVersionSettingEditResponseExport()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptScriptAndVersionSettingEditResponseExportsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExport{}),
			DiscriminatorValue: "worker",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExport{}),
			DiscriminatorValue: "durable-object",
		},
	)
}

// A named Worker entrypoint export (`type: worker`). Worker entrypoints are always
// live (`state: created`) and carry no storage or lifecycle fields. The optional
// `cache` block overrides the Worker's global `cache_options.enabled` for this
// entrypoint.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExport struct {
	// Marks this entry as a Worker entrypoint export.
	Type ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportType `json:"type" api:"required"`
	// Cache override for this entrypoint. Overrides the Worker's global
	// `cache_options.enabled` for this entrypoint only.
	Cache ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportCache `json:"cache"`
	// Live export. May be omitted; defaults to `created`.
	State ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportState `json:"state"`
	JSON  scriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportJSON  `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportJSON contains
// the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExport]
type scriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportJSON struct {
	Type        apijson.Field
	Cache       apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExport) implementsScriptScriptAndVersionSettingEditResponseExport() {
}

// Marks this entry as a Worker entrypoint export.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportType string

const (
	ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportTypeWorker ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportType = "worker"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportTypeWorker:
		return true
	}
	return false
}

// Cache override for this entrypoint. Overrides the Worker's global
// `cache_options.enabled` for this entrypoint only.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportCache struct {
	// Whether caching is enabled for this entrypoint.
	Enabled bool                                                                         `json:"enabled" api:"required"`
	JSON    scriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportCacheJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportCacheJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportCache]
type scriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportCacheJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportCacheJSON) RawJSON() string {
	return r.raw
}

// Live export. May be omitted; defaults to `created`.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportState string

const (
	ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportStateCreated ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportState = "created"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsWorkersWorkerExportStateCreated:
		return true
	}
	return false
}

// A live Durable Object export (`state: created`, the default). The platform
// auto-provisions the namespace on first deploy, matches it on subsequent deploys,
// and never mutates or deletes it as a side effect of a code-only change.
// `storage` is required; `renamed_to`, `transferred_to` and `transfer_from` are
// not allowed on a live entry.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExport struct {
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportStorage `json:"storage" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportType `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container string `json:"container"`
	// Live export. May be omitted; defaults to `created`.
	State ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportState `json:"state"`
	JSON  scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportJSON  `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExport]
type scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportJSON struct {
	Storage     apijson.Field
	Type        apijson.Field
	Container   apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExport) implementsScriptScriptAndVersionSettingEditResponseExport() {
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportStorage string

const (
	ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportStorageSqlite   ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportStorage = "sqlite"
	ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportStorageLegacyKV ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportStorage = "legacy-kv"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportStorage) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportStorageSqlite, ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportStorageLegacyKV:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportType string

const (
	ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportTypeDurableObject ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportType = "durable-object"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportTypeDurableObject:
		return true
	}
	return false
}

// Live export. May be omitted; defaults to `created`.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportState string

const (
	ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportStateCreated ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportState = "created"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExportStateCreated:
		return true
	}
	return false
}

// A `deleted` tombstone: retires the provisioned namespace for this class and all
// of its data. The class must be absent from the uploaded code and no other Worker
// in the account may bind to the namespace, otherwise the deploy is rejected. No
// other fields are allowed. Deletion is irreversible.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExport struct {
	// Tombstone that deletes the namespace.
	State ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExportState `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExportType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExportJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExportJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExport]
type scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExportJSON struct {
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExport) implementsScriptScriptAndVersionSettingEditResponseExport() {
}

// Tombstone that deletes the namespace.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExportState string

const (
	ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExportStateDeleted ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExportState = "deleted"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExportState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExportStateDeleted:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExportType string

const (
	ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExportTypeDurableObject ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExportType = "durable-object"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExportType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectDeletedExportTypeDurableObject:
		return true
	}
	return false
}

// A `renamed` tombstone: rewrites the provisioned namespace's class name from this
// map key to `renamed_to`. The source class may stay in code during the rollout
// window (an info notice is emitted). `storage`, `transferred_to` and
// `transfer_from` are not allowed.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExport struct {
	// Tombstone that renames the namespace's class.
	State ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExportState `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExportType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExportJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExportJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExport]
type scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExportJSON struct {
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExport) implementsScriptScriptAndVersionSettingEditResponseExport() {
}

// Tombstone that renames the namespace's class.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExportState string

const (
	ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExportStateRenamed ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExportState = "renamed"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExportState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExportStateRenamed:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExportType string

const (
	ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExportTypeDurableObject ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExportType = "durable-object"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExportType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectRenamedExportTypeDurableObject:
		return true
	}
	return false
}

// A `transferred` tombstone (source side of a two-phase transfer): hands ownership
// of the provisioned namespace to another script in the same account, named by
// `transferred_to`. The target must have already deployed a matching
// `expecting-transfer` entry. The source class may stay in code during the rollout
// window (an info notice is emitted). `storage`, `renamed_to` and `transfer_from`
// are not allowed.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExport struct {
	// Tombstone that transfers the namespace to another script.
	State ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExportState `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExportType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExportJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExportJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExport]
type scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExportJSON struct {
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExport) implementsScriptScriptAndVersionSettingEditResponseExport() {
}

// Tombstone that transfers the namespace to another script.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExportState string

const (
	ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExportStateTransferred ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExportState = "transferred"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExportState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExportStateTransferred:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExportType string

const (
	ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExportTypeDurableObject ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExportType = "durable-object"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExportType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectTransferredExportTypeDurableObject:
		return true
	}
	return false
}

// The target side of a two-phase transfer (`state: expecting-transfer`). Declares
// that this script expects to receive a namespace for this class from the
// `transfer_from` script. This is a live entry, not a tombstone: bindings resolve
// through the source's namespace until the source commits with a `transferred`
// tombstone. `storage` and `transfer_from` are required; `renamed_to` and
// `transferred_to` are not allowed.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExport struct {
	// Target side of a two-phase transfer.
	State ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportState `json:"state" api:"required"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportStorage `json:"storage" api:"required"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom string `json:"transfer_from" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportType `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object once the transfer settles. Valid only on live entries.
	Container string                                                                                          `json:"container"`
	JSON      scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExport]
type scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportJSON struct {
	State        apijson.Field
	Storage      apijson.Field
	TransferFrom apijson.Field
	Type         apijson.Field
	Container    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExport) implementsScriptScriptAndVersionSettingEditResponseExport() {
}

// Target side of a two-phase transfer.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportState string

const (
	ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportStateExpectingTransfer ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportState = "expecting-transfer"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportStateExpectingTransfer:
		return true
	}
	return false
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportStorage string

const (
	ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportStorageSqlite   ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportStorage = "sqlite"
	ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportStorageLegacyKV ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportStorage = "legacy-kv"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportStorage) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportStorageSqlite, ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportStorageLegacyKV:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportType string

const (
	ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportTypeDurableObject ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportType = "durable-object"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsWorkersDurableObjectExpectingTransferExportTypeDurableObject:
		return true
	}
	return false
}

// Marks this entry as a Worker entrypoint export.
type ScriptScriptAndVersionSettingEditResponseExportsType string

const (
	ScriptScriptAndVersionSettingEditResponseExportsTypeWorker        ScriptScriptAndVersionSettingEditResponseExportsType = "worker"
	ScriptScriptAndVersionSettingEditResponseExportsTypeDurableObject ScriptScriptAndVersionSettingEditResponseExportsType = "durable-object"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsTypeWorker, ScriptScriptAndVersionSettingEditResponseExportsTypeDurableObject:
		return true
	}
	return false
}

// Live export. May be omitted; defaults to `created`.
type ScriptScriptAndVersionSettingEditResponseExportsState string

const (
	ScriptScriptAndVersionSettingEditResponseExportsStateCreated           ScriptScriptAndVersionSettingEditResponseExportsState = "created"
	ScriptScriptAndVersionSettingEditResponseExportsStateDeleted           ScriptScriptAndVersionSettingEditResponseExportsState = "deleted"
	ScriptScriptAndVersionSettingEditResponseExportsStateRenamed           ScriptScriptAndVersionSettingEditResponseExportsState = "renamed"
	ScriptScriptAndVersionSettingEditResponseExportsStateTransferred       ScriptScriptAndVersionSettingEditResponseExportsState = "transferred"
	ScriptScriptAndVersionSettingEditResponseExportsStateExpectingTransfer ScriptScriptAndVersionSettingEditResponseExportsState = "expecting-transfer"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsStateCreated, ScriptScriptAndVersionSettingEditResponseExportsStateDeleted, ScriptScriptAndVersionSettingEditResponseExportsStateRenamed, ScriptScriptAndVersionSettingEditResponseExportsStateTransferred, ScriptScriptAndVersionSettingEditResponseExportsStateExpectingTransfer:
		return true
	}
	return false
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type ScriptScriptAndVersionSettingEditResponseExportsStorage string

const (
	ScriptScriptAndVersionSettingEditResponseExportsStorageSqlite   ScriptScriptAndVersionSettingEditResponseExportsStorage = "sqlite"
	ScriptScriptAndVersionSettingEditResponseExportsStorageLegacyKV ScriptScriptAndVersionSettingEditResponseExportsStorage = "legacy-kv"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsStorage) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsStorageSqlite, ScriptScriptAndVersionSettingEditResponseExportsStorageLegacyKV:
		return true
	}
	return false
}

// Summary of the declarative exports reconciliation that ran on this upload.
// Populated only when the uploaded metadata included an `exports` block. Durable
// Object entries drive reconciliation; `type: worker` entries do not contribute to
// this summary.
type ScriptScriptAndVersionSettingEditResponseExportsReconciliation struct {
	// Class names for which a new namespace was provisioned.
	Created []string `json:"created" api:"required"`
	// Class names whose namespace was deleted by a `deleted` tombstone.
	Deleted []string `json:"deleted" api:"required"`
	// Non-blocking info entries (stale tombstones, tombstone applied with class still
	// in code). See `exports_reconciliation_info`.
	Info []ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfo `json:"info" api:"required"`
	// Source class names whose tombstone entry is now stale and safe to delete from
	// `exports` (no remaining referencing scripts).
	RemovableEntries []string `json:"removable_entries" api:"required"`
	// Applied `renamed` tombstones.
	Renamed []ScriptScriptAndVersionSettingEditResponseExportsReconciliationRenamed `json:"renamed" api:"required"`
	// Phase-1 transfer hints recorded on the target side.
	TransferPending []ScriptScriptAndVersionSettingEditResponseExportsReconciliationTransferPending `json:"transfer_pending" api:"required"`
	// Committed `transferred` tombstones (phase-2).
	Transferred []ScriptScriptAndVersionSettingEditResponseExportsReconciliationTransferred `json:"transferred" api:"required"`
	// Class names whose provisioned namespace was mutated in place.
	Updated []string `json:"updated" api:"required"`
	// Non-blocking warnings. See `exports_reconciliation_warning`.
	Warnings []ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarning `json:"warnings" api:"required"`
	JSON     scriptScriptAndVersionSettingEditResponseExportsReconciliationJSON      `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseExportsReconciliationJSON contains the
// JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseExportsReconciliation]
type scriptScriptAndVersionSettingEditResponseExportsReconciliationJSON struct {
	Created          apijson.Field
	Deleted          apijson.Field
	Info             apijson.Field
	RemovableEntries apijson.Field
	Renamed          apijson.Field
	TransferPending  apijson.Field
	Transferred      apijson.Field
	Updated          apijson.Field
	Warnings         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseExportsReconciliation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseExportsReconciliationJSON) RawJSON() string {
	return r.raw
}

// A non-blocking reconciliation info entry. Emitted for stale tombstones (a no-op
// on this deploy) and for tombstones applied with the source class still in code
// (the supported zero-downtime rollout pattern).
type ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfo struct {
	// The class name the info entry is about.
	Class string `json:"class" api:"required"`
	// Human-readable explanation.
	Message string `json:"message" api:"required"`
	// Stable, machine-readable tag identifying which reconciliation scenario produced
	// an error, warning, or info entry. Clients may branch on this value instead of
	// parsing `message`.
	Scenario ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario `json:"scenario" api:"required"`
	// The provisioned namespace the entry relates to, when applicable.
	NamespaceID string `json:"namespace_id" format:"uuid"`
	// Other Workers in the account that still bind to the affected class. Advisory:
	// while non-empty the tombstone is not yet safe to remove — redeploy these Workers
	// with bindings re-pointed first.
	ReferencingScripts []string                                                               `json:"referencing_scripts"`
	JSON               scriptScriptAndVersionSettingEditResponseExportsReconciliationInfoJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseExportsReconciliationInfoJSON contains
// the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfo]
type scriptScriptAndVersionSettingEditResponseExportsReconciliationInfoJSON struct {
	Class              apijson.Field
	Message            apijson.Field
	Scenario           apijson.Field
	NamespaceID        apijson.Field
	ReferencingScripts apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseExportsReconciliationInfoJSON) RawJSON() string {
	return r.raw
}

// Stable, machine-readable tag identifying which reconciliation scenario produced
// an error, warning, or info entry. Clients may branch on this value instead of
// parsing `message`.
type ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario string

const (
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioCodeClassNotInExports                     ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "code_class_not_in_exports"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioProvisionedClassMissingFromConfig         ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "provisioned_class_missing_from_config"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioConfigExportNotInCode                     ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "config_export_not_in_code"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioConfigReferencesNonexistentClass          ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "config_references_nonexistent_class"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioOrphanedProvisionedNamespace              ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "orphaned_provisioned_namespace"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioStorageTypeMismatch                       ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "storage_type_mismatch"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioFreeTierRequiresSqlite                    ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "free_tier_requires_sqlite"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioInvalidExport                             ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "invalid_export"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTombstoneDeleteClassStillInCode           ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "tombstone_delete_class_still_in_code"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTombstoneDeleteBlockedByExternalBindings  ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "tombstone_delete_blocked_by_external_bindings"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTombstoneRenamedToOccupied                ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "tombstone_renamed_to_occupied"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTransferredPendingNotFound                ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "transferred_pending_not_found"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTransferredTargetMissing                  ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "transferred_target_missing"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTransferredTargetMismatch                 ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "transferred_target_mismatch"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioPhaseOneTransferSourceMissing             ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "phase_one_transfer_source_missing"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioPhaseOneTransferSourceNamespaceMissing    ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "phase_one_transfer_source_namespace_missing"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioPhaseOneTransferTargetClassProvisioned    ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "phase_one_transfer_target_class_provisioned"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioPhaseOneTransferAfterCommitMismatch       ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "phase_one_transfer_after_commit_mismatch"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioPhaseOneTransferDuplicate                 ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "phase_one_transfer_duplicate"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioPhaseOneTransferTargetInDispatchNamespace ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "phase_one_transfer_target_in_dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioPhaseOneTransferSourceInDispatchNamespace ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "phase_one_transfer_source_in_dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTransferredSourceInDispatchNamespace      ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "transferred_source_in_dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTransferredTargetInDispatchNamespace      ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "transferred_target_in_dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioContainerUndeclaredReference              ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "container_undeclared_reference"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioContainerClassNotDurableObject            ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "container_class_not_durable_object"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioContainerWiringInconsistent               ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "container_wiring_inconsistent"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioContainerMultipleDurableObjects           ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "container_multiple_durable_objects"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTransferContainerParityMismatch           ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "transfer_container_parity_mismatch"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTransferContainerParityMismatchOnCommit   ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "transfer_container_parity_mismatch_on_commit"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTombstoneClassStillInCode                 ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "tombstone_class_still_in_code"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioStaleTombstone                            ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "stale_tombstone"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTransferReceiveAlreadyApplied             ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "transfer_receive_already_applied"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTransferReceiveCleanupComplete            ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario = "transfer_receive_cleanup_complete"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenario) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioCodeClassNotInExports, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioProvisionedClassMissingFromConfig, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioConfigExportNotInCode, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioConfigReferencesNonexistentClass, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioOrphanedProvisionedNamespace, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioStorageTypeMismatch, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioFreeTierRequiresSqlite, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioInvalidExport, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTombstoneDeleteClassStillInCode, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTombstoneDeleteBlockedByExternalBindings, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTombstoneRenamedToOccupied, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTransferredPendingNotFound, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTransferredTargetMissing, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTransferredTargetMismatch, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioPhaseOneTransferSourceMissing, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioPhaseOneTransferSourceNamespaceMissing, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioPhaseOneTransferTargetClassProvisioned, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioPhaseOneTransferAfterCommitMismatch, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioPhaseOneTransferDuplicate, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioPhaseOneTransferTargetInDispatchNamespace, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioPhaseOneTransferSourceInDispatchNamespace, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTransferredSourceInDispatchNamespace, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTransferredTargetInDispatchNamespace, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioContainerUndeclaredReference, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioContainerClassNotDurableObject, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioContainerWiringInconsistent, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioContainerMultipleDurableObjects, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTransferContainerParityMismatch, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTransferContainerParityMismatchOnCommit, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTombstoneClassStillInCode, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioStaleTombstone, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTransferReceiveAlreadyApplied, ScriptScriptAndVersionSettingEditResponseExportsReconciliationInfoScenarioTransferReceiveCleanupComplete:
		return true
	}
	return false
}

// A single applied `renamed` tombstone.
type ScriptScriptAndVersionSettingEditResponseExportsReconciliationRenamed struct {
	// The original (source) class name.
	From string `json:"from" api:"required"`
	// The new class name (`renamed_to`).
	To   string                                                                    `json:"to" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseExportsReconciliationRenamedJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseExportsReconciliationRenamedJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseExportsReconciliationRenamed]
type scriptScriptAndVersionSettingEditResponseExportsReconciliationRenamedJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseExportsReconciliationRenamed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseExportsReconciliationRenamedJSON) RawJSON() string {
	return r.raw
}

// A single phase-1 transfer hint recorded on the target side (a live
// `expecting-transfer` entry).
type ScriptScriptAndVersionSettingEditResponseExportsReconciliationTransferPending struct {
	// The target-side class name awaiting transfer.
	Class string `json:"class" api:"required"`
	// The source script the namespace will be transferred from.
	From string                                                                            `json:"from" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseExportsReconciliationTransferPendingJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseExportsReconciliationTransferPendingJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseExportsReconciliationTransferPending]
type scriptScriptAndVersionSettingEditResponseExportsReconciliationTransferPendingJSON struct {
	Class       apijson.Field
	From        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseExportsReconciliationTransferPending) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseExportsReconciliationTransferPendingJSON) RawJSON() string {
	return r.raw
}

// A single committed `transferred` tombstone (phase-2 commit).
type ScriptScriptAndVersionSettingEditResponseExportsReconciliationTransferred struct {
	// The source class name that was transferred.
	Class string `json:"class" api:"required"`
	// The transfer phase. Currently always `committed`.
	Phase ScriptScriptAndVersionSettingEditResponseExportsReconciliationTransferredPhase `json:"phase" api:"required"`
	// The destination script that now owns the namespace.
	To   string                                                                        `json:"to" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponseExportsReconciliationTransferredJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseExportsReconciliationTransferredJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseExportsReconciliationTransferred]
type scriptScriptAndVersionSettingEditResponseExportsReconciliationTransferredJSON struct {
	Class       apijson.Field
	Phase       apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseExportsReconciliationTransferred) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseExportsReconciliationTransferredJSON) RawJSON() string {
	return r.raw
}

// The transfer phase. Currently always `committed`.
type ScriptScriptAndVersionSettingEditResponseExportsReconciliationTransferredPhase string

const (
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationTransferredPhaseCommitted ScriptScriptAndVersionSettingEditResponseExportsReconciliationTransferredPhase = "committed"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsReconciliationTransferredPhase) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsReconciliationTransferredPhaseCommitted:
		return true
	}
	return false
}

// A non-blocking reconciliation warning. Reserved: no scenario populates this
// array today (`code_class_not_in_exports` is surfaced as info and
// `provisioned_class_missing_from_config` is a hard error). Clients should still
// surface any entries that appear.
type ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarning struct {
	// The class name the warning is about.
	Class string `json:"class" api:"required"`
	// Human-readable explanation of the warning.
	Message string `json:"message" api:"required"`
	// Stable, machine-readable tag identifying which reconciliation scenario produced
	// an error, warning, or info entry. Clients may branch on this value instead of
	// parsing `message`.
	Scenario ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario `json:"scenario" api:"required"`
	// The provisioned namespace the warning relates to, when applicable.
	NamespaceID string                                                                    `json:"namespace_id" format:"uuid"`
	JSON        scriptScriptAndVersionSettingEditResponseExportsReconciliationWarningJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseExportsReconciliationWarningJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarning]
type scriptScriptAndVersionSettingEditResponseExportsReconciliationWarningJSON struct {
	Class       apijson.Field
	Message     apijson.Field
	Scenario    apijson.Field
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseExportsReconciliationWarningJSON) RawJSON() string {
	return r.raw
}

// Stable, machine-readable tag identifying which reconciliation scenario produced
// an error, warning, or info entry. Clients may branch on this value instead of
// parsing `message`.
type ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario string

const (
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioCodeClassNotInExports                     ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "code_class_not_in_exports"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioProvisionedClassMissingFromConfig         ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "provisioned_class_missing_from_config"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioConfigExportNotInCode                     ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "config_export_not_in_code"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioConfigReferencesNonexistentClass          ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "config_references_nonexistent_class"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioOrphanedProvisionedNamespace              ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "orphaned_provisioned_namespace"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioStorageTypeMismatch                       ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "storage_type_mismatch"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioFreeTierRequiresSqlite                    ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "free_tier_requires_sqlite"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioInvalidExport                             ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "invalid_export"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTombstoneDeleteClassStillInCode           ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "tombstone_delete_class_still_in_code"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTombstoneDeleteBlockedByExternalBindings  ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "tombstone_delete_blocked_by_external_bindings"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTombstoneRenamedToOccupied                ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "tombstone_renamed_to_occupied"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTransferredPendingNotFound                ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "transferred_pending_not_found"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTransferredTargetMissing                  ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "transferred_target_missing"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTransferredTargetMismatch                 ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "transferred_target_mismatch"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioPhaseOneTransferSourceMissing             ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "phase_one_transfer_source_missing"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioPhaseOneTransferSourceNamespaceMissing    ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "phase_one_transfer_source_namespace_missing"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioPhaseOneTransferTargetClassProvisioned    ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "phase_one_transfer_target_class_provisioned"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioPhaseOneTransferAfterCommitMismatch       ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "phase_one_transfer_after_commit_mismatch"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioPhaseOneTransferDuplicate                 ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "phase_one_transfer_duplicate"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioPhaseOneTransferTargetInDispatchNamespace ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "phase_one_transfer_target_in_dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioPhaseOneTransferSourceInDispatchNamespace ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "phase_one_transfer_source_in_dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTransferredSourceInDispatchNamespace      ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "transferred_source_in_dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTransferredTargetInDispatchNamespace      ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "transferred_target_in_dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioContainerUndeclaredReference              ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "container_undeclared_reference"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioContainerClassNotDurableObject            ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "container_class_not_durable_object"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioContainerWiringInconsistent               ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "container_wiring_inconsistent"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioContainerMultipleDurableObjects           ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "container_multiple_durable_objects"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTransferContainerParityMismatch           ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "transfer_container_parity_mismatch"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTransferContainerParityMismatchOnCommit   ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "transfer_container_parity_mismatch_on_commit"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTombstoneClassStillInCode                 ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "tombstone_class_still_in_code"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioStaleTombstone                            ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "stale_tombstone"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTransferReceiveAlreadyApplied             ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "transfer_receive_already_applied"
	ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTransferReceiveCleanupComplete            ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario = "transfer_receive_cleanup_complete"
)

func (r ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenario) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioCodeClassNotInExports, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioProvisionedClassMissingFromConfig, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioConfigExportNotInCode, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioConfigReferencesNonexistentClass, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioOrphanedProvisionedNamespace, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioStorageTypeMismatch, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioFreeTierRequiresSqlite, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioInvalidExport, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTombstoneDeleteClassStillInCode, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTombstoneDeleteBlockedByExternalBindings, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTombstoneRenamedToOccupied, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTransferredPendingNotFound, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTransferredTargetMissing, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTransferredTargetMismatch, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioPhaseOneTransferSourceMissing, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioPhaseOneTransferSourceNamespaceMissing, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioPhaseOneTransferTargetClassProvisioned, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioPhaseOneTransferAfterCommitMismatch, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioPhaseOneTransferDuplicate, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioPhaseOneTransferTargetInDispatchNamespace, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioPhaseOneTransferSourceInDispatchNamespace, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTransferredSourceInDispatchNamespace, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTransferredTargetInDispatchNamespace, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioContainerUndeclaredReference, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioContainerClassNotDurableObject, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioContainerWiringInconsistent, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioContainerMultipleDurableObjects, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTransferContainerParityMismatch, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTransferContainerParityMismatchOnCommit, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTombstoneClassStillInCode, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioStaleTombstone, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTransferReceiveAlreadyApplied, ScriptScriptAndVersionSettingEditResponseExportsReconciliationWarningsScenarioTransferReceiveCleanupComplete:
		return true
	}
	return false
}

// Limits to apply for this Worker.
type ScriptScriptAndVersionSettingEditResponseLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
	CPUMs int64 `json:"cpu_ms"`
	// The number of subrequests this Worker can make per request.
	Subrequests int64                                               `json:"subrequests"`
	JSON        scriptScriptAndVersionSettingEditResponseLimitsJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseLimitsJSON contains the JSON metadata
// for the struct [ScriptScriptAndVersionSettingEditResponseLimits]
type scriptScriptAndVersionSettingEditResponseLimitsJSON struct {
	CPUMs       apijson.Field
	Subrequests apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseLimitsJSON) RawJSON() string {
	return r.raw
}

// Migrations to apply for Durable Objects associated with this Worker.
type ScriptScriptAndVersionSettingEditResponseMigrations struct {
	// This field can have the runtime type of [[]string].
	DeletedClasses interface{} `json:"deleted_classes"`
	// This field can have the runtime type of [[]string].
	NewClasses interface{} `json:"new_classes"`
	// This field can have the runtime type of [[]string].
	NewSqliteClasses interface{} `json:"new_sqlite_classes"`
	// This field can have the runtime type of [[]SingleStepMigrationRenamedClass].
	RenamedClasses interface{} `json:"renamed_classes"`
	// This field can have the runtime type of [[]MigrationStep].
	Steps interface{} `json:"steps"`
	// This field can have the runtime type of [[]SingleStepMigrationTransferredClass].
	TransferredClasses interface{}                                             `json:"transferred_classes"`
	JSON               scriptScriptAndVersionSettingEditResponseMigrationsJSON `json:"-"`
	union              ScriptScriptAndVersionSettingEditResponseMigrationsUnion
}

// scriptScriptAndVersionSettingEditResponseMigrationsJSON contains the JSON
// metadata for the struct [ScriptScriptAndVersionSettingEditResponseMigrations]
type scriptScriptAndVersionSettingEditResponseMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewSqliteClasses   apijson.Field
	RenamedClasses     apijson.Field
	Steps              apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r scriptScriptAndVersionSettingEditResponseMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptScriptAndVersionSettingEditResponseMigrations) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptScriptAndVersionSettingEditResponseMigrations{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptScriptAndVersionSettingEditResponseMigrationsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SingleStepMigration],
// [ScriptScriptAndVersionSettingEditResponseMigrationsWorkersMultipleStepMigrations].
func (r ScriptScriptAndVersionSettingEditResponseMigrations) AsUnion() ScriptScriptAndVersionSettingEditResponseMigrationsUnion {
	return r.union
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by [SingleStepMigration] or
// [ScriptScriptAndVersionSettingEditResponseMigrationsWorkersMultipleStepMigrations].
type ScriptScriptAndVersionSettingEditResponseMigrationsUnion interface {
	implementsScriptScriptAndVersionSettingEditResponseMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptScriptAndVersionSettingEditResponseMigrationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SingleStepMigration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseMigrationsWorkersMultipleStepMigrations{}),
		},
	)
}

type ScriptScriptAndVersionSettingEditResponseMigrationsWorkersMultipleStepMigrations struct {
	JSON scriptScriptAndVersionSettingEditResponseMigrationsWorkersMultipleStepMigrationsJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseMigrationsWorkersMultipleStepMigrationsJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseMigrationsWorkersMultipleStepMigrations]
type scriptScriptAndVersionSettingEditResponseMigrationsWorkersMultipleStepMigrationsJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseMigrationsWorkersMultipleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseMigrationsWorkersMultipleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseMigrationsWorkersMultipleStepMigrations) implementsScriptScriptAndVersionSettingEditResponseMigrations() {
}

// Observability settings for the Worker.
type ScriptScriptAndVersionSettingEditResponseObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled bool `json:"enabled" api:"required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate float64 `json:"head_sampling_rate" api:"nullable"`
	// Log settings for the Worker.
	Logs ScriptScriptAndVersionSettingEditResponseObservabilityLogs `json:"logs" api:"nullable"`
	// Whether query strings are removed from request URLs in logs and traces.
	RedactQueryString bool `json:"redact_query_string"`
	// Trace settings for the Worker.
	Traces ScriptScriptAndVersionSettingEditResponseObservabilityTraces `json:"traces" api:"nullable"`
	JSON   scriptScriptAndVersionSettingEditResponseObservabilityJSON   `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseObservabilityJSON contains the JSON
// metadata for the struct [ScriptScriptAndVersionSettingEditResponseObservability]
type scriptScriptAndVersionSettingEditResponseObservabilityJSON struct {
	Enabled           apijson.Field
	HeadSamplingRate  apijson.Field
	Logs              apijson.Field
	RedactQueryString apijson.Field
	Traces            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseObservabilityJSON) RawJSON() string {
	return r.raw
}

// Log settings for the Worker.
type ScriptScriptAndVersionSettingEditResponseObservabilityLogs struct {
	// Whether logs are enabled for the Worker.
	Enabled bool `json:"enabled" api:"required"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs bool `json:"invocation_logs" api:"required"`
	// A list of destinations where logs will be exported to.
	Destinations []string `json:"destinations"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate float64 `json:"head_sampling_rate" api:"nullable"`
	// Whether log persistence is enabled for the Worker.
	Persist bool                                                           `json:"persist"`
	JSON    scriptScriptAndVersionSettingEditResponseObservabilityLogsJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseObservabilityLogsJSON contains the JSON
// metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseObservabilityLogs]
type scriptScriptAndVersionSettingEditResponseObservabilityLogsJSON struct {
	Enabled          apijson.Field
	InvocationLogs   apijson.Field
	Destinations     apijson.Field
	HeadSamplingRate apijson.Field
	Persist          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// Trace settings for the Worker.
type ScriptScriptAndVersionSettingEditResponseObservabilityTraces struct {
	// A list of destinations where traces will be exported to.
	Destinations []string `json:"destinations"`
	// Whether traces are enabled for the Worker.
	Enabled bool `json:"enabled"`
	// The sampling rate for traces. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate float64 `json:"head_sampling_rate" api:"nullable"`
	// Whether trace persistence is enabled for the Worker.
	Persist bool `json:"persist"`
	// Controls how inbound trace context (traceparent/tracestate) headers on incoming
	// requests are handled. "authenticated" honors inbound trace context only when
	// accompanied by a valid trace auth token. "accept" unconditionally accepts
	// inbound trace context. Requires the trace propagation feature to be enabled.
	// Returns null when the trace propagation feature is not enabled for the account.
	PropagationPolicy ScriptScriptAndVersionSettingEditResponseObservabilityTracesPropagationPolicy `json:"propagation_policy" api:"nullable"`
	JSON              scriptScriptAndVersionSettingEditResponseObservabilityTracesJSON              `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseObservabilityTracesJSON contains the
// JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseObservabilityTraces]
type scriptScriptAndVersionSettingEditResponseObservabilityTracesJSON struct {
	Destinations      apijson.Field
	Enabled           apijson.Field
	HeadSamplingRate  apijson.Field
	Persist           apijson.Field
	PropagationPolicy apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseObservabilityTraces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseObservabilityTracesJSON) RawJSON() string {
	return r.raw
}

// Controls how inbound trace context (traceparent/tracestate) headers on incoming
// requests are handled. "authenticated" honors inbound trace context only when
// accompanied by a valid trace auth token. "accept" unconditionally accepts
// inbound trace context. Requires the trace propagation feature to be enabled.
// Returns null when the trace propagation feature is not enabled for the account.
type ScriptScriptAndVersionSettingEditResponseObservabilityTracesPropagationPolicy string

const (
	ScriptScriptAndVersionSettingEditResponseObservabilityTracesPropagationPolicyAuthenticated ScriptScriptAndVersionSettingEditResponseObservabilityTracesPropagationPolicy = "authenticated"
	ScriptScriptAndVersionSettingEditResponseObservabilityTracesPropagationPolicyAccept        ScriptScriptAndVersionSettingEditResponseObservabilityTracesPropagationPolicy = "accept"
)

func (r ScriptScriptAndVersionSettingEditResponseObservabilityTracesPropagationPolicy) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseObservabilityTracesPropagationPolicyAuthenticated, ScriptScriptAndVersionSettingEditResponseObservabilityTracesPropagationPolicyAccept:
		return true
	}
	return false
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
type ScriptScriptAndVersionSettingEditResponsePlacement struct {
	// TCP host and port for targeted placement.
	Host string `json:"host"`
	// HTTP hostname for targeted placement.
	Hostname string `json:"hostname"`
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode ScriptScriptAndVersionSettingEditResponsePlacementMode `json:"mode"`
	// Cloud region for targeted placement in format 'provider:region'.
	Region string `json:"region"`
	// This field can have the runtime type of
	// [[]ScriptScriptAndVersionSettingEditResponsePlacementObjectTarget].
	Target interface{}                                            `json:"target"`
	JSON   scriptScriptAndVersionSettingEditResponsePlacementJSON `json:"-"`
	union  ScriptScriptAndVersionSettingEditResponsePlacementUnion
}

// scriptScriptAndVersionSettingEditResponsePlacementJSON contains the JSON
// metadata for the struct [ScriptScriptAndVersionSettingEditResponsePlacement]
type scriptScriptAndVersionSettingEditResponsePlacementJSON struct {
	Host        apijson.Field
	Hostname    apijson.Field
	Mode        apijson.Field
	Region      apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r scriptScriptAndVersionSettingEditResponsePlacementJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptScriptAndVersionSettingEditResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptScriptAndVersionSettingEditResponsePlacement{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptScriptAndVersionSettingEditResponsePlacementUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ScriptScriptAndVersionSettingEditResponsePlacementMode],
// [ScriptScriptAndVersionSettingEditResponsePlacementRegion],
// [ScriptScriptAndVersionSettingEditResponsePlacementHostname],
// [ScriptScriptAndVersionSettingEditResponsePlacementHost],
// [ScriptScriptAndVersionSettingEditResponsePlacementObject],
// [ScriptScriptAndVersionSettingEditResponsePlacementObject],
// [ScriptScriptAndVersionSettingEditResponsePlacementObject],
// [ScriptScriptAndVersionSettingEditResponsePlacementObject].
func (r ScriptScriptAndVersionSettingEditResponsePlacement) AsUnion() ScriptScriptAndVersionSettingEditResponsePlacementUnion {
	return r.union
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
//
// Union satisfied by [ScriptScriptAndVersionSettingEditResponsePlacementMode],
// [ScriptScriptAndVersionSettingEditResponsePlacementRegion],
// [ScriptScriptAndVersionSettingEditResponsePlacementHostname],
// [ScriptScriptAndVersionSettingEditResponsePlacementHost],
// [ScriptScriptAndVersionSettingEditResponsePlacementObject],
// [ScriptScriptAndVersionSettingEditResponsePlacementObject],
// [ScriptScriptAndVersionSettingEditResponsePlacementObject] or
// [ScriptScriptAndVersionSettingEditResponsePlacementObject].
type ScriptScriptAndVersionSettingEditResponsePlacementUnion interface {
	implementsScriptScriptAndVersionSettingEditResponsePlacement()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptScriptAndVersionSettingEditResponsePlacementUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptScriptAndVersionSettingEditResponsePlacementMode{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptScriptAndVersionSettingEditResponsePlacementRegion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptScriptAndVersionSettingEditResponsePlacementHostname{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptScriptAndVersionSettingEditResponsePlacementHost{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptScriptAndVersionSettingEditResponsePlacementObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptScriptAndVersionSettingEditResponsePlacementObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptScriptAndVersionSettingEditResponsePlacementObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptScriptAndVersionSettingEditResponsePlacementObject{}),
		},
	)
}

type ScriptScriptAndVersionSettingEditResponsePlacementMode struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode ScriptScriptAndVersionSettingEditResponsePlacementModeMode `json:"mode" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponsePlacementModeJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponsePlacementModeJSON contains the JSON
// metadata for the struct [ScriptScriptAndVersionSettingEditResponsePlacementMode]
type scriptScriptAndVersionSettingEditResponsePlacementModeJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponsePlacementMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponsePlacementModeJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponsePlacementMode) implementsScriptScriptAndVersionSettingEditResponsePlacement() {
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptScriptAndVersionSettingEditResponsePlacementModeMode string

const (
	ScriptScriptAndVersionSettingEditResponsePlacementModeModeSmart ScriptScriptAndVersionSettingEditResponsePlacementModeMode = "smart"
)

func (r ScriptScriptAndVersionSettingEditResponsePlacementModeMode) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponsePlacementModeModeSmart:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponsePlacementRegion struct {
	// Cloud region for targeted placement in format 'provider:region'.
	Region string                                                       `json:"region" api:"required"`
	JSON   scriptScriptAndVersionSettingEditResponsePlacementRegionJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponsePlacementRegionJSON contains the JSON
// metadata for the struct
// [ScriptScriptAndVersionSettingEditResponsePlacementRegion]
type scriptScriptAndVersionSettingEditResponsePlacementRegionJSON struct {
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponsePlacementRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponsePlacementRegionJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponsePlacementRegion) implementsScriptScriptAndVersionSettingEditResponsePlacement() {
}

type ScriptScriptAndVersionSettingEditResponsePlacementHostname struct {
	// HTTP hostname for targeted placement.
	Hostname string                                                         `json:"hostname" api:"required"`
	JSON     scriptScriptAndVersionSettingEditResponsePlacementHostnameJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponsePlacementHostnameJSON contains the JSON
// metadata for the struct
// [ScriptScriptAndVersionSettingEditResponsePlacementHostname]
type scriptScriptAndVersionSettingEditResponsePlacementHostnameJSON struct {
	Hostname    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponsePlacementHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponsePlacementHostnameJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponsePlacementHostname) implementsScriptScriptAndVersionSettingEditResponsePlacement() {
}

type ScriptScriptAndVersionSettingEditResponsePlacementHost struct {
	// TCP host and port for targeted placement.
	Host string                                                     `json:"host" api:"required"`
	JSON scriptScriptAndVersionSettingEditResponsePlacementHostJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponsePlacementHostJSON contains the JSON
// metadata for the struct [ScriptScriptAndVersionSettingEditResponsePlacementHost]
type scriptScriptAndVersionSettingEditResponsePlacementHostJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponsePlacementHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponsePlacementHostJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponsePlacementHost) implementsScriptScriptAndVersionSettingEditResponsePlacement() {
}

type ScriptScriptAndVersionSettingEditResponsePlacementObject struct {
	// Targeted placement mode.
	Mode ScriptScriptAndVersionSettingEditResponsePlacementObjectMode `json:"mode" api:"required"`
	// Cloud region for targeted placement in format 'provider:region'.
	Region string                                                       `json:"region" api:"required"`
	JSON   scriptScriptAndVersionSettingEditResponsePlacementObjectJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponsePlacementObjectJSON contains the JSON
// metadata for the struct
// [ScriptScriptAndVersionSettingEditResponsePlacementObject]
type scriptScriptAndVersionSettingEditResponsePlacementObjectJSON struct {
	Mode        apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponsePlacementObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponsePlacementObjectJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponsePlacementObject) implementsScriptScriptAndVersionSettingEditResponsePlacement() {
}

// Targeted placement mode.
type ScriptScriptAndVersionSettingEditResponsePlacementObjectMode string

const (
	ScriptScriptAndVersionSettingEditResponsePlacementObjectModeTargeted ScriptScriptAndVersionSettingEditResponsePlacementObjectMode = "targeted"
)

func (r ScriptScriptAndVersionSettingEditResponsePlacementObjectMode) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponsePlacementObjectModeTargeted:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type ScriptScriptAndVersionSettingEditResponseUsageModel string

const (
	ScriptScriptAndVersionSettingEditResponseUsageModelStandard ScriptScriptAndVersionSettingEditResponseUsageModel = "standard"
	ScriptScriptAndVersionSettingEditResponseUsageModelBundled  ScriptScriptAndVersionSettingEditResponseUsageModel = "bundled"
	ScriptScriptAndVersionSettingEditResponseUsageModelUnbound  ScriptScriptAndVersionSettingEditResponseUsageModel = "unbound"
)

func (r ScriptScriptAndVersionSettingEditResponseUsageModel) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseUsageModelStandard, ScriptScriptAndVersionSettingEditResponseUsageModelBundled, ScriptScriptAndVersionSettingEditResponseUsageModelUnbound:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponse struct {
	// Annotations for the Worker version. Annotations are not inherited across
	// settings updates; omitting this field means the new version will have no
	// annotations.
	Annotations ScriptScriptAndVersionSettingGetResponseAnnotations `json:"annotations"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings []ScriptScriptAndVersionSettingGetResponseBinding `json:"bindings"`
	// Global CacheW configuration for the Worker. When caching is on, the platform
	// provisions a `cloudflare.app` zone for the Worker. A `type: worker` entry in the
	// `exports` map can override this value for a single entrypoint.
	CacheOptions ScriptScriptAndVersionSettingGetResponseCacheOptions `json:"cache_options"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate string `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Summary of the declarative exports reconciliation that ran on this upload.
	// Populated only when the uploaded metadata included an `exports` block. Durable
	// Object entries drive reconciliation; `type: worker` entries do not contribute to
	// this summary.
	ExportsReconciliation ScriptScriptAndVersionSettingGetResponseExportsReconciliation `json:"exports_reconciliation"`
	// Limits to apply for this Worker.
	Limits ScriptScriptAndVersionSettingGetResponseLimits `json:"limits"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Observability settings for the Worker.
	Observability ScriptScriptAndVersionSettingGetResponseObservability `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
	Placement ScriptScriptAndVersionSettingGetResponsePlacement `json:"placement"`
	// Tags associated with the Worker.
	Tags []string `json:"tags" api:"nullable"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ConsumerScript `json:"tail_consumers" api:"nullable"`
	// Usage model for the Worker invocations.
	UsageModel ScriptScriptAndVersionSettingGetResponseUsageModel `json:"usage_model"`
	JSON       scriptScriptAndVersionSettingGetResponseJSON       `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseJSON contains the JSON metadata for the
// struct [ScriptScriptAndVersionSettingGetResponse]
type scriptScriptAndVersionSettingGetResponseJSON struct {
	Annotations           apijson.Field
	Bindings              apijson.Field
	CacheOptions          apijson.Field
	CompatibilityDate     apijson.Field
	CompatibilityFlags    apijson.Field
	ExportsReconciliation apijson.Field
	Limits                apijson.Field
	Logpush               apijson.Field
	Observability         apijson.Field
	Placement             apijson.Field
	Tags                  apijson.Field
	TailConsumers         apijson.Field
	UsageModel            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

// Annotations for the Worker version. Annotations are not inherited across
// settings updates; omitting this field means the new version will have no
// annotations.
type ScriptScriptAndVersionSettingGetResponseAnnotations struct {
	// Human-readable message about the version. Truncated to 1000 bytes if longer.
	WorkersMessage string `json:"workers/message"`
	// User-provided identifier for the version. Maximum 100 bytes.
	WorkersTag string `json:"workers/tag"`
	// Operation that triggered the creation of the version. This is read-only and set
	// by the server.
	WorkersTriggeredBy string                                                  `json:"workers/triggered_by"`
	JSON               scriptScriptAndVersionSettingGetResponseAnnotationsJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseAnnotationsJSON contains the JSON
// metadata for the struct [ScriptScriptAndVersionSettingGetResponseAnnotations]
type scriptScriptAndVersionSettingGetResponseAnnotationsJSON struct {
	WorkersMessage     apijson.Field
	WorkersTag         apijson.Field
	WorkersTriggeredBy apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseAnnotations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseAnnotationsJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources.
type ScriptScriptAndVersionSettingGetResponseBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsType `json:"type" api:"required"`
	// Identifier of the D1 database to bind to.
	//
	// Deprecated: This property has been renamed to `database_id`.
	ID string `json:"id"`
	// This field can have the runtime type of [interface{}].
	Algorithm interface{} `json:"algorithm"`
	// This field can have the runtime type of [[]string].
	AllowedDestinationAddresses interface{} `json:"allowed_destination_addresses"`
	// This field can have the runtime type of [[]string].
	AllowedSenderAddresses interface{} `json:"allowed_sender_addresses"`
	// ID of the Flagship app to bind to for feature flag evaluation.
	AppID string `json:"app_id"`
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name"`
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// Identifier of the D1 database to bind to.
	DatabaseID string `json:"database_id"`
	// The name of the dataset to bind to.
	Dataset string `json:"dataset"`
	// Destination address for the email.
	DestinationAddress string `json:"destination_address" format:"email"`
	// The dispatch namespace the Durable Object script belongs to.
	DispatchNamespace string `json:"dispatch_namespace"`
	// Entrypoint to invoke on the target Worker.
	Entrypoint string `json:"entrypoint"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptScriptAndVersionSettingGetResponseBindingsFormat `json:"format"`
	// Enables Gateway identity for the binding. Requires network_id to be
	// "cf1:network" and cannot be combined with tunnel_id.
	Identity ScriptScriptAndVersionSettingGetResponseBindingsIdentity `json:"identity"`
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name"`
	// The user-chosen instance name. Must exist at deploy time. The worker can search,
	// chat, update, and manage items/jobs on this instance.
	InstanceName string `json:"instance_name"`
	// This field can have the runtime type of [interface{}].
	Json interface{} `json:"json"`
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction ScriptScriptAndVersionSettingGetResponseBindingsJurisdiction `json:"jurisdiction"`
	// This field can have the runtime type of [interface{}].
	KeyJwk interface{} `json:"key_jwk"`
	// The namespace the instance belongs to. Defaults to "default" if omitted.
	// Customers who don't use namespaces can simply omit this field.
	Namespace string `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// Identifier of the network to bind to. Only "cf1:network" is currently supported.
	// Mutually exclusive with tunnel_id.
	NetworkID string `json:"network_id"`
	// The old name of the inherited binding. If set, the binding will be renamed from
	// `old_name` to `name` in the new version. If not set, the binding will keep the
	// same name between versions.
	OldName string `json:"old_name"`
	// This field can have the runtime type of
	// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutbound].
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
	// Identifier of the VPC service to bind to.
	ServiceID string `json:"service_id"`
	// This field can have the runtime type of
	// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimitSimple].
	Simple interface{} `json:"simple"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id"`
	// The text value to use.
	Text string `json:"text"`
	// UUID of the Cloudflare Tunnel to bind to. Mutually exclusive with network_id.
	TunnelID string `json:"tunnel_id"`
	// This field can have the runtime type of
	// [[]ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage].
	Usages interface{} `json:"usages"`
	// Identifier for the version to inherit the binding from, which can be the version
	// ID or the literal "latest" to inherit from the latest version. Defaults to
	// inheriting the binding from the latest version.
	VersionID string `json:"version_id"`
	// Name of the Workflow to bind to.
	WorkflowName string                                              `json:"workflow_name"`
	JSON         scriptScriptAndVersionSettingGetResponseBindingJSON `json:"-"`
	union        ScriptScriptAndVersionSettingGetResponseBindingsUnion
}

// scriptScriptAndVersionSettingGetResponseBindingJSON contains the JSON metadata
// for the struct [ScriptScriptAndVersionSettingGetResponseBinding]
type scriptScriptAndVersionSettingGetResponseBindingJSON struct {
	Name                        apijson.Field
	Type                        apijson.Field
	ID                          apijson.Field
	Algorithm                   apijson.Field
	AllowedDestinationAddresses apijson.Field
	AllowedSenderAddresses      apijson.Field
	AppID                       apijson.Field
	BucketName                  apijson.Field
	CertificateID               apijson.Field
	ClassName                   apijson.Field
	DatabaseID                  apijson.Field
	Dataset                     apijson.Field
	DestinationAddress          apijson.Field
	DispatchNamespace           apijson.Field
	Entrypoint                  apijson.Field
	Environment                 apijson.Field
	Format                      apijson.Field
	Identity                    apijson.Field
	IndexName                   apijson.Field
	InstanceName                apijson.Field
	Json                        apijson.Field
	Jurisdiction                apijson.Field
	KeyJwk                      apijson.Field
	Namespace                   apijson.Field
	NamespaceID                 apijson.Field
	NetworkID                   apijson.Field
	OldName                     apijson.Field
	Outbound                    apijson.Field
	Part                        apijson.Field
	Pipeline                    apijson.Field
	QueueName                   apijson.Field
	ScriptName                  apijson.Field
	SecretName                  apijson.Field
	Service                     apijson.Field
	ServiceID                   apijson.Field
	Simple                      apijson.Field
	StoreID                     apijson.Field
	Text                        apijson.Field
	TunnelID                    apijson.Field
	Usages                      apijson.Field
	VersionID                   apijson.Field
	WorkflowName                apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r scriptScriptAndVersionSettingGetResponseBindingJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptScriptAndVersionSettingGetResponseBinding) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptScriptAndVersionSettingGetResponseBinding{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptScriptAndVersionSettingGetResponseBindingsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearch],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchNamespace],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMessaging],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowser],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDataBlob],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindInherit],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindImages],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMedia],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPipelines],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimit],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSendEmail],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTextBlob],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecret],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindFlagship],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKey],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWorkflow],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWasmModule],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCService],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetwork].
func (r ScriptScriptAndVersionSettingGetResponseBinding) AsUnion() ScriptScriptAndVersionSettingGetResponseBindingsUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources.
//
// Union satisfied by
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearch],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchNamespace],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMessaging],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowser],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDataBlob],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindInherit],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindImages],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMedia],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPipelines],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimit],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSendEmail],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTextBlob],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecret],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindFlagship],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKey],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWorkflow],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWasmModule],
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCService]
// or
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetwork].
type ScriptScriptAndVersionSettingGetResponseBindingsUnion interface {
	implementsScriptScriptAndVersionSettingGetResponseBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptScriptAndVersionSettingGetResponseBindingsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearch{}),
			DiscriminatorValue: "ai_search",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchNamespace{}),
			DiscriminatorValue: "ai_search_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMessaging{}),
			DiscriminatorValue: "messaging",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowser{}),
			DiscriminatorValue: "browser",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDataBlob{}),
			DiscriminatorValue: "data_blob",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindInherit{}),
			DiscriminatorValue: "inherit",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindImages{}),
			DiscriminatorValue: "images",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMedia{}),
			DiscriminatorValue: "media",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPipelines{}),
			DiscriminatorValue: "pipelines",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimit{}),
			DiscriminatorValue: "ratelimit",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSendEmail{}),
			DiscriminatorValue: "send_email",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTextBlob{}),
			DiscriminatorValue: "text_blob",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecret{}),
			DiscriminatorValue: "secrets_store_secret",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindFlagship{}),
			DiscriminatorValue: "flagship",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWorkflow{}),
			DiscriminatorValue: "workflow",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWasmModule{}),
			DiscriminatorValue: "wasm_module",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCService{}),
			DiscriminatorValue: "vpc_service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetwork{}),
			DiscriminatorValue: "vpc_network",
		},
	)
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeAI ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "ai"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearch struct {
	// The user-chosen instance name. Must exist at deploy time. The worker can search,
	// chat, update, and manage items/jobs on this instance.
	InstanceName string `json:"instance_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchType `json:"type" api:"required"`
	// The namespace the instance belongs to. Defaults to "default" if omitted.
	// Customers who don't use namespaces can simply omit this field.
	Namespace string                                                                         `json:"namespace"`
	JSON      scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearch]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchJSON struct {
	InstanceName apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	Namespace    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearch) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchTypeAISearch ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchType = "ai_search"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchTypeAISearch:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The user-chosen namespace name. Must exist before deploy -- Wrangler handles
	// auto-creation on deploy failure (R2 bucket pattern). The "default" namespace is
	// auto-created by config-api for new accounts. Grants full access (CRUD + search +
	// chat) to all instances within the namespace.
	Namespace string `json:"namespace" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchNamespaceType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchNamespaceJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchNamespace]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchNamespace) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchNamespaceType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchNamespaceTypeAISearchNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchNamespaceType = "ai_search_namespace"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAISearchNamespaceTypeAISearchNamespace:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMessaging struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The Messaging namespace to bind to.
	Namespace string `json:"namespace" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMessagingType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMessagingJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMessagingJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMessaging]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMessagingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMessaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMessagingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMessaging) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMessagingType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMessagingTypeMessaging ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMessagingType = "messaging"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMessagingType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMessagingTypeMessaging:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineJSON struct {
	Dataset     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeAssets ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "assets"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowser]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowser) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserTypeBrowser ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserType = "browser"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	DatabaseID string `json:"database_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type `json:"type" api:"required"`
	// Identifier of the D1 database to bind to.
	//
	// Deprecated: This property has been renamed to `database_id`.
	ID   string                                                                   `json:"id"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1JSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1JSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1JSON struct {
	DatabaseID  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1JSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeD1 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "d1"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDataBlob struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The name of the file containing the data content. Only accepted for
	// `service worker syntax` Workers.
	Part string `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDataBlobType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDataBlobJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDataBlobJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDataBlob]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDataBlobJSON struct {
	Name        apijson.Field
	Part        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDataBlob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDataBlobJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDataBlob) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDataBlobType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDataBlobTypeDataBlob ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDataBlobType = "data_blob"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDataBlobType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDataBlobTypeDataBlob:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The name of the dispatch namespace.
	Namespace string `json:"namespace" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType `json:"type" api:"required"`
	// Outbound worker.
	Outbound ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutbound `json:"outbound"`
	JSON     scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceJSON     `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params []ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundParam `json:"params"`
	// Outbound worker.
	Worker ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker `json:"worker"`
	JSON   scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON   `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutbound]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON) RawJSON() string {
	return r.raw
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundParam struct {
	// Name of the parameter.
	Name string                                                                                               `json:"name" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundParamJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundParamJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundParam]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundParamJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundParam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundParamJSON) RawJSON() string {
	return r.raw
}

// Outbound worker.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Entrypoint to invoke on the outbound worker.
	Entrypoint string `json:"entrypoint"`
	// Environment of the outbound worker.
	Environment string `json:"environment"`
	// Name of the outbound worker.
	Service string                                                                                                `json:"service"`
	JSON    scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON struct {
	Entrypoint  apijson.Field
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType `json:"type" api:"required"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
	// The dispatch namespace the Durable Object script belongs to.
	DispatchNamespace string `json:"dispatch_namespace"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string                                                                                       `json:"script_name"`
	JSON       scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON struct {
	Name              apijson.Field
	Type              apijson.Field
	ClassName         apijson.Field
	DispatchNamespace apijson.Field
	Environment       apijson.Field
	NamespaceID       apijson.Field
	ScriptName        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeHyperdrive ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindInherit struct {
	// The name of the inherited binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindInheritType `json:"type" api:"required"`
	// The old name of the inherited binding. If set, the binding will be renamed from
	// `old_name` to `name` in the new version. If not set, the binding will keep the
	// same name between versions.
	OldName string `json:"old_name"`
	// Identifier for the version to inherit the binding from, which can be the version
	// ID or the literal "latest" to inherit from the latest version. Defaults to
	// inheriting the binding from the latest version.
	VersionID string                                                                        `json:"version_id"`
	JSON      scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindInheritJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindInheritJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindInherit]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindInheritJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	OldName     apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindInherit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindInheritJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindInherit) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindInheritType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindInheritTypeInherit ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindInheritType = "inherit"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindInheritType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindInheritTypeInherit:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindImages struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindImagesType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindImagesJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindImagesJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindImages]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindImagesJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindImages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindImagesJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindImages) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindImagesType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindImagesTypeImages ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindImagesType = "images"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindImagesType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindImagesTypeImages:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json interface{} `json:"json" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonJSON struct {
	Json        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeJson ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "json"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeKVNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMedia struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMediaType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMediaJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMediaJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMedia]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMediaJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMedia) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMediaJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMedia) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMediaType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMediaTypeMedia ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMediaType = "media"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMediaType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMediaTypeMedia:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateJSON struct {
	CertificateID apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The text value to use.
	Text string `json:"text" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypePlainText ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Name of the Pipeline to bind to.
	Pipeline string `json:"pipeline" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPipelinesType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPipelinesJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPipelinesJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPipelines]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPipelinesJSON struct {
	Name        apijson.Field
	Pipeline    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPipelines) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPipelinesJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPipelines) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPipelinesType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPipelinesTypePipelines ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPipelinesType = "pipelines"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeQueue ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "queue"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimit struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Identifier of the rate limit namespace to bind to.
	NamespaceID string `json:"namespace_id" api:"required"`
	// The rate limit configuration.
	Simple ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimitSimple `json:"simple" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimitType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimitJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimitJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimit]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimitJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Simple      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimitJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimit) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The rate limit configuration.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimitSimple struct {
	// The limit (requests per period).
	Limit float64 `json:"limit" api:"required"`
	// The period in seconds.
	Period int64 `json:"period" api:"required"`
	// Duration in seconds to apply the mitigation action after the rate limit is
	// exceeded. Valid values are 0 (disabled), 10, or multiples of 60 up to 86400.
	// Must be greater than or equal to the period when non-zero.
	MitigationTimeout int64                                                                                 `json:"mitigation_timeout"`
	JSON              scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimitSimpleJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimitSimpleJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimitSimple]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimitSimpleJSON struct {
	Limit             apijson.Field
	Period            apijson.Field
	MitigationTimeout apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimitSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimitSimpleJSON) RawJSON() string {
	return r.raw
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimitType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimitTypeRatelimit ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimitType = "ratelimit"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimitType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindRatelimitTypeRatelimit:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType `json:"type" api:"required"`
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJurisdiction `json:"jurisdiction"`
	JSON         scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJSON         `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJSON struct {
	BucketName   apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	Jurisdiction apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeR2Bucket ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

// The
// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
// of the R2 bucket.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJurisdiction string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJurisdictionEu          ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJurisdiction = "eu"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJurisdictionFedramp     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJurisdiction = "fedramp"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJurisdictionFedrampHigh ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJurisdiction = "fedramp-high"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJurisdictionUs          ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJurisdiction = "us"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJurisdiction) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJurisdictionEu, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJurisdictionFedramp, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJurisdictionFedrampHigh, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJurisdictionUs:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeSecretText ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "secret_text"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSendEmail struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSendEmailType `json:"type" api:"required"`
	// List of allowed destination addresses.
	AllowedDestinationAddresses []string `json:"allowed_destination_addresses" format:"email"`
	// List of allowed sender addresses.
	AllowedSenderAddresses []string `json:"allowed_sender_addresses" format:"email"`
	// Destination address for the email.
	DestinationAddress string                                                                          `json:"destination_address" format:"email"`
	JSON               scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSendEmailJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSendEmailJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSendEmail]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSendEmailJSON struct {
	Name                        apijson.Field
	Type                        apijson.Field
	AllowedDestinationAddresses apijson.Field
	AllowedSenderAddresses      apijson.Field
	DestinationAddress          apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSendEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSendEmailJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSendEmail) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSendEmailType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSendEmailTypeSendEmail ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSendEmailType = "send_email"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSendEmailType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSendEmailTypeSendEmail:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Name of Worker to bind to.
	Service string `json:"service" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType `json:"type" api:"required"`
	// Entrypoint to invoke on the target Worker.
	Entrypoint string `json:"entrypoint"`
	// Optional environment if the Worker utilizes one.
	Environment string                                                                        `json:"environment"`
	JSON        scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceJSON struct {
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	Entrypoint  apijson.Field
	Environment apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeService ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "service"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTextBlob struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The name of the file containing the text content. Only accepted for
	// `service worker syntax` Workers.
	Part string `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTextBlobType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTextBlobJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTextBlobJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTextBlob]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTextBlobJSON struct {
	Name        apijson.Field
	Part        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTextBlob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTextBlobJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTextBlob) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTextBlobType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTextBlobTypeTextBlob ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTextBlobType = "text_blob"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTextBlobType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTextBlobTypeTextBlob:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeJSON struct {
	IndexName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeVectorize ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "vectorize"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Name of the secret in the store.
	SecretName string `json:"secret_name" api:"required"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecret]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretJSON struct {
	Name        apijson.Field
	SecretName  apijson.Field
	StoreID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecret) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindFlagship struct {
	// ID of the Flagship app to bind to for feature flag evaluation.
	AppID string `json:"app_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindFlagshipType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindFlagshipJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindFlagshipJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindFlagship]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindFlagshipJSON struct {
	AppID       apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindFlagship) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindFlagshipJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindFlagship) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindFlagshipType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindFlagshipTypeFlagship ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindFlagshipType = "flagship"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindFlagshipType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindFlagshipTypeFlagship:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm" api:"required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyFormat `json:"format" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyType `json:"type" api:"required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage `json:"usages" api:"required"`
	JSON   scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKey]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKey) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyFormat string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyFormatRaw   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyFormat = "raw"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyFormatPkcs8 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyFormat = "pkcs8"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyFormatSpki  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyFormat = "spki"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyFormatJwk   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyFormatRaw, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyFormatPkcs8, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyFormatSpki, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyTypeSecretKey ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyType = "secret_key"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageEncrypt    ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage = "encrypt"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageDecrypt    ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage = "decrypt"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageSign       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage = "sign"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageVerify     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage = "verify"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageDeriveKey  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage = "deriveKey"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageDeriveBits ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage = "deriveBits"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageWrapKey    ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage = "wrapKey"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageUnwrapKey  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageEncrypt, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageDecrypt, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageSign, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageVerify, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageDeriveKey, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageDeriveBits, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageWrapKey, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWorkflowType `json:"type" api:"required"`
	// Name of the Workflow to bind to.
	WorkflowName string `json:"workflow_name" api:"required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName string `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName string                                                                         `json:"script_name"`
	JSON       scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWorkflowJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWorkflowJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWorkflow]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWorkflowJSON struct {
	Name         apijson.Field
	Type         apijson.Field
	WorkflowName apijson.Field
	ClassName    apijson.Field
	ScriptName   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWorkflowJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWorkflow) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWorkflowType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWorkflowTypeWorkflow ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWorkflowType = "workflow"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWasmModule struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The name of the file containing the WebAssembly module content. Only accepted
	// for `service worker syntax` Workers.
	Part string `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWasmModuleType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWasmModuleJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWasmModuleJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWasmModule]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWasmModuleJSON struct {
	Name        apijson.Field
	Part        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWasmModule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWasmModuleJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWasmModule) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWasmModuleType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWasmModuleTypeWasmModule ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWasmModuleType = "wasm_module"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWasmModuleType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindWasmModuleTypeWasmModule:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCService struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Identifier of the VPC service to bind to.
	ServiceID string `json:"service_id" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCServiceType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCServiceJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCServiceJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCService]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCServiceJSON struct {
	Name        apijson.Field
	ServiceID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCServiceJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCService) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCServiceType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCServiceTypeVPCService ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCServiceType = "vpc_service"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCServiceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCServiceTypeVPCService:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetwork struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetworkType `json:"type" api:"required"`
	// Enables Gateway identity for the binding. Requires network_id to be
	// "cf1:network" and cannot be combined with tunnel_id.
	Identity ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetworkIdentity `json:"identity"`
	// Identifier of the network to bind to. Only "cf1:network" is currently supported.
	// Mutually exclusive with tunnel_id.
	NetworkID string `json:"network_id"`
	// UUID of the Cloudflare Tunnel to bind to. Mutually exclusive with network_id.
	TunnelID string                                                                           `json:"tunnel_id"`
	JSON     scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetworkJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetworkJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetwork]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetworkJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Identity    apijson.Field
	NetworkID   apijson.Field
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetworkJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetwork) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetworkType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetworkTypeVPCNetwork ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetworkType = "vpc_network"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetworkType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetworkTypeVPCNetwork:
		return true
	}
	return false
}

// Enables Gateway identity for the binding. Requires network_id to be
// "cf1:network" and cannot be combined with tunnel_id.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetworkIdentity string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetworkIdentityRuntimeEmailAlpha ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetworkIdentity = "runtime-email-alpha"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetworkIdentity) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVPCNetworkIdentityRuntimeEmailAlpha:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsTypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeAISearch               ScriptScriptAndVersionSettingGetResponseBindingsType = "ai_search"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeAISearchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsType = "ai_search_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeMessaging              ScriptScriptAndVersionSettingGetResponseBindingsType = "messaging"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeBrowser                ScriptScriptAndVersionSettingGetResponseBindingsType = "browser"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeDataBlob               ScriptScriptAndVersionSettingGetResponseBindingsType = "data_blob"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeInherit                ScriptScriptAndVersionSettingGetResponseBindingsType = "inherit"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeImages                 ScriptScriptAndVersionSettingGetResponseBindingsType = "images"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeMedia                  ScriptScriptAndVersionSettingGetResponseBindingsType = "media"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsTypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsTypePipelines              ScriptScriptAndVersionSettingGetResponseBindingsType = "pipelines"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeRatelimit              ScriptScriptAndVersionSettingGetResponseBindingsType = "ratelimit"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeSendEmail              ScriptScriptAndVersionSettingGetResponseBindingsType = "send_email"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeService                ScriptScriptAndVersionSettingGetResponseBindingsType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeTextBlob               ScriptScriptAndVersionSettingGetResponseBindingsType = "text_blob"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsType = "version_metadata"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeSecretsStoreSecret     ScriptScriptAndVersionSettingGetResponseBindingsType = "secrets_store_secret"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeFlagship               ScriptScriptAndVersionSettingGetResponseBindingsType = "flagship"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeSecretKey              ScriptScriptAndVersionSettingGetResponseBindingsType = "secret_key"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeWorkflow               ScriptScriptAndVersionSettingGetResponseBindingsType = "workflow"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeWasmModule             ScriptScriptAndVersionSettingGetResponseBindingsType = "wasm_module"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeVPCService             ScriptScriptAndVersionSettingGetResponseBindingsType = "vpc_service"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeVPCNetwork             ScriptScriptAndVersionSettingGetResponseBindingsType = "vpc_network"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsTypeAI, ScriptScriptAndVersionSettingGetResponseBindingsTypeAISearch, ScriptScriptAndVersionSettingGetResponseBindingsTypeAISearchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsTypeMessaging, ScriptScriptAndVersionSettingGetResponseBindingsTypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsTypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsTypeBrowser, ScriptScriptAndVersionSettingGetResponseBindingsTypeD1, ScriptScriptAndVersionSettingGetResponseBindingsTypeDataBlob, ScriptScriptAndVersionSettingGetResponseBindingsTypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsTypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsTypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsTypeInherit, ScriptScriptAndVersionSettingGetResponseBindingsTypeImages, ScriptScriptAndVersionSettingGetResponseBindingsTypeJson, ScriptScriptAndVersionSettingGetResponseBindingsTypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsTypeMedia, ScriptScriptAndVersionSettingGetResponseBindingsTypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsTypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsTypePipelines, ScriptScriptAndVersionSettingGetResponseBindingsTypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsTypeRatelimit, ScriptScriptAndVersionSettingGetResponseBindingsTypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsTypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsTypeSendEmail, ScriptScriptAndVersionSettingGetResponseBindingsTypeService, ScriptScriptAndVersionSettingGetResponseBindingsTypeTextBlob, ScriptScriptAndVersionSettingGetResponseBindingsTypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsTypeVersionMetadata, ScriptScriptAndVersionSettingGetResponseBindingsTypeSecretsStoreSecret, ScriptScriptAndVersionSettingGetResponseBindingsTypeFlagship, ScriptScriptAndVersionSettingGetResponseBindingsTypeSecretKey, ScriptScriptAndVersionSettingGetResponseBindingsTypeWorkflow, ScriptScriptAndVersionSettingGetResponseBindingsTypeWasmModule, ScriptScriptAndVersionSettingGetResponseBindingsTypeVPCService, ScriptScriptAndVersionSettingGetResponseBindingsTypeVPCNetwork:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptScriptAndVersionSettingGetResponseBindingsFormat string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsFormatRaw   ScriptScriptAndVersionSettingGetResponseBindingsFormat = "raw"
	ScriptScriptAndVersionSettingGetResponseBindingsFormatPkcs8 ScriptScriptAndVersionSettingGetResponseBindingsFormat = "pkcs8"
	ScriptScriptAndVersionSettingGetResponseBindingsFormatSpki  ScriptScriptAndVersionSettingGetResponseBindingsFormat = "spki"
	ScriptScriptAndVersionSettingGetResponseBindingsFormatJwk   ScriptScriptAndVersionSettingGetResponseBindingsFormat = "jwk"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsFormat) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsFormatRaw, ScriptScriptAndVersionSettingGetResponseBindingsFormatPkcs8, ScriptScriptAndVersionSettingGetResponseBindingsFormatSpki, ScriptScriptAndVersionSettingGetResponseBindingsFormatJwk:
		return true
	}
	return false
}

// Enables Gateway identity for the binding. Requires network_id to be
// "cf1:network" and cannot be combined with tunnel_id.
type ScriptScriptAndVersionSettingGetResponseBindingsIdentity string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsIdentityRuntimeEmailAlpha ScriptScriptAndVersionSettingGetResponseBindingsIdentity = "runtime-email-alpha"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsIdentity) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsIdentityRuntimeEmailAlpha:
		return true
	}
	return false
}

// The
// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
// of the R2 bucket.
type ScriptScriptAndVersionSettingGetResponseBindingsJurisdiction string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsJurisdictionEu          ScriptScriptAndVersionSettingGetResponseBindingsJurisdiction = "eu"
	ScriptScriptAndVersionSettingGetResponseBindingsJurisdictionFedramp     ScriptScriptAndVersionSettingGetResponseBindingsJurisdiction = "fedramp"
	ScriptScriptAndVersionSettingGetResponseBindingsJurisdictionFedrampHigh ScriptScriptAndVersionSettingGetResponseBindingsJurisdiction = "fedramp-high"
	ScriptScriptAndVersionSettingGetResponseBindingsJurisdictionUs          ScriptScriptAndVersionSettingGetResponseBindingsJurisdiction = "us"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsJurisdiction) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsJurisdictionEu, ScriptScriptAndVersionSettingGetResponseBindingsJurisdictionFedramp, ScriptScriptAndVersionSettingGetResponseBindingsJurisdictionFedrampHigh, ScriptScriptAndVersionSettingGetResponseBindingsJurisdictionUs:
		return true
	}
	return false
}

// Global CacheW configuration for the Worker. When caching is on, the platform
// provisions a `cloudflare.app` zone for the Worker. A `type: worker` entry in the
// `exports` map can override this value for a single entrypoint.
type ScriptScriptAndVersionSettingGetResponseCacheOptions struct {
	// Whether caching is enabled for this Worker.
	Enabled bool `json:"enabled" api:"required"`
	// Whether cached responses are shared across Worker version uploads. This is
	// independent of `enabled`. It can stay true while caching is off, so the
	// preference survives turning caching off and back on.
	CrossVersionCache bool                                                     `json:"cross_version_cache"`
	JSON              scriptScriptAndVersionSettingGetResponseCacheOptionsJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseCacheOptionsJSON contains the JSON
// metadata for the struct [ScriptScriptAndVersionSettingGetResponseCacheOptions]
type scriptScriptAndVersionSettingGetResponseCacheOptionsJSON struct {
	Enabled           apijson.Field
	CrossVersionCache apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseCacheOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseCacheOptionsJSON) RawJSON() string {
	return r.raw
}

// A single entry in the `exports` map, keyed by export name (a `WorkerEntrypoint`
// class name, a Durable Object class name, or `default` for the Worker's default
// export). The `type` discriminator selects the top-level shape: `worker`
// entrypoint entries may carry `cache` configuration, while `durable-object`
// entries are further refined by the optional `state` field (default `created`).
// Tombstone states (`deleted`, `renamed`, `transferred`) express destructive
// lifecycle operations declaratively; `expecting-transfer` is the live target side
// of a transfer. The server validates the exact per-(type, state) field
// combinations; fields not listed for a variant are rejected.
type ScriptScriptAndVersionSettingGetResponseExport struct {
	// Marks this entry as a Worker entrypoint export.
	Type ScriptScriptAndVersionSettingGetResponseExportsType `json:"type" api:"required"`
	// This field can have the runtime type of
	// [ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportCache].
	Cache interface{} `json:"cache"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container string `json:"container"`
	// Live export. May be omitted; defaults to `created`.
	State ScriptScriptAndVersionSettingGetResponseExportsState `json:"state"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage ScriptScriptAndVersionSettingGetResponseExportsStorage `json:"storage"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom string                                             `json:"transfer_from"`
	JSON         scriptScriptAndVersionSettingGetResponseExportJSON `json:"-"`
	union        ScriptScriptAndVersionSettingGetResponseExportsUnion
}

// scriptScriptAndVersionSettingGetResponseExportJSON contains the JSON metadata
// for the struct [ScriptScriptAndVersionSettingGetResponseExport]
type scriptScriptAndVersionSettingGetResponseExportJSON struct {
	Type         apijson.Field
	Cache        apijson.Field
	Container    apijson.Field
	State        apijson.Field
	Storage      apijson.Field
	TransferFrom apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r scriptScriptAndVersionSettingGetResponseExportJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptScriptAndVersionSettingGetResponseExport) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptScriptAndVersionSettingGetResponseExport{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptScriptAndVersionSettingGetResponseExportsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExport],
// [ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExport],
// [ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExport],
// [ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExport],
// [ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExport],
// [ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExport].
func (r ScriptScriptAndVersionSettingGetResponseExport) AsUnion() ScriptScriptAndVersionSettingGetResponseExportsUnion {
	return r.union
}

// A single entry in the `exports` map, keyed by export name (a `WorkerEntrypoint`
// class name, a Durable Object class name, or `default` for the Worker's default
// export). The `type` discriminator selects the top-level shape: `worker`
// entrypoint entries may carry `cache` configuration, while `durable-object`
// entries are further refined by the optional `state` field (default `created`).
// Tombstone states (`deleted`, `renamed`, `transferred`) express destructive
// lifecycle operations declaratively; `expecting-transfer` is the live target side
// of a transfer. The server validates the exact per-(type, state) field
// combinations; fields not listed for a variant are rejected.
//
// Union satisfied by
// [ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExport],
// [ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExport],
// [ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExport],
// [ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExport],
// [ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExport]
// or
// [ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExport].
type ScriptScriptAndVersionSettingGetResponseExportsUnion interface {
	implementsScriptScriptAndVersionSettingGetResponseExport()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptScriptAndVersionSettingGetResponseExportsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExport{}),
			DiscriminatorValue: "worker",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExport{}),
			DiscriminatorValue: "durable-object",
		},
	)
}

// A named Worker entrypoint export (`type: worker`). Worker entrypoints are always
// live (`state: created`) and carry no storage or lifecycle fields. The optional
// `cache` block overrides the Worker's global `cache_options.enabled` for this
// entrypoint.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExport struct {
	// Marks this entry as a Worker entrypoint export.
	Type ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportType `json:"type" api:"required"`
	// Cache override for this entrypoint. Overrides the Worker's global
	// `cache_options.enabled` for this entrypoint only.
	Cache ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportCache `json:"cache"`
	// Live export. May be omitted; defaults to `created`.
	State ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportState `json:"state"`
	JSON  scriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportJSON  `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportJSON contains
// the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExport]
type scriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportJSON struct {
	Type        apijson.Field
	Cache       apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExport) implementsScriptScriptAndVersionSettingGetResponseExport() {
}

// Marks this entry as a Worker entrypoint export.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportType string

const (
	ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportTypeWorker ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportType = "worker"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportTypeWorker:
		return true
	}
	return false
}

// Cache override for this entrypoint. Overrides the Worker's global
// `cache_options.enabled` for this entrypoint only.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportCache struct {
	// Whether caching is enabled for this entrypoint.
	Enabled bool                                                                        `json:"enabled" api:"required"`
	JSON    scriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportCacheJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportCacheJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportCache]
type scriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportCacheJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportCacheJSON) RawJSON() string {
	return r.raw
}

// Live export. May be omitted; defaults to `created`.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportState string

const (
	ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportStateCreated ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportState = "created"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsWorkersWorkerExportStateCreated:
		return true
	}
	return false
}

// A live Durable Object export (`state: created`, the default). The platform
// auto-provisions the namespace on first deploy, matches it on subsequent deploys,
// and never mutates or deletes it as a side effect of a code-only change.
// `storage` is required; `renamed_to`, `transferred_to` and `transfer_from` are
// not allowed on a live entry.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExport struct {
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportStorage `json:"storage" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportType `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container string `json:"container"`
	// Live export. May be omitted; defaults to `created`.
	State ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportState `json:"state"`
	JSON  scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportJSON  `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExport]
type scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportJSON struct {
	Storage     apijson.Field
	Type        apijson.Field
	Container   apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExport) implementsScriptScriptAndVersionSettingGetResponseExport() {
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportStorage string

const (
	ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportStorageSqlite   ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportStorage = "sqlite"
	ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportStorageLegacyKV ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportStorage = "legacy-kv"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportStorage) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportStorageSqlite, ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportStorageLegacyKV:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportType string

const (
	ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportTypeDurableObject ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportType = "durable-object"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportTypeDurableObject:
		return true
	}
	return false
}

// Live export. May be omitted; defaults to `created`.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportState string

const (
	ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportStateCreated ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportState = "created"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExportStateCreated:
		return true
	}
	return false
}

// A `deleted` tombstone: retires the provisioned namespace for this class and all
// of its data. The class must be absent from the uploaded code and no other Worker
// in the account may bind to the namespace, otherwise the deploy is rejected. No
// other fields are allowed. Deletion is irreversible.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExport struct {
	// Tombstone that deletes the namespace.
	State ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExportState `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExportType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExportJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExportJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExport]
type scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExportJSON struct {
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExport) implementsScriptScriptAndVersionSettingGetResponseExport() {
}

// Tombstone that deletes the namespace.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExportState string

const (
	ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExportStateDeleted ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExportState = "deleted"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExportState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExportStateDeleted:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExportType string

const (
	ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExportTypeDurableObject ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExportType = "durable-object"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExportType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectDeletedExportTypeDurableObject:
		return true
	}
	return false
}

// A `renamed` tombstone: rewrites the provisioned namespace's class name from this
// map key to `renamed_to`. The source class may stay in code during the rollout
// window (an info notice is emitted). `storage`, `transferred_to` and
// `transfer_from` are not allowed.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExport struct {
	// Tombstone that renames the namespace's class.
	State ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExportState `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExportType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExportJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExportJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExport]
type scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExportJSON struct {
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExport) implementsScriptScriptAndVersionSettingGetResponseExport() {
}

// Tombstone that renames the namespace's class.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExportState string

const (
	ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExportStateRenamed ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExportState = "renamed"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExportState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExportStateRenamed:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExportType string

const (
	ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExportTypeDurableObject ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExportType = "durable-object"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExportType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectRenamedExportTypeDurableObject:
		return true
	}
	return false
}

// A `transferred` tombstone (source side of a two-phase transfer): hands ownership
// of the provisioned namespace to another script in the same account, named by
// `transferred_to`. The target must have already deployed a matching
// `expecting-transfer` entry. The source class may stay in code during the rollout
// window (an info notice is emitted). `storage`, `renamed_to` and `transfer_from`
// are not allowed.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExport struct {
	// Tombstone that transfers the namespace to another script.
	State ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExportState `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExportType `json:"type" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExportJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExportJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExport]
type scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExportJSON struct {
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExport) implementsScriptScriptAndVersionSettingGetResponseExport() {
}

// Tombstone that transfers the namespace to another script.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExportState string

const (
	ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExportStateTransferred ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExportState = "transferred"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExportState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExportStateTransferred:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExportType string

const (
	ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExportTypeDurableObject ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExportType = "durable-object"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExportType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectTransferredExportTypeDurableObject:
		return true
	}
	return false
}

// The target side of a two-phase transfer (`state: expecting-transfer`). Declares
// that this script expects to receive a namespace for this class from the
// `transfer_from` script. This is a live entry, not a tombstone: bindings resolve
// through the source's namespace until the source commits with a `transferred`
// tombstone. `storage` and `transfer_from` are required; `renamed_to` and
// `transferred_to` are not allowed.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExport struct {
	// Target side of a two-phase transfer.
	State ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportState `json:"state" api:"required"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportStorage `json:"storage" api:"required"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom string `json:"transfer_from" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportType `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object once the transfer settles. Valid only on live entries.
	Container string                                                                                         `json:"container"`
	JSON      scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExport]
type scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportJSON struct {
	State        apijson.Field
	Storage      apijson.Field
	TransferFrom apijson.Field
	Type         apijson.Field
	Container    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExport) implementsScriptScriptAndVersionSettingGetResponseExport() {
}

// Target side of a two-phase transfer.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportState string

const (
	ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportStateExpectingTransfer ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportState = "expecting-transfer"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportStateExpectingTransfer:
		return true
	}
	return false
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportStorage string

const (
	ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportStorageSqlite   ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportStorage = "sqlite"
	ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportStorageLegacyKV ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportStorage = "legacy-kv"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportStorage) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportStorageSqlite, ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportStorageLegacyKV:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportType string

const (
	ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportTypeDurableObject ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportType = "durable-object"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsWorkersDurableObjectExpectingTransferExportTypeDurableObject:
		return true
	}
	return false
}

// Marks this entry as a Worker entrypoint export.
type ScriptScriptAndVersionSettingGetResponseExportsType string

const (
	ScriptScriptAndVersionSettingGetResponseExportsTypeWorker        ScriptScriptAndVersionSettingGetResponseExportsType = "worker"
	ScriptScriptAndVersionSettingGetResponseExportsTypeDurableObject ScriptScriptAndVersionSettingGetResponseExportsType = "durable-object"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsTypeWorker, ScriptScriptAndVersionSettingGetResponseExportsTypeDurableObject:
		return true
	}
	return false
}

// Live export. May be omitted; defaults to `created`.
type ScriptScriptAndVersionSettingGetResponseExportsState string

const (
	ScriptScriptAndVersionSettingGetResponseExportsStateCreated           ScriptScriptAndVersionSettingGetResponseExportsState = "created"
	ScriptScriptAndVersionSettingGetResponseExportsStateDeleted           ScriptScriptAndVersionSettingGetResponseExportsState = "deleted"
	ScriptScriptAndVersionSettingGetResponseExportsStateRenamed           ScriptScriptAndVersionSettingGetResponseExportsState = "renamed"
	ScriptScriptAndVersionSettingGetResponseExportsStateTransferred       ScriptScriptAndVersionSettingGetResponseExportsState = "transferred"
	ScriptScriptAndVersionSettingGetResponseExportsStateExpectingTransfer ScriptScriptAndVersionSettingGetResponseExportsState = "expecting-transfer"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsStateCreated, ScriptScriptAndVersionSettingGetResponseExportsStateDeleted, ScriptScriptAndVersionSettingGetResponseExportsStateRenamed, ScriptScriptAndVersionSettingGetResponseExportsStateTransferred, ScriptScriptAndVersionSettingGetResponseExportsStateExpectingTransfer:
		return true
	}
	return false
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type ScriptScriptAndVersionSettingGetResponseExportsStorage string

const (
	ScriptScriptAndVersionSettingGetResponseExportsStorageSqlite   ScriptScriptAndVersionSettingGetResponseExportsStorage = "sqlite"
	ScriptScriptAndVersionSettingGetResponseExportsStorageLegacyKV ScriptScriptAndVersionSettingGetResponseExportsStorage = "legacy-kv"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsStorage) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsStorageSqlite, ScriptScriptAndVersionSettingGetResponseExportsStorageLegacyKV:
		return true
	}
	return false
}

// Summary of the declarative exports reconciliation that ran on this upload.
// Populated only when the uploaded metadata included an `exports` block. Durable
// Object entries drive reconciliation; `type: worker` entries do not contribute to
// this summary.
type ScriptScriptAndVersionSettingGetResponseExportsReconciliation struct {
	// Class names for which a new namespace was provisioned.
	Created []string `json:"created" api:"required"`
	// Class names whose namespace was deleted by a `deleted` tombstone.
	Deleted []string `json:"deleted" api:"required"`
	// Non-blocking info entries (stale tombstones, tombstone applied with class still
	// in code). See `exports_reconciliation_info`.
	Info []ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfo `json:"info" api:"required"`
	// Source class names whose tombstone entry is now stale and safe to delete from
	// `exports` (no remaining referencing scripts).
	RemovableEntries []string `json:"removable_entries" api:"required"`
	// Applied `renamed` tombstones.
	Renamed []ScriptScriptAndVersionSettingGetResponseExportsReconciliationRenamed `json:"renamed" api:"required"`
	// Phase-1 transfer hints recorded on the target side.
	TransferPending []ScriptScriptAndVersionSettingGetResponseExportsReconciliationTransferPending `json:"transfer_pending" api:"required"`
	// Committed `transferred` tombstones (phase-2).
	Transferred []ScriptScriptAndVersionSettingGetResponseExportsReconciliationTransferred `json:"transferred" api:"required"`
	// Class names whose provisioned namespace was mutated in place.
	Updated []string `json:"updated" api:"required"`
	// Non-blocking warnings. See `exports_reconciliation_warning`.
	Warnings []ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarning `json:"warnings" api:"required"`
	JSON     scriptScriptAndVersionSettingGetResponseExportsReconciliationJSON      `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseExportsReconciliationJSON contains the
// JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseExportsReconciliation]
type scriptScriptAndVersionSettingGetResponseExportsReconciliationJSON struct {
	Created          apijson.Field
	Deleted          apijson.Field
	Info             apijson.Field
	RemovableEntries apijson.Field
	Renamed          apijson.Field
	TransferPending  apijson.Field
	Transferred      apijson.Field
	Updated          apijson.Field
	Warnings         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseExportsReconciliation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseExportsReconciliationJSON) RawJSON() string {
	return r.raw
}

// A non-blocking reconciliation info entry. Emitted for stale tombstones (a no-op
// on this deploy) and for tombstones applied with the source class still in code
// (the supported zero-downtime rollout pattern).
type ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfo struct {
	// The class name the info entry is about.
	Class string `json:"class" api:"required"`
	// Human-readable explanation.
	Message string `json:"message" api:"required"`
	// Stable, machine-readable tag identifying which reconciliation scenario produced
	// an error, warning, or info entry. Clients may branch on this value instead of
	// parsing `message`.
	Scenario ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario `json:"scenario" api:"required"`
	// The provisioned namespace the entry relates to, when applicable.
	NamespaceID string `json:"namespace_id" format:"uuid"`
	// Other Workers in the account that still bind to the affected class. Advisory:
	// while non-empty the tombstone is not yet safe to remove — redeploy these Workers
	// with bindings re-pointed first.
	ReferencingScripts []string                                                              `json:"referencing_scripts"`
	JSON               scriptScriptAndVersionSettingGetResponseExportsReconciliationInfoJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseExportsReconciliationInfoJSON contains
// the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfo]
type scriptScriptAndVersionSettingGetResponseExportsReconciliationInfoJSON struct {
	Class              apijson.Field
	Message            apijson.Field
	Scenario           apijson.Field
	NamespaceID        apijson.Field
	ReferencingScripts apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseExportsReconciliationInfoJSON) RawJSON() string {
	return r.raw
}

// Stable, machine-readable tag identifying which reconciliation scenario produced
// an error, warning, or info entry. Clients may branch on this value instead of
// parsing `message`.
type ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario string

const (
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioCodeClassNotInExports                     ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "code_class_not_in_exports"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioProvisionedClassMissingFromConfig         ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "provisioned_class_missing_from_config"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioConfigExportNotInCode                     ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "config_export_not_in_code"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioConfigReferencesNonexistentClass          ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "config_references_nonexistent_class"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioOrphanedProvisionedNamespace              ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "orphaned_provisioned_namespace"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioStorageTypeMismatch                       ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "storage_type_mismatch"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioFreeTierRequiresSqlite                    ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "free_tier_requires_sqlite"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioInvalidExport                             ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "invalid_export"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTombstoneDeleteClassStillInCode           ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "tombstone_delete_class_still_in_code"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTombstoneDeleteBlockedByExternalBindings  ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "tombstone_delete_blocked_by_external_bindings"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTombstoneRenamedToOccupied                ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "tombstone_renamed_to_occupied"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTransferredPendingNotFound                ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "transferred_pending_not_found"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTransferredTargetMissing                  ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "transferred_target_missing"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTransferredTargetMismatch                 ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "transferred_target_mismatch"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioPhaseOneTransferSourceMissing             ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "phase_one_transfer_source_missing"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioPhaseOneTransferSourceNamespaceMissing    ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "phase_one_transfer_source_namespace_missing"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioPhaseOneTransferTargetClassProvisioned    ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "phase_one_transfer_target_class_provisioned"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioPhaseOneTransferAfterCommitMismatch       ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "phase_one_transfer_after_commit_mismatch"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioPhaseOneTransferDuplicate                 ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "phase_one_transfer_duplicate"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioPhaseOneTransferTargetInDispatchNamespace ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "phase_one_transfer_target_in_dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioPhaseOneTransferSourceInDispatchNamespace ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "phase_one_transfer_source_in_dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTransferredSourceInDispatchNamespace      ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "transferred_source_in_dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTransferredTargetInDispatchNamespace      ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "transferred_target_in_dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioContainerUndeclaredReference              ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "container_undeclared_reference"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioContainerClassNotDurableObject            ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "container_class_not_durable_object"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioContainerWiringInconsistent               ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "container_wiring_inconsistent"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioContainerMultipleDurableObjects           ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "container_multiple_durable_objects"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTransferContainerParityMismatch           ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "transfer_container_parity_mismatch"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTransferContainerParityMismatchOnCommit   ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "transfer_container_parity_mismatch_on_commit"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTombstoneClassStillInCode                 ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "tombstone_class_still_in_code"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioStaleTombstone                            ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "stale_tombstone"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTransferReceiveAlreadyApplied             ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "transfer_receive_already_applied"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTransferReceiveCleanupComplete            ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario = "transfer_receive_cleanup_complete"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenario) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioCodeClassNotInExports, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioProvisionedClassMissingFromConfig, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioConfigExportNotInCode, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioConfigReferencesNonexistentClass, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioOrphanedProvisionedNamespace, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioStorageTypeMismatch, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioFreeTierRequiresSqlite, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioInvalidExport, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTombstoneDeleteClassStillInCode, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTombstoneDeleteBlockedByExternalBindings, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTombstoneRenamedToOccupied, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTransferredPendingNotFound, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTransferredTargetMissing, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTransferredTargetMismatch, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioPhaseOneTransferSourceMissing, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioPhaseOneTransferSourceNamespaceMissing, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioPhaseOneTransferTargetClassProvisioned, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioPhaseOneTransferAfterCommitMismatch, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioPhaseOneTransferDuplicate, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioPhaseOneTransferTargetInDispatchNamespace, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioPhaseOneTransferSourceInDispatchNamespace, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTransferredSourceInDispatchNamespace, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTransferredTargetInDispatchNamespace, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioContainerUndeclaredReference, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioContainerClassNotDurableObject, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioContainerWiringInconsistent, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioContainerMultipleDurableObjects, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTransferContainerParityMismatch, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTransferContainerParityMismatchOnCommit, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTombstoneClassStillInCode, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioStaleTombstone, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTransferReceiveAlreadyApplied, ScriptScriptAndVersionSettingGetResponseExportsReconciliationInfoScenarioTransferReceiveCleanupComplete:
		return true
	}
	return false
}

// A single applied `renamed` tombstone.
type ScriptScriptAndVersionSettingGetResponseExportsReconciliationRenamed struct {
	// The original (source) class name.
	From string `json:"from" api:"required"`
	// The new class name (`renamed_to`).
	To   string                                                                   `json:"to" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseExportsReconciliationRenamedJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseExportsReconciliationRenamedJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseExportsReconciliationRenamed]
type scriptScriptAndVersionSettingGetResponseExportsReconciliationRenamedJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseExportsReconciliationRenamed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseExportsReconciliationRenamedJSON) RawJSON() string {
	return r.raw
}

// A single phase-1 transfer hint recorded on the target side (a live
// `expecting-transfer` entry).
type ScriptScriptAndVersionSettingGetResponseExportsReconciliationTransferPending struct {
	// The target-side class name awaiting transfer.
	Class string `json:"class" api:"required"`
	// The source script the namespace will be transferred from.
	From string                                                                           `json:"from" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseExportsReconciliationTransferPendingJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseExportsReconciliationTransferPendingJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseExportsReconciliationTransferPending]
type scriptScriptAndVersionSettingGetResponseExportsReconciliationTransferPendingJSON struct {
	Class       apijson.Field
	From        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseExportsReconciliationTransferPending) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseExportsReconciliationTransferPendingJSON) RawJSON() string {
	return r.raw
}

// A single committed `transferred` tombstone (phase-2 commit).
type ScriptScriptAndVersionSettingGetResponseExportsReconciliationTransferred struct {
	// The source class name that was transferred.
	Class string `json:"class" api:"required"`
	// The transfer phase. Currently always `committed`.
	Phase ScriptScriptAndVersionSettingGetResponseExportsReconciliationTransferredPhase `json:"phase" api:"required"`
	// The destination script that now owns the namespace.
	To   string                                                                       `json:"to" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponseExportsReconciliationTransferredJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseExportsReconciliationTransferredJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseExportsReconciliationTransferred]
type scriptScriptAndVersionSettingGetResponseExportsReconciliationTransferredJSON struct {
	Class       apijson.Field
	Phase       apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseExportsReconciliationTransferred) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseExportsReconciliationTransferredJSON) RawJSON() string {
	return r.raw
}

// The transfer phase. Currently always `committed`.
type ScriptScriptAndVersionSettingGetResponseExportsReconciliationTransferredPhase string

const (
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationTransferredPhaseCommitted ScriptScriptAndVersionSettingGetResponseExportsReconciliationTransferredPhase = "committed"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsReconciliationTransferredPhase) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsReconciliationTransferredPhaseCommitted:
		return true
	}
	return false
}

// A non-blocking reconciliation warning. Reserved: no scenario populates this
// array today (`code_class_not_in_exports` is surfaced as info and
// `provisioned_class_missing_from_config` is a hard error). Clients should still
// surface any entries that appear.
type ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarning struct {
	// The class name the warning is about.
	Class string `json:"class" api:"required"`
	// Human-readable explanation of the warning.
	Message string `json:"message" api:"required"`
	// Stable, machine-readable tag identifying which reconciliation scenario produced
	// an error, warning, or info entry. Clients may branch on this value instead of
	// parsing `message`.
	Scenario ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario `json:"scenario" api:"required"`
	// The provisioned namespace the warning relates to, when applicable.
	NamespaceID string                                                                   `json:"namespace_id" format:"uuid"`
	JSON        scriptScriptAndVersionSettingGetResponseExportsReconciliationWarningJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseExportsReconciliationWarningJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarning]
type scriptScriptAndVersionSettingGetResponseExportsReconciliationWarningJSON struct {
	Class       apijson.Field
	Message     apijson.Field
	Scenario    apijson.Field
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseExportsReconciliationWarningJSON) RawJSON() string {
	return r.raw
}

// Stable, machine-readable tag identifying which reconciliation scenario produced
// an error, warning, or info entry. Clients may branch on this value instead of
// parsing `message`.
type ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario string

const (
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioCodeClassNotInExports                     ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "code_class_not_in_exports"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioProvisionedClassMissingFromConfig         ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "provisioned_class_missing_from_config"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioConfigExportNotInCode                     ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "config_export_not_in_code"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioConfigReferencesNonexistentClass          ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "config_references_nonexistent_class"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioOrphanedProvisionedNamespace              ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "orphaned_provisioned_namespace"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioStorageTypeMismatch                       ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "storage_type_mismatch"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioFreeTierRequiresSqlite                    ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "free_tier_requires_sqlite"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioInvalidExport                             ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "invalid_export"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTombstoneDeleteClassStillInCode           ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "tombstone_delete_class_still_in_code"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTombstoneDeleteBlockedByExternalBindings  ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "tombstone_delete_blocked_by_external_bindings"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTombstoneRenamedToOccupied                ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "tombstone_renamed_to_occupied"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTransferredPendingNotFound                ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "transferred_pending_not_found"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTransferredTargetMissing                  ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "transferred_target_missing"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTransferredTargetMismatch                 ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "transferred_target_mismatch"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioPhaseOneTransferSourceMissing             ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "phase_one_transfer_source_missing"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioPhaseOneTransferSourceNamespaceMissing    ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "phase_one_transfer_source_namespace_missing"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioPhaseOneTransferTargetClassProvisioned    ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "phase_one_transfer_target_class_provisioned"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioPhaseOneTransferAfterCommitMismatch       ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "phase_one_transfer_after_commit_mismatch"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioPhaseOneTransferDuplicate                 ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "phase_one_transfer_duplicate"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioPhaseOneTransferTargetInDispatchNamespace ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "phase_one_transfer_target_in_dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioPhaseOneTransferSourceInDispatchNamespace ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "phase_one_transfer_source_in_dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTransferredSourceInDispatchNamespace      ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "transferred_source_in_dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTransferredTargetInDispatchNamespace      ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "transferred_target_in_dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioContainerUndeclaredReference              ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "container_undeclared_reference"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioContainerClassNotDurableObject            ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "container_class_not_durable_object"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioContainerWiringInconsistent               ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "container_wiring_inconsistent"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioContainerMultipleDurableObjects           ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "container_multiple_durable_objects"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTransferContainerParityMismatch           ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "transfer_container_parity_mismatch"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTransferContainerParityMismatchOnCommit   ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "transfer_container_parity_mismatch_on_commit"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTombstoneClassStillInCode                 ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "tombstone_class_still_in_code"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioStaleTombstone                            ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "stale_tombstone"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTransferReceiveAlreadyApplied             ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "transfer_receive_already_applied"
	ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTransferReceiveCleanupComplete            ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario = "transfer_receive_cleanup_complete"
)

func (r ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenario) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioCodeClassNotInExports, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioProvisionedClassMissingFromConfig, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioConfigExportNotInCode, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioConfigReferencesNonexistentClass, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioOrphanedProvisionedNamespace, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioStorageTypeMismatch, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioFreeTierRequiresSqlite, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioInvalidExport, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTombstoneDeleteClassStillInCode, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTombstoneDeleteBlockedByExternalBindings, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTombstoneRenamedToOccupied, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTransferredPendingNotFound, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTransferredTargetMissing, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTransferredTargetMismatch, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioPhaseOneTransferSourceMissing, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioPhaseOneTransferSourceNamespaceMissing, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioPhaseOneTransferTargetClassProvisioned, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioPhaseOneTransferAfterCommitMismatch, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioPhaseOneTransferDuplicate, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioPhaseOneTransferTargetInDispatchNamespace, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioPhaseOneTransferSourceInDispatchNamespace, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTransferredSourceInDispatchNamespace, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTransferredTargetInDispatchNamespace, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioContainerUndeclaredReference, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioContainerClassNotDurableObject, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioContainerWiringInconsistent, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioContainerMultipleDurableObjects, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTransferContainerParityMismatch, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTransferContainerParityMismatchOnCommit, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTombstoneClassStillInCode, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioStaleTombstone, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTransferReceiveAlreadyApplied, ScriptScriptAndVersionSettingGetResponseExportsReconciliationWarningsScenarioTransferReceiveCleanupComplete:
		return true
	}
	return false
}

// Limits to apply for this Worker.
type ScriptScriptAndVersionSettingGetResponseLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
	CPUMs int64 `json:"cpu_ms"`
	// The number of subrequests this Worker can make per request.
	Subrequests int64                                              `json:"subrequests"`
	JSON        scriptScriptAndVersionSettingGetResponseLimitsJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseLimitsJSON contains the JSON metadata
// for the struct [ScriptScriptAndVersionSettingGetResponseLimits]
type scriptScriptAndVersionSettingGetResponseLimitsJSON struct {
	CPUMs       apijson.Field
	Subrequests apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseLimitsJSON) RawJSON() string {
	return r.raw
}

// Migrations to apply for Durable Objects associated with this Worker.
type ScriptScriptAndVersionSettingGetResponseMigrations struct {
	// This field can have the runtime type of [[]string].
	DeletedClasses interface{} `json:"deleted_classes"`
	// This field can have the runtime type of [[]string].
	NewClasses interface{} `json:"new_classes"`
	// This field can have the runtime type of [[]string].
	NewSqliteClasses interface{} `json:"new_sqlite_classes"`
	// This field can have the runtime type of [[]SingleStepMigrationRenamedClass].
	RenamedClasses interface{} `json:"renamed_classes"`
	// This field can have the runtime type of [[]MigrationStep].
	Steps interface{} `json:"steps"`
	// This field can have the runtime type of [[]SingleStepMigrationTransferredClass].
	TransferredClasses interface{}                                            `json:"transferred_classes"`
	JSON               scriptScriptAndVersionSettingGetResponseMigrationsJSON `json:"-"`
	union              ScriptScriptAndVersionSettingGetResponseMigrationsUnion
}

// scriptScriptAndVersionSettingGetResponseMigrationsJSON contains the JSON
// metadata for the struct [ScriptScriptAndVersionSettingGetResponseMigrations]
type scriptScriptAndVersionSettingGetResponseMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewSqliteClasses   apijson.Field
	RenamedClasses     apijson.Field
	Steps              apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r scriptScriptAndVersionSettingGetResponseMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptScriptAndVersionSettingGetResponseMigrations) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptScriptAndVersionSettingGetResponseMigrations{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptScriptAndVersionSettingGetResponseMigrationsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SingleStepMigration],
// [ScriptScriptAndVersionSettingGetResponseMigrationsWorkersMultipleStepMigrations].
func (r ScriptScriptAndVersionSettingGetResponseMigrations) AsUnion() ScriptScriptAndVersionSettingGetResponseMigrationsUnion {
	return r.union
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by [SingleStepMigration] or
// [ScriptScriptAndVersionSettingGetResponseMigrationsWorkersMultipleStepMigrations].
type ScriptScriptAndVersionSettingGetResponseMigrationsUnion interface {
	implementsScriptScriptAndVersionSettingGetResponseMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptScriptAndVersionSettingGetResponseMigrationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SingleStepMigration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseMigrationsWorkersMultipleStepMigrations{}),
		},
	)
}

type ScriptScriptAndVersionSettingGetResponseMigrationsWorkersMultipleStepMigrations struct {
	JSON scriptScriptAndVersionSettingGetResponseMigrationsWorkersMultipleStepMigrationsJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseMigrationsWorkersMultipleStepMigrationsJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseMigrationsWorkersMultipleStepMigrations]
type scriptScriptAndVersionSettingGetResponseMigrationsWorkersMultipleStepMigrationsJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseMigrationsWorkersMultipleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseMigrationsWorkersMultipleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseMigrationsWorkersMultipleStepMigrations) implementsScriptScriptAndVersionSettingGetResponseMigrations() {
}

// Observability settings for the Worker.
type ScriptScriptAndVersionSettingGetResponseObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled bool `json:"enabled" api:"required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate float64 `json:"head_sampling_rate" api:"nullable"`
	// Log settings for the Worker.
	Logs ScriptScriptAndVersionSettingGetResponseObservabilityLogs `json:"logs" api:"nullable"`
	// Whether query strings are removed from request URLs in logs and traces.
	RedactQueryString bool `json:"redact_query_string"`
	// Trace settings for the Worker.
	Traces ScriptScriptAndVersionSettingGetResponseObservabilityTraces `json:"traces" api:"nullable"`
	JSON   scriptScriptAndVersionSettingGetResponseObservabilityJSON   `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseObservabilityJSON contains the JSON
// metadata for the struct [ScriptScriptAndVersionSettingGetResponseObservability]
type scriptScriptAndVersionSettingGetResponseObservabilityJSON struct {
	Enabled           apijson.Field
	HeadSamplingRate  apijson.Field
	Logs              apijson.Field
	RedactQueryString apijson.Field
	Traces            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseObservabilityJSON) RawJSON() string {
	return r.raw
}

// Log settings for the Worker.
type ScriptScriptAndVersionSettingGetResponseObservabilityLogs struct {
	// Whether logs are enabled for the Worker.
	Enabled bool `json:"enabled" api:"required"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs bool `json:"invocation_logs" api:"required"`
	// A list of destinations where logs will be exported to.
	Destinations []string `json:"destinations"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate float64 `json:"head_sampling_rate" api:"nullable"`
	// Whether log persistence is enabled for the Worker.
	Persist bool                                                          `json:"persist"`
	JSON    scriptScriptAndVersionSettingGetResponseObservabilityLogsJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseObservabilityLogsJSON contains the JSON
// metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseObservabilityLogs]
type scriptScriptAndVersionSettingGetResponseObservabilityLogsJSON struct {
	Enabled          apijson.Field
	InvocationLogs   apijson.Field
	Destinations     apijson.Field
	HeadSamplingRate apijson.Field
	Persist          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// Trace settings for the Worker.
type ScriptScriptAndVersionSettingGetResponseObservabilityTraces struct {
	// A list of destinations where traces will be exported to.
	Destinations []string `json:"destinations"`
	// Whether traces are enabled for the Worker.
	Enabled bool `json:"enabled"`
	// The sampling rate for traces. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate float64 `json:"head_sampling_rate" api:"nullable"`
	// Whether trace persistence is enabled for the Worker.
	Persist bool `json:"persist"`
	// Controls how inbound trace context (traceparent/tracestate) headers on incoming
	// requests are handled. "authenticated" honors inbound trace context only when
	// accompanied by a valid trace auth token. "accept" unconditionally accepts
	// inbound trace context. Requires the trace propagation feature to be enabled.
	// Returns null when the trace propagation feature is not enabled for the account.
	PropagationPolicy ScriptScriptAndVersionSettingGetResponseObservabilityTracesPropagationPolicy `json:"propagation_policy" api:"nullable"`
	JSON              scriptScriptAndVersionSettingGetResponseObservabilityTracesJSON              `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseObservabilityTracesJSON contains the
// JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseObservabilityTraces]
type scriptScriptAndVersionSettingGetResponseObservabilityTracesJSON struct {
	Destinations      apijson.Field
	Enabled           apijson.Field
	HeadSamplingRate  apijson.Field
	Persist           apijson.Field
	PropagationPolicy apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseObservabilityTraces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseObservabilityTracesJSON) RawJSON() string {
	return r.raw
}

// Controls how inbound trace context (traceparent/tracestate) headers on incoming
// requests are handled. "authenticated" honors inbound trace context only when
// accompanied by a valid trace auth token. "accept" unconditionally accepts
// inbound trace context. Requires the trace propagation feature to be enabled.
// Returns null when the trace propagation feature is not enabled for the account.
type ScriptScriptAndVersionSettingGetResponseObservabilityTracesPropagationPolicy string

const (
	ScriptScriptAndVersionSettingGetResponseObservabilityTracesPropagationPolicyAuthenticated ScriptScriptAndVersionSettingGetResponseObservabilityTracesPropagationPolicy = "authenticated"
	ScriptScriptAndVersionSettingGetResponseObservabilityTracesPropagationPolicyAccept        ScriptScriptAndVersionSettingGetResponseObservabilityTracesPropagationPolicy = "accept"
)

func (r ScriptScriptAndVersionSettingGetResponseObservabilityTracesPropagationPolicy) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseObservabilityTracesPropagationPolicyAuthenticated, ScriptScriptAndVersionSettingGetResponseObservabilityTracesPropagationPolicyAccept:
		return true
	}
	return false
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
type ScriptScriptAndVersionSettingGetResponsePlacement struct {
	// TCP host and port for targeted placement.
	Host string `json:"host"`
	// HTTP hostname for targeted placement.
	Hostname string `json:"hostname"`
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode ScriptScriptAndVersionSettingGetResponsePlacementMode `json:"mode"`
	// Cloud region for targeted placement in format 'provider:region'.
	Region string `json:"region"`
	// This field can have the runtime type of
	// [[]ScriptScriptAndVersionSettingGetResponsePlacementObjectTarget].
	Target interface{}                                           `json:"target"`
	JSON   scriptScriptAndVersionSettingGetResponsePlacementJSON `json:"-"`
	union  ScriptScriptAndVersionSettingGetResponsePlacementUnion
}

// scriptScriptAndVersionSettingGetResponsePlacementJSON contains the JSON metadata
// for the struct [ScriptScriptAndVersionSettingGetResponsePlacement]
type scriptScriptAndVersionSettingGetResponsePlacementJSON struct {
	Host        apijson.Field
	Hostname    apijson.Field
	Mode        apijson.Field
	Region      apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r scriptScriptAndVersionSettingGetResponsePlacementJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptScriptAndVersionSettingGetResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptScriptAndVersionSettingGetResponsePlacement{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptScriptAndVersionSettingGetResponsePlacementUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ScriptScriptAndVersionSettingGetResponsePlacementMode],
// [ScriptScriptAndVersionSettingGetResponsePlacementRegion],
// [ScriptScriptAndVersionSettingGetResponsePlacementHostname],
// [ScriptScriptAndVersionSettingGetResponsePlacementHost],
// [ScriptScriptAndVersionSettingGetResponsePlacementObject],
// [ScriptScriptAndVersionSettingGetResponsePlacementObject],
// [ScriptScriptAndVersionSettingGetResponsePlacementObject],
// [ScriptScriptAndVersionSettingGetResponsePlacementObject].
func (r ScriptScriptAndVersionSettingGetResponsePlacement) AsUnion() ScriptScriptAndVersionSettingGetResponsePlacementUnion {
	return r.union
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
//
// Union satisfied by [ScriptScriptAndVersionSettingGetResponsePlacementMode],
// [ScriptScriptAndVersionSettingGetResponsePlacementRegion],
// [ScriptScriptAndVersionSettingGetResponsePlacementHostname],
// [ScriptScriptAndVersionSettingGetResponsePlacementHost],
// [ScriptScriptAndVersionSettingGetResponsePlacementObject],
// [ScriptScriptAndVersionSettingGetResponsePlacementObject],
// [ScriptScriptAndVersionSettingGetResponsePlacementObject] or
// [ScriptScriptAndVersionSettingGetResponsePlacementObject].
type ScriptScriptAndVersionSettingGetResponsePlacementUnion interface {
	implementsScriptScriptAndVersionSettingGetResponsePlacement()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptScriptAndVersionSettingGetResponsePlacementUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptScriptAndVersionSettingGetResponsePlacementMode{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptScriptAndVersionSettingGetResponsePlacementRegion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptScriptAndVersionSettingGetResponsePlacementHostname{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptScriptAndVersionSettingGetResponsePlacementHost{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptScriptAndVersionSettingGetResponsePlacementObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptScriptAndVersionSettingGetResponsePlacementObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptScriptAndVersionSettingGetResponsePlacementObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptScriptAndVersionSettingGetResponsePlacementObject{}),
		},
	)
}

type ScriptScriptAndVersionSettingGetResponsePlacementMode struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode ScriptScriptAndVersionSettingGetResponsePlacementModeMode `json:"mode" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponsePlacementModeJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponsePlacementModeJSON contains the JSON
// metadata for the struct [ScriptScriptAndVersionSettingGetResponsePlacementMode]
type scriptScriptAndVersionSettingGetResponsePlacementModeJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponsePlacementMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponsePlacementModeJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponsePlacementMode) implementsScriptScriptAndVersionSettingGetResponsePlacement() {
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptScriptAndVersionSettingGetResponsePlacementModeMode string

const (
	ScriptScriptAndVersionSettingGetResponsePlacementModeModeSmart ScriptScriptAndVersionSettingGetResponsePlacementModeMode = "smart"
)

func (r ScriptScriptAndVersionSettingGetResponsePlacementModeMode) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponsePlacementModeModeSmart:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponsePlacementRegion struct {
	// Cloud region for targeted placement in format 'provider:region'.
	Region string                                                      `json:"region" api:"required"`
	JSON   scriptScriptAndVersionSettingGetResponsePlacementRegionJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponsePlacementRegionJSON contains the JSON
// metadata for the struct
// [ScriptScriptAndVersionSettingGetResponsePlacementRegion]
type scriptScriptAndVersionSettingGetResponsePlacementRegionJSON struct {
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponsePlacementRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponsePlacementRegionJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponsePlacementRegion) implementsScriptScriptAndVersionSettingGetResponsePlacement() {
}

type ScriptScriptAndVersionSettingGetResponsePlacementHostname struct {
	// HTTP hostname for targeted placement.
	Hostname string                                                        `json:"hostname" api:"required"`
	JSON     scriptScriptAndVersionSettingGetResponsePlacementHostnameJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponsePlacementHostnameJSON contains the JSON
// metadata for the struct
// [ScriptScriptAndVersionSettingGetResponsePlacementHostname]
type scriptScriptAndVersionSettingGetResponsePlacementHostnameJSON struct {
	Hostname    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponsePlacementHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponsePlacementHostnameJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponsePlacementHostname) implementsScriptScriptAndVersionSettingGetResponsePlacement() {
}

type ScriptScriptAndVersionSettingGetResponsePlacementHost struct {
	// TCP host and port for targeted placement.
	Host string                                                    `json:"host" api:"required"`
	JSON scriptScriptAndVersionSettingGetResponsePlacementHostJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponsePlacementHostJSON contains the JSON
// metadata for the struct [ScriptScriptAndVersionSettingGetResponsePlacementHost]
type scriptScriptAndVersionSettingGetResponsePlacementHostJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponsePlacementHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponsePlacementHostJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponsePlacementHost) implementsScriptScriptAndVersionSettingGetResponsePlacement() {
}

type ScriptScriptAndVersionSettingGetResponsePlacementObject struct {
	// Targeted placement mode.
	Mode ScriptScriptAndVersionSettingGetResponsePlacementObjectMode `json:"mode" api:"required"`
	// Cloud region for targeted placement in format 'provider:region'.
	Region string                                                      `json:"region" api:"required"`
	JSON   scriptScriptAndVersionSettingGetResponsePlacementObjectJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponsePlacementObjectJSON contains the JSON
// metadata for the struct
// [ScriptScriptAndVersionSettingGetResponsePlacementObject]
type scriptScriptAndVersionSettingGetResponsePlacementObjectJSON struct {
	Mode        apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponsePlacementObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponsePlacementObjectJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponsePlacementObject) implementsScriptScriptAndVersionSettingGetResponsePlacement() {
}

// Targeted placement mode.
type ScriptScriptAndVersionSettingGetResponsePlacementObjectMode string

const (
	ScriptScriptAndVersionSettingGetResponsePlacementObjectModeTargeted ScriptScriptAndVersionSettingGetResponsePlacementObjectMode = "targeted"
)

func (r ScriptScriptAndVersionSettingGetResponsePlacementObjectMode) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponsePlacementObjectModeTargeted:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type ScriptScriptAndVersionSettingGetResponseUsageModel string

const (
	ScriptScriptAndVersionSettingGetResponseUsageModelStandard ScriptScriptAndVersionSettingGetResponseUsageModel = "standard"
	ScriptScriptAndVersionSettingGetResponseUsageModelBundled  ScriptScriptAndVersionSettingGetResponseUsageModel = "bundled"
	ScriptScriptAndVersionSettingGetResponseUsageModelUnbound  ScriptScriptAndVersionSettingGetResponseUsageModel = "unbound"
)

func (r ScriptScriptAndVersionSettingGetResponseUsageModel) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseUsageModelStandard, ScriptScriptAndVersionSettingGetResponseUsageModelBundled, ScriptScriptAndVersionSettingGetResponseUsageModelUnbound:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParams struct {
	// Identifier.
	AccountID param.Field[string]                                          `path:"account_id" api:"required"`
	Settings  param.Field[ScriptScriptAndVersionSettingEditParamsSettings] `json:"settings"`
}

func (r ScriptScriptAndVersionSettingEditParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type ScriptScriptAndVersionSettingEditParamsSettings struct {
	// Annotations for the Worker version. Annotations are not inherited across
	// settings updates; omitting this field means the new version will have no
	// annotations.
	Annotations param.Field[ScriptScriptAndVersionSettingEditParamsSettingsAnnotations] `json:"annotations"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings param.Field[[]ScriptScriptAndVersionSettingEditParamsSettingsBindingUnion] `json:"bindings"`
	// Global CacheW configuration for the Worker. When caching is on, the platform
	// provisions a `cloudflare.app` zone for the Worker. A `type: worker` entry in the
	// `exports` map can override this value for a single entrypoint.
	CacheOptions param.Field[ScriptScriptAndVersionSettingEditParamsSettingsCacheOptions] `json:"cache_options"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Declarative exports for the Worker. Worker entrypoint entries (`type: worker`)
	// carry cache configuration for that entrypoint.
	Exports param.Field[map[string]ScriptScriptAndVersionSettingEditParamsSettingsExportsUnion] `json:"exports"`
	// Limits to apply for this Worker.
	Limits param.Field[ScriptScriptAndVersionSettingEditParamsSettingsLimits] `json:"limits"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[ScriptScriptAndVersionSettingEditParamsSettingsMigrationsUnion] `json:"migrations"`
	// Observability settings for the Worker.
	Observability param.Field[ScriptScriptAndVersionSettingEditParamsSettingsObservability] `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
	Placement param.Field[ScriptScriptAndVersionSettingEditParamsSettingsPlacementUnion] `json:"placement"`
	// Tags associated with the Worker.
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]ConsumerScriptParam] `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[ScriptScriptAndVersionSettingEditParamsSettingsUsageModel] `json:"usage_model"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Annotations for the Worker version. Annotations are not inherited across
// settings updates; omitting this field means the new version will have no
// annotations.
type ScriptScriptAndVersionSettingEditParamsSettingsAnnotations struct {
	// Human-readable message about the version. Truncated to 1000 bytes if longer.
	WorkersMessage param.Field[string] `json:"workers/message"`
	// User-provided identifier for the version. Maximum 100 bytes.
	WorkersTag param.Field[string] `json:"workers/tag"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsAnnotations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources.
type ScriptScriptAndVersionSettingEditParamsSettingsBinding struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsType] `json:"type" api:"required"`
	// Identifier of the D1 database to bind to.
	//
	// Deprecated: This property has been renamed to `database_id`.
	ID                          param.Field[string]      `json:"id"`
	Algorithm                   param.Field[interface{}] `json:"algorithm"`
	AllowedDestinationAddresses param.Field[interface{}] `json:"allowed_destination_addresses"`
	AllowedSenderAddresses      param.Field[interface{}] `json:"allowed_sender_addresses"`
	// ID of the Flagship app to bind to for feature flag evaluation.
	AppID param.Field[string] `json:"app_id"`
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name"`
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id"`
	// The exported class name of the Durable Object.
	ClassName param.Field[string] `json:"class_name"`
	// Identifier of the D1 database to bind to.
	DatabaseID param.Field[string] `json:"database_id"`
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset"`
	// Destination address for the email.
	DestinationAddress param.Field[string] `json:"destination_address" format:"email"`
	// The dispatch namespace the Durable Object script belongs to.
	DispatchNamespace param.Field[string] `json:"dispatch_namespace"`
	// Entrypoint to invoke on the target Worker.
	Entrypoint param.Field[string] `json:"entrypoint"`
	// The environment of the script_name to bind to.
	Environment param.Field[string] `json:"environment"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsFormat] `json:"format"`
	// Enables Gateway identity for the binding. Requires network_id to be
	// "cf1:network" and cannot be combined with tunnel_id.
	Identity param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsIdentity] `json:"identity"`
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name"`
	// The user-chosen instance name. Must exist at deploy time. The worker can search,
	// chat, update, and manage items/jobs on this instance.
	InstanceName param.Field[string]      `json:"instance_name"`
	Json         param.Field[interface{}] `json:"json"`
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsJurisdiction] `json:"jurisdiction"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string]      `json:"key_base64"`
	KeyJwk    param.Field[interface{}] `json:"key_jwk"`
	// The namespace the instance belongs to. Defaults to "default" if omitted.
	// Customers who don't use namespaces can simply omit this field.
	Namespace param.Field[string] `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id"`
	// Identifier of the network to bind to. Only "cf1:network" is currently supported.
	// Mutually exclusive with tunnel_id.
	NetworkID param.Field[string] `json:"network_id"`
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
	// Identifier of the VPC service to bind to.
	ServiceID param.Field[string]      `json:"service_id"`
	Simple    param.Field[interface{}] `json:"simple"`
	// ID of the store containing the secret.
	StoreID param.Field[string] `json:"store_id"`
	// The text value to use.
	Text param.Field[string] `json:"text"`
	// UUID of the Cloudflare Tunnel to bind to. Mutually exclusive with network_id.
	TunnelID param.Field[string]      `json:"tunnel_id"`
	Usages   param.Field[interface{}] `json:"usages"`
	// Identifier for the version to inherit the binding from, which can be the version
	// ID or the literal "latest" to inherit from the latest version. Defaults to
	// inheriting the binding from the latest version.
	VersionID param.Field[string] `json:"version_id"`
	// Name of the Workflow to bind to.
	WorkflowName param.Field[string] `json:"workflow_name"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBinding) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// A binding to allow the Worker to communicate with resources.
//
// Satisfied by
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAI],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearch],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearchNamespace],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMessaging],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngine],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssets],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowser],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDataBlob],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespace],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespace],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdrive],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindInherit],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindImages],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJson],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespace],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMedia],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificate],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainText],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPipelines],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueue],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindRatelimit],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2Bucket],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretText],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSendEmail],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindService],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTextBlob],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorize],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadata],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecret],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindFlagship],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKey],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWorkflow],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWasmModule],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCService],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCNetwork],
// [ScriptScriptAndVersionSettingEditParamsSettingsBinding].
type ScriptScriptAndVersionSettingEditParamsSettingsBindingUnion interface {
	implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion()
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAI) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeAI ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "ai"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearch struct {
	// The user-chosen instance name. Must exist at deploy time. The worker can search,
	// chat, update, and manage items/jobs on this instance.
	InstanceName param.Field[string] `json:"instance_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearchType] `json:"type" api:"required"`
	// The namespace the instance belongs to. Defaults to "default" if omitted.
	// Customers who don't use namespaces can simply omit this field.
	Namespace param.Field[string] `json:"namespace"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearch) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearchType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearchTypeAISearch ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearchType = "ai_search"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearchType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearchTypeAISearch:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearchNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The user-chosen namespace name. Must exist before deploy -- Wrangler handles
	// auto-creation on deploy failure (R2 bucket pattern). The "default" namespace is
	// auto-created by config-api for new accounts. Grants full access (CRUD + search +
	// chat) to all instances within the namespace.
	Namespace param.Field[string] `json:"namespace" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearchNamespaceType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearchNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearchNamespace) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearchNamespaceType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearchNamespaceTypeAISearchNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearchNamespaceType = "ai_search_namespace"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAISearchNamespaceTypeAISearchNamespace:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMessaging struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The Messaging namespace to bind to.
	Namespace param.Field[string] `json:"namespace" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMessagingType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMessaging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMessaging) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMessagingType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMessagingTypeMessaging ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMessagingType = "messaging"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMessagingType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMessagingTypeMessaging:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngine) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngine) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssets) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeAssets ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "assets"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowser) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserTypeBrowser ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserType = "browser"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	DatabaseID param.Field[string] `json:"database_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type] `json:"type" api:"required"`
	// Identifier of the D1 database to bind to.
	//
	// Deprecated: This property has been renamed to `database_id`.
	ID param.Field[string] `json:"id"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeD1 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "d1"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDataBlob struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The name of the file containing the data content. Only accepted for
	// `service worker syntax` Workers.
	Part param.Field[string] `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDataBlobType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDataBlob) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDataBlob) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDataBlobType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDataBlobTypeDataBlob ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDataBlobType = "data_blob"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDataBlobType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDataBlobTypeDataBlob:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The name of the dispatch namespace.
	Namespace param.Field[string] `json:"namespace" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType] `json:"type" api:"required"`
	// Outbound worker.
	Outbound param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutbound] `json:"outbound"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespace) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params param.Field[[]ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutboundParam] `json:"params"`
	// Outbound worker.
	Worker param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutboundWorker] `json:"worker"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutboundParam struct {
	// Name of the parameter.
	Name param.Field[string] `json:"name" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutboundParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Entrypoint to invoke on the outbound worker.
	Entrypoint param.Field[string] `json:"entrypoint"`
	// Environment of the outbound worker.
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker.
	Service param.Field[string] `json:"service"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType] `json:"type" api:"required"`
	// The exported class name of the Durable Object.
	ClassName param.Field[string] `json:"class_name"`
	// The dispatch namespace the Durable Object script belongs to.
	DispatchNamespace param.Field[string] `json:"dispatch_namespace"`
	// The environment of the script_name to bind to.
	Environment param.Field[string] `json:"environment"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespace) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID param.Field[string] `json:"id" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdrive) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdrive) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeHyperdrive ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindInherit struct {
	// The name of the inherited binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindInheritType] `json:"type" api:"required"`
	// The old name of the inherited binding. If set, the binding will be renamed from
	// `old_name` to `name` in the new version. If not set, the binding will keep the
	// same name between versions.
	OldName param.Field[string] `json:"old_name"`
	// Identifier for the version to inherit the binding from, which can be the version
	// ID or the literal "latest" to inherit from the latest version. Defaults to
	// inheriting the binding from the latest version.
	VersionID param.Field[string] `json:"version_id"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindInherit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindInherit) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindInheritType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindInheritTypeInherit ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindInheritType = "inherit"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindInheritType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindInheritTypeInherit:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindImages struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindImagesType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindImages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindImages) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindImagesType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindImagesTypeImages ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindImagesType = "images"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindImagesType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindImagesTypeImages:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json param.Field[interface{}] `json:"json" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJson) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeJson ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "json"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespace) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeKVNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMedia struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMediaType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMedia) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMedia) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMediaType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMediaTypeMedia ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMediaType = "media"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMediaType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMediaTypeMedia:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificate) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The text value to use.
	Text param.Field[string] `json:"text" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainText) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypePlainText ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Name of the Pipeline to bind to.
	Pipeline param.Field[string] `json:"pipeline" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPipelinesType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPipelines) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPipelines) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPipelinesType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPipelinesTypePipelines ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPipelinesType = "pipelines"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueue) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeQueue ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "queue"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindRatelimit struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Identifier of the rate limit namespace to bind to.
	NamespaceID param.Field[string] `json:"namespace_id" api:"required"`
	// The rate limit configuration.
	Simple param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindRatelimitSimple] `json:"simple" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindRatelimitType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindRatelimit) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The rate limit configuration.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindRatelimitSimple struct {
	// The limit (requests per period).
	Limit param.Field[float64] `json:"limit" api:"required"`
	// The period in seconds.
	Period param.Field[int64] `json:"period" api:"required"`
	// Duration in seconds to apply the mitigation action after the rate limit is
	// exceeded. Valid values are 0 (disabled), 10, or multiples of 60 up to 86400.
	// Must be greater than or equal to the period when non-zero.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindRatelimitSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindRatelimitType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindRatelimitTypeRatelimit ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindRatelimitType = "ratelimit"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindRatelimitType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindRatelimitTypeRatelimit:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType] `json:"type" api:"required"`
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketJurisdiction] `json:"jurisdiction"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2Bucket) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2Bucket) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeR2Bucket ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

// The
// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
// of the R2 bucket.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketJurisdiction string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketJurisdictionEu          ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketJurisdiction = "eu"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketJurisdictionFedramp     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketJurisdiction = "fedramp"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketJurisdictionFedrampHigh ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketJurisdiction = "fedramp-high"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketJurisdictionUs          ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketJurisdiction = "us"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketJurisdiction) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketJurisdictionEu, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketJurisdictionFedramp, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketJurisdictionFedrampHigh, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketJurisdictionUs:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The secret value to use.
	Text param.Field[string] `json:"text" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretText) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeSecretText ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "secret_text"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSendEmail struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSendEmailType] `json:"type" api:"required"`
	// List of allowed destination addresses.
	AllowedDestinationAddresses param.Field[[]string] `json:"allowed_destination_addresses" format:"email"`
	// List of allowed sender addresses.
	AllowedSenderAddresses param.Field[[]string] `json:"allowed_sender_addresses" format:"email"`
	// Destination address for the email.
	DestinationAddress param.Field[string] `json:"destination_address" format:"email"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSendEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSendEmail) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSendEmailType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSendEmailTypeSendEmail ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSendEmailType = "send_email"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSendEmailType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSendEmailTypeSendEmail:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindService struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType] `json:"type" api:"required"`
	// Entrypoint to invoke on the target Worker.
	Entrypoint param.Field[string] `json:"entrypoint"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindService) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeService ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "service"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTextBlob struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The name of the file containing the text content. Only accepted for
	// `service worker syntax` Workers.
	Part param.Field[string] `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTextBlobType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTextBlob) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTextBlob) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTextBlobType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTextBlobTypeTextBlob ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTextBlobType = "text_blob"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTextBlobType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTextBlobTypeTextBlob:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorize) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeVectorize ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "vectorize"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadata) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Name of the secret in the store.
	SecretName param.Field[string] `json:"secret_name" api:"required"`
	// ID of the store containing the secret.
	StoreID param.Field[string] `json:"store_id" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecretType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecret) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecretType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindFlagship struct {
	// ID of the Flagship app to bind to for feature flag evaluation.
	AppID param.Field[string] `json:"app_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindFlagshipType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindFlagship) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindFlagship) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindFlagshipType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindFlagshipTypeFlagship ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindFlagshipType = "flagship"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindFlagshipType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindFlagshipTypeFlagship:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm param.Field[interface{}] `json:"algorithm" api:"required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormat] `json:"format" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyType] `json:"type" api:"required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages param.Field[[]ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage] `json:"usages" api:"required"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string] `json:"key_base64"`
	// Key data in
	// [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key)
	// format. Required if `format` is "jwk".
	KeyJwk param.Field[interface{}] `json:"key_jwk"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKey) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormat string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormatRaw   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormat = "raw"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormatPkcs8 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormat = "pkcs8"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormatSpki  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormat = "spki"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormatJwk   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormatRaw, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormatPkcs8, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormatSpki, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyTypeSecretKey ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyType = "secret_key"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageEncrypt    ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage = "encrypt"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageDecrypt    ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage = "decrypt"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageSign       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage = "sign"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageVerify     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage = "verify"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageDeriveKey  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage = "deriveKey"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageDeriveBits ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage = "deriveBits"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageWrapKey    ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage = "wrapKey"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageUnwrapKey  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageEncrypt, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageDecrypt, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageSign, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageVerify, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageDeriveKey, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageDeriveBits, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageWrapKey, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWorkflowType] `json:"type" api:"required"`
	// Name of the Workflow to bind to.
	WorkflowName param.Field[string] `json:"workflow_name" api:"required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName param.Field[string] `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWorkflow) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWorkflow) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWorkflowType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWorkflowTypeWorkflow ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWorkflowType = "workflow"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWasmModule struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The name of the file containing the WebAssembly module content. Only accepted
	// for `service worker syntax` Workers.
	Part param.Field[string] `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWasmModuleType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWasmModule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWasmModule) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWasmModuleType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWasmModuleTypeWasmModule ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWasmModuleType = "wasm_module"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWasmModuleType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindWasmModuleTypeWasmModule:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCService struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Identifier of the VPC service to bind to.
	ServiceID param.Field[string] `json:"service_id" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCServiceType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCService) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCServiceType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCServiceTypeVPCService ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCServiceType = "vpc_service"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCServiceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCServiceTypeVPCService:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCNetwork struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCNetworkType] `json:"type" api:"required"`
	// Enables Gateway identity for the binding. Requires network_id to be
	// "cf1:network" and cannot be combined with tunnel_id.
	Identity param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCNetworkIdentity] `json:"identity"`
	// Identifier of the network to bind to. Only "cf1:network" is currently supported.
	// Mutually exclusive with tunnel_id.
	NetworkID param.Field[string] `json:"network_id"`
	// UUID of the Cloudflare Tunnel to bind to. Mutually exclusive with network_id.
	TunnelID param.Field[string] `json:"tunnel_id"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCNetwork) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCNetworkType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCNetworkTypeVPCNetwork ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCNetworkType = "vpc_network"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCNetworkType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCNetworkTypeVPCNetwork:
		return true
	}
	return false
}

// Enables Gateway identity for the binding. Requires network_id to be
// "cf1:network" and cannot be combined with tunnel_id.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCNetworkIdentity string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCNetworkIdentityRuntimeEmailAlpha ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCNetworkIdentity = "runtime-email-alpha"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCNetworkIdentity) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVPCNetworkIdentityRuntimeEmailAlpha:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeAISearch               ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "ai_search"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeAISearchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "ai_search_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeMessaging              ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "messaging"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeBrowser                ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "browser"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeDataBlob               ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "data_blob"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeInherit                ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "inherit"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeImages                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "images"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeMedia                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "media"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypePipelines              ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "pipelines"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeRatelimit              ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "ratelimit"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeSendEmail              ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "send_email"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeTextBlob               ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "text_blob"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "version_metadata"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeSecretsStoreSecret     ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "secrets_store_secret"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeFlagship               ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "flagship"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeSecretKey              ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "secret_key"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeWorkflow               ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "workflow"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeWasmModule             ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "wasm_module"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeVPCService             ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "vpc_service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeVPCNetwork             ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "vpc_network"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeAISearch, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeAISearchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeMessaging, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeBrowser, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeDataBlob, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeInherit, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeImages, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeMedia, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypePipelines, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeRatelimit, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeSendEmail, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeTextBlob, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeVersionMetadata, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeSecretsStoreSecret, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeFlagship, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeSecretKey, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeWorkflow, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeWasmModule, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeVPCService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeVPCNetwork:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsFormat string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsFormatRaw   ScriptScriptAndVersionSettingEditParamsSettingsBindingsFormat = "raw"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsFormatPkcs8 ScriptScriptAndVersionSettingEditParamsSettingsBindingsFormat = "pkcs8"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsFormatSpki  ScriptScriptAndVersionSettingEditParamsSettingsBindingsFormat = "spki"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsFormatJwk   ScriptScriptAndVersionSettingEditParamsSettingsBindingsFormat = "jwk"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsFormat) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsFormatRaw, ScriptScriptAndVersionSettingEditParamsSettingsBindingsFormatPkcs8, ScriptScriptAndVersionSettingEditParamsSettingsBindingsFormatSpki, ScriptScriptAndVersionSettingEditParamsSettingsBindingsFormatJwk:
		return true
	}
	return false
}

// Enables Gateway identity for the binding. Requires network_id to be
// "cf1:network" and cannot be combined with tunnel_id.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsIdentity string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsIdentityRuntimeEmailAlpha ScriptScriptAndVersionSettingEditParamsSettingsBindingsIdentity = "runtime-email-alpha"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsIdentity) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsIdentityRuntimeEmailAlpha:
		return true
	}
	return false
}

// The
// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
// of the R2 bucket.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsJurisdiction string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsJurisdictionEu          ScriptScriptAndVersionSettingEditParamsSettingsBindingsJurisdiction = "eu"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsJurisdictionFedramp     ScriptScriptAndVersionSettingEditParamsSettingsBindingsJurisdiction = "fedramp"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsJurisdictionFedrampHigh ScriptScriptAndVersionSettingEditParamsSettingsBindingsJurisdiction = "fedramp-high"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsJurisdictionUs          ScriptScriptAndVersionSettingEditParamsSettingsBindingsJurisdiction = "us"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsJurisdiction) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsJurisdictionEu, ScriptScriptAndVersionSettingEditParamsSettingsBindingsJurisdictionFedramp, ScriptScriptAndVersionSettingEditParamsSettingsBindingsJurisdictionFedrampHigh, ScriptScriptAndVersionSettingEditParamsSettingsBindingsJurisdictionUs:
		return true
	}
	return false
}

// Global CacheW configuration for the Worker. When caching is on, the platform
// provisions a `cloudflare.app` zone for the Worker. A `type: worker` entry in the
// `exports` map can override this value for a single entrypoint.
type ScriptScriptAndVersionSettingEditParamsSettingsCacheOptions struct {
	// Whether caching is enabled for this Worker.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Whether cached responses are shared across Worker version uploads. This is
	// independent of `enabled`. It can stay true while caching is off, so the
	// preference survives turning caching off and back on.
	CrossVersionCache param.Field[bool] `json:"cross_version_cache"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsCacheOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A single entry in the `exports` map, keyed by export name (a `WorkerEntrypoint`
// class name, a Durable Object class name, or `default` for the Worker's default
// export). The `type` discriminator selects the top-level shape: `worker`
// entrypoint entries may carry `cache` configuration, while `durable-object`
// entries are further refined by the optional `state` field (default `created`).
// Tombstone states (`deleted`, `renamed`, `transferred`) express destructive
// lifecycle operations declaratively; `expecting-transfer` is the live target side
// of a transfer. The server validates the exact per-(type, state) field
// combinations; fields not listed for a variant are rejected.
type ScriptScriptAndVersionSettingEditParamsSettingsExports struct {
	// Marks this entry as a Worker entrypoint export.
	Type  param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsType] `json:"type" api:"required"`
	Cache param.Field[interface{}]                                                `json:"cache"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container param.Field[string] `json:"container"`
	// The destination class name. Must differ from the source class (the map key) and
	// must be declared as a live (`created`) entry in the same `exports` map.
	// Write-only: never present in GET responses.
	RenamedTo param.Field[string] `json:"renamed_to"`
	// Live export. May be omitted; defaults to `created`.
	State param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsState] `json:"state"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsStorage] `json:"storage"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom param.Field[string] `json:"transfer_from"`
	// The destination script name. Must be in the same account and the same
	// dispatch-namespace context (or both non-dispatch). Cross-dispatch-namespace
	// transfers are rejected. Write-only: never present in GET responses.
	TransferredTo param.Field[string] `json:"transferred_to"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExports) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExports) implementsScriptScriptAndVersionSettingEditParamsSettingsExportsUnion() {
}

// A single entry in the `exports` map, keyed by export name (a `WorkerEntrypoint`
// class name, a Durable Object class name, or `default` for the Worker's default
// export). The `type` discriminator selects the top-level shape: `worker`
// entrypoint entries may carry `cache` configuration, while `durable-object`
// entries are further refined by the optional `state` field (default `created`).
// Tombstone states (`deleted`, `renamed`, `transferred`) express destructive
// lifecycle operations declaratively; `expecting-transfer` is the live target side
// of a transfer. The server validates the exact per-(type, state) field
// combinations; fields not listed for a variant are rejected.
//
// Satisfied by
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersWorkerExport],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExport],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectDeletedExport],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectRenamedExport],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectTransferredExport],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExport],
// [ScriptScriptAndVersionSettingEditParamsSettingsExports].
type ScriptScriptAndVersionSettingEditParamsSettingsExportsUnion interface {
	implementsScriptScriptAndVersionSettingEditParamsSettingsExportsUnion()
}

// A named Worker entrypoint export (`type: worker`). Worker entrypoints are always
// live (`state: created`) and carry no storage or lifecycle fields. The optional
// `cache` block overrides the Worker's global `cache_options.enabled` for this
// entrypoint.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersWorkerExport struct {
	// Marks this entry as a Worker entrypoint export.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersWorkerExportType] `json:"type" api:"required"`
	// Cache override for this entrypoint. Overrides the Worker's global
	// `cache_options.enabled` for this entrypoint only.
	Cache param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersWorkerExportCache] `json:"cache"`
	// Live export. May be omitted; defaults to `created`.
	State param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersWorkerExportState] `json:"state"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersWorkerExport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersWorkerExport) implementsScriptScriptAndVersionSettingEditParamsSettingsExportsUnion() {
}

// Marks this entry as a Worker entrypoint export.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersWorkerExportType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersWorkerExportTypeWorker ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersWorkerExportType = "worker"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersWorkerExportType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersWorkerExportTypeWorker:
		return true
	}
	return false
}

// Cache override for this entrypoint. Overrides the Worker's global
// `cache_options.enabled` for this entrypoint only.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersWorkerExportCache struct {
	// Whether caching is enabled for this entrypoint.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersWorkerExportCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Live export. May be omitted; defaults to `created`.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersWorkerExportState string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersWorkerExportStateCreated ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersWorkerExportState = "created"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersWorkerExportState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersWorkerExportStateCreated:
		return true
	}
	return false
}

// A live Durable Object export (`state: created`, the default). The platform
// auto-provisions the namespace on first deploy, matches it on subsequent deploys,
// and never mutates or deletes it as a side effect of a code-only change.
// `storage` is required; `renamed_to`, `transferred_to` and `transfer_from` are
// not allowed on a live entry.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExport struct {
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportStorage] `json:"storage" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportType] `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container param.Field[string] `json:"container"`
	// Live export. May be omitted; defaults to `created`.
	State param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportState] `json:"state"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExport) implementsScriptScriptAndVersionSettingEditParamsSettingsExportsUnion() {
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportStorage string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportStorageSqlite   ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportStorage = "sqlite"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportStorageLegacyKV ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportStorage = "legacy-kv"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportStorage) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportStorageSqlite, ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportStorageLegacyKV:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportTypeDurableObject ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportType = "durable-object"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportTypeDurableObject:
		return true
	}
	return false
}

// Live export. May be omitted; defaults to `created`.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportState string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportStateCreated ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportState = "created"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExportStateCreated:
		return true
	}
	return false
}

// A `deleted` tombstone: retires the provisioned namespace for this class and all
// of its data. The class must be absent from the uploaded code and no other Worker
// in the account may bind to the namespace, otherwise the deploy is rejected. No
// other fields are allowed. Deletion is irreversible.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectDeletedExport struct {
	// Tombstone that deletes the namespace.
	State param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectDeletedExportState] `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectDeletedExportType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectDeletedExport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectDeletedExport) implementsScriptScriptAndVersionSettingEditParamsSettingsExportsUnion() {
}

// Tombstone that deletes the namespace.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectDeletedExportState string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectDeletedExportStateDeleted ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectDeletedExportState = "deleted"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectDeletedExportState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectDeletedExportStateDeleted:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectDeletedExportType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectDeletedExportTypeDurableObject ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectDeletedExportType = "durable-object"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectDeletedExportType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectDeletedExportTypeDurableObject:
		return true
	}
	return false
}

// A `renamed` tombstone: rewrites the provisioned namespace's class name from this
// map key to `renamed_to`. The source class may stay in code during the rollout
// window (an info notice is emitted). `storage`, `transferred_to` and
// `transfer_from` are not allowed.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectRenamedExport struct {
	// The destination class name. Must differ from the source class (the map key) and
	// must be declared as a live (`created`) entry in the same `exports` map.
	// Write-only: never present in GET responses.
	RenamedTo param.Field[string] `json:"renamed_to" api:"required"`
	// Tombstone that renames the namespace's class.
	State param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectRenamedExportState] `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectRenamedExportType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectRenamedExport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectRenamedExport) implementsScriptScriptAndVersionSettingEditParamsSettingsExportsUnion() {
}

// Tombstone that renames the namespace's class.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectRenamedExportState string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectRenamedExportStateRenamed ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectRenamedExportState = "renamed"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectRenamedExportState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectRenamedExportStateRenamed:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectRenamedExportType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectRenamedExportTypeDurableObject ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectRenamedExportType = "durable-object"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectRenamedExportType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectRenamedExportTypeDurableObject:
		return true
	}
	return false
}

// A `transferred` tombstone (source side of a two-phase transfer): hands ownership
// of the provisioned namespace to another script in the same account, named by
// `transferred_to`. The target must have already deployed a matching
// `expecting-transfer` entry. The source class may stay in code during the rollout
// window (an info notice is emitted). `storage`, `renamed_to` and `transfer_from`
// are not allowed.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectTransferredExport struct {
	// Tombstone that transfers the namespace to another script.
	State param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectTransferredExportState] `json:"state" api:"required"`
	// The destination script name. Must be in the same account and the same
	// dispatch-namespace context (or both non-dispatch). Cross-dispatch-namespace
	// transfers are rejected. Write-only: never present in GET responses.
	TransferredTo param.Field[string] `json:"transferred_to" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectTransferredExportType] `json:"type" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectTransferredExport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectTransferredExport) implementsScriptScriptAndVersionSettingEditParamsSettingsExportsUnion() {
}

// Tombstone that transfers the namespace to another script.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectTransferredExportState string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectTransferredExportStateTransferred ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectTransferredExportState = "transferred"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectTransferredExportState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectTransferredExportStateTransferred:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectTransferredExportType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectTransferredExportTypeDurableObject ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectTransferredExportType = "durable-object"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectTransferredExportType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectTransferredExportTypeDurableObject:
		return true
	}
	return false
}

// The target side of a two-phase transfer (`state: expecting-transfer`). Declares
// that this script expects to receive a namespace for this class from the
// `transfer_from` script. This is a live entry, not a tombstone: bindings resolve
// through the source's namespace until the source commits with a `transferred`
// tombstone. `storage` and `transfer_from` are required; `renamed_to` and
// `transferred_to` are not allowed.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExport struct {
	// Target side of a two-phase transfer.
	State param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportState] `json:"state" api:"required"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportStorage] `json:"storage" api:"required"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom param.Field[string] `json:"transfer_from" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportType] `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object once the transfer settles. Valid only on live entries.
	Container param.Field[string] `json:"container"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExport) implementsScriptScriptAndVersionSettingEditParamsSettingsExportsUnion() {
}

// Target side of a two-phase transfer.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportState string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportStateExpectingTransfer ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportState = "expecting-transfer"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportStateExpectingTransfer:
		return true
	}
	return false
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportStorage string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportStorageSqlite   ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportStorage = "sqlite"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportStorageLegacyKV ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportStorage = "legacy-kv"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportStorage) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportStorageSqlite, ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportStorageLegacyKV:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportTypeDurableObject ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportType = "durable-object"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsWorkersDurableObjectExpectingTransferExportTypeDurableObject:
		return true
	}
	return false
}

// Marks this entry as a Worker entrypoint export.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsTypeWorker        ScriptScriptAndVersionSettingEditParamsSettingsExportsType = "worker"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsTypeDurableObject ScriptScriptAndVersionSettingEditParamsSettingsExportsType = "durable-object"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsTypeWorker, ScriptScriptAndVersionSettingEditParamsSettingsExportsTypeDurableObject:
		return true
	}
	return false
}

// Live export. May be omitted; defaults to `created`.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsState string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsStateCreated           ScriptScriptAndVersionSettingEditParamsSettingsExportsState = "created"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsStateDeleted           ScriptScriptAndVersionSettingEditParamsSettingsExportsState = "deleted"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsStateRenamed           ScriptScriptAndVersionSettingEditParamsSettingsExportsState = "renamed"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsStateTransferred       ScriptScriptAndVersionSettingEditParamsSettingsExportsState = "transferred"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsStateExpectingTransfer ScriptScriptAndVersionSettingEditParamsSettingsExportsState = "expecting-transfer"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsState) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsStateCreated, ScriptScriptAndVersionSettingEditParamsSettingsExportsStateDeleted, ScriptScriptAndVersionSettingEditParamsSettingsExportsStateRenamed, ScriptScriptAndVersionSettingEditParamsSettingsExportsStateTransferred, ScriptScriptAndVersionSettingEditParamsSettingsExportsStateExpectingTransfer:
		return true
	}
	return false
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsStorage string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsStorageSqlite   ScriptScriptAndVersionSettingEditParamsSettingsExportsStorage = "sqlite"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsStorageLegacyKV ScriptScriptAndVersionSettingEditParamsSettingsExportsStorage = "legacy-kv"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsStorage) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsStorageSqlite, ScriptScriptAndVersionSettingEditParamsSettingsExportsStorageLegacyKV:
		return true
	}
	return false
}

// Summary of the declarative exports reconciliation that ran on this upload.
// Populated only when the uploaded metadata included an `exports` block. Durable
// Object entries drive reconciliation; `type: worker` entries do not contribute to
// this summary.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliation struct {
	// Class names for which a new namespace was provisioned.
	Created param.Field[[]string] `json:"created" api:"required"`
	// Class names whose namespace was deleted by a `deleted` tombstone.
	Deleted param.Field[[]string] `json:"deleted" api:"required"`
	// Non-blocking info entries (stale tombstones, tombstone applied with class still
	// in code). See `exports_reconciliation_info`.
	Info param.Field[[]ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfo] `json:"info" api:"required"`
	// Source class names whose tombstone entry is now stale and safe to delete from
	// `exports` (no remaining referencing scripts).
	RemovableEntries param.Field[[]string] `json:"removable_entries" api:"required"`
	// Applied `renamed` tombstones.
	Renamed param.Field[[]ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationRenamed] `json:"renamed" api:"required"`
	// Phase-1 transfer hints recorded on the target side.
	TransferPending param.Field[[]ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationTransferPending] `json:"transfer_pending" api:"required"`
	// Committed `transferred` tombstones (phase-2).
	Transferred param.Field[[]ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationTransferred] `json:"transferred" api:"required"`
	// Class names whose provisioned namespace was mutated in place.
	Updated param.Field[[]string] `json:"updated" api:"required"`
	// Non-blocking warnings. See `exports_reconciliation_warning`.
	Warnings param.Field[[]ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarning] `json:"warnings" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A non-blocking reconciliation info entry. Emitted for stale tombstones (a no-op
// on this deploy) and for tombstones applied with the source class still in code
// (the supported zero-downtime rollout pattern).
type ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfo struct {
	// The class name the info entry is about.
	Class param.Field[string] `json:"class" api:"required"`
	// Human-readable explanation.
	Message param.Field[string] `json:"message" api:"required"`
	// Stable, machine-readable tag identifying which reconciliation scenario produced
	// an error, warning, or info entry. Clients may branch on this value instead of
	// parsing `message`.
	Scenario param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario] `json:"scenario" api:"required"`
	// The provisioned namespace the entry relates to, when applicable.
	NamespaceID param.Field[string] `json:"namespace_id" format:"uuid"`
	// Other Workers in the account that still bind to the affected class. Advisory:
	// while non-empty the tombstone is not yet safe to remove — redeploy these Workers
	// with bindings re-pointed first.
	ReferencingScripts param.Field[[]string] `json:"referencing_scripts"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Stable, machine-readable tag identifying which reconciliation scenario produced
// an error, warning, or info entry. Clients may branch on this value instead of
// parsing `message`.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioCodeClassNotInExports                     ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "code_class_not_in_exports"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioProvisionedClassMissingFromConfig         ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "provisioned_class_missing_from_config"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioConfigExportNotInCode                     ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "config_export_not_in_code"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioConfigReferencesNonexistentClass          ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "config_references_nonexistent_class"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioOrphanedProvisionedNamespace              ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "orphaned_provisioned_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioStorageTypeMismatch                       ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "storage_type_mismatch"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioFreeTierRequiresSqlite                    ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "free_tier_requires_sqlite"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioInvalidExport                             ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "invalid_export"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTombstoneDeleteClassStillInCode           ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "tombstone_delete_class_still_in_code"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTombstoneDeleteBlockedByExternalBindings  ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "tombstone_delete_blocked_by_external_bindings"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTombstoneRenamedToOccupied                ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "tombstone_renamed_to_occupied"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTransferredPendingNotFound                ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "transferred_pending_not_found"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTransferredTargetMissing                  ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "transferred_target_missing"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTransferredTargetMismatch                 ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "transferred_target_mismatch"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioPhaseOneTransferSourceMissing             ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "phase_one_transfer_source_missing"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioPhaseOneTransferSourceNamespaceMissing    ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "phase_one_transfer_source_namespace_missing"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioPhaseOneTransferTargetClassProvisioned    ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "phase_one_transfer_target_class_provisioned"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioPhaseOneTransferAfterCommitMismatch       ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "phase_one_transfer_after_commit_mismatch"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioPhaseOneTransferDuplicate                 ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "phase_one_transfer_duplicate"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioPhaseOneTransferTargetInDispatchNamespace ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "phase_one_transfer_target_in_dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioPhaseOneTransferSourceInDispatchNamespace ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "phase_one_transfer_source_in_dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTransferredSourceInDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "transferred_source_in_dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTransferredTargetInDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "transferred_target_in_dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioContainerUndeclaredReference              ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "container_undeclared_reference"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioContainerClassNotDurableObject            ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "container_class_not_durable_object"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioContainerWiringInconsistent               ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "container_wiring_inconsistent"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioContainerMultipleDurableObjects           ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "container_multiple_durable_objects"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTransferContainerParityMismatch           ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "transfer_container_parity_mismatch"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTransferContainerParityMismatchOnCommit   ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "transfer_container_parity_mismatch_on_commit"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTombstoneClassStillInCode                 ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "tombstone_class_still_in_code"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioStaleTombstone                            ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "stale_tombstone"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTransferReceiveAlreadyApplied             ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "transfer_receive_already_applied"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTransferReceiveCleanupComplete            ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario = "transfer_receive_cleanup_complete"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenario) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioCodeClassNotInExports, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioProvisionedClassMissingFromConfig, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioConfigExportNotInCode, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioConfigReferencesNonexistentClass, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioOrphanedProvisionedNamespace, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioStorageTypeMismatch, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioFreeTierRequiresSqlite, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioInvalidExport, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTombstoneDeleteClassStillInCode, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTombstoneDeleteBlockedByExternalBindings, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTombstoneRenamedToOccupied, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTransferredPendingNotFound, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTransferredTargetMissing, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTransferredTargetMismatch, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioPhaseOneTransferSourceMissing, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioPhaseOneTransferSourceNamespaceMissing, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioPhaseOneTransferTargetClassProvisioned, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioPhaseOneTransferAfterCommitMismatch, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioPhaseOneTransferDuplicate, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioPhaseOneTransferTargetInDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioPhaseOneTransferSourceInDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTransferredSourceInDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTransferredTargetInDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioContainerUndeclaredReference, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioContainerClassNotDurableObject, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioContainerWiringInconsistent, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioContainerMultipleDurableObjects, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTransferContainerParityMismatch, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTransferContainerParityMismatchOnCommit, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTombstoneClassStillInCode, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioStaleTombstone, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTransferReceiveAlreadyApplied, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationInfoScenarioTransferReceiveCleanupComplete:
		return true
	}
	return false
}

// A single applied `renamed` tombstone.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationRenamed struct {
	// The original (source) class name.
	From param.Field[string] `json:"from" api:"required"`
	// The new class name (`renamed_to`).
	To param.Field[string] `json:"to" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationRenamed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A single phase-1 transfer hint recorded on the target side (a live
// `expecting-transfer` entry).
type ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationTransferPending struct {
	// The target-side class name awaiting transfer.
	Class param.Field[string] `json:"class" api:"required"`
	// The source script the namespace will be transferred from.
	From param.Field[string] `json:"from" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationTransferPending) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A single committed `transferred` tombstone (phase-2 commit).
type ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationTransferred struct {
	// The source class name that was transferred.
	Class param.Field[string] `json:"class" api:"required"`
	// The transfer phase. Currently always `committed`.
	Phase param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationTransferredPhase] `json:"phase" api:"required"`
	// The destination script that now owns the namespace.
	To param.Field[string] `json:"to" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationTransferred) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The transfer phase. Currently always `committed`.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationTransferredPhase string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationTransferredPhaseCommitted ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationTransferredPhase = "committed"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationTransferredPhase) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationTransferredPhaseCommitted:
		return true
	}
	return false
}

// A non-blocking reconciliation warning. Reserved: no scenario populates this
// array today (`code_class_not_in_exports` is surfaced as info and
// `provisioned_class_missing_from_config` is a hard error). Clients should still
// surface any entries that appear.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarning struct {
	// The class name the warning is about.
	Class param.Field[string] `json:"class" api:"required"`
	// Human-readable explanation of the warning.
	Message param.Field[string] `json:"message" api:"required"`
	// Stable, machine-readable tag identifying which reconciliation scenario produced
	// an error, warning, or info entry. Clients may branch on this value instead of
	// parsing `message`.
	Scenario param.Field[ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario] `json:"scenario" api:"required"`
	// The provisioned namespace the warning relates to, when applicable.
	NamespaceID param.Field[string] `json:"namespace_id" format:"uuid"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Stable, machine-readable tag identifying which reconciliation scenario produced
// an error, warning, or info entry. Clients may branch on this value instead of
// parsing `message`.
type ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioCodeClassNotInExports                     ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "code_class_not_in_exports"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioProvisionedClassMissingFromConfig         ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "provisioned_class_missing_from_config"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioConfigExportNotInCode                     ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "config_export_not_in_code"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioConfigReferencesNonexistentClass          ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "config_references_nonexistent_class"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioOrphanedProvisionedNamespace              ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "orphaned_provisioned_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioStorageTypeMismatch                       ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "storage_type_mismatch"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioFreeTierRequiresSqlite                    ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "free_tier_requires_sqlite"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioInvalidExport                             ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "invalid_export"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTombstoneDeleteClassStillInCode           ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "tombstone_delete_class_still_in_code"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTombstoneDeleteBlockedByExternalBindings  ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "tombstone_delete_blocked_by_external_bindings"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTombstoneRenamedToOccupied                ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "tombstone_renamed_to_occupied"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTransferredPendingNotFound                ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "transferred_pending_not_found"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTransferredTargetMissing                  ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "transferred_target_missing"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTransferredTargetMismatch                 ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "transferred_target_mismatch"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioPhaseOneTransferSourceMissing             ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "phase_one_transfer_source_missing"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioPhaseOneTransferSourceNamespaceMissing    ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "phase_one_transfer_source_namespace_missing"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioPhaseOneTransferTargetClassProvisioned    ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "phase_one_transfer_target_class_provisioned"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioPhaseOneTransferAfterCommitMismatch       ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "phase_one_transfer_after_commit_mismatch"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioPhaseOneTransferDuplicate                 ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "phase_one_transfer_duplicate"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioPhaseOneTransferTargetInDispatchNamespace ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "phase_one_transfer_target_in_dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioPhaseOneTransferSourceInDispatchNamespace ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "phase_one_transfer_source_in_dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTransferredSourceInDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "transferred_source_in_dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTransferredTargetInDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "transferred_target_in_dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioContainerUndeclaredReference              ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "container_undeclared_reference"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioContainerClassNotDurableObject            ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "container_class_not_durable_object"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioContainerWiringInconsistent               ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "container_wiring_inconsistent"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioContainerMultipleDurableObjects           ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "container_multiple_durable_objects"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTransferContainerParityMismatch           ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "transfer_container_parity_mismatch"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTransferContainerParityMismatchOnCommit   ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "transfer_container_parity_mismatch_on_commit"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTombstoneClassStillInCode                 ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "tombstone_class_still_in_code"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioStaleTombstone                            ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "stale_tombstone"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTransferReceiveAlreadyApplied             ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "transfer_receive_already_applied"
	ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTransferReceiveCleanupComplete            ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario = "transfer_receive_cleanup_complete"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenario) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioCodeClassNotInExports, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioProvisionedClassMissingFromConfig, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioConfigExportNotInCode, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioConfigReferencesNonexistentClass, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioOrphanedProvisionedNamespace, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioStorageTypeMismatch, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioFreeTierRequiresSqlite, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioInvalidExport, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTombstoneDeleteClassStillInCode, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTombstoneDeleteBlockedByExternalBindings, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTombstoneRenamedToOccupied, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTransferredPendingNotFound, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTransferredTargetMissing, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTransferredTargetMismatch, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioPhaseOneTransferSourceMissing, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioPhaseOneTransferSourceNamespaceMissing, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioPhaseOneTransferTargetClassProvisioned, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioPhaseOneTransferAfterCommitMismatch, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioPhaseOneTransferDuplicate, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioPhaseOneTransferTargetInDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioPhaseOneTransferSourceInDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTransferredSourceInDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTransferredTargetInDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioContainerUndeclaredReference, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioContainerClassNotDurableObject, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioContainerWiringInconsistent, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioContainerMultipleDurableObjects, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTransferContainerParityMismatch, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTransferContainerParityMismatchOnCommit, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTombstoneClassStillInCode, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioStaleTombstone, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTransferReceiveAlreadyApplied, ScriptScriptAndVersionSettingEditParamsSettingsExportsReconciliationWarningsScenarioTransferReceiveCleanupComplete:
		return true
	}
	return false
}

// Limits to apply for this Worker.
type ScriptScriptAndVersionSettingEditParamsSettingsLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
	CPUMs param.Field[int64] `json:"cpu_ms"`
	// The number of subrequests this Worker can make per request.
	Subrequests param.Field[int64] `json:"subrequests"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsLimits) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Migrations to apply for Durable Objects associated with this Worker.
type ScriptScriptAndVersionSettingEditParamsSettingsMigrations struct {
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

func (r ScriptScriptAndVersionSettingEditParamsSettingsMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsMigrations) implementsScriptScriptAndVersionSettingEditParamsSettingsMigrationsUnion() {
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by [workers.SingleStepMigrationParam],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsMigrationsWorkersMultipleStepMigrations],
// [ScriptScriptAndVersionSettingEditParamsSettingsMigrations].
type ScriptScriptAndVersionSettingEditParamsSettingsMigrationsUnion interface {
	implementsScriptScriptAndVersionSettingEditParamsSettingsMigrationsUnion()
}

type ScriptScriptAndVersionSettingEditParamsSettingsMigrationsWorkersMultipleStepMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]MigrationStepParam] `json:"steps"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsMigrationsWorkersMultipleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsMigrationsWorkersMultipleStepMigrations) implementsScriptScriptAndVersionSettingEditParamsSettingsMigrationsUnion() {
}

// Observability settings for the Worker.
type ScriptScriptAndVersionSettingEditParamsSettingsObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
	// Log settings for the Worker.
	Logs param.Field[ScriptScriptAndVersionSettingEditParamsSettingsObservabilityLogs] `json:"logs"`
	// Whether query strings are removed from request URLs in logs and traces.
	RedactQueryString param.Field[bool] `json:"redact_query_string"`
	// Trace settings for the Worker.
	Traces param.Field[ScriptScriptAndVersionSettingEditParamsSettingsObservabilityTraces] `json:"traces"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsObservability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Log settings for the Worker.
type ScriptScriptAndVersionSettingEditParamsSettingsObservabilityLogs struct {
	// Whether logs are enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs param.Field[bool] `json:"invocation_logs" api:"required"`
	// A list of destinations where logs will be exported to.
	Destinations param.Field[[]string] `json:"destinations"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
	// Whether log persistence is enabled for the Worker.
	Persist param.Field[bool] `json:"persist"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsObservabilityLogs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Trace settings for the Worker.
type ScriptScriptAndVersionSettingEditParamsSettingsObservabilityTraces struct {
	// A list of destinations where traces will be exported to.
	Destinations param.Field[[]string] `json:"destinations"`
	// Whether traces are enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled"`
	// The sampling rate for traces. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
	// Whether trace persistence is enabled for the Worker.
	Persist param.Field[bool] `json:"persist"`
	// Controls how inbound trace context (traceparent/tracestate) headers on incoming
	// requests are handled. "authenticated" honors inbound trace context only when
	// accompanied by a valid trace auth token. "accept" unconditionally accepts
	// inbound trace context. Requires the trace propagation feature to be enabled.
	// Returns null when the trace propagation feature is not enabled for the account.
	PropagationPolicy param.Field[ScriptScriptAndVersionSettingEditParamsSettingsObservabilityTracesPropagationPolicy] `json:"propagation_policy"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsObservabilityTraces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls how inbound trace context (traceparent/tracestate) headers on incoming
// requests are handled. "authenticated" honors inbound trace context only when
// accompanied by a valid trace auth token. "accept" unconditionally accepts
// inbound trace context. Requires the trace propagation feature to be enabled.
// Returns null when the trace propagation feature is not enabled for the account.
type ScriptScriptAndVersionSettingEditParamsSettingsObservabilityTracesPropagationPolicy string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsObservabilityTracesPropagationPolicyAuthenticated ScriptScriptAndVersionSettingEditParamsSettingsObservabilityTracesPropagationPolicy = "authenticated"
	ScriptScriptAndVersionSettingEditParamsSettingsObservabilityTracesPropagationPolicyAccept        ScriptScriptAndVersionSettingEditParamsSettingsObservabilityTracesPropagationPolicy = "accept"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsObservabilityTracesPropagationPolicy) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsObservabilityTracesPropagationPolicyAuthenticated, ScriptScriptAndVersionSettingEditParamsSettingsObservabilityTracesPropagationPolicyAccept:
		return true
	}
	return false
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
type ScriptScriptAndVersionSettingEditParamsSettingsPlacement struct {
	// TCP host and port for targeted placement.
	Host param.Field[string] `json:"host"`
	// HTTP hostname for targeted placement.
	Hostname param.Field[string] `json:"hostname"`
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode param.Field[ScriptScriptAndVersionSettingEditParamsSettingsPlacementMode] `json:"mode"`
	// Cloud region for targeted placement in format 'provider:region'.
	Region param.Field[string]      `json:"region"`
	Target param.Field[interface{}] `json:"target"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsPlacement) implementsScriptScriptAndVersionSettingEditParamsSettingsPlacementUnion() {
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
//
// Satisfied by
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacementMode],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacementRegion],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacementHostname],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacementHost],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacementObject],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacementObject],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacementObject],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsPlacementObject],
// [ScriptScriptAndVersionSettingEditParamsSettingsPlacement].
type ScriptScriptAndVersionSettingEditParamsSettingsPlacementUnion interface {
	implementsScriptScriptAndVersionSettingEditParamsSettingsPlacementUnion()
}

type ScriptScriptAndVersionSettingEditParamsSettingsPlacementMode struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode param.Field[ScriptScriptAndVersionSettingEditParamsSettingsPlacementModeMode] `json:"mode" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsPlacementMode) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsPlacementMode) implementsScriptScriptAndVersionSettingEditParamsSettingsPlacementUnion() {
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptScriptAndVersionSettingEditParamsSettingsPlacementModeMode string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsPlacementModeModeSmart ScriptScriptAndVersionSettingEditParamsSettingsPlacementModeMode = "smart"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsPlacementModeMode) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsPlacementModeModeSmart:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsPlacementRegion struct {
	// Cloud region for targeted placement in format 'provider:region'.
	Region param.Field[string] `json:"region" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsPlacementRegion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsPlacementRegion) implementsScriptScriptAndVersionSettingEditParamsSettingsPlacementUnion() {
}

type ScriptScriptAndVersionSettingEditParamsSettingsPlacementHostname struct {
	// HTTP hostname for targeted placement.
	Hostname param.Field[string] `json:"hostname" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsPlacementHostname) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsPlacementHostname) implementsScriptScriptAndVersionSettingEditParamsSettingsPlacementUnion() {
}

type ScriptScriptAndVersionSettingEditParamsSettingsPlacementHost struct {
	// TCP host and port for targeted placement.
	Host param.Field[string] `json:"host" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsPlacementHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsPlacementHost) implementsScriptScriptAndVersionSettingEditParamsSettingsPlacementUnion() {
}

type ScriptScriptAndVersionSettingEditParamsSettingsPlacementObject struct {
	// Targeted placement mode.
	Mode param.Field[ScriptScriptAndVersionSettingEditParamsSettingsPlacementObjectMode] `json:"mode" api:"required"`
	// Cloud region for targeted placement in format 'provider:region'.
	Region param.Field[string] `json:"region" api:"required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsPlacementObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsPlacementObject) implementsScriptScriptAndVersionSettingEditParamsSettingsPlacementUnion() {
}

// Targeted placement mode.
type ScriptScriptAndVersionSettingEditParamsSettingsPlacementObjectMode string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsPlacementObjectModeTargeted ScriptScriptAndVersionSettingEditParamsSettingsPlacementObjectMode = "targeted"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsPlacementObjectMode) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsPlacementObjectModeTargeted:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type ScriptScriptAndVersionSettingEditParamsSettingsUsageModel string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsUsageModelStandard ScriptScriptAndVersionSettingEditParamsSettingsUsageModel = "standard"
	ScriptScriptAndVersionSettingEditParamsSettingsUsageModelBundled  ScriptScriptAndVersionSettingEditParamsSettingsUsageModel = "bundled"
	ScriptScriptAndVersionSettingEditParamsSettingsUsageModelUnbound  ScriptScriptAndVersionSettingEditParamsSettingsUsageModel = "unbound"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsUsageModel) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsUsageModelStandard, ScriptScriptAndVersionSettingEditParamsSettingsUsageModelBundled, ScriptScriptAndVersionSettingEditParamsSettingsUsageModelUnbound:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseEnvelope struct {
	Errors   []ScriptScriptAndVersionSettingEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ScriptScriptAndVersionSettingEditResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ScriptScriptAndVersionSettingEditResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success ScriptScriptAndVersionSettingEditResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    scriptScriptAndVersionSettingEditResponseEnvelopeJSON    `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [ScriptScriptAndVersionSettingEditResponseEnvelope]
type scriptScriptAndVersionSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptScriptAndVersionSettingEditResponseEnvelopeErrors struct {
	Code             int64                                                         `json:"code" api:"required"`
	Message          string                                                        `json:"message" api:"required"`
	DocumentationURL string                                                        `json:"documentation_url"`
	Source           ScriptScriptAndVersionSettingEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptScriptAndVersionSettingEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseEnvelopeErrors]
type scriptScriptAndVersionSettingEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptScriptAndVersionSettingEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                            `json:"pointer"`
	JSON    scriptScriptAndVersionSettingEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseEnvelopeErrorsSource]
type scriptScriptAndVersionSettingEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptScriptAndVersionSettingEditResponseEnvelopeMessages struct {
	Code             int64                                                           `json:"code" api:"required"`
	Message          string                                                          `json:"message" api:"required"`
	DocumentationURL string                                                          `json:"documentation_url"`
	Source           ScriptScriptAndVersionSettingEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptScriptAndVersionSettingEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseEnvelopeMessages]
type scriptScriptAndVersionSettingEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptScriptAndVersionSettingEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                              `json:"pointer"`
	JSON    scriptScriptAndVersionSettingEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseEnvelopeMessagesSource]
type scriptScriptAndVersionSettingEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptScriptAndVersionSettingEditResponseEnvelopeSuccess bool

const (
	ScriptScriptAndVersionSettingEditResponseEnvelopeSuccessTrue ScriptScriptAndVersionSettingEditResponseEnvelopeSuccess = true
)

func (r ScriptScriptAndVersionSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ScriptScriptAndVersionSettingGetResponseEnvelope struct {
	Errors   []ScriptScriptAndVersionSettingGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ScriptScriptAndVersionSettingGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ScriptScriptAndVersionSettingGetResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success ScriptScriptAndVersionSettingGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    scriptScriptAndVersionSettingGetResponseEnvelopeJSON    `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ScriptScriptAndVersionSettingGetResponseEnvelope]
type scriptScriptAndVersionSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptScriptAndVersionSettingGetResponseEnvelopeErrors struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           ScriptScriptAndVersionSettingGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             scriptScriptAndVersionSettingGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ScriptScriptAndVersionSettingGetResponseEnvelopeErrors]
type scriptScriptAndVersionSettingGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptScriptAndVersionSettingGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    scriptScriptAndVersionSettingGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseEnvelopeErrorsSource]
type scriptScriptAndVersionSettingGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScriptScriptAndVersionSettingGetResponseEnvelopeMessages struct {
	Code             int64                                                          `json:"code" api:"required"`
	Message          string                                                         `json:"message" api:"required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           ScriptScriptAndVersionSettingGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             scriptScriptAndVersionSettingGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseEnvelopeMessages]
type scriptScriptAndVersionSettingGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScriptScriptAndVersionSettingGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    scriptScriptAndVersionSettingGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseEnvelopeMessagesSource]
type scriptScriptAndVersionSettingGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScriptScriptAndVersionSettingGetResponseEnvelopeSuccess bool

const (
	ScriptScriptAndVersionSettingGetResponseEnvelopeSuccessTrue ScriptScriptAndVersionSettingGetResponseEnvelopeSuccess = true
)

func (r ScriptScriptAndVersionSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
