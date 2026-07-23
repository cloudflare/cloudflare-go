// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
)

// BetaWorkerVersionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaWorkerVersionService] method instead.
type BetaWorkerVersionService struct {
	Options []option.RequestOption
}

// NewBetaWorkerVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBetaWorkerVersionService(opts ...option.RequestOption) (r *BetaWorkerVersionService) {
	r = &BetaWorkerVersionService{}
	r.Options = opts
	return
}

// Create a new version.
func (r *BetaWorkerVersionService) New(ctx context.Context, workerID string, params BetaWorkerVersionNewParams, opts ...option.RequestOption) (res *Version, err error) {
	var env BetaWorkerVersionNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if workerID == "" {
		err = errors.New("missing required worker_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/workers/%s/versions", params.AccountID, workerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List all versions for a Worker.
func (r *BetaWorkerVersionService) List(ctx context.Context, workerID string, params BetaWorkerVersionListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Version], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if workerID == "" {
		err = errors.New("missing required worker_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/workers/%s/versions", params.AccountID, workerID)
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

// List all versions for a Worker.
func (r *BetaWorkerVersionService) ListAutoPaging(ctx context.Context, workerID string, params BetaWorkerVersionListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Version] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, workerID, params, opts...))
}

// Delete a version.
func (r *BetaWorkerVersionService) Delete(ctx context.Context, workerID string, versionID string, body BetaWorkerVersionDeleteParams, opts ...option.RequestOption) (res *BetaWorkerVersionDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if workerID == "" {
		err = errors.New("missing required worker_id parameter")
		return nil, err
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/workers/%s/versions/%s", body.AccountID, workerID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Get details about a specific version.
func (r *BetaWorkerVersionService) Get(ctx context.Context, workerID string, versionID string, params BetaWorkerVersionGetParams, opts ...option.RequestOption) (res *Version, err error) {
	var env BetaWorkerVersionGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if workerID == "" {
		err = errors.New("missing required worker_id parameter")
		return nil, err
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/workers/%s/versions/%s", params.AccountID, workerID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type Version struct {
	// Version identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the version was created.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// The integer version number, starting from one.
	Number int64 `json:"number" api:"required"`
	// All routable URLs that always point to this version. Does not include alias
	// URLs, since aliases can be updated to point to a different version.
	URLs []string `json:"urls" api:"required" format:"uri"`
	// Metadata about the version.
	Annotations VersionAnnotations `json:"annotations"`
	// Configuration for assets within a Worker.
	//
	// [`_headers`](https://developers.cloudflare.com/workers/static-assets/headers/#custom-headers)
	// and
	// [`_redirects`](https://developers.cloudflare.com/workers/static-assets/redirects/)
	// files should be included as modules named `_headers` and `_redirects` with
	// content type `text/plain`.
	Assets VersionAssets `json:"assets"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings []VersionBinding `json:"bindings"`
	// Global CacheW configuration for the Worker. When caching is on, the platform
	// provisions a `cloudflare.app` zone for the Worker. A `type: worker` entry in the
	// `exports` map can override this value for a single entrypoint.
	CacheOptions VersionCacheOptions `json:"cache_options"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate string `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// List of containers attached to a Worker. Containers can only be attached to
	// Durable Object classes of this Worker script.
	Containers []VersionContainer `json:"containers"`
	// Declarative exports for the version, including Durable Object classes (with
	// their `storage` backend) and named Worker entrypoints. On reads, tombstoned
	// lifecycle entries are omitted, so only live exports (`created` and
	// `expecting-transfer`) are returned. `exports` and `migrations` are mutually
	// exclusive on upload.
	Exports map[string]VersionExport `json:"exports"`
	// Summary of the declarative exports reconciliation that ran on this upload.
	// Populated only when the uploaded metadata included an `exports` block. Durable
	// Object entries drive reconciliation; `type: worker` entries do not contribute to
	// this summary.
	ExportsReconciliation VersionExportsReconciliation `json:"exports_reconciliation"`
	// Resource limits enforced at runtime.
	Limits VersionLimits `json:"limits"`
	// The name of the main module in the `modules` array (e.g. the name of the module
	// that exports a `fetch` handler).
	MainModule string `json:"main_module"`
	// Durable Object migration tag. Set when the version is deployed. Omitted if the
	// version has not been deployed or the Worker does not use Durable Objects.
	MigrationTag string `json:"migration_tag"`
	// Migrations for Durable Objects associated with the version. Migrations are
	// applied when the version is deployed.
	Migrations VersionMigrations `json:"migrations"`
	// Code, sourcemaps, and other content used at runtime.
	//
	// This includes
	// [`_headers`](https://developers.cloudflare.com/workers/static-assets/headers/#custom-headers)
	// and
	// [`_redirects`](https://developers.cloudflare.com/workers/static-assets/redirects/)
	// files used to configure
	// [Static Assets](https://developers.cloudflare.com/workers/static-assets/).
	// `_headers` and `_redirects` files should be included as modules named `_headers`
	// and `_redirects` with content type `text/plain`.
	Modules []VersionModule `json:"modules"`
	// The list of npm packages that were installed and used when this Worker version
	// was built.
	PackageDependencies []VersionPackageDependency `json:"package_dependencies"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
	Placement VersionPlacement `json:"placement"`
	// The client used to create the version.
	Source string `json:"source"`
	// Time in milliseconds spent on
	// [Worker startup](https://developers.cloudflare.com/workers/platform/limits/#worker-startup-time).
	StartupTimeMs int64 `json:"startup_time_ms"`
	// Usage model for the version.
	//
	// Deprecated: deprecated
	UsageModel VersionUsageModel `json:"usage_model"`
	JSON       versionJSON       `json:"-"`
}

// versionJSON contains the JSON metadata for the struct [Version]
type versionJSON struct {
	ID                    apijson.Field
	CreatedOn             apijson.Field
	Number                apijson.Field
	URLs                  apijson.Field
	Annotations           apijson.Field
	Assets                apijson.Field
	Bindings              apijson.Field
	CacheOptions          apijson.Field
	CompatibilityDate     apijson.Field
	CompatibilityFlags    apijson.Field
	Containers            apijson.Field
	Exports               apijson.Field
	ExportsReconciliation apijson.Field
	Limits                apijson.Field
	MainModule            apijson.Field
	MigrationTag          apijson.Field
	Migrations            apijson.Field
	Modules               apijson.Field
	PackageDependencies   apijson.Field
	Placement             apijson.Field
	Source                apijson.Field
	StartupTimeMs         apijson.Field
	UsageModel            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Version) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionJSON) RawJSON() string {
	return r.raw
}

// Metadata about the version.
type VersionAnnotations struct {
	// Human-readable message about the version. Truncated to 1000 bytes if longer.
	WorkersMessage string `json:"workers/message"`
	// User-provided identifier for the version. Maximum 100 bytes.
	WorkersTag string `json:"workers/tag"`
	// Operation that triggered the creation of the version.
	WorkersTriggeredBy string                 `json:"workers/triggered_by"`
	JSON               versionAnnotationsJSON `json:"-"`
}

// versionAnnotationsJSON contains the JSON metadata for the struct
// [VersionAnnotations]
type versionAnnotationsJSON struct {
	WorkersMessage     apijson.Field
	WorkersTag         apijson.Field
	WorkersTriggeredBy apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *VersionAnnotations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionAnnotationsJSON) RawJSON() string {
	return r.raw
}

// Configuration for assets within a Worker.
//
// [`_headers`](https://developers.cloudflare.com/workers/static-assets/headers/#custom-headers)
// and
// [`_redirects`](https://developers.cloudflare.com/workers/static-assets/redirects/)
// files should be included as modules named `_headers` and `_redirects` with
// content type `text/plain`.
type VersionAssets struct {
	// Configuration for assets within a Worker.
	Config VersionAssetsConfig `json:"config"`
	// Token provided upon successful upload of all files from a registered manifest.
	JWT  string            `json:"jwt"`
	JSON versionAssetsJSON `json:"-"`
}

// versionAssetsJSON contains the JSON metadata for the struct [VersionAssets]
type versionAssetsJSON struct {
	Config      apijson.Field
	JWT         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionAssets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionAssetsJSON) RawJSON() string {
	return r.raw
}

// Configuration for assets within a Worker.
type VersionAssetsConfig struct {
	// Determines the redirects and rewrites of requests for HTML content.
	HTMLHandling VersionAssetsConfigHTMLHandling `json:"html_handling"`
	// Determines the response when a request does not match a static asset, and there
	// is no Worker script.
	NotFoundHandling VersionAssetsConfigNotFoundHandling `json:"not_found_handling"`
	// Contains a list path rules to control routing to either the Worker or assets.
	// Glob (\*) and negative (!) rules are supported. Rules must start with either '/'
	// or '!/'. At least one non-negative rule must be provided, and negative rules
	// have higher precedence than non-negative rules.
	RunWorkerFirst VersionAssetsConfigRunWorkerFirstUnion `json:"run_worker_first"`
	JSON           versionAssetsConfigJSON                `json:"-"`
}

// versionAssetsConfigJSON contains the JSON metadata for the struct
// [VersionAssetsConfig]
type versionAssetsConfigJSON struct {
	HTMLHandling     apijson.Field
	NotFoundHandling apijson.Field
	RunWorkerFirst   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VersionAssetsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionAssetsConfigJSON) RawJSON() string {
	return r.raw
}

// Determines the redirects and rewrites of requests for HTML content.
type VersionAssetsConfigHTMLHandling string

const (
	VersionAssetsConfigHTMLHandlingAutoTrailingSlash  VersionAssetsConfigHTMLHandling = "auto-trailing-slash"
	VersionAssetsConfigHTMLHandlingForceTrailingSlash VersionAssetsConfigHTMLHandling = "force-trailing-slash"
	VersionAssetsConfigHTMLHandlingDropTrailingSlash  VersionAssetsConfigHTMLHandling = "drop-trailing-slash"
	VersionAssetsConfigHTMLHandlingNone               VersionAssetsConfigHTMLHandling = "none"
)

func (r VersionAssetsConfigHTMLHandling) IsKnown() bool {
	switch r {
	case VersionAssetsConfigHTMLHandlingAutoTrailingSlash, VersionAssetsConfigHTMLHandlingForceTrailingSlash, VersionAssetsConfigHTMLHandlingDropTrailingSlash, VersionAssetsConfigHTMLHandlingNone:
		return true
	}
	return false
}

// Determines the response when a request does not match a static asset, and there
// is no Worker script.
type VersionAssetsConfigNotFoundHandling string

const (
	VersionAssetsConfigNotFoundHandlingNone                  VersionAssetsConfigNotFoundHandling = "none"
	VersionAssetsConfigNotFoundHandling404Page               VersionAssetsConfigNotFoundHandling = "404-page"
	VersionAssetsConfigNotFoundHandlingSinglePageApplication VersionAssetsConfigNotFoundHandling = "single-page-application"
)

func (r VersionAssetsConfigNotFoundHandling) IsKnown() bool {
	switch r {
	case VersionAssetsConfigNotFoundHandlingNone, VersionAssetsConfigNotFoundHandling404Page, VersionAssetsConfigNotFoundHandlingSinglePageApplication:
		return true
	}
	return false
}

// Contains a list path rules to control routing to either the Worker or assets.
// Glob (\*) and negative (!) rules are supported. Rules must start with either '/'
// or '!/'. At least one non-negative rule must be provided, and negative rules
// have higher precedence than non-negative rules.
//
// Union satisfied by [VersionAssetsConfigRunWorkerFirstArray] or
// [shared.UnionBool].
type VersionAssetsConfigRunWorkerFirstUnion interface {
	ImplementsVersionAssetsConfigRunWorkerFirstUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionAssetsConfigRunWorkerFirstUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionAssetsConfigRunWorkerFirstArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type VersionAssetsConfigRunWorkerFirstArray []string

func (r VersionAssetsConfigRunWorkerFirstArray) ImplementsVersionAssetsConfigRunWorkerFirstUnion() {}

// A binding to allow the Worker to communicate with resources.
type VersionBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsType `json:"type" api:"required"`
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
	Format VersionBindingsFormat `json:"format"`
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
	Jurisdiction VersionBindingsJurisdiction `json:"jurisdiction"`
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
	// [VersionBindingsWorkersBindingKindDispatchNamespaceOutbound].
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
	// [VersionBindingsWorkersBindingKindRatelimitSimple].
	Simple interface{} `json:"simple"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id"`
	// The text value to use.
	Text string `json:"text"`
	// UUID of the Cloudflare Tunnel to bind to. Mutually exclusive with network_id.
	TunnelID string `json:"tunnel_id"`
	// This field can have the runtime type of
	// [[]VersionBindingsWorkersBindingKindSecretKeyUsage].
	Usages interface{} `json:"usages"`
	// Identifier for the version to inherit the binding from, which can be the version
	// ID or the literal "latest" to inherit from the latest version. Defaults to
	// inheriting the binding from the latest version.
	VersionID string `json:"version_id"`
	// Name of the Workflow to bind to.
	WorkflowName string             `json:"workflow_name"`
	JSON         versionBindingJSON `json:"-"`
	union        VersionBindingsUnion
}

// versionBindingJSON contains the JSON metadata for the struct [VersionBinding]
type versionBindingJSON struct {
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

func (r versionBindingJSON) RawJSON() string {
	return r.raw
}

func (r *VersionBinding) UnmarshalJSON(data []byte) (err error) {
	*r = VersionBinding{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [VersionBindingsUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [VersionBindingsWorkersBindingKindAI],
// [VersionBindingsWorkersBindingKindAISearch],
// [VersionBindingsWorkersBindingKindAISearchNamespace],
// [VersionBindingsWorkersBindingKindAnalyticsEngine],
// [VersionBindingsWorkersBindingKindAssets],
// [VersionBindingsWorkersBindingKindBrowser],
// [VersionBindingsWorkersBindingKindD1],
// [VersionBindingsWorkersBindingKindDataBlob],
// [VersionBindingsWorkersBindingKindDispatchNamespace],
// [VersionBindingsWorkersBindingKindDurableObjectNamespace],
// [VersionBindingsWorkersBindingKindHyperdrive],
// [VersionBindingsWorkersBindingKindInherit],
// [VersionBindingsWorkersBindingKindImages],
// [VersionBindingsWorkersBindingKindJson],
// [VersionBindingsWorkersBindingKindKVNamespace],
// [VersionBindingsWorkersBindingKindMedia],
// [VersionBindingsWorkersBindingKindMTLSCertificate],
// [VersionBindingsWorkersBindingKindPlainText],
// [VersionBindingsWorkersBindingKindPipelines],
// [VersionBindingsWorkersBindingKindQueue],
// [VersionBindingsWorkersBindingKindRatelimit],
// [VersionBindingsWorkersBindingKindR2Bucket],
// [VersionBindingsWorkersBindingKindSecretText],
// [VersionBindingsWorkersBindingKindSendEmail],
// [VersionBindingsWorkersBindingKindService],
// [VersionBindingsWorkersBindingKindTextBlob],
// [VersionBindingsWorkersBindingKindVectorize],
// [VersionBindingsWorkersBindingKindVersionMetadata],
// [VersionBindingsWorkersBindingKindSecretsStoreSecret],
// [VersionBindingsWorkersBindingKindFlagship],
// [VersionBindingsWorkersBindingKindSecretKey],
// [VersionBindingsWorkersBindingKindWorkflow],
// [VersionBindingsWorkersBindingKindWasmModule],
// [VersionBindingsWorkersBindingKindVPCService],
// [VersionBindingsWorkersBindingKindVPCNetwork].
func (r VersionBinding) AsUnion() VersionBindingsUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources.
//
// Union satisfied by [VersionBindingsWorkersBindingKindAI],
// [VersionBindingsWorkersBindingKindAISearch],
// [VersionBindingsWorkersBindingKindAISearchNamespace],
// [VersionBindingsWorkersBindingKindAnalyticsEngine],
// [VersionBindingsWorkersBindingKindAssets],
// [VersionBindingsWorkersBindingKindBrowser],
// [VersionBindingsWorkersBindingKindD1],
// [VersionBindingsWorkersBindingKindDataBlob],
// [VersionBindingsWorkersBindingKindDispatchNamespace],
// [VersionBindingsWorkersBindingKindDurableObjectNamespace],
// [VersionBindingsWorkersBindingKindHyperdrive],
// [VersionBindingsWorkersBindingKindInherit],
// [VersionBindingsWorkersBindingKindImages],
// [VersionBindingsWorkersBindingKindJson],
// [VersionBindingsWorkersBindingKindKVNamespace],
// [VersionBindingsWorkersBindingKindMedia],
// [VersionBindingsWorkersBindingKindMTLSCertificate],
// [VersionBindingsWorkersBindingKindPlainText],
// [VersionBindingsWorkersBindingKindPipelines],
// [VersionBindingsWorkersBindingKindQueue],
// [VersionBindingsWorkersBindingKindRatelimit],
// [VersionBindingsWorkersBindingKindR2Bucket],
// [VersionBindingsWorkersBindingKindSecretText],
// [VersionBindingsWorkersBindingKindSendEmail],
// [VersionBindingsWorkersBindingKindService],
// [VersionBindingsWorkersBindingKindTextBlob],
// [VersionBindingsWorkersBindingKindVectorize],
// [VersionBindingsWorkersBindingKindVersionMetadata],
// [VersionBindingsWorkersBindingKindSecretsStoreSecret],
// [VersionBindingsWorkersBindingKindFlagship],
// [VersionBindingsWorkersBindingKindSecretKey],
// [VersionBindingsWorkersBindingKindWorkflow],
// [VersionBindingsWorkersBindingKindWasmModule],
// [VersionBindingsWorkersBindingKindVPCService] or
// [VersionBindingsWorkersBindingKindVPCNetwork].
type VersionBindingsUnion interface {
	implementsVersionBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionBindingsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindAI{}),
			DiscriminatorValue: "ai",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindAISearch{}),
			DiscriminatorValue: "ai_search",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindAISearchNamespace{}),
			DiscriminatorValue: "ai_search_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindAnalyticsEngine{}),
			DiscriminatorValue: "analytics_engine",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindAssets{}),
			DiscriminatorValue: "assets",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindBrowser{}),
			DiscriminatorValue: "browser",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindD1{}),
			DiscriminatorValue: "d1",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindDataBlob{}),
			DiscriminatorValue: "data_blob",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindDispatchNamespace{}),
			DiscriminatorValue: "dispatch_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindDurableObjectNamespace{}),
			DiscriminatorValue: "durable_object_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindHyperdrive{}),
			DiscriminatorValue: "hyperdrive",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindInherit{}),
			DiscriminatorValue: "inherit",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindImages{}),
			DiscriminatorValue: "images",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindKVNamespace{}),
			DiscriminatorValue: "kv_namespace",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindMedia{}),
			DiscriminatorValue: "media",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindMTLSCertificate{}),
			DiscriminatorValue: "mtls_certificate",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindPlainText{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindPipelines{}),
			DiscriminatorValue: "pipelines",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindQueue{}),
			DiscriminatorValue: "queue",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindRatelimit{}),
			DiscriminatorValue: "ratelimit",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindR2Bucket{}),
			DiscriminatorValue: "r2_bucket",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindSecretText{}),
			DiscriminatorValue: "secret_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindSendEmail{}),
			DiscriminatorValue: "send_email",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindService{}),
			DiscriminatorValue: "service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindTextBlob{}),
			DiscriminatorValue: "text_blob",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindVectorize{}),
			DiscriminatorValue: "vectorize",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindVersionMetadata{}),
			DiscriminatorValue: "version_metadata",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindSecretsStoreSecret{}),
			DiscriminatorValue: "secrets_store_secret",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindFlagship{}),
			DiscriminatorValue: "flagship",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindSecretKey{}),
			DiscriminatorValue: "secret_key",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindWorkflow{}),
			DiscriminatorValue: "workflow",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindWasmModule{}),
			DiscriminatorValue: "wasm_module",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindVPCService{}),
			DiscriminatorValue: "vpc_service",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionBindingsWorkersBindingKindVPCNetwork{}),
			DiscriminatorValue: "vpc_network",
		},
	)
}

type VersionBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindAIType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindAIJSON `json:"-"`
}

// versionBindingsWorkersBindingKindAIJSON contains the JSON metadata for the
// struct [VersionBindingsWorkersBindingKindAI]
type versionBindingsWorkersBindingKindAIJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindAI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindAIJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindAI) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindAIType string

const (
	VersionBindingsWorkersBindingKindAITypeAI VersionBindingsWorkersBindingKindAIType = "ai"
)

func (r VersionBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindAISearch struct {
	// The user-chosen instance name. Must exist at deploy time. The worker can search,
	// chat, update, and manage items/jobs on this instance.
	InstanceName string `json:"instance_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindAISearchType `json:"type" api:"required"`
	// The namespace the instance belongs to. Defaults to "default" if omitted.
	// Customers who don't use namespaces can simply omit this field.
	Namespace string                                        `json:"namespace"`
	JSON      versionBindingsWorkersBindingKindAISearchJSON `json:"-"`
}

// versionBindingsWorkersBindingKindAISearchJSON contains the JSON metadata for the
// struct [VersionBindingsWorkersBindingKindAISearch]
type versionBindingsWorkersBindingKindAISearchJSON struct {
	InstanceName apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	Namespace    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindAISearch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindAISearchJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindAISearch) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindAISearchType string

const (
	VersionBindingsWorkersBindingKindAISearchTypeAISearch VersionBindingsWorkersBindingKindAISearchType = "ai_search"
)

func (r VersionBindingsWorkersBindingKindAISearchType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindAISearchTypeAISearch:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindAISearchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The user-chosen namespace name. Must exist before deploy -- Wrangler handles
	// auto-creation on deploy failure (R2 bucket pattern). The "default" namespace is
	// auto-created by config-api for new accounts. Grants full access (CRUD + search +
	// chat) to all instances within the namespace.
	Namespace string `json:"namespace" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindAISearchNamespaceType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindAISearchNamespaceJSON `json:"-"`
}

// versionBindingsWorkersBindingKindAISearchNamespaceJSON contains the JSON
// metadata for the struct [VersionBindingsWorkersBindingKindAISearchNamespace]
type versionBindingsWorkersBindingKindAISearchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindAISearchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindAISearchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindAISearchNamespace) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindAISearchNamespaceType string

const (
	VersionBindingsWorkersBindingKindAISearchNamespaceTypeAISearchNamespace VersionBindingsWorkersBindingKindAISearchNamespaceType = "ai_search_namespace"
)

func (r VersionBindingsWorkersBindingKindAISearchNamespaceType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindAISearchNamespaceTypeAISearchNamespace:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindAnalyticsEngineType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindAnalyticsEngineJSON `json:"-"`
}

// versionBindingsWorkersBindingKindAnalyticsEngineJSON contains the JSON metadata
// for the struct [VersionBindingsWorkersBindingKindAnalyticsEngine]
type versionBindingsWorkersBindingKindAnalyticsEngineJSON struct {
	Dataset     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindAnalyticsEngine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindAnalyticsEngineJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindAnalyticsEngine) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindAnalyticsEngineType string

const (
	VersionBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine VersionBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r VersionBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindAssetsType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindAssetsJSON `json:"-"`
}

// versionBindingsWorkersBindingKindAssetsJSON contains the JSON metadata for the
// struct [VersionBindingsWorkersBindingKindAssets]
type versionBindingsWorkersBindingKindAssetsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindAssets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindAssetsJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindAssets) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindAssetsType string

const (
	VersionBindingsWorkersBindingKindAssetsTypeAssets VersionBindingsWorkersBindingKindAssetsType = "assets"
)

func (r VersionBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindBrowserType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindBrowserJSON `json:"-"`
}

// versionBindingsWorkersBindingKindBrowserJSON contains the JSON metadata for the
// struct [VersionBindingsWorkersBindingKindBrowser]
type versionBindingsWorkersBindingKindBrowserJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindBrowserJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindBrowser) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindBrowserType string

const (
	VersionBindingsWorkersBindingKindBrowserTypeBrowser VersionBindingsWorkersBindingKindBrowserType = "browser"
)

func (r VersionBindingsWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	DatabaseID string `json:"database_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindD1Type `json:"type" api:"required"`
	// Identifier of the D1 database to bind to.
	//
	// Deprecated: This property has been renamed to `database_id`.
	ID   string                                  `json:"id"`
	JSON versionBindingsWorkersBindingKindD1JSON `json:"-"`
}

// versionBindingsWorkersBindingKindD1JSON contains the JSON metadata for the
// struct [VersionBindingsWorkersBindingKindD1]
type versionBindingsWorkersBindingKindD1JSON struct {
	DatabaseID  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindD1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindD1JSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindD1) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindD1Type string

const (
	VersionBindingsWorkersBindingKindD1TypeD1 VersionBindingsWorkersBindingKindD1Type = "d1"
)

func (r VersionBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindDataBlob struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The name of the file containing the data content. Only accepted for
	// `service worker syntax` Workers.
	Part string `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type VersionBindingsWorkersBindingKindDataBlobType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindDataBlobJSON `json:"-"`
}

// versionBindingsWorkersBindingKindDataBlobJSON contains the JSON metadata for the
// struct [VersionBindingsWorkersBindingKindDataBlob]
type versionBindingsWorkersBindingKindDataBlobJSON struct {
	Name        apijson.Field
	Part        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindDataBlob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindDataBlobJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindDataBlob) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindDataBlobType string

const (
	VersionBindingsWorkersBindingKindDataBlobTypeDataBlob VersionBindingsWorkersBindingKindDataBlobType = "data_blob"
)

func (r VersionBindingsWorkersBindingKindDataBlobType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindDataBlobTypeDataBlob:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The name of the dispatch namespace.
	Namespace string `json:"namespace" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindDispatchNamespaceType `json:"type" api:"required"`
	// Outbound worker.
	Outbound VersionBindingsWorkersBindingKindDispatchNamespaceOutbound `json:"outbound"`
	JSON     versionBindingsWorkersBindingKindDispatchNamespaceJSON     `json:"-"`
}

// versionBindingsWorkersBindingKindDispatchNamespaceJSON contains the JSON
// metadata for the struct [VersionBindingsWorkersBindingKindDispatchNamespace]
type versionBindingsWorkersBindingKindDispatchNamespaceJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindDispatchNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindDispatchNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindDispatchNamespace) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindDispatchNamespaceType string

const (
	VersionBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace VersionBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r VersionBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type VersionBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params []VersionBindingsWorkersBindingKindDispatchNamespaceOutboundParam `json:"params"`
	// Outbound worker.
	Worker VersionBindingsWorkersBindingKindDispatchNamespaceOutboundWorker `json:"worker"`
	JSON   versionBindingsWorkersBindingKindDispatchNamespaceOutboundJSON   `json:"-"`
}

// versionBindingsWorkersBindingKindDispatchNamespaceOutboundJSON contains the JSON
// metadata for the struct
// [VersionBindingsWorkersBindingKindDispatchNamespaceOutbound]
type versionBindingsWorkersBindingKindDispatchNamespaceOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindDispatchNamespaceOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindDispatchNamespaceOutboundJSON) RawJSON() string {
	return r.raw
}

type VersionBindingsWorkersBindingKindDispatchNamespaceOutboundParam struct {
	// Name of the parameter.
	Name string                                                              `json:"name" api:"required"`
	JSON versionBindingsWorkersBindingKindDispatchNamespaceOutboundParamJSON `json:"-"`
}

// versionBindingsWorkersBindingKindDispatchNamespaceOutboundParamJSON contains the
// JSON metadata for the struct
// [VersionBindingsWorkersBindingKindDispatchNamespaceOutboundParam]
type versionBindingsWorkersBindingKindDispatchNamespaceOutboundParamJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindDispatchNamespaceOutboundParam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindDispatchNamespaceOutboundParamJSON) RawJSON() string {
	return r.raw
}

// Outbound worker.
type VersionBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Entrypoint to invoke on the outbound worker.
	Entrypoint string `json:"entrypoint"`
	// Environment of the outbound worker.
	Environment string `json:"environment"`
	// Name of the outbound worker.
	Service string                                                               `json:"service"`
	JSON    versionBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON `json:"-"`
}

// versionBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON contains
// the JSON metadata for the struct
// [VersionBindingsWorkersBindingKindDispatchNamespaceOutboundWorker]
type versionBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON struct {
	Entrypoint  apijson.Field
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type VersionBindingsWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindDurableObjectNamespaceType `json:"type" api:"required"`
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
	ScriptName string                                                      `json:"script_name"`
	JSON       versionBindingsWorkersBindingKindDurableObjectNamespaceJSON `json:"-"`
}

// versionBindingsWorkersBindingKindDurableObjectNamespaceJSON contains the JSON
// metadata for the struct
// [VersionBindingsWorkersBindingKindDurableObjectNamespace]
type versionBindingsWorkersBindingKindDurableObjectNamespaceJSON struct {
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

func (r *VersionBindingsWorkersBindingKindDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindDurableObjectNamespace) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	VersionBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace VersionBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r VersionBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID string `json:"id" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindHyperdriveType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindHyperdriveJSON `json:"-"`
}

// versionBindingsWorkersBindingKindHyperdriveJSON contains the JSON metadata for
// the struct [VersionBindingsWorkersBindingKindHyperdrive]
type versionBindingsWorkersBindingKindHyperdriveJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindHyperdrive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindHyperdriveJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindHyperdrive) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindHyperdriveType string

const (
	VersionBindingsWorkersBindingKindHyperdriveTypeHyperdrive VersionBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r VersionBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindInherit struct {
	// The name of the inherited binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindInheritType `json:"type" api:"required"`
	// The old name of the inherited binding. If set, the binding will be renamed from
	// `old_name` to `name` in the new version. If not set, the binding will keep the
	// same name between versions.
	OldName string `json:"old_name"`
	// Identifier for the version to inherit the binding from, which can be the version
	// ID or the literal "latest" to inherit from the latest version. Defaults to
	// inheriting the binding from the latest version.
	VersionID string                                       `json:"version_id"`
	JSON      versionBindingsWorkersBindingKindInheritJSON `json:"-"`
}

// versionBindingsWorkersBindingKindInheritJSON contains the JSON metadata for the
// struct [VersionBindingsWorkersBindingKindInherit]
type versionBindingsWorkersBindingKindInheritJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	OldName     apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindInherit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindInheritJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindInherit) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindInheritType string

const (
	VersionBindingsWorkersBindingKindInheritTypeInherit VersionBindingsWorkersBindingKindInheritType = "inherit"
)

func (r VersionBindingsWorkersBindingKindInheritType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindInheritTypeInherit:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindImages struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindImagesType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindImagesJSON `json:"-"`
}

// versionBindingsWorkersBindingKindImagesJSON contains the JSON metadata for the
// struct [VersionBindingsWorkersBindingKindImages]
type versionBindingsWorkersBindingKindImagesJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindImages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindImagesJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindImages) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindImagesType string

const (
	VersionBindingsWorkersBindingKindImagesTypeImages VersionBindingsWorkersBindingKindImagesType = "images"
)

func (r VersionBindingsWorkersBindingKindImagesType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindImagesTypeImages:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json interface{} `json:"json" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindJsonType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindJsonJSON `json:"-"`
}

// versionBindingsWorkersBindingKindJsonJSON contains the JSON metadata for the
// struct [VersionBindingsWorkersBindingKindJson]
type versionBindingsWorkersBindingKindJsonJSON struct {
	Json        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindJsonJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindJson) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindJsonType string

const (
	VersionBindingsWorkersBindingKindJsonTypeJson VersionBindingsWorkersBindingKindJsonType = "json"
)

func (r VersionBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindKVNamespaceType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindKVNamespaceJSON `json:"-"`
}

// versionBindingsWorkersBindingKindKVNamespaceJSON contains the JSON metadata for
// the struct [VersionBindingsWorkersBindingKindKVNamespace]
type versionBindingsWorkersBindingKindKVNamespaceJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindKVNamespaceJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindKVNamespace) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindKVNamespaceType string

const (
	VersionBindingsWorkersBindingKindKVNamespaceTypeKVNamespace VersionBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r VersionBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindMedia struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindMediaType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindMediaJSON `json:"-"`
}

// versionBindingsWorkersBindingKindMediaJSON contains the JSON metadata for the
// struct [VersionBindingsWorkersBindingKindMedia]
type versionBindingsWorkersBindingKindMediaJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindMedia) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindMediaJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindMedia) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindMediaType string

const (
	VersionBindingsWorkersBindingKindMediaTypeMedia VersionBindingsWorkersBindingKindMediaType = "media"
)

func (r VersionBindingsWorkersBindingKindMediaType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindMediaTypeMedia:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindMTLSCertificateType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindMTLSCertificateJSON `json:"-"`
}

// versionBindingsWorkersBindingKindMTLSCertificateJSON contains the JSON metadata
// for the struct [VersionBindingsWorkersBindingKindMTLSCertificate]
type versionBindingsWorkersBindingKindMTLSCertificateJSON struct {
	CertificateID apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindMTLSCertificate) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindMTLSCertificateType string

const (
	VersionBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate VersionBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r VersionBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The text value to use.
	Text string `json:"text" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindPlainTextType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindPlainTextJSON `json:"-"`
}

// versionBindingsWorkersBindingKindPlainTextJSON contains the JSON metadata for
// the struct [VersionBindingsWorkersBindingKindPlainText]
type versionBindingsWorkersBindingKindPlainTextJSON struct {
	Name        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindPlainText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindPlainTextJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindPlainText) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindPlainTextType string

const (
	VersionBindingsWorkersBindingKindPlainTextTypePlainText VersionBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r VersionBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Name of the Pipeline to bind to.
	Pipeline string `json:"pipeline" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindPipelinesType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindPipelinesJSON `json:"-"`
}

// versionBindingsWorkersBindingKindPipelinesJSON contains the JSON metadata for
// the struct [VersionBindingsWorkersBindingKindPipelines]
type versionBindingsWorkersBindingKindPipelinesJSON struct {
	Name        apijson.Field
	Pipeline    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindPipelines) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindPipelinesJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindPipelines) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindPipelinesType string

const (
	VersionBindingsWorkersBindingKindPipelinesTypePipelines VersionBindingsWorkersBindingKindPipelinesType = "pipelines"
)

func (r VersionBindingsWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindQueueType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindQueueJSON `json:"-"`
}

// versionBindingsWorkersBindingKindQueueJSON contains the JSON metadata for the
// struct [VersionBindingsWorkersBindingKindQueue]
type versionBindingsWorkersBindingKindQueueJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindQueueJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindQueue) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindQueueType string

const (
	VersionBindingsWorkersBindingKindQueueTypeQueue VersionBindingsWorkersBindingKindQueueType = "queue"
)

func (r VersionBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindRatelimit struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Identifier of the rate limit namespace to bind to.
	NamespaceID string `json:"namespace_id" api:"required"`
	// The rate limit configuration.
	Simple VersionBindingsWorkersBindingKindRatelimitSimple `json:"simple" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindRatelimitType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindRatelimitJSON `json:"-"`
}

// versionBindingsWorkersBindingKindRatelimitJSON contains the JSON metadata for
// the struct [VersionBindingsWorkersBindingKindRatelimit]
type versionBindingsWorkersBindingKindRatelimitJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Simple      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindRatelimitJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindRatelimit) implementsVersionBinding() {}

// The rate limit configuration.
type VersionBindingsWorkersBindingKindRatelimitSimple struct {
	// The limit (requests per period).
	Limit float64 `json:"limit" api:"required"`
	// The period in seconds.
	Period int64 `json:"period" api:"required"`
	// Duration in seconds to apply the mitigation action after the rate limit is
	// exceeded. Valid values are 0 (disabled), 10, or multiples of 60 up to 86400.
	// Must be greater than or equal to the period when non-zero.
	MitigationTimeout int64                                                `json:"mitigation_timeout"`
	JSON              versionBindingsWorkersBindingKindRatelimitSimpleJSON `json:"-"`
}

// versionBindingsWorkersBindingKindRatelimitSimpleJSON contains the JSON metadata
// for the struct [VersionBindingsWorkersBindingKindRatelimitSimple]
type versionBindingsWorkersBindingKindRatelimitSimpleJSON struct {
	Limit             apijson.Field
	Period            apijson.Field
	MitigationTimeout apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindRatelimitSimple) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindRatelimitSimpleJSON) RawJSON() string {
	return r.raw
}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindRatelimitType string

const (
	VersionBindingsWorkersBindingKindRatelimitTypeRatelimit VersionBindingsWorkersBindingKindRatelimitType = "ratelimit"
)

func (r VersionBindingsWorkersBindingKindRatelimitType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindRatelimitTypeRatelimit:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName string `json:"bucket_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindR2BucketType `json:"type" api:"required"`
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction VersionBindingsWorkersBindingKindR2BucketJurisdiction `json:"jurisdiction"`
	JSON         versionBindingsWorkersBindingKindR2BucketJSON         `json:"-"`
}

// versionBindingsWorkersBindingKindR2BucketJSON contains the JSON metadata for the
// struct [VersionBindingsWorkersBindingKindR2Bucket]
type versionBindingsWorkersBindingKindR2BucketJSON struct {
	BucketName   apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	Jurisdiction apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindR2BucketJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindR2Bucket) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindR2BucketType string

const (
	VersionBindingsWorkersBindingKindR2BucketTypeR2Bucket VersionBindingsWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r VersionBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

// The
// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
// of the R2 bucket.
type VersionBindingsWorkersBindingKindR2BucketJurisdiction string

const (
	VersionBindingsWorkersBindingKindR2BucketJurisdictionEu          VersionBindingsWorkersBindingKindR2BucketJurisdiction = "eu"
	VersionBindingsWorkersBindingKindR2BucketJurisdictionFedramp     VersionBindingsWorkersBindingKindR2BucketJurisdiction = "fedramp"
	VersionBindingsWorkersBindingKindR2BucketJurisdictionFedrampHigh VersionBindingsWorkersBindingKindR2BucketJurisdiction = "fedramp-high"
)

func (r VersionBindingsWorkersBindingKindR2BucketJurisdiction) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindR2BucketJurisdictionEu, VersionBindingsWorkersBindingKindR2BucketJurisdictionFedramp, VersionBindingsWorkersBindingKindR2BucketJurisdictionFedrampHigh:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindSecretTextType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindSecretTextJSON `json:"-"`
}

// versionBindingsWorkersBindingKindSecretTextJSON contains the JSON metadata for
// the struct [VersionBindingsWorkersBindingKindSecretText]
type versionBindingsWorkersBindingKindSecretTextJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindSecretText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindSecretTextJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindSecretText) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindSecretTextType string

const (
	VersionBindingsWorkersBindingKindSecretTextTypeSecretText VersionBindingsWorkersBindingKindSecretTextType = "secret_text"
)

func (r VersionBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindSendEmail struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindSendEmailType `json:"type" api:"required"`
	// List of allowed destination addresses.
	AllowedDestinationAddresses []string `json:"allowed_destination_addresses" format:"email"`
	// List of allowed sender addresses.
	AllowedSenderAddresses []string `json:"allowed_sender_addresses" format:"email"`
	// Destination address for the email.
	DestinationAddress string                                         `json:"destination_address" format:"email"`
	JSON               versionBindingsWorkersBindingKindSendEmailJSON `json:"-"`
}

// versionBindingsWorkersBindingKindSendEmailJSON contains the JSON metadata for
// the struct [VersionBindingsWorkersBindingKindSendEmail]
type versionBindingsWorkersBindingKindSendEmailJSON struct {
	Name                        apijson.Field
	Type                        apijson.Field
	AllowedDestinationAddresses apijson.Field
	AllowedSenderAddresses      apijson.Field
	DestinationAddress          apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindSendEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindSendEmailJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindSendEmail) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindSendEmailType string

const (
	VersionBindingsWorkersBindingKindSendEmailTypeSendEmail VersionBindingsWorkersBindingKindSendEmailType = "send_email"
)

func (r VersionBindingsWorkersBindingKindSendEmailType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindSendEmailTypeSendEmail:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindService struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Name of Worker to bind to.
	Service string `json:"service" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindServiceType `json:"type" api:"required"`
	// Entrypoint to invoke on the target Worker.
	Entrypoint string `json:"entrypoint"`
	// Optional environment if the Worker utilizes one.
	Environment string                                       `json:"environment"`
	JSON        versionBindingsWorkersBindingKindServiceJSON `json:"-"`
}

// versionBindingsWorkersBindingKindServiceJSON contains the JSON metadata for the
// struct [VersionBindingsWorkersBindingKindService]
type versionBindingsWorkersBindingKindServiceJSON struct {
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	Entrypoint  apijson.Field
	Environment apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindServiceJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindService) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindServiceType string

const (
	VersionBindingsWorkersBindingKindServiceTypeService VersionBindingsWorkersBindingKindServiceType = "service"
)

func (r VersionBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindTextBlob struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The name of the file containing the text content. Only accepted for
	// `service worker syntax` Workers.
	Part string `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type VersionBindingsWorkersBindingKindTextBlobType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindTextBlobJSON `json:"-"`
}

// versionBindingsWorkersBindingKindTextBlobJSON contains the JSON metadata for the
// struct [VersionBindingsWorkersBindingKindTextBlob]
type versionBindingsWorkersBindingKindTextBlobJSON struct {
	Name        apijson.Field
	Part        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindTextBlob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindTextBlobJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindTextBlob) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindTextBlobType string

const (
	VersionBindingsWorkersBindingKindTextBlobTypeTextBlob VersionBindingsWorkersBindingKindTextBlobType = "text_blob"
)

func (r VersionBindingsWorkersBindingKindTextBlobType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindTextBlobTypeTextBlob:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindVectorizeType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindVectorizeJSON `json:"-"`
}

// versionBindingsWorkersBindingKindVectorizeJSON contains the JSON metadata for
// the struct [VersionBindingsWorkersBindingKindVectorize]
type versionBindingsWorkersBindingKindVectorizeJSON struct {
	IndexName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindVectorizeJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindVectorize) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindVectorizeType string

const (
	VersionBindingsWorkersBindingKindVectorizeTypeVectorize VersionBindingsWorkersBindingKindVectorizeType = "vectorize"
)

func (r VersionBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindVersionMetadataType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindVersionMetadataJSON `json:"-"`
}

// versionBindingsWorkersBindingKindVersionMetadataJSON contains the JSON metadata
// for the struct [VersionBindingsWorkersBindingKindVersionMetadata]
type versionBindingsWorkersBindingKindVersionMetadataJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindVersionMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindVersionMetadataJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindVersionMetadata) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindVersionMetadataType string

const (
	VersionBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata VersionBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r VersionBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Name of the secret in the store.
	SecretName string `json:"secret_name" api:"required"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindSecretsStoreSecretType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindSecretsStoreSecretJSON `json:"-"`
}

// versionBindingsWorkersBindingKindSecretsStoreSecretJSON contains the JSON
// metadata for the struct [VersionBindingsWorkersBindingKindSecretsStoreSecret]
type versionBindingsWorkersBindingKindSecretsStoreSecretJSON struct {
	Name        apijson.Field
	SecretName  apijson.Field
	StoreID     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindSecretsStoreSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindSecretsStoreSecretJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindSecretsStoreSecret) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindSecretsStoreSecretType string

const (
	VersionBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret VersionBindingsWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r VersionBindingsWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindFlagship struct {
	// ID of the Flagship app to bind to for feature flag evaluation.
	AppID string `json:"app_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindFlagshipType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindFlagshipJSON `json:"-"`
}

// versionBindingsWorkersBindingKindFlagshipJSON contains the JSON metadata for the
// struct [VersionBindingsWorkersBindingKindFlagship]
type versionBindingsWorkersBindingKindFlagshipJSON struct {
	AppID       apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindFlagship) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindFlagshipJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindFlagship) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindFlagshipType string

const (
	VersionBindingsWorkersBindingKindFlagshipTypeFlagship VersionBindingsWorkersBindingKindFlagshipType = "flagship"
)

func (r VersionBindingsWorkersBindingKindFlagshipType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindFlagshipTypeFlagship:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm" api:"required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format VersionBindingsWorkersBindingKindSecretKeyFormat `json:"format" api:"required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindSecretKeyType `json:"type" api:"required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []VersionBindingsWorkersBindingKindSecretKeyUsage `json:"usages" api:"required"`
	JSON   versionBindingsWorkersBindingKindSecretKeyJSON    `json:"-"`
}

// versionBindingsWorkersBindingKindSecretKeyJSON contains the JSON metadata for
// the struct [VersionBindingsWorkersBindingKindSecretKey]
type versionBindingsWorkersBindingKindSecretKeyJSON struct {
	Algorithm   apijson.Field
	Format      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Usages      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindSecretKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindSecretKeyJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindSecretKey) implementsVersionBinding() {}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type VersionBindingsWorkersBindingKindSecretKeyFormat string

const (
	VersionBindingsWorkersBindingKindSecretKeyFormatRaw   VersionBindingsWorkersBindingKindSecretKeyFormat = "raw"
	VersionBindingsWorkersBindingKindSecretKeyFormatPkcs8 VersionBindingsWorkersBindingKindSecretKeyFormat = "pkcs8"
	VersionBindingsWorkersBindingKindSecretKeyFormatSpki  VersionBindingsWorkersBindingKindSecretKeyFormat = "spki"
	VersionBindingsWorkersBindingKindSecretKeyFormatJwk   VersionBindingsWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r VersionBindingsWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindSecretKeyFormatRaw, VersionBindingsWorkersBindingKindSecretKeyFormatPkcs8, VersionBindingsWorkersBindingKindSecretKeyFormatSpki, VersionBindingsWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindSecretKeyType string

const (
	VersionBindingsWorkersBindingKindSecretKeyTypeSecretKey VersionBindingsWorkersBindingKindSecretKeyType = "secret_key"
)

func (r VersionBindingsWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindSecretKeyUsage string

const (
	VersionBindingsWorkersBindingKindSecretKeyUsageEncrypt    VersionBindingsWorkersBindingKindSecretKeyUsage = "encrypt"
	VersionBindingsWorkersBindingKindSecretKeyUsageDecrypt    VersionBindingsWorkersBindingKindSecretKeyUsage = "decrypt"
	VersionBindingsWorkersBindingKindSecretKeyUsageSign       VersionBindingsWorkersBindingKindSecretKeyUsage = "sign"
	VersionBindingsWorkersBindingKindSecretKeyUsageVerify     VersionBindingsWorkersBindingKindSecretKeyUsage = "verify"
	VersionBindingsWorkersBindingKindSecretKeyUsageDeriveKey  VersionBindingsWorkersBindingKindSecretKeyUsage = "deriveKey"
	VersionBindingsWorkersBindingKindSecretKeyUsageDeriveBits VersionBindingsWorkersBindingKindSecretKeyUsage = "deriveBits"
	VersionBindingsWorkersBindingKindSecretKeyUsageWrapKey    VersionBindingsWorkersBindingKindSecretKeyUsage = "wrapKey"
	VersionBindingsWorkersBindingKindSecretKeyUsageUnwrapKey  VersionBindingsWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r VersionBindingsWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindSecretKeyUsageEncrypt, VersionBindingsWorkersBindingKindSecretKeyUsageDecrypt, VersionBindingsWorkersBindingKindSecretKeyUsageSign, VersionBindingsWorkersBindingKindSecretKeyUsageVerify, VersionBindingsWorkersBindingKindSecretKeyUsageDeriveKey, VersionBindingsWorkersBindingKindSecretKeyUsageDeriveBits, VersionBindingsWorkersBindingKindSecretKeyUsageWrapKey, VersionBindingsWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindWorkflowType `json:"type" api:"required"`
	// Name of the Workflow to bind to.
	WorkflowName string `json:"workflow_name" api:"required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName string `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName string                                        `json:"script_name"`
	JSON       versionBindingsWorkersBindingKindWorkflowJSON `json:"-"`
}

// versionBindingsWorkersBindingKindWorkflowJSON contains the JSON metadata for the
// struct [VersionBindingsWorkersBindingKindWorkflow]
type versionBindingsWorkersBindingKindWorkflowJSON struct {
	Name         apijson.Field
	Type         apijson.Field
	WorkflowName apijson.Field
	ClassName    apijson.Field
	ScriptName   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindWorkflowJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindWorkflow) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindWorkflowType string

const (
	VersionBindingsWorkersBindingKindWorkflowTypeWorkflow VersionBindingsWorkersBindingKindWorkflowType = "workflow"
)

func (r VersionBindingsWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindWasmModule struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The name of the file containing the WebAssembly module content. Only accepted
	// for `service worker syntax` Workers.
	Part string `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type VersionBindingsWorkersBindingKindWasmModuleType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindWasmModuleJSON `json:"-"`
}

// versionBindingsWorkersBindingKindWasmModuleJSON contains the JSON metadata for
// the struct [VersionBindingsWorkersBindingKindWasmModule]
type versionBindingsWorkersBindingKindWasmModuleJSON struct {
	Name        apijson.Field
	Part        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindWasmModule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindWasmModuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindWasmModule) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindWasmModuleType string

const (
	VersionBindingsWorkersBindingKindWasmModuleTypeWasmModule VersionBindingsWorkersBindingKindWasmModuleType = "wasm_module"
)

func (r VersionBindingsWorkersBindingKindWasmModuleType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindWasmModuleTypeWasmModule:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindVPCService struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// Identifier of the VPC service to bind to.
	ServiceID string `json:"service_id" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindVPCServiceType `json:"type" api:"required"`
	JSON versionBindingsWorkersBindingKindVPCServiceJSON `json:"-"`
}

// versionBindingsWorkersBindingKindVPCServiceJSON contains the JSON metadata for
// the struct [VersionBindingsWorkersBindingKindVPCService]
type versionBindingsWorkersBindingKindVPCServiceJSON struct {
	Name        apijson.Field
	ServiceID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindVPCService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindVPCServiceJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindVPCService) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindVPCServiceType string

const (
	VersionBindingsWorkersBindingKindVPCServiceTypeVPCService VersionBindingsWorkersBindingKindVPCServiceType = "vpc_service"
)

func (r VersionBindingsWorkersBindingKindVPCServiceType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindVPCServiceTypeVPCService:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindVPCNetwork struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindVPCNetworkType `json:"type" api:"required"`
	// Identifier of the network to bind to. Only "cf1:network" is currently supported.
	// Mutually exclusive with tunnel_id.
	NetworkID string `json:"network_id"`
	// UUID of the Cloudflare Tunnel to bind to. Mutually exclusive with network_id.
	TunnelID string                                          `json:"tunnel_id"`
	JSON     versionBindingsWorkersBindingKindVPCNetworkJSON `json:"-"`
}

// versionBindingsWorkersBindingKindVPCNetworkJSON contains the JSON metadata for
// the struct [VersionBindingsWorkersBindingKindVPCNetwork]
type versionBindingsWorkersBindingKindVPCNetworkJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	NetworkID   apijson.Field
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionBindingsWorkersBindingKindVPCNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionBindingsWorkersBindingKindVPCNetworkJSON) RawJSON() string {
	return r.raw
}

func (r VersionBindingsWorkersBindingKindVPCNetwork) implementsVersionBinding() {}

// The kind of resource that the binding provides.
type VersionBindingsWorkersBindingKindVPCNetworkType string

const (
	VersionBindingsWorkersBindingKindVPCNetworkTypeVPCNetwork VersionBindingsWorkersBindingKindVPCNetworkType = "vpc_network"
)

func (r VersionBindingsWorkersBindingKindVPCNetworkType) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindVPCNetworkTypeVPCNetwork:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type VersionBindingsType string

const (
	VersionBindingsTypeAI                     VersionBindingsType = "ai"
	VersionBindingsTypeAISearch               VersionBindingsType = "ai_search"
	VersionBindingsTypeAISearchNamespace      VersionBindingsType = "ai_search_namespace"
	VersionBindingsTypeAnalyticsEngine        VersionBindingsType = "analytics_engine"
	VersionBindingsTypeAssets                 VersionBindingsType = "assets"
	VersionBindingsTypeBrowser                VersionBindingsType = "browser"
	VersionBindingsTypeD1                     VersionBindingsType = "d1"
	VersionBindingsTypeDataBlob               VersionBindingsType = "data_blob"
	VersionBindingsTypeDispatchNamespace      VersionBindingsType = "dispatch_namespace"
	VersionBindingsTypeDurableObjectNamespace VersionBindingsType = "durable_object_namespace"
	VersionBindingsTypeHyperdrive             VersionBindingsType = "hyperdrive"
	VersionBindingsTypeInherit                VersionBindingsType = "inherit"
	VersionBindingsTypeImages                 VersionBindingsType = "images"
	VersionBindingsTypeJson                   VersionBindingsType = "json"
	VersionBindingsTypeKVNamespace            VersionBindingsType = "kv_namespace"
	VersionBindingsTypeMedia                  VersionBindingsType = "media"
	VersionBindingsTypeMTLSCertificate        VersionBindingsType = "mtls_certificate"
	VersionBindingsTypePlainText              VersionBindingsType = "plain_text"
	VersionBindingsTypePipelines              VersionBindingsType = "pipelines"
	VersionBindingsTypeQueue                  VersionBindingsType = "queue"
	VersionBindingsTypeRatelimit              VersionBindingsType = "ratelimit"
	VersionBindingsTypeR2Bucket               VersionBindingsType = "r2_bucket"
	VersionBindingsTypeSecretText             VersionBindingsType = "secret_text"
	VersionBindingsTypeSendEmail              VersionBindingsType = "send_email"
	VersionBindingsTypeService                VersionBindingsType = "service"
	VersionBindingsTypeTextBlob               VersionBindingsType = "text_blob"
	VersionBindingsTypeVectorize              VersionBindingsType = "vectorize"
	VersionBindingsTypeVersionMetadata        VersionBindingsType = "version_metadata"
	VersionBindingsTypeSecretsStoreSecret     VersionBindingsType = "secrets_store_secret"
	VersionBindingsTypeFlagship               VersionBindingsType = "flagship"
	VersionBindingsTypeSecretKey              VersionBindingsType = "secret_key"
	VersionBindingsTypeWorkflow               VersionBindingsType = "workflow"
	VersionBindingsTypeWasmModule             VersionBindingsType = "wasm_module"
	VersionBindingsTypeVPCService             VersionBindingsType = "vpc_service"
	VersionBindingsTypeVPCNetwork             VersionBindingsType = "vpc_network"
)

func (r VersionBindingsType) IsKnown() bool {
	switch r {
	case VersionBindingsTypeAI, VersionBindingsTypeAISearch, VersionBindingsTypeAISearchNamespace, VersionBindingsTypeAnalyticsEngine, VersionBindingsTypeAssets, VersionBindingsTypeBrowser, VersionBindingsTypeD1, VersionBindingsTypeDataBlob, VersionBindingsTypeDispatchNamespace, VersionBindingsTypeDurableObjectNamespace, VersionBindingsTypeHyperdrive, VersionBindingsTypeInherit, VersionBindingsTypeImages, VersionBindingsTypeJson, VersionBindingsTypeKVNamespace, VersionBindingsTypeMedia, VersionBindingsTypeMTLSCertificate, VersionBindingsTypePlainText, VersionBindingsTypePipelines, VersionBindingsTypeQueue, VersionBindingsTypeRatelimit, VersionBindingsTypeR2Bucket, VersionBindingsTypeSecretText, VersionBindingsTypeSendEmail, VersionBindingsTypeService, VersionBindingsTypeTextBlob, VersionBindingsTypeVectorize, VersionBindingsTypeVersionMetadata, VersionBindingsTypeSecretsStoreSecret, VersionBindingsTypeFlagship, VersionBindingsTypeSecretKey, VersionBindingsTypeWorkflow, VersionBindingsTypeWasmModule, VersionBindingsTypeVPCService, VersionBindingsTypeVPCNetwork:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type VersionBindingsFormat string

const (
	VersionBindingsFormatRaw   VersionBindingsFormat = "raw"
	VersionBindingsFormatPkcs8 VersionBindingsFormat = "pkcs8"
	VersionBindingsFormatSpki  VersionBindingsFormat = "spki"
	VersionBindingsFormatJwk   VersionBindingsFormat = "jwk"
)

func (r VersionBindingsFormat) IsKnown() bool {
	switch r {
	case VersionBindingsFormatRaw, VersionBindingsFormatPkcs8, VersionBindingsFormatSpki, VersionBindingsFormatJwk:
		return true
	}
	return false
}

// The
// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
// of the R2 bucket.
type VersionBindingsJurisdiction string

const (
	VersionBindingsJurisdictionEu          VersionBindingsJurisdiction = "eu"
	VersionBindingsJurisdictionFedramp     VersionBindingsJurisdiction = "fedramp"
	VersionBindingsJurisdictionFedrampHigh VersionBindingsJurisdiction = "fedramp-high"
)

func (r VersionBindingsJurisdiction) IsKnown() bool {
	switch r {
	case VersionBindingsJurisdictionEu, VersionBindingsJurisdictionFedramp, VersionBindingsJurisdictionFedrampHigh:
		return true
	}
	return false
}

// Global CacheW configuration for the Worker. When caching is on, the platform
// provisions a `cloudflare.app` zone for the Worker. A `type: worker` entry in the
// `exports` map can override this value for a single entrypoint.
type VersionCacheOptions struct {
	// Whether caching is enabled for this Worker.
	Enabled bool `json:"enabled" api:"required"`
	// Whether cached responses are shared across Worker version uploads. This is
	// independent of `enabled`. It can stay true while caching is off, so the
	// preference survives turning caching off and back on.
	CrossVersionCache bool                    `json:"cross_version_cache"`
	JSON              versionCacheOptionsJSON `json:"-"`
}

// versionCacheOptionsJSON contains the JSON metadata for the struct
// [VersionCacheOptions]
type versionCacheOptionsJSON struct {
	Enabled           apijson.Field
	CrossVersionCache apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *VersionCacheOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionCacheOptionsJSON) RawJSON() string {
	return r.raw
}

// Container configuration for a Worker.
type VersionContainer struct {
	// Select which Durable Object class should get this container attached.
	ClassName string               `json:"class_name" api:"required"`
	JSON      versionContainerJSON `json:"-"`
}

// versionContainerJSON contains the JSON metadata for the struct
// [VersionContainer]
type versionContainerJSON struct {
	ClassName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionContainer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionContainerJSON) RawJSON() string {
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
type VersionExport struct {
	// Marks this entry as a Worker entrypoint export.
	Type VersionExportsType `json:"type" api:"required"`
	// This field can have the runtime type of
	// [VersionExportsWorkersWorkerExportCache].
	Cache interface{} `json:"cache"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container string `json:"container"`
	// Live export. May be omitted; defaults to `created`.
	State VersionExportsState `json:"state"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage VersionExportsStorage `json:"storage"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom string            `json:"transfer_from"`
	JSON         versionExportJSON `json:"-"`
	union        VersionExportsUnion
}

// versionExportJSON contains the JSON metadata for the struct [VersionExport]
type versionExportJSON struct {
	Type         apijson.Field
	Cache        apijson.Field
	Container    apijson.Field
	State        apijson.Field
	Storage      apijson.Field
	TransferFrom apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r versionExportJSON) RawJSON() string {
	return r.raw
}

func (r *VersionExport) UnmarshalJSON(data []byte) (err error) {
	*r = VersionExport{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [VersionExportsUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [VersionExportsWorkersWorkerExport],
// [VersionExportsWorkersDurableObjectExport],
// [VersionExportsWorkersDurableObjectDeletedExport],
// [VersionExportsWorkersDurableObjectRenamedExport],
// [VersionExportsWorkersDurableObjectTransferredExport],
// [VersionExportsWorkersDurableObjectExpectingTransferExport].
func (r VersionExport) AsUnion() VersionExportsUnion {
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
// Union satisfied by [VersionExportsWorkersWorkerExport],
// [VersionExportsWorkersDurableObjectExport],
// [VersionExportsWorkersDurableObjectDeletedExport],
// [VersionExportsWorkersDurableObjectRenamedExport],
// [VersionExportsWorkersDurableObjectTransferredExport] or
// [VersionExportsWorkersDurableObjectExpectingTransferExport].
type VersionExportsUnion interface {
	implementsVersionExport()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionExportsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionExportsWorkersWorkerExport{}),
			DiscriminatorValue: "worker",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionExportsWorkersDurableObjectExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionExportsWorkersDurableObjectDeletedExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionExportsWorkersDurableObjectRenamedExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionExportsWorkersDurableObjectTransferredExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionExportsWorkersDurableObjectExpectingTransferExport{}),
			DiscriminatorValue: "durable-object",
		},
	)
}

// A named Worker entrypoint export (`type: worker`). Worker entrypoints are always
// live (`state: created`) and carry no storage or lifecycle fields. The optional
// `cache` block overrides the Worker's global `cache_options.enabled` for this
// entrypoint.
type VersionExportsWorkersWorkerExport struct {
	// Marks this entry as a Worker entrypoint export.
	Type VersionExportsWorkersWorkerExportType `json:"type" api:"required"`
	// Cache override for this entrypoint. Overrides the Worker's global
	// `cache_options.enabled` for this entrypoint only.
	Cache VersionExportsWorkersWorkerExportCache `json:"cache"`
	// Live export. May be omitted; defaults to `created`.
	State VersionExportsWorkersWorkerExportState `json:"state"`
	JSON  versionExportsWorkersWorkerExportJSON  `json:"-"`
}

// versionExportsWorkersWorkerExportJSON contains the JSON metadata for the struct
// [VersionExportsWorkersWorkerExport]
type versionExportsWorkersWorkerExportJSON struct {
	Type        apijson.Field
	Cache       apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionExportsWorkersWorkerExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionExportsWorkersWorkerExportJSON) RawJSON() string {
	return r.raw
}

func (r VersionExportsWorkersWorkerExport) implementsVersionExport() {}

// Marks this entry as a Worker entrypoint export.
type VersionExportsWorkersWorkerExportType string

const (
	VersionExportsWorkersWorkerExportTypeWorker VersionExportsWorkersWorkerExportType = "worker"
)

func (r VersionExportsWorkersWorkerExportType) IsKnown() bool {
	switch r {
	case VersionExportsWorkersWorkerExportTypeWorker:
		return true
	}
	return false
}

// Cache override for this entrypoint. Overrides the Worker's global
// `cache_options.enabled` for this entrypoint only.
type VersionExportsWorkersWorkerExportCache struct {
	// Whether caching is enabled for this entrypoint.
	Enabled bool                                       `json:"enabled" api:"required"`
	JSON    versionExportsWorkersWorkerExportCacheJSON `json:"-"`
}

// versionExportsWorkersWorkerExportCacheJSON contains the JSON metadata for the
// struct [VersionExportsWorkersWorkerExportCache]
type versionExportsWorkersWorkerExportCacheJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionExportsWorkersWorkerExportCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionExportsWorkersWorkerExportCacheJSON) RawJSON() string {
	return r.raw
}

// Live export. May be omitted; defaults to `created`.
type VersionExportsWorkersWorkerExportState string

const (
	VersionExportsWorkersWorkerExportStateCreated VersionExportsWorkersWorkerExportState = "created"
)

func (r VersionExportsWorkersWorkerExportState) IsKnown() bool {
	switch r {
	case VersionExportsWorkersWorkerExportStateCreated:
		return true
	}
	return false
}

// A live Durable Object export (`state: created`, the default). The platform
// auto-provisions the namespace on first deploy, matches it on subsequent deploys,
// and never mutates or deletes it as a side effect of a code-only change.
// `storage` is required; `renamed_to`, `transferred_to` and `transfer_from` are
// not allowed on a live entry.
type VersionExportsWorkersDurableObjectExport struct {
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage VersionExportsWorkersDurableObjectExportStorage `json:"storage" api:"required"`
	// Marks this entry as a Durable Object export.
	Type VersionExportsWorkersDurableObjectExportType `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container string `json:"container"`
	// Live export. May be omitted; defaults to `created`.
	State VersionExportsWorkersDurableObjectExportState `json:"state"`
	JSON  versionExportsWorkersDurableObjectExportJSON  `json:"-"`
}

// versionExportsWorkersDurableObjectExportJSON contains the JSON metadata for the
// struct [VersionExportsWorkersDurableObjectExport]
type versionExportsWorkersDurableObjectExportJSON struct {
	Storage     apijson.Field
	Type        apijson.Field
	Container   apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionExportsWorkersDurableObjectExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionExportsWorkersDurableObjectExportJSON) RawJSON() string {
	return r.raw
}

func (r VersionExportsWorkersDurableObjectExport) implementsVersionExport() {}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type VersionExportsWorkersDurableObjectExportStorage string

const (
	VersionExportsWorkersDurableObjectExportStorageSqlite   VersionExportsWorkersDurableObjectExportStorage = "sqlite"
	VersionExportsWorkersDurableObjectExportStorageLegacyKV VersionExportsWorkersDurableObjectExportStorage = "legacy-kv"
)

func (r VersionExportsWorkersDurableObjectExportStorage) IsKnown() bool {
	switch r {
	case VersionExportsWorkersDurableObjectExportStorageSqlite, VersionExportsWorkersDurableObjectExportStorageLegacyKV:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type VersionExportsWorkersDurableObjectExportType string

const (
	VersionExportsWorkersDurableObjectExportTypeDurableObject VersionExportsWorkersDurableObjectExportType = "durable-object"
)

func (r VersionExportsWorkersDurableObjectExportType) IsKnown() bool {
	switch r {
	case VersionExportsWorkersDurableObjectExportTypeDurableObject:
		return true
	}
	return false
}

// Live export. May be omitted; defaults to `created`.
type VersionExportsWorkersDurableObjectExportState string

const (
	VersionExportsWorkersDurableObjectExportStateCreated VersionExportsWorkersDurableObjectExportState = "created"
)

func (r VersionExportsWorkersDurableObjectExportState) IsKnown() bool {
	switch r {
	case VersionExportsWorkersDurableObjectExportStateCreated:
		return true
	}
	return false
}

// A `deleted` tombstone: retires the provisioned namespace for this class and all
// of its data. The class must be absent from the uploaded code and no other Worker
// in the account may bind to the namespace, otherwise the deploy is rejected. No
// other fields are allowed. Deletion is irreversible.
type VersionExportsWorkersDurableObjectDeletedExport struct {
	// Tombstone that deletes the namespace.
	State VersionExportsWorkersDurableObjectDeletedExportState `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type VersionExportsWorkersDurableObjectDeletedExportType `json:"type" api:"required"`
	JSON versionExportsWorkersDurableObjectDeletedExportJSON `json:"-"`
}

// versionExportsWorkersDurableObjectDeletedExportJSON contains the JSON metadata
// for the struct [VersionExportsWorkersDurableObjectDeletedExport]
type versionExportsWorkersDurableObjectDeletedExportJSON struct {
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionExportsWorkersDurableObjectDeletedExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionExportsWorkersDurableObjectDeletedExportJSON) RawJSON() string {
	return r.raw
}

func (r VersionExportsWorkersDurableObjectDeletedExport) implementsVersionExport() {}

// Tombstone that deletes the namespace.
type VersionExportsWorkersDurableObjectDeletedExportState string

const (
	VersionExportsWorkersDurableObjectDeletedExportStateDeleted VersionExportsWorkersDurableObjectDeletedExportState = "deleted"
)

func (r VersionExportsWorkersDurableObjectDeletedExportState) IsKnown() bool {
	switch r {
	case VersionExportsWorkersDurableObjectDeletedExportStateDeleted:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type VersionExportsWorkersDurableObjectDeletedExportType string

const (
	VersionExportsWorkersDurableObjectDeletedExportTypeDurableObject VersionExportsWorkersDurableObjectDeletedExportType = "durable-object"
)

func (r VersionExportsWorkersDurableObjectDeletedExportType) IsKnown() bool {
	switch r {
	case VersionExportsWorkersDurableObjectDeletedExportTypeDurableObject:
		return true
	}
	return false
}

// A `renamed` tombstone: rewrites the provisioned namespace's class name from this
// map key to `renamed_to`. The source class may stay in code during the rollout
// window (an info notice is emitted). `storage`, `transferred_to` and
// `transfer_from` are not allowed.
type VersionExportsWorkersDurableObjectRenamedExport struct {
	// Tombstone that renames the namespace's class.
	State VersionExportsWorkersDurableObjectRenamedExportState `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type VersionExportsWorkersDurableObjectRenamedExportType `json:"type" api:"required"`
	JSON versionExportsWorkersDurableObjectRenamedExportJSON `json:"-"`
}

// versionExportsWorkersDurableObjectRenamedExportJSON contains the JSON metadata
// for the struct [VersionExportsWorkersDurableObjectRenamedExport]
type versionExportsWorkersDurableObjectRenamedExportJSON struct {
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionExportsWorkersDurableObjectRenamedExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionExportsWorkersDurableObjectRenamedExportJSON) RawJSON() string {
	return r.raw
}

func (r VersionExportsWorkersDurableObjectRenamedExport) implementsVersionExport() {}

// Tombstone that renames the namespace's class.
type VersionExportsWorkersDurableObjectRenamedExportState string

const (
	VersionExportsWorkersDurableObjectRenamedExportStateRenamed VersionExportsWorkersDurableObjectRenamedExportState = "renamed"
)

func (r VersionExportsWorkersDurableObjectRenamedExportState) IsKnown() bool {
	switch r {
	case VersionExportsWorkersDurableObjectRenamedExportStateRenamed:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type VersionExportsWorkersDurableObjectRenamedExportType string

const (
	VersionExportsWorkersDurableObjectRenamedExportTypeDurableObject VersionExportsWorkersDurableObjectRenamedExportType = "durable-object"
)

func (r VersionExportsWorkersDurableObjectRenamedExportType) IsKnown() bool {
	switch r {
	case VersionExportsWorkersDurableObjectRenamedExportTypeDurableObject:
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
type VersionExportsWorkersDurableObjectTransferredExport struct {
	// Tombstone that transfers the namespace to another script.
	State VersionExportsWorkersDurableObjectTransferredExportState `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type VersionExportsWorkersDurableObjectTransferredExportType `json:"type" api:"required"`
	JSON versionExportsWorkersDurableObjectTransferredExportJSON `json:"-"`
}

// versionExportsWorkersDurableObjectTransferredExportJSON contains the JSON
// metadata for the struct [VersionExportsWorkersDurableObjectTransferredExport]
type versionExportsWorkersDurableObjectTransferredExportJSON struct {
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionExportsWorkersDurableObjectTransferredExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionExportsWorkersDurableObjectTransferredExportJSON) RawJSON() string {
	return r.raw
}

func (r VersionExportsWorkersDurableObjectTransferredExport) implementsVersionExport() {}

// Tombstone that transfers the namespace to another script.
type VersionExportsWorkersDurableObjectTransferredExportState string

const (
	VersionExportsWorkersDurableObjectTransferredExportStateTransferred VersionExportsWorkersDurableObjectTransferredExportState = "transferred"
)

func (r VersionExportsWorkersDurableObjectTransferredExportState) IsKnown() bool {
	switch r {
	case VersionExportsWorkersDurableObjectTransferredExportStateTransferred:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type VersionExportsWorkersDurableObjectTransferredExportType string

const (
	VersionExportsWorkersDurableObjectTransferredExportTypeDurableObject VersionExportsWorkersDurableObjectTransferredExportType = "durable-object"
)

func (r VersionExportsWorkersDurableObjectTransferredExportType) IsKnown() bool {
	switch r {
	case VersionExportsWorkersDurableObjectTransferredExportTypeDurableObject:
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
type VersionExportsWorkersDurableObjectExpectingTransferExport struct {
	// Target side of a two-phase transfer.
	State VersionExportsWorkersDurableObjectExpectingTransferExportState `json:"state" api:"required"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage VersionExportsWorkersDurableObjectExpectingTransferExportStorage `json:"storage" api:"required"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom string `json:"transfer_from" api:"required"`
	// Marks this entry as a Durable Object export.
	Type VersionExportsWorkersDurableObjectExpectingTransferExportType `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object once the transfer settles. Valid only on live entries.
	Container string                                                        `json:"container"`
	JSON      versionExportsWorkersDurableObjectExpectingTransferExportJSON `json:"-"`
}

// versionExportsWorkersDurableObjectExpectingTransferExportJSON contains the JSON
// metadata for the struct
// [VersionExportsWorkersDurableObjectExpectingTransferExport]
type versionExportsWorkersDurableObjectExpectingTransferExportJSON struct {
	State        apijson.Field
	Storage      apijson.Field
	TransferFrom apijson.Field
	Type         apijson.Field
	Container    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *VersionExportsWorkersDurableObjectExpectingTransferExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionExportsWorkersDurableObjectExpectingTransferExportJSON) RawJSON() string {
	return r.raw
}

func (r VersionExportsWorkersDurableObjectExpectingTransferExport) implementsVersionExport() {}

// Target side of a two-phase transfer.
type VersionExportsWorkersDurableObjectExpectingTransferExportState string

const (
	VersionExportsWorkersDurableObjectExpectingTransferExportStateExpectingTransfer VersionExportsWorkersDurableObjectExpectingTransferExportState = "expecting-transfer"
)

func (r VersionExportsWorkersDurableObjectExpectingTransferExportState) IsKnown() bool {
	switch r {
	case VersionExportsWorkersDurableObjectExpectingTransferExportStateExpectingTransfer:
		return true
	}
	return false
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type VersionExportsWorkersDurableObjectExpectingTransferExportStorage string

const (
	VersionExportsWorkersDurableObjectExpectingTransferExportStorageSqlite   VersionExportsWorkersDurableObjectExpectingTransferExportStorage = "sqlite"
	VersionExportsWorkersDurableObjectExpectingTransferExportStorageLegacyKV VersionExportsWorkersDurableObjectExpectingTransferExportStorage = "legacy-kv"
)

func (r VersionExportsWorkersDurableObjectExpectingTransferExportStorage) IsKnown() bool {
	switch r {
	case VersionExportsWorkersDurableObjectExpectingTransferExportStorageSqlite, VersionExportsWorkersDurableObjectExpectingTransferExportStorageLegacyKV:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type VersionExportsWorkersDurableObjectExpectingTransferExportType string

const (
	VersionExportsWorkersDurableObjectExpectingTransferExportTypeDurableObject VersionExportsWorkersDurableObjectExpectingTransferExportType = "durable-object"
)

func (r VersionExportsWorkersDurableObjectExpectingTransferExportType) IsKnown() bool {
	switch r {
	case VersionExportsWorkersDurableObjectExpectingTransferExportTypeDurableObject:
		return true
	}
	return false
}

// Marks this entry as a Worker entrypoint export.
type VersionExportsType string

const (
	VersionExportsTypeWorker        VersionExportsType = "worker"
	VersionExportsTypeDurableObject VersionExportsType = "durable-object"
)

func (r VersionExportsType) IsKnown() bool {
	switch r {
	case VersionExportsTypeWorker, VersionExportsTypeDurableObject:
		return true
	}
	return false
}

// Live export. May be omitted; defaults to `created`.
type VersionExportsState string

const (
	VersionExportsStateCreated           VersionExportsState = "created"
	VersionExportsStateDeleted           VersionExportsState = "deleted"
	VersionExportsStateRenamed           VersionExportsState = "renamed"
	VersionExportsStateTransferred       VersionExportsState = "transferred"
	VersionExportsStateExpectingTransfer VersionExportsState = "expecting-transfer"
)

func (r VersionExportsState) IsKnown() bool {
	switch r {
	case VersionExportsStateCreated, VersionExportsStateDeleted, VersionExportsStateRenamed, VersionExportsStateTransferred, VersionExportsStateExpectingTransfer:
		return true
	}
	return false
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type VersionExportsStorage string

const (
	VersionExportsStorageSqlite   VersionExportsStorage = "sqlite"
	VersionExportsStorageLegacyKV VersionExportsStorage = "legacy-kv"
)

func (r VersionExportsStorage) IsKnown() bool {
	switch r {
	case VersionExportsStorageSqlite, VersionExportsStorageLegacyKV:
		return true
	}
	return false
}

// Summary of the declarative exports reconciliation that ran on this upload.
// Populated only when the uploaded metadata included an `exports` block. Durable
// Object entries drive reconciliation; `type: worker` entries do not contribute to
// this summary.
type VersionExportsReconciliation struct {
	// Class names for which a new namespace was provisioned.
	Created []string `json:"created" api:"required"`
	// Class names whose namespace was deleted by a `deleted` tombstone.
	Deleted []string `json:"deleted" api:"required"`
	// Non-blocking info entries (stale tombstones, tombstone applied with class still
	// in code). See `exports_reconciliation_info`.
	Info []VersionExportsReconciliationInfo `json:"info" api:"required"`
	// Source class names whose tombstone entry is now stale and safe to delete from
	// `exports` (no remaining referencing scripts).
	RemovableEntries []string `json:"removable_entries" api:"required"`
	// Applied `renamed` tombstones.
	Renamed []VersionExportsReconciliationRenamed `json:"renamed" api:"required"`
	// Phase-1 transfer hints recorded on the target side.
	TransferPending []VersionExportsReconciliationTransferPending `json:"transfer_pending" api:"required"`
	// Committed `transferred` tombstones (phase-2).
	Transferred []VersionExportsReconciliationTransferred `json:"transferred" api:"required"`
	// Class names whose provisioned namespace was mutated in place.
	Updated []string `json:"updated" api:"required"`
	// Non-blocking warnings. See `exports_reconciliation_warning`.
	Warnings []VersionExportsReconciliationWarning `json:"warnings" api:"required"`
	JSON     versionExportsReconciliationJSON      `json:"-"`
}

// versionExportsReconciliationJSON contains the JSON metadata for the struct
// [VersionExportsReconciliation]
type versionExportsReconciliationJSON struct {
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

func (r *VersionExportsReconciliation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionExportsReconciliationJSON) RawJSON() string {
	return r.raw
}

// A non-blocking reconciliation info entry. Emitted for stale tombstones (a no-op
// on this deploy) and for tombstones applied with the source class still in code
// (the supported zero-downtime rollout pattern).
type VersionExportsReconciliationInfo struct {
	// The class name the info entry is about.
	Class string `json:"class" api:"required"`
	// Human-readable explanation.
	Message string `json:"message" api:"required"`
	// Stable, machine-readable tag identifying which reconciliation scenario produced
	// an error, warning, or info entry. Clients may branch on this value instead of
	// parsing `message`.
	Scenario VersionExportsReconciliationInfoScenario `json:"scenario" api:"required"`
	// The provisioned namespace the entry relates to, when applicable.
	NamespaceID string `json:"namespace_id" format:"uuid"`
	// Other Workers in the account that still bind to the affected class. Advisory:
	// while non-empty the tombstone is not yet safe to remove — redeploy these Workers
	// with bindings re-pointed first.
	ReferencingScripts []string                             `json:"referencing_scripts"`
	JSON               versionExportsReconciliationInfoJSON `json:"-"`
}

// versionExportsReconciliationInfoJSON contains the JSON metadata for the struct
// [VersionExportsReconciliationInfo]
type versionExportsReconciliationInfoJSON struct {
	Class              apijson.Field
	Message            apijson.Field
	Scenario           apijson.Field
	NamespaceID        apijson.Field
	ReferencingScripts apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *VersionExportsReconciliationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionExportsReconciliationInfoJSON) RawJSON() string {
	return r.raw
}

// Stable, machine-readable tag identifying which reconciliation scenario produced
// an error, warning, or info entry. Clients may branch on this value instead of
// parsing `message`.
type VersionExportsReconciliationInfoScenario string

const (
	VersionExportsReconciliationInfoScenarioCodeClassNotInExports                     VersionExportsReconciliationInfoScenario = "code_class_not_in_exports"
	VersionExportsReconciliationInfoScenarioProvisionedClassMissingFromConfig         VersionExportsReconciliationInfoScenario = "provisioned_class_missing_from_config"
	VersionExportsReconciliationInfoScenarioConfigExportNotInCode                     VersionExportsReconciliationInfoScenario = "config_export_not_in_code"
	VersionExportsReconciliationInfoScenarioConfigReferencesNonexistentClass          VersionExportsReconciliationInfoScenario = "config_references_nonexistent_class"
	VersionExportsReconciliationInfoScenarioOrphanedProvisionedNamespace              VersionExportsReconciliationInfoScenario = "orphaned_provisioned_namespace"
	VersionExportsReconciliationInfoScenarioStorageTypeMismatch                       VersionExportsReconciliationInfoScenario = "storage_type_mismatch"
	VersionExportsReconciliationInfoScenarioFreeTierRequiresSqlite                    VersionExportsReconciliationInfoScenario = "free_tier_requires_sqlite"
	VersionExportsReconciliationInfoScenarioInvalidExport                             VersionExportsReconciliationInfoScenario = "invalid_export"
	VersionExportsReconciliationInfoScenarioTombstoneDeleteClassStillInCode           VersionExportsReconciliationInfoScenario = "tombstone_delete_class_still_in_code"
	VersionExportsReconciliationInfoScenarioTombstoneDeleteBlockedByExternalBindings  VersionExportsReconciliationInfoScenario = "tombstone_delete_blocked_by_external_bindings"
	VersionExportsReconciliationInfoScenarioTombstoneRenamedToOccupied                VersionExportsReconciliationInfoScenario = "tombstone_renamed_to_occupied"
	VersionExportsReconciliationInfoScenarioTransferredPendingNotFound                VersionExportsReconciliationInfoScenario = "transferred_pending_not_found"
	VersionExportsReconciliationInfoScenarioTransferredTargetMissing                  VersionExportsReconciliationInfoScenario = "transferred_target_missing"
	VersionExportsReconciliationInfoScenarioTransferredTargetMismatch                 VersionExportsReconciliationInfoScenario = "transferred_target_mismatch"
	VersionExportsReconciliationInfoScenarioPhaseOneTransferSourceMissing             VersionExportsReconciliationInfoScenario = "phase_one_transfer_source_missing"
	VersionExportsReconciliationInfoScenarioPhaseOneTransferSourceNamespaceMissing    VersionExportsReconciliationInfoScenario = "phase_one_transfer_source_namespace_missing"
	VersionExportsReconciliationInfoScenarioPhaseOneTransferTargetClassProvisioned    VersionExportsReconciliationInfoScenario = "phase_one_transfer_target_class_provisioned"
	VersionExportsReconciliationInfoScenarioPhaseOneTransferAfterCommitMismatch       VersionExportsReconciliationInfoScenario = "phase_one_transfer_after_commit_mismatch"
	VersionExportsReconciliationInfoScenarioPhaseOneTransferDuplicate                 VersionExportsReconciliationInfoScenario = "phase_one_transfer_duplicate"
	VersionExportsReconciliationInfoScenarioPhaseOneTransferTargetInDispatchNamespace VersionExportsReconciliationInfoScenario = "phase_one_transfer_target_in_dispatch_namespace"
	VersionExportsReconciliationInfoScenarioPhaseOneTransferSourceInDispatchNamespace VersionExportsReconciliationInfoScenario = "phase_one_transfer_source_in_dispatch_namespace"
	VersionExportsReconciliationInfoScenarioTransferredSourceInDispatchNamespace      VersionExportsReconciliationInfoScenario = "transferred_source_in_dispatch_namespace"
	VersionExportsReconciliationInfoScenarioTransferredTargetInDispatchNamespace      VersionExportsReconciliationInfoScenario = "transferred_target_in_dispatch_namespace"
	VersionExportsReconciliationInfoScenarioContainerUndeclaredReference              VersionExportsReconciliationInfoScenario = "container_undeclared_reference"
	VersionExportsReconciliationInfoScenarioContainerClassNotDurableObject            VersionExportsReconciliationInfoScenario = "container_class_not_durable_object"
	VersionExportsReconciliationInfoScenarioContainerWiringInconsistent               VersionExportsReconciliationInfoScenario = "container_wiring_inconsistent"
	VersionExportsReconciliationInfoScenarioContainerMultipleDurableObjects           VersionExportsReconciliationInfoScenario = "container_multiple_durable_objects"
	VersionExportsReconciliationInfoScenarioTransferContainerParityMismatch           VersionExportsReconciliationInfoScenario = "transfer_container_parity_mismatch"
	VersionExportsReconciliationInfoScenarioTransferContainerParityMismatchOnCommit   VersionExportsReconciliationInfoScenario = "transfer_container_parity_mismatch_on_commit"
	VersionExportsReconciliationInfoScenarioTombstoneClassStillInCode                 VersionExportsReconciliationInfoScenario = "tombstone_class_still_in_code"
	VersionExportsReconciliationInfoScenarioStaleTombstone                            VersionExportsReconciliationInfoScenario = "stale_tombstone"
	VersionExportsReconciliationInfoScenarioTransferReceiveAlreadyApplied             VersionExportsReconciliationInfoScenario = "transfer_receive_already_applied"
	VersionExportsReconciliationInfoScenarioTransferReceiveCleanupComplete            VersionExportsReconciliationInfoScenario = "transfer_receive_cleanup_complete"
)

func (r VersionExportsReconciliationInfoScenario) IsKnown() bool {
	switch r {
	case VersionExportsReconciliationInfoScenarioCodeClassNotInExports, VersionExportsReconciliationInfoScenarioProvisionedClassMissingFromConfig, VersionExportsReconciliationInfoScenarioConfigExportNotInCode, VersionExportsReconciliationInfoScenarioConfigReferencesNonexistentClass, VersionExportsReconciliationInfoScenarioOrphanedProvisionedNamespace, VersionExportsReconciliationInfoScenarioStorageTypeMismatch, VersionExportsReconciliationInfoScenarioFreeTierRequiresSqlite, VersionExportsReconciliationInfoScenarioInvalidExport, VersionExportsReconciliationInfoScenarioTombstoneDeleteClassStillInCode, VersionExportsReconciliationInfoScenarioTombstoneDeleteBlockedByExternalBindings, VersionExportsReconciliationInfoScenarioTombstoneRenamedToOccupied, VersionExportsReconciliationInfoScenarioTransferredPendingNotFound, VersionExportsReconciliationInfoScenarioTransferredTargetMissing, VersionExportsReconciliationInfoScenarioTransferredTargetMismatch, VersionExportsReconciliationInfoScenarioPhaseOneTransferSourceMissing, VersionExportsReconciliationInfoScenarioPhaseOneTransferSourceNamespaceMissing, VersionExportsReconciliationInfoScenarioPhaseOneTransferTargetClassProvisioned, VersionExportsReconciliationInfoScenarioPhaseOneTransferAfterCommitMismatch, VersionExportsReconciliationInfoScenarioPhaseOneTransferDuplicate, VersionExportsReconciliationInfoScenarioPhaseOneTransferTargetInDispatchNamespace, VersionExportsReconciliationInfoScenarioPhaseOneTransferSourceInDispatchNamespace, VersionExportsReconciliationInfoScenarioTransferredSourceInDispatchNamespace, VersionExportsReconciliationInfoScenarioTransferredTargetInDispatchNamespace, VersionExportsReconciliationInfoScenarioContainerUndeclaredReference, VersionExportsReconciliationInfoScenarioContainerClassNotDurableObject, VersionExportsReconciliationInfoScenarioContainerWiringInconsistent, VersionExportsReconciliationInfoScenarioContainerMultipleDurableObjects, VersionExportsReconciliationInfoScenarioTransferContainerParityMismatch, VersionExportsReconciliationInfoScenarioTransferContainerParityMismatchOnCommit, VersionExportsReconciliationInfoScenarioTombstoneClassStillInCode, VersionExportsReconciliationInfoScenarioStaleTombstone, VersionExportsReconciliationInfoScenarioTransferReceiveAlreadyApplied, VersionExportsReconciliationInfoScenarioTransferReceiveCleanupComplete:
		return true
	}
	return false
}

// A single applied `renamed` tombstone.
type VersionExportsReconciliationRenamed struct {
	// The original (source) class name.
	From string `json:"from" api:"required"`
	// The new class name (`renamed_to`).
	To   string                                  `json:"to" api:"required"`
	JSON versionExportsReconciliationRenamedJSON `json:"-"`
}

// versionExportsReconciliationRenamedJSON contains the JSON metadata for the
// struct [VersionExportsReconciliationRenamed]
type versionExportsReconciliationRenamedJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionExportsReconciliationRenamed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionExportsReconciliationRenamedJSON) RawJSON() string {
	return r.raw
}

// A single phase-1 transfer hint recorded on the target side (a live
// `expecting-transfer` entry).
type VersionExportsReconciliationTransferPending struct {
	// The target-side class name awaiting transfer.
	Class string `json:"class" api:"required"`
	// The source script the namespace will be transferred from.
	From string                                          `json:"from" api:"required"`
	JSON versionExportsReconciliationTransferPendingJSON `json:"-"`
}

// versionExportsReconciliationTransferPendingJSON contains the JSON metadata for
// the struct [VersionExportsReconciliationTransferPending]
type versionExportsReconciliationTransferPendingJSON struct {
	Class       apijson.Field
	From        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionExportsReconciliationTransferPending) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionExportsReconciliationTransferPendingJSON) RawJSON() string {
	return r.raw
}

// A single committed `transferred` tombstone (phase-2 commit).
type VersionExportsReconciliationTransferred struct {
	// The source class name that was transferred.
	Class string `json:"class" api:"required"`
	// The transfer phase. Currently always `committed`.
	Phase VersionExportsReconciliationTransferredPhase `json:"phase" api:"required"`
	// The destination script that now owns the namespace.
	To   string                                      `json:"to" api:"required"`
	JSON versionExportsReconciliationTransferredJSON `json:"-"`
}

// versionExportsReconciliationTransferredJSON contains the JSON metadata for the
// struct [VersionExportsReconciliationTransferred]
type versionExportsReconciliationTransferredJSON struct {
	Class       apijson.Field
	Phase       apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionExportsReconciliationTransferred) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionExportsReconciliationTransferredJSON) RawJSON() string {
	return r.raw
}

// The transfer phase. Currently always `committed`.
type VersionExportsReconciliationTransferredPhase string

const (
	VersionExportsReconciliationTransferredPhaseCommitted VersionExportsReconciliationTransferredPhase = "committed"
)

func (r VersionExportsReconciliationTransferredPhase) IsKnown() bool {
	switch r {
	case VersionExportsReconciliationTransferredPhaseCommitted:
		return true
	}
	return false
}

// A non-blocking reconciliation warning. Reserved: no scenario populates this
// array today (`code_class_not_in_exports` is surfaced as info and
// `provisioned_class_missing_from_config` is a hard error). Clients should still
// surface any entries that appear.
type VersionExportsReconciliationWarning struct {
	// The class name the warning is about.
	Class string `json:"class" api:"required"`
	// Human-readable explanation of the warning.
	Message string `json:"message" api:"required"`
	// Stable, machine-readable tag identifying which reconciliation scenario produced
	// an error, warning, or info entry. Clients may branch on this value instead of
	// parsing `message`.
	Scenario VersionExportsReconciliationWarningsScenario `json:"scenario" api:"required"`
	// The provisioned namespace the warning relates to, when applicable.
	NamespaceID string                                  `json:"namespace_id" format:"uuid"`
	JSON        versionExportsReconciliationWarningJSON `json:"-"`
}

// versionExportsReconciliationWarningJSON contains the JSON metadata for the
// struct [VersionExportsReconciliationWarning]
type versionExportsReconciliationWarningJSON struct {
	Class       apijson.Field
	Message     apijson.Field
	Scenario    apijson.Field
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionExportsReconciliationWarning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionExportsReconciliationWarningJSON) RawJSON() string {
	return r.raw
}

// Stable, machine-readable tag identifying which reconciliation scenario produced
// an error, warning, or info entry. Clients may branch on this value instead of
// parsing `message`.
type VersionExportsReconciliationWarningsScenario string

const (
	VersionExportsReconciliationWarningsScenarioCodeClassNotInExports                     VersionExportsReconciliationWarningsScenario = "code_class_not_in_exports"
	VersionExportsReconciliationWarningsScenarioProvisionedClassMissingFromConfig         VersionExportsReconciliationWarningsScenario = "provisioned_class_missing_from_config"
	VersionExportsReconciliationWarningsScenarioConfigExportNotInCode                     VersionExportsReconciliationWarningsScenario = "config_export_not_in_code"
	VersionExportsReconciliationWarningsScenarioConfigReferencesNonexistentClass          VersionExportsReconciliationWarningsScenario = "config_references_nonexistent_class"
	VersionExportsReconciliationWarningsScenarioOrphanedProvisionedNamespace              VersionExportsReconciliationWarningsScenario = "orphaned_provisioned_namespace"
	VersionExportsReconciliationWarningsScenarioStorageTypeMismatch                       VersionExportsReconciliationWarningsScenario = "storage_type_mismatch"
	VersionExportsReconciliationWarningsScenarioFreeTierRequiresSqlite                    VersionExportsReconciliationWarningsScenario = "free_tier_requires_sqlite"
	VersionExportsReconciliationWarningsScenarioInvalidExport                             VersionExportsReconciliationWarningsScenario = "invalid_export"
	VersionExportsReconciliationWarningsScenarioTombstoneDeleteClassStillInCode           VersionExportsReconciliationWarningsScenario = "tombstone_delete_class_still_in_code"
	VersionExportsReconciliationWarningsScenarioTombstoneDeleteBlockedByExternalBindings  VersionExportsReconciliationWarningsScenario = "tombstone_delete_blocked_by_external_bindings"
	VersionExportsReconciliationWarningsScenarioTombstoneRenamedToOccupied                VersionExportsReconciliationWarningsScenario = "tombstone_renamed_to_occupied"
	VersionExportsReconciliationWarningsScenarioTransferredPendingNotFound                VersionExportsReconciliationWarningsScenario = "transferred_pending_not_found"
	VersionExportsReconciliationWarningsScenarioTransferredTargetMissing                  VersionExportsReconciliationWarningsScenario = "transferred_target_missing"
	VersionExportsReconciliationWarningsScenarioTransferredTargetMismatch                 VersionExportsReconciliationWarningsScenario = "transferred_target_mismatch"
	VersionExportsReconciliationWarningsScenarioPhaseOneTransferSourceMissing             VersionExportsReconciliationWarningsScenario = "phase_one_transfer_source_missing"
	VersionExportsReconciliationWarningsScenarioPhaseOneTransferSourceNamespaceMissing    VersionExportsReconciliationWarningsScenario = "phase_one_transfer_source_namespace_missing"
	VersionExportsReconciliationWarningsScenarioPhaseOneTransferTargetClassProvisioned    VersionExportsReconciliationWarningsScenario = "phase_one_transfer_target_class_provisioned"
	VersionExportsReconciliationWarningsScenarioPhaseOneTransferAfterCommitMismatch       VersionExportsReconciliationWarningsScenario = "phase_one_transfer_after_commit_mismatch"
	VersionExportsReconciliationWarningsScenarioPhaseOneTransferDuplicate                 VersionExportsReconciliationWarningsScenario = "phase_one_transfer_duplicate"
	VersionExportsReconciliationWarningsScenarioPhaseOneTransferTargetInDispatchNamespace VersionExportsReconciliationWarningsScenario = "phase_one_transfer_target_in_dispatch_namespace"
	VersionExportsReconciliationWarningsScenarioPhaseOneTransferSourceInDispatchNamespace VersionExportsReconciliationWarningsScenario = "phase_one_transfer_source_in_dispatch_namespace"
	VersionExportsReconciliationWarningsScenarioTransferredSourceInDispatchNamespace      VersionExportsReconciliationWarningsScenario = "transferred_source_in_dispatch_namespace"
	VersionExportsReconciliationWarningsScenarioTransferredTargetInDispatchNamespace      VersionExportsReconciliationWarningsScenario = "transferred_target_in_dispatch_namespace"
	VersionExportsReconciliationWarningsScenarioContainerUndeclaredReference              VersionExportsReconciliationWarningsScenario = "container_undeclared_reference"
	VersionExportsReconciliationWarningsScenarioContainerClassNotDurableObject            VersionExportsReconciliationWarningsScenario = "container_class_not_durable_object"
	VersionExportsReconciliationWarningsScenarioContainerWiringInconsistent               VersionExportsReconciliationWarningsScenario = "container_wiring_inconsistent"
	VersionExportsReconciliationWarningsScenarioContainerMultipleDurableObjects           VersionExportsReconciliationWarningsScenario = "container_multiple_durable_objects"
	VersionExportsReconciliationWarningsScenarioTransferContainerParityMismatch           VersionExportsReconciliationWarningsScenario = "transfer_container_parity_mismatch"
	VersionExportsReconciliationWarningsScenarioTransferContainerParityMismatchOnCommit   VersionExportsReconciliationWarningsScenario = "transfer_container_parity_mismatch_on_commit"
	VersionExportsReconciliationWarningsScenarioTombstoneClassStillInCode                 VersionExportsReconciliationWarningsScenario = "tombstone_class_still_in_code"
	VersionExportsReconciliationWarningsScenarioStaleTombstone                            VersionExportsReconciliationWarningsScenario = "stale_tombstone"
	VersionExportsReconciliationWarningsScenarioTransferReceiveAlreadyApplied             VersionExportsReconciliationWarningsScenario = "transfer_receive_already_applied"
	VersionExportsReconciliationWarningsScenarioTransferReceiveCleanupComplete            VersionExportsReconciliationWarningsScenario = "transfer_receive_cleanup_complete"
)

func (r VersionExportsReconciliationWarningsScenario) IsKnown() bool {
	switch r {
	case VersionExportsReconciliationWarningsScenarioCodeClassNotInExports, VersionExportsReconciliationWarningsScenarioProvisionedClassMissingFromConfig, VersionExportsReconciliationWarningsScenarioConfigExportNotInCode, VersionExportsReconciliationWarningsScenarioConfigReferencesNonexistentClass, VersionExportsReconciliationWarningsScenarioOrphanedProvisionedNamespace, VersionExportsReconciliationWarningsScenarioStorageTypeMismatch, VersionExportsReconciliationWarningsScenarioFreeTierRequiresSqlite, VersionExportsReconciliationWarningsScenarioInvalidExport, VersionExportsReconciliationWarningsScenarioTombstoneDeleteClassStillInCode, VersionExportsReconciliationWarningsScenarioTombstoneDeleteBlockedByExternalBindings, VersionExportsReconciliationWarningsScenarioTombstoneRenamedToOccupied, VersionExportsReconciliationWarningsScenarioTransferredPendingNotFound, VersionExportsReconciliationWarningsScenarioTransferredTargetMissing, VersionExportsReconciliationWarningsScenarioTransferredTargetMismatch, VersionExportsReconciliationWarningsScenarioPhaseOneTransferSourceMissing, VersionExportsReconciliationWarningsScenarioPhaseOneTransferSourceNamespaceMissing, VersionExportsReconciliationWarningsScenarioPhaseOneTransferTargetClassProvisioned, VersionExportsReconciliationWarningsScenarioPhaseOneTransferAfterCommitMismatch, VersionExportsReconciliationWarningsScenarioPhaseOneTransferDuplicate, VersionExportsReconciliationWarningsScenarioPhaseOneTransferTargetInDispatchNamespace, VersionExportsReconciliationWarningsScenarioPhaseOneTransferSourceInDispatchNamespace, VersionExportsReconciliationWarningsScenarioTransferredSourceInDispatchNamespace, VersionExportsReconciliationWarningsScenarioTransferredTargetInDispatchNamespace, VersionExportsReconciliationWarningsScenarioContainerUndeclaredReference, VersionExportsReconciliationWarningsScenarioContainerClassNotDurableObject, VersionExportsReconciliationWarningsScenarioContainerWiringInconsistent, VersionExportsReconciliationWarningsScenarioContainerMultipleDurableObjects, VersionExportsReconciliationWarningsScenarioTransferContainerParityMismatch, VersionExportsReconciliationWarningsScenarioTransferContainerParityMismatchOnCommit, VersionExportsReconciliationWarningsScenarioTombstoneClassStillInCode, VersionExportsReconciliationWarningsScenarioStaleTombstone, VersionExportsReconciliationWarningsScenarioTransferReceiveAlreadyApplied, VersionExportsReconciliationWarningsScenarioTransferReceiveCleanupComplete:
		return true
	}
	return false
}

// Resource limits enforced at runtime.
type VersionLimits struct {
	// CPU time limit in milliseconds.
	CPUMs int64 `json:"cpu_ms"`
	// Subrequest limit per request.
	Subrequests int64             `json:"subrequests"`
	JSON        versionLimitsJSON `json:"-"`
}

// versionLimitsJSON contains the JSON metadata for the struct [VersionLimits]
type versionLimitsJSON struct {
	CPUMs       apijson.Field
	Subrequests apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionLimitsJSON) RawJSON() string {
	return r.raw
}

// Migrations for Durable Objects associated with the version. Migrations are
// applied when the version is deployed.
type VersionMigrations struct {
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
	TransferredClasses interface{}           `json:"transferred_classes"`
	JSON               versionMigrationsJSON `json:"-"`
	union              VersionMigrationsUnion
}

// versionMigrationsJSON contains the JSON metadata for the struct
// [VersionMigrations]
type versionMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewSqliteClasses   apijson.Field
	RenamedClasses     apijson.Field
	Steps              apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r versionMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r *VersionMigrations) UnmarshalJSON(data []byte) (err error) {
	*r = VersionMigrations{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [VersionMigrationsUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [SingleStepMigration],
// [VersionMigrationsWorkersMultipleStepMigrations].
func (r VersionMigrations) AsUnion() VersionMigrationsUnion {
	return r.union
}

// Migrations for Durable Objects associated with the version. Migrations are
// applied when the version is deployed.
//
// Union satisfied by [SingleStepMigration] or
// [VersionMigrationsWorkersMultipleStepMigrations].
type VersionMigrationsUnion interface {
	implementsVersionMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionMigrationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SingleStepMigration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionMigrationsWorkersMultipleStepMigrations{}),
		},
	)
}

type VersionMigrationsWorkersMultipleStepMigrations struct {
	JSON versionMigrationsWorkersMultipleStepMigrationsJSON `json:"-"`
}

// versionMigrationsWorkersMultipleStepMigrationsJSON contains the JSON metadata
// for the struct [VersionMigrationsWorkersMultipleStepMigrations]
type versionMigrationsWorkersMultipleStepMigrationsJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionMigrationsWorkersMultipleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionMigrationsWorkersMultipleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r VersionMigrationsWorkersMultipleStepMigrations) implementsVersionMigrations() {}

type VersionModule struct {
	// The base64-encoded module content.
	ContentBase64 string `json:"content_base64" api:"required" format:"byte"`
	// The content type of the module.
	ContentType string `json:"content_type" api:"required"`
	// The name of the module.
	Name string            `json:"name" api:"required"`
	JSON versionModuleJSON `json:"-"`
}

// versionModuleJSON contains the JSON metadata for the struct [VersionModule]
type versionModuleJSON struct {
	ContentBase64 apijson.Field
	ContentType   apijson.Field
	Name          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VersionModule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionModuleJSON) RawJSON() string {
	return r.raw
}

type VersionPackageDependency struct {
	// The exact version that was resolved and installed by the package manager.
	InstalledVersion string `json:"installedVersion" api:"required"`
	// The npm package name.
	Name string `json:"name" api:"required"`
	// The version constraint as written in package.json.
	PackageJsonVersion string                       `json:"packageJsonVersion" api:"required"`
	JSON               versionPackageDependencyJSON `json:"-"`
}

// versionPackageDependencyJSON contains the JSON metadata for the struct
// [VersionPackageDependency]
type versionPackageDependencyJSON struct {
	InstalledVersion   apijson.Field
	Name               apijson.Field
	PackageJsonVersion apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *VersionPackageDependency) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionPackageDependencyJSON) RawJSON() string {
	return r.raw
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
type VersionPlacement struct {
	// TCP host and port for targeted placement.
	Host string `json:"host"`
	// HTTP hostname for targeted placement.
	Hostname string `json:"hostname"`
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode VersionPlacementMode `json:"mode"`
	// Cloud region for targeted placement in format 'provider:region'.
	Region string `json:"region"`
	// This field can have the runtime type of [[]VersionPlacementObjectTarget].
	Target interface{}          `json:"target"`
	JSON   versionPlacementJSON `json:"-"`
	union  VersionPlacementUnion
}

// versionPlacementJSON contains the JSON metadata for the struct
// [VersionPlacement]
type versionPlacementJSON struct {
	Host        apijson.Field
	Hostname    apijson.Field
	Mode        apijson.Field
	Region      apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r versionPlacementJSON) RawJSON() string {
	return r.raw
}

func (r *VersionPlacement) UnmarshalJSON(data []byte) (err error) {
	*r = VersionPlacement{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [VersionPlacementUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [VersionPlacementMode],
// [VersionPlacementRegion], [VersionPlacementHostname], [VersionPlacementHost],
// [VersionPlacementObject], [VersionPlacementObject], [VersionPlacementObject],
// [VersionPlacementObject].
func (r VersionPlacement) AsUnion() VersionPlacementUnion {
	return r.union
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
//
// Union satisfied by [VersionPlacementMode], [VersionPlacementRegion],
// [VersionPlacementHostname], [VersionPlacementHost], [VersionPlacementObject],
// [VersionPlacementObject], [VersionPlacementObject] or [VersionPlacementObject].
type VersionPlacementUnion interface {
	implementsVersionPlacement()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionPlacementUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionPlacementMode{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionPlacementRegion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionPlacementHostname{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionPlacementHost{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionPlacementObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionPlacementObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionPlacementObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionPlacementObject{}),
		},
	)
}

type VersionPlacementMode struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode VersionPlacementModeMode `json:"mode" api:"required"`
	JSON versionPlacementModeJSON `json:"-"`
}

// versionPlacementModeJSON contains the JSON metadata for the struct
// [VersionPlacementMode]
type versionPlacementModeJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionPlacementMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionPlacementModeJSON) RawJSON() string {
	return r.raw
}

func (r VersionPlacementMode) implementsVersionPlacement() {}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type VersionPlacementModeMode string

const (
	VersionPlacementModeModeSmart VersionPlacementModeMode = "smart"
)

func (r VersionPlacementModeMode) IsKnown() bool {
	switch r {
	case VersionPlacementModeModeSmart:
		return true
	}
	return false
}

type VersionPlacementRegion struct {
	// Cloud region for targeted placement in format 'provider:region'.
	Region string                     `json:"region" api:"required"`
	JSON   versionPlacementRegionJSON `json:"-"`
}

// versionPlacementRegionJSON contains the JSON metadata for the struct
// [VersionPlacementRegion]
type versionPlacementRegionJSON struct {
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionPlacementRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionPlacementRegionJSON) RawJSON() string {
	return r.raw
}

func (r VersionPlacementRegion) implementsVersionPlacement() {}

type VersionPlacementHostname struct {
	// HTTP hostname for targeted placement.
	Hostname string                       `json:"hostname" api:"required"`
	JSON     versionPlacementHostnameJSON `json:"-"`
}

// versionPlacementHostnameJSON contains the JSON metadata for the struct
// [VersionPlacementHostname]
type versionPlacementHostnameJSON struct {
	Hostname    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionPlacementHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionPlacementHostnameJSON) RawJSON() string {
	return r.raw
}

func (r VersionPlacementHostname) implementsVersionPlacement() {}

type VersionPlacementHost struct {
	// TCP host and port for targeted placement.
	Host string                   `json:"host" api:"required"`
	JSON versionPlacementHostJSON `json:"-"`
}

// versionPlacementHostJSON contains the JSON metadata for the struct
// [VersionPlacementHost]
type versionPlacementHostJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionPlacementHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionPlacementHostJSON) RawJSON() string {
	return r.raw
}

func (r VersionPlacementHost) implementsVersionPlacement() {}

type VersionPlacementObject struct {
	// Targeted placement mode.
	Mode VersionPlacementObjectMode `json:"mode" api:"required"`
	// Cloud region for targeted placement in format 'provider:region'.
	Region string                     `json:"region" api:"required"`
	JSON   versionPlacementObjectJSON `json:"-"`
}

// versionPlacementObjectJSON contains the JSON metadata for the struct
// [VersionPlacementObject]
type versionPlacementObjectJSON struct {
	Mode        apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionPlacementObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionPlacementObjectJSON) RawJSON() string {
	return r.raw
}

func (r VersionPlacementObject) implementsVersionPlacement() {}

// Targeted placement mode.
type VersionPlacementObjectMode string

const (
	VersionPlacementObjectModeTargeted VersionPlacementObjectMode = "targeted"
)

func (r VersionPlacementObjectMode) IsKnown() bool {
	switch r {
	case VersionPlacementObjectModeTargeted:
		return true
	}
	return false
}

// Usage model for the version.
type VersionUsageModel string

const (
	VersionUsageModelStandard VersionUsageModel = "standard"
	VersionUsageModelBundled  VersionUsageModel = "bundled"
	VersionUsageModelUnbound  VersionUsageModel = "unbound"
)

func (r VersionUsageModel) IsKnown() bool {
	switch r {
	case VersionUsageModelStandard, VersionUsageModelBundled, VersionUsageModelUnbound:
		return true
	}
	return false
}

type VersionParam struct {
	// Metadata about the version.
	Annotations param.Field[VersionAnnotationsParam] `json:"annotations"`
	// Configuration for assets within a Worker.
	//
	// [`_headers`](https://developers.cloudflare.com/workers/static-assets/headers/#custom-headers)
	// and
	// [`_redirects`](https://developers.cloudflare.com/workers/static-assets/redirects/)
	// files should be included as modules named `_headers` and `_redirects` with
	// content type `text/plain`.
	Assets param.Field[VersionAssetsParam] `json:"assets"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings param.Field[[]VersionBindingsUnionParam] `json:"bindings"`
	// Global CacheW configuration for the Worker. When caching is on, the platform
	// provisions a `cloudflare.app` zone for the Worker. A `type: worker` entry in the
	// `exports` map can override this value for a single entrypoint.
	CacheOptions param.Field[VersionCacheOptionsParam] `json:"cache_options"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// List of containers attached to a Worker. Containers can only be attached to
	// Durable Object classes of this Worker script.
	Containers param.Field[[]VersionContainerParam] `json:"containers"`
	// Declarative exports for the version, including Durable Object classes (with
	// their `storage` backend) and named Worker entrypoints. On reads, tombstoned
	// lifecycle entries are omitted, so only live exports (`created` and
	// `expecting-transfer`) are returned. `exports` and `migrations` are mutually
	// exclusive on upload.
	Exports param.Field[map[string]VersionExportsUnionParam] `json:"exports"`
	// Resource limits enforced at runtime.
	Limits param.Field[VersionLimitsParam] `json:"limits"`
	// The name of the main module in the `modules` array (e.g. the name of the module
	// that exports a `fetch` handler).
	MainModule param.Field[string] `json:"main_module"`
	// Migrations for Durable Objects associated with the version. Migrations are
	// applied when the version is deployed.
	Migrations param.Field[VersionMigrationsUnionParam] `json:"migrations"`
	// Code, sourcemaps, and other content used at runtime.
	//
	// This includes
	// [`_headers`](https://developers.cloudflare.com/workers/static-assets/headers/#custom-headers)
	// and
	// [`_redirects`](https://developers.cloudflare.com/workers/static-assets/redirects/)
	// files used to configure
	// [Static Assets](https://developers.cloudflare.com/workers/static-assets/).
	// `_headers` and `_redirects` files should be included as modules named `_headers`
	// and `_redirects` with content type `text/plain`.
	Modules param.Field[[]VersionModuleParam] `json:"modules"`
	// The list of npm packages that were installed and used when this Worker version
	// was built.
	PackageDependencies param.Field[[]VersionPackageDependencyParam] `json:"package_dependencies"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
	Placement param.Field[VersionPlacementUnionParam] `json:"placement"`
	// Usage model for the version.
	//
	// Deprecated: deprecated
	UsageModel param.Field[VersionUsageModel] `json:"usage_model"`
}

func (r VersionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Metadata about the version.
type VersionAnnotationsParam struct {
	// Human-readable message about the version. Truncated to 1000 bytes if longer.
	WorkersMessage param.Field[string] `json:"workers/message"`
	// User-provided identifier for the version. Maximum 100 bytes.
	WorkersTag param.Field[string] `json:"workers/tag"`
}

func (r VersionAnnotationsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker.
//
// [`_headers`](https://developers.cloudflare.com/workers/static-assets/headers/#custom-headers)
// and
// [`_redirects`](https://developers.cloudflare.com/workers/static-assets/redirects/)
// files should be included as modules named `_headers` and `_redirects` with
// content type `text/plain`.
type VersionAssetsParam struct {
	// Configuration for assets within a Worker.
	Config param.Field[VersionAssetsConfigParam] `json:"config"`
	// Token provided upon successful upload of all files from a registered manifest.
	JWT param.Field[string] `json:"jwt"`
}

func (r VersionAssetsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker.
type VersionAssetsConfigParam struct {
	// Determines the redirects and rewrites of requests for HTML content.
	HTMLHandling param.Field[VersionAssetsConfigHTMLHandling] `json:"html_handling"`
	// Determines the response when a request does not match a static asset, and there
	// is no Worker script.
	NotFoundHandling param.Field[VersionAssetsConfigNotFoundHandling] `json:"not_found_handling"`
	// Contains a list path rules to control routing to either the Worker or assets.
	// Glob (\*) and negative (!) rules are supported. Rules must start with either '/'
	// or '!/'. At least one non-negative rule must be provided, and negative rules
	// have higher precedence than non-negative rules.
	RunWorkerFirst param.Field[VersionAssetsConfigRunWorkerFirstUnionParam] `json:"run_worker_first"`
}

func (r VersionAssetsConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Contains a list path rules to control routing to either the Worker or assets.
// Glob (\*) and negative (!) rules are supported. Rules must start with either '/'
// or '!/'. At least one non-negative rule must be provided, and negative rules
// have higher precedence than non-negative rules.
//
// Satisfied by [workers.VersionAssetsConfigRunWorkerFirstArrayParam],
// [shared.UnionBool].
type VersionAssetsConfigRunWorkerFirstUnionParam interface {
	ImplementsVersionAssetsConfigRunWorkerFirstUnionParam()
}

type VersionAssetsConfigRunWorkerFirstArrayParam []string

func (r VersionAssetsConfigRunWorkerFirstArrayParam) ImplementsVersionAssetsConfigRunWorkerFirstUnionParam() {
}

// A binding to allow the Worker to communicate with resources.
type VersionBindingParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsType] `json:"type" api:"required"`
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
	Format param.Field[VersionBindingsFormat] `json:"format"`
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name"`
	// The user-chosen instance name. Must exist at deploy time. The worker can search,
	// chat, update, and manage items/jobs on this instance.
	InstanceName param.Field[string]      `json:"instance_name"`
	Json         param.Field[interface{}] `json:"json"`
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction param.Field[VersionBindingsJurisdiction] `json:"jurisdiction"`
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

func (r VersionBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingParam) implementsVersionBindingsUnionParam() {}

// A binding to allow the Worker to communicate with resources.
//
// Satisfied by [workers.VersionBindingsWorkersBindingKindAIParam],
// [workers.VersionBindingsWorkersBindingKindAISearchParam],
// [workers.VersionBindingsWorkersBindingKindAISearchNamespaceParam],
// [workers.VersionBindingsWorkersBindingKindAnalyticsEngineParam],
// [workers.VersionBindingsWorkersBindingKindAssetsParam],
// [workers.VersionBindingsWorkersBindingKindBrowserParam],
// [workers.VersionBindingsWorkersBindingKindD1Param],
// [workers.VersionBindingsWorkersBindingKindDataBlobParam],
// [workers.VersionBindingsWorkersBindingKindDispatchNamespaceParam],
// [workers.VersionBindingsWorkersBindingKindDurableObjectNamespaceParam],
// [workers.VersionBindingsWorkersBindingKindHyperdriveParam],
// [workers.VersionBindingsWorkersBindingKindInheritParam],
// [workers.VersionBindingsWorkersBindingKindImagesParam],
// [workers.VersionBindingsWorkersBindingKindJsonParam],
// [workers.VersionBindingsWorkersBindingKindKVNamespaceParam],
// [workers.VersionBindingsWorkersBindingKindMediaParam],
// [workers.VersionBindingsWorkersBindingKindMTLSCertificateParam],
// [workers.VersionBindingsWorkersBindingKindPlainTextParam],
// [workers.VersionBindingsWorkersBindingKindPipelinesParam],
// [workers.VersionBindingsWorkersBindingKindQueueParam],
// [workers.VersionBindingsWorkersBindingKindRatelimitParam],
// [workers.VersionBindingsWorkersBindingKindR2BucketParam],
// [workers.VersionBindingsWorkersBindingKindSecretTextParam],
// [workers.VersionBindingsWorkersBindingKindSendEmailParam],
// [workers.VersionBindingsWorkersBindingKindServiceParam],
// [workers.VersionBindingsWorkersBindingKindTextBlobParam],
// [workers.VersionBindingsWorkersBindingKindVectorizeParam],
// [workers.VersionBindingsWorkersBindingKindVersionMetadataParam],
// [workers.VersionBindingsWorkersBindingKindSecretsStoreSecretParam],
// [workers.VersionBindingsWorkersBindingKindFlagshipParam],
// [workers.VersionBindingsWorkersBindingKindSecretKeyParam],
// [workers.VersionBindingsWorkersBindingKindWorkflowParam],
// [workers.VersionBindingsWorkersBindingKindWasmModuleParam],
// [workers.VersionBindingsWorkersBindingKindVPCServiceParam],
// [workers.VersionBindingsWorkersBindingKindVPCNetworkParam],
// [VersionBindingParam].
type VersionBindingsUnionParam interface {
	implementsVersionBindingsUnionParam()
}

type VersionBindingsWorkersBindingKindAIParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindAIType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindAIParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindAIParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindAISearchParam struct {
	// The user-chosen instance name. Must exist at deploy time. The worker can search,
	// chat, update, and manage items/jobs on this instance.
	InstanceName param.Field[string] `json:"instance_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindAISearchType] `json:"type" api:"required"`
	// The namespace the instance belongs to. Defaults to "default" if omitted.
	// Customers who don't use namespaces can simply omit this field.
	Namespace param.Field[string] `json:"namespace"`
}

func (r VersionBindingsWorkersBindingKindAISearchParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindAISearchParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindAISearchNamespaceParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The user-chosen namespace name. Must exist before deploy -- Wrangler handles
	// auto-creation on deploy failure (R2 bucket pattern). The "default" namespace is
	// auto-created by config-api for new accounts. Grants full access (CRUD + search +
	// chat) to all instances within the namespace.
	Namespace param.Field[string] `json:"namespace" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindAISearchNamespaceType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindAISearchNamespaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindAISearchNamespaceParam) implementsVersionBindingsUnionParam() {
}

type VersionBindingsWorkersBindingKindAnalyticsEngineParam struct {
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindAnalyticsEngineType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindAnalyticsEngineParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindAnalyticsEngineParam) implementsVersionBindingsUnionParam() {
}

type VersionBindingsWorkersBindingKindAssetsParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindAssetsType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindAssetsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindAssetsParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindBrowserParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindBrowserType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindBrowserParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindBrowserParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindD1Param struct {
	// Identifier of the D1 database to bind to.
	DatabaseID param.Field[string] `json:"database_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindD1Type] `json:"type" api:"required"`
	// Identifier of the D1 database to bind to.
	//
	// Deprecated: This property has been renamed to `database_id`.
	ID param.Field[string] `json:"id"`
}

func (r VersionBindingsWorkersBindingKindD1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindD1Param) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindDataBlobParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The name of the file containing the data content. Only accepted for
	// `service worker syntax` Workers.
	Part param.Field[string] `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type param.Field[VersionBindingsWorkersBindingKindDataBlobType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindDataBlobParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindDataBlobParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindDispatchNamespaceParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The name of the dispatch namespace.
	Namespace param.Field[string] `json:"namespace" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindDispatchNamespaceType] `json:"type" api:"required"`
	// Outbound worker.
	Outbound param.Field[VersionBindingsWorkersBindingKindDispatchNamespaceOutboundParam] `json:"outbound"`
}

func (r VersionBindingsWorkersBindingKindDispatchNamespaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindDispatchNamespaceParam) implementsVersionBindingsUnionParam() {
}

type VersionBindingsWorkersBindingKindDurableObjectNamespaceParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindDurableObjectNamespaceType] `json:"type" api:"required"`
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

func (r VersionBindingsWorkersBindingKindDurableObjectNamespaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindDurableObjectNamespaceParam) implementsVersionBindingsUnionParam() {
}

type VersionBindingsWorkersBindingKindHyperdriveParam struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID param.Field[string] `json:"id" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindHyperdriveType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindHyperdriveParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindHyperdriveParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindInheritParam struct {
	// The name of the inherited binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindInheritType] `json:"type" api:"required"`
	// The old name of the inherited binding. If set, the binding will be renamed from
	// `old_name` to `name` in the new version. If not set, the binding will keep the
	// same name between versions.
	OldName param.Field[string] `json:"old_name"`
	// Identifier for the version to inherit the binding from, which can be the version
	// ID or the literal "latest" to inherit from the latest version. Defaults to
	// inheriting the binding from the latest version.
	VersionID param.Field[string] `json:"version_id"`
}

func (r VersionBindingsWorkersBindingKindInheritParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindInheritParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindImagesParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindImagesType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindImagesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindImagesParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindJsonParam struct {
	// JSON data to use.
	Json param.Field[interface{}] `json:"json" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindJsonType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindJsonParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindJsonParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindKVNamespaceParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindKVNamespaceType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindKVNamespaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindKVNamespaceParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindMediaParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindMediaType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindMediaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindMediaParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindMTLSCertificateParam struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindMTLSCertificateType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindMTLSCertificateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindMTLSCertificateParam) implementsVersionBindingsUnionParam() {
}

type VersionBindingsWorkersBindingKindPlainTextParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The text value to use.
	Text param.Field[string] `json:"text" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindPlainTextType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindPlainTextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindPlainTextParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindPipelinesParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Name of the Pipeline to bind to.
	Pipeline param.Field[string] `json:"pipeline" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindPipelinesType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindPipelinesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindPipelinesParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindQueueParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindQueueType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindQueueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindQueueParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindRatelimitParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Identifier of the rate limit namespace to bind to.
	NamespaceID param.Field[string] `json:"namespace_id" api:"required"`
	// The rate limit configuration.
	Simple param.Field[VersionBindingsWorkersBindingKindRatelimitSimpleParam] `json:"simple" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindRatelimitType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindRatelimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindRatelimitParam) implementsVersionBindingsUnionParam() {}

// The rate limit configuration.
type VersionBindingsWorkersBindingKindRatelimitSimpleParam struct {
	// The limit (requests per period).
	Limit param.Field[float64] `json:"limit" api:"required"`
	// The period in seconds.
	Period param.Field[int64] `json:"period" api:"required"`
	// Duration in seconds to apply the mitigation action after the rate limit is
	// exceeded. Valid values are 0 (disabled), 10, or multiples of 60 up to 86400.
	// Must be greater than or equal to the period when non-zero.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
}

func (r VersionBindingsWorkersBindingKindRatelimitSimpleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VersionBindingsWorkersBindingKindR2BucketParam struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindR2BucketType] `json:"type" api:"required"`
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction param.Field[VersionBindingsWorkersBindingKindR2BucketJurisdiction] `json:"jurisdiction"`
}

func (r VersionBindingsWorkersBindingKindR2BucketParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindR2BucketParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindSecretTextParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The secret value to use.
	Text param.Field[string] `json:"text" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindSecretTextType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindSecretTextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindSecretTextParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindSendEmailParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindSendEmailType] `json:"type" api:"required"`
	// List of allowed destination addresses.
	AllowedDestinationAddresses param.Field[[]string] `json:"allowed_destination_addresses" format:"email"`
	// List of allowed sender addresses.
	AllowedSenderAddresses param.Field[[]string] `json:"allowed_sender_addresses" format:"email"`
	// Destination address for the email.
	DestinationAddress param.Field[string] `json:"destination_address" format:"email"`
}

func (r VersionBindingsWorkersBindingKindSendEmailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindSendEmailParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindServiceParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindServiceType] `json:"type" api:"required"`
	// Entrypoint to invoke on the target Worker.
	Entrypoint param.Field[string] `json:"entrypoint"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
}

func (r VersionBindingsWorkersBindingKindServiceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindServiceParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindTextBlobParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The name of the file containing the text content. Only accepted for
	// `service worker syntax` Workers.
	Part param.Field[string] `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type param.Field[VersionBindingsWorkersBindingKindTextBlobType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindTextBlobParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindTextBlobParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindVectorizeParam struct {
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindVectorizeType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindVectorizeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindVectorizeParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindVersionMetadataParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindVersionMetadataType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindVersionMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindVersionMetadataParam) implementsVersionBindingsUnionParam() {
}

type VersionBindingsWorkersBindingKindSecretsStoreSecretParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Name of the secret in the store.
	SecretName param.Field[string] `json:"secret_name" api:"required"`
	// ID of the store containing the secret.
	StoreID param.Field[string] `json:"store_id" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindSecretsStoreSecretType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindSecretsStoreSecretParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindSecretsStoreSecretParam) implementsVersionBindingsUnionParam() {
}

type VersionBindingsWorkersBindingKindFlagshipParam struct {
	// ID of the Flagship app to bind to for feature flag evaluation.
	AppID param.Field[string] `json:"app_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindFlagshipType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindFlagshipParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindFlagshipParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindSecretKeyParam struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm param.Field[interface{}] `json:"algorithm" api:"required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[VersionBindingsWorkersBindingKindSecretKeyFormat] `json:"format" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindSecretKeyType] `json:"type" api:"required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages param.Field[[]VersionBindingsWorkersBindingKindSecretKeyUsage] `json:"usages" api:"required"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string] `json:"key_base64"`
	// Key data in
	// [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key)
	// format. Required if `format` is "jwk".
	KeyJwk param.Field[interface{}] `json:"key_jwk"`
}

func (r VersionBindingsWorkersBindingKindSecretKeyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindSecretKeyParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindWorkflowParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindWorkflowType] `json:"type" api:"required"`
	// Name of the Workflow to bind to.
	WorkflowName param.Field[string] `json:"workflow_name" api:"required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName param.Field[string] `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r VersionBindingsWorkersBindingKindWorkflowParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindWorkflowParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindWasmModuleParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The name of the file containing the WebAssembly module content. Only accepted
	// for `service worker syntax` Workers.
	Part param.Field[string] `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type param.Field[VersionBindingsWorkersBindingKindWasmModuleType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindWasmModuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindWasmModuleParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindVPCServiceParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Identifier of the VPC service to bind to.
	ServiceID param.Field[string] `json:"service_id" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindVPCServiceType] `json:"type" api:"required"`
}

func (r VersionBindingsWorkersBindingKindVPCServiceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindVPCServiceParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindVPCNetworkParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindVPCNetworkType] `json:"type" api:"required"`
	// Identifier of the network to bind to. Only "cf1:network" is currently supported.
	// Mutually exclusive with tunnel_id.
	NetworkID param.Field[string] `json:"network_id"`
	// UUID of the Cloudflare Tunnel to bind to. Mutually exclusive with network_id.
	TunnelID param.Field[string] `json:"tunnel_id"`
}

func (r VersionBindingsWorkersBindingKindVPCNetworkParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindVPCNetworkParam) implementsVersionBindingsUnionParam() {}

// Global CacheW configuration for the Worker. When caching is on, the platform
// provisions a `cloudflare.app` zone for the Worker. A `type: worker` entry in the
// `exports` map can override this value for a single entrypoint.
type VersionCacheOptionsParam struct {
	// Whether caching is enabled for this Worker.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Whether cached responses are shared across Worker version uploads. This is
	// independent of `enabled`. It can stay true while caching is off, so the
	// preference survives turning caching off and back on.
	CrossVersionCache param.Field[bool] `json:"cross_version_cache"`
}

func (r VersionCacheOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Container configuration for a Worker.
type VersionContainerParam struct {
	// Select which Durable Object class should get this container attached.
	ClassName param.Field[string] `json:"class_name" api:"required"`
}

func (r VersionContainerParam) MarshalJSON() (data []byte, err error) {
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
type VersionExportParam struct {
	// Marks this entry as a Worker entrypoint export.
	Type  param.Field[VersionExportsType] `json:"type" api:"required"`
	Cache param.Field[interface{}]        `json:"cache"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container param.Field[string] `json:"container"`
	// The destination class name. Must differ from the source class (the map key) and
	// must be declared as a live (`created`) entry in the same `exports` map.
	// Write-only: never present in GET responses.
	RenamedTo param.Field[string] `json:"renamed_to"`
	// Live export. May be omitted; defaults to `created`.
	State param.Field[VersionExportsState] `json:"state"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage param.Field[VersionExportsStorage] `json:"storage"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom param.Field[string] `json:"transfer_from"`
	// The destination script name. Must be in the same account and the same
	// dispatch-namespace context (or both non-dispatch). Cross-dispatch-namespace
	// transfers are rejected. Write-only: never present in GET responses.
	TransferredTo param.Field[string] `json:"transferred_to"`
}

func (r VersionExportParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionExportParam) implementsVersionExportsUnionParam() {}

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
// Satisfied by [workers.VersionExportsWorkersWorkerExportParam],
// [workers.VersionExportsWorkersDurableObjectExportParam],
// [workers.VersionExportsWorkersDurableObjectDeletedExportParam],
// [workers.VersionExportsWorkersDurableObjectRenamedExportParam],
// [workers.VersionExportsWorkersDurableObjectTransferredExportParam],
// [workers.VersionExportsWorkersDurableObjectExpectingTransferExportParam],
// [VersionExportParam].
type VersionExportsUnionParam interface {
	implementsVersionExportsUnionParam()
}

// A named Worker entrypoint export (`type: worker`). Worker entrypoints are always
// live (`state: created`) and carry no storage or lifecycle fields. The optional
// `cache` block overrides the Worker's global `cache_options.enabled` for this
// entrypoint.
type VersionExportsWorkersWorkerExportParam struct {
	// Marks this entry as a Worker entrypoint export.
	Type param.Field[VersionExportsWorkersWorkerExportType] `json:"type" api:"required"`
	// Cache override for this entrypoint. Overrides the Worker's global
	// `cache_options.enabled` for this entrypoint only.
	Cache param.Field[VersionExportsWorkersWorkerExportCacheParam] `json:"cache"`
	// Live export. May be omitted; defaults to `created`.
	State param.Field[VersionExportsWorkersWorkerExportState] `json:"state"`
}

func (r VersionExportsWorkersWorkerExportParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionExportsWorkersWorkerExportParam) implementsVersionExportsUnionParam() {}

// Cache override for this entrypoint. Overrides the Worker's global
// `cache_options.enabled` for this entrypoint only.
type VersionExportsWorkersWorkerExportCacheParam struct {
	// Whether caching is enabled for this entrypoint.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
}

func (r VersionExportsWorkersWorkerExportCacheParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A live Durable Object export (`state: created`, the default). The platform
// auto-provisions the namespace on first deploy, matches it on subsequent deploys,
// and never mutates or deletes it as a side effect of a code-only change.
// `storage` is required; `renamed_to`, `transferred_to` and `transfer_from` are
// not allowed on a live entry.
type VersionExportsWorkersDurableObjectExportParam struct {
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage param.Field[VersionExportsWorkersDurableObjectExportStorage] `json:"storage" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[VersionExportsWorkersDurableObjectExportType] `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container param.Field[string] `json:"container"`
	// Live export. May be omitted; defaults to `created`.
	State param.Field[VersionExportsWorkersDurableObjectExportState] `json:"state"`
}

func (r VersionExportsWorkersDurableObjectExportParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionExportsWorkersDurableObjectExportParam) implementsVersionExportsUnionParam() {}

// A `deleted` tombstone: retires the provisioned namespace for this class and all
// of its data. The class must be absent from the uploaded code and no other Worker
// in the account may bind to the namespace, otherwise the deploy is rejected. No
// other fields are allowed. Deletion is irreversible.
type VersionExportsWorkersDurableObjectDeletedExportParam struct {
	// Tombstone that deletes the namespace.
	State param.Field[VersionExportsWorkersDurableObjectDeletedExportState] `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[VersionExportsWorkersDurableObjectDeletedExportType] `json:"type" api:"required"`
}

func (r VersionExportsWorkersDurableObjectDeletedExportParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionExportsWorkersDurableObjectDeletedExportParam) implementsVersionExportsUnionParam() {}

// A `renamed` tombstone: rewrites the provisioned namespace's class name from this
// map key to `renamed_to`. The source class may stay in code during the rollout
// window (an info notice is emitted). `storage`, `transferred_to` and
// `transfer_from` are not allowed.
type VersionExportsWorkersDurableObjectRenamedExportParam struct {
	// The destination class name. Must differ from the source class (the map key) and
	// must be declared as a live (`created`) entry in the same `exports` map.
	// Write-only: never present in GET responses.
	RenamedTo param.Field[string] `json:"renamed_to" api:"required"`
	// Tombstone that renames the namespace's class.
	State param.Field[VersionExportsWorkersDurableObjectRenamedExportState] `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[VersionExportsWorkersDurableObjectRenamedExportType] `json:"type" api:"required"`
}

func (r VersionExportsWorkersDurableObjectRenamedExportParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionExportsWorkersDurableObjectRenamedExportParam) implementsVersionExportsUnionParam() {}

// A `transferred` tombstone (source side of a two-phase transfer): hands ownership
// of the provisioned namespace to another script in the same account, named by
// `transferred_to`. The target must have already deployed a matching
// `expecting-transfer` entry. The source class may stay in code during the rollout
// window (an info notice is emitted). `storage`, `renamed_to` and `transfer_from`
// are not allowed.
type VersionExportsWorkersDurableObjectTransferredExportParam struct {
	// Tombstone that transfers the namespace to another script.
	State param.Field[VersionExportsWorkersDurableObjectTransferredExportState] `json:"state" api:"required"`
	// The destination script name. Must be in the same account and the same
	// dispatch-namespace context (or both non-dispatch). Cross-dispatch-namespace
	// transfers are rejected. Write-only: never present in GET responses.
	TransferredTo param.Field[string] `json:"transferred_to" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[VersionExportsWorkersDurableObjectTransferredExportType] `json:"type" api:"required"`
}

func (r VersionExportsWorkersDurableObjectTransferredExportParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionExportsWorkersDurableObjectTransferredExportParam) implementsVersionExportsUnionParam() {
}

// The target side of a two-phase transfer (`state: expecting-transfer`). Declares
// that this script expects to receive a namespace for this class from the
// `transfer_from` script. This is a live entry, not a tombstone: bindings resolve
// through the source's namespace until the source commits with a `transferred`
// tombstone. `storage` and `transfer_from` are required; `renamed_to` and
// `transferred_to` are not allowed.
type VersionExportsWorkersDurableObjectExpectingTransferExportParam struct {
	// Target side of a two-phase transfer.
	State param.Field[VersionExportsWorkersDurableObjectExpectingTransferExportState] `json:"state" api:"required"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage param.Field[VersionExportsWorkersDurableObjectExpectingTransferExportStorage] `json:"storage" api:"required"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom param.Field[string] `json:"transfer_from" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[VersionExportsWorkersDurableObjectExpectingTransferExportType] `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object once the transfer settles. Valid only on live entries.
	Container param.Field[string] `json:"container"`
}

func (r VersionExportsWorkersDurableObjectExpectingTransferExportParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionExportsWorkersDurableObjectExpectingTransferExportParam) implementsVersionExportsUnionParam() {
}

// Summary of the declarative exports reconciliation that ran on this upload.
// Populated only when the uploaded metadata included an `exports` block. Durable
// Object entries drive reconciliation; `type: worker` entries do not contribute to
// this summary.
type VersionExportsReconciliationParam struct {
	// Class names for which a new namespace was provisioned.
	Created param.Field[[]string] `json:"created" api:"required"`
	// Class names whose namespace was deleted by a `deleted` tombstone.
	Deleted param.Field[[]string] `json:"deleted" api:"required"`
	// Non-blocking info entries (stale tombstones, tombstone applied with class still
	// in code). See `exports_reconciliation_info`.
	Info param.Field[[]VersionExportsReconciliationInfoParam] `json:"info" api:"required"`
	// Source class names whose tombstone entry is now stale and safe to delete from
	// `exports` (no remaining referencing scripts).
	RemovableEntries param.Field[[]string] `json:"removable_entries" api:"required"`
	// Applied `renamed` tombstones.
	Renamed param.Field[[]VersionExportsReconciliationRenamedParam] `json:"renamed" api:"required"`
	// Phase-1 transfer hints recorded on the target side.
	TransferPending param.Field[[]VersionExportsReconciliationTransferPendingParam] `json:"transfer_pending" api:"required"`
	// Committed `transferred` tombstones (phase-2).
	Transferred param.Field[[]VersionExportsReconciliationTransferredParam] `json:"transferred" api:"required"`
	// Class names whose provisioned namespace was mutated in place.
	Updated param.Field[[]string] `json:"updated" api:"required"`
	// Non-blocking warnings. See `exports_reconciliation_warning`.
	Warnings param.Field[[]VersionExportsReconciliationWarningParam] `json:"warnings" api:"required"`
}

func (r VersionExportsReconciliationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A non-blocking reconciliation info entry. Emitted for stale tombstones (a no-op
// on this deploy) and for tombstones applied with the source class still in code
// (the supported zero-downtime rollout pattern).
type VersionExportsReconciliationInfoParam struct {
	// The class name the info entry is about.
	Class param.Field[string] `json:"class" api:"required"`
	// Human-readable explanation.
	Message param.Field[string] `json:"message" api:"required"`
	// Stable, machine-readable tag identifying which reconciliation scenario produced
	// an error, warning, or info entry. Clients may branch on this value instead of
	// parsing `message`.
	Scenario param.Field[VersionExportsReconciliationInfoScenario] `json:"scenario" api:"required"`
	// The provisioned namespace the entry relates to, when applicable.
	NamespaceID param.Field[string] `json:"namespace_id" format:"uuid"`
	// Other Workers in the account that still bind to the affected class. Advisory:
	// while non-empty the tombstone is not yet safe to remove — redeploy these Workers
	// with bindings re-pointed first.
	ReferencingScripts param.Field[[]string] `json:"referencing_scripts"`
}

func (r VersionExportsReconciliationInfoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A single applied `renamed` tombstone.
type VersionExportsReconciliationRenamedParam struct {
	// The original (source) class name.
	From param.Field[string] `json:"from" api:"required"`
	// The new class name (`renamed_to`).
	To param.Field[string] `json:"to" api:"required"`
}

func (r VersionExportsReconciliationRenamedParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A single phase-1 transfer hint recorded on the target side (a live
// `expecting-transfer` entry).
type VersionExportsReconciliationTransferPendingParam struct {
	// The target-side class name awaiting transfer.
	Class param.Field[string] `json:"class" api:"required"`
	// The source script the namespace will be transferred from.
	From param.Field[string] `json:"from" api:"required"`
}

func (r VersionExportsReconciliationTransferPendingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A single committed `transferred` tombstone (phase-2 commit).
type VersionExportsReconciliationTransferredParam struct {
	// The source class name that was transferred.
	Class param.Field[string] `json:"class" api:"required"`
	// The transfer phase. Currently always `committed`.
	Phase param.Field[VersionExportsReconciliationTransferredPhase] `json:"phase" api:"required"`
	// The destination script that now owns the namespace.
	To param.Field[string] `json:"to" api:"required"`
}

func (r VersionExportsReconciliationTransferredParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A non-blocking reconciliation warning. Reserved: no scenario populates this
// array today (`code_class_not_in_exports` is surfaced as info and
// `provisioned_class_missing_from_config` is a hard error). Clients should still
// surface any entries that appear.
type VersionExportsReconciliationWarningParam struct {
	// The class name the warning is about.
	Class param.Field[string] `json:"class" api:"required"`
	// Human-readable explanation of the warning.
	Message param.Field[string] `json:"message" api:"required"`
	// Stable, machine-readable tag identifying which reconciliation scenario produced
	// an error, warning, or info entry. Clients may branch on this value instead of
	// parsing `message`.
	Scenario param.Field[VersionExportsReconciliationWarningsScenario] `json:"scenario" api:"required"`
	// The provisioned namespace the warning relates to, when applicable.
	NamespaceID param.Field[string] `json:"namespace_id" format:"uuid"`
}

func (r VersionExportsReconciliationWarningParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Resource limits enforced at runtime.
type VersionLimitsParam struct {
	// CPU time limit in milliseconds.
	CPUMs param.Field[int64] `json:"cpu_ms"`
	// Subrequest limit per request.
	Subrequests param.Field[int64] `json:"subrequests"`
}

func (r VersionLimitsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Migrations for Durable Objects associated with the version. Migrations are
// applied when the version is deployed.
type VersionMigrationsParam struct {
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

func (r VersionMigrationsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionMigrationsParam) implementsVersionMigrationsUnionParam() {}

// Migrations for Durable Objects associated with the version. Migrations are
// applied when the version is deployed.
//
// Satisfied by [workers.SingleStepMigrationParam],
// [workers.VersionMigrationsWorkersMultipleStepMigrationsParam],
// [VersionMigrationsParam].
type VersionMigrationsUnionParam interface {
	implementsVersionMigrationsUnionParam()
}

type VersionMigrationsWorkersMultipleStepMigrationsParam struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]MigrationStepParam] `json:"steps"`
}

func (r VersionMigrationsWorkersMultipleStepMigrationsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionMigrationsWorkersMultipleStepMigrationsParam) implementsVersionMigrationsUnionParam() {
}

type VersionModuleParam struct {
	// The base64-encoded module content.
	ContentBase64 param.Field[string] `json:"content_base64" api:"required" format:"byte"`
	// The content type of the module.
	ContentType param.Field[string] `json:"content_type" api:"required"`
	// The name of the module.
	Name param.Field[string] `json:"name" api:"required"`
}

func (r VersionModuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VersionPackageDependencyParam struct {
	// The exact version that was resolved and installed by the package manager.
	InstalledVersion param.Field[string] `json:"installedVersion" api:"required"`
	// The npm package name.
	Name param.Field[string] `json:"name" api:"required"`
	// The version constraint as written in package.json.
	PackageJsonVersion param.Field[string] `json:"packageJsonVersion" api:"required"`
}

func (r VersionPackageDependencyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
type VersionPlacementParam struct {
	// TCP host and port for targeted placement.
	Host param.Field[string] `json:"host"`
	// HTTP hostname for targeted placement.
	Hostname param.Field[string] `json:"hostname"`
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode param.Field[VersionPlacementMode] `json:"mode"`
	// Cloud region for targeted placement in format 'provider:region'.
	Region param.Field[string]      `json:"region"`
	Target param.Field[interface{}] `json:"target"`
}

func (r VersionPlacementParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionPlacementParam) implementsVersionPlacementUnionParam() {}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
//
// Satisfied by [workers.VersionPlacementModeParam],
// [workers.VersionPlacementRegionParam], [workers.VersionPlacementHostnameParam],
// [workers.VersionPlacementHostParam], [workers.VersionPlacementObjectParam],
// [workers.VersionPlacementObjectParam], [workers.VersionPlacementObjectParam],
// [workers.VersionPlacementObjectParam], [VersionPlacementParam].
type VersionPlacementUnionParam interface {
	implementsVersionPlacementUnionParam()
}

type VersionPlacementModeParam struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode param.Field[VersionPlacementModeMode] `json:"mode" api:"required"`
}

func (r VersionPlacementModeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionPlacementModeParam) implementsVersionPlacementUnionParam() {}

type VersionPlacementRegionParam struct {
	// Cloud region for targeted placement in format 'provider:region'.
	Region param.Field[string] `json:"region" api:"required"`
}

func (r VersionPlacementRegionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionPlacementRegionParam) implementsVersionPlacementUnionParam() {}

type VersionPlacementHostnameParam struct {
	// HTTP hostname for targeted placement.
	Hostname param.Field[string] `json:"hostname" api:"required"`
}

func (r VersionPlacementHostnameParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionPlacementHostnameParam) implementsVersionPlacementUnionParam() {}

type VersionPlacementHostParam struct {
	// TCP host and port for targeted placement.
	Host param.Field[string] `json:"host" api:"required"`
}

func (r VersionPlacementHostParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionPlacementHostParam) implementsVersionPlacementUnionParam() {}

type VersionPlacementObjectParam struct {
	// Targeted placement mode.
	Mode param.Field[VersionPlacementObjectMode] `json:"mode" api:"required"`
	// Cloud region for targeted placement in format 'provider:region'.
	Region param.Field[string] `json:"region" api:"required"`
}

func (r VersionPlacementObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionPlacementObjectParam) implementsVersionPlacementUnionParam() {}

type BetaWorkerVersionDeleteResponse struct {
	Errors   []BetaWorkerVersionDeleteResponseError   `json:"errors" api:"required"`
	Messages []BetaWorkerVersionDeleteResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success BetaWorkerVersionDeleteResponseSuccess `json:"success" api:"required"`
	JSON    betaWorkerVersionDeleteResponseJSON    `json:"-"`
}

// betaWorkerVersionDeleteResponseJSON contains the JSON metadata for the struct
// [BetaWorkerVersionDeleteResponse]
type betaWorkerVersionDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerVersionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerVersionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerVersionDeleteResponseError struct {
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           BetaWorkerVersionDeleteResponseErrorsSource `json:"source"`
	JSON             betaWorkerVersionDeleteResponseErrorJSON    `json:"-"`
}

// betaWorkerVersionDeleteResponseErrorJSON contains the JSON metadata for the
// struct [BetaWorkerVersionDeleteResponseError]
type betaWorkerVersionDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BetaWorkerVersionDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerVersionDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerVersionDeleteResponseErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    betaWorkerVersionDeleteResponseErrorsSourceJSON `json:"-"`
}

// betaWorkerVersionDeleteResponseErrorsSourceJSON contains the JSON metadata for
// the struct [BetaWorkerVersionDeleteResponseErrorsSource]
type betaWorkerVersionDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerVersionDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerVersionDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerVersionDeleteResponseMessage struct {
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           BetaWorkerVersionDeleteResponseMessagesSource `json:"source"`
	JSON             betaWorkerVersionDeleteResponseMessageJSON    `json:"-"`
}

// betaWorkerVersionDeleteResponseMessageJSON contains the JSON metadata for the
// struct [BetaWorkerVersionDeleteResponseMessage]
type betaWorkerVersionDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BetaWorkerVersionDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerVersionDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerVersionDeleteResponseMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    betaWorkerVersionDeleteResponseMessagesSourceJSON `json:"-"`
}

// betaWorkerVersionDeleteResponseMessagesSourceJSON contains the JSON metadata for
// the struct [BetaWorkerVersionDeleteResponseMessagesSource]
type betaWorkerVersionDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerVersionDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerVersionDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BetaWorkerVersionDeleteResponseSuccess bool

const (
	BetaWorkerVersionDeleteResponseSuccessTrue BetaWorkerVersionDeleteResponseSuccess = true
)

func (r BetaWorkerVersionDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case BetaWorkerVersionDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type BetaWorkerVersionNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Version   VersionParam        `json:"version" api:"required"`
	// If true, a deployment will be created that sends 100% of traffic to the new
	// version.
	Deploy param.Field[bool] `query:"deploy"`
}

func (r BetaWorkerVersionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Version)
}

// URLQuery serializes [BetaWorkerVersionNewParams]'s query parameters as
// `url.Values`.
func (r BetaWorkerVersionNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type BetaWorkerVersionNewResponseEnvelope struct {
	Errors   []BetaWorkerVersionNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []BetaWorkerVersionNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   Version                                        `json:"result" api:"required"`
	// Whether the API call was successful.
	Success BetaWorkerVersionNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    betaWorkerVersionNewResponseEnvelopeJSON    `json:"-"`
}

// betaWorkerVersionNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [BetaWorkerVersionNewResponseEnvelope]
type betaWorkerVersionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerVersionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerVersionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerVersionNewResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code" api:"required"`
	Message          string                                           `json:"message" api:"required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           BetaWorkerVersionNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             betaWorkerVersionNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// betaWorkerVersionNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [BetaWorkerVersionNewResponseEnvelopeErrors]
type betaWorkerVersionNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BetaWorkerVersionNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerVersionNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerVersionNewResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    betaWorkerVersionNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// betaWorkerVersionNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [BetaWorkerVersionNewResponseEnvelopeErrorsSource]
type betaWorkerVersionNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerVersionNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerVersionNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerVersionNewResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code" api:"required"`
	Message          string                                             `json:"message" api:"required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           BetaWorkerVersionNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             betaWorkerVersionNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// betaWorkerVersionNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [BetaWorkerVersionNewResponseEnvelopeMessages]
type betaWorkerVersionNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BetaWorkerVersionNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerVersionNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerVersionNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    betaWorkerVersionNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// betaWorkerVersionNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [BetaWorkerVersionNewResponseEnvelopeMessagesSource]
type betaWorkerVersionNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerVersionNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerVersionNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BetaWorkerVersionNewResponseEnvelopeSuccess bool

const (
	BetaWorkerVersionNewResponseEnvelopeSuccessTrue BetaWorkerVersionNewResponseEnvelopeSuccess = true
)

func (r BetaWorkerVersionNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BetaWorkerVersionNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BetaWorkerVersionListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Current page.
	Page param.Field[int64] `query:"page"`
	// Items per-page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [BetaWorkerVersionListParams]'s query parameters as
// `url.Values`.
func (r BetaWorkerVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type BetaWorkerVersionDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type BetaWorkerVersionGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Whether to include the `modules` property of the version in the response, which
	// contains code and sourcemap content and may add several megabytes to the
	// response size.
	Include param.Field[BetaWorkerVersionGetParamsInclude] `query:"include"`
}

// URLQuery serializes [BetaWorkerVersionGetParams]'s query parameters as
// `url.Values`.
func (r BetaWorkerVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Whether to include the `modules` property of the version in the response, which
// contains code and sourcemap content and may add several megabytes to the
// response size.
type BetaWorkerVersionGetParamsInclude string

const (
	BetaWorkerVersionGetParamsIncludeModules BetaWorkerVersionGetParamsInclude = "modules"
)

func (r BetaWorkerVersionGetParamsInclude) IsKnown() bool {
	switch r {
	case BetaWorkerVersionGetParamsIncludeModules:
		return true
	}
	return false
}

type BetaWorkerVersionGetResponseEnvelope struct {
	Errors   []BetaWorkerVersionGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []BetaWorkerVersionGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   Version                                        `json:"result" api:"required"`
	// Whether the API call was successful.
	Success BetaWorkerVersionGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    betaWorkerVersionGetResponseEnvelopeJSON    `json:"-"`
}

// betaWorkerVersionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [BetaWorkerVersionGetResponseEnvelope]
type betaWorkerVersionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerVersionGetResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code" api:"required"`
	Message          string                                           `json:"message" api:"required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           BetaWorkerVersionGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             betaWorkerVersionGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// betaWorkerVersionGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [BetaWorkerVersionGetResponseEnvelopeErrors]
type betaWorkerVersionGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BetaWorkerVersionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerVersionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerVersionGetResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    betaWorkerVersionGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// betaWorkerVersionGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [BetaWorkerVersionGetResponseEnvelopeErrorsSource]
type betaWorkerVersionGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerVersionGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerVersionGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerVersionGetResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code" api:"required"`
	Message          string                                             `json:"message" api:"required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           BetaWorkerVersionGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             betaWorkerVersionGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// betaWorkerVersionGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [BetaWorkerVersionGetResponseEnvelopeMessages]
type betaWorkerVersionGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BetaWorkerVersionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerVersionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BetaWorkerVersionGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    betaWorkerVersionGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// betaWorkerVersionGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [BetaWorkerVersionGetResponseEnvelopeMessagesSource]
type betaWorkerVersionGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BetaWorkerVersionGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r betaWorkerVersionGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BetaWorkerVersionGetResponseEnvelopeSuccess bool

const (
	BetaWorkerVersionGetResponseEnvelopeSuccessTrue BetaWorkerVersionGetResponseEnvelopeSuccess = true
)

func (r BetaWorkerVersionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BetaWorkerVersionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
