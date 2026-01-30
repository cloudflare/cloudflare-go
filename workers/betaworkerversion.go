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

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v6/shared"
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
		return
	}
	if workerID == "" {
		err = errors.New("missing required worker_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/workers/%s/versions", params.AccountID, workerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all versions for a Worker.
func (r *BetaWorkerVersionService) List(ctx context.Context, workerID string, params BetaWorkerVersionListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Version], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workerID == "" {
		err = errors.New("missing required worker_id parameter")
		return
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
		return
	}
	if workerID == "" {
		err = errors.New("missing required worker_id parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/workers/%s/versions/%s", body.AccountID, workerID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get details about a specific version.
func (r *BetaWorkerVersionService) Get(ctx context.Context, workerID string, versionID string, params BetaWorkerVersionGetParams, opts ...option.RequestOption) (res *Version, err error) {
	var env BetaWorkerVersionGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workerID == "" {
		err = errors.New("missing required worker_id parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/workers/%s/versions/%s", params.AccountID, workerID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Version struct {
	// Version identifier.
	ID string `json:"id,required" format:"uuid"`
	// When the version was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The integer version number, starting from one.
	Number int64 `json:"number,required"`
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
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate string `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Resource limits enforced at runtime.
	Limits VersionLimits `json:"limits"`
	// The name of the main module in the `modules` array (e.g. the name of the module
	// that exports a `fetch` handler).
	MainModule string `json:"main_module"`
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
	ID                 apijson.Field
	CreatedOn          apijson.Field
	Number             apijson.Field
	Annotations        apijson.Field
	Assets             apijson.Field
	Bindings           apijson.Field
	CompatibilityDate  apijson.Field
	CompatibilityFlags apijson.Field
	Limits             apijson.Field
	MainModule         apijson.Field
	Migrations         apijson.Field
	Modules            apijson.Field
	Placement          apijson.Field
	Source             apijson.Field
	StartupTimeMs      apijson.Field
	UsageModel         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Version) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionJSON) RawJSON() string {
	return r.raw
}

// Metadata about the version.
type VersionAnnotations struct {
	// Human-readable message about the version.
	WorkersMessage string `json:"workers/message"`
	// User-provided identifier for the version.
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
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsType `json:"type,required"`
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
	Format VersionBindingsFormat `json:"format"`
	// Name of the Vectorize index to bind to.
	IndexName string `json:"index_name"`
	// JSON data to use.
	Json string `json:"json"`
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction VersionBindingsJurisdiction `json:"jurisdiction"`
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
	// This field can have the runtime type of
	// [VersionBindingsWorkersBindingKindRatelimitSimple].
	Simple interface{} `json:"simple"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id"`
	// The text value to use.
	Text string `json:"text"`
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
	Simple                      apijson.Field
	StoreID                     apijson.Field
	Text                        apijson.Field
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
// [VersionBindingsWorkersBindingKindSecretKey],
// [VersionBindingsWorkersBindingKindWorkflow],
// [VersionBindingsWorkersBindingKindWasmModule].
func (r VersionBinding) AsUnion() VersionBindingsUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources.
//
// Union satisfied by [VersionBindingsWorkersBindingKindAI],
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
// [VersionBindingsWorkersBindingKindSecretKey],
// [VersionBindingsWorkersBindingKindWorkflow] or
// [VersionBindingsWorkersBindingKindWasmModule].
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
	)
}

type VersionBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindAIType `json:"type,required"`
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

type VersionBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset string `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindAnalyticsEngineType `json:"type,required"`
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
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindAssetsType `json:"type,required"`
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
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindBrowserType `json:"type,required"`
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
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindD1Type `json:"type,required"`
	JSON versionBindingsWorkersBindingKindD1JSON `json:"-"`
}

// versionBindingsWorkersBindingKindD1JSON contains the JSON metadata for the
// struct [VersionBindingsWorkersBindingKindD1]
type versionBindingsWorkersBindingKindD1JSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
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
	Name string `json:"name,required"`
	// The name of the file containing the data content. Only accepted for
	// `service worker syntax` Workers.
	Part string `json:"part,required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type VersionBindingsWorkersBindingKindDataBlobType `json:"type,required"`
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
	Name string `json:"name,required"`
	// The name of the dispatch namespace.
	Namespace string `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindDispatchNamespaceType `json:"type,required"`
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
	Params []string `json:"params"`
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

// Outbound worker.
type VersionBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
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
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindDurableObjectNamespaceType `json:"type,required"`
	// The exported class name of the Durable Object.
	ClassName string `json:"class_name"`
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
	Name        apijson.Field
	Type        apijson.Field
	ClassName   apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindHyperdriveType `json:"type,required"`
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
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindInheritType `json:"type,required"`
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
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindImagesType `json:"type,required"`
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
	Json string `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindJsonType `json:"type,required"`
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
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindKVNamespaceType `json:"type,required"`
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

type VersionBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindMTLSCertificateType `json:"type,required"`
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
	Name string `json:"name,required"`
	// The text value to use.
	Text string `json:"text,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindPlainTextType `json:"type,required"`
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
	Name string `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline string `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindPipelinesType `json:"type,required"`
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
	Name string `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName string `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindQueueType `json:"type,required"`
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
	Name string `json:"name,required"`
	// Identifier of the rate limit namespace to bind to.
	NamespaceID string `json:"namespace_id,required"`
	// The rate limit configuration.
	Simple VersionBindingsWorkersBindingKindRatelimitSimple `json:"simple,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindRatelimitType `json:"type,required"`
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
	Limit float64 `json:"limit,required"`
	// The period in seconds.
	Period int64                                                `json:"period,required"`
	JSON   versionBindingsWorkersBindingKindRatelimitSimpleJSON `json:"-"`
}

// versionBindingsWorkersBindingKindRatelimitSimpleJSON contains the JSON metadata
// for the struct [VersionBindingsWorkersBindingKindRatelimitSimple]
type versionBindingsWorkersBindingKindRatelimitSimpleJSON struct {
	Limit       apijson.Field
	Period      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindR2BucketType `json:"type,required"`
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
	VersionBindingsWorkersBindingKindR2BucketJurisdictionEu      VersionBindingsWorkersBindingKindR2BucketJurisdiction = "eu"
	VersionBindingsWorkersBindingKindR2BucketJurisdictionFedramp VersionBindingsWorkersBindingKindR2BucketJurisdiction = "fedramp"
)

func (r VersionBindingsWorkersBindingKindR2BucketJurisdiction) IsKnown() bool {
	switch r {
	case VersionBindingsWorkersBindingKindR2BucketJurisdictionEu, VersionBindingsWorkersBindingKindR2BucketJurisdictionFedramp:
		return true
	}
	return false
}

type VersionBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindSecretTextType `json:"type,required"`
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
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindSendEmailType `json:"type,required"`
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
	Name string `json:"name,required"`
	// Name of Worker to bind to.
	Service string `json:"service,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindServiceType `json:"type,required"`
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
	Name string `json:"name,required"`
	// The name of the file containing the text content. Only accepted for
	// `service worker syntax` Workers.
	Part string `json:"part,required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type VersionBindingsWorkersBindingKindTextBlobType `json:"type,required"`
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
	IndexName string `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindVectorizeType `json:"type,required"`
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
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindVersionMetadataType `json:"type,required"`
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
	Name string `json:"name,required"`
	// Name of the secret in the store.
	SecretName string `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID string `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindSecretsStoreSecretType `json:"type,required"`
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

type VersionBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm interface{} `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format VersionBindingsWorkersBindingKindSecretKeyFormat `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindSecretKeyType `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages []VersionBindingsWorkersBindingKindSecretKeyUsage `json:"usages,required"`
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
	Name string `json:"name,required"`
	// The kind of resource that the binding provides.
	Type VersionBindingsWorkersBindingKindWorkflowType `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName string `json:"workflow_name,required"`
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
	Name string `json:"name,required"`
	// The name of the file containing the WebAssembly module content. Only accepted
	// for `service worker syntax` Workers.
	Part string `json:"part,required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type VersionBindingsWorkersBindingKindWasmModuleType `json:"type,required"`
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

// The kind of resource that the binding provides.
type VersionBindingsType string

const (
	VersionBindingsTypeAI                     VersionBindingsType = "ai"
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
	VersionBindingsTypeSecretKey              VersionBindingsType = "secret_key"
	VersionBindingsTypeWorkflow               VersionBindingsType = "workflow"
	VersionBindingsTypeWasmModule             VersionBindingsType = "wasm_module"
)

func (r VersionBindingsType) IsKnown() bool {
	switch r {
	case VersionBindingsTypeAI, VersionBindingsTypeAnalyticsEngine, VersionBindingsTypeAssets, VersionBindingsTypeBrowser, VersionBindingsTypeD1, VersionBindingsTypeDataBlob, VersionBindingsTypeDispatchNamespace, VersionBindingsTypeDurableObjectNamespace, VersionBindingsTypeHyperdrive, VersionBindingsTypeInherit, VersionBindingsTypeImages, VersionBindingsTypeJson, VersionBindingsTypeKVNamespace, VersionBindingsTypeMTLSCertificate, VersionBindingsTypePlainText, VersionBindingsTypePipelines, VersionBindingsTypeQueue, VersionBindingsTypeRatelimit, VersionBindingsTypeR2Bucket, VersionBindingsTypeSecretText, VersionBindingsTypeSendEmail, VersionBindingsTypeService, VersionBindingsTypeTextBlob, VersionBindingsTypeVectorize, VersionBindingsTypeVersionMetadata, VersionBindingsTypeSecretsStoreSecret, VersionBindingsTypeSecretKey, VersionBindingsTypeWorkflow, VersionBindingsTypeWasmModule:
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
	VersionBindingsJurisdictionEu      VersionBindingsJurisdiction = "eu"
	VersionBindingsJurisdictionFedramp VersionBindingsJurisdiction = "fedramp"
)

func (r VersionBindingsJurisdiction) IsKnown() bool {
	switch r {
	case VersionBindingsJurisdictionEu, VersionBindingsJurisdictionFedramp:
		return true
	}
	return false
}

// Resource limits enforced at runtime.
type VersionLimits struct {
	// CPU time limit in milliseconds.
	CPUMs int64             `json:"cpu_ms,required"`
	JSON  versionLimitsJSON `json:"-"`
}

// versionLimitsJSON contains the JSON metadata for the struct [VersionLimits]
type versionLimitsJSON struct {
	CPUMs       apijson.Field
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
	ContentBase64 string `json:"content_base64,required" format:"byte"`
	// The content type of the module.
	ContentType string `json:"content_type,required"`
	// The name of the module.
	Name string            `json:"name,required"`
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
	Mode VersionPlacementModeMode `json:"mode,required"`
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
	Region string                     `json:"region,required"`
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
	Hostname string                       `json:"hostname,required"`
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
	Host string                   `json:"host,required"`
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
	Mode VersionPlacementObjectMode `json:"mode,required"`
	// Cloud region for targeted placement in format 'provider:region'.
	Region string                     `json:"region,required"`
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
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
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
	// Human-readable message about the version.
	WorkersMessage param.Field[string] `json:"workers/message"`
	// User-provided identifier for the version.
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
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsType] `json:"type,required"`
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
	Format param.Field[VersionBindingsFormat] `json:"format"`
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name"`
	// JSON data to use.
	Json param.Field[string] `json:"json"`
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction param.Field[VersionBindingsJurisdiction] `json:"jurisdiction"`
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
	Service param.Field[string]      `json:"service"`
	Simple  param.Field[interface{}] `json:"simple"`
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

func (r VersionBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingParam) implementsVersionBindingsUnionParam() {}

// A binding to allow the Worker to communicate with resources.
//
// Satisfied by [workers.VersionBindingsWorkersBindingKindAIParam],
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
// [workers.VersionBindingsWorkersBindingKindSecretKeyParam],
// [workers.VersionBindingsWorkersBindingKindWorkflowParam],
// [workers.VersionBindingsWorkersBindingKindWasmModuleParam],
// [VersionBindingParam].
type VersionBindingsUnionParam interface {
	implementsVersionBindingsUnionParam()
}

type VersionBindingsWorkersBindingKindAIParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindAIType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindAIParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindAIParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindAnalyticsEngineParam struct {
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindAnalyticsEngineType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindAnalyticsEngineParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindAnalyticsEngineParam) implementsVersionBindingsUnionParam() {
}

type VersionBindingsWorkersBindingKindAssetsParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindAssetsType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindAssetsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindAssetsParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindBrowserParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindBrowserType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindBrowserParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindBrowserParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindD1Param struct {
	// Identifier of the D1 database to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindD1Type] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindD1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindD1Param) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindDataBlobParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The name of the file containing the data content. Only accepted for
	// `service worker syntax` Workers.
	Part param.Field[string] `json:"part,required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type param.Field[VersionBindingsWorkersBindingKindDataBlobType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindDataBlobParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindDataBlobParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindDispatchNamespaceParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The name of the dispatch namespace.
	Namespace param.Field[string] `json:"namespace,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindDispatchNamespaceType] `json:"type,required"`
	// Outbound worker.
	Outbound param.Field[VersionBindingsWorkersBindingKindDispatchNamespaceOutboundParam] `json:"outbound"`
}

func (r VersionBindingsWorkersBindingKindDispatchNamespaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindDispatchNamespaceParam) implementsVersionBindingsUnionParam() {
}

// Outbound worker.
type VersionBindingsWorkersBindingKindDispatchNamespaceOutboundParam struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params param.Field[[]string] `json:"params"`
	// Outbound worker.
	Worker param.Field[VersionBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerParam] `json:"worker"`
}

func (r VersionBindingsWorkersBindingKindDispatchNamespaceOutboundParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker.
type VersionBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerParam struct {
	// Environment of the outbound worker.
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker.
	Service param.Field[string] `json:"service"`
}

func (r VersionBindingsWorkersBindingKindDispatchNamespaceOutboundWorkerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VersionBindingsWorkersBindingKindDurableObjectNamespaceParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindDurableObjectNamespaceType] `json:"type,required"`
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

func (r VersionBindingsWorkersBindingKindDurableObjectNamespaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindDurableObjectNamespaceParam) implementsVersionBindingsUnionParam() {
}

type VersionBindingsWorkersBindingKindHyperdriveParam struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID param.Field[string] `json:"id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindHyperdriveType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindHyperdriveParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindHyperdriveParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindInheritParam struct {
	// The name of the inherited binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindInheritType] `json:"type,required"`
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
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindImagesType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindImagesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindImagesParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindJsonParam struct {
	// JSON data to use.
	Json param.Field[string] `json:"json,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindJsonType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindJsonParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindJsonParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindKVNamespaceParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindKVNamespaceType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindKVNamespaceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindKVNamespaceParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindMTLSCertificateParam struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindMTLSCertificateType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindMTLSCertificateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindMTLSCertificateParam) implementsVersionBindingsUnionParam() {
}

type VersionBindingsWorkersBindingKindPlainTextParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The text value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindPlainTextType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindPlainTextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindPlainTextParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindPipelinesParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Pipeline to bind to.
	Pipeline param.Field[string] `json:"pipeline,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindPipelinesType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindPipelinesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindPipelinesParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindQueueParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindQueueType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindQueueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindQueueParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindRatelimitParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Identifier of the rate limit namespace to bind to.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
	// The rate limit configuration.
	Simple param.Field[VersionBindingsWorkersBindingKindRatelimitSimpleParam] `json:"simple,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindRatelimitType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindRatelimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindRatelimitParam) implementsVersionBindingsUnionParam() {}

// The rate limit configuration.
type VersionBindingsWorkersBindingKindRatelimitSimpleParam struct {
	// The limit (requests per period).
	Limit param.Field[float64] `json:"limit,required"`
	// The period in seconds.
	Period param.Field[int64] `json:"period,required"`
}

func (r VersionBindingsWorkersBindingKindRatelimitSimpleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VersionBindingsWorkersBindingKindR2BucketParam struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindR2BucketType] `json:"type,required"`
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
	Name param.Field[string] `json:"name,required"`
	// The secret value to use.
	Text param.Field[string] `json:"text,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindSecretTextType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindSecretTextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindSecretTextParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindSendEmailParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindSendEmailType] `json:"type,required"`
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
	Name param.Field[string] `json:"name,required"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindServiceType] `json:"type,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
}

func (r VersionBindingsWorkersBindingKindServiceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindServiceParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindTextBlobParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The name of the file containing the text content. Only accepted for
	// `service worker syntax` Workers.
	Part param.Field[string] `json:"part,required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type param.Field[VersionBindingsWorkersBindingKindTextBlobType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindTextBlobParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindTextBlobParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindVectorizeParam struct {
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindVectorizeType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindVectorizeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindVectorizeParam) implementsVersionBindingsUnionParam() {}

type VersionBindingsWorkersBindingKindVersionMetadataParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindVersionMetadataType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindVersionMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindVersionMetadataParam) implementsVersionBindingsUnionParam() {
}

type VersionBindingsWorkersBindingKindSecretsStoreSecretParam struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// Name of the secret in the store.
	SecretName param.Field[string] `json:"secret_name,required"`
	// ID of the store containing the secret.
	StoreID param.Field[string] `json:"store_id,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindSecretsStoreSecretType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindSecretsStoreSecretParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindSecretsStoreSecretParam) implementsVersionBindingsUnionParam() {
}

type VersionBindingsWorkersBindingKindSecretKeyParam struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm param.Field[interface{}] `json:"algorithm,required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[VersionBindingsWorkersBindingKindSecretKeyFormat] `json:"format,required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindSecretKeyType] `json:"type,required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages param.Field[[]VersionBindingsWorkersBindingKindSecretKeyUsage] `json:"usages,required"`
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
	Name param.Field[string] `json:"name,required"`
	// The kind of resource that the binding provides.
	Type param.Field[VersionBindingsWorkersBindingKindWorkflowType] `json:"type,required"`
	// Name of the Workflow to bind to.
	WorkflowName param.Field[string] `json:"workflow_name,required"`
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
	Name param.Field[string] `json:"name,required"`
	// The name of the file containing the WebAssembly module content. Only accepted
	// for `service worker syntax` Workers.
	Part param.Field[string] `json:"part,required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type param.Field[VersionBindingsWorkersBindingKindWasmModuleType] `json:"type,required"`
}

func (r VersionBindingsWorkersBindingKindWasmModuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionBindingsWorkersBindingKindWasmModuleParam) implementsVersionBindingsUnionParam() {}

// Resource limits enforced at runtime.
type VersionLimitsParam struct {
	// CPU time limit in milliseconds.
	CPUMs param.Field[int64] `json:"cpu_ms,required"`
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
	ContentBase64 param.Field[string] `json:"content_base64,required" format:"byte"`
	// The content type of the module.
	ContentType param.Field[string] `json:"content_type,required"`
	// The name of the module.
	Name param.Field[string] `json:"name,required"`
}

func (r VersionModuleParam) MarshalJSON() (data []byte, err error) {
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
	Mode param.Field[VersionPlacementModeMode] `json:"mode,required"`
}

func (r VersionPlacementModeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionPlacementModeParam) implementsVersionPlacementUnionParam() {}

type VersionPlacementRegionParam struct {
	// Cloud region for targeted placement in format 'provider:region'.
	Region param.Field[string] `json:"region,required"`
}

func (r VersionPlacementRegionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionPlacementRegionParam) implementsVersionPlacementUnionParam() {}

type VersionPlacementHostnameParam struct {
	// HTTP hostname for targeted placement.
	Hostname param.Field[string] `json:"hostname,required"`
}

func (r VersionPlacementHostnameParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionPlacementHostnameParam) implementsVersionPlacementUnionParam() {}

type VersionPlacementHostParam struct {
	// TCP host and port for targeted placement.
	Host param.Field[string] `json:"host,required"`
}

func (r VersionPlacementHostParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionPlacementHostParam) implementsVersionPlacementUnionParam() {}

type VersionPlacementObjectParam struct {
	// Targeted placement mode.
	Mode param.Field[VersionPlacementObjectMode] `json:"mode,required"`
	// Cloud region for targeted placement in format 'provider:region'.
	Region param.Field[string] `json:"region,required"`
}

func (r VersionPlacementObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VersionPlacementObjectParam) implementsVersionPlacementUnionParam() {}

type BetaWorkerVersionDeleteResponse struct {
	Errors   []BetaWorkerVersionDeleteResponseError   `json:"errors,required"`
	Messages []BetaWorkerVersionDeleteResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success BetaWorkerVersionDeleteResponseSuccess `json:"success,required"`
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
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
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
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
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
	AccountID param.Field[string] `path:"account_id,required"`
	Version   VersionParam        `json:"version,required"`
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
	Errors   []BetaWorkerVersionNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []BetaWorkerVersionNewResponseEnvelopeMessages `json:"messages,required"`
	Result   Version                                        `json:"result,required"`
	// Whether the API call was successful.
	Success BetaWorkerVersionNewResponseEnvelopeSuccess `json:"success,required"`
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
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
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
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
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
	AccountID param.Field[string] `path:"account_id,required"`
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
	AccountID param.Field[string] `path:"account_id,required"`
}

type BetaWorkerVersionGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
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
	Errors   []BetaWorkerVersionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []BetaWorkerVersionGetResponseEnvelopeMessages `json:"messages,required"`
	Result   Version                                        `json:"result,required"`
	// Whether the API call was successful.
	Success BetaWorkerVersionGetResponseEnvelopeSuccess `json:"success,required"`
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
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
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
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
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
