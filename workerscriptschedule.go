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

// Updates Cron Triggers for a Worker.
func (r *WorkerScriptScheduleService) Update(ctx context.Context, accountID string, scriptName string, body WorkerScriptScheduleUpdateParams, opts ...option.RequestOption) (res *WorkerScriptScheduleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptScheduleUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/schedules", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Cron Triggers for a Worker.
func (r *WorkerScriptScheduleService) List(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *WorkerScriptScheduleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptScheduleListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/schedules", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerScriptScheduleUpdateResponse struct {
	Schedules []WorkerScriptScheduleUpdateResponseSchedule `json:"schedules"`
	JSON      workerScriptScheduleUpdateResponseJSON       `json:"-"`
}

// workerScriptScheduleUpdateResponseJSON contains the JSON metadata for the struct
// [WorkerScriptScheduleUpdateResponse]
type workerScriptScheduleUpdateResponseJSON struct {
	Schedules   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptScheduleUpdateResponseSchedule struct {
	CreatedOn  interface{}                                    `json:"created_on"`
	Cron       interface{}                                    `json:"cron"`
	ModifiedOn interface{}                                    `json:"modified_on"`
	JSON       workerScriptScheduleUpdateResponseScheduleJSON `json:"-"`
}

// workerScriptScheduleUpdateResponseScheduleJSON contains the JSON metadata for
// the struct [WorkerScriptScheduleUpdateResponseSchedule]
type workerScriptScheduleUpdateResponseScheduleJSON struct {
	CreatedOn   apijson.Field
	Cron        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleUpdateResponseSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptScheduleListResponse struct {
	Schedules []WorkerScriptScheduleListResponseSchedule `json:"schedules"`
	JSON      workerScriptScheduleListResponseJSON       `json:"-"`
}

// workerScriptScheduleListResponseJSON contains the JSON metadata for the struct
// [WorkerScriptScheduleListResponse]
type workerScriptScheduleListResponseJSON struct {
	Schedules   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptScheduleListResponseSchedule struct {
	CreatedOn  interface{}                                  `json:"created_on"`
	Cron       interface{}                                  `json:"cron"`
	ModifiedOn interface{}                                  `json:"modified_on"`
	JSON       workerScriptScheduleListResponseScheduleJSON `json:"-"`
}

// workerScriptScheduleListResponseScheduleJSON contains the JSON metadata for the
// struct [WorkerScriptScheduleListResponseSchedule]
type workerScriptScheduleListResponseScheduleJSON struct {
	CreatedOn   apijson.Field
	Cron        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleListResponseSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptScheduleUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WorkerScriptScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WorkerScriptScheduleUpdateResponseEnvelope struct {
	Errors   []WorkerScriptScheduleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptScheduleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptScheduleUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptScheduleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptScheduleUpdateResponseEnvelopeJSON    `json:"-"`
}

// workerScriptScheduleUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [WorkerScriptScheduleUpdateResponseEnvelope]
type workerScriptScheduleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptScheduleUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    workerScriptScheduleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptScheduleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [WorkerScriptScheduleUpdateResponseEnvelopeErrors]
type workerScriptScheduleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptScheduleUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    workerScriptScheduleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptScheduleUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [WorkerScriptScheduleUpdateResponseEnvelopeMessages]
type workerScriptScheduleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptScheduleUpdateResponseEnvelopeSuccess bool

const (
	WorkerScriptScheduleUpdateResponseEnvelopeSuccessTrue WorkerScriptScheduleUpdateResponseEnvelopeSuccess = true
)

type WorkerScriptScheduleListResponseEnvelope struct {
	Errors   []WorkerScriptScheduleListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptScheduleListResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptScheduleListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptScheduleListResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptScheduleListResponseEnvelopeJSON    `json:"-"`
}

// workerScriptScheduleListResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerScriptScheduleListResponseEnvelope]
type workerScriptScheduleListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptScheduleListResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    workerScriptScheduleListResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptScheduleListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [WorkerScriptScheduleListResponseEnvelopeErrors]
type workerScriptScheduleListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptScheduleListResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    workerScriptScheduleListResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptScheduleListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [WorkerScriptScheduleListResponseEnvelopeMessages]
type workerScriptScheduleListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptScheduleListResponseEnvelopeSuccess bool

const (
	WorkerScriptScheduleListResponseEnvelopeSuccessTrue WorkerScriptScheduleListResponseEnvelopeSuccess = true
)
