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

// WorkersForPlatformDispatchNamespaceScriptService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkersForPlatformDispatchNamespaceScriptService] method instead.
type WorkersForPlatformDispatchNamespaceScriptService struct {
	Options []option.RequestOption
	Content *WorkersForPlatformDispatchNamespaceScriptContentService
}

// NewWorkersForPlatformDispatchNamespaceScriptService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewWorkersForPlatformDispatchNamespaceScriptService(opts ...option.RequestOption) (r *WorkersForPlatformDispatchNamespaceScriptService) {
	r = &WorkersForPlatformDispatchNamespaceScriptService{}
	r.Options = opts
	r.Content = NewWorkersForPlatformDispatchNamespaceScriptContentService(opts...)
	return
}

// Upload a worker module to a Workers for Platforms namespace.
func (r *WorkersForPlatformDispatchNamespaceScriptService) Update(ctx context.Context, dispatchNamespace string, scriptName string, params WorkersForPlatformDispatchNamespaceScriptUpdateParams, opts ...option.RequestOption) (res *WorkersForPlatformDispatchNamespaceScriptUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkersForPlatformDispatchNamespaceScriptUpdateResponseEnvelope
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
func (r *WorkersForPlatformDispatchNamespaceScriptService) Delete(ctx context.Context, dispatchNamespace string, scriptName string, params WorkersForPlatformDispatchNamespaceScriptDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

// Fetch information about a script uploaded to a Workers for Platforms namespace.
func (r *WorkersForPlatformDispatchNamespaceScriptService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query WorkersForPlatformDispatchNamespaceScriptGetParams, opts ...option.RequestOption) (res *WorkersForPlatformDispatchNamespaceScriptGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkersForPlatformDispatchNamespaceScriptGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkersForPlatformDispatchNamespaceScriptUpdateResponse struct {
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
	TailConsumers []WorkersForPlatformDispatchNamespaceScriptUpdateResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                                      `json:"usage_model"`
	JSON       workersForPlatformDispatchNamespaceScriptUpdateResponseJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptUpdateResponseJSON contains the JSON
// metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptUpdateResponse]
type workersForPlatformDispatchNamespaceScriptUpdateResponseJSON struct {
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

func (r *WorkersForPlatformDispatchNamespaceScriptUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersForPlatformDispatchNamespaceScriptUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a script that will consume logs from the attached Worker.
type WorkersForPlatformDispatchNamespaceScriptUpdateResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                                  `json:"namespace"`
	JSON      workersForPlatformDispatchNamespaceScriptUpdateResponseTailConsumerJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptUpdateResponseTailConsumerJSON contains
// the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptUpdateResponseTailConsumer]
type workersForPlatformDispatchNamespaceScriptUpdateResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptUpdateResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersForPlatformDispatchNamespaceScriptUpdateResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

// Details about a worker uploaded to a Workers for Platforms namespace.
type WorkersForPlatformDispatchNamespaceScriptGetResponse struct {
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Name of the Workers for Platforms dispatch namespace.
	DispatchNamespace string `json:"dispatch_namespace"`
	// When the script was last modified.
	ModifiedOn time.Time                                                  `json:"modified_on" format:"date-time"`
	Script     WorkersForPlatformDispatchNamespaceScriptGetResponseScript `json:"script"`
	JSON       workersForPlatformDispatchNamespaceScriptGetResponseJSON   `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptGetResponseJSON contains the JSON
// metadata for the struct [WorkersForPlatformDispatchNamespaceScriptGetResponse]
type workersForPlatformDispatchNamespaceScriptGetResponseJSON struct {
	CreatedOn         apijson.Field
	DispatchNamespace apijson.Field
	ModifiedOn        apijson.Field
	Script            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersForPlatformDispatchNamespaceScriptGetResponseJSON) RawJSON() string {
	return r.raw
}

type WorkersForPlatformDispatchNamespaceScriptGetResponseScript struct {
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
	TailConsumers []WorkersForPlatformDispatchNamespaceScriptGetResponseScriptTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                                         `json:"usage_model"`
	JSON       workersForPlatformDispatchNamespaceScriptGetResponseScriptJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptGetResponseScriptJSON contains the JSON
// metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptGetResponseScript]
type workersForPlatformDispatchNamespaceScriptGetResponseScriptJSON struct {
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

func (r *WorkersForPlatformDispatchNamespaceScriptGetResponseScript) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersForPlatformDispatchNamespaceScriptGetResponseScriptJSON) RawJSON() string {
	return r.raw
}

// A reference to a script that will consume logs from the attached Worker.
type WorkersForPlatformDispatchNamespaceScriptGetResponseScriptTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                                     `json:"namespace"`
	JSON      workersForPlatformDispatchNamespaceScriptGetResponseScriptTailConsumerJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptGetResponseScriptTailConsumerJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptGetResponseScriptTailConsumer]
type workersForPlatformDispatchNamespaceScriptGetResponseScriptTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptGetResponseScriptTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersForPlatformDispatchNamespaceScriptGetResponseScriptTailConsumerJSON) RawJSON() string {
	return r.raw
}

type WorkersForPlatformDispatchNamespaceScriptUpdateParams struct {
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
	Metadata param.Field[WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadata] `json:"metadata"`
}

func (r WorkersForPlatformDispatchNamespaceScriptUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadata struct {
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
	Migrations param.Field[WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrations] `json:"migrations"`
	Placement  param.Field[WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataPlacement]  `json:"placement"`
	// List of strings to use as tags for this Worker
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataTailConsumer] `json:"tail_consumers"`
	// Usage model to apply to invocations.
	UsageModel param.Field[WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataUsageModel] `json:"usage_model"`
	// Key-value pairs to use as tags for this version of this Worker
	VersionTags param.Field[interface{}] `json:"version_tags"`
}

func (r WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrations],
// [WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrations].
type WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrations interface {
	implementsWorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrations()
}

// A single set of migrations to apply.
type WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrations) implementsWorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStep] `json:"steps"`
}

func (r WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrations) implementsWorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataPlacementMode] `json:"mode"`
}

func (r WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataPlacementMode string

const (
	WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataPlacementModeSmart WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Usage model to apply to invocations.
type WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataUsageModel string

const (
	WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataUsageModelBundled WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataUsageModel = "bundled"
	WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataUsageModelUnbound WorkersForPlatformDispatchNamespaceScriptUpdateParamsMetadataUsageModel = "unbound"
)

type WorkersForPlatformDispatchNamespaceScriptUpdateResponseEnvelope struct {
	Errors   []WorkersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkersForPlatformDispatchNamespaceScriptUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    workersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeJSON    `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptUpdateResponseEnvelope]
type workersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeErrors struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeErrors]
type workersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeMessages struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeMessages]
type workersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeSuccess bool

const (
	WorkersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeSuccessTrue WorkersForPlatformDispatchNamespaceScriptUpdateResponseEnvelopeSuccess = true
)

type WorkersForPlatformDispatchNamespaceScriptDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// If set to true, delete will not be stopped by associated service binding,
	// durable object, or other binding. Any of these associated bindings/durable
	// objects will be deleted along with the script.
	Force param.Field[bool] `query:"force"`
}

// URLQuery serializes [WorkersForPlatformDispatchNamespaceScriptDeleteParams]'s
// query parameters as `url.Values`.
func (r WorkersForPlatformDispatchNamespaceScriptDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkersForPlatformDispatchNamespaceScriptGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type WorkersForPlatformDispatchNamespaceScriptGetResponseEnvelope struct {
	Errors   []WorkersForPlatformDispatchNamespaceScriptGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkersForPlatformDispatchNamespaceScriptGetResponseEnvelopeMessages `json:"messages,required"`
	// Details about a worker uploaded to a Workers for Platforms namespace.
	Result WorkersForPlatformDispatchNamespaceScriptGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success WorkersForPlatformDispatchNamespaceScriptGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    workersForPlatformDispatchNamespaceScriptGetResponseEnvelopeJSON    `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptGetResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptGetResponseEnvelope]
type workersForPlatformDispatchNamespaceScriptGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersForPlatformDispatchNamespaceScriptGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkersForPlatformDispatchNamespaceScriptGetResponseEnvelopeErrors struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptGetResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptGetResponseEnvelopeErrors]
type workersForPlatformDispatchNamespaceScriptGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersForPlatformDispatchNamespaceScriptGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkersForPlatformDispatchNamespaceScriptGetResponseEnvelopeMessages struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptGetResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptGetResponseEnvelopeMessages]
type workersForPlatformDispatchNamespaceScriptGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersForPlatformDispatchNamespaceScriptGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkersForPlatformDispatchNamespaceScriptGetResponseEnvelopeSuccess bool

const (
	WorkersForPlatformDispatchNamespaceScriptGetResponseEnvelopeSuccessTrue WorkersForPlatformDispatchNamespaceScriptGetResponseEnvelopeSuccess = true
)
