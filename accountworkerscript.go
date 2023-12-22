// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

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
func (r *AccountWorkerScriptService) WorkerScriptListWorkers(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *ScriptResponseCollectionCLvvmMt, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ScriptResponseCollectionCLvvmMt struct {
	Errors   []ScriptResponseCollectionCLvvmMtError   `json:"errors"`
	Messages []ScriptResponseCollectionCLvvmMtMessage `json:"messages"`
	Result   []ScriptResponseCollectionCLvvmMtResult  `json:"result"`
	// Whether the API call was successful
	Success ScriptResponseCollectionCLvvmMtSuccess `json:"success"`
	JSON    scriptResponseCollectionCLvvmMtJSON    `json:"-"`
}

// scriptResponseCollectionCLvvmMtJSON contains the JSON metadata for the struct
// [ScriptResponseCollectionCLvvmMt]
type scriptResponseCollectionCLvvmMtJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptResponseCollectionCLvvmMt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ScriptResponseCollectionCLvvmMtError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    scriptResponseCollectionCLvvmMtErrorJSON `json:"-"`
}

// scriptResponseCollectionCLvvmMtErrorJSON contains the JSON metadata for the
// struct [ScriptResponseCollectionCLvvmMtError]
type scriptResponseCollectionCLvvmMtErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptResponseCollectionCLvvmMtError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ScriptResponseCollectionCLvvmMtMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    scriptResponseCollectionCLvvmMtMessageJSON `json:"-"`
}

// scriptResponseCollectionCLvvmMtMessageJSON contains the JSON metadata for the
// struct [ScriptResponseCollectionCLvvmMtMessage]
type scriptResponseCollectionCLvvmMtMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptResponseCollectionCLvvmMtMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ScriptResponseCollectionCLvvmMtResult struct {
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
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                    `json:"usage_model"`
	JSON       scriptResponseCollectionCLvvmMtResultJSON `json:"-"`
}

// scriptResponseCollectionCLvvmMtResultJSON contains the JSON metadata for the
// struct [ScriptResponseCollectionCLvvmMtResult]
type scriptResponseCollectionCLvvmMtResultJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Etag        apijson.Field
	Logpush     apijson.Field
	ModifiedOn  apijson.Field
	UsageModel  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptResponseCollectionCLvvmMtResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ScriptResponseCollectionCLvvmMtSuccess bool

const (
	ScriptResponseCollectionCLvvmMtSuccessTrue ScriptResponseCollectionCLvvmMtSuccess = true
)

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
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                      `json:"usage_model"`
	JSON       accountWorkerScriptUpdateResponseResultJSON `json:"-"`
}

// accountWorkerScriptUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountWorkerScriptUpdateResponseResult]
type accountWorkerScriptUpdateResponseResultJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Etag        apijson.Field
	Logpush     apijson.Field
	ModifiedOn  apijson.Field
	UsageModel  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerScriptUpdateResponseSuccess bool

const (
	AccountWorkerScriptUpdateResponseSuccessTrue AccountWorkerScriptUpdateResponseSuccess = true
)

type AccountWorkerScriptUpdateParams struct {
	// Rollback to provided deployment based on deployment ID. Request body will be
	// ignored. You can learn more about deployments
	// [here](https://developers.cloudflare.com/workers/platform/deployments/).
	RollbackTo param.Field[string] `query:"rollback_to" format:"uuid"`
	// Rollback message to be associated with deployment. Only parsed with
	// `rollback_to` query parameter
	Message param.Field[interface{}] `json:"\"message\""`
	// Worker script.
	SecondFileJs param.Field[string] `json:"\"second-file.js\""`
	// Worker script.
	WorkerJs param.Field[string] `json:"\"worker.js\""`
	// Metadata for script such as bindings in JSON format. Main module needs to be
	// specified with `main_module`.
	Metadata param.Field[string] `json:"metadata"`
}

func (r AccountWorkerScriptUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountWorkerScriptUpdateParams]'s query parameters as
// `url.Values`.
func (r AccountWorkerScriptUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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
