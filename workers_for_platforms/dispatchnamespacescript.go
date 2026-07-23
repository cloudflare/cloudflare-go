// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

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
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/workers"
	"github.com/tidwall/gjson"
)

// DispatchNamespaceScriptService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDispatchNamespaceScriptService] method instead.
type DispatchNamespaceScriptService struct {
	Options     []option.RequestOption
	AssetUpload *DispatchNamespaceScriptAssetUploadService
	Content     *DispatchNamespaceScriptContentService
	Settings    *DispatchNamespaceScriptSettingService
	Bindings    *DispatchNamespaceScriptBindingService
	Secrets     *DispatchNamespaceScriptSecretService
	Tags        *DispatchNamespaceScriptTagService
}

// NewDispatchNamespaceScriptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptService(opts ...option.RequestOption) (r *DispatchNamespaceScriptService) {
	r = &DispatchNamespaceScriptService{}
	r.Options = opts
	r.AssetUpload = NewDispatchNamespaceScriptAssetUploadService(opts...)
	r.Content = NewDispatchNamespaceScriptContentService(opts...)
	r.Settings = NewDispatchNamespaceScriptSettingService(opts...)
	r.Bindings = NewDispatchNamespaceScriptBindingService(opts...)
	r.Secrets = NewDispatchNamespaceScriptSecretService(opts...)
	r.Tags = NewDispatchNamespaceScriptTagService(opts...)
	return
}

// Upload a worker module to a Workers for Platforms namespace. You can find more
// about the multipart metadata on our docs:
// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.
func (r *DispatchNamespaceScriptService) Update(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptUpdateParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptUpdateResponse, err error) {
	var env DispatchNamespaceScriptUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return nil, err
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Delete a worker from a Workers for Platforms namespace. This call has no
// response body on a successful delete.
func (r *DispatchNamespaceScriptService) Delete(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptDeleteParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptDeleteResponse, err error) {
	var env DispatchNamespaceScriptDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return nil, err
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Fetch information about a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptGetParams, opts ...option.RequestOption) (res *Script, err error) {
	var env DispatchNamespaceScriptGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return nil, err
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Details about a worker uploaded to a Workers for Platforms namespace.
type Script struct {
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Name of the Workers for Platforms dispatch namespace.
	DispatchNamespace string `json:"dispatch_namespace"`
	// When the script was last modified.
	ModifiedOn time.Time      `json:"modified_on" format:"date-time"`
	Script     workers.Script `json:"script"`
	JSON       scriptJSON     `json:"-"`
}

// scriptJSON contains the JSON metadata for the struct [Script]
type scriptJSON struct {
	CreatedOn         apijson.Field
	DispatchNamespace apijson.Field
	ModifiedOn        apijson.Field
	Script            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Script) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptUpdateResponse struct {
	StartupTimeMs int64 `json:"startup_time_ms" api:"required"`
	// The name used to identify the script.
	ID string `json:"id"`
	// Global CacheW configuration for the Worker. When caching is on, the platform
	// provisions a `cloudflare.app` zone for the Worker. A `type: worker` entry in the
	// `exports` map can override this value for a single entrypoint.
	CacheOptions DispatchNamespaceScriptUpdateResponseCacheOptions `json:"cache_options"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate string `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags []string `json:"compatibility_flags"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The entry point for the script.
	EntryPoint string `json:"entry_point"`
	// Hashed script content, can be used in a If-None-Match header when updating.
	Etag string `json:"etag"`
	// Declarative exports for the Worker's most recent version, including Durable
	// Object classes (with their `storage` backend) and named Worker entrypoints.
	// Tombstoned lifecycle entries are omitted, so only live exports (`created` and
	// `expecting-transfer`) are returned.
	Exports map[string]DispatchNamespaceScriptUpdateResponseExport `json:"exports"`
	// The names of handlers exported as part of the default export.
	Handlers []string `json:"handlers"`
	// Whether a Worker contains assets.
	HasAssets bool `json:"has_assets"`
	// Whether a Worker contains modules.
	HasModules bool `json:"has_modules"`
	// The client most recently used to deploy this Worker.
	LastDeployedFrom string `json:"last_deployed_from"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// The tag of the Durable Object migration that was most recently applied for this
	// Worker.
	MigrationTag string `json:"migration_tag"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Named exports, such as Durable Object class implementations and named
	// entrypoints.
	NamedHandlers []DispatchNamespaceScriptUpdateResponseNamedHandler `json:"named_handlers"`
	// Observability settings for the Worker.
	Observability DispatchNamespaceScriptUpdateResponseObservability `json:"observability"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
	Placement DispatchNamespaceScriptUpdateResponsePlacement `json:"placement"`
	// Deprecated: deprecated
	PlacementMode DispatchNamespaceScriptUpdateResponsePlacementMode `json:"placement_mode"`
	// Deprecated: deprecated
	PlacementStatus DispatchNamespaceScriptUpdateResponsePlacementStatus `json:"placement_status"`
	// The immutable ID of the script.
	Tag string `json:"tag"`
	// Tags associated with the Worker.
	Tags []string `json:"tags" api:"nullable"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []workers.ConsumerScript `json:"tail_consumers" api:"nullable"`
	// Usage model for the Worker invocations.
	UsageModel DispatchNamespaceScriptUpdateResponseUsageModel `json:"usage_model"`
	JSON       dispatchNamespaceScriptUpdateResponseJSON       `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseJSON contains the JSON metadata for the
// struct [DispatchNamespaceScriptUpdateResponse]
type dispatchNamespaceScriptUpdateResponseJSON struct {
	StartupTimeMs      apijson.Field
	ID                 apijson.Field
	CacheOptions       apijson.Field
	CompatibilityDate  apijson.Field
	CompatibilityFlags apijson.Field
	CreatedOn          apijson.Field
	EntryPoint         apijson.Field
	Etag               apijson.Field
	Exports            apijson.Field
	Handlers           apijson.Field
	HasAssets          apijson.Field
	HasModules         apijson.Field
	LastDeployedFrom   apijson.Field
	Logpush            apijson.Field
	MigrationTag       apijson.Field
	ModifiedOn         apijson.Field
	NamedHandlers      apijson.Field
	Observability      apijson.Field
	Placement          apijson.Field
	PlacementMode      apijson.Field
	PlacementStatus    apijson.Field
	Tag                apijson.Field
	Tags               apijson.Field
	TailConsumers      apijson.Field
	UsageModel         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Global CacheW configuration for the Worker. When caching is on, the platform
// provisions a `cloudflare.app` zone for the Worker. A `type: worker` entry in the
// `exports` map can override this value for a single entrypoint.
type DispatchNamespaceScriptUpdateResponseCacheOptions struct {
	// Whether caching is enabled for this Worker.
	Enabled bool `json:"enabled" api:"required"`
	// Whether cached responses are shared across Worker version uploads. This is
	// independent of `enabled`. It can stay true while caching is off, so the
	// preference survives turning caching off and back on.
	CrossVersionCache bool                                                  `json:"cross_version_cache"`
	JSON              dispatchNamespaceScriptUpdateResponseCacheOptionsJSON `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseCacheOptionsJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptUpdateResponseCacheOptions]
type dispatchNamespaceScriptUpdateResponseCacheOptionsJSON struct {
	Enabled           apijson.Field
	CrossVersionCache apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseCacheOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseCacheOptionsJSON) RawJSON() string {
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
type DispatchNamespaceScriptUpdateResponseExport struct {
	// Marks this entry as a Worker entrypoint export.
	Type DispatchNamespaceScriptUpdateResponseExportsType `json:"type" api:"required"`
	// This field can have the runtime type of
	// [DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportCache].
	Cache interface{} `json:"cache"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container string `json:"container"`
	// Live export. May be omitted; defaults to `created`.
	State DispatchNamespaceScriptUpdateResponseExportsState `json:"state"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage DispatchNamespaceScriptUpdateResponseExportsStorage `json:"storage"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom string                                          `json:"transfer_from"`
	JSON         dispatchNamespaceScriptUpdateResponseExportJSON `json:"-"`
	union        DispatchNamespaceScriptUpdateResponseExportsUnion
}

// dispatchNamespaceScriptUpdateResponseExportJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptUpdateResponseExport]
type dispatchNamespaceScriptUpdateResponseExportJSON struct {
	Type         apijson.Field
	Cache        apijson.Field
	Container    apijson.Field
	State        apijson.Field
	Storage      apijson.Field
	TransferFrom apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r dispatchNamespaceScriptUpdateResponseExportJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptUpdateResponseExport) UnmarshalJSON(data []byte) (err error) {
	*r = DispatchNamespaceScriptUpdateResponseExport{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DispatchNamespaceScriptUpdateResponseExportsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExport],
// [DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExport],
// [DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExport],
// [DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExport],
// [DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExport],
// [DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExport].
func (r DispatchNamespaceScriptUpdateResponseExport) AsUnion() DispatchNamespaceScriptUpdateResponseExportsUnion {
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
// [DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExport],
// [DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExport],
// [DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExport],
// [DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExport],
// [DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExport]
// or
// [DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExport].
type DispatchNamespaceScriptUpdateResponseExportsUnion interface {
	implementsDispatchNamespaceScriptUpdateResponseExport()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptUpdateResponseExportsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExport{}),
			DiscriminatorValue: "worker",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExport{}),
			DiscriminatorValue: "durable-object",
		},
	)
}

// A named Worker entrypoint export (`type: worker`). Worker entrypoints are always
// live (`state: created`) and carry no storage or lifecycle fields. The optional
// `cache` block overrides the Worker's global `cache_options.enabled` for this
// entrypoint.
type DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExport struct {
	// Marks this entry as a Worker entrypoint export.
	Type DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportType `json:"type" api:"required"`
	// Cache override for this entrypoint. Overrides the Worker's global
	// `cache_options.enabled` for this entrypoint only.
	Cache DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportCache `json:"cache"`
	// Live export. May be omitted; defaults to `created`.
	State DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportState `json:"state"`
	JSON  dispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportJSON  `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExport]
type dispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportJSON struct {
	Type        apijson.Field
	Cache       apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExport) implementsDispatchNamespaceScriptUpdateResponseExport() {
}

// Marks this entry as a Worker entrypoint export.
type DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportType string

const (
	DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportTypeWorker DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportType = "worker"
)

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportTypeWorker:
		return true
	}
	return false
}

// Cache override for this entrypoint. Overrides the Worker's global
// `cache_options.enabled` for this entrypoint only.
type DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportCache struct {
	// Whether caching is enabled for this entrypoint.
	Enabled bool                                                                     `json:"enabled" api:"required"`
	JSON    dispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportCacheJSON `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportCacheJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportCache]
type dispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportCacheJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportCacheJSON) RawJSON() string {
	return r.raw
}

// Live export. May be omitted; defaults to `created`.
type DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportState string

const (
	DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportStateCreated DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportState = "created"
)

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportState) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseExportsWorkersWorkerExportStateCreated:
		return true
	}
	return false
}

// A live Durable Object export (`state: created`, the default). The platform
// auto-provisions the namespace on first deploy, matches it on subsequent deploys,
// and never mutates or deletes it as a side effect of a code-only change.
// `storage` is required; `renamed_to`, `transferred_to` and `transfer_from` are
// not allowed on a live entry.
type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExport struct {
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportStorage `json:"storage" api:"required"`
	// Marks this entry as a Durable Object export.
	Type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportType `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container string `json:"container"`
	// Live export. May be omitted; defaults to `created`.
	State DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportState `json:"state"`
	JSON  dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportJSON  `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExport]
type dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportJSON struct {
	Storage     apijson.Field
	Type        apijson.Field
	Container   apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExport) implementsDispatchNamespaceScriptUpdateResponseExport() {
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportStorage string

const (
	DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportStorageSqlite   DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportStorage = "sqlite"
	DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportStorageLegacyKV DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportStorage = "legacy-kv"
)

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportStorage) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportStorageSqlite, DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportStorageLegacyKV:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportType string

const (
	DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportTypeDurableObject DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportType = "durable-object"
)

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportTypeDurableObject:
		return true
	}
	return false
}

// Live export. May be omitted; defaults to `created`.
type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportState string

const (
	DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportStateCreated DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportState = "created"
)

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportState) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExportStateCreated:
		return true
	}
	return false
}

// A `deleted` tombstone: retires the provisioned namespace for this class and all
// of its data. The class must be absent from the uploaded code and no other Worker
// in the account may bind to the namespace, otherwise the deploy is rejected. No
// other fields are allowed. Deletion is irreversible.
type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExport struct {
	// Tombstone that deletes the namespace.
	State DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExportState `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExportType `json:"type" api:"required"`
	JSON dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExportJSON `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExportJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExport]
type dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExportJSON struct {
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExportJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExport) implementsDispatchNamespaceScriptUpdateResponseExport() {
}

// Tombstone that deletes the namespace.
type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExportState string

const (
	DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExportStateDeleted DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExportState = "deleted"
)

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExportState) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExportStateDeleted:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExportType string

const (
	DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExportTypeDurableObject DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExportType = "durable-object"
)

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExportType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectDeletedExportTypeDurableObject:
		return true
	}
	return false
}

// A `renamed` tombstone: rewrites the provisioned namespace's class name from this
// map key to `renamed_to`. The source class may stay in code during the rollout
// window (an info notice is emitted). `storage`, `transferred_to` and
// `transfer_from` are not allowed.
type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExport struct {
	// Tombstone that renames the namespace's class.
	State DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExportState `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExportType `json:"type" api:"required"`
	JSON dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExportJSON `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExportJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExport]
type dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExportJSON struct {
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExportJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExport) implementsDispatchNamespaceScriptUpdateResponseExport() {
}

// Tombstone that renames the namespace's class.
type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExportState string

const (
	DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExportStateRenamed DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExportState = "renamed"
)

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExportState) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExportStateRenamed:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExportType string

const (
	DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExportTypeDurableObject DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExportType = "durable-object"
)

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExportType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectRenamedExportTypeDurableObject:
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
type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExport struct {
	// Tombstone that transfers the namespace to another script.
	State DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExportState `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExportType `json:"type" api:"required"`
	JSON dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExportJSON `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExportJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExport]
type dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExportJSON struct {
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExportJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExport) implementsDispatchNamespaceScriptUpdateResponseExport() {
}

// Tombstone that transfers the namespace to another script.
type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExportState string

const (
	DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExportStateTransferred DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExportState = "transferred"
)

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExportState) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExportStateTransferred:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExportType string

const (
	DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExportTypeDurableObject DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExportType = "durable-object"
)

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExportType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectTransferredExportTypeDurableObject:
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
type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExport struct {
	// Target side of a two-phase transfer.
	State DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportState `json:"state" api:"required"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportStorage `json:"storage" api:"required"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom string `json:"transfer_from" api:"required"`
	// Marks this entry as a Durable Object export.
	Type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportType `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object once the transfer settles. Valid only on live entries.
	Container string                                                                                      `json:"container"`
	JSON      dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportJSON `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExport]
type dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportJSON struct {
	State        apijson.Field
	Storage      apijson.Field
	TransferFrom apijson.Field
	Type         apijson.Field
	Container    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExport) implementsDispatchNamespaceScriptUpdateResponseExport() {
}

// Target side of a two-phase transfer.
type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportState string

const (
	DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportStateExpectingTransfer DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportState = "expecting-transfer"
)

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportState) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportStateExpectingTransfer:
		return true
	}
	return false
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportStorage string

const (
	DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportStorageSqlite   DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportStorage = "sqlite"
	DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportStorageLegacyKV DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportStorage = "legacy-kv"
)

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportStorage) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportStorageSqlite, DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportStorageLegacyKV:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportType string

const (
	DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportTypeDurableObject DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportType = "durable-object"
)

func (r DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseExportsWorkersDurableObjectExpectingTransferExportTypeDurableObject:
		return true
	}
	return false
}

// Marks this entry as a Worker entrypoint export.
type DispatchNamespaceScriptUpdateResponseExportsType string

const (
	DispatchNamespaceScriptUpdateResponseExportsTypeWorker        DispatchNamespaceScriptUpdateResponseExportsType = "worker"
	DispatchNamespaceScriptUpdateResponseExportsTypeDurableObject DispatchNamespaceScriptUpdateResponseExportsType = "durable-object"
)

func (r DispatchNamespaceScriptUpdateResponseExportsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseExportsTypeWorker, DispatchNamespaceScriptUpdateResponseExportsTypeDurableObject:
		return true
	}
	return false
}

// Live export. May be omitted; defaults to `created`.
type DispatchNamespaceScriptUpdateResponseExportsState string

const (
	DispatchNamespaceScriptUpdateResponseExportsStateCreated           DispatchNamespaceScriptUpdateResponseExportsState = "created"
	DispatchNamespaceScriptUpdateResponseExportsStateDeleted           DispatchNamespaceScriptUpdateResponseExportsState = "deleted"
	DispatchNamespaceScriptUpdateResponseExportsStateRenamed           DispatchNamespaceScriptUpdateResponseExportsState = "renamed"
	DispatchNamespaceScriptUpdateResponseExportsStateTransferred       DispatchNamespaceScriptUpdateResponseExportsState = "transferred"
	DispatchNamespaceScriptUpdateResponseExportsStateExpectingTransfer DispatchNamespaceScriptUpdateResponseExportsState = "expecting-transfer"
)

func (r DispatchNamespaceScriptUpdateResponseExportsState) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseExportsStateCreated, DispatchNamespaceScriptUpdateResponseExportsStateDeleted, DispatchNamespaceScriptUpdateResponseExportsStateRenamed, DispatchNamespaceScriptUpdateResponseExportsStateTransferred, DispatchNamespaceScriptUpdateResponseExportsStateExpectingTransfer:
		return true
	}
	return false
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type DispatchNamespaceScriptUpdateResponseExportsStorage string

const (
	DispatchNamespaceScriptUpdateResponseExportsStorageSqlite   DispatchNamespaceScriptUpdateResponseExportsStorage = "sqlite"
	DispatchNamespaceScriptUpdateResponseExportsStorageLegacyKV DispatchNamespaceScriptUpdateResponseExportsStorage = "legacy-kv"
)

func (r DispatchNamespaceScriptUpdateResponseExportsStorage) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseExportsStorageSqlite, DispatchNamespaceScriptUpdateResponseExportsStorageLegacyKV:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateResponseNamedHandler struct {
	// The names of handlers exported as part of the named export.
	Handlers []string `json:"handlers"`
	// The name of the export.
	Name string                                                `json:"name"`
	JSON dispatchNamespaceScriptUpdateResponseNamedHandlerJSON `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseNamedHandlerJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptUpdateResponseNamedHandler]
type dispatchNamespaceScriptUpdateResponseNamedHandlerJSON struct {
	Handlers    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseNamedHandler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseNamedHandlerJSON) RawJSON() string {
	return r.raw
}

// Observability settings for the Worker.
type DispatchNamespaceScriptUpdateResponseObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled bool `json:"enabled" api:"required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate float64 `json:"head_sampling_rate" api:"nullable"`
	// Log settings for the Worker.
	Logs DispatchNamespaceScriptUpdateResponseObservabilityLogs `json:"logs" api:"nullable"`
	// Trace settings for the Worker.
	Traces DispatchNamespaceScriptUpdateResponseObservabilityTraces `json:"traces" api:"nullable"`
	JSON   dispatchNamespaceScriptUpdateResponseObservabilityJSON   `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseObservabilityJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptUpdateResponseObservability]
type dispatchNamespaceScriptUpdateResponseObservabilityJSON struct {
	Enabled          apijson.Field
	HeadSamplingRate apijson.Field
	Logs             apijson.Field
	Traces           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseObservabilityJSON) RawJSON() string {
	return r.raw
}

// Log settings for the Worker.
type DispatchNamespaceScriptUpdateResponseObservabilityLogs struct {
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
	Persist bool                                                       `json:"persist"`
	JSON    dispatchNamespaceScriptUpdateResponseObservabilityLogsJSON `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseObservabilityLogsJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptUpdateResponseObservabilityLogs]
type dispatchNamespaceScriptUpdateResponseObservabilityLogsJSON struct {
	Enabled          apijson.Field
	InvocationLogs   apijson.Field
	Destinations     apijson.Field
	HeadSamplingRate apijson.Field
	Persist          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// Trace settings for the Worker.
type DispatchNamespaceScriptUpdateResponseObservabilityTraces struct {
	// A list of destinations where traces will be exported to.
	Destinations []string `json:"destinations"`
	// Whether traces are enabled for the Worker.
	Enabled bool `json:"enabled"`
	// The sampling rate for traces. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate float64 `json:"head_sampling_rate" api:"nullable"`
	// Whether trace persistence is enabled for the Worker.
	Persist bool `json:"persist"`
	// Controls how inbound trace context (traceparent/tracestate) headers on incoming
	// requests are handled. "authenticated" (default) honors inbound trace context
	// only when accompanied by a valid trace auth token. "accept" unconditionally
	// accepts inbound trace context. Requires the trace propagation feature to be
	// enabled.
	PropagationPolicy DispatchNamespaceScriptUpdateResponseObservabilityTracesPropagationPolicy `json:"propagation_policy"`
	JSON              dispatchNamespaceScriptUpdateResponseObservabilityTracesJSON              `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseObservabilityTracesJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptUpdateResponseObservabilityTraces]
type dispatchNamespaceScriptUpdateResponseObservabilityTracesJSON struct {
	Destinations      apijson.Field
	Enabled           apijson.Field
	HeadSamplingRate  apijson.Field
	Persist           apijson.Field
	PropagationPolicy apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseObservabilityTraces) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseObservabilityTracesJSON) RawJSON() string {
	return r.raw
}

// Controls how inbound trace context (traceparent/tracestate) headers on incoming
// requests are handled. "authenticated" (default) honors inbound trace context
// only when accompanied by a valid trace auth token. "accept" unconditionally
// accepts inbound trace context. Requires the trace propagation feature to be
// enabled.
type DispatchNamespaceScriptUpdateResponseObservabilityTracesPropagationPolicy string

const (
	DispatchNamespaceScriptUpdateResponseObservabilityTracesPropagationPolicyAuthenticated DispatchNamespaceScriptUpdateResponseObservabilityTracesPropagationPolicy = "authenticated"
	DispatchNamespaceScriptUpdateResponseObservabilityTracesPropagationPolicyAccept        DispatchNamespaceScriptUpdateResponseObservabilityTracesPropagationPolicy = "accept"
)

func (r DispatchNamespaceScriptUpdateResponseObservabilityTracesPropagationPolicy) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseObservabilityTracesPropagationPolicyAuthenticated, DispatchNamespaceScriptUpdateResponseObservabilityTracesPropagationPolicyAccept:
		return true
	}
	return false
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
type DispatchNamespaceScriptUpdateResponsePlacement struct {
	// TCP host and port for targeted placement.
	Host string `json:"host"`
	// HTTP hostname for targeted placement.
	Hostname string `json:"hostname"`
	// The last time the script was analyzed for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	LastAnalyzedAt time.Time `json:"last_analyzed_at" format:"date-time"`
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode DispatchNamespaceScriptUpdateResponsePlacementMode `json:"mode"`
	// Cloud region for targeted placement in format 'provider:region'.
	Region string `json:"region"`
	// Status of
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Status DispatchNamespaceScriptUpdateResponsePlacementStatus `json:"status"`
	// This field can have the runtime type of
	// [[]DispatchNamespaceScriptUpdateResponsePlacementObjectTarget].
	Target interface{}                                        `json:"target"`
	JSON   dispatchNamespaceScriptUpdateResponsePlacementJSON `json:"-"`
	union  DispatchNamespaceScriptUpdateResponsePlacementUnion
}

// dispatchNamespaceScriptUpdateResponsePlacementJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptUpdateResponsePlacement]
type dispatchNamespaceScriptUpdateResponsePlacementJSON struct {
	Host           apijson.Field
	Hostname       apijson.Field
	LastAnalyzedAt apijson.Field
	Mode           apijson.Field
	Region         apijson.Field
	Status         apijson.Field
	Target         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r dispatchNamespaceScriptUpdateResponsePlacementJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptUpdateResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	*r = DispatchNamespaceScriptUpdateResponsePlacement{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DispatchNamespaceScriptUpdateResponsePlacementUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DispatchNamespaceScriptUpdateResponsePlacementObject],
// [DispatchNamespaceScriptUpdateResponsePlacementObject],
// [DispatchNamespaceScriptUpdateResponsePlacementObject],
// [DispatchNamespaceScriptUpdateResponsePlacementObject],
// [DispatchNamespaceScriptUpdateResponsePlacementObject],
// [DispatchNamespaceScriptUpdateResponsePlacementObject],
// [DispatchNamespaceScriptUpdateResponsePlacementObject],
// [DispatchNamespaceScriptUpdateResponsePlacementObject].
func (r DispatchNamespaceScriptUpdateResponsePlacement) AsUnion() DispatchNamespaceScriptUpdateResponsePlacementUnion {
	return r.union
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
//
// Union satisfied by [DispatchNamespaceScriptUpdateResponsePlacementObject],
// [DispatchNamespaceScriptUpdateResponsePlacementObject],
// [DispatchNamespaceScriptUpdateResponsePlacementObject],
// [DispatchNamespaceScriptUpdateResponsePlacementObject],
// [DispatchNamespaceScriptUpdateResponsePlacementObject],
// [DispatchNamespaceScriptUpdateResponsePlacementObject],
// [DispatchNamespaceScriptUpdateResponsePlacementObject] or
// [DispatchNamespaceScriptUpdateResponsePlacementObject].
type DispatchNamespaceScriptUpdateResponsePlacementUnion interface {
	implementsDispatchNamespaceScriptUpdateResponsePlacement()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptUpdateResponsePlacementUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptUpdateResponsePlacementObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptUpdateResponsePlacementObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptUpdateResponsePlacementObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptUpdateResponsePlacementObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptUpdateResponsePlacementObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptUpdateResponsePlacementObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptUpdateResponsePlacementObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptUpdateResponsePlacementObject{}),
		},
	)
}

type DispatchNamespaceScriptUpdateResponsePlacementObject struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode DispatchNamespaceScriptUpdateResponsePlacementObjectMode `json:"mode" api:"required"`
	// The last time the script was analyzed for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	LastAnalyzedAt time.Time `json:"last_analyzed_at" format:"date-time"`
	// Status of
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Status DispatchNamespaceScriptUpdateResponsePlacementObjectStatus `json:"status"`
	JSON   dispatchNamespaceScriptUpdateResponsePlacementObjectJSON   `json:"-"`
}

// dispatchNamespaceScriptUpdateResponsePlacementObjectJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptUpdateResponsePlacementObject]
type dispatchNamespaceScriptUpdateResponsePlacementObjectJSON struct {
	Mode           apijson.Field
	LastAnalyzedAt apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponsePlacementObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponsePlacementObjectJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptUpdateResponsePlacementObject) implementsDispatchNamespaceScriptUpdateResponsePlacement() {
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateResponsePlacementObjectMode string

const (
	DispatchNamespaceScriptUpdateResponsePlacementObjectModeSmart DispatchNamespaceScriptUpdateResponsePlacementObjectMode = "smart"
)

func (r DispatchNamespaceScriptUpdateResponsePlacementObjectMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponsePlacementObjectModeSmart:
		return true
	}
	return false
}

// Status of
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateResponsePlacementObjectStatus string

const (
	DispatchNamespaceScriptUpdateResponsePlacementObjectStatusSuccess                 DispatchNamespaceScriptUpdateResponsePlacementObjectStatus = "SUCCESS"
	DispatchNamespaceScriptUpdateResponsePlacementObjectStatusUnsupportedApplication  DispatchNamespaceScriptUpdateResponsePlacementObjectStatus = "UNSUPPORTED_APPLICATION"
	DispatchNamespaceScriptUpdateResponsePlacementObjectStatusInsufficientInvocations DispatchNamespaceScriptUpdateResponsePlacementObjectStatus = "INSUFFICIENT_INVOCATIONS"
)

func (r DispatchNamespaceScriptUpdateResponsePlacementObjectStatus) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponsePlacementObjectStatusSuccess, DispatchNamespaceScriptUpdateResponsePlacementObjectStatusUnsupportedApplication, DispatchNamespaceScriptUpdateResponsePlacementObjectStatusInsufficientInvocations:
		return true
	}
	return false
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateResponsePlacementMode string

const (
	DispatchNamespaceScriptUpdateResponsePlacementModeSmart    DispatchNamespaceScriptUpdateResponsePlacementMode = "smart"
	DispatchNamespaceScriptUpdateResponsePlacementModeTargeted DispatchNamespaceScriptUpdateResponsePlacementMode = "targeted"
)

func (r DispatchNamespaceScriptUpdateResponsePlacementMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponsePlacementModeSmart, DispatchNamespaceScriptUpdateResponsePlacementModeTargeted:
		return true
	}
	return false
}

// Status of
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateResponsePlacementStatus string

const (
	DispatchNamespaceScriptUpdateResponsePlacementStatusSuccess                 DispatchNamespaceScriptUpdateResponsePlacementStatus = "SUCCESS"
	DispatchNamespaceScriptUpdateResponsePlacementStatusUnsupportedApplication  DispatchNamespaceScriptUpdateResponsePlacementStatus = "UNSUPPORTED_APPLICATION"
	DispatchNamespaceScriptUpdateResponsePlacementStatusInsufficientInvocations DispatchNamespaceScriptUpdateResponsePlacementStatus = "INSUFFICIENT_INVOCATIONS"
)

func (r DispatchNamespaceScriptUpdateResponsePlacementStatus) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponsePlacementStatusSuccess, DispatchNamespaceScriptUpdateResponsePlacementStatusUnsupportedApplication, DispatchNamespaceScriptUpdateResponsePlacementStatusInsufficientInvocations:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type DispatchNamespaceScriptUpdateResponseUsageModel string

const (
	DispatchNamespaceScriptUpdateResponseUsageModelStandard DispatchNamespaceScriptUpdateResponseUsageModel = "standard"
	DispatchNamespaceScriptUpdateResponseUsageModelBundled  DispatchNamespaceScriptUpdateResponseUsageModel = "bundled"
	DispatchNamespaceScriptUpdateResponseUsageModelUnbound  DispatchNamespaceScriptUpdateResponseUsageModel = "unbound"
)

func (r DispatchNamespaceScriptUpdateResponseUsageModel) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseUsageModelStandard, DispatchNamespaceScriptUpdateResponseUsageModelBundled, DispatchNamespaceScriptUpdateResponseUsageModelUnbound:
		return true
	}
	return false
}

type DispatchNamespaceScriptDeleteResponse = interface{}

type DispatchNamespaceScriptUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// JSON-encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[DispatchNamespaceScriptUpdateParamsMetadata] `json:"metadata" api:"required"`
	// When set to "strict", the upload will fail if any `inherit` type bindings cannot
	// be resolved against the previous version of the script. Without this,
	// unresolvable inherit bindings are silently dropped.
	BindingsInherit param.Field[DispatchNamespaceScriptUpdateParamsBindingsInherit] `query:"bindings_inherit"`
	// An array of modules (often JavaScript files) comprising a Worker script. At
	// least one module must be present and referenced in the metadata as `main_module`
	// or `body_part` by filename.<br/>Possible Content-Type(s) are:
	// `application/javascript+module`, `text/javascript+module`,
	// `application/javascript`, `text/javascript`, `text/x-python`,
	// `text/x-python-requirement`, `application/wasm`, `text/plain`,
	// `application/octet-stream`, `application/source-map`.
	Files param.Field[[]io.Reader] `json:"files" format:"binary"`
}

func (r DispatchNamespaceScriptUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [DispatchNamespaceScriptUpdateParams]'s query parameters as
// `url.Values`.
func (r DispatchNamespaceScriptUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// JSON-encoded metadata about the uploaded parts and Worker configuration.
type DispatchNamespaceScriptUpdateParamsMetadata struct {
	// Configuration for assets within a Worker.
	Assets param.Field[DispatchNamespaceScriptUpdateParamsMetadataAssets] `json:"assets"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings param.Field[[]DispatchNamespaceScriptUpdateParamsMetadataBindingUnion] `json:"bindings"`
	// Name of the uploaded file that contains the script (e.g. the file adding a
	// listener to the `fetch` event). Indicates a `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
	// Global CacheW configuration for the Worker. When caching is on, the platform
	// provisions a `cloudflare.app` zone for the Worker. A `type: worker` entry in the
	// `exports` map can override this value for a single entrypoint.
	CacheOptions param.Field[DispatchNamespaceScriptUpdateParamsMetadataCacheOptions] `json:"cache_options"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Declarative exports for the Worker. Worker entrypoint entries (`type: worker`)
	// carry cache configuration for that entrypoint.
	Exports param.Field[map[string]DispatchNamespaceScriptUpdateParamsMetadataExportsUnion] `json:"exports"`
	// Retain assets which exist for a previously uploaded Worker version; used in lieu
	// of providing a completion token. An explicit `assets` upload takes precedence
	// over `keep_assets`.
	KeepAssets param.Field[bool] `json:"keep_assets"`
	// List of binding types to keep from previous_upload.
	KeepBindings param.Field[[]string] `json:"keep_bindings"`
	// Limits to apply for this Worker.
	Limits param.Field[DispatchNamespaceScriptUpdateParamsMetadataLimits] `json:"limits"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Name of the uploaded file that contains the main module (e.g. the file exporting
	// a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[DispatchNamespaceScriptUpdateParamsMetadataMigrationsUnion] `json:"migrations"`
	// Observability settings for the Worker.
	Observability param.Field[DispatchNamespaceScriptUpdateParamsMetadataObservability] `json:"observability"`
	// The list of npm packages that were installed and used when this Worker version
	// was built.
	PackageDependencies param.Field[[]DispatchNamespaceScriptUpdateParamsMetadataPackageDependency] `json:"package_dependencies"`
	// Configuration for
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
	Placement param.Field[DispatchNamespaceScriptUpdateParamsMetadataPlacementUnion] `json:"placement"`
	// List of strings to use as tags for this Worker.
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]workers.ConsumerScriptParam] `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[DispatchNamespaceScriptUpdateParamsMetadataUsageModel] `json:"usage_model"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker.
type DispatchNamespaceScriptUpdateParamsMetadataAssets struct {
	// Configuration for assets within a Worker.
	Config param.Field[DispatchNamespaceScriptUpdateParamsMetadataAssetsConfig] `json:"config"`
	// Token provided upon successful upload of all files from a registered manifest.
	JWT param.Field[string] `json:"jwt"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker.
type DispatchNamespaceScriptUpdateParamsMetadataAssetsConfig struct {
	// The contents of a \_headers file (used to attach custom headers on asset
	// responses).
	Headers param.Field[string] `json:"_headers"`
	// The contents of a \_redirects file (used to apply redirects or proxy paths ahead
	// of asset serving).
	Redirects param.Field[string] `json:"_redirects"`
	// Determines the redirects and rewrites of requests for HTML content.
	HTMLHandling param.Field[DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling] `json:"html_handling"`
	// Determines the response when a request does not match a static asset, and there
	// is no Worker script.
	NotFoundHandling param.Field[DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling] `json:"not_found_handling"`
	// Contains a list path rules to control routing to either the Worker or assets.
	// Glob (\*) and negative (!) rules are supported. Rules must start with either '/'
	// or '!/'. At least one non-negative rule must be provided, and negative rules
	// have higher precedence than non-negative rules.
	RunWorkerFirst param.Field[DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstUnion] `json:"run_worker_first"`
	// When true and the incoming request matches an asset, that will be served instead
	// of invoking the Worker script. When false, requests will always invoke the
	// Worker script.
	//
	// Deprecated: deprecated
	ServeDirectly param.Field[bool] `json:"serve_directly"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataAssetsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the redirects and rewrites of requests for HTML content.
type DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling string

const (
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingAutoTrailingSlash  DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "auto-trailing-slash"
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingForceTrailingSlash DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "force-trailing-slash"
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingDropTrailingSlash  DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "drop-trailing-slash"
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingNone               DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling = "none"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandling) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingAutoTrailingSlash, DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingForceTrailingSlash, DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingDropTrailingSlash, DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigHTMLHandlingNone:
		return true
	}
	return false
}

// Determines the response when a request does not match a static asset, and there
// is no Worker script.
type DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling string

const (
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingNone                  DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling = "none"
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling404Page               DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling = "404-page"
	DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingSinglePageApplication DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling = "single-page-application"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingNone, DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandling404Page, DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigNotFoundHandlingSinglePageApplication:
		return true
	}
	return false
}

// Contains a list path rules to control routing to either the Worker or assets.
// Glob (\*) and negative (!) rules are supported. Rules must start with either '/'
// or '!/'. At least one non-negative rule must be provided, and negative rules
// have higher precedence than non-negative rules.
//
// Satisfied by
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstArray],
// [shared.UnionBool].
type DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstUnion interface {
	ImplementsDispatchNamespaceScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstUnion()
}

type DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstArray []string

func (r DispatchNamespaceScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstArray) ImplementsDispatchNamespaceScriptUpdateParamsMetadataAssetsConfigRunWorkerFirstUnion() {
}

// A binding to allow the Worker to communicate with resources.
type DispatchNamespaceScriptUpdateParamsMetadataBinding struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsType] `json:"type" api:"required"`
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
	Format param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsFormat] `json:"format"`
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name"`
	// The user-chosen instance name. Must exist at deploy time. The worker can search,
	// chat, update, and manage items/jobs on this instance.
	InstanceName param.Field[string]      `json:"instance_name"`
	Json         param.Field[interface{}] `json:"json"`
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsJurisdiction] `json:"jurisdiction"`
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

func (r DispatchNamespaceScriptUpdateParamsMetadataBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBinding) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// A binding to allow the Worker to communicate with resources.
//
// Satisfied by
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAI],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearch],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearchNamespace],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowser],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDataBlob],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindInherit],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindImages],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJson],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMedia],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelines],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindRatelimit],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSendEmail],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindService],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTextBlob],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindFlagship],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKey],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflow],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWasmModule],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCService],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCNetwork],
// [DispatchNamespaceScriptUpdateParamsMetadataBinding].
type DispatchNamespaceScriptUpdateParamsMetadataBindingUnion interface {
	implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion()
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAI) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAI DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType = "ai"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAIType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAITypeAI:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearch struct {
	// The user-chosen instance name. Must exist at deploy time. The worker can search,
	// chat, update, and manage items/jobs on this instance.
	InstanceName param.Field[string] `json:"instance_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearchType] `json:"type" api:"required"`
	// The namespace the instance belongs to. Defaults to "default" if omitted.
	// Customers who don't use namespaces can simply omit this field.
	Namespace param.Field[string] `json:"namespace"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearch) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearchType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearchTypeAISearch DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearchType = "ai_search"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearchType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearchTypeAISearch:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearchNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The user-chosen namespace name. Must exist before deploy -- Wrangler handles
	// auto-creation on deploy failure (R2 bucket pattern). The "default" namespace is
	// auto-created by config-api for new accounts. Grants full access (CRUD + search +
	// chat) to all instances within the namespace.
	Namespace param.Field[string] `json:"namespace" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearchNamespaceType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearchNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearchNamespace) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearchNamespaceType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearchNamespaceTypeAISearchNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearchNamespaceType = "ai_search_namespace"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearchNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAISearchNamespaceTypeAISearchNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngine) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType = "analytics_engine"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAnalyticsEngineTypeAnalyticsEngine:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssets) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAssets DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType = "assets"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindAssetsTypeAssets:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowser struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowser) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserTypeBrowser DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserType = "browser"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindBrowserTypeBrowser:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1 struct {
	// Identifier of the D1 database to bind to.
	DatabaseID param.Field[string] `json:"database_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type] `json:"type" api:"required"`
	// Identifier of the D1 database to bind to.
	//
	// Deprecated: This property has been renamed to `database_id`.
	ID param.Field[string] `json:"id"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeD1 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type = "d1"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1Type) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindD1TypeD1:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDataBlob struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The name of the file containing the data content. Only accepted for
	// `service worker syntax` Workers.
	Part param.Field[string] `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDataBlobType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDataBlob) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDataBlob) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDataBlobType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDataBlobTypeDataBlob DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDataBlobType = "data_blob"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDataBlobType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDataBlobTypeDataBlob:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The name of the dispatch namespace.
	Namespace param.Field[string] `json:"namespace" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType] `json:"type" api:"required"`
	// Outbound worker.
	Outbound param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound] `json:"outbound"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespace) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType = "dispatch_namespace"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters.
	Params param.Field[[]DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundParam] `json:"params"`
	// Outbound worker.
	Worker param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker] `json:"worker"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundParam struct {
	// Name of the parameter.
	Name param.Field[string] `json:"name" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Entrypoint to invoke on the outbound worker.
	Entrypoint param.Field[string] `json:"entrypoint"`
	// Environment of the outbound worker.
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker.
	Service param.Field[string] `json:"service"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType] `json:"type" api:"required"`
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

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespace) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType = "durable_object_namespace"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceTypeDurableObjectNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive struct {
	// Identifier of the Hyperdrive connection to bind to.
	ID param.Field[string] `json:"id" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdrive) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType = "hyperdrive"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindHyperdriveTypeHyperdrive:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindInherit struct {
	// The name of the inherited binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindInheritType] `json:"type" api:"required"`
	// The old name of the inherited binding. If set, the binding will be renamed from
	// `old_name` to `name` in the new version. If not set, the binding will keep the
	// same name between versions.
	OldName param.Field[string] `json:"old_name"`
	// Identifier for the version to inherit the binding from, which can be the version
	// ID or the literal "latest" to inherit from the latest version. Defaults to
	// inheriting the binding from the latest version.
	VersionID param.Field[string] `json:"version_id"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindInherit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindInherit) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindInheritType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindInheritTypeInherit DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindInheritType = "inherit"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindInheritType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindInheritTypeInherit:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindImages struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindImagesType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindImages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindImages) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindImagesType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindImagesTypeImages DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindImagesType = "images"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindImagesType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindImagesTypeImages:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJson struct {
	// JSON data to use.
	Json param.Field[interface{}] `json:"json" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJson) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeJson DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType = "json"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindJsonTypeJson:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespace) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType = "kv_namespace"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindKVNamespaceTypeKVNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMedia struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMediaType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMedia) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMedia) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMediaType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMediaTypeMedia DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMediaType = "media"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMediaType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMediaTypeMedia:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificate) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType = "mtls_certificate"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindMTLSCertificateTypeMTLSCertificate:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The text value to use.
	Text param.Field[string] `json:"text" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainText) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType = "plain_text"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPlainTextTypePlainText:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelines struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Name of the Pipeline to bind to.
	Pipeline param.Field[string] `json:"pipeline" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelinesType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelines) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelines) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelinesType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelinesTypePipelines DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelinesType = "pipelines"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelinesType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindPipelinesTypePipelines:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueue) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeQueue DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType = "queue"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindQueueTypeQueue:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindRatelimit struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Identifier of the rate limit namespace to bind to.
	NamespaceID param.Field[string] `json:"namespace_id" api:"required"`
	// The rate limit configuration.
	Simple param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindRatelimitSimple] `json:"simple" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindRatelimitType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindRatelimit) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The rate limit configuration.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindRatelimitSimple struct {
	// The limit (requests per period).
	Limit param.Field[float64] `json:"limit" api:"required"`
	// The period in seconds.
	Period param.Field[int64] `json:"period" api:"required"`
	// Duration in seconds to apply the mitigation action after the rate limit is
	// exceeded. Valid values are 0 (disabled), 10, or multiples of 60 up to 86400.
	// Must be greater than or equal to the period when non-zero.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindRatelimitSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindRatelimitType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindRatelimitTypeRatelimit DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindRatelimitType = "ratelimit"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindRatelimitType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindRatelimitTypeRatelimit:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType] `json:"type" api:"required"`
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketJurisdiction] `json:"jurisdiction"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2Bucket) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeR2Bucket DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType = "r2_bucket"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketTypeR2Bucket:
		return true
	}
	return false
}

// The
// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
// of the R2 bucket.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketJurisdiction string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketJurisdictionEu          DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketJurisdiction = "eu"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketJurisdictionFedramp     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketJurisdiction = "fedramp"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketJurisdictionFedrampHigh DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketJurisdiction = "fedramp-high"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketJurisdiction) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketJurisdictionEu, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketJurisdictionFedramp, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindR2BucketJurisdictionFedrampHigh:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The secret value to use.
	Text param.Field[string] `json:"text" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretText) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeSecretText DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType = "secret_text"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretTextTypeSecretText:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSendEmail struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSendEmailType] `json:"type" api:"required"`
	// List of allowed destination addresses.
	AllowedDestinationAddresses param.Field[[]string] `json:"allowed_destination_addresses" format:"email"`
	// List of allowed sender addresses.
	AllowedSenderAddresses param.Field[[]string] `json:"allowed_sender_addresses" format:"email"`
	// Destination address for the email.
	DestinationAddress param.Field[string] `json:"destination_address" format:"email"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSendEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSendEmail) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSendEmailType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSendEmailTypeSendEmail DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSendEmailType = "send_email"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSendEmailType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSendEmailTypeSendEmail:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindService struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType] `json:"type" api:"required"`
	// Entrypoint to invoke on the target Worker.
	Entrypoint param.Field[string] `json:"entrypoint"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindService) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeService DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType = "service"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindServiceTypeService:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTextBlob struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The name of the file containing the text content. Only accepted for
	// `service worker syntax` Workers.
	Part param.Field[string] `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTextBlobType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTextBlob) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTextBlob) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTextBlobType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTextBlobTypeTextBlob DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTextBlobType = "text_blob"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTextBlobType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindTextBlobTypeTextBlob:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize struct {
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorize) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeVectorize DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType = "vectorize"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVectorizeTypeVectorize:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadata) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType = "version_metadata"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVersionMetadataTypeVersionMetadata:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Name of the secret in the store.
	SecretName param.Field[string] `json:"secret_name" api:"required"`
	// ID of the store containing the secret.
	StoreID param.Field[string] `json:"store_id" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType = "secrets_store_secret"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretTypeSecretsStoreSecret:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindFlagship struct {
	// ID of the Flagship app to bind to for feature flag evaluation.
	AppID param.Field[string] `json:"app_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindFlagshipType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindFlagship) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindFlagship) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindFlagshipType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindFlagshipTypeFlagship DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindFlagshipType = "flagship"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindFlagshipType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindFlagshipTypeFlagship:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm param.Field[interface{}] `json:"algorithm" api:"required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat] `json:"format" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyType] `json:"type" api:"required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages param.Field[[]DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage] `json:"usages" api:"required"`
	// Base64-encoded key data. Required if `format` is "raw", "pkcs8", or "spki".
	KeyBase64 param.Field[string] `json:"key_base64"`
	// Key data in
	// [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key)
	// format. Required if `format` is "jwk".
	KeyJwk param.Field[interface{}] `json:"key_jwk"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKey) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatRaw   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "raw"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatPkcs8 DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "pkcs8"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatSpki  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "spki"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatJwk   DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat = "jwk"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatRaw, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatPkcs8, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatSpki, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyFormatJwk:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyTypeSecretKey DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyType = "secret_key"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyTypeSecretKey:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageEncrypt    DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "encrypt"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDecrypt    DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "decrypt"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageSign       DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "sign"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageVerify     DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "verify"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveKey  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "deriveKey"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveBits DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "deriveBits"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageWrapKey    DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "wrapKey"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageUnwrapKey  DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage = "unwrapKey"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsage) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageEncrypt, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDecrypt, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageSign, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageVerify, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveKey, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageDeriveBits, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageWrapKey, DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindSecretKeyUsageUnwrapKey:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflow struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflowType] `json:"type" api:"required"`
	// Name of the Workflow to bind to.
	WorkflowName param.Field[string] `json:"workflow_name" api:"required"`
	// Class name of the Workflow. Should only be provided if the Workflow belongs to
	// this script.
	ClassName param.Field[string] `json:"class_name"`
	// Script name that contains the Workflow. If not provided, defaults to this script
	// name.
	ScriptName param.Field[string] `json:"script_name"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflow) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflow) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflowType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflowTypeWorkflow DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflowType = "workflow"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflowType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWorkflowTypeWorkflow:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWasmModule struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The name of the file containing the WebAssembly module content. Only accepted
	// for `service worker syntax` Workers.
	Part param.Field[string] `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWasmModuleType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWasmModule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWasmModule) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWasmModuleType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWasmModuleTypeWasmModule DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWasmModuleType = "wasm_module"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWasmModuleType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindWasmModuleTypeWasmModule:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCService struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Identifier of the VPC service to bind to.
	ServiceID param.Field[string] `json:"service_id" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCServiceType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCService) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCServiceType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCServiceTypeVPCService DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCServiceType = "vpc_service"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCServiceType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCServiceTypeVPCService:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCNetwork struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCNetworkType] `json:"type" api:"required"`
	// Identifier of the network to bind to. Only "cf1:network" is currently supported.
	// Mutually exclusive with tunnel_id.
	NetworkID param.Field[string] `json:"network_id"`
	// UUID of the Cloudflare Tunnel to bind to. Mutually exclusive with network_id.
	TunnelID param.Field[string] `json:"tunnel_id"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCNetwork) implementsDispatchNamespaceScriptUpdateParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCNetworkType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCNetworkTypeVPCNetwork DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCNetworkType = "vpc_network"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCNetworkType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsWorkersBindingKindVPCNetworkTypeVPCNetwork:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAI                     DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "ai"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAISearch               DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "ai_search"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAISearchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "ai_search_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAnalyticsEngine        DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "analytics_engine"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAssets                 DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "assets"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeBrowser                DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "browser"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeD1                     DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "d1"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeDataBlob               DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "data_blob"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeDurableObjectNamespace DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "durable_object_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeHyperdrive             DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "hyperdrive"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeInherit                DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "inherit"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeImages                 DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "images"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeJson                   DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "json"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeKVNamespace            DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "kv_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeMedia                  DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "media"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeMTLSCertificate        DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "mtls_certificate"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypePlainText              DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "plain_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypePipelines              DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "pipelines"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeQueue                  DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "queue"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeRatelimit              DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "ratelimit"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeR2Bucket               DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "r2_bucket"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeSecretText             DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "secret_text"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeSendEmail              DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "send_email"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeService                DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeTextBlob               DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "text_blob"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeVectorize              DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "vectorize"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeVersionMetadata        DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "version_metadata"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeSecretsStoreSecret     DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "secrets_store_secret"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeFlagship               DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "flagship"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeSecretKey              DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "secret_key"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeWorkflow               DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "workflow"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeWasmModule             DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "wasm_module"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeVPCService             DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "vpc_service"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeVPCNetwork             DispatchNamespaceScriptUpdateParamsMetadataBindingsType = "vpc_network"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAI, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAISearch, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAISearchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAnalyticsEngine, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeAssets, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeBrowser, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeD1, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeDataBlob, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeDurableObjectNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeHyperdrive, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeInherit, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeImages, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeJson, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeKVNamespace, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeMedia, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeMTLSCertificate, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypePlainText, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypePipelines, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeQueue, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeRatelimit, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeR2Bucket, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeSecretText, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeSendEmail, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeService, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeTextBlob, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeVectorize, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeVersionMetadata, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeSecretsStoreSecret, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeFlagship, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeSecretKey, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeWorkflow, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeWasmModule, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeVPCService, DispatchNamespaceScriptUpdateParamsMetadataBindingsTypeVPCNetwork:
		return true
	}
	return false
}

// Data format of the key.
// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
type DispatchNamespaceScriptUpdateParamsMetadataBindingsFormat string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsFormatRaw   DispatchNamespaceScriptUpdateParamsMetadataBindingsFormat = "raw"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsFormatPkcs8 DispatchNamespaceScriptUpdateParamsMetadataBindingsFormat = "pkcs8"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsFormatSpki  DispatchNamespaceScriptUpdateParamsMetadataBindingsFormat = "spki"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsFormatJwk   DispatchNamespaceScriptUpdateParamsMetadataBindingsFormat = "jwk"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsFormat) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsFormatRaw, DispatchNamespaceScriptUpdateParamsMetadataBindingsFormatPkcs8, DispatchNamespaceScriptUpdateParamsMetadataBindingsFormatSpki, DispatchNamespaceScriptUpdateParamsMetadataBindingsFormatJwk:
		return true
	}
	return false
}

// The
// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
// of the R2 bucket.
type DispatchNamespaceScriptUpdateParamsMetadataBindingsJurisdiction string

const (
	DispatchNamespaceScriptUpdateParamsMetadataBindingsJurisdictionEu          DispatchNamespaceScriptUpdateParamsMetadataBindingsJurisdiction = "eu"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsJurisdictionFedramp     DispatchNamespaceScriptUpdateParamsMetadataBindingsJurisdiction = "fedramp"
	DispatchNamespaceScriptUpdateParamsMetadataBindingsJurisdictionFedrampHigh DispatchNamespaceScriptUpdateParamsMetadataBindingsJurisdiction = "fedramp-high"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataBindingsJurisdiction) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataBindingsJurisdictionEu, DispatchNamespaceScriptUpdateParamsMetadataBindingsJurisdictionFedramp, DispatchNamespaceScriptUpdateParamsMetadataBindingsJurisdictionFedrampHigh:
		return true
	}
	return false
}

// Global CacheW configuration for the Worker. When caching is on, the platform
// provisions a `cloudflare.app` zone for the Worker. A `type: worker` entry in the
// `exports` map can override this value for a single entrypoint.
type DispatchNamespaceScriptUpdateParamsMetadataCacheOptions struct {
	// Whether caching is enabled for this Worker.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Whether cached responses are shared across Worker version uploads. This is
	// independent of `enabled`. It can stay true while caching is off, so the
	// preference survives turning caching off and back on.
	CrossVersionCache param.Field[bool] `json:"cross_version_cache"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataCacheOptions) MarshalJSON() (data []byte, err error) {
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
type DispatchNamespaceScriptUpdateParamsMetadataExports struct {
	// Marks this entry as a Worker entrypoint export.
	Type  param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsType] `json:"type" api:"required"`
	Cache param.Field[interface{}]                                            `json:"cache"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container param.Field[string] `json:"container"`
	// The destination class name. Must differ from the source class (the map key) and
	// must be declared as a live (`created`) entry in the same `exports` map.
	// Write-only: never present in GET responses.
	RenamedTo param.Field[string] `json:"renamed_to"`
	// Live export. May be omitted; defaults to `created`.
	State param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsState] `json:"state"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsStorage] `json:"storage"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom param.Field[string] `json:"transfer_from"`
	// The destination script name. Must be in the same account and the same
	// dispatch-namespace context (or both non-dispatch). Cross-dispatch-namespace
	// transfers are rejected. Write-only: never present in GET responses.
	TransferredTo param.Field[string] `json:"transferred_to"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExports) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExports) implementsDispatchNamespaceScriptUpdateParamsMetadataExportsUnion() {
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
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersWorkerExport],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExport],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectDeletedExport],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectRenamedExport],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectTransferredExport],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExport],
// [DispatchNamespaceScriptUpdateParamsMetadataExports].
type DispatchNamespaceScriptUpdateParamsMetadataExportsUnion interface {
	implementsDispatchNamespaceScriptUpdateParamsMetadataExportsUnion()
}

// A named Worker entrypoint export (`type: worker`). Worker entrypoints are always
// live (`state: created`) and carry no storage or lifecycle fields. The optional
// `cache` block overrides the Worker's global `cache_options.enabled` for this
// entrypoint.
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersWorkerExport struct {
	// Marks this entry as a Worker entrypoint export.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersWorkerExportType] `json:"type" api:"required"`
	// Cache override for this entrypoint. Overrides the Worker's global
	// `cache_options.enabled` for this entrypoint only.
	Cache param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersWorkerExportCache] `json:"cache"`
	// Live export. May be omitted; defaults to `created`.
	State param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersWorkerExportState] `json:"state"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersWorkerExport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersWorkerExport) implementsDispatchNamespaceScriptUpdateParamsMetadataExportsUnion() {
}

// Marks this entry as a Worker entrypoint export.
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersWorkerExportType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersWorkerExportTypeWorker DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersWorkerExportType = "worker"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersWorkerExportType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersWorkerExportTypeWorker:
		return true
	}
	return false
}

// Cache override for this entrypoint. Overrides the Worker's global
// `cache_options.enabled` for this entrypoint only.
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersWorkerExportCache struct {
	// Whether caching is enabled for this entrypoint.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersWorkerExportCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Live export. May be omitted; defaults to `created`.
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersWorkerExportState string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersWorkerExportStateCreated DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersWorkerExportState = "created"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersWorkerExportState) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersWorkerExportStateCreated:
		return true
	}
	return false
}

// A live Durable Object export (`state: created`, the default). The platform
// auto-provisions the namespace on first deploy, matches it on subsequent deploys,
// and never mutates or deletes it as a side effect of a code-only change.
// `storage` is required; `renamed_to`, `transferred_to` and `transfer_from` are
// not allowed on a live entry.
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExport struct {
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportStorage] `json:"storage" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportType] `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container param.Field[string] `json:"container"`
	// Live export. May be omitted; defaults to `created`.
	State param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportState] `json:"state"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExport) implementsDispatchNamespaceScriptUpdateParamsMetadataExportsUnion() {
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportStorage string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportStorageSqlite   DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportStorage = "sqlite"
	DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportStorageLegacyKV DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportStorage = "legacy-kv"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportStorage) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportStorageSqlite, DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportStorageLegacyKV:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportTypeDurableObject DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportType = "durable-object"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportTypeDurableObject:
		return true
	}
	return false
}

// Live export. May be omitted; defaults to `created`.
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportState string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportStateCreated DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportState = "created"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportState) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExportStateCreated:
		return true
	}
	return false
}

// A `deleted` tombstone: retires the provisioned namespace for this class and all
// of its data. The class must be absent from the uploaded code and no other Worker
// in the account may bind to the namespace, otherwise the deploy is rejected. No
// other fields are allowed. Deletion is irreversible.
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectDeletedExport struct {
	// Tombstone that deletes the namespace.
	State param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectDeletedExportState] `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectDeletedExportType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectDeletedExport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectDeletedExport) implementsDispatchNamespaceScriptUpdateParamsMetadataExportsUnion() {
}

// Tombstone that deletes the namespace.
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectDeletedExportState string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectDeletedExportStateDeleted DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectDeletedExportState = "deleted"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectDeletedExportState) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectDeletedExportStateDeleted:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectDeletedExportType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectDeletedExportTypeDurableObject DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectDeletedExportType = "durable-object"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectDeletedExportType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectDeletedExportTypeDurableObject:
		return true
	}
	return false
}

// A `renamed` tombstone: rewrites the provisioned namespace's class name from this
// map key to `renamed_to`. The source class may stay in code during the rollout
// window (an info notice is emitted). `storage`, `transferred_to` and
// `transfer_from` are not allowed.
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectRenamedExport struct {
	// The destination class name. Must differ from the source class (the map key) and
	// must be declared as a live (`created`) entry in the same `exports` map.
	// Write-only: never present in GET responses.
	RenamedTo param.Field[string] `json:"renamed_to" api:"required"`
	// Tombstone that renames the namespace's class.
	State param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectRenamedExportState] `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectRenamedExportType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectRenamedExport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectRenamedExport) implementsDispatchNamespaceScriptUpdateParamsMetadataExportsUnion() {
}

// Tombstone that renames the namespace's class.
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectRenamedExportState string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectRenamedExportStateRenamed DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectRenamedExportState = "renamed"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectRenamedExportState) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectRenamedExportStateRenamed:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectRenamedExportType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectRenamedExportTypeDurableObject DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectRenamedExportType = "durable-object"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectRenamedExportType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectRenamedExportTypeDurableObject:
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
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectTransferredExport struct {
	// Tombstone that transfers the namespace to another script.
	State param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectTransferredExportState] `json:"state" api:"required"`
	// The destination script name. Must be in the same account and the same
	// dispatch-namespace context (or both non-dispatch). Cross-dispatch-namespace
	// transfers are rejected. Write-only: never present in GET responses.
	TransferredTo param.Field[string] `json:"transferred_to" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectTransferredExportType] `json:"type" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectTransferredExport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectTransferredExport) implementsDispatchNamespaceScriptUpdateParamsMetadataExportsUnion() {
}

// Tombstone that transfers the namespace to another script.
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectTransferredExportState string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectTransferredExportStateTransferred DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectTransferredExportState = "transferred"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectTransferredExportState) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectTransferredExportStateTransferred:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectTransferredExportType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectTransferredExportTypeDurableObject DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectTransferredExportType = "durable-object"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectTransferredExportType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectTransferredExportTypeDurableObject:
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
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExport struct {
	// Target side of a two-phase transfer.
	State param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportState] `json:"state" api:"required"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStorage] `json:"storage" api:"required"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom param.Field[string] `json:"transfer_from" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportType] `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object once the transfer settles. Valid only on live entries.
	Container param.Field[string] `json:"container"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExport) implementsDispatchNamespaceScriptUpdateParamsMetadataExportsUnion() {
}

// Target side of a two-phase transfer.
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportState string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStateExpectingTransfer DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportState = "expecting-transfer"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportState) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStateExpectingTransfer:
		return true
	}
	return false
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStorage string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStorageSqlite   DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStorage = "sqlite"
	DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStorageLegacyKV DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStorage = "legacy-kv"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStorage) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStorageSqlite, DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStorageLegacyKV:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportTypeDurableObject DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportType = "durable-object"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsWorkersDurableObjectExpectingTransferExportTypeDurableObject:
		return true
	}
	return false
}

// Marks this entry as a Worker entrypoint export.
type DispatchNamespaceScriptUpdateParamsMetadataExportsType string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsTypeWorker        DispatchNamespaceScriptUpdateParamsMetadataExportsType = "worker"
	DispatchNamespaceScriptUpdateParamsMetadataExportsTypeDurableObject DispatchNamespaceScriptUpdateParamsMetadataExportsType = "durable-object"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsTypeWorker, DispatchNamespaceScriptUpdateParamsMetadataExportsTypeDurableObject:
		return true
	}
	return false
}

// Live export. May be omitted; defaults to `created`.
type DispatchNamespaceScriptUpdateParamsMetadataExportsState string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsStateCreated           DispatchNamespaceScriptUpdateParamsMetadataExportsState = "created"
	DispatchNamespaceScriptUpdateParamsMetadataExportsStateDeleted           DispatchNamespaceScriptUpdateParamsMetadataExportsState = "deleted"
	DispatchNamespaceScriptUpdateParamsMetadataExportsStateRenamed           DispatchNamespaceScriptUpdateParamsMetadataExportsState = "renamed"
	DispatchNamespaceScriptUpdateParamsMetadataExportsStateTransferred       DispatchNamespaceScriptUpdateParamsMetadataExportsState = "transferred"
	DispatchNamespaceScriptUpdateParamsMetadataExportsStateExpectingTransfer DispatchNamespaceScriptUpdateParamsMetadataExportsState = "expecting-transfer"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsState) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsStateCreated, DispatchNamespaceScriptUpdateParamsMetadataExportsStateDeleted, DispatchNamespaceScriptUpdateParamsMetadataExportsStateRenamed, DispatchNamespaceScriptUpdateParamsMetadataExportsStateTransferred, DispatchNamespaceScriptUpdateParamsMetadataExportsStateExpectingTransfer:
		return true
	}
	return false
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type DispatchNamespaceScriptUpdateParamsMetadataExportsStorage string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsStorageSqlite   DispatchNamespaceScriptUpdateParamsMetadataExportsStorage = "sqlite"
	DispatchNamespaceScriptUpdateParamsMetadataExportsStorageLegacyKV DispatchNamespaceScriptUpdateParamsMetadataExportsStorage = "legacy-kv"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsStorage) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsStorageSqlite, DispatchNamespaceScriptUpdateParamsMetadataExportsStorageLegacyKV:
		return true
	}
	return false
}

// Summary of the declarative exports reconciliation that ran on this upload.
// Populated only when the uploaded metadata included an `exports` block. Durable
// Object entries drive reconciliation; `type: worker` entries do not contribute to
// this summary.
type DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliation struct {
	// Class names for which a new namespace was provisioned.
	Created param.Field[[]string] `json:"created" api:"required"`
	// Class names whose namespace was deleted by a `deleted` tombstone.
	Deleted param.Field[[]string] `json:"deleted" api:"required"`
	// Non-blocking info entries (stale tombstones, tombstone applied with class still
	// in code). See `exports_reconciliation_info`.
	Info param.Field[[]DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfo] `json:"info" api:"required"`
	// Source class names whose tombstone entry is now stale and safe to delete from
	// `exports` (no remaining referencing scripts).
	RemovableEntries param.Field[[]string] `json:"removable_entries" api:"required"`
	// Applied `renamed` tombstones.
	Renamed param.Field[[]DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationRenamed] `json:"renamed" api:"required"`
	// Phase-1 transfer hints recorded on the target side.
	TransferPending param.Field[[]DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationTransferPending] `json:"transfer_pending" api:"required"`
	// Committed `transferred` tombstones (phase-2).
	Transferred param.Field[[]DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationTransferred] `json:"transferred" api:"required"`
	// Class names whose provisioned namespace was mutated in place.
	Updated param.Field[[]string] `json:"updated" api:"required"`
	// Non-blocking warnings. See `exports_reconciliation_warning`.
	Warnings param.Field[[]DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarning] `json:"warnings" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A non-blocking reconciliation info entry. Emitted for stale tombstones (a no-op
// on this deploy) and for tombstones applied with the source class still in code
// (the supported zero-downtime rollout pattern).
type DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfo struct {
	// The class name the info entry is about.
	Class param.Field[string] `json:"class" api:"required"`
	// Human-readable explanation.
	Message param.Field[string] `json:"message" api:"required"`
	// Stable, machine-readable tag identifying which reconciliation scenario produced
	// an error, warning, or info entry. Clients may branch on this value instead of
	// parsing `message`.
	Scenario param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario] `json:"scenario" api:"required"`
	// The provisioned namespace the entry relates to, when applicable.
	NamespaceID param.Field[string] `json:"namespace_id" format:"uuid"`
	// Other Workers in the account that still bind to the affected class. Advisory:
	// while non-empty the tombstone is not yet safe to remove — redeploy these Workers
	// with bindings re-pointed first.
	ReferencingScripts param.Field[[]string] `json:"referencing_scripts"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Stable, machine-readable tag identifying which reconciliation scenario produced
// an error, warning, or info entry. Clients may branch on this value instead of
// parsing `message`.
type DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioCodeClassNotInExports                     DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "code_class_not_in_exports"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioProvisionedClassMissingFromConfig         DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "provisioned_class_missing_from_config"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioConfigExportNotInCode                     DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "config_export_not_in_code"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioConfigReferencesNonexistentClass          DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "config_references_nonexistent_class"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioOrphanedProvisionedNamespace              DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "orphaned_provisioned_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioStorageTypeMismatch                       DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "storage_type_mismatch"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioFreeTierRequiresSqlite                    DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "free_tier_requires_sqlite"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioInvalidExport                             DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "invalid_export"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTombstoneDeleteClassStillInCode           DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "tombstone_delete_class_still_in_code"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTombstoneDeleteBlockedByExternalBindings  DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "tombstone_delete_blocked_by_external_bindings"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTombstoneRenamedToOccupied                DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "tombstone_renamed_to_occupied"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTransferredPendingNotFound                DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "transferred_pending_not_found"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTransferredTargetMissing                  DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "transferred_target_missing"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTransferredTargetMismatch                 DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "transferred_target_mismatch"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioPhaseOneTransferSourceMissing             DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "phase_one_transfer_source_missing"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioPhaseOneTransferSourceNamespaceMissing    DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "phase_one_transfer_source_namespace_missing"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioPhaseOneTransferTargetClassProvisioned    DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "phase_one_transfer_target_class_provisioned"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioPhaseOneTransferAfterCommitMismatch       DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "phase_one_transfer_after_commit_mismatch"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioPhaseOneTransferDuplicate                 DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "phase_one_transfer_duplicate"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioPhaseOneTransferTargetInDispatchNamespace DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "phase_one_transfer_target_in_dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioPhaseOneTransferSourceInDispatchNamespace DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "phase_one_transfer_source_in_dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTransferredSourceInDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "transferred_source_in_dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTransferredTargetInDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "transferred_target_in_dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioContainerUndeclaredReference              DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "container_undeclared_reference"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioContainerClassNotDurableObject            DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "container_class_not_durable_object"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioContainerWiringInconsistent               DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "container_wiring_inconsistent"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioContainerMultipleDurableObjects           DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "container_multiple_durable_objects"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTransferContainerParityMismatch           DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "transfer_container_parity_mismatch"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTransferContainerParityMismatchOnCommit   DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "transfer_container_parity_mismatch_on_commit"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTombstoneClassStillInCode                 DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "tombstone_class_still_in_code"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioStaleTombstone                            DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "stale_tombstone"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTransferReceiveAlreadyApplied             DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "transfer_receive_already_applied"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTransferReceiveCleanupComplete            DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario = "transfer_receive_cleanup_complete"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenario) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioCodeClassNotInExports, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioProvisionedClassMissingFromConfig, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioConfigExportNotInCode, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioConfigReferencesNonexistentClass, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioOrphanedProvisionedNamespace, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioStorageTypeMismatch, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioFreeTierRequiresSqlite, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioInvalidExport, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTombstoneDeleteClassStillInCode, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTombstoneDeleteBlockedByExternalBindings, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTombstoneRenamedToOccupied, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTransferredPendingNotFound, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTransferredTargetMissing, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTransferredTargetMismatch, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioPhaseOneTransferSourceMissing, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioPhaseOneTransferSourceNamespaceMissing, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioPhaseOneTransferTargetClassProvisioned, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioPhaseOneTransferAfterCommitMismatch, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioPhaseOneTransferDuplicate, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioPhaseOneTransferTargetInDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioPhaseOneTransferSourceInDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTransferredSourceInDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTransferredTargetInDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioContainerUndeclaredReference, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioContainerClassNotDurableObject, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioContainerWiringInconsistent, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioContainerMultipleDurableObjects, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTransferContainerParityMismatch, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTransferContainerParityMismatchOnCommit, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTombstoneClassStillInCode, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioStaleTombstone, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTransferReceiveAlreadyApplied, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationInfoScenarioTransferReceiveCleanupComplete:
		return true
	}
	return false
}

// A single applied `renamed` tombstone.
type DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationRenamed struct {
	// The original (source) class name.
	From param.Field[string] `json:"from" api:"required"`
	// The new class name (`renamed_to`).
	To param.Field[string] `json:"to" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationRenamed) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A single phase-1 transfer hint recorded on the target side (a live
// `expecting-transfer` entry).
type DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationTransferPending struct {
	// The target-side class name awaiting transfer.
	Class param.Field[string] `json:"class" api:"required"`
	// The source script the namespace will be transferred from.
	From param.Field[string] `json:"from" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationTransferPending) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A single committed `transferred` tombstone (phase-2 commit).
type DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationTransferred struct {
	// The source class name that was transferred.
	Class param.Field[string] `json:"class" api:"required"`
	// The transfer phase. Currently always `committed`.
	Phase param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationTransferredPhase] `json:"phase" api:"required"`
	// The destination script that now owns the namespace.
	To param.Field[string] `json:"to" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationTransferred) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The transfer phase. Currently always `committed`.
type DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationTransferredPhase string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationTransferredPhaseCommitted DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationTransferredPhase = "committed"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationTransferredPhase) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationTransferredPhaseCommitted:
		return true
	}
	return false
}

// A non-blocking reconciliation warning. Reserved: no scenario populates this
// array today (`code_class_not_in_exports` is surfaced as info and
// `provisioned_class_missing_from_config` is a hard error). Clients should still
// surface any entries that appear.
type DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarning struct {
	// The class name the warning is about.
	Class param.Field[string] `json:"class" api:"required"`
	// Human-readable explanation of the warning.
	Message param.Field[string] `json:"message" api:"required"`
	// Stable, machine-readable tag identifying which reconciliation scenario produced
	// an error, warning, or info entry. Clients may branch on this value instead of
	// parsing `message`.
	Scenario param.Field[DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario] `json:"scenario" api:"required"`
	// The provisioned namespace the warning relates to, when applicable.
	NamespaceID param.Field[string] `json:"namespace_id" format:"uuid"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Stable, machine-readable tag identifying which reconciliation scenario produced
// an error, warning, or info entry. Clients may branch on this value instead of
// parsing `message`.
type DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario string

const (
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioCodeClassNotInExports                     DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "code_class_not_in_exports"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioProvisionedClassMissingFromConfig         DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "provisioned_class_missing_from_config"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioConfigExportNotInCode                     DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "config_export_not_in_code"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioConfigReferencesNonexistentClass          DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "config_references_nonexistent_class"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioOrphanedProvisionedNamespace              DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "orphaned_provisioned_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioStorageTypeMismatch                       DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "storage_type_mismatch"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioFreeTierRequiresSqlite                    DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "free_tier_requires_sqlite"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioInvalidExport                             DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "invalid_export"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTombstoneDeleteClassStillInCode           DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "tombstone_delete_class_still_in_code"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTombstoneDeleteBlockedByExternalBindings  DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "tombstone_delete_blocked_by_external_bindings"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTombstoneRenamedToOccupied                DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "tombstone_renamed_to_occupied"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTransferredPendingNotFound                DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "transferred_pending_not_found"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTransferredTargetMissing                  DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "transferred_target_missing"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTransferredTargetMismatch                 DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "transferred_target_mismatch"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioPhaseOneTransferSourceMissing             DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "phase_one_transfer_source_missing"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioPhaseOneTransferSourceNamespaceMissing    DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "phase_one_transfer_source_namespace_missing"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioPhaseOneTransferTargetClassProvisioned    DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "phase_one_transfer_target_class_provisioned"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioPhaseOneTransferAfterCommitMismatch       DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "phase_one_transfer_after_commit_mismatch"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioPhaseOneTransferDuplicate                 DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "phase_one_transfer_duplicate"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioPhaseOneTransferTargetInDispatchNamespace DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "phase_one_transfer_target_in_dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioPhaseOneTransferSourceInDispatchNamespace DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "phase_one_transfer_source_in_dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTransferredSourceInDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "transferred_source_in_dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTransferredTargetInDispatchNamespace      DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "transferred_target_in_dispatch_namespace"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioContainerUndeclaredReference              DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "container_undeclared_reference"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioContainerClassNotDurableObject            DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "container_class_not_durable_object"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioContainerWiringInconsistent               DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "container_wiring_inconsistent"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioContainerMultipleDurableObjects           DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "container_multiple_durable_objects"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTransferContainerParityMismatch           DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "transfer_container_parity_mismatch"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTransferContainerParityMismatchOnCommit   DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "transfer_container_parity_mismatch_on_commit"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTombstoneClassStillInCode                 DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "tombstone_class_still_in_code"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioStaleTombstone                            DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "stale_tombstone"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTransferReceiveAlreadyApplied             DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "transfer_receive_already_applied"
	DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTransferReceiveCleanupComplete            DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario = "transfer_receive_cleanup_complete"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenario) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioCodeClassNotInExports, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioProvisionedClassMissingFromConfig, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioConfigExportNotInCode, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioConfigReferencesNonexistentClass, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioOrphanedProvisionedNamespace, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioStorageTypeMismatch, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioFreeTierRequiresSqlite, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioInvalidExport, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTombstoneDeleteClassStillInCode, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTombstoneDeleteBlockedByExternalBindings, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTombstoneRenamedToOccupied, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTransferredPendingNotFound, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTransferredTargetMissing, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTransferredTargetMismatch, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioPhaseOneTransferSourceMissing, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioPhaseOneTransferSourceNamespaceMissing, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioPhaseOneTransferTargetClassProvisioned, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioPhaseOneTransferAfterCommitMismatch, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioPhaseOneTransferDuplicate, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioPhaseOneTransferTargetInDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioPhaseOneTransferSourceInDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTransferredSourceInDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTransferredTargetInDispatchNamespace, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioContainerUndeclaredReference, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioContainerClassNotDurableObject, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioContainerWiringInconsistent, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioContainerMultipleDurableObjects, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTransferContainerParityMismatch, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTransferContainerParityMismatchOnCommit, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTombstoneClassStillInCode, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioStaleTombstone, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTransferReceiveAlreadyApplied, DispatchNamespaceScriptUpdateParamsMetadataExportsReconciliationWarningsScenarioTransferReceiveCleanupComplete:
		return true
	}
	return false
}

// Limits to apply for this Worker.
type DispatchNamespaceScriptUpdateParamsMetadataLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
	CPUMs param.Field[int64] `json:"cpu_ms"`
	// The number of subrequests this Worker can make per request.
	Subrequests param.Field[int64] `json:"subrequests"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataLimits) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Migrations to apply for Durable Objects associated with this Worker.
type DispatchNamespaceScriptUpdateParamsMetadataMigrations struct {
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

func (r DispatchNamespaceScriptUpdateParamsMetadataMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataMigrations) ImplementsDispatchNamespaceScriptUpdateParamsMetadataMigrationsUnion() {
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by [workers.SingleStepMigrationParam],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations],
// [DispatchNamespaceScriptUpdateParamsMetadataMigrations].
type DispatchNamespaceScriptUpdateParamsMetadataMigrationsUnion interface {
	ImplementsDispatchNamespaceScriptUpdateParamsMetadataMigrationsUnion()
}

type DispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]workers.MigrationStepParam] `json:"steps"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersMultipleStepMigrations) ImplementsDispatchNamespaceScriptUpdateParamsMetadataMigrationsUnion() {
}

// Observability settings for the Worker.
type DispatchNamespaceScriptUpdateParamsMetadataObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
	// Log settings for the Worker.
	Logs param.Field[DispatchNamespaceScriptUpdateParamsMetadataObservabilityLogs] `json:"logs"`
	// Trace settings for the Worker.
	Traces param.Field[DispatchNamespaceScriptUpdateParamsMetadataObservabilityTraces] `json:"traces"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObservability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Log settings for the Worker.
type DispatchNamespaceScriptUpdateParamsMetadataObservabilityLogs struct {
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

func (r DispatchNamespaceScriptUpdateParamsMetadataObservabilityLogs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Trace settings for the Worker.
type DispatchNamespaceScriptUpdateParamsMetadataObservabilityTraces struct {
	// A list of destinations where traces will be exported to.
	Destinations param.Field[[]string] `json:"destinations"`
	// Whether traces are enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled"`
	// The sampling rate for traces. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
	// Whether trace persistence is enabled for the Worker.
	Persist param.Field[bool] `json:"persist"`
	// Controls how inbound trace context (traceparent/tracestate) headers on incoming
	// requests are handled. "authenticated" (default) honors inbound trace context
	// only when accompanied by a valid trace auth token. "accept" unconditionally
	// accepts inbound trace context. Requires the trace propagation feature to be
	// enabled.
	PropagationPolicy param.Field[DispatchNamespaceScriptUpdateParamsMetadataObservabilityTracesPropagationPolicy] `json:"propagation_policy"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataObservabilityTraces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls how inbound trace context (traceparent/tracestate) headers on incoming
// requests are handled. "authenticated" (default) honors inbound trace context
// only when accompanied by a valid trace auth token. "accept" unconditionally
// accepts inbound trace context. Requires the trace propagation feature to be
// enabled.
type DispatchNamespaceScriptUpdateParamsMetadataObservabilityTracesPropagationPolicy string

const (
	DispatchNamespaceScriptUpdateParamsMetadataObservabilityTracesPropagationPolicyAuthenticated DispatchNamespaceScriptUpdateParamsMetadataObservabilityTracesPropagationPolicy = "authenticated"
	DispatchNamespaceScriptUpdateParamsMetadataObservabilityTracesPropagationPolicyAccept        DispatchNamespaceScriptUpdateParamsMetadataObservabilityTracesPropagationPolicy = "accept"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataObservabilityTracesPropagationPolicy) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataObservabilityTracesPropagationPolicyAuthenticated, DispatchNamespaceScriptUpdateParamsMetadataObservabilityTracesPropagationPolicyAccept:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsMetadataPackageDependency struct {
	// The exact version that was resolved and installed by the package manager.
	InstalledVersion param.Field[string] `json:"installedVersion" api:"required"`
	// The npm package name.
	Name param.Field[string] `json:"name" api:"required"`
	// The version constraint as written in package.json.
	PackageJsonVersion param.Field[string] `json:"packageJsonVersion" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataPackageDependency) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
type DispatchNamespaceScriptUpdateParamsMetadataPlacement struct {
	// TCP host and port for targeted placement.
	Host param.Field[string] `json:"host"`
	// HTTP hostname for targeted placement.
	Hostname param.Field[string] `json:"hostname"`
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode param.Field[DispatchNamespaceScriptUpdateParamsMetadataPlacementMode] `json:"mode"`
	// Cloud region for targeted placement in format 'provider:region'.
	Region param.Field[string]      `json:"region"`
	Target param.Field[interface{}] `json:"target"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataPlacement) implementsDispatchNamespaceScriptUpdateParamsMetadataPlacementUnion() {
}

// Configuration for
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Specify mode='smart' for Smart Placement, or one of region/hostname/host.
//
// Satisfied by
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataPlacementObject],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataPlacementObject],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataPlacementObject],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataPlacementObject],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataPlacementObject],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataPlacementObject],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataPlacementObject],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsMetadataPlacementObject],
// [DispatchNamespaceScriptUpdateParamsMetadataPlacement].
type DispatchNamespaceScriptUpdateParamsMetadataPlacementUnion interface {
	implementsDispatchNamespaceScriptUpdateParamsMetadataPlacementUnion()
}

type DispatchNamespaceScriptUpdateParamsMetadataPlacementObject struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	Mode param.Field[DispatchNamespaceScriptUpdateParamsMetadataPlacementObjectMode] `json:"mode" api:"required"`
}

func (r DispatchNamespaceScriptUpdateParamsMetadataPlacementObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsMetadataPlacementObject) implementsDispatchNamespaceScriptUpdateParamsMetadataPlacementUnion() {
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateParamsMetadataPlacementObjectMode string

const (
	DispatchNamespaceScriptUpdateParamsMetadataPlacementObjectModeSmart DispatchNamespaceScriptUpdateParamsMetadataPlacementObjectMode = "smart"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataPlacementObjectMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataPlacementObjectModeSmart:
		return true
	}
	return false
}

// Status of
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateParamsMetadataPlacementObjectStatus string

const (
	DispatchNamespaceScriptUpdateParamsMetadataPlacementObjectStatusSuccess                 DispatchNamespaceScriptUpdateParamsMetadataPlacementObjectStatus = "SUCCESS"
	DispatchNamespaceScriptUpdateParamsMetadataPlacementObjectStatusUnsupportedApplication  DispatchNamespaceScriptUpdateParamsMetadataPlacementObjectStatus = "UNSUPPORTED_APPLICATION"
	DispatchNamespaceScriptUpdateParamsMetadataPlacementObjectStatusInsufficientInvocations DispatchNamespaceScriptUpdateParamsMetadataPlacementObjectStatus = "INSUFFICIENT_INVOCATIONS"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataPlacementObjectStatus) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataPlacementObjectStatusSuccess, DispatchNamespaceScriptUpdateParamsMetadataPlacementObjectStatusUnsupportedApplication, DispatchNamespaceScriptUpdateParamsMetadataPlacementObjectStatusInsufficientInvocations:
		return true
	}
	return false
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateParamsMetadataPlacementMode string

const (
	DispatchNamespaceScriptUpdateParamsMetadataPlacementModeSmart    DispatchNamespaceScriptUpdateParamsMetadataPlacementMode = "smart"
	DispatchNamespaceScriptUpdateParamsMetadataPlacementModeTargeted DispatchNamespaceScriptUpdateParamsMetadataPlacementMode = "targeted"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataPlacementMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataPlacementModeSmart, DispatchNamespaceScriptUpdateParamsMetadataPlacementModeTargeted:
		return true
	}
	return false
}

// Status of
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
type DispatchNamespaceScriptUpdateParamsMetadataPlacementStatus string

const (
	DispatchNamespaceScriptUpdateParamsMetadataPlacementStatusSuccess                 DispatchNamespaceScriptUpdateParamsMetadataPlacementStatus = "SUCCESS"
	DispatchNamespaceScriptUpdateParamsMetadataPlacementStatusUnsupportedApplication  DispatchNamespaceScriptUpdateParamsMetadataPlacementStatus = "UNSUPPORTED_APPLICATION"
	DispatchNamespaceScriptUpdateParamsMetadataPlacementStatusInsufficientInvocations DispatchNamespaceScriptUpdateParamsMetadataPlacementStatus = "INSUFFICIENT_INVOCATIONS"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataPlacementStatus) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataPlacementStatusSuccess, DispatchNamespaceScriptUpdateParamsMetadataPlacementStatusUnsupportedApplication, DispatchNamespaceScriptUpdateParamsMetadataPlacementStatusInsufficientInvocations:
		return true
	}
	return false
}

// Usage model for the Worker invocations.
type DispatchNamespaceScriptUpdateParamsMetadataUsageModel string

const (
	DispatchNamespaceScriptUpdateParamsMetadataUsageModelStandard DispatchNamespaceScriptUpdateParamsMetadataUsageModel = "standard"
	DispatchNamespaceScriptUpdateParamsMetadataUsageModelBundled  DispatchNamespaceScriptUpdateParamsMetadataUsageModel = "bundled"
	DispatchNamespaceScriptUpdateParamsMetadataUsageModelUnbound  DispatchNamespaceScriptUpdateParamsMetadataUsageModel = "unbound"
)

func (r DispatchNamespaceScriptUpdateParamsMetadataUsageModel) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsMetadataUsageModelStandard, DispatchNamespaceScriptUpdateParamsMetadataUsageModelBundled, DispatchNamespaceScriptUpdateParamsMetadataUsageModelUnbound:
		return true
	}
	return false
}

// When set to "strict", the upload will fail if any `inherit` type bindings cannot
// be resolved against the previous version of the script. Without this,
// unresolvable inherit bindings are silently dropped.
type DispatchNamespaceScriptUpdateParamsBindingsInherit string

const (
	DispatchNamespaceScriptUpdateParamsBindingsInheritStrict DispatchNamespaceScriptUpdateParamsBindingsInherit = "strict"
)

func (r DispatchNamespaceScriptUpdateParamsBindingsInherit) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsBindingsInheritStrict:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateResponseEnvelope struct {
	Errors   []DispatchNamespaceScriptUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DispatchNamespaceScriptUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   DispatchNamespaceScriptUpdateResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success DispatchNamespaceScriptUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    dispatchNamespaceScriptUpdateResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptUpdateResponseEnvelope]
type dispatchNamespaceScriptUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptUpdateResponseEnvelopeErrors struct {
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           DispatchNamespaceScriptUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dispatchNamespaceScriptUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptUpdateResponseEnvelopeErrors]
type dispatchNamespaceScriptUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    dispatchNamespaceScriptUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptUpdateResponseEnvelopeErrorsSource]
type dispatchNamespaceScriptUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptUpdateResponseEnvelopeMessages struct {
	Code             int64                                                       `json:"code" api:"required"`
	Message          string                                                      `json:"message" api:"required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           DispatchNamespaceScriptUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dispatchNamespaceScriptUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptUpdateResponseEnvelopeMessages]
type dispatchNamespaceScriptUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    dispatchNamespaceScriptUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptUpdateResponseEnvelopeMessagesSource]
type dispatchNamespaceScriptUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DispatchNamespaceScriptUpdateResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptUpdateResponseEnvelopeSuccessTrue DispatchNamespaceScriptUpdateResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// If set to true, delete will not be stopped by associated service binding,
	// durable object, or other binding. Any of these associated bindings/durable
	// objects will be deleted along with the script.
	Force param.Field[bool] `query:"force"`
}

// URLQuery serializes [DispatchNamespaceScriptDeleteParams]'s query parameters as
// `url.Values`.
func (r DispatchNamespaceScriptDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DispatchNamespaceScriptDeleteResponseEnvelope struct {
	Errors   []DispatchNamespaceScriptDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DispatchNamespaceScriptDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DispatchNamespaceScriptDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DispatchNamespaceScriptDeleteResponse                `json:"result" api:"nullable"`
	JSON    dispatchNamespaceScriptDeleteResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptDeleteResponseEnvelope]
type dispatchNamespaceScriptDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptDeleteResponseEnvelopeErrors struct {
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           DispatchNamespaceScriptDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dispatchNamespaceScriptDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dispatchNamespaceScriptDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptDeleteResponseEnvelopeErrors]
type dispatchNamespaceScriptDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    dispatchNamespaceScriptDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceScriptDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptDeleteResponseEnvelopeErrorsSource]
type dispatchNamespaceScriptDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptDeleteResponseEnvelopeMessages struct {
	Code             int64                                                       `json:"code" api:"required"`
	Message          string                                                      `json:"message" api:"required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           DispatchNamespaceScriptDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dispatchNamespaceScriptDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dispatchNamespaceScriptDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptDeleteResponseEnvelopeMessages]
type dispatchNamespaceScriptDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    dispatchNamespaceScriptDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceScriptDeleteResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptDeleteResponseEnvelopeMessagesSource]
type dispatchNamespaceScriptDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DispatchNamespaceScriptDeleteResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptDeleteResponseEnvelopeSuccessTrue DispatchNamespaceScriptDeleteResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DispatchNamespaceScriptGetResponseEnvelope struct {
	Errors   []DispatchNamespaceScriptGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DispatchNamespaceScriptGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Details about a worker uploaded to a Workers for Platforms namespace.
	Result Script `json:"result" api:"required"`
	// Whether the API call was successful.
	Success DispatchNamespaceScriptGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    dispatchNamespaceScriptGetResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptGetResponseEnvelope]
type dispatchNamespaceScriptGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptGetResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           DispatchNamespaceScriptGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dispatchNamespaceScriptGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dispatchNamespaceScriptGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptGetResponseEnvelopeErrors]
type dispatchNamespaceScriptGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    dispatchNamespaceScriptGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceScriptGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptGetResponseEnvelopeErrorsSource]
type dispatchNamespaceScriptGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptGetResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           DispatchNamespaceScriptGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dispatchNamespaceScriptGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dispatchNamespaceScriptGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptGetResponseEnvelopeMessages]
type dispatchNamespaceScriptGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    dispatchNamespaceScriptGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceScriptGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptGetResponseEnvelopeMessagesSource]
type dispatchNamespaceScriptGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DispatchNamespaceScriptGetResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptGetResponseEnvelopeSuccessTrue DispatchNamespaceScriptGetResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
