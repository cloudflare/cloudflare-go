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

// AccountWorkerDispatchNamespaceScriptService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountWorkerDispatchNamespaceScriptService] method instead.
type AccountWorkerDispatchNamespaceScriptService struct {
	Options  []option.RequestOption
	Content  *AccountWorkerDispatchNamespaceScriptContentService
	Settings *AccountWorkerDispatchNamespaceScriptSettingService
}

// NewAccountWorkerDispatchNamespaceScriptService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountWorkerDispatchNamespaceScriptService(opts ...option.RequestOption) (r *AccountWorkerDispatchNamespaceScriptService) {
	r = &AccountWorkerDispatchNamespaceScriptService{}
	r.Options = opts
	r.Content = NewAccountWorkerDispatchNamespaceScriptContentService(opts...)
	r.Settings = NewAccountWorkerDispatchNamespaceScriptSettingService(opts...)
	return
}

// Upload a worker module to a Workers for Platforms namespace.
func (r *AccountWorkerDispatchNamespaceScriptService) New(ctx context.Context, accountIdentifier string, dispatchNamespace string, scriptName string, body AccountWorkerDispatchNamespaceScriptNewParams, opts ...option.RequestOption) (res *AccountWorkerDispatchNamespaceScriptNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", accountIdentifier, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetch information about a script uploaded to a Workers for Platforms namespace.
func (r *AccountWorkerDispatchNamespaceScriptService) Get(ctx context.Context, accountIdentifier string, dispatchNamespace string, scriptName string, opts ...option.RequestOption) (res *AccountWorkerDispatchNamespaceScriptGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", accountIdentifier, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a worker from a Workers for Platforms namespace. This call has no
// response body on a successful delete.
func (r *AccountWorkerDispatchNamespaceScriptService) Delete(ctx context.Context, accountIdentifier string, dispatchNamespace string, scriptName string, body AccountWorkerDispatchNamespaceScriptDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s", accountIdentifier, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

type AccountWorkerDispatchNamespaceScriptNewResponse struct {
	Errors   []AccountWorkerDispatchNamespaceScriptNewResponseError   `json:"errors"`
	Messages []AccountWorkerDispatchNamespaceScriptNewResponseMessage `json:"messages"`
	Result   AccountWorkerDispatchNamespaceScriptNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountWorkerDispatchNamespaceScriptNewResponseSuccess `json:"success"`
	JSON    accountWorkerDispatchNamespaceScriptNewResponseJSON    `json:"-"`
}

// accountWorkerDispatchNamespaceScriptNewResponseJSON contains the JSON metadata
// for the struct [AccountWorkerDispatchNamespaceScriptNewResponse]
type accountWorkerDispatchNamespaceScriptNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDispatchNamespaceScriptNewResponseError struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accountWorkerDispatchNamespaceScriptNewResponseErrorJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptNewResponseErrorJSON contains the JSON
// metadata for the struct [AccountWorkerDispatchNamespaceScriptNewResponseError]
type accountWorkerDispatchNamespaceScriptNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDispatchNamespaceScriptNewResponseMessage struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accountWorkerDispatchNamespaceScriptNewResponseMessageJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptNewResponseMessageJSON contains the JSON
// metadata for the struct [AccountWorkerDispatchNamespaceScriptNewResponseMessage]
type accountWorkerDispatchNamespaceScriptNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDispatchNamespaceScriptNewResponseResult struct {
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
	TailConsumers []AccountWorkerDispatchNamespaceScriptNewResponseResultTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                                    `json:"usage_model"`
	JSON       accountWorkerDispatchNamespaceScriptNewResponseResultJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptNewResponseResultJSON contains the JSON
// metadata for the struct [AccountWorkerDispatchNamespaceScriptNewResponseResult]
type accountWorkerDispatchNamespaceScriptNewResponseResultJSON struct {
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

func (r *AccountWorkerDispatchNamespaceScriptNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A reference to a script that will consume logs from the attached Worker.
type AccountWorkerDispatchNamespaceScriptNewResponseResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                                `json:"namespace"`
	JSON      accountWorkerDispatchNamespaceScriptNewResponseResultTailConsumerJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptNewResponseResultTailConsumerJSON contains
// the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptNewResponseResultTailConsumer]
type accountWorkerDispatchNamespaceScriptNewResponseResultTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptNewResponseResultTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerDispatchNamespaceScriptNewResponseSuccess bool

const (
	AccountWorkerDispatchNamespaceScriptNewResponseSuccessTrue AccountWorkerDispatchNamespaceScriptNewResponseSuccess = true
)

// Details about a worker uploaded to a Workers for Platforms namespace.
type AccountWorkerDispatchNamespaceScriptGetResponse struct {
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Name of the Workers for Platforms dispatch namespace.
	DispatchNamespace string `json:"dispatch_namespace"`
	// When the script was last modified.
	ModifiedOn time.Time                                             `json:"modified_on" format:"date-time"`
	Script     AccountWorkerDispatchNamespaceScriptGetResponseScript `json:"script"`
	JSON       accountWorkerDispatchNamespaceScriptGetResponseJSON   `json:"-"`
}

// accountWorkerDispatchNamespaceScriptGetResponseJSON contains the JSON metadata
// for the struct [AccountWorkerDispatchNamespaceScriptGetResponse]
type accountWorkerDispatchNamespaceScriptGetResponseJSON struct {
	CreatedOn         apijson.Field
	DispatchNamespace apijson.Field
	ModifiedOn        apijson.Field
	Script            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDispatchNamespaceScriptGetResponseScript struct {
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
	TailConsumers []AccountWorkerDispatchNamespaceScriptGetResponseScriptTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                                    `json:"usage_model"`
	JSON       accountWorkerDispatchNamespaceScriptGetResponseScriptJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptGetResponseScriptJSON contains the JSON
// metadata for the struct [AccountWorkerDispatchNamespaceScriptGetResponseScript]
type accountWorkerDispatchNamespaceScriptGetResponseScriptJSON struct {
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

func (r *AccountWorkerDispatchNamespaceScriptGetResponseScript) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A reference to a script that will consume logs from the attached Worker.
type AccountWorkerDispatchNamespaceScriptGetResponseScriptTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                                `json:"namespace"`
	JSON      accountWorkerDispatchNamespaceScriptGetResponseScriptTailConsumerJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptGetResponseScriptTailConsumerJSON contains
// the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptGetResponseScriptTailConsumer]
type accountWorkerDispatchNamespaceScriptGetResponseScriptTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptGetResponseScriptTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// This interface is a union satisfied by one of the following:
// [AccountWorkerDispatchNamespaceScriptNewParamsVariant0],
// [AccountWorkerDispatchNamespaceScriptNewParamsVariant1].
type AccountWorkerDispatchNamespaceScriptNewParams interface {
	ImplementsAccountWorkerDispatchNamespaceScriptNewParams()
}

type AccountWorkerDispatchNamespaceScriptNewParamsVariant0 struct {
	// A module comprising a Worker script, often a javascript file. Multiple modules
	// may be provided as separate named parts, but at least one module must be present
	// and referenced in the metadata as `main_module` or `body_part` by part name.
	AnyPartName param.Field[[]io.Reader] `json:"<any part name>" format:"binary"`
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[AccountWorkerDispatchNamespaceScriptNewParamsVariant0Metadata] `json:"metadata"`
}

func (r AccountWorkerDispatchNamespaceScriptNewParamsVariant0) MarshalMultipart() (data []byte, contentType string, err error) {
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

func (AccountWorkerDispatchNamespaceScriptNewParamsVariant0) ImplementsAccountWorkerDispatchNamespaceScriptNewParams() {

}

// JSON encoded metadata about the uploaded parts and Worker configuration.
type AccountWorkerDispatchNamespaceScriptNewParamsVariant0Metadata struct {
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
	Migrations param.Field[AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrations] `json:"migrations"`
	Placement  param.Field[AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataPlacement]  `json:"placement"`
	// List of strings to use as tags for this Worker
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataTailConsumer] `json:"tail_consumers"`
	// Usage model to apply to invocations.
	UsageModel param.Field[AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataUsageModel] `json:"usage_model"`
	// Key-value pairs to use as tags for this version of this Worker
	VersionTags param.Field[interface{}] `json:"version_tags"`
}

func (r AccountWorkerDispatchNamespaceScriptNewParamsVariant0Metadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrations],
// [AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrations].
type AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrations interface {
	implementsAccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrations()
}

// A single set of migrations to apply.
type AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrations) implementsAccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrations() {
}

type AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrationsStep] `json:"steps"`
}

func (r AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrations) implementsAccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrations() {
}

type AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataPlacementMode] `json:"mode"`
}

func (r AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataPlacementMode string

const (
	AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataPlacementModeSmart AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Usage model to apply to invocations.
type AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataUsageModel string

const (
	AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataUsageModelBundled AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataUsageModel = "bundled"
	AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataUsageModelUnbound AccountWorkerDispatchNamespaceScriptNewParamsVariant0MetadataUsageModel = "unbound"
)

type AccountWorkerDispatchNamespaceScriptNewParamsVariant1 struct {
	// Rollback message to be associated with this deployment. Only parsed when query
	// param `"rollback_to"` is present.
	Message param.Field[string] `json:"message"`
}

func (r AccountWorkerDispatchNamespaceScriptNewParamsVariant1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AccountWorkerDispatchNamespaceScriptNewParamsVariant1) ImplementsAccountWorkerDispatchNamespaceScriptNewParams() {

}

type AccountWorkerDispatchNamespaceScriptDeleteParams struct {
	// If set to true, delete will not be stopped by associated service binding,
	// durable object, or other binding. Any of these associated bindings/durable
	// objects will be deleted along with the script.
	Force param.Field[bool] `query:"force"`
}

// URLQuery serializes [AccountWorkerDispatchNamespaceScriptDeleteParams]'s query
// parameters as `url.Values`.
func (r AccountWorkerDispatchNamespaceScriptDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
