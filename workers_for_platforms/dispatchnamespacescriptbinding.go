// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/tidwall/gjson"
)

// DispatchNamespaceScriptBindingService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDispatchNamespaceScriptBindingService] method instead.
type DispatchNamespaceScriptBindingService struct {
	Options []option.RequestOption
}

// NewDispatchNamespaceScriptBindingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptBindingService(opts ...option.RequestOption) (r *DispatchNamespaceScriptBindingService) {
	r = &DispatchNamespaceScriptBindingService{}
	r.Options = opts
	return
}

// Fetch script bindings from a script uploaded to a Workers for Platforms
// namespace.
func (r *DispatchNamespaceScriptBindingService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptBindingGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[DispatchNamespaceScriptBindingGetResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/bindings", query.AccountID, dispatchNamespace, scriptName)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Fetch script bindings from a script uploaded to a Workers for Platforms
// namespace.
func (r *DispatchNamespaceScriptBindingService) GetAutoPaging(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptBindingGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DispatchNamespaceScriptBindingGetResponse] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, dispatchNamespace, scriptName, query, opts...))
}

// A binding to allow the Worker to communicate with resources
type DispatchNamespaceScriptBindingGetResponse struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseType `json:"type,required"`
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
	// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutbound].
	Outbound interface{} `json:"outbound"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string `json:"script_name"`
	// Name of Worker to bind to.
	Service string `json:"service"`
	// The text value to use.
	Text  string                                        `json:"text"`
	JSON  dispatchNamespaceScriptBindingGetResponseJSON `json:"-"`
	union DispatchNamespaceScriptBindingGetResponseUnion
}

// dispatchNamespaceScriptBindingGetResponseJSON contains the JSON metadata for the
// struct [DispatchNamespaceScriptBindingGetResponse]
type dispatchNamespaceScriptBindingGetResponseJSON struct {
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

func (r dispatchNamespaceScriptBindingGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptBindingGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DispatchNamespaceScriptBindingGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DispatchNamespaceScriptBindingGetResponseUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata].
func (r DispatchNamespaceScriptBindingGetResponse) AsUnion() DispatchNamespaceScriptBindingGetResponseUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize]
// or
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata].
type DispatchNamespaceScriptBindingGetResponseUnion interface {
	implementsDispatchNamespaceScriptBindingGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptBindingGetResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "browser_rendering",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "tail_consumer",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "version_metadata",
		},
	)
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAI) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeAI                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "ai"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeAssets                 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "assets"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeD1                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "d1"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeHyperdrive             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeJson                   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "json"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeKVNamespace            DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypePlainText              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeQueue                  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "queue"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeR2Bucket               DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeSecretText             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeService                DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "service"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeTailConsumer           DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeVectorize              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeAI, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeAssets, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeD1, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeHyperdrive, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeJson, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeKVNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypePlainText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeQueue, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeR2Bucket, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeSecretText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeService, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeTailConsumer, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeVectorize, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAITypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineJSON struct {
	Dataset     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngine) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeAI                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "ai"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeAssets                 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "assets"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeD1                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "d1"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeHyperdrive             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeJson                   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "json"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypePlainText              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeQueue                  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "queue"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeSecretText             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeService                DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "service"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeTailConsumer           DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeVectorize              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeAI, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeAssets, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeD1, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeHyperdrive, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeJson, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypePlainText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeQueue, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeSecretText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeService, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeTailConsumer, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeVectorize, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAnalyticsEngineTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssets) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeAI                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "ai"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeAssets                 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "assets"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeD1                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "d1"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeHyperdrive             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeJson                   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "json"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypePlainText              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeQueue                  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "queue"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeSecretText             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeService                DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "service"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeTailConsumer           DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeVectorize              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeAI, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeAssets, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeD1, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeHyperdrive, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeJson, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypePlainText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeQueue, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeSecretText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeService, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeTailConsumer, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeVectorize, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindAssetsTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRendering) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeAI                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "ai"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeAssets                 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "assets"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeD1                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "d1"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeHyperdrive             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeJson                   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "json"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypePlainText              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeQueue                  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "queue"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeSecretText             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeService                DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "service"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeTailConsumer           DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeVectorize              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeAI, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeAssets, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeD1, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeHyperdrive, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeJson, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypePlainText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeQueue, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeSecretText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeService, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeTailConsumer, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeVectorize, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindBrowserRenderingTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1JSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1JSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1JSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1JSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeAI                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "ai"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeAssets                 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "assets"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeD1                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "d1"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeHyperdrive             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeJson                   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "json"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeKVNamespace            DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypePlainText              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "plain_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeQueue                  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "queue"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeR2Bucket               DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeSecretText             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "secret_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeService                DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "service"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeTailConsumer           DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeVectorize              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "vectorize"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeAI, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeAssets, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeD1, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeHyperdrive, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeJson, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeKVNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypePlainText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeQueue, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeR2Bucket, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeSecretText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeService, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeTailConsumer, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeVectorize, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindD1TypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to.
	Namespace string `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType `json:"type,required"`
	// Outbound worker.
	Outbound DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutbound `json:"outbound"`
	JSON     dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceJSON     `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespace) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeAI                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "ai"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeAssets                 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "assets"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeD1                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "d1"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeHyperdrive             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeJson                   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "json"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypePlainText              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeQueue                  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "queue"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeSecretText             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeService                DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "service"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeTailConsumer           DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeVectorize              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeAI, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeAssets, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeD1, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeHyperdrive, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeJson, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypePlainText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeQueue, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeSecretText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeService, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeTailConsumer, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeVectorize, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

// Outbound worker.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params []string `json:"params"`
	// Outbound worker.
	Worker DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorker `json:"worker"`
	JSON   dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundJSON   `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutbound]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Environment of the outbound worker.
	Environment string `json:"environment"`
	// Name of the outbound worker.
	Service string                                                                                         `json:"service"`
	JSON    dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorkerJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorkerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorker]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDispatchNamespaceOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace struct {
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType `json:"type,required"`
	// The environment of the script_name to bind to.
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this
	// Worker.
	ScriptName string                                                                                `json:"script_name"`
	JSON       dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespace) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeAI                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "ai"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeAssets                 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "assets"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeD1                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "d1"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeHyperdrive             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeJson                   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "json"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypePlainText              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeQueue                  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "queue"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeSecretText             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeService                DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "service"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeTailConsumer           DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeVectorize              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeAI, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeAssets, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeD1, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeHyperdrive, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeJson, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypePlainText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeQueue, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeSecretText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeService, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeTailConsumer, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeVectorize, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindDurableObjectNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdrive) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeAI                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "ai"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeAssets                 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "assets"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeD1                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "d1"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeHyperdrive             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeJson                   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "json"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypePlainText              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeQueue                  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "queue"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeSecretText             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeService                DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "service"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeTailConsumer           DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeVectorize              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeAI, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeAssets, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeD1, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeHyperdrive, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeJson, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypePlainText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeQueue, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeSecretText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeService, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeTailConsumer, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeVectorize, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindHyperdriveTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson struct {
	// JSON data to use.
	Json string `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonJSON struct {
	Json        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJson) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeAI                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "ai"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeAssets                 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "assets"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeD1                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "d1"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeHyperdrive             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeJson                   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "json"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypePlainText              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeQueue                  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "queue"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeSecretText             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeService                DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "service"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeTailConsumer           DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeVectorize              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeAI, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeAssets, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeD1, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeHyperdrive, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeJson, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypePlainText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeQueue, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeSecretText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeService, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeTailConsumer, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeVectorize, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindJsonTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespace) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeAI                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "ai"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeAssets                 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "assets"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeD1                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "d1"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeHyperdrive             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeJson                   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "json"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypePlainText              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeQueue                  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "queue"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeSecretText             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeService                DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "service"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeTailConsumer           DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeVectorize              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeAI, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeAssets, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeD1, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeHyperdrive, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeJson, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypePlainText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeQueue, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeSecretText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeService, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeTailConsumer, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeVectorize, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindKVNamespaceTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateJSON struct {
	CertificateID apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificate) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeAI                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "ai"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeAssets                 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "assets"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeD1                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "d1"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeHyperdrive             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeJson                   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "json"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypePlainText              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeQueue                  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "queue"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeSecretText             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeService                DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "service"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeTailConsumer           DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeVectorize              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeAI, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeAssets, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeD1, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeHyperdrive, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeJson, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypePlainText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeQueue, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeSecretText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeService, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeTailConsumer, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeVectorize, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindMTLSCertificateTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The text value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainText) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeAI                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "ai"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeAssets                 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "assets"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeD1                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "d1"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeHyperdrive             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeJson                   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "json"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypePlainText              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeQueue                  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "queue"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeSecretText             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeService                DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "service"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeTailConsumer           DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeVectorize              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeAI, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeAssets, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeD1, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeHyperdrive, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeJson, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypePlainText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeQueue, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeSecretText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeService, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeTailConsumer, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeVectorize, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindPlainTextTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueue) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeAI                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "ai"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeAssets                 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "assets"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeD1                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "d1"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeHyperdrive             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeJson                   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "json"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypePlainText              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeQueue                  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "queue"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeSecretText             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeService                DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "service"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeTailConsumer           DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeVectorize              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeAI, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeAssets, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeD1, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeHyperdrive, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeJson, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypePlainText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeQueue, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeSecretText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeService, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeTailConsumer, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeVectorize, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindQueueTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2Bucket) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeAI                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "ai"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeAssets                 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "assets"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeD1                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "d1"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeHyperdrive             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeJson                   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "json"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypePlainText              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeQueue                  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "queue"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeSecretText             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeService                DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "service"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeTailConsumer           DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeVectorize              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeAI, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeAssets, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeD1, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeHyperdrive, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeJson, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypePlainText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeQueue, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeSecretText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeService, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeTailConsumer, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeVectorize, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindR2BucketTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The secret value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretText) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeAI                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "ai"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeAssets                 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "assets"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeD1                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "d1"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeHyperdrive             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeJson                   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "json"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypePlainText              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeQueue                  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "queue"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeSecretText             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeService                DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "service"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeTailConsumer           DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeVectorize              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeAI, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeAssets, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeD1, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeHyperdrive, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeJson, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypePlainText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeQueue, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeSecretText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeService, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeTailConsumer, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeVectorize, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindSecretTextTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindService) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeAI                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "ai"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeAssets                 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "assets"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeD1                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "d1"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeHyperdrive             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeJson                   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "json"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypePlainText              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeQueue                  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "queue"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeSecretText             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeService                DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "service"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeTailConsumer           DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeVectorize              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeAI, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeAssets, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeD1, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeHyperdrive, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeJson, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypePlainText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeQueue, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeSecretText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeService, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeTailConsumer, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeVectorize, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindServiceTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Tail Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerJSON struct {
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumer) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeAI                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "ai"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeAssets                 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "assets"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeD1                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "d1"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeHyperdrive             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeJson                   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "json"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypePlainText              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeQueue                  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "queue"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeSecretText             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeService                DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "service"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeTailConsumer           DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeVectorize              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeAI, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeAssets, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeD1, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeHyperdrive, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeJson, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypePlainText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeQueue, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeSecretText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeService, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeTailConsumer, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeVectorize, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindTailConsumerTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeJSON struct {
	IndexName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorize) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeAI                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "ai"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeAssets                 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "assets"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeD1                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "d1"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeHyperdrive             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeJson                   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "json"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypePlainText              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeQueue                  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "queue"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeSecretText             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeService                DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "service"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeTailConsumer           DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeVectorize              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeAI, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeAssets, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeD1, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeHyperdrive, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeJson, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypePlainText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeQueue, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeSecretText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeService, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeTailConsumer, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeVectorize, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVectorizeTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata]
type dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadata) implementsDispatchNamespaceScriptBindingGetResponse() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeAI                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "ai"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeAssets                 DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "assets"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeD1                     DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "d1"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeHyperdrive             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeJson                   DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "json"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypePlainText              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeQueue                  DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "queue"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeSecretText             DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeService                DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "service"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeTailConsumer           DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeVectorize              DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeAI, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeAssets, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeD1, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeHyperdrive, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeJson, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypePlainText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeQueue, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeSecretText, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeService, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeTailConsumer, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeVectorize, DispatchNamespaceScriptBindingGetResponseWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseType string

const (
	DispatchNamespaceScriptBindingGetResponseTypeAI                     DispatchNamespaceScriptBindingGetResponseType = "ai"
	DispatchNamespaceScriptBindingGetResponseTypeAnalyticsEngine        DispatchNamespaceScriptBindingGetResponseType = "analytics_engine"
	DispatchNamespaceScriptBindingGetResponseTypeAssets                 DispatchNamespaceScriptBindingGetResponseType = "assets"
	DispatchNamespaceScriptBindingGetResponseTypeBrowserRendering       DispatchNamespaceScriptBindingGetResponseType = "browser_rendering"
	DispatchNamespaceScriptBindingGetResponseTypeD1                     DispatchNamespaceScriptBindingGetResponseType = "d1"
	DispatchNamespaceScriptBindingGetResponseTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseTypeHyperdrive             DispatchNamespaceScriptBindingGetResponseType = "hyperdrive"
	DispatchNamespaceScriptBindingGetResponseTypeJson                   DispatchNamespaceScriptBindingGetResponseType = "json"
	DispatchNamespaceScriptBindingGetResponseTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseType = "mtls_certificate"
	DispatchNamespaceScriptBindingGetResponseTypePlainText              DispatchNamespaceScriptBindingGetResponseType = "plain_text"
	DispatchNamespaceScriptBindingGetResponseTypeQueue                  DispatchNamespaceScriptBindingGetResponseType = "queue"
	DispatchNamespaceScriptBindingGetResponseTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseTypeSecretText             DispatchNamespaceScriptBindingGetResponseType = "secret_text"
	DispatchNamespaceScriptBindingGetResponseTypeService                DispatchNamespaceScriptBindingGetResponseType = "service"
	DispatchNamespaceScriptBindingGetResponseTypeTailConsumer           DispatchNamespaceScriptBindingGetResponseType = "tail_consumer"
	DispatchNamespaceScriptBindingGetResponseTypeVectorize              DispatchNamespaceScriptBindingGetResponseType = "vectorize"
	DispatchNamespaceScriptBindingGetResponseTypeVersionMetadata        DispatchNamespaceScriptBindingGetResponseType = "version_metadata"
)

func (r DispatchNamespaceScriptBindingGetResponseType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseTypeAI, DispatchNamespaceScriptBindingGetResponseTypeAnalyticsEngine, DispatchNamespaceScriptBindingGetResponseTypeAssets, DispatchNamespaceScriptBindingGetResponseTypeBrowserRendering, DispatchNamespaceScriptBindingGetResponseTypeD1, DispatchNamespaceScriptBindingGetResponseTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseTypeHyperdrive, DispatchNamespaceScriptBindingGetResponseTypeJson, DispatchNamespaceScriptBindingGetResponseTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseTypeMTLSCertificate, DispatchNamespaceScriptBindingGetResponseTypePlainText, DispatchNamespaceScriptBindingGetResponseTypeQueue, DispatchNamespaceScriptBindingGetResponseTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseTypeSecretText, DispatchNamespaceScriptBindingGetResponseTypeService, DispatchNamespaceScriptBindingGetResponseTypeTailConsumer, DispatchNamespaceScriptBindingGetResponseTypeVectorize, DispatchNamespaceScriptBindingGetResponseTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
