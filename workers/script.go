// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// ScriptService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewScriptService] method instead.
type ScriptService struct {
	Options     []option.RequestOption
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
	r.Schedules = NewScriptScheduleService(opts...)
	r.Tail = NewScriptTailService(opts...)
	r.Content = NewScriptContentService(opts...)
	r.Settings = NewScriptSettingService(opts...)
	r.Deployments = NewScriptDeploymentService(opts...)
	r.Versions = NewScriptVersionService(opts...)
	return
}

// Upload a worker module.
func (r *ScriptService) Update(ctx context.Context, scriptName string, params ScriptUpdateParams, opts ...option.RequestOption) (res *Script, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptUpdateResponseEnvelope
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
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

// Fetch raw script content for your worker. Note this is the original script
// content, not JSON encoded.
func (r *ScriptService) Get(ctx context.Context, scriptName string, query ScriptGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "undefined")}, opts...)
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
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Specifies the placement mode for the Worker (e.g. 'smart').
	PlacementMode string `json:"placement_mode"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ConsumerScript `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string     `json:"usage_model"`
	JSON       scriptJSON `json:"-"`
}

// scriptJSON contains the JSON metadata for the struct [Script]
type scriptJSON struct {
	ID            apijson.Field
	CreatedOn     apijson.Field
	Etag          apijson.Field
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

type ScriptSetting struct {
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ConsumerScript  `json:"tail_consumers"`
	JSON          scriptSettingJSON `json:"-"`
}

// scriptSettingJSON contains the JSON metadata for the struct [ScriptSetting]
type scriptSettingJSON struct {
	Logpush       apijson.Field
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

type ScriptSettingParam struct {
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]ConsumerScriptParam] `json:"tail_consumers"`
}

func (r ScriptSettingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScriptUpdateParamsBody struct {
	AnyPartName param.Field[interface{}] `json:"<any part name>,required"`
	Metadata    param.Field[interface{}] `json:"metadata,required"`
	// Rollback message to be associated with this deployment. Only parsed when query
	// param `"rollback_to"` is present.
	Message param.Field[string] `json:"message"`
}

func (r ScriptUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsBody) implementsWorkersScriptUpdateParamsBodyUnion() {}

// Satisfied by [workers.ScriptUpdateParamsBodyObject],
// [workers.ScriptUpdateParamsBodyObject], [ScriptUpdateParamsBody].
type ScriptUpdateParamsBodyUnion interface {
	implementsWorkersScriptUpdateParamsBodyUnion()
}

type ScriptUpdateParamsBodyObject struct {
	// A module comprising a Worker script, often a javascript file. Multiple modules
	// may be provided as separate named parts, but at least one module must be present
	// and referenced in the metadata as `main_module` or `body_part` by part name.
	// Source maps may also be included using the `application/source-map` content
	// type.
	AnyPartName param.Field[[]io.Reader] `json:"<any part name>" format:"binary"`
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[ScriptUpdateParamsBodyObjectMetadata] `json:"metadata"`
}

func (r ScriptUpdateParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsBodyObject) implementsWorkersScriptUpdateParamsBodyUnion() {}

// JSON encoded metadata about the uploaded parts and Worker configuration.
type ScriptUpdateParamsBodyObjectMetadata struct {
	// List of bindings available to the worker.
	Bindings param.Field[[]interface{}] `json:"bindings"`
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
	// List of binding types to keep from previous_upload.
	KeepBindings param.Field[[]string] `json:"keep_bindings"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[ScriptUpdateParamsBodyObjectMetadataMigrationsUnion] `json:"migrations"`
	Placement  param.Field[PlacementConfigurationParam]                         `json:"placement"`
	// List of strings to use as tags for this Worker
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]ConsumerScriptParam] `json:"tail_consumers"`
	// Usage model to apply to invocations.
	UsageModel param.Field[ScriptUpdateParamsBodyObjectMetadataUsageModel] `json:"usage_model"`
	// Key-value pairs to use as tags for this version of this Worker
	VersionTags param.Field[interface{}] `json:"version_tags"`
}

func (r ScriptUpdateParamsBodyObjectMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Migrations to apply for Durable Objects associated with this Worker.
type ScriptUpdateParamsBodyObjectMetadataMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag             param.Field[string]      `json:"old_tag"`
	DeletedClasses     param.Field[interface{}] `json:"deleted_classes,required"`
	NewClasses         param.Field[interface{}] `json:"new_classes,required"`
	RenamedClasses     param.Field[interface{}] `json:"renamed_classes,required"`
	TransferredClasses param.Field[interface{}] `json:"transferred_classes,required"`
	Steps              param.Field[interface{}] `json:"steps,required"`
}

func (r ScriptUpdateParamsBodyObjectMetadataMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptUpdateParamsBodyObjectMetadataMigrations) implementsWorkersScriptUpdateParamsBodyObjectMetadataMigrationsUnion() {
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by [workers.SingleStepMigrationParam],
// [workers.SteppedMigrationParam],
// [ScriptUpdateParamsBodyObjectMetadataMigrations].
type ScriptUpdateParamsBodyObjectMetadataMigrationsUnion interface {
	implementsWorkersScriptUpdateParamsBodyObjectMetadataMigrationsUnion()
}

// Usage model to apply to invocations.
type ScriptUpdateParamsBodyObjectMetadataUsageModel string

const (
	ScriptUpdateParamsBodyObjectMetadataUsageModelBundled ScriptUpdateParamsBodyObjectMetadataUsageModel = "bundled"
	ScriptUpdateParamsBodyObjectMetadataUsageModelUnbound ScriptUpdateParamsBodyObjectMetadataUsageModel = "unbound"
)

func (r ScriptUpdateParamsBodyObjectMetadataUsageModel) IsKnown() bool {
	switch r {
	case ScriptUpdateParamsBodyObjectMetadataUsageModelBundled, ScriptUpdateParamsBodyObjectMetadataUsageModelUnbound:
		return true
	}
	return false
}

type ScriptUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ScriptUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Script                              `json:"result"`
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
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScriptGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
