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

// AccountWorkerScriptService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountWorkerScriptService]
// method instead.
type AccountWorkerScriptService struct {
	Options     []option.RequestOption
	Schedules   *AccountWorkerScriptScheduleService
	Tails       *AccountWorkerScriptTailService
	UsageModels *AccountWorkerScriptUsageModelService
	Content     *AccountWorkerScriptContentService
	ContentV2   *AccountWorkerScriptContentV2Service
	Settings    *AccountWorkerScriptSettingService
}

// NewAccountWorkerScriptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountWorkerScriptService(opts ...option.RequestOption) (r *AccountWorkerScriptService) {
	r = &AccountWorkerScriptService{}
	r.Options = opts
	r.Schedules = NewAccountWorkerScriptScheduleService(opts...)
	r.Tails = NewAccountWorkerScriptTailService(opts...)
	r.UsageModels = NewAccountWorkerScriptUsageModelService(opts...)
	r.Content = NewAccountWorkerScriptContentService(opts...)
	r.ContentV2 = NewAccountWorkerScriptContentV2Service(opts...)
	r.Settings = NewAccountWorkerScriptSettingService(opts...)
	return
}

// Fetch raw script content for your worker. Note this is the original script
// content, not JSON encoded.
func (r *AccountWorkerScriptService) Get(ctx context.Context, accountIdentifier string, scriptName string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "undefined")}, opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", accountIdentifier, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a worker module.
func (r *AccountWorkerScriptService) Update(ctx context.Context, accountIdentifier string, scriptName string, params AccountWorkerScriptUpdateParams, opts ...option.RequestOption) (res *AccountWorkerScriptUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", accountIdentifier, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Delete your worker. This call has no response body on a successful delete.
func (r *AccountWorkerScriptService) Delete(ctx context.Context, accountIdentifier string, scriptName string, body AccountWorkerScriptDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s", accountIdentifier, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Fetch a list of uploaded workers.
func (r *AccountWorkerScriptService) WorkerScriptListWorkers(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountWorkerScriptWorkerScriptListWorkersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountWorkerScriptUpdateResponse struct {
	Errors   []AccountWorkerScriptUpdateResponseError   `json:"errors"`
	Messages []AccountWorkerScriptUpdateResponseMessage `json:"messages"`
	Result   AccountWorkerScriptUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountWorkerScriptUpdateResponseSuccess `json:"success"`
	JSON    accountWorkerScriptUpdateResponseJSON    `json:"-"`
}

// accountWorkerScriptUpdateResponseJSON contains the JSON metadata for the struct
// [AccountWorkerScriptUpdateResponse]
type accountWorkerScriptUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptUpdateResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountWorkerScriptUpdateResponseErrorJSON `json:"-"`
}

// accountWorkerScriptUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountWorkerScriptUpdateResponseError]
type accountWorkerScriptUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptUpdateResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountWorkerScriptUpdateResponseMessageJSON `json:"-"`
}

// accountWorkerScriptUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountWorkerScriptUpdateResponseMessage]
type accountWorkerScriptUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptUpdateResponseResult struct {
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
	TailConsumers []AccountWorkerScriptUpdateResponseResultTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                      `json:"usage_model"`
	JSON       accountWorkerScriptUpdateResponseResultJSON `json:"-"`
}

// accountWorkerScriptUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountWorkerScriptUpdateResponseResult]
type accountWorkerScriptUpdateResponseResultJSON struct {
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

func (r *AccountWorkerScriptUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A reference to a script that will consume logs from the attached Worker.
type AccountWorkerScriptUpdateResponseResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                  `json:"namespace"`
	JSON      accountWorkerScriptUpdateResponseResultTailConsumerJSON `json:"-"`
}

// accountWorkerScriptUpdateResponseResultTailConsumerJSON contains the JSON
// metadata for the struct [AccountWorkerScriptUpdateResponseResultTailConsumer]
type accountWorkerScriptUpdateResponseResultTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptUpdateResponseResultTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerScriptUpdateResponseSuccess bool

const (
	AccountWorkerScriptUpdateResponseSuccessTrue AccountWorkerScriptUpdateResponseSuccess = true
)

type AccountWorkerScriptWorkerScriptListWorkersResponse struct {
	Errors   []AccountWorkerScriptWorkerScriptListWorkersResponseError   `json:"errors"`
	Messages []AccountWorkerScriptWorkerScriptListWorkersResponseMessage `json:"messages"`
	Result   []AccountWorkerScriptWorkerScriptListWorkersResponseResult  `json:"result"`
	// Whether the API call was successful
	Success AccountWorkerScriptWorkerScriptListWorkersResponseSuccess `json:"success"`
	JSON    accountWorkerScriptWorkerScriptListWorkersResponseJSON    `json:"-"`
}

// accountWorkerScriptWorkerScriptListWorkersResponseJSON contains the JSON
// metadata for the struct [AccountWorkerScriptWorkerScriptListWorkersResponse]
type accountWorkerScriptWorkerScriptListWorkersResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptWorkerScriptListWorkersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptWorkerScriptListWorkersResponseError struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    accountWorkerScriptWorkerScriptListWorkersResponseErrorJSON `json:"-"`
}

// accountWorkerScriptWorkerScriptListWorkersResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountWorkerScriptWorkerScriptListWorkersResponseError]
type accountWorkerScriptWorkerScriptListWorkersResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptWorkerScriptListWorkersResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptWorkerScriptListWorkersResponseMessage struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    accountWorkerScriptWorkerScriptListWorkersResponseMessageJSON `json:"-"`
}

// accountWorkerScriptWorkerScriptListWorkersResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountWorkerScriptWorkerScriptListWorkersResponseMessage]
type accountWorkerScriptWorkerScriptListWorkersResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptWorkerScriptListWorkersResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptWorkerScriptListWorkersResponseResult struct {
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
	TailConsumers []AccountWorkerScriptWorkerScriptListWorkersResponseResultTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                                       `json:"usage_model"`
	JSON       accountWorkerScriptWorkerScriptListWorkersResponseResultJSON `json:"-"`
}

// accountWorkerScriptWorkerScriptListWorkersResponseResultJSON contains the JSON
// metadata for the struct
// [AccountWorkerScriptWorkerScriptListWorkersResponseResult]
type accountWorkerScriptWorkerScriptListWorkersResponseResultJSON struct {
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

func (r *AccountWorkerScriptWorkerScriptListWorkersResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A reference to a script that will consume logs from the attached Worker.
type AccountWorkerScriptWorkerScriptListWorkersResponseResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                                   `json:"namespace"`
	JSON      accountWorkerScriptWorkerScriptListWorkersResponseResultTailConsumerJSON `json:"-"`
}

// accountWorkerScriptWorkerScriptListWorkersResponseResultTailConsumerJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptWorkerScriptListWorkersResponseResultTailConsumer]
type accountWorkerScriptWorkerScriptListWorkersResponseResultTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptWorkerScriptListWorkersResponseResultTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerScriptWorkerScriptListWorkersResponseSuccess bool

const (
	AccountWorkerScriptWorkerScriptListWorkersResponseSuccessTrue AccountWorkerScriptWorkerScriptListWorkersResponseSuccess = true
)

// This interface is a union satisfied by one of the following:
// [AccountWorkerScriptUpdateParamsVariant0],
// [AccountWorkerScriptUpdateParamsVariant1].
type AccountWorkerScriptUpdateParams interface {
	ImplementsAccountWorkerScriptUpdateParams()
}

type AccountWorkerScriptUpdateParamsVariant0 struct {
	// Rollback to provided deployment based on deployment ID. Request body will only
	// parse a "message" part. You can learn more about deployments
	// [here](https://developers.cloudflare.com/workers/platform/deployments/).
	RollbackTo param.Field[string] `query:"rollback_to"`
	// A module comprising a Worker script, often a javascript file. Multiple modules
	// may be provided as separate named parts, but at least one module must be present
	// and referenced in the metadata as `main_module` or `body_part` by part name.
	AnyPartName param.Field[[]io.Reader] `json:"<any part name>" format:"binary"`
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata param.Field[AccountWorkerScriptUpdateParamsVariant0Metadata] `json:"metadata"`
}

func (r AccountWorkerScriptUpdateParamsVariant0) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [AccountWorkerScriptUpdateParamsVariant0]'s query parameters
// as `url.Values`.
func (r AccountWorkerScriptUpdateParamsVariant0) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func (AccountWorkerScriptUpdateParamsVariant0) ImplementsAccountWorkerScriptUpdateParams() {

}

// JSON encoded metadata about the uploaded parts and Worker configuration.
type AccountWorkerScriptUpdateParamsVariant0Metadata struct {
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
	Migrations param.Field[AccountWorkerScriptUpdateParamsVariant0MetadataMigrations] `json:"migrations"`
	Placement  param.Field[AccountWorkerScriptUpdateParamsVariant0MetadataPlacement]  `json:"placement"`
	// List of strings to use as tags for this Worker
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]AccountWorkerScriptUpdateParamsVariant0MetadataTailConsumer] `json:"tail_consumers"`
	// Usage model to apply to invocations.
	UsageModel param.Field[AccountWorkerScriptUpdateParamsVariant0MetadataUsageModel] `json:"usage_model"`
	// Key-value pairs to use as tags for this version of this Worker
	VersionTags param.Field[interface{}] `json:"version_tags"`
}

func (r AccountWorkerScriptUpdateParamsVariant0Metadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrations],
// [AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrations].
type AccountWorkerScriptUpdateParamsVariant0MetadataMigrations interface {
	implementsAccountWorkerScriptUpdateParamsVariant0MetadataMigrations()
}

// A single set of migrations to apply.
type AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrations) implementsAccountWorkerScriptUpdateParamsVariant0MetadataMigrations() {
}

type AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrationsStep] `json:"steps"`
}

func (r AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrations) implementsAccountWorkerScriptUpdateParamsVariant0MetadataMigrations() {
}

type AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r AccountWorkerScriptUpdateParamsVariant0MetadataMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptUpdateParamsVariant0MetadataPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[AccountWorkerScriptUpdateParamsVariant0MetadataPlacementMode] `json:"mode"`
}

func (r AccountWorkerScriptUpdateParamsVariant0MetadataPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type AccountWorkerScriptUpdateParamsVariant0MetadataPlacementMode string

const (
	AccountWorkerScriptUpdateParamsVariant0MetadataPlacementModeSmart AccountWorkerScriptUpdateParamsVariant0MetadataPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type AccountWorkerScriptUpdateParamsVariant0MetadataTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r AccountWorkerScriptUpdateParamsVariant0MetadataTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Usage model to apply to invocations.
type AccountWorkerScriptUpdateParamsVariant0MetadataUsageModel string

const (
	AccountWorkerScriptUpdateParamsVariant0MetadataUsageModelBundled AccountWorkerScriptUpdateParamsVariant0MetadataUsageModel = "bundled"
	AccountWorkerScriptUpdateParamsVariant0MetadataUsageModelUnbound AccountWorkerScriptUpdateParamsVariant0MetadataUsageModel = "unbound"
)

type AccountWorkerScriptUpdateParamsVariant1 struct {
	// Rollback to provided deployment based on deployment ID. Request body will only
	// parse a "message" part. You can learn more about deployments
	// [here](https://developers.cloudflare.com/workers/platform/deployments/).
	RollbackTo param.Field[string] `query:"rollback_to"`
	// Rollback message to be associated with this deployment. Only parsed when query
	// param `"rollback_to"` is present.
	Message param.Field[string] `json:"message"`
}

func (r AccountWorkerScriptUpdateParamsVariant1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountWorkerScriptUpdateParamsVariant1]'s query parameters
// as `url.Values`.
func (r AccountWorkerScriptUpdateParamsVariant1) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func (AccountWorkerScriptUpdateParamsVariant1) ImplementsAccountWorkerScriptUpdateParams() {

}

type AccountWorkerScriptDeleteParams struct {
	// If set to true, delete will not be stopped by associated service binding,
	// durable object, or other binding. Any of these associated bindings/durable
	// objects will be deleted along with the script.
	Force param.Field[bool] `query:"force"`
}

// URLQuery serializes [AccountWorkerScriptDeleteParams]'s query parameters as
// `url.Values`.
func (r AccountWorkerScriptDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
