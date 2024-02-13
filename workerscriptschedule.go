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

// WorkerScriptScheduleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWorkerScriptScheduleService]
// method instead.
type WorkerScriptScheduleService struct {
	Options []option.RequestOption
}

// NewWorkerScriptScheduleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkerScriptScheduleService(opts ...option.RequestOption) (r *WorkerScriptScheduleService) {
	r = &WorkerScriptScheduleService{}
	r.Options = opts
	return
}

// Fetches Cron Triggers for a Worker.
func (r *WorkerScriptScheduleService) WorkerCronTriggerGetCronTriggers(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/schedules", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates Cron Triggers for a Worker.
func (r *WorkerScriptScheduleService) WorkerCronTriggerUpdateCronTriggers(ctx context.Context, accountID string, scriptName string, body WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersParams, opts ...option.RequestOption) (res *WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/schedules", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponse struct {
	Schedules []WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseSchedule `json:"schedules"`
	JSON      workerScriptScheduleWorkerCronTriggerGetCronTriggersResponseJSON       `json:"-"`
}

// workerScriptScheduleWorkerCronTriggerGetCronTriggersResponseJSON contains the
// JSON metadata for the struct
// [WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponse]
type workerScriptScheduleWorkerCronTriggerGetCronTriggersResponseJSON struct {
	Schedules   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseSchedule struct {
	CreatedOn  interface{}                                                              `json:"created_on"`
	Cron       interface{}                                                              `json:"cron"`
	ModifiedOn interface{}                                                              `json:"modified_on"`
	JSON       workerScriptScheduleWorkerCronTriggerGetCronTriggersResponseScheduleJSON `json:"-"`
}

// workerScriptScheduleWorkerCronTriggerGetCronTriggersResponseScheduleJSON
// contains the JSON metadata for the struct
// [WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseSchedule]
type workerScriptScheduleWorkerCronTriggerGetCronTriggersResponseScheduleJSON struct {
	CreatedOn   apijson.Field
	Cron        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponse struct {
	Schedules []WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseSchedule `json:"schedules"`
	JSON      workerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseJSON       `json:"-"`
}

// workerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseJSON contains the
// JSON metadata for the struct
// [WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponse]
type workerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseJSON struct {
	Schedules   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseSchedule struct {
	CreatedOn  interface{}                                                                 `json:"created_on"`
	Cron       interface{}                                                                 `json:"cron"`
	ModifiedOn interface{}                                                                 `json:"modified_on"`
	JSON       workerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseScheduleJSON `json:"-"`
}

// workerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseScheduleJSON
// contains the JSON metadata for the struct
// [WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseSchedule]
type workerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseScheduleJSON struct {
	CreatedOn   apijson.Field
	Cron        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelope struct {
	Errors   []WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeJSON    `json:"-"`
}

// workerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelope]
type workerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeErrors struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    workerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeErrors]
type workerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeMessages struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    workerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeMessages]
type workerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeSuccess bool

const (
	WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeSuccessTrue WorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseEnvelopeSuccess = true
)

type WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelope struct {
	Errors   []WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeJSON    `json:"-"`
}

// workerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelope]
type workerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeErrors struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    workerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeErrors]
type workerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeMessages struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    workerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeMessages]
type workerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeSuccess bool

const (
	WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeSuccessTrue WorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseEnvelopeSuccess = true
)
