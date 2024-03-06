// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apiform"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// WorkerScriptService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWorkerScriptService] method
// instead.
type WorkerScriptService struct {
	Options    []option.RequestOption
	Bindings   *WorkerScriptBindingService
	Schedules  *WorkerScriptScheduleService
	Tail       *WorkerScriptTailService
	UsageModel *WorkerScriptUsageModelService
	Content    *WorkerScriptContentService
	ContentV2  *WorkerScriptContentV2Service
	Settings   *WorkerScriptSettingService
}

// NewWorkerScriptService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkerScriptService(opts ...option.RequestOption) (r *WorkerScriptService) {
	r = &WorkerScriptService{}
	r.Options = opts
	r.Bindings = NewWorkerScriptBindingService(opts...)
	r.Schedules = NewWorkerScriptScheduleService(opts...)
	r.Tail = NewWorkerScriptTailService(opts...)
	r.UsageModel = NewWorkerScriptUsageModelService(opts...)
	r.Content = NewWorkerScriptContentService(opts...)
	r.ContentV2 = NewWorkerScriptContentV2Service(opts...)
	r.Settings = NewWorkerScriptSettingService(opts...)
	return
}

// Upload a worker module.
func (r *WorkerScriptService) Update(ctx context.Context, scriptName string, params WorkerScriptUpdateParams, opts ...option.RequestOption) (res *WorkerScriptUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of uploaded workers.
func (r *WorkerScriptService) List(ctx context.Context, query WorkerScriptListParams, opts ...option.RequestOption) (res *[]WorkerScriptListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete your worker. This call has no response body on a successful delete.
func (r *WorkerScriptService) Delete(ctx context.Context, scriptName string, params WorkerScriptDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

// Fetch raw script content for your worker. Note this is the original script
// content, not JSON encoded.
func (r *WorkerScriptService) Get(ctx context.Context, scriptName string, query WorkerScriptGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "undefined")}, opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type WorkerScriptUpdateResponse struct {
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
	// Deprecated. Deployment metadata for internal usage.
	PipelineHash string `json:"pipeline_hash"`
	// Specifies the placement mode for the Worker (e.g. 'smart').
	PlacementMode string `json:"placement_mode"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkerScriptUpdateResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                         `json:"usage_model"`
	JSON       workerScriptUpdateResponseJSON `json:"-"`
}

// workerScriptUpdateResponseJSON contains the JSON metadata for the struct
// [WorkerScriptUpdateResponse]
type workerScriptUpdateResponseJSON struct {
	ID            apijson.Field
	CreatedOn     apijson.Field
	Etag          apijson.Field
	Logpush       apijson.Field
	ModifiedOn    apijson.Field
	PipelineHash  apijson.Field
	PlacementMode apijson.Field
	TailConsumers apijson.Field
	UsageModel    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerScriptUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a script that will consume logs from the attached Worker.
type WorkerScriptUpdateResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                     `json:"namespace"`
	JSON      workerScriptUpdateResponseTailConsumerJSON `json:"-"`
}

// workerScriptUpdateResponseTailConsumerJSON contains the JSON metadata for the
// struct [WorkerScriptUpdateResponseTailConsumer]
type workerScriptUpdateResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUpdateResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptUpdateResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptListResponse struct {
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
	// Deprecated. Deployment metadata for internal usage.
	PipelineHash string `json:"pipeline_hash"`
	// Specifies the placement mode for the Worker (e.g. 'smart').
	PlacementMode string `json:"placement_mode"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkerScriptListResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                       `json:"usage_model"`
	JSON       workerScriptListResponseJSON `json:"-"`
}

// workerScriptListResponseJSON contains the JSON metadata for the struct
// [WorkerScriptListResponse]
type workerScriptListResponseJSON struct {
	ID            apijson.Field
	CreatedOn     apijson.Field
	Etag          apijson.Field
	Logpush       apijson.Field
	ModifiedOn    apijson.Field
	PipelineHash  apijson.Field
	PlacementMode apijson.Field
	TailConsumers apijson.Field
	UsageModel    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerScriptListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptListResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a script that will consume logs from the attached Worker.
type WorkerScriptListResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                   `json:"namespace"`
	JSON      workerScriptListResponseTailConsumerJSON `json:"-"`
}

// workerScriptListResponseTailConsumerJSON contains the JSON metadata for the
// struct [WorkerScriptListResponseTailConsumer]
type workerScriptListResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptListResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptListResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Rollback to provided deployment based on deployment ID. Request body will only
	// parse a "message" part. You can learn more about deployments
	// [here](https://developers.cloudflare.com/workers/platform/deployments/).
	RollbackTo param.Field[string] `query:"rollback_to"`
	// A module comprising a Worker script, often a javascript file. Multiple modules
	// may be provided as separate named parts, but at least one module must be present
	// and referenced in the metadata as `main_module` or `body_part` by part name.
	AnyPartName param.Field[[]io.Reader] `json:"<any part name>" format:"binary"`
	// Rollback message to be associated with this deployment. Only parsed when query
	// param `"rollback_to"` is present.
	Message param.Field[string] `json:"message"`
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[WorkerScriptUpdateParamsMetadata] `json:"metadata"`
}

func (r WorkerScriptUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [WorkerScriptUpdateParams]'s query parameters as
// `url.Values`.
func (r WorkerScriptUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// JSON encoded metadata about the uploaded parts and Worker configuration.
type WorkerScriptUpdateParamsMetadata struct {
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
	Migrations param.Field[WorkerScriptUpdateParamsMetadataMigrations] `json:"migrations"`
	Placement  param.Field[WorkerScriptUpdateParamsMetadataPlacement]  `json:"placement"`
	// List of strings to use as tags for this Worker
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]WorkerScriptUpdateParamsMetadataTailConsumer] `json:"tail_consumers"`
	// Usage model to apply to invocations.
	UsageModel param.Field[WorkerScriptUpdateParamsMetadataUsageModel] `json:"usage_model"`
	// Key-value pairs to use as tags for this version of this Worker
	VersionTags param.Field[interface{}] `json:"version_tags"`
}

func (r WorkerScriptUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [WorkerScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrations],
// [WorkerScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrations].
type WorkerScriptUpdateParamsMetadataMigrations interface {
	implementsWorkerScriptUpdateParamsMetadataMigrations()
}

// A single set of migrations to apply.
type WorkerScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrations struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]WorkerScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkerScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r WorkerScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrations) implementsWorkerScriptUpdateParamsMetadataMigrations() {
}

type WorkerScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkerScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkerScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]WorkerScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStep] `json:"steps"`
}

func (r WorkerScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrations) implementsWorkerScriptUpdateParamsMetadataMigrations() {
}

type WorkerScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]WorkerScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkerScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r WorkerScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkerScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkerScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptUpdateParamsMetadataPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[WorkerScriptUpdateParamsMetadataPlacementMode] `json:"mode"`
}

func (r WorkerScriptUpdateParamsMetadataPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkerScriptUpdateParamsMetadataPlacementMode string

const (
	WorkerScriptUpdateParamsMetadataPlacementModeSmart WorkerScriptUpdateParamsMetadataPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkerScriptUpdateParamsMetadataTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r WorkerScriptUpdateParamsMetadataTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Usage model to apply to invocations.
type WorkerScriptUpdateParamsMetadataUsageModel string

const (
	WorkerScriptUpdateParamsMetadataUsageModelBundled WorkerScriptUpdateParamsMetadataUsageModel = "bundled"
	WorkerScriptUpdateParamsMetadataUsageModelUnbound WorkerScriptUpdateParamsMetadataUsageModel = "unbound"
)

type WorkerScriptUpdateResponseEnvelope struct {
	Errors   []WorkerScriptUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptUpdateResponseEnvelopeJSON    `json:"-"`
}

// workerScriptUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerScriptUpdateResponseEnvelope]
type workerScriptUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptUpdateResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    workerScriptUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerScriptUpdateResponseEnvelopeErrors]
type workerScriptUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptUpdateResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    workerScriptUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WorkerScriptUpdateResponseEnvelopeMessages]
type workerScriptUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerScriptUpdateResponseEnvelopeSuccess bool

const (
	WorkerScriptUpdateResponseEnvelopeSuccessTrue WorkerScriptUpdateResponseEnvelopeSuccess = true
)

type WorkerScriptListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type WorkerScriptListResponseEnvelope struct {
	Errors   []WorkerScriptListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptListResponseEnvelopeMessages `json:"messages,required"`
	Result   []WorkerScriptListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptListResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptListResponseEnvelopeJSON    `json:"-"`
}

// workerScriptListResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerScriptListResponseEnvelope]
type workerScriptListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptListResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    workerScriptListResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerScriptListResponseEnvelopeErrors]
type workerScriptListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptListResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    workerScriptListResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WorkerScriptListResponseEnvelopeMessages]
type workerScriptListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerScriptListResponseEnvelopeSuccess bool

const (
	WorkerScriptListResponseEnvelopeSuccessTrue WorkerScriptListResponseEnvelopeSuccess = true
)

type WorkerScriptDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// If set to true, delete will not be stopped by associated service binding,
	// durable object, or other binding. Any of these associated bindings/durable
	// objects will be deleted along with the script.
	Force param.Field[bool] `query:"force"`
}

// URLQuery serializes [WorkerScriptDeleteParams]'s query parameters as
// `url.Values`.
func (r WorkerScriptDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkerScriptGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
