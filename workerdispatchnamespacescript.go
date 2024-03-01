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

// WorkerDispatchNamespaceScriptService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewWorkerDispatchNamespaceScriptService] method instead.
type WorkerDispatchNamespaceScriptService struct {
	Options []option.RequestOption
}

// NewWorkerDispatchNamespaceScriptService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWorkerDispatchNamespaceScriptService(opts ...option.RequestOption) (r *WorkerDispatchNamespaceScriptService) {
	r = &WorkerDispatchNamespaceScriptService{}
	r.Options = opts
	return
}

// Upload a worker module to a Workers for Platforms namespace.
func (r *WorkerDispatchNamespaceScriptService) Update(ctx context.Context, dispatchNamespace string, scriptName string, params WorkerDispatchNamespaceScriptUpdateParams, opts ...option.RequestOption) (res *WorkerDispatchNamespaceScriptUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerDispatchNamespaceScriptUpdateResponseEnvelope
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
func (r *WorkerDispatchNamespaceScriptService) Delete(ctx context.Context, dispatchNamespace string, scriptName string, params WorkerDispatchNamespaceScriptDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

// Fetch information about a script uploaded to a Workers for Platforms namespace.
func (r *WorkerDispatchNamespaceScriptService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query WorkerDispatchNamespaceScriptGetParams, opts ...option.RequestOption) (res *WorkerDispatchNamespaceScriptGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerDispatchNamespaceScriptGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerDispatchNamespaceScriptUpdateResponse struct {
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
	TailConsumers []WorkerDispatchNamespaceScriptUpdateResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                          `json:"usage_model"`
	JSON       workerDispatchNamespaceScriptUpdateResponseJSON `json:"-"`
}

// workerDispatchNamespaceScriptUpdateResponseJSON contains the JSON metadata for
// the struct [WorkerDispatchNamespaceScriptUpdateResponse]
type workerDispatchNamespaceScriptUpdateResponseJSON struct {
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

func (r *WorkerDispatchNamespaceScriptUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A reference to a script that will consume logs from the attached Worker.
type WorkerDispatchNamespaceScriptUpdateResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                      `json:"namespace"`
	JSON      workerDispatchNamespaceScriptUpdateResponseTailConsumerJSON `json:"-"`
}

// workerDispatchNamespaceScriptUpdateResponseTailConsumerJSON contains the JSON
// metadata for the struct
// [WorkerDispatchNamespaceScriptUpdateResponseTailConsumer]
type workerDispatchNamespaceScriptUpdateResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDispatchNamespaceScriptUpdateResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details about a worker uploaded to a Workers for Platforms namespace.
type WorkerDispatchNamespaceScriptGetResponse struct {
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Name of the Workers for Platforms dispatch namespace.
	DispatchNamespace string `json:"dispatch_namespace"`
	// When the script was last modified.
	ModifiedOn time.Time                                      `json:"modified_on" format:"date-time"`
	Script     WorkerDispatchNamespaceScriptGetResponseScript `json:"script"`
	JSON       workerDispatchNamespaceScriptGetResponseJSON   `json:"-"`
}

// workerDispatchNamespaceScriptGetResponseJSON contains the JSON metadata for the
// struct [WorkerDispatchNamespaceScriptGetResponse]
type workerDispatchNamespaceScriptGetResponseJSON struct {
	CreatedOn         apijson.Field
	DispatchNamespace apijson.Field
	ModifiedOn        apijson.Field
	Script            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *WorkerDispatchNamespaceScriptGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDispatchNamespaceScriptGetResponseScript struct {
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
	TailConsumers []WorkerDispatchNamespaceScriptGetResponseScriptTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                             `json:"usage_model"`
	JSON       workerDispatchNamespaceScriptGetResponseScriptJSON `json:"-"`
}

// workerDispatchNamespaceScriptGetResponseScriptJSON contains the JSON metadata
// for the struct [WorkerDispatchNamespaceScriptGetResponseScript]
type workerDispatchNamespaceScriptGetResponseScriptJSON struct {
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

func (r *WorkerDispatchNamespaceScriptGetResponseScript) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A reference to a script that will consume logs from the attached Worker.
type WorkerDispatchNamespaceScriptGetResponseScriptTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                         `json:"namespace"`
	JSON      workerDispatchNamespaceScriptGetResponseScriptTailConsumerJSON `json:"-"`
}

// workerDispatchNamespaceScriptGetResponseScriptTailConsumerJSON contains the JSON
// metadata for the struct
// [WorkerDispatchNamespaceScriptGetResponseScriptTailConsumer]
type workerDispatchNamespaceScriptGetResponseScriptTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDispatchNamespaceScriptGetResponseScriptTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDispatchNamespaceScriptUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A module comprising a Worker script, often a javascript file. Multiple modules
	// may be provided as separate named parts, but at least one module must be present
	// and referenced in the metadata as `main_module` or `body_part` by part name.
	AnyPartName param.Field[[]io.Reader] `json:"<any part name>" format:"binary"`
	// Rollback message to be associated with this deployment. Only parsed when query
	// param `"rollback_to"` is present.
	Message param.Field[string] `json:"message"`
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[WorkerDispatchNamespaceScriptUpdateParamsMetadata] `json:"metadata"`
}

func (r WorkerDispatchNamespaceScriptUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// JSON encoded metadata about the uploaded parts and Worker configuration.
type WorkerDispatchNamespaceScriptUpdateParamsMetadata struct {
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
	Migrations param.Field[WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrations] `json:"migrations"`
	Placement  param.Field[WorkerDispatchNamespaceScriptUpdateParamsMetadataPlacement]  `json:"placement"`
	// List of strings to use as tags for this Worker
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]WorkerDispatchNamespaceScriptUpdateParamsMetadataTailConsumer] `json:"tail_consumers"`
	// Usage model to apply to invocations.
	UsageModel param.Field[WorkerDispatchNamespaceScriptUpdateParamsMetadataUsageModel] `json:"usage_model"`
	// Key-value pairs to use as tags for this version of this Worker
	VersionTags param.Field[interface{}] `json:"version_tags"`
}

func (r WorkerDispatchNamespaceScriptUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrations],
// [WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrations].
type WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrations interface {
	implementsWorkerDispatchNamespaceScriptUpdateParamsMetadataMigrations()
}

// A single set of migrations to apply.
type WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrations) implementsWorkerDispatchNamespaceScriptUpdateParamsMetadataMigrations() {
}

type WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStep] `json:"steps"`
}

func (r WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrations) implementsWorkerDispatchNamespaceScriptUpdateParamsMetadataMigrations() {
}

type WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkerDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerDispatchNamespaceScriptUpdateParamsMetadataPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[WorkerDispatchNamespaceScriptUpdateParamsMetadataPlacementMode] `json:"mode"`
}

func (r WorkerDispatchNamespaceScriptUpdateParamsMetadataPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkerDispatchNamespaceScriptUpdateParamsMetadataPlacementMode string

const (
	WorkerDispatchNamespaceScriptUpdateParamsMetadataPlacementModeSmart WorkerDispatchNamespaceScriptUpdateParamsMetadataPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkerDispatchNamespaceScriptUpdateParamsMetadataTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r WorkerDispatchNamespaceScriptUpdateParamsMetadataTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Usage model to apply to invocations.
type WorkerDispatchNamespaceScriptUpdateParamsMetadataUsageModel string

const (
	WorkerDispatchNamespaceScriptUpdateParamsMetadataUsageModelBundled WorkerDispatchNamespaceScriptUpdateParamsMetadataUsageModel = "bundled"
	WorkerDispatchNamespaceScriptUpdateParamsMetadataUsageModelUnbound WorkerDispatchNamespaceScriptUpdateParamsMetadataUsageModel = "unbound"
)

type WorkerDispatchNamespaceScriptUpdateResponseEnvelope struct {
	Errors   []WorkerDispatchNamespaceScriptUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerDispatchNamespaceScriptUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerDispatchNamespaceScriptUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerDispatchNamespaceScriptUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerDispatchNamespaceScriptUpdateResponseEnvelopeJSON    `json:"-"`
}

// workerDispatchNamespaceScriptUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [WorkerDispatchNamespaceScriptUpdateResponseEnvelope]
type workerDispatchNamespaceScriptUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDispatchNamespaceScriptUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDispatchNamespaceScriptUpdateResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    workerDispatchNamespaceScriptUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// workerDispatchNamespaceScriptUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [WorkerDispatchNamespaceScriptUpdateResponseEnvelopeErrors]
type workerDispatchNamespaceScriptUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDispatchNamespaceScriptUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDispatchNamespaceScriptUpdateResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    workerDispatchNamespaceScriptUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// workerDispatchNamespaceScriptUpdateResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [WorkerDispatchNamespaceScriptUpdateResponseEnvelopeMessages]
type workerDispatchNamespaceScriptUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDispatchNamespaceScriptUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerDispatchNamespaceScriptUpdateResponseEnvelopeSuccess bool

const (
	WorkerDispatchNamespaceScriptUpdateResponseEnvelopeSuccessTrue WorkerDispatchNamespaceScriptUpdateResponseEnvelopeSuccess = true
)

type WorkerDispatchNamespaceScriptDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// If set to true, delete will not be stopped by associated service binding,
	// durable object, or other binding. Any of these associated bindings/durable
	// objects will be deleted along with the script.
	Force param.Field[bool] `query:"force"`
}

// URLQuery serializes [WorkerDispatchNamespaceScriptDeleteParams]'s query
// parameters as `url.Values`.
func (r WorkerDispatchNamespaceScriptDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkerDispatchNamespaceScriptGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type WorkerDispatchNamespaceScriptGetResponseEnvelope struct {
	Errors   []WorkerDispatchNamespaceScriptGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerDispatchNamespaceScriptGetResponseEnvelopeMessages `json:"messages,required"`
	// Details about a worker uploaded to a Workers for Platforms namespace.
	Result WorkerDispatchNamespaceScriptGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success WorkerDispatchNamespaceScriptGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerDispatchNamespaceScriptGetResponseEnvelopeJSON    `json:"-"`
}

// workerDispatchNamespaceScriptGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [WorkerDispatchNamespaceScriptGetResponseEnvelope]
type workerDispatchNamespaceScriptGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDispatchNamespaceScriptGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDispatchNamespaceScriptGetResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    workerDispatchNamespaceScriptGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workerDispatchNamespaceScriptGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [WorkerDispatchNamespaceScriptGetResponseEnvelopeErrors]
type workerDispatchNamespaceScriptGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDispatchNamespaceScriptGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDispatchNamespaceScriptGetResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    workerDispatchNamespaceScriptGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workerDispatchNamespaceScriptGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [WorkerDispatchNamespaceScriptGetResponseEnvelopeMessages]
type workerDispatchNamespaceScriptGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDispatchNamespaceScriptGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerDispatchNamespaceScriptGetResponseEnvelopeSuccess bool

const (
	WorkerDispatchNamespaceScriptGetResponseEnvelopeSuccessTrue WorkerDispatchNamespaceScriptGetResponseEnvelopeSuccess = true
)
