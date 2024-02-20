// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// WorkerQueueConsumerService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWorkerQueueConsumerService]
// method instead.
type WorkerQueueConsumerService struct {
	Options []option.RequestOption
}

// NewWorkerQueueConsumerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkerQueueConsumerService(opts ...option.RequestOption) (r *WorkerQueueConsumerService) {
	r = &WorkerQueueConsumerService{}
	r.Options = opts
	return
}

// Creates a new consumer for a queue.
func (r *WorkerQueueConsumerService) New(ctx context.Context, accountID string, name string, body WorkerQueueConsumerNewParams, opts ...option.RequestOption) (res *WorkerQueueConsumerNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerQueueConsumerNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/queues/%s/consumers", accountID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the consumer for a queue, or creates one if it does not exist.
func (r *WorkerQueueConsumerService) Update(ctx context.Context, accountID string, name string, consumerName string, body WorkerQueueConsumerUpdateParams, opts ...option.RequestOption) (res *WorkerQueueConsumerUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerQueueConsumerUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/queues/%s/consumers/%s", accountID, name, consumerName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the consumers for a queue.
func (r *WorkerQueueConsumerService) List(ctx context.Context, accountID string, name string, opts ...option.RequestOption) (res *[]WorkerQueueConsumerListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerQueueConsumerListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/queues/%s/consumers", accountID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes the consumer for a queue.
func (r *WorkerQueueConsumerService) Delete(ctx context.Context, accountID string, name string, consumerName string, opts ...option.RequestOption) (res *WorkerQueueConsumerDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerQueueConsumerDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/queues/%s/consumers/%s", accountID, name, consumerName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerQueueConsumerNewResponse struct {
	CreatedOn       interface{}                            `json:"created_on"`
	DeadLetterQueue string                                 `json:"dead_letter_queue"`
	Environment     interface{}                            `json:"environment"`
	QueueName       interface{}                            `json:"queue_name"`
	ScriptName      interface{}                            `json:"script_name"`
	Settings        WorkerQueueConsumerNewResponseSettings `json:"settings"`
	JSON            workerQueueConsumerNewResponseJSON     `json:"-"`
}

// workerQueueConsumerNewResponseJSON contains the JSON metadata for the struct
// [WorkerQueueConsumerNewResponse]
type workerQueueConsumerNewResponseJSON struct {
	CreatedOn       apijson.Field
	DeadLetterQueue apijson.Field
	Environment     apijson.Field
	QueueName       apijson.Field
	ScriptName      apijson.Field
	Settings        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WorkerQueueConsumerNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueConsumerNewResponseSettings struct {
	BatchSize     float64                                    `json:"batch_size"`
	MaxRetries    float64                                    `json:"max_retries"`
	MaxWaitTimeMs float64                                    `json:"max_wait_time_ms"`
	JSON          workerQueueConsumerNewResponseSettingsJSON `json:"-"`
}

// workerQueueConsumerNewResponseSettingsJSON contains the JSON metadata for the
// struct [WorkerQueueConsumerNewResponseSettings]
type workerQueueConsumerNewResponseSettingsJSON struct {
	BatchSize     apijson.Field
	MaxRetries    apijson.Field
	MaxWaitTimeMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerQueueConsumerNewResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueConsumerUpdateResponse struct {
	CreatedOn       interface{}                               `json:"created_on"`
	DeadLetterQueue interface{}                               `json:"dead_letter_queue"`
	Environment     interface{}                               `json:"environment"`
	QueueName       interface{}                               `json:"queue_name"`
	ScriptName      interface{}                               `json:"script_name"`
	Settings        WorkerQueueConsumerUpdateResponseSettings `json:"settings"`
	JSON            workerQueueConsumerUpdateResponseJSON     `json:"-"`
}

// workerQueueConsumerUpdateResponseJSON contains the JSON metadata for the struct
// [WorkerQueueConsumerUpdateResponse]
type workerQueueConsumerUpdateResponseJSON struct {
	CreatedOn       apijson.Field
	DeadLetterQueue apijson.Field
	Environment     apijson.Field
	QueueName       apijson.Field
	ScriptName      apijson.Field
	Settings        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WorkerQueueConsumerUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueConsumerUpdateResponseSettings struct {
	BatchSize     float64                                       `json:"batch_size"`
	MaxRetries    float64                                       `json:"max_retries"`
	MaxWaitTimeMs float64                                       `json:"max_wait_time_ms"`
	JSON          workerQueueConsumerUpdateResponseSettingsJSON `json:"-"`
}

// workerQueueConsumerUpdateResponseSettingsJSON contains the JSON metadata for the
// struct [WorkerQueueConsumerUpdateResponseSettings]
type workerQueueConsumerUpdateResponseSettingsJSON struct {
	BatchSize     apijson.Field
	MaxRetries    apijson.Field
	MaxWaitTimeMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerQueueConsumerUpdateResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueConsumerListResponse struct {
	CreatedOn   interface{}                             `json:"created_on"`
	Environment interface{}                             `json:"environment"`
	QueueName   interface{}                             `json:"queue_name"`
	Service     interface{}                             `json:"service"`
	Settings    WorkerQueueConsumerListResponseSettings `json:"settings"`
	JSON        workerQueueConsumerListResponseJSON     `json:"-"`
}

// workerQueueConsumerListResponseJSON contains the JSON metadata for the struct
// [WorkerQueueConsumerListResponse]
type workerQueueConsumerListResponseJSON struct {
	CreatedOn   apijson.Field
	Environment apijson.Field
	QueueName   apijson.Field
	Service     apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueConsumerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueConsumerListResponseSettings struct {
	BatchSize     float64                                     `json:"batch_size"`
	MaxRetries    float64                                     `json:"max_retries"`
	MaxWaitTimeMs float64                                     `json:"max_wait_time_ms"`
	JSON          workerQueueConsumerListResponseSettingsJSON `json:"-"`
}

// workerQueueConsumerListResponseSettingsJSON contains the JSON metadata for the
// struct [WorkerQueueConsumerListResponseSettings]
type workerQueueConsumerListResponseSettingsJSON struct {
	BatchSize     apijson.Field
	MaxRetries    apijson.Field
	MaxWaitTimeMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerQueueConsumerListResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [WorkerQueueConsumerDeleteResponseUnknown],
// [WorkerQueueConsumerDeleteResponseArray] or [shared.UnionString].
type WorkerQueueConsumerDeleteResponse interface {
	ImplementsWorkerQueueConsumerDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WorkerQueueConsumerDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type WorkerQueueConsumerDeleteResponseArray []interface{}

func (r WorkerQueueConsumerDeleteResponseArray) ImplementsWorkerQueueConsumerDeleteResponse() {}

type WorkerQueueConsumerNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WorkerQueueConsumerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WorkerQueueConsumerNewResponseEnvelope struct {
	Errors   []WorkerQueueConsumerNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerQueueConsumerNewResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerQueueConsumerNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WorkerQueueConsumerNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WorkerQueueConsumerNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       workerQueueConsumerNewResponseEnvelopeJSON       `json:"-"`
}

// workerQueueConsumerNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerQueueConsumerNewResponseEnvelope]
type workerQueueConsumerNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueConsumerNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueConsumerNewResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    workerQueueConsumerNewResponseEnvelopeErrorsJSON `json:"-"`
}

// workerQueueConsumerNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [WorkerQueueConsumerNewResponseEnvelopeErrors]
type workerQueueConsumerNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueConsumerNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueConsumerNewResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    workerQueueConsumerNewResponseEnvelopeMessagesJSON `json:"-"`
}

// workerQueueConsumerNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [WorkerQueueConsumerNewResponseEnvelopeMessages]
type workerQueueConsumerNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueConsumerNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerQueueConsumerNewResponseEnvelopeSuccess bool

const (
	WorkerQueueConsumerNewResponseEnvelopeSuccessTrue WorkerQueueConsumerNewResponseEnvelopeSuccess = true
)

type WorkerQueueConsumerNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       workerQueueConsumerNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// workerQueueConsumerNewResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [WorkerQueueConsumerNewResponseEnvelopeResultInfo]
type workerQueueConsumerNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueConsumerNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueConsumerUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WorkerQueueConsumerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WorkerQueueConsumerUpdateResponseEnvelope struct {
	Errors   []WorkerQueueConsumerUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerQueueConsumerUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerQueueConsumerUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WorkerQueueConsumerUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WorkerQueueConsumerUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       workerQueueConsumerUpdateResponseEnvelopeJSON       `json:"-"`
}

// workerQueueConsumerUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerQueueConsumerUpdateResponseEnvelope]
type workerQueueConsumerUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueConsumerUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueConsumerUpdateResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    workerQueueConsumerUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// workerQueueConsumerUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [WorkerQueueConsumerUpdateResponseEnvelopeErrors]
type workerQueueConsumerUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueConsumerUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueConsumerUpdateResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    workerQueueConsumerUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// workerQueueConsumerUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [WorkerQueueConsumerUpdateResponseEnvelopeMessages]
type workerQueueConsumerUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueConsumerUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerQueueConsumerUpdateResponseEnvelopeSuccess bool

const (
	WorkerQueueConsumerUpdateResponseEnvelopeSuccessTrue WorkerQueueConsumerUpdateResponseEnvelopeSuccess = true
)

type WorkerQueueConsumerUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                 `json:"total_count"`
	JSON       workerQueueConsumerUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// workerQueueConsumerUpdateResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [WorkerQueueConsumerUpdateResponseEnvelopeResultInfo]
type workerQueueConsumerUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueConsumerUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueConsumerListResponseEnvelope struct {
	Errors   []interface{}                     `json:"errors,required,nullable"`
	Messages []interface{}                     `json:"messages,required,nullable"`
	Result   []WorkerQueueConsumerListResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WorkerQueueConsumerListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WorkerQueueConsumerListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       workerQueueConsumerListResponseEnvelopeJSON       `json:"-"`
}

// workerQueueConsumerListResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerQueueConsumerListResponseEnvelope]
type workerQueueConsumerListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueConsumerListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerQueueConsumerListResponseEnvelopeSuccess bool

const (
	WorkerQueueConsumerListResponseEnvelopeSuccessTrue WorkerQueueConsumerListResponseEnvelopeSuccess = true
)

type WorkerQueueConsumerListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	TotalPages interface{}                                           `json:"total_pages"`
	JSON       workerQueueConsumerListResponseEnvelopeResultInfoJSON `json:"-"`
}

// workerQueueConsumerListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [WorkerQueueConsumerListResponseEnvelopeResultInfo]
type workerQueueConsumerListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueConsumerListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueConsumerDeleteResponseEnvelope struct {
	Errors   []WorkerQueueConsumerDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerQueueConsumerDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerQueueConsumerDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WorkerQueueConsumerDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WorkerQueueConsumerDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       workerQueueConsumerDeleteResponseEnvelopeJSON       `json:"-"`
}

// workerQueueConsumerDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerQueueConsumerDeleteResponseEnvelope]
type workerQueueConsumerDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueConsumerDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueConsumerDeleteResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    workerQueueConsumerDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// workerQueueConsumerDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [WorkerQueueConsumerDeleteResponseEnvelopeErrors]
type workerQueueConsumerDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueConsumerDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueConsumerDeleteResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    workerQueueConsumerDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// workerQueueConsumerDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [WorkerQueueConsumerDeleteResponseEnvelopeMessages]
type workerQueueConsumerDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueConsumerDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerQueueConsumerDeleteResponseEnvelopeSuccess bool

const (
	WorkerQueueConsumerDeleteResponseEnvelopeSuccessTrue WorkerQueueConsumerDeleteResponseEnvelopeSuccess = true
)

type WorkerQueueConsumerDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                 `json:"total_count"`
	JSON       workerQueueConsumerDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// workerQueueConsumerDeleteResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [WorkerQueueConsumerDeleteResponseEnvelopeResultInfo]
type workerQueueConsumerDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueConsumerDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
