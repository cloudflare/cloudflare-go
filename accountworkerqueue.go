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

// AccountWorkerQueueService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountWorkerQueueService] method
// instead.
type AccountWorkerQueueService struct {
	Options   []option.RequestOption
	Consumers *AccountWorkerQueueConsumerService
}

// NewAccountWorkerQueueService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountWorkerQueueService(opts ...option.RequestOption) (r *AccountWorkerQueueService) {
	r = &AccountWorkerQueueService{}
	r.Options = opts
	r.Consumers = NewAccountWorkerQueueConsumerService(opts...)
	return
}

// Get information about a specific queue.
func (r *AccountWorkerQueueService) Get(ctx context.Context, accountIdentifier string, name string, opts ...option.RequestOption) (res *AccountWorkerQueueGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/queues/%s", accountIdentifier, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a queue.
func (r *AccountWorkerQueueService) Update(ctx context.Context, accountIdentifier string, name string, body AccountWorkerQueueUpdateParams, opts ...option.RequestOption) (res *AccountWorkerQueueUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/queues/%s", accountIdentifier, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns the queues owned by an account.
func (r *AccountWorkerQueueService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountWorkerQueueListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/queues", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a queue.
func (r *AccountWorkerQueueService) Delete(ctx context.Context, accountIdentifier string, name string, opts ...option.RequestOption) (res *AccountWorkerQueueDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/queues/%s", accountIdentifier, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new queue.
func (r *AccountWorkerQueueService) QueueNewQueue(ctx context.Context, accountIdentifier string, body AccountWorkerQueueQueueNewQueueParams, opts ...option.RequestOption) (res *AccountWorkerQueueQueueNewQueueResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/queues", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountWorkerQueueGetResponse struct {
	Errors     []AccountWorkerQueueGetResponseError    `json:"errors"`
	Messages   []AccountWorkerQueueGetResponseMessage  `json:"messages"`
	Result     AccountWorkerQueueGetResponseResult     `json:"result"`
	ResultInfo AccountWorkerQueueGetResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountWorkerQueueGetResponseSuccess `json:"success"`
	JSON    accountWorkerQueueGetResponseJSON    `json:"-"`
}

// accountWorkerQueueGetResponseJSON contains the JSON metadata for the struct
// [AccountWorkerQueueGetResponse]
type accountWorkerQueueGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueGetResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountWorkerQueueGetResponseErrorJSON `json:"-"`
}

// accountWorkerQueueGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountWorkerQueueGetResponseError]
type accountWorkerQueueGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueGetResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountWorkerQueueGetResponseMessageJSON `json:"-"`
}

// accountWorkerQueueGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountWorkerQueueGetResponseMessage]
type accountWorkerQueueGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueGetResponseResult struct {
	Consumers           interface{}                             `json:"consumers"`
	ConsumersTotalCount interface{}                             `json:"consumers_total_count"`
	CreatedOn           interface{}                             `json:"created_on"`
	ModifiedOn          interface{}                             `json:"modified_on"`
	Producers           interface{}                             `json:"producers"`
	ProducersTotalCount interface{}                             `json:"producers_total_count"`
	QueueID             interface{}                             `json:"queue_id"`
	QueueName           string                                  `json:"queue_name"`
	JSON                accountWorkerQueueGetResponseResultJSON `json:"-"`
}

// accountWorkerQueueGetResponseResultJSON contains the JSON metadata for the
// struct [AccountWorkerQueueGetResponseResult]
type accountWorkerQueueGetResponseResultJSON struct {
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

func (r *AccountWorkerQueueGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueGetResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       accountWorkerQueueGetResponseResultInfoJSON `json:"-"`
}

// accountWorkerQueueGetResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountWorkerQueueGetResponseResultInfo]
type accountWorkerQueueGetResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueGetResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerQueueGetResponseSuccess bool

const (
	AccountWorkerQueueGetResponseSuccessTrue AccountWorkerQueueGetResponseSuccess = true
)

type AccountWorkerQueueUpdateResponse struct {
	Errors     []AccountWorkerQueueUpdateResponseError    `json:"errors"`
	Messages   []AccountWorkerQueueUpdateResponseMessage  `json:"messages"`
	Result     AccountWorkerQueueUpdateResponseResult     `json:"result"`
	ResultInfo AccountWorkerQueueUpdateResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountWorkerQueueUpdateResponseSuccess `json:"success"`
	JSON    accountWorkerQueueUpdateResponseJSON    `json:"-"`
}

// accountWorkerQueueUpdateResponseJSON contains the JSON metadata for the struct
// [AccountWorkerQueueUpdateResponse]
type accountWorkerQueueUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueUpdateResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountWorkerQueueUpdateResponseErrorJSON `json:"-"`
}

// accountWorkerQueueUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountWorkerQueueUpdateResponseError]
type accountWorkerQueueUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueUpdateResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountWorkerQueueUpdateResponseMessageJSON `json:"-"`
}

// accountWorkerQueueUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountWorkerQueueUpdateResponseMessage]
type accountWorkerQueueUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueUpdateResponseResult struct {
	CreatedOn  interface{}                                `json:"created_on"`
	ModifiedOn interface{}                                `json:"modified_on"`
	QueueID    interface{}                                `json:"queue_id"`
	QueueName  string                                     `json:"queue_name"`
	JSON       accountWorkerQueueUpdateResponseResultJSON `json:"-"`
}

// accountWorkerQueueUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountWorkerQueueUpdateResponseResult]
type accountWorkerQueueUpdateResponseResultJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	QueueID     apijson.Field
	QueueName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueUpdateResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       accountWorkerQueueUpdateResponseResultInfoJSON `json:"-"`
}

// accountWorkerQueueUpdateResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountWorkerQueueUpdateResponseResultInfo]
type accountWorkerQueueUpdateResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueUpdateResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerQueueUpdateResponseSuccess bool

const (
	AccountWorkerQueueUpdateResponseSuccessTrue AccountWorkerQueueUpdateResponseSuccess = true
)

type AccountWorkerQueueListResponse struct {
	Errors     interface{}                              `json:"errors,nullable"`
	Messages   interface{}                              `json:"messages,nullable"`
	Result     []AccountWorkerQueueListResponseResult   `json:"result"`
	ResultInfo AccountWorkerQueueListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountWorkerQueueListResponseSuccess `json:"success"`
	JSON    accountWorkerQueueListResponseJSON    `json:"-"`
}

// accountWorkerQueueListResponseJSON contains the JSON metadata for the struct
// [AccountWorkerQueueListResponse]
type accountWorkerQueueListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueListResponseResult struct {
	Consumers           interface{}                              `json:"consumers"`
	ConsumersTotalCount interface{}                              `json:"consumers_total_count"`
	CreatedOn           interface{}                              `json:"created_on"`
	ModifiedOn          interface{}                              `json:"modified_on"`
	Producers           interface{}                              `json:"producers"`
	ProducersTotalCount interface{}                              `json:"producers_total_count"`
	QueueID             interface{}                              `json:"queue_id"`
	QueueName           string                                   `json:"queue_name"`
	JSON                accountWorkerQueueListResponseResultJSON `json:"-"`
}

// accountWorkerQueueListResponseResultJSON contains the JSON metadata for the
// struct [AccountWorkerQueueListResponseResult]
type accountWorkerQueueListResponseResultJSON struct {
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

func (r *AccountWorkerQueueListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueListResponseResultInfo struct {
	Count      interface{}                                  `json:"count"`
	Page       interface{}                                  `json:"page"`
	PerPage    interface{}                                  `json:"per_page"`
	TotalCount interface{}                                  `json:"total_count"`
	TotalPages interface{}                                  `json:"total_pages"`
	JSON       accountWorkerQueueListResponseResultInfoJSON `json:"-"`
}

// accountWorkerQueueListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountWorkerQueueListResponseResultInfo]
type accountWorkerQueueListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerQueueListResponseSuccess bool

const (
	AccountWorkerQueueListResponseSuccessTrue AccountWorkerQueueListResponseSuccess = true
)

type AccountWorkerQueueDeleteResponse struct {
	Errors     []AccountWorkerQueueDeleteResponseError    `json:"errors"`
	Messages   []AccountWorkerQueueDeleteResponseMessage  `json:"messages"`
	Result     interface{}                                `json:"result,nullable"`
	ResultInfo AccountWorkerQueueDeleteResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountWorkerQueueDeleteResponseSuccess `json:"success"`
	JSON    accountWorkerQueueDeleteResponseJSON    `json:"-"`
}

// accountWorkerQueueDeleteResponseJSON contains the JSON metadata for the struct
// [AccountWorkerQueueDeleteResponse]
type accountWorkerQueueDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueDeleteResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountWorkerQueueDeleteResponseErrorJSON `json:"-"`
}

// accountWorkerQueueDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountWorkerQueueDeleteResponseError]
type accountWorkerQueueDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueDeleteResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountWorkerQueueDeleteResponseMessageJSON `json:"-"`
}

// accountWorkerQueueDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountWorkerQueueDeleteResponseMessage]
type accountWorkerQueueDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       accountWorkerQueueDeleteResponseResultInfoJSON `json:"-"`
}

// accountWorkerQueueDeleteResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountWorkerQueueDeleteResponseResultInfo]
type accountWorkerQueueDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerQueueDeleteResponseSuccess bool

const (
	AccountWorkerQueueDeleteResponseSuccessTrue AccountWorkerQueueDeleteResponseSuccess = true
)

type AccountWorkerQueueQueueNewQueueResponse struct {
	Errors     []AccountWorkerQueueQueueNewQueueResponseError    `json:"errors"`
	Messages   []AccountWorkerQueueQueueNewQueueResponseMessage  `json:"messages"`
	Result     AccountWorkerQueueQueueNewQueueResponseResult     `json:"result"`
	ResultInfo AccountWorkerQueueQueueNewQueueResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountWorkerQueueQueueNewQueueResponseSuccess `json:"success"`
	JSON    accountWorkerQueueQueueNewQueueResponseJSON    `json:"-"`
}

// accountWorkerQueueQueueNewQueueResponseJSON contains the JSON metadata for the
// struct [AccountWorkerQueueQueueNewQueueResponse]
type accountWorkerQueueQueueNewQueueResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueQueueNewQueueResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueQueueNewQueueResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountWorkerQueueQueueNewQueueResponseErrorJSON `json:"-"`
}

// accountWorkerQueueQueueNewQueueResponseErrorJSON contains the JSON metadata for
// the struct [AccountWorkerQueueQueueNewQueueResponseError]
type accountWorkerQueueQueueNewQueueResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueQueueNewQueueResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueQueueNewQueueResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accountWorkerQueueQueueNewQueueResponseMessageJSON `json:"-"`
}

// accountWorkerQueueQueueNewQueueResponseMessageJSON contains the JSON metadata
// for the struct [AccountWorkerQueueQueueNewQueueResponseMessage]
type accountWorkerQueueQueueNewQueueResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueQueueNewQueueResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueQueueNewQueueResponseResult struct {
	CreatedOn  interface{}                                       `json:"created_on"`
	ModifiedOn interface{}                                       `json:"modified_on"`
	QueueID    interface{}                                       `json:"queue_id"`
	QueueName  string                                            `json:"queue_name"`
	JSON       accountWorkerQueueQueueNewQueueResponseResultJSON `json:"-"`
}

// accountWorkerQueueQueueNewQueueResponseResultJSON contains the JSON metadata for
// the struct [AccountWorkerQueueQueueNewQueueResponseResult]
type accountWorkerQueueQueueNewQueueResponseResultJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	QueueID     apijson.Field
	QueueName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueQueueNewQueueResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueQueueNewQueueResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       accountWorkerQueueQueueNewQueueResponseResultInfoJSON `json:"-"`
}

// accountWorkerQueueQueueNewQueueResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountWorkerQueueQueueNewQueueResponseResultInfo]
type accountWorkerQueueQueueNewQueueResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueQueueNewQueueResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerQueueQueueNewQueueResponseSuccess bool

const (
	AccountWorkerQueueQueueNewQueueResponseSuccessTrue AccountWorkerQueueQueueNewQueueResponseSuccess = true
)

type AccountWorkerQueueUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountWorkerQueueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountWorkerQueueQueueNewQueueParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountWorkerQueueQueueNewQueueParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
