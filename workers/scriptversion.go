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

	"github.com/cloudflare/cloudflare-go/v7/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
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
		return nil, err
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/versions", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List of Worker Versions. The first version in the list is the latest version.
func (r *ScriptVersionService) List(ctx context.Context, scriptName string, params ScriptVersionListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[ScriptVersionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return nil, err
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

// Retrieves detailed information about a specific version of a Workers script.
func (r *ScriptVersionService) Get(ctx context.Context, scriptName string, versionID string, query ScriptVersionGetParams, opts ...option.RequestOption) (res *ScriptVersionGetResponse, err error) {
	var env ScriptVersionGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return nil, err
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/versions/%s", query.AccountID, scriptName, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ScriptVersionNewResponse struct {
	Resources ScriptVersionNewResponseResources `json:"resources" api:"required"`
	// Unique identifier for the version.
	ID string `json:"id"`
	// Summary of the declarative exports reconciliation that ran on this upload.
	// Populated only when the uploaded metadata included an `exports` block. Durable
	// Object entries drive reconciliation; `type: worker` entries do not contribute to
	// this summary.
	ExportsReconciliation ScriptVersionNewResponseExportsReconciliation `json:"exports_reconciliation"`
	Metadata              ScriptVersionNewResponseMetadata              `json:"metadata"`
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
	Resources             apijson.Field
	ID                    apijson.Field
	ExportsReconciliation apijson.Field
	Metadata              apijson.Field
	Number                apijson.Field
	StartupTimeMs         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
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
	Bindings ScriptVersionNewResponseResourcesBindings `json:"bindings"`
	Script   ScriptVersionNewResponseResourcesScript   `json:"script"`
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

// List of bindings attached to a Worker. You can find more about bindings on our
// docs:
// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
type ScriptVersionNewResponseResourcesBindings struct {
	JSON scriptVersionNewResponseResourcesBindingsJSON `json:"-"`
}

// scriptVersionNewResponseResourcesBindingsJSON contains the JSON metadata for the
// struct [ScriptVersionNewResponseResourcesBindings]
type scriptVersionNewResponseResourcesBindingsJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesBindingsJSON) RawJSON() string {
	return r.raw
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
	// Declarative exports for this version, including Durable Object classes (with
	// their `storage` backend) and named Worker entrypoints. Tombstoned lifecycle
	// entries are omitted, so only live exports (`created` and `expecting-transfer`)
	// are returned.
	Exports map[string]ScriptVersionNewResponseResourcesScriptRuntimeExport `json:"exports"`
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
	Exports            apijson.Field
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

// A single entry in the `exports` map, keyed by export name (a `WorkerEntrypoint`
// class name, a Durable Object class name, or `default` for the Worker's default
// export). The `type` discriminator selects the top-level shape: `worker`
// entrypoint entries may carry `cache` configuration, while `durable-object`
// entries are further refined by the optional `state` field (default `created`).
// Tombstone states (`deleted`, `renamed`, `transferred`) express destructive
// lifecycle operations declaratively; `expecting-transfer` is the live target side
// of a transfer. The server validates the exact per-(type, state) field
// combinations; fields not listed for a variant are rejected.
type ScriptVersionNewResponseResourcesScriptRuntimeExport struct {
	// Marks this entry as a Worker entrypoint export.
	Type ScriptVersionNewResponseResourcesScriptRuntimeExportsType `json:"type" api:"required"`
	// This field can have the runtime type of
	// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportCache].
	Cache interface{} `json:"cache"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container string `json:"container"`
	// Live export. May be omitted; defaults to `created`.
	State ScriptVersionNewResponseResourcesScriptRuntimeExportsState `json:"state"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage ScriptVersionNewResponseResourcesScriptRuntimeExportsStorage `json:"storage"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom string                                                   `json:"transfer_from"`
	JSON         scriptVersionNewResponseResourcesScriptRuntimeExportJSON `json:"-"`
	union        ScriptVersionNewResponseResourcesScriptRuntimeExportsUnion
}

// scriptVersionNewResponseResourcesScriptRuntimeExportJSON contains the JSON
// metadata for the struct [ScriptVersionNewResponseResourcesScriptRuntimeExport]
type scriptVersionNewResponseResourcesScriptRuntimeExportJSON struct {
	Type         apijson.Field
	Cache        apijson.Field
	Container    apijson.Field
	State        apijson.Field
	Storage      apijson.Field
	TransferFrom apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r scriptVersionNewResponseResourcesScriptRuntimeExportJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptVersionNewResponseResourcesScriptRuntimeExport) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptVersionNewResponseResourcesScriptRuntimeExport{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptVersionNewResponseResourcesScriptRuntimeExportsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExport],
// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExport],
// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExport],
// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExport],
// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExport],
// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExport].
func (r ScriptVersionNewResponseResourcesScriptRuntimeExport) AsUnion() ScriptVersionNewResponseResourcesScriptRuntimeExportsUnion {
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
// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExport],
// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExport],
// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExport],
// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExport],
// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExport]
// or
// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExport].
type ScriptVersionNewResponseResourcesScriptRuntimeExportsUnion interface {
	implementsScriptVersionNewResponseResourcesScriptRuntimeExport()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptVersionNewResponseResourcesScriptRuntimeExportsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExport{}),
			DiscriminatorValue: "worker",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExport{}),
			DiscriminatorValue: "durable-object",
		},
	)
}

// A named Worker entrypoint export (`type: worker`). Worker entrypoints are always
// live (`state: created`) and carry no storage or lifecycle fields. The optional
// `cache` block overrides the Worker's global `cache_options.enabled` for this
// entrypoint.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExport struct {
	// Marks this entry as a Worker entrypoint export.
	Type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportType `json:"type" api:"required"`
	// Cache override for this entrypoint. Overrides the Worker's global
	// `cache_options.enabled` for this entrypoint only.
	Cache ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportCache `json:"cache"`
	// Live export. May be omitted; defaults to `created`.
	State ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportState `json:"state"`
	JSON  scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportJSON  `json:"-"`
}

// scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExport]
type scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportJSON struct {
	Type        apijson.Field
	Cache       apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExport) implementsScriptVersionNewResponseResourcesScriptRuntimeExport() {
}

// Marks this entry as a Worker entrypoint export.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportType string

const (
	ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportTypeWorker ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportType = "worker"
)

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportTypeWorker:
		return true
	}
	return false
}

// Cache override for this entrypoint. Overrides the Worker's global
// `cache_options.enabled` for this entrypoint only.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportCache struct {
	// Whether caching is enabled for this entrypoint.
	Enabled bool                                                                              `json:"enabled" api:"required"`
	JSON    scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportCacheJSON `json:"-"`
}

// scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportCacheJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportCache]
type scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportCacheJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportCacheJSON) RawJSON() string {
	return r.raw
}

// Live export. May be omitted; defaults to `created`.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportState string

const (
	ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportStateCreated ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportState = "created"
)

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportState) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersWorkerExportStateCreated:
		return true
	}
	return false
}

// A live Durable Object export (`state: created`, the default). The platform
// auto-provisions the namespace on first deploy, matches it on subsequent deploys,
// and never mutates or deletes it as a side effect of a code-only change.
// `storage` is required; `renamed_to`, `transferred_to` and `transfer_from` are
// not allowed on a live entry.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExport struct {
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStorage `json:"storage" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportType `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container string `json:"container"`
	// Live export. May be omitted; defaults to `created`.
	State ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportState `json:"state"`
	JSON  scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportJSON  `json:"-"`
}

// scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExport]
type scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportJSON struct {
	Storage     apijson.Field
	Type        apijson.Field
	Container   apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExport) implementsScriptVersionNewResponseResourcesScriptRuntimeExport() {
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStorage string

const (
	ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStorageSqlite   ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStorage = "sqlite"
	ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStorageLegacyKV ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStorage = "legacy-kv"
)

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStorage) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStorageSqlite, ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStorageLegacyKV:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportType string

const (
	ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportTypeDurableObject ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportType = "durable-object"
)

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportTypeDurableObject:
		return true
	}
	return false
}

// Live export. May be omitted; defaults to `created`.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportState string

const (
	ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStateCreated ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportState = "created"
)

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportState) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStateCreated:
		return true
	}
	return false
}

// A `deleted` tombstone: retires the provisioned namespace for this class and all
// of its data. The class must be absent from the uploaded code and no other Worker
// in the account may bind to the namespace, otherwise the deploy is rejected. No
// other fields are allowed. Deletion is irreversible.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExport struct {
	// Tombstone that deletes the namespace.
	State ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportState `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportType `json:"type" api:"required"`
	JSON scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportJSON `json:"-"`
}

// scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExport]
type scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportJSON struct {
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExport) implementsScriptVersionNewResponseResourcesScriptRuntimeExport() {
}

// Tombstone that deletes the namespace.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportState string

const (
	ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportStateDeleted ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportState = "deleted"
)

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportState) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportStateDeleted:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportType string

const (
	ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportTypeDurableObject ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportType = "durable-object"
)

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportTypeDurableObject:
		return true
	}
	return false
}

// A `renamed` tombstone: rewrites the provisioned namespace's class name from this
// map key to `renamed_to`. The source class may stay in code during the rollout
// window (an info notice is emitted). `storage`, `transferred_to` and
// `transfer_from` are not allowed.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExport struct {
	// Tombstone that renames the namespace's class.
	State ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportState `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportType `json:"type" api:"required"`
	JSON scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportJSON `json:"-"`
}

// scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExport]
type scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportJSON struct {
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExport) implementsScriptVersionNewResponseResourcesScriptRuntimeExport() {
}

// Tombstone that renames the namespace's class.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportState string

const (
	ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportStateRenamed ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportState = "renamed"
)

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportState) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportStateRenamed:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportType string

const (
	ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportTypeDurableObject ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportType = "durable-object"
)

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportTypeDurableObject:
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
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExport struct {
	// Tombstone that transfers the namespace to another script.
	State ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportState `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportType `json:"type" api:"required"`
	JSON scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportJSON `json:"-"`
}

// scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExport]
type scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportJSON struct {
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExport) implementsScriptVersionNewResponseResourcesScriptRuntimeExport() {
}

// Tombstone that transfers the namespace to another script.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportState string

const (
	ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportStateTransferred ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportState = "transferred"
)

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportState) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportStateTransferred:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportType string

const (
	ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportTypeDurableObject ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportType = "durable-object"
)

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportTypeDurableObject:
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
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExport struct {
	// Target side of a two-phase transfer.
	State ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportState `json:"state" api:"required"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStorage `json:"storage" api:"required"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom string `json:"transfer_from" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportType `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object once the transfer settles. Valid only on live entries.
	Container string                                                                                               `json:"container"`
	JSON      scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportJSON `json:"-"`
}

// scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportJSON
// contains the JSON metadata for the struct
// [ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExport]
type scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportJSON struct {
	State        apijson.Field
	Storage      apijson.Field
	TransferFrom apijson.Field
	Type         apijson.Field
	Container    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExport) implementsScriptVersionNewResponseResourcesScriptRuntimeExport() {
}

// Target side of a two-phase transfer.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportState string

const (
	ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStateExpectingTransfer ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportState = "expecting-transfer"
)

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportState) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStateExpectingTransfer:
		return true
	}
	return false
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStorage string

const (
	ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStorageSqlite   ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStorage = "sqlite"
	ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStorageLegacyKV ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStorage = "legacy-kv"
)

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStorage) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStorageSqlite, ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStorageLegacyKV:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportType string

const (
	ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportTypeDurableObject ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportType = "durable-object"
)

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportTypeDurableObject:
		return true
	}
	return false
}

// Marks this entry as a Worker entrypoint export.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsType string

const (
	ScriptVersionNewResponseResourcesScriptRuntimeExportsTypeWorker        ScriptVersionNewResponseResourcesScriptRuntimeExportsType = "worker"
	ScriptVersionNewResponseResourcesScriptRuntimeExportsTypeDurableObject ScriptVersionNewResponseResourcesScriptRuntimeExportsType = "durable-object"
)

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsType) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesScriptRuntimeExportsTypeWorker, ScriptVersionNewResponseResourcesScriptRuntimeExportsTypeDurableObject:
		return true
	}
	return false
}

// Live export. May be omitted; defaults to `created`.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsState string

const (
	ScriptVersionNewResponseResourcesScriptRuntimeExportsStateCreated           ScriptVersionNewResponseResourcesScriptRuntimeExportsState = "created"
	ScriptVersionNewResponseResourcesScriptRuntimeExportsStateDeleted           ScriptVersionNewResponseResourcesScriptRuntimeExportsState = "deleted"
	ScriptVersionNewResponseResourcesScriptRuntimeExportsStateRenamed           ScriptVersionNewResponseResourcesScriptRuntimeExportsState = "renamed"
	ScriptVersionNewResponseResourcesScriptRuntimeExportsStateTransferred       ScriptVersionNewResponseResourcesScriptRuntimeExportsState = "transferred"
	ScriptVersionNewResponseResourcesScriptRuntimeExportsStateExpectingTransfer ScriptVersionNewResponseResourcesScriptRuntimeExportsState = "expecting-transfer"
)

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsState) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesScriptRuntimeExportsStateCreated, ScriptVersionNewResponseResourcesScriptRuntimeExportsStateDeleted, ScriptVersionNewResponseResourcesScriptRuntimeExportsStateRenamed, ScriptVersionNewResponseResourcesScriptRuntimeExportsStateTransferred, ScriptVersionNewResponseResourcesScriptRuntimeExportsStateExpectingTransfer:
		return true
	}
	return false
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type ScriptVersionNewResponseResourcesScriptRuntimeExportsStorage string

const (
	ScriptVersionNewResponseResourcesScriptRuntimeExportsStorageSqlite   ScriptVersionNewResponseResourcesScriptRuntimeExportsStorage = "sqlite"
	ScriptVersionNewResponseResourcesScriptRuntimeExportsStorageLegacyKV ScriptVersionNewResponseResourcesScriptRuntimeExportsStorage = "legacy-kv"
)

func (r ScriptVersionNewResponseResourcesScriptRuntimeExportsStorage) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseResourcesScriptRuntimeExportsStorageSqlite, ScriptVersionNewResponseResourcesScriptRuntimeExportsStorageLegacyKV:
		return true
	}
	return false
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

// Summary of the declarative exports reconciliation that ran on this upload.
// Populated only when the uploaded metadata included an `exports` block. Durable
// Object entries drive reconciliation; `type: worker` entries do not contribute to
// this summary.
type ScriptVersionNewResponseExportsReconciliation struct {
	// Class names for which a new namespace was provisioned.
	Created []string `json:"created" api:"required"`
	// Class names whose namespace was deleted by a `deleted` tombstone.
	Deleted []string `json:"deleted" api:"required"`
	// Non-blocking info entries (stale tombstones, tombstone applied with class still
	// in code). See `exports_reconciliation_info`.
	Info []ScriptVersionNewResponseExportsReconciliationInfo `json:"info" api:"required"`
	// Source class names whose tombstone entry is now stale and safe to delete from
	// `exports` (no remaining referencing scripts).
	RemovableEntries []string `json:"removable_entries" api:"required"`
	// Applied `renamed` tombstones.
	Renamed []ScriptVersionNewResponseExportsReconciliationRenamed `json:"renamed" api:"required"`
	// Phase-1 transfer hints recorded on the target side.
	TransferPending []ScriptVersionNewResponseExportsReconciliationTransferPending `json:"transfer_pending" api:"required"`
	// Committed `transferred` tombstones (phase-2).
	Transferred []ScriptVersionNewResponseExportsReconciliationTransferred `json:"transferred" api:"required"`
	// Class names whose provisioned namespace was mutated in place.
	Updated []string `json:"updated" api:"required"`
	// Non-blocking warnings. See `exports_reconciliation_warning`.
	Warnings []ScriptVersionNewResponseExportsReconciliationWarning `json:"warnings" api:"required"`
	JSON     scriptVersionNewResponseExportsReconciliationJSON      `json:"-"`
}

// scriptVersionNewResponseExportsReconciliationJSON contains the JSON metadata for
// the struct [ScriptVersionNewResponseExportsReconciliation]
type scriptVersionNewResponseExportsReconciliationJSON struct {
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

func (r *ScriptVersionNewResponseExportsReconciliation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseExportsReconciliationJSON) RawJSON() string {
	return r.raw
}

// A non-blocking reconciliation info entry. Emitted for stale tombstones (a no-op
// on this deploy) and for tombstones applied with the source class still in code
// (the supported zero-downtime rollout pattern).
type ScriptVersionNewResponseExportsReconciliationInfo struct {
	// The class name the info entry is about.
	Class string `json:"class" api:"required"`
	// Human-readable explanation.
	Message string `json:"message" api:"required"`
	// Stable, machine-readable tag identifying which reconciliation scenario produced
	// an error, warning, or info entry. Clients may branch on this value instead of
	// parsing `message`.
	Scenario ScriptVersionNewResponseExportsReconciliationInfoScenario `json:"scenario" api:"required"`
	// The provisioned namespace the entry relates to, when applicable.
	NamespaceID string `json:"namespace_id" format:"uuid"`
	// Other Workers in the account that still bind to the affected class. Advisory:
	// while non-empty the tombstone is not yet safe to remove — redeploy these Workers
	// with bindings re-pointed first.
	ReferencingScripts []string                                              `json:"referencing_scripts"`
	JSON               scriptVersionNewResponseExportsReconciliationInfoJSON `json:"-"`
}

// scriptVersionNewResponseExportsReconciliationInfoJSON contains the JSON metadata
// for the struct [ScriptVersionNewResponseExportsReconciliationInfo]
type scriptVersionNewResponseExportsReconciliationInfoJSON struct {
	Class              apijson.Field
	Message            apijson.Field
	Scenario           apijson.Field
	NamespaceID        apijson.Field
	ReferencingScripts apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptVersionNewResponseExportsReconciliationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseExportsReconciliationInfoJSON) RawJSON() string {
	return r.raw
}

// Stable, machine-readable tag identifying which reconciliation scenario produced
// an error, warning, or info entry. Clients may branch on this value instead of
// parsing `message`.
type ScriptVersionNewResponseExportsReconciliationInfoScenario string

const (
	ScriptVersionNewResponseExportsReconciliationInfoScenarioCodeClassNotInExports                     ScriptVersionNewResponseExportsReconciliationInfoScenario = "code_class_not_in_exports"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioProvisionedClassMissingFromConfig         ScriptVersionNewResponseExportsReconciliationInfoScenario = "provisioned_class_missing_from_config"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioConfigExportNotInCode                     ScriptVersionNewResponseExportsReconciliationInfoScenario = "config_export_not_in_code"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioConfigReferencesNonexistentClass          ScriptVersionNewResponseExportsReconciliationInfoScenario = "config_references_nonexistent_class"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioOrphanedProvisionedNamespace              ScriptVersionNewResponseExportsReconciliationInfoScenario = "orphaned_provisioned_namespace"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioStorageTypeMismatch                       ScriptVersionNewResponseExportsReconciliationInfoScenario = "storage_type_mismatch"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioFreeTierRequiresSqlite                    ScriptVersionNewResponseExportsReconciliationInfoScenario = "free_tier_requires_sqlite"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioInvalidExport                             ScriptVersionNewResponseExportsReconciliationInfoScenario = "invalid_export"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioTombstoneDeleteClassStillInCode           ScriptVersionNewResponseExportsReconciliationInfoScenario = "tombstone_delete_class_still_in_code"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioTombstoneDeleteBlockedByExternalBindings  ScriptVersionNewResponseExportsReconciliationInfoScenario = "tombstone_delete_blocked_by_external_bindings"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioTombstoneRenamedToOccupied                ScriptVersionNewResponseExportsReconciliationInfoScenario = "tombstone_renamed_to_occupied"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioTransferredPendingNotFound                ScriptVersionNewResponseExportsReconciliationInfoScenario = "transferred_pending_not_found"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioTransferredTargetMissing                  ScriptVersionNewResponseExportsReconciliationInfoScenario = "transferred_target_missing"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioTransferredTargetMismatch                 ScriptVersionNewResponseExportsReconciliationInfoScenario = "transferred_target_mismatch"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioPhaseOneTransferSourceMissing             ScriptVersionNewResponseExportsReconciliationInfoScenario = "phase_one_transfer_source_missing"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioPhaseOneTransferSourceNamespaceMissing    ScriptVersionNewResponseExportsReconciliationInfoScenario = "phase_one_transfer_source_namespace_missing"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioPhaseOneTransferTargetClassProvisioned    ScriptVersionNewResponseExportsReconciliationInfoScenario = "phase_one_transfer_target_class_provisioned"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioPhaseOneTransferAfterCommitMismatch       ScriptVersionNewResponseExportsReconciliationInfoScenario = "phase_one_transfer_after_commit_mismatch"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioPhaseOneTransferDuplicate                 ScriptVersionNewResponseExportsReconciliationInfoScenario = "phase_one_transfer_duplicate"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioPhaseOneTransferTargetInDispatchNamespace ScriptVersionNewResponseExportsReconciliationInfoScenario = "phase_one_transfer_target_in_dispatch_namespace"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioPhaseOneTransferSourceInDispatchNamespace ScriptVersionNewResponseExportsReconciliationInfoScenario = "phase_one_transfer_source_in_dispatch_namespace"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioTransferredSourceInDispatchNamespace      ScriptVersionNewResponseExportsReconciliationInfoScenario = "transferred_source_in_dispatch_namespace"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioTransferredTargetInDispatchNamespace      ScriptVersionNewResponseExportsReconciliationInfoScenario = "transferred_target_in_dispatch_namespace"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioContainerUndeclaredReference              ScriptVersionNewResponseExportsReconciliationInfoScenario = "container_undeclared_reference"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioContainerClassNotDurableObject            ScriptVersionNewResponseExportsReconciliationInfoScenario = "container_class_not_durable_object"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioContainerWiringInconsistent               ScriptVersionNewResponseExportsReconciliationInfoScenario = "container_wiring_inconsistent"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioContainerMultipleDurableObjects           ScriptVersionNewResponseExportsReconciliationInfoScenario = "container_multiple_durable_objects"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioTransferContainerParityMismatch           ScriptVersionNewResponseExportsReconciliationInfoScenario = "transfer_container_parity_mismatch"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioTransferContainerParityMismatchOnCommit   ScriptVersionNewResponseExportsReconciliationInfoScenario = "transfer_container_parity_mismatch_on_commit"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioTombstoneClassStillInCode                 ScriptVersionNewResponseExportsReconciliationInfoScenario = "tombstone_class_still_in_code"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioStaleTombstone                            ScriptVersionNewResponseExportsReconciliationInfoScenario = "stale_tombstone"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioTransferReceiveAlreadyApplied             ScriptVersionNewResponseExportsReconciliationInfoScenario = "transfer_receive_already_applied"
	ScriptVersionNewResponseExportsReconciliationInfoScenarioTransferReceiveCleanupComplete            ScriptVersionNewResponseExportsReconciliationInfoScenario = "transfer_receive_cleanup_complete"
)

func (r ScriptVersionNewResponseExportsReconciliationInfoScenario) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseExportsReconciliationInfoScenarioCodeClassNotInExports, ScriptVersionNewResponseExportsReconciliationInfoScenarioProvisionedClassMissingFromConfig, ScriptVersionNewResponseExportsReconciliationInfoScenarioConfigExportNotInCode, ScriptVersionNewResponseExportsReconciliationInfoScenarioConfigReferencesNonexistentClass, ScriptVersionNewResponseExportsReconciliationInfoScenarioOrphanedProvisionedNamespace, ScriptVersionNewResponseExportsReconciliationInfoScenarioStorageTypeMismatch, ScriptVersionNewResponseExportsReconciliationInfoScenarioFreeTierRequiresSqlite, ScriptVersionNewResponseExportsReconciliationInfoScenarioInvalidExport, ScriptVersionNewResponseExportsReconciliationInfoScenarioTombstoneDeleteClassStillInCode, ScriptVersionNewResponseExportsReconciliationInfoScenarioTombstoneDeleteBlockedByExternalBindings, ScriptVersionNewResponseExportsReconciliationInfoScenarioTombstoneRenamedToOccupied, ScriptVersionNewResponseExportsReconciliationInfoScenarioTransferredPendingNotFound, ScriptVersionNewResponseExportsReconciliationInfoScenarioTransferredTargetMissing, ScriptVersionNewResponseExportsReconciliationInfoScenarioTransferredTargetMismatch, ScriptVersionNewResponseExportsReconciliationInfoScenarioPhaseOneTransferSourceMissing, ScriptVersionNewResponseExportsReconciliationInfoScenarioPhaseOneTransferSourceNamespaceMissing, ScriptVersionNewResponseExportsReconciliationInfoScenarioPhaseOneTransferTargetClassProvisioned, ScriptVersionNewResponseExportsReconciliationInfoScenarioPhaseOneTransferAfterCommitMismatch, ScriptVersionNewResponseExportsReconciliationInfoScenarioPhaseOneTransferDuplicate, ScriptVersionNewResponseExportsReconciliationInfoScenarioPhaseOneTransferTargetInDispatchNamespace, ScriptVersionNewResponseExportsReconciliationInfoScenarioPhaseOneTransferSourceInDispatchNamespace, ScriptVersionNewResponseExportsReconciliationInfoScenarioTransferredSourceInDispatchNamespace, ScriptVersionNewResponseExportsReconciliationInfoScenarioTransferredTargetInDispatchNamespace, ScriptVersionNewResponseExportsReconciliationInfoScenarioContainerUndeclaredReference, ScriptVersionNewResponseExportsReconciliationInfoScenarioContainerClassNotDurableObject, ScriptVersionNewResponseExportsReconciliationInfoScenarioContainerWiringInconsistent, ScriptVersionNewResponseExportsReconciliationInfoScenarioContainerMultipleDurableObjects, ScriptVersionNewResponseExportsReconciliationInfoScenarioTransferContainerParityMismatch, ScriptVersionNewResponseExportsReconciliationInfoScenarioTransferContainerParityMismatchOnCommit, ScriptVersionNewResponseExportsReconciliationInfoScenarioTombstoneClassStillInCode, ScriptVersionNewResponseExportsReconciliationInfoScenarioStaleTombstone, ScriptVersionNewResponseExportsReconciliationInfoScenarioTransferReceiveAlreadyApplied, ScriptVersionNewResponseExportsReconciliationInfoScenarioTransferReceiveCleanupComplete:
		return true
	}
	return false
}

// A single applied `renamed` tombstone.
type ScriptVersionNewResponseExportsReconciliationRenamed struct {
	// The original (source) class name.
	From string `json:"from" api:"required"`
	// The new class name (`renamed_to`).
	To   string                                                   `json:"to" api:"required"`
	JSON scriptVersionNewResponseExportsReconciliationRenamedJSON `json:"-"`
}

// scriptVersionNewResponseExportsReconciliationRenamedJSON contains the JSON
// metadata for the struct [ScriptVersionNewResponseExportsReconciliationRenamed]
type scriptVersionNewResponseExportsReconciliationRenamedJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseExportsReconciliationRenamed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseExportsReconciliationRenamedJSON) RawJSON() string {
	return r.raw
}

// A single phase-1 transfer hint recorded on the target side (a live
// `expecting-transfer` entry).
type ScriptVersionNewResponseExportsReconciliationTransferPending struct {
	// The target-side class name awaiting transfer.
	Class string `json:"class" api:"required"`
	// The source script the namespace will be transferred from.
	From string                                                           `json:"from" api:"required"`
	JSON scriptVersionNewResponseExportsReconciliationTransferPendingJSON `json:"-"`
}

// scriptVersionNewResponseExportsReconciliationTransferPendingJSON contains the
// JSON metadata for the struct
// [ScriptVersionNewResponseExportsReconciliationTransferPending]
type scriptVersionNewResponseExportsReconciliationTransferPendingJSON struct {
	Class       apijson.Field
	From        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseExportsReconciliationTransferPending) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseExportsReconciliationTransferPendingJSON) RawJSON() string {
	return r.raw
}

// A single committed `transferred` tombstone (phase-2 commit).
type ScriptVersionNewResponseExportsReconciliationTransferred struct {
	// The source class name that was transferred.
	Class string `json:"class" api:"required"`
	// The transfer phase. Currently always `committed`.
	Phase ScriptVersionNewResponseExportsReconciliationTransferredPhase `json:"phase" api:"required"`
	// The destination script that now owns the namespace.
	To   string                                                       `json:"to" api:"required"`
	JSON scriptVersionNewResponseExportsReconciliationTransferredJSON `json:"-"`
}

// scriptVersionNewResponseExportsReconciliationTransferredJSON contains the JSON
// metadata for the struct
// [ScriptVersionNewResponseExportsReconciliationTransferred]
type scriptVersionNewResponseExportsReconciliationTransferredJSON struct {
	Class       apijson.Field
	Phase       apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseExportsReconciliationTransferred) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseExportsReconciliationTransferredJSON) RawJSON() string {
	return r.raw
}

// The transfer phase. Currently always `committed`.
type ScriptVersionNewResponseExportsReconciliationTransferredPhase string

const (
	ScriptVersionNewResponseExportsReconciliationTransferredPhaseCommitted ScriptVersionNewResponseExportsReconciliationTransferredPhase = "committed"
)

func (r ScriptVersionNewResponseExportsReconciliationTransferredPhase) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseExportsReconciliationTransferredPhaseCommitted:
		return true
	}
	return false
}

// A non-blocking reconciliation warning. Reserved: no scenario populates this
// array today (`code_class_not_in_exports` is surfaced as info and
// `provisioned_class_missing_from_config` is a hard error). Clients should still
// surface any entries that appear.
type ScriptVersionNewResponseExportsReconciliationWarning struct {
	// The class name the warning is about.
	Class string `json:"class" api:"required"`
	// Human-readable explanation of the warning.
	Message string `json:"message" api:"required"`
	// Stable, machine-readable tag identifying which reconciliation scenario produced
	// an error, warning, or info entry. Clients may branch on this value instead of
	// parsing `message`.
	Scenario ScriptVersionNewResponseExportsReconciliationWarningsScenario `json:"scenario" api:"required"`
	// The provisioned namespace the warning relates to, when applicable.
	NamespaceID string                                                   `json:"namespace_id" format:"uuid"`
	JSON        scriptVersionNewResponseExportsReconciliationWarningJSON `json:"-"`
}

// scriptVersionNewResponseExportsReconciliationWarningJSON contains the JSON
// metadata for the struct [ScriptVersionNewResponseExportsReconciliationWarning]
type scriptVersionNewResponseExportsReconciliationWarningJSON struct {
	Class       apijson.Field
	Message     apijson.Field
	Scenario    apijson.Field
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionNewResponseExportsReconciliationWarning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionNewResponseExportsReconciliationWarningJSON) RawJSON() string {
	return r.raw
}

// Stable, machine-readable tag identifying which reconciliation scenario produced
// an error, warning, or info entry. Clients may branch on this value instead of
// parsing `message`.
type ScriptVersionNewResponseExportsReconciliationWarningsScenario string

const (
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioCodeClassNotInExports                     ScriptVersionNewResponseExportsReconciliationWarningsScenario = "code_class_not_in_exports"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioProvisionedClassMissingFromConfig         ScriptVersionNewResponseExportsReconciliationWarningsScenario = "provisioned_class_missing_from_config"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioConfigExportNotInCode                     ScriptVersionNewResponseExportsReconciliationWarningsScenario = "config_export_not_in_code"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioConfigReferencesNonexistentClass          ScriptVersionNewResponseExportsReconciliationWarningsScenario = "config_references_nonexistent_class"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioOrphanedProvisionedNamespace              ScriptVersionNewResponseExportsReconciliationWarningsScenario = "orphaned_provisioned_namespace"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioStorageTypeMismatch                       ScriptVersionNewResponseExportsReconciliationWarningsScenario = "storage_type_mismatch"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioFreeTierRequiresSqlite                    ScriptVersionNewResponseExportsReconciliationWarningsScenario = "free_tier_requires_sqlite"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioInvalidExport                             ScriptVersionNewResponseExportsReconciliationWarningsScenario = "invalid_export"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioTombstoneDeleteClassStillInCode           ScriptVersionNewResponseExportsReconciliationWarningsScenario = "tombstone_delete_class_still_in_code"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioTombstoneDeleteBlockedByExternalBindings  ScriptVersionNewResponseExportsReconciliationWarningsScenario = "tombstone_delete_blocked_by_external_bindings"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioTombstoneRenamedToOccupied                ScriptVersionNewResponseExportsReconciliationWarningsScenario = "tombstone_renamed_to_occupied"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioTransferredPendingNotFound                ScriptVersionNewResponseExportsReconciliationWarningsScenario = "transferred_pending_not_found"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioTransferredTargetMissing                  ScriptVersionNewResponseExportsReconciliationWarningsScenario = "transferred_target_missing"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioTransferredTargetMismatch                 ScriptVersionNewResponseExportsReconciliationWarningsScenario = "transferred_target_mismatch"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioPhaseOneTransferSourceMissing             ScriptVersionNewResponseExportsReconciliationWarningsScenario = "phase_one_transfer_source_missing"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioPhaseOneTransferSourceNamespaceMissing    ScriptVersionNewResponseExportsReconciliationWarningsScenario = "phase_one_transfer_source_namespace_missing"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioPhaseOneTransferTargetClassProvisioned    ScriptVersionNewResponseExportsReconciliationWarningsScenario = "phase_one_transfer_target_class_provisioned"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioPhaseOneTransferAfterCommitMismatch       ScriptVersionNewResponseExportsReconciliationWarningsScenario = "phase_one_transfer_after_commit_mismatch"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioPhaseOneTransferDuplicate                 ScriptVersionNewResponseExportsReconciliationWarningsScenario = "phase_one_transfer_duplicate"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioPhaseOneTransferTargetInDispatchNamespace ScriptVersionNewResponseExportsReconciliationWarningsScenario = "phase_one_transfer_target_in_dispatch_namespace"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioPhaseOneTransferSourceInDispatchNamespace ScriptVersionNewResponseExportsReconciliationWarningsScenario = "phase_one_transfer_source_in_dispatch_namespace"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioTransferredSourceInDispatchNamespace      ScriptVersionNewResponseExportsReconciliationWarningsScenario = "transferred_source_in_dispatch_namespace"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioTransferredTargetInDispatchNamespace      ScriptVersionNewResponseExportsReconciliationWarningsScenario = "transferred_target_in_dispatch_namespace"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioContainerUndeclaredReference              ScriptVersionNewResponseExportsReconciliationWarningsScenario = "container_undeclared_reference"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioContainerClassNotDurableObject            ScriptVersionNewResponseExportsReconciliationWarningsScenario = "container_class_not_durable_object"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioContainerWiringInconsistent               ScriptVersionNewResponseExportsReconciliationWarningsScenario = "container_wiring_inconsistent"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioContainerMultipleDurableObjects           ScriptVersionNewResponseExportsReconciliationWarningsScenario = "container_multiple_durable_objects"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioTransferContainerParityMismatch           ScriptVersionNewResponseExportsReconciliationWarningsScenario = "transfer_container_parity_mismatch"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioTransferContainerParityMismatchOnCommit   ScriptVersionNewResponseExportsReconciliationWarningsScenario = "transfer_container_parity_mismatch_on_commit"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioTombstoneClassStillInCode                 ScriptVersionNewResponseExportsReconciliationWarningsScenario = "tombstone_class_still_in_code"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioStaleTombstone                            ScriptVersionNewResponseExportsReconciliationWarningsScenario = "stale_tombstone"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioTransferReceiveAlreadyApplied             ScriptVersionNewResponseExportsReconciliationWarningsScenario = "transfer_receive_already_applied"
	ScriptVersionNewResponseExportsReconciliationWarningsScenarioTransferReceiveCleanupComplete            ScriptVersionNewResponseExportsReconciliationWarningsScenario = "transfer_receive_cleanup_complete"
)

func (r ScriptVersionNewResponseExportsReconciliationWarningsScenario) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseExportsReconciliationWarningsScenarioCodeClassNotInExports, ScriptVersionNewResponseExportsReconciliationWarningsScenarioProvisionedClassMissingFromConfig, ScriptVersionNewResponseExportsReconciliationWarningsScenarioConfigExportNotInCode, ScriptVersionNewResponseExportsReconciliationWarningsScenarioConfigReferencesNonexistentClass, ScriptVersionNewResponseExportsReconciliationWarningsScenarioOrphanedProvisionedNamespace, ScriptVersionNewResponseExportsReconciliationWarningsScenarioStorageTypeMismatch, ScriptVersionNewResponseExportsReconciliationWarningsScenarioFreeTierRequiresSqlite, ScriptVersionNewResponseExportsReconciliationWarningsScenarioInvalidExport, ScriptVersionNewResponseExportsReconciliationWarningsScenarioTombstoneDeleteClassStillInCode, ScriptVersionNewResponseExportsReconciliationWarningsScenarioTombstoneDeleteBlockedByExternalBindings, ScriptVersionNewResponseExportsReconciliationWarningsScenarioTombstoneRenamedToOccupied, ScriptVersionNewResponseExportsReconciliationWarningsScenarioTransferredPendingNotFound, ScriptVersionNewResponseExportsReconciliationWarningsScenarioTransferredTargetMissing, ScriptVersionNewResponseExportsReconciliationWarningsScenarioTransferredTargetMismatch, ScriptVersionNewResponseExportsReconciliationWarningsScenarioPhaseOneTransferSourceMissing, ScriptVersionNewResponseExportsReconciliationWarningsScenarioPhaseOneTransferSourceNamespaceMissing, ScriptVersionNewResponseExportsReconciliationWarningsScenarioPhaseOneTransferTargetClassProvisioned, ScriptVersionNewResponseExportsReconciliationWarningsScenarioPhaseOneTransferAfterCommitMismatch, ScriptVersionNewResponseExportsReconciliationWarningsScenarioPhaseOneTransferDuplicate, ScriptVersionNewResponseExportsReconciliationWarningsScenarioPhaseOneTransferTargetInDispatchNamespace, ScriptVersionNewResponseExportsReconciliationWarningsScenarioPhaseOneTransferSourceInDispatchNamespace, ScriptVersionNewResponseExportsReconciliationWarningsScenarioTransferredSourceInDispatchNamespace, ScriptVersionNewResponseExportsReconciliationWarningsScenarioTransferredTargetInDispatchNamespace, ScriptVersionNewResponseExportsReconciliationWarningsScenarioContainerUndeclaredReference, ScriptVersionNewResponseExportsReconciliationWarningsScenarioContainerClassNotDurableObject, ScriptVersionNewResponseExportsReconciliationWarningsScenarioContainerWiringInconsistent, ScriptVersionNewResponseExportsReconciliationWarningsScenarioContainerMultipleDurableObjects, ScriptVersionNewResponseExportsReconciliationWarningsScenarioTransferContainerParityMismatch, ScriptVersionNewResponseExportsReconciliationWarningsScenarioTransferContainerParityMismatchOnCommit, ScriptVersionNewResponseExportsReconciliationWarningsScenarioTombstoneClassStillInCode, ScriptVersionNewResponseExportsReconciliationWarningsScenarioStaleTombstone, ScriptVersionNewResponseExportsReconciliationWarningsScenarioTransferReceiveAlreadyApplied, ScriptVersionNewResponseExportsReconciliationWarningsScenarioTransferReceiveCleanupComplete:
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
	ScriptVersionNewResponseMetadataSourceCfCli        ScriptVersionNewResponseMetadataSource = "cf_cli"
	ScriptVersionNewResponseMetadataSourceDashTemplate ScriptVersionNewResponseMetadataSource = "dash_template"
	ScriptVersionNewResponseMetadataSourceIntegration  ScriptVersionNewResponseMetadataSource = "integration"
	ScriptVersionNewResponseMetadataSourceQuickEditor  ScriptVersionNewResponseMetadataSource = "quick_editor"
	ScriptVersionNewResponseMetadataSourcePlayground   ScriptVersionNewResponseMetadataSource = "playground"
	ScriptVersionNewResponseMetadataSourceWorkersci    ScriptVersionNewResponseMetadataSource = "workersci"
)

func (r ScriptVersionNewResponseMetadataSource) IsKnown() bool {
	switch r {
	case ScriptVersionNewResponseMetadataSourceUnknown, ScriptVersionNewResponseMetadataSourceAPI, ScriptVersionNewResponseMetadataSourceWrangler, ScriptVersionNewResponseMetadataSourceTerraform, ScriptVersionNewResponseMetadataSourceDash, ScriptVersionNewResponseMetadataSourceCfCli, ScriptVersionNewResponseMetadataSourceDashTemplate, ScriptVersionNewResponseMetadataSourceIntegration, ScriptVersionNewResponseMetadataSourceQuickEditor, ScriptVersionNewResponseMetadataSourcePlayground, ScriptVersionNewResponseMetadataSourceWorkersci:
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
	ScriptVersionListResponseMetadataSourceCfCli        ScriptVersionListResponseMetadataSource = "cf_cli"
	ScriptVersionListResponseMetadataSourceDashTemplate ScriptVersionListResponseMetadataSource = "dash_template"
	ScriptVersionListResponseMetadataSourceIntegration  ScriptVersionListResponseMetadataSource = "integration"
	ScriptVersionListResponseMetadataSourceQuickEditor  ScriptVersionListResponseMetadataSource = "quick_editor"
	ScriptVersionListResponseMetadataSourcePlayground   ScriptVersionListResponseMetadataSource = "playground"
	ScriptVersionListResponseMetadataSourceWorkersci    ScriptVersionListResponseMetadataSource = "workersci"
)

func (r ScriptVersionListResponseMetadataSource) IsKnown() bool {
	switch r {
	case ScriptVersionListResponseMetadataSourceUnknown, ScriptVersionListResponseMetadataSourceAPI, ScriptVersionListResponseMetadataSourceWrangler, ScriptVersionListResponseMetadataSourceTerraform, ScriptVersionListResponseMetadataSourceDash, ScriptVersionListResponseMetadataSourceCfCli, ScriptVersionListResponseMetadataSourceDashTemplate, ScriptVersionListResponseMetadataSourceIntegration, ScriptVersionListResponseMetadataSourceQuickEditor, ScriptVersionListResponseMetadataSourcePlayground, ScriptVersionListResponseMetadataSourceWorkersci:
		return true
	}
	return false
}

type ScriptVersionGetResponse struct {
	Resources ScriptVersionGetResponseResources `json:"resources" api:"required"`
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
	Bindings ScriptVersionGetResponseResourcesBindings `json:"bindings"`
	Script   ScriptVersionGetResponseResourcesScript   `json:"script"`
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

// List of bindings attached to a Worker. You can find more about bindings on our
// docs:
// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
type ScriptVersionGetResponseResourcesBindings struct {
	JSON scriptVersionGetResponseResourcesBindingsJSON `json:"-"`
}

// scriptVersionGetResponseResourcesBindingsJSON contains the JSON metadata for the
// struct [ScriptVersionGetResponseResourcesBindings]
type scriptVersionGetResponseResourcesBindingsJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesBindings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesBindingsJSON) RawJSON() string {
	return r.raw
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
	// Declarative exports for this version, including Durable Object classes (with
	// their `storage` backend) and named Worker entrypoints. Tombstoned lifecycle
	// entries are omitted, so only live exports (`created` and `expecting-transfer`)
	// are returned.
	Exports map[string]ScriptVersionGetResponseResourcesScriptRuntimeExport `json:"exports"`
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
	Exports            apijson.Field
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

// A single entry in the `exports` map, keyed by export name (a `WorkerEntrypoint`
// class name, a Durable Object class name, or `default` for the Worker's default
// export). The `type` discriminator selects the top-level shape: `worker`
// entrypoint entries may carry `cache` configuration, while `durable-object`
// entries are further refined by the optional `state` field (default `created`).
// Tombstone states (`deleted`, `renamed`, `transferred`) express destructive
// lifecycle operations declaratively; `expecting-transfer` is the live target side
// of a transfer. The server validates the exact per-(type, state) field
// combinations; fields not listed for a variant are rejected.
type ScriptVersionGetResponseResourcesScriptRuntimeExport struct {
	// Marks this entry as a Worker entrypoint export.
	Type ScriptVersionGetResponseResourcesScriptRuntimeExportsType `json:"type" api:"required"`
	// This field can have the runtime type of
	// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportCache].
	Cache interface{} `json:"cache"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container string `json:"container"`
	// Live export. May be omitted; defaults to `created`.
	State ScriptVersionGetResponseResourcesScriptRuntimeExportsState `json:"state"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage ScriptVersionGetResponseResourcesScriptRuntimeExportsStorage `json:"storage"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom string                                                   `json:"transfer_from"`
	JSON         scriptVersionGetResponseResourcesScriptRuntimeExportJSON `json:"-"`
	union        ScriptVersionGetResponseResourcesScriptRuntimeExportsUnion
}

// scriptVersionGetResponseResourcesScriptRuntimeExportJSON contains the JSON
// metadata for the struct [ScriptVersionGetResponseResourcesScriptRuntimeExport]
type scriptVersionGetResponseResourcesScriptRuntimeExportJSON struct {
	Type         apijson.Field
	Cache        apijson.Field
	Container    apijson.Field
	State        apijson.Field
	Storage      apijson.Field
	TransferFrom apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r scriptVersionGetResponseResourcesScriptRuntimeExportJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptVersionGetResponseResourcesScriptRuntimeExport) UnmarshalJSON(data []byte) (err error) {
	*r = ScriptVersionGetResponseResourcesScriptRuntimeExport{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ScriptVersionGetResponseResourcesScriptRuntimeExportsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExport],
// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExport],
// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExport],
// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExport],
// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExport],
// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExport].
func (r ScriptVersionGetResponseResourcesScriptRuntimeExport) AsUnion() ScriptVersionGetResponseResourcesScriptRuntimeExportsUnion {
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
// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExport],
// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExport],
// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExport],
// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExport],
// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExport]
// or
// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExport].
type ScriptVersionGetResponseResourcesScriptRuntimeExportsUnion interface {
	implementsScriptVersionGetResponseResourcesScriptRuntimeExport()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptVersionGetResponseResourcesScriptRuntimeExportsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExport{}),
			DiscriminatorValue: "worker",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExport{}),
			DiscriminatorValue: "durable-object",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExport{}),
			DiscriminatorValue: "durable-object",
		},
	)
}

// A named Worker entrypoint export (`type: worker`). Worker entrypoints are always
// live (`state: created`) and carry no storage or lifecycle fields. The optional
// `cache` block overrides the Worker's global `cache_options.enabled` for this
// entrypoint.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExport struct {
	// Marks this entry as a Worker entrypoint export.
	Type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportType `json:"type" api:"required"`
	// Cache override for this entrypoint. Overrides the Worker's global
	// `cache_options.enabled` for this entrypoint only.
	Cache ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportCache `json:"cache"`
	// Live export. May be omitted; defaults to `created`.
	State ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportState `json:"state"`
	JSON  scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportJSON  `json:"-"`
}

// scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExport]
type scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportJSON struct {
	Type        apijson.Field
	Cache       apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExport) implementsScriptVersionGetResponseResourcesScriptRuntimeExport() {
}

// Marks this entry as a Worker entrypoint export.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportType string

const (
	ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportTypeWorker ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportType = "worker"
)

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportTypeWorker:
		return true
	}
	return false
}

// Cache override for this entrypoint. Overrides the Worker's global
// `cache_options.enabled` for this entrypoint only.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportCache struct {
	// Whether caching is enabled for this entrypoint.
	Enabled bool                                                                              `json:"enabled" api:"required"`
	JSON    scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportCacheJSON `json:"-"`
}

// scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportCacheJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportCache]
type scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportCacheJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportCacheJSON) RawJSON() string {
	return r.raw
}

// Live export. May be omitted; defaults to `created`.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportState string

const (
	ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportStateCreated ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportState = "created"
)

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportState) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersWorkerExportStateCreated:
		return true
	}
	return false
}

// A live Durable Object export (`state: created`, the default). The platform
// auto-provisions the namespace on first deploy, matches it on subsequent deploys,
// and never mutates or deletes it as a side effect of a code-only change.
// `storage` is required; `renamed_to`, `transferred_to` and `transfer_from` are
// not allowed on a live entry.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExport struct {
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStorage `json:"storage" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportType `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container string `json:"container"`
	// Live export. May be omitted; defaults to `created`.
	State ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportState `json:"state"`
	JSON  scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportJSON  `json:"-"`
}

// scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExport]
type scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportJSON struct {
	Storage     apijson.Field
	Type        apijson.Field
	Container   apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExport) implementsScriptVersionGetResponseResourcesScriptRuntimeExport() {
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStorage string

const (
	ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStorageSqlite   ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStorage = "sqlite"
	ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStorageLegacyKV ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStorage = "legacy-kv"
)

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStorage) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStorageSqlite, ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStorageLegacyKV:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportType string

const (
	ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportTypeDurableObject ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportType = "durable-object"
)

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportTypeDurableObject:
		return true
	}
	return false
}

// Live export. May be omitted; defaults to `created`.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportState string

const (
	ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStateCreated ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportState = "created"
)

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportState) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExportStateCreated:
		return true
	}
	return false
}

// A `deleted` tombstone: retires the provisioned namespace for this class and all
// of its data. The class must be absent from the uploaded code and no other Worker
// in the account may bind to the namespace, otherwise the deploy is rejected. No
// other fields are allowed. Deletion is irreversible.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExport struct {
	// Tombstone that deletes the namespace.
	State ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportState `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportType `json:"type" api:"required"`
	JSON scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportJSON `json:"-"`
}

// scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExport]
type scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportJSON struct {
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExport) implementsScriptVersionGetResponseResourcesScriptRuntimeExport() {
}

// Tombstone that deletes the namespace.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportState string

const (
	ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportStateDeleted ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportState = "deleted"
)

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportState) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportStateDeleted:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportType string

const (
	ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportTypeDurableObject ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportType = "durable-object"
)

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectDeletedExportTypeDurableObject:
		return true
	}
	return false
}

// A `renamed` tombstone: rewrites the provisioned namespace's class name from this
// map key to `renamed_to`. The source class may stay in code during the rollout
// window (an info notice is emitted). `storage`, `transferred_to` and
// `transfer_from` are not allowed.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExport struct {
	// Tombstone that renames the namespace's class.
	State ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportState `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportType `json:"type" api:"required"`
	JSON scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportJSON `json:"-"`
}

// scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExport]
type scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportJSON struct {
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExport) implementsScriptVersionGetResponseResourcesScriptRuntimeExport() {
}

// Tombstone that renames the namespace's class.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportState string

const (
	ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportStateRenamed ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportState = "renamed"
)

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportState) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportStateRenamed:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportType string

const (
	ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportTypeDurableObject ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportType = "durable-object"
)

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectRenamedExportTypeDurableObject:
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
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExport struct {
	// Tombstone that transfers the namespace to another script.
	State ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportState `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportType `json:"type" api:"required"`
	JSON scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportJSON `json:"-"`
}

// scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExport]
type scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportJSON struct {
	State       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExport) implementsScriptVersionGetResponseResourcesScriptRuntimeExport() {
}

// Tombstone that transfers the namespace to another script.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportState string

const (
	ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportStateTransferred ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportState = "transferred"
)

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportState) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportStateTransferred:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportType string

const (
	ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportTypeDurableObject ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportType = "durable-object"
)

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectTransferredExportTypeDurableObject:
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
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExport struct {
	// Target side of a two-phase transfer.
	State ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportState `json:"state" api:"required"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStorage `json:"storage" api:"required"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom string `json:"transfer_from" api:"required"`
	// Marks this entry as a Durable Object export.
	Type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportType `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object once the transfer settles. Valid only on live entries.
	Container string                                                                                               `json:"container"`
	JSON      scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportJSON `json:"-"`
}

// scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportJSON
// contains the JSON metadata for the struct
// [ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExport]
type scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportJSON struct {
	State        apijson.Field
	Storage      apijson.Field
	TransferFrom apijson.Field
	Type         apijson.Field
	Container    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExport) implementsScriptVersionGetResponseResourcesScriptRuntimeExport() {
}

// Target side of a two-phase transfer.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportState string

const (
	ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStateExpectingTransfer ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportState = "expecting-transfer"
)

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportState) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStateExpectingTransfer:
		return true
	}
	return false
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStorage string

const (
	ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStorageSqlite   ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStorage = "sqlite"
	ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStorageLegacyKV ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStorage = "legacy-kv"
)

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStorage) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStorageSqlite, ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportStorageLegacyKV:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportType string

const (
	ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportTypeDurableObject ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportType = "durable-object"
)

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesScriptRuntimeExportsWorkersDurableObjectExpectingTransferExportTypeDurableObject:
		return true
	}
	return false
}

// Marks this entry as a Worker entrypoint export.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsType string

const (
	ScriptVersionGetResponseResourcesScriptRuntimeExportsTypeWorker        ScriptVersionGetResponseResourcesScriptRuntimeExportsType = "worker"
	ScriptVersionGetResponseResourcesScriptRuntimeExportsTypeDurableObject ScriptVersionGetResponseResourcesScriptRuntimeExportsType = "durable-object"
)

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsType) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesScriptRuntimeExportsTypeWorker, ScriptVersionGetResponseResourcesScriptRuntimeExportsTypeDurableObject:
		return true
	}
	return false
}

// Live export. May be omitted; defaults to `created`.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsState string

const (
	ScriptVersionGetResponseResourcesScriptRuntimeExportsStateCreated           ScriptVersionGetResponseResourcesScriptRuntimeExportsState = "created"
	ScriptVersionGetResponseResourcesScriptRuntimeExportsStateDeleted           ScriptVersionGetResponseResourcesScriptRuntimeExportsState = "deleted"
	ScriptVersionGetResponseResourcesScriptRuntimeExportsStateRenamed           ScriptVersionGetResponseResourcesScriptRuntimeExportsState = "renamed"
	ScriptVersionGetResponseResourcesScriptRuntimeExportsStateTransferred       ScriptVersionGetResponseResourcesScriptRuntimeExportsState = "transferred"
	ScriptVersionGetResponseResourcesScriptRuntimeExportsStateExpectingTransfer ScriptVersionGetResponseResourcesScriptRuntimeExportsState = "expecting-transfer"
)

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsState) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesScriptRuntimeExportsStateCreated, ScriptVersionGetResponseResourcesScriptRuntimeExportsStateDeleted, ScriptVersionGetResponseResourcesScriptRuntimeExportsStateRenamed, ScriptVersionGetResponseResourcesScriptRuntimeExportsStateTransferred, ScriptVersionGetResponseResourcesScriptRuntimeExportsStateExpectingTransfer:
		return true
	}
	return false
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type ScriptVersionGetResponseResourcesScriptRuntimeExportsStorage string

const (
	ScriptVersionGetResponseResourcesScriptRuntimeExportsStorageSqlite   ScriptVersionGetResponseResourcesScriptRuntimeExportsStorage = "sqlite"
	ScriptVersionGetResponseResourcesScriptRuntimeExportsStorageLegacyKV ScriptVersionGetResponseResourcesScriptRuntimeExportsStorage = "legacy-kv"
)

func (r ScriptVersionGetResponseResourcesScriptRuntimeExportsStorage) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseResourcesScriptRuntimeExportsStorageSqlite, ScriptVersionGetResponseResourcesScriptRuntimeExportsStorageLegacyKV:
		return true
	}
	return false
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
	ScriptVersionGetResponseMetadataSourceCfCli        ScriptVersionGetResponseMetadataSource = "cf_cli"
	ScriptVersionGetResponseMetadataSourceDashTemplate ScriptVersionGetResponseMetadataSource = "dash_template"
	ScriptVersionGetResponseMetadataSourceIntegration  ScriptVersionGetResponseMetadataSource = "integration"
	ScriptVersionGetResponseMetadataSourceQuickEditor  ScriptVersionGetResponseMetadataSource = "quick_editor"
	ScriptVersionGetResponseMetadataSourcePlayground   ScriptVersionGetResponseMetadataSource = "playground"
	ScriptVersionGetResponseMetadataSourceWorkersci    ScriptVersionGetResponseMetadataSource = "workersci"
)

func (r ScriptVersionGetResponseMetadataSource) IsKnown() bool {
	switch r {
	case ScriptVersionGetResponseMetadataSourceUnknown, ScriptVersionGetResponseMetadataSourceAPI, ScriptVersionGetResponseMetadataSourceWrangler, ScriptVersionGetResponseMetadataSourceTerraform, ScriptVersionGetResponseMetadataSourceDash, ScriptVersionGetResponseMetadataSourceCfCli, ScriptVersionGetResponseMetadataSourceDashTemplate, ScriptVersionGetResponseMetadataSourceIntegration, ScriptVersionGetResponseMetadataSourceQuickEditor, ScriptVersionGetResponseMetadataSourcePlayground, ScriptVersionGetResponseMetadataSourceWorkersci:
		return true
	}
	return false
}

type ScriptVersionNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// JSON-encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[ScriptVersionNewParamsMetadata] `json:"metadata" api:"required"`
	// When set to "strict", the upload will fail if any `inherit` type bindings cannot
	// be resolved against the previous version of the Worker. Without this,
	// unresolvable inherit bindings are silently dropped.
	BindingsInherit param.Field[ScriptVersionNewParamsBindingsInherit] `query:"bindings_inherit"`
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

// URLQuery serializes [ScriptVersionNewParams]'s query parameters as `url.Values`.
func (r ScriptVersionNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// JSON-encoded metadata about the uploaded parts and Worker configuration.
type ScriptVersionNewParamsMetadata struct {
	// Name of the uploaded file that contains the main module (e.g. the file exporting
	// a `fetch` handler). Indicates a `module syntax` Worker, which is required for
	// Version Upload.
	MainModule  param.Field[string]                                    `json:"main_module" api:"required"`
	Annotations param.Field[ScriptVersionNewParamsMetadataAnnotations] `json:"annotations"`
	// List of bindings attached to a Worker. You can find more about bindings on our
	// docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Bindings param.Field[[]ScriptVersionNewParamsMetadataBindingUnion] `json:"bindings"`
	// Global CacheW configuration for the Worker. When caching is on, the platform
	// provisions a `cloudflare.app` zone for the Worker. A `type: worker` entry in the
	// `exports` map can override this value for a single entrypoint.
	CacheOptions param.Field[ScriptVersionNewParamsMetadataCacheOptions] `json:"cache_options"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Declarative exports for this version. Worker entrypoint entries (`type: worker`)
	// carry cache configuration for that entrypoint.
	Exports param.Field[map[string]ScriptVersionNewParamsMetadataExportsUnion] `json:"exports"`
	// List of binding types to keep from previous_upload.
	KeepBindings param.Field[[]string] `json:"keep_bindings"`
	// The list of npm packages that were installed and used when this Worker version
	// was built.
	PackageDependencies param.Field[[]ScriptVersionNewParamsMetadataPackageDependency] `json:"package_dependencies"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[ScriptVersionNewParamsMetadataUsageModel] `json:"usage_model"`
}

func (r ScriptVersionNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptVersionNewParamsMetadataAnnotations struct {
	// Associated alias for a version.
	WorkersAlias param.Field[string] `json:"workers/alias"`
	// Full commit SHA associated with the version, detected from the CI environment.
	// Maximum 64 bytes.
	WorkersCommitSha param.Field[string] `json:"workers/commit_sha"`
	// Human-readable message about the version. Truncated to 1000 bytes if longer.
	WorkersMessage param.Field[string] `json:"workers/message"`
	// Number of the pull or merge request associated with the version, detected from
	// the CI environment. Maximum 20 bytes.
	WorkersPullRequestNumber param.Field[string] `json:"workers/pull_request_number"`
	// Title of the pull or merge request associated with the version, detected from
	// the CI environment. Maximum 512 bytes.
	WorkersPullRequestTitle param.Field[string] `json:"workers/pull_request_title"`
	// URL of the pull or merge request associated with the version, detected from the
	// CI environment. Maximum 512 bytes.
	WorkersPullRequestURL param.Field[string] `json:"workers/pull_request_url"`
	// URL of the source repository the version was built from, detected from the CI
	// environment. Maximum 512 bytes.
	WorkersRepositoryURL param.Field[string] `json:"workers/repository_url"`
	// User-provided identifier for the version. Maximum 100 bytes.
	WorkersTag param.Field[string] `json:"workers/tag"`
}

func (r ScriptVersionNewParamsMetadataAnnotations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources.
type ScriptVersionNewParamsMetadataBinding struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsType] `json:"type" api:"required"`
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
	Format param.Field[ScriptVersionNewParamsMetadataBindingsFormat] `json:"format"`
	// Enables Gateway identity for the binding. Requires network_id to be
	// "cf1:network" and cannot be combined with tunnel_id.
	Identity param.Field[ScriptVersionNewParamsMetadataBindingsIdentity] `json:"identity"`
	// Name of the Vectorize index to bind to.
	IndexName param.Field[string] `json:"index_name"`
	// The user-chosen instance name. Must exist at deploy time. The worker can search,
	// chat, update, and manage items/jobs on this instance.
	InstanceName param.Field[string]      `json:"instance_name"`
	Json         param.Field[interface{}] `json:"json"`
	// The
	// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
	// of the R2 bucket.
	Jurisdiction param.Field[ScriptVersionNewParamsMetadataBindingsJurisdiction] `json:"jurisdiction"`
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

func (r ScriptVersionNewParamsMetadataBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBinding) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// A binding to allow the Worker to communicate with resources.
//
// Satisfied by
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAI],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearch],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearchNamespace],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMessaging],
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
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMedia],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificate],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainText],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelines],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueue],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindRatelimit],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2Bucket],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretText],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSendEmail],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindService],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTextBlob],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorize],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadata],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecret],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindFlagship],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKey],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflow],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWasmModule],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCService],
// [workers.ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCNetwork],
// [ScriptVersionNewParamsMetadataBinding].
type ScriptVersionNewParamsMetadataBindingUnion interface {
	implementsScriptVersionNewParamsMetadataBindingUnion()
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAI struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAIType] `json:"type" api:"required"`
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

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearch struct {
	// The user-chosen instance name. Must exist at deploy time. The worker can search,
	// chat, update, and manage items/jobs on this instance.
	InstanceName param.Field[string] `json:"instance_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearchType] `json:"type" api:"required"`
	// The namespace the instance belongs to. Defaults to "default" if omitted.
	// Customers who don't use namespaces can simply omit this field.
	Namespace param.Field[string] `json:"namespace"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearch) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearchType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearchTypeAISearch ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearchType = "ai_search"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearchType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearchTypeAISearch:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearchNamespace struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The user-chosen namespace name. Must exist before deploy -- Wrangler handles
	// auto-creation on deploy failure (R2 bucket pattern). The "default" namespace is
	// auto-created by config-api for new accounts. Grants full access (CRUD + search +
	// chat) to all instances within the namespace.
	Namespace param.Field[string] `json:"namespace" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearchNamespaceType] `json:"type" api:"required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearchNamespace) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearchNamespace) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearchNamespaceType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearchNamespaceTypeAISearchNamespace ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearchNamespaceType = "ai_search_namespace"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearchNamespaceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAISearchNamespaceTypeAISearchNamespace:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMessaging struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The Messaging namespace to bind to.
	Namespace param.Field[string] `json:"namespace" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMessagingType] `json:"type" api:"required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMessaging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMessaging) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMessagingType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMessagingTypeMessaging ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMessagingType = "messaging"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMessagingType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMessagingTypeMessaging:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngine struct {
	// The name of the dataset to bind to.
	Dataset param.Field[string] `json:"dataset" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAnalyticsEngineType] `json:"type" api:"required"`
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
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindAssetsType] `json:"type" api:"required"`
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
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindBrowserType] `json:"type" api:"required"`
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
	DatabaseID param.Field[string] `json:"database_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindD1Type] `json:"type" api:"required"`
	// Identifier of the D1 database to bind to.
	//
	// Deprecated: This property has been renamed to `database_id`.
	ID param.Field[string] `json:"id"`
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
	Name param.Field[string] `json:"name" api:"required"`
	// The name of the file containing the data content. Only accepted for
	// `service worker syntax` Workers.
	Part param.Field[string] `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDataBlobType] `json:"type" api:"required"`
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
	Name param.Field[string] `json:"name" api:"required"`
	// The name of the dispatch namespace.
	Namespace param.Field[string] `json:"namespace" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceType] `json:"type" api:"required"`
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
	Params param.Field[[]ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundParam] `json:"params"`
	// Outbound worker.
	Worker param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker] `json:"worker"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundParam struct {
	// Name of the parameter.
	Name param.Field[string] `json:"name" api:"required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDispatchNamespaceOutboundWorker struct {
	// Entrypoint to invoke on the outbound worker.
	Entrypoint param.Field[string] `json:"entrypoint"`
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
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindDurableObjectNamespaceType] `json:"type" api:"required"`
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
	ID param.Field[string] `json:"id" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindHyperdriveType] `json:"type" api:"required"`
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
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindInheritType] `json:"type" api:"required"`
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
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindImagesType] `json:"type" api:"required"`
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
	Json param.Field[interface{}] `json:"json" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindJsonType] `json:"type" api:"required"`
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
	Name param.Field[string] `json:"name" api:"required"`
	// Namespace identifier tag.
	NamespaceID param.Field[string] `json:"namespace_id" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindKVNamespaceType] `json:"type" api:"required"`
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

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMedia struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMediaType] `json:"type" api:"required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMedia) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMedia) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMediaType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMediaTypeMedia ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMediaType = "media"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMediaType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMediaTypeMedia:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificate struct {
	// Identifier of the certificate to bind to.
	CertificateID param.Field[string] `json:"certificate_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindMTLSCertificateType] `json:"type" api:"required"`
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
	Name param.Field[string] `json:"name" api:"required"`
	// The text value to use.
	Text param.Field[string] `json:"text" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPlainTextType] `json:"type" api:"required"`
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
	Name param.Field[string] `json:"name" api:"required"`
	// Name of the Pipeline to bind to.
	Pipeline param.Field[string] `json:"pipeline" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindPipelinesType] `json:"type" api:"required"`
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
	Name param.Field[string] `json:"name" api:"required"`
	// Name of the Queue to bind to.
	QueueName param.Field[string] `json:"queue_name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindQueueType] `json:"type" api:"required"`
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

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindRatelimit struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Identifier of the rate limit namespace to bind to.
	NamespaceID param.Field[string] `json:"namespace_id" api:"required"`
	// The rate limit configuration.
	Simple param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindRatelimitSimple] `json:"simple" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindRatelimitType] `json:"type" api:"required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindRatelimit) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The rate limit configuration.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindRatelimitSimple struct {
	// The limit (requests per period).
	Limit param.Field[float64] `json:"limit" api:"required"`
	// The period in seconds.
	Period param.Field[int64] `json:"period" api:"required"`
	// Duration in seconds to apply the mitigation action after the rate limit is
	// exceeded. Valid values are 0 (disabled), 10, or multiples of 60 up to 86400.
	// Must be greater than or equal to the period when non-zero.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindRatelimitSimple) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindRatelimitType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindRatelimitTypeRatelimit ScriptVersionNewParamsMetadataBindingsWorkersBindingKindRatelimitType = "ratelimit"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindRatelimitType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindRatelimitTypeRatelimit:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2Bucket struct {
	// R2 bucket to bind to.
	BucketName param.Field[string] `json:"bucket_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketType] `json:"type" api:"required"`
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
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdictionEu          ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdiction = "eu"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdictionFedramp     ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdiction = "fedramp"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdictionFedrampHigh ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdiction = "fedramp-high"
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdictionUs          ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdiction = "us"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdiction) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdictionEu, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdictionFedramp, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdictionFedrampHigh, ScriptVersionNewParamsMetadataBindingsWorkersBindingKindR2BucketJurisdictionUs:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretText struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The secret value to use.
	Text param.Field[string] `json:"text" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretTextType] `json:"type" api:"required"`
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
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSendEmailType] `json:"type" api:"required"`
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
	Name param.Field[string] `json:"name" api:"required"`
	// Name of Worker to bind to.
	Service param.Field[string] `json:"service" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindServiceType] `json:"type" api:"required"`
	// Entrypoint to invoke on the target Worker.
	Entrypoint param.Field[string] `json:"entrypoint"`
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
	Name param.Field[string] `json:"name" api:"required"`
	// The name of the file containing the text content. Only accepted for
	// `service worker syntax` Workers.
	Part param.Field[string] `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindTextBlobType] `json:"type" api:"required"`
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
	IndexName param.Field[string] `json:"index_name" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVectorizeType] `json:"type" api:"required"`
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
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVersionMetadataType] `json:"type" api:"required"`
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
	Name param.Field[string] `json:"name" api:"required"`
	// Name of the secret in the store.
	SecretName param.Field[string] `json:"secret_name" api:"required"`
	// ID of the store containing the secret.
	StoreID param.Field[string] `json:"store_id" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretsStoreSecretType] `json:"type" api:"required"`
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

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindFlagship struct {
	// ID of the Flagship app to bind to for feature flag evaluation.
	AppID param.Field[string] `json:"app_id" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindFlagshipType] `json:"type" api:"required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindFlagship) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindFlagship) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindFlagshipType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindFlagshipTypeFlagship ScriptVersionNewParamsMetadataBindingsWorkersBindingKindFlagshipType = "flagship"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindFlagshipType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindFlagshipTypeFlagship:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKey struct {
	// Algorithm-specific key parameters.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).
	Algorithm param.Field[interface{}] `json:"algorithm" api:"required"`
	// Data format of the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).
	Format param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyFormat] `json:"format" api:"required"`
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyType] `json:"type" api:"required"`
	// Allowed operations with the key.
	// [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).
	Usages param.Field[[]ScriptVersionNewParamsMetadataBindingsWorkersBindingKindSecretKeyUsage] `json:"usages" api:"required"`
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
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWorkflowType] `json:"type" api:"required"`
	// Name of the Workflow to bind to.
	WorkflowName param.Field[string] `json:"workflow_name" api:"required"`
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
	Name param.Field[string] `json:"name" api:"required"`
	// The name of the file containing the WebAssembly module content. Only accepted
	// for `service worker syntax` Workers.
	Part param.Field[string] `json:"part" api:"required"`
	// The kind of resource that the binding provides.
	//
	// Deprecated: deprecated
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindWasmModuleType] `json:"type" api:"required"`
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

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCService struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// Identifier of the VPC service to bind to.
	ServiceID param.Field[string] `json:"service_id" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCServiceType] `json:"type" api:"required"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCService) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCServiceType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCServiceTypeVPCService ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCServiceType = "vpc_service"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCServiceType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCServiceTypeVPCService:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCNetwork struct {
	// A JavaScript variable name for the binding.
	Name param.Field[string] `json:"name" api:"required"`
	// The kind of resource that the binding provides.
	Type param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCNetworkType] `json:"type" api:"required"`
	// Enables Gateway identity for the binding. Requires network_id to be
	// "cf1:network" and cannot be combined with tunnel_id.
	Identity param.Field[ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCNetworkIdentity] `json:"identity"`
	// Identifier of the network to bind to. Only "cf1:network" is currently supported.
	// Mutually exclusive with tunnel_id.
	NetworkID param.Field[string] `json:"network_id"`
	// UUID of the Cloudflare Tunnel to bind to. Mutually exclusive with network_id.
	TunnelID param.Field[string] `json:"tunnel_id"`
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCNetwork) implementsScriptVersionNewParamsMetadataBindingUnion() {
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCNetworkType string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCNetworkTypeVPCNetwork ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCNetworkType = "vpc_network"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCNetworkType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCNetworkTypeVPCNetwork:
		return true
	}
	return false
}

// Enables Gateway identity for the binding. Requires network_id to be
// "cf1:network" and cannot be combined with tunnel_id.
type ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCNetworkIdentity string

const (
	ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCNetworkIdentityRuntimeEmailAlpha ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCNetworkIdentity = "runtime-email-alpha"
)

func (r ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCNetworkIdentity) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsWorkersBindingKindVPCNetworkIdentityRuntimeEmailAlpha:
		return true
	}
	return false
}

// The kind of resource that the binding provides.
type ScriptVersionNewParamsMetadataBindingsType string

const (
	ScriptVersionNewParamsMetadataBindingsTypeAI                     ScriptVersionNewParamsMetadataBindingsType = "ai"
	ScriptVersionNewParamsMetadataBindingsTypeAISearch               ScriptVersionNewParamsMetadataBindingsType = "ai_search"
	ScriptVersionNewParamsMetadataBindingsTypeAISearchNamespace      ScriptVersionNewParamsMetadataBindingsType = "ai_search_namespace"
	ScriptVersionNewParamsMetadataBindingsTypeMessaging              ScriptVersionNewParamsMetadataBindingsType = "messaging"
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
	ScriptVersionNewParamsMetadataBindingsTypeMedia                  ScriptVersionNewParamsMetadataBindingsType = "media"
	ScriptVersionNewParamsMetadataBindingsTypeMTLSCertificate        ScriptVersionNewParamsMetadataBindingsType = "mtls_certificate"
	ScriptVersionNewParamsMetadataBindingsTypePlainText              ScriptVersionNewParamsMetadataBindingsType = "plain_text"
	ScriptVersionNewParamsMetadataBindingsTypePipelines              ScriptVersionNewParamsMetadataBindingsType = "pipelines"
	ScriptVersionNewParamsMetadataBindingsTypeQueue                  ScriptVersionNewParamsMetadataBindingsType = "queue"
	ScriptVersionNewParamsMetadataBindingsTypeRatelimit              ScriptVersionNewParamsMetadataBindingsType = "ratelimit"
	ScriptVersionNewParamsMetadataBindingsTypeR2Bucket               ScriptVersionNewParamsMetadataBindingsType = "r2_bucket"
	ScriptVersionNewParamsMetadataBindingsTypeSecretText             ScriptVersionNewParamsMetadataBindingsType = "secret_text"
	ScriptVersionNewParamsMetadataBindingsTypeSendEmail              ScriptVersionNewParamsMetadataBindingsType = "send_email"
	ScriptVersionNewParamsMetadataBindingsTypeService                ScriptVersionNewParamsMetadataBindingsType = "service"
	ScriptVersionNewParamsMetadataBindingsTypeTextBlob               ScriptVersionNewParamsMetadataBindingsType = "text_blob"
	ScriptVersionNewParamsMetadataBindingsTypeVectorize              ScriptVersionNewParamsMetadataBindingsType = "vectorize"
	ScriptVersionNewParamsMetadataBindingsTypeVersionMetadata        ScriptVersionNewParamsMetadataBindingsType = "version_metadata"
	ScriptVersionNewParamsMetadataBindingsTypeSecretsStoreSecret     ScriptVersionNewParamsMetadataBindingsType = "secrets_store_secret"
	ScriptVersionNewParamsMetadataBindingsTypeFlagship               ScriptVersionNewParamsMetadataBindingsType = "flagship"
	ScriptVersionNewParamsMetadataBindingsTypeSecretKey              ScriptVersionNewParamsMetadataBindingsType = "secret_key"
	ScriptVersionNewParamsMetadataBindingsTypeWorkflow               ScriptVersionNewParamsMetadataBindingsType = "workflow"
	ScriptVersionNewParamsMetadataBindingsTypeWasmModule             ScriptVersionNewParamsMetadataBindingsType = "wasm_module"
	ScriptVersionNewParamsMetadataBindingsTypeVPCService             ScriptVersionNewParamsMetadataBindingsType = "vpc_service"
	ScriptVersionNewParamsMetadataBindingsTypeVPCNetwork             ScriptVersionNewParamsMetadataBindingsType = "vpc_network"
)

func (r ScriptVersionNewParamsMetadataBindingsType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsTypeAI, ScriptVersionNewParamsMetadataBindingsTypeAISearch, ScriptVersionNewParamsMetadataBindingsTypeAISearchNamespace, ScriptVersionNewParamsMetadataBindingsTypeMessaging, ScriptVersionNewParamsMetadataBindingsTypeAnalyticsEngine, ScriptVersionNewParamsMetadataBindingsTypeAssets, ScriptVersionNewParamsMetadataBindingsTypeBrowser, ScriptVersionNewParamsMetadataBindingsTypeD1, ScriptVersionNewParamsMetadataBindingsTypeDataBlob, ScriptVersionNewParamsMetadataBindingsTypeDispatchNamespace, ScriptVersionNewParamsMetadataBindingsTypeDurableObjectNamespace, ScriptVersionNewParamsMetadataBindingsTypeHyperdrive, ScriptVersionNewParamsMetadataBindingsTypeInherit, ScriptVersionNewParamsMetadataBindingsTypeImages, ScriptVersionNewParamsMetadataBindingsTypeJson, ScriptVersionNewParamsMetadataBindingsTypeKVNamespace, ScriptVersionNewParamsMetadataBindingsTypeMedia, ScriptVersionNewParamsMetadataBindingsTypeMTLSCertificate, ScriptVersionNewParamsMetadataBindingsTypePlainText, ScriptVersionNewParamsMetadataBindingsTypePipelines, ScriptVersionNewParamsMetadataBindingsTypeQueue, ScriptVersionNewParamsMetadataBindingsTypeRatelimit, ScriptVersionNewParamsMetadataBindingsTypeR2Bucket, ScriptVersionNewParamsMetadataBindingsTypeSecretText, ScriptVersionNewParamsMetadataBindingsTypeSendEmail, ScriptVersionNewParamsMetadataBindingsTypeService, ScriptVersionNewParamsMetadataBindingsTypeTextBlob, ScriptVersionNewParamsMetadataBindingsTypeVectorize, ScriptVersionNewParamsMetadataBindingsTypeVersionMetadata, ScriptVersionNewParamsMetadataBindingsTypeSecretsStoreSecret, ScriptVersionNewParamsMetadataBindingsTypeFlagship, ScriptVersionNewParamsMetadataBindingsTypeSecretKey, ScriptVersionNewParamsMetadataBindingsTypeWorkflow, ScriptVersionNewParamsMetadataBindingsTypeWasmModule, ScriptVersionNewParamsMetadataBindingsTypeVPCService, ScriptVersionNewParamsMetadataBindingsTypeVPCNetwork:
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

// Enables Gateway identity for the binding. Requires network_id to be
// "cf1:network" and cannot be combined with tunnel_id.
type ScriptVersionNewParamsMetadataBindingsIdentity string

const (
	ScriptVersionNewParamsMetadataBindingsIdentityRuntimeEmailAlpha ScriptVersionNewParamsMetadataBindingsIdentity = "runtime-email-alpha"
)

func (r ScriptVersionNewParamsMetadataBindingsIdentity) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsIdentityRuntimeEmailAlpha:
		return true
	}
	return false
}

// The
// [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions)
// of the R2 bucket.
type ScriptVersionNewParamsMetadataBindingsJurisdiction string

const (
	ScriptVersionNewParamsMetadataBindingsJurisdictionEu          ScriptVersionNewParamsMetadataBindingsJurisdiction = "eu"
	ScriptVersionNewParamsMetadataBindingsJurisdictionFedramp     ScriptVersionNewParamsMetadataBindingsJurisdiction = "fedramp"
	ScriptVersionNewParamsMetadataBindingsJurisdictionFedrampHigh ScriptVersionNewParamsMetadataBindingsJurisdiction = "fedramp-high"
	ScriptVersionNewParamsMetadataBindingsJurisdictionUs          ScriptVersionNewParamsMetadataBindingsJurisdiction = "us"
)

func (r ScriptVersionNewParamsMetadataBindingsJurisdiction) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataBindingsJurisdictionEu, ScriptVersionNewParamsMetadataBindingsJurisdictionFedramp, ScriptVersionNewParamsMetadataBindingsJurisdictionFedrampHigh, ScriptVersionNewParamsMetadataBindingsJurisdictionUs:
		return true
	}
	return false
}

// Global CacheW configuration for the Worker. When caching is on, the platform
// provisions a `cloudflare.app` zone for the Worker. A `type: worker` entry in the
// `exports` map can override this value for a single entrypoint.
type ScriptVersionNewParamsMetadataCacheOptions struct {
	// Whether caching is enabled for this Worker.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Whether cached responses are shared across Worker version uploads. This is
	// independent of `enabled`. It can stay true while caching is off, so the
	// preference survives turning caching off and back on.
	CrossVersionCache param.Field[bool] `json:"cross_version_cache"`
}

func (r ScriptVersionNewParamsMetadataCacheOptions) MarshalJSON() (data []byte, err error) {
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
type ScriptVersionNewParamsMetadataExports struct {
	// Marks this entry as a Worker entrypoint export.
	Type  param.Field[ScriptVersionNewParamsMetadataExportsType] `json:"type" api:"required"`
	Cache param.Field[interface{}]                               `json:"cache"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container param.Field[string] `json:"container"`
	// The destination class name. Must differ from the source class (the map key) and
	// must be declared as a live (`created`) entry in the same `exports` map.
	// Write-only: never present in GET responses.
	RenamedTo param.Field[string] `json:"renamed_to"`
	// Live export. May be omitted; defaults to `created`.
	State param.Field[ScriptVersionNewParamsMetadataExportsState] `json:"state"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage param.Field[ScriptVersionNewParamsMetadataExportsStorage] `json:"storage"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom param.Field[string] `json:"transfer_from"`
	// The destination script name. Must be in the same account and the same
	// dispatch-namespace context (or both non-dispatch). Cross-dispatch-namespace
	// transfers are rejected. Write-only: never present in GET responses.
	TransferredTo param.Field[string] `json:"transferred_to"`
}

func (r ScriptVersionNewParamsMetadataExports) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataExports) implementsScriptVersionNewParamsMetadataExportsUnion() {
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
// Satisfied by [workers.ScriptVersionNewParamsMetadataExportsWorkersWorkerExport],
// [workers.ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExport],
// [workers.ScriptVersionNewParamsMetadataExportsWorkersDurableObjectDeletedExport],
// [workers.ScriptVersionNewParamsMetadataExportsWorkersDurableObjectRenamedExport],
// [workers.ScriptVersionNewParamsMetadataExportsWorkersDurableObjectTransferredExport],
// [workers.ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExport],
// [ScriptVersionNewParamsMetadataExports].
type ScriptVersionNewParamsMetadataExportsUnion interface {
	implementsScriptVersionNewParamsMetadataExportsUnion()
}

// A named Worker entrypoint export (`type: worker`). Worker entrypoints are always
// live (`state: created`) and carry no storage or lifecycle fields. The optional
// `cache` block overrides the Worker's global `cache_options.enabled` for this
// entrypoint.
type ScriptVersionNewParamsMetadataExportsWorkersWorkerExport struct {
	// Marks this entry as a Worker entrypoint export.
	Type param.Field[ScriptVersionNewParamsMetadataExportsWorkersWorkerExportType] `json:"type" api:"required"`
	// Cache override for this entrypoint. Overrides the Worker's global
	// `cache_options.enabled` for this entrypoint only.
	Cache param.Field[ScriptVersionNewParamsMetadataExportsWorkersWorkerExportCache] `json:"cache"`
	// Live export. May be omitted; defaults to `created`.
	State param.Field[ScriptVersionNewParamsMetadataExportsWorkersWorkerExportState] `json:"state"`
}

func (r ScriptVersionNewParamsMetadataExportsWorkersWorkerExport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataExportsWorkersWorkerExport) implementsScriptVersionNewParamsMetadataExportsUnion() {
}

// Marks this entry as a Worker entrypoint export.
type ScriptVersionNewParamsMetadataExportsWorkersWorkerExportType string

const (
	ScriptVersionNewParamsMetadataExportsWorkersWorkerExportTypeWorker ScriptVersionNewParamsMetadataExportsWorkersWorkerExportType = "worker"
)

func (r ScriptVersionNewParamsMetadataExportsWorkersWorkerExportType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataExportsWorkersWorkerExportTypeWorker:
		return true
	}
	return false
}

// Cache override for this entrypoint. Overrides the Worker's global
// `cache_options.enabled` for this entrypoint only.
type ScriptVersionNewParamsMetadataExportsWorkersWorkerExportCache struct {
	// Whether caching is enabled for this entrypoint.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
}

func (r ScriptVersionNewParamsMetadataExportsWorkersWorkerExportCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Live export. May be omitted; defaults to `created`.
type ScriptVersionNewParamsMetadataExportsWorkersWorkerExportState string

const (
	ScriptVersionNewParamsMetadataExportsWorkersWorkerExportStateCreated ScriptVersionNewParamsMetadataExportsWorkersWorkerExportState = "created"
)

func (r ScriptVersionNewParamsMetadataExportsWorkersWorkerExportState) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataExportsWorkersWorkerExportStateCreated:
		return true
	}
	return false
}

// A live Durable Object export (`state: created`, the default). The platform
// auto-provisions the namespace on first deploy, matches it on subsequent deploys,
// and never mutates or deletes it as a side effect of a code-only change.
// `storage` is required; `renamed_to`, `transferred_to` and `transfer_from` are
// not allowed on a live entry.
type ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExport struct {
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage param.Field[ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportStorage] `json:"storage" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportType] `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object. When set, the namespace is container-enabled. Valid
	// only on live entries.
	Container param.Field[string] `json:"container"`
	// Live export. May be omitted; defaults to `created`.
	State param.Field[ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportState] `json:"state"`
}

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExport) implementsScriptVersionNewParamsMetadataExportsUnion() {
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportStorage string

const (
	ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportStorageSqlite   ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportStorage = "sqlite"
	ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportStorageLegacyKV ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportStorage = "legacy-kv"
)

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportStorage) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportStorageSqlite, ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportStorageLegacyKV:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportType string

const (
	ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportTypeDurableObject ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportType = "durable-object"
)

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportTypeDurableObject:
		return true
	}
	return false
}

// Live export. May be omitted; defaults to `created`.
type ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportState string

const (
	ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportStateCreated ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportState = "created"
)

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportState) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExportStateCreated:
		return true
	}
	return false
}

// A `deleted` tombstone: retires the provisioned namespace for this class and all
// of its data. The class must be absent from the uploaded code and no other Worker
// in the account may bind to the namespace, otherwise the deploy is rejected. No
// other fields are allowed. Deletion is irreversible.
type ScriptVersionNewParamsMetadataExportsWorkersDurableObjectDeletedExport struct {
	// Tombstone that deletes the namespace.
	State param.Field[ScriptVersionNewParamsMetadataExportsWorkersDurableObjectDeletedExportState] `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[ScriptVersionNewParamsMetadataExportsWorkersDurableObjectDeletedExportType] `json:"type" api:"required"`
}

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectDeletedExport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectDeletedExport) implementsScriptVersionNewParamsMetadataExportsUnion() {
}

// Tombstone that deletes the namespace.
type ScriptVersionNewParamsMetadataExportsWorkersDurableObjectDeletedExportState string

const (
	ScriptVersionNewParamsMetadataExportsWorkersDurableObjectDeletedExportStateDeleted ScriptVersionNewParamsMetadataExportsWorkersDurableObjectDeletedExportState = "deleted"
)

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectDeletedExportState) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataExportsWorkersDurableObjectDeletedExportStateDeleted:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptVersionNewParamsMetadataExportsWorkersDurableObjectDeletedExportType string

const (
	ScriptVersionNewParamsMetadataExportsWorkersDurableObjectDeletedExportTypeDurableObject ScriptVersionNewParamsMetadataExportsWorkersDurableObjectDeletedExportType = "durable-object"
)

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectDeletedExportType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataExportsWorkersDurableObjectDeletedExportTypeDurableObject:
		return true
	}
	return false
}

// A `renamed` tombstone: rewrites the provisioned namespace's class name from this
// map key to `renamed_to`. The source class may stay in code during the rollout
// window (an info notice is emitted). `storage`, `transferred_to` and
// `transfer_from` are not allowed.
type ScriptVersionNewParamsMetadataExportsWorkersDurableObjectRenamedExport struct {
	// The destination class name. Must differ from the source class (the map key) and
	// must be declared as a live (`created`) entry in the same `exports` map.
	// Write-only: never present in GET responses.
	RenamedTo param.Field[string] `json:"renamed_to" api:"required"`
	// Tombstone that renames the namespace's class.
	State param.Field[ScriptVersionNewParamsMetadataExportsWorkersDurableObjectRenamedExportState] `json:"state" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[ScriptVersionNewParamsMetadataExportsWorkersDurableObjectRenamedExportType] `json:"type" api:"required"`
}

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectRenamedExport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectRenamedExport) implementsScriptVersionNewParamsMetadataExportsUnion() {
}

// Tombstone that renames the namespace's class.
type ScriptVersionNewParamsMetadataExportsWorkersDurableObjectRenamedExportState string

const (
	ScriptVersionNewParamsMetadataExportsWorkersDurableObjectRenamedExportStateRenamed ScriptVersionNewParamsMetadataExportsWorkersDurableObjectRenamedExportState = "renamed"
)

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectRenamedExportState) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataExportsWorkersDurableObjectRenamedExportStateRenamed:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptVersionNewParamsMetadataExportsWorkersDurableObjectRenamedExportType string

const (
	ScriptVersionNewParamsMetadataExportsWorkersDurableObjectRenamedExportTypeDurableObject ScriptVersionNewParamsMetadataExportsWorkersDurableObjectRenamedExportType = "durable-object"
)

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectRenamedExportType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataExportsWorkersDurableObjectRenamedExportTypeDurableObject:
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
type ScriptVersionNewParamsMetadataExportsWorkersDurableObjectTransferredExport struct {
	// Tombstone that transfers the namespace to another script.
	State param.Field[ScriptVersionNewParamsMetadataExportsWorkersDurableObjectTransferredExportState] `json:"state" api:"required"`
	// The destination script name. Must be in the same account and the same
	// dispatch-namespace context (or both non-dispatch). Cross-dispatch-namespace
	// transfers are rejected. Write-only: never present in GET responses.
	TransferredTo param.Field[string] `json:"transferred_to" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[ScriptVersionNewParamsMetadataExportsWorkersDurableObjectTransferredExportType] `json:"type" api:"required"`
}

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectTransferredExport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectTransferredExport) implementsScriptVersionNewParamsMetadataExportsUnion() {
}

// Tombstone that transfers the namespace to another script.
type ScriptVersionNewParamsMetadataExportsWorkersDurableObjectTransferredExportState string

const (
	ScriptVersionNewParamsMetadataExportsWorkersDurableObjectTransferredExportStateTransferred ScriptVersionNewParamsMetadataExportsWorkersDurableObjectTransferredExportState = "transferred"
)

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectTransferredExportState) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataExportsWorkersDurableObjectTransferredExportStateTransferred:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptVersionNewParamsMetadataExportsWorkersDurableObjectTransferredExportType string

const (
	ScriptVersionNewParamsMetadataExportsWorkersDurableObjectTransferredExportTypeDurableObject ScriptVersionNewParamsMetadataExportsWorkersDurableObjectTransferredExportType = "durable-object"
)

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectTransferredExportType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataExportsWorkersDurableObjectTransferredExportTypeDurableObject:
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
type ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExport struct {
	// Target side of a two-phase transfer.
	State param.Field[ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportState] `json:"state" api:"required"`
	// Durable Object storage backend. `sqlite` is the recommended (and only) backend
	// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
	// already exists as KV-backed; the `exports` flow never provisions a new
	// `legacy-kv` namespace.
	Storage param.Field[ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStorage] `json:"storage" api:"required"`
	// The source script name to receive the namespace from. Must be in the same
	// account and dispatch-namespace context. Present on reads for
	// `expecting-transfer` entries.
	TransferFrom param.Field[string] `json:"transfer_from" api:"required"`
	// Marks this entry as a Durable Object export.
	Type param.Field[ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportType] `json:"type" api:"required"`
	// Name of the container (declared in the upload's `metadata.containers`) that
	// backs this Durable Object once the transfer settles. Valid only on live entries.
	Container param.Field[string] `json:"container"`
}

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExport) implementsScriptVersionNewParamsMetadataExportsUnion() {
}

// Target side of a two-phase transfer.
type ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportState string

const (
	ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStateExpectingTransfer ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportState = "expecting-transfer"
)

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportState) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStateExpectingTransfer:
		return true
	}
	return false
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStorage string

const (
	ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStorageSqlite   ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStorage = "sqlite"
	ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStorageLegacyKV ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStorage = "legacy-kv"
)

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStorage) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStorageSqlite, ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportStorageLegacyKV:
		return true
	}
	return false
}

// Marks this entry as a Durable Object export.
type ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportType string

const (
	ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportTypeDurableObject ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportType = "durable-object"
)

func (r ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataExportsWorkersDurableObjectExpectingTransferExportTypeDurableObject:
		return true
	}
	return false
}

// Marks this entry as a Worker entrypoint export.
type ScriptVersionNewParamsMetadataExportsType string

const (
	ScriptVersionNewParamsMetadataExportsTypeWorker        ScriptVersionNewParamsMetadataExportsType = "worker"
	ScriptVersionNewParamsMetadataExportsTypeDurableObject ScriptVersionNewParamsMetadataExportsType = "durable-object"
)

func (r ScriptVersionNewParamsMetadataExportsType) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataExportsTypeWorker, ScriptVersionNewParamsMetadataExportsTypeDurableObject:
		return true
	}
	return false
}

// Live export. May be omitted; defaults to `created`.
type ScriptVersionNewParamsMetadataExportsState string

const (
	ScriptVersionNewParamsMetadataExportsStateCreated           ScriptVersionNewParamsMetadataExportsState = "created"
	ScriptVersionNewParamsMetadataExportsStateDeleted           ScriptVersionNewParamsMetadataExportsState = "deleted"
	ScriptVersionNewParamsMetadataExportsStateRenamed           ScriptVersionNewParamsMetadataExportsState = "renamed"
	ScriptVersionNewParamsMetadataExportsStateTransferred       ScriptVersionNewParamsMetadataExportsState = "transferred"
	ScriptVersionNewParamsMetadataExportsStateExpectingTransfer ScriptVersionNewParamsMetadataExportsState = "expecting-transfer"
)

func (r ScriptVersionNewParamsMetadataExportsState) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataExportsStateCreated, ScriptVersionNewParamsMetadataExportsStateDeleted, ScriptVersionNewParamsMetadataExportsStateRenamed, ScriptVersionNewParamsMetadataExportsStateTransferred, ScriptVersionNewParamsMetadataExportsStateExpectingTransfer:
		return true
	}
	return false
}

// Durable Object storage backend. `sqlite` is the recommended (and only) backend
// for new namespaces. `legacy-kv` is accepted only for a class whose namespace
// already exists as KV-backed; the `exports` flow never provisions a new
// `legacy-kv` namespace.
type ScriptVersionNewParamsMetadataExportsStorage string

const (
	ScriptVersionNewParamsMetadataExportsStorageSqlite   ScriptVersionNewParamsMetadataExportsStorage = "sqlite"
	ScriptVersionNewParamsMetadataExportsStorageLegacyKV ScriptVersionNewParamsMetadataExportsStorage = "legacy-kv"
)

func (r ScriptVersionNewParamsMetadataExportsStorage) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsMetadataExportsStorageSqlite, ScriptVersionNewParamsMetadataExportsStorageLegacyKV:
		return true
	}
	return false
}

type ScriptVersionNewParamsMetadataPackageDependency struct {
	// The exact version that was resolved and installed by the package manager.
	InstalledVersion param.Field[string] `json:"installedVersion" api:"required"`
	// The npm package name.
	Name param.Field[string] `json:"name" api:"required"`
	// The version constraint as written in package.json.
	PackageJsonVersion param.Field[string] `json:"packageJsonVersion" api:"required"`
}

func (r ScriptVersionNewParamsMetadataPackageDependency) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

// When set to "strict", the upload will fail if any `inherit` type bindings cannot
// be resolved against the previous version of the Worker. Without this,
// unresolvable inherit bindings are silently dropped.
type ScriptVersionNewParamsBindingsInherit string

const (
	ScriptVersionNewParamsBindingsInheritStrict ScriptVersionNewParamsBindingsInherit = "strict"
)

func (r ScriptVersionNewParamsBindingsInherit) IsKnown() bool {
	switch r {
	case ScriptVersionNewParamsBindingsInheritStrict:
		return true
	}
	return false
}

type ScriptVersionNewResponseEnvelope struct {
	Errors   []ScriptVersionNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ScriptVersionNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ScriptVersionNewResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success ScriptVersionNewResponseEnvelopeSuccess `json:"success" api:"required"`
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
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
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
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
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
	AccountID param.Field[string] `path:"account_id" api:"required"`
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
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ScriptVersionGetResponseEnvelope struct {
	Errors   []ScriptVersionGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ScriptVersionGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ScriptVersionGetResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success ScriptVersionGetResponseEnvelopeSuccess `json:"success" api:"required"`
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
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
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
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
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
