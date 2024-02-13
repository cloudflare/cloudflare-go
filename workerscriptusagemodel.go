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
func (r *WorkerScriptUsageModelService) WorkerScriptFetchUsageModel(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *WorkerScriptUsageModelWorkerScriptFetchUsageModelResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelope
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
func (r *WorkerScriptUsageModelService) WorkerScriptUpdateUsageModel(ctx context.Context, accountID string, scriptName string, body WorkerScriptUsageModelWorkerScriptUpdateUsageModelParams, opts ...option.RequestOption) (res *WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/usage-model", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerScriptUsageModelWorkerScriptFetchUsageModelResponse struct {
	UsageModel interface{}                                                   `json:"usage_model"`
	JSON       workerScriptUsageModelWorkerScriptFetchUsageModelResponseJSON `json:"-"`
}

// workerScriptUsageModelWorkerScriptFetchUsageModelResponseJSON contains the JSON
// metadata for the struct
// [WorkerScriptUsageModelWorkerScriptFetchUsageModelResponse]
type workerScriptUsageModelWorkerScriptFetchUsageModelResponseJSON struct {
	UsageModel  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelWorkerScriptFetchUsageModelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponse struct {
	UsageModel interface{}                                                    `json:"usage_model"`
	JSON       workerScriptUsageModelWorkerScriptUpdateUsageModelResponseJSON `json:"-"`
}

// workerScriptUsageModelWorkerScriptUpdateUsageModelResponseJSON contains the JSON
// metadata for the struct
// [WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponse]
type workerScriptUsageModelWorkerScriptUpdateUsageModelResponseJSON struct {
	UsageModel  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelope struct {
	Errors   []WorkerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptUsageModelWorkerScriptFetchUsageModelResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeJSON    `json:"-"`
}

// workerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [WorkerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelope]
type workerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeErrors struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    workerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [WorkerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeErrors]
type workerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeMessages struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    workerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [WorkerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeMessages]
type workerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeSuccess bool

const (
	WorkerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeSuccessTrue WorkerScriptUsageModelWorkerScriptFetchUsageModelResponseEnvelopeSuccess = true
)

type WorkerScriptUsageModelWorkerScriptUpdateUsageModelParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WorkerScriptUsageModelWorkerScriptUpdateUsageModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelope struct {
	Errors   []WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeJSON    `json:"-"`
}

// workerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelope]
type workerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeErrors struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    workerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeErrors]
type workerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeMessages struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    workerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeMessages]
type workerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeSuccess bool

const (
	WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeSuccessTrue WorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseEnvelopeSuccess = true
)
