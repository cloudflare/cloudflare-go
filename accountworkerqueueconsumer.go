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

// AccountWorkerQueueConsumerService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountWorkerQueueConsumerService] method instead.
type AccountWorkerQueueConsumerService struct {
	Options []option.RequestOption
}

// NewAccountWorkerQueueConsumerService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerQueueConsumerService(opts ...option.RequestOption) (r *AccountWorkerQueueConsumerService) {
	r = &AccountWorkerQueueConsumerService{}
	r.Options = opts
	return
}

// Updates the consumer for a queue, or creates one if it does not exist.
func (r *AccountWorkerQueueConsumerService) Update(ctx context.Context, accountIdentifier string, name string, consumerName string, body AccountWorkerQueueConsumerUpdateParams, opts ...option.RequestOption) (res *AccountWorkerQueueConsumerUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/queues/%s/consumers/%s", accountIdentifier, name, consumerName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns the consumers for a queue.
func (r *AccountWorkerQueueConsumerService) List(ctx context.Context, accountIdentifier string, name string, opts ...option.RequestOption) (res *AccountWorkerQueueConsumerListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/queues/%s/consumers", accountIdentifier, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes the consumer for a queue.
func (r *AccountWorkerQueueConsumerService) Delete(ctx context.Context, accountIdentifier string, name string, consumerName string, opts ...option.RequestOption) (res *AccountWorkerQueueConsumerDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/queues/%s/consumers/%s", accountIdentifier, name, consumerName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new consumer for a queue.
func (r *AccountWorkerQueueConsumerService) QueueNewQueueConsumer(ctx context.Context, accountIdentifier string, name string, body AccountWorkerQueueConsumerQueueNewQueueConsumerParams, opts ...option.RequestOption) (res *AccountWorkerQueueConsumerQueueNewQueueConsumerResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/queues/%s/consumers", accountIdentifier, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountWorkerQueueConsumerUpdateResponse struct {
	Errors     []AccountWorkerQueueConsumerUpdateResponseError    `json:"errors"`
	Messages   []AccountWorkerQueueConsumerUpdateResponseMessage  `json:"messages"`
	Result     AccountWorkerQueueConsumerUpdateResponseResult     `json:"result"`
	ResultInfo AccountWorkerQueueConsumerUpdateResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountWorkerQueueConsumerUpdateResponseSuccess `json:"success"`
	JSON    accountWorkerQueueConsumerUpdateResponseJSON    `json:"-"`
}

// accountWorkerQueueConsumerUpdateResponseJSON contains the JSON metadata for the
// struct [AccountWorkerQueueConsumerUpdateResponse]
type accountWorkerQueueConsumerUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueConsumerUpdateResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountWorkerQueueConsumerUpdateResponseErrorJSON `json:"-"`
}

// accountWorkerQueueConsumerUpdateResponseErrorJSON contains the JSON metadata for
// the struct [AccountWorkerQueueConsumerUpdateResponseError]
type accountWorkerQueueConsumerUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueConsumerUpdateResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountWorkerQueueConsumerUpdateResponseMessageJSON `json:"-"`
}

// accountWorkerQueueConsumerUpdateResponseMessageJSON contains the JSON metadata
// for the struct [AccountWorkerQueueConsumerUpdateResponseMessage]
type accountWorkerQueueConsumerUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueConsumerUpdateResponseResult struct {
	CreatedOn       interface{}                                            `json:"created_on"`
	DeadLetterQueue interface{}                                            `json:"dead_letter_queue"`
	Environment     interface{}                                            `json:"environment"`
	QueueName       interface{}                                            `json:"queue_name"`
	ScriptName      interface{}                                            `json:"script_name"`
	Settings        AccountWorkerQueueConsumerUpdateResponseResultSettings `json:"settings"`
	JSON            accountWorkerQueueConsumerUpdateResponseResultJSON     `json:"-"`
}

// accountWorkerQueueConsumerUpdateResponseResultJSON contains the JSON metadata
// for the struct [AccountWorkerQueueConsumerUpdateResponseResult]
type accountWorkerQueueConsumerUpdateResponseResultJSON struct {
	CreatedOn       apijson.Field
	DeadLetterQueue apijson.Field
	Environment     apijson.Field
	QueueName       apijson.Field
	ScriptName      apijson.Field
	Settings        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueConsumerUpdateResponseResultSettings struct {
	BatchSize     float64                                                    `json:"batch_size"`
	MaxRetries    float64                                                    `json:"max_retries"`
	MaxWaitTimeMs float64                                                    `json:"max_wait_time_ms"`
	JSON          accountWorkerQueueConsumerUpdateResponseResultSettingsJSON `json:"-"`
}

// accountWorkerQueueConsumerUpdateResponseResultSettingsJSON contains the JSON
// metadata for the struct [AccountWorkerQueueConsumerUpdateResponseResultSettings]
type accountWorkerQueueConsumerUpdateResponseResultSettingsJSON struct {
	BatchSize     apijson.Field
	MaxRetries    apijson.Field
	MaxWaitTimeMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerUpdateResponseResultSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueConsumerUpdateResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       accountWorkerQueueConsumerUpdateResponseResultInfoJSON `json:"-"`
}

// accountWorkerQueueConsumerUpdateResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountWorkerQueueConsumerUpdateResponseResultInfo]
type accountWorkerQueueConsumerUpdateResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerUpdateResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerQueueConsumerUpdateResponseSuccess bool

const (
	AccountWorkerQueueConsumerUpdateResponseSuccessTrue AccountWorkerQueueConsumerUpdateResponseSuccess = true
)

type AccountWorkerQueueConsumerListResponse struct {
	Errors     []interface{}                                    `json:"errors,nullable"`
	Messages   []interface{}                                    `json:"messages,nullable"`
	Result     []AccountWorkerQueueConsumerListResponseResult   `json:"result"`
	ResultInfo AccountWorkerQueueConsumerListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountWorkerQueueConsumerListResponseSuccess `json:"success"`
	JSON    accountWorkerQueueConsumerListResponseJSON    `json:"-"`
}

// accountWorkerQueueConsumerListResponseJSON contains the JSON metadata for the
// struct [AccountWorkerQueueConsumerListResponse]
type accountWorkerQueueConsumerListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueConsumerListResponseResult struct {
	CreatedOn   interface{}                                          `json:"created_on"`
	Environment interface{}                                          `json:"environment"`
	QueueName   interface{}                                          `json:"queue_name"`
	Service     interface{}                                          `json:"service"`
	Settings    AccountWorkerQueueConsumerListResponseResultSettings `json:"settings"`
	JSON        accountWorkerQueueConsumerListResponseResultJSON     `json:"-"`
}

// accountWorkerQueueConsumerListResponseResultJSON contains the JSON metadata for
// the struct [AccountWorkerQueueConsumerListResponseResult]
type accountWorkerQueueConsumerListResponseResultJSON struct {
	CreatedOn   apijson.Field
	Environment apijson.Field
	QueueName   apijson.Field
	Service     apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueConsumerListResponseResultSettings struct {
	BatchSize     float64                                                  `json:"batch_size"`
	MaxRetries    float64                                                  `json:"max_retries"`
	MaxWaitTimeMs float64                                                  `json:"max_wait_time_ms"`
	JSON          accountWorkerQueueConsumerListResponseResultSettingsJSON `json:"-"`
}

// accountWorkerQueueConsumerListResponseResultSettingsJSON contains the JSON
// metadata for the struct [AccountWorkerQueueConsumerListResponseResultSettings]
type accountWorkerQueueConsumerListResponseResultSettingsJSON struct {
	BatchSize     apijson.Field
	MaxRetries    apijson.Field
	MaxWaitTimeMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerListResponseResultSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueConsumerListResponseResultInfo struct {
	Count      interface{}                                          `json:"count"`
	Page       interface{}                                          `json:"page"`
	PerPage    interface{}                                          `json:"per_page"`
	TotalCount interface{}                                          `json:"total_count"`
	TotalPages interface{}                                          `json:"total_pages"`
	JSON       accountWorkerQueueConsumerListResponseResultInfoJSON `json:"-"`
}

// accountWorkerQueueConsumerListResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountWorkerQueueConsumerListResponseResultInfo]
type accountWorkerQueueConsumerListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerQueueConsumerListResponseSuccess bool

const (
	AccountWorkerQueueConsumerListResponseSuccessTrue AccountWorkerQueueConsumerListResponseSuccess = true
)

type AccountWorkerQueueConsumerDeleteResponse struct {
	Errors     []AccountWorkerQueueConsumerDeleteResponseError    `json:"errors"`
	Messages   []AccountWorkerQueueConsumerDeleteResponseMessage  `json:"messages"`
	Result     []interface{}                                      `json:"result,nullable"`
	ResultInfo AccountWorkerQueueConsumerDeleteResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountWorkerQueueConsumerDeleteResponseSuccess `json:"success"`
	JSON    accountWorkerQueueConsumerDeleteResponseJSON    `json:"-"`
}

// accountWorkerQueueConsumerDeleteResponseJSON contains the JSON metadata for the
// struct [AccountWorkerQueueConsumerDeleteResponse]
type accountWorkerQueueConsumerDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueConsumerDeleteResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountWorkerQueueConsumerDeleteResponseErrorJSON `json:"-"`
}

// accountWorkerQueueConsumerDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountWorkerQueueConsumerDeleteResponseError]
type accountWorkerQueueConsumerDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueConsumerDeleteResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountWorkerQueueConsumerDeleteResponseMessageJSON `json:"-"`
}

// accountWorkerQueueConsumerDeleteResponseMessageJSON contains the JSON metadata
// for the struct [AccountWorkerQueueConsumerDeleteResponseMessage]
type accountWorkerQueueConsumerDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueConsumerDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       accountWorkerQueueConsumerDeleteResponseResultInfoJSON `json:"-"`
}

// accountWorkerQueueConsumerDeleteResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountWorkerQueueConsumerDeleteResponseResultInfo]
type accountWorkerQueueConsumerDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerQueueConsumerDeleteResponseSuccess bool

const (
	AccountWorkerQueueConsumerDeleteResponseSuccessTrue AccountWorkerQueueConsumerDeleteResponseSuccess = true
)

type AccountWorkerQueueConsumerQueueNewQueueConsumerResponse struct {
	Errors     []AccountWorkerQueueConsumerQueueNewQueueConsumerResponseError    `json:"errors"`
	Messages   []AccountWorkerQueueConsumerQueueNewQueueConsumerResponseMessage  `json:"messages"`
	Result     AccountWorkerQueueConsumerQueueNewQueueConsumerResponseResult     `json:"result"`
	ResultInfo AccountWorkerQueueConsumerQueueNewQueueConsumerResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountWorkerQueueConsumerQueueNewQueueConsumerResponseSuccess `json:"success"`
	JSON    accountWorkerQueueConsumerQueueNewQueueConsumerResponseJSON    `json:"-"`
}

// accountWorkerQueueConsumerQueueNewQueueConsumerResponseJSON contains the JSON
// metadata for the struct
// [AccountWorkerQueueConsumerQueueNewQueueConsumerResponse]
type accountWorkerQueueConsumerQueueNewQueueConsumerResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerQueueNewQueueConsumerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueConsumerQueueNewQueueConsumerResponseError struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    accountWorkerQueueConsumerQueueNewQueueConsumerResponseErrorJSON `json:"-"`
}

// accountWorkerQueueConsumerQueueNewQueueConsumerResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountWorkerQueueConsumerQueueNewQueueConsumerResponseError]
type accountWorkerQueueConsumerQueueNewQueueConsumerResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerQueueNewQueueConsumerResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueConsumerQueueNewQueueConsumerResponseMessage struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    accountWorkerQueueConsumerQueueNewQueueConsumerResponseMessageJSON `json:"-"`
}

// accountWorkerQueueConsumerQueueNewQueueConsumerResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountWorkerQueueConsumerQueueNewQueueConsumerResponseMessage]
type accountWorkerQueueConsumerQueueNewQueueConsumerResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerQueueNewQueueConsumerResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueConsumerQueueNewQueueConsumerResponseResult struct {
	CreatedOn       interface{}                                                           `json:"created_on"`
	DeadLetterQueue string                                                                `json:"dead_letter_queue"`
	Environment     interface{}                                                           `json:"environment"`
	QueueName       interface{}                                                           `json:"queue_name"`
	ScriptName      interface{}                                                           `json:"script_name"`
	Settings        AccountWorkerQueueConsumerQueueNewQueueConsumerResponseResultSettings `json:"settings"`
	JSON            accountWorkerQueueConsumerQueueNewQueueConsumerResponseResultJSON     `json:"-"`
}

// accountWorkerQueueConsumerQueueNewQueueConsumerResponseResultJSON contains the
// JSON metadata for the struct
// [AccountWorkerQueueConsumerQueueNewQueueConsumerResponseResult]
type accountWorkerQueueConsumerQueueNewQueueConsumerResponseResultJSON struct {
	CreatedOn       apijson.Field
	DeadLetterQueue apijson.Field
	Environment     apijson.Field
	QueueName       apijson.Field
	ScriptName      apijson.Field
	Settings        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerQueueNewQueueConsumerResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueConsumerQueueNewQueueConsumerResponseResultSettings struct {
	BatchSize     float64                                                                   `json:"batch_size"`
	MaxRetries    float64                                                                   `json:"max_retries"`
	MaxWaitTimeMs float64                                                                   `json:"max_wait_time_ms"`
	JSON          accountWorkerQueueConsumerQueueNewQueueConsumerResponseResultSettingsJSON `json:"-"`
}

// accountWorkerQueueConsumerQueueNewQueueConsumerResponseResultSettingsJSON
// contains the JSON metadata for the struct
// [AccountWorkerQueueConsumerQueueNewQueueConsumerResponseResultSettings]
type accountWorkerQueueConsumerQueueNewQueueConsumerResponseResultSettingsJSON struct {
	BatchSize     apijson.Field
	MaxRetries    apijson.Field
	MaxWaitTimeMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerQueueNewQueueConsumerResponseResultSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerQueueConsumerQueueNewQueueConsumerResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                               `json:"total_count"`
	JSON       accountWorkerQueueConsumerQueueNewQueueConsumerResponseResultInfoJSON `json:"-"`
}

// accountWorkerQueueConsumerQueueNewQueueConsumerResponseResultInfoJSON contains
// the JSON metadata for the struct
// [AccountWorkerQueueConsumerQueueNewQueueConsumerResponseResultInfo]
type accountWorkerQueueConsumerQueueNewQueueConsumerResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerQueueConsumerQueueNewQueueConsumerResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerQueueConsumerQueueNewQueueConsumerResponseSuccess bool

const (
	AccountWorkerQueueConsumerQueueNewQueueConsumerResponseSuccessTrue AccountWorkerQueueConsumerQueueNewQueueConsumerResponseSuccess = true
)

type AccountWorkerQueueConsumerUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountWorkerQueueConsumerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountWorkerQueueConsumerQueueNewQueueConsumerParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountWorkerQueueConsumerQueueNewQueueConsumerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
