// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *WorkerScriptScheduleService) Update(ctx context.Context, scriptName string, params WorkerScriptScheduleUpdateParams, opts ...option.RequestOption) (res *WorkerScriptScheduleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptScheduleUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/schedules", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Cron Triggers for a Worker.
func (r *WorkerScriptScheduleService) Get(ctx context.Context, scriptName string, query WorkerScriptScheduleGetParams, opts ...option.RequestOption) (res *WorkerScriptScheduleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptScheduleGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/schedules", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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

func (r workerScriptScheduleUpdateResponseJSON) RawJSON() string {
	return r.raw
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

func (r workerScriptScheduleUpdateResponseScheduleJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptScheduleGetResponse struct {
	Schedules []WorkerScriptScheduleGetResponseSchedule `json:"schedules"`
	JSON      workerScriptScheduleGetResponseJSON       `json:"-"`
}

// workerScriptScheduleGetResponseJSON contains the JSON metadata for the struct
// [WorkerScriptScheduleGetResponse]
type workerScriptScheduleGetResponseJSON struct {
	Schedules   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptScheduleGetResponseJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptScheduleGetResponseSchedule struct {
	CreatedOn  interface{}                                 `json:"created_on"`
	Cron       interface{}                                 `json:"cron"`
	ModifiedOn interface{}                                 `json:"modified_on"`
	JSON       workerScriptScheduleGetResponseScheduleJSON `json:"-"`
}

// workerScriptScheduleGetResponseScheduleJSON contains the JSON metadata for the
// struct [WorkerScriptScheduleGetResponseSchedule]
type workerScriptScheduleGetResponseScheduleJSON struct {
	CreatedOn   apijson.Field
	Cron        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleGetResponseSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptScheduleGetResponseScheduleJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptScheduleUpdateParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
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

func (r workerScriptScheduleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r workerScriptScheduleUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r workerScriptScheduleUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerScriptScheduleUpdateResponseEnvelopeSuccess bool

const (
	WorkerScriptScheduleUpdateResponseEnvelopeSuccessTrue WorkerScriptScheduleUpdateResponseEnvelopeSuccess = true
)

type WorkerScriptScheduleGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type WorkerScriptScheduleGetResponseEnvelope struct {
	Errors   []WorkerScriptScheduleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptScheduleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptScheduleGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptScheduleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptScheduleGetResponseEnvelopeJSON    `json:"-"`
}

// workerScriptScheduleGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerScriptScheduleGetResponseEnvelope]
type workerScriptScheduleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptScheduleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptScheduleGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    workerScriptScheduleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptScheduleGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [WorkerScriptScheduleGetResponseEnvelopeErrors]
type workerScriptScheduleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptScheduleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptScheduleGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    workerScriptScheduleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptScheduleGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [WorkerScriptScheduleGetResponseEnvelopeMessages]
type workerScriptScheduleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptScheduleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptScheduleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerScriptScheduleGetResponseEnvelopeSuccess bool

const (
	WorkerScriptScheduleGetResponseEnvelopeSuccessTrue WorkerScriptScheduleGetResponseEnvelopeSuccess = true
)
