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
	return
}

// Fetch a list of uploaded workers.
func (r *WorkerScriptService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]WorkerScriptListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete your worker. This call has no response body on a successful delete.
func (r *WorkerScriptService) Delete(ctx context.Context, accountID string, scriptName string, body WorkerScriptDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Fetch raw script content for your worker. Note this is the original script
// content, not JSON encoded.
func (r *WorkerScriptService) Get(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "undefined")}, opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a worker module.
func (r *WorkerScriptService) Replace(ctx context.Context, accountID string, scriptName string, params WorkerScriptReplaceParams, opts ...option.RequestOption) (res *WorkerScriptReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

type WorkerScriptReplaceResponse struct {
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
	TailConsumers []WorkerScriptReplaceResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                          `json:"usage_model"`
	JSON       workerScriptReplaceResponseJSON `json:"-"`
}

// workerScriptReplaceResponseJSON contains the JSON metadata for the struct
// [WorkerScriptReplaceResponse]
type workerScriptReplaceResponseJSON struct {
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

func (r *WorkerScriptReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A reference to a script that will consume logs from the attached Worker.
type WorkerScriptReplaceResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                      `json:"namespace"`
	JSON      workerScriptReplaceResponseTailConsumerJSON `json:"-"`
}

// workerScriptReplaceResponseTailConsumerJSON contains the JSON metadata for the
// struct [WorkerScriptReplaceResponseTailConsumer]
type workerScriptReplaceResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptReplaceResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// Whether the API call was successful
type WorkerScriptListResponseEnvelopeSuccess bool

const (
	WorkerScriptListResponseEnvelopeSuccessTrue WorkerScriptListResponseEnvelopeSuccess = true
)

type WorkerScriptDeleteParams struct {
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

// This interface is a union satisfied by one of the following:
// [WorkerScriptReplaceParamsVariant0], [WorkerScriptReplaceParamsVariant1].
type WorkerScriptReplaceParams interface {
	ImplementsWorkerScriptReplaceParams()
}

type WorkerScriptReplaceParamsVariant0 struct {
	// Rollback to provided deployment based on deployment ID. Request body will only
	// parse a "message" part. You can learn more about deployments
	// [here](https://developers.cloudflare.com/workers/platform/deployments/).
	RollbackTo param.Field[string] `query:"rollback_to"`
	// A module comprising a Worker script, often a javascript file. Multiple modules
	// may be provided as separate named parts, but at least one module must be present
	// and referenced in the metadata as `main_module` or `body_part` by part name.
	AnyPartName param.Field[[]io.Reader] `json:"<any part name>" format:"binary"`
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[WorkerScriptReplaceParamsVariant0Metadata] `json:"metadata"`
}

func (r WorkerScriptReplaceParamsVariant0) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [WorkerScriptReplaceParamsVariant0]'s query parameters as
// `url.Values`.
func (r WorkerScriptReplaceParamsVariant0) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func (WorkerScriptReplaceParamsVariant0) ImplementsWorkerScriptReplaceParams() {

}

// JSON encoded metadata about the uploaded parts and Worker configuration.
type WorkerScriptReplaceParamsVariant0Metadata struct {
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
	Migrations param.Field[WorkerScriptReplaceParamsVariant0MetadataMigrations] `json:"migrations"`
	Placement  param.Field[WorkerScriptReplaceParamsVariant0MetadataPlacement]  `json:"placement"`
	// List of strings to use as tags for this Worker
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]WorkerScriptReplaceParamsVariant0MetadataTailConsumer] `json:"tail_consumers"`
	// Usage model to apply to invocations.
	UsageModel param.Field[WorkerScriptReplaceParamsVariant0MetadataUsageModel] `json:"usage_model"`
	// Key-value pairs to use as tags for this version of this Worker
	VersionTags param.Field[interface{}] `json:"version_tags"`
}

func (r WorkerScriptReplaceParamsVariant0Metadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSingleStepMigrations],
// [WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSteppedMigrations].
type WorkerScriptReplaceParamsVariant0MetadataMigrations interface {
	implementsWorkerScriptReplaceParamsVariant0MetadataMigrations()
}

// A single set of migrations to apply.
type WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSingleStepMigrations) implementsWorkerScriptReplaceParamsVariant0MetadataMigrations() {
}

type WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSteppedMigrationsStep] `json:"steps"`
}

func (r WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSteppedMigrations) implementsWorkerScriptReplaceParamsVariant0MetadataMigrations() {
}

type WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkerScriptReplaceParamsVariant0MetadataMigrationsWorkersSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptReplaceParamsVariant0MetadataPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[WorkerScriptReplaceParamsVariant0MetadataPlacementMode] `json:"mode"`
}

func (r WorkerScriptReplaceParamsVariant0MetadataPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkerScriptReplaceParamsVariant0MetadataPlacementMode string

const (
	WorkerScriptReplaceParamsVariant0MetadataPlacementModeSmart WorkerScriptReplaceParamsVariant0MetadataPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkerScriptReplaceParamsVariant0MetadataTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r WorkerScriptReplaceParamsVariant0MetadataTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Usage model to apply to invocations.
type WorkerScriptReplaceParamsVariant0MetadataUsageModel string

const (
	WorkerScriptReplaceParamsVariant0MetadataUsageModelBundled WorkerScriptReplaceParamsVariant0MetadataUsageModel = "bundled"
	WorkerScriptReplaceParamsVariant0MetadataUsageModelUnbound WorkerScriptReplaceParamsVariant0MetadataUsageModel = "unbound"
)

type WorkerScriptReplaceParamsVariant1 struct {
	// Rollback to provided deployment based on deployment ID. Request body will only
	// parse a "message" part. You can learn more about deployments
	// [here](https://developers.cloudflare.com/workers/platform/deployments/).
	RollbackTo param.Field[string] `query:"rollback_to"`
	// Rollback message to be associated with this deployment. Only parsed when query
	// param `"rollback_to"` is present.
	Message param.Field[string] `json:"message"`
}

func (r WorkerScriptReplaceParamsVariant1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [WorkerScriptReplaceParamsVariant1]'s query parameters as
// `url.Values`.
func (r WorkerScriptReplaceParamsVariant1) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func (WorkerScriptReplaceParamsVariant1) ImplementsWorkerScriptReplaceParams() {

}

type WorkerScriptReplaceResponseEnvelope struct {
	Errors   []WorkerScriptReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptReplaceResponseEnvelopeJSON    `json:"-"`
}

// workerScriptReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerScriptReplaceResponseEnvelope]
type workerScriptReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptReplaceResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    workerScriptReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerScriptReplaceResponseEnvelopeErrors]
type workerScriptReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptReplaceResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    workerScriptReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WorkerScriptReplaceResponseEnvelopeMessages]
type workerScriptReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptReplaceResponseEnvelopeSuccess bool

const (
	WorkerScriptReplaceResponseEnvelopeSuccessTrue WorkerScriptReplaceResponseEnvelopeSuccess = true
)
