// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

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
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/cloudflare/cloudflare-go/v4/workers"
	"github.com/tidwall/gjson"
)

// DispatchNamespaceScriptSettingService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDispatchNamespaceScriptSettingService] method instead.
type DispatchNamespaceScriptSettingService struct {
	Options []option.RequestOption
}

// NewDispatchNamespaceScriptSettingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptSettingService(opts ...option.RequestOption) (r *DispatchNamespaceScriptSettingService) {
	r = &DispatchNamespaceScriptSettingService{}
	r.Options = opts
	return
}

// Patch script metadata, such as bindings
func (r *DispatchNamespaceScriptSettingService) Edit(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptSettingEditParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptSettingEditResponse, err error) {
	var env DispatchNamespaceScriptSettingEditResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get script settings from a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptSettingService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptSettingGetParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptSettingGetResponse, err error) {
	var env DispatchNamespaceScriptSettingGetResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DispatchNamespaceScriptSettingEditResponse struct {
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings []DispatchNamespaceScriptSettingEditResponseBinding `json:"bindings"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate string `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Limits to apply for this Worker.
	Limits DispatchNamespaceScriptSettingEditResponseLimits `json:"limits"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations DispatchNamespaceScriptSettingEditResponseMigrations `json:"migrations"`
	// Observability settings for the Worker.
	Observability DispatchNamespaceScriptSettingEditResponseObservability `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement DispatchNamespaceScriptSettingEditResponsePlacement `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []workers.ConsumerScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel DispatchNamespaceScriptSettingEditResponseUsageModel `json:"usage_model"`
	JSON       dispatchNamespaceScriptSettingEditResponseJSON       `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptSettingEditResponse]
type dispatchNamespaceScriptSettingEditResponseJSON struct {
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

func (r *DispatchNamespaceScriptSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources
type DispatchNamespaceScriptSettingEditResponseBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsType `json:"type,required"`
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
	// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutbound].
	Outbound interface{} `json:"outbound"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string `json:"script_name"`
	// Name of Worker to bind to.
	Service string `json:"service"`
	// The text value to use.
	Text  string                                                `json:"text"`
	JSON  dispatchNamespaceScriptSettingEditResponseBindingJSON `json:"-"`
	union DispatchNamespaceScriptSettingEditResponseBindingsUnion
}

// dispatchNamespaceScriptSettingEditResponseBindingJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptSettingEditResponseBinding]
type dispatchNamespaceScriptSettingEditResponseBindingJSON struct {
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

func (r dispatchNamespaceScriptSettingEditResponseBindingJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptSettingEditResponseBinding) UnmarshalJSON(data []byte) (err error) {
	*r = DispatchNamespaceScriptSettingEditResponseBinding{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DispatchNamespaceScriptSettingEditResponseBindingsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata].
func (r DispatchNamespaceScriptSettingEditResponseBinding) AsUnion() DispatchNamespaceScriptSettingEditResponseBindingsUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize]
// or
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata].
type DispatchNamespaceScriptSettingEditResponseBindingsUnion interface {
	implementsDispatchNamespaceScriptSettingEditResponseBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptSettingEditResponseBindingsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "version_metadata",
		},
	)
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeService                DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeAI, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeD1, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeJson, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeService, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAITypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineJSON struct {
	Dataset     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeService                DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeAI, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeD1, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeJson, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeService, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngineTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeService                DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeAI, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeD1, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeJson, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeService, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssetsTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeService                DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeAI, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeD1, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeJson, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeService, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRenderingTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1JSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1JSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1JSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1JSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeService                DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeAI, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeD1, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeJson, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeService, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1TypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to.
	Namespace string `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType `json:"type,required"`
	// Outbound worker.
	Outbound DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutbound `json:"outbound"`
	JSON     dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceJSON     `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeService                DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeAI, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeD1, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeJson, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeService, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

// Outbound worker.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params []string `json:"params"`
	// Outbound worker.
	Worker DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker `json:"worker"`
	JSON   dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON   `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutbound]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment string `json:"environment"`
	// Name of the outbound worker.
	Service string                                                                                                  `json:"service"`
	JSON    dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace struct {
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType `json:"type,required"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string                                                                                         `json:"script_name"`
	JSON       dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeService                DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAI, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeD1, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeJson, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeService, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeService                DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeAI, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeD1, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeJson, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeService, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdriveTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json string `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonJSON struct {
	Json        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeService                DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeAI, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeD1, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeJson, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeService, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJsonTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeService                DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeAI, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeD1, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeJson, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeService, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateJSON struct {
	CertificateID apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeService                DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeAI, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeD1, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeJson, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeService, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificateTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The text value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeService                DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeAI, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeD1, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeJson, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeService, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainTextTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeService                DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeAI, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeD1, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeJson, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeService, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueueTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeService                DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeAI, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeD1, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeJson, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeService, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2BucketTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The secret value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeService                DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeAI, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeD1, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeJson, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeService, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretTextTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeService                DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeAI, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeD1, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeJson, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeService, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindServiceTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerJSON struct {
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeService                DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeAI, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeD1, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeJson, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeService, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumerTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeJSON struct {
	IndexName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeService                DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeAI, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeD1, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeJson, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeService, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorizeTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata) implementsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeService                DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeAI, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeD1, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeJson, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeService, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsTypeAI                     DispatchNamespaceScriptSettingEditResponseBindingsType = "ai"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditResponseBindingsType = "analytics_engine"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeAssets                 DispatchNamespaceScriptSettingEditResponseBindingsType = "assets"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeBrowserRendering       DispatchNamespaceScriptSettingEditResponseBindingsType = "browser_rendering"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeD1                     DispatchNamespaceScriptSettingEditResponseBindingsType = "d1"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeDispatchNamespace      DispatchNamespaceScriptSettingEditResponseBindingsType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeHyperdrive             DispatchNamespaceScriptSettingEditResponseBindingsType = "hyperdrive"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeJson                   DispatchNamespaceScriptSettingEditResponseBindingsType = "json"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeKVNamespace            DispatchNamespaceScriptSettingEditResponseBindingsType = "kv_namespace"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeMTLSCertificate        DispatchNamespaceScriptSettingEditResponseBindingsType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditResponseBindingsTypePlainText              DispatchNamespaceScriptSettingEditResponseBindingsType = "plain_text"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeQueue                  DispatchNamespaceScriptSettingEditResponseBindingsType = "queue"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeR2Bucket               DispatchNamespaceScriptSettingEditResponseBindingsType = "r2_bucket"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeSecretText             DispatchNamespaceScriptSettingEditResponseBindingsType = "secret_text"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeService                DispatchNamespaceScriptSettingEditResponseBindingsType = "service"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeTailConsumer           DispatchNamespaceScriptSettingEditResponseBindingsType = "tail_consumer"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeVectorize              DispatchNamespaceScriptSettingEditResponseBindingsType = "vectorize"
	DispatchNamespaceScriptSettingEditResponseBindingsTypeVersionMetadata        DispatchNamespaceScriptSettingEditResponseBindingsType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsTypeAI, DispatchNamespaceScriptSettingEditResponseBindingsTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditResponseBindingsTypeAssets, DispatchNamespaceScriptSettingEditResponseBindingsTypeBrowserRendering, DispatchNamespaceScriptSettingEditResponseBindingsTypeD1, DispatchNamespaceScriptSettingEditResponseBindingsTypeDispatchNamespace, DispatchNamespaceScriptSettingEditResponseBindingsTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditResponseBindingsTypeHyperdrive, DispatchNamespaceScriptSettingEditResponseBindingsTypeJson, DispatchNamespaceScriptSettingEditResponseBindingsTypeKVNamespace, DispatchNamespaceScriptSettingEditResponseBindingsTypeMTLSCertificate, DispatchNamespaceScriptSettingEditResponseBindingsTypePlainText, DispatchNamespaceScriptSettingEditResponseBindingsTypeQueue, DispatchNamespaceScriptSettingEditResponseBindingsTypeR2Bucket, DispatchNamespaceScriptSettingEditResponseBindingsTypeSecretText, DispatchNamespaceScriptSettingEditResponseBindingsTypeService, DispatchNamespaceScriptSettingEditResponseBindingsTypeTailConsumer, DispatchNamespaceScriptSettingEditResponseBindingsTypeVectorize, DispatchNamespaceScriptSettingEditResponseBindingsTypeVersionMetadata:
		return true
	}
	return false
}

// Limits to apply for this Worker.
type DispatchNamespaceScriptSettingEditResponseLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
	CPUMs int64                                                `json:"cpu_ms"`
	JSON  dispatchNamespaceScriptSettingEditResponseLimitsJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseLimitsJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptSettingEditResponseLimits]
type dispatchNamespaceScriptSettingEditResponseLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseLimitsJSON) RawJSON() string {
	return r.raw
}

// Migrations to apply for Durable Objects associated with this Worker.
type DispatchNamespaceScriptSettingEditResponseMigrations struct {
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
	// This field can have the runtime type of
	// [[]workers.SingleStepMigrationRenamedClass].
	RenamedClasses interface{} `json:"renamed_classes"`
	// This field can have the runtime type of [[]workers.MigrationStep].
	Steps interface{} `json:"steps"`
	// This field can have the runtime type of
	// [[]workers.SingleStepMigrationTransferredClass].
	TransferredClasses interface{}                                              `json:"transferred_classes"`
	JSON               dispatchNamespaceScriptSettingEditResponseMigrationsJSON `json:"-"`
	union              DispatchNamespaceScriptSettingEditResponseMigrationsUnion
}

// dispatchNamespaceScriptSettingEditResponseMigrationsJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingEditResponseMigrations]
type dispatchNamespaceScriptSettingEditResponseMigrationsJSON struct {
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

func (r dispatchNamespaceScriptSettingEditResponseMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptSettingEditResponseMigrations) UnmarshalJSON(data []byte) (err error) {
	*r = DispatchNamespaceScriptSettingEditResponseMigrations{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DispatchNamespaceScriptSettingEditResponseMigrationsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [workers.SingleStepMigration],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrations].
func (r DispatchNamespaceScriptSettingEditResponseMigrations) AsUnion() DispatchNamespaceScriptSettingEditResponseMigrationsUnion {
	return r.union
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by [workers.SingleStepMigration] or
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrations].
type DispatchNamespaceScriptSettingEditResponseMigrationsUnion interface {
	ImplementsDispatchNamespaceScriptSettingEditResponseMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptSettingEditResponseMigrationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(workers.SingleStepMigration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrations{}),
		},
	)
}

type DispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []workers.MigrationStep                                                               `json:"steps"`
	JSON  dispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrationsJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrationsJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrations]
type dispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseMigrationsWorkersMultipleStepMigrations) ImplementsDispatchNamespaceScriptSettingEditResponseMigrations() {
}

// Observability settings for the Worker.
type DispatchNamespaceScriptSettingEditResponseObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled bool `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate float64                                                     `json:"head_sampling_rate,nullable"`
	JSON             dispatchNamespaceScriptSettingEditResponseObservabilityJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseObservabilityJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseObservability]
type dispatchNamespaceScriptSettingEditResponseObservabilityJSON struct {
	Enabled          apijson.Field
	HeadSamplingRate apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseObservabilityJSON) RawJSON() string {
	return r.raw
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptSettingEditResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode DispatchNamespaceScriptSettingEditResponsePlacementMode `json:"mode"`
	JSON dispatchNamespaceScriptSettingEditResponsePlacementJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponsePlacementJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingEditResponsePlacement]
type dispatchNamespaceScriptSettingEditResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptSettingEditResponsePlacementMode string

const (
	DispatchNamespaceScriptSettingEditResponsePlacementModeSmart DispatchNamespaceScriptSettingEditResponsePlacementMode = "smart"
)

func (r DispatchNamespaceScriptSettingEditResponsePlacementMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponsePlacementModeSmart:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type DispatchNamespaceScriptSettingEditResponseUsageModel string

const (
	DispatchNamespaceScriptSettingEditResponseUsageModelStandard DispatchNamespaceScriptSettingEditResponseUsageModel = "standard"
)

func (r DispatchNamespaceScriptSettingEditResponseUsageModel) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseUsageModelStandard:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponse struct {
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings []DispatchNamespaceScriptSettingGetResponseBinding `json:"bindings"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate string `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Limits to apply for this Worker.
	Limits DispatchNamespaceScriptSettingGetResponseLimits `json:"limits"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations DispatchNamespaceScriptSettingGetResponseMigrations `json:"migrations"`
	// Observability settings for the Worker.
	Observability DispatchNamespaceScriptSettingGetResponseObservability `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement DispatchNamespaceScriptSettingGetResponsePlacement `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []workers.ConsumerScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel DispatchNamespaceScriptSettingGetResponseUsageModel `json:"usage_model"`
	JSON       dispatchNamespaceScriptSettingGetResponseJSON       `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseJSON contains the JSON metadata for the
// struct [DispatchNamespaceScriptSettingGetResponse]
type dispatchNamespaceScriptSettingGetResponseJSON struct {
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

func (r *DispatchNamespaceScriptSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources
type DispatchNamespaceScriptSettingGetResponseBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsType `json:"type,required"`
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
	// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutbound].
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
	JSON  dispatchNamespaceScriptSettingGetResponseBindingJSON `json:"-"`
	union DispatchNamespaceScriptSettingGetResponseBindingsUnion
}

// dispatchNamespaceScriptSettingGetResponseBindingJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptSettingGetResponseBinding]
type dispatchNamespaceScriptSettingGetResponseBindingJSON struct {
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

func (r dispatchNamespaceScriptSettingGetResponseBindingJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptSettingGetResponseBinding) UnmarshalJSON(data []byte) (err error) {
	*r = DispatchNamespaceScriptSettingGetResponseBinding{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DispatchNamespaceScriptSettingGetResponseBindingsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata].
func (r DispatchNamespaceScriptSettingGetResponseBinding) AsUnion() DispatchNamespaceScriptSettingGetResponseBindingsUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize]
// or
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata].
type DispatchNamespaceScriptSettingGetResponseBindingsUnion interface {
	implementsDispatchNamespaceScriptSettingGetResponseBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptSettingGetResponseBindingsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "version_metadata",
		},
	)
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeService                DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeAI, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeD1, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeJson, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeService, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAITypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineJSON struct {
	Dataset     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeService                DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeAI, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeD1, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeJson, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeService, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngineTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeService                DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeAI, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeD1, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeJson, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeService, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssetsTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeService                DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeAI, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeD1, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeJson, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeService, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRenderingTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1JSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1JSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1JSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1JSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeService                DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeAI, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeD1, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeJson, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeService, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1TypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to.
	Namespace string `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType `json:"type,required"`
	// Outbound worker.
	Outbound DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutbound `json:"outbound"`
	JSON     dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceJSON     `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeService                DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeAI, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeD1, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeJson, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeService, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

// Outbound worker.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params []string `json:"params"`
	// Outbound worker.
	Worker DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker `json:"worker"`
	JSON   dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON   `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutbound]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment string `json:"environment"`
	// Name of the outbound worker.
	Service string                                                                                                 `json:"service"`
	JSON    dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace struct {
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType `json:"type,required"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string                                                                                        `json:"script_name"`
	JSON       dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeService                DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAI, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeD1, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeJson, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeService, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeService                DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeAI, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeD1, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeJson, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeService, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdriveTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json string `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonJSON struct {
	Json        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeService                DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeAI, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeD1, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeJson, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeService, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJsonTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeService                DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeAI, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeD1, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeJson, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeService, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateJSON struct {
	CertificateID apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeService                DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeAI, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeD1, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeJson, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeService, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificateTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The text value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeService                DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeAI, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeD1, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeJson, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeService, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainTextTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeService                DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeAI, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeD1, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeJson, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeService, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueueTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeService                DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeAI, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeD1, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeJson, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeService, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2BucketTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The secret value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeService                DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeAI, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeD1, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeJson, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeService, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretTextTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeService                DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeAI, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeD1, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeJson, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeService, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindServiceTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerJSON struct {
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeService                DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeAI, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeD1, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeJson, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeService, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumerTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeJSON struct {
	IndexName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeService                DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeAI, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeD1, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeJson, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeService, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorizeTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata) implementsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeService                DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeAI, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeD1, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeJson, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeService, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsTypeAI                     DispatchNamespaceScriptSettingGetResponseBindingsType = "ai"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeAnalyticsEngine        DispatchNamespaceScriptSettingGetResponseBindingsType = "analytics_engine"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeAssets                 DispatchNamespaceScriptSettingGetResponseBindingsType = "assets"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeBrowserRendering       DispatchNamespaceScriptSettingGetResponseBindingsType = "browser_rendering"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeD1                     DispatchNamespaceScriptSettingGetResponseBindingsType = "d1"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeDispatchNamespace      DispatchNamespaceScriptSettingGetResponseBindingsType = "dispatch_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsType = "durable_object_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeHyperdrive             DispatchNamespaceScriptSettingGetResponseBindingsType = "hyperdrive"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeJson                   DispatchNamespaceScriptSettingGetResponseBindingsType = "json"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeKVNamespace            DispatchNamespaceScriptSettingGetResponseBindingsType = "kv_namespace"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeMTLSCertificate        DispatchNamespaceScriptSettingGetResponseBindingsType = "mtls_certificate"
	DispatchNamespaceScriptSettingGetResponseBindingsTypePlainText              DispatchNamespaceScriptSettingGetResponseBindingsType = "plain_text"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeQueue                  DispatchNamespaceScriptSettingGetResponseBindingsType = "queue"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeR2Bucket               DispatchNamespaceScriptSettingGetResponseBindingsType = "r2_bucket"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeSecretText             DispatchNamespaceScriptSettingGetResponseBindingsType = "secret_text"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeService                DispatchNamespaceScriptSettingGetResponseBindingsType = "service"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeTailConsumer           DispatchNamespaceScriptSettingGetResponseBindingsType = "tail_consumer"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeVectorize              DispatchNamespaceScriptSettingGetResponseBindingsType = "vectorize"
	DispatchNamespaceScriptSettingGetResponseBindingsTypeVersionMetadata        DispatchNamespaceScriptSettingGetResponseBindingsType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsTypeAI, DispatchNamespaceScriptSettingGetResponseBindingsTypeAnalyticsEngine, DispatchNamespaceScriptSettingGetResponseBindingsTypeAssets, DispatchNamespaceScriptSettingGetResponseBindingsTypeBrowserRendering, DispatchNamespaceScriptSettingGetResponseBindingsTypeD1, DispatchNamespaceScriptSettingGetResponseBindingsTypeDispatchNamespace, DispatchNamespaceScriptSettingGetResponseBindingsTypeDurableObjectNamespace, DispatchNamespaceScriptSettingGetResponseBindingsTypeHyperdrive, DispatchNamespaceScriptSettingGetResponseBindingsTypeJson, DispatchNamespaceScriptSettingGetResponseBindingsTypeKVNamespace, DispatchNamespaceScriptSettingGetResponseBindingsTypeMTLSCertificate, DispatchNamespaceScriptSettingGetResponseBindingsTypePlainText, DispatchNamespaceScriptSettingGetResponseBindingsTypeQueue, DispatchNamespaceScriptSettingGetResponseBindingsTypeR2Bucket, DispatchNamespaceScriptSettingGetResponseBindingsTypeSecretText, DispatchNamespaceScriptSettingGetResponseBindingsTypeService, DispatchNamespaceScriptSettingGetResponseBindingsTypeTailConsumer, DispatchNamespaceScriptSettingGetResponseBindingsTypeVectorize, DispatchNamespaceScriptSettingGetResponseBindingsTypeVersionMetadata:
		return true
	}
	return false
}

// Limits to apply for this Worker.
type DispatchNamespaceScriptSettingGetResponseLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
	CPUMs int64                                               `json:"cpu_ms"`
	JSON  dispatchNamespaceScriptSettingGetResponseLimitsJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseLimitsJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptSettingGetResponseLimits]
type dispatchNamespaceScriptSettingGetResponseLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseLimitsJSON) RawJSON() string {
	return r.raw
}

// Migrations to apply for Durable Objects associated with this Worker.
type DispatchNamespaceScriptSettingGetResponseMigrations struct {
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
	// This field can have the runtime type of
	// [[]workers.SingleStepMigrationRenamedClass].
	RenamedClasses interface{} `json:"renamed_classes"`
	// This field can have the runtime type of [[]workers.MigrationStep].
	Steps interface{} `json:"steps"`
	// This field can have the runtime type of
	// [[]workers.SingleStepMigrationTransferredClass].
	TransferredClasses interface{}                                             `json:"transferred_classes"`
	JSON               dispatchNamespaceScriptSettingGetResponseMigrationsJSON `json:"-"`
	union              DispatchNamespaceScriptSettingGetResponseMigrationsUnion
}

// dispatchNamespaceScriptSettingGetResponseMigrationsJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingGetResponseMigrations]
type dispatchNamespaceScriptSettingGetResponseMigrationsJSON struct {
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

func (r dispatchNamespaceScriptSettingGetResponseMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptSettingGetResponseMigrations) UnmarshalJSON(data []byte) (err error) {
	*r = DispatchNamespaceScriptSettingGetResponseMigrations{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DispatchNamespaceScriptSettingGetResponseMigrationsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [workers.SingleStepMigration],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrations].
func (r DispatchNamespaceScriptSettingGetResponseMigrations) AsUnion() DispatchNamespaceScriptSettingGetResponseMigrationsUnion {
	return r.union
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by [workers.SingleStepMigration] or
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrations].
type DispatchNamespaceScriptSettingGetResponseMigrationsUnion interface {
	ImplementsDispatchNamespaceScriptSettingGetResponseMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptSettingGetResponseMigrationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(workers.SingleStepMigration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrations{}),
		},
	)
}

type DispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []workers.MigrationStep                                                              `json:"steps"`
	JSON  dispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrationsJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrationsJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrations]
type dispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseMigrationsWorkersMultipleStepMigrations) ImplementsDispatchNamespaceScriptSettingGetResponseMigrations() {
}

// Observability settings for the Worker.
type DispatchNamespaceScriptSettingGetResponseObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled bool `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate float64                                                    `json:"head_sampling_rate,nullable"`
	JSON             dispatchNamespaceScriptSettingGetResponseObservabilityJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseObservabilityJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingGetResponseObservability]
type dispatchNamespaceScriptSettingGetResponseObservabilityJSON struct {
	Enabled          apijson.Field
	HeadSamplingRate apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseObservabilityJSON) RawJSON() string {
	return r.raw
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptSettingGetResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode DispatchNamespaceScriptSettingGetResponsePlacementMode `json:"mode"`
	JSON dispatchNamespaceScriptSettingGetResponsePlacementJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponsePlacementJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingGetResponsePlacement]
type dispatchNamespaceScriptSettingGetResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptSettingGetResponsePlacementMode string

const (
	DispatchNamespaceScriptSettingGetResponsePlacementModeSmart DispatchNamespaceScriptSettingGetResponsePlacementMode = "smart"
)

func (r DispatchNamespaceScriptSettingGetResponsePlacementMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponsePlacementModeSmart:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type DispatchNamespaceScriptSettingGetResponseUsageModel string

const (
	DispatchNamespaceScriptSettingGetResponseUsageModelStandard DispatchNamespaceScriptSettingGetResponseUsageModel = "standard"
)

func (r DispatchNamespaceScriptSettingGetResponseUsageModel) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseUsageModelStandard:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParams struct {
	// Identifier
	AccountID param.Field[string]                                           `path:"account_id,required"`
	Settings  param.Field[DispatchNamespaceScriptSettingEditParamsSettings] `json:"settings"`
}

func (r DispatchNamespaceScriptSettingEditParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type DispatchNamespaceScriptSettingEditParamsSettings struct {
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings param.Field[[]DispatchNamespaceScriptSettingEditParamsSettingsBindingUnion] `json:"bindings"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Limits to apply for this Worker.
	Limits param.Field[DispatchNamespaceScriptSettingEditParamsSettingsLimits] `json:"limits"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[DispatchNamespaceScriptSettingEditParamsSettingsMigrationsUnion] `json:"migrations"`
	// Observability settings for the Worker.
	Observability param.Field[DispatchNamespaceScriptSettingEditParamsSettingsObservability] `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Placement param.Field[DispatchNamespaceScriptSettingEditParamsSettingsPlacement] `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]workers.ConsumerScriptParam] `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[DispatchNamespaceScriptSettingEditParamsSettingsUsageModel] `json:"usage_model"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
type DispatchNamespaceScriptSettingEditParamsSettingsBinding struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsType] `json:"type,required"`
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

func (r DispatchNamespaceScriptSettingEditParamsSettingsBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBinding) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAI],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngine],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssets],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRendering],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespace],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespace],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdrive],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJson],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespace],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificate],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainText],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueue],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2Bucket],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretText],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindService],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumer],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorize],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadata],
// [DispatchNamespaceScriptSettingEditParamsSettingsBinding].
type DispatchNamespaceScriptSettingEditParamsSettingsBindingUnion interface {
	implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion()
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAI) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAITypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngine) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngine) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngineTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssets) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssetsTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRendering struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRendering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRendering) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRenderingTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1TypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace to bind to.
	Namespace param.Field[string] `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType] `json:"type,required"`
	// Outbound worker.
	Outbound param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutbound] `json:"outbound"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespace) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

// Outbound worker.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params param.Field[[]string] `json:"params"`
	// Outbound worker.
	Worker param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutboundWorker] `json:"worker"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker.
	Service param.Field[string] `json:"service"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespace struct {
	// The exported class name of the Durable Object.
	ClassName param.Field[string] `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType] `json:"type,required"`
	// The environment of the script_name to bind to.
	Environment param.Field[string] `json:"environment"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespace) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdrive) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdrive) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdriveTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json param.Field[string] `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJson) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJsonTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespace) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificate) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificateTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The text value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainText) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainTextTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueue) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueueTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2Bucket) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2Bucket) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2BucketTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The secret value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretText) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretTextTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindService) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindServiceTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumer) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumerTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorize) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorizeTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadata) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsSettingsBindingsType string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeAI                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "ai"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeAnalyticsEngine        DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "analytics_engine"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeAssets                 DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "assets"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeBrowserRendering       DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "browser_rendering"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeD1                     DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "d1"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeDispatchNamespace      DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "dispatch_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "durable_object_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeHyperdrive             DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "hyperdrive"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeJson                   DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "json"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeKVNamespace            DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "kv_namespace"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeMTLSCertificate        DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "mtls_certificate"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypePlainText              DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "plain_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeQueue                  DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "queue"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeR2Bucket               DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "r2_bucket"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeSecretText             DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "secret_text"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeService                DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "service"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeTailConsumer           DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "tail_consumer"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeVectorize              DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "vectorize"
	DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeVersionMetadata        DispatchNamespaceScriptSettingEditParamsSettingsBindingsType = "version_metadata"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeAI, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeAnalyticsEngine, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeAssets, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeBrowserRendering, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeD1, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeDispatchNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeDurableObjectNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeHyperdrive, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeJson, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeKVNamespace, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeMTLSCertificate, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypePlainText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeQueue, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeR2Bucket, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeSecretText, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeService, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeTailConsumer, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeVectorize, DispatchNamespaceScriptSettingEditParamsSettingsBindingsTypeVersionMetadata:
		return true
	}
	return false
}

// Limits to apply for this Worker.
type DispatchNamespaceScriptSettingEditParamsSettingsLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
	CPUMs param.Field[int64] `json:"cpu_ms"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsLimits) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Migrations to apply for Durable Objects associated with this Worker.
type DispatchNamespaceScriptSettingEditParamsSettingsMigrations struct {
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

func (r DispatchNamespaceScriptSettingEditParamsSettingsMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsMigrations) ImplementsDispatchNamespaceScriptSettingEditParamsSettingsMigrationsUnion() {
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by [workers.SingleStepMigrationParam],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsSettingsMigrationsWorkersMultipleStepMigrations],
// [DispatchNamespaceScriptSettingEditParamsSettingsMigrations].
type DispatchNamespaceScriptSettingEditParamsSettingsMigrationsUnion interface {
	ImplementsDispatchNamespaceScriptSettingEditParamsSettingsMigrationsUnion()
}

type DispatchNamespaceScriptSettingEditParamsSettingsMigrationsWorkersMultipleStepMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]workers.MigrationStepParam] `json:"steps"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsMigrationsWorkersMultipleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsMigrationsWorkersMultipleStepMigrations) ImplementsDispatchNamespaceScriptSettingEditParamsSettingsMigrationsUnion() {
}

// Observability settings for the Worker.
type DispatchNamespaceScriptSettingEditParamsSettingsObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsObservability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptSettingEditParamsSettingsPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode param.Field[DispatchNamespaceScriptSettingEditParamsSettingsPlacementMode] `json:"mode"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptSettingEditParamsSettingsPlacementMode string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsPlacementModeSmart DispatchNamespaceScriptSettingEditParamsSettingsPlacementMode = "smart"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsPlacementMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsPlacementModeSmart:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type DispatchNamespaceScriptSettingEditParamsSettingsUsageModel string

const (
	DispatchNamespaceScriptSettingEditParamsSettingsUsageModelStandard DispatchNamespaceScriptSettingEditParamsSettingsUsageModel = "standard"
)

func (r DispatchNamespaceScriptSettingEditParamsSettingsUsageModel) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSettingsUsageModelStandard:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceScriptSettingEditResponse                `json:"result"`
	JSON    dispatchNamespaceScriptSettingEditResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingEditResponseEnvelope]
type dispatchNamespaceScriptSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptSettingEditResponseEnvelopeSuccessTrue DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceScriptSettingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceScriptSettingGetResponse                `json:"result"`
	JSON    dispatchNamespaceScriptSettingGetResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptSettingGetResponseEnvelope]
type dispatchNamespaceScriptSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptSettingGetResponseEnvelopeSuccessTrue DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
