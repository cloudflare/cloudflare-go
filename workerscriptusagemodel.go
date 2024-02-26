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

// Updates the Usage Model for a given Worker. Requires a Workers Paid
// subscription.
func (r *WorkerScriptUsageModelService) Update(ctx context.Context, scriptName string, params WorkerScriptUsageModelUpdateParams, opts ...option.RequestOption) (res *WorkerScriptUsageModelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptUsageModelUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/usage-model", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the Usage Model for a given Worker.
func (r *WorkerScriptUsageModelService) Get(ctx context.Context, scriptName string, query WorkerScriptUsageModelGetParams, opts ...option.RequestOption) (res *WorkerScriptUsageModelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptUsageModelGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/usage-model", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerScriptUsageModelUpdateResponse struct {
	UsageModel interface{}                              `json:"usage_model"`
	JSON       workerScriptUsageModelUpdateResponseJSON `json:"-"`
}

// workerScriptUsageModelUpdateResponseJSON contains the JSON metadata for the
// struct [WorkerScriptUsageModelUpdateResponse]
type workerScriptUsageModelUpdateResponseJSON struct {
	UsageModel  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type WorkerScriptUsageModelUpdateParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r WorkerScriptUsageModelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WorkerScriptUsageModelUpdateResponseEnvelope struct {
	Errors   []WorkerScriptUsageModelUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptUsageModelUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptUsageModelUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptUsageModelUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptUsageModelUpdateResponseEnvelopeJSON    `json:"-"`
}

// workerScriptUsageModelUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [WorkerScriptUsageModelUpdateResponseEnvelope]
type workerScriptUsageModelUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptUsageModelUpdateResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    workerScriptUsageModelUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptUsageModelUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [WorkerScriptUsageModelUpdateResponseEnvelopeErrors]
type workerScriptUsageModelUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptUsageModelUpdateResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    workerScriptUsageModelUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptUsageModelUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [WorkerScriptUsageModelUpdateResponseEnvelopeMessages]
type workerScriptUsageModelUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUsageModelUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptUsageModelUpdateResponseEnvelopeSuccess bool

const (
	WorkerScriptUsageModelUpdateResponseEnvelopeSuccessTrue WorkerScriptUsageModelUpdateResponseEnvelopeSuccess = true
)

type WorkerScriptUsageModelGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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
