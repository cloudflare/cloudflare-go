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

// WorkerQueueService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWorkerQueueService] method
// instead.
type WorkerQueueService struct {
	Options   []option.RequestOption
	Consumers *WorkerQueueConsumerService
}

// NewWorkerQueueService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkerQueueService(opts ...option.RequestOption) (r *WorkerQueueService) {
	r = &WorkerQueueService{}
	r.Options = opts
	r.Consumers = NewWorkerQueueConsumerService(opts...)
	return
}

// Updates a queue.
func (r *WorkerQueueService) Update(ctx context.Context, accountID string, name string, body WorkerQueueUpdateParams, opts ...option.RequestOption) (res *WorkerQueueUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerQueueUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/queues/%s", accountID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the queues owned by an account.
func (r *WorkerQueueService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]WorkerQueueListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerQueueListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/queues", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a queue.
func (r *WorkerQueueService) Delete(ctx context.Context, accountID string, name string, opts ...option.RequestOption) (res *WorkerQueueDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerQueueDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/queues/%s", accountID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a specific queue.
func (r *WorkerQueueService) Get(ctx context.Context, accountID string, name string, opts ...option.RequestOption) (res *WorkerQueueGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerQueueGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/queues/%s", accountID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a new queue.
func (r *WorkerQueueService) QueueNewQueue(ctx context.Context, accountID string, body WorkerQueueQueueNewQueueParams, opts ...option.RequestOption) (res *WorkerQueueQueueNewQueueResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerQueueQueueNewQueueResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/queues", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerQueueUpdateResponse struct {
	CreatedOn  interface{}                   `json:"created_on"`
	ModifiedOn interface{}                   `json:"modified_on"`
	QueueID    interface{}                   `json:"queue_id"`
	QueueName  string                        `json:"queue_name"`
	JSON       workerQueueUpdateResponseJSON `json:"-"`
}

// workerQueueUpdateResponseJSON contains the JSON metadata for the struct
// [WorkerQueueUpdateResponse]
type workerQueueUpdateResponseJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	QueueID     apijson.Field
	QueueName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueListResponse struct {
	Consumers           interface{}                 `json:"consumers"`
	ConsumersTotalCount interface{}                 `json:"consumers_total_count"`
	CreatedOn           interface{}                 `json:"created_on"`
	ModifiedOn          interface{}                 `json:"modified_on"`
	Producers           interface{}                 `json:"producers"`
	ProducersTotalCount interface{}                 `json:"producers_total_count"`
	QueueID             interface{}                 `json:"queue_id"`
	QueueName           string                      `json:"queue_name"`
	JSON                workerQueueListResponseJSON `json:"-"`
}

// workerQueueListResponseJSON contains the JSON metadata for the struct
// [WorkerQueueListResponse]
type workerQueueListResponseJSON struct {
	Consumers           apijson.Field
	ConsumersTotalCount apijson.Field
	CreatedOn           apijson.Field
	ModifiedOn          apijson.Field
	Producers           apijson.Field
	ProducersTotalCount apijson.Field
	QueueID             apijson.Field
	QueueName           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *WorkerQueueListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [WorkerQueueDeleteResponseUnknown],
// [WorkerQueueDeleteResponseArray] or [shared.UnionString].
type WorkerQueueDeleteResponse interface {
	ImplementsWorkerQueueDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WorkerQueueDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type WorkerQueueDeleteResponseArray []interface{}

func (r WorkerQueueDeleteResponseArray) ImplementsWorkerQueueDeleteResponse() {}

type WorkerQueueGetResponse struct {
	Consumers           interface{}                `json:"consumers"`
	ConsumersTotalCount interface{}                `json:"consumers_total_count"`
	CreatedOn           interface{}                `json:"created_on"`
	ModifiedOn          interface{}                `json:"modified_on"`
	Producers           interface{}                `json:"producers"`
	ProducersTotalCount interface{}                `json:"producers_total_count"`
	QueueID             interface{}                `json:"queue_id"`
	QueueName           string                     `json:"queue_name"`
	JSON                workerQueueGetResponseJSON `json:"-"`
}

// workerQueueGetResponseJSON contains the JSON metadata for the struct
// [WorkerQueueGetResponse]
type workerQueueGetResponseJSON struct {
	Consumers           apijson.Field
	ConsumersTotalCount apijson.Field
	CreatedOn           apijson.Field
	ModifiedOn          apijson.Field
	Producers           apijson.Field
	ProducersTotalCount apijson.Field
	QueueID             apijson.Field
	QueueName           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *WorkerQueueGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueQueueNewQueueResponse struct {
	CreatedOn  interface{}                          `json:"created_on"`
	ModifiedOn interface{}                          `json:"modified_on"`
	QueueID    interface{}                          `json:"queue_id"`
	QueueName  string                               `json:"queue_name"`
	JSON       workerQueueQueueNewQueueResponseJSON `json:"-"`
}

// workerQueueQueueNewQueueResponseJSON contains the JSON metadata for the struct
// [WorkerQueueQueueNewQueueResponse]
type workerQueueQueueNewQueueResponseJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	QueueID     apijson.Field
	QueueName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueQueueNewQueueResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WorkerQueueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WorkerQueueUpdateResponseEnvelope struct {
	Errors   []WorkerQueueUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerQueueUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerQueueUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WorkerQueueUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WorkerQueueUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       workerQueueUpdateResponseEnvelopeJSON       `json:"-"`
}

// workerQueueUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerQueueUpdateResponseEnvelope]
type workerQueueUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueUpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    workerQueueUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// workerQueueUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerQueueUpdateResponseEnvelopeErrors]
type workerQueueUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueUpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    workerQueueUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// workerQueueUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WorkerQueueUpdateResponseEnvelopeMessages]
type workerQueueUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerQueueUpdateResponseEnvelopeSuccess bool

const (
	WorkerQueueUpdateResponseEnvelopeSuccessTrue WorkerQueueUpdateResponseEnvelopeSuccess = true
)

type WorkerQueueUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       workerQueueUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// workerQueueUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [WorkerQueueUpdateResponseEnvelopeResultInfo]
type workerQueueUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueListResponseEnvelope struct {
	Errors   []interface{}             `json:"errors,required,nullable"`
	Messages []interface{}             `json:"messages,required,nullable"`
	Result   []WorkerQueueListResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WorkerQueueListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WorkerQueueListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       workerQueueListResponseEnvelopeJSON       `json:"-"`
}

// workerQueueListResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerQueueListResponseEnvelope]
type workerQueueListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerQueueListResponseEnvelopeSuccess bool

const (
	WorkerQueueListResponseEnvelopeSuccessTrue WorkerQueueListResponseEnvelopeSuccess = true
)

type WorkerQueueListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	TotalPages interface{}                                   `json:"total_pages"`
	JSON       workerQueueListResponseEnvelopeResultInfoJSON `json:"-"`
}

// workerQueueListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [WorkerQueueListResponseEnvelopeResultInfo]
type workerQueueListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueDeleteResponseEnvelope struct {
	Errors   []WorkerQueueDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerQueueDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerQueueDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WorkerQueueDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WorkerQueueDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       workerQueueDeleteResponseEnvelopeJSON       `json:"-"`
}

// workerQueueDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerQueueDeleteResponseEnvelope]
type workerQueueDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueDeleteResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    workerQueueDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// workerQueueDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerQueueDeleteResponseEnvelopeErrors]
type workerQueueDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueDeleteResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    workerQueueDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// workerQueueDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WorkerQueueDeleteResponseEnvelopeMessages]
type workerQueueDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerQueueDeleteResponseEnvelopeSuccess bool

const (
	WorkerQueueDeleteResponseEnvelopeSuccessTrue WorkerQueueDeleteResponseEnvelopeSuccess = true
)

type WorkerQueueDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       workerQueueDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// workerQueueDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [WorkerQueueDeleteResponseEnvelopeResultInfo]
type workerQueueDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueGetResponseEnvelope struct {
	Errors   []WorkerQueueGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerQueueGetResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerQueueGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WorkerQueueGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WorkerQueueGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       workerQueueGetResponseEnvelopeJSON       `json:"-"`
}

// workerQueueGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerQueueGetResponseEnvelope]
type workerQueueGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    workerQueueGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workerQueueGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerQueueGetResponseEnvelopeErrors]
type workerQueueGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    workerQueueGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workerQueueGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WorkerQueueGetResponseEnvelopeMessages]
type workerQueueGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerQueueGetResponseEnvelopeSuccess bool

const (
	WorkerQueueGetResponseEnvelopeSuccessTrue WorkerQueueGetResponseEnvelopeSuccess = true
)

type WorkerQueueGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                      `json:"total_count"`
	JSON       workerQueueGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// workerQueueGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [WorkerQueueGetResponseEnvelopeResultInfo]
type workerQueueGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueQueueNewQueueParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WorkerQueueQueueNewQueueParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WorkerQueueQueueNewQueueResponseEnvelope struct {
	Errors   []WorkerQueueQueueNewQueueResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerQueueQueueNewQueueResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerQueueQueueNewQueueResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WorkerQueueQueueNewQueueResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WorkerQueueQueueNewQueueResponseEnvelopeResultInfo `json:"result_info"`
	JSON       workerQueueQueueNewQueueResponseEnvelopeJSON       `json:"-"`
}

// workerQueueQueueNewQueueResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerQueueQueueNewQueueResponseEnvelope]
type workerQueueQueueNewQueueResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueQueueNewQueueResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueQueueNewQueueResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    workerQueueQueueNewQueueResponseEnvelopeErrorsJSON `json:"-"`
}

// workerQueueQueueNewQueueResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [WorkerQueueQueueNewQueueResponseEnvelopeErrors]
type workerQueueQueueNewQueueResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueQueueNewQueueResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerQueueQueueNewQueueResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    workerQueueQueueNewQueueResponseEnvelopeMessagesJSON `json:"-"`
}

// workerQueueQueueNewQueueResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [WorkerQueueQueueNewQueueResponseEnvelopeMessages]
type workerQueueQueueNewQueueResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueQueueNewQueueResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerQueueQueueNewQueueResponseEnvelopeSuccess bool

const (
	WorkerQueueQueueNewQueueResponseEnvelopeSuccessTrue WorkerQueueQueueNewQueueResponseEnvelopeSuccess = true
)

type WorkerQueueQueueNewQueueResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       workerQueueQueueNewQueueResponseEnvelopeResultInfoJSON `json:"-"`
}

// workerQueueQueueNewQueueResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [WorkerQueueQueueNewQueueResponseEnvelopeResultInfo]
type workerQueueQueueNewQueueResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerQueueQueueNewQueueResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
