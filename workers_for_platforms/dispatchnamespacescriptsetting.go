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
	Type string `json:"type,required"`
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
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata{}),
		},
	)
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                     `json:"type,required"`
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

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                                  `json:"type,required"`
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

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                         `json:"type,required"`
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

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindBrowserRendering struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                                   `json:"type,required"`
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

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                     `json:"type,required"`
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

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to.
	Namespace string `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type string `json:"type,required"`
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
	Type string `json:"type,required"`
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

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                             `json:"type,required"`
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

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json string `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                       `json:"type,required"`
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

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type string                                                                              `json:"type,required"`
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

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                                  `json:"type,required"`
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

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The text value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type string                                                                            `json:"type,required"`
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

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                        `json:"type,required"`
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

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                           `json:"type,required"`
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

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The secret value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type string                                                                             `json:"type,required"`
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

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type string                                                                          `json:"type,required"`
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

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type string                                                                               `json:"type,required"`
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

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                            `json:"type,required"`
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

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                                  `json:"type,required"`
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
	Type string `json:"type,required"`
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
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDurableObjectNamespace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata{}),
		},
	)
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                    `json:"type,required"`
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

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                                 `json:"type,required"`
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

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                        `json:"type,required"`
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

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindBrowserRendering struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                                  `json:"type,required"`
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

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                    `json:"type,required"`
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

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to.
	Namespace string `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type string `json:"type,required"`
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
	Type string `json:"type,required"`
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

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                            `json:"type,required"`
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

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json string `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                      `json:"type,required"`
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

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type string                                                                             `json:"type,required"`
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

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                                 `json:"type,required"`
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

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The text value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type string                                                                           `json:"type,required"`
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

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                       `json:"type,required"`
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

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                          `json:"type,required"`
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

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The secret value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type string                                                                            `json:"type,required"`
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

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type string                                                                         `json:"type,required"`
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

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type string                                                                              `json:"type,required"`
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

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                           `json:"type,required"`
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

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type string                                                                                 `json:"type,required"`
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
	Type param.Field[string] `json:"type,required"`
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
	Type param.Field[string] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAI) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[string] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngine) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAnalyticsEngine) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[string] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindAssets) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRendering struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[string] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRendering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindBrowserRendering) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[string] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindD1) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace to bind to.
	Namespace param.Field[string] `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type param.Field[string] `json:"type,required"`
	// Outbound worker.
	Outbound param.Field[DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespaceOutbound] `json:"outbound"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindDispatchNamespace) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
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
	Type param.Field[string] `json:"type,required"`
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

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[string] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdrive) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindHyperdrive) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json param.Field[string] `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[string] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindJson) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[string] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindKVNamespace) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[string] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindMTLSCertificate) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The text value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[string] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindPlainText) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[string] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindQueue) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[string] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2Bucket) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindR2Bucket) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The secret value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[string] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindSecretText) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[string] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindService) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[string] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindTailConsumer) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[string] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVectorize) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
}

type DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[string] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsBindingsWorkersBindingKindVersionMetadata) implementsDispatchNamespaceScriptSettingEditParamsSettingsBindingUnion() {
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
