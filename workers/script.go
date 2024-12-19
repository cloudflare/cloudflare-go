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
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// ScriptService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptService] method instead.
type ScriptService struct {
	Options     []option.RequestOption
	Assets      *ScriptAssetService
	Subdomain   *ScriptSubdomainService
	Schedules   *ScriptScheduleService
	Tail        *ScriptTailService
	Content     *ScriptContentService
	Settings    *ScriptSettingService
	Deployments *ScriptDeploymentService
	Versions    *ScriptVersionService
}

// NewScriptService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScriptService(opts ...option.RequestOption) (r *ScriptService) {
	r = &ScriptService{}
	r.Options = opts
	r.Assets = NewScriptAssetService(opts...)
	r.Subdomain = NewScriptSubdomainService(opts...)
	r.Schedules = NewScriptScheduleService(opts...)
	r.Tail = NewScriptTailService(opts...)
	r.Content = NewScriptContentService(opts...)
	r.Settings = NewScriptSettingService(opts...)
	r.Deployments = NewScriptDeploymentService(opts...)
	r.Versions = NewScriptVersionService(opts...)
	return
}

// Upload a worker module. You can find more about the multipart metadata on our
// docs:
// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/.
func (r *ScriptService) Update(ctx context.Context, scriptName string, params ScriptUpdateParams, opts ...option.RequestOption) (res *ScriptUpdateResponse, err error) {
	var env ScriptUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of uploaded workers.
func (r *ScriptService) List(ctx context.Context, query ScriptListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Script], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts", query.AccountID)
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

// Fetch a list of uploaded workers.
func (r *ScriptService) ListAutoPaging(ctx context.Context, query ScriptListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Script] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete your worker. This call has no response body on a successful delete.
func (r *ScriptService) Delete(ctx context.Context, scriptName string, params ScriptDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

// Fetch raw script content for your worker. Note this is the original script
// content, not JSON encoded.
func (r *ScriptService) Get(ctx context.Context, scriptName string, query ScriptGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "undefined")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Script struct {
	// The id of the script in the Workers system. Usually the script name.
	ID string `json:"id"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Hashed script content, can be used in a If-None-Match header when updating.
	Etag string `json:"etag"`
	// Whether a Worker contains assets.
	HasAssets bool `json:"has_assets"`
	// Whether a Worker contains modules.
	HasModules bool `json:"has_modules"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Specifies the placement mode for the Worker (e.g. 'smart').
	PlacementMode string `json:"placement_mode"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ConsumerScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel ScriptUsageModel `json:"usage_model"`
	JSON       scriptJSON       `json:"-"`
}

// scriptJSON contains the JSON metadata for the struct [Script]
type scriptJSON struct {
	ID            apijson.Field
	CreatedOn     apijson.Field
	Etag          apijson.Field
	HasAssets     apijson.Field
	HasModules    apijson.Field
	Logpush       apijson.Field
	ModifiedOn    apijson.Field
	PlacementMode apijson.Field
	TailConsumers apijson.Field
	UsageModel    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Script) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptJSON) RawJSON() string {
	return r.raw
}

// Usage model for the Worker invocations.
type ScriptUsageModel string

const (
	ScriptUsageModelBundled ScriptUsageModel = "bundled"
	ScriptUsageModelUnbound ScriptUsageModel = "unbound"
)

func (r ScriptUsageModel) IsKnown() bool {
	switch r {
	case ScriptUsageModelBundled, ScriptUsageModelUnbound:
		return true
	}
	return false
}

type ScriptSetting struct {
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Observability settings for the Worker.
	Observability ScriptSettingObservability `json:"observability"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ConsumerScript  `json:"tail_consumers"`
	JSON          scriptSettingJSON `json:"-"`
}

// scriptSettingJSON contains the JSON metadata for the struct [ScriptSetting]
type scriptSettingJSON struct {
	Logpush       apijson.Field
	Observability apijson.Field
	TailConsumers apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptSetting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingJSON) RawJSON() string {
	return r.raw
}

// Observability settings for the Worker.
type ScriptSettingObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled bool `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate float64                        `json:"head_sampling_rate,nullable"`
	JSON             scriptSettingObservabilityJSON `json:"-"`
}

// scriptSettingObservabilityJSON contains the JSON metadata for the struct
// [ScriptSettingObservability]
type scriptSettingObservabilityJSON struct {
	Enabled          apijson.Field
	HeadSamplingRate apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScriptSettingObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingObservabilityJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingParam struct {
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Observability settings for the Worker.
	Observability param.Field[ScriptSettingObservabilityParam] `json:"observability"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]ConsumerScriptParam] `json:"tail_consumers"`
}

func (r ScriptSettingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Observability settings for the Worker.
type ScriptSettingObservabilityParam struct {
	// Whether observability is enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
}

func (r ScriptSettingObservabilityParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptUpdateResponse struct {
	// The id of the script in the Workers system. Usually the script name.
	ID string `json:"id"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Hashed script content, can be used in a If-None-Match header when updating.
	Etag string `json:"etag"`
	// Whether a Worker contains assets.
	HasAssets bool `json:"has_assets"`
	// Whether a Worker contains modules.
	HasModules bool `json:"has_modules"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Specifies the placement mode for the Worker (e.g. 'smart').
	PlacementMode string `json:"placement_mode"`
	StartupTimeMs int64  `json:"startup_time_ms"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ConsumerScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel ScriptUpdateResponseUsageModel `json:"usage_model"`
	JSON       scriptUpdateResponseJSON       `json:"-"`
}

// scriptUpdateResponseJSON contains the JSON metadata for the struct
// [ScriptUpdateResponse]
type scriptUpdateResponseJSON struct {
	ID            apijson.Field
	CreatedOn     apijson.Field
	Etag          apijson.Field
	HasAssets     apijson.Field
	HasModules    apijson.Field
	Logpush       apijson.Field
	ModifiedOn    apijson.Field
	PlacementMode apijson.Field
	StartupTimeMs apijson.Field
	TailConsumers apijson.Field
	UsageModel    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Usage model for the Worker invocations.
type ScriptUpdateResponseUsageModel string

const (
	ScriptUpdateResponseUsageModelBundled ScriptUpdateResponseUsageModel = "bundled"
	ScriptUpdateResponseUsageModelUnbound ScriptUpdateResponseUsageModel = "unbound"
)

func (r ScriptUpdateResponseUsageModel) IsKnown() bool {
	switch r {
	case ScriptUpdateResponseUsageModelBundled, ScriptUpdateResponseUsageModelUnbound:
		return true
	}
	return false
}

type ScriptUpdateParams struct {
	// Identifier
	AccountID param.Field[string]         `path:"account_id,required"`
	Body      ScriptUpdateParamsBodyUnion `json:"body,required"`
	// Rollback to provided deployment based on deployment ID. Request body will only
	// parse a "message" part. You can learn more about deployments
	// [here](https://developers.cloudflare.com/workers/platform/deployments/).
	RollbackTo param.Field[string] `query:"rollback_to"`
}

func (r ScriptUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [ScriptUpdateParams]'s query parameters as `url.Values`.
func (r ScriptUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ScriptUpdateParamsBody struct {
	// Rollback message to be associated with this deployment. Only parsed when query
	// param `"rollback_to"` is present.
	Message  param.Field[string]      `json:"message"`
	Metadata param.Field[interface{}] `json:"metadata"`
}

func (r ScriptUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsBody) implementsWorkersScriptUpdateParamsBodyUnion() {}

// Satisfied by [workers.ScriptUpdateParamsBodyMetadata],
// [workers.ScriptUpdateParamsBodyMessage], [ScriptUpdateParamsBody].
type ScriptUpdateParamsBodyUnion interface {
	implementsWorkersScriptUpdateParamsBodyUnion()
}

type ScriptUpdateParamsBodyMetadata struct {
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata    param.Field[ScriptUpdateParamsBodyMetadataMetadata] `json:"metadata,required"`
	ExtraFields map[string][]io.Reader                              `json:"-,extras"`
}

func (r ScriptUpdateParamsBodyMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsBodyMetadata) implementsWorkersScriptUpdateParamsBodyUnion() {}

// JSON encoded metadata about the uploaded parts and Worker configuration.
type ScriptUpdateParamsBodyMetadataMetadata struct {
	// Configuration for assets within a Worker
	Assets param.Field[ScriptUpdateParamsBodyMetadataMetadataAssets] `json:"assets"`
	// List of bindings available to the worker.
	Bindings param.Field[[]ScriptUpdateParamsBodyMetadataMetadataBinding] `json:"bindings"`
	// Name of the part in the multipart request that contains the script (e.g. the
	// file adding a listener to the `fetch` event). Indicates a
	// `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
	// Date indicating targeted support in the Workers runtime. Backwards incompatible
	// fixes to the runtime following this date will not affect this Worker.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Flags that enable or disable certain features in the Workers runtime. Used to
	// enable upcoming features or opt in or out of specific changes not included in a
	// `compatibility_date`.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Retain assets which exist for a previously uploaded Worker version; used in lieu
	// of providing a completion token.
	KeepAssets param.Field[bool] `json:"keep_assets"`
	// List of binding types to keep from previous_upload.
	KeepBindings param.Field[[]string] `json:"keep_bindings"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[ScriptUpdateParamsBodyMetadataMetadataMigrationsUnion] `json:"migrations"`
	// Observability settings for the Worker.
	Observability param.Field[ScriptUpdateParamsBodyMetadataMetadataObservability] `json:"observability"`
	Placement     param.Field[PlacementConfigurationParam]                         `json:"placement"`
	// List of strings to use as tags for this Worker.
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]ConsumerScriptParam] `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[ScriptUpdateParamsBodyMetadataMetadataUsageModel] `json:"usage_model"`
	// Key-value pairs to use as tags for this version of this Worker.
	VersionTags param.Field[map[string]string] `json:"version_tags"`
}

func (r ScriptUpdateParamsBodyMetadataMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker
type ScriptUpdateParamsBodyMetadataMetadataAssets struct {
	// Configuration for assets within a Worker.
	Config param.Field[ScriptUpdateParamsBodyMetadataMetadataAssetsConfig] `json:"config"`
	// Token provided upon successful upload of all files from a registered manifest.
	JWT param.Field[string] `json:"jwt"`
}

func (r ScriptUpdateParamsBodyMetadataMetadataAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker.
type ScriptUpdateParamsBodyMetadataMetadataAssetsConfig struct {
	// Determines the redirects and rewrites of requests for HTML content.
	HTMLHandling param.Field[ScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandling] `json:"html_handling"`
	// Determines the response when a request does not match a static asset, and there
	// is no Worker script.
	NotFoundHandling param.Field[ScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandling] `json:"not_found_handling"`
	// When true and the incoming request matches an asset, that will be served instead
	// of invoking the Worker script. When false, requests will always invoke the
	// Worker script.
	ServeDirectly param.Field[bool] `json:"serve_directly"`
}

func (r ScriptUpdateParamsBodyMetadataMetadataAssetsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the redirects and rewrites of requests for HTML content.
type ScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandling string

const (
	ScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandlingAutoTrailingSlash  ScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandling = "auto-trailing-slash"
	ScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandlingForceTrailingSlash ScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandling = "force-trailing-slash"
	ScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandlingDropTrailingSlash  ScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandling = "drop-trailing-slash"
	ScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandlingNone               ScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandling = "none"
)

func (r ScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandling) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandlingAutoTrailingSlash, ScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandlingForceTrailingSlash, ScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandlingDropTrailingSlash, ScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandlingNone:
		return true
	}
	return false
}

// Determines the response when a request does not match a static asset, and there
// is no Worker script.
type ScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandling string

const (
	ScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandlingNone                  ScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandling = "none"
	ScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandling404Page               ScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandling = "404-page"
	ScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandlingSinglePageApplication ScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandling = "single-page-application"
)

func (r ScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandling) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandlingNone, ScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandling404Page, ScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandlingSinglePageApplication:
		return true
	}
	return false
}

type ScriptUpdateParamsBodyMetadataMetadataBinding struct {
	// Name of the binding variable.
	Name param.Field[string] `json:"name"`
	// Type of binding. You can find more about bindings on our docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Type        param.Field[string]    `json:"type"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r ScriptUpdateParamsBodyMetadataMetadataBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Migrations to apply for Durable Objects associated with this Worker.
type ScriptUpdateParamsBodyMetadataMetadataMigrations struct {
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

func (r ScriptUpdateParamsBodyMetadataMetadataMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsBodyMetadataMetadataMigrations) implementsWorkersScriptUpdateParamsBodyMetadataMetadataMigrationsUnion() {
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by [workers.SingleStepMigrationParam],
// [workers.ScriptUpdateParamsBodyMetadataMetadataMigrationsWorkersMultipleStepMigrations],
// [ScriptUpdateParamsBodyMetadataMetadataMigrations].
type ScriptUpdateParamsBodyMetadataMetadataMigrationsUnion interface {
	implementsWorkersScriptUpdateParamsBodyMetadataMetadataMigrationsUnion()
}

type ScriptUpdateParamsBodyMetadataMetadataMigrationsWorkersMultipleStepMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]MigrationStepParam] `json:"steps"`
}

func (r ScriptUpdateParamsBodyMetadataMetadataMigrationsWorkersMultipleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsBodyMetadataMetadataMigrationsWorkersMultipleStepMigrations) implementsWorkersScriptUpdateParamsBodyMetadataMetadataMigrationsUnion() {
}

// Observability settings for the Worker.
type ScriptUpdateParamsBodyMetadataMetadataObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
}

func (r ScriptUpdateParamsBodyMetadataMetadataObservability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Usage model for the Worker invocations.
type ScriptUpdateParamsBodyMetadataMetadataUsageModel string

const (
	ScriptUpdateParamsBodyMetadataMetadataUsageModelBundled ScriptUpdateParamsBodyMetadataMetadataUsageModel = "bundled"
	ScriptUpdateParamsBodyMetadataMetadataUsageModelUnbound ScriptUpdateParamsBodyMetadataMetadataUsageModel = "unbound"
)

func (r ScriptUpdateParamsBodyMetadataMetadataUsageModel) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsBodyMetadataMetadataUsageModelBundled, ScriptUpdateParamsBodyMetadataMetadataUsageModelUnbound:
		return true
	}
	return false
}

type ScriptUpdateParamsBodyMessage struct {
	// Rollback message to be associated with this deployment. Only parsed when query
	// param `"rollback_to"` is present.
	Message param.Field[string] `json:"message"`
}

func (r ScriptUpdateParamsBodyMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsBodyMessage) implementsWorkersScriptUpdateParamsBodyUnion() {}

type ScriptUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ScriptUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptUpdateResponse                `json:"result"`
	JSON    scriptUpdateResponseEnvelopeJSON    `json:"-"`
}

// scriptUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptUpdateResponseEnvelope]
type scriptUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptUpdateResponseEnvelopeSuccess bool

const (
	ScriptUpdateResponseEnvelopeSuccessTrue ScriptUpdateResponseEnvelopeSuccess = true
)

func (r ScriptUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// If set to true, delete will not be stopped by associated service binding,
	// durable object, or other binding. Any of these associated bindings/durable
	// objects will be deleted along with the script.
	Force param.Field[bool] `query:"force"`
}

// URLQuery serializes [ScriptDeleteParams]'s query parameters as `url.Values`.
func (r ScriptDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ScriptGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
