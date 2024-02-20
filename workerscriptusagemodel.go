// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// WorkerScriptUsageModelService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWorkerScriptUsageModelService]
// method instead.
type WorkerScriptUsageModelService struct {
	Options []option.RequestOption
}

// NewWorkerScriptUsageModelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkerScriptUsageModelService(opts ...option.RequestOption) (r *WorkerScriptUsageModelService) {
	r = &WorkerScriptUsageModelService{}
	r.Options = opts
	return
}

// Fetches the Usage Model for a given Worker.
func (r *WorkerScriptUsageModelService) Get(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *WorkerScriptUsageModelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptUsageModelGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/usage-model", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the Usage Model for a given Worker. Requires a Workers Paid
// subscription.
func (r *WorkerScriptUsageModelService) Replace(ctx context.Context, accountID string, scriptName string, body WorkerScriptUsageModelReplaceParams, opts ...option.RequestOption) (res *WorkerScriptUsageModelReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptUsageModelReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/usage-model", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerScriptUsageModelGetResponse struct {
	UsageModel interface{}                           `json:"usage_model"`
	JSON       workerScriptUsageModelGetResponseJSON `json:"-"`
}

// workerScriptUsageModelGetResponseJSON contains the JSON metadata for the struct
// [WorkerScriptUsageModelGetResponse]
type workerScriptUsageModelGetResponseJSON struct {
	UsageModel  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptUsageModelReplaceResponse struct {
	UsageModel interface{}                               `json:"usage_model"`
	JSON       workerScriptUsageModelReplaceResponseJSON `json:"-"`
}

// workerScriptUsageModelReplaceResponseJSON contains the JSON metadata for the
// struct [WorkerScriptUsageModelReplaceResponse]
type workerScriptUsageModelReplaceResponseJSON struct {
	UsageModel  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptUsageModelGetResponseEnvelope struct {
	Errors   []WorkerScriptUsageModelGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptUsageModelGetResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptUsageModelGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptUsageModelGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptUsageModelGetResponseEnvelopeJSON    `json:"-"`
}

// workerScriptUsageModelGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerScriptUsageModelGetResponseEnvelope]
type workerScriptUsageModelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptUsageModelGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    workerScriptUsageModelGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptUsageModelGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [WorkerScriptUsageModelGetResponseEnvelopeErrors]
type workerScriptUsageModelGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptUsageModelGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    workerScriptUsageModelGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptUsageModelGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [WorkerScriptUsageModelGetResponseEnvelopeMessages]
type workerScriptUsageModelGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptUsageModelGetResponseEnvelopeSuccess bool

const (
	WorkerScriptUsageModelGetResponseEnvelopeSuccessTrue WorkerScriptUsageModelGetResponseEnvelopeSuccess = true
)

type WorkerScriptUsageModelReplaceParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WorkerScriptUsageModelReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WorkerScriptUsageModelReplaceResponseEnvelope struct {
	Errors   []WorkerScriptUsageModelReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptUsageModelReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptUsageModelReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptUsageModelReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptUsageModelReplaceResponseEnvelopeJSON    `json:"-"`
}

// workerScriptUsageModelReplaceResponseEnvelopeJSON contains the JSON metadata for
// the struct [WorkerScriptUsageModelReplaceResponseEnvelope]
type workerScriptUsageModelReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptUsageModelReplaceResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    workerScriptUsageModelReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptUsageModelReplaceResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [WorkerScriptUsageModelReplaceResponseEnvelopeErrors]
type workerScriptUsageModelReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptUsageModelReplaceResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    workerScriptUsageModelReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptUsageModelReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [WorkerScriptUsageModelReplaceResponseEnvelopeMessages]
type workerScriptUsageModelReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptUsageModelReplaceResponseEnvelopeSuccess bool

const (
	WorkerScriptUsageModelReplaceResponseEnvelopeSuccessTrue WorkerScriptUsageModelReplaceResponseEnvelopeSuccess = true
)
