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

	"github.com/cloudflare/cloudflare-go/v4/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
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

// Patch metadata or config, such as bindings or usage model
func (r *ScriptScriptAndVersionSettingService) Edit(ctx context.Context, scriptName string, params ScriptScriptAndVersionSettingEditParams, opts ...option.RequestOption) (res *ScriptScriptAndVersionSettingEditResponse, err error) {
	var env ScriptScriptAndVersionSettingEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get metadata and config, such as bindings or usage model
func (r *ScriptScriptAndVersionSettingService) Get(ctx context.Context, scriptName string, query ScriptScriptAndVersionSettingGetParams, opts ...option.RequestOption) (res *ScriptScriptAndVersionSettingGetResponse, err error) {
	var env ScriptScriptAndVersionSettingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScriptScriptAndVersionSettingEditResponse struct {
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings []ScriptScriptAndVersionSettingEditResponseBinding `json:"bindings"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate string `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Limits to apply for this Worker.
	Limits ScriptScriptAndVersionSettingEditResponseLimits `json:"limits"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations ScriptScriptAndVersionSettingEditResponseMigrations `json:"migrations"`
	// Observability settings for the Worker.
	Observability ScriptScriptAndVersionSettingEditResponseObservability `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement ScriptScriptAndVersionSettingEditResponsePlacement `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ConsumerScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel ScriptScriptAndVersionSettingEditResponseUsageModel `json:"usage_model"`
	JSON       scriptScriptAndVersionSettingEditResponseJSON       `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseJSON contains the JSON metadata for the
// struct [ScriptScriptAndVersionSettingEditResponse]
type scriptScriptAndVersionSettingEditResponseJSON struct {
	Bindings           apijson.Field
	CompatibilityDate  apijson.Field
	CompatibilityFlags apijson.Field
	Limits             apijson.Field
	Logpush            apijson.Field
	Migrations         apijson.Field
	Observability      apijson.Field
	Placement          apijson.Field
	Tags               apijson.Field
	TailConsumers      apijson.Field
	UsageModel         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources
type ScriptScriptAndVersionSettingEditResponseBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsType `json:"type,required"`
	// Identifier of the D1 database to bind to.
	ID string `json:"id"`
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
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name"`
	// JSON data to use.
	Json string `json:"json"`
	// Namespace to bind to.
	Namespace string `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// This field can have the runtime type of
	// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutbound].
	Outbound interface{} `json:"outbound"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string `json:"script_name"`
	// Name of Worker to bind to.
	Service string `json:"service"`
	// The text value to use.
	Text  string                                               `json:"text"`
	JSON  scriptScriptAndVersionSettingEditResponseBindingJSON `json:"-"`
	union ScriptScriptAndVersionSettingEditResponseBindingsUnion
}

// scriptScriptAndVersionSettingEditResponseBindingJSON contains the JSON metadata
// for the struct [ScriptScriptAndVersionSettingEditResponseBinding]
type scriptScriptAndVersionSettingEditResponseBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	ID            apijson.Field
	BucketName    apijson.Field
	CertificateID apijson.Field
	ClassName     apijson.Field
	Dataset       apijson.Field
	Environment   apijson.Field
	IndexName     apijson.Field
	Json          apijson.Field
	Namespace     apijson.Field
	NamespaceID   apijson.Field
	Outbound      apijson.Field
	QueueName     apijson.Field
	ScriptName    apijson.Field
	Service       apijson.Field
	Text          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata].
func (r ScriptScriptAndVersionSettingEditResponseBinding) AsUnion() ScriptScriptAndVersionSettingEditResponseBindingsUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer],
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize]
// or
// [workers.ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata].
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
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "version_metadata",
		},
	)
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType `json:"type,required"`
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
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeService                ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeAI, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeD1, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeJson, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeService, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAITypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType `json:"type,required"`
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
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeService                ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeAI, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeD1, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeJson, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeService, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType `json:"type,required"`
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
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeService                ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeAI, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeD1, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeJson, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeService, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindAssetsTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType `json:"type,required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRendering) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeService                ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeAI, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeD1, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeJson, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeService, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type `json:"type,required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1JSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1JSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1JSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
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
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeService                ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeAI, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeD1, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeJson, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeService, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindD1TypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to.
	Namespace string `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType `json:"type,required"`
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
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeService                ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeAI, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeD1, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeJson, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeService, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

// Outbound worker.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params []string `json:"params"`
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

// Outbound worker.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
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
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType `json:"type,required"`
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
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeService                ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAI, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeD1, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeJson, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeService, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType `json:"type,required"`
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
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeService                ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeAI, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeD1, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeJson, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeService, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json string `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType `json:"type,required"`
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
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeService                ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeAI, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeD1, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeJson, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeService, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindJsonTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType `json:"type,required"`
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
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeService                ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeAI, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeD1, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeJson, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeService, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType `json:"type,required"`
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
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeService                ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeAI, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeD1, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeJson, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeService, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The text value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType `json:"type,required"`
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
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeService                ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeAI, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeD1, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeJson, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeService, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindPlainTextTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType `json:"type,required"`
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
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeService                ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeAI, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeD1, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeJson, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeService, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindQueueTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType `json:"type,required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2Bucket]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeService                ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeAI, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeD1, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeJson, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeService, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindR2BucketTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The secret value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType `json:"type,required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretText]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
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
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeService                ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeAI, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeD1, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeJson, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeService, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindSecretTextTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType `json:"type,required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindService]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
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
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeService                ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeAI, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeD1, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeJson, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeService, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindServiceTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType `json:"type,required"`
	JSON scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer]
type scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerJSON struct {
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumer) implementsScriptScriptAndVersionSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeService                ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeAI, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeD1, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeJson, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeService, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType `json:"type,required"`
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
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeService                ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeAI, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeD1, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeJson, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeService, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVectorizeTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType `json:"type,required"`
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
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeService                ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeAI, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeD1, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeJson, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeService, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditResponseBindingsType string

const (
	ScriptScriptAndVersionSettingEditResponseBindingsTypeAI                     ScriptScriptAndVersionSettingEditResponseBindingsType = "ai"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditResponseBindingsType = "analytics_engine"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeAssets                 ScriptScriptAndVersionSettingEditResponseBindingsType = "assets"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeBrowserRendering       ScriptScriptAndVersionSettingEditResponseBindingsType = "browser_rendering"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeD1                     ScriptScriptAndVersionSettingEditResponseBindingsType = "d1"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeDispatchNamespace      ScriptScriptAndVersionSettingEditResponseBindingsType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditResponseBindingsType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeHyperdrive             ScriptScriptAndVersionSettingEditResponseBindingsType = "hyperdrive"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeJson                   ScriptScriptAndVersionSettingEditResponseBindingsType = "json"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeKVNamespace            ScriptScriptAndVersionSettingEditResponseBindingsType = "kv_namespace"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeMTLSCertificate        ScriptScriptAndVersionSettingEditResponseBindingsType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditResponseBindingsTypePlainText              ScriptScriptAndVersionSettingEditResponseBindingsType = "plain_text"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeQueue                  ScriptScriptAndVersionSettingEditResponseBindingsType = "queue"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeR2Bucket               ScriptScriptAndVersionSettingEditResponseBindingsType = "r2_bucket"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeSecretText             ScriptScriptAndVersionSettingEditResponseBindingsType = "secret_text"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeService                ScriptScriptAndVersionSettingEditResponseBindingsType = "service"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeTailConsumer           ScriptScriptAndVersionSettingEditResponseBindingsType = "tail_consumer"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeVectorize              ScriptScriptAndVersionSettingEditResponseBindingsType = "vectorize"
	ScriptScriptAndVersionSettingEditResponseBindingsTypeVersionMetadata        ScriptScriptAndVersionSettingEditResponseBindingsType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditResponseBindingsType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseBindingsTypeAI, ScriptScriptAndVersionSettingEditResponseBindingsTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditResponseBindingsTypeAssets, ScriptScriptAndVersionSettingEditResponseBindingsTypeBrowserRendering, ScriptScriptAndVersionSettingEditResponseBindingsTypeD1, ScriptScriptAndVersionSettingEditResponseBindingsTypeDispatchNamespace, ScriptScriptAndVersionSettingEditResponseBindingsTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditResponseBindingsTypeHyperdrive, ScriptScriptAndVersionSettingEditResponseBindingsTypeJson, ScriptScriptAndVersionSettingEditResponseBindingsTypeKVNamespace, ScriptScriptAndVersionSettingEditResponseBindingsTypeMTLSCertificate, ScriptScriptAndVersionSettingEditResponseBindingsTypePlainText, ScriptScriptAndVersionSettingEditResponseBindingsTypeQueue, ScriptScriptAndVersionSettingEditResponseBindingsTypeR2Bucket, ScriptScriptAndVersionSettingEditResponseBindingsTypeSecretText, ScriptScriptAndVersionSettingEditResponseBindingsTypeService, ScriptScriptAndVersionSettingEditResponseBindingsTypeTailConsumer, ScriptScriptAndVersionSettingEditResponseBindingsTypeVectorize, ScriptScriptAndVersionSettingEditResponseBindingsTypeVersionMetadata:
		return true
	}
	return false
}

// Limits to apply for this Worker.
type ScriptScriptAndVersionSettingEditResponseLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
	CPUMs int64                                               `json:"cpu_ms"`
	JSON  scriptScriptAndVersionSettingEditResponseLimitsJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseLimitsJSON contains the JSON metadata
// for the struct [ScriptScriptAndVersionSettingEditResponseLimits]
type scriptScriptAndVersionSettingEditResponseLimitsJSON struct {
	CPUMs       apijson.Field
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
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
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
	NewTag             apijson.Field
	OldTag             apijson.Field
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
// Possible runtime types of the union are [workers.SingleStepMigration],
// [workers.ScriptScriptAndVersionSettingEditResponseMigrationsWorkersMultipleStepMigrations].
func (r ScriptScriptAndVersionSettingEditResponseMigrations) AsUnion() ScriptScriptAndVersionSettingEditResponseMigrationsUnion {
	return r.union
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by [workers.SingleStepMigration] or
// [workers.ScriptScriptAndVersionSettingEditResponseMigrationsWorkersMultipleStepMigrations].
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
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []MigrationStep                                                                      `json:"steps"`
	JSON  scriptScriptAndVersionSettingEditResponseMigrationsWorkersMultipleStepMigrationsJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseMigrationsWorkersMultipleStepMigrationsJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingEditResponseMigrationsWorkersMultipleStepMigrations]
type scriptScriptAndVersionSettingEditResponseMigrationsWorkersMultipleStepMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
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
	Enabled bool `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate float64                                                    `json:"head_sampling_rate,nullable"`
	JSON             scriptScriptAndVersionSettingEditResponseObservabilityJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseObservabilityJSON contains the JSON
// metadata for the struct [ScriptScriptAndVersionSettingEditResponseObservability]
type scriptScriptAndVersionSettingEditResponseObservabilityJSON struct {
	Enabled          apijson.Field
	HeadSamplingRate apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponseObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponseObservabilityJSON) RawJSON() string {
	return r.raw
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptScriptAndVersionSettingEditResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode ScriptScriptAndVersionSettingEditResponsePlacementMode `json:"mode"`
	JSON scriptScriptAndVersionSettingEditResponsePlacementJSON `json:"-"`
}

// scriptScriptAndVersionSettingEditResponsePlacementJSON contains the JSON
// metadata for the struct [ScriptScriptAndVersionSettingEditResponsePlacement]
type scriptScriptAndVersionSettingEditResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingEditResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingEditResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptScriptAndVersionSettingEditResponsePlacementMode string

const (
	ScriptScriptAndVersionSettingEditResponsePlacementModeSmart ScriptScriptAndVersionSettingEditResponsePlacementMode = "smart"
)

func (r ScriptScriptAndVersionSettingEditResponsePlacementMode) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponsePlacementModeSmart:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type ScriptScriptAndVersionSettingEditResponseUsageModel string

const (
	ScriptScriptAndVersionSettingEditResponseUsageModelStandard ScriptScriptAndVersionSettingEditResponseUsageModel = "standard"
)

func (r ScriptScriptAndVersionSettingEditResponseUsageModel) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditResponseUsageModelStandard:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponse struct {
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings []ScriptScriptAndVersionSettingGetResponseBinding `json:"bindings"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate string `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Limits to apply for this Worker.
	Limits ScriptScriptAndVersionSettingGetResponseLimits `json:"limits"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations ScriptScriptAndVersionSettingGetResponseMigrations `json:"migrations"`
	// Observability settings for the Worker.
	Observability ScriptScriptAndVersionSettingGetResponseObservability `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement ScriptScriptAndVersionSettingGetResponsePlacement `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ConsumerScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel ScriptScriptAndVersionSettingGetResponseUsageModel `json:"usage_model"`
	JSON       scriptScriptAndVersionSettingGetResponseJSON       `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseJSON contains the JSON metadata for the
// struct [ScriptScriptAndVersionSettingGetResponse]
type scriptScriptAndVersionSettingGetResponseJSON struct {
	Bindings           apijson.Field
	CompatibilityDate  apijson.Field
	CompatibilityFlags apijson.Field
	Limits             apijson.Field
	Logpush            apijson.Field
	Migrations         apijson.Field
	Observability      apijson.Field
	Placement          apijson.Field
	Tags               apijson.Field
	TailConsumers      apijson.Field
	UsageModel         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources
type ScriptScriptAndVersionSettingGetResponseBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsType `json:"type,required"`
	// Identifier of the D1 database to bind to.
	ID string `json:"id"`
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
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name"`
	// JSON data to use.
	Json string `json:"json"`
	// Namespace to bind to.
	Namespace string `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// This field can have the runtime type of
	// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutbound].
	Outbound interface{} `json:"outbound"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string `json:"script_name"`
	// Name of Worker to bind to.
	Service string `json:"service"`
	// The text value to use.
	Text  string                                              `json:"text"`
	JSON  scriptScriptAndVersionSettingGetResponseBindingJSON `json:"-"`
	union ScriptScriptAndVersionSettingGetResponseBindingsUnion
}

// scriptScriptAndVersionSettingGetResponseBindingJSON contains the JSON metadata
// for the struct [ScriptScriptAndVersionSettingGetResponseBinding]
type scriptScriptAndVersionSettingGetResponseBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	ID            apijson.Field
	BucketName    apijson.Field
	CertificateID apijson.Field
	ClassName     apijson.Field
	Dataset       apijson.Field
	Environment   apijson.Field
	IndexName     apijson.Field
	Json          apijson.Field
	Namespace     apijson.Field
	NamespaceID   apijson.Field
	Outbound      apijson.Field
	QueueName     apijson.Field
	ScriptName    apijson.Field
	Service       apijson.Field
	Text          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata].
func (r ScriptScriptAndVersionSettingGetResponseBinding) AsUnion() ScriptScriptAndVersionSettingGetResponseBindingsUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer],
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize]
// or
// [workers.ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata].
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
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "version_metadata",
		},
	)
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType `json:"type,required"`
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
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeService                ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeAI, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeD1, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeJson, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeService, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAITypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType `json:"type,required"`
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
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeService                ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeAI, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeD1, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeJson, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeService, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType `json:"type,required"`
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
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeService                ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeAI, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeD1, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeJson, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeService, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindAssetsTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType `json:"type,required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRendering) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeService                ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeAI, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeD1, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeJson, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeService, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type `json:"type,required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1JSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1JSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1JSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
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
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeService                ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeAI, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeD1, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeJson, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeService, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindD1TypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to.
	Namespace string `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType `json:"type,required"`
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
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeService                ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeAI, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeD1, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeJson, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeService, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

// Outbound worker.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params []string `json:"params"`
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

// Outbound worker.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
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
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType `json:"type,required"`
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
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeService                ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAI, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeD1, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeJson, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeService, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType `json:"type,required"`
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
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeService                ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeAI, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeD1, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeJson, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeService, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json string `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType `json:"type,required"`
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
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeService                ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeAI, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeD1, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeJson, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeService, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindJsonTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType `json:"type,required"`
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
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeService                ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeAI, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeD1, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeJson, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeService, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType `json:"type,required"`
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
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeService                ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeAI, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeD1, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeJson, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeService, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The text value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType `json:"type,required"`
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
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeService                ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeAI, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeD1, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeJson, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeService, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindPlainTextTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType `json:"type,required"`
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
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeService                ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeAI, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeD1, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeJson, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeService, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindQueueTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType `json:"type,required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2Bucket]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeService                ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeAI, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeD1, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeJson, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeService, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindR2BucketTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The secret value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType `json:"type,required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretText]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
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
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeService                ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeAI, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeD1, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeJson, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeService, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindSecretTextTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType `json:"type,required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindService]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
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
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeService                ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeAI, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeD1, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeJson, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeService, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindServiceTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType `json:"type,required"`
	JSON scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer]
type scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerJSON struct {
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerJSON) RawJSON() string {
	return r.raw
}

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumer) implementsScriptScriptAndVersionSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeService                ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeAI, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeD1, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeJson, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeService, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType `json:"type,required"`
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
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeService                ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeAI, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeD1, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeJson, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeService, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVectorizeTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType `json:"type,required"`
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
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeService                ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeAI, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeD1, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeJson, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeService, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingGetResponseBindingsType string

const (
	ScriptScriptAndVersionSettingGetResponseBindingsTypeAI                     ScriptScriptAndVersionSettingGetResponseBindingsType = "ai"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeAnalyticsEngine        ScriptScriptAndVersionSettingGetResponseBindingsType = "analytics_engine"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeAssets                 ScriptScriptAndVersionSettingGetResponseBindingsType = "assets"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeBrowserRendering       ScriptScriptAndVersionSettingGetResponseBindingsType = "browser_rendering"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeD1                     ScriptScriptAndVersionSettingGetResponseBindingsType = "d1"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeDispatchNamespace      ScriptScriptAndVersionSettingGetResponseBindingsType = "dispatch_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeDurableObjectNamespace ScriptScriptAndVersionSettingGetResponseBindingsType = "durable_object_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeHyperdrive             ScriptScriptAndVersionSettingGetResponseBindingsType = "hyperdrive"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeJson                   ScriptScriptAndVersionSettingGetResponseBindingsType = "json"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeKVNamespace            ScriptScriptAndVersionSettingGetResponseBindingsType = "kv_namespace"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeMTLSCertificate        ScriptScriptAndVersionSettingGetResponseBindingsType = "mtls_certificate"
	ScriptScriptAndVersionSettingGetResponseBindingsTypePlainText              ScriptScriptAndVersionSettingGetResponseBindingsType = "plain_text"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeQueue                  ScriptScriptAndVersionSettingGetResponseBindingsType = "queue"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeR2Bucket               ScriptScriptAndVersionSettingGetResponseBindingsType = "r2_bucket"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeSecretText             ScriptScriptAndVersionSettingGetResponseBindingsType = "secret_text"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeService                ScriptScriptAndVersionSettingGetResponseBindingsType = "service"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeTailConsumer           ScriptScriptAndVersionSettingGetResponseBindingsType = "tail_consumer"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeVectorize              ScriptScriptAndVersionSettingGetResponseBindingsType = "vectorize"
	ScriptScriptAndVersionSettingGetResponseBindingsTypeVersionMetadata        ScriptScriptAndVersionSettingGetResponseBindingsType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingGetResponseBindingsType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseBindingsTypeAI, ScriptScriptAndVersionSettingGetResponseBindingsTypeAnalyticsEngine, ScriptScriptAndVersionSettingGetResponseBindingsTypeAssets, ScriptScriptAndVersionSettingGetResponseBindingsTypeBrowserRendering, ScriptScriptAndVersionSettingGetResponseBindingsTypeD1, ScriptScriptAndVersionSettingGetResponseBindingsTypeDispatchNamespace, ScriptScriptAndVersionSettingGetResponseBindingsTypeDurableObjectNamespace, ScriptScriptAndVersionSettingGetResponseBindingsTypeHyperdrive, ScriptScriptAndVersionSettingGetResponseBindingsTypeJson, ScriptScriptAndVersionSettingGetResponseBindingsTypeKVNamespace, ScriptScriptAndVersionSettingGetResponseBindingsTypeMTLSCertificate, ScriptScriptAndVersionSettingGetResponseBindingsTypePlainText, ScriptScriptAndVersionSettingGetResponseBindingsTypeQueue, ScriptScriptAndVersionSettingGetResponseBindingsTypeR2Bucket, ScriptScriptAndVersionSettingGetResponseBindingsTypeSecretText, ScriptScriptAndVersionSettingGetResponseBindingsTypeService, ScriptScriptAndVersionSettingGetResponseBindingsTypeTailConsumer, ScriptScriptAndVersionSettingGetResponseBindingsTypeVectorize, ScriptScriptAndVersionSettingGetResponseBindingsTypeVersionMetadata:
		return true
	}
	return false
}

// Limits to apply for this Worker.
type ScriptScriptAndVersionSettingGetResponseLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
	CPUMs int64                                              `json:"cpu_ms"`
	JSON  scriptScriptAndVersionSettingGetResponseLimitsJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseLimitsJSON contains the JSON metadata
// for the struct [ScriptScriptAndVersionSettingGetResponseLimits]
type scriptScriptAndVersionSettingGetResponseLimitsJSON struct {
	CPUMs       apijson.Field
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
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
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
	NewTag             apijson.Field
	OldTag             apijson.Field
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
// Possible runtime types of the union are [workers.SingleStepMigration],
// [workers.ScriptScriptAndVersionSettingGetResponseMigrationsWorkersMultipleStepMigrations].
func (r ScriptScriptAndVersionSettingGetResponseMigrations) AsUnion() ScriptScriptAndVersionSettingGetResponseMigrationsUnion {
	return r.union
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by [workers.SingleStepMigration] or
// [workers.ScriptScriptAndVersionSettingGetResponseMigrationsWorkersMultipleStepMigrations].
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
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []MigrationStep                                                                     `json:"steps"`
	JSON  scriptScriptAndVersionSettingGetResponseMigrationsWorkersMultipleStepMigrationsJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseMigrationsWorkersMultipleStepMigrationsJSON
// contains the JSON metadata for the struct
// [ScriptScriptAndVersionSettingGetResponseMigrationsWorkersMultipleStepMigrations]
type scriptScriptAndVersionSettingGetResponseMigrationsWorkersMultipleStepMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
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
	Enabled bool `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate float64                                                   `json:"head_sampling_rate,nullable"`
	JSON             scriptScriptAndVersionSettingGetResponseObservabilityJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseObservabilityJSON contains the JSON
// metadata for the struct [ScriptScriptAndVersionSettingGetResponseObservability]
type scriptScriptAndVersionSettingGetResponseObservabilityJSON struct {
	Enabled          apijson.Field
	HeadSamplingRate apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponseObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponseObservabilityJSON) RawJSON() string {
	return r.raw
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptScriptAndVersionSettingGetResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode ScriptScriptAndVersionSettingGetResponsePlacementMode `json:"mode"`
	JSON scriptScriptAndVersionSettingGetResponsePlacementJSON `json:"-"`
}

// scriptScriptAndVersionSettingGetResponsePlacementJSON contains the JSON metadata
// for the struct [ScriptScriptAndVersionSettingGetResponsePlacement]
type scriptScriptAndVersionSettingGetResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScriptAndVersionSettingGetResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScriptAndVersionSettingGetResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptScriptAndVersionSettingGetResponsePlacementMode string

const (
	ScriptScriptAndVersionSettingGetResponsePlacementModeSmart ScriptScriptAndVersionSettingGetResponsePlacementMode = "smart"
)

func (r ScriptScriptAndVersionSettingGetResponsePlacementMode) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponsePlacementModeSmart:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type ScriptScriptAndVersionSettingGetResponseUsageModel string

const (
	ScriptScriptAndVersionSettingGetResponseUsageModelStandard ScriptScriptAndVersionSettingGetResponseUsageModel = "standard"
)

func (r ScriptScriptAndVersionSettingGetResponseUsageModel) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingGetResponseUsageModelStandard:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParams struct {
	// Identifier.
	AccountID param.Field[string]                                          `path:"account_id,required"`
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
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings param.Field[[]ScriptScriptAndVersionSettingEditParamsSettingsBindingUnion] `json:"bindings"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
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
	Placement param.Field[ScriptScriptAndVersionSettingEditParamsSettingsPlacement] `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]ConsumerScriptParam] `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[ScriptScriptAndVersionSettingEditParamsSettingsUsageModel] `json:"usage_model"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
type ScriptScriptAndVersionSettingEditParamsSettingsBinding struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsType] `json:"type,required"`
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

func (r ScriptScriptAndVersionSettingEditParamsSettingsBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBinding) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAI],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngine],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssets],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRendering],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespace],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespace],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdrive],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJson],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespace],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificate],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainText],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueue],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2Bucket],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretText],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindService],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumer],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorize],
// [workers.ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadata],
// [ScriptScriptAndVersionSettingEditParamsSettingsBinding].
type ScriptScriptAndVersionSettingEditParamsSettingsBindingUnion interface {
	implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion()
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType] `json:"type,required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAI) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAITypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType] `json:"type,required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngine) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngine) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType] `json:"type,required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssets) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRendering struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType] `json:"type,required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRendering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRendering) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type] `json:"type,required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace to bind to.
	Namespace param.Field[string] `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType] `json:"type,required"`
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
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

// Outbound worker.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params param.Field[[]string] `json:"params"`
	// Outbound worker.
	Worker param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutboundWorker] `json:"worker"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker.
	Service param.Field[string] `json:"service"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespace struct {
	// The exported class name of the Durable Object.
	ClassName param.Field[string] `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType] `json:"type,required"`
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
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType] `json:"type,required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdrive) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdrive) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json param.Field[string] `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType] `json:"type,required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJson) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType] `json:"type,required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespace) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType] `json:"type,required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificate) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The text value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType] `json:"type,required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainText) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType] `json:"type,required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueue) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType] `json:"type,required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2Bucket) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2Bucket) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The secret value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType] `json:"type,required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretText) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType] `json:"type,required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindService) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType] `json:"type,required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumer) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType] `json:"type,required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorize) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeVersionMetadata:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType] `json:"type,required"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadata) implementsScriptScriptAndVersionSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptScriptAndVersionSettingEditParamsSettingsBindingsType string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeAI                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "ai"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeAnalyticsEngine        ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "analytics_engine"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeAssets                 ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "assets"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeBrowserRendering       ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "browser_rendering"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeD1                     ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "d1"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeDispatchNamespace      ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "dispatch_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeDurableObjectNamespace ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "durable_object_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeHyperdrive             ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "hyperdrive"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeJson                   ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "json"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeKVNamespace            ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "kv_namespace"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeMTLSCertificate        ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "mtls_certificate"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypePlainText              ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "plain_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeQueue                  ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "queue"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeR2Bucket               ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "r2_bucket"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeSecretText             ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "secret_text"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeService                ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "service"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeTailConsumer           ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "tail_consumer"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeVectorize              ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "vectorize"
	ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeVersionMetadata        ScriptScriptAndVersionSettingEditParamsSettingsBindingsType = "version_metadata"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsBindingsType) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeAI, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeAnalyticsEngine, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeAssets, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeBrowserRendering, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeD1, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeDispatchNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeDurableObjectNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeHyperdrive, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeJson, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeKVNamespace, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeMTLSCertificate, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypePlainText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeQueue, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeR2Bucket, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeSecretText, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeService, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeTailConsumer, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeVectorize, ScriptScriptAndVersionSettingEditParamsSettingsBindingsTypeVersionMetadata:
		return true
	}
	return false
}

// Limits to apply for this Worker.
type ScriptScriptAndVersionSettingEditParamsSettingsLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
	CPUMs param.Field[int64] `json:"cpu_ms"`
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
	Enabled param.Field[bool] `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsObservability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptScriptAndVersionSettingEditParamsSettingsPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode param.Field[ScriptScriptAndVersionSettingEditParamsSettingsPlacementMode] `json:"mode"`
}

func (r ScriptScriptAndVersionSettingEditParamsSettingsPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type ScriptScriptAndVersionSettingEditParamsSettingsPlacementMode string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsPlacementModeSmart ScriptScriptAndVersionSettingEditParamsSettingsPlacementMode = "smart"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsPlacementMode) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsPlacementModeSmart:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type ScriptScriptAndVersionSettingEditParamsSettingsUsageModel string

const (
	ScriptScriptAndVersionSettingEditParamsSettingsUsageModelStandard ScriptScriptAndVersionSettingEditParamsSettingsUsageModel = "standard"
)

func (r ScriptScriptAndVersionSettingEditParamsSettingsUsageModel) IsKnown() bool {
	switch r {
	case ScriptScriptAndVersionSettingEditParamsSettingsUsageModelStandard:
		return true
	}
	return false
}

type ScriptScriptAndVersionSettingEditResponseEnvelope struct {
	Errors   []ScriptScriptAndVersionSettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptScriptAndVersionSettingEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ScriptScriptAndVersionSettingEditResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptScriptAndVersionSettingEditResponse                `json:"result"`
	JSON    scriptScriptAndVersionSettingEditResponseEnvelopeJSON    `json:"-"`
}

// scriptScriptAndVersionSettingEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [ScriptScriptAndVersionSettingEditResponseEnvelope]
type scriptScriptAndVersionSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
	Code             int64                                                         `json:"code,required"`
	Message          string                                                        `json:"message,required"`
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
	Code             int64                                                           `json:"code,required"`
	Message          string                                                          `json:"message,required"`
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
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptScriptAndVersionSettingGetResponseEnvelope struct {
	Errors   []ScriptScriptAndVersionSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptScriptAndVersionSettingGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ScriptScriptAndVersionSettingGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptScriptAndVersionSettingGetResponse                `json:"result"`
	JSON    scriptScriptAndVersionSettingGetResponseEnvelopeJSON    `json:"-"`
}

// scriptScriptAndVersionSettingGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ScriptScriptAndVersionSettingGetResponseEnvelope]
type scriptScriptAndVersionSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
	Code             int64                                                        `json:"code,required"`
	Message          string                                                       `json:"message,required"`
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
	Code             int64                                                          `json:"code,required"`
	Message          string                                                         `json:"message,required"`
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
