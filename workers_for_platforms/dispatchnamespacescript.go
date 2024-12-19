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
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/cloudflare/cloudflare-go/v4/workers"
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a worker from a Workers for Platforms namespace. This call has no
// response body on a successful delete.
func (r *DispatchNamespaceScriptService) Delete(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

// Fetch information about a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptGetParams, opts ...option.RequestOption) (res *Script, err error) {
	var env DispatchNamespaceScriptGetResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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
	TailConsumers []workers.ConsumerScript `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel DispatchNamespaceScriptUpdateResponseUsageModel `json:"usage_model"`
	JSON       dispatchNamespaceScriptUpdateResponseJSON       `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseJSON contains the JSON metadata for the
// struct [DispatchNamespaceScriptUpdateResponse]
type dispatchNamespaceScriptUpdateResponseJSON struct {
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

func (r *DispatchNamespaceScriptUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Usage model for the Worker invocations.
type DispatchNamespaceScriptUpdateResponseUsageModel string

const (
	DispatchNamespaceScriptUpdateResponseUsageModelBundled DispatchNamespaceScriptUpdateResponseUsageModel = "bundled"
	DispatchNamespaceScriptUpdateResponseUsageModelUnbound DispatchNamespaceScriptUpdateResponseUsageModel = "unbound"
)

func (r DispatchNamespaceScriptUpdateResponseUsageModel) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateResponseUsageModelBundled, DispatchNamespaceScriptUpdateResponseUsageModelUnbound:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParams struct {
	// Identifier
	AccountID param.Field[string]                          `path:"account_id,required"`
	Body      DispatchNamespaceScriptUpdateParamsBodyUnion `json:"body,required"`
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

type DispatchNamespaceScriptUpdateParamsBody struct {
	// Rollback message to be associated with this deployment. Only parsed when query
	// param `"rollback_to"` is present.
	Message  param.Field[string]      `json:"message"`
	Metadata param.Field[interface{}] `json:"metadata"`
}

func (r DispatchNamespaceScriptUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsBody) implementsWorkersForPlatformsDispatchNamespaceScriptUpdateParamsBodyUnion() {
}

// Satisfied by
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsBodyMetadata],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsBodyMessage],
// [DispatchNamespaceScriptUpdateParamsBody].
type DispatchNamespaceScriptUpdateParamsBodyUnion interface {
	implementsWorkersForPlatformsDispatchNamespaceScriptUpdateParamsBodyUnion()
}

type DispatchNamespaceScriptUpdateParamsBodyMetadata struct {
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata    param.Field[DispatchNamespaceScriptUpdateParamsBodyMetadataMetadata] `json:"metadata,required"`
	ExtraFields map[string][]io.Reader                                               `json:"-,extras"`
}

func (r DispatchNamespaceScriptUpdateParamsBodyMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsBodyMetadata) implementsWorkersForPlatformsDispatchNamespaceScriptUpdateParamsBodyUnion() {
}

// JSON encoded metadata about the uploaded parts and Worker configuration.
type DispatchNamespaceScriptUpdateParamsBodyMetadataMetadata struct {
	// Configuration for assets within a Worker
	Assets param.Field[DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssets] `json:"assets"`
	// List of bindings available to the worker.
	Bindings param.Field[[]DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataBinding] `json:"bindings"`
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
	Migrations param.Field[DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataMigrationsUnion] `json:"migrations"`
	// Observability settings for the Worker.
	Observability param.Field[DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataObservability] `json:"observability"`
	Placement     param.Field[workers.PlacementConfigurationParam]                                  `json:"placement"`
	// List of strings to use as tags for this Worker.
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]workers.ConsumerScriptParam] `json:"tail_consumers"`
	// Usage model for the Worker invocations.
	UsageModel param.Field[DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataUsageModel] `json:"usage_model"`
	// Key-value pairs to use as tags for this version of this Worker.
	VersionTags param.Field[map[string]string] `json:"version_tags"`
}

func (r DispatchNamespaceScriptUpdateParamsBodyMetadataMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker
type DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssets struct {
	// Configuration for assets within a Worker.
	Config param.Field[DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfig] `json:"config"`
	// Token provided upon successful upload of all files from a registered manifest.
	JWT param.Field[string] `json:"jwt"`
}

func (r DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for assets within a Worker.
type DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfig struct {
	// Determines the redirects and rewrites of requests for HTML content.
	HTMLHandling param.Field[DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandling] `json:"html_handling"`
	// Determines the response when a request does not match a static asset, and there
	// is no Worker script.
	NotFoundHandling param.Field[DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandling] `json:"not_found_handling"`
	// When true and the incoming request matches an asset, that will be served instead
	// of invoking the Worker script. When false, requests will always invoke the
	// Worker script.
	ServeDirectly param.Field[bool] `json:"serve_directly"`
}

func (r DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the redirects and rewrites of requests for HTML content.
type DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandling string

const (
	DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandlingAutoTrailingSlash  DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandling = "auto-trailing-slash"
	DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandlingForceTrailingSlash DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandling = "force-trailing-slash"
	DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandlingDropTrailingSlash  DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandling = "drop-trailing-slash"
	DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandlingNone               DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandling = "none"
)

func (r DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandling) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandlingAutoTrailingSlash, DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandlingForceTrailingSlash, DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandlingDropTrailingSlash, DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigHTMLHandlingNone:
		return true
	}
	return false
}

// Determines the response when a request does not match a static asset, and there
// is no Worker script.
type DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandling string

const (
	DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandlingNone                  DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandling = "none"
	DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandling404Page               DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandling = "404-page"
	DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandlingSinglePageApplication DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandling = "single-page-application"
)

func (r DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandling) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandlingNone, DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandling404Page, DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataAssetsConfigNotFoundHandlingSinglePageApplication:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataBinding struct {
	// Name of the binding variable.
	Name param.Field[string] `json:"name"`
	// Type of binding. You can find more about bindings on our docs:
	// https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.
	Type        param.Field[string]    `json:"type"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Migrations to apply for Durable Objects associated with this Worker.
type DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataMigrations struct {
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

func (r DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataMigrations) ImplementsWorkersForPlatformsDispatchNamespaceScriptUpdateParamsBodyMetadataMetadataMigrationsUnion() {
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by [workers.SingleStepMigrationParam],
// [workers_for_platforms.DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataMigrationsWorkersMultipleStepMigrations],
// [DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataMigrations].
type DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataMigrationsUnion interface {
	ImplementsWorkersForPlatformsDispatchNamespaceScriptUpdateParamsBodyMetadataMetadataMigrationsUnion()
}

type DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataMigrationsWorkersMultipleStepMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]workers.MigrationStepParam] `json:"steps"`
}

func (r DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataMigrationsWorkersMultipleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataMigrationsWorkersMultipleStepMigrations) ImplementsWorkersForPlatformsDispatchNamespaceScriptUpdateParamsBodyMetadataMetadataMigrationsUnion() {
}

// Observability settings for the Worker.
type DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataObservability struct {
	// Whether observability is enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
}

func (r DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataObservability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Usage model for the Worker invocations.
type DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataUsageModel string

const (
	DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataUsageModelBundled DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataUsageModel = "bundled"
	DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataUsageModelUnbound DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataUsageModel = "unbound"
)

func (r DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataUsageModel) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataUsageModelBundled, DispatchNamespaceScriptUpdateParamsBodyMetadataMetadataUsageModelUnbound:
		return true
	}
	return false
}

type DispatchNamespaceScriptUpdateParamsBodyMessage struct {
	// Rollback message to be associated with this deployment. Only parsed when query
	// param `"rollback_to"` is present.
	Message param.Field[string] `json:"message"`
}

func (r DispatchNamespaceScriptUpdateParamsBodyMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptUpdateParamsBodyMessage) implementsWorkersForPlatformsDispatchNamespaceScriptUpdateParamsBodyUnion() {
}

type DispatchNamespaceScriptUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceScriptUpdateResponse                `json:"result"`
	JSON    dispatchNamespaceScriptUpdateResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptUpdateResponseEnvelope]
type dispatchNamespaceScriptUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
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
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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

type DispatchNamespaceScriptGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceScriptGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptGetResponseEnvelopeSuccess `json:"success,required"`
	// Details about a worker uploaded to a Workers for Platforms namespace.
	Result Script                                         `json:"result"`
	JSON   dispatchNamespaceScriptGetResponseEnvelopeJSON `json:"-"`
}

// dispatchNamespaceScriptGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptGetResponseEnvelope]
type dispatchNamespaceScriptGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
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
