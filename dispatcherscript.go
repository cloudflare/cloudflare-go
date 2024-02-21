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

// DispatcherScriptService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDispatcherScriptService] method
// instead.
type DispatcherScriptService struct {
	Options []option.RequestOption
}

// NewDispatcherScriptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDispatcherScriptService(opts ...option.RequestOption) (r *DispatcherScriptService) {
	r = &DispatcherScriptService{}
	r.Options = opts
	return
}

// Upload a worker module to a Workers for Platforms namespace.
func (r *DispatcherScriptService) Update(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, body DispatcherScriptUpdateParams, opts ...option.RequestOption) (res *DispatcherScriptUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DispatcherScriptUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a worker from a Workers for Platforms namespace. This call has no
// response body on a successful delete.
func (r *DispatcherScriptService) Delete(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, body DispatcherScriptDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Fetch information about a script uploaded to a Workers for Platforms namespace.
func (r *DispatcherScriptService) Get(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, opts ...option.RequestOption) (res *DispatcherScriptGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DispatcherScriptGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DispatcherScriptUpdateResponse struct {
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
	TailConsumers []DispatcherScriptUpdateResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                             `json:"usage_model"`
	JSON       dispatcherScriptUpdateResponseJSON `json:"-"`
}

// dispatcherScriptUpdateResponseJSON contains the JSON metadata for the struct
// [DispatcherScriptUpdateResponse]
type dispatcherScriptUpdateResponseJSON struct {
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

func (r *DispatcherScriptUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A reference to a script that will consume logs from the attached Worker.
type DispatcherScriptUpdateResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                         `json:"namespace"`
	JSON      dispatcherScriptUpdateResponseTailConsumerJSON `json:"-"`
}

// dispatcherScriptUpdateResponseTailConsumerJSON contains the JSON metadata for
// the struct [DispatcherScriptUpdateResponseTailConsumer]
type dispatcherScriptUpdateResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatcherScriptUpdateResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details about a worker uploaded to a Workers for Platforms namespace.
type DispatcherScriptGetResponse struct {
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Name of the Workers for Platforms dispatch namespace.
	DispatchNamespace string `json:"dispatch_namespace"`
	// When the script was last modified.
	ModifiedOn time.Time                         `json:"modified_on" format:"date-time"`
	Script     DispatcherScriptGetResponseScript `json:"script"`
	JSON       dispatcherScriptGetResponseJSON   `json:"-"`
}

// dispatcherScriptGetResponseJSON contains the JSON metadata for the struct
// [DispatcherScriptGetResponse]
type dispatcherScriptGetResponseJSON struct {
	CreatedOn         apijson.Field
	DispatchNamespace apijson.Field
	ModifiedOn        apijson.Field
	Script            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DispatcherScriptGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DispatcherScriptGetResponseScript struct {
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
	TailConsumers []DispatcherScriptGetResponseScriptTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                `json:"usage_model"`
	JSON       dispatcherScriptGetResponseScriptJSON `json:"-"`
}

// dispatcherScriptGetResponseScriptJSON contains the JSON metadata for the struct
// [DispatcherScriptGetResponseScript]
type dispatcherScriptGetResponseScriptJSON struct {
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

func (r *DispatcherScriptGetResponseScript) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A reference to a script that will consume logs from the attached Worker.
type DispatcherScriptGetResponseScriptTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                            `json:"namespace"`
	JSON      dispatcherScriptGetResponseScriptTailConsumerJSON `json:"-"`
}

// dispatcherScriptGetResponseScriptTailConsumerJSON contains the JSON metadata for
// the struct [DispatcherScriptGetResponseScriptTailConsumer]
type dispatcherScriptGetResponseScriptTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatcherScriptGetResponseScriptTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DispatcherScriptUpdateParams struct {
	// A module comprising a Worker script, often a javascript file. Multiple modules
	// may be provided as separate named parts, but at least one module must be present
	// and referenced in the metadata as `main_module` or `body_part` by part name.
	AnyPartName param.Field[[]io.Reader] `json:"<any part name>" format:"binary"`
	// Rollback message to be associated with this deployment. Only parsed when query
	// param `"rollback_to"` is present.
	Message param.Field[string] `json:"message"`
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[DispatcherScriptUpdateParamsMetadata] `json:"metadata"`
}

func (r DispatcherScriptUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type DispatcherScriptUpdateParamsMetadata struct {
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
	Migrations param.Field[DispatcherScriptUpdateParamsMetadataMigrations] `json:"migrations"`
	Placement  param.Field[DispatcherScriptUpdateParamsMetadataPlacement]  `json:"placement"`
	// List of strings to use as tags for this Worker
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]DispatcherScriptUpdateParamsMetadataTailConsumer] `json:"tail_consumers"`
	// Usage model to apply to invocations.
	UsageModel param.Field[DispatcherScriptUpdateParamsMetadataUsageModel] `json:"usage_model"`
	// Key-value pairs to use as tags for this version of this Worker
	VersionTags param.Field[interface{}] `json:"version_tags"`
}

func (r DispatcherScriptUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [DispatcherScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrations],
// [DispatcherScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrations].
type DispatcherScriptUpdateParamsMetadataMigrations interface {
	implementsDispatcherScriptUpdateParamsMetadataMigrations()
}

// A single set of migrations to apply.
type DispatcherScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]DispatcherScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]DispatcherScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r DispatcherScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatcherScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrations) implementsDispatcherScriptUpdateParamsMetadataMigrations() {
}

type DispatcherScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r DispatcherScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatcherScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r DispatcherScriptUpdateParamsMetadataMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatcherScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]DispatcherScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStep] `json:"steps"`
}

func (r DispatcherScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatcherScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrations) implementsDispatcherScriptUpdateParamsMetadataMigrations() {
}

type DispatcherScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]DispatcherScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]DispatcherScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r DispatcherScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatcherScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r DispatcherScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatcherScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r DispatcherScriptUpdateParamsMetadataMigrationsWorkersSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatcherScriptUpdateParamsMetadataPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[DispatcherScriptUpdateParamsMetadataPlacementMode] `json:"mode"`
}

func (r DispatcherScriptUpdateParamsMetadataPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type DispatcherScriptUpdateParamsMetadataPlacementMode string

const (
	DispatcherScriptUpdateParamsMetadataPlacementModeSmart DispatcherScriptUpdateParamsMetadataPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type DispatcherScriptUpdateParamsMetadataTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r DispatcherScriptUpdateParamsMetadataTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Usage model to apply to invocations.
type DispatcherScriptUpdateParamsMetadataUsageModel string

const (
	DispatcherScriptUpdateParamsMetadataUsageModelBundled DispatcherScriptUpdateParamsMetadataUsageModel = "bundled"
	DispatcherScriptUpdateParamsMetadataUsageModelUnbound DispatcherScriptUpdateParamsMetadataUsageModel = "unbound"
)

type DispatcherScriptUpdateResponseEnvelope struct {
	Errors   []DispatcherScriptUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatcherScriptUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   DispatcherScriptUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DispatcherScriptUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatcherScriptUpdateResponseEnvelopeJSON    `json:"-"`
}

// dispatcherScriptUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DispatcherScriptUpdateResponseEnvelope]
type dispatcherScriptUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatcherScriptUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DispatcherScriptUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    dispatcherScriptUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// dispatcherScriptUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DispatcherScriptUpdateResponseEnvelopeErrors]
type dispatcherScriptUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatcherScriptUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DispatcherScriptUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    dispatcherScriptUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// dispatcherScriptUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DispatcherScriptUpdateResponseEnvelopeMessages]
type dispatcherScriptUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatcherScriptUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DispatcherScriptUpdateResponseEnvelopeSuccess bool

const (
	DispatcherScriptUpdateResponseEnvelopeSuccessTrue DispatcherScriptUpdateResponseEnvelopeSuccess = true
)

type DispatcherScriptDeleteParams struct {
	// If set to true, delete will not be stopped by associated service binding,
	// durable object, or other binding. Any of these associated bindings/durable
	// objects will be deleted along with the script.
	Force param.Field[bool] `query:"force"`
}

// URLQuery serializes [DispatcherScriptDeleteParams]'s query parameters as
// `url.Values`.
func (r DispatcherScriptDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DispatcherScriptGetResponseEnvelope struct {
	Errors   []DispatcherScriptGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatcherScriptGetResponseEnvelopeMessages `json:"messages,required"`
	// Details about a worker uploaded to a Workers for Platforms namespace.
	Result DispatcherScriptGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success DispatcherScriptGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatcherScriptGetResponseEnvelopeJSON    `json:"-"`
}

// dispatcherScriptGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DispatcherScriptGetResponseEnvelope]
type dispatcherScriptGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatcherScriptGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DispatcherScriptGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    dispatcherScriptGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dispatcherScriptGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DispatcherScriptGetResponseEnvelopeErrors]
type dispatcherScriptGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatcherScriptGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DispatcherScriptGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    dispatcherScriptGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dispatcherScriptGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DispatcherScriptGetResponseEnvelopeMessages]
type dispatcherScriptGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatcherScriptGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DispatcherScriptGetResponseEnvelopeSuccess bool

const (
	DispatcherScriptGetResponseEnvelopeSuccessTrue DispatcherScriptGetResponseEnvelopeSuccess = true
)
